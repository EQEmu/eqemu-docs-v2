## EVENT_BOT_COMMAND

``` lua

function EVENT_BOT_COMMAND(e) {
	eq.debug("bot_command " .. e.bot_command);
	eq.debug("args " .. e.args);
}
```
## EVENT_CAST

``` lua

function EVENT_CAST(e) {
	eq.debug("spell " .. e.spell);
	eq.debug("caster_id " .. e.caster_id);
	eq.debug("caster_level " .. e.caster_level);
}
```
## EVENT_CAST_BEGIN

``` lua

function EVENT_CAST_BEGIN(e) {
	eq.debug("spell " .. e.spell);
	eq.debug("caster_id " .. e.caster_id);
	eq.debug("caster_level " .. e.caster_level);
}
```
## EVENT_CAST_ON

``` lua

function EVENT_CAST_ON(e) {
	eq.debug("spell " .. e.spell);
	eq.debug("caster_id " .. e.caster_id);
	eq.debug("caster_level " .. e.caster_level);
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
}
```
## EVENT_CONSIDER_CORPSE

``` lua

function EVENT_CONSIDER_CORPSE(e) {
	eq.debug("corpse_entity_id " .. e.corpse_entity_id);
}
```
## EVENT_DEATH

``` lua

function EVENT_DEATH(e) {
	eq.debug("other " .. e.other);
	eq.debug("damage " .. e.damage);
	eq.debug("spell " .. e.spell);
	eq.debug("skill " .. e.skill);
}
```
## EVENT_DEATH_COMPLETE

``` lua

function EVENT_DEATH_COMPLETE(e) {
	eq.debug("other " .. e.other);
	eq.debug("damage " .. e.damage);
	eq.debug("spell " .. e.spell);
	eq.debug("skill " .. e.skill);
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
## EVENT_GROUP_CHANGE

``` lua

function EVENT_GROUP_CHANGE(e) {
}
```
## EVENT_LANGUAGE_SKILL_UP

``` lua

function EVENT_LANGUAGE_SKILL_UP(e) {
	eq.debug("skill_id " .. e.skill_id);
	eq.debug("skill_value " .. e.skill_value);
	eq.debug("skill_max " .. e.skill_max);
	eq.debug("is_tradeskill " .. e.is_tradeskill);
}
```
## EVENT_LEAVE_AREA

``` lua

function EVENT_LEAVE_AREA(e) {
	eq.debug("area_id " .. e.area_id);
	eq.debug("area_type " .. e.area_type);
}
```
## EVENT_LEVEL_UP

``` lua

function EVENT_LEVEL_UP(e) {
}
```
## EVENT_LOOT

``` lua

function EVENT_LOOT(e) {
	eq.debug("item " .. e.item);
	eq.debug("corpse " .. e.corpse);
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
## EVENT_TARGET_CHANGE

``` lua

function EVENT_TARGET_CHANGE(e) {
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
	eq.debug("zone_id " .. e.zone_id);
}
```
