# merc_subtypes

## Relationships

```mermaid
erDiagram
    merc_subtypes {
        varchar merc_subtype_id
    }
    merc_inventory {
        varchar item_id
        varchar merc_subtype_id
    }
    merc_templates {
        varchar dbstring
        varchar merc_template_id
        varchar name_type_id
        varchar merc_npc_type_id
        varchar merc_subtype_id
        varchar merc_type_id
    }
    merc_subtypes ||--o{ merc_inventory : "Has-Many"
    merc_subtypes ||--o{ merc_templates : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | merc_subtype_id | [merc_inventory](../../schema/mercenaries/merc_inventory.md) | merc_subtype_id |
| Has-Many | merc_subtype_id | [merc_templates](../../schema/mercenaries/merc_templates.md) | merc_subtype_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merc_subtype_id | int | Unique Mercenary Subtype Identifier |
| class_id | int | [Class Identifier](../../../../server/player/class-list) |
| tier_id | tinyint | Tier Identifier |
| confidence_id | tinyint | Confidence Identifier |

