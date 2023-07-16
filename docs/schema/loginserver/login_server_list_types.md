# login_server_list_types

!!! info
	This page was last generated 2023.07.15

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9naW5fc2VydmVyX2xpc3RfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBsb2dpbl93b3JsZF9zZXJ2ZXJzIHtcbiAgICAgICAgaW50IGxvZ2luX3NlcnZlcl9saXN0X3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBsYXN0X2lwX2FkZHJlc3NcbiAgICAgICAgaW50IGxvZ2luX3NlcnZlcl9hZG1pbl9pZFxuICAgIH1cbiAgICBsb2dpbl9zZXJ2ZXJfbGlzdF90eXBlcyB8fC0tb3sgbG9naW5fd29ybGRfc2VydmVycyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9naW5fc2VydmVyX2xpc3RfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBsb2dpbl93b3JsZF9zZXJ2ZXJzIHtcbiAgICAgICAgaW50IGxvZ2luX3NlcnZlcl9saXN0X3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBsYXN0X2lwX2FkZHJlc3NcbiAgICAgICAgaW50IGxvZ2luX3NlcnZlcl9hZG1pbl9pZFxuICAgIH1cbiAgICBsb2dpbl9zZXJ2ZXJfbGlzdF90eXBlcyB8fC0tb3sgbG9naW5fd29ybGRfc2VydmVycyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9naW5fc2VydmVyX2xpc3RfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBsb2dpbl93b3JsZF9zZXJ2ZXJzIHtcbiAgICAgICAgaW50IGxvZ2luX3NlcnZlcl9saXN0X3R5cGVfaWRcbiAgICAgICAgdmFyY2hhciBsYXN0X2lwX2FkZHJlc3NcbiAgICAgICAgaW50IGxvZ2luX3NlcnZlcl9hZG1pbl9pZFxuICAgIH1cbiAgICBsb2dpbl9zZXJ2ZXJfbGlzdF90eXBlcyB8fC0tb3sgbG9naW5fd29ybGRfc2VydmVycyA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [login_world_servers](../../schema/loginserver/login_world_servers.md) | login_server_list_type_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique List Type Identifier |
| description | varchar | Description |

