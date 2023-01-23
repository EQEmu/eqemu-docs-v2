# eventlog

!!! info
	This page was last generated 2023.01.22

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXZlbnRsb2cge1xuICAgICAgICB2YXJjaGFyIGNoYXJuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIGFjY291bnRpZFxuICAgICAgICB2YXJjaGFyIGFjY291bnRuYW1lXG4gICAgfVxuICAgIGFjY291bnQge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICB9XG4gICAgZXZlbnRsb2cgfHwtLW97IGFjY291bnQgOiBPbmUtdG8tT25lXG4gICAgZXZlbnRsb2cgfHwtLW97IGFjY291bnQgOiBPbmUtdG8tT25lXG4gICAgZXZlbnRsb2cgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXZlbnRsb2cge1xuICAgICAgICB2YXJjaGFyIGNoYXJuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIGFjY291bnRpZFxuICAgICAgICB2YXJjaGFyIGFjY291bnRuYW1lXG4gICAgfVxuICAgIGFjY291bnQge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICB9XG4gICAgZXZlbnRsb2cgfHwtLW97IGFjY291bnQgOiBPbmUtdG8tT25lXG4gICAgZXZlbnRsb2cgfHwtLW97IGFjY291bnQgOiBPbmUtdG8tT25lXG4gICAgZXZlbnRsb2cgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXZlbnRsb2cge1xuICAgICAgICB2YXJjaGFyIGNoYXJuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIGFjY291bnRpZFxuICAgICAgICB2YXJjaGFyIGFjY291bnRuYW1lXG4gICAgfVxuICAgIGFjY291bnQge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICB9XG4gICAgZXZlbnRsb2cgfHwtLW97IGFjY291bnQgOiBPbmUtdG8tT25lXG4gICAgZXZlbnRsb2cgfHwtLW97IGFjY291bnQgOiBPbmUtdG8tT25lXG4gICAgZXZlbnRsb2cgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | accountid | [account](../../schema/account/account.md) | id |
| One-to-One | accountname | [account](../../schema/account/account.md) | name |
| One-to-One | charname | [character_data](../../schema/characters/character_data.md) | name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Event Identifier |
| accountname | varchar | [Account Name](../../schema/account/account.md) |
| accountid | int | [Account Identifier](../../schema/account/account.md) |
| status | int | [Status](../../../../server/player/status-levels) |
| charname | varchar | [Character Name](../../schema/characters/character_data.md) |
| target | varchar | Target |
| time | timestamp | TIme Timestamp |
| descriptiontype | varchar | Description Type |
| description | text | Description |
| event_nid | int | Event Identifier |

