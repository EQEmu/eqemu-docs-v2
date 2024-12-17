!!! info end

    Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua){:target="perl_event"} for latest definitions and Quest examples

## EVENT_AGGRO_SAY

``` lua

function EVENT_AGGRO_SAY(e) {
	eq.debug("other " .. e.other);
	eq.debug("message " .. e.message);
	eq.debug("language " .. e.language);
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
## EVENT_COMBAT

``` lua

function EVENT_COMBAT(e) {
	eq.debug("other " .. e.other);
	eq.debug("joined " .. e.joined);
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
	eq.debug("skill_id " .. e.skill_id);
	eq.debug("killed_entity_id " .. e.killed_entity_id);
	eq.debug("combat_start_time " .. e.combat_start_time);
	eq.debug("combat_end_time " .. e.combat_end_time);
	eq.debug("damage_received " .. e.damage_received);
	eq.debug("healing_received " .. e.healing_received);
	eq.debug("corpse " .. e.corpse);
	eq.debug("killed " .. e.killed);
}
```
## EVENT_DEATH_COMPLETE

``` lua

function EVENT_DEATH_COMPLETE(e) {
	eq.debug("other " .. e.other);
	eq.debug("killer_id " .. e.killer_id);
	eq.debug("damage " .. e.damage);
	eq.debug("spell " .. e.spell);
	eq.debug("skill_id " .. e.skill_id);
	eq.debug("killed_entity_id " .. e.killed_entity_id);
	eq.debug("combat_start_time " .. e.combat_start_time);
	eq.debug("combat_end_time " .. e.combat_end_time);
	eq.debug("damage_received " .. e.damage_received);
	eq.debug("healing_received " .. e.healing_received);
	eq.debug("corpse " .. e.corpse);
	eq.debug("killed " .. e.killed);
}
```
## EVENT_DEATH_ZONE

``` lua

function EVENT_DEATH_ZONE(e) {
	eq.debug("other " .. e.other);
	eq.debug("killer_id " .. e.killer_id);
	eq.debug("damage " .. e.damage);
	eq.debug("spell " .. e.spell);
	eq.debug("skill_id " .. e.skill_id);
	eq.debug("killed_entity_id " .. e.killed_entity_id);
	eq.debug("combat_start_time " .. e.combat_start_time);
	eq.debug("combat_end_time " .. e.combat_end_time);
	eq.debug("damage_received " .. e.damage_received);
	eq.debug("healing_received " .. e.healing_received);
	eq.debug("corpse " .. e.corpse);
	eq.debug("killed " .. e.killed);
}
```
## EVENT_DESPAWN

``` lua

function EVENT_DESPAWN(e) {
}
```
## EVENT_DESPAWN_ZONE

``` lua

function EVENT_DESPAWN_ZONE(e) {
	eq.debug("other " .. e.other);
}
```
## EVENT_ENTER

``` lua

function EVENT_ENTER(e) {
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
## EVENT_EXIT

``` lua

function EVENT_EXIT(e) {
	eq.debug("other " .. e.other);
}
```
## EVENT_FEIGN_DEATH

``` lua

function EVENT_FEIGN_DEATH(e) {
	eq.debug("other " .. e.other);
}
```
## EVENT_HATE_LIST

``` lua

function EVENT_HATE_LIST(e) {
	eq.debug("other " .. e.other);
	eq.debug("joined " .. e.joined);
}
```
## EVENT_HP

``` lua

function EVENT_HP(e) {
	eq.debug("hp_event " .. e.hp_event);
	eq.debug("inc_hp_event " .. e.inc_hp_event);
}
```
## EVENT_KILLED_MERIT

``` lua

function EVENT_KILLED_MERIT(e) {
	eq.debug("other " .. e.other);
}
```
## EVENT_LEAVE_AREA

``` lua

function EVENT_LEAVE_AREA(e) {
	eq.debug("area_id " .. e.area_id);
	eq.debug("area_type " .. e.area_type);
}
```
## EVENT_LOOT_ADDED

``` lua

function EVENT_LOOT_ADDED(e) {
	eq.debug("item " .. e.item);
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_name " .. e.item_name);
	eq.debug("item_charges " .. e.item_charges);
	eq.debug("augment_one " .. e.augment_one);
	eq.debug("augment_two " .. e.augment_two);
	eq.debug("augment_three " .. e.augment_three);
	eq.debug("augment_four " .. e.augment_four);
	eq.debug("augment_five " .. e.augment_five);
	eq.debug("augment_six " .. e.augment_six);
}
```
## EVENT_LOOT_ZONE

``` lua

function EVENT_LOOT_ZONE(e) {
	eq.debug("other " .. e.other);
	eq.debug("item " .. e.item);
	eq.debug("corpse " .. e.corpse);
}
```
## EVENT_PAYLOAD

``` lua

function EVENT_PAYLOAD(e) {
	eq.debug("payload_id " .. e.payload_id);
	eq.debug("payload_value " .. e.payload_value);
}
```
## EVENT_POPUP_RESPONSE

``` lua

function EVENT_POPUP_RESPONSE(e) {
	eq.debug("other " .. e.other);
	eq.debug("popup_id " .. e.popup_id);
}
```
## EVENT_PROXIMITY_SAY

``` lua

function EVENT_PROXIMITY_SAY(e) {
	eq.debug("other " .. e.other);
	eq.debug("message " .. e.message);
	eq.debug("language " .. e.language);
}
```
## EVENT_SAY

``` lua

function EVENT_SAY(e) {
	eq.debug("other " .. e.other);
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
## EVENT_SLAY

``` lua

function EVENT_SLAY(e) {
	eq.debug("other " .. e.other);
}
```
## EVENT_SPAWN

``` lua

function EVENT_SPAWN(e) {
}
```
## EVENT_SPAWN_ZONE

``` lua

function EVENT_SPAWN_ZONE(e) {
	eq.debug("other " .. e.other);
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
	eq.debug("other " .. e.other);
	eq.debug("task_id " .. e.task_id);
}
```
## EVENT_TICK

``` lua

function EVENT_TICK(e) {
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
## EVENT_TRADE

``` lua

function EVENT_TRADE(e) {
	eq.debug("other " .. e.other);
	eq.debug("platinum " .. e.platinum);
	eq.debug("gold " .. e.gold);
	eq.debug("silver " .. e.silver);
	eq.debug("copper " .. e.copper);
	eq.debug("trade " .. e.trade);
}
```
## EVENT_WAYPOINT_ARRIVE

``` lua

function EVENT_WAYPOINT_ARRIVE(e) {
	eq.debug("other " .. e.other);
	eq.debug("wp " .. e.wp);
}
```
## EVENT_WAYPOINT_DEPART

``` lua

function EVENT_WAYPOINT_DEPART(e) {
	eq.debug("other " .. e.other);
	eq.debug("wp " .. e.wp);
}
```
