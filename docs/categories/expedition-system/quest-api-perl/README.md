# Quest API \(Perl\)

> Note: `HASHREF` designates a _**reference**_ to `SVt_PVHV` hash type

## Global methods

| return type | function |
| ---: | :--- |
| void | add\_expedition\_lockout\_all\_clients\(string expedition\_name, string event\_name, int seconds, string uuid = ""\) |
| void | add\_expedition\_lockout\_by\_char\_id\(int character\_id, string expedition\_name, string event\_name, int seconds, string uuid = ""\) |
| Expedition\* | get\_expedition\(\) |
| Expedition\* | get\_expedition\_by\_char\_id\(int character\_id\) |
| Expedition\* | get\_expedition\_by\_dz\_id\(int dz\_id\) |
| Expedition\* | get\_expedition\_by\_zone\_instance\(int zone\_id, int instance\_id\) |
| HASHREF | get\_expedition\_lockout\_by\_char\_id\(int character\_id, string expedition\_name, string event\_name\) |
| HASHREF | get\_expedition\_lockouts\_by\_char\_id\(int character\_id, string expedition\_name = ""\) |
| void | remove\_all\_expedition\_lockouts\_by\_char\_id\(int character\_id, string expedition\_name = ""\) |
| void | remove\_expedition\_lockout\_by\_char\_id\(int character\_id, string expedition\_name, string event\_name\) |

## Client methods

| return type | function |
| ---: | :--- |
| void | AddExpeditionLockout\(string expedition\_name, string event\_name, uint32 seconds, string uuid = ""\) |
| void | AddExpeditionLockoutDuration\(string expedition\_name, string event\_name, int seconds, string uuid = ""\) |
| Expedition\* | CreateExpedition\(string zone\_short\_name, int zone\_version, int duration, string expedition\_name, int min\_players, int max\_players, bool disable\_messages = false\) |
| Expedition\* | GetExpedition\(\) |
| HASHREF | GetExpeditionLockouts\(string expedition\_name = ""\) |
| string | GetLockoutExpeditionUUID\(string expedition\_name, string event\_name\) |
| bool | HasExpeditionLockout\(string expedition\_name, string event\_name\) |
| void | MovePCDynamicZone\(int zone\_id, int zone\_version = -1, bool msg\_if\_invalid = true\) |
| void | MovePCDynamicZone\(string zone\_short\_name, int zone\_version = -1, bool msg\_if\_invalid = true\) |
| void | RemoveAllExpeditionLockouts\(string expedition\_name = ""\) |
| void | RemoveExpeditionLockout\(string expedition\_name, string event\_name\) |

## Group methods

| return type | function |
| ---: | :--- |
| bool | DoesAnyMemberHaveExpeditionLockout\(string expedition\_name, string event\_name, int max\_member\_check\_count = 0\) |

## Raid methods

| return type | function |
| ---: | :--- |
| bool | DoesAnyMemberHaveExpeditionLockout\(string expedition\_name, string event\_name, int max\_member\_check\_count = 0\) |

## Expedition methods

| return type | function |
| ---: | :--- |
| void | AddLockout\(string event\_name, uint32 seconds\_duration\) |
| void | AddLockoutDuration\(string event\_name, int seconds, bool members\_only = true\) |
| void | AddReplayLockout\(int seconds\_duration\) |
| void | AddReplayLockoutDuration\(int seconds\_duration, bool members\_only = true\) |
| uint32 | GetDynamicZoneID\(\) |
| uint32 | GetID\(\) |
| int | GetInstanceID\(\) |
| string | GetLeaderName\(\) |
| HASHREF | GetLockouts\(\) |
| string | GetLootEventByNPCTypeID\(uint32 npc\_type\_id\) |
| string | GetLootEventBySpawnID\(uint32 spawn\_id\) |
| uint32 | GetMemberCount\(\) |
| HASHREF | GetMembers\(\) |
| string | GetName\(\) |
| int | GetSecondsRemaining\(\) |
| string | GetUUID\(\) |
| int | GetZoneID\(\) |
| string | GetZoneName\(\) |
| int | GetZoneVersion\(\) |
| bool | HasLockout\(string event\_name\) |
| bool | HasReplayLockout\(\) |
| bool | IsLocked\(\) |
| void | RemoveCompass\(\) |
| void | RemoveLockout\(string event\_name\) |
| void | SetCompass\(int zone\_id, float x, float y, float z\) |
| void | SetCompass\(string zone\_short\_name, float x, float y, float z\) |
| void | SetLocked\(bool value, ExpeditionLockMessage lock\_msg = ExpeditionLockMessage::None, uint32 msg\_color = Chat::Yellow\) |
| void | SetLootEventByNPCTypeID\(uint32\_t npc\_type\_id, string event\_name\) |
| void | SetLootEventBySpawnID\(uint32\_t spawn\_id, string event\_name\) |
| void | SetReplayLockoutOnMemberJoin\(bool value\) |
| void | SetSafeReturn\(uint32\_t zone\_id, float x, float y, float z, float heading\) |
| void | SetSafeReturn\(string zone\_short\_name, float x, float y, float z, float heading\) |
| void | SetZoneInLocation\(float x, float y, float z, float heading\) |
| void | SetSecondsRemaining\(uint32 seconds\_remaining\) |
| void | UpdateLockoutDuration\(string event\_name, uint32\_t seconds, bool members\_only = true\) |

## Expedition constants <a id="expedition-constants"></a>

#### `ExpeditionLockMessage` <a id="expedition-lock-message"></a>

| constant | value | lock message |
| :---: | :---: | :--- |
| None | 0 |  |
| Close | 1 | "Your expedition is nearing its close. You cannot bring any additional people into your expedition at this time." |
| Begin | 2 | "The trial has begun. You cannot bring any additional people into your expedition at this time." |

