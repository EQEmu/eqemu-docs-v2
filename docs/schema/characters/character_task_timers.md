# character_task_timers

!!! info
	This page was last generated 2022.06.17

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Task Timer Identifier |
| character_id | int | [Character Identifer](character_data.md) |
| task_id | int | [Task Identifier](../../schema/tasks/tasks.md) |
| timer_type | int | Timer Type (0 = Replay, 1 = Request) |
| expire_time | datetime | Expire Time |

