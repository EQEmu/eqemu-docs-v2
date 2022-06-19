# logsys_categories

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9nc3lzX2NhdGVnb3JpZXMge1xuICAgICAgICB2YXJjaGFyIGRpc2NvcmRfd2ViaG9va19pZFxuICAgIH1cbiAgICBkaXNjb3JkX3dlYmhvb2tzIHtcbiAgICAgICAgdmFyY2hhciBpZFxuICAgIH1cbiAgICBsb2dzeXNfY2F0ZWdvcmllcyB8fC0tb3sgZGlzY29yZF93ZWJob29rcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9nc3lzX2NhdGVnb3JpZXMge1xuICAgICAgICB2YXJjaGFyIGRpc2NvcmRfd2ViaG9va19pZFxuICAgIH1cbiAgICBkaXNjb3JkX3dlYmhvb2tzIHtcbiAgICAgICAgdmFyY2hhciBpZFxuICAgIH1cbiAgICBsb2dzeXNfY2F0ZWdvcmllcyB8fC0tb3sgZGlzY29yZF93ZWJob29rcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9nc3lzX2NhdGVnb3JpZXMge1xuICAgICAgICB2YXJjaGFyIGRpc2NvcmRfd2ViaG9va19pZFxuICAgIH1cbiAgICBkaXNjb3JkX3dlYmhvb2tzIHtcbiAgICAgICAgdmFyY2hhciBpZFxuICAgIH1cbiAgICBsb2dzeXNfY2F0ZWdvcmllcyB8fC0tb3sgZGlzY29yZF93ZWJob29rcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


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

