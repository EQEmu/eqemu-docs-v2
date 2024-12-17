=== "Perl (30)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Merc){:target="Merc"} for latest definitions and Quest examples

    ``` perl
    $merc->GetCostFormula();
    $merc->GetGroup();
    $merc->GetHatedCount();
    $merc->GetMaxMeleeRangeToTarget(Mob* target);
    $merc->GetMercenaryCharacterID();
    $merc->GetMercenaryID();
    $merc->GetMercenaryNameType();
    $merc->GetMercenaryOwner();
    $merc->GetMercenarySubtype();
    $merc->GetMercenaryTemplateID();
    $merc->GetMercenaryType();
    $merc->GetOwner();
    $merc->GetOwnerOrSelf();
    $merc->GetProficiencyID();
    $merc->GetStance();
    $merc->GetTierID();
    $merc->HasOrMayGetAggro();
    $merc->IsMercenaryCaster();
    $merc->IsMercenaryCasterCombatRange(Mob* target);
    $merc->IsSitting();
    $merc->IsStanding();
    $merc->ScaleStats(int scale_percentage);
    $merc->ScaleStats(int scale_percentage, bool set_to_max);
    $merc->SendPayload(int payload_id, string payload_value);
    $merc->SetTarget(Mob* target);
    $merc->Signal(int signal_id);
    $merc->Sit();
    $merc->Stand();
    $merc->Suspend();
    $merc->UseDiscipline(uint16 spell_id, uint16 target_id);
    ```
=== "Lua (30)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Merc){:target="Merc"} for latest definitions and Quest examples

    ``` lua
    merc:GetCostFormula();
    merc:GetGroup();
    merc:GetHatedCount();
    merc:GetMaxMeleeRangeToTarget(Mob target);
    merc:GetMercenaryCharacterID();
    merc:GetMercenaryID();
    merc:GetMercenaryNameType();
    merc:GetMercenaryOwner();
    merc:GetMercenarySubtype();
    merc:GetMercenaryTemplateID();
    merc:GetMercenaryType();
    merc:GetOwner();
    merc:GetOwnerOrSelf();
    merc:GetProficiencyID();
    merc:GetStance();
    merc:GetTierID();
    merc:HasOrMayGetAggro();
    merc:IsMercenaryCaster();
    merc:IsMercenaryCasterCombatRange(Mob target);
    merc:IsSitting();
    merc:IsStanding();
    merc:ScaleStats(int scale_percentage);
    merc:ScaleStats(int scale_percentage, bool set_to_max);
    merc:SendPayload(int payload_id, string payload_value);
    merc:SetTarget(Mob target);
    merc:Signal(int signal_id);
    merc:Sit();
    merc:Stand();
    merc:Suspend();
    merc:UseDiscipline(uint16 spell_id, uint16 target_id);
    ```
