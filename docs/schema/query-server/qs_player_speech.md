# qs_player_speech

!!! info
	This page was last generated 2022.06.17

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Speech Identifier |
| from | varchar | [From Character Identifier](character_data.md) |
| to | varchar | [To Character Identifier](character_data.md) |
| message | varchar | Message |
| minstatus | smallint | [Minimum Status](../../../../server/player/status-levels) |
| guilddbid | int | [Guild Database Identifier](guilds.md) |
| type | tinyint | Type |
| timerecorded | timestamp | Time Recorded Timestamp |

