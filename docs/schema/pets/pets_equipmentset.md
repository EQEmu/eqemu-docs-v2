# pets_equipmentset

## Relationships

```mermaid
erDiagram
    pets_equipmentset {
        int set_id
    }
    pets {
        int npcID
        varchar type
        int equipmentset
    }
    pets_equipmentset_entries {
        int item_id
        int set_id
    }
    pets_equipmentset ||--o{ pets : "Has-Many"
    pets_equipmentset ||--o{ pets_equipmentset_entries : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | set_id | [pets](../../schema/pets/pets.md) | equipmentset |
| Has-Many | set_id | [pets_equipmentset_entries](../../schema/pets/pets_equipmentset_entries.md) | set_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| set_id | int | Unique Pet Equipment Set Identifier |
| setname | varchar | Pet Equipment Set Name |
| nested_set | int | Nested Set Identifier |

