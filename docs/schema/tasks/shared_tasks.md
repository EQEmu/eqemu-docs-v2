# shared_tasks

!!! info
	This page was last generated 2022.06.17

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | bigint | Unique Shared Task Identifier |
| task_id | int | [Task Identifier](tasks.md) |
| accepted_time | datetime | Accepted Time |
| expire_time | datetime | Expire Time |
| completion_time | datetime | Completion Time |
| is_locked | tinyint | Is Locked: 0 = False, 1 = True |

