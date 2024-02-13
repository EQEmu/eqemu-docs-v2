# bot_pet_inventories

## Relationships

```mermaid
erDiagram
    bot_pet_inventories {
        varchar pets_index
        varchar item_id
    }
    bot_pets {
        varchar bot_id
        varchar pets_index
        varchar spell_id
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
    bot_pet_inventories ||--o{ bot_pets : "One-to-One"
    bot_pet_inventories ||--o{ items : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | pets_index | [bot_pets](../../schema/bots/bot_pets.md) | pets_index |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| pet_inventories_index | int | Unique Bot Pet Inventory Identifier |
| pets_index | int | [Bot Pet Identifier](bot_pets.md) |
| item_id | int | [Item Identifier](../../schema/items/items.md) |

