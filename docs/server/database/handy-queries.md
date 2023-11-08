# Handy Queries

The purpose of this document is to provide a list of useful queries that can be used to help with various tasks.

If you think a query could be helpful and added to this list, suggest it in EQEmu server help.

## Find All Player Accounts By Player Name

```sql
set @player_name = 'Player';

SELECT
    *
FROM
    account
WHERE
    id = (select account_id from character_data where name = @player_name);
```

## Find All Accounts by IP tied to Character Name

```sql
set @player_name = 'Player';

SELECT
    account_ip.accid,
    account_ip.ip,
    account.name,
    account.charname,
    account_ip.lastused
FROM
    account_ip
    INNER JOIN account ON account_ip.accid = account.id -- Join account table to get account data
    WHERE
        account_ip.ip IN ( -- Get all IP's via accid 
            SELECT
                account_ip.ip
            FROM
                account_ip
            WHERE
                account_ip.accid = ( -- Get account_id by player name
                    SELECT
                        character_data.account_id
                    FROM
                        character_data
                    WHERE
                        name = @player_name))
            GROUP BY
                account.id
            ORDER BY
                account_ip.lastused DESC;
```

## Find All Accounts by IP tied to Character Name Concat Result

```sql
set @player_name = 'Player';

SELECT
    DISTINCT GROUP_CONCAT(account.id)
FROM
    account_ip
        INNER JOIN account ON account_ip.accid = account.id -- Join account table to get account data
WHERE
        account_ip.ip IN ( -- Get all IP's via accid 
        SELECT
            account_ip.ip
        FROM
            account_ip
        WHERE
                account_ip.accid = ( -- Get account_id by player name
                SELECT
                    character_data.account_id
                FROM
                    character_data
                WHERE
                    name = @player_name))
ORDER BY
    account_ip.lastused DESC;
```

## Find All Spells for Sale

The following query looks for all spells that are for sale by NPCs.  This is useful for finding spells that are not currently for sale, or for finding spells that are for sale within an era.

```sql
SELECT
    i.id,
    i.Name,
    m.merchantid,
    m.min_expansion as merchant_min,
    m.max_expansion as merchant_max,
    st.max_expansion as spawn2_max,
    st.min_expansion as spawn2_min,
    st.enabled as spawn2_enabled,
    st.zone,
    n.name,
    n.lastname
FROM
    items i,
    merchantlist m,
    npc_types n,
    spawnentry se,
    spawn2 st
WHERE
    m.merchantid = n.id AND
    se.npcID = n.id AND
    se.spawngroupID = st.spawngroupID AND
    m.item = i.id AND
    st.enabled = '1' AND
    se.chance > '0' AND
    st.version = '0' AND
    n.id < '347000' AND
    (m.max_expansion >= '9' or m.max_expansion = '-1') AND
    (m.min_expansion <= '9' or m.min_expansion = '-1') AND
    (st.max_expansion >= '9' or st.max_expansion = '-1') AND
    (st.min_expansion <= '9' or st.min_expansion = '-1') AND
    (i.Name LIKE 'Song:%' OR i.Name LIKE 'Tome of %' OR i.Name LIKE 'Spell: %');
```

## Find Spells that Teleport to Zones with x Expansion

Below example looks for spells that are greater than expansion 8 (Omens of War) that teleport to zones that are greater than expansion 8 (Omens of War).

```sql
SELECT
    id,
    teleport_zone,
    `name`,
    (( SELECT expansion FROM zone WHERE teleport_zone = short_name LIMIT 1 ) - 1 ) AS expansion 
FROM
    spells_new 
WHERE
    teleport_zone = ( SELECT short_name FROM zone WHERE teleport_zone = short_name LIMIT 1 ) 
    AND (( SELECT expansion FROM zone WHERE teleport_zone = short_name LIMIT 1 ) - 1 ) > 8;
```

## Find Spells that Teleport to Zones with x Expansion (item, zone, npc assoc)

```sql
SELECT
    merchantid,
    item,
    min_expansion,
    (select expansion - 1 from zone where short_name = (select teleport_zone from spells_new where id = (select scrolleffect from items where id = item)) limit 1) as proposed_expansion,
    (select `name` from items where id = item) as item_name,
    ( SELECT NAME FROM npc_types WHERE merchant_id = merchantid LIMIT 1 ) AS npc,
    (
    SELECT
        zone 
    FROM
        spawn2 
    WHERE
        spawngroupID = ( SELECT spawngroupID FROM spawnentry WHERE npcID = ( SELECT id FROM npc_types WHERE merchant_id = merchantid LIMIT 1 ) LIMIT 1 ) 
    ) AS zone 
FROM
    merchantlist 
WHERE
    item IN (
    SELECT
        id 
    FROM
        items 
    WHERE
        scrolleffect IN (
        SELECT
            id 
        FROM
            spells_new 
        WHERE
            teleport_zone = ( SELECT short_name FROM zone WHERE teleport_zone = short_name LIMIT 1 ) 
            AND (( SELECT expansion FROM zone WHERE teleport_zone = short_name LIMIT 1 ) - 1 ) > 8 
        ) 
    )
    ORDER BY min_expansion;
```

## NPC's by Zone and Version

```sql
select
  *
from
  npc_types
where
  id IN (
    select
      spawnentry.npcID
    from
      spawnentry
      join spawn2 on spawn2.spawngroupID = spawnentry.spawngroupID
    where
      spawn2.zone = 'crushbone'
      and spawn2.version = 0
  );
```

## Online Characters

Filters by characters saved within the past 10 minutes

```sql
SELECT
    character_data.account_id,
    character_data.name,
    character_data.zone_id,
    COALESCE((select zone.short_name from zone where zoneidnumber = character_data.zone_id LIMIT 1), "Not Found") as zone_name,
    character_data.zone_instance,
    character_data.x,
    character_data.y,
    character_data.z,
    character_data.heading,
    COALESCE((select guild_id from guild_members where char_id = character_data.id), 0) as guild_id,
    (select guilds.name from guilds where id = ((select guild_id from guild_members where char_id = character_data.id))) as guild_name,
    FROM_UNIXTIME(character_data.last_login) AS date_time
FROM
    character_data
WHERE
    last_login > (UNIX_TIMESTAMP() - 600)
ORDER BY character_data.name;
```

## Online Accounts Unique Addresses

Filters by characters saved within the past 10 minutes

```sql
SELECT
    character_data.account_id,
    character_data.name,
    character_data.zone_id,
    character_data.zone_instance,
    character_data.x,
    character_data.y,
    character_data.z,
    character_data.heading,
    COALESCE((select guild_id from guild_members where char_id = character_data.id), 0) as guild_id,
    (select guilds.name from guilds where id = ((select guild_id from guild_members where char_id = character_data.id))) as guild_name,
    FROM_UNIXTIME(character_data.last_login) AS date_time,
        (select ip from account_ip where account_ip.accid = character_data.account_id LIMIT 1) as ip_address,
        count(*) as count
FROM
    character_data
WHERE
    last_login > (UNIX_TIMESTAMP() - 600)
        group by ip_address
ORDER BY count desc;
```

## Local empty id's in a table

Displays all empty id's in a given table

```
SELECT
 CONCAT(z.expected, IF(z.got-1>z.expected, CONCAT(' thru ',z.got-1), '')) AS missing
FROM (
 SELECT
  @rownum:=@rownum+1 AS expected,
  IF(@rownum=id, 0, @rownum:=id) AS got
 FROM
  (SELECT @rownum:=0) AS a
  JOIN tablename
  ORDER BY id
 ) AS z
WHERE z.got!=0;
```

## List total number of epics by class/epic type

Lists total epics by class/epic type

```
SELECT
    c.class_name AS Class,
    i.Name AS Epic_Name,
    count(c.class_name) AS Count
FROM
    character_data a,
    inventory b,
    class_list c,
    items i,
    account z
WHERE
    a.id = b.charid AND
    a.class = c.class_id AND
    b.itemid = i.id AND
    a.account_id = z.id AND
    z.`status` BETWEEN '0' AND '15' AND
    a.is_deleted = '0' AND
    i.epicitem
GROUP BY Epic_Name
ORDER BY Class;
```
