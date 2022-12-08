=== "Perl (27)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Bot){:target="Bot"} for latest definitions and Quest examples

        Last generated 2022.12.07

    ``` perl
    $bot->AddBotItem(uint16 slot_id, uint32 item_id, uint16 charges, bool attuned, uint32 aug1, uint32 aug2, uint32 aug3, uint32 aug4, uint32 aug5, uint32 aug6);
    $bot->AddBotItem(uint16 slot_id, uint32 item_id);
    $bot->AddBotItem(uint16 slot_id, uint32 item_id, uint16 charges);
    $bot->AddBotItem(uint16 slot_id, uint32 item_id, uint16 charges, bool attuned);
    $bot->AddBotItem(uint16 slot_id, uint32 item_id, uint16 charges, bool attuned, uint32 aug1);
    $bot->AddBotItem(uint16 slot_id, uint32 item_id, uint16 charges, bool attuned, uint32 aug1, uint32 aug2);
    $bot->AddBotItem(uint16 slot_id, uint32 item_id, uint16 charges, bool attuned, uint32 aug1, uint32 aug2, uint32 aug3);
    $bot->AddBotItem(uint16 slot_id, uint32 item_id, uint16 charges, bool attuned, uint32 aug1, uint32 aug2, uint32 aug3, uint32 aug4);
    $bot->AddBotItem(uint16 slot_id, uint32 item_id, uint16 charges, bool attuned, uint32 aug1, uint32 aug2, uint32 aug3, uint32 aug4, uint32 aug5);
    $bot->CountBotItem(item_id);
    $bot->GetBotItem(uint16 slot_id);
    $bot->GetBotItemIDBySlot(uint16 slot_id);
    $bot->GetExpansionBitmask();
    $bot->GetOwner();
    $bot->HasBotItem(uint32 item_id);
    $bot->HasBotSpellEntry(uint16 spellid);
    $bot->OwnerMessage(string message);
    $bot->ReloadBotDataBuckets();
    $bot->ReloadBotOwnerDataBuckets();
    $bot->ReloadBotSpellSettings();
    $bot->ReloadBotSpells();
    $bot->RemoveBotItem(uint32 item_id);
    $bot->SendPayload(int payload_id);
    $bot->SendPayload(int payload_id, string payload_value);
    $bot->SetExpansionBitmask(int expansion_bitmask, bool save);
    $bot->SetExpansionBitmask(int expansion_bitmask);
    $bot->Signal(int signal_id);
    ```
=== "Lua (27)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Bot){:target="Bot"} for latest definitions and Quest examples

        Last generated 2022.12.07

    ``` lua
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges, bool attuned, uint32 augment_one, uint32 augment_two, uint32 augment_three, uint32 augment_four);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges, bool attuned);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges, bool attuned, uint32 augment_one);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges, bool attuned, uint32 augment_one, uint32 augment_two);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges, bool attuned, uint32 augment_one, uint32 augment_two, uint32 augment_three);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges, bool attuned, uint32 augment_one, uint32 augment_two, uint32 augment_three, uint32 augment_four, uint32 augment_five);
    bot:AddBotItem(uint16 slot_id, uint32 item_id, int16 charges, bool attuned, uint32 augment_one, uint32 augment_two, uint32 augment_three, uint32 augment_four, uint32 augment_five, uint32 augment_six);
    bot:AddBotItem(uint16 slot_id, uint32 item_id);
    bot:CountBotItem(item_id);
    bot:GetBotItem(uint16 slot_id);
    bot:GetBotItemIDBySlot(uint16 slot_id);
    bot:GetExpansionBitmask();
    bot:GetOwner();
    bot:HasBotItem(uint32 item_id);
    bot:HasBotSpellEntry(uint16 spellid);
    bot:OwnerMessage(string message);
    bot:ReloadBotDataBuckets();
    bot:ReloadBotOwnerDataBuckets();
    bot:ReloadBotSpellSettings();
    bot:ReloadBotSpells();
    bot:RemoveBotItem(uint32 item_id);
    bot:SendPayload(int payload_id);
    bot:SendPayload(int payload_id, string payload_value);
    bot:SetExpansionBitmask(int expansion_bitmask, bool save);
    bot:SetExpansionBitmask(int expansion_bitmask);
    bot:Signal(int signal_id);
    ```
