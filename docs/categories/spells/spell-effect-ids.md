# Spell Effect IDs



|  |  |  |  |  |  |
| :--- | :--- | :--- | :--- | :--- | :--- |
| ID | Spell Effect | Description | Base Value | Limit Value | Max Value |
| 0 | SE\_CurrentHP | Direct Damage/Healing, duration allows for HoT/DoT | Amt \(-\)DD\(+\)Heal | Target Restrictions  | Max Amt |
| 1 | SE\_ArmorClass | +/- AC | Amt | none | Max Amt |
| 2 | SE\_ATK | +/- ATK | Amt | none | Max Amt |
| 3 | SE\_MovementSpeed | +/- Movement Speed | Amt | none | Max Amt |
| 4 | SE\_STR | +/- STR | Amt | none | Max Amt |
| 5 | SE\_DEX | +/- DEX | Amt | none | Max Amt |
| 6 | SE\_AGI | +/- AGI | Amt | none | Max Amt |
| 7 | SE\_STA | +/- STA | Amt | none | Max Amt |
| 8 | SE\_INT | +/- INT | Amt | none | Max Amt |
| 9 | SE\_WIS | +/- WIS | Amt | none | Max Amt |
| 10 | SE\_CHA | +/- CHA \(BLANK SPACER if 0 in spell effect\) | Amt | none | Max Amt |
| 11 | SE\_AttackSpeed | &gt;100/&lt;100 Attack Speed \(standard 'haste or 'slow\) |  Attack Speed  | none | Max Amt |
| 12 | SE\_Invisibility | Invisibility \(Unstable\)         Invisible Level corresponds to what type of see invisible will detect. | Invisible Level | none | none |
| 13 | SE\_SeeInvis | See Invisibility                                                              Invisible Lv \(1\) = Standard invsibile  | Invisible Level | none | none |
| 14 | SE\_WaterBreathing | Immune to drowning | ? \(Seen from 1 - 3\) | none | none |
| 15 | SE\_CurrentMana | Direct +/- mana, duration allows for mana over time | Amt | none | Max Amt |
| 16 | SE\_NPCFrenzy | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 17 | SE\_NPCAwareness | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 18 | SE\_Lull | Seen in conjunction with SE\_ChangeFrenzyRad and SE\_Harmony, this effect does nothing alone. | 0 | none | none |
| 19 | SE\_AddFaction | Adjusts NPC 'con' \(Alliance spell line\) | Faction Amt | none | none |
| 20 | SE\_Blind | Turn PC screen blank, causes NPC to flee if NOT in melee range | ? \(Usually -1/1\) | none | none |
| 21 | SE\_Stun | Stun for a duration | Stun Time | ? \(Always &lt;= Base\) | Max Level |
| 22 | SE\_Charm | Control an NPC/Player for a duration.   Mechanics  | ? | none | Max Level |
| 23 | SE\_Fear | Cause your target to flee.   Mechanics  | ? | none | Max Level |
| 24 | SE\_Stamina | Increased stamina \(yellow bar\) 'Invigor'  - ? Current Effect \(Does not raise endurance on live\) | ? \(Always Negative\) | none | none |
| 25 | SE\_BindAffinity | Sets Bind point.                                     Bind Point ID \(1=Primary, 2=Secondary 3=Tertiary\)  | Bind Point ID | none | none |
| 26 | SE\_Gate | Transport to Bind point                           Bind Point ID \(1=Primary, 2=Secondary 3=Tertiary\)  | Success Chance | Bind Point ID | none |
| 27 | SE\_CancelMagic | Chance to remove any buff effect.   Mechanics  | Level Modifier | none | none |
| 28 | SE\_InvisVsUndead | Invisiblity vs Undead \(Unstable\)                                          \(Only Invisiblity Level \(1\) exists\) | Invisible Level | none | none |
| 29 | SE\_InvisVsAnimals | Invisiblity vs Animal \(Unstable\)                                          \(Only Invisiblity Level \(1\) exists\) | Invisible Level | none | none |
| 30 | SE\_ChangeFrenzyRad | Sets NPC Aggro Radius to the value of the spell effect Mechanics  | Aggro Radius | none | Max Level |
| 31 | SE\_Mez | Stuns target till duration ends or broken \(melee/spell dmg\)  Mechanics  | Stacking Value | none | Max Level |
| 32 | SE\_SummonItem | Summon an item. | Item ID | none | none |
| 33 | SE\_SummonPet | Summon Pet                  \('Teleport Zone' Field in spells\_new table contains name of pet\) | ? \(Always 1\) | none | none |
| 34 | SE\_Confuse | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 35 | SE\_DiseaseCounter | Determines potency of determental disease spells \(+\) or potency of cures \(-\) | Counter Amt | none | Max Amt |
| 36 | SE\_PoisonCounter | Determines potency of determental poision spells \(+\) or potency of cures \(-\) | Counter Amt | none | Max Amt |
| 37 | SE\_DetectHostile | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 38 | SE\_DetectMagic | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 39 | SE\_DetectPoison | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 40 | SE\_DivineAura | Invulnerable to spells and melee, you can not cast or melee while under this effect. | ? \(Usually 1\) | none | none |
| 41 | SE\_Destroy | Instantly kill target | ? \(Always 1\) | none | none |
| 42 | SE\_ShadowStep | Warps player a short distance in a random direction | ? \(Usually 1\) | none | none |
| 43 | SE\_Berserk | Sets client as 'Berserk' allowing for chance to crippling blow. \[Custom\] \[NOT\] | 1 | none | none |
| 44 | SE\_Lycanthropy | Used as a place holder effect for preventing certain buffs from stacking | Stacking Value | none | none |
| 45 | SE\_Vampirism | Heals you for a percent of your melee damage done to target.  \[Custom\] \[NOT\] | Amt % \(+/-\) | none | none |
| 46 | SE\_ResistFire | +/- Fire Resist | Amt | none | Max Amt |
| 47 | SE\_ResistCold | +/- Cold Resist | Amt | none | Max Amt |
| 48 | SE\_ResistPoison | +/- Poison Resist | Amt | none | Max Amt |
| 49 | SE\_ResistDisease | +/- Disease Resist | Amt | none | Max Amt |
| 50 | SE\_ResistMagic | +/- Magic Resist | Amt | none | Max Amt |
| 51 | SE\_DetectTraps | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 52 | SE\_SenseDead | Point player in direction of nearest Undead NPC | ? \(Always 1\) | none | ? |
| 53 | SE\_SenseSummoned | Point player in direction of nearest Summoned NPC | ? \(Always 1\) | none | none |
| 54 | SE\_SenseAnimals | Point player in direction of nearest Animal NPC | ? \(Always 1\) | none | none |
| 55 | SE\_Rune | Absorb melee damage until a maxium amount then fades | Rune Amt | none | Max Amt |
| 56 | SE\_TrueNorth | Point player in North direction. | ? \(Always 1\) | none | none |
| 57 | SE\_Levitate | Levitation \(Take no fall damage\) | ? \(Usually 1\) | none | none |
| 58 | SE\_Illusion | Changes visual race of target | Race ID \(-1=Gender\) | Texture ID | Helm ID |
| 59 | SE\_DamageShield | Damage taken upon successful melee hit | Amt | none | Max Amt |
| 60 | SE\_TransferItem | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 61 | SE\_Identify | Displays in text the 'lore' field from the items table | ? \(Always 1\) | none | none |
| 62 | SE\_ItemID | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 63 | SE\_WipeHateList | Chance to wipe the NPC's hatelist | Success Chance | none | Max Chance |
| 64 | SE\_SpinTarget | Spins Target and stuns | Spin Duration | ? \(Always &lt;= base\) | Max Level |
| 65 | SE\_InfraVision | Improved night vision | ? \(Always 1\) | none | none |
| 66 | SE\_UltraVision | Better improved night vision | ? \(Always 1\) | none | none |
| 67 | SE\_EyeOfZomm | Transfers your vision and control to a temporary NPC | ? \(Always 1\) | none | none |
| 68 | SE\_ReclaimPet | Kills your pet in exchange for mana. \(Returns 75% of pet spell actual mana cost\) | ? \(Always 1\) | none | none |
| 69 | SE\_TotalHP | Increases Max HP \(standard 'HP Buffs'\) | Amt | none | Max Amt |
| 70 | SE\_CorpseBomb | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 71 | SE\_NecPet | Summon Pet                  \('Teleport Zone' Field in spells\_new table contains name of pet\) |  |  |  |
| 72 | SE\_PreserveCorpse | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 73 | SE\_BindSight | See from the perspective of the target that this is cast on. | ? \(Always 1\) | none | none |
| 74 | SE\_FeignDeath | Fall to the ground and clear hate list | Success Chance | ? \(usually 0\) | Max Chance |
| 75 | SE\_VoiceGraft | Speak through your pet | ? \(Always 1\) | none | none |
| 76 | SE\_Sentinel | Creates a proximity zone where cast that alerts caster if NPC's or Players enter it. | ? \(Always 1\) | none | none |
| 77 | SE\_LocateCorpse | Turn player in direction of targeted players corpse | ? \(Always 1\) | none | none |
| 78 | SE\_AbsorbMagicAtt | 'Spell Rune' Absorb spell damage until a maxium amount then fades | Rune Amt | none | none |
| 79 | SE\_CurrentHPOnce | Direct Damage/Healing  | Amt \(-\)DD\(+\)Heal |  Target Restrictions  | Max Amt |
| 80 | SE\_EnchantLight | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 81 | SE\_Revive | When cast on corpse restore experience and teleports player to corpse | Exp Amt | none | none |
| 82 | SE\_SummonPC | Summon a player to caster | ? \(Always 1\) | none | none |
| 83 | SE\_Teleport | Teleport Group to defined location.    \(Corridinates base1=x base2=y base3=z base4=h\) | Coordinates | none | none |
| 84 | SE\_TossUp | Shoots player into the air. | Dist, \(Always Neg\) | none | none |
| 85 | SE\_WeaponProc | Add melee proc | Proc Spell ID | Rate Modifier | none |
| 86 | SE\_Harmony | Set NPC assist radius to spells value.  Mechanics  | Assist Radius | none | Max Level |
| 87 | SE\_MagnifyVision | Zoom players vision | Magnify Amt | none | none |
| 88 | SE\_Succor | Teleport Group/Self to defined location or to safe point in zone \(2% fail rate\) | Coord. \(-1 = Safe\) | none | none |
| 89 | SE\_ModelSize | Change size by percent \(grow/shrink\) | Percent | ? | ? |
| 90 | SE\_Cloak | Invisibility | Invisibile Level | none | none |
| 91 | SE\_SummonCorpse | Summon targets corpse to caster. | Target Level Max | none | none |
| 92 | SE\_InstantHate | Adds/Removes set amount of your 'hate' instantly from target | Hate Amt | none | ?Always\(+\) |
| 93 | SE\_StopRain | Stops zone weather related rain | ? \(Always 1\) | none | none |
| 94 | SE\_NegateIfCombat | Removes buff if player casts or does any combat skill | ? \(Usually 1\) | none | none |
| 95 | SE\_Sacrifice | Kills player and creates 'Essence Emerald', corpse can not be rezed | ? \(Always 1\) | none | none |
| 96 | SE\_Silence | Prevents spell casting | ? \(Usually 1\) | none | none |
| 97 | SE\_ManaPool | Increase/Decrease max mana pool | Amt | none | Max Amt |
| 98 | SE\_AttackSpeed2 | Stackable Haste/Slow that does not go over haste cap |  Attack Speed  | none | Max Amt |
| 99 | SE\_Root | Immobilizes target.  Mechanics  | ? \(Usually -10000\) | none | none |
| 100 | SE\_HealOverTime | Heal over time \(will stack with an HoT from SE\_CurrentHP\) | Amt |  Target Restrictions  | Max Amt |
| 101 | SE\_CompleteHeal | Heal for '7500' HP with a buff icon that blocks the same effect from taking hold. | Heal Multiplier \(?\) | none | ? |
| 102 | SE\_Fearless | Immune to fear effect | ? \(Always 1\) | none | ? |
| 103 | SE\_CallPet | Summon pet to owner | ? \(Always 1\) | none | ? |
| 104 | SE\_Translocate | Creates a confirmation box to teleport player to location of bind point | Corridinates/Bind ID | none | none |
| 105 | SE\_AntiGate | Prevent the use of gate spells.  \(Base is likely the \# of Bind Point ID's it can block\) | ? \(Seen 1 - 3\) | none | none |
| 106 | SE\_SummonBSTPet | Summon Pet                  \('Teleport Zone' Field in spells\_new table contains name of pet\) | ? \(Always 1\) | none | none |
| 107 | SE\_AlterNPCLevel | +/- to NPC level, will revert to base line when effect fades. \[Custom\] \[NOT\] | Amt | none | none |
| 108 | SE\_Familiar | Summon Pet                  \('Teleport Zone' Field in spells\_new table contains name of pet\) | ? \(Usually 0\) | none | none |
| 109 | SE\_SummonItemIntoBag | If first effect is \(SE\_SummonItem\) and item is bag, this will place items into bag | Item ID | none | none |
| 110 | SE\_IncreaseArchery | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 111 | SE\_ResistAll | +/- all resist values | Amt | none | Max Amt |
| 112 | SE\_CastingLevel | +/- 'Casting Level' which will determine fizzel rate | Amt | none | ?\(Usually 0\) |
| 113 | SE\_SummonHorse | Summon a mount | ? \(Always 1\) | none | none |
| 114 | SE\_ChangeAggro | +/- Percent modifier to spell and melee hate | Amt % | none | none |
| 115 | SE\_Hunger | Prevents hunger/thirst checks \(Ie. You shouldn't need food/drink with this effect\) | ? \(Always 1\) | none | none |
| 116 | SE\_CurseCounter | Determines potency of determental curse spells \(+\) or potency of cures \(-\) | Counter Amt | none | Max Amt |
| 117 | SE\_MagicWeapon | Allows non-magic weapons to be considered 'magical' | ? \(Always 1\) | none | none |
| 118 | SE\_Amplification | Increase modifier from singing skill | Amt | none | Max Amt |
| 119 | SE\_AttackSpeed3 | Stackable Haste/Slow does go over haste cap. |  Attack Speed  | none | Max Amt |
| 120 | SE\_HealRate | +/- Modfies by % the casters base heal value for incomming spells.   Heal Modifiers  | Amt % | none | none |
| 121 | SE\_ReverseDS | Damage/Heal on entity with effect every time entity does a melee hit. | Amt | none | none |
| 122 | SE\_ReduceSkill | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 123 | SE\_Screech | Spell Blocker \(If have buff with value of +1 will block any spell w/ Screech w/ -1 value\)  Buff Stacking  | Value \(usually 1/-1\) | none | none |
| 124 | SE\_ImprovedDamage | Modifies by % casters base spell damage value  Spell Damage Modifiers  | Min % | Max % | none |
| 125 | SE\_ImprovedHeal | Modifies by % the casters base heal value  Heal Modifiers  | Min % | Max % | none |
| 126 | SE\_SpellResistReduction | Modifies by % chance for casters spell resist rate. | Min % | Max % | none |
| 127 | SE\_IncreaseSpellHaste | Modifies by % spell casting time. | Amt % | none | none |
| 128 | SE\_IncreaseSpellDuration | Modifies by % buff duration. | Amt % | none | none |
| 129 | SE\_IncreaseRange | Modifies by % spell casting range. | Amt % | none | none |
| 130 | SE\_SpellHateMod | Modifies by % hate from spells/abilities. | Min % | Max % | none |
| 131 | SE\_ReduceReagentCost | Modifies by % chance not use a reagent. | Min % | Max % | none |
| 132 | SE\_ReduceManaCost | Modifies by % spell mana cost. | Min % | Max % | none |
| 133 | SE\_FcStunTimeMod | Modifies by % stun duration. \[Custom\] \[NOT\] | Amt % | none | none |
| 134 | SE\_LimitMaxLevel | Limit to Max Level \(% decrease is amount lost if over cap\) | Max Level | % Decrease | none |
| 135 | SE\_LimitResist | Limit to Resist Types, \(+\) Include \(-\) Exclude |  Resist Types  | none | none |
| 136 | SE\_LimitTarget | Limit to Target Type. \(+\) Include \(-\) Exclude |  Target Types  | none | none |
| 137 | SE\_LimitEffect | Limit to Spell Effect. \(+\) Include \(-\) Exclude | Spell Effect ID | none | none |
| 138 | SE\_LimitSpellType | Limit to Beneficial\(1\) OR Determental\(0\) spells \(goodEffect in spells\_new \) | 1=Good 0=Bad | none | none |
| 139 | SE\_LimitSpell | Limit to Spell ID \(+\) Include \(-\) Exclude | Spell ID | none | none |
| 140 | SE\_LimitMinDur | Limit to spells with a minium duration. | Duration | none | none |
| 141 | SE\_LimitInstant | Limit to spells that are instant cast. \(1 = Instant Only\) \(0 = Exclude Instant\) | Jan-00 | none | none |
| 142 | SE\_LimitMinLevel | Limit to spells above a specific level. | Level | none | none |
| 143 | SE\_LimitCastTimeMin | Limit to spells with a minium cast time. | Cast Time | none | none |
| 144 | SE\_LimitCastTimeMax | Limit to spells with a maximum cast time. \[Custom\] \[NOT\] | Cast Time | none | none |
| 145 | SE\_Teleport2 | Teleport to defined location. Used by 'Banisher' npcs, AoE spells that are cast on players. | Corridinates | none | none |
| 146 | SE\_ElectricityResist | Electricity Resist? There is no resist type for this. | Amt | none | none |
| 147 | SE\_PercentalHeal | Heal/\(Damage\) for percent value based on targets max HP | Amt\(neg for dmg\) | none | Max Amt |
| 148 | SE\_StackingCommand\_Block | Prevents buff from taking hold if criteria met. \(SLOT = 'formula - 201'\)  Buff Stacking  | Spell Effect | none | &lt;  Amt |
| 149 | SE\_StackingCommand\_Overwrite | Allows buff from taking hold if criteria met. \(SLOT = 'formula - 201'\)  Buff Stacking  | Spell Effect | none | &lt;  Amt |
| 150 | SE\_DeathSave | If under 15% HP, this buff has chance to heal the owner.  Mechanics  | \(1=Partial/2=Full\) | Level Max | Heal Amt |
| 151 | SE\_SuspendPet | Places a pet in temporary storage.  | ? \(0/1\) | none | none |
| 152 | SE\_TemporaryPets | Creates a temporary pet that will fade after duration. \(stacks with regular pets\) | Amount of Pets | none | Duration |
| 153 | SE\_BalanceHP | Balances groups HP \(Penalty modifies the damage amount being distributed\). | Penalty | Max HP taken/player | none |
| 154 | SE\_DispelDetrimental | Dispels only detrimental effects.   Mechanics  | Level Modifer | none | none |
| 155 | SE\_SpellCritDmgIncrease | Modifies by % critical spell damage amount.  Spell Damage Modifiers   \[NOT\]  | Crit Damage Mod % | none | none |
| 156 | SE\_IllusionCopy | Turns caster into mirror image of target. | ?\(Seen as 0/1/30\) | none | ?\(usually 0\) |
| 157 | SE\_SpellDamageShield | Casters will take damage from spell landing on target. | Amt \(negative\) | none | none |
| 158 | SE\_Reflect | Reflects the casters spell back at the cast. | Reflect Chance % | ? \(usuallly 0\) | Max |
| 159 | SE\_AllStats | +/- \(STR, DEX, STA, CHA, WIS, INT\) | Amt | none | none |
| 160 | SE\_MakeDrunk | Gives client drunk effect if below tolerance level  \(Effect currently handled entirely client side\) | Tolerance | none | none |
| 161 | SE\_MitigateSpellDamage | Reduces incomming spell damage by % up to a max value. | Mitigation % | Max Amt reduced | Rune Amt |
| 162 | SE\_MitigateMeleeDamage | Reduces incomming melee damage by % up to a max value. | Mitigation % | Max Amt reduced | Rune Amt |
| 163 | SE\_NegateAttacks | Complete or Partial block of melee / spell damage. \(Max= Max Amt Dmg Blocked per hit\) | Number of Blocks | none | Max Amt Blocked |
| 164 | SE\_AppraiseLDonChest | Gives message if LDON chest is trapped / safe. | ? | none | App. Skill |
| 165 | SE\_DisarmLDoNTrap | Attempts to disarm an LDON trap. | ? | none | Disarm Skill |
| 166 | SE\_UnlockLDoNChest | Attemp to unlock an LDON chest | ? | none | Unlock Skill |
| 167 | SE\_PetPowerIncrease | Increases statistics and level of the player's pet. | Power Level | none | none |
| 168 | SE\_MeleeMitigation | Modifies melee damage by percent. \(+\)Take more DMG \(-\) Take less Damage | Mitigation % | none | none |
| 169 | SE\_CriticalHitChance | Modifies melee critical hit chance by percent. \(+\) Increase Chance \(-\) Decrease Chance | Percent amount to increase critical hit chance. Note: this is a multiplier to critical hit chance, not an amount of percentage points to add to critical hit chance. |  Skills  \(-1=ALL\) | none |
| 170 | SE\_SpellCritChance | Modifies by % chance to critical direct damage spells.  Spell Damage Modifiers  | Crit Chance % | none | none |
| 171 | SE\_CrippBlowChance | Modifies crippling blow chance by percent. \(Must have a critical hit chance to Crip\) | Crip Blow Chance % | none | none |
| 172 | SE\_AvoidMeleeChance | Modifies chance to avoid melee \('miss'\) \(+\) Increase Chance \(-\) Decrease Chance | Avoidance Chance % | none | none |
| 173 | SE\_RiposteChance | Modifies chance to riposte \(+\) Increase \(-\) Decrease | Riposte Chance % | none | none |
| 174 | SE\_DodgeChance | Modifies chance to dodge \(+\) Increase \(-\) Decrease | Dodge Chance % | none | none |
| 175 | SE\_ParryChance | Modifies chance to parry \(+\) Increase \(-\) Decrease | Parry Chance % | none | none |
| 176 | SE\_DualWieldChance | Modifies chance to dual wield \(+\) Increase \(-\) Decrease | DW Chance % | none | none |
| 177 | SE\_DoubleAttackChance | Modifies chance to double attack \(+\) Increase \(-\) Decrease | Double Atk Chance % | none | none |
| 178 | SE\_MeleeLifetap | \(+\) Heals you for a % of your melee damage done to target. \(-\) Dmgs you for % | Amt % \(+/-\) | none | none |
| 179 | SE\_AllInstrumentMod | +/- Bard \(Singing, Brass, Percusion, Wind, String\) modifiers by %. | Amt % | none | none |
| 180 | SE\_ResistSpellChance | Chance to resist any spell. | Resist Chance % | none | Max |
| 181 | SE\_ResistFearChance | Chance to resist fear spells. | Resist Chance | none | none |
| 182 | SE\_HundredHands | Modifies weapon delay by percent. \(stacks with other hastes\) | Amt % | none | none |
| 183 | SE\_MeleeSkillCheck | Increases chance to hit \(Unclear exactly how this works on live\). | Amt % |  Skills \(-1=ALL\) | none |
| 184 | SE\_HitChance | Modifies chance to hit with a specific skill. | Chance % |  Skills \(-1=ALL\) | none |
| 185 | SE\_DamageModifier | Modifies damage amount by percent for a specific skill. | Amt % |  Skills \(-1=ALL\) | none |
| 186 | SE\_MinDamageModifier | Modifies minimum damage amount by percent for a specific skill. | Amt % |  Skills \(-1=ALL\) | none |
| 187 | SE\_BalanceMana | Balances groups mana. \(Penalty modifies the mana amount being distributed\). | Penalty | Max mana taken/pl | none |
| 188 | SE\_IncreaseBlockChance | Modifies chance to block \(+\) Increase \(-\) Decrease | Block Chance % | none | none |
| 189 | SE\_CurrentEndurance | +/- Instant endurance or over time \(If duration is set\) | Amt | none | Max Amt |
| 190 | SE\_EndurancePool | +/- Total endurance pool. | Amt | none | Max Amt |
| 191 | SE\_Amnesia | Silence verse melee abilities that use endurance / disciplines. | Usually 1 | none | none |
| 192 | SE\_Hate | +/- Instant Hate or Hate over time if duration set. | Amt | none | none |
| 193 | SE\_SkillAttack | Melee/Skill damage \(utilizing all melee modifiers/bonuses\) where base value is as if you were swinging a weapon with that damage amount using a specific skill. \(skill field in spells\_new defines which  Skills is used\). | Weapon DMG Amt | Chance to Hit mod | UNKNOWN |
|  |  |  |  |  | Min DMG Amt? |
| 194 | SE\_FadingMemories | To chance to be removeed from all hate lists and set to invisible. | Fade Chance | ? \(Seen 0/75\) | none |
| 195 | SE\_StunResist | Chance to resist a stun from BASH/KICK. | Chance % | none | none |
| 196 | SE\_Strikethrough | +/- Strikethrough chance \(bypassing an opponent's special defenses, such as dodge, block, parry, and riposte.\) | Chance % | none | none |
| 197 | SE\_SkillDamageTaken | Modifies damage taken by % from specific skills \(+\) More dmg taken \(-\) Less dmg taken | Mitigation % |  Skills  \(-1=ALL\) | none |
| 198 | SE\_CurrentEnduranceOnce | +/- Instant endurance. | Amt | none | none |
| 199 | SE\_Taunt | Chance to taunt and instant hate | Chance % Taunt | Amt Hate | none |
| 200 | SE\_ProcChance | Increase chance to proc from weapons. | Chance % | none | none |
| 201 | SE\_RangedProc | Add spell proc to a ranged attack. | Proc Spell ID | Chance % | none |
| 202 | SE\_IllusionOther | Allows next casted Illusion Buff \(self-only\) to be cast on a targeted player in group. | none | none | none |
| 203 | SE\_MassGroupBuff | Allows next casted Group Buff to hit all players and pets within a large radius from caster at 2x mana cost. | 1 | none | none |
| 204 | SE\_GroupFearImmunity | Provides immunity to fear for group. \(Base \* 10 = Duration\) \[No\] | Duration | none | none |
| 205 | SE\_Rampage | Does a single round of AE Melee attack \(Set in EMU as distance 30 from caster\). | ? \(Always = 1\) | none | none |
| 206 | SE\_AETaunt | Area of Effect Taunt \(Places caster top of hatelist +1 hate\) Hardcoded 40 range | Hate Add to taunt | none | Range override |
| 207 | SE\_FleshToBone | Turns Meat / Body parts items into bone chips | 1 | none | none |
| 208 | SE\_PurgePoison | UNKNOWN  \[NOT\] |  |  |  |
| 209 | SE\_DispelBeneficial | Dispels only beneficial effects.  Mechanics  | Level Modifier | none | none |
| 210 | SE\_PetShield | Shield Effect \(Share damage between pet and owner\) for duration. | Duration \(value \* 12\) | none | none |
| 211 | SE\_AEMelee | Do an area of effect melee attack \(ie. AE Rampage\). Not implemented for NPC's | Duration \(value \* 12\) | none | none |
| 212 | SE\_FrenziedDevastation | Increase spell critical chance and mana cost 2x for DD spells. Spell Damage Modifiers  \[NOT\] | Critcal Chance | none | none |
| 213 | SE\_PetMaxHP | +/- owner's pets Max HP by percent. | Amt % | none | none |
| 214 | SE\_MaxHPChange | +/- Max HP by percent. | Amt % \(value / 100\) | none | none |
| 215 | SE\_PetAvoidance | +/- owner's pets chance to avoid melee. | Amt | none | none |
| 216 | SE\_Accuracy | Increase chance to hit by percent \(15 Accuray = 1%\) \[Only\] | Amt |  Skills \(-1=ALL\) | none |
| 217 | SE\_HeadShot | Damage done by 'HeadShot' ability when proced  \(Humaniod target hit w/ arrow\) | ? \(0\) | Amt Dmg | none |
| 218 | SE\_PetCriticalHit | +/- owner's pet/swarm pet chance to critical hit. | Amt  | none | none |
| 219 | SE\_SlayUndead | Chance to do increased damage verse undead. \(Chance = Value/1000\) | Dmg Mod % | Chance | none |
| 220 | SE\_SkillDamageAmount | Add flat amount of damage when a specific melee skill is used. | Amt |  Skills \(-1=ALL\) | none |
| 221 | SE\_Packrat | +/- item weight reduction by percent. | Amt % | none | none |
| 222 | SE\_BlockBehind | Modifies chance to block from behind \(+\) Increase \(-\) Decrease | Chance % | none | none |
| 223 | SE\_DoubleRiposte | +/- Chance to do an additional riposte attack \(after a successful riposte\) \[NOT\] | Chance % | none | none |
| 224 | SE\_GiveDoubleRiposte | +/- Chance to do an additional riposte attack \(after a successful riposte\) \[Only\] | Chance % |  Skills   | none |
| 225 | SE\_GiveDoubleAttack | Allows any class to double attack at a set % chance or +/- chance if can innately DA. | Chance % | none | none |
| 226 | SE\_TwoHandBash | Bash with a two handed weapon. \(This works entirely client side\) | ? \(1\) | none | none |
| 227 | SE\_ReduceSkillTimer | +/- Reuse time on skill abilities \(ie. Kick, Bash, Frenzy\) | Time \(seconds\) |  Skills   | none |
| 228 | SE\_ReduceFallDamage | Reduce the damage that you take from falling. |  |  |  |
| 229 | SE\_PersistantCasting | Chance to continue casting while stunned. | Chance % | none | none |
| 230 | SE\_ExtendedShielding | Increase range of /shield ability. |  |  |  |
| 231 | SE\_StunBashChance | Modify chance to land a stun using BASH | Chance % | none | none |
| 232 | SE\_DivineSave | Chance to avoid death.   Mechanics  | Chance % | Spell ID | none |
| 233 | SE\_Metabolism | Modifies food / drink consumption rates. | Consumption Mod | none | none |
| 234 | SE\_ReduceApplyPoisonTime | Reduces the time to apply poison |  |  |  |
| 235 | SE\_ChannelChanceSpells | +/- Chance to channel a spell \(avoid interupt\).  \[NOT\] | Chance % | none | none |
| 236 | SE\_FreePet | UKNONW EFFECT \[NOT\] |  |  |  |
| 237 | SE\_GivePetGroupTarget | Pet Affinity, allow owner's pet to receive group buffs. | 1 | none | none |
| 238 | SE\_IllusionPersistence | Persistence to your illusionions, causing them to last until you die or the illusion is forcibly removed. | 1 | none | none |
| 239 | SE\_FeignedCastOnChance | Ability gives you an increasing chance for your feigned deaths to not be revealed by spells cast upon you. |  |  |  |
| 240 | SE\_StringUnbreakable | Related to above. |  |  |  |
| 241 | SE\_ImprovedReclaimEnergy | Modifies amount of mana returned from from SE\_ReclaimPet | Amt % | none | none |
| 242 | SE\_IncreaseChanceMemwipe | +/- Chance to successfully memblur target. | Amt % | none | none |
| 243 | SE\_CharmBreakChance | Modifies charm break chance.  Mechanics  | Amt % | none | none |
| 244 | SE\_RootBreakChance | Modifies chance of the casters root breaking.  Mechanics  | Chance % | none | none |
| 245 | SE\_TrapCircumvention | Decreases the chance that you will set off a trap when opening a chest |  |  |  |
| 246 | SE\_SetBreathLevel | Modifies lung capacity. | Amt | none | none |
| 247 | SE\_RaiseSkillCap | Adds skill over the skill cap. |  |  |  |
| 248 | SE\_SecondaryForte | Gives you a 2nd specialize skill that can go past 50 to 100. |  |  |  |
| 249 | SE\_SecondaryDmgInc | Allows off hand weapon to recieve a damage bonus. | 1 | none | none |
| 250 | SE\_SpellProcChance | Modify chance to do a proc from a 'proc spell buff | Chance % | none | none |
| 251 | SE\_ConsumeProjectile | Chance NOT to consume a projectile \(archery/throwing\). | Chance % | none | none |
| 252 | SE\_FrontalBackstabChance | Chance to perform a full backstab while in front of the target. | Chance % | none | none |
| 253 | SE\_FrontalBackstabMinDmg | Allow a frontal backstab for mininum damage. | 1 | none | none |
| 254 | SE\_Blank | Used as a spacer after last spell effect slot in spell file. | 0 | none | none |
| 255 | SE\_ShieldDuration | Increases duration of /shield |  |  |  |
| 256 | SE\_ShroudofStealth | Rogue improved invsible |  |  |  |
| 257 | SE\_PetDiscipline | Give owner's pet, et /hold |  |  |  |
| 258 | SE\_TripleBackstab | Chance to perform a triple backstab. | Chance % | none | none |
| 259 | SE\_CombatStability | +/- AC Soft caps | Amt | none | none |
| 260 | SE\_AddSingingMod | Add modifiers to bard instruments or singing abilities. \(ItemType = Wind, Brass ect\). | Amt | ItemType | none |
| 261 | SE\_SongModCap | Raises cap of modifiers for bard instrument or singing abilities. \[NOT\] | Amt | none | none |
| 262 | SE\_RaiseStatCap | +/- stat caps. \(Base2=  STR: 0, STA: 1, AGI: 2, DEX: 3, WIS: 4, INT: 5, CHA: 6, MR: 7, CR: 8, FR: 9, PR: 10, DR: 11, COR: 12\) | Amt | Stat Type | Max |
| 263 | SE\_TradeSkillMastery |  Lets you raise more than one tradeskill above master. |  |  |  |
| 264 | SE\_HastenedAASkill | Reduces reuse time on AA skills  \[Use\] |  |  |  |
| 265 | SE\_MasteryofPast | Spell levels less than the base values level can not be fizzled. | Level | none | none |
| 266 | SE\_ExtraAttackChance | Chance to do an extra attack with 2 Handed weapons only. | Chance % | none | none |
| 267 | SE\_PetDiscipline2 | /pet focus, /pet no cast |  |  |  |
| 268 | SE\_ReduceTradeskillFail | Reduces chance to fail with given tradeskill by a percent chance |  |  |  |
| 269 | SE\_MaxBindWound | Increase max HP you can bind wound by percent | Amt % | none | none |
| 270 | SE\_BardSongRange | +/- Range of bard beneficial songs. | Amt | none | none |
| 271 | SE\_BaseMovementSpeed | Modifies basemove speed, doesn't stack with other move modfiers \(Ie AA - Run 3\) | Amt % | none | none |
| 272 | SE\_CastingLevel2 | +/- Casting Level' which will determine fizzel rate.  | Amt  | none | none |
| 273 | SE\_CriticalDoTChance | Modifies by % chance to critical Damage over time spells  Spell Damage Modifiers  | Crit Chance % | none | none |
| 274 | SE\_CriticalHealChance | Modifies by % chance to critical heal spell  Heal Modifiers  | Crit Chance % | none | none |
| 275 | SE\_CriticalMend | Chance to peform a critical mend \(Monk Mend ability\). | Crit Chance % | none | none |
| 276 | SE\_Ambidexterity | Increase chance to duel weild by adding bonus 'duel weild skill' amount. | Skill Amt | none | none |
| 277 | SE\_UnfailingDivinity | Gives second chance for SE\_DivineSave to fire and if successful gives a modified heal amt.   Mechanics  | Heal Modifier | none | none |
| 278 | SE\_FinishingBlow | Damage done by 'Finishing Blow' ability when proced  \(Target &lt; 10% HP\) | Chance \(500 = 5%\) | Amt Dmg | none |
| 279 | SE\_Flurry | Chance to do a melee flurry. | Chance % | none | none |
| 280 | SE\_PetFlurry | Chance for owner's pet or swarm pet to flurry. | Chance % | none | none |
| 281 | SE\_FeignedMinion | Ability allows you to instruct your pet to feign death via the '/pet feign' command. \(value = succeed chance\) | Chance % |  |  |
| 282 | SE\_ImprovedBindWound | Increase bind wound  healing amount by percent. | Amt % | none | none |
| 283 | SE\_DoubleSpecialAttack | Chance to perform second special attack as monk. \(ie Flying Kick, Tiger Claw\) | Chance % | none | none |
| 284 | SE\_LoHSetHeal | UKNONW EFFECT \[NOT\] |  |  |  |
| 285 | SE\_NimbleEvasion | UKNONW EFFECT |  |  |  |
| 286 | SE\_FcDamageAmt | +/- spell damage to casters spell.  Spell Damage Modifiers  | Amt | none | ? |
| 287 | SE\_SpellDurationIncByTic | Increase spell duration by buff tick amount. | Amt Ticks | none | none |
| 288 | SE\_SpecialAttackKBProc | Chance to to do a knockback spell proc from special attacks \[AA\] \(Base Chance = 25%\) | Chance Mod % |  Skills   | none |
| 289 | SE\_CastOnFadeEffect | Triggers a spell only if fades after the full duration. \(Typically on spells that can be cured\) | Spell ID | none | none |
| 290 | SE\_IncreaseRunSpeedCap | Increase run speed over the run speed cap. | Amt | none | none |
| 291 | SE\_Purify | Chance to dispel all determental effects.   Mechanics  | Level Modifier | none | none |
| 292 | SE\_StrikeThrough2 | +/- Strikethrough chance \(bypassing an opponent's special defenses, such as dodge, block, parry, and riposte.\) | Amt % | none | none |
| 293 | SE\_FrontalStunResist | Chance to resist a stun from BASH/KICK if facing target. | Chance % | none | none |
| 294 | SE\_CriticalSpellChance | Mod % chance to critical DD spells AND mod crit spell damage.  Spell Damage Modifiers  | Crit Chance % | Crit Damage Mod % | none |
| 295 | SE\_ReduceTimerSpecial | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 296 | SE\_FcSpellVulnerability | \(Debuff/Buff\) Modifies spell damage by % on target. Spell Damage Modifiers  | Min % | Max % | none |
| 297 | SE\_FcDamageAmtIncoming | \(Debuff/Buff\) Modifies spell/skill damage by flat amount on damage.  Spell Damage Modifiers  | Amt | none | ?\(0 or=Amt\) |
| 298 | SE\_ChangeHeight | Shrink by percent. | Amt % | none | none |
| 299 | SE\_WakeTheDead | Creates swarm pet from corpse's for a set duration \(seconds\)\(Base likely amt pets\) | ? \(Seen 1 and 3\) | none | Duration |
| 300 | SE\_Doppelganger | Creates swarm pet that is a mirror image of caster | Amt Pets | none | Duration |
| 301 | SE\_ArcheryDamageModifier | Modify archery damage by percent. | Amt % | none | none |
| 302 | SE\_FcDamagePctCrit | Modifies by % casters base spell damage value \(dmg can crit\) Spell Damage Modifiers  | Min % | Max % | none |
| 303 | SE\_FcDamageAmtCrit | +/- spell damage to casters spell \(dmg can crit\).  Spell Damage Modifiers  | Amt | none | ? |
| 304 | SE\_OffhandRiposteFail | Chance for target not to riposte an an attack made from your off hand weapon. | Chance % | none | none |
| 305 | SE\_MitigateDamageShield | Reduce incomming damage from damage shield using your off hand weapon. | Amt % | none | Max |
| 306 | SE\_ArmyOfTheDead | This ability calls up to five shades of nearby corpses back to life to serve the necromancer.  |  |  |  |
| 307 | SE\_Appraisal | his ability allows you to estimate the selling price of an item you are holding on your cursor. |  |  |  |
| 308 | SE\_SuspendMinion | Store a backup pet that can be unsuspended. | 1 | none | none |
| 309 | SE\_GateCastersBindPoint | Teleports group to casters bind point.          Bind Point ID \(1=Primary, 2=Secondary 3=Tertiary\) | Bind Point ID | none | none |
| 310 | SE\_ReduceReuseTimer | Reduce the reuse timer on disciplines by seconds. \(Base value set in milliseconds\) | Amt Time | none | none |
| 311 | SE\_LimitCombatSkills | Limit to exclude \(discs and combat procs = 0\) or \(spells = 1\) | 0/1 | none | none |
| 312 | SE\_Sanctuary | Places caster at bottom of all hate lists, effect fades if caster cast spell on targets other than self. | ? \(1\) | none | none |
| 313 | SE\_ForageAdditionalItems | Chance to forage additional items using 'forage' ability. | Chance % | none | none |
| 314 | SE\_Invisibility2 | Invisibility \(Stable\)      Invisible Level corresponds to what type of see invisible will detect. | Invisible Level | none | none |
| 315 | SE\_InvisVsUndead2 | Invisibility vs Undead \(Stable\)      Invisible Level corresponds to what type of see invisible will detect. | Invisible Level | none | none |
| 316 | SE\_ImprovedInvisAnimals | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 317 | SE\_ItemHPRegenCapIncrease | Increase HP regen from items over cap. | Amt | none | none |
| 318 | SE\_ItemManaRegenCapIncrease | Increase Mana regen from items over cap. | Amt | none | none |
| 319 | SE\_CriticalHealOverTime | Modifies % chance to do a critical heal over time  Heal Modifiers  | Crit Chance % | none | none |
| 320 | SE\_ShieldBlock | Chance to block an attack while shield is equiped. | Chance % | none | none |
| 321 | SE\_ReduceHate | +/- Instant hate | Amt | none | none |
| 322 | SE\_GateToHomeCity | Gate to original starting city. | 1 | none | none |
| 323 | SE\_DefensiveProc | Add a chance to proc from any incoming melee swing. | Spell ID | Proc Rate Mod | ? |
| 324 | SE\_HPToMana | Casted spells will use HP instead of Mana with a coversion penalty rate. | Conversion Rate % | none | none |
| 325 | SE\_ChanceInvsBreakToAoE | \[AA\] Increasing chance to remain hidden when they are an indirect target of an AoE spell. |  |  |  |
| 326 | SE\_SpellSlotIncrease | Increase physical amount of spell slots. |  |  |  |
| 327 | SE\_MysticalAttune | Increases max amount of buffs a player can have. | Amt | none | none |
| 328 | SE\_DelayDeath | Increase the amount of HP under zero that can be lost before actual death occurs. | Amt HP | none | none |
| 329 | SE\_ManaAbsorbPercentDamage | Reduces incoming damage by % and converts that amount to mana loss. | Mitigation % | none | none |
| 330 | SE\_CriticalDamageMob | Modifies damage done from a critical melee hit. | Dmg Modifier % |  Skills \(-1=ALL\) | none |
| 331 | SE\_Salvage | Increase chance to salvage from tradeskills by percent. | Chance % | none | none |
| 332 | SE\_SummonToCorpse | \[AA\] Ressurect spell with no penalty or exp gain \(Can still res after\). | ? | none | none |
| 333 | SE\_CastOnRuneFadeEffect | Triggers a spell when a spell with a rune amount is used up and fades. | Spell ID | none | none |
| 334 | SE\_BardAEDot | Area of effect damage over time song, damage only done if target is NOT moving. | Amt +/- | none | none |
| 335 | SE\_BlockNextSpellFocus | Chance to block next spell that meets the focus limits defined with this effect in it. | Chance % | none | none |
| 336 | SE\_IllusionaryTarget | UNKNOWN EFFECT \[NOT\] |  |  |  |
| 337 | SE\_PercentXPIncrease | Modify amount of experience gained. | Amt % | none | none |
| 338 | SE\_SummonAndResAllCorpses | Summon and ressurect all corpses for 100% exp. | ? \(Seen=70\) | none | none |
| 339 | SE\_TriggerOnCast | Cast an additional spell on spell use if the spell casted meets focus limits. | Chance % | Spell ID | none |
| 340 | SE\_SpellTrigger | Chance to add an additional spell to the target. \(If multiple effects in same spell where % add up to 100, then one effect must fire\) | Chance % | Spell ID | none |
| 341 | SE\_ItemAttackCapIncrease | Increase the cap to the amount of attack that can gained from items. | Amt | none | none |
| 342 | SE\_ImmuneFleeing | Prevent NPC from fleeing at low health. | 1 | none | none |
| 343 | SE\_InterruptCasting | Chance to interrupt targets spell casting can be instant or per buff tick. | Chance % | none | none |
| 344 | SE\_ChannelChanceItems | Modify chance to successfully channel from item click casts. | Chance % | none | none |
| 345 | SE\_AssassinationLevel | Maximum level of target that 'Assassination' will proc on. | Level | none | none |
| 346 | SE\_HeadShotLevel | Maximum level of target that 'HeadShot' will proc on. | Level | none | none |
| 347 | SE\_DoubleRangedAttack | Chance to perform a double ranged attack \(Throw/Archery\) will consume projectile. | Chance % | none | none |
| 348 | SE\_LimitManaMin | Limit to spells above a minium amount of mana. | Amt | none | none |
| 349 | SE\_ShieldEquipHateMod | Hate modifier to if shield equiped. \(Shield Specialist AA\) This may not be correct | Amt % | none | none |
| 350 | SE\_ManaBurn | Drains mana for damage/heal at a defined ratio up to a defined maximum amount of mana. Ratio \(-\) | Max Mana | Mana/HP Ratio / 10 | none |
| 351 | SE\_PersistentEffect | Aura effects, grants a buff automatically to those within a radius of caster. | ? | none | none |
| 352 | SE\_IncreaseTrapCount | UKNOWN EFFECT |  |  |  |
| 353 | SE\_AdditionalAura | Increase number of aura slots. | Amt Slots | none | none |
| 354 | SE\_DeactivateAllTraps | UKNOWN EFFECT |  |  |  |
| 355 | SE\_LearnTrap | UKNOWN EFFECT |  |  |  |
| 356 | SE\_ChangeTriggerType | UKNOWN EFFECT \[NOT\] |  |  |  |
| 357 | SE\_FcMute | Chance to silence casting of spells that contain specific spell effects. \(Effects determined by focus limits\) | Chance % | none | none |
| 358 | SE\_CurrentManaOnce | +/- Instant mana | Amt | none | none |
| 359 | SE\_PassiveSenseTrap | UKNOWN EFFECT |  |  |  |
| 360 | SE\_ProcOnKillShot | Triggers a spell after a kill shot. \(Max = Min Level of Target\) | Chance % | Spell ID | Min Level |
| 361 | SE\_SpellOnDeath | Triggers when the owner of the buff is killed. | Chance % | Spell ID | none |
| 362 | SE\_PotionBeltSlots | 'Quick Draw' expands the potion belt by one additional available item slot per rank. |  |  |  |
| 363 | SE\_BandolierSlots | 'Battle Ready' expands the bandolier by one additional save slot per rank. |  |  |  |
| 364 | SE\_TripleAttackChance | +/- chance to triple attack. | Chance % | none | none |
| 365 | SE\_ProcOnSpellKillShot | Chance to trigger a spell on kill when the kill is caused by the specific spell with this effect in it. | Chance % | Spell ID | Min Level |
| 366 | SE\_ShieldEquipDmgMod | Damage modifier to melee if shield equiped. \(Shield Specialist AA\) This may not be correct | Amt % | ? | none |
| 367 | SE\_SetBodyType | Change body type of target. |  Body\_Types  | none | none |
| 368 | SE\_FactionMod | Increases faction with base1 \(faction id, live won't match up w/ ours\) by base2. | Faction ID | Faction Mod | none |
| 369 | SE\_CorruptionCounter | Determines potency of determental curse spells \(+\) or potency of cures \(-\) | Counter Amt | none | Max |
| 370 | SE\_ResistCorruption | +/- Corruption Resist | Amt | none | none |
| 371 | SE\_AttackSpeed4 | Stackable slow effect |  Attack Speed  | none | Amt Max |
| 372 | SE\_ForageSkill |  |  |  |  |
| 373 | SE\_CastOnFadeEffectAlways | Triggers a spell when buff fades after natural duration OR from rune/numhits fades. | Spell ID | none | none |
| 374 | SE\_ApplyEffect | Chance to add an additional spell to the target. | Chance % | Spell ID | none |
| 375 | SE\_DotCritDmgIncrease | Modifies by % critical dot damage.  Spell Damage Modifiers  | Amt % | none | none |
| 376 | SE\_Fling | UNKNOWN EFFECT | ? | ? | none |
| 377 | SE\_CastOnFadeEffectNPC | Triggers a spell when buff fades after natural duration. \(On live these are in player spells cast on NPCs\) | Spell ID | none | none |
| 378 | SE\_SpellEffectResistChance | Chance to resist a specific spell effect type \(Ie. 5% chance to Resist DD spells\) | % Chance | Spell Effect ID | none |
| 379 | SE\_ShadowStepDirectional | Push forward for specific amount of units. \(Handled client side\) | Distance Units | none | none |
| 380 | SE\_Knockdown | Push back and up effect. \(Handled client side\) | Push Back | Push Up | none |
| 381 | SE\_KnockTowardCaster | Summon a player to caster at a set distance away from caster. \(Uses a reverse knockback effect\) | Distance Units | none | none |
| 382 | SE\_NegateSpellEffect | Completely remove a specified spell effect bonus. \(Ie. Remove all Aggro bonus\) | ? | Spell Effect ID | none |
| 383 | SE\_SympatheticProc | Chance to proc a a spell off of a regularly casted spell. \(Typically found as item focus\) | Chance Mod | Spell ID | none  |
| 384 | SE\_Leap | Jump toward your target. | Distance | none | none |
| 385 | SE\_LimitSpellGroup | Limit focus effects by spell group. \(spellgroup field in spells\_new\) \(+\) Include \(-\) Exclude | Spellgroup ID | none | none |
| 386 | SE\_CastOnCurer | Trigger a spell on yourself when you cure a target. | Spell ID | none | none |
| 387 | SE\_CastOnCure | Trigger a spell on yourself when you are cured. | Spell ID | none | none |
| 388 | SE\_SummonCorpseZone | Summon a corpse from any zone. | 1 | none | none |
| 389 | SE\_FcTimerRefresh | Reset recast timers \(ungrey spell gems\) | 1 | none | none |
| 390 | SE\_FcTimerLockout | Increase time on reset timer. | Amt | none | none |
| 391 | SE\_MeleeVulnerability | +/- melee mitigation \(+\) weakness \(-\) mitigation \(Live SPA defined SE\_LimitManaMax is not correct\) | Amt % | none | none |
| 392 | SE\_FcHealAmt | +/- flat amount to casters heal spells  Heal Modifiers  | Amt | none | none |
| 393 | SE\_FcHealPctIncoming | \(Buff/Debuff\) Modfies by % the casters base heal value for incomming spells.  Heal Modifiers  | Amt % | none | none |
| 394 | SE\_FcHealAmtIncoming | \(Buff/Debuff\) +/- flat amount the casters base heal value for incomming spells. Heal Modifiers  | Amt | none | none |
| 395 | SE\_FcHealPctCritIncoming | \(Buff/Debuff\) Modifies by % chance to receive a critical heal on incomming spells.  Heal Modifiers  | Amt % | none | none |
| 396 | SE\_FcHealAmtCrit | +/- flat amount to casters heal spells \(amt can critical\)  Heal Modifiers  | Amt | none | none |
| 397 | SE\_PetMeleeMitigation | +/- AC to owner's pet. | Amt | none | none |
| 398 | SE\_SwarmPetDuration | Increases duration of swarm pets by seconds. \(Base value set in miliseconds\) | Time | none | none |
| 399 | SE\_FcTwincast | Chance to cast the same spell 2x from a single cast. | Chance % | none | none |
| 400 | SE\_HealGroupFromMana | Group heal, drains mana from caster and coverts that to the amount of HP healed at a defined ratio. | Max Mana Drain | Mana/HP Ratio / 10 | none |
| 401 | SE\_ManaDrainWithDmg | Drains targets mana and decreases hit points based on a defined ratio of hp/mana drained. | Max Mana Drained | HP/Mana Ratio / 10 | Max ? |
| 402 | SE\_EndDrainWithDmg | Drains targets endurance and decreases hit points based on a defined ratio of hp/endur drained. | Max Endur Drained | HP/Endur Ratio / 10 | none |
| 403 | SE\_LimitSpellClass | Limits to specific types of spell categories. \(3=Cures,3=Offensive, 6=Lifetap\) \(+\) Include \(-\) Exclude | Category | none | none |
| 404 | SE\_LimitSpellSubclass | Limits to specific types of spell categories. \(UNDEFINED\) \(+\) Include \(-\) Exclude | Category | none | none |
| 405 | SE\_TwoHandBluntBlock | Modifies chance to block if 2 Hand Blunt equiped \(+\) Increase \(-\) Decrease | Chance % | none | none |
| 406 | SE\_CastonNumHitFade | Triggers a spell when a spells numhit counter is depleted. | Spell ID | none | none |
| 407 | SE\_CastonFocusEffect | Triggers a spell if focus limits are met \(ie Triggers when a focus effects is applied\) | Spell ID | none | noen |
| 408 | SE\_LimitHPPercent | Caps maximum HP to % or a defined amount, which ever is lowest. | Cap % | Cap Amt | none |
| 409 | SE\_LimitManaPercent | Caps maximum Mana to % or a defined amount, which ever is lowest. | Cap % | Cap Amt | none |
| 410 | SE\_LimitEndPercent | Caps maximum Endurance to % or a defined amount, which ever is lowest. | Cap % | Cap Amt | none |
| 411 | SE\_LimitClass | Limits to spells of a certain player class \(Uses Bitmask, the class value in spell dbase is 1 bitmask higher in relation to item class value\) | Class Bitmask | none | none |
| 412 | SE\_LimitRace | Limits to spells cast by a certain race \[NOT\] | Race ID | none | none |
| 413 | SE\_FcBaseEffects | Modifies base values of certain spell effects.  Mechanics  Partially implemented | Amt % | none | none |
| 414 | SE\_LimitCastingSkill | Limit to spells that use a specific casting skill. \(skill field in spells new\) \(+\) Include \(-\) Exclude |  Skills  | none | none |
| 415 | SE\_FFItemClass | UKNOWN EFFECT \[NOT\] |  |  |  |
| 416 | SE\_ACv2 | +/- AC \(stacks with other AC buffs\) | Amt | none | Max |
| 417 | SE\_ManaRegen\_v2 | +/- mana regen \(stacks with other mana regen buffs\) | Amt | none | Max |
| 418 | SE\_SkillDamageAmount2 | Add flat amount of damage when a specific melee skill is used. | Amt |  Skills \(-1=ALL\) | none |
| 419 | SE\_AddMeleeProc | Add melee proc | Proc Spell ID | Rate Modifier | none |
| 420 | SE\_FcLimitUse | Increase numhits count by percent. \[Custom\] \[NOT\] | Amt % | none | none |
| 421 | SE\_FcIncreaseNumHits | Increase numhits count by flat amount. | Amt | none | none |
| 422 | SE\_LimitUseMin | Limit to spells with a minimum number of numhit counters.  \[Custom\] \[NOT\] | Amt | none | none |
| 423 | SE\_LimitUseType | Limit to spells by  Numhit\_Types  \[Custom\] \[NOT\] |  Numhit\_Types  | none | none |
| 424 | SE\_GravityEffect | Pulls target\(s\) toward caster at a set pace to a specific distance away from caster. | Distance From Caster | Force of Pull | none |
| 425 | SE\_Display | Gives illusion flying dragon \(unclear how this works\) | 1 | none | none |
| 426 | SE\_IncreaseExtTargetWindow | Increases the capacity of your extended target window |  |  |  |
| 427 | SE\_SkillProc | Chance to proc a spell when using a specific skill \(ie Proc from a Taunt or Kick\). | Spell ID | Rate Modifier | none |
| 428 | SE\_LimitToSkill | Limits what skill will effect a SkillProc. \(Always place as next effect after ID 427/429\) |  Skills   | none | none |
| 429 | SE\_SkillProcSuccess | Chance to proc a spell when using a specific skill if it successfully hits the target. | Spell ID | Rate Modifier | none |
| 430 | SE\_PostEffect | Alter vision ? UNKNOWN EFFECT | ? | ? | ? |
| 431 | SE\_PostEffectData | Tint vision ? UNKNOWN EFFECT | RGB | ? | ? |
| 432 | SE\_ExpandMaxActiveTrophyBen | UKNOWN EFFECTâ€‹ \[NOT\] |  |  |  |
| 433 | SE\_CriticalDotDecay | Chance to critical DoT with effect depreciation.  Spell Damage Modifiers \(Removed from live 7-14 ?\) | Chance % | Decay % | Max Level |
| 434 | SE\_CriticalHealDecay | Chance to critical Heal with effect depreciation.  Heal Modifiers \(Removed from live 7-14 ?\) | Chance % | Decay % | Max Level |
| 435 | SE\_CriticalRegenDecay | Chance to critical Regen with effect depreciation.  Heal Modifiers \(Removed from live 7-14 ?\) | Chance % | Decay % | Max Level |
| 436 | SE\_BeneficialCountDownHold | UKNOWN EFFECTâ€‹ |  |  |  |
| 437 | SE\_TeleporttoAnchor | Teleport Guild Hall Anchor | ? | none | none |
| 438 | SE\_TranslocatetoAnchor | Translocate Primary Anchor | ? | none | none |
| 439 | SE\_Assassination | Damage done by 'Assassination' ability when proced  \(Humaniod target hit w/ Backstab\) | ? \(0\) | Amt Dmg | none |
| 440 | SE\_FinishingBlowLvl | Maximum level of target that 'Finishing Blow' will proc on. | Level | ? \(Seen 200\) | none |
| 441 | SE\_DistanceRemoval | Fades buff if owner of buff moves specified amount of distance from location where buff was applied. | Distance | none | none |
| 442 | SE\_TriggerOnReqTarget | Triggers a spell on target when a specific condition is met on that target. Buff fades after trigger. | Spell ID |  Target Condition  | none |
| 443 | SE\_TriggerOnReqCaster | Triggers a spell on target when a specific condition is met on that target. Buff fades after trigger. \(All spells that use this are self only\) | Spell ID |  Target Condition  | none |
| 444 | SE\_ImprovedTaunt | Locks aggro on caster and decreases other players aggro by % up to a specified level. | Max level | Aggro Mod | none |
| 445 | SE\_AddMercSlot | \[AA\] Allows you to conscript additional mercs. | Amt ? | none | none |
| 446 | SE\_AStacker | Buff stacking blocker  Buff Stacking  | Stacking Priority | none | none |
| 447 | SE\_BStacker | Buff stacking blocker  Buff Stacking  | Stacking Priority | none | none |
| 448 | SE\_CStacker | Buff stacking blocker  Buff Stacking  | Stacking Priority | none | none |
| 449 | SE\_DStacker | Buff stacking blocker  Buff Stacking  | Stacking Priority | none | none |
| 450 | SE\_MitigateDotDamage | Reduces incomming dotl damage by % up to a max value. | Mitigation % | Max Amt Reduced | Rune Amt |
| 451 | SE\_MeleeThresholdGuard | Partial Melee Rune that only is lowered if melee hits are over a defined amount \(limit\) of damage | Mitigation % | Min Hit | Rune Amt |
| 452 | SE\_SpellThresholdGuard | Partial Spell Rune that only is lowered if spell dmg is over a defined amount \(limit\) of damage | Mitigation % | Min Hit | Rune Amt |
| 453 | SE\_TriggerMeleeThreshold | Trigger spell when specified amount of melee damage is taken in a single hit, then fade buff. | Spell ID | Damage Amt | none |
| 454 | SE\_TriggerSpellThreshold | Trigger spell when specified amount of spell damage is taken in a single hit, then fade buff. | Spell ID | Damage Amt | none |
| 455 | SE\_AddHatePct | Modifies amount of hate you have on target by percent over instantly, | Amt % | none | none |
| 456 | SE\_AddHateOverTimePct | Modifies amount of hate you have on target by percent over time, | Amt % | none | none |
| 457 | SE\_ResourceTap | Coverts a percent of spell dmg from dmg spells \(DD/DOT\) to hp/mana/end. | % Coversion | 0=H/1=M/2=E | Max Amt |
| 458 | SE\_FactionModPct | Modifies faction gains and losses by percent. | Amt % | none | none |
| 459 | SE\_DamageModifier2 | Modifies damage amount by percent for a specific skill. | Amt % | Skills \(-1=ALL\) |  |

