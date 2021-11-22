---
description: Ideal for local LAN setups
---

# Auto Create Login Accounts

**Ideal for local LAN setups**, you can enable this setting. If any login request comes in from a game client, if the account doesn't already exist, it will **auto create the account on first attempt**. Every subsequent attempt will use the **created username and password combination **

{% code title="login.json snippet" %}
```javascript
{
  "account": {
    // ideal for local LAN setups, if you want a login attempt to automatically create an account
    // this will automatically create the account using the username and password if it doesn't exist
    "auto_create_accounts": true
  },
}
```
{% endcode %}

## Resetting Account Passwords

When using this configuration, it's easy for friends to forget their password or even your own after a long period. For this situation you can use the CLI interface to bail yourself out of this situation

```
./loginserver login-user:update-credentials test mynewpassword

> Executing CLI Command

[LoginServer] [Info] ChangeLoginserverUserCredentials account [test] source_loginserver [local] credentials updated!
```
