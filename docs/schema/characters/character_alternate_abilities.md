# character_alternate_abilities

## Relationships

```mermaid
erDiagram
    character_alternate_abilities {
        smallintunsigned aa_id
        intunsigned id
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    aa_ranks {
        intunsigned id
    }
    character_alternate_abilities ||--o{ character_data : "One-to-One"
    character_alternate_abilities ||--o{ aa_ranks : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | aa_id | [aa_ranks](../../schema/aas/aa_ranks.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | [Character Identifier](character_data.md) |
| aa_id | smallint | [AA Identifier](../../schema/aas/aa_ability.md) |
| aa_value | smallint | AA Value |
| charges | smallint | Charges |

