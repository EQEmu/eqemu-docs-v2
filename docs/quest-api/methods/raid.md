=== "Perl"

    ``` perl
    $raid->BalanceHP(int32 penalty, uint32 group_id);
    $raid->CastGroupSpell(Mob* caster, uint16 spell_id, uint32 group_id);
    $raid->DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name, [int max_check_count = 0]);
    $raid->GetClientByIndex(uint16 raid_index);
    $raid->GetGroup(string name);
    $raid->GetHighestLevel();
    $raid->GetID();
    $raid->GetLowestLevel();
    $raid->GetMember(int raid_index);
    $raid->GetTotalRaidDamage([Mob* other = nullptr]);
    $raid->GroupCount(uint32 group_id);
    $raid->IsGroupLeader(string name);
    $raid->IsLeader(string name);
    $raid->IsRaidMember(string name);
    $raid->RaidCount();
    $raid->SplitExp(uint32 experience, [Mob* other = nullptr]);
    $raid->SplitMoney(uint32 gid, uint32 copper, uint32 silver, uint32 gold, uint32 platinum);
    $raid->TeleportGroup(Mob* sender, uint32 zone_id, float x, float y, float z, float heading, uint32 group_id);
    $raid->TeleportRaid(Mob* sender, uint32 zone_id, float x, float y, float z, float heading);
    ```
=== "Lua"

    ``` lua
    raid:BalanceHP(int penalty, uint32 group_id);
    raid:CastGroupSpell(Mob caster, int spell_id, uint32 group_id);
    raid:DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name);
    raid:DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name, int max_check_count);
    raid:GetClientByIndex(int index);
    raid:GetGroup(string c);
    raid:GetGroup(Client c);
    raid:GetGroupNumber(index);
    raid:GetHighestLevel();
    raid:GetID();
    raid:GetLowestLevel();
    raid:GetMember(int index);
    raid:GetTotalRaidDamage(Mob other);
    raid:GroupCount(uint32 group_id);
    raid:IsGroupLeader(string name);
    raid:IsLeader(string c);
    raid:IsLeader(Client c);
    raid:IsRaidMember(string name);
    raid:RaidCount();
    raid:SplitExp(uint32 exp, Mob other);
    raid:SplitMoney(uint32 gid, uint32 copper, uint32 silver, uint32 gold, uint32 platinum, Client splitter);
    raid:SplitMoney(uint32 gid, uint32 copper, uint32 silver, uint32 gold, uint32 platinum);
    raid:TeleportGroup(Mob sender, uint32 zone_id, uint32 instance_id, float x, float y, float z, float h, uint32 group_id);
    raid:TeleportRaid(Mob sender, uint32 zone_id, uint32 instance_id, float x, float y, float z, float h);
    ```
