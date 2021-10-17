# Zone Header Identifiers

The zone header contains the zone type, fog data, safe coordinates, max Z, gravity, time type, weather data, sky, underworld, minclip and maxclip, and suspend buffs flag.

To modify these values you can use the following method.

{% tabs %}
{% tab title="Perl" %}
```perl
quest::UpdateZoneHeader(identifier, value);
```
{% endtab %}

{% tab title="Lua" %}
```lua
eq.update_zone_header(identifier, value);
```
{% endtab %}
{% endtabs %}

The values are as follows.

| Identifier | Notes |
| :--- | :--- |
| ztype | [Zone Type](zone-types.md) |
| fog\_red | Fog's Red Value \(0 to 255\) |
| fog\_green | Fog's Green Value \(0 to 255\) |
| fog\_blue | Fog's Blue Value \(0 to 255\) |
| fog\_minclip | Fog's Minimum Clipping Distance |
| fog\_maxclip | Fog's Maximum Clipping Distance |
| fog\_density | Fog's Density \(0.0 to 1.0\) |
| gravity | Zone Gravity \(0.4 Usually\) |
| time\_type | Zone Time Type |
| rain\_chance | Rain Chance \(0 to 100\) |
| rain\_duration | Rain Duration \(0 to 100\) |
| snow\_chance | Snow Chance  \(0 to 100\) |
| snow\_duration | Snow Duration \(0 to 100\) |
| sky | Zone Sky |
| safe\_x | Zone Safe X Coordinate |
| safe\_y | Zone Safe Y Coordinate |
| safe\_z | Zone Safe Z Coordinate |
| max\_z | Zone Max Z Coordinate |
| underworld | Zone Underworld Z Coordinate |
| minclip | Zone Minimum Clipping Distance |
| maxclip | Zone Maximum Clipping Distance |
| suspendbuffs | Zone Suspend Buffs Flag \(0 = Off, 1 = On\) |

