# ip_exemptions

!!! info
	This page was last generated 2022.09.16

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgaXBfZXhlbXB0aW9ucyB7XG4gICAgICAgIHZhcmNoYXIgZXhlbXB0aW9uX2lwXG4gICAgfVxuICAgIGFjY291bnRfaXAge1xuICAgICAgICB2YXJjaGFyIGlwXG4gICAgICAgIGludCBhY2NpZFxuICAgIH1cbiAgICBpcF9leGVtcHRpb25zIHx8LS1veyBhY2NvdW50X2lwIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgaXBfZXhlbXB0aW9ucyB7XG4gICAgICAgIHZhcmNoYXIgZXhlbXB0aW9uX2lwXG4gICAgfVxuICAgIGFjY291bnRfaXAge1xuICAgICAgICB2YXJjaGFyIGlwXG4gICAgICAgIGludCBhY2NpZFxuICAgIH1cbiAgICBpcF9leGVtcHRpb25zIHx8LS1veyBhY2NvdW50X2lwIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgaXBfZXhlbXB0aW9ucyB7XG4gICAgICAgIHZhcmNoYXIgZXhlbXB0aW9uX2lwXG4gICAgfVxuICAgIGFjY291bnRfaXAge1xuICAgICAgICB2YXJjaGFyIGlwXG4gICAgICAgIGludCBhY2NpZFxuICAgIH1cbiAgICBpcF9leGVtcHRpb25zIHx8LS1veyBhY2NvdW50X2lwIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | exemption_ip | [account_ip](../../schema/account/account_ip.md) | ip |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| exemption_id | int | Exemption Identifier |
| exemption_ip | varchar | [Exemption IP Address](../../schema/account/account_ip.md) |
| exemption_amount | int | Exemption Amount |

