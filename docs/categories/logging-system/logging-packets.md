# Logging Packets

## Packet Logging

Useful for developers, it can be very handy to see what activity is happening realtime. We now have the ability to sync it to any output format simultaneously.

## Packet Categories

```text
select * from logsys_categories where log_category_description like '%packet%';
+-----------------+------------------------------------+----------------+-------------+--------------+
| log_category_id | log_category_description           | log_to_console | log_to_file | log_to_gmsay |
+-----------------+------------------------------------+----------------+-------------+--------------+
|               5 | Packet: Client -> Server           |              0 |           0 |            0 |
|              39 | Packet: Server -> Client           |              0 |           0 |            0 |
|              40 | Packet: Client -> Server Unhandled |              0 |           0 |            0 |
|              41 | Packet: Server -> Client With Dump |              0 |           0 |            0 |
|              42 | Packet: Server -> Client With Dump |              0 |           0 |            0 |
+-----------------+------------------------------------+----------------+-------------+--------------+
5 rows in set (0.04 sec)
```

## **Enable in Game \(Example\)**

Below would enable client -&gt; server direction packet logging to the client

```text
#logs set gmsay 5 1
```

## Server to Client

All server packets that are sent to the client, the only packet types that aren't sent to the client while in game are the message packets themselves because that would cause a loop

![](https://github.com/EQEmu/Server/wiki/images/llm7EXY-opt.gif?raw=true)

## Client to Server

Packets being sent from client to server

![](https://github.com/EQEmu/Server/wiki/images/8t4tkrB.gif?raw=true)

## Unhandled Packets

In this example - a client clicking a 'Krono' is unhandled by the server, at which it dumps a packet of data sent from the client

![](https://github.com/EQEmu/Server/wiki/images/XkPDXb9.gif?raw=true)

## Packet Logging Levels

* The main differentiator when deciding how much information you want to see with each packet, you use two different but very similar categories for each type. It all comes down to whether or not you want to see the packet dump data with the packet or not. As of this time, Packet: Client -&gt; Server Unhandled always dumps the packet because they are far less frequent.
  * Packet: Client -&gt; Server
  * Packet: Server -&gt; Client
  * Packet: Server -&gt; Client With Dump
  * Packet: Server -&gt; Client With Dump

## Packet Information \(With Dump\)

Below is an example of a full packet payload being output to the client

![](https://github.com/EQEmu/Server/wiki/images/C9SnDRD.gif?raw=true)

