# discord_webhooks

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZGlzY29yZF93ZWJob29rcyB7XG4gICAgICAgIHZhcmNoYXIgaWRcbiAgICB9XG4gICAgbG9nc3lzX2NhdGVnb3JpZXMge1xuICAgICAgICB2YXJjaGFyIGRpc2NvcmRfd2ViaG9va19pZFxuICAgIH1cbiAgICBkaXNjb3JkX3dlYmhvb2tzIHx8LS1veyBsb2dzeXNfY2F0ZWdvcmllcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZGlzY29yZF93ZWJob29rcyB7XG4gICAgICAgIHZhcmNoYXIgaWRcbiAgICB9XG4gICAgbG9nc3lzX2NhdGVnb3JpZXMge1xuICAgICAgICB2YXJjaGFyIGRpc2NvcmRfd2ViaG9va19pZFxuICAgIH1cbiAgICBkaXNjb3JkX3dlYmhvb2tzIHx8LS1veyBsb2dzeXNfY2F0ZWdvcmllcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZGlzY29yZF93ZWJob29rcyB7XG4gICAgICAgIHZhcmNoYXIgaWRcbiAgICB9XG4gICAgbG9nc3lzX2NhdGVnb3JpZXMge1xuICAgICAgICB2YXJjaGFyIGRpc2NvcmRfd2ViaG9va19pZFxuICAgIH1cbiAgICBkaXNjb3JkX3dlYmhvb2tzIHx8LS1veyBsb2dzeXNfY2F0ZWdvcmllcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


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

