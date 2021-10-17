# Windows

{% hint style="info" %}
Be sure to adjust your eqhost file as [indicated for all clients](./).
{% endhint %}

### Generic Connection Instructions

The are the generic instructions for configuring your client for connection to an EQEmu Server.

{% hint style="warning" %}
Do not launch your EverQuest client prior to making these adjustments, as it will attempt to patch and you will have to reinstall the client from scratch!
{% endhint %}

Navigate to your EverQuest client directory and right-click on the eqgame.exe application.

Choose the _**Send To**_ option, and choose _**Desktop \(create shortcut\)**_.

![](../../.gitbook/assets/screenshot4%20%281%29.jpg)

Right-click on the shortcut you created on your desktop, and select the _**Properties**_ option.

![](../../.gitbook/assets/screenshot_10_2_19__9_39_am.png)

In the _**Target**_ field, add " patchme", without the quotation marks, to the end of the line:

1. If there are spaces in your file path:  `"C:\Program Files\Everquest\eqgame.exe" patchme` 
2. If there are no spaces in your file path:  
  
   `C:\Everquest\eqgame.exe patchme`

The final result can look like this:

![](../../.gitbook/assets/shortcut_properties.png)

### 

