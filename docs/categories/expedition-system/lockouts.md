# Lockouts

There are two types of lockout replay timers that need to be distinguished

* Replay Timer - A replay timer on the expedition itself
* Event Timer - A replay timer on some event inside an expedition

## Replay Lockouts

`Replay Timer` lockouts use a hard-coded event name

> From live December 14, 2016 patch notes
>
> Dynamic zone replay times not associated to a particular event will now display 'Replay Timer' in the Event Name column of the expedition window rather than a blank space.

These are lockouts for requesting or joining an expedition. Any party members with this type of lockout will cause an expedition creation request to be denied. Any members invited to an expedition that has had a `Replay Timer` lockout added will receive that lockout immediately with its original duration upon joining \(can be disabled via quest API\).

Players with a `Replay Timer` lockout may only be re-invited to the same expedition that the timer was originally assigned from. Players that have a `Replay Timer` from another expedition can not be added.

## Event Lockouts

Event lockout timers are lockouts for events inside an expedition.

Event lockouts added during an expedition are assigned to:

* The expedition internally \(non-inherited lockout\)
* All current members of the expedition
* Any non-expedition clients currently inside the expedition's Dynamic Zone

A player that has an event lockout that the expedition doesn't already have \(either inherited or not\) may not be added until the expedition completes that event and receives the lockout.

New members added to an expedition do not receive these non-inherited event lockouts immediately on joining. These lockouts are only added to new players when they zone into the expedition's dynamic zone instance. Players receive a warning for any lockouts they're missing that will be added when invited.

If a player already has an event lockout from another expedition and is added to an expedition after the event is completed inside it, that player's lockout timer will be synchronized with the expedition's lockout once they zone in.

> Character lockouts for the expedition will only be synchronized if the character's lockout timer is shorter than the expedition's timer

## Implementation Details

When an expedition creation request is made the lockouts of the requester's group/raid leader \(or the requester if solo\) are used for validation. If any member has an additional lockout for the expedition that the leader doesn't then the request fails.

If the request doesn't fail the expedition inherits the lockouts from the leader and makes a copy of them internally. These are the lockouts that quests can use to determine which events should be available inside the dz through initialization scripts.

Lockouts added by quests during an expedition are added internally to the expedition as well as to all current expedition members. Clients inside the dz that are not members of the expedition will also receive the lockouts to prevent exploiting. These lockouts store the UUID of the current expedition so they can be distinguished from inherited lockouts.

When a character is invited to an expedition the character's lockouts are compared with the internal expedition lockouts to determine if they may join. If the character has any lockouts that the expedition doesn't then the invite will fail.

When a character zones into an expedition's dynamic zone their lockouts are synchronized with the expedition's internal lockouts. Only lockouts added during the current expedition are synchronized. The character will receive new lockouts for any they didn't already have with the current expedition's UUID.

If the character already has an event lockout from another expedition \(the character's lockout UUID differs from the current expedition's UUID\) then their lockout timer for that event is updated to the expedition's current expire time but the character lockout UUID remains unchanged.

> Updating character lockout timers on dz entry even when the UUID is from another expedition is live behavior as of 2020-11-08.

