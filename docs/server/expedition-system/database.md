# Database

The expedition system uses the following database tables

| table name | description |
| :--- | :--- |
| expeditions | Current active expeditions |
| expedition_lockouts | Internal lockouts for current active expeditions |
| expedition_members | Characters in active expeditions |
| character_expedition_lockouts | Character lockouts |
| dynamic_zones | Extends instances with additional dz data (compass, safereturn, zone-in) |

### expeditions

| column | type | description |
| :--- | :--- | :--- |
| id | unsigned int | Expedition ID (auto incremented) |
| uuid | varchar(36) | Expedition UUID generated on creation used for certain lockout conflict checks |
| dynamic_zone_id | unsigned int | dynamic_zones ID used by the expedition |
| expedition_name | varchar(128) | Name of expedition |
| leader_id | unsigned int | Current leader's character ID |
| min_players | tinyint | Minimum players required to request expedition (1-255) |
| max_players | tinyint | Maximum players allowed in expedition (1-255) |
| add_replay_on_join | tinyint | (bool) If added members automatically receive `Replay Timer` lockout when joining an expedition that has it (default: 1) |
| is_locked | tinyint | (bool) Allow adding new members (0: unlocked, 1: locked) |

### expedition_lockouts

| column | type | description |
| :--- | :--- | :--- |
| id | unsigned int | Auto incremented ID |
| expedition_id | unsigned int | Expedition ID the lockout belongs to |
| event_name | varchar(256) | Lockout event name |
| expire_time | datetime | Time the lockout expires |
| duration | unsigned int | Original duration (seconds) the lockout was added with |
| from_expedition_uuid | varchar(36) | Original source expedition uuid the lockout was assigned in (will differ from current expedition if inherited by leader on creation) |

### expedition_members

| column | type | description |
| :--- | :--- | :--- |
| id | unsigned int | Auto incremented ID |
| expedition_id | unsigned int | Expedition ID the member belongs to |
| character_id | unsigned int | Member character ID |
| is_current_member | tinyint | (bool) Character is current member or past member of expedition |

### character_expedition_lockouts

| column | type | description |
| :--- | :--- | :--- |
| id | unsigned int | Auto incremented ID |
| character_id | unsigned int | Character ID lockout belongs to |
| expedition_name | varchar(128) | Lockout expedition name |
| event_name | varchar(256) | Lockout event name |
| expire_time | datetime | Time the lockout expires |
| duration | unsigned int | Original duration (seconds) the lockout was added with |
| from_expedition_uuid | varchar(36) | Original source expedition uuid the lockout was assigned in |

### dynamic_zones

| column | type | description |
| :--- | :--- | :--- |
| id | unsigned int | Dynamic Zone ID (auto incremented) |
| instance_id | int | ID for zone instance in `instance_list` table |
| type | tinyint | 0: None 1: Expedition 2: Tutorial 3: Task 4: Mission (Shared Task) 5: Quest |
| compass_zone_id | unsigned int | Zone ID the compass should be drawn in (0: no compass) |
| compass_zone_x | float | Compass x coordinate |
| compass_zone_y | float | Compass y coordinate |
| compass_zone_z | float | Compass z coordinate |
| safe_return_zone_id | unsigned int | Zone ID characters are moved to when removed from a dz (0: no safe return) |
| safe_return_x | float | Safe return x coordinate |
| safe_return_y | float | Safe return y coordinate |
| safe_return_z | float | Safe return z coordinate |
| safe_return_heading | float | Safe return heading |
| zone_in_x | float | Zone in x coordinate when character moved via MovePCDynamicZone() |
| zone_in_y | float | Zone in y coordinate when character moved via MovePCDynamicZone() |
| zone_in_z | float | Zone in z coordinate when character moved via MovePCDynamicZone() |
| zone_in_heading | float | Zone in heading when character moved via MovePCDynamicZone() |
| has_zone_in | tinyint | (bool) If zone in override coordinates are valid (zone's default zone-in location is used if not) |

