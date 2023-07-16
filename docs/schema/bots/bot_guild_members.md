# bot_guild_members

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2d1aWxkX21lbWJlcnMge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIGd1aWxkX2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgc21hbGxpbnQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG93bmVyX2lkXG4gICAgfVxuICAgIGd1aWxkcyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgbGVhZGVyXG4gICAgICAgIHZhcmNoYXIgbW90ZF9zZXR0ZXJcbiAgICB9XG4gICAgYm90X2d1aWxkX21lbWJlcnMgfHwtLW97IGJvdF9kYXRhIDogXCJPbmUtdG8tT25lXCJcbiAgICBib3RfZ3VpbGRfbWVtYmVycyB8fC0tb3sgZ3VpbGRzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2d1aWxkX21lbWJlcnMge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIGd1aWxkX2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgc21hbGxpbnQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG93bmVyX2lkXG4gICAgfVxuICAgIGd1aWxkcyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgbGVhZGVyXG4gICAgICAgIHZhcmNoYXIgbW90ZF9zZXR0ZXJcbiAgICB9XG4gICAgYm90X2d1aWxkX21lbWJlcnMgfHwtLW97IGJvdF9kYXRhIDogXCJPbmUtdG8tT25lXCJcbiAgICBib3RfZ3VpbGRfbWVtYmVycyB8fC0tb3sgZ3VpbGRzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2d1aWxkX21lbWJlcnMge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIGd1aWxkX2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgc21hbGxpbnQgem9uZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBib3RfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG93bmVyX2lkXG4gICAgfVxuICAgIGd1aWxkcyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgbGVhZGVyXG4gICAgICAgIHZhcmNoYXIgbW90ZF9zZXR0ZXJcbiAgICB9XG4gICAgYm90X2d1aWxkX21lbWJlcnMgfHwtLW97IGJvdF9kYXRhIDogXCJPbmUtdG8tT25lXCJcbiAgICBib3RfZ3VpbGRfbWVtYmVycyB8fC0tb3sgZ3VpbGRzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_id | [bot_data](../../schema/bots/bot_data.md) | bot_id |
| One-to-One | guild_id | [guilds](../../schema/guilds/guilds.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| bot_id | int | [Bot Identifier](bot_data.md) |
| guild_id | mediumint | [Guild Identifier](../../schema/guilds/guilds.md) |
| rank | tinyint | [Guild Rank](../../../../categories/player/guild-ranks) |
| tribute_enable | tinyint | Tribute Enabled: 0 = False, 1= True |
| total_tribute | int | Total Tribute |
| last_tribute | int | Last Tribute |
| banker | tinyint | Banker: 0 = False, 1 = True |
| public_note | text | Public Note |
| alt | tinyint | Alt: 0 = False, 1 = True |

