# spawngroup

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25ncm91cCB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgc3Bhd24yIHtcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgIH1cbiAgICBzcGF3bmdyb3VwIHx8LS1veyBzcGF3bjIgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25ncm91cCB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgc3Bhd24yIHtcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgIH1cbiAgICBzcGF3bmdyb3VwIHx8LS1veyBzcGF3bjIgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | [spawn2](../../schema/spawns/spawn2.md) | spawngroupID |


## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Spawn Group Identifier |
| name | varchar | Name |
| spawn_limit | tinyint | Spawn Limit |
| dist | float | Distance |
| max_x | float | Max X Coordinate |
| min_x | float | Min X Coordinate |
| max_y | float | Max Y Coordinate |
| min_y | float | Min Y Coordinate |
| delay | int | Roaming Delay |
| mindelay | int | Minimum Delay |
| despawn | tinyint | [Despawn Type](../../../../server/npc/spawns/npc-despawn-types) |
| despawn_timer | int | Despawn Timer in Seconds |
| wp_spawns | tinyint |  |

