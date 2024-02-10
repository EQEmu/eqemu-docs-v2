# npc_emotes

## Relationships

```mermaid
erDiagram
    npc_emotes {
        intunsigned emoteid
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
    npc_emotes ||--o{ npc_types : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | emoteid | [npc_types](../../schema/npcs/npc_types.md) | emoteid |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique NPC Emote Identifier |
| emoteid | int | Emote Identifier |
| event_ | tinyint | [Emote Event Type Identifier](../../../../server/npc/emote-event-types) |
| type | tinyint | [Emote Type Identifier](../../../../server/npc/emote-types) |
| text | varchar | Text |

