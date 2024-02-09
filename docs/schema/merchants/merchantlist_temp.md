# merchantlist_temp

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    merchantlist_temp {
        intunsigned itemid
        intunsigned npcid
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
    merchantlist_temp ||--o{ items : "One-to-One"
    merchantlist_temp ||--o{ npc_types : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | itemid | [items](../../schema/items/items.md) | id |
| One-to-One | npcid | [npc_types](../../schema/npcs/npc_types.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| npcid | int | [NPC Type Identifier](../../schema/npcs/npc_types.md) |
| slot | int | Slot |
| zone_id | int | [Zone Identifier](../../../../server/zones/zone-list) |
| instance_id | int | [Instance Identifier](../../schema/instances/instance_list.md) |
| itemid | int | [Item Identifier](../../schema/items/items.md) |
| charges | int | Charges |

