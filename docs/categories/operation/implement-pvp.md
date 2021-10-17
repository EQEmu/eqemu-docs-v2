# Implement PvP

Note: PVP Is currently a work in progress.

To emulate the original Everquest PVP servers, the only rule you have to adjust is the PVPSettings one.

### World:PVPSettings

| Value | Name | Desc |
| :--- | :--- | :--- |
| 0 | Disabled | Default, No PVP Enabled |
| 1 | Rallos Zek | Not yet supported |
| 2 | TZ/VZ | Partial Support |
| 4 | Sullon Zek | Partial support |
| 6 | Discord | Partial Support |
| 7+ | Discord+ \(Any value above 6\) | Discord without the no-drop restrictions removed |

Rules from this point onward are optional, and allow you the ability to customize your PVP server beyond the classic PVP rules. If you enable PVPSettings then adjust the below, most values will override defaults. Some exceptions include PVPLootCoin \(once true by above PVPSettings, it cannot be turned off without setting PVPSettings to 0\).

### World:PVPMinLevel

minimum level pvp is enabled

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| 0 | 1 | ? | 6 |

### World:PVPUseDeityBasedPVP

use deity based pvp alignment \(agnostic is neutral\)

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| false | false | false | true |

### World:PVPUseTeamsBySizeBasedPVP

use racial size based pvp alignment \(drakkin is human group\)

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| false | false | true | false |

The teams are as follows:

* Team 1: Human, Barbarian, Erudite, Drakkin
* Team 2: Gnome, Halfling, Dwarf, Froglok
* Team 3: High Elf, Wood Elf, Half Elf, Vahshir
* Team 4: Dark Elf, Ogre, Troll, Iksar
* **Note: The above teams may be inaccurate. If you can source a reference to correct this, please let Shin know**

### World:PVPLevelDifference

Players with a difference greater than value will not be attackable

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| 0 | 4 | ? | 6 |

### World:PVPLoseExperienceLevelDifference

Players lose experience if killed by a player within level difference

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| 0 | 0 | ? | 5 |

### World:PVPPetDamageMitigation

Pet damage is mitigated by this amount

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| 50 | 100 | ? | ? |

### World:PVPMeleeMitigation

Melee is mitigated by this amount

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| 67 | 100 | ? | ? |

### World:PVPSpellMitigation

Spells are mitigated by this amount

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| 67 | 67 | ? | ? |

### World:PVPRangedMitigation

Ranged attacks are mitigated by this amount

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| 80 | 100 | ? | ? |

### World:PVPCanLootCoin

Can players loot coin from player corpses?

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| false | true | true | true |

### Spells:PVPRootBreakFromSpells

Chance for root to break when cast on by a client \(20% more than native root\)

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| 75 | 75 | 75 | 75 |

### Character:PVPRespawnManaPercent

Percent of mana to respawn with

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| 100 | 0 | 0 | 0 |

### Inventory:PVPCanLootNoTrade

Can players loot no trade items from player corpses?

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| false | false | false | false |

### Inventory:PVPCanLootContainer

Can players loot bags/containers from player corpses?

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| false | false | false | false |

### Inventory:PVPLootableEquipSlots

When looting a player's corpse, equipped items in slots listed here are lootable. 0 means all slots are lootable

| Default | RZ | TV/VZ | SZ |
| :--- | :--- | :--- | :--- |
| 0 | 0 | 0 | 0 |

