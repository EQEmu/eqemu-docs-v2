# dynamic_zone_templates

!!! info
	This page was last generated 2023.01.22

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZHluYW1pY196b25lX3RlbXBsYXRlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIHRhc2tzIHtcbiAgICAgICAgaW50dW5zaWduZWQgZHpfdGVtcGxhdGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdGlueWludCB0eXBlXG4gICAgfVxuICAgIGR5bmFtaWNfem9uZV90ZW1wbGF0ZXMgfHwtLW97IHRhc2tzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZHluYW1pY196b25lX3RlbXBsYXRlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIHRhc2tzIHtcbiAgICAgICAgaW50dW5zaWduZWQgZHpfdGVtcGxhdGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdGlueWludCB0eXBlXG4gICAgfVxuICAgIGR5bmFtaWNfem9uZV90ZW1wbGF0ZXMgfHwtLW97IHRhc2tzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZHluYW1pY196b25lX3RlbXBsYXRlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIHRhc2tzIHtcbiAgICAgICAgaW50dW5zaWduZWQgZHpfdGVtcGxhdGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdGlueWludCB0eXBlXG4gICAgfVxuICAgIGR5bmFtaWNfem9uZV90ZW1wbGF0ZXMgfHwtLW97IHRhc2tzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [tasks](../../schema/tasks/tasks.md) | dz_template_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Dynamic Zone Template Unique Identifier |
| zone_id | int | [Dynamic Zone Template Zone Identifier](../../../../server/zones/zone-list) |
| zone_version | int | Zone Version |
| name | varchar | Dynamic Zone Template Name |
| min_players | int | Minimum Players |
| max_players | int | Maximum Players |
| duration_seconds | int | Duration in Seconds |
| dz_switch_id | int | [Dynamic Zone Switch Identifier](../../schema/doors/doors.md) |
| compass_zone_id | int | [Compass Zone Identifier](../../../../server/zones/zone-list) |
| compass_x | float | Compass X Coordinate |
| compass_y | float | Compass Y Coordinate |
| compass_z | float | Compass Z Coordinatet |
| return_zone_id | int | [Return Zone Identifier](../../../../server/zones/zone-list) |
| return_x | float | Return X Coordinate |
| return_y | float | Return Y Coordinate |
| return_z | float | Return Z Coordinate |
| return_h | float | Return Heading Coordinate |
| override_zone_in | tinyint | Override Zone In: 0 = False, 1 = True |
| zone_in_x | float | Zone In X Coordinate |
| zone_in_y | float | Zone In Y Coordinate |
| zone_in_z | float | Zone In Z Coordinate |
| zone_in_h | float | Zone In Heading Coordinate |

