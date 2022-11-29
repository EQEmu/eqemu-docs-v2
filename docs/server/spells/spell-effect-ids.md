# Spell Effect IDs

| ID | SPA | Spire SPA Effect Name | Description | Base | Limit | Max | Notes |
| :-- | :-- | :-- | :-- | :-- | :-- | :-- | :-- |
| 0 | SE_CurrentHP | Current HP | Modify targets hit points by amount, repeates every tic if in a buff. Heals and damage. | amount | target restriction id | max amount (use positive value) | Negative base value for damage | Positive base value for healing|
| 1 | SE_ArmorClass | AC | Modify AC by amount | amount | none | amount | |
| 2 | SE_ATK | ATK | Modify ATK by amount | amount | none | amount | |
| 3 | SE_MovementSpeed | Movement Rate | Modify movement speed by amount | amount | none | amount | |
| 4 | SE_STR | STR | Modify STR by amount | amount | none | amount | |
| 5 | SE_DEX | DEX | Modify DEX by amount | amount | none | amount | |
| 6 | SE_AGI | AGI | Modify AGI by amount | amount | none | amount | |
| 7 | SE_STA | STA | Modify STA by amount | amount | none | amount | |
| 8 | SE_INT | INT | Modify INT by amount | amount | none | amount | |
| 9 | SE_WIS | WIS | Modify WIS by amount | amount | none | amount | |
| 10 | SE_CHA | CHA | Modify CHA by amount | amount | none | amount | |
| 11 | SE_AttackSpeed | Attack Speed | Modify attack speed by percent | percent haste or slow | none | none | Base greater than 100 for haste (120 = 20 pct haste) |  Base less than 100 for slow (80 = 20 pct slow)|
| 12 | SE_Invisibility | Invisibility: Unstable | Apply invsibility that can drop before duration ends | invisibility level | none | none | Invisibility level determines what level of see invisible can detect it.|
| 13 | SE_SeeInvis | See Invisible | Apply see invisbile which will allow you to see invisible entities | see invisible level | none | none | See Invisible level determines what level of invisible it can see.|
| 14 | SE_WaterBreathing | Water Breathing | Immune to drowning | 1 (increase for stacking overwrite) | none | none | |
| 15 | SE_CurrentMana | Mana | Modify mana by amount, repeates every tic if in a buff. | amount | none | max amount (use positive value) | |
| 16 | SE_NPCFrenzy | NPC Frenzy | | | | | |
| 17 | SE_NPCAwareness | NPC Awareness | | | | | |
| 18 | SE_Lull | Pacify | Seen in harmony and lull spells. No coded functionality. | 1 | none | none | |
| 19 | SE_AddFaction | NPC Faction | Modify your faction with an NPC by amount | amount | none | none | |
| 20 | SE_Blind | Blindness | Remove vision from clients or cause NPC's to flee if not in melee range. | 1 (increase for stacking overwrite) | none | none | |
| 21 | SE_Stun | Stun | Interrupts spell casting and prevents target from doing any actions for duration | duration ms | pvp duration ms | max target level | |
| 22 | SE_Charm | Charm | Control another entity as your pet | Unknown set to 1 | none | max target level | |
| 23 | SE_Fear | Fear | Causes the entity to run away until duration ends | Unknown set to 1 | none | max target level | |
| 24 | SE_Stamina | Stamina Loss | Modify endurance upkeep by amount while using disciplines that drain endurance | amount | none | none | Positive value will reduce endurance upkeep | Negative value will increase endurance upkeep | Live no longer uses this spell effect for endurance regeneration.|
| 25 | SE_BindAffinity | Bind Affinity | Bind location for gate spell. | bind id (Set to 1, 2, or 3) | none | none | Bind id allows you set alternate bind points. Bind Point ID (1=Primary, 2=Secondary 3=Tertiary)|
| 26 | SE_Gate | Gate | Chance to teleport to bind location. | success chance | bind id  (2 or 3) | none | If limit is not set, you will gate to primary bind location. Bind Point ID (1=Primary, 2=Secondary 3=Tertiary)|
| 27 | SE_CancelMagic | Dispel Magic | Chance to removes detrimental and beneficial buffs | chance level modifier | none | none | Success chance is based on level difference of caster and caster of the buff, base value raises the casters level by the base amount.|
| 28 | SE_InvisVsUndead | Invisibility to Undead: Unstable | Apply invsibility verse undead that can drop before duration ends | invisibility level | none | none | Invisibility level determines what level of see invisible can detect it.|
| 29 | SE_InvisVsAnimals | Invisibility to Animals: Unstable | Apply invsibility verse animals that can drop before duration ends | invisibility level | none | none | Invisibility level determines what level of see invisible can detect it.|
| 30 | SE_ChangeFrenzyRad | NPC Aggro Radius | Set NPC aggro radius | amount | none | max target level | |
| 31 | SE_Mez | Mesmerize | Stuns target until duration ends or target takes damage. | 1 (increase for stacking overwrite) | none | max target level | Higher value of stacking type will always override the lower value. Used if you want one type of mez to overrite another.|
| 32 | SE_SummonItem | Summon Item | Summon an item. | item id | none | stack amount | |
| 33 | SE_SummonPet | Summon Pet | Summon a pet. | 1 | none | none | Set 'Teleport Zone' field to be the same as 'type' field in of the pet you want in the pets table.|
| 34 | SE_Confuse | Confuse | | | | | |
| 35 | SE_DiseaseCounter | Disease Counter | Determines potency of determental disease spells or potency of cures. | amount | none | none | Set to positive values for potency of detrimental spells | Set to negative value for potency of cure spells.|
| 36 | SE_PoisonCounter | Poison Counter | Determines potency of determental poison spells or potency of cures. | amount | none | none | Set to positive values for potency of detrimental spells | Set to negative value for potency of cure spells.|
| 37 | SE_DetectHostile | Detect Hostile | | | | | |
| 38 | SE_DetectMagic | Detect Magic | | | | | |
| 39 | SE_TwinCastBlocker | Twincast Blocker | Prevent this spell from being twincast. | 1 | none | none | |
| 40 | SE_DivineAura | Invulnerability | Invulnerable to spells and melee, you can not cast or melee while under this effect. | 1 | none | none | |
| 41 | SE_Destroy | Destroy | Instantly kill target. | 1 | none | none | |
| 42 | SE_ShadowStep | Shadow Step | Warps player a short distance in a random direction. | Unknown (Seen 1 to 50) | none | none | This effect is handled by the client. Changing the base value does not appear to have any affect.|
| 43 | SE_Berserk | Berserk | Sets entity as 'Berserk' allowing for chance to crippling blow regardless of hit points and class. | 1 | none | none | This is an unused live spell effect. Custom Spell Effect may be subject to change if live reuses the SPA.|
| 44 | SE_Lycanthropy | Stacking:  Delayed Heal Marker | Used as stacking checker in healing effects that use spell triggers. | stacking overwrite value | none | none | |
| 45 | SE_Vampirism | Vampirism | Heal for a percentage for your melee damage on target. | Percentage | none | none | This is an unused live spell effect. Custom Spell Effect may be subject to change if live reuses the SPA.|
| 46 | SE_ResistFire | Fire Resist | Modify fire resist by amount | amout | none | max amount | |
| 47 | SE_ResistCold | Cold Resist | Modify cold resist by amount | amout | none | max amount | |
| 48 | SE_ResistPoison | Poison Resist | Modify poison resist by amount | amout | none | max amount | |
| 49 | SE_ResistDisease | Disease Resist | Modify disease resist by amount | amout | none | max amount | |
| 50 | SE_ResistMagic | Magic Resist | Modify magic resist by amount | amout | none | max amount | |
| 51 | SE_DetectTraps | Detect Traps | | | | | |
| 52 | SE_SenseDead | Detect Undead | Point player in direction of nearest Undead NPC | 1 | none | none | |
| 53 | SE_SenseSummoned | Detect Summoned | Point player in direction of nearest Summoned NPC | 1 | none | none | |
| 54 | SE_SenseAnimals | Detect Animals | Point player in direction of nearest Animal NPC | 1 | none | none | |
| 55 | SE_Rune | Rune | Absorb all melee damage until a maxium amount of damage is taken then fades. | amount | none | max amount | |
| 56 | SE_TrueNorth | True North | Points player in north direction. | 1 | none | none | |
| 57 | SE_Levitate | Levitation | Float above ground and take no fall damage. | 1 (increase for stacking overwrite) | levitate while moving (Set to 1) | none | Levitate while moving is seen on Live 'Flying Mounts'|
| 58 | SE_Illusion | Illusion | Change appearance. | race id or gender id | texture id (see notes) | helmet id (see notes) | Illusions have complicated rules. See https://docs.eqemu.io/server/categories/spells/illusion-spell-guidelines|
| 59 | SE_DamageShield | Damage Shield | Take damage for the damage shield amount when meleeing a target with his effect. | amount | none | max amount | |
| 60 | SE_TransferItem | Transfer Item | | | | | |
| 61 | SE_Identify | Identify | Provides information about the item that your target is holding. | 1 | none | none | To use, hold item on cursor and cast spell on self or target. The 'lore' field from items table is displayed in chat window.|
| 62 | SE_ItemID | Item ID | | | | | |
| 63 | SE_WipeHateList | Memblur | Chance to wipe hate list of target, repeates every tic if in a buff. | percent chance | none | unknown | Actual chance to memory blur is much higher than the spells base value, caster level and CHA modifiers are added get the final calculated percent chance|
| 64 | SE_SpinTarget | Spin Stun | Spins and stuns target | duration ms | pvp duration ms | max target level | |
| 65 | SE_InfraVision | Infravision | Improved night vision | 1 | none | none | |
| 66 | SE_UltraVision | Ultravision | Vastly improved night vision | 1 | none | none | |
| 67 | SE_EyeOfZomm | Eye Of Zomm | Transfers your vision and control to a summoned temporary pet | 1 | none | none | Set 'Teleport Zone' field to be the same as 'type' field in of the pet you want in the pets table.|
| 68 | SE_ReclaimPet | Reclaim Energy | Kills your pet in exchange for mana. Returns 75 percent of pet spell actual mana cost. | 1 | none | none | |
| 69 | SE_TotalHP | Max HP | Modify maximum hit points by amount. | amount | none | max amount | |
| 70 | SE_CorpseBomb | Corpse Bomb | | | | | |
| 71 | SE_NecPet | Create Undead Pet | Summon a pet. | 1 | none | none | Set 'Teleport Zone' field to be the same as 'type' field in of the pet you want in the pets table.|
| 72 | SE_PreserveCorpse | Preserve Corpse | | | | | |
| 73 | SE_BindSight | Bind Sight | Transfers your vision to your target. | 1 | none | none | |
| 74 | SE_FeignDeath | Feign Death | Fall to the ground and have a chance to loss aggro from engaged NPCs. | success chance | none | none | |
| 75 | SE_VoiceGraft | Voice Graft | Speak through your pet | 1 | none | none | |
| 76 | SE_Sentinel | Sentinel | Creates a proximity zone where cast that alerts caster if NPC's or Players enter it. | 1 | none | none | Not implemented on EQEMU|
| 77 | SE_LocateCorpse | Locate Corpse | Points player in direction of their target's corpse. | 1 | none | none | |
| 78 | SE_AbsorbMagicAtt | Spell Rune | Absorb all spell damage until a maxium amount of damage is taken then fades. | amount | none | max amount | |
| 79 | SE_CurrentHPOnce | Current HP Once | Modify hit points by amount. Instant heals and direct damage. | amount | target restriction id | max amount | Negative base value for damage | Positive base value for healing|
| 80 | SE_EnchantLight | Enchant Light | | | | | |
| 81 | SE_Revive | Resurrect | Summon player to corpse and restore a percentage of experience. | percentage exp | none | none | |
| 82 | SE_SummonPC | Summon Player | Summon a player to casters location. | 1 | none | none | |
| 83 | SE_Teleport | Teleport | Teleport to another zone or location. | coordinate(x,y,z,h) | none | none | Set 'Teleport Zone' field to zone short name OR set to 'same' to teleport within same zone. To set all xyzh cooridinates, you have use the following. Use this effectid  only once in first effect slot . Cooridinates defined as effect_base_value1=x effect_base_value2=y effect_base_value3=z effect_base_value4=h|
| 84 | SE_TossUp | Gravity Flux | Toss up into the air. | distance (negative) | none | none | |
| 85 | SE_WeaponProc | Add Melee Proc | Add proc to melee | spellid | rate modifer | none | |
| 86 | SE_Harmony | Reaction Radius | Set NPC assist radius | amount | none | none | |
| 87 | SE_MagnifyVision | Magnification | Zoom players vision | amount | none | none | |
| 88 | SE_Succor | Evacuate | Teleport Group/Self to a defined location or to safe point in zone with a 2 percent fail rate | coordinate(x,y,z,h) | none | none | Set 'Teleport Zone' field to zone short name OR set to 'same' to evac within same zone. To set all xyzh cooridinates, you have use the following. Use this effectid only once in first effect slot . Cooridinates defined as effect_base_value1=x effect_base_value2=y effect_base_value3=z effect_base_value4=h|
| 89 | SE_ModelSize | Player Size | Modify targets size by percent or set to a specific model size. | percent shrink or grow | model size | unknown | Base greater than 100 for growth (120 = 20 pct growth) |  Base less than 100 for shrink (80 = 20 pct shrink) | To set to a specific model size, set base to 100 or 0 and then set limit to model size value.|
| 90 | SE_Cloak | Ignore Pet | Ignore pet | 1 | none | none | Not implemented on EQEMU|
| 91 | SE_SummonCorpse | Summon Corpse | Summon targets corpse to caster. | max target level | | | |
| 92 | SE_InstantHate | Hate | Add or remove a set amount of hate instantly from target. | amount | none | max amount | Positive value increases hate. | Negative value decreases hate.|
| 93 | SE_StopRain | Control Weather | Stops zone weather related rain. | 1 | none | none | |
| 94 | SE_NegateIfCombat | Make Fragile | Removes buff if player casts or does any combat skill. | 1 | none | none | |
| 95 | SE_Sacrifice | Sacrifice | Kills player and creates 'Essence Emerald', corpse can not be resurrected | 1 | none | none | |
| 96 | SE_Silence | Silence | Prevents spell casting. | 1 (increase for stacking overwrite) | none | unknown | |
| 97 | SE_ManaPool | Max Mana | Modify max mana pool by amount. | amount | none | max amount | |
| 98 | SE_AttackSpeed2 | Attack Speed: Does not exceed cap | Modify attack speed by percent. Stacks with other Attack Speed effects. Does need exceed haste cap. | percent haste or slow | none | none | Base greater than 100 for haste (120 = 20 pct haste) |  Base less than 100 for slow (80 = 20 pct slow)|
| 99 | SE_Root | Root | Immobilize target. | -10000 | none | none | |
| 100 | SE_HealOverTime | Heal Over Time | Heal over time. Stacks with other heal over time effects. | amount | target restriction id | max amount | |
| 101 | SE_CompleteHeal | Complete Heal: With Recast Blocker Buff | Heal for baseline of 7500 HP and apply a buff icon that blocks the same effect from taking hold until it fades. | heal amount multiplier | none | max heal amount multipler | |
| 102 | SE_Fearless | Fear Immunity | Immune to fear effect. | 1 | none | none | |
| 103 | SE_CallPet | Summon Pet | Summon Pet to owner. | 1 | none | none | |
| 104 | SE_Translocate | Translocate | Teleport your target to a specific location or bind. | coordinate(x,y,z,h) or Bind Point ID | none | none | Set 'Teleport Zone' field to zone short name OR set to 'same' to evac within same zone. To set all xyzh cooridinates, you have use the following. Use this effectid only once in first effect slot . If 'Teleport_Zone' field is not set, then will send to  bind point id, set base value to Bind Point ID (1=Primary, 2=Secondary 3=Tertiary)|
| 105 | SE_AntiGate | Inhibit Gate | Prevent target from casting gate | Seen 1 to 3 | none | none | Unclear what base value determines. May be related to Bind Point IDs.|
| 106 | SE_SummonBSTPet | Summon Warder | Summon a beastlord pet. | 1 | none | none | Set 'Teleport Zone' field to be the same as 'type' field in of the pet you want in the pets table.|
| 107 | SE_AlterNPCLevel | Alter NPC Level | Change NPC level by amount. Level returns to original level when buff fades. | amount | none | none | This is a no longer used on live. Custom Spell Effect may be subject to change if live reuses the SPA.|
| 108 | SE_Familiar | Summon Familiar | Summon a familiar. | 0 or 1 | none | none | Set 'Teleport Zone' field to be the same as 'type' field in of the pet you want in the pets table.|
| 109 | SE_SummonItemIntoBag | Summon into Bag | Summons item into a summoned bag.  | item id | none | none | To use this the first effectid must be SPA 32 SE_SummonItem and this must be a bag such as Phantom Satchel ID 17310. Then use this effectid to summon items that will go into that bag.|
| 110 | SE_IncreaseArchery | Increase Archery | | | | | |
| 111 | SE_ResistAll | All Resists | Modify all resists by amount. | amount | none | max amount | |
| 112 | SE_CastingLevel | Casting Level | Modify casting level by amount, this will determine fizzel rate. | amount | none | none | |
| 113 | SE_SummonHorse | Summon Mount | Summon a mount | 1 | none | none | Set 'Teleport Zone' field to be the same as 'filename' field in of the mount you want in the horses table.|
| 114 | SE_ChangeAggro | Hate Multiplier | Modify hate generated by percent | percent hate modifer | none | none | |
| 115 | SE_Hunger | Food | Resets hunger counter preventing hunger and thirst checks. | 1 | none | none | Set to positive values for potency of detrimental spells | Set to negative value for potency of cure spells.|
| 116 | SE_CurseCounter | Curse Counter | Determines potency of determental curse spells or potency of cures. | amount | | | |
| 117 | SE_MagicWeapon | Make Weapons Magical | Allows non-magic weapons to be considered magical | 1 | none | none | |
| 118 | SE_Amplification | Singing Amplification | Modifies bard 'singing modifier' by percent. | percent | none | none | Recasting this effect will cause it to focus itself, increasing its potency.|
| 119 | SE_AttackSpeed3 | Attack Speed: Overhaste | Modify attack speed by percent. Stacks with other Attack Speed effects. Can exceed haste cap. | percent haste or slow | none | none | Base greater than 100 for haste (120 = 20 pct haste) |  Base less than 100 for slow (80 = 20 pct slow)|
| 120 | SE_HealRate | Incoming Healing Effectiveness | Modify incoming heals by percentage. | percent healing | none | none | |
| 121 | SE_ReverseDS | Reverse Damage Shield | Heal for the reverse damage shields amount when meleeing a target with his effect. | amount | | | Negative value will cause the reverse damage shield to heal.|
| 122 | SE_ReduceSkill | Reduce Skill | pending | pending | pending | pending | not implemented|
| 123 | SE_Screech | Stacking:  Screech | If a buff has a Screech base value of +1 that buff will block any other buff that contains Screech with a base value of -1 | 1 or -1 | none | none | |
| 124 | SE_ImprovedDamage | Focus: Spell Damage | Modify outgoing spell damage by percentage. | min percent | none | max percent | Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.|
| 125 | SE_ImprovedHeal | Focus: Healing | Modify outgoing spell healing by percentage. | min percent | none | max percent | Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.|
| 126 | SE_SpellResistReduction | Focus: Spell Resist Rate | Modify outgoing spell resistance rate by percentage. | min percent | none | none | |
| 127 | SE_IncreaseSpellHaste | Focus: Spell Cast Time | Modify outgoing spell casting time by percentage. | percent | none | none | |
| 128 | SE_IncreaseSpellDuration | Focus: Spell Duration | Modify outgoing spell buff duration by percentage. | percent | none | none | |
| 129 | SE_IncreaseRange | Focus: Spell Range | Modify outgoing spell casting range percentage. | percent | none | none | |
| 130 | SE_SpellHateMod | Focus: Spell and Bash Hate | Modify outgoing spell and bash hate by percentage. | min percent | none | max percent | Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value. | Special case: For bash hate to be focused, need to add focus limit SPA 137 and set it to 999.|
| 131 | SE_ReduceReagentCost | Focus: Chance of Using Reagent | Modify outgoing spells chance to not consume reagent by percentage | min percent | none | max percent | Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.|
| 132 | SE_ReduceManaCost | Focus: Spell Mana Cost | Reduce outgoing spells mana cost by percentage | min percent | none | max percent | Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.|
| 133 | SE_FcStunTimeMod | Focus: Spell Stun Duration | Modify outgoing spell stun duration by percentage. | percent | none | none | |
| 134 | SE_LimitMaxLevel | Limit:  Max Level | Max level of spell that can be focused, if 'limit' is set then decrease effectiviness by the limit values percent per level over max level. | max level | effectiviness percent | none | |
| 135 | SE_LimitResist | Limit:  Resist | Resist Type(s) that a spell focus can require or exclude. | resist type | none | none | Include set value to positive | Exclude set value to negative|
| 136 | SE_LimitTarget | Limit:  Target | Target Type(s) that a spell focus can require or exclude. | target type | none | none | Include set value to positive | Exclude set value to negative|
| 137 | SE_LimitEffect | Limit:  Effect | Spell effect(s) that a spell focus can require or exclude. | spell effect ID | none | none | Include set value to positive | Exclude set value to negative|
| 138 | SE_LimitSpellType | Limit:  SpellType | Only allow focus spells that are Beneficial or Detrimental. | 0=Detrimental,  1=Beneficial | none | none | |
| 139 | SE_LimitSpell | Limit:  Spell | Specific spell id(s) that a spell focus can require or exclude. | spell ID | none | none | Include set value to positive | Exclude set value to negative|
| 140 | SE_LimitMinDur | Limit:  Min Duration | Mininum duration of spell that can be focused. | tics | none | none | Set duration in tics, 1 tick is 6 seconds of game time|
| 141 | SE_LimitInstant | Limit:  Instant spells only | Include or exclude if an instant cast spell can be focused. | 0=Exclude if Instant, 1=Allow only if Instant | none | none | |
| 142 | SE_LimitMinLevel | Limit:  Min Level | Mininum level of spell that can be focused. | level | none | none | |
| 143 | SE_LimitCastTimeMin | Limit:  Min Cast Time | Mininum cast time of spell that can be focused. | milliseconds | none | none | |
| 144 | SE_LimitCastTimeMax | Limit:  Max Cast Time | Max cast time of spell that can be focused | milliseconds | none | none | |
| 145 | SE_Teleport2 | Banish | Teleports targets to a defined location or to safe point in zone | coordinate(x,y,z,h) | none | none | Set 'Teleport Zone' field to zone short name  OR set to 'same' to teleport within same zone. To set all xyzh cooridinates, you have use the following. Use this effectid only once in first effect slot . Cooridinates defined as effect_base_value1=x effect_base_value2=y effect_base_value3=z effect_base_value4=h|
| 146 | SE_ElectricityResist | Portal Locations | pending | pending | pending | pending | |
| 147 | SE_PercentalHeal | Percent HP Heal | Modify targets hit points for percent value of the targets max HP. Heal or damage. | percentage | none | max amount of hit points | Negative base value for damage | Positive base value for healing|
| 148 | SE_StackingCommand_Block | Stacking:  Block | Effect is found on buff and is used to prevent another buff from taking hold if specific criteria is met. | spell effect id | none | Block if less than this value. | Formula - 201 = Slot to block | Exampe: AC buff with a value of 500. You want to block any other AC spell from going into Slot 1 that is less than 1000 AC. Slot 1 : Effect SE_AC(SPA 1) with Base=500, Slot 2:  Effect SE_StackingCommand_Block (SPA 148) with Base= SE_AC(SPA 1) Max = 1000 and Formula = 201 (apply to the first spell slot--remember that we use base 0)|
| 149 | SE_StackingCommand_Overwrite | Stacking:  Overwrite | Effect is found on buff and is used to overwrite another buff if specific criteria is met. | spell effect id | none | Overwrite if less than this value. | Formula - 201 = Slot to block | Exampe: AC buff with a value of 1200. You want to overwrite any other AC spell in Slot 1 that is less than 1000 AC. Slot 1 : Effect SE_AC(SPA 1) with Base=1200 , Slot 2:  Effect SE_StackingCommand_Overwrite (SPA 149) with Base= SE_AC(SPA 1) Max = 1000 and Formula = 201 (apply to the first spell slot--remember that we use base 0)|
| 150 | SE_DeathSave | Death Save | When this effect is applied to a target the owner of the buff is given a chance at receiving a heal when under 15 percent hit points. | 1 = 300 HP Healed,  2 = 8000 HP Healed | min target level to apply override heal amount | override heal amount | If max value is set as heal amount this value will be used instead as the heal amount if the owner of the buff is the mininum level specified in limit field. Chance to receive a heal is determined by the owner of the buffs Charisma. [Chance = ((Charisma * 3) +1) / 10) ] . SPA 277 gives a second chance to be healed if you fail.|
| 151 | SE_SuspendPet | Suspend Pet | Places a pet in temporary storage. | save type | none | none | Save Types, 0 = save pet with no buffs or equipment, 1 = save pet with no buffs or equipment, 2 = unknown.  SPA 308 allows for suspended pets to be resummoned after zoning.|
| 152 | SE_TemporaryPets | Summon a Pet Swarm | Summon temporary pet(s) that will fade after duration, will stack with regular pets. | amount of pets | none | duration seconds | Set 'Teleport Zone' field to be the same as 'type' field in of the pet you want in the pets table.|
| 153 | SE_BalanceHP | Balance Party HP | Balances groups HP with a percent modifier to the damage being distributed. | percent modifier | max HP taken from player | none | Positive base value increases damage being  distrubuted | Negative base value decreases the damage being distributed|
| 154 | SE_DispelDetrimental | Dispel Detrimental | Dispels detrimental buffs. | percent chance x 10 | none | none | Actual percent chance is calculated as base / 10|
| 155 | SE_SpellCritDmgIncrease | Spell Critical Damage | Modifies critical spell damage by percent | percent modifier | none | none | |
| 156 | SE_IllusionCopy | Illusion:  Target | Turns caster into mirror image of target. | Seen 0,1,30 | none | none | Unknown what base values represent.|
| 157 | SE_SpellDamageShield | Spell Damage Shield | Inflicts non-melee damage on caster of a spell. | amount damage shield (negative) | none | unknown | Spells must have 'feedbackable' field set to a value otherwise they will not be affected by spell damage shields.|
| 158 | SE_Reflect | Reflect Spell | Reflect casted detrimental spell back at caster. | percent chance | resist modifier | percent of base damage modifier | Spells must have 'reflectable' field set to a value otherwise they will not be reflected. Resist modifer, positive value reduces the resist rate, negative value increases the resist rate. Percent of base damage modifer, max greater than 100 for damage mod (120 = 20 pct increase in damage) |  max less than 100 for damage mod (80 = 20 pct decrease in damage)|
| 159 | SE_AllStats | All Stats | Modify all base stats by percent. (STR, AGI, DEX, WIS, INT, CHA) | amount | none | max amount | Effect currently handled entirely client side.|
| 160 | SE_MakeDrunk | Drunk | Intoxicate if tolerance under the base value. | amount | none | max amount | |
| 161 | SE_MitigateSpellDamage | Mitigate Spell Damage Rune | Mitigate incoming spell damage by percentage until rune fades. | percent mitigation | max damage absorbed per hit | rune amount | Special: If this effect is placed on item as worn effect or as an AA, it will provide stackable percent spell mitigation for the base value.|
| 162 | SE_MitigateMeleeDamage | Mitigate Melee Damage Rune | Mitigate incoming melee damage by percentage until rune fades. | percent mitigation | max damage absorbed per hit | rune amount | |
| 163 | SE_NegateAttacks | Absorb Damage | Complete or partially block incoming spell and melee damage | amount of blocked hits | none | max amount of damage blocked per hit | |
| 164 | SE_AppraiseLDonChest | Sense LDoN Chest | Attempt to sense the presence of a cursed trap on the targeted object. | 1 | none | skill check value | |
| 165 | SE_DisarmLDoNTrap | Disarm LDoN Trap | Attempt to disarm a cursed trap on the targeted object.  | 1 | none | skill check value | |
| 166 | SE_UnlockLDoNChest | Unlock LDoN Chest | Attempt to destroy any cursed lock present on the targeted object. | 1 | none | skill check value | |
| 167 | SE_PetPowerIncrease | Focus: Pet Power | Increase statistics and level of the pet when summoned. | power value | none | none | Pet power can be scaled automatically if 'petpower' field in pets table is set to 0 or -1, if the power field is set to anything it will look to find the cooresponding pet in the table with same power for that 'type'.|
| 168 | SE_MeleeMitigation | Defensive | Modify incoming melee damage by percent. | percent modifer | none | none | Negative base value decreases damage taken | Positive base value increases damage taken|
| 169 | SE_CriticalHitChance | Critical Melee Chance | Modify melee critical hit chance by skill. | percent modifer | skill type (-1 = all skill types) | none | |
| 170 | SE_SpellCritChance | Spell Critical Chance | Modifies spell critical chance by percent. | percent modifier | none | none | Must have a chance to perform critical hits in order to have a chance to crippling blow.|
| 171 | SE_CrippBlowChance | Crippling Blow Chance | Modify melee crippling blow chance | percent modifer | none | none | |
| 172 | SE_AvoidMeleeChance | Evasion | Modify chance to avoid melee. | percent modifer | none | none | |
| 173 | SE_RiposteChance | Riposte | Modify chance to riposte. | percent modifer | none | none | |
| 174 | SE_DodgeChance | Dodge | Modify chance to dodge. | percent modifer | none | none | |
| 175 | SE_ParryChance | Parry | Modify chance to parry. | percent modifer | none | none | |
| 176 | SE_DualWieldChance | Dual Wield | Modify dual weild chance. | percent modifer | none | none | |
| 177 | SE_DoubleAttackChance | Double Attack | Modify double attack chance. | percent modifer | none | none | Positive value will heal you | Negative value will damage you|
| 178 | SE_MeleeLifetap | Melee Lifetap | Heal for a percentage for your melee damage on target. | percentage | none | none | |
| 179 | SE_AllInstrumentMod | All Instrument Modifier | Set modifier value for all instrument and singing modifers that will be used if higher then the respective item modifers. | modifier percentage | none | none | |
| 180 | SE_ResistSpellChance | Resist Spell Chance | Modify chance to resist incoming spells by percent. | percent modifer | none | none | |
| 181 | SE_ResistFearChance | Resist Fear Spell Chance | Modify chance to resist incoming fear spells by percent. | percent modifer | none | none | |
| 182 | SE_HundredHands | Attack Delay Reducation | Modify attack delay by percent. | percent modifier | none | none | Negative value reduces delay, example -115 is calculated as a 15 percent reduction (-115/100). Positive value increases delay, 300 is calculated as a 30 percent increase in delay|
| 183 | SE_MeleeSkillCheck | Melee Skill Chance | Unknown intended effect. | percent chance | skill type (-1 = all skill types) | none | No longer used on live. It provides no benefits on eqemu.|
| 184 | SE_HitChance | Chance to Hit | Modify chance to hit by skill | percent modifer | skill type (-1 = all skill types) | none | |
| 185 | SE_DamageModifier | Skills Damage Modifier | Modify melee damage by skill. | percent modifer | skill type (-1 = all skill types) | none | |
| 186 | SE_MinDamageModifier | Skills Minimum Damage Modifier | Modify melee minimum damage by skill. | percent modifer | skill type (-1 = all skill types) | none | |
| 187 | SE_BalanceMana | Balance Party Mana | Balances groups mana with a percent modifier to the damage being distributed. | percent modifer | max mana taken from player | none | Positive base value increases damage being distributed | Negative base value decreases the damage being distributed|
| 188 | SE_IncreaseBlockChance | Chance to block | Modify chance to block melee. | percent modifer | none | none | |
| 189 | SE_CurrentEndurance | Endurance | Modify targets endurance by amount, repeates every tic if in a buff. | amount | none | max amount | |
| 190 | SE_EndurancePool | Max Endurance | Modify max endurance pool by amount. | amount | none | max amount | |
| 191 | SE_Amnesia | Amnesia | Prevents disciplines use. | 1 (increase for stacking overwrite) | none | none | |
| 192 | SE_Hate | Hate | Modify targets hate toward you by amount, repeates every tic if in a buff. | amount | none | max amount | |
| 193 | SE_SkillAttack | Skill Attack | Perform a combat round using a specific skill at a set weapon damage and chance to hit modifier. | weapon damage | chance to hit modifier | unknown | Skill used to perform combat round is determined by the 'skill' field in spells table.|
| 194 | SE_FadingMemories | Fade | Remove from hate lists and make invisible. Can set max level of NPCs that can be affected. | success chance | max level (ROF2 era) | max level (modern era) | Support for max level requires Rule (Spells, UseFadingMemoriesMaxLevel) to be true. If used from limit field, then it set as the level, ie. max level of 75 would use limit value of 75. If set from max field, max level 75 would use max value of 1075, if you want to set it so it checks a level range above the spell target then for it to only work on mobs 5 levels or below you set max value to 5.|
| 195 | SE_StunResist | Stun Resist | Modifies chance to resist a stun from a bash or kick by percent | percent modifier | none | none | |
| 196 | SE_StrikeThrough | Strikethrough | Modify chance to strikethrough by percent, bypassing an opponent's special defenses, such as dodge, block, parry, and riposte | percent modifier | none | none | |
| 197 | SE_SkillDamageTaken | Skill Damage Taken | Modify damage taken by percent from specific skill. | percent modifier | skill type (-1 = all skill types) | none | |
| 198 | SE_CurrentEnduranceOnce | Instant Endurance | Modify targets endurance by amount. | amount | none | max amount | Negative base value decreases damage taken | Positive base value increases damage taken|
| 199 | SE_Taunt | Taunt | Chance to taunt target and apply instant hate. | taunt success chance | amount hate added | | |
| 200 | SE_ProcChance | Worn Proc Chance | Modify worn weapon proc chance by percent. | percent modifier | none | none | |
| 201 | SE_RangedProc | Ranged Proc | Add proc to ranged attacks | spellid | rate modifer | none | |
| 202 | SE_IllusionOther | Project Illusion | Allows next casted self only illusion buff to be cast on a targeted player in group. | none | none | none | |
| 203 | SE_MassGroupBuff | Mass Group Buff | Allows next casted Group Buff to hit all players and pets within a large radius from caster at double the mana cost. | 1 | none | none | |
| 204 | SE_GroupFearImmunity | Group Fear Immunity | Provides immunity to fear for group.  | duration | none | none | Duration is calculated as base value * 10. Thus, value of 1 would be 10 seconds. This is not a buff and gives no icon.|
| 205 | SE_Rampage | AE Rampage | Perform a primary slot combat rounds on all creatures within a 40 foot radius. | number of attack rounds | max entities hit per round | aoe range override | On live base is always set to 1, if more than one attack is a spell it uses this SPA in multiple slots. Limit value can be used to set a max amount of entities able to be attacked per round.|
| 206 | SE_AETaunt | AE Taunt | Taunts all creatures within a 40 foot radius, placing you 'base values' points of hate higher than your opponents' previously most hated target. | added hate | none | aoe range override | |
| 207 | SE_FleshToBone | Flesh to Bone | Turns meat or body parts items into bone chips. | 1 | none | none | |
| 208 | SE_PurgePoison | Purge Poison | | | | | |
| 209 | SE_DispelBeneficial | Dispel Beneficial | Dispels beneficial buffs. | percent chance x 10 | none | none | Actual percent chance is calculated as base / 10|
| 210 | SE_PetShield | Pet Shield | Allows pet to use 'shield ability' on owner to reduce 50 percent of damage taken by owner for duration | Duration multiplier 1=12 seconds, 2=24 ect | mitigation on pet owner override | mitigation on pet overide | Special: limit and max values are not on live, they can be used to give mitigation penalties or bonuses to shielder or shielded.|
| 211 | SE_AEMelee | AE Melee | Perform a primary slot combat rounds on all creatures within a 40 foot radius for a duration. | Duration multiplier 1=12 seconds, 2=24 ect | none | none | Only implemented for clients.|
| 212 | SE_FrenziedDevastation | Frenzied Devastation | Increase spell critical chance and while present all direct damage spells cost double the amount of mana. | 1 | chance modifier | none | Live no longer uses the effect in this way. It is now a focus effect.|
| 213 | SE_PetMaxHP | Pet Max HP | Modifies your pets maximum HP by percent. | percent modifier | none | none | This effect goes on the pet owner and then the benefit is applied to the pet.|
| 214 | SE_MaxHPChange | Change Max HP | Modify your maximum HP by percent. | percent modifier | none | none | Base value is divided by 100 to get actual percentage. Example, for 10 percent max HP increase, base value should be 1000.|
| 215 | SE_PetAvoidance | Pet Avoidance | Modifies your pets chance to avoid melee by percent. | percent modifier | none | none | This effect goes on the pet owner and then the benefit is applied to the pet.|
| 216 | SE_Accuracy | Accuracy | Modify your chance to hit by modifying accuracy by amount. | amount accuracy | skill type (-1 = all skill types) | none | AA version of this is not skill limited. Differs from SPA 184, which is a multiplier of your total accuracy.|
| 217 | SE_HeadShot | Headshot | Grants your archery attacks a chance to deal extra damage to humanoid NPC targets. | percent chance | damage amount | none | Used with SPA 346 which limits headshot by level and adds a bonus chance.|
| 218 | SE_PetCriticalHit | Pet Crit Melee | Modifies your pets chance to perform a critical melee hit by percent. | percent modifier | none | none | This effect goes on the pet owner and then the benefit is applied to the pet.|
| 219 | SE_SlayUndead | Slay Undead | Chance to do increased damage verse undead. | damage percent modifier | chance | none | Actual chance will be your limit value / 10. Example a 14 percent chance would require limit value of 140. Damage modifier baseline is 100, Base greater than 100 for increased damage  (120 = 20 pct damage increase) |  Base less than 100 for reduced damage (80 = 20 pct damage reduction).|
| 220 | SE_SkillDamageAmount | Skill Damage Bonus | Add a flat amount of damage when a specific melee skill is used. | amount | skill type (-1 = all skill types) | none | |
| 221 | SE_Packrat | Reduce Weight | Modify weight of your inventory by percent. | percent modifier | none | none | |
| 222 | SE_BlockBehind | Block Behind | Modify chance to block from behind. | percent modifer | none | none | |
| 223 | SE_DoubleRiposte | Double Riposte | Chance to do an additional riposte attack after a successful riposte. | percent chance | none | none | No longer used on live.|
| 224 | SE_GiveDoubleRiposte | Additional Riposte | Chance to do an additional riposte attack, or skill based riposte like Flying Kick. | percent chance | skill type | none | If limit value is set you can riposte using a specific special attack skill, like 'flying kick'. You can not have multiple skills that can riposte, thus limited to use in only one effect.|
| 225 | SE_GiveDoubleAttack | Double Attack | Allows any class to double attack at a set percent chance or modifty chance if class can innately double attack. | percent chance or modifer | none | none | |
| 226 | SE_TwoHandBash | Two Hand bash | Bash with a two handed weapon. | 1 | none | none | Handled client side.|
| 227 | SE_ReduceSkillTimer | Base Refresh Timer | Reduce base refresh timer of skill | time seconds (positive) | skill type | none | |
| 228 | SE_ReduceFallDamage | Reduce Fall Dmg | Reduce the damage that you take from falling by percent. | percent modifer | none | none | |
| 229 | SE_PersistantCasting | Cast Through Stun | Chance to continue casting while stunned. | percent chance | none | none | |
| 230 | SE_ExtendedShielding | Increase Shield Distance | Modify the range of your /shield ability by an amount of distance | distance | none | none | |
| 231 | SE_StunBashChance | Stun Bash Chance | Modify chance to land a stun using bash skill by percent | percent modifer | none | none | |
| 232 | SE_DivineSave | Divine Save | Chance to return to life with the blessing Touch of the Divine when you would otherwise die.   | percent chance | spellid | none | This effect triggers upon death, where base value gives you percent chance to cast Touch of the Divine which is an Invulnerability, heal, HoT and purify effect. Limit value can be used to add an additional spell being applied on death, usually this is a heal.|
| 233 | SE_Metabolism | Metabolism | Modifies food and drink consumption rates. | percent modifer | none | none | Positive value decreases consumption rate | Negative value increase consumption rate.|
| 234 | SE_ReduceApplyPoisonTime | Poison Mastery | Decrease poison application time. | time | none | none | Reducation time calculated as base /1000. Example, 2.5 second reduction would be a base value of 2500.|
| 235 | SE_ChannelChanceSpells | Focus Channelling | Modify chance to channel a spell and avoid being interupted by percent. | percent modifer | none | none | No longer used on live.|
| 236 | SE_FreePet | Free Pet | | | | | |
| 237 | SE_GivePetGroupTarget | Pet Affinity | Allows summoned pets to receive group buffs. | 1 | none | none | This effect goes on the pet owner and then the benefit is applied to the pet.|
| 238 | SE_IllusionPersistence | Permanent Illusion | Illusions will persist through zoning and last up to 16 hours. If base value increased to two, will persist through death. | 1 or 2 | none | none | 1=Persist through zoning, 2=Persist through death|
| 239 | SE_FeignedCastOnChance | Feign Death Through Spell Hit  | Chance to feign death through spell hit. | percent chance | none | none | If spell is resisted your chance is multipled by two.|
| 240 | SE_StringUnbreakable | String Unbreakable | | | | | |
| 241 | SE_ImprovedReclaimEnergy | Improve Reclaim Energy | Modify amount of mana returned from from reclaim energy by percent. | percent modifier | none | none | |
| 242 | SE_IncreaseChanceMemwipe | Increase Chance Memwipe | Modify the chance to wipe hate with memory blurr by percent. | percent modifier | none | none | Actual chance to memory blur is much higher than the memory blurs spells base value, caster level and CHA modifiers are added get the final calculated percent chance. This effect modifiers that final percent chance.|
| 243 | SE_CharmBreakChance | Charm Break Chance | Modify chance of charm breaking early by percentage. | percent modifier | none | none | |
| 244 | SE_RootBreakChance | Root Break Chance | Modify chance of the casters root being broken by percentage. | percent modifier | none | none | Modifies the base line root break chance. The benefit is given to any player casting on that NPC with the root, opposed to only the caster of the root.|
| 245 | SE_TrapCircumvention | Trap Circumvention | Decreases the chance that you will set off a trap when opening a chest or other similar container by percentage. | percent modifier | none | none | |
| 246 | SE_SetBreathLevel | Lung Capacity | Modify the amount of air you can hold in your lungs by percent | percent modifier | none | none | Should work client side. No server side support.|
| 247 | SE_RaiseSkillCap | Increase SkillCap | Increase skill cap. | amount | skill type | none | |
| 248 | SE_SecondaryForte | Extra Specialization | Gives you a second specialize skill that can go past 50 to 100. | 100 | none | none | Changing base value will not alter this effect.|
| 249 | SE_SecondaryDmgInc | Offhand Min Damage Bonus | Modify offhand weapons minimum damage bonus by percentage. | percent modifier | none | none | |
| 250 | SE_SpellProcChance | Spell Proc Chance | Modify proc chance of combat procs gained by spells or AAs by percent. | percent modifier | none | none | |
| 251 | SE_ConsumeProjectile | Endless Quiver | Chance not to consume a projectile when usimg archery or throwing skill. | percent chance | none | none | |
| 252 | SE_FrontalBackstabChance | Backstab from Front | Chance to perform a full backstab while in front of the target. | percent chance | none | none | |
| 253 | SE_FrontalBackstabMinDmg | Chaotic Stab | Allow a frontal backstab for mininum damage. | 1 | none | none | |
| 254 | SE_Blank | No Spell | Default value for an unused effect slot. | none | none | none | Do not replace this effect.|
| 255 | SE_ShieldDuration | Shielding Duration | Extends the duration of your /shield ability. | seconds | none | none | |
| 256 | SE_ShroudofStealth | Shroud Of Stealth | Rogue improved invsible. | 1 | none | none | |
| 257 | SE_PetDiscipline | Give Pet Hold | Gives pet command, pet /hold | 1 | none | none | SPA 267 with a limit value of 15 is required now to obtain pet /hold. |
| 258 | SE_TripleBackstab | Triple Backstab | Chance to perform a triple backstab. | percent chance | none | none | |
| 259 | SE_CombatStability | AC Softcap Limit | Modify AC soft cap by amount. | amount | none | none | |
| 260 | SE_AddSingingMod | Instrument Modifier | Set modifier value a specific instrument/singing skills that will be used if higher then the respective item modifier for that skill. | percent modifier | Item Type ID | none | Item Type IDs, 23=Woodwind, 24=Strings, 25=Brass, 26=Percussions, 50=Singing, 51=All instruments|
| 261 | SE_SongModCap | Song Cap | Modify max song modifier cap by amount. | amount | none | none | Song cap is set in spells table as field 'song_cap'. Not used on live.|
| 262 | SE_RaiseStatCap | Raise Stat Cap | Modify stat cap type by amount. | amount | stat type id | none | Stat type id, STR=0, STA=1, AGI=2, DEX=3, WIS=4, INT=5, CHA=6, MR=7, CR=8, FR=9, PR=10, DR=11, COR=12|
| 263 | SE_TradeSkillMastery | Tradeskill Masteries | Allows you to raise additional standard tradeskill (Baking, Blacksmithing, Brewing, Fletching, Jewelcraft, Pottery, or Tailoring) from its initial Specialization cap of 200 up to 250. | amount of skills that can be raised (max=6) | none | none | |
| 264 | SE_HastenedAASkill | Reduce AA Timer | Reduces reuse time on AA skills. | reducation amount seconds | aa id | none | This can be only set as an AA ability. |
| 265 | SE_MasteryofPast | No Fizzle | Makes it impossible for you to fizzle spells at our below the spell level specified by the base value. | level | none | none | |
| 266 | SE_ExtraAttackChance | Add Extra Attack: 2H Primary | Gives your double attacks a percent chance to perform an extra attack with 2-handed primary weapon. | percent chance | number of attacks | none | |
| 267 | SE_AddPetCommand | Add Pet Commands | Enables multilpe different pet commands based on limit value. | 1 | pet command type | none | Full list of command types found in common.h|
| 268 | SE_ReduceTradeskillFail | Tradeskill Failure Rate | Modify chance to fail with given tradeskill by a percent. | chance modifier | none | none | |
| 269 | SE_MaxBindWound | Bandage Percent Limit | Modify max HP you can bind wound by percent. | percent modifier | none | none | |
| 270 | SE_BardSongRange | Bard Song Range | Modify range of beneficial bard songs by percent. | percent modifier | none | none | |
| 271 | SE_BaseMovementSpeed | Base Run Speed | Modify base rune speed by percentage. | percent modifier | none | none | Does not stack with run speed modifiers.|
| 272 | SE_CastingLevel2 | Casting Level | Modify effective casting level by amount. | level amount | none | none | Live decription: This affects, spells that get stronger or last longer based on your level, stacking priority on targets, likelihood that spells that dispel effects will succeed, likelihood that spells that cure blindness will succeed, likelihood that spells that sense, disarm, or pick locked traps will succeed.|
| 273 | SE_CriticalDoTChance | Critical DoT | Modifies chance to critical damage over time spells. | percent chance | none | none | |
| 274 | SE_CriticalHealChance | Critical Heal | Modify chance to critical heal spells. | percent chance | none | none | |
| 275 | SE_CriticalMend | Critical Mend | Modify chance to critical the monks mend ability. | percent chance | none | none | |
| 276 | SE_Ambidexterity | Dual Wield Skill Amount | Modify your chance to successfully dual wield by increasing dual weild skill by amount. | amount | none | none | |
| 277 | SE_UnfailingDivinity | Extra DI Chance | Gives second chance for a death save to fire and if successful gives a modified heal amount. | heal modifier percent | none | none | This works with Death Save SPA 150.|
| 278 | SE_FinishingBlow | Finishing Blow | Grants melee attacks a chance to deal massive damage to NPC targets with 10% or less health. | percent chance | damage amount | none | Actual chance is calculated as base value / 10. Example for 50 percent chance, set base to 500. Use with SPA 440 to set max level of NPC that can be affected by finishing blow.|
| 279 | SE_Flurry | Flurry Chance | Chance to do a melee flurry, performing two extra attacks. | percent chance | none | none | |
| 280 | SE_PetFlurry | Pet Flurry Chance | Chance for pet to perform a melee flurry. | percent chance | none | none | This effect goes on the pet owner and then the benefit is applied to the pet.|
| 281 | SE_FeignedMinion | Give Pet Feign | | | | | |
| 282 | SE_ImprovedBindWound | Bandage Amount | Modify bind wound healing amount by percent. | percent modifier | none | none | |
| 283 | SE_DoubleSpecialAttack | Special Attack Chain | Chance to perform second special skill attack as monk.  | percent chance | none | none | |
| 284 | SE_LoHSetHeal | LoH Set Heal | | | | | |
| 285 | SE_NimbleEvasion | NoMove Check Sneak | | | | | |
| 286 | SE_FcDamageAmt | Focus: Spell Damage Amount | Modify spell damage by a flat amount. | amount | none | none | |
| 287 | SE_SpellDurationIncByTic | Focus: Buff Duration by Tics | Modify spell buff duration by tics.  | duration tics | none | none | 1 tic = 6 seconds, set base in tics|
| 288 | SE_SkillAttackProc | Add Proc From Skill Attack | Chance to cast a spell when a skill attack is performed. | chance percent | skill type | none | Chance is calculated as base value / 10, example 20 percent chance would be a value of 200. For AA's the proc spell ID is the 'spell'  field used in the aa_ranks table. |
| 289 | SE_CastOnFadeEffect | Cast Spell On Fade | Cast a spell only if buff containing this effect fades after the full duration.   | spellid | none | none | Typically seen on spells that can be cured.|
| 290 | SE_IncreaseRunSpeedCap | Movement Cap | Increase movement speed cap. | amount | none | none | |
| 291 | SE_Purify | Purify | Remove up specified amount of detiremental spells, excluding charm, fear, resurrection, and revival sickness. | amount spells removed | none | none | |
| 292 | SE_StrikeThrough2 | Strikethrough (v292) | Modify chance to strikethrough by percent, bypassing an opponent's special defenses, such as dodge, block, parry, and riposte | percent modifier | none | none | |
| 293 | SE_FrontalStunResist | Frontal Stun Resist | Chance to resist a stun from a bash or kick if facing target. | percent chance | none | none | |
| 294 | SE_CriticalSpellChance | Spell Crit Chance | Modify chance to critical direct damage spells and modify spell critical spell damage by percent | critical chance | critical damage percent modifier | none | |
| 295 | SE_ReduceTimerSpecial | Reduce Timer Special | | | | | |
| 296 | SE_FcSpellVulnerability | Focus: Incoming Spell Damage | Modify incoming spell damage taken by percent. | min percent modifier | none | max percent modifier | Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.|
| 297 | SE_FcDamageAmtIncoming | Focus: Incoming Spell Damage Amt | Modify incoming spell damage taken by a flat amount. | amount | none | none | |
| 298 | SE_ChangeHeight | Pet Size | Modify size by percent. | percent modifier | none | none | |
| 299 | SE_WakeTheDead | Wake the Dead | Create temporary pet from nearby corpses. | 1 | none | duration seconds | Maximum range for corpse from caster is 250 units.|
| 300 | SE_Doppelganger | Doppelganger | Create a temporary pets that mirrors you appearance. | amount of pets | none | duration seconds | |
| 301 | SE_ArcheryDamageModifier | Archery Damage Modifer | Modify archery base damage by percent. | percent modifier | none | none | |
| 302 | SE_FcDamagePctCrit | Focus: Spell Damage (v302 before crit) | Modify spell damage by percent. Damage is applied before critical calculation. | min percent modifier | none | max percent modifier | Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.|
| 303 | SE_FcDamageAmtCrit | Focus: Spell Damage Amt (v303 before crit) | Modify spell damage by a flat amount. Damage is applied before critical  calculation.. | amount | none | none | |
| 304 | SE_OffhandRiposteFail | Secondary Riposte | Modify chance to avoid an offhand riposte by percent. | percent modifier | none | none | |
| 305 | SE_MitigateDamageShield | Damage Shield Mitigation | Modify incoming damage from damage shield using your off hand weapon by amount. | amount | none | none | For spells/items set to positive value will reduce the damage shield amount, for AA's set this value to negative for reducation, this is converted to positive in source code. This is how live has it.|
| 306 | SE_ArmyOfTheDead | Army of the Dead | Create temporary pets from nearby corpses. Can only spawn one pet per corpse up a maximum amount of pets. | amount of pets | none | duration seconds | Maximum range for corpse from caster is 250 units.|
| 307 | SE_Appraisal | Appraisal | Roughly estimates the selling price of the item you are holding. | | | | |
| 308 | SE_ZoneSuspendMinion | Zone Suspend Minion | Allow suspended pets to be resummoned upon zoning. | 1 | none | none | |
| 309 | SE_GateCastersBindpoint | Gate Caster's Bindpoint | Teleports target or group members to casters bind point. | bind id (Set to 1, 2, or 3) | none | none | |
| 310 | SE_ReduceReuseTimer | Decrease Reuse Timer | Reduce spell recast time and disciple reuse time by amount. | time ms | none | none | Positive value reduces reuse timer. Note: You can set to negative to increase reuse timer, but client will not display it properly.|
| 311 | SE_LimitCombatSkills | Limit:  Combat Skills Not Allowed | Include or exclude combat skills or procs from being focused. |  0=Exclude if proc 1=Allow only if proc | none | none | |
| 312 | SE_Sanctuary | Sanctuary | Places caster at bottom hate list, effect fades if caster casts spell on targets other than self. | 1 | none | none | |
| 313 | SE_ForageAdditionalItems | Forage Master | Chance to forage additional items using  forage ability. | percent chance | none | none | |
| 314 | SE_Invisibility2 | Improved Invisibility | Apply invsibility that will last until duration ends | invisibility level | none | none | Invisibility level determines what level of see invisible can detect it.|
| 315 | SE_InvisVsUndead2 | Improved Invisibility Vs Undead | Apply invsibility verse undead that will last until duration ends | invisibility level | none | none | Invisibility level determines what level of see invisible can detect it.|
| 316 | SE_ImprovedInvisAnimals | Improved Invisibility Vs Animals | Apply invsibility verse animal that will last until duration ends | invisibility level | none | none | Invisibility level determines what level of see invisible can detect it.|
| 317 | SE_ItemHPRegenCapIncrease | Worn Regen Cap | Increase HP regen from items over cap by amount. | amount | none | none | |
| 318 | SE_ItemManaRegenCapIncrease | Worn Mana Cap | Increase Mana regen from items over cap by amount. | amount | none | none | |
| 319 | SE_CriticalHealOverTime | Critical HP Regen | Modifies chance to perform a critical heal over time by percent. | percent modifier | none | none | |
| 320 | SE_ShieldBlock | Shield Block Chance | Chance to block an attack while shield is equiped. | percent chance | none | none | |
| 321 | SE_ReduceHate | Reduce Target Hate | Remove or add a set amount of hate instantly from target. | amount | none | max amount | Positive value decreases hate. | Negative value increases hate.|
| 322 | SE_GateToHomeCity | Gate to Starting City | Gate to original starting city. | 1 | none | none | |
| 323 | SE_DefensiveProc | Add Defensive Proc | Add defensive proc that triggers from incoming melee hits. | spellid | rate modifer | none | |
| 324 | SE_HPToMana | HP for Mana | Casted spells will use HP instead of Mana with a conversion penalty rate. | conversion rate percent | | | |
| 325 | SE_NoBreakAESneak | No Break AE Sneak | Chance to remain hidden when hit by an area effect spell. | percent chance | none | none | |
| 326 | SE_SpellSlotIncrease | Spell Slots | Increase spell gems by amount. | amount | none | none | Client has to support the value you use.|
| 327 | SE_MysticalAttune | Buff Slots | Increase max spell buff slots amount. | amount | none | none | Client has to support the value you use.|
| 328 | SE_DelayDeath | Negative HP Limit | Increases the amount of damage you can take before dying after falling unconscious at 0 health. | amount | none | none | Positive value to increase effective negative hit points.|
| 329 | SE_ManaAbsorbPercentDamage | Mana Shield Absorb Damage | Reduces incoming damage by percent and converts that amount to mana loss. | mitigation percent | none | none | |
| 330 | SE_CriticalDamageMob | Critical Melee Damage | Modifies damage done from a critical melee hit by percent. | percent modifier | skill type (-1 = all skill types) | none | |
| 331 | SE_Salvage | Item Recovery | Modify chance to salvage from tradeskills by percent. | percent modifier | none | none | Positive values increase chance to salvage | Negative values decrase chance to salvage.|
| 332 | SE_SummonToCorpse | Summon to Corpse | Summons a player back to their corpse but does not restore any lost experience. Coprse can still be resurrected. | seen 0 or 1 | none | none | |
| 333 | SE_CastOnRuneFadeEffect | Trigger Spell On Rune Fade | Cast a spell when rune amount is used up and fades. | spellid | none | none | This effect needs to go on a spell containing a rune effect.|
| 334 | SE_BardAEDot | Bard AE Dot | Use on area of effect damage over time songs. Damage is only done if target is not moving. | amount | none | amount max | |
| 335 | SE_BlockNextSpellFocus | Focus: Block Next Spell | Chance to block next spell that meets the focus limits. | percent chance | none | none | |
| 336 | SE_IllusionaryTarget | Illusionary Target | | | | | |
| 337 | SE_PercentXPIncrease | Experience | Modify amount of experience gained. | percent modifier | none | none | |
| 338 | SE_SummonAndResAllCorpses | Expedient Recovery | Summon and ressurect all corpses for 100% experience. | Seen at 70 | | | Unknown what base value represents.|
| 339 | SE_TriggerOnCast | Focus: Trigger on Cast | Chance to cast an additional spell on the target when the spell you are casting meets the focus limits. | percent chance | spellid | none | |
| 340 | SE_SpellTrigger | Spell Trigger: Only One Spell Cast | Chance to cast a spell on the target. When multiple of this effects are present, only one is cast. | percent chance | spellid | none | When multiple of this effect exist on the same spell, only one spell will be selected from the list to be cast. For best results, the total percent chance should equal 100%. Example, Slot 1: Cast Ice Nuke 20%, Slot2: Cast Fire Nuke 50%, Slot3 Cast Magic Nuke 30%. When the spell is cast, only one of these spells be triggered on the target.|
| 341 | SE_ItemAttackCapIncrease | Worn Attack Cap | Increase ATK from items over cap by amount. | amount | none | none | |
| 342 | SE_ImmuneFleeing | Prevent Flee on Low Health | Prevent NPC from fleeing at low health. | 1 | none | none | |
| 343 | SE_InterruptCasting | Spell Interrupt | Chance to interrupt targets spell casting can be instant or per buff tick. | percent chance | none | none | |
| 344 | SE_ChannelChanceItems | Item Channeling | Modify chance to channel a spell from items and avoid being interupted by percent. | percent modifer | none | none | |
| 345 | SE_AssassinateLevel | Assassinate Max Level | Set max target level and bonus proc chance for assissinate. | max target level | proc chance bonus | none | |
| 346 | SE_HeadShotLevel | Headshot Max Level | Set max target level and bonus proc chance for assissinate. | max target level | proc chance bonus | none | |
| 347 | SE_DoubleRangedAttack | Double Ranged Attack | Chance to perform an additional ranged attack. | percent chance | none | none | Will consume an additional ammo item.|
| 348 | SE_LimitManaMin | Limit:  Min Mana | Mininum mana of spell that can be focused. | mana amount | none | none | |
| 349 | SE_ShieldEquipDmgMod | Damage With Shield | Modify melee damage and hate when having a shield equiped by percentage. | percent modifier | none | none | |
| 350 | SE_ManaBurn | Manaburn | Instantly drains mana for damage at a defined ratio up to a defined maximum amount of mana. | max amount of mana drained | percent of mana converted to damage | none | Limit value if set to negative will result in damage | Limit value if set to positive will result in a heal. Once this affect is applied, you can not apply another mana burn until the buff with this effect fades. Example, if base is set to 1000 and limit is to -50, spell will drain 1000 mana and do -500 damage to the target. Calc: (1000/-50)= -500.|
| 351 | SE_PersistentEffect | Persistent Effect | Create an aura that will provide a persistent effect around your character. | unknown | none | unknown | Set 'Teleport Zone' field to be the same as 'name' field in of the pet you want in the auras table.|
| 352 | SE_IncreaseTrapCount | Trap Count | Increase trap count by amount. | amount | none | none | |
| 353 | SE_AdditionalAura | Aura Count | Increase aura count by amount. | amount | none | none | |
| 354 | SE_DeactivateAllTraps | Deactivate All Traps | | | | | |
| 355 | SE_LearnTrap | Learn Trap | | | | | |
| 356 | SE_ChangeTriggerType | Change Trigger Type | | | | | |
| 357 | SE_FcMute | Focus: Mute | Chance to prevents casting of spell if limits are met. | chance | none | none | |
| 358 | SE_CurrentManaOnce | Mana Once | Modify mana by amount. Instant only. | amount | none | max amount (use positive value) | |
| 359 | SE_PassiveSenseTrap | Passive Sense Trap | Grants you a chance to innately sense traps when you near them.T his passive chance is only half as effective as your active Sense Trap skill. | | | | |
| 360 | SE_ProcOnKillShot | Trigger Spell On Kill Shot | Chance to cast a spell if you are the one responsible for the death of a challenging foe. | percent chance | spellid | minimum target level | Typical use case is a self only buff when triggered.|
| 361 | SE_SpellOnDeath | Trigger Spell On Death | Chance to casts a spell when the entity with this effect is killed. | percent chance | spellid | none | Typical use case is casting an area of effect upon death. Placing self only beneficial spells such as heals will not work due to player already being dead.|
| 362 | SE_PotionBeltSlots | Potion Belt Slots | | | | | |
| 363 | SE_BandolierSlots | Bandolier Slots | | | | | |
| 364 | SE_TripleAttackChance | Triple Attack Chance | Modify chance to triple attack by percentage. | percent modifier | none | none | |
| 365 | SE_ProcOnSpellKillShot | Trigger Spell on Spell Kill Shot | Chance to cast a spell when your target is killed and the kill is caused by the specific spell with this effect in it. | percent chance | spellid | none | This effect is typical found on direct damage spells.|
| 366 | SE_GroupShielding | Group Shielding | | | | | |
| 367 | SE_SetBodyType | Modify Body Type | Set body type of target. | body type | none | none | |
| 368 | SE_FactionMod | Modify Faction | | | | | |
| 369 | SE_CorruptionCounter | Corruption Counter | Determines potency of determental corruption spells or potency of cures. | amount | none | none | Set to positive values for potency of detrimental spells | Set to negative value for potency of cure spells.|
| 370 | SE_ResistCorruption | Corruption Resist | Modify fire corruption by amount | amout | none | max amount | |
| 371 | SE_AttackSpeed4 | Attack Speed: Inhibit Melee | Decrease the remaining portion of your targets attack speed value not lowered by standard slows, this stacks with other slow effects. | slow amount | none | none | This effect works differently than other slows. Base should always be positive.  Example: (SPA 11) Sha's Legacy 65% slow + (SPA 371) Lassitude 25% slow, Sha's Legacy is calculated as 100 - 35 = 65% slow, therefore the remaining attack speed is (35). Lassitude will now decrease the remaining value (35) by 25% = 16.5%. The total slowed value on the target would be 65% + 16.5% = 81.25% slow. If SPA 371 is only slow effect on target, then it will be slowed for full base value.|
| 372 | SE_ForageSkill | Grant Foraging | Grants forage skill. | 1 | none | none | |
| 373 | SE_CastOnFadeEffectAlways | Cast Spell On Fade (v373) | Cast a spell only if buff containing this effect fades after the full duration.  | spellid | none | none | |
| 374 | SE_ApplyEffect | Spell Trigger: Apply Each Spell | Chance to cast a spell on the target. When multiple of this effect are present, each is applied. | percent chance | spellid | none | |
| 375 | SE_DotCritDmgIncrease | Critical DoT Damage | Modify damage from critical damage over time by percent. | percent modifier | none | none | |
| 376 | SE_Fling | Fling | Knocbkack to target. Handled client side. | | | | |
| 377 | SE_CastOnFadeEffectNPC | Cast Spell On Fade (v377) | Cast a spell only if buff containing this effect fades after the full duration.  | spellid | none | none | |
| 378 | SE_SpellEffectResistChance | Spell Effect Resist Chance | Chance to resist specific spell effect. | chance modifier | spell effect id | none | |
| 379 | SE_ShadowStepDirectional | Directional Shadowstep | Teleport a specified distance and direction. | distance unit | direction degrees | | This effect is handled client side. Unclear how base value equates to actual in game distance movement. Limit directional values example,  0: Shadowstep Forward,  90: Shadowstep Right,  180:Shadowstep Back,  270: Shadowstep Left|
| 380 | SE_Knockdown | Knockback | Knockback effect. | push up? | push back? | none | Handled by client.|
| 381 | SE_KnockTowardCaster | Fling Target to Caster | Knockback to caster. Handled client side. | | | |    |
| 382 | SE_NegateSpellEffect | Negate Spell Effect | Negates specific spell effect benefits for the duration of the debuff and prevent non-duration spell effect from working. | negate spell effect type | spell effect id | none | |
| 383 | SE_SympatheticProc | Focus: Proc on Spell Cast | Cast on spell use with proc rate determined by cast time. | proc rate modifier | spellid | none | Typically found on item focus effects. Longer cast time spells are adjusted to have higher proc rates.|
| 384 | SE_Leap | Fling Caster to Target | Teleport to location that is a specificed distance away. | distance | none | none | |
| 385 | SE_LimitSpellGroup | Limit:  SpellGroup | Spell group(s) that a spell focus can require or exclude, base1: spellgroup id, Include: Positive Exclude: Negative | spellgroup id | none | none | Include set value to positive | Exclude set value to negative|
| 386 | SE_CastOnCurer | Trigger Spell on Curer | Casts a spell on the curer of the afflicition.  | spellid | none | none | This effect goes on the spell that needs curing.|
| 387 | SE_CastOnCure | Trigger Spell on Cure  | Casts a spell on the entity being cured of the afflicition.  | spellid | none | none | This effect goes on the spell that needs curing.|
| 388 | SE_SummonCorpseZone | Summon All Corpses | Summons all of the corpses of your targeted group member to that target's location. | 1 | none | none | |
| 389 | SE_FcTimerRefresh | Focus: Spell Gem Refresh | Reset all recast spell gem timers, can limit to reset only specific spell gem timers. | 1 | none | none | Applied from casted spells only.|
| 390 | SE_FcTimerLockout | Focus: Spell Gem Lockout | Set a spell gem to be on recast timer. | recast duration milliseconds | none | none | Applied from casted spells only.|
| 391 | SE_LimitManaMax | Limit:  Max Mana | Mininum mana of spell that can be focused | mana amount | none | none | |
| 392 | SE_FcHealAmt | Focus: Healing Amount (v392) | Modify healing by a flat amount. | amount | none | none | |
| 393 | SE_FcHealPctIncoming | Focus: Incoming Healing Effectiveness (v392) | Modify incoming heals by a percentage. | percent modifier | none | none | |
| 394 | SE_FcHealAmtIncoming | Focus: Incoming Healing Amount | Modify incoming heals by a flat amount. | amount | none | none | |
| 395 | SE_FcHealPctCritIncoming | Focus: Incoming Healing Effectiveness (v395) | Modify incoming heals by a percentage. | percent modifier | none | none | |
| 396 | SE_FcHealAmtCrit | Focus: Healing Amount (v396 before crit) | Modify incoming heals by a flat amount. Healing is added before critical calculation. | amount | none | none | |
| 397 | SE_PetMeleeMitigation | Pet Mitigation | Modify pets incoming damage mitigation by percentage. | percent modifier | none | none | |
| 398 | SE_SwarmPetDuration | Focus: Swarm Pet Duration | Increase swarm pet duration by amount time. | duration ms | none | none | |
| 399 | SE_FcTwincast | Focus: Twincast Chance | Chance to cast a spell a second time. | percent chance | none | none | |
| 400 | SE_HealGroupFromMana | Heal Group From Mana | Drains your mana and heals your target for an amount determined by ratio modifier for each point of mana drained. | max amount mana drained | ratio of HP gain per 1 mana drained | none | Ratio is calculated as value / 10, example if you want to heal 13 HP for every 1 mana, set base to 130.|
| 401 | SE_ManaDrainWithDmg | Damage for Amount Mana Drained | Drains your mana and damages a target for an amount determined by ratio modifier for each point of mana drained. | max amount mana drained | ratio of HP of damage per 1 mana drained | | Ratio is calculated as value / 10, example if you want to damage for l 13 HP for every 1 mana, set base to 130.|
| 402 | SE_EndDrainWithDmg | Damage for Amount Endurance Drained  | Drains your mana and damages a target for an amount determined by ratio modifier for each point of mana drained. | max amount endurance drained | ratio of HP of damage per 1 endurance drained | | Ratio is calculated as value / 10, example if you want to damage for l 13 HP for every 1 endurance, set base to 130.|
| 403 | SE_LimitSpellClass | Limit:  SpellClass | Spell Category' using table field 'spell_class' that a spell focus can require or exclude. | spell_class id | none | none | Include set value to positive | Exclude set value to negative. Definations for spell_class are not known.|
| 404 | SE_LimitSpellSubclass | Limit:  SpellSubclass | Spell Category Subclass' using table field 'spell_subclass' that a spell focus can require or exclude. | spell_subclass id | none | none | Include set value to positive | Exclude set value to negative. Definations for spell_subclass are not known.|
| 405 | SE_TwoHandBluntBlock | Staff Block Chance | Chance to block an attack while two handed blunt weapon is equiped. | percent chance | none | none | |
| 406 | SE_CastonNumHitFade | Trigger Spell on Hit Count Fade  | Cast a spell when a spell buffs hit counter is depleted. | spellid | none | none | |
| 407 | SE_CastonFocusEffect | Trigger Spell on Focus Effect Success | Casts a spell if the spells focus limits are met. | spellid | none | none | This effect goes on a spell with a focus effect. |
| 408 | SE_LimitHPPercent | Heal Up To Percent Limit | Cap HP at lowest of percent or amount. Can not be healed beyond this limit. | percent HP | amount HP | none | |
| 409 | SE_LimitManaPercent | Restore Mana Up To Percent Limit | Cap Mana at lowest of percent or amount. Can have mana restored beyond this limit. | percent HP | amount HP | none | |
| 410 | SE_LimitEndPercent | Restore Endurance Up To Percent Limit | Cap HP at lowest of percent or amount. Can not be healed beyond this limit. | percent HP | amount HP | none | |
| 411 | SE_LimitClass | Limit:  PlayerClass | Class(es) that can use the spell focus. | class bitmask | none | none | The class value in dbase is +1 in relation to item class value, set as you would item for multiple classes.|
| 412 | SE_LimitRace | Limit:  Race | Race that can use the spell focus. | race id | none | none | Not used in any known live spells. Use only single race at a time.|
| 413 | SE_FcBaseEffects | Focus: Base Spell Value | Modify spell effect value after formula calculations but before other focuses. | percent modifier | none | none | Used to set bard instrument modifiers. Can be used to focus many effects otherwise not able to be focused by other methods. If modifying spells that are buffs, value must be intervals of 10% and starting at greater than 10%. For example, if a 10% increase in rune amount is required, set base value to 11, that will be calculated as base value * 1.10, resulting in a 10% increase. For a 20% increase set to 12.|
| 414 | SE_LimitCastingSkill | Limit:  CastingSkill | Spell and singing skills(s) that a spell focus can require or exclude. | skill type | none | none | Include set value to positive | Exclude set value to negative.|
| 415 | SE_FFItemClass | Limit:  ItemClass | Limits focuses to be applied only from item click. | item ItemType | item SubType | item Slots | Not used on live. Details, base: item ItemType (-1 to include for all ItemTypes, -1000 to exclude clicks from getting the focus, or exclude specific SubTypes or Slots if set), limit: item SubType (-1 for all SubTypes), max: item Slots (bitmask of valid slots, -1 ALL slots), See comments in Mob::CalcFocusEffect for more details. Special: Can be used with SPA 310 reduce item click recast and SPA 127, 500,5001 to reduce item click casting time.|
| 416 | SE_ACv2 | AC (v416) | Modify AC by amount | amount | none | amount | |
| 417 | SE_ManaRegen_v2 | Mana Regen (v417) | Modify mana by amount, repeates every tic if in a buff. | amount | none | max amount (use positive value) | |
| 418 | SE_SkillDamageAmount2 | Skill Damage Bonus (v418) | Add a flat amount of damage when a specific melee skill is used. | amount | skill type (-1 = all skill types) | none | |
| 419 | SE_AddMeleeProc | Add Melee Proc (v419)  | Add proc to melee. | spellid | rate modifer | none | |
| 420 | SE_FcLimitUse | Focus: Hit Count | Modify buff hit count by percent. | percent modifier | none | none | Not used in any known live spells. Hit count is set using field 'numhits' in spells new.|
| 421 | SE_FcIncreaseNumHits | Focus: Hit Count Amount | Modify buff hit count by flat amount. | hit count amount | none | none | Hit count is set using field 'numhits' in spells new.|
| 422 | SE_LimitUseMin | Limit:  Minimum Hit Count | Minium amount of hit count for a spell to be focused. | hit count amount | nonw | none | Hit count is set using field 'numhits' in spells new.|
| 423 | SE_LimitUseType | Limit:  Hit Count Type | Focus will only affect if has this hit count type. | hit count type | nonw | none | hit type is set using field 'numhitstype' in spells_new.  1:Incoming Hit Attempts, 2:Outgoing Hit Attempts, 3:Incoming Spells,  4:Outgoing Spells, 5: Outgoing Hit Successes, 6:Incoming Hit Successes, 7:Matching Spells, 8:Incoming Hits Or Spells, 9:Reflected Spells,  10:Defensive Proc Casts,  11: Offensive Proc Casts.|
| 424 | SE_GravityEffect | Gravitate | Causes a target that is within specific distance of you to gravitate toward or away from you with set amount of force over a period of time. | Force | Distance | unknown | Negative base value will pull target | Positive base value will push target. Max value on live is a level modifier, unclear what it is in our current spell file.|
| 425 | SE_Display | Fly | Grants flying animation to certain illusions. | | | | |
| 426 | SE_IncreaseExtTargetWindow | AddExtTargetSlots | Increases the capacity of your extended target window | | | | |
| 427 | SE_SkillProc | Skill Proc On Attempt | Add proc to a specific combat or ability skill is used. | spellid | rate modifer | none | Example, can add a proc to taunt which will have a chance to fire each time you use skill.|
| 428 | SE_LimitToSkill | Limit To Skill | Limits what combat skills will trigger a skill (SPA 426, 429), melee (SPA 419) or ranged (SPA 201) proc. | skill type | none | none | This needs to be placed after the proc SPA in a spell to be checked properly.|
| 429 | SE_SkillProcSuccess | Skill Proc On Success | Add proc to a specific combat or ability skill when that skill is succesfully used. | spellid | rate modifer | none | Example, can add a proc to taunt which will have a chance to fire if you taunt successsfully.|
| 430 | SE_PostEffect | PostEffect | | | | | |
| 431 | SE_PostEffectData | PostEffectData | | | | | |
| 432 | SE_ExpandMaxActiveTrophyBen | Expand Max Active Trophy Benefits | | | | | |
| 433 | SE_CriticalDotDecay | Critical DoT Decay | Increase critical dot chance, effect decays based on level of spell it effects. | percent chance | decay modifier | none | Live no longer uses this effect, replaced after ROF2 with different effect. Effects were introduced during VoA.|
| 434 | SE_CriticalHealDecay | Critical Heal Decay | Increase critical heal chance, effect decays based on level of spell it effects. | percent chance | decay modifier | none | Live no longer uses this effect, replaced after ROF2 with different effect. Effects were introduced during VoA.|
| 435 | SE_CriticalRegenDecay | Critic Heal Over Time Decay | Increase critical heal over time chance, effect decays based on level of spell it effects. | percent chance | decay modifier | none | Live no longer uses this effect, replaced after ROF2 with different effect. Effects were introduced during VoA.|
| 436 | SE_BeneficialCountDownHold | Toggle Freeze Buff Timers | | | | | |
| 437 | SE_TeleporttoAnchor | Teleport to Anchor | | | | | |
| 438 | SE_TranslocatetoAnchor | Translocate to Anchor | | | | | |
| 439 | SE_Assassinate | Assassinate | Grants  backstab and critical throwing attacks a chance to deal massive damage to targets when attacking from behind. | percent chance | damage amount | none | Use with SPA 345 to set max target level and bonus chance.|
| 440 | SE_FinishingBlowLvl | Finishing Blow Max Level | Sets the max level and max hit point ratio an NPC must be to receive a Finishing Blow. | max target level | hit point percent to trigger | none | Limit value is calculated as limit / 10. To set FB to trigger below 10 pct HP, set limit to 100. If multiple version of this affect exist it will use highest HP percent.|
| 441 | SE_DistanceRemoval | Distance Buff Removal | Buff is removed from target when target moves specified amount of distance away from where initially hit. | distance amount | none | none | |
| 442 | SE_TriggerOnReqTarget | Trigger Spell on Target Requirement | Cast a spell when Target Requirement conditions are met. | spellid | Target Restriction ID | none | See enum SpellRestriction in spdat.h for IDs. This trigger spell is usually cast on a target.|
| 443 | SE_TriggerOnReqCaster | Trigger Spell on Caster Requirement | Cast a spell when Caster Requirement conditions are met | spellid | Target Restriction ID | none | See enum SpellRestriction in spdat.h for IDs. This trigger spell is usually cast on self.|
| 444 | SE_ImprovedTaunt | Improved Taunt | Causes target to locks aggro on the caster and decreases other players hate by percentage on NPC targets below the specified max level.  | 999 | other players percent of hate generation | max target level | Limit greater than 100  increased hate generation on other players  (120 = 20 pct increased hate generation) |  Limit less than 100 for decreased hate generation on other players (80 = 20 pct decrease in hate generation)|
| 445 | SE_AddMercSlot | Add Merc Slot | | | | | |
| 446 | SE_AStacker | A Stacker | Buff blocker A | stacking value | none | none | "Buffs containing these effects can block each other from taking hold via the following. Does not matter what slot the effect is in. (B) Blocks any buffs from taking hold with (A) in it. (C) Blocks any buff from taking hold with (B) in it.|
| (D) Blocks any buff from taking hold with (C) in it. When checking same type (ie A vs A), the higher base effect value will determine which takes hold."|
| 447 | SE_BStacker | B Stacker | Buff blocker B | stacking value | none | none | "Buffs containing these effects can block each other from taking hold via the following. Does not matter what slot the effect is in. (B) Blocks any buffs from taking hold with (A) in it. (C) Blocks any buff from taking hold with (B) in it.|
| (D) Blocks any buff from taking hold with (C) in it. When checking same type (ie A vs A), the higher base effect value will determine which takes hold."|
| 448 | SE_CStacker | C Stacker | Buff blocker C | stacking value | none | none | "Buffs containing these effects can block each other from taking hold via the following. Does not matter what slot the effect is in. (B) Blocks any buffs from taking hold with (A) in it. (C) Blocks any buff from taking hold with (B) in it.|
| (D) Blocks any buff from taking hold with (C) in it. When checking same type (ie A vs A), the higher base effect value will determine which takes hold."|
| 449 | SE_DStacker | D Stacker | Buff blocker D | stacking value | none | none | "Buffs containing these effects can block each other from taking hold via the following. Does not matter what slot the effect is in. (B) Blocks any buffs from taking hold with (A) in it. (C) Blocks any buff from taking hold with (B) in it.|
| (D) Blocks any buff from taking hold with (C) in it. When checking same type (ie A vs A), the higher base effect value will determine which takes hold."|
| 450 | SE_MitigateDotDamage | Mitigate Damage Over Time Rune | Mitigate incoming damage from damage over time spells by percentage until rune fades. | percent mitigation | max damage absorbed per hit | rune amount | Special: If this effect is placed on item as worn effect or as an AA, it will provide stackable percent damage over time mitigation for the base value.|
| 451 | SE_MeleeThresholdGuard | Melee Threshold Guard | Partial Melee Rune that only is lowered if melee hits are over the specified amount of damage. | percent mitigation | minimum damage to be lowered | rune amount | |
| 452 | SE_SpellThresholdGuard | Spell Threshold Guard | Partial Spell Rune that only is lowered if spell hits are over the specified amount of damage. | percent mitigation | minimum damage to be lowered | rune amount | |
| 453 | SE_TriggerMeleeThreshold | Trigger Spell on Melee Threshold | Cast a spell when a specified amount of melee damage taken in a single hit. | spellid | amount of melee damage | none | |
| 454 | SE_TriggerSpellThreshold | Trigger Spell on Spell Threshold | Cast a spell when a specified amount of spell damage is taken in a single hit. | spellid | amount of spell damage | none | |
| 455 | SE_AddHatePct | Add Hate Percent | Modify targets total hate toward the caster by percentage. | percent modifier | none | none | Positive value increased hate. | Negative value decreases hate.|
| 456 | SE_AddHateOverTimePct | Add Hate Over Time Percent | Modify targets total hate toward the caster by percentage,  repeates every tic if in a buff.  | percent modifier | none | none | Positive value increases hate. | Negative value decreases hate.|
| 457 | SE_ResourceTap | Resource Tap | Return a percentage of Spell Damage as (HP/Mana/Endurance), up to a maximum amount per spell.  | percent coverted | 0=HP, 1=Mana,2=Endurance | max amount resource returned | Conversion percent calculated as value / 10, example to convert 85 percent of damage to mana set base value to 850 and limit to 1.|
| 458 | SE_FactionModPct | Faction Pct | Modify faction gains and losses by percent. | percent modifier | none | none | |
| 459 | SE_DamageModifier2 | Skill Damage Modifier (v459) | Modify melee damage by skill. | percent modifer | skill type (-1 = all skill types) | none | |
| 460 | SE_Ff_Override_NotFocusable | Limit:  Include Non-Focusable | Allow spell to be focused even if flagged with 'not_focusable' in spell table. | 1 | none | none | |
| 461 | SE_ImprovedDamage2 | Focus: Spell Damage (v461) | Modify outgoing spell damage by percentage. | min percent | none | max percent | Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.|
| 462 | SE_FcDamageAmt2 | Focus: Spell Damage Amount (v462) | Modify outgoing spell damage by a flat amount | amount | none | none | |
| 463 | SE_Shield_Target | Shield Target | | | | | |
| 464 | SE_PC_Pet_Rampage | PC Pet Rampage | Percent chance for a pet to do a rampage melee attack with damage modifier each melee round. | percent chance | damage modifier | none | Limit greater than 100 for increased damage (120 = 20 pct increased damage) |  Limit less than 100 for decreased damage (80 = 20 pct decreased damage)|
| 465 | SE_PC_Pet_AE_Rampage | PC Pet AE Rampage | Percent chance for a pet to do a AE rampage melee attack with damage modifier each melee round. | percent chance | damage modifier | none | Limit greater than 100 for increased damage (120 = 20 pct increased damage) |  Limit less than 100 for decreased damage (80 = 20 pct decreased damage)|
| 466 | SE_PC_Pet_Flurry_Chance | PC Pet Flurry Chance | Percent chance for a pet to do flurry from double attack hit. | percent chance | none | none | |
| 467 | SE_DS_Mitigation_Amount | Damage Shield Mitigation Amount | Modify incoming damage shield damage by a flat amount. | amount | none | none | |
| 468 | SE_DS_Mitigation_Percentage | Damage Shield Mitigation Percentage | Modify incoming damage shield damage by percentage. | percent modifier | none | none | |
| 469 | SE_Chance_Best_in_Spell_Grp | Spell Trigger: Best in Spell Group: Only One Spell Cast | Chance to cast highest scribed spell within a spell group. When multiple of this effects are present, only one is cast. | percent chance | spellgroup id | | When multiple of this effect exist on the same spell, only one spell will be selected from the list to be cast. For best results, the total percent chance should equal 100%. Example, Slot 1: Cast Ice Nuke spellgroup 20%, Slot2: Cast Fire Nuke spellgroup 50%, Slot3 Cast Magic Nuke spellgroup 30%. When the spell is cast, only one of these spells be triggered on the target.|
| 470 | SE_Trigger_Best_in_Spell_Grp | Spell Trigger: Best in Spell Group: Apply Each | Chance to cast highest scribed spell within a spell group. Each spell has own chance. | percent chance | spellgroup id | | |
| 471 | SE_Double_Melee_Round | Double Melee Round (PC Only) | Percent chance to repeat primary weapon round with a percent damage modifier. | percent chance | percent damage modifier | none | |
| 472 | SE_Buy_AA_Rank | Toggle Passive AA Rank | Used in AA abilities that have Enable/Disable toggle. Spell on Disabled Rank has this effect in it. | 1 | none | none | Certain AA, like Weapon Stance line use a special toggle Hotkey to enable or disable the AA's passive abilities.  This is occurs by doing the following. Each 'rank' of Weapon Stance is actually 2 actual ranks.   First rank is always the Disabled version which cost X amount of AA, this is the rank that SPA 472 goes in. Second rank is the Enabled version which cost 0 AA.  When you buy the first rank, you make a hotkey that on live say 'Weapon Stance Disabled', if you clik that it then BUYS the  next rank of AA (cost 0) which switches the hotkey to 'Enabled Weapon Stance' and you are given the passive buff effects.  If you click the Enabled hotkey, it causes you to lose an AA rank and once again be disabled. Thus, you are switching between  two AA ranks. Thefore when creating an AA using this ability, you need generate both ranks. Follow the same pattern for additional ranks. See aa.cpp for further details.|
| 473 | SE_Double_Backstab_Front | Double Backstab From Front | Chance to double backstab from front. | percent chance | none | none | |
| 474 | SE_Pet_Crit_Melee_Damage_Pct_Owner | Pet Crit Melee Damage | Modifies your pets critical melee damage. | percent modifier | none | none | This effect goes on the pet owner and then the benefit is applied to the pet.|
| 475 | SE_Trigger_Spell_Non_Item | Trigger Spell: Not Cast From Items: Apply Each | Chance to cast a spell on a target if not cast by item click. Each spell has own chance.  | chance pecent | spellid | | |
| 476 | SE_Weapon_Stance | Weapon Stance | Apply a specific spell buffs automatically if depending 2Hander, Shield or Duel Wield is equiped. | spellid | 0=2H, 1=Shield, 2=DW | none | On live this is an AA, on emu can use as item or spell buff effect. Live AA uses a toggle system to turn on and off weapons stance. which requires use of SPA 472 SE_Buy_AA_Rank within in the AA. Details can be found in aa.cpp|
| 477 | SE_Hatelist_To_Top_Index | Move to Top of Rampage List | Chance to be set to top of rampage list | chacne percent | none | none | |
| 478 | SE_Hatelist_To_Tail_Index | Move to Bottom of Rampage List | Chance to be set to bottom of rampage list | percent chance | none | none | |
| 479 | SE_Ff_Value_Min | Limit:  Base Value Min | Minimum base value of a spell effect within a spell attempting to be focused. | effect_base_value | spell effect id | none | Example, only allow focus if spell has Effect ID 0 which is a Heal and the heals effect value is less than 5000.|
| 480 | SE_Ff_Value_Max | Limit:  Base Value Max | Maximum base value of a spell effect within a spell attempting to be focused. | effect_base_value | spell effect id | none | Example, only allow focus if spell has Effect ID 0 which is a Heal and the heals effect value is greater than 5000.|
| 481 | SE_Fc_Cast_Spell_On_Land | Focus: Trigger Spell on Spell Landing | Cast spell if hit by an incoming spell and that incoming spell meets limit requirements. | spellid | none | none | Example, everytime you are hit with a fire nuke, cast a heal on yourself.|
| 482 | SE_Skill_Base_Damage_Mod | Base Hit Damage | Modify base melee damage by percent. | percent modifier | skill type (-1 = ALL skill types) | none | |
| 483 | SE_Fc_Spell_Damage_Pct_IncomingPC | Focus: Incoming Spell Damage (v483) | Modify incoming spell damage taken by percent. | min percent modifier | none | max percent modifier | Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.|
| 484 | SE_Fc_Spell_Damage_Amt_IncomingPC | Focus: Incoming Spell Damage Amount (v484) | Modify incoming spell damage taken by a flat amount. | amount | none | none | |
| 485 | SE_Ff_CasterClass | Limit: Caster Class | Caster of the spell on target with a focus effect that is checked by incoming spells must be specified class(es). | class bitmask | none | none | Set multiple classes same as would for items. This is only used with focus effects that check against incoming spells.|
| 486 | SE_Ff_Same_Caster | Limit: Caster | Caster of spell on target with a focus effect that is checked by incoming spells. | 0=Must be different caster 1=Must be same caster | none | none | This is only used with focus effects that check against incoming spells.|
| 487 | SE_Extend_Tradeskill_Cap | Extend Tradeskill Cap | | | | | |
| 488 | SE_Defender_Melee_Force_Pct_PC | Push Taken | | | | | |
| 489 | SE_Worn_Endurance_Regen_Cap | Worn Endurance Regen Cap | Modify endurance regen cap from items. | amount | none | none | |
| 490 | SE_Ff_ReuseTimeMin | Limit:  Reuse Time Min | Minimum recast time of a spell that can be focused. | recast time | none | none | |
| 491 | SE_Ff_ReuseTimeMax | Limit:  Reuse Time Max | Maximum recast time of a spell that can be focused. | recast time | none | none | |
| 492 | SE_Ff_Endurance_Min | Limit:  Endurance Min | Minimum endurance cost of a spell that can be focused | endurance cost | none | none | |
| 493 | SE_Ff_Endurance_Max | Limit:  Endurance Max | Maximum endurance cost of a spell that can be focused. | endurance cost | none | none | |
| 494 | SE_Pet_Add_Atk | Pet Add ATK | Modifies your pets ATK by amount. | amount | none | none | This effect goes on the pet owner and then the benefit is applied to the pet.|
| 495 | SE_Ff_DurationMax | Limit:  Duration Max | Maximum duration of spell that can be focused. | tics | | | |
| 496 | SE_Critical_Melee_Damage_Mod_Max | Critical Melee Damage: No Stack | Modifies damage done from a critical melee hit by percent. | percent modifier | skill type (-1 = all skill types) | none | |
| 497 | SE_Ff_FocusCastProcNoBypass | Limit:  Proc No Bypass | | | | | |
| 498 | SE_AddExtraAttackPct_1h_Primary | Add Extra Attack: 1H Primary | Gives your double attacks a percent chance to perform an extra attack with 1-handed primary weapon. | percent chance | number of attacks | none | |
| 499 | SE_AddExtraAttackPct_1h_Secondary | Add Extra Attack: 1H Secondary | Gives your double attacks a percent chance to perform an extra attack with 1-handed secondary weapon. | percent chance | number of attacks | none | |
| 500 | SE_Fc_CastTimeMod2 | Focus: Spell Haste (v500, no cap)  | Modify cast time by percent. | percent modifier | none | none | Can reduce cast time to 0.|
| 501 | SE_Fc_CastTimeAmt | Focus: Spell Cast Time | Modify cast time by flat amount. | time ms | none | none | Can reduce cast time to 0.|
| 502 | SE_Fearstun | Fearstun | Stun with a max level limit. Fear restrictions apply. Normal stun restrictions do not apply. | duration ms | PC duration ms | target max level | Max value can be calculated in two ways, to set a max level that target can be affected the formula is value - 1000, example if max level is 80, then set as 1080. If you want to set max level to be relative to caster, example only affects entities that are 3 more less level higher than caster, then set max value to 3.|
| 503 | SE_Melee_Damage_Position_Mod | Rear Arc Melee Damage Mod | Modify melee damage by percent if done from Front or Behind opponent. | percent modifier | 0=back, 1=front | | |
| 504 | SE_Melee_Damage_Position_Amt | Rear Arc Melee Damage Amt | Modify melee damage by flat amount if done from Front or Behind opponen. | amount | 0=back, 1=front | | |
| 505 | SE_Damage_Taken_Position_Mod | Rear Arc Damage Taken Mod | Modify incoming melee damage by percent if damageis taken from Front or Behind. | percent modifier | 0=back, 1=front | | |
| 506 | SE_Damage_Taken_Position_Amt | Rear Arc Damage Taken Amt | Modify incoming melee damage by flat amount if damage is taken from your Front or Behind. | amount | 0=back, 1=front | | |
| 507 | SE_Fc_Amplify_Mod | Focus: Spell Heal Damage and DoT | Modify spell damage, healing and damge over time by percent. | percent modifier | none | none | |
| 508 | SE_Fc_Amplify_Amt | Focus: Spell Heal Damage and DoT Amount | Modify spell damage, healing and damge over time by a flat amount. | amount | none | none | |
| 509 | SE_Health_Transfer | Health Transfer | Exchange health for damage or healing on a target. | casters percent HP change | percent of casters HP to damage | none | Used in newer versions of  Lifeburn and Act of Valor. Base is calculated as value / 100, where a reducation of 75 percent of casters HP would be value 750. Limit is calculated as value / 10, if value is set to 1000 then damage will be 100 percent of casters HP. Negative base decreases casters HP, negative limit decreases targets HP, positive limit heals target.|
| 510 | SE_Fc_ResistIncoming | Focus: Incoming Resist Modifier | Modify incoming spells resist modifier by amount. | amount | none | none | Positive value will lower resist modifier. Every 2 pts of value equals a resist rate decrease of 1 percent. Example, base -750 and limit -10000, will consumes 75% of your current health and deals 1000% of that health as direct damage.|
| 511 | SE_Ff_FocusTimerMin | Limit: Focus Reuse Timer | Sets a recast time until focus can be used again. | 1 | time ms | none | Example, set limit to 1500, then this focus can only trigger once every 1.5 seconds.|
| 512 | SE_Proc_Timer_Modifier | Proc Reuse Timer | Sets a recast time until a proc can fire again. | 1 | time ms | none | Example, set limit to 1500, then this proc can only trigger once every 1.5 seconds.|
| 513 | SE_Mana_Max_Percent | Mana Max Percent | | | | | |
| 514 | SE_Endurance_Max_Percent | Endurance Max Percent | | | | | |
| 515 | SE_AC_Avoidance_Max_Percent | AC Avoidance Max Percent | Modify AC avoidance by percent. | percent modifier | none | none | Calculated as base value / 100 is actual percent. Example for 7.14 percent modifier, set to 714.|
| 516 | SE_AC_Mitigation_Max_Percent | AC Mitigation Max Percent | Modify AC mitigation by percent. | percent modifier | none | none | Calculated as base value / 100 is actual percent. Example for 7.14 percent modifier, set to 714.|
| 517 | SE_Attack_Offense_Max_Percent | Attack Offense Max Percent | Modify ATK offense by percent. | percent modifier | none | none | Calculated as base value / 100 is actual percent. Example for 7.14 percent modifier, set to 714.|
| 518 | SE_Attack_Accuracy_Max_Percent | Attack Accuracy Max Percent | Modify ATK accuracy by percent. | percent modifier | none | none | Calculated as base value / 100 is actual percent. Example for 7.14 percent modifier, set to 714.|
| 519 | SE_Luck_Amount | Luck Amount | | | | | |
| 520 | SE_Luck_Percent | Luck Percent | | | | | |
| 521 | SE_Endurance_Absorb_Pct_Damage | Absorb Damage with Endurance | Reduces a percent of incoming damage using endurance, drains endurance at a ratio. | mitigation percentage | ratio | none | Base calculated as value / 100,  example if base 2000 then mitigation is 20 percent. Limit is calculated as value / 10000, example if limit 500 then ratio will be 0.05 Endurance reduced per 1 Hit Point of damage.|
| 522 | SE_Instant_Mana_Pct | Instant Mana Percent | Modify mana by percent of  maximum mana, or flat amount, whichever is lower. | percent of max mana | max amount | none | Negative base value increases mana | Positive base value increases mana. Base is calculated as value / 100. Example if base 150 and max mana 1000, effect will extracts 1.5% of your maximum mana, or 1000, whichever is lower.|
| 523 | SE_Instant_Endurance_Pct | Instant Endurance Percent | Modify endurance by percent of  maximum endurance, or flat amount, whichever is lower.  | percent of max endurance | max amount | none | Negative base value increases endurance | Positive base value increases endurance. Base is calculated as value / 100. Example if base 150 and max endurance 1000, effect will extracts 1.5% of your maximum endurance, or 1000, whichever is lower.|
| 524 | SE_Duration_HP_Pct | Duration HP Percent | Modify hit points by percent of  maximum hit points, or flat amount, whichever is lower, repeates every tic if in a buff. | percent of max hit points | max amount | none | Negative base value increases hit points | Positive base value increases hit points. Base is calculated as value / 100. Example if base 150 and max hit points 1000, effect will extracts 1.5% of your maximum hit points, or 1000, whichever is lower per tic.|
| 525 | SE_Duration_Mana_Pct | Duration Mana Percent | Modify mana by percent of  maximum mana, or flat amount, whichever is lower, repeates every tic if in a buff. | percent of max mana | max amount | none | Negative base value increases mana | Positive base value increases mana. Base is calculated as value / 100. Example if base 150 and max mana 1000, effect will extracts 1.5% of your maximum mana, or 1000, whichever is lower per tic.|
| 526 | SE_Duration_Endurance_Pct | Duration Endurance Percent | Modify endurance by percent of  maximum endurance, or flat amount, whichever is lower, repeates every tic if in a buff. | percent of max endurance | max amount | none | Negative base value increases endurance | Positive base value increases endurance. Base is calculated as value / 100. Example if base 150 and max endurance 1000, effect will extracts 1.5% of your maximum endurance, or 1000, whichever is lowe per tic.|
