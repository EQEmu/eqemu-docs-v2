# aa_rank_effects

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWFfcmFua19lZmZlY3RzIHtcbiAgICAgICAgaW50dW5zaWduZWQgcmFua19pZFxuICAgIH1cbiAgICBhYV9yYW5rcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIGFhX3JhbmtfZWZmZWN0cyB8fC0tb3sgYWFfcmFua3MgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWFfcmFua19lZmZlY3RzIHtcbiAgICAgICAgaW50dW5zaWduZWQgcmFua19pZFxuICAgIH1cbiAgICBhYV9yYW5rcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIGFhX3JhbmtfZWZmZWN0cyB8fC0tb3sgYWFfcmFua3MgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWFfcmFua19lZmZlY3RzIHtcbiAgICAgICAgaW50dW5zaWduZWQgcmFua19pZFxuICAgIH1cbiAgICBhYV9yYW5rcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIGFhX3JhbmtfZWZmZWN0cyB8fC0tb3sgYWFfcmFua3MgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | rank_id | [aa_ranks](../../schema/aas/aa_ranks.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| rank_id | int | [Rank Identifier](aa_ranks.md) |
| slot | int | AA Slot |
| effect_id | int | [Spell Effect Identifier](../../../../server/spells/spell-effect-ids) |
| base1 | int | First Base Value |
| base2 | int | Second Base Value |

