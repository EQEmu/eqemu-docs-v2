# Introduction

<div style="text-align:center">
    <div style="width: 600px; height: 400px; position: relative; top: 0; left: 25%;">
        <img 
            style="position: absolute; top: 1%; left: 1%; border-radius: 180px"
            width="600" src="https://user-images.githubusercontent.com/3319450/87238998-55010c00-c3cf-11ea-8db5-3be25a868ac8.png" alt="AkkStack">
        <img 
            style="position: absolute; top: 260px; left: 22%; border-radius: 180px"
            src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQtksTtU5AvlNwK2Iq4tctmy9aVU6t6U7iSsITW59EmzQ&s" alt="Docker">
    </div>
</div>


<hr>

**AkkStack** is a simple **Docker Compose** environment that is augmented with **developer** and **operator** focused
tooling for running **EverQuest Emulator servers**.

You can have an entire server running within **minutes**, configured and ready to go for **development** or **production** use!

This is what I've used in production, battle-tested, for almost 5+ years. I've worked through a lot of issues to give you
the final stable product.

It's what I've also used exclusively for development as well as many other developers!

<hr>

## Features

- [x] **[Perl 5.24.4 x64](https://strawberryperl.com/)** (Quests)
- [x] **[Lua Server Runtime](https://www.lua.org/about.html)** (Quests)
- [x] **[MariaDB x64 10.x](https://mariadb.org/)**
- [x] **[Latest PEQ Database](http://db.projecteq.net/)**
- [x] **[Latest PEQ Quests](https://github.com/ProjectEQ/projecteqquests)**
- [x] **[V2 Server Side Maps](https://github.com/Akkadius/eqemu-maps)**
- [x] **[Optimized and latest server binaries](https://github.com/EQEmu/Server/releases) (Stable)**
- [x] **[Spire Web Admin / Content Editor](https://github.com/akkadius/spire)**
- [x] **[ProjectEQ Editor](https://github.com/ProjectEQ/peqphpeditor)**
- [x] **SSH in the eqemu-server container on port 2222**
- [x] **make menu for managing the eqemu-server container**
- [x] **make menu for managing the host-level container environment**
- [x] **Cron Jobs for the eqemu-server container**
- [x] **Startup Scripts for the eqemu-server container**
- [x] **Docker Compose v2 for easy management**
- [x] **Developer and Operator focused tooling**
- [x] **Developer tuned environment for easy development**
- [x] **Operator tuned environment for easy management**

### Containerized Services

| **Service**      | **Description**                                                                                                                                                                              |
|------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **eqemu-server** | Runs the emulator server and all services related to the emulator server                                                                                                                     |
| **mariadb**      | MySQL service                                                                                                                                                                                |
| **phpmyadmin**   | (Optional) PhpMyAdmin which is automatically configured behind a password proxy                                                                                                              |
| **peq-editor**   | (Optional) PEQ Editor which is automatically configured                                                                                                                                      |
| **ftp-quests**   | (Optional) An FTP instance fully ready to be used to remotely edit quests                                                                                                                    |
| **backup-cron**  | (Optional) A container built to automatically backup (Dropbox API) the entire deployment and perform database and quest snapshots for with different retention schedules defined in **.env** |

### Spire Web Admin

[Spire](https://github.com/akkadius/spire) is a powerhouse web admin panel as well as rich server content editor. It is continually being developed and new features being added regularly.

<img src="https://github.com/Akkadius/akk-stack/assets/3319450/caa916d6-349f-4f86-8a5f-c504a8b1e1ac" style="border-radius: 10px">

### PEQ Editor

Automatically configured with pre-set admin password, listens on port **8081** by default

<img src="https://user-images.githubusercontent.com/3319450/87240902-3dcc1980-c3e3-11ea-9d1e-746e217b4459.png">

### PhpMyAdmin

Automatically configured **PhpMyAdmin** instance with pre-set admin password (Behind a password protected proxy) listens on port 8082 by default

<img src="https://user-images.githubusercontent.com/3319450/87240916-63f1b980-c3e3-11ea-8dd8-93bca87f54ec.png">

<img src="https://user-images.githubusercontent.com/3319450/87240919-6f44e500-c3e3-11ea-8c56-6fe0e5ecef89.png">

### CLI Menus

#### Embedded Server Management CLI

Embedded server management CLI makes for quick and easy management of the server

<img src="https://user-images.githubusercontent.com/3319450/87240603-7c140980-c3e0-11ea-9e92-ce18edcfad29.gif" style="border-radius: 10px">

#### In Container Command Menu

A **make** menu to manage the in-container environment, need to be in home directory **cd ~/** to use

<img src="https://user-images.githubusercontent.com/3319450/87240694-779c2080-c3e1-11ea-8330-26d8add10e5f.gif" style="border-radius: 10px">

#### Host-Level Command Menu

A **make** menu to manage the host-level container environment

<img src="https://user-images.githubusercontent.com/3319450/87240726-bfbb4300-c3e1-11ea-80ac-e53bfa3386f4.gif" style="border-radius: 10px">

### SSH

**eqemu-server** starts with SSH server, the installation provides a generated 30+ character password, persistent keys through reboot. 

!!!info 
    Default port is **2222**

### Cron Jobs

Cronjob support via crontab is available in the **eqemu-server** container. 

Edit **~/assets/cron/crontab.cron** directly and the file watcher will install new crontab changes. (You cannot use crontab -e)

For example, PEQ has a configured database dump cron that feeds https://db.projecteq.net

```
eqemu@f30bb0b5bd3c:~$ cat ~/assets/cron/crontab.cron 
0 3 * * * cd /home/eqemu/server && ./scripts/peq-dump.sh && curl -F 'data=@/tmp/peq-latest.zip' http://db.projecteq.net/api/v1/dump/upload\?key\=apikey && rm /tmp/peq-latest.zip
```

### Startup Scripts

The **eqemu-server** container will start applications or scripts in **~/server/startup/*** folder. 

!!! warning 
    Do not try to run eqemu services here as they are managed by Spire

This is useful for running custom scripts or applications on startup, for example, a Discord bot

```
eqemu@f30bb0b5bd3c:~$ ls -lsh ~/server/startup
total 8.7M
8.7M -rwxr-xr-x 1 eqemu eqemu 8.7M Jun 10  2019 discordeq
```

### Automated Backups

There is a **backup-cron** container that is optional to use, it will automatically backup the entire deployment and perform database and quest snapshots for with different retention schedules defined in **.env**

<img src="https://github.com/Akkadius/akk-stack/assets/3319450/4f2240cb-4e5a-4433-b7dc-1f22a7b8a40b" style="border-radius: 10px">
