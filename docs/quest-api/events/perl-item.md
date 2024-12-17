!!! info end

    Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl){:target="perl_event"} for latest definitions and Quest examples

## EVENT_AUGMENT_INSERT

``` perl

sub EVENT_AUGMENT_INSERT {
}
```
## EVENT_AUGMENT_ITEM

``` perl

sub EVENT_AUGMENT_ITEM {
}
```
## EVENT_AUGMENT_REMOVE

``` perl

sub EVENT_AUGMENT_REMOVE {
}
```
## EVENT_DESTROY_ITEM

``` perl

sub EVENT_DESTROY_ITEM {
}
```
## EVENT_DROP_ITEM

``` perl

sub EVENT_DROP_ITEM {
	quest::debug("quantity " . $quantity);
	quest::debug("itemname " . $itemname);
	quest::debug("itemid " . $itemid);
	quest::debug("spell_id " . $spell_id);
	quest::debug("slotid " . $slotid);
}
```
## EVENT_EQUIP_ITEM

``` perl

sub EVENT_EQUIP_ITEM {
}
```
## EVENT_ITEM_CLICK

``` perl

sub EVENT_ITEM_CLICK {
	quest::debug("itemid " . $itemid);
	quest::debug("itemname " . $itemname);
	quest::debug("slotid " . $slotid);
	quest::debug("spell_id " . $spell_id);
	quest::debug("spell " . $spell);
}
```
## EVENT_ITEM_CLICK_CAST

``` perl

sub EVENT_ITEM_CLICK_CAST {
	quest::debug("itemid " . $itemid);
	quest::debug("itemname " . $itemname);
	quest::debug("slotid " . $slotid);
	quest::debug("spell_id " . $spell_id);
	quest::debug("spell " . $spell);
}
```
## EVENT_ITEM_ENTER_ZONE

``` perl

sub EVENT_ITEM_ENTER_ZONE {
	quest::debug("itemid " . $itemid);
	quest::debug("itemname " . $itemname);
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
## EVENT_SCALE_CALC

``` perl

sub EVENT_SCALE_CALC {
	quest::debug("itemid " . $itemid);
	quest::debug("itemname " . $itemname);
}
```
## EVENT_TIMER

``` perl

sub EVENT_TIMER {
	quest::debug("timer " . $timer);
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
## EVENT_UNAUGMENT_ITEM

``` perl

sub EVENT_UNAUGMENT_ITEM {
}
```
## EVENT_UNEQUIP_ITEM

``` perl

sub EVENT_UNEQUIP_ITEM {
}
```
## EVENT_WEAPON_PROC

``` perl

sub EVENT_WEAPON_PROC {
}
```
