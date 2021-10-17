# Logging for Developers

The main log function used everywhere in the code funnels through `Out` or `Log.Out` and is macroed to simply `Log`

### Example

```text
void Out(Logs::DebugLevel debug_level, uint16 log_category, std::string message, ...);

Usage:

Log(Logs::General, Logs::Zone_Server, "Loading server configuration...");
```

The macro is necessary because we check for whether the category is enabled or not everywhere we call a log function, this was to save massive overhead of pushing/popping strings on the call stack through a function name only to return immediately out.

```text
#define Log(debug_level, log_category, message, ...) do {\
	if (LogSys.log_settings[log_category].is_category_enabled == 1)\
		LogSys.Out(debug_level, log_category, message, ##__VA_ARGS__);\
} while (0)
```

## Adding New Categories

* For development purposes, you may want to add a new category, this is very simple to do
* An example of a category being added can be seen at this commit: [How to add a category](https://github.com/EQEmu/Server/commit/a46c0ee7e2dcf094c4b0e4d9cb91525443c19c5b)
* An entry needs to be added to the enum and the constant to add a new category - and a database update is not required for it to be functional, but it is required for folks to be able to set the default settings in their server

## Adding Default Values

In `eqemu_logsys.cpp` you can set default values for the log system initialization routine

* [https://github.com/EQEmu/Server/blob/master/common/eqemu\_logsys.cpp\#L90](https://github.com/EQEmu/Server/blob/master/common/eqemu_logsys.cpp#L90)

```text
void EQEmuLogSys::LoadLogSettingsDefaults()
{
	/* Get Executable platform currently running this code (Zone/World/etc) */
	log_platform = GetExecutablePlatformInt();

	/* Zero out Array */
	memset(log_settings, 0, sizeof(LogSettings) * Logs::LogCategory::MaxCategoryID);

	/* Set Defaults */
	log_settings[Logs::World_Server].log_to_console = Logs::General;
	log_settings[Logs::Zone_Server].log_to_console = Logs::General;
	log_settings[Logs::QS_Server].log_to_console = Logs::General;
	log_settings[Logs::UCS_Server].log_to_console = Logs::General;
	log_settings[Logs::Crash].log_to_console = Logs::General;
	log_settings[Logs::MySQLError].log_to_console = Logs::General;
	log_settings[Logs::Login_Server].log_to_console = Logs::General;
	log_settings[Logs::Headless_Client].log_to_console = Logs::General;

	/* Set Category enabled status on defaults */
	log_settings[Logs::World_Server].is_category_enabled = 1;
	log_settings[Logs::Zone_Server].is_category_enabled = 1;
	log_settings[Logs::QS_Server].is_category_enabled = 1;
	log_settings[Logs::UCS_Server].is_category_enabled = 1;
	log_settings[Logs::Crash].is_category_enabled = 1;
	log_settings[Logs::MySQLError].is_category_enabled = 1;
	log_settings[Logs::Login_Server].is_category_enabled = 1;
...
```

## Default Category Colors

Color settings for GMSay and console respectively can be managed in eqemu\_logsys.cpp

