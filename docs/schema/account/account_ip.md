# account_ip

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWNjb3VudF9pcCB7XG4gICAgICAgIGludCBhY2NpZFxuICAgICAgICB2YXJjaGFyIGlwXG4gICAgfVxuICAgIGJhbm5lZF9pcHMge1xuICAgICAgICB2YXJjaGFyIGlwX2FkZHJlc3NcbiAgICB9XG4gICAgZ21faXBzIHtcbiAgICAgICAgaW50IGFjY291bnRfaWRcbiAgICAgICAgdmFyY2hhciBpcF9hZGRyZXNzXG4gICAgfVxuICAgIGlwX2V4ZW1wdGlvbnMge1xuICAgICAgICB2YXJjaGFyIGV4ZW1wdGlvbl9pcFxuICAgIH1cbiAgICBhY2NvdW50X2lwIHx8LS1veyBiYW5uZWRfaXBzIDogT25lLXRvLU9uZVxuICAgIGFjY291bnRfaXAgfHwtLW97IGdtX2lwcyA6IE9uZS10by1PbmVcbiAgICBhY2NvdW50X2lwIHx8LS1veyBpcF9leGVtcHRpb25zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWNjb3VudF9pcCB7XG4gICAgICAgIGludCBhY2NpZFxuICAgICAgICB2YXJjaGFyIGlwXG4gICAgfVxuICAgIGJhbm5lZF9pcHMge1xuICAgICAgICB2YXJjaGFyIGlwX2FkZHJlc3NcbiAgICB9XG4gICAgZ21faXBzIHtcbiAgICAgICAgaW50IGFjY291bnRfaWRcbiAgICAgICAgdmFyY2hhciBpcF9hZGRyZXNzXG4gICAgfVxuICAgIGlwX2V4ZW1wdGlvbnMge1xuICAgICAgICB2YXJjaGFyIGV4ZW1wdGlvbl9pcFxuICAgIH1cbiAgICBhY2NvdW50X2lwIHx8LS1veyBiYW5uZWRfaXBzIDogT25lLXRvLU9uZVxuICAgIGFjY291bnRfaXAgfHwtLW97IGdtX2lwcyA6IE9uZS10by1PbmVcbiAgICBhY2NvdW50X2lwIHx8LS1veyBpcF9leGVtcHRpb25zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWNjb3VudF9pcCB7XG4gICAgICAgIGludCBhY2NpZFxuICAgICAgICB2YXJjaGFyIGlwXG4gICAgfVxuICAgIGJhbm5lZF9pcHMge1xuICAgICAgICB2YXJjaGFyIGlwX2FkZHJlc3NcbiAgICB9XG4gICAgZ21faXBzIHtcbiAgICAgICAgaW50IGFjY291bnRfaWRcbiAgICAgICAgdmFyY2hhciBpcF9hZGRyZXNzXG4gICAgfVxuICAgIGlwX2V4ZW1wdGlvbnMge1xuICAgICAgICB2YXJjaGFyIGV4ZW1wdGlvbl9pcFxuICAgIH1cbiAgICBhY2NvdW50X2lwIHx8LS1veyBiYW5uZWRfaXBzIDogT25lLXRvLU9uZVxuICAgIGFjY291bnRfaXAgfHwtLW97IGdtX2lwcyA6IE9uZS10by1PbmVcbiAgICBhY2NvdW50X2lwIHx8LS1veyBpcF9leGVtcHRpb25zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | ip | [banned_ips](../../schema/admin/banned_ips.md) | ip_address |
| One-to-One | ip | [gm_ips](../../schema/admin/gm_ips.md) | ip_address |
| One-to-One | ip | [ip_exemptions](../../schema/admin/ip_exemptions.md) | exemption_ip |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| accid | int | [Account Identifier](account.md) |
| ip | varchar | IP Address |
| count | int | Number of times logged in from this IP |
| lastused | timestamp | Timestamp of when account was last logged in |

