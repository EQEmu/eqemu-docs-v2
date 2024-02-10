# merc_npc_types

## Relationships

```mermaid
erDiagram
    merc_npc_types {
        varchar merc_npc_type_id
    }
    merc_armorinfo {
        varchar armortint_id
        varchar merc_npc_type_id
    }
    merc_stats {
        varchar merc_npc_type_id
    }
    merc_templates {
        varchar dbstring
        varchar merc_template_id
        varchar name_type_id
        varchar merc_npc_type_id
        varchar merc_subtype_id
        varchar merc_type_id
    }
    merc_weaponinfo {
        varchar merc_npc_type_id
    }
    merc_npc_types ||--o{ merc_armorinfo : "Has-Many"
    merc_npc_types ||--o{ merc_stats : "Has-Many"
    merc_npc_types ||--o{ merc_templates : "Has-Many"
    merc_npc_types ||--o{ merc_weaponinfo : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | merc_npc_type_id | [merc_armorinfo](../../schema/mercenaries/merc_armorinfo.md) | merc_npc_type_id |
| Has-Many | merc_npc_type_id | [merc_stats](../../schema/mercenaries/merc_stats.md) | merc_npc_type_id |
| Has-Many | merc_npc_type_id | [merc_templates](../../schema/mercenaries/merc_templates.md) | merc_npc_type_id |
| Has-Many | merc_npc_type_id | [merc_weaponinfo](../../schema/mercenaries/merc_weaponinfo.md) | merc_npc_type_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merc_npc_type_id | int | Mercenary NPC Type Identifier |
| proficiency_id | tinyint | Proficiency Identifier |
| tier_id | tinyint | Tier Identifier |
| class_id | int | [Class Identifier](../../../../server/player/class-list) |
| name | varchar | Name |

