---
description: This page further details the use of the EQEmu Task System.
---

# Task System

The EQEmu Task System allows implementation of quests that utilize the clients Task Window (alt+Q).  Many aspects of this system can be found throughout this wiki; this guide attempts to organize that information on a single page to inform the reader who wishes to understand and utilize the Task System.

Much of the Task System functionality can be found in [/zones/task.cpp](https://github.com/EQEmu/Server/blob/master/zone/tasks.cpp) and should be referenced by those comfortable with looking at the EQEmu source code.

{% hint style="info" %}
While every attempt will be made to keep this page updated, please reference the additional supporting documentation found in the [Database Schema](https://eqemu.gitbook.io/database-schema/), [Quest API](https://eqemu.gitbook.io/quest-api/), and [Changelog](https://eqemu.gitbook.io/changelog/).
{% endhint %}

| Section Description                                                     | Link                                                                                                                  |
| ----------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| Database Schema:  All of the DB tables for the task system              | [Database Schema](https://eqemu.gitbook.io/server/categories/how-to-guides/task-system-guide#database-schema)         |
| GM Commands:  in-game GM commands for the task system                   | [GM Commands](https://eqemu.gitbook.io/server/categories/how-to-guides/task-system-guide#gm-commands)                 |
| Task System Rules:  server rules effecting the task system              | [Task System Rules](https://eqemu.gitbook.io/server/categories/how-to-guides/task-system-guide#task-system-rules)     |
| Logging Options: use with the Logging System                            | [Logging Options](https://eqemu.gitbook.io/server/categories/how-to-guides/task-system-guide#logging-options)         |
| Quest API: quest methods and EVENT triggers relating to the task system | [Quest API](https://eqemu.gitbook.io/server/categories/how-to-guides/task-system-guide#quest-api-for-the-task-system) |

## Database Schema

The most up to date information regarding the database schema can be found in the Database Schema Space for [Tasks](https://eqemu.gitbook.io/database-schema/categories/tasks).  The information presented here is current as of this writing.

### Tasks Table

#### Field Descriptions

| Column           | Data Type | Description                                                                                                                                  |
| ---------------- | --------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| id               | int       | Unique Task Identifier; use 1 - 9999, a value of 0 should NOT be used                                                                        |
| type             | tinyint   | [Task Type](https://eqemu.gitbook.io/server/categories/types/task-types); 0=Task, 1=Shared, 2=Quest, 3=Expedition                            |
| duration         | int       | Duration in seconds; if 0, the task has no time limit                                                                                        |
| duration_code    | tinyint   | [Duration Code](https://eqemu.gitbook.io/server/categories/reference-lists/task-duration-codes); 0=None, 1=Short, 2=Medium, 3=Long           |
| title            | varchar   | Title                                                                                                                                        |
| description      | text      | Description; only the active task description displays                                                                                       |
| reward           | varchar   | Reward Description                                                                                                                           |
| rewardid         | int       | Item ID number, or a reference to a list of item ID numbers (from goallists table)                                                           |
| cashreward       | int       | Cash Reward, listed in copper (system will convert to plat, gold, etc.)                                                                      |
| xpreward         | int       | Experience Reward; for level-based, set a negative value using the following formula:  max level \* 100 + experience percent                 |
| rewardmethod     | tinyint   | Reward Method: 0 = Single Item ID, 1 = List of Items (in the goallist table), 2 = Quest Controlled (reward granted through NPC quest script) |
| minlevel         | tinyint   | Minimum Level to obtain tasks                                                                                                                |
| maxlevel         | tinyint   | Maximum Level to obtain tasks                                                                                                                |
| repeatable       | tinyint   | Repeatable: 0 = False, 1 = True                                                                                                              |
| faction_reward   | int       | Faction Reward                                                                                                                               |
| completion_emote | varchar   | Completion Emote                                                                                                                             |

### Task_Activities Table

#### Field Descriptions

| Column               | Data Type | Description                                                                                                                                                                                                                                                    |
| -------------------- | --------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| taskid               | int       | Task Identifier (ID from Tasks table); each activity for a task will utilize the same ID.                                                                                                                                                                      |
| activityid           | int       | Activity Identifier: Starts at 0, **must be sequential**                                                                                                                                                                                                       |
| step                 | int       | Step: 0 = Always Available, >0 = Must Complete Previous ("stepped" task)                                                                                                                                                                                       |
| activitytype         | tinyint   | [Activity Type](https://eqemu.gitbook.io/server/categories/types/task-activity-types) (or see below)                                                                                                                                                           |
| target_name          | varchar   | Target Name                                                                                                                                                                                                                                                    |
| item_list            | varchar   | Item Identifier List                                                                                                                                                                                                                                           |
| skill_list           | varchar   | Skill Identifier List                                                                                                                                                                                                                                          |
| spell_list           | varchar   | Spell Identifier List                                                                                                                                                                                                                                          |
| description_override | varchar   | Description Override                                                                                                                                                                                                                                           |
| goalid               | int       | Goal Identifier or Goal List Identifier                                                                                                                                                                                                                        |
| goalmethod           | int       | <p>Goal Method: </p><p>0 = Single Value (npc_type ID or item ID)</p><p>1 = List (refers to the goallist table to use a list of npc_type IDs or Item IDs)</p><p>2 = Under control of quest system</p>                                                           |
| goalcount            | int       | Goal Count (how many of the things in the goal id must be done)                                                                                                                                                                                                |
| delivertonpc         | int       | Deliver To NPC: 0 = No Delivery NPC, >0 = npc_type ID                                                                                                                                                                                                          |
| zones                | varchar   | [Zones List](https://eqemu.gitbook.io/server/categories/reference-lists/zones)--use the zoneidnumber field; "0" displays "ALL".  Activities will only be counted if completed in the assigned zone. Can also be used as "Touch" for visiting an assigned zone. |
| optional             | tinyint   | Optional: 0 = False, 1 = True; if a task activity is optional, the task will be complete when all non-optional task activities are completed.                                                                                                                  |

#### Activity Types

{% hint style="info" %}
Please note that **quest::updatetaskactivity **could be used in many EVENT triggers if you did not want to use the task system as described in the table below.
{% endhint %}

| Type ID | <p>Type <br>Name</p> | Information and usage                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| ------- | -------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1       | Deliver              | <p>To have this activity automatically handled by the task system without the need to write any Perl quest code, set:</p><ul><li>goalmethod 0</li><li>goalid to the id of the item you want delivered</li><li>delivertonpc to the npc_types id you want the item delivered to</li><li>zoneid to the zoneidnumber the npc is in (or 0 for all)</li></ul><p>If you wish to allow the player to deliver any out of a group of items, then you can use a list. To do this, create a list in the goallist table containing the allowable items. Then set the goalmethod to 1 and put the list number in the goalid field.</p>                                                                           |
| 2       | Kill                 | <p>To have this activity automatically handled by the task system without the need to write any Perl quest code, set:</p><ul><li>goalmethod to 0</li><li>goalid to the id of the npc that must be killed (from the npc_types table).</li><li>zoneid should be set to 0 (any zone), or the zoneidnumber if the mobs must be killed in a particular zone to count.</li></ul><p>If you wish to allow the player to kill any out of a group of NPC types, then you can use a list as described for the Deliver activity. For common mobs, you will probably always be using lists, since there are a lot of entries for mobs with the same name in the npc_types table which differ by level, etc.</p> |
| 3       | Loot                 | <p>To have this activity automatically handled by the task system without the need to write any Perl quest code, set:</p><ul><li>goalmethod to 0</li><li>goalid to the id of the item that must be looted (from the items table).</li><li>zoneid should be set to 0 (any zone), or the zoneidnumber if the item must be looted in a particular zone to count.</li></ul><p>If you wish to allow the player to loot any one of a group of items, then you can use a list as described for the Deliver activity.</p>                                                                                                                                                                                  |
| 4       | Speak With           | <p>To have this activity automatically handled by the task system without the need to write any Perl quest code, set:</p><ul><li>goalmethod to 0</li><li>goalid to the id of the NPC that must be spoken to (from the npc_types table). </li><li>zoneid should be set to 0 (any zone), or the zoneidnumber that the NPC is in.</li></ul><p>If you wish to allow the player to speak to any one of a group of NPCs, then you can use a list as described for the Deliver activity.</p>                                                                                                                                                                                                              |
| 5       | Explore              | <p>To have this activity automatically handled by the task system without the need to write any Perl quest code, you will need to add a row to the proximity table. Once you have added the row and given it a unique exploreid number, in the activity table, set:</p><ul><li>goalmethod to 0</li><li>goalid to the exploreid of your newly created row in the proximity table.</li><li>zoneid should be set to the zoneidnumber that the proximity is in.</li></ul><p>It should be possible to set up multiple proximities in the same zone, using lists, such that any of the proximities will trigger the activity complete.</p>                                                               |
| 6       | Tradeskill           | <p>To have this activity automatically handled by the task system, set:</p><ul><li>goalmethod to 0</li><li>goalid to the id of the item that must be created (from the items table).</li><li>zoneid should be set to 0 (any zone), or the zoneidnumber that the items must be created in.</li></ul><p>If you wish to allow the player to create any one of a group of items, then you can use a list as described for the Deliver activity. Sometimes there are several items with the same name that can be created from a tradeskill, CLASS 1 Wood Point arrows being one of them.</p>                                                                                                           |
| 7       | Fish                 | <p>To have this activity automatically handled by the task system, set:</p><ul><li>goalmethod to 0</li><li>goalid to the id of the item that must be fished (from the items table).</li><li>zoneid should be set to 0 (any zone), or the zoneidnumber that the items must be fished in.</li></ul><p>If you wish to allow the player to fish any one of a group of items, then you can use a list as described for the Deliver activity.</p>                                                                                                                                                                                                                                                        |
| 8       | Forage               | <p>To have this activity automatically handled by the task system, set:</p><ul><li>goalmethod to 0</li><li>goalid to the id of the item that must be foraged (from the items table).</li><li>zoneid should be set to 0 (any zone), or the zoneidnumber that the items must be fished in.</li></ul><p>If you wish to allow the player to forage any one of a group of items, then you can use a list as described for the Deliver activity.</p>                                                                                                                                                                                                                                                     |
| 9       | Use                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| 10      | Use                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| 11      | Touch                | <p>To have this activity automatically handled by the task system, set:</p><ul><li>goalmethod to 0</li></ul><p>Since the zoneid field is used to specify the zoneidnumber, the goalid field is not used. It is also not possible to specify a list of zones.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| 100     | Give Cash            | <p>To have this activity automatically handled by the task system, set:</p><ul><li>goalmethod to 0</li><li>npctypeid to the id of the NPC that the cash should be given to</li><li>zoneid should be set to 0 (any zone), or the zoneidnumber that the NPC is found in.</li></ul>                                                                                                                                                                                                                                                                                                                                                                                                                   |

### Character_Tasks Table

#### Field Descriptions

| Column       | Data Type | Description                                                                                                       |
| ------------ | --------- | ----------------------------------------------------------------------------------------------------------------- |
| charid       | int       | Character Identifier; ID from character_data                                                                      |
| taskid       | int       | Task Identifier; ID from tasks table                                                                              |
| slot         | int       | <p>Slot; the slot number (up to 19) that causes the activity to </p><p>display in order to the client</p>         |
| type         | tinyint   | [Task Type](https://eqemu.gitbook.io/server/categories/types/task-types); 0=Task, 1=Shared, 2=Quest, 3=Expedition |
| acceptedtime | int       | Accepted Time UNIX Timestamp                                                                                      |

### Character_Activities Table

#### Field Descriptions

| Column     | Data Type | Description                                         |
| ---------- | --------- | --------------------------------------------------- |
| charid     | int       | Character Identifier; ID from character_data        |
| taskid     | int       | Task Identifier; ID from tasks table                |
| activityid | int       | Activity Identifier, ID from tasks table            |
| donecount  | int       | Done Count (IE 'items looted', 'npcs killed', etc.) |
| completed  | tinyint   | Completed: 0 = False, 1 = True                      |

### Completed Tasks

{% hint style="warning" %}
If the rule **TaskSystem:RecordCompletedTasks** is set to **false**, then this table is not used.
{% endhint %}

#### Field Descriptions

| Column        | Data Type | Description                                        |
| ------------- | --------- | -------------------------------------------------- |
| charid        | int       | Character Identifier; ID from character_data table |
| completedtime | int       | Completed Time UNIX Timestamp                      |
| taskid        | int       | Task Identifier; ID from tasks table               |
| activityid    | int       | Activity Identifier; ID from tasks table           |

{% hint style="info" %}
This table may have duplicate entries for repeatable tasks.  To display Optional Tasks, be sure to set the rule **TaskSystem:RecordCompletedOptionalTasks** to **true**.
{% endhint %}

### Goallists Table

#### Field Descriptions

| Column | Data Type | Description                                          |
| ------ | --------- | ---------------------------------------------------- |
| listid | int       | Goal List Identifier; repeat ID for multiple entries |
| entry  | int       | Entry Identifier; either npc_type id or item id      |

{% hint style="info" %}
Example: list id 10 for a "kill" mission could have three rows, all with list id 10, and entries of 2001, 2102, and 2103 (for Fippy Darkpaw, Kraxz Darkpaw, and Grarrax Darkpaw).
{% endhint %}

### Proximities Table

{% hint style="warning" %}
The rule **TaskSystem:EnableTaskProximity** must be set to **true **to allow the use of proximities independent of proximities defined by quest scripts for npcs (IE for "explore" tasks).
{% endhint %}

#### Field Descriptions

| Column    | Data Type | Description                                                                                                     |
| --------- | --------- | --------------------------------------------------------------------------------------------------------------- |
| zoneid    | int       | [Zone Identifier](https://eqemu.gitbook.io/server/categories/reference-lists/zones); use the zoneidnumber field |
| exploreid | int       | Explore Identifier; match the goalid in the activity table for type 5 (explore) activities                      |
| minx      | float     | Minimum X Coordinate                                                                                            |
| maxx      | float     | Maximum X Coordinate                                                                                            |
| miny      | float     | Minimum Y Coordinate                                                                                            |
| maxy      | float     | Maximum Y Coordinate                                                                                            |
| minz      | float     | Minimum Z Coordinate                                                                                            |
| maxz      | float     | Maximum Z Coordinate                                                                                            |

{% hint style="danger" %}
Overlapping proximities will only trigger proximity detection for the proximity with the lowest-numbered **exploreid**.
{% endhint %}

### Tasksets Table

{% hint style="info" %}
Task Sets are intended to make it easier to write quest NPCs that can offer a lot of tasks. While this functionality can also be achieved with quest::assigntask, it is important to note that players can remove tasks when writing your scripts.  A Task Set could require a series of tasks in a progression.
{% endhint %}

#### Field Descriptions

| Column | Data Type | Description                              |
| ------ | --------- | ---------------------------------------- |
| id     | int       | Unique Task Set Entry Identifier         |
| taskid | int       | Task Identifier; ID from the tasks table |

### Character_enabledtasks Table

#### Field Descriptions

| Column | Data Type | Description                                  |
| ------ | --------- | -------------------------------------------- |
| charid | int       | Character Identifier; ID from character_data |
| taskid | int       | Task Identifier; ID from tasks table         |

## GM Commands

| Command               | Description                                                           |
| --------------------- | --------------------------------------------------------------------- |
| #task show            | Shows the active tasks for the targeted client                        |
| #task update \[count] | Updates the specified activity by 1, or by \[count]                   |
| #task reloadall       | Reload all information relating to tasks from DB (all zones)          |
| #task reload task     | Reload all information for the specified task from the DB (all zones) |
| #task reload lists    | Reload all information from the goallists table (all zones)           |
| #task reload prox     | Reload all information from the proximities table (all zones)         |
| #task reload sets     | Reload task sets (all zones)                                          |

## Task System Rules

| Rule Name                                    | Default Value | Description                                                                                                                |
| -------------------------------------------- | ------------- | -------------------------------------------------------------------------------------------------------------------------- |
| TaskSystem:KeepOneRecordPerCompletedTask     | True          | Keeps one record for each completed task (the client displays the most recent 50; use quest::istaskcompleted for lookups). |
| TaskSystem:EnableTaskProximity               | True          | Makes use of the 'proximities' table, independent of quest::set_proximitiy                                                 |
| TaskSystem:RecordCompletedOptionalActivities | True          | Makes use of the 'character_activities' table for optional tasks                                                           |
| TaskSystem:PeriodicCheckTimer                | 5             | Seconds between checks for failed tasks. Also used by the 'Touch' activity                                                 |
| TaskSystem:RecordCompletedTasks              | True          | Makes use of the 'character_activities' table for non-optional tasks                                                       |
| TaskSystem:EnableTaskSystem                  | True          | Globally enable or disable the Task system                                                                                 |
| QueryServ:PlayerLogTaskUpdates               | False         | Logs Player Task Updates                                                                                                   |

## Logging Options

{% hint style="info" %}
These options are present in the Logging System through [logsys_categories](https://eqemu.gitbook.io/database-schema/categories/admin/logsys_categories) or in-game via command #logs set \[output] \[type] \[level]
{% endhint %}

```sql
1 LOG_CATEGORY( TASKS )
2 LOG_TYPE( TASKS, GLOBALLOAD, DISABLED )
3 LOG_TYPE( TASKS, CLIENTLOAD, DISABLED )
4 LOG_TYPE( TASKS, UPDATE, DISABLED )
5 LOG_TYPE( TASKS, CLIENTSAVE, DISABLED )
6 LOG_TYPE( TASKS, PACKETS, DISABLED )
7 LOG_TYPE( TASKS, PROXIMITY, DISABLED )
```

TASKS UPDATE is used to print out debug information when checking if activities are completed.

## Quest API (for the task system)

These are the **quest::\[methods]** and **EVENT** triggers that are related to the Task System.

{% hint style="info" %}
Be sure to reference the [Quest API](https://eqemu.gitbook.io/quest-api/) in the event that functionality is expanded.
{% endhint %}

### taskselector

      **Parameter(s):**

      task_id _(int)_

      **Usage:**

      Used to bring up the Task Selector Window with the specified tasks available for selection (from 1 to 40 task_id(s)). Note that when the task selector is brought up via this method, no check is made as to whether the character has the tasks enabled.

      **Example:**

```perl
sub EVENT_SAY {
     if ($text=~/hail/i) {
          quest::say("Heyo! Looking for an exciting [task] to fill the time?");
     }
     if ($text=~/task/i) {
          #:: Select task 103
          quest::taskselector(103);
     }
}
```

### task_setselector

      **Parameter(s):**

      task_set_id _(int)_

      **Usage:**

      Sets the Task Set, by provided Task Set ID.

      **Example:**

```perl
#:: Set task set 202
quest::task_setselector(202);
```

### istaskactive

      **Parameter(s):**

      task_id _(int)_

      **Usage:**

      Used to determine if a task is active, by Task ID.

      **Example:**

```perl
#:: Check if task 212 is active
quest::istaskactive(212); #:: Returns bool
```

### istaskactivityactive

      **Parameter(s):**

      task_id _(int)_, activity_id _(int)_

      **Usage:**

      Used to determine if a task activity is active, by Task ID and Activity ID.

      **Example:**

```perl
#:: Check if Activity 9 of Task 212 is active 
quest::istaskactivityactive(212, 9); #:: Returns bool
```

### updatetaskactivity

      **Parameter(s):**

      task_id _(int)_, activity_id _(int)_, count _(int)_, ignore_quest_update _(bool)_

      **Usage:**

      Used to increment the done count of the specified task is active. The parameter for ignore_quest_update is optional, and defaults to false; the parameter for count default to 1 (increase count by 1).

      **Example:**

```perl
#:: Update Task ID 219 activity 9, by 1 count
quest::updatetaskactivity(216,9);
```

### resettaskactivity

      **Parameter(s):**

      task_id _(int)_, activity_id _(int)_

      **Usage:**

      Sets the done count to 0 for the specified Task ID

      **Example:**

```perl
#:: Reset the activity done count for task 202
quest::resettaskactivity(202);
```

### taskexplorearea

      **Parameter(s):**

      explore_id _(int)_

      **Usage:**

      Used to mark any explore activities (type 5), which have the numeric value Explore ID in their Goal ID field, and for which the Zone ID of the activity is either 0 or this zone, as completed.

      **Example:**

```perl
#:: Mark Task 21 complete
quest::taskexplorearea(21);
```

### assigntask

      **Parameter(s):**

      task_id _(int)_, npcid _(int)_, enforce_level_requirement = false _(bool)_

      **Usage:**

      Used to assign a task to a client, optionally this can include the NPC ID and whether or not to enforce the level requirement specified in the DB.

      **Example:**

```perl
#:: Assign Task 102
quest::assigntask(102);
#:: Assign Task 103 with level requirement enforced
quest::assigntask(103, 1);
#:: Assign Task 104 and include the NPC ID
$client->AssignTask(104, $npc->GetID());
#:: Assign Task 105, include the NPC ID with level requirement enforced
$client->AssignTask(105, $npc->GetID(), 1);
```

### failtask

      **Parameter(s):**

      task_id _(int)_

      **Usage:**

      Fails the task, by Task ID, for the client character that triggered the event.

      **Example:**

```perl
#:: Fail Task 216
quest::failtask(216);
```

### tasktimeleft

      **Parameter(s):**

      task_id _(int)_

      **Usage:**

      Returns the amount of time left, in seconds, before the specified task runs out. -1 is returned if there is no time limit, or if the player does not have the task.

      **Example:**

```perl
#:: Create a scalar variable to store the amount of time left for Task ID 22
$timeremaining = quest::tasktimeleft(22);  #:: Returns int
if ($timeremaining => 1) {
     quest::say("You have $timeremaining seconds left--better hurry!");
}
else {
     quest::say("Sorry $name, but you're out of time!");
}
```

### istaskcompleted

      **Parameter(s):**

      task_id (int)

      **Usage:**

      Used to determine if a task is completed, by Task ID.

      **Example:**

```perl
#:: Check if task 212 is completed
quest::istaskcompleted(212); #:: Returns bool
```

### enabletask

      **Parameter(s):**

      task_id _(int)_

      **Usage:**

      Enables a task.

      **Example:**

```perl
#:: Match if the player has an active task
if ($task != 0) {
     #:: Create a scalar variable to store the active speak-to activity
     $activity = quest::activespeakactivity($task);
     #:: Mark the activity as complete
     quest::updatetaskactivity($task, $activity);
     quest::say("Well done!");
     #:: Offer the next task, if there is one
     if (!quest::istaskactive($task)) {
          #:: Disable the task so that it cannot be repeated
          quest::disabletask($task);
          #:: Match if there are tasks remaining in the Task Set
          if ($task != quest::lasttaskinset(200)) {
               quest::say("Well done, I have another task if you are willing.");
               #:: Enable the next Task in the Task Set
               quest::enabletask(quest::nexttaskinset(200, $task));
          }
          #:: Match if there are no tasks remaining in the task set
          else {
               quest::say("Thank you for cleansing Qeynos Hills!");
          }
     }
}
```

### disabletask

      **Parameter(s):**

      task_id _(int)_

      **Usage:**

      Disables a task so that it is not available to a client character. Useful if you do not want someone to repeat a task.

      **Example:**

```perl
#:: Match if the player has an active task
if ($task != 0) {
     #:: Create a scalar variable to store the active speak-to activity
     $activity = quest::activespeakactivity($task);
     #:: Mark the activity as complete
     quest::updatetaskactivity($task, $activity);
     quest::say("Well done!");
     #:: Offer the next task, if there is one
     if (!quest::istaskactive($task)) {
          #:: Disable the task so that it cannot be repeated
          quest::disabletask($task);
          #:: Match if there are tasks remaining in the Task Set
          if ($task != quest::lasttaskinset(200)) {
               quest::say("Well done, I have another task if you are willing.");
               #:: Enable the next Task in the Task Set
               quest::enabletask(quest::nexttaskinset(200, $task));
          }
          #:: Match if there are no tasks remaining in the task set
          else {
               quest::say("Thank you for cleansing Qeynos Hills!");
          }
     }
}
```

### istaskenabled

      **Parameter(s):**

      task_id (int)

      **Usage:**

      Used to determine if a task is enabled, by Task ID.

      **Example:**

```perl
#:: Check if task 212 is enabled
quest::istaskenabled(212); #:: Returns bool
```

### enabledtaskcount

      **Parameter(s):**

      task_set _(int)_

      **Usage:**

      Counts the enabled tasks in the specified Task Set.

      **Example:**

```perl
#:: Match if there are no enabled tasks in Task Set 10
if (quest::enabledtaskcount(10) == 0) {
     quest::enabletask(50, 51, 52);
}
```

### firsttaskinset

      **Parameter(s):**

      task_set _(int)_

      **Usage:**

      Returns the first task in the specified Task Set.

      **Example:**

```perl
sub EVENT_SAY {
     #:: If the player hasn't completed the last task in the Task Set
     if (!quest::istaskcompleted(quest::lasttaskinset(200))) {
          #:: If the player has no tasks enabled for this task set, enable the first one
          if (quest::enabledtaskcount(200) == 0) {
               quest::say("You have not done any of my tasks before!");
               #:: Enable the first task in Task Set 200
               quest::enabletask(quest::firsttaskinset(200));
          }
     }
}
```

### lasttaskinset

      **Parameter(s):**

      task_set _(int)_

      **Usage:**

      Returns the last task in the specified Task Set.

      **Example:**

```perl
#:: Find the Task ID of the last task in Task Set 200
quest::lasttaskinset(200); #:: Returns int
```

### nexttaskinset

      **Parameter(s):**

      task_set _(int)_, task_id _(int)_

      **Usage:**

      Returns the next task in the specified Task Set that comes after the specified Task ID.

      **Example:**

```perl
sub EVENT_SAY {
     #:: Create a scalar variable to store a task integer
     $task = quest::activespeaktask();
     #:: Match if there is an active task to speak to the NPC
     if ($task != 0) {
          #:: Match if there are no active tasks for the current speaking task
          if (!quest::istaskactive($task)) {
               #:: Disable the current speaking task
               quest::disabletask($task);
               #:: Match if the current speaking task is NOT the last task in the Task Set
               if ($task != quest::lasttaskinset(200)) {
                    quest::say("Well done, I have another task if you are willing.");
                    #:: Enable the next task in Task Set 200
                    quest::enabletask(quest::nexttaskinset(200, $task));
               }
               else {
                    quest::say("Thank you for cleansing Qeynos Hills!");
               }
          }
     }
}
```

### activespeaktask

      **Parameter(s):**

      None.

      **Usage:**

      Returns the Task ID of the lowest numbered task slot if the player who triggered the event has an active task with an active activity to speak to the NPC (returns 0 if not).

      **Example:**

      If you have task id 150: activity 0--kill three rats, activity 1--talk with NPC, activity 2--kill four fire beetles, activity 3--talk with NPC.

```perl
sub EVENT_SAY {
     #:: Match text for hail, case insensitive
     if ($text=~/hail/i) {
          #:: Create a scalar variable to store the current task that has a speak activity
          $speaktask = quest::activespeaktask(); #:: Returns int
          #:: Match if there's an active task with a speak activity
          if ($speaktask => 1) {
               quest::say("You have a task to speak with me and its ID is $speaktask.");
          }
          #:: Match if there's NO active task with a speak activity
          else {
               quest::say("You do not have any tasks with a speaking activity at this time.");
          }
     }
}
```

### activespeakactivity

      **Parameter(s):**

      task_id _(int)_

      **Usage:**

      Returns the Activity ID of the lowest numbered active activity to speak with an NPC in the specified task.

      **Example:**

      If you have task id 150: activity 0--kill three rats, activity 1--talk with NPC, activity 2--kill four fire beetles, activity 3--talk with NPC.

```perl
sub EVENT_SAY {
     #:: Match text for hail, case insensitive
     if ($text=~/hail/i) {
          #:: Create a scalar variable to store the current speak activity for task 150
          $speakactivity = quest::activespeakactivity(150); #:: Returns int
          #:: Match if the player is on the first speak activity
          if ($speakactivity == 1) {
               quest::say("I really hate rats.");
          }
          #:: Match if the player is on the second speak activity
          elsif ($speakactivity == 3) {
               quest::say("I really hate fire beetles.");
          }
     }
}
```

### activetasksinset

      **Parameter(s):**

      task_set _(int)_

      **Usage:**

      Returns the number of tasks in the given TaskSet that the player has active.

      **Example:**

      You have a TaskSet "20", which consists of three tasks--200, 201, 202.

```perl
sub EVENT_SAY {
     #:: Create a scalar variable to store the number of active tasks in TaskSet 20
     $activetaskcount = quest::activetasksinset(20);  #:: returns int
     #:: Match text for hail, case insensitive
     if ($text=~/hail/i) {
          quest::say("You have $activetaskcount tasks remaining.");
     }
}
```

### completedtasksinset

      **Parameter(s):**

      task_set _(int)_

      **Usage:**

      Returns the number of tasks in the given Task Set that the player has completed.

      **Example:**

```perl
sub EVENT_SAY {
     #:: Match text for "tasks", case insensitive
     if ($text=~/tasks/i) {
          #:: Create a scalar variable to store the count of completed tasks in task set "200"
          my $TasksComplete = quest::completedtasksinset(200); #:: Returns int
          quest::say("You have completed $TasksComplete tasks.");
     }
}
```

### EVENT_TASKACCEPTED

#### Trigger

* when a player accepts a task from the task selector window.

{% hint style="info" %}
If the assigntask(taskid) function is used to forcibly assign a task to a player, then this sub will also be called with task_id set to taskid if assigntask is successful.
{% endhint %}

### EVENT_TASK_COMPLETE

#### Trigger

* when a player completes a task.

#### Exports

| Name        | Type | Usage                                     |
| ----------- | ---- | ----------------------------------------- |
| donecount   | int  | `quest::say($donecount); # returns int`   |
| activity_id | int  | `quest::say($activity_id); # returns int` |
| task_id     | int  | `quest::say($task_id); # returns int`     |

### EVENT_TASK_FAIL

#### Trigger

* When a player fails a task.

#### Exports

| Name    | Type | Usage                                 |
| ------- | ---- | ------------------------------------- |
| task_id | int  | `quest::say($task_id); # returns int` |

### EVENT_TASK_STAGE_COMPLETE

#### Trigger

* When a task stage is completed.

#### Exports

| Name        | Type | Usage                                     |
| ----------- | ---- | ----------------------------------------- |
| activity_id | int  | `quest::say($activity_id); # returns int` |
| task_id     | int  | `quest::say($task_id); # returns int`     |

#### Example

* In this example, when a player completes a task, it triggers the event and tries to match with task ID 212; if it matches, a yellow message is displayed to the client.

{% tabs %}
{% tab title="Perl" %}
```perl
sub EVENT_TASK_STAGE_COMPLETE {
    #:: Match task id 212
    if ($task_id == 212) {
        $client->Message(15,"The zombie presence seems somewhat lessened, and perhaps they have been quelled...for the time being.");
    }
}
```
{% endtab %}

{% tab title="Lua" %}
```lua
function event_task_stage_complete(e)
    --:: Match task id 212
    if (e.task_id == 212) then
        e.self:Message(15, "The zombie presence seems somewhat lessened, and perhaps they have been quelled...for the time being.");
    end
end
```
{% endtab %}
{% endtabs %}

### EVENT_TASK_UPDATE

#### Trigger

* when a player's task is updated.

#### Exports

| Name        | Type | Usage                                     |
| ----------- | ---- | ----------------------------------------- |
| donecount   | int  | `quest::say($donecount); # returns int`   |
| activity_id | int  | `quest::say($activity_id); # returns int` |
| task_id     | int  | `quest::say($task_id); # returns int`     |
