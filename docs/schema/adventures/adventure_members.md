# adventure_members

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    adventure_members {
        intunsigned id
        intunsigned charid
    }
    adventure_details {
        intunsigned id
        smallintunsigned adventure_id
        int instance_id
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    adventure_members ||--o{ adventure_details : "One-to-One"
    adventure_members ||--o{ character_data : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | [adventure_details](../../schema/adventures/adventure_details.md) | id |
| One-to-One | charid | [character_data](../../schema/characters/character_data.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | [Adventure Identifier](adventure_details.md) |
| charid | int | [Character Identifier](../../schema/characters/character_data.md) |

