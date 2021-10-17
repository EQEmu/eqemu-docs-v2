# petitions

| Column | Data Type | Description |
| :--- | :--- | :--- |
| dib | int | Unknown |
| petid | int | Unique Petition Entry Identifier |
| charname | varchar | [Character Name](../../../schema/categories/admin/character_data.md) |
| accountname | varchar | [Account Name](../../../schema/categories/admin/account.md) |
| lastgm | varchar | Last GM |
| petitiontext | text | Petition Text |
| gmtext | text | GM Text |
| zone | varchar | [Zone Short Name](../../../../categories/zones/zone-list) |
| urgency | int | Urgency |
| charclass | int | [Character Class](../../../../categories/player/class-list) |
| charrace | int | [Character Race](../../../../categories/npc/race-list) |
| charlevel | int | Character Level |
| checkouts | int | Checkouts |
| unavailables | int | Unavailables |
| ischeckedout | tinyint | Is Checked Out: 0 = False, 1 = True |
| senttime | bigint | Sent Time UNIX Timestamp |

