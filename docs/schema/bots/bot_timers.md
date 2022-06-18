# bot_timers

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3RpbWVycyB7XG4gICAgICAgIGludHVuc2lnbmVkIGJvdF9pZFxuICAgICAgICAgYm90X2RhdGFcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBzcGVsbHNfaWRcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgICAgIHNtYWxsaW50IHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgb3duZXJfaWRcbiAgICB9XG4gICAgYm90X3RpbWVycyB8fC0tb3sgYm90X2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3RpbWVycyB7XG4gICAgICAgIGludHVuc2lnbmVkIGJvdF9pZFxuICAgICAgICAgYm90X2RhdGFcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBzcGVsbHNfaWRcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgICAgIHNtYWxsaW50IHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgb3duZXJfaWRcbiAgICB9XG4gICAgYm90X3RpbWVycyB8fC0tb3sgYm90X2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3RpbWVycyB7XG4gICAgICAgIGludHVuc2lnbmVkIGJvdF9pZFxuICAgICAgICAgYm90X2RhdGFcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBzcGVsbHNfaWRcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgICAgIHNtYWxsaW50IHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgb3duZXJfaWRcbiAgICB9XG4gICAgYm90X3RpbWVycyB8fC0tb3sgYm90X2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_data | [bot_data](../../schema/bots/bot_data.md) | bot_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| bot_id | int | [Unique Bot Identifier](bot_data.md) |
| timer_id | int | Timer Identifier |
| timer_value | int | Timer Expiration |

