# tool_gearup_armor_sets

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB7XG4gICAgICAgIHZhcmNoYXIgaXRlbV9pZFxuICAgIH1cbiAgICBpdGVtcyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50IGJvb2tcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICAgICAgaW50IGljb25cbiAgICAgICAgc21hbGxpbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIGludCBwcm9jZWZmZWN0XG4gICAgICAgIGludCBzY3JvbGxlZmZlY3RcbiAgICAgICAgaW50IHdvcm5lZmZlY3RcbiAgICB9XG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB7XG4gICAgICAgIHZhcmNoYXIgaXRlbV9pZFxuICAgIH1cbiAgICBpdGVtcyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50IGJvb2tcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICAgICAgaW50IGljb25cbiAgICAgICAgc21hbGxpbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIGludCBwcm9jZWZmZWN0XG4gICAgICAgIGludCBzY3JvbGxlZmZlY3RcbiAgICAgICAgaW50IHdvcm5lZmZlY3RcbiAgICB9XG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB7XG4gICAgICAgIHZhcmNoYXIgaXRlbV9pZFxuICAgIH1cbiAgICBpdGVtcyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50IGJvb2tcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICAgICAgaW50IGljb25cbiAgICAgICAgc21hbGxpbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIGludCBwcm9jZWZmZWN0XG4gICAgICAgIGludCBzY3JvbGxlZmZlY3RcbiAgICAgICAgaW50IHdvcm5lZmZlY3RcbiAgICB9XG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| class | tinyint | [Class](../../../../categories/player/class-list) |
| level | smallint | Level |
| slot | tinyint | [Slot Identifier](../../../../server/inventory/inventory-slots) |
| item_id | int | [Item Identifier](items.md) |
| score | mediumint | Score |
| expansion | tinyint | [Expansion](../../../../server/operation/expansion-list) |

