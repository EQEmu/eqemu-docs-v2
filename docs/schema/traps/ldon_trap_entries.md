# ldon_trap_entries

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbGRvbl90cmFwX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCB0cmFwX2lkXG4gICAgfVxuICAgIGxkb25fdHJhcF90ZW1wbGF0ZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBzbWFsbGludHVuc2lnbmVkIHNwZWxsX2lkXG4gICAgfVxuICAgIGxkb25fdHJhcF9lbnRyaWVzIHx8LS1veyBsZG9uX3RyYXBfdGVtcGxhdGVzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbGRvbl90cmFwX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCB0cmFwX2lkXG4gICAgfVxuICAgIGxkb25fdHJhcF90ZW1wbGF0ZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBzbWFsbGludHVuc2lnbmVkIHNwZWxsX2lkXG4gICAgfVxuICAgIGxkb25fdHJhcF9lbnRyaWVzIHx8LS1veyBsZG9uX3RyYXBfdGVtcGxhdGVzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbGRvbl90cmFwX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCB0cmFwX2lkXG4gICAgfVxuICAgIGxkb25fdHJhcF90ZW1wbGF0ZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBzbWFsbGludHVuc2lnbmVkIHNwZWxsX2lkXG4gICAgfVxuICAgIGxkb25fdHJhcF9lbnRyaWVzIHx8LS1veyBsZG9uX3RyYXBfdGVtcGxhdGVzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | trap_id | [ldon_trap_templates](../../schema/traps/ldon_trap_templates.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique LDoN Trap Entry Identifier |
| trap_id | int | [Trap Identifier](ldon_trap_templates.md) |

