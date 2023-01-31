# Getting Started with EQGZI

## What is EQGZI?

As of **01/31/2023**, EQGZI is a program that turns a 3D modeled project into an everquest EQG zone file. It uses the OBJ format as an intermediary export, and supports: Lights, Environment Emitters, Material/Shader definitions, Region, and Spawn placement data.

## Requirements

- [Blender 3.0 or greater](https://www.blender.org/download/). This is the tool for editing zones.
- [EQGZI-Manager](https://github.com/xackery/eqgzi-manager/releases/latest). This is a standalone exe that manages EQGZI.


## Installation

A video of the following steps
<iframe width="100%" height="600" src="https://www.youtube.com/embed/G16H-qWBZ20" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

- Download [EQGZI-Manager](https://github.com/xackery/eqgzi-manager/releases/latest) and move it to a place you wish to keep your customized zones.
- Double click eqgzi-manager.exe
- Open eqgzi-manager, and click `Download EQGZI & Manager`
- Ensure the `Blender: path` is set. If not, press Detect after [installing Blender 3.0 or greater](https://www.blender.org/download/).
- Click `Create New Zone`
    - Type in the shortname of any zone, e.g. `clz`
- Click `Create clz.eqg`. See if the status says "Created clz.eqg" or not after
    - Click `Open clz folder` and see if an out folder is inside it. Open it, and see if clz.eqg was successfully made
- Your environment is now set up!


## Further reading and configuration

- [Shader List](shader-list.md) - A list of shaders
- [Blender Custom Properties](blender-custom-properties.md) - A list of custom properties supported by EQGZI inside Blender