# Client Methods

## `void` AddExpeditionLockout\(string expedition\_name, string event\_name, uint32 seconds, string uuid = ""\) <a id="client-add-lockout"></a>

Assigns the specified lockout to the client. If the client already has the lockout it's replaced with this one. If a uuid isn't supplied one is generated

```perl
# Add 10 minute lockouts
$client->AddExpeditionLockout("Anguish, the Fallen Palace", "Ture", 600);
$client->AddExpeditionLockout("Anguish, the Fallen Palace", "Keldovan the Harrier", 600);
```

## `void` AddExpeditionLockoutDuration\(string expedition\_name, string event\_name, int seconds, string uuid = ""\) <a id="client-add-lockout-duration"></a>

Adds `seconds` to the specified client lockout duration. If `seconds` is negative then the duration of the lockout is reduced \(limited to 0\).

If the client does not have the lockout and `seconds` is positive, a new lockout is created with the specified uuid. If a uuid isn't supplied one is generated.

```perl
# Decrease Ture lockout by 5 minutes
$client->AddExpeditionLockoutDuration("Anguish, the Fallen Palace", "Ture", -300);
# Increase keldovan lockout by 5 minutes
$client->AddExpeditionLockoutDuration("Anguish, the Fallen Palace", "Keldovan the Harrier", 300);
```

## `Expedition*` CreateExpedition\(string zone\_short\_name, int zone\_version, int duration, string expedition\_name, int min\_players, int max\_players, bool disable\_messages = false\) <a id="client-create-expedition"></a>

This initiates an expedition creation request by the client. All members in the client's current group or raid \(if in one\) are validated for the expedition request. On success an expedition object is returned and all members of the client's group/raid are added to the expedition. Returns nullptr if the request is rejected

This method takes both the instance zone details and expedition details

Passing `true` for the optional `disable_messages` parameter disables all expedition rejection messages sent to the leader \(member lockout conflicts, etc\)

> Note: Expedition creation follows behavior introduced with the live September 16, 2020 patch. The expedition may be created even if the group or raid exceeds the expedition's max player requirement, but only members up to the max are added to it. Players are added based on priority of raid group number followed by ungrouped players. The raid leader is added first. Players inside groups or raid groups may not be added in displayed order.

```perl
my $dz = $client->CreateExpedition("anguish", 0, 3600, "Anguish, the Fallen Palace", 6, 54);
if ($dz) {
  quest::debug(sprintf("expedition [%d] created", $dz->GetID()));
} else {
  quest::debug(sprintf("expedition request failed"));
}
```

## `Expedition*` GetExpedition\(\) <a id="client-get-expedition"></a>

Return client's current expedition if client is in one, otherwise returns nullptr

```perl
my $dz = $client->GetExpedition();
if ($dz) {
  quest::debug(sprintf("client in expedition [%d]", $dz->GetID()));
} else {
  quest::debug(sprintf("client is not in an expedition"));
}
```

## `HASHREF` GetExpeditionLockouts\(string expedition\_name = ""\) <a id="client-get-lockouts"></a>

Returns a hash reference keyed on expedition name containing an event hash reference with client's current lockouts. If expedition\_name is supplied then just an event hash reference is returned for that expedition. The value of event hash references is the lockout's remaining duration. Returns empty value if client doesn't have any lockouts

```perl
my $lockouts = $client->GetExpeditionLockouts();
foreach my $expedition_name (keys %{ $lockouts }) {
  my $events = $lockouts->{$expedition_name};
  foreach my $event_name (keys %{ $events }) {
    my $remaining = $events->{$event_name};
    quest::debug("expedition: [$expedition_name] event: [$event_name] seconds remaining: [$remaining]");
  }
}
```

```perl
# filtered on expedition name
my $lockouts = $client->GetExpeditionLockouts("Anguish, the Fallen Palace");
foreach my $event_name (keys %{ $lockouts }) {
  my $remaining = $lockouts->{$event_name};
  quest::debug("anguish event: [$event_name] seconds remaining: [$remaining]");
}
```

## `string` GetLockoutExpeditionUUID\(string expedition\_name, string event\_name\) <a id="client-get-lockout-uuid"></a>

Returns uuid string of the expedition the lockout was received in. Returns an empty string if client doesn't have the lockout

```perl
my $uuid = $client->GetLockoutExpeditionUUID("Anguish, the Fallen Palace", "Ture");
quest::debug("client received lockout in expedition uuid: [$uuid]");
```

## `bool` HasExpeditionLockout\(string expedition\_name, string event\_name\) <a id="client-has-lockout"></a>

Returns true of the client has the lockout, false otherwise

```perl
my $has_lockout = $client->HasExpeditionLockout("Anguish, the Fallen Palace", "Ture");
quest::debug("client has lockout: [$has_lockout]");
```

## `void` MovePCDynamicZone\(int zone\_id, int zone\_version = -1, bool msg\_if\_invalid = true\) <a id="client-move-pc-zoneid"></a>

## `void` MovePCDynamicZone\(string zone\_short\_name, int zone\_version = -1, bool msg\_if\_invalid = true\) <a id="client-move-pc-zonename"></a>

Attempts to move the client to the specified zone if the client is part of a system \(expedition or otherwise\) with an associated dynamic zone instance in the target zone.

If a valid `zone_version` \(0 or higher\) is supplied then the move will only be allowed if a player's dz to `zone_id` is also for that `zone_version`.

If the client is not associated with a dz for the specified zone, the client is not moved. If `msg_if_invalid` is true \(default\) the client is sent a message that the way is blocked.

If the client is associated with multiple dz systems to the same zone, the client's `DynamicZoneSwitchListWnd` interface window is invoked to allow the client to choose which dynamic zone to enter.

On success, the player is moved to the dynamic zone's zone in coordinates \(set on creation or afterwards\) If no zone in coordinates are set, the player is moved to the zone's default safe coords.

`MovePCInstance` should be used if finer control of zone in location is required, but care must be taken to validate the client is part of an instance for the target zone

```perl
$client->MovePCDynamicZone(317); # by zone id
$client->MovePCDynamicZone("anguish"); # by zone short name
```

## `void` RemoveAllExpeditionLockouts\(string expedition\_name = ""\) <a id="client-remove-all-lockouts"></a>

Removes all expedition lockouts from the character, optionally filtered on an expedition name

```perl
# removes all lockouts
$client->RemoveAllExpeditionLockouts();
# removes all anguish lockouts
$client->RemoveAllExpeditionLockouts("Anguish, the Fallen Palace");
```

## `void` RemoveExpeditionLockout\(string expedition\_name, string event\_name\) <a id="client-remove-lockout"></a>

Removes the specified lockout from the character. No-op if the character doesn't have the lockout

```perl
$client->RemoveExpeditionLockout("Anguish, the Fallen Palace", "Ture");
```

