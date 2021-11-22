---
description: >-
  This section covers information regarding the Alternative Advancement (AA)
  system of the EQEmu Server.
---

# AAs

The Alternative Advancement system was introduced during the Luclin expansion of EverQuest.  AAs were gained by contributing some portion of gained experience to the AA experience pool, diverting experience from the standard leveling experience pool.

### Client Considerations

The behavior of most AAs is determined by [spell effects](../spells/spell-effect-ids.md).  It is important that server operators who have customized AAs distribute an updated spells file to ensure proper behavior.  

!!! info
    If you would like additional information regarding Spells, please refer to the [Spells Guide](../spells/).  Having a thorough understanding of the spells system prior to editing AAs is recommended.

Additionally, the appearance of AAs or customized AAs in-game can be modified through the server operator's choice of [AA Categories](aa-categories.md).  These categories determine the tab on which the AA will appear in game.

### Server Considerations

A number of database tables are used to create the AA system.  You can, of course, [customize AAs](customizing-aas.md) for your world.  

!!! info
    Customizing AAs is typically achieved through direct edits / queries to your database, as their behaviors can vary from using spell effects, to rank effects, and so on. 

!!! warning
    Directly editing your database can have unintended consequences. Be sure to create a backup.

Understanding the interaction of the database tables will surely help, should you choose to customize:

| Table | Description |
| :--- | :--- |
| [aa_ability](https://eqemu.gitbook.io/database-schema/server/aas/aa_ability) | the name of each AA, class/race/deity restrictions, enable/disable, [AA type](aa-types.md) setting |
| [aa_actions](https://eqemu.gitbook.io/database-schema/server/aas/aa_actions) | assigned spellid, timers, [AA target types](aa-target-types.md) setting |
| [aa_effects](https://eqemu.gitbook.io/database-schema/server/aas/aa_effects) | links the spell effects and spell base values to each AA |
| [aa_rank_effects](https://eqemu.gitbook.io/database-schema/server/aas/aa_rank_effects) | modifies spell effects and base values for each AA as ranks increase |
| [aa_rank_prereqs](https://eqemu.gitbook.io/database-schema/server/aas/aa_rank_prereqs) | determines any prerequisites required before an AA purchase is allowed |
| [aa_ranks](https://eqemu.gitbook.io/database-schema/server/aas/aa_ranks) | settings for each AA cost and progression through its ranks |
| [aa_required_level_cost](https://eqemu.gitbook.io/database-schema/server/aas/aa_required_level_cost) | Handles certain AAs such as Harm Touch and Lay On Hands which can progress with or without AA experience cost, based solely on level |
| [aa_timers](https://eqemu.gitbook.io/database-schema/server/aas/aa_timers) | a character data table that indicates when a time-limited AA will again be available to the player |



