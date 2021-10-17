# Quest API (Perl)

> Note: `HASHREF` designates a _**reference**_ to `SVt_PVHV` hash type

## Global methods

| return type | function |
| ---: | :--- |
| void | add_expedition_lockout_all_clients(string expedition_name, string event_name, int seconds, string uuid = "") |
| void | add_expedition_lockout_by_char_id(int character_id, string expedition_name, string event_name, int seconds, string uuid = "") |
| Expedition* | get_expedition() |
| Expedition* | get_expedition_by_char_id(int character_id) |
| Expedition* | get_expedition_by_dz_id(int dz_id) |
| Expedition* | get_expedition_by_zone_instance(int zone_id, int instance_id) |
| HASHREF | get_expedition_lockout_by_char_id(int character_id, string expedition_name, string event_name) |
| HASHREF | get_expedition_lockouts_by_char_id(int character_id, string expedition_name = "") |
| void | remove_all_expedition_lockouts_by_char_id(int character_id, string expedition_name = "") |
| void | remove_expedition_lockout_by_char_id(int character_id, string expedition_name, string event_name) |

## Client methods

| return type | function |
| ---: | :--- |
| void | AddExpeditionLockout(string expedition_name, string event_name, uint32 seconds, string uuid = "") |
| void | AddExpeditionLockoutDuration(string expedition_name, string event_name, int seconds, string uuid = "") |
| Expedition* | CreateExpedition(string zone_short_name, int zone_version, int duration, string expedition_name, int min_players, int max_players, bool disable_messages = false) |
| Expedition* | GetExpedition() |
| HASHREF | GetExpeditionLockouts(string expedition_name = "") |
| string | GetLockoutExpeditionUUID(string expedition_name, string event_name) |
| bool | HasExpeditionLockout(string expedition_name, string event_name) |
| void | MovePCDynamicZone(int zone_id, int zone_version = -1, bool msg_if_invalid = true) |
| void | MovePCDynamicZone(string zone_short_name, int zone_version = -1, bool msg_if_invalid = true) |
| void | RemoveAllExpeditionLockouts(string expedition_name = "") |
| void | RemoveExpeditionLockout(string expedition_name, string event_name) |

## Group methods

| return type | function |
| ---: | :--- |
| bool | DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name, int max_member_check_count = 0) |

## Raid methods

| return type | function |
| ---: | :--- |
| bool | DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name, int max_member_check_count = 0) |

## Expedition methods

| return type | function |
| ---: | :--- |
| void | AddLockout(string event_name, uint32 seconds_duration) |
| void | AddLockoutDuration(string event_name, int seconds, bool members_only = true) |
| void | AddReplayLockout(int seconds_duration) |
| void | AddReplayLockoutDuration(int seconds_duration, bool members_only = true) |
| uint32 | GetDynamicZoneID() |
| uint32 | GetID() |
| int | GetInstanceID() |
| string | GetLeaderName() |
| HASHREF | GetLockouts() |
| string | GetLootEventByNPCTypeID(uint32 npc_type_id) |
| string | GetLootEventBySpawnID(uint32 spawn_id) |
| uint32 | GetMemberCount() |
| HASHREF | GetMembers() |
| string | GetName() |
| int | GetSecondsRemaining() |
| string | GetUUID() |
| int | GetZoneID() |
| string | GetZoneName() |
| int | GetZoneVersion() |
| bool | HasLockout(string event_name) |
| bool | HasReplayLockout() |
| bool | IsLocked() |
| void | RemoveCompass() |
| void | RemoveLockout(string event_name) |
| void | SetCompass(int zone_id, float x, float y, float z) |
| void | SetCompass(string zone_short_name, float x, float y, float z) |
| void | SetLocked(bool value, ExpeditionLockMessage lock_msg = ExpeditionLockMessage::None, uint32 msg_color = Chat::Yellow) |
| void | SetLootEventByNPCTypeID(uint32_t npc_type_id, string event_name) |
| void | SetLootEventBySpawnID(uint32_t spawn_id, string event_name) |
| void | SetReplayLockoutOnMemberJoin(bool value) |
| void | SetSafeReturn(uint32_t zone_id, float x, float y, float z, float heading) |
| void | SetSafeReturn(string zone_short_name, float x, float y, float z, float heading) |
| void | SetZoneInLocation(float x, float y, float z, float heading) |
| void | SetSecondsRemaining(uint32 seconds_remaining) |
| void | UpdateLockoutDuration(string event_name, uint32_t seconds, bool members_only = true) |

## Expedition constants <a id="expedition-constants"></a>

#### `ExpeditionLockMessage` <a id="expedition-lock-message"></a>

| constant | value | lock message |
| :---: | :---: | :--- |
| None | 0 |  |
| Close | 1 | "Your expedition is nearing its close. You cannot bring any additional people into your expedition at this time." |
| Begin | 2 | "The trial has begun. You cannot bring any additional people into your expedition at this time." |

