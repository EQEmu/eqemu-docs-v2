# Spells

## Getting Started

So you would like to understand spells? It's a bit of a journey, so let's get started.

There are tens of thousands of records in your [spells_new](https://eqemu.gitbook.io/database-schema/tables/spells_new) table, and making sense of them all can be a little daunting (at least at first).

Each spell has MANY fields related to it--and you're probably wondering what all of these fields do. You may even be up against some [Client Limitations](https://eqemu.gitbook.io/server/categories/reference-lists/client-spell-id-limitations), or perhaps you've just discovered that AAs are actually spells! Don't get too worked up!! (╯°□°）╯彡 ┻━┻

We'll explore these topics below.

### Slots

The first concept is that many fields are related to each other. We have commonly referred to these related fields as "Slots", and most understand the concept that slots are somehow related. There are 12 of these slots available, and the values in your database determine what happens when the spell is cast.

It might be easiest to view in simplified table, rather than as a series of fields. It will also be useful to disregard many of the fields so that we can focus on understanding the mechanics of the spell, before we try to understand all of the fields in the spells_new database table. Below is a simplified table of the spell "Greater Healing", that has a default Spell ID of 15:

| **Slot** | **Effect ID** | **Effect Base Value** | **Max** | **Formula** |
| -------- | ------------- | --------------------- | ------- | ----------- |
| 1        | 0             | 140                   | 300     | 7           |
| 2        | 254           | 0                     | 0       | 100         |
| 3        | 254           | 0                     | 0       | 100         |
| 4        | 254           | 0                     | 0       | 100         |
| 5        | 254           | 0                     | 0       | 100         |
| 6        | 254           | 0                     | 0       | 100         |
| 7        | 254           | 0                     | 0       | 100         |
| 8        | 254           | 0                     | 0       | 100         |
| 9        | 254           | 0                     | 0       | 100         |
| 10       | 254           | 0                     | 0       | 100         |
| 11       | 254           | 0                     | 0       | 100         |
| 12       | 254           | 0                     | 0       | 100         |

This simplified table shows that Slot 1 is where all the action is in the Greater Healing spell. Slots 2-12 in this case just have default values, as they are unused. If you're looking in your database at your spells_new table, you will see that each header row from table above appears as a database field for each slot (IE effectid1, effectid2...effectid12; max1, max2...max12).

### Spell Effects

There is an entire page on the Wiki describing each [Spell Effect](https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids). A spell effect, in simplest terms, determines what is modified, or what is caused to happen, by the spell. It is important to understand that each slot, as shown above, can have a Spell Effect enumerated in the Effect ID field--that is to say that each Spell can have one or more spell effects that occur when the spell is cast.

Dissecting the information presented in the Greater Healing spell table above, we see that Spell Effect 0 is listed in the first slot (effectid1 field in your database). Cross-referencing with the Spell Effect list, we find that Spell Effect 0 is "SE_CurrentHP", or Spell Effect Current Hit Points. We can therefore surmise the the Greater Healing spell must have something to do with hit points.

### Base Values

Base Values are integers that are used in coordination with the Spell Effect to determine an outcome. Each Spell Effect does something slightly different with its base value, as demonstrated in the Spell Effects table.

In the case of our Greater Healing spell, we see a Base Value of 140 in Slot 1 (effect_base_value1 field in your database). Since we know the Spell Effect will impact Current Hit Points, and since 140 is a positive value (as opposed to -140, for instance), we can say that the Healing spell will increase current hitpoints by 140--it will heal (add), rather than damage (subtract) from our current hit points.

Something isn't quite adding up yet, however: we know that as you level up, your healing power seems to increase. This mechanic can be explained by examining some additional fields.

### Max Values

Max Values are integers that are also used in coordination with the Spell Effect to determine an outcome. Each Spell Effect's Max Value can do something slightly different, as demonstrated in the Spell Effect table.

In the case of our Greater Healing spell, we see a Max Value of 300 in Slot 1 (max1 field in your database). This means that the most this spell will adjust Current Hit Points by, is 300. This is indicated by "Max Amt" in the Max Value column of the Spell Effect table for SE_CurrentHP. Instead, if "Max Level" were indicated for Max Value of this spell effect, we would surmise that the spell would continue to scale as the player character gained levels, with each new level adding to the amount of hit points that could be restored.

So we have established that the Greater Healing spell is capable of changing current hit points in the amount of 140 to a maximum amount of 300--but how does it scale as you increase in levels from the first level you get the spell, until you max out the effectiveness of the spell and are able to heal 300 hit points?

### Formula

Formula values are integers that are used in coordination with the Spell Effect to determine its outcome as well. Each Spell Effect can use a different formula, and those are found in the [spell_effects.cpp](https://github.com/EQEmu/Server/blob/master/zone/spell_effects.cpp) source file. A wiki page also details those [Base Value Formulas](https://eqemu.gitbook.io/server/categories/spells/base-value-formulas) if you're not comfortable looking at the code.

In the case of our Greater Healing spell, we see that the value for Formula in Slot 1 is 7 (formula1 field in your database). Referencing the source, we find that the spell calc is Base Value + User Level _ _\* Formula. If we assume that we are a level 20 Cleric, casting the Greater Healing spell: 140 + 20 \* 7 = 280. This of course assumes that there are no focus effects being taken into consideration. If we are instead a level 29 Druid, we would hit the spell's max value: 140 + 29 \* 7 = 343--which is greater than the 300 hit point Max Value. As a result, the Druid would only heal 300 hit points using this spell.

### Summary

It's important to understand the concept that at least one Spell Effect is in use for each spell. Each spell effect has a number of parameters that may behave differently from one spell effect to the next, even though they occupy the same field in the database--pay close attention to the Base Value, Limit Value, and Max Value for the Spell Effect, as they can drastically change the behavior of a spell.

## Dissecting more Complex Spells

Every spell that exists can be deciphered simply by following the guidelines as laid out above. Let's look at a couple more spells to hone our skills.

### Koadic's Endless Intellect

Many players will remember the endless requests for KEI--players always needed more mana. Let's examine the spell:

| **Slot** | **Effect ID** | **Effect Base Value** | **Max** | **Formula** |
| -------- | ------------- | --------------------- | ------- | ----------- |
| 1        | 97            | 250                   | 250     | 100         |
| 2        | 15            | 14                    | 14      | 100         |
| 3        | 8             | 25                    | 25      | 100         |
| 4        | 9             | 25                    | 25      | 100         |
| 5        | 254           | 0                     | 0       | 100         |
| 6        | 254           | 0                     | 0       | 100         |
| 7        | 254           | 0                     | 0       | 100         |
| 8        | 254           | 0                     | 0       | 100         |
| 9        | 254           | 0                     | 0       | 100         |
| 10       | 254           | 0                     | 0       | 100         |
| 11       | 254           | 0                     | 0       | 100         |
| 12       | 254           | 0                     | 0       | 100         |

We see that slots 1 through 4 have something going on, and slots 5 through 12 are not being utilized. Let's take a look at the parameters.

#### Spell Effects

Cross reference the Spell Effects table.

`Slot 1: 97 - SE_ManaPool: Increase/Decrease max mana pool `\
`Slot 2: 15 - SE_CurrentMana: Direct +/- mana, duration allows for mana over time `\
`Slot 3: 8 - SE_INT: +/- INT `\
`Slot 4: 9 - SE_WIS: +/- WIS `\


#### Effect Base Value

`Slot 1: SE_ManaPool - +250 `\
`Slot 2: SE_CurrentMana - +14 `\
`Slot 3: SE_INT - +25 `\
`Slot 4: SE_WIS - +25 `\


#### Max

`Slot 1: SE_ManaPool - 250 `\
`Slot 2: SE_CurrentMana - 14 `\
`Slot 3: SE_INT - 25 `\
`Slot 4: SE_WIS - 25 `\


#### Formula

All slots use Formula 100. When we analyze the [spell_effects.cpp](https://github.com/EQEmu/Server/blob/master/zone/spell_effects.cpp) source file, we see the corresponding formula establishes that the base value is used (result = ubase).

Since the Spell Effect Current Mana (SE_CurrentMana) allows for duration, we see that you will receive 14 mana per tick while the spell is active.

#### Summary

Koadic's Endless Intellect:

| **Slot** | **Description**              |
| -------- | ---------------------------- |
| 1        | Increase Mana Pool by 250    |
| 2        | Increase Mana by 14 per tick |
| 3        | Increase INT by 25           |
| 4        | Increase WIS by 25           |

### Protection of the Glades

Most players will recall being a caster class and requesting POTG along with their KEI. Let's take a look at the spell and find out why:

| **Slot** | **Effect ID** | **Effect Base Value** | **Max** | **Formula** |
| -------- | ------------- | --------------------- | ------- | ----------- |
| 1        | 1             | 25                    | 80      | 100         |
| 2        | 69            | 290                   | 485     | 100         |
| 3        | 79            | 290                   | 485     | 100         |
| 4        | 15            | 6                     | 0       | 100         |
| 5        | 148           | 1                     | 1081    | 201         |
| 6        | 148           | 69                    | 2486    | 202         |
| 7        | 254           | 0                     | 0       | 100         |
| 8        | 254           | 0                     | 0       | 100         |
| 9        | 254           | 0                     | 0       | 100         |
| 10       | 254           | 0                     | 0       | 100         |
| 11       | 254           | 0                     | 0       | 100         |
| 12       | 254           | 0                     | 0       | 100         |

We see that slots 1 through 6 have something going on, and slots 7 through 12 are not being utilized. Let's take a look at the parameters.

#### Spell Effects

Cross reference the Spell Effects table.

`Slot 1: 1 - SE_ArmorClass: +/- AC `\
`Slot 2: 69 - SE_TotalHP: Increases Max HP (standard 'HP Buffs') `\
`Slot 3: 79 - SE_CurrentHPOnce: Direct Damage/Healing `\
`Slot 4: 15 - SE_CurrentMana: Direct +/- mana, duration allows for mana over time `\
`Slot 5: 148 - SE_StackingCommand_Block: Prevents buff from taking hold if criteria met. (SLOT = 'formula - 201') Buff Stacking `\
`Slot 6: 148 - SE_StackingCommand_Block: Prevents buff from taking hold if criteria met. (SLOT = 'formula - 201') Buff Stacking `\


#### Effect Base Value

`Slot 1: SE_ArmorClass - +25 `\
`Slot 2: SE_TotalHP - +290 `\
`Slot 3: SE_CurrentHPOnce - +290 (heal up the new-found reservoir of hit points) `\
`Slot 4: SE_CurrentMana - +6 (remember duration allows for mana over time) `\
`Slot 5: SE_StackingCommand_Block - 1 (block spells using SE_ArmorClass) `\
`Slot 6: SE_StackingCommand_Block - 69 (block spells using SE_TotalHP) `\


#### Max

`Slot 1: SE_ArmorClass - 80 `\
`Slot 2: SE_TotalHP - 485 `\
`Slot 3: SE_CurrentHPOnce - 485 `\
`Slot 4: SE_CurrentMana - 0 (not used) `\
`Slot 5: SE_StackingCommand_Block - Less than 1081 (block spells with less than 1081 AC addition)`\
`Slot 6: SE_StackingCommand_Block - Less than 2486 (block spells with less than 2486 HP addition)`

#### Formula

All slots 1 through 4 use Formula 100. When we analyze the [spell_effects.cpp](https://github.com/EQEmu/Server/blob/master/zone/spell_effects.cpp) source file, we see the corresponding formula establishes that the base value is used (result = ubase).

Since the Spell Effect Current Mana (SE_CurrentMana) allows for duration, we see that you will receive 6 mana per tick while the spell is active.

Slots 5 and 6, utilizing SE_StackingCommand_Block, will apply to the slot indicated by calculating 'formula - 201'. So slot 5 applies to slot 1 (201 - 201 = 0, remember that computers count 0, 1, 2...), and slot 6 applies to slot 2 (202 - 201 = 1, remembering that computers count 0, 1, 2...).

#### Summary

Protection of the Glades:

| **Slot** | **Description**                                                 |
| -------- | --------------------------------------------------------------- |
| 1        | Increase Armor Class by 25 - 80                                 |
| 2        | Increase Hit Points by 290 - 485                                |
| 3        | Increase Hit Points when cast by 290 - 485                      |
| 4        | Increase Mana by 6 per tick                                     |
| 5        | Block new spell if slot 1 has AC effect less than 1081          |
| 6        | Block new spell if slot 2 has increase HP effect less than 2486 |

## Understanding the Remaining Fields

A dizzying array of fields exists that have not been covered above. Many of these fields are quite self-explanatory, such as the level at which each class can use the spell, the deities allowed to use the spell, cast-on messages, etc. This section is meant to cover the remaining fields as questions arise about their functionality. Remember to reference the information on [spells_new](https://eqemu.gitbook.io/database-schema/tables/spells_new) found on this wiki, and the information contained in the [spdat.h](https://github.com/EQEmu/Server/blob/master/common/spdat.h) source file, as well as the other helpful wiki pages for [Base Value Formulas](https://eqemu.gitbook.io/server/categories/spells/base-value-formulas), [Spell Target Restrictions](https://eqemu.gitbook.io/server/categories/spells/spell-target-restrictions), [Spell Resist Types](spell-resist-types.md), and [Damage Shield Types](https://eqemu.gitbook.io/server/categories/spells/damage-shield-types).

## Reagents

Spells that require reagents have items (by Item ID) listed in the components fields (component1, component2...component4). Examples of such spells would be Elemental: Earth (default Spell ID 401), which requires Item ID 10015 - Malachite, or Diamondskin (default Spell ID 394), which requires Item ID 10028 - Peridot. Some spells require a focus item, but do not consume the item when cast, such as Flame Lick (default Spell ID 239), which requires a 10307 - Fire Beetle Eye to cast, but does not consume it. Note the relationship between the component fields, the component count fields, and the NoexpendReagent fields, which utilize slots--just like Spell Effects discussed above:

| **Slot** | **Component** | **Component Count** | **No Expend Reagent**                        |
| -------- | ------------- | ------------------- | -------------------------------------------- |
| 1        | ItemID        | Count of the ItemID | Used as Focus item, or "Needed But Not Used" |
| 2        | ItemID        | Count of the ItemID | Used as Focus item, or "Needed But Not Used" |
| 3        | ItemID        | Count of the ItemID | Used as Focus item, or "Needed But Not Used" |
| 4        | ItemID        | Count of the ItemID | Used as Focus item, or "Needed But Not Used" |

#### Elemental: Earth

| **Slot** | **Component** | **Component Count** | **No Expend Reagent** | **Notes**                                                  |
| -------- | ------------- | ------------------- | --------------------- | ---------------------------------------------------------- |
| 1        | 10015         | 1                   | 6361                  | Use 10015 - Malachite x1, focus with 6361 - Shovel of Ponz |
| 2        | -1            | 1                   | -1                    | Nothing                                                    |
| 3        | -1            | 1                   | -1                    | Nothing                                                    |
| 4        | -1            | 1                   | -1                    | Nothing                                                    |

#### Diamondskin

| **Slot** | **Component** | **Component Count** | **No Expend Reagent** | **Notes**              |
| -------- | ------------- | ------------------- | --------------------- | ---------------------- |
| 1        | 10028         | 1                   | -1                    | Use 10028 - Peridot x1 |
| 2        | -1            | 1                   | -1                    | Nothing                |
| 3        | -1            | 1                   | -1                    | Nothing                |
| 4        | -1            | 1                   | -1                    | Nothing                |

#### Flame Lick

| **Slot** | **Component** | **Component Count** | **No Expend Reagent** | **Notes**                                 |
| -------- | ------------- | ------------------- | --------------------- | ----------------------------------------- |
| 1        | -1            | 1                   | 10307                 | Require a 10307 - Fire Beetle Eye to cast |
| 2        | -1            | 1                   | -1                    | Nothing                                   |
| 3        | -1            | 1                   | -1                    | Nothing                                   |
| 4        | -1            | 1                   | -1                    | Nothing                                   |

## Buff Duration

Two fields impact the duration of buffs: buffduration (in ticks) and buffduration formula. Understanding the formula allows you to calculate the duration of the buff. The formula can be found in the [spells.cpp](https://github.com/EQEmu/Server/blob/master/zone/spells.cpp) source file.

### Buff Duration Formula 0

Not a buff! This formula is used for instant spells. Ex. [Greater Healing](http://lucy.allakhazam.com/spell.html?id=15\&source=Live)

### Buff Duration Formula 1

**If your level is greater than 3:**

Choose the lesser integer of the two values: Level ÷ 2, or the buffduration field.

Example: [Heat Blood](http://lucy.allakhazam.com/spell.html?id=360\&source=Live), Spell ID 360, buffduration field = 6.

_As a level 10 NEC:_

```
Level ÷ 2 = X  
10 ÷ 2 = X  
Is X < 6?  
Yes, 5 is less than 6
Smallest Answer: 5 ticks
```

_As a level 28 SHD:_

```
Level ÷ 2 = X
28 ÷ 2 = X
Is X < 6?  
No, 14 is greater than 6
Smallest Answer: 6 ticks
```

**If your level is less than 3:**

1 tick.

### Buff Duration Formula 2

**If your level is greater than 3:**

Choose the lesser integer of the two values: Level ÷ 2 + 5, or the buffduration field.

Example: [Root](http://lucy.allakhazam.com/spell.html?id=230\&source=Live), Spell ID 230, buffduration field = 8.

_As a level 4 WIZ:_

```
Level ÷ 2 + 5 = X  
4 ÷ 2 + 5 = X  
Is X < 8?  
Yes, 7 is less than 8.
Smallest Answer: 7 ticks
```

_As a level 34 NEC:_

```
Level ÷ 2 + 5= X  
34 ÷ 2 + 5 = X  
Is X < 8?  
No, 19 is greater than 8
Smallest Answer: 8 ticks
```

**If your level is less than 3:**

6 ticks.

### Buff Duration Formula 3

Choose the lesser integer of the two values: Level x 30, or the buffduration field.

Example: [Invisibility](http://lucy.allakhazam.com/spell.html?id=42\&source=Live), Spell ID 42, buffduration field = 200.

_As a level 4 ENC:_

```
Level x 30 = X  
4 x 30 = X  
Is X < 200?  
Yes, 120 is less than 200.
Smallest Answer: 120 ticks
```

_As a level 16 WIZ:_

```
Level x 30 = X  
16 x 30 = X  
Is X < 200?  
No, 480 is greater than 200.
Smallest Answer: 200 ticks
```

### Buff Duration Formula 4

50, unless buffduration is lower.

Example: [Resurrection Sickness](http://lucy.allakhazam.com/spell.html?id=757\&source=Live), Spell ID 757, buffduration field = 50.

```
buffduration = X
Is X < 50?  
No, they are the same.
Smallest Answer: 50 ticks
```

Example: [Song of the Deep Seas](http://lucy.allakhazam.com/spell.html?id=1973\&source=Live), Spell ID 1973, buffduration field = 25.

```
buffduration = X
Is X < 50?  
Yes, 25 is less than 50.
Smallest Answer: 25 ticks
```

### Buff Duration Formula 5

2, unless buffduration is lower.

Example: [Flash of Light](http://lucy.allakhazam.com/spell.html?id=201\&source=Live), Spell ID 201, buffduration field = 2.

```
buffduration = X
Is X < 2?  
No, they are the same.
Smallest Answer: 2 ticks
```

Example: [The Unspoken Word](http://lucy.allakhazam.com/spell.html?id=1545\&source=Live), Spell ID 1545, buffduration field = 1.

```
buffduration = X
Is X < 2?  
Yes, 1 is less than 2.
Smallest Answer: 1 ticks
```

### Buff Duration Formula 6

Choose the lesser integer of the two values: Level ÷ 2 + 2, or the buffduration field.

Example: [Drowsy](http://lucy.allakhazam.com/spell.html?id=270\&source=Live), Spell ID 270, buffduration field = 35.

_As a level 6 SHM:_

```
Level ÷ 2 + 2 = X  
6 ÷ 2 + 2 = X  
Is X < 35?  
Yes, 5 is less than 35.
Smallest Answer: 5 ticks
```

_As a level 70 SHM:_

```
Level ÷ 2 + 2 = X  
70 ÷ 2 + 2 = X  
Is X < 35?  
No, 37 is greater than 35
Smallest Answer: 35 ticks
```

### Buff Duration Formula 7

Choose the lesser integer of the two values: Level, or the buffduration field.

Example: [Fear](http://lucy.allakhazam.com/spell.html?id=229\&source=Live), Spell ID 229, buffduration field = 3.

_As a level 2 NEC:_

```
Level = X  
Is X < 3?  
Yes, 2 is less than 3.
Smallest Answer: 2 ticks
```

_As a level 12 SHD:_

```
Level = X  
Is X < 3?  
No, 12 is greater than 3.
Smallest Answer: 3 ticks
```

### Buff Duration Formula 8

Choose the lesser integer of the two values: Level + 10, or the buffduration field.

Example: [Vampiric Embrace](http://lucy.allakhazam.com/spell.html?id=359\&source=Live), Spell ID 359, buffduration field = 75.

_As a level 7 NEC:_

```
Level + 10 = X  
7 + 10 = X
Is X < 75?  
Yes, 17 is less than 75.
Smallest Answer: 17 ticks
```

_As a level 70 SHD:_

```
Level + 10 = X 
70 + 10 = X 
Is X < 75?  
No, 80 is greater than 75.
Smallest Answer: 75 ticks
```

### Buff Duration Formula 9

Choose the lesser integer of the two values: 2 x Level + 10, or the buffduration field.

Example: [Ensnare](http://lucy.allakhazam.com/spell.html?id=512\&source=Live), Spell ID 512, buffduration field = 140.

_As a level 26 DRU:_

```
2 x Level + 10 = X  
2 x 26 + 10 = X
Is X < 140?  
Yes, 62 is less than 140.
Smallest Answer: 62 ticks
```

_As a level 70 RNG:_

```
2 x Level + 10 = X  
2 x 70 + 10 = X
Is X < 140?  
No, 150 is greater than 140.
Smallest Answer: 140 ticks
```

### Buff Duration Formula 10

Choose the lesser integer of the two values: 3 x Level + 10, or the buffduration field.

Example: [Regeneration](http://lucy.allakhazam.com/spell.html?id=144\&source=Live), Spell ID 144, buffduration field = 205.

_As a level 34 DRU:_

```
3 x Level + 10 = X  
3 x 34 + 10 = X
Is X < 205?  
Yes, 112 is less than 205.
Smallest Answer: 112 ticks
```

_As a level 70 DRU:_

```
3 x Level + 10 = X  
3 x 70 + 10 = X
Is X < 205?  
No, 220 is greater than 205.
Smallest Answer: 205 ticks
```

### Buff Duration Formula 11

Choose the lesser integer of the two values: 30 x (Level + 3), or the buffduration field.

Example: [See Invisible](http://lucy.allakhazam.com/spell.html?id=80\&source=Live), Spell ID 80, buffduration field = 270.

_As a level 4 WIZ:_

```
30 x (Level + 3) = X  
30 x (4 + 3) = X
Is X < 270?  
Yes, 210 is less than 270.
Smallest Answer: 210 ticks
```

_As a level 32 RNG:_

```
30 x (Level + 3) = X  
30 x (32 + 3) = X
Is X < 270?  
No, 1050 is greater than 270.
Smallest Answer: 270 ticks
```

### Buff Duration Formula 12

**If your level is greater than 7:**

Choose the lesser integer of the two values: Level ÷ 4, or the buffduration field.

Used by some focuses and disciplines.

**If your level is less than 7:**

1 tick.

### Buff Duration Formula 13

Choose the lesser integer of the two values: 4 x Level + 10, or the buffduration field.

Used by some poisons and wards.

### Buff Duration Formula 14

Choose the lesser integer of the two values: 5 x (Level + 2), or the buffduration field.

Used by a ridiculous secretion.

### Buff Duration Formula 15

Choose the lesser integer of the two values: 10 x (Level + 10), or the buffduration field.

Used by some potions and powerful spells.

### Buff Duration Formula 50

Permanent. Used for permanent buffs that might get stripped due to other forces (IE casting when using [Perfected Invisibility](http://lucy.allakhazam.com/spell.html?id=13219\&source=Live)).

### Buff Duration Formula 51

Permanent. Used for auras and the like.
