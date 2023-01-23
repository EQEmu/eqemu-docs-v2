# bot_inspect_messages

!!! info
	This page was last generated 2023.01.22

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgaW50dW5zaWduZWQgb3duZXJfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIHNtYWxsaW50IHpvbmVfaWRcbiAgICB9XG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgaW50dW5zaWduZWQgb3duZXJfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIHNtYWxsaW50IHpvbmVfaWRcbiAgICB9XG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgaW50dW5zaWduZWQgb3duZXJfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIHNtYWxsaW50IHpvbmVfaWRcbiAgICB9XG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_id | [bot_data](../../schema/bots/bot_data.md) | bot_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| bot_id | int | [Bot Identifier](bot_data.md) |
| inspect_message | varchar | Inspect Message |

