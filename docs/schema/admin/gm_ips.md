# gm_ips

!!! info
	This page was last generated 2022.09.10

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ21faXBzIHtcbiAgICAgICAgaW50IGFjY291bnRfaWRcbiAgICAgICAgdmFyY2hhciBpcF9hZGRyZXNzXG4gICAgfVxuICAgIGFjY291bnQge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgfVxuICAgIGFjY291bnRfaXAge1xuICAgICAgICB2YXJjaGFyIGlwXG4gICAgICAgIGludCBhY2NpZFxuICAgIH1cbiAgICBnbV9pcHMgfHwtLW97IGFjY291bnQgOiBPbmUtdG8tT25lXG4gICAgZ21faXBzIHx8LS1veyBhY2NvdW50X2lwIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ21faXBzIHtcbiAgICAgICAgaW50IGFjY291bnRfaWRcbiAgICAgICAgdmFyY2hhciBpcF9hZGRyZXNzXG4gICAgfVxuICAgIGFjY291bnQge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgfVxuICAgIGFjY291bnRfaXAge1xuICAgICAgICB2YXJjaGFyIGlwXG4gICAgICAgIGludCBhY2NpZFxuICAgIH1cbiAgICBnbV9pcHMgfHwtLW97IGFjY291bnQgOiBPbmUtdG8tT25lXG4gICAgZ21faXBzIHx8LS1veyBhY2NvdW50X2lwIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ21faXBzIHtcbiAgICAgICAgaW50IGFjY291bnRfaWRcbiAgICAgICAgdmFyY2hhciBpcF9hZGRyZXNzXG4gICAgfVxuICAgIGFjY291bnQge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgfVxuICAgIGFjY291bnRfaXAge1xuICAgICAgICB2YXJjaGFyIGlwXG4gICAgICAgIGludCBhY2NpZFxuICAgIH1cbiAgICBnbV9pcHMgfHwtLW97IGFjY291bnQgOiBPbmUtdG8tT25lXG4gICAgZ21faXBzIHx8LS1veyBhY2NvdW50X2lwIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | account_id | [account](../../schema/account/account.md) | id |
| Has-Many | ip_address | [account_ip](../../schema/account/account_ip.md) | ip |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| name | varchar | [Character Name](../../schema/characters/character_data.md) |
| account_id | int | [Account Identifier](../../schema/account/account.md) |
| ip_address | varchar | [IP Address](../../schema/account/account_ip.md) |

