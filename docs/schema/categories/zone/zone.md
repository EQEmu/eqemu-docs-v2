# zone

| Column | Data Type | Description |
| :--- | :--- | :--- |
| short\_name | varchar | Short Name |
| id | int | Unique Entry Identifier |
| file\_name | varchar | File Name |
| long\_name | text | Long Name |
| map\_file\_name | varchar | Map File Name |
| safe\_x | float | Safe X Coordinate |
| safe\_y | float | Safe Y Coordinate |
| safe\_z | float | Safe Z Coordinate |
| graveyard\_id | float | [Graveyard Identifier](../../../schema/categories/zone/graveyard.md) |
| min\_level | tinyint | Minimum Level |
| min\_status | tinyint | [Minimum Status](../../../../categories/player/status-levels) |
| zoneidnumber | int | [Unique Zone Identifier](../../../../categories/zones/zone-list) |
| version | tinyint | Version |
| timezone | int | Timezone |
| maxclients | int | Maximum Clients |
| ruleset | int | [Ruleset Identifier](../../../schema/categories/zone/rule_sets.md) |
| note | varchar | Note |
| underworld | float | Bottom Z to represent when the player is under the world |
| minclip | float | Minimum Clipping Distance |
| maxclip | float | Maximum Clipping Distance |
| fog\_minclip | float | Fog Minimum Clipping Distance |
| fog\_maxclip | float | Fog Maximum Clipping Distance |
| fog\_blue | tinyint | Fog Blue Level: 0 = None, 255 = Max |
| fog\_red | tinyint | Fog Red Level: 0 = None, 255 = Max |
| fog\_green | tinyint | Fog Green Level: 0 = None, 255 = Max |
| sky | tinyint | Sky type the client will present as the backdrop |
| ztype | tinyint | This field is sent directly to the client on zone-in, most zones are set to 0, 1 or 255. |
| zone\_exp\_multiplier | decimal | This will multiply the XP to this percentage value \(decimal based, 100% = 1.0\) |
| walkspeed | float | Walkspeed in this zone |
| time\_type | tinyint | This value varies depending on the zone but it is sent to the client on zone in.  Most starting zones/newbie areas have this value set to 2, dungeons tyically have this set to 0, some zones break from the norm and have values greater than 2, \(akanon = 3, blackburrow = 5, cazicthule = 5, crushbone = 5, erudnint = 4, kaladima = 3, etc.\) |
| fog\_red1 | tinyint | Fog Red Level 1: 0 = None, 255 = Max |
| fog\_green1 | tinyint | Fog Green Level 1: 0 = None, 255 = Max |
| fog\_blue1 | tinyint | Fog Blue Level 1: 0 = None, 255 = Max |
| fog\_minclip1 | float | Fog Minimum Clipping Distance 1 |
| fog\_maxclip1 | float | Fog Maximum Clipping Distance 1 |
| fog\_red2 | tinyint | Fog Red Level 2: 0 = None, 255 = Max |
| fog\_green2 | tinyint | Fog Green Level 2: 0 = None, 255 = Max |
| fog\_blue2 | tinyint | Fog Blue Level 2: 0 = None, 255 = Max |
| fog\_minclip2 | float | Fog Minimum Clipping Distance 2 |
| fog\_maxclip2 | float | Fog Maximum Clipping Distance 2 |
| fog\_red3 | tinyint | Fog Red Level 3: 0 = None, 255 = Max |
| fog\_green3 | tinyint | Fog Green Level 3: 0 = None, 255 = Max |
| fog\_blue3 | tinyint | Fog Blue Level 3: 0 = None, 255 = Max |
| fog\_minclip3 | float | Fog Minimum Clipping Distance 3 |
| fog\_maxclip3 | float | Fog Maximum Clipping Distance 4 |
| fog\_red4 | tinyint | Fog Red Level 4: 0 = None, 255 = Max |
| fog\_green4 | tinyint | Fog Green Level 4: 0 = None, 255 = Max |
| fog\_blue4 | tinyint | Fog Blue Level 4: 0 = None, 255 = Max |
| fog\_minclip4 | float | Fog Minimum Clipping Distance 4 |
| fog\_maxclip4 | float | Fog Maximum Clipping Distance 4 |
| fog\_density | float | This is the intensity of the fog, this should be a number between 0-1, most commonly used is .1 or .33 |
| flag\_needed | varchar | [Flag Required](zone_flags.md) |
| canbind | tinyint | Can Bind: 0 = False, 1 = True \(for Caster\), 2 = True \(for All\) |
| cancombat | tinyint | Can Combat: 0 = False, 1 = True |
| canlevitate | tinyint | Can Levitate: 0 = False, 1 = True \(Does not affect those with \#gm on\) |
| castoutdoor | tinyint | Cast Outdoors: 0 = False, 1 = True |
| hotzone | tinyint | Hotzone: 0 = False, 1 = True |
| insttype | tinyint | Instance Type |
| shutdowndelay | bigint | Shutdown Delay |
| peqzone | tinyint | \#peqzone: 0 = False, 1 = True |
| expansion | tinyint | [Expansion](../../../../categories/operation/expansion-list) |
| suspendbuffs | tinyint | Suspend Buffs: 0 = False, 1 = True |
| rain\_chance1 | int | Rain Chance 1 |
| rain\_chance2 | int | Rain Chance 2 |
| rain\_chance3 | int | Rain Chance 3 |
| rain\_chance4 | int | Rain Chance 4 |
| rain\_duration1 | int | Rain Duration 1 |
| rain\_duration2 | int | Rain Duration 2 |
| rain\_duration3 | int | Rain Duration 3 |
| rain\_duration4 | int | Rain Duration 4 |
| snow\_chance1 | int | Snow Chance 1 |
| snow\_chance2 | int | Snow Chance 2 |
| snow\_chance3 | int | Snow Chance 3 |
| snow\_chance4 | int | Snow Chance 4 |
| snow\_duration1 | int | Snow Duration 1 |
| snow\_duration2 | int | Snow Duration 2 |
| snow\_duration3 | int | Snow Duration 3 |
| snow\_duration4 | int | Snow Duration 4 |
| gravity | float | Gravity |
| type | int | Type \(0 = Unknown, 1 = Regular, 2 = Instanced, 3 = Hybrid, 4 = Raid, 5 = City\) |
| skylock | tinyint | Sky Lock |
| fast\_regen\_hp | int | Fast Regen Health |
| fast\_regen\_mana | int | Fast Regen Mana |
| fast\_regen\_endurance | int | Fast Regen Endurance |
| npc\_max\_aggro\_dist | int | NPC Max Aggro Distance |
| max\_movement\_update\_range | int | Max Movement Update Range |

