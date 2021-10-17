# Shared Tasks

Shared tasks use the existing task system quest API with the added availability of dynamic zone instance creation. The only major difference where care needs to be taken in scripts is with player task update events which will trigger for all members of a shared task. Any script that depends on a single task update or stage completed event to perform an action will need to guard against multiple updates from members.

Note that mixing shared task types with other types is not supported via the [taskselector](https://docs.eqemu.io/server/categories/task-system-guide#taskselector) APIs since shared tasks need to run through separate validation for player count and member level requirements. If any shared task type is included in a task selector offer, all other task types will be dropped.

The [EVENT\_TASK ACCEPTED](https://docs.eqemu.io/server/categories/task-system-guide#event_taskaccepted) event will only occur if shared task creation was successful after passing all validation requirements defined in the `tasks` database entry.

## Dynamic Zones

Dynamic zones are the same underlying instancing system that expedition's use. These may be created for shared tasks through the client API method:

```lua
Client::CreateTaskDynamicZone(int taskid, LUA_TTABLE dynamic_zone_info)
```

{% hint style="warning" %}
This method is currently only available to the Lua API
{% endhint %}

This method creates a dynamic zone for a client assigned to the specified task id. If the client does not have a task for the task id or the specified task id is not a shared task, then the dynamic zone is not created. If dynamic zone creation is successful then all members of the client's shared task are added to it. Multiple dynamic zones may be created for a shared task to support custom servers.

Dynamic zone creation data is passed through the `dynamic_zone_info` parameter similar to the [Client::CreateExpedition](https://docs.eqemu.io/server/categories/expedition-system/quest-api-lua/client-methods#expedition-createexpedition-lua_ttable-expedition_info) API. Passing a duration for the instance data will have no effect. The shared task system will automatically override the duration to the shared task's remaining time.

{% hint style="warning" %}
A shared task set to `0` duration in the database for unlimited time is not supported and the dynamic zone will be set to a max duration of 24 hours \(suspected max of live shared tasks\)
{% endhint %}

| key name | supported keys | description |
| :--- | :--- | :--- |
| `instance` | zone, version | dz instance details |
| `compass` | zone, x, y, z | `(optional)` compass coordinates |
| `safereturn` | zone, x, y, z, h | `(optional)` safe return coordinates |
| `zonein` | x, y, z, h | `(optional)` dz zone-in coordinates |

#### instance keys <a id="instance-keys"></a>

| key | type | description |
| :--- | :--- | :--- |
| `zone` | int or string | zone id or zone short\_name of the expedition's dz instance |
| `version` | int | dz instance version |
| `duration` | int | \(optional\) no effect |

#### compass, safereturn, and zonein keys <a id="compass-safereturn-and-zonein-keys"></a>

| key | type | description |
| :--- | :--- | :--- |
| `zone` | int or string | zone id or zone short name |
| `x` | float | `(optional)` x coordinate \(default: `0`\) |
| `y` | float | `(optional)` y coordinate \(default: `0`\) |
| `z` | float | `(optional)` z coordinate \(default: `0`\) |
| `h` | float | `(optional)` heading \(default: `0`\) |

```lua
local shared_task_id = 4795 -- some database task id with type set to shared task

function event_task_accepted(e)
  -- create dz for the accepted task on the client that accepted it
  if e.task_id == shared_task_id then
    -- dz duration will be overridden by time remaining on the shared task
    local dz = {
      instance   = { zone="thundercrest", version = 11 },
      compass    = { zone="broodlands", x=1241.88, y=511.147, z=23.4192 },
      safereturn = { zone="broodlands", x=1242.0, y=526.0, z=27.0, h=0.0 }
    }
    e.other:CreateTaskDynamicZone(e.task_id, dz)
  end
end
```

