!!! info end

    Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl){:target="perl_event"} for latest definitions and Quest examples

    Last generated 2023.07.15

## 

``` perl

sub  {
}
```
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
}
```
## EVENT_CAST_BEGIN

``` perl

sub EVENT_CAST_BEGIN {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("caster_level " . $caster_level);
}
```
## EVENT_CAST_ON

``` perl

sub EVENT_CAST_ON {
	quest::debug("spell_id " . $spell_id);
	quest::debug("caster_id " . $caster_id);
	quest::debug("caster_level " . $caster_level);
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
}
```
## EVENT_DEATH

``` perl

sub EVENT_DEATH {
	quest::debug("killer_id " . $killer_id);
	quest::debug("killer_damage " . $killer_damage);
	quest::debug("killer_spell " . $killer_spell);
	quest::debug("killer_skill " . $killer_skill);
	quest::debug("killed_corpse_id " . $killed_corpse_id);
	quest::debug("killed_entity_id " . $killed_entity_id);
	quest::debug("killed_bot_id " . $killed_bot_id);
	quest::debug("killed_npc_id " . $killed_npc_id);
	quest::debug("killed_x " . $killed_x);
	quest::debug("killed_y " . $killed_y);
	quest::debug("killed_z " . $killed_z);
	quest::debug("killed_h " . $killed_h);
}
```
## EVENT_DEATH_COMPLETE

``` perl

sub EVENT_DEATH_COMPLETE {
	quest::debug("killer_id " . $killer_id);
	quest::debug("killer_damage " . $killer_damage);
	quest::debug("killer_spell " . $killer_spell);
	quest::debug("killer_skill " . $killer_skill);
	quest::debug("killed_corpse_id " . $killed_corpse_id);
	quest::debug("killed_entity_id " . $killed_entity_id);
	quest::debug("killed_bot_id " . $killed_bot_id);
	quest::debug("killed_npc_id " . $killed_npc_id);
	quest::debug("killed_x " . $killed_x);
	quest::debug("killed_y " . $killed_y);
	quest::debug("killed_z " . $killed_z);
	quest::debug("killed_h " . $killed_h);
}
```
## EVENT_DESPAWN

``` perl

sub EVENT_DESPAWN {
	quest::debug("despawned_entity_id " . $despawned_entity_id);
	quest::debug("despawned_bot_id " . $despawned_bot_id);
	quest::debug("despawned_npc_id " . $despawned_npc_id);
}
```
## EVENT_ITEM

``` perl

sub EVENT_ITEM {
	quest::debug("QuestItem " . $QuestItem);
	quest::debug("copper " . $copper);
	quest::debug("silver " . $silver);
	quest::debug("gold " . $gold);
	quest::debug("platinum " . $platinum);
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
## EVENT_UNEQUIP_ITEM_BOT

``` perl

sub EVENT_UNEQUIP_ITEM_BOT {
	quest::debug("item_id " . $item_id);
	quest::debug("item_quantity " . $item_quantity);
	quest::debug("slot_id " . $slot_id);
	quest::debug("item " . $item);
}
```
## EVENT_USE_SKILL

``` perl

sub EVENT_USE_SKILL {
	quest::debug("skill_id " . $skill_id);
	quest::debug("skill_level " . $skill_level);
}
```
