# dynamic_zones

!!! info
	This page was last generated 2022.06.17

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Dynamic Zone Identifier |
| instance_id | int | [Instance Identifier](instance_list.md) |
| type | tinyint | Type |
| uuid | varchar | UUID |
| name | varchar | Name |
| leader_id | int | [Leader Character Identifier](character_data.md) |
| min_players | int | Minimum Players |
| max_players | int | Maximum Players |
| compass_zone_id | int | [Compass Zone Identifier](../../../../server/zones/zone-list) |
| compass_x | float | Compass X Coordinate |
| compass_y | float | Compass Y Coordinate |
| compass_z | float | Compass Z Coordinate |
| safe_return_zone_id | int | [Safe Return Zone Identifier](../../../../server/zones/zone-list) |
| safe_return_x | float | Safe Return X Coordinate |
| safe_return_y | float | Safe Return Y Coordinate |
| safe_return_z | float | Safe Return Z Coordinate |
| safe_return_heading | float | Safe Return Heading Coordinate |
| zone_in_x | float | Zone In X Coordinate |
| zone_in_y | float | Zone In Y Coordinate |
| zone_in_z | float | Zone In Z Coordinate |
| zone_in_heading | float | Zone In Heading Coordinate |
| has_zone_in | tinyint | Has Zone In: 0 = False, 1 = True |

