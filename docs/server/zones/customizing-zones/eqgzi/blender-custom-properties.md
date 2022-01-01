# Custom Properties in EQGZI

Custom Properties are defined inside Blender. It is a way to convey custom EverQuest related information inside a .blend file to be set to EverQuest

## Material Custom Properties

- **fx**: Shader to use on this material. Examples of shaders can be seen in the [Shader List](shader-list.md)
- **e_* properties**: Each of these have a property value that are a type followed by a value, for example, the custom property `e_TextureDiffuse0` may have the value `2 ra_watertest_c_01.dds` set on it, defining a type of 2, and value of ra_watertest_c_01.dds.

## Object Custom Properties

### Object Region Custom Properties

A region is a cubic area that creates unique behavior within it. A classic region example is water or lava.

Note that all regions MUST have a prefix inside their object's name, the object name prefixes are: AWT_: water, ALV_: lava, APK_: pvp, ATP_: zoneline, ASL_: ice, APV_: generic

Note also that all regions MUST be an empty cube. (Add, Empty -> Cube)

- **unknowna**: with ATP_: zoneidnumber (unverified)
- **unknownb**: TODO: identify
- **unknownc**: TODO: identify

### Object Emitter Custom Properties

Particle effects are called emitters in EQ, and are defined by a zone_EnvironmentEmitter.txt file.

- **emit_id**: value based on [Environment Emitters](/server/zones/environment-emitters/)
- **emit_duration**: how long in ms an emitter should occur for. By default 90000000

### Object Sound Custom Properties

Sound custom properties impact the zone.emt file that tells the zone how to emit sounds, both one shots, and ambient sounds.

- **sound**: file name of a sound, default none.wav
- **sound_active**: when the sound should be played, default 0 (always). Other options: 1 (daytime), 2 (nighttime)
- **sound_volume**: loudness of a sound, default 1.0 (max volume), range of 0.0 to 1.0
- **sound_fade_in**: fade in in ms for a sound, default 0 (never fade)
- **sound_fade_out**: fade out in ms for a sound, default 0 (never fade)
- **sound_type**: loop parameter, default 0 (constantly), Other options: 1 (delayed repeat)
- **sound_radius**: when person is this radius from the origin play at full sound, default 15.0 (15m range)
- **sound_distance**: max distance from origin the sound can be heard, default 50.0 (50m range)
- **sound_rand_distance**: distance in meters the sound can be played at randomly, default 0
- **sound_trigger_range**: distance from origin for a player to trigger the sound being played, default 50.0 (50m range)
- **sound_min_repeat_delay**: minimum delay before the sound repeats playback in ms, default 0
- **sound_max_repeat_delay**: maximum delay before the sound repeats playback in ms, default 0
- **sound_xmi_index**: xmi is a classic playback system, typically just use, default 0
- **sound_echo**: echo level, default 0, range 0 to 1.0
- **sound_env_toggle**: set whether playback can be controlled with environment sounds in options window, default 1

### Object Spawn Custom Properties

This is an optional feature added to assist with zone generation. When you define these custom properties on any object, inside your sql\ subfolder sql files will generate based on defined information below. Each field is what you'd expect in a database perspective


- **spawn2_id**:
- **spawn2_pathgrid**:
- **spawn2_respawntime**:
- **spawn2_spawngroupid**:
- **spawn2_variance**:
- **spawn2_version**:

- **spawngroup_id**: Only required field to generate a spawngroup
- **spawngroup_delay**: 
- **spawngroup_despawn**: 
- **spawngroup_despawn_timer**: 
- **spawngroup_dist**: 
- **spawngroup_max_x**: 
- **spawngroup_max_y**: 
- **spawngroup_min_x**: 
- **spawngroup_min_y**: 
- **spawngroup_mindelay**: 
- **spawngroup_name**: 
- **spawngroup_spawn_limit**: 
- **spawngroup_sql**: 
- **spawngroup_wp_spawns**: 
- **spawngroup_wp_spawns**: 
