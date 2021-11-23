# spawnentry

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBzcGF3bmdyb3VwSURcbiAgICAgICAgaW50IG5wY0lEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIG5wY190eXBlcyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBhbHRfY3VycmVuY3lfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY2hhbnRfaWRcbiAgICAgICAgaW50IG5wY19mYWN0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG5wY19zcGVsbHNfaWRcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGVtb3RlaWRcbiAgICAgICAgaW50dW5zaWduZWQgYXJtb3J0aW50X2lkXG4gICAgfVxuICAgIHNwYXduZW50cnkgfHwtLW97IHNwYXduZ3JvdXAgOiBPbmUtdG8tT25lXG4gICAgc3Bhd25lbnRyeSB8fC0tb3sgbnBjX3R5cGVzIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBzcGF3bmdyb3VwSURcbiAgICAgICAgaW50IG5wY0lEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIG5wY190eXBlcyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBhbHRfY3VycmVuY3lfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY2hhbnRfaWRcbiAgICAgICAgaW50IG5wY19mYWN0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG5wY19zcGVsbHNfaWRcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGVtb3RlaWRcbiAgICAgICAgaW50dW5zaWduZWQgYXJtb3J0aW50X2lkXG4gICAgfVxuICAgIHNwYXduZW50cnkgfHwtLW97IHNwYXduZ3JvdXAgOiBPbmUtdG8tT25lXG4gICAgc3Bhd25lbnRyeSB8fC0tb3sgbnBjX3R5cGVzIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | spawngroupID | [spawngroup](../../schema/spawns/spawngroup.md) | id |
| One-to-One | npcID | [npc_types](../../schema/npcs/npc_types.md) | id |


## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| spawngroupID | int | [Unique Spawn Group Identifier](spawngroup.md) |
| npcID | int | [NPC Type Identifier](../../schema/npcs/npc_types.md) |
| chance | smallint | Chance: 0 = Never, 100 = Always |
| condition_value_filter | mediumint |  |

