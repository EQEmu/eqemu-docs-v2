# adventure_stats

!!! info
	This page was last generated 2022.09.10

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWR2ZW50dXJlX3N0YXRzIHtcbiAgICAgICAgaW50dW5zaWduZWQgcGxheWVyX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICB9XG4gICAgYWR2ZW50dXJlX3N0YXRzIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWR2ZW50dXJlX3N0YXRzIHtcbiAgICAgICAgaW50dW5zaWduZWQgcGxheWVyX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICB9XG4gICAgYWR2ZW50dXJlX3N0YXRzIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWR2ZW50dXJlX3N0YXRzIHtcbiAgICAgICAgaW50dW5zaWduZWQgcGxheWVyX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICB9XG4gICAgYWR2ZW50dXJlX3N0YXRzIHx8LS1veyBjaGFyYWN0ZXJfZGF0YSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | player_id | [character_data](../../schema/characters/character_data.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| player_id | int | [Character Identifier](../../schema/characters/character_data.md) |
| guk_wins | mediumint | Deepest Guk Wins |
| mir_wins | mediumint | Miragul's Menagerie Wins |
| mmc_wins | mediumint | Mistmoore Catacombs Wins |
| ruj_wins | mediumint | Rujarkian Hills Wins |
| tak_wins | mediumint | Takish-Hiz Wins |
| guk_losses | mediumint | Deepest Guk Losses |
| mir_losses | mediumint | Miragul's Menagerie Losses |
| mmc_losses | mediumint | Mistmoore Catacombs Losses |
| ruj_losses | mediumint | Rujarkian Hills Losses |
| tak_losses | mediumint | Takish-Hiz Losses |

