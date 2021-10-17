# Adjust Maximum Level

So you want to take your server to the next level? There are a lot of old posts about this, and the ones Google likes to show are all missing a crucial detail. There are four variables that contain this data. They are all in the rule\_values table. I don't see a page for this, so here it goes.

_In this example I'm setting the Max Level to 100; change this value as needed._

**1\)** Open a MySQL Admin tool \(Heidi, OmniDB, MySQL Workbench, etc.\) and connect to the eqemu database for your server.

**2\)** Run the following SQL Statements:

`UPDATE eqemu.rule_values SET rule_value = 100 WHERE rule_name = "Character:MaxLevel"`

`UPDATE eqemu.rule_values SET value = 100 WHERE rule_name = "Character:MaxExpLevel"`

**3\)** Verify

`SELECT * FROM rule_values WHERE rule_name = "Character:MaxLevel" OR rule_name = "Character:MaxExpLevel"`

You should have 4 records returned. If everything looks correct, proceed to step 4.

**4\)** Restart your server.

_Thanks to Akkadius for the clue on multiple variables._

