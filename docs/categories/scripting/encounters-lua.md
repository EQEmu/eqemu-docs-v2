---
description: >-
  This page describes using scripted Encounters, which is applicable to Lua
  scripts only (not perl).
---

# Encounters \(Lua\)

#### Encounter Scripts

Encounter scripts are quest scripts that are only loaded after explicitly called with one of the following:

```text
eq.load_encounter("encounter_name");
eq.load_encounter_with_data("encounter_name", "some data string");
```

They will load one and only one from the following location. Which ever it finds first in the following order:

* ./quests/zone/encounters/name.lua
* ./quests/global/encounters/name.lua

Encounter scripts listen for specific events from other script types with the following functions:

```text
Void register_npc_event(String name, Integer evt, Integer npc_id, luafunction func);
Void register_player_event(String name, Integer evt, luafunction func);
Void register_item_event(String name, Integer evt, Integer item_id, luafunction func);
Void register_spell_event(String name, Integer evt, Integer spell_id, luafunction func);

Note: Passing a value of -1 for npc, item or spell id to watch will watch every npc, item or spell for those events.
```

Encounters can be unloaded with the following in the same way they are loaded:

```text
eq.unload_encounter("encounter_name");
eq.unload_encounter_with_data("encounter_name", "some data string");
```

Unloading the encounter will automatically unhook all functions that are currently hooked to that named encounter.

Note:

* Encounter scripts cannot properly catch EVENT\_COMMAND or EVENT\_TRADE unless an existing quest is already listening for them.
* Encounter scripts also run before any normal script and will catch return values.

