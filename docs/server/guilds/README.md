---
description: This page describes the guild management system on your EQEmu Server.
---

# Guilds

The internal structure of guild management is based on the Rain of Fear 2 (RoF2) client.  This means that the internal structure is based on the 8 guild ranks and rank names.  The server will translate these to older clients when the player is using Titanium or Underfoot.

The guild management system has not been tested with other clients.  If your player base uses another client, please post an [issue](https://github.com/EQEmu/Server/issues) as it may be possible to add the appropriate opcodes.

As of February 2024, the following **features are implemented**:

* Customizable Guild Ranks (RoF2)
* Permissions (RoF2)
* Notes
* Guild Creation using Alt-G (UF/RoF2)
* Tributes
* Guild Bank
* Guild LFG.  Works from the command line and the Bulletin Boards in Guild Lobby (requires queryserver to be running)

**Not Yet Implemented:**

- Guild Trophies
- Real Estate
- Banners

The system has been tested against Titanium, Underfoot and Rain of Fear 2.  Functionality is limited to the client feature set.

## Server Rules for Guild Management

There are several rules that allow for customization of guild functionality.

| Category | Rule Name | Type | Default | Description |
| :--- | :--- | :--- | :--- | :--- |
| Guild | TributePlatConversionRate | int | 10 | The conversion rate of platinum donations.  Default is 10 guild favor to 1 platinum |
| Guild | TributeTime | int | 600000 | Time in ms for guild tributes.  Default is 10 mins. |
| Guild | TributeTimeRefreshInterval | int | 180000 | Time in ms to send timer updates to all guild members. Default is 3 mins. |
| Guild | UseCharacterMaxLevelForGuildTributes | bool | true | Guild Tributes will adhere to Character:MaxLevel.  Default is true. |

## Scripting

If you have existing quest scripts that use client:GuildRank() or $client->GuildRank(), these will need to be updated to reflect the RoF2 guild rank numbers.

## Tributes

As of release 22.44.5, the guild tribute data can be sourced from content_database.  Prior to this release, the guild tribute must be sourced from the peq_content tributes and tribute_level tables.  These are included within a new install, and can also be sourced from the peq content [database](https://db.projecteq.net/)

## Defaults

There have been many questions about the default behaviour of guild permissions.  Care was taken to align with Live.  If you encounter any discrepancies, please report them as an issue on github.  Also, please checkout the [Default](./defaults.md) section.

## Player Event Logging

The following player events are available.

|Rule Name|Default|Rentention|
|:---|:---|:---|
|Guild Tribute Donate Item|Enabled|7 days|
|Guild Tribute Donate Platinum|Enabled|7 days|

## Known Issues

There is a display issue within the Guild Management Window across all three clients (Ti/UF/RoF2).  Care has been taken to alleviate the issue, however it is explained here for reference.

When a player opens the guild management window and has 'Show Offline' unchecked, they may notice that the 'Notes' and 'Tribute' tabs are missing an online client.  If they toggle the 'Show Offline' checkbox, the Guild Management Window will update itself and correct the display issue. 

The guild tribute window may show a level of 0 when the tribute level is higher than the CharacterMaxLevel.  This is a display bug only and will not allow the tribute to be used.
