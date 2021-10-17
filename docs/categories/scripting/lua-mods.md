---
description: EQEMU provides a lua mod system to let you override default game logic
---

# Lua Mods

Inside the directory [https://github.com/EQEmu/Server/tree/master/utils/mods](https://github.com/EQEmu/Server/tree/master/utils/mods) you can find example scripts on how to override combat to a legacy eqemu formula. This page helps details the process to set up adding lua hooks, so you can customize this to your server's needs

In your eqemu server root directory \(where zone, world, shared\_memory etc exists\) create a new folder called mods

On linux, ensure mods and scripts are chmod 755

Ensure load\_order.txt exists under mods, and refers to any hooks that need to be loaded.

Most functions have a e.IgnoreDefault boolean that by default is false. If set to true, you can override all default logic and inject your own.

\#reloadscript can be used to reload the scripts in your zone to reflect your latest changes. Be sure to \#reloadallscripts once you finalize your setup and want it to go across all zones.

| Function | Description |
| :--- | :--- |
| CheckHitChance |  |
| AvoidDamage |  |
| MeleeMitigation |  |
| TryCriticalHit |  |
| ApplyDamageTable |  |
| CommonOutgoingHitSuccess |  |
| GetExperienceForKill |  |
| GetEXPForLevel |  |
| GetRequiredAAExperience |  |

