# eventlog

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXZlbnRsb2cge1xuICAgICAgICBpbnR1bnNpZ25lZCBhY2NvdW50aWRcbiAgICAgICAgdmFyY2hhciBhY2NvdW50bmFtZVxuICAgICAgICB2YXJjaGFyIGNoYXJuYW1lXG4gICAgfVxuICAgIGFjY291bnQge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBldmVudGxvZyB8fC0tb3sgYWNjb3VudCA6IE9uZS10by1PbmVcbiAgICBldmVudGxvZyB8fC0tb3sgYWNjb3VudCA6IE9uZS10by1PbmVcbiAgICBldmVudGxvZyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXZlbnRsb2cge1xuICAgICAgICBpbnR1bnNpZ25lZCBhY2NvdW50aWRcbiAgICAgICAgdmFyY2hhciBhY2NvdW50bmFtZVxuICAgICAgICB2YXJjaGFyIGNoYXJuYW1lXG4gICAgfVxuICAgIGFjY291bnQge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBldmVudGxvZyB8fC0tb3sgYWNjb3VudCA6IE9uZS10by1PbmVcbiAgICBldmVudGxvZyB8fC0tb3sgYWNjb3VudCA6IE9uZS10by1PbmVcbiAgICBldmVudGxvZyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXZlbnRsb2cge1xuICAgICAgICBpbnR1bnNpZ25lZCBhY2NvdW50aWRcbiAgICAgICAgdmFyY2hhciBhY2NvdW50bmFtZVxuICAgICAgICB2YXJjaGFyIGNoYXJuYW1lXG4gICAgfVxuICAgIGFjY291bnQge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBldmVudGxvZyB8fC0tb3sgYWNjb3VudCA6IE9uZS10by1PbmVcbiAgICBldmVudGxvZyB8fC0tb3sgYWNjb3VudCA6IE9uZS10by1PbmVcbiAgICBldmVudGxvZyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


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

