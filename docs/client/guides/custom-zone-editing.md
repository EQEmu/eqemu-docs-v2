---
description: This guide covers adding or editing zones in the client.
---

# Custom Zone Editing

When you make a new zone, or want to edit an existing one, you need to override some hard coded client side data points. There's two routes to do this: eq-core-dll injection, or eqgame patching

## EQ-Core-DLL Injection

This flow involves generating a dinput8.dll that goes in your EQ folder. You'll need to distribute this to your clients, one option for doing so is using [eqemupatcher](https://github.com/xackery/eqemupatcher).

Note that you can provide an existing zone ID and edit values for zones that are already in eqgame (e.g. change long name or short name refs), or you can add a new, unique zone id to create from scratch zones (The example 787 entry is a new zone id). Also note, zones below 200 are not replacable due to how the hook eq-core uses hooks into them. If you want to replace a zone below 200, you'll need to either edit the eqgame directly, or clone the zone to a higher new zoneidnumber.

- Follow the [guide on the eq-core-dll repo](https://github.com/xackery/eq-core-dll/) to build your DLL
- Enable custom zones via `bool areCustomZonesEnabled = true;` in _options.h
- Add a new entry for your zone:
```
static ZoneEntry Zones[] = {
    // zoneType, zoneID, zoneShortName, zoneLongName, eqStrID, zoneFlags2, x, y, z
    ZoneEntry(0, 787, "hollows", "Darkened Hollows", 35154, 4, 0, 0, 0),
};
```
- Update your zone table with SQL, you will see fields such as zoneidnumber, short_name, long_name. Make sure they match the eq-core-dll entries. (This affects the "You have entered" message after zoning)
- Edit eqstr_us.txt and ensure the eqStrID exists, and has your zone long name in it. (This affects the text when at character select informing a player of their current zone)

## EQGame Patching

This guide is not yet covered in detail, but essentially you use a hex editor and search for the zone short name in your eqgame.exe, you'll see data near it similar to the info provided for injection, and can edit the values. You'll still need to distribute the patched eqgame with something like [eqemupatcher](https://github.com/xackery/eqemupatcher), but keep in mind eqemupatcher checks the CRC of eqgame, so it may need additional fixing for that too.  

## EQGame Zone Cloning

To clone a zone, you can use the following steps:
- Let's assume for this example you want to clone East Commonlands to a new zone called "North Commonlands". So ecommons.s3d to ncommons.s3d

- Pick a zoneidnumber unused above 786. Let's use 787

- In your zones SQL table, find the ecommons row and duplicate it. Change the id to a new unique number, zoneidnumber to 787, short_name to ncommons, long_name to North Commonlands

- Copy the s3d/eqg's related to your original zone's shortname, e.g. ecommons.s3d and any other files related to ncommons.s3d. (Same with .eqg)

- Use eqzip, quail-gui or another tool and open the s3d/eqg archive, find the ecommons.wld inside it and rename it to ncommons.wld. (In the case of an eqg, there'll be a .zon file. but it shouldn't need to be renamed)

- Edit eq-core-dll and add the _option entry to match the settings noted above, 787, ncommons, etc

- Edit eqstr_us.txt and add a new entry for North Commonlands 