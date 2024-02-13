!!! info end

    Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua){:target="perl_event"} for latest definitions and Quest examples

## EVENT_SPELL_EFFECT_BUFF_TIC_CLIENT

``` lua

function EVENT_SPELL_EFFECT_BUFF_TIC_CLIENT(e) {
	eq.debug("target " .. e.target);
	eq.debug("spell_id " .. e.spell_id);
	eq.debug("caster_id " .. e.caster_id);
	eq.debug("tics_remaining " .. e.tics_remaining);
	eq.debug("caster_level " .. e.caster_level);
	eq.debug("buff_slot " .. e.buff_slot);
	eq.debug("spell " .. e.spell);
}
```
## EVENT_SPELL_EFFECT_CLIENT

``` lua

function EVENT_SPELL_EFFECT_CLIENT(e) {
	eq.debug("target " .. e.target);
	eq.debug("spell_id " .. e.spell_id);
	eq.debug("caster_id " .. e.caster_id);
	eq.debug("tics_remaining " .. e.tics_remaining);
	eq.debug("caster_level " .. e.caster_level);
	eq.debug("buff_slot " .. e.buff_slot);
	eq.debug("spell " .. e.spell);
}
```
## EVENT_SPELL_EFFECT_TRANSLOCATE_COMPLETE

``` lua

function EVENT_SPELL_EFFECT_TRANSLOCATE_COMPLETE(e) {
	eq.debug("target " .. e.target);
}
```
## EVENT_SPELL_FADE

``` lua

function EVENT_SPELL_FADE(e) {
	eq.debug("target " .. e.target);
	eq.debug("spell_id " .. e.spell_id);
	eq.debug("caster_id " .. e.caster_id);
	eq.debug("tics_remaining " .. e.tics_remaining);
	eq.debug("caster_level " .. e.caster_level);
	eq.debug("buff_slot " .. e.buff_slot);
	eq.debug("spell " .. e.spell);
}
```
