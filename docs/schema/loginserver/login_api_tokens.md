# login_api_tokens

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Login API Token Identifier |
| token | varchar | Token Identifier |
| can_write | int | Can Write: 0 = False, 1 = True |
| can_read | int | Can Read: 0 = False, 1 = True |
| created_at | datetime | Created At Date |
| updated_at | datetime | Updated At Date |

