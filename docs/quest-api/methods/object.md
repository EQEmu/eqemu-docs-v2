=== "Perl (46)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Object){:target="Object"} for latest definitions and Quest examples

    ``` perl
    $object->ClearEntityVariables();
    $object->ClearUser();
    $object->Close();
    $object->Delete();
    $object->Delete(bool reset_state);
    $object->DeleteEntityVariable(string variable_name);
    $object->DeleteItem(uint8_t index);
    $object->Depop();
    $object->EntityVariableExists(string variable_name);
    $object->GetDBID();
    $object->GetEntityVariable(variable_name);
    $object->GetEntityVariables();
    $object->GetHeading();
    $object->GetID();
    $object->GetIcon();
    $object->GetItemID();
    $object->GetModelName();
    $object->GetSize();
    $object->GetSolidType();
    $object->GetTiltX();
    $object->GetTiltY();
    $object->GetType();
    $object->GetX();
    $object->GetY();
    $object->GetZ();
    $object->IsGroundSpawn();
    $object->IsObject();
    $object->Repop();
    $object->Save();
    $object->SetEntityVariable(string variable_name, string variable_value);
    $object->SetHeading(float heading);
    $object->SetID(uint16_t set_id);
    $object->SetIcon(uint32_t icon);
    $object->SetItemID(uint32_t itemid);
    $object->SetLocation(float x, float y, float z);
    $object->SetModelName(string name);
    $object->SetSize(float size);
    $object->SetSolidType(uint16_t type);
    $object->SetTiltX(float tilt_x);
    $object->SetTiltY(float tilt_y);
    $object->SetType(uint32_t type);
    $object->SetX(float x);
    $object->SetY(float y);
    $object->SetZ(float z);
    $object->StartDecay();
    $object->VarSave();
    ```
=== "Lua (37)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Object){:target="Object"} for latest definitions and Quest examples

    ``` lua
    object:ClearEntityVariables();
    object:ClearUser();
    object:Close();
    object:Delete();
    object:Delete(bool reset_state);
    object:DeleteEntityVariable(string variable_name);
    object:DeleteItem(int index);
    object:Depop();
    object:EntityVariableExists(string variable_name);
    object:GetDBID();
    object:GetEntityVariable(variable_name);
    object:GetEntityVariables();
    object:GetHeading();
    object:GetID();
    object:GetIcon();
    object:GetItemID();
    object:GetModelName();
    object:GetType();
    object:GetX();
    object:GetY();
    object:GetZ();
    object:IsGroundSpawn();
    object:Repop();
    object:Save();
    object:SetEntityVariable(string variable_name, string variable_value);
    object:SetHeading(float h);
    object:SetID(int user);
    object:SetIcon(uint32 icon);
    object:SetItemID(uint32 item_id);
    object:SetLocation(float x, float y, float z);
    object:SetModelName(string name);
    object:SetType(uint32 type);
    object:SetX(float x);
    object:SetY(float y);
    object:SetZ(float z);
    object:StartDecay();
    object:VarSave();
    ```
