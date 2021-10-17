# Group/Raid Methods

## Group methods

### `bool` DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name, int max_check_count = 0)

Returns true if any group member has the specified lockout.

Passing a non-zero `max_check_count` will limit to checking that number of group members in no guaranteed order.

See Raid::DoesAnyMemberHaveExpeditionLockout for example. \


## Raid methods

### `bool` DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name, int max_check_count = 0)

Returns true if any raid member has the specified lockout.

Passing a non-zero `max_check_count` will limit to checking that number of raid members. The order that raid members are checked is leader first, followed by raid member group number, and finally ungrouped members. The order of group members within raid groups is not guaranteed to be the same as displayed.

```lua
local member_has_lockout = false

if e.other:GetRaid().valid then
  member_has_lockout = e.other:GetRaid():DoesAnyMemberHaveExpeditionLockout("foo", "bar")
elseif e.other:GetGroup().valid then
  member_has_lockout = e.other:GetGroup():DoesAnyMemberHaveExpeditionLockout("foo", "bar")
else
  member_has_lockout = e.other:HasExpeditionLockout("foo", "bar")
end

if member_has_lockout then
  e.self:Say("Insert custom quest dialogue if someone has a lockout")
else
  local expedition = e.other:CreateExpedition("zone", 0, 36000, "foo", 6, 54)
end
```
