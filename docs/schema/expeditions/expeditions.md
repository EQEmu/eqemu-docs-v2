# expeditions

!!! info
	This page was last generated 2022.09.10

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXhwZWRpdGlvbnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBkeW5hbWljX3pvbmVfaWRcbiAgICB9XG4gICAgZHluYW1pY196b25lcyB7XG4gICAgICAgIGludCBpbnN0YW5jZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBjb21wYXNzX3pvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc2FmZV9yZXR1cm5fem9uZV9pZFxuICAgIH1cbiAgICBleHBlZGl0aW9ucyB8fC0tb3sgZHluYW1pY196b25lcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXhwZWRpdGlvbnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBkeW5hbWljX3pvbmVfaWRcbiAgICB9XG4gICAgZHluYW1pY196b25lcyB7XG4gICAgICAgIGludCBpbnN0YW5jZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBjb21wYXNzX3pvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc2FmZV9yZXR1cm5fem9uZV9pZFxuICAgIH1cbiAgICBleHBlZGl0aW9ucyB8fC0tb3sgZHluYW1pY196b25lcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZXhwZWRpdGlvbnMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBkeW5hbWljX3pvbmVfaWRcbiAgICB9XG4gICAgZHluYW1pY196b25lcyB7XG4gICAgICAgIGludCBpbnN0YW5jZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICBpbnR1bnNpZ25lZCBjb21wYXNzX3pvbmVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgc2FmZV9yZXR1cm5fem9uZV9pZFxuICAgIH1cbiAgICBleHBlZGl0aW9ucyB8fC0tb3sgZHluYW1pY196b25lcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | dynamic_zone_id | [dynamic_zones](../../schema/tasks/shared_task_dynamic_zones.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Expedition Identifier |
| dynamic_zone_id | int | [Dynamic Zone Identifier](../../schema/dynamic-zones/dynamic_zones.md) |
| add_replay_on_join | tinyint | Add Replay On Join: 0 = False, 1 = True |
| is_locked | tinyint | Is Locked: 0 = False, 1 = True |

