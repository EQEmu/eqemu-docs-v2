# Developer Blog - Server Optimizations

**Author(s)** Akkadius

Several times the topic of server size optimizations and the problems we had in the past have come up.

I wanted to take a moment to discuss the close mob list optimizations we've made in the past year and how it has impacted the server.

I'll also use this to discuss some of the other optimizations we've made over the years as well.

| Before                                                                                                 | After                                                                                                  |
|--------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------|
| ![ezgif-4-3bc7d388a2](https://github.com/user-attachments/assets/d37ce5b9-3112-4476-9b42-ffd306a09589) | ![ezgif-4-9c62a2f608](https://github.com/user-attachments/assets/7f8b5e5c-b81d-4c53-acfb-3f4ec4584f1b) |
| ![ezgif-4-3bc7d388a2](https://github.com/user-attachments/assets/7cb3210b-8259-4f3b-84d3-55c096a21c1a) | ![ezgif-4-9c62a2f608](https://github.com/user-attachments/assets/90056450-8a19-4cc8-b297-f41a5dfc03e9) |

Whenever the system was fully rolled out, it dropped CPU utilization roughly 80% and network send dropped dramatically
as well. I'll break it down in the sections below.

## Problem 1) Entity List Looping

EQEmu had a bit of a scaling issue when it came to number of entities in a zone. Everything every entity does used to
get sent to every single entity regardless of distance. Core functions would calculate against every entity regardless
of distance. Resulting in a tremendous amount of wasted CPU cycles, overhead and computation.

This adds up tremendously as many parts of the code are looping over the entity list sometimes many hundreds or
thousands of times a second performing expensive checks over (N) NPC's when all you care about is potentially 1-10
entities closest to you.

Imagine you have a zone with 500-1000 entities between clients and NPCs. Here's what's happening against all of those
entities as we loop through the entire entity list. The below list are examples and does not contain everything.

* **NPC to NPC Aggro Scanning** - NPC aggro scanning would hit all entities and check multiple times per second per entity. This
  goes for client aggro scanning and NPC aggro scanning.
* **Client to MPC Aggro Scanning** - Clients would multiple times per second, per client loop through the entire entity list to determine whether or not they should aggro an NPC.
* **Auras** - When auras are being processed, they loop all entities, performing distance checks before determining
* **Animations** - We would send animation packets to every client in the zone. When all that matters are clients closest to us.
* **AE Spells** - When AE spells of any sort are processed, rains, waves, the entire entity list is being looped (600
  NPCs) versus maybe the 4 relevant entities around you and distance checks done for every entity.
* **AE Taunt** - Lots of special ability checks, distance checks, LOS
* **Rampage** - Checking aggro, NPC, entity type.
* **Say** - When an NPC says anything, it would otherwise loop the entire entity list to calculate distance.
* **NPC Guard Assist** - Every NPC scans every single NPC for assisting during combat, performing expensive LOS checks,
  faction, distance checks before determining if an NPC should assist.
* **NPC Yell for Help** - (AIYellForHelp) Similar to assist, NPC scans list for distance, checking assist range to
  determine if NPC can help another NPC. These checks can be intensely made many times a second.
* **NPC Beneficial Cast Checks** - (AICheckCloseBeneficialSpells) When NPC's are checking to cast beneficial spells on
  other NPC's, buffs, heals, etc. They are looping through the entire entity list, performing distance checks, LOS,
  reverse faction, spell checks before ultimately determining if the NPC should cast the spell. This happens multiple
  times per second per entity.
* **QueueCloseClients** - Packet sending function used almost everywhere, would otherwise loop through the entire entity
  list and calculate distance before sending the packet. For example this is used in the following packets - damage
  update, death, client update, emote, yell for help, say, open door, send level appearance, spell effect, anim, spell
  anim, spell interrupt, begin cast and more.
* **Position Updates** - Position updates are sent zone wide to every client regardless of distance as NPC's move across
  the map, as players move across the map. These can be sent far more conservatively - only sending position updates zone wide when we really need to.

After thinking through all of this, it can be easy to see where all of this adds up, you can start to see where all of these interactions add up tremendously in overhead and wasted CPU cycles.

### Solution - Close Mob Lists

**Why don't we just act against the closest relevant NPCs?**

Instead of always looping through the entity list excessively and tens of thousands of times a second in some cases, how
can we reduce that cost by upwards of 100x?

We keep a list of close mobs relevant to each mob (client, npc, bot etc.) and that's achieved through dynamic scanning.

Below is a visual representation of what used to happen in the codebase relative to each entity versus how things behave today. Almost everything in the code that involved checking nearby NPC's would loop through all entities and perform expensive calculations. Today we just use the close mob list which reduces a ton of computational overhead.

| Before                                                                                                 | After                                                                                                  |
|--------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------|
| ![ezgif-4-3bc7d388a2](https://github.com/user-attachments/assets/7cb3210b-8259-4f3b-84d3-55c096a21c1a) | ![ezgif-4-9c62a2f608](https://github.com/user-attachments/assets/90056450-8a19-4cc8-b297-f41a5dfc03e9) |

#### Scanning

Let's dive into how the mechanics work.

Each mob has a list called `m_close_mobs` and utilize the following method in the source to scan.

```cpp  
void EntityList::ScanCloseMobs(Mob *scanning_mob)  
```  

The way scanning works is say you have 1000 entities in a zone and a client is building its close mobs list. The default
scanning range is defined by `Range:MobCloseScanDistance` and it is 600 units. Anything within that scan range gets put
within the `close_mobs` mob variable list.

The mob will also add itself to other mobs close lists as it becomes in range of another mob, this is to ensure accuracy
of the close list of every mob. Most of the time, entities have less than ~10-50 on their close list relevant to them.

Checks are made to ensure we don't add the same mob twice in any scenario.

##### Relevant Code

Here are the relevant mob member variables

```cpp
std::unordered_map<uint16, Mob *> m_close_mobs;
Timer                             m_scan_close_mobs_timer;
Timer                             m_mob_check_moving_timer;
```

Relevant methods

```cpp
Mob::CheckScanCloseMobsMovingTimer();
Mob::ScanCloseMobProcess();
Mob::GetCloseMobList(float distance = 0);
EntityList::GetCloseMobList(Mob *mob, float distance);
```

**GetCloseMobList** usages

Here is the reformatted list showing just the file and the function:

- **attack.cpp**
    - Attack
- **effects.cpp**
    - AEAttack
    - AESpell
    - AETaunt
    - MassGroupBuff
- **entity.cpp**
    - GetTargetsForVirusEffect
    - QuestJournalledSayClose
    - QueueCloseClients
- **mob.cpp**
    - Say
- **npc.cpp**
    - AICheckCloseBeneficialSpells
    - AIYellForHelp
- **spells.cpp**
    - SpellFinished
- **aura.cpp**
    - ProcessEnterTrap
    - ProcessExitTrap
    - ProcessOnAllFriendlies
    - ProcessOnAllGroupMembers
    - ProcessOnGroupMembersPets
    - ProcessSpawns
    - ProcessTotem
- **client.cpp**
    - ClientToNpcAggroProcess
- **mob.cpp**
    - IsCloseToBanker
    - SetBottomRampageList
    - SetTopRampageList
- **npc.cpp**
    - DoNpcToNpcAggroScan

#### Initial Scan

When an entity is created, we need to call the function - initial scan is invoked differently depending on the entity.

| **Entity Type** | **Where**               |  
|-----------------|-------------------------|  
| Client          | Client::CompleteConnect |  
| NPC             | EntityList::AddNPC      |  

#### Ongoing Scan

During the duration of an entities life, there is an ongoing scan based on a dynamic timer. This timer is also different
depending on whether an NPC is idle or moving.

| **Entity Type** | **Where**       |  
|-----------------|-----------------|  
| Client          | Client::Process |  
| NPC             | NPC::Process    |  

```cpp
Mob::ScanCloseMobProcess()
```

For both client / npc the code is called in each entities respective Mob/Client::Process() branches.

#### Scanning Dynamic Interval Timer

Scanning frequency changes depending on whether an entity is moving or idle. This is to keep computational expenses very
low.

This is determined by `Mob::mob_close_scan_timer` and calculted during invocation of `Mob::CheckScanCloseMobsMovingTimer`

| **Entity Type** | **Idle** | **Moving**      |  
|-----------------|----------|-----------------|  
| Mob             | 1 minute | 6 seconds (tic) |  

For NPC's we have a `Mob::m_mob_check_moving_timer` which is throttled to once a second to keep computational checks light
to check if an NPC is moving or not to change the scan timer.

#### Book Keeping

In order to ensure we don't have dangling pointers, crashes we have to clean this close mob list up when entities die.

In the mob destructor we cleanup

```cpp
Mob::~Mob() {
  // ... destructor code
  entity_list.RemoveMobFromCloseLists(this);
  m_close_mobs.clear();
```

#### **What if an NPC is beyond the scan range?**

Simple. Every time `Mob::GetCloseMobList()` is called, by default it will return the close mob list if you pass in a distance less than the specified rule value `Range:MobCloseScanDistance`. If your code passes in a distance greater than this scan range, it will default to using the entire entity list. This solves for edge cases where you have a zone wide spell, zone wide AOE, aggro, taunt or otherwise.

## Problem 2) Zone Wide Position Updates

Our server code used to send position updates for all NPC's and clients to the player zone wide. This resulted in excess of packet sending and waste of CPU overhead.

We now only send position updates to entities close to clients and do full zone updates under certain circumstances.

### Simple Solution - Only Send Position Updates to Close Mobs

While the simple answer is to only send position updates to close mobs, there's a few things we have to do to ensure we don't fall out of sync under certain edge cases.

### Solution - Bulk Send Position Updates

When we enter a zone we get a full send of the entity list of spawn locations, appearance etc.

After that initial send we only get sent zone wide updates under specific circumstances.

This is handled by `Client::CheckSendBulkClientPositionUpdate()` after a client has traveled enough distance to warrant a full zone sync all mobs (Clients and NPCs).

For clients, the client requires a position update roughly every 10 seconds, so if a client is out of range we just send one every 10 seconds.

**Client::Process**

```cpp
if (!IsMoving() && m_position_update_timer.Check()) {
  SentPositionPacket(0.0f, 0.0f, 0.0f, 0.0f, 0);
}
```

Here are the ranges specified in our movement updates

Here is the reformatted list, including the range mentioned in the function signature:

- **client.cpp**
    - CheckSendBulkClientPositionUpdate - `ClientRangeAny`
- **client_packet.cpp**
    - Handle_OP_ClientUpdate - (Not specified in the usage snippet)
- **mob.cpp**
    - GMMove (1st usage) - `ClientRangeAny`
    - GMMove (2nd usage) - `ClientRangeAny`
- **mob_movement_manager.cpp**
    - StopMovingCommand - `ClientRangeCloseMedium`
    - MoveToCommand (1st usage) - `ClientRangeCloseMedium`
    - MoveToCommand (2nd usage) - `ClientRangeCloseMedium`
    - MoveToCommand (3rd usage) - `ClientRangeCloseMedium`
    - EvadeCombatCommand - `ClientRangeCloseMedium`
    - TeleportToCommand - `ClientRangeAny`
    - SwimToCommand (1st usage) - `ClientRangeCloseMedium`
    - SwimToCommand (2nd usage) - `ClientRangeCloseMedium`
    - SwimToCommand (3rd usage) - `ClientRangeCloseMedium`
    - FlyToCommand (1st usage) - `ClientRangeCloseMedium`
    - FlyToCommand (2nd usage) - `ClientRangeCloseMedium`
    - FlyToCommand (3rd usage) - `ClientRangeCloseMedium`
    - RotateToCommand (1st usage) - `ClientRangeCloseMedium`
    - RotateToCommand (2nd usage) - `ClientRangeCloseMedium`
- **spells.cpp**
    - Spin - `ClientRangeAny`
- **movement.cpp**
    - command_movement - `ClientRangeAny`
- **mob_movement_manager.h**
    - SendCommandToClients (Declaration) - (Range not specified in this snippet)

#### Problem 3) Sending Packet Messages to Relevant Distances

This one was done long before close mob lists, but we went in and implemented sending updates by range and implemented them as configurable rules in the source.

This made cost savings but not as dramatic as the close mob list optimizations themselves. They compliment the close mob changes because the packet send functions can utilize the same list instead of looping the whole entity list.

```cpp
RULE_INT(Range, Say, 15, "The range that is required before /say or hail messages will work to an NPC");
RULE_INT(Range, Emote, 135, "The packet range in which emote messages are sent'");
RULE_INT(Range, BeginCast, 200, "The packet range in which begin cast messages are sent");
RULE_INT(Range, Anims, 135, "The packet range in which begin cast messages are sent");
RULE_INT(Range, SpellParticles, 135, "The packet range in which spell particles are sent");
RULE_INT(Range, DamageMessages, 50, "The packet range in which damage messages are sent (non-crit)");
RULE_INT(Range, SpellMessages, 75, "The packet range in which spell damage messages are sent");
RULE_INT(Range, SongMessages, 75, "The packet range in which song messages are sent");
RULE_INT(Range, StunMessages, 75, "The packet range in which stun messages are sent");
RULE_INT(Range, ClientPositionUpdates, 300, "Distance in which the own changed position is communicated to other clients");
RULE_INT(Range, CriticalDamage, 80, "The packet range in which critical hit messages are sent");
RULE_INT(Range, MobCloseScanDistance, 600, "Close scan distance");
RULE_INT(Range, MaxDistanceToClickDoors, 100, "Max distance that a client can click a door from (Client says 'You can't reach that' at roughly ;25-50 for most doors)");
```

