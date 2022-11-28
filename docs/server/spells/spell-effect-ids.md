# Spell Effect IDs

Click [here](https://docs.google.com/spreadsheets/d/1gpnJXk5C-wurxvbi_O6EoywtRtWm9qkRdhGfy5Xo_sY/edit?usp=sharing) for a more detailed and up to date spreadsheet of spell effect information.

|  |  |  |  |  |  |
| :--- | :--- | :--- | :--- | :--- | :--- |
| ID | Spell Effect | Description | Base Value | Limit Value | Max Value |
| 0 | SE_CurrentHP | Direct Damage/Healing, duration allows for HoT/DoT | Amt (-)DD(+)Heal | Target Restrictions  | Max Amt |
| 1 | SE_ArmorClass | +/- AC | Amt | none | Max Amt |
| 2 | SE_ATK | +/- ATK | Amt | none | Max Amt |
| 3 | SE_MovementSpeed | +/- Movement Speed | Amt | none | Max Amt |
| 4 | SE_STR | +/- STR | Amt | none | Max Amt |
| 5 | SE_DEX | +/- DEX | Amt | none | Max Amt |
| 6 | SE_AGI | +/- AGI | Amt | none | Max Amt |
| 7 | SE_STA | +/- STA | Amt | none | Max Amt |
| 8 | SE_INT | +/- INT | Amt | none | Max Amt |
| 9 | SE_WIS | +/- WIS | Amt | none | Max Amt |
| 10 | SE_CHA | +/- CHA (BLANK SPACER if 0 in spell effect) | Amt | none | Max Amt |
| 11 | SE_AttackSpeed | &gt;100/&lt;100 Attack Speed (standard 'haste or 'slow) |  Attack Speed  | none | Max Amt |
| 12 | SE_Invisibility | Invisibility (Unstable)         Invisible Level corresponds to what type of see invisible will detect. | Invisible Level | none | none |
| 13 | SE_SeeInvis | See Invisibility                                                              Invisible Lv (1) = Standard invsibile  | Invisible Level | none | none |
| 14 | SE_WaterBreathing | Immune to drowning | ? (Seen from 1 - 3) | none | none |
| 15 | SE_CurrentMana | Direct +/- mana, duration allows for mana over time | Amt | none | Max Amt |
| 16 | SE_NPCFrenzy | UNKNOWN EFFECT [NOT] |  |  |  |
| 17 | SE_NPCAwareness | UNKNOWN EFFECT [NOT] |  |  |  |
| 18 | SE_Lull | Seen in conjunction with SE_ChangeFrenzyRad and SE_Harmony, this effect does nothing alone. | 0 | none | none |
| 19 | SE_AddFaction | Adjusts NPC 'con' (Alliance spell line) | Faction Amt | none | none |
| 20 | SE_Blind | Turn PC screen blank, causes NPC to flee if NOT in melee range | ? (Usually -1/1) | none | none |
| 21 | SE_Stun | Stun for a duration | Stun Time | ? (Always &lt;= Base) | Max Level |
| 22 | SE_Charm | Control an NPC/Player for a duration.   Mechanics  | ? | none | Max Level |
| 23 | SE_Fear | Cause your target to flee.   Mechanics  | ? | none | Max Level |
| 24 | SE_Stamina | Increased stamina (yellow bar) 'Invigor'  - ? Current Effect (Does not raise endurance on live) | ? (Always Negative) | none | none |
| 25 | SE_BindAffinity | Sets Bind point.                                     Bind Point ID (1=Primary, 2=Secondary 3=Tertiary)  | Bind Point ID | none | none |
| 26 | SE_Gate | Transport to Bind point                           Bind Point ID (1=Primary, 2=Secondary 3=Tertiary)  | Success Chance | Bind Point ID | none |
| 27 | SE_CancelMagic | Chance to remove any buff effect.   Mechanics  | Level Modifier | none | none |
| 28 | SE_InvisVsUndead | Invisiblity vs Undead (Unstable)                                          (Only Invisiblity Level (1) exists) | Invisible Level | none | none |
| 29 | SE_InvisVsAnimals | Invisiblity vs Animal (Unstable)                                          (Only Invisiblity Level (1) exists) | Invisible Level | none | none |
| 30 | SE_ChangeFrenzyRad | Sets NPC Aggro Radius to the value of the spell effect Mechanics  | Aggro Radius | none | Max Level |
| 31 | SE_Mez | Stuns target till duration ends or broken (melee/spell dmg)  Mechanics  | Stacking Value | none | Max Level |
| 32 | SE_SummonItem | Summon an item. | Item ID | none | none |
| 33 | SE_SummonPet | Summon Pet                  ('Teleport Zone' Field in spells_new table contains name of pet) | ? (Always 1) | none | none |
| 34 | SE_Confuse | UNKNOWN EFFECT [NOT] |  |  |  |
| 35 | SE_DiseaseCounter | Determines potency of determental disease spells (+) or potency of cures (-) | Counter Amt | none | Max Amt |
| 36 | SE_PoisonCounter | Determines potency of determental poision spells (+) or potency of cures (-) | Counter Amt | none | Max Amt |
| 37 | SE_DetectHostile | UNKNOWN EFFECT [NOT] |  |  |  |
| 38 | SE_DetectMagic | UNKNOWN EFFECT [NOT] |  |  |  |
| 39 | SE_DetectPoison | UNKNOWN EFFECT [NOT] |  |  |  |
| 40 | SE_DivineAura | Invulnerable to spells and melee, you can not cast or melee while under this effect. | ? (Usually 1) | none | none |
| 41 | SE_Destroy | Instantly kill target | ? (Always 1) | none | none |
| 42 | SE_ShadowStep | Warps player a short distance in a random direction | ? (Usually 1) | none | none |
| 43 | SE_Berserk | Sets client as 'Berserk' allowing for chance to crippling blow. [Custom] [NOT] | 1 | none | none |
| 44 | SE_Lycanthropy | Used as a place holder effect for preventing certain buffs from stacking | Stacking Value | none | none |
| 45 | SE_Vampirism | Heals you for a percent of your melee damage done to target.  [Custom] [NOT] | Amt % (+/-) | none | none |
| 46 | SE_ResistFire | +/- Fire Resist | Amt | none | Max Amt |
| 47 | SE_ResistCold | +/- Cold Resist | Amt | none | Max Amt |
| 48 | SE_ResistPoison | +/- Poison Resist | Amt | none | Max Amt |
| 49 | SE_ResistDisease | +/- Disease Resist | Amt | none | Max Amt |
| 50 | SE_ResistMagic | +/- Magic Resist | Amt | none | Max Amt |
| 51 | SE_DetectTraps | UNKNOWN EFFECT [NOT] |  |  |  |
| 52 | SE_SenseDead | Point player in direction of nearest Undead NPC | ? (Always 1) | none | ? |
| 53 | SE_SenseSummoned | Point player in direction of nearest Summoned NPC | ? (Always 1) | none | none |
| 54 | SE_SenseAnimals | Point player in direction of nearest Animal NPC | ? (Always 1) | none | none |
| 55 | SE_Rune | Absorb melee damage until a maxium amount then fades | Rune Amt | none | Max Amt |
| 56 | SE_TrueNorth | Point player in North direction. | ? (Always 1) | none | none |
| 57 | SE_Levitate | Levitation (Take no fall damage) | ? (Usually 1) | none | none |
| 58 | SE_Illusion | Changes visual race of target | Race ID (-1=Gender) | Texture ID | Helm ID |
| 59 | SE_DamageShield | Damage taken upon successful melee hit | Amt | none | Max Amt |
| 60 | SE_TransferItem | UNKNOWN EFFECT [NOT] |  |  |  |
| 61 | SE_Identify | Displays in text the 'lore' field from the items table | ? (Always 1) | none | none |
| 62 | SE_ItemID | UNKNOWN EFFECT [NOT] |  |  |  |
| 63 | SE_WipeHateList | Chance to wipe the NPC's hatelist | Success Chance | none | Max Chance |
| 64 | SE_SpinTarget | Spins Target and stuns | Spin Duration | ? (Always &lt;= base) | Max Level |
| 65 | SE_InfraVision | Improved night vision | ? (Always 1) | none | none |
| 66 | SE_UltraVision | Better improved night vision | ? (Always 1) | none | none |
| 67 | SE_EyeOfZomm | Transfers your vision and control to a temporary NPC | ? (Always 1) | none | none |
| 68 | SE_ReclaimPet | Kills your pet in exchange for mana. (Returns 75% of pet spell actual mana cost) | ? (Always 1) | none | none |
| 69 | SE_TotalHP | Increases Max HP (standard 'HP Buffs') | Amt | none | Max Amt |
| 70 | SE_CorpseBomb | UNKNOWN EFFECT [NOT] |  |  |  |
| 71 | SE_NecPet | Summon Pet                  ('Teleport Zone' Field in spells_new table contains name of pet) |  |  |  |
| 72 | SE_PreserveCorpse | UNKNOWN EFFECT [NOT] |  |  |  |
| 73 | SE_BindSight | See from the perspective of the target that this is cast on. | ? (Always 1) | none | none |
| 74 | SE_FeignDeath | Fall to the ground and clear hate list | Success Chance | ? (usually 0) | Max Chance |
| 75 | SE_VoiceGraft | Speak through your pet | ? (Always 1) | none | none |
| 76 | SE_Sentinel | Creates a proximity zone where cast that alerts caster if NPC's or Players enter it. | ? (Always 1) | none | none |
| 77 | SE_LocateCorpse | Turn player in direction of targeted players corpse | ? (Always 1) | none | none |
| 78 | SE_AbsorbMagicAtt | 'Spell Rune' Absorb spell damage until a maxium amount then fades | Rune Amt | none | none |
| 79 | SE_CurrentHPOnce | Direct Damage/Healing  | Amt (-)DD(+)Heal |  Target Restrictions  | Max Amt |
| 80 | SE_EnchantLight | UNKNOWN EFFECT [NOT] |  |  |  |
| 81 | SE_Revive | When cast on corpse restore experience and teleports player to corpse | Exp Amt | none | none |
| 82 | SE_SummonPC | Summon a player to caster | ? (Always 1) | none | none |
| 83 | SE_Teleport | Teleport Group to defined location.    (Corridinates base1=x base2=y base3=z base4=h) | Coordinates | none | none |
| 84 | SE_TossUp | Shoots player into the air. | Dist, (Always Neg) | none | none |
| 85 | SE_WeaponProc | Add melee proc | Proc Spell ID | Rate Modifier | none |
| 86 | SE_Harmony | Set NPC assist radius to spells value.  Mechanics  | Assist Radius | none | Max Level |
| 87 | SE_MagnifyVision | Zoom players vision | Magnify Amt | none | none |
| 88 | SE_Succor | Teleport Group/Self to defined location or to safe point in zone (2% fail rate) | Coord. (-1 = Safe) | none | none |
| 89 | SE_ModelSize | Change size by percent (grow/shrink) | Percent | ? | ? |
| 90 | SE_Cloak | Invisibility | Invisibile Level | none | none |
| 91 | SE_SummonCorpse | Summon targets corpse to caster. | Target Level Max | none | none |
| 92 | SE_InstantHate | Adds/Removes set amount of your 'hate' instantly from target | Hate Amt | none | ?Always(+) |
| 93 | SE_StopRain | Stops zone weather related rain | ? (Always 1) | none | none |
| 94 | SE_NegateIfCombat | Removes buff if player casts or does any combat skill | ? (Usually 1) | none | none |
| 95 | SE_Sacrifice | Kills player and creates 'Essence Emerald', corpse can not be rezed | ? (Always 1) | none | none |
| 96 | SE_Silence | Prevents spell casting | ? (Usually 1) | none | none |
| 97 | SE_ManaPool | Increase/Decrease max mana pool | Amt | none | Max Amt |
| 98 | SE_AttackSpeed2 | Stackable Haste/Slow that does not go over haste cap |  Attack Speed  | none | Max Amt |
| 99 | SE_Root | Immobilizes target.  Mechanics  | ? (Usually -10000) | none | none |
| 100 | SE_HealOverTime | Heal over time (will stack with an HoT from SE_CurrentHP) | Amt |  Target Restrictions  | Max Amt |
| 101 | SE_CompleteHeal | Heal for '7500' HP with a buff icon that blocks the same effect from taking hold. | Heal Multiplier (?) | none | ? |
| 102 | SE_Fearless | Immune to fear effect | ? (Always 1) | none | ? |
| 103 | SE_CallPet | Summon pet to owner | ? (Always 1) | none | ? |
| 104 | SE_Translocate | Creates a confirmation box to teleport player to location of bind point | Corridinates/Bind ID | none | none |
| 105 | SE_AntiGate | Prevent the use of gate spells.  (Base is likely the # of Bind Point ID's it can block) | ? (Seen 1 - 3) | none | none |
| 106 | SE_SummonBSTPet | Summon Pet                  ('Teleport Zone' Field in spells_new table contains name of pet) | ? (Always 1) | none | none |
| 107 | SE_AlterNPCLevel | +/- to NPC level, will revert to base line when effect fades. [Custom] [NOT] | Amt | none | none |
| 108 | SE_Familiar | Summon Pet                  ('Teleport Zone' Field in spells_new table contains name of pet) | ? (Usually 0) | none | none |
| 109 | SE_SummonItemIntoBag | If first effect is (SE_SummonItem) and item is bag, this will place items into bag | Item ID | none | none |
| 110 | SE_IncreaseArchery | UNKNOWN EFFECT [NOT] |  |  |  |
| 111 | SE_ResistAll | +/- all resist values | Amt | none | Max Amt |
| 112 | SE_CastingLevel | +/- 'Casting Level' which will determine fizzel rate | Amt | none | ?(Usually 0) |
| 113 | SE_SummonHorse | Summon a mount | ? (Always 1) | none | none |
| 114 | SE_ChangeAggro | +/- Percent modifier to spell and melee hate | Amt % | none | none |
| 115 | SE_Hunger | Prevents hunger/thirst checks (Ie. You shouldn't need food/drink with this effect) | ? (Always 1) | none | none |
| 116 | SE_CurseCounter | Determines potency of determental curse spells (+) or potency of cures (-) | Counter Amt | none | Max Amt |
| 117 | SE_MagicWeapon | Allows non-magic weapons to be considered 'magical' | ? (Always 1) | none | none |
| 118 | SE_Amplification | Increase modifier from singing skill | Amt | none | Max Amt |
| 119 | SE_AttackSpeed3 | Stackable Haste/Slow does go over haste cap. |  Attack Speed  | none | Max Amt |
| 120 | SE_HealRate | +/- Modfies by % the casters base heal value for incomming spells.   Heal Modifiers  | Amt % | none | none |
| 121 | SE_ReverseDS | Damage/Heal on entity with effect every time entity does a melee hit. | Amt | none | none |
| 122 | SE_ReduceSkill | UNKNOWN EFFECT [NOT] |  |  |  |
| 123 | SE_Screech | Spell Blocker (If have buff with value of +1 will block any spell w/ Screech w/ -1 value)  Buff Stacking  | Value (usually 1/-1) | none | none |
| 124 | SE_ImprovedDamage | Modifies by % casters base spell damage value  Spell Damage Modifiers  | Min % | Max % | none |
| 125 | SE_ImprovedHeal | Modifies by % the casters base heal value  Heal Modifiers  | Min % | Max % | none |
| 126 | SE_SpellResistReduction | Modifies by % chance for casters spell resist rate. | Min % | Max % | none |
| 127 | SE_IncreaseSpellHaste | Modifies by % spell casting time. | Amt % | none | none |
| 128 | SE_IncreaseSpellDuration | Modifies by % buff duration. | Amt % | none | none |
| 129 | SE_IncreaseRange | Modifies by % spell casting range. | Amt % | none | none |
| 130 | SE_SpellHateMod | Modifies by % hate from spells/abilities. | Min % | Max % | none |
| 131 | SE_ReduceReagentCost | Modifies by % chance not use a reagent. | Min % | Max % | none |
| 132 | SE_ReduceManaCost | Modifies by % spell mana cost. | Min % | Max % | none |
| 133 | SE_FcStunTimeMod | Modifies by % stun duration. [Custom] [NOT] | Amt % | none | none |
| 134 | SE_LimitMaxLevel | Limit to Max Level (% decrease is amount lost if over cap) | Max Level | % Decrease | none |
| 135 | SE_LimitResist | Limit to Resist Types, (+) Include (-) Exclude |  Resist Types  | none | none |
| 136 | SE_LimitTarget | Limit to Target Type. (+) Include (-) Exclude |  Target Types  | none | none |
| 137 | SE_LimitEffect | Limit to Spell Effect. (+) Include (-) Exclude | Spell Effect ID | none | none |
| 138 | SE_LimitSpellType | Limit to Beneficial(1) OR Determental(0) spells (goodEffect in spells_new ) | 1=Good 0=Bad | none | none |
| 139 | SE_LimitSpell | Limit to Spell ID (+) Include (-) Exclude | Spell ID | none | none |
| 140 | SE_LimitMinDur | Limit to spells with a minium duration. | Duration | none | none |
| 141 | SE_LimitInstant | Limit to spells that are instant cast. (1 = Instant Only) (0 = Exclude Instant) | Jan-00 | none | none |
| 142 | SE_LimitMinLevel | Limit to spells above a specific level. | Level | none | none |
| 143 | SE_LimitCastTimeMin | Limit to spells with a minium cast time. | Cast Time | none | none |
| 144 | SE_LimitCastTimeMax | Limit to spells with a maximum cast time. [Custom] [NOT] | Cast Time | none | none |
| 145 | SE_Teleport2 | Teleport to defined location. Used by 'Banisher' npcs, AoE spells that are cast on players. | Corridinates | none | none |
| 146 | SE_ElectricityResist | Electricity Resist? There is no resist type for this. | Amt | none | none |
| 147 | SE_PercentalHeal | Heal/(Damage) for percent value based on targets max HP | Amt(neg for dmg) | none | Max Amt |
| 148 | SE_StackingCommand_Block | Prevents buff from taking hold if criteria met. (SLOT = 'formula - 201')  Buff Stacking  | Spell Effect | none | &lt;  Amt |
| 149 | SE_StackingCommand_Overwrite | Allows buff from taking hold if criteria met. (SLOT = 'formula - 201')  Buff Stacking  | Spell Effect | none | &lt;  Amt |
| 150 | SE_DeathSave | If under 15% HP, this buff has chance to heal the owner.  Mechanics  | (1=Partial/2=Full) | Level Max | Heal Amt |
| 151 | SE_SuspendPet | Places a pet in temporary storage.  | ? (0/1) | none | none |
| 152 | SE_TemporaryPets | Creates a temporary pet that will fade after duration. (stacks with regular pets) | Amount of Pets | none | Duration |
| 153 | SE_BalanceHP | Balances groups HP (Penalty modifies the damage amount being distributed). | Penalty | Max HP taken/player | none |
| 154 | SE_DispelDetrimental | Dispels only detrimental effects.   Mechanics  | Level Modifer | none | none |
| 155 | SE_SpellCritDmgIncrease | Modifies by % critical spell damage amount.  Spell Damage Modifiers   [NOT]  | Crit Damage Mod % | none | none |
| 156 | SE_IllusionCopy | Turns caster into mirror image of target. | ?(Seen as 0/1/30) | none | ?(usually 0) |
| 157 | SE_SpellDamageShield | Casters will take damage from spell landing on target. | Amt (negative) | none | none |
| 158 | SE_Reflect | Reflects the casters spell back at the cast. | Reflect Chance % | ? (usuallly 0) | Max |
| 159 | SE_AllStats | +/- (STR, DEX, STA, CHA, WIS, INT) | Amt | none | none |
| 160 | SE_MakeDrunk | Gives client drunk effect if below tolerance level  (Effect currently handled entirely client side) | Tolerance | none | none |
| 161 | SE_MitigateSpellDamage | Reduces incomming spell damage by % up to a max value. | Mitigation % | Max Amt reduced | Rune Amt |
| 162 | SE_MitigateMeleeDamage | Reduces incomming melee damage by % up to a max value. | Mitigation % | Max Amt reduced | Rune Amt |
| 163 | SE_NegateAttacks | Complete or Partial block of melee / spell damage. (Max= Max Amt Dmg Blocked per hit) | Number of Blocks | none | Max Amt Blocked |
| 164 | SE_AppraiseLDonChest | Gives message if LDON chest is trapped / safe. | ? | none | App. Skill |
| 165 | SE_DisarmLDoNTrap | Attempts to disarm an LDON trap. | ? | none | Disarm Skill |
| 166 | SE_UnlockLDoNChest | Attemp to unlock an LDON chest | ? | none | Unlock Skill |
| 167 | SE_PetPowerIncrease | Increases statistics and level of the player's pet. | Power Level | none | none |
| 168 | SE_MeleeMitigation | Modifies melee damage by percent. (+)Take more DMG (-) Take less Damage | Mitigation % | none | none |
| 169 | SE_CriticalHitChance | Modifies melee critical hit chance by percent. (+) Increase Chance (-) Decrease Chance | Percent amount to increase critical hit chance. Note: this is a multiplier to critical hit chance, not an amount of percentage points to add to critical hit chance. |  Skills  (-1=ALL) | none |
| 170 | SE_SpellCritChance | Modifies by % chance to critical direct damage spells.  Spell Damage Modifiers  | Crit Chance % | none | none |
| 171 | SE_CrippBlowChance | Modifies crippling blow chance by percent. (Must have a critical hit chance to Crip) | Crip Blow Chance % | none | none |
| 172 | SE_AvoidMeleeChance | Modifies chance to avoid melee ('miss') (+) Increase Chance (-) Decrease Chance | Avoidance Chance % | none | none |
| 173 | SE_RiposteChance | Modifies chance to riposte (+) Increase (-) Decrease | Riposte Chance % | none | none |
| 174 | SE_DodgeChance | Modifies chance to dodge (+) Increase (-) Decrease | Dodge Chance % | none | none |
| 175 | SE_ParryChance | Modifies chance to parry (+) Increase (-) Decrease | Parry Chance % | none | none |
| 176 | SE_DualWieldChance | Modifies chance to dual wield (+) Increase (-) Decrease | DW Chance % | none | none |
| 177 | SE_DoubleAttackChance | Modifies chance to double attack (+) Increase (-) Decrease | Double Atk Chance % | none | none |
| 178 | SE_MeleeLifetap | (+) Heals you for a % of your melee damage done to target. (-) Dmgs you for % | Amt % (+/-) | none | none |
| 179 | SE_AllInstrumentMod | +/- Bard (Singing, Brass, Percusion, Wind, String) modifiers by %. | Amt % | none | none |
| 180 | SE_ResistSpellChance | Chance to resist any spell. | Resist Chance % | none | Max |
| 181 | SE_ResistFearChance | Chance to resist fear spells. | Resist Chance | none | none |
| 182 | SE_HundredHands | Modifies weapon delay by percent. (stacks with other hastes) | Amt % | none | none |
| 183 | SE_MeleeSkillCheck | Increases chance to hit (Unclear exactly how this works on live). | Amt % |  Skills (-1=ALL) | none |
| 184 | SE_HitChance | Modifies chance to hit with a specific skill. | Chance % |  Skills (-1=ALL) | none |
| 185 | SE_DamageModifier | Modifies damage amount by percent for a specific skill. | Amt % |  Skills (-1=ALL) | none |
| 186 | SE_MinDamageModifier | Modifies minimum damage amount by percent for a specific skill. | Amt % |  Skills (-1=ALL) | none |
| 187 | SE_BalanceMana | Balances groups mana. (Penalty modifies the mana amount being distributed). | Penalty | Max mana taken/pl | none |
| 188 | SE_IncreaseBlockChance | Modifies chance to block (+) Increase (-) Decrease | Block Chance % | none | none |
| 189 | SE_CurrentEndurance | +/- Instant endurance or over time (If duration is set) | Amt | none | Max Amt |
| 190 | SE_EndurancePool | +/- Total endurance pool. | Amt | none | Max Amt |
| 191 | SE_Amnesia | Silence verse melee abilities that use endurance / disciplines. | Usually 1 | none | none |
| 192 | SE_Hate | +/- Instant Hate or Hate over time if duration set. | Amt | none | none |
| 193 | SE_SkillAttack | Melee/Skill damage (utilizing all melee modifiers/bonuses) where base value is as if you were swinging a weapon with that damage amount using a specific skill. (skill field in spells_new defines which  Skills is used). | Weapon DMG Amt | Chance to Hit mod | UNKNOWN |
|  |  |  |  |  | Min DMG Amt? |
| 194 | SE_FadingMemories | To chance to be removeed from all hate lists and set to invisible. | Fade Chance | ? (Seen 0/75) | none |
| 195 | SE_StunResist | Chance to resist a stun from BASH/KICK. | Chance % | none | none |
| 196 | SE_Strikethrough | +/- Strikethrough chance (bypassing an opponent's special defenses, such as dodge, block, parry, and riposte.) | Chance % | none | none |
| 197 | SE_SkillDamageTaken | Modifies damage taken by % from specific skills (+) More dmg taken (-) Less dmg taken | Mitigation % |  Skills  (-1=ALL) | none |
| 198 | SE_CurrentEnduranceOnce | +/- Instant endurance. | Amt | none | none |
| 199 | SE_Taunt | Chance to taunt and instant hate | Chance % Taunt | Amt Hate | none |
| 200 | SE_ProcChance | Increase chance to proc from weapons. | Chance % | none | none |
| 201 | SE_RangedProc | Add spell proc to a ranged attack. | Proc Spell ID | Chance % | none |
| 202 | SE_IllusionOther | Allows next casted Illusion Buff (self-only) to be cast on a targeted player in group. | none | none | none |
| 203 | SE_MassGroupBuff | Allows next casted Group Buff to hit all players and pets within a large radius from caster at 2x mana cost. | 1 | none | none |
| 204 | SE_GroupFearImmunity | Provides immunity to fear for group. (Base * 10 = Duration) [No] | Duration | none | none |
| 205 | SE_Rampage | Does a single round of AE Melee attack (Set in EMU as distance 30 from caster). | ? (Always = 1) | none | none |
| 206 | SE_AETaunt | Area of Effect Taunt (Places caster top of hatelist +1 hate) Hardcoded 40 range | Hate Add to taunt | none | Range override |
| 207 | SE_FleshToBone | Turns Meat / Body parts items into bone chips | 1 | none | none |
| 208 | SE_PurgePoison | UNKNOWN  [NOT] |  |  |  |
| 209 | SE_DispelBeneficial | Dispels only beneficial effects.  Mechanics  | Level Modifier | none | none |
| 210 | SE_PetShield | Shield Effect (Share damage between pet and owner) for duration. | Duration (value * 12) | none | none |
| 211 | SE_AEMelee | Do an area of effect melee attack (ie. AE Rampage). Not implemented for NPC's | Duration (value * 12) | none | none |
| 212 | SE_FrenziedDevastation | Increase spell critical chance and mana cost 2x for DD spells. Spell Damage Modifiers  [NOT] | Critcal Chance | none | none |
| 213 | SE_PetMaxHP | +/- owner's pets Max HP by percent. | Amt % | none | none |
| 214 | SE_MaxHPChange | +/- Max HP by percent. | Amt % (value / 100) | none | none |
| 215 | SE_PetAvoidance | +/- owner's pets chance to avoid melee. | Amt | none | none |
| 216 | SE_Accuracy | Increase chance to hit by percent (15 Accuray = 1%) [Only] | Amt |  Skills (-1=ALL) | none |
| 217 | SE_HeadShot | Damage done by 'HeadShot' ability when proced  (Humaniod target hit w/ arrow) | ? (0) | Amt Dmg | none |
| 218 | SE_PetCriticalHit | +/- owner's pet/swarm pet chance to critical hit. | Amt  | none | none |
| 219 | SE_SlayUndead | Chance to do increased damage verse undead. (Chance = Value/1000) | Dmg Mod % | Chance | none |
| 220 | SE_SkillDamageAmount | Add flat amount of damage when a specific melee skill is used. | Amt |  Skills (-1=ALL) | none |
| 221 | SE_Packrat | +/- item weight reduction by percent. | Amt % | none | none |
| 222 | SE_BlockBehind | Modifies chance to block from behind (+) Increase (-) Decrease | Chance % | none | none |
| 223 | SE_DoubleRiposte | +/- Chance to do an additional riposte attack (after a successful riposte) [NOT] | Chance % | none | none |
| 224 | SE_GiveDoubleRiposte | +/- Chance to do an additional riposte attack (after a successful riposte) [Only] | Chance % |  Skills   | none |
| 225 | SE_GiveDoubleAttack | Allows any class to double attack at a set % chance or +/- chance if can innately DA. | Chance % | none | none |
| 226 | SE_TwoHandBash | Bash with a two handed weapon. (This works entirely client side) | ? (1) | none | none |
| 227 | SE_ReduceSkillTimer | +/- Reuse time on skill abilities (ie. Kick, Bash, Frenzy) | Time (seconds) |  Skills   | none |
| 228 | SE_ReduceFallDamage | Reduce the damage that you take from falling. |  |  |  |
| 229 | SE_PersistantCasting | Chance to continue casting while stunned. | Chance % | none | none |
| 230 | SE_ExtendedShielding | Increase range of /shield ability. |  |  |  |
| 231 | SE_StunBashChance | Modify chance to land a stun using BASH | Chance % | none | none |
| 232 | SE_DivineSave | Chance to avoid death.   Mechanics  | Chance % | Spell ID | none |
| 233 | SE_Metabolism | Modifies food / drink consumption rates. | Consumption Mod | none | none |
| 234 | SE_ReduceApplyPoisonTime | Reduces the time to apply poison |  |  |  |
| 235 | SE_ChannelChanceSpells | +/- Chance to channel a spell (avoid interupt).  [NOT] | Chance % | none | none |
| 236 | SE_FreePet | UKNONW EFFECT [NOT] |  |  |  |
| 237 | SE_GivePetGroupTarget | Pet Affinity, allow owner's pet to receive group buffs. | 1 | none | none |
| 238 | SE_IllusionPersistence | Persistence to your illusionions, causing them to last until you die or the illusion is forcibly removed. | 1 | none | none |
| 239 | SE_FeignedCastOnChance | Ability gives you an increasing chance for your feigned deaths to not be revealed by spells cast upon you. |  |  |  |
| 240 | SE_StringUnbreakable | Related to above. |  |  |  |
| 241 | SE_ImprovedReclaimEnergy | Modifies amount of mana returned from from SE_ReclaimPet | Amt % | none | none |
| 242 | SE_IncreaseChanceMemwipe | +/- Chance to successfully memblur target. | Amt % | none | none |
| 243 | SE_CharmBreakChance | Modifies charm break chance.  Mechanics  | Amt % | none | none |
| 244 | SE_RootBreakChance | Modifies chance of the casters root breaking.  Mechanics  | Chance % | none | none |
| 245 | SE_TrapCircumvention | Decreases the chance that you will set off a trap when opening a chest |  |  |  |
| 246 | SE_SetBreathLevel | Modifies lung capacity. | Amt | none | none |
| 247 | SE_RaiseSkillCap | Adds skill over the skill cap. |  |  |  |
| 248 | SE_SecondaryForte | Gives you a 2nd specialize skill that can go past 50 to 100. |  |  |  |
| 249 | SE_SecondaryDmgInc | Allows off hand weapon to recieve a damage bonus. | 1 | none | none |
| 250 | SE_SpellProcChance | Modify chance to do a proc from a 'proc spell buff | Chance % | none | none |
| 251 | SE_ConsumeProjectile | Chance NOT to consume a projectile (archery/throwing). | Chance % | none | none |
| 252 | SE_FrontalBackstabChance | Chance to perform a full backstab while in front of the target. | Chance % | none | none |
| 253 | SE_FrontalBackstabMinDmg | Allow a frontal backstab for mininum damage. | 1 | none | none |
| 254 | SE_Blank | Used as a spacer after last spell effect slot in spell file. | 0 | none | none |
| 255 | SE_ShieldDuration | Increases duration of /shield |  |  |  |
| 256 | SE_ShroudofStealth | Rogue improved invsible |  |  |  |
| 257 | SE_PetDiscipline | Give owner's pet, et /hold |  |  |  |
| 258 | SE_TripleBackstab | Chance to perform a triple backstab. | Chance % | none | none |
| 259 | SE_CombatStability | +/- AC Soft caps | Amt | none | none |
| 260 | SE_AddSingingMod | Add modifiers to bard instruments or singing abilities. (ItemType = Wind, Brass ect). | Amt | ItemType | none |
| 261 | SE_SongModCap | Raises cap of modifiers for bard instrument or singing abilities. [NOT] | Amt | none | none |
| 262 | SE_RaiseStatCap | +/- stat caps. (Base2=  STR: 0, STA: 1, AGI: 2, DEX: 3, WIS: 4, INT: 5, CHA: 6, MR: 7, CR: 8, FR: 9, PR: 10, DR: 11, COR: 12) | Amt | Stat Type | Max |
| 263 | SE_TradeSkillMastery |  Lets you raise more than one tradeskill above master. |  |  |  |
| 264 | SE_HastenedAASkill | Reduces reuse time on AA skills  [Use] |  |  |  |
| 265 | SE_MasteryofPast | Spell levels less than the base values level can not be fizzled. | Level | none | none |
| 266 | SE_ExtraAttackChance | Chance to do an extra attack with 2 Handed weapons only. | Chance % | none | none |
| 267 | SE_PetDiscipline2 | /pet focus, /pet no cast |  |  |  |
| 268 | SE_ReduceTradeskillFail | Reduces chance to fail with given tradeskill by a percent chance |  |  |  |
| 269 | SE_MaxBindWound | Increase max HP you can bind wound by percent | Amt % | none | none |
| 270 | SE_BardSongRange | +/- Range of bard beneficial songs. | Amt | none | none |
| 271 | SE_BaseMovementSpeed | Modifies basemove speed, doesn't stack with other move modfiers (Ie AA - Run 3) | Amt % | none | none |
| 272 | SE_CastingLevel2 | +/- Casting Level' which will determine fizzel rate.  | Amt  | none | none |
| 273 | SE_CriticalDoTChance | Modifies by % chance to critical Damage over time spells  Spell Damage Modifiers  | Crit Chance % | none | none |
| 274 | SE_CriticalHealChance | Modifies by % chance to critical heal spell  Heal Modifiers  | Crit Chance % | none | none |
| 275 | SE_CriticalMend | Chance to peform a critical mend (Monk Mend ability). | Crit Chance % | none | none |
| 276 | SE_Ambidexterity | Increase chance to duel weild by adding bonus 'duel weild skill' amount. | Skill Amt | none | none |
| 277 | SE_UnfailingDivinity | Gives second chance for SE_DivineSave to fire and if successful gives a modified heal amt.   Mechanics  | Heal Modifier | none | none |
| 278 | SE_FinishingBlow | Damage done by 'Finishing Blow' ability when proced  (Target &lt; 10% HP) | Chance (500 = 5%) | Amt Dmg | none |
| 279 | SE_Flurry | Chance to do a melee flurry. | Chance % | none | none |
| 280 | SE_PetFlurry | Chance for owner's pet or swarm pet to flurry. | Chance % | none | none |
| 281 | SE_FeignedMinion | Ability allows you to instruct your pet to feign death via the '/pet feign' command. (value = succeed chance) | Chance % |  |  |
| 282 | SE_ImprovedBindWound | Increase bind wound  healing amount by percent. | Amt % | none | none |
| 283 | SE_DoubleSpecialAttack | Chance to perform second special attack as monk. (ie Flying Kick, Tiger Claw) | Chance % | none | none |
| 284 | SE_LoHSetHeal | UKNONW EFFECT [NOT] |  |  |  |
| 285 | SE_NimbleEvasion | UKNONW EFFECT |  |  |  |
| 286 | SE_FcDamageAmt | +/- spell damage to casters spell.  Spell Damage Modifiers  | Amt | none | ? |
| 287 | SE_SpellDurationIncByTic | Increase spell duration by buff tick amount. | Amt Ticks | none | none |
| 288 | SE_SpecialAttackKBProc | Chance to to do a knockback spell proc from special attacks [AA] (Base Chance = 25%) | Chance Mod % |  Skills   | none |
| 289 | SE_CastOnFadeEffect | Triggers a spell only if fades after the full duration. (Typically on spells that can be cured) | Spell ID | none | none |
| 290 | SE_IncreaseRunSpeedCap | Increase run speed over the run speed cap. | Amt | none | none |
| 291 | SE_Purify | Chance to dispel all determental effects.   Mechanics  | Level Modifier | none | none |
| 292 | SE_StrikeThrough2 | +/- Strikethrough chance (bypassing an opponent's special defenses, such as dodge, block, parry, and riposte.) | Amt % | none | none |
| 293 | SE_FrontalStunResist | Chance to resist a stun from BASH/KICK if facing target. | Chance % | none | none |
| 294 | SE_CriticalSpellChance | Mod % chance to critical DD spells AND mod crit spell damage.  Spell Damage Modifiers  | Crit Chance % | Crit Damage Mod % | none |
| 295 | SE_ReduceTimerSpecial | UNKNOWN EFFECT [NOT] |  |  |  |
| 296 | SE_FcSpellVulnerability | (Debuff/Buff) Modifies spell damage by % on target. Spell Damage Modifiers  | Min % | Max % | none |
| 297 | SE_FcDamageAmtIncoming | (Debuff/Buff) Modifies spell/skill damage by flat amount on damage.  Spell Damage Modifiers  | Amt | none | ?(0 or=Amt) |
| 298 | SE_ChangeHeight | Shrink by percent. | Amt % | none | none |
| 299 | SE_WakeTheDead | Creates swarm pet from corpse's for a set duration (seconds)(Base likely amt pets) | ? (Seen 1 and 3) | none | Duration |
| 300 | SE_Doppelganger | Creates swarm pet that is a mirror image of caster | Amt Pets | none | Duration |
| 301 | SE_ArcheryDamageModifier | Modify archery damage by percent. | Amt % | none | none |
| 302 | SE_FcDamagePctCrit | Modifies by % casters base spell damage value (dmg can crit) Spell Damage Modifiers  | Min % | Max % | none |
| 303 | SE_FcDamageAmtCrit | +/- spell damage to casters spell (dmg can crit).  Spell Damage Modifiers  | Amt | none | ? |
| 304 | SE_OffhandRiposteFail | Chance for target not to riposte an an attack made from your off hand weapon. | Chance % | none | none |
| 305 | SE_MitigateDamageShield | Reduce incomming damage from damage shield using your off hand weapon. | Amt % | none | Max |
| 306 | SE_ArmyOfTheDead | This ability calls up to five shades of nearby corpses back to life to serve the necromancer.  |  |  |  |
| 307 | SE_Appraisal | his ability allows you to estimate the selling price of an item you are holding on your cursor. |  |  |  |
| 308 | SE_SuspendMinion | Store a backup pet that can be unsuspended. | 1 | none | none |
| 309 | SE_GateCastersBindPoint | Teleports group to casters bind point.          Bind Point ID (1=Primary, 2=Secondary 3=Tertiary) | Bind Point ID | none | none |
| 310 | SE_ReduceReuseTimer | Reduce the reuse timer on disciplines by seconds. (Base value set in milliseconds) | Amt Time | none | none |
| 311 | SE_LimitCombatSkills | Limit to exclude (discs and combat procs = 0) or (spells = 1) | 0/1 | none | none |
| 312 | SE_Sanctuary | Places caster at bottom of all hate lists, effect fades if caster cast spell on targets other than self. | ? (1) | none | none |
| 313 | SE_ForageAdditionalItems | Chance to forage additional items using 'forage' ability. | Chance % | none | none |
| 314 | SE_Invisibility2 | Invisibility (Stable)      Invisible Level corresponds to what type of see invisible will detect. | Invisible Level | none | none |
| 315 | SE_InvisVsUndead2 | Invisibility vs Undead (Stable)      Invisible Level corresponds to what type of see invisible will detect. | Invisible Level | none | none |
| 316 | SE_ImprovedInvisAnimals | UNKNOWN EFFECT [NOT] |  |  |  |
| 317 | SE_ItemHPRegenCapIncrease | Increase HP regen from items over cap. | Amt | none | none |
| 318 | SE_ItemManaRegenCapIncrease | Increase Mana regen from items over cap. | Amt | none | none |
| 319 | SE_CriticalHealOverTime | Modifies % chance to do a critical heal over time  Heal Modifiers  | Crit Chance % | none | none |
| 320 | SE_ShieldBlock | Chance to block an attack while shield is equiped. | Chance % | none | none |
| 321 | SE_ReduceHate | +/- Instant hate | Amt | none | none |
| 322 | SE_GateToHomeCity | Gate to original starting city. | 1 | none | none |
| 323 | SE_DefensiveProc | Add a chance to proc from any incoming melee swing. | Spell ID | Proc Rate Mod | ? |
| 324 | SE_HPToMana | Casted spells will use HP instead of Mana with a coversion penalty rate. | Conversion Rate % | none | none |
| 325 | SE_ChanceInvsBreakToAoE | [AA] Increasing chance to remain hidden when they are an indirect target of an AoE spell. |  |  |  |
| 326 | SE_SpellSlotIncrease | Increase physical amount of spell slots. |  |  |  |
| 327 | SE_MysticalAttune | Increases max amount of buffs a player can have. | Amt | none | none |
| 328 | SE_DelayDeath | Increase the amount of HP under zero that can be lost before actual death occurs. | Amt HP | none | none |
| 329 | SE_ManaAbsorbPercentDamage | Reduces incoming damage by % and converts that amount to mana loss. | Mitigation % | none | none |
| 330 | SE_CriticalDamageMob | Modifies damage done from a critical melee hit. | Dmg Modifier % |  Skills (-1=ALL) | none |
| 331 | SE_Salvage | Increase chance to salvage from tradeskills by percent. | Chance % | none | none |
| 332 | SE_SummonToCorpse | [AA] Ressurect spell with no penalty or exp gain (Can still res after). | ? | none | none |
| 333 | SE_CastOnRuneFadeEffect | Triggers a spell when a spell with a rune amount is used up and fades. | Spell ID | none | none |
| 334 | SE_BardAEDot | Area of effect damage over time song, damage only done if target is NOT moving. | Amt +/- | none | none |
| 335 | SE_BlockNextSpellFocus | Chance to block next spell that meets the focus limits defined with this effect in it. | Chance % | none | none |
| 336 | SE_IllusionaryTarget | UNKNOWN EFFECT [NOT] |  |  |  |
| 337 | SE_PercentXPIncrease | Modify amount of experience gained. | Amt % | none | none |
| 338 | SE_SummonAndResAllCorpses | Summon and ressurect all corpses for 100% exp. | ? (Seen=70) | none | none |
| 339 | SE_TriggerOnCast | Cast an additional spell on spell use if the spell casted meets focus limits. | Chance % | Spell ID | none |
| 340 | SE_SpellTrigger | Chance to add an additional spell to the target. (If multiple effects in same spell where % add up to 100, then one effect must fire) | Chance % | Spell ID | none |
| 341 | SE_ItemAttackCapIncrease | Increase the cap to the amount of attack that can gained from items. | Amt | none | none |
| 342 | SE_ImmuneFleeing | Prevent NPC from fleeing at low health. | 1 | none | none |
| 343 | SE_InterruptCasting | Chance to interrupt targets spell casting can be instant or per buff tick. | Chance % | none | none |
| 344 | SE_ChannelChanceItems | Modify chance to successfully channel from item click casts. | Chance % | none | none |
| 345 | SE_AssassinationLevel | Maximum level of target that 'Assassination' will proc on. | Level | none | none |
| 346 | SE_HeadShotLevel | Maximum level of target that 'HeadShot' will proc on. | Level | none | none |
| 347 | SE_DoubleRangedAttack | Chance to perform a double ranged attack (Throw/Archery) will consume projectile. | Chance % | none | none |
| 348 | SE_LimitManaMin | Limit to spells above a minium amount of mana. | Amt | none | none |
| 349 | SE_ShieldEquipHateMod | Hate modifier to if shield equiped. (Shield Specialist AA) This may not be correct | Amt % | none | none |
| 350 | SE_ManaBurn | Drains mana for damage/heal at a defined ratio up to a defined maximum amount of mana. Ratio (-) | Max Mana | Mana/HP Ratio / 10 | none |
| 351 | SE_PersistentEffect | Aura effects, grants a buff automatically to those within a radius of caster. | ? | none | none |
| 352 | SE_IncreaseTrapCount | UKNOWN EFFECT |  |  |  |
| 353 | SE_AdditionalAura | Increase number of aura slots. | Amt Slots | none | none |
| 354 | SE_DeactivateAllTraps | UKNOWN EFFECT |  |  |  |
| 355 | SE_LearnTrap | UKNOWN EFFECT |  |  |  |
| 356 | SE_ChangeTriggerType | UKNOWN EFFECT [NOT] |  |  |  |
| 357 | SE_FcMute | Chance to silence casting of spells that contain specific spell effects. (Effects determined by focus limits) | Chance % | none | none |
| 358 | SE_CurrentManaOnce | +/- Instant mana | Amt | none | none |
| 359 | SE_PassiveSenseTrap | UKNOWN EFFECT |  |  |  |
| 360 | SE_ProcOnKillShot | Triggers a spell after a kill shot. (Max = Min Level of Target) | Chance % | Spell ID | Min Level |
| 361 | SE_SpellOnDeath | Triggers when the owner of the buff is killed. | Chance % | Spell ID | none |
| 362 | SE_PotionBeltSlots | 'Quick Draw' expands the potion belt by one additional available item slot per rank. |  |  |  |
| 363 | SE_BandolierSlots | 'Battle Ready' expands the bandolier by one additional save slot per rank. |  |  |  |
| 364 | SE_TripleAttackChance | +/- chance to triple attack. | Chance % | none | none |
| 365 | SE_ProcOnSpellKillShot | Chance to trigger a spell on kill when the kill is caused by the specific spell with this effect in it. | Chance % | Spell ID | Min Level |
| 366 | SE_ShieldEquipDmgMod | Damage modifier to melee if shield equiped. (Shield Specialist AA) This may not be correct | Amt % | ? | none |
| 367 | SE_SetBodyType | Change body type of target. |  Body_Types  | none | none |
| 368 | SE_FactionMod | Increases faction with base1 (faction id, live won't match up w/ ours) by base2. | Faction ID | Faction Mod | none |
| 369 | SE_CorruptionCounter | Determines potency of determental curse spells (+) or potency of cures (-) | Counter Amt | none | Max |
| 370 | SE_ResistCorruption | +/- Corruption Resist | Amt | none | none |
| 371 | SE_AttackSpeed4 | Stackable slow effect |  Attack Speed  | none | Amt Max |
| 372 | SE_ForageSkill |  |  |  |  |
| 373 | SE_CastOnFadeEffectAlways | Triggers a spell when buff fades after natural duration OR from rune/numhits fades. | Spell ID | none | none |
| 374 | SE_ApplyEffect | Chance to add an additional spell to the target. | Chance % | Spell ID | none |
| 375 | SE_DotCritDmgIncrease | Modifies by % critical dot damage.  Spell Damage Modifiers  | Amt % | none | none |
| 376 | SE_Fling | UNKNOWN EFFECT | ? | ? | none |
| 377 | SE_CastOnFadeEffectNPC | Triggers a spell when buff fades after natural duration. (On live these are in player spells cast on NPCs) | Spell ID | none | none |
| 378 | SE_SpellEffectResistChance | Chance to resist a specific spell effect type (Ie. 5% chance to Resist DD spells) | % Chance | Spell Effect ID | none |
| 379 | SE_ShadowStepDirectional | Push forward for specific amount of units. (Handled client side) | Distance Units | none | none |
| 380 | SE_Knockdown | Push back and up effect. (Handled client side) | Push Back | Push Up | none |
| 381 | SE_KnockTowardCaster | Summon a player to caster at a set distance away from caster. (Uses a reverse knockback effect) | Distance Units | none | none |
| 382 | SE_NegateSpellEffect | Completely remove a specified spell effect bonus. (Ie. Remove all Aggro bonus) | ? | Spell Effect ID | none |
| 383 | SE_SympatheticProc | Chance to proc a a spell off of a regularly casted spell. (Typically found as item focus) | Chance Mod | Spell ID | none  |
| 384 | SE_Leap | Jump toward your target. | Distance | none | none |
| 385 | SE_LimitSpellGroup | Limit focus effects by spell group. (spellgroup field in spells_new) (+) Include (-) Exclude | Spellgroup ID | none | none |
| 386 | SE_CastOnCurer | Trigger a spell on yourself when you cure a target. | Spell ID | none | none |
| 387 | SE_CastOnCure | Trigger a spell on yourself when you are cured. | Spell ID | none | none |
| 388 | SE_SummonCorpseZone | Summon a corpse from any zone. | 1 | none | none |
| 389 | SE_FcTimerRefresh | Reset recast timers (ungrey spell gems) | 1 | none | none |
| 390 | SE_FcTimerLockout | Increase time on reset timer. | Amt | none | none |
| 391 | SE_MeleeVulnerability | +/- melee mitigation (+) weakness (-) mitigation (Live SPA defined SE_LimitManaMax is not correct) | Amt % | none | none |
| 392 | SE_FcHealAmt | +/- flat amount to casters heal spells  Heal Modifiers  | Amt | none | none |
| 393 | SE_FcHealPctIncoming | (Buff/Debuff) Modfies by % the casters base heal value for incomming spells.  Heal Modifiers  | Amt % | none | none |
| 394 | SE_FcHealAmtIncoming | (Buff/Debuff) +/- flat amount the casters base heal value for incomming spells. Heal Modifiers  | Amt | none | none |
| 395 | SE_FcHealPctCritIncoming | (Buff/Debuff) Modifies by % chance to receive a critical heal on incomming spells.  Heal Modifiers  | Amt % | none | none |
| 396 | SE_FcHealAmtCrit | +/- flat amount to casters heal spells (amt can critical)  Heal Modifiers  | Amt | none | none |
| 397 | SE_PetMeleeMitigation | +/- AC to owner's pet. | Amt | none | none |
| 398 | SE_SwarmPetDuration | Increases duration of swarm pets by seconds. (Base value set in miliseconds) | Time | none | none |
| 399 | SE_FcTwincast | Chance to cast the same spell 2x from a single cast. | Chance % | none | none |
| 400 | SE_HealGroupFromMana | Group heal, drains mana from caster and coverts that to the amount of HP healed at a defined ratio. | Max Mana Drain | Mana/HP Ratio / 10 | none |
| 401 | SE_ManaDrainWithDmg | Drains targets mana and decreases hit points based on a defined ratio of hp/mana drained. | Max Mana Drained | HP/Mana Ratio / 10 | Max ? |
| 402 | SE_EndDrainWithDmg | Drains targets endurance and decreases hit points based on a defined ratio of hp/endur drained. | Max Endur Drained | HP/Endur Ratio / 10 | none |
| 403 | SE_LimitSpellClass | Limits to specific types of spell categories. (3=Cures,3=Offensive, 6=Lifetap) (+) Include (-) Exclude | Category | none | none |
| 404 | SE_LimitSpellSubclass | Limits to specific types of spell categories. (UNDEFINED) (+) Include (-) Exclude | Category | none | none |
| 405 | SE_TwoHandBluntBlock | Modifies chance to block if 2 Hand Blunt equiped (+) Increase (-) Decrease | Chance % | none | none |
| 406 | SE_CastonNumHitFade | Triggers a spell when a spells numhit counter is depleted. | Spell ID | none | none |
| 407 | SE_CastonFocusEffect | Triggers a spell if focus limits are met (ie Triggers when a focus effects is applied) | Spell ID | none | noen |
| 408 | SE_LimitHPPercent | Caps maximum HP to % or a defined amount, which ever is lowest. | Cap % | Cap Amt | none |
| 409 | SE_LimitManaPercent | Caps maximum Mana to % or a defined amount, which ever is lowest. | Cap % | Cap Amt | none |
| 410 | SE_LimitEndPercent | Caps maximum Endurance to % or a defined amount, which ever is lowest. | Cap % | Cap Amt | none |
| 411 | SE_LimitClass | Limits to spells of a certain player class (Uses Bitmask, the class value in spell dbase is 1 bitmask higher in relation to item class value) | Class Bitmask | none | none |
| 412 | SE_LimitRace | Limits to spells cast by a certain race [NOT] | Race ID | none | none |
| 413 | SE_FcBaseEffects | Modifies base values of certain spell effects.  Mechanics  Partially implemented | Amt % | none | none |
| 414 | SE_LimitCastingSkill | Limit to spells that use a specific casting skill. (skill field in spells new) (+) Include (-) Exclude |  Skills  | none | none |
| 415 | SE_FFItemClass | UKNOWN EFFECT [NOT] |  |  |  |
| 416 | SE_ACv2 | +/- AC (stacks with other AC buffs) | Amt | none | Max |
| 417 | SE_ManaRegen_v2 | +/- mana regen (stacks with other mana regen buffs) | Amt | none | Max |
| 418 | SE_SkillDamageAmount2 | Add flat amount of damage when a specific melee skill is used. | Amt |  Skills (-1=ALL) | none |
| 419 | SE_AddMeleeProc | Add melee proc | Proc Spell ID | Rate Modifier | none |
| 420 | SE_FcLimitUse | Increase numhits count by percent. [Custom] [NOT] | Amt % | none | none |
| 421 | SE_FcIncreaseNumHits | Increase numhits count by flat amount. | Amt | none | none |
| 422 | SE_LimitUseMin | Limit to spells with a minimum number of numhit counters.  [Custom] [NOT] | Amt | none | none |
| 423 | SE_LimitUseType | Limit to spells by  Numhit_Types  [Custom] [NOT] |  Numhit_Types  | none | none |
| 424 | SE_GravityEffect | Pulls target(s) toward caster at a set pace to a specific distance away from caster. | Distance From Caster | Force of Pull | none |
| 425 | SE_Display | Gives illusion flying dragon (unclear how this works) | 1 | none | none |
| 426 | SE_IncreaseExtTargetWindow | Increases the capacity of your extended target window |  |  |  |
| 427 | SE_SkillProc | Chance to proc a spell when using a specific skill (ie Proc from a Taunt or Kick). | Spell ID | Rate Modifier | none |
| 428 | SE_LimitToSkill | Limits what skill will effect a SkillProc. (Always place as next effect after ID 427/429) |  Skills   | none | none |
| 429 | SE_SkillProcSuccess | Chance to proc a spell when using a specific skill if it successfully hits the target. | Spell ID | Rate Modifier | none |
| 430 | SE_PostEffect | Alter vision ? UNKNOWN EFFECT | ? | ? | ? |
| 431 | SE_PostEffectData | Tint vision ? UNKNOWN EFFECT | RGB | ? | ? |
| 432 | SE_ExpandMaxActiveTrophyBen | UKNOWN EFFECTâ€‹ [NOT] |  |  |  |
| 433 | SE_CriticalDotDecay | Chance to critical DoT with effect depreciation.  Spell Damage Modifiers (Removed from live 7-14 ?) | Chance % | Decay % | Max Level |
| 434 | SE_CriticalHealDecay | Chance to critical Heal with effect depreciation.  Heal Modifiers (Removed from live 7-14 ?) | Chance % | Decay % | Max Level |
| 435 | SE_CriticalRegenDecay | Chance to critical Regen with effect depreciation.  Heal Modifiers (Removed from live 7-14 ?) | Chance % | Decay % | Max Level |
| 436 | SE_BeneficialCountDownHold | UKNOWN EFFECTâ€‹ |  |  |  |
| 437 | SE_TeleporttoAnchor | Teleport Guild Hall Anchor | ? | none | none |
| 438 | SE_TranslocatetoAnchor | Translocate Primary Anchor | ? | none | none |
| 439 | SE_Assassination | Damage done by 'Assassination' ability when proced  (Humaniod target hit w/ Backstab) | ? (0) | Amt Dmg | none |
| 440 | SE_FinishingBlowLvl | Maximum level of target that 'Finishing Blow' will proc on. | Level | ? (Seen 200) | none |
| 441 | SE_DistanceRemoval | Fades buff if owner of buff moves specified amount of distance from location where buff was applied. | Distance | none | none |
| 442 | SE_TriggerOnReqTarget | Triggers a spell on target when a specific condition is met on that target. Buff fades after trigger. | Spell ID |  Target Condition  | none |
| 443 | SE_TriggerOnReqCaster | Triggers a spell on target when a specific condition is met on that target. Buff fades after trigger. (All spells that use this are self only) | Spell ID |  Target Condition  | none |
| 444 | SE_ImprovedTaunt | Locks aggro on caster and decreases other players aggro by % up to a specified level. | Max level | Aggro Mod | none |
| 445 | SE_AddMercSlot | [AA] Allows you to conscript additional mercs. | Amt ? | none | none |
| 446 | SE_AStacker | Buff stacking blocker  Buff Stacking  | Stacking Priority | none | none |
| 447 | SE_BStacker | Buff stacking blocker  Buff Stacking  | Stacking Priority | none | none |
| 448 | SE_CStacker | Buff stacking blocker  Buff Stacking  | Stacking Priority | none | none |
| 449 | SE_DStacker | Buff stacking blocker  Buff Stacking  | Stacking Priority | none | none |
| 450 | SE_MitigateDotDamage | Reduces incomming dotl damage by % up to a max value. | Mitigation % | Max Amt Reduced | Rune Amt |
| 451 | SE_MeleeThresholdGuard | Partial Melee Rune that only is lowered if melee hits are over a defined amount (limit) of damage | Mitigation % | Min Hit | Rune Amt |
| 452 | SE_SpellThresholdGuard | Partial Spell Rune that only is lowered if spell dmg is over a defined amount (limit) of damage | Mitigation % | Min Hit | Rune Amt |
| 453 | SE_TriggerMeleeThreshold | Trigger spell when specified amount of melee damage is taken in a single hit, then fade buff. | Spell ID | Damage Amt | none |
| 454 | SE_TriggerSpellThreshold | Trigger spell when specified amount of spell damage is taken in a single hit, then fade buff. | Spell ID | Damage Amt | none |
| 455 | SE_AddHatePct | Modifies amount of hate you have on target by percent over instantly, | Amt % | none | none |
| 456 | SE_AddHateOverTimePct | Modifies amount of hate you have on target by percent over time, | Amt % | none | none |
| 457 | SE_ResourceTap | Coverts a percent of spell dmg from dmg spells (DD/DOT) to hp/mana/end. | % Coversion | 0=H/1=M/2=E | Max Amt |
| 458 | SE_FactionModPct | Modifies faction gains and losses by percent. | Amt % | none | none |
| 459 | SE_DamageModifier2 | Modifies damage amount by percent for a specific skill. | Amt % | Skills (-1=ALL) |  |

