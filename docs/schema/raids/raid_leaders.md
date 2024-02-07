# raid_leaders

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcmFpZF9sZWFkZXJzIHtcbiAgICAgICAgaW50dW5zaWduZWQgcmlkXG4gICAgfVxuICAgIHJhaWRfZGV0YWlscyB7XG4gICAgICAgIGludCByYWlkaWRcbiAgICB9XG4gICAgcmFpZF9sZWFkZXJzIHx8LS1veyByYWlkX2RldGFpbHMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcmFpZF9sZWFkZXJzIHtcbiAgICAgICAgaW50dW5zaWduZWQgcmlkXG4gICAgfVxuICAgIHJhaWRfZGV0YWlscyB7XG4gICAgICAgIGludCByYWlkaWRcbiAgICB9XG4gICAgcmFpZF9sZWFkZXJzIHx8LS1veyByYWlkX2RldGFpbHMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcmFpZF9sZWFkZXJzIHtcbiAgICAgICAgaW50dW5zaWduZWQgcmlkXG4gICAgfVxuICAgIHJhaWRfZGV0YWlscyB7XG4gICAgICAgIGludCByYWlkaWRcbiAgICB9XG4gICAgcmFpZF9sZWFkZXJzIHx8LS1veyByYWlkX2RldGFpbHMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | rid | [raid_details](../../schema/raids/raid_details.md) | raidid |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| gid | int | [Group Identifier](../../schema/groups/group_id.md) |
| rid | int | [Raid Identifier](raid_details.md) |
| marknpc | varchar | Mark NPC: 0 = False, 1 = True |
| maintank | varchar | Main Tank: 0 = False, 1 = True |
| assist | varchar | Assist: 0 = False, 1 = True |
| puller | varchar | Puller: 0 = False, 1 = True |
| leadershipaa | tinyblob | Leadership AA |
| mentoree | varchar | Mentoree: 0 = False, 1 = True |
| mentor_percent | int | Mentore Percent: 0 = None, 100 = Max |

