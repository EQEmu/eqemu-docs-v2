---
description: EQEMU provides a lua mod system to let you override default game logic
---

# Lua Mods

Inside the directory [https://github.com/EQEmu/Server/tree/master/utils/mods](https://github.com/EQEmu/Server/tree/master/utils/mods) you can find example scripts on how to override combat to a legacy eqemu formula. This page helps details the process to set up adding lua hooks, so you can customize this to your server's needs

In your eqemu server root directory (where zone, world, shared_memory etc exists) create a new folder called mods

On linux, ensure mods and scripts are chmod 755

Ensure load_order.txt exists under mods, and refers to any hooks that need to be loaded.

Most functions have a e.IgnoreDefault boolean that by default is false. If set to true, you can override all default logic and inject your own.

You can use #reload quest can be used to reload the scripts in your zone to reflect your latest changes. Be sure to #reload world once you finalize your setup and want it to go across all zones.

# Function List

| Function | Description |
| :--- | :--- |
| [CheckHitChance](https://github.com/EQEmu/Server/blob/master/utils/mods/legacy_combat.lua#L97) | Check if an attacker hit a defender. e.hit will have hit information|
| AvoidDamage | Check if an attacker avoided damage, no example available |
| [MeleeMitigation](https://github.com/EQEmu/Server/blob/master/utils/mods/legacy_combat.lua#L65) | Process melee mitigation (damage reduction) |
| [TryCriticalHit](https://github.com/EQEmu/Server/blob/master/utils/mods/legacy_combat.lua#L255) | Try executing a critical hit |
| ApplyDamageTable | Adjust damage based on a table, no example available |
| [CommonOutgoingHitSuccess](https://github.com/EQEmu/Server/blob/master/utils/mods/legacy_combat.lua#L961) | This is the final value of an outgoing hit success. Note that if any critical hits occured, if this value is modified, the crit won't match up as the text for the crit would already be displayed. |
| [GetExperienceForKill](https://github.com/EQEmu/Server/blob/master/utils/mods/classic_wow_experience.lua#L9) | Get experience earned for a kill, and modify the amount |
| [GetEXPForLevel](https://github.com/EQEmu/Server/blob/master/utils/mods/classic_wow_experience.lua#L49) | Get raw experience normally rewarded for leveling |
| [GetRequiredAAExperience](https://github.com/EQEmu/Server/blob/master/utils/mods/classic_wow_experience.lua#L4) | Get how much experience is needed for an AA gain |

