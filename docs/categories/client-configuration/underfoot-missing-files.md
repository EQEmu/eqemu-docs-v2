# Underfoot Missing Files

## Description

**Getting the Underfoot (UF) client is simple, since it can be purchased and downloaded from Steam. The only downside to the Underfoot client is that it is missing certain files that you will need to replace to make the client as fully functional as the SoD client.**

> Once you purchase Underfoot from Steam and have it downloaded, you should make a backup copy of the entire folder immediately and store it somewhere else on your system. Next you should make a copy of the folder where ever you want it to be to played from, C:\EQ\Underfoot is a good location. Playing from the Steam folder can ruin your install if Steam updates the download, so you will not want to do that.

* You will need at least a Titanium or older install for the zones that have been changed or no longer exist on Live. Some examples of these zones are East and West Commonlands, Nektulos, North and South Ro, The Oasis, and East and West Freeport. The only place to find these files is from a Titanium or older install. Seeds of Destruction (SoD) and newer clients do not have these files at all. 
* The other files you will need are files that should be in a normal EverQuest install but are missing from the Steam download for some reason. Most, but all of these files can also be taken from a Titanium install, but as more things are discovered it may be required to get them from a newer source. Once Underfoot has all of its missing files, it should run as well as the SoD client.
* There is no way to help you get the older missing files except from the original discs such as Titanium or another box set. Titanium is the best source, but anything from the Gold box or newer should have most of them. If you are really desperate, you could likely salvage most of the files from the Trilogy discs and then the rest from the PoP expansion discs.
* For the missing/corrupt files that can still be gotten from a Live install, you can now get a fully patched to live client by signing up for the 14 day free trial. 

#### Update August 18, 2011 <a href="ac_update-august-18-2011" id="ac_update-august-18-2011"></a>

> The old Escape To Norrath trial was never useful because it was only a partial client. But Sony has now changed the trial to be a trial on the actual live servers. Go make a new SOE account and hit the trial button. It will download the full client (6.5GB). To speed up the download, you can take your existing UF install and copy all of those files to the folder where you told the live install to go. That will extremely shorten your download time. Note that the steps below describe how to copy your UF install folder and create a new Live folder that will be used to attain the files needed for UF. NEVER patch your UF install folder or it will fail to work with EQEmu!

**These steps were used on August 18, 2011: **

1. Locate the Underfoot install. Mine resides in C:\EQ\Everquest_UF\\.
2. Make a new folder for the Live install, C:\EQ\Everquest_Live\ is a good location.
3. Copy everything from the Underfoot folder to the Live folder.
4. While that is copying, open your browser and go to http://www.everquest.com∞ and click on Everquest and then look for the button for the free trial.
5. Create a new account. An existing account will not work. When it tells you to download the installer, ignore it.
6. Go to the Live folder you made and double click on the launchpad.exe program.
7. Login with the new account you just created.
8. Wait for it to finish patching.
9. You are done. I recommend noting the expiration date of your trial and patching again one last time before it expires to make sure you have everything as of that date.

## **Missing Files**

* Here is a list compiled from different posts of files missing in the UF client that can be copied over from Live or another client like SoF or even some from Titanium:
  * **dbstr_us.txt** - This one is the most important, otherwise many AAs and spell descriptions will not show.
  * **dkm_anims.eqg** - This one is important for playing Male Drakkin, otherwise they don't animate.

These files are required for revamped or removed zones to show everything correctly, and can all be found in a Titanium installation. These are all original zones that were changed and it is likely that even an original EverQuest disc would work, but no one has reported testing it out.

```
 File Date     File Size  File Name
2005-10-17 11:37        178030  blackburrow_2_obj.s3d
2005-10-17 11:30          2520  freportw_sounds.eff
2005-10-17 11:28        163048  freportn.xmi
2005-10-17 11:29       1809356  freportn_chr.s3d
2005-10-17 11:29           102  freportn_sndbnk.eff
2005-10-17 11:29          1344  freportn_sounds.eff
2005-10-17 11:29         68840  freportw.xmi
2005-10-17 11:30       2628677  freportw_chr.s3d
2005-10-17 11:29           184  freportw_sndbnk.eff
2005-10-17 11:35            24  freportw_assets.txt
2005-10-17 11:28         26548  nro.xmi
2005-10-17 11:28       2569236  nro_chr.s3d
2005-10-17 11:34            84  nro_chr.txt
2005-10-17 11:33        301167  nro_obj2.s3d
2005-10-17 11:28           137  nro_sndbnk.eff
2005-10-17 11:28          2940  nro_sounds.eff
2005-10-17 11:30         73800  nro2_chr.s3d
2005-10-17 11:30          2856  oasis_sounds.eff
2005-10-17 11:29         73802  oasis2_chr.s3d
2005-10-17 11:33        373067  oasis_2_obj.s3d
2005-10-17 11:30       2544081  oasis_chr.s3d
2005-10-17 11:34            52  oasis_chr.txt
2005-10-17 11:29           153  oasis_sndbnk.eff
2005-10-17 11:39         19505  oot_chr2.s3d
2005-10-17 11:39        512580  rathemtn_chr2.s3d
2005-10-17 11:37          2856  soltemple_sounds.eff
2005-10-17 11:36        244516  soltemple_chr.s3d
2005-10-17 11:37            85  soltemple_sndbnk.eff
2005-10-17 11:30          2772  sro_sounds.eff
2005-10-17 11:33        297528  sro_2_obj.s3d
2005-10-17 11:29       2484461  sro_chr.s3d
2005-10-17 11:35            21  sro_chr.txt
2005-10-17 11:29           178  sro_sndbnk.eff
```

* **These files are required because they are simply missing or incomplete from the Underfoot download. Some can be obtained from Titanium and some from SoD. The Live client is also a valid source for all of the below files.**

```
  File Date     File Size  File Name
--------->From Titanium, SoD or Live<-----------
2005-10-17 11:34       1944364  bas.eqg
2005-10-17 11:28       2415064  blackburrow_chr.s3d
2005-10-17 11:39       2743860  bothunder.mp3
2005-10-17 11:39       7127397  bothunder.s3d
2005-10-17 11:39       3090804  bothunder_chr.s3d
2005-10-17 11:41            34  bothunder_chr.txt
2005-10-17 11:39       3988692  bothunder_obj.s3d
2005-10-17 11:38            77  dreadlands_chr.txt
2005-10-17 11:38       3144939  frozenshadow_chr.s3d
2005-10-17 11:30        873955  gequip3.s3d
2005-10-17 11:44           262  gukg_chr.txt
2005-10-17 11:42       2239485  gukh.s3d
2005-10-17 12:42           335  gunthak_chr.txt
2005-10-17 11:40       7481382  pofire_chr.s3d
2005-10-17 11:40       8906873  postorms.s3d
2005-10-17 11:41       2431425  powater_chr.s3d
2005-10-17 11:45         17836  qvic.emt
2005-10-17 11:44       3446081  qvic.s3d
2005-10-17 11:45           212  qvic_chr.txt
2005-10-17 11:44       3088488  qvic_obj.s3d
----->From SoD or Live (these are from Live)<------
2010-03-26 21:07       2269493  dest_sphere_shield.eqg
2010-03-26 16:16       1702782  discordtower.zon
2010-03-26 16:14       8825512  dkm_anims.eqg
2010-03-26 16:34       6904245  freeportacademy.eqg
2010-03-26 20:16            55  poknowledge_assets.txt
2011-02-01 01:15        176612  spellsnew.eff
2011-02-01 09:50        158261  RaceData.txt
2011-02-02 17:37       4552251  dbstr_us.txt
2011-02-03 22:27      18971769  spells_us.txt
2011-02-04 20:32       1402768  mpu.eqg
```

* Alternatively to copying these files over 1 by 1 from another client, you can simply select all files in another client such as a EQLive patched client install folder and paste them into the Underfoot folder. Then, when prompted if you want to replace existing files, just set it to say no to all, and only missing files will be copied over. Then, the only thing that has to be manually copied over is the dbstr_us.txt file.
* If there are more missing files that are not listed here, feel free to update this Wiki page to add the additional file names to the list, or any extra helpful details.
