# Player Buffer Scripts

### Discipline Trainer

This script will allow an NPC to train all available disciplines up to the player's current level.

{% code title="Dave_the_Disciplinarian.pl" %}
```perl
sub EVENT_SAY {
	#:: Match text for "hail", case insensitive
	if ($text=~/hail/i) {
		#:: Send a message that only the client can see, in yellow (15) text
		$client->Message(15, "With just one look, I can see that you have forgotten the finer points of combat, $name.  Would you like me to [" . quest::saylink("teach") . "] you how to perform the skills that require more discipline than the basics?");
	}
	#:: Match text for "teach", case insensitive
	elsif ($text=~/teach/i) {
		#:: Clear out any existing disciplines (optional)
		quest::untraindiscs();
		#:: Train all disciplines up to the user's level
		quest::traindiscs($ulevel, 0);
		#:: Send a message that only the client can see, in yellow (15) text
		$client->Message(15, "You look like a more fierce combatant already! Go out and test your new abilities!");
		#:: Play a Ding! sound
		quest::ding();
	}
}

sub EVENT_ITEM {
	#:: Return unused items since we don't expect any handins
	plugin::returnUnusedItems();
}
```
{% endcode %}

If you would prefer that your players do not have to interact with an NPC, it is quite simple to add all new disciplines each time a player levels using a script.  You would simply add this snippet to your global quest script file.

This script will train all new disciplines when a player levels up:

{% tabs %}
{% tab title="Perl" %}
{% code title="global/global_player.pl" %}
```perl
sub EVENT_LEVEL_UP {
	#:: Train all disciplines, maximum set to player's level, minimum set to the level prior
	quest::traindiscs($ulevel,$ulevel - 1);
}
```
{% endcode %}
{% endtab %}

{% tab title="Lua" %}
```
function event_level_up(e)
			-- Train all disciplines up to current level
			eq.train_discs(e.self:GetLevel());
end
```
{% endtab %}
{% endtabs %}

### Spell Scriber

This script will scribe all available spells (or songs) up to the player's current level.

{% code title="Skippy_the_Scribe.pl" %}
```perl
sub EVENT_SAY {
	#:: Match text for "hail", case insensitive
	if ($text=~/hail/i) {
		#:: Separate response for melee classes--match classes using a string comparison (eq)
		if ($class eq 'Berserker' || $class eq 'Monk' || $class eq 'Rogue' || $class eq 'Warrior') {
			#:: Send a message that only the client can see, in yellow (15) text
			$client->Message(15, "Young fighter, I am the greatest spell scribe Norrath has ever seen--I do not waste my time on brutes like you!");
		}
		#:: Separate response for bards, who are always special
		elsif ($class eq 'Bard') {
			#:: Send a message that only the client can see, in yellow (15) text
			$client->Message(15, "With just one look, I can see that your songbook is lacking, $name.  Would you like me to [" . quest::saylink("scribe") . "] all of the known $class songs for you?");
		}
		#:: Separate response for casting classes
		elsif ($class eq 'Beastlord' || $class eq 'Cleric' || $class eq 'Druid' || $class eq 'Enchanter' || $class eq 'Magician' || $class eq 'Necromancer' || $class eq 'Paladin' || $class eq 'Ranger' || $class eq 'Shadowknight' || $class eq 'Shaman' || $class eq 'Wizard') {
			#:: Send a message that only the client can see, in yellow (15) text
			$client->Message(15, "With just one look, I can see that your spellbook is lacking, $name.  Would you like me to [" . quest::saylink("scribe") . "] all of the known $class spells for you?");
		}
	}
	#:: Match text for "scribe", case insensitive
	elsif ($text=~/scribe/i) {
		if ($class eq 'Bard' || $class eq 'Beastlord' || $class eq 'Cleric' || $class eq 'Druid' || $class eq 'Enchanter' || $class eq 'Magician' || $class eq 'Necromancer' || $class eq 'Paladin' || $class eq 'Ranger' || $class eq 'Shadowknight' || $class eq 'Shaman' || $class eq 'Wizard') {
			#:: Clear out any existing spells
			quest::unscribespells();
			#:: Scribe all spells up to the user's level
			quest::scribespells($ulevel, 0);
			#:: Send a message that only the client can see, in yellow (15) text
			$client->Message(15, "You look like a more powerful caster already! Go out and test your new spells!");
			#:: Play a Ding! sound
			quest::ding();
		}
		elsif ($class eq 'Berserker' || $class eq 'Monk' || $class eq 'Rogue' || $class eq 'Warrior') {
			$client->Message(15, "Begone, $class--lest I turn you into froglok tad!");
		}
	}
}

sub EVENT_ITEM {
	#:: Return unused items since we don't expect any handins
	plugin::returnUnusedItems();
}
```
{% endcode %}

As with disciplines, if you would prefer that your players do not have to interact with an NPC, it is quite simple to add all new spells/songs each time a player levels adding a snippet to your global quest file.  

This script will scribe all new spells/songs when a player levels up.

{% tabs %}
{% tab title="Perl" %}
{% code title="global/global_player.pl" %}
```perl
sub EVENT_LEVEL_UP {
	#:: Scribe all spells/songs, maximum set to player's level, minimum set to the level prior
	quest::scribespells($ulevel,$ulevel - 1);
}
```
{% endcode %}
{% endtab %}

{% tab title="Lua" %}
```
function event_level_up(e)
			-- Scribe all spells up to current level
			eq.scribe_spells(e.self:GetLevel());
end
```
{% endtab %}
{% endtabs %}

### Skill Maxer

This script will set all available skills to their maximum amount at the player's current level.

{% code title="Scotty_the_Skilled.pl" %}
```perl
sub EVENT_SAY {
	#:: Match text for "hail", case insensitive
	if ($text=~/hail/i) {
		#:: Send a message that only the client can see, in yellow (15) text
		$client->Message(15, "Hello, $name!  It appears you could use some help with your [" . quest::saylink("skills") . "].  Would you like me to teach you?");
	}
	#:: Match text for "skills", case insensitive
	elsif ($text=~/skills/i) {
		#:: Set available (non-trade, non-casting specialization) skills to maximum for race/class at current level
		foreach my $skill ( 0 .. 42, 48 .. 54, 70 .. 74 ) {
			next unless $client->CanHaveSkill($skill);
			#:: Create a scalar variable to store each skill's maximum skill level at the player's current level
			my $maxSkill = $client->MaxSkill($skill, $client->GetClass(), $ulevel);
			#:: Check that the player's skill does not already exceed the maximum skill based on level
			next unless $maxSkill > $client->GetRawSkill($skill);
			#:: Set the skill to the maximum
			$client->SetSkill($skill, $maxSkill);
		}
		#:: Send a message that only the client can see, in yellow (15) text
		$client->Message(15, "You look like a more capable $class already! Go out and test your new skills!");
		#:: Play a Ding! sound
		quest::ding();
	}
}

sub EVENT_ITEM {
	#:: Return unused items since we don't expect any handins
	plugin::returnUnusedItems();
}
```
{% endcode %}

As with spells and disciplines, if you would prefer that your players do not have to interact with an NPC, it is quite simple to max all skills each time a player levels by adding a snippet to your global quest file.

This script will max all skills when a player levels up.

{% tabs %}
{% tab title="Perl" %}
{% code title="global/global_player.pl" %}
```perl
sub EVENT_LEVEL_UP {
	#:: Set available (non-trade, non-casting specialization) skills to maximum for race/class at current level
	foreach my $skill ( 0 .. 42, 48 .. 54, 70 .. 74 ) {
		next unless $client->CanHaveSkill($skill);
		#:: Create a scalar variable to store each skill's maximum skill level at the player's current level
		my $maxSkill = $client->MaxSkill($skill, $client->GetClass(), $ulevel);
		#:: Check that the player's skill does not already exceed the maximum skill based on level
		next unless $maxSkill > $client->GetRawSkill($skill);
		#:: Set the skill to the maximum
		$client->SetSkill($skill, $maxSkill);
	}
}
```
{% endcode %}
{% endtab %}

{% tab title="Lua" %}
```
function event_level_up(e)
     local skills = { 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 48, 49, 50, 51, 52, 53, 54, 56, 62, 66, 67, 70, 71, 72, 73, 74 };
     for i, curskill in ipairs(skills) do
          local maxskill = e.self:MaxSkill(curskill);
          if (e.self:CanHaveSkill(curskill) == false) then
               --Do nothing
          elseif (maxskill <= e.self:GetRawSkill(curskill)) then
               --Do nothing
          else
               --Do Training
               e.self:SetSkill(curskill, maxskill);
          end
     end
end
```
{% endtab %}
{% endtabs %}

### Player Buffer

This script will provide buffs and cures to the player and/or the player's pet, or the player's group and their pets, and will bind the player to the area. The NPC will charge the player for the good buffs.

Credit to Maze who created the original version of this script, which has been adapted for use on this wiki.

{% code title="Canary_Morris.pl" %}
```perl
#:: Create a scalar vaiable to store the maximum level of buffs
$maxlevelbuffs = 70;
#:: Create a hash to define spells cast at each 10 levels, and their cost
#:: Keys are by each 10 levels, with differentiation for spells ranked as generic vs. advanced
#:: Values are each spell, by id, and the cost in platinum for each rank
%buffshash = (
	#:: Hash for Warriors
	1 => 
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10costadvanced	=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},
	#:: Hash for Clerics
	2 =>
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},
	#:: Hash for Paladins
	3 =>
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},
	#:: Hash for Ranger
	4 =>
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},
	#:: Hash for Shadow Knight
	5 =>
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},
	#:: Hash for Druid
	6 =>
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},
	#:: Hash for Monk
	7 =>
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},
	#:: Hash for Bard
	8 =>
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},
	#:: Hash for Rogue
	9 =>
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},
	#:: Hash for Shaman
	10 =>
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},
	#:: Hash for Necromancer
	11 =>
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},
	#:: Hash for Wizard
	12 =>
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},	
	#:: Hash for Magician
	13 =>
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},
	#:: Hash for Enchanter
	14 =>
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},
	#:: Hash for Beastlord
	15 => 
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	},
	#:: Hash for Berserker
	16 => 
	{
		level10generic		=>	[276,278,219,368,146,148,279,129],
		level10costgeneric	=>	0,
		level10advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level10costadvanced	=>	100,
		level20generic		=>	[276,278,219,368,146,148,279,129],
		level20costgeneric	=>	0,
		level20advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level20costadvanced	=>	100,
		level30generic		=>	[276,278,219,368,146,148,279,129],
		level30costgeneric	=>	0,
		level30advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level30costadvanced	=>	100,
		level40generic		=>	[276,278,219,368,146,148,279,129],
		level40costgeneric	=>	0,
		level40advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level40costadvanced	=>	100,
		level50generic		=>	[276,278,219,368,146,148,279,129],
		level50costgeneric	=>	0,
		level50advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level50costadvanced	=>	100,
		level60generic		=>	[276,278,219,368,146,148,279,129],
		level60costgeneric	=>	0,
		level60advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level60costadvanced	=>	100,
		level70generic		=>	[276,278,219,368,146,148,279,129],
		level70costgeneric	=>	0,
		level70advanced		=>	[3692,2505,423,356,172,64,145,1693,15,60,61,46,2524],
		level70costadvanced	=>	100,
	}
);


sub EVENT_SAY {
	#:: Create a scalar variable to store the cost of group buffs
	my $groupbuffmaxcost = 0;
	#:: Match text for "hail", case insensitive
	if ($text =~/Hail/i) { 
		quest::say("Greetings $name. I have been sent to aid you on your journey! Would you like some [".quest::saylink("lesser buffs",1)."] for ".$buffshash{$client->GetClass()}->{"level".round_up($ulevel)."costgeneric"}."pp, or some [".quest::saylink("greater blessings",1)."] for ".$buffshash{$client->GetClass()}->{"level".round_up($ulevel)."costadvanced"}."pp?");
		#:: Check to see if the player is grouped
		if ($client->IsGrouped()) {
			#:: Create a scalar variable to store the group
			$buffgroup = $entity_list->GetGroupByClient($client);
			#:: Create a loop to count how many members are in the group
			for ($count = 0; $count < $buffgroup->GroupCount(); $count++) {
				#:: Create a scalar variable to store the group member count
				my $groupmember = $buffgroup->GetMember($count);
				#:: Make sure that each group member is in zone before charging for them
				if ($groupmember->GetZoneID() == $zoneid) {
					#:: Create a scalar variable to store the cost of the buffs
					$groupbuffmaxcost += ($buffshash{$groupmember->GetClass()}->{"level".round_up($groupmember->GetLevel())."costadvanced"});
				}
			}
			quest::say("Since you are grouped, you can also get the best [".quest::saylink("group blessings",1)."] for ".$groupbuffmaxcost."pp");
		}
		quest::say("To [".quest::saylink ("cure", 1)."] you and/or your pet is free.");
		quest::say("Or I can [".quest::saylink ("bind", 1)."] you to this location at no cost.");
	}
	#:: Match text for "lesser buffs", case insensitive
	elsif ($text=~/lesser buffs/i) {
		#:: Make sure the player has enough plat for the buffs
		if ($client->GetCarriedMoney() >= (($buffshash{$client->GetClass()}->{"level".round_up($ulevel)."costgeneric"}) * 1000)) {
			#:: Create an array of the buffs to be cast on the player
			my @buffstocast = @{$buffshash{$client->GetClass()}->{"level".round_up($ulevel)."generic"}};
			#:: Create a loop to cast each buff in the hash for the pet (if it exists) and then the player
			foreach my $buff (@buffstocast) {
				#:: Check to see if the player has a pet
				if ($client->GetPetID()) {
					#:: Create a scalar variable to store the pet by ID
					my $pcpet = $entity_list->GetMobByID($client->GetPetID());
					#:: Buff the pet
					$client->SpellFinished ($buff, $pcpet);
				}
				#:: Buff the player
				quest::selfcast ($buff);
			}
			#:: Take payment for the buffs
			$client->TakeMoneyFromPP((($buffshash{$client->GetClass()}->{"level".round_up($ulevel)."costgeneric"}) * 1000), 1) unless ($buffshash{$client->GetClass()}->{"level".round_up($ulevel)."costgeneric"}) == 0;
		}
		else {
			#:: Match if the player does not have enough plat
			quest::say("I charge ".($buffshash{$client->GetClass()}->{"level".round_up($ulevel)."costgeneric"})."pp for lesser buffs, you don't appear to have enough on you.");
		}
	}
	#:: Match text for "greater blessings", case insensitive
	elsif ($text =~/greater blessings/i) {
		#:: Make sure the player has enough plat for the buffs
		if ($client->GetCarriedMoney() >= (($buffshash{$client->GetClass()}->{"level".round_up($ulevel)."costadvanced"}) * 1000)) {
			#:: Create an array of the buffs to be cast on the player
			my @buffstocast = @{$buffshash{$client->GetClass()}->{"level".round_up($ulevel)."advanced"}};
			#:: Create a loop to cast each buff in the hash for the pet (if it exists) and then the player
			foreach my $buff (@buffstocast) {
				#:: Check to see if the player has a pet
				if ($client->GetPetID()) {
					#:: Create a scalar variable to store the pet by ID
					my $pcpet = $entity_list->GetMobByID($client->GetPetID());
					#:: Buff the pet
					$client->SpellFinished ($buff, $pcpet);
				}
				#:: Buff the player
				quest::selfcast ($buff);
			}
			#:: Take payment for the buffs
			$client->TakeMoneyFromPP((($buffshash{$client->GetClass()}->{"level".round_up($ulevel)."costadvanced"}) * 1000), 1) unless ($buffshash{$client->GetClass()}->{"level".round_up($ulevel)."costadvanced"}) == 0;
		}
		else {
			#:: Match if the player does not have enough plat
			quest::say("I charge ".($buffshash{$client->GetClass()}->{"level".round_up($ulevel)."costadvanced"})."pp for greater blessings, you don't appear to have enough on you.");
		}
	}
	#:: Match text for "group blessings", case insensitive, and if the player is grouped
	elsif ($text =~ /group blessings/i && $client->IsGrouped()) {
		#:: Make sure the player has enough plat for the buffs
		if ($client->GetCarriedMoney() >= (($buffshash{$client->GetClass()}->{"level".$maxlevelbuffs."costadvanced"}) * 1000 * $buffgroup->GroupCount())) {
			#:: Create a loop to count the group members
			for ($count = 0; $count < $buffgroup->GroupCount(); $count++) {
				#:: Create a scalar variable to store the group member count
				my $groupmember = $buffgroup->GetMember($count);
				#:: Make sure that each group member is in zone before charging for them
				if ($groupmember->GetZoneID() == $zoneid) {
					#:: Take payment for the buffs
					$client->TakeMoneyFromPP((($buffshash{$groupmember->GetClass()}->{"level".round_up($groupmember->GetLevel())."costadvanced"}) * 1000), 1) unless ($buffshash{$groupmember->GetClass()}->{"level".round_up($groupmember->GetLevel())."costadvanced"}) == 0;
					#:: Create an array of the buffs to be cast on each player
					my @buffstocast = @{$buffshash{$groupmember->GetClass()}->{"level".round_up($groupmember->GetLevel())."advanced"}};
					#:: Create a loop to cast each buff in the hash for each group member pet (if any exist) and then each group member
					foreach my $buff (@buffstocast) {
						#:: Check to see if the player has a pet
						if ($groupmember->GetPetID()) {
							#:: Create a scalar variable to store the pet by ID
							my $groupmemberpet = $entity_list->GetMobByID($groupmember->GetPetID());
							#:: Buff the pet
							$groupmember->SpellFinished ($buff, $groupmemberpet);
						}
						#:: Buff the group members
						$groupmember->SpellFinished ($buff, $groupmember);
					}
				#:: Create an on-screen message for each group member so they know who to thank for the buffs
				$groupmember->SendMarqueeMessage(15, 510, 1, 1, 3000, "You've been buffed compliments of ".$client->GetCleanName()."!");
				}			
			}
		}
		else {
			#:: Match if the player does not have enough plat
			quest::say("I charge a maximum of ".(($buffshash{$client->GetClass()}->{"level".$maxlevelbuffs."costadvanced"}) * $buffgroup->GroupCount())."pp to cast buffs on your entire group.  You don't appear to have enough!");
		}	
	}
	#:: Match text for "cure", case insensitive
	elsif ($text=~/cure/i) { 
		#:: Check to see if the player has a pet
		if ($client->GetPetID()) {
			#:: Create a scalar variable to store the pet by ID
			my $pppet = $entity_list->GetMobByID($client->GetPetID());
			#:: Cure the pet
			$client->SpellFinished (6594, $pppet);
		}
		#:: Cure the player
		$client->SpellFinished(6594); 
	}
	#:: Match text for "bind", case insensitive
	elsif ($text=~/bind/i) {
		#:: Bind the player
		quest::selfcast(35);
	}
}

#:: This is a subroutine to round up to the nearest 10 levels
sub round_up {
	my $n = shift;
	my $scale = 10**int(log($n)/log(10));
	$n = 9 if $scale == 1;
	if ($n > $scale) {
		$n = int($n/$scale+1)*$scale;
	}
	$n;
}
```
{% endcode %}
