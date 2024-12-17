---
description: >-
  This page describes the Rules and Rule Values for your EQEmu Server. These
  rules are found in the rule_values table.
---

# Server Rules

### Editing rules

To edit rules, you can use the [`#rules set [Name] [Value]`](https://docs.eqemu.io/server/operation/in-game-command-reference) command to set a rule.

You can then use [`#reload rules`](https://docs.eqemu.io/server/operation/loading-server-data/) to force all zones to reload the new rules.

!!! warning
     The `notes` field in the database is set by the server software.  If you adjust the values in the field, they will be overwritten and restored to default values when you update your server binaries.


!!! info
    If you would like to improve upon the descriptions or notes in the Server Rules table, please submit a pull request on the [ruletypes](https://github.com/EQEmu/Server/blob/master/common/ruletypes.h) header file.

## AA Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| AA:ExpPerPoint | 23976503 | Amount of experience per AA. Is the same as the amount of experience to go from level 51 to level 52 |
| AA:NormalizedAAEnabled | false | TSS+ change to AA that normalizes AA experience to a fixed # of white con kills independent of level |
| AA:NormalizedAANumberOfWhiteConPerAA | 25 | The number of white con kills per AA point |
| AA:ModernAAScalingEnabled | false | Are we linearly scaling AA experience based on total # of earned AA? |
| AA:ModernAAScalingStartPercent | 1000 | 1000% or 10x AA experience at the start of the scaling range |
| AA:ModernAAScalingAAMinimum | 0 | The minimum number of earned AA before AA experience scaling begins |
| AA:ModernAAScalingAALimit | 4000 | The number of earned AA when AA experience scaling ends |
| AA:SoundForAAEarned | false | Play sound when AA point earned |
| AA:UnusedAAPointCap | -1 | Cap for Unused AA Points.  Default: -1.  NOTE: DO NOT LOWER THIS WITHOUT KNOWING WHAT YOU ARE DOING. MAY RESULT IN PLAYERS LOSING AAs. |
| AA:MaxAAEXPPerKill | -1 | Maximum AA EXP per Kill (3425214 is about 7%) - Default: -1 will disable the check |

## Adventure Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Adventure:MinNumberForGroup | 2 | Minimum members for adventure group |
| Adventure:MaxNumberForGroup | 6 | Maximum members for adventure group |
| Adventure:MaxLevelRange | 9 | Maximum level range for adventure |
| Adventure:NumberKillsForBossSpawn | 45 | Number of adventure kills to make the boss spawn |
| Adventure:DistanceForRescueAccept | 10000.0 | Distance for adventure rescue accept |
| Adventure:DistanceForRescueComplete | 2500.0 | Distance for adventure rescue complete |
| Adventure:ItemIDToEnablePorts | 41000 | ItemID to enable adventure ports. 0 to disable, otherwise using a LDoN portal will require the user to have this item |
| Adventure:LDoNTrapDistanceUse | 625 | LDoN trap distance use |
| Adventure:LDoNBaseTrapDifficulty | 15.0 | LDoN base trap difficulty |
| Adventure:LDoNCriticalFailTrapThreshold | 10.0 | LDoN critical fail trap threshold |
| Adventure:LDoNLootCountModifier | 10 | LDoN Loot Count Modifier, lower is better (default is 10) |

## Aggro Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Aggro:SmartAggroList | true | Smart aggro list attempts to choose targets in a much smarter fashion, prefering players to pets, sitting and critically injured players to normal players, and players in melee range to players not |
| Aggro:SittingAggroMod | 35 | Aggro increase against sitting targets. 35=35% |
| Aggro:MeleeRangeAggroMod | 10 | Aggro increase against targets in melee range. 10=10% |
| Aggro:CurrentTargetAggroMod | 0 | Aggro increase against current target. 0% = prefer the current target to any other. Makes it harder for our NPC to switch targets |
| Aggro:CriticallyWoundedAggroMod | 100 | Aggro increase against critical wounded targets |
| Aggro:SpellAggroMod | 100 | Aggro increase for spells |
| Aggro:PetSpellAggroMod | 10 | Aggro increase for pet spells |
| Aggro:TunnelVisionAggroMod | 0.75 | People not currently the top hate generate this much hate on a Tunnel Vision mob |
| Aggro:MaxScalingProcAggro | 400 | Set to -1 for no limit. Maximum amount of aggro that HP scaling SPA effect in a proc will add |
| Aggro:IntAggroThreshold | 75 | Int lesser or equal the value will aggro regardless of level difference |
| Aggro:AllowTickPulling | false | tick pulling is an exploit in an NPC's call for help fixed sometime in 2006 on live |
| Aggro:MinAggroLevel | 18 | Minimum level for use with UseLevelAggro |
| Aggro:UseLevelAggro | true | MinAggroLevel rule value+ and Undead will aggro regardless of level difference. This will disabled Rule:IntAggroThreshold if set to true |
| Aggro:ClientAggroCheckMovingInterval | 1000 | Interval in which clients actually check for aggro while moving - in milliseconds - this should be lower than ClientAggroCheckIdleInterval |
| Aggro:ClientAggroCheckIdleInterval | 6000 | Interval in which clients actually check for aggro while idle - in milliseconds - this should be higher than ClientAggroCheckMovingInterval |
| Aggro:PetAttackRange | 40000.0 | Maximum squared range /pet attack works at default is 200 |
| Aggro:NPCAggroMaxDistanceEnabled | true | If enabled, NPC's will drop aggro beyond 600 units or what is defined at the zone level |
| Aggro:AggroPlayerPets | false | If enabled, NPCs will aggro player pets |
| Aggro:UndeadAlwaysAggro | true | should undead always aggro? |
| Aggro:BardAggroCap | 40 | per song bard aggro cap. |
| Aggro:InitialAggroBonus | 100 | Initial Aggro Bonus, Default: 100 |
| Aggro:InitialPetAggroBonus | 100 | Initial Pet Aggro Bonus, Default 100 |

## Analytics Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Analytics:CrashReporting | true | Automatic crash reporting analytics for EQEmu Server developers |

## Bazaar Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Bazaar:AuditTrail | false | Setting whether a path to the trader should be displayed in the bazaar |
| Bazaar:MaxSearchResults | 50 | Maximum number of search results in Bazaar |
| Bazaar:EnableWarpToTrader | true | Setting whether teleport to the selected trader should be active |
| Bazaar:MaxBarterSearchResults | 200 | The maximum results returned in the /barter search |
| Bazaar:ParcelDeliveryCostMod | 0.20 | Cost of parcel delivery for a bazaar purchase as a percentage of item cost. Default is 20% of item cost. RoF+ Only. |
| Bazaar:VoucherDeliveryCost | 200 | Number of vouchers for direct delivery for a bazaar purchase. Default is 200 vouchers. RoF+ Only. |
| Bazaar:EnableParcelDelivery | true | Enable bazaar purchases via parcel delivery.  Default is True. |
| Bazaar:MaxBuyerInventorySearchResults | 200 | Maximum number of search results when a Buyer searches the global item list. Default is 200. RoF+ Only. |

## Bots Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Bots:Enabled | false | Enable of disable bot functionality, default is false |
| Bots:BotExpansionSettings | 16383 | Sets the expansion settings for bot use. Defaults to all expansions enabled up to TSS |
| Bots:AllowCamelCaseNames | false | Allows the use of 'MyBot' type names |
| Bots:AllowBotEquipAnyRaceGear | false | Allows Bots to wear Equipment even if their race is not valid |
| Bots:CommandSpellRank | 1 | Filters bot command spells by rank. 1, 2 and 3 are valid filters - any other number allows all ranks |
| Bots:CreationLimit | 150 | Number of bots that each account can create |
| Bots:FinishBuffing | false | Allow for buffs to complete even if the bot caster is out of mana. Only affects buffing out of combat |
| Bots:GroupBuffing | false | Bots will cast single target buffs as group buffs, default is false for single. Does not make single target buffs work for MGB |
| Bots:HealRotationMaxMembers | 24 | Maximum number of heal rotation members |
| Bots:HealRotationMaxTargets | 12 | Maximum number of heal rotation targets |
| Bots:ManaRegen | 2.0 | Adjust mana regen. Acts as a final multiplier, stacks with Rule Character:ManaRegenMultiplier. |
| Bots:PreferNoManaCommandSpells | true | Give sorting priority to newer no-mana spells (i.e., 'Bind Affinity') |
| Bots:QuestableSpawnLimit | false | Optional quest method to manage bot spawn limits using the quest_globals name bot_spawn_limit, see: /bazaar/Aediles_Thrall.pl |
| Bots:SpawnLimit | 71 | Number of bots a character can have spawned at one time, You + 71 bots is a 12 group pseudo-raid |
| Bots:BotGroupXP | false | Determines whether client gets experience for bots outside their group |
| Bots:BotLevelsWithOwner | false | Auto-updates spawned bots as owner levels/de-levels (false is original behavior) |
| Bots:BotCharacterLevel | 0 | If level is greater that value player can spawn bots if BotCharacterLevelEnabled is true |
| Bots:CasterStopMeleeLevel | 13 | Level at which caster bots stop melee attacks |
| Bots:AllowOwnerOptionAutoDefend | true | When option is enabled, bots will defend their owner on enemy aggro |
| Bots:LeashDistance | 562500.0f | Distance a bot is allowed to travel from leash owner before being pulled back (squared value) |
| Bots:AllowApplyPoisonCommand | true | Allows the use of the bot command 'applypoison' |
| Bots:AllowApplyPotionCommand | true | Allows the use of the bot command 'applypotion' |
| Bots:RestrictApplyPotionToRogue | true | Restricts the bot command 'applypotion' to rogue-usable potions (i.e., poisons) |
| Bots:OldRaceRezEffects | false | Older clients had ID 757 for races with high starting STR, but it doesn't seem used anymore |
| Bots:ResurrectionSickness | true | Use Resurrection Sickness based on Resurrection spell cast, set to false to disable Resurrection Sickness. |
| Bots:OldResurrectionSicknessSpell | 757 | 757 is Default Old Resurrection Sickness Spell |
| Bots:ResurrectionSicknessSpell | 756 | 756 is Default Resurrection Sickness Spell |
| Bots:AllowPickpocketCommand | true | Allows the use of the bot command 'pickpocket' |
| Bots:BotHealOnLevel | false | Setting whether a bot should heal completely when leveling. Default FALSE. |
| Bots:AutosaveIntervalSeconds | 300 | Number of seconds after which a timer is triggered which stores the bot data. The value 0 means no periodic automatic saving. |
| Bots:CazicTouchBotsOwner | true | Default True. Cazic Touch/DT will hit bot owner rather than bot. |
| Bots:BotsClickItemsMinLvl | 1 | Minimum level for bots to be able to use ^clickitem. Default 1. |
| Bots:BotsCanClickItems | true | Enables the ability for bots to click items they have equipped. Default TRUE |
| Bots:CanClickMageEpicV1 | true | Whether or not bots are allowed to click Mage Epic 1.0. Default TRUE |
| Bots:BotsIgnoreLevelBasedHasteCaps | false | Ignores hard coded level based haste caps. |
| Bots:BotsHasteCap | 100 | Haste cap for non-v3(over haste) haste |
| Bots:BotsHastev3Cap | 25 | Haste cap for v3(over haste) haste |

## Bugs Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Bugs:ReportingSystemActive | true | Activates bug reporting |

## Channels Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Channels:RequiredStatusAdmin | 251 | Required status to administer chat channels |
| Channels:RequiredStatusListAll | 251 | Required status to list all chat channels |
| Channels:DeleteTimer | 1440 | Empty password protected channels will be deleted after this many minutes |

## Character Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Character:MaxLevel | 65 | Sets the highest level for players that can be reached through experience |
| Character:PerCharacterQglobalMaxLevel | false | Check for qglobal 'CharMaxLevel' character qglobal (Type 5, \"\, if player tries to level beyond that point, it will not go beyond that level |
| Character:PerCharacterBucketMaxLevel | false | Check for data bucket 'CharMaxLevel', if player tries to level beyond that point, it will not go beyond that level |
| Character:MaxExpLevel | 0 | Defines the maximum level that can be reached through experience |
| Character:DeathExpLossLevel | 10 | Any level equal to or greater than this will lose experience at death |
| Character:DeathExpLossMaxLevel | 255 | Every higher level will no longer lose experience at death |
| Character:DeathItemLossLevel | 10 | From this level on, items are left in the corpse when LeaveCorpses is activated |
| Character:DeathExpLossMultiplier | 3 | Adjust how much experience is lost. Default 3.5% (0=0.5%, 1=1.5%, 2=2.5%, 3=3.5%, 4=4.5%, 5=5.5%, 6=6.5%, 7=7.5%, 8=8.5%, 9=9.5%, 10=11%) |
| Character:DeathKeepLevel | false | Players can not drop below 0% experience from death. |
| Character:UseDeathExpLossMult | false | Setting to control whether DeathExpLossMultiplier or the code default is used: (Level x Level / 18.0) x 12000 |
| Character:UseOldRaceRezEffects | false | Older clients had ID 757 for races with high starting STR, but it doesn't seem used anymore |
| Character:CorpseDecayTime | 604800000 | Time after which the corpse decays (milliseconds) DEFAULT: 604800000 (7 Days) |
| Character:EmptyCorpseDecayTime | 10800000 | Time after which an empty corpse decays (milliseconds) DEFAULT: 10800000 (3 Hours) |
| Character:CorpseResTime | 10800000 | Time after which the corpse can no longer be resurrected (milliseconds) DEFAULT: 10800000 (3 Hours) |
| Character:DuelCorpseResTime | 600000 | Time before cant res corpse after a duel (milliseconds) DEFAULT: 600000 (10 Minutes) |
| Character:CorpseOwnerOnlineTime | 30000 | How often corpse will check if its owner is online DEFAULT: 30000 (30 Seconds) |
| Character:LeaveCorpses | true | Setting whether you leave a corpse behind |
| Character:LeaveNakedCorpses | false | Setting whether you leave a corpse without items |
| Character:MaxDraggedCorpses | 2 | Maximum number of corpses you can drag at once |
| Character:DragCorpseDistance | 400 | If a player is using /corpsedrag and moving, the corpse will not move until the player exceeds this distance |
| Character:FinalExpMultiplier | 1 | Added on top of everything else, easy for setting EXP events |
| Character:ExpMultiplier | 0.5 | If greater than 0, the experience gained is multiplied by this value.  |
| Character:AAExpMultiplier | 0.5 | If greater than 0, the AA experience gained is multiplied by this value.  |
| Character:GroupExpMultiplier | 0.5 | The experience in a group is multiplied by this value in addition to the group multiplier. The group multiplier is: 2 members=x 1.2, 3=x1.4, 4=x1.6, 5=x1.8, 6=x2.16 |
| Character:RaidExpMultiplier | 0.2 | The experience gained in raids is multiplied by (1-RaidExpMultiplier)  |
| Character:FinalRaidExpMultiplier | 1.0 | Multiplies all raid experience by this value. Default: 1.0 |
| Character:UseXPConScaling | true | When activated, the experience is modified depending on the difference between player level and NPC level. The values from the rules GreenModifier to RedModifier are used |
| Character:ShowExpValues | 0 | Show experience values. 0=normal, 1=show raw experience values, 2=show raw experience values and percent |
| Character:GreenModifier | 20 | The experience obtained for green con mobs is multiplied by value/100 |
| Character:LightBlueModifier | 40 | The experience obtained for light-blue con mobs is multiplied by value/100 |
| Character:BlueModifier | 90 | The experience obtained for blue con mobs is multiplied by value/100 |
| Character:WhiteModifier | 100 | The experience obtained for white con mobs is multiplied by value/100 |
| Character:YellowModifier | 125 | The experience obtained for yellow con mobs is multiplied by value/100 |
| Character:RedModifier | 150 | The experience obtained for red con mobs is multiplied by value/100 |
| Character:AutosaveIntervalS | 300 | Number of seconds after which a timer is triggered which stores the character data. The value 0 means no periodic automatic saving. |
| Character:HPRegenMultiplier | 100 | The hitpoint regeneration is multiplied by value/100 (up to the caps) |
| Character:ManaRegenMultiplier | 100 | The mana regeneration is multiplied by value/100 (up to the caps) |
| Character:EnduranceRegenMultiplier | 100 | The endurance regeneration is multiplied by value/100 (up to the caps) |
| Character:OldMinMana | false | This is used for servers that want to follow older skill cap formulas so they can still have some regen w/o mediate |
| Character:HealOnLevel | false | Setting whether a player should heal completely when leveling |
| Character:FeignKillsPet | false | Setting whether Feign Death kills the player pet |
| Character:ItemManaRegenCap | 15 | Limit on mana regeneration granted by items |
| Character:ItemHealthRegenCap | 30 | Limit on health regeneration granted by items |
| Character:ItemDamageShieldCap | 30 | Limit on damage shields granted by items |
| Character:ItemAccuracyCap | 150 | Limit on accuracy granted by items |
| Character:ItemAvoidanceCap | 100 | Limit on avoidance granted by items |
| Character:ItemCombatEffectsCap | 100 | Limit on combat effects granted by items |
| Character:ItemShieldingCap | 35 | Limit on shielding granted by items |
| Character:ItemSpellShieldingCap | 35 | Limit on spell shielding granted by items |
| Character:ItemDoTShieldingCap | 35 | Limit on DoT shielding granted by items |
| Character:ItemStunResistCap | 35 | Limit on resistance granted by items |
| Character:ItemStrikethroughCap | 35 | Limit on strikethrough granted by items |
| Character:ItemATKCap | 250 | Limit on ATK granted by items |
| Character:ItemHealAmtCap | 250 | Limit on heal amount granted by items |
| Character:ItemSpellDmgCap | 250 | Limit on spell damage granted by items |
| Character:ItemClairvoyanceCap | 250 | Limit on clairvoyance granted by items |
| Character:ItemDSMitigationCap | 50 | Limit on damageshield mitigation granted by items |
| Character:ItemEnduranceRegenCap | 15 | Limit on endurance regeneration granted by items |
| Character:ItemExtraDmgCap | 150 | Cap for bonuses to melee skills like Bash, Frenzy, etc. |
| Character:HasteCap | 100 | Haste cap for non-v3(over haste) haste |
| Character:Hastev3Cap | 25 | Haste cap for v3(over haste) haste |
| Character:HeroicStatsUseDataBucketsToScale | false | Allows scaling the benefits a character receives from Heroic Stats using Data Buckets. Stacks with other Heroic Stats Mulitplier Rules. |
| Character:HeroicIntelligenceIncreaseSpellDmgMultiplier | 0.00 | Allows Heroic Intelligence to increase a Players Worn Spell Damage Stat from Equipment, for example, setting this rule to 1.00 will always grant 1 Spell Damage per 1 Heroic Intelligence |
| Character:HeroicWisdomIncreaseHealAmtMultiplier | 0.00 | Allows Heroic Wisdom to increase a Players Worn Heal Amount Stat from Equipment, for example, setting this rule to 1.00 will always grant 1 Heal Amount per 1 Heroic Wisdom |
| Character:HeroicStrengthMultiplier | 1.00 | Multplier scales benefits from Heroic Strength. Grants 25 Base Endurance, 0.05 Endurance Regen, 1 Melee Damage each Hit, and 1 Shield AC per 10 Heroic Strength. |
| Character:HeroicStaminaMultiplier | 1.00 | Multplier scales benefits from Heroic Stamina. Grants 25 Base Endurance, 0.05 Endurance Regen, 100 Base HP, and 0.5 HP Regen per 10 Heroic Stamina. |
| Character:HeroicAgilityMultiplier | 1.00 | Multplier scales benefits from Heroic Agility. Grants 25 Base Endurance, 0.05 Endurance Regen, and 1 Avoidance AC per 10 Heroic Agility. (Rule does not change Dodge Chance) |
| Character:HeroicDexterityMultiplier | 1.00 | Multplier scales benefits from Heroic Dexterity. Grants 25 Base Endurance, 0.05 Endurance Regen, and 1 Archery/Throwing Damage each hit per 10 Heroic Dexterity. (Rule does not change Assassinate/Headshot/Block/Parry/Riposte Chances) |
| Character:HeroicWisdomMultiplier | 1.00 | Multplier scales benefits from Heroic Wisdom. Grants 250 Base Mana, 1 Mana Regen per 25 Heroic Wisdom. |
| Character:HeroicIntelligenceMultiplier | 1.00 | Multplier scales benefits from Heroic Intelligence. Grants 250 Base Mana, 1 Mana Regen per 25 Heroic Intelligence. |
| Character:SkillUpModifier | 100 | The probability for a skill-up is multiplied by value/100 |
| Character:SharedBankPlat | false | Shared bank platinum. Off by default to prevent duplication |
| Character:BindAnywhere | false | Allows players to bind their soul anywhere in the world |
| Character:RestRegenEnabled | true | Setting to activate out-of-combat regeneration |
| Character:RestRegenTimeToActivate | 30 | Time in seconds for rest state regen to kick in |
| Character:RestRegenRaidTimeToActivate | 300 | Time in seconds for rest state regen to kick in with a raid target |
| Character:KillsPerGroupLeadershipAA | 250 | Minimum number of dark blue mobs that must be killed to get a Group Leadership AA |
| Character:KillsPerRaidLeadershipAA | 250 | Minimum number of dark blue mobs that must be killed to get a Raid Leadership AAA |
| Character:MaxFearDurationForPlayerCharacter | 4 | Maximum number of tics a player can be feared. 1 tic equls 6 seconds |
| Character:MaxCharmDurationForPlayerCharacter | 15 | Maximum number of tics a player can be charmed. 1 tic equls 6 seconds |
| Character:BaseHPRegenBonusRaces | 4352 | A bitmask of race(s) that receive the regen bonus. Iksar (4096) & Troll (256) = 4352. See common/races.h for the bitmask values |
| Character:SoDClientUseSoDHPManaEnd | false | Setting this to true will allow SoD clients to use the SoD HP/Mana/End formulas and previous clients will use the old formulas |
| Character:UseRaceClassExpBonuses | true | Setting this to true will enable Class and Racial experience rate bonuses |
| Character:UseOldRaceExpPenalties | false | Setting this to true will enable racial experience penalties for Iksar, Troll, Ogre, and Barbarian, as well as the bonus for Halflings |
| Character:UseOldClassExpPenalties | false | Setting this to true will enable old class experience penalties for Paladin, SK, Ranger, Bard, Monk, Wizard, Enchanter, Magician, and Necromancer, as well as the bonus for Rogues and Warriors |
| Character:RespawnFromHover | false | Setting whether the respawn window should be used |
| Character:RespawnFromHoverTimer | 300 | Respawn Window countdown timer, in seconds |
| Character:UseNewStatsWindow | true | Setting whether the new Stats window, which displays all information, should be used |
| Character:ItemCastsUseFocus | false | If true, this allows item clickies to use focuses that have limited maximum levels on them |
| Character:MinStatusForNoDropExemptions | 80 | This allows status x and higher to trade no drop items |
| Character:SkillCapMaxLevel | 75 | Sets the Maximum Level used for Skill Caps (from skill_caps table). -1 makes it use MaxLevel rule value. It is set to 75 because PEQ only has skill caps up to that level, and grabbing the players' skill past 75 will return 0, breaking all skills past that level. This helps servers with obsurd level caps (75+ level cap) function without any modifications |
| Character:StatCap | 0 | If StatCap > 0 then this value is used. If it is 0, the value of the following code is used: If Level < 61: 255. If Level >= 61 and the client SoF or newer: 255 + 5 x (level -60).  If the client is older than SoF and the level < 71 then: 255 + x (level-60). In all other cases: 330. |
| Character:CheckCursorEmptyWhenLooting | true | If true, a player cannot loot a corpse (player or NPC) with an item on their cursor |
| Character:MaintainIntoxicationAcrossZones | true | If true, alcohol effects are maintained across zoning and logging out/in |
| Character:EnableDiscoveredItems | true | If enabled, it enables EVENT_DISCOVER_ITEM and also saves character names and timestamps for the first time an item is discovered |
| Character:EnableXTargetting | true | Enable Extended Targeting Window, for users with UF and later clients |
| Character:EnableAggroMeter | true | Enable Aggro Meter, for users with RoF and later clients |
| Character:KeepLevelOverMax | false | Don't de-level a character that has somehow gone over the level cap |
| Character:FoodLossPerUpdate | 32 | How much food/water you lose per stamina update |
| Character:EnableHungerPenalties | false | Being hungry/thirsty has negative effects -- it does appear normal live servers do not have penalties |
| Character:EnableFoodRequirement | true | If disabled, food is no longer required |
| Character:BaseInstrumentSoftCap | 36 | Softcap for instrument mods, 36 commonly referred to as 3.6 as well |
| Character:UseSpellFileSongCap | true | When they removed the AA that increased the cap they removed the above and just use the spell field |
| Character:BaseRunSpeedCap | 158 | Base Run Speed Cap, on live it's 158% which will give you a runspeed of 1.580 hard capped to 225 |
| Character:EnvironmentDamageMulipliter | 1 | Multiplier for environmental damage like fall damage. |
| Character:UnmemSpellsOnDeath | true | Setting whether at death all memorized Spells are forgotten |
| Character:TradeskillUpAlchemy | 2.0 | Alchemy skillup rate adjustment. Lower is faster |
| Character:TradeskillUpBaking | 2.0 | Baking skillup rate adjustment. Lower is faster |
| Character:TradeskillUpBlacksmithing | 2.0 | Blacksmithing skillup rate adjustment. Lower is faster |
| Character:TradeskillUpBrewing | 3.0 | Brewing skillup rate adjustment. Lower is faster |
| Character:TradeskillUpFletching | 2.0 | Fletching skillup rate adjustment. Lower is faster |
| Character:TradeskillUpJewelcrafting | 2.0 | Jewelcrafting skillup rate adjustment. Lower is faster |
| Character:TradeskillUpMakePoison | 2.0 | Make Poison skillup rate adjustment. Lower is faster |
| Character:TradeskillUpPottery | 4.0 | Pottery skillup rate adjustment. Lower is faster |
| Character:TradeskillUpResearch | 1.0 | Research skillup rate adjustment. Lower is faster |
| Character:TradeskillUpTinkering | 2.0 | Tinkering skillup rate adjustment. Lower is faster |
| Character:TradeskillUpTailoring | 2.0 | Tailoring skillup rate adjustment. Lower is faster |
| Character:MarqueeHPUpdates | false | Will show health percentage in center of screen if health lesser than 100% |
| Character:IksarCommonTongue | 95 | Starting value for Common Tongue for Iksars |
| Character:OgreCommonTongue | 95 | Starting value for Common Tongue for Ogres |
| Character:TrollCommonTongue | 95 | Starting value for Common Tongue for Trolls |
| Character:ActiveInvSnapshots | false | Takes a periodic snapshot of inventory contents from online players |
| Character:InvSnapshotMinIntervalM | 180 | Minimum time between inventory snapshots (minutes) |
| Character:InvSnapshotMinRetryM | 30 | Time to re-attempt an inventory snapshot after a failure  (minutes) |
| Character:InvSnapshotHistoryD | 30 | Time to keep snapshot entries (days) |
| Character:RestrictSpellScribing | false | Setting whether to restrict spell scribing to allowable races/classes of spell scroll |
| Character:UseStackablePickPocketing | true | Allows stackable pickpocketed items to stack instead of only being allowed in empty inventory slots |
| Character:AllowMQTarget | false | Disables putting players in the 'hackers' list for targeting beyond the clip plane or attempting to target something untargetable |
| Character:UseOldBindWound | false | Uses the original bind wound behavior |
| Character:GrantHoTTOnCreate | false | Grant Health of Target's Target leadership AA on character creation |
| Character:UseOldConSystem | false | Setting whether the pre SoF era consider system should be used |
| Character:OPClientUpdateVisualDebug | false | Shows a pulse and forward directional particle each time the client sends its position to server |
| Character:AllowCrossClassTrainers | false | If the value is true, a player can also train with other class Guildmasters. |
| Character:PetsUseReagents | true | Conjuring pets consumes reagents |
| Character:DismountWater | true | Dismount horses when entering water |
| Character:NoSkillsOnHorse | false | Enabling this will prevent Bind Wound and Foraging while on a mount |
| Character:UseNoJunkFishing | false | Disregards junk items when fishing |
| Character:SoftDeletes | true | When characters are deleted in character select, they are only soft deleted |
| Character:DefaultGuild | 0 | If not 0, new characters placed into the guild # indicated |
| Character:DefaultGuildRank | 8 | Default guild rank when Character:DefaultGuild is a non-0 value, default is 8 (Recruit) |
| Character:ProcessFearedProximity | false | Processes proximity checks when feared |
| Character:EnableCharacterEXPMods | false | Enables character zone-based experience modifiers. |
| Character:PVPEnableGuardFactionAssist | true | Enables faction based assisting against the aggresor in pvp. |
| Character:SkillUpFromItems | true | Allow Skill ups from clickable items |
| Character:EnableTestBuff | false | Allow the use of /testbuff |
| Character:UseResurrectionSickness | true | Use Resurrection Sickness based on Resurrection spell cast, set to false to disable Resurrection Sickness. |
| Character:OldResurrectionSicknessSpellID | 757 | 757 is Default Old Resurrection Sickness Spell ID |
| Character:ResurrectionSicknessSpellID | 756 | 756 is Default Resurrection Sickness Spell ID |
| Character:EnableBardMelody | true | Enable Bard /melody by default, to disable change to false for a classic experience. |
| Character:EnableRangerAutoFire | true | Enable Ranger /autofire by default, to disable change to false for a classic experience. |
| Character:EnableTGB | true | Enable /tgb (Target Group Buff) by default, to disable change to false for a classic experience. |
| Character:SkillUpMaximumChancePercentage | 25 | Maximum chance to improve a combat skill, before skill-specific modifiers.  This should be greater than SkillUpMinimumChancePercentage. |
| Character:SkillUpMinimumChancePercentage | 2 | Minimum chance to improve a combat skill, after skill-specific modifiers.  This should be lesser than SkillUpMaximumChancePercentage. |
| Character:WarriorTrackingDistanceMultiplier | 0 | If you want warriors to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:ClericTrackingDistanceMultiplier | 0 | If you want clerics to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:PaladinTrackingDistanceMultiplier | 0 | If you want paladins to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:RangerTrackingDistanceMultiplier | 12 | If you want rangers to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:ShadowKnightTrackingDistanceMultiplier | 0 | If you want shadow knights to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:DruidTrackingDistanceMultiplier | 10 | If you want druids to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:MonkTrackingDistanceMultiplier | 0 | If you want monks to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:BardTrackingDistanceMultiplier | 7 | If you want bards to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:RogueTrackingDistanceMultiplier | 0 | If you want rogues to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:ShamanTrackingDistanceMultiplier | 0 | If you want shaman to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:NecromancerTrackingDistanceMultiplier | 0 | If you want necromancers to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:WizardTrackingDistanceMultiplier | 0 | If you want wizards to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:MagicianTrackingDistanceMultiplier | 0 | If you want magicians to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:EnchanterTrackingDistanceMultiplier | 0 | If you want enchanters to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:BeastlordTrackingDistanceMultiplier | 0 | If you want beastlords to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:BerserkerTrackingDistanceMultiplier | 0 | If you want berserkers to be able to track, increase this above 0.  0 disables tracking packets. |
| Character:OnInviteReceiveAlreadyinGroupMessage | true | If you want clients to receive a message when trying to invite a player into a group that is currently in another group. |
| Character:PetZoneWithOwner | true | Should Pets Zone with Owner |
| Character:FullManaOnDeath | true | On death set mana to full |
| Character:FullEndurOnDeath | true | On death set endurance to full |
| Character:ExperiencePercentCapPerKill | -1 | Caps the percentage of experience that can be gained per kill. -1 disables the cap; 0 blocks all (non-aa) xp. |
| Character:EnableGroupEXPModifier | true | Enable or disable the group experience modifier in group, default is true |
| Character:EnableGroupMemberEXPModifier | true | Enable or disable the group member experience modifier based on number of players in group, default is true |
| Character:GroupMemberEXPModifier | 0.2 | Sets the group experience modifier per members between 2 and 5, default is 0.2 |
| Character:FullGroupEXPModifier | 2.16 | Sets the group experience modifier for a full group, default is 2.16 |
| Character:IgnoreLevelBasedHasteCaps | false | Ignores hard coded level based haste caps. |
| Character:EnableRaidEXPModifier | true | Enable or disable the raid experience modifier, default is true |
| Character:EnableRaidMemberEXPModifier | true | Enable or disable the raid experience modifier based on members in raid, default is true |
| Character:LeaveCursorMoneyOnCorpse | false | Enable or disable leaving cursor money on player corpses |
| Character:ItemExtraSkillDamageCalcAsPercent | false | If enabled, apply Item Extra Skill Damage as Percentage-based modifiers |
| Character:UseForageCommonFood | true | If enabled, use the common foods specified in the code. |
| Character:ClearXTargetDelay | 10 | Seconds between uses of the #clearxtargets command (Set to 0 to disable) |
| Character:PreventMountsFromZoning | false | Enable to prevent mounts from zoning - Prior to December 15, 2004 this is enabled. |
| Character:GroupInvitesRequireTarget | false | Enable to require players to have invitee on target (Disables /invite name) - Classic Style |
| Character:PlayerTradingLoreFeedback | true | If enabled, during a player to player trade, if lore items exist, it will output which items. |
| Character:MendAlwaysSucceedValue | 199 | Value at which mend will always succeed its skill check. Default: 199 |
| Character:SneakAlwaysSucceedOver100 | false | When sneak skill is over 100, always succeed sneak/hide. Default: false |
| Character:BandolierSwapDelay | 0 | Bandolier swap delay in milliseconds, default is 0 |

## Chat Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Chat:ServerWideOOC | true | Enable server wide ooc-chat |
| Chat:ServerWideAuction | true | Enable server wide auction-chat |
| Chat:EnableVoiceMacros | true | Enable voice macros |
| Chat:EnableMailKeyIPVerification | true | Setting whether the authenticity of the client should be verified via its IP address when accessing the InGame mailbox |
| Chat:EnableAntiSpam | true | Enable anti-spam system for chat |
| Chat:SuppressCommandErrors | false | Do not suppress command errors by default |
| Chat:ChannelsIgnoreNameFilter | false | Ignore name filtering when creating new chat channels |
| Chat:MaxPermanentPlayerChannels | 0 | Maximum number of permanent chat channels a player can make. Default 0. |
| Chat:MinStatusToBypassAntiSpam | 100 | Minimum status to bypass the anti-spam system |
| Chat:MinimumMessagesPerInterval | 4 | Minimum number of chat messages allowed per interval. The karma value is added to this value |
| Chat:MaximumMessagesPerInterval | 12 | Maximum value of chat messages allowed per interval |
| Chat:MaxMessagesBeforeKick | 20 | If an attempt is made to send more than the maximum allowed number of chat messages per interval, the client will be disconnected after this absolute number of messages |
| Chat:IntervalDurationMS | 60000 | Interval length in milliseconds |
| Chat:KarmaUpdateIntervalMS | 1200000 | Karma update interval in milliseconds |
| Chat:KarmaGlobalChatLimit | 72 | Amount of karma you need to be able to talk in ooc/auction/chat below the level limit |
| Chat:GlobalChatLevelLimit | 8 | Level limit you need to of reached to talk in ooc/auction/chat if your karma is too low |
| Chat:AutoInjectSaylinksToSay | true | Automatically injects saylinks into dialogue that has [brackets in them] |
| Chat:AutoInjectSaylinksToClientMessage | true | Automatically injects saylinks into dialogue that has [brackets in them] |
| Chat:QuestDialogueUsesDialogueWindow | false | Pipes all quest dialogue to dialogue window |
| Chat:DialogueWindowAnimatesNPCsIfNoneSet | true | If there is no animation specified in the dialogue window markdown then it will choose a random greet animation such as wave or salute |

## Cheat Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Cheat:MQWarpDetectionDistanceFactor | 9.0 | clients move at 4.4 about if in a straight line but with movement and to acct for lag we raise it a bit |
| Cheat:MQWarpExemptStatus | -1 | Required status level to exempt the MQWarpDetector. Set to -1 to disable this feature. |
| Cheat:MQZoneExemptStatus | -1 | Required status level to exempt the MQZoneDetector. Set to -1 to disable this feature. |
| Cheat:MQGateExemptStatus | -1 | Required status level to exempt the MQGateDetector. Set to -1 to disable this feature. |
| Cheat:MQGhostExemptStatus | -1 | Required status level to exempt the MQGhostDetector. Set to -1 to disable this feature. |
| Cheat:MQFastMemExemptStatus | -1 | Required status level to exempt the MQFastMemDetector. Set to -1 to disable this feature. |
| Cheat:EnableMQWarpDetector | true | Enable the MQWarp Detector. Set to False to disable this feature. |
| Cheat:EnableMQZoneDetector | true | Enable the MQZone Detector. Set to False to disable this feature. |
| Cheat:EnableMQGateDetector | true | Enable the MQGate Detector. Set to False to disable this feature. |
| Cheat:EnableMQGhostDetector | true | Enable the MQGhost Detector. Set to False to disable this feature. |
| Cheat:EnableMQFastMemDetector | true | Enable the MQFastMem Detector. Set to False to disable this feature. |
| Cheat:MarkMQWarpLT | false | Mark clients makeing smaller warps |

## Client Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Client:UseLiveFactionMessage | false | Allows players to see detailed faction adjustments as on the live servers |
| Client:UseLiveBlockedMessage | false | Setting whether detailed spell block messages should be used as on the live servers |

## Combat Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Combat:AERampageMaxDistance | 70 | Max AERampage range (% of max combat distance) |
| Combat:PetBaseCritChance | 0 | Pet base crit chance |
| Combat:NPCBashKickLevel | 6 | The level that NPCcan KICK/BASH |
| Combat:NPCKickStunLevel | 1 | The level that NPC has a chance to stun with kick. |
| Combat:PCKickStunLevel | 56 | The level that PC has a chance to stun with kick. |
| Combat:MeleeCritDifficulty | 8900 | Value against which is rolled to check if a melee crit is triggered. Lower is easier |
| Combat:ArcheryCritDifficulty | 3400 | Value against which is rolled to check if an archery crit is triggered. Lower is easier |
| Combat:ThrowingCritDifficulty | 1100 | Value against which is rolled to check if a throwing crit is triggered. Lower is easier |
| Combat:NPCCanCrit | false | Setting whether an NPC can land critical hits |
| Combat:UseIntervalAC | true | Switch whether bonuses, armour class, multipliers, classes and caps should be considered in the calculation of damage values |
| Combat:PetAttackMagicLevel | 10 | Level at which pets can cause magic damage, no longer used |
| Combat:NPCAttackMagicLevel | 10 | Level at which NPC and pets can cause magic damage |
| Combat:LevelDifferenceRollCheck | -1 | Level Difference to enable LeverDifferenceRollBonus for MeleeMitigation - Default: -1 is disabled, 20 is common |
| Combat:LevelDifferenceRollBonus | 0.5 | Roll Bonus/Detrement if using LevelDifferenceRollCheck |
| Combat:EnableFearPathing | true | Setting whether to use pathing during fear |
| Combat:FleeGray | true | If true FleeGrayHPRatio will be used |
| Combat:FleeGrayHPRatio | 50 | HP percentage when a Gray NPC begins to flee |
| Combat:FleeGrayMaxLevel | 18 | NPC above this level won't do gray/green con flee |
| Combat:FleeHPRatio | 25 | HP percentage when a NPC begins to flee |
| Combat:FleeIfNotAlone | false | If false, mobs won't flee if other mobs are in combat with it |
| Combat:AdjustProcPerMinute | true | Adapt the average proc rate to the speed of the weapon |
| Combat:AvgProcsPerMinute | 2.0 | Average proc rate per minute |
| Combat:ProcPerMinDexContrib | 0.075 | Increases the probability of a proc increased by DEX by the value indicated |
| Combat:BaseProcChance | 0.035 | Base chance for procs |
| Combat:ProcDexDivideBy | 11000 | Divisor for the probability of a proc increased by dexterity |
| Combat:MinRangedAttackDist | 25 | Minimum Distance to use Ranged Attacks |
| Combat:ArcheryBonusRequiresStationary | true | does the 2x archery bonus chance require a stationary npc |
| Combat:ArcheryBonusLevelRequirement | 51 | Level requirement when the 2x archery bonus will be enabled. The default is 51. |
| Combat:ArcheryNPCMultiplier | 1.0 | Value is multiplied by the regular dmg to get the archery dmg |
| Combat:AssistNoTargetSelf | true | When assisting a target that does not have a target: true = target self, false = leave target as was before assist (false = live like) |
| Combat:MaxRampageTargets | 3 | Maximum number of people hit with rampage |
| Combat:DefaultRampageTargets | 1 | Default number of people to hit with rampage |
| Combat:RampageHitsTarget | false | Rampage will hit the target if it still has targets left |
| Combat:MaxFlurryHits | 2 | Maximum number of extra hits from flurry |
| Combat:MinHastedDelay | 400 | Minimum hasted combat delay |
| Combat:AvgDefProcsPerMinute | 2.0 | Average defense procs per minute |
| Combat:DefProcPerMinAgiContrib | 0.075 | How much agility contributes to defensive proc rate |
| Combat:NPCFlurryChance | 20 | Chance for NPC to flurry |
| Combat:TauntOverLevel | 1 | Allows you to taunt NPC's over warriors level |
| Combat:TauntOverAggro | 0 | + amount over hate_top it will add before any bonus hate. |
| Combat:TauntChanceBonus | 0 | Bonus to taunt chance |
| Combat:ClassicTauntSystem | false | Enable to use the pre 2006 taunt system. |
| Combat:TauntSkillFalloff | 0.33 | For every taunt skill point that's not maxed you lose this percentage chance to taunt |
| Combat:EXPFromDmgShield | false | Determine if damage from a damage shield counts for experience gain |
| Combat:QuiverHasteCap | 1000 | Quiver haste cap 1000 on live for a while, currently 700 on live |
| Combat:BerserkerFrenzyStart | 35 | Percentage Health Points below which Warrior and Berserker start frenzy |
| Combat:BerserkerFrenzyEnd | 45 | Percentage Health Points above which Warrior and Berserker end frenzy |
| Combat:OneProcPerWeapon | true | If enabled, One proc per weapon per round |
| Combat:ProjectileDmgOnImpact | true | If enabled, projectiles (i.e. arrows) will hit on impact, instead of instantly |
| Combat:MeleePush | true | Enable melee push |
| Combat:MeleePushChance | 50 | NPC chance the target will be pushed. Made up, 100 actually isn't that bad |
| Combat:MeleePushForceClientPercent | 0.00 | Percent to add or remove from push for players |
| Combat:MeleePushForcePetPercent | 0.00 | Percent to add or remove from push for pets |
| Combat:NPCtoNPCPush | false | Disabled prevents NPC to NPC pushing per the 2013+ patch. |
| Combat:UseLiveCombatRounds | true | Turn this false if you don't want to worry about fixing up combat rounds for NPCs |
| Combat:NPCAssistCap | 5 | Maximum number of NPC that will assist another NPC at once |
| Combat:NPCAssistCapTimer | 6000 | Time a NPC will take to clear assist aggro cap space (milliseconds) |
| Combat:UseRevampHandToHand | false | Use h2h revamped dmg/delays I believe this was implemented during SoF |
| Combat:ClassicMasterWu | false | Classic Master Wu uses a random special, modern doesn't |
| Combat:HitBoxMod | 1.00 | Added to test hit boxes. |
| Combat:LevelToStopDamageCaps | 0 | Level to stop damage caps. 1 will effectively disable them, 20 should give basically same results as old incorrect system |
| Combat:LevelToStopACTwinkControl | 50 | Level to stop armorclass twink control. 1 will effectively disable it, 50 should give basically same results as current system |
| Combat:ClassicNPCBackstab | false | True disables NPC facestab - NPC get normal attack if not behind |
| Combat:UseNPCDamageClassLevelMods | true | Uses GetClassLevelDamageMod calc in npc_scale_manager |
| Combat:UseExtendedPoisonProcs | false | Allow old school poisons to last until characrer zones, at a lower proc rate |
| Combat:EnableSneakPull | false | Enable implementation of Sneak Pull |
| Combat:SneakPullAssistRange | 400 | Modified range of assist for sneak pull |
| Combat:Classic2HBAnimation | false | 2HB will use the 2 hand piercing animation instead of the overhead slashing animation |
| Combat:ArcheryConsumesAmmo | true | Set to false to disable Archery Ammo Consumption |
| Combat:ThrowingConsumesAmmo | true | Set to false to disable Throwing Ammo Consumption |
| Combat:UseLiveRiposteMechanics | false | Set to true to disable SPA 173 SE_RiposteChance from making those with the effect on them immune to enrage, can longer riposte from a riposte. |
| Combat:FrontalStunImmunityClasses | 0 | Bitmask for Classes than have frontal stun immunity, No Races (0) by default. |
| Combat:NPCsUseFrontalStunImmunityClasses | false | Enable or disable NPCs using frontal stun immunity Classes from Combat:FrontalStunImmunityClasses, false by default. |
| Combat:FrontalStunImmunityRaces | 512 | Bitmask for Races than have frontal stun immunity, Ogre (512) only by default. |
| Combat:NPCsUseFrontalStunImmunityRaces | true | Enable or disable NPCs using frontal stun immunity Races from Combat:FrontalStunImmunityRaces, true by default. |
| Combat:AssassinateOnlyHumanoids | true | Enable or disable Assassinate only being allowed on Humanoids, true by default. |
| Combat:AssassinateLevelRequirement | 60 | Level requirement to enable assassinate attempts on backstabs. The default is 60. |
| Combat:HeadshotOnlyHumanoids | true | Enable or disable Headshot only being allowed on Humanoids, true by default. |
| Combat:EnableWarriorShielding | true | Enable or disable Warrior Shielding Ability (/shield), true by default. |
| Combat:BackstabIgnoresElemental | false | Enable or disable Elemental weapon damage affecting backstab damage, false by default. |
| Combat:BackstabIgnoresBane | false | Enable or disable Bane weapon damage affecting backstab damage, false by default. |
| Combat:DoubleBackstabLevelRequirement | 55 | Level requirement to enable double backstab attempts. The default is 55. |
| Combat:SummonMeleeRange | true | Enable or disable summoning of a player when already in melee range of the summoner. |
| Combat:ArcheryHitPenalty | 0 | Archery has a hit penalty to try to help balance it with the plethora of long term +hit modifiers for it - Default: 0 |
| Combat:ArcheryBaseDamageBonus | 1 | Percentage modifier to base archery Damage 0.5=50% base damage, 1=100%,2=200% - Default: 1 |
| Combat:WaterMatchRequiredForAutoFireLoS | true | Enable/Disable the requirement of both the attacker/victim being both in or out of water for AutoFire LoS to pass. |
| Combat:ExtraAllowedKickClassesBitmask | 0 | Bitmask for allowing extra classes beyond Warrior, Ranger, Beastlord, and Berserker to kick, No Extra Classes (0) by default |
| Combat:MaxProcs | 4 | Adjustable maximum number of procs per round, the hard cap is MAX_PROCS (11). Requires mob repop or client zone when changed |
| Combat:FinishingBlowOnlyWhenFleeing | false | Enable to only allow Finishing Blow when fleeing (Original Style Finishing Blow) |
| Combat:ClassicTripleAttack | false | enable to use non-skill based classic triple attack. Originally it was Warrior Only but was expanded, can use the TripleAttackChance to tune the classes out. |
| Combat:ClassicTripleAttackChanceWarrior | 100 | Innate Chance for Warrior to Triple Attack after a Double Attack (125 = 12.5%). DEFAULT: 100 |
| Combat:ClassicTripleAttackChanceMonk | 100 | Innate Chance for Monk to Triple Attack after a Double Attack (200 = 20%). DEFAULT: 100 |
| Combat:ClassicTripleAttackChanceBerserker | 100 | Innate Chance for Berserker to Triple Attack after a Double Attack (200 = 20%). DEFAULT: 100 |
| Combat:ClassicTripleAttackChanceRanger | 100 | Innate Chance for Ranger to Triple Attack after a Double Attack (200 = 20%). DEFAULT: 100 |
| Combat:StunChance | 12 | Percent chance that client will be stunned when mob is behind player. DEFAULT: 12 |
| Combat:StunDuration | 2000 | Duration of stuns in ms. DEFAULT: 2000 |
| Combat:ClientStunMessage | false | Client stunning NPC produces message. DEFAULT false |
| Combat:BashTwoHanderUseShoulderAC | false | Enable to use shoulder AC for bash calculations when two hander is equipped. Unproven if accurate DEFAULT: false |
| Combat:BashACBonusDivisor | 25.0 | this divides the AC value contribution to bash damage, lower to increase damage |
| Combat:UseMobStaticOffenseSkill | false | Toggle to enabled the use of a static offense skill for Mobs.  DEFAULT: false |
| Combat:UseEnhancedMobStaticWeaponSkill | false | Toggle to enabled the use of an enhanced (slightly higher hit rate) static weapon skill for Mobs.  DEFAULT: false |
| Combat:PCAttackPowerScaling | 100 | Applies scaling to PC Attack Power (75 = 75%).  DEFAULT: 100 to not adjust existing Servers |
| Combat:PCAccuracyAvoidanceMod2Scale | 100 | Scale Factor for PC Accuracy and Avoidance (Mod2, found on items).  Found a value of 100 to make both too strong (75 = x0.75).  DEFAULT: 100 to not adjust existing Servers |
| Combat:AllowRaidTargetBlind | false | Toggle to allow raid targets to be blinded, default is false (Live-like) |
| Combat:RogueBackstabHasteCorrection | false | Toggle to enable correction for Haste impacting Backstab DPS too much.  DEFAULT: false |
| Combat:LegacyComputeDefense | false | Trim AGI Scaling of defense mostly for lower levels to help compensate for the newer agi based defense system. Default: False |
| Combat:SlayDamageMultiplier | 1.0 | Slay Damage Adjustment - Multiply final slay damage by this value. Default: 1.0 |
| Combat:SlayRateMultiplier | 1.0 | Slay Rate Adjustments - Multiply final slay rate check by this value. Default: 1.0 |
| Combat:MaximumLevelStunsCripplingBlow | 55 | Maximum level that Crippling Blows will stun a npc. Default: 55 |
| Combat:ArcheryBaseDamage | 0 | Archery base damage, default is 0 |
| Combat:BackstabBaseDamage | 0 | Backstab base damage, default is 0 |
| Combat:BashBaseDamage | 2 | Bash base damage, default is 2 |
| Combat:DragonPunchBaseDamage | 12 | Dragon Punch base damage, default is 12 |
| Combat:EagleStrikeBaseDamage | 7 | Eagle Strike base damage, default is 7 |
| Combat:FlyingKickBaseDamage | 25 | Flying Kick base damage, default is 25 |
| Combat:FrenzyBaseDamage | 10 | Frenzy base damage, default is 10 |
| Combat:KickBaseDamage | 3 | Kick base damage, default is 3 |
| Combat:RoundKickBaseDamage | 5 | Round Kick base damage, default is 5 |
| Combat:ThrowingBaseDamage | 0 | Throwing base damage, default is 0 |
| Combat:TigerClawBaseDamage | 4 | Tiger Claw base damage, default is 4 |

## Command Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Command:DyeCommandRequiresDyes | false | Enable this to require a Prismatic Dye (32557) each time someone uses #dye. |
| Command:HideMeCommandDisablesTells | true | Disable this to allow tells to be received when using #hideme. |

## Console Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Console:SessionTimeOut | 600000 | Amount of time in ms for the console session to time out |

## Doors Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Doors:RequireKeyOnCursor | false | Enable this to require pre-keyring keys to be on player cursor to open doors. |

## DynamicZone Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| DynamicZone:ClientRemovalDelayMS | 60000 | Delay (milliseconds) until a client is teleported out of dynamic zone after being removed as member |
| DynamicZone:EmptyShutdownEnabled | true | Enable early instance shutdown for dynamic zones that have no members |
| DynamicZone:EmptyShutdownDelaySeconds | 1500 | Seconds to set dynamic zone instance expiration if early shutdown enabled |
| DynamicZone:EnableInDynamicZoneStatus | false | Enables the 'In Dynamic Zone' member status in dynamic zone window. If false (live-like) players inside the dynamic zone will show as 'Online' |
| DynamicZone:WorldProcessRate | 6000 | Timer interval (milliseconds) that systems check their dynamic zone states |

## EventLog Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| EventLog:RecordSellToMerchant | false | Record sales from a player to an NPC merchant in eventlog table |
| EventLog:RecordBuyFromMerchant | false | Record purchases by a player from an NPC merchant in eventlog table |

## Expansion Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Expansion:CurrentExpansion | -1 | The current expansion enabled for the server [-1 = ALL, 0 = Classic, 1 = Kunark etc.] |
| Expansion:UseCurrentExpansionAAOnly | false | When true will only load AA ranks that match CurrentExpansion rule |
| Expansion:AutoGrantAAExpansion | -1 | Expansion to auto grant AAs up to, [-1 = Disabled, 0 = Classic, 1 = Kunark etc.] |

## Expedition Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Expedition:MinStatusToBypassPlayerCountRequirements | 80 | Minimum GM status to bypass minimum player requirements for Expedition creation |
| Expedition:AlwaysNotifyNewLeaderOnChange | false | Always notify clients when made expedition leader. If false (live-like) new leaders are only notified when made leader via /dzmakeleader |
| Expedition:LockoutDurationMultiplier | 1.0 | Multiplies lockout duration by this value when new lockouts are added |
| Expedition:ChooseLeaderCooldownTime | 2000 | Cooldown time (milliseconds) between choosing a new leader for automatic leader changes |

## Faction Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Faction:AllyFactionMinimum | 1100 | Minimum faction for ally |
| Faction:WarmlyFactionMinimum | 750 | Minimum faction for warmly |
| Faction:KindlyFactionMinimum | 500 | Minimum faction for kindly |
| Faction:AmiablyFactionMinimum | 100 | Minimum faction for amiably |
| Faction:IndifferentlyFactionMinimum | 0 | Minimum faction for indifferently |
| Faction:ApprehensivelyFactionMinimum | -100 | Minimum faction for apprehensively |
| Faction:DubiouslyFactionMinimum | -500 | Minimum faction for dubiously |
| Faction:ThreateninglyFactionMinimum | -750 | Minimum faction for threateningly |

## GM Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| GM:MinStatusToSummonItem | 250 | Minimum required status to summon items |
| GM:MinStatusToZoneAnywhere | 250 | Minimum required status to zone anywhere |
| GM:MinStatusToLevelTarget | 100 | Minimum required status to set the level of a player |
| GM:MinStatusToBypassLockedServer | 100 | Players >= this status can log in to the server even when it is locked |
| GM:MinStatusToBypassCheckSumVerification | 100 | Players >= this status can bypass the eqgame.exe and spells_us.txt checksum verification |

## Guild Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Guild:PlayerCreationAllowed | false | Allow players to create a guild using the window in Underfoot+ |
| Guild:PlayerCreationLimit | 1 | Only allow use of the UF+ window if the account has < than this number of guild leaders on it |
| Guild:PlayerCreationRequiredStatus | 0 | Required status to create a guild |
| Guild:PlayerCreationRequiredLevel | 0 | Required level of the player attempting to create the guild |
| Guild:PlayerCreationRequiredTime | 0 | Time needed online on the account to create a guild (in minutes) |
| Guild:TributeTime | 600000 | Time in ms for guild tributes.  Default is 10 mins. |
| Guild:TributeTimeRefreshInterval | 180000 | Time in ms to send all guild members a Tribute Time refresh. Default is 3 mins. |
| Guild:TributePlatConversionRate | 10 | The conversion rate of platinum donations.  Default is 10 guild favor to 1 platinum. |
| Guild:UseCharacterMaxLevelForGuildTributes | true | Guild Tributes will adhere to Character:MaxLevel.  Default is true. |

## HotReload Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| HotReload:QuestsRepopWithReload | true | When a hot reload is triggered, the zone will repop |
| HotReload:QuestsRepopWhenPlayersNotInCombat | true | When a hot reload is triggered, the zone will repop when no clients are in combat |
| HotReload:QuestsResetTimersWithReload | true | When a hot reload is triggered, quest timers will be reset |
| HotReload:QuestsAutoReloadGlobalScripts | false | When a quest, plugin, or global script changes, auto reload. |

## Instances Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Instances:ReservedInstances | 100 | Number of instance IDs which are reserved for globals. This value should not be changed while a server is running |
| Instances:RecycleInstanceIds | true | Setting whether free instance IDs should be recycled to prevent them from gradually running out at 32k |
| Instances:GuildHallExpirationDays | 90 | Amount of days before a Guild Hall instance expires |

## Inventory Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Inventory:EnforceAugmentRestriction | true | Forces augment slot restrictions |
| Inventory:EnforceAugmentUsability | true | Forces augmented item usability |
| Inventory:EnforceAugmentWear | true | Forces augment wear slot validation |
| Inventory:DeleteTransformationMold | true | False if you want mold to last forever |
| Inventory:AllowAnyWeaponTransformation | false | Weapons can use any weapon transformation |
| Inventory:TransformSummonedBags | false | Transforms summoned bags into disenchanted ones instead of deleting |
| Inventory:AllowMultipleOfSameAugment | false | Allows multiple of the same augment to be placed in an item via #augmentitem or MQ2, set to true to allow |
| Inventory:AlternateAugmentationSealer | 53 | Allows RoF+ clients to augment items from a special container type |
| Inventory:LazyLoadBank | true | Don't load bank during zoning, only when in proximinity to a banker. May increase zone speed and stability |

## Items Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Items:DisableAttuneable | false | Enable this to disable Attuneable Items |
| Items:DisableBardFocusEffects | false | Enable this to disable Bard Focus Effects on Items |
| Items:DisableLore | false | Enable this to disable Lore Items |
| Items:DisableNoDrop | false | Enable this to disable No Drop Items |
| Items:DisableNoPet | false | Enable this to disable No Pet Items |
| Items:DisableNoRent | false | Enable this to disable No Rent Items |
| Items:DisableNoTransfer | false | Enable this to disable No Transfer Items |
| Items:DisablePotionBelt | false | Enable this to disable Potion Belt Items |
| Items:DisableSpellFocusEffects | false | Enable this to disable Spell Focus Effects on Items |
| Items:SummonItemAllowInvisibleAugments | false | Enable this to allow augments to be put in invisible augment slots of items in Client::SummonItem |
| Items:AugmentItemAllowInvisibleAugments | false | Enable this to allow augments to be put in invisible augment slots by players |

## Logging Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Logging:PrintFileFunctionAndLine | false | Ex: [World Server] [net.cpp::main:309] Loading variables... |
| Logging:WorldGMSayLogging | true | Relay worldserver logging to zone processes via GM say output |
| Logging:PlayerEventsQSProcess | false | Have query server process player events instead of world. Useful when wanting to use a dedicated server and database for processing player events on separate disk |
| Logging:BatchPlayerEventProcessIntervalSeconds | 5 | This is the interval in which player events are processed in world or qs |
| Logging:BatchPlayerEventProcessChunkSize | 10000 | This is the cap of events that can be inserted into the queue before a force flush. This is to keep from hitting MySQL max_allowed_packet and killing the connection |

## Mail Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Mail:EnableMailSystem | true | Setting whether the mail system is activated. If false, client won't bring up the Mail window |
| Mail:ExpireTrash | 0 | Setting when the mail trash is emptied. Time in seconds. 0 will delete all messages in the trash when the mailserver starts |
| Mail:ExpireRead | 31536000 | Setting when read mails expire. 31536000=1 Year. Set to -1 for never |
| Mail:ExpireUnread | 31536000 | Setting when unread mails expire. 31536000=1 Year. Set to -1 for never |

## Map Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Map:FixPathingZOnSendTo | false | Try to repair Z coordinates in the SendTo routine as well |
| Map:FixZWhenPathing | true | Automatically fix NPC Z coordinates when moving/pathing/engaged (Far less CPU intensive than its predecessor) |
| Map:MobZVisualDebug | false | Displays spell effects determining whether or not NPC is hitting Best Z calcs (blue for hit, red for miss) |
| Map:MobPathingVisualDebug | false | Displays nodes in pathing points in realtime to help with visual debugging |
| Map:FixPathingZMaxDeltaSendTo | 20 | At runtime in SendTo: maximum change in Z to allow the BestZ code to apply |
| Map:FindBestZHeightAdjust | 1 | Adds this to the current Z before seeking the best Z position |

## Merchant Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Merchant:UsePriceMod | true | Use faction/charisma price modifiers |
| Merchant:UseClassicPriceMod | false | Must also set UsePriceMod. Negates other rules for vendor price mods. |
| Merchant:SellCostMod | 1.05 | Modifier for NPC sell price |
| Merchant:BuyCostMod | 0.95 | Modifier for NPC buy price |
| Merchant:PriceBonusPct | 4 | Determines maximum price bonus from having good faction/CHA. Value is a percent |
| Merchant:PricePenaltyPct | 4 | Determines maximum price penalty from having bad faction/CHA. Value is a percent |
| Merchant:ChaBonusMod | 3.45 | Determines CHA cap, from 104 CHA. 3.45 is 132 CHA at apprehensive. 0.34 is 400 CHA at apprehensive |
| Merchant:ChaPenaltyMod | 1.52 | Determines CHA bottom, up to 102 CHA. 1.52 is 37 CHA at apprehensive. 0.98 is 0 CHA at apprehensive |
| Merchant:EnableAltCurrencySell | true | Enables the ability to resell items to alternate currency merchants |
| Merchant:AllowCorpse | false | Setting whether dealers leave a corpse behind  |

## Mercs Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Mercs:SuspendIntervalMS | 10000 | Time interval for merc suspend (milliseconds) |
| Mercs:UpkeepIntervalMS | 180000 | Time interval for merc upkeep (milliseconds) |
| Mercs:SuspendIntervalS | 10 | Time interval for merc suspend command (seconds) |
| Mercs:AllowMercs | false | Allow the use of mercs |
| Mercs:ChargeMercPurchaseCost | false | Turns Mercenary purchase costs on or off |
| Mercs:ChargeMercUpkeepCost | false | Turns Mercenary upkeep costs on or off |
| Mercs:AggroRadius | 100 | Determines the distance from which a merc will aggro group member's target(also used to determine the distance at which a healer merc will begin healing a group member) |
| Mercs:AggroRadiusPuller | 25 | Determines the distance from which a merc will aggro group member's target, if they have the group role of puller (also used to determine the distance at which a healer merc will begin healing a group member, if they have the group role of puller) |
| Mercs:ResurrectRadius | 50 | Determines the distance from which a healer merc will attempt to resurrect a group member's corpse |
| Mercs:ScaleRate | 100 | Merc scale factor |
| Mercs:AllowMercSuspendInCombat | true | Allow merc suspend in combat |
| Mercs:MercsIgnoreLevelBasedHasteCaps | false | Ignores hard coded level based haste caps. |
| Mercs:MercsHasteCap | 100 | Haste cap for non-v3(over haste) haste |
| Mercs:MercsHastev3Cap | 25 | Haste cap for v3(over haste) haste |

## NPC Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| NPC:MinorNPCCorpseDecayTime | 450000 | NPC corpse decay time, if NPC below level 55 (milliseconds) |
| NPC:MajorNPCCorpseDecayTime | 1500000 | NPC corpse decay time, if NPC equal or greater than level 55 (milliseconds) |
| NPC:CorpseUnlockTimer | 150000 | Time after which corpses are unlocked for everyone to loot (milliseconds) |
| NPC:EmptyNPCCorpseDecayTime | 0 | NPC corpse decay time, if no items are left on the corpse (milliseconds) |
| NPC:UseItemBonusesForNonPets | true | Switch whether item bonuses should be used for NPCs who are not pets |
| NPC:UseBaneDamage | false | If NPCs can't inherently hit the target we don't add bane/magic dmg which isn't exactly the same as PCs |
| NPC:SayPauseTimeInSec | 5 | Time span in which an NPC pauses his movement after a Say event without aggro (seconds) |
| NPC:OOCRegen | 0 | Enable out-of-combat regeneration for NPC |
| NPC:BuffFriends | false | Setting whether NPC should buff other NPC |
| NPC:EnableNPCQuestJournal | false | Setting whether the NPC Quest Journal is active |
| NPC:LastFightingDelayMovingMin | 10000 | Minimum time before mob goes home after all aggro loss (milliseconds) |
| NPC:LastFightingDelayMovingMax | 20000 | Maximum time before mob goes home after all aggro loss (milliseconds) |
| NPC:SmartLastFightingDelayMoving | true | When true, mobs that started going home previously will do so again immediately if still on FD hate list |
| NPC:ReturnNonQuestNoDropItems | false | Returns NO DROP items on NPC that don't have an EVENT_TRADE sub in their script |
| NPC:ReturnQuestItemsFromNonQuestNPCs | false | Returns Quest items traded to NPCs that are not flagged as a Quest NPC |
| NPC:StartEnrageValue | 9 |  Percentage HP that an NPC will begin to enrage |
| NPC:LiveLikeEnrage | false | If set to true then only player controlled pets will enrage |
| NPC:EnableMeritBasedFaction | false | If set to true, faction will be given in the same way as experience (solo/group/raid) |
| NPC:NPCToNPCAggroTimerMin | 500 | Minimum time span after which one NPC aggro another NPC (milliseconds) |
| NPC:NPCToNPCAggroTimerMax | 6000 | Maximum time span after which one NPC aggro another NPC (milliseconds) |
| NPC:UseClassAsLastName | true | Uses class archetype as LastName for NPC with none |
| NPC:NewLevelScaling | true | Better level scaling, use old if new formulas would break your server |
| NPC:NPCBackstabMod | 1.9 | Multiplier for NPC Backstab, Higher = Lower backstab amount |
| NPC:NPCGatePercent | 20 |  Percentage at which the NPC Will attempt to gate at |
| NPC:NPCGateNearBind | false | Will NPC attempt to gate when near bind location? |
| NPC:NPCGateDistanceBind | 75 | Distance from bind before NPC will attempt to gate |
| NPC:NPCHealOnGate | true | Will the NPC Heal on Gate |
| NPC:UseMeditateBasedManaRegen | false | Based NPC ooc regen on Meditate skill |
| NPC:NPCHealOnGateAmount | 25 | How much the NPC will heal on gate if enabled |
| NPC:AnimalsOpenDoors | true | Determines or not whether animals open doors or not when they approach them |
| NPC:MaxRaceID | 732 | Maximum Race ID, RoF2 by default supports up to 732 |
| NPC:DisableLastNames | false | Enable to disable NPC Last Names |
| NPC:NPCIgnoreLevelBasedHasteCaps | false | Ignores hard coded level based haste caps. |
| NPC:NPCHasteCap | 150 | Haste cap for non-v3(over haste) haste |
| NPC:NPCHastev3Cap | 25 | Haste cap for v3(over haste) haste |

## Network Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Network:ResendDelayBaseMS | 100 | Base delay for resending data in EQStreamManager (milliseconds) |
| Network:ResendDelayFactor | 1.5 | Multiplier for the base delay when resending data in EQStreamManager |
| Network:ResendDelayMinMS | 300 | Minimum timespan between two send retries (milliseconds) |
| Network:ResendDelayMaxMS | 5000 | Maximum timespan between two send retries (milliseconds) |
| Network:ClientDataRate | 0.0 | KB / sec, 0.0 disabled |
| Network:CompressZoneStream | true | Setting whether the zone stream should be compressed for transmission |

## Parcel Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Parcel:EnableParcelMerchants | true | Enable or Disable Parcel Merchants.  Requires RoF+ Clients. |
| Parcel:EnableDirectToInventoryDelivery | false | Enable or Disable RoF2 bazaar purchases to be delivered directly to the buyer's inventory. |
| Parcel:DeleteOnDuplicate | false | Delete retrieved item if it creates a lore conflict. |
| Parcel:EnablePruning | false | Enable the automatic pruning of sent parcels.  Uses rule ParcelPruneDelay for prune delay. |
| Parcel:ParcelDeliveryDelay | 30000 | Sets the time that a player must wait between sending parcels. |
| Parcel:ParcelMaxItems | 50 | The maximum number of parcels a player is allowed to have in their mailbox. |
| Parcel:ParcelPruneDelay | 30 | The number of days after which a parcel is deleted. Items are lost! |

## Pathing Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Pathing:Find | true | Enable pathing for FindPerson requests from the client |
| Pathing:Fear | true | Enable pathing for fear |
| Pathing:NavmeshStepSize | 100.0f | Step size for the movement manager |
| Pathing:ShortMovementUpdateRange | 130.0f | Range for short movement updates |
| Pathing:MaxNavmeshNodes | 4092 | Maximum navmesh nodes in a traversable path |

## Pets Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Pets:AttackCommandRange | 150 | Range at which a pet will respond to attack commands |
| Pets:UnTargetableSwarmPet | false | Setting whether swarm pets should be targetable |
| Pets:PetPowerLevelCap | 10 | Maximum number of levels a player pet can go up with pet power |
| Pets:CanTakeNoDrop | false | Setting whether anyone can give no-drop items to pets |
| Pets:CanTakeQuestItems | true | Setting whether anyone can give quest items to pets |
| Pets:LivelikeBreakCharmOnInvis | true | Default: true will break charm on any type of invis (hide/ivu/iva/etc) false will only break if the pet can not see you (ex. you have an undead pet and cast IVU |
| Pets:ClientPetsUseOwnerNameInLastName | true | Disable this to keep client pet's last names from being owner_name's pet |
| Pets:PetTauntRange | 150 | Range at which a pet will taunt targets. |

## QueryServ Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| QueryServ:PlayerLogChat | false | Log player chat |
| QueryServ:PlayerLogTrades | false | Log player trades |
| QueryServ:PlayerDropItems | false | Log player dropping items |
| QueryServ:PlayerLogHandins | false | Log player hand ins |
| QueryServ:PlayerLogNPCKills | false | Log player NPC kills |
| QueryServ:PlayerLogDeletes | false | Log player deletes |
| QueryServ:PlayerLogMoves | false | Log player moves |
| QueryServ:PlayerLogMerchantTransactions | false | Log merchant transactions |
| QueryServ:PlayerLogZone | false | Log player zone events |
| QueryServ:PlayerLogDeaths | false | Log player deaths |
| QueryServ:PlayerLogConnectDisconnect | false | Logs player connect/disconnect state |
| QueryServ:PlayerLogLevels | false | Log player leveling/deleveling |
| QueryServ:PlayerLogAARate | false | Log player AA experience rates |
| QueryServ:PlayerLogQGlobalUpdate | false | Log player QGlobal updates |
| QueryServ:PlayerLogTaskUpdates | false | Log player Task updates |
| QueryServ:PlayerLogAAPurchases | false | Log player AA purchases |
| QueryServ:PlayerLogTradeSkillEvents | false | Log player tradeskill transactions |
| QueryServ:PlayerLogIssuedCommandes | false | Log player issued commands |
| QueryServ:PlayerLogAlternateCurrencyTransactions | false | Log player alternate currency transactions |

## Range Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Range:Say | 15 | The range that is required before /say or hail messages will work to an NPC |
| Range:Emote | 135 | The packet range in which emote messages are sent' |
| Range:BeginCast | 200 | The packet range in which begin cast messages are sent |
| Range:Anims | 135 | The packet range in which begin cast messages are sent |
| Range:SpellParticles | 135 | The packet range in which spell particles are sent |
| Range:DamageMessages | 50 | The packet range in which damage messages are sent (non-crit) |
| Range:SpellMessages | 75 | The packet range in which spell damage messages are sent |
| Range:SongMessages | 75 | The packet range in which song messages are sent |
| Range:StunMessages | 75 | The packet range in which stun messages are sent |
| Range:ClientPositionUpdates | 300 | Distance in which the own changed position is communicated to other clients |
| Range:CriticalDamage | 80 | The packet range in which critical hit messages are sent |
| Range:MobCloseScanDistance | 600 | Close scan distance |
| Range:MaxDistanceToClickDoors | 100 | Max distance that a client can click a door from (Client says 'You can't reach that' at roughly 25-50 for most doors) |

## Skills Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Skills:MaxTrainTradeskills | 21 | Highest level for trade skills that can be taught by the trainer |
| Skills:MaxTrainResearch | 21 | Highest level for training research from a GM. |
| Skills:UseLimitTradeskillSearchSkillDiff | true | Enables the limit for the maximum difference between trivial and skill for recipe searches and favorites |
| Skills:TrivialTradeskillCombinesNoFail | false | Enable to make all trivial tradeskill combines unable to fail |
| Skills:MaxTradeskillSearchSkillDiff | 50 | The maximum difference in skill between the trivial of an item and the skill of the player if the trivial is higher than the skill. Recipes that have not been learnt or made at least once via the Experiment mode will be removed from searches based on this criteria. |
| Skills:MaxTrainSpecializations | 50 | Maximum level a GM trainer will train casting specializations |
| Skills:SwimmingStartValue | 100 | Start value of swimming skill |
| Skills:TrainSenseHeading | false | Switch whether SenseHeading is trained by use |
| Skills:SenseHeadingStartValue | 200 | Start value of sense heading skill |
| Skills:SelfLanguageLearning | true | Enabling self-learning of languages |
| Skills:RequireTomeHandin | false | Disable click-to-learn and force hand in to Guild Master |
| Skills:TradeSkillClamp | 0 | Legacy tradeskills would clamp at 252 regardless of item modifiers and skill combination. DEFAULT: 0 will bypass clamp. Legacy value 252 |
| Skills:UseAltSinisterStrikeFormula | false | Enabling will utilize a formula derived from 2004 monkey business post which makes the AA actually worth something. |
| Skills:TrackingAutoRefreshSkillUps | true | Disable to prevent tracking auto-refresh from giving skill-ups. Classic Style |
| Skills:MaximumTauntDistance | 150 | Maximum player taunt distance. |

## Spells Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Spells:AllowExtraDmgSkill | false | Allow ExtraDmgSkill from Items, Spells, and AAs to apply ExtraDmgAmt when the ExtraDmgSkill matches the casted Spells Skill Type. |
| Spells:BaseCritChance | 0 | Base percentage chance that everyone has to crit a spell |
| Spells:BaseCritRatio | 100 | Base percentage bonus to damage on a successful spell crit. 100=2xdamage |
| Spells:WizCritLevel | 12 | Level wizards first get spell crits |
| Spells:WizCritChance | 7 | Wizards crit chance, on top of BaseCritChance |
| Spells:WizCritRatio | 0 | Wizards crit bonus, on top of BaseCritRatio (should be 0 for Live-like) |
| Spells:TranslocateTimeLimit | 0 | If not zero, time in seconds to accept a Translocate |
| Spells:SacrificeMinLevel | 46 | First level the spell Sacrifice will work on |
| Spells:SacrificeMaxLevel | 69 | Last level the spell Sacrifice will work on |
| Spells:SacrificeItemID | 9963 | Item ID of the item Sacrifice will return (defaults to an Essence Emerald) |
| Spells:EnableSpellGlobals | false | If enabled, spells check the spell_globals table and compare character data from their quest globals before allowing the spell to scribe with scribespells/traindiscs |
| Spells:EnableSpellBuckets | false | If enabled, spells check the spell_buckets table and compare character data from their data buckets before allowing the spell to scribe with scribespells/traindiscs |
| Spells:MaxBuffSlotsNPC | 60 | Maximum number of NPC buff slots. The default value is the limit of the Titanium client |
| Spells:MaxSongSlotsNPC | 0 | Maximum number of NPC song slots. NPC don't have songs, so it should be 0 |
| Spells:MaxDiscSlotsNPC | 0 | Maximum number of NPC disc slots. NPC don't have discs, so it should be 0 |
| Spells:MaxTotalSlotsNPC | 60 | Maximum total of NPC slots. The default value is the limit of the Titanium client |
| Spells:MaxTotalSlotsPET | 30 | Maximum total of pet slots. The default value is the limit of the Titanium client |
| Spells:ReflectType | 4 | Reflect type. 0=disabled, 1=single target player spells only, 2=all player spells, 3=all single target spells, 4=all spells |
| Spells:ReflectMessagesClose | true | True (Live functionality) is for Reflect messages to show to players within close proximity. False shows just player reflecting |
| Spells:LiveLikeFocusEffects | true | Determines whether specific healing, dmg and mana reduction focuses are randomized |
| Spells:BaseImmunityLevel | 55 | The level that targets start to be immune to stun, fear and mez spells with a maximum level of 0 |
| Spells:NPCIgnoreBaseImmunity | true | Whether or not NPC get to ignore the BaseImmunityLevel for their spells |
| Spells:AvgSpellProcsPerMinute | 6.0 | Adjust rate for sympathetic spell procs |
| Spells:ResistFalloff | 67 | Maximum that level that will adjust our resist chance based on level modifiers |
| Spells:CharismaEffectiveness | 10 | Determines how much resist modification charisma applies to charm/pacify checks. Default 10 CHA = -1 resist mod |
| Spells:CharismaEffectivenessCap | 255 | Determines how much resist modification charisma applies to charm/pacify checks. Default 10 CHA = -1 resist mod |
| Spells:CharismaCharmDuration | false | Allow CHA resist mod to extend charm duration |
| Spells:CharmBreakCheckChance | 25 | Determines chance for a charm break check to occur each buff tick |
| Spells:CharmDisablesSpecialAbilities | false | When charm is cast on an NPC, strip their special abilities |
| Spells:RootBreakFromSpells | 55 | Chance for root to break when cast on |
| Spells:DeathSaveCharismaMod | 3 | Determines how much charisma effects chance of death save firing |
| Spells:DivineInterventionHeal | 8000 | Divine intervention heal amount |
| Spells:AdditiveBonusWornType | 0 | Calc worn bonuses to add together (instead of taking highest) if set to THIS worn type. (2=Will covert live items automatically) |
| Spells:UseCHAScribeHack | false | ScribeSpells and TrainDiscs quest functions will ignore entries where field 12 is CHA |
| Spells:BuffLevelRestrictions | true | Buffs will not land on low level toons like live |
| Spells:RootBreakCheckChance | 70 | Determines chance for a root break check to occur each buff tick |
| Spells:FearBreakCheckChance | 70 | Determines chance for a fear break check to occur each buff tick |
| Spells:SuccorFailChance | 2 | Determines chance for a succor spell not to teleport an invidual player |
| Spells:FRProjectileItem_Titanium | 1113 | Item id for Titanium clients for Fire 'spell projectile' |
| Spells:FRProjectileItem_SOF | 80684 | Item id for SOF clients for Fire 'spell projectile' |
| Spells:FRProjectileItem_NPC | 80684 | Item id for NPC Fire 'spell projectile' |
| Spells:UseLiveSpellProjectileGFX | false | Use spell projectile graphics set in the spells_new table (player_1). Server must be using UF+ spell file |
| Spells:FocusCombatProcs | false | Allow all combat procs to receive focus effects |
| Spells:PreNerfBardAEDoT | false | Allow bard AOE dots to damage targets when moving |
| Spells:AI_SpellCastFinishedFailRecast | 800 | AI spell recast time  when an spell is cast but fails, ie if stunned (milliseconds) |
| Spells:AI_EngagedNoSpellMinRecast | 500 | AI spell recast time check when no spell is cast while engaged. Min time in random (milliseconds) |
| Spells:AI_EngagedNoSpellMaxRecast | 1000 | AI spell recast time check when no spell is cast engaged. Mmaximum time in random (milliseconds) |
| Spells:AI_EngagedBeneficialSelfChance | 100 | Chance during first AI Cast check to do a beneficial spell on self |
| Spells:AI_EngagedBeneficialOtherChance | 25 | Chance during second AI Cast check to do a beneficial spell on others |
| Spells:AI_EngagedDetrimentalChance | 20 | Chance during third AI Cast check to do a determental spell on others |
| Spells:AI_PursueNoSpellMinRecast | 500 | AI spell recast time check when no spell is cast while chasing target. Mmin time in random (milliseconds) |
| Spells:AI_PursueNoSpellMaxRecast | 2000 | AI spell recast time check when no spell is cast while chasing target. Maximum time in random (milliseconds) |
| Spells:AI_PursueDetrimentalChance | 90 | Chance while chasing target to cast a detrimental spell |
| Spells:AI_IdleNoSpellMinRecast | 6000 | AI spell recast time check when no spell is cast while idle. Mmin time in random (milliseconds) |
| Spells:AI_IdleNoSpellMaxRecast | 60000 | AI spell recast time check when no spell is cast while chasing target. Maximum time in random (milliseconds) |
| Spells:AI_IdleBeneficialChance | 100 | Chance while idle to do a beneficial spell on self or others |
| Spells:AI_HealHPPct | 50 | Hitpoint percentage at which NPC starts healing when max_hp of the spell is not set (inside and outside combat) |
| Spells:SHDProcIDOffByOne | true | Pre June 2009 SHD spell procs were off by 1, they stopped doing this in June 2009 (UF+ spell files need this false) |
| Spells:Jun182014HundredHandsRevamp | false | This should be true for if you import a spell file newer than June 18, 2014 |
| Spells:SwarmPetTargetLock | false | Use old method of swarm pets target locking till target dies then despawning |
| Spells:NPC_UseFocusFromSpells | true | Allow NPC to use most spell derived focus effects |
| Spells:NPC_UseFocusFromItems | false | Allow NPC to use most item derived focus effects |
| Spells:UseAdditiveFocusFromWornSlot | false | Allows an additive focus effect to be calculated from worn slot. Does not apply limits checks. Can only have one additive focus rule be true. |
| Spells:UseAdditiveFocusFromWornSlotWithLimits | false | Allows an additive focus effect to be calculated from worn slot. Applies normal limit checks. Can only have one additive focus rule be true. |
| Spells:AlwaysSendTargetsBuffs | false | Ignore Leadership Alternate Abilities level if true |
| Spells:FlatItemExtraSpellAmt | false | Allow SpellDmg stat to affect all spells, regardless of cast time/cooldown/etc |
| Spells:IgnoreSpellDmgLvlRestriction | false | Ignore the 5 level spread on applying SpellDmg |
| Spells:ItemExtraSpellAmtCalcAsPercent | false | Apply the Item stats Spell Dmg and Heal Amount as Percentage-based modifiers instead of flat additions |
| Spells:BreakFeignDeathWhenCastOn | 2.0 | Percentage that Feign Death will break when you resist a spell |
| Spells:BreakSneakWhenCastOn | 2.0 | Percentage that Sneak will break when you resist a spell |
| Spells:EnableResistSoftCap | false | Enabled resist softcap and can be adjusted by rule, SpellResistSoftCap |
| Spells:SpellResistSoftCap | 255 | Softcap for spell resists. |
| Spells:AllowItemTGB | false | Target group buff (/tgb) doesn't work with items on live, custom servers want it though |
| Spells:NPCInnateProcOverride | true | NPC innate procs override the target type to single target |
| Spells:OldRainTargets | false | Use old incorrectly implemented maximum targets for rains |
| Spells:CallOfTheHeroAggroClearDist | 250.0 | Distance at which CoTH will wipe aggro. To disable and always enable aggro wipe on any distance of CoTH, set to 0. |
| Spells:NPCSpellPush | false | Enable spell push on NPCs |
| Spells:July242002PetResists | true | Enable Pets using PCs resist change from July 24 2002 |
| Spells:CazicTouchTargetsPetOwner | true | If True, causes Cazic Touch to swap targets from pet to pet owner if a pet is tanking. |
| Spells:PreventFactionWarOnCharmBreak | false | Enable spell interupts and dot removal on charm break to prevent faction wars. |
| Spells:AllowDoubleInvis | true | Allows you to cast invisibility spells on a player that is already invisible, live like behavior. |
| Spells:AllowSpellMemorizeFromItem | false | Allows players to memorize spells by right-clicking spell scrolls |
| Spells:InvisRequiresGroup | false | Invis requires the the target to be in group. |
| Spells:ClericInnateHealFocus | 5 | Clerics on live get a 5 pct innate heal focus |
| Spells:DOTsScaleWithSpellDmg | false | Allow SpellDmg stat to affect DoT spells |
| Spells:HOTsScaleWithHealAmt | false | Allow HealAmt stat to affect HoT spells |
| Spells:CompoundLifetapHeals | true | True: Lifetap heals calculate damage bonuses and then heal bonuses.  False:  Lifetaps heal using the amount damaged to mob. |
| Spells:UseFadingMemoriesMaxLevel | false | Enables to limit field in spell data to set the max level that over which an NPC will ignore fading memories effect and not lose aggro. |
| Spells:FixBeaconHeading | false | Beacon spells use casters heading to fix live bug.  False: Live like heading always 0. |
| Spells:UseSpellImpliedTargeting | false | Replicates EQ2-style targeting behavior for spells. Spells will 'pass through' inappropriate targets to target's target if it is appropriate. |
| Spells:BuffsFadeOnDeath | true | Disable to keep buffs from fading on death |
| Spells:IllusionsAlwaysPersist | false | Allows Illusions to persist beyond death and zoning always. |
| Spells:UseItemCastMessage | false | Enable to use the \"item begins to glow\" messages when casting from an item. |
| Spells:TargetsTargetRequiresCombatRange | true | Disable to remove combat range requirement from Target's Target Spell Target Type |
| Spells:NPCBuffLevelRestrictions | false | Impose BuffLevelRestrictions on NPCs if true |
| Spells:ResurrectionEffectBlock | 2 | 0 = allow overwrites/rule disabled. If set to 1 = Block all buffs that would overwrite Resurrection Effects. If set to 2 = Will not overwrite Resurrection Effects, instead moves new buff to an empty slot if available. Default is 2. |
| Spells:WaterMatchRequiredForLoS | true | Enable/Disable the requirement of both the attacker/victim being both in or out of water for spells LoS to pass. |
| Spells:WizardCritMinimumRandomRatio | 20 | The minimum value for the random range which Wizards and Caster DPS Mercs innately have for spell crit ratio. Set to 20 for vanilla values. |
| Spells:WizardCritMaximumRandomRatio | 70 | The maximum value for the random range which Wizards and Caster DPS Mercs innately have for spell crit ratio. Set to 70 for vanilla values. |
| Spells:DefensiveProcPenaltyLevelGap | 6 | Defensive Proc Penalty Level Gap where procs start losing their proc rate at RuleR(Spells, DefensiveProcPenaltyModifier)% per level difference |
| Spells:DefensiveProcPenaltyLevelGapModifier | 10.0f | Defensive Proc Penalty Level Gap Modifier where procs start losing their proc rate at defined % after RuleI(Spells, DefensiveProcLevelGap) level difference |
| Spells:DOTBonusDamageSplitOverDuration | true | Disable to have Damage Over Time total bonus damage added to each tick instead of divided across duration |
| Spells:HOTBonusHealingSplitOverDuration | true | Disable to have Heal Over Time total bonus healing added to each tick instead of divided across duration |
| Spells:UseLegacyFizzleCode | false | Enable will turn on the legacy fizzle code which is far stricter and more accurate to 2001/2002 testing. |
| Spells:LegacyManaburn | false | Enable to have the legacy manaburn system from 2003 and earlier. |
| Spells:LegacyManaburnCap | 9492 | Adjusted the hard cap (Normal or Crit) for the Legacy Manaburn system. DEFAULT: 9492 |
| Spells:EvacClearAggroInSameZone | false | Enable to clear aggro on clients when evacing in same zone. |
| Spells:CharmAggroOverLevel | false | Enabling this rule will cause Charm casts over level to show resisted and cause aggro. Early EQ style. |
| Spells:RequireMnemonicRetention | true | Enabling will require spell slots 9-12 to have the appropriate Mnemonic Retention AA learned. |
| Spells:EvacClearCharmPet | false | Enable to have evac in zone clear charm from charm pets and detach buffs. |
| Spells:ManaTapsRequireNPCMana | false | Enabling will require target to have mana to tap. Default off as many npc's are caster class with 0 mana and need fixed. |
| Spells:HarmTouchCritRatio | 200 | Harmtouch crit bonus, on top of BaseCritRatio |
| Spells:UseClassicHarmTouchDamage | false | Use pre 2007 Harm Touch calculations - Default: False |
| Spells:UseClassicSpellFocus | false | Enabling will tell the server to handle random focus damage as classic spell imports lack the limit values. |
| Spells:ManaTapsOnAnyClass | false | Enabling this will allow you to cast mana taps on any class, this will bypass ManaTapsRequireNPCMana rule. |
| Spells:HealAmountMessageFilterThreshold | 100 | Lifetaps below this threshold will not have a message sent to the client (Heal will still process) 0 to Disable. |
| Spells:SnareOverridesSpeedBonuses | false | Enabling will allow snares to override any speed bonuses the entity may have. Default: False |
| Spells:TargetedAOEMaxTargets | 4 | Max number of targets a Targeted AOE spell can cast on. Set to 0 for no limit. |
| Spells:PointBlankAOEMaxTargets | 0 | Max number of targets a Point-Blank AOE spell can cast on. Set to 0 for no limit. |
| Spells:DefaultAOEMaxTargets | 0 | Max number of targets that an AOE spell which does not meet other descriptions can cast on. Set to 0 for no limit. |
| Spells:AllowFocusOnSkillDamageSpells | false | Allow focus effects 185, 459, and 482 to enhance SkillAttack spell effect 193 |

## TaskSystem Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| TaskSystem:EnableTaskSystem | true | Globally enable or disable the Task system |
| TaskSystem:PeriodicCheckTimer | 5 | Seconds between checks for failed tasks. Also used by the 'Touch' activity |
| TaskSystem:RecordCompletedTasks | true | Record completed tasks |
| TaskSystem:RecordCompletedOptionalActivities | false | Record completed optional activities |
| TaskSystem:KeepOneRecordPerCompletedTask | true | Keep only one record per completed task |
| TaskSystem:EnableTaskProximity | true | Enable task proximity system |
| TaskSystem:RequestCooldownTimerSeconds | 15 | Seconds between allowing characters to request tasks (live-like default: 15 seconds) |
| TaskSystem:ExpRewardsIgnoreLevelBasedEXPMods | false | Rewarding Level Based Exp will ignore Rule LevelBasedEXPMods if true |
| TaskSystem:SharedTasksWorldProcessRate | 6000 | Timer interval (milliseconds) that shared tasks are processed in world |
| TaskSystem:SharedTasksTerminateTimerMS | 120000 | Delay (milliseconds) until a shared task is terminated if requirements are no longer met after member removal (default: 2 minutes) |
| TaskSystem:UpdateOneElementPerTask | true | If true (live-like) task updates only increment the first matching activity. If false all matching elements will be incremented. |
| TaskSystem:MaxUpdateMessages | 50 | Maximum update messages for non-GiveCash activity types in IncrementDoneCount |

## Watermap Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Watermap:CheckForWaterOnSendTo | false | Checks if a mob has moved into/out of water on SendTo |
| Watermap:CheckForWaterWhenFishing | false | Only lets a player fish near water (if a water map exists for the zone) |
| Watermap:FishingRodLength | 30 | How far in front of player water must be for fishing to work |
| Watermap:FishingLineLength | 100 | If water is more than this far below the player, it is considered too far to fish |
| Watermap:FishingLineStepSize | 1 | Basic step size for fishing calc, too small and it will eat cpu, too large and it will miss potential water |

## World Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| World:ZoneAutobootTimeoutMS | 60000 | Time out for automatic booting of zones in milliseconds |
| World:UseBannedIPsTable | false | Toggle whether or not to check incoming client connections against the banned_ips table. Set this value to false to disable this feature |
| World:EnableTutorialButton | true | Setting whether the Tutorial button should be active. At least in RoF2 you can always press the button, but it loses its effect |
| World:EnableReturnHomeButton | true | Setting whether the Return Home button should be active |
| World:MaxLevelForTutorial | 10 | The highest level with which you can enter the tutorial |
| World:TutorialZoneID | 189 | Zone ID of the tutorial |
| World:MinOfflineTimeToReturnHome | 21600 | Minimum offline time to activate the Return Home button. 21600 seconds is 6 Hours |
| World:MaxClientsPerIP | -1 | Maximum number of clients allowed to connect per IP address if account status is < AddMaxClientsStatus. Default value: -1 (feature disabled) |
| World:ExemptMaxClientsStatus | -1 | Exempt accounts from the MaxClientsPerIP and AddMaxClientsStatus rules, if their status is >= this value. Default value: -1 (feature disabled) |
| World:AddMaxClientsPerIP | -1 | Maximum number of clients allowed to connect per IP address if account status is < ExemptMaxClientsStatus. Default value: -1 (feature disabled) |
| World:AddMaxClientsStatus | -1 | Accounts with status >= this rule will be allowed to use the amount of accounts defined in the AddMaxClientsPerIP. Default value: -1 (feature disabled) |
| World:MaxClientsSetByStatus | false | If true, IP Limiting will be set to the status on the account as long as the status is > MaxClientsPerIP |
| World:EnableIPExemptions | false | If true, ip_exemptions table is used, if there is no entry for the IP it will default to RuleI(World, MaxClientsPerIP) |
| World:ClearTempMerchantlist | true | Clears temp merchant items when world boots |
| World:GMAccountIPList | false | Check IP list against GM accounts. This increases the security of GM accounts, e.g. if you only allow localhost '127.0.0.1' for GM accounts. Think carefully about what you enter! |
| World:MinGMAntiHackStatus | 1 | Minimum status to check against AntiHack list |
| World:SoFStartZoneID | -1 | Sets the Starting Zone for SoF Clients separate from Titanium Clients (-1 is disabled) |
| World:TitaniumStartZoneID | -1 | Sets the Starting Zone for Titanium Clients (-1 is disabled). Replaces the old method |
| World:ExpansionSettings | 16383 | Sets the expansion settings bitmask for the server, This is sent on login to world and affects client expansion settings. Defaults to all expansions enabled up to TSS, value is bitmask |
| World:CharacterSelectExpansionSettings | -1 | Sets the expansion settings bitmask for character select if you wish to override. -1 is off |
| World:UseClientBasedExpansionSettings | true | If true it will overrule World, ExpansionSettings and set someone's expansion based on the client they're using |
| World:PVPSettings | 0 | Sets the PVP settings for the server. 1=Rallos Zek RuleSet, 2=Tallon/Vallon Zek Ruleset, 4=Sullon Zek Ruleset, 6=Discord Ruleset, anything above 6 is the Discord Ruleset without the no-drop restrictions removed. NOTE: edit IsAttackAllowed in Zone-table to accomodate for these rules |
| World:PVPMinLevel | 0 | Minimum level to pvp |
| World:StartZoneSameAsBindOnCreation | true | Should the start zone always be the same location as your bind? |
| World:EnforceCharacterLimitAtLogin | false | Enforce the limit for characters that are online at login |
| World:EnableDevTools | true | Enable or Disable the Developer Tools globally (Most of the time you want this enabled) |
| World:EnableChecksumVerification | false | Enable or Disable the Checksum Verification for eqgame.exe and spells_us.txt |
| World:MaximumQuestErrors | 30 | Changes the maximum number of quest errors that can be displayed in #questerrors, default is 30 |
| World:BootHour | 0 | Sets the in-game hour world will set when it first boots. 0-24 are valid options, where 0 disables this rule |
| World:UseItemLinksForKeyRing | false | Uses item links for Key Ring Listing instead of item name |
| World:UseOldShadowKnightClassExport | true | Disable to have Shadowknight show as Shadow Knight (live-like) |
| World:EnableAutoLogin | false | Enables or disables auto login of characters, allowing people to log characters in directly from loginserver to ingame |
| World:EnablePVPRegions | true | Enables or disables PVP Regions automatically setting your PVP flag |
| World:CustomFilesAdminLevel | 20 | Admin level at which custom file key is not required when CustomFilesKey is specified |

## Zone Rules
| Rule Name | Default Value | Description |
| :--- | :--- | :--- |
| Zone:ClientLinkdeadMS | 90000 | The time a client remains link dead on the server after a sudden disconnection (milliseconds) |
| Zone:GraveyardTimeMS | 1200000 | Time until a player corpse is moved to a zone's graveyard, if one is specified for the zone (milliseconds) |
| Zone:EnableShadowrest | 1 | Enables or disables the Shadowrest zone feature for player corpses. Default is turned on |
| Zone:AutoShutdownDelay | 60000 | How long a dynamic zone stays loaded while empty (milliseconds) |
| Zone:PEQZoneReuseTime | 900 | Seconds between two uses of the #peqzone command (Set to 0 to disable) |
| Zone:PEQZoneDebuff1 | 4454 | First debuff casted by #peqzone Default is Cursed Keeper's Blight |
| Zone:PEQZoneDebuff2 | 2209 | Second debuff casted by #peqzone Default is Tendrils of Apathy |
| Zone:UsePEQZoneDebuffs | true | Setting if the command #peqzone applies the defined debuffs |
| Zone:PEQZoneHPRatio | 75 | Required HP Ratio to use #peqzone |
| Zone:HotZoneBonus | 0.75 | Value which is added to the experience multiplier. This also applies to AA experience. |
| Zone:EbonCrystalItemID | 40902 | Item ID for Ebon Crystal |
| Zone:RadiantCrystalItemID | 40903 | Item ID for Radiant Crystal |
| Zone:LevelBasedEXPMods | false | Allows you to use the level_exp_mods table in consideration to your players experience hits |
| Zone:WeatherTimer | 600 | Weather timer when no duration is available |
| Zone:EnableLoggedOffReplenishments | true | 'Replenish mana/hp/end if logged off for MinOfflineTimeToReplenishments |
| Zone:MinOfflineTimeToReplenishments | 21600 | Minimum time a player must be offline before LoggedOffReplenishments becomes active (seconds) |
| Zone:UseZoneController | true | Enables the ability to use persistent quest based zone controllers (zone_controller.pl/lua) |
| Zone:EnableZoneControllerGlobals | false | Enables the ability to use quest globals with the zone controller NPC |
| Zone:GlobalLootMultiplier | 1 | Sets Global Loot drop multiplier for database based drops, useful for double, triple loot etc |
| Zone:KillProcessOnDynamicShutdown | true | When process has booted a zone and has hit its zone shut down timer, it will hard kill the process to free memory back to the OS |
| Zone:SpawnEventMin | 3 | When strict is set in spawn_events, specifies the max EQ minutes into the trigger hour a spawn_event will fire. Going below 3 may cause the spawn_event to not fire. |
| Zone:ForageChance | 25 | Chance of foraging from zone table vs global table |
| Zone:FishingChance | 399 | Chance of fishing from zone table vs global table |
| Zone:AllowCrossZoneSpellsOnBots | false | Set to true to allow cross zone spells (cast/remove) to affect bots |
| Zone:AllowCrossZoneSpellsOnMercs | false | Set to true to allow cross zone spells (cast/remove) to affect mercenaries |
| Zone:AllowCrossZoneSpellsOnPets | false | Set to true to allow cross zone spells (cast/remove) to affect pets |
