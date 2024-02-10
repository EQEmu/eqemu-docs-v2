# discord_webhooks

## Relationships

```mermaid
erDiagram
    discord_webhooks {
        int id
    }
    logsys_categories {
        int discord_webhook_id
    }
    discord_webhooks ||--o{ logsys_categories : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [logsys_categories](../../schema/admin/logsys_categories.md) | discord_webhook_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Webhook Identifier |
| webhook_name | varchar | Webhook Name |
| webhook_url | varchar | Webhook URL |
| created_at | datetime | Created At |
| deleted_at | datetime | Deleted At |

