# completed_shared_tasks

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY29tcGxldGVkX3NoYXJlZF90YXNrcyB7XG4gICAgICAgIGludCB0YXNrX2lkXG4gICAgfVxuICAgIHRhc2tzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdGlueWludCB0eXBlXG4gICAgICAgIGludHVuc2lnbmVkIGR6X3RlbXBsYXRlX2lkXG4gICAgfVxuICAgIGNvbXBsZXRlZF9zaGFyZWRfdGFza3MgfHwtLW97IHRhc2tzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY29tcGxldGVkX3NoYXJlZF90YXNrcyB7XG4gICAgICAgIGludCB0YXNrX2lkXG4gICAgfVxuICAgIHRhc2tzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdGlueWludCB0eXBlXG4gICAgICAgIGludHVuc2lnbmVkIGR6X3RlbXBsYXRlX2lkXG4gICAgfVxuICAgIGNvbXBsZXRlZF9zaGFyZWRfdGFza3MgfHwtLW97IHRhc2tzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY29tcGxldGVkX3NoYXJlZF90YXNrcyB7XG4gICAgICAgIGludCB0YXNrX2lkXG4gICAgfVxuICAgIHRhc2tzIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdGlueWludCB0eXBlXG4gICAgICAgIGludHVuc2lnbmVkIGR6X3RlbXBsYXRlX2lkXG4gICAgfVxuICAgIGNvbXBsZXRlZF9zaGFyZWRfdGFza3MgfHwtLW97IHRhc2tzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | task_id | [tasks](../../schema/tasks/tasks.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | bigint | [Shared Task Identifier](shared_tasks.md) |
| task_id | int | [Shared Task Identifier](shared_tasks.md) |
| accepted_time | datetime | Accepted Time |
| expire_time | datetime | Expire TIme |
| completion_time | datetime | Completion TIme |
| is_locked | tinyint | Is Locked: 0 = False, 1 = True |

