## EVENT_AGGRO

``` perl

sub EVENT_AGGRO {
}
```
## EVENT_AGGRO_SAY

``` perl

sub EVENT_AGGRO_SAY {
	quest::debug("data " . $data);
	quest::debug("text " . $text);
	quest::debug("langid " . $langid);
}
```
## EVENT_ATTACK

``` perl

sub EVENT_ATTACK {
}
```
## EVENT_CAST

``` perl

sub EVENT_CAST {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("caster_level " . $caster_level);
}
```
## EVENT_CAST_BEGIN

``` perl

sub EVENT_CAST_BEGIN {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("caster_level " . $caster_level);
}
```
## EVENT_CAST_ON

``` perl

sub EVENT_CAST_ON {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("caster_level " . $caster_level);
}
```
## EVENT_COMBAT

``` perl

sub EVENT_COMBAT {
	quest::debug("combat_state " . $combat_state);
}
```
## EVENT_DEATH

``` perl

sub EVENT_DEATH {
	quest::debug("killer_id " . $killer_id);
	quest::debug("killer_damage " . $killer_damage);
	quest::debug("killer_spell " . $killer_spell);
	quest::debug("killer_skill " . $killer_skill);
	quest::debug("killed_corpse_id " . $killed_corpse_id);
	quest::debug("killed_npc_id " . $killed_npc_id);
	quest::debug("killed_x " . $killed_x);
	quest::debug("killed_y " . $killed_y);
	quest::debug("killed_z " . $killed_z);
	quest::debug("killed_h " . $killed_h);
}
```
## EVENT_DEATH_COMPLETE

``` perl

sub EVENT_DEATH_COMPLETE {
	quest::debug("killer_id " . $killer_id);
	quest::debug("killer_damage " . $killer_damage);
	quest::debug("killer_spell " . $killer_spell);
	quest::debug("killer_skill " . $killer_skill);
	quest::debug("killed_corpse_id " . $killed_corpse_id);
	quest::debug("killed_npc_id " . $killed_npc_id);
	quest::debug("killed_x " . $killed_x);
	quest::debug("killed_y " . $killed_y);
	quest::debug("killed_z " . $killed_z);
	quest::debug("killed_h " . $killed_h);
}
```
## EVENT_ENTER

``` perl

sub EVENT_ENTER {
}
```
## EVENT_ENTER_AREA

``` perl

sub EVENT_ENTER_AREA {
	quest::debug("area_id " . $area_id);
	quest::debug("area_type " . $area_type);
}
```
## EVENT_EXIT

``` perl

sub EVENT_EXIT {
}
```
## EVENT_FEIGN_DEATH

``` perl

sub EVENT_FEIGN_DEATH {
}
```
## EVENT_HATE_LIST

``` perl

sub EVENT_HATE_LIST {
	quest::debug("hate_state " . $hate_state);
}
```
## EVENT_HP

``` perl

sub EVENT_HP {
	quest::debug("hpevent " . $hpevent);
	quest::debug("inchpevent " . $inchpevent);
}
```
## EVENT_ITEM

``` perl

sub EVENT_ITEM {
	quest::debug("QuestItem " . $QuestItem);
	quest::debug("copper " . $copper);
	quest::debug("silver " . $silver);
	quest::debug("gold " . $gold);
	quest::debug("platinum " . $platinum);
}
```
## EVENT_KILLED_MERIT

``` perl

sub EVENT_KILLED_MERIT {
}
```
## EVENT_LEAVE_AREA

``` perl

sub EVENT_LEAVE_AREA {
	quest::debug("area_id " . $area_id);
	quest::debug("area_type " . $area_type);
}
```
## EVENT_LOOT_ZONE

``` perl

sub EVENT_LOOT_ZONE {
	quest::debug("looted_id " . $looted_id);
	quest::debug("looted_charges " . $looted_charges);
	quest::debug("corpse " . $corpse);
	quest::debug("corpse_id " . $corpse_id);
}
```
## EVENT_NPC_SLAY

``` perl

sub EVENT_NPC_SLAY {
	quest::debug("killed " . $killed);
}
```
## EVENT_PAYLOAD

``` perl

sub EVENT_PAYLOAD {
	quest::debug("payload_id " . $payload_id);
	quest::debug("payload_value " . $payload_value);
}
```
## EVENT_POPUPRESPONSE

``` perl

sub EVENT_POPUPRESPONSE {
	quest::debug("popupid " . $popupid);
}
```
## EVENT_PROXIMITY_SAY

``` perl

sub EVENT_PROXIMITY_SAY {
	quest::debug("data " . $data);
	quest::debug("text " . $text);
	quest::debug("langid " . $langid);
}
```
## EVENT_SAY

``` perl

sub EVENT_SAY {
	quest::debug("data " . $data);
	quest::debug("text " . $text);
	quest::debug("langid " . $langid);
}
```
## EVENT_SIGNAL

``` perl

sub EVENT_SIGNAL {
	quest::debug("signal " . $signal);
}
```
## EVENT_SLAY

``` perl

sub EVENT_SLAY {
}
```
## EVENT_SPAWN

``` perl

sub EVENT_SPAWN {
}
```
## EVENT_TARGET_CHANGE

``` perl

sub EVENT_TARGET_CHANGE {
}
```
## EVENT_TASKACCEPTED

``` perl

sub EVENT_TASKACCEPTED {
	quest::debug("task_id " . $task_id);
}
```
## EVENT_TICK

``` perl

sub EVENT_TICK {
}
```
## EVENT_TIMER

``` perl

sub EVENT_TIMER {
	quest::debug("timer " . $timer);
}
```
## EVENT_WAYPOINT_ARRIVE

``` perl

sub EVENT_WAYPOINT_ARRIVE {
	quest::debug("wp " . $wp);
}
```
## EVENT_WAYPOINT_DEPART

``` perl

sub EVENT_WAYPOINT_DEPART {
	quest::debug("wp " . $wp);
}
```
