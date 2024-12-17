=== "Perl (57)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=QuestItem){:target="QuestItem"} for latest definitions and Quest examples

    ``` perl
    $questitem->AddEXP(uint32 exp);
    $questitem->ClearTimers();
    $questitem->Clone();
    $questitem->ContainsAugmentByID(uint32_t item_id);
    $questitem->CountAugmentByID(uint32_t item_id);
    $questitem->DeleteCustomData(string identifier);
    $questitem->GetAugment(uint8 slot_id);
    $questitem->GetAugmentIDs();
    $questitem->GetAugmentItemID(uint8 slot_id);
    $questitem->GetAugmentType();
    $questitem->GetCharges();
    $questitem->GetColor();
    $questitem->GetComment();
    $questitem->GetCustomData(identifier);
    $questitem->GetCustomDataString();
    $questitem->GetEXP();
    $questitem->GetID();
    $questitem->GetItem(uint8 slot_id);
    $questitem->GetItem();
    $questitem->GetItemID(uint8 slot_id);
    $questitem->GetItemLink();
    $questitem->GetItemScriptID();
    $questitem->GetKillsNeeded(uint8 current_level);
    $questitem->GetMaxEvolveLevel();
    $questitem->GetName();
    $questitem->GetPrice();
    $questitem->GetTaskDeliveredCount();
    $questitem->GetTotalItemCount();
    $questitem->GetUnscaledItem();
    $questitem->IsAmmo();
    $questitem->IsAttuned();
    $questitem->IsAugmentable();
    $questitem->IsAugmented();
    $questitem->IsEquipable(int16 slot_id);
    $questitem->IsEquipable(uint16 race_bitmask, uint16 class_bitmask);
    $questitem->IsExpendable();
    $questitem->IsInstanceNoDrop();
    $questitem->IsStackable();
    $questitem->IsType(int type);
    $questitem->IsWeapon();
    $questitem->ItemSay(string text);
    $questitem->ItemSay(string text, uint8 language_id);
    $questitem->RemoveTaskDeliveredItems();
    $questitem->SetAttuned(bool is_attuned);
    $questitem->SetCharges(int16_t charges);
    $questitem->SetColor(uint32 color);
    $questitem->SetCustomData(string identifier, bool value);
    $questitem->SetCustomData(string identifier, string value);
    $questitem->SetCustomData(string identifier, int value);
    $questitem->SetCustomData(string identifier, float value);
    $questitem->SetEXP(uint32 exp);
    $questitem->SetInstanceNoDrop(bool is_attuned);
    $questitem->SetPrice(uint32 price);
    $questitem->SetScale(float scale_multiplier);
    $questitem->SetScaling(bool is_scaling);
    $questitem->SetTimer(string timer_name, uint32 timer);
    $questitem->StopTimer(string timer_name);
    ```
