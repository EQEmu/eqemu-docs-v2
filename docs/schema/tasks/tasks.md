# tasks

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram
[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdGFza3Mge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICB0YXNrX2FjdGl2aXRpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCB0YXNraWRcbiAgICAgICAgaW50dW5zaWduZWQgZGVsaXZlcnRvbnBjXG4gICAgICAgIGludHVuc2lnbmVkIGdvYWxpZFxuICAgIH1cbiAgICB0YXNrc2V0cyB7XG4gICAgICAgIGludHVuc2lnbmVkIHRhc2tpZFxuICAgIH1cbiAgICB0YXNrcyB8fC0tb3sgdGFza19hY3Rpdml0aWVzIDogSGFzLU1hbnlcbiAgICB0YXNrcyB8fC0tb3sgdGFza3NldHMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdGFza3Mge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICB0YXNrX2FjdGl2aXRpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCB0YXNraWRcbiAgICAgICAgaW50dW5zaWduZWQgZGVsaXZlcnRvbnBjXG4gICAgICAgIGludHVuc2lnbmVkIGdvYWxpZFxuICAgIH1cbiAgICB0YXNrc2V0cyB7XG4gICAgICAgIGludHVuc2lnbmVkIHRhc2tpZFxuICAgIH1cbiAgICB0YXNrcyB8fC0tb3sgdGFza19hY3Rpdml0aWVzIDogSGFzLU1hbnlcbiAgICB0YXNrcyB8fC0tb3sgdGFza3NldHMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [task_activities](../../schema/tasks/task_activities.md) | taskid |
| Has-Many | id | [tasksets](../../schema/tasks/tasksets.md) | taskid |


## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Task Identifier |
| type | tinyint | [Task Type](../../../../server/task-system-guide/task-types) |
| duration | int | Duration |
| duration_code | tinyint | [Duration Code](../../../../server/task-system-guide/task-duration-codes) |
| title | varchar | Title |
| description | text | Description |
| reward | varchar | Reward Description |
| rewardid | int | [Reward Item Identifier](../../schema/items/items.md) |
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

