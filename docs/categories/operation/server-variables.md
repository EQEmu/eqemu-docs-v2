---
description: >-
  This page describes the Variables for your EQEmu Server. These variables are
  found in the variables table.
---

# Server Variables

### Editing Rules

To edit variables, some are tied to admin \#commands that will refresh the variable cache \(noted on description below\), otherwise you must restart world and zone to have them take full effect.

<table>
  <thead>
    <tr>
      <th style="text-align:left">Type</th>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Example Value</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left">String</td>
      <td style="text-align:left">hotfix_name</td>
      <td style="text-align:left"></td>
      <td style="text-align:left">This is used as a temporary cache for #hotfix. It is recommended to not
        touch this unless advised by a developer.</td>
    </tr>
    <tr>
      <td style="text-align:left">Integer</td>
      <td style="text-align:left">loglevel</td>
      <td style="text-align:left">7</td>
      <td style="text-align:left">Related to Commands, Merchants, Trades, Loot. Used for setting verbosity
        of log level for above events. Special rule for &gt;9, not recommended
        to use unless advised by a developer.</td>
    </tr>
    <tr>
      <td style="text-align:left">String</td>
      <td style="text-align:left">MOTD</td>
      <td style="text-align:left"></td>
      <td style="text-align:left">Message of the Day to display on login. Use #motd to set this while in
        game.</td>
    </tr>
    <tr>
      <td style="text-align:left">Boolean</td>
      <td style="text-align:left">new_faction_conversion</td>
      <td style="text-align:left">true</td>
      <td style="text-align:left">Script ran against quests folder to convert new faction values</td>
    </tr>
    <tr>
      <td style="text-align:left">Integer</td>
      <td style="text-align:left">DisableNoRent</td>
      <td style="text-align:left">0</td>
      <td style="text-align:left">when set to 1, shared memory will ignore no rent flag when loading items</td>
    </tr>
    <tr>
      <td style="text-align:left">Integer</td>
      <td style="text-align:left">DisableNoDrop</td>
      <td style="text-align:left">0</td>
      <td style="text-align:left">when set to 1, shared memory will ignore no drop flag when loading items</td>
    </tr>
    <tr>
      <td style="text-align:left">Integer</td>
      <td style="text-align:left">DisableLore</td>
      <td style="text-align:left">0</td>
      <td style="text-align:left">when set to 1, shared memory will ignore lore flag when loading items</td>
    </tr>
    <tr>
      <td style="text-align:left">Integer</td>
      <td style="text-align:left">DisableNoTransfer</td>
      <td style="text-align:left">0</td>
      <td style="text-align:left">when set to 1, shared memory will ignore no transfer flag when loading
        items</td>
    </tr>
    <tr>
      <td style="text-align:left">String</td>
      <td style="text-align:left">GuildCreation</td>
      <td style="text-align:left"></td>
      <td style="text-align:left">Creates a 30 minute founder guild creation system. Appears a delimited
        list of numbers needed. This is unsupported.</td>
    </tr>
    <tr>
      <td style="text-align:left">Integer</td>
      <td style="text-align:left">HonorLSWorldAdmin</td>
      <td style="text-align:left">0</td>
      <td style="text-align:left">when set to 1, CheckAuth() will honor login server world admin. This is
        unlikely to be used in majority of cases.</td>
    </tr>
    <tr>
      <td style="text-align:left">Integer</td>
      <td style="text-align:left">PvPItem</td>
      <td style="text-align:left">0</td>
      <td style="text-align:left">
        <p>if variable ServerType is set to 1, and variable PVPReward is set to 3,
          and this value is &gt;0 and &lt; 200000, then..
          <br />if PvPItem is set to 1, one item can be looted (unconfirmed)</p>
        <p>if PVPItem is set &gt; 1, value is an itemid that will be created on corpse.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Integer</td>
      <td style="text-align:left">PVPReward</td>
      <td style="text-align:left">0</td>
      <td style="text-align:left">
        <p>If set to 1, then one item can be looted (unconfirmed)</p>
        <p>If set to 2, player can loot all items from corpse.</p>
        <p>If set to 3, PVPItem logic is in effect</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">String</td>
      <td style="text-align:left">Rules</td>
      <td style="text-align:left"></td>
      <td style="text-align:left">If rules exists as a record, the value is sent to player as white text,
        freezing player until they accept rules. #serverrules and #acceptrules
        need to be set to a level allowed by a player to enforce this system</td>
    </tr>
    <tr>
      <td style="text-align:left">String</td>
      <td style="text-align:left">RuleSet</td>
      <td style="text-align:left"></td>
      <td style="text-align:left">Sets active ruleset. By default not needed</td>
    </tr>
    <tr>
      <td style="text-align:left">Integer</td>
      <td style="text-align:left">ServerType</td>
      <td style="text-align:left">0</td>
      <td style="text-align:left">Legacy system. If you set to 1, can enable PVP system</td>
    </tr>
  </tbody>
</table>

