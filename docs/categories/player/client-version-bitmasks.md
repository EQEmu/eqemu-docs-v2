# Client Version Bitmasks

| **Token** | **Value \(int\)** | **Value \(hex\)** |
| :--- | :--- | :--- |
| bitUnknown | 0 | 0x00000000 |
| bitClient62\* | 1 | 0x00000001 |
| bitTitanium | 2 | 0x00000002 |
| bitSoF | 4 | 0x00000004 |
| bitSoD | 8 | 0x00000008 |
| bitUF | 16 | 0x00000010 |
| bitRoF | 32 | 0x00000020 |
| bitRoF2 | 64 | 0x00000040 |
| maskUnknown | 0 | 0x00000000 |
| maskTitaniumAndEarlier | 3 | 0x00000003 |
| maskSoFAndEarlier | 7 | 0x00000007 |
| maskSoDAndEarlier | 15 | 0x0000000F |
| maskUFAndEarlier | 31 | 0x0000001F |
| maskRoFAndEarlier | 63 | 0x0000003F |
| maskSoFAndLater | 4294967292 | 0xFFFFFFFC |
| maskSoDAndLater | 4294967288 | 0xFFFFFFF8 |
| maskUFAndLater | 4294967280 | 0xFFFFFFF0 |
| maskRoFAndLater | 4294967264 | 0xFFFFFFE0 |
| maskRoF2AndLater | 4294967232 | 0xFFFFFFC0 |
| maskAllClients | 4294967295 | 0xFFFFFFFF |

`*` Client is no longer supported. Value is kept as a placeholder.

