# guild_bank

## Relationships

```mermaid
erDiagram
    guild_bank {
        varchar donator
        intunsigned guildid
        intunsigned itemid
    }
    character_data {
        intunsigned id
        varchar name
        varchar nane
        intunsigned zone_instance
        intunsigned zone_id
    }
    guilds {
        int id
        int leader
        varchar motd_setter
    }
    items {
        int id
        int book
        varchar name
        int recasttype
        int icon
        mediumint bardeffect
        int clickeffect
        int focuseffect
        int proceffect
        int scrolleffect
        int worneffect
    }
    guild_bank ||--o{ character_data : "One-to-One"
    guild_bank ||--o{ guilds : "One-to-One"
    guild_bank ||--o{ items : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | donator | [character_data](../../schema/characters/character_data.md) | name |
| One-to-One | guildid | [guilds](../../schema/guilds/guilds.md) | id |
| One-to-One | itemid | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| guildid | int | [Guild Identifier](guilds.md) |
| area | tinyint | Area |
| slot | int | Slot Identifier |
| itemid | int | [Item Identifier](../../schema/items/items.md) |
| qty | int | Quantity |
| donator | varchar | [Character Identifier](../../schema/characters/character_data.md) |
| permissions | tinyint | Permissions |
| whofor | varchar | Who For |

