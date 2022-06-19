# shared_task_activity_state

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc2hhcmVkX3Rhc2tfYWN0aXZpdHlfc3RhdGUge1xuICAgICAgICBiaWdpbnQgc2hhcmVkX3Rhc2tfaWRcbiAgICAgICAgaW50IGFjdGl2aXR5X2lkXG4gICAgfVxuICAgIHRhc2tzX2FjdGl2aXRpZXMge1xuICAgICAgICB2YXJjaGFyIGFjdGl2aXR5aWRcbiAgICB9XG4gICAgc2hhcmVkX3Rhc2tzIHtcbiAgICAgICAgYmlnaW50IGlkXG4gICAgICAgIGludCB0YXNrX2lkXG4gICAgfVxuICAgIHNoYXJlZF90YXNrX2FjdGl2aXR5X3N0YXRlIHx8LS1veyB0YXNrc19hY3Rpdml0aWVzIDogT25lLXRvLU9uZVxuICAgIHNoYXJlZF90YXNrX2FjdGl2aXR5X3N0YXRlIHx8LS1veyBzaGFyZWRfdGFza3MgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc2hhcmVkX3Rhc2tfYWN0aXZpdHlfc3RhdGUge1xuICAgICAgICBiaWdpbnQgc2hhcmVkX3Rhc2tfaWRcbiAgICAgICAgaW50IGFjdGl2aXR5X2lkXG4gICAgfVxuICAgIHRhc2tzX2FjdGl2aXRpZXMge1xuICAgICAgICB2YXJjaGFyIGFjdGl2aXR5aWRcbiAgICB9XG4gICAgc2hhcmVkX3Rhc2tzIHtcbiAgICAgICAgYmlnaW50IGlkXG4gICAgICAgIGludCB0YXNrX2lkXG4gICAgfVxuICAgIHNoYXJlZF90YXNrX2FjdGl2aXR5X3N0YXRlIHx8LS1veyB0YXNrc19hY3Rpdml0aWVzIDogT25lLXRvLU9uZVxuICAgIHNoYXJlZF90YXNrX2FjdGl2aXR5X3N0YXRlIHx8LS1veyBzaGFyZWRfdGFza3MgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc2hhcmVkX3Rhc2tfYWN0aXZpdHlfc3RhdGUge1xuICAgICAgICBiaWdpbnQgc2hhcmVkX3Rhc2tfaWRcbiAgICAgICAgaW50IGFjdGl2aXR5X2lkXG4gICAgfVxuICAgIHRhc2tzX2FjdGl2aXRpZXMge1xuICAgICAgICB2YXJjaGFyIGFjdGl2aXR5aWRcbiAgICB9XG4gICAgc2hhcmVkX3Rhc2tzIHtcbiAgICAgICAgYmlnaW50IGlkXG4gICAgICAgIGludCB0YXNrX2lkXG4gICAgfVxuICAgIHNoYXJlZF90YXNrX2FjdGl2aXR5X3N0YXRlIHx8LS1veyB0YXNrc19hY3Rpdml0aWVzIDogT25lLXRvLU9uZVxuICAgIHNoYXJlZF90YXNrX2FjdGl2aXR5X3N0YXRlIHx8LS1veyBzaGFyZWRfdGFza3MgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | activity_id | tasks_activities | activityid |
| One-to-One | shared_task_id | [shared_tasks](../../schema/tasks/shared_tasks.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| shared_task_id | bigint | [Shared Task Identifier](shared_tasks.md) |
| activity_id | int | [Activity Identifier](task_activities.md) |
| done_count | int | Done Count |
| updated_time | datetime | Updated Time |
| completed_time | datetime | Completed Time |

