# Entity Lists

## Description

**\* Manipulating entity lists is EXTREMELY powerful and if you were smart you would read this article.**

This article also assumes you have read other articles and have a relatively basic understanding of the quest/scripting engine.

To start off, an entity could be any of the following:

* **NPC** ($npc)
* **Mob** ($mob Almost the same thing as NPC, can be used interchangeably) 
* **Client** ($client)
* **Door** ($door - not implied)
* **Object** ($object - not implied)
* **Corpse** ($corpse - not implied)

In this simple example below you can always assume that the **client** is the player who **'triggered'** the script to occur, and the **NPC** who is bound to the script that has triggered:[?](http://wiki.eqemulator.org/p?Entity_Lists\_-\_How_to_Use_Them\&frm=Ultimate_Perl_Reference#)

{% tabs %}
{% tab title="Perl" %}
```perl
sub EVENT_SAY {
	if($text=~/hail/i) {
		$client->Message(15, "This is a client message.");
		$npc->Say("Hello there $name.");
	}
}
```
{% endtab %}

{% tab title="Lua" %}
```lua
function event_say(e)
	if (e.message:findi("hail")) then
		e.other:Message(15, "This is a client message.");
		e.self:Say("Hello there " .. e.other:GetCleanName() . ".");
	end
end
```
{% endtab %}
{% endtabs %}

Some of you may already know this as common knowledge, but in order for this to cover everyone's knowledge levels I will preamble with the basics a bit.

## What are Entity Objects?

Entity objects are used most commonly as a way encorporate the simple example above, I have one **$client** that **triggered** the script and a **$npc** that was triggered when I **hailed** the npc, but what they were really intended for is so you can flexibly 'select' any NPC in the zone when trying to perform any specific function.

We can essentially use a **Controller NPC** or any **NPC **for that matter to control any other NPC in the zone using entity objects, you just need to know how to select them and use them to your extreme advantage.

**Crushbone**

I am going to use Crushbone as a simple example here.

This zone by default has 119 NPC's approximately. Most if not all of them are Orc's and some prisoners and such.

| ![](http://everquest.allakhazam.com/scenery/crushbone-scenery.jpg) | ![](http://everquest.allakhazam.com/scenery/crushbone-ironorc.jpg) | ![](http://everquest.allakhazam.com/scenery/crushbone-trainer.jpg) |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |

* Below is a list of NPC's in the zone, we're simply using this as a reference of what kind of NPC's are in the zone. There are other ways to get this information using plugins and handy tools but for now simply for illustration purposes we want to take a quick survery of what we have in the zone and how many NPC's of certain name types we have.
* Below is simpy the spawn data, it does not mean that all of these NPC's are what is spawned at all times in the database, these are the unique entries that are in the database. Once again there are about 119 NPC's with the zone fully popped.

{% tabs %}
{% tab title="Query" %}
```sql
SELECT npc_types.id, COUNT(npc_types.id) AS `npc_count`,
npc_types.`name`, npc_types.lastname, npc_types.`level` FROM(spawn2 ,npc_types)
INNER JOIN spawnentry ON npc_types.id = spawnentry.npcID
AND spawnentry.spawngroupID = spawn2.spawngroupID
WHERE spawn2.zone = 'crushbone' GROUP BY `npc_types`.`name`;
```
{% endtab %}

{% tab title="Result" %}
```
+-------+-----------+-----------------------------+----------+-------+
| id    | npc_count | name                        | lastname | level |
+-------+-----------+-----------------------------+----------+-------+
| 58059 |         1 | Ambassador_DVinn            |          |    20 |
| 58019 |         1 | an_elven_priest             |          |     5 |
| 58007 |         9 | an_elven_slave              |          |     3 |
| 58043 |         1 | an_orc_thaumaturgist        |          |    12 |
| 58020 |         9 | a_dwarven_slave             |          |     2 |
| 58018 |         3 | a_dwarven_smith             |          |     3 |
| 58032 |         1 | Emperor_Crush               |          |    18 |
| 58056 |         1 | Kelynn                      |          |     9 |
| 58028 |         1 | Lord_Darish                 |          |    14 |
| 58015 |       114 | orc_centurion               |          |     9 |
| 58025 |         6 | orc_emissary                |          |    15 |
| 58008 |        33 | orc_legionnaire             |          |    12 |
| 58024 |         9 | orc_oracle                  |          |    13 |
| 58001 |        25 | orc_pawn                    |          |     1 |
| 58042 |         1 | Orc_Scoutsman               |          |    12 |
| 58016 |        15 | orc_slaver                  |          |     5 |
| 58040 |         1 | orc_taskmaster              |          |    14 |
| 58035 |         1 | orc_trainer                 |          |    13 |
| 58047 |         1 | Orc_Warden                  |          |    14 |
| 58002 |         1 | orc_warlord                 |          |    16 |
| 58017 |         1 | Retlon_Brenclog             |          |    16 |
| 58010 |         1 | Rondo_Dunfire               |          |     9 |
| 58030 |         1 | royal_guard                 |          |    15 |
| 58058 |         1 | The_Fabled_Ambassador_Dvinn |          |    25 |
| 58057 |         1 | The_Fabled_Emperor_Crush    |          |    25 |
| 58031 |         1 | The_Prophet                 |          |    17 |
+-------+-----------+-----------------------------+----------+-------+
26 rows in set
```
{% endtab %}
{% endtabs %}

## Manipulating the Entity List

Now, we can start getting to the real nuts and bolts of practical application here. Let's say I'm on a mission to liven up a custom version of my zone, and whenever Emperor Crush spawns, I want all of my Orc's in the zone to howl and beat their chest or swing their arm or do something provocative as well as make an emote we would use an example such as below.

{% tabs %}
{% tab title="Perl" %}
```perl
sub EVENT_SPAWN {
	quest::shout("Graghhhh! Who goes there!"); #:: Ruh roh, Emperor is pissed I think!
	quest::settimer("get_zone_wriled_up", 3); #:: Use a timer so that if we #repop the zone, we give everything a little bit of time to pop before trying to manipulate the rest of the NPC's
}

sub EVENT_TIMER {
	if($timer eq "get_zone_wriled_up") { #:: Triggering timer event
		quest::shout("lol"); #:: Some test thing message I left in apparently.
		my @nlist = $entity_list->GetNPCList(); #:: Use Entity List Object to get this
		foreach my $n (@nlist){ #:: Iterate through each NPC in the @nlist array
			if($n->GetCleanName()=~/orc/i){ #:: We are doing a simple string comparison, GetCleanName should be getting the name of the NPC, we want to match if it has the name 'orc' in it                $n->Say("Yargghhh"); #::: Have each NPC say 'Yargghhh' that matches the name 'orc'
				$n->Say("Yargghhh");
				$n->DoAnim(8); #:: Have each NPC do an attack aniamtion
			}
		}
		quest::stoptimer($timer); #:: Stop the timer, $timer is an exported object from the source as this event is triggered
	}
}
```
{% endtab %}

{% tab title="Lua" %}
```lua
function event_spawn(e)
	e.self:Shout("Graghhhh! Who goes there!"); --:: Ruh roh, Emperor is pissed I think!
	eq.set_timer("get_zone_wriled_up", 3);  --:: Use a timer so that if we #repop the zone, we give everything a little bit of time to pop before trying to manipulate the rest of the NPC's
end

function event_timer(e)
	if (e.timer == "get_zone_wriled_up") then --:: Triggering timer event
		e.self:Shout("lol"); --:: Some test thing message I left in apparently.
		local nlist = entity_list:GetNPCList(); --:: Use Entity List Object to get this
		for i, npc in ipairs(nlist) do --:: Iterate through each NPC in the @nlist array
			if (npc:GetCleanName():findi("orc")) then --:: We are doing a simple string comparison, GetCleanName should be getting the name of the NPC, we want to match if it has the name 'orc' in it                $n->Say("Yargghhh"); #::: Have each NPC say 'Yargghhh' that matches the name 'orc'
				npc:Say("Yargghhh");
				npc:DoAnim(8); --:: Have each NPC do an attack aniamtion
			end
		end
		eq.stop_timer("get_zone_wriled_up"); --:: Stop the timer
	end
end
```
{% endtab %}
{% endtabs %}

* In this very quick example you can see that I used a timer as a delay to spawn the NPC's, and then I used a simply entity list iteration to have NPC's do what I want them to do

![](http://i.imgur.com/gAoOpRc.jpg)

As you can see they Orc's are all performing animation 8 which is a simple attack animation, for animtions see the [Animation Reference List](../npc/animations.md).

![](https://firebasestorage.googleapis.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-LnsZsY32yEVJrtyioON%2Fuploads%2F2aCj5Q0vKj7xBbTeJDQj%2Ffile.png?alt=media)

You also see orcs nearby you are saying **'Yargghh''**

## **Iteration Example**

* **Now, lets say I want to use the same iteration loop to do stuff to other NPC's. Let's say I want to have the slaves shout 'HELP!' - this is getting a bit ridiculous isn't it? Who cares!**

{% tabs %}
{% tab title="Perl" %}
```perl
sub EVENT_SPAWN {
	quest::shout("Graghhhh! Who goes there!");
	quest::settimer("get_zone_wriled_up", 3);
}

sub EVENT_TIMER {
	if($timer eq "get_zone_wriled_up") {
		quest::shout("lol");
		my @nlist = $entity_list->GetNPCList();
		foreach my $n (@nlist){
			if($n->GetCleanName()=~/orc/i){
				$n->Say("Yargghhh");
				$n->DoAnim(8);
			}
			if($n->GetCleanName()=~/slave/i){
				$n->Shout("HELP!");
			}
		}
		quest::stoptimer($timer);
	}
}
```
{% endtab %}

{% tab title="Lua" %}
```lua
function event_spawn(e)
	e.self:Shout("Graghhhh! Who goes there!");
	eq.set_timer("get_zone_wriled_up", 3);
end

function event_timer(e)
	if (e.timer == "get_zone_wriled_up") then
		e.self:Shout("lol");
		local nlist = entity_list:GetNPCList();
		for i, npc in ipairs(nlist) do
			if (npc:GetCleanName():findi("orc")) then
				npc:Say("Yargghhh");
				npc:DoAnim(8);
			end
			if (npc:GetCleanName():findi("slave")) then
				npc:Shout("HELP!");
			end
		end
		eq.stop_timer("get_zone_wriled_up");
	end
end
```
{% endtab %}
{% endtabs %}

*  Now, I see exactly what we would expect, the slaves shouting for help, of course you could just have the elven scripts shout in their script but we're using this simply as an example

![](https://firebasestorage.googleapis.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-LnsZsY32yEVJrtyioON%2Fuploads%2FNFb0vJedzuIFIiYTaF16%2Ffile.png?alt=media)

* The same exact thing applies for clients, you can use any $client object listed in the [Quest API](https://eqemu.gitbook.io/quest-api/) manipulate the entity list and trigger any action in any event that you want. It is a less frequent thing to want to use a full list of clients to do something, but you can use it for a variety of things depending on the cool ideas that you come up with. But for examples sake I want to use an example where maybe I want to freeze all of the clients in a zone in a boss fight with **Emperor Crushbone **for many a few seconds at a time

## Client Example

{% tabs %}
{% tab title="Perl" %}
```perl
sub EVENT_COMBAT {
	if($combat_state == 1) { #::: Attack state is on for NPC
		my @clist = $entity_list->GetClientList();
		foreach my $c (@clist) {
			$c->Message(15, "You are frozen!");
			$c->Freeze();
		}
	}
}
```
{% endtab %}

{% tab title="Lua" %}
```lua
function event_combat(e)
	if (e.joined) then
		local client_list = entity_list:GetClientList();
		for i, client in ipairs(client_list) do
			client:Message(15, "You are frozen!");
			client:Freeze();
		end
	end
end
```
{% endtab %}
{% endtabs %}

* And you can see below that when I attack **Emperor Crush,** I and every single other $client in the zone would be frozen until I did something to unfreeze the client. I literally had to force exit my game client because I lost complete control of my client and I could not do anything at this point. So it would be a good idea to set a timer for 5 seconds to also have all clients **$client->UnFreeze();** once again to gain control.You may not want to actually use this example in your quests as there are spells that can achieve the same effect, this is simply for demonstrational purposes only.

![](https://firebasestorage.googleapis.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-LnsZsY32yEVJrtyioON%2Fuploads%2FOJPIBKNj2ObnBXcMdvUg%2Ffile.png?alt=media)

## Using Entity Selectors

We've already used some selectors in previous examples, but to to give you an idea, below are some of the selectors that we can use to grab things and do stuff with them:

| Method                                                  |
| ------------------------------------------------------- |
| GetClientByAccID(accid)                                 |
| GetClientByCharID(iCharID)                              |
| GetClientByID(id)                                       |
| GetClientByName(name)                                   |
| GetClientByWID(iWID)                                    |
| GetClientList()                                         |
| GetCorpseByID(id)                                       |
| GetCorpseByName(name)                                   |
| GetCorpseByOwner(client)                                |
| GetCorpseList()                                         |
| GetDoorsByDBID(id)                                      |
| GetDoorsByID(id)                                        |
| GetDoorsList()                                          |
| GetGroupByClient(client)                                |
| GetGroupByID(id)                                        |
| GetGroupByLeaderName(leader)                            |
| GetGroupByMob(mob)                                      |
| GetMob(name)                                            |
| GetMobByID(id)                                          |
| GetMobByNpcTypeID(get_id)                               |
| GetMobID(id)                                            |
| GetMobList()                                            |
| GetNPCByID(id)                                          |
| GetNPCByNPCTypeID(npc_id)                               |
| GetNPCList()                                            |
| GetObjectByDBID()                                       |
| GetObjectByID()                                         |
| GetObjectList()                                         |
| GetRaidByClient(client)                                 |
| GetRaidByID(id)                                         |
| GetRandomClient(x, y, z, range, ClientToExclude = NULL) |

So let's say, I wanted to get something by any of these above? I use examples very similar to the one above, except for getting a single 'entity' does not require me to 'loop' through all the entries in a foreach or things of the like. If I wanted to for example just get ONE NPC, a very common method to do this is to use **$entity_list->GetNPCByNPCTypeID(YourNPCIDGoesHere);**

* **You would use this example if you are specifically trying to select a unique NPC in the zone that there is only one NPC ID of that NPC, to knock out two birds in one stone, I'm going to make this specific NPC attack a player by using a player selector.**
  * **In this example I will have Emperor Crush get Lord Darnish NPCID (58028) attack the player as well just incase the player hacked his way to the tower**

{% tabs %}
{% tab title="Perl" %}
```perl
sub EVENT_COMBAT {
	if($combat_state == 1) { #::: Attack state is on for NPC
		my $lord_darnish = $entity_list->GetNPCByNPCTypeID(58028); #::: Lord Darnish
		$lord_darnish->Attack($client); #::: We use $client because it is implied in combat
	}
}
```
{% endtab %}

{% tab title="Lua" %}
```lua
function event_combat(e)
	if (e.joined) then Attack state is on for NPC
		local lord_darnish = entity_list:GetNPCByNPCTypeID(58028); -- Lord Darnish
		lord_darnish:Attack(e.self);
	end
end
```
{% endtab %}
{% endtabs %}

* **Once again if you are trying to select similar NPC's you need to use an entity list iteration (foreach loop) and then use entity objects to filter on their data, few one line examples below:**
  * **if($n->GetClass() == 1) #::: Get all NPC's that are warriors**
  * **if($n->GetCleanName()=\~/orc/i) #::: Check to see if entity has 'orc' in its name**
  * **if($n->GetNPCTypeID() == 58000) #::: Check for NPC's that have 58000 for an NPC Type ID**
  * **if($n->GetRace() == 54) #::: Would select NPC's that have race 54**

## **More on Selectors**

* Keep in mind, you are not limited to the examples I have illustrated, there are many ways you can define a single object and how you define them is up to you.
* To make use of the $lord_darnish selector example above, I could not even define $lord_darnish, I just did that for readability and examples
  * When I do this: **$entity_list->GetNPCByNPCTypeID(58028)** (I can now use another pointer right here and go straight to my object)
    * So: **$entity_list->GetNPCByNPCTypeID(58028)->Attack($client); - Would be equivalent to my example above**
* When using an entity_list selector, you are simply grabbing that entity object and depending on that type, $client/$npc/$object/$door/etc you can access its methods
  * So For example if I use an entity selector that gets a client, all of $client is available to use off of that object
    * **$entity_list->GetClientByName("Akkadius")->Duck();** - Would make me duck ($client method)
    * **$entity_list->GetNPCByNPCTypeID(58028)->WipeHateList();** - Wipes that particular NPC's hatelist
* At this point just knowing this little bit of knowledge you have all kinds of possibilities in your arsenal
