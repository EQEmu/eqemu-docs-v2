# bot_inspect_messages

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgc21hbGxpbnQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBvd25lcl9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBzcGVsbHNfaWRcbiAgICB9XG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgc21hbGxpbnQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBvd25lcl9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBzcGVsbHNfaWRcbiAgICB9XG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgc21hbGxpbnQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBvd25lcl9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBzcGVsbHNfaWRcbiAgICB9XG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_id | [bot_data](../../schema/bots/bot_data.md) | bot_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| bot_id | int | [Bot Identifier](bot_data.md) |
| inspect_message | varchar | Inspect Message |

