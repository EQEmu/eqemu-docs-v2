# rule_sets

## Relationships

```mermaid
erDiagram
    rule_sets {
        tinyintunsigned ruleset_id
    }
    rule_values {
        tinyintunsigned ruleset_id
    }
    rule_sets ||--o{ rule_values : "Has-Many"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | ruleset_id | [rule_values](../../schema/rules/rule_values.md) | ruleset_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| ruleset_id | tinyint | Unique Rule Set Identifier |
| name | varchar | Name |

