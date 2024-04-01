---
description: This page describes how to load models into EQ
---

# Loading models into EQ

Loading NPC models in EQ comes in various fashions. This guide breaks down each way and the steps to do so.

## Use global load for NPC/Player models

This is probably the least recommended route. You may think to yourself, "I can just put all models in global load I add, and it'll always be available!" But, there's memory limits and zone load time problems with this mindset. Typically, it is recommended to ONLY put models that need to be visible in every zone in this file. This includes things like mounts, familiars, or illusions.

Entries looks like this:
```
frg,gukbottom_chr
dev,swampofnohope_chr
aro,aro
```

Note that the upper most number needs to align to number of entries in the file.

This is a comma separated list of model short name, and file this model exists.

[check out this eq model dump](https://gist.github.com/xackery/0efb7d9a588ba470c6c728b88c735c11). You can remove the suffix (.eqg/.s3d) and place it in here.

If you aren't sure the model's short name, [Shendare's Race Inventory](http://www.shendare.com/EQ/Emu/EQRI/RoF2_EQRaces.htm) may give you a hint of it's name.

## Use zone load for NPC/Player Models

Zone load is a more recommended way to add models. This method will cause EQ to only load the model when you enter said zone. The setup is similar to global.

You'll notice like abysmal.s3d, you have a file named abysmal_chr.txt:
```
11
RTN,RTN
TNM,TNM_CHR
TNF,TNF_CHR
BRL,BRL_CHR
BOX,BOX_CHR
CST,CST_CHR
SHP,SHP_CHR
tac,tac_chr
tmb,tmb_chr
ttb,ttb_chr
twf,twf_chr
```

Note that the 11 on top should equal the total number of entries in the file +1

.EQG follows the exact same pattern.

This pattern is identical to the globalload style, so you can use the [same eq model dump](https://gist.github.com/xackery/0efb7d9a588ba470c6c728b88c735c11) to get entry data.
