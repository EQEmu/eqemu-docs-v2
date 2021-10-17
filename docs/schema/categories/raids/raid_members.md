# raid\_members

| Column | Data Type | Description |
| :--- | :--- | :--- |
| raidid | int | [Raid Identifier](raid_details.md) |
| charid | int | [Character Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/raids/character_data.md) |
| groupid | int | [Group Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/raids/group_id.md) |
| \_class | tinyint | [Class](https://eqemu.gitbook.io/server/categories/player/class-list) |
| level | tinyint | Level |
| name | varchar | Name |
| isgroupleader | tinyint | Is Group Leader: 0 = False, 1 = True |
| israidleader | tinyint | Is Raid Leader: 0 = False, 1 = True |
| islooter | tinyint | Is Looter: 0 = False, 1 = True |

