# qs_player_trade_record

!!! info
	This page was last generated 2023.06.12

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX3RyYWRlX3JlY29yZCB7XG4gICAgICAgIGludCBjaGFyMV9pZFxuICAgICAgICBpbnQgY2hhcjJfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgIH1cbiAgICBxc19wbGF5ZXJfdHJhZGVfcmVjb3JkIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IFwiT25lLXRvLU9uZVwiXG4gICAgcXNfcGxheWVyX3RyYWRlX3JlY29yZCB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX3RyYWRlX3JlY29yZCB7XG4gICAgICAgIGludCBjaGFyMV9pZFxuICAgICAgICBpbnQgY2hhcjJfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgIH1cbiAgICBxc19wbGF5ZXJfdHJhZGVfcmVjb3JkIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IFwiT25lLXRvLU9uZVwiXG4gICAgcXNfcGxheWVyX3RyYWRlX3JlY29yZCB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX3RyYWRlX3JlY29yZCB7XG4gICAgICAgIGludCBjaGFyMV9pZFxuICAgICAgICBpbnQgY2hhcjJfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgIH1cbiAgICBxc19wbGF5ZXJfdHJhZGVfcmVjb3JkIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IFwiT25lLXRvLU9uZVwiXG4gICAgcXNfcGxheWVyX3RyYWRlX3JlY29yZCB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char1_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | char2_id | [character_data](../../schema/characters/character_data.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| trade_id | int | Unique Trade Identifier |
| time | timestamp | Time Timestamp |
| char1_id | int | [Character 1 Identifier](../../schema/characters/character_data.md) |
| char1_pp | int | Character 1 Platinum |
| char1_gp | int | Character 1 Gold |
| char1_sp | int | Character 1 Silver |
| char1_cp | int | Character 1 Copper |
| char1_items | mediumint | [Character 1 Item Identifier](../../schema/items/items.md) |
| char2_id | int | [Character 2 Identifier](../../schema/characters/character_data.md) |
| char2_pp | int | Character 2 Platinum |
| char2_gp | int | Character 2 Gold |
| char2_sp | int | Character 2 Silver |
| char2_cp | int | Character 2 Copper |
| char2_items | mediumint | [Character 2 Item Identifier](../../schema/items/items.md) |

