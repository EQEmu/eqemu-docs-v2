!!! info
    Automated cron-based backups that upload to Dropbox using Dropbox API

<img src="https://github.com/Akkadius/akk-stack/assets/3319450/4f2240cb-4e5a-4433-b7dc-1f22a7b8a40b" style="border-radius: 10px">

### Initialize Backups

To get started, you need to run the uploader script in the backup-cron container for the first time to initialize your application

As of July 2021 this guide has changed to Dropbox's new auth mechanisms to include more configuration in the OAuth flow.

```
docker-compose exec backup-cron dropbox_uploader.sh
```

Follow the instructions prompted from running the command

```

 This is the first time you run this script, please follow the instructions:

(note: Dropbox will change their API on 2021-09-30.
When using dropbox_uploader.sh configured in the past with the old API, have a look at README.md, before continue.)

 1) Open the following URL in your Browser, and log in using your account: https://www.dropbox.com/developers/apps
 2) Click on "Create App", then select "Choose an API: Scoped Access"
 3) "Choose the type of access you need: App folder"
 4) Enter the "App Name" that you prefer (e.g. MyUploader1167208717053), must be unique

 Now, click on the "Create App" button.

 5) Now the new configuration is opened, switch to tab "permissions" and check "files.metadata.read/write" and "files.content.read/write"
 Now, click on the "Submit" button.

 6) Now to tab "settings" and provide the following information:
 App key: dmz4wbjsnghfkwj
 App secret: iq26gmwnlsnwj48
  Open the following URL in your Browser and allow suggested permissions: https://www.dropbox.com/oauth2/authorize?client_id=dmz4wbjsnghfkwj&token_access_type=offline&response_type=code
 Please provide the access code: Bun8T-9NG2kAAAAAAABF0by79e-VuivtOXRtHkS10KA                                                                                               

 > App key: xxx
 > App secret: 'xxx
 > Access code: 'Bun8T-9NG2kAAAAAAABF0by79e-xxx'. Looks ok? [y/N]: y
   The configuration has been saved.
```

Once you go through the steps of creating your application. Do not forget to set scopes on your app to be able to write and read files. You MUST follow the prompts above in order otherwise you will run into issues.

<img src="https://user-images.githubusercontent.com/3319450/174466660-b9db68db-5a3e-4877-b55d-1ceaa249bb6c.png" style="border-radius: 15px">

<img src="https://user-images.githubusercontent.com/3319450/174466869-b06d9170-3fd6-4057-85a6-238b905fc7d8.png" style="border-radius: 15px">

Your configuration gets written to `.dropbox_uploader` which resides at the root of your deployment. This is a sensitive file and is not to be checked into any sort of version control and is used by the `backup-cron` container

### Validate it Works!

Run `make backup-dropbox-list`

```
make backup-dropbox-list
docker-compose up -d backup-cron
docker-compose exec backup-cron dropbox_uploader.sh list
 > Listing "/"... DONE
```

If it shows `> Listing "/"... DONE` then it is initialized successfully

You can test by running a backup

```
make backup-dropbox-database
docker-compose exec backup-cron ./backup/backup-database.sh
# Dumping database and compressing
peq-06-19-2022.sql
# Uploading database snapshot
 > Uploading "/tmp/peq-06-19-2022.tar.gz" to "/backups/database-snapshots/peq-06-19-2022.tar.gz"... DONE
# Truncating backups/database-snapshots days back 7
# Cleaning up...
```

### Backup Configuration

Backup retention configurable in `.env`

Your deployment name is what your backups will be prepended to when they get uploaded to Dropbox

```
# DEPLOYMENT_NAME=peq-production
# BACKUP_RETENTION_DAYS_DB_SNAPSHOTS=10
# BACKUP_RETENTION_DAYS_DEPLOYMENT=35
# BACKUP_RETENTION_DAYS_QUEST_SNAPSHOTS=7
```

Crons defined in `backup/crontab.cron`

Crons are configured to run on a variance so that not all deployments fire backups at the same time

| **Backup Type** | **Description** | **Schedule** |
|---|---|---|
| Deployment | Deployment consists of the entire akk-stack folder (server, database etc.). If you ever experienced catastrophic failure or needed to restore the entire setup, simply restoring the deployment folder will get you back up and running | Once a week at 1AM on a random variance of 1800 seconds |
| Quests | A simple snapshot of the quests folder | Once a day at 1M on a random variance of 1800 seconds |
| Database | A simple snapshot of the database | Once a day at 1M on a random variance of 1800 seconds |

### Running Backups Manually

Bash into the `backup-cron` service; assuming your OAUTH token is valid and everything works

```
root@host:/opt/eqemu-servers/peq-production# docker-compose exec backup-cron bash
```

```
backup-cron@backup-cron:~$ dropbox_uploader.sh list peq-production
 > Listing "/peq-production"... DONE
 [D]  database-snapshots
 [D]  deployment-backups
 [D]  quest-snapshots
```

**Database Snapshots**

```
backup-cron@backup-cron:~$ dropbox_uploader.sh list peq-production/database-snapshots
 > Listing "/peq-production/database-snapshots"... DONE
 [F] 182189205 peq-07-02-2020.tar.gz
 [F] 182222834 peq-07-03-2020.tar.gz
 [F] 182263995 peq-07-04-2020.tar.gz
 [F] 182300144 peq-07-05-2020.tar.gz
 [F] 182394017 peq-07-06-2020.tar.gz
 [F] 182464528 peq-07-07-2020.tar.gz
 [F] 182465093 peq-07-08-2020.tar.gz
 [F] 182527952 peq-07-09-2020.tar.gz
 [F] 182574977 peq-07-10-2020.tar.gz
 [F] 182566469 peq-07-11-2020.tar.gz
 [F] 182661537 peq-07-12-2020.tar.gz
 ...
```

**Deployment Snapshots**

(Includes entire deployment folder)

```
 backup-cron@backup-cron:~$ dropbox_uploader.sh list peq-production/deployment-backups
 > Listing "/peq-production/deployment-backups"... DONE
 [F] 3309179293 deployment-07-02-2020.tar.gz
 [F] 2357754207 deployment-07-05-2020.tar.gz
 [F] 2364156848 deployment-07-12-2020.tar.gz
 ...
```

***Quest Snapshots***

```
backup-cron@backup-cron:~$ dropbox_uploader.sh list peq-production/quest-snapshots
 > Listing "/peq-production/quest-snapshots"... DONE
 [F] 29464443 quests-07-07-2020.tar.gz
 [F] 29464443 quests-07-08-2020.tar.gz
 [F] 29464443 quests-07-09-2020.tar.gz
 [F] 29464443 quests-07-10-2020.tar.gz
 [F] 29464443 quests-07-11-2020.tar.gz
 [F] 29464443 quests-07-12-2020.tar.gz
 ...
```