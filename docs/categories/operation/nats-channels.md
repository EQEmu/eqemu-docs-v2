# NATS Channels



**Note that NATS is currently on an experimental branch**

IN and OUT are from the perspective of a eqemu. So, if IN is yes, that means it is expected that a third party client will request messages to this channel, If OUT is yes, that means you can subscribe as a third party and get a feed of data.

#### Global Scoped

* **global.admin\_message.out** - eqproto::AdminMessage - Admin related communication. This is called from both zone and world, and may contain sensitive information, designed to be sent to an administrator-only channel. \(hacker, zone bootup, new account creations, etc\)

#### World Scoped

* **world.command\_message.in** - eqproto::CommandMessage - Send a command. Commands are as follows:
* who - Who is currently online
* **world.channel\_message.out** - eqproto::ChannelMessage - Any world-wide messages are broadcasted

#### Zone Scoped

* **zone.ecommons.channel\_message.in** - eqproto::ChannelMessage - Send a channel message to zone
* **zone.ecommons.channel\_message.in** - eqproto::CommandMessage - Send a command message to zone
* zone.ecommons.entity.event.\#\|No\|Yes\|EntityEvent\|Subscribe to events from entity \#
* zone.ecommons.entity.list\|Yes\|No\|Entities\|Request reply of a list of all entities
* zone.ecommons.channel\_message\|Yes\|No\|ChannelMessage\|Send a channel message to all provided zone

