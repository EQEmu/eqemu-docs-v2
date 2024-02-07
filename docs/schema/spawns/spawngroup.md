# spawngroup

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25ncm91cCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBzcGF3bjIge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHBhdGhncmlkXG4gICAgICAgIGludCBzcGF3bmdyb3VwSURcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICBzbWFsbGludCB2ZXJzaW9uXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICBzcGF3bmdyb3VwIHx8LS1veyBzcGF3bjIgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25ncm91cCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBzcGF3bjIge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHBhdGhncmlkXG4gICAgICAgIGludCBzcGF3bmdyb3VwSURcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICBzbWFsbGludCB2ZXJzaW9uXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICBzcGF3bmdyb3VwIHx8LS1veyBzcGF3bjIgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25ncm91cCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBzcGF3bjIge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHBhdGhncmlkXG4gICAgICAgIGludCBzcGF3bmdyb3VwSURcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICBzbWFsbGludCB2ZXJzaW9uXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICBzcGF3bmdyb3VwIHx8LS1veyBzcGF3bjIgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


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
| wp_spawns | tinyint | Waypoint Spawns: 0 = False, 1 = True |

