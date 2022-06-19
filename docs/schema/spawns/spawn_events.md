# spawn_events

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25fZXZlbnRzIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgY29uZF9pZFxuICAgICAgICB2YXJjaGFyIHpvbmVcbiAgICB9XG4gICAgc3Bhd25fY29uZGl0aW9ucyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgICAgICBtZWRpdW1pbnQgdmFsdWVcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCB2ZXJzaW9uXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICB9XG4gICAgc3Bhd25fZXZlbnRzIHx8LS1veyBzcGF3bl9jb25kaXRpb25zIDogT25lLXRvLU9uZVxuICAgIHNwYXduX2V2ZW50cyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25fZXZlbnRzIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgY29uZF9pZFxuICAgICAgICB2YXJjaGFyIHpvbmVcbiAgICB9XG4gICAgc3Bhd25fY29uZGl0aW9ucyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgICAgICBtZWRpdW1pbnQgdmFsdWVcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCB2ZXJzaW9uXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICB9XG4gICAgc3Bhd25fZXZlbnRzIHx8LS1veyBzcGF3bl9jb25kaXRpb25zIDogT25lLXRvLU9uZVxuICAgIHNwYXduX2V2ZW50cyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgc3Bhd25fZXZlbnRzIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgY29uZF9pZFxuICAgICAgICB2YXJjaGFyIHpvbmVcbiAgICB9XG4gICAgc3Bhd25fY29uZGl0aW9ucyB7XG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGlkXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgICAgICBtZWRpdW1pbnQgdmFsdWVcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgICAgICBpbnQgem9uZWlkbnVtYmVyXG4gICAgICAgIHRpbnlpbnR1bnNpZ25lZCB2ZXJzaW9uXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICB9XG4gICAgc3Bhd25fZXZlbnRzIHx8LS1veyBzcGF3bl9jb25kaXRpb25zIDogT25lLXRvLU9uZVxuICAgIHNwYXduX2V2ZW50cyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


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

