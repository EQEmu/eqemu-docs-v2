# Client Configuration

EQEmulator is a custom, completely from-scratch, open source server implementation for EverQuest, built mostly on C++. EQEmulator is an open-source project, released under the GNU General Public License.

{% hint style="warning" %}
**EverQuest clients are the intellectual property of the Daybreak Game Company, LLC.** Copies of Daybreak Game Company intellectual property are not sourced through the EQEmulator project, nor should you provide unlicensed copies through any EQEmulator asset.
{% endhint %}

This page is intended to assist you in configuring your client to connect to an EQEmu server.

## All Operating Systems

* Navigate to your EverQuest directory and modify the eqhost file:

**Titanium** or **Secrets of Faydwer** Clients:

```text
[LoginServer]
Host=login.eqemulator.net:5998
```

**Seeds of Destruction**, **Underfoot**, or **Rain of Fear** Clients:

```text
[LoginServer]
Host=login.eqemulator.net:5999
```

## Operating System-specific Instructions

Once you've made the above changes to your eqhost file, you can follow the [**Windows Guide**](windows-client-configuration.md) or the [**mac OS Guide**](macos-mojave-client-configuration.md) for further customization.  

## Server-Specific Instructions

Each server will likely have some small modifications that you are required to make to your client.  These modifications correlate to the decisions the server operator made regarding their world, represented by the changes they have made to their database.

While we cannot cover every server, we can provide tips for the modified files that you should seek out for each new server for which you plan to connect.  

{% hint style="info" %}
It is highly recommended that you create a specific client folder for each server, so that you do not have to keep modifying these files.
{% endhint %}

#### Spell File:  spells\_%s.txt

This file contains the client-sided descriptions and effects for spells.  If the server operator has made changes to the spells in their world, you will likely NOT be able to see spell effects if you do not update this file.  This file goes in your main EverQuest directory.  

Example file name:  `spells_us.txt`

#### Spell Gems:  gemicons0%s.tga

These files are typically modified to show "old" icons, such as on a Time-Locked Progression server, or a "Classic" world.  These files need to be modified in your ../uifiles/default directory.

There are several files: `gemicons01.tga` through `gemicons03.tga`.  

#### Spell Icons:  spells0%s.tga

These files are typically modified to show "old" icons, such as on a Time-Locked Progression server, or a "Classic" world.   These files need to be modified in your ../uifiles/default directory.

There are several files:  `spells01.tga` through `spells07.tga`. 

#### Strings:  dbstr\_%s.txt

Database Strings are typically modified to customize things like alternate currency or AA descriptions, among many other uses, on a customized server.  This file goes in your main EverQuest directory.

Example file name:  `dbstr_us.txt`



