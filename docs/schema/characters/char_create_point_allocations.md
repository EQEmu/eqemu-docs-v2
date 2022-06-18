# char_create_point_allocations

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcl9jcmVhdGVfcG9pbnRfYWxsb2NhdGlvbnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBjaGFyX2NyZWF0ZV9jb21iaW5hdGlvbnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBhbGxvY2F0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIHN0YXJ0X3pvbmVcbiAgICB9XG4gICAgY2hhcl9jcmVhdGVfcG9pbnRfYWxsb2NhdGlvbnMgfHwtLW97IGNoYXJfY3JlYXRlX2NvbWJpbmF0aW9ucyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcl9jcmVhdGVfcG9pbnRfYWxsb2NhdGlvbnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBjaGFyX2NyZWF0ZV9jb21iaW5hdGlvbnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBhbGxvY2F0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIHN0YXJ0X3pvbmVcbiAgICB9XG4gICAgY2hhcl9jcmVhdGVfcG9pbnRfYWxsb2NhdGlvbnMgfHwtLW97IGNoYXJfY3JlYXRlX2NvbWJpbmF0aW9ucyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcl9jcmVhdGVfcG9pbnRfYWxsb2NhdGlvbnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBjaGFyX2NyZWF0ZV9jb21iaW5hdGlvbnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBhbGxvY2F0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIHN0YXJ0X3pvbmVcbiAgICB9XG4gICAgY2hhcl9jcmVhdGVfcG9pbnRfYWxsb2NhdGlvbnMgfHwtLW97IGNoYXJfY3JlYXRlX2NvbWJpbmF0aW9ucyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [char_create_combinations](../../schema/characters/char_create_combinations.md) | allocation_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Allocation Identifier |
| base_str | int | Base Strength |
| base_sta | int | Base Stamina |
| base_dex | int | Base Dexterity |
| base_agi | int | Base Agility |
| base_int | int | Base Intelligence |
| base_wis | int | Base Wisdom |
| base_cha | int | Base Charisma |
| alloc_str | int | Allocated Strength |
| alloc_sta | int | Allocated Stamina |
| alloc_dex | int | Allocated Dexterity |
| alloc_agi | int | Allocated Agility |
| alloc_int | int | Allocated Intelligence |
| alloc_wis | int | Allocated Wisdom |
| alloc_cha | int | Allocated Charisma |

