---
description: This page describes the evolving items currently available on your EQEmu Server.  This requires release 22.xx.
---

# Evolving Items

The evolving items system is only available on the Rain of Fear 2 (RoF2) Client.  It requires Server Release 22.62.0+

- Allows for evolving item support for RoF2 only.
- Has a 30 sec timer after equipping an evolving item before it can start to evolve. Time is configurable with a Rule.
XP gain is Rule configurable for raid, solo, and group xp.
- Enables the Evolving XP Transfer in Corathus Creep. This is limited at present as I have little information. It currently allows for xp transfer between items that have the same type, and subtype. I can and will update this as feedback comes in.
- PlayerEvent logging is enabled for various functions, namely progression of an item and xp transfer between items.
- Evolving item status remains with the item, and will transfer if an tradeable item is traded. 
- There are two gm commands to monitor statis. #evolve target and #evolve item


The system is controlled using the following server rules:

| Category | Rule Name | Default | Description |
| :--- | :--- | :--- | :--- |
| EvolvingItems | PercentOfSoloExperience | 0.1 | Percentage of solo experience allocated to evolving items that require experience. Default is 0.1% | 
| EvolvingItems | PercentOfGroupExperience | 0.1 | Percentage of group experience allocated to evolving items that require experience.  Default is 0.1%| 
| EvolvingItems | PercentOfRaidExperience | 0.1 | Percentage of solo experience allocated to evolving items that require experience.  Default is 0.1%| 
| EvolvingItems | DelayUponEquipping | 30000 | Delay in ms before an evolving item will earn rewards after equipping.  Default is 30000ms or 30s. | 

## Installation

This update requires that your opcodes be updated.  If you are using Spire, this can be done from within that tool, otherwise the opcodes are located at within the source at /utils/patches/patches_RoF2.conf.  
  
The opcode in question is: OP_EvolveItem=0x7cfb

Opcodes need to be reloaded, or a server restart to take effect.

There are two tables utilized in the functionality.

#character_evolving_items
This table contains character specific data surrounding the evolving items in the characters possession. If a character trades, deletes, sells, etc an evolving item the deleted_at column is updated.  This allows for a history of items that the character has had.  Rows are not deleted.

#items_evolving_details
This table is sourced from the content database and contains the generic evolution data regarding an item.  There are several generic columns, though there are three that should be understood:

These two tables are the most important to enable actual items.

**type**
References the how the item evolves.  Types are sourced in the namespace EvolvingItems::Types and are listed below. This list will be expanded as more is known about evolving items.

|Type|Description|
|:---:|:---|
|1|Experience based evolution|
|2|Number of Kills (not yet implemented)|
|3|Specific Mob Race|
|4|Specific Zone ID|

* Type 99 is a placeholder and identifies items not yet configured.

|SubType|Description|
|:---:|:---|
|0|When type = 1, use all experience|
|1|When type = 1, use only solo experience|
|2|When type = 1, use only group experience|
|3|When type = 1, use only raid experience|
|x|When type = 2, NOT YET IMPLEMENTED|
|x|When type = 3, the Mob Race ID of the kill|
|x|When type = 4, the Zone ID where the kill must occur|

|required_amount|Description|
|:---:|:---|
|x|The amount required of the type/subtype combination to evolve that level of the item|

There are two items with presets to demonstrate

**Evolution ID 1211**
Level 1 is item id 85612 requiring 100,000 of experience from any source
Level 2 is item id 85613 requiring 200,000 of experience from any source
Level 3 is item id 85614 requiring 300,000 of experience from any source

**Evolving ID 1066**
Level 1 is item id 89550 requring 500 kills of race id 274
Level 2 is item id 89551 requring 1000 kills of race id 274
Level 3 is item id 89552 requring 2000 kills of race id 274

## XP Transfer

There is an witch in Corathus Creep who will preform basic transfer of evolving item experience.  The current implementation allows for the following:

- If sub_types are the same, the conversion rate is 100% otherwise it is 30%
- The current amount is the total of all the levels currently obtained

If there are any suggests, please create an issue and I can work to add more functionality.

## Admin Commands

There are several admin commands to assist GMs.

- \#evolve item item_id  
  This allows the GM to query the details of an item.
  Example: #evolve item 85613

- \#evolve target  
  Please ensure to have a player targeted.
  This is used to display the evolving details for all items that the player has in their possession.

## Player Event Logging

The following player events are available.

|Rule Name|Default|Retention|
|:---|:---|:---|
|Evolve Item|Enabled|7 days|

The playerevent has a status field which captures a more detailed explanation.  
Status messages such as:  
- Evolved Item due to obtaining progression - Old Evolve Item removed from inventory.  
- Evolved Item due to obtaining progression - New Evolve Item placed in inventory.  
- Transfer XP - Original FROM Evolve Item removed from inventory.  
- Transfer XP - Updated FROM item placed in inventory.  
- Transfer XP - Original TO Evolve Item removed from inventory.  
- Transfer XP - Updated TO Evolve item placed in inventory.  

## Creating new items

If you would like to create a new item, not part of the exisiting evolving items already established, please follow these steps.

- Create a new 'item' in the items tables.  This is found in your content database.
- There are 4 fields in the items table that play a role for evolving items.  

|Field|Description|
|:---|:---|
|evoitem| 1 = evolving item, 0 = not an evolving item|
|evoid| a unique id within all evolving items, and shared between all items of the same evolving group|
|evolvinglevel| the evolving level of this item within the evolving group id (1, 2, 3, etc)|
|evomax| the max evolving level of the evolving group id|

For example, for the items with evoid = 1066

|Item ID|evoitem|evoid|evolvinglevel|evomax|
|:---|:---|:---|:---|:---|
|89550|1|1066|1|3|
|89551|1|1066|2|3|
|89552|1|1066|3|3|

- To enable the new item(s) just created, you need to rerun shared_memory and restart your server.
- Create the records in items_evolving_details as noted above.

That should be it.  
If you have ideas about other ways to evolve an item, please post an issue at https://github.com/EQEmu/Server/issues

## Known Issues

There are currently no known issues.
