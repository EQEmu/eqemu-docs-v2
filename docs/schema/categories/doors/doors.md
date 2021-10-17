# doors

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Entry Identifier |
| doorid | smallint | Unique Door Identifier |
| zone | varchar | [Zone Short Name](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| version | smallint | Zone Version: -1 For All |
| name | varchar | This is the name of the door, such as 'IT11161' or 'POPCREATE501', for names of objects you can see. |
| pos\_y | float | Door Y Coordinate |
| pos\_x | float | Door X Coordinate |
| pos\_z | float | Door Z Coordinate |
| heading | float | Door Heading Coordinate |
| opentype | smallint | [Door Open Type](https://eqemu.gitbook.io/server/categories/zones/door-open-types) |
| guild | smallint | [Guild Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/doors/guilds.md) |
| lockpick | smallint | Lockpicking Skill Required: -1 = Unpickable |
| keyitem | int | [Item Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/doors/items.md) |
| nokeyring | tinyint | No Key Ring: 0 = False, 1 = True |
| triggerdoor | smallint | Trigger Door: 0 For Current Door or use a Unique Door Identifier |
| triggertype | smallint | Trigger Type: 1 = Open a Type 255 door, 255 = Will Not Open |
| disable\_timer | tinyint | Disable Timer |
| doorisopen | smallint | Door Is Open: 0 = False, 1 = True |
| door\_param | int | Door Parameter |
| dest\_zone | varchar | [Zone Short Name](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| dest\_instance | int | [Destination Instance](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/doors/instance_list.md) |
| dest\_x | float | Destination X Coordinate |
| dest\_y | float | Destination Y Coordinate |
| dest\_z | float | Destination Z Coordinate |
| dest\_heading | float | Destination Heading Coordinate |
| invert\_state | int | This column will basically behave like such: if the door has a click type and it is to raise up like a door, it will be raised on spawn of the door. Meaning it is inverted. Another example: If a [Door Open Type](https://eqemu.gitbook.io/server/categories/zones/door-open-types) is set to a spinning object on click, you could set this to 1 to have the door be spinning on spawn. |
| incline | int | Incline |
| size | smallint | Size |
| buffer | float | Unused |
| client\_version\_mask | int | [Client Version Mask](https://eqemu.gitbook.io/server/categories/player/client-version-bitmasks) |
| is\_ldon\_door | smallint | Is LDoN Door: 0 = False, 1 = True |

