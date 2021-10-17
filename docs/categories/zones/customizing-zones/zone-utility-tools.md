---
description: This page organizes various tools for manipulation of Zones.
---

# Zone Utility Tools

**Download**

* [https://ci.appveyor.com/api/projects/KimLS/zone-utilities/artifacts/build\_x64.zip](https://ci.appveyor.com/api/projects/KimLS/zone-utilities/artifacts/build_x64.zip)

**Source**

* [https://github.com/EQEmu/zone-utilities](https://github.com/EQEmu/zone-utilities)

**Contents**

```text
awater.exe
azone.exe
map_edit.exe
pfs.exe
```

## Azone

* Azone is a binary responsible for generating our .map files by reading geometry files, for more information on these see [Maps Introduction](https://eqemu.gitbook.io/server/categories/maps).
* Azone will look within the current directory for each zone\_name you pass it and attempt to open the files with one of three loaders in the following order:
  * EQG Standard
  * EQG Terrain \(v4\)
  * S3D

**Usage**

```text
./azone nektulos tutorialb qeytoqrg
```

This will load and output the following files

* nektulos.eqg -&gt; nektulos.map
* tutorialb.eqg -&gt; tutorialb.map
* qeytoqrg.s3d -&gt; qeytoqrg.map

## Awater

* Awater reads in a geometry file and outputs a water map file that can be loaded by the EQEmu server software and is then used for area detection purposes.
* Water maps are a bit of a misnomer as they handle more than water volume data but rather all marked area volumes

**Usage**

```text
./awater nektulos tutorialb qeytoqrg
```

This will load and output the following files

* nektulos.eqg -&gt; nektulos.wtr
* tutorialb.eqg -&gt; tutorialb.wtr
* qeytoqrg.s3d -&gt; qeytoqrg.wtr

Each of these **.wtr** files may then be copied to the server's maps directory to be used by the server.

## PFS

PFS is a command line utility for manipulating pfs \(S3D/EQG\) files. It works similarly to S3DSpy but from a command line.

The usage is modeled loosely after 7-Zip's command line utility for familiarity:

```markup
pfs [<switches>...] <command> <command args>... <archive_name> [<file_names>...]
<Switches>
 -i=dir: Set input directory
 -o=dir: Set output directory
<Commands>
 a: Add files to archive
 d: Delete files from the archive
 e: Extract files from the archive
 l: List contents of the archive
 <Command Args>
  arg1: Only search for files with this extension, may use * as a wildcard meaning all extensions
 u: Update files of the archive
```

## Map View

Map View is a utility that will attempt to load map and **.wtr** files and then render them to a 3D view

**Usage**

```text
./map_view nektulos
```

This will attempt to load **nektulos.map** and **nektulos.wtr**

#### Controls

You may use the mouse and WSAD for movement.

Holding shift will increase the speed you move within the world drastically.

* N toggles rendering of non-collidable geometry.
* C toggles rendering of collidable geometry.
* V toggles rendering of area volumes.

#### Troubleshooting

If you're having trouble with getting a zone to render here are some things to check:

* Your graphics card and drivers support OpenGL 3.3 or 3.0 \(you need to specify within CMake for a special 3.0 build\)
* The program can see the shaders/ directory and has the shaders files within it.
* The program can see the map and wtr files you are attempting to load.

