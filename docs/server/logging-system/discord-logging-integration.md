This is meant to be purely a one-way, highly performing implementation of publishing data to Discord.

* Strictly one-way
* Highly performing
* Extends the EQEmu Logging system to include another category called
* Messages get batched in multiples of the same type to prevent from sending many individual messages
* Runs through UCS. If you do not use UCS; this integration will not send messages to Discord
* Runs a queue listener in its own thread on UCS so even there is rate limiting congestion, it will sleep and retry on its own thread not to interrupt the main UCS thread or networking threads
* Has new Quest API methods to leverage this functionality for logging events to player facing channels or administrative facing channels. For example maybe you want to have the hackers logging category be send to an admin channel and you want an in game world wide emote to broadcast to a player announcement channel when a certain event happens
* Optimizations to EQEmu Core Logging system

![image](https://user-images.githubusercontent.com/3319450/171353504-7b9f4bb1-87f5-444d-8e1a-0ea6189d42ea.png)

### New Table (discord_webhooks)

This table pertains to webhook mappings that will correspond to `discord_webhook_id` in the `logsys_categories` table

```sql
MariaDB [peq]> describe discord_webhooks;
+--------------+--------------+------+-----+---------+----------------+
| Field        | Type         | Null | Key | Default | Extra          |
+--------------+--------------+------+-----+---------+----------------+
| id           | int(11)      | NO   | PRI | NULL    | auto_increment |
| webhook_name | varchar(100) | YES  |     | NULL    |                |
| webhook_url  | varchar(255) | YES  |     | NULL    |                |
| created_at   | datetime     | YES  |     | NULL    |                |
| deleted_at   | datetime     | YES  |     | NULL    |                |
+--------------+--------------+------+-----+---------+----------------+
```

### New Fields (logsys_categories)

```sql
MariaDB [peq]> describe logsys_categories;
+--------------------------+--------------+------+-----+---------+-------+
| Field                    | Type         | Null | Key | Default | Extra |
+--------------------------+--------------+------+-----+---------+-------+
| log_category_id          | int(11)      | NO   | PRI | NULL    |       |
| log_category_description | varchar(150) | YES  |     | NULL    |       |
| log_to_console           | smallint(11) | YES  |     | 0       |       |
| log_to_file              | smallint(11) | YES  |     | 0       |       |
| log_to_gmsay             | smallint(11) | YES  |     | 0       |       |
| log_to_discord           | smallint(11) | YES  |     | 0       |       |
| discord_webhook_id       | int(11)      | YES  |     | 0       |       |
+--------------------------+--------------+------+-----+---------+-------+
7 rows in set (0.003 sec)
```

### Quest API

**Lua**

```lua
eq.discord_send(std::string webhook_name, std::string message);
```

**Perl**

```perl
quest::discordsend(string webhook_name, string message);
```

### How to Use

You first need will need a webhook to send messages to your Discord Channel. Assuming you are an admin of a Discord and/or ability to make Webhook integrations. You can follow this Discord guide https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks

Once you have the webhook you can install it in your `discord_webhooks` table like so. You will want to give you webhook a name if you want to later call it by the quest api

```sql
MariaDB [peq]> select * from discord_webhooks;
+----+--------------+--------------------------------------------------------------------------------------------------------------------------+------------+------------+
| id | webhook_name | webhook_url                                                                                                              | created_at | deleted_at |
+----+--------------+--------------------------------------------------------------------------------------------------------------------------+------------+------------+
|  1 | test         | https://discord.com/api/webhooks/971964565959114793/xxx-xxx                                                              | NULL       | NULL       |
|  2 | quest_api    | https://discord.com/api/webhooks/971964720372396062/xxx-xxx                                                              | NULL       | NULL       |
+----+--------------+--------------------------------------------------------------------------------------------------------------------------+------------+------------+
2 rows in set (0.001 sec)

```

Once you have a webhook created, you can use it in either a Quest API method or hook it into any of your logging categories per normal.

For example, you can see I have logging categories setup just to test the integration

```sql
MariaDB [peq]> select log_category_id, log_category_description, log_to_discord, discord_webhook_id from logsys_categories where discord_webhook_id > 0;
+-----------------+--------------------------+----------------+--------------------+
| log_category_id | log_category_description | log_to_discord | discord_webhook_id |
+-----------------+--------------------------+----------------+--------------------+
|              36 | MySQL Query              |              1 |                  1 |
|              53 | Info                     |              1 |                  1 |
+-----------------+--------------------------+----------------+--------------------+
2 rows in set (0.002 sec)
```

And then in my test channel you can see both Info and MySQL messages being broadcasted

![image](https://user-images.githubusercontent.com/3319450/171356608-14438726-8a84-4e07-9f9b-587ea1227054.png)
 

