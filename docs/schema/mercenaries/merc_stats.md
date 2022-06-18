# merc_stats

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19zdGF0cyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfbnBjX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY19ucGNfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX25wY190eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfc3RhdHMgfHwtLW97IG1lcmNfbnBjX3R5cGVzIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19zdGF0cyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfbnBjX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY19ucGNfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX25wY190eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfc3RhdHMgfHwtLW97IG1lcmNfbnBjX3R5cGVzIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY19zdGF0cyB7XG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNfbnBjX3R5cGVfaWRcbiAgICB9XG4gICAgbWVyY19ucGNfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX25wY190eXBlX2lkXG4gICAgfVxuICAgIG1lcmNfc3RhdHMgfHwtLW97IG1lcmNfbnBjX3R5cGVzIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | merc_npc_type_id | [merc_npc_types](../../schema/mercenaries/merc_npc_types.md) | merc_npc_type_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merc_npc_type_id | int | [Mercenary NPC Type Identifier](merc_npc_types.md) |
| clientlevel | tinyint | Client Level |
| level | tinyint | Level |
| hp | int | Health |
| mana | int | Mana |
| AC | smallint | Armor Class |
| ATK | mediumint | Attack |
| STR | mediumint | Strength |
| STA | mediumint | Stamina |
| DEX | mediumint | Dexterity |
| AGI | mediumint | Agility |
| _INT | mediumint | Intelligence |
| WIS | mediumint | Wisdom |
| CHA | mediumint | Charisma |
| MR | smallint | Magic Resistance |
| CR | smallint | Cold Resistance |
| DR | smallint | Disease Resistance |
| FR | smallint | Fire Resistance |
| PR | smallint | Poison Resistance |
| Corrup | smallint | Corruption Resistance |
| mindmg | int | Minimum Damage |
| maxdmg | int | Maximum Damage |
| attack_count | smallint | Attack Count |
| attack_speed | tinyint | Attack Speed: The lower the number, the faster the NPC hits. (Deprecated) |
| attack_delay | tinyint | Attack Delay: Delay between the attack arounds in 10ths of a second. |
| special_abilities | text | Special Abilities |
| Accuracy | mediumint | Accuracy |
| hp_regen_rate | int | Health Regeneration Rate |
| mana_regen_rate | int | Mana Regeneration Rate |
| runspeed | float | Run Speed |
| statscale | int | Stat Scale: 50 = 50%, 100 = 100%, 150 = 150% |
| spellscale | float | Spell Scale: 50 = 50%, 100 = 100%, 150 = 150% |
| healscale | float | Heal Scale: 50 = 50%, 100 = 100%, 150 = 150% |

