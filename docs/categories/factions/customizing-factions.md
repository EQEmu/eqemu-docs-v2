---
description: This page describes customizing factions on your EQEmu server
---

# Customizing Factions

Every faction that exists on a server is listed in the [`faction_list`](https://eqemu.gitbook.io/database-schema/categories/factions/faction_list) table. This table includes a unique id, a name, and a starting \(base\) faction. This starting value is combined with any race/class/deity adjustments that apply \(see [`faction_list_mod`](https://eqemu.gitbook.io/database-schema/categories/factions/faction_list_mod) table\) to create the starting faction for characters when they are born. As characters interact with the world, they can gain/lose faction through their actions, be that killing or questing. These values are continuously updated in the [`faction_values`](https://eqemu.gitbook.io/database-schema/categories/factions/faction_values) table.

{% hint style="info" %}
Note that the `faction_values` table contains _character_ data, even though it does not follow the Database Schema Guidelines \(it should be `character_faction_values`\).
{% endhint %}

{% hint style="success" %}
As of [06 March 2019](https://eqemu.gitbook.io/changelog/years/2019#3-1-2019) we have imported the client faction list and we now use these renumbered factions. The client does not have a base for any of the factions, all of the offsets are realized in faction\_list\_mod instead. This does not mean that the base cannot be altered for custom results.
{% endhint %}

Each faction has a minimum and a maximum value, at which point no more faction can be lost or gained. This information is stored in the `faction_base_data` table. The default _base_ maximum is 2000 and the _base_ minimum is -2000 if there is no entry. The amount a character can earn is based on these constants **and** the faction's initial base value. For example, a faction that starts out at 0 \(base value in `faction_list`\) is a simple case, and would use the maximum of 2000 and minimum of -2000.

The character can earn or lose faction until his _earned_ faction reaches the faction's min or max value. The base value for a faction impacts this in cases where that base is non zero. For example, if a faction starts out at -500, the character can earn 2500 points--500 to make up for initial impressions, and 2000 more to get into good graces. If a faction starts you out at a positive value--for instance, 300--you can only earn 1700 points, but you can lose 2300 points!

A character's unmodified faction value at any time is calculated by the following formula:

> Current Unmodified Faction = faction\_list 'base' integer + faction\_values 'current\_value' integer

{% hint style="info" %}
Again, note that the integer in the 'current\_value' field of the `faction_values` table represents the factions gains and hits the player has received from killing NPCs or completing quests.
{% endhint %}

This faction value is then adjusted by any impact caused by the character's race/class/deity combination. These modifications are stored in the `faction_list_mod` table. These adjustments cannot be made up for by any actions or deeds:  they are permanent marks on the character's record, and can therefore impact how high \(or low\) a character can get with the faction in question. They can be temporarily masked by things like illusions, but cannot be completely erased.

To get your final \(modified\) faction, add the race/class/deity modifier to the Current Unmodified Faction value:

> Current Faction = \[Current Unmodified Faction\] + faction\_list\_mod 'mod' integer

### Examples of Current Faction

Imagine that we create a character named "Snake", who is an Ogre Berserker that worships Cazic Thule.  

#### Unmodified Case:

The NPCs using faction ID 368 - Dervish Cutthroats enjoy a base value of -100 \(in the `faction_list` table\).  Since they have no faction list modifiers \(based on race/class/deity\), everyone, including Snake, would start off conning "Apprehensive":  

> faction\_list 'base' integer = -100  
> faction\_values 'current\_value' integer = 0  
> faction\_list\_mod 'mod' integer = N/A
>
> Result:  -100 Current Faction \("Apprehensive"\)

The maximum number of faction points possible \(stored in the `faction_values`table\), is +2100, and the minimum number of faction points possible, is -1900.   

Assuming that there is a faction hit of -10 for killing a Dervish Cutthroat, and our intrepid adventurer, Snake, killed a derv:  Snake's faction value would be -10, to which we would add the base value of -100, and Snake would consider "Dubious" to Dervish Cutthroats.

If Snake received +1200 faction points, he would consider "Ally".  If Snake received -651 points of faction hits, he would consider "Scowls, Ready to Attack".  From our example above, if Snake killed 190 dervs, he would reach the minimum faction value possible.  

#### Modified Case:

The NPCs using faction ID 361 - Ashen Order have a base value of 0, but they DO have a modifier based on both race and class \(but none on deity\):

![Class and Race modifiers for faction 361 - Ashen Order](../../.gitbook/assets/ashen_order-flm.png)

Above we see a rather cryptic list of modifiers to faction 361 - Ashen Order.  Note that in the `mod_name` column, we see the [Classes](https://eqemu.gitbook.io/server/categories/player/class-list) to be modified \(c11, c5, c7, c9\) as well as the [Races](https://eqemu.gitbook.io/server/categories/npc/race-list) to be modified \(r10, r128, r6, r9\), and their corresponding modifier values in the 'mod' column.

With no 'base' value \(in the `faction_list` table\), the maximum number of faction points possible \(stored in the `faction_values`table\), is +2000, and the minimum number of faction points possible, is -2000.   

Recall that Snake is an Ogre Berserker.  The Berserker class suffers no penalty in this case, but the Ogre race \(ID 10\) suffers a massive -1000 penalty.  Therefore, before Snake has done any killing or questing \(meaning, his faction value for Ashen Order would not exist\), NPCs with this faction assigned would consider him "Scowls, Ready to Attack".

> faction\_list 'base' integer = 0  
> faction\_values 'current\_value' integer = 0  
> faction\_list\_mod 'mod' integer = -1000
>
> Result:  -1000 Current Faction \("Scowls, Ready to Attack"\)

Assuming Snake took a faction hit of -10 for each NPC assigned to Ashen Order, he could kill 100 monks to reach the lowest possible Current Faction.

If instead, Snake turned over a new leaf and _raised_ his faction with Ashen Order, the highest his Current Faction could reach would be 1000, and he would consider "Warmly" \(he could never attain "Ally"\).  

### Consider Levels from Faction Values

Faction values for each level are as follows. The number on the left is the TOTAL value as seen in the table above \(+ or minus any effects like illusions, etc\):

| Total Faction Value | Faction CON |
| :--- | :--- |
| 1100 -&gt; ABOVE | ALLY |
| 750 -&gt; 1099 | WARMLY |
| 500 -&gt; 749 | KINDLY |
| 100 -&gt; 499 | AMIABLE |
| 0 -&gt; 99 | INDIFFERENT |
| -100 -&gt; -1 | APPREHENSIVE |
| -500 -&gt; -101 | DUBIOUS |
| -750 -&gt; -501 | THREATENLY |
| BELOW -&gt; -751 | SCOWLS |

