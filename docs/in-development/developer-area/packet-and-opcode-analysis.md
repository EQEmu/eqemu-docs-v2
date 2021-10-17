# Packet and OpCode Analysis

Expanding features of currently supported Everquest clients is not a trivial matter, and is a task made more complex without packet captures from the time when that client was in use. It is still possible to identify data structures for older clients by examining captures from the most recent client.

The information that follows is a collection of basic information I’ve learned about capturing packets, interpreting them, finding opcodes, and how to test them in the emulator. The information not included here is that it’s a ton of trial and error, testing theories, basic debugging skill, familiarity with live debugging of processes without debug symbols, etc. Depending on what you want to do, this is not a trivial task.

## Capturing Packets from Live

There are a variety of tools available for capturing network traffic. My preference is [Microsoft Message Analyzer](https://www.microsoft.com/en-us/download/details.aspx?id=44226). [Wireshark](https://www.wireshark.org) is another very popular choice. If you don't have any experience with these types of tools, take some time and familiarize yourself with a few of them before picking one.

The capture tool can be installed on the same computer that you will use to play Everquest, or it can be installed on a separate machine that sits between your client and the Everquest servers. Neither poses a risk of being detected by Daybreak as packet logging is a passive operation. The benefit to using a proxy device is that you can modify packets before they reach your client machine. For the purposes of this article, I am only focusing on same box installations.

## Getting a capture

Before you launch the capture tool, make sure you've shut down as many running processes on your machine as possible (mail clients, RDP sessions, chat programs, etc.) This will reduce noise by reducing the amount of potential network traffic. You can also filter this out later, but I prefer to start with as small a set of data as possible.

In your capture tool, select to log only UDP protocol traffic. You can further simplify this by filtering the capture to be UDP packets from your machine's IP address to EQ's servers, and packets from EQ's servers to your IP address.

Sample filtering rule for Microsoft Network Analyzer to reduce noise: `UDP and !TCP and !SSDP and !ICMPv6 and !DNS and !LLMNR and !WSDiscovery and !IPv6NoNxt and !IPv6 and *Destination != 10.0.0.255 and *Destination != yourmachineiphere and *Source != 192.168.2.225`

Another approach: `UDP and (*Source in 69.174.0.0/16 or *Destination in 69.174.0.0/16)`

When I’m setting up for a capture, I initially log the entire login sequence from authentication, through character select, and into the server. That captures a large amount of data that is helpful for identifying the specific patch (the last time the client was patched), the server’s entire list of guilds, character select information, etc. This approach does log your account name and password. If you plan on sharing this capture file, make sure you delete those packets, and the packets that list your character names.

To configure Message Anaylzer to properly export for the extractor, first you need to right click the header column of the preview pane, and go to Add Column... Navigate to TCP->Segment->Payload. Right click and select Add as column. Now, right click a row in the Payload section column, select Display Binary Values As -> Hex. Now, To export the capture in Message Analyzer, go to Session -> Analysis Grid -> Export -> All, in the drop down choose Tab Delimited (\*.txt)

Then, I will move my character to the area of the game I want to investigate. If, for example, I wanted to learn more about how Shrouds work I would move to a Shroud “vendor”. Then I would start a capture session, interact with the vendor, and save the capture. Screenshots are also helpful, especially for cases where you can see information on the screen. This will help pinpoint specific values (data type and placement) in the packet. I usually repeat that process at least one more time to pinpoint which packets relate to shrouds and which were just noise from other people in the zone.

Sometimes the data for a particular action is sent once and cached in memory, so it can be helpful to zone and/or log off and back on to get more information about how the server and client interact. It is also useful to get a capture from a different server to isolate if values are constant across servers, or if they are server specific. For instance, zone IDs will be constant, but player and guild IDs will vary.

## ShowEQ Source

ShowEQ is a valuable second source of information. While they do not map every Opcode in each ShowEQ patch, they do map enough of them to help you pin down some additional information.

Navigate to the ShowEQ repo’s [Tags](https://sourceforge.net/p/seq/svn/HEAD/tree/showeq/tags/) folder. Here you will find a list of all releases, each corresponding to a patch from Verant/Sony/Daybreak.

Look at the Date column to see when each patch was released. Select one of the folders and then select the conf subfolder. The only files we are concerned with are worldopcodes.xml and zoneopcodes.xml. These contain the world and zone opcodes, but not all of them were updated for the latest patch. Only the ones ShowEQ needs were updated. Thankfully, if you look inside each file you can see the date/time each opcode was updated.

## Reading a capture

You cannot directly use a packet capture. Most packets will be compressed and most likely transmitted as Combined or Oversized packets. Making sense of them requires understanding the network-level packet structures so that you can decompress them and break them into individual packets. The only good way to do this is via code. I’m sure there is an existing toolset to do this, but I opted for writing one. You can find the source [here](https://github.com/daerath/EQPacketParser). It is fairly basic in function, which is perfectly fine for my purposes (so far).

Once the packets have been uncompressed and split apart, the opcodes from ShowEQ come into play. The tool maps packet opcodes to known opcodes and indicates packet direction (to or from the server). It also outputs each packet in byte format, hex format, and as a string. This helps make some calculations easier and the raw string output can enable faster identification of packet structure if it contains string values.

Unfortunately, most Opcodes are not mapped each time a patch comes out, and because Daybreak still changes Opcodes with each patch, most packets are not immediately identifiable. Therefore, it is extremely important to focus your packet captures to the smallest possible sample size.

Sample output of an Emulator server OP_ChannelMessage packet:\


> Idx: 16\
> Size:56\
> From: Server\
> OpCode: OP_ChannelMessage::bc,33\
> Bytes:0,9,1,248,188,51,78,105,103,101,108,0,0,0,0,0,0,0,0,0,0,8,0,0,0,0,0,0,0,0,100,0,0,0,112,97,99,107,101,116,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0\
> Hex:00,09,01,f8,bc,33,4e,69,67,65,6c,00,00,00,00,00,00,00,00,00,00,08,00,00,00,00,00,00,00,00,64,00,00,00,70,61,63,6b,65,74,00,00,00,00,00,00,00,00,00,00,00,00,00,00,00,00\
> String: ??3Nigel d packet  \
>

In the above case, we know the format because it is a well-known type, and it is from an Emulator server, but if we did not, the next step would be to identify suspected values. When there are visible words, you usually have a preceding string length (e.g. 8,0,0,0 for an 8-character string), or if not, a trailing null terminator (0).

## Opcode identification in the client

Finding an opcode in the current client is easy. It’s in the packet capture. Finding the corresponding opcode in an emulator supported client is more complex. If ShowEQ doesn’t have a mapping, and if the emulator doesn’t already have that code identified, the only option remaining is to decompile eqgame.exe and try to find it. I am going to very briefly touch on this because it is a much more complex topic and requires knowledge of assembly language and debugging tools such as [windbg](http://windbg.org), [windbgx](https://www.microsoft.com/en-us/p/windbg-preview/9pgjgd53tn86?activetab=pivot:overviewtab), [IDA Pro](https://www.hex-rays.com/products/ida/index.shtml), etc. Opcodes appear in the following structure:\


> cmp edx,1659h\
> ja loc\_13FF3ED\
> jz loc\_13FF512

They may also appear like this:\


> cmp \[ebp+var\_5354], 6786h\
> ja short loc\_4C1BD8\
> cmp \[ebp+var\_5354], 6786h\
> jz short loc\_4C90DB

They also appear in multiple sections with login-based opcodes (e.g. guild list, password mismatch, etc.) typically being in an entirely separate part of the executable. The other opcodes are in blocks (like those above) with jumps to other blocks of opcodes (if smaller than X, goto location ABC, else goto location ZYX).

The way I’ve done this, and there may be a better approach, is to first look in the existing client for that opcode. Then I look at the assembly instructions that are executed by that opcode to learn more about what happens when that opcode is received. From there I’ll open the version of the client that is supported by the emulator and attempt to find the same pattern of instructions.

Once you’ve identified a probable list of opcodes in the target client version, and you’ve picked apart the various fields in the packet, you can quickly attempt to simulate it using an emulator server and some quest NPCs.

## Sending packets using LUA quest files

The fastest way to identify fields in a packet, or to attempt to test specific opcodes is to use quest files. You can edit a quest NPC to send packets of any type, with whatever opcode you specify, and then cause the NPC to send them to your client.

The risk is that you may crash your client. A lot. However, I’ve found this to be a very effective way to prove some theories without having to modify the emulator source. For a real-world example, I’m trying to pick apart the Purchase Property window for Neighborhoods. I know the opcode and I have samples from Live but sending that packet to the RoF client always results in a crash. That packet has strings that would be allocated in memory, so my theory is that the packet format changed and that changed where the string length values were located. To test this, I zeroed out the packet and sent it.

No crash. Then I started adding values back in and was able to get the window to display. All the values are totally wrong, so I’m not done yet, but that gave me enough information to focus live debugging of the client to pinpoint the exact issue.

Put another way. This is not a trivial task (for me).

Example .LUA code:\


> local pack = Packet(0x67C9, 6, true);\
> pack:WriteInt8(0,195);\
> pack:WriteInt8(1,219);\
> pack:WriteInt8(2,248);\
> pack:WriteInt8(3,7);\
> e.other:QueuePacket(pack);

In the above, I’m sending a packet and I’m creating it by writing individual byte values. You can write different types, but I found that when working against packet captures and semi-unknown structures, it’s simpler to write the packet one byte at a time. This code can be placed into an NPC “say” handler, and then you can trigger it by hailing an NPC. Tweak a value, reload quests, and repeat.
