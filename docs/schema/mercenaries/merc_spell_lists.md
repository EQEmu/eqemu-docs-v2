# merc_spell_lists

!!! info
	This page was last generated 2022.06.17

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19zcGVsbF9saXN0cyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfc3BlbGxfbGlzdF9pZFxuICAgIH1cbiAgICBtZXJjX3NwZWxsX2xpc3RfZW50cmllcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfc3BlbGxfbGlzdF9pZFxuICAgIH1cbiAgICBtZXJjX3NwZWxsX2xpc3RzIHx8LS1veyBtZXJjX3NwZWxsX2xpc3RfZW50cmllcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19zcGVsbF9saXN0cyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfc3BlbGxfbGlzdF9pZFxuICAgIH1cbiAgICBtZXJjX3NwZWxsX2xpc3RfZW50cmllcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfc3BlbGxfbGlzdF9pZFxuICAgIH1cbiAgICBtZXJjX3NwZWxsX2xpc3RzIHx8LS1veyBtZXJjX3NwZWxsX2xpc3RfZW50cmllcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19zcGVsbF9saXN0cyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfc3BlbGxfbGlzdF9pZFxuICAgIH1cbiAgICBtZXJjX3NwZWxsX2xpc3RfZW50cmllcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfc3BlbGxfbGlzdF9pZFxuICAgIH1cbiAgICBtZXJjX3NwZWxsX2xpc3RzIHx8LS1veyBtZXJjX3NwZWxsX2xpc3RfZW50cmllcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | merc_spell_list_id | [merc_spell_list_entries](../../schema/mercenaries/merc_spell_list_entries.md) | merc_spell_list_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merc_spell_list_id | int | Unique Mercenary Spell List Identifier |
| class_id | int | [Class Identifier](../../../../categories/player/class-list) |
| proficiency_id | tinyint | Proficiency Identifier |
| name | varchar | Name |

