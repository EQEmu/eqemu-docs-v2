# qs\_player\_speech

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Speech Identifier |
| from | varchar | [From Character Identifier](https://github.com/EQEmu/docs-db-schema/tree/774e95edd473c84dafd6fe13b9b699f6b84a7ce8/docs/schema/categories/query_server/character_data.md) |
| to | varchar | [To Character Identifier](https://github.com/EQEmu/docs-db-schema/tree/774e95edd473c84dafd6fe13b9b699f6b84a7ce8/docs/schema/categories/query_server/character_data.md) |
| message | varchar | Message |
| minstatus | smallint | [Minimum Status](https://eqemu.gitbook.io/server/categories/player/status-levels) |
| guilddbid | int | [Guild Database Identifier](https://github.com/EQEmu/docs-db-schema/tree/774e95edd473c84dafd6fe13b9b699f6b84a7ce8/docs/schema/categories/query_server/guilds.md) |
| type | tinyint | Type |
| timerecorded | timestamp | Time Recorded Timestamp |

