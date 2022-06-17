# lootdrop_entries

!!! info
	This page was last generated 2022.06.17

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdGRyb3BfZW50cmllcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGxvb3Rkcm9wX2lkXG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGxvb3Rkcm9wIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgbG9vdGRyb3BfZW50cmllcyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG4gICAgbG9vdGRyb3BfZW50cmllcyB8fC0tb3sgbG9vdGRyb3AgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdGRyb3BfZW50cmllcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGxvb3Rkcm9wX2lkXG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGxvb3Rkcm9wIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgbG9vdGRyb3BfZW50cmllcyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG4gICAgbG9vdGRyb3BfZW50cmllcyB8fC0tb3sgbG9vdGRyb3AgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdGRyb3BfZW50cmllcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGxvb3Rkcm9wX2lkXG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGxvb3Rkcm9wIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgbG9vdGRyb3BfZW50cmllcyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG4gICAgbG9vdGRyb3BfZW50cmllcyB8fC0tb3sgbG9vdGRyb3AgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |
| One-to-One | lootdrop_id | [lootdrop](../../schema/loot/lootdrop.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| lootdrop_id | int | [Lootdrop Identifier](lootdrop.md) |
| item_id | int | [Item Identifier](items.md) |
| item_charges | smallint | Item Charges |
| equip_item | tinyint | Equip Item: 0 = False, 1 = True |
| chance | float | Chance: 0 = Never, 100 = Always |
| disabled_chance | float | Disabled Chance: 0 = Never, 100 = Always |
| trivial_min_level | smallint | Trivial Minimum Level |
| trivial_max_level | smallint | Trivial Maximum Level |
| multiplier | tinyint | Multiplier |
| npc_min_level | smallint | NPC Minimum Level |
| npc_max_level | smallint | NPC Maximum Level |

