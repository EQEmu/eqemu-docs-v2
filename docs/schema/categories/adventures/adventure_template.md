# adventure\_template

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | [Adventure Identifier](adventure_details.md) |
| zone | varchar | [Zone Short Name](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| zone\_version | tinyint | Zone Version |
| is\_hard | tinyint | Hard: 0 = False, 1 = True |
| is\_raid | tinyint | Raid: 0 = False, 1 = True |
| min\_level | tinyint | Minimum Level |
| max\_level | tinyint | Maximum Level |
| type | tinyint | Type |
| type\_data | int | Type Data |
| type\_count | smallint | Type Count |
| assa\_x | float | Assassination X Coordinate |
| assa\_y | float | Assassination Y Coordinate |
| assa\_z | float | Assassination Z Coordinate |
| assa\_h | float | Assassination Heading Coordinate |
| text | varchar | Text |
| duration | int | Duration |
| zone\_in\_time | int | Zone In Duration |
| win\_points | smallint | LDoN Points for Winning |
| lose\_points | smallint | LDoN Points for Losing |
| theme | tinyint | [LDoN Theme](https://eqemu.gitbook.io/server/categories/zones/ldon-themes) |
| zone\_in\_zone\_id | smallint | [Zone In Zone Identifier](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| zone\_in\_x | float | Zone In X Coordinate |
| zone\_in\_y | float | Zone In Y Coordinate |
| zone\_in\_object\_id | smallint | Zone In Object Identifier |
| dest\_x | float | Destination X Coordinate |
| dest\_y | float | Destination Y Coordinate |
| dest\_z | float | Destination Z Coordinate |
| dest\_h | float | Destination Heading Coordinate |
| graveyard\_zone\_id | int | [Zone Identifier](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| graveyard\_x | float | Graveyard X Coordinate |
| graveyard\_y | float | Graveyard Y Coordinate |
| graveyard\_z | float | Graveyard Z Coordinate |
| graveyard\_radius | float | Graveyard Radius |

