# Inventory Slots



### Server-recognized slots

#### Equipment slots \(worn\)

* Charm = 0
* Ear1 = 1
* Head = 2
* Face = 3
* Ear2 = 4
* Neck = 5
* Shoulders = 6
* Arms = 7
* Back = 8
* Wrist1 = 9
* Wrist2 = 10
* Range = 11
* Hands = 12
* Primary = 13
* Secondary = 14
* Finger1 = 15
* Finger2 = 16
* Chest = 17
* Legs = 18
* Feet = 19
* Waist = 20
* PowerSource = 21
* Ammo = 22

![Equipment Slots](../../.gitbook/assets/image.png)

#### General slots \(personal\)

* General1 = 23
* General2 = 24
* General3 = 25
* General4 = 26
* General5 = 27
* General6 = 28
* General7 = 29
* General8 = 30
* General9 = 31
* General10 = 32

![General Slots](../../.gitbook/assets/image%20%288%29.png)

Bag slots in general inventory are:

* General1 : 251 -&gt; 260
* General2 : 261 -&gt; 270
* General3 : 271 -&gt; 280
* General4 : 281 -&gt; 290
* General5 : 291 -&gt; 300
* General6 : 301 -&gt; 310
* General7 : 311 -&gt; 320
* General8 : 321 -&gt; 330
* General9 : 331 -&gt; 340
* General10 : 341 -&gt; 350

#### Cursor slot

* Cursor = 33

Bag slots in cursor inventory are:

* Cursor : 351 -&gt; 360

#### Tribute

Tribute items are slots 400 -&gt; 404, these items are not visible, but are counted for stats/effects.

#### Bank

Bank slots are 2000 -&gt; 2023

Bags in the bank are:

* 2000 : 2031 -&gt; 2040
* 2001 : 2041 -&gt; 2050
* 2002 : 2051 -&gt; 2060
* 2003 : 2061 -&gt; 2070
* 2004 : 2071 -&gt; 2080
* 2005 : 2081 -&gt; 2090
* 2006 : 2091 -&gt; 2100
* 2007 : 2101 -&gt; 2110
* 2008 : 2111 -&gt; 2120
* 2009 : 2121 -&gt; 2130
* 2010 : 2131 -&gt; 2140
* 2011 : 2141 -&gt; 2150
* 2012 : 2151 -&gt; 2160
* 2013 : 2161 -&gt; 2170
* 2014 : 2171 -&gt; 2180
* 2015 : 2181 -&gt; 2190
* 2016 : 2191 -&gt; 2200
* 2017 : 2201 -&gt; 2210
* 2018 : 2211 -&gt; 2220
* 2019 : 2221 -&gt; 2230
* 2020 : 2231 -&gt; 2240
* 2021 : 2241 -&gt; 2250
* 2022 : 2251 -&gt; 2260
* 2023 : 2261 -&gt; 2270

#### Shared Bank

Shared bank slots are 2500 and 2501

Note: These are stored in the sharedbank table, not the inventory table.

Bags in the shared bank are:

* 2500 : 2531 -&gt; 2540
* 2501 : 2541 -&gt; 2550

![](../../.gitbook/assets/image%20%286%29.png)

**Note: Not all clients support all of the server-recognized slots. Care should be taken when attempting to hard-code slot values over the use of server-based free slot requests.**

#### Quest use references

* [Perl](https://github.com/EQEmu/Server/wiki/Perl-Inventory-Slot-Identifiers)
* [Lua](https://github.com/EQEmu/Server/wiki/Lua-Constants)

### Developer information on client-specific slot support

\(in-work\)

**SoF Slots**

#### Equipped slots

\(somebody should make this prettier some day\)

* SLOT\_CHARM= 0
* SLOT\_EAR01= 1
* SLOT\_HEAD= 2
* SLOT\_FACE= 3
* SLOT\_EAR02= 4
* SLOT\_NECK= 5
* SLOT\_SHOULDER= 6
* SLOT\_ARMS= 7
* SLOT\_BACK= 8
* SLOT\_BRACER01= 9
* SLOT\_BRACER02= 10
* SLOT\_RANGE= 11
* SLOT\_HANDS= 12
* SLOT\_PRIMARY= 13
* SLOT\_SECONDARY= 14
* SLOT\_RING01= 15
* SLOT\_RING02= 16
* SLOT\_CHEST= 17
* SLOT\_LEGS= 18
* SLOT\_FEET= 19
* SLOT\_WAIST= 20
* SLOT\_POWERSOURCE= 21
* SLOT\_AMMO= 22

#### Inventory Slots

NOTE: Numbering for personal inventory goes top to bottom, then left to right

It's the opposite for inside bags: left to right, then top to bottom

Example:

inventory:containers:

* 1 51 2
* 2 63 4
* 3 75 6
* 4 87 8
*   9 10

#### Personal Inventory

Personal inventory slots 23 through 30.

Bags in personal inventory are:

* 23: 262-&gt;271
* 24: 272-&gt;281
* 25: 282-&gt;291
* 26: 292-&gt;301
* 27: 302-&gt;311
* 28: 312-&gt;321
* 29: 322-&gt;331
* 30: 332-&gt;341

#### Cursor

Cursor is slot 31, and the bag slots for the cursor are 342-&gt;351 \(haven't confirmed slots inside bag on cursor for SoF yet\).

#### SoF Bank

Bank slots in SoF are 2000-&gt;2023

Bags in the bank are:

* 2000 2004 2008 2012 2016 2020
* 2001 2005 2009 2013 2017 2021
* 2002 2006 2010 2014 2018 2022
* 2003 2007 2011 2015 2019 2023

Slots inside those bags are:

Note: Other than the addition of the 8 new slots, the main bank slots are the same as Titanium, but slots inside bags are +1 to what Titanium uses.

* 2000 = 2032 - 2041
* 2001 = 2042 - 2051
* 2002 = 2052 - 2061
* 2003 = 2062 - 2071
* 2004 = 2072 - 2061
* 2005 = 2082 - 2091
* 2006 = 2092 - 2101
* 2007 = 2102 - 2111
* 2008 = 2112 - 2121
* 2009 = 2122 - 2131
* 2010 = 2132 - 2141
* 2011 = 2142 - 2151
* 2012 = 2152 - 2161
* 2013 = 2162 - 2171
* 2014 = 2172 - 2181
* 2015 = 2182 - 2191
* 8 New Bank Slots:
* 2016 = 2192 - 2201
* 2017 = 2202 - 2211
* 2018 = 2212 - 2221
* 2019 = 2222 - 2231
* 2020 = 2232 - 2241
* 2021 = 2242 - 2251
* 2022 = 2252 - 2261
* 2023 = 2262 - 2271

#### SoF Shared Bank

Shared bank slots in SoF are 2500 and 2501

Note: These are the same as in Titanium, but the slots for inside bags are +1 to the ones in Titanium

Bags in the shared bank are:

* 2500 = 2532 - 2541
* 2501 = 2542 - 2551

