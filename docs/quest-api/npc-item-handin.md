## Overview

This update significantly simplifies the NPC item hand-in system by consolidating logic into the source, ensuring consistent behavior across both Perl and Lua. Additionally, it introduces robust item return handling, improves logging to aid in debugging, and adds native support for multi-quest NPCs.

![Image](https://github.com/user-attachments/assets/d7a55891-9ace-4a1a-99e7-e9710f4e2414)

## Key Improvements

### Catch-All System for Returning Items

![Image](https://github.com/user-attachments/assets/09d97498-eb4f-496d-aaf0-c6e98e53c928)

Previously, item loss could occur in various scenarios. The new system mitigates these issues:

- **Empty `EVENT_TRADE` handlers:** Players no longer lose items when handing in to NPCs without a valid event handler.
- **Accidental hand-ins to pets:** Items are returned instead of being lost when mistakenly given to a pet instead of an NPC.
- **Missing `return_items` calls in scripts:** Even if a quest script does not explicitly return items, the system ensures they are not lost.
- **NPCs lacking proper `check_handin` logic:** Items are not consumed without proper validation.
- **Preserving original item instances:** Previously, NPCs returned a new item rather than the exact instance handed in (losing augments, attuned properties, and serial numbers). Now, the exact item instance is returned with all properties intact.
- **Augment loss in Perl scripts:** Previously, augments were lost when using item ID-based summons. This update ensures augments remain intact by returning the original item instance.
- **Disable** Catch-All behavior can be disabled by rule **Items:AlwaysReturnHandins** (Default: true) (Not recommended)

### Item Stacks

- Item stacks can now be processed without being lost!

### Full Item Instance Return

- Items can be handed in with full bags, equipped augments and will be returned in full!

### Unified Handling Across Perl and Lua

- Previously, hand-in logic was separately maintained in both Perl and Lua.
- Now, both scripting languages pass through existing plugins and are handled directly in the source.
- This reduces potential defects, simplifies feature additions, and ensures consistent behavior.

### Improved Player Event Tracking

- Ensures that every item hand-in is captured, even if the quest does not explicitly handle it.
- Provides better player event records for debugging and tracking purposes.

### Consistent Handling of Hand-ins

- Item hand-ins are now processed the same way in both Perl and Lua.
- Money hand-ins follow the same process in both scripting languages.
- NPCs that support multiple simultaneous quests are now handled natively in the source.
- The hand-in logic has been simplified for better maintainability.
- Comprehensive logging allows for deeper debugging of hand-in issues.

### Multi-Quest NPC Support

- A new `multiquest_enabled` column has been added to the `npc_types` database table.
- This allows NPCs to accept multiple quest item combinations natively without requiring external script logic.
- The system ensures the correct quest is matched, and items are not lost due to ambiguous hand-ins.
- NPCs that can accept multiple different item sets will now function correctly.
- If a valid hand-in combination is found, the NPC will proceed as expected.
- If no valid combination is found, all handed-in items are returned.
- The feature is turned on per-NPC either through the `npc_types` field `multiquest_enabled` or by calling **MultiQuestEnable()** via NPC method

<video controls style="width:50%" autoplay>
    <source src="https://github.com/user-attachments/assets/5c65d784-cd82-45a3-9b58-b258bec6f116" type="video/mp4">
</video>

=== "Simple Hand-in"

	=== "Lua"

		```lua
		function event_trade(e)
			if (eq.handin({[1001] = 1})) then -- Cloth Cap (1001)
				e.self:Say("Thank you for the cloth cap!");
			end
		end
		```

	=== "Perl"

		```perl
		sub EVENT_ITEM {
		    if (quest::handin({1001 => 1})) { # Cloth Cap (1001)
		        quest::say("Thank you for the cloth cap!");
		    }
		}
		```

=== "Money"

	=== "Lua"

		```lua
		function event_trade(e)
		    if (eq.handin({platinum = 1000})) then
		        e.self:Say("Thank you for the platinum!");
		    end
		end
		```

	=== "Perl"

		```perl
		sub EVENT_ITEM {
		    if (quest::handin({"platinum" => 1000})) {
		        quest::say("Thank you for the platinum!");
		    }
		}
		```

=== "Money (All)"

	=== "Lua"

		```lua
		function event_trade(e)
		    if (eq.handin({platinum = 1000, gold = 100, silver = 100, copper = 100})) then
		        e.self:Say("Thank you for the platinum, gold, silver and copper!");
		    end
		end
		```

	=== "Perl"

		```perl
		sub EVENT_ITEM {
		    if (quest::handin({"platinum" => 1000, "gold" => 100, "silver" => 100, "copper" => 100})) {
		        quest::say("Thank you for the platinum, gold, silver and copper!");
		    }
		}
		```

=== "Item + Money"

	=== "Lua"

		```lua
		function event_trade(e)
		    if (eq.handin({[1001] = 1, platinum = 1000})) then -- Cloth Cap (1001)
		        e.self:Say("Thank you for the cloth cap and platinum!");
		    end
		end
		```

	=== "Perl"

		```perl
		sub EVENT_ITEM {
		    if (quest::handin({1001 => 1, "platinum" => 1000})) { # Cloth Cap (1001)
		        quest::say("Thank you for the cloth cap and platinum!");
		    }
		}
		```

=== "Item Table"

	=== "Lua"

		```lua
		function event_trade(e)
		    local handin_requirements = {
		        { 
		            required_handin = { [1001] = 1, [13005] = 4 }, -- Cloth cap, Iron ration (4)
		            exp = 100,
		            result_item = 2300, -- Journeyman's Boots
		            message = "Thank you!"
		        },
		        { 
		            required_handin = { platinum = 10000 },
		            exp = 100,
		            result_item = 13401, -- Manastone
		            message = "Thank you!"
		        }
		    }

		    for _, handin in ipairs(handin_requirements) do
		        if eq.handin(handin.required_handin) then
		            if handin.message then e.self:Say(handin.message) end
		            if handin.result_item then e.other:SummonItem(handin.result_item) end
		            if handin.exp then e.other:AddEXP(handin.exp) end
		            return
		        end
		    end
		end
		```

	=== "Perl"

		```perl
		sub EVENT_ITEM {
		  my @handin_requirements = (
		    { 
		      required_handin => { 1001 => 1, 13005 => 4 }, # cloth cap, iron ration (4)
		      exp => 100,
		      result_item => 2300, # Journeyman's Boots
		      message => "Thank you!"
		    },
		    { 
		      required_handin => { "platinum" => 10000 },
		      exp => 100,
		      result_item => 13401, # Manastone
		      message => "Thank you!"
		    }
		  );

		  foreach my $handin (@handin_requirements) {
		    if (quest::handin($handin->{required_handin})) {
		      quest::say($handin->{message}) if $handin->{message};
		      quest::summonitem($handin->{result_item}) if $handin->{result_item};
		      quest::exp($handin->{exp}) if $handin->{exp};
		      return;
		    }
		  }
		}
		```

=== "Item Stacks"

	=== "Lua"

		```lua
		function event_trade(e)
		    if (eq.handin({[13005] = 20})) then -- Iron Rations (13005) (20 count)
		        e.self:Say("Thank you for all of the rations! I was SO hungry.");
		    end
		end
		```

	=== "Perl"

		```perl
		sub EVENT_ITEM {
			if (quest::handin({13005 => 20})) { # Iron Rations (13005) (20 count)
		    	quest::say("Thank you for all of the rations! I was SO hungry.");
		  	}
		}
		```

=== "Item Stacks (Looped)"

	=== "Lua"

		```lua
		function event_trade(e)
			local count = 0
		    while (eq.handin({[13005] = 1})) do
		        count = count + 1
		    end
			if count > 0 then
		        e.self:Say("Thank you for the " .. count .. " rations! We could feed a village!")
		    end
		end
		```

	=== "Perl"

		```perl
		sub EVENT_ITEM {
		  my $count = 0;
		  while (quest::handin({13005 => 1})) {
		    $count++;
		  }
		  if ($count > 0) {
		    quest::say("Thank you for the $count rations! We could feed a village!");
		  }
		}
		```

	!!! info 

		Scenario 1) Player hands in a stack of 1,000 iron rations, they get processed one for each in the stack handed in

		NPC says, 'Thank you for the 1000 rations! We could feed a village!'

	!!! info

		Scenario 2) Player hands in multiple different stacks equalling 1,055 rations.

		![Image](https://github.com/user-attachments/assets/c65ab72b-add1-4140-b96f-50ff6559ebf3)

		NPC says, 'Thank you for the 1055 rations! We could feed a village!'

=== "Multi-Questing"

	=== "Lua"

		```lua
		function event_trade(e)
			e.self:MultiQuestEnable(); 
			if (eq.handin({[12268] = 1, [7100] = 1, platinum = 325})) then
				e.self:Say("The time to trade has come!! I am now rich and you are now fast. Take the Journeyman Boots and run like the wind.");
				e.other:QuestReward(e.self, 0, 0, 0, 0, 2300, 1250);
			end
		end
		```

	=== "Perl"

		```perl
		sub EVENT_ITEM {
		    $npc->MultiQuestEnable();
		    if (quest::handin({12268 => 1, 7100 => 1, "platinum" => 325})) {
		        quest::say("The time to trade has come!! I am now rich and you are now fast. Take the Journeyman Boots and run like the wind.");
		        quest::summonitem(2300);
		        quest::exp(1250);
		    }
		}
		```

## Hand-In Examples

[ ] TODO: Add examples @Akkadius

## Logging Enhancements

- A new logging category `Logs::NpcHandin` has been introduced.
- This captures all item hand-ins, making it easier to debug and track player transactions.
- Logs can be output to both the console and GM chat for real-time monitoring.

## Related Plugin Changes

This update includes corresponding changes to quest plugins. See the associated plugin update: [ProjectEQ Quest Plugin Changes #1403](https://github.com/ProjectEQ/projecteqquests/pull/1403)

## Conclusion

This overhaul simplifies item hand-ins, unifies handling across scripting languages, prevents item loss, improves logging for debugging, and adds multi-quest support. Users and developers can expect a more stable and predictable experience when handling items in quests.

