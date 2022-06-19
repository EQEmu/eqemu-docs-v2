# vw_bot_character_mobs

!!! info
	This page was last generated 2022.06.19

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| mob_type | varchar | Mob Type: B = Bot, C = Client |
| id | int | Unique View Bot Character Mob Identifier |
| name | varchar | Name |
| class | tinyint | [Class](../../../../categories/player/class-list) |
| level | int | Level |
| last_login | int | Last Login Unix Timestamp |
| zone_id | int | [Zone Identifier](../../../../server/zones/zone-list) |
| deleted_at | datetime | Deleted At |

