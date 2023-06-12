# qs_player_move_record_entries

!!! info
	This page was last generated 2023.06.12

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX21vdmVfcmVjb3JkX2VudHJpZXMge1xuICAgICAgICBpbnQgYXVnXzFcbiAgICAgICAgaW50IGF1Z18yXG4gICAgICAgIGludCBhdWdfM1xuICAgICAgICBpbnQgYXVnXzRcbiAgICAgICAgaW50IGF1Z181XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgYm9va1xuICAgICAgICBpbnQgaWNvblxuICAgICAgICBzbWFsbGludCBiYXJkZWZmZWN0XG4gICAgICAgIGludCBjbGlja2VmZmVjdFxuICAgICAgICBpbnQgZm9jdXNlZmZlY3RcbiAgICAgICAgaW50IHByb2NlZmZlY3RcbiAgICAgICAgaW50IHNjcm9sbGVmZmVjdFxuICAgICAgICBpbnQgd29ybmVmZmVjdFxuICAgICAgICBpbnQgcmVjYXN0dHlwZVxuICAgIH1cbiAgICBxc19wbGF5ZXJfbW92ZV9yZWNvcmRfZW50cmllcyB8fC0tb3sgaXRlbXMgOiBcIk9uZS10by1PbmVcIlxuICAgIHFzX3BsYXllcl9tb3ZlX3JlY29yZF9lbnRyaWVzIHx8LS1veyBpdGVtcyA6IFwiT25lLXRvLU9uZVwiXG4gICAgcXNfcGxheWVyX21vdmVfcmVjb3JkX2VudHJpZXMgfHwtLW97IGl0ZW1zIDogXCJPbmUtdG8tT25lXCJcbiAgICBxc19wbGF5ZXJfbW92ZV9yZWNvcmRfZW50cmllcyB8fC0tb3sgaXRlbXMgOiBcIk9uZS10by1PbmVcIlxuICAgIHFzX3BsYXllcl9tb3ZlX3JlY29yZF9lbnRyaWVzIHx8LS1veyBpdGVtcyA6IFwiT25lLXRvLU9uZVwiXG4gICAgcXNfcGxheWVyX21vdmVfcmVjb3JkX2VudHJpZXMgfHwtLW97IGl0ZW1zIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX21vdmVfcmVjb3JkX2VudHJpZXMge1xuICAgICAgICBpbnQgYXVnXzFcbiAgICAgICAgaW50IGF1Z18yXG4gICAgICAgIGludCBhdWdfM1xuICAgICAgICBpbnQgYXVnXzRcbiAgICAgICAgaW50IGF1Z181XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgYm9va1xuICAgICAgICBpbnQgaWNvblxuICAgICAgICBzbWFsbGludCBiYXJkZWZmZWN0XG4gICAgICAgIGludCBjbGlja2VmZmVjdFxuICAgICAgICBpbnQgZm9jdXNlZmZlY3RcbiAgICAgICAgaW50IHByb2NlZmZlY3RcbiAgICAgICAgaW50IHNjcm9sbGVmZmVjdFxuICAgICAgICBpbnQgd29ybmVmZmVjdFxuICAgICAgICBpbnQgcmVjYXN0dHlwZVxuICAgIH1cbiAgICBxc19wbGF5ZXJfbW92ZV9yZWNvcmRfZW50cmllcyB8fC0tb3sgaXRlbXMgOiBcIk9uZS10by1PbmVcIlxuICAgIHFzX3BsYXllcl9tb3ZlX3JlY29yZF9lbnRyaWVzIHx8LS1veyBpdGVtcyA6IFwiT25lLXRvLU9uZVwiXG4gICAgcXNfcGxheWVyX21vdmVfcmVjb3JkX2VudHJpZXMgfHwtLW97IGl0ZW1zIDogXCJPbmUtdG8tT25lXCJcbiAgICBxc19wbGF5ZXJfbW92ZV9yZWNvcmRfZW50cmllcyB8fC0tb3sgaXRlbXMgOiBcIk9uZS10by1PbmVcIlxuICAgIHFzX3BsYXllcl9tb3ZlX3JlY29yZF9lbnRyaWVzIHx8LS1veyBpdGVtcyA6IFwiT25lLXRvLU9uZVwiXG4gICAgcXNfcGxheWVyX21vdmVfcmVjb3JkX2VudHJpZXMgfHwtLW97IGl0ZW1zIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcXNfcGxheWVyX21vdmVfcmVjb3JkX2VudHJpZXMge1xuICAgICAgICBpbnQgYXVnXzFcbiAgICAgICAgaW50IGF1Z18yXG4gICAgICAgIGludCBhdWdfM1xuICAgICAgICBpbnQgYXVnXzRcbiAgICAgICAgaW50IGF1Z181XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgYm9va1xuICAgICAgICBpbnQgaWNvblxuICAgICAgICBzbWFsbGludCBiYXJkZWZmZWN0XG4gICAgICAgIGludCBjbGlja2VmZmVjdFxuICAgICAgICBpbnQgZm9jdXNlZmZlY3RcbiAgICAgICAgaW50IHByb2NlZmZlY3RcbiAgICAgICAgaW50IHNjcm9sbGVmZmVjdFxuICAgICAgICBpbnQgd29ybmVmZmVjdFxuICAgICAgICBpbnQgcmVjYXN0dHlwZVxuICAgIH1cbiAgICBxc19wbGF5ZXJfbW92ZV9yZWNvcmRfZW50cmllcyB8fC0tb3sgaXRlbXMgOiBcIk9uZS10by1PbmVcIlxuICAgIHFzX3BsYXllcl9tb3ZlX3JlY29yZF9lbnRyaWVzIHx8LS1veyBpdGVtcyA6IFwiT25lLXRvLU9uZVwiXG4gICAgcXNfcGxheWVyX21vdmVfcmVjb3JkX2VudHJpZXMgfHwtLW97IGl0ZW1zIDogXCJPbmUtdG8tT25lXCJcbiAgICBxc19wbGF5ZXJfbW92ZV9yZWNvcmRfZW50cmllcyB8fC0tb3sgaXRlbXMgOiBcIk9uZS10by1PbmVcIlxuICAgIHFzX3BsYXllcl9tb3ZlX3JlY29yZF9lbnRyaWVzIHx8LS1veyBpdGVtcyA6IFwiT25lLXRvLU9uZVwiXG4gICAgcXNfcGxheWVyX21vdmVfcmVjb3JkX2VudHJpZXMgfHwtLW97IGl0ZW1zIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | aug_1 | [items](../../schema/items/items.md) | id |
| One-to-One | aug_2 | [items](../../schema/items/items.md) | id |
| One-to-One | aug_3 | [items](../../schema/items/items.md) | id |
| One-to-One | aug_4 | [items](../../schema/items/items.md) | id |
| One-to-One | aug_5 | [items](../../schema/items/items.md) | id |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| event_id | int | Unique Event Identifier |
| from_slot | mediumint | [From Slot Identifier](../../../../server/inventory/inventory-slots) |
| to_slot | mediumint | [To Slot Identifier](../../../../server/inventory/inventory-slots) |
| item_id | int | [Item Identifier](../../schema/items/items.md) |
| charges | mediumint | Charges |
| aug_1 | int | Augment Slot 1 |
| aug_2 | int | Augment Slot 2 |
| aug_3 | int | Augment Slot 3 |
| aug_4 | int | Augment Slot 4 |
| aug_5 | int | Augment Slot 5 |

