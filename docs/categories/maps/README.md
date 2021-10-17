# Maps

Maps in EQEmu are used to do many things, we have different files that are responsible for different functions.  Note that these are not the maps utilized inside your client to display your location--these are the maps for the zones themselves.

## File Structure

* maps/base/zoneshortname.map
* maps/nav/zoneshortname.nav
* maps/water/zoneshortname.wtr
* maps/path/zoneshortname.path \(Deprecated\)

## Base Maps \(.map\)

* Base map files are responsible for a few core critical things
  * **Line of Sight \(LOS\)** eg: Can this spell be casted if there is a wall or object between me and the target?
  * **Calculating Best Z** The server is constantly doing Z calculations to keep NPC's on level plane, to prevent from dipping into the ground which was a much larger issue in the past
* Maps are generated via [Azone](https://github.com/EQEmu/zone-utilities)

## Water Maps \(.wtr\)

* Water maps are responsible for one if not obvious, determining whether or not a client is inside water. Server side we determine different combat logic and pathing logic when a mob and/or client is in the water
* Water maps takes a point \(x, y, z\) and determines what type of region said point is marked as eg water, lava, normal
* Water maps are generated via [Awater](https://github.com/EQEmu/zone-utilities)

## Navigation Mesh \(.nav\)

Navmesh is modern navigation mesh technology, we use it server side to determine shortest path to a target in NPC AI decision making processes, it's what the server uses to determine what NPC's can walk on and they will strictly adhere to this mesh when making pathing decisions. The end result is a very smooth path-making decisions

#### Example of this in game

{% embed url="https://www.youtube.com/watch?v=ujtqipXAP1E" %}

Another explanation of Navmesh on [Stack Overflow](https://gamedev.stackexchange.com/a/15395)

