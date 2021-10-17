# Placing Objects

## Description

* So, have you ever wanted to know how to place a trade skill object? Well, it's really quite simple. You really only need to worry about the icon, model file, location, and heading.
* Doors in the database are stored in the[ object](http://wiki.eqemulator.org/p?object\&frm=Placing_Objects) table

## Requirements

* SQL tools from MySQL (Just get them, there very useful for everything.)
* Calculator (For heading)
* Basic SQL knowledge (Not necessary but useful)

```sql
INSERT INTO object VALUES (ItemID, ZoneID, LocYPos, LocXPos, Height, Heading, 0, 0, 'Model', Type, Icon, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0);
```

## **LocYPos, LocXPos, and Height**

* While in game, do the command /loc. Now I belive that loc is listed as x,y,z but objects use the cords from the real map. So x and y are backwards in /loc z is height (Not in real 3d, in real 3d x and z are the level plain and y is up and down.)

## **Icon**

* When you click a trade skill container it has an icon. This is the icon list.

| **Icon ID** | **Icon **     |
| ----------- | ------------- |
| 892         | Loom          |
| 1112        | Pottery Wheel |
| 1113        | Pottery Kiln  |
| 1114        | Stove         |
| 1115        | Forge         |
| 1116        | Brew barrel   |
| 1142        | Augment Bath  |

## **Model Files (objectname)**

* This is what the container will look like. Do I really need to clarify?

| **Model File ID** | **Model**     |
| ----------------- | ------------- |
| IT66\_ACTORDEF    | Forge         |
| IT69\_ACTORDEF    | Stove         |
| IT70\_ACTORDEF    | Brew barrel   |
| IT73\_ACTORDEF    | Pottery Kiln  |
| IT74\_ACTORDEF    | Pottery wheel |
| IT128\_ACTORDEF   | Loom          |
| IT10714\_ACTORDEF | Augment Bath  |

## **Heading**

* Use #loc, it displays heading ^-^

## **Type**

* This is the heart and soul of the clickable. This tells what it does.
* If the column is blank, it is non-applicable.
* See [Object Types Reference List](../object-types.md)
