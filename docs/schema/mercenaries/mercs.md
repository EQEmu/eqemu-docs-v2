# mercs

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    mercs {
        intunsigned OwnerCharacterID
        intunsigned TemplateID
        intunsigned MercID
        tinyintunsigned StanceID
    }
    merc_buffs {
        intunsigned MercId
        intunsigned SpellId
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    merc_spell_list_entries {
        varchar merc_spell_list_id
        varchar spell_id
        varchar stance_id
    }
    merc_stance_entries {
        varchar stance_id
    }
    merc_templates {
        varchar dbstring
        varchar merc_template_id
        varchar name_type_id
        varchar merc_npc_type_id
        varchar merc_subtype_id
        varchar merc_type_id
    }
    mercs ||--o{ merc_buffs : "Has-Many"
    mercs ||--o{ character_data : "One-to-One"
    mercs ||--o{ merc_spell_list_entries : "Has-Many"
    mercs ||--o{ merc_stance_entries : "Has-Many"
    mercs ||--o{ merc_templates : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | MercID | [merc_buffs](../../schema/mercenaries/merc_buffs.md) | MercId |
| One-to-One | OwnerCharacterID | [character_data](../../schema/characters/character_data.md) | id |
| Has-Many | StanceID | [merc_spell_list_entries](../../schema/mercenaries/merc_spell_list_entries.md) | stance_id |
| Has-Many | StanceID | [merc_stance_entries](../../schema/mercenaries/merc_stance_entries.md) | stance_id |
| One-to-One | TemplateID | [merc_templates](../../schema/mercenaries/merc_templates.md) | merc_template_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| MercID | int | Unique Mercenary Identifier |
| OwnerCharacterID | int | [Owner Character Identifier](../../schema/characters/character_data.md) |
| Slot | tinyint | Slot |
| Name | varchar | Name |
| TemplateID | int | [Template Identifier](merc_templates.md) |
| SuspendedTime | int | Suspended Time UNIX Timestamp |
| IsSuspended | tinyint | Is Suspended: 0 = False, 1 = True |
| TimerRemaining | int | Timer Remaining in Seconds |
| Gender | tinyint | [Gender](../../../../server/npc/genders) |
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

