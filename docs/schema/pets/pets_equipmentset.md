# pets_equipmentset

## Relationship Diagram
[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcGV0c19lcXVpcG1lbnRzZXQge1xuICAgICAgICBpbnQgc2V0X2lkXG4gICAgfVxuICAgIHBldHNfZXF1aXBtZW50c2V0X2VudHJpZXMge1xuICAgICAgICBpbnQgc2V0X2lkXG4gICAgfVxuICAgIHBldHNfZXF1aXBtZW50c2V0IHx8LS1veyBwZXRzX2VxdWlwbWVudHNldF9lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcGV0c19lcXVpcG1lbnRzZXQge1xuICAgICAgICBpbnQgc2V0X2lkXG4gICAgfVxuICAgIHBldHNfZXF1aXBtZW50c2V0X2VudHJpZXMge1xuICAgICAgICBpbnQgc2V0X2lkXG4gICAgfVxuICAgIHBldHNfZXF1aXBtZW50c2V0IHx8LS1veyBwZXRzX2VxdWlwbWVudHNldF9lbnRyaWVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

## Relationships
| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | set_id | pets_equipmentset_entries | set_id |


## Schema
| Column | Data Type | Description |
| :--- | :--- | :--- |
| set_id | int | Unique Pet Equipment Set Identifier |
| setname | varchar | Pet Equipment Set Name |
| nested_set | int | Nested Set Identifier |

