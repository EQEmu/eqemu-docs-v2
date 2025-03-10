# rule_values

## Relationships

```mermaid
erDiagram
    rule_values {
        tinyintunsigned ruleset_id
    }
    rule_sets {
        tinyintunsigned ruleset_id
    }
    rule_values ||--o{ rule_sets : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | ruleset_id | [rule_sets](../../schema/rules/rule_sets.md) | ruleset_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| ruleset_id | tinyint | [Rule Set Identifier](rule_sets.md) |
| rule_name | varchar | Rule Name |
| rule_value | varchar | Rule Value |
| notes | text | Notes |

