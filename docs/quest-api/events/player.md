## EVENT_BOT_COMMAND

=== "EVENT_BOT_COMMAND"

    ``` lua

	function EVENT_BOT_COMMAND(e) {
		eq.debug("bot_command " .. bot_command);
		eq.debug("args " .. args);
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
## EVENT_CLICK_DOOR

=== "EVENT_CLICK_DOOR"

    ``` lua

	function EVENT_CLICK_DOOR(e) {
		eq.debug("door " .. door);
	}
    ```
## EVENT_CLICK_OBJECT

=== "EVENT_CLICK_OBJECT"

    ``` lua

	function EVENT_CLICK_OBJECT(e) {
		eq.debug("object " .. object);
	}
    ```
## EVENT_COMBINE

=== "EVENT_COMBINE"

    ``` lua

	function EVENT_COMBINE(e) {
		eq.debug("container_slot " .. container_slot);
	}
    ```
## EVENT_COMBINE_FAILURE

=== "EVENT_COMBINE_FAILURE"

    ``` lua

	function EVENT_COMBINE_FAILURE(e) {
		eq.debug("recipe_id " .. recipe_id);
		eq.debug("recipe_name " .. recipe_name);
	}
    ```
## EVENT_COMBINE_SUCCESS

=== "EVENT_COMBINE_SUCCESS"

    ``` lua

	function EVENT_COMBINE_SUCCESS(e) {
		eq.debug("recipe_id " .. recipe_id);
		eq.debug("recipe_name " .. recipe_name);
	}
    ```
## EVENT_COMBINE_VALIDATE

=== "EVENT_COMBINE_VALIDATE"

    ``` lua

	function EVENT_COMBINE_VALIDATE(e) {
		eq.debug("recipe_id " .. recipe_id);
		eq.debug("validate_type " .. validate_type);
		eq.debug("zone_id " .. zone_id);
		eq.debug("tradeskill_id " .. tradeskill_id);
	}
    ```
## EVENT_COMMAND

=== "EVENT_COMMAND"

    ``` lua

	function EVENT_COMMAND(e) {
		eq.debug("command " .. command);
		eq.debug("args " .. args);
	}
    ```
## EVENT_CONNECT

=== "EVENT_CONNECT"

    ``` lua

	function EVENT_CONNECT(e) {
	}
    ```
## EVENT_CONSIDER

=== "EVENT_CONSIDER"

    ``` lua

	function EVENT_CONSIDER(e) {
		eq.debug("entity_id " .. entity_id);
	}
    ```
## EVENT_CONSIDER_CORPSE

=== "EVENT_CONSIDER_CORPSE"

    ``` lua

	function EVENT_CONSIDER_CORPSE(e) {
		eq.debug("corpse_entity_id " .. corpse_entity_id);
	}
    ```
## EVENT_DEATH

=== "EVENT_DEATH"

    ``` lua

	function EVENT_DEATH(e) {
		eq.debug("other " .. other);
		eq.debug("damage " .. damage);
		eq.debug("spell " .. spell);
		eq.debug("skill " .. skill);
	}
    ```
## EVENT_DEATH_COMPLETE

=== "EVENT_DEATH_COMPLETE"

    ``` lua

	function EVENT_DEATH_COMPLETE(e) {
		eq.debug("other " .. other);
		eq.debug("damage " .. damage);
		eq.debug("spell " .. spell);
		eq.debug("skill " .. skill);
	}
    ```
## EVENT_DISCONNECT

=== "EVENT_DISCONNECT"

    ``` lua

	function EVENT_DISCONNECT(e) {
	}
    ```
## EVENT_DISCOVER_ITEM

=== "EVENT_DISCOVER_ITEM"

    ``` lua

	function EVENT_DISCOVER_ITEM(e) {
		eq.debug("item " .. item);
	}
    ```
## EVENT_DUEL_LOSE

=== "EVENT_DUEL_LOSE"

    ``` lua

	function EVENT_DUEL_LOSE(e) {
		eq.debug("other " .. other);
	}
    ```
## EVENT_DUEL_WIN

=== "EVENT_DUEL_WIN"

    ``` lua

	function EVENT_DUEL_WIN(e) {
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
## EVENT_ENTER_ZONE

=== "EVENT_ENTER_ZONE"

    ``` lua

	function EVENT_ENTER_ZONE(e) {
	}
    ```
## EVENT_ENVIRONMENTAL_DAMAGE

=== "EVENT_ENVIRONMENTAL_DAMAGE"

    ``` lua

	function EVENT_ENVIRONMENTAL_DAMAGE(e) {
		eq.debug("env_damage " .. env_damage);
		eq.debug("env_damage_type " .. env_damage_type);
		eq.debug("env_final_damage " .. env_final_damage);
	}
    ```
## EVENT_FEIGN_DEATH

=== "EVENT_FEIGN_DEATH"

    ``` lua

	function EVENT_FEIGN_DEATH(e) {
		eq.debug("other " .. other);
	}
    ```
## EVENT_FISH_FAILURE

=== "EVENT_FISH_FAILURE"

    ``` lua

	function EVENT_FISH_FAILURE(e) {
	}
    ```
## EVENT_FISH_START

=== "EVENT_FISH_START"

    ``` lua

	function EVENT_FISH_START(e) {
	}
    ```
## EVENT_FISH_SUCCESS

=== "EVENT_FISH_SUCCESS"

    ``` lua

	function EVENT_FISH_SUCCESS(e) {
		eq.debug("item " .. item);
	}
    ```
## EVENT_FORAGE_FAILURE

=== "EVENT_FORAGE_FAILURE"

    ``` lua

	function EVENT_FORAGE_FAILURE(e) {
	}
    ```
## EVENT_FORAGE_SUCCESS

=== "EVENT_FORAGE_SUCCESS"

    ``` lua

	function EVENT_FORAGE_SUCCESS(e) {
		eq.debug("item " .. item);
	}
    ```
## EVENT_GROUP_CHANGE

=== "EVENT_GROUP_CHANGE"

    ``` lua

	function EVENT_GROUP_CHANGE(e) {
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
## EVENT_LEVEL_UP

=== "EVENT_LEVEL_UP"

    ``` lua

	function EVENT_LEVEL_UP(e) {
	}
    ```
## EVENT_LOOT

=== "EVENT_LOOT"

    ``` lua

	function EVENT_LOOT(e) {
		eq.debug("item " .. item);
		eq.debug("corpse " .. corpse);
	}
    ```
## EVENT_PLAYER_PICKUP

=== "EVENT_PLAYER_PICKUP"

    ``` lua

	function EVENT_PLAYER_PICKUP(e) {
		eq.debug("item " .. item);
	}
    ```
## EVENT_POPUP_RESPONSE

=== "EVENT_POPUP_RESPONSE"

    ``` lua

	function EVENT_POPUP_RESPONSE(e) {
		eq.debug("popup_id " .. popup_id);
	}
    ```
## EVENT_RESPAWN

=== "EVENT_RESPAWN"

    ``` lua

	function EVENT_RESPAWN(e) {
		eq.debug("option " .. option);
		eq.debug("resurrect " .. resurrect);
	}
    ```
## EVENT_SAY

=== "EVENT_SAY"

    ``` lua

	function EVENT_SAY(e) {
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
## EVENT_TARGET_CHANGE

=== "EVENT_TARGET_CHANGE"

    ``` lua

	function EVENT_TARGET_CHANGE(e) {
	}
    ```
## EVENT_TASK_COMPLETE

=== "EVENT_TASK_COMPLETE"

    ``` lua

	function EVENT_TASK_COMPLETE(e) {
		eq.debug("count " .. count);
		eq.debug("activity_id " .. activity_id);
		eq.debug("task_id " .. task_id);
	}
    ```
## EVENT_TASK_FAIL

=== "EVENT_TASK_FAIL"

    ``` lua

	function EVENT_TASK_FAIL(e) {
		eq.debug("task_id " .. task_id);
	}
    ```
## EVENT_TASK_STAGE_COMPLETE

=== "EVENT_TASK_STAGE_COMPLETE"

    ``` lua

	function EVENT_TASK_STAGE_COMPLETE(e) {
		eq.debug("task_id " .. task_id);
		eq.debug("activity_id " .. activity_id);
	}
    ```
## EVENT_TASK_UPDATE

=== "EVENT_TASK_UPDATE"

    ``` lua

	function EVENT_TASK_UPDATE(e) {
		eq.debug("count " .. count);
		eq.debug("activity_id " .. activity_id);
		eq.debug("task_id " .. task_id);
	}
    ```
## EVENT_TEST_BUFF

=== "EVENT_TEST_BUFF"

    ``` lua

	function EVENT_TEST_BUFF(e) {
	}
    ```
## EVENT_TIMER

=== "EVENT_TIMER"

    ``` lua

	function EVENT_TIMER(e) {
		eq.debug("timer " .. timer);
	}
    ```
## EVENT_UNHANDLED_OPCODE

=== "EVENT_UNHANDLED_OPCODE"

    ``` lua

	function EVENT_UNHANDLED_OPCODE(e) {
		eq.debug("packet " .. packet);
		eq.debug("connecting " .. connecting);
	}
    ```
## EVENT_USE_SKILL

=== "EVENT_USE_SKILL"

    ``` lua

	function EVENT_USE_SKILL(e) {
		eq.debug("skill_id " .. skill_id);
		eq.debug("skill_level " .. skill_level);
	}
    ```
## EVENT_WARP

=== "EVENT_WARP"

    ``` lua

	function EVENT_WARP(e) {
		eq.debug("from_x " .. from_x);
		eq.debug("from_y " .. from_y);
		eq.debug("from_z " .. from_z);
	}
    ```
## EVENT_ZONE

=== "EVENT_ZONE"

    ``` lua

	function EVENT_ZONE(e) {
		eq.debug("zone_id " .. zone_id);
	}
    ```
