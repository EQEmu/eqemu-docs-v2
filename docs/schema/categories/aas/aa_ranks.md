# aa\_ranks

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | [AA Identifier](aa_ability.md) |
| upper\_hotkey\_sid | int | Upper Hotkey SID |
| lower\_hotkey\_sid | int | Lower Hotkey SID |
| title\_sid | int | Title SID |
| desc\_sid | int | Description SID |
| cost | int | Cost in AA Points |
| level\_req | int | Level Required |
| spell | int | [Spell Identifier](https://github.com/EQEmu/docs-db-schema/tree/e0eb157dbf5563b03c0faf391abc87ec69239f4a/docs/schema/categories/aas/spells_new.md) |
| spell\_type | int | [Spell Type](https://eqemu.gitbook.io/server/categories/spells/spell-types) |
| recast\_time | int | Recast Timer |
| expansion | int | [Expansion Identifier](https://eqemu.gitbook.io/server/categories/operation/expansion-list) |
| prev\_id | int | Previous Rank Identifier |
| next\_id | int | Next Rank Identifier |

