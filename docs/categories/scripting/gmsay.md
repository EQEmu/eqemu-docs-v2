# GMSay

## Examples

{% tabs %}
{% tab title="Perl" %}
```perl
quest::gmsay("Test") # Will send white text to all players in the current zone with an admin status of >= 80
quest::gmsay("Test", 13) # Will send red text to all players in the current zone with an admin status of >= 80
quest::gmsay("Test", 13, 1) # Will send red text to all players with an admin status of >= 80 in all zones
quest::gmsay("Test", 13, 1, 30) # Will send red text to all players with an admin status of >= 80 in all zones who are in the guild with an ID of 30
quest::gmsay("Test", 13, 1, 30, 0) # Will send red text to all players (minstatus == 0) in all zones who are in the guild with an ID of 30
```
{% endtab %}
{% endtabs %}

