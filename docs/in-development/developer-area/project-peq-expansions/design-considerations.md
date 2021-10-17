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
* Spells \(spells\_us.txt\) will **not** have expansion settings:
  * Various clients have spell count limitations
  * spells\_us.txt should be part of a standard "patch day" download when an expansion is launched
* Spell gems and spell book icons will **not** have expansion settings
* Models will **not** have expansion settings
  * Feel free to customize using eqclient.ini pre-Luclin models
  * Server operators can provide any models they wish
* Client Maps are enabled based on expansion settings through the client
  * The actual map content will not be controlled with expansion settings
  * Server operators can choose to provide any maps they wish
* Skill Caps???
* db\_str???
* base\_data???

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

* aa\_ability
  * Here we likely need a min/max expansion setting
  * Don't forget HT/LoH
* aa\_actions
* aa\_effects
* aa\_ranks
* aa\_rank\_effects
* aa\_rank\_prereqs
* aa\_required\_level\_cost
* aa\_timers
* altadv\_vars
  * Were AAs regrouped during our timeline?  We may need an expansion setting on this expansion setting.

### Account

Accounts have an expansion setting on them that opens race / class combos.  **This should probably be examined at part of this development cycle.**

* account
* account\_flags
* account\_ip
* account\_rewards
  * Veteran rewards will need min/max expansion settings
* sharedbank
  * Does this use the account expansion setting?

### Admin

* banned\_ips
* bot\_command\_settings
* bug\_reports
* bugs
* chatchannels
  * Chatchannels could make use of a rule setting to be turned on at the appropriate time for status &lt; X.
* class\_skill
* command\_settings
* db\_version
* discovered\_items
* eqtime
* eventlog
* fear\_hints
* gm\_ips
* hackers
* ip\_exemptions
* level\_exp\_mods
* logsys\_categories
* name\_filter
* perl\_event\_export\_settings
* petitions
* profanity\_list
* races
* reports
* saylink
* start\_zones
  * Requires min/max expansion settings
* starting\_items
  * Requires min/max expansion settings
* variables
* veteran\_reward\_templates

### Adventures

These tables are all related to LDON.  Presumably gating would occur with either zone access or npc\_types for Adventures.

* adventure\_details
* adventure\_members
* adventure\_stats
* adventure\_template
* adventure\_template\_entry
* adventure\_template\_entry\_flavor

### Alternate Currency

The actual currencies were introduced on some timeline.  Given that an item\_id is assigned and that each currency is a table entry, likely no min/max expansion setting is required.

* alternate\_currency

### Books

Text for items.  Item availability would be controlled by lootdrop and merchantlist entries; likely no min/max expansion setting is required.

* books

### Bots

Please review, Uleat.

### Buyers

No action needed.

* buyer

### Characters

* character\_activities
* character\_alternate\_abilities
* character\_alt\_currency
* character\_auras
* character\_bandolier
  * Is access to the bandolier currently observing expansion setting?
* character\_bind
* character\_buffs
* character\_corpses
* character\_corpse\_items
* character\_currency
* character\_data
* character\_disciplines
* character\_enabledtasks
* character\_inspect\_messages
* character\_item\_recast
* character\_languages
* character\_leadership\_abilities
* character\_material
* character\_memmed\_spells
* character\_pet\_buffs
* character\_pet\_info
* character\_pet\_inventory
* character\_potionbelt
  * Is access to the potion belt currently observing expansion setting?
* character\_skills
* character\_spells
* character\_tasks
* character\_task\_lockouts
* character\_tribute
* char\_create\_combinations
  * Uses existing expansion setting.
* char\_create\_point\_allocations
  * Requires min/max expansion settings
* char\_recipe\_list
* friends
* keyring
  * Is access to the key ring currently observing expansion setting?
* lfguild
* mail
* player\_titlesets

### Client Files

Please note the Client section above.

* base\_data
* db\_str
* skill\_caps
* spells\_new
  * NO expansion setting.

### Data Storage

No action needed.

* data\_buckets
* quest\_globals

### Doors

Certainly doors change throughout the various--keyitem\(s\) came and went, destination settings changed \(IE tox vs toxxulia PoK book\).

* doors
  * Requires min/max expansion setting

### Factions

Generally factions are gated by access to mobs.  Some factions were added to npc\_types as time went on, so under the NPC tables we'll want to look at min/max expansion settings.  Noudess should probably confirm that faction\_list\_mod didn't change, especially.

* client\_faction\_associations
* client\_faction\_names
* client\_server\_faction\_map
* custom\_faction\_mappings
* faction\_base\_data
* faction\_list
* faction\_list\_mod
* faction\_values

### Graveyards

Graveyards were gated by access to zones.  No min/max expansion setting required.

* graveyard

### Ground Spawns

Ground spawns, like fishing, certainly changed over time.  A min/max expansion setting is required.

* ground\_spawn
  * Requires min/max expansion setting

### Groups

Player table.  No min/max expansion setting required.

* group\_id
* group\_leaders

### Guilds

The guild hall was added during DoN, some guild management tools were added during LoY.  

* guilds
* guild\_bank
* guild\_ranks
* guild\_members
* guild\_relations

### Grids

Grids do not need expansion settings as the associated spawn2 point would gate the settings.

* grid
* grid\_entries

### Horses

Horses were introduced and evolved over time, but likely they would be gated through quest scripts and merchantlist entries.

* horses

### Instances

Player table.  No min/max expansion setting required.

* instance\_list
* instance\_list\_player

### Inventory

Player table.  No min/max expansion setting required.

* inventory
* inventory\_snapshots
* inventory\_versions

### Items

Items would be gated by lootdrop entries and merchantlists.  Some consideration could be given to nerfed items, but ultimately the server operator should probably make that decision.  

* items
* item\_tick

### Loginserver

World utility.  No min/max expansion setting required.

* login\_accounts
* login\_api\_tokens
* login\_server\_admins
* login\_server\_types
* login\_world\_servers

### Loot

Loot requires min/max expansion settings on multiple fronts:

* global\_loot
  * Requires min/max expansion settings
  * Example: Box of Abu-Kar
* lootdrop
  * Requires min/max expansion settings
  * Example: Manastone
* lootdrop\_entries
* loottable
  * Requires min/max expansion settings
  * Example: Manastone
* loottable\_entries

### Mercenaries

Mercenaries did evolve over time, however, are gated by zone/quest access.

* mercs
* merc\_armorinfo
* merc\_buffs
* merc\_inventory
* merc\_merchantlist\_entries
* merc\_merchant\_templates
* merc\_merchant\_template\_entries
* merc\_name\_types
* merc\_npc\_types
* merc\_spell\_lists
* merc\_spell\_list\_entries
* merc\_stance\_entries
* merc\_stats
* merc\_subtypes
* merc\_templates
* merc\_types
* merc\_weaponinfo

### Merchants

Merchantlists certainly changed over time.

* merchantlist
  * Requires min/max expansion settings
  * Example:  any spell scrolls
* merchantlist\_temp

### NPCs

NPCs require min/max expansion settings on at least factions.  

* npc\_emotes
* npc\_faction
  * Requires min/max expansion settings
  * Example: Primordial Malice
* npc\_faction\_entries
* npc\_scale\_global\_base
* npc\_spells
* npc\_spells\_effects
* npc\_spells\_effects\_entries
* npc\_spells\_entries
* npc\_types
* npc\_types\_metadata
* npc\_types\_tint
* proximities

### Objects

Objects require min/max expansion settings--zones changed over time.

* object
  * Requires min/max expansion settings
  * Example: LDON camps
* object\_contents

### Pets

Pets would be gated by spell/aa access.

* pets
* pets\_equipmentset
* pets\_equipmentset\_entries

### Query Server

World utility.  No min/max expansions required.

* qs\_merchant\_transaction\_record
* qs\_merchant\_transaction\_record\_entries
* qs\_player\_aa\_rate\_hourly
* qs\_player\_delete\_record
* qs\_player\_delete\_record\_entries
* qs\_player\_events
* qs\_player\_handin\_record
* qs\_player\_handin\_record\_entries
* qs\_player\_move\_record
* qs\_player\_move\_record\_entries
* qs\_player\_npc\_kill\_record
* qs\_player\_npc\_kill\_record\_entries
* qs\_player\_speech
* qs\_player\_trade\_record
* qs\_player\_trade\_record\_entries

### Raids

Player table.  No min/max expansion settings necessary.

* raid\_details
* raid\_leaders
* raid\_members

### Rules

Rules certainly changed over time--the question is whether or not this should be handled with various rule sets, or make use of a min/max expansion setting.

* rule\_sets
* rule\_values

### Spawns

Spawns certainly evolve over time.

* respawn\_times
  * This may require min/max expansion setting
* spawn2
  * Min/Max expansion setting required.
* spawnentry
* spawngroup
* spawn\_conditions
* spawn\_condition\_values
* spawn\_events

### Spells

Spells evolved over time.  Please see "Client" section above--spells will not get a min/max expansion setting.

* auras
* blocked\_spells
* damageshieldtypes
* spell\_buckets
* spell\_globals
* spells\_new

### Tasks

These would be gated by quest scripts.

* completed\_tasks
* goallists
* tasks
* tasksets
* task\_activities
* task\_replay\_groups

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
* tradeskill\_recipe
  * Requires min/max expansion setting
* tradeskill\_recipe\_entries

### Traps

Some traps did evolve over time.

* ldon\_trap\_entries
* ldon\_trap\_templates
* traps
  * Add min/max expansion settings

### Tributes

Tributes were added during GoD and evolved over time.  Obviously merchantlist and lootdrop expansion settings would restrict item access to be used as tribute.  Is this effected by world expansion setting?

* tributes
* tribute\_levels

### Zones

Zones evolved over time--access to zones as well as zone points.  Min/Max expansion settings are required.

* launcher
* launcher\_zones
* zone
  * Requires min/max expansion setting \(rather than using status / level\)
* zoneserver\_auth
* zone\_flags
* zone\_points
  * Requires min/max expansion settings
  * Example:  highpass vs. highpasshold
* zone\_server

## Combat

* [https://github.com/mackal/EQMechanics/wiki](https://github.com/mackal/EQMechanics/wiki)

```text
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

