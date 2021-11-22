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

```perl
my $member_has_lockout = 0;

my $raid = $client->GetRaid();
my $group = $client->GetGroup();

if ($raid) {
  $member_has_lockout = $raid->DoesAnyMemberHaveExpeditionLockout("foo", "bar");
} elsif ($group) {
  $member_has_lockout = $group->DoesAnyMemberHaveExpeditionLockout("foo", "bar");
} else {
  $member_has_lockout = $client->HasExpeditionLockout("foo", "bar");
}

quest::debug("member has lockout: [$member_has_lockout]");
```
