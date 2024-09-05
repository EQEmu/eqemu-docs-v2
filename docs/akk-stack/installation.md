# Installation

## Pre-requisites

### Install Docker

It doesn't matter what Linux OS you use as long as it has **Docker** and **Docker Compose**, my personal recommendation is **Debian**.

- [x] Install [Docker](https://docs.docker.com/engine/install/debian/)
- [x] Install docker compose (listed below)
- [x] Follow the instructions to [run Docker as a non-root user](https://docs.docker.com/engine/install/linux-postinstall/)


!!! warning

    Do not run Docker as root, run docker as the user you are logged in as and add your user to the `docker` group

### Installing Docker Compose

```
sudo curl -SL https://github.com/docker/compose/releases/latest/download/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose && sudo chmod +x /usr/local/bin/docker-compose
```

Confirm that it's working

```
docker-compose -v
Docker Compose version v2.2.3
```

## Installing Akk-Stack

First clone the repository somewhere on your server, in this case I'm going to clone it to an `/opt/eqemu-servers` folder in a Debian Linux host with Docker installed

```
git clone https://github.com/Akkadius/akk-stack.git && cd akk-stack
```

```
$ git clone https://github.com/Akkadius/akk-stack.git
Cloning into 'akk-stack'...
remote: Enumerating objects: 57, done.
remote: Counting objects: 100% (57/57), done.
remote: Compressing objects: 100% (42/42), done.
remote: Total 782 (delta 14), reused 52 (delta 11), pack-reused 725
Receiving objects: 100% (782/782), 101.94 KiB | 7.28 MiB/s, done.
Resolving deltas: 100% (437/437), done.
```

### Initialize the Environment

There are a ton of configuration variables available in the **.env** file that is produced from running the next command, we will get into that later. The key thing here is that it creates the base **.env** and scrambles all of the password fields in the environment. 

!!! warning

    Do not run `make init-reset-env` after the environment has been initialized. Things will no longer work properly.

```shell
make init-reset-env
```

Will output the following.

```
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

<img src="https://user-images.githubusercontent.com/3319450/87240353-7289a200-c3de-11ea-8afe-1b0a5ad8400e.gif" style="border-radius: 10px">


## Post-Install

### Start / Stop

To start the server, simply use the **make up** command from the root of the **akk-stack** directory

```
make up
```

To stop the server, simply use the **make down** command from the root of the **akk-stack** directory

```
make down
```

### Deployment Info

To print a handy list of passwords and access URL's, simply use **make info** at the host level of the deployment

```
make info
```

!!! info
    Note this may look different depending on the services you have booted

```
----------------------------------
> Server Info
----------------------------------
> Akkas Test Bed (LDL)
----------------------------------
> Passwords
----------------------------------
MARIADB_PASSWORD=wNh6CrKiVq6oy5FPjQIMr7M18oAC7Ii
MARIADB_ROOT_PASSWORD=jrwe7Jv6sZRgrYig5vgoTEvBX6XGgxb
SERVER_PASSWORD=UXWrUXWv4MPZsJpUgbWuEPJn59ksNpc
PHPMYADMIN_PASSWORD=L6buMu5dzfIkhNTjh7LeMsxNdFfLUrA
PEQ_EDITOR_PASSWORD=jufMcw584ZDK3JRNJf4JB8z0e3Whoma
FTP_QUESTS_PASSWORD=gtL1yKDmZyC4eK9X85ZAytGdUVEgN62
----------------------------------
> IP
----------------------------------
IP_ADDRESS=192.168.65.62
----------------------------------
> Quests FTP  | 192.168.65.62:21 | quests / gtL1yKDmZyC4eK9X85ZAytGdUVEgN62
----------------------------------
> Web Interfaces
----------------------------------
> PEQ Editor  | http://192.168.65.62:8081 | admin / jufMcw584ZDK3JRNJf4JB8z0e3Whoma
> PhpMyAdmin  | http://192.168.65.62:8082 | admin / L6buMu5dzfIkhNTjh7LeMsxNdFfLUrA
> EQEmu Admin | http://192.168.65.62:3000 | admin / 2c9a88fa8470a70168080e5dbc8446
----------------------------------
> Spire Backend Development  | http://192.168.65.62:3010 | 
> Spire Frontend Development | http://192.168.65.62:8080 | 
----------------------------------
```
