---
description: EQEmu Grid Wander Types
---

# Wander Types

{% hint style="info" %}
These Wander Types are assigned to a [Grid](https://eqemu.gitbook.io/database-schema/categories/grids) to cause the behavior listed below as the NPC wanders through Grid Waypoint Entries.
{% endhint %}

<table>
  <thead>
    <tr>
      <th style="text-align:left"><b>ID</b>
      </th>
      <th style="text-align:left">Wander Type Name</th>
      <th style="text-align:left"><b>Wander Type Description</b>
      </th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left">0</td>
      <td style="text-align:left">GridCircular</td>
      <td style="text-align:left">Circle the points in order (IE 1 through 10); start back at first.</td>
    </tr>
    <tr>
      <td style="text-align:left">1</td>
      <td style="text-align:left">GridRandom10</td>
      <td style="text-align:left">Pick 10 random points from the grid set and cycle through.</td>
    </tr>
    <tr>
      <td style="text-align:left">2</td>
      <td style="text-align:left">GridRandom</td>
      <td style="text-align:left">Random; every point is randomly selected from the entire grid set.</td>
    </tr>
    <tr>
      <td style="text-align:left">3</td>
      <td style="text-align:left">GridPatrol</td>
      <td style="text-align:left">Patrol; walk way points in order (i.e. 1 through 10) and then reverse
        order (i.e. 10 through 1).</td>
    </tr>
    <tr>
      <td style="text-align:left">4</td>
      <td style="text-align:left">GridOneWayRepop</td>
      <td style="text-align:left">Go to the end and depop with spawn timer; walk waypoints in order (i.e.
        1 through 10), depop at final point, and repop at initial point after spawn
        timer.</td>
    </tr>
    <tr>
      <td style="text-align:left">5</td>
      <td style="text-align:left">GridRand5LoS</td>
      <td style="text-align:left">Pick random closest 5 and pick one that&apos;s in line of sight.</td>
    </tr>
    <tr>
      <td style="text-align:left">6</td>
      <td style="text-align:left">GridOneWayDepop</td>
      <td style="text-align:left">Go the end and depop without spawn timer; walk waypoints in order (i.e.
        1 through 10) and then depop.</td>
    </tr>
    <tr>
      <td style="text-align:left">7</td>
      <td style="text-align:left">GridCenterPoint</td>
      <td style="text-align:left">Initial point as the center, fan out to each grid point creating a &quot;star
        burst&quot; like pattern (1 - 2 - 1 - 4 - 1 - 3 - 1 - 5, etc.).</td>
    </tr>
    <tr>
      <td style="text-align:left">8</td>
      <td style="text-align:left">GridRandomCenterPoint</td>
      <td style="text-align:left">
        <p>Initial point as the center, fan out to a random grid point.</p>
        <p></p>
        <p>Causes an NPC to alternate between a random waypoint in grid_entries and
          a random waypoint marked with the center point column set to true. If no
          waypoints are marked as a center point, this wander type will not work.
          There is no numbering requirement or limit for center points--you can have
          as many as you need.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">9</td>
      <td style="text-align:left">GridRandomPath</td>
      <td style="text-align:left">Randomly select a waypoint but follow path to it instead of walk directly
        to it ignoring walls.</td>
    </tr>
  </tbody>
</table>

