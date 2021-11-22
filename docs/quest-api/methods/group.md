=== "Perl"

    ``` perl
    $group->CastGroupSpell(Mob* caster, uint16 spell_id);
    $group->DisbandGroup();
    $group->DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name, [int max_check_count = 0]);
    $group->GetHighestLevel();
    $group->GetID();
    $group->GetLeader();
    $group->GetLeaderName();
    $group->GetMember(int group_index);
    $group->GetTotalGroupDamage(Mob* other);
    $group->GroupCount();
    $group->GroupMessage(Mob* sender, uint8 language, string message);
    $group->IsGroupMember(client);
    $group->IsLeader(Mob* target);
    $group->SendHPPacketsFrom(Mob* new_member);
    $group->SendHPPacketsTo(Mob* new_member);
    $group->SetLeader(Mob* new_leader);
    $group->SplitExp(uint32 exp, Mob* other);
    $group->SplitMoney(uint32 copper, uint32 silver, uint32 gold, uint32 platinum);
    $group->TeleportGroup(Mob* sender, uint32 zone_id, float x, float y, float z, float heading);
    ```
=== "Lua"

    ``` lua
    group:CastGroupSpell(Mob caster, int spell_id);
    group:DisbandGroup();
    group:DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name);
    group:DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name, int max_check_count);
    group:GetHighestLevel();
    group:GetID();
    group:GetLeader();
    group:GetLeaderName();
    group:GetLowestLevel();
    group:GetMember(int index);
    group:GetTotalGroupDamage(Mob other);
    group:GroupCount();
    group:GroupMessage(Mob sender, int language, string message);
    group:IsGroupMember(Mob mob);
    group:IsLeader(Mob leader);
    group:SetLeader(Mob leader);
    group:SplitExp(uint32 exp, Mob other);
    group:SplitMoney(uint32 copper, uint32 silver, uint32 gold, uint32 platinum, Client splitter);
    group:SplitMoney(uint32 copper, uint32 silver, uint32 gold, uint32 platinum);
    group:TeleportGroup(Mob sender, uint32 zone_id, uint32 instance_id, float x, float y, float z, float h);
    ```
