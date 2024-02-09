# quest_globals

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    quest_globals {
        int charid
        varchar name
        int npcid
        int zoneid
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    item_tick {
        varchar it_itemid
        varchar it_qglobal
    }
    merc_merchant_templates {
        varchar merc_merchant_template_id
        varchar qglobal
    }
    spell_globals {
        varchar qglobal
        int spellid
    }
    npc_types {
        intunsigned alt_currency_id
        int id
        int npc_faction_id
        text name
        intunsigned loottable_id
        intunsigned merchant_id
        intunsigned emoteid
        intunsigned adventure_template_id
        intunsigned armortint_id
        intunsigned npc_spells_id
        intunsigned npc_spells_effects_id
        intunsigned trap_template
    }
    quest_globals ||--o{ character_data : "One-to-One"
    quest_globals ||--o{ item_tick : "Has-Many"
    quest_globals ||--o{ merc_merchant_templates : "Has-Many"
    quest_globals ||--o{ spell_globals : "Has-Many"
    quest_globals ||--o{ npc_types : "One-to-One"


```

```mermaid
erDiagram
    quest_globals {
        int charid
        varchar name
        int npcid
        int zoneid
    }
    zone {
        int zoneidnumber
        varchar short_name
        tinyintunsigned version
        varchar content_flags
        varchar content_flags_disabled
    }
    quest_globals ||--o{ zone : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | charid | [character_data](../../schema/characters/character_data.md) | id |
| Has-Many | name | [item_tick](../../schema/items/item_tick.md) | it_qglobal |
| Has-Many | name | [merc_merchant_templates](../../schema/mercenaries/merc_merchant_templates.md) | qglobal |
| Has-Many | name | [spell_globals](../../schema/spells/spell_globals.md) | qglobal |
| One-to-One | npcid | [npc_types](../../schema/npcs/npc_types.md) | id |
| One-to-One | zoneid | [zone](../../schema/zone/zone.md) | zoneidnumber |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| charid | int | [Character Identifier](../../schema/characters/character_data.md) |
| npcid | int | [NPC Type Identifier](../../schema/npcs/npc_types.md) |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| name | varchar | Name |
| value | varchar | Value |
| expdate | int | Expiration Date UNIX Timestamp |

