# bot_timers

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3RpbWVycyB7XG4gICAgICAgIHZhcmNoYXIgYm90X2RhdGFcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIHNwZWxsc19pZFxuICAgICAgICB2YXJjaGFyIHpvbmVfaWRcbiAgICAgICAgdmFyY2hhciBvd25lcl9pZFxuICAgIH1cbiAgICBib3RfdGltZXJzIHx8LS1veyBib3RfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3RpbWVycyB7XG4gICAgICAgIHZhcmNoYXIgYm90X2RhdGFcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIHNwZWxsc19pZFxuICAgICAgICB2YXJjaGFyIHpvbmVfaWRcbiAgICAgICAgdmFyY2hhciBvd25lcl9pZFxuICAgIH1cbiAgICBib3RfdGltZXJzIHx8LS1veyBib3RfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X3RpbWVycyB7XG4gICAgICAgIHZhcmNoYXIgYm90X2RhdGFcbiAgICAgICAgdmFyY2hhciBib3RfaWRcbiAgICB9XG4gICAgYm90X2RhdGEge1xuICAgICAgICB2YXJjaGFyIGJvdF9pZFxuICAgICAgICB2YXJjaGFyIHNwZWxsc19pZFxuICAgICAgICB2YXJjaGFyIHpvbmVfaWRcbiAgICAgICAgdmFyY2hhciBvd25lcl9pZFxuICAgIH1cbiAgICBib3RfdGltZXJzIHx8LS1veyBib3RfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


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

