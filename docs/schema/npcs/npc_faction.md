# npc_faction

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX2ZhY3Rpb24ge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHByaW1hcnlmYWN0aW9uXG4gICAgfVxuICAgIG5wY19mYWN0aW9uX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBucGNfZmFjdGlvbl9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBmYWN0aW9uX2lkXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBucGNfZmFjdGlvbiB8fC0tb3sgbnBjX2ZhY3Rpb25fZW50cmllcyA6IFwiSGFzLU1hbnlcIlxuICAgIG5wY19mYWN0aW9uIHx8LS1veyBmYWN0aW9uX2xpc3QgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX2ZhY3Rpb24ge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHByaW1hcnlmYWN0aW9uXG4gICAgfVxuICAgIG5wY19mYWN0aW9uX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBucGNfZmFjdGlvbl9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBmYWN0aW9uX2lkXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBucGNfZmFjdGlvbiB8fC0tb3sgbnBjX2ZhY3Rpb25fZW50cmllcyA6IFwiSGFzLU1hbnlcIlxuICAgIG5wY19mYWN0aW9uIHx8LS1veyBmYWN0aW9uX2xpc3QgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX2ZhY3Rpb24ge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHByaW1hcnlmYWN0aW9uXG4gICAgfVxuICAgIG5wY19mYWN0aW9uX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBucGNfZmFjdGlvbl9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBmYWN0aW9uX2lkXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBucGNfZmFjdGlvbiB8fC0tb3sgbnBjX2ZhY3Rpb25fZW50cmllcyA6IFwiSGFzLU1hbnlcIlxuICAgIG5wY19mYWN0aW9uIHx8LS1veyBmYWN0aW9uX2xpc3QgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [npc_faction_entries](../../schema/npcs/npc_faction_entries.md) | npc_faction_id |
| One-to-One | primaryfaction | [faction_list](../../schema/factions/faction_list.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique NPC Faction Identifier |
| name | tinytext | Name |
| primaryfaction | int | [Primary Faction Identifier](../../schema/factions/faction_list.md) |
| ignore_primary_assist | tinyint | Ignore Primary Assist: 0 = False, &gt;0 = True |

