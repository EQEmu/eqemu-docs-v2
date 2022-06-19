# npc_spells

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX3NwZWxscyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIGludHVuc2lnbmVkIHBhcmVudF9saXN0XG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciB6b25lX2lkXG4gICAgICAgIHZhcmNoYXIgb3duZXJfaWRcbiAgICB9XG4gICAgYm90X3NwZWxsc19lbnRyaWVzIHtcbiAgICAgICAgdmFyY2hhciBucGNfc3BlbGxzX2lkXG4gICAgICAgIHZhcmNoYXIgc3BlbGxfaWRcbiAgICAgICAgdmFyY2hhciB0eXBlXG4gICAgICAgIHZhcmNoYXIgc3BlbGxpZFxuICAgIH1cbiAgICBucGNfc3BlbGxzX2VudHJpZXMge1xuICAgICAgICBpbnQgbnBjX3NwZWxsc19pZFxuICAgICAgICBzbWFsbGludHVuc2lnbmVkIHNwZWxsaWRcbiAgICB9XG4gICAgbnBjX3NwZWxscyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIGludHVuc2lnbmVkIHBhcmVudF9saXN0XG4gICAgfVxuICAgIG5wY19zcGVsbHMgfHwtLW97IGJvdF9kYXRhIDogSGFzLU1hbnlcbiAgICBucGNfc3BlbGxzIHx8LS1veyBib3Rfc3BlbGxzX2VudHJpZXMgOiBIYXMtTWFueVxuICAgIG5wY19zcGVsbHMgfHwtLW97IG5wY19zcGVsbHNfZW50cmllcyA6IEhhcy1NYW55XG4gICAgbnBjX3NwZWxscyB8fC0tb3sgYm90X2RhdGEgOiBIYXMtTWFueVxuICAgIG5wY19zcGVsbHMgfHwtLW97IGJvdF9zcGVsbHNfZW50cmllcyA6IEhhcy1NYW55XG4gICAgbnBjX3NwZWxscyB8fC0tb3sgbnBjX3NwZWxscyA6IE9uZS10by1PbmVcbiAgICBucGNfc3BlbGxzIHx8LS1veyBucGNfc3BlbGxzX2VudHJpZXMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX3NwZWxscyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIGludHVuc2lnbmVkIHBhcmVudF9saXN0XG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciB6b25lX2lkXG4gICAgICAgIHZhcmNoYXIgb3duZXJfaWRcbiAgICB9XG4gICAgYm90X3NwZWxsc19lbnRyaWVzIHtcbiAgICAgICAgdmFyY2hhciBucGNfc3BlbGxzX2lkXG4gICAgICAgIHZhcmNoYXIgc3BlbGxfaWRcbiAgICAgICAgdmFyY2hhciB0eXBlXG4gICAgICAgIHZhcmNoYXIgc3BlbGxpZFxuICAgIH1cbiAgICBucGNfc3BlbGxzX2VudHJpZXMge1xuICAgICAgICBpbnQgbnBjX3NwZWxsc19pZFxuICAgICAgICBzbWFsbGludHVuc2lnbmVkIHNwZWxsaWRcbiAgICB9XG4gICAgbnBjX3NwZWxscyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIGludHVuc2lnbmVkIHBhcmVudF9saXN0XG4gICAgfVxuICAgIG5wY19zcGVsbHMgfHwtLW97IGJvdF9kYXRhIDogSGFzLU1hbnlcbiAgICBucGNfc3BlbGxzIHx8LS1veyBib3Rfc3BlbGxzX2VudHJpZXMgOiBIYXMtTWFueVxuICAgIG5wY19zcGVsbHMgfHwtLW97IG5wY19zcGVsbHNfZW50cmllcyA6IEhhcy1NYW55XG4gICAgbnBjX3NwZWxscyB8fC0tb3sgYm90X2RhdGEgOiBIYXMtTWFueVxuICAgIG5wY19zcGVsbHMgfHwtLW97IGJvdF9zcGVsbHNfZW50cmllcyA6IEhhcy1NYW55XG4gICAgbnBjX3NwZWxscyB8fC0tb3sgbnBjX3NwZWxscyA6IE9uZS10by1PbmVcbiAgICBucGNfc3BlbGxzIHx8LS1veyBucGNfc3BlbGxzX2VudHJpZXMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX3NwZWxscyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIGludHVuc2lnbmVkIHBhcmVudF9saXN0XG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciB6b25lX2lkXG4gICAgICAgIHZhcmNoYXIgb3duZXJfaWRcbiAgICB9XG4gICAgYm90X3NwZWxsc19lbnRyaWVzIHtcbiAgICAgICAgdmFyY2hhciBucGNfc3BlbGxzX2lkXG4gICAgICAgIHZhcmNoYXIgc3BlbGxfaWRcbiAgICAgICAgdmFyY2hhciB0eXBlXG4gICAgICAgIHZhcmNoYXIgc3BlbGxpZFxuICAgIH1cbiAgICBucGNfc3BlbGxzX2VudHJpZXMge1xuICAgICAgICBpbnQgbnBjX3NwZWxsc19pZFxuICAgICAgICBzbWFsbGludHVuc2lnbmVkIHNwZWxsaWRcbiAgICB9XG4gICAgbnBjX3NwZWxscyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgICAgIGludHVuc2lnbmVkIHBhcmVudF9saXN0XG4gICAgfVxuICAgIG5wY19zcGVsbHMgfHwtLW97IGJvdF9kYXRhIDogSGFzLU1hbnlcbiAgICBucGNfc3BlbGxzIHx8LS1veyBib3Rfc3BlbGxzX2VudHJpZXMgOiBIYXMtTWFueVxuICAgIG5wY19zcGVsbHMgfHwtLW97IG5wY19zcGVsbHNfZW50cmllcyA6IEhhcy1NYW55XG4gICAgbnBjX3NwZWxscyB8fC0tb3sgYm90X2RhdGEgOiBIYXMtTWFueVxuICAgIG5wY19zcGVsbHMgfHwtLW97IGJvdF9zcGVsbHNfZW50cmllcyA6IEhhcy1NYW55XG4gICAgbnBjX3NwZWxscyB8fC0tb3sgbnBjX3NwZWxscyA6IE9uZS10by1PbmVcbiAgICBucGNfc3BlbGxzIHx8LS1veyBucGNfc3BlbGxzX2VudHJpZXMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [bot_data](../../schema/bots/bot_data.md) | spells_id |
| Has-Many | id | [bot_spells_entries](../../schema/bots/bot_spells_entries.md) | npc_spells_id |
| Has-Many | id | [npc_spells_entries](../../schema/npcs/npc_spells_entries.md) | npc_spells_id |
| Has-Many | parent_list | [bot_data](../../schema/bots/bot_data.md) | spells_id |
| Has-Many | parent_list | [bot_spells_entries](../../schema/bots/bot_spells_entries.md) | npc_spells_id |
| One-to-One | parent_list | [npc_spells](../../schema/npcs/npc_spells.md) | parent_list |
| Has-Many | parent_list | [npc_spells_entries](../../schema/npcs/npc_spells_entries.md) | npc_spells_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique NPC Spell Set Identifier |
| name | tinytext | NPC Spell Set Name |
| parent_list | int | Inherit all the spells from this list, and merge them with our spells. Only one level of inheritance is allowed, so your parent's parent will not be included. |
| attack_proc | smallint | The combat proc that an NPC with this spell set will add to their list of procs. [Spell Identifier](../../schema/spells/spells_new.md) |
| proc_chance | tinyint | Proc Chance: 0 = Never, 100 = Always |
| range_proc | smallint | The ranged proc that an NPC with this spell set will add to their list of procs. |
| rproc_chance | smallint | Ranged Proc Chance: 0 = Never, 100 = Always |
| defensive_proc | smallint | The defensive proc that an NPC with this spell set will add to their list of procs. |
| dproc_chance | smallint | Defensive Proc Chance: 0 = Never, 100 = Always |
| fail_recast | int | Fail Recast |
| engaged_no_sp_recast_min | int | Engaged No Spell Recast Minimum (Unused) |
| engaged_no_sp_recast_max | int | Engaged No Spell Recast Maximum (Unused) |
| engaged_b_self_chance | tinyint | Engaged B Self Chance (Unused) |
| engaged_b_other_chance | tinyint | Engaged B Other Chance (Unused) |
| engaged_d_chance | tinyint | Engaged D Chance (Unused) |
| pursue_no_sp_recast_min | int | Pursue No Spell Recast Minimum (Unused) |
| pursue_no_sp_recast_max | int | Pursue No Spell Recast Maximum (Unused) |
| pursue_d_chance | tinyint | Pursue D Chance (Unused) |
| idle_no_sp_recast_min | int | Idle No Spell Recast Minimum (Unused) |
| idle_no_sp_recast_max | int | Idle No Spell Recast Maximum (Unused) |
| idle_b_chance | tinyint | Idle B Chance (Unused) |

