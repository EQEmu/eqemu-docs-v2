# blocked_spells

!!! info
	This page was last generated 2022.09.10

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYmxvY2tlZF9zcGVsbHMge1xuICAgICAgICBtZWRpdW1pbnR1bnNpZ25lZCBzcGVsbGlkXG4gICAgICAgIGludCB6b25laWRcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bTJcbiAgICAgICAgaW50IHR5cGVkZXNjbnVtXG4gICAgICAgIHZhcmNoYXIgdGVsZXBvcnRfem9uZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBibG9ja2VkX3NwZWxscyB8fC0tb3sgc3BlbGxzX25ldyA6IE9uZS10by1PbmVcbiAgICBibG9ja2VkX3NwZWxscyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYmxvY2tlZF9zcGVsbHMge1xuICAgICAgICBtZWRpdW1pbnR1bnNpZ25lZCBzcGVsbGlkXG4gICAgICAgIGludCB6b25laWRcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bTJcbiAgICAgICAgaW50IHR5cGVkZXNjbnVtXG4gICAgICAgIHZhcmNoYXIgdGVsZXBvcnRfem9uZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBibG9ja2VkX3NwZWxscyB8fC0tb3sgc3BlbGxzX25ldyA6IE9uZS10by1PbmVcbiAgICBibG9ja2VkX3NwZWxscyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYmxvY2tlZF9zcGVsbHMge1xuICAgICAgICBtZWRpdW1pbnR1bnNpZ25lZCBzcGVsbGlkXG4gICAgICAgIGludCB6b25laWRcbiAgICB9XG4gICAgc3BlbGxzX25ldyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bVxuICAgICAgICBpbnQgZWZmZWN0ZGVzY251bTJcbiAgICAgICAgaW50IHR5cGVkZXNjbnVtXG4gICAgICAgIHZhcmNoYXIgdGVsZXBvcnRfem9uZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBibG9ja2VkX3NwZWxscyB8fC0tb3sgc3BlbGxzX25ldyA6IE9uZS10by1PbmVcbiAgICBibG9ja2VkX3NwZWxscyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | spellid | [spells_new](../../schema/spells/spells_new.md) | id |
| One-to-One | zoneid | [zone](../../schema/zone/zone.md) | zoneidnumber |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Blocked Spells Identifier |
| spellid | mediumint | [Spell Identifier](spells_new.md) |
| type | tinyint | [Blocked Spell Type](../../../../server/spells/blocked-spell-types) |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| x | float | X Coordinate |
| y | float | Y Coordinate |
| z | float | Z Coordinate |
| x_diff | float | X Radius |
| y_diff | float | Y Radius |
| z_diff | float | Z Radius |
| message | varchar | Message when blocked |
| description | varchar | Blocked spells description |

