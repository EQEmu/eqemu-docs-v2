# Database Backup Tool

World has a database backup utility wrapped as a CLI \(Command Line Interface\) utility

It's what powers the ProjectEQ archives micro-service [http://db.projecteq.net/](http://db.projecteq.net/) 

Dumper wrapper script [https://github.com/EQEmu/Server/blob/master/utils/sql/peq-dump/peq-dump.sh](https://github.com/EQEmu/Server/blob/master/utils/sql/peq-dump/peq-dump.sh) 

## Features

* Selectively backup certain types of tables \(content tables, player tables, login tables, state tables, system tables etc\)
* Does not table lock by default; making it easy to backup a server while it is online and players are running on it
* Supports table locking \(Can impact server playability\)
* Supports specifying a custom dump path
* Supports piping the dump output to console instead of writing directly to a file
* Supports option for compressing the backup \(only works with file output\) as long as you have a supported compression utility installed
  * 7Zip \(Linux, Windows\)
  * Tar \(Linux\)

### CLI Options

```text
eqemu@e5311a8e9505:~/server$ ./bin/world database:dump

Command

database:dump

Options

  --all
  --content-tables
  --login-tables
  --player-tables
  --state-tables
  --system-tables
  --query-serv-tables
  --table-structure-only
  --table-lock
  --dump-path=
  --dump-output-to-console
  --drop-table-syntax-only
  --compress
```

### Schema Source

The schema source for the table groupings can be found in the source at [https://github.com/EQEmu/Server/blob/master/common/database\_schema.h\#L27](https://github.com/EQEmu/Server/blob/master/common/database_schema.h#L27) 

## Example Output\(s\)

The below shows a few different types of examples

### Player Tables and Compression

```text
eqemu@e5311a8e9505:~/server$ ./bin/world database:dump --compress --player-tables
[World] [Info] MySQL installed [mysql  Ver 15.1 Distrib 10.3.22-MariaDB, for debian-linux-gnu (x86_64) using readline 5.2]
[World] [Info] Database [peq] Host [mariadb] Username [eqemu]
[World] [Info] Dumping Tables [account account_ip account_flags account_rewards adventure_details adventure_stats buyer char_recipe_list character_activities character_alt_currency character_alternate_abilities character_auras character_bandolier character_bind character_buffs character_corpse_items character_corpses character_currency character_data character_disciplines character_enabledtasks character_expedition_lockouts character_exp_modifiers character_inspect_messages character_item_recast character_languages character_leadership_abilities character_material character_memmed_spells character_pet_buffs character_pet_info character_pet_inventory character_potionbelt character_skills character_spells character_tasks character_tribute completed_tasks data_buckets discovered_items faction_values friends guild_bank guild_members guild_ranks guild_relations guilds instance_list_player inventory inventory_snapshots keyring mail petitions player_titlesets quest_globals sharedbank spell_buckets spell_globals timers titles trader trader_audit zone_flags ]
[World] [Info] Database dump created at [backups/peq-2021-05-17-player.sql]
[World] [Info] Compression requested... Compressing dump [backups/peq-2021-05-17-player.sql]
[World] [Info] Compressed dump created at [backups/peq-2021-05-17-player.tar.gz]
```

### Append Multiple Table Categories \(System, State, Login\)

You can specify as many or as little table groupings as possible

```text
eqemu@e5311a8e9505:~/server$ ./bin/world database:dump --compress --system-tables --state-tables --login-tables
[World] [Info] MySQL installed [mysql  Ver 15.1 Distrib 10.3.22-MariaDB, for debian-linux-gnu (x86_64) using readline 5.2]
[World] [Info] Database [peq] Host [mariadb] Username [eqemu]
[World] [Info] Dumping Tables [chatchannels command_settings content_flags db_str eqtime launcher launcher_zones spawn_condition_values spawn_events level_exp_mods logsys_categories name_filter perl_event_export_settings profanity_list rule_sets rule_values variables db_version inventory_versions adventure_members banned_ips bug_reports bugs dynamic_zones eventlog expedition_lockouts expedition_members expeditions gm_ips group_id group_leaders hackers ip_exemptions instance_list item_tick lfguild merchantlist_temp object_contents raid_details raid_leaders raid_members reports respawn_times saylink server_scheduled_events login_accounts login_api_tokens login_server_admins login_server_list_types login_world_servers ]
[World] [Info] Database dump created at [backups/peq-2021-05-17-system-state-login.sql]
[World] [Info] Compression requested... Compressing dump [backups/peq-2021-05-17-system-state-login.sql]
[World] [Info] Compressed dump created at [backups/peq-2021-05-17-system-state-login.tar.gz]
```

