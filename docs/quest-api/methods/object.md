=== "Perl (42)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=Object){:target="Object"} for latest definitions and Quest examples

        Last generated 2021.11.21 22:14:50

    ``` perl
    $object->ClearUser();
    $object->Close();
    $object->Delete([bool reset_state = false]);
    $object->DeleteItem(uint8 index);
    $object->Depop();
    $object->EntityVariableExists(string key);
    $object->GetDBID();
    $object->GetEntityVariable(string key);
    $object->GetHeading();
    $object->GetID();
    $object->GetIcon();
    $object->GetItemID();
    $object->GetModelName();
    $object->GetSize();
    $object->GetSize();
    $object->GetSize();
    $object->GetSolidType();
    $object->GetType();
    $object->GetX();
    $object->GetY();
    $object->GetZ();
    $object->IsGroundSpawn();
    $object->IsObject();
    $object->Repop();
    $object->Save();
    $object->SetEntityVariable(string key, string var);
    $object->SetHeading(float heading);
    $object->SetID(uint16 id);
    $object->SetIcon(uint32 icon);
    $object->SetItemID(uint32 item_id);
    $object->SetLocation(float x, float y, float z);
    $object->SetModelName(string name);
    $object->SetSize(float size);
    $object->SetSolidType(uint16 type);
    $object->SetTiltX(float tilt_x);
    $object->SetTiltY(float tilt_y);
    $object->SetType(uint32 type);
    $object->SetX(float x);
    $object->SetY(float y);
    $object->SetZ(float z);
    $object->StartDecay();
    $object->VarSave();
    ```
=== "Lua (34)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Object){:target="Object"} for latest definitions and Quest examples

        Last generated 2021.11.21 22:14:50

    ``` lua
    object:ClearUser();
    object:Close();
    object:Delete(bool reset_state);
    object:Delete();
    object:DeleteItem(int index);
    object:Depop();
    object:EntityVariableExists(string name);
    object:GetDBID();
    object:GetEntityVariable(name);
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
    object:SetEntityVariable(string name, string value);
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
