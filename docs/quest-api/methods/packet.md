=== "Lua (33)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua&type=Packet){:target="Packet"} for latest definitions and Quest examples

    ``` lua
    packet:GetOpcode();
    packet:GetProtocolOpcode();
    packet:GetRawOpcode();
    packet:GetSize();
    packet:GetWritePosition();
    packet:ReadDouble(int offset);
    packet:ReadFixedLengthString(int offset, int string_length);
    packet:ReadFloat(int offset);
    packet:ReadInt16(offset);
    packet:ReadInt32(offset);
    packet:ReadInt64(int offset);
    packet:ReadInt8(offset);
    packet:ReadString(int offset);
    packet:SetOpcode(int op);
    packet:SetRawOpcode(int op);
    packet:SetWritePosition(int offset);
    packet:WriteDouble(double value);
    packet:WriteDouble(int offset, double value);
    packet:WriteFixedLengthString(int offset, string value, int string_length);
    packet:WriteFixedLengthString(string value);
    packet:WriteFloat(float value);
    packet:WriteFloat(int offset, float value);
    packet:WriteInt16(int value);
    packet:WriteInt16(int offset, int value);
    packet:WriteInt32(int offset, int value);
    packet:WriteInt32(int value);
    packet:WriteInt64(int offset, int64 value);
    packet:WriteInt64(int64 value);
    packet:WriteInt8(int offset, int value);
    packet:WriteInt8(int value);
    packet:WriteString(int offset, string value);
    packet:WriteString(string value);
    packet:operator=(const o);
    ```
