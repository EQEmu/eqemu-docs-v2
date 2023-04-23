# account_rewards

!!! info
	This page was last generated 2023.04.23

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWNjb3VudF9yZXdhcmRzIHtcbiAgICAgICAgaW50dW5zaWduZWQgcmV3YXJkX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGFjY291bnRfaWRcbiAgICB9XG4gICAgdmV0ZXJhbl9yZXdhcmRfdGVtcGxhdGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgY2xhaW1faWRcbiAgICAgICAgaW50dW5zaWduZWQgaXRlbV9pZFxuICAgIH1cbiAgICBhY2NvdW50X3Jld2FyZHMgfHwtLW97IHZldGVyYW5fcmV3YXJkX3RlbXBsYXRlcyA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWNjb3VudF9yZXdhcmRzIHtcbiAgICAgICAgaW50dW5zaWduZWQgcmV3YXJkX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGFjY291bnRfaWRcbiAgICB9XG4gICAgdmV0ZXJhbl9yZXdhcmRfdGVtcGxhdGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgY2xhaW1faWRcbiAgICAgICAgaW50dW5zaWduZWQgaXRlbV9pZFxuICAgIH1cbiAgICBhY2NvdW50X3Jld2FyZHMgfHwtLW97IHZldGVyYW5fcmV3YXJkX3RlbXBsYXRlcyA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYWNjb3VudF9yZXdhcmRzIHtcbiAgICAgICAgaW50dW5zaWduZWQgcmV3YXJkX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGFjY291bnRfaWRcbiAgICB9XG4gICAgdmV0ZXJhbl9yZXdhcmRfdGVtcGxhdGVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgY2xhaW1faWRcbiAgICAgICAgaW50dW5zaWduZWQgaXRlbV9pZFxuICAgIH1cbiAgICBhY2NvdW50X3Jld2FyZHMgfHwtLW97IHZldGVyYW5fcmV3YXJkX3RlbXBsYXRlcyA6IFwiT25lLXRvLU9uZVwiXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | reward_id | [veteran_reward_templates](../../schema/admin/veteran_reward_templates.md) | claim_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| account_id | int | [Account Identifier](account.md) |
| reward_id | int | [Veteran Reward Identifier](../../schema/admin/veteran_reward_templates.md) |
| amount | int | Amount |

