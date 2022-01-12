# Creating Animated Textures

Let's take this blender scenario:

Material: test

color image texture: test_c.dds

When you set color image inside blender, you're setting the e_Diffuse0 property (you can view this inside eqgzi-gui). 

If you create inside the same directory your zone's .blend file is at, e.g. test_c.txt, you can create an animation sequence:

An example file is this:

test_c.txt
```
3
500
test_c.dds
test02_c.dds
test03_c.dds
```


- The top row is the number of textures listed (in this case, 3)
- 500 is the time in milliseconds to show each texture.
- The remaining lines are parsed, and should match the first line (# of entries), if it is off, your client will crash

EQGZI during conversion will parse all files listed in this and inject them into the eqg, as well as the text file

That's it, test in game and see if the material is animating.