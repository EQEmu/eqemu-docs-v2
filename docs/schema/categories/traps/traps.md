# traps

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Trap Identifier |
| zone | varchar | [Zone Short Name](../../../../categories/zones/zone-list) |
| version | smallint | Version: -1 For All |
| x | int | X Coordinate |
| y | int | Y Coordinate |
| z | int | Z Coordinate |
| chance | tinyint | Chance: 0 = None, 100 = Always |
| maxzdiff | float | Max Z Difference |
| radius | float | Trap Radius |
| effect | int | [Trap Type](../../../../categories/zones/trap-types) |
| effectvalue | int | Effect Value: \(Based on Trap Type\) 0 = [Spell Identifier](../../../schema/categories/spells/spells_new.md), 1 = Radius, 2 = [NPC Type Identifier](../../../schema/categories/npcs/npc_types.md), 3 = [NPC Type Identifier](../../../schema/categories/npcs/npc_types.md), 4 = Minimum Damage |
| effectvalue2 | int | Effect Value 2: \(Based on Trap Type\) 0 = Unused, 1 = \(0 = Everything Will Aggro, 1 = Only KoS Will Agro\), 2 = Number of NPCs, 3 = Number of NPCs, 4 = Maximum Damage |
| message | varchar | Message |
| skill | int | [Skill Required](../../../../categories/player/skills) |
| level | mediumint | Level |
| respawn_time | int | Respawn Timer in Seconds |
| respawn_var | int | Random Respawn Timer Variance in Seconds |
| triggered_number | tinyint | Triggered Member |
| group | tinyint | Group |
| despawn_when_triggered | tinyint | Despawn When Triggered: 0 = False, 1 = True |
| undetectable | tinyint | Undetectable: 0 = False, 1= True |

