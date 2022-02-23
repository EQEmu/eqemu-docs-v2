# raid_members

!!! info
	This page was last generated 2022.02.23

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| raidid | int | [Raid Identifier](raid_details.md) |
| charid | int | [Character Identifier](../../schema/characters/character_data.md) |
| groupid | int | [Group Identifier](../../schema/groups/group_id.md) |
| _class | tinyint | [Class](../../../../server/player/class-list) |
| level | tinyint | Level |
| name | varchar | Name |
| isgroupleader | tinyint | Is Group Leader: 0 = False, 1 = True |
| israidleader | tinyint | Is Raid Leader: 0 = False, 1 = True |
| islooter | tinyint | Is Looter: 0 = False, 1 = True |

