# perl_event_export_settings

!!! info
	This page was last generated 2024.02.07

## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| event_id | int | [Unique Perl Event Identifier](https://docs.eqemu.io/quest-api/events/) |
| event_description | varchar | Event Description |
| export_qglobals | smallint | Export QGlobals: 0 = False, 1 = True |
| export_mob | smallint | Export Mob: 0 = False, 1 = True |
| export_zone | smallint | Export Zone: 0 = False, 1 = True |
| export_item | smallint | Export Item: 0 = False, 1 = True |
| export_event | smallint | Export Event: 0 = False, 1 = True |

