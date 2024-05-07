# Dev Container Linux Setup

!!! info
     Check out [https://github.com/Akkadius/akk-stack](../../akk-stack/introduction.md) for ways to host your server, this instruction is more for development purposes. Akk-stack also offers dev environments.

Dev containers uses Docker and VSCode to create a development environment that is consistent across all developers. This is a great way to ensure that everyone is using the same tools and versions. You can also use this to develop on Windows, Mac, or Linux.

## Prerequisites

1. Install [Docker](https://www.docker.com/products/docker-desktop)
1. Install [VSCode](https://code.visualstudio.com/)

## Setup 1: Inside a dev container

This is recommended because you get autocomplete, debugging, and other features built in once your environment is spun up.

1. Start vscode
1. Clone this repository
1. Open the repository in VSCode
1. Click on the green button in the bottom left corner of vscode and select "Reopen in Container"
1. Wait for the container to build and start
1. Open a terminal in VSCode (ctrl+~)
1. Run `git submodule update --init --recursive` to pull in submodules
1. Use the build link on the bottom status bar, then select the GCC bottom option, or:
	1. Alternatively, run `make cmake` to configure environment.
	1. Follow this with `make build` command after.
1. Verify in your `build/bin/` folder that zone, world, etc exist

## Setup 2: Inject the database

1. Open a terminal in VSCode (ctrl+~) or your preferred terminal
1. Ensure your path is .devcontainer, if not run `cd .devcontainer`
1. Run `make inject-mariadb` to get a fresh peq snapshot from base

## Setup 3: Clone quests

1. Start a new vscode window, and git clone the quests folder into build/bin/quests

## Setup 4: Prep the environment

1. Open a terminal in VSCode (ctrl+~) or your preferred terminal
1. Ensure your path is .devcontainer, if not run `cd .devcontainer`
1. Run `make prep` to copy the required files to the build directory. This can be ran multiple times to "reset" the environment
1. Inside build/bin/eqemu_config.json edit the server name to your server's name, ideally with the keyword "dev" inside it.

## Setup 5: Connect to the database

1. Use Beekeeper or Heidi or your preferred SQL client to connect to the database
1. The settings: host: 127.0.0.1:3306, username: peq, database: peq, password: peqpass. You should see all tables populated

## Setup 6: Run shared memory

1. Open a terminal in VSCode (ctrl+~) or your preferred terminal
1. Ensure your path is .devcontainer, if not run `cd .devcontainer`
1. Run `make shared`. On success, it should return to the terminal at end
1. Above should only need to be ran once, or if you modify items, skill caps, spells, and other expansion tweaks, it may need to be ran again while the server is fully shut down

## Setup 7: Run world

1. Open a terminal in VSCode (ctrl+~) or your preferred terminal
1. Ensure your path is .devcontainer, if not run `cd .devcontainer`
1. Run `make world`. On success, it should be idle and not return to an empty prompt

## Setup 8: Run zone

1. Open a new terminal in VSCode (ctrl+~) or your preferred terminal
1. Ensure your path is .devcontainer, if not run `cd .devcontainer`
1. Run `make zone`. On success, it should be idle. If you switch back to world, you should see an event about zone connecting

## Setup 9: Log in and set yourself GM

1. Using a rof2 client where eqhost.txt is set to projecteq's login will make your debox server appear under `[D] Your Devbox Name`
1. When you first log in, the world console should say like "SetOnline Online Status [foo] (1) status [Online]
1. Ensure your path is .devcontainer, if not run `cd .devcontainer`
1. You can run `make gm-foo` to update your status to 255, replacing foo with your name


## Setup 10: Optional, get maps

Maps are a big download so unless you care about navmesh, water, los testing, this can be skipped.

1. Ensure your path is .devcontainer, if not run `cd .devcontainer`
1. Simply run `make maps`. This will download maps to base/maps, and the make prep command you ran earlier symlinks it.


## Debugging

You can optionally debug by attaching via gdb to the running processes. This is only recommended if you are familiar with gdb and debugging in general

1. Add breakpoints in zone or world that you want to debug
1. Press CTRL+SHIFT+D or the debug icon on left pane, and select a (gdb) [processType] for what you want to debug
1. When your code reaches said break point, you will get a call stack and variable dump on left
1. While the break point is triggered, all code is frozen, so a connection client will slowly disconnect if you're not quick to resume the breakpoint

### Quick and Dirty Alternative Test

If you're looking to compile binaries and not mess with the full dev container situation, you can use in linux or a WSL environment:

1. Ensure your path is .devcontainer, if not run `cd .devcontainer`
1. Run `make docker-cmake` to configure environment
1. Run `make docker-build` to build the project
1. Verify in your `build/bin/` folder that zone, world, etc exist
