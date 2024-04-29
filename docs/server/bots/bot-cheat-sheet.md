!!! info end

    This page lists a quick reference and highlights how to use bots. Check out [bot commands](/server/bots/bot-commands) for a more detailed list of commands that bots can use.

## Getting Started

There's a one time command to create a bot. Use `^create help`. You'll get a popup listing the options, e.g. `^create Xackery 2 6 1` creates a bot named Xackery, class cleric, race dark elf, gender female

Now, you can spawn it with `^spawn Xackery`. To make it despawn, you use `^camp` while targetting the bot, or `^camp Xackery` to have it camp without targetting it.

Once a bot is spawned, you can invite it to your group via `/inv Xackery` as if it's a normal player.

Use `^botstopmeleelevel #` to set a level to tell the bot to stop meleeing. E.g. set it to 0 so they never engage in melee.

I suggest hotkeying `^summon all` to quickly get bots unstuck and at your location. Since bots aggro, it's suggested to just spam this so you have better control. It's also a way to leash a bot who engages too aggressively.

I suggest hotkeying `^attack` as well, this is handy for having the bots engage first

Using `^invgive` while holding an item on your cursor is faster than trading with each bot individually.

Other handy commands to consider are: `^cure disease`, `^pull`, `^guard`, `^invis living`, `^waterbreathing`, `^levitation`, `^track`, `^taunt`, `^resistance`, `^resurrect`, `^rune`, `^portal`, `^picklock`, `^movementspeed`, `^lull`, `^mesmerize`, `^charm` 

### Spells

Spells have two types of logic. By default, all spells go to the `^spells` command. Default behavior is just to use every spell on this list. You can press the add saylinks by each spell to add them to spell settings explained below.

Spell setting is a mode to let you granularly control spells, first, you need to opt in to spell settings by typing `^enforcespellsettings`

Now that you've opted in, you can use `^spellsettings` to see the list of spells you can enable/disable, or remove. By default, spells you pressed Add on from ^spells goes to this list enabled, so you don't need to do anything, But, pressing the `Enabled` links will flip them to disabled, and `Remove` will remove them from the list so they're back in the `^spells` section.

There are many more advanced ways to use bots, including healrotations, stances, disciplines, but for a quick start this should get you started.
