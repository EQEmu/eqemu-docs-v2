# object_contents

## Relationships

```mermaid
erDiagram
    object_contents {
        intunsigned itemid
        intunsigned parentid
        mediumintunsigned augslot1
        mediumintunsigned augslot2
        mediumintunsigned augslot3
        mediumintunsigned augslot4
        mediumintunsigned augslot5
        mediumint augslot6
        intunsigned zoneid
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
    object {
        varchar content_flags
        varchar content_flags_disabled
        int itemid
        int id
        smallint version
        intunsigned zoneid
    }
    zone {
        int zoneidnumber
        varchar short_name
        tinyintunsigned version
        varchar content_flags
        varchar content_flags_disabled
    }
    object_contents ||--o{ items : "One-to-One"
    object_contents ||--o{ items : "One-to-One"
    object_contents ||--o{ items : "One-to-One"
    object_contents ||--o{ items : "One-to-One"
    object_contents ||--o{ items : "One-to-One"
    object_contents ||--o{ items : "One-to-One"
    object_contents ||--o{ items : "One-to-One"
    object_contents ||--o{ object : "One-to-One"
    object_contents ||--o{ zone : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | augslot1 | [items](../../schema/items/items.md) | id |
| One-to-One | augslot2 | [items](../../schema/items/items.md) | id |
| One-to-One | augslot3 | [items](../../schema/items/items.md) | id |
| One-to-One | augslot4 | [items](../../schema/items/items.md) | id |
| One-to-One | augslot5 | [items](../../schema/items/items.md) | id |
| One-to-One | augslot6 | [items](../../schema/items/items.md) | id |
| One-to-One | itemid | [items](../../schema/items/items.md) | id |
| One-to-One | parentid | [object](../../schema/objects/object.md) | id |
| One-to-One | zoneid | [zone](../../schema/zone/zone.md) | zoneidnumber |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| parentid | int | [Object Identifier](object.md) |
| bagidx | int | Bag Index |
| itemid | int | [Item Identifier](../../schema/items/items.md) |
| charges | smallint | Charges |
| droptime | datetime | Drop Time |
| augslot1 | mediumint | Augment Slot 1 |
| augslot2 | mediumint | Augment Slot 2 |
| augslot3 | mediumint | Augment Slot 3 |
| augslot4 | mediumint | Augment Slot 4 |
| augslot5 | mediumint | Augment Slot 5 |
| augslot6 | mediumint | Augment Slot 6 |

