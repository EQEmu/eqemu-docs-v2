---
description: This article describes the ability to use multiple database and data sources
---

# Multi Tenancy



{% hint style="warning" %}
**Warning **This feature is for **advanced server operators**
{% endhint %}

## What is Multi Tenancy in Software?

There are several types of multi-tenancy in Software, the one we're going to be talking about implemented in the server backend is described as simply being able to use different database connections for different types of data

{% hint style="info" %}
Multi-tenant architecture hosts data in multiple databases. This model is relatively complex in terms of management and maintenance, but tenants can be separated by a chosen criterion.
{% endhint %}

![](<../../.gitbook/assets/image (19).png>)

## How is Multi-Tenancy Used in the Server?

As of this writing today; you can optionally declare a separate database connection other than your main database connection that can be a source of **content data**

**Content data** includes world data such as **NPCs, Zones, Merchants, Spawns, Doors, Objects** etc.

### Configuration

Today you can declare your database configuration like so; everything will load out of this connection by default, using a content database connection is **entirely optional**

```javascript
cat eqemu_config.json | jq '.server.database'
{
  "db": "peq",
  "host": "mariadb",
  "port": "3306",
  "username": "eqemu",
  "password": "eqemu"
}
```

Now if you wanted to configure a content database source you would declare another database connection section that resembles the following

```javascript
cat eqemu_config.json | jq '.server.content_database'
{
  "db": "peq_content",
  "host": "content-cdn.projecteq.net",
  "port": "16033",
  "username": "content-user",
  "password": "content-password"
}
```

## Why Would I Want to Use This?

The notion of having a content database source is very powerful because it solves for the following use cases

* **Developing content on a local development server** which has access to a remote content database connection shared with production, allowing a developer to constantly **repop zones, refresh scripts locally** without impacting the state of the **production zone server**
* A developer can **develop scripts locally**, create database entries that are assigned an ID that will be used in production without having to reconcile data and duplicate ID's in an isolated local database up to a remote database which is very kludgy and cumbersome
* It **decouples the responsibilities of table types** in the source, content data should not be written to systemically from a server to store player state, zone state, world state and only hold mostly read only data for content
* It allows **many production servers** to use the same exact data source; **heavily **eliminating duplication of work and maintenance of keeping several different content datasets up to date. 
* It allows someone to make **different variants / flavors** of the same server using **different local rulesets**
* It enables the [Project PEQ Expansions](../../in-development/project-peq-expansions/) project to be able to have many **standing replicas of each Expansion / Era;** enabling PEQ developers to jump between eras and maintain era accuracy

![Many development servers using the same content database](<../../.gitbook/assets/image (18).png>)

## How Do I Know What Tables Belong to Which Category?

Database table schema reference is maintained in the source at **./common/database_schema.h** and all table types are defined under **DatabaseSchema::GetContentTables()**

{% embed url="https://github.com/EQEmu/Server/blob/master/common/database_schema.h#L165" %}

