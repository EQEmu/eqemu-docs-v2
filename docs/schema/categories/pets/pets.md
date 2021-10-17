# pets

| Column | Data Type | Description |
| :--- | :--- | :--- |
| type | varchar | [NPC Type Name](../../../schema/categories/pets/npc_types.md) |
| petpower | int | Pet Power |
| npcID | int | [NPC Type Identifier](../../../schema/categories/pets/npc_types.md) |
| temp | tinyint | Temporary: 0 = False, 1 = True |
| petcontrol | tinyint | Pet Control: 0 = No Control, 1 = No Attack Control, 2 = Full Control |
| petnaming | tinyint | Pet Naming: 0 = Soandsos Pet, 1 = Soandsos Familiar, 2 = Soandsos Warder, 3 = Random Naming \(i.e. Gobaner\), 4 = Keeps name from npc\_types table |
| monsterflag | tinyint | Monster Flag: 0 = False, 1 = True |
| equipmentset | int | [Pet Equipment Set Identifier](../../../schema/categories/pets/pet_equipmentset.md) |

