# bot_spell_casting_chances

!!! info
	This page was last generated 2023.06.12

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB7XG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCBzcGVsbF90eXBlX2luZGV4XG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCBzdGFuY2VfaW5kZXhcbiAgICB9XG4gICAgYm90X3NwZWxsc19lbnRyaWVzIHtcbiAgICAgICAgaW50IG5wY19zcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB0eXBlXG4gICAgICAgIHNtYWxsaW50IHNwZWxsaWRcbiAgICB9XG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB8fC0tb3sgYm90X3NwZWxsc19lbnRyaWVzIDogXCJIYXMtTWFueVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB7XG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCBzcGVsbF90eXBlX2luZGV4XG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCBzdGFuY2VfaW5kZXhcbiAgICB9XG4gICAgYm90X3NwZWxsc19lbnRyaWVzIHtcbiAgICAgICAgaW50IG5wY19zcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB0eXBlXG4gICAgICAgIHNtYWxsaW50IHNwZWxsaWRcbiAgICB9XG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB8fC0tb3sgYm90X3NwZWxsc19lbnRyaWVzIDogXCJIYXMtTWFueVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB7XG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCBzcGVsbF90eXBlX2luZGV4XG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCBzdGFuY2VfaW5kZXhcbiAgICB9XG4gICAgYm90X3NwZWxsc19lbnRyaWVzIHtcbiAgICAgICAgaW50IG5wY19zcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB0eXBlXG4gICAgICAgIHNtYWxsaW50IHNwZWxsaWRcbiAgICB9XG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB8fC0tb3sgYm90X3NwZWxsc19lbnRyaWVzIDogXCJIYXMtTWFueVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | spell_type_index | [bot_spells_entries](../../schema/bots/bot_spells_entries.md) | type |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Bot Spell Casting Chance Identifier |
| spell_type_index | tinyint | [Spell Type Identifier](../../../../categories/spells/spell-types) |
| class_id | tinyint | [Class Identifier](../../../../categories/player/class-list) |
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

