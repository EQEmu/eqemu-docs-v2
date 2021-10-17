---
description: >-
  This page describes the conversion of quest globals (qglobals) into Data
  Buckets.
---

# Converting QGlobals

## Changing Quest Globals to Data Buckets

As many may know, Quest Globals are stored using a name, type, value, and duration.  Due to this we can do a 1:1 conversion of all 8 types of globals.

### Types of Globals and their specificities:

| Type | NPC ID | Player | Zone |
| :--- | :--- | :--- | :--- |
| 0 | Current | Current | Current |
| 1 | All | Current | Current |
| 2 | Current | All | Current |
| 3 | All | All | Current |
| 4 | Current | Current | All |
| 5 | All | Current | All |
| 6 | Current | All | All |
| 7 | All | All | All |

### Shortcuts for Globals

This list is intended to make it easier for you to assign globals to the desired entity type.  

| Shortcut | Value |
| :--- | :--- |
| All NPCs | 1 |
| All Players | 2 |
| All NPCs and Players | 3 |
| All Zones | 4 |
| All NPCs and Zones | 5 |
| All Players and Zones | 6 |
| All NPCs, Players and Zones | 7 |

### Conversion

The way we will be storing this data is by using the bucket's name.  As you can see, types 0 through six have similar names, based on NPC/Player/Zone specificity. Type 7 is global, so the bucket name is just the qglobal name with no specificity.

**Conversion SQL:**

```sql
-- Type 0 (NPC, Player, and Zone Specific):
INSERT INTO
  data_buckets (`key`, `value`, `expires`)
SELECT
  CONCAT(npcid, '-', charid, '-', zoneid, '-', name) AS `name`,
  `value`,
  `expdate`
FROM
  quest_globals
WHERE
  zoneid > 0 AND npcid > 0 AND charid > 0;
 
 
-- Type 1 (Player and Zone Specific):
INSERT INTO
  data_buckets (`key`, `value`, `expires`)
SELECT
  CONCAT('0-', charid, '-', zoneid, '-', name) AS `name`,
  `value`,
  `expdate`
FROM
  quest_globals
WHERE
  zoneid > 0 AND npcid = 0 AND charid > 0;
 
 
-- Type 2 (NPC and Zone Specific):
INSERT INTO
  data_buckets (`key`, `value`, `expires`)
SELECT
  CONCAT(npcid, '-0-', zoneid, '-', name) AS `name`,
  `value`,
  `expdate`
FROM
  quest_globals
WHERE
  zoneid > 0 AND npcid > 0 AND charid = 0;
 
 
-- Type 3 (Zone Specific):
INSERT INTO
  data_buckets (`key`, `value`, `expires`)
SELECT
  CONCAT('0-0-', zoneid, '-', name) AS `name`,
  `value`,
  `expdate`
FROM
  quest_globals
WHERE
  zoneid > 0 AND npcid = 0 AND charid = 0;
 
 
-- Type 4 (NPC and Player Specific):
INSERT INTO
  data_buckets (`key`, `value`, `expires`)
SELECT
  CONCAT(npcid, '-', charid, '-0-', name) AS `name`,
  `value`,
  `expdate`
FROM
  quest_globals
WHERE
  zoneid = 0 AND npcid > 0 AND charid > 0;
 
 
-- Type 5 (Character Specific):
INSERT INTO
  data_buckets (`key`, `value`, `expires`)
SELECT
  CONCAT('0-', charid, '-0-', name) AS `name`,
  `value`,
  `expdate`
FROM
  quest_globals
WHERE
  zoneid = 0 AND npcid = 0 AND charid > 0;
 
 
-- Type 6 (NPC Specific):
INSERT INTO
  data_buckets (`key`, `value`, `expires`)
SELECT
  CONCAT(npcid, '-0-0-', name) AS `name`,
  `value`,
  `expdate`
FROM
  quest_globals
WHERE
  zoneid = 0 AND npcid > 0 AND charid = 0;
 
 
-- Type 7 (Global):
INSERT INTO
  data_buckets (`key`, `value`, `expires`)
SELECT
  `name`,
  `value`,
  `expdate`
FROM
  quest_globals
WHERE
  zoneid = 0 AND npcid = 0 AND charid = 0;
```

### Requirements

#### Lua:

The Lua module necessary for using the Lua script \(lua\_modules/buckets.lua\):

{% tabs %}
{% tab title="buckets.lua" %}
```lua

local data_buckets = {}
 
function data_buckets.SetData(npc, client, name, value, options, duration)
    local npcid = 0;
    local charid = 0;
    local zoneid = eq.get_zone_id();
    local bucket_name = "";
    if (options > 7) then -- If somehow you put options above 7, it defaults to type 7
        options = 7;
    end
    if (options == 0) then
        if (npc ~= nil) then
            npcid = npc:CastToNPC():GetNPCTypeID();
        end
       
        if (client ~= nil) then
            charid = client:CastToClient():CharacterID();
        end
        bucket_name = npcid .. "-" .. charid .. "-" .. zoneid .. "-" .. name;
    elseif (options > 0 and options < 7) then
        if (options == 1 or options == 3 or options == 5) then
            if (client ~= nil) then
                charid = client:CharacterID();
            end
        end
       
        if (options == 2 or options == 3 or options == 6) then
            if (npc ~= nil) then
                npcid = npc:GetNPCTypeID();
            end
        end
       
        if (options > 3) then
            zoneid = 0;
        end
        bucket_name = npcid .. "-" .. charid .. "-" .. zoneid .. "-" .. name;
    elseif (options == 7) then
        bucket_name = name;
    end
    if (duration ~= nil) then
        eq.set_data(bucket_name, tostring(value), tostring(duration));
    else
        eq.set_data(bucket_name, tostring(value));
    end
end
 
function data_buckets.GetData(npc, client, name, options)
    local npcid = 0;
    local zoneid = eq.get_zone_id();
    local charid = 0;
    local bucket_name = "";
    if (options < 0 or options > 7) then
        options = 7;
    end
   
    if (options == 0) then
        if (npc ~= nil) then
            npcid = npc:GetNPCTypeID();
        end
       
        if (client ~= nil) then
            charid = client:CharacterID();
        end
        bucket_name = npcid .. "-" .. charid .. "-" .. zoneid .. "-" .. name;
    elseif (options > 0 and options < 7) then
        if (options == 1 or options == 3 or options == 5) then
            if (client ~= nil) then
                charid = client:CharacterID();
            end
        end
       
        if (options == 2 or options == 3 or options == 6) then
            if (npc ~= nil) then
                npcid = npc:GetNPCTypeID();
            end
        end
       
        if (options > 3) then
            zoneid = 0;
        end
        bucket_name = npcid .. "-" .. charid .. "-" .. zoneid .. "-" .. name;
    elseif (options == 7) then
        bucket_name = name;
    end
    return eq.get_data(bucket_name);
end
 
function data_buckets.DeleteData(npc, client, name, options)
    local npcid = 0;
    local zoneid = eq.get_zone_id();
    local charid = 0;
    if (options < 0 or options > 7) then
        options = 7;
    end
   
    if (options == 0) then
        if (npc ~= nil) then
            npcid = npc:GetNPCTypeID();
        end
       
        if (client ~= nil) then
            charid = client:CharacterID();
        end
        bucket_name = npcid .. "-" .. charid .. "-" .. zoneid .. "-" .. name;
    elseif (options > 0 and options < 7) then
        if (options == 1 or options == 3 or options == 5) then
            if (client ~= nil) then
                charid = client:CharacterID();
            end
        end
       
        if (options == 2 or options == 3 or options == 6) then
            if (npc ~= nil) then
                npcid = npc:GetNPCTypeID();
            end
        end
       
        if (options > 3) then
            zoneid = 0;
        end
        bucket_name = npcid .. "-" .. charid .. "-" .. zoneid .. "-" .. name;
    elseif (options == 7) then
        bucket_name = name;
    end
    eq.delete_data(bucket_name);
end
 
return data_buckets;
```
{% endtab %}
{% endtabs %}

**Example**

```lua
function event_say(e)
    local data_buckets = require("data_buckets");
    if (e.message:findi("set")) then
        for options = 0, 7, 1 do
            data_buckets.SetData(e.self, e.other, "Options Test " .. options, options, options);
        end
    elseif (e.message:findi("read")) then
        for options = 0, 7, 1 do
            e.self:Say(data_buckets.GetData(e.self, e.other, "Options Test " .. options, options));
        end
    elseif (e.message:findi("delete")) then
        for options = 0, 7, 1 do
            data_buckets.DeleteData(e.self, e.other, "Options Test " .. options, options);
        end
    end
end
```

#### Perl

The Perl plugin necessary for using the Perl script \(plugins/buckets.pl\):

{% tabs %}
{% tab title="buckets.pl" %}
```perl
#Type    ID      Player      Zone
#0      Current     Current     Current
#1      All         Current     Current
#2      Current     All         Current
#3      All         All         Current
#4      Current     Current     All
#5      All         Current     All
#6      Current     All         All
#7      All         All         All
 
sub SetData {
    my @data = @_;
    my ($npcid, $charid, $zoneid) = (0, 0, plugin::val('zoneid'));
    my ($name, $value, $type, $duration) = ($data[0], $data[1], $data[2], (defined $data[3] ? $data[3] : 0));
    if ($type == 0) {
        my $npc = plugin::val('npc');
        $npcid = $npc->GetNPCTypeID();
        $charid = plugin::val('charid');
        if ($duration > 0) {
            quest::set_data("$npcid-$charid-$zoneid-$name", $value, $duration);
        } else {
            quest::set_data("$npcid-$charid-$zoneid-$name", $value);
        }
    } elsif ($type ~~ [1..6]) {
        if ($type & 1) {
            $charid = plugin::val('charid');
        }
        if ($type & 2) {
            my $npc = plugin::val('npc');
            $npcid = $npc->GetNPCTypeID();
        }
        if ($type & 4) {
            $zoneid = 0;
        }
        if ($duration > 0) {
            quest::set_data("$npcid-$charid-$zoneid-$name", $value, $duration);
        } else {
            quest::set_data("$npcid-$charid-$zoneid-$name", $value);
        }
    } elsif ($type == 7) {
        if ($duration > 0) {
            quest::set_data("$name", $value, $duration);
        } else {
            quest::set_data("$name", $value);
        }
    }
}
 
sub GetData {
    my ($npcid, $charid, $zoneid) = (0, 0, plugin::val('zoneid'));
    my ($name, $type) = (shift, shift);
    if ($type == 0) {
        my $npc = plugin::val('npc');
        $npcid = $npc->GetNPCTypeID();
        $charid = plugin::val('charid');
        plugin::Message("$npcid-$charid-$zoneid-$name: " . quest::get_data("$npcid-$charid-$zoneid-$name"));
        return quest::get_data("$npcid-$charid-$zoneid-$name");
    } elsif ($type ~~ [1..6]) {
        if ($type & 1) {
            $charid = plugin::val('charid');
        }
        if ($type & 2) {
            my $npc = plugin::val('npc');
            $npcid = $npc->GetNPCTypeID();
        }
        if ($type & 4) {
            $zoneid = 0;
        }
        plugin::Message("$npcid-$charid-$zoneid-$name: " . quest::get_data("$npcid-$charid-$zoneid-$name"));
        return quest::get_data("$npcid-$charid-$zoneid-$name");    
    } elsif ($type == 7) {
        return quest::get_data("$name");
    }
}
 
sub DeleteData {
    my ($npcid, $charid, $zoneid) = (0, 0, plugin::val('zoneid'));
    my ($name, $type) = (shift, shift);
    if ($type == 0) {
        my $npc = plugin::val('npc');
        $npcid = $npc->GetNPCTypeID();
        $charid = plugin::val('charid');
        quest::delete_data("$npcid-$charid-$zoneid-$name");
    } elsif ($type ~~ [1..6]) {
        if ($type & 1) {
            $charid = plugin::val('charid');
        }
        if ($type & 2) {
            my $npc = plugin::val('npc');
            $npcid = $npc->GetNPCTypeID();
        }
        if ($type & 4) {
            $zoneid = 0;
        }
        quest::delete_data("$npcid-$charid-$zoneid-$name");    
    } elsif ($type == 7) {
        quest::delete_data("$name");
    }
}
```
{% endtab %}
{% endtabs %}

**Example:**

```perl
sub EVENT_SAY {
    if ($text=~/Set/i) {
        for (my $type = 0; $type < 8; $type++) {
            plugin::SetData("Type Test $type", "Blah blah $type", $type);
        }
    } elsif ($text=~/Read/i) {
        for (my $type = 0; $type < 8; $type++) {
            plugin::GetData("Type Test $type", $type);
        }
    } elsif ($text=~/Delete/i) {
        for (my $type = 0; $type < 8; $type++) {
            plugin::DeleteData("Type Test $type", $type);
        }
    }
}
```

