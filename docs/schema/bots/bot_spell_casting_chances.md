# bot_spell_casting_chances

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB7XG4gICAgICAgIHZhcmNoYXIgc3BlbGxfdHlwZV9pbmRleFxuICAgICAgICB2YXJjaGFyIHN0YW5jZV9pbmRleFxuICAgIH1cbiAgICBib3Rfc3BlbGxzX2VudHJpZXMge1xuICAgICAgICB2YXJjaGFyIHR5cGVcbiAgICAgICAgdmFyY2hhciBzcGVsbGlkXG4gICAgICAgIHZhcmNoYXIgbnBjX3NwZWxsc19pZFxuICAgICAgICB2YXJjaGFyIHNwZWxsX2lkXG4gICAgfVxuICAgIGJvdF9zcGVsbF9jYXN0aW5nX2NoYW5jZXMgfHwtLW97IGJvdF9zcGVsbHNfZW50cmllcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB7XG4gICAgICAgIHZhcmNoYXIgc3BlbGxfdHlwZV9pbmRleFxuICAgICAgICB2YXJjaGFyIHN0YW5jZV9pbmRleFxuICAgIH1cbiAgICBib3Rfc3BlbGxzX2VudHJpZXMge1xuICAgICAgICB2YXJjaGFyIHR5cGVcbiAgICAgICAgdmFyY2hhciBzcGVsbGlkXG4gICAgICAgIHZhcmNoYXIgbnBjX3NwZWxsc19pZFxuICAgICAgICB2YXJjaGFyIHNwZWxsX2lkXG4gICAgfVxuICAgIGJvdF9zcGVsbF9jYXN0aW5nX2NoYW5jZXMgfHwtLW97IGJvdF9zcGVsbHNfZW50cmllcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB7XG4gICAgICAgIHZhcmNoYXIgc3BlbGxfdHlwZV9pbmRleFxuICAgICAgICB2YXJjaGFyIHN0YW5jZV9pbmRleFxuICAgIH1cbiAgICBib3Rfc3BlbGxzX2VudHJpZXMge1xuICAgICAgICB2YXJjaGFyIHR5cGVcbiAgICAgICAgdmFyY2hhciBzcGVsbGlkXG4gICAgICAgIHZhcmNoYXIgbnBjX3NwZWxsc19pZFxuICAgICAgICB2YXJjaGFyIHNwZWxsX2lkXG4gICAgfVxuICAgIGJvdF9zcGVsbF9jYXN0aW5nX2NoYW5jZXMgfHwtLW97IGJvdF9zcGVsbHNfZW50cmllcyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | spell_type_index | [bot_spells_entries](../../schema/bots/bot_spells_entries.md) | type |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Bot Spell Casting Chance Identifier |
| spell_type_index | tinyint | [Spell Type Identifier](../../../../categories/spells/spell-types) |
| class_id | tinyint | [Class Identifier](../../../../server/player/class-list) |
| stance_index | tinyint | [Stance Type Identifier](../../../../categories/bots/stance-types) |
| nHSND_value | tinyint | Negative Healer/Slower/Nuker/Doter Value |
| pH_value | tinyint | Positive Healer Value |
| pS_value | tinyint | Positive Slower Value |
| pHS_value | tinyint | Positive Healer/Slower Value |
| pN_value | tinyint | Positive Nuker Value |
| pHN_value | tinyint | Positive Healer/Nuker Value |
| pSN_value | tinyint | Positive Slower/Nuker Value |
| pHSN_value | tinyint | Positive Healer/Slower/Nuker Value |
| pD_value | tinyint | Positive Doter Value |
| pHD_value | tinyint | Positive Healer/Doter Value |
| pSD_value | tinyint | Positive Slower/Doter Value |
| pHSD_value | tinyint | Positive Healer/Slower/Doter Value |
| pND_value | tinyint | Positive Nuker/Doter Value |
| pHND_value | tinyint | Positive Healer/Nuker/Doter Value |
| pSND_value | tinyint | Positive Slower/Nuker/Doter Value |
| pHSND_value | tinyint | Positive Healer/Slower/Nuker/Doter Value |

