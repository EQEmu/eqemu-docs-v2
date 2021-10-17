# Ground Up Windows Install

{% hint style="info" %}
This guide is intended for advanced users who want to configure EQEmu "from the ground up" for a windows server. It is not the recommended route, we suggest using the [Windows-Server install ](server-installation-windows.md)instead for most cases

This guide provides instructions for compiling 32-bit Windows server binaries and setting up a local development server

Read through this guide before starting to ensure an understanding of the process
{% endhint %}

{% hint style="info" %}
The target computer must have an active internet connection at the time of installation.

Please direct any questions to our server support channel in \[[discord](https://discord.gg/QHsm7CD)]

_Screenshots may vary depending on options selected and program graphical user interface changes._
{% endhint %}

## Compiler Setup

The current c/c++ support standard of the EQEmulator server code base mandates the use of Visual Studio 2017 or later compilers.

Visual Studio 2017 is the current EQEmulator standard for binary compilation. Please ensure that your system meets the \[[Visual Studio 2017 Minimum System Requirements](https://docs.microsoft.com/en-us/visualstudio/productinfo/vs2017-system-requirements-vs#visual-studio-2017-system-requirements)].

Visual Studio 2019 may also be used..though, less is known about the stability of this platform with the EQEmulator code base.

This setup assumes an install on a 64-bit Windows operating system with 32-bit target binaries.

## Verify System Environment Variable %Path% Length

Sometimes an automated server installation will fail due to the %Path% variable being full. This can happen with a manual installation as well.

Since this guide installs more programs than are required for server operation alone, verifying the length of %Path% is critical before we start

### Finding Your Current Path

* Click on your `Start` button
* Type "environment variable" into your `Search programs and files` text box
* Click on `Edit system environment variables`
* Click the `Environment Variables` button on the window that pops up
* Ensure the variable `Path` is selected in the system variables section (bottom text box)
* Click the `Edit` button

You may check the length of your %Path% variable by copying the `Variable value` contents and pasting them into a text editor that supports "selection" count.

![](https://user-images.githubusercontent.com/3311166/60465078-3b56b180-9c1e-11e9-89c5-0137a6ed84ba.png)

Registry values are only allocated 1024 bytes of storage. However, environmental variables may contain up to 2048 bytes through the use of an alias.

If your selection count is greater than 768 characters, you may need to setup an alias to prevent corruption of the %Path% variable.

## Required Programs

Some of the pre-requisites for compiling binaries are the same as running a server.

If you have already installed any of the following, the download and installation requirement should be omitted:

*   Visual Studio 2017 Community Edition \[[select Visual Studio 2017 Community Edition](https://visualstudio.microsoft.com/vs/older-downloads/)]

    _Note: Microsoft now requires a user account to download Visual Studio. Clicking the Visual Studio link above will take you to the "older versions" page. Clicking the_ `Download` _button on that page will prompt you to log in or create an account._
*   Visual Studio 2019 Community Edition \[[alternative download](https://visualstudio.microsoft.com/vs/)]

    _Note: Only install one version of Visual Studio_
* Perl v5.12.3.1204 (32-bit) \[[download](https://github.com/EQEmu/eqemu.github.com/raw/master/downloads/ActivePerl-5.12.3.1204-MSWin32-x86-294330.msi)]
* CMake (64-bit) \[[download](https://github.com/Kitware/CMake/releases/download/v3.14.5/cmake-3.14.5-win64-x64.msi)]
* Git (64-bit) \[[download](https://git-scm.com/download/win)]
*   TortoiseGit (64-bit - optional) \[[download](https://download.tortoisegit.org/tgit/2.8.0.0/TortoiseGit-2.8.0.0-64bit.msi)]

    _Note: TortoiseGit is a menu-driven, add-on gui interface for Git. Though optional, this instructional provides for its use._

Some programs may be able to use newer versions, or even their latest releases, without issue. But, this is not the case with Perl and (later) dependencies.

The above list of programs is known to work for compiling working server binaries.

### Install Visual Studio

During the install process, ensure the option for `Desktop development with C++` is checked.

![](https://user-images.githubusercontent.com/3311166/60468475-b40e3b80-9c27-11e9-8b2b-462bd0f22165.png)

{% hint style="info" %}
This package is required by Visual Studio to compile c/c++ code and by CMake to determine available compiler options. It will also cause CMake file generation to fail, if not enabled
{% endhint %}

If you selected Visual Studio 2017 Community Edition, you will need to update to the most current version.

{% hint style="info" %}
This requirement is not needed for Visual Studio 2019 installations..but, it is a good idea to have the most up-to-date compiler features
{% endhint %}

### Install Perl

This installation is self-explanatory.

_Note: It is recommended that you install in the root folder (_`c:\`_) to avoid possible issues._

### Install CMake

This installation is self-explanatory.

### Install Git

This installation is self-explanatory.

### Install TortoiseGit

This installation is self-explanatory.

### Restart Computer

You will need to restart your computer to ensure that all of the %Path% additions are loaded into memory.

### Acquiring the Code

At this point, you will need to make a decision on how you want to manage your code.

There are two options:

1. Create a local repository from the parent EQEmulator project that can be updated, managed and maintained (recommended)
2.  Create a local repository from a fork of the EQEmulator project that you manage (optional - select only if you want to contribute back to the parent project)

    _Note: If you choose to create a fork of the EQEmulator repository, you will need to create a _\[[github.com](https://github.com)] _account._

If you chose option 1, create a sub-folder called `git-eqemulator` in the root folder of c: drive.

If you chose option 2, create a sub-folder called `git-<git-username>` in the root folder of c: drive. (example: git username is `Pavlov`, folder name would be `git-pavlov`)

The purpose of this folder is to facilitate code management. We'll refer to this as the `account`folder.\


Go to the EQEmulator server code repository web page at [https://github.com/EQEmu/Server](https://github.com/EQEmu/Server).\


If you chose option 2 and are creating a fork, click on the fork button to add the repository to your github account. You should be redirected to your fork's main repository page.

![](https://user-images.githubusercontent.com/3311166/60550151-0e290280-9cf5-11e9-966e-e8d1fec1c80e.png)

Finally, click the `Clone or download` button, then `Open in Desktop` button to create a local code repository on your computer. Select `TortoiseGit` for the program to use. When prompted for where to create it, select the `account` folder created above.

![](https://user-images.githubusercontent.com/3311166/60550164-21d46900-9cf5-11e9-882b-439dea49c8e8.png)

You should now have a managed local code repository on your computer.

_Note: It is helpful to create a shortcut to the_ `account` _folder and place it on your desktop._\


### Install Submodules

To install the required submodules:

* Open your `account` folder
* Right-click the `Server` folder and select `Git Sync`
*   Find the `Submodule` button and click the attached arrow to activate the drop-down menu

    ![](https://user-images.githubusercontent.com/3311166/60761743-9c1a2b80-a01c-11e9-8dea-8609a1f5aa5d.png)
* Select `Submodule Init` - this should initiate the action
* Click the `Submodule` button arrow again to activate the menu
* Select `Submodule Update` - again, initiating action

Submodules are now installed.\


### Install Dependencies

To install the required dependencies:

* Download the \[[dependencies](https://github.com/EQEmu/eqemu.github.com/raw/master/downloads/WindowsDependencies_x86.zip)] file
* Navigate down to `c:\<account>\Server\dependencies`
* Copy the downloaded file into the folder
* Unpack the file

Dependencies are now installed.\


### Install VCPkg

To install VCPkg:

* Download the \[[vcpkg](https://github.com/EQEmu/Server/releases/download/v1.2/vcpkg-export-x86.zip)] file
* Navigate down to `c:\<account>\Server`
* Create a folder called `vcpkg`
* Navigate down to `c:\<account>\Server\vcpkg`
* Copy the downloaded file into the folder
* Unpack the file

VCPkg is now installed.\


## Running CMake

CMake's default options are adequate to configure and generate the files needed for Visual Studio.

There are two folder locations that you will need to provide:

* `Where is the source code:`
* `Where to build the binaries:`

For the source code, type-in or navigate to your `c:/<account>/Server` folder.

The easiest way to define the build folder is to copy the source path and paste it in. Then, add `/build` to the end of the path so that you have `c:/<account>/Server/build`.

Once CMake knows where to look, click the `Configure` button. You will get a pop-up window stating that the `build` folder does not exist. Click `OK` to create it.\


The next window will be for compiler selection. Ensure that the version of Visual Studio that you installed is selected (`Visual Studio 15 2017` or `Visual Studio 16 2019`). Leave the `Use default native compilers` option as-is. If you are using a newer version of CMake that has the `Optional platform for generator` drop-down box, ensure that `Win32` is selected. Finally, click `Finish` to proceed.

![](../../.gitbook/assets/cmake_compiler_updated\_2.png)

_Note: CMake will download the appropriate vcpkg and install its contents. User intervention is no longer required._\


You should now have a list of unconfigured options (in red) displayed in the main window of CMake:

![](https://user-images.githubusercontent.com/3311166/60629816-54e62d80-9dc5-11e9-89ab-8961e94c9491.png)

The following list contains the most common options of interest to the majority of users:

* `EQEMU_BUILD_CLIENT_FILES` \[_enabled_] Builds binaries used to import/export client support files
* `EQEMU_BUILD_LOGIN` \[_disabled_] Builds the login server (this guide makes use of the login server - change this option to _**enabled**_)
* `EQEMU_BUILD_LUA` \[_enabled_] Compiles server code with Lua support
* `EQEMU_BUILD_PERL` \[_enabled_] Compiles server code with Perl support
* `EQEMU_DEBUG_LEVEL` \[_5_] Determines what additional messaging and debugging code is enabled/disabled (_12_ is max)
* `EQEMU_ENABLE_BOTS` \[_disabled_] Compiles server code with Bot support (user choice)

{% hint style="info" %}
Ensure that you set `EQEMU_BUILD_LOGIN` to **enabled**
{% endhint %}

Once you have set the options that you would like for your server, click `Configure` again.

Since we set an option (login server) that requires additional settings, more unconfigured options have appeared. In this case, the open ssl library:

![](https://user-images.githubusercontent.com/3311166/60629837-67606700-9dc5-11e9-8b8c-560e629eae27.png)

These new options are only file path definitions. No additional changes need to be made. Click `Configure` one last time.

![](https://user-images.githubusercontent.com/3311166/60629846-76471980-9dc5-11e9-9b30-d7059103ed89.png)

{% hint style="info" %}
Regardless of option settings, anytime that you have red (unconfigured) entries in your options list, you will need to click `Configure` to ensure that the settings are applied to the current CMake file generation template.
{% endhint %}

You can now click the `Generate` button.

If file generation was successful, you should see "Generating done" at the bottom of the CMake window.

You are now ready to open Visual Studio and compile your code!

## Compiling Code

To open the CMake-generated solution file:

* Navigate through the desktop `account` folder shortcut to `c:\<account>\Server\build`
* Right-click the `EQEmu.sln` file
* Select `Open with` >> (`Microsoft Visual Studio 2017` or `Microsoft Visual Studio 2019`)

Upon loading, Intellisense will begin mapping out the entire project. Allow a few seconds for this process to finish. The lower left-hand corner will display "Ready" when this process has completed

Visual Studio does not honor CMake's `CMAKE_BUILD_TYPE` variable. You will need to manually set this to the desired build type before compiling your code.

![](https://user-images.githubusercontent.com/3311166/60749564-bfd06980-9f69-11e9-9688-a2efd8485ea5.png)

There are 4 options:

* `Debug` - not recommended for production servers unless conditions warrant its use (i.e., load testing, trouble-shooting, etc...)
* `RelWithDebugInfo` - standard compile for production servers (provides debug symbols without the added overhead of extra code, memory buffers and stl assertions)
* `Release` - similar to `RelWithDebugInfo`..but, without access to debug symbols (not recommended)
* `MinSizeRel` - Same as `Release` with the exception of trading off faster code for smaller size (not recommended)

{% hint style="info" %}
**Debug** _is _preferred_ for a local test server_
{% endhint %}

To compile your server code, you have two choices:

* Use the menu path under the top menu (`Build` >> `Build`)
* Use the menu path by right-clicking under `Solution Explorer` (`Solution 'EQEmu'` >> `Build`)

![](https://user-images.githubusercontent.com/3311166/60749565-c3fc8700-9f69-11e9-8b54-fa331602d8c3.png)

{% hint style="info" %}
Both paths result in the same action. Use whichever you are more comfortable with
{% endhint %}

The compiled code will be located in the `c:\<account>\Server\build\bin\<build_type>`folder.

### Code Maintenance

* Working Tree Changes
* Master Branch Code Updates
* Dependencies and Submodules Updates
* Branches and Stashes
* Reverting Changes
* Resetting to Previous Commits

