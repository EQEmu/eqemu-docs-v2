# npc_scale_global_base

!!! info
	This page was last generated 2024.02.07

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| type | int | Type: 0 = Trash, 1 = Named, 2 = Raid |
| level | int | Level |
| zone_id_list | text | [Zone Identifier List](../../../../server/zones/zone-list) (Pipe (|) Separated) |
| instance_version_list | text | Instance Version List (Pipe (|) Separated) |
| ac | int | Armor Class |
| hp | bigint | Health |
| accuracy | int | Accuracy |
| slow_mitigation | int | Slow Mitigation |
| attack | int | Attack |
| strength | int | Strength |
| stamina | int | Stamina |
| dexterity | int | Dexterity |
| agility | int | Agility |
| intelligence | int | Intelligence |
| wisdom | int | Wisdom |
| charisma | int | Charisma |
| magic_resist | int | Magic Resistance |
| cold_resist | int | Cold Resistance |
| fire_resist | int | Fire Resistance |
| poison_resist | int | Poison Resistance |
| disease_resist | int | Disease Resistance |
| corruption_resist | int | Corruption Resistance |
| physical_resist | int | Physical Resistance |
| min_dmg | int | Minimum Damage |
| max_dmg | int | Maximum Damage |
| hp_regen_rate | bigint | Health Regeneration Rate |
| hp_regen_per_second | bigint | Health Regeneration Per Second |
| attack_delay | int | Attack Delay: Delay between the attack arounds in 10ths of a second. |
| spell_scale | int | Spell Scale: 50 = 50%, 100 = 100%, 150 = 150% |
| heal_scale | int | Heal Scale: 50 = 50%, 100 = 100%, 150 = 150% |
| avoidance | int | Avoidance |
| heroic_strikethrough | int | Heroic Strikethrough |
| special_abilities | text | Special Abilities |

