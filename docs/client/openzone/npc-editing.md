1. Output s3d with Lantern Extractor as gltf.

1. Import gltf into Blender

1. Rename the animations in the format detailed in the Openzone documentation. 

1. Add some “jack” bones between root and pe (I use an IK chain and some contraints and then bake the animations)

1. Export as fbx with “Better fbx importer & exporter” plugin (cost me $28, but works. Built in fbx exporter didn’t)

1. Convert to an8 with “ObjtoAn8” program that I got here https://github.com/carnivoroussociety/GoldEditor

1. In anim8or, you need to rename the mesh from “obj” or whatever it is to the format detailed in the Openzone docs. 

1. Put in Openzone “creature” folder. Then add it to zone in Openzone and export as s3d. 

1. Replace s3d in client files with exported s3d. If your model is the wrong orientation, size, or height you can adjust it in Blender until it comes out right in game. 