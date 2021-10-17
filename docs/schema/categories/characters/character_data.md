# character\_data

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Character Identifier |
| account\_id | int | [Account Identifier](../../../schema/categories/account/account.md) |
| name | varchar | Name |
| last\_name | varchar | Last Name |
| title | varchar | Title |
| suffix | varchar | Suffix |
| zone\_id | int | [Zone Identifier](../../../../categories/zones/zone-list) |
| zone\_instance | int | Zone Instance Identifier |
| y | float | Y Coordinate |
| x | float | X Coordinate |
| z | float | Z Coordinate |
| heading | float | Heading Coordinate |
| gender | tinyint | [Gender](../../../../categories/npc/genders) |
| race | smallint | [Race](../../../../categories/npc/race-list) |
| class | tinyint | [Class](../../../../categories/player/class-list) |
| level | int | Level |
| deity | int | [Deity](../../../../categories/player/deity-list) |
| birthday | int | UNIX Timestamp of Birthday |
| last\_login | int | UNIX Timestamp of Last Login |
| time\_played | int | Time Played |
| level2 | tinyint | Level 2 |
| anon | tinyint | Anon: 0 = False, 1 = Anonymous, 2 = Roleplaying |
| gm | tinyint | GM: 0 = False, 1 = True |
| face | int | Face |
| hair\_color | tinyint | Hair Color |
| hair\_style | tinyint | Hair Style |
| beard | tinyint | Beard |
| beard\_color | tinyint | Beard Color |
| eye\_color\_1 | tinyint | Eye Color 1 |
| eye\_color\_2 | tinyint | Eye Color 2 |
| drakkin\_heritage | int | Drakkin Heritage |
| drakkin\_tattoo | int | Drakkin Tattoo |
| drakkin\_details | int | Drakkin Details |
| ability\_time\_seconds | tinyint | Ability Timer in Seconds |
| ability\_number | tinyint | [Ability Number](../../../schema/categories/characters/aa_ability.md) |
| ability\_time\_minutes | tinyint | Ability Timer in Minutes |
| ability\_time\_hours | tinyint | Ability Timer in Hours |
| exp | int | Experience |
| aa\_points\_spent | int | AA Points Spent |
| aa\_exp | int | AA Experience |
| aa\_points | int | AA Points |
| group\_leadership\_exp | int | Group Leadership Experience |
| raid\_leadership\_exp | int | Raid Leadership Experience |
| group\_leadership\_points | int | Group Leadership Points |
| raid\_leadership\_points | int | Raid Leadership Points |
| points | int | Points |
| cur\_hp | int | Health |
| mana | int | Mana |
| endurance | int | Endurance |
| intoxication | int | Intoxication |
| str | int | Strength |
| sta | int | Stamina |
| cha | int | Charisma |
| dex | int | Dexterity |
| int | int | Intelligence |
| agi | int | Agility |
| wis | int | Wisdom |
| zone\_change\_count | int | Zone Change Count |
| toxicity | int | Toxicity |
| hunger\_level | int | Hunger Level |
| thirst\_level | int | Thirst Level |
| ability\_up | int | Ability Up |
| ldon\_points\_guk | int | LDoN Points - [Deepest Guk](../../../../categories/zones/ldon-themes) |
| ldon\_points\_mir | int | LDoN Points - [Miragul's Menagerie](../../../../categories/zones/ldon-themes) |
| ldon\_points\_mmc | int | LDoN Points - [Mistmoore Catacombs](../../../../categories/zones/ldon-themes) |
| ldon\_points\_ruj | int | LDoN Points - [Rujarkian Hills](../../../../categories/zones/ldon-themes) |
| ldon\_points\_tak | int | LDoN Points - [Takish-Hiz](../../../../categories/zones/ldon-themes) |
| ldon\_points\_available | int | LDoN Points Available |
| tribute\_time\_remaining | int | Tribute Time Remaining |
| career\_tribute\_points | int | Career Tribute Points |
| tribute\_points | int | Tribute Points |
| tribute\_active | int | Tribute Active: 0 = False, 1 = True |
| pvp\_status | tinyint | PVP Status: 0 = False, 1 = True |
| pvp\_kills | int | PVP Kills |
| pvp\_deaths | int | PVP Deaths |
| pvp\_current\_points | int | PVP Points |
| pvp\_career\_points | int | Career PVP Points |
| pvp\_best\_kill\_streak | int | Best Kill Streak |
| pvp\_worst\_death\_streak | int | Worse Death Streak |
| pvp\_current\_kill\_streak | int | Current Kill Streak |
| pvp2 | int | PVP Status: 0 = False, 1 = True |
| pvp\_type | int | PVP Type |
| show\_helm | int | Show Helm: 0 = False, 1 = True |
| group\_auto\_consent | tinyint | Group Auto Consent: 0 = False, 1 = True |
| raid\_auto\_consent | tinyint | Raid Auto Consent: 0 = False, 1 = True |
| guild\_auto\_consent | tinyint | Guild Auto Consent: 0 = False, 1 = True |
| leadership\_exp\_on | tinyint | Leadership Experience On : 0 = False, 1 = True |
| RestTimer | int | Rest Timer |
| air\_remaining | int | Air Remaining |
| autosplit\_enabled | int | Autosplit Enabled: 0 = False, 1 = True |
| lfp | tinyint | Looking For Party: 0 = False, 1 = True |
| lfg | tinyint | Looking For Guild: 0 = False, 1 = True |
| mailkey | char | Mail Key |
| xtargets | tinyint | XTargets |
| firstlogon | tinyint | First Logon |
| e\_aa\_effects | int | Experience AA Effects |
| e\_percent\_to\_aa | int | Experience Percentage to AA Points |
| e\_expended\_aa\_spent | int | Expended AA Points Spent |
| aa\_points\_spent\_old | int | AA Points Spent Old |
| aa\_points\_old | int | AA Points Old |
| e\_last\_invsnapshot | int | Last Inventory Snapshot |
| deleted\_at | datetime |  |

