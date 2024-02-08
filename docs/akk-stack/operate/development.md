# Development 

!!! info
    
    This page will go detail tools and processes around using akk-stack for server development, spire development, etc.

## EQEmu Server Development

When **akk-stack** is installed, it by default uses pre-compiled binaries for those who have no interest in making modifications to the server. This uses stock binaries that are available via the release pipeline [here](https://github.com/EQEmu/Server/releases) and can be browsed in [Spire's release analytics](https://spire.akkadius.com/dev/releases)

### Compiling 

!!! info

    **Compiling for server development** within **akk-stack** has some extra features that are available to you out of the box that will save developers a tremendous amount of time.

If you are interested in compiling your own binaries, you can do so by following the instructions detailed below.

#### Initialize Development Build

Initialize your development build by using `make init-dev-build` when shelled into the home directory of the server container.

``` 
cd ~/ && make init-dev-build
```

**What this does**

* Initialize the build environment for you and set up the necessary tools to compile the server. This includes setting up the necessary dependencies, and setting up the necessary environment variables to compile the server.
* It uses a specific version of **[Perl](https://www.perl.org/)** that is known to work with the server, the **eqemu-server** container image is already pre-baked with all of the sane defaults for this to work flawlessly out of the box. It is decoupled from system Perl and is not affected by system Perl updates.
* It **disables compilation optimizations** and enables debugging symbols to be able to debug the server effectively and speeds up compilation tremendously because we're not doing optimization passes in the compiler.
* [Ccache](https://ccache.dev/) is also used to speed up compilation times, and the build is done using the **Ninja** build system. This is also pre-baked into the **eqemu-server** container image and you don't have to do any configuration beyond running the make command.
* Uses the **[gold](https://en.wikipedia.org/wiki/Gold_(linker))** linker which is known to be faster than the default linker that comes with the compiler.
* Uses the **[Clang](https://clang.llvm.org/)** compiler which is known to be faster than the default GCC compiler that comes with the compiler.
* Uses **Ninja** as the build system which is known to be faster than the default **make** build system that comes with the compiler.


```
> Initializing EQEmu Server Build

cd ~/code && \
	git submodule init && \
	git submodule update && \
	rm -rf build && \
	mkdir -p build && \
	cd build && \
	cmake -DEQEMU_BUILD_LOGIN=ON \
		-DEQEMU_BUILD_TESTS=ON \
		-DCMAKE_CXX_COMPILER_LAUNCHER=ccache \
		-DPERL_LIBRARY=/opt/eqemu-perl/lib/5.32.1/x86_64-linux-thread-multi/CORE/libperl.so \
		-DPERL_EXECUTABLE=/opt/eqemu-perl/bin/perl \
		-DPERL_INCLUDE_PATH=/opt/eqemu-perl/lib/5.32.1/x86_64-linux-thread-multi/CORE/ \
		-DCMAKE_EXE_LINKER_FLAGS=-fuse-ld=gold \
		-DCMAKE_CXX_FLAGS_RELWITHDEBINFO:STRING="-O0 -g -DNDEBUG -Os -Wno-everything" -G "Ninja" ..
-- The C compiler identification is Clang 14.0.6
-- The CXX compiler identification is Clang 14.0.6
-- Detecting C compiler ABI info
-- Detecting C compiler ABI info - done
... truncated output
```

#### Running the Build

**akk-stack** comes configured with a very simple alias that can be used anywhere in the shell to invoke the build, simply by typing `n`

```
n
```

Under the hood, this is just a simple alias that changes to the build directory and runs the build using the `ninja` build system.

``` 
alias n='cd ~/code/build && ninja -j$(expr $(nproc) - 2)'
```

You will see an output that resembles the following:

```
eqemu@66636230306d:~/code/build$ n
[90/130] Building CXX object zone/CMakeFiles/zone.dir/zone_base_data.cpp.o
```