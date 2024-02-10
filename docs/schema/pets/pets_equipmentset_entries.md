# pets_equipmentset_entries

## Relationships

```mermaid
erDiagram
    pets_equipmentset_entries {
        int item_id
        int set_id
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
    pets_equipmentset {
        int set_id
    }
    pets_equipmentset_entries ||--o{ items : "One-to-One"
    pets_equipmentset_entries ||--o{ pets_equipmentset : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |
| One-to-One | set_id | [pets_equipmentset](../../schema/pets/pets_equipmentset.md) | set_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| set_id | int | [Pet Equipment Set Identifier](pets_equipmentset.md) |
| slot | int | Slot |
| item_id | int | [Item Identifier](../../schema/items/items.md) |

