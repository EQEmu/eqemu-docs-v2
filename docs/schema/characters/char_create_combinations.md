# char_create_combinations

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcl9jcmVhdGVfY29tYmluYXRpb25zIHtcbiAgICAgICAgaW50dW5zaWduZWQgc3RhcnRfem9uZVxuICAgICAgICBpbnR1bnNpZ25lZCBhbGxvY2F0aW9uX2lkXG4gICAgfVxuICAgIGNoYXJfY3JlYXRlX3BvaW50X2FsbG9jYXRpb25zIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgdmFyY2hhciB6b25lX3Nob3J0X25hbWVcbiAgICAgICAgdmFyY2hhciB6b25laWRudW1lclxuICAgIH1cbiAgICBjaGFyX2NyZWF0ZV9jb21iaW5hdGlvbnMgfHwtLW97IGNoYXJfY3JlYXRlX3BvaW50X2FsbG9jYXRpb25zIDogT25lLXRvLU9uZVxuICAgIGNoYXJfY3JlYXRlX2NvbWJpbmF0aW9ucyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcl9jcmVhdGVfY29tYmluYXRpb25zIHtcbiAgICAgICAgaW50dW5zaWduZWQgc3RhcnRfem9uZVxuICAgICAgICBpbnR1bnNpZ25lZCBhbGxvY2F0aW9uX2lkXG4gICAgfVxuICAgIGNoYXJfY3JlYXRlX3BvaW50X2FsbG9jYXRpb25zIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgdmFyY2hhciB6b25lX3Nob3J0X25hbWVcbiAgICAgICAgdmFyY2hhciB6b25laWRudW1lclxuICAgIH1cbiAgICBjaGFyX2NyZWF0ZV9jb21iaW5hdGlvbnMgfHwtLW97IGNoYXJfY3JlYXRlX3BvaW50X2FsbG9jYXRpb25zIDogT25lLXRvLU9uZVxuICAgIGNoYXJfY3JlYXRlX2NvbWJpbmF0aW9ucyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcl9jcmVhdGVfY29tYmluYXRpb25zIHtcbiAgICAgICAgaW50dW5zaWduZWQgc3RhcnRfem9uZVxuICAgICAgICBpbnR1bnNpZ25lZCBhbGxvY2F0aW9uX2lkXG4gICAgfVxuICAgIGNoYXJfY3JlYXRlX3BvaW50X2FsbG9jYXRpb25zIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc1xuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NfZGlzYWJsZWRcbiAgICAgICAgdmFyY2hhciB6b25lX3Nob3J0X25hbWVcbiAgICAgICAgdmFyY2hhciB6b25laWRudW1lclxuICAgIH1cbiAgICBjaGFyX2NyZWF0ZV9jb21iaW5hdGlvbnMgfHwtLW97IGNoYXJfY3JlYXRlX3BvaW50X2FsbG9jYXRpb25zIDogT25lLXRvLU9uZVxuICAgIGNoYXJfY3JlYXRlX2NvbWJpbmF0aW9ucyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | allocation_id | [char_create_point_allocations](../../schema/characters/char_create_point_allocations.md) | id |
| One-to-One | start_zone | [zone](../../schema/zone/zone.md) | zoneidnumber |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| allocation_id | int | [Allocation Identifier](char_create_point_allocations.md) |
| race | int | [Race](../../../../server/npc/race-list) |
| class | int | [Class](../../../../server/player/class-list) |
| deity | int | [Deity](../../../../server/player/deity-list) |
| start_zone | int | [Start Zone Identifier](../../../../server/zones/zone-list) |
| expansions_req | int | [Expansions Required](../../../../server/operation/expansion-bitmasks) |

