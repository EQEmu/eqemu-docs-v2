## EVENT_SPELL_EFFECT_BUFF_TIC_CLIENT

=== "EVENT_SPELL_EFFECT_BUFF_TIC_CLIENT"

    ``` lua

	function EVENT_SPELL_EFFECT_BUFF_TIC_CLIENT(e) {
		eq.debug("target " .. target);
		eq.debug("spell_id " .. spell_id);
		eq.debug("caster_id " .. caster_id);
		eq.debug("tics_remaining " .. tics_remaining);
		eq.debug("caster_level " .. caster_level);
		eq.debug("buff_slot " .. buff_slot);
		eq.debug("spell " .. spell);
	}
    ```
## EVENT_SPELL_EFFECT_CLIENT

=== "EVENT_SPELL_EFFECT_CLIENT"

    ``` lua

	function EVENT_SPELL_EFFECT_CLIENT(e) {
		eq.debug("target " .. target);
		eq.debug("spell_id " .. spell_id);
		eq.debug("caster_id " .. caster_id);
		eq.debug("tics_remaining " .. tics_remaining);
		eq.debug("caster_level " .. caster_level);
		eq.debug("buff_slot " .. buff_slot);
		eq.debug("spell " .. spell);
	}
    ```
## EVENT_SPELL_EFFECT_TRANSLOCATE_COMPLETE

=== "EVENT_SPELL_EFFECT_TRANSLOCATE_COMPLETE"

    ``` lua

	function EVENT_SPELL_EFFECT_TRANSLOCATE_COMPLETE(e) {
		eq.debug("target " .. target);
	}
    ```
## EVENT_SPELL_FADE

=== "EVENT_SPELL_FADE"

    ``` lua

	function EVENT_SPELL_FADE(e) {
		eq.debug("target " .. target);
		eq.debug("spell_id " .. spell_id);
		eq.debug("caster_id " .. caster_id);
		eq.debug("tics_remaining " .. tics_remaining);
		eq.debug("caster_level " .. caster_level);
		eq.debug("buff_slot " .. buff_slot);
		eq.debug("spell " .. spell);
	}
    ```
