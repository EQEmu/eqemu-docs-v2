# bot_inventories

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    bot_inventories {
        varchar bot_id
        varchar item_id
        varchar augment_1
        varchar augment_2
        varchar augment_3
        varchar augment_4
        varchar augment_5
        varchar augment_6
    }
    bot_data {
        varchar bot_id
        varchar owner_id
        varchar spells_id
        varchar zone_id
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
    bot_inventories ||--o{ bot_data : "One-to-One"
    bot_inventories ||--o{ items : "One-to-One"
    bot_inventories ||--o{ items : "One-to-One"
    bot_inventories ||--o{ items : "One-to-One"
    bot_inventories ||--o{ items : "One-to-One"
    bot_inventories ||--o{ items : "One-to-One"
    bot_inventories ||--o{ items : "One-to-One"
    bot_inventories ||--o{ items : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_id | [bot_data](../../schema/bots/bot_data.md) | bot_id |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |
| One-to-One | augment_1 | [items](../../schema/items/items.md) | id |
| One-to-One | augment_2 | [items](../../schema/items/items.md) | id |
| One-to-One | augment_3 | [items](../../schema/items/items.md) | id |
| One-to-One | augment_4 | [items](../../schema/items/items.md) | id |
| One-to-One | augment_5 | [items](../../schema/items/items.md) | id |
| One-to-One | augment_6 | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| inventories_index | int | Unique Bot Inventory Identifier |
| bot_id | int | [Bot Identifier](bot_data.md) |
| slot_id | mediumint | [Slot Identifier](../../../../categories/inventory/inventory-slots) |
| item_id | int | [Item Identifier](../items/items.md) |
| inst_charges | smallint | Charges |
| inst_color | int | Color |
| inst_no_drop | tinyint | No Drop: 0 = False, 1=  True |
| inst_custom_data | text | Custom Data |
| ornament_icon | int | Ornamentation Icon |
| ornament_id_file | int | Ornamentation Item Texture |
| ornament_hero_model | int | Ornamentation Hero's Forge Model |
| augment_1 | mediumint | Augment Slot 1 |
| augment_2 | mediumint | Augment Slot 2 |
| augment_3 | mediumint | Augment Slot 3 |
| augment_4 | mediumint | Augment Slot 4 |
| augment_5 | mediumint | Augment Slot 5 |
| augment_6 | mediumint | Augment Slot 6 |

