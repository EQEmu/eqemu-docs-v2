# bot_inspect_messages

## Relationships

```mermaid
erDiagram
    bot_inspect_messages {
        varchar bot_id
    }
    bot_data {
        varchar bot_id
        varchar owner_id
        varchar spells_id
        varchar zone_id
    }
    bot_inspect_messages ||--o{ bot_data : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_id | [bot_data](../../schema/bots/bot_data.md) | bot_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| bot_id | int | [Bot Identifier](bot_data.md) |
| inspect_message | varchar | Inspect Message |

