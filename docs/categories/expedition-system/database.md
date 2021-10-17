# Database

The expedition system uses the following database tables

| table name | description |
| :--- | :--- |
| expeditions | Current active expeditions |
| expedition\_lockouts | Internal lockouts for current active expeditions |
| expedition\_members | Characters in active expeditions |
| character\_expedition\_lockouts | Character lockouts |
| dynamic\_zones | Extends instances with additional dz data \(compass, safereturn, zone-in\) |

### expeditions

| column | type | description |
| :--- | :--- | :--- |
| id | unsigned int | Expedition ID \(auto incremented\) |
| uuid | varchar\(36\) | Expedition UUID generated on creation used for certain lockout conflict checks |
| dynamic\_zone\_id | unsigned int | dynamic\_zones ID used by the expedition |
| expedition\_name | varchar\(128\) | Name of expedition |
| leader\_id | unsigned int | Current leader's character ID |
| min\_players | tinyint | Minimum players required to request expedition \(1-255\) |
| max\_players | tinyint | Maximum players allowed in expedition \(1-255\) |
| add\_replay\_on\_join | tinyint | \(bool\) If added members automatically receive `Replay Timer` lockout when joining an expedition that has it \(default: 1\) |
| is\_locked | tinyint | \(bool\) Allow adding new members \(0: unlocked, 1: locked\) |

### expedition\_lockouts

| column | type | description |
| :--- | :--- | :--- |
| id | unsigned int | Auto incremented ID |
| expedition\_id | unsigned int | Expedition ID the lockout belongs to |
| event\_name | varchar\(256\) | Lockout event name |
| expire\_time | datetime | Time the lockout expires |
| duration | unsigned int | Original duration \(seconds\) the lockout was added with |
| from\_expedition\_uuid | varchar\(36\) | Original source expedition uuid the lockout was assigned in \(will differ from current expedition if inherited by leader on creation\) |

### expedition\_members

| column | type | description |
| :--- | :--- | :--- |
| id | unsigned int | Auto incremented ID |
| expedition\_id | unsigned int | Expedition ID the member belongs to |
| character\_id | unsigned int | Member character ID |
| is\_current\_member | tinyint | \(bool\) Character is current member or past member of expedition |

### character\_expedition\_lockouts

| column | type | description |
| :--- | :--- | :--- |
| id | unsigned int | Auto incremented ID |
| character\_id | unsigned int | Character ID lockout belongs to |
| expedition\_name | varchar\(128\) | Lockout expedition name |
| event\_name | varchar\(256\) | Lockout event name |
| expire\_time | datetime | Time the lockout expires |
| duration | unsigned int | Original duration \(seconds\) the lockout was added with |
| from\_expedition\_uuid | varchar\(36\) | Original source expedition uuid the lockout was assigned in |

### dynamic\_zones

| column | type | description |
| :--- | :--- | :--- |
| id | unsigned int | Dynamic Zone ID \(auto incremented\) |
| instance\_id | int | ID for zone instance in `instance_list` table |
| type | tinyint | 0: None 1: Expedition 2: Tutorial 3: Task 4: Mission \(Shared Task\) 5: Quest |
| compass\_zone\_id | unsigned int | Zone ID the compass should be drawn in \(0: no compass\) |
| compass\_zone\_x | float | Compass x coordinate |
| compass\_zone\_y | float | Compass y coordinate |
| compass\_zone\_z | float | Compass z coordinate |
| safe\_return\_zone\_id | unsigned int | Zone ID characters are moved to when removed from a dz \(0: no safe return\) |
| safe\_return\_x | float | Safe return x coordinate |
| safe\_return\_y | float | Safe return y coordinate |
| safe\_return\_z | float | Safe return z coordinate |
| safe\_return\_heading | float | Safe return heading |
| zone\_in\_x | float | Zone in x coordinate when character moved via MovePCDynamicZone\(\) |
| zone\_in\_y | float | Zone in y coordinate when character moved via MovePCDynamicZone\(\) |
| zone\_in\_z | float | Zone in z coordinate when character moved via MovePCDynamicZone\(\) |
| zone\_in\_heading | float | Zone in heading when character moved via MovePCDynamicZone\(\) |
| has\_zone\_in | tinyint | \(bool\) If zone in override coordinates are valid \(zone's default zone-in location is used if not\) |

