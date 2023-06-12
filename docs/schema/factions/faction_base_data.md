# faction_base_data

!!! info
	This page was last generated 2023.06.12

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl9iYXNlX2RhdGEge1xuICAgICAgICBzbWFsbGludCBjbGllbnRfZmFjdGlvbl9pZFxuICAgIH1cbiAgICBmYWN0aW9uX2xpc3Qge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgZmFjdGlvbl9iYXNlX2RhdGEgfHwtLW97IGZhY3Rpb25fbGlzdCA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl9iYXNlX2RhdGEge1xuICAgICAgICBzbWFsbGludCBjbGllbnRfZmFjdGlvbl9pZFxuICAgIH1cbiAgICBmYWN0aW9uX2xpc3Qge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgZmFjdGlvbl9iYXNlX2RhdGEgfHwtLW97IGZhY3Rpb25fbGlzdCA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl9iYXNlX2RhdGEge1xuICAgICAgICBzbWFsbGludCBjbGllbnRfZmFjdGlvbl9pZFxuICAgIH1cbiAgICBmYWN0aW9uX2xpc3Qge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgZmFjdGlvbl9iYXNlX2RhdGEgfHwtLW97IGZhY3Rpb25fbGlzdCA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | client_faction_id | [faction_list](../../schema/factions/faction_list.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| client_faction_id | smallint | [Client Faction Identifier](faction_list.md) |
| min | smallint | Minimum Faction |
| max | smallint | Maximum Faction |
| unk_hero1 | smallint | Unknown |
| unk_hero2 | smallint | Unknown |
| unk_hero3 | smallint | Unknown |

