# Client Methods

## `void` AddExpeditionLockout\(string expedition\_name, string event\_name, uint32 seconds, string uuid = ""\)

Assigns the specified lockout to the client. If the client already has the lockout it's replaced with this one. If a uuid isn't supplied one is generated

```lua
e.other:AddExpeditionLockout("Anguish, the Fallen Palace", "Ture", 300)
```

## `void` AddExpeditionLockoutDuration\(string expedition\_name, string event\_name, int seconds, string uuid = ""\)

Adds `seconds` to the specified client lockout duration. If `seconds` is negative then the duration of the lockout is reduced \(limited to 0\).

If the client does not have the lockout and `seconds` is positive, a new lockout is created with the specified uuid. If a uuid isn't supplied one is generated.

```lua
 -- Decrease Ture lockout by 5 minutes
e.other:AddExpeditionLockoutDuration("Anguish, the Fallen Palace", "Ture", -300)
 -- Increase keldovan lockout by 5 minutes
e.other:AddExpeditionLockoutDuration("Anguish, the Fallen Palace", "Keldovan the Harrier", 300)
```

## `Expedition*` CreateExpedition\(string zone\_short\_name, int zone\_version, int duration, string expedition\_name, int min\_players, int max\_players, bool disable\_messages = false\)

This initiates an expedition creation request by the client. All members in the client's current group or raid \(if in one\) are validated for the expedition request. On success an expedition object is returned and all members of the client's group/raid are added to the expedition. Returns nullptr if the request is rejected

This method takes both the instance zone details and expedition details

Passing `true` for the optional `disable_messages` parameter disables all expedition rejection messages sent to the leader \(member lockout conflicts, etc\)

> Note: Expedition creation follows behavior introduced with the live September 16, 2020 patch. The expedition may be created even if the group or raid exceeds the expedition's max player requirement, but only members up to the max are added to it. Players are added based on priority of raid group number followed by ungrouped players. The raid leader is added first. Players inside groups or raid groups may not be added in displayed order.

```lua
function event_say(e)
  if e.message == "request" then
    local expedition = e.other:CreateExpedition("anguish", 0, 36000, "Anguish, the Fallen Palace", 6, 54)
    if expedition.valid then
      eq.debug("Expedition creation was succesful")
    end
  end
end
```

## `Expedition*` CreateExpedition\(LUA\_TTABLE expedition\_info\)

This is an alternate expedition request method that uses a table containing all creation data.

This table supports adding all dynamic zone data at creation time rather than post-creation. This has the benefit of only hitting the database once to insert the expedition's dz data.

Option data is placed in the table as hash keys. The `expedition` and `instance` keys are required to define creation

Keys marked as `optional` may be excluded if not used

| key name | supported keys | description |
| :--- | :--- | :--- |
| `expedition` | name, min\_players, max\_players, disable\_messages | expedition options |
| `instance` | zone, version, duration | dz instance details |
| `compass` | zone, x, y, z | `(optional)` compass coordinates |
| `safereturn` | zone, x, y, z, h | `(optional)` safe return coordinates |
| `zonein` | x, y, z, h | `(optional)` dz zone-in coordinates |

#### expedition keys

| key | type | description |
| :--- | :---: | :--- |
| `name` | string | expedition name |
| `min_players` | int | minimum players to request expedition |
| `max_players` | int | maximum expedition players |
| `disable_messages` | bool | `(optional)` disable leader messages on request |

#### instance keys

| key | type | description |
| :--- | :---: | :--- |
| `zone` | int or string | zone id or zone short\_name of the expedition's dz instance |
| `version` | int | dz instance version |
| `duration` | int | dz expire time \(seconds\) |

#### compass, safereturn, and zonein keys

| key | type | description |
| :--- | :---: | :--- |
| `zone` | int or string | zone id or zone short name |
| `x` | float | `(optional)` x coordinate \(default: `0`\) |
| `y` | float | `(optional)` y coordinate \(default: `0`\) |
| `z` | float | `(optional)` z coordinate \(default: `0`\) |
| `h` | float | `(optional)` heading \(default: `0`\) |

```lua
local anguish_info = {
  expedition = { name="Anguish, the Fallen Palace", min_players=6, max_players=54 },
  instance   = { zone="anguish", version=0, duration=21600 },
  compass    = { zone=300, x=1353.15, y=1712.19, z=109.001 },
  safereturn = { zone="wallofslaughter", x=1349.13, y=1715.00, z=123.81, h=0 },
  zonein     = { x=-9, y=-2466, z=-79, h=0 }
}
local dz = e.other:CreateExpedition(anguish_info)
if dz.valid then
  eq.debug("Expedition creation was successful")
end
```

## `Expedition*` GetExpedition\(\)

Return client's current expedition if client is in one, otherwise returns nullptr

```lua
local dz = e.other:GetExpedition()
if not dz.valid then
  eq.debug("Client is not in an expedition")
end
```

## `LUA_TTABLE` GetExpeditionLockouts\(string expedition\_name = ""\)

Returns a hash table keyed on expedition name containing an event hash table with client's current lockouts. If expedition\_name is supplied then just an event hash table is returned for that expedition. The value of event hash tables is the lockout's remaining duration. Returns empty table if client doesn't have any lockouts

```lua
local lockouts = e.other:GetLockouts()
for expedition_name,event_lockouts in pairs(lockouts) do
  for event_name,remaining in pairs(event_lockouts) do
    eq.debug(string.format("expedition: (%s) event: (%s) remaining: (%s) s",
      expedition_name, event_name, remaining))
  end
end
```

```lua
-- filtered on expedition name
local lockouts = e.other:GetLockouts("Anguish, the Fallen Palace")
for event_name,remaining in pairs(lockouts) do
  eq.debug(string.format("Anguish lockouts event: (%s) remaining: (%s) s",
    event_name, remaining))
end
```

## `string` GetLockoutExpeditionUUID\(string expedition\_name, string event\_name\)

Returns uuid string of the expedition the lockout was received in. Returns an empty string if client doesn't have the lockout

```lua
local uuid = e.other:GetLockoutExpeditionUUID("Anguish, the Fallen Palace", "Ture")
if uuid ~= "" then
  eq.debug(string.format("Client received lockout in expedition uuid %s", uuid))
end
```

## `bool` HasExpeditionLockout\(string expedition\_name, string event\_name\)

Returns true of the client has the lockout, false otherwise

```lua
local has_lockout = e.other:HasExpeditionLockout("Anguish, the Fallen Palace", "Ture")
if has_lockout then
  eq.debug("Client has the lockout")
else
  eq.debug("Client does not have the lockout")
end
```

## `void` MovePCDynamicZone\(int zone\_id, int zone\_version = -1, bool msg\_if\_invalid = true\)

## `void` MovePCDynamicZone\(string zone\_short\_name, int zone\_version = -1, bool msg\_if\_invalid = true\)

Attempts to move the client to the specified zone if the client is part of a system \(expedition or otherwise\) with an associated dynamic zone instance in the target zone.

If a valid `zone_version` \(0 or higher\) is supplied then the move will only be allowed if a player's dz to `zone_id` is also for that `zone_version`.

If the client is not associated with a dz for the specified zone, the client is not moved. If `msg_if_invalid` is true \(default\) the client is sent a message that the way is blocked.

If the client is associated with multiple dz systems to the same zone, the client's `DynamicZoneSwitchListWnd` interface window is invoked to allow the client to choose which dynamic zone to enter.

On success, the player is moved to the dynamic zone's zone in coordinates \(set on creation or afterwards\) If no zone in coordinates are set, the player is moved to the zone's default safe coords.

`MovePCInstance` should be used if finer control of zone in location is required, but care must be taken to validate the client is part of an instance for the target zone

```lua
e.other:MovePCDynamicZone(317) -- by zone id
e.other:MovePCDynamicZone("anguish") -- by zone short name
```

## `void` RemoveAllExpeditionLockouts\(string expedition\_name = ""\)

Removes all expedition lockouts from the character, optionally filtered on an expedition name

```lua
-- removes all lockouts
e.other:RemoveAllExpeditionLockouts(e.other:CharacterID())
-- removes all anguish lockouts
e.other:RemoveAllExpeditionLockouts(e.other:CharacterID(), "Anguish, the Fallen Palace")
```

## `void` RemoveExpeditionLockout\(string expedition\_name, string event\_name\)

Removes the specified lockout from the character. No-op if the character doesn't have the lockout

```lua
e.other:RemoveExpeditionLockout("Anguish, the Fallen Palace", "Ture")
```

