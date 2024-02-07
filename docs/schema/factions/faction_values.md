# faction_values

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl92YWx1ZXMge1xuICAgICAgICBpbnQgY2hhcl9pZFxuICAgICAgICBpbnQgZmFjdGlvbl9pZFxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBmYWN0aW9uX3ZhbHVlcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBcIk9uZS10by1PbmVcIlxuICAgIGZhY3Rpb25fdmFsdWVzIHx8LS1veyBmYWN0aW9uX2xpc3QgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl92YWx1ZXMge1xuICAgICAgICBpbnQgY2hhcl9pZFxuICAgICAgICBpbnQgZmFjdGlvbl9pZFxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBmYWN0aW9uX3ZhbHVlcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBcIk9uZS10by1PbmVcIlxuICAgIGZhY3Rpb25fdmFsdWVzIHx8LS1veyBmYWN0aW9uX2xpc3QgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl92YWx1ZXMge1xuICAgICAgICBpbnQgY2hhcl9pZFxuICAgICAgICBpbnQgZmFjdGlvbl9pZFxuICAgIH1cbiAgICBjaGFyYWN0ZXJfZGF0YSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2luc3RhbmNlXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBmYWN0aW9uX3ZhbHVlcyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBcIk9uZS10by1PbmVcIlxuICAgIGZhY3Rpb25fdmFsdWVzIHx8LS1veyBmYWN0aW9uX2xpc3QgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


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

