---
description: This page describes the settings and options in your eqemu_config.json file.
---

# Configure your eqemu\_config

{% hint style="warning" %}
If you make use of the various services listed below, be sure to **open the corresponding ports** on your server / firewall / router to TCP and UDP traffic!
{% endhint %}

{% hint style="danger" %}
STOP!  Did you read the line above?  It's a **really** important tip.
{% endhint %}

| Legend |
| :--- |
| \*Required |
| Not required |

### Config Example

* Config format: .json
* Below is a working base config example from our installer:
* [https://github.com/Akkadius/eqemu-install-v2/blob/master/eqemu\_config.json](https://github.com/Akkadius/eqemu-install-v2/blob/master/eqemu_config.json)

### world

| Variable | Default | Description |
| :--- | :--- | :--- |
| \*shortname |  | This is the short name of your server, this shows up in a client's .ini file eg: servername\_charactername.ini |
| \*longname |  | This is the long name of your server, this shows up in a on the Loginserver |
| address |  | Not required, but binds the server to this address, default is to listen on all addresses |
| localaddress |  | Not required, but recommended to set for LAN setups so other local clients can connect properly |
| maxclients | -1 | This sets the max amount of clients that can connect to your server, -1 is unlimited |
| locked | false | This determines whether the server starts up locked or not, it takes a minimum status of 100 to get through locked state |
| key |  | This key is used for encryption between your server processes - make sure this is unique and random to your server |

#### telnet

* Subsection for world telnet sub-service

| Variable | Default | Description |
| :--- | :--- | :--- |
| \*ip | 127.0.0.1 | Telnet server IP \(0.0.0.0 would bind all addresses\) |
| \*port | 9000 | Telnet server port |
| \*enabled | false | Enables the telnet service |

### database

* Required for a connection to the MySQL database

| Variable | Default | Description |
| :--- | :--- | :--- |
| \*username | eq | MySQL username |
| \*password | eq | MySQL password |
| \*host | localhost | MySQL host |
| \*port | 3306 | MySQL port |
| \*db | eq | MySQL database name |

### qsdatabase

* Required for a connection to the query server MySQL database if you decide to use a different and external source and/or database for it \(Logging purposes or otherwise\)

| Variable | Default | Description |
| :--- | :--- | :--- |
| \*username | eq | MySQL username |
| \*password | eq | MySQL password |
| \*host | localhost | MySQL host |
| \*port | 3306 | MySQL port |
| \*db | eq | MySQL database name |

### mailserver

* Required for UCS Mailserver service

| Variable | Default | Description |
| :--- | :--- | :--- |
| \*host |  | Mailserver hostname - just use \(0.0.0.0\) |
| port | 7778 | Mailserver port |

### chatserver

* Required for UCS Chatserver service

| Variable | Default | Description |
| :--- | :--- | :--- |
| \*host |  | Chatserver hostname - just use \(0.0.0.0\) |
| port | 7778 | Chatserver port |

### zones

* Zone level configuration

| Variable | Default | Description |
| :--- | :--- | :--- |
| \*defaultstatus | 0 | Default status on account creation |

#### ports

| Variable | Default | Description |
| :--- | :--- | :--- |
| low | 7000 | Starting TCP/UDP port assignment for zones |
| high | 7999 | Ending TCP/UDP port assignment for zones |

### loginserver

| Variable | Default | Description |
| :--- | :--- | :--- |
| \*host | login.eqemulator.net | This is the hostname of the loginserver endpoint |
| \*port | 5998 | Loginserver endpoint port |
| legacy | 0 | When set to 1, will connect to old netcode \(IE current public LS\) |
| account |  | Account forum username for public EQEmu authentication, this is used with worldserver registration |
| password |  | Account forum password |

{% hint style="warning" %}
Your server will need to use port 5998 to communicate with the public EQEmu loginserver.  Do NOT change that value as you would for your eqhost file, based on client.
{% endhint %}

* Note: Multiple Loginserver endpoints can be established by using the same configuration above, however declaring your loginserver subsections with a number, below is an example

```text
  "world" : {
	   "loginserver1" : {
			"account" : "",
			"host" : "login.eqemulator.net",
			"legacy" : "1",
			"password" : "",
			"port" : "5998"
	   },
	   "loginserver2" : {
			"account" : "",
			"host" : "myloginserver.net",
			"password" : "",
			"port" : "5998"
	   },
  },
```

### directories

* Most of these are optional, but configurable settings for the server

| Variable | Default | Description |
| :--- | :--- | :--- |
| maps | Maps/ | Maps directory |
| quests | quests/ | Quests Directory |
| plugins | plugins/ | Plugins Directory |
| lua\_modules | lua\_modules/ | LUA Modules Directory |
| patches | ./ | Patches Directory |
| shared\_memory | shared/ | Shared Memory Directory |
| logs | logs/ | Logs Directory |

