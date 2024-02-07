# tool_gearup_armor_sets

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCByZWNhc3R0eXBlXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBtZWRpdW1pbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgYm9va1xuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIGludCBwcm9jZWZmZWN0XG4gICAgICAgIGludCBzY3JvbGxlZmZlY3RcbiAgICAgICAgaW50IHdvcm5lZmZlY3RcbiAgICAgICAgaW50IGljb25cbiAgICB9XG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB8fC0tb3sgaXRlbXMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCByZWNhc3R0eXBlXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBtZWRpdW1pbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgYm9va1xuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIGludCBwcm9jZWZmZWN0XG4gICAgICAgIGludCBzY3JvbGxlZmZlY3RcbiAgICAgICAgaW50IHdvcm5lZmZlY3RcbiAgICAgICAgaW50IGljb25cbiAgICB9XG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB8fC0tb3sgaXRlbXMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCByZWNhc3R0eXBlXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBtZWRpdW1pbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgYm9va1xuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIGludCBwcm9jZWZmZWN0XG4gICAgICAgIGludCBzY3JvbGxlZmZlY3RcbiAgICAgICAgaW50IHdvcm5lZmZlY3RcbiAgICAgICAgaW50IGljb25cbiAgICB9XG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB8fC0tb3sgaXRlbXMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| class | tinyint | [Class](../../../../server/player/class-list) |
| level | smallint | Level |
| slot | tinyint | [Slot Identifier](../../../../server/inventory/inventory-slots) |
| item_id | int | [Item Identifier](../../schema/items/items.md) |
| score | mediumint | Score |
| expansion | tinyint | [Expansion](../../../../server/operation/expansion-list) |

