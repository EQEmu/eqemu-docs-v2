# account_flags

!!! info
	This page was last generated 2022.09.16

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWNjb3VudF9mbGFncyB7XG4gICAgICAgIGludHVuc2lnbmVkIHBfYWNjaWRcbiAgICB9XG4gICAgYWNjb3VudCB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgYWNjb3VudF9mbGFncyB8fC0tb3sgYWNjb3VudCA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWNjb3VudF9mbGFncyB7XG4gICAgICAgIGludHVuc2lnbmVkIHBfYWNjaWRcbiAgICB9XG4gICAgYWNjb3VudCB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgYWNjb3VudF9mbGFncyB8fC0tb3sgYWNjb3VudCA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWNjb3VudF9mbGFncyB7XG4gICAgICAgIGludHVuc2lnbmVkIHBfYWNjaWRcbiAgICB9XG4gICAgYWNjb3VudCB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgYWNjb3VudF9mbGFncyB8fC0tb3sgYWNjb3VudCA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | p_accid | [account](../../schema/account/account.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| p_accid | int | [Account Identifier](account.md) |
| p_flag | varchar | Name |
| p_value | varchar | Value |

