# shared_task_members

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc2hhcmVkX3Rhc2tfbWVtYmVycyB7XG4gICAgICAgIGJpZ2ludCBjaGFyYWN0ZXJfaWRcbiAgICAgICAgYmlnaW50IHNoYXJlZF90YXNrX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBzaGFyZWRfdGFza3Mge1xuICAgICAgICBiaWdpbnQgaWRcbiAgICAgICAgaW50IHRhc2tfaWRcbiAgICB9XG4gICAgc2hhcmVkX3Rhc2tfbWVtYmVycyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgc2hhcmVkX3Rhc2tfbWVtYmVycyB8fC0tb3sgc2hhcmVkX3Rhc2tzIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc2hhcmVkX3Rhc2tfbWVtYmVycyB7XG4gICAgICAgIGJpZ2ludCBjaGFyYWN0ZXJfaWRcbiAgICAgICAgYmlnaW50IHNoYXJlZF90YXNrX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBzaGFyZWRfdGFza3Mge1xuICAgICAgICBiaWdpbnQgaWRcbiAgICAgICAgaW50IHRhc2tfaWRcbiAgICB9XG4gICAgc2hhcmVkX3Rhc2tfbWVtYmVycyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgc2hhcmVkX3Rhc2tfbWVtYmVycyB8fC0tb3sgc2hhcmVkX3Rhc2tzIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc2hhcmVkX3Rhc2tfbWVtYmVycyB7XG4gICAgICAgIGJpZ2ludCBjaGFyYWN0ZXJfaWRcbiAgICAgICAgYmlnaW50IHNoYXJlZF90YXNrX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBzaGFyZWRfdGFza3Mge1xuICAgICAgICBiaWdpbnQgaWRcbiAgICAgICAgaW50IHRhc2tfaWRcbiAgICB9XG4gICAgc2hhcmVkX3Rhc2tfbWVtYmVycyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgc2hhcmVkX3Rhc2tfbWVtYmVycyB8fC0tb3sgc2hhcmVkX3Rhc2tzIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | character_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | shared_task_id | [shared_tasks](../../schema/tasks/shared_tasks.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| shared_task_id | bigint | [Shared Task Identifier](shared_tasks.md) |
| character_id | bigint | [Character Identifier](../../schema/characters/character_data.md) |
| is_leader | tinyint | Is Leader: 0 = False, 1 = True |

