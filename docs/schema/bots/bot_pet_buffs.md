# bot_pet_buffs

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3BldF9idWZmcyB7XG4gICAgICAgIHZhcmNoYXIgcGV0c19pbmRleFxuICAgICAgICB2YXJjaGFyIHNwZWxsX2lkXG4gICAgfVxuICAgIGJvdF9wZXRzIHtcbiAgICAgICAgdmFyY2hhciBwZXRzX2luZGV4XG4gICAgICAgIHZhcmNoYXIgYm90X2lkXG4gICAgICAgIHZhcmNoYXIgc3BlbGxfaWRcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIHRlbGVwb3J0X3pvbmVcbiAgICAgICAgaW50IGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW0yXG4gICAgICAgIGludCB0eXBlZGVzY251bVxuICAgIH1cbiAgICBib3RfcGV0X2J1ZmZzIHx8LS1veyBib3RfcGV0cyA6IE9uZS10by1PbmVcbiAgICBib3RfcGV0X2J1ZmZzIHx8LS1veyBzcGVsbHNfbmV3IDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3BldF9idWZmcyB7XG4gICAgICAgIHZhcmNoYXIgcGV0c19pbmRleFxuICAgICAgICB2YXJjaGFyIHNwZWxsX2lkXG4gICAgfVxuICAgIGJvdF9wZXRzIHtcbiAgICAgICAgdmFyY2hhciBwZXRzX2luZGV4XG4gICAgICAgIHZhcmNoYXIgYm90X2lkXG4gICAgICAgIHZhcmNoYXIgc3BlbGxfaWRcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIHRlbGVwb3J0X3pvbmVcbiAgICAgICAgaW50IGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW0yXG4gICAgICAgIGludCB0eXBlZGVzY251bVxuICAgIH1cbiAgICBib3RfcGV0X2J1ZmZzIHx8LS1veyBib3RfcGV0cyA6IE9uZS10by1PbmVcbiAgICBib3RfcGV0X2J1ZmZzIHx8LS1veyBzcGVsbHNfbmV3IDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3BldF9idWZmcyB7XG4gICAgICAgIHZhcmNoYXIgcGV0c19pbmRleFxuICAgICAgICB2YXJjaGFyIHNwZWxsX2lkXG4gICAgfVxuICAgIGJvdF9wZXRzIHtcbiAgICAgICAgdmFyY2hhciBwZXRzX2luZGV4XG4gICAgICAgIHZhcmNoYXIgYm90X2lkXG4gICAgICAgIHZhcmNoYXIgc3BlbGxfaWRcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIHRlbGVwb3J0X3pvbmVcbiAgICAgICAgaW50IGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW1cbiAgICAgICAgaW50IGVmZmVjdGRlc2NudW0yXG4gICAgICAgIGludCB0eXBlZGVzY251bVxuICAgIH1cbiAgICBib3RfcGV0X2J1ZmZzIHx8LS1veyBib3RfcGV0cyA6IE9uZS10by1PbmVcbiAgICBib3RfcGV0X2J1ZmZzIHx8LS1veyBzcGVsbHNfbmV3IDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | pets_index | [bot_pets](../../schema/bots/bot_pets.md) | pets_index |
| One-to-One | spell_id | [spells_new](../../schema/spells/spells_new.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| pet_buffs_index | int | Unique Bot Pet Buffs Identifier |
| pets_index | int | [Bot Pet Identifier](bot_pets.md) |
| spell_id | int | [Spell Identifier](../../../schema/spells/spells_new.md) |
| caster_level | int | Caster Level |
| duration | int | Duration of buff |

