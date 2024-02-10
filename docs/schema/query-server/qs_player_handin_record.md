# qs_player_handin_record

## Relationships

```mermaid
erDiagram
    qs_player_handin_record {
        int char_id
        int npc_id
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
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
    qs_player_handin_record ||--o{ character_data : "One-to-One"
    qs_player_handin_record ||--o{ npc_types : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | npc_id | [npc_types](../../schema/npcs/npc_types.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| handin_id | int | Unique Handin Identifier |
| time | timestamp | Time Timestamp |
| quest_id | int | [Quest Identifier](../../schema/tasks/tasks.md) |
| char_id | int | [Character Identifier](../../schema/characters/character_data.md) |
| char_pp | int | Character Platinum |
| char_gp | int | Character Gold |
| char_sp | int | Character Silver |
| char_cp | int | Character Copper |
| char_items | mediumint | [Character Items Identifier](../../schema/items/items.md) |
| npc_id | int | [NPC Type Identifier](../../schema/npcs/npc_types.md) |
| npc_pp | int | NPC Platinum |
| npc_gp | int | NPC Gold |
| npc_sp | int | NPC Silver |
| npc_cp | int | NPC Copper |
| npc_items | mediumint | [NPC Item Identifier](../../schema/items/items.md) |

