---
description: Have an old Login Server from years ago? Here's your reference guide...
---

# Migrating from Legacy Login Server

## Migrating the Config

There is no straight migration path for the configuration file, you will simply have to reference the values you had in your `login.ini` and look at any that applies to the options in the new `login.json` format that is show in [Configuration](broken-reference)

## Migrating the Database

To migrate from a Legacy Login server database, you can use the provided `.sql` that you may find also available in the source at `loginserver/login_util/login_old_to_new_schema_convert.sql`

{% hint style="info" %}
Take note of the old table names in the script, if you used different table names (Because they were previously configurable) you will need to change the table names to reflect
{% endhint %}

### Migration SQL

{% code title="login_old_to_new_schema_convert.sql" %}
```sql
-- Because the old / legacy schema was mostly inconsistent with naming and overall data structure, we have
-- migrated to a schema that follows our modern conventions and meanwhile fixes quite a few bugs that
-- were present as well

-- Login Accounts

INSERT INTO
  login_accounts (
    id,
    account_name,
    account_password,
    account_email,
    source_loginserver,
    last_ip_address,
    last_login_date,
    created_at
  )
SELECT
  LoginServerID,
  AccountName,
  AccountPassword,
  AccountEmail,
  'local',
  LastIPAddress,
  LastLoginDate,
  AccountCreateDate
FROM
  tblLoginServerAccounts;

-- Server Admins

INSERT INTO
  login_server_admins (
    id,
    account_name,
    account_password,
    first_name,
    last_name,
    email,
    registration_date,
    registration_ip_address
  )
SELECT
  ServerAdminID,
  AccountName,
  AccountPassword,
  FirstName,
  LastName,
  Email,
  RegistrationDate,
  RegistrationIPAddr
FROM
  tblServerAdminRegistration;

-- World Servers

INSERT INTO
  login_world_servers (
    id,
    long_name,
    short_name,
    tag_description,
    login_server_list_type_id,
    last_login_date,
    last_ip_address,
    login_server_admin_id,
    is_server_trusted,
    note
  )
SELECT
  `ServerID`,
  `ServerLongName`,
  `ServerShortName`,
  `ServerTagDescription`,
  `ServerListTypeID`,
  `ServerLastLoginDate`,
  `ServerLastIPAddr`,
  `ServerAdminID`,
  `ServerTrusted`,
  `Note`
FROM
  tblWorldServerRegistration;
```
{% endcode %}
