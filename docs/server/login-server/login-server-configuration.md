---
description: Login Server configuration options
---

# Configuration

The login server configuration file is stored in the server root and is named `login.json` below is the most recent configuration per **September 17, 2019**

!!! warning
      **Warning** Be careful with how you configure your Login server as it some settings may not be ideal for your needs


```javascript
{
  "database": {
    "host": "127.0.0.1", // database host
    "port": "3306", // database port
    "db": "peq", // database name
    "user": "eqemu", // database user
    "password": "eqemu" // database password
  },
  "account": {
    // ideal for local LAN setups, if you want a login attempt to automatically create an account
    // this will automatically create the account using the username and password if it doesn't exist
    "auto_create_accounts": true
  },
  "worldservers": {
    "unregistered_allowed": true, // allows worldservers to connect to your loginserver without server admin authentication
    "reject_duplicate_servers": false // if enabled, rejects duplicate worldservers
  },
  "web_api": {
    "enabled": true, // enable/disable embedded webserver api
    "port": 6000 // the port you want the web api to serve on (recommended not to change)
  },
  "security": {
    "mode": 14, // encryption mode (dont touch) (14=scrypt)
    "allow_password_login": true, // allows users to login via password, most cases, leave this on
    "allow_token_login": true // allows token based login directly from launching game
  },
  "logging": {
    "trace": false, // For debugging general packet messaging
    "world_trace": false, // For debugging world to loginserver messaging
    "dump_packets_in": false, // for debugging inbound packets
    "dump_packets_out": false // for debugging outbound packets
  },
  "client_configuration": {
    "titanium_port": 5998, // don't change
    "titanium_opcodes": "login_opcodes.conf", // opcodes for the titanium era clients
    "sod_port": 5999, // don't change
    "sod_opcodes": "login_opcodes_sod.conf" // opcodes for sod and higher era clients
  }
}
```

