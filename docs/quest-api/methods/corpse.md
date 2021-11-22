=== "Perl (34)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Corpse){:target="Corpse"} for latest definitions and Quest examples

        Last generated 2021.11.21 22:16:22

    ``` perl
    $corpse->AddItem(uint32 item_id, uint16 charges, [unt16 slot = 0]);
    $corpse->AddLooter(Mob* who);
    $corpse->AllowMobLoot(Mob* them, uint8 slot);
    $corpse->CanMobLoot(int character_id);
    $corpse->CastRezz(uint16 spell_id, [Mob* caster = nullptr]);
    $corpse->CompleteRezz();
    $corpse->CountItem(uint32 item_id);
    $corpse->CountItems();
    $corpse->Delete();
    $corpse->GetCharID();
    $corpse->GetCopper();
    $corpse->GetDBID();
    $corpse->GetDecayTime();
    $corpse->GetFirstSlotByItemID(uint32 item_id);
    $corpse->GetGold();
    $corpse->GetItemIDBySlot(uint16 loot_slot);
    $corpse->GetLootList();
    $corpse->GetOwnerName();
    $corpse->GetPlatinum();
    $corpse->GetSilver();
    $corpse->GetWornItem(equipSlot);
    $corpse->HasItem(uint32 item_id);
    $corpse->IsEmpty();
    $corpse->IsLocked();
    $corpse->IsRezzed();
    $corpse->Lock();
    $corpse->RemoveCash();
    $corpse->RemoveItem(uint16 loot_slot);
    $corpse->RemoveItemByID(uint32 item_id, [int quantity = 1]);
    $corpse->ResetLooter();
    $corpse->SetCash(uint16 copper, uint16 silver, uint16 gold, uint16 platinum);
    $corpse->SetDecayTimer(uint32 decay_time);
    $corpse->Summon(Client* client, bool is_spell);
    $corpse->UnLock();
    ```
=== "Lua (36)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Corpse){:target="Corpse"} for latest definitions and Quest examples

        Last generated 2021.11.21 22:16:22

    ``` lua
    corpse:AddItem(uint32 itemnum, uint16 charges, int16 slot, uint32 aug1, uint32 aug2, uint32 aug3, uint32 aug4, uint32 aug5);
    corpse:AddLooter(Mob who);
    corpse:AllowMobLoot(Mob them, uint8 slot);
    corpse:Bury();
    corpse:CanMobLoot(int charid);
    corpse:CountItem(uint32 item_id);
    corpse:CountItems();
    corpse:Delete();
    corpse:Depop();
    corpse:GetCharID();
    corpse:GetCopper();
    corpse:GetDBID();
    corpse:GetDecayTime();
    corpse:GetFirstSlotByItemID(uint32 item_id);
    corpse:GetGold();
    corpse:GetItemIDBySlot(uint16 loot_slot);
    corpse:GetLootList(State* L);
    corpse:GetOwnerName();
    corpse:GetPlatinum();
    corpse:GetSilver();
    corpse:GetWornItem(int16 equipSlot);
    corpse:HasItem(uint32 item_id);
    corpse:IsEmpty();
    corpse:IsLocked();
    corpse:IsRezzed();
    corpse:Lock();
    corpse:RemoveCash();
    corpse:RemoveItem(uint16 lootslot);
    corpse:RemoveItemByID(uint32 item_id);
    corpse:RemoveItemByID(uint32 item_id, int quantity);
    corpse:ResetLooter();
    corpse:Save();
    corpse:SetCash(uint32 copper, uint32 silver, uint32 gold, uint32 platinum);
    corpse:SetDecayTimer(uint32 decaytime);
    corpse:Summon(Client client, spell, checkdistance);
    corpse:UnLock();
    ```
