# tasks

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Task Identifier |
| type | tinyint | [Task Type](../../../../categories/task-system-guide/task-types) |
| duration | int | Duration |
| duration_code | tinyint | [Duration Code](../../../../categories/task-system-guide/task-duration-codes) |
| title | varchar | Title |
| description | text | Description |
| reward | varchar | Reward Description |
| rewardid | int | [Reward Item Identifier](../../../schema/categories/items/items.md) |
| cashreward | int | Cash Reward in Copper |
| xpreward | int | Experience Reward |
| rewardmethod | tinyint | Reward Method: 0 = Single Item ID, 1 = List of Items, 2 = Quest Controlled |
| reward_radiant_crystals | int |  |
| reward_ebon_crystals | int |  |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| level_spread | int |  |
| min_players | int |  |
| max_players | int |  |
| repeatable | tinyint | Repeatable: 0 = False, 1 = True |
| faction_reward | int | Faction Reward |
| completion_emote | varchar | Completion Emote |
| replay_timer_seconds | int |  |
| request_timer_seconds | int |  |

