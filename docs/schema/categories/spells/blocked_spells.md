# blocked\_spells

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Blocked Spells Identifier |
| spellid | mediumint | [Spell Identifier](spells_new.md) |
| type | tinyint | [Blocked Spell Type](https://eqemu.gitbook.io/server/categories/spells/blocked-spell-types) |
| zoneid | int | [Zone Identifier](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| x | float | X Coordinate |
| y | float | Y Coordinate |
| z | float | Z Coordinate |
| x\_diff | float | X Radius |
| y\_diff | float | Y Radius |
| z\_diff | float | Z Radius |
| message | varchar | Message when blocked |
| description | varchar | Blocked spells description |

