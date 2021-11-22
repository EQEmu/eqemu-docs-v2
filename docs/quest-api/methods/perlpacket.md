=== "Perl (18)"

    !!! info end

        Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl&type=PerlPacket){:target="PerlPacket"} for latest definitions and Quest examples

    ``` perl
    $perlpacket->DESTROY();
    $perlpacket->FromArray(numbers, length);
    $perlpacket->GetByte(pos);
    $perlpacket->GetFloat(pos);
    $perlpacket->GetLong(pos);
    $perlpacket->GetShort(pos);
    $perlpacket->Resize(len);
    $perlpacket->SendTo(who);
    $perlpacket->SendToAll();
    $perlpacket->SetByte(pos, val);
    $perlpacket->SetEQ1319(pos, part13, part19);
    $perlpacket->SetEQ1913(pos, part19, part13);
    $perlpacket->SetFloat(pos, val);
    $perlpacket->SetLong(pos, val);
    $perlpacket->SetOpcode(opcode);
    $perlpacket->SetShort(pos, val);
    $perlpacket->SetString(pos, str);
    $perlpacket->Zero();
    ```
