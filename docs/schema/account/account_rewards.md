# account_rewards

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    account_rewards {
        intunsigned account_id
        intunsigned reward_id
    }
    veteran_reward_templates {
        intunsigned claim_id
        intunsigned item_id
    }
    account_rewards ||--o{ veteran_reward_templates : "One-to-One"


```


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

