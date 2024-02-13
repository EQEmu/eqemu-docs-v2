# guild_relations

## Relationships

```mermaid
erDiagram
    guild_relations {
        mediumintunsigned guild1
        mediumintunsigned guild2
    }
    guilds {
        int id
        int leader
        varchar motd_setter
    }
    guild_relations ||--o{ guilds : "One-to-One"
    guild_relations ||--o{ guilds : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | guild1 | [guilds](../../schema/guilds/guilds.md) | id |
| One-to-One | guild2 | [guilds](../../schema/guilds/guilds.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| guild1 | mediumint | [Unique Guild Identifier 1](guilds.md) |
| guild2 | mediumint | [Unique Guild Identifier 2](guilds.md) |
| relation | tinyint | Relation |

