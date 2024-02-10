=== "Perl (26)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Group){:target="Group"} for latest definitions and Quest examples

    ``` perl
    $group->CastGroupSpell(Mob* caster, uint16 spell_id);
    $group->DisbandGroup();
    $group->DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name, int max_check_count);
    $group->DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name);
    $group->GetAverageLevel();
    $group->GetHighestLevel();
    $group->GetID();
    $group->GetLeader();
    $group->GetLeaderName();
    $group->GetLowestLevel();
    $group->GetMember(int member_index);
    $group->GetTotalGroupDamage(Mob* other);
    $group->GroupCount();
    $group->GroupMessage(Mob* sender, uint8_t language, string message);
    $group->GroupMessage(Mob* sender, string message);
    $group->IsGroupMember(Mob* c);
    $group->IsGroupMember(string name);
    $group->IsLeader(Mob* c);
    $group->IsLeader(string name);
    $group->SendHPPacketsFrom(Mob* new_member);
    $group->SendHPPacketsTo(Mob* new_member);
    $group->SetLeader(Mob* new_leader);
    $group->SplitExp(uint32_t exp, Mob* other);
    $group->SplitMoney(uint32 copper, uint32 silver, uint32 gold, uint32 platinum);
    $group->SplitMoney(uint32 copper, uint32 silver, uint32 gold, uint32 platinum, Client* splitter);
    $group->TeleportGroup(Mob* sender, uint32 zone_id, float x, float y, float z, float heading);
    ```
=== "Lua (24)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Group){:target="Group"} for latest definitions and Quest examples

    ``` lua
    group:CastGroupSpell(Mob caster, int spell_id);
    group:DisbandGroup();
    group:DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name, int max_check_count);
    group:DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name);
    group:GetAverageLevel();
    group:GetHighestLevel();
    group:GetID();
    group:GetLeader();
    group:GetLeaderName();
    group:GetLowestLevel();
    group:GetMember(int member_index);
    group:GetTotalGroupDamage(Mob other);
    group:GroupCount();
    group:GroupMessage(Mob sender, uint8 language, string message);
    group:GroupMessage(Mob sender, string message);
    group:IsGroupMember(Mob c);
    group:IsGroupMember(string name);
    group:IsLeader(Mob c);
    group:IsLeader(string name);
    group:SetLeader(Mob c);
    group:SplitExp(uint64 exp, Mob other);
    group:SplitMoney(uint32 copper, uint32 silver, uint32 gold, uint32 platinum, Client splitter);
    group:SplitMoney(uint32 copper, uint32 silver, uint32 gold, uint32 platinum);
    group:TeleportGroup(Mob sender, uint32 zone_id, uint32 instance_id, float x, float y, float z, float h);
    ```
