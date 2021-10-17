---
description: An explanation of the use of Data Buckets in EQEmu.
---

# Using Data Buckets

## What are Data Buckets?

* Data buckets are a replacement to the well-known `qglobals`, but they are far more performant, reliable and simpler to use
* You can use data buckets to store values unique to anything you would like, for example
  * NPC based flags
  * Zone based flags
  * Character based flags 
  * etc.

## Functions

Data buckets exist in these 3 main functions in both Perl and LUA

```
get_data(std::string bucket_key)
set_data(std::string bucket_key, std::string bucket_value, std::string expires_in)
delete_data(std::string bucket_key)
```

## Storage

* Data buckets are stored in the \[\[data_buckets]] table and has a very simple structure
  * **id**
  * **key**
  * **value**
  * **expires**
* Expired data bucket rows will not be queryable via the quest API, they may exist in a table until 5-10 minutes have past and the server will garbage collect and wipe the table clean of expired buckets

## Usage Examples

### Perl

In this simple example you can see that we are keying by character id to set the flag to make this unique per player, we are setting it with the value of 70, and since we didn't set the optional value of unix time it never expires

```perl
if ($text =~/character-flag-test/i) {
    $key = $client->CharacterID() . "-some-epic-flag";
    quest::set_data($key, 70);
    quest::say("You have traveled far! You have a mighty (" . quest::get_data($key) . ") epic points!");
}
```

![](https://user-images.githubusercontent.com/3319450/42416821-8cef6690-823e-11e8-8b78-caec1a213f71.png)

### LUA

In this example below you can see that we set a simple global flag, we immediately access it, delete it and then try to unsuccessfully access it again because the bucket data had already been deleted

```lua
if (e.message:findi("test")) then
    e.self:Say("This is a test!")
    eq.set_data("lua_test_key", "lua_value");
    e.self:Say("We just set some bucket data with '".. eq.get_data("lua_test_key") .. "'");
    eq.delete_data("lua_test_key");
    e.self:Say("I'm going to try and access the value again... '".. eq.get_data("lua_test_key") .. "'");
end
```

![](https://user-images.githubusercontent.com/3319450/42416831-a726d71e-823e-11e8-92d5-cad58040e00c.png)

## Ways to Key Buckets

Keying is simply a way to uniquely identify a flag, if you want to make some data unique to a player, then you would need something to key uniquely to that player, such as their **character_id**. If you wanted to set a flag uniquely for a NPC for example, you could use the **npc_type_id**, for a zone you could use the **zone_id**. All of these circumstances are completely up to you and you have the entire Quest API to grab something that can make something unique!

Some of the examples below should give you some ideas!

**By Character**

```perl
$key   = $client->CharacterID() . "-some-flag";
$value = 70;
quest::set_data($key, $value);
```

**By Door (And Zone)**

```perl
sub EVENT_CLICKDOOR {
    quest::say($doorid);
    if ($doorid == 4) {
        $key   = $zoneid . '-' . $doorid . '-last-person-to-click-door";
        $value = $client->GetCleanName();

        if (quest::get_data($key)) {
            $client->Message(15, "You know... the last person to click this door was '" . quest::get_data($key) . "'");
        }

        quest::set_data($key, $value);
    }
}
```

**Result**

![](https://user-images.githubusercontent.com/3319450/42416886-d2caadae-823f-11e8-9b49-a6bd21bf8f73.png)

**Database Result**

![](https://user-images.githubusercontent.com/3319450/42416888-dfd7a966-823f-11e8-8fc6-1829ed0c35c3.png)

**By NPC**

```perl
sub EVENT_DEATH_COMPLETE {
    $key = $npc->GetNPCTypeID() . "-death-count";
    quest::set_data($key, quest::get_data($key) + 1);

    $death_count = quest::get_data($key);
    quest::shout("Man! I've died (" . $death_count . ") times in my lifetime!");
}
```

**Result**

![](https://user-images.githubusercontent.com/3319450/42416932-0eab7a32-8241-11e8-8b64-14d76a2c596f.png)

**Database**

![](https://user-images.githubusercontent.com/3319450/42416937-21a01aee-8241-11e8-9c29-9a2e1ed4356d.png)

## Expiration Examples

* Below in this LUA example we will count the number of times a player has talked to an NPC, first by checking if we've got a bucket set at all, if not we will set an expiration time on it. Each time we call set_data, it will not over-ride the original expiration time unless we pass a new time parameter

```lua
function event_say(e)
    if (e.message:findi("hail")) then

        -- Set unique key for the bucket
        local key = e.other:GetCleanName() .. "_times_talked";

        -- If the bucket is empty, we need to set it
        -- The first time we will set an expiration on this (86400 seconds)
        if (eq.get_data(key) == "") then
            eq.set_data(key, '1', 86400);
        end

        local times_talked = tonumber(eq.get_data(key));

        e.self:Say("You know... You've talked to me " .. times_talked .. " time(s) today, get a life will ya!");

        -- Increment times talked
        eq.set_data(key, tostring(times_talked + 1));
    end

end
```

**Result**

![](https://user-images.githubusercontent.com/3319450/42417093-750ed16c-8245-11e8-91f2-4746d2568ddb.png)

**Database**

![](https://user-images.githubusercontent.com/3319450/42417095-89907c12-8245-11e8-9d5e-0090c0c24527.png)

### Acceptable Time Formats

We have the ability to use time shorthands if need-be, the following are acceptable time inputs

| Input | Time Result |
| ----- | ----------- |
| 15s   | 15 seconds  |
| s15   | 15 seconds  |
| 60m   | 60 minutes  |
| 7d    | 7 days      |
| 1y    | 1 year      |
| 600   | 600 seconds |

### Perl Expiration

* To set an expiration time in Perl, very similarly to the LUA example above, you would simply call your `set_data` function with an expiration flag as your 3rd parameter like so

```perl
quest::set_data("my_example_flag", "some_value", 3600); # 3600 seconds = 1 hour (Expire in 1 hour)
```

## Benchmarks

* Below are some simple benchmarks used to calculate performance. While even these numbers could be greatly optimized yet, these are plenty good for most use cases that server operators need. If you need even faster temporary data storage within the context of a zone, I would suggest using \[\[Entity Variables]] as they operate purely in memory

![image](https://user-images.githubusercontent.com/3319450/42429025-83c1d260-82fc-11e8-804f-f7490ac23600.png)

```perl
sub EVENT_SAY {
    use Time::HiRes;
    my $start = [ Time::HiRes::gettimeofday() ];

    if ($text =~ /random-write/i) {
        my $iterations = 1000;
        my $key_range = 100;
        quest::debug("Testing random-write... Iterations: (" . plugin::commify($iterations) . ") Key Range: " . $key_range);
        for ($i = 0; $i < $iterations; $i++) {
            quest::set_data("key_" . int(rand($key_range)), &generate_random_string(100));
        }
    }

    if ($text =~ /sequential-write/i) {
        my $iterations = 1000;
        quest::debug("Testing sequential-write... Iterations: (" . plugin::commify($iterations) . ")");
        for ($i = 0; $i < $iterations; $i++) {
            quest::set_data("key_" . $i, &generate_random_string(100), 15);
        }
    }

    if ($text =~ /sequential-read/i) {
        my $iterations = 1000;
        quest::debug("Testing sequential-read... Iterations: (" . plugin::commify($iterations) . ")");
        for ($i = 0; $i < $iterations; $i++) {
            $data = quest::get_data("key_" . $i);
            # if ($data ne "") {
            #     quest::say("Data for $i : $data");
            # }
        }
    }

    if ($text =~ /random-read/i) {
        my $iterations = 1000;
        quest::debug("Testing random-read... Iterations: (" . plugin::commify($iterations) . ")");
        for ($i = 0; $i < $iterations; $i++) {
            $data = quest::get_data("key_" . int(rand($iterations)));
        }
    }

    my $elapsed = Time::HiRes::tv_interval($start);
    quest::debug("Operation took: " . $elapsed);
}
```

### Practical Example

In South Qeynos, you will find an NPC named Vicus Nonad.  Vicus is a tax collector, but because of his terrible cold, he needs a player to help make the rounds to collect taxes.  The early implementation of this quest script utilized quest globals, and below is an example of replacing the quest global functionality with Data Buckets.

{% tabs %}
{% tab title="Vicus_Nonad.pl" %}
```perl
sub EVENT_SPAWN {
	#:: Set a timer "cough" to repeat every 350 seconds (5 min 50 sec)
	quest::settimer("cough",350);
}

sub EVENT_TIMER {
	#:: Match the "cough" timer
	if ($timer eq "cough") {
		quest::emote("coughs and wheezes.");
	}
}

sub EVENT_SAY {
	if ($text=~/hail/i) {
		quest::say("Greetings, $name.  My name is Vicus Nonad. <cough>  I am the official tax collector for the fine city of Qeynos. <cough>  I serve the will of Antonius Bayle, our glorious leader.  <cough>  <cough>  Please excuse my [" . quest::saylink("cough") . "].  <cough>");
	}
	elsif ($text=~/cough/i) {
		quest::say("Oh, <cough> I am sorry, but it seems I have fallen a bit ill.  I was caught out in the rain the other day and my chills have gotten the best of me. <cough>  If only someone would [" . quest::saylink("help with today's collections") . "]..  <cough>");
	}
	elsif ($text=~/help with today's collections/i) {
		#:: Data bucket to verify quest has been started appropriately
		$key = $client->CharacterID() . "-tax-collection";
		#:: Set a data bucket, quest started
		quest::set_data($key, 1);
		quest::say("Oh thank <cough> you so <cough> <cough> much <cough>..  Here is the official collection box.  Please collect from each merchant on the <cough> [" . quest::saylink("list") . "].  Then bring me back the combined total of all your collections.");
		#:: Give a 17012 - Tax Collection Box
		quest::summonitem(17012);
	}
	elsif ($text=~/list/i) {
		quest::say("Oh.  <cough>  I am sorry..  I forgot to give it to you.  Here you go.  Be sure to give that back when your job is finished.  <cough>");
		#:: Give a 18009 - List of Debtors
		quest::summonitem(18009);
	}
}

sub EVENT_ITEM {
	#:: Match a 18009 - List of Debtors and 13181 - Full Tax Collection Box
	if (plugin::takeItems(13181 => 1, 18009 => 1)) {
		quest::say("<cough> Great! Thank you so much. Here is a small gratuity for a job well done. Thank you again. <cough> Antonius Bayle and the People of Qeynos appreciate all yo have done.");
		#:: Delete the data bucket
		$key = $client->CharacterID() . "-tax-collection";
		quest::delete_data($key);
		#:: Give a random reward: 13053 - Brass Ring, 10010 - Copper Amulet, 10018 - Hematite, 10017 - Turquoise
		quest::summonitem(quest::ChooseRandom(13053, 10010, 10018, 10017));
		#:: Ding!
		quest::ding();
		#:: Set factions
		quest::faction(219,10);		#:: + Antonius Bayle
		quest::faction(262,4);		#:: + Guards of Qeynos
		quest::faction(304,-5);		#:: - Ring of Scale
		quest::faction(273,-10);	#:: - Kane Bayle
		quest::faction(291,10);		#:: + Merchants of Qeynos
		#:: Grant a small amount of experience
		quest::exp(500);
		#:: Create a hash for storing cash - 200 to 300cp
		my %cash = plugin::RandomCash(200,300);
		#:: Grant a random cash reward
		quest::givecash($cash{copper},$cash{silver},$cash{gold},$cash{platinum});
	}
	#:: Match a 13181 - Full Tax Collection Box
	elsif (plugin::takeItems(13181 => 1)) {
		quest::say("Very good <cough> work. But I need both the full tax collection box and the list of debtors. You did get the [" . quest::saylink("list") . "] from me before you left, right? <cough>");
		#:: Return a 13181 - Full Tax Collection Box
		quest::summonitem(13181);
	}
	#:: Match a 18009 - List of Debtors
	elsif (plugin::takeItems(18009 => 1)) {
		quest::say("Very good <cough> work. But I need both the full tax collection box and the list of debtors. You did get the [" . quest::saylink("list") . "] from me before you left, right? <cough>");
	}
	#:: Return unused items
	plugin::returnUnusedItems();
}

```
{% endtab %}
{% endtabs %}

At line 22 above, we see the implementation of the Data Bucket Key.  This replaces the use of `quest::setglobal`.  In the Database, we see the corresponding key:

![](../../.gitbook/assets/\_mysql\_5\_6\_45\__fvproject_fvproject_data_buckets.png)

Note that we clean up the key upon handing in the requisite tax money and list with the following (line 41 above):

```perl
		$key = $client->CharacterID() . "-tax-collection";
		quest::delete_data($key);
```

The original portion of the script for the Quest Global would have been as follows:

```perl
#:: Match for "help", case insensitive
elsif ($text=~/help/i) {
			quest::say("Oh thank <cough> you so <cough> <cough> much <cough>..  Here is the official collection box.  Please collect from each merchant on the <cough> [list].  Then bring me back the combined total of all your collections.");
			#:: Give a 17012 - Tax Collection Box
			quest::summonitem(17012);
			#:: Set the qglobal "tax_collection", to a value of "0", for all NPCs and Zones, and last forever
			quest::setglobal("tax_collection", 0, 5, "F");
}
```

While this implementation may seem to be easier, it should be noted that on a server with many players, running a query through a hash of globals for each event that triggers a look up can cause serious performance issues.  Imagine a hash being created for every global qglobal entry (the "5" in the script above), for EVERY event trigger!

Our intrepid adventurer is required to do this quest in order, by speaking with Vicus prior to the merchants in Qeynos who have to pay taxes.  To enforce this, we verify that the user has the appropriate data bucket key set (line 9) before offering the dialogue that results in the tax payment:

{% tabs %}
{% tab title="Mar_Sedder.pl" %}
```perl
sub EVENT_SAY {
	if ($text=~/hail/i) {
		quest::say("Hail, $name.  What brings you to the Tin Soldier?  We have the finest in previously owned weapons and armor.  Feel free to browse.");
	}
	elsif ($text=~/tax collection/i) {
		#:: Data bucket to verify quest has been started appropriately
		$key = $client->CharacterID() . "-tax-collection";
		#:: Match if the key exists
		if (quest::get_data($key)) {
			quest::say("Here are the taxes, $name. Boy, you call the guards and they take their time to show up but be a few days late on your taxes and they send the goons after you. Sheesh!");
			#:: Give a 13171 - Sedder's Tax Payment
			quest::summonitem(13171);
			#:: Set faction
			quest::faction(291,-1);		#:: - Merchants of Qeynos
		}
	}
}

sub EVENT_ITEM {
	#:: Return unused items
	plugin::returnUnusedItems();
}
```
{% endtab %}
{% endtabs %}

