## Overview

This update significantly simplifies the NPC item hand-in system by consolidating logic into the source, ensuring consistent behavior across both Perl and Lua. Additionally, it introduces robust item return handling, improves logging to aid in debugging, and adds native support for multi-quest NPCs.

![Image](https://github.com/user-attachments/assets/d7a55891-9ace-4a1a-99e7-e9710f4e2414)

## Key Improvements

## Catch-All System for Returning Items

![Image](https://github.com/user-attachments/assets/09d97498-eb4f-496d-aaf0-c6e98e53c928)

Previously, item loss could occur in various scenarios. The new system mitigates these issues:

- **Empty `EVENT_TRADE` handlers:** Players no longer lose items when handing in to NPCs without a valid event handler.
- **Accidental hand-ins to pets:** Items are returned instead of being lost when mistakenly given to a pet instead of an NPC.
- **Missing `return_items` calls in scripts:** Even if a quest script does not explicitly return items, the system ensures they are not lost.
- **NPCs lacking proper `check_handin` logic:** Items are not consumed without proper validation.
- **Preserving original item instances:** Previously, NPCs returned a new item rather than the exact instance handed in (losing augments, attuned properties, and serial numbers). Now, the exact item instance is returned with all properties intact.
- **Augment loss in Perl scripts:** Previously, augments were lost when using item ID-based summons. This update ensures augments remain intact by returning the original item instance.
- **Disable** Catch-All behavior can be disabled by rule **Items:AlwaysReturnHandins** (Default: true) (Not recommended)

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

**Hasten Bootstrutter Example**

=== "Lua"

	```lua
	function event_trade(e)
		local item_lib = require("items");
		e.self:MultiQuestEnable(); 
		if (item_lib.check_turn_in(e.trade, {item1 = 12268, item2 = 7100, platinum = 325})) then
			e.self:Say("The time to trade has come!! I am now rich and you are now fast. Take the Journeyman Boots and run like the wind.");
			e.other:QuestReward(e.self, 0, 0, 0, 0, 2300, 1250);
		end
		item_lib.return_items(e.self, e.other, e.trade)
	end
	```

=== "Perl"

	```perl
	sub EVENT_ITEM {
	    $npc->MultiQuestEnable();
	    if (plugin::check_handin(\%itemcount, 12268 => 1, 7100 => 1, "platinum" => 325)) {
	        quest::say("The time to trade has come!! I am now rich and you are now fast. Take the Journeyman Boots and run like the wind.");
	        quest::summonitem(2300);
	        quest::exp(1250);
	    }
	    plugin::return_items(\%itemcount);
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

