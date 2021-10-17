---
description: This page lists the logging categories available in EQEmu.
---

# Logging Categories

Logging categories have their settings stored in the [logsys\_categories](https://eqemu.gitbook.io/database-schema/categories/admin/logsys_categories) table, but the categories are defined in [common/eqemu\_logsys.h](https://github.com/EQEmu/Server/blob/master/common/eqemu_logsys.h)

{% hint style="info" %}
The default settings for the "log-to" categories \(Console, File, GM Say\) are listed below.  You can modify these values for your own server in your `logsys_categories` table.
{% endhint %}

| ID | Description | Console | File | GMSay | Notes |
| :--- | :--- | :--- | :--- | :--- | :--- |
| 1 | AA | 0 | 0 | 0 | Use of AA actions \(IE activated, cast\) |
| 2 | AI | 0 | 0 | 0 | Logs the AI for mercs and bots \(IE spell casting, aggro, etc.\) |
| 3 | Aggro | 0 | 0 | 0 | Mob aggro, prox and range |
| 4 | Attack | 0 | 0 | 0 | Weapon skill and weapon in use, and mob aggro |
| 5 | Packet :: Client -&gt; Server | 0 | 0 | 0 | See [Logging Packets](logging-packets.md) |
| 6 | Combat | 0 | 0 | 0 | Combat damage \(IE weapon, base, thrown\) |
| 7 | Commands | 0 | 0 | 0 | Commands that require account status \(IE GM\) |
| 8 | Crash | 1 | 1 | 0 | Stacktrace, exceptions, timestamp |
| 9 | Debug | 0 | 0 | 0 | Logs failures and errors \(IE loading assets, login\) |
| 10 | Doors | 0 | 0 | 0 | Use of doors \(IE click, open, trigger\) |
| 11 | Error | 1 | 1 | 0 | Logs failures and errors \(IE loading assets, DB connection\) |
| 12 | Guilds | 0 | 0 | 0 | Guild activities \(IE update MOTD, DB queries\) |
| 13 | Inventory | 0 | 0 | 0 | Activities with inventory \(IE items, slots, shared inventory\) |
| 14 | Launcher | 1 | 1 | 0 | World / Zone launch requests |
| 15 | Netcode | 0 | 0 | 0 | Logs opcode issues \(IE login, world connection\) |
| 16 | Normal | 1 | 1 | 0 | Just normal things \(like /say\) |
| 17 | Object | 0 | 0 | 0 | Logging for objects \(IE loom\) |
| 18 | Pathing | 0 | 0 | 0 | Pathing waypoint logging \(IE arrive, pause\) |
| 19 | QS Server | 0 | 0 | 0 | Log query server data \(as defined by rules\) |
| 20 | Quests | 0 | 0 | 0 | Outputs quest errors to log |
| 21 | Rules | 0 | 0 | 0 | Logs the rule system \(IE load rules\) |
| 22 | Skills | 0 | 0 | 0 | Logs skill messages \(IE failed combine\) |
| 23 | Spawns | 0 | 0 | 0 | All about spawns \(IE ground, spawn2, conditions\) |
| 24 | Spells | 0 | 0 | 0 |  |
| 25 | Status | 1 | 1 | 0 | DB / Host / Port log |
| 26 | TCP Connection | 0 | 0 | 0 | Logs TCP connect / disconnect |
| 27 | Tasks | 0 | 0 | 0 |  |
| 28 | Tradeskills | 0 | 0 | 0 |  |
| 29 | Trading | 0 | 0 | 0 |  |
| 30 | Tribute | 0 | 0 | 0 |  |
| 31 | UCS Server | 1 | 1 | 0 |  |
| 32 | WebInterface Server | 1 | 1 | 0 |  |
| 33 | World Server | 1 | 1 | 0 |  |
| 34 | Zone Server | 1 | 1 | 0 |  |
| 35 | MySQL Error | 1 | 1 | 1 |  |
| 36 | MySQL Query | 0 | 0 | 0 |  |
| 37 | Mercenaries | 0 | 0 | 0 |  |
| 38 | Quest Debug | 0 | 0 | 1 |  |
| 39 | Packet :: Server -&gt; Client | 0 | 0 | 0 | See [Logging Packets](logging-packets.md) |
| 40 | Packet :: Client -&gt; Server Unhandled | 0 | 0 | 0 | See [Logging Packets](logging-packets.md) |
| 41 | Packet :: Server -&gt; Client \(Dump\) | 0 | 0 | 0 | See [Logging Packets](logging-packets.md) |
| 42 | Packet :: Client -&gt; Server \(Dump\) | 0 | 0 | 0 | See [Logging Packets](logging-packets.md) |
| 43 | Login Server | 0 | 0 | 0 |  |
| 44 | Client Login | 1 | 1 | 0 |  |
| 45 | Headless Client | 0 | 0 | 0 |  |
| 46 | HP Update | 0 | 0 | 0 |  |
| 47 | FixZ | 0 | 0 | 0 |  |
| 48 | Food | 0 | 0 | 0 |  |
| 49 | Traps | 0 | 0 | 0 |  |
| 50 | NPC Roam Box | 0 | 0 | 0 |  |
| 51 | NPC Scaling | 0 | 0 | 0 |  |
| 52 | Mob Appearance | 0 | 0 | 0 |  |
| 53 | Info | 1 | 1 | 0 |  |
| 54 | Warning | 1 | 0 | 0 |  |
| 55 | Critical | 1 | 0 | 0 |  |
| 56 | Emergency | 1 | 0 | 0 |  |
| 57 | Alert | 1 | 0 | 0 |  |
| 58 | Notice | 1 | 0 | 0 |  |
| 59 | AI Scan Close | 0 | 0 | 0 |  |
| 60 | AI Yell For Help | 0 | 0 | 0 |  |
| 61 | AI Cast Beneficial Close | 0 | 0 | 0 |  |
| 62 | AOE Cast | 0 | 0 | 0 |  |
| 63 | Entity Management | 0 | 0 | 0 |  |
| 64 | Flee | 0 | 0 | 0 |  |
| 65 | Aura | 0 | 0 | 0 |  |
| 66 | HotReload | 1 | 0 | 1 |  |

