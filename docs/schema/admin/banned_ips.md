# banned_ips

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYmFubmVkX2lwcyB7XG4gICAgICAgIHZhcmNoYXIgaXBfYWRkcmVzc1xuICAgIH1cbiAgICBhY2NvdW50X2lwIHtcbiAgICAgICAgaW50IGFjY2lkXG4gICAgICAgIHZhcmNoYXIgaXBcbiAgICB9XG4gICAgYmFubmVkX2lwcyB8fC0tb3sgYWNjb3VudF9pcCA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYmFubmVkX2lwcyB7XG4gICAgICAgIHZhcmNoYXIgaXBfYWRkcmVzc1xuICAgIH1cbiAgICBhY2NvdW50X2lwIHtcbiAgICAgICAgaW50IGFjY2lkXG4gICAgICAgIHZhcmNoYXIgaXBcbiAgICB9XG4gICAgYmFubmVkX2lwcyB8fC0tb3sgYWNjb3VudF9pcCA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgYmFubmVkX2lwcyB7XG4gICAgICAgIHZhcmNoYXIgaXBfYWRkcmVzc1xuICAgIH1cbiAgICBhY2NvdW50X2lwIHtcbiAgICAgICAgaW50IGFjY2lkXG4gICAgICAgIHZhcmNoYXIgaXBcbiAgICB9XG4gICAgYmFubmVkX2lwcyB8fC0tb3sgYWNjb3VudF9pcCA6IFwiSGFzLU1hbnlcIlxuXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOnRydWUsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjp0cnVlfQ==){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | ip_address | [account_ip](../../schema/account/account_ip.md) | ip |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| ip_address | varchar | [IP Address](../../schema/account/account_ip.md) |
| notes | varchar | Ban reason |

