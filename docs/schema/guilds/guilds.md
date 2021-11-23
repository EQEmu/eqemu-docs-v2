# guilds

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBndWlsZF9iYW5rIHtcbiAgICAgICAgaW50dW5zaWduZWQgZ3VpbGRpZFxuICAgIH1cbiAgICBndWlsZF9yYW5rcyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGd1aWxkX21lbWJlcnMge1xuICAgICAgICBpbnQgY2hhcl9pZFxuICAgICAgICBtZWRpdW1pbnR1bnNpZ25lZCBndWlsZF9pZFxuICAgIH1cbiAgICBndWlsZHMgfHwtLW97IGd1aWxkX2JhbmsgOiBIYXMtTWFueVxuICAgIGd1aWxkcyB8fC0tb3sgZ3VpbGRfcmFua3MgOiBIYXMtTWFueVxuICAgIGd1aWxkcyB8fC0tb3sgZ3VpbGRfbWVtYmVycyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBndWlsZF9iYW5rIHtcbiAgICAgICAgaW50dW5zaWduZWQgZ3VpbGRpZFxuICAgIH1cbiAgICBndWlsZF9yYW5rcyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGd1aWxkX21lbWJlcnMge1xuICAgICAgICBpbnQgY2hhcl9pZFxuICAgICAgICBtZWRpdW1pbnR1bnNpZ25lZCBndWlsZF9pZFxuICAgIH1cbiAgICBndWlsZHMgfHwtLW97IGd1aWxkX2JhbmsgOiBIYXMtTWFueVxuICAgIGd1aWxkcyB8fC0tb3sgZ3VpbGRfcmFua3MgOiBIYXMtTWFueVxuICAgIGd1aWxkcyB8fC0tb3sgZ3VpbGRfbWVtYmVycyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBndWlsZF9iYW5rIHtcbiAgICAgICAgaW50dW5zaWduZWQgZ3VpbGRpZFxuICAgIH1cbiAgICBndWlsZF9yYW5rcyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGd1aWxkX21lbWJlcnMge1xuICAgICAgICBpbnQgY2hhcl9pZFxuICAgICAgICBtZWRpdW1pbnR1bnNpZ25lZCBndWlsZF9pZFxuICAgIH1cbiAgICBndWlsZHMgfHwtLW97IGd1aWxkX2JhbmsgOiBIYXMtTWFueVxuICAgIGd1aWxkcyB8fC0tb3sgZ3VpbGRfcmFua3MgOiBIYXMtTWFueVxuICAgIGd1aWxkcyB8fC0tb3sgZ3VpbGRfbWVtYmVycyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [guild_bank](../../schema/guilds/guild_bank.md) | guildid |
| Has-Many | id | [guild_ranks](../../schema/guilds/guild_ranks.md) | guild_id |
| Has-Many | id | [guild_members](../../schema/guilds/guild_members.md) | guild_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Guild Identifier |
| name | varchar | Name |
| leader | int | [Character Identifier](../../schema/characters/character_data.md) |
| minstatus | smallint | [Minimum Status](../../../../server/player/status-levels) |
| motd | text | Message of the Day |
| tribute | int | Tribute |
| motd_setter | varchar | [Character Identifier](../../schema/characters/character_data.md) |
| channel | varchar | Channel |
| url | varchar | Website URL |

