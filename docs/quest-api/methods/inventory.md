=== "Perl (19)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Inventory){:target="Inventory"} for latest definitions and Quest examples

        Last generated 2021.11.21 22:07:34

    ``` perl
    $inventory->CanItemFitInContainer(ItemInstance item_to_check, ItemInstance container_to_check);
    $inventory->CheckNoDrop(int16 slot_id);
    $inventory->DeleteItem(int16 slot_id, [uint8 quantity = 0]);
    $inventory->FindFreeSlot(bool is_for_bag, bool try_cursor, [uint8 min_size = 0, bool is_arrow = false]);
    $inventory->GetBagIndex(int16 slot_id);
    $inventory->GetItem(int16 slot_id);
    $inventory->GetMaterialFromSlot(int16 slot_id);
    $inventory->GetSlotByItemInst(ItemInstance item);
    $inventory->GetSlotFromMaterial(uint8 material);
    $inventory->GetSlotID(int16 slot_id, [uint8 bag_index]);
    $inventory->HasItem(uint32 item_id, [uint8 quantity, uint8 where]);
    $inventory->HasItemByLoreGroup(uint32 loregroup, [uint8 where]);
    $inventory->HasItemByUse(uint8 use, uint8 quantity, [uint8 where]);
    $inventory->HasSpaceForItem(ItemInstance item_to_check, uint8 quantity);
    $inventory->PopItem(int16 slot_id);
    $inventory->PushCursor(ItemInstance item);
    $inventory->PutItem(int16 slot_id, ItemInstance item);
    $inventory->SupportsContainers(int16 slot_id);
    $inventory->SwapItem(int16 source_slot_id, int16 destination_slot_id);
    ```
=== "Lua (29)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Inventory){:target="Inventory"} for latest definitions and Quest examples

        Last generated 2021.11.21 22:07:34

    ``` lua
    inventory:CalcBagIdx(slot_id);
    inventory:CalcMaterialFromSlot(equipslot);
    inventory:CalcSlotFromMaterial(material);
    inventory:CalcSlotId(slot_id, bag_slot);
    inventory:CalcSlotId(slot_id);
    inventory:CanItemFitInContainer(Item item, Item container);
    inventory:CheckNoDrop(int slot_id);
    inventory:DeleteItem(int slot_id);
    inventory:DeleteItem(int slot_id, int quantity);
    inventory:FindFreeSlot(bool for_bag, bool try_cursor);
    inventory:FindFreeSlot(bool for_bag, bool try_cursor, min_size);
    inventory:FindFreeSlot(bool for_bag, bool try_cursor, min_size, bool is_arrow);
    inventory:GetItem(int slot_id, int bag_slot);
    inventory:GetItem(int slot_id);
    inventory:GetSlotByItemInst(ItemInst inst);
    inventory:HasItem(item_id, quantity);
    inventory:HasItem(item_id);
    inventory:HasItem(item_id, quantity, where);
    inventory:HasItemByLoreGroup(uint32 loregroup);
    inventory:HasItemByLoreGroup(uint32 loregroup, where);
    inventory:HasItemByUse(use);
    inventory:HasItemByUse(use, uint8 quantity, uint8 where);
    inventory:HasItemByUse(use, uint8 quantity);
    inventory:HasSpaceForItem(Item item, int quantity);
    inventory:PopItem(int slot_id);
    inventory:PushCursor(ItemInst item);
    inventory:PutItem(slot_id, ItemInst item);
    inventory:SupportsContainers(int slot_id);
    inventory:SwapItem(int source_slot, int destination_slot);
    ```
