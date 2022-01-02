# loottable_entries

!!! info
	This page was last generated 2022.01.01

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdHRhYmxlX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBsb290ZHJvcF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICB9XG4gICAgbG9vdGRyb3BfZW50cmllcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGxvb3Rkcm9wX2lkXG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGxvb3R0YWJsZV9lbnRyaWVzIHx8LS1veyBsb290ZHJvcF9lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdHRhYmxlX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBsb290ZHJvcF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICB9XG4gICAgbG9vdGRyb3BfZW50cmllcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGxvb3Rkcm9wX2lkXG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGxvb3R0YWJsZV9lbnRyaWVzIHx8LS1veyBsb290ZHJvcF9lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdHRhYmxlX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBsb290ZHJvcF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICB9XG4gICAgbG9vdGRyb3BfZW50cmllcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGxvb3Rkcm9wX2lkXG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGxvb3R0YWJsZV9lbnRyaWVzIHx8LS1veyBsb290ZHJvcF9lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | lootdrop_id | [lootdrop_entries](../../schema/loot/lootdrop_entries.md) | lootdrop_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| loottable_id | int | [Loottable Identifier](loottable.md) |
| lootdrop_id | int | [Lootdrop Identifier](lootdrop.md) |
| multiplier | tinyint | Multiplier |
| droplimit | tinyint | Maximum Drops |
| mindrop | tinyint | Minimum Drops |
| probability | float | Probability: 0 = Never, 100 = Always |

