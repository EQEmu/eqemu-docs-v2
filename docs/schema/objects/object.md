# object

!!! info
	This page was last generated 2022.05.11

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgb2JqZWN0IHtcbiAgICAgICAgaW50IGl0ZW1pZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25laWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgIHpvbmVpZHVudW1iZXJcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIG9iamVjdCB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcbiAgICBvYmplY3QgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgb2JqZWN0IHtcbiAgICAgICAgaW50IGl0ZW1pZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25laWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgIHpvbmVpZHVudW1iZXJcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIG9iamVjdCB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcbiAgICBvYmplY3QgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgb2JqZWN0IHtcbiAgICAgICAgaW50IGl0ZW1pZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25laWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgIHpvbmVpZHVudW1iZXJcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIG9iamVjdCB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcbiAgICBvYmplY3QgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | zoneid | [zone](../../schema/zone/zone.md) | zoneidnumber |
| One-to-One | itemid | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Object Identifier |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| version | smallint | Version: -1 For All |
| xpos | float | X Coordinate |
| ypos | float | Y Coordinate |
| zpos | float | Z Coordinate |
| heading | float | Heading Coordinate |
| itemid | int | [Item Identifier](../../schema/items/items.md) |
| charges | smallint | Charges |
| objectname | varchar | Object Name |
| type | int | [Type](../../../../server/zones/object-types) |
| icon | int | Icon |
| unknown08 | mediumint | Unknown |
| unknown10 | mediumint | Unknown |
| unknown20 | int | Unknown |
| unknown24 | int | Unknown |
| unknown60 | int | Unknown |
| unknown64 | int | Unknown |
| unknown68 | int | Unknown |
| unknown72 | int | Unknown |
| unknown76 | int | Unknown |
| unknown84 | int | Unknown |
| size | float | Size |
| tilt_x | float | Tilt X |
| tilt_y | float | Tilt Y |
| display_name | varchar | Display Name |
| min_expansion | tinyint |  |
| max_expansion | tinyint |  |
| content_flags | varchar |  |
| content_flags_disabled | varchar |  |

