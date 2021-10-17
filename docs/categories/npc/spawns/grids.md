# Grids

* **\#wpinfo** - Target an NPC and use this to retrieve WP information assigned to it.
* **\#gassign** grid\_num - Target an NPC and use this to assign it's spawn2 point to grid grid\_num.
* **\#grid** add/delete grid\_num wandertype pausetype - Adds grid grid\_num with the wandertype/pausetype, or deletes grid number grid\_num.
* **\#wp** add/delete grid\_num pause wp\_num -h - Adds waypoint number wp\_num to grid grid\_num with a pause of pause on the location you're standing at, or deletes waypoint number wp\_num in grid number grid\_num. Use the -h switch to use your current heading as the heading for that waypoint.
* **\#wpadd** \[pause\] \[-h\] - Adds a waypoint to a targeted NPC. If no grid is assigned, a new one will be created. Pause \(default 0\) and the heading switch are optional. As with \#wp add, the heading switch will use your current heading.

## Wander Type:

* 0: Circular. NPC will cycle waypoints in order, then head back to the first waypoint.
* 1: Random 10. NPC will pick a random of the nearest 10 waypoints and go to it.
* 2: Random. NPCwill pick a random waypoint in the grid and go to it.
* 3: Patrol. NPC will walk the waypoints to the end, then turn and backtrack.
* 4: One Way. NPC will walk the waypoints to the end, then repop.
* 5: Random 5 LoS. NPC will pick a random of the nearest 5 waypoints within line of sight.
* 6: One Way. NPC will walk the waypoints to the end, then Depop.
* 7: Center Point.  NPC will treat wp0 as center and move between it and a random wp.
* 8: Random Center Point.  NPC will alternate between a random waypoint and a random waypoint that is marked as center point.
* 9: Random Path.  Pick random waypoints, but follow path instead of heading there directly.

## Pause Type:

* 0: Random half. NPC pauses for half of it's pause time + random of half it's pause time.
* 1: Full. NPC pauses for full pause time.
* 2: Random. NPC pauses for random of it's pause time.

## Note:

* To setup an NPC to a grid, simply \#gassign an NPC spawned by a certain spawn2 point, and all NPC's spawned by that point will adhere to the selected grid. An NPC assigned to a grid will simply start wandering it after spawning.

