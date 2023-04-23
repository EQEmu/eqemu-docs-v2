# aa_rank_prereqs

!!! info
	This page was last generated 2023.01.30

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWFfcmFua19wcmVyZXFzIHtcbiAgICAgICAgaW50IGFhX2lkXG4gICAgICAgIGludHVuc2lnbmVkIHJhbmtfaWRcbiAgICB9XG4gICAgYWFfcmFua3Mge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBhYV9yYW5rX3ByZXJlcXMgfHwtLW97IGFhX3JhbmtzIDogXCJPbmUtdG8tT25lXCJcbiAgICBhYV9yYW5rX3ByZXJlcXMgfHwtLW97IGFhX3JhbmtzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWFfcmFua19wcmVyZXFzIHtcbiAgICAgICAgaW50IGFhX2lkXG4gICAgICAgIGludHVuc2lnbmVkIHJhbmtfaWRcbiAgICB9XG4gICAgYWFfcmFua3Mge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBhYV9yYW5rX3ByZXJlcXMgfHwtLW97IGFhX3JhbmtzIDogXCJPbmUtdG8tT25lXCJcbiAgICBhYV9yYW5rX3ByZXJlcXMgfHwtLW97IGFhX3JhbmtzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWFfcmFua19wcmVyZXFzIHtcbiAgICAgICAgaW50IGFhX2lkXG4gICAgICAgIGludHVuc2lnbmVkIHJhbmtfaWRcbiAgICB9XG4gICAgYWFfcmFua3Mge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBhYV9yYW5rX3ByZXJlcXMgfHwtLW97IGFhX3JhbmtzIDogXCJPbmUtdG8tT25lXCJcbiAgICBhYV9yYW5rX3ByZXJlcXMgfHwtLW97IGFhX3JhbmtzIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | aa_id | [aa_ranks](../../schema/aas/aa_ranks.md) | id |
| One-to-One | rank_id | [aa_ranks](../../schema/aas/aa_ranks.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| rank_id | int | [Rank Identifier](aa_ranks.md) |
| aa_id | int | [AA Identifier](aa_ability.md) |
| points | int | Cost in AA Points |

