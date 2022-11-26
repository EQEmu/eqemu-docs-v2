# Bot Casting Logic

## Description 

Bot Logic is detailed below, Bots will prefer to cast spells in a specific order depending on if they have aggro, or need a heal.

Below is the current logic for each Bot Class.

 Class | Casting Logic |
| --- | --- |
|Bard| Escape -> HateRedux -> Slow -> Debuff -> Combat Buff Song -> DoT -> Nuke |
|Beastlord |  Slow -> Heal Self -> Debuff -> Pet -> Heal Pet -> Combat Buff -> DoT -> Nuke -> Heal Others |
|Berserker|Escape -> Heal Self -> Debuff -> Combat Buff -> Combat Buff Song -> DoT -> Nuke |
|Cleric| Escape -> Heal Self -> Heal Others -> De-buff -> Combat Buff -> Nuke|
|Druid| Escape -> Heal Self -> Heal Others -> Debuff -> Combat Buff -> DoT -> Nuke |
|Enchanter| Mez -> Escape -> Slow -> Debuff -> Combat Buff -> DoT -> Nuke |
|Magician | Pet -> Debuff -> Combat Buff -> Heal Pet -> Nuke |
|Monk | Escape -> Heal Self -> Debuff -> Combat Buff -> Combat Buff Song -> DoT -> Nuke |
|Necromancer | Escape -> Pet -> Debuff -> Lifetap -> Combat Buff -> DoT -> Heal Pet -> Nuke |
|Paladin | Heal Self -> Combat Buff -> Debuff -> Heal Others -> Nuke  |
|Ranger | Escape -> Heal Self -> Combat Buff -> DoT -> Nuke -> Heal Others |
|Rogue | Escape -> Heal Self -> Debuff -> Combat Buff -> Combat Buff Song -> DoT -> Nuke|
|Shadow Knight | Lifetap -> Combat Buff -> Debuff -> DoT -> Nuke |
|Shaman | Escape -> Slow -> Heal Self -> Heal Others -> Debuff -> Combat Buff -> DoT -> Nuke -> Heal Pet |
|Warrior| Heal Self -> Combat Buff -> Debuff -> Combat Buff Song -> DoT -> Nuke |
|Wizard | Escape -> Debuff -> Combat Buff -> Nuke |