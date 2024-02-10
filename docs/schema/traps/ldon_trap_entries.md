# ldon_trap_entries

## Relationships

```mermaid
erDiagram
    ldon_trap_entries {
        intunsigned trap_id
    }
    ldon_trap_templates {
        intunsigned id
        smallintunsigned spell_id
    }
    ldon_trap_entries ||--o{ ldon_trap_templates : "One-to-One"


```


| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | trap_id | [ldon_trap_templates](../../schema/traps/ldon_trap_templates.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique LDoN Trap Entry Identifier |
| trap_id | int | [Trap Identifier](ldon_trap_templates.md) |

