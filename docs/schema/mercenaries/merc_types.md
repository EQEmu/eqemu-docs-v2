# merc_types

!!! info
	This page was last generated 2022.06.17

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfdHlwZV9pZFxuICAgIH1cbiAgICBtZXJjX3RlbXBsYXRlcyB7XG4gICAgICAgIHRpbnlpbnQgbmFtZV90eXBlX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfbnBjX3R5cGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190ZW1wbGF0ZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX3N1YnR5cGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfdHlwZXMgfHwtLW97IG1lcmNfdGVtcGxhdGVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfdHlwZV9pZFxuICAgIH1cbiAgICBtZXJjX3RlbXBsYXRlcyB7XG4gICAgICAgIHRpbnlpbnQgbmFtZV90eXBlX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfbnBjX3R5cGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190ZW1wbGF0ZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX3N1YnR5cGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfdHlwZXMgfHwtLW97IG1lcmNfdGVtcGxhdGVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfdHlwZV9pZFxuICAgIH1cbiAgICBtZXJjX3RlbXBsYXRlcyB7XG4gICAgICAgIHRpbnlpbnQgbmFtZV90eXBlX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfbnBjX3R5cGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190ZW1wbGF0ZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX3N1YnR5cGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfdHlwZXMgfHwtLW97IG1lcmNfdGVtcGxhdGVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | merc_type_id | [merc_templates](../../schema/mercenaries/merc_templates.md) | merc_type_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merc_type_id | int | Unique Mercenary Type Identifier |
| race_id | int | [Race Identifier](../../../../categories/npc/race-list) |
| proficiency_id | tinyint | Proficiency Identifier |
| dbstring | varchar | [DBString Identifier](db_str.md) |
| clientversion | int | [Client Version](../../../../categories/player/client-version-bitmasks) |

