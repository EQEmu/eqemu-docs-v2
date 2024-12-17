=== "Lua (51)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Door){:target="Door"} for latest definitions and Quest examples

    ``` lua
    door:CreateDatabaseEntry();
    door:ForceClose(Mob sender, bool alt_mode);
    door:ForceClose(Mob sender);
    door:ForceOpen(Mob sender);
    door:ForceOpen(Mob sender, bool alt_mode);
    door:GetClientVersionMask();
    door:GetDestinationHeading();
    door:GetDestinationInstanceID();
    door:GetDestinationX();
    door:GetDestinationY();
    door:GetDestinationZ();
    door:GetDestinationZoneName();
    door:GetDisableTimer();
    door:GetDoorDBID();
    door:GetDoorID();
    door:GetDoorName();
    door:GetDoorParam();
    door:GetDzSwitchID();
    door:GetGuildID();
    door:GetHeading();
    door:GetID();
    door:GetIncline();
    door:GetInvertState();
    door:GetKeyItem();
    door:GetLockPick();
    door:GetNoKeyring();
    door:GetOpenType();
    door:GetSize();
    door:GetTriggerDoorID();
    door:GetTriggerType();
    door:GetX();
    door:GetY();
    door:GetZ();
    door:HasDestinationZone();
    door:IsDestinationZoneSame();
    door:IsDoorBlacklisted();
    door:IsLDoNDoor();
    door:SetDisableTimer(bool flag);
    door:SetDoorName(string name);
    door:SetHeading(float h);
    door:SetIncline(uint32 incline);
    door:SetInvertState(int invert_state);
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
