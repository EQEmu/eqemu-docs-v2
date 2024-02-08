!!! info

    This page shows the structure of the **akk-stack** file system and how to interact with it.

## Standard Server Directories

Within the **eqemu-server** container, the fileystem is ephermeral, meaning it will not persist through a reboot. The
following directories noted that are shared to the host and will persist through a reboot.

| Directory                | Details                                                                                                                                                                         |
|--------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **/opt/eqemu-perl/**     | Where our eqemu-specific version of Perl is installed to. You usually don't need to do anything here.                                                                           |
| **~/.ssh/**              | (Shared to the host) Where SSH client keys are persisted if you use them to clone custom Git repo's using SSH etc. Also used to store keys of users who access **eqemu-server** |
| **~/assets**             | (Shared to the host @./assets) This is for common server utilities, management scripts, crontab, these are usually a part of akk-stack exclusively                              |
| **~/code**               | (Shared to the host @./code) Where EQEmu Server source code is located if you wish to compile yourself (not required)                                                           |
| **~/code/build**         | Where the **CMake** configuration is held for **make** or **ninja**                                                                                                             |
| **~/code/build/bin**     | If you compile, where executables are built to and symlinked to **~/server/bin**                                                                                                |
| **~/server**             | (Shared to the host @./server) Where the main server folder is located.                                                                                                         |
| **~/server/backups**     | Where server automatic database backups go when there are server schema updates. These can be used to revert prior to an update if needed. Rare necessity but here if needed.   |
| **~/server/maps**        | Where server side maps are located                                                                                                                                              |
| **~/server/quests**      | Where server side quests are located                                                                                                                                            |
| **~/server/shared**      | Where server side shared memory mapped files are located                                                                                                                        |
| **~/server/logs**        | Where server logs are located                                                                                                                                                   |
| **~/server/plugins**     | Where server quest perl plugins are located                                                                                                                                     |
| **~/server/lua_modules** | Where server quest lua plugins are located                                                                                                                                      |
| **~/server/assets**      | Where server client patches, opcodes, misc are located                                                                                                                          |
| **~/server/bin**         | Where server executables are located                                                                                                                                            |
| **~/server/startup**     | Where server startup scripts are located, for running custom services not managed by Spire on container bootup                                                                  |

## Symlinked Directories

The following files are symlinked within the container.

These symlinks are created on bootup of the container and are not ones you need to manage yourself and are here for convenience.

This means that when the **~/code** directory is updated via **git pull** the opcodes and patch files are automatically updated.

When you compile new binaries, the symlinks are automatically updated to the new binaries in the **~/server/bin** directory.


``` 
~/assets/scripts/create-symlinks.pl 
	Symlinking Source: /home/eqemu/code/build/bin/export_client_files Target: /home/eqemu/server/bin/export_client_files
	Symlinking Source: /home/eqemu/code/build/bin/import_client_files Target: /home/eqemu/server/bin/import_client_files
	Symlinking Source: /home/eqemu/code/build/bin/loginserver Target: /home/eqemu/server/bin/loginserver
	Symlinking Source: /home/eqemu/code/build/bin/queryserv Target: /home/eqemu/server/bin/queryserv
	Symlinking Source: /home/eqemu/code/build/bin/shared_memory Target: /home/eqemu/server/bin/shared_memory
	Symlinking Source: /home/eqemu/code/build/bin/ucs Target: /home/eqemu/server/bin/ucs
	Symlinking Source: /home/eqemu/code/build/bin/world Target: /home/eqemu/server/bin/world
	Symlinking Source: /home/eqemu/code/build/bin/zone Target: /home/eqemu/server/bin/zone
# Symlinking patches
	Symlinking Source: /home/eqemu/code/utils/patches/patch_RoF.conf Target: /home/eqemu/server/assets/patches/patch_RoF.conf
	Symlinking Source: /home/eqemu/code/utils/patches/patch_SoD.conf Target: /home/eqemu/server/assets/patches/patch_SoD.conf
	Symlinking Source: /home/eqemu/code/utils/patches/patch_UF.conf Target: /home/eqemu/server/assets/patches/patch_UF.conf
	Symlinking Source: /home/eqemu/code/utils/patches/patch_RoF2.conf Target: /home/eqemu/server/assets/patches/patch_RoF2.conf
	Symlinking Source: /home/eqemu/code/utils/patches/patch_Titanium.conf Target: /home/eqemu/server/assets/patches/patch_Titanium.conf
	Symlinking Source: /home/eqemu/code/utils/patches/patch_SoF.conf Target: /home/eqemu/server/assets/patches/patch_SoF.conf
# Symlinking opcodes
	Symlinking Source: /home/eqemu/code/utils/patches/opcodes.conf Target: /home/eqemu/server/assets/opcodes/opcodes.conf
	Symlinking Source: /home/eqemu/code/utils/patches/mail_opcodes.conf Target: /home/eqemu/server/assets/opcodes/mail_opcodes.conf
# Symlinking plugins
	Symlinking Source: /home/eqemu/server/quests/plugins/ Target: /home/eqemu/server
# Symlinking lua_modules
	Symlinking Source: /home/eqemu/server/quests/lua_modules/ Target: /home/eqemu/server
```