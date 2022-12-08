=== "Perl (15)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=QuestItem){:target="QuestItem"} for latest definitions and Quest examples

        Last generated 2022.12.07

    ``` perl
    $questitem->ContainsAugmentByID(uint32_t item_id);
    $questitem->CountAugmentByID(uint32_t item_id);
    $questitem->GetAugment(int slot_id);
    $questitem->GetCharges();
    $questitem->GetID();
    $questitem->GetName();
    $questitem->GetTaskDeliveredCount();
    $questitem->IsAttuned();
    $questitem->IsStackable();
    $questitem->IsType(int type);
    $questitem->ItemSay(string text, int language_id);
    $questitem->ItemSay(string text);
    $questitem->RemoveTaskDeliveredItems();
    $questitem->SetCharges(int16_t charges);
    $questitem->SetScale(float scale_multiplier);
    ```
