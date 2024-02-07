# tributes

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdHJpYnV0ZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICB0cmlidXRlX2xldmVscyB7XG4gICAgICAgIGludHVuc2lnbmVkIGl0ZW1faWRcbiAgICAgICAgaW50dW5zaWduZWQgdHJpYnV0ZV9pZFxuICAgIH1cbiAgICB0cmlidXRlcyB8fC0tb3sgdHJpYnV0ZV9sZXZlbHMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdHJpYnV0ZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICB0cmlidXRlX2xldmVscyB7XG4gICAgICAgIGludHVuc2lnbmVkIGl0ZW1faWRcbiAgICAgICAgaW50dW5zaWduZWQgdHJpYnV0ZV9pZFxuICAgIH1cbiAgICB0cmlidXRlcyB8fC0tb3sgdHJpYnV0ZV9sZXZlbHMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdHJpYnV0ZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICB0cmlidXRlX2xldmVscyB7XG4gICAgICAgIGludHVuc2lnbmVkIGl0ZW1faWRcbiAgICAgICAgaW50dW5zaWduZWQgdHJpYnV0ZV9pZFxuICAgIH1cbiAgICB0cmlidXRlcyB8fC0tb3sgdHJpYnV0ZV9sZXZlbHMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [tribute_levels](../../schema/tributes/tribute_levels.md) | tribute_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Tribute Identifier |
| unknown | int | Unknown |
| name | varchar | Name |
| descr | mediumtext | Description |
| isguild | tinyint | Is Guild: 0 = false, 1 = True |

