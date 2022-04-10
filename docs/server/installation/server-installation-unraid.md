# Basic Server Install - Unraid

This guide is intended for advanced users who wish to setup a private EQEmu server within Unraid including the login server using the AkkStack. 

## Prerequisites

### Setting up DevPack GUI

You need to install the DevPack GUI Community Application plugin as it will enable  you to install necessary build tools for installing **AkkStack**

1. Ensure you have the latest version of **Community Applications**
2. Navigate to *Apps* to and search for **DevPack GUI**
3. Install the DevPack GUI to your Unraid server
4. Once installed, navigate to *Plugins* and click on the **Dev Tool** icon
5. Install the following 
      * **gc-#.#.#-x86_64**
      * **guile-#.#.#-x86_64**
      * **make-#.#.#0x86_64**
6. Hit **Apply** & **Done**

### Creating a place for EQEmu to install & run

You need to create a location to pull down the server files. For this tutorial we will use **appdata** as our base since its often on an SSD cache. 

!!! info
    Ensure you have at least **7GB+** of available space before proceeding. 

### Clone the AkkStack

1. Log into your Unraid server
2. Open the terminal for Unraid by clicking on the icon in the navigation bar
3. Clone the AkkStack source code from github
   
```
git clone https://github.com/Akkadius/akk-stack /mnt/user/appdata/eqemu_server && cd /mnt/user/appdata/eqemu_server
```

### Pinning CPUs for EQEmu containers

It is highly desirable that you pin some specific cores of your Unraid server for EQEmu, as it will keep EQEmu from conflicting with other workloads on your server. This must be done in the **docker-compose.yml** file prior to building your images. 

!!! info
    If done after, the images will need to be rebuilt for the change to take affect.

1. Use either `nano` or `vi` text editor
```
nano docker-compose.yml
```
2. Add to the bottom of each service `cpuset: <core1>,<matchedcore>` to pin the AkkStack to specific cores

Example:
```javascript
    cap_add:
      # This is needed for running perf traces
      # - CAP_SYS_ADMIN
      # For GDB traces but it has security implications
      - SYS_PTRACE
    cpu_shares: 900
    cpuset: 5,21
```

!!! info
    Unfortunately, it appears setting a memory limit isn't possible.

### Initialize the environment

There are a ton of configuration variables available in the `.env` file that is produced from running the next command, we will get into that later. The key thing here is that it creates the base `.env` and scrambles all of the password fields in the environment

```
root@Tower:/mnt/user/appdata/eqemu_server# make init-reset-env
make env-transplant
Wrote updated config to [.env]
make env-scramble-secrets
Wrote updated config to [.env]
```

### Initialize network parameters

The next command is going to initialize two large key things in our setup

1. The Unraid server ip address we're going to use
2. The zone port range we're going to use; default is 30 which should be plenty for home use.

Make sure that you only open as many ports as you need on the zone end, because `docker-proxy` will NAT all ports individually in its own docker userland which does take some time when starting and shutting off containers. The more ports you nail up, the longer it takes to start / stop. Since this is a test server, I'm only going to use 30 ports. This `make` command also drives the `eqemu_config.json` port and address parameters as well automatically for you

```
root@Tower:/mnt/user/appdata/eqemu_server# make set-vars port-range-high=7030 ip-address=192.168.1.5
Wrote [IP_ADDRESS] = [192.168.1.5] to [.env]
Wrote [PORT_RANGE_HIGH] = [7030] to [.env]
```

### Set permissions for MariaDB

Since Unraid runs everything as root, the make file for the AkkStack will throw an error when it gets to MariaDB since the `my.cnf` file is wide open to the world. To prevent this from happening, you must change the permissions for the file prior to running `make`

```
chmod 0444 containers/mariadb/my.cnf
```

## Install

From this point you're ready to run the fully automated install with a simple `make install`

An example of what this output looks like below (Sped up) Be patient as this will take a while

<p align="center"><img src="https://user-images.githubusercontent.com/3319450/87240353-7289a200-c3de-11ea-8afe-1b0a5ad8400e.gif"></p>

## Post installation

There are a few cleanup / fixes that must be performed after installation to ensure you can connect.
git 
### Configuring world & login server

`eqemu_config.json` is configured to run in a local environment, not Unraid. There are a few tweaks you need to make to enable the **world server** to communicate properly with the **login server**

1. Open the **eqemu_config.json** file in your favorite linux text editor
```
nano server/eqemu_config.json
```
2. Scroll down to **shortname** and change to your shortname world server of choice, this will be prepended to character ini files
3. Edit the **longname** to what something clever for you
2. Change **localaddress** from `127.0.0.1` to your Unraid server IP. Ex. `192.168.1.5`

Example:
```javascript
"key": "DDC6pffVlPt0dAqYq8LjIh7BQmdeKkwc",
   "shortname": "TEQ",
   "longname": "Test EQEmu Docker Server",
   "localaddress": "192.168.1.5",
   "address": "192.168.1.5"
```
4. Save & exit the text editor
!!! info
    These changes will not take affect until you reboot EQEmu

### Ensure all EQEmu services are started

1. Run `make info` to display IP & port for EQEmu Admin (Occulus)
2. Open the page for EQEMu Admin & login with teh credentials provided above
3. Click on the **power button** and select **Power On**
4. Under *Launcher Options* ensure the following services are toggled to on
      * Run Shared Memory
      * Run Loginserver
      * Run QueryServ
5. Click on **Start Server** for the services to startup

### Docker container cleanup

AkkStack will download several containers and build from them. This has the side effect of polluting Unraid's docker container list & subsequent images. These should be deleted. 

!!! info
    You may delete the containers from the Unraid GUI, but it will not remove the images

1. From your Unraid terminal, click on the **Docker** tab
2. Click on a container with a generated name and click **Remove**
3. Leave the checkmark in the *also remove image* box & click **Yes, delete it!**
4. Repeat above for each generated container

!!! info
    You are likely to encounter the message that the image is unable to be deleted as its in use by another container. This is ok, you can clean these if you wish through the `docker image remove` terminal command.

### Startup on Unraid boot

If you wish to start up EQEmu on Unraid boot, you will need to have a simple script to run it.

1. Ensure you have the latest version of *Community Applications*
2. Navigate to *Apps* to and search for **User Scripts**
3. Install the **User Scripts** plugin to your Unraid server
4. Once installed, navigate to *Plugins* and click on the **User Scripts** icon
5. Click the button the **Add New Script**
6. Enter `start_eqemu_server` for the script name
7. Click on the *gear* icon, and select **Edit Script**
8. Paste the below into your script and click **Save Changes**
```
/mnt/user/appdata/eqemu_server/make up-all
```
9.  Click the drop down *Schedule Disabled* and select the startup option **At First Array Start Only**
10. Click **Done**