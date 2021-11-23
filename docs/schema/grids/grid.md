# grid

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3JpZCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBncmlkX2VudHJpZXMge1xuICAgICAgICBpbnQgZ3JpZGlkXG4gICAgfVxuICAgIGdyaWQgfHwtLW97IGdyaWRfZW50cmllcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3JpZCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBncmlkX2VudHJpZXMge1xuICAgICAgICBpbnQgZ3JpZGlkXG4gICAgfVxuICAgIGdyaWQgfHwtLW97IGdyaWRfZW50cmllcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [grid_entries](../../schema/grids/grid_entries.md) | gridid |


## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Grid Identifier |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| type | int | [Wander Type](../../../../server/npc/spawns/wander-types) |
| type2 | int | [Pause Type](../../../../server/npc/spawns/pause-types) |

