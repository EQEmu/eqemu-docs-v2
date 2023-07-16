# dynamic_zones

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZHluYW1pY196b25lcyB7XG4gICAgICAgIGludCBpbnN0YW5jZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnQgZHpfc3dpdGNoX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGNvbXBhc3Nfem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBzYWZlX3JldHVybl96b25lX2lkXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgfVxuICAgIGRvb3JzIHtcbiAgICAgICAgaW50dW5zaWduZWQgZGVzdF9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgICAgIGludCBkel9zd2l0Y2hfaWRcbiAgICAgICAgaW50IGtleWl0ZW1cbiAgICAgICAgdmFyY2hhciB6b25lXG4gICAgICAgIHZhcmNoYXIgZGVzdF96b25lXG4gICAgICAgIHNtYWxsaW50IHZlcnNpb25cbiAgICB9XG4gICAgZHluYW1pY196b25lX21lbWJlcnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBjaGFyYWN0ZXJfaWRcbiAgICAgICAgaW50dW5zaWduZWQgZHluYW1pY196b25lX2lkXG4gICAgfVxuICAgIGluc3RhbmNlX2xpc3Qge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgaW50dW5zaWduZWQgem9uZVxuICAgIH1cbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyB6b25lIDogXCJPbmUtdG8tT25lXCJcbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyBkb29ycyA6IFwiSGFzLU1hbnlcIlxuICAgIGR5bmFtaWNfem9uZXMgfHwtLW97IGR5bmFtaWNfem9uZV9tZW1iZXJzIDogXCJIYXMtTWFueVwiXG4gICAgZHluYW1pY196b25lcyB8fC0tb3sgaW5zdGFuY2VfbGlzdCA6IFwiT25lLXRvLU9uZVwiXG4gICAgZHluYW1pY196b25lcyB8fC0tb3sgem9uZSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZHluYW1pY196b25lcyB7XG4gICAgICAgIGludCBpbnN0YW5jZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnQgZHpfc3dpdGNoX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGNvbXBhc3Nfem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBzYWZlX3JldHVybl96b25lX2lkXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgfVxuICAgIGRvb3JzIHtcbiAgICAgICAgaW50dW5zaWduZWQgZGVzdF9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgICAgIGludCBkel9zd2l0Y2hfaWRcbiAgICAgICAgaW50IGtleWl0ZW1cbiAgICAgICAgdmFyY2hhciB6b25lXG4gICAgICAgIHZhcmNoYXIgZGVzdF96b25lXG4gICAgICAgIHNtYWxsaW50IHZlcnNpb25cbiAgICB9XG4gICAgZHluYW1pY196b25lX21lbWJlcnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBjaGFyYWN0ZXJfaWRcbiAgICAgICAgaW50dW5zaWduZWQgZHluYW1pY196b25lX2lkXG4gICAgfVxuICAgIGluc3RhbmNlX2xpc3Qge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgaW50dW5zaWduZWQgem9uZVxuICAgIH1cbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyB6b25lIDogXCJPbmUtdG8tT25lXCJcbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyBkb29ycyA6IFwiSGFzLU1hbnlcIlxuICAgIGR5bmFtaWNfem9uZXMgfHwtLW97IGR5bmFtaWNfem9uZV9tZW1iZXJzIDogXCJIYXMtTWFueVwiXG4gICAgZHluYW1pY196b25lcyB8fC0tb3sgaW5zdGFuY2VfbGlzdCA6IFwiT25lLXRvLU9uZVwiXG4gICAgZHluYW1pY196b25lcyB8fC0tb3sgem9uZSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZHluYW1pY196b25lcyB7XG4gICAgICAgIGludCBpbnN0YW5jZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnQgZHpfc3dpdGNoX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGNvbXBhc3Nfem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBzYWZlX3JldHVybl96b25lX2lkXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgfVxuICAgIGRvb3JzIHtcbiAgICAgICAgaW50dW5zaWduZWQgZGVzdF9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgICAgIGludCBkel9zd2l0Y2hfaWRcbiAgICAgICAgaW50IGtleWl0ZW1cbiAgICAgICAgdmFyY2hhciB6b25lXG4gICAgICAgIHZhcmNoYXIgZGVzdF96b25lXG4gICAgICAgIHNtYWxsaW50IHZlcnNpb25cbiAgICB9XG4gICAgZHluYW1pY196b25lX21lbWJlcnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBjaGFyYWN0ZXJfaWRcbiAgICAgICAgaW50dW5zaWduZWQgZHluYW1pY196b25lX2lkXG4gICAgfVxuICAgIGluc3RhbmNlX2xpc3Qge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgaW50dW5zaWduZWQgem9uZVxuICAgIH1cbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyB6b25lIDogXCJPbmUtdG8tT25lXCJcbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyBkb29ycyA6IFwiSGFzLU1hbnlcIlxuICAgIGR5bmFtaWNfem9uZXMgfHwtLW97IGR5bmFtaWNfem9uZV9tZW1iZXJzIDogXCJIYXMtTWFueVwiXG4gICAgZHluYW1pY196b25lcyB8fC0tb3sgaW5zdGFuY2VfbGlzdCA6IFwiT25lLXRvLU9uZVwiXG4gICAgZHluYW1pY196b25lcyB8fC0tb3sgem9uZSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | compass_zone_id | [zone](../../schema/zone/zone.md) | zoneidnumber |
| Has-Many | dz_switch_id | [doors](../../schema/doors/doors.md) | dz_switch_id |
| Has-Many | id | [dynamic_zone_members](../../schema/dynamic-zones/dynamic_zone_members.md) | dynamic_zone_id |
| One-to-One | instance_id | [instance_list](../../schema/instances/instance_list.md) | id |
| One-to-One | safe_return_zone_id | [zone](../../schema/zone/zone.md) | zoneidnumber |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Dynamic Zone Identifier |
| instance_id | int | [Instance Identifier](../../schema/instances/instance_list.md) |
| type | tinyint | Type |
| uuid | varchar | UUID |
| name | varchar | Name |
| leader_id | int | [Leader Character Identifier](../../schema/characters/character_data.md) |
| min_players | int | Minimum Players |
| max_players | int | Maximum Players |
| dz_switch_id | int | Dynamic Zone Switch Identifier](../../schema/doors/doors.md) |
| compass_zone_id | int | [Compass Zone Identifier](../../../../server/zones/zone-list) |
| compass_x | float | Compass X Coordinate |
| compass_y | float | Compass Y Coordinate |
| compass_z | float | Compass Z Coordinate |
| safe_return_zone_id | int | [Safe Return Zone Identifier](../../../../server/zones/zone-list) |
| safe_return_x | float | Safe Return X Coordinate |
| safe_return_y | float | Safe Return Y Coordinate |
| safe_return_z | float | Safe Return Z Coordinate |
| safe_return_heading | float | Safe Return Heading Coordinate |
| zone_in_x | float | Zone In X Coordinate |
| zone_in_y | float | Zone In Y Coordinate |
| zone_in_z | float | Zone In Z Coordinate |
| zone_in_heading | float | Zone In Heading Coordinate |
| has_zone_in | tinyint | Has Zone In: 0 = False, 1 = True |

