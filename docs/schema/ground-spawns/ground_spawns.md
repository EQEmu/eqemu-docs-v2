# ground_spawns

!!! info
	This page was last generated 2022.06.17

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3JvdW5kX3NwYXducyB7XG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVpZFxuICAgICAgICBpbnR1bnNpZ25lZCBpdGVtXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICAgem9uZWlkdW51bWJlclxuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgIH1cbiAgICBpdGVtcyB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBncm91bmRfc3Bhd25zIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuICAgIGdyb3VuZF9zcGF3bnMgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3JvdW5kX3NwYXducyB7XG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVpZFxuICAgICAgICBpbnR1bnNpZ25lZCBpdGVtXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICAgem9uZWlkdW51bWJlclxuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgIH1cbiAgICBpdGVtcyB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBncm91bmRfc3Bhd25zIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuICAgIGdyb3VuZF9zcGF3bnMgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3JvdW5kX3NwYXducyB7XG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVpZFxuICAgICAgICBpbnR1bnNpZ25lZCBpdGVtXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICAgem9uZWlkdW51bWJlclxuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgIH1cbiAgICBpdGVtcyB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBncm91bmRfc3Bhd25zIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuICAgIGdyb3VuZF9zcGF3bnMgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | zoneid | [zone](../../schema/zone/zone.md) | zoneidnumber |
| One-to-One | item | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Ground Spawn Identifier |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| version | smallint | Version: -1 For All |
| max_x | float | Maximum X Coordinate |
| max_y | float | Maximum Y Coordinate |
| max_z | float | Maximum Z Coordinate |
| min_x | float | Minimum X Coordinate |
| min_y | float | Minimum Y Coordinate |
| heading | float | Heading Coordinate |
| name | varchar | Name |
| item | int | [Item Identifier](items.md) |
| max_allowed | int | Max Allowed |
| comment | varchar | Comment |
| respawn_timer | int | Respawn Timer in Seconds |
| min_expansion | tinyint | [Minimum Expansion](../../../../server/operation/expansion-list) |
| max_expansion | tinyint | [Maximum Expansion](../../../../server/operation/expansion-list) |
| content_flags | varchar | Content Flags Required to be Enabled |
| content_flags_disabled | varchar | Content Flags Required to be Disabled |

