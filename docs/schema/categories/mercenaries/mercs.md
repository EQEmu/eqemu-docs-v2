# mercs

| Column | Data Type | Description |
| :--- | :--- | :--- |
| MercID | int | Unique Mercenary Identifier |
| OwnerCharacterID | int | [Owner Character Identifier](../../../schema/categories/mercenaries/character_data.md) |
| Slot | tinyint | Slot |
| Name | varchar | Name |
| TemplateID | int | [Template Identifier](merc_templates.md) |
| SuspendedTime | int | Suspended Time UNIX Timestamp |
| IsSuspended | tinyint | Is Suspended: 0 = False, 1 = True |
| TimerRemaining | int | Timer Remaining in Seconds |
| Gender | tinyint | [Gender](../../../../categories/npc/genders) |
| MercSize | float | Mercenary Size |
| StanceID | tinyint | [Stance Type Identifier](../../../../categories/bots/stance-types) |
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

