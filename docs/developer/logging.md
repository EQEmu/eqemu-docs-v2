# Logging for Developers

## Using Logs in Source

Using logs are very simple and are macro-driven. We use macros because we check for whether or not the log is enabled or has the level of logging enabled before even attempting to try to create strings and allocate them on the stack and waste precious resources discarding them right away. Preprocessor macros (Aliases below) allow us to do this cleanly and performant and makes for a nice developer experience.

Our logging aliases use CPP fmt library conveniently under the hood - so you don't have to think about what your data type bindings look like.

!!! example

    ```cpp
    LogDynamicZones("Caching [{}] dynamic zone(s) took [{}s]", dynamic_zone_cache.size(), bench.elapsed());
    ```

    ```cpp
    LogDynamicZones("Purging [{}] dynamic zone(s)", dz_ids.size());
    ```

    ```cpp
    LogLootDetail(
        "NPC [{}] does not meet loot_drop level requirements (min_level) level [{}] current [{}] for item [{}]",
        GetCleanName(),
        loot_drop.npc_min_level,
        GetLevel(),
        database.CreateItemLink(loot_drop.item_id)
    );
    ```

    ```cpp
    LogLoot(
        "[NPC::AddLootDrop] NPC [{}] Item ({}) [{}] charges [{}] chance [{}] trivial min/max [{}/{}] npc min/max [{}/{}]",
        GetName(),
        item2->ID,
        linker.GenerateLink(),
        loot_drop.item_charges,
        loot_drop.chance,
        loot_drop.trivial_min_level,
        loot_drop.trivial_max_level,
        loot_drop.npc_min_level,
        loot_drop.npc_max_level
    );
    ```

    ```cpp
    LogInfo("Client Files Export Utility");
    ```

## Logging Aliases

Logging aliases are maintained in **eqemu_logsys_log_aliases.h**

!!! info

    === "Aliases"
    
        All aliases have a "Detail" equivalent (Logging level 3)
    
        ```cpp
        LogInfo(message, ...);
        LogDebug(message, ...);
        LogAA(message, ...);
        LogAI(message, ...);
        LogAggro(message, ...);
        LogAttack(message, ...);
        LogPacketClientServer(message, ...);
        LogCombat(message, ...);
        LogCommands(message, ...);
        LogCrash(message, ...);
        LogDoors(message, ...);
        LogGroup(message, ...);
        LogGuilds(message, ...);
        LogInventory(message, ...);
        LogLauncher(message, ...);
        LogNetcode(message, ...);
        LogNormal(message, ...);
        LogObject(message, ...);
        LogPathing(message, ...);
        LogQSServer(message, ...);
        LogQuests(message, ...);
        LogRules(message, ...);
        LogSkills(message, ...);
        LogSpawns(message, ...);
        LogSpells(message, ...);
        LogTCPConnection(message, ...);
        LogTasks(message, ...);
        LogTradeskills(message, ...);
        LogTrading(message, ...);
        LogTribute(message, ...);
        LogMySQLError(message, ...);
        LogMySQLQuery(message, ...);
        LogMercenaries(message, ...);
        LogQuestDebug(message, ...);
        LogLoginserver(message, ...);
        LogClientLogin(message, ...);
        LogHeadlessClient(message, ...);
        LogHPUpdate(message, ...);
        LogFixZ(message, ...);
        LogFood(message, ...);
        LogTraps(message, ...);
        LogNPCRoamBox(message, ...);
        LogNPCScaling(message, ...);
        LogMobAppearance(message, ...);
        LogStatus(message, ...);
        LogAIScanClose(message, ...);
        LogAIYellForHelp(message, ...);
        LogAICastBeneficialClose(message, ...);
        LogAoeCast(message, ...);
        LogEntityManagement(message, ...);
        LogFlee(message, ...);
        LogAura(message, ...);
        LogHotReload(message, ...);
        LogMerchants(message, ...);
        LogZonePoints(message, ...);
        LogExpeditions(message, ...);
        LogDynamicZones(message, ...);
        LogCheatList(message, ...);
        LogClientList(message, ...);
        LogDiaWind(message, ...);
        LogHTTP(message, ...);
        ```

## Adding New Logging Categories

For development purposes, you may want to add a new category, this is very simple to do 

An example of a category being added can be seen at this commit: [How to add a category](https://github.com/EQEmu/EQEmu/commit/a46c0ee7e2dcf094c4b0e4d9cb91525443c19c5b). 

Once you've added the code to the mentioned sections, the next time world boots world will inject the logging categories if they don't exist in the table.

## Adding Default Values

In **eqemu_logsys.cpp** you can set default values for the log system initialization routine

[https://github.com/EQEmu/EQEmu/blob/master/common/eqemu_logsys.cpp](https://github.com/EQEmu/EQEmu/blob/master/common/eqemu_logsys.cpp)

=== "eqemu_logsys.cpp"

    ```cpp
    void EQEmuLogSys::LoadLogSettingsDefaults()
    {
        /**
         * Set Defaults
         */
        log_settings[Logs::WorldServer].log_to_console    = static_cast<uint8>(Logs::General);
        log_settings[Logs::ZoneServer].log_to_console     = static_cast<uint8>(Logs::General);
        log_settings[Logs::QSServer].log_to_console       = static_cast<uint8>(Logs::General);
        log_settings[Logs::UCSServer].log_to_console      = static_cast<uint8>(Logs::General);
        log_settings[Logs::Crash].log_to_console          = static_cast<uint8>(Logs::General);
        log_settings[Logs::MySQLError].log_to_console     = static_cast<uint8>(Logs::General);
        log_settings[Logs::Loginserver].log_to_console    = static_cast<uint8>(Logs::General);
        log_settings[Logs::HeadlessClient].log_to_console = static_cast<uint8>(Logs::General);
        log_settings[Logs::NPCScaling].log_to_gmsay       = static_cast<uint8>(Logs::General);
        log_settings[Logs::HotReload].log_to_gmsay        = static_cast<uint8>(Logs::General);
        log_settings[Logs::HotReload].log_to_console      = static_cast<uint8>(Logs::General);
        log_settings[Logs::Loot].log_to_gmsay             = static_cast<uint8>(Logs::General);
        log_settings[Logs::Scheduler].log_to_console      = static_cast<uint8>(Logs::General);
        log_settings[Logs::Cheat].log_to_console          = static_cast<uint8>(Logs::General);
        log_settings[Logs::HTTP].log_to_console           = static_cast<uint8>(Logs::General);
        log_settings[Logs::HTTP].log_to_gmsay             = static_cast<uint8>(Logs::General);
    ...
    ```

Similarly, when you add defaults here. World will inject these defaults into the database for users who pick up your new logging addition.

## Adding Aliases

When you create a new category, we later added support for logging aliases which you must add for ease of use.

=== "eqemu_logsys_log_aliases.h"

    ```cpp 
    #define LogSaylink(message, ...) do {\
        if (LogSys.log_settings[Logs::Saylink].is_category_enabled == 1)\
            OutF(LogSys, Logs::General, Logs::Saylink, __FILE__, __func__, __LINE__, message, ##__VA_ARGS__);\
    } while (0)
    
    #define LogSaylinkDetail(message, ...) do {\
        if (LogSys.log_settings[Logs::Saylink].is_category_enabled == 1)\
            OutF(LogSys, Logs::Detail, Logs::Saylink, __FILE__, __func__, __LINE__, message, ##__VA_ARGS__);\
    } while (0)
    ```

Using said alias then becomes very simple

!!! example

        ```cpp
        LogSaylink("Loaded [{}] saylinks into cache", saylinks.size());
        ```
