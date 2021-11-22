# character_data

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Character Identifier |
| account_id | int | [Account Identifier](../../schema/account/account.md) |
| name | varchar | Name |
| last_name | varchar | Last Name |
| title | varchar | Title |
| suffix | varchar | Suffix |
| zone_id | int | [Zone Identifier](../../../../server/zones/zone-list) |
| zone_instance | int | Zone Instance Identifier |
| y | float | Y Coordinate |
| x | float | X Coordinate |
| z | float | Z Coordinate |
| heading | float | Heading Coordinate |
| gender | tinyint | [Gender](../../../../server/npc/genders) |
| race | smallint | [Race](../../../../server/npc/race-list) |
| class | tinyint | [Class](../../../../server/player/class-list) |
| level | int | Level |
| deity | int | [Deity](../../../../server/player/deity-list) |
| birthday | int | UNIX Timestamp of Birthday |
| last_login | int | UNIX Timestamp of Last Login |
| time_played | int | Time Played |
| level2 | tinyint | Level 2 |
| anon | tinyint | Anon: 0 = False, 1 = Anonymous, 2 = Roleplaying |
| gm | tinyint | GM: 0 = False, 1 = True |
| face | int | Face |
| hair_color | tinyint | Hair Color |
| hair_style | tinyint | Hair Style |
| beard | tinyint | Beard |
| beard_color | tinyint | Beard Color |
| eye_color_1 | tinyint | Eye Color 1 |
| eye_color_2 | tinyint | Eye Color 2 |
| drakkin_heritage | int | Drakkin Heritage |
| drakkin_tattoo | int | Drakkin Tattoo |
| drakkin_details | int | Drakkin Details |
| ability_time_seconds | tinyint | Ability Timer in Seconds |
| ability_number | tinyint | [Ability Number](../../schema/aas/aa_ability.md) |
| ability_time_minutes | tinyint | Ability Timer in Minutes |
| ability_time_hours | tinyint | Ability Timer in Hours |
| exp | int | Experience |
| aa_points_spent | int | AA Points Spent |
| aa_exp | int | AA Experience |
| aa_points | int | AA Points |
| group_leadership_exp | int | Group Leadership Experience |
| raid_leadership_exp | int | Raid Leadership Experience |
| group_leadership_points | int | Group Leadership Points |
| raid_leadership_points | int | Raid Leadership Points |
| points | int | Points |
| cur_hp | int | Health |
| mana | int | Mana |
| endurance | int | Endurance |
| intoxication | int | Intoxication |
| str | int | Strength |
| sta | int | Stamina |
| cha | int | Charisma |
| dex | int | Dexterity |
| int | int | Intelligence |
| agi | int | Agility |
| wis | int | Wisdom |
| zone_change_count | int | Zone Change Count |
| toxicity | int | Toxicity |
| hunger_level | int | Hunger Level |
| thirst_level | int | Thirst Level |
| ability_up | int | Ability Up |
| ldon_points_guk | int | LDoN Points - [Deepest Guk](../../../../server/zones/ldon-themes) |
| ldon_points_mir | int | LDoN Points - [Miragul's Menagerie](../../../../server/zones/ldon-themes) |
| ldon_points_mmc | int | LDoN Points - [Mistmoore Catacombs](../../../../server/zones/ldon-themes) |
| ldon_points_ruj | int | LDoN Points - [Rujarkian Hills](../../../../server/zones/ldon-themes) |
| ldon_points_tak | int | LDoN Points - [Takish-Hiz](../../../../server/zones/ldon-themes) |
| ldon_points_available | int | LDoN Points Available |
| tribute_time_remaining | int | Tribute Time Remaining |
| career_tribute_points | int | Career Tribute Points |
| tribute_points | int | Tribute Points |
| tribute_active | int | Tribute Active: 0 = False, 1 = True |
| pvp_status | tinyint | PVP Status: 0 = False, 1 = True |
| pvp_kills | int | PVP Kills |
| pvp_deaths | int | PVP Deaths |
| pvp_current_points | int | PVP Points |
| pvp_career_points | int | Career PVP Points |
| pvp_best_kill_streak | int | Best Kill Streak |
| pvp_worst_death_streak | int | Worse Death Streak |
| pvp_current_kill_streak | int | Current Kill Streak |
| pvp2 | int | PVP Status: 0 = False, 1 = True |
| pvp_type | int | PVP Type |
| show_helm | int | Show Helm: 0 = False, 1 = True |
| group_auto_consent | tinyint | Group Auto Consent: 0 = False, 1 = True |
| raid_auto_consent | tinyint | Raid Auto Consent: 0 = False, 1 = True |
| guild_auto_consent | tinyint | Guild Auto Consent: 0 = False, 1 = True |
| leadership_exp_on | tinyint | Leadership Experience On : 0 = False, 1 = True |
| RestTimer | int | Rest Timer |
| air_remaining | int | Air Remaining |
| autosplit_enabled | int | Autosplit Enabled: 0 = False, 1 = True |
| lfp | tinyint | Looking For Party: 0 = False, 1 = True |
| lfg | tinyint | Looking For Guild: 0 = False, 1 = True |
| mailkey | char | Mail Key |
| xtargets | tinyint | XTargets |
| firstlogon | tinyint | First Logon |
| e_aa_effects | int | Experience AA Effects |
| e_percent_to_aa | int | Experience Percentage to AA Points |
| e_expended_aa_spent | int | Expended AA Points Spent |
| aa_points_spent_old | int | AA Points Spent Old |
| aa_points_old | int | AA Points Old |
| e_last_invsnapshot | int | Last Inventory Snapshot |
| deleted_at | datetime |  |

