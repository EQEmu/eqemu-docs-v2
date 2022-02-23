# zone_points

!!! info
	This page was last generated 2022.02.23

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Zone Point Identifier |
| zone | varchar | [Zone Short Name](../../../../server/zones/zone-list) |
| version | int | Version |
| number | smallint | Represents the iterator field sent in the struct ZonePoint_Entry, zone points for the current zone are sent when client zones in (during Client::Handle_Connect_OP_ReqClientSpawn in client_packet.cpp).  This number field must be unique and also could have a hardcoded equivalent in the client, eg. client is expecting a specific number value for a zone point or teleport/object pad, such as in Erudin (erudnext). |
| y | float | Y Coordinate |
| x | float | X Coordinate |
| z | float | Z Coordinate |
| heading | float | Heading Coordinate |
| target_y | float | Target Y Coordinate |
| target_x | float | Target X Coordinate |
| target_z | float | Target Z Coordinate |
| target_heading | float | Target Heading Coordinate |
| zoneinst | smallint | [Instance Identifier](../../../server/instances/instance_list) |
| target_zone_id | int | [Target Zone Identifier](../../../../server/zones/zone-list) |
| target_instance | int | Target Instance Identifier |
| buffer | float | Zone Point Buffer |
| client_version_mask | int | [Client Version Mask](../../../../server/player/client-version-bitmasks) |
| min_expansion | tinyint |  |
| max_expansion | tinyint |  |
| content_flags | varchar |  |
| content_flags_disabled | varchar |  |
| is_virtual | tinyint |  |
| height | int |  |
| width | int |  |

