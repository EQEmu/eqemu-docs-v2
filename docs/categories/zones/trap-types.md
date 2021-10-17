# Trap Types

| **Type** | **Effect** | **Effect Value** | **Effect Value 2** | **Notes** |
| :--- | :--- | :--- | :--- | :--- |
| Casting Trap | 0 | Spell ID | Unused | Casts a spell on the player that triggered the trap. |
| Alarm Trap | 1 | Range of effect | 0 \(all NPCs aggro\) or 1 \(only KOS NPCs aggro\) | Causes NPCs to aggro the player who triggered the trap. |
| Mystic Spawn \(NPC Trap\) | 2 | npc\_type ID of the mob spawned | number of mobs spawned | Spawns NPCs at a close spot in a small area |
| Bandit Spawn \(NPC Trap2\) | 3 | npc\_type ID of the mob spawned | number of mobs spawned | Spawns NPCs at a close spot in a large area |
| Damage Trap | 4 | minimum amount of damage | maximum amount of damage | Causes damage to the player that triggered the trap. |

### Trap Message Defaults

| **Type** | **Effect** | **Default Message** |
| :--- | :--- | :--- |
| Casting Trap | 0 | "\_\_\_\_\_ triggers a trap!" |
| Alarm Trap | 1 | "A loud alarm rings out through the air..." |
| Mystic Spawn | 2 | "The air shimmers..." |
| Bandit Spawn | 3 | "A bandit leaps out from behind a tree!" |
| Damage Trap | 4 | "\_\_\_\_\_ triggers a trap!" |

