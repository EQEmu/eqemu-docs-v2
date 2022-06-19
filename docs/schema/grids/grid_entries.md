# grid_entries

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3JpZF9lbnRyaWVzIHtcbiAgICAgICAgaW50IHpvbmVpZFxuICAgICAgICBpbnQgZ3JpZGlkXG4gICAgICAgIHZhcmNoYXIgZ3JpZF9pZFxuICAgIH1cbiAgICBncmlkIHtcbiAgICAgICAgaW50IHpvbmVpZFxuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICB9XG4gICAgZ3JpZF9lbnRyaWVzIHx8LS1veyBncmlkIDogT25lLXRvLU9uZVxuICAgIGdyaWRfZW50cmllcyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3JpZF9lbnRyaWVzIHtcbiAgICAgICAgaW50IHpvbmVpZFxuICAgICAgICBpbnQgZ3JpZGlkXG4gICAgICAgIHZhcmNoYXIgZ3JpZF9pZFxuICAgIH1cbiAgICBncmlkIHtcbiAgICAgICAgaW50IHpvbmVpZFxuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICB9XG4gICAgZ3JpZF9lbnRyaWVzIHx8LS1veyBncmlkIDogT25lLXRvLU9uZVxuICAgIGdyaWRfZW50cmllcyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3JpZF9lbnRyaWVzIHtcbiAgICAgICAgaW50IHpvbmVpZFxuICAgICAgICBpbnQgZ3JpZGlkXG4gICAgICAgIHZhcmNoYXIgZ3JpZF9pZFxuICAgIH1cbiAgICBncmlkIHtcbiAgICAgICAgaW50IHpvbmVpZFxuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICB9XG4gICAgZ3JpZF9lbnRyaWVzIHx8LS1veyBncmlkIDogT25lLXRvLU9uZVxuICAgIGdyaWRfZW50cmllcyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | grid_id | [grid](../../schema/grids/grid.md) | id |
| One-to-One | zoneid | [zone](../../schema/zone/zone.md) | zoneidnumber |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| gridid | int | [Grid Identifier](grid.md) |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| number | int | Waypoint Identifier |
| x | float | X Coordinate |
| y | float | Y Coordinate |
| z | float | Z Coordinate |
| heading | float | Heading Coordinate |
| pause | int | Pause in Seconds |
| centerpoint | tinyint | Center Point: 0 = False, 1 = True |

