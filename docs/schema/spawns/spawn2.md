# spawn2

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd24yIHtcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgIH1cbiAgICBzcGF3bmdyb3VwIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBzcGF3bmVudHJ5IHtcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgICAgICBpbnQgbnBjSURcbiAgICB9XG4gICAgc3Bhd24yIHx8LS1veyBzcGF3bmdyb3VwIDogT25lLXRvLU9uZVxuICAgIHNwYXduMiB8fC0tb3sgc3Bhd25lbnRyeSA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd24yIHtcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgIH1cbiAgICBzcGF3bmdyb3VwIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBzcGF3bmVudHJ5IHtcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgICAgICBpbnQgbnBjSURcbiAgICB9XG4gICAgc3Bhd24yIHx8LS1veyBzcGF3bmdyb3VwIDogT25lLXRvLU9uZVxuICAgIHNwYXduMiB8fC0tb3sgc3Bhd25lbnRyeSA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | spawngroupID | [spawngroup](../../schema/spawns/spawngroup.md) | id |
| Has-Many | spawngroupID | [spawnentry](../../schema/spawns/spawnentry.md) | spawngroupID |


## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Spawn2 Entry Identifier |
| spawngroupID | int | [Unique Spawngroup Identifier](spawngroup.md) |
| zone | varchar | [Zone Short Name](../../../../server/zones/zone-list) |
| version | smallint | Version |
| x | float | X Coordinate |
| y | float | Y Coordinate |
| z | float | Z Coordinate |
| heading | float | Heading Coordinate |
| respawntime | int | Respawn Time in Seconds |
| variance | int | Variance in Seconds |
| pathgrid | int | [Path Grid Identifier](../../schema/grids/grid.md) |
| path_when_zone_idle | tinyint |  |
| _condition | mediumint | Condition |
| cond_value | mediumint | Condition Value |
| enabled | tinyint | Enabled: 0 = False, 1 = True |
| animation | tinyint | [Animation](../../../../server/npc/npc-animation-types) |
| min_expansion | tinyint |  |
| max_expansion | tinyint |  |
| content_flags | varchar |  |
| content_flags_disabled | varchar |  |

