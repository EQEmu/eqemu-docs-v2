---
description: >-
  This page details the process of creating an NPC from scratch
---

Here are the blockers you need to overcome to create a custom NPC:

- [Inject a new model name to eqgame](#inject-a-new-model-name-to-eqgame)
- [Increase max race ID server side](#increase-max-race-id-server-side)
- [Add the new NPC to either global or zone loading](#add-the-new-npc-to-either-global-or-zone-loading)

## Inject a new model name to eqgame

Adding new model names is not easily exposed. You need to do a client side edit to inject it, unfortunately. Some ways to achieve this include:

- Edit the eqgame.exe directly (not documented yet)
- Add a dll hook to inject the new NPCs via dinput8.dll (recommended)

To inject the new NPCs via dinput8.dll, one of the easier routes is to use [eq-core-dll](https://github.com/xackery/eq-core-dll). The setup process for this is on the README of the link, follow the steps, contact the discord on eqemu if any help, a number of folks are familiar with the process.

From now on, we're going to continue the wiki with the shortname of `WTF` and race ID of `733`, which is the next unused ID for RoF2 by default.


Once set up, inside _options.h you want to add the new NPC shortname and new unique ID, similar to this:

```cpp
// areCustomNPCsEnabled if set to true will allow the NPCs defined in NPCs[] to be injected in game
bool areCustomNPCsEnabled = false;

// NPC Entry:
// raceID is the index. If it's a new NPC, start at 733. You'll need to update the rule NPC:MaxRaceID
// GenderID ranges from 0 to 2 usually
// modelName is the race's shortname tag
// raceMask is a range of flags, typically using 8 is safe for most NPCs, but e.g. 1 = drivable boat, 2 = ridable boat, etc
// dbStrID if left to 1 reverts to raceID as it's ID, otherwise you can custom set one, and it'll look up dbStr for info
static NPCEntry NPCs[] = {
    // raceID, genderID, modelName, raceMask, dbStrID
    NPCEntry(733, 2, "WTF", 3, 1),
};
```

Change areCustomNPCsEnabled to true, and add your new NPC to the list. You can add as many as you want, but you'll need to update the MaxRaceID rule to match the highest raceID you use.

[An example of existing race IDs and their values can be found in this gist](https://gist.github.com/xackery/f39d14b93004dae295e861f387a7374e)

Once you build the solution, copy dinput8.dll to your EQ directory and you've now injected a new model name to eqgame.

# Increase max race ID server side

This step requires you to edit the rules table. You can do this via spire, peq editor, database, or via the server console. I'll assume you've played enough with eqemu to know how to edit rules already, so the changes are:

- Edit NPC:MaxRaceID from 732 to your new value, 733, or higher if you added more than one NPC
- Reload rules in game via #reload rule

# Add the new NPC to either global or zone loading

