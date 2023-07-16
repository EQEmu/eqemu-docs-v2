# tool_gearup_armor_sets

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgc21hbGxpbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgYm9va1xuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50IHByb2NlZmZlY3RcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICAgICAgaW50IHNjcm9sbGVmZmVjdFxuICAgICAgICBpbnQgd29ybmVmZmVjdFxuICAgICAgICBpbnQgaWNvblxuICAgIH1cbiAgICB0b29sX2dlYXJ1cF9hcm1vcl9zZXRzIHx8LS1veyBpdGVtcyA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgc21hbGxpbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgYm9va1xuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50IHByb2NlZmZlY3RcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICAgICAgaW50IHNjcm9sbGVmZmVjdFxuICAgICAgICBpbnQgd29ybmVmZmVjdFxuICAgICAgICBpbnQgaWNvblxuICAgIH1cbiAgICB0b29sX2dlYXJ1cF9hcm1vcl9zZXRzIHx8LS1veyBpdGVtcyA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgc21hbGxpbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgYm9va1xuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50IHByb2NlZmZlY3RcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICAgICAgaW50IHNjcm9sbGVmZmVjdFxuICAgICAgICBpbnQgd29ybmVmZmVjdFxuICAgICAgICBpbnQgaWNvblxuICAgIH1cbiAgICB0b29sX2dlYXJ1cF9hcm1vcl9zZXRzIHx8LS1veyBpdGVtcyA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


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

