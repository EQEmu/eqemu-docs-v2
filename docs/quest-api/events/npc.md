## EVENT_AGGRO_SAY

=== "EVENT_AGGRO_SAY"

    ``` lua

	function EVENT_AGGRO_SAY(e) {
		eq.debug("other " .. other);
		eq.debug("message " .. message);
		eq.debug("language " .. language);
	}
    ```
## EVENT_CAST

=== "EVENT_CAST"

    ``` lua

	function EVENT_CAST(e) {
		eq.debug("spell " .. spell);
	}
    ```
## EVENT_CAST_BEGIN

=== "EVENT_CAST_BEGIN"

    ``` lua

	function EVENT_CAST_BEGIN(e) {
		eq.debug("spell " .. spell);
	}
    ```
## EVENT_CAST_ON

=== "EVENT_CAST_ON"

    ``` lua

	function EVENT_CAST_ON(e) {
		eq.debug("spell " .. spell);
	}
    ```
## EVENT_COMBAT

=== "EVENT_COMBAT"

    ``` lua

	function EVENT_COMBAT(e) {
		eq.debug("other " .. other);
		eq.debug("joined " .. joined);
	}
    ```
## EVENT_DEATH

=== "EVENT_DEATH"

    ``` lua

	function EVENT_DEATH(e) {
		eq.debug("other " .. other);
		eq.debug("damage " .. damage);
		eq.debug("spell " .. spell);
		eq.debug("skill_id " .. skill_id);
	}
    ```
## EVENT_DEATH_COMPLETE

=== "EVENT_DEATH_COMPLETE"

    ``` lua

	function EVENT_DEATH_COMPLETE(e) {
		eq.debug("other " .. other);
		eq.debug("damage " .. damage);
		eq.debug("spell " .. spell);
		eq.debug("skill_id " .. skill_id);
	}
    ```
## EVENT_DEATH_ZONE

=== "EVENT_DEATH_ZONE"

    ``` lua

	function EVENT_DEATH_ZONE(e) {
	}
    ```
## EVENT_ENTER

=== "EVENT_ENTER"

    ``` lua

	function EVENT_ENTER(e) {
		eq.debug("other " .. other);
	}
    ```
## EVENT_ENTER_AREA

=== "EVENT_ENTER_AREA"

    ``` lua

	function EVENT_ENTER_AREA(e) {
		eq.debug("area_id " .. area_id);
		eq.debug("area_type " .. area_type);
	}
    ```
## EVENT_EXIT

=== "EVENT_EXIT"

    ``` lua

	function EVENT_EXIT(e) {
		eq.debug("other " .. other);
	}
    ```
## EVENT_FEIGN_DEATH

=== "EVENT_FEIGN_DEATH"

    ``` lua

	function EVENT_FEIGN_DEATH(e) {
		eq.debug("other " .. other);
	}
    ```
## EVENT_HATE_LIST

=== "EVENT_HATE_LIST"

    ``` lua

	function EVENT_HATE_LIST(e) {
		eq.debug("other " .. other);
		eq.debug("joined " .. joined);
	}
    ```
## EVENT_HP

=== "EVENT_HP"

    ``` lua

	function EVENT_HP(e) {
		eq.debug("hp_event " .. hp_event);
		eq.debug("inc_hp_event " .. inc_hp_event);
	}
    ```
## EVENT_KILLED_MERIT

=== "EVENT_KILLED_MERIT"

    ``` lua

	function EVENT_KILLED_MERIT(e) {
		eq.debug("other " .. other);
	}
    ```
## EVENT_LEAVE_AREA

=== "EVENT_LEAVE_AREA"

    ``` lua

	function EVENT_LEAVE_AREA(e) {
		eq.debug("area_id " .. area_id);
		eq.debug("area_type " .. area_type);
	}
    ```
## EVENT_LOOT_ZONE

=== "EVENT_LOOT_ZONE"

    ``` lua

	function EVENT_LOOT_ZONE(e) {
		eq.debug("other " .. other);
		eq.debug("item " .. item);
		eq.debug("corpse " .. corpse);
	}
    ```
## EVENT_POPUP_RESPONSE

=== "EVENT_POPUP_RESPONSE"

    ``` lua

	function EVENT_POPUP_RESPONSE(e) {
		eq.debug("other " .. other);
		eq.debug("popup_id " .. popup_id);
	}
    ```
## EVENT_PROXIMITY_SAY

=== "EVENT_PROXIMITY_SAY"

    ``` lua

	function EVENT_PROXIMITY_SAY(e) {
		eq.debug("other " .. other);
		eq.debug("message " .. message);
		eq.debug("language " .. language);
	}
    ```
## EVENT_SAY

=== "EVENT_SAY"

    ``` lua

	function EVENT_SAY(e) {
		eq.debug("other " .. other);
		eq.debug("message " .. message);
		eq.debug("language " .. language);
	}
    ```
## EVENT_SIGNAL

=== "EVENT_SIGNAL"

    ``` lua

	function EVENT_SIGNAL(e) {
		eq.debug("signal " .. signal);
	}
    ```
## EVENT_SLAY

=== "EVENT_SLAY"

    ``` lua

	function EVENT_SLAY(e) {
		eq.debug("other " .. other);
	}
    ```
## EVENT_SPAWN

=== "EVENT_SPAWN"

    ``` lua

	function EVENT_SPAWN(e) {
	}
    ```
## EVENT_SPAWN_ZONE

=== "EVENT_SPAWN_ZONE"

    ``` lua

	function EVENT_SPAWN_ZONE(e) {
	}
    ```
## EVENT_TARGET_CHANGE

=== "EVENT_TARGET_CHANGE"

    ``` lua

	function EVENT_TARGET_CHANGE(e) {
		eq.debug("other " .. other);
	}
    ```
## EVENT_TASK_ACCEPTED

=== "EVENT_TASK_ACCEPTED"

    ``` lua

	function EVENT_TASK_ACCEPTED(e) {
		eq.debug("other " .. other);
		eq.debug("task_id " .. task_id);
	}
    ```
## EVENT_TICK

=== "EVENT_TICK"

    ``` lua

	function EVENT_TICK(e) {
	}
    ```
## EVENT_TIMER

=== "EVENT_TIMER"

    ``` lua

	function EVENT_TIMER(e) {
		eq.debug("timer " .. timer);
	}
    ```
## EVENT_TRADE

=== "EVENT_TRADE"

    ``` lua

	function EVENT_TRADE(e) {
		eq.debug("other " .. other);
		eq.debug("platinum " .. platinum);
		eq.debug("gold " .. gold);
		eq.debug("silver " .. silver);
		eq.debug("copper " .. copper);
		eq.debug("trade " .. trade);
	}
    ```
## EVENT_WAYPOINT_ARRIVE

=== "EVENT_WAYPOINT_ARRIVE"

    ``` lua

	function EVENT_WAYPOINT_ARRIVE(e) {
		eq.debug("other " .. other);
		eq.debug("wp " .. wp);
	}
    ```
## EVENT_WAYPOINT_DEPART

=== "EVENT_WAYPOINT_DEPART"

    ``` lua

	function EVENT_WAYPOINT_DEPART(e) {
		eq.debug("other " .. other);
		eq.debug("wp " .. wp);
	}
    ```
