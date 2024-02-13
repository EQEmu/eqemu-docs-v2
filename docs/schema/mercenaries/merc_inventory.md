# merc_inventory

## Relationships

```mermaid
erDiagram
    merc_inventory {
        varchar item_id
        varchar merc_subtype_id
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
    merc_subtypes {
        varchar merc_subtype_id
    }
    merc_inventory ||--o{ items : "One-to-One"
    merc_inventory ||--o{ merc_subtypes : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |
| One-to-One | merc_subtype_id | [merc_subtypes](../../schema/mercenaries/merc_subtypes.md) | merc_subtype_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merc_inventory_id | int | Unique Mercenary Inventory Identifier |
| merc_subtype_id | int | [Mercenary Subtype Identifier](merc_subtypes.md) |
| item_id | int | [Item Identifier](../../schema/items/items.md) |
| min_level | int | Minimum Level |
| max_level | int | Maximum Level |

