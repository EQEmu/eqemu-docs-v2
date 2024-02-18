!!! info

    If you're new to programming or scripting, many applications use event-driven architectural patterns. 

    The EverQuest Emulator scripting engine is almost exclusively event driven and events are triggered through various player-triggered or environment-triggered mechanisms. 

!!! note 

    For the most up to date set of events see the **events** section in the left navigation pane

### Default Exports

* These exports are the default event objects exported along with these event types.

#### Bot
`e.self`

#### Encounter
`e.name`

#### Item
`e.self` and `e.owner`

#### NPC
`e.self`

#### Player
`e.self`

#### Spell
`e.self`

### EVENT_AGGRO

#### Trigger

* When a mob aggros a client.

Often used for flavor text--just remember that every NPC_Type has an EmoteID that can often handle this behavior.

#### Example

* In this example, the NPC will say "Time to die PlayerName." when the NPC is aggro'd by the player.

=== "Perl"

    ```perl
    sub EVENT_AGGRO {
        quest::say("Time to die!");
    }
    ```

=== "Lua"

    ```lua
    function event_aggro(e)
        e.self:Say("Time to die!");
    end
    ```

### EVENT_AGGRO_SAY

#### Trigger

* When a mob is targeted, the player types something, and NPC is in combat.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| data | int | `quest::say($data); # returns int` |
| text | int | `quest::say($text); # returns int` |
| langid | int | `quest::say($langid); # returns int` |

#### Example

* In this example, the NPC, if in combat, would say the names of everyone on its hate list, and include both the amount of damage the entity has done, as well as the amount of hate the entity has generated.

=== "Perl"

    ```perl
    sub EVENT_AGGRO_SAY {
        #:: Match "fight", case insensitive, if the NPC aggro
        if ($text=~/fight/i) {
            quest::say("I am fighting!");
        }
    }
    ```

=== "Lua"

    ```lua
    function event_aggro_say(e)
        if (e.message:findi("fight")) then
            e.self:Say("I am fighting!");
        end
    end
    ```

### EVENT_ATTACK

#### Trigger

* When the NPC is attacked.

Note the subtle difference from EVENT_AGGRO, which is triggered when the NPC is aggro'd (which could occur through bad faction, for instance).

#### Example

* In this example, the NPC will say "Time to die PlayerName." when the NPC is attacked by the player.

=== "Perl"
    ```perl
    sub EVENT_ATTACK {
        quest::say("Time to die $name.");
    }
    ```

=== "Lua"
    ```lua
    function event_attack(e)
        e.self:Say("Time to die!");
    end
    ```

### EVENT_AUGMENT_ITEM

#### Trigger

* When a client augments an item.

You would likely use this event in your global player.pl file.

#### Example

* In this example, a message in yellow text is displayed to the client when the player adds an augment to an item.


=== "Perl"
    ```perl
    sub EVENT_AUGMENT_ITEM{
        $client->Message(15, "Yay, it fit!");
    }
    ```

=== "Lua"
    ```lua
    function event_augment_item(e)
        e.owner:Message(15, "Yay, it fit!");
    end
    ```

### EVENT_AUGMENT_INSERT

#### Trigger

* When a client inserts an augment into an item.

You would likely use this event in your global player.pl file.

#### Example

* In this example, a message in yellow text is displayed to the client when the player puts the augment into the augment slot of an item.


=== "Perl"
    ```perl
    sub EVENT_AUGMENT_INSERT {
        $client->Message(15, "Yay, it fit!");
    }
    ```

=== "Lua"
    ```lua
    function event_augment_insert(e)
        e.owner:Message(15, "Yay, it fit!");
    end
    ```

### EVENT_AUGMENT_REMOVE

#### Trigger

* When a client removes an augment from an item.  You would likely use this event in your global player.pl file.

#### Example

* In this example, a message in yellow text is displayed to the client when the player removes an augment from an item.


=== "Perl"
    ```perl
    sub EVENT_AUGMENT_REMOVE {
        $client->Message(15, "Yay, you pulled it out!");
    }
    ```


=== "Lua"
    ```lua
    function event_augment_remove(e)
        e.owner:Message(15, "Yay, you pulled it out!");
    end
    ```



### EVENT_BOT_COMMAND

#### Trigger

* When a player says anything with the **^** command token preceding it that isn't handled by a system bot command.
* Exception: `^help` will invoke a special call so that scripted bot commands may be included in search listing.

Used in player files to handle script-based bot command processing.

#### Exports

| Name | Type | Details |
| :--- | :--- | :--- |
| bot_command | string | A single, de-tokenized bot command name |
| args (perl) | string | An unseparated text line of additional parameters |
| args (lua) | string array | An array of single-word text parameters |

#### Example

* These examples are a working template for adding your own bot commands.

Note: The scripting apis do not fully support the Bot object and scripting options will be limited.


=== "Perl"
    ```perl
    sub EVENT_BOT_COMMAND {
        
        my %bot_command_data = (
            "scriptbotcommand1" => [0, "perl script-based supplementation test bot command 1"],
            "scriptbotcommand2" => [100, "perl script-based supplementation test bot command 2"],
            "scriptbotcommand3" => [250, "perl script-based supplementation test bot command 3"]
        );
        
        if ($bot_command eq "help" || $bot_command eq "?") { # always a supplemental call from within the system bot_command_help handler
            my $available_bot_commands = 0;
            
            foreach my $bot_command_name (sort {$a cmp $b} keys %bot_command_data) {
                if ($status >= $bot_command_data{$bot_command_name}[0] && ($args eq "0" || $bot_command_name=~/$args/i)) {
                    $client->Message(15, "^$bot_command_name - $bot_command_data{$bot_command_name}[1]");
                    $available_bot_commands = $available_bot_commands + 1;
                }
            }
            
            return $available_bot_commands;
        }
        elsif (exists($bot_command_data{$bot_command}) && $status >= $bot_command_data{$bot_command}[0]) {
            if ($bot_command eq "scriptbotcommand1") {
                $client->Message(14, "doing '$bot_command'...");
                return 1;
            }
            elsif ($bot_command eq "scriptbotcommand2") {
                $client->Message(14, "doing '$bot_command'...");
                return 1;
            }
            elsif ($bot_command eq "scriptbotcommand3") {
                $client->Message(14, "doing '$bot_command'...");
                return 1;
            }
        }
        
        return 0;
    }
    ```


=== "Lua"
    ```lua
    function event_bot_command(e)
        
        local bot_command_data = {
            ["scriptbotcommand1"] = {0, "lua script-based supplementation test bot command 1"},
            ["scriptbotcommand2"] = {100, "lua script-based supplementation test bot command 2"},
            ["scriptbotcommand3"] = {250, "lua script-based supplementation test bot command 3"}
        }
        
        if (e.bot_command == "help" or e.bot_command == "?") then -- always a supplemental call from within the system bot_command_help handler
            local available_bot_commands = 0;
            
            local key_sort = {}
            for key in pairs(bot_command_data) do
                table.insert(key_sort, key);
            end
            table.sort(key_sort);
            
            for index, key in ipairs(key_sort) do
                if (e.self:Admin() >= bot_command_data[key][1] and (e.args[1] == nil or key:findi(e.args[1]))) then
                    e.self:Message(15, "^" .. key .. " - " .. bot_command_data[key][2]);
                    available_bot_commands = available_bot_commands + 1;
                end
            end
            
            return available_bot_commands;
        elseif (bot_command_data[e.bot_command] ~= nil and e.self:Admin() >= bot_command_data[e.bot_command][1]) then
            if (e.bot_command == "scriptbotcommand1") then
                e.self:Message(14, "doing '" .. e.bot_command .. "'...");
                return 1;
            elseif (e.bot_command == "scriptbotcommand2") then
                e.self:Message(14, "doing '" .. e.bot_command .. "'...");
                return 1;
            elseif (e.bot_command == "scriptbotcommand3") then
                e.self:Message(14, "doing '" .. e.bot_command .. "'...");
                return 1;
            end
        end
        
        return 0;
    end
    ```



### EVENT_CAST

#### Trigger

* When a client casts a spell.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| spell_id | int | `quest::say($spell_id); # returns int` |

#### Example

* In this example, the player would emote upon a successful cast.


=== "Perl"
    ```perl
    sub EVENT_CAST {
        quest::me("regains his concentration and casts his spell.");
    }
    ```


=== "Lua"
    ```lua
    function event_cast(e)
        e.self:Emote("regains his concentration and casts his spell.");
    end
    ```



### EVENT_CAST_BEGIN

#### Trigger

* When a client begins to cast a spell.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| spell_id | int | `quest::say($spell_id); # returns int` |

#### Example

* In this example, the player would emote if they begin casting the gate spell.  You would likely place this particular snippet into your global player.pl file.


=== "Perl"
    ```perl
    sub EVENT_CAST_BEGIN {
        #:: Match if $spell_id is 36 - Gate
        if ($spell_id == 36) {
            quest::me("begins casting the Gate spell.");
        }
    }
    ```


=== "Lua"
    ```lua
    function event_cast_begin(e)
        --:: Match if spell_id is 36 - Hate
        local spell_id = e.spell:GetID();
        if (spell_id == 36) then
            e.self:Emote("begins casting the Gate spell.");
        end
    end
    ```



### EVENT_CAST_ON

#### Trigger

* When a player casts a spell on a player or NPC.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| spell_id | int | `quest::say($spell_id); # returns int` |

#### Example

* In this example, if the player casts Banish Summoned on an NPC with a Summoned body type, the NPC will be killed. If you were placing this snippet in an NPC's quest script, you likely wouldn't bother matching body type.


=== "Perl"
    ```perl
    sub EVENT_CAST_ON {
        #:: Match if $spell_id is 116 - Banish Summoned
        if ($spell_id == 116) {
            #:: Match if the NPC's body type is 28 - Summoned Creature
            if ($mob->GetBodyType() == 28) {
                $npc->Kill();
            } 
            else {
                $client->Message(13, "This spell only effects summoned creatures");
            }
        }
    }
    ```


=== "Lua"
    ```lua
    function event_cast_on(e)
        --:: Match if spell_id is 116 - Banish Summoned
        local spell_id = e.spell:GetID();
        if (spell_id == 116) then
            --:: Match if the NPC's body type is 28 - Summoned Creature
            local body_type = e.self:GetBodyType();
            if (body_type == 28) then
                e.self:Kill();
            else
                if (e.self:IsClient()) then
                    e.self:Message(13, "This spell only effects summoned creatures.");
                end
            end
        end
    end
    ```



### EVENT_CLICKDOOR

#### Trigger

* When the client clicks on a door object.

Note that you would likely use this event in the zone player.pl file. Since doors have open types and destination fields stored in the database, most "simple" doors do not require a separate quest script. An example of a "simple" door would be any door that requires a single keyitem (by Item ID) to open, like the door to the basement in Befallen.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| doorid | int | `quest::say($doorid); # returns int` |
| version | int | `quest::say($version); # returns int` |

#### Example

* In this example, a player who is part of a Deepest Guk Adventure would be teleported to Deepest Guk when they click the doorway found in the Hollow Log.


=== "Perl"
    ```perl
    sub EVENT_CLICKDOOR {
        #:: Match if doorID is 1 - the door found in the Hollow Log that leads to Deepest Guk Adventures
        if ($doorid == 1) {
            #:: Create a variable to store the player's adventure zone instance ID
            $GukAInstance = quest::GetInstanceID("guka",50);
            #:: Match if the player has an instance
            if ($GukAInstance > 0) {
                #:: Teleport the player to their instance in Deepest Guk at the safe spot
                quest::MovePCInstance(229, $GukAInstance, 101, -841, 2.38);
            } 
            else {
                $client->Message(13, "You are not a part of a Deepest Guk adventure instance!");
            }
        }
    }
    ```


=== "Lua"
    ```lua
    function event_click_door(e)    
        --:: Match if doorID is 1 - the door found in the Hollow Log that leads to Deepest Guk Adventures
        local door_id = e.door:GetDoorID();
        if (door_id == 1) then
            --:: Create a variable to store the player's adventure zone instance ID
            local guka_instance = eq.get_instance_id("guka", 50);
            --:: Match if the player has an instance
            if (guka_instance > 0) then
                --:: Teleport the player to their instance in Deepest Guk at the safe spot
                e.self:MovePCInstance(229, guka_instance, 101, -841, 2.38);
            else
                e.self:Message(13, "You are not a part of a Deepest Guk adventure instance!");
            end
        end
    end
    ```



### EVENT_CLICK_OBJECT

#### Trigger

* When the client clicks on an object.

Note the similarity between this event and Perl EVENT_CLICKDOOR, since it is easy to confuse a door object (like a Plane of Knowledge Book) with Objects (IE Pottery wheels, Brew Barrels, etc.). You would likely use this event in the zone player.pl (or global_player.pl) files.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| objectid | int | `quest::say($objectid); # returns int` |
| clicker_id | int | `quest::say($clicker_id); # returns int` |

#### Example

* In this example, a message is displayed to a player when they open the Ogre Cultural Forge in Oggok.


=== "Perl"
    ```perl
    sub EVENT_CLICK_OBJECT {
        #:: Match to the ogre cultural forge in Oggok by object ID
        if ($objectid == 1075) {
            #:: Check to see if the player who clicked is a race other than Ogre
            if ($race ne "Ogre") {
                #:: Send the client a message in color 1 (gray)
                $client->Message(1,"The foul stench of Ogre overwhelms you as you open the forge.");
            } else {
                $client->Message(1,"Mmmm--dis smells just like home.");
            }
        }
    }
    ```


=== "Lua"
    ```lua
    function event_click_object(e)    
        --:: Match to the ogre cultural forge in Oggok by object ID
        local object_id = e.object:GetID();
        if (object_id == 1075) then
            --:: Check to see if the player who clicked is a race other than Ogre
            local race_id = e.self:GetRace();
            if (race_id != 10) then
                --:: Send the client a message in color 1 (gray)
                e.self:Message(1,"The foul stench of Ogre overwhelms you as you open the forge.");
            else
                e.self:Message(1,"Mmmm--dis smells just like home.");
            end
        end
    end
    ```
    


### EVENT_COMBAT

#### Trigger

* When an NPC enters or leaves combat.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| combat_state | int | `quest::say($combat_state); # returns int` |

#### Example

* In this example, the NPC will say some flavor text when entering combat.


=== "Perl"
    ```perl
    sub EVENT_COMBAT {
        #:: combat state 0 = False, 1 = True
        if ($combat_state == 1) {
            quest::say("Time to die!");
        }
    }
    ```
    

=== "Lua"
    ```lua
    function event_combat(e)
        --:: e.joined is true/false
        if (e.joined) then
            e.self:Say("Time to die!");
        end
    end
    ```



### EVENT_COMBINE_FAILURE

#### Trigger

* When a combine is unsuccessful.  You would likely use this event in your global / player.pl file.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| recipe_id | int | `quest::say($recipe_id); # returns int` |
| recipe_name | int | `quest::say($recipe_name); # returns int` |

#### Example

* In this example, we watch for a player failing the combine for a Hand Made Backpack and then tease them.


=== "Perl"
    ```perl
    sub EVENT_COMBINE_FAILURE {
        #:: Match Recipe 2686: "Hand Made Backpack" by ID
        if ($recipe_id == 2686) {
            #:: Send the client a message in color 15 (yellow)
            $client->Message(15,"Awww...now where are you going to put all of your stuff?");
        }
    }
    ```


=== "Lua"
    ```lua
    function event_combine_failure(e)
        --:: Match Recipe 2686: "Hand Made Backpack" by ID
        if (e.recipe_id == 2686) then
            --:: Send the client a message in color 15 (yellow)
            e.self:Message(15, "Awww...now where are you going to put all of your stuff?");
        end
    end
    ```



### EVENT_COMBINE_SUCCESS

#### Trigger

* When a combine is successful.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| recipe_id | int | `quest::say($recipe_id); # returns int` |
| recipe_name | int | `quest::say($recipe_name); # returns int` |

#### Example

* In this example, we send the client a message when they successfully combine a Hand Made Backpack


=== "Perl"
    ```perl
    sub EVENT_COMBINE_SUCCESS {
        #:: Match Recipe 2686: "Hand Made Backpack" by ID
        if ($recipe_id == 2686) {
            #:: Send the client a message in color 15 (yellow)
            $client->Message(15,"Yay, now you have a place to put all of your stuff!");
        }
    }
    ```


=== "Lua"
    ```lua
    function event_combine_success(e)
        --:: Match Recipe 2686: "Hand Made Backpack" by ID
        if (e.recipe_id == 2686) then
            --:: Send the client a message in color 15 (yellow)
            e.self:Message(15, "Yay, now you have a place to put all of your stuff!");
        end
    end
    ```



### EVENT_COMMAND

#### Trigger

* When a player says anything with the **#** command token preceding it that isn't handled by a system command.
* Exception: `#help` will invoke a special call so that scripted commands may be included in search listing.

Used in player files to handle script-based command processing.

#### Exports

| Name | Type | Details |
| :--- | :--- | :--- |
| command | string | A single, de-tokenized command name |
| args (perl) | string | An unseparated text line of additional parameters |
| args (lua) | string array | An array of single-word text parameters |

#### Example

* These examples are a working template for adding your own commands.


=== "Perl"
    ```perl
    sub EVENT_COMMAND {
        
        my %command_data = (
            "scriptcommand1" => [0, "perl script-based supplementation test command 1"],
            "scriptcommand2" => [100, "perl script-based supplementation test command 2"],
            "scriptcommand3" => [250, "perl script-based supplementation test command 3"]
        );
        
        if ($command eq "help") { # always a supplemental call from within the system command_help handler
            my $available_commands = 0;
            
            foreach my $command_name (sort {$a cmp $b} keys %command_data) {
                if ($status >= $command_data{$command_name}[0] && ($args eq "0" || $command_name=~/$args/i)) {
                    $client->Message(15, "#$command_name - $command_data{$command_name}[1]");
                    $available_commands = $available_commands + 1;
                }
            }
            
            return $available_commands;
        }
        elsif (exists($command_data{$command}) && $status >= $command_data{$command}[0]) {
            if ($command eq "scriptcommand1") {
                $client->Message(14, "doing '$command'...");
                return 1;
            }
            elsif ($command eq "scriptcommand2") {
                $client->Message(14, "doing '$command'...");
                return 1;
            }
            elsif ($command eq "scriptcommand3") {
                $client->Message(14, "doing '$command'...");
                return 1;
            }
        }
        
        return 0;
    }
    ```


=== "Lua"
    ```lua
    function event_command(e)
        
        local command_data = {
            ["scriptcommand1"] = {0, "lua script-based supplementation test command 1"},
            ["scriptcommand2"] = {100, "lua script-based supplementation test command 2"},
            ["scriptcommand3"] = {250, "lua script-based supplementation test command 3"}
        }
        
        if (e.command == "help") then -- always a supplemental call from within the system command_help handler
            local available_commands = 0;
            
            local key_sort = {}
            for key in pairs(command_data) do
                table.insert(key_sort, key);
            end
            table.sort(key_sort);
            
            for index, key in ipairs(key_sort) do
                if (e.self:Admin() >= command_data[key][1] and (e.args[1] == nil or key:findi(e.args[1]))) then
                    e.self:Message(15, "#" .. key .. " - " .. command_data[key][2]);
                    available_commands = available_commands + 1;
                end
            end
            
            return available_commands;
        elseif (command_data[e.command] ~= nil and e.self:Admin() >= command_data[e.command][1]) then
            if (e.command == "scriptcommand1") then
                e.self:Message(14, "doing '" .. e.command .. "'...");
                return 1;
            elseif (e.command == "scriptcommand2") then
                e.self:Message(14, "doing '" .. e.command .. "'...");
                return 1;
            elseif (e.command == "scriptcommand3") then
                e.self:Message(14, "doing '" .. e.command .. "'...");
                return 1;
            end
        end
        
        return eq.DispatchCommands(e); -- needed to invoke lua-based command system
    end
    ```



### EVENT_CONNECT

#### Trigger

* when a player connects to the world.

You would likely be using this event in your global_player.pl.

#### Example

* In this example, veteran AAs are awarded based on accumulated play time.


=== "Perl"
    ```perl
    sub EVENT_CONNECT {
        my %vet_aa = (481 => [31536000, 1, 1],
        482 => [63072000, 1, 1],
        483 => [94608000, 1, 1],
        484 => [126144000, 1, 1],
        485 => [157680000, 1, 1],
        486 => [189216000, 1, 1],
        487 => [220752000, 1, 1],
        511 => [252288000, 1, 1],
        2000 => [283824000, 1, 1],
        8081 => [315360000, 1, 1],
        8130 => [346896000, 1, 1],
        453 => [378432000, 1, 1],
        182 => [409968000, 1, 1],
        600 => [441504000, 1, 1]);
        foreach my $key (keys %vet_aa) {
            if ($vet_aa{$key}[2] && ($vet_aa{$key}[2] || $client->GetAccountAge() >= $vet_aa{$key}[0])) {
                $client->GrantAlternateAdvancementAbility($key, 1);
            }
        }
    }
    ```


=== "Lua"
    ```lua
    --[[ the main key is the ID of the AA
    --   the first set is the age required in seconds
    --   the second is if to ignore the age and grant anyways live test server style
    --   the third is enabled
    --]]
    vet_aa = {
        [481]  = { 31536000, true, true}, -- Lesson of the Devote 1 yr
        [482]  = { 63072000, true, true}, -- Infusion of the Faithful 2 yr
        [483]  = { 94608000, true, true}, -- Chaotic Jester 3 yr
        [484]  = {126144000, true, true}, -- Expedient Recovery 4 yr
        [485]  = {157680000, true, true}, -- Steadfast Servant 5 yr
        [486]  = {189216000, true, true}, -- Staunch Recovery 6 yr
        [487]  = {220752000, true, true}, -- Intensity of the Resolute 7 yr
        [511]  = {252288000, true, true}, -- Throne of Heroes 8 yr
        [2000] = {283824000, true, true}, -- Armor of Experience 9 yr
        [8081] = {315360000, true, true}, -- Summon Resupply Agent 10 yr
        [8130] = {346896000, true, true}, -- Summon Clockwork Banker 11 yr
        [453]  = {378432000, true, true}, -- Summon Permutation Peddler 12 yr
        [182]  = {409968000, true, true}, -- Summon Personal Tribute Master 13 yr
        [600]  = {441504000, true, true}, -- Blessing of the Devoted 14 yr
    }
    
    function event_connect(e)
        local age = e.self:GetAccountAge();
        for aa, v in pairs(vet_aa) do
            if v[3] and (v[2] or age >= v[1]) then
                e.self:GrantAlternateAdvancementAbility(aa, 1)
            end
        end
    end
    ```



### EVENT_DEATH

#### Trigger

* When the NPC dies. Fires before death finishes.

#### Exports

| Name | Type | Details |
| :--- | :--- | :--- |
| client | client | client who killed mob |
| npc | npc | npc that was killed |
| killer_id | int | client ID of killer. (Does not seem castable to mob) |
| killer_damage | int | How much damage was dealt on killing blow |
| killer_spell | int | Spell ID used to kill mob |
| killer_skill | int | Skill ID used to kill mob |
| charid | int | Character ID who killed mob |
| class | string | Class Name who killed mob |
| faction | int | Faction comparison of killed mob vs killer |
| h | float | heading of mob during death |
| hpratio | float | percent health of mob after death (negative value) |
| mlevel | int | level of mob killed |
| mname | string | name of mob killed |
| mobid | int | id of mob killed |
| name | string | name of killer |
| race | string | race of killer |
| status | int | account status of killer |
| uguild_id | int | uguild of killer |
| ulevel | int | level of killer |
| userid | int | user id of killer |
| x | float | x position of killed mob |
| y | float | y position of killed mob |
| z | float | z position of killed mob |
| zonehour | int | hour of zone when mob died |
| zoneid | int | zone id where mob died |
| zoneln | string | long name of zone where mob died |
| zonemin | int | minimum level to enter zone where mob died |
| zonesn | string | short name of zone where mob died |
| zonetime | int | time of zone where mob died |
| zoneweather | int |  |

### EVENT_DEATH_COMPLETE

#### Trigger

* When the NPC dies.

Often used to spawn adds or send signals upon the death of an NPC.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| killer_id | int | `quest::say($killer_id); # returns int` |
| killer_damage | int | `quest::say($killer_damage); # returns int` |
| killer_spell | int | `quest::say($killer_spell); # returns int` |
| killer_skill | int | `quest::say($killer_skill); # returns int` |

#### Example

* In this example, we spawn a fire beetle after the death of our NPC at the NPC's location.


=== "Perl"
    ```perl
    sub EVENT_DEATH_COMPLETE {
        #:: Spawn a 2024 - a_fire_beetle by NPC Type ID, grid 0, guildwarset 0, current X, Y, Z, and heading
        quest::spawn2(2024,0,0,$x,$y,$z,$h);
    }
    ```


=== "Lua"
    ```lua
    function event_death_complete(e)
        --:: Spawn a 2024 - a_fire_beetle by NPC Type ID, grid 0, guildwarset 0, current X, Y, Z, and heading
        eq.spawn2(2024, 0, 0, e.self:GetX(), e.self:GetY(), e.self:GetZ(), e.self:GetHeading());
    end
    ```



### EVENT_DEATH_ZONE

#### Trigger

* When the NPC dies.

**Used by the zone controller.**

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| killer_id | int | `quest::say($killer_id); # returns int` |
| killer_damage | int | `quest::say($killer_damage); # returns int` |
| killer_spell | int | `quest::say($killer_spell); # returns int` |
| killer_skill | int | `quest::say($killer_skill); # returns int` |
| killer_npc_id | int | `quest::say($killer_npc_id); # returns int` |

### EVENT_DESTROY_ITEM

#### Trigger

* When a client destroys an item.

Used mainly for logging purposes.

### EVENT_DISCONNECT

#### Trigger

* When a player disconnects from the world.

Used mainly for logging purposes.

### EVENT_DISCOVER_ITEM

#### Trigger

* When an item is discovered.

Used in conjunction with World Rule EnableDiscoveredItems.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| itemid | int | `quest::say($itemid); # returns int` |

#### Example


=== "Perl"
    ```perl
    sub EVENT_DISCOVER_ITEM {
        #:: Create a scalar variable to store the item link
        $discovereditem = quest::varlink($itemid);
        #:: Shout the discovery to all zones
        quest::shout2("$name has discovered $discovereditem!  Yay!");
    }
    ```


=== "Lua"
    ```lua
    function event_discover_item(e)    
        --:: Create a scalar variable to store the item link
        local item_link = eq.item_link(e.item:GetID());
        --:: Emote the discovery to all zones
        eq.world_emote(335, "$name has discovered " .. item_link .. "!  Yay!");
    end
    ```



### EVENT_DROP_ITEM

#### Trigger

* When a client drops an item.

Mainly used for logging purposes.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| quantity | int | `quest::say($quantity); # returns int` |
| itemname | int | `quest::say($itemname); # returns int` |
| itemid | int | `quest::say($itemid); # returns int` |
| spell_id | int | `quest::say($spell_id); # returns int` |
| slotid | int | `quest::say($slotid); # returns int` |

### EVENT_DUEL_LOSE

#### Trigger

* When a client loses a duel.

You would use this event in the global global_player.pl file.

#### Example

* In this example, we set a data bucket to keep track of a player's dueling failures.


=== "Perl"
    ```perl
    sub EVENT_DUEL_LOSE {
        my $bucket_name = $client->CharacterID() . "-losses";
        my $bucket_value = (quest::get_data($bucket_name) + 1);
        quest::set_data($bucket_name, $bucket_value);
    }
    ```


=== "Lua"
    ```lua
    function event_duel_lose(e)
        local bucket_name = e.self:CharacterID() .. "-losses";
        local bucket_value = (eq.get_data(bucket_name) + 1);
        eq.set_data(bucket_name, bucket_value);
    end
    ```



### EVENT_DUEL_WIN

#### Trigger

* when a client wins a duel.

You would use this event in the global global_player.pl file.

#### Example

* In this example, we set a data bucket to keep track of a player's dueling wins.


=== "Perl"
    ```perl
    sub EVENT_DUEL_WIN {
        my $bucket_name = $client->CharacterID() . "-wins";
        my $bucket_value = (quest::get_data($bucket_name) + 1);
        quest::set_data($bucket_name, $bucket_value);
    }
    ```


=== "Lua"
    ```lua
    function event_duel_win(e)
        local bucket_name = e.self:CharacterID() .. "-wins";
        local bucket_value = (eq.get_data(bucket_name) + 1);
        eq.set_data(bucket_name, bucket_value);
    end
    ```
    


### EVENT_ENTER

#### Trigger

* When a client enters a mob's proximity (as defined by quest::set_proximity(min_x, max_x, min_y, max_y, min_z, max_z)).

#### Example

* In this example, we first set the NPC's proximity; when a player enters the proximity, they become PVP (red).


=== "Perl"
    ```perl
    sub EVENT_SPAWN {
        #:: Create a proximity, 100 units across, 100 units tall, without proximity say
        quest::set_proximity($x - 50, $x + 50, $y - 50, $y + 50, $z - 50, $z + 50, 0);
    }
    
    sub EVENT_ENTER {
        #:: Turn pvp on
        quest::pvp("on");
    }
    ```


=== "Lua"
    ```lua
    function event_spawn(e)
        --:: Create a proximity, 100 units across, 100 units tall, without proximity say
        local x = e.self:GetX();
        local y = e.self:GetY();
        local z = e.self:GetZ();
        eq.set_proximity((x - 50), (x + 50), (y - 50), (y + 50), (z - 50), (z + 50), 0);    
    end
    
    function event_enter(e)
        --:: Turn pvp on
        e.other:SetPVP("on");
    end
    ```
    


### EVENT_ENTER_AREA

#### Trigger

* when a client enters the area of a mob.

### EVENT_ENTERZONE

#### Trigger

* When a player enters the zone.  Likely you will add to the zone player.pl file.

#### Example

* In this example we remove the LDON compass mark.


=== "Perl"
    ```perl
    sub EVENT_ENTERZONE {
        #:: Clear the LDON Compass mark
        $client->ClearCompassMark();
    }
    ```


=== "Lua"
    ```lua
    function event_enterzone(e)
        --:: Clear the LDON Compass mark
        e.self:ClearCompassMark();
    end
    ```



* In this example we create a LDON compass mark


=== "Perl"
    ```perl
    sub EVENT_ENTERZONE {
        #:: Create a scalar for storing the instance ID
        $RujDInstance = quest::GetInstanceID("rujd",50);
        #:: If the instance ID exists, it should be greater than 0--mark the player's compass if it is
        if ($RujDInstance > 0) {
            #:: Create a line on the compass leading the player to X,Y,Z
            $client->MarkCompassLoc(-157.09, 19.31, 100);
        }
    }
    ```


=== "Lua"
    ```lua
    function event_enterzone(e)
        --:: Create a variable for storing the instance ID
        local rujd_instance = eq.get_instance_id("rujd", 50);
        --:: If the instance ID exists, it should be greater than 0--mark the player's compass if it is
        if (rujd_instance > 0) then
            --:: Create a line on the compass leading the player to X,Y,Z
            e.self:MarkCompassLoc(-157.09, 19.31, 100);
        end    
    end
    ```



### EVENT_ENVIRONMENTAL_DAMAGE

#### Trigger

* When taking any sort of environmental damage.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| env_damage | int | `quest::say($env_damage); # returns int` |
| env_damage_type | int | `quest::say($env_damage_type); # returns int` |
| env_final_damage | int | `quest::say($env_final_damage); # returns int` |

### EVENT_EQUIP_ITEM

#### Trigger

* When a player equips an item.

### EVENT_EXIT

#### Trigger

* When a client leaves a mob's proximity (as defined by quest::set_proximity).

#### Example

* In this example, we first set the NPC's proximity; when a player exits the proximity, they are no longer PVP (blue).

#### Example


=== "Perl"
    ```perl
    sub EVENT_SPAWN {
        #:: Create a proximity, 100 units across, 100 units tall, without proximity say
        quest::set_proximity($x - 50, $x + 50, $y - 50, $y + 50, $z - 50, $z + 50, 0);
    }
    
    sub EVENT_EXIT {
        #:: Turn pvp off
        quest::pvp("off");
    }
    ```


=== "Lua"
    ```lua
    function event_spawn(e)
        --:: Create a proximity, 100 units across, 100 units tall, without proximity say
        local x = e.self:GetX();
        local y = e.self:GetY();
        local z = e.self:GetZ();
        eq.set_proximity((x - 50), (x + 50), (y - 50), (y + 50), (z - 50), (z + 50), 0);    
    end
    
    function event_exit(e)
        --:: Turn pvp off
        e.other:SetPVP("off");
    end
    ```



### EVENT_FEIGN_DEATH

#### Trigger

* When a client feigns death.

#### Example


=== "Perl"
    ```perl
    sub EVENT_FEIGN_DEATH {
        #:: See if the player has a pet
        if ($client->GetPetID()) {
            #:: Identify the pet by ID and kill it
            $pet_entity = $entity_list->GetMobByID($client->GetPetID());
            $pet_entity->Kill();
        }
    }
    ```


=== "Lua"
    ```lua
    function event_feign_death(e)
        --:: See if the player has a pet
        local pet_id = e.other:GetPet():GetID();
        if (pet_id > 0) then
            --:: Identify the pet by ID and kill it
            local pet_entity = entity_list:GetMob(pet_id);
            pet_entity:Kill();
        end
    end
    ```
    


### EVENT_FISH_FAILURE

#### Trigger

* When a client fails at fishing.

You would use this event in the zone player.pl file.

#### Example

* In this example, we set a data bucket to keep track of a player's fishing failures.
* You would use this event in the global global_player.pl file.


=== "Perl"
    ```perl
    sub EVENT_FISH_FAILURE {
        my $bucket_name = $client->CharacterID() . "-fish-fail";
        my $bucket_value = (quest::get_data($bucket_name) + 1);
        quest::set_data($bucket_name, $bucket_value);
        $client->Message(1, "Maybe you're using the wrong bait!")
    }
    ```


=== "Lua"
    ```lua
    function event_fish_failure(e)
        local bucket_name = e.self:CharacterID() .. "-fish-fail";
        local bucket_value = (eq.get_data(bucket_name) + 1);
        eq.set_data(bucket_name, bucket_value);
        e.self:Message(1, "Maybe you're using the wrong bait!");
    end
    ```



### EVENT_FISH_START

#### Trigger

* when a client starts fishing.

You would use this event in the zone player.pl file.

#### Example

* In this example, a message is displayed by the client if they start fishing.


=== "Perl"
    ```perl
    sub EVENT_FISH_START {
        $client->Message(1, "You crack a beer and toss your line in.");
    }
    ```


=== "Lua"
    ```lua
    function event_fish_start(e)
        e.self:Message(1, "You crack a beer and toss your line in.");
    end
    ```



### EVENT_FISH_SUCCESS

#### Trigger

* when a client succeeds at fishing.

You would use this event in the zone player.pl file or the global global_player.pl file.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| fished_item | int | `quest::say($fished_item); # returns int` |

#### Example

* In this example, we set a data bucket to keep track of a player's fishing successes.
* You would use this event in the global global_player.pl file.


=== "Perl"
    ```perl
    sub EVENT_FISH_SUCCESS {
        my $bucket_name = $client->CharacterID() . "-fish-success";
        my $bucket_value = (quest::get_data($bucket_name) + 1);
        quest::set_data($bucket_name, $bucket_value);
    }
    ```


=== "Lua"
    ```lua
    function event_fish_success(e)
        local bucket_name = e.self:CharacterID() .. "-fish-success";
        local bucket_value = (eq.get_data(bucket_name) + 1);
        eq.set_data(bucket_name, bucket_value);
        e.self:Message(1, "You caught something!");
    end
    ```



### EVENT_FORAGE_FAILURE

#### Trigger

* When a client fails at foraging.

You would use this event in the zone player.pl or global global_player.pl file.

#### Example

* In this example, we set a data bucket to keep track of a player's foraging failures.


=== "Perl"
    ```perl
    sub EVENT_FORAGE_FAILURE {
        my $bucket_name = $client->CharacterID() . "-forage-fail";
        my $bucket_value = (quest::get_data($bucket_name) + 1);
        quest::set_data($bucket_name, $bucket_value);
        $client->Message(1, "You didn't find anything!")
    }
    ```


=== "Lua"
    ```lua
    function event_forage_failure(e)
        local bucket_name = e.self:CharacterID() .. "-forage-fail";
        local bucket_value = (eq.get_data(bucket_name) + 1);
        eq.set_data(bucket_name, bucket_value);
        e.self:Message(1, "You didn't find anything!");
    end
    ```



### EVENT_FORAGE_SUCCESS

#### Trigger

* when a client succeeds at foraging.

You would use this event in the zone player.pl file or the global global_player.pl file.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| foraged_item | int | `quest::say($foraged_item); # returns int` |

#### Example

* In this example, we set a data bucket to keep track of a player's foraging successes.


=== "Perl"
    ```perl
    sub EVENT_FORAGE_SUCCESS {
        my $bucket_name = $client->CharacterID() . "-forage-success";
        my $bucket_value = (quest::get_data($bucket_name) + 1);
        quest::set_data($bucket_name, $bucket_value);
        $client->Message(1, "You found something!")
    }
    ```


=== "Lua"
    ```lua
    function event_forage_success(e)
        local bucket_name = e.self:CharacterID() .. "-forage-success";
        local bucket_value = (eq.get_data(bucket_name) + 1);
        eq.set_data(bucket_name, bucket_value);
        e.self:Message(1, "You found something!");
    end
    ```



### EVENT_GROUP_CHANGE

#### Trigger

* when a group change occurs.

### EVENT_HATE_LIST

#### Trigger

* When a mob's hate list is changed.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| hate_state | int | `quest::say($hate_state); # returns int` |

#### Example

In this example, some flavor text is added as a player is added and removed from the NPC's hate list.


=== "Perl"
    ```perl
    sub EVENT_HATE_LIST {
        if ($hate_state == 1) {
            quest::say("You're gonna die!");
        } else {
            quest::say("You're no match for my might!");
        }
    }
    ```


=== "Lua"
    ```lua
    function event_hate_list(e)
        if (e.joined) then
            e.self:Say("You're gonna die!");
        else
            e.self:Say("You're no match for my might!");
        end
    end
    ```



### EVENT_HP

#### Trigger

* When a mob's HP dropping below a threshold (as defined by quest::setnexthpevent()).

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| hpevent | int | `quest::say($hpevent); # returns int` |
| inchpevent | int | `quest::say($inchpevent); # returns int--incoming HP event` |

#### Example

* In this example, we set the HP event threshold, and then depop the mob when the threshold is reached.


=== "Perl"
    ```perl
    sub EVENT_SPAWN {
        #:: Set the HP event threshold for 50 percent health
        quest::setnexthpevent(50);
    }
    
    sub EVENT_HP {
        #:: Match when the threshold is met
        if ($hpevent == 50) {
            quest::depop();
        }
    }
    ```


=== "Lua"
    ```lua
    function event_spawn(e)
        --:: Set the HP event threshold for 50 percent health
        eq.set_next_hp_event(50);
    end
    
    function event_hp(e)
        --:: Match when the threshold is met
        if (e.hp_event == 50) then
            e.self:Depop();
        end
    end
    ```



### EVENT_ITEM

#### Trigger

* When an item or money is turned into the mob.

#### Example

* In this example, we turn in 325 Platinum, a Ring of the Ancients, and a Shadowed Rapier in exchange for our Journeyman's Boots


=== "Perl"
    ```perl
    sub EVENT_ITEM {
         #:: Try to take 325 Platinum from the client
        if ($client->TakeMoneyFromPP(325000, 1)) {
            #:: Match turn for 12268 - Ring of the Ancients and 7100 - Shadowed Rapier
            if (plugin::check_handin(%itemcount, 12268 => 1, 7100 => 1)) {
                quest::say("The time to trade has come!! I am now rich and you are now fast. Take the Journeyman Boots and run like the wind.");
                #:: Give a 2300 - Journeyman's Boots
                quest::summonitem(2300);
                #:: Grant a small amount of experience
                quest::exp(1250);
            }
        }
        #:: Return unused items
        plugin::return_items(%itemcount);
    }
    ```


=== "Lua"
    ```lua
    function event_trade(e)
        --:: Require  items library.
        local item_lib = require("items");
        --:: Try to take 325 Platinum from the client
        if (e.other:TakeMoneyFromPP(325000, 1)) then
            --:: Match turn for 12268 - Ring of the Ancients and 7100 - Shadowed Rapier
            if (item_lib.check_turn_in(e.self, e.trade, {item1 = 12268, item2 = 7100})) then
                e.self:Say("The time to trade has come!! I am now rich and you are now fast. Take the Journeyman Boots and run like the wind.");
                --:: Give a 2300 - Journeyman's Boots
                e.other:SummonItem(2300);
                --:: Grant a small amount of experience
                e.other:AddEXP(1250);
            end
        end
        --:: Return unused items
        item_lib.return_items(e.self, e.other, e.trade);
    end
    ```
    


### EVENT_ITEM_CLICK

#### Trigger

* When an item is clicked.

This is a useful script to put in the Global quest scripts directory, so that you can make an item click work anywhere in the world.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| itemid | int | `quest::say($itemid); # returns int` |
| itemname | int | `quest::say($itemname); # returns int` |
| slotid | int | `quest::say($slotid); # returns int` |
| spell_id | int | `quest::say($spell_id); # returns int` |

#### Example

* This example is taken from the Evil Eye Costume Kit, which is part of the Halloween Costume illusion items.
* Note that the scriptfileid field for the item is set to 30073 in the database.
* Note that a corresponding quest file exists at global/items/script_30073.pl.


=== "Perl"
    ```perl
    sub EVENT_ITEM_CLICK {
        #:: Use == for numeric comparison to Item ID 54711 - Evil Eye Costume Kit
        if ($itemid == 54711) {
            #:: Change the player's race to 469 - Evil Eye
            quest::playerrace(469);
        } 
    }
    ```


=== "Lua"
    ```lua
    function event_item_click(e)
        --:: Use == for numeric comparison to Item ID 54711 - Evil Eye Costume Kit
        local item_id = e.self:GetID();
        if (item_id == 54711) then
            --:: Change the player's race to 469 - Evil Eye
            e.owner:ChangeRace(469);
        end
    end
    ```



### EVENT_ITEM_CLICK_CAST

#### Trigger

* When a client casts the click effect on an item.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| itemid | int | `quest::say($itemid); # returns int` |
| itemname | int | `quest::say($itemname); # returns int` |
| slotid | int | `quest::say($slotid); # returns int` |
| spell_id | int | `quest::say($spell_id); # returns int` |

### EVENT_ITEM_ENTER_ZONE

Called when an item that would trigger EVENT_SCALE_CALC is in the inventory when a player zones in.

### EVENT_ITEM_TICK

#### Trigger

* when the click effect of an item ticks.

### EVENT_KILLED_MERIT

#### Trigger

On NPC death and applies to the group that did the most damage to the NPC (IE the group that got XP for the kill, assuming there was XP; or the group that gets loot rights to the NPC, assuming that there was loot). Although not often used, this event gives you the opportunity to assign quest globals or data buckets (IE for character flags), or update tasks.

#### Example


=== "Perl"
    ```perl
    sub EVENT_KILLED_MERIT {
        #:: Get the name of the mob to use in the data bucket
        my $slain = $npc->GetCleanName();
        #:: Set the quest data bucket--this would apply to all group members
        my $bucket_name = $client->CharacterID() . "-Slain";
        quest::set_data($bucket_name, 1);
        #:: Display an emote message to each client in yellow to notify them that they received credit
        $client->Message(15, "You have received credit for killing $slain.");
    }
    ```


=== "Lua"
    ```lua
    function event_killed_merit(e)
        --:: Get the name of the mob to use in the quest global
        local slain = e.self:GetCleanName();
        --:: Set the quest global--this would apply to all group members
        local bucket_name = e.other:CharacterID() .. "-Slain";
        eq.set_data(bucket_name, 1);
        --:: Display an emote message to each client in yellow to notify them that they received credit
        e.other:Message(15, "You have received credit for killing " .. slain .. ".");    
    end
    ```



### EVENT_LEAVE_AREA

#### Trigger

* when a client leaves a mob's area.

### EVENT_LEVEL_UP

#### Trigger

* When the player gains a level.


=== "Perl"
    ```perl
    sub EVENT_LEVEL_UP {
        #:: Congratulate the player
        $client->Message(15, "Congratulations on level $ulevel!");
    }
    ```


=== "Lua"
    ```lua
    function event_level_up(e)
        --:: Congratulate the player
        e.self:Message(15, "Congratulations on level " .. e.self:GetLevel() .. "!");
    end
    ```



### EVENT_LOOT

#### Trigger

* when player successfully loots an item from a corpse.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| looted_id | int | `quest::say($looted_id); # returns int` |
| looted_charges | int | `quest::say($looted_charges); # returns int` |
| corpse | int | `quest::say($corpse); # returns int` |

#### Example

* This example uses the loot event to match a particular item and corpse.

Note that we use the NPC's name (Fippy_Darkpaw) and not the numeric corpse ID (IE 249) or full corpse name (IE Fippy_Darkpaw`s_corpse249).


=== "Perl"
    ```perl
    sub EVENT_LOOT {
        #:: Use == for numeric comparison to Item ID 60396 - Fippy's Paw
        #:: Use eq for string comparison to Fippy_Darkpaw's corpse
        if ($looted_id == 60396 && $corpse eq "Fippy_Darkpaw") {
            $client->Message(15, "The bloody stump of Fippy's paw--it's a lot smaller than you thought it would be.");
        }
    }
    ```


=== "Lua"
    ```lua
    function event_loot(e)
        --:: Use == for numeric comparison to Item ID 60396 - Fippy's Paw
        local looted_id = e.self:GetID();
        local corpse_name = e.corpse:GetCleanName()
        if (looted_id == 60396 and corpse_name:findi("Fippy")) then
            e.owner:Message(15, "The bloody stump of Fippy's paw--it's a lot smaller than you thought it would be.");
        end
    end
    ```



### EVENT_NPC_SLAY

#### Trigger

* When an NPC slays another NPC.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| killed | int | `quest::say($killed); # returns int NPCTypeID` |

#### Example

* In this example, we add some flavor text when the Exterminator kills the rats.


=== "Perl"
    ```perl
    sub EVENT_NPC_SLAY {
        quest::say("Another unworthy opponent. Never cross Mining Guild 628!!");
    }
    ```


=== "Lua"
    ```lua
    function event_slay(e)
        e.self:Say("Another unworthy opponent. Never cross Mining Guild 628!!");
    end
    ```



### EVENT_PLAYER_PICKUP

#### Trigger

* When a player picks up an object from the ground.  You would likely use this event in your zone player.pl file.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| picked_up_id | int | `quest::say($picked_up_id); # returns int` |
| picked_up_entity_id | int | `quest::say($picked_up_entity_id); # returns int` |

#### Example

-- In this example, when the player picks up a Chalice of Conquest, a signal is sent to another NPC


=== "Perl"
    ```perl
    #:: Chalice of Conquest quest
    sub EVENT_PLAYER_PICKUP {
        #:: Match 12274 - Chalice of Conquest, ground spawn created by #Captain_Klunga.pl
        if ($picked_up_id == 12274) {
            #:: Send a signal to Dagnor's Cauldron >> #Captain_Klunga (70072)
            quest::signalwith(70072, 1);
        }
    }
    ```
    

=== "Lua"
    ```lua
    --:: Chalice of Conquest quest
    function event_player_pickup(e)
        --:: Match 12274 - Chalice of Conquest, ground spawn created by #Captain_Klunga.pl
        local picked_up_id = e.item:GetID()
        if (picked_up_id == 12274) then
            --:: Send a signal to Dagnor's Cauldron >> #Captain_Klunga (70072)
            eq.signal(70072, 1);
        end
    end
    ```



### EVENT_POPUPRESPONSE

#### Trigger

* When a player clicks a button on a popup.

Used with quest::popup.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| popupid | int | `quest::say($popupid); # returns int` |

#### Example

* a well-known example from the Guild Lobby portal pool.


=== "Perl"
    ```perl
    sub EVENT_POPUPRESPONSE {
        #:: Triggered by sub EVENT_ENTER: quest::popup('Teleport', 'Teleport to The Plane of Hate?', 666, 1, 0);
        if ($popupid == 666) {
            #:: Teleport the player to hateplaneb
            quest::movepc(186,-393,656,3);
        }
    }
    ```


=== "Lua"
    ```lua
    function event_popup_response(e)    
        --:: Triggered by function event_enter: eq.popup('Teleport', 'Teleport to The Plane of Hate?', 666, 1, 0);
        if (e.popup_id == 666) then
            --:: Teleport the player to hateplaneb
            e.self:MovePC(186, -393, 656, 3);
        end
    end
    ```



### EVENT_PROXIMITY_SAY

#### Trigger

* When the client enters a mob's proximity and uses the appropriate text trigger supplied beneath this event.

Note that you must enable quest::set_proximity().

#### Example

* In this example, we establish a proximity around the NPC and enable the NPC to listen for say messages in the proximity (IE without having the mob targeted, necessarily); we then match for a "hail" message and attack anyone foolish enough to say hello.


=== "Perl"
    ```perl
    sub EVENT_SPAWN {
        #:: Create a proximity, 100 units across, 100 units tall, enable proximity say
        quest::set_proximity($x - 50, $x + 50, $y - 50, $y + 50, $z - 50, $z + 50, 1);
        #:: Also, enable proximity say
        quest::enable_proximity_say();
    }
    
    sub EVENT_PROXIMITY_SAY {
        #:: Match say message for "hail", /i for case insensitive
        if ($text=~/hail/i) {
            quest::say("Hello, $name!");
        }
    }
    ```


=== "Lua"
    ```lua
    function event_spawn(e)
        --:: Create a proximity, 100 units across, 100 units tall, enable proximity say
        local x = e.self:GetX();
        local y = e.self:GetY();
        local z = e.self:GetZ();
        eq.set_proximity((x - 50), (x + 50), (y - 50), (y + 50), (z - 50), (z + 50), 1);    
    end
    
    function event_proximity_say(e)
        --:: Match say message for "hail", case-insensitive
        if (e.message:findi("hail")) then
            e.self:Say("Hello, " .. e.other:GetCleanName() .. "!");
        end
    end
    ```
    


### EVENT_RESPAWN

#### Trigger

* on respawn

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| option | int | `quest::say($option); # returns int` |
| resurrect | int | `quest::say($resurrect); # returns int` |

### EVENT_SAY

#### Trigger

* when a mob is targeted and the player types something.

#### Exports

| Name | Type | Details |
| :--- | :--- | :--- |
| client | client | Client who did say event |
| npc | npc | Npc who is handling say event |
| charid | int | character id of who did say event |
| class | string | class of who did say event |
| data | int | unknown? 124078 |
| faction | int | faction comparison of who did say and npc |
| h | float | heading position of npc |
| hpratio | float | hp ratio e.g. 100 |
| instanceid | int | instance id of zone, typically 0 |
| instanceversion | int | instance version of zone, typically 0 |
| langid | int | language id, common is 0 |
| mlevel | int | mob level of npc |
| mname | string | mob name of npc |
| mobid | int | mob entity id of npc |
| name | string | name of who did say event |
| race | string | race of who did say event |
| status | int | account status of who did say event |
| text | string | Text of who did say event |
| uguild_id | int | guild id of who did say event |
| uguildrank | int | guild rank of who did say event |
| ulevel | int | level of who did say event |
| userid | int | user id of who did say event |
| x | float | x position of npc |
| y | float | y position of npc |
| z | float | z position of npc |
| zonehour | int | hour of zone when mob died |
| zoneid | int | zone id where mob died |
| zoneln | string | long name of zone where mob died |
| zonemin | int | minimum level to enter zone where mob died |
| zonesn | string | short name of zone where mob died |
| zonetime | int | time of zone where mob died |
| zoneweather | int | weather of zone where mob died |

#### Example

* This example is a response to a "hail"


=== "Perl"
    ```perl
    sub EVENT_SAY {
        #:: Checks if the text is like "Hail", the "/i" is for case-insensitive.
        if ($text=~/Hail/i) {
            quest::say("Hello, $name!");
        }
    }
    ```


=== "Lua"
    ```lua
    function event_say(e)
        --:: Checks if the text is like "Hail", case-insensitive
        if (e.message:findi("hail")) then
            e.self:Say("Hello, " .. e.other:GetCleanName() .. "!");
        end    
    end
    ```



* This example additionally checks the language of the "hail", and will only respond to text in that language.


=== "Perl"
    ```perl
    sub EVENT_SAY {
        #:: Checks to see if the language is Thieves Cant (language ID 10)
        if ($langid == 10) {
            #:: Checks if the text is like "Hail", the "i" is for case-insensitive.
            if ($text=~/Hail/i) {
                # Respond, using the same language
                quest::say("Hello, $name!",10);
            }
        }
    }
    ```


=== "Lua"

    
    ```lua
    function event_say(e)
        --:: Checks to see if the language is Thieves Cant (language ID 10)
        if (e.language == 10) then
            --:: Checks if the text is like "Hail", the "i" is for case-insensitive.
            if (e.message:findi("hail")) then
                --:: Respond, using the same language
                e.self:Say("Hello, $name!", 10);
            end
        end
    end
    ```



### EVENT_SCALE_CALC

#### Trigger

* When an item is equipped to scale the item--probably should zone.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| itemid | int | `quest::say($itemid); # returns int` |
| itemname | int | `quest::say($itemname); # returns int` |

#### Example

* In this example, we scale a charm item:  40342 - Charm of Exotic Speech.  We have a script global / items / 40342.pl
* The item has a charmfile and charmfileID assigned


=== "Perl"
    ```perl
    sub EVENT_SCALE_CALC {
        #:: Used for charms that scale with number of rare languages learned
        my $languages_mastered = 0;
        my $scale = 0;
        #:: Check each rare language: Old Erudian through Elder Dragon
        for (my $i = 11; $i <= 22; $i++) {
            #:: Check if the client has mastered the language
            if ($client->GetLanguageSkill($i) == 100) {
                $languages_mastered++;
            }
        }
        $scale = ($languages_mastered / 12);
        $questitem->SetScale($scale);
    }
    ```


=== "Lua"
    ```lua
    function event_scale_calc(e)
        --:: Used for charms that scale with number of rare languages learned
        local languages_mastered = 0;
        local scale = 0;
        --:: Check each rare language: Old Erudian through Elder Dragon
        for i = 11, 22, 1 do
            --:: Check if the client has mastered the language
            if (e.owner:GetLanguageSkill(i) == 100) then
                languages_mastered = (languages_mastered + 1);
            end
        end
        scale = (languages_mastered / 12);
        e.self:SetScale(scale);
    end
    ```



### EVENT_SIGNAL

#### Trigger

* Triggered using quest::signal(NPC_ID, wait_time) or quest::signalwith(NPC_ID,signal_ID,wait_time).

Generally a way to have one NPC cause another NPC to do something. Often used with "controller" NPCs that coordinate events.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| signal | int | `quest::say($signal); # returns int` |

#### Example

* This example increments a counter each time a signal with the appropriate ID is received.


=== "Perl"
    ```perl
    my $count = 0;
    
    sub EVENT_SIGNAL {
        #:: Signal 1 is from the clockwork spiders being killed.
        if ($signal == 1) {
            $count++;
            if ($count == 1) {
                #:: Start a three minute timer to spawn targetable Manaetic Behemoth
                quest::settimer("wake", 180);
            }
        }
        #:: Signal 2 is from the clockwork spiders reaching Manaetic Behemoth.
        elsif ($signal == 2) {
            #:: Reset the count and make them start over.
            $count = 0;
            quest::stoptimer("wake");
        }
    }
    
    sub EVENT_TIMER {
        #:: This uses eq for a string comparison to match the timer "wake".
        #:: Check the count to make sure the clockwork spiders were killed and not just kited. 
        if ($timer eq "wake" && $count >= 12) {
            quest::stoptimer("wake");
            #:: Spawn the targetable version of Manaetic Behemoth in place
            quest::spawn2(206074,0,0,$x,$y,$z,0);
            #:: Depop the untargetable version of Manaetic Behemoth with respawn timer active.
            quest::depop_withtimer();
        }
    }
    ```


=== "Lua"
    ```lua
    count = 0;
    
    function event_signal(e)
        --:: Signal 1 is from the clockwork spiders being killed.
        if (e.signal == 1) then
            count = (count + 1);
            if (count == 1) then
                --:: Start a three minute timer to spawn targetable Manaetic Behemoth
                quest::set_timer("wake", 180);
            end
        --:: Signal 2 is from the clockwork spiders reaching Manaetic Behemoth.
        else if (e.signal == 2) then
            --:: Reset the count and make them start over.
            count = 0;
            eq.stop_timer("wake");
        end
    end
    
    function event_timer(e)
        --:: Check the count to make sure the clockwork spiders were killed and not just kited. 
        if (e.timer == "wake" && count >= 12) then
            eq.stop_timer("wake");
            --:: Spawn the targetable version of Manaetic Behemoth in place
            eq.spawn2(206074, 0, 0, e.self:GetX(), e.self:GetY(), e.self:GetZ(), 0);
            --:: Depop the untargetable version of Manaetic Behemoth with respawn timer active.
            eq.depop_with_timer();
        end
    end
    ```



### EVENT_SLAY

#### Trigger

* whenever an NPC kills a player.

Often this event is used for some flavor messages, or to spawn adds. Do not confuse this event with EVENT_DEATH_COMPLETE or EVENT_DEATH, which are used when a player kills an NPC.

#### Example

* This is a well-known example from Emperor Ssraeshza, who mocks any player that he kills


=== "Perl"
    ```perl
    sub EVENT_SLAY {
        quest::say("Your god has found you lacking.");
    }
    ```


=== "Lua"
    ```lua
    function event_slay(e)
        e.self:Say("Your god has found you lacking.");
    end
    ```
    


### EVENT_SPAWN

#### Trigger

* When the NPC spawns.

This event is often used to start timers, attack player targets, establish NPC HP events or proximities, start dialogues, and more.

#### Example

* In this example, when the mob spawn, we make it run and attack a nearby player while it shouts a war cry.


=== "Perl"
    ```perl
    sub EVENT_SPAWN {
        #:: Set the NPC to run
        quest::SetRunning(1);
        #:: Shout out a war cray
        quest::shout("For Jotenheimr!!!");
        #:: Try to find a random client within 200 units of the NPC
        my $rClient = $entity_list->GetRandomClient($x,$y,$z, 200);
        #:: If there's a random sucker nearby, attack them
        if ($rClient) {
            quest::attack($rClient->GetName());
        }
    }
    ```


=== "Lua"
    ```lua
    function event_spawn(e)
        --:: Set the NPC to run
        e.self:SetRunning(1);
        --:: Shout out a war cray
        e.self:Shout("For Jotenheimr!!!");
        --:: Try to find a random client within 200 units of the NPC
        local random_client = entity_list:GetRandomClient(e.self:GetX(), e.self:GetY(), e.self:GetZ(), 200);
        --:: If there's a random sucker nearby, attack them
        if (random_client ~= nil) then
            e.self:Attack(random_client);
        end
    end
    ```



### EVENT_SPAWN_ZONE

#### Trigger

* When an NPC spawns.

**Used by the zone controller.**

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| spawned_entity_id | int | `quest::say($spawned_entity_id); # returns int` |
| spawned_npc_id | int | `quest::say($spawned_npc_id); # returns int` |

### EVENT_SPELL_EFFECT_CLIENT

#### Trigger

* When the spell lands on a client.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| caster_id | int | `quest::say($caster_id); # returns int` |

### EVENT_SPELL_EFFECT_NPC

#### Trigger

* When the spell lands on an NPC.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| caster_id | int | `quest::say($caster_id); # returns int` |

### EVENT_SPELL_EFFECT_BUFF_TIC_CLIENT

#### Trigger

* When the spell ticks on a client.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| caster_id | int | `quest::say($caster_id); # returns int` |

### EVENT_SPELL_EFFECT_BUFF_TIC_NPC

#### Trigger

* When the spell ticks on an NPC.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| caster_id | int | `quest::say($caster_id); # returns int` |

### EVENT_SPELL_EFFECT_TRANSLOCATE_COMPLETE

#### Trigger

* when translocation is complete.

### EVENT_SPELL_FADE

#### Trigger

* when a spell fades.

### EVENT_TARGET_CHANGE

#### Trigger

* when a mob changes their current target or clears it.

### EVENT_TASKACCEPTED

#### Trigger

* when a player accepts a task from the task selector window.

Typically you would handle this functionality using the task system.

### EVENT_TASK_COMPLETE

#### Trigger

* when a player completes a task.

Typically you would handle this functionality using the task system.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| donecount | int | `quest::say($donecount); # returns int` |
| activity_id | int | `quest::say($activity_id); # returns int` |
| task_id | int | `quest::say($task_id); # returns int` |

### EVENT_TASK_FAIL

#### Trigger

* When a player fails a task.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| task_id | int | `quest::say($task_id); # returns int` |

Typically you would handle this functionality using the task system.

### EVENT_TASK_STAGE_COMPLETE

#### Trigger

* When a task stage is completed.

Typically you would handle this functionality using the task system.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| activity_id | int | `quest::say($activity_id); # returns int` |
| task_id | int | `quest::say($task_id); # returns int` |

#### Example

* In this example, when a player completes a task, it triggers the event and tries to match with task ID 212; if it matches, a yellow message is displayed to the client.


=== "Perl"
    ```perl
    sub EVENT_TASK_STAGE_COMPLETE {
        #:: Match task id 212
        if ($task_id == 212) {
            $client->Message(15,"The zombie presence seems somewhat lessened, and perhaps they have been quelled...for the time being.");
        }
    }
    ```


=== "Lua"
    ```lua
    function event_task_stage_complete(e)
        --:: Match task id 212
        if (e.task_id == 212) then
            e.self:Message(15, "The zombie presence seems somewhat lessened, and perhaps they have been quelled...for the time being.");
        end
    end
    ```



### EVENT_TASK_UPDATE

#### Trigger

* when a player's task is updated.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| donecount | int | `quest::say($donecount); # returns int` |
| activity_id | int | `quest::say($activity_id); # returns int` |
| task_id | int | `quest::say($task_id); # returns int` |

### EVENT_TIMER

#### Trigger

* by a quest::settimer(timer_name,duration_in_seconds) or quest::settimerMS(timer_name,duration_in_milliseconds)

The timer will loop until it is stopped, and EVENT_TIMER will trigger each time that the duration of the timer elapses. Timers can be stopped using the quest::stopalltimers() or quest::stoptimer(timer_name) functions.

#### Exports

| Name | Type | Usage |
| :--- | :--- | :--- |
| timer | int | `quest::say($timer); # returns int` |

#### Example

* This is an example of using a timer with a string name to cause an NPC to depop 30 minutes after it spawns


=== "Perl"
    ```perl
    sub EVENT_SPAWN {
        #:: Start a timer that is named "depop", the duration is 1,800 seconds (30 minutes)
        quest::settimer("depop",1800); 
    }
    
    sub EVENT_TIMER {
        #:: Use eq for string comparison to match timer "depop"
        if ($timer eq "depop") {
            # Stop timer "depop" from looping
            quest::stoptimer("depop"); 
            quest::depop(); 
        }
    }
    ```


=== "Lua"
    ```lua
    function event_spawn(e) {
        --:: Start a timer that is named "depop", the duration is 1,800 seconds (30 minutes)
        eq.set_timer("depop", 1800); 
    end
    
    function event_timer(e) {
        if (e.timer == "depop") then
            --:: Stop timer "depop" from looping
            eq.stop_timer("depop"); 
            eq.depop(); 
        end
    end
    ```
    


### EVENT_TRADE

#### Trigger

* When a client begins a trade.

### EVENT_UNAUGMENT_ITEM

#### Trigger

* When a client removes an augment from an item.

### EVENT_UNEQUIP_ITEM

#### Trigger

* When a client unequips an item.

### EVENT_USE_SKILL

#### Trigger

* When a player uses a skill

#### Exports

| Name | Type | Description |
| :--- | :--- | :--- |
| skill_id | int | `quest::say($skill_id); # returns int` |
| skill_level | int | `quest::say($skill_level); # returns int` |

#### Example

This example would display a message to the client who uses kick and has a high enough skill level.


=== "Perl"
    ```perl
    sub EVENT_USE_SKILL {
        #:: Match if skill is Kick (30) and Skill Level is greater than or equal to 100.
        if ($skill_id == 30 && $skill_level >= 100) {
            $client->Message(315, "You have used Kick. Your skill level is $skill_level");
        }
    }
    ```


=== "Lua"
    ```lua
    function event_use_skill(e)
        --:: Match if skill is Kick (30) and Skill Level is greater than or equal to 100.
        if (e.skill_id == 30 && e.skill_level >= 100) then
            e.self:Message(315, "You have used Kick. Your skill level is " .. skill_level .. ".");
        end
    end
    ```



### EVENT_WAYPOINT_ARRIVE

#### Trigger

* When an NPC arrives at a grid waypoint entry.

#### Exports

| Name | Type | Description |
| :--- | :--- | :--- |
| wp | int | `quest::say($wp); # returns int` |

#### Example

* This example would cause your NPC to speak as it arrives at a particular waypoint.
* Don't forget to count waypoint 0 (the spawn point) when using waypoint events.


=== "Perl"
    ```perl
    sub EVENT_WAYPOINT_ARRIVE {
        #:: If we have arrived at waypoint 10
        if ($wp == 10) {
            quest::say("We're finally here!");
        }
    }
    ```


=== "Lua"
    ```lua
    function event_waypoint_arrive(e)
        --:: If we have arrived at waypoint 10
        if (e.wp == 10) then
            e.self:Say("We're finally here!");
        end
    end
    ```



### EVENT_WAYPOINT_DEPART

#### Trigger

* When an NPC leaves its current grid waypoint entry.

#### Exports

| Name | Type | Description |
| :--- | :--- | :--- |
| wp | int | `quest::say($wp); # returns int` |

#### Example

* This example would cause your NPC to speak as it departs a particular waypoint.
* Don't forget to count waypoint 0 (the spawn point) when using waypoint events.


=== "Perl"
    ```perl
    sub EVENT_WAYPOINT_DEPART {
        # If we departed waypoint 0
        if ($wp == 0) {
            quest::say("And we're off!");
        }
    }
    ```


=== "Lua"
    ```lua
    function event_waypoint_depart(e)
        --:: If we departed waypoint 0
        if (e.wp == 0) then
            e.self:Say("And we're off!");
        end
    end
    ```



### EVENT_WEAPON_PROC

#### Trigger

* When a weapon procs.

### EVENT_ZONE

#### Trigger

* When a player leaves a zone.

#### Exports

| Name | Type | Description |
| :--- | :--- | :--- |
| target_zone_id | int | `quest::say($target_zone_id); # returns int` |

#### Example

* In this example, we blow up pets when a player with a pet zones.


=== "Perl"
    ```perl
    sub EVENT_ZONE {
        #:: Match if a player has a pet
        if ($client->GetPetID()) {
            #:: Get the pet's ID and kill it
            $PetID = $entity_list->GetMobByID($client->GetPetID());
            $PetID->Kill();
        }
    }
    ```


=== "Lua"
    ```lua
    function event_zone(e)
        --:: See if the player has a pet
        local pet_id = e.other:GetPet():GetID();
        if (pet_id > 0) then
            --:: Identify the pet by ID and kill it
            local pet_entity = entity_list:GetMob(pet_id);
            pet_entity:Kill();
        end
    end
    ```
    

