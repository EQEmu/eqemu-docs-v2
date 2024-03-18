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

# Shader References

- Name: AddAlpha_MaxC1.fx

  Example: anguish.eqg, clvgate.mod

- Name: AddAlpha_MaxCB1.fx

    Example: thundercrest.eqg, obj_roof_decor_a.mod

- Name: AddAlpha_MaxCBSG1.fx

    Example: crs.eqg, crs.mod

- Name: AddAlpha_MaxCG1.fx

    Example: tbsequip.eqg, it11107.mod

- Name: AddAlpha_MPLBasicA.fx

    Example: barren.eqg, obp_beam_yellow.mod

- Name: AddAlpha_MPLBasicAT.fx

    Example: pillarsalra.eqg, obj_nat_button_base.mod

- Name: AddAlpha_MPLBumpA.fx

    Example: wallet35.eqg, it12307.mod

- Name: AddAlphaC1Max.fx

    Example: fhalls.eqg, obp_geode01.mod

- Name: Alpha_MaxC1.fx

    Example: fie.eqg, fie.mod

- Name: Alpha_MaxCBS1.fx

    Example: aie.eqg, aie.mod

- Name: Alpha_MaxCBSG1.fx

    Example: delvea.eqg, obp_spderweb_small_d.mod

- Name: Alpha_MaxCBSGE1.fx

    Example: cry.eqg, cry.mds

- Name: Alpha_MaxCE1.fx

    Example: anguish.eqg, obj_cppods1.mod

- Name: Alpha_MPLBasicA.fx

    Example: alkabormare.eqg, obp_mistyfloraa_lod1.mod

- Name: Alpha_MPLBasicAT.fx

    Example: arelis.eqg, obj_weapon_rack_lod1.mod

- Name: Alpha_MPLBumpA.fx

    Example: bertoxtemple.eqg, obj_corral_topb.mod

- Name: Alpha_MPLBumpAT.fx

    Example: bertoxtemple.eqg, obp_plant_cluster.mod

- Name: AlphaSModelC1Max.fx

    Example: fhalls.eqg, obj_smplant01.mod

- Name: AlphaSModelCBGG1Max.fx

    Example: provinggrounds.eqg, mmlttm.mod

- Name: Chroma_MaxC1.fx

    Example: bazaar.eqg, obp_bannerdragon.mod

- Name: Chroma_MaxCB1.fx

    Example: bnr.eqg, obp_bnr_blue00.mod

- Name: Chroma_MaxCBS1.fx

    Example: ahf.eqg, ahf.mds

- Name: Chroma_MaxCBSG1.fx

    Example: bat.eqg, bat.mod

- Name: Chroma_MaxCBSGE1.fx

    Example: genericplate.eqg, it14157.mod

- Name: Chroma_MPLBasicA.fx

    Example: argath.eqg, obp_cropb_lod1.mod

- Name: Chroma_MPLBasicAT.fx

    Example: alkabormare.eqg, obj_dread_bannersm.mod

- Name: Chroma_MPLBumpA.fx

    Example: arena2.eqg, obj_arch_innerlt.mod

- Name: Chroma_MPLBumpAT.fx

    Example: arcstone.eqg, obj_fs_tent_.mod

- Name: Chroma_MPLGBAT.fx

    Example: furniture08.eqg, it20060.mod

- Name: Opaque_AddAlphaC1Max.fx

    Example: anguish.eqg, rasymbol.mod

- Name: Opaque_MaxC1_2UV.fx

    Example: bazaar.eqg, obj_cageb.mod

- Name: Opaque_MaxC1.fx

    Example: bazaar.eqg, obj_cageb.mod

- Name: Opaque_MaxCB1_2UV.fx

    Example: bazaar.eqg, obj_sackb.mod

- Name: Opaque_MaxCB1.fx

    Example: 10annvshield.eqg, it14000.mod

- Name: Opaque_MaxCBE1.fx

    Example: riftseekers.eqg, obj_scustat.mod

- Name: Opaque_MaxCBS_2UV.fx

    Example: gua.eqg, gua.mod

- Name: Opaque_MaxCBS1.fx

    Example: aam.eqg, aam.mds

- Name: Opaque_MaxCBSE1.fx

    Example: anguish.eqg, obj_dias01.mod

- Name: Opaque_MaxCBSGE1.fx

    Example: dodequip.eqg, it10817.mod

- Name: Opaque_MaxCBST2_2UV.fx

    Example: g05.eqg, g05.mod

- Name: Opaque_MaxCE1.fx

    Example: steamfactory.eqg, obj_pipe_elbow.mod

- Name: Opaque_MaxCG1.fx

    Example: anguish.eqg, clvgate.mod

- Name: Opaque_MaxCSG1.fx

    Example: dranikcatacombsa.eqg, obj_pedestalg01.mod

- Name: Opaque_MaxLava.fx

    Example: arcstone.eqg, obj_fs_flame_rock_c1.mod

- Name: Opaque_MaxSMLava2.fx

    Example: ldr.eqg, ldr.mod

- Name: Opaque_MaxWater.fx

    Example: anguish.eqg, obj_skin.mod

- Name: Opaque_MaxWaterFall.fx

    Example: ashengate.eqg, obj_crumblea.mod

- Name: Opaque_MPLBasic.fx

    Example: alkabormare.eqg, obj_everruina.mod

- Name: Opaque_MPLBasicA.fx

    Example: zonein.eqg, obj_libtelportpad.mod

- Name: Opaque_MPLBlend.fx

    Example: alkabormare.eqg, obj_everruina.mod

- Name: Opaque_MPLBlendNoBump.fx

    Example: beastdomain.eqg, obj_beastd_spike.mod

- Name: Opaque_MPLBump.fx

    Example: alkabormare.eqg, obj_dread_bannersm.mod

- Name: Opaque_MPLBump2UV.fx

    Example: arcstone.eqg, obj_fs_flame_rock_c1.mod

- Name: Opaque_MPLBumpA.fx

    Example: furniture19.eqg, it35125.mod

- Name: Opaque_MPLBumpAT.fx

    Example: breedinggrounds.eqg, obj_bg_frdragstatb.mod

- Name: Opaque_MPLFull.fx

    Example: drachnidhive.eqg, obj_glyph_voodo_wrym_a_.mod

- Name: Opaque_MPLFull2UV.fx

    Example: guildhall.eqg, obj_lamp_wall_05.mod

- Name: Opaque_MPLGB.fx

    Example: arcstone.eqg, obj_reliccard.mod

- Name: Opaque_MPLGB2UV.fx

    Example: convorteum.eqg, obj_convringsyel.mod

- Name: Opaque_MPLRB.fx

    Example: arthicrex.eqg, obj_temple_torch.mod

- Name: Opaque_MPLRB2UV.fx

    Example: crystallos.eqg, obj_uppercolearth.mod

- Name: Opaque_MPLSB.fx

    Example: arcstone.eqg, obj_ravenglass_.mod

- Name: Opaque_MPLSB2UV.fx

    Example: cityofbronze.eqg, obj_bronzecityentb.mod

- Name: Opaque_OpaqueRegionCBGG1Max.fx

    Example: dranikhollowsa.eqg, obj_urntall02_bk_00.mod

- Name: Opaque_OpaqueSkinMeshCBGG1Max.fx

    Example: omensequip.eqg, it10762.mod

- Name: Opaque_OpaqueSModelC1Max.fx

    Example: wallofslaughter.eqg, obj_tower02rev.mod

- Name: Opaque_OpaqueSModelCBGG1Max.fx

    Example: chambersf.eqg, obj_spine02.mod

- Name: OpaqueRegionC1Max.fx

    Example: fhalls.eqg, obj_statuepedestal01.mod

- Name: OpaqueRegionCB1Max.fx

    Example: fhalls.eqg, obj_statuepedestal01.mod

- Name: OpaqueSModelCB1Max.fx

    Example: draniksscar.eqg, rc_d_entry_b.mod

- Name: OpaqueSModelCBGG1Max.fx

    Example: fhalls.eqg, obp_geode01.mod

- Name: OpaqueSModelCG1Max.fx

    Example: provinggrounds.eqg, mmlrtm.mod
