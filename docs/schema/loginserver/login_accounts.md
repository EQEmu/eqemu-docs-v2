# login_accounts

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9naW5fYWNjb3VudHMge1xuICAgICAgICB2YXJjaGFyIGFjY291bnRfbmFtZVxuICAgICAgICB2YXJjaGFyIGxhc3RfaXBfYWRkcmVzc1xuICAgIH1cbiAgICBhY2NvdW50IHtcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBhY2NvdW50X2lwIHtcbiAgICAgICAgdmFyY2hhciBpcFxuICAgICAgICBpbnQgYWNjaWRcbiAgICB9XG4gICAgbG9naW5fYWNjb3VudHMgfHwtLW97IGFjY291bnQgOiBPbmUtdG8tT25lXG4gICAgbG9naW5fYWNjb3VudHMgfHwtLW97IGFjY291bnRfaXAgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9naW5fYWNjb3VudHMge1xuICAgICAgICB2YXJjaGFyIGFjY291bnRfbmFtZVxuICAgICAgICB2YXJjaGFyIGxhc3RfaXBfYWRkcmVzc1xuICAgIH1cbiAgICBhY2NvdW50IHtcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBhY2NvdW50X2lwIHtcbiAgICAgICAgdmFyY2hhciBpcFxuICAgICAgICBpbnQgYWNjaWRcbiAgICB9XG4gICAgbG9naW5fYWNjb3VudHMgfHwtLW97IGFjY291bnQgOiBPbmUtdG8tT25lXG4gICAgbG9naW5fYWNjb3VudHMgfHwtLW97IGFjY291bnRfaXAgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9naW5fYWNjb3VudHMge1xuICAgICAgICB2YXJjaGFyIGFjY291bnRfbmFtZVxuICAgICAgICB2YXJjaGFyIGxhc3RfaXBfYWRkcmVzc1xuICAgIH1cbiAgICBhY2NvdW50IHtcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBhY2NvdW50X2lwIHtcbiAgICAgICAgdmFyY2hhciBpcFxuICAgICAgICBpbnQgYWNjaWRcbiAgICB9XG4gICAgbG9naW5fYWNjb3VudHMgfHwtLW97IGFjY291bnQgOiBPbmUtdG8tT25lXG4gICAgbG9naW5fYWNjb3VudHMgfHwtLW97IGFjY291bnRfaXAgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | account_name | [account](../../schema/account/account.md) | name |
| Has-Many | last_ip_address | [account_ip](../../schema/account/account_ip.md) | ip |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Account Identifier |
| account_name | varchar | [Account Name](../../schema/account/account.md) |
| account_password | text | Account Password |
| account_email | varchar | Account Email |
| source_loginserver | varchar | Source Loginserver |
| last_ip_address | varchar | [Last IP Address](../../schema/account/account_ip.md) |
| last_login_date | datetime | Last Login Date |
| created_at | datetime | Created At Date |
| updated_at | datetime | Updated At Date |

