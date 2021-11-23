# spawnentry

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBucGNJRFxuICAgICAgICBpbnQgc3Bhd25ncm91cElEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgYWx0X2N1cnJlbmN5X2lkXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNoYW50X2lkXG4gICAgICAgIGludCBucGNfZmFjdGlvbl9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBucGNfc3BlbGxzX2lkXG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBlbW90ZWlkXG4gICAgICAgIGludHVuc2lnbmVkIGFybW9ydGludF9pZFxuICAgIH1cbiAgICBzcGF3bmVudHJ5IHx8LS1veyBzcGF3bmdyb3VwIDogT25lLXRvLU9uZVxuICAgIHNwYXduZW50cnkgfHwtLW97IG5wY190eXBlcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBucGNJRFxuICAgICAgICBpbnQgc3Bhd25ncm91cElEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgYWx0X2N1cnJlbmN5X2lkXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNoYW50X2lkXG4gICAgICAgIGludCBucGNfZmFjdGlvbl9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBucGNfc3BlbGxzX2lkXG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBlbW90ZWlkXG4gICAgICAgIGludHVuc2lnbmVkIGFybW9ydGludF9pZFxuICAgIH1cbiAgICBzcGF3bmVudHJ5IHx8LS1veyBzcGF3bmdyb3VwIDogT25lLXRvLU9uZVxuICAgIHNwYXduZW50cnkgfHwtLW97IG5wY190eXBlcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25lbnRyeSB7XG4gICAgICAgIGludCBucGNJRFxuICAgICAgICBpbnQgc3Bhd25ncm91cElEXG4gICAgfVxuICAgIHNwYXduZ3JvdXAge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbnBjX3R5cGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgYWx0X2N1cnJlbmN5X2lkXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNoYW50X2lkXG4gICAgICAgIGludCBucGNfZmFjdGlvbl9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBucGNfc3BlbGxzX2lkXG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBlbW90ZWlkXG4gICAgICAgIGludHVuc2lnbmVkIGFybW9ydGludF9pZFxuICAgIH1cbiAgICBzcGF3bmVudHJ5IHx8LS1veyBzcGF3bmdyb3VwIDogT25lLXRvLU9uZVxuICAgIHNwYXduZW50cnkgfHwtLW97IG5wY190eXBlcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

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

