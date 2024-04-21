---
description: This page describes the parcel system on your EQEmu Server.
---

# Parcels

The parcel feature is only available on the Rain of Fear 2 (RoF2) Client.

The system is activated using the following server rules:

| Category | Rule Name | Default | Description |
| :--- | :--- | :--- | :--- |
| Parcel | DeleteOnDuplicate | false | Delete retrieved item if it creates a lore conflict. | 
| Parcel | EnableDirectToInventoryDelivery | false | Enable or Disable RoF2 bazaar purchases to be delivered directly to the buyer's inventory. | 
| Parcel | EnableParcelMerchants | true | Enable or Disable Parcel Merchants.  Requires RoF+ Clients. | 
| Parcel | EnablePruning | false | Enable the automatic pruning of sent parcels.  Uses rule ParcelPruneDelay for prune delay. | 
| Parcel | ParcelDeliveryDelay | 30000 | Sets the time that a player must wait between sending parcels in ms. | 
| Parcel | ParcelMaxItems | 50 | The maximum number of parcels a player is allowed to have in their mailbox. | 
| Parcel | ParcelPruneDelay | 30 | The number of days after which a parcel is deleted. Items are lost! | 

There are Several merchants across Norrath which can be enabled for parcel management. 

## Installation

Parcel functionality requires:

1. SQL Updates.  Please run the sql file located at (https://github.com/EQEmu/Server/tree/master/utils/sql/git/optional).  The file is 2024_04_20_Parcel_Merchant_Updates.sql.  This will update the is_parcel_merchant flag to '1' and also update their display name.  

**Please note:  The sql update uses npc ids.  If you have made changes or modified your npc ids, you may need to change/update.  

2. Opcode updates.  Depending on your installation, opcodes may require manual updating.  They are found in \assets\patches\patch_RoF2.conf within your installation.  The source opcodes are found at https://github.com/EQEmu/Server/tree/master/utils/patches  
   

## Defaults

There are several defaults as noted within the rules above.

- A player can only have 50 parcels at a time.
- A player can send a parcel once every 30 seconds
- There is a pruning system to automatically delete parcels after a specified period of days.  This is disabled by default.
- A player will recieve a message when a parcel is delivered, and on login.

## Admin Commands

There are several admin commands to assist GMs.

- \#parcel listdb [Character Name]  
  This allows the GM to query the character_parcels table to determine what parcels are available for the player identified.

- \#parcel listmemory [Character Name]  
  Similar to above, though queries the in-memory version of the player's parcels.  This is used for troubleshooting primarly.  Further, this command will only function if the Character is in the same zone as the GM.

- \#parcel add [Character Name] [item id] [quantity] [note]  
  Allows the GM to send an item to the Character Name.  The Character does not need to be online, or in the same zone.  
  Furthermore, the GM can send No Trade items.  In order to send money, please use an item id of 99990 and note that money is sent as an amount of copper.  
  Quantity is valid for stackable items, charges on an item, or amount of copper.

- \#parcels details [Character Name]  
  Allows for further querying of parcels assigned to a specific character.

  For Example:

- To send 12p5g3s1c to Player1  
  \#parcel add Player1 99990 12531 Here is your Money!

- To send a Fine Steel Jutte to Player3  
  \#parcel add Player3 6941 1 Enjoy your Jutte



# How it works

A player needs to find a merchant that notes 'Parcels and General Supplies'.  Right clicking on such a merchant provides a new merchant window, with a Parcels Tab.

https://youtu.be/3KB6W33BhE8

## Player Event Logging

The following player events are available.

|Rule Name|Default|Retention|
|:---|:---|:---|
|Parcel Send|Enabled|7 days|
|Parcel Retrieve|Enabled|7 days|
|Parcel Delete (Pruning System)|Enabled|7 days|

## Known Issues

There are currently no known issues.