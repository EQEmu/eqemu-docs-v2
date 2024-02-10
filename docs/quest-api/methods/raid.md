=== "Perl (28)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Raid){:target="Raid"} for latest definitions and Quest examples

    ``` perl
    $raid->BalanceHP(int32_t penalty, uint32_t group_id);
    $raid->CastGroupSpell(Mob* caster, uint16 spell_id, uint32 group_id);
    $raid->DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name, int max_check_count);
    $raid->DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name);
    $raid->GetClientByIndex(uint16_t member_index);
    $raid->GetGroup(string name);
    $raid->GetGroup(Client* client);
    $raid->GetGroupNumber(member_index);
    $raid->GetHighestLevel();
    $raid->GetID();
    $raid->GetLeader();
    $raid->GetLeaderName();
    $raid->GetLowestLevel();
    $raid->GetMember(int member_index);
    $raid->GetTotalRaidDamage(Mob* other);
    $raid->GroupCount(uint32_t group_id);
    $raid->IsGroupLeader(string who);
    $raid->IsGroupLeader(Client* c);
    $raid->IsLeader(Client* c);
    $raid->IsLeader(string name);
    $raid->IsRaidMember(string name);
    $raid->IsRaidMember(Client* c);
    $raid->RaidCount();
    $raid->SplitExp(uint32 experience, Mob* other);
    $raid->SplitMoney(uint32 gid, uint32 copper, uint32 silver, uint32 gold, uint32 platinum, Client* splitter);
    $raid->SplitMoney(uint32 gid, uint32 copper, uint32 silver, uint32 gold, uint32 platinum);
    $raid->TeleportGroup(Mob* sender, uint32 zone_id, float x, float y, float z, float heading, uint32 group_id);
    $raid->TeleportRaid(Mob* sender, uint32 zone_id, float x, float y, float z, float heading);
    ```
=== "Lua (28)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Raid){:target="Raid"} for latest definitions and Quest examples

    ``` lua
    raid:BalanceHP(int penalty, uint32 group_id);
    raid:CastGroupSpell(Mob caster, int spell_id, uint32 group_id);
    raid:DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name, int max_check_count);
    raid:DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name);
    raid:GetClientByIndex(int member_index);
    raid:GetGroup(string c);
    raid:GetGroup(Client c);
    raid:GetGroupNumber(member_index);
    raid:GetHighestLevel();
    raid:GetID();
    raid:GetLeader();
    raid:GetLeaderName();
    raid:GetLowestLevel();
    raid:GetMember(int member_index);
    raid:GetTotalRaidDamage(Mob other);
    raid:GroupCount(uint32 group_id);
    raid:IsGroupLeader(string name);
    raid:IsGroupLeader(Client c);
    raid:IsLeader(string c);
    raid:IsLeader(Client c);
    raid:IsRaidMember(Client c);
    raid:IsRaidMember(string name);
    raid:RaidCount();
    raid:SplitExp(uint64 exp, Mob other);
    raid:SplitMoney(uint32 gid, uint32 copper, uint32 silver, uint32 gold, uint32 platinum, Client splitter);
    raid:SplitMoney(uint32 gid, uint32 copper, uint32 silver, uint32 gold, uint32 platinum);
    raid:TeleportGroup(Mob sender, uint32 zone_id, uint32 instance_id, float x, float y, float z, float h, uint32 group_id);
    raid:TeleportRaid(Mob sender, uint32 zone_id, uint32 instance_id, float x, float y, float z, float h);
    ```
