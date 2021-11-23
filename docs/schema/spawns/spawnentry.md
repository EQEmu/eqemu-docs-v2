# spawnentry

## Relationship Diagram
[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBzcGF3bmdyb3VwSURcbiAgICAgICAgaW50IG5wY0lEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGFsdF9jdXJyZW5jeV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjaGFudF9pZFxuICAgICAgICBpbnQgbnBjX2ZhY3Rpb25faWRcbiAgICAgICAgaW50dW5zaWduZWQgbnBjX3NwZWxsc19pZFxuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgZW1vdGVpZFxuICAgICAgICBpbnR1bnNpZ25lZCBhcm1vcnRpbnRfaWRcbiAgICB9XG4gICAgc3Bhd25lbnRyeSB8fC0tb3sgc3Bhd25ncm91cCA6IE9uZS10by1PbmVcbiAgICBzcGF3bmVudHJ5IHx8LS1veyBucGNfdHlwZXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBzcGF3bmdyb3VwSURcbiAgICAgICAgaW50IG5wY0lEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGFsdF9jdXJyZW5jeV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjaGFudF9pZFxuICAgICAgICBpbnQgbnBjX2ZhY3Rpb25faWRcbiAgICAgICAgaW50dW5zaWduZWQgbnBjX3NwZWxsc19pZFxuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgZW1vdGVpZFxuICAgICAgICBpbnR1bnNpZ25lZCBhcm1vcnRpbnRfaWRcbiAgICB9XG4gICAgc3Bhd25lbnRyeSB8fC0tb3sgc3Bhd25ncm91cCA6IE9uZS10by1PbmVcbiAgICBzcGF3bmVudHJ5IHx8LS1veyBucGNfdHlwZXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | spawngroupID | spawngroup | id |
| One-to-One | npcID | npc_types | id |


## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| spawngroupID | int | [Unique Spawn Group Identifier](spawngroup.md) |
| npcID | int | [NPC Type Identifier](../../schema/npcs/npc_types.md) |
| chance | smallint | Chance: 0 = Never, 100 = Always |
| condition_value_filter | mediumint |  |

