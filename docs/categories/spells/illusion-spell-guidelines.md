# Illusion Spell Guidelines

Illusion Spells follow a couple different rules.

These are influenced by spell base, limit, and max values.

### Gender-Inspecific Illusions

If an Illusion Spell has a max value of 0 on the Illusion effect ID we use the base value as the Race ID, assume the [gender](../npc/genders.md) using a source method, use the limit value as [texture](../npc/textures.md), and the max value as [helmtexture](../npc/textures.md).

Example Spell is Illusion: Human \(582\).

### Gender-Specific Illusions

If an Illusion Spell has a max value greater than 0 on the Illusion effect ID we check if the limit value is 0, if it is we use the base value as the Race ID and the max value minus 1 as the [gender](../npc/genders.md), [texture](../npc/textures.md) and [helmtexture](../npc/textures.md) are assumed.

* 1 is Male
* 2 is Female
* 3 is Neuter

Example Spell is Illusion: Male \(1732\).

### Gender and Texture Specific Spells

If an Illusion Spell has a max value greater than 0 on the Illusion effect ID and the limit value is not 0, we then check if the max value is not equal to 3, if it's not equal to 3 we use the base value as the Race ID, the max value minus 1 as the [gender](../npc/genders.md), the limit value as [texture](../npc/textures.md), and the max value as [helmtexture](../npc/textures.md).

Example Spell is Illusion: Runic Tattoo Nihil Male \(33575\).

### Gender, Texture, and Helmtexture Specific Spells

If an Illusion Spell has a max value greater than 0 and equal to 3, we use the base value as the Race ID, assume the [gender](../npc/genders.md) using a source method, use the limit value as [texture](../npc/textures.md) and [helmtexture](../npc/textures.md).

Example Spell is Illusion: Rallosian Goblin \(32787\).

