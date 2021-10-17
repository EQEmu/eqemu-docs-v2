# Loading Server Data

This page explains how different assets on the server are loaded and by what methods can they be refreshed and/or reloaded

## Alternate Currency Data

* In game command **\#reloadstatic** - will reload alternate currency data from the `alternate_currency` table for just the zone the command was executed in

## Base Data

* Base Data uses shared memory - you most likely aren't editing this
* Base Data is recommended to be ran from server boot-up for drastic changes, however editing existing data can be reloaded server-wide using **\#hotfix**

## Blocked Spells

* Blocked spells data is loaded from the `blocked_spells` table on zone init
* There is no current way to hot reload this data

## Doors

* In game command **\#reloadstatic** - will reload doors from the `doors` table for the respective zone and instance

## Factions

* Factions uses shared memory
* Factions can be hot reloaded in game using **\#hotfix** - keep in mind that this uses shared memory files produced from **shared\_memory** binary located in the **./shared** folder
* New factions need a server reboot, existing factions can use **\#hotfix**, you can use placeholder data to add new factions and safely reload like other shared memory data

## Fishing

* Fishing data is live once entered into the database, no reloading is required

## Foraging

* Forage data is live once entered into the database, no reloading is required

## Grids / Pathing Data

* All grid data is loaded at zone boot-up and any new data is simply reloaded during a **\#repop**

## Ground Spawns

* In game command **\#reloadstatic** - will reload ground spawns from the `ground_spawns` table for the respective zone and instance

## Horses

* Horse data is queried immediately from the `horses` table on request of creating horse from Spell cast

## Items

* Items uses shared memory
* Items can be hot reloaded in game using **\#hotfix** - keep in mind that this uses shared memory files produced from **shared\_memory** binary located in the **./shared** folder
* Items in earlier clients such as Titanium show the change immediately, if you have inspected an item on a later client, then issued a **\#hotfix**, the client caches these results so you will need to camp or zone to see the stat changes however any affects should take affect immediately
* Note: If you are going to build new items on your server on the fly, I recommend creating a big bank of blank or placeholder ID's in your table that you can use in the future. You can't hot reload the server with new items without creating issues after creating a new row entry
* Existing items can be safely edited without a server reload and using the **\#hotfix** command. This takes affect for all zones immediately

## Level EXP Mods

* If you are using the `level_exp_mods` table for customizing percentage difficulty modifiers - these can be hot reloaded using in game command **\#reloadlevelmods**

## Loot

* Loot uses shared memory
* Loot can be hot reloaded in game using **\#hotfix** - keep in mind that this uses shared memory files produced from **shared\_memory** binary located in the **./shared** folder
* Note: Loot assigned in a script/quest does not rely on the database system or to be reloaded from shared memory

## Logging

* Server logging settings can be reloaded in game using **\#logs reload\_all** from the `logsys_categories` table - this takes affect server wide for all processes

## Merchants

* Merchant data is loaded and cached the first time the request is made to a merchant if it wasn't already loaded on zone bootup
* Merchants can be hot reloaded using **\#reloadmerchants**

## NPC Data

* NPC data is live once a **\#repop** is issued after a new edit or creation

## NPC Emotes

* While emotes don't HAVE to be database driven \(most custom servers will just use scripts\) - there is an option to reload the database driven emotes
* In game command **\#reloadstatic** - will reload database driven emote data from the `npc_emotes` table for just the zone the command was executed in

## Objects

* In game command **\#reloadstatic** - will reload objects from the `objects` table for the respective zone and instance

{% hint style="success" %}
Note that you will likely have to zone for your client to update.
{% endhint %}

## Pets

* Pet data is live and usable the moment it is edited in the `pets` table and an `npc_types` entry was already created. Otherwise a **\#repop** will be require to reload the npc\_types data

## Perl Event Exports

* In game command **\#reloadperlexportsettings** will reload all export settings from the `perl_event_export_settings` table and this will take affect for **all current running zones**

## Quests

* In game command **\#rq** or **\#reloadquest** will reload all quest scripts for the zone you are in, LUA or Perl
* Server-wide in game command **\#reloadworld** will perform a quest reload for all zones, however keep in mind this does a repop as well

## Rules

* In game command **\#reloadallrules** will reload rules from the `rule_values` table for all running zone processes and the world process

## Skill Caps

* Skill Caps uses shared memory \(Class skills, caps etc.\)
* Skill caps is recommended to be ran from server boot-up for drastic changes, however editing existing data can be reloaded server-wide using **\#hotfix**

## Spells

* Spells uses shared memory
* Spells can be hot reloaded in game using **\#hotfix** - keep in mind that this uses shared memory files produced from **shared\_memory** binary located in the **./shared** folder
* Note: If you are going to build new spells on your server on the fly, I recommend creating a big bank of blank or placeholder ID's in your table that you can use in the future. You can't hot reload the server with new spells without creating issues after creating a new row entry.
* Existing spells can be safely edited without a server reload and using the **\#hotfix** command. This takes affect for all zones immediately

## Tasks

* All new Task data requires **\#task reloadall** to be executed to reload data from the `tasks` table and the `activities` table.
* Note: If a character has a task that you are testing/building and you add new steps, the server will remove it from your task window

## Traps

* In game command **\#reloadstatic** - will reload traps from the `traps` table for the respective zone and instance
* In game comand **\#reloadtraps** will also reload traps without reloading everything else that **\#reloadstatic** does

## Tradeskills

* Tradeskills query the database direct and do not require reload to take affect

## Zone Points

* In game command **\#reloadstatic** - will zone points from the `zone_points` table for the respective zone and instance - note this may not work for all clients
* In game comand **\#reloadzps** will also reload zone points without reloading everything else that **\#reloadstatic** does

## Zone Data

* Sometimes you want to reload the `zone` table data, this is done via using \#zheader \[zoneshortname\]. If you're in Nexus, use **\#zheader nexus**
* This will resend the zone header packet to the client to update fog / sky / gravity / data / etc.

