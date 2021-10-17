# spawn\_events

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Spawn Event Entry Identifier |
| zone | varchar | [Zone Short Name](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| cond\_id | mediumint | [Spawn Condition Identifier](spawn_conditions.md) |
| name | varchar | Name |
| period | int | Period |
| next\_minute | tinyint | Next Minute |
| next\_hour | tinyint | Next Hour |
| next\_day | tinyint | Next Day |
| next\_month | tinyint | Next Month |
| next\_year | int | Next Year |
| enabled | tinyint | Enabled: 0 = False, 1 = True |
| action | tinyint | [Action Type](https://eqemu.gitbook.io/server/categories/npc/spawns/action-types) |
| argument | mediumint | Argument: \(Based on Action\) 0 = Argument Value |
| strict | tinyint | Strict Date Criteria: 0 = False, 1 = True |

