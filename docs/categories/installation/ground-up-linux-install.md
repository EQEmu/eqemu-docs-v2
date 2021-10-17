# Ground Up Linux Install

{% hint style="info" %}
It is recommended to use the [**Linux-Server install**](server-installation-linux.md)** **in nearly all cases, this is for advanced users/developers. Note the Linux-Server install above DOES include source code building, so it can do all steps below.
{% endhint %}

ubuntu apt-get install requirements. (you can refer to [https://github.com/EQEmu/Server/blob/master/utils/scripts/linux_installer/install.sh#L91](https://github.com/EQEmu/Server/blob/master/utils/scripts/linux_installer/install.sh#L91)) for other environment installs

```
apt-get -y install --no-install-recommends build-essential gcc-5 g++-5 libtool cmake curl debconf-utils  git git-core libio-stringy-perl liblua5.1  liblua5.1-dev libluabind-dev libmysql++  libperl-dev libperl5i-perl libjson-perl libsodium-dev  libmysqlclient-dev libssl-dev lua5.1  minizip make mariadb-client locales  nano open-vm-tools unzip uuid-dev iputils-ping  zlibc wget
```

ubuntu libsodium install (needs 18)

```
wget http://ftp.us.debian.org/debian/pool/main/libs/libsodium/libsodium-dev_1.0.11-2_amd64.deb -O /tmp/libsodium-dev.deb
wget http://ftp.us.debian.org/debian/pool/main/libs/libsodium/libsodium18_1.0.11-2_amd64.deb -O /tmp/libsodium18.deb
dpkg -i /tmp/libsodium*.deb
rm -f /tmp/libsodium-dev.deb
rm -f /tmp/libsodium18.deb
```

you need to install and configure mariadb or mysql yourself. I didn't include this step, since with a ground up install, I assume you have your own means to provide it. 

These environment variables are set just to simplify snippets in the future. Feel free to adjust to your preferred locations

```
export EMUBUILDDIR=~/eqemu/src/build/bin
export EMUSRCDIR=~/eqemu/src
mkdir -p $EMUBUILDDIR
mkdir -p $EQEMUSRCDIR
```

Clone eqemu's source code from github

```
git clone --recurse-submodules https://github.com/EQEmu/Server.git $EMUSRCDIR
```

configure eqemu's source code to generate files into the build dir

```
mkdir -p $EMUSRCDIR/build
cd $EMUSRCDIR/build
cmake -DEQEMU_BUILD_LOGIN=ON -DEQEMU_BUILD_LUA=ON -G 'Unix Makefiles' ..
```

compile eqemu's source code

```
make
```

verify binaries were built

```
ls $EMUBUILDDIR/world
> ~/eqemu/build/bin/world
ls $EMUBUILDDIR/zone
> ~/eqemu/build/bin/zone
```

download eqemu_config.json (edit this file to your environment, especially DB settings)

```
cd $EQEMUBUILDDIR/ 
wget --no-check-certificate https://raw.githubusercontent.com/Akkadius/EQEmuInstall/master/eqemu_config_docker.json -O eqemu_config.json
```

prime database (This step can at times cause cmake to break, just to warn, since linux_login_server_setup likes to overwrite the cmake configuration. Remember remove CMakeCache.txt and rerun cmake to repair if this happens)

```
cd $EQEMUBUILDDIR/
./eqemu_server.pl source_peq_db 
./eqemu_server.pl check_db_updates 
./eqemu_server.pl linux_login_server_setup
```

edit your login.json file to connect to your DB

get utility scripts, opcodes, endless other things

```
cd $EQEMUBUILDDIR/
./eqemu_server.pl fetch_utility_scripts
./eqemu_server.pl opcodes
./eqemu_server.pl maps #this is going to take a while, ~2 gig DL
./eqemu_server.pl opcodes
./eqemu_server.pl plugins
./eqemu_server.pl quests
./eqemu_server.pl lua_modules
```

I suggest at this point going into your $EQEMUBUILDDIR and running ./shared_memory, ./world, ./zone, ./loginserver and reviewing any errors it reports. If no errors, you can run the shell script instead from now on

```
cd $EQEMUBUILDDIR/
./server_start_with_login.sh
```

This likely is STILL missing steps, just to warn.
