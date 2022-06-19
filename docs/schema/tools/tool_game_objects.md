# tool_game_objects

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nYW1lX29iamVjdHMge1xuICAgICAgICB2YXJjaGFyIHpvbmVpZFxuICAgICAgICB2YXJjaGFyIHpvbmVzblxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgaW50IHpvbmVpZG51bWJlclxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIHNob3J0X25hbWVcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICB2YXJjaGFyIHpvbmVfc2hvcnRfbmFtZVxuICAgICAgICB2YXJjaGFyIHpvbmVpZG51bWVyXG4gICAgfVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nYW1lX29iamVjdHMge1xuICAgICAgICB2YXJjaGFyIHpvbmVpZFxuICAgICAgICB2YXJjaGFyIHpvbmVzblxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgaW50IHpvbmVpZG51bWJlclxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIHNob3J0X25hbWVcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICB2YXJjaGFyIHpvbmVfc2hvcnRfbmFtZVxuICAgICAgICB2YXJjaGFyIHpvbmVpZG51bWVyXG4gICAgfVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nYW1lX29iamVjdHMge1xuICAgICAgICB2YXJjaGFyIHpvbmVpZFxuICAgICAgICB2YXJjaGFyIHpvbmVzblxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgaW50IHpvbmVpZG51bWJlclxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIHNob3J0X25hbWVcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICB2YXJjaGFyIHpvbmVfc2hvcnRfbmFtZVxuICAgICAgICB2YXJjaGFyIHpvbmVpZG51bWVyXG4gICAgfVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


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

