!!! info

    Akkadius tips and tricks

I've been asked MySQL performance related questions over the years so I decided to put together a short reference guide for those who want to properly tune their MySQL instance for performance.

You can get very complicated with all of this if you want, but for a majority EQEmu Server use cases there is no need to and I'm going to keep this very simple.

Our EverQuest Emulator server code base is very optimized to run efficient queries and a lot of work has been done to do so over the years; however - it is still a good idea to make sure you have a few knobs tweaked in your MySQL setup

What makes me qualified to give advice on this topic? I've been running MySQL servers for well over 10 years and have ran 100's of MySQL nodes in high scale environments and have seen a thing or two.

## Server Operators

Make sure in your MySQL config `my.cnf` you've allocated enough `innodb_buffer_pool_size` to hold your dataset in memory. There are ways to calculate this in more usage oriented ways but for literally most people just set it to 1GB

```text
innodb_buffer_pool_size=1G
```

Done. Move on, no need to look at any other knobs you just covered a majority of your ground and you can move on with your life.

If you want to know why you're setting this variable, it is one of the most important settings for modern InnoDB storage engines. This buffer pool is memory allocated in which table caches, table indexes are stored and keep all of your table operations fast.

## MariaDB Config Location

The location of your MySQL config varies on your operating system

| OS      | Location |
| ----------- | ----------- |
| **Linux**      | /etc/mysql/my.cnf	       |
| **Windows**   | INSTALLDIR\data\my.cnf or my.ini        |

## Hardware

If you're hosting a sizeable server, make sure you're not running super slow hard disk drives (5400RPM) to be holding your on-disk data for MySQL. Most modern computers and servers are using SSD or NVMe disks these days so if you are rocking this kind of disk for hardware you will have zero issues.

## EverQuest Emulator Developers

Now this section is no longer for server operators but for developers.

* **No Selects in for-loops** Make sure you're not firing selects off constantly or in any hot paths. 
* **Bulk Load Where Possible** Pre-load and bulk load data where necessary and loop through it later instead of doing individual selects inside of for loops. We spent many years unearthing all of the sins we had in the code, we don't need any more!
* **Bulk Load Data in Memory Where Possible** If certain datasets can be held in memory that aren't very large, make use of that instead of querying the database often for data that isn't changing
* **Bulk Insert Where Possible** If you have multiple items being saved in batch, use bulk inserts (chained insert or replace into statements) to save data
