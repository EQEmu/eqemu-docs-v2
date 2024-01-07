# Basic Server Install - Windows

### Operating System Compatibility

* Windows 10 - x64
* Windows 11 - x64

### Requirements

This installer assumes you have a **clean** install of Windows and an Internet Connection 

## Runtime

2-10 minute(s) install depending on your connection

## What's in the Installer

- [x] **[Perl 5.24.4 x64](https://strawberryperl.com/)** (Quests)
- [x] **[Lua Server Runtime](https://www.lua.org/about.html)** (Quests)
- [x] **[MariaDB x64 10.x](https://mariadb.org/)**
- [x] **[Latest PEQ Database](http://db.projecteq.net/)**
- [x] **[Latest PEQ Quests](https://github.com/ProjectEQ/projecteqquests)**
- [x] **[V2 Server Side Maps](https://github.com/Akkadius/eqemu-maps)**
- [x] **[Optimized and latest server binaries](https://github.com/EQEmu/Server/releases) (Stable)**
- [x] **[Spire Web Admin / Content Editor](https://github.com/akkadius/spire)**

!!! info
      More installer info found in the installation details

## Download

First you will need to [download the install script](https://github.com/Akkadius/eqemu-install-v2/releases/download/static/eqemu-install.bat). The install script is a thin batch script that downloads the installer and kicks the install process off.

![image](https://github.com/EQEmu/eqemu-docs-v2/assets/3319450/d77efa93-c648-4688-8956-bc4710b6cf22)

Your browser may warn you that the file is not commonly downloaded, this is because the file is new and has not been downloaded many times yet. You can safely ignore this warning and download the file.

![image](https://github.com/EQEmu/eqemu-docs-v2/assets/3319450/e245bd55-732a-4286-a740-333b50a282e7)

## Run the Installer

You can start the installer anywhere. Assume you run it right out of your downloads folder.

Right click **eqemu-install.bat** and **Run as Administrator** and you will be presented with the following screen.

!!! info Run as Administrator

      The installer only needs administrator privileges to install Perl, MariaDB and configure firewall rules. Everything else should be ran as a normal user.

![image](https://github.com/EQEmu/eqemu-docs-v2/assets/3319450/30c8e725-f72e-4f48-9643-35b7e5424406)

You will be presented with a series of prompts to configure your server installation. For most users, you will not need to change anything and you can accept the defaults.

![image](https://github.com/EQEmu/eqemu-docs-v2/assets/3319450/2e618f17-2f5f-482c-8bcc-256e071ac9bd)

## Installation Details

### Standard Directories

| Directory                            | Details                                                                                  |
|--------------------------------------|------------------------------------------------------------------------------------------|
| C:\Program Files\MariaDB 10.11       | Where MariaDB is installed to. You usually don't need to do anything here.               |
| C:\Strawberry\perl                   | Where Perl is installed to. You usually don't need to do anything here.                  |
| C:\Users\<user\>\code                | Where EQEmu Server source code is located if you wish to compile yourself (not required) |
| C:\Users\<user\>\server              | Where the main server folder is located                                                  |
| C:\Users\<user\>\server\maps        | Where server side maps are located                                                       |
| C:\Users\<user\>\server\quests      | Where server side quests are located                                                     |
| C:\Users\<user\>\server\shared      | Where server side shared memory mapped files are located                                 |
| C:\Users\<user\>\server\logs        | Where server logs are located                                                            |
| C:\Users\<user\>\server\plugins     | Where server quest perl plugins are located                                              |
| C:\Users\<user\>\server\lua_modules | Where server quest lua plugins are located                                              |
| C:\Users\<user\>\server\assets      | Where server client patches, opcodes, misc are located                                   |
| C:\Users\<user\>\server\bin         | Where server executables are located                                   |

### Shortcut Scripts

| Script                        | Purpose                                                                                            |
| ----------------------------- |----------------------------------------------------------------------------------------------------|
| server_start.bat              | Starts the server using Spire launcher                                                             |
| server_stop.bat               | Stops the server using Spire launcher                                                              |
| server_restart.bat            | Restarts the server using Spire launcher                                                           |
| spire_start.bat | Starts the Spire web service                                                                       |
| spire_stop.bat | Stops the Spire web service                                                                        |
| spire_web.bat | Opens the Spire web service in your default browser                                                |
| spire_web_admin.bat | Opens the Spire web service in your default browser and navigates to the server admin panel /admin |

### Ports Forwarded

| Port      | Purpose                     |
| --------- | --------------------------- |
| 9000-9001 | World process               |
| 7000-7500 | Zone processes              |
| 7778      | UCS / Mail                  |
| 5998      | Loginserver (Titanium)      |
| 5999      | Loginserver (SoD+)          |

### Install Finished

When the installation process is done running, you will see **"press any key to continue"** when everything should be installed and ready to go.

The installer will do a few things for you at the end

* Open up your **install_config.yml** this is all of the inputs you provided to the installer or randomly generated passwords created during the install process. You can use this file to reference your passwords and other information.
* Opens up Spire web service in your default browser. This is a web service that allows you to manage your server. You can use this to start/stop/restart your server, edit content, manage accounts, manage zones, and more. Spire is continually ever evolving and is the recommended way to manage your server.
* Opens up the server folder in your file explorer. This is where all of your server files are located. You can use this to edit configuration files, add custom content, and more.

![image](https://github.com/EQEmu/Server/assets/3319450/3408489e-6286-4ef7-a2fc-41cd3949aa6f)

![image](https://github.com/EQEmu/Server/assets/3319450/a2451c34-7715-411e-9404-2c36416e4322)

![image](https://github.com/EQEmu/eqemu-docs-v2/assets/3319450/f5b2c5bc-cd29-423c-80c7-ce1fe4605775)

## Post-Installation

### Change Server Name

!!! info

      Change your server name in **eqemu_config.json** under **shortname** and **longname**

### Set a GM Account (Optional)

To set your account as a GM, you will need to edit the **account** table in your database and set your account status to 255 once you've logged in for the first time. You can use a tool like [HeidiSQL](https://www.heidisql.com/) to connect to your database and edit the table. Navicat is also a great tool for this.

### Login Server

The windows installer comes pre-configured with a **Local Login Server**

!!! info

      The server installer is configured to connect to the both the **EQEmu Public Login Server** and the **Local Login Server**

You can use **both** methods of logging into your server

**Login Methods**

* EQEmu Public Login Server
* Local Login Server (Private LAN Play)

Many server operators choose to keep things simple by only using the public login server

!!! info

      For information regarding the configuration and operation of a **Local Login Server; **please refer to the documentation in the [login server configuration documentation](../../login-server/login-server-configuration/)

### Connecting Clients to a Login Server

!!! info

      Remember that all **EverQuest Clients** on the local network will need to have a modified **eqhost.txt** (Client file) file that points at your Local Login Server

```text
[LoginServer]
Host=192.168.0.1:5999
```

Be sure to adjust the IP address to the IP address of your login server on your local network.

!!! info

      **Port 5998** is required when using the **Titanium** or **Secrets of Faydwer** clients
      **Port 5999** is required when using the **Seeds of Destruction, Underfoot** or **Rain of Fear** client and above

### Starting Server with Login Server (Optional)

You will need to go into Spire admin and click the "Power" button and ensure you have the login server checked. This will start the login server.

![image](https://github.com/EQEmu/Server/assets/3319450/d71bcaa8-493a-49f9-b206-9656c6de5cf3)


### Connecting to Local Login Server

!!! info

      For more Local Login Server configuration options; see [login server configuration documentation](../../login-server/login-server-configuration/)


Connect using your game client and you should see the server select screen

By default; the **login.json** used in the installer is set to **auto_create_accounts** when an account does not exist locally; making it very easy to get up and started. Resetting local login server accounts can be done through the login server command line also referenced in the [login server cli documentation](../../login-server/login-server-cli-management-interface/)

```json
  "account": {
    // ideal for local LAN setups, if you want a login attempt to automatically create an account
    // this will automatically create the account using the username and password if it doesn't exist
    "auto_create_accounts": true
  },
```

![Sample Server Select screen with local loginserver](https://user-images.githubusercontent.com/3319450/34912582-a025e892-f8aa-11e7-8676-2cdd98f6592c.png)
