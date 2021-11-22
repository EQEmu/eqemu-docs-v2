# Zone Header Identifiers

The zone header contains the zone type, fog data, safe coordinates, max Z, gravity, time type, weather data, sky, underworld, minclip and maxclip, and suspend buffs flag.

To modify these values you can use the following method.

=== "Perl"
    ```perl
    quest::UpdateZoneHeader(identifier, value);
    ```
=== "Lua"
    ```lua
    eq.update_zone_header(identifier, value);
    ```

The values are as follows.

| Identifier | Notes |
| :--- | :--- |
| ztype | [Zone Type](zone-types.md) |
| fog_red | Fog's Red Value (0 to 255) |
| fog_green | Fog's Green Value (0 to 255) |
| fog_blue | Fog's Blue Value (0 to 255) |
| fog_minclip | Fog's Minimum Clipping Distance |
| fog_maxclip | Fog's Maximum Clipping Distance |
| fog_density | Fog's Density (0.0 to 1.0) |
| gravity | Zone Gravity (0.4 Usually) |
| time_type | Zone Time Type |
| rain_chance | Rain Chance (0 to 100) |
| rain_duration | Rain Duration (0 to 100) |
| snow_chance | Snow Chance  (0 to 100) |
| snow_duration | Snow Duration (0 to 100) |
| sky | Zone Sky |
| safe_x | Zone Safe X Coordinate |
| safe_y | Zone Safe Y Coordinate |
| safe_z | Zone Safe Z Coordinate |
| max_z | Zone Max Z Coordinate |
| underworld | Zone Underworld Z Coordinate |
| minclip | Zone Minimum Clipping Distance |
| maxclip | Zone Maximum Clipping Distance |
| suspendbuffs | Zone Suspend Buffs Flag (0 = Off, 1 = On) |

