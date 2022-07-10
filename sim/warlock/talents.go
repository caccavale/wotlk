package warlock

import (
	"time"

	"github.com/wowsims/wotlk/sim/core"
	//	"github.com/wowsims/wotlk/sim/core/proto"
	"github.com/wowsims/wotlk/sim/core/stats"
)

func (warlock *Warlock) ApplyTalents() {
	// demonic embrace
	if warlock.Talents.DemonicEmbrace > 0 {
		bonus := 1.01 + float64(warlock.Talents.DemonicEmbrace)*0.03
		warlock.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Stamina,
			ModifiedStat: stats.Stamina,
			Modifier: func(in float64, _ float64) float64 {
				return in * bonus
			},
		})
	}

	// Suppression
	warlock.AddStat(stats.SpellHit, float64(warlock.Talents.Suppression)*core.SpellHitRatingPerHitChance)

	// Add 1% crit per level of backlash.
	warlock.AddStat(stats.SpellCrit, float64(warlock.Talents.Backlash)*core.CritRatingPerCritChance)

	// fel intellect
	if warlock.Talents.FelVitality > 0 {
		bonus := (0.01) * float64(warlock.Talents.FelVitality)
		// Adding a second 3% bonus int->mana dependency
		warlock.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Mana,
			Modifier: func(intellect float64, mana float64) float64 {
				return mana + intellect*15*bonus
			},
		})
		// //  TODO: fel stamina increases max health (might be useful for warlock tanking sim)
		// warlock.AddStatDependency(stats.StatDependency{
		// 	SourceStat:   stats.Stamina,
		// 	ModifiedStat: stats.Health,
		// 	Modifier: func(intellect float64, mana float64) float64 {
		// 		return Health + Stamina*10*bonus
		// 	},
		// })
	}

	warlock.PseudoStats.BonusCritRating += float64(warlock.Talents.DemonicTactics) * 1 * core.CritRatingPerCritChance

	// if !warlock.Options.SacrificeSummon && warlock.Options.Summon != proto.Warlock_Options_NoSummon {
	// 	if warlock.Talents.MasterDemonologist > 0 {
	// 		switch warlock.Options.Summon {
	// 		case proto.Warlock_Options_Imp:
	// 			warlock.PseudoStats.ThreatMultiplier *= 0.96 * float64(warlock.Talents.MasterDemonologist)
	// 		case proto.Warlock_Options_Succubus:
	// 			warlock.PseudoStats.DamageDealtMultiplier *= 1.0 + 0.02*float64(warlock.Talents.MasterDemonologist)
	// 		case proto.Warlock_Options_Felgaurd:
	// 			warlock.PseudoStats.DamageDealtMultiplier *= 1.0 + 0.01*float64(warlock.Talents.MasterDemonologist)
	// 			// 		Felguard - Increases all damage caused by 1% and all resistances by .1 per level.
	// 			// 		Voidwalker - Reduces physical damage taken by 2%.
	// 			// 		Felhunter - Increases all resistances by .2 per level.
	// 		}
	// 	}

	// 	if warlock.Talents.SoulLink {
	// 		warlock.PseudoStats.DamageDealtMultiplier *= 1.05
	// 	}

	// 	// Extract stats for demonic knowledge
	// 	petChar := warlock.Pets[0].GetCharacter()
	// 	bonus := (petChar.GetStat(stats.Stamina) + petChar.GetStat(stats.Intellect)) * (0.04 * float64(warlock.Talents.DemonicKnowledge))
	// 	warlock.AddStat(stats.SpellPower, bonus)
	// }

	// // demonic tactics, applies even without pet out
	// warlock.AddStats(stats.Stats{
	// 	stats.MeleeCrit: float64(warlock.Talents.DemonicTactics) * 1 * core.CritRatingPerCritChance,
	// 	stats.SpellCrit: float64(warlock.Talents.DemonicTactics) * 1 * core.CritRatingPerCritChance,
	// })

	if warlock.Talents.Nightfall > 0 {
		warlock.setupNightfall()
	}

	if warlock.Talents.ShadowEmbrace > 0 {
		warlock.setupShadowEmbrace()
	}

}

func (warlock *Warlock) setupShadowEmbrace() {
	warlock.ShadowEmbraceAura = warlock.RegisterAura(core.Aura{
		Label:     "Shadow Embrace",
		ActionID:  core.ActionID{SpellID: 32391},
		Duration:  time.Second * 12,
		MaxStacks: 3,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.PeriodicShadowDamageDealtMultiplier /= 1.0 + 0.01*float64(warlock.Talents.ShadowEmbrace)*float64(oldStacks)
			aura.Unit.PseudoStats.PeriodicShadowDamageDealtMultiplier *= 1.0 + 0.01*float64(warlock.Talents.ShadowEmbrace)*float64(newStacks)
			// TO DO : Healing over time reduction part
		},
	})

	warlock.RegisterAura(core.Aura{
		Label: "Shadow Embrace Talent",
		//		ActionID: core.ActionID{SpellID: 32394},
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spell == warlock.Shadowbolt { // TODO: also works on Haunt
				warlock.ShadowEmbraceAura.Activate(sim)
				warlock.ShadowEmbraceAura.AddStack(sim)
			}
		},
	})
}

func (warlock *Warlock) setupNightfall() {
	warlock.NightfallProcAura = warlock.RegisterAura(core.Aura{
		Label:    "Nightfall Shadow Trance",
		ActionID: core.ActionID{SpellID: 17941},
		Duration: time.Second * 10,
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			// Check for an instant cast shadowbolt to disable aura
			if spell != warlock.Shadowbolt || spell.CurCast.CastTime != 0 {
				return
			}
			aura.Deactivate(sim)
		},
	})

	warlock.RegisterAura(core.Aura{
		Label: "Nightfall",
		// ActionID: core.ActionID{SpellID: 18095},
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnPeriodicDamageDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spell != warlock.Corruption { // TODO: also works on drain life...
				return
			}
			if sim.RandomFloat("nightfall") > 0.02*float64(warlock.Talents.Nightfall) {
				return
			}
			warlock.NightfallProcAura.Activate(sim)
		},
	})
}

func (warlock *Warlock) applyNightfall(cast *core.Cast) {
	if warlock.NightfallProcAura.IsActive() {
		cast.CastTime = 0
	}
}
