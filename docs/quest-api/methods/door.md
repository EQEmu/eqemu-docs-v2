=== "Lua"

    ``` lua
    door:CreateDatabaseEntry();
    door:ForceClose(Mob sender);
    door:ForceClose(Mob sender, bool alt_mode);
    door:ForceOpen(Mob sender, bool alt_mode);
    door:ForceOpen(Mob sender);
    door:GetDisableTimer();
    door:GetDoorDBID();
    door:GetDoorID();
    door:GetDoorName();
    door:GetHeading();
    door:GetIncline();
    door:GetKeyItem();
    door:GetLockPick();
    door:GetNoKeyring();
    door:GetOpenType();
    door:GetSize();
    door:GetX();
    door:GetY();
    door:GetZ();
    door:SetDisableTimer(bool flag);
    door:SetDoorName(string name);
    door:SetHeading(float h);
    door:SetIncline(uint32 incline);
    door:SetKeyItem(uint32 key);
    door:SetLocation(float x, float y, float z);
    door:SetLockPick(uint32 pick);
    door:SetNoKeyring(int type);
    door:SetOpenType(uint32 type);
    door:SetSize(uint32 sz);
    door:SetX(float x);
    door:SetY(float y);
    door:SetZ(float z);
    ```
