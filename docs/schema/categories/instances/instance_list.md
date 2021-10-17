# instance\_list

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Instance Identifier |
| zone | int | [Zone Identifier](https://eqemu.gitbook.io/server/categories/zones/zone-list) |
| version | tinyint | Version |
| is\_global | tinyint | Is Global: 0 = False, 1 = True |
| start\_time | int | Start Time UNIX Timestamp |
| duration | int | Duration in Seconds |
| never\_expires | tinyint | Never Expires: 0 = False, 1 = True |

