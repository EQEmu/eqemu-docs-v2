# keyring

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAga2V5cmluZyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgICAgIGludCBjaGFyX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIHNtYWxsaW50IGJhcmRlZmZlY3RcbiAgICAgICAgaW50IGNsaWNrZWZmZWN0XG4gICAgICAgIGludCBmb2N1c2VmZmVjdFxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgICAgIGludCByZWNhc3R0eXBlXG4gICAgICAgIGludCBpY29uXG4gICAgICAgIGludCBib29rXG4gICAgfVxuICAgIGtleXJpbmcgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuICAgIGtleXJpbmcgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAga2V5cmluZyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgICAgIGludCBjaGFyX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIHNtYWxsaW50IGJhcmRlZmZlY3RcbiAgICAgICAgaW50IGNsaWNrZWZmZWN0XG4gICAgICAgIGludCBmb2N1c2VmZmVjdFxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgICAgIGludCByZWNhc3R0eXBlXG4gICAgICAgIGludCBpY29uXG4gICAgICAgIGludCBib29rXG4gICAgfVxuICAgIGtleXJpbmcgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuICAgIGtleXJpbmcgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAga2V5cmluZyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgICAgIGludCBjaGFyX2lkXG4gICAgfVxuICAgIGNoYXJhY3Rlcl9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaW5zdGFuY2VcbiAgICAgICAgdmFyY2hhciBuYW5lXG4gICAgICAgIGludHVuc2lnbmVkIHpvbmVfaWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBuYW1lXG4gICAgICAgIHNtYWxsaW50IGJhcmRlZmZlY3RcbiAgICAgICAgaW50IGNsaWNrZWZmZWN0XG4gICAgICAgIGludCBmb2N1c2VmZmVjdFxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgICAgIGludCByZWNhc3R0eXBlXG4gICAgICAgIGludCBpY29uXG4gICAgICAgIGludCBib29rXG4gICAgfVxuICAgIGtleXJpbmcgfHwtLW97IGNoYXJhY3Rlcl9kYXRhIDogT25lLXRvLU9uZVxuICAgIGtleXJpbmcgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| char_id | int | [Character Identifier](character_data.md) |
| item_id | int | [Item Identifier](../../schema/items/items.md) |

