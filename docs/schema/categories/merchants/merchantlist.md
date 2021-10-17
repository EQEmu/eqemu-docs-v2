# merchantlist

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merchantid | int | Merchant Identifier |
| slot | int | Slot |
| item | int | [Item Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/merchants/items.md) |
| faction\_required | smallint | Faction Required |
| level\_required | tinyint | Level Required |
| alt\_currency\_cost | smallint | [Alternate Currency Cost](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/merchants/alternate_currency.md) |
| classes\_required | int | [Classes Required](https://eqemu.gitbook.io/server/categories/player/class-list) |
| probability | int | Probability: 0 = Never, 100 = Always |

