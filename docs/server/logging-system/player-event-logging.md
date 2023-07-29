---
tags:
  - Discord
  - Logs
  - Logging
---


# Player Event Logging

For the longest time, player audit logging or event logging of any type has been sub-par. It is difficult to triage when a player lost an item because we have no adequate out of the box logging. It is difficult to know get visibility into player activity to understand how things are being used or misused

## Goals

* As a server operator, be able to see a lineage of events of a player
* As a server operator, be able to see events by type
* As a server operator, be able to see events by zone
* As a server operator, be able to see record event specific data in a non-structured way but be able to query it in a flexible semi-structured way
* Have player events be recorded in a highly robust, high performing, low overhead way
* Have player events be turned on and off by type
* Have player events automatically truncate by type so data retention can be tweaked and tuned to the server operators unique needs
* Have the player event creation pipeline be easy to extend and adapt from a server operator perspective
* Be able to process events within World or within Queryserv as a dedicated events processor on a dedicated server if the operator so chooses

## Features

* Roughly 40 implemented player events
* One single event table, which heavily simplifies creating events and querying events in a lineage along with other events. Querying the JSON event data can get interesting when running reporting but it is all flexible enough.
* Highly robust event pipeline where World (Default) or QueryServ can process events in batch. This allows zone processes to very minimally and lightly process event creation and the event processor can flush events at a configurable interval (Default 5s) Rule "Logging:BatchPlayerEventProcessIntervalSeconds". The event processor runs in its own dedicated thread and uses a mutex lock to interact with the queue that gets mutated between the main thread and the processing thread
* Highly performant event creation. Events do not incur overhead at the Zone level if they are not enabled
* Highly readable code. All player events are represented as `structs` which are serialized into JSON when the event reaches the database
* Discord Webhook messages unique to each event are supported and can be turned on per-event and routed to Channels of your choosing, player facing or admin facing. Consider that some events could be too noisy. Discord webhook messages are routed by Universal Chat Service (Server)
* New player events are automatically injected when they are added by server developers
* Deprecated player events are automatically removed when marked as deprecated or unimplemented
* Data retention policies. Every event is trimmed hourly at given data retention intervals based on the age of the event. Most events defaults are to be stored for a week. Configurable by the server operator.

![image](https://user-images.githubusercontent.com/3319450/216811907-d7db6aed-5893-41f9-a96f-1c871c0eaa83.png)

![image](https://user-images.githubusercontent.com/3319450/216809359-42a07e5e-2f81-416d-991b-42da3d338750.png)

![image](https://user-images.githubusercontent.com/3319450/216809336-10f6e0d1-4f6b-4ded-8dd4-e64345131bf8.png)

## Deprecates

* The old `hackers` table is now deprecated. These are now their own `POSSIBLE_HACK` events since all of this was player contextual logging to begin with
* The old `eventlog` table is now deprecated. All events covered in here are also covered in the player event logging system and far more

## Table Schema(s)

[player_event_logs](../../schema/admin/player_event_logs.md)

#### Indexes

```
KEY `event_created_at` (`event_type_id`,`created_at`),
KEY `zone_id` (`zone_id`),
KEY `character_id` (`character_id`,`zone_id`) USING BTREE,
KEY `created_at` (`created_at`)
```

### player_event_log_settings

[player_event_log_settings](../../schema/admin/player_event_log_settings.md)

## Currently Handled Events

These are the events currently supported at the time of this writing. It's very likely more will be added later on and you can find the most recent events in https://github.com/EQEmu/Server/blob/master/common/events/player_events.h

```
1	GM Command
2	Zoning
3	AA Gain
4	AA Purchase
5	Forage Success
6	Forage Failure
7	Fish Success
8	Fish Failure
9	Item Destroy
10	Went Online
11	Went Offline
12	Level Gain
13	Level Loss
14	Loot Item
15	Merchant Purchase
16	Merchant Sell
21	Groundspawn Pickup
22	NPC Handin
23	Skill Up
24	Task Accept
25	Task Update
26	Task Complete
27	Trade
29	Say
30	Rez Accepted
31	Death
32	Combine Failure
33	Combine Success
34	Dropped Item
35	Split Money
38	Trader Purchase
39	Trader Sell
42	Discover Item
43	Possible Hack
44	Killed NPC
45	Killed Named NPC
46	Killed Raid NPC
```

## Credits

The system was designed by Akkadius and co-designed by Kinglykrab. Aead and Kinglykrab also helped with the implementation of several events and many Discord formatters. Thank you both for your help <3

## Quest Hand-In Events

In order to capture player hand-ins, you will need special plugin capture code that is included in PEQ Quest PR https://github.com/ProjectEQ/projecteqquests/pull/1372/files

Hand-ins may need further work to capture other code paths

## Player Event Log Explorer (Spire)

A system like this is best served when it can be seen as to what you can do with it with the right tooling. Here is examples of key explorer features. The release of this will likely shortly trail when the server feature is released.

### Rich Data Viewing

Spire translates event data into viewable NPC, Item, Spell etc. cards to have contextual insight into the events and what they contain.

![image](https://user-images.githubusercontent.com/3319450/216812922-547efee0-4a84-4b1d-8637-1a2d13283cba.png)


![image](https://user-images.githubusercontent.com/3319450/216810318-32e0e8f3-1392-497e-a14b-55f6585e7edd.png)

### Event Raw

Any event you can see the richly formatted event or you can inspect the raw event data

![event-raw](https://user-images.githubusercontent.com/3319450/216810779-b14bbb99-7492-497e-afdf-407843ab1275.gif)


### Flexible Filtering

You can filter by event type, zone_id, character_id which are top level filters always available and indexed regardless of event type

![event-filter](https://user-images.githubusercontent.com/3319450/216810442-62040911-f17c-4b41-b032-ec00a440c49d.gif)

### Advanced Event Data Filtering

You can filter by JSON event data

![event-data-filter](https://user-images.githubusercontent.com/3319450/216810590-9802cfb5-a619-4e08-ab0e-342376e4044b.gif)

You can even filter by deeply nested data. For example if you're trying to search for a certain item that was traded by doing a nested wildcard search

![event-deep-filter](https://user-images.githubusercontent.com/3319450/216810644-dea5d9b6-3a8c-4936-afaf-efae03fb1fe5.gif)

## Discord Webhooks

To use Discord webhooks, you will need to make sure the event is handled to begin with in https://github.com/EQEmu/Server/blob/master/common/events/player_event_logs.cpp#L269

You will need to add a webhook to your [discord_webhooks](../../schema/admin/discord_webhooks.md) table. See [Discord Logging](discord-logging.md) for more information

Once you have a webhook added to your table, you can use the webhook id in the [player_event_log_settings](../../schema/admin/player_event_log_settings.md) table to enable the event to be sent to Discord.

## Developer Docs

### Where's the Code?

Most of the code managing player events is located @

```
./common/events
├── player_event_discord_formatter.cpp
├── player_event_discord_formatter.h
├── player_event_logs.cpp
├── player_event_logs.h
└── player_events.h
```

### How Do I Add a New Event?

First you need an event defined. You do that in **player_events.h** in both `EventType` and `EventName`

You'll also need a struct representing the event data that you want to exist in the database, that is also defined in **player_events.h**

It is important to implement the cereal serialization fields because that is what is used to automatically handle transporting the data from zone to world and ultimately into JSON format in the database. Much of this is handled automatically for the developer.

```cpp
struct SayEvent {
	std::string message;
	std::string target;

	// cereal
	template<class Archive>
	void serialize(Archive &ar)
	{
		ar(
			CEREAL_NVP(message),
			CEREAL_NVP(target)
		);
	}
};
```

You will also need to find the place in the code that you would like to record the event. There are two convenient macros that wrap a lot of other logic to keep things performant

```cpp
RecordPlayerEventLog(event_type, event_data) // if client object is implied
RecordPlayerEventLogWithClient(c, event_type, event_data) // if you need to pass in the client object
```
For the message to properly show up in Discord, it will need a proper formatter to format the event data in a visually pleasing way.

Implement a formatter in **./common/player_event_discord_formatter.cpp** that looks similar to

```cpp
std::string PlayerEventDiscordFormatter::FormatDiscoverItemEvent(  
   const PlayerEvent::PlayerEventContainer &c,  
   const PlayerEvent::DiscoverItemEvent &e  
)
```

You can copy other formatter signatures fairly easy to understand what is happening. This will need to be hooked up in the formatter switch in `PlayerEventLogs::GetDiscordPayloadFromEvent` in **common/events/player_event_logs.cpp**

#### Example Diff

Below are two example of an event being implemented / added. These only contain what is required to add events, not format Discord messages

##### More Complex

```cpp
diff --git a/common/events/player_events.h b/common/events/player_events.h  
index af0198ab9..6c988bb41 100644  
--- a/common/events/player_events.h  
+++ b/common/events/player_events.h  
@@ -42,7 +42,7 @@ namespace PlayerEvent {  
      COMBINE_FAILURE,  
      COMBINE_SUCCESS,  
      DROPPED_ITEM,  
-     SPLIT_MONEY, // unimplemented  
+     SPLIT_MONEY,  
      DZ_JOIN, // unimplemented  
      DZ_LEAVE, // unimplemented  
      TRADER_PURCHASE, // unimplemented  
@@ -574,6 +574,27 @@ namespace PlayerEvent {  
         );  
      }  
   };  
+  
+  struct SplitMoneyEvent {  
+     uint32 copper;  
+     uint32 silver;  
+     uint32 gold;  
+     uint32 platinum;  
+     uint64 player_money_balance;  
+  
+     // cereal  
+     template<class Archive>  
+     void serialize(Archive &ar)  
+     {  
+        ar(  
+           CEREAL_NVP(copper),  
+           CEREAL_NVP(silver),  
+           CEREAL_NVP(gold),  
+           CEREAL_NVP(platinum),  
+           CEREAL_NVP(player_money_balance)  
+        );  
+     }  
+  };  
 }  
   
 #endif //EQEMU_PLAYER_EVENTS_H  
diff --git a/zone/groups.cpp b/zone/groups.cpp  
index 910d5b21c..01d7c58bb 100644  
--- a/zone/groups.cpp  
+++ b/zone/groups.cpp  
@@ -26,6 +26,7 @@  
 #include "../common/strings.h"  
 #include "worldserver.h"  
 #include "string_ids.h"  
+#include "../common/events/player_event_logs.h"  
   
 extern EntityList entity_list;  
 extern WorldServer worldserver;  
@@ -177,6 +178,18 @@ void Group::SplitMoney(uint32 copper, uint32 silver, uint32 gold, uint32 platinu  
            true  
         );  
   
+        if (player_event_logs.IsEventEnabled(PlayerEvent::SPLIT_MONEY)) {  
+           auto e = PlayerEvent::SplitMoneyEvent{  
+              .copper = copper_split,  
+              .silver = silver_split,  
+              .gold = gold_split,  
+              .platinum = platinum_split,  
+              .player_money_balance = members[i]->CastToClient()->GetCarriedMoney(),  
+           };  
+  
+           RecordPlayerEventLogWithClient(members[i]->CastToClient(), PlayerEvent::SPLIT_MONEY, e);  
+        }  
+  
         members[i]->CastToClient()->MessageString(  
            Chat::MoneySplit,  
            YOU_RECEIVE_AS_SPLIT,  
diff --git a/zone/raids.cpp b/zone/raids.cpp  
index b27cae0e0..5fbc543c1 100644  
--- a/zone/raids.cpp  
+++ b/zone/raids.cpp  
@@ -17,6 +17,7 @@  
 */  
   
 #include "../common/strings.h"  
+#include "../common/events/player_event_logs.h"  
   
 #include "client.h"  
 #include "entity.h"  
@@ -814,6 +815,19 @@ void Raid::SplitMoney(uint32 gid, uint32 copper, uint32 silver, uint32 gold, uin  
            true  
         );  
   
+        if (player_event_logs.IsEventEnabled(PlayerEvent::SPLIT_MONEY)) {  
+           auto e = PlayerEvent::SplitMoneyEvent{  
+              .copper = copper_split,  
+              .silver = silver_split,  
+              .gold = gold_split,  
+              .platinum = platinum_split,  
+              .player_money_balance = members[i].member->GetCarriedMoney(),  
+           };  
+  
+           RecordPlayerEventLogWithClient(members[i].member, PlayerEvent::SPLIT_MONEY, e);  
+        }  
+  
+  
         members[i].member->MessageString(  
            Chat::MoneySplit,  
            YOU_RECEIVE_AS_SPLIT,
```

#### Simpler

```diff
diff --git a/common/events/player_events.h b/common/events/player_events.h  
index d9c3a98d4..07ee5aee5 100644  
--- a/common/events/player_events.h  
+++ b/common/events/player_events.h  
@@ -17,8 +17,8 @@ namespace PlayerEvent {  
      FISH_SUCCESS,  
      FISH_FAILURE,  
      ITEM_DESTROY,  
-     WENT_ONLINE, // unimplemented  
-     WENT_OFFLINE, // unimplemented  
+     WENT_ONLINE,  
+     WENT_OFFLINE,  
      LEVEL_GAIN, // unimplemented  
      LEVEL_LOSS, // unimplemented  
      LOOT_ITEM, // unimplemented  
diff --git a/zone/client_packet.cpp b/zone/client_packet.cpp  
index ae75922c5..fbf7eca3b 100644  
--- a/zone/client_packet.cpp  
+++ b/zone/client_packet.cpp  
@@ -784,6 +784,7 @@ void Client::CompleteConnect()  
   /* This sub event is for if a player logs in for the first time since entering world. */  
   if (firstlogon == 1) {  
      parse->EventPlayer(EVENT_CONNECT, this, "", 0);  
+     RecordPlayerEventLog(PlayerEvent::WENT_ONLINE, PlayerEvent::EmptyEvent{});  
      /* QS: PlayerLogConnectDisconnect */  
      if (RuleB(QueryServ, PlayerLogConnectDisconnect)) {  
         std::string event_desc = StringFormat("Connect :: Logged into zoneid:%i instid:%i", GetZoneID(), GetInstanceID());  
diff --git a/zone/client_process.cpp b/zone/client_process.cpp  
index c6ff4e02f..2c5df69e8 100644  
--- a/zone/client_process.cpp  
+++ b/zone/client_process.cpp  
@@ -55,6 +55,7 @@  
 #include "zone.h"  
 #include "zonedb.h"  
 #include "../common/zone_store.h"  
+#include "../common/events/player_event_logs.h"  
   
 extern QueryServ* QServ;  
 extern Zone* zone;  
@@ -184,6 +185,7 @@ bool Client::Process() {  
         SetDynamicZoneMemberStatus(DynamicZoneMemberStatus::Offline);  
   
         parse->EventPlayer(EVENT_DISCONNECT, this, "", 0);  
+        RecordPlayerEventLog(PlayerEvent::WENT_OFFLINE, PlayerEvent::EmptyEvent{});  
   
         return false; //delete client  
      }  
@@ -693,6 +695,7 @@ void Client::OnDisconnect(bool hard_disconnect) {  
         MyRaid->MemberZoned(this);  
   
      parse->EventPlayer(EVENT_DISCONNECT, this, "", 0);  
+     RecordPlayerEventLog(PlayerEvent::WENT_OFFLINE, PlayerEvent::EmptyEvent{});  
   
      /* QS: PlayerLogConnectDisconnect */  
      if (RuleB(QueryServ, PlayerLogConnectDisconnect)){
``` 

### Default Event Settings

Default event settings are defined in `PlayerEventLogs::SetSettingsDefaults`

```cpp
void PlayerEventLogs::SetSettingsDefaults()  
{  
   m_settings[PlayerEvent::GM_COMMAND].event_enabled = 1;  
   m_settings[PlayerEvent::ZONING].event_enabled = 1;  
   m_settings[PlayerEvent::AA_GAIN].event_enabled = 1;  
   m_settings[PlayerEvent::AA_PURCHASE].event_enabled = 1;
   // ...truncated
```

This determines if the event is enabled by default and sets the defaults before it gets injected into the operators database automatically on next code update.