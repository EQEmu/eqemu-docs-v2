# tasks

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Task Identifier |
| type | tinyint | [Task Type](https://eqemu.gitbook.io/server/categories/task-system-guide/task-types) |
| duration | int | Duration |
| duration\_code | tinyint | [Duration Code](https://eqemu.gitbook.io/server/categories/task-system-guide/task-duration-codes) |
| title | varchar | Title |
| description | text | Description |
| reward | varchar | Reward Description |
| rewardid | int | [Reward Item Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/tasks/items.md) |
| cashreward | int | Cash Reward in Copper |
| xpreward | int | Experience Reward |
| rewardmethod | tinyint | Reward Method: 0 = Single Item ID, 1 = List of Items, 2 = Quest Controlled |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| repeatable | tinyint | Repeatable: 0 = False, 1 = True |
| faction\_reward | int | Faction Reward |
| completion\_emote | varchar | Completion Emote |

