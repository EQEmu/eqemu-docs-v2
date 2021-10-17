---
description: >-
  This page describes the Rules and Rule Values for your EQEmu Server. These
  rules are found in the rule_values table.
---

# Server Rules

### Editing rules

To edit rules, you can use the \#rules set &lt;key&gt; &lt;value&gt; command to set a rule. 

You can then use \#reloadallrules to force all zones to reload the new rules. 

{% hint style="warning" %}
The `notes` field in the database is set by the server software.  If you adjust the values in the field, they will be overwritten and restored to default values when you update your server binaries.
{% endhint %}

{% hint style="info" %}
If you would like to improve upon the descriptions or notes in the Server Rules table, please submit a pull request on the [ruletypes](https://github.com/EQEmu/Server/blob/master/common/ruletypes.h) header file.
{% endhint %}

| **Type** | **Category** | **Name** | **Default Value** | **Description** |
| :--- | :--- | :--- | :--- | :--- |
| Integer | Character | MaxLevel | 65 |  |
| Boolean | Character | PerCharacterQglobalMaxLevel | FALSE | This will check for qglobal 'CharMaxLevel' character qglobal \(Type 5\), if player tries to level beyond that point, it will not go beyond that level |
| Integer | Character | MaxExpLevel | 0 | Sets the Max Level attainable via Experience |
| Integer | Character | DeathExpLossLevel | 10 | Any level greater than this will lose exp on death |
| Integer | Character | DeathExpLossMaxLevel | 255 | Any level greater than this will no longer lose exp on death |
| Integer | Character | DeathItemLossLevel | 10 |  |
| Integer | Character | DeathExpLossMultiplier | 3 | Adjust how much exp is lost |
| Boolean | Character | UseDeathExpLossMult | FALSE | Adjust to use the above multiplier or to use code default. |
| Integer | Character | CorpseDecayTimeMS | 10800000 |  |
| Integer | Character | CorpseResTimeMS | 10800000 | time before cant res corpse\(3 hours |
| Boolean | Character | LeaveCorpses | TRUE |  |
| Boolean | Character | LeaveNakedCorpses | FALSE |  |
| Integer | Character | MaxDraggedCorpses | 2 |  |
| Decimal | Character | DragCorpseDistance | 400 | If the corpse is &lt;= this distance from the player, it won't move |
| Decimal | Character | ExpMultiplier | 0.5 |  |
| Decimal | Character | AAExpMultiplier | 0.5 |  |
| Decimal | Character | GroupExpMultiplier | 0.5 |  |
| Decimal | Character | RaidExpMultiplier | 0.2 |  |
| Boolean | Character | UseXPConScaling | TRUE |  |
| Integer | Character | ShowExpValues | 0 | 0 - normal, 1 - Show raw experience values, 2 - Show raw experience values AND percent. |
| Integer | Character | LightBlueModifier | 40 |  |
| Integer | Character | BlueModifier | 90 |  |
| Integer | Character | WhiteModifier | 100 |  |
| Integer | Character | YellowModifier | 125 |  |
| Integer | Character | RedModifier | 150 |  |
| Integer | Character | AutosaveIntervalS | 300 | 0=disabled |
| Integer | Character | HPRegenMultiplier | 100 |  |
| Integer | Character | ManaRegenMultiplier | 100 |  |
| Integer | Character | EnduranceRegenMultiplier | 100 |  |
| Integer | Character | ConsumptionMultiplier | 100 | item's hunger restored = this value \* item's food level, 100 = normal, 50 = people eat 2x as fast, 200 = people eat 2x as slow |
| Boolean | Character | HealOnLevel | FALSE |  |
| Boolean | Character | FeignKillsPet | FALSE |  |
| Integer | Character | ItemManaRegenCap | 15 |  |
| Integer | Character | ItemHealthRegenCap | 35 |  |
| Integer | Character | ItemDamageShieldCap | 30 |  |
| Integer | Character | ItemAccuracyCap | 150 |  |
| Integer | Character | ItemAvoidanceCap | 100 |  |
| Integer | Character | ItemCombatEffectsCap | 100 |  |
| Integer | Character | ItemShieldingCap | 35 |  |
| Integer | Character | ItemSpellShieldingCap | 35 |  |
| Integer | Character | ItemDoTShieldingCap | 35 |  |
| Integer | Character | ItemStunResistCap | 35 |  |
| Integer | Character | ItemStrikethroughCap | 35 |  |
| Integer | Character | ItemATKCap | 250 |  |
| Integer | Character | ItemHealAmtCap | 250 |  |
| Integer | Character | ItemSpellDmgCap | 250 |  |
| Integer | Character | ItemClairvoyanceCap | 250 |  |
| Integer | Character | ItemDSMitigationCap | 50 |  |
| Integer | Character | ItemEnduranceRegenCap | 15 |  |
| Integer | Character | ItemExtraDmgCap | 150 | Cap for bonuses to melee skills like Bash, Frenzy, etc |
| Integer | Character | HasteCap | 100 | Haste cap for non-v3\(overhaste\) haste. |
| Integer | Character | SkillUpModifier | 100 | skill ups are at 100% |
| Boolean | Character | SharedBankPlat | FALSE | off by default to prevent duping for now |
| Boolean | Character | BindAnywhere | FALSE |  |
| Integer | Character | RestRegenPercent | 0 | 0 to enable rest state bonus HP and mana regen.  Set to &gt;0 to enable rest state bonus HP and mana regen. |
| Integer | Character | RestRegenTimeToActivate | 30 | Time in seconds for rest state regen to kick in. |
| Integer | Character | RestRegenRaidTimeToActivate | 300 | Time in seconds for rest state regen to kick in with a raid target. |
| Boolean | Character | RestRegenEndurance | FALSE | Whether rest regen will work for endurance or not. |
| Integer | Character | KillsPerGroupLeadershipAA | 250 | Number of dark blues or above per Group Leadership AA |
| Integer | Character | KillsPerRaidLeadershipAA | 250 | Number of dark blues or above per Raid Leadership AA |
| Integer | Character | MaxFearDurationForPlayerCharacter | 4 | 4 tics, each tic calculates every 6 seconds. |
| Integer | Character | MaxCharmDurationForPlayerCharacter | 15 |  |
| Integer | Character | BaseHPRegenBonusRaces | 4352 | a bitmask of race\(s\) that receive the regen bonus. Iksar \(4096\) & Troll \(256\) = 4352. see common/races.h for the bitmask values |
| Boolean | Character | SoDClientUseSoDHPManaEnd | FALSE | Setting this to true will allow SoD clients to use the SoD HP/Mana/End formulas and previous clients will use the old formulas |
| Boolean | Character | UseRaceClassExpBonuses | TRUE | Setting this to true will enable Class and Racial experience rate bonuses |
| Boolean | Character | UseOldRaceExpPenalties | FALSE | Setting this to true will enable racial exp penalties for Iksar, Troll, Ogre, and Barbarian, as well as the bonus for Halflings |
| Boolean | Character | UseOldClassExpPenalties | FALSE | Setting this to true will enable old class exp penalties for Paladin, SK, Ranger, Bard, Monk, Wizard, Enchanter, Magician, and Necromancer, as well as the bonus for Rogues and Warriors |
| Boolean | Character | RespawnFromHover | FALSE | Use Respawn window, or not. |
| Integer | Character | RespawnFromHoverTimer | 300 | Respawn Window countdown timer, in SECONDS |
| Boolean | Character | UseNewStatsWindow | TRUE | New stats window shows everything |
| Boolean | Character | ItemCastsUseFocus | FALSE | If true, this allows item clickies to use focuses that have limited max levels on them |
| Integer | Character | MinStatusForNoDropExemptions | 80 | This allows status x and higher to trade no drop items. |
| Integer | Character | SkillCapMaxLevel | 75 | Sets the Max Level used for Skill Caps \(from skill\_caps table\). -1 makes it use MaxLevel rule value. It is set to 75 because PEQ only has skillcaps up to that level, and grabbing the players' skill past 75 will return 0, breaking all skills past that level. This helps servers with obsurd level caps \(75+ level cap\) function without any modifications. |
| Integer | Character | StatCap | 0 |  |
| Boolean | Character | CheckCursorEmptyWhenLooting | TRUE | If true, a player cannot loot a corpse \(player or NPC\) with an item on their cursor |
| Boolean | Character | MaintainIntoxicationAcrossZones | TRUE | If true, alcohol effects are maintained across zoning and logging out/in. |
| Boolean | Character | EnableDiscoveredItems | TRUE | If enabled, it enables EVENT\_DISCOVER\_ITEM and also saves character names and timestamps for the first time an item is discovered. |
| Boolean | Character | EnableXTargetting | TRUE | Enable Extended Targetting Window, for users with UF and later clients. |
| Boolean | Character | KeepLevelOverMax | FALSE | Don't delevel a character that has somehow gone over the level cap |
| Integer | Character | FoodLossPerUpdate | 35 | How much food/water you lose per stamina update |
| Integer | Character | BaseInstrumentSoftCap | 36 | Softcap for instrument mods, 36 commonly referred to as "3.6" as well. |
| Boolean | Character | UseSpellFileSongCap | TRUE | When they removed the AA that increased the cap they removed the above and just use the spell field |
| Integer | Character | BaseRunSpeedCap | 158 | Base Run Speed Cap, on live it's 158% which will give you a runspeed of 1.580 hard capped to 225. |
| Integer | Character | OrnamentationAugmentType | 20 | Ornamentation Augment Type |
| Decimal | Character | EnvironmentDamageMulipliter | 1 |  |
| Boolean | Character | UnmemSpellsOnDeath | TRUE |  |
| Integer | Character | TradeskillUpAlchemy | 2 | Alchemy skillup rate adjust. Lower is faster. |
| Integer | Character | TradeskillUpBaking | 2 | Baking skillup rate adjust. Lower is faster. |
| Integer | Character | TradeskillUpBlacksmithing | 2 | Blacksmithing skillup rate adjust. Lower is faster. |
| Integer | Character | TradeskillUpBrewing | 3 | Brewing skillup rate adjust. Lower is faster. |
| Integer | Character | TradeskillUpFletching | 2 | Fletching skillup rate adjust. Lower is faster. |
| Integer | Character | TradeskillUpJewelcrafting | 2 | Jewelcrafting skillup rate adjust. Lower is faster. |
| Integer | Character | TradeskillUpMakePoison | 2 | Make Poison skillup rate adjust. Lower is faster. |
| Integer | Character | TradeskillUpPottery | 4 | Pottery skillup rate adjust. Lower is faster. |
| Integer | Character | TradeskillUpResearch | 1 | Research skillup rate adjust. Lower is faster. |
| Integer | Character | TradeskillUpTinkering | 2 | Tinkering skillup rate adjust. Lower is faster. |
| Boolean | Character | SpamHPUpdates | FALSE | if your server has stupid amounts of HP that causes client display issues, turn this on! |
| Boolean | Character | MarqueeHPUpdates | FALSE | Will show Health % in center of screen &lt; 100% |
| Integer | Character | IksarCommonTongue | 95 | 95 By default \(live-like? |
| Integer | Character | OgreCommonTongue | 95 | 95 By default \(live-like? |
| Integer | Character | TrollCommonTongue | 95 | 95 By default \(live-like? |
| Boolean | Character | ActiveInvSnapshots | FALSE | Takes a periodic snapshot of inventory contents from online players |
| Integer | Character | InvSnapshotMinIntervalM | 180 | Minimum time \(in minutes\) between inventory snapshots |
| Integer | Character | InvSnapshotMinRetryM | 30 | Time \(in minutes\) to re-attempt an inventory snapshot after a failure |
| Integer | Character | InvSnapshotHistoryD | 30 | Time \(in days\) to keep snapshot entries |
| Boolean | Character | RestrictSpellScribing | FALSE | Restricts spell scribing to allowable races/classes of spell scroll, if true |
| Boolean | Character | UseStackablePickPocketing | TRUE | Allows stackable pickpocketed items to stack instead of only being allowed in empty inventory slots |
| Boolean | Character | EnableAvoidanceCap | FALSE |  |
| Integer | Character | AvoidanceCap | 750 | 750 Is a pretty good value, seen people dodge all attacks beyond 1,000 Avoidance |
| Boolean | Character | AllowMQTarget | FALSE | Disables putting players in the 'hackers' list for targeting beyond the clip plane or attempting to target something untargetable |
| Boolean | Character | UseOldBindWound | FALSE | Uses the original bind wound behavior |
| Boolean | Character | GrantHoTTOnCreate | FALSE | Grant Health of Target's Target leadership AA on character creation |
| Integer | Expansion | CurrentExpansion | -1 | Sets which server side expansion is enabled. See [https://eqemu.gitbook.io/server/in-development/project-peq-expansions/expansion-content-filtering](https://eqemu.gitbook.io/server/in-development/project-peq-expansions/expansion-content-filtering) for more details. |
| Integer | Mercs | SuspendIntervalMS | 10000 |  |
| Integer | Mercs | UpkeepIntervalMS | 180000 |  |
| Integer | Mercs | SuspendIntervalS | 10 |  |
| Integer | Mercs | UpkeepIntervalS | 180 |  |
| Boolean | Mercs | AllowMercs | FALSE |  |
| Boolean | Mercs | ChargeMercPurchaseCost | FALSE |  |
| Boolean | Mercs | ChargeMercUpkeepCost | FALSE |  |
| Integer | Mercs | AggroRadius | 100 | Determines the distance from which a merc will aggro group member's target\(also used to determine the distance at which a healer merc will begin healing a group member |
| Integer | Mercs | AggroRadiusPuller | 25 | Determines the distance from which a merc will aggro group member's target, if they have the group role of puller \(also used to determine the distance at which a healer merc will begin healing a group member, if they have the group role of puller |
| Integer | Mercs | ResurrectRadius | 50 | Determines the distance from which a healer merc will attempt to resurrect a group member's corpse |
| Integer | Mercs | ScaleRate | 100 |  |
| Integer | Guild | MaxMembers | 2048 |  |
| Boolean | Guild | PlayerCreationAllowed | FALSE | Allow players to create a guild using the window in Underfoot+ |
| Integer | Guild | PlayerCreationLimit | 1 | Only allow use of the UF+ window if the account has &lt; than this number of guild leaders on it |
| Integer | Guild | PlayerCreationRequiredStatus | 0 | Required admin status. |
| Integer | Guild | PlayerCreationRequiredLevel | 0 | Required Level of the player attempting to create the guild. |
| Integer | Guild | PlayerCreationRequiredTime | 0 | Required Time Entitled On Account \(in Minutes\) to create the guild. |
| Integer | Skills | MaxTrainTradeskills | 21 |  |
| Boolean | Skills | UseLimitTradeskillSearchSkillDiff | TRUE |  |
| Integer | Skills | MaxTradeskillSearchSkillDiff | 50 |  |
| Integer | Skills | MaxTrainSpecializations | 50 | Max level a GM trainer will train casting specializations |
| Integer | Skills | SwimmingStartValue | 100 |  |
| Boolean | Skills | TrainSenseHeading | FALSE |  |
| Integer | Skills | SenseHeadingStartValue | 200 |  |
| Decimal | Pets | AttackCommandRange | 150 |  |
| Boolean | Pets | UnTargetableSwarmPet | FALSE |  |
| Decimal | Pets | PetPowerLevelCap | 10 | Max number of levels your pet can go up with pet power |
| Integer | GM | MinStatusToSummonItem | 250 |  |
| Integer | GM | MinStatusToZoneAnywhere | 250 |  |
| Integer | World | ZoneAutobootTimeoutMS | 60000 |  |
| Integer | World | ClientKeepaliveTimeoutMS | 65000 |  |
| Boolean | World | UseBannedIPsTable | FALSE | Toggle whether or not to check incoming client connections against the Banned\_IPs table. Set this value to false to disable this feature. |
| Boolean | World | EnableTutorialButton | TRUE |  |
| Boolean | World | EnableReturnHomeButton | TRUE |  |
| Integer | World | MaxLevelForTutorial | 10 |  |
| Integer | World | TutorialZoneID | 189 |  |
| Integer | World | GuildBankZoneID | 345 |  |
| Integer | World | MinOfflineTimeToReturnHome | 21600 | 21600 seconds is 6 Hours |
| Integer | World | MaxClientsPerIP | -1 | Maximum number of clients allowed to connect per IP address if account status is &lt; AddMaxClientsStatus. Default value: -1 \(feature disabled\) |
| Integer | World | ExemptMaxClientsStatus | -1 | \#ERROR! |
| Integer | World | AddMaxClientsPerIP | -1 | Maximum number of clients allowed to connect per IP address if account status is &lt; ExemptMaxClientsStatus. Default value: -1 \(feature disabled |
| Integer | World | AddMaxClientsStatus | -1 | \#ERROR! |
| Boolean | World | MaxClientsSetByStatus | FALSE | MaxClientsPerIP.  If True, IP Limiting will be set to the status on the account as long as the status is &gt; MaxClientsPerIP |
| Boolean | World | EnableIPExemptions | FALSE | If True, ip\_exemptions table is used, if there is no entry for the IP it will default to RuleI\(World:MaxClientsPerIP\) |
| Boolean | World | ClearTempMerchantlist | TRUE | Clears temp merchant items when world boots. |
| Boolean | World | DeleteStaleCorpeBackups | TRUE | Deletes stale corpse backups older than 2 weeks. |
| Integer | World | AccountSessionLimit | -1 | Max number of characters allowed on at once from a single account \(-1 is disabled |
| Integer | World | ExemptAccountLimitStatus | -1 | Min status required to be exempt from multi-session per account limiting \(-1 is disabled |
| Boolean | World | GMAccountIPList | FALSE | Check ip list against GM Accounts, AntiHack GM Accounts. |
| Integer | World | MinGMAntiHackStatus | 1 | Minimum GM status to check against AntiHack list |
| Integer | World | SoFStartZoneID | -1 | Sets the Starting Zone for SoF Clients separate from Titanium Clients \(-1 is disabled |
| Integer | World | TitaniumStartZoneID | -1 | Sets the Starting Zone for Titanium Clients \(-1 is disabled\). Replaces the old method. |
| Integer | World | ExpansionSettings | 16383 | Sets the expansion settings for the server, This is sent on login to world and affects client expansion settings. Defaults to all expansions enabled up to TSS. |
| Boolean | World | UseClientBasedExpansionSettings | TRUE | if true it will overrule World:ExpansionSettings and set someone's expansion based on the client they're using |
| Integer | World | PVPSettings | 0 | Sets the PVP settings for the server, 1 = Rallos Zek RuleSet, 2 = Tallon/Vallon Zek Ruleset, 4 = Sullon Zek Ruleset, 6 = Discord Ruleset, anything above 6 is the Discord Ruleset without the no-drop restrictions removed. TODO: Edit IsAttackAllowed in Zone to accomodate for these rules. |
| Boolean | World | IsGMPetitionWindowEnabled | FALSE |  |
| Integer | World | FVNoDropFlag | 0 | Sets the Firiona Vie settings on the client. If set to 2, the flag will be set for GMs only, allowing trading of no-drop items. |
| Boolean | World | IPLimitDisconnectAll | FALSE |  |
| Integer | World | TellQueueSize | 20 |  |
| Integer | Zone | NPCPositonUpdateTicCount | 32 | ms between intervals of sending a position update to the entire zone. |
| Integer | Zone | ClientLinkdeadMS | 180000 | the time a client remains link dead on the server after a sudden disconnection |
| Integer | Zone | GraveyardTimeMS | 1200000 | ms time until a player corpse is moved to a zone's graveyard, if one is specified for the zone |
| Boolean | Zone | EnableShadowrest | 1 | enables or disables the shadowrest zone feature for player corpses. Default is turned on. |
| Boolean | Zone | UsePlayerCorpseBackups | TRUE | Keeps backups of player corpses. |
| Integer | Zone | MQWarpExemptStatus | -1 | Required status level to exempt the MQWarpDetector. Set to -1 to disable this feature. |
| Integer | Zone | MQZoneExemptStatus | -1 | Required status level to exempt the MQZoneDetector. Set to -1 to disable this feature. |
| Integer | Zone | MQGateExemptStatus | -1 | Required status level to exempt the MQGateDetector. Set to -1 to disable this feature. |
| Integer | Zone | MQGhostExemptStatus | -1 | Required status level to exempt the MGhostDetector. Set to -1 to disable this feature. |
| Boolean | Zone | EnableMQWarpDetector | TRUE | Enable the MQWarp Detector. Set to False to disable this feature. |
| Boolean | Zone | EnableMQZoneDetector | TRUE | Enable the MQZone Detector. Set to False to disable this feature. |
| Boolean | Zone | EnableMQGateDetector | TRUE | Enable the MQGate Detector. Set to False to disable this feature. |
| Boolean | Zone | EnableMQGhostDetector | TRUE | Enable the MQGhost Detector. Set to False to disable this feature. |
| Decimal | Zone | MQWarpDetectionDistanceFactor | 9 | clients move at 4.4 about if in a straight line but with movement and to acct for lag we raise it a bit |
| Boolean | Zone | MarkMQWarpLT | FALSE |  |
| Integer | Zone | AutoShutdownDelay | 5000 | How long a dynamic zone stays loaded while empty |
| Integer | Zone | PEQZoneReuseTime | 900 | How long, in seconds, until you can reuse the \#peqzone command. |
| Integer | Zone | PEQZoneDebuff1 | 4454 | First debuff casted by \#peqzone Default is Cursed Keeper's Blight. |
| Integer | Zone | PEQZoneDebuff2 | 2209 | Second debuff casted by \#peqzone Default is Tendrils of Apathy. |
| Boolean | Zone | UsePEQZoneDebuffs | TRUE | Will determine if \#peqzone will debuff players or not when used. |
| Decimal | Zone | HotZoneBonus | 0.75 |  |
| Integer | Zone | ReservedInstances | 30 | Will reserve this many instance ids for globals... probably not a good idea to change this while a server is running. |
| Integer | Zone | EbonCrystalItemID | 40902 |  |
| Integer | Zone | RadiantCrystalItemID | 40903 |  |
| Boolean | Zone | LevelBasedEXPMods | FALSE | Allows you to use the level\_exp\_mods table in consideration to your players EXP hits |
| Integer | Zone | WeatherTimer | 600 | Weather timer when no duration is available |
| Boolean | Zone | EnableLoggedOffReplenishments | TRUE |  |
| Integer | Zone | MinOfflineTimeToReplenishments | 21600 | 21600 seconds is 6 Hours |
| Boolean | Zone | UseZoneController | TRUE | Enables the ability to use persistent quest based zone controllers \(zone\_controller.pl/lua |
| Boolean | Zone | EnableZoneControllerGlobals | FALSE | Enables the ability to use quest globals with the zone controller NPC |
| Boolean | Map | FixPathingZWhenLoading | TRUE | increases zone boot times a bit to reduce hopping. |
| Boolean | Map | FixPathingZAtWaypoints | FALSE | alternative to `WhenLoading`, accomplishes the same thing but does it at each waypoint instead of once at boot time. |
| Boolean | Map | FixPathingZWhenMoving | FALSE | very CPU intensive, but helps hopping with widely spaced waypoints. |
| Boolean | Map | FixPathingZOnSendTo | FALSE | try to repair Z coords in the SendTo routine as well. |
| Decimal | Map | FixPathingZMaxDeltaMoving | 20 | at runtime while pathing: max change in Z to allow the BestZ code to apply. |
| Decimal | Map | FixPathingZMaxDeltaWaypoint | 20 | at runtime at each waypoint: max change in Z to allow the BestZ code to apply. |
| Decimal | Map | FixPathingZMaxDeltaSendTo | 20 | at runtime in SendTo: max change in Z to allow the BestZ code to apply. |
| Decimal | Map | FixPathingZMaxDeltaLoading | 45 | while loading each waypoint: max change in Z to allow the BestZ code to apply. |
| Integer | Map | FindBestZHeightAdjust | 1 | Adds this to the current Z before seeking the best Z position |
| Boolean | Pathing | Aggro | TRUE | Enable pathing for aggroed mobs. |
| Boolean | Pathing | AggroReturnToGrid | TRUE | Enable pathing for aggroed roaming mobs returning to their previous waypoint. |
| Boolean | Pathing | Guard | TRUE | Enable pathing for mobs moving to their guard point. |
| Boolean | Pathing | Find | TRUE | Enable pathing for FindPerson requests from the client. |
| Boolean | Pathing | Fear | TRUE | Enable pathing for fear |
| Decimal | Pathing | ZDiffThreshold | 10 | If a mob las LOS to it's target, it will run to it if the Z difference is &lt; this. |
| Integer | Pathing | LOSCheckFrequency | 1000 | A mob will check for LOS to it's target this often \(milliseconds\). |
| Integer | Pathing | RouteUpdateFrequencyShort | 1000 | How often a new route will be calculated if the target has moved. |
| Integer | Pathing | RouteUpdateFrequencyLong | 5000 | How often a new route will be calculated if the target has moved. |
| Integer | Pathing | RouteUpdateFrequencyNodeCount | 5 |  |
| Decimal | Pathing | MinDistanceForLOSCheckShort | 40000 | \(NoRoot\). While following a path, only check for LOS to target within this distance. |
| Decimal | Pathing | MinDistanceForLOSCheckLong | 1000000 | \(NoRoot\). Min distance when initially attempting to acquire the target. |
| Integer | Pathing | MinNodesLeftForLOSCheck | 4 | Only check for LOS when we are down to this many path nodes left to run. |
| Integer | Pathing | MinNodesTraversedForLOSCheck | 3 | Only check for LOS after we have traversed this many path nodes. |
| Integer | Pathing | CullNodesFromStart | 1 | Checks LOS from Start point to second node for this many nodes and removes first node if there is LOS |
| Integer | Pathing | CullNodesFromEnd | 1 | Checks LOS from End point to second to last node for this many nodes and removes last node if there is LOS |
| Decimal | Pathing | CandidateNodeRangeXY | 400 | When searching for path start/end nodes, only nodes within this range will be considered. |
| Decimal | Pathing | CandidateNodeRangeZ | 10 | When searching for path start/end nodes, only nodes within this range will be considered. |
| Boolean | Watermap | CheckWaypointsInWaterWhenLoading | FALSE | Does not apply BestZ as waypoints are loaded if they are in water |
| Boolean | Watermap | CheckForWaterAtWaypoints | FALSE | Check if a mob has moved into/out of water when at waypoints and sets flymode |
| Boolean | Watermap | CheckForWaterWhenMoving | FALSE | Checks if a mob has moved into/out of water each time it's loc is recalculated |
| Boolean | Watermap | CheckForWaterOnSendTo | FALSE | Checks if a mob has moved into/out of water on SendTo |
| Boolean | Watermap | CheckForWaterWhenFishing | FALSE | Only lets a player fish near water \(if a water map exists for the zone |
| Decimal | Watermap | FishingRodLength | 30 | How far in front of player water must be for fishing to work |
| Decimal | Watermap | FishingLineLength | 100 | If water is more than this far below the player, it is considered too far to fish |
| Decimal | Watermap | FishingLineStepSize | 1 | Basic step size for fishing calc, too small and it will eat cpu, too large and it will miss potential water |
| Integer | Spells | AutoResistDiff | 15 |  |
| Decimal | Spells | ResistChance | 2 | chance to resist given no resists and same level |
| Decimal | Spells | ResistMod | 0.4 | multiplier, chance to resist = this \* ResistAmount |
| Decimal | Spells | PartialHitChance | 0.7 | The chance when a spell is resisted that it will partial hit. |
| Decimal | Spells | PartialHitChanceFear | 0.25 | The chance when a fear spell is resisted that it will partial hit. |
| Integer | Spells | BaseCritChance | 0 | base % chance that everyone has to crit a spell |
| Integer | Spells | BaseCritRatio | 100 | base % bonus to damage on a successful spell crit. 100 = 2x damage |
| Integer | Spells | WizCritLevel | 12 | level wizards first get spell crits |
| Integer | Spells | WizCritChance | 7 | wiz's crit chance, on top of BaseCritChance |
| Integer | Spells | WizCritRatio | 0 | wiz's crit bonus, on top of BaseCritRatio \(should be 0 for Live-like |
| Integer | Spells | ResistPerLevelDiff | 85 | 8.5 resist per level difference. |
| Integer | Spells | TranslocateTimeLimit | 0 | If not zero, time in seconds to accept a Translocate. |
| Integer | Spells | SacrificeMinLevel | 46 | first level Sacrifice will work on |
| Integer | Spells | SacrificeMaxLevel | 69 | last level Sacrifice will work on |
| Integer | Spells | SacrificeItemID | 9963 | Item ID of the item Sacrifice will return \(defaults to an EE |
| Boolean | Spells | EnableSpellGlobals | FALSE | If Enabled, spells check the spell\_globals table and compare character data from the quest globals before allowing that spell to scribe with scribespells |
| Integer | Spells | MaxBuffSlotsNPC | 25 |  |
| Integer | Spells | MaxSongSlotsNPC | 10 |  |
| Integer | Spells | MaxDiscSlotsNPC | 1 |  |
| Integer | Spells | MaxTotalSlotsNPC | 36 |  |
| Integer | Spells | MaxTotalSlotsPET | 30 | do not set this higher than 25 until the player profile is removed from the blob |
| Boolean | Spells | EnableBlockedBuffs | TRUE |  |
| Integer | Spells | ReflectType | 1 | 0 = disabled, 1 = single target player spells only, 2 = all player spells, 3 = all single target spells, 4 = all spells |
| Integer | Spells | VirusSpreadDistance | 30 | The distance a viral spell will jump to its next victim |
| Boolean | Spells | LiveLikeFocusEffects | TRUE | Determines whether specific healing, dmg and mana reduction focuses are randomized |
| Integer | Spells | BaseImmunityLevel | 55 | The level that targets start to be immune to stun, fear and mez spells with a max level of 0. |
| Boolean | Spells | NPCIgnoreBaseImmunity | TRUE | Whether or not NPCs get to ignore the BaseImmunityLevel for their spells. |
| Decimal | Spells | AvgSpellProcsPerMinute | 6 | Adjust rate for sympathetic spell procs |
| Integer | Spells | ResistFalloff | 67 | Max that level that will adjust our resist chance based on level modifiers |
| Integer | Spells | CharismaEffectiveness | 10 | Deterimes how much resist modification charisma applies to charm/pacify checks. Default 10 CHA = -1 resist mod. |
| Integer | Spells | CharismaEffectivenessCap | 255 | Deterimes how much resist modification charisma applies to charm/pacify checks. Default 10 CHA = -1 resist mod. |
| Boolean | Spells | CharismaCharmDuration | FALSE | Allow CHA resist mod to extend charm duration. |
| Integer | Spells | CharmBreakCheckChance | 25 | Determines chance for a charm break check to occur each buff tick. |
| Integer | Spells | MaxCastTimeReduction | 50 | Max percent your spell cast time can be reduced by spell haste |
| Integer | Spells | RootBreakFromSpells | 55 | Chance for root to break when cast on. |
| Integer | Spells | DeathSaveCharismaMod | 3 | Determines how much charisma effects chance of death save firing. |
| Integer | Spells | DivineInterventionHeal | 8000 | Divine intervention heal amount. |
| Integer | Spells | AdditiveBonusWornType | 0 | Calc worn bonuses to add together \(instead of taking highest\) if set to THIS worn type. \(2=Will covert live items automatically |
| Boolean | Spells | UseCHAScribeHack | FALSE | ScribeSpells and TrainDiscs quest functions will ignore entries where field 12 is CHA. What's the best way to do this? |
| Boolean | Spells | BuffLevelRestrictions | TRUE | Buffs will not land on low level toons like live |
| Integer | Spells | RootBreakCheckChance | 70 | Determines chance for a root break check to occur each buff tick. |
| Integer | Spells | FearBreakCheckChance | 70 | Determines chance for a fear break check to occur each buff tick. |
| Integer | Spells | SuccorFailChance | 2 | Determines chance for a succor spell not to teleport an invidual player |
| Integer | Spells | FRProjectileItem\_Titanium | 1113 | Item id for Titanium clients for Fire 'spell projectile'. |
| Integer | Spells | FRProjectileItem\_SOF | 80684 | Item id for SOF clients for Fire 'spell projectile'. |
| Integer | Spells | FRProjectileItem\_NPC | 80684 | Item id for NPC Fire 'spell projectile'. |
| Boolean | Spells | UseLiveSpellProjectileGFX | FALSE | Use spell projectile graphics set in the spells\_new table \(player\_1\). Server must be using UF+ spell file. |
| Boolean | Spells | FocusCombatProcs | FALSE | Allow all combat procs to receive focus effects. |
| Boolean | Spells | PreNerfBardAEDoT | FALSE | Allow bard AOE dots to damage targets when moving. |
| Integer | Spells | AI\_SpellCastFinishedFailRecast | 800 | AI spell recast time\(MS\) when an spell is cast but fails \(ie stunned\). |
| Integer | Spells | AI\_EngagedNoSpellMinRecast | 500 | AI spell recast time\(MS\) check when no spell is cast while engaged. \(min time in random |
| Integer | Spells | AI\_EngagedNoSpellMaxRecast | 1000 | AI spell recast time\(MS\) check when no spell is cast engaged.\(max time in random |
| Integer | Spells | AI\_EngagedBeneficialSelfChance | 100 | Chance during first AI Cast check to do a beneficial spell on self. |
| Integer | Spells | AI\_EngagedBeneficialOtherChance | 25 | Chance during second AI Cast check to do a beneficial spell on others. |
| Integer | Spells | AI\_EngagedDetrimentalChance | 20 | Chance during third AI Cast check to do a determental spell on others. |
| Integer | Spells | AI\_PursueNoSpellMinRecast | 500 | AI spell recast time\(MS\) check when no spell is cast while chasing target. \(min time in random |
| Integer | Spells | AI\_PursueNoSpellMaxRecast | 2000 | AI spell recast time\(MS\) check when no spell is cast while chasing target. \(max time in random |
| Integer | Spells | AI\_PursueDetrimentalChance | 90 | Chance while chasing target to cast a detrimental spell. |
| Integer | Spells | AI\_IdleNoSpellMinRecast | 6000 | AI spell recast time\(MS\) check when no spell is cast while idle. \(min time in random |
| Integer | Spells | AI\_IdleNoSpellMaxRecast | 60000 | AI spell recast time\(MS\) check when no spell is cast while chasing target. \(max time in random |
| Integer | Spells | AI\_IdleBeneficialChance | 100 | Chance while idle to do a beneficial spell on self or others. |
| Boolean | Spells | SHDProcIDOffByOne | TRUE | pre June 2009 SHD spell procs were off by 1, they stopped doing this in June 2009 \(so UF+ spell files need this false |
| Boolean | Spells | Jun182014HundredHandsRevamp | FALSE | this should be true for if you import a spell file newer than June 18, 2014 |
| Boolean | Spells | SwarmPetTargetLock | FALSE | Use old method of swarm pets target locking till target dies then despawning. |
| Boolean | Spells | NPC\_UseFocusFromSpells | TRUE | Allow npcs to use most spell derived focus effects. |
| Boolean | Spells | NPC\_UseFocusFromItems | FALSE | Allow npcs to use most item derived focus effects. |
| Boolean | Spells | UseAdditiveFocusFromWornSlot | FALSE | Allows an additive focus effect to be calculated from worn slot. |
| Boolean | Spells | AlwaysSendTargetsBuffs | FALSE | ignore LAA level if true |
| Boolean | Spells | FlatItemExtraSpellAmt | FALSE | allow SpellDmg stat to affect all spells, regardless of cast time/cooldown/etc |
| Boolean | Spells | IgnoreSpellDmgLvlRestriction | FALSE | ignore the 5 level spread on applying SpellDmg |
| Boolean | Spells | AllowItemTGB | FALSE | TGB doesn't work with items on live, custom servers want it though |
| Boolean | Spells | NPCInnateProcOverride | TRUE | NPC innate procs override the target type to single target. |
| Integer | Combat | MeleeBaseCritChance | 0 | The base crit chance for non warriors, NOTE: This will apply to NPCs as well |
| Integer | Combat | WarBerBaseCritChance | 3 | The base crit chance for warriors and berserkers, only applies to clients |
| Integer | Combat | BerserkBaseCritChance | 6 | The bonus base crit chance you get when you're berserk |
| Integer | Combat | NPCBashKickLevel | 6 | The level that npcs can KICK/BASH |
| Integer | Combat | NPCBashKickStunChance | 15 | Percent chance that a bash/kick will stun |
| Integer | Combat | RogueCritThrowingChance | 25 | Rogue throwing crit bonus |
| Integer | Combat | RogueDeadlyStrikeChance | 80 | Rogue chance throwing from behind crit becomes a deadly strike |
| Integer | Combat | RogueDeadlyStrikeMod | 2 | Deadly strike modifier to crit damage |
| Integer | Combat | ClientBaseCritChance | 0 | The base crit chance for all clients, this will stack with warrior's/zerker's crit chance. |
| Boolean | Combat | UseIntervalAC | TRUE |  |
| Integer | Combat | PetAttackMagicLevel | 30 |  |
| Boolean | Combat | EnableFearPathing | TRUE |  |
| Decimal | Combat | FleeMultiplier | 2 | Determines how quickly a NPC will slow down while fleeing. Decrease multiplier to slow NPC down quicker. |
| Integer | Combat | FleeHPRatio | 25 | HP % when a NPC begins to flee. |
| Boolean | Combat | FleeIfNotAlone | FALSE | If false, mobs won't flee if other mobs are in combat with it. |
| Boolean | Combat | AdjustProcPerMinute | TRUE |  |
| Decimal | Combat | AvgProcsPerMinute | 2 |  |
| Decimal | Combat | ProcPerMinDexContrib | 0.075 |  |
| Decimal | Combat | BaseProcChance | 0.035 |  |
| Decimal | Combat | ProcDexDivideBy | 11000 |  |
| Boolean | Combat | AdjustSpecialProcPerMinute | TRUE | Set PPM for special abilities like HeadShot \(Live does this as of 4-14 |
| Decimal | Combat | AvgSpecialProcsPerMinute | 2 | Unclear what best value is atm. |
| Decimal | Combat | BaseHitChance | 69 |  |
| Decimal | Combat | NPCBonusHitChance | 26 |  |
| Decimal | Combat | HitFalloffMinor | 5 | hit will fall off up to 5% over the initial level range |
| Decimal | Combat | HitFalloffModerate | 7 | hit will fall off up to 7% over the three levels after the initial level range |
| Decimal | Combat | HitFalloffMajor | 50 | hit will fall off sharply if we're outside the minor and moderate range |
| Decimal | Combat | HitBonusPerLevel | 1.2 | You gain this % of hit for every level you are above your target |
| Decimal | Combat | WeaponSkillFalloff | 0.33 | For every weapon skill point that's not maxed you lose this % of hit |
| Decimal | Combat | ArcheryHitPenalty | 0.25 | Archery has a hit penalty to try to help balance it with the plethora of long term +hit modifiers for it |
| Decimal | Combat | AgiHitFactor | 0.01 |  |
| Decimal | Combat | MinChancetoHit | 5 | Minimum % chance to hit with regular melee/ranged |
| Decimal | Combat | MaxChancetoHit | 95 | Maximum % chance to hit with regular melee/ranged |
| Integer | Combat | MinRangedAttackDist | 25 | Minimum Distance to use Ranged Attacks |
| Boolean | Combat | ArcheryBonusRequiresStationary | TRUE | does the 2x archery bonus chance require a stationary npc |
| Decimal | Combat | ArcheryBaseDamageBonus | 1 | % Modifier to Base Archery Damage \(.5 = 50% base damage, 1 = 100%, 2 = 200% |
| Decimal | Combat | ArcheryNPCMultiplier | 1 | this is multiplied by the regular dmg to get the archery dmg |
| Boolean | Combat | AssistNoTargetSelf | TRUE | when assisting a target that does not have a target: true = target self, false = leave target as was before assist \(false = live like |
| Integer | Combat | MaxRampageTargets | 3 | max number of people hit with rampage |
| Integer | Combat | DefaultRampageTargets | 1 | default number of people to hit with rampage |
| Boolean | Combat | RampageHitsTarget | FALSE | rampage will hit the target if it still has targets left |
| Integer | Combat | MaxFlurryHits | 2 | max number of extra hits from flurry |
| Integer | Combat | MonkDamageTableBonus | 5 | % bonus monks get to their damage table calcs |
| Integer | Combat | FlyingKickBonus | 25 | % Modifier that this skill gets to str and skill bonuses |
| Integer | Combat | DragonPunchBonus | 20 | % Modifier that this skill gets to str and skill bonuses |
| Integer | Combat | EagleStrikeBonus | 15 | % Modifier that this skill gets to str and skill bonuses |
| Integer | Combat | TigerClawBonus | 10 | % Modifier that this skill gets to str and skill bonuses |
| Integer | Combat | RoundKickBonus | 5 | % Modifier that this skill gets to str and skill bonuses |
| Integer | Combat | FrenzyBonus | 0 | % Modifier to damage |
| Integer | Combat | BackstabBonus | 0 | % Modifier to damage |
| Boolean | Combat | ProcTargetOnly | TRUE | true = procs will only affect our target, false = procs will affect all of our targets |
| Decimal | Combat | NPCACFactor | 2.25 |  |
| Integer | Combat | ClothACSoftcap | 75 |  |
| Integer | Combat | LeatherACSoftcap | 100 |  |
| Integer | Combat | MonkACSoftcap | 120 |  |
| Integer | Combat | ChainACSoftcap | 200 |  |
| Integer | Combat | PlateACSoftcap | 300 |  |
| Decimal | Combat | AAMitigationACFactor | 3 |  |
| Decimal | Combat | WarriorACSoftcapReturn | 0.45 |  |
| Decimal | Combat | KnightACSoftcapReturn | 0.33 |  |
| Decimal | Combat | LowPlateChainACSoftcapReturn | 0.23 |  |
| Decimal | Combat | LowChainLeatherACSoftcapReturn | 0.17 |  |
| Decimal | Combat | CasterACSoftcapReturn | 0.06 |  |
| Decimal | Combat | MiscACSoftcapReturn | 0.3 |  |
| Boolean | Combat | OldACSoftcapRules | FALSE | use old softcaps |
| Boolean | Combat | UseOldDamageIntervalRules | FALSE | use old damage formulas for everything |
| Decimal | Combat | WarACSoftcapReturn | 0.3448 | new AC returns |
| Decimal | Combat | ClrRngMnkBrdACSoftcapReturn | 0.303 |  |
| Decimal | Combat | PalShdACSoftcapReturn | 0.3226 |  |
| Decimal | Combat | DruNecWizEncMagACSoftcapReturn | 0.2 |  |
| Decimal | Combat | RogShmBstBerACSoftcapReturn | 0.25 |  |
| Decimal | Combat | SoftcapFactor | 1.88 |  |
| Decimal | Combat | ACthac0Factor | 0.55 |  |
| Decimal | Combat | ACthac20Factor | 0.55 |  |
| Integer | Combat | HitCapPre20 | 40 | live has it capped at 40 for whatever dumb reason... this is mainly for custom servers |
| Integer | Combat | HitCapPre10 | 20 | live has it capped at 20, see above :p |
| Integer | Combat | MinHastedDelay | 400 | how fast we can get with haste. |
| Decimal | Combat | AvgDefProcsPerMinute | 2 |  |
| Decimal | Combat | DefProcPerMinAgiContrib | 0.075 | How much agility contributes to defensive proc rate |
| Integer | Combat | SpecialAttackACBonus | 15 | Percent amount of damage per AC gained for certain special attacks \(damage = AC\*SpecialAttackACBonus/100\). |
| Integer | Combat | NPCFlurryChance | 20 | Chance for NPC to flurry. |
| Boolean | Combat | TauntOverLevel | 1 | Allows you to taunt NPC's over warriors level. |
| Decimal | Combat | TauntSkillFalloff | 0.33 | For every taunt skill point that's not maxed you lose this % chance to taunt. |
| Boolean | Combat | EXPFromDmgShield | FALSE | Determine if damage from a damage shield counts for EXP gain. |
| Integer | Combat | MonkACBonusWeight | 15 |  |
| Integer | Combat | ClientStunLevel | 55 | This is the level where client kicks and bashes can stun the target |
| Integer | Combat | QuiverHasteCap | 1000 | Quiver haste cap 1000 on live for a while, currently 700 on live |
| Boolean | Combat | UseArcheryBonusRoll | FALSE | Make the 51+ archery bonus require an actual roll |
| Integer | Combat | ArcheryBonusChance | 50 |  |
| Integer | Combat | BerserkerFrenzyStart | 35 |  |
| Integer | Combat | BerserkerFrenzyEnd | 45 |  |
| Boolean | Combat | OneProcPerWeapon | TRUE | If enabled, One proc per weapon per round |
| Boolean | Combat | ProjectileDmgOnImpact | TRUE | If enabled, projectiles \(ie arrows\) will hit on impact, instead of instantly. |
| Boolean | Combat | MeleePush | TRUE | enable melee push |
| Integer | Combat | MeleePushChance | 50 | \(NPCs\) chance the target will be pushed. Made up, 100 actually isn't that bad |
| Boolean | Combat | UseLiveCombatRounds | TRUE | turn this false if you don't want to worry about fixing up combat rounds for NPCs |
| Integer | Combat | NPCAssistCap | 5 | Maxiumium number of NPCs that will assist another NPC at once |
| Integer | Combat | NPCAssistCapTimer | 6000 | Time in milliseconds a NPC will take to clear assist aggro cap space |
| Boolean | Combat | UseRevampHandToHand | FALSE | use h2h revamped dmg/delays I believe this was implemented during SoF |
| Boolean | Combat | ClassicMasterWu | FALSE | classic master wu uses a random special, modern doesn't |
| Integer | NPC | MinorNPCCorpseDecayTimeMS | 450000 | level&lt;55 |
| Integer | NPC | MajorNPCCorpseDecayTimeMS | 1500000 | \#ERROR! |
| Integer | NPC | CorpseUnlockTimer | 150000 |  |
| Integer | NPC | EmptyNPCCorpseDecayTimeMS | 0 |  |
| Boolean | NPC | UseItemBonusesForNonPets | TRUE |  |
| Integer | NPC | SayPauseTimeInSec | 5 |  |
| Integer | NPC | OOCRegen | 0 |  |
| Boolean | NPC | BuffFriends | FALSE |  |
| Boolean | NPC | EnableNPCQuestJournal | FALSE |  |
| Integer | NPC | LastFightingDelayMovingMin | 10000 |  |
| Integer | NPC | LastFightingDelayMovingMax | 20000 |  |
| Boolean | NPC | SmartLastFightingDelayMoving | TRUE |  |
| Boolean | NPC | ReturnNonQuestNoDropItems | FALSE | Returns NO DROP items on NPCs that don't have an EVENT\_TRADE sub in their script |
| Integer | NPC | StartEnrageValue | 9 | % HP that an NPC will begin to enrage |
| Boolean | NPC | LiveLikeEnrage | FALSE | If set to true then only player controlled pets will enrage |
| Boolean | NPC | EnableMeritBasedFaction | FALSE | If set to true, faction will given in the same way as experience \(solo/group/raid |
| Integer | NPC | NPCToNPCAggroTimerMin | 500 |  |
| Integer | NPC | NPCToNPCAggroTimerMax | 6000 |  |
| Boolean | NPC | UseClassAsLastName | TRUE | Uses class archetype as LastName for npcs with none |
| Boolean | NPC | NewLevelScaling | TRUE | Better level scaling, use old if new formulas would break your server |
| Boolean | Aggro | SmartAggroList | TRUE |  |
| Integer | Aggro | SittingAggroMod | 35 | 35% |
| Integer | Aggro | MeleeRangeAggroMod | 10 | 10% |
| Integer | Aggro | CurrentTargetAggroMod | 0 | 0% -- will prefer our current target to any other; makes it harder for our npcs to switch targets. |
| Integer | Aggro | CriticallyWoundedAggroMod | 100 | 100% |
| Integer | Aggro | SpellAggroMod | 100 |  |
| Integer | Aggro | SongAggroMod | 33 |  |
| Integer | Aggro | PetSpellAggroMod | 10 |  |
| Decimal | Aggro | TunnelVisionAggroMod | 0.75 | people not currently the top hate generate this much hate on a Tunnel Vision mob |
| Integer | Aggro | MaxScalingProcAggro | 400 | Set to -1 for no limit. Maxmimum amount of aggro that HP scaling SPA effect in a proc will add. |
| Integer | Aggro | IntAggroThreshold | 75 | Int &lt;= this will aggro regardless of level difference. |
| Boolean | Aggro | AllowTickPulling | FALSE | tick pulling is an exploit in an NPC's call for help fixed sometime in 2006 on live |
| Boolean | Aggro | UseLevelAggro | TRUE | Level 18+ and Undead will aggro regardless of level difference. \(this will disabled Rule:IntAggroThreshold if set to true |
| Boolean | TaskSystem | EnableTaskSystem | TRUE | Globally enable or disable the Task system |
| Integer | TaskSystem | PeriodicCheckTimer | 5 | Seconds between checks for failed tasks. Also used by the 'Touch' activity |
| Boolean | TaskSystem | RecordCompletedTasks | TRUE |  |
| Boolean | TaskSystem | RecordCompletedOptionalActivities | FALSE |  |
| Boolean | TaskSystem | KeepOneRecordPerCompletedTask | TRUE |  |
| Boolean | TaskSystem | EnableTaskProximity | TRUE |  |
| Integer | Bots | AAExpansion | 8 | Bots get AAs through this expansion |
| Boolean | Bots | AllowCamelCaseNames | FALSE | Allows the use of 'MyBot' type names |
| Integer | Bots | CommandSpellRank | 1 | Filters bot command spells by rank \(1, 2 and 3 are valid filters - any other number allows all ranks |
| Integer | Bots | CreationLimit | 150 | Number of bots that each account can create |
| Boolean | Bots | FinishBuffing | FALSE | Allow for buffs to complete even if the bot caster is out of mana. Only affects buffing out of combat. |
| Boolean | Bots | GroupBuffing | FALSE | Bots will cast single target buffs as group buffs, default is false for single. Does not make single target buffs work for MGB. |
| Integer | Bots | HealRotationMaxMembers | 24 | Maximum number of heal rotation members |
| Integer | Bots | HealRotationMaxTargets | 12 | Maximum number of heal rotation targets |
| Decimal | Bots | ManaRegen | 2 | Adjust mana regen for bots, 1 is fast and higher numbers slow it down 3 is about the same as players. |
| Boolean | Bots | PreferNoManaCommandSpells | TRUE | Give sorting priority to newer no-mana spells \(i.e., 'Bind Affinity' |
| Boolean | Bots | QuestableSpawnLimit | FALSE | Optional quest method to manage bot spawn limits using the quest\_globals name bot\_spawn\_limit, see: /bazaar/Aediles\_Thrall.pl |
| Boolean | Bots | QuestableSpells | FALSE | Anita Thrall's \(Anita\_Thrall.pl\) Bot Spell Scriber quests. |
| Integer | Bots | SpawnLimit | 71 | Number of bots a character can have spawned at one time, You + 71 bots is a 12 group raid |
| Boolean | Bots | BotGroupXP | FALSE | Determines whether client gets xp for bots outside their group. |
| Boolean | Bots | BotBardUseOutOfCombatSongs | TRUE | Determines whether bard bots use additional out of combat songs \(optional script |
| Boolean | Bots | BotLevelsWithOwner | FALSE | Auto-updates spawned bots as owner levels/de-levels \(false is original behavior |
| Boolean | Bots | BotCharacterLevelEnabled | FALSE | Enables required level to spawn bots |
| Integer | Bots | BotCharacterLevel | 0 | this value you can spawn bots if BotCharacterLevelEnabled is true.  0 as default \(if level &gt; this value you can spawn bots if BotCharacterLevelEnabled is true |
| Boolean | Chat | ServerWideOOC | TRUE |  |
| Boolean | Chat | ServerWideAuction | TRUE |  |
| Boolean | Chat | EnableVoiceMacros | TRUE |  |
| Boolean | Chat | EnableMailKeyIPVerification | TRUE |  |
| Boolean | Chat | EnableAntiSpam | TRUE |  |
| Boolean | Chat | SuppressCommandErrors | FALSE | Do not suppress by default |
| Integer | Chat | MinStatusToBypassAntiSpam | 100 |  |
| Integer | Chat | MinimumMessagesPerInterval | 4 |  |
| Integer | Chat | MaximumMessagesPerInterval | 12 |  |
| Integer | Chat | MaxMessagesBeforeKick | 20 |  |
| Integer | Chat | IntervalDurationMS | 60000 |  |
| Integer | Chat | KarmaUpdateIntervalMS | 1200000 |  |
| Integer | Chat | KarmaGlobalChatLimit | 72 | amount of karma you need to be able to talk in ooc/auction/chat below the level limit |
| Integer | Chat | GlobalChatLevelLimit | 8 | level limit you need to of reached to talk in ooc/auction/chat if your karma is too low. |
| Boolean | Merchant | UsePriceMod | TRUE | Use faction/charisma price modifiers. |
| Decimal | Merchant | SellCostMod | 1.05 | Modifier for NPC sell price. |
| Decimal | Merchant | BuyCostMod | 0.95 | Modifier for NPC buy price. |
| Integer | Merchant | PriceBonusPct | 4 | Determines maximum price bonus from having good faction/CHA. Value is a percent. |
| Integer | Merchant | PricePenaltyPct | 4 | Determines maximum price penalty from having bad faction/CHA. Value is a percent. |
| Decimal | Merchant | ChaBonusMod | 3.45 | Determines CHA cap, from 104 CHA. 3.45 is 132 CHA at apprehensive. 0.34 is 400 CHA at apprehensive. |
| Decimal | Merchant | ChaPenaltyMod | 1.52 | Determines CHA bottom, up to 102 CHA. 1.52 is 37 CHA at apprehensive. 0.98 is 0 CHA at apprehensive. |
| Boolean | Merchant | EnableAltCurrencySell | TRUE | Enables the ability to resell items to alternate currency merchants |
| Boolean | Bazaar | AuditTrail | FALSE |  |
| Integer | Bazaar | MaxSearchResults | 50 |  |
| Boolean | Bazaar | EnableWarpToTrader | TRUE |  |
| Integer | Bazaar | MaxBarterSearchResults | 200 | The max results returned in the /barter search |
| Boolean | Mail | EnableMailSystem | TRUE | If false, client won't bring up the Mail window. |
| Integer | Mail | ExpireTrash | 0 | Time in seconds. 0 will delete all messages in the trash when the mailserver starts |
| Integer | Mail | ExpireRead | 31536000 | 1 Year. Set to -1 for never |
| Integer | Mail | ExpireUnread | 31536000 | 1 Year. Set to -1 for never |
| Integer | Channels | RequiredStatusAdmin | 251 | Required status to administer chat channels |
| Integer | Channels | RequiredStatusListAll | 251 | Required status to list all chat channels |
| Integer | Channels | DeleteTimer | 1440 | Empty password protected channels will be deleted after this many minutes |
| Boolean | EventLog | RecordSellToMerchant | FALSE | Record sales from a player to an NPC merchant in eventlog table |
| Boolean | EventLog | RecordBuyFromMerchant | FALSE | Record purchases by a player from an NPC merchant in eventlog table |
| Integer | Adventure | MinNumberForGroup | 2 |  |
| Integer | Adventure | MaxNumberForGroup | 6 |  |
| Integer | Adventure | MinNumberForRaid | 18 |  |
| Integer | Adventure | MaxNumberForRaid | 36 |  |
| Integer | Adventure | MaxLevelRange | 9 |  |
| Integer | Adventure | NumberKillsForBossSpawn | 45 |  |
| Decimal | Adventure | DistanceForRescueAccept | 10000 |  |
| Decimal | Adventure | DistanceForRescueComplete | 2500 |  |
| Integer | Adventure | ItemIDToEnablePorts | 41000 | 0 to disable, otherwise using a LDoN portal will require the user to have this item. |
| Integer | Adventure | LDoNTrapDistanceUse | 625 |  |
| Decimal | Adventure | LDoNBaseTrapDifficulty | 15 |  |
| Decimal | Adventure | LDoNCriticalFailTrapThreshold | 10 |  |
| Integer | Adventure | LDoNAdventureExpireTime | 1800 | 30 minutes to expire |
| Integer | AA | ExpPerPoint | 23976503 | Amount of exp per AA. Is the same as the amount of exp to go from level 51 to level 52. |
| Boolean | AA | Stacking | TRUE | Allow AA that belong to the same group to stack on SOF+ clients. |
| Integer | Console | SessionTimeOut | 600000 | Amount of time in ms for the console session to time out |
| Boolean | QueryServ | PlayerLogChat | FALSE | Logs Player Chat |
| Boolean | QueryServ | PlayerLogTrades | FALSE | Logs Player Trades |
| Boolean | QueryServ | PlayerLogHandins | FALSE | Logs Player Handins |
| Boolean | QueryServ | PlayerLogNPCKills | FALSE | Logs Player NPC Kills |
| Boolean | QueryServ | PlayerLogDeletes | FALSE | Logs Player Deletes |
| Boolean | QueryServ | PlayerLogMoves | FALSE | Logs Player Moves |
| Boolean | QueryServ | PlayerLogMerchantTransactions | FALSE | Logs Merchant Transactions |
| Boolean | QueryServ | PlayerLogPCCoordinates | FALSE | Logs Player Coordinates with certain events |
| Boolean | QueryServ | PlayerLogDropItem | FALSE | Logs Player Drop Item |
| Boolean | QueryServ | PlayerLogZone | FALSE | Logs Player Zone Events |
| Boolean | QueryServ | PlayerLogDeaths | FALSE | Logs Player Deaths |
| Boolean | QueryServ | PlayerLogConnectDisconnect | FALSE | Logs Player Connect Disconnect State |
| Boolean | QueryServ | PlayerLogLevels | FALSE | Logs Player Leveling/Deleveling |
| Boolean | QueryServ | PlayerLogAARate | FALSE | Logs Player AA Experience Rates |
| Boolean | QueryServ | PlayerLogQGlobalUpdate | FALSE | Logs Player QGlobal Updates |
| Boolean | QueryServ | PlayerLogTaskUpdates | FALSE | Logs Player Task Updates |
| Boolean | QueryServ | PlayerLogKeyringAddition | FALSE | Log PLayer Keyring additions |
| Boolean | QueryServ | PlayerLogAAPurchases | FALSE | Log Player AA Purchases |
| Boolean | QueryServ | PlayerLogTradeSkillEvents | FALSE | Log Player Tradeskill Transactions |
| Boolean | QueryServ | PlayerLogIssuedCommandes | FALSE | Log Player Issued Commands |
| Boolean | QueryServ | PlayerLogMoneyTransactions | FALSE | Log Player Money Transaction/Splits |
| Boolean | QueryServ | PlayerLogAlternateCurrencyTransactions | FALSE | Log Ploayer Alternate Currency Transactions |
| Boolean | Inventory | EnforceAugmentRestriction | TRUE | Forces augment slot restrictions |
| Boolean | Inventory | EnforceAugmentUsability | TRUE | Forces augmented item usability |
| Boolean | Inventory | EnforceAugmentWear | TRUE | Forces augment wear slot validation |
| Boolean | Inventory | DeleteTransformationMold | TRUE | False if you want mold to last forever |
| Boolean | Inventory | AllowAnyWeaponTransformation | FALSE | Weapons can use any weapon transformation |
| Boolean | Inventory | TransformSummonedBags | FALSE | Transforms summoned bags into disenchanted ones instead of deleting |
| Boolean | Client | UseLiveFactionMessage | FALSE | Allows players to see faction adjustments like Live |
| Boolean | Client | UseLiveBlockedMessage | FALSE | Allows players to see faction adjustments like Live |
| Boolean | Character | ProcessFearedProximity | FALSE | Process proximity checks while feared. |

