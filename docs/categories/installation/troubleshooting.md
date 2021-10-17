# Troubleshooting

## Windows

Below are some common issues that Windows users encounter when installing EQEmu Server.

### Cannot Connect

#### Improperly configured eqemu\_config

Please verify that you followed the instructions in the How-To Guide for [Configuring your eqemu\_config](configure-your-eqemu_config.md).

#### Windows Firewall

Be sure to enable TCP and UDP ports 7000 - 7500.  Run the following commands from your command prompt:

```text
netsh advfirewall firewall add rule name="EQEmu Zones (7000-7500) TCP" dir=in action=allow protocol=TCP localport=7000-7500
```

```text
netsh advfirewall firewall add rule name="EQEmu Zones (7000-7500) UDP" dir=in action=allow protocol=UDP localport=7000-7500
```



## Linux

Below are some common issues that Linux users encounter when installing EQEmu Server.

### Cannot Connect

#### Improperly configured eqemu\_config

Please verify that you followed the instructions in the How-To Guide for [Configuring your eqemu\_config](configure-your-eqemu_config.md).

