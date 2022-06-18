# item_tick

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgaXRlbV90aWNrIHtcbiAgICAgICAgaW50IGl0X2l0ZW1pZFxuICAgICAgICB2YXJjaGFyIGl0X3FnbG9iYWxcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGljb25cbiAgICAgICAgIG5hbWVcbiAgICAgICAgc21hbGxpbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIGludCBwcm9jZWZmZWN0XG4gICAgICAgIGludCBzY3JvbGxlZmZlY3RcbiAgICAgICAgaW50IHdvcm5lZmZlY3RcbiAgICAgICAgaW50IGJvb2tcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICB9XG4gICAgcXVlc3RfZ2xvYmFscyB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgem9uZWlkXG4gICAgICAgIGludCBjaGFyaWRcbiAgICAgICAgaW50IG5wY2lkXG4gICAgfVxuICAgIGl0ZW1fdGljayB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG4gICAgaXRlbV90aWNrIHx8LS1veyBxdWVzdF9nbG9iYWxzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgaXRlbV90aWNrIHtcbiAgICAgICAgaW50IGl0X2l0ZW1pZFxuICAgICAgICB2YXJjaGFyIGl0X3FnbG9iYWxcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGljb25cbiAgICAgICAgIG5hbWVcbiAgICAgICAgc21hbGxpbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIGludCBwcm9jZWZmZWN0XG4gICAgICAgIGludCBzY3JvbGxlZmZlY3RcbiAgICAgICAgaW50IHdvcm5lZmZlY3RcbiAgICAgICAgaW50IGJvb2tcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICB9XG4gICAgcXVlc3RfZ2xvYmFscyB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgem9uZWlkXG4gICAgICAgIGludCBjaGFyaWRcbiAgICAgICAgaW50IG5wY2lkXG4gICAgfVxuICAgIGl0ZW1fdGljayB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG4gICAgaXRlbV90aWNrIHx8LS1veyBxdWVzdF9nbG9iYWxzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgaXRlbV90aWNrIHtcbiAgICAgICAgaW50IGl0X2l0ZW1pZFxuICAgICAgICB2YXJjaGFyIGl0X3FnbG9iYWxcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGljb25cbiAgICAgICAgIG5hbWVcbiAgICAgICAgc21hbGxpbnQgYmFyZGVmZmVjdFxuICAgICAgICBpbnQgY2xpY2tlZmZlY3RcbiAgICAgICAgaW50IGZvY3VzZWZmZWN0XG4gICAgICAgIGludCBwcm9jZWZmZWN0XG4gICAgICAgIGludCBzY3JvbGxlZmZlY3RcbiAgICAgICAgaW50IHdvcm5lZmZlY3RcbiAgICAgICAgaW50IGJvb2tcbiAgICAgICAgaW50IHJlY2FzdHR5cGVcbiAgICB9XG4gICAgcXVlc3RfZ2xvYmFscyB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgICAgICBpbnQgem9uZWlkXG4gICAgICAgIGludCBjaGFyaWRcbiAgICAgICAgaW50IG5wY2lkXG4gICAgfVxuICAgIGl0ZW1fdGljayB8fC0tb3sgaXRlbXMgOiBPbmUtdG8tT25lXG4gICAgaXRlbV90aWNrIHx8LS1veyBxdWVzdF9nbG9iYWxzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | it_itemid | [items](../../schema/items/items.md) | id |
| Has-Many | it_qglobal | [quest_globals](../../schema/data-storage/quest_globals.md) | name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| it_itemid | int | [Item Identifier](items.md) |
| it_chance | int | Chance: 0 = Never, 100 = Always |
| it_level | int | Level |
| it_id | int | [Spell Identifier](../../schema/spells/spells_new.md) |
| it_qglobal | varchar | [Quest Global Identifier](../../schema/data-storage/quest_globals.md) (Deprecated) |
| it_bagslot | tinyint | [Bag Slot](../../../../server/inventory/inventory-slots) |

