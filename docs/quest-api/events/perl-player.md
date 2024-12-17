!!! info end

    Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl){:target="perl_event"} for latest definitions and Quest examples

## EVENT_AA_BUY

``` perl

sub EVENT_AA_BUY {
	quest::debug("aa_cost " . $aa_cost);
	quest::debug("aa_id " . $aa_id);
	quest::debug("aa_previous_id " . $aa_previous_id);
	quest::debug("aa_next_id " . $aa_next_id);
}
```
## EVENT_AA_EXP_GAIN

``` perl

sub EVENT_AA_EXP_GAIN {
	quest::debug("aa_exp_gained " . $aa_exp_gained);
}
```
## EVENT_AA_GAIN

``` perl

sub EVENT_AA_GAIN {
	quest::debug("aa_gained " . $aa_gained);
}
```
## EVENT_AA_LOSS

``` perl

sub EVENT_AA_LOSS {
	quest::debug("aa_lost " . $aa_lost);
}
```
## EVENT_ALT_CURRENCY_LOSS

``` perl

sub EVENT_ALT_CURRENCY_LOSS {
	quest::debug("currency_id " . $currency_id);
	quest::debug("amount " . $amount);
	quest::debug("total " . $total);
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
## EVENT_AUGMENT_INSERT_CLIENT

``` perl

sub EVENT_AUGMENT_INSERT_CLIENT {
	quest::debug("item_id " . $item_id);
	quest::debug("item_slot " . $item_slot);
	quest::debug("augment_id " . $augment_id);
	quest::debug("augment_slot " . $augment_slot);
	quest::debug("item " . $item);
	quest::debug("augment " . $augment);
}
```
## EVENT_AUGMENT_REMOVE_CLIENT

``` perl

sub EVENT_AUGMENT_REMOVE_CLIENT {
	quest::debug("item_id " . $item_id);
	quest::debug("item_slot " . $item_slot);
	quest::debug("augment_id " . $augment_id);
	quest::debug("augment_slot " . $augment_slot);
	quest::debug("destroyed " . $destroyed);
	quest::debug("item " . $item);
	quest::debug("augment " . $augment);
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
## EVENT_BOT_CREATE

``` perl

sub EVENT_BOT_CREATE {
	quest::debug("bot_name " . $bot_name);
	quest::debug("bot_id " . $bot_id);
	quest::debug("bot_race " . $bot_race);
	quest::debug("bot_class " . $bot_class);
	quest::debug("bot_gender " . $bot_gender);
}
```
## EVENT_CAST

``` perl

sub EVENT_CAST {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("caster_level " . $caster_level);
	quest::debug("target_id " . $target_id);
	quest::debug("target " . $target);
	quest::debug("spell " . $spell);
}
```
## EVENT_CAST_BEGIN

``` perl

sub EVENT_CAST_BEGIN {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("caster_level " . $caster_level);
	quest::debug("target_id " . $target_id);
	quest::debug("target " . $target);
	quest::debug("spell " . $spell);
}
```
## EVENT_CAST_ON

``` perl

sub EVENT_CAST_ON {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("caster_level " . $caster_level);
	quest::debug("target_id " . $target_id);
	quest::debug("target " . $target);
	quest::debug("spell " . $spell);
}
```
## EVENT_CLICKDOOR

``` perl

sub EVENT_CLICKDOOR {
	quest::debug("doorid " . $doorid);
	quest::debug("version " . $version);
	quest::debug("door " . $door);
}
```
## EVENT_CLICK_OBJECT

``` perl

sub EVENT_CLICK_OBJECT {
	quest::debug("objectid " . $objectid);
	quest::debug("clicker_id " . $clicker_id);
	quest::debug("object " . $object);
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
	quest::debug("target " . $target);
}
```
## EVENT_CONSIDER_CORPSE

``` perl

sub EVENT_CONSIDER_CORPSE {
	quest::debug("corpse_entity_id " . $corpse_entity_id);
	quest::debug("corpse " . $corpse);
}
```
## EVENT_CRYSTAL_GAIN

``` perl

sub EVENT_CRYSTAL_GAIN {
	quest::debug("ebon_amount " . $ebon_amount);
	quest::debug("radiant_amount " . $radiant_amount);
	quest::debug("is_reclaim " . $is_reclaim);
}
```
## EVENT_CRYSTAL_LOSS

``` perl

sub EVENT_CRYSTAL_LOSS {
	quest::debug("ebon_amount " . $ebon_amount);
	quest::debug("radiant_amount " . $radiant_amount);
	quest::debug("is_reclaim " . $is_reclaim);
}
```
## EVENT_DAMAGE_GIVEN

``` perl

sub EVENT_DAMAGE_GIVEN {
	quest::debug("entity_id " . $entity_id);
	quest::debug("damage " . $damage);
	quest::debug("spell_id " . $spell_id);
	quest::debug("skill_id " . $skill_id);
	quest::debug("is_damage_shield " . $is_damage_shield);
	quest::debug("is_avoidable " . $is_avoidable);
	quest::debug("buff_slot " . $buff_slot);
	quest::debug("is_buff_tic " . $is_buff_tic);
	quest::debug("special_attack " . $special_attack);
	quest::debug("spell " . $spell);
}
```
## EVENT_DAMAGE_TAKEN

``` perl

sub EVENT_DAMAGE_TAKEN {
	quest::debug("entity_id " . $entity_id);
	quest::debug("damage " . $damage);
	quest::debug("spell_id " . $spell_id);
	quest::debug("skill_id " . $skill_id);
	quest::debug("is_damage_shield " . $is_damage_shield);
	quest::debug("is_avoidable " . $is_avoidable);
	quest::debug("buff_slot " . $buff_slot);
	quest::debug("is_buff_tic " . $is_buff_tic);
	quest::debug("special_attack " . $special_attack);
	quest::debug("spell " . $spell);
}
```
## EVENT_DEATH

``` perl

sub EVENT_DEATH {
	quest::debug("killer_id " . $killer_id);
	quest::debug("killer_damage " . $killer_damage);
	quest::debug("killer_spell " . $killer_spell);
	quest::debug("killer_skill " . $killer_skill);
	quest::debug("killed_entity_id " . $killed_entity_id);
	quest::debug("combat_start_time " . $combat_start_time);
	quest::debug("combat_end_time " . $combat_end_time);
	quest::debug("damage_received " . $damage_received);
	quest::debug("healing_received " . $healing_received);
	quest::debug("killed_corpse_id " . $killed_corpse_id);
	quest::debug("killed_x " . $killed_x);
	quest::debug("killed_y " . $killed_y);
	quest::debug("killed_z " . $killed_z);
	quest::debug("killed_h " . $killed_h);
	quest::debug("killed_merc_id " . $killed_merc_id);
	quest::debug("killed_npc_id " . $killed_npc_id);
}
```
## EVENT_DEATH_COMPLETE

``` perl

sub EVENT_DEATH_COMPLETE {
	quest::debug("killer_id " . $killer_id);
	quest::debug("killer_damage " . $killer_damage);
	quest::debug("killer_spell " . $killer_spell);
	quest::debug("killer_skill " . $killer_skill);
	quest::debug("killed_entity_id " . $killed_entity_id);
	quest::debug("combat_start_time " . $combat_start_time);
	quest::debug("combat_end_time " . $combat_end_time);
	quest::debug("damage_received " . $damage_received);
	quest::debug("healing_received " . $healing_received);
	quest::debug("killed_corpse_id " . $killed_corpse_id);
	quest::debug("killed_x " . $killed_x);
	quest::debug("killed_y " . $killed_y);
	quest::debug("killed_z " . $killed_z);
	quest::debug("killed_h " . $killed_h);
	quest::debug("killed_merc_id " . $killed_merc_id);
	quest::debug("killed_npc_id " . $killed_npc_id);
}
```
## EVENT_DESTROY_ITEM_CLIENT

``` perl

sub EVENT_DESTROY_ITEM_CLIENT {
	quest::debug("item_id " . $item_id);
	quest::debug("item_name " . $item_name);
	quest::debug("quantity " . $quantity);
	quest::debug("item " . $item);
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
## EVENT_DROP_ITEM_CLIENT

``` perl

sub EVENT_DROP_ITEM_CLIENT {
	quest::debug("quantity " . $quantity);
	quest::debug("item_name " . $item_name);
	quest::debug("item_id " . $item_id);
	quest::debug("spell_id " . $spell_id);
	quest::debug("slot_id " . $slot_id);
	quest::debug("item " . $item);
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
## EVENT_ENTER_ZONE

``` perl

sub EVENT_ENTER_ZONE {
}
```
## EVENT_ENTITY_VARIABLE_DELETE

``` perl

sub EVENT_ENTITY_VARIABLE_DELETE {
	quest::debug("variable_name " . $variable_name);
	quest::debug("variable_value " . $variable_value);
}
```
## EVENT_ENTITY_VARIABLE_SET

``` perl

sub EVENT_ENTITY_VARIABLE_SET {
	quest::debug("variable_name " . $variable_name);
	quest::debug("variable_value " . $variable_value);
}
```
## EVENT_ENTITY_VARIABLE_UPDATE

``` perl

sub EVENT_ENTITY_VARIABLE_UPDATE {
	quest::debug("variable_name " . $variable_name);
	quest::debug("old_value " . $old_value);
	quest::debug("new_value " . $new_value);
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
	quest::debug("item " . $item);
}
```
## EVENT_EXIT

``` perl

sub EVENT_EXIT {
}
```
## EVENT_EXP_GAIN

``` perl

sub EVENT_EXP_GAIN {
	quest::debug("exp_gained " . $exp_gained);
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
## EVENT_GM_COMMAND

``` perl

sub EVENT_GM_COMMAND {
	quest::debug("message " . $message);
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
	quest::debug("target " . $target);
}
```
## EVENT_ITEM_CLICK_CAST_CLIENT

``` perl

sub EVENT_ITEM_CLICK_CAST_CLIENT {
	quest::debug("slot_id " . $slot_id);
	quest::debug("item_id " . $item_id);
	quest::debug("item_name " . $item_name);
	quest::debug("spell_id " . $spell_id);
	quest::debug("item " . $item);
}
```
## EVENT_ITEM_CLICK_CLIENT

``` perl

sub EVENT_ITEM_CLICK_CLIENT {
	quest::debug("slot_id " . $slot_id);
	quest::debug("item_id " . $item_id);
	quest::debug("item_name " . $item_name);
	quest::debug("spell_id " . $spell_id);
	quest::debug("item " . $item);
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
## EVENT_LEAVE_AREA

``` perl

sub EVENT_LEAVE_AREA {
	quest::debug("area_id " . $area_id);
	quest::debug("area_type " . $area_type);
}
```
## EVENT_LEVEL_DOWN

``` perl

sub EVENT_LEVEL_DOWN {
	quest::debug("levels_lost " . $levels_lost);
}
```
## EVENT_LEVEL_UP

``` perl

sub EVENT_LEVEL_UP {
	quest::debug("levels_gained " . $levels_gained);
}
```
## EVENT_LOOT

``` perl

sub EVENT_LOOT {
	quest::debug("looted_id " . $looted_id);
	quest::debug("looted_charges " . $looted_charges);
	quest::debug("corpse_name " . $corpse_name);
	quest::debug("corpse_id " . $corpse_id);
	quest::debug("corpse " . $corpse);
}
```
## EVENT_MEMORIZE_SPELL

``` perl

sub EVENT_MEMORIZE_SPELL {
	quest::debug("slot_id " . $slot_id);
	quest::debug("spell_id " . $spell_id);
	quest::debug("spell " . $spell);
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
## EVENT_PAYLOAD

``` perl

sub EVENT_PAYLOAD {
	quest::debug("payload_id " . $payload_id);
	quest::debug("payload_value " . $payload_value);
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
## EVENT_READ_ITEM

``` perl

sub EVENT_READ_ITEM {
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
## EVENT_SCRIBE_SPELL

``` perl

sub EVENT_SCRIBE_SPELL {
	quest::debug("slot_id " . $slot_id);
	quest::debug("spell_id " . $spell_id);
	quest::debug("spell " . $spell);
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
## EVENT_SPELL_BLOCKED

``` perl

sub EVENT_SPELL_BLOCKED {
	quest::debug("blocking_spell_id " . $blocking_spell_id);
	quest::debug("cast_spell_id " . $cast_spell_id);
	quest::debug("blocking_spell " . $blocking_spell);
	quest::debug("cast_spell " . $cast_spell);
}
```
## EVENT_TARGET_CHANGE

``` perl

sub EVENT_TARGET_CHANGE {
	quest::debug("target " . $target);
}
```
## EVENT_TASKACCEPTED

``` perl

sub EVENT_TASKACCEPTED {
	quest::debug("task_id " . $task_id);
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
## EVENT_TIMER_PAUSE

``` perl

sub EVENT_TIMER_PAUSE {
	quest::debug("timer " . $timer);
	quest::debug("duration " . $duration);
}
```
## EVENT_TIMER_RESUME

``` perl

sub EVENT_TIMER_RESUME {
	quest::debug("timer " . $timer);
	quest::debug("duration " . $duration);
}
```
## EVENT_TIMER_START

``` perl

sub EVENT_TIMER_START {
	quest::debug("timer " . $timer);
	quest::debug("duration " . $duration);
}
```
## EVENT_TIMER_STOP

``` perl

sub EVENT_TIMER_STOP {
	quest::debug("timer " . $timer);
}
```
## EVENT_UNEQUIP_ITEM_CLIENT

``` perl

sub EVENT_UNEQUIP_ITEM_CLIENT {
	quest::debug("item_id " . $item_id);
	quest::debug("item_quantity " . $item_quantity);
	quest::debug("slot_id " . $slot_id);
	quest::debug("item " . $item);
}
```
## EVENT_UNHANDLED_OPCODE

``` perl

sub EVENT_UNHANDLED_OPCODE {
}
```
## EVENT_UNMEMORIZE_SPELL

``` perl

sub EVENT_UNMEMORIZE_SPELL {
	quest::debug("slot_id " . $slot_id);
	quest::debug("spell_id " . $spell_id);
	quest::debug("spell " . $spell);
}
```
## EVENT_UNSCRIBE_SPELL

``` perl

sub EVENT_UNSCRIBE_SPELL {
	quest::debug("slot_id " . $slot_id);
	quest::debug("spell_id " . $spell_id);
	quest::debug("spell " . $spell);
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
