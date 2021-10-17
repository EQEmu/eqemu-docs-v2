# Bot Groups

## Description 

Bot groups are a **special derivation of player groups**.

A **bot owner** may create a **special group** of **player-owned bots** that are available to employ through the use of the **bot command system**.

## Bot-Group Commands

```text
^botgroupaddmember - Adds a member to a bot-group
^botgroupcreate - Creates a bot-group and designates a leader
^botgroupdelete - Deletes a bot-group and releases its members
^botgrouplist - Lists all of your existing bot-groups
^botgroupload - Loads all members of a bot-group
^botgroupremovemember - Removes a bot from its bot-group
```

Bot-group commands **only control** features of the **bot-group itself**. **Control** of the bots is accomplished through the use of[ Bot Commands](http://wiki.eqemulator.org/p?Bot_Commands&frm=Bot_Groups) .

## Creating a Bot-Group

```text
usage: (<target_leader>) ^botgroupcreate [group_name] ([leader_name])
```

* ^botgroupcreate Mybotgroup - Creates a bot-group named 'Mybotgroup' with the selected bot as its leader
* ^botgroupcreate Mybotgroup Jojo - Creates a bot-group named 'Mybotgroup' with the bot 'Jojo' as its leader

Upon creation of the bot-group, **a group instance is created** on the server and the **bot-group is created** with its leader **added as a member to the database**.

## Adding a Member to a Bot-Group

```text
usage: (<target_leader>) ^botgroupaddmember [member_name] ([leader_name])
```

* ^botgroupaddmember Bozo - Adds bot 'Bozo' to the bot-group owned by the selected bot
* ^botgroupaddmember Bozo Jojo - Adds bot 'Bozo' to the bot-group owned by the bot 'Jojo'

Adding a new member updates both the **group instance** and the **database bot-group table**.

## Removing a Member from a Bot-Group

```text
usage: (<target_member>) ^botgroupremovemember [member_name]
```

* ^botgroupremovemember - Removes the selected bot from its bot-group
* ^botgroupremovemember Bozo - Removes the bot 'Bozo' from its bot-group

Removing a member updates both the **group instance** and the **database bot-group table**.

In the case of a removed bot being the **leader of a bot-group**, **all spawned members** of the bot-group are **disbanded from their current group** and the **database entries are removed** - see "Deleting a Bot-Group"

## Deleting a Bot-Group

```text
usage: ^botgroupdelete [group_name]
```

* ^botgroupdelete Mybotgroup - Disbands all members of the bot-group named 'Mybotgroup' and removes all database entries

**Removing the leader of a bot-group** performs the same action as **using the botgroupdelete** command.

## Loading a Bot-Group

```text
usage: ^botgroupload [group_name]
```

* ^botgroupload Mybotgroup - Spawns all members of the bot-group named 'Mybotgroup'

## Listing Bot-Groups

```text
usage: ^botgrouplist
```

* ^botgrouplist - Lists all of the created bot-groups that you own by name

