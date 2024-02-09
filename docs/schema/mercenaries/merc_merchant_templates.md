# merc_merchant_templates

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    merc_merchant_templates {
        varchar merc_merchant_template_id
        varchar qglobal
    }
    merc_merchant_entries {
        varchar merc_merchant_template_id
        varchar merchant_id
    }
    merc_merchant_template_entries {
        varchar merc_merchant_template_id
        varchar merc_template_id
        varchar merc_merchant_template_entry_id
    }
    quest_globals {
        int charid
        varchar name
        int npcid
        int zoneid
    }
    merc_merchant_templates ||--o{ merc_merchant_entries : "Has-Many"
    merc_merchant_templates ||--o{ merc_merchant_template_entries : "Has-Many"
    merc_merchant_templates ||--o{ quest_globals : "Has-Many"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | merc_merchant_template_id | [merc_merchant_entries](../../schema/mercenaries/merc_merchant_entries.md) | merc_merchant_template_id |
| Has-Many | merc_merchant_template_id | [merc_merchant_template_entries](../../schema/mercenaries/merc_merchant_template_entries.md) | merc_merchant_template_id |
| Has-Many | qglobal | [quest_globals](../../schema/data-storage/quest_globals.md) | name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merc_merchant_template_id | int | Unique Mercenary Merchant Template Identifier |
| name | varchar | Name |
| qglobal | varchar | [Quest Global](../data-storage/quest_globals.md) (Deprecated) |

