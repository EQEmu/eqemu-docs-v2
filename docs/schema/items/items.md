# items

## Relationships

```mermaid
erDiagram
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
    spells_new {
        int id
        int descnum
        int effectdescnum
        int effectdescnum2
        int typedescnum
        varchar teleport_zone
    }
    books {
        int id
    }
    alternate_currency {
        int id
        int item_id
    }
    bot_inventories {
        varchar bot_id
        varchar item_id
        varchar augment_1
        varchar augment_2
        varchar augment_3
        varchar augment_4
        varchar augment_5
        varchar augment_6
    }
    bot_pet_inventories {
        varchar pets_index
        varchar item_id
    }
    items ||--o{ spells_new : "One-to-One"
    items ||--o{ books : "One-to-One"
    items ||--o{ spells_new : "One-to-One"
    items ||--o{ spells_new : "One-to-One"
    items ||--o{ alternate_currency : "Has-Many"
    items ||--o{ bot_inventories : "Has-Many"
    items ||--o{ bot_inventories : "Has-Many"
    items ||--o{ bot_inventories : "Has-Many"
    items ||--o{ bot_inventories : "Has-Many"
    items ||--o{ bot_inventories : "Has-Many"
    items ||--o{ bot_inventories : "Has-Many"
    items ||--o{ bot_inventories : "Has-Many"
    items ||--o{ bot_pet_inventories : "Has-Many"
    items ||--o{ spells_new : "One-to-One"
    items ||--o{ spells_new : "One-to-One"
    items ||--o{ spells_new : "One-to-One"


```

```mermaid
erDiagram
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
    character_bandolier {
        intunsigned id
        intunsigned item_id
    }
    character_corpse_items {
        intunsigned corpse_id
        intunsigned item_id
        intunsigned aug_1
        intunsigned aug_2
        intunsigned aug_3
        intunsigned aug_4
        intunsigned aug_5
        int aug_6
    }
    character_pet_inventory {
        int char_id
        int item_id
    }
    character_potionbelt {
        intunsigned id
        intunsigned icon
        intunsigned item_id
    }
    discovered_items {
        varchar char_name
        intunsigned item_id
    }
    items ||--o{ character_bandolier : "Has-Many"
    items ||--o{ character_corpse_items : "Has-Many"
    items ||--o{ character_pet_inventory : "Has-Many"
    items ||--o{ character_potionbelt : "Has-Many"
    items ||--o{ discovered_items : "Has-Many"


```

```mermaid
erDiagram
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
    ground_spawns {
        varchar content_flags
        varchar content_flags_disabled
        intunsigned item
        smallint version
        intunsigned zoneid
    }
    inventory {
        intunsigned charid
        intunsigned itemid
        mediumintunsigned augslot1
        mediumintunsigned augslot2
        mediumintunsigned augslot3
        mediumintunsigned augslot4
        mediumintunsigned augslot5
        mediumint augslot6
    }
    items ||--o{ doors : "Has-Many"
    items ||--o{ fishing : "Has-Many"
    items ||--o{ forage : "Has-Many"
    items ||--o{ ground_spawns : "Has-Many"
    items ||--o{ inventory : "Has-Many"


```

```mermaid
erDiagram
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
    item_tick {
        varchar it_itemid
        varchar it_qglobal
    }
    keyring {
        int char_id
        int item_id
    }
    lootdrop_entries {
        int item_id
        intunsigned lootdrop_id
    }
    merc_inventory {
        varchar item_id
        varchar merc_subtype_id
    }
    merchantlist {
        varchar content_flags
        varchar content_flags_disabled
        varchar bucket_name
        int item
        varchar merchant_id
        int merchantid
    }
    items ||--o{ item_tick : "Has-Many"
    items ||--o{ keyring : "Has-Many"
    items ||--o{ lootdrop_entries : "Has-Many"
    items ||--o{ merc_inventory : "Has-Many"
    items ||--o{ merchantlist : "Has-Many"


```

```mermaid
erDiagram
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
    merchantlist_temp {
        intunsigned itemid
        intunsigned npcid
    }
    object {
        varchar content_flags
        varchar content_flags_disabled
        int itemid
        int id
        smallint version
        intunsigned zoneid
    }
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
    pets_equipmentset_entries {
        int item_id
        int set_id
    }
    qs_merchant_transaction_record_entries {
        int item_id
        int aug_1
        int aug_2
        int aug_3
        int aug_4
        int aug_5
    }
    items ||--o{ merchantlist_temp : "Has-Many"
    items ||--o{ object : "Has-Many"
    items ||--o{ object_contents : "Has-Many"
    items ||--o{ pets_equipmentset_entries : "Has-Many"
    items ||--o{ qs_merchant_transaction_record_entries : "Has-Many"


```

```mermaid
erDiagram
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
    qs_player_delete_record_entries {
        int item_id
        int aug_1
        int aug_2
        int aug_3
        int aug_4
        int aug_5
    }
    qs_player_handin_record_entries {
        int item_id
        int aug_1
        int aug_2
        int aug_3
        int aug_4
        int aug_5
    }
    qs_player_move_record_entries {
        int item_id
        int aug_1
        int aug_2
        int aug_3
        int aug_4
        int aug_5
    }
    qs_player_trade_record_entries {
        int from_id
        int to_id
        int item_id
        int aug_1
        int aug_2
        int aug_3
        int aug_4
        int aug_5
    }
    starting_items {
        varchar content_flags
        varchar content_flags_disabled
        varchar itemid
        varchar zone_id
        varchar zoneid
    }
    items ||--o{ qs_player_delete_record_entries : "Has-Many"
    items ||--o{ qs_player_delete_record_entries : "Has-Many"
    items ||--o{ qs_player_delete_record_entries : "Has-Many"
    items ||--o{ qs_player_delete_record_entries : "Has-Many"
    items ||--o{ qs_player_delete_record_entries : "Has-Many"
    items ||--o{ qs_player_delete_record_entries : "Has-Many"
    items ||--o{ qs_player_handin_record_entries : "Has-Many"
    items ||--o{ qs_player_handin_record_entries : "Has-Many"
    items ||--o{ qs_player_handin_record_entries : "Has-Many"
    items ||--o{ qs_player_handin_record_entries : "Has-Many"
    items ||--o{ qs_player_handin_record_entries : "Has-Many"
    items ||--o{ qs_player_handin_record_entries : "Has-Many"
    items ||--o{ qs_player_move_record_entries : "Has-Many"
    items ||--o{ qs_player_move_record_entries : "Has-Many"
    items ||--o{ qs_player_move_record_entries : "Has-Many"
    items ||--o{ qs_player_move_record_entries : "Has-Many"
    items ||--o{ qs_player_move_record_entries : "Has-Many"
    items ||--o{ qs_player_move_record_entries : "Has-Many"
    items ||--o{ qs_player_trade_record_entries : "Has-Many"
    items ||--o{ qs_player_trade_record_entries : "Has-Many"
    items ||--o{ qs_player_trade_record_entries : "Has-Many"
    items ||--o{ qs_player_trade_record_entries : "Has-Many"
    items ||--o{ qs_player_trade_record_entries : "Has-Many"
    items ||--o{ qs_player_trade_record_entries : "Has-Many"
    items ||--o{ starting_items : "Has-Many"


```

```mermaid
erDiagram
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
    tool_gearup_armor_sets {
        int item_id
    }
    trader {
        intunsigned char_id
        intunsigned item_id
    }
    tradeskill_recipe_entries {
        int item_id
        int recipe_id
    }
    tribute_levels {
        intunsigned item_id
        intunsigned tribute_id
    }
    veteran_reward_templates {
        intunsigned claim_id
        intunsigned item_id
    }
    items ||--o{ tool_gearup_armor_sets : "Has-Many"
    items ||--o{ trader : "Has-Many"
    items ||--o{ tradeskill_recipe_entries : "Has-Many"
    items ||--o{ tribute_levels : "Has-Many"
    items ||--o{ veteran_reward_templates : "Has-Many"


```

```mermaid
erDiagram
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
    trader_audit {
        varchar buyer
        varchar seller
        varchar itemname
    }
    character_item_recast {
        intunsigned id
        intunsigned recast_type
    }
    items ||--o{ trader_audit : "Has-Many"
    items ||--o{ character_item_recast : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bardeffect | [spells_new](../../schema/spells/spells_new.md) | id |
| One-to-One | book | [books](../../schema/books/books.md) | id |
| One-to-One | clickeffect | [spells_new](../../schema/spells/spells_new.md) | id |
| One-to-One | focuseffect | [spells_new](../../schema/spells/spells_new.md) | id |
| Has-Many | id | [alternate_currency](../../schema/alternate-currency/alternate_currency.md) | item_id |
| Has-Many | id | [bot_inventories](../../schema/bots/bot_inventories.md) | item_id |
| Has-Many | id | [bot_inventories](../../schema/bots/bot_inventories.md) | augment_1 |
| Has-Many | id | [bot_inventories](../../schema/bots/bot_inventories.md) | augment_2 |
| Has-Many | id | [bot_inventories](../../schema/bots/bot_inventories.md) | augment_3 |
| Has-Many | id | [bot_inventories](../../schema/bots/bot_inventories.md) | augment_4 |
| Has-Many | id | [bot_inventories](../../schema/bots/bot_inventories.md) | augment_5 |
| Has-Many | id | [bot_inventories](../../schema/bots/bot_inventories.md) | augment_6 |
| Has-Many | id | [bot_pet_inventories](../../schema/bots/bot_pet_inventories.md) | item_id |
| Has-Many | id | [character_bandolier](../../schema/characters/character_bandolier.md) | item_id |
| Has-Many | id | [character_corpse_items](../../schema/characters/character_corpse_items.md) | item_id |
| Has-Many | id | [character_pet_inventory](../../schema/characters/character_pet_inventory.md) | item_id |
| Has-Many | id | [character_potionbelt](../../schema/characters/character_potionbelt.md) | item_id |
| Has-Many | id | [discovered_items](../../schema/admin/discovered_items.md) | item_id |
| Has-Many | id | [doors](../../schema/doors/doors.md) | keyitem |
| Has-Many | id | [fishing](../../schema/tradeskills/fishing.md) | Itemid |
| Has-Many | id | [forage](../../schema/tradeskills/forage.md) | Itemid |
| Has-Many | id | [ground_spawns](../../schema/ground-spawns/ground_spawns.md) | item |
| Has-Many | id | [inventory](../../schema/mercenaries/merc_inventory.md) | itemid |
| Has-Many | id | [item_tick](../../schema/items/item_tick.md) | it_itemid |
| Has-Many | id | [keyring](../../schema/characters/keyring.md) | item_id |
| Has-Many | id | [lootdrop_entries](../../schema/loot/lootdrop_entries.md) | item_id |
| Has-Many | id | [merc_inventory](../../schema/mercenaries/merc_inventory.md) | item_id |
| Has-Many | id | [merchantlist](../../schema/merchants/merchantlist.md) | item |
| Has-Many | id | [merchantlist_temp](../../schema/merchants/merchantlist_temp.md) | itemid |
| Has-Many | id | [object](../../schema/objects/object.md) | itemid |
| Has-Many | id | [object_contents](../../schema/objects/object_contents.md) | itemid |
| Has-Many | id | [pets_equipmentset_entries](../../schema/pets/pets_equipmentset_entries.md) | item_id |
| Has-Many | id | [qs_merchant_transaction_record_entries](../../schema/query-server/qs_merchant_transaction_record_entries.md) | item_id |
| Has-Many | id | [qs_player_delete_record_entries](../../schema/query-server/qs_player_delete_record_entries.md) | item_id |
| Has-Many | id | [qs_player_delete_record_entries](../../schema/query-server/qs_player_delete_record_entries.md) | aug_1 |
| Has-Many | id | [qs_player_delete_record_entries](../../schema/query-server/qs_player_delete_record_entries.md) | aug_2 |
| Has-Many | id | [qs_player_delete_record_entries](../../schema/query-server/qs_player_delete_record_entries.md) | aug_3 |
| Has-Many | id | [qs_player_delete_record_entries](../../schema/query-server/qs_player_delete_record_entries.md) | aug_4 |
| Has-Many | id | [qs_player_delete_record_entries](../../schema/query-server/qs_player_delete_record_entries.md) | aug_5 |
| Has-Many | id | [qs_player_handin_record_entries](../../schema/query-server/qs_player_handin_record_entries.md) | item_id |
| Has-Many | id | [qs_player_handin_record_entries](../../schema/query-server/qs_player_handin_record_entries.md) | aug_1 |
| Has-Many | id | [qs_player_handin_record_entries](../../schema/query-server/qs_player_handin_record_entries.md) | aug_2 |
| Has-Many | id | [qs_player_handin_record_entries](../../schema/query-server/qs_player_handin_record_entries.md) | aug_3 |
| Has-Many | id | [qs_player_handin_record_entries](../../schema/query-server/qs_player_handin_record_entries.md) | aug_4 |
| Has-Many | id | [qs_player_handin_record_entries](../../schema/query-server/qs_player_handin_record_entries.md) | aug_5 |
| Has-Many | id | [qs_player_move_record_entries](../../schema/query-server/qs_player_move_record_entries.md) | item_id |
| Has-Many | id | [qs_player_move_record_entries](../../schema/query-server/qs_player_move_record_entries.md) | aug_1 |
| Has-Many | id | [qs_player_move_record_entries](../../schema/query-server/qs_player_move_record_entries.md) | aug_2 |
| Has-Many | id | [qs_player_move_record_entries](../../schema/query-server/qs_player_move_record_entries.md) | aug_3 |
| Has-Many | id | [qs_player_move_record_entries](../../schema/query-server/qs_player_move_record_entries.md) | aug_4 |
| Has-Many | id | [qs_player_move_record_entries](../../schema/query-server/qs_player_move_record_entries.md) | aug_5 |
| Has-Many | id | [qs_player_trade_record_entries](../../schema/query-server/qs_player_trade_record_entries.md) | item_id |
| Has-Many | id | [qs_player_trade_record_entries](../../schema/query-server/qs_player_trade_record_entries.md) | aug_1 |
| Has-Many | id | [qs_player_trade_record_entries](../../schema/query-server/qs_player_trade_record_entries.md) | aug_2 |
| Has-Many | id | [qs_player_trade_record_entries](../../schema/query-server/qs_player_trade_record_entries.md) | aug_3 |
| Has-Many | id | [qs_player_trade_record_entries](../../schema/query-server/qs_player_trade_record_entries.md) | aug_4 |
| Has-Many | id | [qs_player_trade_record_entries](../../schema/query-server/qs_player_trade_record_entries.md) | aug_5 |
| Has-Many | id | [starting_items](../../schema/admin/starting_items.md) | itemid |
| Has-Many | id | [tool_gearup_armor_sets](../../schema/tools/tool_gearup_armor_sets.md) | item_id |
| Has-Many | id | [trader](../../schema/trader/trader.md) | item_id |
| Has-Many | id | [tradeskill_recipe_entries](../../schema/tradeskills/tradeskill_recipe_entries.md) | item_id |
| Has-Many | id | [tribute_levels](../../schema/tributes/tribute_levels.md) | item_id |
| Has-Many | id | [veteran_reward_templates](../../schema/admin/veteran_reward_templates.md) | item_id |
| Has-Many | name | [trader_audit](../../schema/trader/trader_audit.md) | itemname |
| One-to-One | proceffect | [spells_new](../../schema/spells/spells_new.md) | id |
| Has-Many | recasttype | [character_item_recast](../../schema/characters/character_item_recast.md) | recast_type |
| One-to-One | scrolleffect | [spells_new](../../schema/spells/spells_new.md) | id |
| One-to-One | worneffect | [spells_new](../../schema/spells/spells_new.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Item Identifier |
| minstatus | smallint | [Minimum Status](../../../../server/player/status-levels) |
| Name | varchar | Name |
| aagi | int | Agility: -128 to 127 |
| ac | int | Armor Class |
| accuracy | int | Accuracy |
| acha | int | Charisma: -128 to 127 |
| adex | int | Dexterity: -128 to 127 |
| aint | int | Intelligence: -128 to 127 |
| artifactflag | tinyint | Artifact: 0 = False, 1 = True |
| asta | int | Stamina: -128 to 127 |
| astr | int | Strenght: -128 to 127 |
| attack | int | Attack |
| augrestrict | int | [Augment Restriction](../../../../server/items/augment-restrictions) |
| augslot1type | tinyint | [Augment Slot 1 Type](../../../../server/items/augment-types) |
| augslot1visible | tinyint | Augment Slot 1 Visible: 0 = False, 1 = True |
| augslot2type | tinyint | [Augment Slot 2 Type](../../../../server/items/augment-types) |
| augslot2visible | tinyint | Augment Slot 2 Visible: 0 = False, 1 = True |
| augslot3type | tinyint | [Augment Slot 3 Type](../../../../server/items/augment-types) |
| augslot3visible | tinyint | Augment Slot 3 Visible: 0 = False, 1 = True |
| augslot4type | tinyint | [Augment Slot 4 Type](../../../../server/items/augment-types) |
| augslot4visible | tinyint | Augment Slot 4 Visible: 0 = False, 1 = True |
| augslot5type | tinyint | [Augment Slot 5 Type](../../../../server/items/augment-types) |
| augslot5visible | tinyint | Augment Slot 5 Visible: 0 = False, 1 = True |
| augslot6type | tinyint | [Augment Slot 6 Type](../../../../server/items/augment-types) |
| augslot6visible | tinyint | Augment Slot 6 Visible: 0 = False, 1 = True |
| augtype | int | [Augment Type](../../../../server/items/augment-types) |
| avoidance | int | Avoidance |
| awis | int | Wisdom: -128 to 127 |
| bagsize | int | [Bag Size](../../../../server/inventory/bag-sizes) |
| bagslots | int | Bag Slots: 1 = Minimum, 10 = Maximum |
| bagtype | int | [Bag Type](../../../../server/inventory/bag-types) |
| bagwr | int | Bag Weight Reduction: 0 = 0%, 100 = 100% |
| banedmgamt | int | Bane Damage Amount |
| banedmgraceamt | int | Bane Damage Race Amount |
| banedmgbody | int | [Bane Damage Body Type](../../../../server/npc/body-types) |
| banedmgrace | int | [Bane Damage Race](../../../../server/npc/race-list) |
| bardtype | int | [Bard Type](../../../../server/items/bard-types) |
| bardvalue | int | Bard Value |
| book | int | [Book](../../schema/books/books.md) |
| casttime | int | Cast Time in Seconds |
| casttime_ | int | Cast Time in Seconds |
| charmfile | varchar | Charm File |
| charmfileid | varchar | Charm File Identifier |
| classes | int | [Classes](../../../../server/player/class-list) |
| color | int | Color |
| combateffects | varchar | Combat Effects |
| extradmgskill | int | [Extra Damage Skill](../../../../server/player/skills) |
| extradmgamt | int | Extra Damage Amount |
| price | int | Price in Copper |
| cr | int | Cold Resistance: -128 to 127 |
| damage | int | Damage |
| damageshield | int | Damage Shield |
| deity | int | [Deity](../../../../server/player/deity-list) |
| delay | int | Delay |
| augdistiller | int | Augment Distiller Item Identifier |
| dotshielding | int | Damage Over Time Shielding |
| dr | int | Disease Resistance: -128 to 127 |
| clicktype | int | [Click Type](../../../../server/items/click-types) |
| clicklevel2 | int | Click Level 2 |
| elemdmgtype | int | [Elemental Damage Type](../../../../server/items/item-element-types) |
| elemdmgamt | int | Elemental Damage Amount |
| endur | int | Endurance |
| factionamt1 | int | Faction Amount 1 |
| factionamt2 | int | Faction Amount 2 |
| factionamt3 | int | Faction Amount 3 |
| factionamt4 | int | Faction Amount 4 |
| factionmod1 | int | Faction Modifier 1 |
| factionmod2 | int | Faction Modifier 2 |
| factionmod3 | int | Faction Modifier 3 |
| factionmod4 | int | Faction Modifier 4 |
| filename | varchar | File Name |
| focuseffect | int | [Focus Effect Identifier](../../schema/spells/spells_new.md) |
| fr | int | Fire Resistance: -128 to 127 |
| fvnodrop | int | Firiona Vie No Drop: 0 = False, 1 = True |
| haste | int | Haste: 0 = 0%, 255 = 255% |
| clicklevel | int | Click Level |
| hp | int | Health |
| regen | int | Health Regeneration |
| icon | int | Icon |
| idfile | varchar | Item Texture |
| itemclass | int | [Item Class](../../../../server/items/item-class) |
| itemtype | int | [Item Type](../../../../server/items/item-types) |
| ldonprice | int | LDoN Price |
| ldontheme | int | [LDoN Theme](../../../../server/zones/ldon-themes) |
| ldonsold | int | LDoN Sold: 0 = False, 1 = True |
| light | int | Light |
| lore | varchar | Lore Description |
| loregroup | int | [Lore Group](../../../../server/items/item-lore-groups) |
| magic | int | Magic: 0 = False, 1 = True |
| mana | int | Mana |
| manaregen | int | Mana Regeneration |
| enduranceregen | int | Endurance Regeneration |
| material | int | [Material](../../../../server/npc/textures) |
| herosforgemodel | int | Hero's Forge Model |
| maxcharges | int | Maximum Charges |
| mr | int | Magic Resistance: -128 to 127 |
| nodrop | int | No Drop: 0 = True, 1 = False |
| norent | int | No Rent: 0 = True, 1 = False |
| pendingloreflag | tinyint | Pending Lore Flag: 0 = False, 1 = True |
| pr | int | Poison Resistance: -128 to 127 |
| procrate | int | Proc Rate: 0 = 100%, 50 = 150%, 100 = 200% |
| races | int | [Races](../../../../server/npc/race-list) |
| range | int | Range: 0 to 255 |
| reclevel | int | Recommended Level |
| recskill | int | Recommended Skill Level |
| reqlevel | int | Required Level |
| sellrate | float | Sell Rate |
| shielding | int | Shielding: 5 = 5%, 20 = 20%, 50 = 50% |
| size | int | Size |
| skillmodtype | int | [Skill Modifier Type](../../../../server/player/skills) |
| skillmodvalue | int | Skill Modifier Value |
| slots | int | [Slots](../../../../server/inventory/item-slots) |
| clickeffect | int | [Click Effect Identifier](../../schema/spells/spells_new.md) |
| spellshield | int | Spell Shielding |
| strikethrough | int | Strikethrough |
| stunresist | int | Stun Resist |
| summonedflag | tinyint | Unknown |
| tradeskills | int | Tradeskill Item: 0 = False, 1= True |
| favor | int | Favor |
| weight | int | Weight: 10 = 1.0, 25 = 2.5, 100 = 10.0 |
| UNK012 | int | Unknown |
| UNK013 | int | Unknown |
| benefitflag | int | Unknown |
| UNK054 | int | Unknown |
| UNK059 | int | Unknown |
| booktype | int | [Book Language](../../../../server/player/languages) |
| recastdelay | int | Recast Delay in Seconds |
| recasttype | int | Recast Type: -1 = None, &gt;0 = Recast Type used across all items |
| guildfavor | int | Guild Favor |
| UNK123 | int | Unknown |
| UNK124 | int | Unknown |
| attuneable | int | Attuneable: 0 = False, 1 = True |
| nopet | int | No Pet: 0 = False, 1 = True |
| updated | datetime | Updated Datetime |
| comment | varchar | Comment |
| UNK127 | int | Unknown |
| pointtype | int | Unknown |
| potionbelt | int | Potion Belt: 0 = False, 1 = True |
| potionbeltslots | int | Potion Belt Slots |
| stacksize | int | Stack Size |
| notransfer | int | No Transfer: 0 = False, 1 = True |
| stackable | int | Stackable: 0 = False, 1 = True |
| UNK134 | varchar | Unknown |
| UNK137 | int | Unknown |
| proceffect | int | [Proc Effect Identifier](../../schema/spells/spells_new.md) |
| proctype | int | Proc Type: 0 |
| proclevel2 | int | Proc Level 2 |
| proclevel | int | Proc Level |
| UNK142 | int | Unknown |
| worneffect | int | [Worn Effect Identifier](../../schema/spells/spells_new.md) |
| worntype | int | Worn Type: 2 = Worn |
| wornlevel2 | int | Worn Level 2 |
| wornlevel | int | Worn Level |
| UNK147 | int | Unknown |
| focustype | int | Focus Type: 6 = Focus |
| focuslevel2 | int | Focus Level 2 |
| focuslevel | int | Focus Level |
| UNK152 | int | Unknown |
| scrolleffect | int | [Scroll Effect Identifier](../../schema/spells/spells_new.md) |
| scrolltype | int | Scroll Type: 7 = Scroll |
| scrolllevel2 | int | Scroll Level 2 |
| scrolllevel | int | Scroll Level |
| UNK157 | int | Unknown |
| serialized | datetime | Serialized Datetime |
| verified | datetime | Verified Datetime |
| serialization | text | Serialization |
| source | varchar | Source |
| UNK033 | int | Unknown |
| lorefile | varchar | Lore File |
| UNK014 | int | Unknown |
| svcorruption | int | Corruption Resistance: -128 to 127 |
| skillmodmax | int | Skill Modifier Max |
| UNK060 | int | Unknown |
| augslot1unk2 | int | Unknown |
| augslot2unk2 | int | Unknown |
| augslot3unk2 | int | Unknown |
| augslot4unk2 | int | Unknown |
| augslot5unk2 | int | Unknown |
| augslot6unk2 | int | Unknown |
| UNK120 | int | Unknown |
| UNK121 | int | Unknown |
| questitemflag | int | Quest Item: 0 = False, 1 = True |
| UNK132 | text | Unknown |
| clickunk5 | int | Unknown |
| clickunk6 | varchar | Unknown |
| clickunk7 | int | Unknown |
| procunk1 | int | Unknown |
| procunk2 | int | Unknown |
| procunk3 | int | Unknown |
| procunk4 | int | Unknown |
| procunk6 | varchar | Unknown |
| procunk7 | int | Unknown |
| wornunk1 | int | Unknown |
| wornunk2 | int | Unknown |
| wornunk3 | int | Unknown |
| wornunk4 | int | Unknown |
| wornunk5 | int | Unknown |
| wornunk6 | varchar | Unknown |
| wornunk7 | int | Unknown |
| focusunk1 | int | Unknown |
| focusunk2 | int | Unknown |
| focusunk3 | int | Unknown |
| focusunk4 | int | Unknown |
| focusunk5 | int | Unknown |
| focusunk6 | varchar | Unknown |
| focusunk7 | int | Unknown |
| scrollunk1 | int | Unknown |
| scrollunk2 | int | Unknown |
| scrollunk3 | int | Unknown |
| scrollunk4 | int | Unknown |
| scrollunk5 | int | Unknown |
| scrollunk6 | varchar | Unknown |
| scrollunk7 | int | Unknown |
| UNK193 | int | Unknown |
| purity | int | Purity |
| evoitem | int | Evolving Item: 0 = False, 1 = True |
| evoid | int | [Evolving Identifier](items.md) |
| evolvinglevel | int | Evolving Level |
| evomax | int | Evolving Max |
| clickname | varchar | Click Name |
| procname | varchar | Proc Name |
| wornname | varchar | Worn Name |
| focusname | varchar | Focus Name |
| scrollname | varchar | Scroll Name |
| dsmitigation | smallint | Damage Shield Mitigation |
| heroic_str | smallint | Heroic Strength |
| heroic_int | smallint | Heroic Intelligence |
| heroic_wis | smallint | Heroic Wisdom |
| heroic_agi | smallint | Heroic Agility |
| heroic_dex | smallint | Heroic Dexterity |
| heroic_sta | smallint | Heroic Stamina |
| heroic_cha | smallint | Heroic Charisma |
| heroic_pr | smallint | Heroic Poison Resistance |
| heroic_dr | smallint | Heroic Disease Resistance |
| heroic_fr | smallint | Heroic Fire Resistance |
| heroic_cr | smallint | Heroic Cold Resistance |
| heroic_mr | smallint | Heroic Magic Resistance |
| heroic_svcorrup | smallint | Heroic Corruption Resistance |
| healamt | smallint | Heal Amount: 0 to 32767 |
| spelldmg | smallint | Spell Damage: 0 to 32767 |
| clairvoyance | smallint | Clairvoyance |
| backstabdmg | smallint | Backstab Damage |
| created | varchar | Created |
| elitematerial | smallint | Elite Material |
| ldonsellbackrate | smallint | LDoN Sellback Rate |
| scriptfileid | mediumint | Script File Name |
| expendablearrow | smallint | Expendable Arrow: 0 = False, 1 = True |
| powersourcecapacity | mediumint | Powersource Capacity |
| bardeffect | mediumint | [Bard Effect Identifier](../../schema/spells/spells_new.md) |
| bardeffecttype | smallint | [Bard Effect Type](../../../../server/items/bard-types) |
| bardlevel2 | smallint | Bard Level 2 |
| bardlevel | smallint | Bard Level |
| bardunk1 | smallint | Unknown |
| bardunk2 | smallint | Unknown |
| bardunk3 | smallint | Unknown |
| bardunk4 | smallint | Unknown |
| bardunk5 | smallint | Unknown |
| bardname | varchar | Bard Name |
| bardunk7 | smallint | Unknown |
| UNK214 | smallint | Unknown |
| subtype | int | Sub Type |
| UNK220 | int | Unknown |
| UNK221 | int | Unknown |
| heirloom | int | Heirloom: 0 = False, 1 = True |
| UNK223 | int | Unknown |
| UNK224 | int | Unknown |
| UNK225 | int | Unknown |
| UNK226 | int | Unknown |
| UNK227 | int | Unknown |
| UNK228 | int | Unknown |
| UNK229 | int | Unknown |
| UNK230 | int | Unknown |
| UNK231 | int | Unknown |
| UNK232 | int | Unknown |
| UNK233 | int | Unknown |
| UNK234 | int | Unknown |
| placeable | int | Placeable: 0 = False, 1 = True |
| UNK236 | int | Unknown |
| UNK237 | int | Unknown |
| UNK238 | int | Unknown |
| UNK239 | int | Unknown |
| UNK240 | int | Unknown |
| UNK241 | int | Unknown |
| epicitem | int | Epic Item: 0 = False, 1 = True |

