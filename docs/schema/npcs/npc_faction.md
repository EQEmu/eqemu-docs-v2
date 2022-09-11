# npc_faction

!!! info
	This page was last generated 2022.09.10

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX2ZhY3Rpb24ge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHByaW1hcnlmYWN0aW9uXG4gICAgfVxuICAgIG5wY19mYWN0aW9uX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBmYWN0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG5wY19mYWN0aW9uX2lkXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBucGNfZmFjdGlvbiB8fC0tb3sgbnBjX2ZhY3Rpb25fZW50cmllcyA6IEhhcy1NYW55XG4gICAgbnBjX2ZhY3Rpb24gfHwtLW97IGZhY3Rpb25fbGlzdCA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX2ZhY3Rpb24ge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHByaW1hcnlmYWN0aW9uXG4gICAgfVxuICAgIG5wY19mYWN0aW9uX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBmYWN0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG5wY19mYWN0aW9uX2lkXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBucGNfZmFjdGlvbiB8fC0tb3sgbnBjX2ZhY3Rpb25fZW50cmllcyA6IEhhcy1NYW55XG4gICAgbnBjX2ZhY3Rpb24gfHwtLW97IGZhY3Rpb25fbGlzdCA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX2ZhY3Rpb24ge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IHByaW1hcnlmYWN0aW9uXG4gICAgfVxuICAgIG5wY19mYWN0aW9uX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBmYWN0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIG5wY19mYWN0aW9uX2lkXG4gICAgfVxuICAgIGZhY3Rpb25fbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgIH1cbiAgICBucGNfZmFjdGlvbiB8fC0tb3sgbnBjX2ZhY3Rpb25fZW50cmllcyA6IEhhcy1NYW55XG4gICAgbnBjX2ZhY3Rpb24gfHwtLW97IGZhY3Rpb25fbGlzdCA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [npc_faction_entries](../../schema/npcs/npc_faction_entries.md) | npc_faction_id |
| One-to-One | primaryfaction | [faction_list](../../schema/factions/faction_list.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique NPC Faction Identifier |
| name | tinytext | Name |
| primaryfaction | int | [Primary Faction Identifier](../../schema/factions/faction_list.md) |
| ignore_primary_assist | tinyint | Ignore Primary Assist: 0 = False, &gt;0 = True |

