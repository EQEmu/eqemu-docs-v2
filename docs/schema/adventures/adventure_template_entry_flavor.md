# adventure_template_entry_flavor

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

```mermaid
erDiagram
    adventure_template_entry_flavor {
        intunsigned id
    }
    adventure_details {
        intunsigned id
        smallintunsigned adventure_id
        int instance_id
    }
    adventure_template_entry_flavor ||--o{ adventure_details : "One-to-One"


```


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | id | [adventure_details](../../schema/adventures/adventure_details.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | [Adventure Identifier](adventure_details.md) |
| text | varchar | Text |

