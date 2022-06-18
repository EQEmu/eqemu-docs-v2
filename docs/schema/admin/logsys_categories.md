# logsys_categories

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9nc3lzX2NhdGVnb3JpZXMge1xuICAgICAgICBpbnQgZGlzY29yZF93ZWJob29rX2lkXG4gICAgfVxuICAgIGRpc2NvcmRfd2ViaG9va3Mge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbG9nc3lzX2NhdGVnb3JpZXMgfHwtLW97IGRpc2NvcmRfd2ViaG9va3MgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9nc3lzX2NhdGVnb3JpZXMge1xuICAgICAgICBpbnQgZGlzY29yZF93ZWJob29rX2lkXG4gICAgfVxuICAgIGRpc2NvcmRfd2ViaG9va3Mge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbG9nc3lzX2NhdGVnb3JpZXMgfHwtLW97IGRpc2NvcmRfd2ViaG9va3MgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9nc3lzX2NhdGVnb3JpZXMge1xuICAgICAgICBpbnQgZGlzY29yZF93ZWJob29rX2lkXG4gICAgfVxuICAgIGRpc2NvcmRfd2ViaG9va3Mge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbG9nc3lzX2NhdGVnb3JpZXMgfHwtLW97IGRpc2NvcmRfd2ViaG9va3MgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | discord_webhook_id | [discord_webhooks](../../schema/admin/discord_webhooks.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| log_category_id | int | Unique Log Category Identifier |
| log_category_description | varchar | Log Category Description |
| log_to_console | smallint | Log to Console: 0 = False, 1 = True |
| log_to_file | smallint | Log to File: 0 = False, 1 = True |
| log_to_gmsay | smallint | Log to GMSay: 0 = False, 1 = True |
| log_to_discord | smallint | Log To Discord: 0 = False, 1 = True |
| discord_webhook_id | int | Unique Webhook Identifier |

