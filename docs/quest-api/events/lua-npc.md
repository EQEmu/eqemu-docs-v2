!!! info end

    Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua){:target="perl_event"} for latest definitions and Quest examples

    Last generated 2023.07.15

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
}
```
## EVENT_CAST_BEGIN

``` lua

function EVENT_CAST_BEGIN(e) {
	eq.debug("spell " .. e.spell);
}
```
## EVENT_CAST_ON

``` lua

function EVENT_CAST_ON(e) {
	eq.debug("spell " .. e.spell);
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
