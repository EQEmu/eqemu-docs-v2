---
description: This guide covers renaming a playable race.
---

# Renaming Playable Races

Whether you just want to rename one of the existing playable races, or you plan on doing a raceswap mod. You may want to rename one of your playable races. For the sake of this guide we will assume ROF2 client and that we're attempting to rename "Drakkin" to "Zombie".


## Update db_str Table

You will need to locate and update the appropriate fields in your db_str table. The `id` field will be the race id you want to change, type 8 is the race description seen on the character create screen, type 11 is race name, type 12 is plural race name. Using my example above I would do something like this:

`UPDATE db_str SET value='Zombie' WHERE id=522 AND type IN (11, 12);`
`UPDATE db_str SET value='Insert Zombie race description here' WHERE id=522 AND type=8`


## Update eqstr_us.txt

In order to change the 3 letter race code that appears in item tooltips and such you will need to locate `eqstr_us.txt` in your ROF2 directory. Ctrl + f to find the 3 letter race code you wish to change. In my example it's `DRK` for Drakkin. Change this to whichever 3 letter code you want to use such as `ZOM` and save. Keep in mind this file will need to be distributed to your players with your other client files via patcher or other means.


## Update Character Create Icon(s)

If you wish to update the icons for male/female options for your race you can do so by editing the following files:

- `rof2/uifiles/default/MaleRace.tga`
- `rof2/uifiles/default/FemaleRace.tga`

I use gimp but you can use whichever photo editing software you like as long as it supports tga files. Locate the section of the image with your chosen race icon, make your changes, then save. This will also need to be distributed with client files.


## Update CharacterCreate.xml

Locate `rof2/uifiles/default/EQUI_CharacterCreate.xml` and open it in a text editor. As far as I can tell the ROF2 client expects the contents of this file to match what your `db_str` entries are for each of your races. Going back to step 1 of this guide if you changed from "Drakkin" to "Zombie" you will need to open this file and do a find/replace of "Drakkin" and replace it with "Zombie" then save. Once again distribute this modified file with your client files. If you fail to do this step you will see a disabled icon on the character create screen and be unable to create a character with your renamed race.


## Done!

If you did everything correctly you should see the changes reflected in the character create screen including race name, description, and icons. You will also see the updated name in /who results in game as well as your new 3 letter race code in item tooltips.