# banned_ips

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    banned_ips {
        varchar ip_address
    }
    account_ip {
        int accid
        varchar ip
    }
    banned_ips ||--o{ account_ip : "Has-Many"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | ip_address | [account_ip](../../schema/account/account_ip.md) | ip |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| ip_address | varchar | [IP Address](../../schema/account/account_ip.md) |
| notes | varchar | Ban reason |

