# npc_faction

## Relationship Diagram
[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX2ZhY3Rpb24ge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIG5wY19mYWN0aW9uX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBucGNfZmFjdGlvbl9pZFxuICAgIH1cbiAgICBucGNfZmFjdGlvbiB8fC0tb3sgbnBjX2ZhY3Rpb25fZW50cmllcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX2ZhY3Rpb24ge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIG5wY19mYWN0aW9uX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBucGNfZmFjdGlvbl9pZFxuICAgIH1cbiAgICBucGNfZmFjdGlvbiB8fC0tb3sgbnBjX2ZhY3Rpb25fZW50cmllcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | npc_faction_entries | npc_faction_id |


## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique NPC Faction Identifier |
| name | tinytext | Name |
| primaryfaction | int | [Primary Faction Identifier](../../schema/factions/faction_list.md) |
| ignore_primary_assist | tinyint | Ignore Primary Assist: 0 = False, &gt;0 = True |

