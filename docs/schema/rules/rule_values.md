# rule_values

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcnVsZV92YWx1ZXMge1xuICAgICAgICB0aW55aW50dW5zaWduZWQgcnVsZXNldF9pZFxuICAgIH1cbiAgICBydWxlX3NldHMge1xuICAgICAgICB0aW55aW50dW5zaWduZWQgcnVsZXNldF9pZFxuICAgIH1cbiAgICBydWxlX3ZhbHVlcyB8fC0tb3sgcnVsZV9zZXRzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcnVsZV92YWx1ZXMge1xuICAgICAgICB0aW55aW50dW5zaWduZWQgcnVsZXNldF9pZFxuICAgIH1cbiAgICBydWxlX3NldHMge1xuICAgICAgICB0aW55aW50dW5zaWduZWQgcnVsZXNldF9pZFxuICAgIH1cbiAgICBydWxlX3ZhbHVlcyB8fC0tb3sgcnVsZV9zZXRzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcnVsZV92YWx1ZXMge1xuICAgICAgICB0aW55aW50dW5zaWduZWQgcnVsZXNldF9pZFxuICAgIH1cbiAgICBydWxlX3NldHMge1xuICAgICAgICB0aW55aW50dW5zaWduZWQgcnVsZXNldF9pZFxuICAgIH1cbiAgICBydWxlX3ZhbHVlcyB8fC0tb3sgcnVsZV9zZXRzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | ruleset_id | [rule_sets](../../schema/rules/rule_sets.md) | ruleset_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| ruleset_id | tinyint | [Rule Set Identifier](rule_sets.md) |
| rule_name | varchar | Rule Name |
| rule_value | varchar | Rule Value |
| notes | text | Notes |

