# pets

!!! info
	This page was last generated 2022.05.11

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcGV0cyB7XG4gICAgICAgIGludCBucGNJRFxuICAgIH1cbiAgICBucGNfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBhbHRfY3VycmVuY3lfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY2hhbnRfaWRcbiAgICAgICAgaW50IG5wY19mYWN0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG5wY19zcGVsbHNfaWRcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGVtb3RlaWRcbiAgICAgICAgaW50dW5zaWduZWQgYXJtb3J0aW50X2lkXG4gICAgICAgIGludHVuc2lnbmVkIGxvb3R0YWJsZV9pZFxuICAgIH1cbiAgICBwZXRzIHx8LS1veyBucGNfdHlwZXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcGV0cyB7XG4gICAgICAgIGludCBucGNJRFxuICAgIH1cbiAgICBucGNfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBhbHRfY3VycmVuY3lfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY2hhbnRfaWRcbiAgICAgICAgaW50IG5wY19mYWN0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG5wY19zcGVsbHNfaWRcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGVtb3RlaWRcbiAgICAgICAgaW50dW5zaWduZWQgYXJtb3J0aW50X2lkXG4gICAgICAgIGludHVuc2lnbmVkIGxvb3R0YWJsZV9pZFxuICAgIH1cbiAgICBwZXRzIHx8LS1veyBucGNfdHlwZXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcGV0cyB7XG4gICAgICAgIGludCBucGNJRFxuICAgIH1cbiAgICBucGNfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBhbHRfY3VycmVuY3lfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY2hhbnRfaWRcbiAgICAgICAgaW50IG5wY19mYWN0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG5wY19zcGVsbHNfaWRcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludHVuc2lnbmVkIGVtb3RlaWRcbiAgICAgICAgaW50dW5zaWduZWQgYXJtb3J0aW50X2lkXG4gICAgICAgIGludHVuc2lnbmVkIGxvb3R0YWJsZV9pZFxuICAgIH1cbiAgICBwZXRzIHx8LS1veyBucGNfdHlwZXMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | npcID | [npc_types](../../schema/npcs/npc_types.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int |  |
| type | varchar | [NPC Type Name](../../schema/npcs/npc_types.md) |
| petpower | int | Pet Power |
| npcID | int | [NPC Type Identifier](../../schema/npcs/npc_types.md) |
| temp | tinyint | Temporary: 0 = False, 1 = True |
| petcontrol | tinyint | Pet Control: 0 = No Control, 1 = No Attack Control, 2 = Full Control |
| petnaming | tinyint | Pet Naming: 0 = Soandsos Pet, 1 = Soandsos Familiar, 2 = Soandsos Warder, 3 = Random Naming (i.e. Gobaner), 4 = Keeps name from npc_types table |
| monsterflag | tinyint | Monster Flag: 0 = False, 1 = True |
| equipmentset | int | [Pet Equipment Set Identifier](../../schema/pets/pets_equipmentset.md) |

