# bot_owner_options

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X293bmVyX29wdGlvbnMge1xuICAgICAgICB2YXJjaGFyIG93bmVyX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBib3Rfb3duZXJfb3B0aW9ucyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X293bmVyX29wdGlvbnMge1xuICAgICAgICB2YXJjaGFyIG93bmVyX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBib3Rfb3duZXJfb3B0aW9ucyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X293bmVyX29wdGlvbnMge1xuICAgICAgICB2YXJjaGFyIG93bmVyX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBib3Rfb3duZXJfb3B0aW9ucyB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | owner_id | [character_data](../../schema/characters/character_data.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| owner_id | int | [Owner Character Identifier](../../../schema/characters/character_data.md) |
| option_type | smallint | Option Type |
| option_value | smallint | Option Value |

