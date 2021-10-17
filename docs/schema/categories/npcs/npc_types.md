# npc\_types

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique NPC Type Identifier |
| name | text | Name |
| lastname | varchar | Last Name |
| level | tinyint | Level |
| race | smallint | [Race](https://eqemu.gitbook.io/server/categories/npc/race-list) |
| class | tinyint | [Class](https://eqemu.gitbook.io/server/categories/player/class-list) |
| bodytype | int | [Body Type](https://eqemu.gitbook.io/server/categories/npc/body-types) |
| hp | int | Health |
| mana | int | Mana |
| gender | tinyint | [Gender](https://eqemu.gitbook.io/server/categories/npc/genders) |
| texture | tinyint | [Texture](https://eqemu.gitbook.io/server/categories/npc/textures) |
| helmtexture | tinyint | [Helmet Texture](https://eqemu.gitbook.io/server/categories/npc/textures) |
| herosforgemodel | int | Hero's Forge Model |
| size | float | Size |
| hp\_regen\_rate | int | Health Regeneration |
| mana\_regen\_rate | int | Mana Regeneration |
| loottable\_id | int | [Loottable Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/npcs/loottable.md) |
| merchant\_id | int | [Merchant Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/npcs/merchantlist.md) |
| alt\_currency\_id | int | [Alternate Currency Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/npcs/alternate_currency.md) |
| npc\_spells\_id | int | [NPC Spell Set Identifier](npc_spells.md) |
| npc\_spells\_effects\_id | int | [NPC Spell Effects Identifier](npc_spells_effects.md) |
| npc\_faction\_id | int | [NPC Faction Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/npcs/faction_list.md) |
| adventure\_template\_id | int | [Adventure Template Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/npcs/adventure_template.md) |
| trap\_template | int | [Trap Template Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/npcs/traps.md) |
| mindmg | int | Minimum Damage |
| maxdmg | int | Maximum Damage |
| attack\_count | smallint | Attack Count |
| npcspecialattks | varchar | NPC Special Attacks \(Deprecated\) |
| special\_abilities | text | NPC Special Abilities |
| aggroradius | int | Aggro Radius |
| assistradius | int | Assist Radius |
| face | int | Face |
| luclin\_hairstyle | int | Hair Style |
| luclin\_haircolor | int | Hair Color |
| luclin\_eyecolor | int | Eye Color 1 |
| luclin\_eyecolor2 | int | Eye Color 2 |
| luclin\_beardcolor | int | Beard Color |
| luclin\_beard | int | Beard |
| drakkin\_heritage | int | Drakkin Heritage |
| drakkin\_tattoo | int | Drakkin Tattoo |
| drakkin\_details | int | Drakkin Details |
| armortint\_id | int | [Armor Tint Identifier](npc_types_tint.md) |
| armortint\_red | tinyint | Armor Tint Red: 0 = None, 255 = Max |
| armortint\_green | tinyint | Armor Tint Green: 0 = None, 255 = Max |
| armortint\_blue | tinyint | Armor Tint Blue: 0 = None, 255 = Max |
| d\_melee\_texture1 | int | Primary Weapon Texture |
| d\_melee\_texture2 | int | Secondary Weapon Texture |
| ammo\_idfile | varchar | Ammo Texture |
| prim\_melee\_type | tinyint | [Primary Melee Type](https://eqemu.gitbook.io/server/categories/player/skills) |
| sec\_melee\_type | tinyint | [Secondary Melee Type](https://eqemu.gitbook.io/server/categories/player/skills) |
| ranged\_type | tinyint | [Ranged Type](https://eqemu.gitbook.io/server/categories/player/skills) |
| runspeed | float | Run Speed |
| MR | smallint | Magic Resistance |
| CR | smallint | Cold Resistance |
| DR | smallint | Disease Resistance |
| FR | smallint | Fire Resistance |
| PR | smallint | Poison Resistance |
| Corrup | smallint | Corruption Resistance |
| PhR | smallint | Physical Resistance |
| see\_invis | smallint | See Invisible: 0 = False, 1 = True |
| see\_invis\_undead | smallint | See Invisible vs. Undread: 0 = False, 1 = True |
| qglobal | int | Quest Globals: 0 = Disabled, 1 = Enabled \(Deprecated\) |
| AC | smallint | Armor Class |
| npc\_aggro | tinyint | NPC Aggro: 0 = False, 1 = True |
| spawn\_limit | tinyint | Spawn Limit |
| attack\_speed | float | Attack Speed: The lower the number, the faster the NPC hits. \(Deprecated\) |
| attack\_delay | tinyint | Attack Delay: Delay between the attack arounds in 10ths of a second. |
| findable | tinyint | Findable: 0 = False, 1 = True |
| STR | mediumint | Strength |
| STA | mediumint | Stamina |
| DEX | mediumint | Dexterity |
| AGI | mediumint | Agility |
| \_INT | mediumint | Intelligence |
| WIS | mediumint | Wisdom |
| CHA | mediumint | Charisma |
| see\_hide | tinyint | See Hide: 0 = False, 1 = True |
| see\_improved\_hide | tinyint | See Improved Hide: 0 = False, 1 = True |
| trackable | tinyint | Trackable: 0 = False, 1 = True |
| isbot | tinyint | Is Bot: 0 = False, 1 = True |
| exclude | tinyint | Exclude: 0 = False, 1 = True |
| ATK | mediumint | Attack |
| Accuracy | mediumint | Accuracy |
| Avoidance | mediumint | Avoidance |
| slow\_mitigation | smallint | Slow Mitigation |
| version | smallint | Version |
| maxlevel | tinyint | Maximum Level |
| scalerate | int | Scale Rate |
| private\_corpse | tinyint | Private Corpse: 0 = False, 1 = True |
| unique\_spawn\_by\_name | tinyint | Unique Spawn By Name: 0 = False, 1 = True |
| underwater | tinyint | Underwater: 0 = False, 1 = True |
| isquest | tinyint | Is Quest: 0 = False, 1 = True |
| emoteid | int | [Emote Identifier](npc_emotes.md) |
| spellscale | float | Spell Scale: 25 = 25%, 50 = 50%, 100 = 100% |
| healscale | float | Heal Scale: 25 = 25%, 50 = 50%, 100 = 100% |
| no\_target\_hotkey | tinyint | No Target Hotkey: 0 = False, 1 = True |
| raid\_target | tinyint | Raid Target: 0 = False, 1 = True |
| armtexture | tinyint | [Arm Texture](https://eqemu.gitbook.io/server/categories/npc/textures) |
| bracertexture | tinyint | [Bracer Texture](https://eqemu.gitbook.io/server/categories/npc/textures) |
| handtexture | tinyint | [Hand Texture](https://eqemu.gitbook.io/server/categories/npc/textures) |
| legtexture | tinyint | [Leg Texture](https://eqemu.gitbook.io/server/categories/npc/textures) |
| feettexture | tinyint | [Feet Texture](https://eqemu.gitbook.io/server/categories/npc/textures) |
| light | tinyint | Light |
| walkspeed | tinyint | Walk Speed |
| peqid | int | PEQ Identifier |
| unique\_ | tinyint | Unique |
| fixed | tinyint | Fixed |
| ignore\_despawn | tinyint | Ignore Despawn: 0 = False, 1 = True |
| show\_name | tinyint | Show Name: 0 = False, 1 = True |
| untargetable | tinyint | Untargetable: 0 = False, 1 = True |
| charm\_ac | smallint | Charmed Armor Class |
| charm\_min\_dmg | int | Charmed Minimum Damage |
| charm\_max\_dmg | int | Charmed Maximum Damage |
| charm\_attack\_delay | tinyint | Charmed Attack Delay |
| charm\_accuracy\_rating | mediumint | Charmed Accuracy |
| charm\_avoidance\_rating | mediumint | Charmed Avoidance |
| charm\_atk | mediumint | Charmed Attack |
| skip\_global\_loot | tinyint | Skip Global Loot: 0 = False, 1 = True |
| rare\_spawn | tinyint | Rare Spawn: 0 = False, 1 = True |
| stuck\_behavior | tinyint | Stuck Behavior |
| model | smallint | Model |
| flymode | tinyint | [Fly Mode](https://eqemu.gitbook.io/server/categories/npc/fly-modes) |
| always\_aggro | tinyint | Aggro regardless of \_int or level : 0 = False, 1 = True |

