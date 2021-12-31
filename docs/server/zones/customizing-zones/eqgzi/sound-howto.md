# Information relevant to sound emitters in zone in .emt files.

  - sound: file name of a sound, default none.wav
  - sound_active: when the sound should be played, default 0 (always). Other options: 1 (daytime), 2 (nighttime)
  - sound_volume: loudness of a sound, default 1.0 (max volume), range of 0.0 to 1.0
  - sound_fade_in: fade in in ms for a sound, default 0 (never fade)
  - sound_fade_out: fade out in ms for a sound, default 0 (never fade)
  - sound_type: loop parameter, default 0 (constantly), Other options: 1 (delayed repeat)
  - sound_radius: when person is this radius from the origin play at full sound, default 15.0 (15m range)
  - sound_distance: max distance from origin the sound can be heard, default 50.0 (50m range)
  - sound_rand_distance: distance in meters the sound can be played at randomly, default 0
  - sound_trigger_range: distance from origin for a player to trigger the sound being played, default 50.0 (50m range)
  - sound_min_repeat_delay: minimum delay before the sound repeats playback in ms, default 0
  - sound_max_repeat_delay: maximum delay before the sound repeats playback in ms, default 0
  - sound_xmi_index: xmi is a classic playback system, typically just use, default 0
  - sound_echo: echo level, default 0, range 0 to 1.0
  - sound_env_toggle: set whether playback can be controlled with environment sounds in options window, default 1

How these varying fields go in is something like:

```
;?,SoundFile (wav=sound mp3/xmi=music),Unknown (0=OK 1=OK),WhenActive (0=Always 1=Daytime 2=Nighttime),Volume (1.0 = 100%),FadeInMS,FadeOutMS,WavLoopType (0=Constant 1=Delayed Repeat),X,Y,Z,WavFullVolRadius,WavMaxAudibleDist,N onZero = RandomizeLocation,ActivationRange,MinRepeatDelay,M axRepeatDelay,xmiIndex,EchoLevel (50 = Max),IsEnvSound (for option toggle)
```

A few examples (yanked from innothuleb.emt):

```
2,bix_hit.wav,1,0,0.30,0,0,1,71.33,-937.87,13.91,40.00,50.00,0.00,50.00,2500,7500,0,1.00,1,0
2,darkwds1.wav,0,0,0.30,0,0,1,52.10,30.56,9.00,10000.00,10000.00,0.00,10000.00,0,0,0,1.00,1,0
2,jungle_lp.wav,0,0,0.30,0,0,1,60.11,19.28,8.87,7000.00,10000.00,0.00,10000.00,0,0,0,1.00,1,0
2,nightime_background_lp.wav,0,0,0.75,0,0,1,41.61,43.25,7.15,7000.00,10000.00,0.00,10000.00,0,0,0,1.00,1,0
2,fire_torch01_lp.wav,0,0,5.00,0,0,1,-2154.73,-921.32,0.91,10.00,75.00,0.00,75.00,0,0,0,1.00,1,0
```
