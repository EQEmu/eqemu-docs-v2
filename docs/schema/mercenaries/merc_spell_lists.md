# merc_spell_lists

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19zcGVsbF9saXN0cyB7XG4gICAgICAgIHZhcmNoYXIgbWVyY19zcGVsbF9saXN0X2lkXG4gICAgfVxuICAgIG1lcmNfc3BlbGxfbGlzdF9lbnRyaWVzIHtcbiAgICAgICAgdmFyY2hhciBzcGVsbF9pZFxuICAgICAgICB2YXJjaGFyIHN0YW5jZV9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfc3BlbGxfbGlzdF9pZFxuICAgIH1cbiAgICBtZXJjX3NwZWxsX2xpc3RzIHx8LS1veyBtZXJjX3NwZWxsX2xpc3RfZW50cmllcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19zcGVsbF9saXN0cyB7XG4gICAgICAgIHZhcmNoYXIgbWVyY19zcGVsbF9saXN0X2lkXG4gICAgfVxuICAgIG1lcmNfc3BlbGxfbGlzdF9lbnRyaWVzIHtcbiAgICAgICAgdmFyY2hhciBzcGVsbF9pZFxuICAgICAgICB2YXJjaGFyIHN0YW5jZV9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfc3BlbGxfbGlzdF9pZFxuICAgIH1cbiAgICBtZXJjX3NwZWxsX2xpc3RzIHx8LS1veyBtZXJjX3NwZWxsX2xpc3RfZW50cmllcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19zcGVsbF9saXN0cyB7XG4gICAgICAgIHZhcmNoYXIgbWVyY19zcGVsbF9saXN0X2lkXG4gICAgfVxuICAgIG1lcmNfc3BlbGxfbGlzdF9lbnRyaWVzIHtcbiAgICAgICAgdmFyY2hhciBzcGVsbF9pZFxuICAgICAgICB2YXJjaGFyIHN0YW5jZV9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfc3BlbGxfbGlzdF9pZFxuICAgIH1cbiAgICBtZXJjX3NwZWxsX2xpc3RzIHx8LS1veyBtZXJjX3NwZWxsX2xpc3RfZW50cmllcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | merc_spell_list_id | [merc_spell_list_entries](../../schema/mercenaries/merc_spell_list_entries.md) | merc_spell_list_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merc_spell_list_id | int | Unique Mercenary Spell List Identifier |
| class_id | int | [Class Identifier](../../../../server/player/class-list) |
| proficiency_id | tinyint | Proficiency Identifier |
| name | varchar | Name |

