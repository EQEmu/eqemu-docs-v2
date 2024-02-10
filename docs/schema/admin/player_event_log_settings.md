# player_event_log_settings

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | bigint | Unique Event Identifier |
| event_name | varchar | Event Name |
| event_enabled | tinyint | Event Enabled: 0 = False, 1 = True |
| retention_days | int | Retention Days: 0 for Permanent |
| discord_webhook_id | int | [Discord Webhook Identifier](../../schema/admin/discord_webhooks.md) |

