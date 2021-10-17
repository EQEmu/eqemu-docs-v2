---
description: Register your World Server with your Login Server (optional)
---

# World Registration

One of the options in the configuration is to be able to set whether or not world server registration is required. Here were are going to explain the differences in requirements

## Main Referenced Configuration Options

The below snippet is part of the current options that drive registration functionality

{% code title="login.json snippet" %}
```javascript
 { 
  "worldservers": {
    "unregistered_allowed": true, // allows worldservers to connect to your loginserver without server admin authentication
    "reject_duplicate_servers": false // if enabled, rejects duplicate worldservers
  }
}
```
{% endcode %}

## Unregistered World Servers

Most world servers that blindly target a login server are considered **unregistered world servers **these servers are just fine and can operate simply; but whether or not you want to **register** a server to an admin depends upon your needs. Most people in LAN environments will not care about registration and just want to get into the game as quick as possible!

## Registered World Servers

Whether or not you require servers to be registered is up to you and configurable in your config file. However, **most** servers allow **unregistered servers** but allow servers to be **registered** so that those servers have an **authentic owner** and can be promoted to a higher **status**

If a server is registered, the login sever will expect that the world server entry has an associated world admin in the `login_world_servers` table.** If you don't have a server admin account created, you can do so using the CLI**

```
./loginserver world-admin:create serveradmin serverpassword test@test.com

> Executing CLI Command

[LoginServer] [Info] Attempting to create world admin account | username [serveradmin] encryption algorithm [SCrypt] (14)
[LoginServer] [Info] Account creation success for user [serveradmin] encryption algorithm [SCrypt] (14) new admin id [1]
```

#### Result

```
MariaDB [peq]> select * from login_server_admins limit 10;
+----+--------------+-------------------------------------------------------------------------------------------------------+------------+-----------+---------------+---------------------+-------------------------+
| id | account_name | account_password                                                                                      | first_name | last_name | email         | registration_date   | registration_ip_address |
+----+--------------+-------------------------------------------------------------------------------------------------------+------------+-----------+---------------+---------------------+-------------------------+
|  1 | serveradmin  | $7$C6..../....1dvF3Cx9hdtzT2T6kBsAr4sq.53rWo1zjNr2BXDxUx.$fBFmF0JiiHya2f0chj17hJFuzTNLzcrTkVxKpeTmgo3 |            |           | test@test.com | 2019-09-17 07:19:21 |                         |
+----+--------------+-------------------------------------------------------------------------------------------------------+------------+-----------+---------------+---------------------+-------------------------+
1 row in set (0.00 sec)
```

Now we have server admin **1** 

If you have had a server try to login already, it should create a world server entry that is unregistered

```
[LoginServer] [Info] Server [Akkas Docker PEQ Installer (L)] [peqtest] did not login but unregistered servers are allowed
```

We would see an entry as follows

```
MariaDB [peq]> select * from login_world_servers;
+----+--------------------------------+------------+-----------------+---------------------------+---------------------+-----------------+-----------------------+-------------------+------+
| id | long_name                      | short_name | tag_description | login_server_list_type_id | last_login_date     | last_ip_address | login_server_admin_id | is_server_trusted | note |
+----+--------------------------------+------------+-----------------+---------------------------+---------------------+-----------------+-----------------------+-------------------+------+
|  1 | Akkas Docker PEQ Installer (L) | peqtest    |                 |                         3 | 2019-09-17 07:09:18 |                 |                     0 |                 0 | NULL |
+----+--------------------------------+------------+-----------------+---------------------------+---------------------+-----------------+-----------------------+-------------------+------+
1 row in set (0.00 sec)
```

The main thing here is that we need to make sure the `login_server_admin_id` column is set to the `id` of the server owner. As of this writing there is no CLI call or API call to perform these actions and sounds like a great PR for someone to make in the future

Once the `login_server_admin_id` is set to an appropriate server admin owner, as long as the respective authenticating world server tries to connect the next time with a configured `account` and `password` in their `loginserver` block of their `eqemu_config.json` they should authenticate against this table properly and will be authenticated as a registered server

