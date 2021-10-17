---
description: Methods accessed via the Lua `eq` export table
---

# Global Methods

## `void` add\_expedition\_lockout\_all\_clients\(string expedition\_name, string event\_name, int seconds, string uuid = ""\)

Adds the specified lockout to all clients in the current zone. If a client already has the lockout it's replaced with this one. If a uuid is not supplied one is generated and applied to all clients.

```lua
eq.add_expedition_lockout_all_clients("Anguish, the Fallen Palace", "Ture", 300)
```

## `void` add\_expedition\_lockout\_by\_char\_id\(int character\_id, string expedition\_name, string event\_name, int seconds, string uuid = ""\)

Adds the specified lockout to the character with an optional specified uuid. If the client already has the lockout it's replaced with this one. If a uuid is not supplied one is generated when the lockout is added to the character

```lua
-- 5 minute lockout
eq.add_expedition_lockout_by_char_id(e.other:CharacterID(), "Anguish, the Fallen Palace", "Ture", 300)
```

## `Expedition*` get\_expedition\(\)

Returns an Expedition object corresponding to current zone if it's a valid dynamic zone instance being used for an expedition, otherwise returns nullptr.

```lua
local dz = eq.get_expedition()
if dz.valid then
  eq.debug("Current zone is an expedition's dz instance")
end
```

## `Expedition*` get\_expedition\_by\_char\_id\(int character\_id\)

Returns the character's current Expedition if the character is in one, otherwise returns nullptr.

```lua
local dz = eq.get_expedition_by_char_id(e.other:CharacterID())
if dz.null then
  eq.debug("Character is not in an expedition")
end
```

## `Expedition*` get\_expedition\_by\_dz\_id\(int dz\_id\)

Returns the Expedition for the specified dynamic zone id if it's valid, otherwise returns nullptr.

```lua
local dz = eq.get_expedition_by_dz_id(dz_id)
if not dz.valid then
  eq.debug("that dynamic zone id is not in use by an expedition")
end
```

## `Expedition*` get\_expedition\_by\_zone\_instance\(int zone\_id, int instance\_id\)

Returns the Expedition for the specified zone and instance id if it's a valid dynamic zone instance, otherwise returns nullptr.

```lua
local dz = eq.get_expedition_by_zone_instance(zone_id, instance_id)
if dz.valid then
  eq.debug("that zone instance is an expedition dynamic zone")
end
```

## `LUA_TTABLE` get\_expedition\_lockout\_by\_char\_id\(int character\_id, string expedition\_name, string event\_name\)

Returns a Lua table containing lockout details.

The table contains two hash keys `remaining` and `uuid`. T\["remaining"\] is the seconds remaining on the lockout and T\["uuid"\] is the expedition uuid the lockout was originally assigned in.

If the character doesn't have a lockout for the specified expedition event an empty table is returned.

```lua
local lockout = eq.get_expedition_lockout_by_char_id(
  e.other:CharacterID(), "Anguish, the Fallen Palace", "Ture")

if next(lockout) == nil then
  eq.debug("Character does not have a Ture lockout")
else
  eq.debug(string.format("Lockout remaining: (%s)s uuid: (%s)", lockout.remaining, lockout.uuid))
end
```

## `LUA_TTABLE` get\_expedition\_lockouts\_by\_char\_id\(int character\_id, string expedition\_name = ""\)

Returns a Lua table containing all of character's current lockouts keyed by expedition name. If an optional `expedition_name` is passed in it returns a table of lockouts keyed by event name for that expedition.

Each character lockout contains a hash table with `remaining` and `uuid` keys

If the character has no lockouts, the table is empty

```lua
-- without filter on expedition name
local lockouts = eq.get_expedition_lockouts_by_char_id(e.other:CharacterID())

for expedition_name,event_lockouts in pairs(lockouts) do
  for event_name,lockout in pairs(event_lockouts) do
    eq.debug(string.format("expedition: (%s) event: (%s) remaining: (%s)s uuid: (%s)",
      expedition_name, event_name, lockout.remaining, lockout.uuid))
  end
end
```

```lua
-- with filter on expedition name
local lockouts = eq.get_expedition_lockouts_by_char_id(e.other:CharacterID(), "Anguish, the Fallen Palace")

for event_name,lockout in pairs(lockouts) do
  eq.debug(string.format("event: (%s) remaining: (%s)s uuid: (%s)",
    event_name, lockout.remaining, lockout.uuid))
end
```

## `void` remove\_all\_expedition\_lockouts\_by\_char\_id\(int character\_id, string expedition\_name = ""\)

Removes all expedition lockouts from the character, optionally filtered on an expedition name

```lua
-- removes all lockouts
eq.remove_expedition_lockout_by_char_id(e.other:CharacterID())
-- removes all anguish lockouts
eq.remove_expedition_lockout_by_char_id(e.other:CharacterID(), "Anguish, the Fallen Palace")
```

## `void` remove\_expedition\_lockout\_by\_char\_id\(int character\_id, string expedition\_name, string event\_name\)

Removes the specified lockout from the character. No-op if the character doesn't have the lockout

```lua
eq.remove_expedition_lockout_by_char_id(e.other:CharacterID(), "Anguish, the Fallen Palace", "Ture")
```

