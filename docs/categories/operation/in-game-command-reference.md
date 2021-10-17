---
description: >-
  This page lists the commands that are available in-game, based on assigned
  Account Status, for your EQEmu Server.
---

# Commands Reference

Often referred to as "GM Commands", the following commands can be allowed based on the player's Account Status.

| Command | Description |
| :--- | :--- |
| \#acceptrules | \[acceptrules\] - Accept the EQEmu Agreement |
| \#advnpcspawn | \[maketype\|makegroup\|addgroupentry\|addgroupspawn\]\[removegroupspawn\|movespawn\|editgroupbox\|cleargroupbox\] |
| \#aggro | \(range\) \[-v\] - Display aggro information for all mobs 'range' distance from your target. -v is verbose faction info. |
| \#aggrozone | \[aggro\] - Aggro every mob in the zone with X aggro. Default is 0. Not recommend if you are not invulnerable. |
| \#ai | \[factionid/spellslist/con/guard/roambox/stop/start\] - Modify AI on NPC target |
| \#appearance | \[type\] \[value\] - Send an appearance packet for you or your target |
| \#apply shared memory | \[shared memory name\] Tells every zone and world to apply a specific shared memory segment by name |
| \#attack | \[targetname\] - Make your NPC target attack targetname |
| \#augmentitem | Force augments an item. Must have the augment item window open. |
| \#ban | \[name\] - Ban by character name |
| \#beard | Change the beard of your target |
| \#beardcolor | Change the beard color of your target |
| \#bestz | Ask map for a good Z coord for your x,y coords. |
| \#bind | Sets your targets bind spot to their current location |
| \#bot | Type "\#bot help" to the see the list of available commands for bots. |
| \#camerashake | Shakes the camera on everyone's screen globally |
| \#castspell | \[spellid\] - Cast a spell |
| \#chat | \[channel num\] \[message\] - Send a channel message to all zones |
| \#checklos | Check for line of sight to your target |
| \#copycharacter | \[source character name\] \[dest character name\] \[dest account name\] - Create a copy of a character |
| \#corpse | Manipulate corpses, use with no arguments for help |
| \#corpsefix | Attempts to bring corpses from underneath the ground within close proximity of the player |
| \#crashtest | Crash the zoneserver |
| \#cvs | Summary of client versions currently online. |
| \#damage | \[amount\] - Damage your target |
| \#databuckets | \[view\|delete\]\[key\]\[limit = 50\] - View or delete data bucket by key |
| \#date | \[yyyy\] \[mm\] \[dd\] \[HH\] \[MM\] - Set EQ time |
| \#dbspawn2 | \[spawngroup\] \[respawn\] \[variance\] - Spawn an NPC from a predefined row in the spawn2 table |
| \#delacct | \[accountname\] - Delete an account |
| \#deletegraveyard | \[zone name\] - Deletes the graveyard for the specified zone. |
| \#delpetition | \[petition number\] - Delete a petition |
| \#depop | Depop your NPC target |
| \#depopzone | Depop the zone |
| \#details | Change the details of your target \(Drakkin only\) |
| \#devtools | Manages devtools |
| \#disablerecipe | \[recipe id\] - Disabled the specified recipe ID. |
| \#disarmtrap | Analog for LDoN disarm trap for the newer clients since we still don't have it working |
| \#distance | Reports the distance between you and your target. |
| \#doanim | \[animnum\] \[type\] - Send an EmoteAnim for you or your target |
| \#dz | Manage expeditions and dynamic zone instances |
| \#dzkickplayers | Removes all players from current expedition \(/kick players alternative for pre-RoF clients\) |
| \#editmassrespawn | \[name search\]\[second value\] - Mass \(zone-wide\) NPC respawn timer editing command |
| \#emote | \['name'/'world'/'zone'\] \[type\] \[message\] - Send an emote message |
| \#emotesearch | Searches loaded NPC emotes |
| \#emoteview | Lists all of a NPCs loaded emotes |
| \#enablerecipe | \[recipe id\] - Enables the specified recipe ID. |
| \#endurance | Restores your or your target's endurance |
| \#equipitem | \[slotid\(0-22\)\] - Equip the item on your cursor into the specified slot |
| \#face | \[number of face\] - Sets you or your target's face to face number, temporarily. |
| \#faction | \[find \(criteria\|all\) \| review \(criteria\|all\) \| reset \(id\)\] - Resets player's faction |
| \#findaliases | \[search criteria\] - Searches for available command aliases by alias or command |
| \#findnpctype | \[search criteria\] - Search database NPC types |
| \#findrace | \[search criteria\] - Search for a race |
| \#findspell | \[search criteria\] - Search for a spell |
| \#findzone | \[search criteria\] - Search database for zone |
| \#fixmob | \[nextrace\|prevrace\|gender\|nexttexture\|prevtexture\|nexthelm\|prevhelm\] - Manipulate appearance of your NPC target |
| \#flag | \[status\] \[acctname\] - Refresh your admin status, or set an account's admin status if arguments provided |
| \#flagedit | Edit zone flags on your target |
| \#flags | Displays the flags of you or your target |
| \#flymode | \[0\|1\|2\|3\|4\|5\] - Set your or your player target's flymode to ground\|flying\|levitate\|water\|floating\|levitate running |
| \#fov | Check whether you are behind or in your target's field of view |
| \#freeze | Freeze your target |
| \#gassign | \[id\] - Assign targetted NPC to predefined wandering grid id |
| \#gearup | Developer tool to quickly equip a character |
| \#gender | \[0/1/2\] - Change your or your target's gender to male/female/neuter |
| \#getplayerburriedcorp | Get the target's total number of buried player corpses. |
| \#getvariable | \[varname\] - Get the value of a variable from the database |
| \#ginfo | Get group info on target. |
| \#giveitem | \[itemid\] \[charges\] - Summon an item onto your target's cursor. Charges are optional. |
| \#givemoney | \[pp\] \[gp\] \[sp\] \[cp\] - Gives specified amount of money to the target player. |
| \#globalview | Lists all qglobals in cache if you were to do a quest with this target. |
| \#gm | Turn player target's or your GM flag on or off |
| \#gmspeed | \[on/off\] - Turn GM speed hack on/off for you or your player target |
| \#gmzone | \[zone short name\]\[zone version = 0\]\[identifier = gmzone\] - Zones to a private GM instance |
| \#goto | \[x\] \[y\] \[z\] - Teleport to the provided coordinates or to your target |
| \#grid | \[add/delete\] \[grid\_num\] \[wandertype\] \[pausetype\] - Create/delete a wandering grid |
| \#guild | Guild manipulation commands. Use argument help for more info. |
| \#guildapprove | \[guildapproveid\] - Approve a guild with specified ID \(guild creator receives the id\) |
| \#guildcreate | \[guildname\] - Creates an approval setup for guild name specified |
| \#guildlist | \[guildapproveid\] - Lists character names who have approved the guild specified by the approve id |
| \#hair | Change the hair style of your target |
| \#haircolor | Change the hair color of your target |
| \#haste | \[percentage\] - Set your haste percentage |
| \#hatelist | Display hate list for target. |
| \#heal | Completely heal your target |
| \#helm | Change the helm of your target |
| \#help | \[search term\] - List available commands and their description, specify partial command as argument to search |
| \#heritage | Change the heritage of your target \(Drakkin only\) |
| \#heromodel | \[hero model\]\[slot\] - Full set of Hero's Forge Armor appearance. If slot is set, sends exact model just to slot. |
| \#hideme | \[on/off\] - Hide yourself from spawn lists. |
| \#hotfix | \[hotfix name\] - Reloads shared memory into a hotfix. Equivalent to load shared memory followed by apply shared memory |
| \#hp | Refresh your HP bar from the server. |
| \#incstat | Increases or Decreases a client's stats permanently. |
| \#instance | Modify Instances |
| \#interrogateinv | use \[help\] argument for available options |
| \#interrupt | \[message id\] \[color\] - Interrupt your casting. Arguments are optional. |
| \#invsnapshot | Manipulates inventory snapshots for your current target |
| \#invul | \[on/off\] - Turn player target's or your invulnerable flag on or off |
| \#ipban | \[IP address\] - Ban IP by character name |
| \#iplookup | \[charname\] - Look up IP address of charname |
| \#iteminfo | Get information about the item on your cursor |
| \#itemsearch | \[search criteria\] - Search for an item |
| \#kick | \[charname\] - Disconnect charname |
| \#kill | Kill your target |
| \#killallnpcs | \[npc name\] - Kills all NPCs by search name. Leave blank for all attackable NPCs |
| \#lastname | \[new lastname\] - Set your or your player target's lastname |
| \#level | \[level\] - Set your or your target's level |
| \#list | \[npcs\|players\|corpses\|doors\|objects\] \[search\] - Search entities |
| \#listnpcs | \[name/range\] - Search NPCs |
| \#listpetition | List petitions |
| \#load shared memory | \[shared memory name\] - Reloads shared memory and uses the input as output |
| \#loc | Print out your or your target's current location and heading |
| \#lock | Lock the worldserver |
| \#logs | Manage anything to do with logs |
| \#logtest | Performs log performance testing |
| \#makepet | \[level\] \[class\] \[race\] \[texture\] - Make a pet |
| \#mana | Fill your or your target's mana |
| \#maxskills | Maxes skills for you |
| \#memspell | \[slotid\] \[spellid\] - Memorize spellid in the specified slot |
| \#merchant close shop | Closes a merchant's shop |
| \#merchant open shop | Opens a merchant's shop |
| \#modifynpcstat | Modifys a NPC's stats |
| \#motd | \[new motd\] - Set message of the day |
| \#movechar | \[charname\] \[zonename\] - Move charname to zonename |
| \#movement | Various movement commands |
| \#myskills | Show details about your current skill levels |
| \#mysql | Mysql CLI, see 'help' for options. |
| \#mysqltest | MySQL bench test |
| \#mystats | Show details about you or your pet. |
| \#name | \[newname\] - Rename your player target |
| \#netstats | Gets the network stats for a stream |
| \#network | Admin commands for the UDP network interface |
| \#npccast | \[targetname/entityid\] \[spellid\] - Causes NPC target to cast spellid on targetname/entityid |
| \#npcedit | \[column\] \[value\] - Mega NPC editing command |
| \#npceditmass | \[name-search\]\[column\]\[value\] - Mass \(zone-wide\) NPC data editing command |
| \#npcemote | \[message\] - Make your NPC target emote a message. |
| \#npcloot | \[show/money/add/remove\] \[itemid/all/money: pp gp sp cp\] - Manipulate the loot an NPC is carrying |
| \#npcsay | \[message\] - Make your NPC target say a message. |
| \#npcshout | \[message\] - Make your NPC target shout a message. |
| \#npcspawn | \[create/add/update/remove/delete\] - Manipulate spawn DB |
| \#npcspecialattk | \[flagchar\] \[perm\] - Set NPC special attack flags. Flags are E\(nrage\) F\(lurry\) R\(ampage\) S\(ummon\). |
| \#npcstats | Show stats about target NPC |
| \#npctype cache | \[id\|all\] - Clears the NPC type cache for either the ID or ALL NPCs |
| \#npctypespawn | \[npctypeid\] \[factionid\] - Spawn an NPC from the db |
| \#nudge | Nudge your target's current position by specific values |
| \#nukebuffs | Strip all buffs on you or your target |
| \#nukeitem | \[itemid\] - Remove itemid from your player target's inventory |
| \#object | List\|Add\|Edit\|Move\|Rotate\|Copy\|Save\|Undo\|Delete - Manipulate static and tradeskill objects within the zone |
| \#oocmute | \[1/0\] - Mutes OOC chat |
| \#opcode | OPCODE management |
| \#path | View and edit pathing |
| \#peekinv | \[worn/cursor/inv/bank/trade/all\] - Print out contents of your player target's inventory |
| \#peqzone | \[zonename\] - Go to specified zone, if you have &gt; 75% health |
| \#permaclass | \[classnum\] - Change your or your player target's class \(target is disconnected\) |
| \#permagender | \[gendernum\] - Change your or your player target's gender \(zone to take effect\) |
| \#permarace | \[racenum\] - Change your or your player target's race \(zone to take effect\) |
| \#petitioninfo | \[petition number\] - Get info about a petition |
| \#petname | \[newname\] - Temporarilyrenames your pet. Leave name blank to restore the original name. |
| \#pf | Useless information. |
| \#picklock | Used to pick locks. |
| \#profanity | Manage censored language |
| \#profiledump | Dump profiling info to logs |
| \#profilereset | Reset profiling info |
| \#proximity | Shows NPC proximity |
| \#push | Lets you do spell push |
| \#pvp | \[on/off\] - Set your or your player target's PVP status |
| \#qglobal | \[on/off/view\] - Toggles qglobal functionality on an NPC |
| \#questerrors | Shows quest errors |
| \#race | \[racenum\] - Change you or your target's race. Use \#race 0 to return to normal. This is temporary, you revert back to normal upon zoning. |
| \#raidloot | \[Leader\|Groupleader\|Selected\|All\] - Sets your raid loot settings if you have permission to do so |
| \#randomfeatures | Randomly changes the Facial Features of your target |
| \#refreshgroup | Refreshes Group. |
| \#reloadaa | Reloads AA data |
| \#reloadallrules | Executes a reload of all rules |
| \#reloademote | Reloads NPC emotes |
| \#reloadlevelmods |  |
| \#reloadmerchants |  |
| \#reloadperlexportsettings |  |
| \#reloadqst | Clear quest cache |
| \#reloadrulesworld | Executes a reload of all rules in world specifically |
| \#reloadstatic | Reload Static Zone Data |
| \#reloadtitles | Reload player titles from the database |
| \#reloadtraps | Repops all traps in the current zone |
| \#reloadworld | Reloads all scripts and repops all NPC's globally. |
| \#repop | Repop the zone with optional delay. Using \#repop force forcefully resets spawn timers. |
| \#resetaa | Resets a Player's AA in their profile. |
| \#resetaa timer | Command to reset AA cooldown timers |
| \#revoke | \[charname\] \[1/0\] - Makes charname unable to talk on OOC and unable to use /tell. |
| \#roambox | Manages roambox settings for an NPC |
| \#rules | \(subcommand\) - Manage server rules |
| \#save | Force your player or player corpse target to be saved to the database |
| \#scale | Handles NPC scaling |
| \#scribespell | \[spellid\] - Scribe specified spell in your target's spell book. |
| \#scribespells | \[maxlevel\], \[minlevel\] - Scribe all spells for you or your player target that are usable by them, up to level specified. |
| \#sendzonespawns | Refresh spawn list for all clients in zone |
| \#sensetrap | Analog for LDoN sense trap for the newer clients since we still don't have it working |
| \#serverinfo | Get OS info about server host |
| \#serverrules | Read this server's rules |
| \#setaapts | Set your or your player target's available AA points |
| \#setaaxp | \[amount\] - Adds the amount of AA experience you specified to you or your target. |
| \#setadventurepoints | Set you or your player target's available adventure points |
| \#setanim | \[animnum\] - Set target's appearance to animnum |
| \#setcrystals | Set your or your player targets available radiant or ebon crystals |
| \#setfaction | \[faction number\] - Sets targeted NPC's faction in the database |
| \#setgraveyard | \[zone name\] - Creates a graveyard for the specified zone based on your target's LOC. |
| \#setlanguage | \[language ID\] \[value\] - Set your target's language skillnum to value |
| \#setlsinfo | \[email\] \[password\] - Set login server email address and password. |
| \#setpass | \[accountname\] \[password\] - Set local password for accountname |
| \#setpvppoints | Set your or your player targets PVP points |
| \#setskill | \[skillnum\] \[value\] - Set your target's skill skillnum to value |
| \#setskillall | \[value\] - Set all of your target's skills to value |
| \#setstartzone | \[zoneid\] - Set target's starting zone. Set to zero to allow the player to use /setstartcity |
| \#setstat | Sets the stats to a specific value. |
| \#setxp | \[value\] - Set your or your player target's experience |
| \#showbonusstats | \[item\|spell\|all\] - Shows bonus stats for target from items or spells. Shows both by default. |
| \#showbuffs | List buffs active on your target or you if no target |
| \#shownpcgloballoot | Show GlobalLoot entries on this NPC |
| \#shownumhits | Shows buffs numhits for yourself |
| \#showskills | Show the values of your or your player target's skills |
| \#showstats | Show details about you or your target |
| \#showzonegloballoot | Show GlobalLoot entries on this zone |
| \#showzonepoints | Show zone points for current zone |
| \#shutdown | Shut this zone process down |
| \#size | \[size\] - Change size of you or your target |
| \#spawn | \[name\] \[race\] \[level\] \[material\] \[hp\] \[gender\] \[class\] \[priweapon\] \[secweapon\] \[merchantid\] - Spawn an NPC |
| \#spawneditmass | Mass editing spawn command |
| \#spawnfix | Find targeted NPC in database based on its X/Y/heading and update the database to make it spawn at your current location/heading. |
| \#spawnstatus | Show respawn timer status |
| \#spellinfo | \[spellid\] - Get detailed info about a spell |
| \#spoff | Sends OP\_ManaChange |
| \#spon | Sends OP\_MemorizeSpell |
| \#stun | \[duration\] - Stuns you or your target for duration |
| \#summon | \[charname\] - Summons your player/npc/corpse target, or charname if specified |
| \#summonburriedplayerc | Summons the target's oldest buried corpse, if any exist. |
| \#summonitem | \[itemid\] \[charges\] - Summon an item onto your cursor. Charges are optional. |
| \#suspend | Suspend by character name and for specified number of days |
| \#task | \(subcommand\) - Task system commands |
| \#tattoo | Change the tattoo of your target \(Drakkin only\) |
| \#tempname | \[newname\] - Temporarily renames your target. Leave name blank to restore the original name. |
| \#test | Test command |
| \#texture | \[texture\] \[helmtexture\] - Change your or your target's appearance, use 255 to show equipment |
| \#time | \[HH\] \[MM\] - Set EQ time |
| \#timers | Display persisten timers for target |
| \#timezone | \[HH\] \[MM\] - Set timezone. Minutes are optional |
| \#title | \[text\] - Set your or your player target's title |
| \#titlesuffix | \[text\] \[1 = create title table row\] - Set your or your player target's title suffix |
| \#traindisc | \[level\] - Trains all the disciplines usable by the target, up to level specified. \(may freeze client for a few seconds\) |
| \#trapinfo | Gets information about the traps currently spawned in the zone |
| \#tune | Calculate ideal statistical values related to combat |
| \#ucs | Attempts to reconnect to the UCS server |
| \#undyeme | Remove dye from all of your armor slots |
| \#unfreeze | Unfreeze your target |
| \#unlock | Unlock the worldserver |
| \#unscribespell | \[spellid\] - Unscribe specified spell from your target's spell book. |
| \#unscribespells | Clear out your or your player target's spell book. |
| \#untraindisc | \[spellid\] - Untrain specified discipline from your target |
| \#untraindiscs | Untrains all disciplines from your target |
| \#uptime | \[zone server id\] - Get uptime of worldserver, or zone server if argument provided |
| \#version | Display current version of EQEmu server |
| \#viewnpctype | \[npctype id\] - Show info about an npctype |
| \#viewpetition | \[petition number\] - View a petition |
| \#wc | \[wear slot\] \[material\] - Sends an OP\_WearChange for your target |
| \#weather | \[0/1/2/3\] \(Off/Rain/Snow/Manual\) - Change the weather |
| \#who | \[search\] |
| \#worldshutdown | Shut down world and all zones |
| \#worldwide | Performs world-wide GM functions such as cast \(can be extended for other commands\). Use caution. |
| \#wp | \[add/delete\] \[grid\_num\] \[pause\] \[wp\_num\] - Add/delete a waypoint to/from a wandering grid |
| \#wpadd | \[pause\] - Add your current location as a waypoint to your NPC targets AI path |
| \#wpinfo | Show waypoint info about your NPC target |
| \#xtargets | Show your targets Extended Targets and optionally set how many xtargets they can have |
| \#zclip | \[min\] \[max\] - modifies and resends zhdr packet |
| \#zcolor | \[red\] \[green\] \[blue\] - Change sky color |
| \#zheader | \[zonename\] - Load zheader for zonename from the database |
| \#zone | \[zonename\] \[x\] \[y\] \[z\] - Go to specified zone \(co-ordinates optional\) |
| \#zonebootup | \[ZoneServerID\] \[shortname\] - Make a zone server boot a specific zone |
| \#zoneinstance | \[instanceid\] \[x\] \[y\] \[z\] - Go to specified instance zone \(co-ordinates optional\) |
| \#zonelock | \[list/lock/unlock\] - Set/query lock flag for zoneservers |
| \#zoneshutdown | \[shortname\] - Shut down a zone server |
| \#zonespawn | Not implemented |
| \#zonestatus | Show connected zoneservers, synonymous with /servers |
| \#zopp | Troubleshooting command. Sends a fake item packet to you. No server reference is created. |
| \#zsafecoords | \[x\] \[y\] \[z\] - Set safe co-ordinates |
| \#zsave | Saves zheader to the database |
| \#zsky | \[skytype\] - Change zone sky type |
| \#zstats | Show info about zone header |
| \#zunderworld | \[zcoord\] - Sets the underworld using zcoord |
| \#zuwcoords | \[zcoord\] - Set underworld zcoord |

