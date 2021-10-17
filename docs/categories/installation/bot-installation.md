# Bot Installation

## Description

The preferable method of installing bots is to use the automatic updater script to install the required schema and update for any required changes.

Sometimes, this does not seem to work for some people and the process will need to be performed manually.

These instructions assume a database free of any bots schema.

This will not work for Akkadius's repack!

## Requirements

* Git pull the latest code changes from the git repository: [https://github.com/EQEmu/Server/](https://github.com/EQEmu/Server/)
* Run cmake and delete the cache, then select configure
* Check the 'EQEMU\_ENABLE\_BOTS' option and select configure again
* Select generate, then exit cmake
* Open visual studio and compile all of the required source code
* Copy the new binaries to your server directory

## Recommended Actions

At this point, I would recommend manually running eqemu\_update.pl and update your standard database \(option 3\) before continuing.

Doing this should ensure that you have working system components.

## Installing the Base Schema

Locate the file **2015\_09\_30\_bots.sql** and apply it or copy the contents of this file into the query tab of a database management program, like HeidiSQL, and run the query.

```text
https://github.com/EQEmu/Server/blob/master/utils/sql/git/bots/required/2015_09_30_bots.sql
```

This should apply the 'base' schema required for bots.

## Modifying Existing Database

Run each of the following queries by itself to modify the existing database.

```text
UPDATE `spawn2` SET `enabled` = 1 WHERE `id` IN (59297,59298);
```

```text
ALTER TABLE `guild_members` DROP PRIMARY KEY;
```

```text
ALTER TABLE `group_id` DROP PRIMARY KEY;
```

```text
ALTER TABLE `group_id` ADD PRIMARY KEY USING BTREE(`groupid`, `charid`, `name`, `ismerc`);
```

```text
INSERT INTO `command_settings` VALUES ('bot', '0', '');
```

```text
INSERT INTO `rule_values` VALUES
('1', 'Bots:AAExpansion', '8', 'The expansion through which bots will obtain AAs'),
('1', 'Bots:CreationLimit', '150', 'Number of bots that each account can create'),
('1', 'Bots:FinishBuffing', 'false', 'Allow for buffs to complete even if the bot caster is out of mana.  Only affects buffing out of combat.'),
('1', 'Bots:GroupBuffing', 'false', 'Bots will cast single target buffs as group buffs, default is false for single. Does not make single target buffs work for MGB.'),
('1', 'Bots:ManaRegen', '3.0', 'Adjust mana regen for bots, 1 is fast and higher numbers slow it down 3 is about the same as players.'),
('1', 'Bots:QuestableSpawnLimit', 'false', 'Optional quest method to manage bot spawn limits using the quest_globals name bot_spawn_limit, see: /bazaar/Aediles_Thrall.pl'),
('1', 'Bots:QuestableSpells', 'false', 'Anita Thrall\'s (Anita_Thrall.pl) Bot Spell Scriber quests.'),
('1', 'Bots:SpawnLimit', '71', 'Number of bots a character can have spawned at one time, You + 71 bots is a 12 group raid');
```

## Adding and Updating \`db\_version\`.\`bots\_version\`

These must be added for the versioning system to work.

```text
ALTER TABLE `db_version` ADD `bots_version` int(11) DEFAULT '0' AFTER `version`;
```

```text
UPDATE `db_version` SET `bots_version` = '9000';
```

## Applying Version Updates to Your Bots Database

I recommend letting the automatic updater take control at this point.

If you choose to apply all updates manually, you will need to check the manifest file to find the version number to update to: [https://github.com/EQEmu/Server/blob/master/utils/sql/git/bots/bots\_db\_update\_manifest.txt](https://github.com/EQEmu/Server/blob/master/utils/sql/git/bots/bots_db_update_manifest.txt)

\(You must apply the updates in the order that they appear in the manifest or database corruption could ensue.\)

## **Automatic Updater**

* Again, manually run eqemu\_update.pl to access the options menu
* Select option '4' to queue pending bots database updates
* Select option '4' again to apply these updates

With any luck, you should have a working BOTS server!

