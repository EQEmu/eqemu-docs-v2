# Bot Commands

## Description

The **bot command system** has been redesigned and is now a clone of the existing **EQEmu command system**.

Instead of using the **old operator** and **command tokens** \#bot command, use the **new operator** with the **new command tokens** \(i.e., ^command\).

A **redirect** has also been added to the server command interpreter that will allow the use of \#bot command..but, only the **new command tokens** may be used.

Many of the **commands** have been **reviewed**, **re-coded** and **improved upon**, where possible.

Since bots are an **on-going project**, some bot commands and features **may be programmed** into the server code..but, **not enabled** or **fully-realized** at this time.

As the actual bot code is **updated**, more of those commands and features can be **enabled**, as well as the addition of even **more commands**.

Please use the ^findaliases command to locate abridged versions of command names.

## Bot Commands

```text
^actionable - Lists actionable command arguments and use descriptions
^aggressive - Orders a bot to use a aggressive discipline
^applypoison – Orders a level 18 or greater rogue bot to apply poison to its weapon
^applypotion – Order a bot to consume a potion
^attack - Orders bots to attack a designated target
^bindaffinity - Orders a bot to attempt an affinity binding
^bot - Lists the available bot management [subcommands]
^botappearance - Lists the available bot appearance [subcommands]
^botbeardcolor - Changes the beard color of a bot
^botbeardstyle - Changes the beard style of a bot
^botcamp - Orders a bot(s) to camp
^botclone - Creates a copy of a bot
^botcreate - Creates a new bot
^botdelete - Deletes all record of a bot
^botdetails - Changes the Drakkin details of a bot
^botdyearmor - Changes the color of a bot's (bots') armor
^boteyes - Changes the eye colors of a bot
^botface - Changes the facial appearance of your bot
^botfollowdistance - Changes the follow distance(s) of a bot(s)
^botgroup - Lists the available bot-group [subcommands]
^botgroupaddmember - Adds a member to a bot-group
^botgroupcreate - Creates a bot-group and designates a leader
^botgroupdelete - Deletes a bot-group and releases its members
^botgrouplist - Lists all of your existing bot-groups
^botgroupload - Loads all members of a bot-group
^botgroupremovemember - Removes a bot from its bot-group
^bothaircolor - Changes the hair color of a bot
^bothairstyle - Changes the hairstyle of a bot
^botheritage - Changes the Drakkin heritage of a bot
^botinspectmessage - Changes the inspect message of a bot
^botlist - Lists the bots that you own
^botoutofcombat - Toggles your bot between standard and out-of-combat spell/skill use - if any specialized behaviors exist
^botreport - Orders a bot to report its readiness
^botspawn - Spawns a created bot
^botstance - Changes the stance of a bot
^botstopmeleelevel – Sets a bot’s level to stop melee
^botsuffix – Sets a bot’s name suffix
^botsummon - Summons bot(s) to your location
^botsurname – Sets a bot’s surname
^bottattoo - Changes the Drakkin tattoo of a bot
^bottitle – Sets a bot’s title
^bottogglearcher - Toggles a archer bot between melee and ranged weapon use
^bottogglehelm - Toggles the helm visibility of a bot between shown and hidden
^botupdate - Updates a bot to reflect any level changes that you have experienced
^botwoad - Changes the Barbarian woad of a bot
^charm - Attempts to have a bot charm your target
^circle - Orders a Druid bot to open a magical doorway to a specified destination
^cure - Orders a bot to remove any ailments
^defensive - Orders a bot to use a defensive discipline
^depart - Orders a bot to open a magical doorway to a specified destination
^escape - Orders a bot to send a target group to a safe location within the zone
^evacuate – Order a capable bot to evacuate group – NOT YET FULLY IMPLEMENTED
^findaliases - Find available aliases for a bot command
^follow - Orders bots to follow a designated target
^guard - Orders bots to guard their current positions
^healrotation - Lists the available bot heal rotation [subcommands]
^healrotationadaptivetargeting - Enables or disables adaptive targeting within the heal rotation instance
^healrotationaddmember - Adds a bot to a heal rotation instance
^healrotationaddtarget - Adds target to a heal rotation instance
^healrotationadjustcritical - Adjusts the critial HP limit of the heal rotation instance's Class Armor Type criteria
^healrotationadjustsafe - Adjusts the safe HP limit of the heal rotation instance's Class Armor Type criteria
^healrotationcastingoverride - Enables or disables casting overrides within the heal rotation instance
^healrotationchangeinterval - Changes casting interval between members within the heal rotation instance
^healrotationclearhot - Removes all heal over time instances from a heal rotation instance
^healrotationcleartargets - Removes all targets from a heal rotation instance
^healrotationcreate - Creates a bot heal rotation instance and designates a leader
^healrotationdelete – Deletes a bot heal rotation instance
^healrotationfastheals - Enables or disables fast heals within the heal rotation instance
^healrotationlist - Reports heal rotation instance(s) information
^healrotationremovemember - Removes a bot from a heal rotation instance
^healrotationremovetarget - Removes target from a heal rotations instance
^healrotationresetlimits - Resets all Class Armor Type HP limit criteria in a heal rotation to its default value
^healrotationsave - Saves a heal rotation
^healrotationsethot - Sets a heal rotation members heal over time spell
^healrotationstart - Starts a heal rotation
^healrotationstop - Stops a heal rotation
^help - List available commands and their description - specify partial command as argument to search
^hold - Suspends a bot's AI processing until released
^identify - Orders a bot to cast an item identification spell
^inventory - Lists the available bot inventory [subcommands]
^inventorygive - Gives the item on your cursor to a bot
^inventorylist - Lists all items in a bot's inventory
^inventoryremove - Removes an item from a bot's inventory
^inventorywindow – Displays a bot’s inventory
^invisibility - Orders a bot to cast a cloak of invisibility, or allow them to be seen
^itemuse – Generates list of spawned bot’s who can use item on player’s cursor
^levitation - Orders a bot to cast a levitation spell
^lull - Orders a bot to cast a pacification spell
^mesmerize - Orders a bot to cast a mesmerization spell
^movementspeed - Orders a bot to cast a movement speed enhancement spell
^owneroption – Generates list of options owner can set for targeted bot
^pet - Lists the available bot pet [subcommands]
^petgetlost – Dismisses a bot’s pet
^petremove - Orders a bot to remove its pet
^petsettype - Orders a Magician bot to use a specified pet type
^picklock - Orders a capable bot to pick the lock of the closest door
^precombat – Sets bots pre-combat on/off to heal/cure/buff prior to engagement
^portal - Orders a Wizard bot to open a magical doorway to a specified destination
^pull - Orders a designated bot to 'pull' an enemy
^release - Releases a suspended bot's AI processing (with hate list wipe)
^resistance - Orders a bot to cast a specified resistance buff
^resurrect - Orders a bot to resurrect a player's (players') corpse(s)
^rune - Orders a bot to cast a rune of protection
^sendhome - Orders a bot to open a magical doorway home
^size - Orders a bot to change a player's size
^succor – Orders a capable bot to cast succor - NOT YET FULLY IMPLEMENTED
^summoncorpse - Orders a bot to summon a corpse to its feet
^suspend – Pauses a bot’s AI processing
^taunt - Toggles taunt use by a bot
^track - Orders a capable bot to track enemies
^waterbreathing - Orders a bot to cast a water breathing spell
```

Using the '**help**' or '**usage**' argument after a command will show the **proper formatting** and **available options** for it.

## Actionable Bots

With specific exception, bot commands are designed to work using an '**actionable**' argument.

Some of these actionable arguments also require the use of an '**actionable name**' parameter.

The use an 'actionable' bot argument provides much **greater flexibility** and **control** over a bot or groups of bots and **eliminates the overhead of programming** multiple selection criteria into a command.

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

Only those **bots owned by the commanding player** can be **selected** for any **bot command use**.

## Example

```text
usage: (<friendly_target>) ^follow ([option: reset]) [actionable: byname | ownergroup | botgroup | namesgroup | healrotation | spawned] ([actionable_name])
```

* ^follow reset spawned - resets all spawned bots to follow their default assignments
* ^follow byname Jojo - Set the bot 'Jojo' to follow the selected friendly target
* ^follow botgroup Mybotgroup - Sets all spawned members of the bot-group 'Mybotgroup' to follow the selected friendly target
* ^follow ownergroup - Sets all bots within the owner's group to follow the selected friendly target

**Target selection is optional** if the argument preceding the command is enclosed in parentheses. In this case, an **omitted actionable argument** should default to the **bot's owner**.

If there is **no argument** preceding the command, then the selected target is **not required**, and hence, ignored.

Optional '**options**' and '**actionable**' arguments are also **enclosed within parentheses**.

