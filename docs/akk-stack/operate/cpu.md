# CPU

Akk Stack employs a few different methods to protect the CPU from being over-utilized and to ensure that the server is always responsive.

## CPU Limiting

The **eqemu-server** container is limited to 90% of the CPU by default. This is to ensure that the server is always responsive and that the host is not over-utilized. This can be changed in the **docker-compose.yml** file.

```
cat docker-compose.yml | grep shares
    cpu_shares: 900
```

https://docs.docker.com/config/containers/resource_constraints/#configure-the-default-cfs-scheduler

## CPU Watchdog

The **eqemu-server** container has a watchdog that will kill processes that are running high for too long. This is to protect from common occurrences where a server operator accidently creates a loop in a script causing the zone process CPU to spike to 100% and become entirely unresponsive. The watchdog is designed to detect and kill these processes.

You can check the logs if you suspect a process is triggering the watchdog from the root of the **eqemu-server** home directory.

```
cat process-kill.log
```

```
Sat Jul 11 20:52:47 CDT 2020 [process-watcher] Killed process [21143] [./bin/zone] for taking too much CPU time [43.50]
```