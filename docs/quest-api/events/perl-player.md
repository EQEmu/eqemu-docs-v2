## EVENT_AA_BUY

``` perl

sub EVENT_AA_BUY {
	quest::debug("aa_cost " . $aa_cost);
	quest::debug("aa_id " . $aa_id);
	quest::debug("aa_previous_id " . $aa_previous_id);
	quest::debug("aa_next_id " . $aa_next_id);
}
```
## EVENT_AA_GAIN

``` perl

sub EVENT_AA_GAIN {
	quest::debug("aa_gained " . $aa_gained);
}
```
## EVENT_ALT_CURRENCY_MERCHANT_BUY

``` perl

sub EVENT_ALT_CURRENCY_MERCHANT_BUY {
	quest::debug("currency_id " . $currency_id);
	quest::debug("npc_id " . $npc_id);
	quest::debug("merchant_id " . $merchant_id);
	quest::debug("item_id " . $item_id);
	quest::debug("item_cost " . $item_cost);
}
```
## EVENT_ALT_CURRENCY_MERCHANT_SELL

``` perl

sub EVENT_ALT_CURRENCY_MERCHANT_SELL {
	quest::debug("currency_id " . $currency_id);
	quest::debug("npc_id " . $npc_id);
	quest::debug("merchant_id " . $merchant_id);
	quest::debug("item_id " . $item_id);
	quest::debug("item_cost " . $item_cost);
}
```
## EVENT_BOT_COMMAND

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
## EVENT_CLICKDOOR

``` perl

sub EVENT_CLICKDOOR {
	quest::debug("doorid " . $doorid);
	quest::debug("version " . $version);
}
```
## EVENT_CLICK_OBJECT

``` perl

sub EVENT_CLICK_OBJECT {
	quest::debug("objectid " . $objectid);
	quest::debug("clicker_id " . $clicker_id);
}
```
## EVENT_COMBINE

``` perl

sub EVENT_COMBINE {
	quest::debug("container_slot " . $container_slot);
}
```
## EVENT_COMBINE_FAILURE

``` perl

sub EVENT_COMBINE_FAILURE {
	quest::debug("recipe_id " . $recipe_id);
	quest::debug("recipe_name " . $recipe_name);
}
```
## EVENT_COMBINE_SUCCESS

``` perl

sub EVENT_COMBINE_SUCCESS {
	quest::debug("recipe_id " . $recipe_id);
	quest::debug("recipe_name " . $recipe_name);
}
```
## EVENT_COMBINE_VALIDATE

``` perl

sub EVENT_COMBINE_VALIDATE {
	quest::debug("recipe_id " . $recipe_id);
	quest::debug("validate_type " . $validate_type);
	quest::debug("zone_id " . $zone_id);
	quest::debug("tradeskill_id " . $tradeskill_id);
}
```
## EVENT_COMMAND

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

``` perl

sub EVENT_CONNECT {
}
```
## EVENT_CONSIDER

``` perl

sub EVENT_CONSIDER {
	quest::debug("entity_id " . $entity_id);
}
```
## EVENT_CONSIDER_CORPSE

``` perl

sub EVENT_CONSIDER_CORPSE {
	quest::debug("corpse_entity_id " . $corpse_entity_id);
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
## EVENT_DISCONNECT

``` perl

sub EVENT_DISCONNECT {
}
```
## EVENT_DISCOVER_ITEM

``` perl

sub EVENT_DISCOVER_ITEM {
	quest::debug("itemid " . $itemid);
}
```
## EVENT_DUEL_LOSE

``` perl

sub EVENT_DUEL_LOSE {
}
```
## EVENT_DUEL_WIN

``` perl

sub EVENT_DUEL_WIN {
}
```
## EVENT_ENTER_ZONE

``` perl

sub EVENT_ENTER_ZONE {
}
```
## EVENT_ENVIRONMENTAL_DAMAGE

``` perl

sub EVENT_ENVIRONMENTAL_DAMAGE {
	quest::debug("env_damage " . $env_damage);
	quest::debug("env_damage_type " . $env_damage_type);
	quest::debug("env_final_damage " . $env_final_damage);
}
```
## EVENT_EQUIP_ITEM_CLIENT

``` perl

sub EVENT_EQUIP_ITEM_CLIENT {
	quest::debug("item_id " . $item_id);
	quest::debug("item_quantity " . $item_quantity);
	quest::debug("slot_id " . $slot_id);
}
```
## EVENT_FEIGN_DEATH

``` perl

sub EVENT_FEIGN_DEATH {
}
```
## EVENT_FISH_FAILURE

``` perl

sub EVENT_FISH_FAILURE {
}
```
## EVENT_FISH_START

``` perl

sub EVENT_FISH_START {
}
```
## EVENT_FISH_SUCCESS

``` perl

sub EVENT_FISH_SUCCESS {
	quest::debug("fished_item " . $fished_item);
}
```
## EVENT_FORAGE_FAILURE

``` perl

sub EVENT_FORAGE_FAILURE {
}
```
## EVENT_FORAGE_SUCCESS

``` perl

sub EVENT_FORAGE_SUCCESS {
	quest::debug("foraged_item " . $foraged_item);
}
```
## EVENT_GROUP_CHANGE

``` perl

sub EVENT_GROUP_CHANGE {
	quest::debug("grouped " . $grouped);
	quest::debug("raided " . $raided);
}
```
## EVENT_INSPECT

``` perl

sub EVENT_INSPECT {
	quest::debug("target_id " . $target_id);
}
```
## EVENT_LANGUAGE_SKILL_UP

``` perl

sub EVENT_LANGUAGE_SKILL_UP {
	quest::debug("skill_id " . $skill_id);
	quest::debug("skill_value " . $skill_value);
	quest::debug("skill_max " . $skill_max);
}
```
## EVENT_LEVEL_UP

``` perl

sub EVENT_LEVEL_UP {
}
```
## EVENT_LOOT

``` perl

sub EVENT_LOOT {
	quest::debug("looted_id " . $looted_id);
	quest::debug("looted_charges " . $looted_charges);
	quest::debug("corpse " . $corpse);
	quest::debug("corpse_id " . $corpse_id);
}
```
## EVENT_MERCHANT_BUY

``` perl

sub EVENT_MERCHANT_BUY {
	quest::debug("npc_id " . $npc_id);
	quest::debug("merchant_id " . $merchant_id);
	quest::debug("item_id " . $item_id);
	quest::debug("item_quantity " . $item_quantity);
	quest::debug("item_cost " . $item_cost);
}
```
## EVENT_MERCHANT_SELL

``` perl

sub EVENT_MERCHANT_SELL {
	quest::debug("npc_id " . $npc_id);
	quest::debug("merchant_id " . $merchant_id);
	quest::debug("item_id " . $item_id);
	quest::debug("item_quantity " . $item_quantity);
	quest::debug("item_cost " . $item_cost);
}
```
## EVENT_PLAYER_PICKUP

``` perl

sub EVENT_PLAYER_PICKUP {
	quest::debug("picked_up_id " . $picked_up_id);
	quest::debug("picked_up_entity_id " . $picked_up_entity_id);
}
```
## EVENT_POPUPRESPONSE

``` perl

sub EVENT_POPUPRESPONSE {
	quest::debug("popupid " . $popupid);
}
```
## EVENT_RESPAWN

``` perl

sub EVENT_RESPAWN {
	quest::debug("option " . $option);
	quest::debug("resurrect " . $resurrect);
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
## EVENT_SKILL_UP

``` perl

sub EVENT_SKILL_UP {
	quest::debug("skill_id " . $skill_id);
	quest::debug("skill_value " . $skill_value);
	quest::debug("skill_max " . $skill_max);
	quest::debug("is_tradeskill " . $is_tradeskill);
}
```
## EVENT_SPELL_EFFECT_BOT

``` perl

sub EVENT_SPELL_EFFECT_BOT {
	quest::debug("payload_id " . $payload_id);
	quest::debug("payload_value " . $payload_value);
}
```
## EVENT_TARGET_CHANGE

``` perl

sub EVENT_TARGET_CHANGE {
}
```
## EVENT_TASK_BEFORE_UPDATE

``` perl

sub EVENT_TASK_BEFORE_UPDATE {
	quest::debug("donecount " . $donecount);
	quest::debug("activity_id " . $activity_id);
	quest::debug("task_id " . $task_id);
}
```
## EVENT_TASK_COMPLETE

``` perl

sub EVENT_TASK_COMPLETE {
	quest::debug("donecount " . $donecount);
	quest::debug("activity_id " . $activity_id);
	quest::debug("task_id " . $task_id);
}
```
## EVENT_TASK_FAIL

``` perl

sub EVENT_TASK_FAIL {
	quest::debug("task_id " . $task_id);
}
```
## EVENT_TASK_STAGE_COMPLETE

``` perl

sub EVENT_TASK_STAGE_COMPLETE {
	quest::debug("task_id " . $task_id);
	quest::debug("activity_id " . $activity_id);
}
```
## EVENT_TASK_UPDATE

``` perl

sub EVENT_TASK_UPDATE {
	quest::debug("donecount " . $donecount);
	quest::debug("activity_id " . $activity_id);
	quest::debug("task_id " . $task_id);
}
```
## EVENT_TEST_BUFF

``` perl

sub EVENT_TEST_BUFF {
}
```
## EVENT_TIMER

``` perl

sub EVENT_TIMER {
	quest::debug("timer " . $timer);
}
```
## EVENT_UNEQUIP_ITEM_CLIENT

``` perl

sub EVENT_UNEQUIP_ITEM_CLIENT {
	quest::debug("item_id " . $item_id);
	quest::debug("item_quantity " . $item_quantity);
	quest::debug("slot_id " . $slot_id);
}
```
## EVENT_UNHANDLED_OPCODE

``` perl

sub EVENT_UNHANDLED_OPCODE {
}
```
## EVENT_USE_SKILL

``` perl

sub EVENT_USE_SKILL {
	quest::debug("skill_id " . $skill_id);
	quest::debug("skill_level " . $skill_level);
}
```
## EVENT_WARP

``` perl

sub EVENT_WARP {
	quest::debug("from_x " . $from_x);
	quest::debug("from_y " . $from_y);
	quest::debug("from_z " . $from_z);
}
```
## EVENT_ZONE

``` perl

sub EVENT_ZONE {
	quest::debug("from_zone_id " . $from_zone_id);
	quest::debug("from_instance_id " . $from_instance_id);
	quest::debug("from_instance_version " . $from_instance_version);
	quest::debug("target_zone_id " . $target_zone_id);
	quest::debug("target_instance_id " . $target_instance_id);
	quest::debug("target_instance_version " . $target_instance_version);
}
```
