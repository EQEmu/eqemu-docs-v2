# character_expedition_lockouts

!!! info
	This page was last generated 2022.06.17

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Lockout Identifier |
| character_id | int | [Character Identifier](character_data.md) |
| expedition_name | varchar | Expedition Name |
| event_name | varchar | Event Name |
| expire_time | datetime | Expire Time |
| duration | int | Duration in Seconds |
| from_expedition_uuid | varchar | From Expedition UUID |

