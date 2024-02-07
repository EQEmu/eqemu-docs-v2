### Updating Spire

From within the **eqemu-server** container at the root of the home directory, you can run the following command to update the Spire server binaries. Do note that Spire can also be updated from within the Spire admin panel.

```
make update-admin-panel
```

It will kill the currently running panel, cycle it out, start it up. This is not service affecting for running servers with a launcher running.

<img src="https://github.com/Akkadius/spire/assets/3319450/b0576ea5-0ce4-4062-a456-768985faca4b" style="border-radius: 10px">

### Updating Server Binaries

Updating server binaries is as simple as running **update** in the **eqemu-server** shell, it will change directory to the source directory, git pull and run a build which will be immediately available the next time you boot a process

You can update using server release binaries by running

```
update
```

You can update by compiling the source yourself by running

```
update-source
```

You can also update binaries within the Spire Admin panel. 

![image](https://github.com/Akkadius/akk-stack/assets/3319450/69b54e00-ede1-4c05-af72-3a2c50a592ca)