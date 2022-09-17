# dynamic_zones

!!! info
	This page was last generated 2022.09.16

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZHluYW1pY196b25lcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIGludHVuc2lnbmVkIGNvbXBhc3Nfem9uZV9pZFxuICAgICAgICBpbnQgZHpfc3dpdGNoX2lkXG4gICAgICAgIGludCBpbnN0YW5jZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBzYWZlX3JldHVybl96b25lX2lkXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgfVxuICAgIGRvb3JzIHtcbiAgICAgICAgaW50IGR6X3N3aXRjaF9pZFxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgICAgIGludCBrZXlpdGVtXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgICAgICB2YXJjaGFyIGRlc3Rfem9uZVxuICAgICAgICBzbWFsbGludCB2ZXJzaW9uXG4gICAgICAgIGludHVuc2lnbmVkIGRlc3RfaW5zdGFuY2VcbiAgICB9XG4gICAgZHluYW1pY196b25lX21lbWJlcnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBkeW5hbWljX3pvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgY2hhcmFjdGVyX2lkXG4gICAgfVxuICAgIGluc3RhbmNlX2xpc3Qge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZVxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgIH1cbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuICAgIGR5bmFtaWNfem9uZXMgfHwtLW97IGRvb3JzIDogSGFzLU1hbnlcbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyBkeW5hbWljX3pvbmVfbWVtYmVycyA6IEhhcy1NYW55XG4gICAgZHluYW1pY196b25lcyB8fC0tb3sgaW5zdGFuY2VfbGlzdCA6IE9uZS10by1PbmVcbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZHluYW1pY196b25lcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIGludHVuc2lnbmVkIGNvbXBhc3Nfem9uZV9pZFxuICAgICAgICBpbnQgZHpfc3dpdGNoX2lkXG4gICAgICAgIGludCBpbnN0YW5jZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBzYWZlX3JldHVybl96b25lX2lkXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgfVxuICAgIGRvb3JzIHtcbiAgICAgICAgaW50IGR6X3N3aXRjaF9pZFxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgICAgIGludCBrZXlpdGVtXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgICAgICB2YXJjaGFyIGRlc3Rfem9uZVxuICAgICAgICBzbWFsbGludCB2ZXJzaW9uXG4gICAgICAgIGludHVuc2lnbmVkIGRlc3RfaW5zdGFuY2VcbiAgICB9XG4gICAgZHluYW1pY196b25lX21lbWJlcnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBkeW5hbWljX3pvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgY2hhcmFjdGVyX2lkXG4gICAgfVxuICAgIGluc3RhbmNlX2xpc3Qge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZVxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgIH1cbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuICAgIGR5bmFtaWNfem9uZXMgfHwtLW97IGRvb3JzIDogSGFzLU1hbnlcbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyBkeW5hbWljX3pvbmVfbWVtYmVycyA6IEhhcy1NYW55XG4gICAgZHluYW1pY196b25lcyB8fC0tb3sgaW5zdGFuY2VfbGlzdCA6IE9uZS10by1PbmVcbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZHluYW1pY196b25lcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIGludHVuc2lnbmVkIGNvbXBhc3Nfem9uZV9pZFxuICAgICAgICBpbnQgZHpfc3dpdGNoX2lkXG4gICAgICAgIGludCBpbnN0YW5jZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBzYWZlX3JldHVybl96b25lX2lkXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgfVxuICAgIGRvb3JzIHtcbiAgICAgICAgaW50IGR6X3N3aXRjaF9pZFxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgICAgIGludCBrZXlpdGVtXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgICAgICB2YXJjaGFyIGRlc3Rfem9uZVxuICAgICAgICBzbWFsbGludCB2ZXJzaW9uXG4gICAgICAgIGludHVuc2lnbmVkIGRlc3RfaW5zdGFuY2VcbiAgICB9XG4gICAgZHluYW1pY196b25lX21lbWJlcnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBkeW5hbWljX3pvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgY2hhcmFjdGVyX2lkXG4gICAgfVxuICAgIGluc3RhbmNlX2xpc3Qge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgem9uZVxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgIH1cbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuICAgIGR5bmFtaWNfem9uZXMgfHwtLW97IGRvb3JzIDogSGFzLU1hbnlcbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyBkeW5hbWljX3pvbmVfbWVtYmVycyA6IEhhcy1NYW55XG4gICAgZHluYW1pY196b25lcyB8fC0tb3sgaW5zdGFuY2VfbGlzdCA6IE9uZS10by1PbmVcbiAgICBkeW5hbWljX3pvbmVzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


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

