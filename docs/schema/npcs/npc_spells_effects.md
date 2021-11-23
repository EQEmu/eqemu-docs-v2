# npc_spells_effects

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram
[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX3NwZWxsc19lZmZlY3RzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgbnBjX3NwZWxsc19lZmZlY3RzX2VudHJpZXMge1xuICAgICAgICBpbnQgbnBjX3NwZWxsc19lZmZlY3RzX2lkXG4gICAgfVxuICAgIG5wY19zcGVsbHNfZWZmZWN0cyB8fC0tb3sgbnBjX3NwZWxsc19lZmZlY3RzX2VudHJpZXMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX3NwZWxsc19lZmZlY3RzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgbnBjX3NwZWxsc19lZmZlY3RzX2VudHJpZXMge1xuICAgICAgICBpbnQgbnBjX3NwZWxsc19lZmZlY3RzX2lkXG4gICAgfVxuICAgIG5wY19zcGVsbHNfZWZmZWN0cyB8fC0tb3sgbnBjX3NwZWxsc19lZmZlY3RzX2VudHJpZXMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [npc_spells_effects_entries](../../schema/npcs/npc_spells_effects_entries.md) | npc_spells_effects_id |


## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Spell Effects Identifier |
| name | tinytext | Name |
| parent_list | int | Parent List |

