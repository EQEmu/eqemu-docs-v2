# spawnentry

!!! info
	This page was last generated 2022.02.23

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBucGNJRFxuICAgICAgICBpbnQgc3Bhd25ncm91cElEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGFsdF9jdXJyZW5jeV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjaGFudF9pZFxuICAgICAgICBpbnQgbnBjX2ZhY3Rpb25faWRcbiAgICAgICAgaW50dW5zaWduZWQgbnBjX3NwZWxsc19pZFxuICAgICAgICBpbnR1bnNpZ25lZCBlbW90ZWlkXG4gICAgICAgIGludHVuc2lnbmVkIGFybW9ydGludF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICB9XG4gICAgc3Bhd25lbnRyeSB8fC0tb3sgc3Bhd25ncm91cCA6IE9uZS10by1PbmVcbiAgICBzcGF3bmVudHJ5IHx8LS1veyBucGNfdHlwZXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBucGNJRFxuICAgICAgICBpbnQgc3Bhd25ncm91cElEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGFsdF9jdXJyZW5jeV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjaGFudF9pZFxuICAgICAgICBpbnQgbnBjX2ZhY3Rpb25faWRcbiAgICAgICAgaW50dW5zaWduZWQgbnBjX3NwZWxsc19pZFxuICAgICAgICBpbnR1bnNpZ25lZCBlbW90ZWlkXG4gICAgICAgIGludHVuc2lnbmVkIGFybW9ydGludF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICB9XG4gICAgc3Bhd25lbnRyeSB8fC0tb3sgc3Bhd25ncm91cCA6IE9uZS10by1PbmVcbiAgICBzcGF3bmVudHJ5IHx8LS1veyBucGNfdHlwZXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBucGNJRFxuICAgICAgICBpbnQgc3Bhd25ncm91cElEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGFsdF9jdXJyZW5jeV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjaGFudF9pZFxuICAgICAgICBpbnQgbnBjX2ZhY3Rpb25faWRcbiAgICAgICAgaW50dW5zaWduZWQgbnBjX3NwZWxsc19pZFxuICAgICAgICBpbnR1bnNpZ25lZCBlbW90ZWlkXG4gICAgICAgIGludHVuc2lnbmVkIGFybW9ydGludF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICB9XG4gICAgc3Bhd25lbnRyeSB8fC0tb3sgc3Bhd25ncm91cCA6IE9uZS10by1PbmVcbiAgICBzcGF3bmVudHJ5IHx8LS1veyBucGNfdHlwZXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

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
| min_expansion | tinyint | [Expansion Identifier](https://docs.eqemu.io/developer/project-peq-expansions/expansion-list-reference/), Default: -1 (Enabled for All) |
| max_expansion | tinyint | [Expansion Identifier](https://docs.eqemu.io/developer/project-peq-expansions/expansion-list-reference/), Default: -1 (Enabled for All) |
| content_flags | varchar | See [Expansion and Content Filtering](https://docs.eqemu.io/developer/project-peq-expansions/expansion-content-filtering/), Default: NULL (Loaded in all scenarios)  |
| content_flags_disabled | varchar | See [Expansion and Content Filtering](https://docs.eqemu.io/developer/project-peq-expansions/expansion-content-filtering/), Default: NULL (Loaded in all scenarios)  |
