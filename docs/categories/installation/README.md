---
description: This page describes installing EQEmu Server
---

# Server Installation

### So you want to be a server operator?

Welcome to the world of EQEmulator. Chances are that you're wondering what it would be like to craft your own world. Perhaps you don't like how someone runs the world they have created, perhaps you would like to customize a world because nothing like \[insert your idea\] exists. Maybe you just want to run around a world like some kind of god and slay all of those nasty beasties that had the upper hand when you played EverQuest in 1999. There's really no right or wrong reason to create a world--but certainly there are ways that it can go horribly wrong, and your decisions will likely be the cause.

This guide is nothing more than the beginning of the journey to creating your own world. This guide is an attempt to impart some wisdom on those willing to receive it, and is written in the hope that you will read it carefully, thoroughly, and consider things carefully before charging into the unknown and wasting hours of your time by making a bad decision.

### What is a World Server?

It's probably a good idea to take a step back and make a few observations before we get into some deeper concepts. A world server, in the simplest terms, is an application that runs on a computer, and communicates with a database of information. 

When Neo joins Morpheus in the Construct Program in the movie _The Matrix_, their conversation parallels much of what you're hoping to create: you have a program \(EQEmu server\) and you can load anything you want \(from the database\). To extend the analogy, the world server, like The Matrix, has "rules" that determine how everything works. To paraphrase Morpheus, some rules on your World Server can be bent, and some rules can be broken.

* Installing using scripted installation:  
  * [Basic Server Install - Linux](server-installation-linux.md)
  * [Basic Server Install - Windows](server-installation-windows.md)

### About the Server Application.

All of the code that you see for EQEmu Server Github is the what makes the world work--it's The Matrix. When you run the installer, you get a copy of all of that code, and the installation script makes sure it's all compiled for you. If you want to break some rules, there's a really good chance that you're going to have to learn how to change the code, and then you can compile your changes, and create your own custom application with your changes.

* For Linux Users who want to break rules:  check out the [Ground Up Linux Install](ground-up-linux-install.md) guide.
* For Windows Users who want to break rules:  check out the [Ground Up Windows Install](ground-up-windows-install.md) guide.

### About the Database.

If you want shelves full of weapons to storm the government building where they're holding Morpheus captive, you'll need a database that contains that information. When you run the installer, you get a copy of the PEQ database. The PEQ database is the database created, modified and maintained by the folks at [ProjectEQ.net](https://projecteq.net/). Thousands and thousands of hours have likely been spent on the PEQ Database, and the people that run the project were incredibly generous and made their database open source so that anyone could make use of it. 

{% hint style="success" %}
Daily database dumps of the PEQ Database are [available here](http://db.projecteq.net/)
{% endhint %}

The PEQ database may not provide the exact data that makes **your world** exactly the way **YOU** want it, but they have saved you countless hours of creating [zones](../zones/), [NPCs](../npc/), [pathing grids](../npc/spawns/grids.md), and more. This is something to keep in mind before you spend your time complaining about the database content--they give freely, and at least to some degree, you have benefitted from their hard work. If you want to customize your world, you'll have to do some hard work too; and once you've done this work, please consider sharing it with the community, as so many others have done--someone **just like you** will benefit.

### Making the World Server talk to the Database

Consider this your first step in customizing your world once you've run the installer script.  You will need to [configure your eqemu\_config.json](configure-your-eqemu_config.md) file to set your world's name, the name of your database, and more.  Follow the instructions carefully so that your decisions do not become the cause of something going horribly wrong.

{% hint style="warning" %}
Be sure to back up any files before you start to modify them!  
{% endhint %}

### Where we go from there, is a choice I leave to you

| Choice | Description |
| :--- | :--- |
| [Operation](https://eqemu.gitbook.io/server/categories/operation) | Learn the ins-and-outs of operating your new world |
| [PEQ DB Editor](https://eqemu.gitbook.io/server/categories/installation/install-peq-database-editor) | How to install the PEQ DB editor \(so you can customize your world\). |
| [Scripting](https://eqemu.gitbook.io/server/categories/scripting) | Learn to customize NPC interactions through scripts |
| [Client Config](../client-configuration/) | You'll have to configure your client to connect to your world |
| [Custom NPCs](../npc/customizing-npcs/) | Learn to create a custom NPC |
| [Help!](https://discord.gg/QHsm7CD) | Join the discussion on Discord--look for **\#support server** |



