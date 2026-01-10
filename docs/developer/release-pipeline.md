# Release Pipeline

## Introduction

Server binary releases are simple to create. The heavy lifting of releases are automatically handled by the build pipeline but require developers to execute a few small steps in order to execute them.

## Versioning

EQEmulator Server roughly follows semantic versioning https://semver.org/. 

```
MAJOR version when you make incompatible API changes
MINOR version when you add functionality in a backwards compatible manner
PATCH version when you make backwards compatible bug fixes
```

We started at v22.1.0 because the project has been around for 22 years.

Version gets defined in **package.json** and it needs to be changed in order for the build pipeline to even consider building a release.

```json
{
  "name": "eqemu-server",
  "version": "22.1.0",
  "repository": {
    "type": "git",
    "url": "https://github.com/EQEmu/EQEmu.git"
  }
}
```

Version also gets updated in **common/version.h** it's not as important as **package.json** - but it should be set to at least relatively identify what local version people are running within their codebase.

```cpp
#define CURRENT_VERSION "22.1.0-dev" // always append -dev to the current version for custom-builds
```

While `CURRENT_VERSION` should be updated manually for releases, it always gets written during the build pipeline.

## Changelog

The changelog informs users of what was changed within the release - so they can be aware of expected changes. 

It is one of the most effective and important tools to communicate change with users.

Changelog is kept in **CHANGELOG.md** and follows https://keepachangelog.com/en/1.0.0/

### What is a changelog?

A changelog is a file which contains a curated, chronologically ordered list of notable changes for each version of a project.

### Why keep a changelog?

To make it easier for users and contributors to see precisely what notable changes have been made between each release (or version) of the project.

### Who needs a changelog?

People do. Whether consumers or developers, the end users of software are human beings who care about what's in the software. When the software changes, people want to know why and how.

### What makes a good changelog?

**Guiding Principles**

* Changelogs are for humans, not machines.
* There should be an entry for every single version.
* The same types of changes should be grouped.
* Versions and sections should be linkable.
* The latest version comes first.
* The release date of each version is displayed.
* Mention whether you follow Semantic Versioning.

**Types of changes**

* **Added** for new features.
* **Changed** for changes in existing functionality.
* **Deprecated** for soon-to-be removed features.
* **Removed** for now removed features.
* **Fixed** for any bug fixes.
* **Security** in case of vulnerabilities.

### Contents

To build the base of the changelog, we built a tool to build out the scaffolding of the changelog contents. 

This provides pull requests, suggest putting them in a section by them self called "Pull Requests" and providing other high level information above.

Choose the amount of days since the previous release to capture changes since then.

You may need to do some manual massaging of the content to make sure you don't have duplicate of what was potentially released in a version that was put out the same day or yesterday.

![image](https://user-images.githubusercontent.com/3319450/214199167-f1430a4c-a784-401c-8329-b8d1c5a5c802.png)

### Example Contents Diff

Below is an example of the changes needed to make a release

```diff
diff --git a/CHANGELOG.md b/CHANGELOG.md
index 323dea8de..0d2ac4028 100644
--- a/CHANGELOG.md
+++ b/CHANGELOG.md
@@ -1,3 +1,10 @@
+## [22.1.1] - 01/23/2022
+
+### Fixes
+
+* Fix botgrouplist to display unique entries. ([#2785](https://github.com/EQEmu/EQEmu/pull/2785)) ([Aeadoin](https://github.com/Aeadoin)) 2023-01-23
+* Fix scenario where dereferenced object could be null. ([#2784](https://github.com/EQEmu/EQEmu/pull/2784)) ([Aeadoin](https://github.com/Aeadoin)) 2023-01-23
+
 ## [22.1.0] - 01/22/2022
 
 This is a first release using the new build system. Changelog entry representative of last year. Subsequent releases will consist of incremental changes since the last release.
diff --git a/common/version.h b/common/version.h
index 79ad6b81b..5c47be7f1 100644
--- a/common/version.h
+++ b/common/version.h
@@ -25,7 +25,7 @@
 
 // Build variables
 // these get injected during the build pipeline
-#define CURRENT_VERSION "22.1.0-dev" // always append -dev to the current version for custom-builds
+#define CURRENT_VERSION "22.1.1-dev" // always append -dev to the current version for custom-builds
 #define LOGIN_VERSION "0.8.0"
 #define COMPILE_DATE    __DATE__
 #define COMPILE_TIME    __TIME__
diff --git a/package.json b/package.json
index 8f51d0ec9..4eaa1d797 100644
--- a/package.json
+++ b/package.json
@@ -1,6 +1,6 @@
 {
   "name": "eqemu-server",
-  "version": "22.1.0",
+  "version": "22.1.1",
   "repository": {
     "type": "git",
     "url": "https://github.com/EQEmu/EQEmu.git"
```

## Releasing

Once you have the above changes ready to go, merge it into master it will trigger a pipeline build and the pipeline will take care of the rest!

See [Build Pipeline](build-pipeline.md)

![image](https://user-images.githubusercontent.com/3319450/214204709-6ddd4383-99c3-49cc-9ae6-9797ab37d8bc.png)

![image](https://user-images.githubusercontent.com/3319450/214204806-0b139dc7-4a5e-4f33-b22b-37d82c91fbef.png)

## Release Analytics

We have an in-house tool built to track and manage releases and crash analytics for server developers.

It can be found at http://spire.akkadius.com/dev/releases

![image](https://user-images.githubusercontent.com/3319450/214204952-1f66fafa-16e4-4e10-8a3a-f4691610eabc.png)

