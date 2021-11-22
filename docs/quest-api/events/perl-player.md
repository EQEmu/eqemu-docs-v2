## EVENT_BOT_COMMAND

=== "EVENT_BOT_COMMAND"

``` perl

sub EVENT_BOT_COMMAND {
	quest::debug("bot_command " . $bot_command);
	quest::debug("args " . $args);
	quest::debug("data " . $data);
	quest::debug("text " . $text);
	quest::debug("langid " . $langid);
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
## EVENT_CLICKDOOR

=== "EVENT_CLICKDOOR"

``` perl

sub EVENT_CLICKDOOR {
	quest::debug("doorid " . $doorid);
	quest::debug("version " . $version);
}
```
## EVENT_CLICK_DOOR

=== "EVENT_CLICK_DOOR"

``` perl

sub EVENT_CLICK_DOOR {
}
```
## EVENT_CLICK_OBJECT

=== "EVENT_CLICK_OBJECT"

``` perl

sub EVENT_CLICK_OBJECT {
	quest::debug("objectid " . $objectid);
	quest::debug("clicker_id " . $clicker_id);
}
```
## EVENT_COMBINE

=== "EVENT_COMBINE"

``` perl

sub EVENT_COMBINE {
	quest::debug("container_slot " . $container_slot);
}
```
## EVENT_COMBINE_FAILURE

=== "EVENT_COMBINE_FAILURE"

``` perl

sub EVENT_COMBINE_FAILURE {
	quest::debug("recipe_id " . $recipe_id);
	quest::debug("recipe_name " . $recipe_name);
}
```
## EVENT_COMBINE_SUCCESS

=== "EVENT_COMBINE_SUCCESS"

``` perl

sub EVENT_COMBINE_SUCCESS {
	quest::debug("recipe_id " . $recipe_id);
	quest::debug("recipe_name " . $recipe_name);
}
```
## EVENT_COMBINE_VALIDATE

=== "EVENT_COMBINE_VALIDATE"

``` perl

sub EVENT_COMBINE_VALIDATE {
	quest::debug("recipe_id " . $recipe_id);
	quest::debug("validate_type " . $validate_type);
	quest::debug("zone_id " . $zone_id);
	quest::debug("tradeskill_id " . $tradeskill_id);
}
```
## EVENT_COMMAND

=== "EVENT_COMMAND"

``` perl

sub EVENT_COMMAND {
	quest::debug("command " . $command);
	quest::debug("args " . $args);
	quest::debug("data " . $data);
	quest::debug("text " . $text);
	quest::debug("langid " . $langid);
}
```
## EVENT_CONNECT

=== "EVENT_CONNECT"

``` perl

sub EVENT_CONNECT {
}
```
## EVENT_CONSIDER

=== "EVENT_CONSIDER"

``` perl

sub EVENT_CONSIDER {
	quest::debug("entity_id " . $entity_id);
}
```
## EVENT_CONSIDER_CORPSE

=== "EVENT_CONSIDER_CORPSE"

``` perl

sub EVENT_CONSIDER_CORPSE {
	quest::debug("corpse_entity_id " . $corpse_entity_id);
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
## EVENT_DISCONNECT

=== "EVENT_DISCONNECT"

``` perl

sub EVENT_DISCONNECT {
}
```
## EVENT_DISCOVER_ITEM

=== "EVENT_DISCOVER_ITEM"

``` perl

sub EVENT_DISCOVER_ITEM {
	quest::debug("itemid " . $itemid);
}
```
## EVENT_DUEL_LOSE

=== "EVENT_DUEL_LOSE"

``` perl

sub EVENT_DUEL_LOSE {
}
```
## EVENT_DUEL_WIN

=== "EVENT_DUEL_WIN"

``` perl

sub EVENT_DUEL_WIN {
}
```
## EVENT_ENTER_ZONE

=== "EVENT_ENTER_ZONE"

``` perl

sub EVENT_ENTER_ZONE {
}
```
## EVENT_ENVIRONMENTAL_DAMAGE

=== "EVENT_ENVIRONMENTAL_DAMAGE"

``` perl

sub EVENT_ENVIRONMENTAL_DAMAGE {
	quest::debug("env_damage " . $env_damage);
	quest::debug("env_damage_type " . $env_damage_type);
	quest::debug("env_final_damage " . $env_final_damage);
}
```
## EVENT_FEIGN_DEATH

=== "EVENT_FEIGN_DEATH"

``` perl

sub EVENT_FEIGN_DEATH {
}
```
## EVENT_FISH_FAILURE

=== "EVENT_FISH_FAILURE"

``` perl

sub EVENT_FISH_FAILURE {
}
```
## EVENT_FISH_START

=== "EVENT_FISH_START"

``` perl

sub EVENT_FISH_START {
}
```
## EVENT_FISH_SUCCESS

=== "EVENT_FISH_SUCCESS"

``` perl

sub EVENT_FISH_SUCCESS {
	quest::debug("fished_item " . $fished_item);
}
```
## EVENT_FORAGE_FAILURE

=== "EVENT_FORAGE_FAILURE"

``` perl

sub EVENT_FORAGE_FAILURE {
}
```
## EVENT_FORAGE_SUCCESS

=== "EVENT_FORAGE_SUCCESS"

``` perl

sub EVENT_FORAGE_SUCCESS {
	quest::debug("foraged_item " . $foraged_item);
}
```
## EVENT_GROUP_CHANGE

=== "EVENT_GROUP_CHANGE"

``` perl

sub EVENT_GROUP_CHANGE {
	quest::debug("grouped " . $grouped);
	quest::debug("raided " . $raided);
}
```
## EVENT_LEVEL_UP

=== "EVENT_LEVEL_UP"

``` perl

sub EVENT_LEVEL_UP {
}
```
## EVENT_LOOT

=== "EVENT_LOOT"

``` perl

sub EVENT_LOOT {
	quest::debug("looted_id " . $looted_id);
	quest::debug("looted_charges " . $looted_charges);
	quest::debug("corpse " . $corpse);
	quest::debug("corpse_id " . $corpse_id);
}
```
## EVENT_PLAYER_PICKUP

=== "EVENT_PLAYER_PICKUP"

``` perl

sub EVENT_PLAYER_PICKUP {
	quest::debug("picked_up_id " . $picked_up_id);
	quest::debug("picked_up_entity_id " . $picked_up_entity_id);
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
## EVENT_RESPAWN

=== "EVENT_RESPAWN"

``` perl

sub EVENT_RESPAWN {
	quest::debug("option " . $option);
	quest::debug("resurrect " . $resurrect);
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
## EVENT_TARGET_CHANGE

=== "EVENT_TARGET_CHANGE"

``` perl

sub EVENT_TARGET_CHANGE {
}
```
## EVENT_TASK_COMPLETE

=== "EVENT_TASK_COMPLETE"

``` perl

sub EVENT_TASK_COMPLETE {
	quest::debug("donecount " . $donecount);
	quest::debug("activity_id " . $activity_id);
	quest::debug("task_id " . $task_id);
}
```
## EVENT_TASK_FAIL

=== "EVENT_TASK_FAIL"

``` perl

sub EVENT_TASK_FAIL {
	quest::debug("task_id " . $task_id);
}
```
## EVENT_TASK_STAGE_COMPLETE

=== "EVENT_TASK_STAGE_COMPLETE"

``` perl

sub EVENT_TASK_STAGE_COMPLETE {
	quest::debug("task_id " . $task_id);
	quest::debug("activity_id " . $activity_id);
}
```
## EVENT_TASK_UPDATE

=== "EVENT_TASK_UPDATE"

``` perl

sub EVENT_TASK_UPDATE {
	quest::debug("donecount " . $donecount);
	quest::debug("activity_id " . $activity_id);
	quest::debug("task_id " . $task_id);
}
```
## EVENT_TEST_BUFF

=== "EVENT_TEST_BUFF"

``` perl

sub EVENT_TEST_BUFF {
}
```
## EVENT_TIMER

=== "EVENT_TIMER"

``` perl

sub EVENT_TIMER {
	quest::debug("timer " . $timer);
}
```
## EVENT_UNHANDLED_OPCODE

=== "EVENT_UNHANDLED_OPCODE"

``` perl

sub EVENT_UNHANDLED_OPCODE {
}
```
## EVENT_USE_SKILL

=== "EVENT_USE_SKILL"

``` perl

sub EVENT_USE_SKILL {
	quest::debug("skill_id " . $skill_id);
	quest::debug("skill_level " . $skill_level);
}
```
## EVENT_WARP

=== "EVENT_WARP"

``` perl

sub EVENT_WARP {
	quest::debug("from_x " . $from_x);
	quest::debug("from_y " . $from_y);
	quest::debug("from_z " . $from_z);
}
```
## EVENT_ZONE

=== "EVENT_ZONE"

``` perl

sub EVENT_ZONE {
	quest::debug("target_zone_id " . $target_zone_id);
}
```
