# aa\_ability

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique AA Identifier |
| name | text | Name |
| category | int | [AA Category](https://eqemu.gitbook.io/server/categories/aas/aa-categories) |
| classes | int | [Classes](https://eqemu.gitbook.io/server/categories/player/class-list) Bitmasks |
| races | int | [Races](https://eqemu.gitbook.io/server/categories/npc/race-list) |
| drakkin\_heritage | int | Drakkin Heritage: 127 = All |
| deities | int | [Deities](https://eqemu.gitbook.io/server/categories/player/deity-list) |
| status | int | [Minimum Status](https://eqemu.gitbook.io/server/categories/player/status-levels) |
| type | int | [AA Type](https://eqemu.gitbook.io/server/categories/aas/aa-types) |
| charges | int | Number of Charges |
| grant\_only | tinyint | Grant Only Flag: 0 = No, 1 = Yes |
| first\_rank\_id | int | First Rank Identifier |
| enabled | tinyint | Enabled: 0 = No, 1 = Yes |
| reset\_on\_death | tinyint |  |

