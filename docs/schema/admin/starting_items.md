# starting_items

!!! info
	This page was last generated 2022.06.17

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3RhcnRpbmdfaXRlbXMge1xuICAgICAgICBpbnQgaXRlbWlkXG4gICAgICAgICB6b25lX2lkXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgICB6b25laWR1bnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgIH1cbiAgICBpdGVtcyB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBzdGFydGluZ19pdGVtcyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcbiAgICBzdGFydGluZ19pdGVtcyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3RhcnRpbmdfaXRlbXMge1xuICAgICAgICBpbnQgaXRlbWlkXG4gICAgICAgICB6b25lX2lkXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgICB6b25laWR1bnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgIH1cbiAgICBpdGVtcyB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBzdGFydGluZ19pdGVtcyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcbiAgICBzdGFydGluZ19pdGVtcyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3RhcnRpbmdfaXRlbXMge1xuICAgICAgICBpbnQgaXRlbWlkXG4gICAgICAgICB6b25lX2lkXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgICB6b25laWR1bnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgIH1cbiAgICBpdGVtcyB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBzdGFydGluZ19pdGVtcyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcbiAgICBzdGFydGluZ19pdGVtcyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | zone_id | [zone](../../schema/zone/zone.md) | zoneidnumber |
| One-to-One | itemid | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Starting Items Entry Identifier |
| race | int | [Race](../../../../server/npc/race-list): 0 = All |
| class | int | [Class](../../../../server/player/class-list): 0 = All |
| deityid | int | [Deity](../../../../server/player/deity-list): 0 = Alll |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| itemid | int | [Item Identifier](../../schema/items/items.md) |
| item_charges | tinyint | Item Charges |
| gm | tinyint | GM: 0 = False, 1 = True |
| slot | mediumint | [Slot](../../../../server/inventory/inventory-slots) |
| min_expansion | tinyint | [Minimum Expansion](../../../../server/operation/expansion-list) |
| max_expansion | tinyint | [Maximum Expansion](../../../../server/operation/expansion-list) |
| content_flags | varchar | Content Flags Required to be Enabled |
| content_flags_disabled | varchar | Content Flags Required to be Disabled |

