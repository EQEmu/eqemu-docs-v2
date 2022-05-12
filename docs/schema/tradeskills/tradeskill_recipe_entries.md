# tradeskill_recipe_entries

!!! info
	This page was last generated 2022.05.11

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdHJhZGVza2lsbF9yZWNpcGVfZW50cmllcyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgICAgIGludCByZWNpcGVfaWRcbiAgICB9XG4gICAgdHJhZGVza2lsbF9yZWNpcGUge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgdHJhZGVza2lsbF9yZWNpcGVfZW50cmllcyB8fC0tb3sgdHJhZGVza2lsbF9yZWNpcGUgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdHJhZGVza2lsbF9yZWNpcGVfZW50cmllcyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgICAgIGludCByZWNpcGVfaWRcbiAgICB9XG4gICAgdHJhZGVza2lsbF9yZWNpcGUge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgdHJhZGVza2lsbF9yZWNpcGVfZW50cmllcyB8fC0tb3sgdHJhZGVza2lsbF9yZWNpcGUgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdHJhZGVza2lsbF9yZWNpcGVfZW50cmllcyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgICAgIGludCByZWNpcGVfaWRcbiAgICB9XG4gICAgdHJhZGVza2lsbF9yZWNpcGUge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgdHJhZGVza2lsbF9yZWNpcGVfZW50cmllcyB8fC0tb3sgdHJhZGVza2lsbF9yZWNpcGUgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | recipe_id | [tradeskill_recipe](../../schema/tradeskills/tradeskill_recipe.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Tradeskill Recipe Entry Identifier |
| recipe_id | int | [Unique Tradeskill Recipe Identifier](tradeskill_recipe.md) |
| item_id | int | [Item Identifier](../../schema/items/items.md) |
| successcount | tinyint | Success Count |
| failcount | tinyint | Fail Count |
| componentcount | tinyint | Component Count |
| salvagecount | tinyint | Salvage Count |
| iscontainer | tinyint | Is Container: 0 = False, 1 = True |

