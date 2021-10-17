# Expedition Methods

## `void` AddLockout\(string event\_name, uint32 seconds\_duration\) <a id="dz-add-lockout"></a>

Adds an event lockout timer to the expedition with the specified duration.

The lockout is added to all current members of the expedition whether they're inside the zone or not. This also adds the lockout to any clients inside the dynamic zone that are not part of the expedition \(in the event that they've quit the expedition and haven't been teleported out yet\)

The lockout is also stored internally in the expedition. New members added to the expedition will receive these lockouts once they zone into the expedition's dynamic zone.

The lockout replaces any existing lockout held by the expedition and members.

The lockout is assigned the current expedition's internal uuid for tracking.

> Note: The source adds 60s to the lockout timer sent to clients to compensate for the client displaying lockouts rounded down to the nearest minute. If adding lockout times using live packet data subtract 60s from the time \(live adds 60s as well\)

```lua
function event_death(e)
  local expedition = eq.get_expedition() -- current instance's expedition (we should be inside a dz)
  if not expedition.valid then
    eq.debug("report a bug") -- possible fallback would be to manually iterate clients and add lockouts
  else
    expedition:AddLockout("Keldovan the Harrier", 86400) -- add a 1 day lockout
  end
end
```

## `void` AddLockoutDuration\(string event\_name, int seconds, bool members\_only = true\) <a id="dz-add-lockout-duration"></a>

Adds duration to existing member lockouts for the event. If `seconds` is negative then lockout duration is reduced.

Unlike UpdateLockoutDuration which applies a static lockout duration to all members based on the expedition's internal lockout, this modifies each member's existing lockout duration separately.

If a member does not have the specified event lockout and `seconds` is positive, the lockout is added to the member with the current expedition's uuid.

By default this only updates expedition members. Passing `false` for `members_only` will update the expedition's internal lockout duration as well which will affect any future members added to the expedition.

> Non-expedition members inside the expedition's dynamic zone instance are always updated

```lua
-- inside dz, after killing some progression npcs
function event_death(e)
  local dz = eq.get_expedition()
  if dz.valid then
    -- add durations to each member's current event lockout when specific npcs killed
    -- the expedition's lockout is unaffected, new members will receive original lockout
    if e.self:GetNPCTypeID() == 294042 then
      dz:AddLockout("Some Event", 25200) -- create 7 hour lockout when this npc is killed
    elseif e.self:GetNPCTypeID() == 294043 then
      dz:AddLockoutDuration("Some Event", 57600) -- add 16 hours when this npc is killed
    end
  end
end
```

## `void` AddReplayLockout\(int seconds\_duration\) <a id="dz-add-replay-lockout"></a>

Adds a `Replay Timer` lockout to the expedition. This is a shortcut alternative to adding the lockout by name through the `AddLockout` method.

Replay Timer's use a hardcoded event name and are used as a lockout against requesting the expedition itself.

Unlike normal event lockouts, any members added to the expedition after a replay timer has been assigned will receive the lockout immediately with a reset expiration timer based on the original duration.

```lua
function event_say(e)
  if e.message == "request" then
    local expedition = e.other:CreateExpedition("anguish", 0, 36000, "Anguish, the Fallen Palace", 6, 54)
    if expedition.valid then
      eq.debug("Expedition creation was succesful")
      expedition:AddReplayLockout(7200) -- 2 hour request lockout
    end
  end
end
```

## `void` AddReplayLockoutDuration\(int seconds, bool members\_only = true\) <a id="dz-add-replay-lockout-duration"></a>

Adds duration to expedition's `Replay Timer` lockout. See AddLockoutDuration

```lua
function event_say(e)
  if e.message == "request" then
    local dz = e.other:CreateExpedition("ikkinz", 0, 21600, "Ikkinz, Chambers of Singular Might", 1, 6)
    if dz.valid then
      expedition:AddReplayLockout(3600) -- 1 hour replay lockout
    end
  end
end

-- inside dz, after killing some progression npcs
function event_death(e)
  local dz = eq.get_expedition()
  if dz.valid then
    if e.self:GetNPCTypeID() == 294042 then
      dz:AddReplayLockoutDuration(25200) -- add 7 hours to replay lockout when this npc is killed
    elseif e.self:GetNPCTypeID() == 294043 then
      -- add 16 hours to replay lockout when this npc is killed
      -- expedition lockout is unaffected, new members will receive the original (7 hours - elapsed)
      dz:AddReplayLockoutDuration(57600)
    end
  end
end
```

## `uint32_t` GetDynamicZoneID\(\) <a id="dz-get-dz-id"></a>

Returns the expedition's dynamic zone id

```lua
local dz_id = expedition:GetDynamicZoneID()
```

## `uint32_t` GetID\(\) <a id="dz-get-id"></a>

Returns the expedition's id

```lua
local id = expedition:GetID()
```

## `int` GetInstanceID\(\) <a id="dz-get-instance-id"></a>

Returns the expedition's dynamic zone instance id

```lua
local instance_id = expedition:GetInstanceID()
```

## `string` GetLeaderName\(\) <a id="dz-get-leader-name"></a>

Returns the expedition's current leader name

```lua
local leader_name = expedition:GetLeaderName()
```

## `LUA_TTABLE` GetLockouts\(\) <a id="dz-get-lockouts"></a>

Returns a table containing the expedition's current internal lockouts and time remaining. This includes lockouts that were inherited by the leader on creation and not just lockouts assigned during the current expedition.

```lua
local lockouts = expedition:GetLockouts()
for event_name,remaining in pairs(lockouts) do
  eq.debug(string.format("event: (%s) remaining: (%s)s", event_name, remaining))
end
```

## `string` GetLootEventByNPCTypeID\(uint32 npc\_type\_id\) <a id="dz-get-loot-event-by-npctype"></a>

## `string` GetLootEventBySpawnID\(uint32 spawn\_id\) <a id="dz-get-loot-event-by-spawnid"></a>

Returns an event name associated with an npc type id or spawn id set by the SetLootEventByNPCTypeID or SetLootEventBySpawnID apis.

Returns an empty string if no event set

> Important: This is only valid inside the expedition's dynamic zone instance

```lua
function event_death(e)
  local npc_type_id = e.self:GetNPCTypeID()
  local spawn_id = e.self:GetID()
  local dz = eq.get_expedition() -- current instance's expedition (we should be inside a dz)
  if dz.valid then
    local event_name = dz:GetLootEventByNPCTypeID(npc_type_id)
    eq.debug(string.format("Npc type (%s) was associated with event (%s)", npc_type_id, event_name))
    event_name = dz:GetLootEventBySpawnID(spawn_id)
    eq.debug(string.format("This npc entity id (%s) was associated with event (%s)", spawn_id, event_name))
  end
end
```

## `int` GetMemberCount\(\) <a id="dz-get-member-count"></a>

Returns number of members in expedition

```lua
local count = dz:GetMemberCount()
```

## `LUA_TTABLE` GetMembers\(\) <a id="dz-get-members"></a>

Returns a Lua table containing member names and character ids

```lua
local members = dz:GetMembers()
for name,id in pairs(members) do
  eq.debug(string.format("Member name (%s) character id (%s)", name, id))
end
```

## `string` GetName\(\) <a id="dz-get-name"></a>

Returns name of the expedition

```lua
local expedition_name = dz:GetName()
```

## `int` GetSecondsRemaining\(\) <a id="dz-get-seconds-remaining"></a>

Returns seconds remaining until expedition's dynamic zone expires

```lua
local seconds_remaining = dz:GetSecondsRemaining()
```

## `string` GetUUID\(\) <a id="dz-get-uuid"></a>

Returns expedition's internal uuid assigned on creation

```lua
local uuid = dz:GetUUID()
```

## `int` GetZoneID\(\) <a id="dz-get-zone-id"></a>

Returns zone id of the expedition's dynamic zone

```lua
local zone_id = dz:GetZoneID()
```

## `string` GetZoneName\(\) <a id="dz-get-zone-name"></a>

Returns zone short name of the expedition's dynamic zone

```lua
local zone_short_name = dz:GetZoneName()
```

## `int` GetZoneVersion\(\) <a id="dz-get-zone-version"></a>

Returns zone version of the expedition's dynamic zone instance

```lua
local zone_version = dz:GetZoneVersion()
```

## `bool` HasLockout\(string event\_name\) <a id="dz-has-lockout"></a>

Returns whether the expedition internally has the specified event lockout. This is the primary function that should be used for initialization scripts to determine which events are available to an expedition.

```lua
function init()
  local dz = eq.get_expedition()
  if dz.valid then
    local has_lockout = dz:HasLockout("Ture")
    if not has_lockout then
      eq.debug("Expedition doesn't have this lockout, we can spawn this event")
    end
  end
end
```

## `bool` HasReplayLockout\(\) <a id="dz-has-replay-lockout"></a>

Returns whether the expedition internally has a `Replay Timer` lockout set. Shortcut alternative to checking `HasLockout`

```lua
local has_replay_lockout = dz:HasReplayLockout()
```

## `bool` IsLocked\(\) <a id="dz-is-locked"></a>

Returns if the expedition is currently locked by `SetLocked`

```lua
local dz = eq.get_expedition()
if dz.valid then
  if dz:IsLocked() then
    eq.debug("expedition is locked")
  else
    eq.debug("expedition is not locked")
  end
end
```

## `void` RemoveCompass\(\) <a id="dz-remove-compass"></a>

Removes any compass set by `SetCompass`

```lua
dz:RemoveCompass()
```

## `void` RemoveLockout\(string event\_name\) <a id="dz-remove-lockout"></a>

Removes the lockout from the expedition and from any current members

```lua
dz:RemoveLockout("Ture")
```

## `void` SetCompass\(int zone\_id, float x, float y, float z\) <a id="dz-set-compass-zoneid"></a>

## `void` SetCompass\(string zone\_short\_name, float x, float y, float z\) <a id="dz-set-compass-zonename"></a>

Set the location of the expedition's dynamic zone compass. This only applies to non-instances outside of the dz.

```lua
function event_say(e)
  if e.message == "request" then
    local dz = e.other:CreateExpedition("anguish", 0, 36000, "Anguish, the Fallen Palace", 6, 54)
    if dz.valid then
      -- set compass to the anguish door inside wallofslaughter by zone id
      dz:SetCompass(300, 1353.15, 1712.19, 109.001)
      -- set compass to the anguish door inside wallofslaughter by zone name
      dz:SetCompass("wallofslaughter", 1353.15, 1712.19, 109.001)
    end
  end
end
```

## `void` SetLocked\(bool locked, ExpeditionLockMessage lock\_msg = ExpeditionLockMessage::None, uint32\_t msg\_color = Chat::Yellow\) <a id="dz-set-locked"></a>

If true, lock an expedition to prevent new members from being added. False unlocks.

Passing a `lock_msg` sends an expedition locked message to clients inside the expedition's dynamic zone. See ExpeditionLockMessage for types and messages

```lua
function event_death(e)
  local dz = eq.get_expedition()
  if dz.valid then
    -- no new members may be added to this expedition until it's unlocked
    dz:SetLocked(true, ExpeditionLockMessage.Close)
  end
end
```

## `void` SetLootEventByNPCTypeID\(uint32\_t npc\_type\_id, string event\_name\) <a id="dz-set-loot-event-by-npctype"></a>

> Important: This must be set from a script inside the expedition's dynamic zone instance

Associates an event with the specified npc type id to prevent it from being looted by characters that didn't receive the event lockout from the current expedition.

This is an important initialization step to prevent exploits in expeditions

```lua
function init()
  local dz = eq.get_expedition()
  if dz.valid then
    -- associate the "Keldovan the Harrier" event with his npc type id
    dz:SetLootEventByNPCTypeID(317005, 'Keldovan the Harrier')
  end
end
```

## `void` SetLootEventBySpawnID\(uint32\_t spawn\_id, string event\_name\) <a id="dz-set-loot-event-by-spawnid"></a>

> Important: This must be set from a script inside the expedition's dynamic zone instance

Associates an event with the specified entity spawn id to prevent it from being looted by characters that didn't receive the event lockout from the current expedition.

This is an important initialization step to prevent exploits in expeditions

Unlike `SetLootEventByNPCTypeID` this associates the event with specific spawns inside the dynamic zone. This is for cases where multiple npc types may exist like spawnable chests which should be associated with the event that spawned them.

```lua
function event_death(e)
  local dz = eq.get_expedition()
  if dz.valid then
    dz:AddLockout("Lower Globe of Discordant Energy")

    -- associate this spawned chest with the event
    local chest = eq.spawn2(317087, 0, 0, -301 ,702, -201, 0);
    dz:SetLootEventBySpawnID(chest:GetID(), "Lower Globe of Discordant Energy")
  end
end
```

## `void` SetReplayLockoutOnMemberJoin\(bool value\) <a id="dz-set-replay-lockout-on-member-join"></a>

This toggles whether `Replay Timer` lockouts are automatically added to new members that join an expedition

```lua
function event_say(e)
  if e.message == "request" then
    local dz = e.other:CreateExpedition("anguish", 0, 36000, "Anguish, the Fallen Palace", 6, 54)
    if dz.valid then
      dz:SetReplayLockoutOnMemberJoin(false)
    end
  end
end
```

## `void` SetSafeReturn\(uint32\_t zone\_id, float x, float y, float z, float heading\) <a id="dz-set-safe-return-zoneid"></a>

## `void` SetSafeReturn\(string zone\_short\_name, float x, float y, float z, float heading\) <a id="dz-set-safe-return-zonename"></a>

Set the expedition's dynamic zone safe return zone and coordinates. This is the location clients are moved to when an expedition expires or for clients that are automatically teleported out of the zone by the kick timer if they quit the expedition

```lua
function event_say(e)
  if e.message == "request" then
    local dz = e.other:CreateExpedition("anguish", 0, 36000, "Anguish, the Fallen Palace", 6, 54)
    if dz.valid then
      -- by zone id
      dz:SetSafeReturn(300, 1349.13, 1715.00, 123.81, 0)
      -- or by zone name
      dz:SetSafeReturn("wallofslaughter", 1349.13, 1715.00, 123.81, 0)
    end
  end
end
```

## `void` SetSecondsRemaining\(uint32 seconds\_remaining\) <a id="dz-set-seconds-remaining"></a>

Set seconds remaining on expedition's dynamic zone before it expires. This currently only supports reducing time. It will have no effect if the new seconds remaining is larger than the current time remaining.

> Note: this method is async, GetSecondsRemaining will not immediately show the change.

```lua
function event_timer(e)
  if e.timer == "event_fail_timer" then
    eq.depop()
    local dz = eq.get_expedition()
    if dz.valid then
      dz:SetSecondsRemaining(300) -- close down expedition in 5 minutes
    end
  end
end
```

## `void` SetZoneInLocation\(float x, float y, float z, float heading\) <a id="dz-set-zonein-location"></a>

Override the zone-in location of the expedition's dynamic zone used by `Client::MovePCDynamicZone`

```lua
function event_say(e)
  if e.message == "request" then
    local dz = e.other:CreateExpedition("anguish", 0, 36000, "Anguish, the Fallen Palace", 6, 54)
    if dz.valid then
      dz:SetZoneInLocation(-9, -2466, -79, 0)
    end
  end
end
```

## `void` UpdateLockoutDuration\(string event\_name, uint32\_t seconds, bool members\_only = true\) <a id="dz-update-lockout-duration"></a>

Sets the lockout of all members in the expedition to the specified duration based on the expedition's internal lockout start time.

Unlike `AddLockout` which replaces lockouts with a new lockout timer, this modifies the duration of an existing lockout relative to its original start time.

Calling this on an expedition that does not have the internal lockout has no effect. All members essentially receive an updated lockout copied from the expedition's internal timer.

> Prefer using AddLockoutDuration to update individual member lockouts

The expedition's internal lockout is NOT modified by default. New members added to the expedition that zone in to the dz will receive the lockout as originally added. This can be overridden by passing false for the `members_only` parameter.

> Note: Member lockouts updated with this method are assigned the current expedition's uuid

```lua
function event_say(e)
  if e.message == "request" then
    local dz = e.other:CreateExpedition("anguish", 0, 36000, "Anguish, the Fallen Palace", 6, 54)
    if dz.valid then
      expedition:AddReplayLockout(7200) -- 2 hour replay lockout
    end
  end
end

-- inside dz, after killing some progression npc
function event_death(e)
  local dz = eq.get_expedition()
  if dz.valid then
    -- copies expedition's lockout to all current members with modified duration.
    -- member lockout durations will be (7 hours - elapsed time since added)
    dz:UpdateLockoutDuration("Replay Timer", 25200) -- 7 hours
  end
end
```

