# zone_points

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    zone_points {
        varchar content_flags
        varchar content_flags_disabled
        intunsigned target_instance
        varchar zone
        intunsigned target_zone_id
        int version
    }
    content_flags {
        varchar flag_name
    }
    instance_list {
        int id
        tinyintunsigned version
        intunsigned zone
    }
    zone {
        int zoneidnumber
        varchar short_name
        tinyintunsigned version
        varchar content_flags
        varchar content_flags_disabled
    }
    zone_points ||--o{ content_flags : "One-to-One"
    zone_points ||--o{ content_flags : "One-to-One"
    zone_points ||--o{ instance_list : "One-to-One"
    zone_points ||--o{ zone : "One-to-One"
    zone_points ||--o{ zone : "One-to-One"
    zone_points ||--o{ zone : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | content_flags | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| One-to-One | content_flags_disabled | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| One-to-One | target_instance | [instance_list](../../schema/instances/instance_list.md) | id |
| One-to-One | target_zone_id | [zone](../../schema/zone/zone.md) | zoneidnumber |
| One-to-One | version | [zone](../../schema/zone/zone.md) | version |
| One-to-One | zone | [zone](../../schema/zone/zone.md) | short_name |


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
| min_expansion | tinyint | [Minimum Expansion](../../../../server/operation/expansion-list) |
| max_expansion | tinyint | [Maximum Expansion](../../../../server/operation/expansion-list) |
| content_flags | varchar | Content Flags Required to be Enabled |
| content_flags_disabled | varchar | Content Flags Required to be Disabled |
| is_virtual | tinyint | Is Virtual: 0 = False, 1 = True |
| height | int | Height |
| width | int | Width |

