# Updating Server Assets

We have an automated way for pulling down certain assets and maintaining parts of the server

The main script we use for this is **eqemu\_server.pl** in the server folder

**eqemu\_server.pl** is a required part of the server binary set.

World will download a fresh copy on bootup and then run the script to make sure no database schema updates need to be ran, it also does a handful of other very powerful things

### EQEmu Server Menu

* First start by going to your server directory and either running the script and going into the interactive prompt, or you can run commands to do things that you would like the script to do

```text
[Update] No script update necessary...

>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
>>> EQEmu Server Main Menu >>>>>>>>>>>>
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

 [database]     Enter database management menu
 [assets]       Manage server assets
 [new_server]   New folder EQEmu/PEQ install - Assumes MySQL/Perl installed
 [setup_bots]   Enables bots on server - builds code and database requirements

 exit

Enter a command #>
```

### Assets Menu

* The assets menu will show a variety of options depending on what OS you are running the script on \(Windows or Linux\)

Example: perl eqemu\_server.pl maps

```text

 [maps]                 Download latest maps
 [opcodes]              Download opcodes (Patches for eq clients)
 [quests]               Download latest quests
 [plugins]              Download latest plugins
 [lua_modules]          Download latest lua_modules
 [utility_scripts]      Download utility scripts to run and operate the EQEmu Server
>>> Windows
 [windows_server_download]      Updates server via latest 'stable' code
 [windows_server_latest]        Updates server via latest commit 'unstable'
 [windows_server_download_bots] Updates server (bots) from latest 'stable'
 [fetch_dlls]                   Grabs dll's needed to run windows binaries
 [setup_loginserver]            Sets up loginserver for Windows
```

#### Asset Commands

| Command | Action |
| :--- | :--- |
| maps | Download latest maps |
| opcodes | Download latest opcodes \(Patches for eq clients\) |
| quests | Download latest quests |
| plugins | Download latest plugins |
| lua\_modules | Download latest lua\_modules |
| utility\_scripts | Download utility scripts |
| setup\_loginserver | Windows Loginserver setup |
| linux\_login\_server\_setup | Loginserver setup for Linux |

#### Database Maintenance Commands

| Command | Action |
| :--- | :--- |
| backup\_database | Back up database to backups/ |
| backup\_database\_compressed | Back up database \(zip/tar\) to backups/ |
| backup\_player\_tables | Backs up only player tables to backsup/ |
| check\_db\_updates | Checks for database updates manually \(should not be needed\) |
| check\_bot\_db\_updates | Checks for bot database updates manually \(should not be needed\) |
| remove\_duplicate\_rules | Older databases used to have tons of duplicate rules - this can be ran to purge them for the new inheritance rules |
| drop\_db\_bots\_schema | Removes bot database schema |

#### Updating Windows Binaries - Stable

* **perl eqemu\_server.pl windows\_server\_download** will download EQEmu development team approved 'stable' binaries - you can check the date on the files as to when they were built

Example:

```text
[Copy] folder doesn't exist, creating 'updates_staged/'
[Update] No script update necessary...
[Update] Fetching Latest Windows Binaries...
[Download] Saved: (updates_staged/master_windows_build.zip) from https://raw.githubusercontent.com/Akkadius/EQEmuInstall/master/master_windows_build.zip
[Update] Fetched Latest Windows Binaries...
[Update] Extracting... ---
[Unzip] Extracting...
[Update] Installing :: common.lib
[Update] Installing :: eqlaunch.exe
[Update] Installing :: eqlaunch.ilk
[Update] Installing :: eqlaunch.pdb
[Update] Installing :: export_client_files.exe
[Update] Installing :: export_client_files.ilk
[Update] Installing :: export_client_files.pdb
[Update] Installing :: import_client_files.exe
[Update] Installing :: import_client_files.ilk
[Update] Installing :: import_client_files.pdb
[Update] Installing :: libeay32.dll
[Update] Installing :: libsodium.dll
[Update] Installing :: loginserver.exe
[Update] Installing :: loginserver.ilk
[Update] Installing :: loginserver.pdb
[Update] Installing :: luabind.lib
[Update] Installing :: queryserv.exe
[Update] Installing :: queryserv.ilk
[Update] Installing :: queryserv.pdb
[Update] Installing :: shared_memory.exe
[Update] Installing :: shared_memory.ilk
[Update] Installing :: shared_memory.pdb
[Update] Installing :: ucs.exe
[Update] Installing :: ucs.ilk
[Update] Installing :: ucs.pdb
[Update] Installing :: world.exe
[Update] Installing :: world.ilk
[Update] Installing :: world.pdb
[Update] Installing :: zone.exe
[Update] Installing :: zone.ilk
[Update] Installing :: zone.pdb
[Update] Done
```

#### Updating Windows Binaries - Unstable - Latest

* **perl eqemu\_server.pl windows\_server\_latest** - will download the latest compiled binaries from our **AppVeyor CI integration**

Example:

```text
[Update] No script update necessary...
[Update] Fetching Latest Windows Binaries (unstable) from Appveyor...
[Download] Saved: (updates_staged/master_windows_build_pdb.zip) from https://ci.appveyor.com/api/projects/KimLS/server/artifacts/build_x86_pdb.zip
[Download] Saved: (updates_staged/master_windows_build.zip) from https://ci.appveyor.com/api/projects/KimLS/server/artifacts/build_x86.zip
[Update] Fetched Latest Windows Binaries (unstable) from Appveyor...
[Update] Extracting... ---
[Unzip] Extracting...
[Unzip] Extracting...
[Update] Installing :: eqlaunch.exe
[Update] Installing :: eqlaunch.pdb
[Update] Installing :: export_client_files.exe
[Update] Installing :: export_client_files.pdb
[Update] Installing :: import_client_files.exe
[Update] Installing :: import_client_files.pdb
[Update] Installing :: loginserver.exe
[Update] Installing :: loginserver.pdb
[Update] Installing :: queryserv.exe
[Update] Installing :: queryserv.pdb
[Update] Installing :: shared_memory.exe
[Update] Installing :: shared_memory.pdb
[Update] Installing :: tests.exe
[Update] Installing :: tests.pdb
[Update] Installing :: ucs.exe
[Update] Installing :: ucs.pdb
[Update] Installing :: world.exe
[Update] Installing :: world.pdb
[Update] Installing :: zone.exe
[Update] Installing :: zone.pdb
[Update] Done
```

#### New Server Option

* This command is detailed and used in [Development Server Setup](../installation/development-server-setup.md)

