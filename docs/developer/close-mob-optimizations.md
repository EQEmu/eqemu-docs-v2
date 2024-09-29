I've been asked about close mob optimizations a few times so I finally wanted to spend some time to clearly break it all down and explain to folks passing in conversations the problem, solutions and pieces around it.

| Before | After |
|--|--|
| ![ezgif-4-3bc7d388a2](https://github.com/user-attachments/assets/d37ce5b9-3112-4476-9b42-ffd306a09589) | ![ezgif-4-9c62a2f608](https://github.com/user-attachments/assets/7f8b5e5c-b81d-4c53-acfb-3f4ec4584f1b) |

Whenever the system was fully rolled out, it dropped CPU utilization roughly 80% and network send dropped dramatically
as well. I'll break it down in the sections below.

# The Problem

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

# Solution(s)

Now you can start to see where all of these interactions add up tremendously in overhead and wasted CPU cycles.

**Why don't we just act against the closest relevant NPCs?**

Instead of always looping through the entity list excessively and tens of thousands of times a second in some cases, how
can we reduce that cost by upwards of 100x?

We keep a list of close mobs relevant to each mob (client, npc, bot etc.) and that's achieved through dynamic scanning.  

| Before | After |
|--|--|
| ![ezgif-4-3bc7d388a2](https://github.com/user-attachments/assets/7cb3210b-8259-4f3b-84d3-55c096a21c1a) | ![ezgif-4-9c62a2f608](https://github.com/user-attachments/assets/90056450-8a19-4cc8-b297-f41a5dfc03e9) |

## Scanning

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

```
Mob::ScanCloseMobProcess()
```

For both client / npc the code is called in each entities respective Mob/Client::Process() branches.

```cpp  
Mob::ScanCloseMobProcess()  
```  

### Scanning Dynamic Interval Timer

Scanning frequency changes depending on whether an entity is moving or idle. This is to keep computational expenses very
low. This is determined by `Client::mob_close_scan_timer` and `Mob::mob_close_scan_timer`

| **Entity Type** | **Idle** | **Moving**      |  
|-----------------|----------|-----------------|  
| Client          | 1 minute | 6 seconds (tic) |  
| NPC             | 1 minute | 6 seconds (tic) |  

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


