# grid

!!! info
	This page was last generated 2022.06.17

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3JpZCB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgem9uZWlkXG4gICAgfVxuICAgIGdyaWRfZW50cmllcyB7XG4gICAgICAgIGludCBncmlkaWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgICB6b25laWR1bnVtYmVyXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgfVxuICAgIGdyaWQgfHwtLW97IGdyaWRfZW50cmllcyA6IEhhcy1NYW55XG4gICAgZ3JpZCB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3JpZCB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgem9uZWlkXG4gICAgfVxuICAgIGdyaWRfZW50cmllcyB7XG4gICAgICAgIGludCBncmlkaWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgICB6b25laWR1bnVtYmVyXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgfVxuICAgIGdyaWQgfHwtLW97IGdyaWRfZW50cmllcyA6IEhhcy1NYW55XG4gICAgZ3JpZCB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3JpZCB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgem9uZWlkXG4gICAgfVxuICAgIGdyaWRfZW50cmllcyB7XG4gICAgICAgIGludCBncmlkaWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgICB6b25laWR1bnVtYmVyXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgfVxuICAgIGdyaWQgfHwtLW97IGdyaWRfZW50cmllcyA6IEhhcy1NYW55XG4gICAgZ3JpZCB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [grid_entries](../../schema/grids/grid_entries.md) | gridid |
| One-to-One | zoneid | [zone](../../schema/zone/zone.md) | zoneidnumber |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Grid Identifier |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| type | int | [Wander Type](../../../../server/npc/spawns/wander-types) |
| type2 | int | [Pause Type](../../../../server/npc/spawns/pause-types) |

