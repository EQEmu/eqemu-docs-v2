---
description: Embedded Web REST API
---

# REST API

The login server has an embedded web server that servers an API on port 6000 unless otherwise specified in the config

{% code title="login.json" %}
```javascript
{
  "web_api": {
    "enabled": true, // enable/disable embedded webserver api
    "port": 6000 // the port you want the web api to serve on (recommended not to change)
  },
}
```
{% endcode %}

You have the option to disable it, or change the port altogether

## API Tokens

The login server requires API tokens to interact with the API, they can be generated via the CLI

### Using the CLI Interface

```bash
eqemu@dc25a75287d7:~/server$ ./loginserver web-api-token:create --write --read
```

The token will persist to the `login_api_tokens` table

```
MariaDB [peq]> select * from login_api_tokens;
+----+--------------------------------------+-----------+----------+---------------------+---------------------+
| id | token                                | can_write | can_read | created_at          | updated_at          |
+----+--------------------------------------+-----------+----------+---------------------+---------------------+
|  1 | b1bb9eb7-416c-487a-bb91-aedc59e0c57d |         1 |        1 | 2019-09-17 04:23:30 | 2019-09-17 04:23:30 |
+----+--------------------------------------+-----------+----------+---------------------+---------------------+
```

{% hint style="warning" %}
To create read only users, simply specify **only** the **--read **flag when creating users
{% endhint %}

{% hint style="info" %}
**Note **As a general rule, API calls that are **GET **are going to require **--read **while **POST** HTTP calls will require **--write**
{% endhint %}

## API Endpoints

The login server has a small handful of endpoints which have been created, it is easy to add more as needed

## PHP Client

There is an example of a PHP client that interacts with this API located here

{% embed url="https://github.com/Akkadius/eqemu-loginserver-php-api-client" %}

## Login Accounts

{% swagger baseUrl="http://loginserver:6000" path="/v1/account/create" method="post" summary="Login Account Create" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="path" name="Authorization" type="string" %}
Authorization: Bearer <token>
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
[
	{
		"local_ip" : "127.0.0.1",
		"players_online" : 0,
		"remote_ip" : "",
		"server_list_id" : 3,
		"server_long_name" : "Akkas Docker PEQ Installer (L)",
		"server_short_name" : "Akkas Docker PEQ Installer (L)",
		"server_status" : -2,
		"zones_booted" : 0
	}
]
```
{% endswagger-response %}
{% endswagger %}

{% code title="Request Body" %}
```
{
    "username": "test",
    "password": "test",
    "email": "<optional>"
}
```
{% endcode %}

{% swagger baseUrl="http://loginserver:6000" path="/v1/account/credentials/validate/local" method="post" summary="Validate Login Account Credentials" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="path" name="Authorization" type="string" %}
Authorization: Bearer <token>
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
{
    "data": {
        "account_id": 3,
    },
    "message": "Credentials valid!"
}
```
{% endswagger-response %}
{% endswagger %}

{% code title="Request Body" %}
```javascript
{
    "username": "test",
    "password": "test",
    "email": "<optional>"
}
```
{% endcode %}

{% swagger baseUrl="http://loginserver:6000" path="/v1/account/credentials/update/local" method="post" summary="Update Login Account Credentials" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="path" name="Authorization" type="string" %}
Authorization: Bearer <token>
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
{
    "message": "Loginserver account credentials updated!"
}
```
{% endswagger-response %}

{% swagger-response status="400" description="" %}
```
{
    "error": "Failed to update loginserver account credentials!"
}
```
{% endswagger-response %}
{% endswagger %}

```javascript
{
    "username": "test",
    "password": "newpassword"
}
```

## World Servers

{% swagger baseUrl="http://loginserver:6000" path="/v1/servers/list" method="get" summary="Get World Server List" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="path" name="Authorization" type="string" %}
Authorization: Bearer <token>
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
[
	{
		"local_ip" : "127.0.0.1",
		"players_online" : 0,
		"remote_ip" : "",
		"server_list_id" : 3,
		"server_long_name" : "Akkas Docker PEQ Installer (L)",
		"server_short_name" : "Akkas Docker PEQ Installer (L)",
		"server_status" : -2,
		"zones_booted" : 0
	}
]
```
{% endswagger-response %}
{% endswagger %}
