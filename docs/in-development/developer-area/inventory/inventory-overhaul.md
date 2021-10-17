# Inventory Overhaul

**Phase 1a: Re-enumerate 'typePossessions' slots \[development 100%, implementation: 100%\]**

| Token | Old Value | New Value |
| :--- | :--- | :--- |
| slotCharm | 0 | 0 |
| slotEar1 | 1 | 1 |
| slotHead | 2 | 2 |
| slotFace | 3 | 3 |
| slotEar2 | 4 | 4 |
| slotNeck | 5 | 5 |
| slotShoulders | 6 | 6 |
| slotArms | 7 | 7 |
| slotBack | 8 | 8 |
| slotWrist1 | 9 | 9 |
| slotWrist2 | 10 | 10 |
| slotRange | 11 | 11 |
| slotHands | 12 | 12 |
| slotPrimary | 13 | 13 |
| slotSecondary | 14 | 14 |
| slotFinger1 | 15 | 15 |
| slotFinger2 | 16 | 16 |
| slotChest | 17 | 17 |
| slotLegs | 18 | 18 |
| slotFeet | 19 | 19 |
| slotWaist | 20 | 20 |
| slotPowerSource | 9999 | 21\* |
| slotAmmo | 21 | 22\* |
| slotGeneral1 | 22 | 23\* |
| slotGeneral2 | 23 | 24\* |
| slotGeneral3 | 24 | 25\* |
| slotGeneral4 | 25 | 26\* |
| slotGeneral5 | 26 | 27\* |
| slotGeneral6 | 27 | 28\* |
| slotGeneral7 | 28 | 29\* |
| slotGeneral8 | 29 | 30\* |
| slotGeneral9 | n/a | 31\* |
| slotGeneral10 | n/a | 32\* |
| slotCursor | 30 | 33\* |

**Phase 1b: Insert General9 and General10 bag slots into current General bag slot range \[development: 100% complete, implementation: 100%\]**

| Token | Old Begin | Old End Value | New Begin Value | New End Value |
| :--- | :--- | :--- | :--- | :--- |
| GENERAL1BAGSLOTS | 251 | 260 | 251 | 260 |
| GENERAL2BAGSLOTS | 261 | 270 | 261 | 270 |
| GENERAL3BAGSLOTS | 271 | 280 | 271 | 280 |
| GENERAL4BAGSLOTS | 281 | 290 | 281 | 290 |
| GENERAL5BAGSLOTS | 291 | 300 | 291 | 300 |
| GENERAL6BAGSLOTS | 301 | 310 | 301 | 310 |
| GENERAL7BAGSLOTS | 311 | 320 | 311 | 320 |
| GENERAL8BAGSLOTS | 321 | 330 | 321 | 330 |
| GENERAL9BAGSLOTS | n/a | n/a | 331\* | 340\* |
| GENERAL10BAGSLOTS | n/a | n/a | 341\* | 350\* |
| CURSORBAGSLOTS | 331 | 340 | 351\* | 360\* |

**Phase 2: Modify inventory to allow for supported client expansion changes and GM status \[development: 100% complete, implementation: 100%\]**

* expansion bitmask affects the various clients differently..but, these changes in client behavior are not accounted for
* GM flag state typically overrides expansion bitmask and these changes in client behavior also need to be accounted for

**Phase ?: Rework class ItemInstQueue for proper behavior \[development: 0%\]**

* this is the 'cursor buffer' used for server operations and does not reflect any client behavior
* it will fix a lot of remaining cursor issues - especially RoF+ based ones
* separate type possessions cursor \(slot 33\) from type limbo slots \(slots {0..35}\)
* implement server bag slots for type limbo main slots \(36 \* 10 = 360 allocation requirement\)
* should be doable without a sql update

**Phase ?: Change merchant lists to similar method as corpse's InternalLoot\_Struct \[development: 0%\]**

* merchants do not like to have their 'slot' enumerations non-contiguous
* remove the slot identifier from database schema
* enumerate merchant items upon loading and as players sell items to them
* create per-client merchant windows and track them
* as players transact per merchant window ids, cross-reference those back to merchant master list ids

**Phase ?: Increase number of allowed bag slots and remap all bag slots to an unused range \[development: 0%\]**

* hard-code the max server-allowed bag slot size to 100 slots \(should account for anything we would ever support\)
* remap _all_ bags slots into a new range .. 10000+ should be available
* update EQDictionary to reflect the new maximum sizes

**Phase ?: Implement class ItemInstance::origin\_tag \[development: 100% complete, implementation: 0% complete\]**

* a 16-byte char field that represents a base64-encoded unique identifier for every item \(12 bytes of int data encoded into 16 ascii bytes\)
* used in numerous systems, such as: bazaar/barter, hotkey assignments, etc...
* does not use the same encoding hash as SocketLib [https://github.com/EQEmu/Server/tree/master/common/SocketLib](https://github.com/EQEmu/Server/tree/master/common/SocketLib)
* implementation allows for multiple encoding hash references

**Phase ?: Rework bazaar code to make proper use of ItemInstance::origin\_tag property \[development: 80% complete\]**

* this will also fix the 'broken' system for newer clients

