## EVENT_AUGMENT_INSERT

=== "EVENT_AUGMENT_INSERT"

    ``` lua

	function EVENT_AUGMENT_INSERT(e) {
		eq.debug("item " .. item);
		eq.debug("slot_id " .. slot_id);
	}
    ```
## EVENT_AUGMENT_ITEM

=== "EVENT_AUGMENT_ITEM"

    ``` lua

	function EVENT_AUGMENT_ITEM(e) {
		eq.debug("aug " .. aug);
		eq.debug("slot_id " .. slot_id);
	}
    ```
## EVENT_AUGMENT_REMOVE

=== "EVENT_AUGMENT_REMOVE"

    ``` lua

	function EVENT_AUGMENT_REMOVE(e) {
		eq.debug("item " .. item);
		eq.debug("slot_id " .. slot_id);
		eq.debug("destroyed " .. destroyed);
	}
    ```
## EVENT_DESTROY_ITEM

=== "EVENT_DESTROY_ITEM"

    ``` lua

	function EVENT_DESTROY_ITEM(e) {
	}
    ```
## EVENT_DROP_ITEM

=== "EVENT_DROP_ITEM"

    ``` lua

	function EVENT_DROP_ITEM(e) {
	}
    ```
## EVENT_EQUIP_ITEM

=== "EVENT_EQUIP_ITEM"

    ``` lua

	function EVENT_EQUIP_ITEM(e) {
		eq.debug("slot_id " .. slot_id);
	}
    ```
## EVENT_ITEM_CLICK

=== "EVENT_ITEM_CLICK"

    ``` lua

	function EVENT_ITEM_CLICK(e) {
		eq.debug("slot_id " .. slot_id);
	}
    ```
## EVENT_ITEM_CLICK_CAST

=== "EVENT_ITEM_CLICK_CAST"

    ``` lua

	function EVENT_ITEM_CLICK_CAST(e) {
		eq.debug("slot_id " .. slot_id);
	}
    ```
## EVENT_ITEM_ENTER_ZONE

=== "EVENT_ITEM_ENTER_ZONE"

    ``` lua

	function EVENT_ITEM_ENTER_ZONE(e) {
	}
    ```
## EVENT_ITEM_TICK

=== "EVENT_ITEM_TICK"

    ``` lua

	function EVENT_ITEM_TICK(e) {
	}
    ```
## EVENT_LOOT

=== "EVENT_LOOT"

    ``` lua

	function EVENT_LOOT(e) {
		eq.debug("corpse " .. corpse);
	}
    ```
## EVENT_SCALE_CALC

=== "EVENT_SCALE_CALC"

    ``` lua

	function EVENT_SCALE_CALC(e) {
	}
    ```
## EVENT_TIMER

=== "EVENT_TIMER"

    ``` lua

	function EVENT_TIMER(e) {
		eq.debug("timer " .. timer);
	}
    ```
## EVENT_UNAUGMENT_ITEM

=== "EVENT_UNAUGMENT_ITEM"

    ``` lua

	function EVENT_UNAUGMENT_ITEM(e) {
		eq.debug("aug " .. aug);
		eq.debug("slot_id " .. slot_id);
	}
    ```
## EVENT_UNEQUIP_ITEM

=== "EVENT_UNEQUIP_ITEM"

    ``` lua

	function EVENT_UNEQUIP_ITEM(e) {
		eq.debug("slot_id " .. slot_id);
	}
    ```
## EVENT_WEAPON_PROC

=== "EVENT_WEAPON_PROC"

    ``` lua

	function EVENT_WEAPON_PROC(e) {
		eq.debug("target " .. target);
		eq.debug("spell " .. spell);
	}
    ```
