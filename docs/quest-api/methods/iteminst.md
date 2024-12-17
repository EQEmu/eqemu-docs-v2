=== "Lua (56)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=ItemInst){:target="ItemInst"} for latest definitions and Quest examples

    ``` lua
    iteminst:AddExp(uint32 exp);
    iteminst:ClearTimers();
    iteminst:Clone();
    iteminst:ContainsAugmentByID(uint32 item_id);
    iteminst:CountAugmentByID(uint32 item_id);
    iteminst:DeleteCustomData(const std);
    iteminst:GetAugment(uint8 slot_id);
    iteminst:GetAugmentItemID(uint8 slot_id);
    iteminst:GetAugmentType();
    iteminst:GetCharges();
    iteminst:GetColor();
    iteminst:GetCustomData(const std);
    iteminst:GetCustomDataString();
    iteminst:GetExp();
    iteminst:GetID();
    iteminst:GetItem();
    iteminst:GetItem(uint8 slot_id);
    iteminst:GetItemID(uint8 slot_id);
    iteminst:GetItemLink();
    iteminst:GetItemScriptID();
    iteminst:GetKillsNeeded(uint8 current_level);
    iteminst:GetMaxEvolveLvl();
    iteminst:GetName();
    iteminst:GetPrice();
    iteminst:GetTaskDeliveredCount();
    iteminst:GetTotalItemCount();
    iteminst:GetUnscaledItem();
    iteminst:IsAmmo();
    iteminst:IsAttuned();
    iteminst:IsAugmentable();
    iteminst:IsAugmented();
    iteminst:IsEquipable(uint16 race_bitmask, uint16 class_bitmask);
    iteminst:IsEquipable(int16 slot_id);
    iteminst:IsExpendable();
    iteminst:IsInstNoDrop();
    iteminst:IsStackable();
    iteminst:IsType(int item_class);
    iteminst:IsWeapon();
    iteminst:ItemSay(string text // @categories Inventory and Items);
    iteminst:ItemSay(string text, uint8 language_id // @categories Inventory and Items);
    iteminst:RemoveTaskDeliveredItems();
    iteminst:SetAttuned(bool flag);
    iteminst:SetCharges(int charges);
    iteminst:SetColor(uint32 color);
    iteminst:SetCustomData(const string &identifier, std);
    iteminst:SetCustomData(const std);
    iteminst:SetCustomData(const std);
    iteminst:SetCustomData(const std);
    iteminst:SetExp(uint32 exp);
    iteminst:SetInstNoDrop(bool flag);
    iteminst:SetPrice(uint32 price);
    iteminst:SetScale(double scale_factor);
    iteminst:SetScaling(bool v);
    iteminst:SetTimer(string name, uint32 time);
    iteminst:StopTimer(string name);
    iteminst:operator=(const o);
    ```
