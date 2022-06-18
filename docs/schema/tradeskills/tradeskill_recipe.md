# tradeskill_recipe

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdHJhZGVza2lsbF9yZWNpcGUge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBjb250ZW50X2ZsYWdzIHtcbiAgICAgICAgdmFyY2hhciBmbGFnX25hbWVcbiAgICB9XG4gICAgY29udGVudF9mbGFncyB7XG4gICAgICAgIHZhcmNoYXIgZmxhZ19uYW1lXG4gICAgfVxuICAgIHRyYWRlc2tpbGxfcmVjaXBlX2VudHJpZXMge1xuICAgICAgICBpbnQgaXRlbV9pZFxuICAgICAgICBpbnQgcmVjaXBlX2lkXG4gICAgfVxuICAgIHRyYWRlc2tpbGxfcmVjaXBlIHx8LS1veyBjb250ZW50X2ZsYWdzIDogT25lLXRvLU9uZVxuICAgIHRyYWRlc2tpbGxfcmVjaXBlIHx8LS1veyBjb250ZW50X2ZsYWdzIDogT25lLXRvLU9uZVxuICAgIHRyYWRlc2tpbGxfcmVjaXBlIHx8LS1veyB0cmFkZXNraWxsX3JlY2lwZV9lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdHJhZGVza2lsbF9yZWNpcGUge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBjb250ZW50X2ZsYWdzIHtcbiAgICAgICAgdmFyY2hhciBmbGFnX25hbWVcbiAgICB9XG4gICAgY29udGVudF9mbGFncyB7XG4gICAgICAgIHZhcmNoYXIgZmxhZ19uYW1lXG4gICAgfVxuICAgIHRyYWRlc2tpbGxfcmVjaXBlX2VudHJpZXMge1xuICAgICAgICBpbnQgaXRlbV9pZFxuICAgICAgICBpbnQgcmVjaXBlX2lkXG4gICAgfVxuICAgIHRyYWRlc2tpbGxfcmVjaXBlIHx8LS1veyBjb250ZW50X2ZsYWdzIDogT25lLXRvLU9uZVxuICAgIHRyYWRlc2tpbGxfcmVjaXBlIHx8LS1veyBjb250ZW50X2ZsYWdzIDogT25lLXRvLU9uZVxuICAgIHRyYWRlc2tpbGxfcmVjaXBlIHx8LS1veyB0cmFkZXNraWxsX3JlY2lwZV9lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgdHJhZGVza2lsbF9yZWNpcGUge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBjb250ZW50X2ZsYWdzIHtcbiAgICAgICAgdmFyY2hhciBmbGFnX25hbWVcbiAgICB9XG4gICAgY29udGVudF9mbGFncyB7XG4gICAgICAgIHZhcmNoYXIgZmxhZ19uYW1lXG4gICAgfVxuICAgIHRyYWRlc2tpbGxfcmVjaXBlX2VudHJpZXMge1xuICAgICAgICBpbnQgaXRlbV9pZFxuICAgICAgICBpbnQgcmVjaXBlX2lkXG4gICAgfVxuICAgIHRyYWRlc2tpbGxfcmVjaXBlIHx8LS1veyBjb250ZW50X2ZsYWdzIDogT25lLXRvLU9uZVxuICAgIHRyYWRlc2tpbGxfcmVjaXBlIHx8LS1veyBjb250ZW50X2ZsYWdzIDogT25lLXRvLU9uZVxuICAgIHRyYWRlc2tpbGxfcmVjaXBlIHx8LS1veyB0cmFkZXNraWxsX3JlY2lwZV9lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | content_flags | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| One-to-One | content_flags_disabled | [content_flags](../../schema/flagging/content_flags.md) | flag_name |
| Has-Many | id | [tradeskill_recipe_entries](../../schema/tradeskills/tradeskill_recipe_entries.md) | recipe_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Tradeskill Recipe Identifier |
| name | varchar | Recipe Name |
| tradeskill | smallint | [Tradeskill](../../../../server/player/skills) |
| skillneeded | smallint | Skill Level Needed |
| trivial | smallint | Trivial Skill Level |
| nofail | tinyint | No Fail: 0 = False, 1 = True |
| replace_container | tinyint | Replace Container: 0 = False, 1 = True |
| notes | tinytext | Notes |
| must_learn | tinyint | Must Learn: 0 = False, 1 = True |
| quest | tinyint | Quest Controlled: 0 = False, 1 = True |
| enabled | tinyint | Enabled: 0 = False, 1 = True |
| min_expansion | tinyint | [Minimum Expansion](../../../../server/operation/expansion-list) |
| max_expansion | tinyint | [Maximum Expansion](../../../../server/operation/expansion-list) |
| content_flags | varchar | Content Flags Required to be Enabled |
| content_flags_disabled | varchar | Content Flags Required to be Disabled |

