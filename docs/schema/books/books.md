# books

!!! info
	This page was last generated 2022.09.10

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm9va3Mge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGJvb2tcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludCBpY29uXG4gICAgICAgIHNtYWxsaW50IGJhcmRlZmZlY3RcbiAgICAgICAgaW50IGNsaWNrZWZmZWN0XG4gICAgICAgIGludCBmb2N1c2VmZmVjdFxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgICAgIGludCByZWNhc3R0eXBlXG4gICAgfVxuICAgIGJvb2tzIHx8LS1veyBpdGVtcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm9va3Mge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGJvb2tcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludCBpY29uXG4gICAgICAgIHNtYWxsaW50IGJhcmRlZmZlY3RcbiAgICAgICAgaW50IGNsaWNrZWZmZWN0XG4gICAgICAgIGludCBmb2N1c2VmZmVjdFxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgICAgIGludCByZWNhc3R0eXBlXG4gICAgfVxuICAgIGJvb2tzIHx8LS1veyBpdGVtcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm9va3Mge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGJvb2tcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludCBpY29uXG4gICAgICAgIHNtYWxsaW50IGJhcmRlZmZlY3RcbiAgICAgICAgaW50IGNsaWNrZWZmZWN0XG4gICAgICAgIGludCBmb2N1c2VmZmVjdFxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgICAgIGludCByZWNhc3R0eXBlXG4gICAgfVxuICAgIGJvb2tzIHx8LS1veyBpdGVtcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


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

