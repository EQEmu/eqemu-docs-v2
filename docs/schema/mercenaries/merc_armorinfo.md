# merc_armorinfo

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19hcm1vcmluZm8ge1xuICAgICAgICBpbnR1bnNpZ25lZCBhcm1vcnRpbnRfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgIH1cbiAgICBucGNfdHlwZXNfdGludCB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIG1lcmNfbnBjX3R5cGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgIH1cbiAgICBtZXJjX2FybW9yaW5mbyB8fC0tb3sgbnBjX3R5cGVzX3RpbnQgOiBIYXMtTWFueVxuICAgIG1lcmNfYXJtb3JpbmZvIHx8LS1veyBtZXJjX25wY190eXBlcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19hcm1vcmluZm8ge1xuICAgICAgICBpbnR1bnNpZ25lZCBhcm1vcnRpbnRfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgIH1cbiAgICBucGNfdHlwZXNfdGludCB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIG1lcmNfbnBjX3R5cGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgIH1cbiAgICBtZXJjX2FybW9yaW5mbyB8fC0tb3sgbnBjX3R5cGVzX3RpbnQgOiBIYXMtTWFueVxuICAgIG1lcmNfYXJtb3JpbmZvIHx8LS1veyBtZXJjX25wY190eXBlcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19hcm1vcmluZm8ge1xuICAgICAgICBpbnR1bnNpZ25lZCBhcm1vcnRpbnRfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgIH1cbiAgICBucGNfdHlwZXNfdGludCB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIG1lcmNfbnBjX3R5cGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgIH1cbiAgICBtZXJjX2FybW9yaW5mbyB8fC0tb3sgbnBjX3R5cGVzX3RpbnQgOiBIYXMtTWFueVxuICAgIG1lcmNfYXJtb3JpbmZvIHx8LS1veyBtZXJjX25wY190eXBlcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


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
| helmtexture | tinyint | [Helmet Texture](../../../../categories/npc/textures) |
| armortint_id | int | [Armor Tint Identifier](../../schema/npcs/npc_types_tint.md) |
| armortint_red | tinyint | Armor Tint Red: 0 = None, 255 = Max |
| armortint_green | tinyint | Armor Tint Green: 0 = None, 255 = Max |
| armortint_blue | tinyint | Armor Tint Blue: 0 = None, 255 = Max |

