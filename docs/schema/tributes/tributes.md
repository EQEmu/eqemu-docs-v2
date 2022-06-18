# tributes

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdHJpYnV0ZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICB0cmlidXRlX2xldmVscyB7XG4gICAgICAgIGludHVuc2lnbmVkIHRyaWJ1dGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgaXRlbV9pZFxuICAgIH1cbiAgICB0cmlidXRlcyB8fC0tb3sgdHJpYnV0ZV9sZXZlbHMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdHJpYnV0ZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICB0cmlidXRlX2xldmVscyB7XG4gICAgICAgIGludHVuc2lnbmVkIHRyaWJ1dGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgaXRlbV9pZFxuICAgIH1cbiAgICB0cmlidXRlcyB8fC0tb3sgdHJpYnV0ZV9sZXZlbHMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdHJpYnV0ZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICB0cmlidXRlX2xldmVscyB7XG4gICAgICAgIGludHVuc2lnbmVkIHRyaWJ1dGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgaXRlbV9pZFxuICAgIH1cbiAgICB0cmlidXRlcyB8fC0tb3sgdHJpYnV0ZV9sZXZlbHMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [tribute_levels](../../schema/tributes/tribute_levels.md) | tribute_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Tribute Identifier |
| unknown | int | Unknown |
| name | varchar | Name |
| descr | mediumtext | Description |
| isguild | tinyint | Is Guild: 0 = false, 1 = True |

