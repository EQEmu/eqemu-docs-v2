# npc_spells

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique NPC Spell Set Identifier |
| name | tinytext | NPC Spell Set Name |
| parent_list | int | Inherit all the spells from this list, and merge them with our spells. Only one level of inheritance is allowed, so your parent's parent will not be included. |
| attack_proc | smallint | The combat proc that an NPC with this spell set will add to their list of procs. [Spell Identifier](../../schema/spells/spells_new.md) |
| proc_chance | tinyint | Proc Chance: 0 = Never, 100 = Always |
| range_proc | smallint | The ranged proc that an NPC with this spell set will add to their list of procs. |
| rproc_chance | smallint | Ranged Proc Chance: 0 = Never, 100 = Always |
| defensive_proc | smallint | The defensive proc that an NPC with this spell set will add to their list of procs. |
| dproc_chance | smallint | Defensive Proc Chance: 0 = Never, 100 = Always |
| fail_recast | int | Fail Recast |
| engaged_no_sp_recast_min | int |  |
| engaged_no_sp_recast_max | int |  |
| engaged_b_self_chance | tinyint |  |
| engaged_b_other_chance | tinyint |  |
| engaged_d_chance | tinyint |  |
| pursue_no_sp_recast_min | int |  |
| pursue_no_sp_recast_max | int |  |
| pursue_d_chance | tinyint |  |
| idle_no_sp_recast_min | int |  |
| idle_no_sp_recast_max | int |  |
| idle_b_chance | tinyint |  |

