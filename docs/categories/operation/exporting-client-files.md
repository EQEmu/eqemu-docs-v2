---
description: >-
  This page describes the process of exporting the various client files that are
  needed to keep the client in sync with your EQEmu server.
---

# Exporting Client Files

Your EQEmu Server comes with a built in script to automatically generate a number of key files.  Using this guide will help you get set up to output those files so that you can share them with your players.

{% hint style="warning" %}
As always, make backups of your files before replacing or updating them!
{% endhint %}

## Client Files

When we're discussion "Client Files", we are specifically referring to some key files that keep your client and server in sync:

| File Name | Content |
| :--- | :--- |
| BaseData.txt | An ASCII export of your [base\_data](https://eqemu.gitbook.io/database-schema/categories/client-files/base_data) table, which contains information regarding class "base" stats, such as HP, Mana and Endurance by level. |
| dbstr\_us.txt | An ASCII export of your [db\_str](https://eqemu.gitbook.io/database-schema/categories/client-files/db_str) table, which contains the text that is displayed to a player character from various sources, such as spells or trades. |
| SkillCaps.txt | An ASCII export of your [skill\_caps](https://eqemu.gitbook.io/database-schema/categories/client-files/skill_caps) table, which contains information about skills \(1HS, 1HB, Abjuration, etc.\) by class and level. |
| spells\_us.txt | An ASCII export of your [spells\_new](https://eqemu.gitbook.io/database-schema/categories/client-files/spells_new) table, which contains information about all of the [spells](../spells/) in use in your world. |

{% hint style="info" %}
Be sure to read out [Client Spell Limitations](../spells/client-spell-id-limitations.md) before exporting your client files!
{% endhint %}

### Exporting for the First Time Requires an Extra Step

Since it's your first time, be sure to go slow and use lots of ...care.  You probably do not have a folder that the export script will look for; this folder is where the resultant files will be placed after the export script is run.  You will need to create this folder in the root directory of your EQEmu Server Install.

{% hint style="info" %}
You only need to do this the first time--afterwards, you can skip to the section below and Run the Export Script.
{% endhint %}

**First, navigate to your root server folder:**

```text
cd /home/eqemu/server
```

This example uses the default install directory; obviously you would navigate to your server directory if you chose to install the server in a different location.

**Then, create the necessary directory:**

```text
mkdir export
```

### Run the Export Script

In your main server folder, you will see a symlink to the client file exporter script.  Run this script to create your files.  

{% hint style="warning" %}
When you run this script, it will happily overwrite the contents of your server/export directory.  If you want to keep old copies of your Client Files, you should back them up before executing.
{% endhint %}

```text
./export_client_files
```

If you receive an error that the folder doesn't exist, please go back to the "Exporting for the First Time..." step above.

If you were successful, navigate to your _export_ folder and you should see four files listed:

```text
[root@localhost server]# cd export
[root@localhost export]# ls
BaseData.txt  dbstr_us.txt  SkillCaps.txt  spells_us.txt
```

You can now distribute these files to your players so that their client will match your EQEmu World Server settings.

