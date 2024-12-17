!!! info end

    Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua){:target="perl_event"} for latest definitions and Quest examples

## EVENT_AA_BUY

``` lua

function EVENT_AA_BUY(e) {
	eq.debug("aa_cost " .. e.aa_cost);
	eq.debug("aa_id " .. e.aa_id);
	eq.debug("aa_previous_id " .. e.aa_previous_id);
	eq.debug("aa_next_id " .. e.aa_next_id);
}
```
## EVENT_AA_EXP_GAIN

``` lua

function EVENT_AA_EXP_GAIN(e) {
	eq.debug("aa_exp_gained " .. e.aa_exp_gained);
}
```
## EVENT_AA_GAIN

``` lua

function EVENT_AA_GAIN(e) {
	eq.debug("aa_gained " .. e.aa_gained);
}
```
## EVENT_AA_LOSS

``` lua

function EVENT_AA_LOSS(e) {
	eq.debug("aa_lost " .. e.aa_lost);
}
```
## EVENT_ALT_CURRENCY_GAIN

``` lua

function EVENT_ALT_CURRENCY_GAIN(e) {
	eq.debug("currency_id " .. e.currency_id);
	eq.debug("amount " .. e.amount);
	eq.debug("total " .. e.total);
}
```
## EVENT_ALT_CURRENCY_LOSS

``` lua

function EVENT_ALT_CURRENCY_LOSS(e) {
	eq.debug("currency_id " .. e.currency_id);
	eq.debug("amount " .. e.amount);
	eq.debug("total " .. e.total);
}
```
## EVENT_ALT_CURRENCY_MERCHANT_BUY

``` lua

function EVENT_ALT_CURRENCY_MERCHANT_BUY(e) {
	eq.debug("currency_id " .. e.currency_id);
	eq.debug("npc_id " .. e.npc_id);
	eq.debug("merchant_id " .. e.merchant_id);
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_cost " .. e.item_cost);
}
```
## EVENT_ALT_CURRENCY_MERCHANT_SELL

``` lua

function EVENT_ALT_CURRENCY_MERCHANT_SELL(e) {
	eq.debug("currency_id " .. e.currency_id);
	eq.debug("npc_id " .. e.npc_id);
	eq.debug("merchant_id " .. e.merchant_id);
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_cost " .. e.item_cost);
}
```
## EVENT_AUGMENT_INSERT_CLIENT

``` lua

function EVENT_AUGMENT_INSERT_CLIENT(e) {
	eq.debug("item " .. e.item);
	eq.debug("augment " .. e.augment);
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_slot " .. e.item_slot);
	eq.debug("augment_id " .. e.augment_id);
	eq.debug("augment_slot " .. e.augment_slot);
}
```
## EVENT_AUGMENT_REMOVE_CLIENT

``` lua

function EVENT_AUGMENT_REMOVE_CLIENT(e) {
	eq.debug("item " .. e.item);
	eq.debug("augment " .. e.augment);
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_slot " .. e.item_slot);
	eq.debug("augment_id " .. e.augment_id);
	eq.debug("augment_slot " .. e.augment_slot);
	eq.debug("destroyed " .. e.destroyed);
}
```
## EVENT_BOT_COMMAND

``` lua

function EVENT_BOT_COMMAND(e) {
	eq.debug("bot_command " .. e.bot_command);
	eq.debug("args " .. e.args);
}
```
## EVENT_BOT_CREATE

``` lua

function EVENT_BOT_CREATE(e) {
	eq.debug("bot_name " .. e.bot_name);
	eq.debug("bot_id " .. e.bot_id);
	eq.debug("bot_race " .. e.bot_race);
	eq.debug("bot_class " .. e.bot_class);
	eq.debug("bot_gender " .. e.bot_gender);
}
```
## EVENT_CAST

``` lua

function EVENT_CAST(e) {
	eq.debug("spell " .. e.spell);
	eq.debug("caster_id " .. e.caster_id);
	eq.debug("caster_level " .. e.caster_level);
	eq.debug("target_id " .. e.target_id);
	eq.debug("target " .. e.target);
}
```
## EVENT_CAST_BEGIN

``` lua

function EVENT_CAST_BEGIN(e) {
	eq.debug("spell " .. e.spell);
	eq.debug("caster_id " .. e.caster_id);
	eq.debug("caster_level " .. e.caster_level);
	eq.debug("target_id " .. e.target_id);
	eq.debug("target " .. e.target);
}
```
## EVENT_CAST_ON

``` lua

function EVENT_CAST_ON(e) {
	eq.debug("spell " .. e.spell);
	eq.debug("caster_id " .. e.caster_id);
	eq.debug("caster_level " .. e.caster_level);
	eq.debug("target_id " .. e.target_id);
	eq.debug("target " .. e.target);
}
```
## EVENT_CLICK_DOOR

``` lua

function EVENT_CLICK_DOOR(e) {
	eq.debug("door " .. e.door);
}
```
## EVENT_CLICK_OBJECT

``` lua

function EVENT_CLICK_OBJECT(e) {
	eq.debug("object " .. e.object);
}
```
## EVENT_COMBINE

``` lua

function EVENT_COMBINE(e) {
	eq.debug("container_slot " .. e.container_slot);
}
```
## EVENT_COMBINE_FAILURE

``` lua

function EVENT_COMBINE_FAILURE(e) {
	eq.debug("recipe_id " .. e.recipe_id);
	eq.debug("recipe_name " .. e.recipe_name);
}
```
## EVENT_COMBINE_SUCCESS

``` lua

function EVENT_COMBINE_SUCCESS(e) {
	eq.debug("recipe_id " .. e.recipe_id);
	eq.debug("recipe_name " .. e.recipe_name);
}
```
## EVENT_COMBINE_VALIDATE

``` lua

function EVENT_COMBINE_VALIDATE(e) {
	eq.debug("recipe_id " .. e.recipe_id);
	eq.debug("validate_type " .. e.validate_type);
	eq.debug("zone_id " .. e.zone_id);
	eq.debug("tradeskill_id " .. e.tradeskill_id);
}
```
## EVENT_COMMAND

``` lua

function EVENT_COMMAND(e) {
	eq.debug("command " .. e.command);
	eq.debug("args " .. e.args);
}
```
## EVENT_CONNECT

``` lua

function EVENT_CONNECT(e) {
}
```
## EVENT_CONSIDER

``` lua

function EVENT_CONSIDER(e) {
	eq.debug("entity_id " .. e.entity_id);
	eq.debug("other " .. e.other);
}
```
## EVENT_CONSIDER_CORPSE

``` lua

function EVENT_CONSIDER_CORPSE(e) {
	eq.debug("corpse_entity_id " .. e.corpse_entity_id);
	eq.debug("corpse " .. e.corpse);
}
```
## EVENT_CRYSTAL_GAIN

``` lua

function EVENT_CRYSTAL_GAIN(e) {
	eq.debug("ebon_amount " .. e.ebon_amount);
	eq.debug("radiant_amount " .. e.radiant_amount);
	eq.debug("is_reclaim " .. e.is_reclaim);
}
```
## EVENT_CRYSTAL_LOSS

``` lua

function EVENT_CRYSTAL_LOSS(e) {
	eq.debug("ebon_amount " .. e.ebon_amount);
	eq.debug("radiant_amount " .. e.radiant_amount);
	eq.debug("is_reclaim " .. e.is_reclaim);
}
```
## EVENT_DAMAGE_GIVEN

``` lua

function EVENT_DAMAGE_GIVEN(e) {
	eq.debug("entity_id " .. e.entity_id);
	eq.debug("damage " .. e.damage);
	eq.debug("spell_id " .. e.spell_id);
	eq.debug("skill_id " .. e.skill_id);
	eq.debug("is_damage_shield " .. e.is_damage_shield);
	eq.debug("is_avoidable " .. e.is_avoidable);
	eq.debug("buff_slot " .. e.buff_slot);
	eq.debug("is_buff_tic " .. e.is_buff_tic);
	eq.debug("special_attack " .. e.special_attack);
	eq.debug("other " .. e.other);
}
```
## EVENT_DAMAGE_TAKEN

``` lua

function EVENT_DAMAGE_TAKEN(e) {
	eq.debug("entity_id " .. e.entity_id);
	eq.debug("damage " .. e.damage);
	eq.debug("spell_id " .. e.spell_id);
	eq.debug("skill_id " .. e.skill_id);
	eq.debug("is_damage_shield " .. e.is_damage_shield);
	eq.debug("is_avoidable " .. e.is_avoidable);
	eq.debug("buff_slot " .. e.buff_slot);
	eq.debug("is_buff_tic " .. e.is_buff_tic);
	eq.debug("special_attack " .. e.special_attack);
	eq.debug("other " .. e.other);
}
```
## EVENT_DEATH

``` lua

function EVENT_DEATH(e) {
	eq.debug("other " .. e.other);
	eq.debug("killer_id " .. e.killer_id);
	eq.debug("damage " .. e.damage);
	eq.debug("spell " .. e.spell);
	eq.debug("skill " .. e.skill);
	eq.debug("killed_entity_id " .. e.killed_entity_id);
	eq.debug("combat_start_time " .. e.combat_start_time);
	eq.debug("combat_end_time " .. e.combat_end_time);
	eq.debug("damage_received " .. e.damage_received);
	eq.debug("healing_received " .. e.healing_received);
}
```
## EVENT_DEATH_COMPLETE

``` lua

function EVENT_DEATH_COMPLETE(e) {
	eq.debug("other " .. e.other);
	eq.debug("killer_id " .. e.killer_id);
	eq.debug("damage " .. e.damage);
	eq.debug("spell " .. e.spell);
	eq.debug("skill " .. e.skill);
	eq.debug("killed_entity_id " .. e.killed_entity_id);
	eq.debug("combat_start_time " .. e.combat_start_time);
	eq.debug("combat_end_time " .. e.combat_end_time);
	eq.debug("damage_received " .. e.damage_received);
	eq.debug("healing_received " .. e.healing_received);
}
```
## EVENT_DESTROY_ITEM_CLIENT

``` lua

function EVENT_DESTROY_ITEM_CLIENT(e) {
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_name " .. e.item_name);
	eq.debug("quantity " .. e.quantity);
	eq.debug("item " .. e.item);
}
```
## EVENT_DISCONNECT

``` lua

function EVENT_DISCONNECT(e) {
}
```
## EVENT_DISCOVER_ITEM

``` lua

function EVENT_DISCOVER_ITEM(e) {
	eq.debug("item " .. e.item);
}
```
## EVENT_DROP_ITEM_CLIENT

``` lua

function EVENT_DROP_ITEM_CLIENT(e) {
	eq.debug("slot_id " .. e.slot_id);
	eq.debug("quantity " .. e.quantity);
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_name " .. e.item_name);
	eq.debug("spell_id " .. e.spell_id);
	eq.debug("item " .. e.item);
}
```
## EVENT_DUEL_LOSE

``` lua

function EVENT_DUEL_LOSE(e) {
	eq.debug("other " .. e.other);
}
```
## EVENT_DUEL_WIN

``` lua

function EVENT_DUEL_WIN(e) {
	eq.debug("other " .. e.other);
}
```
## EVENT_ENTER

``` lua

function EVENT_ENTER(e) {
}
```
## EVENT_ENTER_AREA

``` lua

function EVENT_ENTER_AREA(e) {
	eq.debug("area_id " .. e.area_id);
	eq.debug("area_type " .. e.area_type);
}
```
## EVENT_ENTER_ZONE

``` lua

function EVENT_ENTER_ZONE(e) {
}
```
## EVENT_ENTITY_VARIABLE_DELETE

``` lua

function EVENT_ENTITY_VARIABLE_DELETE(e) {
	eq.debug("variable_name " .. e.variable_name);
	eq.debug("variable_value " .. e.variable_value);
	eq.debug("old_value " .. e.old_value);
	eq.debug("new_value " .. e.new_value);
}
```
## EVENT_ENTITY_VARIABLE_SET

``` lua

function EVENT_ENTITY_VARIABLE_SET(e) {
	eq.debug("variable_name " .. e.variable_name);
	eq.debug("variable_value " .. e.variable_value);
	eq.debug("old_value " .. e.old_value);
	eq.debug("new_value " .. e.new_value);
}
```
## EVENT_ENTITY_VARIABLE_UPDATE

``` lua

function EVENT_ENTITY_VARIABLE_UPDATE(e) {
	eq.debug("variable_name " .. e.variable_name);
	eq.debug("variable_value " .. e.variable_value);
	eq.debug("old_value " .. e.old_value);
	eq.debug("new_value " .. e.new_value);
}
```
## EVENT_ENVIRONMENTAL_DAMAGE

``` lua

function EVENT_ENVIRONMENTAL_DAMAGE(e) {
	eq.debug("env_damage " .. e.env_damage);
	eq.debug("env_damage_type " .. e.env_damage_type);
	eq.debug("env_final_damage " .. e.env_final_damage);
}
```
## EVENT_EQUIP_ITEM_CLIENT

``` lua

function EVENT_EQUIP_ITEM_CLIENT(e) {
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_quantity " .. e.item_quantity);
	eq.debug("slot_id " .. e.slot_id);
	eq.debug("item " .. e.item);
}
```
## EVENT_EXIT

``` lua

function EVENT_EXIT(e) {
}
```
## EVENT_EXP_GAIN

``` lua

function EVENT_EXP_GAIN(e) {
	eq.debug("exp_gained " .. e.exp_gained);
}
```
## EVENT_FEIGN_DEATH

``` lua

function EVENT_FEIGN_DEATH(e) {
	eq.debug("other " .. e.other);
}
```
## EVENT_FISH_FAILURE

``` lua

function EVENT_FISH_FAILURE(e) {
}
```
## EVENT_FISH_START

``` lua

function EVENT_FISH_START(e) {
}
```
## EVENT_FISH_SUCCESS

``` lua

function EVENT_FISH_SUCCESS(e) {
	eq.debug("item " .. e.item);
}
```
## EVENT_FORAGE_FAILURE

``` lua

function EVENT_FORAGE_FAILURE(e) {
}
```
## EVENT_FORAGE_SUCCESS

``` lua

function EVENT_FORAGE_SUCCESS(e) {
	eq.debug("item " .. e.item);
}
```
## EVENT_GM_COMMAND

``` lua

function EVENT_GM_COMMAND(e) {
	eq.debug("message " .. e.message);
}
```
## EVENT_GROUP_CHANGE

``` lua

function EVENT_GROUP_CHANGE(e) {
}
```
## EVENT_INSPECT

``` lua

function EVENT_INSPECT(e) {
	eq.debug("other " .. e.other);
}
```
## EVENT_ITEM_CLICK_CAST_CLIENT

``` lua

function EVENT_ITEM_CLICK_CAST_CLIENT(e) {
	eq.debug("slot_id " .. e.slot_id);
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_name " .. e.item_name);
	eq.debug("spell_id " .. e.spell_id);
	eq.debug("item " .. e.item);
}
```
## EVENT_ITEM_CLICK_CLIENT

``` lua

function EVENT_ITEM_CLICK_CLIENT(e) {
	eq.debug("slot_id " .. e.slot_id);
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_name " .. e.item_name);
	eq.debug("spell_id " .. e.spell_id);
	eq.debug("item " .. e.item);
}
```
## EVENT_LANGUAGE_SKILL_UP

``` lua

function EVENT_LANGUAGE_SKILL_UP(e) {
	eq.debug("skill_id " .. e.skill_id);
	eq.debug("skill_value " .. e.skill_value);
	eq.debug("skill_max " .. e.skill_max);
}
```
## EVENT_LDON_POINTS_GAIN

``` lua

function EVENT_LDON_POINTS_GAIN(e) {
	eq.debug("theme_id " .. e.theme_id);
	eq.debug("points " .. e.points);
}
```
## EVENT_LDON_POINTS_LOSS

``` lua

function EVENT_LDON_POINTS_LOSS(e) {
	eq.debug("theme_id " .. e.theme_id);
	eq.debug("points " .. e.points);
}
```
## EVENT_LEAVE_AREA

``` lua

function EVENT_LEAVE_AREA(e) {
	eq.debug("area_id " .. e.area_id);
	eq.debug("area_type " .. e.area_type);
}
```
## EVENT_LEVEL_DOWN

``` lua

function EVENT_LEVEL_DOWN(e) {
	eq.debug("levels_lost " .. e.levels_lost);
}
```
## EVENT_LEVEL_UP

``` lua

function EVENT_LEVEL_UP(e) {
	eq.debug("levels_gained " .. e.levels_gained);
}
```
## EVENT_LOOT

``` lua

function EVENT_LOOT(e) {
	eq.debug("item " .. e.item);
	eq.debug("corpse " .. e.corpse);
}
```
## EVENT_MEMORIZE_SPELL

``` lua

function EVENT_MEMORIZE_SPELL(e) {
	eq.debug("slot_id " .. e.slot_id);
	eq.debug("spell_id " .. e.spell_id);
	eq.debug("spell " .. e.spell);
}
```
## EVENT_MERCHANT_BUY

``` lua

function EVENT_MERCHANT_BUY(e) {
	eq.debug("npc_id " .. e.npc_id);
	eq.debug("merchant_id " .. e.merchant_id);
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_quantity " .. e.item_quantity);
	eq.debug("item_cost " .. e.item_cost);
}
```
## EVENT_MERCHANT_SELL

``` lua

function EVENT_MERCHANT_SELL(e) {
	eq.debug("npc_id " .. e.npc_id);
	eq.debug("merchant_id " .. e.merchant_id);
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_quantity " .. e.item_quantity);
	eq.debug("item_cost " .. e.item_cost);
}
```
## EVENT_PAYLOAD

``` lua

function EVENT_PAYLOAD(e) {
	eq.debug("payload_id " .. e.payload_id);
	eq.debug("payload_value " .. e.payload_value);
}
```
## EVENT_PLAYER_PICKUP

``` lua

function EVENT_PLAYER_PICKUP(e) {
	eq.debug("item " .. e.item);
}
```
## EVENT_POPUP_RESPONSE

``` lua

function EVENT_POPUP_RESPONSE(e) {
	eq.debug("popup_id " .. e.popup_id);
}
```
## EVENT_READ_ITEM

``` lua

function EVENT_READ_ITEM(e) {
	eq.debug("text_file " .. e.text_file);
	eq.debug("item_id " .. e.item_id);
	eq.debug("book_text " .. e.book_text);
	eq.debug("can_cast " .. e.can_cast);
	eq.debug("can_scribe " .. e.can_scribe);
	eq.debug("slot_id " .. e.slot_id);
	eq.debug("target_id " .. e.target_id);
	eq.debug("type " .. e.type);
	eq.debug("item " .. e.item);
}
```
## EVENT_RESPAWN

``` lua

function EVENT_RESPAWN(e) {
	eq.debug("option " .. e.option);
	eq.debug("resurrect " .. e.resurrect);
}
```
## EVENT_SAY

``` lua

function EVENT_SAY(e) {
	eq.debug("message " .. e.message);
	eq.debug("language " .. e.language);
}
```
## EVENT_SCRIBE_SPELL

``` lua

function EVENT_SCRIBE_SPELL(e) {
	eq.debug("slot_id " .. e.slot_id);
	eq.debug("spell_id " .. e.spell_id);
	eq.debug("spell " .. e.spell);
}
```
## EVENT_SIGNAL

``` lua

function EVENT_SIGNAL(e) {
	eq.debug("signal " .. e.signal);
}
```
## EVENT_SKILL_UP

``` lua

function EVENT_SKILL_UP(e) {
	eq.debug("skill_id " .. e.skill_id);
	eq.debug("skill_value " .. e.skill_value);
	eq.debug("skill_max " .. e.skill_max);
	eq.debug("is_tradeskill " .. e.is_tradeskill);
}
```
## EVENT_SPELL_BLOCKED

``` lua

function EVENT_SPELL_BLOCKED(e) {
	eq.debug("blocking_spell_id " .. e.blocking_spell_id);
	eq.debug("cast_spell_id " .. e.cast_spell_id);
	eq.debug("blocking_spell " .. e.blocking_spell);
	eq.debug("cast_spell " .. e.cast_spell);
}
```
## EVENT_TARGET_CHANGE

``` lua

function EVENT_TARGET_CHANGE(e) {
	eq.debug("other " .. e.other);
}
```
## EVENT_TASK_ACCEPTED

``` lua

function EVENT_TASK_ACCEPTED(e) {
	eq.debug("task_id " .. e.task_id);
}
```
## EVENT_TASK_BEFORE_UPDATE

``` lua

function EVENT_TASK_BEFORE_UPDATE(e) {
	eq.debug("count " .. e.count);
	eq.debug("activity_id " .. e.activity_id);
	eq.debug("task_id " .. e.task_id);
}
```
## EVENT_TASK_COMPLETE

``` lua

function EVENT_TASK_COMPLETE(e) {
	eq.debug("count " .. e.count);
	eq.debug("activity_id " .. e.activity_id);
	eq.debug("task_id " .. e.task_id);
}
```
## EVENT_TASK_FAIL

``` lua

function EVENT_TASK_FAIL(e) {
	eq.debug("task_id " .. e.task_id);
}
```
## EVENT_TASK_STAGE_COMPLETE

``` lua

function EVENT_TASK_STAGE_COMPLETE(e) {
	eq.debug("task_id " .. e.task_id);
	eq.debug("activity_id " .. e.activity_id);
}
```
## EVENT_TASK_UPDATE

``` lua

function EVENT_TASK_UPDATE(e) {
	eq.debug("count " .. e.count);
	eq.debug("activity_id " .. e.activity_id);
	eq.debug("task_id " .. e.task_id);
}
```
## EVENT_TEST_BUFF

``` lua

function EVENT_TEST_BUFF(e) {
}
```
## EVENT_TIMER

``` lua

function EVENT_TIMER(e) {
	eq.debug("timer " .. e.timer);
}
```
## EVENT_TIMER_PAUSE

``` lua

function EVENT_TIMER_PAUSE(e) {
	eq.debug("timer " .. e.timer);
	eq.debug("duration " .. e.duration);
}
```
## EVENT_TIMER_RESUME

``` lua

function EVENT_TIMER_RESUME(e) {
	eq.debug("timer " .. e.timer);
	eq.debug("duration " .. e.duration);
}
```
## EVENT_TIMER_START

``` lua

function EVENT_TIMER_START(e) {
	eq.debug("timer " .. e.timer);
	eq.debug("duration " .. e.duration);
}
```
## EVENT_TIMER_STOP

``` lua

function EVENT_TIMER_STOP(e) {
	eq.debug("timer " .. e.timer);
}
```
## EVENT_UNEQUIP_ITEM_CLIENT

``` lua

function EVENT_UNEQUIP_ITEM_CLIENT(e) {
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_quantity " .. e.item_quantity);
	eq.debug("slot_id " .. e.slot_id);
	eq.debug("item " .. e.item);
}
```
## EVENT_UNHANDLED_OPCODE

``` lua

function EVENT_UNHANDLED_OPCODE(e) {
	eq.debug("packet " .. e.packet);
	eq.debug("connecting " .. e.connecting);
}
```
## EVENT_UNMEMORIZE_SPELL

``` lua

function EVENT_UNMEMORIZE_SPELL(e) {
	eq.debug("slot_id " .. e.slot_id);
	eq.debug("spell_id " .. e.spell_id);
	eq.debug("spell " .. e.spell);
}
```
## EVENT_UNSCRIBE_SPELL

``` lua

function EVENT_UNSCRIBE_SPELL(e) {
	eq.debug("slot_id " .. e.slot_id);
	eq.debug("spell_id " .. e.spell_id);
	eq.debug("spell " .. e.spell);
}
```
## EVENT_USE_SKILL

``` lua

function EVENT_USE_SKILL(e) {
	eq.debug("skill_id " .. e.skill_id);
	eq.debug("skill_level " .. e.skill_level);
}
```
## EVENT_WARP

``` lua

function EVENT_WARP(e) {
	eq.debug("from_x " .. e.from_x);
	eq.debug("from_y " .. e.from_y);
	eq.debug("from_z " .. e.from_z);
}
```
## EVENT_ZONE

``` lua

function EVENT_ZONE(e) {
	eq.debug("from_zone_id " .. e.from_zone_id);
	eq.debug("from_instance_id " .. e.from_instance_id);
	eq.debug("from_instance_version " .. e.from_instance_version);
	eq.debug("zone_id " .. e.zone_id);
	eq.debug("instance_id " .. e.instance_id);
	eq.debug("instance_version " .. e.instance_version);
}
```
