# REST API

The login server has an embedded web server that servers an API on port 6000 unless otherwise specified in the config.

The API is very basic in its current form but serves very simple use cases.

!!! warning

    For security reasons, its recommended that this port is exposed behind a secure HTTP proxy as the API surface itself is not encrypted and credentials are sent over the wire. Use responsibily.

=== "login.json"
      ```json
      {
        "web_api": {
          "enabled": true, // enable/disable embedded webserver api
          "port": 6000 // the port you want the web api to serve on (recommended not to change)
        },
      }
      ```

You have the option to disable it, or change the port altogether

## API Tokens

The login server requires API tokens to interact with the API, they can be generated via the CLI

## Using the CLI Interface

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

!!! warning
      To create read only users, simply specify **only** the --read flag when creating users


!!! info
      As a general rule, API calls that are GET are going to require --read while **POST** HTTP calls will require **--write**


## API Endpoints

The login server has a small handful of endpoints which have been created, it is easy to add more as needed

## PHP Client

There is an example of a PHP client that interacts with this API located here

[https://github.com/Akkadius/eqemu-loginserver-php-api-client](https://github.com/Akkadius/eqemu-loginserver-php-api-client)

## Accounts

!!! example

    BaseURL http://loginserver:6000

    Header **Authorization**: Bearer <token\>

    === "POST /v1/account/create"

        Create local loginserver account

        **Request Body**

        ```json
        {
            "username": "test",
            "password": "test",
            "email": "<optional>"
        }
        ```

        !!! responses

            === "200"

                ```json
                {
                    "data": {
                        "account_id": 1,
                    },
                    "message": "Account created successfully!"
                }
                ```
    
            === "400"
    
                ```json
                {
                    "error": "Account failed to create!"
                }
                ```

    === "POST /v1/accounts/credentials/validate/local"

        Validate local Loginserver account credentials

        **Request Body**

        ```json
        {
            "username": "test",
            "password": "test",
            "email": "<optional>"
        }
        ```

        !!! responses

            === "200"

                ```json
                {
                    "data": {
                        "account_id": 3,
                    },
                    "message": "Credentials valid!"
                }
                ```

            === "400"

                ```json
                {
                    "error": "Credentials invalid!"
                }
                ```

    === "POST /v1/account/credentials/update/local"

        Update local Loginserver account credentials

        **Request Body**

        ```json
        {
            "username": "test",
            "password": "test",
        }
        ```

        !!! responses

            === "200"

                ```json
                {
                    "message": "Loginserver account credentials updated!"
                }
                ```

            === "400"

                ```json
                {
                    "error": "Failed to find associated loginserver account!"
                }
                ```

                ```json
                {
                    "error": "Failed to update loginserver account credentials!"
                }
                ```


## World Servers

!!! example
    
    BaseURL http://loginserver:6000

    Header **Authorization**: Bearer <token\>

    === "GET /v1/servers/list"

        Lists connected worldservers (memory)

        **Result**

        ```json
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
