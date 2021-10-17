# Player Housing

Player Housing was introduced in the House of Thule expansion. It can be accessed from the Guild Lobby and allows for players to purchase houses and place items in their houses. Houses do require upkeep to maintain.

There is a single zone for all neighborhood instances (Sunrise Hills). House zones are separate and appear to be spread across a variety of .zon files (e.g. phinterior1a1.zon) based on the type of home purchased.

The Place button in the Real Estate Items window is enabled on a zone-by-zone basis (probably). It is likely part of one of the zoning packets. For example, if you have the window open and click on a placeable item, the button remains disabled in the Guild Lobby. Repeat the process in a neighborhood or inside a residence and the button becomes enabled.

To my knowledge, no work has been done to add player housing to EQEMU. If it has, it hasn't been made public as no emu server has this capability as far as I am aware.

This page has packet information from Live as of 12/28/2018. While it is likely that the OpCodes have changed since House of Thule was released, my hope is that the structure of each packet has remained largely unchanged.

It looks like a single opcode (0x125C) is used for multiple packet types. At first I thought this was a code that caused a specific window to appear, but I've seen the same code used for the neighborhood search window (to select a specific neighborhood) and for the street search window inside a specific neighborhood. I have yet to see this code for other neighborhood or item placement actions.

Packet captures specific to trophy data is blocked at the moment. None of my characters on live have trophies, and I think that's why I'm not seeing a response from the server. There's nothing to send back. Unless I can find someone with at least one trophy (ideally a guild and a personal one), I'll have to spend some time questing to get one.

\[TODO: Add image of neighborhood search window]

### Current Issues

The primary issue facing Player Housing is that the placement window is disabled. I don't know if there is a missing packet on client initialization, one that appears only when entering a zone that supports placing objects, a client status (Paid or not) level, or if the copy of RoF2 used by players on the Emu simply doesn't support it (perhaps a bug that impacted that version).

This leaves the options for implementing player housing as a script-based approach. It may be possible to enable some client windows once objects are placed, but in the worst case some kind of interface can be created. It is unlikely that would be added to the official emu source, so any old RoF2 captures that include Neighborhood zones would be helpful.

Hint. Hint.

### Known / Suspected OpCodes in RoF and RoF2 client

The RoF2 opcodes are suspect. It could be that they didn't change across several patches, or that the opcode dumps in showeq simply didn't update those. I could only find two that directly related to player housing in the opcode files, but I suspect there are more that didn't get logged.

| Client | Name             | Description                                                                                                 |
| ------ | ---------------- | ----------------------------------------------------------------------------------------------------------- |
| RoF    | OP_HouseAddress  | 0x6786                                                                                                      |
|        | OP_HouseContents | <p>0x7CDA (11/28/2012)<br>0x49B7 (03/13/2013)<br>0x6419 (04/17/2013)</p>                                    |
| RoF2   | OP_HouseAddress  | 0x48AE (04/17/2013) - ShowEQ updated zoneopcodes.xml 04/22/2013,06/24/2013,07/18/2013,08/30/2013,10/09/2013 |
|        | OP_HouseContents | 0x6419 (04/17/2013) - ShowEQ update zoneopcodes.xml 04/22/2013,06/24/2013,07/18/2013,08/30/2013,10/09/2013  |
| CotF   | OP_HouseAddress  | <p>0x4F13 (11/06/13)<br>0x7B0D (01/22/2014)</p>                                                             |
|        | OP_HouseContents | <p>0x0588 (11/06/13)<br>0x41B3 (01/22/2014)</p>                                                             |

### Player Housing Zones

| Short Name                   | Long Name                    | Zone ID Number |
| ---------------------------- | ---------------------------- | -------------- |
| neighborhood                 | Sunrise Hills                | 712            |
| phinterior3a1                | House Interior               | 714            |
| phinterior1a1                | House Interior               | 715            |
| phinterior3a2                | House Interior               | 716            |
| phinterior3a3                | House Interior               | 717            |
| phinterior1a2                | House Interior               | 718            |
| phinterior1a3                | House Interior               | 719            |
| phinterior1b1                | Dragon House Interior        | 720            |
| phinterior1d1                | Dragon House Interior        | 723            |
| guildhalllrg                 | Grand Guild Hall             | 737            |
| guildhallsml                 | Greater Guild Hall           | 738            |
| plhogrinteriors1a1           | One Bedroom House Interior   | 739            |
| plhogrinteriors1a2           | One Bedroom House Interior   | 740            |
| plhogrinteriors3a1           | Three Bedroom House Interior | 741            |
| plhogrinteriors3a2           | Three Bedroom House Interior | 742            |
| plhogrinteriors3b1           | Three Bedroom House Interior | 743            |
| plhogrinteriors3b2           | Three Bedroom House Interior | 744            |
| plhdkeinteriors1a1           | One Bedroom House Interior   | 745            |
| plhdkeinteriors1a2           | One Bedroom House Interior   | 746            |
| plhdkeinteriors1a3           | One Bedroom House Interior   | 747            |
| plhdkeinteriors3a1           | Three Bedroom House Interior | 748            |
| plhdkeinteriors3a2           | Three Bedroom House Interior | 749            |
| plhdkeinteriors3a3           | Three Bedroom House Interior | 750            |
| guildhall3                   | Modest Guild Hall            | 751            |
| phinteriortree               | Evantil's Abode              | 766            |
| plhbixieint (Post RoF2)      | Bixie Hive                   | 774            |
| plhpirateshipint (Post RoF2) | Pirate Ship                  | 786            |

### Click on door to view Neighborhood List (Client -> Server)

| Name                   | Number of Bytes | Description                                             |
| ---------------------- | --------------- | ------------------------------------------------------- |
| Network OpCode         | 2               | 0x09                                                    |
| Packet Sequence Number | 2               |                                                         |
| Application OpCode     | 2               | <p>OP_ClickDoor<br>- Live: 0x2254<br>- RoF2: 0x3A8F</p> |
| Door Id                | 1               | 77 for the neighborhood "door" gate in Guild Lobby      |
| Unknown                | 1               | 0                                                       |
| Unknown                | 1               | 0                                                       |
| Unknown                | 1               | 0                                                       |
| picklockskill          | 1               | 0                                                       |
| Unknown5               | 3               | 0                                                       |
| Item Id                | 4               | 255,255,255,255                                         |
| Player Id              | 2               |                                                         |
| Unknown                | 2               |                                                         |

### Neighborhood List (Server -> Client)

**20190126 - Verified OpCode and Packet structure in RoF**

| Name                               | Number of Bytes          | Description                                                                                                                                          |
| ---------------------------------- | ------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| Network OpCode                     | 2                        | 0x0D                                                                                                                                                 |
| Packet Sequence Number             | 2                        | Identifies order of fragments                                                                                                                        |
| Total length of combined fragments | 3                        | Total size of all fragments. This is only present in the first fragment.                                                                             |
| Application OpCode                 | 2                        | <p>RoF: 0x67C9<br>Live (20181217): 0x5C12<br>Live (20190118): 0x1659</p>                                                                             |
| Unknown                            | 2                        | An id? Different across servers                                                                                                                      |
| Window ID                          | 7                        | <p>Always 1,0,6,0,0,0,0<br>This value is used to branch code paths between the Neighborhood list and Property marker windows.</p>                    |
| Total # of Neighborhoods           | 2                        |                                                                                                                                                      |
| Unknown                            | 5                        | Always 1,1,0,0,0                                                                                                                                     |
| Neighborhood ID                    | 4                        |                                                                                                                                                      |
| Neighborhood Name Length           | 4                        |                                                                                                                                                      |
| Neighborhood Name                  | Neighborhood Name Length | Not null terminated                                                                                                                                  |
| Guild ID                           | 4                        | 0 if no guild is associated with this neighborhood, otherwise it is the Guild ID                                                                     |
| Unknown                            | 4                        | Only seen on Live. If RoF, this field is not present in the packet: If Guild ID is 0, this is also 0. If Guild ID is not 0, this is always 147,0,0,0 |
| Max Player Plots                   | 4                        | Always 71                                                                                                                                            |
| Used Player Plots                  | 4                        |                                                                                                                                                      |
| Max Guild Plots                    | 4                        | Always 4                                                                                                                                             |
| Used Guild Plots                   | 4                        |                                                                                                                                                      |

Once the first neighborhood is parsed, the format repeats for the next Neighborhood ID -> Used Guild Plots. This continues until the last neighborhood is parsed.

### Searching for neighborhoods (or clicking on a specific one) (Client -> Server)

| Name                         | Number of Bytes | Description                                                                                                                                                                     |
| ---------------------------- | --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Network OpCode               | 2               | 0x09                                                                                                                                                                            |
| Packet Sequence Number       | 2               |                                                                                                                                                                                 |
| Application OpCode           | 2               | <p>RoF: 0x67C9<br>Live (20181217): 0x5C12<br>Live (20190118): 0x1659</p>                                                                                                        |
| Search Type                  | 4               | <p>0 when searching for a player or guild.<br>2 when clicking to view details of a specific neighborhood.<br>4 when inside a neighborhood and searching specific addresses.</p> |
| Neighborhood ID              | 4               | 1 if searching for a player. The Neighborhood ID if clicking on a neighborhood entry.                                                                                           |
| Unknown                      | 4               | 0s                                                                                                                                                                              |
| Neighborhood Search Criteria | 80              | A neighborhood's name. 79 bytes plus 1 for the null terminator.                                                                                                                 |
| Player Search Criteria       | 64              | A player's name. 63 bytes plus 1 for the null terminator.                                                                                                                       |
| Guild Search Criteria        | 64              | A guild's name. 63 bytes plus 1 for the null terminator.                                                                                                                        |
| Unknown                      | 2               | Always 1,1                                                                                                                                                                      |
| Unknown                      | 2               | Always 132,0 (Seen in RoF, but not in Live)                                                                                                                                     |

### Response to click to get neighborhood details (Server -> Client)

| Name                   | Number of Bytes    | Description                                                              |
| ---------------------- | ------------------ | ------------------------------------------------------------------------ |
| Network OpCode         | 2                  | 0x09                                                                     |
| Packet Sequence Number | 2                  |                                                                          |
| Application OpCode     | 2                  | <p>RoF: 0x67C9<br>Live (20181217): 0x5C12<br>Live (20190118): 0x1659</p> |
| Unknown                | 4                  | 173,153,9,0                                                              |
| Unknown                | 1                  | 10                                                                       |
| Unknown                | 4                  | 2,0,0,0                                                                  |
| # of Players           | 2                  |                                                                          |
| Player Name Length     | 4                  |                                                                          |
| Player Name            | Player Name Length | Not null terminated                                                      |
| # of Guilds            | 2                  |                                                                          |
| Guild ID               | 4                  | If # of guilds = 0, this field is not in the packet.                     |
| Guild Terminator       | 4                  | 147,0,0,0 : If # of guilds = 0, this field is not in the packet.         |

If # of Players is greater than 0, the fields Player Name Length and Player Name repeat. The same is true for # of Guilds.

### Click on a property marker (Server -> Client) (this is further along, but still has some unknowns)

| Name                   | Number of Bytes       | Description                                                                                                                       |
| ---------------------- | --------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| Network OpCode         | 2                     | 0x09                                                                                                                              |
| Packet Sequence Number | 2                     |                                                                                                                                   |
| Application OpCode     | 2                     | <p>RoF: 0x67C9<br>Live (20181217): 0x5C12<br>Live (20190118): 0x1659</p>                                                          |
| Unknown                | 2                     | An id? Different across servers                                                                                                   |
| Window ID              | 7                     | <p>Always 1,0,4,3,0,0,0<br>This value is used to branch code paths between the Neighborhood list and Property marker windows.</p> |
| Unknown                | 6                     |                                                                                                                                   |
| Neighborhood ID        | 4                     |                                                                                                                                   |
| Unknown                | 4                     |                                                                                                                                   |
| Unknown                | 4                     |                                                                                                                                   |
| Guild ID               | 4                     |                                                                                                                                   |
| Unknown                | 4                     | 123,0,0,0 (may not be constant)                                                                                                   |
| Unknown                | 12                    |                                                                                                                                   |
| Street Address Length  | 4                     |                                                                                                                                   |
| Street Address         | Street Address Length | Not null terminated.                                                                                                              |
| Unknown Name Length    | 4                     |                                                                                                                                   |
| Unknown Value          | Unknown Name Length   |                                                                                                                                   |
| Owner Name Length      | 4                     |                                                                                                                                   |
| Owner Name             | Owner Name Length     |                                                                                                                                   |
| Capacity               | 4                     |                                                                                                                                   |
| Vault                  | 4                     |                                                                                                                                   |
| Pets                   | 4                     |                                                                                                                                   |
| Unknown                | 4                     |                                                                                                                                   |
| Unknown                | 1                     |                                                                                                                                   |

### Real Estate Items window, Click on Trophies button (Client -> Server)

Possible new opcode: 0xA02F Length: 342 bytes

| Name                   | Number of Bytes | Description |
| ---------------------- | --------------- | ----------- |
| Network OpCode         | 2               | 0x09        |
| Packet Sequence Number | 2               |             |
| Application OpCode     | 2               | 0xA02F      |
| Unknown                | 4               | 0           |
| Unknown                | 4               | 1           |
| Unknown                | 328             | 0           |

### OP_NewZone - phinterior3a3 (Server -> Client)

Matches the RoF2 NewZone_Struct, but 20 bytes of new data that comes after the Blooming member variable in the RoF2 struct.

| Name                       | Number of Bytes | Description                                                                                                                              |
| -------------------------- | --------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| Network OpCode             | 2               | 0x09                                                                                                                                     |
| Packet Sequence Number     | 2               |                                                                                                                                          |
| Application OpCode         | 2               | <p>RoF2: 0x1795<br>Live: 0x2F3B</p>                                                                                                      |
| CharName                   | 64              |                                                                                                                                          |
| Zone Short Name            | 128             |                                                                                                                                          |
| Zone Long Name             | 128             |                                                                                                                                          |
| Unknown326                 | 174             | All 0s.                                                                                                                                  |
| ztype                      | 1               |                                                                                                                                          |
| fog_red                    | 4               |                                                                                                                                          |
| fog_green                  | 4               |                                                                                                                                          |
| fog_blue                   | 4               |                                                                                                                                          |
| unknown507                 | 1               |                                                                                                                                          |
| fox_minclip                | 4               |                                                                                                                                          |
| fog_maxclip                | 4               |                                                                                                                                          |
| Gravity                    | 4               |                                                                                                                                          |
| TimeType                   | 1               |                                                                                                                                          |
| rain_chance                | 4               |                                                                                                                                          |
| rain_duration              | 4               |                                                                                                                                          |
| snow_chance                | 4               |                                                                                                                                          |
| snow_duration              | 4               |                                                                                                                                          |
| ZoneTimeZone               | 1               |                                                                                                                                          |
| Sky                        | 1               |                                                                                                                                          |
| Unknown539                 | 1               | See if this changes across zones.                                                                                                        |
| WaterMidi                  | 4               |                                                                                                                                          |
| DayMidi                    | 4               |                                                                                                                                          |
| NightMidi                  | 4               |                                                                                                                                          |
| zone_exp_multiplier        | 4               |                                                                                                                                          |
| safe_y                     | 4               |                                                                                                                                          |
| safe_x                     | 4               |                                                                                                                                          |
| safe_z                     | 4               |                                                                                                                                          |
| min_z                      | 4               |                                                                                                                                          |
| max_z                      | 4               |                                                                                                                                          |
| underworld                 | 4               |                                                                                                                                          |
| minclip                    | 4               |                                                                                                                                          |
| maxclip                    | 4               |                                                                                                                                          |
| ForageLow                  | 4               |                                                                                                                                          |
| ForageMedium               | 4               |                                                                                                                                          |
| ForageHigh                 | 4               |                                                                                                                                          |
| FishingLow                 | 4               |                                                                                                                                          |
| FishingMedium              | 4               |                                                                                                                                          |
| FishingHigh                | 4               |                                                                                                                                          |
| sky_lock                   | 4               |                                                                                                                                          |
| graveyard_timer            | 4               |                                                                                                                                          |
| scriptIDHour               | 4               |                                                                                                                                          |
| scriptIDMinute             | 4               |                                                                                                                                          |
| scriptIDTick               | 4               |                                                                                                                                          |
| scriptIDOnPlayerDeath      | 4               |                                                                                                                                          |
| scriptIDOnNPCDeath         | 4               |                                                                                                                                          |
| scriptIDPlayerEnteringZone | 4               |                                                                                                                                          |
| scriptIDOnZonePop          | 4               |                                                                                                                                          |
| scriptIDNPCLoot            | 4               |                                                                                                                                          |
| scriptIDAdventureFailed    | 4               |                                                                                                                                          |
| CanExploreTasks            | 4               |                                                                                                                                          |
| UnknownFlag                | 4               | Appears to be 1 for neighborhood related zones, but setting this to 1 and sending to RoF2 did not enable the Place button in the client. |
| scriptIDOnFishing          | 4               |                                                                                                                                          |
| scriptIDOnForage           | 4               |                                                                                                                                          |
| zone_short_name2\[32]      | 32              | name of the housing file (e.g. phinterior3a3)                                                                                            |
| WeatherString\[32]         | 32              |                                                                                                                                          |
| SkyString2\[32]            | 32              |                                                                                                                                          |
| SkyRelated2                | 4               |                                                                                                                                          |
| WeatherString2\[32]        | 32              |                                                                                                                                          |
| WeatherChangeTime          | 4               |                                                                                                                                          |
| Climate                    | 4               |                                                                                                                                          |
| NPCAggroMaxDist            | 4               |                                                                                                                                          |
| FilterID                   | 4               |                                                                                                                                          |
| ZoneID                     | 2               | 205,2 Housing instance?                                                                                                                  |
| ZoneInstance               | 2               |                                                                                                                                          |
| ScriptNPCReceivedAnItem    | 4               |                                                                                                                                          |
| bCheck                     | 4               |                                                                                                                                          |
| scriptIDSomething          | 4               |                                                                                                                                          |
| scriptIDSomething2         | 4               |                                                                                                                                          |
| scriptIDSOmething3         | 4               |                                                                                                                                          |
| SuspendBuffs               | 4               |                                                                                                                                          |
| LavaDamage                 | 4               |                                                                                                                                          |
| MinLavaDamage              | 4               |                                                                                                                                          |
| bDisallowManaStone         | 1               |                                                                                                                                          |
| bNoBind                    | 1               |                                                                                                                                          |
| bNoAttack                  | 1               |                                                                                                                                          |
| bNoCallOfHero              | 1               |                                                                                                                                          |
| bNoFlux                    | 1               |                                                                                                                                          |
| bNoFear                    | 1               |                                                                                                                                          |
| fall_damage                | 1               |                                                                                                                                          |
| unknown863                 | 1               | 0                                                                                                                                        |
| FastRegenHP                | 4               | 180                                                                                                                                      |
| FastRegenMana              | 4               | 180                                                                                                                                      |
| FastRegenEndurance         | 4               | 180                                                                                                                                      |
| CanPlaceCampsite           | 4               | Seen 2                                                                                                                                   |
| CanPlaceGuildBanner        | 4               | Seen 2                                                                                                                                   |
| FogDensity                 | 4               | Seen 0.33                                                                                                                                |
| bAdjustGamma               | 4               |                                                                                                                                          |
| TimeStringID               | 4               |                                                                                                                                          |
| bNoMercenaries             | 4               |                                                                                                                                          |
| FishingRelated             | 4               | Seen -1                                                                                                                                  |
| ForageRelated              | 4               | Seen -1                                                                                                                                  |
| bNoLevitate                | 4               |                                                                                                                                          |
| Blooming                   | 4               | Seen 1                                                                                                                                   |
| Unk916                     | 4               | Seen 0                                                                                                                                   |
| Unk920                     | 4               | Seen 100                                                                                                                                 |
| Unk924                     | 4               | Seen 0                                                                                                                                   |
| Unk928                     | 4               | Seen 1                                                                                                                                   |
| Unk932                     | 4               | Seen 1                                                                                                                                   |

TODO: In-zone clicks for the spring, purchasing a property, placing items in a property, identifying RoF opcodes, updating with additional vendors, fixing Z for existing vendors (some are floating), adding house interior zones (first attempt results in a client crash), etc.

### Objects (Houses)

#### plhousingexta.eqg

* IT20000
* IT20001
* IT20002
* IT20003
* IT20004
* IT20005

#### plhousingextb.eqg

* IT20006
* IT20007
* IT20008
* IT20009

#### phexterior1a.eqg

* IT20017
* IT20018
* IT20019
* IT20020
* IT20021
* IT20022
* IT20023
* IT20024
* IT20025
* IT20026

#### neighborhood.eqg

* IT20027 - Sign
