# npc_spells_effects_entries

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX3NwZWxsc19lZmZlY3RzX2VudHJpZXMge1xuICAgICAgICBpbnQgbnBjX3NwZWxsc19lZmZlY3RzX2lkXG4gICAgfVxuICAgIG5wY19zcGVsbHNfZWZmZWN0cyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIG5wY19zcGVsbHNfZWZmZWN0c19lbnRyaWVzIHx8LS1veyBucGNfc3BlbGxzX2VmZmVjdHMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX3NwZWxsc19lZmZlY3RzX2VudHJpZXMge1xuICAgICAgICBpbnQgbnBjX3NwZWxsc19lZmZlY3RzX2lkXG4gICAgfVxuICAgIG5wY19zcGVsbHNfZWZmZWN0cyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIG5wY19zcGVsbHNfZWZmZWN0c19lbnRyaWVzIHx8LS1veyBucGNfc3BlbGxzX2VmZmVjdHMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX3NwZWxsc19lZmZlY3RzX2VudHJpZXMge1xuICAgICAgICBpbnQgbnBjX3NwZWxsc19lZmZlY3RzX2lkXG4gICAgfVxuICAgIG5wY19zcGVsbHNfZWZmZWN0cyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIG5wY19zcGVsbHNfZWZmZWN0c19lbnRyaWVzIHx8LS1veyBucGNfc3BlbGxzX2VmZmVjdHMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | npc_spells_effects_id | [npc_spells_effects](../../schema/npcs/npc_spells_effects.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Spell Effect Entry Identifier |
| npc_spells_effects_id | int | [NPC Spells Effects Identifier](npc_spells_effects.md) |
| spell_effect_id | smallint | [Spell Effect Identifier](../../../../server/spells/spell-effect-ids) |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| se_base | int | Spell Effect Base |
| se_limit | int | Spell Effect Limit |
| se_max | int | Spell Effect Maximum |

