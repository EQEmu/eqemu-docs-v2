=== "Perl (21)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=PerlPacket){:target="PerlPacket"} for latest definitions and Quest examples

    ``` perl
    $perlpacket->DESTROY();
    $perlpacket->FromArray(reference avref, uint32_t length);
    $perlpacket->GetByte(uint32_t pos);
    $perlpacket->GetFloat(uint32_t pos);
    $perlpacket->GetLong(pos);
    $perlpacket->GetShort(uint32_t pos);
    $perlpacket->Resize(uint32_t len);
    $perlpacket->SendTo(Client* who);
    $perlpacket->SendToAll();
    $perlpacket->SetByte(uint32_t pos, uint8_t val);
    $perlpacket->SetEQ1319(uint32_t pos, float part13, float part19);
    $perlpacket->SetEQ1913(uint32_t pos, float part19, float part13);
    $perlpacket->SetFloat(uint32_t pos, float val);
    $perlpacket->SetLong(uint32_t pos, uint32_t val);
    $perlpacket->SetOpcode(string opcode);
    $perlpacket->SetShort(uint32_t pos, uint16_t val);
    $perlpacket->SetString(uint32_t pos, char* str);
    $perlpacket->Zero();
    $perlpacket->new(string class_name);
    $perlpacket->new(string class_name, string opcode, uint32_t len);
    $perlpacket->new(string class_name, string opcode);
    ```
