!!! info end

    This page lists the commands that are available in-game, based on assigned Account Status, for your EQEmu Server.

    Last Generated: 2022.11.1


| Command | Description | Status Level |
| :--- | :--- | :--- |
| #acceptrules | [acceptrules] - Accept the EQEmu Agreement | Player (0) |
| #advnpcspawn | [maketype&#124;makegroup&#124;addgroupentry&#124;addgroupspawn][removegroupspawn&#124;movespawn&#124;editgroupbox&#124;cleargroupbox] | GMLeadAdmin (150) |
| #aggro | [Distance] [-v] - Display aggro information for all mobs 'Distance' distance from your target. (-v is verbose Faction Information) | QuestTroupe (80) |
| #aggrozone | [aggro] - Aggro every mob in the zone with X aggro. Default is 0. Not recommend if you're not invulnerable. | GMAdmin (100) |
| #ai | [factionid/spellslist/con/guard/roambox/stop/start] - Modify AI on NPC target | GMAdmin (100) |
| #appearance | [type] [value] - Send an appearance packet for you or your target | GMLeadAdmin (150) |
| #appearanceeffects | [view] [set] [remove] appearance effects. | GMAdmin (100) |
| #apply_shared_memory | [shared_memory_name] - Tells every zone and world to apply a specific shared memory segment by name. | GMImpossible (250) |
| #attack | [Entity Name] - Make your NPC target attack an entity by name | GMLeadAdmin (150) |
| #augmentitem | Force augments an item. Must have the augment item window open. | GMImpossible (250) |
| #ban | [Character Name] [Reason] - Ban by character name | GMLeadAdmin (150) |
| #bind | Sets your targets bind spot to their current location | GMMgmt (200) |
| #bot | Type \#bot help\ or \^help\ to the see the list of available commands for bots. | Player (0) |
| #camerashake | [Duration (Milliseconds)] [Intensity (1-10)] - Shakes the camera on everyone's screen globally. | QuestTroupe (80) |
| #castspell | [Spell ID] [Instant (0 = False, 1 = True, Default is 1 if Unused)] - Cast a spell | Guide (50) |
| #chat | [channel num] [message] - Send a channel message to all zones | GMMgmt (200) |
| #checklos | Check for line of sight to your target | Guide (50) |
| #copycharacter | [source_char_name] [dest_char_name] [dest_account_name] - Copies character to destination account | GMImpossible (250) |
| #corpse | Manipulate corpses, use with no arguments for help | Guide (50) |
| #corpsefix | Attempts to bring corpses from underneath the ground within close proximity of the player | Player (0) |
| #countitem | [Item ID] - Counts the specified Item ID in your or your target's inventory | GMLeadAdmin (150) |
| #cvs | Summary of client versions currently online. | GMMgmt (200) |
| #damage | [Amount] - Damage yourself or your target | GMAdmin (100) |
| #databuckets | View&#124;Delete [key] [limit]- View data buckets, limit 50 default or Delete databucket by key | QuestTroupe (80) |
| #date | [Year] [Month] [Day] [Hour] [Minute] - Set EQ time (Hour and Minute are optional) | EQSupport (90) |
| #dbspawn2 | [Spawngroup ID] [Respawn] [Variance] [Condition ID] [Condition Minimum] - Spawn an NPC from a predefined row in the spawn2 table, Respawn and Variance are in Seconds (condition is optional) | GMAdmin (100) |
| #delacct | [accountname] - Delete an account | GMLeadAdmin (150) |
| #delpetition | [petition number] - Delete a petition | ApprenticeGuide (20) |
| #depop | [Start Spawn Timer] - Depop your NPC target and optionally start their spawn timer (false by default) | Guide (50) |
| #depopzone | [Start Spawn Timers] - Depop the zone and optionally start spawn timers (false by default) | GMAdmin (100) |
| #devtools | [Enable&#124;Disable] - Manages Developer Tools (send no parameter for menu) | GMMgmt (200) |
| #disablerecipe | [Recipe ID] - Disables a Recipe | QuestTroupe (80) |
| #disarmtrap | Analog for ldon disarm trap for the newer clients since we still don't have it working. | QuestTroupe (80) |
| #distance | Reports the distance between you and your target. | QuestTroupe (80) |
| #door | Door editing command | QuestTroupe (80) |
| #doanim | [Animation ID&#124;Animation Name] [Speed] - Send an animation by ID or name at the specified speed to you or your target (Speed is optional) | Guide (50) |
| #dye | [slot&#124;'help'] [red] [green] [blue] [use_tint] - Dyes the specified armor slot to Red, Green, and Blue provided, allows you to bypass darkness limits. | ApprenticeGuide (20) |
| #dz | Manage expeditions and dynamic zone instances | QuestTroupe (80) |
| #dzkickplayers | Removes all players from current expedition. (/kickplayers alternative for pre-RoF clients) | Player (0) |
| #editmassrespawn | [name-search] [second-value] - Mass (Zone wide) NPC respawn timer editing command | GMAdmin (100) |
| #emote | [Name&#124;World&#124;Zone] [type] [message] - Send an emote message by name, to the world, or to your zone (^ separator allows multiple messages to be sent at once) | QuestTroupe (80) |
| #emotesearch | [Search Criteria] - Search for NPC Emotes | QuestTroupe (80) |
| #emoteview | Lists all NPC Emotes | QuestTroupe (80) |
| #emptyinventory | Clears your or your target's entire inventory (Equipment, General, Bank, and Shared Bank) | GMImpossible (250) |
| #enablerecipe | [Recipe ID] - Enables a Recipe | QuestTroupe (80) |
| #endurance | Restores your or your target's endurance. | Guide (50) |
| #equipitem | [slotid(0-21)] - Equip the item on your cursor into the specified slot | Guide (50) |
| #faction | [Find (criteria &#124; all ) &#124; Review (criteria &#124; all) &#124; Reset (id)] - Resets Player's Faction | QuestTroupe (80) |
| #factionassociation | [factionid] [amount] - triggers a faction hits via association | GMLeadAdmin (150) |
| #feature | Change your or your target's feature's temporarily | QuestTroupe (80) |
| #findaa | [Search Criteria] - Search for an AA | Guide (50) |
| #findaliases | [Search Criteria]- Searches for available command aliases, by alias or command | Player (0) |
| #findclass | [Search Criteria] - Search for a class | Guide (50) |
| #findfaction | [Search Criteria] - Search for a faction | Guide (50) |
| #findnpctype | [Search Criteria] - Search database NPC types | GMAdmin (100) |
| #findrace | [Search Criteria] - Search for a race | Guide (50) |
| #findrecipe | [Search Criteria] - Search for a recipe | Guide (50) |
| #findskill | [Search Criteria] - Search for a skill | Guide (50) |
| #findspell | [Search Criteria] - Search for a spell | Guide (50) |
| #findtask | [Search Criteria] - Search for a task | Guide (50) |
| #findzone | [Search Criteria] - Search database zones | GMAdmin (100) |
| #fixmob | [race&#124;gender&#124;texture&#124;helm&#124;face&#124;hair&#124;haircolor&#124;beard&#124;beardcolor&#124;heritage&#124;tattoo&#124;detail] [next&#124;prev] - Manipulate appearance of your target | QuestTroupe (80) |
| #flag | [Status] [Account Name] - Refresh your admin status, or set an account's Admin status if arguments provided | Player (0) |
| #flagedit | Edit zone flags on your target. Use #flagedit help for more info. | GMAdmin (100) |
| #flags | displays the Zone Flags of you or your target | Player (0) |
| #flymode | [0/1/2/3/4/5] - Set your or your player target's flymode to ground/flying/levitate/water/floating/levitate_running | Guide (50) |
| #fov | Check wether you're behind or in your target's field of view | QuestTroupe (80) |
| #freeze | Freeze your target | QuestTroupe (80) |
| #gassign | [Grid ID] - Assign targetted NPC to predefined wandering grid id | GMAdmin (100) |
| #gearup | Developer tool to quickly equip a character | GMMgmt (200) |
| #gender | [0/1/2] - Change your or your target's gender to male/female/neuter | Guide (50) |
| #getplayerburiedcorpsecount | Get your or your target's total number of buried player corpses. | GMAdmin (100) |
| #getvariable | [Variable Name] - Get the value of a variable from the database | GMMgmt (200) |
| #ginfo | get group info on target. | ApprenticeGuide (20) |
| #giveitem | [itemid] [charges] - Summon an item onto your target's cursor. Charges are optional. | GMMgmt (200) |
| #givemoney | [Platinum] [Gold] [Silver] [Copper] - Gives specified amount of money to you or your player target | GMMgmt (200) |
| #globalview | Lists all qglobals in cache if you were to do a quest with this target. | QuestTroupe (80) |
| #gm | [On&#124;Off] - Modify your or your target's GM Flag | QuestTroupe (80) |
| #gmspeed | [On&#124;Off] - Turn GM Speed On or Off for you or your player target | GMAdmin (100) |
| #gmzone | [Zone ID&#124;Zone Short Name] [Version] [Instance Identifier] - Zones to a private GM instance (Version defaults to 0 and Instance Identifier defaults to 'gmzone' if not used) | GMAdmin (100) |
| #godmode | [on/off] - Turns on/off hideme, gmspeed, invul, and flymode. | GMMgmt (200) |
| #goto | [playername] or [x y z] [h] - Teleport to the provided coordinates or to your target | Steward (10) |
| #grid | [add/delete] [grid_num] [wandertype] [pausetype] - Create/delete a wandering grid | GMAreas (170) |
| #guild | Guild manipulation commands. Use argument help for more info. | Steward (10) |
| #guildapprove | [guildapproveid] - Approve a guild with specified ID (guild creator receives the id) | Player (0) |
| #guildcreate | [guildname] - Creates an approval setup for guild name specified | Player (0) |
| #guildlist | [guildapproveid] - Lists character names who have approved the guild specified by the approve id | Player (0) |
| #haste | [percentage] - Set your haste percentage | GMAdmin (100) |
| #hatelist | Display hate list for NPC. | QuestTroupe (80) |
| #heal | Completely heal your target | Steward (10) |
| #help | [Search Criteria] - List available commands and their description, specify partial command as argument to search | Player (0) |
| #heromodel | [hero model] [slot] - Full set of Hero's Forge Armor appearance. If slot is set, sends exact model just to slot. | GMMgmt (200) |
| #hideme | [on/off] - Hide yourself from spawn lists. | QuestTroupe (80) |
| #hotfix | [hotfix_name] - Reloads shared memory into a hotfix, equiv to load_shared_memory followed by apply_shared_memory | GMImpossible (250) |
| #hp | Refresh your HP bar from the server. | Player (0) |
| #incstat | Increases or Decreases a client's stats permanently. | GMMgmt (200) |
| #instance | Modify Instances | GMMgmt (200) |
| #interrogateinv | use [help] argument for available options | Player (0) |
| #interrupt | [message id] [color] - Interrupt your casting. Arguments are optional. | Guide (50) |
| #invsnapshot | Manipulates inventory snapshots for your current target | QuestTroupe (80) |
| #invul | [On&#124;Off]] - Turn player target's or your invulnerable flag on or off | QuestTroupe (80) |
| #ipban | [IP] - Ban IP | GMMgmt (200) |
| #iplookup | [charname] - Look up IP address of charname | GMMgmt (200) |
| #iteminfo | Get information about the item on your cursor | Steward (10) |
| #itemsearch | [Search Criteria] - Search for an item | Steward (10) |
| #kick | [Character Name] - Disconnect a player by name | GMLeadAdmin (150) |
| #kill | Kill your target | GMAdmin (100) |
| #killallnpcs | [npc_name] - Kills all npcs by search name, leave blank for all attackable NPC's | GMMgmt (200) |
| #lastname | [Last Name] - Set your or your player target's last name (use \-1\ to remove last name) | Guide (50) |
| #level | [Level] - Set your target's level | Steward (10) |
| #list | [npcs&#124;players&#124;corpses&#124;doors&#124;objects] [search] - Search entities | ApprenticeGuide (20) |
| #listpetition | List petitions | Guide (50) |
| #lootsim | [npc_type_id] [loottable_id] [iterations] - Runs benchmark simulations using real loot logic to report numbers and data | GMImpossible (250) |
| #load_shared_memory | [shared_memory_name] - Reloads shared memory and uses the input as output | GMImpossible (250) |
| #loc | Print out your or your target's current location and heading | Player (0) |
| #logs | Manage anything to do with logs | GMImpossible (250) |
| #makepet | [Pet Name] - Make a pet | Guide (50) |
| #mana | Fill your or your target's mana | Guide (50) |
| #maxskills | Maxes skills for you. | GMMgmt (200) |
| #memspell | [Spell ID] [Spell Gem] - Memorize a Spell by ID to the specified Spell Gem for you or your target | Guide (50) |
| #merchant_close_shop | Closes a merchant shop | GMAdmin (100) |
| #merchant_open_shop | Opens a merchants shop | GMAdmin (100) |
| #modifynpcstat | [Stat] [Value] - Modifies an NPC's stats temporarily. | GMLeadAdmin (150) |
| #motd | [Message of the Day] - Set Message of the Day (leave empty to have no Message of the Day) | GMLeadAdmin (150) |
| #movechar | [Character ID&#124;Character Name] [Zone ID&#124;Zone Short Name] - Move an offline character to the specified zone | Guide (50) |
| #movement | Various movement commands | GMMgmt (200) |
| #myskills | Show details about your current skill levels | Player (0) |
| #mysql | [Help&#124;Query] [SQL Query] - Mysql CLI, see 'Help' for options. | GMImpossible (250) |
| #mystats | Show details about you or your pet | Guide (50) |
| #name | [New Name] - Rename your player target | GMLeadAdmin (150) |
| #netstats | Gets the network stats for a stream. | GMMgmt (200) |
| #network | Admin commands for the udp network interface. | GMImpossible (250) |
| #npccast | [targetname/entityid] [spellid] - Causes NPC target to cast spellid on targetname/entityid | QuestTroupe (80) |
| #npcedit | [column] [value] - Mega NPC editing command | GMAdmin (100) |
| #npceditmass | [name-search] [column] [value] - Mass (Zone wide) NPC data editing command | GMAdmin (100) |
| #npcemote | [Message] - Make your NPC target emote a message. | GMLeadAdmin (150) |
| #npcloot | Manipulate the loot an NPC is carrying. Use #npcloot help for more information. | QuestTroupe (80) |
| #npcsay | [Message] - Make your NPC target say a message. | GMLeadAdmin (150) |
| #npcshout | [Message] - Make your NPC target shout a message. | GMLeadAdmin (150) |
| #npcspawn | [create/add/update/remove/delete] - Manipulate spawn DB | GMAreas (170) |
| #npcstats | Show stats about target NPC | QuestTroupe (80) |
| #npctypespawn | [NPC ID] [Faction ID] - Spawn an NPC by ID from the database with an option of setting its Faction ID | Steward (10) |
| #nudge | Nudge your target's current position by specific values | QuestTroupe (80) |
| #nukebuffs | [Beneficial&#124;Detrimental&#124;Help] - Strip all buffs by type on you or your target (no argument to remove all buffs) | Guide (50) |
| #nukeitem | [Item ID] - Removes the specified Item ID from you or your player target's inventory | GMLeadAdmin (150) |
| #object | List&#124;Add&#124;Edit&#124;Move&#124;Rotate&#124;Copy&#124;Save&#124;Undo&#124;Delete - Manipulate static and tradeskill objects within the zone | GMAdmin (100) |
| #oocmute | [0&#124;1] - Enable or Disable Server OOC | GMMgmt (200) |
| #opcode | Reloads all server patches | GMImpossible (250) |
| #path | view and edit pathing | GMMgmt (200) |
| #peekinv | [equip/gen/cursor/poss/limbo/curlim/trib/bank/shbank/allbank/trade/world/all] - Print out contents of your player target's inventory | GMAdmin (100) |
| #peqzone | [Zone ID&#124;Zone Short Name] - Teleports you to the specified zone if you meet the requirements. | Player (0) |
| #peqzone_flags | displays the PEQZone Flags of you or your target | Player (0) |
| #permaclass | [Class ID] - Change your or your player target's class, changed client is disconnected | QuestTroupe (80) |
| #permagender | [Gender ID] - Change your or your player target's gender | QuestTroupe (80) |
| #permarace | [Race ID] - Change your or your player target's race | QuestTroupe (80) |
| #petitems | View your pet's items if you have one | ApprenticeGuide (20) |
| #petitioninfo | [petition number] - Get info about a petition | ApprenticeGuide (20) |
| #picklock | Analog for ldon pick lock for the newer clients since we still don't have it working. | Player (0) |
| #profanity | Manage censored language. | GMLeadAdmin (150) |
| #push | [Back Push] [Up Push] - Lets you do spell push on an NPC | GMLeadAdmin (150) |
| #proximity | Shows NPC proximity | GMLeadAdmin (150) |
| #pvp | [On&#124;Off] - Set you or your player target's PVP status | GMAdmin (100) |
| #qglobal | [On&#124;Off&#124;View] - Toggles quest global functionality for your NPC target | GMAdmin (100) |
| #questerrors | Shows quest errors. | GMAdmin (100) |
| #race | [racenum] - Change your or your target's race. Use racenum 0 to return to normal | Guide (50) |
| #raidloot | [All&#124;GroupLeader&#124;RaidLeader&#124;Selected] - Sets your Raid Loot Type if you have permission to do so. | Player (0) |
| #randomfeatures | Temporarily randomizes the Facial Features of your target | QuestTroupe (80) |
| #refreshgroup | Refreshes Group for you or your player target. | Player (0) |
| #reload | Reloads different types of server data globally, use no argument for help menu. | GMMgmt (200) |
| #removeitem | [Item ID] [Amount] - Removes the specified Item ID by Amount from you or your player target's inventory (Amount defaults to 1 if not used) | GMAdmin (100) |
| #repop | [Force] - Repop the zone with optional force repop | GMAdmin (100) |
| #resetaa | Resets a Player's AA in their profile and refunds spent AA's to unspent, may disconnect player. | GMMgmt (200) |
| #resetaa_timer | [All&#124;Timer ID] - Command to reset AA cooldown timers for you or your player target. | GMMgmt (200) |
| #resetdisc_timer | [All&#124;Timer ID] - Command to reset discipline timers. | GMMgmt (200) |
| #revoke | [Character Name] [0&#124;1] - Revokes or unrevokes a player's ability to talk in OOC by name (0 = Unrevoke, 1 = Revoke) | GMMgmt (200) |
| #roambox | [Remove&#124;Set] [Box Size] [Delay (Milliseconds)] - Remove or set an NPC's roambox size and delay | GMMgmt (200) |
| #rules | (subcommand) - Manage server rules | GMImpossible (250) |
| #save | Force your player or player corpse target to be saved to the database | Guide (50) |
| #scale | Handles npc scaling | GMLeadAdmin (150) |
| #scribespell | [Spell ID] - Scribe a spell by ID to your or your target's spell book. | GMCoder (180) |
| #scribespells | [Max level] [Min level] - Scribe all spells for you or your player target that are usable by them, up to level specified. (may freeze client for a few seconds) | GMLeadAdmin (150) |
| #sendzonespawns | Refresh spawn list for all clients in zone | GMLeadAdmin (150) |
| #sensetrap | Analog for ldon sense trap for the newer clients since we still don't have it working. | Player (0) |
| #serverinfo | Get CPU, Operating System, and Process Information about the server | GMMgmt (200) |
| #serverlock | [0&#124;1] - Lock or Unlock the World Server (0 = Unlocked, 1 = Locked) | GMLeadAdmin (150) |
| #serverrules | Read this server's rules | Player (0) |
| #setaapts | [AA&#124;Group&#124;Raid] [AA Amount] - Set your or your player target's Available AA Points by Type | GMAdmin (100) |
| #setaaxp | [AA&#124;Group&#124;Raid] [AA Experience] - Set your or your player target's AA Experience by Type | GMAdmin (100) |
| #setadventurepoints | [Theme] [Points] - Set your or your player target's available Adventure Points by Theme | GMLeadAdmin (150) |
| #setaltcurrency | [Currency ID] [Amount] - Set your or your target's available Alternate Currency by Currency ID | GMAdmin (100) |
| #setanim | [Animation ID (IDs are 0 to 4)] - Set target's appearance to Animation ID | GMMgmt (200) |
| #setcrystals | [value] - Set your or your player target's available radiant or ebon crystals | GMAdmin (100) |
| #setendurance | [Endurance] - Set your or your target's Endurance | GMAdmin (100) |
| #setfaction | [Faction ID] - Sets targeted NPC's faction in the database | GMAreas (170) |
| #sethp | [Health] - Set your or your target's Health | GMAdmin (100) |
| #setlanguage | [Language ID] [Value] - Set your or your target's Language by ID to Value | Guide (50) |
| #setlsinfo | [Email] [Password] - Set loginserver email address and password (if supported by loginserver) | Steward (10) |
| #setmana | [Mana] - Set your or your target's Mana | GMAdmin (100) |
| #setpass | [Account Name] [Password] - Set local password by account name | GMLeadAdmin (150) |
| #setpvppoints | [Amount] - Set your or your player target's PVP points | GMAdmin (100) |
| #setskill | [skillnum] [value] - Set your target's skill skillnum to value | Guide (50) |
| #setskillall | [Skill Level] - Set all of your or your target's skills to the specified skill level | Guide (50) |
| #setstartzone | [Zone ID&#124;Zone Short Name] - Sets your or your target's starting zone (Use '0' or 'Reset' to allow the player use of /setstartcity) | QuestTroupe (80) |
| #setstat | Sets the stats to a specific value. | Max (255) |
| #setxp | [value] - Set your or your player target's experience | GMAdmin (100) |
| #showbonusstats | [item&#124;spell&#124;all] Shows bonus stats for target from items or spells. Shows both by default. | Guide (50) |
| #showbuffs | List buffs active on your target or you if no target | Guide (50) |
| #shownumhits | Shows buffs numhits for yourself. | Player (0) |
| #shownpcgloballoot | Show global loot entries for your target NPC | Guide (50) |
| #showskills | [Start Skill ID] [All] - Show the values of your or your player target's skills in a popup 50 at a time, use 'all' as second argument to show non-usable skill's values | Guide (50) |
| #showspellslist | Shows spell list of targeted NPC | GMAdmin (100) |
| #showstats | Show details about you or your target | Guide (50) |
| #showzonegloballoot | Show global loot entries for your current zone | Guide (50) |
| #showzonepoints | Show zone points for current zone | Guide (50) |
| #shutdown | Shut this zone process down | GMLeadAdmin (150) |
| #spawn | [name] [race] [level] [material] [hp] [gender] [class] [priweapon] [secweapon] [merchantid] - Spawn an NPC | Steward (10) |
| #spawneditmass | [Search Criteria] [Edit Option] [Edit Value] [Apply] Mass editing spawn command (Apply is optional, 0 = False, 1 = True, default is False) | GMLeadAdmin (150) |
| #spawnfix | Find targeted NPC in database based on its X/Y/heading and update the database to make it spawn at your current location/heading. | GMAreas (170) |
| #spawnstatus | [All&#124;Disabled&#124;Enabled&#124;Spawn ID] - Show respawn timer status | GMAdmin (100) |
| #spellinfo | [spellid] - Get detailed info about a spell | Steward (10) |
| #stun | [duration] - Stuns you or your target for duration | GMAdmin (100) |
| #summon | [Character Name] - Summons your corpse, NPC, or player target, or by character name if specified | QuestTroupe (80) |
| #summonburiedplayercorpse | Summons the target's oldest buried corpse, if any exist. | GMAdmin (100) |
| #summonitem | [itemid] [charges] - Summon an item onto your cursor. Charges are optional. | GMMgmt (200) |
| #suspend | [name] [days] [reason] - Suspend by character name and for specificed number of days | GMLeadAdmin (150) |
| #task | (subcommand) - Task system commands | GMLeadAdmin (150) |
| #tempname | [newname] - Temporarily renames your target. Leave name blank to restore the original name. | GMAdmin (100) |
| #petname | [newname] - Temporarily renames your pet. Leave name blank to restore the original name. | GMAdmin (100) |
| #texture | [Texture] [Helmet Texture] - Change your or your target's texture (Helmet Texture defaults to 0 if not used) | Steward (10) |
| #time | [Hour] [Minute] - Set world time to specified time | EQSupport (90) |
| #timers | Display persistent timers for target | GMMgmt (200) |
| #timezone | [Hour] [Minutes] - Set timezone (Minutes are optional) | EQSupport (90) |
| #title | [Title] - Set your or your player target's title (use \-1\ to remove title) | Guide (50) |
| #titlesuffix | [Title Suffix] - Set your or your player target's title suffix (use \-1\ to remove title suffix) | Guide (50) |
| #traindisc | [level] - Trains all the disciplines usable by the target, up to level specified. (may freeze client for a few seconds) | GMLeadAdmin (150) |
| #trapinfo | Gets infomation about the traps currently spawned in the zone. | QuestTroupe (80) |
| #tune | Calculate statistical values related to combat. | GMAdmin (100) |
| #undye | Remove dye from all of your or your target's armor slots | GMAdmin (100) |
| #undyeme | Remove dye from all of your armor slots | Player (0) |
| #unfreeze | Unfreeze your target | QuestTroupe (80) |
| #unmemspell | [Spell ID] - Unmemorize a Spell by ID for you or your target | Guide (50) |
| #unmemspells |  Unmemorize all spells for you or your target | Guide (50) |
| #unscribespell | [Spell ID] - Unscribe a spell from your or your target's spell book by Spell ID | GMCoder (180) |
| #unscribespells | Clear out your or your player target's spell book. | GMCoder (180) |
| #untraindisc | [Spell ID] - Untrain your or your target's discipline by Spell ID | GMCoder (180) |
| #untraindiscs | Untrains all disciplines from your target. | GMCoder (180) |
| #updatechecksum | update client checksum | GMImpossible (250) |
| #uptime | [zone server id] - Get uptime of worldserver, or zone server if argument provided | Steward (10) |
| #version | Display current version of EQEmu server | Player (0) |
| #viewcurrencies | View your or your target's currencies | GMAdmin (100) |
| #viewnpctype | [NPC ID] - Show stats for an NPC by NPC ID | GMAdmin (100) |
| #viewpetition | [petition number] - View a petition | ApprenticeGuide (20) |
| #viewrecipe | [Recipe ID] - Show a recipe's entries | GMAdmin (100) |
| #viewzoneloot | [item id] - Allows you to search a zone's loot for a specific item ID. (0 shows all loot in the zone) | QuestTroupe (80) |
| #wc | [wear slot] [material] - Sends an OP_WearChange for your target | GMMgmt (200) |
| #weather | [0/1/2/3] (Off/Rain/Snow/Manual) - Change the weather | QuestTroupe (80) |
| #who | [search] | ApprenticeGuide (20) |
| #worldshutdown | Shut down world and all zones | GMMgmt (200) |
| #wp | [add&#124;delete] [grid_id] [pause] [waypoint_id] [-h] - Add or delete a waypoint by grid ID. (-h to use current heading) | GMAreas (170) |
| #wpadd | [pause] [-h] - Add your current location as a waypoint to your NPC target's AI path. (-h to use current heading) | GMAreas (170) |
| #wpinfo | Show waypoint info about your NPC target | GMAreas (170) |
| #worldwide | Performs world-wide GM functions such as cast (can be extended for other commands). Use caution | GMImpossible (250) |
| #xtargets | [New Max XTargets] - Show your or your target's XTargets and optionally set max XTargets. | GMImpossible (250) |
| #zclip | [Minimum Clip] [Maximum Clip] [Fog Minimum Clip] [Fog Maximum Clip] [Permanent (0 = False, 1 = True)] - Change zone clipping | QuestTroupe (80) |
| #zcolor | [Red] [Green] [Blue] [Permanent (0 = False, 1 = True)] - Change sky color | QuestTroupe (80) |
| #zheader | [Zone ID&#124;Zone Short Name] [Version] - Load a zone header from the database | QuestTroupe (80) |
| #zone | [Zone ID&#124;Zone Short Name] [X] [Y] [Z] - Teleport to specified Zone by ID or Short Name (coordinates are optional) | Guide (50) |
| #zonebootup | [ZoneServerID] [shortname] - Make a zone server boot a specific zone | GMLeadAdmin (150) |
| #zoneinstance | [Instance ID] [X] [Y] [Z] - Teleport to specified Instance by ID (coordinates are optional) | Guide (50) |
| #zonelock | [List&#124;Lock&#124;Unlock] [Zone ID&#124;Zone Short Name] - Set or get lock status of a Zone by ID or Short Name | GMAdmin (100) |
| #zoneshutdown | [shortname] - Shut down a zone server | GMLeadAdmin (150) |
| #zonestatus | Show connected zoneservers, synonymous with /servers | GMLeadAdmin (150) |
| #zopp | Troubleshooting command - Sends a fake item packet to you. No server reference is created. | GMImpossible (250) |
| #zsafecoords | [X] [Y] [Z] [Heading] [Permanent (0 = False, 1 = True)] - Set the current zone's safe coordinates | QuestTroupe (80) |
| #zsave |  Saves zheader to the database | QuestTroupe (80) |
| #zsky | [Sky Type] [Permanent (0 = False, 1 = True)] - Change zone sky type | QuestTroupe (80) |
| #zstats | Show info about zone header | QuestTroupe (80) |
| #zunderworld | [Z] [Permanent (0 = False, 1 = True)] - Change zone underworld Z | QuestTroupe (80) |
