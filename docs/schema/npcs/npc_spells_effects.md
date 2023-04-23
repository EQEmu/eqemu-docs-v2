# npc_spells_effects

!!! info
	This page was last generated 2023.04.23

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX3NwZWxsc19lZmZlY3RzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgbnBjX3NwZWxsc19lZmZlY3RzX2VudHJpZXMge1xuICAgICAgICBpbnQgbnBjX3NwZWxsc19lZmZlY3RzX2lkXG4gICAgfVxuICAgIG5wY19zcGVsbHNfZWZmZWN0cyB8fC0tb3sgbnBjX3NwZWxsc19lZmZlY3RzX2VudHJpZXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX3NwZWxsc19lZmZlY3RzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgbnBjX3NwZWxsc19lZmZlY3RzX2VudHJpZXMge1xuICAgICAgICBpbnQgbnBjX3NwZWxsc19lZmZlY3RzX2lkXG4gICAgfVxuICAgIG5wY19zcGVsbHNfZWZmZWN0cyB8fC0tb3sgbnBjX3NwZWxsc19lZmZlY3RzX2VudHJpZXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX3NwZWxsc19lZmZlY3RzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgbnBjX3NwZWxsc19lZmZlY3RzX2VudHJpZXMge1xuICAgICAgICBpbnQgbnBjX3NwZWxsc19lZmZlY3RzX2lkXG4gICAgfVxuICAgIG5wY19zcGVsbHNfZWZmZWN0cyB8fC0tb3sgbnBjX3NwZWxsc19lZmZlY3RzX2VudHJpZXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [npc_spells_effects_entries](../../schema/npcs/npc_spells_effects_entries.md) | npc_spells_effects_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Spell Effects Identifier |
| name | tinytext | Name |
| parent_list | int | Parent List |

