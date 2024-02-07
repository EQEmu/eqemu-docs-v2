# spawn_events

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25fZXZlbnRzIHtcbiAgICAgICAgdmFyY2hhciB6b25lXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGNvbmRfaWRcbiAgICB9XG4gICAgc3Bhd25fY29uZGl0aW9ucyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBzcGF3bl9ldmVudHMgfHwtLW97IHNwYXduX2NvbmRpdGlvbnMgOiBcIk9uZS10by1PbmVcIlxuICAgIHNwYXduX2V2ZW50cyB8fC0tb3sgem9uZSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25fZXZlbnRzIHtcbiAgICAgICAgdmFyY2hhciB6b25lXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGNvbmRfaWRcbiAgICB9XG4gICAgc3Bhd25fY29uZGl0aW9ucyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBzcGF3bl9ldmVudHMgfHwtLW97IHNwYXduX2NvbmRpdGlvbnMgOiBcIk9uZS10by1PbmVcIlxuICAgIHNwYXduX2V2ZW50cyB8fC0tb3sgem9uZSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25fZXZlbnRzIHtcbiAgICAgICAgdmFyY2hhciB6b25lXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGNvbmRfaWRcbiAgICB9XG4gICAgc3Bhd25fY29uZGl0aW9ucyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBzcGF3bl9ldmVudHMgfHwtLW97IHNwYXduX2NvbmRpdGlvbnMgOiBcIk9uZS10by1PbmVcIlxuICAgIHNwYXduX2V2ZW50cyB8fC0tb3sgem9uZSA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | cond_id | [spawn_conditions](../../schema/spawns/spawn_conditions.md) | id |
| One-to-One | zone | [zone](../../schema/zone/zone.md) | short_name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Spawn Event Entry Identifier |
| zone | varchar | [Zone Short Name](../../../../server/zones/zone-list) |
| cond_id | mediumint | [Spawn Condition Identifier](spawn_conditions.md) |
| name | varchar | Name |
| period | int | Period |
| next_minute | tinyint | Next Minute |
| next_hour | tinyint | Next Hour |
| next_day | tinyint | Next Day |
| next_month | tinyint | Next Month |
| next_year | int | Next Year |
| enabled | tinyint | Enabled: 0 = False, 1 = True |
| action | tinyint | [Action Type](../../../../server/npc/spawns/action-types) |
| argument | mediumint | Argument: (Based on Action) 0 = Argument Value |
| strict | tinyint | Strict Date Criteria: 0 = False, 1 = True |

