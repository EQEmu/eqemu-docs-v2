# blocked_spells

!!! info
	This page was last generated 2021.11.23

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Blocked Spells Identifier |
| spellid | mediumint | [Spell Identifier](spells_new.md) |
| type | tinyint | [Blocked Spell Type](../../../../server/spells/blocked-spell-types) |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| x | float | X Coordinate |
| y | float | Y Coordinate |
| z | float | Z Coordinate |
| x_diff | float | X Radius |
| y_diff | float | Y Radius |
| z_diff | float | Z Radius |
| message | varchar | Message when blocked |
| description | varchar | Blocked spells description |

