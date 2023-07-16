# merc_types

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgfVxuICAgIGRiX3N0ciB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBtZXJjX3RlbXBsYXRlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfdGVtcGxhdGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgICAgICB0aW55aW50IG5hbWVfdHlwZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX3N1YnR5cGVfaWRcbiAgICB9XG4gICAgbWVyY190eXBlcyB8fC0tb3sgZGJfc3RyIDogXCJPbmUtdG8tT25lXCJcbiAgICBtZXJjX3R5cGVzIHx8LS1veyBtZXJjX3RlbXBsYXRlcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgfVxuICAgIGRiX3N0ciB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBtZXJjX3RlbXBsYXRlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfdGVtcGxhdGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgICAgICB0aW55aW50IG5hbWVfdHlwZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX3N1YnR5cGVfaWRcbiAgICB9XG4gICAgbWVyY190eXBlcyB8fC0tb3sgZGJfc3RyIDogXCJPbmUtdG8tT25lXCJcbiAgICBtZXJjX3R5cGVzIHx8LS1veyBtZXJjX3RlbXBsYXRlcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY190eXBlcyB7XG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgfVxuICAgIGRiX3N0ciB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBtZXJjX3RlbXBsYXRlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfdGVtcGxhdGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY190eXBlX2lkXG4gICAgICAgIHZhcmNoYXIgZGJzdHJpbmdcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY19ucGNfdHlwZV9pZFxuICAgICAgICB0aW55aW50IG5hbWVfdHlwZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX3N1YnR5cGVfaWRcbiAgICB9XG4gICAgbWVyY190eXBlcyB8fC0tb3sgZGJfc3RyIDogXCJPbmUtdG8tT25lXCJcbiAgICBtZXJjX3R5cGVzIHx8LS1veyBtZXJjX3RlbXBsYXRlcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


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

