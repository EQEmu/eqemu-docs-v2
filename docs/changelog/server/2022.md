# 2022

## Februrary - October

**Akkadius** [Crash] Fix reload crashes 2022-09-29  
**Michael Cook (mackal)** eqlaunch wasn't loading paths ([#2461](https://github.com/EQEmu/Server/pull/2461)) 2022-09-29  
**Akkadius** [Hotfix] Fix lua mod load path 2022-09-29  
**Akkadius** [Logging] Remove loginserver unhandled error ([#2458](https://github.com/EQEmu/Server/pull/2458)) 2022-09-28  
**Kinglykrab** [Tasks] Use zone currencies instead of hard-coded enum. ([#2459](https://github.com/EQEmu/Server/pull/2459)) 2022-09-28  
**Akkadius** [Crash] Websocket Crash fix race when fetching log categories ([#2456](https://github.com/EQEmu/Server/pull/2456)) 2022-09-28  
**Aeadoin** [Feature] Change Lifetap Emotes to be filterable. ([#2454](https://github.com/EQEmu/Server/pull/2454)) 2022-09-28  
**Akkadius** [Crash] Fix spawn race condition shown by #repop ([#2455](https://github.com/EQEmu/Server/pull/2455)) 2022-09-28  
**Akkadius** [Database] Add fallback migration for logsys columns ([#2457](https://github.com/EQEmu/Server/pull/2457)) 2022-09-28  
**Akkadius** [Hotfix] Force collation on conversion script 2022-09-28  
**Akkadius** [Logging] Add stack trace in code paths that shouldn't occur ([#2453](https://github.com/EQEmu/Server/pull/2453)) 2022-09-28  
**Akkadius** [File Paths] Implement Path Manager ([#2440](https://github.com/EQEmu/Server/pull/2440)) 2022-09-28  
**Akkadius** [Logging] Netcode Logging Unify ([#2443](https://github.com/EQEmu/Server/pull/2443)) 2022-09-28  
**Akkadius** [World CLI] Refactor world CLI to be easier to reason about ([#2441](https://github.com/EQEmu/Server/pull/2441)) 2022-09-28  
**Akkadius** [Library] Bump httplib to 0.11.2 ([#2442](https://github.com/EQEmu/Server/pull/2442)) 2022-09-28  
**Kinglykrab** [Quest API] Add GetGMStatus() to Perl/Lua. ([#2448](https://github.com/EQEmu/Server/pull/2448)) 2022-09-28  
**Kinglykrab** [Quest API] Add Merchant Events to Perl/Lua. ([#2452](https://github.com/EQEmu/Server/pull/2452)) 2022-09-28  
**hg** [Tasks] Schema simplification ([#2449](https://github.com/EQEmu/Server/pull/2449)) 2022-09-28  
**JJ** [Database Updates] Typo in manifest 9207 ([#2451](https://github.com/EQEmu/Server/pull/2451)) 2022-09-27  
**Aeadoin** [Bug Fix] Fix Swarm Pet Flurry/Rampages Messages ([#2444](https://github.com/EQEmu/Server/pull/2444)) 2022-09-25  
**Akkadius** [eqemu_server.pl] Remove non-working fetch_latest_windows_binaries() ([#2445](https://github.com/EQEmu/Server/pull/2445)) 2022-09-25  
**Akkadius** [Command] Fix #copycharacter command crash ([#2446](https://github.com/EQEmu/Server/pull/2446)) 2022-09-25  
**Aeadoin** [Feature] Allow Focus Effects to be Filtered out. ([#2447](https://github.com/EQEmu/Server/pull/2447)) 2022-09-25  
**Michael Cook (mackal)** Define is _WINDOWS not WINDOWS ([#2439](https://github.com/EQEmu/Server/pull/2439)) 2022-09-21  
**hg** [Quest API] Let HasQuestSub check encounters ([#2435](https://github.com/EQEmu/Server/pull/2435)) 2022-09-20  
**Aeadoin** [Feature] Add Type 49545 to Spell Resistrictions ([#2436](https://github.com/EQEmu/Server/pull/2436)) 2022-09-20  
**Aeadoin** [Bug Fix] Fix Spellinfo Command to work with SpellIDs above int16 ([#2437](https://github.com/EQEmu/Server/pull/2437)) 2022-09-20  
**hg** [Tasks] Let task reward find free bag slots ([#2431](https://github.com/EQEmu/Server/pull/2431)) 2022-09-18  
**Akkadius** [Pathing] Fix pathing z-correctness for certain models ([#2430](https://github.com/EQEmu/Server/pull/2430)) 2022-09-11  
**Aeadoin** [Feature] Update HateMod used by SPA 114 to Int32\. ([#2428](https://github.com/EQEmu/Server/pull/2428)) 2022-09-08  
**hg** [Tasks] Add rule to update multiple task elements ([#2427](https://github.com/EQEmu/Server/pull/2427)) 2022-09-06  
**hg** (github-desktop-hgtw/master) [Commands] Make #damage require a target ([#2426](https://github.com/EQEmu/Server/pull/2426)) 2022-09-05  
**Michael Cook (mackal)** I guess we'll go plural ([#2425](https://github.com/EQEmu/Server/pull/2425)) 2022-09-05  
**Akkadius** [Hotfix] Add Bazaar portal discs to SQL 2022-09-05  
**Akkadius** [Zone Points] Fix zone point heading data ([#2415](https://github.com/EQEmu/Server/pull/2415)) 2022-09-05  
**Akkadius** [Doors] Fix door target zone heading data ([#2414](https://github.com/EQEmu/Server/pull/2414)) 2022-09-05  
**Akkadius** [Character Starting Points] Fix headings data ([#2413](https://github.com/EQEmu/Server/pull/2413)) 2022-09-05  
**Akkadius** [Hotfix] Faction associations file naming / lock consistency 2022-09-05  
**Michael Cook (mackal)** Clang was complaining about these ([#2421](https://github.com/EQEmu/Server/pull/2421)) 2022-09-03  
**hg** [Tasks] Only allow shared task completion once ([#2422](https://github.com/EQEmu/Server/pull/2422)) 2022-09-03  
**hg** [Tasks] Make Task Selector Cooldown Optional ([#2420](https://github.com/EQEmu/Server/pull/2420)) 2022-09-03  
**Michael Cook (mackal)** [Bug Fix] Shared Memory Faction Association Typo ([#2419](https://github.com/EQEmu/Server/pull/2419)) 2022-09-03  
**Michael Cook (mackal)** [Cleanup] Rework Lua QuestReward to not use try/catch blocks ([#2417](https://github.com/EQEmu/Server/pull/2417)) 2022-09-03  
**Kinglykrab** [NPC Scaling] Recalculate Skills and Reload Spells on Level Change ([#2416](https://github.com/EQEmu/Server/pull/2416)) 2022-09-03  
**Michael Cook (mackal)** Unsure how this seding messed up ([#2418](https://github.com/EQEmu/Server/pull/2418)) 2022-09-03  
**Michael Cook (mackal)** [Feature] Faction Association ([#2408](https://github.com/EQEmu/Server/pull/2408)) 2022-09-03  
**hg** [Tasks] Change zone task data container ([#2410](https://github.com/EQEmu/Server/pull/2410)) 2022-09-03  
**Akkadius** (akkadius/heading-fix-doors-data) [Zoning] Fix zoning logic issues ([#2412](https://github.com/EQEmu/Server/pull/2412)) 2022-09-03  
**hg** [Bug Fix] Avoid erase in discord queue range loop ([#2411](https://github.com/EQEmu/Server/pull/2411)) 2022-09-02  
**hg** [Bug Fix] Fix memory leak in ucs ([#2409](https://github.com/EQEmu/Server/pull/2409)) 2022-09-02  
**hg** [Tasks] Tweak task update messages ([#2406](https://github.com/EQEmu/Server/pull/2406)) 2022-09-02  
**hg** [Tasks] Remove delivered task items from trades ([#2405](https://github.com/EQEmu/Server/pull/2405)) 2022-09-02  
**hg** [Zone] Add missing safe_heading assignment ([#2407](https://github.com/EQEmu/Server/pull/2407)) 2022-09-02  
**hg** [Tasks] Replace task goals with explicit fields ([#2402](https://github.com/EQEmu/Server/pull/2402)) 2022-09-01  
**Michael Cook (mackal)** [Bug Fix] Resolve logic error in Raid::QueueClients ([#2404](https://github.com/EQEmu/Server/pull/2404)) 2022-09-01  
**Akkadius** [Code Cleanup] Zone Data Loading Refactor ([#2388](https://github.com/EQEmu/Server/pull/2388)) 2022-09-01  
**Michael Cook (mackal)** Fix issues with Client::SetHideMe ([#2403](https://github.com/EQEmu/Server/pull/2403)) 2022-09-01  
**Michael Cook (mackal)** [Repositories] Add more precise types to repository generator ([#2391](https://github.com/EQEmu/Server/pull/2391)) 2022-08-31  
**Kinglykrab** [Commands] Add #findrecipe and #viewrecipe Commands. ([#2401](https://github.com/EQEmu/Server/pull/2401)) 2022-08-31  
**Michael** [Bug] Loot Drop Randomization adjustment ([#2368](https://github.com/EQEmu/Server/pull/2368)) 2022-08-31  
**Aeadoin** [Feature] Implement Heroic Strikethrough to NPCs ([#2395](https://github.com/EQEmu/Server/pull/2395)) 2022-08-31  
**Michael Cook (mackal)** [Utility] Add std::string_view overloads for std::from_chars ([#2392](https://github.com/EQEmu/Server/pull/2392)) 2022-08-31  
**Akkadius** [Zone] Fix and simplify zone shutdown logic ([#2390](https://github.com/EQEmu/Server/pull/2390)) 2022-08-30  
**Michael Cook (mackal)** Use macro to generate correct format specifier ([#2400](https://github.com/EQEmu/Server/pull/2400)) 2022-08-28  
**hg** [Bug Fix] Fix loading world shared task state ([#2398](https://github.com/EQEmu/Server/pull/2398)) 2022-08-28  
**Kinglykrab** [Quest API] Allow CreateInstance to be used without a Client initiator. ([#2399](https://github.com/EQEmu/Server/pull/2399)) 2022-08-28  
**Kinglykrab** [Quest API] Add Recipe Methods ([#2393](https://github.com/EQEmu/Server/pull/2393)) 2022-08-23  
**Michael Cook (mackal)** [Manifest] Its not_empty not notempty ([#2394](https://github.com/EQEmu/Server/pull/2394)) 2022-08-22  
**Kinglykrab** [Bug Fix] Fix Silent Saylinks Sending Message to Others. ([#2389](https://github.com/EQEmu/Server/pull/2389)) 2022-08-21  
**Akkadius** [Expansions] Zone expansion consistency changes ([#2380](https://github.com/EQEmu/Server/pull/2380)) 2022-08-21  
**Kinglykrab** [Feature] Add Guild Chat to Console. ([#2387](https://github.com/EQEmu/Server/pull/2387)) 2022-08-21  
**hg** [Tasks] Implement task activity prerequisites ([#2374](https://github.com/EQEmu/Server/pull/2374)) 2022-08-21  
**Kinglykrab** [Bug Fix] Fix Duplicate Silent Saylink Messages ([#2386](https://github.com/EQEmu/Server/pull/2386)) 2022-08-21  
**Akkadius** [Repository] Modernize character recipe list ([#2385](https://github.com/EQEmu/Server/pull/2385)) 2022-08-21  
**Akkadius** [Tasks] Data validation for zone_version ([#2381](https://github.com/EQEmu/Server/pull/2381)) 2022-08-21  
**Aeadoin** [Feature] Change Mana Costs to use Signed Int ([#2384](https://github.com/EQEmu/Server/pull/2384)) 2022-08-21  
**Kinglykrab** [Bug Fix] Fix Strings::Money Missing Conditions. ([#2383](https://github.com/EQEmu/Server/pull/2383)) 2022-08-21  
**Michael Cook (mackal)** Update CURRENT_BINARY_DATABASE_VERSION ([#2382](https://github.com/EQEmu/Server/pull/2382)) 2022-08-20  
**Kinglykrab** [Quest API] Add Goto Player Teleport Methods. ([#2379](https://github.com/EQEmu/Server/pull/2379)) 2022-08-20  
**Kinglykrab** [Commands] Command Status Reload and Helper Method ([#2377](https://github.com/EQEmu/Server/pull/2377)) 2022-08-20  
**Kinglykrab** [Feature] Instance Version Specific Experience Modifiers ([#2376](https://github.com/EQEmu/Server/pull/2376)) 2022-08-20  
**Akkadius** [Loot] Add #lootsim (Loot Simulator) command ([#2375](https://github.com/EQEmu/Server/pull/2375)) 2022-08-20  
**Kinglykrab** [Bug Fix] Fix Bot Group Loading ([#2366](https://github.com/EQEmu/Server/pull/2366)) 2022-08-20  
**Michael Cook (mackal)** Fix issue with trap auras casting on caster ([#2378](https://github.com/EQEmu/Server/pull/2378)) 2022-08-18  
**Kinglykrab** [Saylinks] Convert all GM Command Saylinks to Silent Saylinks. ([#2373](https://github.com/EQEmu/Server/pull/2373)) 2022-08-13  
**Akkadius** [Saylinks] Add Silent helper ([#2372](https://github.com/EQEmu/Server/pull/2372)) 2022-08-13  
**Michael** [Bug] UINT32 EmoteID ([#2369](https://github.com/EQEmu/Server/pull/2369)) 2022-08-13  
**Akkadius** [Repositories] Add GetMaxId, Count ([#2371](https://github.com/EQEmu/Server/pull/2371)) 2022-08-13  
**Akkadius** [Doors] Improvements to door manipulation ([#2370](https://github.com/EQEmu/Server/pull/2370)) 2022-08-13  
**Kinglykrab** [Quest API] Add GetSkillDmgAmt() to Perl. ([#2365](https://github.com/EQEmu/Server/pull/2365)) 2022-08-10  
**Aeadoin** [Feature] Change GetSkillDmgAmt to int32 ([#2364](https://github.com/EQEmu/Server/pull/2364)) 2022-08-09  
**Kinglykrab** [Rules] Add Rule to Enable Tells with #hideme ([#2358](https://github.com/EQEmu/Server/pull/2358)) 2022-08-04  
**Akkadius** [Hotfix] Remove expansion field from account for those who have it ([#2357](https://github.com/EQEmu/Server/pull/2357)) 2022-08-01  
**hg** [Tasks] Add task accept packet validation ([#2354](https://github.com/EQEmu/Server/pull/2354)) 2022-07-31  
**hg** [Quest API] Fix lua task selector count when over max ([#2353](https://github.com/EQEmu/Server/pull/2353)) 2022-07-31  
**Kinglykrab** [Typo] Remove CanTradeFVNoDropItem() Duplicate ([#2352](https://github.com/EQEmu/Server/pull/2352)) 2022-07-31  
**hg** [Tasks] Make #task reloadall not quit shared tasks ([#2351](https://github.com/EQEmu/Server/pull/2351)) 2022-07-31  
**Kinglykrab** [Quest API] Add GetBotItem() and GetBotItemIDBySlot() to Perl/Lua. ([#2350](https://github.com/EQEmu/Server/pull/2350)) 2022-07-31  
**Michael Cook (mackal)** [Code Cleanup] Resolve some warnings in loginserver/world_server.cpp ([#2347](https://github.com/EQEmu/Server/pull/2347)) 2022-07-31  
**hg** [Shared Tasks] Fix shared task message target ([#2349](https://github.com/EQEmu/Server/pull/2349)) 2022-07-31  
**hg** [Shared Tasks] Avoid erasing shared tasks while iterating ([#2348](https://github.com/EQEmu/Server/pull/2348)) 2022-07-31  
**Akkadius** [Hotfix] Fix potential race for crash dumps (Linux) 2022-07-31  
**Akkadius** [Hotfix] SQL Update 2022-07-31  
**Michael Cook (mackal)** [Bug Fix] Limit merchant temp item list to zone and instance ([#2346](https://github.com/EQEmu/Server/pull/2346)) 2022-07-30  
**hg** [Dynamic Zones] Implement dz templates ([#2345](https://github.com/EQEmu/Server/pull/2345)) 2022-07-30  
**hg** [Tasks] Use dz switch id for task touch events ([#2344](https://github.com/EQEmu/Server/pull/2344)) 2022-07-30  
**hg** [Dynamic Zones] Implement dz switch id ([#2343](https://github.com/EQEmu/Server/pull/2343)) 2022-07-30  
**hg** [Shared Tasks] Enforce task reqs on player removal ([#2342](https://github.com/EQEmu/Server/pull/2342)) 2022-07-30  
**hg** [Shared Tasks] Cleanup shared task request and remove ([#2341](https://github.com/EQEmu/Server/pull/2341)) 2022-07-30  
**hg** [Shared Tasks] Implement task timer groups ([#2340](https://github.com/EQEmu/Server/pull/2340)) 2022-07-30  
**hg** [Shared Tasks] Implement Activity Locking ([#2339](https://github.com/EQEmu/Server/pull/2339)) 2022-07-30  
**Michael** [Feature] GM State Change Persistance ([#2328](https://github.com/EQEmu/Server/pull/2328)) 2022-07-30  
**hg** [Quest API] Use Floating Point for CameraEffect Intensity ([#2337](https://github.com/EQEmu/Server/pull/2337)) 2022-07-30  
**hg** [Tasks] Send Client Message for All Solo Task Updates ([#2336](https://github.com/EQEmu/Server/pull/2336)) 2022-07-30  
**hg** [Saylinks] Inject Saylinks in MessageClose API ([#2335](https://github.com/EQEmu/Server/pull/2335)) 2022-07-30  
**Kinglykrab** [Quest API] Add IsRareSpawn() to Perl/Lua. ([#2338](https://github.com/EQEmu/Server/pull/2338)) 2022-07-30  
**Paul Coene** [Roambox] Improve Path Finding ([#2324](https://github.com/EQEmu/Server/pull/2324)) 2022-07-30  
**Michael Cook (mackal)** [Feature] NPCs with bows and arrows do ranged attacks ([#2322](https://github.com/EQEmu/Server/pull/2322)) 2022-07-30  
**Michael** [Rules] Add adjustment for zone forage. ([#2330](https://github.com/EQEmu/Server/pull/2330)) 2022-07-30  
**Aeadoin** [Feature] Change mana_used to int32 ([#2321](https://github.com/EQEmu/Server/pull/2321)) 2022-07-30  
**hg** [Tasks] Support Raw NPC Names in Task Goal List ([#2333](https://github.com/EQEmu/Server/pull/2333)) 2022-07-30  
**hg** [Tasks] Use CashReward for Tasks ([#2332](https://github.com/EQEmu/Server/pull/2332)) 2022-07-30  
**Trent** [Rules] Add Keep Level on Death ([#2319](https://github.com/EQEmu/Server/pull/2319)) 2022-07-30  
**hg** [Tasks] Add Task Reward Points Field ([#2317](https://github.com/EQEmu/Server/pull/2317)) 2022-07-30  
**Quintinon** [Rules] Update logic checks everywhere for FVNoDropFlag. ([#2179](https://github.com/EQEmu/Server/pull/2179)) 2022-07-30  
**Kinglykrab** [Rules] Add Rule to allow Assassinate on non-Humanoid body types. ([#2331](https://github.com/EQEmu/Server/pull/2331)) 2022-07-29  
**Kinglykrab** [Rules] Add Rule to allow Headshots on non-Humanoid body types. ([#2329](https://github.com/EQEmu/Server/pull/2329)) 2022-07-29  
**Akkadius** [Hotfix] Shared Memory Protection Fixes 2022-07-27  
**Michael Cook (mackal)** [CPP] Update C++ standard to C++17 ([#2308](https://github.com/EQEmu/Server/pull/2308)) 2022-07-27  
**Quintinon** [Memory Leak] This mem leak was missed due to merge issues in the previous PRs. ([#2314](https://github.com/EQEmu/Server/pull/2314)) 2022-07-27  
**hg** [Quest API] Allow scripts to prevent door click ([#2327](https://github.com/EQEmu/Server/pull/2327)) 2022-07-27  
**hg** [Netcode] Adjust first packet for compress flag ([#2326](https://github.com/EQEmu/Server/pull/2326)) 2022-07-27  
**Akkadius** [Client] Fix IsMoving for Client ([#2318](https://github.com/EQEmu/Server/pull/2318)) 2022-07-27  
**hg** [Saylinks] Refactor saylink injection ([#2315](https://github.com/EQEmu/Server/pull/2315)) 2022-07-27  
**Akkadius** [CI] Hook tests back up ([#2316](https://github.com/EQEmu/Server/pull/2316)) 2022-07-27  
**Akkadius** [Content Filter] Fix Runtime Filtering When Set to -1 (All) ([#2313](https://github.com/EQEmu/Server/pull/2313)) 2022-07-27  
**Paul Coene** Fix bestz to work on client or target. ([#2323](https://github.com/EQEmu/Server/pull/2323)) 2022-07-23  
**hg** [Tasks] Reward clients on shared task completion sync ([#2306](https://github.com/EQEmu/Server/pull/2306)) 2022-07-16  
**Akkadius** [Hotfix] fix manifest 2022-07-16  
**Michael Cook (mackal)** [Feature] Implement OP_CashReward ([#2307](https://github.com/EQEmu/Server/pull/2307)) 2022-07-14  
**KimLS** Fix windows build for strings.cpp 2022-07-14  
**Michael Cook (mackal)** [Bug Fix] Remove StringUtilTest::EscapeStringMemoryTest ([#2310](https://github.com/EQEmu/Server/pull/2310)) 2022-07-14  
**Akkadius** [Tasks] Zone Version Matching ([#2303](https://github.com/EQEmu/Server/pull/2303)) 2022-07-14  
**Akkadius** [Zone] Deprecate Zone `expansion` Field ([#2297](https://github.com/EQEmu/Server/pull/2297)) 2022-07-14  
**Akkadius** [World] World Bootup Consolidation ([#2294](https://github.com/EQEmu/Server/pull/2294)) 2022-07-14  
**Akkadius** [Hotfix] Fix merge issue 2022-07-14  
**Paul Coene** [DoT Messages] Add DoT messages for mob->PC casts, fixed others to use correct str. ([#2289](https://github.com/EQEmu/Server/pull/2289)) 2022-07-14  
**neckkola** [Login] Added OP_ExpansionPacketData for RoF2 and update payload for Titanium ([#2186](https://github.com/EQEmu/Server/pull/2186)) 2022-07-14  
**hg** [Tasks] Apply full duration mission replay timers ([#2299](https://github.com/EQEmu/Server/pull/2299)) 2022-07-14  
**Akkadius** [Logs] #logs list Improvements ([#2302](https://github.com/EQEmu/Server/pull/2302)) 2022-07-14  
**Akkadius** [Code Cleanup] Remove Unused EQEMU_DEPOP_INVALIDATES_CACHE ([#2292](https://github.com/EQEmu/Server/pull/2292)) 2022-07-14  
**Akkadius** [World] Add more descriptive LS auth erroring ([#2293](https://github.com/EQEmu/Server/pull/2293)) 2022-07-14  
**Akkadius** [Crash] Linux Crash Dump Improvements ([#2296](https://github.com/EQEmu/Server/pull/2296)) 2022-07-14  
**hg** [Tasks] Place task item rewards in free slots ([#2300](https://github.com/EQEmu/Server/pull/2300)) 2022-07-14  
**hg** [Tasks] Fix #task command crash on bad input ([#2301](https://github.com/EQEmu/Server/pull/2301)) 2022-07-14  
**Akkadius** [Code Cleanup] Remove use of bzero since it is deprecated for memset ([#2295](https://github.com/EQEmu/Server/pull/2295)) 2022-07-14  
**Akkadius** [Logs] Fix GMSay Log Regression ([#2298](https://github.com/EQEmu/Server/pull/2298)) 2022-07-14  
**Kinglykrab** [Validation] Add Size Validation to #hotfix. ([#2304](https://github.com/EQEmu/Server/pull/2304)) 2022-07-14  
**Akkadius** [Strings] Refactor Strings Usage ([#2305](https://github.com/EQEmu/Server/pull/2305)) 2022-07-14  
**hg** [Quest API] Fix missing arg in perl set_proximity ([#2291](https://github.com/EQEmu/Server/pull/2291)) 2022-07-08  
**Michael Cook (mackal)** Fix string comparison issues in Client::SendZoneFlagInfo ([#2290](https://github.com/EQEmu/Server/pull/2290)) 2022-07-08  
**Akkadius** [Hotfix] Windows compile fix take 3 (final) 2022-07-06  
**Akkadius** [Hotfix] Possible windows compile fix take 2 2022-07-06  
**Akkadius** [Hotfix] Possible windows compile fix 2022-07-06  
**Akkadius** [Server Maintenance Script] Improvements to Downloading - Empty File Detection ([#2282](https://github.com/EQEmu/Server/pull/2282)) 2022-07-06  
**Akkadius** [Backups] Use World CLI for Database Backups ([#2286](https://github.com/EQEmu/Server/pull/2286)) 2022-07-06  
**Akkadius** [Server] Configuration Issues Checker (LAN Detection) ([#2283](https://github.com/EQEmu/Server/pull/2283)) 2022-07-06  
**Akkadius** [Logging] Table Injection - Member Variable Cleanup ([#2281](https://github.com/EQEmu/Server/pull/2281)) 2022-07-06  
**Paul Coene** [Bug Fix] NPCs were getting weapon proc added twice ([#2277](https://github.com/EQEmu/Server/pull/2277)) 2022-07-06  
**Kinglykrab** [Quest API] Perl Doors Fix. ([#2288](https://github.com/EQEmu/Server/pull/2288)) 2022-07-05  
**Kinglykrab** [Quest API] Add missing methods/package.adds to Perl API. ([#2287](https://github.com/EQEmu/Server/pull/2287)) 2022-07-05  
**Kinglykrab** [Bug Fix] Fix Spell Bucket and Spell Global Logic Checks. ([#2285](https://github.com/EQEmu/Server/pull/2285)) 2022-07-04  
**hg** [Quest API] Use binding library for perl apis ([#2216](https://github.com/EQEmu/Server/pull/2216)) 2022-07-03  
**Michael** [Cleanup] Update to EQEmu [#2253](https://github.com/EQEmu/Server/pull/2253) to clean up message strings ([#2279](https://github.com/EQEmu/Server/pull/2279)) 2022-07-03  
**Akkadius** [Hotfix] Move discord_webhooks to state tables because we don't want webhooks being exported 2022-07-03  
**Akkadius** [Hotfix] Add discord_webhooks to server tables 2022-07-03  
**Quintinon** [Bug Fix] Fix miscellaneous memory leaks related to EQApplicationPacket and it's pBuffer ([#2262](https://github.com/EQEmu/Server/pull/2262)) 2022-07-02  
**Paul Coene** [Feature] Add humanoid and non-wielded restrictions to pick pocket ([#2276](https://github.com/EQEmu/Server/pull/2276)) 2022-07-02  
**Quintinon** [Bug Fix] Handle memory leaks from return value of Client::GetTraderItems() ([#2266](https://github.com/EQEmu/Server/pull/2266)) 2022-07-02  
**Quintinon** [Cleanup] Fix unintended copies in zone/zonedb.cpp by changing auto to auto& ([#2271](https://github.com/EQEmu/Server/pull/2271)) 2022-07-02  
**Michael** [Feature] Bind Wound and Forage while mounted. ([#2257](https://github.com/EQEmu/Server/pull/2257)) 2022-07-02  
**Michael** [Bug Fix] Add required distance to CoTH before aggro wipe ([#2253](https://github.com/EQEmu/Server/pull/2253)) 2022-07-02  
**Kinglykrab** [Telnet] Add guildsay to console commands and Guild Channel to QueueMessage. ([#2263](https://github.com/EQEmu/Server/pull/2263)) 2022-07-02  
**Quintinon** [Cleanup] Cleanup code smells and compiler warnings in common/shareddb ([#2270](https://github.com/EQEmu/Server/pull/2270)) 2022-07-02  
**Kinglykrab** [Bots] Cleanup ^inventoryremove, ^inventorylist, and ^list Commands and bot groups. ([#2273](https://github.com/EQEmu/Server/pull/2273)) 2022-07-02  
**Paul Coene** [Fix] Boats should never get FixZ'd ([#2246](https://github.com/EQEmu/Server/pull/2246)) 2022-07-02  
**Kinglykrab** [Quest API] Add TrackNPC to Perl/Lua. ([#2272](https://github.com/EQEmu/Server/pull/2272)) 2022-06-29  
**hg** [Bug Fix] Fix empty spawned merchants ([#2275](https://github.com/EQEmu/Server/pull/2275)) 2022-06-28  
**Kinglykrab** [Spells] Target's Target Combat Range Rule ([#2274](https://github.com/EQEmu/Server/pull/2274)) 2022-06-24  
**Quintinon** [Bug Fix] Delete NpcType Struct returned by Bot::CreateDefaultNPCTypeStructForBot() when unused ([#2267](https://github.com/EQEmu/Server/pull/2267)) 2022-06-18  
**Quintinon** [Bug Fix] Free return value of ZoneDatabase::LoadTraderItemWithCharges() ([#2264](https://github.com/EQEmu/Server/pull/2264)) 2022-06-18  
**Quintinon** [Bug Fix] Correct type signed/unsigned int when reading item from database in shareddb ([#2269](https://github.com/EQEmu/Server/pull/2269)) 2022-06-15  
**Kinglykrab** [Loot] Remove unnecessary loot error messages. ([#2261](https://github.com/EQEmu/Server/pull/2261)) 2022-06-12  
**Akkadius** [Hotfix] Correct database call to point to the content_db connection 2022-06-12  
**Michael** [Bug Fix] Tradeskill Item 0 Error ([#2256](https://github.com/EQEmu/Server/pull/2256)) 2022-06-10  
**Michael** [Bug Fix] Tradeskill Autocombine MinSkill ([#2260](https://github.com/EQEmu/Server/pull/2260)) 2022-06-10  
**Michael** [Quest API] Expand SaveGuardSpot ([#2258](https://github.com/EQEmu/Server/pull/2258)) 2022-06-10  
**Kinglykrab** [Bug Fix] Stop skill ups on Charmed NPCs. ([#2249](https://github.com/EQEmu/Server/pull/2249)) 2022-06-09  
**Akkadius** [Discord Integration] Native Discord Integration ([#2140](https://github.com/EQEmu/Server/pull/2140)) 2022-06-09  
**hg** [Bug Fix] Fix stack leaks in Lua events [#2254](https://github.com/EQEmu/Server/pull/2254) 2022-06-09  
**Akkadius** [Hotfix] Flipped positive / negative values for legacy_combat.lua 2022-06-09  
**Kinglykrab** [Messages] Convert messages from Spells to FocusEffect where necessary. ([#2243](https://github.com/EQEmu/Server/pull/2243)) 2022-06-08  
**Michael Cook (mackal)** Fix memory leaks found by Quint ([#2248](https://github.com/EQEmu/Server/pull/2248)) 2022-06-07  
**Michael** [Bug Fix] Bazaar Search MYSQL Error ([#2252](https://github.com/EQEmu/Server/pull/2252)) 2022-06-07  
**Michael** [Bug Fix] Hacker_Str was causing sql errors - Non Escaped ([#2251](https://github.com/EQEmu/Server/pull/2251)) 2022-06-07  
**Kinglykrab** [Rules] Add Frontal Stun Immunity Rules. ([#2217](https://github.com/EQEmu/Server/pull/2217)) 2022-06-06  
**Quintinon** [Bug Fix] Fix two invalid data accesses in zone/client.cpp ([#2238](https://github.com/EQEmu/Server/pull/2238)) 2022-06-06  
**Kinglykrab** [Bug Fix] Fix issue where #advnpcspawn addspawn does not add spawn sometimes. ([#2247](https://github.com/EQEmu/Server/pull/2247)) 2022-06-06  
**Michael Cook (mackal)** [Memory Leak] Fix leak of CommandRecords in commandlist ([#2244](https://github.com/EQEmu/Server/pull/2244)) 2022-06-06  
**Kinglykrab** [Commands] Add BestZ and Region Data to #loc ([#2245](https://github.com/EQEmu/Server/pull/2245)) 2022-06-05  
**Kinglykrab** [Rules] Add Rule to Disable NPC Last Names. ([#2227](https://github.com/EQEmu/Server/pull/2227)) 2022-06-04  
**Kinglykrab** [Cleanup] Cleanup Haste references and Lua API calls for unsigned to signed. ([#2240](https://github.com/EQEmu/Server/pull/2240)) 2022-06-04  
**Kinglykrab** [Cleanup] Add Validation to varchar number item fields. ([#2241](https://github.com/EQEmu/Server/pull/2241)) 2022-06-04  
**Kinglykrab** [Commands] Cleanup #spawneditmass Command. ([#2229](https://github.com/EQEmu/Server/pull/2229)) 2022-06-04  
**Quintinon** [Bug Fix] Correct (probably) unintended bitwise AND instead of logical AND ([#2239](https://github.com/EQEmu/Server/pull/2239)) 2022-06-01  
**Quintinon** [Combat] Fix shield calculation ([#2234](https://github.com/EQEmu/Server/pull/2234)) 2022-06-01  
**Quintinon** [Code Cleanup] Resharper Warnings ([#2235](https://github.com/EQEmu/Server/pull/2235)) 2022-06-01  
**Kinglykrab** [Rules] Add Rules to disable various item functionalities and cleanup data types. ([#2225](https://github.com/EQEmu/Server/pull/2225)) 2022-06-01  
**Kinglykrab** [Bug Fix] Fix MovePC in #zone and #zoneinstance Commands. ([#2236](https://github.com/EQEmu/Server/pull/2236)) 2022-06-01  
**Kinglykrab** [Commands] Cleanup #date Command. ([#2228](https://github.com/EQEmu/Server/pull/2228)) 2022-06-01  
**Akkadius** [Loading] Zone Version Loading Fixes ([#2233](https://github.com/EQEmu/Server/pull/2233)) 2022-05-31  
**Akkadius** [Bug Fix] Adjustment for nullptr crash ([#2232](https://github.com/EQEmu/Server/pull/2232)) 2022-05-31  
**Akkadius** [Bug Fix] Fix null pointer crash on zones that have not booted a zone yet with #reload commands or anything that calls GetZoneDescription ([#2231](https://github.com/EQEmu/Server/pull/2231)) 2022-05-31  
**Akkadius** [Tasks] Fix validation loading ([#2230](https://github.com/EQEmu/Server/pull/2230)) 2022-05-31  
**titanium-forever** [Database Backup] Enable database dump of bot data ([#2221](https://github.com/EQEmu/Server/pull/2221)) 2022-06-01  
**Kinglykrab** [Quest API] Fix parameters in some Perl worldwide methods. ([#2224](https://github.com/EQEmu/Server/pull/2224)) 2022-05-31  
**Kinglykrab** [Bug Fix] Fix Legacy Combat Lua Script ([#2226](https://github.com/EQEmu/Server/pull/2226)) 2022-05-30  
**Kinglykrab** [Hot Fix] Fix Linux compile due to missing include. ([#2223](https://github.com/EQEmu/Server/pull/2223)) 2022-05-30  
**Kinglykrab** [INT64] Further int64 cleanup in Perl SetHP() and GetSpellHPBonuses() in Perl/Lua. ([#2222](https://github.com/EQEmu/Server/pull/2222)) 2022-05-29  
**Kinglykrab** [Commands] Cleanup #nudge Command. ([#2220](https://github.com/EQEmu/Server/pull/2220)) 2022-05-29  
**Kinglykrab** [Commands] Cleanup #emptyinventory Command. ([#2219](https://github.com/EQEmu/Server/pull/2219)) 2022-05-29  
**Kinglykrab** [INT64] Fix int64 for OOC Regen and GetHP(), GetMaxHP(), GetItemHPBonuses() in Perl/Lua. ([#2218](https://github.com/EQEmu/Server/pull/2218)) 2022-05-29  
**Kinglykrab** [Money Messages] Cleanup quest::givecash(), split, and task reward messages. ([#2205](https://github.com/EQEmu/Server/pull/2205)) 2022-05-29  
**Kinglykrab** [Cleanup] Cleanup spell and max level bucket logic. ([#2181](https://github.com/EQEmu/Server/pull/2181)) 2022-05-28  
**Kinglykrab** [Commands] Cleanup #npcedit, #lastname, #title, and #titlesuffix Commands. ([#2215](https://github.com/EQEmu/Server/pull/2215)) 2022-05-28  
**Kinglykrab** [Bug Fix] Fix IP Exemptions. ([#2189](https://github.com/EQEmu/Server/pull/2189)) 2022-05-27  
**Paul Coene** [Bug Fix] IsDamage test for lifetap was not complete. ([#2213](https://github.com/EQEmu/Server/pull/2213)) 2022-05-27  
**Kinglykrab** [Commands] Cleanup #findaliases and #help Commands. ([#2204](https://github.com/EQEmu/Server/pull/2204)) 2022-05-27  
**Kinglykrab** [Commands] Cleanup #zone and #zoneinstance Commands. ([#2202](https://github.com/EQEmu/Server/pull/2202)) 2022-05-27  
**Kinglykrab** [Rules] Add Spells:BuffsFadeOnDeath. ([#2200](https://github.com/EQEmu/Server/pull/2200)) 2022-05-27  
**Kinglykrab** [Commands] Cleanup #findclass and #findrace Commands. ([#2211](https://github.com/EQEmu/Server/pull/2211)) 2022-05-27  
**Kinglykrab** [Commands] Cleanup #corpsefix Command. ([#2197](https://github.com/EQEmu/Server/pull/2197)) 2022-05-27  
**Kinglykrab** [Commands] Cleanup #level Command. ([#2203](https://github.com/EQEmu/Server/pull/2203)) 2022-05-27  
**Kinglykrab** [Rules] Add Spells:IllusionsAlwaysPersist. ([#2199](https://github.com/EQEmu/Server/pull/2199)) 2022-05-27  
**Kinglykrab** [Regen] Fix possible overflow in CalcHPRegenCap(). ([#2185](https://github.com/EQEmu/Server/pull/2185)) 2022-05-27  
**Kinglykrab** [Commands] Cleanup #oocmute Command. ([#2191](https://github.com/EQEmu/Server/pull/2191)) 2022-05-27  
**Michael Cook (mackal)** Revert "[Aggro] Rooted mobs will add other hated targets to Hate list ([#2180](https://github.com/EQEmu/Server/pull/2180))" ([#2214](https://github.com/EQEmu/Server/pull/2214)) 2022-05-27  
**Paul Coene** [Aggro] Rooted mobs will add other hated targets to Hate list ([#2180](https://github.com/EQEmu/Server/pull/2180)) 2022-05-27  
**titanium-forever** Create user directory during account creation to ensure default files are copied to profile from /etc/skel ([#2176](https://github.com/EQEmu/Server/pull/2176)) 2022-05-27  
**Kinglykrab** [Bug Fix] Fix bot compile locking client on server enter. ([#2210](https://github.com/EQEmu/Server/pull/2210)) 2022-05-26  
**Kinglykrab** [Commands] Fix typos in #ban and #ipban Commands. ([#2209](https://github.com/EQEmu/Server/pull/2209)) 2022-05-25  
**Chris** [Bug Fix] Blocked spells max spell id increased ([#2207](https://github.com/EQEmu/Server/pull/2207)) 2022-05-25  
**Kinglykrab** [Bug Fix] Fix HP Regen Per Second. ([#2206](https://github.com/EQEmu/Server/pull/2206)) 2022-05-25  
**Kinglykrab** [Bot Commands] Use Account Status Constants. ([#2201](https://github.com/EQEmu/Server/pull/2201)) 2022-05-23  
**Kinglykrab** [Commands] Remove unused/broken #deletegraveyard and #setgraveyard Commands. ([#2198](https://github.com/EQEmu/Server/pull/2198)) 2022-05-23  
**Kinglykrab** [Commands] Cleanup #motd Command. ([#2190](https://github.com/EQEmu/Server/pull/2190)) 2022-05-23  
**Kinglykrab** [Rules] Cleanup all unused rules. ([#2184](https://github.com/EQEmu/Server/pull/2184)) 2022-05-23  
**Kinglykrab** [Commands] #bind Typo. ([#2196](https://github.com/EQEmu/Server/pull/2196)) 2022-05-22  
**Kinglykrab** [Commands] Cleanup #kill Command. ([#2195](https://github.com/EQEmu/Server/pull/2195)) 2022-05-22  
**Kinglykrab** [Bug Fix] Fix bot guild removal. ([#2194](https://github.com/EQEmu/Server/pull/2194)) 2022-05-22  
**Kinglykrab** [Commands] Consolidate #lock and #unlock Commands into #serverlock. ([#2193](https://github.com/EQEmu/Server/pull/2193)) 2022-05-22  
**Akkadius** Schema consistency fixes ([#2192](https://github.com/EQEmu/Server/pull/2192)) 2022-05-21  
**Kinglykrab** [Cleanup] Move Client::Undye() to client.cpp from #path Command. ([#2188](https://github.com/EQEmu/Server/pull/2188)) 2022-05-21  
**Paul Coene** [Bug Fix] Fix duplicate and missing messages due to innate in spells ([#2170](https://github.com/EQEmu/Server/pull/2170)) 2022-05-20  
**Kinglykrab** [Titles] Cleanup titles, title suffix, and last name methods. ([#2174](https://github.com/EQEmu/Server/pull/2174)) 2022-05-19  
**Kinglykrab** [Quest API] Add CheckNameFilter to Perl/Lua. ([#2175](https://github.com/EQEmu/Server/pull/2175)) 2022-05-19  
**Kinglykrab** [Cleanup] Quest API push methods using invalid types. ([#2172](https://github.com/EQEmu/Server/pull/2172)) 2022-05-15  
**Kinglykrab** [Cleanup] Remove unused methods. ([#2171](https://github.com/EQEmu/Server/pull/2171)) 2022-05-15  
**hg** Save eyes in #npcedit featuresave ([#2178](https://github.com/EQEmu/Server/pull/2178)) 2022-05-15  
**Kinglykrab** [Quest API] Add TaskSelector to Perl/Lua. ([#2177](https://github.com/EQEmu/Server/pull/2177)) 2022-05-15  
**hg** [Opcode] Implement SetFace opcode ([#2167](https://github.com/EQEmu/Server/pull/2167)) 2022-05-11  
**Kinglykrab** [Quest API] Add GetHealAmount() and GetSpellDamage() to Perl/Lua. ([#2165](https://github.com/EQEmu/Server/pull/2165)) 2022-05-11  
**Kinglykrab** [Cleanup] Cleanup #kick message. ([#2164](https://github.com/EQEmu/Server/pull/2164)) 2022-05-10  
**Kinglykrab** [Hot Fix] Off by on in Merchant Loading. ([#2166](https://github.com/EQEmu/Server/pull/2166)) 2022-05-10  
**Kinglykrab** [Commands] #reload Command Overhaul. ([#2162](https://github.com/EQEmu/Server/pull/2162)) 2022-05-10  
**Akkadius** [int64] Hate Fixes ([#2163](https://github.com/EQEmu/Server/pull/2163)) 2022-05-09  
**Kinglykrab** [Merchants] Add Merchant Data Bucket capability. ([#2160](https://github.com/EQEmu/Server/pull/2160)) 2022-05-09  
**Kinglykrab** [Cleanup] Possible issues with variable/parameter name equality. ([#2161](https://github.com/EQEmu/Server/pull/2161)) 2022-05-09  
**Akkadius** [Hotfix] Fix regression caused by [#2129](https://github.com/EQEmu/Server/pull/2129) 2022-05-09  
**Michael Cook (mackal)** Fix out of bounds issues with SPA 288 ([#2157](https://github.com/EQEmu/Server/pull/2157)) 2022-05-08  
**Kinglykrab** [Bug FIx] Fix #repop Command. ([#2159](https://github.com/EQEmu/Server/pull/2159)) 2022-05-08  
**Kinglykrab** [Bug Fix] Make Perl TakeMoneyFromPP int64 ([#2158](https://github.com/EQEmu/Server/pull/2158)) 2022-05-08  
**Kinglykrab** [Bug Fix] Fix possible issue where variables have the same name. ([#2156](https://github.com/EQEmu/Server/pull/2156)) 2022-05-08  
**Akkadius** [int64] Windows Compile Fixes ([#2155](https://github.com/EQEmu/Server/pull/2155)) 2022-05-07  
**Kinglykrab** [Commands] Cleanup #cvs Command. ([#2153](https://github.com/EQEmu/Server/pull/2153)) 2022-05-07  
**Akkadius** [int64] Support for HP / Mana / End / Damage / Hate ([#2091](https://github.com/EQEmu/Server/pull/2091)) 2022-05-07  
**Kinglykrab** [Commands] Cleanup #ban, #ipban, #flag, #kick, #setlsinfo, and #setpass Commands. ([#2104](https://github.com/EQEmu/Server/pull/2104)) 2022-05-07  
**Akkadius** [UCS] Auto Client Reconnection ([#2154](https://github.com/EQEmu/Server/pull/2154)) 2022-05-07  
**Akkadius** [Hotfix] Fix DB version merge 2022-05-07  
**Akkadius** [Tasks] Implement Task Goal Match List ([#2097](https://github.com/EQEmu/Server/pull/2097)) 2022-05-07  
**Kinglykrab** [Commands] Cleanup #time and #timezone Command. ([#2147](https://github.com/EQEmu/Server/pull/2147)) 2022-05-07  
**Kinglykrab** [Bug Fix] NPC::GetNPCStat has no default return. ([#2150](https://github.com/EQEmu/Server/pull/2150)) 2022-05-07  
**Kinglykrab** [Bug Fix] Lua GetBlockNextSpell() no return. ([#2151](https://github.com/EQEmu/Server/pull/2151)) 2022-05-07  
**Kinglykrab** [Commands] Cleanup #reloadzps Command. ([#2129](https://github.com/EQEmu/Server/pull/2129)) 2022-05-07  
**Kinglykrab** [Commands] Cleanup #reloadtraps Command. ([#2126](https://github.com/EQEmu/Server/pull/2126)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #reloadlevelmods Command. ([#2122](https://github.com/EQEmu/Server/pull/2122)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #npctype_cache Command. ([#2109](https://github.com/EQEmu/Server/pull/2109)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #npcspecialattk Command. ([#2108](https://github.com/EQEmu/Server/pull/2108)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #ucs Command. ([#2149](https://github.com/EQEmu/Server/pull/2149)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #reloadaa Command. ([#2120](https://github.com/EQEmu/Server/pull/2120)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #profanity Command. ([#2113](https://github.com/EQEmu/Server/pull/2113)) 2022-05-06  
**Kinglykrab** [Commands] Remove unused #bestz and #pf Commands. ([#2112](https://github.com/EQEmu/Server/pull/2112)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #makepet Command. ([#2105](https://github.com/EQEmu/Server/pull/2105)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #npcemote Command. ([#2106](https://github.com/EQEmu/Server/pull/2106)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #npcsay and #npcshout Commands. ([#2107](https://github.com/EQEmu/Server/pull/2107)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #npctypespawn Command. ([#2110](https://github.com/EQEmu/Server/pull/2110)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #reloadmerchants Command. ([#2123](https://github.com/EQEmu/Server/pull/2123)) 2022-05-06  
**Kinglykrab** [Commmands] Cleanup #questerrors Command. ([#2116](https://github.com/EQEmu/Server/pull/2116)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #randomizefeatures Command. ([#2118](https://github.com/EQEmu/Server/pull/2118)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #refreshgroup Command. ([#2119](https://github.com/EQEmu/Server/pull/2119)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #push Command. ([#2114](https://github.com/EQEmu/Server/pull/2114)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #qglobal Command. ([#2115](https://github.com/EQEmu/Server/pull/2115)) 2022-05-06  
**Kinglykrab** [Bug Fix] Resolve subroutine redefinition due to bot methods. ([#2117](https://github.com/EQEmu/Server/pull/2117)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #resetaa_timer Command. ([#2131](https://github.com/EQEmu/Server/pull/2131)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #revoke Command. ([#2134](https://github.com/EQEmu/Server/pull/2134)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #reloadallrules Command. ([#2121](https://github.com/EQEmu/Server/pull/2121)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #reloadperlexportsettings Command. ([#2124](https://github.com/EQEmu/Server/pull/2124)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #reloadtitles Command. ([#2125](https://github.com/EQEmu/Server/pull/2125)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #reloadworld and #repop Command. ([#2127](https://github.com/EQEmu/Server/pull/2127)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #summonburiedplayercorpse Command. ([#2146](https://github.com/EQEmu/Server/pull/2146)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #trapinfo Command. ([#2148](https://github.com/EQEmu/Server/pull/2148)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #reloadrulesworld Command. ([#2128](https://github.com/EQEmu/Server/pull/2128)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #reloadstatic Command. ([#2130](https://github.com/EQEmu/Server/pull/2130)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #resetaa Command. ([#2132](https://github.com/EQEmu/Server/pull/2132)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #sensetrap Command. ([#2137](https://github.com/EQEmu/Server/pull/2137)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #serverinfo Command. ([#2138](https://github.com/EQEmu/Server/pull/2138)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #resetdisc_timer Command. ([#2133](https://github.com/EQEmu/Server/pull/2133)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #roambox Command. ([#2135](https://github.com/EQEmu/Server/pull/2135)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #save Command. ([#2136](https://github.com/EQEmu/Server/pull/2136)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #serverrules Command. ([#2139](https://github.com/EQEmu/Server/pull/2139)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #shownpcgloballoot and #showzonegloballoot Command. ([#2141](https://github.com/EQEmu/Server/pull/2141)) 2022-05-06  
**Kinglykrab** [Commands] Add #feature Command. ([#2142](https://github.com/EQEmu/Server/pull/2142)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #spawnfix Command. ([#2143](https://github.com/EQEmu/Server/pull/2143)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #summon Command. ([#2145](https://github.com/EQEmu/Server/pull/2145)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #gassign Command. ([#2101](https://github.com/EQEmu/Server/pull/2101)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #spawnstatus Command. ([#2144](https://github.com/EQEmu/Server/pull/2144)) 2022-05-06  
**Kinglykrab** [Commands] Cleanup #attack Command. ([#2103](https://github.com/EQEmu/Server/pull/2103)) 2022-05-03  
**Kinglykrab** [Commands] Cleanup #freeze and #unfreeze Commands. ([#2102](https://github.com/EQEmu/Server/pull/2102)) 2022-05-03  
**Kinglykrab** [Commands] Cleanup #getvariable Command. ([#2100](https://github.com/EQEmu/Server/pull/2100)) 2022-05-03  
**Kinglykrab** [Quest API] Expand Bot quest API functionality. ([#2096](https://github.com/EQEmu/Server/pull/2096)) 2022-05-03  
**Kinglykrab** [Quest API] Perl Money Fixes. ([#2098](https://github.com/EQEmu/Server/pull/2098)) 2022-05-03  
**Kinglykrab** [Quest API] Add commify to Perl/Lua. ([#2099](https://github.com/EQEmu/Server/pull/2099)) 2022-05-03  
**Paul Coene** [Combat] Fix Frenzy vs opponents immune to non-magic ([#2095](https://github.com/EQEmu/Server/pull/2095)) 2022-05-02  
**hg** [Repository] Cast floats to avoid grid repository warnings ([#2094](https://github.com/EQEmu/Server/pull/2094)) 2022-05-01  
**Akkadius** [Drone] Speed up drone builds ([#2092](https://github.com/EQEmu/Server/pull/2092)) 2022-05-01  
**hg** Remove already defined method ([#2093](https://github.com/EQEmu/Server/pull/2093)) 2022-05-01  
**Kinglykrab** [Commands] Add additional #peqzone functionality. ([#2085](https://github.com/EQEmu/Server/pull/2085)) 2022-05-01  
**Akkadius** [Combat] Basic Combat Recording ([#2090](https://github.com/EQEmu/Server/pull/2090)) 2022-05-01  
**Kinglykrab** [Bots] Bot::PerformTradeWithClient Cleanup. ([#2084](https://github.com/EQEmu/Server/pull/2084)) 2022-05-01  
**Akkadius** [Refactor] Simplify NPC Loading ([#2087](https://github.com/EQEmu/Server/pull/2087)) 2022-05-01  
**Akkadius** [Compile] Decrease build times using unity build strategy ([#2089](https://github.com/EQEmu/Server/pull/2089)) 2022-05-01  
**KayenEQ** [Spells] Update to target types Beam and Cone to ignore invalid targets. ([#2080](https://github.com/EQEmu/Server/pull/2080)) 2022-05-01  
**Akkadius** [Logging] Update BUILD_LOGGING=false Blank Aliases ([#2083](https://github.com/EQEmu/Server/pull/2083)) 2022-05-01  
**Akkadius** [Regen] Implement Per Second HP Regen for NPCs ([#2086](https://github.com/EQEmu/Server/pull/2086)) 2022-05-01  
**Akkadius** [Code Cleanup] Remove this-> in code where its implied ([#2088](https://github.com/EQEmu/Server/pull/2088)) 2022-05-01  
**Kinglykrab** [Bots] Remove unused methods. ([#2082](https://github.com/EQEmu/Server/pull/2082)) 2022-04-30  
**Kinglykrab** [Quest API] Add multiple inventory method short hands to client. ([#2078](https://github.com/EQEmu/Server/pull/2078)) 2022-04-30  
**Kinglykrab** [Bots] Fix ^dyearmor command math. ([#2081](https://github.com/EQEmu/Server/pull/2081)) 2022-04-30  
**Kinglykrab** [Quest API] Add AddPlatinum(), GetCarriedPlatinum() and TakePlatinum() to Perl/Lua. ([#2079](https://github.com/EQEmu/Server/pull/2079)) 2022-04-30  
**nytmyr** [Quest API] Add EVENT_SKILL_UP & EVENT_LANGUAGE_SKILL_UP to Perl/Lua ([#2076](https://github.com/EQEmu/Server/pull/2076)) 2022-04-25  
**cybernine186** Bug Fix for WorldServer::HandleMessage, CZUpdateType_NPC ([#2074](https://github.com/EQEmu/Server/pull/2074)) 2022-04-21  
**Paul Coene** [Bug Fix] Restore missing messages for lifetap and dmg spells. ([#2057](https://github.com/EQEmu/Server/pull/2057)) 2022-04-13  
**Kinglykrab** [Commands] Cleanup #task Command. ([#2071](https://github.com/EQEmu/Server/pull/2071)) 2022-04-13  
**KayenEQ** [Bug Fix] Blocked spells max spell id increased ([#2073](https://github.com/EQEmu/Server/pull/2073)) 2022-04-13  
**KayenEQ** [API] Methods for getting more information on quest timers. ([#2060](https://github.com/EQEmu/Server/pull/2060)) 2022-04-13  
**KayenEQ** [Bug Fix] Instrument Mods should not affect spells that change model size. ([#2072](https://github.com/EQEmu/Server/pull/2072)) 2022-04-13  
**Kinglykrab** [Quest API] Add GetBotListByCharacterID() to Perl/Lua. ([#2069](https://github.com/EQEmu/Server/pull/2069)) 2022-04-02  
**Paul Coene** [Bug Fix] Fix recipient sound (vtell) on non-player races ([#2066](https://github.com/EQEmu/Server/pull/2066)) 2022-04-02  
**Kinglykrab** [Bug Fix] Clear title/suffix bug fix. ([#2068](https://github.com/EQEmu/Server/pull/2068)) 2022-04-02  
**KayenEQ** [Bug Fix] Bard Invisible causing display issues. ([#2067](https://github.com/EQEmu/Server/pull/2067)) 2022-04-01  
**Kinglykrab** [Quest API] Add GetBotListByClientName(client_name) to Perl/Lua. ([#2064](https://github.com/EQEmu/Server/pull/2064)) 2022-03-23  
**Natedog2012** [Bug Fix] #peqzone no longer bypass Handle_OP_ZoneChange ([#2063](https://github.com/EQEmu/Server/pull/2063)) 2022-03-19  
**Paul Coene** [Bug Fix] Fix for being able to skill up on corspe. ([#2058](https://github.com/EQEmu/Server/pull/2058)) 2022-03-19  
**Paul Coene** [Bug Fix] manifest for db version 9176 had incorrect field name([#2062](https://github.com/EQEmu/Server/pull/2062)) 2022-03-19  
**Natedog2012** [Bug Fix] Force NPCs to respect special ability 24 and 50 when set on player pets ([#2059](https://github.com/EQEmu/Server/pull/2059)) 2022-03-15  
**Paul Coene** [Logging] Fix log messages to final damage values ([#2056](https://github.com/EQEmu/Server/pull/2056)) 2022-03-14  
**Kinglykrab** [Bug Fix] Fix possible crash with zone name methods. ([#2055](https://github.com/EQEmu/Server/pull/2055)) 2022-03-13  
**Kinglykrab** [Quest API] Add AddItem() to Perl/Lua. ([#2054](https://github.com/EQEmu/Server/pull/2054)) 2022-03-13  
**KayenEQ** [API] GetNPCStat default better naming ([#2053](https://github.com/EQEmu/Server/pull/2053)) 2022-03-12  
**Kinglykrab** [Quest API] Allow EVENT_ZONE to be parsed as non-zero to prevent zoning. ([#2052](https://github.com/EQEmu/Server/pull/2052)) 2022-03-12  
**Kinglykrab** [Quest API] Add EVENT_CAST_ON exports to EVENT_CAST and EVENT_CAST_BEGIN. ([#2051](https://github.com/EQEmu/Server/pull/2051)) 2022-03-12  
**Kinglykrab** [Quest API] Export killed XYZH to EVENT_DEATH_ZONE in Perl. ([#2050](https://github.com/EQEmu/Server/pull/2050)) 2022-03-12  
**catapultam-habeo** [Bots] Update Bot Heal & Damage methods to more closely match Clients + Bugfixes ([#2045](https://github.com/EQEmu/Server/pull/2045)) 2022-03-11  
**Akkadius** [Repositories] Update repositories ([#2040](https://github.com/EQEmu/Server/pull/2040)) 2022-03-11  
**Akkadius** [Bug FIx] Saylink Collation Database Edge Case ([#2039](https://github.com/EQEmu/Server/pull/2039)) 2022-03-11  
**KayenEQ** [Bug Fix] #tune command various fixes ([#2046](https://github.com/EQEmu/Server/pull/2046)) 2022-03-11  
**KayenEQ** [API] GetNPCStat can now return default stat values. ([#2048](https://github.com/EQEmu/Server/pull/2048)) 2022-03-11  
**Kinglykrab** [Quest API] Add caster_id and caster_level export to EVENT_CAST_ON in Perl/Lua. ([#2049](https://github.com/EQEmu/Server/pull/2049)) 2022-03-10  
**Kinglykrab** [Bug Fix] Spell Buckets/Globals did not allow string-based values. ([#2043](https://github.com/EQEmu/Server/pull/2043)) 2022-03-09  
**KayenEQ** [API] Perl functions to set invulnerable to and modify environmental damage. ([#2044](https://github.com/EQEmu/Server/pull/2044)) 2022-03-08  
**KayenEQ** [Bug Fix] Invisible will display as dropped now on air pets when they attack. ([#2042](https://github.com/EQEmu/Server/pull/2042)) 2022-03-07  
**JJ** [Database] Update 2022_01_10_checksum_verification.sql ([#2041](https://github.com/EQEmu/Server/pull/2041)) 2022-03-07  
**KayenEQ** [Bug Fix] PR 2032 would lock client on casting fail as written ([#2038](https://github.com/EQEmu/Server/pull/2038)) 2022-03-06  
**Akkadius** [Command] Fix #killallnpcs from crashing ([#2037](https://github.com/EQEmu/Server/pull/2037)) 2022-03-06  
**Randy Girard** [Content Filtering] Updates contents flags to be checked at runtime. ([#1940](https://github.com/EQEmu/Server/pull/1940)) 2022-03-06  
**Akkadius** [Database] Add Primary ID Keys to Tables ([#2036](https://github.com/EQEmu/Server/pull/2036)) 2022-03-06  
**Paul Coene** [Feature] Client Checksum Verification (Resubmit old 1678) ([#1922](https://github.com/EQEmu/Server/pull/1922)) 2022-03-06  
**catapultam-habeo** [Feature] EQ2-style implied targeting for spells. ([#2032](https://github.com/EQEmu/Server/pull/2032)) 2022-03-06  
**Paul Coene** [AI] Spell Type (1024) InCombatBuff were spam casting ([#2030](https://github.com/EQEmu/Server/pull/2030)) 2022-03-06  
**catapultam-habeo** [Feature] Allow pets to zone with permanent (buffdurationformula 50) buffs to maintain them through zone transitions ([#2035](https://github.com/EQEmu/Server/pull/2035)) 2022-03-06  
**Paul Coene** [Bug Fix] Bandolier didn't recognize source weapon on cursor ([#2026](https://github.com/EQEmu/Server/pull/2026)) 2022-03-06  
**Paul Coene** [Bug Fix] Fixed several instances of incorrect comparision - & executes after == ([#2025](https://github.com/EQEmu/Server/pull/2025)) 2022-03-06  
**neckkola** [Bots] Fix bot spawn when bot id = char_id ([#1984](https://github.com/EQEmu/Server/pull/1984)) 2022-03-06  
**catapultam-habeo** [Bots] Apply Spells:IgnoreSpellDmgLvlRestriction to bots ([#2024](https://github.com/EQEmu/Server/pull/2024)) 2022-03-06  
**KayenEQ** [Bug Fix] Missing break ([#2031](https://github.com/EQEmu/Server/pull/2031)) 2022-03-04  
**Michael Cook (mackal)** Fix order of operations issues in worldshutdown command ([#2029](https://github.com/EQEmu/Server/pull/2029)) 2022-03-04  
**Kinglykrab** [Bug Fix] Objects::GetTiltX() and Objects::GetTiltY() Perl Croak Typos. ([#2028](https://github.com/EQEmu/Server/pull/2028)) 2022-03-03  
**Kinglykrab** [Bug Fix] Doors::GetSize() Perl Croak Typo. ([#2027](https://github.com/EQEmu/Server/pull/2027)) 2022-03-03  
**KayenEQ** [Spells] Fixes for numhits type 7 counter incrementing incorrectly. ([#2022](https://github.com/EQEmu/Server/pull/2022)) 2022-03-01  
**Kinglykrab** [Commands] Cleanup #worldwide command. ([#2021](https://github.com/EQEmu/Server/pull/2021)) 2022-02-28  
**catapultam-habeo** Rule to apply Spell Dmg and Heal Amount stats as a percentage instead of flat value. ([#2017](https://github.com/EQEmu/Server/pull/2017)) 2022-02-28  
**Kinglykrab** [Bug Fix] quest::MovePCInstance() Arguments Fix. ([#2020](https://github.com/EQEmu/Server/pull/2020)) 2022-02-27  
**Kinglykrab** [Bug Fix] Spell Buckets/Globals SQL Escape. ([#2019](https://github.com/EQEmu/Server/pull/2019)) 2022-02-26  
**KayenEQ** [Spells] SPA 79 SE_CurrentHPOnce now will check for focus, critical and partial resist checks, except for buffs. ([#2018](https://github.com/EQEmu/Server/pull/2018)) 2022-02-24  
**Kinglykrab** [Quest API] Add EVENT_EQUIP_ITEM_CLIENT and EVENT_UNEQUIP_ITEM_CLIENT to Perl/Lua. ([#2015](https://github.com/EQEmu/Server/pull/2015)) 2022-02-22  
**Akkadius** [GM Command] Fix Crash Issue and Validation with #zclip ([#2014](https://github.com/EQEmu/Server/pull/2014)) 2022-02-21  
**Kinglykrab** [Bug Fix] Alleviate some lag with crosszone/worldwide spell casting. ([#2016](https://github.com/EQEmu/Server/pull/2016)) 2022-02-20  
**KayenEQ** [Bug Fix] checking casting_spell_slot before its defined is bad ([#2013](https://github.com/EQEmu/Server/pull/2013)) 2022-02-20  
**KayenEQ** [API] perl added GetNPCStat(identifier) ([#2012](https://github.com/EQEmu/Server/pull/2012)) 2022-02-20  
**KayenEQ** [Spells] Bard songs from item clickies should not require components ([#2011](https://github.com/EQEmu/Server/pull/2011)) 2022-02-18  
**KayenEQ** [Bug Fix] Fix for castspell command ([#2010](https://github.com/EQEmu/Server/pull/2010)) 2022-02-17  
**Kinglykrab** [Commands] Bug fix for #logs command. ([#2008](https://github.com/EQEmu/Server/pull/2008)) 2022-02-17  
**KayenEQ** [API] Fix for SetBuffDuration function to check bard slots. ([#2009](https://github.com/EQEmu/Server/pull/2009)) 2022-02-17  
**Kinglykrab** [Quest API] Add GetRandomMob() and GetRandomNPC() to Perl/Lua. ([#2006](https://github.com/EQEmu/Server/pull/2006)) 2022-02-17  
**cybernine186** [Bug] Fixed trade items record log ([#2003](https://github.com/EQEmu/Server/pull/2003)) 2022-02-16  
**KayenEQ** [Spells] AE Duration effect (Rains) will now work with Target Ring and PBAE spells. ([#2000](https://github.com/EQEmu/Server/pull/2000)) 2022-02-16  
**Kinglykrab** [Commands] Cleanup #unscribespell Command. ([#1998](https://github.com/EQEmu/Server/pull/1998)) 2022-02-16  
**Kinglykrab** [Commands] Cleanup #untraindisc Command. ([#1996](https://github.com/EQEmu/Server/pull/1996)) 2022-02-16  
**KayenEQ** fix for meloday cast bar issue ([#2005](https://github.com/EQEmu/Server/pull/2005)) 2022-02-15  
**KayenEQ** invis updates ([#2004](https://github.com/EQEmu/Server/pull/2004)) 2022-02-15  
**KayenEQ** [API] Apply spells with custom buff durations and adjust existing spell buff durations. ([#1997](https://github.com/EQEmu/Server/pull/1997)) 2022-02-15  
**KayenEQ** Update bonuses.cpp ([#2002](https://github.com/EQEmu/Server/pull/2002)) 2022-02-15  
**KayenEQ** [Spells] Invisibility updates and rework ([#1991](https://github.com/EQEmu/Server/pull/1991)) 2022-02-15  
**KayenEQ** bug fix ([#2001](https://github.com/EQEmu/Server/pull/2001)) 2022-02-15  
**Kinglykrab** [Commands] Cleanup #showskills Command. ([#1994](https://github.com/EQEmu/Server/pull/1994)) 2022-02-14  
**Kinglykrab** [Commands] Cleanup #setskillall Command. ([#1992](https://github.com/EQEmu/Server/pull/1992)) 2022-02-14  
**KayenEQ** [Combat] /shield command "too far away message" ([#1999](https://github.com/EQEmu/Server/pull/1999)) 2022-02-14  
**KayenEQ** [Bug Fix] Edge case AA reset timer issue fixes ([#1995](https://github.com/EQEmu/Server/pull/1995)) 2022-02-13  
**KayenEQ** bot crash fix ([#1993](https://github.com/EQEmu/Server/pull/1993)) 2022-02-13  
**KimLS** Small change to order of a couple of files in zone cmakefile, added the gm commands to their own source group so they are cleanly separated in any IDEs cmake supports 2022-02-12  
**KimLS** Extra extern c warning squash, XS macro expands to XS_EXTERNAL which under c++ expands to extern c 2022-02-11  
**KimLS** What is this random semicoloned if, removed to silence warning 2022-02-11  
**KimLS** Handle default case of GetDynamicZoneTypeName to silence warning 2022-02-11  
**KimLS** Remove obsolete register spec on int to silence warning 2022-02-11  
**KayenEQ** [Bug Fix] Fix for Bot command casting ([#1990](https://github.com/EQEmu/Server/pull/1990)) 2022-02-12  
**Kinglykrab** [Commands] Cleanup #ai Command. ([#1980](https://github.com/EQEmu/Server/pull/1980)) 2022-02-11  
**KayenEQ** [Spells] Fix for AA recast timers not resetting properly ([#1989](https://github.com/EQEmu/Server/pull/1989)) 2022-02-11  
**KayenEQ** revert completed ([#1988](https://github.com/EQEmu/Server/pull/1988)) 2022-02-11  
**KimLS** Fix for passing std::string to vsprintf 2022-02-11  
**KayenEQ** group message ([#1987](https://github.com/EQEmu/Server/pull/1987)) 2022-02-10  
**Kinglykrab** [Commands] Cleanup #npcloot Command. ([#1974](https://github.com/EQEmu/Server/pull/1974)) 2022-02-10  
**Kinglykrab** [Commands] Cleanup #logs Command. ([#1969](https://github.com/EQEmu/Server/pull/1969)) 2022-02-10  
**Kinglykrab** [Commands] Cleanup #flagedit Command. ([#1968](https://github.com/EQEmu/Server/pull/1968)) 2022-02-10  
**Kinglykrab** [Commands] Cleanup #name Command. ([#1977](https://github.com/EQEmu/Server/pull/1977)) 2022-02-10  
**Kinglykrab** [Commands] Cleanup #hatelist Command. ([#1976](https://github.com/EQEmu/Server/pull/1976)) 2022-02-10  
**Kinglykrab** [Commands] Cleanup #netstats Command. ([#1970](https://github.com/EQEmu/Server/pull/1970)) 2022-02-10  
**Kinglykrab** [Commands] Cleanup #version Command. ([#1967](https://github.com/EQEmu/Server/pull/1967)) 2022-02-10  
**Kinglykrab** [Commands] Cleanup #undye and #undyeme Commands. ([#1966](https://github.com/EQEmu/Server/pull/1966)) 2022-02-10  
**Kinglykrab** [Commands] Cleanup #timers Command. ([#1965](https://github.com/EQEmu/Server/pull/1965)) 2022-02-10  
**Kinglykrab** [Quest API] Add GetEnvironmentalDamageName() to Perl/Lua. ([#1964](https://github.com/EQEmu/Server/pull/1964)) 2022-02-10  
**KayenEQ** [Bug Fix] Fix for PR1954 target restriction with npcpc_only_flag from groupbuffs ([#1986](https://github.com/EQEmu/Server/pull/1986)) 2022-02-10  
**KayenEQ** [Bug Fix] PR 1982 ([#1985](https://github.com/EQEmu/Server/pull/1985)) 2022-02-09  
**KayenEQ** fixed ([#1983](https://github.com/EQEmu/Server/pull/1983)) 2022-02-09  
**KayenEQ** [Spells] Support for SPA 194 SE_FadingMemories to use max level checks on aggroed mobs ([#1979](https://github.com/EQEmu/Server/pull/1979)) 2022-02-09  
**KayenEQ** [Bug Fix] Bard update fixes 1 ([#1982](https://github.com/EQEmu/Server/pull/1982)) 2022-02-09  
**Kinglykrab** [Quest API] Add AddAISpellEffect(spell_effect_id, base_value, limit_value, max_value) and RemoveAISpellEffect(spell_effect_id) to Lua. ([#1981](https://github.com/EQEmu/Server/pull/1981)) 2022-02-08  
**KayenEQ** (origin/release) [API] Perl functions added to apply spell effects directly to NPCs without requiring buffs. ([#1975](https://github.com/EQEmu/Server/pull/1975)) 2022-02-08  
**KayenEQ** [Spells] Allow damage spells to heal if quest based spell mitigation is over 100 pct. ([#1978](https://github.com/EQEmu/Server/pull/1978)) 2022-02-08  
**KayenEQ** procs silence ([#1973](https://github.com/EQEmu/Server/pull/1973)) 2022-02-08  
**KayenEQ** [Bug Fix] Summon Companion causing pets to warps away. ([#1972](https://github.com/EQEmu/Server/pull/1972)) 2022-02-08  
**KayenEQ** [Spells] Fix for AA and Discipline recast timers being set on spell casting failure. ([#1971](https://github.com/EQEmu/Server/pull/1971)) 2022-02-08  
**KayenEQ** [Spells] Major update to Bard song pulsing, Bard item clicks while singing, and spell casting restriction logic. ([#1954](https://github.com/EQEmu/Server/pull/1954)) 2022-02-07  
**Kinglykrab** (tag: Automated-Release-(Windows-x64)) [Quest API] Add inventory->CountItemEquippedByID(item_id) and inventory->HasItemEquippedByID(item_id) to Perl/Lua. ([#1963](https://github.com/EQEmu/Server/pull/1963)) 2022-02-06  
**Akkadius** [Maintenance Script] Pull from different maps mirror for now 2022-02-05  
**KayenEQ** escape fix for different target types ([#1962](https://github.com/EQEmu/Server/pull/1962)) 2022-02-04  
**KayenEQ** [Spells] SPA 311 SE_LimitCombatSkills should prevent focusing of procs even if proc is a 'casted' spell. ([#1961](https://github.com/EQEmu/Server/pull/1961)) 2022-02-04  
**KayenEQ** [Spells] Illusions will now persist onto the corpse when mob is killed. ([#1960](https://github.com/EQEmu/Server/pull/1960)) 2022-02-04  
**Kinglykrab** [Bug Fix] NPC::CountItem and Corpse::CountItem 0 Charge Item Fix. ([#1959](https://github.com/EQEmu/Server/pull/1959)) 2022-02-04  
**KayenEQ** [Bug Fix] Illusions will now properly display armor to other clients when they zone in. ([#1958](https://github.com/EQEmu/Server/pull/1958)) 2022-02-03  
**KayenEQ** [Spells] Swarm pet aggro logic fix ([#1956](https://github.com/EQEmu/Server/pull/1956)) 2022-02-03  
**Kinglykrab** [Commands] #ginfo Cleanup. ([#1955](https://github.com/EQEmu/Server/pull/1955)) 2022-02-03  
**Kinglykrab** [Commands] Cleanup #npceditmass command. ([#1957](https://github.com/EQEmu/Server/pull/1957)) 2022-02-03  
**KayenEQ** [Spells] Support for 'HateAdded' spell field to apply negative values to reduce hate. ([#1953](https://github.com/EQEmu/Server/pull/1953)) 2022-02-02

## 1/01/2022

**Akkadius** [Doors] Ignore Doors that Have Non-Zero Trigger or Door Param ([#1899](https://github.com/EQEmu/Server/pull/1899))  
**j883376** [Spells] Allow GMs to remove buffs from any target ([#1907](https://github.com/EQEmu/Server/pull/1907))  
**JJ** [Cleanup] holdzones not used. ([#1852](https://github.com/EQEmu/Server/pull/1852))  
**KayenEQ** [API] mob->AppearanceEffects improved functionality. ([#1821](https://github.com/EQEmu/Server/pull/1821))  
**KayenEQ** [API] mob->SpellEffect small perl fix ([#1820](https://github.com/EQEmu/Server/pull/1820))  
**KayenEQ** [Bug Fix] Bind Sight will now function properly ([#1825](https://github.com/EQEmu/Server/pull/1825))  
**KayenEQ** [Bug Fix] Hero Forge armor graphics not displaying properly to other clients in zone. ([#1883](https://github.com/EQEmu/Server/pull/1883))  
**KayenEQ** [Bug Fix] Numhits now display instantly on cast. ([#1826](https://github.com/EQEmu/Server/pull/1826))  
**KayenEQ** [Spells] Update to dispel to use SPA 272 SE_CastingLevel2 ([#1886](https://github.com/EQEmu/Server/pull/1886))  
**KayenEQ** [Spells] Minor update to SPA 502 Fearstuns max calculation. ([#1889](https://github.com/EQEmu/Server/pull/1889))  
**KayenEQ** [Features] Appearance Effects will now be sent to clients upon zone in. GM commands. ([#1874](https://github.com/EQEmu/Server/pull/1874))  
**KayenEQ** [Spells] Update to SPA 205 SE_Rampage ([#1882](https://github.com/EQEmu/Server/pull/1882))  
**KayenEQ** [Rule] Added rule to disable SPA 173 from making player immune to enrage. ([#1897](https://github.com/EQEmu/Server/pull/1897))  
**KayenEQ** [Spells] Bard AA clicks should not receive song modifiers. ([#1824](https://github.com/EQEmu/Server/pull/1824))  
**KayenEQ** [Spells] Eye of Zomm will now despawn and stack properly ([#1849](https://github.com/EQEmu/Server/pull/1849))  
**KayenEQ** [Spells] Fixed issue with permanent Illusions not being consistent when zoning. ([#1876](https://github.com/EQEmu/Server/pull/1876))  
**KayenEQ** [Spells] Implemented SPA 245 SE_TrapCircumvention ([#1885](https://github.com/EQEmu/Server/pull/1885))  
**KayenEQ** [Spells] Implemented SPA 281 SE_PetFeignMinion ([#1900](https://github.com/EQEmu/Server/pull/1900))  
**KayenEQ** [Spells] Rework of SPA 288 SE_SkillAttackProc ([#1893](https://github.com/EQEmu/Server/pull/1893))  
**KayenEQ** [Spells] SPA 310 SE_ReduceReuseTimer will now work on spell recast time ([#1856](https://github.com/EQEmu/Server/pull/1856))  
**KayenEQ** [Spells] Throwing procs fixed and other proc updates ([#1871](https://github.com/EQEmu/Server/pull/1871))  
**KayenEQ** [Spells] Updates and fixes to targeted focus effects ([#1870](https://github.com/EQEmu/Server/pull/1870))  
**KayenEQ** [Spells] Update SPA 238 SE_IllusionPersistence allow illusions to persist through deaths at higher AA ranks. ([#1884](https://github.com/EQEmu/Server/pull/1884))  
**KayenEQ** [Spells] Update to SPA 296 and 483 item and AA support ([#1857](https://github.com/EQEmu/Server/pull/1857))  
**KayenEQ** [Spells] Update to SPA 297 and 484 to support focus from AA and items. ([#1858](https://github.com/EQEmu/Server/pull/1858))  
**KayenEQ** [Spells] Update to SPA 440 SE_FinishingBlowMaxLevel limit value sets HP ratio for FB ([#1890](https://github.com/EQEmu/Server/pull/1890))  
**KayenEQ** [Spells] Update to SPA 58 SE_Levitate to support limit value ([#1875](https://github.com/EQEmu/Server/pull/1875))  
**KayenEQ** [Bug Fix] Root code error passing wrong variable into IsDetrimental check ([#1892](https://github.com/EQEmu/Server/pull/1892))  
**KayenEQ** [Bug Fix] Bard instrument modifiers applying to effects when it shouldn't under certain conditions. ([#1822](https://github.com/EQEmu/Server/pull/1822))  
**KayenEQ** [Spells] Update to SPA 89 SE_ChangeSize to support directly limit value that sets model size ([#1877](https://github.com/EQEmu/Server/pull/1877))  
**KayenEQ** [Bug Fix] validspell check to fix reported crash bug. ([#1895](https://github.com/EQEmu/Server/pull/1895))  
**Kinglykrab** [Bug Fix] Charm Break Invisibility Fix. ([#1855](https://github.com/EQEmu/Server/pull/1855))  
**Kinglykrab** [Bug Fix] Fix #guild rename, #killallnpcs, and #worldwide message errors. ([#1904](https://github.com/EQEmu/Server/pull/1904))  
**Kinglykrab** [Bug Fix] Fix Perl Croak for GetEnt() ([#1898](https://github.com/EQEmu/Server/pull/1898))  
**Kinglykrab** [Bug Fix] Fix possible crash in #givemoney. ([#1828](https://github.com/EQEmu/Server/pull/1828))  
**Kinglykrab** [Bug Fix] Fix possible crash with #killallnpcs. ([#1846](https://github.com/EQEmu/Server/pull/1846))  
**Kinglykrab** [Commands] Add #countitem Command. ([#1842](https://github.com/EQEmu/Server/pull/1842))  
**Kinglykrab** [Commands] Add #petitems Command. ([#1823](https://github.com/EQEmu/Server/pull/1823))  
**Kinglykrab** [Commands] Add #removeitem Command. ([#1847](https://github.com/EQEmu/Server/pull/1847))  
**Kinglykrab** [Commands] Add #setaltcurrency Command. ([#1850](https://github.com/EQEmu/Server/pull/1850))  
**Kinglykrab** [Commands] Add #setendurance Command. ([#1841](https://github.com/EQEmu/Server/pull/1841))  
**Kinglykrab** [Commands] Add #sethp Command. ([#1840](https://github.com/EQEmu/Server/pull/1840))  
**Kinglykrab** [Commands] Add #setmana Command. ([#1839](https://github.com/EQEmu/Server/pull/1839))  
**Kinglykrab** [Commands] Add #unmemspell and #unmemspells Commands. ([#1867](https://github.com/EQEmu/Server/pull/1867))  
**Kinglykrab** [Commands] Add #viewcurrencies Command. ([#1844](https://github.com/EQEmu/Server/pull/1844))  
**Kinglykrab** [Commands] Cleanup #bind Command. ([#1829](https://github.com/EQEmu/Server/pull/1829))  
**Kinglykrab** [Commands] Cleanup #disablerecipe Command. ([#1816](https://github.com/EQEmu/Server/pull/1816))  
**Kinglykrab** [Commands] Cleanup #enablerecipe Command. ([#1817](https://github.com/EQEmu/Server/pull/1817))  
**Kinglykrab** [Commands] Cleanup #flymode Command. ([#1845](https://github.com/EQEmu/Server/pull/1845))  
**Kinglykrab** [Commands] Cleanup #gender Command. ([#1832](https://github.com/EQEmu/Server/pull/1832))  
**Kinglykrab** [Commands] Cleanup #getplayerburiedcorpsecount Command. ([#1818](https://github.com/EQEmu/Server/pull/1818))  
**Kinglykrab** [Commands] Cleanup #gm Command. ([#1830](https://github.com/EQEmu/Server/pull/1830))  
**Kinglykrab** [Commands] Cleanup #gmspeed Command. ([#1831](https://github.com/EQEmu/Server/pull/1831))  
**Kinglykrab** [Commands] Cleanup #gmzone Command. ([#1836](https://github.com/EQEmu/Server/pull/1836))  
**Kinglykrab** [Commands] Cleanup #guild Command. ([#1880](https://github.com/EQEmu/Server/pull/1880))  
**Kinglykrab** [Commands] Cleanup #movechar Command. ([#1838](https://github.com/EQEmu/Server/pull/1838))  
**Kinglykrab** [Commands] Cleanup #mysql Command. ([#1837](https://github.com/EQEmu/Server/pull/1837))  
**Kinglykrab** [Commands] Cleanup #setadventurepoints Command. ([#1901](https://github.com/EQEmu/Server/pull/1901))  
**Kinglykrab** [Commands] Cleanup #setstartzone Command. ([#1853](https://github.com/EQEmu/Server/pull/1853))  
**Kinglykrab** [Commands] Cleanup #texture Command. ([#1835](https://github.com/EQEmu/Server/pull/1835))  
**Kinglykrab** [Commands] Cleanup #title Command. ([#1833](https://github.com/EQEmu/Server/pull/1833))  
**Kinglykrab** [Commands] Cleanup #titlesuffix Command. ([#1834](https://github.com/EQEmu/Server/pull/1834))  
**Kinglykrab** [Commands] Cleanup #wpinfo Command. ([#1866](https://github.com/EQEmu/Server/pull/1866))  
**Kinglykrab** (origin/master, master) [Bug Fix] Cleanup Perl croaks for Spire parser. ([#1908](https://github.com/EQEmu/Server/pull/1908))  
**Kinglykrab** [Quest API] Add GetBodyTypeName() to Perl/Lua. ([#1863](https://github.com/EQEmu/Server/pull/1863))  
**Kinglykrab** [Quest API] Add GetFactionName() to Perl/Lua. ([#1859](https://github.com/EQEmu/Server/pull/1859))  
**Kinglykrab** [Quest API] Add GetLanguageName() to Perl/Lua. ([#1860](https://github.com/EQEmu/Server/pull/1860))  
**Kinglykrab** [Quest API] Add GetLDoNThemeName() to Perl/Lua. ([#1861](https://github.com/EQEmu/Server/pull/1861))  
**Michael Cook (mackal)** Switch server to use new style ManaChange_Struct ([#1843](https://github.com/EQEmu/Server/pull/1843))  
**mmcgarvey** [Database] Escape reserved mysql keyword rank w/ backticks ([#1862](https://github.com/EQEmu/Server/pull/1862))  
**mmcgarvey** [Shared Tasks] Cross Zone Remove Fix ([#1740](https://github.com/EQEmu/Server/pull/1740))  
**mmcgarvey** [Skills] Make Tracking Skill Configurable ([#1784](https://github.com/EQEmu/Server/pull/1784))  
**mmcgarvey** [Spells] Instant Heals honor IgnoreSpellDmgLvlRestriction ([#1888](https://github.com/EQEmu/Server/pull/1888))  
**Natedog2012** [Bug Fix] SendSpellBarEnable sends correct slotid to fix spellbar on RoF2 ([#1848](https://github.com/EQEmu/Server/pull/1848))  
**Natedog2012** Client will give 1 second window to start casting at the end of DA effect but we interrupt it and need to allow spellbar active after ([#1894](https://github.com/EQEmu/Server/pull/1894))  
**Natedog2012** Do not set teleport doors to Open ([#1786](https://github.com/EQEmu/Server/pull/1786))  
**Natedog2012** [Quest API] Add ResetCastbarCooldownBySlot / ResetCastbarCooldownBySpellID / ResetAllCastbarCooldowns ([#1873](https://github.com/EQEmu/Server/pull/1873))  
**Natedog2012** Re-enable spellbar and reset Discipline timer when stopping casts in EVENT_CAST_BEGIN ([#1891](https://github.com/EQEmu/Server/pull/1891))  
**Natedog2012** RemoveAllAppearanceEffects sends all relevant data so NPC appearance doesn't get altered. ([#1864](https://github.com/EQEmu/Server/pull/1864))  
**Natedog2012** SetPetID after we assign the new NPC an ID ([#1851](https://github.com/EQEmu/Server/pull/1851))  
**Natedog2012** [Skills] RoF+ allows other classes to have feign death if set in skill_caps ([#1902](https://github.com/EQEmu/Server/pull/1902))  
**Paul Coene** [Bug Fix] Pick Lock was allowing skillups on doors above player skill ([#1815](https://github.com/EQEmu/Server/pull/1815))  
**Paul Coene** [Combat] Allow npcs/pets to kick vs opponents requiring magic weapons if wearing magic booties. ([#1868](https://github.com/EQEmu/Server/pull/1868))  
**Paul Coene** [Doors] Add new rule enabling classic "key on cursor" for pre keyring keys ([#1869](https://github.com/EQEmu/Server/pull/1869))
