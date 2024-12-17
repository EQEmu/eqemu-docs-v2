!!! info end

    Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua){:target="perl_event"} for latest definitions and Quest examples

## EVENT_AGGRO_SAY

``` lua

function EVENT_AGGRO_SAY(e) {
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
	eq.debug("skill " .. e.skill);
	eq.debug("killed_entity_id " .. e.killed_entity_id);
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
}
```
## EVENT_DESPAWN

``` lua

function EVENT_DESPAWN(e) {
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
## EVENT_EQUIP_ITEM_BOT

``` lua

function EVENT_EQUIP_ITEM_BOT(e) {
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_quantity " .. e.item_quantity);
	eq.debug("slot_id " .. e.slot_id);
	eq.debug("item " .. e.item);
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
## EVENT_SPELL_BLOCKED

``` lua

function EVENT_SPELL_BLOCKED(e) {
}
```
## EVENT_TARGET_CHANGE

``` lua

function EVENT_TARGET_CHANGE(e) {
	eq.debug("other " .. e.other);
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
## EVENT_UNEQUIP_ITEM_BOT

``` lua

function EVENT_UNEQUIP_ITEM_BOT(e) {
	eq.debug("item_id " .. e.item_id);
	eq.debug("item_quantity " .. e.item_quantity);
	eq.debug("slot_id " .. e.slot_id);
	eq.debug("item " .. e.item);
}
```
## EVENT_USE_SKILL

``` lua

function EVENT_USE_SKILL(e) {
	eq.debug("skill_id " .. e.skill_id);
	eq.debug("skill_level " .. e.skill_level);
}
```
