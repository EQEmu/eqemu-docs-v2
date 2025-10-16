!!! info end

    This page lists the Bot commands that are available in-game, based on assigned Account Status, for your EQEmu Server.

## Description

The bot command system has been redesigned and is now a clone of the existing EQEmu command system.

Instead of using the old operator and command tokens #bot command, use the new operator with the new command tokens (i.e., ^command).

A redirect has also been added to the server command interpreter that will allow the use of #bot command..but, only the new command tokens may be used.

Many of the commands have been reviewed, re-coded and improved upon, where possible.

Since bots are an on-going project, some bot commands and features may be programmed into the server code..but, not enabled or fully-realized at this time.

As the actual bot code is updated, more of those commands and features can be enabled, as well as the addition of even more commands.

Please use the ^findaliases command to locate abridged versions of command names.

Using the 'help' or 'usage' argument after a command will show the proper formatting and available options for it.

## Actionable Bots

With specific exceptions, bot commands are designed to work using an 'actionable' argument.

Some of these actionable arguments also require the use of an 'actionable name' parameter.

The use an 'actionable' bot argument provides much greater flexibility and control over a bot or groups of bots and eliminates the overhead of programming multiple selection criteria into a command.

```text
target - selects target as single bot .. use ^command [target] or imply by empty actionable argument
byname [name] - selects single bot by name
ownergroup - selects all bots in the owner's group;
botgroup [name] - selects members of a bot-group by its name
targetgroup - selects all bots in target's group
namesgroup [name] - selects all bots in name's group
healrotation [name] - selects all member and target bots of a heal rotation where name is a member
healrotationmembers [name] - selects all member bots of a heal rotation where name is a member
healrotationtargets [name] - selects all target bots of a heal rotation where name is a member
spawned - selects all spawned bots
all - selects all spawned bots .. argument use indicates en masse database updating
```

Only those bots owned by the commanding player can be selected for any bot command use.

## Example

```text
Usage: (<friendly_target>) ^follow ([option: reset]) [actionable: byname | ownergroup | botgroup | namesgroup | healrotation | spawned] ([actionable_name])
```

* `^follow reset spawned` - resets all spawned bots to follow their default assignments
* `^follow byname Jojo` - Set the bot 'Jojo' to follow the selected friendly target
* `^follow botgroup Mybotgroup` - Sets all spawned members of the bot-group 'Mybotgroup' to follow the selected friendly target
* `^follow ownergroup` - Sets all bots within the owner's group to follow the selected friendly target

Target selection is optional if the argument preceding the command is enclosed in parentheses. In this case, an omitted actionable argument should default to the bot's owner.

If there is no argument preceding the command, then the selected target is not required, and hence, ignored.

Optional 'options' and 'actionable' arguments are also enclosed within parentheses.





| Command | Description | Status Level |
| :--- | :--- | :--- |
| ^actionable | Lists actionable command arguments and use descriptions | Player (0) |
| ^applypoison | Applies cursor-held poison to a rogue bot's weapon | Player (0) |
| ^attack | Orders bots to attack a designated target | Player (0) |
| ^behindmob | Toggles whether or not your bot tries to stay behind a mob | Player (0) |
| ^blockedbuffs | Set, view and clear blocked buffs for the selected bot(s) | Player (0) |
| ^blockedpetbuffs | Set, view and clear blocked pet buffs for the selected bot(s) | Player (0) |
| ^bot | Lists the available bot management [subcommands] | Player (0) |
| ^botappearance | Lists the available bot appearance [subcommands] | Player (0) |
| ^botbeardcolor | Changes the beard color of a bot | Player (0) |
| ^botbeardstyle | Changes the beard style of a bot | Player (0) |
| ^botcamp | Orders a bot(s) to camp | Player (0) |
| ^botclone | Creates a copy of a bot | GMMgmt (200) |
| ^botcreate | Creates a new bot | Player (0) |
| ^botdelete | Deletes all record of a bot | Player (0) |
| ^botdetails | Changes the Drakkin details of a bot | Player (0) |
| ^botdyearmor | Changes the color of a bot's (bots') armor | Player (0) |
| ^boteyes | Changes the eye colors of a bot | Player (0) |
| ^botface | Changes the facial appearance of your bot | Player (0) |
| ^botfollowdistance | Changes the follow distance(s) of a bot(s) | Player (0) |
| ^bothaircolor | Changes the hair color of a bot | Player (0) |
| ^bothairstyle | Changes the hairstyle of a bot | Player (0) |
| ^botheritage | Changes the Drakkin heritage of a bot | Player (0) |
| ^botinspectmessage | Changes the inspect message of a bot | Player (0) |
| ^botlist | Lists the bots that you own | Player (0) |
| ^botreport | Orders a bot to report its readiness | Player (0) |
| ^botsettings | Lists settings related to spell types and bot combat | Player (0) |
| ^botspawn | Spawns a created bot | Player (0) |
| ^botstance | Changes the stance of a bot | Player (0) |
| ^botstopmeleelevel | Sets the level a caster or spell-casting fighter bot will stop melee combat | Player (0) |
| ^botsuffix | Sets a bots suffix | Player (0) |
| ^botsummon | Summons bot(s) to your location | Player (0) |
| ^botsurname | Sets a bots surname (last name) | Player (0) |
| ^bottattoo | Changes the Drakkin tattoo of a bot | Player (0) |
| ^bottogglehelm | Toggles the helm visibility of a bot between shown and hidden | Player (0) |
| ^bottoggleranged | Toggles a ranged bot between melee and ranged weapon use | Player (0) |
| ^bottitle | Sets a bots title | Player (0) |
| ^botupdate | Updates a bot to reflect any level changes that you have experienced | Player (0) |
| ^botwoad | Changes the Barbarian woad of a bot | Player (0) |
| ^cast | Tells the first found specified bot to cast the given spell type | Player (0) |
| ^discipline | Uses aggressive/defensive disciplines or can specify spell ID | Player (0) |
| ^distanceranged | Controls the range casters and ranged will try to stay away from a mob | Player (0) |
| ^classracelist | Lists the classes and races and their appropriate IDs | Player (0) |
| ^clickitem | Orders your targeted bot to click the item in the provided inventory slot. | Player (0) |
| ^copysettings | Copies settings from one bot to another | Player (0) |
| ^defaultsettings | Restores a bot back to default settings | Player (0) |
| ^depart | Orders a bot to open a magical doorway to a specified destination | Player (0) |
| ^enforcespellsettings | Toggles your Bot to cast only spells in their spell settings list. | Player (0) |
| ^findaliases | Find available aliases for a bot command | Player (0) |
| ^follow | Orders bots to follow a designated target (option 'chain' auto-links eligible spawned bots) | Player (0) |
| ^guard | Orders bots to guard their current positions | Player (0) |
| ^healrotation | Lists the available bot heal rotation [subcommands] | Player (0) |
| ^healrotationadaptivetargeting | Enables or disables adaptive targeting within the heal rotation instance | Player (0) |
| ^healrotationaddmember | Adds a bot to a heal rotation instance | Player (0) |
| ^healrotationaddtarget | Adds target to a heal rotation instance | Player (0) |
| ^healrotationadjustcritical | Adjusts the critial HP limit of the heal rotation instance's Class Armor Type criteria | Player (0) |
| ^healrotationadjustsafe | Adjusts the safe HP limit of the heal rotation instance's Class Armor Type criteria | Player (0) |
| ^healrotationcastingoverride | Enables or disables casting overrides within the heal rotation instance | Player (0) |
| ^healrotationchangeinterval | Changes casting interval between members within the heal rotation instance | Player (0) |
| ^healrotationclearhot | Clears the HOT of a heal rotation instance | Player (0) |
| ^healrotationcleartargets | Removes all targets from a heal rotation instance | Player (0) |
| ^healrotationcreate | Creates a bot heal rotation instance and designates a leader | Player (0) |
| ^healrotationdelete | Deletes a bot heal rotation entry by leader | Player (0) |
| ^healrotationfastheals | Enables or disables fast heals within the heal rotation instance | Player (0) |
| ^healrotationlist | Reports heal rotation instance(s) information | Player (0) |
| ^healrotationremovemember | Removes a bot from a heal rotation instance | Player (0) |
| ^healrotationremovetarget | Removes target from a heal rotations instance | Player (0) |
| ^healrotationresetlimits | Resets all Class Armor Type HP limit criteria in a heal rotation to its default value | Player (0) |
| ^healrotationsave | Saves a bot heal rotation entry by leader | Player (0) |
| ^healrotationsethot | Sets the HOT in a heal rotation instance | Player (0) |
| ^healrotationstart | Starts a heal rotation | Player (0) |
| ^healrotationstop | Stops a heal rotation | Player (0) |
| ^help | List available commands and their description - specify partial command as argument to search | Player (0) |
| ^hold | Prevents a bot from attacking until released | Player (0) |
| ^illusionblock | Control whether or not illusion effects will land on the bot if casted by another player or bot | Player (0) |
| ^inventory | Lists the available bot inventory [subcommands] | Player (0) |
| ^inventorygive | Gives the item on your cursor to a bot | Player (0) |
| ^inventorylist | Lists all items in a bot's inventory | Player (0) |
| ^inventoryremove | Removes an item from a bot's inventory | Player (0) |
| ^inventorywindow | Displays all items in a bot's inventory in a pop-up window | Player (0) |
| ^itemuse | Elicits a report from spawned bots that can use the item on your cursor (option 'empty' yields only empty slots) | Player (0) |
| ^maxmeleerange | Toggles whether your bot is at max melee range or not. This will disable all special abilities, including taunt. | Player (0) |
| ^owneroption | Sets options available to bot owners | Player (0) |
| ^pet | Lists the available bot pet [subcommands] | Player (0) |
| ^petgetlost | Orders a bot to remove its summoned pet | Player (0) |
| ^petremove | Orders a bot to remove its charmed pet | Player (0) |
| ^petsettype | Orders a Magician bot to use a specified pet type | Player (0) |
| ^picklock | Orders a capable bot to pick the lock of the closest door | Player (0) |
| ^pickpocket | Orders a capable bot to pickpocket a NPC | Player (0) |
| ^precombat | Sets flag used to determine pre-combat behavior | Player (0) |>|_precombat) ||
| ^pull | Orders a designated bot to 'pull' an enemy | Player (0) |>|_pull) ||
| ^release | Releases a suspended bot's AI processing (with hate list wipe) | Player (0) |>|_release) ||
| ^setassistee | Sets your target to be the person your bots assist. Bots will always assist you before others | Player (0) |>|_set_assistee) ||
| ^sithppercent | HP threshold for a bot to start sitting in combat if allowed | Player (0) |>|_sit_hp_percent) ||
| ^sitincombat | Toggles whether or a not a bot will attempt to med or sit to heal in combat | Player (0) |>|_sit_in_combat) ||
| ^sitmanapercent | Mana threshold for a bot to start sitting in combat if allowed | Player (0) |>|_sit_mana_percent) ||
| ^spellaggrochecks | Toggles whether or not bots will cast a spell type if they think it will get them aggro | Player (0) |>|_spell_aggro_checks) ||
| ^spellannouncecasts | Turn on or off cast announcements by spell type | Player (0) |>|_spell_announce_cast) ||
| ^spellengagedpriority | Controls the order of casts by spell type when engaged in combat | Player (0) |>|_spell_engaged_priority) ||
| ^spelldelays | Controls the delay between casts for a specific spell type | Player (0) |>|_spell_delays) ||
| ^spellholds | Controls whether a bot holds the specified spell type or not | Player (0) |>|_spell_holds) ||
| ^spellidlepriority | Controls the order of casts by spell type when out of combat | Player (0) |>|_spell_idle_priority) ||
| ^spellmaxhppct | Controls at what HP percent a bot will start casting different spell types | Player (0) |>|_spell_max_hp_pct) ||
| ^spellmaxmanapct | Controls at what mana percent a bot will start casting different spell types | Player (0) |>|_spell_max_mana_pct) ||
| ^spellmaxthresholds | Controls the minimum target HP threshold for a spell to be cast for a specific type | Player (0) |>|_spell_max_thresholds) ||
| ^spellminhppct | Controls at what HP percent a bot will stop casting different spell types | Player (0) |>|_spell_min_hp_pct) ||
| ^spellminmanapct | Controls at what mana percent a bot will stop casting different spell types | Player (0) |>|_spell_min_mana_pct) ||
| ^spellminthresholds | Controls the maximum target HP threshold for a spell to be cast for a specific type | Player (0) |>|_spell_min_thresholds) ||
| ^spellresistlimits | Controls the resist limits for bots to cast spells on their target | Player (0) |>|_spell_resist_limits) ||
| ^spellpursuepriority | Controls the order of casts by spell type when pursuing in combat | Player (0) |>|_spell_pursue_priority) ||
| ^spelltargetcount | Sets the required target amount for group/AE spells by spell type | Player (0) |>|_spell_target_count) ||
| ^spellinfo | Opens a dialogue window with spell info | Player (0) |>|_dialogue_window) ||
| ^spells | Lists all Spells learned by the Bot. | Player (0) |>|_spell_list) ||
| ^spellsettings | Lists a bot's spell setting entries | Player (0) |>|_spell_settings_list) ||
| ^spellsettingsadd | Add a bot spell setting entry | Player (0) |>|_spell_settings_add) ||
| ^spellsettingsdelete | Delete a bot spell setting entry | Player (0) |>|_spell_settings_delete) ||
| ^spellsettingstoggle | Toggle a bot spell use | Player (0) |>|_spell_settings_toggle) ||
| ^spellsettingsupdate | Update a bot spell setting entry | Player (0) |>|_spell_settings_update) ||
| ^spelltypeids | Lists spelltypes by ID | Player (0) |>|_spelltype_ids) ||
| ^spelltypenames | Lists spelltypes by shortname | Player (0) |>|_spelltype_names) ||
| ^suspend | Suspends a bot's AI processing until released | Player (0) |>|_suspend) ||
| ^taunt | Toggles taunt use by a bot | Player (0) |>|_taunt) ||
| ^timer | Checks or clears timers of the chosen type. | GMMgmt (200) |>|_timer) ||
| ^track | Orders a capable bot to track enemies | Player (0) |>|_track) ||
| ^viewcombos | Views bot race class combinations | Player (0) |>|_view_combos)
