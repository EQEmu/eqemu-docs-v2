# bot_timers

!!! info
	This page was last generated 2022.09.16

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3RpbWVycyB7XG4gICAgICAgIGludHVuc2lnbmVkIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIGJvdF9kYXRhXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgICAgIGludHVuc2lnbmVkIHNwZWxsc19pZFxuICAgICAgICBzbWFsbGludCB6b25lX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG93bmVyX2lkXG4gICAgfVxuICAgIGJvdF90aW1lcnMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3RpbWVycyB7XG4gICAgICAgIGludHVuc2lnbmVkIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIGJvdF9kYXRhXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgICAgIGludHVuc2lnbmVkIHNwZWxsc19pZFxuICAgICAgICBzbWFsbGludCB6b25lX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG93bmVyX2lkXG4gICAgfVxuICAgIGJvdF90aW1lcnMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3RpbWVycyB7XG4gICAgICAgIGludHVuc2lnbmVkIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIGJvdF9kYXRhXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgICAgIGludHVuc2lnbmVkIHNwZWxsc19pZFxuICAgICAgICBzbWFsbGludCB6b25lX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG93bmVyX2lkXG4gICAgfVxuICAgIGJvdF90aW1lcnMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


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

