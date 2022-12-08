=== "Perl (98)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=EntityList){:target="EntityList"} for latest definitions and Quest examples

        Last generated 2022.12.07

    ``` perl
    $entitylist->CanAddHateForMob(Mob* target);
    $entitylist->Clear();
    $entitylist->ClearClientPetitionQueue();
    $entitylist->ClearFeignAggro(Mob* target);
    $entitylist->DeleteNPCCorpses();
    $entitylist->DeletePlayerCorpses();
    $entitylist->DoubleAggro(Mob* who);
    $entitylist->Fighting(Mob* target);
    $entitylist->FindDoor(uint32_t door_id);
    $entitylist->GetBotByID(uint32_t bot_id);
    $entitylist->GetBotByName(string bot_name);
    $entitylist->GetBotList();
    $entitylist->GetBotListByCharacterID(uint32_t character_id);
    $entitylist->GetBotListByCharacterID(uint32_t character_id, uint8_t class_id);
    $entitylist->GetBotListByClientName(string client_name, uint8 class_id);
    $entitylist->GetBotListByClientName(string client_name);
    $entitylist->GetClientByAccID(uint32_t account_id);
    $entitylist->GetClientByCharID(uint32_t character_id);
    $entitylist->GetClientByID(uint16_t client_id);
    $entitylist->GetClientByName(string name);
    $entitylist->GetClientByWID(uint32_t wid);
    $entitylist->GetClientList();
    $entitylist->GetCorpseByID(uint16_t id);
    $entitylist->GetCorpseByName(string name);
    $entitylist->GetCorpseByOwner(Client* client);
    $entitylist->GetCorpseList();
    $entitylist->GetDoorsByDBID(uint32_t database_id);
    $entitylist->GetDoorsByDoorID(uint32_t door_id);
    $entitylist->GetDoorsByID(uint32_t entity_id);
    $entitylist->GetDoorsList();
    $entitylist->GetGroupByClient(Client* client);
    $entitylist->GetGroupByID(uint32_t id);
    $entitylist->GetGroupByLeaderName(string leader_name);
    $entitylist->GetGroupByMob(Mob* mob);
    $entitylist->GetMob(string name);
    $entitylist->GetMobByID(uint16_t mob_id);
    $entitylist->GetMobByNpcTypeID(uint32_t npc_type_id);
    $entitylist->GetMobID(uint16_t mob_id);
    $entitylist->GetMobList();
    $entitylist->GetNPCByID(uint16_t id);
    $entitylist->GetNPCByNPCTypeID(uint32_t npc_id);
    $entitylist->GetNPCBySpawnID(uint32_t spawn_id);
    $entitylist->GetNPCList();
    $entitylist->GetObjectByDBID(uint32_t database_id);
    $entitylist->GetObjectByID(uint32_t entity_id);
    $entitylist->GetObjectList();
    $entitylist->GetRaidByClient(Client* client);
    $entitylist->GetRaidByID(uint32_t id);
    $entitylist->GetRandomBot(float x, float y, float z, float distance, exclude_bot);
    $entitylist->GetRandomBot();
    $entitylist->GetRandomBot(float x, float y, float z, float distance);
    $entitylist->GetRandomClient();
    $entitylist->GetRandomClient(float x, float y, float z, float distance);
    $entitylist->GetRandomClient(float x, float y, float z, float distance, exclude_client);
    $entitylist->GetRandomMob(float x, float y, float z, float distance);
    $entitylist->GetRandomMob();
    $entitylist->GetRandomMob(float x, float y, float z, float distance, exclude_mob);
    $entitylist->GetRandomNPC(float x, float y, float z, float distance);
    $entitylist->GetRandomNPC(float x, float y, float z, float distance, exclude_npc);
    $entitylist->GetRandomNPC();
    $entitylist->HalveAggro(Mob* who);
    $entitylist->IsMobSpawnedByNpcTypeID(uint32_t npc_type_id);
    $entitylist->MakeNameUnique(char* name);
    $entitylist->Marquee(uint32 type, string message);
    $entitylist->Marquee(uint32 type, string message, uint32 duration);
    $entitylist->Marquee(uint32 type, uint32 priority, uint32 fade_in, uint32 fade_out, uint32 duration, string message);
    $entitylist->Message(uint32 guild_id, uint32 color_type, string message);
    $entitylist->MessageClose(Mob* sender, bool skip_sender, float distance, uint32 color_type, string message);
    $entitylist->MessageGroup(Mob* sender, bool skip_close, uint32_t emote_color_type, string message);
    $entitylist->MessageStatus(uint32 guild_id, int to_minstatus, uint32 color_type, string message);
    $entitylist->OpenDoorsNear(Mob* opener);
    $entitylist->RemoveAllClients();
    $entitylist->RemoveAllCorpses();
    $entitylist->RemoveAllDoors();
    $entitylist->RemoveAllGroups();
    $entitylist->RemoveAllMobs();
    $entitylist->RemoveAllNPCs();
    $entitylist->RemoveAllObjects();
    $entitylist->RemoveAllTraps();
    $entitylist->RemoveClient(uint16_t delete_id);
    $entitylist->RemoveCorpse(uint16_t delete_id);
    $entitylist->RemoveDoor(uint16_t delete_id);
    $entitylist->RemoveEntity(uint16_t id);
    $entitylist->RemoveFromHateLists(Mob* mob, bool set_to_one);
    $entitylist->RemoveFromHateLists(Mob* mob);
    $entitylist->RemoveFromTargets(Mob* mob);
    $entitylist->RemoveGroup(uint32_t delete_id);
    $entitylist->RemoveMob(uint16_t delete_id);
    $entitylist->RemoveNPC(uint16_t delete_id);
    $entitylist->RemoveNumbers(char* name);
    $entitylist->RemoveObject(uint16_t delete_id);
    $entitylist->RemoveTrap(uint16_t delete_id);
    $entitylist->ReplaceWithTarget(Mob* old_mob, Mob* new_target);
    $entitylist->SignalAllBotsByOwnerCharacterID(uint32_t character_id, int signal_id);
    $entitylist->SignalAllClients(int signal_id);
    $entitylist->SignalBotByBotID(uint32_t bot_id, int signal_id);
    $entitylist->SignalBotByBotName(string bot_name, int signal_id);
    $entitylist->SignalMobsByNPCID(uint32 npc_type_id, int signal_id);
    ```
=== "Lua (86)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=EntityList){:target="EntityList"} for latest definitions and Quest examples

        Last generated 2022.12.07

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
    entitylist:GetBotListByCharacterID(uint32 character_id);
    entitylist:GetBotListByCharacterID(uint32 character_id, uint8 class_id);
    entitylist:GetBotListByClientName(string client_name);
    entitylist:GetBotListByClientName(string client_name, uint8 class_id);
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
    entitylist:GetRandomBot(float x, float y, float z, float distance, exclude_bot);
    entitylist:GetRandomBot(float x, float y, float z, float distance);
    entitylist:GetRandomBot();
    entitylist:GetRandomClient();
    entitylist:GetRandomClient(float x, float y, float z, float distance, exclude_client);
    entitylist:GetRandomClient(float x, float y, float z, float distance);
    entitylist:GetRandomMob();
    entitylist:GetRandomMob(float x, float y, float z, float distance, exclude_mob);
    entitylist:GetRandomMob(float x, float y, float z, float distance);
    entitylist:GetRandomNPC(float x, float y, float z, float distance, exclude_npc);
    entitylist:GetRandomNPC(float x, float y, float z, float distance);
    entitylist:GetRandomNPC();
    entitylist:GetShuffledClientList();
    entitylist:GetSpawnByID(uint32 id);
    entitylist:GetSpawnList();
    entitylist:HalveAggro(Mob who);
    entitylist:IsMobSpawnedByNpcTypeID(int npc_type);
    entitylist:MakeNameUnique(string name);
    entitylist:Marquee(uint32 type, string message);
    entitylist:Marquee(uint32 type, string message, uint32 duration);
    entitylist:Marquee(uint32 type, uint32 priority, uint32 fade_in, uint32 fade_out, uint32 duration, string message);
    entitylist:Message(uint32 guild_dbid, uint32 type, string message);
    entitylist:MessageClose(Mob sender, bool skip_sender, float dist, uint32 type, string message);
    entitylist:MessageGroup(Mob who, bool skip_close, uint32 type, string message);
    entitylist:MessageStatus(uint32 guild_dbid, int min_status, uint32 type, string message);
    entitylist:OpenDoorsNear(Mob opener);
    entitylist:RemoveFromHateLists(Mob who, bool set_to_one);
    entitylist:RemoveFromHateLists(Mob who);
    entitylist:RemoveFromTargets(Mob mob);
    entitylist:RemoveFromTargets(Mob mob, bool RemoveFromXTargets);
    entitylist:RemoveNumbers(string name);
    entitylist:ReplaceWithTarget(Mob target, Mob new_target);
    entitylist:SignalAllBotsByOwnerCharacterID(uint32 character_id, int signal_id);
    entitylist:SignalAllClients(int signal_id);
    entitylist:SignalBotByBotID(uint32 bot_id, int signal_id);
    entitylist:SignalBotByBotName(string bot_name, int signal_id);
    entitylist:SignalMobsByNPCID(uint32 npc_id, int signal_id);
    ```
