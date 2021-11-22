## EVENT_AGGRO

=== "EVENT_AGGRO"

``` perl

sub EVENT_AGGRO {
}
```
## EVENT_AGGRO_SAY

=== "EVENT_AGGRO_SAY"

``` perl

sub EVENT_AGGRO_SAY {
	quest::debug("data " . $data);
	quest::debug("text " . $text);
	quest::debug("langid " . $langid);
}
```
## EVENT_ATTACK

=== "EVENT_ATTACK"

``` perl

sub EVENT_ATTACK {
}
```
## EVENT_CAST

=== "EVENT_CAST"

``` perl

sub EVENT_CAST {
	quest::debug("spell_id " . $spell_id);
}
```
## EVENT_CAST_BEGIN

=== "EVENT_CAST_BEGIN"

``` perl

sub EVENT_CAST_BEGIN {
	quest::debug("spell_id " . $spell_id);
}
```
## EVENT_CAST_ON

=== "EVENT_CAST_ON"

``` perl

sub EVENT_CAST_ON {
	quest::debug("spell_id " . $spell_id);
}
```
## EVENT_COMBAT

=== "EVENT_COMBAT"

``` perl

sub EVENT_COMBAT {
	quest::debug("combat_state " . $combat_state);
}
```
## EVENT_DEATH

=== "EVENT_DEATH"

``` perl

sub EVENT_DEATH {
	quest::debug("killer_id " . $killer_id);
	quest::debug("killer_damage " . $killer_damage);
	quest::debug("killer_spell " . $killer_spell);
	quest::debug("killer_skill " . $killer_skill);
}
```
## EVENT_DEATH_COMPLETE

=== "EVENT_DEATH_COMPLETE"

``` perl

sub EVENT_DEATH_COMPLETE {
	quest::debug("killer_id " . $killer_id);
	quest::debug("killer_damage " . $killer_damage);
	quest::debug("killer_spell " . $killer_spell);
	quest::debug("killer_skill " . $killer_skill);
}
```
## EVENT_DEATH_ZONE

=== "EVENT_DEATH_ZONE"

``` perl

sub EVENT_DEATH_ZONE {
	quest::debug("killer_id " . $killer_id);
	quest::debug("killer_damage " . $killer_damage);
	quest::debug("killer_spell " . $killer_spell);
	quest::debug("killer_skill " . $killer_skill);
	quest::debug("killed_npc_id " . $killed_npc_id);
}
```
## EVENT_FEIGN_DEATH

=== "EVENT_FEIGN_DEATH"

``` perl

sub EVENT_FEIGN_DEATH {
}
```
## EVENT_HATE_LIST

=== "EVENT_HATE_LIST"

``` perl

sub EVENT_HATE_LIST {
	quest::debug("hate_state " . $hate_state);
}
```
## EVENT_HP

=== "EVENT_HP"

``` perl

sub EVENT_HP {
	quest::debug("hpevent " . $hpevent);
	quest::debug("inchpevent " . $inchpevent);
	quest::debug("hpevent " . $hpevent);
	quest::debug("inchpevent " . $inchpevent);
}
```
## EVENT_ITEM

=== "EVENT_ITEM"

``` perl

sub EVENT_ITEM {
	quest::debug("copper " . $copper);
	quest::debug("silver " . $silver);
	quest::debug("gold " . $gold);
	quest::debug("platinum " . $platinum);
}
```
## EVENT_KILLED_MERIT

=== "EVENT_KILLED_MERIT"

``` perl

sub EVENT_KILLED_MERIT {
}
```
## EVENT_LOOT_ZONE

=== "EVENT_LOOT_ZONE"

``` perl

sub EVENT_LOOT_ZONE {
	quest::debug("looted_id " . $looted_id);
	quest::debug("looted_charges " . $looted_charges);
	quest::debug("corpse " . $corpse);
	quest::debug("corpse_id " . $corpse_id);
}
```
## EVENT_NPC_SLAY

=== "EVENT_NPC_SLAY"

``` perl

sub EVENT_NPC_SLAY {
	quest::debug("killed " . $killed);
}
```
## EVENT_POPUPRESPONSE

=== "EVENT_POPUPRESPONSE"

``` perl

sub EVENT_POPUPRESPONSE {
	quest::debug("popupid " . $popupid);
}
```
## EVENT_POPUP_RESPONSE

=== "EVENT_POPUP_RESPONSE"

``` perl

sub EVENT_POPUP_RESPONSE {
}
```
## EVENT_PROXIMITY_SAY

=== "EVENT_PROXIMITY_SAY"

``` perl

sub EVENT_PROXIMITY_SAY {
	quest::debug("data " . $data);
	quest::debug("text " . $text);
	quest::debug("langid " . $langid);
}
```
## EVENT_SAY

=== "EVENT_SAY"

``` perl

sub EVENT_SAY {
	quest::debug("data " . $data);
	quest::debug("text " . $text);
	quest::debug("langid " . $langid);
}
```
## EVENT_SIGNAL

=== "EVENT_SIGNAL"

``` perl

sub EVENT_SIGNAL {
	quest::debug("signal " . $signal);
}
```
## EVENT_SLAY

=== "EVENT_SLAY"

``` perl

sub EVENT_SLAY {
}
```
## EVENT_SPAWN

=== "EVENT_SPAWN"

``` perl

sub EVENT_SPAWN {
}
```
## EVENT_SPAWN_ZONE

=== "EVENT_SPAWN_ZONE"

``` perl

sub EVENT_SPAWN_ZONE {
	quest::debug("spawned_entity_id " . $spawned_entity_id);
	quest::debug("spawned_npc_id " . $spawned_npc_id);
}
```
## EVENT_TARGET_CHANGE

=== "EVENT_TARGET_CHANGE"

``` perl

sub EVENT_TARGET_CHANGE {
}
```
## EVENT_TASKACCEPTED

=== "EVENT_TASKACCEPTED"

``` perl

sub EVENT_TASKACCEPTED {
	quest::debug("task_id " . $task_id);
}
```
## EVENT_TASK_ACCEPTED

=== "EVENT_TASK_ACCEPTED"

``` perl

sub EVENT_TASK_ACCEPTED {
}
```
## EVENT_TICK

=== "EVENT_TICK"

``` perl

sub EVENT_TICK {
}
```
## EVENT_TIMER

=== "EVENT_TIMER"

``` perl

sub EVENT_TIMER {
	quest::debug("timer " . $timer);
}
```
## EVENT_TRADE

=== "EVENT_TRADE"

``` perl

sub EVENT_TRADE {
}
```
## EVENT_WAYPOINT_ARRIVE

=== "EVENT_WAYPOINT_ARRIVE"

``` perl

sub EVENT_WAYPOINT_ARRIVE {
	quest::debug("wp " . $wp);
}
```
## EVENT_WAYPOINT_DEPART

=== "EVENT_WAYPOINT_DEPART"

``` perl

sub EVENT_WAYPOINT_DEPART {
	quest::debug("wp " . $wp);
}
```
