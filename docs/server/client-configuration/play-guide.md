# Play Guide

!!! info
  
    This guide will take you through the steps necessary to play on an EQEmulator Server. <a href="ac_this-guide-will-take-you-through-the-steps-necessary-to-play-on-an-eqemulator-server" id="ac_this-guide-will-take-you-through-the-steps-necessary-to-play-on-an-eqemulator-server"></a>

## Supported Clients

!!! warning
    
    **EverQuest clients are the intellectual property of the Daybreak Game Company, LLC.** Copies of Daybreak Game Company intellectual property are not sourced through the EQEmulator project, nor should you provide unlicensed copies through any EQEmulator asset.


|Titanium Edition|Secrets of Faydwer|Seeds of Destruction|Underfoot|Rain of Fear (Most used)|
|:---:|:---:|:---:|:---:|:---:|
| ![](https://user-images.githubusercontent.com/3319450/143334304-4faf5cf8-6ed9-4b47-a0e2-938cc3f68e57.png){: style="height:150px; width: 125px;"} | ![image](https://user-images.githubusercontent.com/3319450/143334432-e6e9eaef-b141-4b05-9607-ceb38dcf717d.png){: style="height:150px; width: 125px;"} | ![image](https://user-images.githubusercontent.com/3319450/143334455-420ee97d-ed5e-4f21-a824-48371831c604.png){: style="height:150px; width: 125px;"} | ![image](https://user-images.githubusercontent.com/3319450/143334476-4b699dec-6a1b-4690-be7f-64eec22cd60c.png){: style="height:150px; width: 125px;"} | ![image](https://user-images.githubusercontent.com/3319450/143334498-810f76b6-7f18-4723-a02a-d50e11af96d1.png){: style="height:150px; width: 250px;"} |

- [x] **Everquest** Titanium Edition (Retail from the CDs)
- [x] **Everquest** Secrets of Faydwer (Retail from the CDs/DVDs)
- [x] **Everquest** Seeds of Destruction (No longer available by legal means)
- [x] **Everquest** Underfoot (No longer available by legal means)
- [x] **Everquest** Rain of Fear (RoF - Build Date of Dec 10 2012 17:35:44) Download-Only from Steam (No longer available by legal means)
- [x] **Everquest** Rain of Fear (RoF2 - Build Date of May 10 2013 23:30:08) Download-Only from Steam (No longer available by legal means)

!!! note
  
    Some servers may only allow earlier version of clients (Left earliest, right latest)

!!! info
  
    Most servers will allow most clients simultaneously, the earlier clients may not have the features or zone capabilities that later clients do, however.

## Step 1) Installing 

### Titanium & Secrets of Faydwer

!!! info
    
    Install Titanium or Secrets of Faydwer from the disks or Seeds of Destruction/Underfoot from your local copy the Steam download. **Do NOT patch live**.

    Titanium/SoF Note Some users with AMD Athlon(tm) 64 X2 Dual Core Processors have reported a speed issue with the client.

    If you still have Dual Core CPU issues or other issues with starting Everquest, see the[ Frequently Asked Questions](frequently-asked-questions.md). Note that the SoD and later clients do not have Multi-CPU issues like previous clients.

### Seeds of Destruction, Underfoot, Rain of Fear

!!! warning

    It is HIGHLY recommended that you make a backup copy of your SoD/UF/RoF/RoF2 download folder before taking any further actions in the case that your download gets corrupted or patched by mistake. 

    Do not attempt to run Everquest using the shortcut on your Desktop before making a backup or it will patch your client and break it for EQEmu!

    To make a backup, simply copy the installed Everquest folder to some other place on your hard drive. Naming it something like C:\Everquest_SoD or C:\Everquest_RoF2 would make it easy to identify.

!!! note

    If using the SoD, UF, RoF, or RoF2 client from Steam, it will get installed into one of the following folders by default:

    **On 32bit Windows:**

    `C:\Program Files\Steam\steamapps\common\EverQuest`

    **Or for 64bit Vista/7:**

    `C:\Program Files (x86)\Steam\steamapps\common\EverQuest`

## Step 2) Forum Account

If you have not already done so, you will need to register an account with the EQEmulator Forums∞. Click the "Register" link at the top left of the page. Follow the prompts that are displayed and enter the required information: username, password, email address, etc. Make sure you use a valid email address as your new account needs to be verified. Check your email and validate your account. If you do not receive a validation code, please PM an Administrator in the IRC channel.

## Step 3) Login Server Account

You will now need to create a **Login Server account** that you will use log into the EQEmulator Loginserver.

Please go to the Login Server page (located under Miscellaneous in your UserCP).

[Select "Create a LS Account."](http://www.eqemulator.org/account/?CreateLS)

Fill out the fields and proceed. 

!!! info

    Please note that your EQEmulator Forum account and EQEmulator Login Server account are totally seperate; you should use a different username/password combination for each.

    Your Login Server Account passwords are now recoverable as of March 2014, but you must setup verification right away.

    [Verification](http://www.eqemulator.org/account/?SMS)

## Step 4) EQHost File

!!! info

    In order for your client to connect to the **EQEmulator Login Server** (and not SOE's EQLive Login Server), you will need to change your **eqhost.txt** file to point at the correct location. 

    Go into your EverQuest directory and locate the eqhost.txt file. 

    Replace the contents of the file with the following (Remove any trailing spaces)

!!! eqhost

    As of November 2021; most of the community has been going through ProjectEQ Loginserver for stability reasons even when your Loginserver accounts are EQEmulator accounts
  
    === "Titanium & Secrets of Faydwer"

    ```text
    [LoginServer]
    # Host=login.eqemulator.net:5998
    Host=login.projecteq.net:5998
    ```

    === "Seeds of Destruction, Underfoot, Rain of Fear"

    ```text
    [LoginServer]
    # Host=login.eqemulator.net:5999
    Host=login.projecteq.net:5999
    ```

## Step 5) Creating Shortcut

There are the generic instructions for configuring your client for connection to an EQEmu Server.

???+ warning

    Do not launch your EverQuest client prior to making these adjustments, as it will attempt to patch and you will have to reinstall the client from scratch!

Navigate to your EverQuest client directory and right-click on the eqgame.exe application.

!!! Shortcut

    Choose the **Send To** option, and choose **Desktop (create shortcut)**.

    ![image](../../gitbook/assets/something.jpg){: style="height:300px; width: auto;"}

!!! Properties

    Right-click on the shortcut you created on your desktop, and select the _**Properties**_ option.

    ![](../../gitbook/assets/screenshot_10_2_19__9_39_am.png){: style="height:300px; width: auto;"}

    In the **Target** field, add ` patchme`, to the end of the line

    **Examples**

    `"C:\Program Files\Everquest\eqgame.exe" patchme`

    `C:\Everquest\eqgame.exe patchme`

## Step 6) Playing

That's it! You're ready to play now. Double-click that shortcut you just created, and login using the Login Server Account that you created in Step 3.


!!! info 

    If there are any issues or questions after completing the steps in this guide, please review the Troubleshooting FAQ below and the [Frequently Asked Questions](frequently-asked-questions.md) page. You should also search the Forums for answers before posting any questions. Another option for support is to join the EQEmulator IRC Support Channel and ask a question, and then stay logged in there and wait for a response (which can take a while).

## Community

Find us on Discord!

<iframe src="https://discord.com/widget?id=212663220849213441&theme=dark" width="300" height="500" allowtransparency="true" frameborder="0" sandbox="allow-popups allow-popups-to-escape-sandbox allow-same-origin allow-scripts"></iframe>


## FAQs

!!! question

    ### Can I use the anniversary edition or other expansions?

    No. As the it says above, you can only use Titanium, Secrets of Faydwer, Seeds of Destruction, Underfoot, or Rain of Fear at this time. 

    Other versions like platinum, anniversary, and trilogy **will not work**. 

    EQEmu compatibility is created for very specific versions of the eqgame.exe file.

    Every time SOE releases a patch, it is a different version.

!!! question

    ### Can I use third party software like MacroQuest?

    Check the rules for the server you play on.

!!! question

    ### Can I use custom UIs?

    Yes! Most custom UIs work with EMU, however there may be a few flaws, where as live uses corrupt resist and the rest timer where EMU doesn't, which would show "unknown" in your skin where it doesn't exist yet in EMU.

!!! question

    ### Is EverQuest Emulator Free?

    100% free to play, however you have to acquire a compatible version of Everquest by yourself.

!!! question

    ### Can I download EverQuest Titanium or Secrets of Faydwer?

    No, only the retail installations of either of the supported versions will work. Illegal downloads of them are not supported by the team here and are a bannable offense if discussed on the forums.

!!! question

    ### Is there a list of servers I can see and if they are legit, semi-legit, or non-legit and other stuff?

    Yes see [this](http://www.eqemulator.org/index.php?pageid=serverlist) page for EQEmulator Loginserver 

    See http://login.projecteq.net/ for the ProjectEQ Loginserver
