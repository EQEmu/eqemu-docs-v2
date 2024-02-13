=== "Perl (33)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Doors){:target="Doors"} for latest definitions and Quest examples

    ``` perl
    $doors->CreateDatabaseEntry();
    $doors->ForceClose(Mob* sender);
    $doors->ForceClose(Mob* sender, bool alt_mode);
    $doors->ForceOpen(Mob* sender);
    $doors->ForceOpen(Mob* sender, bool alt_mode);
    $doors->GetDisableTimer();
    $doors->GetDoorDBID();
    $doors->GetDoorID();
    $doors->GetHeading();
    $doors->GetID();
    $doors->GetIncline();
    $doors->GetKeyItem();
    $doors->GetLockPick();
    $doors->GetModelName();
    $doors->GetNoKeyring();
    $doors->GetOpenType();
    $doors->GetSize();
    $doors->GetX();
    $doors->GetY();
    $doors->GetZ();
    $doors->SetDisableTimer(bool disable_timer);
    $doors->SetHeading(float heading);
    $doors->SetIncline(uint32_t incline);
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
