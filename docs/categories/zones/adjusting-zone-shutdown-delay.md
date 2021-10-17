---
description: How to configure how long it takes for zones typically to shutdown
---

# Adjusting Zone Shutdown Delay

### Dynamic Zones

In a typical server installation; the server launcher boots what are called **dynamic zones**. These are zones that are sitting idle, not assigned to a zone but are fully ready to be assigned an actual zone from the **world server**. Once a client requests a zone request to the world server, it will check its local list of connected zones to see if the zone is already booted before making the decision to make a zone assignment

### Shutdown Delay

Typically most zones in a default PEQ database installation have a shutdown delay of 60 minutes, if you look in the **zone **table this value is configured in **milliseconds. **Below this is illustrated in minutes to make it easier to read

Shutdown delays are configured two ways, through the rule **Zone:AutoShutdownDelay** or the value in the zone table. Whichever one is higher is the one that the zone uses to enforce shutting back down

The shutdown delay timer only starts when the last client has left the zone resulting in 0 clients remaining in a zone and the timer resets when any client zones back in. When this timer expires the process is killed and your launcher will re-spawn another process to take its place

```
select
  short_name,
  ROUND(shutdowndelay / 1000 / 60) as shutdown_minutes
from
  zone
ORDER BY
  shutdowndelay desc
limit
  10;
  
+------------+------------------+
| short_name | shutdown_minutes |
+------------+------------------+
| kithicor   |             1440 |
| airplane   |              720 |
| tacvi      |              300 |
| arena2     |               60 |
| arena      |               60 |
| arcstone   |               60 |
| apprentice |               60 |
| anguish    |               60 |
| akheva     |               60 |
| arttest    |               60 |
+------------+------------------+
```
