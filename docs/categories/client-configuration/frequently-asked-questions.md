# Frequently Asked Questions

## EQEmulator Support

**What is the Everquest Emulator?**

* EQEmu is a server program to allow Admins to run their own Private or Public servers.

**Can I play EQEmu without setting up a server?**

* Yes. If you only want to play the game, you don't have to run a server.

**Are there any costs to play the Everquest Emulator?**

* No, there are no costs to play.

**What is required to play on an Everquest Emulator Server?**

* The only requirement is a fresh install of Everquest Titanium \(disks\), Secrets of Faydwer \(disks\), Seeds of Destruction \(Steam\), Underfoot \(Steam\), or Rain of Fear \(Steam\) unpatched.

**Can I use the separate Expansion CDs if I have all of the Expansions included in Titanium?**

* No, only Everquest Titanium Pack \(Discs\) or Secrets of Faydwer \(Discs\) will work for the emulator.  The rest of the compatible clients are all Digital Downloads from Steam. Nothing else will work.

**Is there any place I can download an Everquest Cient for EQEmulator?**

* No, not legally. You may only acquire Titanium and Secrets of Faydwer.

**Is the Emulator buggy?**

* There are some bugs in the emulator, but nothing major enough to keep people from having fun. New updates/fixes are added often.

**Is the Emulator Similar to EQLive?**

* There are many servers and each one is different, but some do simulate live closely up through DoN.

**Are there servers with special or custom setups?**

* Yes, there are many servers that have increased experience rates, custom loot, quest, zones, mobs, encounters, etc.

**How many servers are there for the emulator?**

* There are dozens of different public servers and also a few private servers as well.

**Is there a way to tell which servers are the most popular?**

* Yes, on the Login Server it lists all servers and each server shows the current population. You can judge popularity pretty well by what the population is.  You can also view the Server List on the EQEmulator website for details.

**Where can I find more information about the Everquest Emulator?**

* You can find most of the information you need by reading through the Wiki. Another good place to find important information is in this list of useful resource Links:
* [http://www.eqemulator.net/forums/showthread.php?t=26075](http://www.eqemulator.net/forums/showthread.php?t=26075)

## Player Support

**I am running Windows Vista/7 and am having problems saving my eqhost.txt file after editing it as the Guide explained.**

* Due to the built-in UAC in Windows Vista & 7, you must either disable UAC or use the following work around to edit the eqhost.txt file:

1. Click Start -&gt; All Programs -&gt; Accessories
2. Right-click Notepad
3. Choose "Run as Administrator"
4. Click "Continue" at the UAC prompt
5. Browse to and modify your eqhost.txt file as desired

**When I try to open EQ, the window blinks black and then closes and I get an error, but EQ never even starts.**

* Most likely you are having a resolution or drivers issue. Try adjusting your EQ or Desktop Resolutions or updating your video drivers. It may just be that EQ is trying to start in a resolution that it does not support. If setting your desktop to 1024X768 does not work, then you may need to set it in your eqclient.ini. There should be a file named OptionsEditor.exe in your Everquest folder. Open that program and set your resolution to 1024X768 and see if that helps. If not, run it again and try other resolutions. Once you have it working, you should be able to adjust your resolution from in game and then /camp out to save the new settings.
* Another possible resolution to this issue is sometimes to try another video card driver for your video card. You may want to try using an older driver or newer one to see if that makes a difference. You may need to run a driver cleaner application when installing your new \(or old\) video card drivers, depending on the instructions for changing drivers on your specific video card.
* If none of the above resolves your resolution, crash, or lockup problems, try the suggestions below:
  * Right click on your Everquest icon and go under compatibility mode. There is a checkbox for "runs in 640x480 resolution". Click this. It does not mean the game will run in 640x480 but the intro screen will. Without that one checked, you may not even be able to get into the game.
  * Now, change these under your video mode. If EQ starts in windowed mode, or if the character select is in windowed mode do this; center the windowed mode window up on your screen, to where it covers the entire screen. Now Hit ALT-ENTER and you will go full screen before you go into the game. Only then enter the world, not in windowed mode. The reason for this is if you go full screen it sometimes crashes and I would rather crash before my character enters the world.
  * Make a copy of your eqclient.ini then delete what you have under the video mode section and change it to this:

```text
[VideoMode]
Width=1920
Height=1200
FullscreenBitsPerPixel=32
FullscreenRefreshRate=60
WindowedWidth=1920
WindowedHeight=1200
```

* When attempting to run more than one instance of Everquest it tries to open and then gives an error.
* You don't have enough memory to start another instance of EQ. Try adjusting your EQ settings lower. Turning off Luclin models helps the most.

**I am running Windows Vista and cannot get EQ to open or I get disconnected often.**

* Try running in Compatibility mode. See the link below to directions on how to do that:
* [http://www.eqemulator.net/forums/showthread.php?t=25327](http://www.eqemulator.net/forums/showthread.php?t=25327)

**EQ Opens and starts to patch, but my login doesn't work.**

* You messed up on the install setup. You need to uninstall EQ and delete the directory and follow this guide exactly:
*  [Play Guide: Getting Started](play-guide.md)

**I can get to the login screen, but I get an error saying that the connection timed out when I try to log in.**

* Either the login server is down, or more likely you need to adjust your eqhost.txt file as described in the play guide. Make certain that there are no extra spaces or carriage returns in the file. Also make sure you didn't accidentally save the file as eqhost.txt.txt or it won't work at all. Also, if you are trying to use EQEmu with an EQLive \(Patched\) installation instead of directly from the Titanium or Secrets of Faydwer Disks without patching, you will get a connection timed out error as well. In this case, though, it takes considerably longer before the error pops up after trying to log in. When the problem is with your network connection, the LS, or your eqhost.txt file, it should only take about 30 seconds or less to return with the error. But, if you are running a patched version of EQ, it can take minutes before the error pops up, if it ever does at all.

**I can get to the login screen, but I get an error about my username or password being wrong.**

* Either you didn't follow the play guide and setup your login server accounts properly, or your password is too long. Try using a password that is 6 characters long.

**I lost my Login Server and/or Forum Account username and/or password. What can I do about that?**

* Recovering usernames and/or passwords is covered in the LoginInformationRecovery page.

**I can log into the Server Select screen, but when I get there, no servers are showing up. It is just a blank list.**

* This usually means there is a network problem. Either your network is blocking port 5998 or your router may just be having issues that need to be corrected. If you live in a dorm or are playing from a public network, it is probably going through a proxy and that means they are most likely blocking port 5998, so you won't be able to play. If you are playing from home and have control of your internet connection, then the problem is probably either with your router/modem or with your PC. First, make sure that you don't have a firewall or any other security programs running that might block your connectivity in any way. Next, make sure you aren't forwarding port 5998 to any PCs on your network \(this includes port range forwarding that would include port 5998 in the range\). If not, next you want to try powering down your modem and router. Then power up your modem and wait until it gets a connection. Next, power up your router and wait for it to connect. Last, reboot your PC and try connecting to EQEmu again. If that still isn't working, then there is a good chance that you need to reset your router or router/modem back to factory defaults. Most modems have a sunken in button that you can hold down to reset it to defaults. Check your product documentation for how to do it. Once that is done, set up your router again and try connecting to EQEmu again. It should be working now, but if not, there may be an issue with the routing from your ISP or your router may just be going bad.

**I can log into the Server Select screen, but when I try to connect, the screen goes black and then goes to the login screen again.**

* There are multiple reasons that could be causing this. Here are some of them:

1. Most likely you are running the wrong version of Everquest. Only Legal copies of Everquest completely unpatched will work. See our list of supported ClientVersions to validate your client should be compatible.
2. You may want to try other servers to make sure that it isn't a problem on that particular server. Some servers only support older clients, but Titanium should work for all servers.
3. If you live in a dorm or are playing at work, you might be using a proxy that is blocking port 5998, which means you won't be able to connect.
4. If you are running a P2P or Torrent program, it might be using all of your bandwidth. Try turning them all off.
5. You may have a firewall that is blocking EQ. Disable all firewalls.
6. If you are running a Dual Core CPU or Windows Vista, you will probably need to read the following post and apply at least one of the fixes listed in it: http://www.eqemulator.net/forums/showthread.php?t=25327∞
7. You may need to adjust your Windows Desktop and/or your Everquest Resolution settings.
8. Make sure your drivers are up to date for your Video Card.
9. If you have sound enabled, make sure your sound card is working.
10. Try resetting your router. If that doesn't help, try connecting your PC directly to your cable/dsl modem and see if that helps.
11. If you previously had another version of EQ installed and installed SoF over the old install \(even if you uninstalled EQ before-hand\), you will need to uninstall EQ again and delete the Everquest install folder completely, and then install SoF again from scratch. Otherwise, simply installing SoF to a new/different folder than your current EQ install works just fine.

**I am having weird choppiness/warping when I play and/or getting disconnected very often.**

* Most likely you have a Dual or Quad Core CPU. You will want to read the following post for some suggestions:
* [http://www.eqemulator.net/forums/showthread.php?t=25327](http://www.eqemulator.net/forums/showthread.php?t=25327)

**I tried making a character, but when I was done, it just disconnected me.**

* This is because there is a time limit at character select/creation. Try being faster when making your character.

**Sometimes when I play my ping times are very high and/or I lag really bad.**

* Either you have a slow connection, or are downloading large files, or the server may be having resource issues. This is more likely when many players are on the server.

**I found a bug on the server that I play on. What should I do?**

* Most likely the bug you found is specific to that server. If so, post it on that server's forums or /petition a GM about it.

**Often when I am playing 2 or more characters at once, some of them will get disconnected almost every time I zone.**

* This is because you shouldn't play more than 1 character per account at the same time. You can easily make multiple accounts.

**The server I play on Limits the number of players allowed per-IP. More than 1 person plays from my IP. How can we get it increased?**

* EQEmu has no control over the IP limit settings of individual servers. Please make the IP limit increase request to the GM\(s\)/Admin\(s\) of the server you play on. Check their website/forums and post there with your request if possible. To find their website, please either check the MOTD on the server, ask in /ooc, or refer to the server list here and click the "view" link next to the server you play on:
* [http://www.eqemulator.org/index.php?pageid=serverlist](http://www.eqemulator.org/index.php?pageid=serverlist)
* If the server has a website, it is most likely on that page.

**Everquest works fine, but all of the fonts are hard to read due to being a bit blocky or garbled. How do I correct this so fonts look normal?**

* If your fonts are showing up oddly, you are probably running an LCD monitor and not running EQ at your monitor's native resolution. Try running in Windowed mode and see if that corrects the font issues \(ALT+ENTER\). If that fixes it, then just make sure when you are running in full screen mode, you are running at the native resolution of your monitor.

**I run Windows Vista/7 and am unable to type certain characters in the chat window such as: + = \[ \] { } " '. How is this fixed?**

1. Open Device Manager
2. Click on "Keyboards" to open the tree.
3. Right-Click on "Microsoft eHome MCIR 109 Keyboard" and choose "Uninstall", then click "Ok".
4. Right-Click on "Microsoft eHome MCIR Keyboard" and choose "Uninstall", then click "Ok".
5. Right-Click on "Microsoft eHome Remote Control Keyboard keys" and choose "Uninstall", then click "Ok".
6. Click on "Human Interface Devices" to open the tree.
7. Right-Click on "Microsoft eHome Infrared Transceiver" and choose "Disable", then click "Ok".
8. Restart Everquest.

**I am still having a problem with playing that wasn't answered here. What should I do?**

* First try reading and searching the forums. If you still can't find the solution, you should make a post in one of the support threads with as many details as possible about your issue. Make sure to post in the right section that is related to your issue \(General Support, Spell Support, Windows Server Support, Linux Server Support or Mini Login Support\).

## Server Administrator Support

**I setup the server, but when I click the start.bat shortcut, the window opens and closes so fast that I can't even read it.**

* You should try running the world.exe from the command prompt and watch for errors from there.

**I ran the world.exe from the command prompt and it stops after giving an error about being unable to connect to the database.**

* Make sure you have the config file setup properly. If you are certain that it is set correctly, then make sure you ran the following command from the mysql prompt in your database and set your new password to what you want it to be:
* set password for 'root'@'localhost' = OLD\_PASSWORD\('newpwd'\);

**My server loads up and everything looks ok accept I can't cast any detrimental spells and mobs won't aggro even if they are KoS.**

* This is because you are missing the .map file. These files need to be put in your eqemu/maps directory. They tell the server the 3D layout of the zone and they are not the same as the map files in your Everquest directory for in game maps. 

**My server loads up and I can log in, but I am unable to move or barely at all unless I have levitate on.**

* Make sure you have all of the required SQL added to your server. Check through the change logs for the required SQL.

**My server loads up and I can log in from my LAN, but everyone outside of my network gets disconnected when they try to connect to the server.**

* Make sure you have Port RANGE Forwarding setup to forward ports 7000 to 7500, 9000 to 9000 and optionally 9080 to 9080 to your server's LAN IP Address.

**I edited items, or factions and they don't seem to be showing up or working properly.**

* These changes only take effect after the server is restarted. Some other changes require zones to be reset to take effect.

**I have heard that Linux is better than Windows for running Emulator servers. Is it recommended for everyone to use?**

* If you are a technical person, then yes, you would probably benefit from running your server on Linux. If you are not technical, then you probably want to stick with Windows. The main benefit of Linux is greater stability and less system resource usage.

**I run a Windows server and my Logs directory is huge. Is there a way to turn off logging?**

* You can simply rename the eqemu/logs directory to stop the logs, or you can edit the log.ini file by following the wiki for it.

**I am having a problem with some quests. If I turn in an item, the NPC doesn't react at all and just eats the item.**

* Make sure you have moved the Plugins files from your Quests directory into the eqemu/plugins folder.

**My server was running fine before I updated to the latest Source/Binaries and now something is not working properly.**

* Make sure that you added all required SQL files as noted in the changelog.txt file. Check the SVN folder here for the full list:
* http://code.google.com/p/projecteqemu/source/browse/\#svn/trunk/EQEmuServer/utils/sql/svn∞
* You will want to source in any updates that occurred between the revision your server was last running and the revision it is running now. For example; If your server was running Rev1143 and you upgraded it to Rev1200, you will need to source the 1195\_account\_suspendeduntil.sql into your db and can optionally source 1144\_optional\_rule\_return\_nodrop.sql as well.

**My EQEmu Web Tool gives an error when trying to connect to it and does not prompt for login.**

* First, make sure you have the HTTP service enabled in your eqemu\_config.xml
* If it is enabled and still not functioning, make sure you have the mime.types file in your server folder.
* Also, make sure you have port 9080 open in your router/firewall and/or forwarded.
* Last, make sure your EQEmu server is started, or the tool will not function.

**My EQEmu Web Tool gives a login prompt, but I am not able to authenticate.**

* First, you must be logging into the Web Tool on a GM account \(using your normal Login Server account name\), as normal accounts will not work.

Second, you must set a password in your account table password field for the account you want access on. You can hash the password so it is encrypted.

More to be added as needed.

