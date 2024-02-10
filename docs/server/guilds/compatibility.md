---
description: This page describes the backwards compatibility for the guild management system on your EQEmu Server.
---

### Ranks

Guild ranks will translate as follows when the server sends rank to the client.

|Guild Rank Name|RoF2|Titanium|Underfoot|
|:---|:---:|:---:|:---:|
|Leader|1|2|2|
|Senior Officer|2|1|1|
|Officer|3|1|1|
|Senior Member|4|0|0|
|Member|5|0|0|
|Junior Member|6|0|0|
|Initate|7|0|0|
|Recruit|8|0|0|

What does this mean?

If you are using a mixture of clients, the change of ranks will mirror the above translation.

Example:

If an RoF2 client changes the rank of Player A from Recruit to Junior Member, a Ti/UF client will show Player A as a member before and after the change.  

If a Ti/UF client changes the rank of Player A from Member to Officer, a RoF 2 client will show Player A as Officer.  

Player A will then be afforded the permissions as noted within the RoF2 client for the rank.

### Permissions

When creating a new guild, the following permissions are set as a default.

![image](https://github.com/EQEmu/Server/assets/65987027/5e5d4a21-8b44-4f05-b185-774141ef0f18)


### Guild Bank

By default, the guild leader is a guild banker.  This is accomplished based on the client type.

The Ti/Underfoot Client has a checkbox to identify a guild member as a banker.
The RoF2 Client uses the permission system to determine a guild members access to the guild bank.

Please note:  The system will convert RoF2 bank permissions to Ti/UF banker status when the following four permissions are set.

- Bank:Deposit Items
- Bank:Promote Items
- Bank:View Items
- Bank:Withdraw Items

Therefore, if your player base uses RoF2 generally though a few use Underfoot, those Underfoot players will have banker status based on the above noted RoF2 permissions.  The guild leader will not have to login with Underfoot and set their banker status.

Please note:  The internal structure of eqemu uses the RoF2 rank system.  Therefore if a member has the rank of Junior Member, and that rank has the 4 permissions noted above, when they log in with a UF client, they will have banker access, even though their UF guild rank will display as 'member' in the client as noted in the above translation table.
