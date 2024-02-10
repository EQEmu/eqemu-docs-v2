# grid_entries

## Relationships

```mermaid
erDiagram
    grid_entries {
        int gridid
        varchar grid_id
        int zoneid
    }
    grid {
        int id
        int zoneid
    }
    zone {
        int zoneidnumber
        varchar short_name
        tinyintunsigned version
        varchar content_flags
        varchar content_flags_disabled
    }
    grid_entries ||--o{ grid : "One-to-One"
    grid_entries ||--o{ zone : "One-to-One"


```


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

