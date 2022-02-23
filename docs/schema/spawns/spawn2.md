# spawn2

!!! info
	This page was last generated 2022.02.23

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd24yIHtcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgIH1cbiAgICBzcGF3bmdyb3VwIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIHNwYXduZW50cnkge1xuICAgICAgICBpbnQgbnBjSURcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgIH1cbiAgICBzcGF3bjIgfHwtLW97IHNwYXduZ3JvdXAgOiBPbmUtdG8tT25lXG4gICAgc3Bhd24yIHx8LS1veyBzcGF3bmVudHJ5IDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd24yIHtcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgIH1cbiAgICBzcGF3bmdyb3VwIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIHNwYXduZW50cnkge1xuICAgICAgICBpbnQgbnBjSURcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgIH1cbiAgICBzcGF3bjIgfHwtLW97IHNwYXduZ3JvdXAgOiBPbmUtdG8tT25lXG4gICAgc3Bhd24yIHx8LS1veyBzcGF3bmVudHJ5IDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd24yIHtcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgIH1cbiAgICBzcGF3bmdyb3VwIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIHNwYXduZW50cnkge1xuICAgICAgICBpbnQgbnBjSURcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgIH1cbiAgICBzcGF3bjIgfHwtLW97IHNwYXduZ3JvdXAgOiBPbmUtdG8tT25lXG4gICAgc3Bhd24yIHx8LS1veyBzcGF3bmVudHJ5IDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

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

