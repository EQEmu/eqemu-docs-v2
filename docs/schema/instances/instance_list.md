# instance_list

!!! info
	This page was last generated 2021.11.23

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgaW5zdGFuY2VfbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lXG4gICAgfVxuICAgIGluc3RhbmNlX2xpc3RfcGxheWVyIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgICB6b25laWR1bnVtYmVyXG4gICAgfVxuICAgIGluc3RhbmNlX2xpc3QgfHwtLW97IGluc3RhbmNlX2xpc3RfcGxheWVyIDogSGFzLU1hbnlcbiAgICBpbnN0YW5jZV9saXN0IHx8LS1veyB6b25lIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgaW5zdGFuY2VfbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lXG4gICAgfVxuICAgIGluc3RhbmNlX2xpc3RfcGxheWVyIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgICB6b25laWR1bnVtYmVyXG4gICAgfVxuICAgIGluc3RhbmNlX2xpc3QgfHwtLW97IGluc3RhbmNlX2xpc3RfcGxheWVyIDogSGFzLU1hbnlcbiAgICBpbnN0YW5jZV9saXN0IHx8LS1veyB6b25lIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgaW5zdGFuY2VfbGlzdCB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCB6b25lXG4gICAgfVxuICAgIGluc3RhbmNlX2xpc3RfcGxheWVyIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICB9XG4gICAgem9uZSB7XG4gICAgICAgICB6b25laWR1bnVtYmVyXG4gICAgfVxuICAgIGluc3RhbmNlX2xpc3QgfHwtLW97IGluc3RhbmNlX2xpc3RfcGxheWVyIDogSGFzLU1hbnlcbiAgICBpbnN0YW5jZV9saXN0IHx8LS1veyB6b25lIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [instance_list_player](../../schema/instances/instance_list_player.md) | id |
| Has-Many | zone | [zone](../../schema/zone/zone.md) | zoneidunumber |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Instance Identifier |
| zone | int | [Zone Identifier](../../../../server/zones/zone-list) |
| version | tinyint | Version |
| is_global | tinyint | Is Global: 0 = False, 1 = True |
| start_time | int | Start Time UNIX Timestamp |
| duration | int | Duration in Seconds |
| never_expires | tinyint | Never Expires: 0 = False, 1 = True |

