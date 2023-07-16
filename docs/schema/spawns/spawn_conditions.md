# spawn_conditions

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25fY29uZGl0aW9ucyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICBzcGF3bl9jb25kaXRpb25fdmFsdWVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaW5zdGFuY2VfaWRcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciB6b25lXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgfVxuICAgIHNwYXduX2NvbmRpdGlvbnMgfHwtLW97IHNwYXduX2NvbmRpdGlvbl92YWx1ZXMgOiBcIk9uZS10by1PbmVcIlxuICAgIHNwYXduX2NvbmRpdGlvbnMgfHwtLW97IHpvbmUgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25fY29uZGl0aW9ucyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICBzcGF3bl9jb25kaXRpb25fdmFsdWVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaW5zdGFuY2VfaWRcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciB6b25lXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgfVxuICAgIHNwYXduX2NvbmRpdGlvbnMgfHwtLW97IHNwYXduX2NvbmRpdGlvbl92YWx1ZXMgOiBcIk9uZS10by1PbmVcIlxuICAgIHNwYXduX2NvbmRpdGlvbnMgfHwtLW97IHpvbmUgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25fY29uZGl0aW9ucyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICBzcGF3bl9jb25kaXRpb25fdmFsdWVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaW5zdGFuY2VfaWRcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciB6b25lXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgfVxuICAgIHNwYXduX2NvbmRpdGlvbnMgfHwtLW97IHNwYXduX2NvbmRpdGlvbl92YWx1ZXMgOiBcIk9uZS10by1PbmVcIlxuICAgIHNwYXduX2NvbmRpdGlvbnMgfHwtLW97IHpvbmUgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | [spawn_condition_values](../../schema/spawns/spawn_condition_values.md) | id |
| One-to-One | zone | [zone](../../schema/zone/zone.md) | short_name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| zone | varchar | [Zone Short Name](../../../../server/zones/zone-list) |
| id | mediumint | Spawn Condition Identifier |
| value | mediumint | Value |
| onchange | tinyint | [On Change Type](../../../../server/npc/spawns/on-change-types) |
| name | varchar | Name |

