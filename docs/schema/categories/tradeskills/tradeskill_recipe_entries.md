# tradeskill\_recipe\_entries

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Tradeskill Recipe Entry Identifier |
| recipe\_id | int | [Unique Tradeskill Recipe Identifier](tradeskill_recipe.md) |
| item\_id | int | [Item Identifier](../../../schema/categories/tradeskills/items.md) |
| successcount | tinyint | Success Count |
| failcount | tinyint | Fail Count |
| componentcount | tinyint | Component Count |
| salvagecount | tinyint | Salvage Count |
| iscontainer | tinyint | Is Container: 0 = False, 1 = True |

