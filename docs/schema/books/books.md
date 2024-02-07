# books

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm9va3Mge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIG1lZGl1bWludCBiYXJkZWZmZWN0XG4gICAgICAgIGludCBib29rXG4gICAgICAgIGludCBjbGlja2VmZmVjdFxuICAgICAgICBpbnQgZm9jdXNlZmZlY3RcbiAgICAgICAgaW50IHByb2NlZmZlY3RcbiAgICAgICAgaW50IHNjcm9sbGVmZmVjdFxuICAgICAgICBpbnQgd29ybmVmZmVjdFxuICAgICAgICBpbnQgaWNvblxuICAgIH1cbiAgICBib29rcyB8fC0tb3sgaXRlbXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm9va3Mge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIG1lZGl1bWludCBiYXJkZWZmZWN0XG4gICAgICAgIGludCBib29rXG4gICAgICAgIGludCBjbGlja2VmZmVjdFxuICAgICAgICBpbnQgZm9jdXNlZmZlY3RcbiAgICAgICAgaW50IHByb2NlZmZlY3RcbiAgICAgICAgaW50IHNjcm9sbGVmZmVjdFxuICAgICAgICBpbnQgd29ybmVmZmVjdFxuICAgICAgICBpbnQgaWNvblxuICAgIH1cbiAgICBib29rcyB8fC0tb3sgaXRlbXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm9va3Mge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIG1lZGl1bWludCBiYXJkZWZmZWN0XG4gICAgICAgIGludCBib29rXG4gICAgICAgIGludCBjbGlja2VmZmVjdFxuICAgICAgICBpbnQgZm9jdXNlZmZlY3RcbiAgICAgICAgaW50IHByb2NlZmZlY3RcbiAgICAgICAgaW50IHNjcm9sbGVmZmVjdFxuICAgICAgICBpbnQgd29ybmVmZmVjdFxuICAgICAgICBpbnQgaWNvblxuICAgIH1cbiAgICBib29rcyB8fC0tb3sgaXRlbXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


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

