# bot_inspect_messages

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgIH1cbiAgICBib3RfZGF0YSB7XG4gICAgICAgIHZhcmNoYXIgYm90X2lkXG4gICAgICAgIHZhcmNoYXIgc3BlbGxzX2lkXG4gICAgICAgIHZhcmNoYXIgem9uZV9pZFxuICAgICAgICB2YXJjaGFyIG93bmVyX2lkXG4gICAgfVxuICAgIGJvdF9pbnNwZWN0X21lc3NhZ2VzIHx8LS1veyBib3RfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgIH1cbiAgICBib3RfZGF0YSB7XG4gICAgICAgIHZhcmNoYXIgYm90X2lkXG4gICAgICAgIHZhcmNoYXIgc3BlbGxzX2lkXG4gICAgICAgIHZhcmNoYXIgem9uZV9pZFxuICAgICAgICB2YXJjaGFyIG93bmVyX2lkXG4gICAgfVxuICAgIGJvdF9pbnNwZWN0X21lc3NhZ2VzIHx8LS1veyBib3RfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2luc3BlY3RfbWVzc2FnZXMge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgIH1cbiAgICBib3RfZGF0YSB7XG4gICAgICAgIHZhcmNoYXIgYm90X2lkXG4gICAgICAgIHZhcmNoYXIgc3BlbGxzX2lkXG4gICAgICAgIHZhcmNoYXIgem9uZV9pZFxuICAgICAgICB2YXJjaGFyIG93bmVyX2lkXG4gICAgfVxuICAgIGJvdF9pbnNwZWN0X21lc3NhZ2VzIHx8LS1veyBib3RfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_id | [bot_data](../../schema/bots/bot_data.md) | bot_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| bot_id | int | [Bot Identifier](bot_data.md) |
| inspect_message | varchar | Inspect Message |

