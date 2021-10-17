# Expedition Methods

## `void` AddLockout\(string event\_name, uint32 seconds\_duration\) <a id="dz-add-lockout"></a>

Adds an event lockout timer to the expedition with the specified duration.

The lockout is added to all current members of the expedition whether they're inside the zone or not. This also adds the lockout to any clients inside the dynamic zone that are not part of the expedition \(in the event that they've quit the expedition and haven't been teleported out yet\)

The lockout is also stored internally in the expedition. New members added to the expedition will receive these lockouts once they zone into the expedition's dynamic zone.

The lockout replaces any existing lockout held by the expedition and members.

The lockout is assigned the current expedition's internal uuid for tracking.

> Note: The source adds 60s to the lockout timer sent to clients to compensate for the client displaying lockouts rounded down to the nearest minute. If adding lockout times using live packet data subtract 60s from the time \(live adds 60s as well\)

```perl
my $dz = $client->GetExpedition();
if ($dz) {
  $dz->AddLockout("Expedition Event Name", 3600); # 1 hour lockout
}
```

## `void` AddLockoutDuration\(string event\_name, int seconds, bool members\_only = true\) <a id="dz-add-lockout-duration"></a>

Adds duration to existing member lockouts for the event. If `seconds` is negative then lockout duration is reduced.

Unlike UpdateLockoutDuration which applies a static lockout duration to all members based on the expedition's internal lockout, this modifies each member's existing lockout duration separately.

If a member does not have the specified event lockout and `seconds` is positive, the lockout is added to the member with the current expedition's uuid.

By default this only updates expedition members. Passing `false` for `members_only` will update the expedition's internal lockout duration as well which will affect any future members added to the expedition.

> Non-expedition members inside the expedition's dynamic zone instance are always updated

```perl
my $dz = $client->GetExpedition();
if ($dz) {
  $dz->AddLockoutDuration("Expedition Event Name", 3600); # add 1 hour to lockout
}
```

## `void` AddReplayLockout\(int seconds\_duration\) <a id="dz-add-replay-lockout"></a>

Adds a `Replay Timer` lockout to the expedition. This is a shortcut alternative to adding the lockout by name through the `AddLockout` method.

Replay Timer's use a hardcoded event name and are used as a lockout against requesting the expedition itself.

Unlike normal event lockouts, any members added to the expedition after a replay timer has been assigned will receive the lockout immediately with a reset expiration timer based on the original duration.

```perl
my $dz = $client->GetExpedition();
if ($dz) {
  $dz->AddReplayLockout(7200); # 2 hour replay timer
}
```

## `void` AddReplayLockoutDuration\(int seconds, bool members\_only = true\) <a id="dz-add-replay-lockout-duration"></a>

Adds duration to expedition's `Replay Timer` lockout. See AddLockoutDuration

```perl
my $dz = $client->GetExpedition();
if ($dz) {
  $dz->AddReplayLockoutDuration(7200); # add 2 hours to replay timer
}
```

## `uint32_t` GetDynamicZoneID\(\) <a id="dz-get-dz-id"></a>

Returns the expedition's dynamic zone id

```perl
my $dz_id = $expedition->GetDynamicZoneID();
```

## `uint32_t` GetID\(\) <a id="dz-get-id"></a>

Returns the expedition's id

```perl
my $id = $expedition->GetID();
```

## `int` GetInstanceID\(\) <a id="dz-get-instance-id"></a>

Returns the expedition's dynamic zone instance id

```perl
my $id = $expedition->GetInstanceID();
```

## `string` GetLeaderName\(\) <a id="dz-get-leader-name"></a>

Returns the expedition's current leader name

```perl
my $leader_name = expedition:GetLeaderName();
```

## `HASHREF` GetLockouts\(\) <a id="dz-get-lockouts"></a>

Returns a hash reference containing the expedition's current internal lockouts and time remaining. This includes lockouts that were inherited by the leader on creation and not just lockouts assigned during the current expedition.

```perl
my $dz = $client->GetExpedition();
if ($dz) {
  my $lockouts = $dz->GetLockouts();
  foreach my $event_name (keys %{ $lockouts }) {
    quest::debug("event: [$event_name] seconds remaining: [$lockouts->{$event_name}]");
  }
}
```

## `string` GetLootEventByNPCTypeID\(uint32 npc\_type\_id\) <a id="dz-get-loot-event-by-npctype"></a>

## `string` GetLootEventBySpawnID\(uint32 spawn\_id\) <a id="dz-get-loot-event-by-spawnid"></a>

Returns an event name associated with an npc type id or spawn id set by the SetLootEventByNPCTypeID or SetLootEventBySpawnID apis.

Returns an empty string if no event set

> Important: This is only valid inside the expedition's dynamic zone instance

```perl
my $dz = quest::get_expedition();
if ($dz) {
  quest::debug("npc type id 12345 event: " . $dz->GetLootEventByNPCTypeID(12345));
  quest::debug("spawn id 20 event: " . $dz->GetLootEventBySpawnID(20));
}
```

## `int` GetMemberCount\(\) <a id="dz-get-member-count"></a>

Returns number of members in expedition

```perl
my $count = $dz->GetMemberCount();
```

## `HASHREF` GetMembers\(\) <a id="dz-get-members"></a>

Returns a hash reference containing member names and character ids

```perl
my $dz = $client->GetExpedition();
if ($dz) {
  my $members = $dz->GetMembers();
  foreach my $member_name (keys %{ $members }) {
    quest::debug("char name: [$member_name] char id: [$members->{$member_name}]");
  }
}
```

## `string` GetName\(\) <a id="dz-get-name"></a>

Returns name of the expedition

```perl
my $dz = $client->GetExpedition();
if ($dz) {
  quest::debug("expedition name: " . $dz->GetName());
}
```

## `int` GetSecondsRemaining\(\) <a id="dz-get-seconds-remaining"></a>

Returns seconds remaining until expedition's dynamic zone expires

```perl
my $dz = $client->GetExpedition();
if ($dz) {
  quest::debug("expedition seconds remaining: " . $dz->GetSecondsRemaining());
}
```

## `string` GetUUID\(\) <a id="dz-get-uuid"></a>

Returns expedition's internal uuid assigned on creation

```perl
my $dz = $client->GetExpedition();
if ($dz) {
  quest::debug("expedition uuid: " . $dz->GetUUID());
}
```

## `int` GetZoneID\(\) <a id="dz-get-zone-id"></a>

Returns zone id of the expedition's dynamic zone

```perl
my $dz = $client->GetExpedition();
if ($dz) {
  quest::debug("expedition's dz zone id: " . $dz->GetZoneID());
}
```

## `string` GetZoneName\(\) <a id="dz-get-zone-name"></a>

Returns zone short name of the expedition's dynamic zone

```perl
my $dz = $client->GetExpedition();
if ($dz) {
  quest::debug("expedition's dz zone name: " . $dz->GetZoneName());
}
```

## `int` GetZoneVersion\(\) <a id="dz-get-zone-version"></a>

Returns zone version of the expedition's dynamic zone instance

```perl
my $dz = $client->GetExpedition();
if ($dz) {
  quest::debug("expedition's dz zone version: " . $dz->GetZoneVersion());
}
```

## `bool` HasLockout\(string event\_name\) <a id="dz-has-lockout"></a>

Returns whether the expedition internally has the specified event lockout. This is the primary function that should be used for initialization scripts to determine which events are available to an expedition.

```perl
my $dz = quest::get_expedition();
if ($dz) {
  my $has_lockout = $dz->HasLockout("Ture") ? 1 : 0;
  quest::debug("expedition has Ture lockout: " . $has_lockout);
}
```

## `bool` HasReplayLockout\(\) <a id="dz-has-replay-lockout"></a>

Returns whether the expedition internally has a `Replay Timer` lockout set. Shortcut alternative to checking `HasLockout`

```perl
my $dz = quest::get_expedition();
if ($dz) {
  my $has_lockout = $dz->HasReplayLockout() ? 1 : 0;
  quest::debug("expedition has Replay lockout: " . $has_lockout);
}
```

## `bool` IsLocked\(\) <a id="dz-is-locked"></a>

Returns if the expedition is currently locked by `SetLocked`

```perl
my $dz = quest::get_expedition();
if ($dz) {
  if ($dz->IsLocked()) {
    quest::debug("expedition is locked");
  } else {
    quest::debug("expedition is not locked");
  }
}
```

## `void` RemoveCompass\(\) <a id="dz-remove-compass"></a>

Removes any compass set by `SetCompass`

```perl
$dz->RemoveCompass();
```

## `void` RemoveLockout\(string event\_name\) <a id="dz-remove-lockout"></a>

Removes the lockout from the expedition and from any current members

```perl
$dz->RemoveLockout("Ture");
```

## `void` SetCompass\(int zone\_id, float x, float y, float z\) <a id="dz-set-compass-zoneid"></a>

## `void` SetCompass\(string zone\_short\_name, float x, float y, float z\) <a id="dz-set-compass-zonename"></a>

Set the location of the expedition's dynamic zone compass. This only applies to non-instances outside of the dz.

```perl
$dz->SetCompass(300, 1353.15, 1712.19, 109.001); # by zone id
$dz->SetCompass("wallofslaughter", 1353.15, 1712.19, 109.001); # by zone name
```

## `void` SetLocked\(bool locked, ExpeditionLockMessage lock\_msg = ExpeditionLockMessage::None, uint32\_t msg\_color = Chat::Yellow\) <a id="dz-set-locked"></a>

If true, lock an expedition to prevent new members from being added. False unlocks.

Passing a `lock_msg` sends an expedition locked message to clients inside the expedition's dynamic zone. See ExpeditionLockMessage for types and messages

```perl
$dz->SetLocked(true, ExpeditionLockMessage::None); # lock with no message
$dz->SetLocked(true, ExpeditionLockMessage::Close); # lock with "closing" message
```

## `void` SetLootEventByNPCTypeID\(uint32\_t npc\_type\_id, string event\_name\) <a id="dz-set-loot-event-by-npctype"></a>

> Important: This must be set from a script inside the expedition's dynamic zone instance

Associates an event with the specified npc type id to prevent it from being looted by characters that didn't receive the event lockout from the current expedition.

This is an important initialization step to prevent exploits in expeditions

```perl
$dz->SetLootEventByNPCTypeID(317005, 'Keldovan the Harrier');
```

## `void` SetLootEventBySpawnID\(uint32\_t spawn\_id, string event\_name\) <a id="dz-set-loot-event-by-spawnid"></a>

> Important: This must be set from a script inside the expedition's dynamic zone instance

Associates an event with the specified entity spawn id to prevent it from being looted by characters that didn't receive the event lockout from the current expedition.

This is an important initialization step to prevent exploits in expeditions

Unlike `SetLootEventByNPCTypeID` this associates the event with specific spawns inside the dynamic zone. This is for cases where multiple npc types may exist like spawnable chests which should be associated with the event that spawned them.

```perl
my $chest_id = quest::spawn2(317087, 0, 0, -301 ,702, -201, 0);
$dz->SetLootEventBySpawnID($chest_id, "Lower Globe of Discordant Energy");
```

## `void` SetReplayLockoutOnMemberJoin\(bool value\) <a id="dz-set-replay-lockout-on-member-join"></a>

This toggles whether `Replay Timer` lockouts are automatically added to new members that join an expedition

```perl
$dz->SetReplayLockoutOnMemberJoin(0);
```

## `void` SetSafeReturn\(uint32\_t zone\_id, float x, float y, float z, float heading\) <a id="dz-set-safe-return-zoneid"></a>

## `void` SetSafeReturn\(string zone\_short\_name, float x, float y, float z, float heading\) <a id="dz-set-safe-return-zonename"></a>

Set the expedition's dynamic zone safe return zone and coordinates. This is the location clients are moved to when an expedition expires or for clients that are automatically teleported out of the zone by the kick timer if they quit the expedition

```perl
$dz->SetSafeReturn("wallofslaughter", 1349.13, 1715.00, 123.81, 0);
```

## `void` SetSecondsRemaining\(uint32 seconds\_remaining\) <a id="dz-set-seconds-remaining"></a>

Set seconds remaining on expedition's dynamic zone before it expires. This currently only supports reducing time. It will have no effect if the new seconds remaining is larger than the current time remaining.

> Note: this method is async, GetSecondsRemaining will not immediately show the change.

```perl
$dz->SetSecondsRemaining(300); # close down expedition in 5 minutes
```

## `void` SetZoneInLocation\(float x, float y, float z, float heading\) <a id="dz-set-zonein-location"></a>

Override the zone-in location of the expedition's dynamic zone used by `Client::MovePCDynamicZone`

```perl
$dz->SetZoneInLocation(-9, -2466, -79, 0);
```

## `void` UpdateLockoutDuration\(string event\_name, uint32\_t seconds, bool members\_only = true\) <a id="dz-update-lockout-duration"></a>

Sets the lockout of all members in the expedition to the specified duration based on the expedition's internal lockout start time.

Unlike `AddLockout` which replaces lockouts with a new lockout timer, this modifies the duration of an existing lockout relative to its original start time.

Calling this on an expedition that does not have the internal lockout has no effect. All members essentially receive an updated lockout copied from the expedition's internal timer.

> Prefer using AddLockoutDuration to update individual member lockouts

The expedition's internal lockout is NOT modified by default. New members added to the expedition that zone in to the dz will receive the lockout as originally added. This can be overridden by passing false for the `members_only` parameter.

> Note: Member lockouts updated with this method are assigned the current expedition's uuid

```perl
$dz->UpdateLockoutDuration("Replay Timer", 25200) # 7 hours
```

