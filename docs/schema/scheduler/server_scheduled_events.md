# server_scheduled_events

!!! info
	This page was last generated 2023.04.23

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Scheduled Event Identifier |
| description | varchar | Description |
| event_type | varchar | Event Type |
| event_data | text | Event Data |
| minute_start | int | Minute Start |
| hour_start | int | Hour Start |
| day_start | int | Day Start |
| month_start | int | Month Start |
| year_start | int | Year Start |
| minute_end | int | Minute End |
| hour_end | int | Hour End |
| day_end | int | Day End |
| month_end | int | Month End |
| year_end | int | Year End |
| cron_expression | varchar | CRON Expression |
| created_at | datetime | Created At |
| deleted_at | datetime | Deleted At |

