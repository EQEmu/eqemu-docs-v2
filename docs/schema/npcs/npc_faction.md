# npc_faction

!!! info
	This page was last generated 2022.02.23

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX2ZhY3Rpb24ge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbnBjX2ZhY3Rpb25fZW50cmllcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG5wY19mYWN0aW9uX2lkXG4gICAgfVxuICAgIG5wY19mYWN0aW9uIHx8LS1veyBucGNfZmFjdGlvbl9lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX2ZhY3Rpb24ge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbnBjX2ZhY3Rpb25fZW50cmllcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG5wY19mYWN0aW9uX2lkXG4gICAgfVxuICAgIG5wY19mYWN0aW9uIHx8LS1veyBucGNfZmFjdGlvbl9lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX2ZhY3Rpb24ge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbnBjX2ZhY3Rpb25fZW50cmllcyB7XG4gICAgICAgIGludHVuc2lnbmVkIG5wY19mYWN0aW9uX2lkXG4gICAgfVxuICAgIG5wY19mYWN0aW9uIHx8LS1veyBucGNfZmFjdGlvbl9lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [npc_faction_entries](../../schema/npcs/npc_faction_entries.md) | npc_faction_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique NPC Faction Identifier |
| name | tinytext | Name |
| primaryfaction | int | [Primary Faction Identifier](../../schema/factions/faction_list.md) |
| ignore_primary_assist | tinyint | Ignore Primary Assist: 0 = False, &gt;0 = True |

