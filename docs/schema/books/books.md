# books

!!! info
	This page was last generated 2023.04.23

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm9va3Mge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICAgICAgaW50IGJvb2tcbiAgICAgICAgc21hbGxpbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgICAgIGludCBpY29uXG4gICAgfVxuICAgIGJvb2tzIHx8LS1veyBpdGVtcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm9va3Mge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICAgICAgaW50IGJvb2tcbiAgICAgICAgc21hbGxpbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgICAgIGludCBpY29uXG4gICAgfVxuICAgIGJvb2tzIHx8LS1veyBpdGVtcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm9va3Mge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICAgICAgaW50IGJvb2tcbiAgICAgICAgc21hbGxpbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgICAgIGludCBpY29uXG4gICAgfVxuICAgIGJvb2tzIHx8LS1veyBpdGVtcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [items](../../schema/items/items.md) | book |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Book Identifier |
| name | varchar | Unique Book Identifier |
| txtfile | text | The text in the book. ` Represents line spaces, `` is two line spaces, ``` is three line spaces, etc. (13 lines per book page) |
| language | int | [Language](../../../../server/player/languages) |

