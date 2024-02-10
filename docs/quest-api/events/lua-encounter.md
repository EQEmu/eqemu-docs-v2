!!! info end

    Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua){:target="perl_event"} for latest definitions and Quest examples

## EVENT_ENCOUNTER_LOAD

``` lua

function EVENT_ENCOUNTER_LOAD(e) {
	eq.debug("encounter " .. e.encounter);
	eq.debug("data " .. e.data);
}
```
## EVENT_ENCOUNTER_UNLOAD

``` lua

function EVENT_ENCOUNTER_UNLOAD(e) {
	eq.debug("data " .. e.data);
}
```
## EVENT_TIMER

``` lua

function EVENT_TIMER(e) {
	eq.debug("timer " .. e.timer);
}
```
