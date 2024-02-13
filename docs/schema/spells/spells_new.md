# spells_new

## Relationships

```mermaid
erDiagram
    spells_new {
        int id
        int descnum
        int effectdescnum
        int effectdescnum2
        int typedescnum
        varchar teleport_zone
    }
    db_str {
        int id
    }
    auras {
        int npc_type
        int spell_id
    }
    blocked_spells {
        mediumintunsigned spellid
        int zoneid
    }
    bot_pets {
        varchar bot_id
        varchar pets_index
        varchar spell_id
    }
    bot_pet_buffs {
        varchar pets_index
        varchar spell_id
    }
    spells_new ||--o{ db_str : "One-to-One"
    spells_new ||--o{ db_str : "One-to-One"
    spells_new ||--o{ db_str : "One-to-One"
    spells_new ||--o{ db_str : "One-to-One"
    spells_new ||--o{ auras : "One-to-One"
    spells_new ||--o{ blocked_spells : "Has-Many"
    spells_new ||--o{ bot_pets : "Has-Many"
    spells_new ||--o{ bot_pet_buffs : "Has-Many"


```

```mermaid
erDiagram
    spells_new {
        int id
        int descnum
        int effectdescnum
        int effectdescnum2
        int typedescnum
        varchar teleport_zone
    }
    bot_spells_entries {
        varchar type
        varchar npc_spells_id
        varchar spell_id
        varchar spellid
    }
    character_auras {
        int id
        int spell_id
    }
    character_buffs {
        intunsigned character_id
        smallintunsigned spell_id
    }
    character_disciplines {
        intunsigned id
        smallintunsigned disc_id
    }
    character_pet_buffs {
        int char_id
        int spell_id
    }
    spells_new ||--o{ bot_spells_entries : "Has-Many"
    spells_new ||--o{ character_auras : "Has-Many"
    spells_new ||--o{ character_buffs : "Has-Many"
    spells_new ||--o{ character_disciplines : "Has-Many"
    spells_new ||--o{ character_pet_buffs : "Has-Many"


```

```mermaid
erDiagram
    spells_new {
        int id
        int descnum
        int effectdescnum
        int effectdescnum2
        int typedescnum
        varchar teleport_zone
    }
    character_pet_info {
        int char_id
        int spell_id
    }
    character_spells {
        intunsigned id
        smallintunsigned spell_id
    }
    damageshieldtypes {
        intunsigned spellid
    }
    items {
        int id
        int book
        varchar name
        int recasttype
        int icon
        mediumint bardeffect
        int clickeffect
        int focuseffect
        int proceffect
        int scrolleffect
        int worneffect
    }
    ldon_trap_templates {
        intunsigned id
        smallintunsigned spell_id
    }
    spells_new ||--o{ character_pet_info : "Has-Many"
    spells_new ||--o{ character_spells : "Has-Many"
    spells_new ||--o{ damageshieldtypes : "Has-Many"
    spells_new ||--o{ items : "Has-Many"
    spells_new ||--o{ items : "Has-Many"
    spells_new ||--o{ items : "Has-Many"
    spells_new ||--o{ items : "Has-Many"
    spells_new ||--o{ items : "Has-Many"
    spells_new ||--o{ items : "Has-Many"
    spells_new ||--o{ ldon_trap_templates : "Has-Many"


```

```mermaid
erDiagram
    spells_new {
        int id
        int descnum
        int effectdescnum
        int effectdescnum2
        int typedescnum
        varchar teleport_zone
    }
    merc_spell_list_entries {
        varchar merc_spell_list_id
        varchar spell_id
        varchar stance_id
    }
    npc_spells_entries {
        int npc_spells_id
        smallintunsigned spellid
    }
    spell_buckets {
        varchar key
        bigintunsigned spellid
    }
    spell_globals {
        varchar qglobal
        int spellid
    }
    spells_new ||--o{ merc_spell_list_entries : "Has-Many"
    spells_new ||--o{ npc_spells_entries : "Has-Many"
    spells_new ||--o{ spell_buckets : "Has-Many"
    spells_new ||--o{ spell_globals : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | descnum | [db_str](../../schema/client-files/db_str.md) | id |
| One-to-One | effectdescnum | [db_str](../../schema/client-files/db_str.md) | id |
| One-to-One | effectdescnum2 | [db_str](../../schema/client-files/db_str.md) | id |
| One-to-One | typedescnum | [db_str](../../schema/client-files/db_str.md) | id |
| One-to-One | id | [auras](../../schema/spells/auras.md) | spell_id |
| Has-Many | id | [blocked_spells](../../schema/spells/blocked_spells.md) | spellid |
| Has-Many | id | [bot_pets](../../schema/bots/bot_pets.md) | spell_id |
| Has-Many | id | [bot_pet_buffs](../../schema/bots/bot_pet_buffs.md) | spell_id |
| Has-Many | id | [bot_spells_entries](../../schema/bots/bot_spells_entries.md) | spellid |
| Has-Many | id | [character_auras](../../schema/characters/character_auras.md) | spell_id |
| Has-Many | id | [character_buffs](../../schema/characters/character_buffs.md) | spell_id |
| Has-Many | id | [character_disciplines](../../schema/characters/character_disciplines.md) | disc_id |
| Has-Many | id | [character_pet_buffs](../../schema/characters/character_pet_buffs.md) | spell_id |
| Has-Many | id | [character_pet_info](../../schema/characters/character_pet_info.md) | spell_id |
| Has-Many | id | [character_spells](../../schema/characters/character_spells.md) | spell_id |
| Has-Many | id | [damageshieldtypes](../../schema/spells/damageshieldtypes.md) | spellid |
| Has-Many | id | [items](../../schema/items/items.md) | bardeffect |
| Has-Many | id | [items](../../schema/items/items.md) | clickeffect |
| Has-Many | id | [items](../../schema/items/items.md) | focuseffect |
| Has-Many | id | [items](../../schema/items/items.md) | proceffect |
| Has-Many | id | [items](../../schema/items/items.md) | scrolleffect |
| Has-Many | id | [items](../../schema/items/items.md) | worneffect |
| Has-Many | id | [ldon_trap_templates](../../schema/traps/ldon_trap_templates.md) | spell_id |
| Has-Many | id | [merc_spell_list_entries](../../schema/mercenaries/merc_spell_list_entries.md) | spell_id |
| Has-Many | id | [npc_spells_entries](../../schema/npcs/npc_spells_entries.md) | spellid |
| Has-Many | id | [spell_buckets](../../schema/spells/spell_buckets.md) | spellid |
| Has-Many | id | [spell_globals](../../schema/spells/spell_globals.md) | spellid |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Spell Identifier |
| name | varchar | Name |
| player_1 | varchar | Player_1 |
| teleport_zone | varchar | The zone you are teleporting to or the [NPC Name](../../schema/npcs/npc_types.md) you want to spawn. |
| you_cast | varchar | The message sent to others when you cast the spell. |
| other_casts | varchar | The message seen when someone around you casts the spell. |
| cast_on_you | varchar | The message received when the spell is cast on you. |
| cast_on_other | varchar | The message recieved when the spell is cast on another. |
| spell_fades | varchar | The message recieved when the spell fades. |
| range | int | Range |
| aoerange | int | Area of Effect Range |
| pushback | int | Push Back |
| pushup | int | Push Up |
| cast_time | int | Cast Time in Milliseconds |
| recovery_time | int | Recovery Time in Seconds |
| recast_time | int | Recast Time in Seconds |
| buffdurationformula | int | [Buff Duration Formula](../../../../server/spells/buff-duration-formulas) |
| buffduration | int | Buff Duration |
| AEDuration | int | Area of Effect Duration |
| mana | int | Mana Cost |
| effect_base_value1 | int | Effect Base Value 1 |
| effect_base_value2 | int | Effect Base Value 2 |
| effect_base_value3 | int | Effect Base Value 3 |
| effect_base_value4 | int | Effect Base Value 4 |
| effect_base_value5 | int | Effect Base Value 5 |
| effect_base_value6 | int | Effect Base Value 6 |
| effect_base_value7 | int | Effect Base Value 7 |
| effect_base_value8 | int | Effect Base Value 8 |
| effect_base_value9 | int | Effect Base Value 9 |
| effect_base_value10 | int | Effect Base Value 10 |
| effect_base_value11 | int | Effect Base Value 11 |
| effect_base_value12 | int | Effect Base Value 12 |
| effect_limit_value1 | int | Effect Limit Value 1 |
| effect_limit_value2 | int | Effect Limit Value 2 |
| effect_limit_value3 | int | Effect Limit Value 3 |
| effect_limit_value4 | int | Effect Limit Value 4 |
| effect_limit_value5 | int | Effect Limit Value 5 |
| effect_limit_value6 | int | Effect Limit Value 6 |
| effect_limit_value7 | int | Effect Limit Value 7 |
| effect_limit_value8 | int | Effect Limit Value 8 |
| effect_limit_value9 | int | Effect Limit Value 9 |
| effect_limit_value10 | int | Effect Limit Value 10 |
| effect_limit_value11 | int | Effect Limit Value 11 |
| effect_limit_value12 | int | Effect Limit Value 12 |
| max1 | int | Max 1 |
| max2 | int | Max 2 |
| max3 | int | Max 3 |
| max4 | int | Max 4 |
| max5 | int | Max 5 |
| max6 | int | Max 6 |
| max7 | int | Max 7 |
| max8 | int | Max 8 |
| max9 | int | Max 9 |
| max10 | int | Max 10 |
| max11 | int | Max 11 |
| max12 | int | Max 12 |
| icon | int | Icon |
| memicon | int | Memmed Icon |
| components1 | int | [Item Identifier](../../schema/items/items.md) |
| components2 | int | [Item Identifier](../../schema/items/items.md) |
| components3 | int | [Item Identifier](../../schema/items/items.md) |
| components4 | int | [Item Identifier](../../schema/items/items.md) |
| component_counts1 | int | Component Count 1 |
| component_counts2 | int | Component Count 2 |
| component_counts3 | int | Component Count 3 |
| component_counts4 | int | Component Count 4 |
| NoexpendReagent1 | int | If it is a number between 1-4 it means component number 1-4 is a focus and not to expend it. If it is a valid item ID it means this item is a focus as well. |
| NoexpendReagent2 | int | If it is a number between 1-4 it means component number 1-4 is a focus and not to expend it. If it is a valid item ID it means this item is a focus as well. |
| NoexpendReagent3 | int | If it is a number between 1-4 it means component number 1-4 is a focus and not to expend it. If it is a valid item ID it means this item is a focus as well. |
| NoexpendReagent4 | int | If it is a number between 1-4 it means component number 1-4 is a focus and not to expend it. If it is a valid item ID it means this item is a focus as well. |
| formula1 | int | [Formula 1](../../../../server/spells/base-value-formulas) |
| formula2 | int | [Formula 2](../../../../server/spells/base-value-formulas) |
| formula3 | int | [Formula 3](../../../../server/spells/base-value-formulas) |
| formula4 | int | [Formula 4](../../../../server/spells/base-value-formulas) |
| formula5 | int | [Formula 5](../../../../server/spells/base-value-formulas) |
| formula6 | int | [Formula 6](../../../../server/spells/base-value-formulas) |
| formula7 | int | [Formula 7](../../../../server/spells/base-value-formulas) |
| formula8 | int | [Formula 8](../../../../server/spells/base-value-formulas) |
| formula9 | int | [Formula 9](../../../../server/spells/base-value-formulas) |
| formula10 | int | [Formula 10](../../../../server/spells/base-value-formulas) |
| formula11 | int | [Formula 11](../../../../server/spells/base-value-formulas) |
| formula12 | int | [Formula 12](../../../../server/spells/base-value-formulas) |
| LightType | int | Light Type |
| goodEffect | int | Good Effect: 0 = Detrimental, 1 = Beneficial, 2 = Beneficial Group Only |
| Activated | int | Activated |
| resisttype | int | [Resist Type](../../../../server/spells/resist-types) |
| effectid1 | int | [Effect Identifier 1](../../../../server/spells/spell-effect-ids) |
| effectid2 | int | [Effect Identifier 2](../../../../server/spells/spell-effect-ids) |
| effectid3 | int | [Effect Identifier 3](../../../../server/spells/spell-effect-ids) |
| effectid4 | int | [Effect Identifier 4](../../../../server/spells/spell-effect-ids) |
| effectid5 | int | [Effect Identifier 5](../../../../server/spells/spell-effect-ids) |
| effectid6 | int | [Effect Identifier 6](../../../../server/spells/spell-effect-ids) |
| effectid7 | int | [Effect Identifier 7](../../../../server/spells/spell-effect-ids) |
| effectid8 | int | [Effect Identifier 8](../../../../server/spells/spell-effect-ids) |
| effectid9 | int | [Effect Identifier 9](../../../../server/spells/spell-effect-ids) |
| effectid10 | int | [Effect Identifier 10](../../../../server/spells/spell-effect-ids) |
| effectid11 | int | [Effect Identifier 11](../../../../server/spells/spell-effect-ids) |
| effectid12 | int | [Effect Identifier 12](../../../../server/spells/spell-effect-ids) |
| targettype | int | [Target Type](../../../../server/spells/target-types) |
| basediff | int | Base Difficult Fizzle Adjustment |
| skill | int | [Skill Identifier](../../../../server/player/skills) |
| zonetype | int | [Zone Type](../../../../server/zones/zone-types) |
| EnvironmentType | int | [Environment Type](../../../../server/spells/environment-types) |
| TimeOfDay | int | [Time of Day Type](../../../../server/spells/time-of-day-types) |
| classes1 | int | Required Level for Warrior |
| classes2 | int | Required Level for Cleric |
| classes3 | int | Required Level for Paladin |
| classes4 | int | Required Level for Ranger |
| classes5 | int | Required Level for Shadow Knight |
| classes6 | int | Required Level for Druid |
| classes7 | int | Required Level for Monk |
| classes8 | int | Required Level for Bard |
| classes9 | int | Required Level for Rogue |
| classes10 | int | Required Level for Shaman |
| classes11 | int | Required Level for Necromancer |
| classes12 | int | Required Level for Wizard |
| classes13 | int | Required Level for Magician |
| classes14 | int | Required Level for Enchanter |
| classes15 | int | Required Level for Beastlord |
| classes16 | int | Required Level for Berserker |
| CastingAnim | int | Casting Animation |
| TargetAnim | int | Target Animation |
| TravelType | int | Travel Type (Unused) |
| SpellAffectIndex | int | Spell Affect Index |
| disallow_sit | int | Disallow Sit: 0 = False, 1 = True |
| deities0 | int | [Deity List](../../../../server/player/deity-list) |
| deities1 | int | [Deity List](../../../../server/player/deity-list) |
| deities2 | int | [Deity List](../../../../server/player/deity-list) |
| deities3 | int | [Deity List](../../../../server/player/deity-list) |
| deities4 | int | [Deity List](../../../../server/player/deity-list) |
| deities5 | int | [Deity List](../../../../server/player/deity-list) |
| deities6 | int | [Deity List](../../../../server/player/deity-list) |
| deities7 | int | [Deity List](../../../../server/player/deity-list) |
| deities8 | int | [Deity List](../../../../server/player/deity-list) |
| deities9 | int | [Deity List](../../../../server/player/deity-list) |
| deities10 | int | [Deity List](../../../../server/player/deity-list) |
| deities11 | int | [Deity List](../../../../server/player/deity-list) |
| deities12 | int | [Deity List](../../../../server/player/deity-list) |
| deities13 | int | [Deity List](../../../../server/player/deity-list) |
| deities14 | int | [Deity List](../../../../server/player/deity-list) |
| deities15 | int | [Deity List](../../../../server/player/deity-list) |
| deities16 | int | [Deity List](../../../../server/player/deity-list) |
| field142 | int | Unknown |
| field143 | int | Unknown |
| new_icon | int | New Icon |
| spellanim | int | Spell Animation |
| uninterruptable | int | Uninterruptable: 0 = False, 1 = True |
| ResistDiff | int | Resist Difference |
| dot_stacking_exempt | int | Damage Over Time Stacking Exempt: 0 = False, 1 = True |
| deleteable | int | Deleteable: 0 = False, 1 = True |
| RecourseLink | int | [Recourse Spell Identifier](spells_new.md) |
| no_partial_resist | int | No Partial Resist: 0 = False, 1 = True |
| field152 | int | Unknown |
| field153 | int | Unknown |
| short_buff_box | int | Short Buff Box: 0 = False, 1 = True |
| descnum | int | [Description Number](../../schema/client-files/db_str.md) |
| typedescnum | int | [Type Description Number](../../schema/client-files/db_str.md) |
| effectdescnum | int | [Effect Description Number](../../schema/client-files/db_str.md) |
| effectdescnum2 | int | [Effect Description Number 2](../../schema/client-files/db_str.md) |
| npc_no_los | int | NPC No Line of Sight: 0 = False, 1 = True |
| field160 | int | Unknown |
| reflectable | int | Reflectable: 0 = False, 1 = True |
| bonushate | int | Bonus Hate |
| field163 | int | Unknown |
| field164 | int | Unknown |
| ldon_trap | int | [LDoN Trap Identifier](../../schema/traps/traps.md) |
| EndurCost | int | Endurance Cost |
| EndurTimerIndex | int | Endurance Timer |
| IsDiscipline | int | Is Discipline: 0 = False, 1 = True |
| field169 | int | Unknown |
| field170 | int | Unknown |
| field171 | int | Unknown |
| field172 | int | Unknown |
| HateAdded | int | Hate Added |
| EndurUpkeep | int | Endurance Upkeep |
| numhitstype | int | [Number of Hits Type](../../../../server/spells/numhit-types) |
| numhits | int | Number of Hits |
| pvpresistbase | int | PVP Resist Base |
| pvpresistcalc | int | PVP Resist Calc |
| pvpresistcap | int | PVP Resist Cap |
| spell_category | int | [Spell Category](../../../../server/spells/spell-groups) |
| pvp_duration | int | PVP Duration |
| pvp_duration_cap | int | PVP Duration Cap |
| pcnpc_only_flag | int | PC/NPC Only Flag: 0 = Not Applicable, 1 = PCs and Mercs, 2 = NPCs |
| cast_not_standing | int | Cast Not Standing: 0 = False, 1 = True |
| can_mgb | int | Can Mass Group Buff: 0 = False, 1 = True |
| nodispell | int | No Dispell: 0 = False, 1 = True |
| npc_category | int | [NPC Spell Category Identifier](../../../../server/spells/npc-spell-categories) |
| npc_usefulness | int | NPC Usefulness |
| MinResist | int | Minimum Resistance |
| MaxResist | int | Maximum Resistance |
| viral_targets | int | Viral Targets |
| viral_timer | int | Viral Timer |
| nimbuseffect | int | Nimbus Effect |
| ConeStartAngle | int | Cone Start Angle |
| ConeStopAngle | int | Cone Stop Angle |
| sneaking | int | Sneaking: 0 = False, 1 = True |
| not_extendable | int | Not Extendable: 0 = False, 1 = True |
| field198 | int | Unknown |
| field199 | int | Unknown |
| suspendable | int | Suspendable: 0 = False, 1 = True |
| viral_range | int | Viral Range |
| songcap | int | Song Cap |
| field203 | int | Unknown |
| field204 | int | Unknown |
| no_block | int | No Block: 0 = False, 1 = True |
| field206 | int | Unknown |
| spellgroup | int | [Spell Group](../../../../server/spells/spell-groups) |
| rank | int | Rank |
| field209 | int | Unknown |
| field210 | int | Unknown |
| CastRestriction | int | [Cast Restrictions](../../../../server/spells/spell-target-restrictions) |
| allowrest | int | Allow Rest: 0 = False, 1 = True |
| InCombat | int | In Combat: 0 = False, 1 = True |
| OutofCombat | int | Out Of Combat: 0 = False, 1 = True |
| field215 | int | Unknown |
| field216 | int | Unknown |
| field217 | int | Unknown |
| aemaxtargets | int | Area of Effect Max Targets |
| maxtargets | int | Max Targets |
| field220 | int | Unknown |
| field221 | int | Unknown |
| field222 | int | Unknown |
| field223 | int | Unknown |
| persistdeath | int | Persist Death: 0 = False, 1 = True |
| field225 | int | Unknown |
| field226 | int | Unknown |
| min_dist | float | Minimum Distance |
| min_dist_mod | float | Minimum Distance Modifier |
| max_dist | float | Maximum Distance |
| max_dist_mod | float | Maximum Distance Modifier |
| min_range | int | Minimum Range |
| field232 | int | Unknown |
| field233 | int | Unknown |
| field234 | int | Unknown |
| field235 | int | Unknown |
| field236 | int | Unknown |

