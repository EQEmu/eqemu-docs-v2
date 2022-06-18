# spawngroup

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25ncm91cCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBzcGF3bjIge1xuICAgICAgICB2YXJjaGFyIHpvbmVcbiAgICAgICAgc21hbGxpbnQgdmVyc2lvblxuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgICAgICBtZWRpdW1pbnR1bnNpZ25lZCBfY29uZGl0aW9uXG4gICAgICAgIG1lZGl1bWludCBjb25kX3ZhbHVlXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgaW50IHBhdGhncmlkXG4gICAgfVxuICAgIHNwYXduZ3JvdXAgfHwtLW97IHNwYXduMiA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25ncm91cCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBzcGF3bjIge1xuICAgICAgICB2YXJjaGFyIHpvbmVcbiAgICAgICAgc21hbGxpbnQgdmVyc2lvblxuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgICAgICBtZWRpdW1pbnR1bnNpZ25lZCBfY29uZGl0aW9uXG4gICAgICAgIG1lZGl1bWludCBjb25kX3ZhbHVlXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgaW50IHBhdGhncmlkXG4gICAgfVxuICAgIHNwYXduZ3JvdXAgfHwtLW97IHNwYXduMiA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25ncm91cCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBzcGF3bjIge1xuICAgICAgICB2YXJjaGFyIHpvbmVcbiAgICAgICAgc21hbGxpbnQgdmVyc2lvblxuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHNwYXduZ3JvdXBJRFxuICAgICAgICBtZWRpdW1pbnR1bnNpZ25lZCBfY29uZGl0aW9uXG4gICAgICAgIG1lZGl1bWludCBjb25kX3ZhbHVlXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgaW50IHBhdGhncmlkXG4gICAgfVxuICAgIHNwYXduZ3JvdXAgfHwtLW97IHNwYXduMiA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

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

