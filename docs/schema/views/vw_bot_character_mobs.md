# vw_bot_character_mobs

!!! info
	This page was last generated 2023.07.15

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| mob_type | varchar | Mob Type: B = Bot, C = Client |
| id | int | Unique View Bot Character Mob Identifier |
| name | varchar | Name |
| class | decimal | [Class](../../../../server/player/class-list) |
| level | int | Level |
| last_login | int | Last Login Unix Timestamp |
| zone_id | decimal | [Zone Identifier](../../../../server/zones/zone-list) |
| deleted_at | datetime | Deleted At |

