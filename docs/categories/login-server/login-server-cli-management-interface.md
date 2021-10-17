---
description: >-
  Managing and interacting with your Login Server from the command line
  interface
---

# CLI Management Interface

In the latest overhaul of the login server, an enhanced command-line management interface was introduced to easily interface with basic management functionality needs. If you want to do more enhanced functionality such as integrating with the HTTP API, you can reference [API Endpoints](broken-reference)

## CLI Menu

![](<../../.gitbook/assets/image (17).png>)

### Help Menu

The login server command line help menu can be accessed by simply running the server binary from the command line with the option `--help`

```
eqemu@dc25a75287d7:~/server$ ./loginserver --help

> EQEmulator [LoginServer] CLI Menu

login-user
  login-user:check-credentials           Check user login credentials
  login-user:check-external-credentials  Check user external login credentials
  login-user:create                      Creates Local Loginserver Account
  login-user:update-credentials          Change user login credentials
web-api-token
  web-api-token:create                   Creates Loginserver API Token
  web-api-token:list                     Lists Loginserver API Tokens
world-admin
  world-admin:create                     Creates Loginserver World Administrator Account
```

### Command Example(s)

To see the required options and/or arguments for a command, simply type out the command to get the required context

```
loginserver login-user:create

Command

login-user:create {username} {password}

Options
  --email=*
```
