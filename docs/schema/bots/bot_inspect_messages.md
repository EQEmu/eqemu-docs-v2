# bot_inspect_messages

!!! info
	This page was last generated 2023.01.30

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIHNtYWxsaW50IHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgb3duZXJfaWRcbiAgICB9XG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMgfHwtLW97IGJvdF9kYXRhIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIHNtYWxsaW50IHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgb3duZXJfaWRcbiAgICB9XG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMgfHwtLW97IGJvdF9kYXRhIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIHNtYWxsaW50IHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgb3duZXJfaWRcbiAgICB9XG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMgfHwtLW97IGJvdF9kYXRhIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_id | [bot_data](../../schema/bots/bot_data.md) | bot_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| bot_id | int | [Bot Identifier](bot_data.md) |
| inspect_message | varchar | Inspect Message |

