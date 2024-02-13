=== "Perl (26)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Spawn){:target="Spawn"} for latest definitions and Quest examples

    ``` perl
    $spawn->Depop();
    $spawn->Disable();
    $spawn->Enable();
    $spawn->ForceDespawn();
    $spawn->GetCurrentNPCID();
    $spawn->GetHeading();
    $spawn->GetID();
    $spawn->GetKillCount();
    $spawn->GetRespawnTimer();
    $spawn->GetSpawnCondition();
    $spawn->GetSpawnGroupID();
    $spawn->GetVariance();
    $spawn->GetX();
    $spawn->GetY();
    $spawn->GetZ();
    $spawn->IsEnabled();
    $spawn->IsNPCPointerValid();
    $spawn->LoadGrid();
    $spawn->Repop();
    $spawn->Repop(uint32 delay);
    $spawn->Reset();
    $spawn->SetCurrentNPCID(uint32 npc_id);
    $spawn->SetNPCPointer(NPC* n);
    $spawn->SetRespawnTimer(uint32 new_respawn_time);
    $spawn->SetTimer(uint32 duration);
    $spawn->SetVariance(uint32 new_variance);
    ```
=== "Lua (26)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Spawn){:target="Spawn"} for latest definitions and Quest examples

    ``` lua
    spawn:CurrentNPCID();
    spawn:Depop();
    spawn:Disable();
    spawn:Enable();
    spawn:Enabled();
    spawn:ForceDespawn();
    spawn:GetHeading();
    spawn:GetID();
    spawn:GetKillCount();
    spawn:GetSpawnCondition();
    spawn:GetVariance();
    spawn:GetX();
    spawn:GetY();
    spawn:GetZ();
    spawn:LoadGrid();
    spawn:NPCPointerValid();
    spawn:Repop();
    spawn:Repop(uint32 delay);
    spawn:Reset();
    spawn:RespawnTimer();
    spawn:SetCurrentNPCID(uint32 nid);
    spawn:SetNPCPointer(NPC n);
    spawn:SetRespawnTimer(uint32 newrespawntime);
    spawn:SetTimer(uint32 duration);
    spawn:SetVariance(uint32 newvariance);
    spawn:SpawnGroupID();
    ```
