## EVENT_ENCOUNTER_LOAD

=== "EVENT_ENCOUNTER_LOAD"

    ``` lua

	function EVENT_ENCOUNTER_LOAD(e) {
		eq.debug("encounter " .. encounter);
		eq.debug("data " .. data);
	}
    ```
## EVENT_ENCOUNTER_UNLOAD

=== "EVENT_ENCOUNTER_UNLOAD"

    ``` lua

	function EVENT_ENCOUNTER_UNLOAD(e) {
		eq.debug("data " .. data);
	}
    ```
## EVENT_TIMER

=== "EVENT_TIMER"

    ``` lua

	function EVENT_TIMER(e) {
		eq.debug("timer " .. timer);
	}
    ```
