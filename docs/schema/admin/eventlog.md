# eventlog

!!! info
	This page was last generated 2022.06.17

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Event Identifier |
| accountname | varchar | [Account Name](account.md) |
| accountid | int | [Account Identifier](account.md) |
| status | int | [Status](../../../../server/player/status-levels) |
| charname | varchar | [Character Name](character_data.md) |
| target | varchar | Target |
| time | timestamp | TIme Timestamp |
| descriptiontype | varchar | Description Type |
| description | text | Description |
| event_nid | int | Event Identifier |

