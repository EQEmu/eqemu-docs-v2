# faction_values

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl92YWx1ZXMge1xuICAgICAgICBpbnQgY2hhcl9pZFxuICAgICAgICBpbnQgZmFjdGlvbl9pZFxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBmYWN0aW9uX3ZhbHVlcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgZmFjdGlvbl92YWx1ZXMgfHwtLW97IGZhY3Rpb25fbGlzdCA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl92YWx1ZXMge1xuICAgICAgICBpbnQgY2hhcl9pZFxuICAgICAgICBpbnQgZmFjdGlvbl9pZFxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBmYWN0aW9uX3ZhbHVlcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgZmFjdGlvbl92YWx1ZXMgfHwtLW97IGZhY3Rpb25fbGlzdCA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl92YWx1ZXMge1xuICAgICAgICBpbnQgY2hhcl9pZFxuICAgICAgICBpbnQgZmFjdGlvbl9pZFxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIHZhcmNoYXIgbmFuZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBmYWN0aW9uX3ZhbHVlcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgZmFjdGlvbl92YWx1ZXMgfHwtLW97IGZhY3Rpb25fbGlzdCA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | faction_id | [faction_list](../../schema/factions/faction_list.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| char_id | int | [Character Identifier](../../schema/characters/character_data.md) |
| faction_id | int | [Faction Identifier](faction_list.md) |
| current_value | smallint | Current Value |
| temp | tinyint | Temporary: 0 = False, 1 = True |

