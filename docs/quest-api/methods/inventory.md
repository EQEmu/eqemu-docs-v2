=== "Perl (32)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Inventory){:target="Inventory"} for latest definitions and Quest examples

    ``` perl
    $inventory->CanItemFitInContainer(EQ::ItemInstance* item_to_check, EQ::ItemInstance* container_to_check);
    $inventory->CheckNoDrop(int16_t slot_id);
    $inventory->CountAugmentEquippedByID(uint32_t item_id);
    $inventory->CountItemEquippedByID(uint32_t item_id);
    $inventory->DeleteItem(int16_t slot_id);
    $inventory->DeleteItem(int16_t slot_id, uint8_t quantity);
    $inventory->FindFreeSlot(bool is_for_bag, bool try_cursor);
    $inventory->FindFreeSlot(bool is_for_bag, bool try_cursor, uint8_t min_size);
    $inventory->FindFreeSlot(bool is_for_bag, bool try_cursor, uint8_t min_size, bool is_arrow);
    $inventory->GetAugmentIDsBySlotID(int16 slot_id);
    $inventory->GetBagIndex(int16_t slot_id);
    $inventory->GetItem(int16_t slot_id);
    $inventory->GetMaterialFromSlot(int16_t slot_id);
    $inventory->GetSlotByItemInst(EQ::ItemInstance* item);
    $inventory->GetSlotFromMaterial(uint8_t material);
    $inventory->GetSlotID(int16_t slot_id);
    $inventory->GetSlotID(int16_t slot_id, uint8_t bag_index);
    $inventory->HasAugmentEquippedByID(uint32_t item_id);
    $inventory->HasItem(uint32_t item_id, uint8_t quantity, uint8_t where_to_look);
    $inventory->HasItem(uint32_t item_id);
    $inventory->HasItem(uint32_t item_id, uint8_t quantity);
    $inventory->HasItemByLoreGroup(uint32_t loregroup, uint8_t where_to_look);
    $inventory->HasItemByLoreGroup(uint32_t loregroup);
    $inventory->HasItemByUse(uint8_t item_use, uint8_t quantity, uint8_t where_to_look);
    $inventory->HasItemByUse(uint8_t item_use, uint8_t quantity);
    $inventory->HasItemEquippedByID(uint32_t item_id);
    $inventory->HasSpaceForItem(EQ::ItemInstance* item_to_check, uint8_t quantity);
    $inventory->PopItem(int16_t slot_id);
    $inventory->PushCursor(EQ::ItemInstance* item);
    $inventory->PutItem(int16_t slot_id, EQ::ItemInstance* item);
    $inventory->SupportsContainers(int16_t slot_id);
    $inventory->SwapItem(int16_t source_slot_id, int16_t destination_slot_id);
    ```
=== "Lua (34)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Inventory){:target="Inventory"} for latest definitions and Quest examples

    ``` lua
    inventory:CalcBagIdx(slot_id);
    inventory:CalcMaterialFromSlot(equipslot);
    inventory:CalcSlotFromMaterial(material);
    inventory:CalcSlotId(slot_id);
    inventory:CalcSlotId(slot_id, bag_slot);
    inventory:CanItemFitInContainer(Item item, Item container);
    inventory:CheckNoDrop(int slot_id);
    inventory:CountAugmentEquippedByID(uint32 item_id);
    inventory:CountItemEquippedByID(uint32 item_id);
    inventory:DeleteItem(int slot_id);
    inventory:DeleteItem(int slot_id, int quantity);
    inventory:FindFreeSlot(bool for_bag, bool try_cursor);
    inventory:FindFreeSlot(bool for_bag, bool try_cursor, min_size);
    inventory:FindFreeSlot(bool for_bag, bool try_cursor, min_size, bool is_arrow);
    inventory:GetAugmentIDsBySlotID(int16 slot_id);
    inventory:GetItem(int slot_id);
    inventory:GetItem(int slot_id, int bag_slot);
    inventory:GetSlotByItemInst(ItemInst inst);
    inventory:HasAugmentEquippedByID(uint32 item_id);
    inventory:HasItem(item_id, quantity);
    inventory:HasItem(item_id, quantity, where);
    inventory:HasItem(item_id);
    inventory:HasItemByLoreGroup(uint32 loregroup);
    inventory:HasItemByLoreGroup(uint32 loregroup, where);
    inventory:HasItemByUse(use, uint8 quantity, uint8 where);
    inventory:HasItemByUse(use, uint8 quantity);
    inventory:HasItemByUse(use);
    inventory:HasItemEquippedByID(uint32 item_id);
    inventory:HasSpaceForItem(Item item, int quantity);
    inventory:PopItem(int slot_id);
    inventory:PushCursor(ItemInst item);
    inventory:PutItem(slot_id, ItemInst item);
    inventory:SupportsContainers(int slot_id);
    inventory:SwapItem(int source_slot, int destination_slot);
    ```
