# spawn_events

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25fZXZlbnRzIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgY29uZF9pZFxuICAgICAgICB2YXJjaGFyIHpvbmVcbiAgICB9XG4gICAgc3Bhd25fY29uZGl0aW9ucyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgIH1cbiAgICBzcGF3bl9ldmVudHMgfHwtLW97IHNwYXduX2NvbmRpdGlvbnMgOiBPbmUtdG8tT25lXG4gICAgc3Bhd25fZXZlbnRzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25fZXZlbnRzIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgY29uZF9pZFxuICAgICAgICB2YXJjaGFyIHpvbmVcbiAgICB9XG4gICAgc3Bhd25fY29uZGl0aW9ucyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgIH1cbiAgICBzcGF3bl9ldmVudHMgfHwtLW97IHNwYXduX2NvbmRpdGlvbnMgOiBPbmUtdG8tT25lXG4gICAgc3Bhd25fZXZlbnRzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25fZXZlbnRzIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgY29uZF9pZFxuICAgICAgICB2YXJjaGFyIHpvbmVcbiAgICB9XG4gICAgc3Bhd25fY29uZGl0aW9ucyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgIH1cbiAgICBzcGF3bl9ldmVudHMgfHwtLW97IHNwYXduX2NvbmRpdGlvbnMgOiBPbmUtdG8tT25lXG4gICAgc3Bhd25fZXZlbnRzIHx8LS1veyB6b25lIDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


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

