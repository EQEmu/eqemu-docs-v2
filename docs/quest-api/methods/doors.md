=== "Perl (51)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Doors){:target="Doors"} for latest definitions and Quest examples

    ``` perl
    $doors->CreateDatabaseEntry();
    $doors->ForceClose(Mob* sender);
    $doors->ForceClose(Mob* sender, bool alt_mode);
    $doors->ForceOpen(Mob* sender);
    $doors->ForceOpen(Mob* sender, bool alt_mode);
    $doors->GetClientVersionMask();
    $doors->GetDestinationHeading();
    $doors->GetDestinationInstanceID();
    $doors->GetDestinationX();
    $doors->GetDestinationY();
    $doors->GetDestinationZ();
    $doors->GetDestinationZoneName();
    $doors->GetDisableTimer();
    $doors->GetDoorDBID();
    $doors->GetDoorID();
    $doors->GetDoorParam();
    $doors->GetDzSwitchID();
    $doors->GetGuildID();
    $doors->GetHeading();
    $doors->GetID();
    $doors->GetIncline();
    $doors->GetInvertState();
    $doors->GetKeyItem();
    $doors->GetLockPick();
    $doors->GetModelName();
    $doors->GetNoKeyring();
    $doors->GetOpenType();
    $doors->GetSize();
    $doors->GetTriggerDoorID();
    $doors->GetTriggerType();
    $doors->GetX();
    $doors->GetY();
    $doors->GetZ();
    $doors->HasDestinationZone();
    $doors->IsDestinationZoneSame();
    $doors->IsDoorBlacklisted();
    $doors->IsLDoNDoor();
    $doors->SetDisableTimer(bool disable_timer);
    $doors->SetHeading(float heading);
    $doors->SetIncline(uint32_t incline);
    $doors->SetInvertState(int invert_state);
    $doors->SetKeyItem(uint32_t key_item_id);
    $doors->SetLocation(float x, float y, float z);
    $doors->SetLockPick(uint32_t lockpick_type);
    $doors->SetModelName(string name);
    $doors->SetNoKeyring(uint8_t no_key_ring);
    $doors->SetOpenType(uint32_t open_type);
    $doors->SetSize(uint32_t size);
    $doors->SetX(float x);
    $doors->SetY(float y);
    $doors->SetZ(float z);
    ```
