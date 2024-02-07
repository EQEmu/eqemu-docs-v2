# Services

See the following table for a list of services that are available in the **akk-stack**

Optional services are enabled through the **.env** file

| **Service**      | **Description**                                                                                                                                                                              |
|------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **eqemu-server** | Runs the emulator server and all services related to the emulator server                                                                                                                     |
| **mariadb**      | MySQL service                                                                                                                                                                                |
| **phpmyadmin**   | **(Optional)** PhpMyAdmin which is automatically configured behind a password proxy                                                                                                              |
| **peq-editor**   | **(Optional)** PEQ Editor which is automatically configured                                                                                                                                      |
| **ftp-quests**   | **(Optional)** An FTP instance fully ready to be used to remotely edit quests                                                                                                                    |
| **backup-cron**  | **(Optional)** A container built to automatically backup (Dropbox API) the entire deployment and perform database and quest snapshots for with different retention schedules defined in **.env** |

## Enabling Services

In your **.env** file you can enable or disable services by setting the **ENABLE_** variable to **true** or **false** and will be started when you run **make up**

```
###########################################################
# Services
###########################################################
ENABLE_BACKUP_CRON=false
ENABLE_FTP_QUESTS=false
ENABLE_PEQ_EDITOR=false
ENABLE_PHPMYADMIN=false
```

## Service Lifetime

By default each container / service in the **docker-compose.yml** is configured to restart unless stopped, meaning if the server restarts the Docker daemon will boot the services you had started initially which is the default behavior of this stack

**Spire** and the **eqemu-server** entrypoint bootup script is designed to start the emulator server services when the server first comes up, so if you need to bring the whole host down, everything will **come back up on reboot automatically**
