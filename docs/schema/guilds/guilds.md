# guilds

!!! info
	This page was last generated 2022.05.11

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGd1aWxkX2Jhbmsge1xuICAgICAgICBpbnR1bnNpZ25lZCBndWlsZGlkXG4gICAgfVxuICAgIGd1aWxkX3JhbmtzIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgZ3VpbGRfaWRcbiAgICB9XG4gICAgZ3VpbGRfbWVtYmVycyB7XG4gICAgICAgIGludCBjaGFyX2lkXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGd1aWxkcyB8fC0tb3sgZ3VpbGRfYmFuayA6IEhhcy1NYW55XG4gICAgZ3VpbGRzIHx8LS1veyBndWlsZF9yYW5rcyA6IEhhcy1NYW55XG4gICAgZ3VpbGRzIHx8LS1veyBndWlsZF9tZW1iZXJzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGd1aWxkX2Jhbmsge1xuICAgICAgICBpbnR1bnNpZ25lZCBndWlsZGlkXG4gICAgfVxuICAgIGd1aWxkX3JhbmtzIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgZ3VpbGRfaWRcbiAgICB9XG4gICAgZ3VpbGRfbWVtYmVycyB7XG4gICAgICAgIGludCBjaGFyX2lkXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGd1aWxkcyB8fC0tb3sgZ3VpbGRfYmFuayA6IEhhcy1NYW55XG4gICAgZ3VpbGRzIHx8LS1veyBndWlsZF9yYW5rcyA6IEhhcy1NYW55XG4gICAgZ3VpbGRzIHx8LS1veyBndWlsZF9tZW1iZXJzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGd1aWxkX2Jhbmsge1xuICAgICAgICBpbnR1bnNpZ25lZCBndWlsZGlkXG4gICAgfVxuICAgIGd1aWxkX3JhbmtzIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgZ3VpbGRfaWRcbiAgICB9XG4gICAgZ3VpbGRfbWVtYmVycyB7XG4gICAgICAgIGludCBjaGFyX2lkXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGd1aWxkcyB8fC0tb3sgZ3VpbGRfYmFuayA6IEhhcy1NYW55XG4gICAgZ3VpbGRzIHx8LS1veyBndWlsZF9yYW5rcyA6IEhhcy1NYW55XG4gICAgZ3VpbGRzIHx8LS1veyBndWlsZF9tZW1iZXJzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

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

