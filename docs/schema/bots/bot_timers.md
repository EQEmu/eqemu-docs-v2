# bot_timers

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3RpbWVycyB7XG4gICAgICAgIHZhcmNoYXIgYm90X2RhdGFcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgICAgIGludHVuc2lnbmVkIHNwZWxsc19pZFxuICAgICAgICBzbWFsbGludCB6b25lX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG93bmVyX2lkXG4gICAgfVxuICAgIGJvdF90aW1lcnMgfHwtLW97IGJvdF9kYXRhIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3RpbWVycyB7XG4gICAgICAgIHZhcmNoYXIgYm90X2RhdGFcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgICAgIGludHVuc2lnbmVkIHNwZWxsc19pZFxuICAgICAgICBzbWFsbGludCB6b25lX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG93bmVyX2lkXG4gICAgfVxuICAgIGJvdF90aW1lcnMgfHwtLW97IGJvdF9kYXRhIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3RpbWVycyB7XG4gICAgICAgIHZhcmNoYXIgYm90X2RhdGFcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgYm90X2lkXG4gICAgICAgIGludHVuc2lnbmVkIHNwZWxsc19pZFxuICAgICAgICBzbWFsbGludCB6b25lX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG93bmVyX2lkXG4gICAgfVxuICAgIGJvdF90aW1lcnMgfHwtLW97IGJvdF9kYXRhIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_data | [bot_data](../../schema/bots/bot_data.md) | bot_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| bot_id | int | [Unique Bot Identifier](bot_data.md) |
| timer_id | int | Timer Identifier |
| timer_value | int | Timer Expiration |

