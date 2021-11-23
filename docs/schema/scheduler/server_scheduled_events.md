# server_scheduled_events

!!! info
	This page was last generated 2021.11.23

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int |  |
| description | varchar |  |
| event_type | varchar |  |
| event_data | text |  |
| minute_start | int |  |
| hour_start | int |  |
| day_start | int |  |
| month_start | int |  |
| year_start | int |  |
| minute_end | int |  |
| hour_end | int |  |
| day_end | int |  |
| month_end | int |  |
| year_end | int |  |
| cron_expression | varchar |  |
| created_at | datetime |  |
| deleted_at | datetime |  |

