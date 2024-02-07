# tool_game_objects

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nYW1lX29iamVjdHMge1xuICAgICAgICBpbnQgem9uZWlkXG4gICAgICAgIHZhcmNoYXIgem9uZXNuXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICB2YXJjaGFyIHNob3J0X25hbWVcbiAgICAgICAgaW50IHpvbmVpZG51bWJlclxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgfVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogXCJPbmUtdG8tT25lXCJcbiAgICB0b29sX2dhbWVfb2JqZWN0cyB8fC0tb3sgem9uZSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nYW1lX29iamVjdHMge1xuICAgICAgICBpbnQgem9uZWlkXG4gICAgICAgIHZhcmNoYXIgem9uZXNuXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICB2YXJjaGFyIHNob3J0X25hbWVcbiAgICAgICAgaW50IHpvbmVpZG51bWJlclxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgfVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogXCJPbmUtdG8tT25lXCJcbiAgICB0b29sX2dhbWVfb2JqZWN0cyB8fC0tb3sgem9uZSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nYW1lX29iamVjdHMge1xuICAgICAgICBpbnQgem9uZWlkXG4gICAgICAgIHZhcmNoYXIgem9uZXNuXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICB2YXJjaGFyIHNob3J0X25hbWVcbiAgICAgICAgaW50IHpvbmVpZG51bWJlclxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgfVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogXCJPbmUtdG8tT25lXCJcbiAgICB0b29sX2dhbWVfb2JqZWN0cyB8fC0tb3sgem9uZSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


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

