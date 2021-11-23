# task_activities

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdGFza19hY3Rpdml0aWVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgdGFza2lkXG4gICAgICAgIGludHVuc2lnbmVkIGRlbGl2ZXJ0b25wY1xuICAgICAgICBpbnR1bnNpZ25lZCBnb2FsaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGFsdF9jdXJyZW5jeV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjaGFudF9pZFxuICAgICAgICBpbnQgbnBjX2ZhY3Rpb25faWRcbiAgICAgICAgaW50dW5zaWduZWQgbnBjX3NwZWxsc19pZFxuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgZW1vdGVpZFxuICAgICAgICBpbnR1bnNpZ25lZCBhcm1vcnRpbnRfaWRcbiAgICB9XG4gICAgZ29hbGxpc3RzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbGlzdGlkXG4gICAgfVxuICAgIHRhc2tfYWN0aXZpdGllcyB8fC0tb3sgbnBjX3R5cGVzIDogT25lLXRvLU9uZVxuICAgIHRhc2tfYWN0aXZpdGllcyB8fC0tb3sgZ29hbGxpc3RzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdGFza19hY3Rpdml0aWVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgdGFza2lkXG4gICAgICAgIGludHVuc2lnbmVkIGRlbGl2ZXJ0b25wY1xuICAgICAgICBpbnR1bnNpZ25lZCBnb2FsaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGFsdF9jdXJyZW5jeV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjaGFudF9pZFxuICAgICAgICBpbnQgbnBjX2ZhY3Rpb25faWRcbiAgICAgICAgaW50dW5zaWduZWQgbnBjX3NwZWxsc19pZFxuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgZW1vdGVpZFxuICAgICAgICBpbnR1bnNpZ25lZCBhcm1vcnRpbnRfaWRcbiAgICB9XG4gICAgZ29hbGxpc3RzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbGlzdGlkXG4gICAgfVxuICAgIHRhc2tfYWN0aXZpdGllcyB8fC0tb3sgbnBjX3R5cGVzIDogT25lLXRvLU9uZVxuICAgIHRhc2tfYWN0aXZpdGllcyB8fC0tb3sgZ29hbGxpc3RzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | delivertonpc | [npc_types](../../schema/npcs/npc_types.md) | id |
| Has-Many | goalid | [goallists](../../schema/tasks/goallists.md) | listid |


## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| taskid | int | [Task Identifier](tasks.md) |
| activityid | int | Activity Identifier: Starts at 0 |
| step | int | Step: 0 = Always Available, &gt;0 = Must Complete Previous |
| activitytype | tinyint | [Activity Type](../../../../server/task-system-guide/task-activity-types) |
| target_name | varchar | Target Name |
| item_list | varchar | [Item Identifier List](../../schema/items/items.md) |
| skill_list | varchar | [Skill Identifier List](../../../../server/player/skills) |
| spell_list | varchar | [Spell Identifier List](../../schema/spells/spells_new.md) |
| description_override | varchar | Description Override |
| goalid | int | Goal Identifier or [Goal List Identifier](goallists.md) |
| goalmethod | int | Goal Method: 0 = Single Value, 1 = List |
| goalcount | int | Goal Count |
| delivertonpc | int | Deliver To NPC: 0 = No Delivery NPC, &gt;0 = [NPC Type Identifier](../../schema/npcs/npc_types.md) |
| zones | varchar | [Zones List](../../../../server/zones/zone-list) |
| optional | tinyint | Optional: 0 = False, 1 = True |

