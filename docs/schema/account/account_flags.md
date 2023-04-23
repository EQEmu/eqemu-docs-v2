# account_flags

!!! info
	This page was last generated 2023.04.23

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWNjb3VudF9mbGFncyB7XG4gICAgICAgIGludHVuc2lnbmVkIHBfYWNjaWRcbiAgICB9XG4gICAgYWNjb3VudCB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICB9XG4gICAgYWNjb3VudF9mbGFncyB8fC0tb3sgYWNjb3VudCA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWNjb3VudF9mbGFncyB7XG4gICAgICAgIGludHVuc2lnbmVkIHBfYWNjaWRcbiAgICB9XG4gICAgYWNjb3VudCB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICB9XG4gICAgYWNjb3VudF9mbGFncyB8fC0tb3sgYWNjb3VudCA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWNjb3VudF9mbGFncyB7XG4gICAgICAgIGludHVuc2lnbmVkIHBfYWNjaWRcbiAgICB9XG4gICAgYWNjb3VudCB7XG4gICAgICAgIGludCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICB9XG4gICAgYWNjb3VudF9mbGFncyB8fC0tb3sgYWNjb3VudCA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


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

