!!! info end

    This page lists the commands that are available in-game, based on assigned Account Status, for your EQEmu Server.
| Command | Description | Status Level |
| :--- | :--- | :--- |
| #acceptrules | [acceptrules] - Accept the EQEmu Agreement | Player (0) |
| #advnpcspawn | [maketype&#124;makegroup&#124;addgroupentry&#124;addgroupspawn][removegroupspawn&#124;movespawn&#124;editgroupbox&#124;cleargroupbox] | GMLeadAdmin (150) |
| #aggrozone | [aggro] - Aggro every mob in the zone with X aggro. Default is 0. Not recommend if you're not invulnerable. | GMAdmin (100) |
| #ai | [factionid/spellslist/con/guard/roambox/stop/start] - Modify AI on NPC target | GMAdmin (100) |
| #appearance | [type] [value] - Send an appearance packet for you or your target | GMLeadAdmin (150) |
| #appearanceeffects | [Help&#124;Remove&#124;Set&#124;View] - Modify appearance effects on yourself or your target. | GMAdmin (100) |
| #apply_shared_memory | [shared_memory_name] - Tells every zone and world to apply a specific shared memory segment by name. | GMImpossible (250) |
| #attack | [Entity Name] - Make your NPC target attack an entity by name | GMLeadAdmin (150) |
| #augmentitem | Force augments an item. Must have the augment item window open. | GMImpossible (250) |
| #ban | [Character Name] [Reason] - Ban by character name | GMLeadAdmin (150) |
| #bugs | [Close&#124;Delete&#124;Review&#124;Search&#124;View] - Handles player bug reports | QuestTroupe (80) |
| #bot | Type \#bot help\ or \^help\ to the see the list of available commands for bots. | Player (0) |
| #camerashake | [Duration (Milliseconds)] [Intensity (1-10)] - Shakes the camera on everyone's screen globally. | QuestTroupe (80) |
| #castspell | [Spell ID] [Instant (0 = False, 1 = True, Default is 1 if Unused)] - Cast a spell | Guide (50) |
| #chat | [Channel ID] [Message] - Send a channel message to all zones | GMMgmt (200) |
| #clearxtargets | Clears XTargets | Player (0) |
| #copycharacter | [source_char_name] [dest_char_name] [dest_account_name] - Copies character to destination account | GMImpossible (250) |
| #corpse | Manipulate corpses, use with no arguments for help | Guide (50) |
| #corpsefix | Attempts to bring corpses from underneath the ground within close proximity of the player | Player (0) |
| #countitem | [Item ID] - Counts the specified Item ID in your or your target's inventory | GMLeadAdmin (150) |
| #damage | [Amount] - Damage yourself or your target | GMAdmin (100) |
| #databuckets | View&#124;Delete [key] [limit]- View data buckets, limit 50 default or Delete databucket by key | QuestTroupe (80) |
| #dbspawn2 | [Spawngroup ID] [Respawn] [Variance] [Condition ID] [Condition Minimum] - Spawn an NPC from a predefined row in the spawn2 table, Respawn and Variance are in Seconds (condition is optional) | GMAdmin (100) |
| #delacct | [Account ID&#124;Account Name] - Delete an account by ID or Name | GMLeadAdmin (150) |
| #delpetition | [petition number] - Delete a petition | ApprenticeGuide (20) |
| #depop | [Start Spawn Timer] - Depop your NPC target and optionally start their spawn timer (false by default) | Guide (50) |
| #depopzone | [Start Spawn Timers] - Depop the zone and optionally start spawn timers (false by default) | GMAdmin (100) |
| #devtools | [menu&#124;window] [enable&#124;disable] - Manages Developer Tools (send no parameter for menu) | GMMgmt (200) |
| #disablerecipe | [Recipe ID] - Disables a Recipe | QuestTroupe (80) |
| #disarmtrap | Analog for ldon disarm trap for the newer clients since we still don't have it working. | QuestTroupe (80) |
| #door | Door editing command | QuestTroupe (80) |
| #doanim | [Animation ID&#124;Animation Name] [Speed] - Send an animation by ID or name at the specified speed to you or your target (Speed is optional) | Guide (50) |
| #dye | [slot&#124;'help'] [red] [green] [blue] [use_tint] - Dyes the specified armor slot to Red, Green, and Blue provided, allows you to bypass darkness limits. | ApprenticeGuide (20) |
| #dz | Manage expeditions and dynamic zone instances | QuestTroupe (80) |
| #dzkickplayers | Removes all players from current expedition. (/kickplayers alternative for pre-RoF clients) | Player (0) |
| #editmassrespawn | [name-search] [second-value] - Mass (Zone wide) NPC respawn timer editing command | GMAdmin (100) |
| #emote | [Name&#124;World&#124;Zone] [type] [message] - Send an emote message by name, to the world, or to your zone (^ separator allows multiple messages to be sent at once) | QuestTroupe (80) |
| #emptyinventory | Clears your or your target's entire inventory (Equipment, General, Bank, and Shared Bank) | GMImpossible (250) |
| #enablerecipe | [Recipe ID] - Enables a Recipe | QuestTroupe (80) |
| #entityvariable | [clear&#124;delete&#124;set&#124;view] - Modify entity variables for yourself or your target | GMAdmin (100) |
| #exptoggle | [Toggle] - Toggle your or your target's experience gain. | QuestTroupe (80) |
| #faction | [Find (criteria &#124; all ) &#124; Review (criteria &#124; all) &#124; Reset (id)] - Resets Player's Faction | QuestTroupe (80) |
| #factionassociation | [factionid] [amount] - triggers a faction hits via association | GMLeadAdmin (150) |
| #feature | Change your or your target's feature's temporarily | QuestTroupe (80) |
| #size | Change your targets size (alias of #feature size) | QuestTroupe (80) |
| #find | Search command used to find various things | Guide (50) |
| #fish | Fish for an item | QuestTroupe (80) |
| #fixmob | [race&#124;gender&#124;texture&#124;helm&#124;face&#124;hair&#124;haircolor&#124;beard&#124;beardcolor&#124;heritage&#124;tattoo&#124;detail] [next&#124;prev] - Manipulate appearance of your target | QuestTroupe (80) |
| #flagedit | Edit zone flags on your target. Use #flagedit help for more info. | GMAdmin (100) |
| #fleeinfo | - Gives info about whether a NPC will flee or not, using the command issuer as top hate. | QuestTroupe (80) |
| #forage | Forage an item | QuestTroupe (80) |
| #gearup | Developer tool to quickly equip yourself or your target | GMMgmt (200) |
| #giveitem | [itemid] [charges] - Summon an item onto your target's cursor. Charges are optional. | GMMgmt (200) |
| #givemoney | [Platinum] [Gold] [Silver] [Copper] - Gives specified amount of money to you or your player target | GMMgmt (200) |
| #gmzone | [Zone ID&#124;Zone Short Name] [Version] [Instance Identifier] - Zones to a private GM instance (Version defaults to 0 and Instance Identifier defaults to 'gmzone' if not used) | GMAdmin (100) |
| #goto | [playername] or [x y z] [h] - Teleport to the provided coordinates or to your target | Steward (10) |
| #grantaa | [level] - Grants a player all available AA points up the specified level, all AAs are granted if no level is specified. | GMMgmt (200) |
| #grid | [add/delete] [grid_num] [wandertype] [pausetype] - Create/delete a wandering grid | GMAreas (170) |
| #guild | Guild manipulation commands. Use argument help for more info. | Steward (10) |
| #help | [Search Criteria] - List available commands and their description, specify partial command as argument to search | Player (0) |
| #hotfix | [hotfix_name] - Reloads shared memory into a hotfix, equiv to load_shared_memory followed by apply_shared_memory | GMImpossible (250) |
| #hp | Refresh your HP bar from the server. | Player (0) |
| #instance | Modify Instances | GMMgmt (200) |
| #interrogateinv | use [help] argument for available options | Player (0) |
| #interrupt | [Message ID] [Color] - Interrupt your casting. Arguments are optional. | Guide (50) |
| #invsnapshot | Manipulates inventory snapshots for your current target | QuestTroupe (80) |
| #ipban | [IP] - Ban IP | GMMgmt (200) |
| #kick | [Character Name] - Disconnect a player by name | GMLeadAdmin (150) |
| #kill | Kill your target | GMAdmin (100) |
| #killallnpcs | [npc_name] - Kills all npcs by search name, leave blank for all attackable NPC's | GMMgmt (200) |
| #list | [npcs&#124;players&#124;corpses&#124;doors&#124;objects] [search] - Search entities | ApprenticeGuide (20) |
| #lootsim | [npc_type_id] [loottable_id] [iterations] - Runs benchmark simulations using real loot logic to report numbers and data | GMImpossible (250) |
| #load_shared_memory | [shared_memory_name] - Reloads shared memory and uses the input as output | GMImpossible (250) |
| #loc | Print out your or your target's current location and heading | Player (0) |
| #logs | Manage anything to do with logs | GMImpossible (250) |
| #makepet | [Pet Name] - Make a pet | Guide (50) |
| #memspell | [Spell ID] [Spell Gem] - Memorize a Spell by ID to the specified Spell Gem for you or your target | Guide (50) |
| #merchantshop | Closes or opens your target merchant's shop | GMAdmin (100) |
| #modifynpcstat | [Stat] [Value] - Modifies an NPC's stats temporarily. | GMLeadAdmin (150) |
| #movechar | [Character ID&#124;Character Name] [Zone ID&#124;Zone Short Name] - Move an offline character to the specified zone | Guide (50) |
| #movement | Various movement commands | GMMgmt (200) |
| #myskills | Show details about your current skill levels | Player (0) |
| #mysql | [Help&#124;Query] [SQL Query] - Mysql CLI, see 'Help' for options. | GMImpossible (250) |
| #mystats | Show details about you or your pet | Guide (50) |
| #npccast | [targetname/entityid] [spellid] - Causes NPC target to cast spellid on targetname/entityid | QuestTroupe (80) |
| #npcedit | [column] [value] - Mega NPC editing command | GMAdmin (100) |
| #npceditmass | [name-search] [column] [value] - Mass (Zone wide) NPC data editing command | GMAdmin (100) |
| #npcemote | [Message] - Make your NPC target emote a message. | GMLeadAdmin (150) |
| #npcloot | Manipulate the loot an NPC is carrying. Use #npcloot help for more information. | QuestTroupe (80) |
| #npcsay | [Message] - Make your NPC target say a message. | GMLeadAdmin (150) |
| #npcshout | [Message] - Make your NPC target shout a message. | GMLeadAdmin (150) |
| #npcspawn | [create/add/update/remove/delete] - Manipulate spawn DB | GMAreas (170) |
| #npctypespawn | [NPC ID] [Faction ID] - Spawn an NPC by ID from the database with an option of setting its Faction ID | Steward (10) |
| #nudge | Nudge your target's current position by specific values | QuestTroupe (80) |
| #nukebuffs | [Beneficial&#124;Detrimental&#124;Help] - Strip all buffs by type on you or your target (no argument to remove all buffs) | Guide (50) |
| #nukeitem | [Item ID] - Removes the specified Item ID from you or your player target's inventory | GMLeadAdmin (150) |
| #object | List&#124;Add&#124;Edit&#124;Move&#124;Rotate&#124;Copy&#124;Save&#124;Undo&#124;Delete - Manipulate static and tradeskill objects within the zone | GMAdmin (100) |
| #opcode | Reloads all opcodes from server patch files | GMMgmt (200) |
| #parcels | View and edit the parcel system.  Requires parcels to be enabled in rules. | GMMgmt (200) |
| #path | view and edit pathing | GMMgmt (200) |
| #peqzone | [Zone ID&#124;Zone Short Name] - Teleports you to the specified zone if you meet the requirements. | Player (0) |
| #petitems | View your pet's items if you have one | ApprenticeGuide (20) |
| #picklock | Analog for ldon pick lock for the newer clients since we still don't have it working. | Player (0) |
| #profanity | Manage censored language. | GMLeadAdmin (150) |
| #push | [Back Push] [Up Push] - Lets you do spell push on an NPC | GMLeadAdmin (150) |
| #raidloot | [All&#124;GroupLeader&#124;RaidLeader&#124;Selected] - Sets your Raid Loot Type if you have permission to do so. | Player (0) |
| #randomfeatures | Temporarily randomizes the Facial Features of your target | QuestTroupe (80) |
| #refreshgroup | Refreshes Group for you or your player target. | Player (0) |
| #reload | Reloads different types of server data globally, use no argument for help menu. | GMMgmt (200) |
| #rq | Reloads quests (alias of #reload quests). | GMMgmt (200) |
| #rl | Reloads logs (alias of #reload logs). | GMMgmt (200) |
| #removeitem | [Item ID] [Amount] - Removes the specified Item ID by Amount from you or your player target's inventory (Amount defaults to 1 if not used) | GMAdmin (100) |
| #repop | [Force] - Repop the zone with optional force repop | GMAdmin (100) |
| #resetaa | [aa&#124;leadership] - Resets a player's AAs or Leadership AAs and refunds spent AAs (not Leadership AAs) to unspent, may disconnect player. | GMMgmt (200) |
| #resetaa_timer | [All&#124;Timer ID] - Command to reset AA cooldown timers for you or your player target. | GMMgmt (200) |
| #resetdisc_timer | [All&#124;Timer ID] - Command to reset discipline timers. | GMMgmt (200) |
| #revoke | [Character Name] [0&#124;1] - Revokes or unrevokes a player's ability to talk in OOC by name (0 = Unrevoke, 1 = Revoke) | GMMgmt (200) |
| #roambox | [Remove&#124;Set] [Box Size] [Delay (Milliseconds)] - Remove or set an NPC's roambox size and delay | GMMgmt (200) |
| #rules | (subcommand) - Manage server rules | GMImpossible (250) |
| #save | Force your player or player corpse target to be saved to the database | Guide (50) |
| #scale | Handles NPC scaling | GMLeadAdmin (150) |
| #scribespell | [Spell ID] - Scribe a spell by ID to your or your target's spell book. | GMCoder (180) |
| #scribespells | [Max level] [Min level] - Scribe all spells for you or your player target that are usable by them, up to level specified. (may freeze client for a few seconds) | GMLeadAdmin (150) |
| #sendzonespawns | Refresh spawn list for all clients in zone | GMLeadAdmin (150) |
| #sensetrap | Analog for ldon sense trap for the newer clients since we still don't have it working. | Player (0) |
| #serverrules | Show server rules | Player (0) |
| #set | Set command used to set various things | Guide (50) |
| #show | Show command used to show various things | Guide (50) |
| #shutdown | Shut this zone process down | GMLeadAdmin (150) |
| #spawn | [name] [race] [level] [material] [hp] [gender] [class] [priweapon] [secweapon] [merchantid] - Spawn an NPC | Steward (10) |
| #spawneditmass | [Search Criteria] [Edit Option] [Edit Value] [Apply] Mass editing spawn command (Apply is optional, 0 = False, 1 = True, default is False) | GMLeadAdmin (150) |
| #spawnfix | Find targeted NPC in database based on its X/Y/heading and update the database to make it spawn at your current location/heading. | GMAreas (170) |
| #stun | [duration] - Stuns you or your target for duration | GMAdmin (100) |
| #summon | [Character Name] - Summons your corpse, NPC, or player target, or by character name if specified | QuestTroupe (80) |
| #summonburiedplayercorpse | Summons the target's oldest buried corpse, if any exist. | GMAdmin (100) |
| #summonitem | [itemid] [charges] - Summon an item onto your cursor. Charges are optional. | GMMgmt (200) |
| #suspend | [name] [days] [reason] - Suspend by character name and for specificed number of days | GMLeadAdmin (150) |
| #suspendmulti | [Character Name One&#124;Character Name Two&#124;etc] [Days] [Reason] - Suspend multiple characters by name for specified number of days | GMLeadAdmin (150) |
| #takeplatinum | [Platinum] - Takes specified amount of platinum from you or your player target | GMMgmt (200) |
| #task | (subcommand) - Task system commands | GMLeadAdmin (150) |
| #petname | [newname] - Temporarily renames your pet. Leave name blank to restore the original name. | GMAdmin (100) |
| #traindisc | [level] - Trains all the disciplines usable by the target, up to level specified. (may freeze client for a few seconds) | GMLeadAdmin (150) |
| #tune | Calculate statistical values related to combat. | GMAdmin (100) |
| #undye | Remove dye from all of your or your target's armor slots | GMAdmin (100) |
| #unmemspell | [Spell ID] - Unmemorize a Spell by ID for you or your target | Guide (50) |
| #unmemspells |  Unmemorize all spells for you or your target | Guide (50) |
| #unscribespell | [Spell ID] - Unscribe a spell from your or your target's spell book by Spell ID | GMCoder (180) |
| #unscribespells | Clear out your or your player target's spell book. | GMCoder (180) |
| #untraindisc | [Spell ID] - Untrain your or your target's discipline by Spell ID | GMCoder (180) |
| #untraindiscs | Untrains all disciplines from your target. | GMCoder (180) |
| #wc | [Slot ID] [Material] [Hero Forge Model] [Elite Material] - Sets the specified slot for you or your target to a material, Hero Forge Model and Elite Material are optional | GMMgmt (200) |
| #worldshutdown | Shut down world and all zones | GMMgmt (200) |
| #wp | [add&#124;delete] [grid_id] [pause] [waypoint_id] [-h] - Add or delete a waypoint by grid ID. (-h to use current heading) | GMAreas (170) |
| #wpadd | [pause] [-h] - Add your current location as a waypoint to your NPC target's AI path. (-h to use current heading) | GMAreas (170) |
| #worldwide | Performs world-wide GM functions such as cast (can be extended for other commands). Use caution | GMImpossible (250) |
| #zone | [Zone ID&#124;Zone Short Name] [X] [Y] [Z] - Teleport to specified Zone by ID or Short Name (coordinates are optional) | Guide (50) |
| #zonebootup | [ZoneServerID] [shortname] - Make a zone server boot a specific zone | GMLeadAdmin (150) |
| #zoneinstance | [Instance ID] [X] [Y] [Z] - Teleport to specified Instance by ID (coordinates are optional) | Guide (50) |
| #zoneshutdown | [shortname] - Shut down a zone server | GMLeadAdmin (150) |
| #zsave |  Saves zheader to the database | QuestTroupe (80) |

| Command | Subcommand | Usage |
| :--- | :--- | :--- |
| #find | aa | #find aa [Search Criteria] |
| #find | body_type | #find body_type [Search Criteria] |
| #find | bug_category | #find bug_category [Search Criteria] |
| #find | character | #find character [Search Criteria] |
| #find | class | #find class [Search Criteria] |
| #find | comparison_type | #find comparison_type [Search Criteria] |
| #find | currency | #find currency [Search Criteria] |
| #find | deity | #find deity [Search Criteria] |
| #find | emote | #find emote [Search Criteria] |
| #find | faction | #find faction [Search Criteria] |
| #find | item | #find item [Search Criteria] |
| #find | language | #find language [Search Criteria] |
| #find | object_type | #find object_type [Search Criteria] |
| #find | race | #find race [Search Criteria] |
| #find | recipe | #find recipe [Search Criteria] |
| #find | skill | #find skill [Search Criteria] |
| #find | special_ability | #find special_ability [Search Criteria] |
| #find | stance | #find stance [Search Criteria] |
| #find | spell | #find spell [Search Criteria] |
| #find | task | #find task [Search Criteria] |
| #find | zone | #find zone [Search Criteria] |
| #set | aa_exp | #set aa_exp [aa&#124;group&#124;raid] [Amount] |
| #set | aa_points | #set aa_points [aa&#124;group&#124;raid] [Amount] |
| #set | adventure_points | #set adventure_points [Theme ID] [Amount] |
| #set | alternate_currency | #set alternate_currency [Currency ID] [Amount] |
| #set | animation | #set animation [Animation ID] |
| #set | anon | #set anon [Character ID] [Anonymous Flag] or #set anon [Anonymous Flag] |
| #set | auto_login | #set auto_login [0&#124;1] |
| #set | bind_point | #set bind_point |
| #set | checksum | #set checksum |
| #set | class_permanent | #set class_permanent [Class ID] |
| #set | crystals | #set crystals [ebon&#124;radiant] [Amount] |
| #set | date | #set date [Year] [Month] [Day] [Hour] [Minute] (Hour and Minute are optional) |
| #set | endurance | #set endurance [Amount] |
| #set | endurance_full | #set endurance_full |
| #set | exp | #set exp [aa&#124;exp] [Amount] |
| #set | flymode | #set flymode [Flymode ID] |
| #set | frozen | #set frozen [on&#124;off] |
| #set | gender | #set gender [Gender ID] |
| #set | gender_permanent | #set gender_permanent [Gender ID] |
| #set | gm | #set gm [on&#124;off] |
| #set | gm_speed | #set gm_speed [on&#124;off] |
| #set | gm_status | #set gm_status [GM Status] [Account] |
| #set | god_mode | #set god_mode [on&#124;off] |
| #set | haste | #set haste [Percentage] |
| #set | hide_me | #set hide_me [on&#124;off] |
| #set | hero_model | #set hero_model [Hero Model] [Slot] (Slot is optional) |
| #set | hp | #set hp [Amount] |
| #set | hp_full | #set hp_full |
| #set | invulnerable | #set invulnerable |
| #set | language | #set language [Language ID] [Language Level] |
| #set | last_name | #set last_name [Last Name] |
| #set | level | #set level [Level] |
| #set | loginserver_info | #set loginserver_info [Email] [Password] |
| #set | mana | #set mana [Amount] |
| #set | mana_full | #set mana_full |
| #set | motd | #set motd |
| #set | name | #set name |
| #set | ooc_mute | #set ooc_mute |
| #set | password | #set password [Account Name] [Password] (account table password) |
| #set | pvp | #set pvp [on&#124;off] |
| #set | pvp_points | #set pvp_points [Amount] |
| #set | race | #set race [Race ID] |
| #set | race_permanent | #set race_permanent [Race ID] |
| #set | server_locked | #set server_locked [on&#124;off] |
| #set | skill | #set skill [Skill ID] [Skill Level] |
| #set | skill_all | #set skill_all [Skill Level] |
| #set | skill_all_max | #set skill_all_max |
| #set | start_zone | #set endurance [Amount] |
| #set | temporary_name | #set temporary_name [Name] |
| #set | texture | #set texture [Texture ID] |
| #set | time | #set time [Hour] [Minute] |
| #set | time_zone | #set time_zone [Hour] [Minute] |
| #set | title | #set title [Title] |
| #set | title_suffix | #set title_suffix [Title Suffix] |
| #set | temporary_name | #set temporary_name [Name] |
| #set | weather | #set weather [0&#124;1&#124;2&#124;3] |
| #set | zone | #set zone [option] |
| #show | aas | #show aas |
| #show | aa_points | #show aa_points |
| #show | aggro | #show aggro [Distance] [-v] (-v is verbose Faction Information) |
| #show | auto_login | #show auto_login |
| #show | buffs | #show buffs |
| #show | buried_corpse_count | #show buried_corpse_count |
| #show | client_version_summary | #show client_version_summary |
| #show | content_flags | #show content_flags |
| #show | currencies | #show currencies |
| #show | distance | #show distance |
| #show | emotes | #show emotes |
| #show | field_of_view | #show field_of_view |
| #show | flags | #show flags |
| #show | group_info | #show group_info |
| #show | hatelist | #show hatelist |
| #show | inventory | #show inventory |
| #show | ip_lookup | #show ip_lookup |
| #show | line_of_sight | #show line_of_sight |
| #show | network | #show network |
| #show | network_stats | #show network_stats |
| #show | npc_global_loot | #show npc_global_loot |
| #show | npc_stats | #show npc_stats |
| #show | npc_type | #show npc_type [NPC ID] |
| #show | peqzone_flags | #show peqzone_flags |
| #show | petition | #show petition |
| #show | petition_info | #show petition_info |
| #show | proximity | #show proximity |
| #show | quest_errors | #show quest_errors |
| #show | quest_globals | #show quest_globals |
| #show | recipe | #show recipe [Recipe ID] |
| #show | server_info | #show server_info |
| #show | skills | #show skills |
| #show | spawn_status | #show spawn_status [all&#124;disabled&#124;enabled&#124;Spawn ID] |
| #show | special_abilities | #show special_abilities |
| #show | spells | #show spells [disciplines&#124;spells] |
| #show | spells_list | #show spells_list |
| #show | stats | #show stats |
| #show | timers | #show timers |
| #show | traps | #show traps |
| #show | uptime | #show uptime [Zone Server ID] (Zone Server ID is optional) |
| #show | variable | #show variable [Variable Name] |
| #show | version | #show version |
| #show | waypoints | #show waypoints |
| #show | who | #show who [Search Criteria] (Search criteria is optional) |
| #show | xtargets | #show xtargets [Amount] (Amount is optional) |
| #show | zone_data | #show zone_data |
| #show | zone_global_loot | #show zone_global_loot |
| #show | zone_loot | #show zone_loot |
| #show | zone_points | #show zone_points |
| #show | zone_status | #show zone_status |
