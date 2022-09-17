# completed_shared_tasks

!!! info
	This page was last generated 2022.09.16

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY29tcGxldGVkX3NoYXJlZF90YXNrcyB7XG4gICAgICAgIGludCB0YXNrX2lkXG4gICAgfVxuICAgIHRhc2tzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdGlueWludCB0eXBlXG4gICAgICAgIGludHVuc2lnbmVkIGR6X3RlbXBsYXRlX2lkXG4gICAgfVxuICAgIGNvbXBsZXRlZF9zaGFyZWRfdGFza3MgfHwtLW97IHRhc2tzIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY29tcGxldGVkX3NoYXJlZF90YXNrcyB7XG4gICAgICAgIGludCB0YXNrX2lkXG4gICAgfVxuICAgIHRhc2tzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdGlueWludCB0eXBlXG4gICAgICAgIGludHVuc2lnbmVkIGR6X3RlbXBsYXRlX2lkXG4gICAgfVxuICAgIGNvbXBsZXRlZF9zaGFyZWRfdGFza3MgfHwtLW97IHRhc2tzIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY29tcGxldGVkX3NoYXJlZF90YXNrcyB7XG4gICAgICAgIGludCB0YXNrX2lkXG4gICAgfVxuICAgIHRhc2tzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdGlueWludCB0eXBlXG4gICAgICAgIGludHVuc2lnbmVkIGR6X3RlbXBsYXRlX2lkXG4gICAgfVxuICAgIGNvbXBsZXRlZF9zaGFyZWRfdGFza3MgfHwtLW97IHRhc2tzIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | task_id | [tasks](../../schema/tasks/tasks.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | bigint | [Shared Task Identifier](shared_tasks.md) |
| task_id | int | [Shared Task Identifier](shared_tasks.md) |
| accepted_time | datetime | Accepted Time |
| expire_time | datetime | Expire TIme |
| completion_time | datetime | Completion TIme |
| is_locked | tinyint | Is Locked: 0 = False, 1 = True |

