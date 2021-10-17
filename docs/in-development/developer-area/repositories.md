# Repositories

The purpose of this guide to serve as a reference for our source generated repositories and our use of the repository pattern in the server codebase

## What is the Repository Pattern?

"The simplest approach, especially with an existing system, is to create a new Repository implementation for each business object you need to store to or retrieve from your persistence layer. Further, you should only implement the specific methods you are calling in your application."

{% hint style="info" %}
Reference [https://deviq.com/repository-pattern/](https://deviq.com/repository-pattern/)
{% endhint %}

![Repository Pattern Illustration A](<../../.gitbook/assets/image (9).png>)

![Repository Pattern Illustration B ](<../../.gitbook/assets/image (10).png>)

The underlying persistence layer could be anything, file, memory, remote storage, but in most cases relevant to what we are using repositories for right now in this article pertain to the database (MySQL)

### Benefits

* Simplified interaction with the database, instead of manually looking up columns, your IDE will autocomplete the struct fields that you are working with
* Heavily reducing the mental overhead of having to interact with the database at the programmatic level
* It creates an object representation of our tables (persistence layer) with a 1:1 mapping with what is in the source by having row values stored in a struct (Data Transfer Object)
* We reduce manual overhead by having code generation generate our database table representations
* Ease of maintenance; whenever we make schema changes, rerunning the repository updates fields and subsequent methods keeping code changes minimal

### **Methods Implemented in Base Repository**

* Single insertion covered by `InsertOne`
* Single update covered by `UpdateOne`
* Single delete covered by `DeleteOne`
* Single select covered by `FindOne`
* Bulk selection method via filtered `GetWhere(std::string where_filter)`
* Bulk deletion method via filtered `DeleteWhere(std::string where_filter)`
* Bulk insertion methods handled automatically via `InsertMany`
* Select all covered by `All`

### **Code Generation: Repository Generator**

The repository generator can be found in the source via the following path

```bash
./utils/scripts/generators/repository-generator.pl
```

{% tabs %}
{% tab title="Linux bash" %}
```bash
# Command Structure
perl ~/code/utils/scripts/generators/repository-generator.pl [server-location] [table_name|all]

# Generate everything
perl ~/code/utils/scripts/generators/repository-generator.pl ~/server/ all

# Only generate a repository for the account table
perl ~/code/utils/scripts/generators/repository-generator.pl ~/server/ account 
```
{% endtab %}

{% tab title="Windows cmd" %}
```
# Command Structure
# Must be run from the root eqemu source folder
# [server-location] is the full path to server's world.exe location
E:\EQEmu\src>perl utils/scripts/generators/repository-generator.pl [server-location] [table_name|all]

# Generate everything
E:\EQEmu\src>perl utils/scripts/generators/repository-generator.pl E:\EQEmu\build\bin\Debug\ all

# Only generate a repository for the account table
E:\EQEmu\src>perl utils/scripts/generators/repository-generator.pl E:\EQEmu\build\bin\Debug\ account

```
{% endtab %}
{% endtabs %}

The generator works by examining your database tables found in the connection properties registered in the server config **eqemu_config** and will generate a **struct** object that tries to represent the database data types as close as possible

### Extending the Base Repository

The base repository is not meant to be touched manually; it is meant to provide generic scaffolding to the underlying persistence layer with generic methods typically needed

To implement your own customary methods you simply add your own methods in the extended repository, for each repository generated it outputs two files, the base repository (immutable) and the extended repository (mutable)

For example when we generate repositories from a single table we get both

`./common/repositories/base/base_instance_list_repository.h`

and

`./common/repositories/instance_list_repository.h`

The class structure in the non-base repository looks something like this 

```cpp
#include "../database.h"
#include "../string_util.h"
#include "base/base_instance_list_repository.h"

class InstanceListRepository: public BaseInstanceListRepository {
public:

    /**
     * This file was auto generated and can be modified and extended upon
     *
     * Base repository methods are automatically
     * generated in the "base" version of this repository. The base repository
     * is immutable and to be left untouched, while methods in this class
     * are used as extension methods for more specific persistence-layer
     * accessors or mutators.
     *
     * Base Methods (Subject to be expanded upon in time)
     *
     * Note: Not all tables are designed appropriately to fit functionality with all base methods
     *
     * InsertOne
     * UpdateOne
     * DeleteOne
     * FindOne
     * GetWhere(std::string where_filter)
     * DeleteWhere(std::string where_filter)
     * InsertMany
     * All
     *
     * Example custom methods in a repository
     *
     * InstanceListRepository::GetByZoneAndVersion(int zone_id, int zone_version)
     * InstanceListRepository::GetWhereNeverExpires()
     * InstanceListRepository::GetWhereXAndY()
     * InstanceListRepository::DeleteWhereXAndY()
     *
     * Most of the above could be covered by base methods, but if you as a developer
     * find yourself re-using logic for other parts of the code, its best to just make a
     * method that can be re-used easily elsewhere especially if it can use a base repository
     * method and encapsulate filters there
     */

	// Custom extended repository methods here

};
```

Pretty bare right? That's because it should be, this leaves the base repository free to be regenerated later by the repository generator if updates are supplied in the future and allows you to make all of your custom update / delete / find methods in the extended repository without requiring a ton of work to merge or update repositories

## Using Repositories in Server Code

Below we have an example implemented using our CLI menu interface to simply test some code

In **world_server_command_handler.cpp **we've registered a test command for testing repository code 

```bash
/**
 * Register commands
 */
function_map["test:test"]       = &WorldserverCommandHandler::TestCommand;
function_map["test:expansion"]  = &WorldserverCommandHandler::ExpansionTestCommand;
function_map["test:repository"] = &WorldserverCommandHandler::TestRepository;
```

### Base Repository Contents (Truncated)

```cpp
class BaseInstanceListRepository {
public:
	struct InstanceList {
		int id;
		int zone;
		int version;
		int is_global;
		int start_time;
		int duration;
		int never_expires;
	};

	static std::string PrimaryKey()
	{
		return std::string("id");
	}

	static std::vector<std::string> Columns()
	{
		return {
			"id",
			"zone",
			"version",
			"is_global",
			"start_time",
			"duration",
			"never_expires",
		};
	}
	...
```

### Command Code

```cpp
/**
 * @param argc
 * @param argv
 * @param cmd
 * @param description
 */
void TestRepository(int argc, char **argv, argh::parser &cmd, std::string &description)
{
	description = "Test command";

	if (cmd[{"-h", "--help"}]) {
		return;
	}

	/**
	 * Insert one
	 */
	auto instance_list_entry = InstanceListRepository::NewEntity();

	instance_list_entry.zone          = 999;
	instance_list_entry.version       = 1;
	instance_list_entry.is_global     = 1;
	instance_list_entry.start_time    = 0;
	instance_list_entry.duration      = 0;
	instance_list_entry.never_expires = 1;

	auto instance_list_inserted = InstanceListRepository::InsertOne(instance_list_entry);

	LogInfo("Inserted ID is [{}] zone [{}]", instance_list_inserted.id, instance_list_inserted.zone);

	/**
	 * Find one
	 */
	auto found_instance_list = InstanceListRepository::FindOne(instance_list_inserted.id);

	LogInfo("Found ID is [{}] zone [{}]", found_instance_list.id, found_instance_list.zone);

	/**
	 * Update one
	 */
	LogInfo("Updating instance id [{}] zone [{}]", found_instance_list.id, found_instance_list.zone);

	int update_instance_list_count = InstanceListRepository::UpdateOne(found_instance_list);

	found_instance_list.zone = 777;

	LogInfo(
		"Updated instance id [{}] zone [{}] affected [{}]",
		found_instance_list.id,
		found_instance_list.zone,
		update_instance_list_count
	);


	/**
	 * Delete one
	 */
	int deleted = InstanceListRepository::DeleteOne(found_instance_list.id);

	LogInfo("Deleting one instance [{}] deleted count [{}]", found_instance_list.id, deleted);

	/**
	 * Insert many
	 */
	std::vector<InstanceListRepository::InstanceList> instance_lists;

	auto instance_list_entry_bulk = InstanceListRepository::NewEntity();

	instance_list_entry_bulk.zone          = 999;
	instance_list_entry_bulk.version       = 1;
	instance_list_entry_bulk.is_global     = 1;
	instance_list_entry_bulk.start_time    = 0;
	instance_list_entry_bulk.duration      = 0;
	instance_list_entry_bulk.never_expires = 1;

	for (int i = 0; i < 10; i++) {
		instance_lists.push_back(instance_list_entry_bulk);
	}

	/**
	 * Fetch all
	 */
	int inserted_count = InstanceListRepository::InsertMany(instance_lists);

	LogInfo("Bulk insertion test, inserted [{}]", inserted_count);

	for (auto &entry: InstanceListRepository::GetWhere(fmt::format("zone = {}", 999))) {
		LogInfo("Iterating through entry id [{}] zone [{}]", entry.id, entry.zone);
	}

	/**
	 * Delete where
	 */
	int deleted_count = InstanceListRepository::DeleteWhere(fmt::format("zone = {}", 999));

	LogInfo("Bulk deletion test, deleted [{}]", deleted_count);

}
```

### Command Output

```bash
./bin/world test:repository
[MySQL Query] INSERT INTO instance_list (zone, version, is_global, start_time, duration, never_expires)  VALUES (999,1,1,0,0,1) (1 row affected) (0.000942s)
[WorldServer] [Info] Inserted ID is [3669] zone [999]
[MySQL Query] SELECT id, zone, version, is_global, start_time, duration, never_expires FROM instance_list WHERE id = 3669 LIMIT 1 (1 row returned) (0.000274s)
[WorldServer] [Info] Found ID is [3669] zone [999]
[WorldServer] [Info] Updating instance id [3669] zone [999]
[MySQL Query] UPDATE instance_list SET zone = 999, version = 1, is_global = 1, start_time = 0, duration = 0, never_expires = 1 WHERE id = 3669 (1 row affected) (0.000188s)
[WorldServer] [Info] Updated instance id [3669] zone [777] affected [1]
[MySQL Query] DELETE FROM instance_list WHERE id = 3669 (1 row affected) (0.000676s)
[WorldServer] [Info] Deleting one instance [3669] deleted count [1]
[MySQL Query] INSERT INTO instance_list (zone, version, is_global, start_time, duration, never_expires)  VALUES (999,1,1,0,0,1),(999,1,1,0,0,1),(999,1,1,0,0,1),(999,1,1,0,0,1),(999,1,1,0,0,1),(999,1,1,0,0,1),(999,1,1,0,0,1),(999,1,1,0,0,1),(999,1,1,0,0,1),(999,1,1,0,0,1) (10 rows affected) (0.000456s)
[WorldServer] [Info] Bulk insertion test, inserted [10]
[MySQL Query] SELECT id, zone, version, is_global, start_time, duration, never_expires FROM instance_list WHERE zone = 999 (10 rows returned) (0.000219s)
[WorldServer] [Info] Iterating through entry id [3670] zone [999]
[WorldServer] [Info] Iterating through entry id [3671] zone [999]
[WorldServer] [Info] Iterating through entry id [3672] zone [999]
[WorldServer] [Info] Iterating through entry id [3673] zone [999]
[WorldServer] [Info] Iterating through entry id [3674] zone [999]
[WorldServer] [Info] Iterating through entry id [3675] zone [999]
[WorldServer] [Info] Iterating through entry id [3676] zone [999]
[WorldServer] [Info] Iterating through entry id [3677] zone [999]
[WorldServer] [Info] Iterating through entry id [3678] zone [999]
[WorldServer] [Info] Iterating through entry id [3679] zone [999]
[MySQL Query] DELETE FROM instance_list WHERE zone = 999 (10 rows affected) (0.000459s)
[WorldServer] [Info] Bulk deletion test, deleted [10]
```

You can see how we've had to use zero raw SQL to interact with the database in our actual domain logic; we can keep all of the SQL queries and interactions hidden behind methods created in our repositories and simply pass needed methods to those repository methods

## Another Example

Another real use example is where we need some additional criteria to pull some results from the database. Instead of querying for grids by zone using raw SQL we want to encapsulate some of this so it can be easily used in our domain logic. We could use `GetWhere` quickly, but to make a more re-usable method we're going to create some new methods for our pathing grids

* static std::vector GetZoneGrids(int zone_id) 
* static Grid GetGrid(const std::vector \&grids, int grid_id)

{% code title="grid_repository.h" %}
```cpp

	// Custom extended repository methods here

	static std::vector<Grid> GetZoneGrids(int zone_id)
	{
		std::vector<Grid> grids;

		auto results = content_db.QueryDatabase(
			fmt::format(
				"{} WHERE zoneid = {}",
				BaseSelect(),
				zone_id
			)
		);

		for (auto row = results.begin(); row != results.end(); ++row) {
			Grid entry{};

			entry.id     = atoi(row[0]);
			entry.zoneid = atoi(row[1]);
			entry.type   = atoi(row[2]);
			entry.type2  = atoi(row[3]);

			grids.push_back(entry);
		}

		return grids;
	}

	static Grid GetGrid(
		const std::vector<Grid> &grids,
		int grid_id
	)
	{
		for (auto &row : grids) {
			if (row.id == grid_id) {
				return row;
			}
		}

		return NewEntity();
	}
```
{% endcode %}

Since we want these records to be cached at the zone level; we create two properties to eventually load via data that is zone contextual

{% code title="zone.h" %}
```cpp
class Zone {
...
	std::vector<GridRepository::Grid>             zone_grids;
	std::vector<GridEntriesRepository::GridEntry> zone_grid_entries;
```
{% endcode %}

We created a function that during zone initialization we call **LoadGrids **so we can reuse it in other parts of the code if we wanted to reload grid data for any reason

```cpp
void Zone::LoadGrids()
{
	zone_grids        = GridRepository::GetZoneGrids(GetZoneID());
	zone_grid_entries = GridEntriesRepository::GetZoneGridEntries(GetZoneID());
}
```

Since we have this data in memory now using the structure directly mapped from the repositories, we also have this very nice way over iterating over the data in a vector when it comes to load the data into the grids / waypoints themselves

{% code title="waypoints.cpp" %}
```cpp
void NPC::AssignWaypoints ...
...

for (auto &entry : zone->zone_grid_entries) {
		if (entry.gridid == grid_id) {
			wplist new_waypoint{};
			new_waypoint.index       = max_wp;
			new_waypoint.x           = entry.x;
			new_waypoint.y           = entry.y;
			new_waypoint.z           = entry.z;
			new_waypoint.pause       = entry.pause;
			new_waypoint.heading     = entry.heading;
			new_waypoint.centerpoint = entry.centerpoint;

			LogPathing(
				"Loading Grid [{}] number [{}] name [{}]",
				grid_id,
				entry.number,
				GetCleanName()
			);

			Waypoints.push_back(new_waypoint);
			max_wp++;
		}
	}
```
{% endcode %}
