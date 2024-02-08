# tool_gearup_armor_sets

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    tool_gearup_armor_sets {
        int item_id
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
    tool_gearup_armor_sets ||--o{ items : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| class | tinyint | [Class](../../../../server/player/class-list) |
| level | smallint | Level |
| slot | tinyint | [Slot Identifier](../../../../server/inventory/inventory-slots) |
| item_id | int | [Item Identifier](../../schema/items/items.md) |
| score | mediumint | Score |
| expansion | tinyint | [Expansion](../../../../server/operation/expansion-list) |

