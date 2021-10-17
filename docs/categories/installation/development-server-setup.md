# Development Server Setup



* As a developer - you may find the necessity to build a clean server/folder with the latest PEQ database without messing up an existing folder - today this is easy to do and applies for either Linux or Windows
* **This assumes you have Perl / MySQL and the rest of the environment already installed**

**Creating a New Folder**

* First, create your server folder that you wish use, the examples are going to be used with Linux, but the command line is almost identical for Windows
* In Linux - the base installer uses /home/eqemu/ - so let's use that directory as our base

```text
mkdir eqemu_test
cd eqemu_test
```

**Using eqemu\_server.pl**

* Now that we are in our server folder, we can either copy the eqemu\_server.pl from another server folder, or we can pull down a fresh copy from Github
* You can use curl, wget or manually create the file

```text
curl -O https://raw.githubusercontent.com/EQEmu/Server/master/utils/scripts/eqemu_server.pl && chmod 755 eqemu_server.pl && ./eqemu_server.pl new_server
```

```text
wget --no-check-certificate --cache=no https://raw.githubusercontent.com/EQEmu/Server/master/utils/scripts/eqemu_server.pl && chmod 755 eqemu_server.pl && ./eqemu_server.pl new_server
```

**Setting Environment Parameters**

* The script will prompt for a few questions, this is to ask for a valid MySQL user/password so the script can install a new environment properly, it is also going to ask for your new database name so it can associate the new folder with that database

```text
[New Server] For a new server folder install, we assume Perl and MySQL are configured
[New Server] This will install a fresh PEQ Database, with all server assets
[New Server] You will need to supply database credentials to get started...

[Input] MySQL User: eqemu
[Input] MySQL Password: eqemu
[New Server] Success! We have a database connection
[Input] Specify a NEW database name that PEQ will be installed to: peq_new
```

* Note: If you do not supply valid MySQL credentials - the command will halt and not continue. Make sure to also use a simple database name to prevent issues

**Waiting for Installation**

* Next, this is going to kick off installer routines - and build the source in the following directory:
  * **/home/eqemu/eqemu\_test**\_source/Server/build
  * When this is done compiling, the folder you created will by symlinked to this custom build directory
* Once complete, you will see installation summary info presented to you

  ```text
  [New Server] New server folder install complete
  [New Server] Below is your installation info:
  [Install] Installation complete...
  [Install] Server Info (Save somewhere if needed):
  - mysql_eqemu_db_name  peq_new
  - mysql_eqemu_user     eqemu
  - mysql_eqemu_password eqemu
  [Install] Linux Utility Scripts:
  - server_start.sh                      Starts EQEmu server (Quiet) with 30 dynamic zones, UCS & Queryserv, dynamic zones
  - server_start_dev.sh                  Starts EQEmu server with 10 dynamic zones, UCS & Queryserv, dynamic zones all verbose
  - server_stop.sh                       Stops EQEmu Server (No warning)
  - server_status.sh                     Prints the status of the EQEmu Server processes
  [Configure] eqemu_config.json         Edit to change server settings and name
  ```

* Type **exit** out of the menu and return back to the prompt

**Server Start Test**

* If necessary, return back to your original folder and start your respective start script whether you are in Windows or Linux, in Linux we will use ./server\_start.sh

```text
root@debian:/home/eqemu/eqemu_test# ./server_start.sh
[Status] Loading items...
[Status] Loading factions...
[Status] Loading loot...
[Status] Loading skill caps...
[Status] Loading spells...
[Status] Loading base data...
Server started - use server_status.sh to check server status

root@debian:/home/eqemu/eqemu_test#
Akka's Linux Server Launcher
Zones to launch: 30

root@debian:/home/eqemu/eqemu_test# ./server_status.sh

Akka's Linux Server Launcher
World: UP Zones: (30/30) UCS: UP Queryserv: UP
```

That's it!

Your new server partition should be ready for you to use and manipulate however you need

