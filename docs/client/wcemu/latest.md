# WCEmu v1.5.7

## ACTORDEF

Wld actor definition

```c
// The callback function for the actor
  // Argument 1 (%s): The callback function
CALLBACK "1"
// The bounds reference for the actor
  // Argument 1 (%d): The bounds reference
BOUNDSREF 1
// The current action of the actor
  // Argument 1 (%s): The current action
CURRENTACTION? "1"
// The location of the actor
  // Argument 1 (%0.8e): The x coordinate
  // Argument 2 (%0.8e): The y coordinate
  // Argument 3 (%0.8e): The z coordinate
  // Argument 4 (%d): The x rotation
  // Argument 5 (%d): The y rotation
  // Argument 6 (%d): The z rotation
LOCATION? 1.12345678 2.12345678 3.12345678 4 5 6
// The active geometry of the actor
  // Argument 1 (%s): The active geometry
ACTIVEGEOMETRY? "1"
// The number of actions for the actor
  // Argument 1 (%d): The number of actions
NUMACTIONS 1
	// Action entry
	ACTION
	// Unknown entry 1
 	 // Argument 1 (%d): value of unk1
	UNK1 1
	// Number of levels of detail
 	 // Argument 1 (%d): Number of levels of detail
	NUMLEVELSOFDETAIL 1
		// Level of detail entry
		LEVELOFDETAIL
		// Sprite entry tag
 		 // Argument 1 (%s): sprite tag
		SPRITE "1"
		// Sprite index
 		 // Argument 1 (%d): Sprite index
		SPRITEINDEX 1
		// Minimum distance to render LOD
 		 // Argument 1 (%0.8e): Minimum distance
		MINDISTANCE 1.12345678
// Ignored in RoF2. 0x80 flag. This gets ignored if ActorInst doesn't have it. Likely need to use hierarchysprite flag for things like boats
  // Argument 1 (%d): 0: no, 1: yes
USEMODELCOLLIDER 1
// Unknown property 2
  // Argument 1 (%d): Unknown property 2
USERDATA 1
```



## ACTORINST

Wld actor instance

```c
// Reference to the actor's sprite tag
  // Argument 1 (%s): Actor's sprite reference tag
DEFINITION "1"
// The current action of the actor
  // Argument 1 (%s): The current action
CURRENTACTION? "1"
// The location of the actor
  // Argument 1 (%0.8e): The x coordinate
  // Argument 2 (%0.8e): The y coordinate
  // Argument 3 (%0.8e): The z coordinate
  // Argument 4 (%d): The x rotation
  // Argument 5 (%d): The y rotation
  // Argument 6 (%d): The z rotation
LOCATION? 1.12345678 2.12345678 3.12345678 4 5 6
// Radius around the actor instance for bounds
  // Argument 1 (%0.8e): Radius
BOUNDINGRADIUS? 1.12345678
// Scale factor of the actor instance
  // Argument 1 (%0.8e): Scale factor amount
SCALEFACTOR? 1.12345678
// Has a sound tag attached?
  // Argument 1 (%s): NULL if empty, sound tag
SOUND? "1"
// Is actor instance active?
  // Argument 1 (%d): NULL if empty, 1 if set to true
ACTIVE? 1
// Uses sprite volume?
  // Argument 1 (%d): NULL empty
SPRITEVOLUMEONLY? 1
// References an RGB Track?
  // Argument 1 (%s): NULL if not set, tag otherwise
DMRGBTRACK? "1"
// Reference to sphere tag
  // Argument 1 (%s): tag of sphere
SPHERE "1"
// Radius of sphere
  // Argument 1 (%0.8e): radius of sphere
SPHERERADIUS 1.12345678
// Use a bounding box
  // Argument 1 (%d): 0: false, 1: true
USEBOUNDINGBOX 1
// Unknown property 2
  // Argument 1 (%d): Unknown property 2
USERDATA 1
```



## AMBIENTLIGHT

Wld Ambient Light

```c
LIGHT 1.12345678
REGIONLIST 1 1
```



## BLITSPRITEDEF

Wld Blit Sprite

```c
// Sprite tag
  // Argument 1 (%s): Name of tag
SPRITE "1"
// Method for rendering
  // Argument 1 (%s): Rendering method
RENDERMETHOD "1"
// Is Transparent
  // Argument 1 (%d): 0: false, 1: true
TRANSPARENT 1
```



## DMSPRITEDEF

Wld DM sprite definition

```c
// The index of the tag
  // Argument 1 (%d): The index of the tag
TAGINDEX 1
// Fragment 1
  // Argument 1 (%d): fragment index
FRAGMENT1 1
// Material palette tag
  // Argument 1 (%s): Tag
MATERIALPALETTE "1"
// Fragment 3
  // Argument 1 (%d): fragment 3
FRAGMENT3 1
// center?
  // Argument 1 (%0.8e): The x coordinate
  // Argument 2 (%0.8e): The y coordinate
  // Argument 3 (%0.8e): The z coordinate
CENTER? 1.12345678 2.12345678 3.12345678
// params1
  // Argument 1 (%d): params1
  // Argument 2 (%d): params1
  // Argument 3 (%d): params1
PARAMS1? 1 2 3
// The number of vertices in the sprite
  // Argument 1 (%d): The number of vertices
NUMVERTICES 1
	// The coordinates of a vertex
 	 // Argument 1 (%0.8e): The x coordinate
 	 // Argument 2 (%0.8e): The y coordinate
 	 // Argument 3 (%0.8e): The z coordinate
	XYZ 1.12345678 2.12345678 3.12345678
// The number of texture coords
  // Argument 1 (%d): The number of tex coords
NUMTEXCOORDS 1
	// The coordinates of a texture normal
 	 // Argument 1 (%0.8e): The u coordinate
 	 // Argument 2 (%0.8e): The v coordinate
	UV 1.12345678 2.12345678
// The number of vertex normals in the sprite
  // Argument 1 (%d): The number of vertex normals
NUMNORMALS 1
	// The coordinates of a texture normal
 	 // Argument 1 (%0.8e): The x coordinate
 	 // Argument 2 (%0.8e): The y coordinate
 	 // Argument 3 (%0.8e): The z coordinate
	XYZ 1.12345678 2.12345678 3.12345678
// The number of vertex colors in the sprite
  // Argument 1 (%d): The number of vertex colors
NUMCOLORS 1
	// The coordinates of a vertex
 	 // Argument 1 (%d): Red
 	 // Argument 2 (%d): Green
 	 // Argument 3 (%d): Blue
 	 // Argument 4 (%d): Alpha
	RGBA 1 2 3 4
// The number of face2s in the sprite
  // Argument 1 (%d): The number of face2s
NUMFACES 1
	// DM Face 2 Entries
	DMFACE
	// face flags
 	 // Argument 1 (%d): face flags
	FLAG 1
	// face data
 	 // Argument 1 (%d): Index 0 of face data
 	 // Argument 2 (%d): Index 1 of face data
 	 // Argument 3 (%d): Index 2 of face data
 	 // Argument 4 (%d): Index 3 of face data
	DATA 1 2 3 4
	// Triangle indexes
 	 // Argument 1 (%d): Index 0 of triangle
 	 // Argument 2 (%d): Index 1 of triangle
 	 // Argument 3 (%d): Index 2 of triangle
	TRIANGLE 1 2 3
// The number of mesh operations in the sprite
  // Argument 1 (%d): The number of mesh operations
NUMMESHOPS 1
	// A mesh operation
 	 // Argument 1 (%d): arg 0
 	 // Argument 2 (%d): arg 1
 	 // Argument 3 (%0.8e): arg 2
 	 // Argument 4 (%d): arg 3
 	 // Argument 5 (%d): arg 4
	MESHOP 1 2 3.12345678 4 5
// The skin assignment groups
  // Argument 1 (%d): The size of the group
  // Argument 2 (%s): The data of the group
SKINASSIGNMENTGROUPS 1 2 3
// data 8 information
  // Argument 1 (%d): 8 info
DATA8 1
// The face material groups
  // Argument 1 (%d): The size of the group
  // Argument 2 (%d): The data of the group
  // Argument 3 (%d): The data of the group
FACEMATERIALGROUPS 1 2 3
// The vertex material groups
  // Argument 1 (%d): The size of the group
  // Argument 2 (%d): The data of the group
  // Argument 3 (%d): The data of the group
VERTEXMATERIALGROUPS 1 2 3
// params2
  // Argument 1 (%d): params2
  // Argument 2 (%d): params2
  // Argument 3 (%d): params2
PARAMS2? 1 2 3
```



## DMSPRITEDEF2

Wld DM sprite definition

```c
// The index of the tag
  // Argument 1 (%d): The index of the tag
TAGINDEX 1
// The center offset of the sprite
  // Argument 1 (%0.8e): The x offset
  // Argument 2 (%0.8e): The y offset
  // Argument 3 (%0.8e): The z offset
CENTEROFFSET 1.12345678 2.12345678 3.12345678
// The number of vertices in the sprite
  // Argument 1 (%d): The number of vertices
NUMVERTICES 1
	// The coordinates of a vertex
 	 // Argument 1 (%0.8e): The x coordinate
 	 // Argument 2 (%0.8e): The y coordinate
 	 // Argument 3 (%0.8e): The z coordinate
	XYZ 1.12345678 2.12345678 3.12345678
// The number of UVs in the sprite
  // Argument 1 (%d): The number of UVs
NUMUVS 1
	// UV entry
 	 // Argument 1 (%0.8e): U on UV map
 	 // Argument 2 (%0.8e): V on UV map
	UV 1.12345678 2.12345678
// The number of vertex normals in the sprite
  // Argument 1 (%d): The number of vertex normals
NUMVERTEXNORMALS 1
	// The coordinates of a texture normal
 	 // Argument 1 (%0.8e): The x coordinate
 	 // Argument 2 (%0.8e): The y coordinate
 	 // Argument 3 (%0.8e): The z coordinate
	XYZ 1.12345678 2.12345678 3.12345678
// The number of vertex colors in the sprite
  // Argument 1 (%d): The number of vertex colors
NUMVERTEXCOLORS 1
	// The coordinates of a vertex
 	 // Argument 1 (%d): Red
 	 // Argument 2 (%d): Green
 	 // Argument 3 (%d): Blue
 	 // Argument 4 (%d): Alpha
	RGBA 1 2 3 4
// The skin assignment groups
  // Argument 1 (%d): The size of the group
  // Argument 2 (%s): The data of the group
SKINASSIGNMENTGROUPS 1 2 3
// The material palette used by the sprite
  // Argument 1 (%s): The name of the material palette
MATERIALPALETTE "1"
// The DM track instance
  // Argument 1 (%s): The track instance
DMTRACKINST "1"
// The polyhedron definition
  // Argument 1 (%s): The definition of the polyhedron
POLYHEDRON "1"
// The definition reference
  // Argument 1 (%s): The definition
DEFINITION "1"
// The number of face2s in the sprite
  // Argument 1 (%d): The number of face2s
NUMFACE2S 1
	// DM Face 2 Entries
	DMFACE2
	// Is face passable?
 	 // Argument 1 (%d): Is face passable?
	PASSABLE 1
	// Triangle indexes
 	 // Argument 1 (%d): Index 0 of triangle
 	 // Argument 2 (%d): Index 1 of triangle
 	 // Argument 3 (%d): Index 2 of triangle
	TRIANGLE 1 2 3
// The number of mesh operations in the sprite
  // Argument 1 (%d): The number of mesh operations
NUMMESHOPS 1
	// A mesh operation
 	 // Argument 1 (%d): arg 0
 	 // Argument 2 (%d): arg 1
 	 // Argument 3 (%0.8e): arg 2
 	 // Argument 4 (%d): arg 3
 	 // Argument 5 (%d): arg 4
	MESHOP 1 2 3.12345678 4 5
// The face material groups
  // Argument 1 (%d): The size of the group
  // Argument 2 (%d): The data of the group
  // Argument 3 (%d): The data of the group
FACEMATERIALGROUPS 1 2 3
// The vertex material groups
  // Argument 1 (%d): The size of the group
  // Argument 2 (%d): The data of the group
  // Argument 3 (%d): The data of the group
VERTEXMATERIALGROUPS 1 2 3
// The minimum bounding box coordinates
  // Argument 1 (%0.8e): The x coordinate
  // Argument 2 (%0.8e): The y coordinate
  // Argument 3 (%0.8e): The z coordinate
BOUNDINGBOXMIN 1.12345678 2.12345678 3.12345678
// The maximum bounding box coordinates
  // Argument 1 (%0.8e): The x coordinate
  // Argument 2 (%0.8e): The y coordinate
  // Argument 3 (%0.8e): The z coordinate
BOUNDINGBOXMAX 1.12345678 2.12345678 3.12345678
// The bounding radius of the sprite
  // Argument 1 (%0.8e): The bounding radius
BOUNDINGRADIUS 1.12345678
// The FPS scale of the sprite
  // Argument 1 (%d): The FPS scale
FPSCALE 1
// The hex one flag
  // Argument 1 (%d): The hex one flag
HEXONEFLAG 1
// The hex two flag
  // Argument 1 (%d): The hex two flag
HEXTWOFLAG 1
// The hex four thousand flag
  // Argument 1 (%d): The hex four thousand flag
HEXFOURTHOUSANDFLAG 1
// The hex eight thousand flag
  // Argument 1 (%d): The hex eight thousand flag
HEXEIGHTTHOUSANDFLAG 1
// The hex ten thousand flag
  // Argument 1 (%d): The hex ten thousand flag
HEXTENTHOUSANDFLAG 1
// The hex twenty thousand flag
  // Argument 1 (%d): The hex twenty thousand flag
HEXTWENTYTHOUSANDFLAG 1
```



## DMTRACKDEF2

Wld DM Track Def 2

```c
SLEEP 1
PARAM2 1
FPSCALE 1
SIZE6 1
NUMFRAMES 1
	NUMVERTICES 1
		XYZ 1.12345678 2.12345678 3.12345678
```



## EQGANIDEF

EQG Animation Definition

```c
VERSION 1
STRICT 1
NUMBONES 1
	BONE "1"
	NUMFRAMES 1
		FRAME
		MILLISECONDS 1
		TRANSLATION 1.12345678 2.12345678 3.12345678
		ROTATION 1.12345678 2.12345678 3.12345678 4.12345678
		SCALE 1.12345678 2.12345678 3.12345678
```



## EQGLAYERDEF

EQG Layer Definition

```c
VERSION 1
NUMLAYERS 1
	LAYER
	MATERIAL "1"
	DIFFUSE "1"
	NORMAL "1"
```



## EQGSKINNEDMODELDEF

EQG Model Definition

```c
VERSION 1
NUMMATERIALS 1
	MATERIAL "1"
	SHADERTAG "1"
	HEXONEFLAG 1
	NUMPROPERTIES 1
		PROPERTY "1" 2 "3"
	ANIMSLEEP 1
	NUMANIMTEXTURES 1
		TEXTURE "1"
	NUMBONES 1
		BONE "1"
		NEXT 1
		CHILDREN 1
		CHILDINDEX 1
		PIVOT 1.12345678 2.12345678 3.12345678
		QUATERNION 1.12345678 2.12345678 3.12345678 4.12345678
		SCALE 1.12345678 2.12345678 3.12345678
	NUMMODELS 1
		MODEL "1"
		MAINPIECE 1
		NUMVERTICES 1
			VERTEX
			XYZ 1.12345678 2.12345678 3.12345678
			UV 1.12345678 2.12345678
			UV2 1.12345678 2.12345678
			NORMAL 1.12345678 2.12345678 3.12345678
			TINT 1 2 3 4
		NUMFACES 1
			FACE
			TRIANGLE 1 2 3
			MATERIAL "1"
			PASSABLE 1
			TRANSPARENT 1
			COLLISIONREQUIRED 1
			CULLED 1
			DEGENERATE 1
		NUMWEIGHTS 1
			WEIGHT 1 2.12345678 3 4.12345678 5 6.12345678 7 8.12345678
```



## EQGMODELDEF

EQG Model Definition

```c
VERSION 1
NUMMATERIALS 1
	MATERIAL "1"
	SHADERTAG "1"
	HEXONEFLAG 1
	NUMPROPERTIES 1
		PROPERTY "1" 2 "3"
	ANIMSLEEP 1
	NUMANIMTEXTURES 1
		TEXTURE "1"
	NUMVERTICES 1
		VERTEX
		XYZ 1.12345678 2.12345678 3.12345678
		UV 1.12345678 2.12345678
		UV2 1.12345678 2.12345678
		NORMAL 1.12345678 2.12345678 3.12345678
		TINT 1 2 3 4
	NUMFACES 1
		FACE
		TRIANGLE 1 2 3
		MATERIAL "1"
		PASSABLE 1
		TRANSPARENT 1
		COLLISIONREQUIRED 1
		CULLED 1
		DEGENERATE 1
	NUMBONES 1
		BONE
		NAME "1"
		NEXT 1
		CHILDREN 1
		CHILDINDEX 1
		PIVOT 1.12345678 2.12345678 3.12345678
		QUATERNION 1.12345678 2.12345678 3.12345678 4.12345678
		SCALE 1.12345678 2.12345678 3.12345678
```



## EQGTERDEF

EQG Model Definition

```c
VERSION 1
NUMMATERIALS 1
	MATERIAL "1"
	SHADERTAG "1"
	HEXONEFLAG 1
	NUMPROPERTIES 1
		PROPERTY "1" 2 "3"
	ANIMSLEEP 1
	NUMANIMTEXTURES 1
		TEXTURE "1"
	NUMVERTICES 1
		VERTEX
		XYZ 1.12345678 2.12345678 3.12345678
		UV 1.12345678 2.12345678
		UV2 1.12345678 2.12345678
		NORMAL 1.12345678 2.12345678 3.12345678
		TINT 1 2 3 4
	NUMFACES 1
		FACE
		TRIANGLE 1 2 3
		MATERIAL "1"
		PASSABLE 1
		TRANSPARENT 1
		COLLISIONREQUIRED 1
		CULLED 1
		DEGENERATE 1
	NUMBONES 1
		BONE
		NAME "1"
		NEXT 1
		CHILDREN 1
		CHILDINDEX 1
		PIVOT 1.12345678 2.12345678 3.12345678
		QUATERNION 1.12345678 2.12345678 3.12345678 4.12345678
		SCALE 1.12345678 2.12345678 3.12345678
```



## GLOBALAMBIENTLIGHTDEF

Wld Global Ambient Light Def is used for setting the global ambient light on WLD files

```c
// Is this a new wld file?
  // Argument 1 (%d): Red
  // Argument 2 (%d): Green
  // Argument 3 (%d): Blue
  // Argument 4 (%d): Alpha
COLOR 1 2 3 4
```



## HIERARCHICALSPRITEDEF

Wld  Hierarchical Sprite Def

```c
NUMDAGS 1
	DAG
	TAG "1"
	SPRITE "1"
	SPRITEINDEX 1
	TRACK "1"
	TRACKINDEX "1"
	SUBDAGLIST 1 1
NUMATTACHEDSKINS 1
	ATTACHEDSKIN
	DMSPRITE "1"
	DMSPRITEINDEX 1
	LINKSKINUPDATESTODAGINDEX 1
POLYHEDRON
DEFINITION "1"
CENTEROFFSET? 1.12345678 2.12345678 3.12345678
BOUNDINGRADIUS? 1.12345678
// also known as HAVEATTACHEDSKINS
HEXTWOHUNDREDFLAG 1
// also known as DAGCOLLISONS
HEXTWENTYTHOUSANDFLAG 1
```



## LIGHTDEFINITION

Wld Light

```c
// Is there a current frame, and what's value
  // Argument 1 (%d): NULL if skipped
CURRENTFRAME? 1
// Number of frames in light
  // Argument 1 (%d): Count of frames
NUMFRAMES 1
	// value of light level frame
 	 // Argument 1 (%0.8e): light level
	LIGHTLEVELS 1.12345678
// Is a sleep value set?
  // Argument 1 (%d): NULL if skipped, sleep value in ms
SLEEP? 1
// Are frames skipped
  // Argument 1 (%d): number of frames to skip
SKIPFRAMES 1
// Number of colors
  // Argument 1 (%d): Count of colors
NUMCOLORS 1
	// Color value
 	 // Argument 1 (%0.8e): R Value of color
 	 // Argument 2 (%0.8e): G Value of color
 	 // Argument 3 (%0.8e): B Value of color
	COLOR 1.12345678 2.12345678 3.12345678
```



## MATERIALDEFINITION

Wld Material

```c
// For tag variations, starts at 0, increases by 1
  // Argument 1 (%d): Index of tag
TAGINDEX 1
// For variations
  // Argument 1 (%d): Variation of tag
VARIATION 1
// Method for rendering
  // Argument 1 (%s): Rendering method
RENDERMETHOD "1"
// RGB Colorizing
  // Argument 1 (%d): Red
  // Argument 2 (%d): Green
  // Argument 3 (%d): Blue
  // Argument 4 (%d): Alpha
RGBPEN 1 2 3 4
// Color brightness
  // Argument 1 (%0.8e): Brightness amount
BRIGHTNESS 1.12345678
// Scaled ambient amount
  // Argument 1 (%0.8e): Scaled ambient amount
SCALEDAMBIENT 1.12345678
// Simple sprite instance section
SIMPLESPRITEINST
// Simple sprite instance tag
  // Argument 1 (%s): Simple sprite instance tag
TAG "1"
// Hex fifty flag
  // Argument 1 (%d): Hex fifty flag
HEXFIFTYFLAG 1
// Pairs of flags?
  // Argument 1 (%d): Pairs 0
  // Argument 2 (%d): Pairs 1
PAIRS? 1 2
// Is material double sided?
  // Argument 1 (%d): 0: False, 1: True
DOUBLESIDED 1
```



## MATERIALPALETTE

Wld Material Palette

```c
// Number of materials in the palette
  // Argument 1 (%d): Number of materials
NUMMATERIALS 1
	// Material tag
 	 // Argument 1 (%s): Tag of material
	MATERIAL "1"
```



## PARTICLECLOUDDEF

Wld Particle Cloud

```c
TAGINDEX 1
BLITTAG "1"
PARTICLETYPE 1
MOVEMENT "1"
HIGHOPACITY 1
FOLLOWITEM 1
SIZE 1
GRAVITYMULTIPLIER 1.12345678
GRAVITY 1.12345678 2.12345678 3.12345678
DURATION 1
SPAWNRADIUS 1.12345678
SPAWNANGLE 1.12345678
LIFESPAN 1
SPAWNVELOCITYMULTIPLIER 1.12345678
SPAWNVELOCITY 1.12345678 2.12345678 3.12345678
SPAWNRATE 1
SPAWNSCALE 1.12345678
TINT 1 2 3 4
SPAWNBOXMIN? 1.12345678 2.12345678 3.12345678
SPAWNBOXMAX? 1.12345678 2.12345678 3.12345678
BOXMIN? 1.12345678 2.12345678 3.12345678
BOXMAX? 1.12345678 2.12345678 3.12345678
HEXEIGHTYFLAG 1
HEXONEHUNDREDFLAG 1
HEXFOURHUNDREDFLAG 1
HEXFOURTHOUSANDFLAG 1
HEXEIGHTTHOUSANDFLAG 1
HEXTENTHOUSANDFLAG 1
HEXTWENTYTHOUSANDFLAG 1
```



## POINTLIGHT

Wld Point Light

```c
LIGHT "1"
STATIC 1
STATICINFLUENCE "1"
HASREGIONS 1
XYZ 1.12345678 2.12345678 3.12345678
RADIUSOFINFLUENCE 1.12345678
```



## POLYHEDRONDEFINITION

Wld Polyhedron Definition

```c
BOUNDINGRADIUS 1.12345678
SCALEFACTOR 1.12345678
NUMVERTICES 1
	XYZ 1.12345678 2.12345678 3.12345678
NUMFACES 1
	VERTEXLIST 1 1
HEXONEFLAG 1
```



## REGION

Wld Region

```c
REVERBVOLUME 1.12345678
REVERBOFFSET 1
REGIONFOG 1
GOURAND2 1
ENCODEDVISIBILITY 1
VISLISTBYTES 1
NUMREGIONVERTEX 1
	XYZ 1.12345678 2.12345678 3.12345678
NUMRENDERVERTICES 1
	XYZ 1.12345678 2.12345678 3.12345678
NUMWALLS 1
	WALL
	NORMALABCD 1.12345678 2.12345678 3.12345678 4.12345678
	NUMVERTICES 1
		XYZ 1.12345678 2.12345678 3.12345678
NUMOBSTACLES 1
	OBSTACLE
	NORMALABCD 1.12345678 2.12345678 3.12345678 4.12345678
	NUMVERTICES 1
		XYZ 1.12345678 2.12345678 3.12345678
NUMCUTTINGOBSTACLES 1
	CUTTINGOBSTACLE
	NORMALABCD 1.12345678 2.12345678 3.12345678 4.12345678
	NUMVERTICES 1
		XYZ 1.12345678 2.12345678 3.12345678
VISTREE
NUMVISNODE 1
	VISNODE
	NORMALABCD 1.12345678 2.12345678 3.12345678 4.12345678
	VISLISTINDEX 1
	FRONTTREE 1
	BACKTREE 1
NUMVISIBLELIST 1
	VISLIST
	REGIONS 1 1
SPHERE 1.12345678 2.12345678 3.12345678 4.12345678
USERDATA "1"
SPRITE "1"
```



## RGBDEFORMATIONTRACKDEF

Wld RGB 

```c
DATA1 1
DATA2 1
SLEEP 1
DATA4 1
RGBDEFORMATIONFRAME
NUMRGBAS 1
	RGBA 1 2 3 4
```



## SIMPLESPRITEDEF

Wld Simple Sprite

```c
// Index of tag
  // Argument 1 (%d): Index of tag
TAGINDEX 1
// Variation of tag
  // Argument 1 (%d): Variation of tag
VARIATION 1
// Should frames be skipped?
  // Argument 1 (%d): 0: false, 1: true
SKIPFRAMES? 1
// Is animated?
  // Argument 1 (%d): 0: false, 1: true
ANIMATED? 1
// Is there a sleep duration (in milliseconds)
  // Argument 1 (%d): NULL for non-value
SLEEP? 1
// Current frame set?
  // Argument 1 (%d): NULL for non-value
CURRENTFRAME? 1
// Number of frames in simple sprite
  // Argument 1 (%d): Number of frames
NUMFRAMES 1
	// Frame tag
 	 // Argument 1 (%s): Frame tag
	FRAME "1"
	// Number of files
 	 // Argument 1 (%d): Count of files
	NUMFILES 1
		// Texture file name
 		 // Argument 1 (%s): tag of file
		FILE "1"
```



## SPRITE2DDEF

Wld Sprite 2d Def

```c
SCALE 1.12345678 2.12345678
SPHERELISTTAG "1"
DEPTHSCALE? 1.12345678
CENTEROFFSET? 1.12345678 2.12345678 3.12345678
BOUNDINGRADIUS? 1.12345678
CURRENTFRAMEREF? 1
SLEEP? 1
NUMPITCHES 1
	PITCH
	PITCHCAP 1
	TOPORBOTTOMVIEW 1
	NUMHEADINGS 1
		HEADING
		HEADINGCAP 1
		NUMFRAMES 1
			FRAME "1"
			NUMFILES 1
				FILE "1"
RENDERMETHOD "1"
RENDERINFO
PEN? "1"
BRIGHTNESS? 1.12345678
SCALEDAMBIENT? 1.12345678
SPRITE? "1"
UVORIGIN? 1.12345678 2.12345678 3.12345678
UAXIS? 1.12345678 2.12345678 3.12345678
VAXIS? 1.12345678 2.12345678 3.12345678
UVCOUNT 1
	UV 1.12345678 2.12345678
TWOSIDED 1
HEXTENFLAG 1
```



## 3DSPRITEDEF

Wld 3d Sprite Definition

```c
CENTEROFFSET? 1.12345678 2.12345678 3.12345678
BOUNDINGRADIUS? 1.12345678
SPHERELIST "1"
NUMVERTICES 1
	XYZ 1.12345678 2.12345678 3.12345678
NUMBSPNODES 1
	BSPNODE
	VERTEXLIST 1 1
RENDERMETHOD "1"
RENDERINFO
PEN? 1
BRIGHTNESS? 1.12345678
SCALEDAMBIENT? 1.12345678
SPRITE? "1"
UVORIGIN? 1.12345678 2.12345678 3.12345678
UAXIS? 1.12345678 2.12345678 3.12345678
VAXIS? 1.12345678 2.12345678 3.12345678
UVCOUNT 1
	UV 1.12345678 2.12345678
TWOSIDED 1
FRONTTREE 1
BACKTREE 1
```



## TRACKDEFINITION

Wld Track

```c
TAGINDEX 1
NUMFRAMES 1
	FRAME 1 2 3 4 5 6 7 8
NUMLEGACYFRAMES 1
	LEGACYFRAME 1 2 3 4 5.12345678 6.12345678 7.12345678 8.12345678
```



## TRACKINSTANCE

Wld Track

```c
TAGINDEX 1
DEFINITION "1"
DEFINITIONINDEX 1
// deprecated, ignored in RoF2
  // Argument 1 (%d): deprecated, ignored in RoF2
INTERPOLATE 1
// deprecated, ignored in RoF2
  // Argument 1 (%d): deprecated, ignored in RoF2
REVERSE 1
SLEEP? 1
```



## WORLDDEF

Wld World definition
This is a collection of properties that defines a world

```c
// Is this a new wld file?
  // Argument 1 (%d): 0: old wld versioning, 1: new wld versioning
NEWWORLD 1
// Should this wce be treated like a zone?
  // Argument 1 (%s): 0: no, 1: yes
ZONE "1"
// Used in eqg parsing for version rebuilding
  // Argument 1 (%d): The version of the eqg file
EQGVERSION? 1
```



## WORLDTREE

Wld World Tree

```c
NUMWORLDNODES 1
	WORLDNODE
	NORMALABCD 1.12345678 2.12345678 3.12345678 4.12345678
	WORLDREGIONTAG "1"
	FRONTTREE 1
	BACKTREE 1
```



## ZONE

Wld Zone

```c
REGIONLIST 1 1
USERDATA "1"
```



