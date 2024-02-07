# Installation

## Pre-requisites

### Install Docker

Linux Host or VM with [Docker Installed](https://docs.docker.com/engine/install/debian/) along with [Docker Compose v2](https://docs.docker.com/compose/install/linux/#install-the-plugin-manually)

It doesn't matter what Linux OS you use as long as it has **Docker** and **Docker Compose**, my personal recommendation is **Debian**.

After you install, make sure you install docker compose (listed below) and follow the instructions to run Docker as non-root

https://docs.docker.com/engine/install/linux-postinstall/

### Installing Docker Compose

With the latest docker compose v2; Docker does not provide the easiest documentation for installing Docker Compose as `docker-compose` instead of `docker compose`.

```
DOCKER_CONFIG=${DOCKER_CONFIG:-$HOME/.docker}
chmod +x $DOCKER_CONFIG/cli-plugins/docker-compose
sudo rm /usr/local/bin/docker-compose
sudo ln -s $DOCKER_CONFIG/cli-plugins/docker-compose /usr/local/bin/docker-compose
```

Confirm that it's working

```
docker-compose -v
Docker Compose version v2.2.3
```

## Installing Akk-Stack

First clone the repository somewhere on your server, in this case I'm going to clone it to an `/opt/eqemu-servers` folder in a Debian Linux host with Docker installed

```
git clone https://github.com/Akkadius/akk-stack.git peq-test-server
```

```
root@host:/opt/eqemu-servers# git clone https://github.com/Akkadius/akk-stack.git peq-test-server
Cloning into 'peq-test-server'...
remote: Enumerating objects: 57, done.
remote: Counting objects: 100% (57/57), done.
remote: Compressing objects: 100% (42/42), done.
remote: Total 782 (delta 14), reused 52 (delta 11), pack-reused 725
Receiving objects: 100% (782/782), 101.94 KiB | 7.28 MiB/s, done.
Resolving deltas: 100% (437/437), done.
```

Change into the new directory that represents your server

```bash
root@host:/opt/eqemu-servers# cd peq-test-server/
```

### Initialize the Environment

There are a ton of configuration variables available in the **.env** file that is produced from running the next command, we will get into that later. The key thing here is that it creates the base **.env** and scrambles all of the password fields in the environment

```shell
root@host:/opt/eqemu-servers# make init-reset-env
make env-transplant
Wrote updated config to [.env]
make env-scramble-secrets
Wrote updated config to [.env]
```

### Initialize Network Parameters

The next command is going to initialize two large key things in our setup 

* The **ip address** we're going to use
* The **zone port range** we're going to use

Make sure that you only open as many ports as you need on the zone end, because `docker-proxy` will NAT all ports individually in its own docker userland which does take some time when starting and shutting off containers. 

The more ports you nail up, the longer it takes to start / stop. Since this is a test server, I'm only going to use 30 ports. 

This `make` command also configures the `eqemu_config.json` port and address parameters as well automatically for you

!!! Terminal
    ```
    make set-vars port-range-high=7030 ip-address=66.70.153.122
    ```

    Yields the following output

    ```
    Wrote [IP_ADDRESS] = [66.70.153.122] to [.env]
    Wrote [PORT_RANGE_HIGH] = [7030] to [.env]
    ```



## Install

From this point you're ready to run the fully automated install with a simple `make install`

An example of what this output looks like below (Sped up)

<img src="https://user-images.githubusercontent.com/3319450/87240353-7289a200-c3de-11ea-8afe-1b0a5ad8400e.gif">

## Post-Install

Now that you're installed we need to look at how we interact with the environment

To gain a bash into the emulator server we have two options, we can come through a docker exec entry or we can SSH into the container

### Deployment Info

To print a handy list of passwords and access URL's, simply use **make info** at the host level of the deployment

```
root@host:/opt/eqemu-servers/peq-test-server# make info
##################################
# Server Info
##################################
# Akkas Docker PEQ Installer
##################################
# Passwords
##################################
MARIADB_PASSWORD=1jo5XUzpY7lYOf5FmJKRBhUfGmnVzBN
MARIADB_ROOT_PASSWORD=mDI8gefiVEGjeiMCUMrZhMmKMWI101B
SERVER_PASSWORD=uVNjjlucE5H9UzUlziZfP16GQvsWJhe
PHPMYADMIN_PASSWORD=tD02XcNGoaIaV82wnnEnenp0V7p58V9
PEQ_EDITOR_PASSWORD=5X5o1E84SXQzjmxN86fLzuBFJyGEjN9
FTP_QUESTS_PASSWORD=Jqx3KxCZFkRA1aPqBJqMTSA1vA8uK4Y
##################################
# IP
##################################
IP_ADDRESS=66.70.153.122
##################################
# Quests FTP  | 66.70.153.122:21 | quests / Jqx3KxCZFkRA1aPqBJqMTSA1vA8uK4Y
##################################
# Web Interfaces
##################################
# PEQ Editor  | http://66.70.153.122:8081 | admin / 5X5o1E84SXQzjmxN86fLzuBFJyGEjN9
# PhpMyAdmin  | http://66.70.153.122:8082 | admin / tD02XcNGoaIaV82wnnEnenp0V7p58V9
# EQEmu Admin | http://66.70.153.122:3000 | admin / 82a71144a51c521283834f99daff5a
##################################
```

### Service Lifetime

By default each container / service in the **docker-compose.yml** is configured to restart unless stopped, meaning if the server restarts the Docker daemon will boot the services you had started initially which is the default behavior of this stack

Occulus and the eqemu-server entrypoint bootup script is designed to start the emulator server services when the server first comes up, so if you need to bring the whole host down, everything will come back up on reboot automatically

### Services to Boot

By default the whole deployment is booted post install, but for production setups maybe you only want the emulator server and the database server only. Simply bring everything down with either **make down** or **docker-compose down**

**make up** will by default only bring up eqemu-server and mariadb

```
root@host:/opt/eqemu-servers/peq-test-server# make up --dry-run
docker-compose up -d eqemu-server mariadb
```

If you want to single boot another service, such as the **peq-editor** simply **docker-compose up -d peq-editor** and you'll have the 2 main services as well as the editor booted

![dc-ps](https://user-images.githubusercontent.com/3319450/87241769-eb432b00-c3eb-11ea-9cbf-f48307981303.gif)

### Accessing the Admin Panel

By default, Occulus runs within the `eqemu-server` service container and is available on port 3000

To access your admin panel bash or ssh into your server and run config to see your web admin password (Or view it in make info mentioned before)

```
eqemu@97b8129b90b4:~$ config | jq '.["web-admin"]'
{
  "application": {
    "key": "dadbeb31-3073-43dc-a359-569737bb2746",
    "admin": {
      "password": "82a71144a51c521283834f99daff5a"
    }
  },
  "launcher": {
    "runLoginserver": false,
    "runQueryServ": false,
    "isRunning": true,
    "minZoneProcesses": 3
  }
}
```

### Updating the Admin Panel

Updating server binaries is as simple as running **make update-admin-panel** in the server shell at the root of home, it will kill the currently running panel, cycle it out, start it up. This is not service affecting for running servers with a launcher running.

### Updating Server Binaries

Updating server binaries is as simple as running **update** in the server shell, it will change directory to the source directory, git pull and run a build which will be immediately available the next time you boot a process

### Running Server Processes While Developing

While developing its easy to jump back and forth between compiling changes and running single processes

If you have camped to character select, you can run **kzone** which will kill all zones and simply typing **z** will boot a zone process in the background but will still display in the foreground of the shell

`world` `ucs` `shared` are all shorthands that also work anywhere in any folder in the shell (See below in compiling and developing)