# grid_entries

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3JpZF9lbnRyaWVzIHtcbiAgICAgICAgaW50IGdyaWRpZFxuICAgICAgICB2YXJjaGFyIGdyaWRfaWRcbiAgICAgICAgaW50IHpvbmVpZFxuICAgIH1cbiAgICBncmlkIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCB6b25laWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCB2ZXJzaW9uXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICB9XG4gICAgZ3JpZF9lbnRyaWVzIHx8LS1veyBncmlkIDogXCJPbmUtdG8tT25lXCJcbiAgICBncmlkX2VudHJpZXMgfHwtLW97IHpvbmUgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3JpZF9lbnRyaWVzIHtcbiAgICAgICAgaW50IGdyaWRpZFxuICAgICAgICB2YXJjaGFyIGdyaWRfaWRcbiAgICAgICAgaW50IHpvbmVpZFxuICAgIH1cbiAgICBncmlkIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCB6b25laWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCB2ZXJzaW9uXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICB9XG4gICAgZ3JpZF9lbnRyaWVzIHx8LS1veyBncmlkIDogXCJPbmUtdG8tT25lXCJcbiAgICBncmlkX2VudHJpZXMgfHwtLW97IHpvbmUgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3JpZF9lbnRyaWVzIHtcbiAgICAgICAgaW50IGdyaWRpZFxuICAgICAgICB2YXJjaGFyIGdyaWRfaWRcbiAgICAgICAgaW50IHpvbmVpZFxuICAgIH1cbiAgICBncmlkIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCB6b25laWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCB2ZXJzaW9uXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICB9XG4gICAgZ3JpZF9lbnRyaWVzIHx8LS1veyBncmlkIDogXCJPbmUtdG8tT25lXCJcbiAgICBncmlkX2VudHJpZXMgfHwtLW97IHpvbmUgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


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

