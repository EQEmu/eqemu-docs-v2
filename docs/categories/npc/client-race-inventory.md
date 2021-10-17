# Client Race Inventory

This work has been completed by [Shendare](https://github.com/Shendare), who has endeavored to create a complete [EverQuest Client Race Inventory](http://www.shendare.com/EQ/Emu/EQRI/) (external links).  It is included (and plagiarized!) here as a reference and to provide the ability for Server Operators to explore models to customize their server.

## Download

You can download the program from the link below:

{% file src="../../.gitbook/assets/eqri.zip" %}
Compiled binary - EQRI.exe
{% endfile %}

## Source

You can download the source files from the link below:

{% file src="../../.gitbook/assets/eqri_source.zip" %}
Source Code
{% endfile %}

## How to use the program

{% hint style="info" %}
Note: Only Titanium, SoF, and RoF2 (Rain of Fear, Steam Release #2) are currently supported!
{% endhint %}

Unzip EQRI.exe into your EQ directory, or create a new folder inside your EQ directory called "EQRI", and unzip the program into that. Technically, it can be unzipped to and run from anywhere, with the following caveats:

If you decided to unzip it outside of your EQ directory, you'll need to create a EQRI.ini file in the folder EQRI.exe is in, telling it where to find your EQ installation, as follows:

As a **relative path**, such as from the EQRI folder inside your EverQuest directory:

`EQPath=..`

Or as a **full path** to your EQ installation, such as:

`EQPath=C:\Program Files\Sony\EverQuest`

Then the program will know where to look for the EQ files if it can't find them in its current directory.

Either way, it's just a matter of running EQRI.exe, either with a double-click or from the command line.

After less than a minute, you'll have a fresh set of references in the EQRI folder detailing the race/model availability as currently configured in your EQ installation (Version_EQRaces.htm and Version_EQZones.htm).

Press any key once the program is finished to exit.

You'll see the important info echoed to the console, while a very detailed log file is generated in the EQRI folder along with the references.

The references are built completely from the EQ source files. If you make a change to a ZoneName_chr.txt file or Resources\GlobalLoad.txt, the changes will be reflected in the results!

NOTE: The program must have write permission in the EQ directory to create the output files in the EQRI folder. In some Windows versions and configurations, this may require you to run the program as an Administrator. Alternatively, you can specify a different write-enabled output folder path in the EQRI.ini file as follows:

`OutPath=C:\OutputFolder`

## How it works

### **Preparation**

1. Reads textual information from eqstr_us.txt and dbstr_us.txt for reference later
2. Reads eqgame.exe's hard-coded list of race/gender model codes by parsing the appropriate block of its machine code
3. Reads eqgame.exe's hard-coded list of zones and their associated information fields (nick, name, expansion, minimum level, etc.) similarly

### **Model Availability Loading**

1. Uses a hard-coded list of Luclin-enabled models (global5\_chr\[2], frog_mount_chr, and globalXXX_chr\[2] for playable races)
2. Parses Resources\GlobalLoad.txt for remaining global model sources
3. Using the Zone list, parses all ZoneNick_chr\[2].s3d files for zone-local models
4. Using the Zone list, parses all ZoneNick_chr.txt files for zone models imported from EQGs, S3Ds, and other zones
5. Looks for any race/gender models that aren't being referenced in a .S3D or ZoneNick_chr.txt file ("orphan models"), and looks for a ModelCode_chr.s3d or ModelCode.eqg file for it

### **Model Information Loading**

For every race/gender model it comes across:

1. Logs the source file it came from and the zone it's available in
2. Parses an S3D file's inner .WLD file to log each model's Texture and Head (AKA HelmTexture) variations
   1. Note: S3D/WLD parsing is monumentally more complicated than it sounds. The WLD format is cryptic, clunky, ambiguous, and difficult to load consistent, reliable information from. Most of the development time on this program was devoted to figuring out how to interpret the WLD data. Mad props to Windcatcher for his wlddoc.pdf and Pascal source code, which made this possible!
3. Parses EQG files and interprets Textures, Heads, and other variations based on texture filenames used
   1. Note: I didn't find file format information for the .MOD and .MDS files that make up EQG race models. Fortunately, the devs were decently consistent with the filename convention for textures used in models, and I was able to extrapolate everything needed for Texture and Head variations from those, with a little interpreting.

### **Reporting**

1. Races: Builds a reference HTML file of all races the client can render, listing model codes, variations, sources, and the zones configured to load each model
2. Zones: Builds a reference HTML file of all zones the client can handle, listing names, expansion requirements, bit flags, minimum level to enter, and what race models are available for rendering
3. Zones: Also lists all global models and sources, along with their variations and what makes them global (Luclin models loaded, Resources\GlobalLoad.txt, etc.)

## Example Output

The output from running the program is too long for GitBook to render in a reasonable time.  You can view sample output at [EverQuest Race Inventory](http://www.shendare.com/EQ/Emu/EQRI/).

### Race Output

The example below demonstrates the output for Race 1 - Human.  

{% tabs %}
{% tab title="Titanium" %}
## Race # 1 - Human, Plural 'Humans' (Playable)

**Model Codes:**

* Male = **HUM**
* Female = **HUF**
* Neutral = N/A

**Model availability:**

* Via Source: globalhum_chr.s3d (Loaded with Luclin models)\
  \- Model **HUM** - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**\

  * **All Zones** (**Global**)
* Via Source: globalhuf_chr.s3d (Loaded with Luclin models)\
  \- Model **HUF** - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**\

  * **All Zones** (**Global**)
* Via Source: global_chr.s3d (Loaded via GlobalLoad.txt)\
  \- Model **HUM** - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**\
  \- Model **HUF** - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**\

  * **All Zones** (**Global**)
{% endtab %}

{% tab title="SoF" %}
## Race # 1 - Human, Plural 'Humans' (Playable)

**Character Creation Description:**

The strength of the Human race lies in its diversity of tought, belief, and profession; though they tend be weaker than many other races. They have no particular specialty with the sword or arcane magic. While that is true, Humans are cunning and possess great ingenuity, which gives them an advantage. The Human mind is sharp enough to adapt to nearly all forms of study and this gives them a great range of options when choosing a profession.

**Model Codes:**

* Male = **HUM**
* Female = **HUF**
* Neutral = N/A

**Model availability:**

* Via Source: globalhum_chr.s3d (Loaded with Luclin models)\
  \- Model **HUM** - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**\

  * **All Zones** (**Global**)
* Via Source: globalhuf_chr.s3d (Loaded with Luclin models)\
  \- Model **HUF** - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**\

  * **All Zones** (**Global**)
* Via Source: global_chr.s3d (Loaded via GlobalLoad.txt)\
  \- Model **HUM** - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**\
  \- Model **HUF** - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**\

  * **All Zones** (**Global**)
{% endtab %}

{% tab title="RoF2" %}
## Race # 1 - Human, Plural 'Humans' (Playable)

**Character Creation Description:**

The strength of the Human race lies in its diversity of thought, belief, and profession; though they tend be weaker than many other races. They have no particular specialty with the sword or arcane magic. While that is true, Humans are cunning and possess great ingenuity, which gives them an advantage. The Human mind is sharp enough to adapt to nearly all forms of study and this gives them a great range of options when choosing a profession.

**Model Codes:**

* Male = **HUM**
* Female = **HUF**
* Neutral = N/A

**Model availability:**

* Via Source: globalhum_chr.s3d (Loaded with Luclin models)\
  \- Model **HUM** - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**\

  * **All Zones** (**Global**)
* Via Source: globalhuf_chr.s3d (Loaded with Luclin models)\
  \- Model **HUF** - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**\

  * **All Zones** (**Global**)
* Via Source: global_chr.s3d (Loaded via GlobalLoad.txt)\
  \- Model **HUM** - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**\
  \- Model **HUF** - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**\

  * **All Zones** (**Global**)
{% endtab %}
{% endtabs %}

### Zone Output

The example below demonstrates the output for zone 0 - Globals:

{% tabs %}
{% tab title="Titanium" %}
## Globally Available Race Models

* **Source:** global7\_chr.s3d (Loaded when Luclin models disabled)
  * **KEF** - Vah Shir Female (Race 130, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **KEM** - Vah Shir Male (Race 130, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalhum_chr.s3d (Loaded with Luclin models)
  * **HUM** - Human Male (Race 1, Gender 0) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**
* **Source:** globalhuf_chr.s3d (Loaded with Luclin models)
  * **HUF** - Human Female (Race 1, Gender 1) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalbam_chr.s3d (Loaded with Luclin models)
  * **BAM** - Barbarian Male (Race 2, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**, Tattoo: **00-08**
* **Source:** globalbaf_chr.s3d (Loaded with Luclin models)
  * **BAF** - Barbarian Female (Race 2, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Tattoo: **00-08**
* **Source:** globalerm_chr.s3d (Loaded with Luclin models)
  * **ERM** - Erudite Male (Race 3, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-04**, Faces: **00-07**, Hair: **00-05**, Beards: **00-05**
* **Source:** globalerf_chr.s3d (Loaded with Luclin models)
  * **ERF** - Erudite Female (Race 3, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-04**, Faces: **00-07**, Hair: **00-08**
* **Source:** globalelm_chr.s3d (Loaded with Luclin models)
  * **ELM** - Wood Elf Male (Race 4, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalelf_chr.s3d (Loaded with Luclin models)
  * **ELF** - Wood Elf Female (Race 4, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalhim_chr.s3d (Loaded with Luclin models)
  * **HIM** - High Elf Male (Race 5, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-03**
* **Source:** globalhif_chr.s3d (Loaded with Luclin models)
  * **HIF** - High Elf Female (Race 5, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globaldam_chr.s3d (Loaded with Luclin models)
  * **DAM** - Dark Elf Male (Race 6, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-03**
* **Source:** globaldaf_chr.s3d (Loaded with Luclin models)
  * **DAF** - Dark Elf Female (Race 6, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalham_chr.s3d (Loaded with Luclin models)
  * **HAM** - Half Elf Male (Race 7, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-03**
* **Source:** globalhaf_chr.s3d (Loaded with Luclin models)
  * **HAF** - Half Elf Female (Race 7, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globaldwm_chr.s3d (Loaded with Luclin models)
  * **DWM** - Dwarf Male (Race 8, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**
* **Source:** globaldwf_chr.s3d (Loaded with Luclin models)
  * **DWF** - Dwarf Female (Race 8, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00 01**
* **Source:** globaltrm_chr.s3d (Loaded with Luclin models)
  * **TRM** - Troll Male (Race 9, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
* **Source:** globaltrf_chr.s3d (Loaded with Luclin models)
  * **TRF** - Troll Female (Race 9, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalogm_chr.s3d (Loaded with Luclin models)
  * **OGM** - Ogre Male (Race 10, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalogf_chr.s3d (Loaded with Luclin models)
  * **OGF** - Ogre Female (Race 10, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalhom_chr.s3d (Loaded with Luclin models)
  * **HOM** - Halfling Male (Race 11, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**
* **Source:** globalhof_chr.s3d (Loaded with Luclin models)
  * **HOF** - Halfling Female (Race 11, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalgnm_chr.s3d (Loaded with Luclin models)
  * **GNM** - Gnome Male (Race 12, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**
* **Source:** globalgnf_chr.s3d (Loaded with Luclin models)
  * **GNF** - Gnome Female (Race 12, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalikm_chr.s3d (Loaded with Luclin models)
  * **IKM** - Iksar Male (Race 128, Gender 0) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalikf_chr.s3d (Loaded with Luclin models)
  * **IKF** - Iksar Female (Race 128, Gender 1) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalkem_chr.s3d (Loaded with Luclin models)
  * **KEM** - Vah Shir Male (Race 130, Gender 0) - Textures: **00-03, 50**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalkef_chr.s3d (Loaded with Luclin models)
  * **KEF** - Vah Shir Female (Race 130, Gender 1) - Textures: **00-03, 50**, Heads: **00-03**, Faces: **00-07**
* **Source:** global5\_chr.s3d (Loaded with Luclin models)
  * **AEL** - Air Elemental (Race 210, Gender 2)
  * **EEL** - Earth Elemental (Race 209, Gender 2)
  * **FEL** - Fire Elemental (Race 212, Gender 2)
  * **HSM** - Horse Male (Race 216, Gender 0) - Textures: **00-03**
  * **WEL** - Water Elemental (Race 211, Gender 2)
* **Source:** frog_mount_chr.s3d (Loaded with Luclin models)
  * **HSF** - Horse Female (Race 216, Gender 1) - Textures: **00-03**
  * **FMT** - Drogmore (Race 348, Gender 2) - Textures: **00-03**
* **Source:** globalfroglok_chr.s3d (Loaded via GlobalLoad.txt)
  * **FRO** - Froglok (Race 26, Gender 2)
  * **FRG** - Froglok (Race 27, Gender 2) - Textures: **00 01**
* **Source:** globalpcfroglok_chr.s3d (Loaded via GlobalLoad.txt)
  * **FRF** - Froglok Female (Race 330, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-09**
  * **FRM** - Froglok Male (Race 330, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-09**
* **Source:** gbn.eqg (Loaded via GlobalLoad.txt)
  * **GBN** - Goblin (Race 433, Gender 2) - Textures: **00-04**
* **Source:** wwf.eqg (Loaded via GlobalLoad.txt)
  * **WWF** - Werewolf (Race 454, Gender 2) - Textures: **00-04**
* **Source:** kbd.eqg (Loaded via GlobalLoad.txt)
  * **KBD** - Kobold (Race 455, Gender 2) - Textures: **00-05**, Heads: **00-05**
* **Source:** fng.eqg (Loaded via GlobalLoad.txt)
  * **FNG** - Sporali (Race 456, Gender 2) - Textures: **00 02 04 06 08 10**, Heads: **00 01, 03**
* **Source:** ork.eqg (Loaded via GlobalLoad.txt)
  * **ORK** - Orc (Race 458, Gender 2) - Textures: **00-09**, Heads: **00-09**
* **Source:** ggy.eqg (Loaded via GlobalLoad.txt)
  * **GGY** - Gargoyle (Race 464, Gender 2) - Textures: **00-03**
* **Source:** eve.eqg (Loaded via GlobalLoad.txt)
  * **EVE** - Evil Eye (Race 469, Gender 2) - Textures: **00-02**
* **Source:** mnr.eqg (Loaded via GlobalLoad.txt)
  * **MNR** - Minotaur (Race 470, Gender 2)
* **Source:** zmm.eqg (Loaded via GlobalLoad.txt)
  * **ZMM** - Zombie Male (Race 471, Gender 0)
* **Source:** fry.eqg (Loaded via GlobalLoad.txt)
  * **FRY** - Fairy (Race 473, Gender 2)
* **Source:** cwr.eqg (Loaded via GlobalLoad.txt)
  * **CWR** - Clockwork Boar (Race 472, Gender 2)
* **Source:** bas.eqg (Loaded via GlobalLoad.txt)
  * **BAS** - Basilisk (Race 436, Gender 2) - Textures: **00 01**
* **Source:** rap_chr.s3d (Loaded via GlobalLoad.txt)
  * **RAP** - Raptor (Race 163, Gender 2) - Textures: **00 01**
* **Source:** skt_chr.s3d (Loaded via GlobalLoad.txt)
  * **SKT** - Skeleton (Race 367, Gender 2) - Textures: **00-04**
* **Source:** global6\_chr.s3d (Loaded via GlobalLoad.txt)
  * **ALL** - Alligator (Race 91, Gender 2) - Textures: **00 01**, Heads: **00 01**
  * **DRK** - Drake (Race 89, Gender 2) - Textures: **00-03**
  * **WOF** - Chokadai (Race 356, Gender 2) - Textures: **00 01**
  * **SKE** - Skeleton (Race 60, Gender 2)
  * **TPN** - Teleport Man (Race 240, Gender 2)
  * **TIG** - Tiger (Race 63, Gender 2)
  * **WOL** - Wolf (Race 42, Gender 2) - Textures: **00-03**
  * **WOE** - Wolf (Race 120, Gender 2) - Textures: **00-03**
* **Source:** global4\_chr.s3d (Loaded via GlobalLoad.txt)
  * **IKS** - Undead Iksar (Race 161, Gender 2) - Heads: **00 01**
  * **SPE** - Spectre (Race 85, Gender 2)
  * **BAF** - Barbarian Female (Race 2, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **BAM** - Barbarian Male (Race 2, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **IKF** - Iksar Female (Race 128, Gender 1) - Textures: **00-04, 10**, Heads: **00-03**, Faces: **00-07**
  * **IKM** - Iksar Male (Race 128, Gender 0) - Textures: **00-04, 10**, Heads: **00-03**, Faces: **00-07**
* **Source:** global_chr.s3d (Loaded via GlobalLoad.txt)
  * **ELE** - Elemental (Race 75, Gender 2) - Textures: **00-03**, Heads: **00 01**
  * **EYE** - Eye (Race 108, Gender 2) - Textures: **00-03**
  * **IVM** - Invisible Man Male (Race 127, Gender 0)
  * **SKE** - Skeleton (Race 60, Gender 2) - Textures: **00 01**
  * **WER** - Werewolf (Race 14, Gender 2)
  * **WOE** - Wolf (Race 120, Gender 2) - Textures: **00 01**
  * **BOAT** - Boat (Race 141, Gender 2)
  * **BAF** - Barbarian Female (Race 2, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **BAM** - Barbarian Male (Race 2, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **DAF** - Dark Elf Female (Race 6, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **DAM** - Dark Elf Male (Race 6, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **DWF** - Dwarf Female (Race 8, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **DWM** - Dwarf Male (Race 8, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **ERF** - Erudite Female (Race 3, Gender 1) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **ERM** - Erudite Male (Race 3, Gender 0) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **GNF** - Gnome Female (Race 12, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-04**, Faces: **00-07**
  * **GNM** - Gnome Male (Race 12, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-04**, Faces: **00-07**
  * **HAF** - Half Elf Female (Race 7, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **HAM** - Half Elf Male (Race 7, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **HOF** - Halfling Female (Race 11, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **HOM** - Halfling Male (Race 11, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **HIF** - High Elf Female (Race 5, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **HIM** - High Elf Male (Race 5, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **HUF** - Human Female (Race 1, Gender 1) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **HUM** - Human Male (Race 1, Gender 0) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **OGF** - Ogre Female (Race 10, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **OGM** - Ogre Male (Race 10, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **TRF** - Troll Female (Race 9, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **TRM** - Troll Male (Race 9, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **ELF** - Wood Elf Female (Race 4, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **ELM** - Wood Elf Male (Race 4, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
* **Source:** global2\_chr.s3d (Loaded via GlobalLoad.txt)
  * **BEA** - Bear (Race 43, Gender 2) - Textures: **00-02**, Heads: **00-02**
  * **BRI** - Bristlebane (Race 153, Gender 2)
  * **CAZ** - Cazic Thule (Race 95, Gender 2) - Heads: **00 01**
  * **ERO** - Erollisi (Race 150, Gender 2)
  * **IMP** - Imp (Race 46, Gender 2)
  * **INN** - Innoruuk (Race 123, Gender 2)
  * **RAL** - Rallos Zek (Race 66, Gender 2) - Textures: **00 01**
  * **SCA** - Scarecrow (Race 82, Gender 2)
  * **SOL** - Solusek Ro (Race 58, Gender 2)
  * **TRI** - Tribunal (Race 151, Gender 2)
  * **TUN** - Tunare (Race 62, Gender 2)
{% endtab %}

{% tab title="SoF" %}
## Globally Available Race Models

* **Source:** global7\_chr.s3d (Loaded when Luclin models disabled)
  * **KEF** - Vah Shir Female (Race 130, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **KEM** - Vah Shir Male (Race 130, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalhum_chr.s3d (Loaded with Luclin models)
  * **HUM** - Human Male (Race 1, Gender 0) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**
* **Source:** globalhuf_chr.s3d (Loaded with Luclin models)
  * **HUF** - Human Female (Race 1, Gender 1) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalbam_chr.s3d (Loaded with Luclin models)
  * **BAM** - Barbarian Male (Race 2, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**, Tattoo: **00-08**
* **Source:** globalbaf_chr.s3d (Loaded with Luclin models)
  * **BAF** - Barbarian Female (Race 2, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Tattoo: **00-08**
* **Source:** globalerm_chr.s3d (Loaded with Luclin models)
  * **ERM** - Erudite Male (Race 3, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-04**, Faces: **00-07**, Hair: **00-05**, Beards: **00-05**
* **Source:** globalerf_chr.s3d (Loaded with Luclin models)
  * **ERF** - Erudite Female (Race 3, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-04**, Faces: **00-07**, Hair: **00-08**
* **Source:** globalelm_chr.s3d (Loaded with Luclin models)
  * **ELM** - Wood Elf Male (Race 4, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalelf_chr.s3d (Loaded with Luclin models)
  * **ELF** - Wood Elf Female (Race 4, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalhim_chr.s3d (Loaded with Luclin models)
  * **HIM** - High Elf Male (Race 5, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-03**
* **Source:** globalhif_chr.s3d (Loaded with Luclin models)
  * **HIF** - High Elf Female (Race 5, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globaldam_chr.s3d (Loaded with Luclin models)
  * **DAM** - Dark Elf Male (Race 6, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-03**
* **Source:** globaldaf_chr.s3d (Loaded with Luclin models)
  * **DAF** - Dark Elf Female (Race 6, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalham_chr.s3d (Loaded with Luclin models)
  * **HAM** - Half Elf Male (Race 7, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-03**
* **Source:** globalhaf_chr.s3d (Loaded with Luclin models)
  * **HAF** - Half Elf Female (Race 7, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globaldwm_chr.s3d (Loaded with Luclin models)
  * **DWM** - Dwarf Male (Race 8, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**
* **Source:** globaldwf_chr.s3d (Loaded with Luclin models)
  * **DWF** - Dwarf Female (Race 8, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00 01**
* **Source:** globaltrm_chr.s3d (Loaded with Luclin models)
  * **TRM** - Troll Male (Race 9, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
* **Source:** globaltrf_chr.s3d (Loaded with Luclin models)
  * **TRF** - Troll Female (Race 9, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalogm_chr.s3d (Loaded with Luclin models)
  * **OGM** - Ogre Male (Race 10, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalogf_chr.s3d (Loaded with Luclin models)
  * **OGF** - Ogre Female (Race 10, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalhom_chr.s3d (Loaded with Luclin models)
  * **HOM** - Halfling Male (Race 11, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**
* **Source:** globalhof_chr.s3d (Loaded with Luclin models)
  * **HOF** - Halfling Female (Race 11, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalgnm_chr.s3d (Loaded with Luclin models)
  * **GNM** - Gnome Male (Race 12, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**
* **Source:** globalgnf_chr.s3d (Loaded with Luclin models)
  * **GNF** - Gnome Female (Race 12, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalikm_chr.s3d (Loaded with Luclin models)
  * **IKM** - Iksar Male (Race 128, Gender 0) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalikf_chr.s3d (Loaded with Luclin models)
  * **IKF** - Iksar Female (Race 128, Gender 1) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalkem_chr.s3d (Loaded with Luclin models)
  * **KEM** - Vah Shir Male (Race 130, Gender 0) - Textures: **00-03, 50**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalkef_chr.s3d (Loaded with Luclin models)
  * **KEF** - Vah Shir Female (Race 130, Gender 1) - Textures: **00-03, 50**, Heads: **00-03**, Faces: **00-07**
* **Source:** global5\_chr.s3d (Loaded with Luclin models)
  * **AEL** - Air Elemental (Race 210, Gender 2)
  * **EEL** - Earth Elemental (Race 209, Gender 2)
  * **FEL** - Fire Elemental (Race 212, Gender 2)
  * **HSM** - Horse Male (Race 216, Gender 0) - Textures: **00-03**
  * **WEL** - Water Elemental (Race 211, Gender 2)
* **Source:** frog_mount_chr.s3d (Loaded with Luclin models)
  * **HSF** - Horse Female (Race 216, Gender 1) - Textures: **00-03**
  * **FMT** - Drogmore (Race 348, Gender 2) - Textures: **00-03**
* **Source:** globalfroglok_chr.s3d (Loaded via GlobalLoad.txt)
  * **FRO** - Froglok (Race 26, Gender 2)
  * **FRG** - Froglok (Race 27, Gender 2) - Textures: **00 01**
* **Source:** globalpcfroglok_chr.s3d (Loaded via GlobalLoad.txt)
  * **FRF** - Froglok Female (Race 330, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-09**
  * **FRM** - Froglok Male (Race 330, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-09**
* **Source:** gbn.eqg (Loaded via GlobalLoad.txt)
  * **GBN** - Goblin (Race 433, Gender 2) - Textures: **00-08**, Heads: **00 06-08**
* **Source:** wwf.eqg (Loaded via GlobalLoad.txt)
  * **WWF** - Werewolf (Race 454, Gender 2) - Textures: **00-04**
* **Source:** kbd.eqg (Loaded via GlobalLoad.txt)
  * **KBD** - Kobold (Race 455, Gender 2) - Textures: **00-05**, Heads: **00-05**
* **Source:** fng.eqg (Loaded via GlobalLoad.txt)
  * **FNG** - Sporali (Race 456, Gender 2) - Textures: **00 02 04 06 08 10 12**, Heads: **00 01, 03**
* **Source:** ork.eqg (Loaded via GlobalLoad.txt)
  * **ORK** - Orc (Race 458, Gender 2) - Textures: **00-09**, Heads: **00-09**
* **Source:** ggy.eqg (Loaded via GlobalLoad.txt)
  * **GGY** - Gargoyle (Race 464, Gender 2) - Textures: **00-03**
* **Source:** eve.eqg (Loaded via GlobalLoad.txt)
  * **EVE** - Evil Eye (Race 469, Gender 2) - Textures: **00-02**
* **Source:** mnr.eqg (Loaded via GlobalLoad.txt)
  * **MNR** - Minotaur (Race 470, Gender 2)
* **Source:** zmf.eqg (Loaded via GlobalLoad.txt)
  * **ZMF** - Zombie Female (Race 471, Gender 1)
* **Source:** zmm.eqg (Loaded via GlobalLoad.txt)
  * **ZMM** - Zombie Male (Race 471, Gender 0)
* **Source:** fry.eqg (Loaded via GlobalLoad.txt)
  * **FRY** - Fairy (Race 473, Gender 2) - Textures: **00 01**, Heads: **00 01**
* **Source:** cwr.eqg (Loaded via GlobalLoad.txt)
  * **CWR** - Clockwork Boar (Race 472, Gender 2)
* **Source:** bas.eqg (Loaded via GlobalLoad.txt)
  * **BAS** - Basilisk (Race 436, Gender 2) - Textures: **00 01**
* **Source:** rap_chr.s3d (Loaded via GlobalLoad.txt)
  * **RAP** - Raptor (Race 163, Gender 2) - Textures: **00 01**
* **Source:** t00.eqg (Loaded via GlobalLoad.txt)
  * **T00** - Bear Trap (Race 503, Gender 2)
* **Source:** t05.eqg (Loaded via GlobalLoad.txt)
  * **T05** - Stone Ring (Race 508, Gender 2)
* **Source:** t07.eqg (Loaded via GlobalLoad.txt)
  * **T07** - Runic Symbol (Race 510, Gender 2)
* **Source:** t09.eqg (Loaded via GlobalLoad.txt)
  * **T09** - Floating Skull (Race 512, Gender 2)
* **Source:** t10.eqg (Loaded via GlobalLoad.txt)
  * **T10** - Spike Trap (Race 513, Gender 2)
* **Source:** t11.eqg (Loaded via GlobalLoad.txt)
  * **T11** - Totem (Race 514, Gender 2)
* **Source:** spt.eqg (Loaded via GlobalLoad.txt)
  * **SPT** - Spectre (Race 485, Gender 2)
* **Source:** dpm_chr.s3d (Loaded via GlobalLoad.txt)
  * **DPM** - Pirate Male (Race 339, Gender 0)
* **Source:** dpf_chr.s3d (Loaded via GlobalLoad.txt)
  * **DPF** - Pirate Female (Race 339, Gender 1)
* **Source:** epm_chr.s3d (Loaded via GlobalLoad.txt)
  * **EPM** - Pirate Male (Race 342, Gender 0)
* **Source:** epf_chr.s3d (Loaded via GlobalLoad.txt)
  * **EPF** - Pirate Female (Race 342, Gender 1)
* **Source:** gpm_chr.s3d (Loaded via GlobalLoad.txt)
  * **GPM** - Pirate Male (Race 338, Gender 0)
* **Source:** gpf_chr.s3d (Loaded via GlobalLoad.txt)
  * **GPF** - Pirate Female (Race 338, Gender 1)
* **Source:** hpf_chr.s3d (Loaded via GlobalLoad.txt)
  * **HPF** - Pirate Female (Race 341, Gender 1)
* **Source:** hpm_chr.s3d (Loaded via GlobalLoad.txt)
  * **HPM** - Pirate Male (Race 341, Gender 0)
* **Source:** opf_chr.s3d (Loaded via GlobalLoad.txt)
  * **OPF** - Pirate Female (Race 340, Gender 1)
* **Source:** opm_chr.s3d (Loaded via GlobalLoad.txt)
  * **OPM** - Pirate Male (Race 340, Gender 0)
* **Source:** tbm_chr.s3d (Loaded via GlobalLoad.txt)
  * **TBM** - Troll Male (Race 331, Gender 0)
* **Source:** tbf_chr.s3d (Loaded via GlobalLoad.txt)
  * **TBF** - Troll Female (Race 331, Gender 1)
* **Source:** wor.eqg (Loaded via GlobalLoad.txt)
  * **WOR** - Worg (Race 580, Gender 2) - Heads: **00-02**
* **Source:** skt_chr.s3d (Loaded via GlobalLoad.txt)
  * **SKT** - Skeleton (Race 367, Gender 2) - Textures: **00-04**
* **Source:** global6\_chr.s3d (Loaded via GlobalLoad.txt)
  * **ALL** - Alligator (Race 91, Gender 2) - Textures: **00 01**, Heads: **00 01**
  * **DRK** - Drake (Race 89, Gender 2) - Textures: **00-03**
  * **WOF** - Chokadai (Race 356, Gender 2) - Textures: **00 01**
  * **SKE** - Skeleton (Race 60, Gender 2)
  * **TPN** - Teleport Man (Race 240, Gender 2)
  * **TIG** - Tiger (Race 63, Gender 2)
  * **WOL** - Wolf (Race 42, Gender 2) - Textures: **00-03**
  * **WOE** - Wolf (Race 120, Gender 2) - Textures: **00-03**
* **Source:** global4\_chr.s3d (Loaded via GlobalLoad.txt)
  * **IKS** - Undead Iksar (Race 161, Gender 2) - Heads: **00 01**
  * **SPE** - Spectre (Race 85, Gender 2)
  * **BAF** - Barbarian Female (Race 2, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **BAM** - Barbarian Male (Race 2, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **IKF** - Iksar Female (Race 128, Gender 1) - Textures: **00-04, 10**, Heads: **00-03**, Faces: **00-07**
  * **IKM** - Iksar Male (Race 128, Gender 0) - Textures: **00-04, 10**, Heads: **00-03**, Faces: **00-07**
* **Source:** global_chr.s3d (Loaded via GlobalLoad.txt)
  * **ELE** - Elemental (Race 75, Gender 2) - Textures: **00-03**, Heads: **00 01**
  * **EYE** - Eye (Race 108, Gender 2) - Textures: **00-03**
  * **IVM** - Invisible Man Male (Race 127, Gender 0)
  * **SKE** - Skeleton (Race 60, Gender 2) - Textures: **00 01**
  * **WER** - Werewolf (Race 14, Gender 2)
  * **WOE** - Wolf (Race 120, Gender 2) - Textures: **00 01**
  * **BOAT** - Boat (Race 141, Gender 2)
  * **BAF** - Barbarian Female (Race 2, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **BAM** - Barbarian Male (Race 2, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **DAF** - Dark Elf Female (Race 6, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **DAM** - Dark Elf Male (Race 6, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **DWF** - Dwarf Female (Race 8, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **DWM** - Dwarf Male (Race 8, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **ERF** - Erudite Female (Race 3, Gender 1) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **ERM** - Erudite Male (Race 3, Gender 0) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **GNF** - Gnome Female (Race 12, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-04**, Faces: **00-07**
  * **GNM** - Gnome Male (Race 12, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-04**, Faces: **00-07**
  * **HAF** - Half Elf Female (Race 7, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **HAM** - Half Elf Male (Race 7, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **HOF** - Halfling Female (Race 11, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **HOM** - Halfling Male (Race 11, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **HIF** - High Elf Female (Race 5, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **HIM** - High Elf Male (Race 5, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **HUF** - Human Female (Race 1, Gender 1) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **HUM** - Human Male (Race 1, Gender 0) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **OGF** - Ogre Female (Race 10, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **OGM** - Ogre Male (Race 10, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **TRF** - Troll Female (Race 9, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **TRM** - Troll Male (Race 9, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **ELF** - Wood Elf Female (Race 4, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **ELM** - Wood Elf Male (Race 4, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
* **Source:** dkm.eqg (Loaded via GlobalLoad.txt)
  * **DKM** - Drakkin Male (Race 522, Gender 0) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-06**, Hair: **00-08**, Beards: **00-11**, Detail: **00-07**, Tattoo: **00-07**, Heritage: **00-07**
* **Source:** dkf.eqg (Loaded via GlobalLoad.txt)
  * **DKF** - Drakkin Female (Race 522, Gender 1) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-06**, Hair: **00-07**, Beards: **00-03**, Detail: **00-07**, Tattoo: **00-07**, Heritage: **00-07**
* **Source:** unm.eqg (Loaded via GlobalLoad.txt)
  * **UNM** - Nightmare/Unicorn (Race 519, Gender 2) - Textures: **00 01**, Heads: **00 01**
* **Source:** hrs.eqg (Loaded via GlobalLoad.txt)
  * **HRS** - Horse (Race 518, Gender 2) - Textures: **00-02**, Heads: **00 03 06**
* **Source:** globalgdb.eqg (Loaded via GlobalLoad.txt)
  * **G00** - Banner (Race 553, Gender 2)
  * **G01** - Banner (Race 554, Gender 2)
  * **G02** - Banner (Race 555, Gender 2)
  * **G03** - Banner (Race 556, Gender 2)
  * **G04** - Banner (Race 557, Gender 2)
* **Source:** glm.eqg (Loaded via GlobalLoad.txt)
  * **GLM** - Golem (Race 374, Gender 2) - Textures: **00-04**
* **Source:** mur.eqg (Loaded via GlobalLoad.txt)
  * **MUR** - Lightning Warrior (Race 407, Gender 2)
* **Source:** dsg.eqg (Loaded via GlobalLoad.txt)
  * **DSG** - Bazu (Race 409, Gender 2)
* **Source:** scu.eqg (Loaded via GlobalLoad.txt)
  * **SCU** - Pyrilen (Race 411, Gender 2)
* **Source:** frd.eqg (Loaded via GlobalLoad.txt)
  * **FRD** - Gelidran (Race 417, Gender 2)
* **Source:** ddv.eqg (Loaded via GlobalLoad.txt)
  * **DDV** - Girplan (Race 419, Gender 2) - Textures: **00-03**
* **Source:** shl.eqg (Loaded via GlobalLoad.txt)
  * **SHL** - Shiliskin (Race 467, Gender 2) - Textures: **00-03**, Heads: **00-03**
* **Source:** bsg.eqg (Loaded via GlobalLoad.txt)
  * **BSG** - Banshee Female (Race 488, Gender 1)
* **Source:** srn.eqg (Loaded via GlobalLoad.txt)
  * **SRN** - Scrykin (Race 495, Gender 2) - Textures: **00-03**, Heads: **00-03**
* **Source:** bxi.eqg (Loaded via GlobalLoad.txt)
  * **BXI** - Bixie (Race 520, Gender 2) - Textures: **00-02**, Heads: **00-02**
* **Source:** drg.eqg (Loaded via GlobalLoad.txt)
  * **DRG** - Dragon (Race 530, Gender 2) - Textures: **00-05**, Heads: **00-05**
* **Source:** mch.eqg (Loaded via GlobalLoad.txt)
  * **MCH** - Chimera (Race 582, Gender 2)
* **Source:** mki.eqg (Loaded via GlobalLoad.txt)
  * **MKI** - Kirin (Race 583, Gender 2)
* **Source:** mpu.eqg (Loaded via GlobalLoad.txt)
  * **MPU** - Puma (Race 584, Gender 2) - Textures: **00 01**
* **Source:** g05.eqg (Loaded via GlobalLoad.txt)
  * **G05** - Banner (Race 586, Gender 2)
* **Source:** i10.eqg (Loaded via GlobalLoad.txt)
  * **I10** - Campfire (Race 567, Gender 2)
* **Source:** global2\_chr.s3d (Loaded via GlobalLoad.txt)
  * **BEA** - Bear (Race 43, Gender 2) - Textures: **00-02**, Heads: **00-02**
  * **BRI** - Bristlebane (Race 153, Gender 2)
  * **CAZ** - Cazic Thule (Race 95, Gender 2) - Heads: **00 01**
  * **ERO** - Erollisi (Race 150, Gender 2)
  * **IMP** - Imp (Race 46, Gender 2)
  * **INN** - Innoruuk (Race 123, Gender 2)
  * **RAL** - Rallos Zek (Race 66, Gender 2) - Textures: **00 01**
  * **SCA** - Scarecrow (Race 82, Gender 2)
  * **SOL** - Solusek Ro (Race 58, Gender 2)
  * **TRI** - Tribunal (Race 151, Gender 2)
  * **TUN** - Tunare (Race 62, Gender 2)
{% endtab %}

{% tab title="RoF2" %}
## Globally Available Race Models

* **Source:** global7\_chr.s3d (Loaded when Luclin models disabled)
  * **KEF** - Vah Shir Female (Race 130, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **KEM** - Vah Shir Male (Race 130, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalhum_chr.s3d (Loaded with Luclin models)
  * **HUM** - Human Male (Race 1, Gender 0) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**
* **Source:** globalhuf_chr.s3d (Loaded with Luclin models)
  * **HUF** - Human Female (Race 1, Gender 1) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalbam_chr.s3d (Loaded with Luclin models)
  * **BAM** - Barbarian Male (Race 2, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**, Tattoo: **00-08**
* **Source:** globalbaf_chr.s3d (Loaded with Luclin models)
  * **BAF** - Barbarian Female (Race 2, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Tattoo: **00-08**
* **Source:** globalerm_chr.s3d (Loaded with Luclin models)
  * **ERM** - Erudite Male (Race 3, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-04**, Faces: **00-07**, Hair: **00-05**, Beards: **00-05**
* **Source:** globalerf_chr.s3d (Loaded with Luclin models)
  * **ERF** - Erudite Female (Race 3, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-04**, Faces: **00-07**, Hair: **00-08**
* **Source:** globalelm_chr.s3d (Loaded with Luclin models)
  * **ELM** - Wood Elf Male (Race 4, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalelf_chr.s3d (Loaded with Luclin models)
  * **ELF** - Wood Elf Female (Race 4, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalhim_chr.s3d (Loaded with Luclin models)
  * **HIM** - High Elf Male (Race 5, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-03**
* **Source:** globalhif_chr.s3d (Loaded with Luclin models)
  * **HIF** - High Elf Female (Race 5, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globaldam_chr.s3d (Loaded with Luclin models)
  * **DAM** - Dark Elf Male (Race 6, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-03**
* **Source:** globaldaf_chr.s3d (Loaded with Luclin models)
  * **DAF** - Dark Elf Female (Race 6, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalham_chr.s3d (Loaded with Luclin models)
  * **HAM** - Half Elf Male (Race 7, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-03**
* **Source:** globalhaf_chr.s3d (Loaded with Luclin models)
  * **HAF** - Half Elf Female (Race 7, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globaldwm_chr.s3d (Loaded with Luclin models)
  * **DWM** - Dwarf Male (Race 8, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**
* **Source:** globaldwf_chr.s3d (Loaded with Luclin models)
  * **DWF** - Dwarf Female (Race 8, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00 01**
* **Source:** globaltrm_chr.s3d (Loaded with Luclin models)
  * **TRM** - Troll Male (Race 9, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
* **Source:** globaltrf_chr.s3d (Loaded with Luclin models)
  * **TRF** - Troll Female (Race 9, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalogm_chr.s3d (Loaded with Luclin models)
  * **OGM** - Ogre Male (Race 10, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalogf_chr.s3d (Loaded with Luclin models)
  * **OGF** - Ogre Female (Race 10, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalhom_chr.s3d (Loaded with Luclin models)
  * **HOM** - Halfling Male (Race 11, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**
* **Source:** globalhof_chr.s3d (Loaded with Luclin models)
  * **HOF** - Halfling Female (Race 11, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalgnm_chr.s3d (Loaded with Luclin models)
  * **GNM** - Gnome Male (Race 12, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**, Beards: **00-05**
* **Source:** globalgnf_chr.s3d (Loaded with Luclin models)
  * **GNF** - Gnome Female (Race 12, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**, Hair: **00-03**
* **Source:** globalikm_chr.s3d (Loaded with Luclin models)
  * **IKM** - Iksar Male (Race 128, Gender 0) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalikf_chr.s3d (Loaded with Luclin models)
  * **IKF** - Iksar Female (Race 128, Gender 1) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalkem_chr.s3d (Loaded with Luclin models)
  * **KEM** - Vah Shir Male (Race 130, Gender 0) - Textures: **00-03, 50**, Heads: **00-03**, Faces: **00-07**
* **Source:** globalkef_chr.s3d (Loaded with Luclin models)
  * **KEF** - Vah Shir Female (Race 130, Gender 1) - Textures: **00-03, 50**, Heads: **00-03**, Faces: **00-07**
* **Source:** global5\_chr.s3d (Loaded with Luclin models)
  * **AEL** - Air Elemental (Race 210, Gender 2)
  * **EEL** - Earth Elemental (Race 209, Gender 2)
  * **FEL** - Fire Elemental (Race 212, Gender 2)
  * **HSM** - Horse Male (Race 216, Gender 0) - Textures: **00-03**
  * **WEL** - Water Elemental (Race 211, Gender 2)
* **Source:** frog_mount_chr.s3d (Loaded with Luclin models)
  * **HSF** - Horse Female (Race 216, Gender 1) - Textures: **00-03**
  * **FMT** - Drogmore (Race 348, Gender 2) - Textures: **00-03**
* **Source:** globalfroglok_chr.s3d (Loaded via GlobalLoad.txt)
  * **FRO** - Froglok (Race 26, Gender 2)
  * **FRG** - Froglok (Race 27, Gender 2) - Textures: **00 01**
* **Source:** globalpcfroglok_chr.s3d (Loaded via GlobalLoad.txt)
  * **FRF** - Froglok Female (Race 330, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-09**
  * **FRM** - Froglok Male (Race 330, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-09**
* **Source:** frogequip.s3d (Loaded via GlobalLoad.txt)
  * **IT4840** - Unrecognized Race Model
  * **IT4841** - Unrecognized Race Model
  * **IT4842** - Unrecognized Race Model
  * **IT4843** - Unrecognized Race Model
  * **IT4844** - Unrecognized Race Model
  * **IT4845** - Unrecognized Race Model
  * **IT4846** - Unrecognized Race Model
  * **IT4870** - Unrecognized Race Model
  * **IT4871** - Unrecognized Race Model
  * **IT4872** - Unrecognized Race Model
  * **IT4873** - Unrecognized Race Model
  * **IT4874** - Unrecognized Race Model
  * **IT4875** - Unrecognized Race Model
  * **IT4876** - Unrecognized Race Model
  * **IT5841** - Unrecognized Race Model
  * **IT5842** - Unrecognized Race Model
  * **IT5843** - Unrecognized Race Model
  * **IT5871** - Unrecognized Race Model
  * **IT5872** - Unrecognized Race Model
  * **IT5873** - Unrecognized Race Model
  * **IT7843** - Unrecognized Race Model
  * **IT7845** - Unrecognized Race Model
  * **IT7847** - Unrecognized Race Model
  * **IT7873** - Unrecognized Race Model
  * **IT7875** - Unrecognized Race Model
  * **IT7877** - Unrecognized Race Model
  * **IT8846** - Unrecognized Race Model
  * **IT8847** - Unrecognized Race Model
  * **IT8876** - Unrecognized Race Model
  * **IT8877** - Unrecognized Race Model
* **Source:** global_obj.s3d (Loaded via GlobalLoad.txt)
  * **POKTELE500** - Unrecognized Race Model
* **Source:** gequip.s3d (Loaded via GlobalLoad.txt)
  * **BBBOARD** - Unrecognized Race Model
  * **GENA00** - Unrecognized Race Model
  * **GENA10** - Unrecognized Race Model
  * **GENA20** - Unrecognized Race Model
  * **GENA30** - Unrecognized Race Model
  * **GENA40** - Unrecognized Race Model
  * **GENB00** - Unrecognized Race Model
  * **GENB10** - Unrecognized Race Model
  * **GENB20** - Unrecognized Race Model
  * **GENB30** - Unrecognized Race Model
  * **GENB40** - Unrecognized Race Model
  * **GENC00** - Unrecognized Race Model
  * **GENC10** - Unrecognized Race Model
  * **GENC20** - Unrecognized Race Model
  * **GENC30** - Unrecognized Race Model
  * **GENC40** - Unrecognized Race Model
  * **GEND00** - Unrecognized Race Model
  * **GEND10** - Unrecognized Race Model
  * **GEND20** - Unrecognized Race Model
  * **GEND30** - Unrecognized Race Model
  * **GEND40** - Unrecognized Race Model
  * **GENE01** - Unrecognized Race Model
  * **GENF00** - Unrecognized Race Model
  * **GENG00** - Unrecognized Race Model
  * **GENH00** - Unrecognized Race Model
  * **GENI00** - Unrecognized Race Model
  * **GENJ00** - Unrecognized Race Model
  * **GENJ10** - Unrecognized Race Model
  * **GENK00** - Unrecognized Race Model
  * **GENK10** - Unrecognized Race Model
  * **GENL00** - Unrecognized Race Model
  * **GENM00** - Unrecognized Race Model
  * **GENN00** - Unrecognized Race Model
  * **GENO00** - Unrecognized Race Model
  * **GENP00** - Unrecognized Race Model
  * **GENQ00** - Unrecognized Race Model
  * **GENR00** - Unrecognized Race Model
  * **GENR10** - Unrecognized Race Model
  * **GENS00** - Unrecognized Race Model
  * **GENS10** - Unrecognized Race Model
  * **GENT00** - Unrecognized Race Model
  * **GENU00** - Unrecognized Race Model
  * **GENU10** - Unrecognized Race Model
  * **GENV00** - Unrecognized Race Model
  * **GENW00** - Unrecognized Race Model
  * **GENW10** - Unrecognized Race Model
  * **GENX00** - Unrecognized Race Model
  * **GENY10** - Unrecognized Race Model
  * **GENY20** - Unrecognized Race Model
  * **GENY30** - Unrecognized Race Model
  * **GENY40** - Unrecognized Race Model
  * **GENZ00** - Unrecognized Race Model
  * **GENZ01** - Unrecognized Race Model
  * **IT1** - Unrecognized Race Model
  * **IT2** - Unrecognized Race Model
  * **IT3** - Unrecognized Race Model
  * **IT4** - Unrecognized Race Model
  * **IT5** - Unrecognized Race Model
  * **IT6** - Unrecognized Race Model
  * **IT7** - Unrecognized Race Model
  * **IT8** - Unrecognized Race Model
  * **IT9** - Unrecognized Race Model
  * **IT10** - Unrecognized Race Model
  * **IT13** - Unrecognized Race Model
  * **IT14** - Unrecognized Race Model
  * **IT15** - Unrecognized Race Model
  * **IT16** - Unrecognized Race Model
  * **IT17** - Unrecognized Race Model
  * **IT18** - Unrecognized Race Model
  * **IT19** - Unrecognized Race Model
  * **IT20** - Unrecognized Race Model
  * **IT21** - Unrecognized Race Model
  * **IT22** - Unrecognized Race Model
  * **IT23** - Unrecognized Race Model
  * **IT24** - Unrecognized Race Model
  * **IT25** - Unrecognized Race Model
  * **IT26** - Unrecognized Race Model
  * **IT27** - Unrecognized Race Model
  * **IT28** - Unrecognized Race Model
  * **IT29** - Unrecognized Race Model
  * **IT30** - Unrecognized Race Model
  * **IT31** - Unrecognized Race Model
  * **IT32** - Unrecognized Race Model
  * **IT33** - Unrecognized Race Model
  * **IT34** - Unrecognized Race Model
  * **IT35** - Unrecognized Race Model
  * **IT36** - Unrecognized Race Model
  * **IT37** - Unrecognized Race Model
  * **IT38** - Unrecognized Race Model
  * **IT39** - Unrecognized Race Model
  * **IT40** - Unrecognized Race Model
  * **IT41** - Unrecognized Race Model
  * **IT42** - Unrecognized Race Model
  * **IT43** - Unrecognized Race Model
  * **IT44** - Unrecognized Race Model
  * **IT45** - Unrecognized Race Model
  * **IT46** - Unrecognized Race Model
  * **IT47** - Unrecognized Race Model
  * **IT48** - Unrecognized Race Model
  * **IT49** - Unrecognized Race Model
  * **IT50** - Unrecognized Race Model
  * **IT51** - Unrecognized Race Model
  * **IT52** - Unrecognized Race Model
  * **IT53** - Unrecognized Race Model
  * **IT54** - Unrecognized Race Model
  * **IT55** - Unrecognized Race Model
  * **IT56** - Unrecognized Race Model
  * **IT57** - Unrecognized Race Model
  * **IT58** - Unrecognized Race Model
  * **IT59** - Unrecognized Race Model
  * **IT60** - Unrecognized Race Model
  * **IT61** - Unrecognized Race Model
  * **IT62** - Unrecognized Race Model
  * **IT63** - Unrecognized Race Model
  * **IT64** - Unrecognized Race Model
  * **IT65** - Unrecognized Race Model
  * **IT66** - Unrecognized Race Model
  * **IT67** - Unrecognized Race Model
  * **IT68** - Unrecognized Race Model
  * **IT69** - Unrecognized Race Model
  * **IT70** - Unrecognized Race Model
  * **IT71** - Unrecognized Race Model
  * **IT72** - Unrecognized Race Model
  * **IT73** - Unrecognized Race Model
  * **IT74** - Unrecognized Race Model
  * **IT75** - Unrecognized Race Model
  * **IT76** - Unrecognized Race Model
  * **IT78** - Unrecognized Race Model
  * **IT79** - Unrecognized Race Model
  * **IT80** - Unrecognized Race Model
  * **IT81** - Unrecognized Race Model
  * **IT82** - Unrecognized Race Model
  * **IT83** - Unrecognized Race Model
  * **IT84** - Unrecognized Race Model
  * **IT85** - Unrecognized Race Model
  * **IT86** - Unrecognized Race Model
  * **IT87** - Unrecognized Race Model
  * **IT88** - Unrecognized Race Model
  * **IT89** - Unrecognized Race Model
  * **IT90** - Unrecognized Race Model
  * **IT91** - Unrecognized Race Model
  * **IT92** - Unrecognized Race Model
  * **IT93** - Unrecognized Race Model
  * **IT94** - Unrecognized Race Model
  * **IT95** - Unrecognized Race Model
  * **IT96** - Unrecognized Race Model
  * **IT97** - Unrecognized Race Model
  * **IT98** - Unrecognized Race Model
  * **IT99** - Unrecognized Race Model
  * **IT100** - Unrecognized Race Model
  * **IT101** - Unrecognized Race Model
  * **IT102** - Unrecognized Race Model
  * **IT103** - Unrecognized Race Model
  * **IT104** - Unrecognized Race Model
  * **IT105** - Unrecognized Race Model
  * **IT106** - Unrecognized Race Model
  * **IT107** - Unrecognized Race Model
  * **IT108** - Unrecognized Race Model
  * **IT109** - Unrecognized Race Model
  * **IT110** - Unrecognized Race Model
  * **IT111** - Unrecognized Race Model
  * **IT112** - Unrecognized Race Model
  * **IT113** - Unrecognized Race Model
  * **IT115** - Unrecognized Race Model
  * **IT117** - Unrecognized Race Model
  * **IT118** - Unrecognized Race Model
  * **IT119** - Unrecognized Race Model
  * **IT120** - Unrecognized Race Model
  * **IT121** - Unrecognized Race Model
  * **IT122** - Unrecognized Race Model
  * **IT123** - Unrecognized Race Model
  * **IT124** - Unrecognized Race Model
  * **IT125** - Unrecognized Race Model
  * **IT126** - Unrecognized Race Model
  * **IT127** - Unrecognized Race Model
  * **IT128** - Unrecognized Race Model
  * **IT129** - Unrecognized Race Model
  * **IT130** - Unrecognized Race Model
  * **IT131** - Unrecognized Race Model
  * **IT132** - Unrecognized Race Model
  * **IT133** - Unrecognized Race Model
  * **IT134** - Unrecognized Race Model
  * **IT135** - Unrecognized Race Model
  * **IT136** - Unrecognized Race Model
  * **IT137** - Unrecognized Race Model
  * **IT138** - Unrecognized Race Model
  * **IT139** - Unrecognized Race Model
  * **IT140** - Unrecognized Race Model
  * **IT141** - Unrecognized Race Model
  * **IT142** - Unrecognized Race Model
  * **IT145** - Unrecognized Race Model
  * **IT146** - Unrecognized Race Model
  * **IT148** - Unrecognized Race Model
  * **IT149** - Unrecognized Race Model
  * **IT150** - Unrecognized Race Model
  * **IT151** - Unrecognized Race Model
  * **IT153** - Unrecognized Race Model
  * **IT154** - Unrecognized Race Model
  * **IT155** - Unrecognized Race Model
  * **IT156** - Unrecognized Race Model
* **Source:** gequip8.s3d (Loaded via GlobalLoad.txt)
  * **IT10524** - Unrecognized Race Model
  * **IT10733** - Unrecognized Race Model
* **Source:** gequip2.s3d (Loaded via GlobalLoad.txt)
  * **IT11** - Unrecognized Race Model
  * **IT12** - Unrecognized Race Model
  * **IT161** - Unrecognized Race Model
  * **IT162** - Unrecognized Race Model
  * **IT163** - Unrecognized Race Model
  * **IT164** - Unrecognized Race Model
  * **IT165** - Unrecognized Race Model
  * **IT166** - Unrecognized Race Model
  * **IT167** - Unrecognized Race Model
  * **IT168** - Unrecognized Race Model
  * **IT169** - Unrecognized Race Model
  * **IT170** - Unrecognized Race Model
  * **IT171** - Unrecognized Race Model
  * **IT172** - Unrecognized Race Model
  * **IT173** - Unrecognized Race Model
  * **IT174** - Unrecognized Race Model
  * **IT175** - Unrecognized Race Model
  * **IT176** - Unrecognized Race Model
  * **IT177** - Unrecognized Race Model
  * **IT178** - Unrecognized Race Model
  * **IT179** - Unrecognized Race Model
  * **IT180** - Unrecognized Race Model
  * **IT181** - Unrecognized Race Model
  * **IT182** - Unrecognized Race Model
  * **IT183** - Unrecognized Race Model
  * **IT184** - Unrecognized Race Model
  * **IT185** - Unrecognized Race Model
  * **IT186** - Unrecognized Race Model
  * **IT187** - Unrecognized Race Model
  * **IT188** - Unrecognized Race Model
  * **IT189** - Unrecognized Race Model
  * **IT190** - Unrecognized Race Model
  * **IT191** - Unrecognized Race Model
  * **IT192** - Unrecognized Race Model
  * **IT193** - Unrecognized Race Model
  * **IT194** - Unrecognized Race Model
  * **IT195** - Unrecognized Race Model
  * **IT196** - Unrecognized Race Model
  * **IT197** - Unrecognized Race Model
  * **IT198** - Unrecognized Race Model
  * **IT228** - Unrecognized Race Model
  * **IT250** - Unrecognized Race Model
  * **IT300** - Unrecognized Race Model
  * **IT301** - Unrecognized Race Model
  * **IT308** - Unrecognized Race Model
  * **IT530** - Unrecognized Race Model
  * **IT535** - Unrecognized Race Model
  * **IT536** - Unrecognized Race Model
  * **IT537** - Unrecognized Race Model
  * **IT540** - Unrecognized Race Model
  * **IT545** - Unrecognized Race Model
  * **IT550** - Unrecognized Race Model
  * **IT555** - Unrecognized Race Model
  * **IT556** - Unrecognized Race Model
  * **IT557** - Unrecognized Race Model
  * **IT560** - Unrecognized Race Model
  * **IT561** - Unrecognized Race Model
  * **IT565** - Unrecognized Race Model
  * **IT566** - Unrecognized Race Model
  * **IT570** - Unrecognized Race Model
  * **IT575** - Unrecognized Race Model
  * **IT580** - Unrecognized Race Model
  * **IT585** - Unrecognized Race Model
  * **IT590** - Unrecognized Race Model
  * **IT595** - Unrecognized Race Model
  * **IT600** - Unrecognized Race Model
  * **IT605** - Unrecognized Race Model
  * **IT610** - Unrecognized Race Model
  * **IT615** - Unrecognized Race Model
  * **IT620** - Unrecognized Race Model
  * **IT625** - Unrecognized Race Model
  * **IT626** - Unrecognized Race Model
  * **IT627** - Unrecognized Race Model
  * **IT628** - Unrecognized Race Model
  * **IT629** - Unrecognized Race Model
  * **IT630** - Unrecognized Race Model
  * **IT635** - Unrecognized Race Model
  * **IT640** - Unrecognized Race Model
  * **IT641** - Unrecognized Race Model
  * **IT645** - Unrecognized Race Model
  * **IT646** - Unrecognized Race Model
  * **IT650** - Unrecognized Race Model
  * **IT655** - Unrecognized Race Model
  * **IT656** - Unrecognized Race Model
* **Source:** gequip5.s3d (Loaded via GlobalLoad.txt)
  * **IT10400** - Unrecognized Race Model
  * **IT10401** - Unrecognized Race Model
  * **IT10402** - Unrecognized Race Model
  * **IT10403** - Unrecognized Race Model
  * **IT10404** - Unrecognized Race Model
  * **IT10405** - Unrecognized Race Model
  * **IT10406** - Unrecognized Race Model
  * **IT10407** - Unrecognized Race Model
  * **IT10408** - Unrecognized Race Model
  * **IT10409** - Unrecognized Race Model
  * **IT10410** - Unrecognized Race Model
  * **IT10411** - Unrecognized Race Model
  * **IT10412** - Unrecognized Race Model
  * **IT10413** - Unrecognized Race Model
  * **IT10501** - Unrecognized Race Model
  * **IT10502** - Unrecognized Race Model
  * **IT10503** - Unrecognized Race Model
  * **IT10504** - Unrecognized Race Model
  * **IT10505** - Unrecognized Race Model
  * **IT10506** - Unrecognized Race Model
  * **IT10507** - Unrecognized Race Model
  * **IT10508** - Unrecognized Race Model
  * **IT10509** - Unrecognized Race Model
  * **IT10510** - Unrecognized Race Model
  * **IT10511** - Unrecognized Race Model
  * **IT10512** - Unrecognized Race Model
  * **IT10513** - Unrecognized Race Model
  * **IT10514** - Unrecognized Race Model
  * **IT10515** - Unrecognized Race Model
  * **IT10516** - Unrecognized Race Model
  * **IT10517** - Unrecognized Race Model
  * **IT10518** - Unrecognized Race Model
  * **IT10519** - Unrecognized Race Model
  * **IT10520** - Unrecognized Race Model
  * **IT10521** - Unrecognized Race Model
  * **IT10522** - Unrecognized Race Model
  * **IT10523** - Unrecognized Race Model
  * **IT10524** - Unrecognized Race Model
  * **IT10525** - Unrecognized Race Model
  * **IT10526** - Unrecognized Race Model
  * **IT10527** - Unrecognized Race Model
  * **IT10528** - Unrecognized Race Model
  * **IT10530** - Unrecognized Race Model
  * **IT10531** - Unrecognized Race Model
  * **IT10532** - Unrecognized Race Model
  * **IT10533** - Unrecognized Race Model
  * **IT10534** - Unrecognized Race Model
  * **IT10535** - Unrecognized Race Model
  * **IT10536** - Unrecognized Race Model
  * **IT10537** - Unrecognized Race Model
  * **IT10538** - Unrecognized Race Model
  * **IT10539** - Unrecognized Race Model
  * **IT10540** - Unrecognized Race Model
  * **IT10541** - Unrecognized Race Model
  * **IT10542** - Unrecognized Race Model
  * **IT10543** - Unrecognized Race Model
  * **IT10544** - Unrecognized Race Model
  * **IT10545** - Unrecognized Race Model
  * **IT10600** - Unrecognized Race Model
  * **IT10601** - Unrecognized Race Model
  * **IT10602** - Unrecognized Race Model
  * **IT10603** - Unrecognized Race Model
  * **IT10604** - Unrecognized Race Model
  * **IT10605** - Unrecognized Race Model
  * **IT10606** - Unrecognized Race Model
  * **IT10607** - Unrecognized Race Model
  * **IT10608** - Unrecognized Race Model
  * **IT10609** - Unrecognized Race Model
  * **IT10610** - Unrecognized Race Model
  * **IT10611** - Unrecognized Race Model
  * **IT10612** - Unrecognized Race Model
  * **IT10613** - Unrecognized Race Model
  * **IT10614** - Unrecognized Race Model
  * **IT10615** - Unrecognized Race Model
  * **IT10616** - Unrecognized Race Model
  * **IT10617** - Unrecognized Race Model
  * **IT10618** - Unrecognized Race Model
  * **IT10619** - Unrecognized Race Model
  * **IT10620** - Unrecognized Race Model
  * **IT10621** - Unrecognized Race Model
  * **IT10622** - Unrecognized Race Model
  * **IT10623** - Unrecognized Race Model
  * **IT10624** - Unrecognized Race Model
  * **IT10625** - Unrecognized Race Model
  * **IT10626** - Unrecognized Race Model
  * **IT10627** - Unrecognized Race Model
  * **IT10628** - Unrecognized Race Model
  * **IT10629** - Unrecognized Race Model
  * **IT10630** - Unrecognized Race Model
  * **IT10631** - Unrecognized Race Model
  * **IT10632** - Unrecognized Race Model
  * **IT10633** - Unrecognized Race Model
  * **IT10634** - Unrecognized Race Model
  * **IT10635** - Unrecognized Race Model
  * **IT10636** - Unrecognized Race Model
  * **IT10637** - Unrecognized Race Model
  * **IT10638** - Unrecognized Race Model
  * **IT10639** - Unrecognized Race Model
  * **IT10640** - Unrecognized Race Model
  * **IT10641** - Unrecognized Race Model
  * **IT10642** - Unrecognized Race Model
  * **IT10643** - Unrecognized Race Model
  * **IT10644** - Unrecognized Race Model
  * **IT10645** - Unrecognized Race Model
  * **IT10646** - Unrecognized Race Model
  * **IT10647** - Unrecognized Race Model
  * **IT10648** - Unrecognized Race Model
  * **IT10649** - Unrecognized Race Model
  * **IT10650** - Unrecognized Race Model
  * **IT10651** - Unrecognized Race Model
  * **IT10652** - Unrecognized Race Model
  * **IT10653** - Unrecognized Race Model
  * **IT10654** - Unrecognized Race Model
  * **IT10655** - Unrecognized Race Model
  * **IT10656** - Unrecognized Race Model
  * **IT10657** - Unrecognized Race Model
  * **IT10658** - Unrecognized Race Model
  * **IT10659** - Unrecognized Race Model
  * **IT10660** - Unrecognized Race Model
  * **IT10661** - Unrecognized Race Model
  * **IT10662** - Unrecognized Race Model
  * **IT10663** - Unrecognized Race Model
  * **IT10664** - Unrecognized Race Model
  * **IT10665** - Unrecognized Race Model
  * **IT10666** - Unrecognized Race Model
  * **IT10667** - Unrecognized Race Model
  * **IT10668** - Unrecognized Race Model
  * **IT10669** - Unrecognized Race Model
  * **IT10670** - Unrecognized Race Model
  * **IT10671** - Unrecognized Race Model
* **Source:** gequip4.s3d (Loaded via GlobalLoad.txt)
  * **IT10015** - Unrecognized Race Model
  * **IT10026** - Unrecognized Race Model
  * **IT10027** - Unrecognized Race Model
  * **IT10028** - Unrecognized Race Model
  * **IT10029** - Unrecognized Race Model
  * **IT10203** - Unrecognized Race Model
  * **IT11013** - Unrecognized Race Model
  * **IT11017** - Unrecognized Race Model
  * **IT11018** - Unrecognized Race Model
  * **IT11019** - Unrecognized Race Model
  * **IT11020** - Unrecognized Race Model
  * **IT11502** - Unrecognized Race Model
  * **IT67367** - Unrecognized Race Model
* **Source:** gequip3.s3d (Loaded via GlobalLoad.txt)
  * **IT10000** - Unrecognized Race Model
  * **IT10001** - Unrecognized Race Model
  * **IT10002** - Unrecognized Race Model
  * **IT10003** - Unrecognized Race Model
  * **IT10004** - Unrecognized Race Model
  * **IT10005** - Unrecognized Race Model
  * **IT10006** - Unrecognized Race Model
  * **IT10007** - Unrecognized Race Model
  * **IT10008** - Unrecognized Race Model
  * **IT10009** - Unrecognized Race Model
  * **IT10010** - Unrecognized Race Model
  * **IT10011** - Unrecognized Race Model
  * **IT10012** - Unrecognized Race Model
  * **IT10013** - Unrecognized Race Model
  * **IT10014** - Unrecognized Race Model
  * **IT10016** - Unrecognized Race Model
  * **IT10017** - Unrecognized Race Model
  * **IT10018** - Unrecognized Race Model
  * **IT10019** - Unrecognized Race Model
  * **IT10020** - Unrecognized Race Model
  * **IT10021** - Unrecognized Race Model
  * **IT10022** - Unrecognized Race Model
  * **IT10023** - Unrecognized Race Model
  * **IT10024** - Unrecognized Race Model
  * **IT10025** - Unrecognized Race Model
  * **IT10100** - Unrecognized Race Model
  * **IT10101** - Unrecognized Race Model
  * **IT10103** - Unrecognized Race Model
  * **IT10104** - Unrecognized Race Model
  * **IT10105** - Unrecognized Race Model
  * **IT10106** - Unrecognized Race Model
  * **IT10107** - Unrecognized Race Model
  * **IT10108** - Unrecognized Race Model
  * **IT10109** - Unrecognized Race Model
  * **IT10200** - Unrecognized Race Model
  * **IT10201** - Unrecognized Race Model
  * **IT10202** - Unrecognized Race Model
  * **IT10300** - Unrecognized Race Model
  * **IT10301** - Unrecognized Race Model
  * **IT11001** - Unrecognized Race Model
  * **IT11002** - Unrecognized Race Model
  * **IT11003** - Unrecognized Race Model
  * **IT11017** - Unrecognized Race Model
  * **IT11500** - Unrecognized Race Model
  * **IT11501** - Unrecognized Race Model
  * **MINIPOM200** - Mini POM (Race 252, Gender 2)
* **Source:** loyequip.s3d (Loaded via GlobalLoad.txt)
  * **IT10672** - Unrecognized Race Model
  * **IT10673** - Unrecognized Race Model
  * **IT10674** - Unrecognized Race Model
  * **IT10675** - Unrecognized Race Model
  * **IT10676** - Unrecognized Race Model
  * **IT10677** - Unrecognized Race Model
  * **IT10678** - Unrecognized Race Model
  * **IT10679** - Unrecognized Race Model
  * **IT10680** - Unrecognized Race Model
  * **IT10681** - Unrecognized Race Model
  * **IT10682** - Unrecognized Race Model
  * **IT10683** - Unrecognized Race Model
  * **IT10685** - Unrecognized Race Model
  * **IT10686** - Unrecognized Race Model
  * **IT10687** - Unrecognized Race Model
  * **IT10688** - Unrecognized Race Model
  * **IT10689** - Unrecognized Race Model
  * **IT10690** - Unrecognized Race Model
  * **IT10691** - Unrecognized Race Model
  * **IT10692** - Unrecognized Race Model
  * **IT10693** - Unrecognized Race Model
  * **IT10694** - Unrecognized Race Model
  * **IT10695** - Unrecognized Race Model
  * **IT10696** - Unrecognized Race Model
  * **IT10697** - Unrecognized Race Model
* **Source:** ldonequip.s3d (Loaded via GlobalLoad.txt)
  * **IT10700** - Unrecognized Race Model
  * **IT10701** - Unrecognized Race Model
  * **IT10702** - Unrecognized Race Model
  * **IT10703** - Unrecognized Race Model
  * **IT10704** - Unrecognized Race Model
  * **IT10705** - Unrecognized Race Model
  * **IT10706** - Unrecognized Race Model
  * **IT10707** - Unrecognized Race Model
  * **IT10708** - Unrecognized Race Model
  * **IT10709** - Unrecognized Race Model
  * **IT10710** - Unrecognized Race Model
  * **IT10711** - Unrecognized Race Model
  * **IT10712** - Unrecognized Race Model
  * **IT10713** - Unrecognized Race Model
  * **IT10714** - Unrecognized Race Model
* **Source:** gatesequip.s3d (Loaded via GlobalLoad.txt)
  * **IT10715** - Unrecognized Race Model
  * **IT10716** - Unrecognized Race Model
  * **IT10717** - Unrecognized Race Model
  * **IT10718** - Unrecognized Race Model
  * **IT10719** - Unrecognized Race Model
  * **IT10720** - Unrecognized Race Model
  * **IT10722** - Unrecognized Race Model
  * **IT10723** - Unrecognized Race Model
  * **IT10724** - Unrecognized Race Model
  * **IT10725** - Unrecognized Race Model
  * **IT10726** - Unrecognized Race Model
  * **IT10727** - Unrecognized Race Model
  * **IT10728** - Unrecognized Race Model
  * **IT10729** - Unrecognized Race Model
  * **IT10730** - Unrecognized Race Model
  * **IT10731** - Unrecognized Race Model
  * **IT10732** - Unrecognized Race Model
* **Source:** rap_chr.s3d (Loaded via GlobalLoad.txt)
  * **RAP** - Raptor (Race 163, Gender 2) - Textures: **00 01**
* **Source:** dpm_chr.s3d (Loaded via GlobalLoad.txt)
  * **DPM** - Pirate Male (Race 339, Gender 0)
* **Source:** dpf_chr.s3d (Loaded via GlobalLoad.txt)
  * **DPF** - Pirate Female (Race 339, Gender 1)
* **Source:** epm_chr.s3d (Loaded via GlobalLoad.txt)
  * **EPM** - Pirate Male (Race 342, Gender 0)
* **Source:** epf_chr.s3d (Loaded via GlobalLoad.txt)
  * **EPF** - Pirate Female (Race 342, Gender 1)
* **Source:** gpm_chr.s3d (Loaded via GlobalLoad.txt)
  * **GPM** - Pirate Male (Race 338, Gender 0)
* **Source:** gpf_chr.s3d (Loaded via GlobalLoad.txt)
  * **GPF** - Pirate Female (Race 338, Gender 1)
* **Source:** hpf_chr.s3d (Loaded via GlobalLoad.txt)
  * **HPF** - Pirate Female (Race 341, Gender 1)
* **Source:** hpm_chr.s3d (Loaded via GlobalLoad.txt)
  * **HPM** - Pirate Male (Race 341, Gender 0)
* **Source:** opf_chr.s3d (Loaded via GlobalLoad.txt)
  * **OPF** - Pirate Female (Race 340, Gender 1)
* **Source:** opm_chr.s3d (Loaded via GlobalLoad.txt)
  * **OPM** - Pirate Male (Race 340, Gender 0)
* **Source:** tbm_chr.s3d (Loaded via GlobalLoad.txt)
  * **TBM** - Troll Male (Race 331, Gender 0)
* **Source:** tbf_chr.s3d (Loaded via GlobalLoad.txt)
  * **TBF** - Troll Female (Race 331, Gender 1)
* **Source:** mmm_chr.s3d (Loaded via GlobalLoad.txt)
  * **MMM** - Vampire Male (Race 360, Gender 0) - Textures: **00 01**, Heads: **00 01**
* **Source:** mmf_chr.s3d (Loaded via GlobalLoad.txt)
  * **MMF** - Vampire Female (Race 360, Gender 1) - Textures: **00 01**, Heads: **00 01**
* **Source:** tnf_chr.s3d (Loaded via GlobalLoad.txt)
  * **TNF** - Nihil Female (Race 385, Gender 1) - Textures: **00-04**
* **Source:** tnm_chr.s3d (Loaded via GlobalLoad.txt)
  * **TNM** - Nihil Male (Race 385, Gender 0) - Textures: **00-04**
* **Source:** skt_chr.s3d (Loaded via GlobalLoad.txt)
  * **SKT** - Skeleton (Race 367, Gender 2) - Textures: **00-04**
* **Source:** global6\_chr.s3d (Loaded via GlobalLoad.txt)
  * **ALL** - Alligator (Race 91, Gender 2) - Textures: **00 01**, Heads: **00 01**
  * **DRK** - Drake (Race 89, Gender 2) - Textures: **00-03**
  * **WOF** - Chokidai (Race 356, Gender 2) - Textures: **00 01**
  * **SKE** - Skeleton (Race 60, Gender 2)
  * **TPN** - Teleport Man (Race 240, Gender 2)
  * **TIG** - Tiger (Race 63, Gender 2)
  * **WOL** - Wolf (Race 42, Gender 2) - Textures: **00-03**
  * **WOE** - Wolf (Race 120, Gender 2) - Textures: **00-03**
* **Source:** global4\_chr.s3d (Loaded via GlobalLoad.txt)
  * **IKS** - Undead Iksar (Race 161, Gender 2) - Heads: **00 01**
  * **SPE** - Spectre (Race 85, Gender 2)
  * **BAF** - Barbarian Female (Race 2, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **BAM** - Barbarian Male (Race 2, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **IKF** - Iksar Female (Race 128, Gender 1) - Textures: **00-04, 10**, Heads: **00-03**, Faces: **00-07**
  * **IKM** - Iksar Male (Race 128, Gender 0) - Textures: **00-04, 10**, Heads: **00-03**, Faces: **00-07**
* **Source:** global_chr.s3d (Loaded via GlobalLoad.txt)
  * **ELE** - Elemental (Race 75, Gender 2) - Textures: **00-03**, Heads: **00 01**
  * **EYE** - Eye (Race 108, Gender 2)
  * **IVM** - Invisible Man of Zomm Male (Race 600, Gender 0)
  * **SKE** - Skeleton (Race 60, Gender 2) - Textures: **00 01**
  * **WER** - Werewolf (Race 14, Gender 2)
  * **WOE** - Wolf (Race 120, Gender 2) - Textures: **00 01**
  * **BOAT** - Boat (Race 141, Gender 2)
  * **BAF** - Barbarian Female (Race 2, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **BAM** - Barbarian Male (Race 2, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **DAF** - Dark Elf Female (Race 6, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **DAM** - Dark Elf Male (Race 6, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **DWF** - Dwarf Female (Race 8, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **DWM** - Dwarf Male (Race 8, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **ERF** - Erudite Female (Race 3, Gender 1) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **ERM** - Erudite Male (Race 3, Gender 0) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **GNF** - Gnome Female (Race 12, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-04**, Faces: **00-07**
  * **GNM** - Gnome Male (Race 12, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-04**, Faces: **00-07**
  * **HAF** - Half Elf Female (Race 7, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **HAM** - Half Elf Male (Race 7, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **HOF** - Halfling Female (Race 11, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **HOM** - Halfling Male (Race 11, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **HIF** - High Elf Female (Race 5, Gender 1) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **HIM** - High Elf Male (Race 5, Gender 0) - Textures: **00-03, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **HUF** - Human Female (Race 1, Gender 1) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **HUM** - Human Male (Race 1, Gender 0) - Textures: **00-04, 10-16**, Heads: **00-03**, Faces: **00-07**
  * **OGF** - Ogre Female (Race 10, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **OGM** - Ogre Male (Race 10, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **TRF** - Troll Female (Race 9, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **TRM** - Troll Male (Race 9, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **ELF** - Wood Elf Female (Race 4, Gender 1) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
  * **ELM** - Wood Elf Male (Race 4, Gender 0) - Textures: **00-03**, Heads: **00-03**, Faces: **00-07**
* **Source:** fgh_chr.s3d (Loaded via GlobalLoad.txt)
  * **FGH** - Froglok Ghost (Race 371, Gender 2)
* **Source:** ila_chr.s3d (Loaded via GlobalLoad.txt)
  * **ILA** - Ikaav (Race 394, Gender 2)
* **Source:** rpt.eqg (Loaded via GlobalLoad.txt)
  * **RPT** - Raptor (Race 609, Gender 2) - Textures: **00-02**
* **Source:** jkr_chr.s3d (Loaded via GlobalLoad.txt)
  * **JKR** - Jokester (Race 384, Gender 2) - Textures: **00-02**
* **Source:** mmy_chr.s3d (Loaded via GlobalLoad.txt)
  * **MMY** - Mummy (Race 368, Gender 2)
* **Source:** global2\_chr.s3d (Loaded via GlobalLoad.txt)
  * **BEA** - Bear (Race 43, Gender 2) - Textures: **00-02**, Heads: **00-02**
  * **BRI** - Bristlebane (Race 153, Gender 2)
  * **CAZ** - Cazic Thule (Race 95, Gender 2) - Heads: **00 01**
  * **ERO** - Erollisi (Race 150, Gender 2)
  * **IMP** - Imp (Race 46, Gender 2)
  * **INN** - Innoruuk (Race 123, Gender 2)
  * **RAL** - Rallos Zek (Race 66, Gender 2) - Textures: **00 01**
  * **SCA** - Scarecrow (Race 82, Gender 2)
  * **SOL** - Solusek Ro (Race 58, Gender 2)
  * **TRI** - Tribunal (Race 151, Gender 2)
  * **TUN** - Tunare (Race 62, Gender 2)
{% endtab %}
{% endtabs %}
