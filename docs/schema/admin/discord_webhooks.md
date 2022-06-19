# discord_webhooks

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZGlzY29yZF93ZWJob29rcyB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBsb2dzeXNfY2F0ZWdvcmllcyB7XG4gICAgICAgIGludCBkaXNjb3JkX3dlYmhvb2tfaWRcbiAgICB9XG4gICAgZGlzY29yZF93ZWJob29rcyB8fC0tb3sgbG9nc3lzX2NhdGVnb3JpZXMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZGlzY29yZF93ZWJob29rcyB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBsb2dzeXNfY2F0ZWdvcmllcyB7XG4gICAgICAgIGludCBkaXNjb3JkX3dlYmhvb2tfaWRcbiAgICB9XG4gICAgZGlzY29yZF93ZWJob29rcyB8fC0tb3sgbG9nc3lzX2NhdGVnb3JpZXMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZGlzY29yZF93ZWJob29rcyB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBsb2dzeXNfY2F0ZWdvcmllcyB7XG4gICAgICAgIGludCBkaXNjb3JkX3dlYmhvb2tfaWRcbiAgICB9XG4gICAgZGlzY29yZF93ZWJob29rcyB8fC0tb3sgbG9nc3lzX2NhdGVnb3JpZXMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

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

