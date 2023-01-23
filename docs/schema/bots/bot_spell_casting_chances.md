# bot_spell_casting_chances

!!! info
	This page was last generated 2023.01.22

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB7XG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCBzcGVsbF90eXBlX2luZGV4XG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCBzdGFuY2VfaW5kZXhcbiAgICB9XG4gICAgYm90X3NwZWxsc19lbnRyaWVzIHtcbiAgICAgICAgaW50IG5wY19zcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB0eXBlXG4gICAgICAgIHNtYWxsaW50IHNwZWxsaWRcbiAgICB9XG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB8fC0tb3sgYm90X3NwZWxsc19lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB7XG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCBzcGVsbF90eXBlX2luZGV4XG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCBzdGFuY2VfaW5kZXhcbiAgICB9XG4gICAgYm90X3NwZWxsc19lbnRyaWVzIHtcbiAgICAgICAgaW50IG5wY19zcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB0eXBlXG4gICAgICAgIHNtYWxsaW50IHNwZWxsaWRcbiAgICB9XG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB8fC0tb3sgYm90X3NwZWxsc19lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB7XG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCBzcGVsbF90eXBlX2luZGV4XG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCBzdGFuY2VfaW5kZXhcbiAgICB9XG4gICAgYm90X3NwZWxsc19lbnRyaWVzIHtcbiAgICAgICAgaW50IG5wY19zcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCB0eXBlXG4gICAgICAgIHNtYWxsaW50IHNwZWxsaWRcbiAgICB9XG4gICAgYm90X3NwZWxsX2Nhc3RpbmdfY2hhbmNlcyB8fC0tb3sgYm90X3NwZWxsc19lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


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
| nHSND_value | tinyint | nHSND Value |
| pH_value | tinyint | pH Value |
| pS_value | tinyint | pS Value |
| pHS_value | tinyint | pHS Value |
| pN_value | tinyint | pN Value |
| pHN_value | tinyint | pHN Value |
| pSN_value | tinyint | pSN Value |
| pHSN_value | tinyint | pHSN Value |
| pD_value | tinyint | pD Value |
| pHD_value | tinyint | pHD Value |
| pSD_value | tinyint | pSD Value |
| pHSD_value | tinyint | pHSD Value |
| pND_value | tinyint | pND Value |
| pHND_value | tinyint | pHND Value |
| pSND_value | tinyint | pSND Value |
| pHSND_value | tinyint | pHSND Value |

