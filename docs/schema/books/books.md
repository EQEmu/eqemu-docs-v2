# books

!!! info
	This page was last generated 2022.05.11

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int |  |
| name | varchar | Unique Book Identifier |
| txtfile | text | The text in the book. ` Represents line spaces, `` is two line spaces, ``` is three line spaces, etc. (13 lines per book page) |
| language | int | [Language](../../../../server/player/languages) |

