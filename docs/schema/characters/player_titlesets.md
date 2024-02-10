# player_titlesets

## Relationships

```mermaid
erDiagram
    player_titlesets {
        intunsigned char_id
        intunsigned title_set
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    titles {
        int char_id
        int title_set
        int item_id
    }
    player_titlesets ||--o{ character_data : "One-to-One"
    player_titlesets ||--o{ titles : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_id | [character_data](../../schema/characters/character_data.md) | id |
| Has-Many | title_set | [titles](../../schema/titles/titles.md) | title_set |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Title Set Identifier |
| char_id | int | [Character Identifier](character_data.md) |
| title_set | int | Title Set |

