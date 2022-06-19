# expedition_lockouts

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXhwZWRpdGlvbl9sb2Nrb3V0cyB7XG4gICAgICAgIGludHVuc2lnbmVkIGV4cGVkaXRpb25faWRcbiAgICB9XG4gICAgZXhwZWRpdGlvbnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBkeW5hbWljX3pvbmVfaWRcbiAgICB9XG4gICAgZXhwZWRpdGlvbl9sb2Nrb3V0cyB8fC0tb3sgZXhwZWRpdGlvbnMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXhwZWRpdGlvbl9sb2Nrb3V0cyB7XG4gICAgICAgIGludHVuc2lnbmVkIGV4cGVkaXRpb25faWRcbiAgICB9XG4gICAgZXhwZWRpdGlvbnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBkeW5hbWljX3pvbmVfaWRcbiAgICB9XG4gICAgZXhwZWRpdGlvbl9sb2Nrb3V0cyB8fC0tb3sgZXhwZWRpdGlvbnMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXhwZWRpdGlvbl9sb2Nrb3V0cyB7XG4gICAgICAgIGludHVuc2lnbmVkIGV4cGVkaXRpb25faWRcbiAgICB9XG4gICAgZXhwZWRpdGlvbnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBkeW5hbWljX3pvbmVfaWRcbiAgICB9XG4gICAgZXhwZWRpdGlvbl9sb2Nrb3V0cyB8fC0tb3sgZXhwZWRpdGlvbnMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | expedition_id | [expeditions](../../schema/expeditions/expeditions.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Expedition Lockout Identifier |
| expedition_id | int | [Expedition Identifier](expeditions.md) |
| event_name | varchar | Event Name |
| expire_time | datetime | Expire Time |
| duration | int | Duration in Seconds |
| from_expedition_uuid | varchar | From Expedition UUID |

