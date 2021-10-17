# Global Methods

## `void` add\_expedition\_lockout\_all\_clients\(string expedition\_name, string event\_name, int seconds, string uuid = ""\) <a id="add-expedition-lockouts-all-clients"></a>

Adds the specified lockout to all clients in the current zone. If a client already has the lockout it's replaced with this one. If a uuid is not supplied one is generated and applied to all clients.

```perl
quest::add_expedition_lockout_all_clients("Added expedition", "Added event", 7200);
```

## `void` add\_expedition\_lockout\_by\_char\_id\(int character\_id, string expedition\_name, string event\_name, int seconds, string uuid = ""\) <a id="add-expedition-lockouts-by-char-id"></a>

Adds the specified lockout to the character with an optional specified uuid. If the client already has the lockout it's replaced with this one. If a uuid is not supplied one is generated when the lockout is added to the character

```perl
quest::add_expedition_lockout_by_char_id($client->CharacterID(), "Added expedition", "Added event", 7200);
```

## `Expedition*` get\_expedition\(\) <a id="get-expedition"></a>

Returns an Expedition object corresponding to current zone if it's a valid dynamic zone instance being used for an expedition, otherwise returns nullptr.

```perl
my $dz = quest::get_expedition();
if ($dz) {
  quest::debug(sprintf("current zone is expedition id [%d]'s dz instance", $dz->GetID()));
}
```

## `Expedition*` get\_expedition\_by\_char\_id\(int character\_id\) <a id="get-expedition-by-char-id"></a>

Returns the character's current Expedition if the character is in one, otherwise returns nullptr.

```perl
my $dz = quest::get_expedition_by_char_id($client->CharacterID());
if ($dz) {
  quest::debug(sprintf("client [%d] is in expedition id [%d]", $client->CharacterID(), $dz->GetID()));
}
```

## `Expedition*` get\_expedition\_by\_dz\_id\(int dz\_id\) <a id="get-expedition-by-dz-id"></a>

Returns the Expedition for the specified dynamic zone id if it's valid, otherwise returns nullptr.

```perl
my $dzid = 0;
my $dz = quest::get_expedition_by_char_id($client->CharacterID());
if ($dz) {
  $dzid = $dz->GetDynamicZoneID();
}

if ($dzid > 0) {
  my $expedition_from_dzid = quest::get_expedition_by_dz_id($dzid);
  if ($expedition_from_dzid) {
    quest::debug(sprintf("dz id [%d] is for expedition id [%d]", $dzid, $expedition_from_dzid->GetID()));
  }
}
```

## `Expedition*` get\_expedition\_by\_zone\_instance\(int zone\_id, int instance\_id\) <a id="get-expedition-by-zone-instance"></a>

Returns the Expedition for the specified zone and instance id if it's a valid dynamic zone instance, otherwise returns nullptr.

```perl
# assuming character is in an anguish expedition
my $zoneid = quest::GetZoneID("anguish");
my $instanceid = quest::GetInstanceIDByCharID("anguish", 0, $client->CharacterID());

my $dz = quest::get_expedition_by_zone_instance($zoneid, $instanceid);
if ($dz) {
  quest::debug(sprintf("zone id [%d] instance id [%d] is for expedition id [%d]", $zoneid, $instanceid, $dz->GetID()));
}
```

## `HASHREF` get\_expedition\_lockout\_by\_char\_id\(int character\_id, string expedition\_name, string event\_name\) <a id="get-expedition-lockout-by-char-id"></a>

Returns a hash reference containing lockout details.

The hash contains the keys `remaining` and `uuid`. `remaining` is the seconds remaining on the lockout and `uuid` is the expedition uuid the lockout was originally assigned in.

If the character doesn't have a lockout for the specified expedition event empty hash values are returned

```perl
my $lockout = quest::get_expedition_lockout_by_char_id($client->CharacterID(), "Anguish, the Fallen Palace", "Ture");
if (keys %{ $lockout }) {
  quest::debug(sprintf("lockout remaining: [%d] seconds, uuid: [%s]", $lockout->{'remaining'}, $lockout->{'uuid'}));
} else {
  quest::debug(sprintf("character [%d] does not have this lockout", $client->CharacterID()));
}
```

## `HASHREF` get\_expedition\_lockouts\_by\_char\_id\(int character\_id, string expedition\_name = ""\) <a id="get-expedition-lockouts-by-char-id"></a>

Returns a hash reference containing all of character's current lockouts keyed by expedition name. If an optional `expedition_name` is passed in it returns a hash reference of lockouts keyed by event name for that expedition.

Each character lockout contains a hash reference with `remaining` and `uuid` keys

If the character has no lockouts, the hash values are empty

```perl
# without filter on expedition name
my $lockouts = quest::get_expedition_lockouts_by_char_id($client->CharacterID());
foreach my $expedition_name (keys %{ $lockouts }) {
  my $events = $lockouts->{$expedition_name};
  foreach my $event_name (keys %{ $events }) {
    my $lockout = $events->{$event_name};
    quest::debug("expedition: [$expedition_name] event: [$event_name] seconds remaining: [$lockout->{'remaining'}] uuid: [$lockout->{'uuid'}]");
  }
}
```

```perl
# with filter on expedition name
my $anguish_lockouts = quest::get_expedition_lockouts_by_char_id($client->CharacterID(), "Anguish, the Fallen Palace");
foreach my $event_name (keys %{ $anguish_lockouts }) {
  my $lockout = $anguish_lockouts->{$event_name};
  quest::debug("event: [$event_name] seconds remaining: [$lockout->{'remaining'}] uuid: [$lockout->{'uuid'}]");
}
```

## `void` remove\_all\_expedition\_lockouts\_by\_char\_id\(int character\_id, string expedition\_name = ""\) <a id="remove-all-expedition-lockouts-by-char-id"></a>

Removes all expedition lockouts from the character, optionally filtered on an expedition name

```perl
# removes all lockouts
quest::remove_all_expedition_lockouts_by_char_id($client->CharacterID());
# removes all anguish lockouts
quest::remove_all_expedition_lockouts_by_char_id($client->CharacterID(), "Anguish, the Fallen Palace");
```

## `void` remove\_expedition\_lockout\_by\_char\_id\(int character\_id, string expedition\_name, string event\_name\) <a id="remove-expedition-lockout-by-char-id"></a>

Removes the specified lockout from the character. No-op if the character doesn't have the lockout

```perl
quest::remove_expedition_lockout_by_char_id($client->CharacterID(), "Added expedition", "Added event");
```

