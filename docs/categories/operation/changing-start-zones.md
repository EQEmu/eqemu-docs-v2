# Changing Start Zones

## Description

You can change the starting zone for characters on a per race/class and deity basis, as well as for Titanium or SOF+ clients. This is separate from the Tutorial zone- this can be enabled/disabled as required. Note: the tutorial may teleport players to POK regardless of what you set.

{% hint style="warning" %}
**NB:** The variable "startzone" in the variables table is no longer used. Any references to this are deprecated.
{% endhint %}

You might wish to have a different starting zone for Titanium clients than SOF+ in case you wish to use a zone that is not available in the Titanium install. If you go wrong somewhere in the process, SOF+ clients will put players in Crescent Reach. Titanium clients will end up in the old default city \(Gnomes in Ak'anon, Wood Elves in Kelethin etc\)

## Method

The method is relatively simple: First, insert a row into the rule\_values table called World:TitaniumStartZoneID \(for Titanium clients\) and one for World:SofStartZoneID \(for all other clients\) with a value that matches the corresponding zoneidnumber of the zone you wish them to start in \(you will also need to provide x,y,z values too. Ensure that the ruleset\_id field for the rule is is set to 1, or if you have specified as default ruleset in the variables table, that the ruleset \_id matches that. The slight trick here is that the variables table uses the ruleset name, but the rule\_values table uses its number,

Secondly, insert a row in start\_zones for each class/deity/race combination for each of the zones you specify. If the start zone is the same for both types of client, you only need one row per combination, but if you have two different start zones depending on client then there will need to be two . e.g. I want a different start zone for Titanium and SOF clients. For a dwarf warrior of Brell Serilis, I would need to have two entries: one with a zone id of the Titanium zone, and one with the SOF+ zone. I would need to repeat this for all combinations.

## Example

I wish all Titanium players to start in Plane of Knowledge, and all SOF+ clients to start in Goru\`kar Mesa. On my server I have defined a rule\_set with a ruleset\_id of **3** and a name of  MyDefault. In my variables table I would insert \(or update\) a row whose varname is RuleSet and has a value of MyDefault.

Then I would create a row in rule\_values with a ruleset\_id of **3**, a rule\_name of World:SofStartZoneID and a rule\_value of 397 \(Goru\`kar Mesa\)

Then I would create a row in rule\_values with a ruleset\_id of **3**, a rule\_name of World:TitmniumStartZoneID and a rule\_value of 20 \(Plane of Knowledge\).

I would then create \(or more likely update\) rows in start\_zones for each class/deity/race combination, one for zone 202, one for zone 397.

