# Quest Loading

!!! info

    This page outlines how quests are loaded by the server

## Notes

The `[ext]` refers to file extension--either .pl or .lua. 

It should be noted that a .lua file takes precedence over a .pl file of the same name.

## Global Scripts

Global scripts were designed to run on top of the scripts mentioned above, meaning if you have a player script in a zone directory and a global player script, they will both execute and not interfere with each other

| Type | Script Name |
| :--- | :--- |
| Bots | quests/global/global_bot.ext |
| Mercenaries | quests/global/global_merc.ext |
| NPCs | quests/global/global_npc.ext |
| Players | quests/global/global_player.ext |

The `global_npc.ext` scripts run in a zone wide context. All NPCs in the same zone can see variables, modify them, and see changes made by the others

## Bots

In order of operations, a Bot's script will be dictated by the first script that it finds below:

| Hierarchy | Example |
| :--- | :--- |
| quests/zoneshortname/v[instance_version]/bot.[ext] | quests/qeynos/v0/bot.pl |
| quests/zoneshortname/bot_v[instance_version].[ext] | quests/qeynos/bot_v1.pl |
| quests/zoneshortname/bot.[ext] | quests/qeynos/bot.lua |
| quests/global/bot.[ext] | quests/global/bot.pl |

## Encounters

Encounters will load a script on the first event that triggers them and will load one and only one from the following location.

In order of operations, an Encounter's script will be dictated by the first script that it finds below:

| Hierarchy | Example |
| :--- | :--- |
| quests/zone/v[instance_version]/encounters/name.[ext] | quests/qeynos/v0/encounters/name.pl |
| quests/zone/encounters/name.[ext] | quests/qeynos/encounters/name.lua |
| quests/global/encounters/name.[ext] | quests/global/encounters/name.pl |

## Items

Items will load a script on the first event that triggers them and will load one and only one from the following location.

In order of operations, an Item's script will be dictated by the first script that it finds below:

| Hierarchy | Example |
| :--- | :--- |
| quests/zone/v[instance_version]/items/item_script.[ext] | quests/qeynos/v0/items/script_30057.pl |
| quests/zone/items/item_script.[ext] | quests/qeynos/items/script_30057.pl |
| quests/global/items/item_script.[ext] | quests/global/items/script_30057.lua |
| quests/zone/items/default.[ext] | quests/qeynos/items/default.lua |
| quests/global/items/default.[ext] | quests/global/items/default.pl |

## Mercenaries

In order of operations, a Mercenary's script will be dictated by the first script that it finds below:

| Hierarchy | Example |
| :--- | :--- |
| quests/zoneshortname/v[instance_version]/merc.[ext] | quests/qeynos/v0/merc.pl |
| quests/zoneshortname/merc_v[instance_version].[ext] | quests/qeynos/merc_v1.pl |
| quests/zoneshortname/merc.[ext] | quests/qeynos/merc.lua |
| quests/global/merc.[ext] | quests/global/merc.pl |

## NPCs

These scripts run in a context based on NPC ID or NPC name.

All NPCs with the same NPC ID or name can see variables, modify them, and see changes made by the others.

If you need variables that are specific to each instance of an NPC ID or name, data buckets or entity variables can be used.

In order of operations, an NPC's script will be dictated by the first script that it finds below:

| Hierarchy | Example |
| :--- | :--- |
| quests/zoneshortname/v[instance_version]/npc_id.[ext] | quests/qeynos/v0/1173.lua |
| quests/zoneshortname/v[instance_version]/npc_name.[ext] | quests/qeynos/v0/Trumpy_Irontoe.pl |
| quests/zoneshortname/npc_id.[ext] | quests/qeynos/1173.lua |
| quests/zoneshortname/npc_name.[ext] | quests/qeynos/Trumpy_Irontoe.pl |
| quests/global/npc_id.[ext] | quests/global/1234.pl |
| quests/global/npc_name.[ext] | quests/global/Nexus_Scion.lua |
| quests/zoneshortname/v[instance_version]/default.[ext] | quests/tutorialb/v0/default.pl |
| quests/zoneshortname/default.[ext] | quests/tutorialb/default.pl |
| quests/global/default.[ext] | quests/global/default.lua |

## Players

In order of operations, a Player's script will be dictated by the first script that it finds below:

| Hierarchy | Example |
| :--- | :--- |
| quests/zoneshortname/v[instance_version]/player.[ext] | quests/qeynos/v0/player.pl |
| quests/zoneshortname/player_v[instance_version].[ext] | quests/qeynos/player_v1.pl |
| quests/zoneshortname/player.[ext] | quests/qeynos/player.lua |
| quests/global/player.[ext] | quests/global/player.pl |

## Spells

Spells will load a script on the first event that triggers them and will load one and only one from the following location.

In order of operations, a Spell's script will be dictated by the first script that it finds below:

| Hierarchy | Example |
| :--- | :--- |
| quests/zone/v[instance_version]/spells/spell_id.[ext] | quests/qeynos/v0/spells/1234.pl |
| quests/zone/spells/spell_id.[ext] | quests/qeynos/spells/1234.pl |
| quests/global/spells/spell_id.[ext] | quests/global/spells/1234.lua |
| quests/zone/spells/default.[ext] | quests/qeynos/spells/default.lua |
| quests/global/spells/default.[ext] | quests/global/spells/default.pl |

