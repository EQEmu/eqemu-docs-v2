# guild_relations

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfcmVsYXRpb25zIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgZ3VpbGQxXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkMlxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgbGVhZGVyXG4gICAgICAgIHZhcmNoYXIgbW90ZF9zZXR0ZXJcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGd1aWxkcyB7XG4gICAgICAgIGludCBsZWFkZXJcbiAgICAgICAgdmFyY2hhciBtb3RkX3NldHRlclxuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgZ3VpbGRfcmVsYXRpb25zIHx8LS1veyBndWlsZHMgOiBPbmUtdG8tT25lXG4gICAgZ3VpbGRfcmVsYXRpb25zIHx8LS1veyBndWlsZHMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfcmVsYXRpb25zIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgZ3VpbGQxXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkMlxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgbGVhZGVyXG4gICAgICAgIHZhcmNoYXIgbW90ZF9zZXR0ZXJcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGd1aWxkcyB7XG4gICAgICAgIGludCBsZWFkZXJcbiAgICAgICAgdmFyY2hhciBtb3RkX3NldHRlclxuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgZ3VpbGRfcmVsYXRpb25zIHx8LS1veyBndWlsZHMgOiBPbmUtdG8tT25lXG4gICAgZ3VpbGRfcmVsYXRpb25zIHx8LS1veyBndWlsZHMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgZ3VpbGRfcmVsYXRpb25zIHtcbiAgICAgICAgbWVkaXVtaW50dW5zaWduZWQgZ3VpbGQxXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkMlxuICAgIH1cbiAgICBndWlsZHMge1xuICAgICAgICBpbnQgbGVhZGVyXG4gICAgICAgIHZhcmNoYXIgbW90ZF9zZXR0ZXJcbiAgICAgICAgaW50IGlkXG4gICAgfVxuICAgIGd1aWxkcyB7XG4gICAgICAgIGludCBsZWFkZXJcbiAgICAgICAgdmFyY2hhciBtb3RkX3NldHRlclxuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgZ3VpbGRfcmVsYXRpb25zIHx8LS1veyBndWlsZHMgOiBPbmUtdG8tT25lXG4gICAgZ3VpbGRfcmVsYXRpb25zIHx8LS1veyBndWlsZHMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagram}

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

