# qs_player_aa_rate_hourly

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    qs_player_aa_rate_hourly {
        int char_id
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    qs_player_aa_rate_hourly ||--o{ character_data : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_id | [character_data](../../schema/characters/character_data.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| char_id | int | [Character Identifier](../../schema/characters/character_data.md) |
| hour_time | int | Hour Time UNIX Timestamp |
| aa_count | varchar | AA Count |

