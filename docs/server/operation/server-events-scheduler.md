The server has currently two process driven schedulers

* World
  * Broadcast Messages (broadcast) (Stateless)
  * Reload World (reload_world) (Stateless)
* Zone
  * Rule Value Change (rule_change) (Stateful)
  * Content Flag Change (content_flag_change) (Stateful)
  * Hotzone Activate (hot_zone_activate) (Stateful)

Zone maintains state, world does not in the current implementation. For current every current zone event activation, there is an event deactivation event which undoes the state change.

World does not maintain state, it is used for single event triggers and supports cron expressions since the events can be kicked off on a routine basis

## Event Loading

When world or zone boots, you will see the initial load of events. Zone will also reload if World has detected any changes to any of the scheduled events in the database once a minute

```
[Zone] [Scheduler] Loaded Event (1) [Hotzone Rotation Day 1-10] type [hot_zone_activate] start [2/3/2021 3:1:00] end [2/16/2021 4:40:00] cron [] created []
[Zone] [Scheduler] Loaded Event (2) [Hotzone Rotation Day 11-20] type [hot_zone_activate] start [2/3/2021 15:1:00] end [2/15/2021 19:7:00] cron [] created []
[Zone] [Scheduler] Loaded Event (3) [Hotzone Rotation Day 21-33] type [hot_zone_activate] start [2/3/2021 3:1:00] end [2/16/2021 4:40:00] cron [] created []
[Zone] [Scheduler] Loaded Event (4) [Broadcast Test] type [broadcast] start [2/16/2021 22:56:00] end [2/16/2021 22:56:00] cron [] created []
[Zone] [Scheduler] Loaded Event (15) [Broadcast Test Cron] type [broadcast] start [2/16/2021 22:56:00] end [2/16/2021 22:56:00] cron [* * * * * *] created []
[Zone] [Scheduler] Loaded Event (16) [Experience Multiplier] type [rule_change] start [2/16/2021 22:56:00] end [2/16/2021 22:56:00] cron [] created []
[Zone] [Scheduler] Loaded Event (18) [Loot Multiplier] type [rule_change] start [2/16/2021 22:56:00] end [2/16/2021 22:56:00] cron [] created []
[Zone] [Scheduler] Loaded Event (19) [Enable Halloween] type [content_flag_change] start [2/15/2021 22:51:00] end [2/15/2021 22:53:00] cron [] created []
[Zone] [Scheduler] Loaded Event (20) [Reload World After Halloween] type [reload_world] start [2/15/2021 22:53:00] end [2/15/2021 22:54:00] cron [] created []
[Zone] [Scheduler] Loaded scheduled events [9]
```

## Event Changes

If new changes are made to the `server_scheduled_events` state table, world will pick it up every minute. If world detects that there are any data changes made to the table, it will notify all zone processes to reload their scheduled events. This makes it easy for server operators to add / modify events live while the server is online

## Server Scheduled Events Model

The database table model is represented as follows

```cpp
struct ServerScheduledEvents {
	int         id;
	std::string description;
	std::string event_type;
	std::string event_data;
	int         minute_start;
	int         hour_start;
	int         day_start;
	int         month_start;
	int         year_start;
	int         minute_end;
	int         hour_end;
	int         day_end;
	int         month_end;
	int         year_end;
	std::string cron_expression;
	std::string created_at;
	std::string deleted_at;
};
```

## Scheduling Methods

The way the scheduling schema works is by start end end times and vary slightly in triggering functionality depending on whether the event is stateless or stateful

**Repeat Events (Stateless)**

Stateless events can be triggered repeatedly such as a broadcast event by using a `cron_expression`. You can also schedule a one time event by setting the start time and end time one minute apart

**Repeat Events (Stateful)**

Events can start and end on intervals but you need to only use the unit of time (for example days) and let everything else be `0` for columns. A `0` value will imply the current time when the scheduler tics

Let's take a hotzone rotation for example as shown in the table below. If I want to rotate 3 sets of hotzones within the month every month I can do it like this

Hotzone Group #1 starts on day 1 and ends on day 10
Hotzone Group #2 starts on day 10 and ends on day 20
Hotzone Group #3 starts on day 20 and ends on day 33 (Yes day 33 never happens but its an outer bounds)

This rotation will continue perpetually because of the wildcard mask-type declaration of the time data. Same thing could be achieved using different time formats, or different months

```
select * from server_scheduled_events where event_type = 'hot_zone_activate';
+----+----------------------------+-------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------+--------------+------------+-----------+-------------+------------+------------+----------+---------+-----------+----------+-----------------+------------+------------+
| id | description                | event_type        | event_data                                                                                                                                                            | minute_start | hour_start | day_start | month_start | year_start | minute_end | hour_end | day_end | month_end | year_end | cron_expression | created_at | deleted_at |
+----+----------------------------+-------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------+--------------+------------+-----------+-------------+------------+------------+----------+---------+-----------+----------+-----------------+------------+------------+
|  1 | Hotzone Rotation Day 1-10  | hot_zone_activate | befallen,charasis,dalnir,frontiermtns,gukbottom,iceclad,lakeofillomen,northkarana,qey2hh1,soldunga,southro,wakening,podisease,velketor,akheva,riwwi,bothunder,poair   |            0 |          0 |         1 |           0 |          0 |          0 |        0 |      10 |         0 |        0 | NULL            | NULL       | NULL       |
|  2 | Hotzone Rotation Day 10-20 | hot_zone_activate | beholder,commons,eastkarana,frozenshadow,guktop,kaesora,lakerathe,oasis,runnyeye,soldungb,stonebrunt,warrens,chardok,ponightmare,thedeep,veksar,paw,kodtaz            |            0 |          0 |        10 |           0 |          0 |          0 |        0 |      20 |         0 |        0 | NULL            | NULL       | NULL       |
|  3 | Hotzone Rotation Day 20-33 | hot_zone_activate | burningwood,crystal,erudsxing,greatdivide,hole,kerraridge,najena,permafrost,sirens,southkarana,timorous,warslikswood,sebilis,nadox,griegsend,barindu,hohonora,powater |            0 |          0 |        20 |           0 |          0 |          0 |        0 |      33 |         0 |        0 | NULL            | NULL       | NULL       |
+----+----------------------------+-------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------+--------------+------------+-----------+-------------+------------+------------+----------+---------+-----------+----------+-----------------+------------+------------+
```

## Example Schedule Data

```
select * from server_scheduled_events;
+----+------------------------------+---------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------+--------------+------------+-----------+-------------+------------+------------+----------+---------+-----------+----------+-----------------+------------+------------+
| id | description                  | event_type          | event_data                                                                                                                                                            | minute_start | hour_start | day_start | month_start | year_start | minute_end | hour_end | day_end | month_end | year_end | cron_expression | created_at | deleted_at |
+----+------------------------------+---------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------+--------------+------------+-----------+-------------+------------+------------+----------+---------+-----------+----------+-----------------+------------+------------+
|  1 | Hotzone Rotation Day 1-10    | hot_zone_activate   | befallen,charasis,dalnir,frontiermtns,gukbottom,iceclad,lakeofillomen,northkarana,qey2hh1,soldunga,southro,wakening,podisease,velketor,akheva,riwwi,bothunder,poair   |            0 |          0 |         1 |           0 |          0 |          0 |        0 |      10 |         0 |        0 | NULL            | NULL       | NULL       |
|  2 | Hotzone Rotation Day 10-20   | hot_zone_activate   | beholder,commons,eastkarana,frozenshadow,guktop,kaesora,lakerathe,oasis,runnyeye,soldungb,stonebrunt,warrens,chardok,ponightmare,thedeep,veksar,paw,kodtaz            |            0 |          0 |        10 |           0 |          0 |          0 |        0 |      20 |         0 |        0 | NULL            | NULL       | NULL       |
|  3 | Hotzone Rotation Day 20-33   | hot_zone_activate   | burningwood,crystal,erudsxing,greatdivide,hole,kerraridge,najena,permafrost,sirens,southkarana,timorous,warslikswood,sebilis,nadox,griegsend,barindu,hohonora,powater |            0 |          0 |        20 |           0 |          0 |          0 |        0 |      33 |         0 |        0 | NULL            | NULL       | NULL       |
|  4 | Broadcast Test               | broadcast           | This is a system test                                                                                                                                                 |            0 |          0 |         0 |           0 |          0 |          0 |        0 |       0 |         0 |        0 | NULL            | NULL       | NULL       |
| 15 | Broadcast Test Cron          | broadcast           | This is a system cron test 2                                                                                                                                          |            0 |          0 |         0 |           0 |          0 |          0 |        0 |       0 |         0 |        0 | * * * * * *     | NULL       | NULL       |
| 16 | Experience Multiplier        | rule_change         | Character:FinalExpMultiplier=10                                                                                                                                       |            0 |          0 |         0 |           0 |          0 |          0 |        0 |       0 |         0 |        0 | NULL            | NULL       | NULL       |
| 18 | Loot Multiplier              | rule_change         | Zone:GlobalLootMultiplier=10                                                                                                                                          |            0 |          0 |         0 |           0 |          0 |          7 |       19 |      15 |         2 |     2021 | NULL            | NULL       | NULL       |
| 19 | Enable Halloween             | content_flag_change | peq_halloween                                                                                                                                                         |           51 |         22 |        15 |           2 |       2021 |         53 |       22 |      15 |         2 |     2021 | NULL            | NULL       | NULL       |
| 20 | Reload World After Halloween | reload_world        | NULL                                                                                                                                                                  |           53 |         22 |        15 |           2 |       2021 |         54 |       22 |      15 |         2 |     2021 | NULL            | NULL       | NULL       |
+----+------------------------------+---------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------+--------------+------------+-----------+-------------+------------+------------+----------+---------+-----------+----------+-----------------+------------+------------+
```

**Example of Content Flags being set for Halloween**

![image](https://user-images.githubusercontent.com/3319450/108163340-7ab46900-70b4-11eb-93c3-308a88e2f9c0.png)

**Example of Rule Event Activation (Detail Logs)**

```
[Zone] [Scheduler] Activating Event [Experience Multiplier] scheduled rule change, setting rule [Character:FinalExpMultiplier] to [10]
[Zone] [Rules] Set rule [Character:FinalExpMultiplier] to value [10]
[Zone] [Scheduler] [ValidateEventReadyToActivate] now_time [1613537959] start_time [1613537940] doesnt_end [true] end_time [94614463305936]
[Zone] [Scheduler] Activating Event [Loot Multiplier] scheduled rule change, setting rule [Zone:GlobalLootMultiplier] to [10]
```

When event deactivates

```
[Zone] [Scheduler] Deactivating event [Loot Multiplier] resetting rules to normal
```

**Example of Broadcasts Being sent from World**

Both the cron expression event and regular event from the table above shown in the image below

![image](https://user-images.githubusercontent.com/3319450/108163705-0d550800-70b5-11eb-89fe-ab7d33d7ee84.png)
