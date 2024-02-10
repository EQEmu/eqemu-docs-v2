!!! info end

    Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl){:target="perl_event"} for latest definitions and Quest examples

## EVENT_SPELL_EFFECT_BOT

``` perl

sub EVENT_SPELL_EFFECT_BOT {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("tics_remaining " . $tics_remaining);
	quest::debug("caster_level " . $caster_level);
	quest::debug("buff_slot " . $buff_slot);
	quest::debug("spell " . $spell);
}
```
## EVENT_SPELL_EFFECT_BUFF_TIC_BOT

``` perl

sub EVENT_SPELL_EFFECT_BUFF_TIC_BOT {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("tics_remaining " . $tics_remaining);
	quest::debug("caster_level " . $caster_level);
	quest::debug("buff_slot " . $buff_slot);
	quest::debug("spell " . $spell);
}
```
## EVENT_SPELL_EFFECT_BUFF_TIC_CLIENT

``` perl

sub EVENT_SPELL_EFFECT_BUFF_TIC_CLIENT {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("tics_remaining " . $tics_remaining);
	quest::debug("caster_level " . $caster_level);
	quest::debug("buff_slot " . $buff_slot);
	quest::debug("spell " . $spell);
}
```
## EVENT_SPELL_EFFECT_BUFF_TIC_NPC

``` perl

sub EVENT_SPELL_EFFECT_BUFF_TIC_NPC {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("tics_remaining " . $tics_remaining);
	quest::debug("caster_level " . $caster_level);
	quest::debug("buff_slot " . $buff_slot);
	quest::debug("spell " . $spell);
}
```
## EVENT_SPELL_EFFECT_CLIENT

``` perl

sub EVENT_SPELL_EFFECT_CLIENT {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("tics_remaining " . $tics_remaining);
	quest::debug("caster_level " . $caster_level);
	quest::debug("buff_slot " . $buff_slot);
	quest::debug("spell " . $spell);
}
```
## EVENT_SPELL_EFFECT_NPC

``` perl

sub EVENT_SPELL_EFFECT_NPC {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("tics_remaining " . $tics_remaining);
	quest::debug("caster_level " . $caster_level);
	quest::debug("buff_slot " . $buff_slot);
	quest::debug("spell " . $spell);
}
```
## EVENT_SPELL_EFFECT_TRANSLOCATE_COMPLETE

``` perl

sub EVENT_SPELL_EFFECT_TRANSLOCATE_COMPLETE {
}
```
## EVENT_SPELL_FADE

``` perl

sub EVENT_SPELL_FADE {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("tics_remaining " . $tics_remaining);
	quest::debug("caster_level " . $caster_level);
	quest::debug("buff_slot " . $buff_slot);
	quest::debug("spell " . $spell);
}
```
