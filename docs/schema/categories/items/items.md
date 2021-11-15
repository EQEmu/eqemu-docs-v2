# items

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique Item Identifier |
| minstatus | smallint | [Minimum Status](../../../../categories/player/status-levels) |
| Name | varchar | Name |
| aagi | int | Agility: -128 to 127 |
| ac | int | Armor Class |
| accuracy | int | Accuracy |
| acha | int | Charisma: -128 to 127 |
| adex | int | Dexterity: -128 to 127 |
| aint | int | Intelligence: -128 to 127 |
| artifactflag | tinyint | Artifact: 0 = False, 1 = True |
| asta | int | Stamina: -128 to 127 |
| astr | int | Strenght: -128 to 127 |
| attack | int | Attack |
| augrestrict | int | [Augment Restriction](../../../../categories/items/augment-restrictions) |
| augslot1type | tinyint | [Augment Slot 1 Type](../../../../categories/items/augment-types) |
| augslot1visible | tinyint | Augment Slot 1 Visible: 0 = False, 1 = True |
| augslot2type | tinyint | [Augment Slot 2 Type](../../../../categories/items/augment-types) |
| augslot2visible | tinyint | Augment Slot 2 Visible: 0 = False, 1 = True |
| augslot3type | tinyint | [Augment Slot 3 Type](../../../../categories/items/augment-types) |
| augslot3visible | tinyint | Augment Slot 3 Visible: 0 = False, 1 = True |
| augslot4type | tinyint | [Augment Slot 4 Type](../../../../categories/items/augment-types) |
| augslot4visible | tinyint | Augment Slot 4 Visible: 0 = False, 1 = True |
| augslot5type | tinyint | [Augment Slot 5 Type](../../../../categories/items/augment-types) |
| augslot5visible | tinyint | Augment Slot 5 Visible: 0 = False, 1 = True |
| augslot6type | tinyint | [Augment Slot 6 Type](../../../../categories/items/augment-types) |
| augslot6visible | tinyint | Augment Slot 6 Visible: 0 = False, 1 = True |
| augtype | int | [Augment Type](../../../../categories/items/augment-types) |
| avoidance | int | Avoidance |
| awis | int | Wisdom: -128 to 127 |
| bagsize | int | [Bag Size](../../../../categories/inventory/bag-sizes) |
| bagslots | int | Bag Slots: 1 = Minimum, 10 = Maximum |
| bagtype | int | [Bag Type](../../../../categories/inventory/bag-types) |
| bagwr | int | Bag Weight Reduction: 0 = 0%, 100 = 100% |
| banedmgamt | int | Bane Damage Amount |
| banedmgraceamt | int | Bane Damage Race Amount |
| banedmgbody | int | [Bane Damage Body Type](../../../../categories/npc/body-types) |
| banedmgrace | int | [Bane Damage Race](../../../../categories/npc/race-list) |
| bardtype | int | [Bard Type](../../../../categories/items/bard-types) |
| bardvalue | int | Bard Value |
| book | int | [Book](../../../schema/categories/books/books.md) |
| casttime | int | Cast Time in Seconds |
| casttime\_ | int | Cast Time in Seconds |
| charmfile | varchar | Charm File |
| charmfileid | varchar | Charm File Identifier |
| classes | int | [Classes](../../../../categories/player/class-list) |
| color | int | Color |
| combateffects | varchar | Combat Effects |
| extradmgskill | int | [Extra Damage Skill](../../../../categories/player/skills) |
| extradmgamt | int | Extra Damage Amount |
| price | int | Price in Copper |
| cr | int | Cold Resistance: -128 to 127 |
| damage | int | Damage |
| damageshield | int | Damage Shield |
| deity | int | [Deity](../../../../categories/player/deity-list) |
| delay | int | Delay |
| augdistiller | int | Augment Distiller Item Identifier |
| dotshielding | int | Damage Over Time Shielding |
| dr | int | Disease Resistance: -128 to 127 |
| clicktype | int | [Click Type](../../../../categories/items/click-types) |
| clicklevel2 | int | Click Level 2 |
| elemdmgtype | int | [Elemental Damage Type](../../../../categories/items/item-element-types) |
| elemdmgamt | int | Elemental Damage Amount |
| endur | int | Endurance |
| factionamt1 | int | Faction Amount 1 |
| factionamt2 | int | Faction Amount 2 |
| factionamt3 | int | Faction Amount 3 |
| factionamt4 | int | Faction Amount 4 |
| factionmod1 | int | Faction Modifier 1 |
| factionmod2 | int | Faction Modifier 2 |
| factionmod3 | int | Faction Modifier 3 |
| factionmod4 | int | Faction Modifier 4 |
| filename | varchar | File Name |
| focuseffect | int | [Focus Effect Identifier](../../../schema/categories/spells/spells_new.md) |
| fr | int | Fire Resistance: -128 to 127 |
| fvnodrop | int | Firiona Vie No Drop: 0 = False, 1 = True |
| haste | int | Haste: 0 = 0%, 255 = 255% |
| clicklevel | int | Click Level |
| hp | int | Health |
| regen | int | Health Regeneration |
| icon | int | Icon |
| idfile | varchar | Item Texture |
| itemclass | int | [Item Class](../../../../categories/items/item-class) |
| itemtype | int | [Item Type](../../../../categories/items/item-types) |
| ldonprice | int | LDoN Price |
| ldontheme | int | [LDoN Theme](../../../../categories/zones/ldon-themes) |
| ldonsold | int | LDoN Sold: 0 = False, 1 = True |
| light | int | Light |
| lore | varchar | Lore Description |
| loregroup | int | [Lore Group](../../../../categories/items/item-lore-groups) |
| magic | int | Magic: 0 = False, 1 = True |
| mana | int | Mana |
| manaregen | int | Mana Regeneration |
| enduranceregen | int | Endurance Regeneration |
| material | int | [Material](../../../../categories/npc/textures) |
| herosforgemodel | int | Hero's Forge Model |
| maxcharges | int | Maximum Charges |
| mr | int | Magic Resistance: -128 to 127 |
| nodrop | int | No Drop: 0 = True, 1 = False |
| norent | int | No Rent: 0 = True, 1 = False |
| pendingloreflag | tinyint | Pending Lore Flag: 0 = False, 1 = True |
| pr | int | Poison Resistance: -128 to 127 |
| procrate | int | Proc Rate: 0 = 100%, 50 = 150%, 100 = 200% |
| races | int | [Races](../../../../categories/npc/race-list) |
| range | int | Range: 0 to 255 |
| reclevel | int | Recommended Level |
| recskill | int | Recommended Skill Level |
| reqlevel | int | Required Level |
| sellrate | float | Sell Rate |
| shielding | int | Shielding: 5 = 5%, 20 = 20%, 50 = 50% |
| size | int | Size |
| skillmodtype | int | [Skill Modifier Type](../../../../categories/player/skills) |
| skillmodvalue | int | Skill Modifier Value |
| slots | int | [Slots](../../../../categories/inventory/item-slots) |
| clickeffect | int | [Click Effect Identifier](../../../schema/categories/spells/spells_new.md) |
| spellshield | int | Spell Shielding |
| strikethrough | int | Strikethrough |
| stunresist | int | Stun Resist |
| summonedflag | tinyint | Unknown |
| tradeskills | int | Tradeskill Item: 0 = False, 1= True |
| favor | int | Favor |
| weight | int | Weight: 10 = 1.0, 25 = 2.5, 100 = 10.0 |
| UNK012 | int | Unknown |
| UNK013 | int | Unknown |
| benefitflag | int | Unknown |
| UNK054 | int | Unknown |
| UNK059 | int | Unknown |
| booktype | int | [Book Language](../../../../categories/player/languages) |
| recastdelay | int | Recast Delay in Seconds |
| recasttype | int | Recast Type: -1 = None, &gt;0 = Recast Type used across all items |
| guildfavor | int | Guild Favor |
| UNK123 | int | Unknown |
| UNK124 | int | Unknown |
| attuneable | int | Attuneable: 0 = False, 1 = True |
| nopet | int | No Pet: 0 = False, 1 = True |
| updated | datetime | Updated Datetime |
| comment | varchar | Comment |
| UNK127 | int | Unknown |
| pointtype | int | Unknown |
| potionbelt | int | Potion Belt: 0 = False, 1 = True |
| potionbeltslots | int | Potion Belt Slots |
| stacksize | int | Stack Size |
| notransfer | int | No Transfer: 0 = False, 1 = True |
| stackable | int | Stackable: 0 = False, 1 = True |
| UNK134 | varchar | Unknown |
| UNK137 | int | Unknown |
| proceffect | int | [Proc Effect Identifier](../../../schema/categories/spells/spells_new.md) |
| proctype | int | Proc Type: 0 |
| proclevel2 | int | Proc Level 2 |
| proclevel | int | Proc Level |
| UNK142 | int | Unknown |
| worneffect | int | [Worn Effect Identifier](../../../schema/categories/spells/spells_new.md) |
| worntype | int | Worn Type: 2 = Worn |
| wornlevel2 | int | Worn Level 2 |
| wornlevel | int | Worn Level |
| UNK147 | int | Unknown |
| focustype | int | Focus Type: 6 = Focus |
| focuslevel2 | int | Focus Level 2 |
| focuslevel | int | Focus Level |
| UNK152 | int | Unknown |
| scrolleffect | int | [Scroll Effect Identifier](../../../schema/categories/spells/spells_new.md) |
| scrolltype | int | Scroll Type: 7 = Scroll |
| scrolllevel2 | int | Scroll Level 2 |
| scrolllevel | int | Scroll Level |
| UNK157 | int | Unknown |
| serialized | datetime | Serialized Datetime |
| verified | datetime | Verified Datetime |
| serialization | text | Serialization |
| source | varchar | Source |
| UNK033 | int | Unknown |
| lorefile | varchar | Lore File |
| UNK014 | int | Unknown |
| svcorruption | int | Corruption Resistance: -128 to 127 |
| skillmodmax | int | Skill Modifier Max |
| UNK060 | int | Unknown |
| augslot1unk2 | int | Unknown |
| augslot2unk2 | int | Unknown |
| augslot3unk2 | int | Unknown |
| augslot4unk2 | int | Unknown |
| augslot5unk2 | int | Unknown |
| augslot6unk2 | int | Unknown |
| UNK120 | int | Unknown |
| UNK121 | int | Unknown |
| questitemflag | int | Quest Item: 0 = False, 1 = True |
| UNK132 | text | Unknown |
| clickunk5 | int | Unknown |
| clickunk6 | varchar | Unknown |
| clickunk7 | int | Unknown |
| procunk1 | int | Unknown |
| procunk2 | int | Unknown |
| procunk3 | int | Unknown |
| procunk4 | int | Unknown |
| procunk6 | varchar | Unknown |
| procunk7 | int | Unknown |
| wornunk1 | int | Unknown |
| wornunk2 | int | Unknown |
| wornunk3 | int | Unknown |
| wornunk4 | int | Unknown |
| wornunk5 | int | Unknown |
| wornunk6 | varchar | Unknown |
| wornunk7 | int | Unknown |
| focusunk1 | int | Unknown |
| focusunk2 | int | Unknown |
| focusunk3 | int | Unknown |
| focusunk4 | int | Unknown |
| focusunk5 | int | Unknown |
| focusunk6 | varchar | Unknown |
| focusunk7 | int | Unknown |
| scrollunk1 | int | Unknown |
| scrollunk2 | int | Unknown |
| scrollunk3 | int | Unknown |
| scrollunk4 | int | Unknown |
| scrollunk5 | int | Unknown |
| scrollunk6 | varchar | Unknown |
| scrollunk7 | int | Unknown |
| UNK193 | int | Unknown |
| purity | int | Purity |
| evoitem | int | Evolving Item: 0 = False, 1 = True |
| evoid | int | [Evolving Identifier](items.md) |
| evolvinglevel | int | Evolving Level |
| evomax | int | Evolving Max |
| clickname | varchar | Click Name |
| procname | varchar | Proc Name |
| wornname | varchar | Worn Name |
| focusname | varchar | Focus Name |
| scrollname | varchar | Scroll Name |
| dsmitigation | smallint | Damage Shield Mitigation |
| heroic\_str | smallint | Heroic Strength |
| heroic\_int | smallint | Heroic Intelligence |
| heroic\_wis | smallint | Heroic Wisdom |
| heroic\_agi | smallint | Heroic Agility |
| heroic\_dex | smallint | Heroic Dexterity |
| heroic\_sta | smallint | Heroic Stamina |
| heroic\_cha | smallint | Heroic Charisma |
| heroic\_pr | smallint | Heroic Poison Resistance |
| heroic\_dr | smallint | Heroic Disease Resistance |
| heroic\_fr | smallint | Heroic Fire Resistance |
| heroic\_cr | smallint | Heroic Cold Resistance |
| heroic\_mr | smallint | Heroic Magic Resistance |
| heroic\_svcorrup | smallint | Heroic Corruption Resistance |
| healamt | smallint | Heal Amount: 0 to 32767 |
| spelldmg | smallint | Spell Damage: 0 to 32767 |
| clairvoyance | smallint | Clairvoyance |
| backstabdmg | smallint | Backstab Damage |
| created | varchar | Created |
| elitematerial | smallint | Elite Material |
| ldonsellbackrate | smallint | LDoN Sellback Rate |
| scriptfileid | smallint | Script File Name |
| expendablearrow | smallint | Expendable Arrow: 0 = False, 1 = True |
| powersourcecapacity | smallint | Powersource Capacity |
| bardeffect | smallint | [Bard Effect Identifier](../../../schema/categories/spells/spells_new.md) |
| bardeffecttype | smallint | [Bard Effect Type](../../../../categories/items/bard-types) |
| bardlevel2 | smallint | Bard Level 2 |
| bardlevel | smallint | Bard Level |
| bardunk1 | smallint | Unknown |
| bardunk2 | smallint | Unknown |
| bardunk3 | smallint | Unknown |
| bardunk4 | smallint | Unknown |
| bardunk5 | smallint | Unknown |
| bardname | varchar | Bard Name |
| bardunk7 | smallint | Unknown |
| UNK214 | smallint | Unknown |
| UNK219 | int | Unknown |
| UNK220 | int | Unknown |
| UNK221 | int | Unknown |
| heirloom | int | Heirloom: 0 = False, 1 = True |
| UNK223 | int | Unknown |
| UNK224 | int | Unknown |
| UNK225 | int | Unknown |
| UNK226 | int | Unknown |
| UNK227 | int | Unknown |
| UNK228 | int | Unknown |
| UNK229 | int | Unknown |
| UNK230 | int | Unknown |
| UNK231 | int | Unknown |
| UNK232 | int | Unknown |
| UNK233 | int | Unknown |
| UNK234 | int | Unknown |
| placeable | int | Placeable: 0 = False, 1 = True |
| UNK236 | int | Unknown |
| UNK237 | int | Unknown |
| UNK238 | int | Unknown |
| UNK239 | int | Unknown |
| UNK240 | int | Unknown |
| UNK241 | int | Unknown |
| epicitem | int | Epic Item: 0 = False, 1 = True |
| subtype | int |  |

