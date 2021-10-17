# Spell Mechanics



**Collection of details on various spell effect mechanics.**

\*Most of the mechanics described in this section were deduced from extensive personal parsing on live server \~ Kayen

\*Term 'resist check' means essessentially the chance for the spell to be resisted. Resist checks can occur for various reasons after the spell has already landed. Failed 'resist check' means as if it was resisted.

**ROOT**

* **SE_Root (99)**
  * Root has an innate chance to break every buff tick.
    * There is a 70% chance to trigger a resist check each tick.
      * If resist check fails root will fade.
      * There is a lower bounds 6% chance per tick of root breaking no matter what.  (If you have 0% chance to resist the spell due to Tash/Green con)               
  * Root has an innate chance to break every time a direct damage spell is applied
    * There is an approximately 70% chance of root breaking from a DD spell if target is greater than your level, equal to your level or (1) level less than you.
    * There is an approximately 55% chance of root breaking from a DD spell if the target is (2) levels less then you.
    * There is an approximately 42% chance of root breaking from a DD spell if target is between (3) and (20) levels less then you.
    * There is an approximately 14% chance of root breaking from a DD spell if target is (20) levels less then you.
    * **SE_RootBreakChance (244)** _'Visicid Roots AA'_** **Reduces the baseline root break chance (ie 70%) by the effect value.
      * When a root is cast by a player with this effect the reduced root break chance is applied directly to the root.
      * The benefit is therefore given to any player casting on that NPC with the root. (opposed to only the caster of the root).
    * When multiple roots are applied to the same target the root in the first slot is always removed first from spell damage.
    * Only roots from determinental spells can be broken by direct damage spells.

\*The following spell effects (Charm, Mez, Pacify) resist chances are influenced in part by this specific **charisma modifier.**

* Charisma Modifer applies a bonus/penalty to your resist chance base on your charisma statistic.
  * For Charisma >= 75 and <= 255, every 10 Charisma is equivelent to a -1 resist modifier. _(-1 = 0,5 % reduction in resist chance)_
  * For Charisma < 75, every 10 Charisma less than 75 is equivelent to a +6 resist modifer. _(6 = 3% increase in resist chance)_

**CHARM**

* **SE_Charm (22)**
  * Charm when is cast initially a resist check is made which is effected by a **charisma modifier**.
  * Charm if successful has an innate chance to break every buff tick. (_Unless flagged in unbreakable in spell table 'powerful_flag_')
    * There is an approximately 25% chance to trigger a resist check each tick to determine if charm will fade.
      * Resist check for charm to fade is NOT effected by Charisma or Base1 value (unknown what that does).
      * There is a lower bounds of approximately 3% chance per tick of charm breaking no matter what.  (If you have 0% chance to resist the spell due to Tash/Green con)  
      * If resist check fails...
        * If player has a bonus from **SE_CharmBreakChance (243)** _'Total Domination AA'  _you have another chance for charm not to break based on the value,
          * Example: _Total Domination Rank 3_ is 50% chance for charm not to break if resist check fails.
        * Otherwise, charm will fade.

**PACIFY**

* **SE_Lull (18) SE_Harmony (89) SE_ChangeFrenzyRad (30)**
  * Pacification abilities chance to be resisted is NOT influenced by the targets magic resistance (or any other resist stat).
    * There is a flat 7.5 % chance to resist that is modified only by the level difference between caster and target.
  * If Pacification is RESISTED then a second resist check is performed.
    * This resist check does use the targets regular resistances and a **charisma modifier.**
      * If the resist check fails (resisted) then target will aggro the caster.
      * If the resist check passes (resisted) then target will NOT aggro the caster.

**MESMERIZATION**

* **SE_Mez (31)**
  * Mez when is cast initially a resist check is made which is effected by a **charisma modifier**.
    * Mez base value of 1 or 2 determines stacking behavior.
      * Mez with value of 2 will always overwrite a mez with a value of 1.

**FEAR**

* **SE_Fear (23)**
  * Fear when is cast initially a resist check is made which is effected by a **charisma modifier** (_Note: This is not the same as Charm/Mez/Paci)_
    * Charisma Modifer applies a bonus/penalty to your resist chance base on your charisma statistic.
      * For Charisma < 100, fear is given a -20 resist modifier. _(-20 = 10 % reduction in resist chance)_
      * For Charisma > 100 to 255, the -20 resist modifier is progresively lowered to zero as you gain more charisma (0 at 255).
      * This modifier DOES NOT apply to Undead Fears.
  * Fear has an innate chance to break every buff tick.
  * There is a 70% chance to trigger a resist check each tick.
    * Resist check for fear to fade is NOT effected by Charisma or Base1 value (unknown what that does).
    * If resist check fails fear will fade.
    * There is NO lower bounds chance per tick of fear breaking.  (If you have 0% chance to resist the spell due to Tash/Green con)     

**DISPELL**

* **SE_CancelMagic (27), SE_Purify (291), SE_DispelDetrimental (124)**, **SE_DispelBeneficial (209)**
  * When dispel is cast every buff starting with the first buff slot is checked to see if it can be removed. _(Buffs can be set immune to dispel in table)_
    * If it fails to dispel a buff it will move on to the next buff and try that until there are no buffs left to check.
    * Dispels often have multiple chances to dispel as in [Annul Magic](http://lucy.allakhazam.com/spell.html?id=1526\&source=Live)
      * Meaning for each 'slot' there is a dispel, is its own independent chance to run through your buff list and remove a buff.
  * To determine if the buff will be removed dispels check the level of the caster agianst the buff level.
    * Buff Level is determined by the level of the caster who that buff.
    * If there is no difference between the casters level and the 'buff level' then there is an approximately 36% chance to have that buff removed.
    * If the level of the caster is less than the the 'buff level' then for each level difference the chance to remove the buff is decreased by 8%.
      * Example.. If buff level is 2 level higher then casters level you have a (36 - 16) = 16% chance to remove the buff.
      * There is always a lower limit of 10% chance to remove the buff regardless of level difference.
    * If the level of the caster is greater than the 'buff level' then for each level difference the chance to remove the buff is increased by 2%.
    * The base effect value of these effects is a level modifier.
      * Annul Magic has a value of (9) . This means when calculating chance to remove the buff we ADD 9 levels to the caster.
      * Example: Caster is Lv 15 and targets 'buff level' 24 and you cast annul magic (9).
        * This means that the level difference is (15+9) - 24 = 0
        * Resutling in a 36% chance to remove the buff.

**6/6/14** - \[**PATCH 5-20-14**] Note: _Updated all spells which use Remove Detrimental and Cancel Beneficial spell effects to use a new method. The chances for those spells to affect their targets have not changed unless otherwise noted_

* Base value now defines a percent chance for success? Where Base / 10 = Chance (Ie. 950 / 10 = 95%)
* _At present time this effect is supported using PRE or POST patch values._

**DEATH SAVE**

* **SE_DeathSave (150)**
  * When a DeathSave buff (_Divine Intervention / Death Pact)_ is applied to a target the owner of the buff is given a chance at receiving a heal when under 15% hit points.
    * This effect always triggers when under 15% HP.
      * The heal amount is determined by the base value \[1=Partial 3000 HP] \[2=Full 8000 HP].
        * If Max Value is set as heal amount this value will be used instead as the heal amount.
      * Chance to receive a heal is determined by the owner of the buffs Charisma.
        * Chance = ((Charisma \* 3) +1) / 10) 
        * If chance is successful then the owner of the buff recieves the heal and the buff FADES.
        * If chance is NOT succesful then the buff FADES.
          * OR If you fail and have an effect from **SE_UnfailingDivinity (277)** _**'**AA Unfailing Divinity'_ you are given a second chance.
            * The chance is the same as the initial check using Charisma.
            * If succesful then the owner of the buff recievies a heal that is modified by the value of **SE_UnfailingDivinity (277)**
              * If value of effect is 20 then your heal is only 20% of the initial value _(Ie. 8000 \* 20 / 100)_
            * If chance is NOT succesful then buff FADES

**DIVINE SAVE**

* **SE_DivineSave (232)**
  * Divine Save provides an opportunity to avoid death. The effect is triggered on the owner of the buff/aa/item with this effect when actual character death occurs (after delay death effects).
    * The base value determines the chance for this effect to occur upon death, if successful...
      * The owner of the effect has their HP restored to (1)
      * A spell is cast on the owner (Spell ID determined by base2/limit). This is typically the heal spell [Touch of the Divine I](http://lucy.allakhazam.com/spell.html?id=4544\&source=Live).
      * The short duration buff [Touch of the Divine](http://lucy.allakhazam.com/spell.html?id=4789\&source=Live) is ALWAYS cast, providing HOT/Removal of Determineal Buffs/Invulnerability.

**SE_FcBaseEffects (413)**

* Focus effect that modifies the base values_ (value before any other modifications are applied)_ of many spell effects _(not typically subject to focus effects). _
* List of known spell effects that are FcBaseEffects modifies
  * SE_SkillAttack - Modifies base weapon damage amount
  * Direct/DoT base spell damage amounts. [ Spell Damage Modifiers](http://wiki.eqemulator.org/p?Spells_DamageModifiers\&frm=Spells_Mechanics)
  * Heal/HoT base spell healing amounts. [ Heal Modifiers](http://wiki.eqemulator.org/p?Spells_HealModifiers\&frm=Spells_Mechanics)
  * SE_Rune - Modifies base rune amount
  * Bard instrument modifiers - NOT IMPLEMENTED  _(Was reports of crashes possibly due to this so remains disabled till can be resolved)_
  * Bard song effects (regular foci do not work on bard songs) - NOT IMPLEMENTED _(Was reports of crashes possibly due to this so remains disabled till can be resolved)_
