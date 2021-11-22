## EVENT_AUGMENT_INSERT

=== "EVENT_AUGMENT_INSERT"

``` perl

sub EVENT_AUGMENT_INSERT {
}
```
## EVENT_AUGMENT_ITEM

=== "EVENT_AUGMENT_ITEM"

``` perl

sub EVENT_AUGMENT_ITEM {
}
```
## EVENT_AUGMENT_REMOVE

=== "EVENT_AUGMENT_REMOVE"

``` perl

sub EVENT_AUGMENT_REMOVE {
}
```
## EVENT_DESTROY_ITEM

=== "EVENT_DESTROY_ITEM"

``` perl

sub EVENT_DESTROY_ITEM {
}
```
## EVENT_DROP_ITEM

=== "EVENT_DROP_ITEM"

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

=== "EVENT_EQUIP_ITEM"

``` perl

sub EVENT_EQUIP_ITEM {
}
```
## EVENT_ITEM_CLICK

=== "EVENT_ITEM_CLICK"

``` perl

sub EVENT_ITEM_CLICK {
	quest::debug("itemid " . $itemid);
	quest::debug("itemname " . $itemname);
	quest::debug("slotid " . $slotid);
	quest::debug("spell_id " . $spell_id);
}
```
## EVENT_ITEM_CLICK_CAST

=== "EVENT_ITEM_CLICK_CAST"

``` perl

sub EVENT_ITEM_CLICK_CAST {
	quest::debug("itemid " . $itemid);
	quest::debug("itemname " . $itemname);
	quest::debug("slotid " . $slotid);
	quest::debug("spell_id " . $spell_id);
}
```
## EVENT_ITEM_ENTER_ZONE

=== "EVENT_ITEM_ENTER_ZONE"

``` perl

sub EVENT_ITEM_ENTER_ZONE {
	quest::debug("itemid " . $itemid);
	quest::debug("itemname " . $itemname);
}
```
## EVENT_ITEM_TICK

=== "EVENT_ITEM_TICK"

``` perl

sub EVENT_ITEM_TICK {
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
## EVENT_SCALE_CALC

=== "EVENT_SCALE_CALC"

``` perl

sub EVENT_SCALE_CALC {
	quest::debug("itemid " . $itemid);
	quest::debug("itemname " . $itemname);
}
```
## EVENT_TIMER

=== "EVENT_TIMER"

``` perl

sub EVENT_TIMER {
	quest::debug("timer " . $timer);
}
```
## EVENT_UNAUGMENT_ITEM

=== "EVENT_UNAUGMENT_ITEM"

``` perl

sub EVENT_UNAUGMENT_ITEM {
}
```
## EVENT_UNEQUIP_ITEM

=== "EVENT_UNEQUIP_ITEM"

``` perl

sub EVENT_UNEQUIP_ITEM {
}
```
## EVENT_WEAPON_PROC

=== "EVENT_WEAPON_PROC"

``` perl

sub EVENT_WEAPON_PROC {
}
```
