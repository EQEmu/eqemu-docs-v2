# Example

This example demonstrates a full shared task selector offer with dynamic zone creation on accept. Note that dz creation can occur at any time and does not have to occur on the accept event but will likely be the most common point.

```lua
local shared_task_id = 4795 -- some database task id with type set to shared task
local non_shared_task_id = 10 -- some database task id with type set to quest

function event_say(e)
  if e.message:findi("hail") then
    -- The non-shared task here will be ignored, mixing shared tasks
    -- with other types is not allowed due to shared task validation
    eq.task_selector({ non_shared_task_id, shared_task_id })
  end
end

function event_task_accepted(e)
  if e.task_id == shared_task_id then
    -- create dz for the shared task on the client that accepted it
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

