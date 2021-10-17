# Expedition System

The expedition system enables the ability to create Dynamic Zone instances managed by the client's Expedition interface window. This system allows clients to add and remove members of a Dynamic Zone while providing built-in event lockout timer support.

The system is currently quest driven and provides quest APIs to create zone instances without the burden of manually tracking and validating lockout conflicts among characters. The system should have zero impact on servers that don't make use of it.

## Expedition requests

When a non-solo expedition request occurs, the leader of the raid or group becomes the lockout authority. If any members of the party have a lockout for the expedition that the leader doesn't already have the expedition request is denied. Denied expedition requests message the party leader \(not the member that requested it\).

If the request is successful, all party members are added to the expedition and it internally inherits the leader's current lockouts. These inherited lockouts are not assigned to members of the expedition. They are used to determine if players may be added to the expedition later based on their lockouts and to allow quests to choose which expedition events should be enabled or disabled.

## Dynamic Zones

Dynamic Zones are a wrapper for instances that extends them with some extra data. In the future these will probably be further modified for use by other systems that make use of dynamic zones.

## Commands

The `#dz` command is provided for GMs and developers to perform various operations. Use the command without arguments in game to see usage \(minimum 80 access level\).

The `#dzkickplayers` command is provided for players so that expedition leaders using pre-RoF clients can have the `/kickplayers exp` functionality to remove all members of an expedition.

## Rules

The following rules are available for server customization

| Category | Rule Name | Type | Default | Description |
| :--- | :--- | :--- | :--- | :--- |
| Expedition | MinStatusToBypassPlayerCountRequirements | int | 80 | Minimum GM status to bypass minimum player requirements for Expedition creation |
| Expedition | EmptyDzShutdownEnabled | bool | true | Enable early instance shutdown after last member of expedition removed |
| Expedition | EmptyDzShutdownDelaySeconds | int | 1500 | Seconds to set dynamic zone instance expiration if early shutdown enabled |
| Expedition | WorldExpeditionProcessRateMS | int | 6000 | Timer interval \(ms\) that world checks expedition states |
| Expedition | AlwaysNotifyNewLeaderOnChange | bool | false | Always notify clients when made expedition leader \(such as for going offline\). If false \(live-like\) new leaders are only notified when made leader via /dzmakeleader |
| Expedition | LockoutDurationMultiplier | float | 1.0 | Multiplies lockout duration by this value when new lockouts are added |
| Expedition | EnableInDynamicZoneStatus | bool | false | Enables the 'In Dynamic Zone' member status in expedition window. If false \(live-like\) players inside the dz will show as 'Online' |
| DynamicZone | ClientRemovalDelayMS | int | 60000 | Delay \(ms\) until a client is teleported out of dynamic zone after being removed as member |

> Note: The `DynamicZone::ClientRemovalDelayMS` rule is in its own category for possible future systems that may make use of dynamic zones. The empty dz shutdown rules may be moved to this category in the future depending on design changes.

## Known issues

* Using /exit with an expedition invite window open may crash the client instead of exiting cleanly
* Invoking the dynamic zone switch list window after reloading the UI may crash clients
  * client bug that still exists on live
  * this window is currently unused since expeditions are the only system that use dynamic zones

