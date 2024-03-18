# Creating your first zone in Blender

If you haven't read the [Getting Started](getting-started.md), It is recommend to start there.

These instructions were written using Blender 2.93.5, but should apply to all versions of blender v2.80 or higher.

If you have Blender and aren't sure what version you are on, you can check that here: 
  
  ![How to find version](https://myimages.bravenet.com/264/906/210/2/PYY_1640762567.png)
  
  Otherwise you will want to visit [here](https://www.blender.org/download/) to get Blender.  It currently shows 3.0, which shouldn't cause much of a change to the process.  


  


## Making an extremely basic zone using Blender and Eqgzi

Now that all setup has hopefully gone well, you are probably hoping to be able to do something useful with it all.  To start, get blender going, you should end up with a scene something like this one:

![1](https://myimages.bravenet.com/264/906/210/2/TZZ_1640765696.png)

This is not a Blender tutorial however, it is more of a how to get blender to do what we want for eqgzi tutorial.  There are many tutorials out there, but for this example we are just going to make a simple box with an open top, the first shape you are shown how to make [here](https://www.youtube.com/watch?v=sW_NnFgIiso).   When you are done making the box, can use the scale tool to make it bigger, or press 'N' on your keyboard to adjust its dimensions directly.  For convenience in future tutorials, I'm going to make the X 500M, Y 500M, and the Z 100M like shown here:

![2](https://myimages.bravenet.com/264/906/210/2/KPP_1640766181.png)

So now we have a fantastical empty gray box, and it will work, but seizure warning if you do turn this into an eqg.  Next step we will need to add a shader/material to the box.  To start, we click on the 'Shader' tab near the top, then click 'New' near the middle (I circled it in red so it'd be easier to spot).

![3](https://myimages.bravenet.com/264/906/210/2/CUQ_1640766605.png)

After clicking new, a whole bunch of stuff will pop up, don't worry too much about this right now.  

![4](https://myimages.bravenet.com/264/906/210/2/GRN_1640766955.png)

For the next step you can pretty much add any dds file you want, but for those not feeling too creative about that, I have included [a dds for you here](https://drive.google.com/file/d/1avrTy3jb2dtEEH6D423yFuiGLHYJE4Z1/view?usp=sharing).  You will want to keep track of this file for later, it will need to be in the same location we save our .blend later.  For now though, just right click on the Add button (to the left of the New button), then under the texture tab, click on Image Texture.

![5](https://myimages.bravenet.com/264/906/210/2/HNK_1640767261.png)

Next use the open button to find wherever you decided to put the dds you will be using.  

![6](https://myimages.bravenet.com/264/906/210/2/ZFO_1640767721.png)

After that you will want to drag the circle next to color and connect it to the Base Color of the Principled BSDF.   While we are at that, **Material 001** is a pretty awful name,  lets rename that to **floor**.

![7](https://myimages.bravenet.com/264/906/210/2/TPL_1640767868.png)

Now we have a very basic box, with a very basic texture, that isn't scaled well at all.  If the scale really bothers you, you can adjust it using UV Editing, but not going to worry about that here.  As far as the basic blend project goes, we are now good to go, for my example we will be saving it as grobb.blend.  I don't think the location particularly matters, but in my case that location is ```C:\EQGZoneImporter\EQGZoneImporterv1_5\_grobb.eqg``` The end folder's name is ```_grobb.eqg```  

A few more files will be needed for the process to work, to save you the trouble of spelunking for them, I have provided them [here](https://drive.google.com/file/d/1Jyj_TpyDqzhkJqtqo-Tn84ptLGBieZrS/view?usp=sharing).  I included grobb.blend as well in case your blender adventure didn't end so well.  Currently, the folder should look something like this:

![8](https://myimages.bravenet.com/264/906/210/2/FXX_1640768958.png)

Some adjustment of the 'convert.bat' file will likely be necessary, unless you decided you loved my file structure so much you opted to emulate it.  For my convert.bat, it looks something like this:

```
set zone=grobb
blender --background %zone%.blend --python C:\EQGZoneImporter\EQGZoneImporterv1_5\convert.py
eqgzi import %zone%.obj %zone%.eqg
copy %zone%_emit.txt ..\%zone%_EnvironmentEmitters.txt
```

The part you will most likely need to adjust as you make new zones will be the first line where you set the zone, but for this first run you will want to ensure you have the right path for ```convert.py```.  

Assuming you edited it properly to point to the right locations, it should now be as simple as running convert.bat, then copying over grobb.eqg to a client you want to try it out in.  I recommend doing it somewhere you have GM permissions, as #goto is going to help a lot right here.  For me personally I like to add the other 2 lines at the end of convert.bat as well to save myself the effort:

```
copy C:\EQGZoneImporter\EQGZoneImporterv1_5\_grobb.eqg\grobb.eqg c:\TestServer\ /y
start "" "c:\Users\..\TEST eqgame.exe - Shortcut.lnk"
```

Assuming all went well, using #goto 0 0 0 should present a scene something like this.

![9](https://myimages.bravenet.com/264/906/210/2/YEE_1640769762.png)

Yes it is very dark, and yes the texture is scaled terribly right now.  We will talk about how to add some lights and a waterfall in the next section.  I have included the [output I got](https://drive.google.com/file/d/13vICQbnd4goipfwovmVChv7UIJOFlRbT/view?usp=sharing) as well if you needed to compare/contrast for whatever reason.

## Todo DZ
  - Add a section on how to add some lights and a waterfall graphic
  - Add a section on how to add a lava region.

## Initial setup:
- Create new pages for each 3d modeling tool, e.g. Maya, 3ds max, and basically the manual process until we write plugins to support each one.
- Start documenting object types in blender, custom properties supported, and how to do various feature injections (with the different 3d modeling pages, perhaps we can generically create this page for all of them to refer to)
