# adventure_template_entry

## Relationships

```mermaid
erDiagram
    adventure_template_entry {
        intunsigned id
        intunsigned template_id
    }
    adventure_details {
        intunsigned id
        smallintunsigned adventure_id
        int instance_id
    }
    adventure_template {
        intunsigned id
        intunsigned graveyard_zone_id
        varchar zone
        smallintunsigned zone_in_zone_id
        tinyintunsigned zone_version
        varchar version
    }
    adventure_template_entry ||--o{ adventure_details : "One-to-One"
    adventure_template_entry ||--o{ adventure_template : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | [adventure_details](../../schema/adventures/adventure_details.md) | id |
| One-to-One | template_id | [adventure_template](../../schema/adventures/adventure_template.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | [Adventure Identifier](adventure_details.md) |
| template_id | int | [Template Identifier](adventure_template.md) |

