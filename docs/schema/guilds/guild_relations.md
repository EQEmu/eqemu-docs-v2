# guild_relations

!!! info
	This page was last generated 2023.06.12

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfcmVsYXRpb25zIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgZ3VpbGQxXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkMlxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGxlYWRlclxuICAgICAgICB2YXJjaGFyIG1vdGRfc2V0dGVyXG4gICAgfVxuICAgIGd1aWxkX3JlbGF0aW9ucyB8fC0tb3sgZ3VpbGRzIDogXCJPbmUtdG8tT25lXCJcbiAgICBndWlsZF9yZWxhdGlvbnMgfHwtLW97IGd1aWxkcyA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfcmVsYXRpb25zIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgZ3VpbGQxXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkMlxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGxlYWRlclxuICAgICAgICB2YXJjaGFyIG1vdGRfc2V0dGVyXG4gICAgfVxuICAgIGd1aWxkX3JlbGF0aW9ucyB8fC0tb3sgZ3VpbGRzIDogXCJPbmUtdG8tT25lXCJcbiAgICBndWlsZF9yZWxhdGlvbnMgfHwtLW97IGd1aWxkcyA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfcmVsYXRpb25zIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgZ3VpbGQxXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkMlxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGxlYWRlclxuICAgICAgICB2YXJjaGFyIG1vdGRfc2V0dGVyXG4gICAgfVxuICAgIGd1aWxkX3JlbGF0aW9ucyB8fC0tb3sgZ3VpbGRzIDogXCJPbmUtdG8tT25lXCJcbiAgICBndWlsZF9yZWxhdGlvbnMgfHwtLW97IGd1aWxkcyA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | guild1 | [guilds](../../schema/guilds/guilds.md) | id |
| One-to-One | guild2 | [guilds](../../schema/guilds/guilds.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| guild1 | mediumint | [Unique Guild Identifier 1](guilds.md) |
| guild2 | mediumint | [Unique Guild Identifier 2](guilds.md) |
| relation | tinyint | Relation |

