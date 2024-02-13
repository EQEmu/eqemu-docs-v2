# command_subsettings

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Subcommand Identifier |
| parent_command | varchar | [Parent Command](../../schema/admin/command_settings.md) |
| sub_command | varchar | Subcommand Identifier |
| access_level | int | [Required Status](../../../../server/player/status-levels) |
| top_level_aliases | varchar | Top Level Aliases |

