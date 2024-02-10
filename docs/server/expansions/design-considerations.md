---
description: >-
  This document describes the challenges and solutions to creating a
  dynamic-expansion supported PEQ baseline.  This is a work-in-progress and your
  comments, feedback, and ideas are certainly welcome.
---

# Design Considerations

## Client

* There will be no mechanism to close off clients based on server expansion settings.  
* Server Operators are welcome to customize client requirements.
* Spells (spells_us.txt) will **not** have expansion settings:
  * Various clients have spell count limitations
  * spells_us.txt should be part of a standard "patch day" download when an expansion is launched
* Spell gems and spell book icons will **not **have expansion settings
* Models will **not** have expansion settings
  * Feel free to customize using eqclient.ini pre-Luclin models
  * Server operators can provide any models they wish
* Client Maps are enabled based on expansion settings through the client
  * The actual map content will not be controlled with expansion settings
  * Server operators can choose to provide any maps they wish
* Skill Caps???
* db_str???
* base_data???

## Quest Scripts

* An expansion setting should be exposed to the quest API
* Scripts should utilize the expansion setting to "mask" portions of the script:

```perl
sub EVENT_SAY {
    if ($text=~/hail/i) {
        if (quest::is_the_ruins_of_kunark_enabled()) {
            quest::say("Something.");
        }
        elsif (quest::is_the_scars_of_velious_enabled()) {
            quest::say("Something different.");
        }
        elsif (quest::is_the_shadows_of_luclin_enabled()) {
            quest::say("Something else entirely.");
        }
        #:: Etc.
    }
}
```

## World

* Zone maps files will need to allow zone version support
  * .nav, .map, .wtr should exist for each **version** of the zone
  * All files should be standard on server install
  * Versions to be defined in the database
  * Examples:  nektulos, lavastorm
* Zone table customization to be covered below
  * Examples: Freeport, Highpass

## Database Tables

### AAs

Obviously AAs were added at each expansion.  As AAs make use of spells, it would be incumbent upon the server operator to include any out-of-era spells in their server files.

* aa_ability
  * Here we likely need a min/max expansion setting
  * Don't forget HT/LoH
* aa_actions
* aa_effects
* aa_ranks
* aa_rank_effects
* aa_rank_prereqs
* aa_required_level_cost
* aa_timers
* altadv_vars
  * Were AAs regrouped during our timeline?  We may need an expansion setting on this expansion setting.

### Account

Accounts have an expansion setting on them that opens race / class combos.  **This should probably be examined at part of this development cycle.**

* account
* account_flags
* account_ip
* account_rewards
  * Veteran rewards will need min/max expansion settings
* sharedbank
  * Does this use the account expansion setting?

### Admin

* banned_ips
* bot_command_settings
* bug_reports
* bugs
* chatchannels
  * Chatchannels could make use of a rule setting to be turned on at the appropriate time for status < X.
* class_skill
* command_settings
* db_version
* discovered_items
* eqtime
* eventlog
* fear_hints
* gm_ips
* hackers
* ip_exemptions
* level_exp_mods
* logsys_categories
* name_filter
* perl_event_export_settings
* petitions
* profanity_list
* races
* reports
* saylink
* start_zones
  * Requires min/max expansion settings
* starting_items
  * Requires min/max expansion settings
* variables
* veteran_reward_templates

### Adventures

These tables are all related to LDON.  Presumably gating would occur with either zone access or npc_types for Adventures.

* adventure_details
* adventure_members
* adventure_stats
* adventure_template
* adventure_template_entry
* adventure_template_entry_flavor

### Alternate Currency

The actual currencies were introduced on some timeline.  Given that an item_id is assigned and that each currency is a table entry, likely no min/max expansion setting is required.

* alternate_currency

### Books

Text for items.  Item availability would be controlled by lootdrop and merchantlist entries; likely no min/max expansion setting is required.

* books

### Bots

Please review, Uleat.

### Buyers

No action needed.

* buyer

### Characters

* character_activities
* character_alternate_abilities
* character_alt_currency
* character_auras
* character_bandolier
  * Is access to the bandolier currently observing expansion setting?
* character_bind
* character_buffs
* character_corpses
* character_corpse_items
* character_currency
* character_data
* character_disciplines
* character_enabledtasks
* character_inspect_messages
* character_item_recast
* character_languages
* character_leadership_abilities
* character_material
* character_memmed_spells
* character_pet_buffs
* character_pet_info
* character_pet_inventory
* character_potionbelt
  * Is access to the potion belt currently observing expansion setting?
* character_skills
* character_spells
* character_tasks
* character_task_lockouts
* character_tribute
* char_create_combinations
  * Uses existing expansion setting.
* char_create_point_allocations
  * Requires min/max expansion settings
* char_recipe_list
* friends
* keyring
  * Is access to the key ring currently observing expansion setting?
* lfguild
* mail
* player_titlesets

### Client Files

Please note the Client section above.

* base_data
* db_str
* skill_caps
* spells_new
  * NO expansion setting.

### Data Storage

No action needed.

* data_buckets
* quest_globals

### Doors

Certainly doors change throughout the various--keyitem(s) came and went, destination settings changed (IE tox vs toxxulia PoK book).

* doors
  * Requires min/max expansion setting

### Factions

Generally factions are gated by access to mobs.  Some factions were added to npc_types as time went on, so under the NPC tables we'll want to look at min/max expansion settings.  Noudess should probably confirm that faction_list_mod didn't change, especially.

* client_faction_associations
* client_faction_names
* client_server_faction_map
* custom_faction_mappings
* faction_base_data
* faction_list
* faction_list_mod
* faction_values

### Graveyards

Graveyards were gated by access to zones.  No min/max expansion setting required.

* graveyard

### Ground Spawns

Ground spawns, like fishing, certainly changed over time.  A min/max expansion setting is required.

* ground_spawn
  * Requires min/max expansion setting

### Groups

Player table.  No min/max expansion setting required.

* group_id
* group_leaders

### Guilds

The guild hall was added during DoN, some guild management tools were added during LoY.  

* guilds
* guild_bank
* guild_ranks
* guild_members
* guild_relations

### Grids

Grids do not need expansion settings as the associated spawn2 point would gate the settings.

* grid
* grid_entries

### Horses

Horses were introduced and evolved over time, but likely they would be gated through quest scripts and merchantlist entries.

* horses

### Instances

Player table.  No min/max expansion setting required.

* instance_list
* instance_list_player

### Inventory

Player table.  No min/max expansion setting required.

* inventory
* inventory_snapshots
* inventory_versions

### Items

Items would be gated by lootdrop entries and merchantlists.  Some consideration could be given to nerfed items, but ultimately the server operator should probably make that decision.  

* items
* item_tick

### Loginserver

World utility.  No min/max expansion setting required.

* login_accounts
* login_api_tokens
* login_server_admins
* login_server_types
* login_world_servers

### Loot

Loot requires min/max expansion settings on multiple fronts:

* global_loot
  * Requires min/max expansion settings
  * Example: Box of Abu-Kar
* lootdrop
  * Requires min/max expansion settings
  * Example: Manastone
* lootdrop_entries
* loottable
  * Requires min/max expansion settings
  * Example: Manastone
* loottable_entries

### Mercenaries

Mercenaries did evolve over time, however, are gated by zone/quest access.

* mercs
* merc_armorinfo
* merc_buffs
* merc_inventory
* merc_merchantlist_entries
* merc_merchant_templates
* merc_merchant_template_entries
* merc_name_types
* merc_npc_types
* merc_spell_lists
* merc_spell_list_entries
* merc_stance_entries
* merc_stats
* merc_subtypes
* merc_templates
* merc_types
* merc_weaponinfo

### Merchants

Merchantlists certainly changed over time.

* merchantlist
  * Requires min/max expansion settings
  * Example:  any spell scrolls
* merchantlist_temp

### NPCs

NPCs require min/max expansion settings on at least factions.  

* npc_emotes
* npc_faction
  * Requires min/max expansion settings
  * Example: Primordial Malice
* npc_faction_entries
* npc_scale_global_base
* npc_spells
* npc_spells_effects
* npc_spells_effects_entries
* npc_spells_entries
* npc_types
* npc_types_metadata
* npc_types_tint
* proximities

### Objects

Objects require min/max expansion settings--zones changed over time.

* object
  * Requires min/max expansion settings
  * Example: LDON camps
* object_contents

### Pets

Pets would be gated by spell/aa access.

* pets
* pets_equipmentset
* pets_equipmentset_entries

### Query Server

World utility.  No min/max expansions required.

* qs_merchant_transaction_record
* qs_merchant_transaction_record_entries
* qs_player_aa_rate_hourly
* qs_player_delete_record
* qs_player_delete_record_entries
* qs_player_events
* qs_player_handin_record
* qs_player_handin_record_entries
* qs_player_move_record
* qs_player_move_record_entries
* qs_player_npc_kill_record
* qs_player_npc_kill_record_entries
* qs_player_speech
* qs_player_trade_record
* qs_player_trade_record_entries

### Raids

Player table.  No min/max expansion settings necessary.

* raid_details
* raid_leaders
* raid_members

### Rules

Rules certainly changed over time--the question is whether or not this should be handled with various rule sets, or make use of a min/max expansion setting.

* rule_sets
* rule_values

### Spawns

Spawns certainly evolve over time.

* respawn_times
  * This may require min/max expansion setting
* spawn2
  * Min/Max expansion setting required.
* spawnentry
* spawngroup
* spawn_conditions
* spawn_condition_values
* spawn_events

### Spells

Spells evolved over time.  Please see "Client" section above--spells will not get a min/max expansion setting.

* auras
* blocked_spells
* damageshieldtypes
* spell_buckets
* spell_globals
* spells_new

### Tasks

These would be gated by quest scripts.

* completed_tasks
* goallists
* tasks
* tasksets
* task_activities
* task_replay_groups

### Timers

Player table.  No min/max expansion settings.

* timers

### Titles

Titles are gated by character AAs / Skills / Class / Gender.  Due to the inclusion of skills, titles would require min/max expansion since a tradeskill can surpass the lowest threshold, even in classic.

* titles
  * Requires min/max expansion setting.

### Tradeskills

Tradeskills evolved over time and will require min/max expansion settings.

* fishing
  * Requires min/max expansion setting
* forage
  * Requires min/max expansion setting
* tradeskill_recipe
  * Requires min/max expansion setting
* tradeskill_recipe_entries

### Traps

Some traps did evolve over time.

* ldon_trap_entries
* ldon_trap_templates
* traps
  * Add min/max expansion settings

### Tributes

Tributes were added during GoD and evolved over time.  Obviously merchantlist and lootdrop expansion settings would restrict item access to be used as tribute.  Is this effected by world expansion setting?

* tributes
* tribute_levels

### Zones

Zones evolved over time--access to zones as well as zone points.  Min/Max expansion settings are required.

* launcher
* launcher_zones
* zone
  * Requires min/max expansion setting (rather than using status / level)
* zoneserver_auth
* zone_flags
* zone_points
  * Requires min/max expansion settings
  * Example:  highpass vs. highpasshold
* zone_server

## Combat

* [https://github.com/mackal/EQMechanics/wiki](https://github.com/mackal/EQMechanics/wiki)

```
if IsClient():
  softcap = 350
  if GetClass() == WARRIOR and GetLevel() > 50:
    softcap = 430
  if (GetClass() == PALADIN or GetClass() == SHADOWKNIGHT or GetClass() == BARD) and GetLevel() > 50
    softcap = 403
  if (GetClass() == RANGER or GetClass() == ROGUE or GetClass() == MONK or GetClass() == BEASTLORD) and GetLevel() > 50
    softcap = 375

  cs_mod = 0
  cs_rank = GetAA(ComabtStability)
  if cs_rank == 1:
    cs_mod = 2
  elif cs_rank == 2:
    cs_mod = 5
  elif cs_rank == 3:
    cs_mod = 10

  softcap = softcap + (softcap * cs_mod / 100)

  if GetAA(PhysicalEnhancement):
    softcap = softcap + (2 * softcap / 100)

  if bonus <= softcap:
    return bonus

  if IsWarrior(): # melee, not priest/casters
    return (bonus - softcap) / 12 + softcap

  return softcap # hardcapped for everyone else
```
