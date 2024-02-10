# bot_stances

## Relationships

```mermaid
erDiagram
    bot_stances {
        varchar bot_id
        varchar stance_id
    }
    bot_data {
        varchar bot_id
        varchar owner_id
        varchar spells_id
        varchar zone_id
    }
    bot_spell_casting_chances {
        varchar spell_type_index
        varchar stance_index
    }
    bot_stances ||--o{ bot_data : "One-to-One"
    bot_stances ||--o{ bot_spell_casting_chances : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_id | [bot_data](../../schema/bots/bot_data.md) | bot_id |
| Has-Many | stance_id | [bot_spell_casting_chances](../../schema/bots/bot_spell_casting_chances.md) | stance_index |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| bot_id | int | [Bot Identifier](bot_data.md) |
| stance_id | tinyint | [Stance Identifier](../../../../server/bots/stance-types) |

