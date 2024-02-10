# trader_audit

## Relationships

```mermaid
erDiagram
    trader_audit {
        varchar buyer
        varchar seller
        varchar itemname
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    items {
        int id
        int book
        varchar name
        int recasttype
        int icon
        mediumint bardeffect
        int clickeffect
        int focuseffect
        int proceffect
        int scrolleffect
        int worneffect
    }
    trader_audit ||--o{ character_data : "One-to-One"
    trader_audit ||--o{ items : "One-to-One"
    trader_audit ||--o{ character_data : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | buyer | [character_data](../../schema/characters/character_data.md) | name |
| One-to-One | itemname | [items](../../schema/items/items.md) | name |
| One-to-One | seller | [character_data](../../schema/characters/character_data.md) | name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| time | datetime | Time |
| seller | varchar | Seller |
| buyer | varchar | Buyer |
| itemname | varchar | [Item Name](../../schema/items/items.md) |
| quantity | int | Quantity |
| totalcost | int | Total Cost |
| trantype | tinyint | Transaction Type |

