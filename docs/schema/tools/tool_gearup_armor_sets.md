# tool_gearup_armor_sets

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBpY29uXG4gICAgICAgIHNtYWxsaW50IGJhcmRlZmZlY3RcbiAgICAgICAgaW50IGNsaWNrZWZmZWN0XG4gICAgICAgIGludCBmb2N1c2VmZmVjdFxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgYm9va1xuICAgICAgICBpbnQgcmVjYXN0dHlwZVxuICAgIH1cbiAgICB0b29sX2dlYXJ1cF9hcm1vcl9zZXRzIHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBpY29uXG4gICAgICAgIHNtYWxsaW50IGJhcmRlZmZlY3RcbiAgICAgICAgaW50IGNsaWNrZWZmZWN0XG4gICAgICAgIGludCBmb2N1c2VmZmVjdFxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgYm9va1xuICAgICAgICBpbnQgcmVjYXN0dHlwZVxuICAgIH1cbiAgICB0b29sX2dlYXJ1cF9hcm1vcl9zZXRzIHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdG9vbF9nZWFydXBfYXJtb3Jfc2V0cyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgICAgIGludCBpY29uXG4gICAgICAgIHNtYWxsaW50IGJhcmRlZmZlY3RcbiAgICAgICAgaW50IGNsaWNrZWZmZWN0XG4gICAgICAgIGludCBmb2N1c2VmZmVjdFxuICAgICAgICBpbnQgcHJvY2VmZmVjdFxuICAgICAgICBpbnQgc2Nyb2xsZWZmZWN0XG4gICAgICAgIGludCB3b3JuZWZmZWN0XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgYm9va1xuICAgICAgICBpbnQgcmVjYXN0dHlwZVxuICAgIH1cbiAgICB0b29sX2dlYXJ1cF9hcm1vcl9zZXRzIHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| class | tinyint | [Class](../../../../categories/player/class-list) |
| level | smallint | Level |
| slot | tinyint | [Slot Identifier](../../../../server/inventory/inventory-slots) |
| item_id | int | [Item Identifier](../../schema/items/items.md) |
| score | mediumint | Score |
| expansion | tinyint | [Expansion](../../../../server/operation/expansion-list) |

