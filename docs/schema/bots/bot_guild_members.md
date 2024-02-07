# bot_guild_members

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2d1aWxkX21lbWJlcnMge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIGd1aWxkX2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBvd25lcl9pZFxuICAgICAgICB2YXJjaGFyIHNwZWxsc19pZFxuICAgICAgICB2YXJjaGFyIHpvbmVfaWRcbiAgICB9XG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBsZWFkZXJcbiAgICAgICAgdmFyY2hhciBtb3RkX3NldHRlclxuICAgIH1cbiAgICBib3RfZ3VpbGRfbWVtYmVycyB8fC0tb3sgYm90X2RhdGEgOiBcIk9uZS10by1PbmVcIlxuICAgIGJvdF9ndWlsZF9tZW1iZXJzIHx8LS1veyBndWlsZHMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2d1aWxkX21lbWJlcnMge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIGd1aWxkX2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBvd25lcl9pZFxuICAgICAgICB2YXJjaGFyIHNwZWxsc19pZFxuICAgICAgICB2YXJjaGFyIHpvbmVfaWRcbiAgICB9XG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBsZWFkZXJcbiAgICAgICAgdmFyY2hhciBtb3RkX3NldHRlclxuICAgIH1cbiAgICBib3RfZ3VpbGRfbWVtYmVycyB8fC0tb3sgYm90X2RhdGEgOiBcIk9uZS10by1PbmVcIlxuICAgIGJvdF9ndWlsZF9tZW1iZXJzIHx8LS1veyBndWlsZHMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2d1aWxkX21lbWJlcnMge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIGd1aWxkX2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBvd25lcl9pZFxuICAgICAgICB2YXJjaGFyIHNwZWxsc19pZFxuICAgICAgICB2YXJjaGFyIHpvbmVfaWRcbiAgICB9XG4gICAgZ3VpbGRzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBsZWFkZXJcbiAgICAgICAgdmFyY2hhciBtb3RkX3NldHRlclxuICAgIH1cbiAgICBib3RfZ3VpbGRfbWVtYmVycyB8fC0tb3sgYm90X2RhdGEgOiBcIk9uZS10by1PbmVcIlxuICAgIGJvdF9ndWlsZF9tZW1iZXJzIHx8LS1veyBndWlsZHMgOiBcIk9uZS10by1PbmVcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


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

