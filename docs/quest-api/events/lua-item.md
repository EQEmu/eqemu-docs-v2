!!! info end

    Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua){:target="perl_event"} for latest definitions and Quest examples

## EVENT_AUGMENT_INSERT

``` lua

function EVENT_AUGMENT_INSERT(e) {
	eq.debug("item " .. e.item);
	eq.debug("slot_id " .. e.slot_id);
}
```
## EVENT_AUGMENT_ITEM

``` lua

function EVENT_AUGMENT_ITEM(e) {
	eq.debug("aug " .. e.aug);
	eq.debug("slot_id " .. e.slot_id);
}
```
## EVENT_AUGMENT_REMOVE

``` lua

function EVENT_AUGMENT_REMOVE(e) {
	eq.debug("item " .. e.item);
	eq.debug("slot_id " .. e.slot_id);
	eq.debug("destroyed " .. e.destroyed);
}
```
## EVENT_DESTROY_ITEM

``` lua

function EVENT_DESTROY_ITEM(e) {
}
```
## EVENT_DROP_ITEM

``` lua

function EVENT_DROP_ITEM(e) {
}
```
## EVENT_EQUIP_ITEM

``` lua

function EVENT_EQUIP_ITEM(e) {
	eq.debug("slot_id " .. e.slot_id);
}
```
## EVENT_ITEM_CLICK

``` lua

function EVENT_ITEM_CLICK(e) {
	eq.debug("slot_id " .. e.slot_id);
}
```
## EVENT_ITEM_CLICK_CAST

``` lua

function EVENT_ITEM_CLICK_CAST(e) {
	eq.debug("slot_id " .. e.slot_id);
}
```
## EVENT_ITEM_ENTER_ZONE

``` lua

function EVENT_ITEM_ENTER_ZONE(e) {
}
```
## EVENT_LOOT

``` lua

function EVENT_LOOT(e) {
	eq.debug("corpse " .. e.corpse);
}
```
## EVENT_SCALE_CALC

``` lua

function EVENT_SCALE_CALC(e) {
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
## EVENT_UNAUGMENT_ITEM

``` lua

function EVENT_UNAUGMENT_ITEM(e) {
	eq.debug("aug " .. e.aug);
	eq.debug("slot_id " .. e.slot_id);
}
```
## EVENT_UNEQUIP_ITEM

``` lua

function EVENT_UNEQUIP_ITEM(e) {
	eq.debug("slot_id " .. e.slot_id);
}
```
## EVENT_WEAPON_PROC

``` lua

function EVENT_WEAPON_PROC(e) {
	eq.debug("target " .. e.target);
	eq.debug("spell " .. e.spell);
}
```
