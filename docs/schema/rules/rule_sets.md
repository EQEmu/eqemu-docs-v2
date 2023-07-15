# rule_sets

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcnVsZV9zZXRzIHtcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHJ1bGVzZXRfaWRcbiAgICB9XG4gICAgcnVsZV92YWx1ZXMge1xuICAgICAgICB0aW55aW50dW5zaWduZWQgcnVsZXNldF9pZFxuICAgIH1cbiAgICBydWxlX3NldHMgfHwtLW97IHJ1bGVfdmFsdWVzIDogXCJIYXMtTWFueVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcnVsZV9zZXRzIHtcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHJ1bGVzZXRfaWRcbiAgICB9XG4gICAgcnVsZV92YWx1ZXMge1xuICAgICAgICB0aW55aW50dW5zaWduZWQgcnVsZXNldF9pZFxuICAgIH1cbiAgICBydWxlX3NldHMgfHwtLW97IHJ1bGVfdmFsdWVzIDogXCJIYXMtTWFueVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcnVsZV9zZXRzIHtcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHJ1bGVzZXRfaWRcbiAgICB9XG4gICAgcnVsZV92YWx1ZXMge1xuICAgICAgICB0aW55aW50dW5zaWduZWQgcnVsZXNldF9pZFxuICAgIH1cbiAgICBydWxlX3NldHMgfHwtLW97IHJ1bGVfdmFsdWVzIDogXCJIYXMtTWFueVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | ruleset_id | [rule_values](../../schema/rules/rule_values.md) | ruleset_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| ruleset_id | tinyint | Unique Rule Set Identifier |
| name | varchar | Name |

