# global_loot

!!! info
	This page was last generated 2022.02.23

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ2xvYmFsX2xvb3Qge1xuICAgICAgICBpbnQgbG9vdHRhYmxlX2lkXG4gICAgfVxuICAgIGxvb3R0YWJsZSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIGdsb2JhbF9sb290IHx8LS1veyBsb290dGFibGUgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ2xvYmFsX2xvb3Qge1xuICAgICAgICBpbnQgbG9vdHRhYmxlX2lkXG4gICAgfVxuICAgIGxvb3R0YWJsZSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIGdsb2JhbF9sb290IHx8LS1veyBsb290dGFibGUgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ2xvYmFsX2xvb3Qge1xuICAgICAgICBpbnQgbG9vdHRhYmxlX2lkXG4gICAgfVxuICAgIGxvb3R0YWJsZSB7XG4gICAgICAgIGludHVuc2lnbmVkIGlkXG4gICAgfVxuICAgIGdsb2JhbF9sb290IHx8LS1veyBsb290dGFibGUgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | loottable_id | [loottable](../../schema/loot/loottable.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Global Loot Identifier |
| description | varchar | Description |
| loottable_id | int | [Loottable Identifier](loottable.md) |
| enabled | tinyint | Enabled: 0 = False, 1 = True |
| min_level | int | Minimum Level |
| max_level | int | Maximum Level |
| rare | tinyint | Rare: 0 = False, 1 = True |
| raid | tinyint | Raid: 0 = False, 1 = True |
| race | mediumtext | [Race](../../../../server/npc/race-list), multiple races supported if  |
| class | mediumtext | [Class](../../../../server/player/class-list), multiple classes supported if  |
| bodytype | mediumtext | [Body Type](../../../../server/npc/body-types), multiple body types supported if  |
| zone | mediumtext | [Zone Short Name](../../../../server/zones/zone-list),, multiple zones supported if  |
| hot_zone | tinyint |  |
| min_expansion | tinyint |  |
| max_expansion | tinyint |  |
| content_flags | varchar |  |
| content_flags_disabled | varchar |  |

