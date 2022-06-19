# pets_equipmentset

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcGV0c19lcXVpcG1lbnRzZXQge1xuICAgICAgICBpbnQgc2V0X2lkXG4gICAgfVxuICAgIHBldHMge1xuICAgICAgICBpbnQgbnBjSURcbiAgICAgICAgaW50IGVxdWlwbWVudHNldFxuICAgICAgICB2YXJjaGFyIHR5cGVcbiAgICB9XG4gICAgcGV0c19lcXVpcG1lbnRzZXRfZW50cmllcyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgICAgIGludCBzZXRfaWRcbiAgICB9XG4gICAgcGV0c19lcXVpcG1lbnRzZXQgfHwtLW97IHBldHMgOiBIYXMtTWFueVxuICAgIHBldHNfZXF1aXBtZW50c2V0IHx8LS1veyBwZXRzX2VxdWlwbWVudHNldF9lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcGV0c19lcXVpcG1lbnRzZXQge1xuICAgICAgICBpbnQgc2V0X2lkXG4gICAgfVxuICAgIHBldHMge1xuICAgICAgICBpbnQgbnBjSURcbiAgICAgICAgaW50IGVxdWlwbWVudHNldFxuICAgICAgICB2YXJjaGFyIHR5cGVcbiAgICB9XG4gICAgcGV0c19lcXVpcG1lbnRzZXRfZW50cmllcyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgICAgIGludCBzZXRfaWRcbiAgICB9XG4gICAgcGV0c19lcXVpcG1lbnRzZXQgfHwtLW97IHBldHMgOiBIYXMtTWFueVxuICAgIHBldHNfZXF1aXBtZW50c2V0IHx8LS1veyBwZXRzX2VxdWlwbWVudHNldF9lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcGV0c19lcXVpcG1lbnRzZXQge1xuICAgICAgICBpbnQgc2V0X2lkXG4gICAgfVxuICAgIHBldHMge1xuICAgICAgICBpbnQgbnBjSURcbiAgICAgICAgaW50IGVxdWlwbWVudHNldFxuICAgICAgICB2YXJjaGFyIHR5cGVcbiAgICB9XG4gICAgcGV0c19lcXVpcG1lbnRzZXRfZW50cmllcyB7XG4gICAgICAgIGludCBpdGVtX2lkXG4gICAgICAgIGludCBzZXRfaWRcbiAgICB9XG4gICAgcGV0c19lcXVpcG1lbnRzZXQgfHwtLW97IHBldHMgOiBIYXMtTWFueVxuICAgIHBldHNfZXF1aXBtZW50c2V0IHx8LS1veyBwZXRzX2VxdWlwbWVudHNldF9lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | set_id | [pets](../../schema/pets/pets.md) | equipmentset |
| Has-Many | set_id | [pets_equipmentset_entries](../../schema/pets/pets_equipmentset_entries.md) | set_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| set_id | int | Unique Pet Equipment Set Identifier |
| setname | varchar | Pet Equipment Set Name |
| nested_set | int | Nested Set Identifier |

