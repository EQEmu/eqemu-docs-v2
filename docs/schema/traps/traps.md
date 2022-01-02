# traps

!!! info
	This page was last generated 2022.01.01

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Trap Identifier |
| zone | varchar | [Zone Short Name](../../../../server/zones/zone-list) |
| version | smallint | Version: -1 For All |
| x | int | X Coordinate |
| y | int | Y Coordinate |
| z | int | Z Coordinate |
| chance | tinyint | Chance: 0 = None, 100 = Always |
| maxzdiff | float | Max Z Difference |
| radius | float | Trap Radius |
| effect | int | [Trap Type](../../../../server/zones/trap-types) |
| effectvalue | int | Effect Value: (Based on Trap Type) 0 = [Spell Identifier](../../schema/spells/spells_new.md), 1 = Radius, 2 = [NPC Type Identifier](../../schema/npcs/npc_types.md), 3 = [NPC Type Identifier](../../schema/npcs/npc_types.md), 4 = Minimum Damage |
| effectvalue2 | int | Effect Value 2: (Based on Trap Type) 0 = Unused, 1 = (0 = Everything Will Aggro, 1 = Only KoS Will Agro), 2 = Number of NPCs, 3 = Number of NPCs, 4 = Maximum Damage |
| message | varchar | Message |
| skill | int | [Skill Required](../../../../server/player/skills) |
| level | mediumint | Level |
| respawn_time | int | Respawn Timer in Seconds |
| respawn_var | int | Random Respawn Timer Variance in Seconds |
| triggered_number | tinyint | Triggered Member |
| group | tinyint | Group |
| despawn_when_triggered | tinyint | Despawn When Triggered: 0 = False, 1 = True |
| undetectable | tinyint | Undetectable: 0 = False, 1= True |
| min_expansion | tinyint |  |
| max_expansion | tinyint |  |
| content_flags | varchar |  |
| content_flags_disabled | varchar |  |

