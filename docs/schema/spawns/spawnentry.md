# spawnentry

!!! info
	This page was last generated 2022.06.17

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBzcGF3bmdyb3VwSURcbiAgICAgICAgaW50IG5wY0lEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGxvb3R0YWJsZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBhbHRfY3VycmVuY3lfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY2hhbnRfaWRcbiAgICAgICAgaW50IG5wY19mYWN0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG5wY19zcGVsbHNfaWRcbiAgICAgICAgaW50dW5zaWduZWQgZW1vdGVpZFxuICAgICAgICBpbnR1bnNpZ25lZCBhcm1vcnRpbnRfaWRcbiAgICB9XG4gICAgc3Bhd25lbnRyeSB8fC0tb3sgc3Bhd25ncm91cCA6IE9uZS10by1PbmVcbiAgICBzcGF3bmVudHJ5IHx8LS1veyBucGNfdHlwZXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBzcGF3bmdyb3VwSURcbiAgICAgICAgaW50IG5wY0lEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGxvb3R0YWJsZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBhbHRfY3VycmVuY3lfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY2hhbnRfaWRcbiAgICAgICAgaW50IG5wY19mYWN0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG5wY19zcGVsbHNfaWRcbiAgICAgICAgaW50dW5zaWduZWQgZW1vdGVpZFxuICAgICAgICBpbnR1bnNpZ25lZCBhcm1vcnRpbnRfaWRcbiAgICB9XG4gICAgc3Bhd25lbnRyeSB8fC0tb3sgc3Bhd25ncm91cCA6IE9uZS10by1PbmVcbiAgICBzcGF3bmVudHJ5IHx8LS1veyBucGNfdHlwZXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBzcGF3bmdyb3VwSURcbiAgICAgICAgaW50IG5wY0lEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGxvb3R0YWJsZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBhbHRfY3VycmVuY3lfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY2hhbnRfaWRcbiAgICAgICAgaW50IG5wY19mYWN0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG5wY19zcGVsbHNfaWRcbiAgICAgICAgaW50dW5zaWduZWQgZW1vdGVpZFxuICAgICAgICBpbnR1bnNpZ25lZCBhcm1vcnRpbnRfaWRcbiAgICB9XG4gICAgc3Bhd25lbnRyeSB8fC0tb3sgc3Bhd25ncm91cCA6IE9uZS10by1PbmVcbiAgICBzcGF3bmVudHJ5IHx8LS1veyBucGNfdHlwZXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

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
| condition_value_filter | mediumint | Condition Value Filter |
| min_expansion | tinyint | [Minimum Expansion](../../../../server/operation/expansion-list) |
| max_expansion | tinyint | [Maximum Expansion](../../../../server/operation/expansion-list) |
| content_flags | varchar | Content Flags Required to be Enabled |
| content_flags_disabled | varchar | Content Flags Required to be Disabled |

