# Textures

| Texture ID | Texture Name |
| :--- | :--- |
| 0 | Naked for Player Races |
| 1 | Cloth for Player Races |
| 2 | Chain for Player Races |
| 3 | Plate for Player Races |
| 4 | Leather for Player Races |
| 5 | NPC Race Only |
| 6 | NPC Race Only |
| 7 | NPC Race Only |
| 8 | NPC Race Only |
| 9 | NPC Race Only |
| 10 | NPC Race Only |
| 11 | Robe 1 for Player Races |
| 12 | Robe 2 for Player Races |
| 13 | Robe 3 for Player Races |
| 14 | Robe 4 for Player Races |
| 15 | Robe 5 for Player Races |
| 16 | Robe 6 for Player Races |
| 17 | Velious Leather 1 |
| 18 | Velious Chain 1 |
| 19 | Velious Plate 1 |
| 20 | Velious Leather 2 |
| 21 | Velious Chain 2 |
| 22 | Velious Plate 2 |
| 23 | Velious Monk |
| 240 | Velious Helmets |

## Velious Textures

The following lines must exist in your client file "eqclient.ini" in order for Velious textures to load properly on the client:

```text
LoadArmor17=TRUE
LoadArmor18=TRUE
LoadArmor19=TRUE
LoadArmor20=TRUE
LoadArmor21=TRUE
LoadArmor22=TRUE
LoadArmor23=TRUE
LoadArmor240=TRUE
```

In order to use Velious helmet textures on an item the **idfile** field should be "IT240" and the **material** field should be 240. The client determines which helmet model to load based on the race and gender of the player that equips it.
