# Developer Blog - Server Optimizations

!!! info

    **Author(s)** Akkadius

    Several times the topic of server size optimizations and the problems we had in the past have come up.

    I wanted to take a moment to write about some of the optimizations we've made to the server codebase to reduce CPU overhead, network send, database calls and improve performance.

    It may not be exhaustive, but it's a good start to understanding some of the notable optimizations we've made over time so we have something to link and reference.

## Optimization - Mob Close Lists

!!! info

    **Year** 2019
  
    Initial PR - https://github.com/EQEmu/Server/pull/940

Many PR's and iterations later, we had quite a few edge cases and bugs to work out. It's been rock solid for a few years now.

| Before                                                                                                 | After                                                                                                  |
|--------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------|
| ![ezgif-4-3bc7d388a2](https://github.com/user-attachments/assets/d37ce5b9-3112-4476-9b42-ffd306a09589) | ![ezgif-4-9c62a2f608](https://github.com/user-attachments/assets/7f8b5e5c-b81d-4c53-acfb-3f4ec4584f1b) |
| ![ezgif-4-3bc7d388a2](https://github.com/user-attachments/assets/7cb3210b-8259-4f3b-84d3-55c096a21c1a) | ![ezgif-4-9c62a2f608](https://github.com/user-attachments/assets/90056450-8a19-4cc8-b297-f41a5dfc03e9) |

Whenever the system was fully rolled out, it dropped CPU utilization roughly 80% and network send dropped dramatically
as well. I'll break it down in the sections below.

EQEmu had a bit of a scaling issue when it came to number of entities in a zone. Everything every entity does used to
get sent to every single entity regardless of distance. Core functions would calculate against every entity regardless
of distance. Resulting in a tremendous amount of wasted CPU cycles, overhead and computation.

This adds up tremendously as many parts of the code are looping over the entity list sometimes many hundreds or
thousands of times a second performing expensive checks over (N) NPC's when all you care about is potentially 1-10
entities closest to you.

Imagine you have a zone with 500-1000 entities between clients and NPCs. Here's what's happening against all of those
entities as we loop through the entire entity list. The below list are examples and does not contain everything.

!!! info
  
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

### Close Mob Lists

**Why don't we just act against the closest relevant mobs?**

Instead of always looping through the entity list excessively and tens of thousands of times a second in some cases, how
can we reduce that cost by upwards of 100x?

We keep a list of close mobs relevant to each mob (client, npc, bot etc.) and that's achieved through dynamic scanning.

Below is a visual representation of what used to happen in the codebase relative to each entity versus how things behave today. Almost everything in the code that involved checking nearby NPC's would loop through all entities and perform expensive calculations. Today we just use the close mob list which reduces a ton of computational overhead.

| Before                                                                                                 | After                                                                                                  |
|--------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------|
| ![ezgif-4-3bc7d388a2](https://github.com/user-attachments/assets/7cb3210b-8259-4f3b-84d3-55c096a21c1a) | ![ezgif-4-9c62a2f608](https://github.com/user-attachments/assets/90056450-8a19-4cc8-b297-f41a5dfc03e9) |

### Scanning

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

### Relevant Code

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

!!! info

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

### Initial Scan

When an entity is created, we need to call the function - initial scan is invoked differently depending on the entity.

| **Entity Type** | **Where**               |  
|-----------------|-------------------------|  
| Client          | Client::CompleteConnect |  
| NPC             | EntityList::AddNPC      |  

### Ongoing Scan

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

### Scanning Dynamic Interval Timer

Scanning frequency changes depending on whether an entity is moving or idle. This is to keep computational expenses very
low.

This is determined by `Mob::mob_close_scan_timer` and calculted during invocation of `Mob::CheckScanCloseMobsMovingTimer`

| **Entity Type** | **Idle** | **Moving**      |  
|-----------------|----------|-----------------|  
| Mob             | 1 minute | 6 seconds (tic) |  

For NPC's we have a `Mob::m_mob_check_moving_timer` which is throttled to once a second to keep computational checks light
to check if an NPC is moving or not to change the scan timer.

### Book Keeping

In order to ensure we don't have dangling pointers, crashes we have to clean this close mob list up when entities die.

In the mob destructor we cleanup

```cpp
Mob::~Mob() {
  // ... destructor code
  entity_list.RemoveMobFromCloseLists(this);
  m_close_mobs.clear();
```

### **What if an NPC is beyond the scan range?**

Simple. Every time `Mob::GetCloseMobList()` is called, by default it will return the close mob list if you pass in a distance less than the specified rule value `Range:MobCloseScanDistance`. If your code passes in a distance greater than this scan range, it will default to using the entire entity list. This solves for edge cases where you have a zone wide spell, zone wide AOE, aggro, taunt or otherwise.

## Optimization - Zone Wide Position Updates

!!! info

    **Year** 2022

Our server code used to send position updates for all NPC's and clients to the player zone wide. This resulted in excess of packet sending and waste of CPU overhead.

We now only send position updates to entities close to clients and do full zone updates under certain circumstances.

### Only Send Position Updates to Close Mobs

While the simple answer is to only send position updates to close mobs, there's a few things we have to do to ensure we don't fall out of sync under certain edge cases.

### Bulk Send Position Updates

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

!!! info
  
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

## Optimization - Sending Packet Messages to Relevant Distances

!!! info

    **Year** 2017
    
    Relevant commit - https://github.com/EQEmu/Server/commit/14d09485eb1fda95982eba7ebf48207729b394a2

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

## Optimization - Send Animation Packets Less

!!! info

    **Year** 2017
    
    Related commit https://github.com/EQEmu/Server/commit/127f51e7587b0d354f0f326f8661a640baf313e2

Notes from original commit

!!! quote

	- HP Updates now only send to others when HP percentage changes (0-100%)
		- HP Updates were sending excessively even during idle zones when HP wasn't changing at all
	- Attack animations now only send once per second versus up to a hundred times a second per Mob/Client
	- 17,000 OP_ClientUpdate packets per second have been observed in combat scenarios, some of the major culprits have been
		throttled without affecting what the client should see
	- Before and After packet differences under similar load/tests (Packets per second)
		- 7,000 - 8,000 	OP_Animation pps	After: 600-800 pps
		- 13,0000 - 17,000 	OP_MobHealth pps	After: 1-10 pps
		- 15,0000 - 20,000 	OP_ClientUpdate pps	After: 500-1,000 pps
	- Packet reports from a 46 client test here:
		https://gist.github.com/Akkadius/28b7ad2fdd82bdd15ea737c68f404346
	- Servers who use Marquee HP updates will also recieve far less packet spam as they will only be sent when HP changes


!!! info

    We've made changes around this code since so you will need to look at the source to see the latest record of how these changes are implemented.

## Optimization - Perl Heavy Exports

!!! info

    **Year** 2015
    
    Related commit - https://github.com/EQEmu/Server/commit/e8d18cb014fc2123518056fe5dbf9d8f17360da6

Commit notes

!!! quote

    - Added Rate limit the rate in which signals are processed for NPC's (.5 seconds instead of .01 seconds)
      Added Perl Export Settings which should heavily reduce the Perl footprint
    - Normally when any sub EVENT_ gets triggered, all kinds of variables have to get exported every single time an event is triggered and
      this can make Perl very slow when events are triggered constantly
      - The two most taxing variable exports are the item variables ($itemcount{} $hasitem{} $oncursor{}) and qglobals ($qglobals{})
      - qglobals can pose to be an issue quickly when global qglobals build up, it is highly recommend to use the GetGlobal() and SetGlobal()
      methods instead as they don't reference the hashmap $qglobals{} that is rebuilt every single time a sub event is triggered
    - A stress test conducted with 10,000 samples shows an excess of time taken to export variables: http://i.imgur.com/NEpW1tS.png
    - After the Perl Export Settings table is implemented, and all exports are shut off you see the following test result:
      http://i.imgur.com/Du5hth9.png
    - The difference of eliminating uneeded exports brings the overhead and footprint of 10,000 triggers from 54 seconds to 2 seconds
    - In a 10,000 sample test (10,000 sub event triggers), exporting item variables adds 12 seconds alone, when item variables are only needed in
      EVENT_ITEM and EVENT_SAY a majority of the time if at all
    - In a 10,000 sample test (10,000 sub event triggers), exporting qglobals with approximately 1,000 global qglobals in the database creates
      about 11-20 seconds of delay on its own (Depending on hardware of course)
    - I've written a parser that has determined which of these exports are needed in which sub routines and have turned off all of the unneeded
      exports in sub routines that do not need them and used it to create the default table that will be installed in the database.
    - The export table is called 'perl_event_export_settings' and it resembles the following structure and contains all current 81 EVENTS
      - If an entry doesn't exist in this table and a new subroutine is added to the source, all exports will be on by default for that routine
  
    ```
    +----------+-----------------------------------------+-----------------+------------+-------------+-------------+--------------+
    | event_id | event_description                       | export_qglobals | export_mob | export_zone | export_item | export_event |
    +----------+-----------------------------------------+-----------------+------------+-------------+-------------+--------------+
    |        0 | EVENT_SAY                               |               1 |          1 |           1 |           1 |            1 |
    |        1 | EVENT_ITEM                              |               1 |          1 |           1 |           0 |            1 |
    |        2 | EVENT_DEATH                             |               1 |          1 |           1 |           0 |            1 |
    |        3 | EVENT_SPAWN                             |               1 |          1 |           1 |           0 |            1 |
    |        4 | EVENT_ATTACK                            |               0 |          1 |           1 |           0 |            1 |
    |        5 | EVENT_COMBAT                            |               1 |          1 |           1 |           0 |            1 |
    +----------+-----------------------------------------+-----------------+------------+-------------+-------------+--------------+
    ```
  
    - If a change is made to this table while the server is live and running, you can hot reload all zone process settings via:
      #reloadperlexportsettings
      - For those who wonder what "exports" are, they are reference to variables that are made available at runtime of the sub event, such as:
        (export_qglobals) (Heavy) : $qglobals https://github.com/EQEmu/Server/blob/master/zone/embparser.cpp#L916
        (export_item) (Heavy) : $itemcount{} $hasitem{} $oncursor{} https://github.com/EQEmu/Server/blob/master/zone/embparser.cpp#L1103
        (export_zone) : $zoneid, $instanceid, $zoneln etc. https://github.com/EQEmu/Server/blob/master/zone/embparser.cpp#L1083
        (export_mob) : $x, $y, $z, $h, $hpratio etc. https://github.com/EQEmu/Server/blob/master/zone/embparser.cpp#L1032
        (export_event) : (event specific) IE: EVENT_SAY ($text) https://github.com/EQEmu/Server/blob/master/zone/embparser.cpp#L1141

## Optimization - Logging Functions moved to Macros

!!! info

    **Year** 2017
  
    Related commit - https://github.com/EQEmu/Server/commit/7aa1d243b0483ad9041537aada44f923bf923390

Original notes

!!! quote

    Reworked how all log calls are made in the source
    
    - Before we used Log.Out, we will now use a macro Log(
      - Before: Log.Out(Logs::General, Logs::Status, "Importing Spells...");
      - After: Log(Logs::General, Logs::Status, "Importing Spells...");
      - The difference is
        1) It's 200-300x faster especially when log statements are inside very hot code paths. We already
           had most hot paths checked before we logged them, but this blankets all existing logging calls now and not just the
           select few we had picked out in the source.
        2) Strings don't get copied to the stack, popped and pushed constantly even when we hit a log statement that
           actually isn't going to log anything.
           - We do an 'if (LogSys.log_settings[log_category].is_category_enabled == 1)' before we call a log function
           in the log macro so the log function doesn't get called at all if we're not logging the category
      - This has increased binary executables roughly 15KB
      - The old extern for EQEmuLogSys is now named LogSys appropriately instead of Log (Ex: LogSys.StartFileLogs())
      - The result keeps logging footprint non-existent for when we're not logging that category

## All Other Performance Changes

For anything not explicitly mentioned in the blog. Search "Performance" or "Optimization" in https://github.com/EQEmu/Server/blob/master/CHANGELOG.md for more