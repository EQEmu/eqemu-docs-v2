---
description: >-
  This page gives information about server management
---

# Server Management

### Content from ProjectEQ and EQEmu

[ProjectEQ](https://projecteq.com) is a team dedicated to the improvement of the Everquest Emulation community by providing era-accurate databases and quests.

[EQEmulator](https://eqemulator.org) is a team dedicated to the EQEmu Server used by nearly all EverQuest private servers.

Under most servers, these two team's efforts are considered the "Official" private server bundles to baseline a new server from.

Other alternatives exist, but are not maintained as much.


## Content Types 

### Server Content

Inside an EQEmu server, a binary folder exists that houses the binaries (or executables) ran to make a server ran. This includes files such as zone.exe (to run zones), world.exe (to connect to a login server and manage zones, and also do character select for a player), ucs.exe (Universal Chat Service, optional and used for the chat channels in game), queryserv.exe (Used for storing additional data in logs such as trade logging)

### Quest Content

Quests are what bring life to NPCs in EverQuest. These quests utilize the Quest API to handle trading, /say interactions with an NPC, movement, and virtually everything else that may be unique to an NPC.

### Database Content

A Database stores the information about anything modifiable inside EverQuest. EQEMu currently only supports MySQL (or MariaDB) based databases.

### Map Content

Maps are used on server side to give routing and mapping information for NPCs to smartly figure out the path it should take to get to a destination, especially when chasing NPCs. Watermaps also exist, that give additional details about how water should be handled in a zone.

## Configuration Inside eqemu_config.json

Inside your server directory, a file exists called eqemu_config.json. This file holds key information needed to run your EverQuest server, such as database connection info, what path to find various content noted above, as well as how to identify the server.

### Long Name

This is how your server is identified when on the Server Select screen inside EverQuest. If your server is not registered, it will have a |T| or |U| prefix on it.

### Short Name

This is how your server's character profiles are stored, it is recommend to keep these short and relevant to your long name.

## Server Management

When managing an EverQuest private server, the following sections are areas worth noting.

### SQL

SQL stands for Server Query Language, and represents a Database. A SQL Database is where all data is stored about EverQuest.

### Shared Memory

Shared Memory creates file hashes that get loaded by World and Zone during bootup and are responsible for data such as items, faction, and base data. Shared Memory typically only needs to be ran once during initial setup, or when a major update occurs.


