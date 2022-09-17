# qs_player_trade_record

!!! info
	This page was last generated 2022.09.16

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX3RyYWRlX3JlY29yZCB7XG4gICAgICAgIGludCBjaGFyMV9pZFxuICAgICAgICBpbnQgY2hhcjJfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgIH1cbiAgICBxc19wbGF5ZXJfdHJhZGVfcmVjb3JkIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcbiAgICBxc19wbGF5ZXJfdHJhZGVfcmVjb3JkIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX3RyYWRlX3JlY29yZCB7XG4gICAgICAgIGludCBjaGFyMV9pZFxuICAgICAgICBpbnQgY2hhcjJfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgIH1cbiAgICBxc19wbGF5ZXJfdHJhZGVfcmVjb3JkIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcbiAgICBxc19wbGF5ZXJfdHJhZGVfcmVjb3JkIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX3RyYWRlX3JlY29yZCB7XG4gICAgICAgIGludCBjaGFyMV9pZFxuICAgICAgICBpbnQgY2hhcjJfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgIH1cbiAgICBxc19wbGF5ZXJfdHJhZGVfcmVjb3JkIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcbiAgICBxc19wbGF5ZXJfdHJhZGVfcmVjb3JkIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


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

