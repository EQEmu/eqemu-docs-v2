# shared_task_activity_state

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc2hhcmVkX3Rhc2tfYWN0aXZpdHlfc3RhdGUge1xuICAgICAgICBiaWdpbnQgc2hhcmVkX3Rhc2tfaWRcbiAgICAgICAgaW50IGFjdGl2aXR5X2lkXG4gICAgfVxuICAgIHRhc2tzX2FjdGl2aXRpZXMge1xuICAgICAgICAgYWN0aXZpdHlpZFxuICAgIH1cbiAgICBzaGFyZWRfdGFza3Mge1xuICAgICAgICBiaWdpbnQgaWRcbiAgICAgICAgaW50IHRhc2tfaWRcbiAgICB9XG4gICAgc2hhcmVkX3Rhc2tfYWN0aXZpdHlfc3RhdGUgfHwtLW97IHRhc2tzX2FjdGl2aXRpZXMgOiBPbmUtdG8tT25lXG4gICAgc2hhcmVkX3Rhc2tfYWN0aXZpdHlfc3RhdGUgfHwtLW97IHNoYXJlZF90YXNrcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc2hhcmVkX3Rhc2tfYWN0aXZpdHlfc3RhdGUge1xuICAgICAgICBiaWdpbnQgc2hhcmVkX3Rhc2tfaWRcbiAgICAgICAgaW50IGFjdGl2aXR5X2lkXG4gICAgfVxuICAgIHRhc2tzX2FjdGl2aXRpZXMge1xuICAgICAgICAgYWN0aXZpdHlpZFxuICAgIH1cbiAgICBzaGFyZWRfdGFza3Mge1xuICAgICAgICBiaWdpbnQgaWRcbiAgICAgICAgaW50IHRhc2tfaWRcbiAgICB9XG4gICAgc2hhcmVkX3Rhc2tfYWN0aXZpdHlfc3RhdGUgfHwtLW97IHRhc2tzX2FjdGl2aXRpZXMgOiBPbmUtdG8tT25lXG4gICAgc2hhcmVkX3Rhc2tfYWN0aXZpdHlfc3RhdGUgfHwtLW97IHNoYXJlZF90YXNrcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc2hhcmVkX3Rhc2tfYWN0aXZpdHlfc3RhdGUge1xuICAgICAgICBiaWdpbnQgc2hhcmVkX3Rhc2tfaWRcbiAgICAgICAgaW50IGFjdGl2aXR5X2lkXG4gICAgfVxuICAgIHRhc2tzX2FjdGl2aXRpZXMge1xuICAgICAgICAgYWN0aXZpdHlpZFxuICAgIH1cbiAgICBzaGFyZWRfdGFza3Mge1xuICAgICAgICBiaWdpbnQgaWRcbiAgICAgICAgaW50IHRhc2tfaWRcbiAgICB9XG4gICAgc2hhcmVkX3Rhc2tfYWN0aXZpdHlfc3RhdGUgfHwtLW97IHRhc2tzX2FjdGl2aXRpZXMgOiBPbmUtdG8tT25lXG4gICAgc2hhcmVkX3Rhc2tfYWN0aXZpdHlfc3RhdGUgfHwtLW97IHNoYXJlZF90YXNrcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

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

