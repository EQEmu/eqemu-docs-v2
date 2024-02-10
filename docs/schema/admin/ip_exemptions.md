# ip_exemptions

## Relationships

```mermaid
erDiagram
    ip_exemptions {
        varchar exemption_ip
    }
    account_ip {
        int accid
        varchar ip
    }
    ip_exemptions ||--o{ account_ip : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | exemption_ip | [account_ip](../../schema/account/account_ip.md) | ip |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| exemption_id | int | Exemption Identifier |
| exemption_ip | varchar | [Exemption IP Address](../../schema/account/account_ip.md) |
| exemption_amount | int | Exemption Amount |

