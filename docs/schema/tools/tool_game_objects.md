# tool_game_objects

!!! info
	This page was last generated 2023.01.30

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nYW1lX29iamVjdHMge1xuICAgICAgICBpbnQgem9uZWlkXG4gICAgICAgIHZhcmNoYXIgem9uZXNuXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCB2ZXJzaW9uXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgfVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogXCJPbmUtdG8tT25lXCJcbiAgICB0b29sX2dhbWVfb2JqZWN0cyB8fC0tb3sgem9uZSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nYW1lX29iamVjdHMge1xuICAgICAgICBpbnQgem9uZWlkXG4gICAgICAgIHZhcmNoYXIgem9uZXNuXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCB2ZXJzaW9uXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgfVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogXCJPbmUtdG8tT25lXCJcbiAgICB0b29sX2dhbWVfb2JqZWN0cyB8fC0tb3sgem9uZSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nYW1lX29iamVjdHMge1xuICAgICAgICBpbnQgem9uZWlkXG4gICAgICAgIHZhcmNoYXIgem9uZXNuXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCB2ZXJzaW9uXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgfVxuICAgIHRvb2xfZ2FtZV9vYmplY3RzIHx8LS1veyB6b25lIDogXCJPbmUtdG8tT25lXCJcbiAgICB0b29sX2dhbWVfb2JqZWN0cyB8fC0tb3sgem9uZSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


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

