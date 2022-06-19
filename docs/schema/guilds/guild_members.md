# guild_members

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfbWVtYmVycyB7XG4gICAgICAgIGludCBjaGFyX2lkXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGxlYWRlclxuICAgICAgICB2YXJjaGFyIG1vdGRfc2V0dGVyXG4gICAgfVxuICAgIGd1aWxkX21lbWJlcnMgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuICAgIGd1aWxkX21lbWJlcnMgfHwtLW97IGd1aWxkcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfbWVtYmVycyB7XG4gICAgICAgIGludCBjaGFyX2lkXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGxlYWRlclxuICAgICAgICB2YXJjaGFyIG1vdGRfc2V0dGVyXG4gICAgfVxuICAgIGd1aWxkX21lbWJlcnMgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuICAgIGd1aWxkX21lbWJlcnMgfHwtLW97IGd1aWxkcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfbWVtYmVycyB7XG4gICAgICAgIGludCBjaGFyX2lkXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGxlYWRlclxuICAgICAgICB2YXJjaGFyIG1vdGRfc2V0dGVyXG4gICAgfVxuICAgIGd1aWxkX21lbWJlcnMgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuICAgIGd1aWxkX21lbWJlcnMgfHwtLW97IGd1aWxkcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


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

