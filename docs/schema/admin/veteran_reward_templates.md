# veteran_reward_templates

## Relationships

```mermaid
erDiagram
    veteran_reward_templates {
        intunsigned claim_id
        intunsigned item_id
    }
    account_rewards {
        intunsigned account_id
        intunsigned reward_id
    }
    items {
        int id
        int book
        varchar name
        int recasttype
        int icon
        mediumint bardeffect
        int clickeffect
        int focuseffect
        int proceffect
        int scrolleffect
        int worneffect
    }
    veteran_reward_templates ||--o{ account_rewards : "Has-Many"
    veteran_reward_templates ||--o{ items : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | claim_id | [account_rewards](../../schema/account/account_rewards.md) | reward_id |
| One-to-One | item_id | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| claim_id | int | Unique Claim Identifier |
| name | varchar | Name |
| item_id | int | [Item Identifier](../../schema/items/items.md) |
| charges | smallint | Charges |
| reward_slot | tinyint | Reward Slot |

