# bot_inventories

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2ludmVudG9yaWVzIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBpdGVtX2lkXG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF8xXG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF8yXG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF8zXG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF80XG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF81XG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF82XG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciB6b25lX2lkXG4gICAgICAgIHZhcmNoYXIgb3duZXJfaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludCBib29rXG4gICAgICAgIGludCByZWNhc3R0eXBlXG4gICAgICAgIGludCBpY29uXG4gICAgICAgIHNtYWxsaW50IGJhcmRlZmZlY3RcbiAgICAgICAgaW50IGNsaWNrZWZmZWN0XG4gICAgICAgIGludCBmb2N1c2VmZmVjdFxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgfVxuICAgIGJvdF9pbnZlbnRvcmllcyB8fC0tb3sgYm90X2RhdGEgOiBPbmUtdG8tT25lXG4gICAgYm90X2ludmVudG9yaWVzIHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcbiAgICBib3RfaW52ZW50b3JpZXMgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuICAgIGJvdF9pbnZlbnRvcmllcyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG4gICAgYm90X2ludmVudG9yaWVzIHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcbiAgICBib3RfaW52ZW50b3JpZXMgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuICAgIGJvdF9pbnZlbnRvcmllcyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG4gICAgYm90X2ludmVudG9yaWVzIHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2ludmVudG9yaWVzIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBpdGVtX2lkXG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF8xXG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF8yXG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF8zXG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF80XG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF81XG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF82XG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciB6b25lX2lkXG4gICAgICAgIHZhcmNoYXIgb3duZXJfaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludCBib29rXG4gICAgICAgIGludCByZWNhc3R0eXBlXG4gICAgICAgIGludCBpY29uXG4gICAgICAgIHNtYWxsaW50IGJhcmRlZmZlY3RcbiAgICAgICAgaW50IGNsaWNrZWZmZWN0XG4gICAgICAgIGludCBmb2N1c2VmZmVjdFxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgfVxuICAgIGJvdF9pbnZlbnRvcmllcyB8fC0tb3sgYm90X2RhdGEgOiBPbmUtdG8tT25lXG4gICAgYm90X2ludmVudG9yaWVzIHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcbiAgICBib3RfaW52ZW50b3JpZXMgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuICAgIGJvdF9pbnZlbnRvcmllcyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG4gICAgYm90X2ludmVudG9yaWVzIHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcbiAgICBib3RfaW52ZW50b3JpZXMgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuICAgIGJvdF9pbnZlbnRvcmllcyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG4gICAgYm90X2ludmVudG9yaWVzIHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2ludmVudG9yaWVzIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBpdGVtX2lkXG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF8xXG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF8yXG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF8zXG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF80XG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF81XG4gICAgICAgIHZhcmNoYXIgYXVnbWVudF82XG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICAgICAgdmFyY2hhciBzcGVsbHNfaWRcbiAgICAgICAgdmFyY2hhciB6b25lX2lkXG4gICAgICAgIHZhcmNoYXIgb3duZXJfaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludCBib29rXG4gICAgICAgIGludCByZWNhc3R0eXBlXG4gICAgICAgIGludCBpY29uXG4gICAgICAgIHNtYWxsaW50IGJhcmRlZmZlY3RcbiAgICAgICAgaW50IGNsaWNrZWZmZWN0XG4gICAgICAgIGludCBmb2N1c2VmZmVjdFxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgfVxuICAgIGJvdF9pbnZlbnRvcmllcyB8fC0tb3sgYm90X2RhdGEgOiBPbmUtdG8tT25lXG4gICAgYm90X2ludmVudG9yaWVzIHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcbiAgICBib3RfaW52ZW50b3JpZXMgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuICAgIGJvdF9pbnZlbnRvcmllcyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG4gICAgYm90X2ludmVudG9yaWVzIHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcbiAgICBib3RfaW52ZW50b3JpZXMgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuICAgIGJvdF9pbnZlbnRvcmllcyB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG4gICAgYm90X2ludmVudG9yaWVzIHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_id | [bot_data](../../schema/bots/bot_data.md) | bot_id |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |
| One-to-One | augment_1 | [items](../../schema/items/items.md) | id |
| One-to-One | augment_2 | [items](../../schema/items/items.md) | id |
| One-to-One | augment_3 | [items](../../schema/items/items.md) | id |
| One-to-One | augment_4 | [items](../../schema/items/items.md) | id |
| One-to-One | augment_5 | [items](../../schema/items/items.md) | id |
| One-to-One | augment_6 | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| inventories_index | int | Unique Bot Inventory Identifier |
| bot_id | int | [Bot Identifier](bot_data.md) |
| slot_id | mediumint | [Slot Identifier](../../../../categories/inventory/inventory-slots) |
| item_id | int | [Item Identifier](../items/items.md) |
| inst_charges | smallint | Charges |
| inst_color | int | Color |
| inst_no_drop | tinyint | No Drop: 0 = False, 1=  True |
| inst_custom_data | text | Custom Data |
| ornament_icon | int | Ornamentation Icon |
| ornament_id_file | int | Ornamentation Item Texture |
| ornament_hero_model | int | Ornamentation Hero's Forge Model |
| augment_1 | mediumint | Augment Slot 1 |
| augment_2 | mediumint | Augment Slot 2 |
| augment_3 | mediumint | Augment Slot 3 |
| augment_4 | mediumint | Augment Slot 4 |
| augment_5 | mediumint | Augment Slot 5 |
| augment_6 | mediumint | Augment Slot 6 |

