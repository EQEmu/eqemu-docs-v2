# Loading Server Data

This page explains how different assets on the server are loaded and by what methods can they be refreshed and/or reloaded

## Alternate Currency Data

* **#reload static** will reload Alternate Currency data from the `alternate_currency` table globally

## Base Data

* Base Data uses shared memory - you most likely aren't editing this
* Base Data is recommended to be ran from server boot-up for drastic changes, however editing existing data can be reloaded server-wide using **#hotfix**

## Blocked Spells

* Blocked spells data is loaded from the `blocked_spells` table on zone init
* There is no current way to hot reload this data

## Content Flags

* **#reload content_flags** will reload Content Flags data from the `content_flags` table globally

## Doors

* **#reload static** will reload Doors data from the `doors` table globally

## Factions

* Factions uses shared memory
* Factions can be hot reloaded in game using **#hotfix** - keep in mind that this uses shared memory files produced from **shared_memory** binary located in the **./shared** folder
* New factions need a server reboot, existing factions can use **#hotfix**, you can use placeholder data to add new factions and safely reload like other shared memory data

## Fishing and Foraging

* Fishing and Foraging data are live once entered into the database, no reloading is required

## Grids and Pathing Data

* All grid data is loaded at zone boot-up and any new data is simply reloaded during a **#repop**

## Ground Spawns

* **#reload static** will reload ground spawns from the `ground_spawns` table globally

## Horses

* Horse data is queried direclty from the `horses` table on request of creating horse from Spell cast

## Items

* Items uses shared memory
* Items can be hot reloaded in game using **#hotfix** - keep in mind that this uses shared memory files produced from **shared_memory** binary located in the **./shared** folder
* Items in earlier clients such as Titanium show the change immediately, if you have inspected an item on a later client, then issued a **#hotfix**, the client caches these results so you will need to camp or zone to see the stat changes however any affects should take affect immediately
* Note: If you are going to build new items on your server on the fly, I recommend creating a big bank of blank or placeholder ID's in your table that you can use in the future. You can't hot reload the server with new items without creating issues after creating a new row entry
* Existing items can be safely edited without a server reload and using the **#hotfix** command. This takes affect for all zones immediately

## Level EXP Mods

* **#reload level_mods** will reload Level Based Experience Modifiers data from the `level_exp_mods` table globally
* Note: Must be enabled with the `Zone:LevelBasedEXPMods` rule to be used.

## Loot

* Loot uses shared memory
* Loot can be hot reloaded in game using **#hotfix** - keep in mind that this uses shared memory files produced from **shared_memory** binary located in the **./shared** folder
* Note: Loot assigned in a script/quest does not rely on the database system or to be reloaded from shared memory

## Logging

* **#logs reload_all** or **#reload logs** will reload Log Settings from the `logsys_categories` table globally

## Merchants

* **#reload merchants** will reload Merchants data from the `merchantlist` table globally.
* Merchant data is loaded and cached the first time the request is made to a merchant if it wasn't already loaded on zone bootup

## NPC Data

* NPC data is live once a **#repop** is issued after a new edit or creation

## NPC Emotes

* **#reload npc_emotes** will reload NPC Emotes from the `npc_emotes` table globally.
* While emotes don't HAVE to be database driven (most custom servers will just use scripts) - there is an option to reload the database driven emotes

## Objects

* **#reload static** will reload Objects data from the `object` table globally

!!! info
      Note that you will likely have to zone for your client to update.

## Pets

* Pet data is live and usable the moment it is edited in the `pets` table and an `npc_types` entry was already created. Otherwise a **#repop** will be require to reload the npc_types data

## Perl Event Exports

* **#reload perl_export** will reload Perl Event Export Settings data from the `perl_event_export_settings` table globally

## Quests

* **#reload quest** or **#reload quest 0** will reload all quest scripts for the zone you are in.
    - **#reload quest 1** will reload all quests for the zone you're in and stop quest timers.
* **#reload world** or **#reload world 0** will reload Quests globally.
    - **#reload world 1** will reload Quests and repop globally.
    - **#reload world 2** will reload Quests and forcefully repop globally.

## Rules

* **#reload rules** will reload Rules data from the `rule_values` table globally

## Skill Caps

* Skill Caps uses shared memory (Class skills, caps etc.)
* Skill caps is recommended to be ran from server boot-up for drastic changes, however editing existing data can be reloaded server-wide using **#hotfix**

## Spells

* Spells uses shared memory
* Spells can be hot reloaded in game using **#hotfix** - keep in mind that this uses shared memory files produced from **shared_memory** binary located in the **./shared** folder
* Note: If you are going to build new spells on your server on the fly, I recommend creating a big bank of blank or placeholder ID's in your table that you can use in the future. You can't hot reload the server with new spells without creating issues after creating a new row entry.
* Existing spells can be safely edited without a server reload and using the **#hotfix** command. This takes affect for all zones immediately

## Tasks

* **#task reloadall** or **#reload tasks** will reload Tasks data from the `tasks` table and the `activities` table globally
* Note: If a character has a task that you are testing/building and you add new steps, the server will remove it from your task window

## Titles
* **#reload titles** will reload Titles data from `titles` table globally

## Traps

* **#reload static** will reload Traps data from the `traps` table globally
* **#reload traps 1** will also reload traps globally without reloading everything else that **#reload static** does

## Tradeskills

* Tradeskills query the database directly and do not require reload to take affect

## Variables

* **#reload variables** will reload Variables data from the `variables` table globally

## Zone Points

* **#reload static** will reload Zone Points data from the `zone_points` table globally
* **#reload zone_points** will also reload Zone Points data globally without reloading everything else that **#reload static** does

## Zone Data

* Sometimes you want to reload the `zone` table data, this is done via using #zheader [zoneshortname]. If you're in Nexus, use **#zheader nexus**
* This will resend the zone header packet to the client to update fog / sky / gravity / data / etc.

## Reload Static Command

* **#reload static** reloads the following globally.
    - Alternate Currencies
    - Doors
    - Objects
    - Ground Spawns
    - NPC Emotes
    - Traps
    - Veteran Rewards
    - Zone Configurations
    - Zone Points

