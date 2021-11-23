# pets

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcGV0cyB7XG4gICAgICAgIGludCBucGNJRFxuICAgIH1cbiAgICBucGNfdHlwZXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgYWx0X2N1cnJlbmN5X2lkXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNoYW50X2lkXG4gICAgICAgIGludCBucGNfZmFjdGlvbl9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBucGNfc3BlbGxzX2lkXG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBlbW90ZWlkXG4gICAgICAgIGludHVuc2lnbmVkIGFybW9ydGludF9pZFxuICAgIH1cbiAgICBwZXRzIHx8LS1veyBucGNfdHlwZXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcGV0cyB7XG4gICAgICAgIGludCBucGNJRFxuICAgIH1cbiAgICBucGNfdHlwZXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50dW5zaWduZWQgYWx0X2N1cnJlbmN5X2lkXG4gICAgICAgIGludHVuc2lnbmVkIG1lcmNoYW50X2lkXG4gICAgICAgIGludCBucGNfZmFjdGlvbl9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBucGNfc3BlbGxzX2lkXG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBlbW90ZWlkXG4gICAgICAgIGludHVuc2lnbmVkIGFybW9ydGludF9pZFxuICAgIH1cbiAgICBwZXRzIHx8LS1veyBucGNfdHlwZXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | npcID | [npc_types](../../schema/npcs/npc_types.md) | id |


## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| type | varchar | [NPC Type Name](../../schema/npcs/npc_types.md) |
| petpower | int | Pet Power |
| npcID | int | [NPC Type Identifier](../../schema/npcs/npc_types.md) |
| temp | tinyint | Temporary: 0 = False, 1 = True |
| petcontrol | tinyint | Pet Control: 0 = No Control, 1 = No Attack Control, 2 = Full Control |
| petnaming | tinyint | Pet Naming: 0 = Soandsos Pet, 1 = Soandsos Familiar, 2 = Soandsos Warder, 3 = Random Naming \(i.e. Gobaner\), 4 = Keeps name from npc_types table |
| monsterflag | tinyint | Monster Flag: 0 = False, 1 = True |
| equipmentset | int | [Pet Equipment Set Identifier](../../schema/pets/pets_equipmentset.md) |

