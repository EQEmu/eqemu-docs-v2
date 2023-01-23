# tool_game_objects

!!! info
	This page was last generated 2023.01.22

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nYW1lX29iamVjdHMge1xuICAgICAgICBpbnQgem9uZWlkXG4gICAgICAgIHZhcmNoYXIgem9uZXNuXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCB2ZXJzaW9uXG4gICAgfVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nYW1lX29iamVjdHMge1xuICAgICAgICBpbnQgem9uZWlkXG4gICAgICAgIHZhcmNoYXIgem9uZXNuXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCB2ZXJzaW9uXG4gICAgfVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nYW1lX29iamVjdHMge1xuICAgICAgICBpbnQgem9uZWlkXG4gICAgICAgIHZhcmNoYXIgem9uZXNuXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCB2ZXJzaW9uXG4gICAgfVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | zoneid | [zone](../../schema/zone/zone.md) | zoneidnumber |
| One-to-One | zonesn | [zone](../../schema/zone/zone.md) | short_name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Tool Game Object Identifier |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| zonesn | varchar | [Zone Short Name](../../../../server/zones/zone-list) |
| object_name | varchar | Object Name |
| file_from | varchar | File From |
| is_global | tinyint | Is Global: 0 = False, 1 = True |

