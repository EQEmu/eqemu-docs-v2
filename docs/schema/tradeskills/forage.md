# forage

!!! info
	This page was last generated 2022.06.17

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZm9yYWdlIHtcbiAgICAgICAgaW50IEl0ZW1pZFxuICAgICAgICBpbnQgem9uZWlkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICAgem9uZWlkdW51bWJlclxuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgIH1cbiAgICBmb3JhZ2UgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuICAgIGZvcmFnZSB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZm9yYWdlIHtcbiAgICAgICAgaW50IEl0ZW1pZFxuICAgICAgICBpbnQgem9uZWlkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICAgem9uZWlkdW51bWJlclxuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgIH1cbiAgICBmb3JhZ2UgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuICAgIGZvcmFnZSB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZm9yYWdlIHtcbiAgICAgICAgaW50IEl0ZW1pZFxuICAgICAgICBpbnQgem9uZWlkXG4gICAgfVxuICAgIGl0ZW1zIHtcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIHpvbmUge1xuICAgICAgICAgem9uZWlkdW51bWJlclxuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgIH1cbiAgICBmb3JhZ2UgfHwtLW97IGl0ZW1zIDogT25lLXRvLU9uZVxuICAgIGZvcmFnZSB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | Itemid | [items](../../schema/items/items.md) | id |
| One-to-One | zoneid | [zone](../../schema/zone/zone.md) | zoneidnumber |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Forage Identifier |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| Itemid | int | [Item Identifier](items.md) |
| level | smallint | Level |
| chance | smallint | Chance: 0 = Never, 100 = Always |
| min_expansion | tinyint | [Minimum Expansion](../../../../server/operation/expansion-list) |
| max_expansion | tinyint | [Maximum Expansion](../../../../server/operation/expansion-list) |
| content_flags | varchar | Content Flags Required to be Enabled |
| content_flags_disabled | varchar | Content Flags Required to be Disabled |

