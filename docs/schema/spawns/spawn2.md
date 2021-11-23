# spawn2

## Relationship Diagram
[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd24yIHtcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgICAgICBpbnQgc3Bhd25ncm91cElEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBzcGF3bmdyb3VwSURcbiAgICAgICAgaW50IG5wY0lEXG4gICAgfVxuICAgIHNwYXduMiB8fC0tb3sgc3Bhd25ncm91cCA6IE9uZS10by1PbmVcbiAgICBzcGF3bjIgfHwtLW97IHNwYXduZW50cnkgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd24yIHtcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgICAgICBpbnQgc3Bhd25ncm91cElEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBzcGF3bmdyb3VwSURcbiAgICAgICAgaW50IG5wY0lEXG4gICAgfVxuICAgIHNwYXduMiB8fC0tb3sgc3Bhd25ncm91cCA6IE9uZS10by1PbmVcbiAgICBzcGF3bjIgfHwtLW97IHNwYXduZW50cnkgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | spawngroupID | spawngroup | id |
| Has-Many | spawngroupID | spawnentry | spawngroupID |


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

