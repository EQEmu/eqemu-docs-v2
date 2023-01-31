# merc_weaponinfo

!!! info
	This page was last generated 2023.01.30

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY193ZWFwb25pbmZvIHtcbiAgICAgICAgaW50IG1lcmNfbnBjX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY19ucGNfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX25wY190eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfd2VhcG9uaW5mbyB8fC0tb3sgbWVyY19ucGNfdHlwZXMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY193ZWFwb25pbmZvIHtcbiAgICAgICAgaW50IG1lcmNfbnBjX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY19ucGNfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX25wY190eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfd2VhcG9uaW5mbyB8fC0tb3sgbWVyY19ucGNfdHlwZXMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY193ZWFwb25pbmZvIHtcbiAgICAgICAgaW50IG1lcmNfbnBjX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY19ucGNfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX25wY190eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfd2VhcG9uaW5mbyB8fC0tb3sgbWVyY19ucGNfdHlwZXMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | merc_npc_type_id | [merc_npc_types](../../schema/mercenaries/merc_npc_types.md) | merc_npc_type_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Mercenary Weapon Info Identifier |
| merc_npc_type_id | int | [Mercenary NPC Type Identifier](merc_npc_types.md) |
| minlevel | tinyint | Minimum Level |
| maxlevel | tinyint | Maximum Level |
| d_melee_texture1 | int | Primary Weapon Texture |
| d_melee_texture2 | int | Secondary Weapon Texture |
| prim_melee_type | tinyint | [Primary Melee Type](../../../../server/player/skills) |
| sec_melee_type | tinyint | [Secondary Melee Type](../../../../categories/player/skills) |

