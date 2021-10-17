---
description: This page describes the faction system on your EQEmu Server.
---

# Factions

The faction system is pretty simply really, but it can also be quite confusing if you don't understand the purpose of some of the tables, or you don't understand how all the tables interact. These pages should help you get a high level understanding of the faction system, the database tables related to faction functionality, and how to use the system to get the results you expect.

## NPC Aggro

It is important to understand that NPCs aggro Players due to faction differences.  These differences can be caused by racial / class / deity mods that are inherit in the player's character choice, or they can be caused by the player losing faction standing with a particular faction.

Read more about the NPC Aggro settings in the [NPC Aggro section](../npc/npc-aggro.md).

## Database Tables

### faction\_list

The [faction\_list ](https://eqemu.gitbook.io/database-schema/categories/factions/faction_list)table contains one entry for each unique faction on the server. This list was previous our own values, but now uses values directly from the client for any non-custom factions. The ID field identifies an in-game faction, but IS NOT directly associated with an NPC. This list simply lists the in game factions and the starting value that all characters start at with this faction before they begin earning or losing faction. DO NOT use these values in npc\_types. See npc\_faction below for details.

### faction\_base\_data

The [faction\_base\_data](https://eqemu.gitbook.io/database-schema/categories/factions/faction_base_data) was imported directly from the client. In includes the lowest a character can go in "earned" faction and the highest amount of "earned" faction they can obtain. The modifiers for race, class and deity are still applied on top of the personal faction to determine the final standing.

### faction\_list\_mod

The [faction\_list\_mod](https://eqemu.gitbook.io/database-schema/categories/factions/faction_list_mod) table stores the static data that indicates what bonuses or penalties race, class and deity have on various factions. For example, a character's standing may be worse with some factions due to that faction hating your deity. Or like in the case of the Kerran race, they dislike Erudites. These adjustments, combined with the starting value for everyone taken from faction\_list, make up where you start in relation to all of the factions on Norrath.

### faction\_values

The [faction\_values](https://eqemu.gitbook.io/database-schema/categories/factions/faction_values) table stores the ever-changing data in regards to each characters actions in the game. A running sum of all of the hit and gains in faction for characters is stored here, one line per char/faction pair. Lines only exist if a character has at least one hit or credit based on in game actions.

These two tables are more about assigning each NPC a primary faction, while also establishing sister factions \(assist worthy\), enemy factions \(attack worthy\), and the hit taken by any PC that were to kill the PC.

### npc\_faction

The ID from the [npc\_faction](https://eqemu.gitbook.io/database-schema/categories/npcs/npc_faction) table is where an NPC gets its entire faction behavior. The ID assigned here, ties an NPC to their primary faction \(from faction\_list\) as well as determines whether not not they will assist their primary faction comrades in battle.

But more than this, it serves as a key into the npc\_faction\_entries table.

### npc\_faction\_entries

The [npc\_faction\_entries](https://eqemu.gitbook.io/database-schema/categories/npcs/npc_faction_entries) table has three major roles. First, entries are placed here that indicate how much faction is gained/lost with various factions if any NPC has this npc\_faction\_id is killed. The second, is that you can list how these NPCS react to other NPCS of various other factions. The third, these entries can also indicate faction gain/loss via the task system. The npc\_faction\_id is specified in the faction\_reward field in the tasks table.

This table collects too much data in one place, and really should be broken down. As it stands right now, if there are 10 different mobs that yield different faction hits when killed, there needs to be 10 different npc\_factions and copies of their behavior records in new npc\_faction\_entries records. If we separated out the hits into a separate table, then we would only need one copy of the behavior records. We'll mark this as potential improvement. \(In my work matching live faction data, it has required a ton of redundancy\).

