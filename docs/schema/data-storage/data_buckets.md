# data_buckets

## Relationships

```mermaid
erDiagram
    data_buckets {
        varchar key
    }
    merchantlist {
        varchar content_flags
        varchar content_flags_disabled
        varchar bucket_name
        int item
        varchar merchant_id
        int merchantid
    }
    spell_buckets {
        varchar key
        bigintunsigned spellid
    }
    data_buckets ||--o{ merchantlist : "Has-Many"
    data_buckets ||--o{ spell_buckets : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | key | [merchantlist](../../schema/merchants/merchantlist.md) | bucket_name |
| Has-Many | key | [spell_buckets](../../schema/spells/spell_buckets.md) | key |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | bigint | Unique Data Bucket Identifier |
| key | varchar | Key |
| value | text | Value |
| expires | int | Expiration UNIX Timestamp |
| character_id | bigint |  |
| npc_id | bigint |  |
| bot_id | bigint |  |

