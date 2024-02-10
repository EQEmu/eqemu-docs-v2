# bot_guild_members

## Relationships

```mermaid
erDiagram
    bot_guild_members {
        varchar bot_id
        varchar guild_id
    }
    bot_data {
        varchar bot_id
        varchar owner_id
        varchar spells_id
        varchar zone_id
    }
    guilds {
        int id
        int leader
        varchar motd_setter
    }
    bot_guild_members ||--o{ bot_data : "One-to-One"
    bot_guild_members ||--o{ guilds : "One-to-One"


```


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

