---
tags:
  - Logs
  - Logging
---


# Logging Categories

Logging categories have their settings stored in the [logsys_categories](https://docs.eqemu.io/schema/admin/logsys_categories/) table, but the categories are defined in [common/eqemu_logsys.h](https://github.com/EQEmu/Server/blob/master/common/eqemu_logsys.h)

!!! info
      The default settings for the "log-to" categories (Console, File, GM Say) are listed below.  You can modify these values for your own server in your `logsys_categories` table.
      The last column in this table (`discord_webhook_id`) will output to that webhook `id` in the `discord_webhooks` table.

| ID | Description | Console | File | GMSay | WebhookID | Notes |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 1 | AA | 0 | 0 | 0 | 0 | Use of AA actions (IE activated, cast) | 
| 2 | AI | 0 | 0 | 0 | 0 | Mob aggro, prox and range | 
| 3 | Aggro | 0 | 0 | 0 | 0 | Mob aggro, prox and range | 
| 4 | Attack | 0 | 0 | 0 | 0 | Currently only shows aggro count on mob death | 
| 6 | Combat | 0 | 0 | 0 | 0 | Combat attacks/misses/damage (IE weapon, base, thrown) | 
| 7 | Commands | 0 | 0 | 0 | 0 | GM and Bot Command loading and usage details | 
| 8 | Crash | 1 | 1 | 0 | 0 | Stacktrace, exceptions, timestamp | 
| 9 | Debug | 0 | 0 | 0 | 0 | Logs failures and errors (IE loading assets, login) and other miscellaneous details | 
| 10 | Doors | 0 | 0 | 0 | 0 | Use of doors (IE click, open, trigger) | 
| 11 | Error | 1 | 0 | 0 | 0 | Logs failures and errors (IE loading assets, DB connection) | 
| 12 | Guilds | 0 | 0 | 0 | 0 | Guild activities (IE update MOTD, DB queries) | 
| 13 | Inventory | 0 | 0 | 0 | 0 | Activities with inventory (IE items, slots, shared inventory) | 
| 14 | Launcher | 0 | 0 | 0 | 0 | World / Zone launch requests | 
| 15 | Netcode | 0 | 0 | 0 | 0 | Logs opcode issues (IE login, world connection) | 
| 17 | Object | 0 | 0 | 0 | 0 | Logs opcode details and issues (IE login, world connection) | 
| 18 | Pathing | 0 | 0 | 0 | 0 | Pathing waypoint logging (IE arrive, pause) | 
| 20 | Quests | 0 | 0 | 0 | 0 | Outputs quest errors to log | 
| 21 | Rules | 0 | 0 | 0 | 0 | Logs the rule system (IE load rules) | 
| 22 | Skills | 0 | 0 | 0 | 0 | Logs skill messages (IE failed combine) | 
| 23 | Spawns | 0 | 0 | 0 | 0 | All about spawns (IE ground, spawn2, conditions) | 
| 24 | Spells | 0 | 0 | 0 | 0 | All about spells (IE stacking checks, spell lands, failures, etc) |  
| 27 | Tasks | 0 | 0 | 0 | 0 | Task information | 
| 28 | Tradeskills | 0 | 0 | 0 | 0 | Tradeskill information (success, failure, modifiers) | 
| 29 | Trading | 0 | 0 | 0 | 0 | Trader info (bazaar, barter) | 
| 30 | Tribute | 0 | 0 | 0 | 0 | Tribute logging | 
| 35 | QueryErr | 1 | 0 | 0 | 0 | MySQL query errors | 
| 36 | Query | 0 | 0 | 0 | 0 | MySQL query details | 
| 37 | Mercenaries | 0 | 0 | 0 | 0 | unused | 
| 38 | Quest Debug | 0 | 0 | 0 | 0 | unused | 
| 44 | Client Login | 0 | 0 | 0 | 0 | Client login/disconnect/kick | 
| 46 | HP Update | 0 | 0 | 0 | 0 | Hitpoint changes | 
| 47 | FixZ | 0 | 0 | 0 | 0 | Z-axis failures and fixes | 
| 48 | Food | 0 | 0 | 0 | 0 | Food/Drink logging | 
| 49 | Traps | 0 | 0 | 0 | 0 | Trap details | 
| 50 | NPC Roam Box | 0 | 0 | 0 | 0 | Roambox info and failures for NPCs | 
| 51 | NPC Scaling | 0 | 0 | 1 | 0 | Scaling info for NPCs | 
| 52 | MobAppearance | 0 | 0 | 0 | 0 | Appearance changes for NPCs | 
| 53 | Info | 1 | 0 | 0 | 0 | Miscellanenous info (data loading, logging info, reloads, connection info) | 
| 54 | Warning | 1 | 0 | 0 | 0 | Warning messages (bootup info, connection conflicts and details) | 
| 59 | AI Scan | 0 | 0 | 0 | 0 | Close Mob entity list scanning | 
| 60 | AI Yell | 0 | 0 | 0 | 0 | Mob yell for help details (social aggro) | 
| 61 | AI CastBeneficial | 0 | 0 | 0 | 0 | NPC buffing of friends | 
| 62 | AOE Cast | 0 | 0 | 0 | 0 | AOE spell details via closemob list | 
| 63 | Entity Management | 0 | 0 | 0 | 0 | Entity List adds/removes | 
| 64 | Flee | 0 | 0 | 0 | 0 | Mob flee details | 
| 65 | Aura | 0 | 0 | 0 | 0 | Aura add/remove info | 
| 66 | HotReload | 1 | 0 | 1 | 0 | Quest hot reload info (reload on the fly) | 
| 67 | Merchants | 0 | 0 | 0 | 0 | Merchant loading / pricing info | 
| 68 | ZonePoints | 0 | 0 | 0 | 0 | Loading details for zone points and virtual zone points | 
| 69 | Loot | 0 | 0 | 1 | 0 | Lootdrop/loottable loading details and errors | 
| 70 | Expeditions | 0 | 0 | 0 | 0 | Expedition details (invite, add, remove, create, accept) | 
| 71 | DynamicZones | 0 | 0 | 0 | 0 | Dynamic Zone details (errors, create, load, saves, updates, etc) | 
| 72 | Scheduler | 1 | 0 | 0 | 0 | Scheduled event scans, loading, unloads (via [Server Events Scheduler](https://docs.eqemu.io/server/operation/server-events-scheduler/)) | 
| 73 | Cheat | 1 | 0 | 0 | 0 | Anticheat logging | 
| 74 | ClientList | 0 | 0 | 0 | 0 | Clientlist updates | 
| 75 | DialogueWindow | 0 | 0 | 0 | 0 | DiaWind popup info | 
| 76 | HTTP | 1 | 0 | 1 | 0 | Sidecar and SQL sourcing info | 
| 77 | Saylink | 0 | 0 | 0 | 0 | Saylink loading and caching | 
| 78 | ChecksumVer | 1 | 0 | 1 | 0 | Checksum validation checks | 
| 79 | CombatRecord | 0 | 0 | 1 | 0 | Combat info (engage, attacks, damage, duration, etc.) | 
| 80 | Hate | 0 | 0 | 0 | 0 | Hatelist scanning and updates | 
| 81 | Discord | 1 | 0 | 0 | 0 | Discord webhook logging | 
| 82 | Faction | 0 | 0 | 0 | 0 | Faction loading, reloading and errors | 
| 83 | Packet S->C | 0 | 0 | 0 | 0 | Packet logging from server to client | 
| 84 | Packet C->S | 0 | 0 | 0 | 0 | Packet logging from client to server | 
| 85 | Packet S->S | 0 | 0 | 0 | 0 | Packet logging between the server | 
| 86 | Bugs | 0 | 0 | 0 | 0 | Bug report create logging | 
| 87 | QuestErrors | 1 | 0 | 1 | 0 | Quest error logging (Lua/Perl) | 
| 88 | PlayerEvents | 0 | 0 | 0 | 0 | Player Event logging (loot, kills, death, merchant, speech, etc) | 
| 89 | DataBuckets | 0 | 0 | 0 | 0 | DataBucket info (deletes, creates, loading, reloading, errors, updates) | 
| 90 | Zoning | 0 | 0 | 0 | 0 | Player zoning info | 
| 91 | EqTime | 1 | 0 | 1 | 0 | In-game time updates and failures | 
| 92 | Corpses | 0 | 0 | 0 | 0 | Corpse logging (buries, rezzes, moves, loading, failures) | 
| 93 | XTargets | 0 | 0 | 0 | 0 | Add/Remove updates for Extended Targets | 
| 94 | EvolveItem | 0 | 0 | 0 | 0 | Evolving Item logging | 
| 95 | PositionUpdate | 0 | 0 | 0 | 0 | NPC position updates sent to clients | 
| 96 | KSM | 0 | 0 | 0 | 0 | KSM info for Linux (memory sharing) | 
| 97 | Bot Settings | 0 | 0 | 0 | 0 | Saving/Loading of bot settings | 
| 98 | Bot Spell Checks | 0 | 0 | 0 | 0 | Spell checks for bots (casting checks, failures, etc) | 
| 99 | Bot Spell Type Checks | 0 | 0 | 0 | 0 | Spell type checks for bots (spell list verification) | 
| 100 | NpcHandin | 1 | 0 | 1 | 0 | NPC Hand-In logging (NPC trading/quest hand-ins) | 
| 101 | ZoneState | 0 | 0 | 0 | 0 | Zone State logging (saving and restoring of the state of a zone on shutdown) | 
| 102 | Net Server <-> Client | 0 | 0 | 0 | 0 | Packet logging between the server and client | 
| 103 | Net TCP | 0 | 0 | 0 | 0 | Server packet logging | 
