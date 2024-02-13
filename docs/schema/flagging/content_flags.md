# content_flags

## Relationships

```mermaid
erDiagram
    content_flags {
        varchar flag_name
    }
    doors {
        varchar content_flags
        varchar content_flags_disabled
        int dz_switch_id
        int keyitem
        varchar zone
        varchar dest_zone
        intunsigned dest_instance
        smallint version
    }
    fishing {
        varchar content_flags
        varchar content_flags_disabled
        int Itemid
        int zoneid
        int npc_id
    }
    forage {
        varchar content_flags
        varchar content_flags_disabled
        int Itemid
        int zoneid
    }
    global_loot {
        varchar content_flags
        varchar content_flags_disabled
        int loottable_id
        mediumtext zone
    }
    ground_spawns {
        varchar content_flags
        varchar content_flags_disabled
        intunsigned item
        smallint version
        intunsigned zoneid
    }
    content_flags ||--o{ doors : "Has-Many"
    content_flags ||--o{ doors : "Has-Many"
    content_flags ||--o{ fishing : "Has-Many"
    content_flags ||--o{ fishing : "Has-Many"
    content_flags ||--o{ forage : "Has-Many"
    content_flags ||--o{ forage : "Has-Many"
    content_flags ||--o{ global_loot : "Has-Many"
    content_flags ||--o{ global_loot : "Has-Many"
    content_flags ||--o{ ground_spawns : "Has-Many"
    content_flags ||--o{ ground_spawns : "Has-Many"


```

```mermaid
erDiagram
    content_flags {
        varchar flag_name
    }
    lootdrop {
        varchar content_flags
        varchar content_flags_disabled
        intunsigned id
    }
    loottable {
        varchar content_flags
        varchar content_flags_disabled
        intunsigned id
    }
    merchantlist {
        varchar content_flags
        varchar content_flags_disabled
        varchar bucket_name
        int item
        varchar merchant_id
        int merchantid
    }
    object {
        varchar content_flags
        varchar content_flags_disabled
        int itemid
        int id
        smallint version
        intunsigned zoneid
    }
    spawn2 {
        varchar content_flags
        varchar content_flags_disabled
        int pathgrid
        int id
        int spawngroupID
        smallint version
        varchar zone
    }
    content_flags ||--o{ lootdrop : "Has-Many"
    content_flags ||--o{ lootdrop : "Has-Many"
    content_flags ||--o{ loottable : "Has-Many"
    content_flags ||--o{ loottable : "Has-Many"
    content_flags ||--o{ merchantlist : "Has-Many"
    content_flags ||--o{ merchantlist : "Has-Many"
    content_flags ||--o{ object : "Has-Many"
    content_flags ||--o{ object : "Has-Many"
    content_flags ||--o{ spawn2 : "Has-Many"
    content_flags ||--o{ spawn2 : "Has-Many"


```

```mermaid
erDiagram
    content_flags {
        varchar flag_name
    }
    spawnentry {
        varchar content_flags
        varchar content_flags_disabled
        int npcID
        int spawngroupID
    }
    start_zones {
        varchar content_flags
        varchar content_flags_disabled
        int start_zone
        int zone_id
    }
    starting_items {
        varchar content_flags
        varchar content_flags_disabled
        varchar itemid
        varchar zone_id
        varchar zoneid
    }
    tradeskill_recipe {
        int id
        varchar content_flags
        varchar content_flags_disabled
    }
    traps {
        varchar content_flags
        varchar content_flags_disabled
        smallintunsigned version
        varchar zone
    }
    content_flags ||--o{ spawnentry : "Has-Many"
    content_flags ||--o{ spawnentry : "Has-Many"
    content_flags ||--o{ start_zones : "Has-Many"
    content_flags ||--o{ start_zones : "Has-Many"
    content_flags ||--o{ starting_items : "Has-Many"
    content_flags ||--o{ starting_items : "Has-Many"
    content_flags ||--o{ tradeskill_recipe : "Has-Many"
    content_flags ||--o{ tradeskill_recipe : "Has-Many"
    content_flags ||--o{ traps : "Has-Many"
    content_flags ||--o{ traps : "Has-Many"


```

```mermaid
erDiagram
    content_flags {
        varchar flag_name
    }
    zone {
        int zoneidnumber
        varchar short_name
        tinyintunsigned version
        varchar content_flags
        varchar content_flags_disabled
    }
    zone_points {
        varchar content_flags
        varchar content_flags_disabled
        intunsigned target_instance
        varchar zone
        intunsigned target_zone_id
        int version
    }
    content_flags ||--o{ zone : "Has-Many"
    content_flags ||--o{ zone : "Has-Many"
    content_flags ||--o{ zone_points : "Has-Many"
    content_flags ||--o{ zone_points : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | flag_name | [doors](../../schema/doors/doors.md) | content_flags |
| Has-Many | flag_name | [doors](../../schema/doors/doors.md) | content_flags_disabled |
| Has-Many | flag_name | [fishing](../../schema/tradeskills/fishing.md) | content_flags |
| Has-Many | flag_name | [fishing](../../schema/tradeskills/fishing.md) | content_flags_disabled |
| Has-Many | flag_name | [forage](../../schema/tradeskills/forage.md) | content_flags |
| Has-Many | flag_name | [forage](../../schema/tradeskills/forage.md) | content_flags_disabled |
| Has-Many | flag_name | [global_loot](../../schema/loot/global_loot.md) | content_flags |
| Has-Many | flag_name | [global_loot](../../schema/loot/global_loot.md) | content_flags_disabled |
| Has-Many | flag_name | [ground_spawns](../../schema/ground-spawns/ground_spawns.md) | content_flags |
| Has-Many | flag_name | [ground_spawns](../../schema/ground-spawns/ground_spawns.md) | content_flags_disabled |
| Has-Many | flag_name | [lootdrop](../../schema/loot/lootdrop.md) | content_flags |
| Has-Many | flag_name | [lootdrop](../../schema/loot/lootdrop.md) | content_flags_disabled |
| Has-Many | flag_name | [loottable](../../schema/loot/loottable.md) | content_flags |
| Has-Many | flag_name | [loottable](../../schema/loot/loottable.md) | content_flags_disabled |
| Has-Many | flag_name | [merchantlist](../../schema/merchants/merchantlist.md) | content_flags |
| Has-Many | flag_name | [merchantlist](../../schema/merchants/merchantlist.md) | content_flags_disabled |
| Has-Many | flag_name | [object](../../schema/objects/object.md) | content_flags |
| Has-Many | flag_name | [object](../../schema/objects/object.md) | content_flags_disabled |
| Has-Many | flag_name | [spawn2](../../schema/spawns/spawn2.md) | content_flags |
| Has-Many | flag_name | [spawn2](../../schema/spawns/spawn2.md) | content_flags_disabled |
| Has-Many | flag_name | [spawnentry](../../schema/spawns/spawnentry.md) | content_flags |
| Has-Many | flag_name | [spawnentry](../../schema/spawns/spawnentry.md) | content_flags_disabled |
| Has-Many | flag_name | [start_zones](../../schema/admin/start_zones.md) | content_flags |
| Has-Many | flag_name | [start_zones](../../schema/admin/start_zones.md) | content_flags_disabled |
| Has-Many | flag_name | [starting_items](../../schema/admin/starting_items.md) | content_flags |
| Has-Many | flag_name | [starting_items](../../schema/admin/starting_items.md) | content_flags_disabled |
| Has-Many | flag_name | [tradeskill_recipe](../../schema/tradeskills/tradeskill_recipe.md) | content_flags |
| Has-Many | flag_name | [tradeskill_recipe](../../schema/tradeskills/tradeskill_recipe.md) | content_flags_disabled |
| Has-Many | flag_name | [traps](../../schema/traps/traps.md) | content_flags |
| Has-Many | flag_name | [traps](../../schema/traps/traps.md) | content_flags_disabled |
| Has-Many | flag_name | [zone](../../schema/zone/zone.md) | content_flags |
| Has-Many | flag_name | [zone](../../schema/zone/zone.md) | content_flags_disabled |
| Has-Many | flag_name | [zone_points](../../schema/zone/zone_points.md) | content_flags |
| Has-Many | flag_name | [zone_points](../../schema/zone/zone_points.md) | content_flags_disabled |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Flag Identifier |
| flag_name | varchar | Flag Name |
| enabled | tinyint | Enabled: 0 = False, 1 = True |
| notes | text | Notes |

