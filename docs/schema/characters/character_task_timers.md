# character_task_timers

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX3Rhc2tfdGltZXJzIHtcbiAgICAgICAgaW50dW5zaWduZWQgY2hhcmFjdGVyX2lkXG4gICAgICAgIGludHVuc2lnbmVkIHRhc2tfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIHRhc2tzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdGlueWludCB0eXBlXG4gICAgfVxuICAgIGNoYXJhY3Rlcl90YXNrX3RpbWVycyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgY2hhcmFjdGVyX3Rhc2tfdGltZXJzIHx8LS1veyB0YXNrcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX3Rhc2tfdGltZXJzIHtcbiAgICAgICAgaW50dW5zaWduZWQgY2hhcmFjdGVyX2lkXG4gICAgICAgIGludHVuc2lnbmVkIHRhc2tfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIHRhc2tzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdGlueWludCB0eXBlXG4gICAgfVxuICAgIGNoYXJhY3Rlcl90YXNrX3RpbWVycyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgY2hhcmFjdGVyX3Rhc2tfdGltZXJzIHx8LS1veyB0YXNrcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX3Rhc2tfdGltZXJzIHtcbiAgICAgICAgaW50dW5zaWduZWQgY2hhcmFjdGVyX2lkXG4gICAgICAgIGludHVuc2lnbmVkIHRhc2tfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lX2lkXG4gICAgfVxuICAgIHRhc2tzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdGlueWludCB0eXBlXG4gICAgfVxuICAgIGNoYXJhY3Rlcl90YXNrX3RpbWVycyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgY2hhcmFjdGVyX3Rhc2tfdGltZXJzIHx8LS1veyB0YXNrcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | character_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | task_id | [tasks](../../schema/tasks/tasks.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Task Timer Identifier |
| character_id | int | [Character Identifer](character_data.md) |
| task_id | int | [Task Identifier](../../schema/tasks/tasks.md) |
| timer_type | int | Timer Type (0 = Replay, 1 = Request) |
| expire_time | datetime | Expire Time |

