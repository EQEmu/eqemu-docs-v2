# zone

| Column | Data Type | Description |
| :--- | :--- | :--- |
| short_name | varchar | Short Name |
| id | int | Unique Entry Identifier |
| file_name | varchar | File Name |
| long_name | text | Long Name |
| map_file_name | varchar | Map File Name |
| safe_x | float | Safe X Coordinate |
| safe_y | float | Safe Y Coordinate |
| safe_z | float | Safe Z Coordinate |
| graveyard_id | float | [Graveyard Identifier](../../../schema/categories/graveyards/graveyard.md) |
| min_level | tinyint | Minimum Level |
| min_status | tinyint | [Minimum Status](../../../../categories/player/status-levels) |
| zoneidnumber | int | [Unique Zone Identifier](../../../../categories/zones/zone-list) |
| version | tinyint | Version |
| timezone | int | Timezone |
| maxclients | int | Maximum Clients |
| ruleset | int | [Ruleset Identifier](../../../schema/categories/rules/rule_sets.md) |
| note | varchar | Note |
| underworld | float | Bottom Z to represent when the player is under the world |
| minclip | float | Minimum Clipping Distance |
| maxclip | float | Maximum Clipping Distance |
| fog_minclip | float | Fog Minimum Clipping Distance |
| fog_maxclip | float | Fog Maximum Clipping Distance |
| fog_blue | tinyint | Fog Blue Level: 0 = None, 255 = Max |
| fog_red | tinyint | Fog Red Level: 0 = None, 255 = Max |
| fog_green | tinyint | Fog Green Level: 0 = None, 255 = Max |
| sky | tinyint | Sky type the client will present as the backdrop |
| ztype | tinyint | This field is sent directly to the client on zone-in, most zones are set to 0, 1 or 255. |
| zone_exp_multiplier | decimal | This will multiply the XP to this percentage value \(decimal based, 100% = 1.0\) |
| walkspeed | float | Walkspeed in this zone |
| time_type | tinyint | This value varies depending on the zone but it is sent to the client on zone in.  Most starting zones/newbie areas have this value set to 2, dungeons tyically have this set to 0, some zones break from the norm and have values greater than 2, \(akanon = 3, blackburrow = 5, cazicthule = 5, crushbone = 5, erudnint = 4, kaladima = 3, etc.\) |
| fog_red1 | tinyint | Fog Red Level 1: 0 = None, 255 = Max |
| fog_green1 | tinyint | Fog Green Level 1: 0 = None, 255 = Max |
| fog_blue1 | tinyint | Fog Blue Level 1: 0 = None, 255 = Max |
| fog_minclip1 | float | Fog Minimum Clipping Distance 1 |
| fog_maxclip1 | float | Fog Maximum Clipping Distance 1 |
| fog_red2 | tinyint | Fog Red Level 2: 0 = None, 255 = Max |
| fog_green2 | tinyint | Fog Green Level 2: 0 = None, 255 = Max |
| fog_blue2 | tinyint | Fog Blue Level 2: 0 = None, 255 = Max |
| fog_minclip2 | float | Fog Minimum Clipping Distance 2 |
| fog_maxclip2 | float | Fog Maximum Clipping Distance 2 |
| fog_red3 | tinyint | Fog Red Level 3: 0 = None, 255 = Max |
| fog_green3 | tinyint | Fog Green Level 3: 0 = None, 255 = Max |
| fog_blue3 | tinyint | Fog Blue Level 3: 0 = None, 255 = Max |
| fog_minclip3 | float | Fog Minimum Clipping Distance 3 |
| fog_maxclip3 | float | Fog Maximum Clipping Distance 4 |
| fog_red4 | tinyint | Fog Red Level 4: 0 = None, 255 = Max |
| fog_green4 | tinyint | Fog Green Level 4: 0 = None, 255 = Max |
| fog_blue4 | tinyint | Fog Blue Level 4: 0 = None, 255 = Max |
| fog_minclip4 | float | Fog Minimum Clipping Distance 4 |
| fog_maxclip4 | float | Fog Maximum Clipping Distance 4 |
| fog_density | float | This is the intensity of the fog, this should be a number between 0-1, most commonly used is .1 or .33 |
| flag_needed | varchar | [Flag Required](zone_flags.md) |
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
| rain_chance1 | int | Rain Chance 1 |
| rain_chance2 | int | Rain Chance 2 |
| rain_chance3 | int | Rain Chance 3 |
| rain_chance4 | int | Rain Chance 4 |
| rain_duration1 | int | Rain Duration 1 |
| rain_duration2 | int | Rain Duration 2 |
| rain_duration3 | int | Rain Duration 3 |
| rain_duration4 | int | Rain Duration 4 |
| snow_chance1 | int | Snow Chance 1 |
| snow_chance2 | int | Snow Chance 2 |
| snow_chance3 | int | Snow Chance 3 |
| snow_chance4 | int | Snow Chance 4 |
| snow_duration1 | int | Snow Duration 1 |
| snow_duration2 | int | Snow Duration 2 |
| snow_duration3 | int | Snow Duration 3 |
| snow_duration4 | int | Snow Duration 4 |
| gravity | float | Gravity |
| type | int | Type \(0 = Unknown, 1 = Regular, 2 = Instanced, 3 = Hybrid, 4 = Raid, 5 = City\) |
| skylock | tinyint | Sky Lock |
| fast_regen_hp | int | Fast Regen Health |
| fast_regen_mana | int | Fast Regen Mana |
| fast_regen_endurance | int | Fast Regen Endurance |
| npc_max_aggro_dist | int | NPC Max Aggro Distance |
| max_movement_update_range | int | Max Movement Update Range |

