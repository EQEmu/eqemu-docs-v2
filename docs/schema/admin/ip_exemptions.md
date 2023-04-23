# ip_exemptions

!!! info
	This page was last generated 2023.04.23

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgaXBfZXhlbXB0aW9ucyB7XG4gICAgICAgIHZhcmNoYXIgZXhlbXB0aW9uX2lwXG4gICAgfVxuICAgIGFjY291bnRfaXAge1xuICAgICAgICB2YXJjaGFyIGlwXG4gICAgICAgIGludCBhY2NpZFxuICAgIH1cbiAgICBpcF9leGVtcHRpb25zIHx8LS1veyBhY2NvdW50X2lwIDogXCJIYXMtTWFueVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgaXBfZXhlbXB0aW9ucyB7XG4gICAgICAgIHZhcmNoYXIgZXhlbXB0aW9uX2lwXG4gICAgfVxuICAgIGFjY291bnRfaXAge1xuICAgICAgICB2YXJjaGFyIGlwXG4gICAgICAgIGludCBhY2NpZFxuICAgIH1cbiAgICBpcF9leGVtcHRpb25zIHx8LS1veyBhY2NvdW50X2lwIDogXCJIYXMtTWFueVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgaXBfZXhlbXB0aW9ucyB7XG4gICAgICAgIHZhcmNoYXIgZXhlbXB0aW9uX2lwXG4gICAgfVxuICAgIGFjY291bnRfaXAge1xuICAgICAgICB2YXJjaGFyIGlwXG4gICAgICAgIGludCBhY2NpZFxuICAgIH1cbiAgICBpcF9leGVtcHRpb25zIHx8LS1veyBhY2NvdW50X2lwIDogXCJIYXMtTWFueVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


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

