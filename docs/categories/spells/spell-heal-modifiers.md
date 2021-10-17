# Spell Heal Modifiers



**Overview of spell effects that modify Healing**

**Key:** **Focus Effect,** _**Item Effect**_**,** ~~**Non-Focus Spell Effect**~~

**INSTANT HEALS**

_\*Heal critical modfier is always set at 100% (2x)_

_\*If multiple of same Focus Effect, highest avialable value is used. \[Stacks] Indicates that if multiple of same effect, the final amount is the total of all effects. _

* **Focus Effects** **that increase/decrease healing by percent.**
  * **SE_FcBaseEffects (413)** Modifies by % the base heal value of the spell cast.
  * **SE_ImprovedHeal** **(125) **Modifies by % the heal value after (413) is applied.
* **Effects on CASTER that increase/decrease healing by a specific amount.**
  * **SE_FcHealAmt (392)** +/- a set heal amount. (This value CAN NOT criticals).
  * **SE_FcHealAmtCrital (396) **+/- a set heal amount. (This value CAN critical)
  * _**Item 'Heal Amount"**_ - Amount specified on the item modified by formula bellow. (This value CAN critical)
    * Value is modified by the 'total cast time' of the spell it is being applied to
    * _Total Cast Time_ is derived from adding which ever is greater(Recast OR Recovery time) to cast time.
    * _Total Cast Time_ (0 - 2.5 seconds). 
      * Heal Amount = Heal Amount / 4
    * _Total Cast Time_ (2.6 - 7.0 seconds)
      * Heal Amount = Heal Amount\*(0.167\*(_Total Cast Time_)/1000))
    * _Total Cast Time_ ( > 7.0 seconds)
      * Heal Amount = Heal Amount / 7
* **Effects on TARGET that increase/decrease healing by a specific amount.**
  * **SE_FcHealAmtIncoming (394) **+/- a set amount of healing. (This value CAN NOT critical)
* **Effects on TARGET that increase/decrease healing by percent.**
  * ~~**SE_HealRate (120)**~~ Modfies by % the casters base heal value. (This value CAN NOT critical) \[Stacks]
  * **SE_FcHealPctIncoming (393)**  Modifies by % the casters base heal value. (This value CAN NOT critical).
* **Effects on CASTER that increase critical spell chance.**
  * ~~**SE_CriticalHealChance (274)**~~ Modfies by % the chance to critical heal.\[Stacks]
  * ~~**SE_CriticalHealDecay (434)**~~ Modifies by % the chance to critical heal. \[Stacks]
    * Effect value decreases if spell cast is over the maximum level specified.
* **Effects on TARGET that increase critical spell chance.**
  * **SE_FcHealPctCritIncoming (395)** Modifies by % the casters chance to critical heal.

**Calculation **

**Critical Chance = **~~**SE_CriticalHealChance**~~** + **~~**SE_CriticalHealDecay**~~** + SE_FcHealPctCritIncoming**

**X =** Heal Amount

**CritMod** = 2 (Assume we critical, no critical set to 1)

**Final Heal Amount = (((X \* SE_FcBaseEffects) \* SE_ImprovedHeal )\*CritMod) +  SE_FcHealAmt + (SE_FcHealAmtCrital \* CritMod) + (Item 'Heal Amount" \* CritMod) + SE_FcHealAmtIncoming **

**+  ((X \* SE_FcBaseEffects) \* **~~**SE_HealRate**~~** )) + ((X \* SE_FcBaseEffects) \* SE_FcHealPctIncoming )) **

**HEAL OVER TIME**

_\*Heal critical modfier is always set at 100% (2x)_

* **Heal Over Time spells are NOT modifed by the following**
  * **Effects on TARGET that increase/decrease healing by a specific amount.**
  * **Effects on CASTER or TARGET that increase/decrease healing by a specific amount.**
* **Effects on CASTER that increase critical spell chance. **
  * ~~**SE_CriticalHealOverTime (319)**~~ Modfies by % the chance to critical heal. \[Stacks]
  * ~~**SE_CriticalRegenlDecay (435)**~~** **Modifies by % the chance to critical heal. \[Stacks]
    * Effect value decreases if spell cast is over the maximum level specified.
* **Effects on TARGET that increase critical spell chance.**
  * **SE_FcHealPctCritIncoming (395) **Modifies by % the casters chance to critical heal.
