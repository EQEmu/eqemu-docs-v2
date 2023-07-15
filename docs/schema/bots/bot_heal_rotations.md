# bot_heal_rotations

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2hlYWxfcm90YXRpb25zIHtcbiAgICAgICAgaW50dW5zaWduZWQgaGVhbF9yb3RhdGlvbl9pbmRleFxuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIHNtYWxsaW50IHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgb3duZXJfaWRcbiAgICB9XG4gICAgYm90X2hlYWxfcm90YXRpb25fbWVtYmVycyB7XG4gICAgICAgIGludHVuc2lnbmVkIGJvdF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBoZWFsX3JvdGF0aW9uX2luZGV4XG4gICAgfVxuICAgIGJvdF9oZWFsX3JvdGF0aW9uX3RhcmdldHMge1xuICAgICAgICBpbnR1bnNpZ25lZCBoZWFsX3JvdGF0aW9uX2luZGV4XG4gICAgfVxuICAgIGJvdF9oZWFsX3JvdGF0aW9ucyB8fC0tb3sgYm90X2RhdGEgOiBcIk9uZS10by1PbmVcIlxuICAgIGJvdF9oZWFsX3JvdGF0aW9ucyB8fC0tb3sgYm90X2hlYWxfcm90YXRpb25fbWVtYmVycyA6IFwiSGFzLU1hbnlcIlxuICAgIGJvdF9oZWFsX3JvdGF0aW9ucyB8fC0tb3sgYm90X2hlYWxfcm90YXRpb25fdGFyZ2V0cyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2hlYWxfcm90YXRpb25zIHtcbiAgICAgICAgaW50dW5zaWduZWQgaGVhbF9yb3RhdGlvbl9pbmRleFxuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIHNtYWxsaW50IHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgb3duZXJfaWRcbiAgICB9XG4gICAgYm90X2hlYWxfcm90YXRpb25fbWVtYmVycyB7XG4gICAgICAgIGludHVuc2lnbmVkIGJvdF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBoZWFsX3JvdGF0aW9uX2luZGV4XG4gICAgfVxuICAgIGJvdF9oZWFsX3JvdGF0aW9uX3RhcmdldHMge1xuICAgICAgICBpbnR1bnNpZ25lZCBoZWFsX3JvdGF0aW9uX2luZGV4XG4gICAgfVxuICAgIGJvdF9oZWFsX3JvdGF0aW9ucyB8fC0tb3sgYm90X2RhdGEgOiBcIk9uZS10by1PbmVcIlxuICAgIGJvdF9oZWFsX3JvdGF0aW9ucyB8fC0tb3sgYm90X2hlYWxfcm90YXRpb25fbWVtYmVycyA6IFwiSGFzLU1hbnlcIlxuICAgIGJvdF9oZWFsX3JvdGF0aW9ucyB8fC0tb3sgYm90X2hlYWxfcm90YXRpb25fdGFyZ2V0cyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2hlYWxfcm90YXRpb25zIHtcbiAgICAgICAgaW50dW5zaWduZWQgaGVhbF9yb3RhdGlvbl9pbmRleFxuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIHNtYWxsaW50IHpvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgb3duZXJfaWRcbiAgICB9XG4gICAgYm90X2hlYWxfcm90YXRpb25fbWVtYmVycyB7XG4gICAgICAgIGludHVuc2lnbmVkIGJvdF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBoZWFsX3JvdGF0aW9uX2luZGV4XG4gICAgfVxuICAgIGJvdF9oZWFsX3JvdGF0aW9uX3RhcmdldHMge1xuICAgICAgICBpbnR1bnNpZ25lZCBoZWFsX3JvdGF0aW9uX2luZGV4XG4gICAgfVxuICAgIGJvdF9oZWFsX3JvdGF0aW9ucyB8fC0tb3sgYm90X2RhdGEgOiBcIk9uZS10by1PbmVcIlxuICAgIGJvdF9oZWFsX3JvdGF0aW9ucyB8fC0tb3sgYm90X2hlYWxfcm90YXRpb25fbWVtYmVycyA6IFwiSGFzLU1hbnlcIlxuICAgIGJvdF9oZWFsX3JvdGF0aW9ucyB8fC0tb3sgYm90X2hlYWxfcm90YXRpb25fdGFyZ2V0cyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_id | [bot_data](../../schema/bots/bot_data.md) | bot_id |
| Has-Many | heal_rotation_index | [bot_heal_rotation_members](../../schema/bots/bot_heal_rotation_members.md) | heal_rotation_index |
| Has-Many | heal_rotation_index | [bot_heal_rotation_targets](../../schema/bots/bot_heal_rotation_targets.md) | heal_rotation_index |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| heal_rotation_index | int | Unique Heal Rotation Identifier |
| bot_id | int | [Bot Identifier](bot_data.md) |
| interval | int | Interval |
| fast_heals | int | Fast Heals |
| adaptive_targeting | int | Adaptive Targeting |
| casting_override | int | Casting Override |
| safe_hp_base | float | Safe Health Base |
| safe_hp_cloth | float | Safe Health Cloth |
| safe_hp_leather | float | Safe Health Leather |
| safe_hp_chain | float | Safe Health Chain |
| safe_hp_plate | float | Safe Health Plate |
| critical_hp_base | float | Critical Health Base |
| critical_hp_cloth | float | Critical Health Cloth |
| critical_hp_leather | float | Critical Health Leather |
| critical_hp_chain | float | Critical Health Chain |
| critical_hp_plate | float | Critical Health Plate |

