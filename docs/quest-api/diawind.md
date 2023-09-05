# DiaWind (Dialogue Window)

Dialogue Window used to be a plugin that is now source-native that provides a simple way to deliver a dialogue experience to players through a popup window instead of through the chat window.

## Note

* Using newlines may cause the dialogue window to not function properly. Example below.

### No Newlines
```pl
sub EVENT_SAY {
    if ($text=~/#a/i) {
        my $dialogMessage = "{title: Title} {button_one: Button One} {button_two: Button Two} wintype:1 Message";
        quest::crosszonedialoguewindowbycharid($client->CharacterID(), $dialogMessage);
    }
}
```
![image](https://github.com/EQEmu/eqemu-docs-v2/assets/89047260/b2718cdd-dfc8-410e-bccb-b502bd727d63)

### Newlines
```pl
sub EVENT_SAY {
    if ($text=~/#a/i) {
        my $dialogMessage = "{title: Title}
        {button_one: Button One}
        {button_two: Button Two}
        wintype:1
        Message";
        quest::crosszonedialoguewindowbycharid($client->CharacterID(), $dialogMessage);
    }
}
```
![image](https://github.com/EQEmu/eqemu-docs-v2/assets/89047260/d319b86b-31b9-475c-b522-30463e080489)


## Features

* Simple custom markdown with features to change window behavior
* Simple html color code support
* NPC animation support
* Window timer support
* Saylink / Bracket support response through buttons
* If there are quest saylinks in the quest::say code, the dialogue window render will strip them out so they display normally in the window.
* If there are multiple saylink responses, they are rendered in the chat window per normal window implementation. Single responses are at the click of a button
* If there is no animation code set in the dialogue window the NPC will animate socially with a handful of different "greet" animations. This can also be shut off behind rule mentioned above

## Enabling

You can enable this feature globally and have it intercept all normal quest dialogue through middleware functionality by enabling **rules**

```
RULE_BOOL(Chat, QuestDialogueUsesDialogueWindow, false, "Pipes all quest dialogue to dialogue window")
```

**QuestDialogueUsesDialogueWindow** will need to be set to true

## Examples

Below is when **QuestDialogueUsesDialogueWindow** is turned on, normal `quest::say` dialogue gets piped to a window.

![image](https://user-images.githubusercontent.com/3319450/132463174-b1156824-b5c1-4acb-8d75-7061d5cc334d.gif)

### Multiple Buttons

![image](https://user-images.githubusercontent.com/3319450/132143042-0b3b1711-988b-40fb-a9aa-192b496c503d.png)

```perl
sub EVENT_SAY {
    $client->DiaWind("
        {title: This is my amazing window title!}
        {button_one: This is one!}
        {button_two: This is two!}
        wintype:1
        Hello!
    ");
}
```

![image](https://user-images.githubusercontent.com/3319450/137598461-4f97855d-b3d2-463d-addc-0b0061eaa744.png)

**Quest API Methods**

```perl
$client->DialogueWindow("markdown");
$client->DiaWind("markdown"); // alias
```

```lua
e.other:DialogueWindow("markdown");
e.other:DiaWind("markdown"); // alias
```

## Markdown

* **{lb}** = Light Blue Color
* **{y}** = Yellow Color
* **{gold}** = Gold Color
* **{g}** = Green Color
* **{r}** = Red Color
* **{gray}** = Gray Color
* **~** = End Color Tag ()
* **[&gt;** = Response Text (Not Visible) - This is what the player responds with, for example: [What else do we need to do?&gt;
* **[]** = Response Text (Visible) - If a NPC has multiple brackets, it will give the player multiple saylinks to click, if there is just one [bracket] inline with the text, the player will respond by clicking 'Yes' on the window
* **+66+** = Animation number - As long as there is a number between two plus signs together, the NPC will perform that animation
* **+salute+** = Animation phrase - This references
* **{bullet}** - Equivalent to a bullet such as the one in this list
* **{in}** - Will indent the text
* **{linebreak}** - Will create a linebreak
* **mysterious** - If the text 'mysterious' is anywhere in the text, it will not show up in the player window, but it will format the window as 'Mysterious Voice tells you'
* **wintype:0/1** - If this is 1 the window will be Yes/No
* **popupid:ID** - If this option is present, it will give the popup window an ID response on "Yes" or button one
* **secondresponseid:id** - If this option is present, it will give the popup window an ID response on "No" or button two
* **noquotes** - This tells the window to not use any quotes for special formatting at times
* **nosound** - No sound affect will play with the window
* **=Timer=** - If a number is specified between the two ==, it will countdown a timer on the popup window

