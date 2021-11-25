# Database Schema

If you are setting up the database from scratch, you can find the latest full table dump at the below link

[https://github.com/EQEmu/Server/blob/master/loginserver/login_util/login_schema.sql](https://github.com/EQEmu/Server/blob/master/loginserver/login_util/login_schema.sql)

If you have setup your server through an installer, chances are these tables were automatically installed and taken care of for you

## Tables


| Table  | Description  |
|---|---|
| [login_accounts](../../../schema/loginserver/login_accounts/) | Where Loginserver accounts are stored |
| [login_api_tokens](../../../schema/loginserver/login_api_tokens/)| (Optional) If using the API, where tokens are stored |
| [login_server_admins](../../../schema/loginserver/login_server_admins/)| Where Loginserver World server admins are stored |
| [login_world_servers](../../../schema/loginserver/login_world_servers/)| Where Worldservers are added when either registered manually or added automatically |
