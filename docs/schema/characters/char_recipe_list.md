# char_recipe_list

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    char_recipe_list {
        int char_id
        int recipe_id
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    tradeskill_recipe {
        int id
        varchar content_flags
        varchar content_flags_disabled
    }
    char_recipe_list ||--o{ character_data : "One-to-One"
    char_recipe_list ||--o{ tradeskill_recipe : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | recipe_id | [tradeskill_recipe](../../schema/tradeskills/tradeskill_recipe.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| char_id | int | [Character Identifier](character_data.md) |
| recipe_id | int | [Recipe Identifier](../../schema/tradeskills/tradeskill_recipe.md) |
| madecount | int | Made Count |

