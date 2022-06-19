# bot_guild_members

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2d1aWxkX21lbWJlcnMge1xuICAgICAgICBpbnQgYm90X2lkXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGJvdF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBvd25lcl9pZFxuICAgICAgICBzbWFsbGludCB6b25lX2lkXG4gICAgfVxuICAgIGd1aWxkcyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgbGVhZGVyXG4gICAgICAgIHZhcmNoYXIgbW90ZF9zZXR0ZXJcbiAgICB9XG4gICAgYm90X2d1aWxkX21lbWJlcnMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuICAgIGJvdF9ndWlsZF9tZW1iZXJzIHx8LS1veyBndWlsZHMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2d1aWxkX21lbWJlcnMge1xuICAgICAgICBpbnQgYm90X2lkXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGJvdF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBvd25lcl9pZFxuICAgICAgICBzbWFsbGludCB6b25lX2lkXG4gICAgfVxuICAgIGd1aWxkcyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgbGVhZGVyXG4gICAgICAgIHZhcmNoYXIgbW90ZF9zZXR0ZXJcbiAgICB9XG4gICAgYm90X2d1aWxkX21lbWJlcnMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuICAgIGJvdF9ndWlsZF9tZW1iZXJzIHx8LS1veyBndWlsZHMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYm90X2d1aWxkX21lbWJlcnMge1xuICAgICAgICBpbnQgYm90X2lkXG4gICAgICAgIG1lZGl1bWludHVuc2lnbmVkIGd1aWxkX2lkXG4gICAgfVxuICAgIGJvdF9kYXRhIHtcbiAgICAgICAgaW50dW5zaWduZWQgc3BlbGxzX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGJvdF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBvd25lcl9pZFxuICAgICAgICBzbWFsbGludCB6b25lX2lkXG4gICAgfVxuICAgIGd1aWxkcyB7XG4gICAgICAgIGludCBpZFxuICAgICAgICBpbnQgbGVhZGVyXG4gICAgICAgIHZhcmNoYXIgbW90ZF9zZXR0ZXJcbiAgICB9XG4gICAgYm90X2d1aWxkX21lbWJlcnMgfHwtLW97IGJvdF9kYXRhIDogT25lLXRvLU9uZVxuICAgIGJvdF9ndWlsZF9tZW1iZXJzIHx8LS1veyBndWlsZHMgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | bot_id | [bot_data](../../schema/bots/bot_data.md) | bot_id |
| One-to-One | guild_id | [guilds](../../schema/guilds/guilds.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| bot_id | int | [Bot Identifier](bot_data.md) |
| guild_id | mediumint | [Guild Identifier](../../schema/guilds/guilds.md) |
| rank | tinyint | [Guild Rank](../../../../categories/player/guild-ranks) |
| tribute_enable | tinyint | Tribute Enabled: 0 = False, 1= True |
| total_tribute | int | Total Tribute |
| last_tribute | int | Last Tribute |
| banker | tinyint | Banker: 0 = False, 1 = True |
| public_note | text | Public Note |
| alt | tinyint | Alt: 0 = False, 1 = True |

