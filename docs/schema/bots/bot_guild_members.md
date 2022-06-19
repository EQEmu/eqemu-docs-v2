# bot_guild_members

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2d1aWxkX21lbWJlcnMge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIGd1aWxkX2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciB6b25lX2lkXG4gICAgICAgIHZhcmNoYXIgb3duZXJfaWRcbiAgICB9XG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBsZWFkZXJcbiAgICAgICAgdmFyY2hhciBtb3RkX3NldHRlclxuICAgIH1cbiAgICBib3RfZ3VpbGRfbWVtYmVycyB8fC0tb3sgYm90X2RhdGEgOiBPbmUtdG8tT25lXG4gICAgYm90X2d1aWxkX21lbWJlcnMgfHwtLW97IGd1aWxkcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2d1aWxkX21lbWJlcnMge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIGd1aWxkX2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciB6b25lX2lkXG4gICAgICAgIHZhcmNoYXIgb3duZXJfaWRcbiAgICB9XG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBsZWFkZXJcbiAgICAgICAgdmFyY2hhciBtb3RkX3NldHRlclxuICAgIH1cbiAgICBib3RfZ3VpbGRfbWVtYmVycyB8fC0tb3sgYm90X2RhdGEgOiBPbmUtdG8tT25lXG4gICAgYm90X2d1aWxkX21lbWJlcnMgfHwtLW97IGd1aWxkcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2d1aWxkX21lbWJlcnMge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIGd1aWxkX2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciB6b25lX2lkXG4gICAgICAgIHZhcmNoYXIgb3duZXJfaWRcbiAgICB9XG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBsZWFkZXJcbiAgICAgICAgdmFyY2hhciBtb3RkX3NldHRlclxuICAgIH1cbiAgICBib3RfZ3VpbGRfbWVtYmVycyB8fC0tb3sgYm90X2RhdGEgOiBPbmUtdG8tT25lXG4gICAgYm90X2d1aWxkX21lbWJlcnMgfHwtLW97IGd1aWxkcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_id | [bot_data](../../schema/bots/bot_data.md) | bot_id |
| One-to-One | guild_id | [guilds](../../schema/guilds/guilds.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| bot_id | int | [Bot Identifier](bot_data.md) |
| guild_id | mediumint | [Guild Identifier](../../../schema/guilds/guilds.md) |
| rank | tinyint | [Guild Rank](../../../../categories/player/guild-ranks) |
| tribute_enable | tinyint | Tribute Enabled: 0 = False, 1= True |
| total_tribute | int | Total Tribute |
| last_tribute | int | Last Tribute |
| banker | tinyint | Banker: 0 = False, 1 = True |
| public_note | text | Public Note |
| alt | tinyint | Alt: 0 = False, 1 = True |

