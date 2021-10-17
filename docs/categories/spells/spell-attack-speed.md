# Spell Attack Speed

* **SE_AttackSpeed**
  * **Standard Haste and Slows**
    * Baseline attack speed is set at a value of 100.
    * This spell effect modifies that value.
      * **Haste**
        * Haste spells will have a value greater than 100.
        * Haste amount is calculated at (Value - 100). 
        * Example: [Swift like the Wind](http://lucy.allakhazam.com/spell.html?id=172\&source=Live) (Value = 160) 
          * 160 - 100 = 60% Haste
        * Haste Caps
          * Lv 60+ = **(100)**
          * Lv 51 - 59 = **(85)**
          * Lv < 50 = **(level + 25)**
      * **Slow**
        * Slow spells will have a value less than 100
        * Slow amount is calculated at (100 - Value)
        * Example: [Sha's Legacy](http://lucy.allakhazam.com/spell.html?id=6828\&source=Live) (Value = 35)
          * 100 - 35 = 65% Slow
        * Stacking: Slow will always overwrite haste in the same spell slot.
      * **Slow Mitigation**
        * Set in npc_types table 'slow_mitigation'
        * Decreases the potency of a slow effect by a percent as defined in the table

| **Range** | **Client Message when cast on an NPC with slow mitigation** |
| --------- | ----------------------------------------------------------- |
| 0 - 25 %  | "Your spell was mostly successful"                          |
| 26 - 74 % |  "Your spell was partially successful"                      |
| 74 - 100% |  "Your spell was slightly successful"                       |

* **SE_AttackSpeed2 (98)**
  * **Stackable Haste (w/ SE_AttackSpeed) that does not go over the haste cap**
    * Calculated same as SE_AttackSpeed
    * Value stacks directly with with SE_AttackSpeed.
    * Example: [Swift like the Wind](http://lucy.allakhazam.com/spell.html?id=172\&source=Live) (60%) + [Composition of Ervaj](http://lucy.allakhazam.com/spell.html?id=1452\&source=Live) **(10%)** = 70% Haste
    * Can also be used a stackable slow effect as seen in [Al'Kabors Spiral of Deadly Force](http://lucy.allakhazam.com/spell.html?id=17974\&source=Live)
* **SE_AttackSpeed3 (119)**
  * **Stackable Haste (w/ SE_AttackSpeed + SE_AttackSpeed2) that can go over the haste cap**
    * The direct spell effect value is added to the haste derived from SE_AttackSpeed and SE_AttackSpeed2 and will exceed the haste cap.
    * Example (**IF CAP AT 100**): [Swift like the Wind](http://lucy.allakhazam.com/spell.html?id=172\&source=Live) (60%) + [Composition of Ervaj](http://lucy.allakhazam.com/spell.html?id=1452\&source=Live) (10%) + Item Haste (30%) = 100% + [Ancient Power](http://lucy.allakhazam.com/spell.html?id=6375\&source=Live) **(10%) =** 110% Haste
    * Can also be used as a stackable slow effect as seen in [Taking Root](http://lucy.allakhazam.com/spell.html?id=31853\&source=Live)
    * Overhaste Caps
      * Lv 51+ = **(25) **
      * Lv 1 - 50 = **(10)**
* **SE_AttackSpeed4**
  * **Stackable slow effect (w/ SE_AttackSpeed)**
    * **When slowed by SE_AttackSpeed:** This effect decreases the remaining portion of your attack speed value NOT lowered by standard slow
    * Example: (SE_AttackSpeed) [Sha's Legacy](http://lucy.allakhazam.com/spell.html?id=6828\&source=Live) 65% slow + (SE_AttackSpeed4) [Lassitude](http://lucy.allakhazam.com/spell.html?id=11785\&source=Live) 25% slow
      * Sha's Legacy is calculated as 100 - 35 = 65% slow, therefore the remaining attack speed is (35)
      * Lassitude will now decrease the remaining value (35) by 25% = 16.5%
      * The total slowed value on the target would be 65% + 16.5% = 81.25% slow
    * **When NOT slowed by SE_AttackSpeed:** This effect acts as a regular slow, decreasing attack speed by the full amount.
    * Note: The base value should always be set as a positive. There are a few examples of spells with negative values, however all of their descriptions indicate it as a slow effect (It is likely developer error)
