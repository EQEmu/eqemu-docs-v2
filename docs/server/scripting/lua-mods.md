---
description: EQEMU provides a lua mod system to let you override default game logic
---

# Lua Mods

Inside the directory [https://github.com/EQEmu/Server/tree/master/utils/mods](https://github.com/EQEmu/Server/tree/master/utils/mods) you can find example scripts on how to override combat to a legacy eqemu formula. This page helps details the process to set up adding lua hooks, so you can customize this to your server's needs

In your eqemu server root directory (where zone, world, shared_memory etc exists) create a new folder called mods

On linux, ensure mods and scripts are chmod 755

Ensure load_order.txt exists under mods, and refers to any hooks that need to be loaded.

Most functions have a e.IgnoreDefault boolean that by default is false. If set to true, you can override all default logic and inject your own.

You can use #reload quest to reload the scripts in your zone to reflect your latest changes. Be sure to #reload world once you finalize your setup and want it to go across all zones.

# Function List

Function|Description
:---|:---
**ApplyDamageTable**|Adjust damage based on a table, no example available.
**AvoidDamage**|Check if an attacker avoided damage, no example available
**CalcSpellEffectValue_formula**|Apply custom spell formulas to how a spell value is deduced. An [example is found in utils](https://github.com/EQEmu/Server/blob/master/utils/mods/spell_formula.lua)
**CheckHitChance**|Check if an attacker hit a defender. e.hit will have hit information. [legacy combat example](https://github.com/EQEmu/Server/blob/master/utils/mods/legacy_combat.lua#L97) and [jamfest example](https://github.com/jamfesteq/quests/blob/main/mods/mighty.lua#L3)
**CommonDamage**|Adjust damage done. [jamfest example](https://github.com/jamfesteq/quests/blob/89e5ca78646f258801463dfcbc6b3800e9e3768a/mods/mighty.lua#L130)
**CommonOutgoingHitSuccess**|This is the final value of an outgoing hit success. Note that if any critical hits occured, if this value is modified, the crit won't match up as the text for the crit would already be displayed. [legacy combat example](https://github.com/EQEmu/Server/blob/master/utils/mods/legacy_combat.lua#L961).
**GetExperienceForKill**|Get experience earned for a kill, and modify the amount. [classic wow example](https://github.com/EQEmu/Server/blob/master/utils/mods/classic_wow_experience.lua#L9). Note for more granular control, you can use SetEXP and SetAAEXP.
**GetEXPForLevel**|Get raw experience normally rewarded for leveling. [classic wow example](https://github.com/EQEmu/Server/blob/master/utils/mods/classic_wow_experience.lua#L49)
**GetRequiredAAExperience**|Get how much experience is needed for an AA gain. [classic wow example](https://github.com/EQEmu/Server/blob/master/utils/mods/classic_wow_experience.lua#L4)
**HealDamage**|Adjust healing done. [jamfest example](https://github.com/jamfesteq/quests/blob/89e5ca78646f258801463dfcbc6b3800e9e3768a/mods/mighty.lua#L86)
**IsImmuneToSpell**|Check if a target is immune to a spell, and inject custom logic and message on resist [jamfest example](https://github.com/jamfesteq/quests/blob/89e5ca78646f258801463dfcbc6b3800e9e3768a/mods/mighty.lua#L244)
**MeleeMitigation**|Process melee mitigation (damage reduction). [legacy combat example](https://github.com/EQEmu/Server/blob/master/utils/mods/legacy_combat.lua#L65). |
**RegisterBug**|When a /bug report occurs, this fires with info. [githubeq](https://github.com/xackery/githubeq) uses this to reroute bugs to github issues. [githubeq example](https://github.com/xackery/githubeq/blob/master/lua/register_bug.lua)|
**SetAAEXP**|Set AA experience earned for a kill. [jamfest example](https://github.com/jamfesteq/quests/blob/main/mods/exp.lua#L60)
**SetEXP**|Set experience earned for a kill. [jamfest example](https://github.com/jamfesteq/quests/blob/main/mods/exp.lua#L2)
**TryCriticalHit**|Try executing a critical hit. [legacy combat example](https://github.com/EQEmu/Server/blob/master/utils/mods/legacy_combat.lua#L255).
**UpdatePersonalFaction**|Adjust faction gain. [jamfest example](https://github.com/jamfesteq/quests/blob/89e5ca78646f258801463dfcbc6b3800e9e3768a/mods/faction.lua#L3C33)