# spawngroup

## Relationship Diagram
[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25ncm91cCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBzcGF3bjIge1xuICAgICAgICBpbnQgc3Bhd25ncm91cElEXG4gICAgICAgIGludCBzcGF3bmdyb3VwSURcbiAgICB9XG4gICAgc3Bhd25ncm91cCB8fC0tb3sgc3Bhd24yIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25ncm91cCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBzcGF3bjIge1xuICAgICAgICBpbnQgc3Bhd25ncm91cElEXG4gICAgICAgIGludCBzcGF3bmdyb3VwSURcbiAgICB9XG4gICAgc3Bhd25ncm91cCB8fC0tb3sgc3Bhd24yIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | spawn2 | spawngroupID |


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

