# bot_heal_rotation_members

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2hlYWxfcm90YXRpb25fbWVtYmVycyB7XG4gICAgICAgIGludHVuc2lnbmVkIGhlYWxfcm90YXRpb25faW5kZXhcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgICAgIGludHVuc2lnbmVkIHNwZWxsc19pZFxuICAgICAgICBpbnR1bnNpZ25lZCBvd25lcl9pZFxuICAgICAgICBzbWFsbGludCB6b25lX2lkXG4gICAgfVxuICAgIGJvdF9oZWFsX3JvdGF0aW9ucyB7XG4gICAgICAgIGludHVuc2lnbmVkIGhlYWxfcm90YXRpb25faW5kZXhcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgfVxuICAgIGJvdF9oZWFsX3JvdGF0aW9uX21lbWJlcnMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuICAgIGJvdF9oZWFsX3JvdGF0aW9uX21lbWJlcnMgfHwtLW97IGJvdF9oZWFsX3JvdGF0aW9ucyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2hlYWxfcm90YXRpb25fbWVtYmVycyB7XG4gICAgICAgIGludHVuc2lnbmVkIGhlYWxfcm90YXRpb25faW5kZXhcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgICAgIGludHVuc2lnbmVkIHNwZWxsc19pZFxuICAgICAgICBpbnR1bnNpZ25lZCBvd25lcl9pZFxuICAgICAgICBzbWFsbGludCB6b25lX2lkXG4gICAgfVxuICAgIGJvdF9oZWFsX3JvdGF0aW9ucyB7XG4gICAgICAgIGludHVuc2lnbmVkIGhlYWxfcm90YXRpb25faW5kZXhcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgfVxuICAgIGJvdF9oZWFsX3JvdGF0aW9uX21lbWJlcnMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuICAgIGJvdF9oZWFsX3JvdGF0aW9uX21lbWJlcnMgfHwtLW97IGJvdF9oZWFsX3JvdGF0aW9ucyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2hlYWxfcm90YXRpb25fbWVtYmVycyB7XG4gICAgICAgIGludHVuc2lnbmVkIGhlYWxfcm90YXRpb25faW5kZXhcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgICAgIGludHVuc2lnbmVkIHNwZWxsc19pZFxuICAgICAgICBpbnR1bnNpZ25lZCBvd25lcl9pZFxuICAgICAgICBzbWFsbGludCB6b25lX2lkXG4gICAgfVxuICAgIGJvdF9oZWFsX3JvdGF0aW9ucyB7XG4gICAgICAgIGludHVuc2lnbmVkIGhlYWxfcm90YXRpb25faW5kZXhcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgfVxuICAgIGJvdF9oZWFsX3JvdGF0aW9uX21lbWJlcnMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuICAgIGJvdF9oZWFsX3JvdGF0aW9uX21lbWJlcnMgfHwtLW97IGJvdF9oZWFsX3JvdGF0aW9ucyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_id | [bot_data](../../schema/bots/bot_data.md) | bot_id |
| One-to-One | heal_rotation_index | [bot_heal_rotations](../../schema/bots/bot_heal_rotations.md) | heal_rotation_index |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| member_index | int | Unique Bot Heal Rotation Member Identifier |
| heal_rotation_index | int | [Heal Rotation Identifier](bot_heal_rotations.md) |
| bot_id | int | [Bot Identifier](bot_data.md) |

