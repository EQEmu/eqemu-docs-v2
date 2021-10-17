# zone\_points

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Zone Point Identifier |
| zone | varchar | [Zone Short Name](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| version | int | Version |
| number | smallint | Represents the iterator field sent in the struct ZonePoint\_Entry, zone points for the current zone are sent when client zones in \(during Client::Handle\_Connect\_OP\_ReqClientSpawn in client\_packet.cpp\).  This number field must be unique and also could have a hardcoded equivalent in the client, eg. client is expecting a specific number value for a zone point or teleport/object pad, such as in Erudin \(erudnext\). |
| y | float | Y Coordinate |
| x | float | X Coordinate |
| z | float | Z Coordinate |
| heading | float | Heading Coordinate |
| target\_y | float | Target Y Coordinate |
| target\_x | float | Target X Coordinate |
| target\_z | float | Target Z Coordinate |
| target\_heading | float | Target Heading Coordinate |
| zoneinst | smallint | [Instance Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/zone/instance_list.d) |
| target\_zone\_id | int | [Target Zone Identifier](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| target\_instance | int | Target Instance Identifier |
| buffer | float | Zone Point Buffer |
| client\_version\_mask | int | [Client Version Mask](https://eqemu.gitbook.io/server/categories/player/client-version-bitmasks) |

