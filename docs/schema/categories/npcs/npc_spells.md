# npc\_spells

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique NPC Spell Set Identifier |
| name | tinytext | NPC Spell Set Name |
| parent\_list | int | Inherit all the spells from this list, and merge them with our spells. Only one level of inheritance is allowed, so your parent's parent will not be included. |
| attack\_proc | smallint | The combat proc that an NPC with this spell set will add to their list of procs. [Spell Identifier](../../../schema/categories/npcs/spells_new.md) |
| proc\_chance | tinyint | Proc Chance: 0 = Never, 100 = Always |
| range\_proc | smallint | The ranged proc that an NPC with this spell set will add to their list of procs. |
| rproc\_chance | smallint | Ranged Proc Chance: 0 = Never, 100 = Always |
| defensive\_proc | smallint | The defensive proc that an NPC with this spell set will add to their list of procs. |
| dproc\_chance | smallint | Defensive Proc Chance: 0 = Never, 100 = Always |
| fail\_recast | int | Fail Recast |
| engaged\_no\_sp\_recast\_min | int |  |
| engaged\_no\_sp\_recast\_max | int |  |
| engaged\_b\_self\_chance | tinyint |  |
| engaged\_b\_other\_chance | tinyint |  |
| engaged\_d\_chance | tinyint |  |
| pursue\_no\_sp\_recast\_min | int |  |
| pursue\_no\_sp\_recast\_max | int |  |
| pursue\_d\_chance | tinyint |  |
| idle\_no\_sp\_recast\_min | int |  |
| idle\_no\_sp\_recast\_max | int |  |
| idle\_b\_chance | tinyint |  |

