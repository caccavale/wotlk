import {
	Consumes,
	Debuffs,
	EquipmentSpec,
	Flask,
	Food,
	Glyphs,
	IndividualBuffs,
	PartyBuffs,
	Potions,
	RaidBuffs,
	RaidTarget,
	TristateEffect
} from '../core/proto/common.js';
import {SavedTalents} from '../core/proto/ui.js';

import {
	RestorationDruid_Options as RestorationDruidOptions,
	RestorationDruid_Rotation as RestorationDruidRotation,
	DruidMajorGlyph,
	DruidMinorGlyph,
} from '../core/proto/druid.js';

import * as Tooltips from '../core/constants/tooltips.js';
import {NO_TARGET} from "../core/proto_utils/utils";

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/wotlk/talent-calc and copy the numbers in the url.
export const StandardTalents = {
	name: 'Phase 1',
	data: SavedTalents.create({
		talentsString: '5032003115331303213305311231--205003012',
		glyphs: Glyphs.create({
			major1: DruidMajorGlyph.GlyphOfFocus,
			major2: DruidMajorGlyph.GlyphOfInsectSwarm,
			major3: DruidMajorGlyph.GlyphOfStarfall,
			minor1: DruidMinorGlyph.GlyphOfTyphoon,
			minor2: DruidMinorGlyph.GlyphOfUnburdenedRebirth,
			minor3: DruidMinorGlyph.GlyphOfTheWild,
		}),
	}),
};

export const DefaultRotation = RestorationDruidRotation.create({
});

export const DefaultOptions = RestorationDruidOptions.create({
	innervateTarget: RaidTarget.create({
		targetIndex: NO_TARGET,
	}),
});

export const DefaultConsumes = Consumes.create({
	defaultPotion: Potions.RunicManaPotion,
	flask: Flask.FlaskOfTheFrostWyrm,
	food: Food.FoodFishFeast,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	bloodlust: true,
	divineSpirit: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	icyTalons: true,
	moonkinAura: TristateEffect.TristateEffectImproved,
	leaderOfThePack: TristateEffect.TristateEffectImproved,
	powerWordFortitude: TristateEffect.TristateEffectImproved,
	sanctifiedRetribution: true,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
	trueshotAura: true,
	wrathOfAirTotem: true,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	vampiricTouch: true,
});

export const DefaultPartyBuffs = PartyBuffs.create({
	heroicPresence: false,
});

export const DefaultDebuffs = Debuffs.create({
	bloodFrenzy: true,
	ebonPlaguebringer: true,
	faerieFire: TristateEffect.TristateEffectImproved,
	heartOfTheCrusader: true,
	judgementOfWisdom: true,
	shadowMastery: true,
	sunderArmor: true,
	totemOfWrath: true,
});

export const OtherDefaults = {
	distanceFromTarget: 18,
};

export const PRE_RAID_PRESET = {
	name: 'Pre-raid Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.fromJsonString(`{ "items": [
		{
			"id": 42554,
			"enchant": 3820,
			"gems": [
				41285,
				40049
			]
		},
		{
			"id": 40680
		},
		{
			"id": 37673,
			"enchant": 3810,
			"gems": [
				42144
			]
		},
		{
			"id": 41610,
			"enchant": 3859
		},
		{
			"id": 39547,
			"enchant": 3832,
			"gems": [
				42144,
				40026
			]
		},
		{
			"id": 37884,
			"enchant": 2332,
			"gems": [
				0
			]
		},
		{
			"id": 39544,
			"enchant": 3604,
			"gems": [
				42144,
				0
			]
		},
		{
			"id": 40696,
			"enchant": 3601,
			"gems": [
				40014,
				39998
			]
		},
		{
			"id": 37854,
			"enchant": 3719
		},
		{
			"id": 44202,
			"enchant": 3606,
			"gems": [
				39998
			]
		},
		{
			"id": 40585
		},
		{
			"id": 43253,
			"gems": [
				40026
			]
		},
		{
			"id": 37873
		},
		{
			"id": 40682
		},
		{
			"id": 45085,
			"enchant": 3834
		},
		{
			"id": 40698
		},
		{
			"id": 40712
		}
	]}`),
};

export const P1_PRESET = {
	name: 'P1 Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 40467,
			"enchant": 3820,
			"gems": [
				41285,
				42144
			]
		},
		{
			"id": 44661,
			"gems": [
				40026
			]
		},
		{
			"id": 40470,
			"enchant": 3810,
			"gems": [
				42144
			]
		},
		{
			"id": 44005,
			"enchant": 3859,
			"gems": [
				40026
			]
		},
		{
			"id": 40469,
			"enchant": 3832,
			"gems": [
				42144,
				39998
			]
		},
		{
			"id": 44008,
			"enchant": 2332,
			"gems": [
				39998,
				0
			]
		},
		{
			"id": 40466,
			"enchant": 3604,
			"gems": [
				39998,
				0
			]
		},
		{
			"id": 40561,
			"enchant": 3601,
			"gems": [
				39998
			]
		},
		{
			"id": 40560,
			"enchant": 3719
		},
		{
			"id": 40519,
			"enchant": 3606
		},
		{
			"id": 40399
		},
		{
			"id": 40080
		},
		{
			"id": 40255
		},
		{
			"id": 40432
		},
		{
			"id": 40395,
			"enchant": 3834
		},
		{
			"id": 40192
		},
		{
			"id": 40321
		}
	]}`),
};