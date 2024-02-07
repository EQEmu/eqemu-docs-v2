# player_event_logs

!!! info
	This page was last generated 2024.02.07

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | bigint | Unique Player Event Log Identifier |
| account_id | bigint | [Account Identifier](../../schema/account/account.md) |
| character_id | bigint | [Character Identifier](../../schema/characters/character_data.md) |
| zone_id | int | [Zone Identifier](../../../../server/zones/zone-list) |
| instance_id | int | [Instance Identifier](../../schema/instances/instance_list.md) |
| x | float | X Coordinate |
| y | float | Y  Coordinate |
| z | float | Z  Coordinate |
| heading | float | Heading Coordinate |
| event_type_id | int | [Event Type Identifier](../../schema/admin/player_event_log_settings.md) |
| event_type_name | varchar | [Event Type Name](../../schema/admin/player_event_log_settings.md) |
| event_data | longtext | Event Data JSON |
| created_at | datetime | Created At |

