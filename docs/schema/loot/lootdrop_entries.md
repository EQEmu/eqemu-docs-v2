# lootdrop_entries

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdGRyb3BfZW50cmllcyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGxvb3Rkcm9wX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGxvb3Rkcm9wX2VudHJpZXMgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdGRyb3BfZW50cmllcyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGxvb3Rkcm9wX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGxvb3Rkcm9wX2VudHJpZXMgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdGRyb3BfZW50cmllcyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGxvb3Rkcm9wX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGxvb3Rkcm9wX2VudHJpZXMgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| lootdrop_id | int | [Lootdrop Identifier](lootdrop.md) |
| item_id | int | [Item Identifier](../../schema/items/items.md) |
| item_charges | smallint | Item Charges |
| equip_item | tinyint | Equip Item: 0 = False, 1 = True |
| chance | float | Chance: 0 = Never, 100 = Always |
| disabled_chance | float | Disabled Chance: 0 = Never, 100 = Always |
| trivial_min_level | smallint |  |
| trivial_max_level | smallint |  |
| multiplier | tinyint | Multiplier |
| npc_min_level | smallint |  |
| npc_max_level | smallint |  |

