=== "Perl (26)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Raid){:target="Raid"} for latest definitions and Quest examples

        Last generated 2023.01.02

    ``` perl
    $raid->BalanceHP(int32_t penalty, uint32_t group_id);
    $raid->CastGroupSpell(Mob* caster, uint16 spell_id, uint32 group_id);
    $raid->DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name, int max_check_count);
    $raid->DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name);
    $raid->GetClientByIndex(uint16_t member_index);
    $raid->GetGroup(Client* client);
    $raid->GetGroup(string name);
    $raid->GetGroupNumber(member_index {);
    $raid->GetHighestLevel();
    $raid->GetID();
    $raid->GetLowestLevel();
    $raid->GetMember(int member_index);
    $raid->GetTotalRaidDamage(Mob* other);
    $raid->GroupCount(uint32_t group_id);
    $raid->IsGroupLeader(Client* c);
    $raid->IsGroupLeader(string who);
    $raid->IsLeader(string name);
    $raid->IsLeader(Client* c);
    $raid->IsRaidMember(Client* c);
    $raid->IsRaidMember(string name);
    $raid->RaidCount();
    $raid->SplitExp(uint32 experience, Mob* other);
    $raid->SplitMoney(uint32 gid, uint32 copper, uint32 silver, uint32 gold, uint32 platinum, Client* splitter);
    $raid->SplitMoney(uint32 gid, uint32 copper, uint32 silver, uint32 gold, uint32 platinum);
    $raid->TeleportGroup(Mob* sender, uint32 zone_id, float x, float y, float z, float heading, uint32 group_id);
    $raid->TeleportRaid(Mob* sender, uint32 zone_id, float x, float y, float z, float heading);
    ```
=== "Lua (26)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Raid){:target="Raid"} for latest definitions and Quest examples

        Last generated 2023.01.02

    ``` lua
    raid:BalanceHP(int penalty, uint32 group_id);
    raid:CastGroupSpell(Mob caster, int spell_id, uint32 group_id);
    raid:DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name, int max_check_count);
    raid:DoesAnyMemberHaveExpeditionLockout(string expedition_name, string event_name);
    raid:GetClientByIndex(int member_index);
    raid:GetGroup(Client c);
    raid:GetGroup(string c);
    raid:GetGroupNumber(member_index);
    raid:GetHighestLevel();
    raid:GetID();
    raid:GetLowestLevel();
    raid:GetMember(int member_index);
    raid:GetTotalRaidDamage(Mob other);
    raid:GroupCount(uint32 group_id);
    raid:IsGroupLeader(Client c);
    raid:IsGroupLeader(string name);
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
