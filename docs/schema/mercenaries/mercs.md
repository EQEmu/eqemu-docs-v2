# mercs

!!! info
	This page was last generated 2022.06.17

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY3Mge1xuICAgICAgICBpbnR1bnNpZ25lZCBNZXJjSURcbiAgICB9XG4gICAgbWVyY19idWZmcyB7XG4gICAgICAgICBNZXJjSURcbiAgICB9XG4gICAgbWVyY3MgfHwtLW97IG1lcmNfYnVmZnMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY3Mge1xuICAgICAgICBpbnR1bnNpZ25lZCBNZXJjSURcbiAgICB9XG4gICAgbWVyY19idWZmcyB7XG4gICAgICAgICBNZXJjSURcbiAgICB9XG4gICAgbWVyY3MgfHwtLW97IG1lcmNfYnVmZnMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY3Mge1xuICAgICAgICBpbnR1bnNpZ25lZCBNZXJjSURcbiAgICB9XG4gICAgbWVyY19idWZmcyB7XG4gICAgICAgICBNZXJjSURcbiAgICB9XG4gICAgbWVyY3MgfHwtLW97IG1lcmNfYnVmZnMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | MercID | [merc_buffs](../../schema/mercenaries/merc_buffs.md) | MercID |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| MercID | int | Unique Mercenary Identifier |
| OwnerCharacterID | int | [Owner Character Identifier](../../../schema/categories/characters/character_data.md) |
| Slot | tinyint | Slot |
| Name | varchar | Name |
| TemplateID | int | [Template Identifier](merc_templates.md) |
| SuspendedTime | int | Suspended Time UNIX Timestamp |
| IsSuspended | tinyint | Is Suspended: 0 = False, 1 = True |
| TimerRemaining | int | Timer Remaining in Seconds |
| Gender | tinyint | [Gender](../../../../categories/npc/genders) |
| MercSize | float | Mercenary Size |
| StanceID | tinyint | [Stance Type Identifier](../../../../server/bots/stance-types) |
| HP | int | Health |
| Mana | int | Mana |
| Endurance | int | Endurance |
| Face | int | Face |
| LuclinHairStyle | int | Hair Style |
| LuclinHairColor | int | Hair Color |
| LuclinEyeColor | int | Eye Color 1 |
| LuclinEyeColor2 | int | Eye Color 2 |
| LuclinBeardColor | int | Beard Color |
| LuclinBeard | int | Beard |
| DrakkinHeritage | int | Drakkin Heritage |
| DrakkinTattoo | int | Drakkin Tattoo |
| DrakkinDetails | int | Drakkin Details |

