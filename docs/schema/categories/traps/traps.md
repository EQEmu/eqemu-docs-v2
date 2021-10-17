# traps

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Trap Identifier |
| zone | varchar | [Zone Short Name](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| version | smallint | Version: -1 For All |
| x | int | X Coordinate |
| y | int | Y Coordinate |
| z | int | Z Coordinate |
| chance | tinyint | Chance: 0 = None, 100 = Always |
| maxzdiff | float | Max Z Difference |
| radius | float | Trap Radius |
| effect | int | [Trap Type](https://eqemu.gitbook.io/server/categories/zones/trap-types) |
| effectvalue | int | Effect Value: \(Based on Trap Type\) 0 = [Spell Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/traps/spells_new.md), 1 = Radius, 2 = [NPC Type Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/traps/npc_types.md), 3 = [NPC Type Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/traps/npc_types.md), 4 = Minimum Damage |
| effectvalue2 | int | Effect Value 2: \(Based on Trap Type\) 0 = Unused, 1 = \(0 = Everything Will Aggro, 1 = Only KoS Will Agro\), 2 = Number of NPCs, 3 = Number of NPCs, 4 = Maximum Damage |
| message | varchar | Message |
| skill | int | [Skill Required](https://eqemu.gitbook.io/server/categories/player/skills) |
| level | mediumint | Level |
| respawn\_time | int | Respawn Timer in Seconds |
| respawn\_var | int | Random Respawn Timer Variance in Seconds |
| triggered\_number | tinyint | Triggered Member |
| group | tinyint | Group |
| despawn\_when\_triggered | tinyint | Despawn When Triggered: 0 = False, 1 = True |
| undetectable | tinyint | Undetectable: 0 = False, 1= True |

