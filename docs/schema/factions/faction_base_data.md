# faction_base_data

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl9iYXNlX2RhdGEge1xuICAgICAgICBzbWFsbGludCBjbGllbnRfZmFjdGlvbl9pZFxuICAgIH1cbiAgICBmYWN0aW9uX2xpc3Qge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgZmFjdGlvbl9iYXNlX2RhdGEgfHwtLW97IGZhY3Rpb25fbGlzdCA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl9iYXNlX2RhdGEge1xuICAgICAgICBzbWFsbGludCBjbGllbnRfZmFjdGlvbl9pZFxuICAgIH1cbiAgICBmYWN0aW9uX2xpc3Qge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgZmFjdGlvbl9iYXNlX2RhdGEgfHwtLW97IGZhY3Rpb25fbGlzdCA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZmFjdGlvbl9iYXNlX2RhdGEge1xuICAgICAgICBzbWFsbGludCBjbGllbnRfZmFjdGlvbl9pZFxuICAgIH1cbiAgICBmYWN0aW9uX2xpc3Qge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgZmFjdGlvbl9iYXNlX2RhdGEgfHwtLW97IGZhY3Rpb25fbGlzdCA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


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

