---
description: This page introduces scripting for your EQEmu Server.
---

# Scripting

Please note that there is extensive documentation of the [EQEmu Quest Api](https://eqemu.gitbook.io/quest-api/) available on this site.  This documentation includes sample scripts describing functions, variables, sub routines and the like.  

The information found here is intended to get someone started down the path of creating their own custom scripts on their world.

## Getting Started

If you are new to the EQEmu Quest API, it can easily be daunting to try to digest the vast array of methods, functions, subroutines, plugins, and varying quest-scripting languages.  Quest scripts in EQEmu are written in [Lua](https://www.lua.org/) or [Perl](https://www.perl.org/), both of which are widely accepted throughout the EQEmu community.

**If you are new to both Lua and Perl**, a good place to get started is the [Beginner's Guide to Perl and Lua.](https://eqemu.gitbook.io/quest-api/methods/beginners-guide)

**If you are adept at either Lua or Perl**, it is best to familiarize yourself with how EQEmu [loads quests scripts](https://eqemu.gitbook.io/quest-api/methods/quest-loading), followed by reviewing the various [Events](https://eqemu.gitbook.io/quest-api/methods/events) that can be triggered in either Perl or Lua, along with  exports of those events \(as well as the [global exports](https://eqemu.gitbook.io/quest-api/perl/exports)\), and then beginning to review the many methods and functions exposed in the API.

## To Perl, or to Lua?

The decision to utilize either Perl or Lua for your quest scripts is entirely up to you.  Both offer nearly all the same functions and methods--where discrepancies exist, they are noted \(for instance the [Spell Methods \(Lua\)](https://eqemu.gitbook.io/quest-api/methods/spell-methods-lua) page describes methods only available in Lua.  Beginners often find Perl easier to understand. Adept Lua programmers might enjoy the ability to load [encounters](encounters-lua.md).  Advanced Perl programmers may find the ability to create or utilize [plugins](https://eqemu.gitbook.io/quest-api/perl/plugins) very desirable.

All attempts have been made to provide corresponding examples in both languages throughout this Wiki.  Many community members have a preference, and certainly there are many members of the community that can provide you guidance for either language--please look for us in the **\#Support-Quest-Scripts** channel on the [EQEmu Discord](https://discord.gg/QHsm7CD) server.

## Some Basics to get you Started

We're frequently asked for certain types of scripts that server operators need to create NPCs in their world that can provide some basic services to players. 

### Player Buffing

The [Player Buffer Scripts](player-buffer-scripts.md) page gives you examples of how to automatically buff players \(for instance, when a player levels up and needs new skills or spells\), as well as examples of how to include NPCs that perform these duties or cast specific buffs when a player interacts with the Buffer NPC.

These scripts range from fairly straightforward to relatively complex--be sure to reference [Event Subroutines](https://eqemu.gitbook.io/quest-api/methods/events) to learn why these scripts fire, [Quest API Functions](https://eqemu.gitbook.io/quest-api/perl/functions) to learn what each function does and how to properly call the function, and finally, [Customizing NPCs](https://eqemu.gitbook.io/server/categories/npc/customizing-npcs) if you would like to learn how to create your own NPC who will utilize any of these quest scripts.

### Player Teleporting

The [Player Teleporter Scripts](https://eqemu.gitbook.io/server/categories/scripting/player-teleporter-scripts) page gives some examples of how you can move players around in game.  Don't forget that there are plenty of examples in the default scripts that were installed with your server, such as the Nexus portals.  

Often server operators would like to teleport users to a specific instance of a zone as well, which would allow you to create raid instances for competing guilds, or create custom dungeon crawls for your players.  The sample script includes something called a "Data Bucket", so you may want to read [Using Data Buckets](using-data-buckets.md) to gain some understanding of this flexible, performant system.

