# guild_ranks

!!! info
	This page was last generated 2023.06.12

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfcmFua3Mge1xuICAgICAgICBtZWRpdW1pbnR1bnNpZ25lZCBndWlsZF9pZFxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGxlYWRlclxuICAgICAgICB2YXJjaGFyIG1vdGRfc2V0dGVyXG4gICAgfVxuICAgIGd1aWxkX3JhbmtzIHx8LS1veyBndWlsZHMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfcmFua3Mge1xuICAgICAgICBtZWRpdW1pbnR1bnNpZ25lZCBndWlsZF9pZFxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGxlYWRlclxuICAgICAgICB2YXJjaGFyIG1vdGRfc2V0dGVyXG4gICAgfVxuICAgIGd1aWxkX3JhbmtzIHx8LS1veyBndWlsZHMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfcmFua3Mge1xuICAgICAgICBtZWRpdW1pbnR1bnNpZ25lZCBndWlsZF9pZFxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGxlYWRlclxuICAgICAgICB2YXJjaGFyIG1vdGRfc2V0dGVyXG4gICAgfVxuICAgIGd1aWxkX3JhbmtzIHx8LS1veyBndWlsZHMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | guild_id | [guilds](../../schema/guilds/guilds.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| guild_id | mediumint | [Guild Identifier](guilds.md) |
| rank | tinyint | Rank Identifier |
| title | varchar | Title |
| can_hear | tinyint | Can Hear: 0 = False, 1 = True |
| can_speak | tinyint | Can Speak: 0 = False, 1 = True |
| can_invite | tinyint | Can Invite: 0 = False, 1 = True |
| can_remove | tinyint | Can Remove: 0 = False, 1 = True |
| can_promote | tinyint | Can Promote: 0 = False, 1 = True |
| can_demote | tinyint | Can Demote: 0 = False, 1 = True |
| can_motd | tinyint | Can MOTD: 0 = False, 1 = True |
| can_warpeace | tinyint | Can War Peace: 0 = False, 1 = True |

