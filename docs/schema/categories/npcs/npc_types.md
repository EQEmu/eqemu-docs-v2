# npc_types

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique NPC Type Identifier |
| name | text | Name |
| lastname | varchar | Last Name |
| level | tinyint | Level |
| race | smallint | [Race](../../../../categories/npc/race-list) |
| class | tinyint | [Class](../../../../categories/player/class-list) |
| bodytype | int | [Body Type](../../../../categories/npc/body-types) |
| hp | int | Health |
| mana | int | Mana |
| gender | tinyint | [Gender](../../../../categories/npc/genders) |
| texture | tinyint | [Texture](../../../../categories/npc/textures) |
| helmtexture | tinyint | [Helmet Texture](../../../../categories/npc/textures) |
| herosforgemodel | int | Hero's Forge Model |
| size | float | Size |
| hp_regen_rate | int | Health Regeneration |
| mana_regen_rate | int | Mana Regeneration |
| loottable_id | int | [Loottable Identifier](../../../schema/categories/loot/loottable.md) |
| merchant_id | int | [Merchant Identifier](../../../schema/categories/merchants/merchantlist.md) |
| alt_currency_id | int | [Alternate Currency Identifier](../../../schema/categories/alternate-currency/alternate_currency.md) |
| npc_spells_id | int | [NPC Spell Set Identifier](npc_spells.md) |
| npc_spells_effects_id | int | [NPC Spell Effects Identifier](npc_spells_effects.md) |
| npc_faction_id | int | [NPC Faction Identifier](../../../schema/categories/factions/faction_list.md) |
| adventure_template_id | int | [Adventure Template Identifier](../../../schema/categories/adventures/adventure_template.md) |
| trap_template | int | [Trap Template Identifier](../../../schema/categories/traps/traps.md) |
| mindmg | int | Minimum Damage |
| maxdmg | int | Maximum Damage |
| attack_count | smallint | Attack Count |
| npcspecialattks | varchar | NPC Special Attacks \(Deprecated\) |
| special_abilities | text | NPC Special Abilities |
| aggroradius | int | Aggro Radius |
| assistradius | int | Assist Radius |
| face | int | Face |
| luclin_hairstyle | int | Hair Style |
| luclin_haircolor | int | Hair Color |
| luclin_eyecolor | int | Eye Color 1 |
| luclin_eyecolor2 | int | Eye Color 2 |
| luclin_beardcolor | int | Beard Color |
| luclin_beard | int | Beard |
| drakkin_heritage | int | Drakkin Heritage |
| drakkin_tattoo | int | Drakkin Tattoo |
| drakkin_details | int | Drakkin Details |
| armortint_id | int | [Armor Tint Identifier](npc_types_tint.md) |
| armortint_red | tinyint | Armor Tint Red: 0 = None, 255 = Max |
| armortint_green | tinyint | Armor Tint Green: 0 = None, 255 = Max |
| armortint_blue | tinyint | Armor Tint Blue: 0 = None, 255 = Max |
| d_melee_texture1 | int | Primary Weapon Texture |
| d_melee_texture2 | int | Secondary Weapon Texture |
| ammo_idfile | varchar | Ammo Texture |
| prim_melee_type | tinyint | [Primary Melee Type](../../../../categories/player/skills) |
| sec_melee_type | tinyint | [Secondary Melee Type](../../../../categories/player/skills) |
| ranged_type | tinyint | [Ranged Type](../../../../categories/player/skills) |
| runspeed | float | Run Speed |
| MR | smallint | Magic Resistance |
| CR | smallint | Cold Resistance |
| DR | smallint | Disease Resistance |
| FR | smallint | Fire Resistance |
| PR | smallint | Poison Resistance |
| Corrup | smallint | Corruption Resistance |
| PhR | smallint | Physical Resistance |
| see_invis | smallint | See Invisible: 0 = False, 1 = True |
| see_invis_undead | smallint | See Invisible vs. Undread: 0 = False, 1 = True |
| qglobal | int | Quest Globals: 0 = Disabled, 1 = Enabled \(Deprecated\) |
| AC | smallint | Armor Class |
| npc_aggro | tinyint | NPC Aggro: 0 = False, 1 = True |
| spawn_limit | tinyint | Spawn Limit |
| attack_speed | float | Attack Speed: The lower the number, the faster the NPC hits. \(Deprecated\) |
| attack_delay | tinyint | Attack Delay: Delay between the attack arounds in 10ths of a second. |
| findable | tinyint | Findable: 0 = False, 1 = True |
| STR | mediumint | Strength |
| STA | mediumint | Stamina |
| DEX | mediumint | Dexterity |
| AGI | mediumint | Agility |
| _INT | mediumint | Intelligence |
| WIS | mediumint | Wisdom |
| CHA | mediumint | Charisma |
| see_hide | tinyint | See Hide: 0 = False, 1 = True |
| see_improved_hide | tinyint | See Improved Hide: 0 = False, 1 = True |
| trackable | tinyint | Trackable: 0 = False, 1 = True |
| isbot | tinyint | Is Bot: 0 = False, 1 = True |
| exclude | tinyint | Exclude: 0 = False, 1 = True |
| ATK | mediumint | Attack |
| Accuracy | mediumint | Accuracy |
| Avoidance | mediumint | Avoidance |
| slow_mitigation | smallint | Slow Mitigation |
| version | smallint | Version |
| maxlevel | tinyint | Maximum Level |
| scalerate | int | Scale Rate |
| private_corpse | tinyint | Private Corpse: 0 = False, 1 = True |
| unique_spawn_by_name | tinyint | Unique Spawn By Name: 0 = False, 1 = True |
| underwater | tinyint | Underwater: 0 = False, 1 = True |
| isquest | tinyint | Is Quest: 0 = False, 1 = True |
| emoteid | int | [Emote Identifier](npc_emotes.md) |
| spellscale | float | Spell Scale: 25 = 25%, 50 = 50%, 100 = 100% |
| healscale | float | Heal Scale: 25 = 25%, 50 = 50%, 100 = 100% |
| no_target_hotkey | tinyint | No Target Hotkey: 0 = False, 1 = True |
| raid_target | tinyint | Raid Target: 0 = False, 1 = True |
| armtexture | tinyint | [Arm Texture](../../../../categories/npc/textures) |
| bracertexture | tinyint | [Bracer Texture](../../../../categories/npc/textures) |
| handtexture | tinyint | [Hand Texture](../../../../categories/npc/textures) |
| legtexture | tinyint | [Leg Texture](../../../../categories/npc/textures) |
| feettexture | tinyint | [Feet Texture](../../../../categories/npc/textures) |
| light | tinyint | Light |
| walkspeed | tinyint | Walk Speed |
| peqid | int | PEQ Identifier |
| unique_ | tinyint | Unique |
| fixed | tinyint | Fixed |
| ignore_despawn | tinyint | Ignore Despawn: 0 = False, 1 = True |
| show_name | tinyint | Show Name: 0 = False, 1 = True |
| untargetable | tinyint | Untargetable: 0 = False, 1 = True |
| charm_ac | smallint | Charmed Armor Class |
| charm_min_dmg | int | Charmed Minimum Damage |
| charm_max_dmg | int | Charmed Maximum Damage |
| charm_attack_delay | tinyint | Charmed Attack Delay |
| charm_accuracy_rating | mediumint | Charmed Accuracy |
| charm_avoidance_rating | mediumint | Charmed Avoidance |
| charm_atk | mediumint | Charmed Attack |
| skip_global_loot | tinyint | Skip Global Loot: 0 = False, 1 = True |
| rare_spawn | tinyint | Rare Spawn: 0 = False, 1 = True |
| stuck_behavior | tinyint | Stuck Behavior |
| model | smallint | Model |
| flymode | tinyint | [Fly Mode](../../../../categories/npc/fly-modes) |
| always_aggro | tinyint | Aggro regardless of _int or level : 0 = False, 1 = True |
| exp_mod | int |  |

