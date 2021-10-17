# Bot Heal Rotations

## Description

Heal Rotations are a **specialty group** composed of **healing bot members** and **healable target mobs**.

Once a heal rotation is made **active**, its pseudo-AI will **augment a healing member's AI** to determine when they should **perform healing tasks**.

The **original heal rotation** code was **less sophisticated** and was **incorporated** into the **bot class** itself.

The **new heal rotation** code is a stand-alone class that has **many options** and **adjustable behavior**.

As a result of the **system change**, it may have **issues** that need to be worked out..many of the features are **experimental**.

Please **post any issues** at [http://www.eqemulator.org/forums/index.php](http://www.eqemulator.org/forums/index.php).

## Heal Rotation Commands

```text
^healrotation - Lists the available bot heal rotation [subcommands]
^healrotationadaptivetargeting - Enables or disables adaptive targeting within the heal rotation instance
^healrotationaddmember - Adds a bot to a heal rotation instance
^healrotationaddtarget - Adds target to a heal rotation instance
^healrotationadjustcritical - Adjusts the critial HP limit of the heal rotation instance's Class Armor Type criteria
^healrotationadjustsafe - Adjusts the safe HP limit of the heal rotation instance's Class Armor Type criteria
^healrotationcastingoverride - Enables or disables casting overrides within the heal rotation instance
^healrotationchangeinterval - Changes casting interval between members within the heal rotation instance
^healrotationclearhot - Clears the HOT of a heal rotation instance
^healrotationcleartargets - Removes all targets from a heal rotation instance
^healrotationcreate - Creates a bot heal rotation instance and designates a leader
^healrotationdelete - Deletes a bot heal rotation entry by leader
^healrotationfastheals - Enables or disables fast heals within the heal rotation instance
^healrotationlist - Reports heal rotation instance(s) information
^healrotationremovemember - Removes a bot from a heal rotation instance
^healrotationremovetarget - Removes target from a heal rotations instance
^healrotationresetlimits - Resets all Class Armor Type HP limit criteria in a heal rotation to its default value
^healrotationsave - Saves a bot heal rotation entry by leader
^healrotationsethot - Sets the HOT in a heal rotation instance
^healrotationstart - Starts a heal rotation
^healrotationstop - Stops a heal rotation
```

All **heal rotation commands** that **require a target** for use refer to the **healing bot member** for determining **heal rotation selection**.

A **member** of one heal rotation can be a **target** in its own or different one..but, not multiple heal rotations.

## Heal Rotation Options

### **Interval**

* The time in seconds between caster promptings of healers in the member pool
* Assuming that all casters are capable of healing in a given cycle, the time for a complete healing cycle would be **\#\_of\_hr\_members** x **interval\_in\_seconds** \(6 members, 5-second interval = 30 second cycle\)
* The healing cycle is not penalized for a caster's inability to cast. In this case, the caster is simply 'skipped' and the next caster is prompted until all casters have been cleared from the cycle
* If no casters are available for a given cycle \(or remainer of,\) then the cycle pool remains empty for the remainder of the interval and a new cycle is started upon its elapsing

### **Fast Heals**

* When enabled, only spells with a casting time of less than or equal to a specified time will be selected for use
* The current time requirement is 2 seconds or less

### **Casting Override**

* Casting override will interrupt any AI spell, even healing, that are being cast by the healing member when it is prompted for cycle casting
* A special heal rotation casting flag is set for any member that successfully casts for the heal rotation so that it may not be interrupted

## Heal Rotation Options - Reactive

### **Adaptive Targeting**

* There are two modes of reactive healing: **Standard** and **Adaptive Targeting**
* **Standard mode** has very basic criteria for determining a 'healable' target: **Tank** \(as marked by group,\) then **Healer**, and finally, **Everyone**
* **Adaptive Targeting mode** uses more advanced criteria for selecting a target:
  * **Heal Frequency:** Average heal recurrence rate for a given time period - good for finding hr targets that have grabbed hate away from a tank
  * **Heal Count:** Highest heal count for a given time period - backup for Heal Frequency
  * **Tank:** Lowest health of marked tanks
  * **Healer:** Lowest health of healers
  * **Critical by Armor Type:** Sorted as less or equal to 'critical' health percentage by armor type, then by ascending armor type
  * **Extended Heal Frequency:** Average heal recurrence rate for a given time period, plus occurrences from the previous time period
  * **Extended Heal Count:** Highest heal count for a given time period, plus the count from the previous time period
  * **Wounded by Armor Type:** Sorted as less than 'safe' health percentage by armor type, then by ascending armor type
  * **Lowest Health Percentage:** Sorted by lowest health percentage, no additional criteria
* Specific armor types only play a role in Adaptive Targeting as Standard is limited to 'Base'-only

## Heal Rotation Options - Proactive

### **HOT \(Heal Override Target\)**

* The HOT is a singular target that you can designate from the target pool for focused healing
* Once set, all reactive criteria are ignored and the target will be healed on interval timing - regardless of current health
* Clearing the HOT reverts the heal rotation back to its reactive healing state
* If saved, the heal rotation entry does not include the HOT target or state

## Creating a Heal Rotation

```text
usage: (<target_creator>) ^healrotationcreate ([creator_name])
    ([interval=5: 1-30(seconds)] [fastheals=off: on | off] [adaptivetargeting=off: on | off] [castingoverride=off: on | off])
```

* ^healrotationcreate - Creates a Heal Rotation with the selected bot as its creator using the default settings \(interval of **5** seconds, fast heals **off**, adaptive targeting **off**, casting override **off**\)
* ^healrotationcreate Jojo 3 off on on - Creates a Heal Rotation with 'Jojo' as its creator using the provided settings \(interval of **3** seconds, fast heals **off**, adaptive targeting **on**, casting override **on**\)

If you have previously used ^healrotationsave on the 'creator' bot, the settings will automatically load when this command is issued. \(See **Saving a Heal Rotation**\)

## Saving a Heal Rotation

```text

```

## Deleting a Heal Rotation

```text

```

## Adding Heal Rotation Pool Members

```text
usage: (<target_member>) ^healrotationaddmember [new_member_name] ([member_name])
```

* ^healrotationaddmember Jojo - Adds healer bot 'Jojo' to selected member bot's heal rotation
* ^healrotationaddmember Jojo Bozo - Adds healer bot 'Jojo' to heal rotation of member bot 'Bozo'

You can adjust the maximum size of the member pool by changing the rule '**Bots:HealRotationMaxMembers**' - current default is **24**.

## Removing Heal Rotation Pool Members

```text
usage: (<target_member>) ^healrotationremovemember ([member_name])
```

* ^healrotationremovemember - Removes selected healer bot from their heal rotation
* ^healrotationremovemember Jojo - Removes healer bot 'Jojo' from its heal rotation

Removing the last member from a heal rotation, \(whether through command use, by camping, or any other means,\) will also clear the target pool and consume the heal rotation instance.

## Adding Heal Rotation Pool Targets

```text
usage: (<target_member>) ^healrotationaddtarget [heal_target_name] ([member_name])
```

* ^healrotationaddtarget Meatshield - Adds target mob 'Meatshield' to selected member bot's heal rotation
* ^healrotationaddtarget Meatshield Bozo - Adds target mob 'Meatshield' to heal rotation of member bot 'Bozo'

You can adjust the maximum size of the target pool by changing the rule '**Bots:HealRotationMaxTargets**' - current default is **12**.

## Removing Heal Rotation Pool Targets

```text
usage: (<target_member>) ^healrotationremovetarget [heal_target_name] ([member_name])
```

* ^healrotationremovetarget Meatshield - Removes target mob 'Meatshield' from selected member bot's heal rotation
* ^healrotationremovetarget Meatshield Bozo - Removes target mob 'Meatshield' from heal rotation of member bot 'Bozo'

## Starting a Heal Rotation

```text

```

## Stopping a Heal Rotation

```text

```

## Setting a HOT

```text

```

## Clearing a HOT

```text

```

## Showing Current Stats of a Heal Rotation

```text

```

