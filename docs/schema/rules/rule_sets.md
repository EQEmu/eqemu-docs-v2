# rule_sets

!!! info
	This page was last generated 2022.05.11

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcnVsZV9zZXRzIHtcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHJ1bGVzZXRfaWRcbiAgICB9XG4gICAgcnVsZV92YWx1ZXMge1xuICAgICAgICB0aW55aW50dW5zaWduZWQgcnVsZXNldF9pZFxuICAgIH1cbiAgICBydWxlX3NldHMgfHwtLW97IHJ1bGVfdmFsdWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcnVsZV9zZXRzIHtcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHJ1bGVzZXRfaWRcbiAgICB9XG4gICAgcnVsZV92YWx1ZXMge1xuICAgICAgICB0aW55aW50dW5zaWduZWQgcnVsZXNldF9pZFxuICAgIH1cbiAgICBydWxlX3NldHMgfHwtLW97IHJ1bGVfdmFsdWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcnVsZV9zZXRzIHtcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHJ1bGVzZXRfaWRcbiAgICB9XG4gICAgcnVsZV92YWx1ZXMge1xuICAgICAgICB0aW55aW50dW5zaWduZWQgcnVsZXNldF9pZFxuICAgIH1cbiAgICBydWxlX3NldHMgfHwtLW97IHJ1bGVfdmFsdWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | ruleset_id | [rule_values](../../schema/rules/rule_values.md) | ruleset_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| ruleset_id | tinyint | Unique Rule Set Identifier |
| name | varchar | Name |

