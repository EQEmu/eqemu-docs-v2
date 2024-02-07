# merc_armorinfo

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19hcm1vcmluZm8ge1xuICAgICAgICB2YXJjaGFyIGFybW9ydGludF9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfbnBjX3R5cGVfaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzX3RpbnQge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBtZXJjX25wY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgbWVyY19ucGNfdHlwZV9pZFxuICAgIH1cbiAgICBtZXJjX2FybW9yaW5mbyB8fC0tb3sgbnBjX3R5cGVzX3RpbnQgOiBcIkhhcy1NYW55XCJcbiAgICBtZXJjX2FybW9yaW5mbyB8fC0tb3sgbWVyY19ucGNfdHlwZXMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19hcm1vcmluZm8ge1xuICAgICAgICB2YXJjaGFyIGFybW9ydGludF9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfbnBjX3R5cGVfaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzX3RpbnQge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBtZXJjX25wY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgbWVyY19ucGNfdHlwZV9pZFxuICAgIH1cbiAgICBtZXJjX2FybW9yaW5mbyB8fC0tb3sgbnBjX3R5cGVzX3RpbnQgOiBcIkhhcy1NYW55XCJcbiAgICBtZXJjX2FybW9yaW5mbyB8fC0tb3sgbWVyY19ucGNfdHlwZXMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19hcm1vcmluZm8ge1xuICAgICAgICB2YXJjaGFyIGFybW9ydGludF9pZFxuICAgICAgICB2YXJjaGFyIG1lcmNfbnBjX3R5cGVfaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzX3RpbnQge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBtZXJjX25wY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgbWVyY19ucGNfdHlwZV9pZFxuICAgIH1cbiAgICBtZXJjX2FybW9yaW5mbyB8fC0tb3sgbnBjX3R5cGVzX3RpbnQgOiBcIkhhcy1NYW55XCJcbiAgICBtZXJjX2FybW9yaW5mbyB8fC0tb3sgbWVyY19ucGNfdHlwZXMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | armortint_id | [npc_types_tint](../../schema/npcs/npc_types_tint.md) | id |
| One-to-One | merc_npc_type_id | [merc_npc_types](../../schema/mercenaries/merc_npc_types.md) | merc_npc_type_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Mercenary Armor Info Identifier |
| merc_npc_type_id | int | [Mercenary NPC Type Identifier](merc_npc_types.md) |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| texture | tinyint | [Texture](../../../../server/npc/textures) |
| helmtexture | tinyint | [Helmet Texture](../../../../server/npc/textures) |
| armortint_id | int | [Armor Tint Identifier](../../schema/npcs/npc_types_tint.md) |
| armortint_red | tinyint | Armor Tint Red: 0 = None, 255 = Max |
| armortint_green | tinyint | Armor Tint Green: 0 = None, 255 = Max |
| armortint_blue | tinyint | Armor Tint Blue: 0 = None, 255 = Max |

