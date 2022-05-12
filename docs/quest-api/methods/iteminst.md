=== "Lua (48)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=ItemInst){:target="ItemInst"} for latest definitions and Quest examples

        Last generated 2022.05.11

    ``` lua
    iteminst:AddExp(uint32 exp);
    iteminst:ClearTimers();
    iteminst:Clone();
    iteminst:ContainsAugmentByID(uint32 item_id);
    iteminst:CountAugmentByID(uint32 item_id);
    iteminst:DeleteCustomData(string identifier);
    iteminst:GetAugment(int slot);
    iteminst:GetAugmentItemID(int slot);
    iteminst:GetAugmentType();
    iteminst:GetCharges();
    iteminst:GetColor();
    iteminst:GetCustomData(identifier);
    iteminst:GetCustomDataString();
    iteminst:GetExp();
    iteminst:GetID();
    iteminst:GetItem();
    iteminst:GetItem(int slot);
    iteminst:GetItemID(int slot);
    iteminst:GetItemScriptID();
    iteminst:GetKillsNeeded(int current_level);
    iteminst:GetMaxEvolveLvl();
    iteminst:GetPrice();
    iteminst:GetTotalItemCount();
    iteminst:GetUnscaledItem(int slot);
    iteminst:IsAmmo();
    iteminst:IsAugmentable();
    iteminst:IsAugmented();
    iteminst:IsEquipable(int race, int class_);
    iteminst:IsEquipable(int slot_id);
    iteminst:IsExpendable();
    iteminst:IsInstNoDrop();
    iteminst:IsStackable();
    iteminst:IsType(int item_class);
    iteminst:IsWeapon();
    iteminst:SetCharges(int charges);
    iteminst:SetColor(uint32 color);
    iteminst:SetCustomData(string identifier, string value);
    iteminst:SetCustomData(string identifier, float value);
    iteminst:SetCustomData(string identifier, int value);
    iteminst:SetCustomData(string identifier, bool value);
    iteminst:SetExp(uint32 exp);
    iteminst:SetInstNoDrop(bool flag);
    iteminst:SetPrice(uint32 price);
    iteminst:SetScale(double scale_factor);
    iteminst:SetScaling(bool v);
    iteminst:SetTimer(string name, uint32 time);
    iteminst:StopTimer(string name);
    iteminst:operator=(const o);
    ```
