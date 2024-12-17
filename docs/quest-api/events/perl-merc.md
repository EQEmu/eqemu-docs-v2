!!! info end

    Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl){:target="perl_event"} for latest definitions and Quest examples

## EVENT_AGGRO_SAY

``` perl

sub EVENT_AGGRO_SAY {
	quest::debug("data " . $data);
	quest::debug("text " . $text);
	quest::debug("langid " . $langid);
}
```
## EVENT_ATTACK

``` perl

sub EVENT_ATTACK {
}
```
## EVENT_CAST

``` perl

sub EVENT_CAST {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("caster_level " . $caster_level);
	quest::debug("target_id " . $target_id);
	quest::debug("target " . $target);
	quest::debug("spell " . $spell);
}
```
## EVENT_CAST_BEGIN

``` perl

sub EVENT_CAST_BEGIN {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("caster_level " . $caster_level);
	quest::debug("target_id " . $target_id);
	quest::debug("target " . $target);
	quest::debug("spell " . $spell);
}
```
## EVENT_CAST_ON

``` perl

sub EVENT_CAST_ON {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("caster_level " . $caster_level);
	quest::debug("target_id " . $target_id);
	quest::debug("target " . $target);
	quest::debug("spell " . $spell);
}
```
## EVENT_COMBAT

``` perl

sub EVENT_COMBAT {
	quest::debug("combat_state " . $combat_state);
}
```
## EVENT_DAMAGE_GIVEN

``` perl

sub EVENT_DAMAGE_GIVEN {
	quest::debug("entity_id " . $entity_id);
	quest::debug("damage " . $damage);
	quest::debug("spell_id " . $spell_id);
	quest::debug("skill_id " . $skill_id);
	quest::debug("is_damage_shield " . $is_damage_shield);
	quest::debug("is_avoidable " . $is_avoidable);
	quest::debug("buff_slot " . $buff_slot);
	quest::debug("is_buff_tic " . $is_buff_tic);
	quest::debug("special_attack " . $special_attack);
	quest::debug("spell " . $spell);
}
```
## EVENT_DAMAGE_TAKEN

``` perl

sub EVENT_DAMAGE_TAKEN {
	quest::debug("entity_id " . $entity_id);
	quest::debug("damage " . $damage);
	quest::debug("spell_id " . $spell_id);
	quest::debug("skill_id " . $skill_id);
	quest::debug("is_damage_shield " . $is_damage_shield);
	quest::debug("is_avoidable " . $is_avoidable);
	quest::debug("buff_slot " . $buff_slot);
	quest::debug("is_buff_tic " . $is_buff_tic);
	quest::debug("special_attack " . $special_attack);
	quest::debug("spell " . $spell);
}
```
## EVENT_DEATH

``` perl

sub EVENT_DEATH {
	quest::debug("killer_id " . $killer_id);
	quest::debug("killer_damage " . $killer_damage);
	quest::debug("killer_spell " . $killer_spell);
	quest::debug("killer_skill " . $killer_skill);
	quest::debug("killed_entity_id " . $killed_entity_id);
	quest::debug("combat_start_time " . $combat_start_time);
	quest::debug("combat_end_time " . $combat_end_time);
	quest::debug("damage_received " . $damage_received);
	quest::debug("healing_received " . $healing_received);
	quest::debug("killed_corpse_id " . $killed_corpse_id);
	quest::debug("killed_x " . $killed_x);
	quest::debug("killed_y " . $killed_y);
	quest::debug("killed_z " . $killed_z);
	quest::debug("killed_h " . $killed_h);
	quest::debug("killed_merc_id " . $killed_merc_id);
	quest::debug("killed_npc_id " . $killed_npc_id);
}
```
## EVENT_DEATH_COMPLETE

``` perl

sub EVENT_DEATH_COMPLETE {
	quest::debug("killer_id " . $killer_id);
	quest::debug("killer_damage " . $killer_damage);
	quest::debug("killer_spell " . $killer_spell);
	quest::debug("killer_skill " . $killer_skill);
	quest::debug("killed_entity_id " . $killed_entity_id);
	quest::debug("combat_start_time " . $combat_start_time);
	quest::debug("combat_end_time " . $combat_end_time);
	quest::debug("damage_received " . $damage_received);
	quest::debug("healing_received " . $healing_received);
	quest::debug("killed_corpse_id " . $killed_corpse_id);
	quest::debug("killed_x " . $killed_x);
	quest::debug("killed_y " . $killed_y);
	quest::debug("killed_z " . $killed_z);
	quest::debug("killed_h " . $killed_h);
	quest::debug("killed_merc_id " . $killed_merc_id);
	quest::debug("killed_npc_id " . $killed_npc_id);
}
```
## EVENT_DESPAWN

``` perl

sub EVENT_DESPAWN {
	quest::debug("despawned_entity_id " . $despawned_entity_id);
	quest::debug("despawned_bot_id " . $despawned_bot_id);
	quest::debug("despawned_merc_id " . $despawned_merc_id);
	quest::debug("despawned_npc_id " . $despawned_npc_id);
}
```
## EVENT_ENTITY_VARIABLE_DELETE

``` perl

sub EVENT_ENTITY_VARIABLE_DELETE {
	quest::debug("variable_name " . $variable_name);
	quest::debug("variable_value " . $variable_value);
}
```
## EVENT_ENTITY_VARIABLE_SET

``` perl

sub EVENT_ENTITY_VARIABLE_SET {
	quest::debug("variable_name " . $variable_name);
	quest::debug("variable_value " . $variable_value);
}
```
## EVENT_ENTITY_VARIABLE_UPDATE

``` perl

sub EVENT_ENTITY_VARIABLE_UPDATE {
	quest::debug("variable_name " . $variable_name);
	quest::debug("old_value " . $old_value);
	quest::debug("new_value " . $new_value);
}
```
## EVENT_NPC_SLAY

``` perl

sub EVENT_NPC_SLAY {
	quest::debug("killed " . $killed);
	quest::debug("killed_npc " . $killed_npc);
}
```
## EVENT_PAYLOAD

``` perl

sub EVENT_PAYLOAD {
	quest::debug("payload_id " . $payload_id);
	quest::debug("payload_value " . $payload_value);
}
```
## EVENT_POPUPRESPONSE

``` perl

sub EVENT_POPUPRESPONSE {
	quest::debug("popupid " . $popupid);
}
```
## EVENT_SAY

``` perl

sub EVENT_SAY {
	quest::debug("data " . $data);
	quest::debug("text " . $text);
	quest::debug("langid " . $langid);
}
```
## EVENT_SIGNAL

``` perl

sub EVENT_SIGNAL {
	quest::debug("signal " . $signal);
}
```
## EVENT_SLAY

``` perl

sub EVENT_SLAY {
}
```
## EVENT_SPAWN

``` perl

sub EVENT_SPAWN {
}
```
## EVENT_SPELL_BLOCKED

``` perl

sub EVENT_SPELL_BLOCKED {
	quest::debug("blocking_spell_id " . $blocking_spell_id);
	quest::debug("cast_spell_id " . $cast_spell_id);
	quest::debug("blocking_spell " . $blocking_spell);
	quest::debug("cast_spell " . $cast_spell);
}
```
## EVENT_TARGET_CHANGE

``` perl

sub EVENT_TARGET_CHANGE {
	quest::debug("target " . $target);
}
```
## EVENT_TIMER

``` perl

sub EVENT_TIMER {
	quest::debug("timer " . $timer);
}
```
## EVENT_TIMER_PAUSE

``` perl

sub EVENT_TIMER_PAUSE {
	quest::debug("timer " . $timer);
	quest::debug("duration " . $duration);
}
```
## EVENT_TIMER_RESUME

``` perl

sub EVENT_TIMER_RESUME {
	quest::debug("timer " . $timer);
	quest::debug("duration " . $duration);
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
## EVENT_USE_SKILL

``` perl

sub EVENT_USE_SKILL {
	quest::debug("skill_id " . $skill_id);
	quest::debug("skill_level " . $skill_level);
}
```
