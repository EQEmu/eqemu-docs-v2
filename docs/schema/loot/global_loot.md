# global_loot

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ2xvYmFsX2xvb3Qge1xuICAgICAgICBpbnQgbG9vdHRhYmxlX2lkXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICB9XG4gICAgY29udGVudF9mbGFncyB7XG4gICAgICAgIHZhcmNoYXIgZmxhZ19uYW1lXG4gICAgfVxuICAgIGxvb3R0YWJsZSB7XG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgZ2xvYmFsX2xvb3QgfHwtLW97IGNvbnRlbnRfZmxhZ3MgOiBPbmUtdG8tT25lXG4gICAgZ2xvYmFsX2xvb3QgfHwtLW97IGNvbnRlbnRfZmxhZ3MgOiBPbmUtdG8tT25lXG4gICAgZ2xvYmFsX2xvb3QgfHwtLW97IGxvb3R0YWJsZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ2xvYmFsX2xvb3Qge1xuICAgICAgICBpbnQgbG9vdHRhYmxlX2lkXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICB9XG4gICAgY29udGVudF9mbGFncyB7XG4gICAgICAgIHZhcmNoYXIgZmxhZ19uYW1lXG4gICAgfVxuICAgIGxvb3R0YWJsZSB7XG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgZ2xvYmFsX2xvb3QgfHwtLW97IGNvbnRlbnRfZmxhZ3MgOiBPbmUtdG8tT25lXG4gICAgZ2xvYmFsX2xvb3QgfHwtLW97IGNvbnRlbnRfZmxhZ3MgOiBPbmUtdG8tT25lXG4gICAgZ2xvYmFsX2xvb3QgfHwtLW97IGxvb3R0YWJsZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ2xvYmFsX2xvb3Qge1xuICAgICAgICBpbnQgbG9vdHRhYmxlX2lkXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICB9XG4gICAgY29udGVudF9mbGFncyB7XG4gICAgICAgIHZhcmNoYXIgZmxhZ19uYW1lXG4gICAgfVxuICAgIGxvb3R0YWJsZSB7XG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgZ2xvYmFsX2xvb3QgfHwtLW97IGNvbnRlbnRfZmxhZ3MgOiBPbmUtdG8tT25lXG4gICAgZ2xvYmFsX2xvb3QgfHwtLW97IGNvbnRlbnRfZmxhZ3MgOiBPbmUtdG8tT25lXG4gICAgZ2xvYmFsX2xvb3QgfHwtLW97IGxvb3R0YWJsZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | content_flags | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| One-to-One | content_flags_disabled | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| One-to-One | loottable_id | [loottable](../../schema/loot/loottable.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Global Loot Identifier |
| description | varchar | Description |
| loottable_id | int | [Loottable Identifier](loottable.md) |
| enabled | tinyint | Enabled: 0 = False, 1 = True |
| min_level | int | Minimum Level |
| max_level | int | Maximum Level |
| rare | tinyint | Rare: 0 = False, 1 = True |
| raid | tinyint | Raid: 0 = False, 1 = True |
| race | mediumtext | [Race](../../../../server/npc/race-list), multiple races supported if  |
| class | mediumtext | [Class](../../../../server/player/class-list), multiple classes supported if  |
| bodytype | mediumtext | [Body Type](../../../../server/npc/body-types), multiple body types supported if  |
| zone | mediumtext | [Zone Short Name](../../../../server/zones/zone-list),, multiple zones supported if  |
| hot_zone | tinyint | Hot Zone: 0 = False, 1 = True |
| min_expansion | tinyint | [Minimum Expansion](../../../../server/operation/expansion-list) |
| max_expansion | tinyint | [Maximum Expansion](../../../../server/operation/expansion-list) |
| content_flags | varchar | Content Flags Required to be Enabled |
| content_flags_disabled | varchar | Content Flags Required to be Disabled |

