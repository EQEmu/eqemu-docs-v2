# character_alt_currency

!!! info
	This page was last generated 2022.06.19

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX2FsdF9jdXJyZW5jeSB7XG4gICAgICAgIGludHVuc2lnbmVkIGNoYXJfaWRcbiAgICAgICAgaW50dW5zaWduZWQgY3VycmVuY3lfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBhbHRlcm5hdGVfY3VycmVuY3kge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGl0ZW1faWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2FsdF9jdXJyZW5jeSB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgY2hhcmFjdGVyX2FsdF9jdXJyZW5jeSB8fC0tb3sgYWx0ZXJuYXRlX2N1cnJlbmN5IDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX2FsdF9jdXJyZW5jeSB7XG4gICAgICAgIGludHVuc2lnbmVkIGNoYXJfaWRcbiAgICAgICAgaW50dW5zaWduZWQgY3VycmVuY3lfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBhbHRlcm5hdGVfY3VycmVuY3kge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGl0ZW1faWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2FsdF9jdXJyZW5jeSB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgY2hhcmFjdGVyX2FsdF9jdXJyZW5jeSB8fC0tb3sgYWx0ZXJuYXRlX2N1cnJlbmN5IDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgY2hhcmFjdGVyX2FsdF9jdXJyZW5jeSB7XG4gICAgICAgIGludHVuc2lnbmVkIGNoYXJfaWRcbiAgICAgICAgaW50dW5zaWduZWQgY3VycmVuY3lfaWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2RhdGEge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgICAgICB2YXJjaGFyIG5hbWVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pbnN0YW5jZVxuICAgICAgICB2YXJjaGFyIG5hbmVcbiAgICAgICAgaW50dW5zaWduZWQgem9uZV9pZFxuICAgIH1cbiAgICBhbHRlcm5hdGVfY3VycmVuY3kge1xuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IGl0ZW1faWRcbiAgICB9XG4gICAgY2hhcmFjdGVyX2FsdF9jdXJyZW5jeSB8fC0tb3sgY2hhcmFjdGVyX2RhdGEgOiBPbmUtdG8tT25lXG4gICAgY2hhcmFjdGVyX2FsdF9jdXJyZW5jeSB8fC0tb3sgYWx0ZXJuYXRlX2N1cnJlbmN5IDogT25lLXRvLU9uZVxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | char_id | [character_data](../../schema/characters/character_data.md) | id |
| One-to-One | currency_id | [alternate_currency](../../schema/alternate-currency/alternate_currency.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| char_id | int | [Character Identifier](character_data.md) |
| currency_id | int | [Currency Identifier](../../schema/alternate-currency/alternate_currency.md) |
| amount | int | Amount |

