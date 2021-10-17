# Special Attacks

## Summon \(1\)

{% hint style="info" %}
Summon, denoted by special attack 1 or "S" \(old code\), allows you to teleport the player to the NPC, or the NPC to the player.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Level 1 | Summon Target to NPC | 0 |
| Level 2 | Summon NPC to Target | 0 |
| Parameter 0 | Cooldown in ms | 6000 |
| Parameter 1 | HP ratio required to summon | 97 |

### **Example:**

```text
1,1,2,10000,90
```

|  |  |
| :--- | :--- |
| 1 | Special Attack "Summon" |
| 1 | Special Attack "Summon" enabled |
| 2 | Summon npc to target enabled |
| 10000 | 10 second cooldown |
| 90 | Requires 90 percent or less HP |

## Enrage \(2\)

{% hint style="info" %}
Enrage, denoted by special attack 2 or "E" \(old code\), allows the NPC to riposte incoming melee attacks.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Parameter 0 | HP percent required to enrage | Rule: NPC:StartEnrageValue |
| Parameter 1 | Duration of Special Attack in ms | 10000 |
| Parameter 2 | Cooldown of Special Attack in ms | 360000 |

### **Example:**

```text
2,1,25,8000,60000
```

|  |  |
| :--- | :--- |
| 2 | Special Attack "Enrage" |
| 1 | Special Attack "Enrage" enabled |
| 25 | Begin Special Attack when HP reaches 25 percent |
| 8000 | Special Attack will last 8 seconds |
| 60000 | Special Attack will cooldown in 1 minute |

## Rampage \(3\)

{% hint style="info" %}
Rampage, denoted by special attack 3 or "R" \(old code\), allows the NPC to attack additional players.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Parameter 0 | Percent chance of Special Attack Proc | 20 percent |
| Parameter 1 | Special Attack Target Count | Rule: Combat:MaxRampageTargets |
| Parameter 2 | Percent of a normal attack damage to deal | 100 |
| Parameter 3 | Flat damage bonus to add to the Special Attack | 0 |
| Parameter 4 | Percent of AC ignored for the Special Attack | 0 |
| Parameter 5 | Flat amount of AC ignored for the Special Attack | 0 |
| Parameter 6 | Percent of natural crit used by this Special Attack | 100 |
| Parameter 7 | Flat crit bonus to add to base crit for the Special Attack | 0 |

### **Example:**

```text
3,1,15,1,75
```

|  |  |
| :--- | :--- |
| 3 | Special Attack "Rampage" |
| 1 | Special Attack "Rampage" enabled |
| 15 | 15 percent chance to proc |
| 1 | 1 target |
| 75 | 75 percent of normal damage |

## Area Rampage \(4\)

{% hint style="info" %}
Area Rampage, denoted by special attack 4 or "r" \(old code\), allows the NPC to attack additional players.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Parameter 0 | Percent chance of Special Attack Proc | 20 percent |
| Parameter 1 | Special Attack Target Count | -1 \(all in range\) |
| Parameter 2 | Percent of a normal attack damage to deal | 100 |
| Parameter 3 | Flat damage bonus to add to the Special Attack | 0 |
| Parameter 4 | Percent of AC ignored for the Special Attack | 0 |
| Parameter 5 | Flat amount of AC ignored for the Special Attack | 0 |
| Parameter 6 | Percent of natural crit used by this Special Attack | 100 |
| Parameter 7 | Flat crit bonus to add to base crit for the Special Attack | 0 |

### **Example:**

```text
4,1,15,-1,75
```

|  |  |
| :--- | :--- |
| 4 | Special Attack "Area Rampage" |
| 1 | Special Attack "Area Rampage" enabled |
| 15 | 15 percent chance to proc |
| -1 | All targets within range |
| 75 | 75 percent of normal damage |

## Flurry \(5\)

{% hint style="info" %}
Flurry, denoted by special attack 5 or "F" \(old code\), allows the NPC to have additional melee attacks against a player.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Parameter 0 | Percent chance of Special Attack Proc | Rule: Combat:NPCFlurryChance |
| Parameter 1 | Special Attack count | Rule: Combat:MaxFlurryHits |
| Parameter 2 | Percent of a normal attack damage to deal | 100 |
| Parameter 3 | Flat damage bonus to add to the Special Attack | 0 |
| Parameter 4 | Percent of AC ignored for the Special Attack | 0 |
| Parameter 5 | Flat amount of AC ignored for the Special Attack | 0 |
| Parameter 6 | Percent of natural crit used by this Special Attack | 100 |
| Parameter 7 | Flat crit bonus to add to base crit for the Special Attack | 0 |

### **Example:**

```text
5,1,25,5,50
```

|  |  |
| :--- | :--- |
| 5 | Special Attack "Flurry" |
| 1 | Special Attack "Flurry" enabled |
| 25 | 25 percent chance to proc |
| 5 | 5 added attacks |
| 50 | 50 percent of normal damage |

## Triple Attack \(6\)

{% hint style="info" %}
Triple Attack, denoted by special attack 6 or "T" \(old code\), allows the NPC to have three melee attacks against a player.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
6,1
```

|  |  |
| :--- | :--- |
| 6 | Special Attack "Triple Attack" |
| 1 | Special Attack "Triple Attack" enabled |

## Quad Attack \(7\)

{% hint style="info" %}
Quad Attack, denoted by special attack 7 or "Q" \(old code\), allows the NPC to have four melee attacks against a player.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
7,1
```

|  |  |
| :--- | :--- |
| 7 | Special Attack "Quad Attack" |
| 1 | Special Attack "Quad Attack" enabled |

## Dual Wield \(8\)

{% hint style="info" %}
Dual Wield, denoted by special attack 8 or "L" \(old code\), allows the NPC to utilize two weapons.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
8,1
```

|  |  |
| :--- | :--- |
| 8 | Special Attack "Dual Wield" |
| 1 | Special Attack "Dual Wield" enabled |

## Bane Attack \(9\)

{% hint style="info" %}
Bane Attack, denoted by special attack 9 or "b" \(old code\), gives the NPC the ability to hit NPCs that require bane weapons to damage.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
9,1
```

|  |  |
| :--- | :--- |
| 9 | Special Attack "Bane Attack" |
| 1 | Special Attack "Bane Attack" enabled |

## Magical Attack \(10\)

{% hint style="info" %}
Magical Attack, denoted by special attack 10 or "m" \(old code\), gives the NPC the ability to hit NPCs that require magical weapons to damage.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
10,1
```

|  |  |
| :--- | :--- |
| 10 | Special Attack "Magical Attack" |
| 1 | Special Attack "Magical Attack" enabled |

## Ranged Attack \(11\)

{% hint style="info" %}
Ranged Attack, denoted by special attack 11 or "Y" \(old code\), gives the NPC the ability to perform ranged attacks if their target is out of melee range.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Parameter 0 | Minimum ranged distance | 25 |
| Parameter 1 | Maximum ranged distance | 250 |
| Parameter 2 | Percent hit chance modifier | 0 |
| Parameter 3 | Percent damage modifier | 0 |

### **Example:**

```text
11,1,50
```

|  |  |
| :--- | :--- |
| 11 | Special Attack "Ranged Attack" |
| 1 | Special Attack "Ranged Attack" enabled |
| 50 | Minimum distance 50 units |

## Unslowable \(12\)

{% hint style="info" %}
Unslowable, denoted by special attack 12 or "U" \(old code\), makes the NPC immune to slow effects.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
12,1
```

|  |  |
| :--- | :--- |
| 12 | Special Attack "Unslowable" |
| 1 | Special Attack "Unslowable" enabled |

## Unmezable \(13\)

{% hint style="info" %}
Unmezable, denoted by special attack 13 or "M" \(old code\), makes the NPC immune to mesmerization effects.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
13,1
```

|  |  |
| :--- | :--- |
| 13 | Special Attack "Unmezable" |
| 1 | Special Attack "Unmezable" enabled |

## Uncharmable \(14\)

{% hint style="info" %}
Uncharmable, denoted by special attack 14 or "C" \(old code\), makes the NPC immune to charm effects.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
14,1
```

|  |  |
| :--- | :--- |
| 14 | Special Attack "Uncharmable" |
| 1 | Special Attack "Uncharmable" enabled |

## Unstunable \(15\)

{% hint style="info" %}
Unstunable, denoted by special attack 15 or "N" \(old code\), makes the NPC immune to stun effects.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
15,1
```

|  |  |
| :--- | :--- |
| 15 | Special Attack "Unstunable" |
| 1 | Special Attack "Unstunable" enabled |

## Unsnareable \(16\)

{% hint style="info" %}
Unsnareable, denoted by special attack 16 or "I" \(old code\), makes the NPC immune to snare effects.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
16,1
```

|  |  |
| :--- | :--- |
| 16 | Special Attack "Unsnareable" |
| 1 | Special Attack "Unsnareable" enabled |

## Unfearable \(17\)

{% hint style="info" %}
Unfearable, denoted by special attack 17 or "D" \(old code\), makes the NPC immune to fear effects.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
17,1
```

|  |  |
| :--- | :--- |
| 17 | Special Attack "Unfearable" |
| 1 | Special Attack "Unfearable" enabled |

## Immune to Dispell \(18\)

{% hint style="info" %}
Immune to Dispell, denoted by special attack 18 or "K" \(old code\), makes the NPC immune to dispelling spells.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
18,1
```

|  |  |
| :--- | :--- |
| 18 | Special Attack "Immune to Dispell" |
| 1 | Special Attack "Immune to Dispell" enabled |

## Immune to Melee \(19\)

{% hint style="info" %}
Immune to Melee, denoted by special attack 19 or "A" \(old code\), makes the NPC immune to all melee damage.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
19,1
```

|  |  |
| :--- | :--- |
| 19 | Special Attack "Immune to Melee" |
| 1 | Special Attack "Immune to Melee" enabled |

## Immune to Magic \(20\)

{% hint style="info" %}
Immune to Magic, denoted by special attack 20 or "B" \(old code\), makes the NPC immune to all magic spells.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
20,1
```

|  |  |
| :--- | :--- |
| 20 | Special Attack "Immune to Magic" |
| 1 | Special Attack "Immune to Magic" enabled |

## Immune to Fleeing \(21\)

{% hint style="info" %}
Immune to Fleeing, denoted by special attack 21 or "f" \(old code\), prevents the NPC from fleeing under any circumstances.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
21,1
```

|  |  |
| :--- | :--- |
| 21 | Special Attack "Immune to Fleeing" |
| 1 | Special Attack "Immune to Fleeing" enabled |

## Immune to Non-Bane Damage \(22\)

{% hint style="info" %}
Immune to Non-Bane Damage, denoted by special attack 22 or "O" \(old code\), prevents the NPC from being damaged by weapons that don't have the bane type matching its bodytype.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
22,1
```

|  |  |
| :--- | :--- |
| 22 | Special Attack "Immune to Non-Bane Damage" |
| 1 | Special Attack "Immune to Non-Bane Damage" enabled |

## Immune to Non-Magical Damage \(23\)

{% hint style="info" %}
Immune to Non-Magical Damage, denoted by special attack 23 or "W" \(old code\), prevents the NPC from being damaged by weapons that are not magical.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
23,1
```

|  |  |
| :--- | :--- |
| 23 | Special Attack "Immune to Non-Bane Damage" |
| 1 | Special Attack "Immune to Non-Bane Damage" enabled |

## Will Not Aggro \(24\)

{% hint style="info" %}
Will Not Aggro, denoted by special attack 24 or "H" \(old code\), prevents a player from getting on an NPC's hate list.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
24,1
```

|  |  |
| :--- | :--- |
| 24 | Special Attack "Will Not Aggro" |
| 1 | Special Attack "Will Not Aggro" enabled |

## Immune to Aggro \(25\)

{% hint style="info" %}
Immune to Aggro, denoted by special attack 25 or "G" \(old code\), prevents the NPC from getting on an NPC's hate list.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
25,1
```

|  |  |
| :--- | :--- |
| 25 | Special Attack "Immune to Aggro" |
| 1 | Special Attack "Immune to Aggro" enabled |

## Resist Ranged Spells \(26\)

{% hint style="info" %}
Resist Ranged Spells, denoted by special attack 26 or "g" \(old code\), prevents the NPC from being damaged from spells cast outside of melee range \("belly caster" mob\).
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
26,1
```

|  |  |
| :--- | :--- |
| 26 | Special Attack "Resist Ranged Spells" |
| 1 | Special Attack "Resist Ranged Spells" enabled |

## See through Feign Death \(27\)

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
27,1
```

|  |  |
| :--- | :--- |
| 27 | Special Attack "See through Feign Death" |
| 1 | Special Attack "See through Feign Death" enabled |

## Immune to Taunt \(28\)

{% hint style="info" %}
Immune to Taunt, denoted by special attack 28 or "i" \(old code\), prevents the NPC from taunted by players.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
28,1
```

|  |  |
| :--- | :--- |
| 28 | Special Attack "Immune to Taunt" |
| 1 | Special Attack "Immune to Taunt" enabled |

## Tunnel Vision \(29\)

{% hint style="info" %}
Tunnel Vision, denoted by special attack 29 or "t" \(old code\), makes anyone not on the top of the hate list generate a different amount of hate.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Parameter 0 | aggro modifier | Rule: Aggro:TunnelVisionAggroMod |

### **Example:**

```text
29,1,0
```

|  |  |
| :--- | :--- |
| 29 | Special Attack "Tunnel Vision" |
| 1 | Special Attack "Tunnel Vision" enabled |
| 0 | aggro modifier 0 |

## Does NOT buff/heal friends \(30\)

{% hint style="info" %}
Does NOT buff/heal friends, denoted by special attack 30 or "n" \(old code\), makes the NPC NOT buff or heal members of the same faction.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
30,1
```

|  |  |
| :--- | :--- |
| 30 | Special Attack "Does NOT buff/heal friends" |
| 1 | Special Attack "Does NOT buff/heal friends" enabled |

## Unpacifiable \(31\)

{% hint style="info" %}
Unpacifiable, denoted by special attack 31 or "p" \(old code\), makes the NPC immune to lull effects.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
31,1
```

|  |  |
| :--- | :--- |
| 31 | Special Attack "Unpacifiable" |
| 1 | Special Attack "Unpacifiable" enabled |

## Leashed \(32\)

{% hint style="info" %}
Leashed, denoted by special attack 32 or "J" \(old code\), makes the NPC return to their aggro point, fully heal, and wipes their hate list if the NPC is pulled out of a particular range.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Parameter 0 | range | aggro range \* aggro range |

### **Example:**

```text
32,1,150
```

|  |  |
| :--- | :--- |
| 32 | Special Attack "Leashed" |
| 1 | Special Attack "Leashed" enabled |
| 150 | Leash at 150 units |

## Tethered \(33\)

{% hint style="info" %}
Tethered, denoted by special attack 33 or "j" \(old code\), is used to leashe the mob to their aggro range.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Parameter 0 | range | aggro range \* aggro range |

### **Example:**

```text
33,1,100
```

|  |  |
| :--- | :--- |
| 33 | Special Attack "Tethered" |
| 1 | Special Attack "Tethered" enabled |
| 100 | Leash at 100 units |

## Destructible Object \(34\)

{% hint style="info" %}
Destructible Object, denoted by special attack 34 or "o" \(old code\), is used on destructing NPCs. \(Deprecated?\)
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
34,1
```

|  |  |
| :--- | :--- |
| 34 | Special Attack "Destructible Object" |
| 1 | Special Attack "Destructible Object" enabled |

## No Harm from Players \(35\)

{% hint style="info" %}
No Harm from Players, denoted by special attack 35 or "Z" \(old code\), prevents players from being able to harm the NPC in any way.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
35,1
```

|  |  |
| :--- | :--- |
| 35 | Special Attack "No Harm from Players" |
| 1 | Special Attack "No Harm from Players" enabled |

## Always Flee \(36\)

{% hint style="info" %}
Always Flee, denoted by special attack 36, makes the NPC flee at low health even if faction allies are near.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
36,1
```

|  |  |
| :--- | :--- |
| 36 | Special Attack "Always Flee" |
| 1 | Special Attack "Always Flee" enabled |

## Flee Percentage \(37\)

{% hint style="info" %}
Flee Percentage, denoted by special attack 37, makes the NPC flee at low health even if faction allies are near, at a given percent.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Level 1 | Percent NPC will flee at |  |
| Parameter 0 | Percent chance to flee |  |

### **Example:**

```text
37,10
```

|  |  |
| :--- | :--- |
| 37 | Special Attack "Flee Percentage" |
| 10 | Flee when health reaches 10 percent |

## Allow Beneficial \(38\)

{% hint style="info" %}
Allow Beneficial, denoted by special attack 38, allows players to cast beneficial spells on the NPC.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
38,1
```

|  |  |
| :--- | :--- |
| 38 | Special Attack "Allow Beneficial" |
| 1 | Special Attack "Allow Beneficial" enabled |

## Disable Melee \(39\)

{% hint style="info" %}
Disable Melee, denoted by special attack 39, makes the NPC unable to melee, but does allow the NPC to aggro.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
39,1
```

|  |  |
| :--- | :--- |
| 39 | Special Attack "Disable Melee" |
| 1 | Special Attack "Disable Melee" enabled |

## Chase Distance \(40\)

{% hint style="info" %}
Chase Distance, denoted by special attack 40, establishes the minimum and maximum distances between the NPC and an aggro player.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Parameter 0 | Maximum chase distance |  |
| Parameter 1 | Minimum chase distance |  |
| Parameter 2 | Ignore line of sight |  |

### **Example:**

```text
40,1,200,20
```

|  |  |
| :--- | :--- |
| 40 | Special Attack "Chase Distance" |
| 1 | Special Attack "Chase Distance" enabled |
| 200 | NPC will chase you if you're greater than 200 units away |
| 20 | NPC will chase you if you're less than 20 units away |

## Allow Tank \(41\)

{% hint style="info" %}
Allow Tank, denoted by special attack 41, sets the NPC to allow other NPCs to take aggro from players.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
41,1
```

|  |  |
| :--- | :--- |
| 41 | Special Attack "Allow Tank" |
| 1 | Special Attack "Allow Tank" enabled |

## Ignore Root Aggro \(42\)

{% hint style="info" %}
Ignore Root Aggro, denoted by special attack 42, sets the NPC to ignore the rules of root aggro--the NPC will not attack the closest player, but rather the player on top of the hate list.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Bool | 1 = on, 0 = off | 0 |

### **Example:**

```text
42,1
```

|  |  |
| :--- | :--- |
| 42 | Special Attack "Ignore Root Aggro" |
| 1 | Special Attack "Ignore Root Aggro" enabled |

## Casting Resist Diff \(43\)

{% hint style="info" %}
Casting Resist Diff, denoted by special attack 43, makes the NPC's spells cast at a different resist level.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Parameter 0 | Flat modifier to the resist diff of the spell | 0 |

### **Example:**

```text
43,1,-200
```

|  |  |
| :--- | :--- |
| 43 | Special Attack "Casting Resist Diff" |
| 1 | Special Attack "Casting Resist Diff" enabled |
| -200 | Add -200 to the resist modifier for the spell cast |

## Counter Avoid Damage \(44\)

{% hint style="info" %}
Counter Avoid Damage, denoted by special attack 44, makes the NPC more likely to hit a player, decreasing their chance for at avoiding melee through dodge/parry/riposte/etc.
{% endhint %}

| Settings |  | Defaults |
| :--- | :--- | :--- |
| Parameter 0 | Flat negative percent modifier to ALL avoidance skills | 0 |
| Parameter 1 | Flat negative percent modifier to Riposte skill | 0 |
| Parameter 2 | Flat negative percent modifier to Parry skill | 0 |
| Parameter 3 | Flat negative percent modifier to Block skill | 0 |
| Parameter 4 | Flat negative percent modifier to Dodge skill | 0 |

### **Examples:\*\***

```text
44,1,50
```

|  |  |
| :--- | :--- |
| 44 | Special Attack "Counter Avoid Damage" |
| 1 | Special Attack "Counter Avoid Damage" enabled |
| 50 | 50 percent reduction to ALL avoidance skills |

```text
44,1,0,0,0,0,50
```

|  |  |
| :--- | :--- |
| 44 | Special Attack "Counter Avoid Damage" |
| 1 | Special Attack "Counter Avoid Damage" enabled |
| 0 | 0 percent reduction to ALL avoidance skills |
| 0 | 0 percent reduction to Riposte avoidance skill |
| 0 | 0 percent reduction to Parry avoidance skill |
| 0 | 0 percent reduction to Block avoidance skill |
| 50 | 50 percent reduction to Dodge avoidance skill |

## Proximity Aggro \(45\)

{% hint style="info" %}
Proximity Aggro, denoted by special attack 45, allows the NPC to engage new clients while in combat if the client is within their proximity.
{% endhint %}

### **Example:**

```text
45,1
```

|  |  |
| :--- | :--- |
| 45 | Special Attack "Proximity Aggro" |
| 1 | Special Attack "Proximity Aggro" enabled |

## Immune Ranged Attacks \(46\)

{% hint style="info" %}
Immune Ranged Attacks, denoted by special attack 46, makes the NPC immune to ranged attacks.
{% endhint %}

### **Example:**

```text
46,1
```

|  |  |
| :--- | :--- |
| 46 | Special Attack "Immune Ranged Attacks" |
| 1 | Special Attack "Immune Ranged Attacks" enabled |

## Immune Client Damage \(47\)

{% hint style="info" %}
Immune Client Damage, denoted by special attack 47, makes the NPC immune to damage by clients.
{% endhint %}

### **Example:**

```text
47,1
```

|  |  |
| :--- | :--- |
| 47 | Special Attack "Immune Client Damage" |
| 1 | Special Attack "Immune Client Damage" enabled |

## Immune NPC Damage \(48\)

{% hint style="info" %}
Immune NPC Damage, denoted by special attack 48, makes the NPC immune to damage by NPCs.
{% endhint %}

### **Example:**

```text
48,1
```

|  |  |
| :--- | :--- |
| 48 | Special Attack "Immune NPC Damage" |
| 1 | Special Attack "Immune NPC Damage" enabled |

## Immune Client Aggro \(49\)

{% hint style="info" %}
Immune Client Aggro, denoted by special attack 49, makes the NPC immune to aggro by clients.
{% endhint %}

### **Example:**

```text
49,1
```

|  |  |
| :--- | :--- |
| 49 | Special Attack "Immune Client Aggro" |
| 1 | Special Attack "Immune Client Aggro" enabled |

## Immune NPC Aggro \(50\)

{% hint style="info" %}
Immune NPC Aggro, denoted by special attack 50, makes the NPC immune to aggro by NPCs.
{% endhint %}

### **Example:**

```text
50,1
```

|  |  |
| :--- | :--- |
| 50 | Special Attack "Immune NPC Aggro" |
| 1 | Special Attack "Immune NPC Aggro" enabled |

