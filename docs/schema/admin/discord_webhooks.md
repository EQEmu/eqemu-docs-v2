# discord_webhooks

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZGlzY29yZF93ZWJob29rcyB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBsb2dzeXNfY2F0ZWdvcmllcyB7XG4gICAgICAgIGludCBkaXNjb3JkX3dlYmhvb2tfaWRcbiAgICB9XG4gICAgZGlzY29yZF93ZWJob29rcyB8fC0tb3sgbG9nc3lzX2NhdGVnb3JpZXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZGlzY29yZF93ZWJob29rcyB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBsb2dzeXNfY2F0ZWdvcmllcyB7XG4gICAgICAgIGludCBkaXNjb3JkX3dlYmhvb2tfaWRcbiAgICB9XG4gICAgZGlzY29yZF93ZWJob29rcyB8fC0tb3sgbG9nc3lzX2NhdGVnb3JpZXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZGlzY29yZF93ZWJob29rcyB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBsb2dzeXNfY2F0ZWdvcmllcyB7XG4gICAgICAgIGludCBkaXNjb3JkX3dlYmhvb2tfaWRcbiAgICB9XG4gICAgZGlzY29yZF93ZWJob29rcyB8fC0tb3sgbG9nc3lzX2NhdGVnb3JpZXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


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

