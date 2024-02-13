=== "Perl (39)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Expedition){:target="Expedition"} for latest definitions and Quest examples

    ``` perl
    $expedition->AddLockout(string event_name, uint32_t seconds);
    $expedition->AddLockoutDuration(string event_name, int seconds);
    $expedition->AddLockoutDuration(string event_name, int seconds, bool members_only);
    $expedition->AddReplayLockout(uint32_t seconds);
    $expedition->AddReplayLockoutDuration(int seconds);
    $expedition->AddReplayLockoutDuration(int seconds, bool members_only);
    $expedition->GetDynamicZoneID();
    $expedition->GetID();
    $expedition->GetInstanceID();
    $expedition->GetLeaderName();
    $expedition->GetLockouts();
    $expedition->GetLootEventByNPCTypeID(uint32_t npc_type_id);
    $expedition->GetLootEventBySpawnID(uint32_t spawn_id);
    $expedition->GetMemberCount();
    $expedition->GetMembers();
    $expedition->GetName();
    $expedition->GetSecondsRemaining();
    $expedition->GetUUID();
    $expedition->GetZoneID();
    $expedition->GetZoneName();
    $expedition->GetZoneVersion();
    $expedition->HasLockout(string event_name);
    $expedition->HasReplayLockout();
    $expedition->IsLocked();
    $expedition->RemoveCompass();
    $expedition->RemoveLockout(string event_name);
    $expedition->SetCompass(scalar zone, float x, float y, float z);
    $expedition->SetLocked(bool locked);
    $expedition->SetLocked(bool locked, int lock_msg);
    $expedition->SetLocked(bool locked, int lock_msg, uint32_t color);
    $expedition->SetLootEventByNPCTypeID(uint32_t npc_type_id, string event_name);
    $expedition->SetLootEventBySpawnID(uint32_t entity_id, string event_name);
    $expedition->SetReplayLockoutOnMemberJoin(bool enable);
    $expedition->SetSafeReturn(scalar zone, float x, float y, float z, float heading);
    $expedition->SetSecondsRemaining(uint32_t seconds_remaining);
    $expedition->SetSwitchID(int dz_switch_id);
    $expedition->SetZoneInLocation(float x, float y, float z, float heading);
    $expedition->UpdateLockoutDuration(string event_name, uint32_t seconds);
    $expedition->UpdateLockoutDuration(string event_name, uint32_t seconds, bool members_only);
    ```
=== "Lua (41)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Expedition){:target="Expedition"} for latest definitions and Quest examples

    ``` lua
    expedition:AddLockout(string event_name, uint32_t seconds);
    expedition:AddLockoutDuration(string event_name, int seconds);
    expedition:AddLockoutDuration(string event_name, int seconds, bool members_only);
    expedition:AddReplayLockout(uint32_t seconds);
    expedition:AddReplayLockoutDuration(int seconds);
    expedition:AddReplayLockoutDuration(int seconds, bool members_only);
    expedition:GetDynamicZoneID();
    expedition:GetID();
    expedition:GetInstanceID();
    expedition:GetLeaderName();
    expedition:GetLockouts();
    expedition:GetLootEventByNPCTypeID(uint32_t npc_type_id);
    expedition:GetLootEventBySpawnID(uint32_t spawn_id);
    expedition:GetMemberCount();
    expedition:GetMembers();
    expedition:GetName();
    expedition:GetSecondsRemaining();
    expedition:GetUUID();
    expedition:GetZoneID();
    expedition:GetZoneName();
    expedition:GetZoneVersion();
    expedition:HasLockout(string event_name);
    expedition:HasReplayLockout();
    expedition:IsLocked();
    expedition:RemoveCompass();
    expedition:RemoveLockout(string event_name);
    expedition:SetCompass(uint32_t zone_id, float x, float y, float z);
    expedition:SetCompass(string zone_name, float x, float y, float z);
    expedition:SetLocked(bool lock_expedition);
    expedition:SetLocked(bool lock_expedition, int lock_msg);
    expedition:SetLocked(bool lock_expedition, int lock_msg, uint32_t msg_color);
    expedition:SetLootEventByNPCTypeID(uint32_t npc_type_id, string event_name);
    expedition:SetLootEventBySpawnID(uint32_t spawn_id, string event_name);
    expedition:SetReplayLockoutOnMemberJoin(bool enable);
    expedition:SetSafeReturn(uint32_t zone_id, float x, float y, float z, float heading);
    expedition:SetSafeReturn(string zone_name, float x, float y, float z, float heading);
    expedition:SetSecondsRemaining(uint32_t seconds_remaining);
    expedition:SetSwitchID(int dz_switch_id);
    expedition:SetZoneInLocation(float x, float y, float z, float heading);
    expedition:UpdateLockoutDuration(string event_name, uint32_t duration);
    expedition:UpdateLockoutDuration(string event_name, uint32_t duration, bool members_only);
    ```
