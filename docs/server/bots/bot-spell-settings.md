# Bot Spell Settings

## Description 

What a Bot casts, can be further defined by the player by using a set of various Bot Commands.

## Using **^spells**

**^spells min_level (Optional Value)** - Will provide a list of all available Bot Spells that your Targeted Bot can cast. Per the example below, you can see a List of Available Spells, along with an option to "Add" the Spell to your Spell Settings.

![](https://user-images.githubusercontent.com/109764533/205463827-808a9a8f-c07f-4699-8944-f6cf02b2bbc8.png)

## Using **^spellsettings**

**^spellsettings** - Will provide a list of all Bot Spells that you have added to your Bots Spell Settings List, using this command you can Toggle the Enabled/Disabled State of Spells, or remove a Spell from the Bots Spell Settings List.

![](https://user-images.githubusercontent.com/109764533/205463846-310f7604-94cc-469a-9ffe-599a794a5ca9.png)

## Using **^enforcespellsettings**

**^enforcespellsettings [true or false]** - By default Enforcing the Spell Settings List is false (disabled). Setting ^enforcespellsettings to true will force your Bot to **only** cast spells that you have added to it's Spell Settings List (viewed/managed with **^spellsettings**) you can add new spells to your Spell Settings List with **^spells** and clickking on "Add"

## Using **^spellsettingsupdate**

**^spellsettingsupdate [Spell ID] [Priority] [Min HP] [Max HP]** - Optionally with this command you can assign a different Priority to an individual Bot Spell (Lower Priority is more likely to be cast), additionally you set a Min/Max HP for when your Bot will cast a spell **(Currently this only works for conversion spells like Canni)** but will eventually be expanded to work for all spells.



