# Entity Variables

## Description

I would hope you read [Entity Lists](entity-lists.md) and understand entities and how they work before reading this. If you have, then it will help understanding what entity variables are and how powerful they can be in helping you do whatever it is you are trying to do

When you are using Perl, and in a script you define an object, you define it like so, **$object = 1; **Easy, right?** **- Yes, but there is a problem. **ANY NPC** who also has this same script loaded could overwrite this variable. When I am trying to set and read data unique to a particular NPC, using this method does not make things all that easy. Given, you COULD use a Perl array accessor unique to the NPCID but that is not quite as flexible to this system that has been developed by KLS to do.

So here I am going to show you some examples and why **entity variables are amazing.**

An entity variable can be used by most commonly $npc/$client. It is just like any other variable to store quickly accessible stuff.

They are stored in **memory of the zone, and persist when an entity is alive**. Meaning if you repop the variables clear, or an NPC/client dies/zones etc.

Picture it like this, these NPCs below have a '**virtual bucket**' of variables you can set. The only thing you need when setting variables is a **variable name to read and write.**

With this, you can get really really creative as to how you pass variables around.

* **$entity->SetEntityVariable("variable_name", "value") - Write**
* **$entity->GetEntityVariable("variable_name") - Read - Would return whatever the variable is set to for that NPC**
* **There are Perl Plugin Shorthands for these functions (For speed of writing scripts)**
  * **plugin::SEV($entity, "variable_name", "value"); - Write**
  * **plugin::REV($entity, "variable_name"); - Returns value, Read**
  * See [Plugins Reference](https://eqemu.gitbook.io/quest-api/perl/plugins)

## Example

* In this simple example, we're looping the same timer every 3 seconds and an NPC is going to say something different every 3 seconds until it gets to the 4th passing of the timer and the NPC shuts his mouth.

{% tabs %}
{% tab title="Perl" %}
```perl
sub EVENT_SAY {
    if($text=~/hail/i) {
        quest::say("hi, starting timer...");
        quest::settimer("test_timer", 3);
        plugin::SEV($client, "passings", 1); #:: Just called it 'passings' because I will be using it to track passings of a timer
    }
}
 
sub EVENT_TIMER {
    if($timer eq "test_timer") {
        my $passings = plugin::REV($client, "passings");
        if($passings == 1) {
			quest::say("Hey how are you?");
		} elsif($passings == 2) {
			quest::say("What do you want?");
		} elsif($passings == 3) {
			quest::say("Are you going to go away?");
		} elsif($passings == 4) { #:: Since we want to stop at 4, we cut off the timer and stop counting the passes
            quest::say("Like seriously, buzz off, I'm busy...");
            quest::stoptimer($timer);
        }
        plugin::SEV($client, "passings", ($passings + 1)); #:: Increment the number so we don't loop the timer forever...
    }
}
```
{% endtab %}
{% endtabs %}
