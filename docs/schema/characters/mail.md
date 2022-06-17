# mail

!!! info
	This page was last generated 2022.06.17

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| msgid | int | Unique Message Identifier |
| charid | int | [Character Identifier](character_data.md) |
| timestamp | int | UNIX Timestamp |
| from | varchar | [From Character Name](character_data.md) |
| subject | varchar | Subject |
| body | text | Body |
| to | text | [To Character Name](character_data.md) |
| status | tinyint | [Status](../../../../server/player/status-levels) |

