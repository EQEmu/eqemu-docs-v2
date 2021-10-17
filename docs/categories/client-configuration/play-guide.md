# Play Guide

## Description

#### This guide will take you through the steps necessary to play on an EQEmulator Server. <a href="ac_this-guide-will-take-you-through-the-steps-necessary-to-play-on-an-eqemulator-server" id="ac_this-guide-will-take-you-through-the-steps-necessary-to-play-on-an-eqemulator-server"></a>

* **Please note: As of this time you MUST have one of the follow versions of Everquest**

| <p><img src="http://wiki.eqemulator.org/l/wa/images/exp_box_art/Titanium.jpg" alt=""></p><p><strong>Everquest: Titanium Edition (Retail from</strong></p><p><strong>the CDs)</strong></p>                                                         | <p><img src="http://wiki.eqemulator.org/l/wa/images/expansions/secrets_of_faydwer.png" alt=""></p><p><strong>Everquest: Secrets of Faydwer (Retail from the CDs/DVDs)</strong></p>                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <p><img src="http://wiki.eqemulator.org/l/wa/images/expansions/seeds_of_destruction.jpg" alt=""></p><p><strong>Everquest: Seeds of Destruction (No longer available by legal means)</strong></p>                                                  | <p><img src="http://wiki.eqemulator.org/l/wa/images/expansions/underfoot.jpg" alt=""></p><p><strong>Everquest: Underfoot (No longer available by legal means)</strong></p>                                                                                           |
| <p><img src="http://wiki.eqemulator.org/l/wa/images/expansions/eqroflogo.png" alt=""></p><p><strong>Everquest: Rain of Fear (RoF - Build Date of Dec 10 2012 17:35:44) Download-Only from Steam (No longer available by legal means)</strong></p> | <p><img src="http://wiki.eqemulator.org/l/wa/images/expansions/eqroflogo.png" alt=""></p><p><strong>Everquest: Rain of Fear (RoF2 - Build Date</strong> <strong>of May 10 2013 23:30:08) Download-Only from Steam  (No longer available by legal means)</strong></p> |

* Some servers may only allow earlier version of clients (Left earliest, right latest)
* Most servers will allow most clients simultaneously, the earlier clients may not have the features or zone capabilities that later clients do, however.

See the ClientFeatures page to view some of the differences between the clients. Also see the [Client Versions](../player/client-version-bitmasks.md) page for details on determining if the client you have is compatible with EQEmu or not.

* Note: If using the Underfoot client, some files may need to be copied from another EQ install as mentioned in the[ ](http://wiki.eqemulator.org/p?UFMissingFilesList\&frm=Frequently_Asked_Questions--Play_Guide%3A_Getting_Started)[Underfoot Missing Files](underfoot-missing-files.md) page and discussed [here](http://www.eqemulator.org/forums/showthread.php?t=31635)
* Note: Do not download the Rain of Fear client from anyplace other than Steam, or you may get an incompatible version of the client!
  * The Rain of Fear client is no longer available from Steam.

## Step 1 - Installing 

### **Tit | SoF**

* Install Titanium or Secrets of Faydwer from the disks or Seeds of Destruction/Underfoot from your local copy the Steam download. Do NOT patch to live. Do not patch at all.
* Titanium/SoF Note Some users with AMD Athlon(tm) 64 X2 Dual Core Processors have reported a speed issue with the client.
* If you still have Dual Core CPU issues or other issues with starting Everquest, see the[ Frequently Asked Questions](frequently-asked-questions.md). Note that the SoD and later clients do not have Multi-CPU issues like previous clients.

### **SoD | UF | RoF | RoF2**

* Note It is HIGHLY recommended that you make a backup copy of your SoD/UF/RoF/RoF2 download folder before taking any further actions in the case that your download gets corrupted or patched by mistake.  Do not attempt to run Everquest using the shortcut on your Desktop before making a backup or it will patch your client and break it for EQEmu!
* To make a backup, simply copy the installed Everquest folder to some other place on your hard drive. Naming it something like C:\Everquest_SoD or C:\Everquest_RoF2 would make it easy to identify.
* Note: If using the SoD, UF, RoF, or RoF2 client from Steam, it will get installed into one of the following folders by default:
* **On 32bit Windows:**
  * C:\Program Files\Steam\steamapps\common\EverQuest
* **Or for 64bit Vista/7:**
  * C:\Program Files (x86)\Steam\steamapps\common\EverQuest

## Step 2 - Forum Account

* If you have not already done so, you will need to register an account with the EQEmulator Forums∞. Click the "Register" link at the top left of the page. Follow the prompts that are displayed and enter the required information: username, password, email address, etc. Make sure you use a valid email address as your new account needs to be verified. Check your email and validate your account. If you do not receive a validation code, please PM an Administrator in the IRC channel.

## Step 3 - Login Server Account

* You will now need to create a **Login Server account** that you will use log into the EQEmulator Loginserver.
  * Please go to the Login Server page (located under Miscellaneous in your UserCP∞).
    * [Select "Create a LS Account."](http://www.eqemulator.org/account/?CreateLS)
    * Fill out the fields and proceed.
    * \***Please note that your EQEmulator Forum account and EQEmulator Login Server account are totally seperate; you should use a different username/password combination for each.**
* Your Login Server Account passwords are now recoverable as of March 2014, but you must setup verification right away.
  * [Verification](http://www.eqemulator.org/account/?SMS)

## Step 4 - EQHost File

* In order for your client to connect to the EQEmulator Login Server (and not SOE's EQLive Login Server), you will need to change your eqhost.txt file to point at the correct location. Go into your EverQuest directory and locate the eqhost.txt file. Replace the contents of the file with the following (Remove any trailing spaces):

### **Tit | SoF**

```
[LoginServer]
Host=login.eqemulator.net:5998
```

### **SoD| UF | RoF | RoF2**

```
[LoginServer]
Host=login.eqemulator.net:5999
```

## Step 5 - Creating Shortcut

* Now, you will need to create a shortcut to launch EverQuest and bypass the EQLive patcher (if this step is not taken, you will launch the patcher and overwrite the EQHost.txt file that you just modified, making you connect to the global EQLive Login Server). To do this, go into your EverQuest directory. Locate the **EQGAME.EXE** file. Right-click on the file and select "Create Shortcut." You can rename this shortcut and drag it anywhere you'd like (most people like to put it on their desktop). Now, right-click that shortcut and select Properties. In the target box, you will see the path to the eqgame.exe file. We need to edit this so that it will pass an argument to that executable when the shortcut is launched. At the end of the text box (after any parentheses), add "patchme" (no quotes). Your target box should look similar to this:
* "C:\Program Files\Everquest\eqgame.exe" patchme
* Or if you do not have spaces in the path, it will look like this without quotes:
* C:\Everquest\eqgame.exe patchme

## **Step 6 - Playing**

* That's it! You're ready to play now. Double-click that shortcut you just created, and login using the Login Server Account that you created in Step 3.
* If there are any issues or questions after completing the steps in this guide, please review the Troubleshooting FAQ below and the [Frequently Asked Questions](frequently-asked-questions.md) page. You should also search the Forums for answers before posting any questions. Another option for support is to join the EQEmulator IRC Support Channel and ask a question, and then stay logged in there and wait for a response (which can take a while).

### Quick FAQs

* **Can I use the anniversary edition or other expansions?**
  * NO. As the it says above, you can only use Titanium, Secrets of Faydwer, Seeds of Destruction, Underfoot, or Rain of Fear at this time. Other versions like platinum, anniversary, and trilogy WILL NOT WORK. EQEmu compatibility is created for very specific versions of the eqgame.exe file.  Every time SOE releases a patch, it is a different version.  See the[ Client_Versions](http://wiki.eqemulator.org/p?Client_Versions\&frm=Frequently_Asked_Questions--Play_Guide%3A_Getting_Started) page for specific version details.
* **Can I use third party software like MacroQuest and Magelo?**
  * Check the rules for the server you play on. Why people would lie about this is beyond me.
* **Can I use custom UIs? **
  * YES! Most custom UIs work with EMU, there may be a few flaws, where as live uses corrupt resist and the rest timer where EMU doesn't, which would show "unknown" in your skin where it doesn't exist yet in EMU.
* **Is EMU free?**
  * 100% free to play, however you have to acquire a compatible version of Everquest by yourself.
* **Can I download EverQuest Titanium or Secrets of Faydwer?**
  * No, only the retail installations of either of the supported versions will work. Illegal downloads of them are not supported by the team here and are a bannable offense if discussed on the forums.
* **Is there a list of servers I can see and if they are legit, semi-legit, or non-legit and other stuff?**
  * Yes see [this](http://www.eqemulator.org/index.php?pageid=serverlist) page.
