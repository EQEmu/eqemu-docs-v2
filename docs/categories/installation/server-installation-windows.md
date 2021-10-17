# Basic Server Install - Windows

### Operating System Compatibility

* Windows 10 - x64

### Requirements

This installer assumes you have a **clean **install of Windows and an Internet Connection 

## Runtime

2-10 minute(s) install depending on your connection

## What's in the Installer

* **Perl 5.24.4 x64**
* **MariaDB x64 10.x**
* **Latest PEQ Database**
* **Latest PEQ Quests**
* **Latest Plugins repository**
* **Visual Studio** Runtimes
* **Automatically added firewall rules**
* **V2 Server Side Maps**
* **Path Files**
* **LUA**
* **Optimized and latest server binaries (Stable)**
* **Loginserver**

{% hint style="info" %}
More installer info found in the installation details
{% endhint %}

## Server Installation

### **Download**

First you will need to download the starter files that will kick off the installation process found here: [EQEmu Server Installer Files](https://github.com/Akkadius/eqemu-install-v2/releases/download/1.0/eqemu_installer_files_x64.zip)

![](<../../.gitbook/assets/image (7).png>)

### Extract

Once you have the files downloaded, you need to make a folder somewhere that you would like your server to reside, the place I chose is **C:\eqemu**

Extract the files inside **eqemu_installer_files_x64.zip** to the folder **C:\eqemu**

![Zip file present in c:\eqemu](<../../.gitbook/assets/image (3).png>)

![Extracted files present in directory](<../../.gitbook/assets/image (1).png>)

### Run the Installer

Once extracted, right click **eqemu_install.bat** and **Run as Administrator;** from this point on the entire install will be automated and run on its own, this is assuming you have an internet connection so the installer can pull down all of the necessary server assets

![](<../../.gitbook/assets/image (4).png>)

## Installation Details

### Installed Components

| Component                       | Details                                                             |
| ------------------------------- | ------------------------------------------------------------------- |
| MariaDB x64 10.x (MySQL Server) | **User** root (Doesn't allow remote connections) **Password** eqemu |
|                                 | Installed to **C:\Program Files\MariaDB 10.0**                      |
| Strawberry Perl 5.24.4 x64      | Installed to **C:\Strawberry\perl\bin**                             |

### **Folders**

| **Folder**  | Purpose                                                  |
| ----------- | -------------------------------------------------------- |
| bin         | Server executables are located                           |
| logs        | Server executable logs are written                       |
| shared      | Server executable shared memory mapped files are located |
| quests      | Server quest scripts zones, npcs, spells, items etc.     |
| plugins     | Server quest plugins for Perl                            |
| lua_modules | Server quest plugins for Lua                             |
| assets      | Server client patches, opcodes, misc                     |
| maps        | Server side geometry map                                 |

### Scripts

| Script                        | Purpose                                                                                                               |
| ----------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| t_start_server.bat            | Starts server; UCS, World, QueryServ and 30 Zone Processes                                                            |
| t_start_server_with_login.bat | Starts server; Loginserver, UCS, World, QueryServ and 30 Zone Processes                                               |
| t_stop_server.bat             | Stops server                                                                                                          |
| t_database_backup.bat         | Helper script to easily create a backup of the MySQL database                                                         |
| t_set_gm_account.bat          | Helper script to easily set an account to GM status (Requires account logged in for the first time)                   |
| win_server_launcher.pl        | Launcher script; not to be ran directly, called from batch scripts and can be configured to adjust zone process count |
| t_server_crash_report.pl      | Helper script to help summarize crashes that have occurred on a server to give to a developer for troubleshooting     |

### Shortcuts

| Shortcuts                     | Description                            |
| ----------------------------- | -------------------------------------- |
| start_export_client_files.exe | Shortcut ./bin/export_client_files.exe |
| start_import_client_files.exe | Shortcut ./bin/import_client_files.exe |
| start_loginserver.exe         | Shortcut ./bin/loginserver.exe         |
| start_queryserv.exe           | Shortcut ./bin/queryserv.exe           |
| start_shared_memory.exe       | Shortcut ./bin/shared_memory.exe       |
| start_ucs.exe                 | Shortcut ./bin/ucs.exe                 |
| start_world.exe               | Shortcut ./bin/world.exe               |
| start_zone.exe                | Shortcut ./bin/zone.exe                |

### Ports Forwarded

| Port      | Purpose                     |
| --------- | --------------------------- |
| 9000-9001 | World process               |
| 7000-7500 | Zone processes              |
| 7778      | UCS / Mail                  |
| 5998      | Loginserver (Titanium)      |
| 5999      | Loginserver (SoD+)          |
| 6000      | Loginserver (HTTP Rest API) |

## Post-Installation

When the installation process is done running, you will see **"press any key to continue"** once the final step of adding firewall rules has been complete that's when everything should be installed and ready to go 

![Your installation should closely resemble above](<../../.gitbook/assets/image (5).png>)

### Change Server Name

{% hint style="info" %}
Change your server name in **eqemu_config.json **under** shortname **and** longname**
{% endhint %}

### Set a GM Account (Optional)

To set your account as a GM account; use the **t_set_gm_account.bat** helper script to set your loginserver account to be flagged. You will need to ensure you login to character select first in order for your account to be in the database

### Login Server Installation

The windows installer comes pre-configured with a **Local Login Server **

{% hint style="success" %}
The server installer is configured to connect to the both the **EQEmu Public Login Server** and the **Local Login Server**

You can use **both** methods of logging into your server

**Login Methods**

* EQEmu Public Login Server
* Local Login Server (Private LAN Play)
{% endhint %}

Many server operators choose to keep things simple by only using the public login server

{% hint style="info" %}
For information regarding the configuration and operation of a **Local Login Server; **please refer to the documentation in the [Login Server Space](https://eqemu.gitbook.io/eqemulator-loginserver/loginserver/configuration)
{% endhint %}

### Connecting Clients to a Login Server

{% hint style="success" %}
Remember that all **EverQuest Clients** on the local network will need to have a modified **eqhost.txt** (Client file) file that points at your Local Login Server
{% endhint %}

{% code title="" %}
```javascript
[LoginServer]
Host=192.168.0.1:5999
```
{% endcode %}

Be sure to adjust the IP address to the IP address of your login server on your local network.

{% hint style="success" %}
**Port 5998** is required when using the **Titanium** or **Secrets of Faydwer **clients\
**Port 5999** is required when using the **Seeds of Destruction, Underfoot** or **Rain of Fear** client and above
{% endhint %}

### Starting Server with Login Server (Optional)

Once you have configured your login server, restart your world server using the helper script; **t_start_server_with_login_server** and you should be good to go!

### Connecting to Local Login Server

{% hint style="info" %}
For more Local Login Server configuration options; see [Login Server Space](https://eqemu.gitbook.io/eqemulator-loginserver/loginserver/configuration)
{% endhint %}

Connect using your game client and you should see the server select screen

By default; the **login.json **used in the installer is set to **auto_create_accounts** when an account does not exist locally; making it very easy to get up and started. Resetting local login server accounts can be done through the login server command line also referenced in the [login server documentation](https://eqemu.gitbook.io/eqemulator-loginserver/loginserver/command-line-management-interface#help-menu)

```javascript
  "account": {
    // ideal for local LAN setups, if you want a login attempt to automatically create an account
    // this will automatically create the account using the username and password if it doesn't exist
    "auto_create_accounts": true
  },
```

![Sample Server Select screen with local loginserver](https://user-images.githubusercontent.com/3319450/34912582-a025e892-f8aa-11e7-8676-2cdd98f6592c.png)
