=== "Perl (34)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Corpse){:target="Corpse"} for latest definitions and Quest examples

    ``` perl
    $corpse->AddItem(uint32 item_id, uint16 charges);
    $corpse->AddItem(uint32 item_id, uint16 charges, uint16 slot);
    $corpse->AddLooter(Mob* who);
    $corpse->AllowMobLoot(Mob* them, uint8_t slot);
    $corpse->CanMobLoot(int character_id);
    $corpse->CastRezz(uint16_t spell_id, Mob* caster);
    $corpse->CompleteRezz();
    $corpse->CountItem(uint32_t item_id);
    $corpse->CountItems();
    $corpse->Delete();
    $corpse->GetCharID();
    $corpse->GetCopper();
    $corpse->GetDBID();
    $corpse->GetDecayTime();
    $corpse->GetGold();
    $corpse->GetLootList();
    $corpse->GetOwnerName();
    $corpse->GetPlatinum();
    $corpse->GetSilver();
    $corpse->GetWornItem(uint16_t equip_slot);
    $corpse->HasItem(uint32_t item_id);
    $corpse->IsEmpty();
    $corpse->IsLocked();
    $corpse->IsRezzed();
    $corpse->Lock();
    $corpse->RemoveItem(uint16_t loot_slot);
    $corpse->RemoveItemByID(uint32_t item_id);
    $corpse->RemoveItemByID(uint32_t item_id, int quantity);
    $corpse->ResetDecayTimer();
    $corpse->ResetLooter();
    $corpse->SetCash(uint16 copper, uint16 silver, uint16 gold, uint16 platinum);
    $corpse->SetDecayTimer(uint32_t decay_time);
    $corpse->Summon(Client* client, bool is_spell);
    $corpse->UnLock();
    ```
=== "Lua (39)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Corpse){:target="Corpse"} for latest definitions and Quest examples

    ``` lua
    corpse:AddItem(uint32 itemnum, uint16 charges, int16 slot, uint32 aug1, uint32 aug2, uint32 aug3, uint32 aug4, uint32 aug5);
    corpse:AddItem(uint32 itemnum, uint16 charges, int16 slot);
    corpse:AddItem(uint32 itemnum, uint16 charges);
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
    corpse:GetLootList();
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
    corpse:ResetDecayTimer();
    corpse:ResetLooter();
    corpse:Save();
    corpse:SetCash(uint32 copper, uint32 silver, uint32 gold, uint32 platinum);
    corpse:SetDecayTimer(uint32 decaytime);
    corpse:Summon(Client client, spell, checkdistance);
    corpse:UnLock();
    ```
