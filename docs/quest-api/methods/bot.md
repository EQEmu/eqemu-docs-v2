=== "Perl (4)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Bot){:target="Bot"} for latest definitions and Quest examples

        Last generated 2022.05.11

    ``` perl
    $bot->AddBotItem(uint16 slot_id, uint32 item_id, [int16 charges = -1], [bool attuned = false], [uint32 augment_one = 0], [uint32 augment_two = 0], [uint32 augment_three = 0], [uint32 augment_four = 0], [uint32 augment_five = 0], [uint32 augment_six = 0]);
    $bot->CountBotItem(uint32 item_id);
    $bot->GetOwner();
    $bot->RemoveBotItem(uint32 item_id);
    ```
=== "Lua (13)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Bot){:target="Bot"} for latest definitions and Quest examples

        Last generated 2022.05.11

    ``` lua
    bot:AddBotItem(uint16 slot_id, uint32 item_id);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges, bool attuned);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges, bool attuned, uint32 augment_one);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges, bool attuned, uint32 augment_one, uint32 augment_two);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges, bool attuned, uint32 augment_one, uint32 augment_two, uint32 augment_three);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges, bool attuned, uint32 augment_one, uint32 augment_two, uint32 augment_three, uint32 augment_four);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges, bool attuned, uint32 augment_one, uint32 augment_two, uint32 augment_three, uint32 augment_four, uint32 augment_five);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges, bool attuned, uint32 augment_one, uint32 augment_two, uint32 augment_three, uint32 augment_four, uint32 augment_five, uint32 augment_six);
    bot:CountBotItem(item_id);
    bot:GetOwner();
    bot:HasBotItem(uint32 item_id);
    bot:RemoveBotItem(uint32 item_id);
    ```
