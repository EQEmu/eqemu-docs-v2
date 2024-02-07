# gm_ips

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ21faXBzIHtcbiAgICAgICAgaW50IGFjY291bnRfaWRcbiAgICAgICAgdmFyY2hhciBpcF9hZGRyZXNzXG4gICAgfVxuICAgIGFjY291bnQge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgfVxuICAgIGFjY291bnRfaXAge1xuICAgICAgICB2YXJjaGFyIGlwXG4gICAgICAgIGludCBhY2NpZFxuICAgIH1cbiAgICBnbV9pcHMgfHwtLW97IGFjY291bnQgOiBcIk9uZS10by1PbmVcIlxuICAgIGdtX2lwcyB8fC0tb3sgYWNjb3VudF9pcCA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ21faXBzIHtcbiAgICAgICAgaW50IGFjY291bnRfaWRcbiAgICAgICAgdmFyY2hhciBpcF9hZGRyZXNzXG4gICAgfVxuICAgIGFjY291bnQge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgfVxuICAgIGFjY291bnRfaXAge1xuICAgICAgICB2YXJjaGFyIGlwXG4gICAgICAgIGludCBhY2NpZFxuICAgIH1cbiAgICBnbV9pcHMgfHwtLW97IGFjY291bnQgOiBcIk9uZS10by1PbmVcIlxuICAgIGdtX2lwcyB8fC0tb3sgYWNjb3VudF9pcCA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ21faXBzIHtcbiAgICAgICAgaW50IGFjY291bnRfaWRcbiAgICAgICAgdmFyY2hhciBpcF9hZGRyZXNzXG4gICAgfVxuICAgIGFjY291bnQge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgfVxuICAgIGFjY291bnRfaXAge1xuICAgICAgICB2YXJjaGFyIGlwXG4gICAgICAgIGludCBhY2NpZFxuICAgIH1cbiAgICBnbV9pcHMgfHwtLW97IGFjY291bnQgOiBcIk9uZS10by1PbmVcIlxuICAgIGdtX2lwcyB8fC0tb3sgYWNjb3VudF9pcCA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | account_id | [account](../../schema/account/account.md) | id |
| Has-Many | ip_address | [account_ip](../../schema/account/account_ip.md) | ip |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| name | varchar | [Character Name](../../schema/characters/character_data.md) |
| account_id | int | [Account Identifier](../../schema/account/account.md) |
| ip_address | varchar | [IP Address](../../schema/account/account_ip.md) |

