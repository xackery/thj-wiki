---
title: "Linux Installation Guide"
summary: "Complete Client Installation for Linux"
images: [images/lutris1.png]
weight: 1
---

# Assumptions and Information

*This guide is written from the perspective of a new user on a fresh-ish installation of [Pop!_OS](https://pop.system76.com/) using Wayland, and will use [Lutris](https://www.lutris.net/) to manage the game. Your mileage may vary based on other distros and system environments.*

# Obtain the client files

*Note: It is recommended that you DO NOT use a flatpak distribution of Steam, for a variety of reasons. You should still be able to get this to work, but some details, such as paths, may vary.*

1. Invoke the steam console with `$ steam steam://open/console`. This should open up Steam and provide a console interface.

2. Enter `download_depot 205710 205711 1926608638440811669` into the console. The expected output is:

    ```
    ] download_depot 205710 205711 1926608638440811669
    Downloading depot 205711 (8204 MB) ...
    ```

3. This download may take a significant period of time and will not report any progress.
    * ⚠ **Important**: *Do not start any other downloads or games through Steam while this is in progress, or it will be aborted silently.*
    * Downloaded files can be found at `~/.steam/debian-installation/ubuntu12_32/steamapps/content/app_205710/depot_205711`
    * If you are running some configuration other than the supported one, try locating these files with `find ~/ -type d -name "depot_205711"`
    * You can check download progress with `ls -lhat ~/.steam/debian-installation/ubuntu12_32/steamapps/content/app_205710/depot_205711 | grep total`; this will show the total size of the downloads. This can also be used to check to see if the download was completed successfully (I got 7.4GB when completed).
    * TODO: Provide a directory checksum here (`tar cf - ~/.steam/debian-installation/ubuntu12_32/steamapps/content/app_205710/depot_205711 --sort=name --mtime='1970-01-01' --owner=0 --group=0 --numeric-owner | sha256sum`)

4. Once the files are fully downloaded, copy them to your intended installation directory.
    * For example `mkdir -p ~/Games/eqemu_thj/rof2 && cp -r ~/.steam/debian-installation/ubuntu12_32/steamapps/content/app_205710/depot_205711 ~/Games/eqemu_thj_rof2`

# Prepare your system

1. The version of WINE that ships with Ubuntu & Pop!_OS is positively ancient. Let's update it.

    * Start by adding the upstream WINE ppa and updating apt.

    ```
    sudo dpkg --add-architecture i386
    sudo mkdir -pm755 /etc/apt/keyrings
    sudo wget -O /etc/apt/keyrings/winehq-archive.key https://dl.winehq.org/wine-builds/winehq.key
    sudo wget -NP /etc/apt/sources.list.d/ https://dl.winehq.org/wine-builds/ubuntu/dists/jammy/winehq-jammy.sources
    sudo apt update
    ```

    * Install WINE

    ```
    sudo apt install --install-recommends winehq-stable dxvk winetricks
    ```

    * You should have WINE version 10
    ```
    wine --version
    ```

3. Install Gamescope. This makes EQ play bit nicer, and we'll use it inside Lutris. Enables fake fullscreen and output scaling, similar to Lossless Scaling does on Windows.
    * You have three options to install it
        1. (*Recommended*) The best way I found was [this github repo](https://github.com/akdor1154/gamescope-pkg), which provides a .deb to install.
        2. Flatpak. Install from Flathub. Not recommended because it can be tricky to use.
        3. Build from source. Outside of the scope of this guide, but you probably want to use [this repo](https://github.com/ChimeraOS/gamescope).

4. Winetricks
    * First, update winetricks. This also initializes the wine prefix we will be using.
    ```
    cd ~/Games/eqemu_thj/rof2
    WINEPREFIX=$(pwd)/.wine sudo winetricks --self-update -y
    ```
    * Install the winetricks verbs we need for RoF2 and the patcher
    ```
    WINEPREFIX=$(pwd)/.wine winetricks -q corefonts d3dcompiler_43 d3dcompiler_47 dxvk vcrun2003 vcrun2005
    WINEPREFIX=$(pwd)/.wine winetricks -q vcrun2008 vcrun2010 vcrun6 vcrun6sp6 vcrun2012 vcrun2013 vcrun2015 oleaut32 dotnet48
    ```
    * *Note: At the time of this writing, winetricks has a bug which prevents correct installation of the dotnet48 verb, which breaks the patcher.*

6. Grab the patcher.
    ```
    wget https://github.com/The-Heroes-Journey-EQEMU/eqemupatcher/releases/latest/download/heroesjourneyeq.exe
    ```

7. Test to make sure everything works so far; the patcher should launch and allow you to enter the game. The game should launch and run, but may have some issues still. Notably, you will not be able to stay connected to THJ in particular.
    ```
    WINEPREFIX=$(pwd)/.wine wine heroesjourneyeq.exe
    ```

# Configure Lutris

1. Install Lutris
    * (*Recommended*) Use the software center (Pop!_Shop for Pop!_OS) to install Lutris. **Make sure you use the .deb verison instead of the Flatpak.**

2. Configure Lutris
    * Game Info: This is mostly cosmetic, but make sure you set the Runner to 'Wine'. (*The images can be found in /Resources/ in your game directory after you patch*)
    ![Game Info](/images/lutris1.png)
    * Game Options: The demonstrated paths are consistent with the previous guide, but will obviously not be exactly correct for you. Adjust as needed!
    ![Game Options](/images/lutris2.png)
    * Runner Options:
        * Enabled 'dgvoodoo2'
        * Disabled 'AMD FidelityFX Super Resolution'
        * Disabled 'Enable BattlEye Anti-Cheat`
        * Disabled 'Enable Easy Anti-Cheat'
        * Enable 'Mouse Warp Override'
        * **Add dinput8 DLL Override** -> `n,b`
    ![Runner Options](/images/lutris3.png)
    * System Options:
        * Enable 'Gamescope'
            * Enable 'Relative Mouse Mode'
            * Set 'Output Resolution' (Your Desktop Resolution),
            * Set 'Game Resolution' (The larger this resolution, the smaller the UI will be)
            * Set 'Window Mode' (Dealer's choice)
        * Enable 'Resture resolution on game exit'
        * Enable 'Restore gamma on game exist'
        * Enable 'Use discrete graphics'
    ![Runner Options](/images/lutris4.png)