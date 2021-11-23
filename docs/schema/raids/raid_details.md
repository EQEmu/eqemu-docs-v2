# raid_details

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcmFpZF9kZXRhaWxzIHtcbiAgICAgICAgaW50IHJhaWRpZFxuICAgIH1cbiAgICByYWlkX2xlYWRlcnMge1xuICAgICAgICBpbnR1bnNpZ25lZCByaWRcbiAgICB9XG4gICAgcmFpZF9tZW1iZXJzIHtcbiAgICAgICAgaW50IHJhaWRpZFxuICAgIH1cbiAgICByYWlkX2RldGFpbHMgfHwtLW97IHJhaWRfbGVhZGVycyA6IEhhcy1NYW55XG4gICAgcmFpZF9kZXRhaWxzIHx8LS1veyByYWlkX21lbWJlcnMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcmFpZF9kZXRhaWxzIHtcbiAgICAgICAgaW50IHJhaWRpZFxuICAgIH1cbiAgICByYWlkX2xlYWRlcnMge1xuICAgICAgICBpbnR1bnNpZ25lZCByaWRcbiAgICB9XG4gICAgcmFpZF9tZW1iZXJzIHtcbiAgICAgICAgaW50IHJhaWRpZFxuICAgIH1cbiAgICByYWlkX2RldGFpbHMgfHwtLW97IHJhaWRfbGVhZGVycyA6IEhhcy1NYW55XG4gICAgcmFpZF9kZXRhaWxzIHx8LS1veyByYWlkX21lbWJlcnMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcmFpZF9kZXRhaWxzIHtcbiAgICAgICAgaW50IHJhaWRpZFxuICAgIH1cbiAgICByYWlkX2xlYWRlcnMge1xuICAgICAgICBpbnR1bnNpZ25lZCByaWRcbiAgICB9XG4gICAgcmFpZF9tZW1iZXJzIHtcbiAgICAgICAgaW50IHJhaWRpZFxuICAgIH1cbiAgICByYWlkX2RldGFpbHMgfHwtLW97IHJhaWRfbGVhZGVycyA6IEhhcy1NYW55XG4gICAgcmFpZF9kZXRhaWxzIHx8LS1veyByYWlkX21lbWJlcnMgOiBIYXMtTWFueVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | raidid | [raid_leaders](../../schema/raids/raid_leaders.md) | rid |
| Has-Many | raidid | [raid_members](../../schema/raids/raid_members.md) | raidid |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| raidid | int | Unique Raid Identifier |
| loottype | int | Loot Type |
| locked | tinyint | Locked: 0 = False, 1 = True |
| motd | varchar | Message of the Day |

