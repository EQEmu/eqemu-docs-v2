=== "Perl (75)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=EntityList){:target="EntityList"} for latest definitions and Quest examples

        Last generated 2021.11.21 22:07:34

    ``` perl
    $entitylist->CanAddHateForMob(Mob* target);
    $entitylist->Clear();
    $entitylist->ClearClientPetitionQueue();
    $entitylist->ClearFeignAggro(Mob* target);
    $entitylist->DeleteNPCCorpses();
    $entitylist->DeletePlayerCorpses();
    $entitylist->DoubleAggro(*Mob target);
    $entitylist->Fighting(Mob* target);
    $entitylist->FindDoor(uint32 door_id);
    $entitylist->GetBotByID(uint32 bot_id);
    $entitylist->GetBotByName(string bot_name);
    $entitylist->GetBotList();
    $entitylist->GetClientByAccID(uint32 account_id);
    $entitylist->GetClientByCharID(uint32 character_id);
    $entitylist->GetClientByID(uint16 client_id);
    $entitylist->GetClientByName(name);
    $entitylist->GetClientByWID(uint32 wid);
    $entitylist->GetClientList();
    $entitylist->GetCorpseByID(id);
    $entitylist->GetCorpseByName(name);
    $entitylist->GetCorpseByOwner(client);
    $entitylist->GetCorpseList();
    $entitylist->GetDoorsByDBID(uint32 database_id);
    $entitylist->GetDoorsByDoorID(uint32 door_id);
    $entitylist->GetDoorsByID(uint32 entity_id);
    $entitylist->GetDoorsList();
    $entitylist->GetGroupByClient(Client* client);
    $entitylist->GetGroupByID(id);
    $entitylist->GetGroupByLeaderName(leader);
    $entitylist->GetGroupByMob(Mob* mob);
    $entitylist->GetMob(name);
    $entitylist->GetMobByID(id);
    $entitylist->GetMobByNpcTypeID(get_id);
    $entitylist->GetMobID(id);
    $entitylist->GetMobList();
    $entitylist->GetNPCByID(id);
    $entitylist->GetNPCByNPCTypeID(npc_id);
    $entitylist->GetNPCBySpawnID(spawn_id);
    $entitylist->GetNPCList();
    $entitylist->GetObjectByDBID(uint32 database_id);
    $entitylist->GetObjectByID(uint32 entity_id);
    $entitylist->GetObjectList();
    $entitylist->GetRaidByClient(client);
    $entitylist->GetRaidByID(id);
    $entitylist->GetRandomClient(float x, float y, float z, float distance, [Client* exclude_client = nullptr]);
    $entitylist->HalveAggro(Mob* target);
    $entitylist->IsMobSpawnedByNpcTypeID(get_id);
    $entitylist->MakeNameUnique(string name);
    $entitylist->Message(uint32 guild_id, uint32 emote_color_type, string message);
    $entitylist->MessageClose(Mob* sender, bool skip_sender, float distance, uint32 emote_color_type, string message);
    $entitylist->MessageGroup(Mob* sender, bool skip_close, uint32 emote_color_type, string message);
    $entitylist->MessageStatus(uint32 guild_id, uint32 emote_color_type, string message);
    $entitylist->OpenDoorsNear(NPC* opener);
    $entitylist->RemoveAllClients();
    $entitylist->RemoveAllCorpses();
    $entitylist->RemoveAllDoors();
    $entitylist->RemoveAllGroups();
    $entitylist->RemoveAllMobs();
    $entitylist->RemoveAllNPCs();
    $entitylist->RemoveAllObjects();
    $entitylist->RemoveAllTraps();
    $entitylist->RemoveClient(delete_id);
    $entitylist->RemoveCorpse(delete_id);
    $entitylist->RemoveDoor(delete_id);
    $entitylist->RemoveEntity(uint16 id);
    $entitylist->RemoveFromHateLists(Mob* mob, [bool set_to_one = false]);
    $entitylist->RemoveFromTargets(Mob* target);
    $entitylist->RemoveGroup(delete_id);
    $entitylist->RemoveMob(delete_id);
    $entitylist->RemoveNPC(delete_id);
    $entitylist->RemoveObject(delete_id);
    $entitylist->RemoveTrap(delete_id);
    $entitylist->ReplaceWithTarget(Mob* old_mob, Mob* new_target);
    $entitylist->SignalAllClients(uint32 data);
    $entitylist->SignalMobsByNPCID(uint32 npc_type_id, int signal_id);
    ```
=== "Lua (66)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=EntityList){:target="EntityList"} for latest definitions and Quest examples

        Last generated 2021.11.21 22:07:34

    ``` lua
    entitylist:CanAddHateForMob(Mob p);
    entitylist:ChannelMessage(Mob from, int channel_num, int language, string message);
    entitylist:ClearClientPetitionQueue();
    entitylist:ClearFeignAggro(Mob who);
    entitylist:DeleteNPCCorpses();
    entitylist:DeletePlayerCorpses();
    entitylist:DoubleAggro(Mob who);
    entitylist:Fighting(Mob who);
    entitylist:FilteredMessageClose(Mob sender, bool skip_sender, float dist, uint32 type, int filter, string message);
    entitylist:FindDoor(uint32 id);
    entitylist:GetBotByID(uint32 bot_id);
    entitylist:GetBotByName(string bot_name);
    entitylist:GetBotList();
    entitylist:GetClientByAccID(uint32 acct_id);
    entitylist:GetClientByCharID(uint32 char_id);
    entitylist:GetClientByID(int id);
    entitylist:GetClientByName(string name);
    entitylist:GetClientByWID(uint32 wid);
    entitylist:GetClientList();
    entitylist:GetCorpseByID(int id);
    entitylist:GetCorpseByName(string name);
    entitylist:GetCorpseByOwner(Client client);
    entitylist:GetCorpseList();
    entitylist:GetDoorsByDBID(uint32 db_id);
    entitylist:GetDoorsByDoorID(uint32 door_id);
    entitylist:GetDoorsByID(int id);
    entitylist:GetDoorsList();
    entitylist:GetGroupByClient(Client client);
    entitylist:GetGroupByID(int id);
    entitylist:GetGroupByLeaderName(string name);
    entitylist:GetGroupByMob(Mob mob);
    entitylist:GetMob(string name);
    entitylist:GetMob(int id);
    entitylist:GetMobByNpcTypeID(int npc_type);
    entitylist:GetMobID(int id);
    entitylist:GetMobList();
    entitylist:GetNPCByID(int id);
    entitylist:GetNPCByNPCTypeID(int npc_type);
    entitylist:GetNPCBySpawnID(uint32 spawn_id);
    entitylist:GetNPCList();
    entitylist:GetObjectByDBID(uint32 db_id);
    entitylist:GetObjectByID(int id);
    entitylist:GetObjectList();
    entitylist:GetRaidByClient(Client client);
    entitylist:GetRaidByID(int id);
    entitylist:GetRandomClient(float x, float y, float z, float dist);
    entitylist:GetRandomClient(float x, float y, float z, float dist, exclude);
    entitylist:GetShuffledClientList();
    entitylist:GetSpawnByID(uint32 id);
    entitylist:GetSpawnList();
    entitylist:HalveAggro(Mob who);
    entitylist:IsMobSpawnedByNpcTypeID(int npc_type);
    entitylist:MakeNameUnique(string name);
    entitylist:Message(uint32 guild_dbid, uint32 type, string message);
    entitylist:MessageClose(Mob sender, bool skip_sender, float dist, uint32 type, string message);
    entitylist:MessageGroup(Mob who, bool skip_close, uint32 type, string message);
    entitylist:MessageStatus(uint32 guild_dbid, int min_status, uint32 type, string message);
    entitylist:OpenDoorsNear(Mob opener);
    entitylist:RemoveFromHateLists(Mob who);
    entitylist:RemoveFromHateLists(Mob who, bool set_to_one);
    entitylist:RemoveFromTargets(Mob mob, bool RemoveFromXTargets);
    entitylist:RemoveFromTargets(Mob mob);
    entitylist:RemoveNumbers(string name);
    entitylist:ReplaceWithTarget(Mob target, Mob new_target);
    entitylist:SignalAllClients(int signal);
    entitylist:SignalMobsByNPCID(uint32 npc_id, int signal);
    ```
