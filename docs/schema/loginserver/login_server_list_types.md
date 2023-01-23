# login_server_list_types

!!! info
	This page was last generated 2023.01.22

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9naW5fc2VydmVyX2xpc3RfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBsb2dpbl93b3JsZF9zZXJ2ZXJzIHtcbiAgICAgICAgaW50IGxvZ2luX3NlcnZlcl9hZG1pbl9pZFxuICAgICAgICBpbnQgbG9naW5fc2VydmVyX2xpc3RfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIGxhc3RfaXBfYWRkcmVzc1xuICAgIH1cbiAgICBsb2dpbl9zZXJ2ZXJfbGlzdF90eXBlcyB8fC0tb3sgbG9naW5fd29ybGRfc2VydmVycyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9naW5fc2VydmVyX2xpc3RfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBsb2dpbl93b3JsZF9zZXJ2ZXJzIHtcbiAgICAgICAgaW50IGxvZ2luX3NlcnZlcl9hZG1pbl9pZFxuICAgICAgICBpbnQgbG9naW5fc2VydmVyX2xpc3RfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIGxhc3RfaXBfYWRkcmVzc1xuICAgIH1cbiAgICBsb2dpbl9zZXJ2ZXJfbGlzdF90eXBlcyB8fC0tb3sgbG9naW5fd29ybGRfc2VydmVycyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9naW5fc2VydmVyX2xpc3RfdHlwZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBsb2dpbl93b3JsZF9zZXJ2ZXJzIHtcbiAgICAgICAgaW50IGxvZ2luX3NlcnZlcl9hZG1pbl9pZFxuICAgICAgICBpbnQgbG9naW5fc2VydmVyX2xpc3RfdHlwZV9pZFxuICAgICAgICB2YXJjaGFyIGxhc3RfaXBfYWRkcmVzc1xuICAgIH1cbiAgICBsb2dpbl9zZXJ2ZXJfbGlzdF90eXBlcyB8fC0tb3sgbG9naW5fd29ybGRfc2VydmVycyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [login_world_servers](../../schema/loginserver/login_world_servers.md) | login_server_list_type_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique List Type Identifier |
| description | varchar | Description |

