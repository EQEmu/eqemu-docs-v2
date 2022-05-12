---
description: A list of shaders for EQG zones
---

These shaders are only applicable to eqg files, s3d uses a different method



# Shader List

|Note|Zone|Material|Shader|Property|Type|Value|Descrption|Triangle|Bits|Location
|---|---|---|---|---|---|---|---|---|---|---
typical texture|dranikhollowsb|tunnel02|Opaque_MaxCB1.fx|e_TextureDiffuse0|2|sp_tunn05.dds|generic texture|1|17|-61.324 256.222 2.408
typical texture|dranikhollowsb|tunnel02|Opaque_MaxCB1.fx|e_TextureNormal0|2|sp_tunn05_n.dds|generic texture|1|17|-61.324 256.222 2.408
waterfall animated|broodlands|waterfalls|Opaque_MaxWaterFall.fx|e_TextureDiffuse0|2|wtr_waterfall_tile.dds|generic texture
waterfall animated|broodlands|waterfalls|Opaque_MaxWaterFall.fx|e_fSlide1X|0|-0.12|movement of waterfall
waterfall animated|broodlands|waterfalls|Opaque_MaxWaterFall.fx|e_fSlide1Y|0|-0.32|movement of waterfall
waterfall animated|broodlands|waterfalls|Opaque_MaxWaterFall.fx|e_fSlide2X|0|0|movement of waterfall
waterfall animated|broodlands|waterfalls|Opaque_MaxWaterFall.fx|e_fSlide2Y|0|-0.5|movement of waterfall
transparent texture|broodlands|volcanocard|Alpha_MaxCBSG1.fx|e_fShininess0|0|1|picture of volcano in bg
transparent texture|broodlands|volcanocard|Alpha_MaxCBSG1.fx|e_TextureDiffuse0|2|volcano_card_c.dds|
transparent texture|broodlands|volcanocard|Alpha_MaxCBSG1.fx|e_TextureNormal0|2|volcano_card_1g_n.dds|lava normal effect?
Chroma texture|broodlands|cnpyedge|Chroma_MaxC1.fx|e_TextureDiffuse0|2|swmp_canopy_trim.dds|"way to use opacity maps,
 white becomes alpha"|96342|17|1306.5 -1204.6 79.4
 Water Opacity|anguish|water|Opaque_MaxWater.fx|e_TextureDiffuse0|2|ra_watertest_c_01.dds
Water Opacity|anguish|water|Opaque_MaxWater.fx|e_TextureNormal0|2|water_n.dds
Water Opacity|anguish|water|Opaque_MaxWater.fx|e_TextureEnvironment0|2|water_e.dds
Water Opacity|anguish|water|Opaque_MaxWater.fx|e_fFresnelBias|0|0.17
Water Opacity|anguish|water|Opaque_MaxWater.fx|e_fFresnelPower|0|10
Water Opacity|anguish|water|Opaque_MaxWater.fx|e_fWaterColor1|3|255 0 0 21
Water Opacity|anguish|water|Opaque_MaxWater.fx|e_fWaterColor2|3|255 0 30 23
Water Opacity|anguish|water|Opaque_MaxWater.fx|e_fReflectionAmount|0|0.5
Water Opacity|anguish|water|Opaque_MaxWater.fx|e_fReflectionColor|3|255 255 255 255
Flowing Water|broodlands|watertable|Opaque_MaxWater.fx|e_TextureNormal0|2|rc_cavewater_n.dds
Flowing Water|broodlands|watertable|Opaque_MaxWater.fx|e_TextureEnvironment0|2|ra_watertest_e_01.dds
Flowing Water|broodlands|watertable|Opaque_MaxWater.fx|e_fFresnelBias|0|0.06
Flowing Water|broodlands|watertable|Opaque_MaxWater.fx|e_fFresnelPower|0|6.35
Flowing Water|broodlands|watertable|Opaque_MaxWater.fx|e_fWaterColor1|3|255 61 93 100
Flowing Water|broodlands|watertable|Opaque_MaxWater.fx|e_fWaterColor2|3|255 96 151 166
Flowing Water|broodlands|watertable|Opaque_MaxWater.fx|e_fReflectionAmount|0|0.01
Flowing Water|broodlands|watertable|Opaque_MaxWater.fx|e_fReflectionColor|3|255 255 255 255
Flowing Water|broodlands|watertable|Opaque_MaxWater.fx|e_fSlide1X|0|0.04
Flowing Water|broodlands|watertable|Opaque_MaxWater.fx|e_fSlide1Y|0|0.04
Flowing Water|broodlands|watertable|Opaque_MaxWater.fx|e_fSlide2X|0|0.03
Flowing Water|broodlands|watertable|Opaque_MaxWater.fx|e_fSlide2Y|0|0.03
Volcano Background|broodlands|volcanocard|Alpha_MaxCBSG1.fx|e_fShininess0|0|1||110911|Permeable,  23|-2096.9 -1309.6 28
Volcano Background|broodlands|volcanocard|Alpha_MaxCBSG1.fx|e_TextureDiffuse0|2|volcano_card_c.dds||||
Volcano Background|broodlands|volcanocard|Alpha_MaxCBSG1.fx|e_TextureNormal0|2|volcano_card_1g_n.dds||||
Flowing Lava|thenest|lavaflow|Opaque_MaxLava.fx|e_TextureDiffuse0|2|kl_lavaTop_c.dds|Flowing lava, in X direction only|308920|19|36.742 -2270.422 -477.900
Flowing Lava|thenest|lavaflow|Opaque_MaxLava.fx|e_TextureDiffuse1|2|kl_lavaBottom_c.dds||||
Flowing Lava|thenest|lavaflow|Opaque_MaxLava.fx|e_TextureNormal0|2|kl_lava_n.dds||||
Flowing Lava|thenest|lavaflow|Opaque_MaxLava.fx|e_fSlide1X|0|0.3||||
Flowing Lava|thenest|lavaflow|Opaque_MaxLava.fx|e_fSlide1Y|0|0||||
Flowing Lava|thenest|lavaflow|Opaque_MaxLava.fx|e_fSlide2X|0|0.2||||
Flowing Lava|thenest|lavaflow|Opaque_MaxLava.fx|e_fSlide2Y|0|0