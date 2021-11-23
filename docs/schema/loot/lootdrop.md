# lootdrop

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdGRyb3Age1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBsb290ZHJvcF9lbnRyaWVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbG9vdGRyb3BfaWRcbiAgICAgICAgaW50IGl0ZW1faWRcbiAgICB9XG4gICAgbG9vdGRyb3AgfHwtLW97IGxvb3Rkcm9wX2VudHJpZXMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdGRyb3Age1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBsb290ZHJvcF9lbnRyaWVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbG9vdGRyb3BfaWRcbiAgICAgICAgaW50IGl0ZW1faWRcbiAgICB9XG4gICAgbG9vdGRyb3AgfHwtLW97IGxvb3Rkcm9wX2VudHJpZXMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [lootdrop_entries](../../schema/loot/lootdrop_entries.md) | lootdrop_id |


## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Lootdrop Identifier |
| name | varchar | Name |
| min_expansion | tinyint |  |
| max_expansion | tinyint |  |
| content_flags | varchar |  |
| content_flags_disabled | varchar |  |

