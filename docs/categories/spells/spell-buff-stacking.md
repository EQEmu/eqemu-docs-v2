# Spell Buff Stacking



* **SE_StackingCommand_Block (148)**
  * Effect is found on buff and is used to prevent another buff from taking hold if specific criteria is met.
    * Base = Spell Effect
    * Formula - 201 = Slot
    * Max = Block if LESS THAN this value.
      * **Exampe:** AC buff with a value of 500. You want to block any other AC spell from going into Slot 1 that is less than 1000 AC.
        * Slot 1 : Effect SE_AC(1) with Base=500 
        * Slot 2:  Effect SE_StackingCommand_Block (148) with Base= SE_AC(1) Max = 1000 and Formula = 201 (apply to the first spell slot--remember that we use base 0)
* **SE_StackingCommand_Overwrite (149)**
  * Effect is found on buff and is used to overwrite another buff if specific criteria is met.
    * Base = Spell Effect
    * Formula - 201 = Slot
    * Max = Overwrite if LESS THAN this value.
      * **Exampe: **AC buff with a value of 1200. You want to overwrite any other AC spell in Slot 1 that is less than 1000 AC.
        * Slot 1 : Effect SE_AC(1) with Base=1200 
        * Slot 2:  Effect SE_StackingCommand_Overwrite (149) with Base= SE_AC(1) Max = 1000 and Formula = 201 (apply to the first spell slot--remember that we use base 0)
* **SE_Screech (123)**
  * Any effect with SE_Screech that has a value of +1 will block any other effect with SE_Screech in it that has a value of -1.
    * Example:
      * Spell [Screech](http://lucy.allakhazam.com/spell.html?id=1383\&source=Live) has SE_Screech with value of -1. 
      * Spell [Screech Immunity](http://lucy.allakhazam.com/spell.html?id=2785\&source=Live) has SE_Screech with value of +1
        * If you have the buff Screech Immunity, the spell Screech will not be able to effect you or any other spell with Screech at (-1)
* **SE_AStacker (446), SE_BStacker (447), SE_CStaker (448), SE_DStaker (447) **
  * Buffs containing these effects can block each other from taking hold via the following. _(Does not matter what slot the effect is in)_
  * (B) Blocks any buffs from taking hold with (A) in it.
  * (C) Blocks any buff from taking hold with (B) in it.
  * (D) Blocks any buff from taking hold with (C) in it.
  * When checking same type (ie A vs A), the higher base effect value will determine which takes hold.
