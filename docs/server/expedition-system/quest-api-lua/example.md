# Example

This is a full example to demonstrate usage. Guard Pineshade at the orc lift in gfaydark provides an expedition to crushbone

`gfaydark/Guard_Pineshade.lua`

```lua
local expedition_name = "Crushbone, DVinn's Stronghold"

local expedition_info = {
  expedition = { name=expedition_name, min_players=1, max_players=6 },
  instance   = { zone="crushbone", version=0, duration=3600 },
  compass    = { zone="gfaydark", x=238.00, y=987.00, z=-24.90 }, -- pointing to guard pineshade
  safereturn = { zone="gfaydark", x=245.84, y=987.93, z=-27.6, h=484.0 }, -- at orc lift in gfay
  zonein     = { x=479.44, y=-500.18, z=5.75, h=421.8 } -- on bridge
}

function event_say(e)
  if e.message:findi("hail") then
    local dz = e.other:GetExpedition()
    if dz.valid and dz:GetName() == expedition_name then
      e.self:Say("Tell me when you're [" .. eq.say_link("ready") .. "] to enter")
    else
      e.self:Say("You may [" .. eq.say_link("request") .. "] the expedition")
    end
  elseif e.message:findi("request") then
    local dz = e.other:CreateExpedition(expedition_info) -- request expedition
    if dz.valid then
      dz:AddReplayLockout(7200) -- immediately add a 2 hour replay lockout on creation
      e.self:Say("Tell me when you're [" .. eq.say_link("ready") .. "] to enter")
    end
  elseif e.message:findi("ready") then
    e.other:MovePCDynamicZone("crushbone")
  end
end
```

`crushbone/script_init.lua`

```lua
local dz = eq.get_expedition()
if dz.valid then
  -- bind unique npc types to events to prevent any looting exploits
  dz:SetLootEventByNPCTypeID(58032, "Emperor Crush") -- npc: Emperor_Crush
  dz:SetLootEventByNPCTypeID(317109, "Ambassador Mata Muram") -- npc: Overlord_Mata_Muram

  -- spawn any events that the expedition doesn't have a lockout for
  -- (can't depop npcs that spawn by default from this script, handle elsewhere)
  if not dz:HasLockout("Ambassador Mata Muram") then
    eq.spawn2(317109, 0, 0, 130.78, -149.09, 88.70, 270.8) -- npc: Overlord_Mata_Muram
  end
end
```

`crushbone/Emperor_Crush.lua`

```lua
function event_spawn(e)
  local dz = eq.get_expedition()
  if dz.valid and dz:HasLockout("Emperor Crush") then
    eq.depop()
  end
end

function event_death(e)
  local dz = eq.get_expedition()
  if dz.valid then
    dz:AddLockout("Emperor Crush", 86400) -- 1 day lockout
  end
end
```

`crushbone/Overlord_Mata_Muram.lua`

```lua
function event_death(e)
  local dz = eq.get_expedition()
  if dz.valid then
    dz:AddLockout("Ambassador Mata Muram", 604800) -- 7 day lockout

    -- spawn a chest and bind its spawn id to the event to prevent loot exploits
    local chest = eq.spawn2(893, 0, 0, e.self:GetX(), e.self:GetY(), e.self:GetZ(), e.self:GetHeading())
    dz:SetLootEventBySpawnID(chest:GetID(), "Ambassador Mata Muram")
  end
end
```

