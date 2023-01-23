# guild_members

!!! info
	This page was last generated 2023.01.22

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfbWVtYmVycyB7XG4gICAgICAgIGludCBjaGFyX2lkXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICB9XG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBsZWFkZXJcbiAgICAgICAgdmFyY2hhciBtb3RkX3NldHRlclxuICAgIH1cbiAgICBndWlsZF9tZW1iZXJzIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcbiAgICBndWlsZF9tZW1iZXJzIHx8LS1veyBndWlsZHMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfbWVtYmVycyB7XG4gICAgICAgIGludCBjaGFyX2lkXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICB9XG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBsZWFkZXJcbiAgICAgICAgdmFyY2hhciBtb3RkX3NldHRlclxuICAgIH1cbiAgICBndWlsZF9tZW1iZXJzIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcbiAgICBndWlsZF9tZW1iZXJzIHx8LS1veyBndWlsZHMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfbWVtYmVycyB7XG4gICAgICAgIGludCBjaGFyX2lkXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICB9XG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBsZWFkZXJcbiAgICAgICAgdmFyY2hhciBtb3RkX3NldHRlclxuICAgIH1cbiAgICBndWlsZF9tZW1iZXJzIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcbiAgICBndWlsZF9tZW1iZXJzIHx8LS1veyBndWlsZHMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | guild_id | [guilds](../../schema/guilds/guilds.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| char_id | int | [Character Identifier](../../schema/characters/character_data.md) |
| guild_id | mediumint | [Guild Identifier](guilds.md) |
| rank | tinyint | [Rank](../../../../server/player/guild-ranks) |
| tribute_enable | tinyint | Tribute Enable: 0 = False, 1 = True |
| total_tribute | int | Total Tribute |
| last_tribute | int | Last Tribute |
| banker | tinyint | Banked: 0 = False, 1 = True |
| public_note | text | Public Note |
| alt | tinyint | Alt: 0 = False, 1 = True |

