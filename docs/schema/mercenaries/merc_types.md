# merc_types

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgfVxuICAgIGRiX3N0ciB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBtZXJjX3RlbXBsYXRlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfdGVtcGxhdGVfaWRcbiAgICAgICAgdGlueWludCBuYW1lX3R5cGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIGRic3RyaW5nXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfc3VidHlwZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY190eXBlcyB8fC0tb3sgZGJfc3RyIDogT25lLXRvLU9uZVxuICAgIG1lcmNfdHlwZXMgfHwtLW97IG1lcmNfdGVtcGxhdGVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgfVxuICAgIGRiX3N0ciB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBtZXJjX3RlbXBsYXRlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfdGVtcGxhdGVfaWRcbiAgICAgICAgdGlueWludCBuYW1lX3R5cGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIGRic3RyaW5nXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfc3VidHlwZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY190eXBlcyB8fC0tb3sgZGJfc3RyIDogT25lLXRvLU9uZVxuICAgIG1lcmNfdHlwZXMgfHwtLW97IG1lcmNfdGVtcGxhdGVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgfVxuICAgIGRiX3N0ciB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBtZXJjX3RlbXBsYXRlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfdGVtcGxhdGVfaWRcbiAgICAgICAgdGlueWludCBuYW1lX3R5cGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIGRic3RyaW5nXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfc3VidHlwZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY190eXBlcyB8fC0tb3sgZGJfc3RyIDogT25lLXRvLU9uZVxuICAgIG1lcmNfdHlwZXMgfHwtLW97IG1lcmNfdGVtcGxhdGVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | dbstring | [db_str](../../schema/client-files/db_str.md) | id |
| Has-Many | merc_type_id | [merc_templates](../../schema/mercenaries/merc_templates.md) | merc_type_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merc_type_id | int | Unique Mercenary Type Identifier |
| race_id | int | [Race Identifier](../../../../categories/npc/race-list) |
| proficiency_id | tinyint | Proficiency Identifier |
| dbstring | varchar | [DBString Identifier](../../schema/client-files/db_str.md) |
| clientversion | int | [Client Version](../../../../categories/player/client-version-bitmasks) |

