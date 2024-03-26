package seed

import (
	"fmt"
	"strings"
)

type Characteristic struct {
	BaseDesc string
	ValDesc  map[string]string
}

type Entry struct {
	Name, Description string
}

type Clan struct {
	Name           string
	Description    string
	Appearance     string
	AssociatedSect []string
	Haven          string
	Background     string
	Character      string
	Discipline     []Discipline
	Weakness       string
	Organizations  string
	Strongholds    []string
	SubClan        []Clan
	IsHighClan     bool
}

type Sect struct {
	Name               string
	Description        string
	Titles             []Entry
	Practices, Rituals string
	Strongholds        []string
}

type Traditions struct {
	Name        string
	Description string
	Tradition   []Entry
}

type Discipline struct {
	Name        string
	Description string
	Abilities   map[string]disciplineAbility
}

type disciplineAbility struct {
	BaseDesc string
	Lvl      int
	System   map[string]string
}

var ClanMap = map[string]Clan{
	"Assamite": {
		Name: "Assamite",
		Description: `The childer of Haqim, known as Assamites to the
		rest of the Kindred, are a silent knife in the dark, an
		order of bloodthirsty assassins who participate in the
		secret wars of the undead by operating as killers for
		hire. Outside the purview of the Sects, the Assamites
		are true independents and mercenaries, hiring out to
		whoever can pay their blood-price and ungoverned by
		the will of Prince or Priscus.n truth, the Assamites are more than simple thugs and killers. Theirs is a complex but insular Clan predi-
		cated upon the three principles of wisdom, sorcery, and
		diablerie. Most Assamites that other vampires encounter
		are members of the warrior caste, however, so Kindred
		society has painted them all with that brush. For their
		part, the Assassins have done nothing to stop this misun-
		derstanding. If it helps them acquire contracts and it oc-
		cludes the true nature of their Clan, the better for them.
		`,
		Appearance: `Older Assamites often come from Middle
		Eastern and North African cultures, though more and
		more young Assamites come from a wider demographic. In traditional
		environments, the Assamites prefer garb appropriate to religious or Clan custom. When in public, however, Assamites wear whatever the locals do, allowing them to fulfill their contracts without anyone no-
		ticing anything amiss. An Assamite’s skin grows darker with age (as opposed to other vampires, whose skin gets paler); particularly ancient
		Assamites are almost ebony in complexion.`,
		Background: ` Those Embraced into Clan Assamite tend
		to fall into two distinct types: The “provincial” members of
		the Clan fit whatever their locality is, and can blend seam-
		lessly in with the people around them. The higher-profile
		“jet-setters” transcend cultures, bolstered by their ability
		to handle interpersonal and intellectual challenges.`,
		Character: `Physical Attributes tend to be
		primary, with some Assamites favoring Social Attri-
		butes to help them get close to their prey. Talents and
		Skills are equally favored, but Knowledges may help
		the wise Assamite in a pinch. Few Assassins cultivate
		extensive Backgrounds, and instead specialize in an
		array of Disciplines that heighten their competence.
		The most accomplished Assamites follow the Clan’s
		unique Path of Enlightenment, and those who don’t
		often have to spend a great deal of effort maintaining
		their Virtues and Humanity`,
		AssociatedSect: []string{
			"Independent",
		},
		Haven: `Assamites often share communal havens
		with others of their local cell, remote structures that
		allow the Assassins to watch the larger domain from a
		distance. These havens are generally well appointed,
		but not so lavish that the whole place can’t be moved
		on short notice. Individual Assamites also tend to keep
		personal hideouts of a much more humble nature, for
		when they need a place to lay low.`,
		Weakness: `Due to the Tremere blood-curse,
		should an Assamite consume the blood of another
		Kindred, she suffers one automatic level of unsoakable lethal damage per blood point imbibed. Diablerie attempts result in automatic aggravated damage, one health level per point of permanent Willpower the victim possesses; the would-be diablerist gains no benefits
		(including Generation reduction) if he survives the process. In addition, Assamites must tithe some of the profits from their contracts to their sires or superiors (generally around 10 percent of all such earnings).`,
		Discipline: []Discipline{
			DisciplineMap[`celerity`], DisciplineMap[`obfuscate`], DisciplineMap[`quietus`],
		},
		Organizations: `An insular, hierarchical organization shapes
		much of Assamite custom. “The Old Man on the Mountain”
		— the master assassin who makes his haven in the mountain fortress of Alamut — is the ultimate authority, and the Clan heeds the orders that trickle down to them with a mix of reverence and terror. Individual and local cells of Assamites known as falaqi frequently have license to act with autonomy, but “turncoats” against the higher cause are rare.`,
		IsHighClan: true,
	},
	"FollowersofSet": {
		Name: "Follower of Set",
		AssociatedSect: []string{
			"Camarilla","Sabbat",
		},
		Description: `Addiction, debasement, corruption, and desperation strike fear into many Kindred who worry that theirs will become an unlife of ruin, but to the Followers of Set, these and more are the tools of the trade. Pimps, pushers, and
		priests, the Setites cater to
		the needs of the desperate,
		and convert them to a ni-
		hilistic cause in doing
		so.he Followers of Set are as much a chthonic reli-
		gion as they are a Clan, though the faith includes the
		Clan. Its mythology is complex and convoluted, an im-
		penetrable pantheon of god-monsters. At the apex of
		this worship stands a syncretism of the Egyptian Lord
		of the Underworld Set and the Greek hydra Typhon,
		as much spiritual guardians of secret places as they are
		the “liberators” of other’s souls`,
		Appearance: `Many older Setites hail from the
		North African and Mediterranean ethnicities native to
		the Serpents’ historical territory, but they freely Em-
		brace from among the mortals of their adopted homes.
		Some long-standing Setite temples are tied to loca-
		tions where “Egyptian” Serpents might seem out of
		place, but where some aspect of serpent mythology is
		53VAMPIRE THE MASQUERADE 20th ANNIVERSARY EDITION
		present, as in Mesoamerica or even far-flung Nordic lo-
		cales, and thus draw their membership from local popu-
		lations. Red hair is considered a mark of Set’s favor.
		`,
		Background: ` Prospective childer for the Followers of
		Set often spend some time involved with a Setite cult,
		so they’re indoctrinated in the mysteries of the Clan
		before becoming one of its Kindred. They may come
		from any cultural origin, though many are outsiders,
		loners, or otherwise marginalized by society, which is
		often what led them to the forbidden fruits offered by
		the cult of Set in the first place.`,
		Character: `Sires choose those who demon-
		strate acumen in Social and Mental Attributes, as the
		Clan needs its proselytizers and its priests to be charismatic
		and quick-witted. Knowledges are almost always primary,
		though Talents may be so instead, especially among those
		Serpents engaged in dealings with others. Setites focus
		on Backgrounds that give influence over others whether
		subtly or overtly, making Allies, Contacts, Influence, Re-
		sources, and sometimes Retainers popular.`,
		Haven: `Where their hidden temples stand, the Set-
		ites make their havens, either individually or commu-
		nally. These may be anything from “churches” with
		never-before-heard-of denominations or they may be
		outright cults that have to hide their existences. The
		secretive Serpents sometimes hide individual havens
		in places where other Kindred don’t often go, such as
		insular ethnic neighborhoods, abandoned domains,
		“the rough part of town,” and so forth. Some Setites
		also haven in secret mystical places that have value to
		the Clan, guarding them from outsiders.`,
		Discipline: []Discipline{
			DisciplineMap[`presence`], DisciplineMap[`obfuscate`], DisciplineMap[`serpentis`],
		},
		Weakness: `Given their origins in darkness, the Ser-
		pents react negatively to bright light: Add two health
		levels to damage caused by exposure to sunlight. Setites
		also lose one die from dice pools for actions taken in
		bright light (police spotlights, stage lights, flares, etc.).`,
		Organizations: `mple or network of cults representing
		a city’s Serpent presence. Setites in the cities rarely
		scheme among each other, preferring to face outward
		threats in unity rather than the schismatic self-interest
		of the Sects. Whispers persist of a massive temple devoted to Set located somewhere in Africa, with a terrible Methuselah who claims to be the childe of Set himself at the head. If this is true, then the Clan’s higher agenda probably originates here, but the Setites
		themselves remain silent on the topic.`,
		IsHighClan: true,
	},
	"Brujah": {
		Name: "Brujah",
		Description: `The legacy of the Brujah is one of halcyon greatness
		marred by their own fiery natures. Theirs was the glory
		of ancient Carthage, but Ventrue treachery in ancient
		Rome brought the dream to an end. Since then, the
		Brujah have borne a grudge. Tonight, the Brujah
		are rebels and
		provocateurs,
		bat-swinging
		hooligans and
		agents of change in a society long crippled by stasis.More so than any other Clan, the Brujah still feel the
		flames of the passions that once inspired them as mortals.
		Clan Brujah loves a cause and is quick to act on a stirring
		speech, accusation of injustice, or a call to arms. This
		connection to passion can be a blessing, but inspiration
		can also yield to the madness and hunger of the Beast.
		`,
		Appearance: `Many Brujah affect styles and manner-
		isms that reflect an attitude of rebellion. Multicolored
		hair, shaven heads, spikes, rivets, fetish gear, and t-shirts
		with bold slogans might appeal to a Brujah. While not
		every Brujah wears the “uniform,” the Rabble often enjoy
		adorning themselves in outfits intended to provoke. Some
		young Brujah prefer mobile devices as their tools of resis-
		tance, and can summon a riot at the touch of a screen.`,
		Background: `As creatures of passion, Brujah often
		Embrace without really thinking much about it, and
		their childer tend to be a disparate lot. Sometimes, the
		Rabble Embrace those who share a similar outlook or
		enthusiasm for a cause as the prospective sire. Other
		times, they inflict the Embrace on those of opposite ide-
		ology, cursing a rival with vampirism as punishment.`,
		Character: `Brujah are usually — though
		not exclusively — drawn from mortals with violent
		or ungovernable personalities. Their Natures and De-
		meanors are often similar, as Brujah have little use for
		guile. They lean toward physical Attributes, with So-
		cial and Mental about equal afterwards. Many favor
		Skills and Talents, but most respect the Knowledges
		that make them more than just thugs. Contacts, Allies,
		and Herd are common Backgrounds.`,
		AssociatedSect: []string{
			"Camarilla","Sabbat",
		},
		Haven: `Brujah may feel kinship to a city, but they rare-
		ly develop such ties to individual locations. Thus, at any
		given time, a Brujah probably has a half-dozen or more
		hideouts, safehouses, and flats available. These are often
		shabby and ill-kept until the Brujah needs them. Brujah
		havens might also have mortals who follow the Brujah’s
		ideology or his cult of personality. This works out fine: It
		never hurts to have a spare vessel in an emergency.`,
		Weakness: `The same passions that inspire Bru-
		jah to greatness or depravity, left unchecked, can send
		them into incandescent rages: The difficulties of rolls
		to resist or guide frenzy (p. 298) are two higher than
		normal. Additionally, a Brujah may never spend Will-
		power to avoid frenzy, though he may spend a point of
		Willpower to end a frenzy that has already begun.`,
		Discipline: []Discipline{
			DisciplineMap[`celerity`], DisciplineMap[`potence`], DisciplineMap[`presence`],
		},
		Organizations: `Certain causes rise and fall in Brujah
		fashion, but some of the more tenured are those who call
		themselves Idealists and Iconoclasts. Iconoclasts want to
		tear it all down, while Idealists enjoy solving problems
		through theory. This last harkens somewhat to the classical roots of the Brujah as philosopher-kings, and most Idealists are among the ranks of Brujah ancillae and elders.`,
		IsHighClan: true,
	},
	"Gangrel": {
		Name:        "Gangrel",
		Description: `A glint of red eyes in the darkness, the scent of a predator’s musk, a flash of fangs, the sound of flesh tearing: These mark the presence of the Gangrel. More than any other Clan, the Gangrel resemble the beasts associated with the legends of vampires: bats, wolves, and other creatures of darkness. Indeed, the Outlanders may develop the ability to transform themselves into these and other, more primal forms. Most are tough and, when pressed, ferocious. And when Gangrel succumb to the depredations of the Beast, they are left with some feature redolent of the animal kingdom.Theirs is a tense relationship with vampire society, and Outlanders are among the most frequent to turn Anarch or Autarkis. In some localities, the Gangrel have collectively abandoned membership in any Sect`,
		Appearance: `Personal presentation is often not high
		on the list of many Gangrel priorities, and a Gangrel’s
		appearance is often more a matter of circumstance than it is of
		active decision. The Clan’s weakness can contribute a great deal to
		their appearance, as does an extended unlife in the places where
		they make their havens, which are frequently short of modern conveniences`,
		AssociatedSect: []string{
			"Independent", "Sabbat", "Camarilla", "Anarch Movement",
		},
		Haven: `Gangrel often lair where they can, taking
		refuge when the sun threatens to rise. Those who do
		maintain permanent havens often lean toward the util-
		itarian: Everything from a cave to a covered alley to an
		illegal squat may serve as a Gangrel haven, usually with
		little demarcating them as any sort of personal territory
		(until it’s too late for the unfortunate interloper).`,
		Background: `Gangrel sire childer like they seek prey:
		 after long hunts during which the prospective childe
		 doesn’t even know she’s being followed. Creating a
		 fledgling means sharing limited resources, so each sire-
		 childe relationship is unique and significant. Outland-
		 ers Embrace because they choose an individual, not
		 out of whim or recklessness. Those who earn their at-
		 tention are hardy, whether physically or emotionally.`,
		Character: `Sufficiency is the Gangrel hallmark, and many have outsider or loner personality archetypes. Physical Attributes are far and away most
		frequent, as are Talents with a smattering of Skills and Knowledges. Many Gangrel focus on Disciplines rather than Backgrounds, preferring to rely on themselves more than others. Gangrel almost never have significant Resources, Influence, or Retainers.`,
		Discipline: []Discipline{
			DisciplineMap[`animalism`], DisciplineMap[`fortitude`], DisciplineMap[`protean`],
		},
		Weakness: `Every time a Gangrel frenzies, she
		acquires a temporary animal Characteristic (which
		may replace an existing temporary one). A patch of
		fur, a brief torpor after feeding, or skittishness around
		crowds — all of these may mar an Outlander after fren-
		zy. Characteristics acquired in Gangrel frenzies need
		not only be physical – they can be behavioral as well.
		Players should work with the Storyteller to determine
		what new animal trait is acquired (whether the frenzy
		involved the fight-or-flight impulse may be relevant).
		Over time, or in an exceptional situation, a particular animal feature may become permanent, with the
		next frenzy adding a new feature.`,
		Organizations: `Regional groups of Gangrel occasion-
		ally assemble in convocations that draw from ethnic or
		cultural influences. These are informal affairs, geared
		more toward sharing information and revelry than advancing any cogent agenda. Aside from these infrequent gatherings, almost all Gangrel organization is very local where it exists at all, from pairs of sire-and-childe through terrifying packs centered around one
		accomplished Outlander.`,
		IsHighClan: true,
	},
	"Giovanni": {
		Name:        "Giovanni",
		Description: `Achieving prominence during the Venetian Renaissance, the Giovanni family built their fortune on the rise of the middle class and the ready profit of banking and Mediterranean trade. However, with the family’s rise came hubris, as its paterfamilias sought ever more power, and with that hubris came horror. To this night, the Giovanni are known for the insular nature of their Clan and the incestuous practices by which they populate it. A few outside families and factions fall under Giovanni auspices, but the vast majority of the Clan comes from the debased mortal family.For the most part, the Giovannni participate little in the Jyhad, pursuing their own agenda of cultivating wealth and building a foundation of power in the lands beyond the veil of death. Outsiders rarely comprehend the goals of the Necromancers, but only the most trusted of the Giovanni know that the Clan wants to plunge the world into a state where the dead and the living co-mingle.`,
		Appearance: ` Outwardly, Giovanni dress with subtlety and
		taste. Much of the Clan comes from the original mortal family,
		and have not only olive Italian complexions, but some amount
		of inherited family features. Those outside the immediate
		family often appear “of a type,” and in the traditional garb of
		their regional family branch.`,
		AssociatedSect: []string{
			"Independent",
		},
		Haven: `The family wealth of the Giovanni is evident in their havens, which may take the form of villas or lavish estates. The Necromancers often have valuables invested in their havens as well, such as galleries of fine art or displays of jewelry. Many Giovanni also maintain secondary havens, where they may have elaborate necromantic crypts or just flats where they can lie low if necessary.
		`,
		Background: `Giovanni of the main family branch have usually spent some amount of time as a ghoul in a practice known as the Proxy Kiss. During this time, the Kindred-to-be learns about the treacherous and jealous  reality of the vampire family. She learns ambition and an unhealthy dose of duplicity, in addition to the family history and customs. Giovanni rarely see much of the outside world on their own terms during the Proxy Kiss period, and often become insular and alienated, while at the same time eager to stand out among the other family ghouls.`,
		Character: `Giovanni vampires often have outgoing, professional Demeanors that hide unpleasant Natures warped in their upbringing. Social or Mental
		Attributes are usually primary, though some of the “soldiers” of the family prefer Physical Attributes. Emphasis likewise usually falls on primary Knowledges or Talents, depending on proclivity. A split in the Clan sees those who favor the practicality of Backgrounds (particularly
		those tied to wealth and exerting influence) diverge from
		those who prefer the forbidden puissance of Disciplines.
		Few Giovanni could be described as well-rounded.`,
		Discipline: []Discipline{
			DisciplineMap[`dominate`], DisciplineMap[`potence`], DisciplineMap[`necromancy`],
		},
		Weakness: `The Kiss of a Giovanni vampire causes
		excruciating pain in mortal vessels who receive it. If
		the Giovanni isn’t careful, her vessel may die of shock
		and agony before being wholly exsanguinated. When a
		Giovanni feeds upon a mortal, she does twice as much
		damage as the Kiss of another vampire would inflict. For
		example, if a Giovanni takes one point of blood from a
		mortal vessel, that victim would suffer two health lev-
		els of damage. As a result, they tend to use blood banks
		and other means of feeding that don’t fight as much.`,
		Organizations: `Like few other Clans, the Necromancers have a top-down organization where policy is made by a (presumed) still-active Clan progenitor, Augustus. The family maintains an enormous palazzo known as the Mausoleum in Venice, where elders and fledglings alike dance to the whims of their ancient puppetmasters. Clan structure is itself familial, with the added complications that degeneracy and vampirism offer. Incest, ancestor worship, necrophilia, cults of guilt, and bizarre relationships in which fathers and grandmothers are their own issue’s childer make a mire of the Clan and family, and fracture many Giovanni long before they become Kindred`,
		IsHighClan:    true,
	},
	"Lasombra": {
		Name: "Lasombra",
		Description: ` To the mind of a Lasombra, it is better to reign in
		hell than to serve in heaven. Fear, frenzy, the power
		to determine whether another lives or dies: these are
		at the root of the power that the Lasombra hold dear.
		Whereas other vampires try vainly to hold the Beast at
		bay or give themselves wholly to it, the Lasombra beat
		the Beast into submission,
		invoking it when it
		suits them but leav-
		ing it trapped inside
		when they wish to
		govern themselves. the Clan was instru-
		mental not only in establishing many of the rituals of
		the Sabbat, but in the institutions that keep it from
		descending into chaos each night.
		Whether they see themselves as God’s instruments
		or as outcasts from His creation, the Lasombra believe
		they have a duty (whether to Sect, Clan, pack, or even
		just themselves), and obligation to their responsibilities
		gives them a wicked sense of purpose.`,
		Appearance: `he Lasombra are frequently attractive.
		Whether through the Spanish, Italian, and Moorish
		stock associated with the Clan, or due to more cosmopolitan modern backgrounds, the Keepers cut a striking figure. Their dress is often conservative or religious,
		drawing on years of ceremony and faithful ritual. Rare
		is the Lasombra who cannot at least manipulate shadows to affect a dramatic entrance or enigmatic pose.`,
		AssociatedSect: []string{
			"Sabbat",
		},
		Haven: ` The obligation of their leadership leads many
		young Lasombra to maintain communal havens with other members of their pack. Wealthier Keepers and
		those who predate the Sect often maintain their own
		havens, whether sinister penthouse suites or sprawling
		Old World villas.`,
		Background: `Prospective sires of Clan Lasombra
		 seek both erudition and ambition in their potential
		 childer. As such, many Lasombra come from profes-
		 sional backgrounds, and display outgoing and even
		 aggressive personalities. Merit in their sires’ eyes takes
		 fledgling Lasombra far, and the Keepers do not hesitate
		 to cull their ranks of flawed, lazy, or boorish childer.`,
		Character: `Lasombra often have disparate
		Natures and Demeanors. Mental or Social Attributes
		are equally likely to be primary. Attributes tend to be
		narrow and specialized, showing individual expertise.
		Keepers cultivate Backgrounds of all types similarly
		in (initially) low quantities, to better diversify them-
		selves`,
		Discipline: []Discipline{
			DisciplineMap[`dominate`], DisciplineMap[`potence`], DisciplineMap[`obteneration`],
		},
		Weakness: `Lasombra vampires cast no reflections.
		Whether in a mirror, in a body of water, on a polished
		surface, or in the rear-view of a taxicab, the image of
		the Keeper does not reflect.`,
		Organizations: `For the Lasombra, the nights of high
		aristocracy never faded, and the titles and offices a
		modern onlooker might associate with history still car-
		ry great weight among the Keepers. A complex system
		of patronage, mentorship, and lineage characterizes
		the Clan, similar to the courts and churches of cen-
		turies gone by. Childer benefit greatly from esteemed
		sires and vice versa, while acts that confer Sect or Clan
		status may also elevate a Lasombra’s peers, so long as
		she associates her success with her fellows.`,
		IsHighClan: true,
	},
	"Malkavian": {
		Name: "Malkvian",
		Appearance: `While Malkavians can come from any
		culture, the madness following the Embrace tends to
		lead them to extremes of self-presentation. Malkavians
		may appear disheveled, injured, or simply dirty. They
		could still be wearing the same clothes from the night
		of their Embrace or they may have stolen clothes from
		a laundromat or a department store during a fit of con-
		fusion or fugue. Of course, Malkavians are just as likely
		to be meticulous and exacting in their appearance, trying obsessively to appear as normal as possible.`,
		Description: `Clan Malkavian is twice damned: once by the curse
		of being Kindred, and again by the turmoil that dis-
		turbs their hearts and minds. Upon the Embrace, ev-
		ery Malkavian is afflicted with an insurmountable
		insanity that fractures her outlook for every night
		thereafter, making her unlife one of madness.
		Some consider this a form of oracular insight,
		while others simply consider them dangerous.Make no mistake: Malkavian insanity is a
		painful, alienating phenomenon, but it occa-
		sionally provides the Lunatics with bursts of in-
		sight or heretofore unknown perspective. Madness
		for the Malkavians may take the form of any clinical
		form of insanity, or it may be a hyperacuity of senses
		others don’t know they have; a supernatural puppe-
		teer pulling the Malkavian’s strings, or a sense that
		the Malkavian is somehow ahead of evolutionary
		schedule. Their precarious stability makes it hard for other
		Kindred (or, indeed, any vessels with whom they may
		meet) to interact with Malkavians. `,
		AssociatedSect: []string{
			"Camarilla",	},
		Haven: `Consistency is rare among Malkavians.
		Quite simply, they establish havens where they think
		to, where they can, and where they can recall. A sig-
		nificant number of Malkavians have literally no home,
		spending each night where exhaustion or the sun’s rays
		leave them. Others may permanently have the presi-
		dential suite in a posh hotel, a squat in the Barrens,
		the dispensary at a county jail, or a broom closet in a
		historical landmark.
		`,
		Background: `Malkavians Embrace with all the caprice
		 one would assume from them. Lunatic childer come from
		 all economic and cultural strata, though most have
		 some sort of hard-luck story or black secret behind
		 them that caused their sire to take note. Truly dam-
		 aged Malkavians who are unaware of the meanings
		 of their actions may not even be aware that they
		 have sired childer, which makes for very difficult
		 entry into Kindred society for these castoffs, many
		 of whom end up among the Caitiff.`,
		Character: `Loner, outsider, and deviant
		concepts and archetypes prevail among the Malka-
		vians, as do Mental Attributes (with an occasional
		Social-primary madman or Physical-primary maniac
		hiding among the ranks). Talents and Knowledges
		are likely most popular among Malkavians. Backgrounds
		tend to be either broad and shallow or few and deep, rep-
		resenting the way the way the Malkavian caroms through
		unlife. Virtues, Humanity, and Paths often tumble quick-
		ly, usually in the wake of Willpower doing the same.
		`,
		Discipline: []Discipline{
			DisciplineMap[`auspex`], DisciplineMap[`dementation`], DisciplineMap[`obsfucate`],
		},
		Weakness: `All members of Clan Malkavian suffer from
		a permanent, incurable derangement. They may acquire and
		recover from other derangements, and may spend Willpower
		to ameliorate the effects of the derangement for a scene, but
		they can never recover from their original derangement.
		`,
		Organizations: ` Rumor is more widespread than truth with
		regard to Malkavian organization, and any number of har-
		rowing tales describe the function of the Clan. Some say
		the Lunatics all share a hive consciousness; others claim
		that this is in fact the lingering consciousness of the Clan’s
		progenitor Malkav. Still others claim that a network of
		madness puts all Malkavians in contact with one anoth-
		er and is the cause of their crippling insanity. If nothing
		else, the Malkavians prove inscrutable and uncanny. One night, each of them frenzies when they see one another,
		while the next night, they all converge at the same time
		at the Sheriff’s haven and accuse him of harboring Sabbat
		spies.`,
		IsHighClan: true,
	},
	"Nosferatu": {
		Name: "Nosferatu",
		Appearance: `Physical horror is the lot of the Nosfer-
		atu, and their unsettling deformations are countless.
		No two Nosferatu share the exact same malforma-
		tion, and the Clan is a freakshow of snarled limbs,
		fanged protrusions, hellish countenances, serpen-
		tine spines, ruined faces, spasmodic appendages, and
		even features not usually seen on the mortal stock
		from which the Nosferatu are drawn. The Sewer
		Rats often hide these disfigurements under shape-
		less robes and rags, but some exult in the discomfort
		their presence causes, and don’t bother disguising
		them. They may even emphasize them`,
		Description: `Those who doubt that the Embrace is a curse need
		look no further than the Nosferatu. Twisted by the
		mark of Caine, members of Clan Nosferatu are warped
		by the Embrace into hideous monsters. As such, they
		skulk and keep to the shadows, and they often rouse the
		ire and mockery of other Kindred for their nightmarish appearances. Still others are so terrified or revolted
		by the Nosferatu that these warped Kindred have little
		social interaction at all. To their credit, the Nosferatu come
		to possess many of the whispered secrets of their reluctant fellows. The Sewer Rats enjoy a grudging respect as
		the information-brokers of the Kindred, given their supernatural acumen at stealth and the fact that many Kindred
		would rather ignore them than acknowledge them.On the whole, the Nosferatu condition is lonely and
		alienating. How they react to the Curse of Caine varies with their outlook and mental stamina, but it’s hard
		to be an object of utter revulsion and not let it shape
		one’s disposition toward one’s “Kindred” in some way.`,
		AssociatedSect: []string{
			"Camarilla","Autarkis", "Anarchs",
		},
		Haven: `Nosferatu Kindred often make their havens
		far from the scorn and spite of other vampires. Whether they construct warrens in the sewers suggested by their
		nickname or they sculpt a sprawling nightmare-nest in
		the spire of a condemned church, Sewer Rats value secrecy and distance from rivals in their havens. Nosferatu of humbler means may well squat in an abandoned
		tenement or a disused alley. So long as it’s away from
		other Kindred, it’s a good haven.`,
		Background: `The Sewer Rats mostly fit into one of
		 two categories. Some Nosferatu Embrace the damaged,
		 flawed, outcast, or vile, feeling some degree of kinship
		 with them. Other Embrace spitefully, dragging the
		 beautiful or privileged into an immortal hell of disfigurement and monstrosity`,
		Character: `Many Nosferatu come from
		loner and outsider concepts, Natures, and Demeanors.
		They favor Physical and Mental Attributes over Social, and they tend to be balanced in their regard for
		Talents, Skills, and Knowledges. Nosferatu prefer to
		specialize rather than generalize in Backgrounds, favoring those that carry favors and information like Allies,
		Contacts, Mentor, and even a bit of Influence`,
		Discipline: []Discipline{
			DisciplineMap[`animalism`], DisciplineMap[`obfuscate`], DisciplineMap[`potence`],
		},
		Weakness: `All Nosferatu have an Appearance
		score of zero, and they may never improve it. Cross it
		off the character sheet. Dice pools that use the Ap-
		pearance Trait are inherently difficult for these hideous
		Kindred.
		`,
		Organizations: `Occasionally, a coterie of Nosferatu
		becomes a brood or cult, often based around a remote
		warren. These can sprawl into vast “kingdoms” of Sew-
		er Rats, who often exist without the knowledge of a do-
		main’s Prince or Archbishop. The Nosferatu are among
		the Kindred most likely to share a communal haven, if
		only for protection in numbers. As well, Nosferatu share
		information with each other via networks — whether
		digital, personal, occult, or something less definable —
		that defy the comprehension of other Kindred`,
		IsHighClan: true,
	},
	"Ravnos": {
		Name: "Ravnos",
		Appearance: `Young Ravnos of-
		ten come from Eastern European Romani
		stock, with a relative paucity of “non-gypsy”
		gadje in the ranks. What few elders of
		the Clan may remain are presumed to
		come from Indian or Middle Eastern
		origins. Given that the Clan is widely
		spread and holds no traditional central
		domain, no consistent look can be said
		to be predominant, and any mendicant
		Kindred might be of Ravnos origin.`,
		Description: `The Ravnos move like the rumors that surround
		them. They are the thief in the night, the raksha
		chased by the wind, the nightmare-dream too fearful
		to be real. Whether associated with the Romani folk
		of Europe or the grave-robbing ghûl of Western Asia,
		Kindred society burdens the Ravnos with prejudices of
		foulness, uncleanliness, and wickedness. With reputations like these, the Ravnos are considered outsiders even among those Kindred who
		do not ally themselves with Sects. Many young
		Ravnos tend toward nomadic unlives, mov-
		ing from one domain to the next or hiding
		on the fringes of established territories
		where they can escape if local Kindred
		sentiment turns against them. The Ravnos practice a unique Discipline known as Chimerstry that convinces their enemies that they see
		things that do not exist. Chimerstry does much to con-
		vince Kindred that the Ravnos trade in lies and misdi-
		rection, but it can also prove to be a Deceiver’s salva-
		tion and ease the vagaries of an outcast unlife`,
		AssociatedSect: []string{
			"Independent",
		},
		Haven: `Many Ravnos take to the
		road instead of establishing perma-
		nent havens, dwelling temporar-
		ily among itinerant communities,
		at roadside rest stops, or even in ve-
		hicles. When a Deceiver does put
		down roots in a domain, his per-
		manent haven is often away from
		high-profile Kindred territories.
		Havens in ethnic ghettos, indus-
		trial outskirts, and isolated geog-
		raphy are safest and most easily
		cultivated for the Ravnos.`,
		Background: `The Ravnos are
		 scattered and suspicious, and those
		 childer who don’t have the tendency toward self-sufficiency don’t last long. In many
		 cases, a Ravnos will either never sire, or sire for com-
		 panionship or safety, with little concern for how well
		 a childe will fare as a vampire. Ravnos rarely seek out
		 childer actively, instead drawing from those whose
		 paths they cross on any given night. As such, the hard-
		 luck drifter reputation tends to follow the Deceivers.`,
		Character: `Outsider and selfish Natures
		are common among the Clan. Physical and Social At-
		tributes predominate, as do Talents and Skills. Savvy
		Ravnos develop a breadth of Backgrounds that can
		give them an edge when they’re on the move or in a
		tight spot, such as stashed Resources, defensible Do-
		main, and a few Allies or Contacts.`,
		Discipline: []Discipline{
			DisciplineMap[`animalism`], DisciplineMap[`chimerstry`], DisciplineMap[`fortitude`],
		},
		Weakness: `A turbulent history makes the Ravnos
		slaves to their vices. Each Ravnos has a penchant for
		some sort of vice — lying, cruelty, or theft, for example.
		When presented with the opportunity to engage in that
		vice, the Ravnos must indulge it unless her player suc-
		ceeds on a Self-Control or Instinct roll.`,
		Organizations: `The Ravnos are a far-flung Clan, with
		little to unite them and an open acknowledgement
		that each Deceiver looks after his own interests first.
		That said, Ravnos often like to make a great show of
		Clan camaraderie and cultural ritual, even though they
		know that promises made to one another are as fleet-
		ing as whispers on a night wind. Deceivers have been
		known to ally against common enemies such as tyrant
		Princes or Sabbat pogroms, but these alliances quickly
		fade once the threat no longer exists.`,
		IsHighClan: true,
	},
	"Toreador": {
		Name: "Toreador",
		Appearance: `Almost to the last, they
		are attractive in some way, whether
		the traditional beauty of a run-
		way model or the dangerous al-
		lure of something more predatory.
		The Degenerates augment their
		physical beauty with a sense of
		personal style, which may take
		the form of expensive cou-
		ture, avant-garde street wear,
		or classical fashions designed
		to emphasize their appeal-
		ing qualities. This isn’t to say
		that ugly Toreador don’t exist.
		Indeed, those gifted with less
		physical beauty often go that
		much further with their choice
		of accoutrements.`,
		Description: `From the Toreador perspective, when the sun fades,
		darkness gives rise to an eternal and wondrous world.
		Everything is fraught with wonder and terror, low poli-
		tics and sensual glories, the profound and the profane,
		and an undeniable undercurrent of the sanguinary.
		These Kindred are the Toreador, and they spend
		unlives ensconced in pleasure. Toreador often
		succumb to ennui, or fight the eventual
		boredom of unchanging immortality
		by playing at rivalries. An excess of
		stimulation turns them into
		slaves to the sensations they
		seek. The most debased
		Toreador can become true
		monsters, sinking to
		unimaginable levels of
		depravity in order to feel
		anything at all. Toreador culture is a mixture of syba-
		rites, dilettantes, and visionaries. Some
		Toreador, with echoes of mortal passion,
		Embrace lovers or “project” progeny who
		seem to fly in the face of every Toreador
		custom. These either don’t last long or
		rise to great prominence as subversives
		and individualists. Ideas, trends, and
		“the next great thing” spread through the Clan, and
		other Kindred often look to the Toreador to guide
		them. The Degenerates know this, and many become
		Harpies, Princes, and other key figures in vampiric so-
		ciety.
		`,
		AssociatedSect: []string{
			"Camarilla",	
		},
		Haven: `The Degenerates
		spare no expense in appoint-
		ing their havens in luxury,
		often with many original
		works of art. It is a point of
		pride among Toreador to
		have an unconventional
		(and thus memorable)
		haven with modern
		comforts; thus, many
		have striking lofts and
		penthouses, while the
		bolder among them renovate or repurpose everything from abandoned
		aquariums or deconsecrated churches to rooftop gar-
		dens or converted warehouse-galleries in fashion-for-
		ward neighborhoods. Share a communal haven? How
		déclassé.`,
		Background: `Many Toreador hail from high-society
		 or “bohemian” backgrounds. Indeed, many are them-
		 selves artists or influential among local art scenes or
		 other subcultures. Actors, singers, musicians, sculptors,
		 poets, playwrights, authors, and creative folk of any
		 stripe may well find a home in the Clan, as do those
		 who serve as patrons to (or travel in the entourages of)
		 those artistic types.
		 `,
		Character: ` Social attributes are almost al-
		ways primary, with an even split among Talents, Skills,
		and Knowledges depending on how the Toreador
		distinguishes herself. Toreador also love to cultivate
		Backgrounds. Allies, Contacts, Resources, Domain,
		Haven, Mentors, Resources, Retainers — all of these
		have great value among Toreador. Wise Toreador may
		choose to develop their Virtues, Humanity, Path, or
		Willpower, because with an unlife of degeneracy, they
		must frequently attend to the ugly business of bringing
		the Beast to heel.`,
		Discipline: []Discipline{
			DisciplineMap[`celerity`], DisciplineMap[`auspex`], DisciplineMap[`presence`],
		},
		Weakness: `When a Toreador experiences some-
		thing truly remarkable — a person, an objet d’art, a
		lovely sunrise — the player must make a Self-Control
		or Instinct roll. Failure means that the
		Kindred finds herself enthralled by the experience. The
		dazzled Toreador cannot act for the duration of the
		scene aside from commenting on or continuing their
		involvement with whatever has captured their atten-
		tion. If the experience no longer affects her (whether
		by moving, being destroyed, or whatever is appropri-
		ate to the situation), the captivation ends. Enraptured
		Toreador may not even defend themselves if attacked,
		though being wounded allows them to make another
		Self-Control or Instinct roll.`,
		Organizations: `Clan Toreador is cliquish and paro-
		chial in its local domains, but very rarely on a level that
		affects Clan-wide custom. Certain Toreador (and a few
		outside the Clan) sometimes use the terms “artiste” and “poseur” when describing individual Toreador,
		often derisively, to describe whether the Degenerate
		in question is one who is seen as creative or simply a
		follower of established trends, but these are certainly
		informal distinctions.`,
		IsHighClan: true,
	},
	"Tremere": {
		Name: "Tremere",
		Appearance: `Tremere often
		have two distinct presentations: a traditional and severe public aspect and a
		much more eldritch mien
		better suited to wielding their
		blood sorceries. When out
		in public or at Kindred
		events, the Tremere favor conservative suits and
		dresses and muted tones.
		When in their chantries
		or convening with others of their Clan,
		they often prefer robes
		decorated with subtle oc-
		cult symbols or garb
		with various folds and
		pockets for their bizarre
		ritual ingredients.`,
		Description: `In nights long lost to the passage of time, the Tremere
		existed, though they were something else. Those early
		Tremere then made a bargain — or wrought a spell,
		or any number of other harrowing methods attributed
		to the Clan — that changed them from what they had
		been into the vampires they are tonight. Some claim
		they stole the Curse of Caine from a torpid Antediluvian, or that they concocted the flawed immortality of the Kindred from the stolen vitae of other vampires. Such mysterious origins, which some describe as
		treacherous or even blasphemous, haunt the Tremere,
		as the other Clans look upon them with mistrust and
		suspicion. Tonight, Clan Tremere is a Clan shaped
		by its practice of blood sorcery. A flex-
		ible Discipline, Thaumaturgy is heav-
		ily entrenched within the Tremere,
		and they maintain cultic havens
		known as chantries to study
		its uses and share secrets
		among each other. To
		the Tremere, blood is
		both sustenance and the
		source of mystical power;
		they gather in their witch-
		houses to further their understanding of
		the vitae that is such a focal point of their
		unlives.`,
		AssociatedSect: []string{
			"Camarilla",
			},
		Haven: `Many Tremere rely on a central chantry the
		Clan maintains in cities where it has a notable pres-
		ence. More solitary Warlocks develop private havens,
		with all of the trappings one might expect from an oc-
		cult scholar, from libraries to alchemical laboratories to
		moonlit balcony observatories and even more sinister
		oubliettes where vivisected “research subjects” bleed
		according to experimental Tremere-controlled stimuli.`,
		Background: `The Tremere draw from a fairly narrow
		 pool of potential acolytes. Those who have an aware-
		 ness of the supernatural, who are driven to succeed,
		 who seek answers that elude less inquisitive individu-
		 als, yet who also have the discipline to heed the edicts
		 of the hierarchy make good Tremere. This isn’t to say
		 that individualists don’t have room in the Clan; rather,
		 those who go their own way may well find themselves
		 leading a chantry — or greeting the sun if their inter-
		 ests don’t align with those of the pyramid.`,
		Character: `Mental Attributes and Knowl-
		edges are prominent among the Tremere. Many have
		high Courage and Willpower, but are somewhat lacking
		in Conscience or Conviction. They often favor Back-
		grounds that heighten their relationship to the Clan, like
		Mentor, Status, and Retainers (whom an accomplished
		Warlock may craft from otherwise inert components).`,
		Discipline: []Discipline{
			DisciplineMap[`auspex`], DisciplineMap[`dominate`], DisciplineMap[`thaumaturgy`],
		},
		Weakness: `Tremere dependency on blood is even
		more pronounced than that of other Kindred. It takes
		only two draughts of another vampire’s blood for a
		Tremere to become blood bound instead of the normal
		three — the first drink counts as if the Tremere had taken
		two drinks.The elders of the Clan are well aware of this, and
		seek to impart loyalty to the Clan by forcing all neonate
		Warlocks to drink of the (transubstantiated) blood of the
		seven Tremere elders soon after their Embrace.`,
		Organizations: `The hallmark of Clan Tremere is “the
		pyramid,” the rigid hierarchy that governs the Clan
		and makes it the most organized of all the Kindred
		lineages. With many levels of membership, internal
		factions, and circles of mystery, the Tremere hierarchy
		presents a unified face to those outside the Clan, and
		is almost as unified behind the scenes. Still, the pyramid inculcates more than its share of paranoia, as both rogue Warlocks and a competitive environment of academic occultism pits each acolyte against her peers to
		the greater accomplishment of the Clan.`,
		IsHighClan: true,
	},
	"Tzimisce": {
		Name: "Tzimisce",
		Description: `A blood moon casts a crimson light over the land
		beyond the forest and something fearsome howls its agony into the night. The Tzimisce call these lands their
		ancestral home. Since time out of mind the Fiends
		have been masters and lords of the domains of much of
		Eastern Europe. But theirs is a proud, selfish Clan for
		which tradition goes only so far despite their aristocrat-
		ic origins. In fact, the Clan claims to have destroyed
		its Antediluvian, and in the wake of that momentous
		event, helped establish the foundations of the Sabbat.Tzimisce practice a strange Discipline known as Vicissitude that allows them to twist the skin and
		bone of their victims. Clan Tzimisce is a Clan of extremes, and
		long, cold nights spent in remote castles have
		turned the Fiends’ perspectives both greatly
		inward and outward. Mystics of the Clan study
		a philosophy of metamorphosis, seeking to discover what lies beyond the state of vampirism.`,
		Appearance: `Given their ability to manipulate their
		physical appearance with Vicissitude, Tzimisce look however they want, and they often want to provoke or
		frighten. Some prefer extreme modifications and experimentations with their bodies that leave them looking
		only vaguely humanoid. Others seek to redefine and even
		transcend the limits of their forms, rebuilding themselves in the images of
		angels, monsters,
		nightmares,
		and things
		even less
		recognizable`,
		AssociatedSect: []string{
			"Sabbat",
		},
		Haven: `Young
		Tzimisce
		are often
		Sabbat
		Priests or
		Ducti and prefer to maintain communal havens with their
		packs. They encourage the pack to live in fearsome loca-
		tions, such as beneath a hospital or morgue, or in the dank
		recesses of a mausoleum. Elders of the Clan sometimes
		have ancestral holdings in the Old World, and the image
		of the vampire on the craggy mountain in a crumbling
		castle owes much to Tzimisce lords. Rarely are these an-
		cient holdings kept to any modern standards of comfort,
		but their lords are strangely hospitable to invited guests
		(and terribly intolerant of uninvited bores).`,
		Background: `Elder Tzimisce, particularly those of
		 the still-landed nobility in hoary old domains, may
		 have family lines from which they Embrace, or they
		 may restrict their occasional siring to the terrified vil-
		 lages suffering in thrall beneath their estates. New
		 World and younger Tzimisce aren’t as discriminating
		 and are more practical. Indeed, many Tzimisce fledg-
		 lings are little more than shock troops, Embraced and
		 warped to the limits of their frames to cause revulsion
		 and revel in bloodshed until put down.`,
		Character: `Few Tzimisce know modera-
		tion, and thus often favor Physical or Mental Attri-
		butes, usually with one extraordinarily high Trait.
		Knowledges are usually favored, though Skills are val-
		ued as well. Mentor, Allies, Domain, and Retainers are
		all quite appropriate to Tzimisce, as are narrow — al-
		most overdeveloped — Disciplines. Their alien mind-
		sets often lead them to follow Paths of Enlightenment.
		The ghastliness to which many Tzimisce are often ex-
		posed during their fledgling years sometimes results in
		complicated derangements.`,
		Discipline: []Discipline{
			DisciplineMap[`animalism`], DisciplineMap[`auspex`], DisciplineMap[`vicissitude`],
		},
		Weakness: `The Tzimisce are inextricably tied to
		their domains of origin, and must rest in the proximity
		of at least two handfuls of “native soil” — earth from a
		place important to her as a mortal, such as the soil from
		her birthplace or the graveyard where she underwent
		her Embrace. Each night spent without this physical
		connection to her land limits all of the Tzimisce’s dice
		pools to one-half, cumulatively, until she has only a
		single die in her pool. The penalty remains until she
		rests for a full day amid her earth once more.`,
		Organizations: `The Fiends are, on the whole, mistrust=ful of other Kindred, especially others of their own Clan.
		As such, Tzimisce organization, such as it is, has a high
		regard for solitude. It is against their nature to be inclusive, and thus they have to work at being Sabbat (though
		this is less difficult for younger Tzimisce of at least somewhat modern outlooks). This is also a key reason so many
		of them undertake Paths of Enlightenment: to give their
		xenophobia purpose, but also to provide some common
		point of reference with others on the Path.`,
		IsHighClan: true,
	},
	"Ventrue": {
		Name: "Ventrue",
		Appearance: `The Ventrue favor conservative
		clothing and reserved presentation, unless
		they’re making a point about power
		or money. Ventrue Princes may
		well wear a circlet or carry a
		scepter as symbols of office,
		while young Blue Bloods
		display their own achieve-
		ment via suits, ties, dresses,
		and accessories that are eas-
		ily overlooked singularly
		but add up to a stunning to-
		tal effect. If a Ventrue has so
		much as a hair out of place,
		it’s because he spent all night
		running down the Society of
		Leopold and demanding the
		Sabbat menace retreat.`,
		Description: `Throughout history, while the other Clans have
		skulked about their petty intrigues, the Ventrue have
		curried favor with Caesar, whispered into the ear of
		Charlemagne, bankrolled the Age of Exploration, and
		even swayed policy in the Holy See. Theirs
		is a legacy of rulership, from Ventrue fledglings starting their
		climb to the top to the mightiest elders whose influence
		spans the world. Long have
		they played kingmaker in the
		shadows in the mortal world, and
		long have they been the Clan of
		Kings among the Kindred. Tonight, the Ventrue are a
		synthesis of the modern and
		the ancient, often in stark
		contrast within the Clan and
		among one another. Theirs is
		money of old, from the vaults
		of Croesus, but their young manipulate stock markets and influence currencies. Elders may command armies or even whole
		governments, while neonates conjure their assets from
		a website or smartphone app. But for all their wealth,
		their distinguished history, and their status among the
		Damned, each and every Ventrue must still seek that
		one resource that makes Kindred society egalitarian:
		precious blood.`,
		AssociatedSect: []string{
			"Camarilla",	},
		Haven: `A Ventrue’s haven
		displays both her great power
		(read: wealth) and distinguished
		tastes. Opulent, grandiose, even
		baroque — these may all apply to
		Ventrue havens. They shun the
		gaudy displays of other Kindred,
		and their style tends less to the
		avant-garde than it does to the
		classical and traditional. To the
		Blue Bloods, a well-maintained haven is an extension of oneself, and for someone to
		see it in less than flawless state implies weakness, distraction, or even madness.`,
		Background: `Anyone who has “made something of
		 herself” may attract the attention of the Ventrue, who
		 judge their childer based on their prominence and suc-
		 cess even before they start to groom them for the Em-
		 brace. Socialites, moneyed family, corporate wunderkinds,
		 military leaders, and even untested newcomers who show
		 great promise are keenly valued among the Blue Bloods`,
		Character: `Ventrue usually have directorial
		or outgoing personality archetypes. Members of Clan
		Ventrue favor Social and Mental Attributes, but any
		Ability category can be primary, reflecting a personal avenue of expertise. Backgrounds go both wide and deep
		for the Ventrue, with almost every Blue Blood possessing
		some amount of Resources, Status, and Herd (particularly the latter, given the Clan weakness). Elders in particular cultivate enviable havens and sprawling Domains.`,
		Discipline: []Discipline{
			DisciplineMap[`dominate`], DisciplineMap[`fortitude`], DisciplineMap[`presence`],
		},
		Weakness: `The Ventrue have rarified tastes, and
		they find only one specific type of mortal blood palat-
		able and vital for them. When a player creates a Ven-
		true character, they should decide with the Storyteller
		what specific type of blood suits the character, and this
		choice is permanent. Blood of other types (even ani-
		mals) simply offers the vampire no blood pool increase,
		no matter how much he consumes — he simply vomits
		it back up. This refined palate may be very narrow or
		very broad — say, the blood of younger sisters, or the
		blood of nude children. Vampiric blood is exempt from
		this restriction`,
		Organizations: `The Clan-wide organization of the
		Ventrue is localized and feudal, with various universal-
		ly understood peerages, vassalages, oaths of fealty, and
		sworn boons taking the place of a rigid hierarchy. Many
		Ventrue style themselves as secret masters of their do-
		mains, consolidating power in longstanding networks
		and forming conspiracies. The Ventrue greatly value
		propriety and honor, and use many forms of address
		and respect — their Laws of Decorum are complex
		and rigid, and could fill several volumes. Almost every
		Ventrue worth his status can recite his lineage at least
		back to the elders, if not to the great Methuselahs.`,
		IsHighClan: true,
	},
}

var SectMap = map[string]Sect{
	"Camarilla": {
		Name: "Camarilla",
		Description: `The Camarilla is also known as “The Ivory Tower,”
		and the Sect lives up to that moniker. Created in the
		15th century, the Camarilla was formed to preserve
		and protect vampire society against the decimation
		brought on by the Inquisition as well as the power drain
		presented by the War of Princes during the Dark Ages.
		The leaders of the Camarilla ruthlessly enforced the
		Tradition of the Masquerade, elevating it to the Sect’s
		highest law, a priority they continue to pursue even
		in modern nights. It is the largest Sect of Kindred, after all, and nearly any city on the globe likely has some Camarilla presence. This expansiveness is partially due to the fact that the Camarilla says that any vampire, regardless of Clan or bloodline, may go to a Prince and claim membership in the Camarilla. While Kindred of any lineage can claim membership, most come from the founding Clans of the Camarilla: the Brujah, Gangrel, Malkavians, Nosferatu,nToreador, Tremere, and Ventrue. These Clans helped
		to create the Camarilla, and they have a seat on the
		Inner Circle. Vampires of other bloodlines
		can attend conclave and meetings, but they are treated
		as minority voices or simply ignored. As a result of this ongoing conflict with the Sabbat, in modern nights the Ivory Tower is crumbling, losing a few bricks here and there while it proudly proclaims itself to be strong and whole.  Neonates feel more and more like serfs to the elder nobility, asked to protect and uphold an organization in which they have little hope of advancement but
		plenty of opportunities for punishment. The ancillae
		end up in the worst position: unable to break the glass
		ceiling of elder dominance, but given enough scraps
		of power to make the younger Kindred sharpen their
		knives in jealousy. The neonates and younger ancillae have an increasing advantage — modern technology.`,
		Practices: `At the heart of the Camarilla lie the Traditions, and
		at the heart of the Traditions lies the Prince. Whether
		in name or deed, the Prince enforces the Traditions
		within their domains and punishes those who violate
		their law.
		 Many Princes hold a regular court, which
		function as a combination of social affair and legal pro-
		ceeding. At court, vampires gossip, politick, and gain
		favor with the Prince, and the Prince and his officers
		handle matters of the law and render judgment or enact policies.When a vampire is found guilty of high crimes (such
		as an egregious breach of the Masquerade, the diablerie
		of another Kindred, or offending the Prince), the Prince
		can call for the Lextalionis. The currency of the Camarilla (and many other independent vampires that deal with them) is a system of favors and debts. Such debts (or “boons”) are not only given and incurred, but are also traded between vampires in a healthy, invisible economy. Conclaves are the center of global Camarilla politics, serving as the highest court of Camarilla legislation, a
		high-level session of Sect-wide policy, and a political
		rally all rolled into one. Only called by a Justicar, all
		Kindred that the Camarilla claims authority over can
		attend a conclave, and the proceedings can be over in
		a few hours or stretch on for many weeks. Given the
		open-door policy of conclaves, security is an issue, and
		the location of a particular conclave might not be an-
		nounced until a few days beforehand.Generally, conclaves are called because of the actions of a powerful individual (such as a Prince) or due to a serious breach of the Traditions, but once called, any Kindred may bring a grievance to the conclave and
		have it addressed.`,
		Titles: []Entry{
			{
				Name: "Prince",
				Description: `The preeminent vampire of a Camarilla city and probably the most numerous position of ultimate au-
				thority among Western Kindred, Princes are the rul-
				ers of given cities. Some Princes are tyrants or absolute
				monarchs of the Damned while others are politically
				feeble puppets propped up by more powerful support-
				ers, but the position of Prince is one acknowledged
				and even (grudgingly) respected by all vampires, even
				those not of the Camarilla.  Prince’s duties and privileges are many, but the
				most important is the interpretation and enforcement
				of the Traditions, particularly the Masquerade. Beyond
				that, a Prince has any individual powers he can claim
				and uphold, such as declaring Elysium, calling a Blood
				Hunt, adjudicating disputes between residents of his
				domain, the right to claim a blood-tribute, and poten-
				tially even the right to name, ignore, or even disband
				the Primogen.`,
			},
			{
				Name: "Harpy",
				Description: `Harpies are the opinion leaders and the
				trend-setters to whom other Kindred look when it
				comes to matters of taste, style, philosophy, or poli-
				tics. A Harpy’s word influences the domain’s attitudes
				and can be a powerful supporter of the status quo or a
				force for insidious change. Harpies are rarely appointed
				directly (and Kindred rarely trust those who are). In-
				stead, a Harpy paradoxically becomes so by acting as
				a Harpy. The Harpy’s role is often intertwined with
				domain politics, and it is a bold or foolish Prince who
				neglects those vampires who represent the cutting edge
				of popular opinion in her domain.`,
			},
			{
				Name: "Keeper Of Elysium",
				Description: `This is a largely honorific title,
				though it has many practical responsibilities. The
				Keeper of Elysium assures that the customs of Elysium
				are observed, and is a caretaker of sites declared Ely-
				sium by a Prince.`,
			},
			{
				Name: "Scourge",
				Description: `In this time, which many elders believe to
				be the Time of Thin Blood, the Scourge is responsible
				for destroying those vampires of 14th and higher Generation. This is a position most often found in paranoid
				domains where elders believe that fighting the symptoms of the Final Nights will stave off the coming of
				Gehenna.`,
			},
			{
				Name: "Sheriff",
				Description: `The Sheriff is the Prince’s right-hand Kindred, responsible for the physical enforcement of
				Princely decree. Some Sheriffs are diligent masters-at-
				arms while others are thuggish, bloody fuckheads who
				abuse their authority to torment those beneath their
				station. A Sheriff may appoint Hounds to assist him
				(or the Prince may appoint them, in the interests of
				curtailing a Sheriff’s overt power).
				`,
			},
			{
				Name: "Primogen",
				Description: `Primogen is a flexible title. In some domains, a Primogen is simply the eldest and most influential Kindred of a given Clan. In others, a Primogen
				is a member of a council of advisors to the Prince. In
				some domains, Princes appoint the Primogen (and
				they may even appoint multiple Primogen of the same
				Clan, to keep that Clan divided) while in other cities,
				Primogen arise from among the most powerful members of that Clan or as a result of popular Clan politics.`,
			},
			{
				Name: "Justicar",
				Description: `hese are the judges appointed by the Inner Circle to be the Camarilla’s eyes, hands and, if necessary, fists. It is the Justicars who decide the punishment for those who have violated the Traditions on a
				widespread level. A Justicar serves for 13 years, and her
				actions may be challenged only by another Justicar.`,
			},
			{
				Name: "Archon",
				Description: `Each Justicar selects a number of minions,
				known as Archons, to act in his name as suits his purposes. Archons are typically chosen from the upper
				ranks of ancillae and occasionally elders of lesser station. Justicars occasionally choose Archons to carry
				out specific missions, and prefer political savvy, insight, and skill over standing and clout. An Archon’s
				position typically lasts for as long as a Justicar wishes to retain her, but no longer than the Justicar’s tenure.`,
			},
			{
				Name: "Alastor",
				Description: `If the Archons and Justicars are the law
				enforcement agencies of the Camarilla, the Alastors
				are the secret police. Moving unseen and unnoticed
				through the Camarilla, they serve a variety of purposes
				at the Inner Circle’s command. Mostly, however, the
				Alastors hunt the most dangerous criminals to the Camarilla — those on the so-called “Red List.” While the
				anonymous existence of the Alastors can be a difficult
				one, it does have its rewards, specifically a more-or-less
				universal immunity to prosecution from local Princes`,
			},
			{
				Name: "Inner Circle",
				Description: `The true hub of the Camarilla,
				this group meets once every 13 years to plan out the
				business and direction of vampire society — as much as
				any group can presume to dictate the doings of a race of
				immortal predators. Every Clan is permitted one representative, usually the eldest member of the Clan, as
				only the eldest may cast the Clan’s vote.`,
			},
		},
		Strongholds: []string{
			"Chicago", "New Orleans", "Vienna", "London", "Paris",
		},
	},
	"Sabbat": {
		Name: "Sabbat",
		Description: `Opposite the Camarilla stand the Sect of inhuman
		vampires known as the Sabbat. The Sabbat follow the Code of Milan Most of the vampiric factions believe that the so-called “Sword of Caine” is a collection of mindless barbarians and ultraviolent
		fiends, or even demon worshippers bent on bringing
		Satan to earth. As such, the Sabbat are vilified all
		throughout vampire society.They reject humanity as a moral basis
		for their lives, so they have turned to other alternatives.
		They adhere to a wide variety of Paths of Enlightenment, philosophical tenets that force the Beast into a narrow channel and allow the Cainite to maintain a day-to-day existence that resembles stability (if not sanity). Further, the Sabbat not only rebel against morality, but also against their own inclina-
		tions towards solitude. They frequently join together
		into packs that act as one part religious cult, one part
		political faction, and one part combat unit. Between moral devotion, pack loyalty, and the need to rebel, Sabbat cities are devoid of the calm, quiet society of the Camarilla court. Tensions are always high in Sabbat “dioceses,” and the Cainite’s surroundings often reflect their explosive natures. In cities controlled by Sabbat, murder, robberies, rape, and assault are commonplace. cause of these tensions, the Sabbat is hardly a uni-
		fied entity, and the Sect is home to numerous factions
		of vampires united under the Sabbat’s banner. One of
		the most feared is the Black Hand, a special militia hid-
		den within the packs of the Sword of Caine. All mem-
		bers of the Black Hand bear a distinguishing mark — a
		permanent, mystical sigil on the palms of their right
		hands. Although this brand may be concealed or made
		over, it may never be removed — once you’re in the
		Hand, you’re a member until Final Death.
		The thread that holds the Sabbat together is a religious zeal to tear down the Antediluvians. When packs or ideologies threaten to explode into all-out civil war, the Sabbat manage to pull together around their common hatred for the Camarilla.`,
		Titles: []Entry{
			{
				Name: "Archbishop",
				Description: `With many of the same powers as the Prince, the Archbishop is the closest analog the Sabbat has to that
				position. An Archbishop is different, though, in that
				the Sabbat is less concerned with enforcing the Traditions and more concerned with waging its holy war
				against the Antediluvians and everyone else. Thus,
				the Archbishop is part spiritual leader and part warlord, advancing the Sword of Caine’s agenda and establishing Cainite primacy. This last is a difficult task
				to undertake, as it’s not simply a question of turning a
				Sabbat city into a living hell and declaring that vampires rule; the fundamentals of the Masquerade and the
				sheer weight of the mortal population means that is
				a war to be waged in stages. Too many Sabbat fail to
				understand this, especially among the young, and lose
				faith in their leadership because they’re too impatient
				to play out the long-game Jyhad. This subversion of ignorance is perhaps the Archbishop’s greatest challenge
				to overcome.`,
			},
			{
				Name:        "Ductus",
				Description: `Leaders of individual packs, Ducti attend to the operation matters of their charges, resembling gang leaders or chiefs of small tribes. The title of Ductus is largely honorific, according recognition to the most accomplished member of a pack. Some authority accompanies the title, but the Ductus who throws his weight around is likely to find his ass dumped unceremoniously in a trash bin, if not staked out to welcome the next sunrise.`,
			},
			{
				Name: "Pack Priest",
				Description: `Priests bear the responsibility for the
				spiritual wellbeing of their packs. Second in command
				to the Ductus, the Pack Priest officiates all the rituals
				observed by the pack, and often creates a few for the
				sole use of the pack. All packs have at least one Pack
				Priest, though some rare and large packs have two.
				In the event that the Ductus is eliminated, the Pack
				Priest becomes a temporary leader until a new leader
				can be appointed by the Bishop, Archbishop, or (in
				autonomous packs) the pack itself.`,
			},
			{
				Name: "Templar",
				Description: ` Also known as Paladins, the Templars are
				an elite force of bodyguards appointed by a Bishop or
				greater leader. Templars serve a variety of duties, always
				in a martial capacity. Most Archbishops keep a cadre
				of Paladins in their retinues to handle delicate matters
				best solved by a judicious application of violence.`,
			},
			{
				Name: "Bishop",
				Description: `Bishops are those immediately below the
				Archbishop in Sabbat domains, with the same high 
				level duties but much narrower scope. A Bishop may
				be in charge of a single aspect of Cainite unlife in the
				domain, or she may be a more general spiritual leader,
				inquisitor against diabolism, military general, or any
				other specific aspect of Sabbat agenda. Some Sabbat
				domains have no Archbishop, but are instead ruled by
				a council of Bishops.`,
			},
			{
				Name: "Priscus",
				Description: `Superficially, a Priscus is similar to a Primogen, but is never appointed by an Archbishop or
				Bishop; Prisci rise to the rank over time and according
				to no specific criteria. A Priscus, politically, is an advisor to the local Sabbat authority figures, with no formal role in Sect politics but great practical influence.`,
			},
			{
				Name: "Cardinal",
				Description: `Cardinals oversee Sabbat affairs in large
				geographical regions. As the superiors of the Archbishops, Cardinals coordinate the Sabbat in their cities and
				direct them in the Jyhad. Further, it is their direct duty
				to bring any cities within their territory under the Sabbat’s sway. Most Sabbat see their Cardinals no more
				than once per year, if at all, as the duties of the office
				keep them in constant communication with Bishops,
				Archbishops, Prisci, and the Regent herself.`,
			},
			{
				Name: "Regent",
				Description: `The Regent is the ultimate authority over
				all of the Sabbat, ruling the Sect from the Sabbat
				stronghold in Mexico City. Each of the Archbishops
				or councils of Bishops is ultimately accountable to the
				Regent.`,
			},
		},
		Strongholds: []string{
			"Detroit", "Miami", "Mexico City", "Montreal", "Madrid",
		},
		Practices: `Everything in the Sabbat points back to two key
		(and often conflicting) principles: freedom and loyalty.
		Their set of laws, the Code of Milan, reinforces these two points. Cainites are free to do what they wish, but they must always remain loyal to the Sword, because only the Sect can protect free vampires from manipulative elders, especially the Antediluvians. Two of the most influential Clans of the Sect, the Lasombra and the Tzimisce, are believed to have destroyed their Antediluvian progenitors and diablerized them, setting the example for the other Sabbat vampire Clans to follow.
		The core of the Sabbat is the pack, a collection of vampires bound by Vinculum and common interest. Packs are such an integral part of Sabbat society that Cainites without a pack are often viewed with suspicion and distrust. Each pack usually fights within itself until some authority rises to the top, such as a priest to conduct rituals and a Ductus to lead in secular matters. Higher positions exist, but unlike
		the Camarilla where all power stems from the Prince,
		in the Sabbat positions tend to build upon the packs`,
		Rituals: `Sabbat vampires go through Creation Rites designed
		to prove themselves True Sabbat. The rite forces the
		vampire to prove himself to his new sire or pack, often
		in a way that strips away much of his Humanity, turning the fledgling from a weak, pathetic human into a
		proper monster worth of the Sect. Out of mockery of human religious belief, many of the rituals (or “ritae”) and conventions of the Sabbat are perversions of those of the Catholic Church.`,
	},
	"AnarchMovement": {
		Name: "Anarch Movement",
		Description: ` Nominally a faction within the Camarilla, most “An-
		archs” are still under the authority of the Ivory Tower.
		The Camarilla would say that the Anarchs are under
		the protection of the Sect, while the Anarchs would
		likely call it oppression. Still, many in the Anarch
		Movement understand the usefulness of the structure,
		with only the most radical calling for total withdrawal
		from the Camarilla.The Anarchs seek to change the
		Ivory Tower from within, turning it into the benevo-
		lent Kindred society it so often claims to be.
		Despite their name, not all Anarchs
		are actually anarchists; many seek to change the Ca-
		marilla or the Sabbat into some sort of new structure,
		usually based on mortal governmental ideas. Most of
		these structures tend to revolve around some form of
		democracy, but variations of neo-feudalism and even
		fascism have been tried (with varying success) through-
		out the Movement. `,
		Practices: `Status is not only an incompatible concept, but
		actively offensive to some Anarchs. Instead, Anarchs
		rely on whether they have personally heard of another
		Anarch, instead of what some other vampire said was
		important. The Status Background for Anarchs relies on reputation: a combination of self-promotion, word-of-mouth, and outright braggadocio that gets the Anarch recognized both inside and out of his
		own barony. It doesn’t correspond to a political title or
		a measure of authority — there are untitled vampires
		(such as the notorious prankster Smiling Jack) with a
		universally higher reputation than Kindred who claim
		the title of Baron. Whether this recognition confers respect or disdain is a secondary concern; as Oscar Wilde
		said, “There is only one thing in life worse than being
		talked about, and that is not being talked about.”
		
		The Anarchs genuinely like to have fun and joke
		with each other. They believe that Kindred existence
		shouldn’t be all backstabbing and bloodletting, but
		rather about enjoying each night to the fullest. Sabbat
		fiends consider this idea to be a disgusting show of humanity while Camarilla cynics believe it to be the folly
		of youth, but the Anarchs don’t care. They frequently
		play games and pull pranks on each other (and on other unsuspecting Kindred), both as tests of courage and
		as a way to blow off some steam`,
		Titles: []Entry{
			{
				Name: "Baron",
				Description: `Much of Anarch philosophy decries the tyranny of Princes, but despite their name, the Anarchs often recognize that leadership in a domain is necessary. Thus, the Baron: a Prince in all but name and reputation.
				Many Barons take great pains to avoid Camarilla-style
				autocracy, but power can corrupt even the most zealous
				freedom fighter. A Baron forever straddles a fine line
				between being a wise leader and a power-mad autocrat,
				and the impractical nature of revolutionary ideology
				dooms many of them when the nightly affairs of the
				domain need attention.`,
			},
		},
		Strongholds: []string{
			"Los Angeles", "San Diego", "San Francisco",
		},
	},
	"Independents": {
		Name: "Independent",
		Description: `Over the past few centuries, the fatal dance between
		the Camarilla and the Sabbat has changed vampire so-
		ciety. Their bloody Jyhad has shaped the secret history
		of the world, destroying millions of mortal lives in the
		process. However, there are some Clans who witness
		both Sects tear at each other’s throats in the name of
		their ancient feud and decide that they would prefer to
		watch from the sidelines. Of course, independent Kindred are vampires first and Clan members second; they don’t all possess mindless,
		unwavering devotion to the ideals of their Clan. Most
		of them focus on their personal goals first, whether or
		not they coincide with their Clan’s goals. This has the
		side-benefit of confusing outside Kindred even more, since it’s never clear whether a particular independent will act with their Clan’s interests in mind, their personal interests, or some hybrid of the two. The Independent clans are: The Assamites, The Followers of Set, The Giovanni and the Ravnos`,
		Strongholds: []string{
			"Alamut(Assamites)", "Cairo(Followers of Set)", "Venice(Giovanni)", "Kolkata(Ravnos)",
		},
	},
}

var TraditionMap = []Traditions{
	{
		Name: "The Traditions",
		Description: `Vampires observe a set of customs that exists some-
		where between being coded into their undead natures
		and a social contract that’s ratified every night among
		the courts of the Damned. `,
		Tradition: []Entry{
			{
				Name: "The Masquerade",
				Description: `Thou shall not reveal thy true nature to those not of the Blood.
				Doing such shall renounce thy claims of Blood`,
			},
			{
				Name: "The Domain",
				Description: `Thy domain is thine own concern.
				All others owe thee respect while in it.
				None may challenge thy word while in thy domain.`,
			},
			{
				Name: "The Progeny",
				Description: `Thou shall only Sire another with the permission of thine Elder.
				If thou createst another without thine Elder’s leave, both thou and thy Progeny shall be slain.`,
			},
			{
				Name: "The Accounting",
				Description: `Those thou create are thine own children.
				Until thy Progeny shall be Released, thou shall command them in all things.
				Their sins are thine to endure.`,
			},
			{
				Name: "Hospitality",
				Description: `Honor one another’s domain.
				When thou comest to a foreign city, thou shall present thyself to the one who ruleth there.
				Without the word of acceptance, thou art nothing.`,
			},
			{
				Name: "Destruction",
				Description: `Thou art forbidden to destroy another of thy kind.
				The right of destruction belongeth only to thine Elder.
				Only the Eldest among thee shall call the Blood Hunt.`,
			},
		},
	},
	{
		Name: "The Code of Milan",
		Description: `Vampires observe a set of customs that exists some-
		where between being coded into their undead natures
		and a social contract that’s ratified every night among
		the courts of the Damned. `,
		Tradition: []Entry{
			{
				Name: "Code 1",
				Description: `The Sabbat shall remain united in its support of the Sect’s Regent.
				If necessary, a new Regent shall be elected. The Regent shall support relief from
				tyranny, granting all Sabbat freedom.`,
			},
			{
				Name:        "Code 2",
				Description: `All Sabbat shall do their best to serve their leaders as long as said leaders serve the will of the Regent.`,
			},
			{
				Name:        "Code 3",
				Description: `All Sabbat shall faithfully observe all the auctoritas ritae.`,
			},
			{
				Name:        "Code 4",
				Description: `All Sabbat shall keep their word of honor to one another.`,
			},
			{
				Name:        "Code 5",
				Description: `All Sabbat shall treat their peers fairly and equally, upholding the strength and unity of the Sabbat. If necessary, they shall provide for the needs of their brethren.`,
			},
			{
				Name:        "Code 6",
				Description: `All Sabbat must put the good of the Sect and the race of Cainites before their own personal needs, despite all costs.`,
			},
			{
				Name:        "Code 7",
				Description: `Those who are not honorable under this code will be considered less than equal and therefore unworthy of assistance.`,
			},
			{
				Name:        "Code 8",
				Description: `As it has always been, so shall it always be. The Lextalionis shall be the model for undying justice by which all Sabbat shall abide.`,
			},
			{
				Name:        "Code 9",
				Description: `All Sabbat shall protect one another from the enemies of the Sect. Personal enemies shall remain a personal responsibility, unless they undermine Sect security.`,
			},
			{
				Name:        "Code 10",
				Description: `All Sect members shall protect Sabbat territory from all other powers.`,
			},
			{
				Name:        "Code 11",
				Description: `The spirit of freedom shall be the fundamental principle of the Sect. All Sabbat shall expect and demand freedom from their leaders.`,
			},
			{
				Name: "Code 12",
				Description: `The Ritus of Monomacy shall be used to settle disputes among all Sabbat.
				`,
			},
			{
				Name: "Code 13",
				Description: `All Sabbat shall support the Black Hand.
				`,
			},
			{
				Name: "Addendum 14",
				Description: `All Sabbat have the right to monitor the behavior and activities of their fellow Sect members in order to maintain freedom and security.
				`,
			},
			{
				Name: "Addendum 15",
				Description: `All Sabbat possess the right to call a council of their peers and immediate leaders.
				`,
			},
			{
				Name: "Addendum 16",
				Description: `All Sabbat shall act against Sect members who use the powers and authority the Sabbat has given them for personal gain at the expense of the Sabbat. Action shall be taken only through accepted means, approved by a quorum of Prisci.
				`,
			},
		},
	},
	{
		Name:        "Auctoritas Ritae",
		Description: ` The overarching term that describes all mandatory rituals within the Sabbat culture. Includes The Blood Bath, The Blood Feast, Creation Rites, Fire Dance, Games of Instincts, Monomacy, the Sermons of Caine, The War Party, and the Wild Hunt`,
		Tradition: []Entry{
			{
				Name: "The Blood Bath",
				Description: `This ritus is performed whenever Sect leaders wish to recognize a Sabbat vampire’s claim to a title, such as Bishop or Cardinal. The Blood Bath formalizes the vampire’s new status in the Sect. As many Sabbat as possible who will serve under this new leader must attend the ceremony, for failing to do so without an adequate reason is a grave slight to the leader in question.Starting with the priest conducting the ritus, attendant
				Sect leaders and other Sabbat take turns coming for-
				ward, kneeling in front of and expressing their en-
				dorsement of or allegiance to the Cainite, and contrib-
				uting a quantity of blood into a large vessel. The newly
				titled vampire gives praise and/or advice to each of the
				vampires present, emphasizing the benefits the Sabbat
				stands to gain through the sharing of her wisdom. She
				then bathes in the blood donated to the pool. Follow-
				ing the ceremony, all vampires present drink from the
				bathing vessel (the blood in which is sometimes conse-
				crated as a Vaulderie), symbolizing that they willingly
				partake of everything the new leader has to offer.`,
			},
			{
				Name: "The Blood Feast",
				Description: `The Blood Feast is a ritual “meal,”
				in which captured vessels are suspended from the ceil-
				ing, bound to sculptures, or otherwise immobilized and
				fed from at the leisure of all vampires present. The feast itself is as much social gathering as it is a structured ritus, and many Sabbat make grand entrances, wearing the best of their finery.`,
			},
			{
				Name: "The Blood Feast",
				Description: `The Blood Feast is a ritual “meal,”
				in which captured vessels are suspended from the ceil-
				ing, bound to sculptures, or otherwise immobilized and
				fed from at the leisure of all vampires present. The feast itself is as much social gathering as it is a structured ritus, and many Sabbat make grand entrances, wearing the best of their finery.`,
			},
			{
				Name: "Creation Rites",
				Description: `After the Embrace, the new vampire is eligible for the Creation Rites only after he has demonstrated his worth to the Sect. The ritus itself
				is quite simple — the priest merely touches a flaming
				brand to the initiate’s head and leads him in an oath of
				allegiance. The ceremony that precedes the Creation
				Rites, however, varies widely, and it is wholly in the
				hands of the Cainite’s sire.`,
			},
			{
				Name:        "Fire Dance",
				Description: `To most vampires, fire is something to be feared and avoided, yet not to the Sabbat. While they still fear it, they are not above turning it loose on their enemies. To be fully Sabbat, one must face the Rötschreck and master it.`,
			},
			{
				Name:        "Games of Instinct",
				Description: `The vampires of the Sabbat engage in numerous sanctioned “games,” adjudicated by their Pack Priests to maintain their predatory edge. These games take various forms, and different packs practice different styles — everything from parodies of children’s games or sports to completely unique vampiric tests of skill can be made into a Game of Instinct. `,
			},
			{
				Name:        "Monomacy",
				Description: `Monomacy is usually practiced by only ranking members of packs, a challenge of ritual combat is issued to the high ranking member simultaneously with the challenge to the rival. The priest then decides whether or not the grudge is worth Monomacy, and whether or not she chooses to preside over the ritual. The challenged vampire may decline. In theory, there is nothing wrong with declining a challenge, but unless the challenger is of such little consequence as to be below the challengee’s notice, declining usually involves a great loss of face.`,
			},
			{
				Name:        "Sermons of Caine",
				Description: `Pack members take turns reciting from the Book of Nod, while the others sit in a semicircle holding lit candles and meditating on the passages. The sermons are sometimes followed by the Vaulderie, and, among more intellectual packs, intense deliberation. Pack members often discuss the passages read during the ritus almost until dawn`,
			},
			{
				Name:        "The War Party",
				Description: `War parties consist of multiple packs that vie for the blood of a non-Sabbat elder. Packs participating in the War Party compete against one another for the privi lege of killing and diablerizing the elder, but rarely do the packs come into deadly conflict with each other, reserving their violence for their target. In preparation for a War Party, the participating packs gather and celebrate. They may also perform the Fire Dance, listen to Sermons of Caine, and participate in a Blood Feast or Vaulderie. The chief of the War Party, usually the most accomplished or highest-ranking priest among the packs, offers the assembled packs the challenge.`,
			},
			{
				Name:        "The Wild Hunt",
				Description: `One of the greatest crimes a Sabbat can commit is to turn traitor, and the Sect protects its secrets. If a member reveals a Sect secret to the enemy, she is punished severely. If a Sabbat leaks information of a vital nature, a priest may call for a Wild Hunt. The Wild Hunt is much like the blood hunt, but ends with the eradication of the offending Sabbat Sect member, as well as anyone — Kindred or kine — who may have knowledge of the betrayal. The gravity of the Wild Hunt depends upon the traitor in question — the packs are expected to police their own ranks, while high-profile turncoats receive the attention of Archbishops, Prisci, Cardinals, and all those who serve them.`,
			},
		},
	},
	{
		Name:        "Ignoblis Ritae",
		Description: ` The overarching term that describes all low or common rituals within the Sabbat culture. These rituals vary widely from pack to pack. Includes the Acceptance Ritus, Contrition Ritus, Stealth Ritus, Sun Dance`,
		Tradition: []Entry{
			{
				Name: "Acceptance Ritus",
				Description: `The Sabbat shall remain united in its support of the Sect’s Regent.
				If necessary, a new Regent shall be elected. The Regent shall support relief from
				tyranny, granting all Sabbat freedom.`,
			},
		},
	},
}

var DisciplineMap = map[string]Discipline{
	"animalism": {
		Name: `Animalism`,
		Description: `The Beast resides within all creatures, from scut-
		tling cockroaches to scabrous rats up through untamed
		wolves and even powerful Kindred elders. Animalism
		allows the vampire to amplify his intensely primordial
		nature. He can not only communicate with animals,
		but can also force his will upon them, directing such
		beasts to do as he commands. As the vampire grows in
		power, he can even control the Beast within mortals
		and other supernaturals.`,
		Abilities: map[string]disciplineAbility{
			"Feral Whisper": {
				BaseDesc: `This power is the basis from which all other Animalism abilities grow. The vampire creates an empathic
				connection with a beast, thereby allowing him to com-
				municate or issue simple commands. The Kindred locks
				eyes with the animal, transmitting his desires through
				sheer force of will. The simpler the creature, the more difficult it becomes to connect with the animal’s Beast. Mammals,
				predatory birds, and larger reptiles are relatively easy
				to communicate with. Insects, invertebrates, and most
				fish are just too simple to connect with.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `No roll is necessary to talk with an animal, but the character must establish eye contact first. Issuing commands requires a Manipulation + Animal Ken roll. The difficulty depends on the
					creature: Predatory mammals (wolves, cats, vampire
					bats) are difficulty 6, other mammals and predatory
					birds (rats, owls) are difficulty 7, and other birds and
					reptiles (pigeons, snakes) are difficulty 8. This difficulty
					is reduced by one if the character speaks to the animal
					in its “native tongue,” and can be adjusted further by
					circumstances and roleplaying skill. One success is sufficient to have a cat follow an
					individual and lead the character to the same location,
					three successes are enough to have a raven spy on a
					target for weeks, and five successes ensure that a grizzly
					ferociously guards the entrance to the character’s wil-
					derness haven for some months`,
				},
			},
			"beckoning": {
				BaseDesc: `The vampire’s connection to the Beast grows strong
				enough that he may call out in the voice of a specific
				type of animal — howling like a wolf, shrilling like a
				raven, etc. This call mystically summons creatures of
				the chosen type. Since each type of animal has a dif-
				ferent call, Beckoning works for only a single species
				at a time.
				All such animals within earshot are summoned, and
				some percentage of them will heed the Beckoning if it
				is successful. While the vampire has no further control
				over the beasts who answer, the animals who do are
				favorably disposed toward him and are at least willing
				to listen to the Kindred’s concerns. (The vampire can
				then use Feral Whispers on individual animals to command them.)`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The player rolls Charisma + Survival (dif-
						ficulty 6) to determine the response to the character’s
						call; consult the table below. Only animals that can
						hear the cry will respond. If the Storyteller decides no
						animals of that type are within earshot, the summons
						goes unanswered.
						The call can be as specific as the player desires. A
						character could call for all bats in the area, for only the
						male bats nearby, or for only the albino bat with the
						notched ear he saw the other night`,
					"1": " A single animal responds.",
					"2": "One-quarter of the animals within earshot respond.",
					"3": "Half of the animals respond.",
					"4": "Most of the animals respond.",
					"5": "All of the animals respond.",
				},
			},
			"Quell the Beast": {
				BaseDesc: `As the supreme predators of the natural world, Kindred are highly attuned to the bestial nature that dwells
				within every mortal heart. A vampire who develops
				this power may assert his will over a mortal (animal or
				human) subject, subduing the Beast within her. This
				quenches all powerful, strong emotions — hope, fury,
				love, fear — within the target. The Kindred must either touch his subject or stare into her eyes to channel
				his will effectively. Different Clans evoke this power in different ways,
				though the effect itself is identical. Tzimisce call it
				Cowing the Beast, since they force the mortal’s weaker
				spirit to shrivel in fear before the Kindred’s own inner
				Beast. Nosferatu refer to it as Song of Serenity, since
				they soothe the subject’s Beast into a state of utter
				complacency, thus allowing them to feed freely. Gan-
				grel know the power as Quell the Beast, and force the
				mortal spirit into a state of fear or apathy as befits the
				individual vampire’s nature.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The player rolls Manipulation + Intimidation if forcing down the Beast through fear, or Manipulation + Empathy if soothing it into complacency. The difficulty of the roll is 7 in either case. This is an extended action requiring as many total successes as the
					target has Willpower. Failure indicates that the player
					must start over from the beginning, while a botch indicates that the vampire may not affect that subject’s
					Beast for the remainder of the scene. When a mortal’s Beast is cowed or soothed, she can
					no longer use or regain Willpower. She ceases all struggles, whether mental or physical. She doesn’t even
					defend herself if assaulted, though the Storyteller may
					allow a Willpower roll if the mortal believes her life is
					truly threatened. To recover from this power, the mortal’s player rolls Willpower (difficulty 6) once per day
					until she accumulates enough successes to equal the
					vampire’s Willpower. Kindred cannot be affected by
					this power.
					Though a vampire’s Beast cannot be cowed with this
					ability, the Storyteller may allow characters to use the
					“soothing” variation of this power to pull a vampire out
					of frenzy.`,
				},
			},
			"Subsume the Spirit": {
				BaseDesc: `By locking his gaze with that of an animal, the vampire may mentally possess the creature. Some elders believe that since animals don’t have souls but spirits, the
				vampire can move his own soul into the animal’s body.
				Many younger vampires think it a matter of transfer-
				ring one’s consciousness into the animal’s mind. In either case, it’s agreed that the beast’s weaker spirit (or
				mind) is pushed aside by the Kindred’s own consciousness. The vampire’s body falls into a motionless state
				akin to torpor while his mind takes control of the animal’s actions, remaining this way until the Kindred’s
				consciousness returns.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The player rolls Manipulation + Animal
					Ken (difficulty 8) as the character looks into the ani-
					mal’s eyes. The number of successes allows the character to employ some mental Disciplines while possessing the animal, as noted below.`,
					"1": "Cannot use Disciplines",
					"2": "Can use Auspex and other sensory powers",
					"3": "Can also use Presence and other powers of emotional manipulation",
					"4": "Can also use Dementation, Dominate, and other powers of mental manipulation",
					"5": "Can also use Chimerstry, Necromancy, Thaumaturgy, and other mystical powers",
				},
			},
			"Drawing Out the Beast": {
				BaseDesc: `At this level of Animalism, the Kindred has a keen
				understanding of the Beast Within, and is able to release his feral urges upon another mortal or vampire.
				The recipient of the vampire’s Beast is instantly overcome by frenzy. This is an unnatural frenzy, however,
				as the victim is channeling the Kindred’s own fury. As
				such, the vampire’s own behavior, expressions, and
				even speech patterns are evident in the subject’s savage actions.
				Gangrel and Tzimisce are especially fond of unleashing their Beasts onto others. Gangrel do so to stir their
				ghouls into inspired heights of savagery during combat.
				Tzimisce care less about who receives their Beast than
				retaining their own composure.If the attempt fails, the character himself immediately enters frenzy. As the character relaxes in expectation of relieving his savage urges, the Beast takes that
				opportunity to dig deeper. In this case, the frenzy lasts
				twice as long as normal and is twice as difficult to shrug
				off; its severity also increases exponentially.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The player must announce his preferred tar-
					get (since it must be someone within sight, Drawing
					Out the Beast cannot be used if the vampire is alone),
					then roll Manipulation + Self-Control/Instinct (diffi-
					culty 8). `,
					"1": `The character transfers the Beast, but
					unleashes it upon a random
					individual`,
					"2": `The character is stunned by the effort
					and may not act next turn, but
					transfers the Beast successfully.
					Alternatively, the character may act
					normally during the turn, but must
					spend a Willpower point or suffer a
					single level of lethal damage`,
					"3": `The character transfers the Beast
					successfully.`,
				},
			},
			"Animal Succulence": {
				BaseDesc: `Most vampires find the blood of animals flat, taste-
				less, and lacking in nutritional value. Some Gangrel
				and Nosferatu, however, have refined their under-
				standing of the spirits of such “lesser prey” to the point
				that they are able to draw much more sustenance from
				beasts than normal Kindred can. This power does not
				allow an elder to subsist solely on the blood of animals,
				but it does allow him to go for extended periods of time
				without taking vitae from humans or other Kindred. Animal Succulence does not increase the blood
				point value of the blood of other supernatural creatures
				(Gangrel, werecreatures, and so on) who have taken
				animal forms, nor does it change the vampire’s feeding
				preferences (such as the Ventrue have)`,
				Lvl: 6,
				System: map[string]string{
					"sys": `No roll is needed; once learned, this power
					is always in effect. Animal Succulence allows a charac-
					ter to count each blood point drawn from an animal as
					two in her blood pool. This does not increase the size
					of the vampire’s blood pool, just the nutritional value
					of animal blood. Animal Succulence does not allow a character to
					completely ignore his craving for the blood of “higher”
					prey; in fact, it heightens his desire for “real food.” Ev-
					ery three times (rounded down) the character drinks
					from an animal, a cumulative +1 difficulty is applied
					to the next Self-Control/Instinct roll the player makes
					when the character is confronted with the possibility
					of dining on human or Kindred blood.`,
				},
			},
			"Shared Soul": {
				BaseDesc: `This power allows a character to probe the mind of
				any one animal she touches. Shared Soul can be very
				disconcerting to both parties involved, as each par-
				ticipant is completely immersed in the thoughts and
				emotions of the other. With enough effort or time,
				each participant can gain a complete understanding of
				the other’s mind. Shared Soul is most often used to
				extract an animal’s memories of a specific event, but
				some Gangrel use this power as a tool in the search
				for enlightenment, feeling they come to a better un-
				derstanding of their own Beasts through rapport with
				true beasts. Too close of a bond, however, can leave
				the two souls entangled after the sharing ends, causing
				the vampire to adopt mannerisms, behavior patterns,
				or even ethics (or lack thereof) similar to those of the
				animal.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The character touches the intended subject
					creature, and the player rolls Perception + Animal Ken
					(difficulty 6). The player spends a Willpower point for
					every turn past the first that contact is maintained. Lo-
					cating a specific memory takes six turns, minus one turn
					for every success on the roll. A complete bond takes 10
					turns, minus one turn for every success on the roll. A
					botch on this roll may, at the Storyteller’s discretion,
					send the vampire into a frenzy or give the character a
					derangement related to the behavior patterns of the
					animal (extreme cowardice if the vampire contacted
					the soul of mouse, bloodlust if the subject was a rabid`,
				},
			},
			"Species Speech": {
				BaseDesc: `The basic power Feral Whispers (Animalism 1) al-
				lows character to communicate with only one animal
				at a time. With Species Speech, a character can enter
				into psychic communion with all creatures of a certain
				species that are present. Species Speech is most often
				used after an application of The Beckoning (Animal-
				ism 2), which can draw a crowd of likely subjects.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The player rolls Manipulation + Animal
					Ken (difficulty 7) to establish contact with the tar-
					geted group of animals. Once the character estab-
					lishes contact, the player makes a second roll to issue
					commands. There is no practical upper limit on the
					number of animals that can be commanded with this
					power, although all of the intended subjects must be in
					the vampire’s immediate vicinity. Only one species of
					animal can be commanded at a time; thus, if a charac-
					ter is standing in the middle of the reptile house at the
					zoo, she could command all of the Komodo dragons,
					all of the boa constrictors, or all of the skinks, but she
					could not simultaneously give orders to every reptile
					or snake present.`,
				},
			},
			"Conquer the Beast": {
				BaseDesc: `Masters of Animalism have a much greater under-
				standing of both beasts in general and the Beast in
				particular. Those who have developed this power can
				master their own Beasts to a degree impossible for lesser
				Kindred to attain. Conquer the Beast allows the vam-
				pire both to control her frenzies and to enter them at
				will. Some elders say that the development of this pow-
				er is one of the first steps on the road to Golconda`,
				Lvl: 7,
				System: map[string]string{
					"sys": `The character can enter frenzy at will. The
					player rolls Willpower (difficulty 7). Success sends the
					character into a controlled frenzy. He can choose his
					targets at will, but gains limited Dominate and wound
					penalty resistance and Rötschreck immunity as per the
					normal frenzy rules. A botch on the roll sends
					the vampire into an uncontrolled frenzy which Con-
					quer the Beast may not be used to end.The player may also roll Willpower (difficulty 9) to
					enable the character to control an involuntary frenzy.
					In this case, a Willpower point must be spent for ev-
					ery turn that the vampire remains in frenzy. The player
					may make Self-Control/Instinct rolls as normal to end a
					frenzy, but if the vampire runs out of Willpower points
					before the frenzy ends, he drops into an uncontrolled
					frenzy again.`,
				},
			},
			"Taunt the Caged Beast": {
				BaseDesc: `Some Kindred are so attuned to the Beast that they
				can unleash it in another individual at will. Vampires
				who have developed this power are able to send ad-
				versaries into frenzy with a finger’s touch and the re-
				sultant momentary contact with the victim’s Beast.
				The physical contact allows the vampire’s own Beast
				to reach out and awaken that of the victim, enraging it
				by threatening its territory.`,
				Lvl: 8,
				System: map[string]string{
					"sys": `The character touches the target. The
					player spends a Willpower point and rolls Manipulation + Empathy (difficulty 7). The victim makes a
					Self-Control/Instinct roll (difficulty 5 + the number
					of successes); failure results in an immediate frenzy. A
					botch causes the character to unleash his own Beast
					and frenzy instead. This power maybe used on those
					individuals who are normally incapable of frenzy, send-
					ing ordinary humans into murderous rages worthy of
					the most bloodthirsty Brujah berserker.`,
					"1": " A single animal responds.",
					"2": "One-quarter of the animals within earshot respond.",
					"3": "Half of the animals respond.",
					"4": "Most of the animals respond.",
					"5": "All of the animals respond.",
				},
			},
			"Unchain the Beast": {
				BaseDesc: `The self-destructive nature of Cainites can be turned
				against them by an elder who possesses this formidable
				power. With a glance, the vampire can awaken the
				Beasts of her enemies, causing physical injury and ex-
				cruciating agony as the victim’s own violent impulses
				manifest in physical form to tear him apart from within.
				A target of this power erupts into a fountain of blood
				and gore as claw and bite wounds from an invisible
				source spontaneously tear his flesh asunder.`,
				Lvl: 9,
				System: map[string]string{
					"sys": `The character makes eye contact with the intended victim. The player spends
						three blood points and rolls Manipulation + Intimida-
						tion (difficulty of the victim’s Self-Control/Instinct +
						4). Each success inflicts one health level of aggravated
						damage, which can be soaked normally. A botch in-
						flicts one health level of lethal damage to the invoking
						character for each “1” rolled. This damage can also be
						soaked normally.
						`,
				},
			},
		},
	},
	"auspex": {
		Name: `Auspex`,
		Description: `Auspex gives the vampire uncanny sensory abilities.
		She starts with the capacity to heighten her natural
		senses significantly, but as she grows in power, she can
		perceive psychic auras and read the thoughts of another
		being. Auspex can also pierce through mental illusions
		such as those created by Obfuscate.However, a vampire with Auspex needs to be careful.
		Her increased sensory sensitivity can cause her to be
		drawn in by beautiful things or stunned by loud nois-
		es or pungent smells. Sudden or dynamic events can
		disorient an Auspex-using character unless her player
		makes a Willpower roll to block them out (difficulty of
		at least 4, although the more potent the source of dis-
		traction, the higher the difficulty). `,
		Abilities: map[string]disciplineAbility{
			"Heightened Senses": {
				BaseDesc: `This power increases the acuity of all of the vam-
				pire’s senses, effectively doubling the clarity and range
				of sight, hearing, and smell. While her senses of taste
				and touch extend no farther than normal, they like-
				wise become far more distinct; the vampire could taste
				the hint of liquor in a victim’s blood or feel the give of
				the board concealing a hollow space in the floor. The
				Kindred may magnify her senses at will, sustaining this
				heightened focus for as long as she desires. Occasionally, this talent provides extrasensory or
				even precognitive insights. These brief, unfocused
				glimpses may be odd premonitions, flashes of empathy,
				or eerie feelings of foreboding. The vampire has no
				control over these perceptions, but with practice can
				learn to interpret them with a fair degree of accuracy.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `It takes a reflexive action to activate this
					ability, but no roll or other cost is required. In certain
					circumstances, dice rolls associated with using the
					character’s sense (such as Perception + Alertness) de-
					crease in difficulty by a number equal to the character’s
					Auspex rating when the power is engaged.
					The Storyteller may also use this power to see if the
					character perceives a threat. In this case, the Story-
					teller privately rolls the character’s unmodified Auspex
					rating, applying whatever difficulty he feels best suits
					the circumstances. For example, sensing that a pistol is
					pointed at the back of the character’s head may require
					a roll of difficulty 5, while the sudden realization that
					a rival for Primogen is planning her assassination may
					require a 9.`,
				},
			},
			"Aura Perception": {
				BaseDesc: `Using this power, the vampire can perceive the psy-
				chic “auras” that radiate from mortals and supernatural
				beings alike. These halos comprise a shifting series of
				colors that take practice to discern with clarity. Even
				the simplest individual has many shifting hues within
				his aura; strong emotions predominate, while momen-
				tary impressions or deep secrets flash through in streaks
				and swirls. The colors change in sympathy with the subject’s
				emotional state, blending into new tones in a con-
				stantly dancing pattern. The stronger the emotions
				involved, the more intense the hues become. A skilled
				vampire can learn much from her subject by reading
				the nuances of color and brilliance in the aura’s flow`,
				Lvl: 2,
				System: map[string]string{
					"sys": `After the character stares at the subject for
					at least a few seconds, the player rolls Perception + Em-
					pathy (difficulty 8); each success indicates how much
					of the subject’s aura the character sees and understands
					(see the table below). A failure indicates that the play
					of colors and patterns yields no prevailing impression.
					A botch indicates a false or erroneous interpretation.`,
					"1": `Can distinguish only the shade (pale
						or bright).`,
					"2": `Can distinguish the main color.`,
					"3": `Can recognize the color patterns.`,
					"4": `Can detect subtle shifts.`,
					"5": `Can identify mixtures of color and
					pattern.`,
				},
			},
			"The Spirit's Touch": {
				BaseDesc: `When someone handles an object for any length
				of time, he leaves a psychic impression on the item.
				A vampire with this level of Auspex can “read” these
				sensations, learning who handled the object, when he
				last held it, and what was done with it recently. These visions are seldom clear
				and detailed, registering more like a kind of “psychic
				snapshot.” Still, the Kindred can learn much even
				from such a glimpse. Although most visions concern
				the last person to handle the item, a long-time owner
				leaves a stronger impression than someone who held
				the object briefly, `,
				Lvl: 3,
				System: map[string]string{
					"sys": `The player rolls Perception + Empathy. The
					difficulty is determined by the age of the impressions
					and the mental and spiritual strength of the person or
					event that left them. Sensing information from a pis-
					tol used for a murder hours ago may require a 4, while
					learning who owned a bloodstained puppet fashioned a
					century ago might be a 9.
					The greater the individual’s emotional connection
					to the object, the stronger the impression he leaves
					on it — and the more information the Kindred can
					glean from it. Events involving strong emotions (a gift-
					giving, a torture, a long family history) likewise leave
					stronger impressions than short or casual contact do.
					Assume that each success offers one piece of informa-
					tion, as per the chart below.`,
					"Botch": `The character is overwhelmed by
					psychic impressions for the next 30
					minutes and unable to act.`,
					"fail": `No information of value.`,
					"1": `Very basic information: the last
					owner’s gender or hair color,
					for instance.`,
					"2": `A second piece of basic information.`,
					"3": `More useful information about the
					last owner, such as age and state of
					mind the last time he used the item.`,
					"4": `The person’s name.`,
					"5": `A wealth of information: nearly
					anything you want to know about the
					person’s relationship with that object
					is available`,
				},
			},
			"Telepathy": {
				BaseDesc: `The vampire projects a portion of her consciousness
				into a nearby mortal’s mind, creating a mental link
				through which she can communicate wordlessly or
				even read the target’s deepest thoughts. The Kindred
				“hears” in her own mind the thoughts plucked from a
				subject as if they were spoken to her.
				This is one of the most potent vampiric abilities,
				since, given time, a Kindred can learn virtually any-
				thing from a subject without him ever knowing. The
				Tremere and Tzimisce in particular find this power es-
				pecially useful in gleaning secrets from others, or for
				directing their mortal followers with silent precision.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The player rolls Intelligence + Subterfuge
					(difficulty of the subject’s current Willpower points).
					Projecting thoughts into the target’s mind requires one
					success. The subject recognizes that the thoughts come
					from somewhere other than his own consciousness,
					though he cannot discern their actual origin without a
					successful Perception + Awareness roll (difficulty equal
					to the vampire’s Manipulation + Subterfuge).
					To read minds, one success must be rolled for each
					item of information plucked or each layer of thought
					pierced. Deep secrets or buried memories are harder to
					obtain than surface emotions or unspoken comments,
					requiring five or more successes to access. Reading thoughts with Telepathy does not com-
					monly work upon the undead mind. A character may
					expend a Willpower point to make the effort, mak-
					ing the roll normally afterward. Likewise, it is equally
					difficult to read the thoughts of other supernatural
					creatures`,
				},
			},
			"Psychic Projection": {
				BaseDesc: `The Kindred with this awesome ability projects her
				senses out of her physical shell, stepping from her body
				as an entity of pure thought. The vampire’s astral form
				is immune to physical damage or fatigue, and can “fly”
				with blinding speed anywhere across the earth — or
				even underground — so long as she remains below the
				moon’s orbit.
				The Kindred’s material form lies in a torpid state
				while her astral self is active, and the vampire isn’t
				aware of anything that befalls her body until she returns
				to it. An ephemeral silver cord connects the Kindred’s
				psyche to her body. If this cord is severed, her consciousness becomes stranded in the astral plane (the realm of
				ghosts, spirits, and shades).`,
				Lvl: 5,
				System: map[string]string{
					"sys": ` Journeying in astral form requires the player
					to expend a point of Willpower and make a Perception
					+ Awareness roll. Difficulty varies depending on the dis-
					tance and complexity of the intended trip; 5 is within
					sight, 7 is nearby or to a familiar location, and 9 reflects a
					trip far from familiar territory (a first journey from North
					America to the Far East; trying to shortcut through the
					earth). The greater the number of successes rolled, the
					more focused the character’s astral presence is, and the
					easier it is for her to reach her desired destination.Failure means the character is unable to separate
					her consciousness from her body, while a botch can
					have nasty consequences — flinging her astral form to
					a random destination on Earth or in the spirit realm,
					arriving in a place where the sun is active (necessitat-
					ing a frenzy roll, although the sunlight doesn’t do any
					damage), or hurtling toward the desired destination so
					forcefully that the silver cord snaps.`,
				},
			},
			"Clairvoyance": {
				BaseDesc: `By using Clairvoyance, a vampire can perceive dis-
				tant events without using Psychic Projection. By con-
				centrating on a familiar person, place, or object, a
				character can observe the subject’s immediate vicinity
				while staying aware of her own surroundings.`,
				Lvl: 6,
				System: map[string]string{
					"sys": ` The player rolls Perception + Empathy (difficulty 6) and describes the target she’s trying to look
						in on. If the roll is successful, the character can then
						perceive the events and environment surrounding the
						desired target for one turn per success. Other Auspex
						powers may be used on the scene being viewed; these are
						rolled normally. Clairvoyance does split the vampire’s
						perceptions between what she is viewing at a distance
						and what is taking place around her. As a result, while
						using this power, a character is at +3 difficulty on all rolls
						relating to actions that affect her physical surroundings.`,
				},
			},
			"Prediction": {
				BaseDesc: `Some people are capable of finishing their friends’
				sentences. Elder vampires with Prediction sometimes begin their friends’ sentences. Prediction is a constant
				low-level telepathic scan of the minds of everyone the
				character is in proximity to. While this power does not
				give the vampire the details of his neighbors’ conscious
				thoughts, it does provide a wealth of cues as to the subjects’ moods, suppressed reflexes, and attitudes toward
				the topic of conversation.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `Whenever the character is in conversation
					and either participant in the discussion makes a Social
					roll, the player may pre-empt the roll to spend a blood
					point and make a Perception + Empathy roll (difficulty of the target’s Manipulation + Subterfuge). Each
					success is an additional die that can be applied to the
					player’s Social roll or subtracted from the dice pool of
					the Social roll being made against the character.`,
				},
			},
			"Telepathic Communication": {
				BaseDesc: `Telepathy allows a character to pick up only
				the surface thoughts of other individuals, and to speak
				to one at a time. With Telepathic Communication,
				a character can form a more powerful link between
				his mind and that of other subjects, allowing them to
				converse in words, concepts, and sensory images at the
				speed of thought (and without the need for Willpower
				expenditure, unlike with Telepathy). Vampires with
				this level of Auspex can act as “switchboard operators,”
				creating a telepathic web that allows all participants to
				share thoughts with some or all other members of the
				network as they choose.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The player rolls Charisma + Empathy (dif-
						ficulty equals the target’s current Willpower points) to
						establish contact, although a willing subject may allow
						the vampire access and thus obviate the need for a roll.
						The maximum range at which a subject may be con-
						tacted and the maximum number of individuals who
						may be linked simultaneously with this power depends
						on the Auspex rating of the vampire who initiates con-
						tact `,
					"auspex 6": `3 subjects, 500 miles/800 kilometers`,
					"auspex 7": `num of subjects equal to Perception rating, 1000 miles/1500 kilometers`,
					"auspex 8": `num of subjects equal to Perception + Empathy, 500 miles, 800 kilometers per point of Intelligence`,
					"auspex 9": `2x Perception + Empathy, 1000 miles, 1500 kilometers per point of Intelligence`,
				},
			},
			"Karmic Sight": {
				BaseDesc: `The power of Aura Perception (Auspex 2) allows a
				vampire to take a brief glimpse at the soul of a sub-
				ject. This power takes Aura Perception several steps
				forward, allowing a vampire who has mastered Auspex
				2 to probe the inner workings of a subject’s mind and
				soul`,
				Lvl: 7,
				System: map[string]string{
					"sys": `The player rolls Perception + Empathy (dif-
						ficulty equals the subject’s current Willpower). The de-
						gree of success determines the information gained.`,
					"Botch": `The character gains a Derangement
					or Psychological/Mental/
					Supernatural Flaw similar to
					one of the target’s for one
					night, at Storyteller discretion.`,
					"1": `As per five successes on an
					Aura Perception roll.`,
					"2": `Subject’s Nature, Demeanor,
					and Humanity or Path can be
					determined.`,
					"3": `Any outside influences on the
					subject’s mind or soul, such as
					Dominate or a demonic pact,
					can be detected.`,
					"4": `Subject’s Willpower, Humanity or
					Path, and Virtue ratings can
					be determined.`,
					"5": `The state of the subject’s
					karma may be determined.
					This is a highly abstract piece
					of information best left to
					Storyteller discretion, but
					should reveal the general
					balance between “good” and
					“bad” actions the subject has
					performed, both recently and
					over the course of his
					existence. If the plot merits
					it, the character may receive
					visions of one or more
					incidents in the subject’s past
					that radically altered his destiny.
					With this degree of success, some
					fate-related Merits and Flaws (e.g.
					Dark Fate) can be identified as well.`,
				},
			},
			"Mirror Reflex": {
				BaseDesc: `This power was developed by a Toreador elder who
				made a fearsome reputation through her fencing prow-
				ess, acting as a hired champion in dozens of Ventrue
				duels. Mirror Reflex is similar to Prediction in that it is
				in essence a low-level telepathic scan of an opponent,
				but this power taps into physical (rather than social)
				reflexes, allowing the character to anticipate an en-
				emy’s moves in personal combat.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `The player spends a blood point and rolls
					Perception + the combat skill the opponent is using
					(difficulty of the subject’s Manipulation + combat skill
					in use). Each success is an additional die that can he ap-
					plied to the character’s dice pools during the next turn
					of combat for any actions taken against the scanned op-
					ponent. The use of Mirror Reflex does take one combat
					action, and the power has a maximum range in yards or
					meters equal to the character’s Willpower rating`,
				},
			},
			"Psychic Assault": {
				BaseDesc: `Psychic Assault is nothing less than a direct mind-to-mind attack which uses the sheer force of an elder’s
				will to overpower his target. Victims of Psychic Assault
				show little outward sign of the attack, save for nosebleeds and expressions of intense agony; all injuries by
				means of this psychic pressure inflicted are internal. A
				medical examination of a mortal victim of a Psychic
				Assault invariably shows the cause of death to be a
				heart attack or aneurysm, while vampires killed with
				this power decay to dust instantly, regardless of age.`,
				Lvl: 8,
				System: map[string]string{
					"sys": `The character must touch or make eye con-
					tact with his target. The player spends three blood
					points (and a Willpower point, if assaulting a vampire
					or other supernatural being) and rolls Manipulation +
					Intimidation in a contested roll against the victim’s
					Willpower. The result depends on the number of net
					successes the attacker rolls.`,
					"Botch": `The target becomes immune to the
					attacker’s Psychic Assault for one
					night per each “1” rolled.`,
					"Fail": `The target is unharmed and may
					determine that a psychic assault is
					taking place by succeeding on a
					Perception + Awareness roll
					(difficulty 6).`,
					"1": `The target is shaken but unharmed.
					He loses a temporary Willpower
					point.`,
					"2": `The target is badly frightened. He
					loses three temporary Willpower
					points and, if a vampire, must roll
					Courage (difficulty of the attacker‘s
					Auspex rating) to avoid Rötschreck`,
					"3": `The target loses six temporary
					Willpower points and, if a vampire,
					must roll Courage as above. If this
					causes him to lose his last temporary
					Willpower point, he loses a
					permanent Willpower point and
					receives three health levels of bashing
					damage (soaked normally).`,
					"4": `The target loses all temporary
					Willpower points and half of his
					permanent Willpower points (round
					down) and suffers three health levels
					of lethal damage (soaked normally).`,
					"5": `The target must roll his permanent
					Willpower (difficulty 7). If he
					succeeds, apply the effect of four
					successes, and is also rendered
					unconscious for the rest of the night.
					If he fails, the Psychic Assault kills
					him instantly.`,
				},
			},
			"False Slumber": {
				BaseDesc: `Possibly the source of many Malkavians’ conviction
				that their sire is alive and well on the astral plane, this
				power allows a Methuselah‘s spirit to leave his body while
				in torpor. While seemingly asleep, the vampire is able to
				project astrally, think, and perceive events normally`,
				Lvl: 9,
				System: map[string]string{
					"sys": `No roll is needed. This power is considered
					to be active whenever the vampire’s body is in torpor,
					and astral travels are handled as per the rules for Psy-
					chic Projection. The vampire is not able to awaken
					physically at will, however — waking from torpor is
					handled per the normal rules for such an action.A vampire with this power whose silver cord is severed in
					astral combat loses all Willpower points, as per the rules for
					astral combat under Psychic Projection, but is not killed.
					Instead, he loses the use of this Auspex power and half of
					his permanent Willpower points. `,
				},
			},
			"Seeing the Unseen": {
				BaseDesc: `Auspex enables Kindred to perceive many
				things beyond the limits of lesser senses.
				Among its many uses, Auspex can detect
				the presence of a supernatural being who is
				hidden from normal sight (a vampire using
				Obfuscate, for example, or a ghost) or pierce
				illusions created by the Discipline of Chim-
				erstry. Note: “Normal sight” includes regular,
				non-Auspex use of the Awareness skill.`,
				Lvl: 1,
				System: map[string]string{
					"Obsfucate": ` When a vampire tries to use her
					heightened perceptions to notice a Kindred
					hidden with Obfuscate, she detects the sub-
					ject’s presence if her Auspex rating is higher
					than his Obfuscate, and she succeeds at a
					Perception + Awareness roll (difficulty equals
					7 minus the number of dots by which her
					Auspex exceeds his Obfuscate). Conversely,
					if the target’s Obfuscate outranks her Auspex,
					he remains undiscovered. If the two ratings
					are equal, both characters make a resisted roll
					of Perception + Awareness (Auspex user)
					against Manipulation + Subterfuge (Obfuscate
					user). The difficulty for both rolls is 7, and the
					character with the most successes wins.`,
					"Chimesrstry": `Likewise, vampires with
					Auspex may seek to penetrate illusions cre-
					ated with Chimerstry. The Auspex-wielder
					must actively seek to pierce the illusion (i.e.,
					the player must tell the Storyteller that his
					character is trying to detect an illusion).
					The Auspex-user and Chimerstry-wielder
					then compare relative ratings, per Obfuscate,
					above. The process is otherwise identical to
					piercing Obfuscate.`,
					"Other Powers": `Since the powers of beings
					like magi and wraiths function differently
					from vampiric Disciplines, a simple compari-
					son of relative ratings isn’t applicable. To keep
					things simple, both characters make a resisted
					roll. The vampire rolls Perception + Aware-
					ness, while the subject rolls Manipulation +
					Subterfuge. Again, the difficulty is 7, and the
					character with the most successes wins.`,
				},
			},
		},
	},
	"celerity": {
		Name: `Celerity`,
		Description: `Not all vampires are slow, meticulous creatures.
		When needed, some vampires can move fast — really
		fast. Celerity allows Assamites, Brujah, and Toreadors
		to move with astonishing swiftness, becoming practi-
		cally a blur. The Assamites use their speed in conjunc-
		tion with stealth to strike quickly and viciously from
		the shadows before they are noticed. Brujah, on the
		other hand, simply like the edge that the power gives
		them against overwhelming odds. The Toreador are
		more inclined to use Celerity to provide an air of un-
		natural grace to live performances or for an extra push
		to complete a masterpiece on time, but they can be
		as quick to draw blood as any assassin or punk when
		angered.`,
		Abilities: map[string]disciplineAbility{
			"Celerity 1-5": {
				BaseDesc: `Normally, a character without Celerity must divide
				their dice if she wants to take multiple actions in a
				single turn. A character using Celerity
				performs his extra actions (including full movement)
				without penalty, gaining a full dice pool for each sepa-
				rate action. Extra actions gained through Celerity may
				not in turn be split into multiple actions, however.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `Each point of Celerity adds one die to every
					Dexterity-related dice roll. In addition, the player can
					spend one blood point to take an extra action up to the
					number of dots he has in Celerity at the beginning of
					the relevant turn; this expenditure can go beyond her
					normal Generation maximum. Any dots used for extra
					actions, however, are no longer available for Dexteri-
					ty-related rolls during that turn. These additional ac-
					tions must be physical (e.g., the vampire cannot use a
					mental Discipline like Dominate multiple times in one
					turn), and extra actions occur at the end of the turn
					(the vampire’s regular action still takes place per her
					initiative roll).
					`,
				},
			},
			"Projectile": {
				BaseDesc: `Despite the fact that a vampire with Celerity moves
				at incredible speeds, any bullets he fires or knives he
				throws while in this state don’t move any faster than
				they normally would. Scientifically minded Kindred
				have been baffled by the phenomenon for centuries,
				but more pragmatic ones have found a way to work
				around it. Projectile enables a vampire to take his pre-
				ternatural speed and transfer it into something he has
				thrown, fired, or launched.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `Projectile requires the expenditure of a blood
					point. In addition, the player must decide how many
					levels of his character’s Celerity he is putting into the
					speed of the launched object. Thus, a character with
					Celerity 6 in addition to Projectile could decide to put three dots’ worth of speed into a knife he is throwing,
					and use the other three dots as dice or potential extra
					actions as per normal. Each dot of Celerity infused into
					a thrown object becomes an automatic success to the
					attack’s damage roll, assuming the weapon or projectile
					actually hits.`,
				},
			},
			"Flower of Death": {
				BaseDesc: `In combat, speed kills. A proper application of Celerity in combat can turn even the meekest Cainite
				into a walking abattoir. How much more deadly, then,
				is a vampire with the ability to utilize his preternatural
				speed to the utmost in combat? Flower of Death allows
				a vampire to take his Celerity and apply it in full to
				each hand-to-hand or melee attack he makes.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `Flower of Death costs four blood points, but
					the spectacular effect is well worth it. Once the power
					is in effect, the vampire’s bonus dice for Dexterity rolls
					get added to every dice pool for attack the character
					makes (even if the roll doesn’t use Dexterity) until the
					end of the scene. Further, even if the Kindred uses some
					of his Celerity dots for extra actions during the scene,
					these extra dice are still available. The effect is limited
					to hand-to-hand or melee weapon attacks — firearms,bows, and other ranged weapons are excluded — but
					does grant the attacker additional dice for damage rolls.
					Flower of Death is not cumulative — it is impossible
					to “layer” uses of the power over one another to create
					astronomical dice pools`,
				},
			},
			"Zephyr": {
				BaseDesc: `Zephyr produces an effect vaguely similar to one of
				the legendary comic book-style uses of enhanced speed,
				allowing its practitioner to run so fast he can run across
				water. Particularly successful applications of Zephyr allow a vampire to go so far as to run up walls and, in at
				least one recorded instance, across a ceiling. Most times, a vampire moving at such a rate of speed
				is barely visible, appearing more as a vampire-shaped
				blur than anything else. Observers must succeed on a
				Perception + Alertness roll (difficulty 7) to get a de-
				cent look at a Kindred zooming past in this fashion`,
				Lvl: 8,
				System: map[string]string{
					"sys": ` Zephyr requires the expenditure of one point
					of blood and one point of Willpower. Unfortunately,
					Zephyr requires such extremes of concentration that it
					cannot be combined with any form of attack, or indeed,
					with most any sort of action at all. If a character using
					Zephyr feels the need to do something else while mov-
					ing at such tremendous speeds, a Willpower roll (difficulty 8) is required. Needless to say, botches at Zephyr
					speed can be spectacular in all the wrong ways`,
				},
			},
		},
	},
	"chimerstry": {
		Name: `Chimerstry`,
		Description: ` The Discipline
		is, fundamentally, an art of conjuration that converts
		the vampire’s will into phantoms that confound the
		senses and technology alike. Even vampires fall under
		the sway of the Ravnos’ illusory world, unless they have
		a strong enough grasp of Auspex. The Ravnos often use this power to swindle and seduce their
		victims into acts that work out badly for the victim. Illusions created by Chimerstry can be seen for what
		they are by a victim who “proves” the illusion’s false-
		hood (e.g., a person who walks up to an illusory wall,
		expresses his disbelief in it, and puts his hand through
		it effectively banishes the illusion), and explicitly in-
		credible illusions are seen as false immediately (e.g.,
		dragons breathing fire or gravity working in reverse).
		Sometimes, frequent targets of Chimerstry end up at-
		tempting to disbelieve everything around them, lead-
		ing to derangements.`,
		Abilities: map[string]disciplineAbility{
			"Ignis Fatuus": {
				BaseDesc: `The vampire may conjure a minor, static mirage that
				confounds one sense. For instance, he may evoke a sul-
				furous stench, the appearance of stigmata, or the shat-
				ter of broken glass. Note that though tactile illusions
				can be felt, they have no real substance; an invisible
				but tactile wall cannot confine anyone, and invisible
				razor-wire causes no real damage. Similarly, the vam-
				pire must know the characteristics of what he’s creat-
				ing. While it’s easy enough to estimate what a knife
				wound might look like, falsifying a person’s voice or a
				photograph of a childhood home requires knowledge
				of the details.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The player spends a point of Willpower
					for the vampire to create this illusion. The volume of
					smells, ambient lighting, smoke clouds, and the like
					are limited to roughly 20 cubic feet (half a cubic me-
					ter) per dot the vampire has in Chimerstry. The illu-
					sion lasts until the vampire leaves its vicinity (such as
					stepping out of the room) or until another person sees
					through it somehow. The Cainite may also end the il-
					lusion at any time with no effort.`,
				},
			},
			"Fata Morgana": {
				BaseDesc: `The Cainite can now create illusions that appeal to
				all the senses, though they remain static. For example,
				the vampire could make a filthy cellar appear as an opulent ballroom, though she could not create a glittering
				chandelier or a score of graceful dancers. Again, the
				illusion has no solid presence, though it’s easy enough
				to fool an enraptured visitor with suggestions of what
				she might expect. A bucket of brackish water is as cool
				as chilled champagne, after all.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The player spends a Willpower point and a
					blood point to create the illusion. These static images
					remain until dispelled, in much the same way that an
					Ignis Fatuus illusion does.`,
				},
			},
			"Apparition": {
				BaseDesc: `Not really a power unto itself, Apparition allows a
				vampire to give motion to an illusion created with Ig-
				nis Fatuus or Fata Morgana. Thus, the Cainite could
				create the illusion of a living being, running water,
				fluttering drapes, or a roaring fire.Once the creator stops concentrating on the illusion,
				it can continue in simple, repetitive motions – roughly
				speaking, anything that can be described in a simple
				sentence, such as a guard walking back and forth in
				front of a steel door. After that, the vampire cannot
				regain control over the illusion – she can either allow
				it to continue moving as ordered, or let it fade as de-
				scribed under Ignis Fatuus.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The creator spends one blood point to make
					the illusion move in one significant way, or in any
					number of subtle ways. For example, the vampire could
					create the illusion of a lurking mugger lurching at her
					victim, or she could create the illusion of a desolate
					street, down which a chill wind blows trash while a
					streetlamp flickers and hums. Taking complicated actions besides maintaining the illusion — that is, anything that would require a dice roll — first requires success on a Willpower roll, resulting in the dissolution of
					the false construct if the roll fails.`,
				},
			},
			"Permanency": {
				BaseDesc: `This power, also used in conjunction with Ignis
				Fatuus or Fata Morgana, allows a mirage to persist even
				when the vampire cannot see it. In this way, Ravnos
				often cloak their temporary havens in false trappings of
				luxury, or ward off trespassers with illusory guard dogs.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The vampire need only spend a blood point,
					and the illusion becomes permanent until dissolved
					(including “programmed” illusions like those created
					by Apparition).`,
				},
			},
			"Horrid Reality": {
				BaseDesc: `Rather than create simple illusions, the vampire
				can now project hallucinations directly into a victim’s
				mind. The target of these illusions believes completely
				that the images are real; a hallucinatory fire can burn
				him, an imaginary noose can strangle him, and an illusory wall can block him. This power affects only one
				person at a time; though others can see the illusion, it
				doesn’t impact them in the same way. Other people
				can try to convince the victim that his terrors are not
				real, but he won’t believe them. Note that targets with
				enough dots in Auspex can still attempt to roll for See-
				ing the Unseen`,
				Lvl: 5,
				System: map[string]string{
					"sys": `A Horrid Realty illusion costs two Willpower points to set in motion and lasts for an entire
					scene (though its effects may last longer; see below).
					If the vampire is trying to injure his victim, his player
					must roll Manipulation + Subterfuge (difficulty of the victim’s Perception + Self-Control/Instinct). Each
					success inflicts one health level of lethal damage on
					the victim that cannot be soaked — the Cainite as-
					saults the victim’s mind and perceptions, not his body.
					If the player wishes to inflict less damage or change it
					to bashing, he may announce a maximum amount of
					damage before rolling the dice. Secondary effects (such
					as frenzy rolls for illusory fire) may also occur.The victim heals all his damage instantaneously if
					he can be convinced that the damage he took was il-
					lusory, but convincing him may take some doing, such
					as with at least two successes on a Charisma + Empathy
					roll (difficulty equal to the Manipulation + Subterfuge
					of the Cainite using Horrid Reality). The target must
					be convinced of the attack’s illusory nature within 24
					hours of its taking place, or it becomes too well estab-
					lished in his memory, and he will have to heal the dam-
					age using blood (if a vampire) or over time (if mortal).
					This power cannot actually kill its victims (though a
					target with a heart condition may well die from fright).
					A victim “killed” by an illusory attack loses conscious-
					ness or enters torpor.`,
				},
			},
			"False Resonance": {
				BaseDesc: `Illusions of living or unliving beings are all well
				and good until someone decides to read the illusion’s
				mind or its aura. The automatic failure to perceive any
				sense of the target’s thoughts or emotions will usually
				be passed off as bad luck, lack of concentration, or
				whatever reason any Kindred might construct to explain why he didn’t succeed in gleaning information
				through supernatural means. A vampire can use False
				Resonance to overlay auras and thoughts on illusions,
				as well as leave a trace that other emotionally resonant
				powers can detect later.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `This power automatically applies to any
					other use of Chimerstry as the user wishes. In effect,
					any attempt to use Auspex, the Dementation power
					Eyes of Chaos, or similar sensory powers that generates
					five or fewer successes will detect an aura, thoughts,
					Demeanor or whatever the power would normally de-
					tect. Thoughts won’t be exceptionally complex, and
					will relate to whatever is going on around the illusion
					in a mundane and simplistic way. Auras will consist
					of colors related to specific emotions (anger, sadness,
					hatred, love, and happiness) and will not show much
					complexity beyond that. Spirit’s Touch can pick up the
					same emotional resonance until the next sunrise.`,
				},
			},
			"Fatuus Mastery": {
				BaseDesc: `A Ravnos with Fatuus Mastery has no restriction
				on how often she may use the first three levels (Ignis
				Fatuus, Fata Morgana, and Apparition) and can maintain or control illusions with minimal concentration
				or fatigue. Kindred who rely on the high cost of Chimerstry to limit a vampire’s ability to use illusions are in
				for a very rude surprise when they encounter a Cainite
				with this power.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `Fatuus Mastery negates the Willpower and
					blood cost for using the first three levels of Chimerstry.
					In addition, the Kindred may direct movement for a
					number of illusions equal to his Intelligence without
					intense concentration. Furthermore, the character can
					maintain the illusion as long as it remains within his
					Willpower rating in miles (or about one and a half
					times that in kilometers), although he may not make
					it react to events around it if he has no way to perceive
					those events.`,
				},
			},
			"Shared Nightmare": {
				BaseDesc: `Even though Horrid Reality is visible to all onlookers, it can only inflict “damage” on one victim. With
				Shared Nightmare, a vampire can inflict her tormented
				visions on a crowd.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `To use this power, the player must spend
					two Willpower points, plus one blood point per target.
					The player rolls Manipulation + Subterfuge once, but
					compares the results against each target individually.
					The difficulty is still each victim’s Perception + Self-
					Control/Instinct.`,
				},
			},
			"Far Fatuus": {
				BaseDesc: `This power allows a Kindred to project illusions to
				any area he can see or visualize. Under most circumstances, accomplishing this requires him to have visited the location in question before he can project illusions there. Although more difficult, a vampire may
				project illusions on the basis of a description, a photo,
				or a video clip`,
				Lvl: 7,
				System: map[string]string{
					"sys": `The difficulty for using Far Fatuus depends
					on the user’s familiarity with the location. The player
					must roll Perception + Subterfuge to affect the loca-
					tion. Once this roll is successful, the vampire may then
					use any other Chimerstry power on that location.`,
					"Difficulty 6": `As familiar as one’s haven; currently
					viewing with Clairvoyance or Psychic
					Projection`,
					"Difficulty 7": `Visited three or more times.`,
					"Difficulty 8": `Visited once; viewing on a live feed`,
					"Difficulty 9": `Described in detail; seen it in a video
					or have a undoctored photo`,
				},
			},
			"Synesthesia": {
				BaseDesc: `A Cainite who masters this power can shuffle oth-
				ers’ senses around to suit his preferences. He can select one target and inflict a serious, disorienting, and
				all-encompassing case of synesthesia upon her, making it impossible for her to interact meaningfully with
				the real world for the power’s duration. The vampire
				has complete control over how the target’s senses work
				and can manipulate them to suit. For example, he may
				decide that she smells all sounds as varieties of nauseating stenches, or more subtly, he may exchange pain for
				pleasure. Used against a crowd, sensations are randomly shuffled, so one man will see what the woman next to him sees, but hears what the man 15 feet behind him
				hears and feels what the child a block away feels. The
				end result is extremely disorienting for all victims.`,
				Lvl: 8,
				System: map[string]string{
					"sys": ` When used against a single victim, the play-
					er must spend one Willpower point and roll Manipula-
					tion + Intimidation (difficulty is victim’s current Willpower points). For use against crowds, the difficulty is
					7, and the power affects everyone within the vampire’s
					line of sight and subtracts one point from Perception
					per success rolled. Victims whose Perception has been
					reduced to zero can only sit down and wait for the disorientation to end. Duration against a single victim is
					determined below. Against a crowd, the power persists
					until sunrise.`,
					"1": `One week`,
					"2": `One month`,
					"3": `Six months`,
					"4": `One year`,
					"5": `Permanent`,
				},
			},
			"Mayaparisatya": {
				BaseDesc: `This expression of Chimerstry allows the Cainite
				to directly alter or create real objects or creatures, although such changes are of finite duration. A vampire
				with this power can transform the air around a rival
				Kindred into fire or render a locked door insubstantial.
				A more harrowing use of this power enables the vampire to force an object out of existence by transforming
				it into nothing more than a wisp of its former reality.`,
				Lvl: 9,
				System: map[string]string{
					"sys": `To use this power, the player must spend 10
					blood points and one permanent Willpower point and
					roll Manipulation + Subterfuge. Difficulty for the roll
					is 6 for affecting inanimate objects, and the victim’s
					Willpower rating for affecting characters. This power
					can affect anything within miles (kilometers) of the
					vampire, as long as the character is aware of the target
					in some way. If used with Far Fatuus, the effects are
					centered on the chosen location. This power can affect
					a number of conscious targets equal to the Kindred’s
					Willpower per use. When dealing with inanimate objects, the number
					of successes determines how drastic the alteration may
					be. No matter how many successes the player rolls, the
					duration is always a scene. This power can affect any
					objects of a type within the vampire’s targeted area.`,
					"1": `Render an object harmless (swords
						won’t cut, firearms won’t shoot),
						create a large volume of obscuring
						smok`,
					"2": `Change an object into another object
					(turn candles into tarantulas, etc.)`,
					"3": `Render the object insubstantial, make
					smoke solid`,
					"4": `Cause drastic changes (stone becomes
						highly flammable)`,
					"5": `Cause the environment to behave
					illogically (gravity twists sideways,
					rivers stand still as hills flow upward)`,
					"6+": `Delete any offending material objects
					from existence. This effect is
					permanent (for usage on concious targets reference concious.`,
					"concious": `When using the power on conscious targets, consult
					the table above for alterations (such as forcing the vic-
					tim into another form or transforming her into a differ-
					ent substance). If using the power to negate the victim’s
					existence, the power inflicts two levels of unsoakable
					aggravated damage per success. If the power doesn’t kill
					the victim, subtract one dot of Strength and Stamina
					per success. The damage must be healed normally, but
					the lost Attributes return at the end of the scene. Vic-
					tims of this power look hazy and insubstantial. Victims
					destroyed with this power simply vanish`,
				},
			},
		},
	},
	"dementation": {
		Name: `Dementation`,
		Description: `Dementation is the Discipline that allows a vampire
		to focus and channel madness into the minds of those
		around him. Though it’s the natural legacy of the Malkavians, practitioners of Dementation need not actually be mad to use the Discipline… but it helps.
		Disturbingly, Dementation doesn’t actually make
		their victims mad, but rather it seems to break down
		the doors to the hidden darkness of the target’s mind,
		releasing into the open whatever is found there. The
		Malkavians claim that this is because insanity is the
		next logical step in mental evolution, a transhuman-
		ist advancement of what modern people consider con-
		sciousness`,
		Abilities: map[string]disciplineAbility{
			"Passion": {
				BaseDesc: `The vampire stirs his victim’s emotions, either
				heightening them to a fevered pitch or blunting them
				until the target is completely desensitized. The Cainite may not choose which emotion is affected; she may
				only amplify or dull emotions already present in the
				target. In this way, a vampire can inflame mild irritation into quivering rage or atrophy true love into casual interest.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The character talks to her victim, and
					the vampire’s player rolls Charisma + Empathy (dif-
					ficulty equals the victim’s Humanity or Path rating).
					The number of successes determines the duration of
					the altered state of feeling. Effects of this power might
					include one- or two-point additions or subtractions
					to difficulties of frenzy rolls, Virtue rolls, rolls to resist
					Presence powers, etc.`,
					"1":  `One turn`,
					"2":  `One hour`,
					"3":  `One night`,
					"4":  `One week`,
					"5":  `One month`,
					"6+": `Three months`,
				},
			},
			"The Haunting": {
				BaseDesc: `The vampire manipulates the sensory centers of his
				victim’s brain, flooding the victim’s senses with visions,
				sounds, scents, or feelings that aren’t really there. The
				images, regardless of the sense to which they appeal,
				are only fleeting “glimpses,” barely perceptible to the
				victim. The vampire using Dementation cannot control what the victim perceives, but may choose which
				sense is affected. The “haunting” effects occur mainly when the vitim is alone, and mostly at night. They may take the
				form of the subject’s repressed fears, guilty memories,
				or anything else that the Storyteller finds dramatically
				appropriate.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `After the vampire speaks to the victim, the
					player spends a blood point and rolls Manipulation +
					Subterfuge (difficulty of his victim’s Perception + Self-Control/Instinct). The number of successes determines
					the length of the sensory “visitations.” The precise effects are up to the Storyteller, though particularly eerie or harrowing apparitions can certainly reduce dice
					pools for a turn or two after the manifestation`,
					"1":  `One night`,
					"2":  `Two nights`,
					"3":  `One week`,
					"4":  `One month`,
					"5":  `Three months`,
					"6+": `One year`,
				},
			},
			"Eyes of Chaos": {
				BaseDesc: `This peculiar power allows the vampire to take advantage of the fleeting clarity hidden in insanity. She
				may scrutinize the “patterns” of a person’s soul, the
				convolutions of a vampire’s inner nature, or even random events in nature itself. The Kindred with this
				power can discern the most well-hidden psychoses, or
				gain insight into a person’s true self. Malkavians with
				this power often have (or claim to have) knowledge of
				the moves and countermoves of the great Jyhad, or the
				patterns of fate. Secrets revealed
				via Eyes of Chaos are never simple facts; they’re tanta-
				lizing symbols adrift in a sea of madness`,
				Lvl: 3,
				System: map[string]string{
					"sys": `This power allows a vampire to determine a
					person’s true Nature, among other things. The vampire
					concentrates for a turn, then her player rolls Perception
					+ Occult. The difficulty depends on the intricacy of the
					pattern. Discerning the Nature of a stranger would be
					difficulty 9, a casual acquaintance would be an 8, and
					an established ally a 6. The vampire could also read
					the message locked in a coded missive (difficulty 7), or
					even see the doings of an invisible hand in such events
					as the pattern of falling leaves (difficulty 6). Almost
					anything might contain some hidden insight, no mat-
					ter how trivial or meaningless. The patterns are pres-
					ent in most things, but are often so intricate they can
					keep a vampire spellbound for hours while she tries to
					understand their message.`,
				},
			},
			"Voice of Madness": {
				BaseDesc: `By merely addressing his victims aloud, the Kindred
				can drive targets into fits of blind rage or fear, forcing
				them to abandon reason and higher thought. Victims
				are plagued by hallucinations of their subconscious demons, and try to flee or destroy their hidden shames.
				Tragedy almost always follows in the wake of this pow-
				er’s use, though offending Malkavians often claim that
				they were merely encouraging people to act “according
				to their natures.” Unfortunately for the vampire concerned, he runs a very real risk of falling prey to his
				own voice’s power.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The player spends a blood point and makes
					a Manipulation + Empathy roll (difficulty 7). One tar-
					get is affected per success, although all potential vic-
					tims must be listening to the vampire’s voice.
					Affected victims fly immediately into frenzy or a
					blind fear like Rötschreck. Kindred or other creatures
					capable of frenzy, such as Lupines, may make a frenzy
					check or Rötschreck test (Storyteller’s choice as to how
					they are affected) at +2 difficulty to resist the power.
					Mortals are automatically affected and don’t remember
					their actions while berserk. The frenzy or fear lasts for
					a scene, though vampires and Lupines may test as usual
					to snap out of it. The vampire using Voice of Madness must also test
					for frenzy or Rötschreck upon invoking this power,
					though his difficulty to resist is one lower than nor-
					mal. If the initial roll to invoke this power is a failure,
					however, the roll to resist the frenzy is one higher than
					normal. If the roll to invoke this power is a botch, the
					frenzy or Rötschreck response is automatic.`,
				},
			},
			"Total Insanity": {
				BaseDesc: `The vampire coaxes the madness from the deepest
				recesses of her target’s mind, focusing it into an overwhelming wave of insanity. This power has driven
				countless victims, vampire and mortal alike, to unfortunate ends.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The Kindred must gain her target’s undi-
					vided attention for at least one full turn to enact this
					power. The player spends a blood point and rolls Ma-
					nipulation + Intimidation (difficulty of her victim’s
					current Willpower points). If the roll is successful, the
					victim is afflicted with five derangements. The number of successes
					determines the duration.`,
					"1":  `One turn`,
					"2":  `One night`,
					"3":  `One week`,
					"4":  `One month`,
					"5+": `One year`,
				},
			},
			"Lingering Malaise": {
				BaseDesc: `While lesser Dementation powers allow a vampire to
				inflict temporary (though often long-lasting) madness
				upon a victim, elders of the Clan have developed the
				ability to infect the minds of their victims with per-
				manent maladies. Lingering Malaise causes permanent
				psychological shifts within the victim, making him, as
				one Gangrel elder remarked, “an honorary Lunatic.”`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The character speaks to his victim for at
					least a minute, describing the derangement that Lin-
					gering Malaise will inflict. The player rolls Manipula-
					tion + Empathy (difficulty equal to the victim’s current
					Willpower points); the victim resists with a Willpower
					roll (using his permanent Willpower at difficulty 8). If
					the user of Lingering Malaise scores more successes, the
					victim gains a permanent derangement chosen by the
					individual who inflicts it. Lingering Malaise may only
					be used to inflict one derangement per night on any
					given victim, though multiple attempts may be made
					until the derangement takes hold.`,
				},
			},
			"Shattered Mirror": {
				BaseDesc: `Although Dementation’s low-level effects are pri-
				marily to initiate or promote insanity rather than to
				create it spontaneously, some of its more potent mani-
				festations are not as subtle. The wielder of this fear-
				some power can transfer her own deranged mindset
				into the psyche of a hapless victim, spreading her own
				brand of insanity like a virus.”`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The vampire must establish eye contact
					with her intended victim to apply this power.
					The player then rolls Charisma + Subterfuge (difficulty
					equal to the target’s current Willpower points) resisted
					by the target’s Wits + Self-Control/Instinct (difficulty
					equal to the Dementation user’s current Willpower
					points). If the aggressor wins, the target gains all of her
					derangements and Mental Flaws for a period of time
					determined by the number of net successes the aggressor scored.`,
					"1":  `One hour`,
					"2":  `One night`,
					"3":  `One week`,
					"4":  `One month`,
					"5+": `six months`,
					"6+": `one year per success over 5`,
				},
			},
			"Restructure": {
				BaseDesc: `The elder with this fearsome power has the ability
				to twist his victims’ psyches at their most basic levels,
				warping their very beings. The subject of Restructure
				retains her memories, but her outlook on life changes
				completely, as if she has undergone a sudden epiphany
				or religious conversion. This effect goes much deeper
				than the implantation of a derangement; it actually
				performs a complete rewrite of the victim’s personality.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `As the description says, this power allows
					the vampire to change his target’s Nature to one more
					suitable to his ends. To accomplish this, the character
					must make eye contact (p. 152) with his intended vic-
					tim. The player rolls Manipulation + Subterfuge (diffi-
					culty equals the victim’s Wits + Subterfuge). If he rolls
					a number of successes equal to or greater than the tar-
					get’s Self-Control/Instinct, the target’s Nature changes
					to whatever the player using Restructure desires. This
					effect is permanent and can be undone only by another
					application of Restructure (though subtle differences
					from the character’s original Nature may still remain,
					as it is impossible for such a fundamental change to occur flawlessly). A botch on this roll changes the character’s own Nature to that of his intended victim.`,
				},
			},
			"Personal Scourge": {
				BaseDesc: `Similar to the Auspex power of Psychic Assault, this fearsome ability allows the elder to turn
				the very strength of her victim’s mind against him,
				inflicting physical harm with the power of his own
				will. Victims of this self-powered attack spontaneously
				erupt in lacerations and bruises, spraying blood in every direction and howling in agony. Those who have
				observed such an attack with Auspex note that the
				victim’s aura swirls with violent psychosis and erupts
				outward in writhing appendages — a sight that can
				make even the most jaded Tzimisce quail`,
				Lvl: 8,
				System: map[string]string{
					"sys": `The vampire must touch or establish eye
					contact with her target. The player rolls Manipulation + Empathy (difficulty equal to the target’s
					Stamina + Self-Control/Instinct) and spends two Willpower points. For a number of turns equal to the number of successes rolled, the victim rolls his own permanent Willpower as lethal damage against himself. This
					damage can be soaked with his own Humanity or Path
					of Enlightenment (difficulty 6); Fortitude does not add
					to this soak dice pool, nor does body armor. He may
					take no actions during this time other than thrashing
					and gibbering; this includes spending blood to heal.`,
				},
			},
			"Lunatic Eruption": {
				BaseDesc: `This fearsome ability is only known to have been
				applied a few times in recorded Kindred history, most
				spectacularly during the final nights of the last battle of
				Carthage. It is effectively a psychic nuclear bomb, used
				to incite every intelligent being within several miles
				(kilometers) into an orgy of bloodlust and rage. It is
				suspected that the Malkavians have used the threat of
				this power as a bargaining chip in several key negotiations`,
				Lvl: 9,
				System: map[string]string{
					"sys": `The player spends four Willpower points
					and rolls Stamina + Intimidation (difficulty 8). The
					area of effect is determined by the number of successes
					scored`,
					`1`: `One city block`,
					`2`: `An entire neighborhood`,
					`3`: `A large downtown area`,
					`4`: `Several neighborhoods`,
					`5`: `An entire metropolitan area
					(approximately 30 miles or
					45 kilometers`,
					`6+`: `An additional 10 miles or 15
					kilometers for every success past 5`,
				},
			},
		},
	},
	"dominate": {
		Name: `Dominate`,
		Description: `Dominate is one of the most dreaded of Disciplines.
		It is a vampire’s ability to influence another person’s
		thoughts and actions through her own force of will.
		Dominate requires that the vampire capture her victim’s gaze; as such, it may be used against
		only one subject at a time. Further, commands must be
		issued verbally, although simple orders may be made
		with signs — for example, a pointed finger and forceful expression to indicate “Go!” However, the subject
		won’t comply if he can’t understand the vampire, no
		matter how powerful the Kindred’s will is. Perhaps unsurprisingly, vampires to which Dominate
		comes naturally tend to be from willful, domineering
		Clans. The Giovanni, Lasombra, Tremere, and Ventrue all consider an iron will to be a boon, and are eager to impose that iron will on any who would move
		against them.`,
		Abilities: map[string]disciplineAbility{
			"Command": {
				BaseDesc: `The vampire locks eyes with the subject and speaks
				a one-word command, which the subject must be obey
				instantly. The order must be clear and straightforward:
				run, agree, fall, yawn, jump, laugh, surrender, stop,
				scream, follow. If the command is at all confusing or
				ambiguous, the subject may respond slowly or perform
				the task poorly. The subject cannot be ordered to do
				something directly harmful to herself, so a command
				like “die” is ineffective.
				The command may be included in a sentence, thereby concealing the power’s use from others. This effort at subtlety still requires the Kindred to make eye
				contact at the proper moment and stress the key word
				slightly.`,
				Lvl: 1,
				System: map[string]string{
					"sys": ` The player rolls Manipulation + Intimida-
					tion (difficulty equals the target’s current Willpower
					points). More successes force the subject to act with
					greater vigor or for a longer duration (continue run-
					ning for a number of turns, go off on a laughing jag,
					scream uncontrollably).
					Remember, too, that being commanded to against
					one’s Nature confounds the use of this power. Being
					told to “sleep!” in a dangerous situation or “attack!”
					in police custody may not have the desired effect, or
					indeed, any effect at all.`,
				},
			},
			"Mesmerize": {
				BaseDesc: `With this power, a vampire can verbally implant a
				false thought or hypnotic suggestion in the subject’s
				subconscious mind. Both Kindred and target must be
				free from distraction, since Mesmerize requires intense
				concentration and precise wording to be effective. Mesmerize allows for anything from simple, precise
				directives (handing over an item) to complex, highly
				involved ones (taking notes of someone’s habits and
				relaying that information at an appointed time). It is
				not useful for planting illusions or false memories (such
				as seeing a rabbit or believing yourself to be on fire). A
				subject can have only one suggestion implanted at any
				time.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The player rolls Manipulation + Leader-
					ship (difficulty equal to the target’s current Willpower
					points). The number of successes determines how well
					the suggestion takes hold in the victim’s subconscious.
					If the vampire scores one or two successes, the subject
					cannot be forced to do anything that seems strange to
					her (she might walk outside, but is unlikely to steal a
					car). At three or four successes, the command is effec-
					tive unless following it endangers the subject. At five
					successes or greater, the vampire can implant nearly
					any sort of command. No matter how strong the Kindred’s will, his com-
					mand cannot force the subject to harm herself directly
					or defy her innate Nature. So, while a vampire who
					scored five successes could make a 98-pound weakling
					attack a 300-pound bouncer, he could not make the
					mortal shoot herself in the head.`,
				},
			},
			"The Forgetful Mind": {
				BaseDesc: `After capturing the subject’s gaze, the vampire delves
				into the subject’s memories, stealing or re-creating
				them at his whim. The Forgetful Mind does not allow
				for telepathic contact; the Kindred operates much like
				a hypnotist, asking directed questions and drawing out
				answers from the subject.The degree of detail used has a direct bearing on how
				strongly the new memories take hold, since the vic-
				tim’s subconscious mind resists the alteration. A sim-
				plistic or incomplete false memory (“You went to the
				movies last night”) crumbles much more quickly than
				does one with more attention to detail.To restore removed memories or sense false ones
				in a subject, the character’s Dominate rating must be
				equal to or higher than that of the vampire who made
				the alteration. In that situation, the player must make
				a Wits + Empathy roll (difficulty equal to the origi-
				nal vampire’s permanent Willpower rating) and score
				more successes than his predecessor did. However, the
				Kindred cannot use The Forgetful Mind to restore his
				own memories if they were stolen in such a way.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The player states what sorts of alteration he
					wants to perform, then rolls Wits + Subterfuge (dif-
					ficulty equal to the target’s current Willpower points).
					Any success pacifies the victim for the amount of time
					it takes the vampire to perform the verbal alteration,
					provided the vampire does not act aggressively toward
					her. The table below indicates the degree of modifica-
					tion possible to the subject’s memory. If the successes
					rolled don’t allow for the extent of change the charac-
					ter desired, the Storyteller reduces the resulting impact
					on the victim’s mind.`,
					`1`: `May remove a single memory;
					lasts one day.`,
					`2`: `May remove, but not alter, memory
					permanently.`,
					`3`: `May make slight changes to memory.`,
					`4`: `May alter or remove entire scene
					from subject’s memory`,
					`5`: `May reconstruct entire periods of
					subject’s life.`,
				},
			},
			"Conditioning": {
				BaseDesc: `Through sustained manipulation, the vampire can
				make a subject more pliant to the Kindred’s will. Over
				time, the victim becomes increasingly susceptible to
				the vampire’s influence while simultaneously growing
				more resistant to the corrupting efforts of other Kin-
				dred. Gaining complete control over a subject’s mind
				is no small task, taking weeks or even months to ac-
				complish.
				Kindred often fill their retainers’ heads with subtle
				whispers and veiled urges, thereby ensuring these mor-
				tals’ loyalty. Yet vampires must pay a high price for
				the minds they ensnare. Servants Dominated in this
				way lose much of their passion and individuality. They
				follow the vampire’s orders quite literally, seldom tak-
				ing initiative or showing any imagination. In the end,
				such retainers become like automatons or the walking
				dead.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The player rolls Charisma + Leadership (dif-
						ficulty equal to the target’s current Willpower points)
						once per scene. Conditioning is an extended action, for
						which the Storyteller secretly determines the number
						of successes required. It typically requires between five
						and 10 times the subject’s Self-Control/Instinct rating. Targets with more empathic Natures may require
						a lower number of successes, while those with willful
						Natures require a higher total. Only through roleplaying may a character discern whether his subject is conditioned successfully. A target may become more tractable even before
						becoming fully conditioned. Once the vampire accumulates half the required number of successes, 
						subsequent uses of Dominate. After being conditioned,
						the target falls so far under the vampire’s influence that
						the Kindred need not make eye contact or even be
						present to retain absolute control. `,
				},
			},
			"Possession": {
				BaseDesc: `At this level of Dominate, the force of the Kindred’s
				psyche is such that it can utterly supplant the mind of
				a mortal subject. Speaking isn’t required, but the vam-
				pire must capture the victim’s gaze. During the psychic
				struggle, the contestants’ eyes are locked on one an-
				other.
				Once the Kindred overwhelms the subject’s mind,
				the vampire moves his consciousness into the victim’s
				body and controls it as easily as he uses his own. The
				mortal falls into a mental fugue while under possession.
				She is aware of events only in a distorted, dreamlike
				fashion. In turn, the vampire’s mind focuses entirely
				on controlling his mortal subject. His own body lies
				in a torpid state, defenseless against any actions made
				toward it. The character may travel as far from his body as he is
				physically able while possessing the mortal. The vam-
				pire may also venture out during the day in the mortal
				form. However, the vampire’s own body must be awake
				to do so, requiring a successful roll to remain awake
				`,
				Lvl: 5,
				System: map[string]string{
					"sys": ` The vampire must completely strip away the
					target’s Willpower prior to possessing her. The player
					spends a Willpower point, then rolls Charisma + Intimidation, while the subject rolls his Willpower in a
					resisted action (difficulty 7 for each). For each success
					the vampire obtains over the victim’s total, the target
					loses a point of temporary Willpower. Only if the attacker botches can the subject escape her fate, since
					this makes the target immune to any further Dominate
					attempts by that vampire for the rest of the story.`,
					`1`: `Cannot use Disciplines`,
					`2`: `Can use Auspex and other sensory
					powers`,
					`3`: `Can also use Presence and other
					powers of emotional manipulation`,
					`4`: `Can also use Dementation,
					Dominate, and other powers
					of mental manipulation`,
					`5`: `Can also use Chimerstry,
					Necromancy, Thaumaturgy,
					and other mystical powers`,
				},
			},
			"Chain the Psyche": {
				BaseDesc: `Not content with merely commanding their subjects,some elders apply this power to ensure obedience from
				recalcitrant victims. Chain the Psyche is a Dominate
				technique that inflicts incapacitating pain on a target
				who attempts to break the vampire’s commands.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The player spends a blood point when her
					character applies Dominate to a subject. Any attempt
					that the subject makes to act against the vampire’s
					implanted commands or to recover stolen memories
					causes intense pain. When such an attempt is made,
					the Storyteller rolls the character‘s Manipulation + Intimidation (difficulty equal to the subject’s Stamina +
					Empathy). Each success equals one turn that the vic-
					tim is unable to act, as she is wracked with agony. Each
					application of Chain the Psyche crushes a number of
					resistance attempts equal to the character’s Manipula-
					tion rating, after which the effect fades.`,
				},
			},
			"Loyalty": {
				BaseDesc: `With this power in effect, the elder’s Dominate is so
				strong that other vampires find it almost impossible to
				break with their own commands. Despite the name,
				Loyalty instills no special feelings in the victim — the
				vampire’s commands are simply implanted far more
				deeply than normal.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `Any other vampire attempting to employ
					Dominate on a subject who has been Dominated by a
					vampire with Loyalty has a +3 difficulty modifier to his
					rolls and must spend an additional Willpower point.`,
				},
			},
			"Obedience": {
				BaseDesc: `While most Kindred must employ Dominate through
				eye contact, some powerful elders may command loy-
				alty with the lightest brush of a hand`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The character can employ all Dominate
					powers through touch instead of eye contact (although
					eye contact still works). Skin contact is necessary —
					simply touching the target’s clothing or something she
					is holding will not suffice. The touch does not have to
					be maintained for the full time it takes to issue a Dominate command, though repeated attempts to Dominate
					a single target require the character to touch the subject again.`,
				},
			},
			"Mass Manipulation": {
				BaseDesc: `A truly skilled elder may command small crowds
				through the use of this power. By manipulating the
				strongest minds within a given group, a gathering may
				be directed to the vampire’s will.
				`,
				Lvl: 7,
				System: map[string]string{
					"sys": `The player declares that he is using this
					power before rolling for the use of another Dominate
					power. The difficulty of the roll is that which would be
					required to Dominate the most resistant member of the
					target group — if he cannot be Dominated, no one in
					his immediate vicinity can. For every success past that
					needed to inflict the desired result on the first target,
					the player may choose one additional target to receive
					the same effect in its entirety. The vampire needs to
					make eye contact only with the initial target.`,
				},
			},
			"Still the Mortal Flesh": {
				BaseDesc: `Despite its name, this power may be employed on
				vampires as well as mortals, and it has left more than
				one unfortunate victim writhing in agony — or unable
				to do even that. A vampire who has developed this
				power is able to override her victim’s body as easily as
				his mind in order to cut off his senses or even stop his
				heart. It is rumored that this power once came more
				easily to the Kindred, but modern medicine has made
				the bodies and spirits of mortals more resistant to such
				manipulations`,
				Lvl: 8,
				System: map[string]string{
					"sys": `The player rolls Manipulation + Medicine
					(difficulty equal to the target’s current Willpower
					points + 2; a difficulty over 10 means that this power
					cannot affect the target at all). The effect lasts for one
					turn per success. The player must choose what func-
					tion of the target’s body is being cut off before rolling.
					She may affect any of the body’s involuntary functions;
					breathing, circulation, perspiration, sight, and hearing
					are all viable targets. While Still the Mortal Flesh is in effect, a vampire can either stop any one of those functions entirely or cause them to fluctuate erratically.`,
				},
			},
			"Far Mastery": {
				BaseDesc: `This refinement of Obedience (though the character
					need not have learned Obedience first) allows the use
					of Dominate on any subject that the vampire is familiar
					with, at any time, over any distance. If the elder knows
					where his target is, he may issue commands as if he
					were standing face- to-face with his intended victim.`,
				Lvl: 8,
				System: map[string]string{
					"sys": `The player spends a Willpower point and
					rolls Perception + Empathy (difficulty equal to the sub-
					ject’s Wits + Stealth) to establish contact. If this roll
					succeeds, Dominate may be used as if the character had
					established eye contact with the target. A second Will-
					power point must be spent in order for a vampire to use
					this power on another vampire or other supernatural
					being.`,
				},
			},
			"Speak Through the Blood": {
				BaseDesc: `The power structures of Methuselahs extend across
				continents and centuries. This power allows such an-
				cients to wield control over their descendants, even
				those far outside their geographic spheres of influence.
				Speak Through the Blood allows an elder to issue com-
				mands to every vampire whose lineage can be traced
				to her — even if the two have never met. Thus, en-
				tire broods act to further the goals of sleeping ancients
				whose existences they may be completely unaware of.
				The vampires affected by this power rarely act directly
				to pursue the command they were given, but over a
				decade or so, their priorities slowly shift until the fulfillment of the Methuselah‘s command is among their
				long-term goals. Speak Through the Blood, because it
				takes effect so slowly, is rarely recognized as an outside
				influence, and its victims rationalize their behavior as
				“growing and changing,” or something to that effect. A vampire who has reached Golconda is not affected
				by this power, and is completely unaware that it has
				been used. Her childer, however, are affected normally
				unless they are also enlightened. Ghouls of the victims
				of this power are also affected, but to a lesser extent.`,
				Lvl: 9,
				System: map[string]string{
					"sys": `The player spends a permanent Willpower
					point and rolls Manipulation + Leadership. The dif-
					ficulty of this roll is equal to four plus the number of
					Generations to which the command is to be passed.
					Unless the character is aware of the location and pres-
					ent agenda of every descendant of his — a highly un-
					likely event — he may only issue general commands, 158 CHAPTER FOUR: DISCIPLINES
					such as “work for the greater glory of Clan Malkavian”
					or “destroy all those who seek to extinguish the light
					of knowledge.” Speak Through the Blood can be used
					by a vampire in torpor. Commands issued through this
					power last for one decade per success on the roll. Dif-
					ficulties over 10 require one additional success for each
					point past 10, making it that much more difficult to
					issue long-lasting commands stretching down to the
					ends of one’s lineage.`,
				},
			},
			"Resist Dominate": {
				BaseDesc: `This refinement of Obedience (though the character
					need not have learned Obedience first) allows the use
					of Dominate on any subject that the vampire is familiar
					with, at any time, over any distance. If the elder knows
					where his target is, he may issue commands as if he
					were standing face- to-face with his intended victim.`,
				Lvl: 0,
				System: map[string]string{
					"Mortals": `Few mortals can hope to resist
					Dominate, as their strength of will nothing
					compared to the supernatural magnetism
					of a vampire. Still, there are extremely rare
					individuals who, due to strong religious
					faith, unique psychic talent, or extraordinary
					mental resolve, can shrug off this Discipline’s
					effects. Only a foolish vampire ignores the
					potential threat such human beings represent.`,
					"Vampires": `It is impossible to Dominate
					another Kindred who is of stronger Blood.
					The vampire must be of an equal or higher
					Generation than the target for the powers to
					be effective. Scholars of the Kindred condi-
					tion suspect that this is one of the protections
					Caine put in place to protect himself from
					the whims of his willful childer. A faction of
					those who believe this theory also maintain
					that this implies that Caine himself employed
					the Discipline of Dominate.`,
					"Nature": `A character’s Nature can have a
					distinct impact on how easily Dominate influences her. A vampire might easily control
					subjects with inherently empathic Natures (Caregiver, Child, Conformist), while those
					whose Natures denote a great degree of inner
					strength (Bravo, Director, Rebel) can be more
					of a challenge. `,
					"Botches": `If a Dominate roll botches, the
					target is rendered immune to future attempts
					by the same vampire for the rest of the story.`,
				},
			},
		},
	},
	"fortitude": {
		Name: "Fortitiude",
		Description: `Although all vampires have an unnatural constitution
		that make them much sturdier than mortals, Fortitude
		bestows a resilience that would make an action movie
		hero envious. Vampires with this Discipline can shrug
		off agonizing trauma and make the most bone-shattering impact look like a flesh wound. The power even offers protection against the traditional banes of vampires,
		such as sunlight and fire, and the Gangrel, Ravnos, and
		Ventrue all find that edge incredibly useful.`,
		Abilities: map[string]disciplineAbility{
			"Fortitude 1-5": {
				BaseDesc: ``,
				Lvl:      5,
				System: map[string]string{
					"sys": `A character’s rating in Fortitude adds to his
					Stamina for the purposes of soaking normal damage
					(bashing and lethal). A character with this Discipline
					may also use his dots in Fortitude to soak aggravated
					damage, though Kindred cannot normally soak things
					like vampire bites, werewolf claws, magical effects, fire,
					sunlight, or massive physical trauma.`,
				},
			},
			"Fortitude 6+": {
				BaseDesc: `Elder vampires progress in Fortitude in the same way
				as Celerity. They can increase their basic
				mastery of the Discipline or to take an alternate power`,
				Lvl: 6,
				System: map[string]string{
					"sys": ``,
				},
			},
			"Personal Armor": {
				BaseDesc: `Nobody likes to get hit, not even Cainites. The
				easiest way to ensure that one is not hit (or shot, or
				stabbed) repeatedly is to take the weapon with which
				one is assaulted away from one’s attacker and break it.
				That’s where Personal Armor comes in. This application of Fortitude, derived from one popular in the 12th
				century, causes anything that strikes a Kindred who
				employs Personal Armor to shatter on impact.
				`,
				Lvl: 6,
				System: map[string]string{
					"sys": `With the expenditure of two blood points,
					a vampire can add preternatural hardness to his flesh.
					Every time an attack is made on the Kindred using Per-
					sonal Armor (one which he fails to dodge), his player
					rolls Fortitude (difficulty 8). If the roll grants more suc-
					cesses than the attacker rolled, then the weapon used
					to make the attack shatters against the vampire’s flesh.
					(“Magical” weapons may be resistant to this effect, at
					the Storyteller’s discretion.) The vampire still takes
					normal damage if the attack is successful, even if the
					weapon shatters in the process, though this damage
					may be soaked. If the attack roll botches, any normal
					weapon automatically shatters. A hand-to-hand attack causes the attacker equal
					damage to that suffered by the defender when Personal
					Armor comes into play. If the attacker misses entirely,
					she still takes one level of bashing damage.
					The effects of this power last for the duration of the
					scene.`,
				},
			},
			"Shared Strength": {
				BaseDesc: `It’s one thing to laugh off bullets, rather another to
				watch the ricochets mow down everyone around you.
				Many Kindred have wished, at one time or another,
				that they could lend their monstrous vitality to those
				around them. Those few vampires who have mastered
				Shared Strength can — if only for a little while.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `Shared Strength duplicates a portion of a
					vampire’s Fortitude (one dot for every point of blood
					the vampire spends) to another being. Activating the
					power requires a Stamina + Survival roll (difficulty 8,
					increased to 9 if the target is not a normal mortal), and
					the expenditure of a point of Willpower. Furthermore,
					the vampire must mark his target by pressing a drop of
					his blood onto the target’s forehead. This stain remains
					visible as long as the power is in effect, the duration of
					which is determined by the initial roll. The target of this power need not be willing to accept the benefit to receive it, and the bestowing vampire can end the effect at any time for no cost. Particularly sadistic Kindred have come up with any number
					of ways in which a target’s “devil’s mark” and supernatural endurance can be used to land him in a great
					deal of trouble.
					A vampire can never bestow more levels of Fortitude
					than he himself possesses.`,
					"1": `One turn`,
					"2": `One scene`,
					"3": `One hour`,
					"4": `One night`,
					"5": `One week`,
					"6": `One month`,
					"7": `One year`,
				},
			},
			"Adamantine": {
				BaseDesc: `Adamantine functions as a more potent version of
				Personal Armor.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `This power mimics the effects of Personal
					Armor, save that the vampire who uses it takes no
					damage from attacks that shatter on her skin.`,
				},
			},
		},
	},
	"necromancy": {
		Name: "Necromancy",
		Description: `Necromancy is both a Discipline and a school of
		blood magic devoted to the command of the souls of the
		dead. It’s similar to Thaumaturgy in that it has several
		“paths” and accompanying “rituals” rather than a strict
		linear progression of powers. The study of Necromancy
		is not widespread among the Kindred, and its practitioners — primarily the Giovanni — are shunned and
		despised for their foul practices (until those practices
		become useful, of course).`,
		Abilities: map[string]disciplineAbility{
			"Necromancy Base": {
				BaseDesc: `Over the centuries, the various schools of vampiric
				Necromancy have evolved and diversified from an earlier form of death magic, leaving several distinct paths
				of necromantic magic available to Cainites. Nearly all
				modern necromancers learn the Sepulchre Path first
				before extending their studies to other paths. The primary Necromancy path increases automatically as the
				character increases her overall Necromancy rating.
				Other paths must be bought separately, using the experience costs for secondary paths`,
				Lvl: 1,
				System: map[string]string{
					"sys": `A Cainite necromancer must learn at least
					three levels in his primary path before learning his first
					level in a secondary Necromancy path. He must then
					master the primary path (all five levels) before acquiring any knowledge of a third path.`,
				},
			},
			"The Sepulchre Path": {
				BaseDesc: `Path in which the vampire can witness, summon, and command the spirits of the dead.`,
				Lvl:      0,
			},
			"The Ash Path": {
				BaseDesc: `Path allows necromancers to peer into the
				lands of the dead, and even affect things there.`,
				Lvl: 0,
			},
			"The Bone Path": {
				BaseDesc: `Path of corpses and ressurection`,
				Lvl:      0,
			},
			"The Cenopath": {
				BaseDesc: `Path for discovering or forging links between
				the living world and the Shadowland`,
				Lvl: 0,
			},
			"The Corpse in the Monster": {
				BaseDesc: `Path that allows the user to fully experience the corpse as a gateway between life and death`,
				Lvl:      0,
			},
			"The Grave's Decay": {
				BaseDesc: `Under this path, a practitioner of
				Necromancy channels the force of decay`,
				Lvl: 1,
			},
			"Path of the Four Humors": {
				BaseDesc: `This antiquated path tapps into the four humors, qualities split along two axes`,
				Lvl:      0,
			},
			"Vitreous Path": {
				BaseDesc: `Path allows a necromancer to control
				and influence the energies pertaining to death.`,
				Lvl: 0,
			},
			"Necromantic Rituals": {
				BaseDesc: `Rituals connected with Necromancy`,
				Lvl:      0,
			},
		},
	},
	"theSepulchrePath": {
		Name: "The Sepulchre Path",
		Description: `Through the Sepulchre path, the vampire can wit-
		ness, summon, and command the spirits of the dead.
		At higher levels, the necromancer can force the ghost
		to remain in a particular place or object, or even dam-
		age wraiths directly. Since many other areas of Nec-
		romancy involve dealing with ghosts, this is the most
		common path for necromancers to start with. If a Kindred uses a Sepulchre Path power in the presence of something of great importance to the
		ghost the power affects, the chances for success in the
		summoning increase dramatically (reduce the difficulty
		by 2).`,
		Abilities: map[string]disciplineAbility{
			"Witness of Death": {
				BaseDesc: `Before it is possible to control the dead, one must
				perceive them. This power allows just that, attuning
				a vampire’s unliving senses to the presence of the incorporeal Under its effects, a necromancer sees ghosts as translucent phantoms gliding among the living and hears
				their whispers and moans. She feels the spectral cold
				of their touch and smells their musty hint of decay. Yet
				one cannot mistake the dead for the living, as they lack
				true substance, and appear dimmer and less real than
				creatures of flesh and blood. When a vampire uses this
				power, her eyes flicker with pale blue fire that only the
				dead can see.
				Ghosts resent being spied upon, and more powerful
				shades may use their own powers to inflict their displeasure on the incautious.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The player rolls Perception + Awareness
					(difficulty 5). Success allows the vampire to perceive
					ghosts as described for the rest of the scene. Failure has no special
					effect, but a botch means the vampire can see only the
					dead for the scene; everything else appears as shapeless,
					dim shadows. While the vampire’s other senses remain
					attuned to the living, he is all but blind in this state
					and suffers a +3 difficulty to most vision-based Perception rolls and attacks. Ghosts notice the glowing eyes
					of a vampire using this power only with a successful
					Perception + Alertness roll (difficulty 7).`,
				},
			},
			"Summon Soul": {
				BaseDesc: `The power of Summon Soul allows a necromancer to
				call a ghost back from the Underworld, for conversational purposes only. In order to perform this feat (and
				indeed, most of the feats in this path), the vampire
				must meet certain conditions: The necromancer must know the name of the
				wraith in question, though an image of the wraith ob-
				tained via Witness of Death, Shroudsight, Auspex, or other supernatural perception
				will suffice.
				An object with which the wraith had some contact
				in life must be in the vicinity, though it need not be
				something of significant importance to the ghost’s liv-
				ing consciousness. A piece of the ghost’s corpse works
				well for this purpose (and even provides a -1 difficulty
				modifier).
				Certain types of ghosts cannot be summoned with
				this power. Vampires who achieved Golconda before
				their Final Deaths, or who were diablerized, are beyond
				the reach of this summons. Likewise, many ghosts of
				the dead cannot be called — they are destroyed, un-
				able to return to the mortal plane, or lost in the eternal
				storm of the Underworld.`,
				Lvl: 2,
				System: map[string]string{
					"sys": ` The player spends one blood point and rolls
					Manipulation + Occult (difficulty equal to 7 or the
					ghost’s Willpower, whichever is higher). The vampire
					must know the name of the ghost and have on hand an
					object the ghost had contact with in life. Provided that
					the target has died and become a ghost, success means
					the shade appears before the necromancer as described
					above. Not everyone becomes a ghost — it requires a
					strong will to persevere in the face of death, and souls
					that have found peace pass on to their eternal rewards.
					Moreover, it is possible for the dead to suffer spiritual
					dissolution and destruction after they become ghosts.
					Vampires know if their summons should have suc-
					ceeded by a feeling of sudden, terrifying descent as they
					reach too far into the great Beyond, so this power can
					be used to determine whether a soul has endured beyond death. While a failure means the vampire wastes
					blood, a botch calls a spirit other than the one sought
					— usually a malevolent ghost known as a Spectre. Such a fiend torments the one who summoned
					it with every wicked power at its disposal.
					Once a ghost is summoned, it may not deliberately
					move out of sight of the vampire, though it can take
					any other actions, including direct attack. The vampire’s player may spend a Willpower point to dismiss
					the ghost at any time (unless he rolled a botch). Otherwise, at the end of the scene, shadows engulf the spirit
					once more and return it to its original location.`,
				},
			},
			"Compel Soul": {
				BaseDesc: `With this power, a vampire can command a ghost to
				do his bidding for a while. Compulsion of the soul is a
				perilous undertaking and, when used improperly, can
				endanger vampire and wraith alike.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The vampire locates and approaches the intended ghost or calls it to his presence with Summon
					Soul. As with the previous power, he must have the
					ghost’s name and an object it handled in life. His player
					then spends one blood point and rolls Manipulation +
					Occult in a resisted roll against the ghost’s Willpower
					(difficulty 6 for both rolls).
					If the vampire wins, the number of net successes determines the degree of control he has over the ghost
					(as described below). Moreover, the vampire’s control
					keeps ghosts that have been called with Summon Soul
					from returning to their original locations at the end of
					the scene. If the ghost wins, the vampire loses a number.
					of Willpower points equal to the ghost’s net successes.
					On a tie, the roll becomes an extended contest that
					continues each turn until one side wins. If the vampire
					botches at any point, the ghost is immune to any use
					of the vampire’s Necromancy for the rest of the scene.
					If the ghost botches, it must obey as if the vampire’s
					player had rolled five net successes.`,
					"1": `The ghost must perform one simple
					task for the vampire that does not
					place it in certain danger. It must
					attend to this task immediately,
					although it can delay the
					compulsion and pursue its own
					business at a cost of one Willpower
					point per scene. The ghost may not
					attack the vampire until this task is
					complete. It is possible to issue the
					task of answering one question, in
					which case the ghost must answer
					truthfully and to the best of
					its knowledge`,
					"2": `The vampire may issue two orders
					or ask two questions as outlined for
					one success. Alternatively, the
					vampire may demand a simple task
					with a real possibility of danger, as
					long as the danger is not certain. The
					ghost may delay this compulsion with
					Willpower.`,
					"3": `The vampire may issue three orders
					as outlined for one success.
					Alternatively, he may demand the
					ghost fulfill one difficult and
					dangerous task or a simple assignment
					that has an extended duration of up
					to one month. The ghost may delay
					such orders with Willpower.`,
					"4": `The vampire may issue four orders,
					as outlined for one success, or
					assign two tasks, as for two successes.
					Alternatively, the vampire may
					command the ghost to perform one
					complex assignment that puts the
					ghost at extreme risk, or perform any
					number of non-threatening tasks
					as the vampire’s slave for up to one
					month (or, if the necromancer spends
					a permanent point of Willpower, for a year and a day). It is possible for
					ghosts to delay individual tasks, but
					not put off enslavement.`,
					"5": `The vampire may issue multiple
					orders that have a sum complexity
					or danger of five successes’ worth.
					Instead, the vampire may order the
					ghost to perform any one action that
					it is capable of executing within one
					month. Such a task can place the
					ghost in immediate peril of
					destruction, or even force it to betray
					and assault loved ones. It is not
					possible for ghosts to delay a task of
					this magnitude with Willpower —
					they must obey.`,
				},
			},
			"Haunting": {
				BaseDesc: `Haunting binds a summoned ghost to a particular
				location or, in extreme cases, an object. The wraith
				cannot leave the area to which the necromancer binds
				it without risking destruction.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The player spends one blood point while
					standing at the location for the haunting or touching
					the intended prison. She then has the ghost brought to
					her by whatever means she desires, though Summon
					Soul is quickest and most reliable. Her player then
					rolls Manipulation + Occult (difficulty is equal to the
					target’s current Willpower points if resisted, to a minimum of 4; otherwise it is 4). The difficulty rises by one
					if the vampire wishes to place the ghost in an object.
					As usual, the difficulty decreases by one if the necromancer has a part of the spirit’s corpse in addition to
					knowing its name (minimum difficulty 3). Each success binds the ghost within the location or
					object for one night. This duration extends to one week
					if the player spends a Willpower point or a year and a
					day for a dot of permanent Willpower. A wraith attempting to leave the area of a haunting must make an
					extended Willpower roll (difficulty 9, four cumulative
					successes necessary in a single scene) or take a level of
					aggravated damage for each roll. If the wraith runs out
					of health levels, it is hurled deep into the Underworld
					to face destruction.`,
				},
			},
			"Torment": {
				BaseDesc: `It is through the use of this power that powerful necromancers convince bound ghosts to behave — or else. Torment allows the vampire to strike a wraith as if he
				himself were in the lands of the dead, inflicting dam-
				age on the wraith’s ectoplasmic form. The vampire
				remains in the real world, however, so he cannot be
				struck in return.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The player rolls Stamina + Empathy (dif-
						ficulty equal to the wraith’s current Willpower points),
						and the vampire reaches out to strike the wraith. Each
						success inflicts a level of lethal damage on the wraith.
						Should the wraith lose all health levels, it immediately
						vanishes into what appears to be a doorway to some
						hideous nightmare realm. Ghosts “destroyed” thus cannot reappear in or near the real world for a month.`,
				},
			},
		},
	},
	"theAshPath": {
		Name: "The Ash Path",
		Description: `The Ash Path allows necromancers to peer into the
		lands of the dead, and even affect things there. Of the
		paths of Necromancy, the Ash Path is the most peril-
		ous to learn, because many of the path’s uses increase a
		necromancer’s vulnerability to wraiths`,
		Abilities: map[string]disciplineAbility{
			"Shroudsight": {
				BaseDesc: `Shroudsight allows a necromancer to see through the
				Shroud, the mystical barrier that separates the living
				world from the Underworld. By using this power, the
				vampire can spot ghostly buildings and items, the landscape of the so-called Shadowlands, and even wraiths
				themselves. However, an observant wraith may notice
				when a vampire suddenly starts staring at him, which
				can lead to unpleasant consequences.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `A simple roll of Perception + Awareness
					(difficulty 7) allows a necromancer to utilize Shroudsight. The effects last for a scene.`,
				},
			},
			"Lifeless Tongues": {
				BaseDesc: `Where Shroudsight allows a necromancer to see
				ghosts, Lifeless Tongues allows her to converse with
				them effortlessly. Once Lifeless Tongues is employed,
				the vampire can carry on a conversation with the deni-
				zens of the ghostly Underworld without spending blood
				or causing the wraiths to expend any effort.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `To use Lifeless Tongues requires a roll of
					Perception + Occult (difficulty 6) and the expenditure
					of a Willpower point.`,
				},
			},
			"Dead Hand": {
				BaseDesc: `Similar to the Sepulchre Path power Torment, Dead
				Hand allows a necromancer to reach across the Shroud
				and affect a ghostly object as if it were in the real world.
				Ghosts are solid to necromancers using this power, and can be attacked. Furthermore, the necromancer
				can pick up ghostly items, scale ghostly architecture
				(giving real-world bystanders the impression that he’s
				climbing on air!), and generally exist in two worlds.
				On the other hand, a necromancer using Dead Hand is
				quite solid to the residents of the Underworld — and
				to whatever hostilities they might have.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The player spends a point of Willpower and
					makes a successful Wits + Occult roll (difficulty 7) to
					activate Dead Hand for one scene. For each additional
					scene the vampire wishes to remain in contact with
					the Underworld, he must spend a point of blood.`,
				},
			},
			"Ex Nihilo": {
				BaseDesc: `Ex Nihilo allows a necromancer to enter the Underworld physically. While in the lands of the dead,
				the vampire is essentially a particularly solid ghost. He
				maintains his normal number of health levels, but can
				be hurt only by things that inflict aggravated damage
				on ghosts (weapons forged from souls, certain ghostly
				powers, etc.). A vampire physically in the Underworld
				can pass through solid objects in the real world (at the
				cost of one health level) and remain “incorporeal” for
				a number of turns equal to her Stamina rating. On the
				other hand, vampires present in the Underworld are
				subject to all of the Underworld’s perils, including ultimate destruction. A vampire killed in the realm of the
				dead is gone forever, beyond even the reach of other
				necromancers.`,
				Lvl: 5,
				System: map[string]string{
					"sys": ` Using Ex Nihilo takes a tremendous toll on
					the necromancer. To activate this power, the vampire
					must first draw a doorway with chalk or blood on any
					available surface. (The vampire may draw doors ahead
					of time for exactly this purpose.) The player must then
					expend two points of Willpower and two points of
					blood before making a Stamina + Occult roll (difficulty 8) as the vampire attempts to open the chalk door
					physically. If the roll succeeds, the door opens and the
					vampire steps through into the Underworld.
					When the vampire wishes to return to the real world,
					he merely needs to concentrate (and the player spends
					another Willpower point and rolls Stamina + Occult,
					difficulty 6). `,
				},
			},
			"Shroud Mastery": {
				BaseDesc: `Shroud Mastery offers the Kindred the ability to
				manipulate the veil between the worlds of the living
				and the dead. By doing so, a necromancer can make
				it easier for bound wraiths in his service to function,
				or make it nearly impossible for ghosts to contact the
				material world.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `To exercise Shroud Mastery, the necromancer expends two points of Willpower, then states
					whether he is attempting to raise or lower the Shroud.
					The player then makes a Willpower roll (difficulty 9).
					Each success on the roll raises or lowers the difficulties
					of all nearby wraiths’ attempts to cross the Shroud in
					any way by one, to a maximum of 10 or a minimum of
					3. The Shroud reverts to its normal strength at a rate of
					one point per hour thereafter.`,
				},
			},
		},
	},
	"theBonePath": {
		Name: "The Bone Path",
		Description: `The Bone Path is concerned primarily with corpses
		and the methods by which dead souls can be restored
		to the living world — temporarily or otherwise.`,
		Abilities: map[string]disciplineAbility{
			"Tremens": {
				BaseDesc: `Tremens allows a necromancer to make the flesh of a
				corpse shift once. An arm might suddenly flop forward,
				a cadaver might sit up, or dead eyes might abruptly
				open. This sort of thing tends to have an impressive
				impact on people who aren’t expecting a departed relative to roll over in his coffin.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `To use Tremens, the necromancer spends
					a single blood point, and the player must succeed on
					a Dexterity + Occult roll (difficulty 6). The more successes that are achieved, the more complicated an action can be effected in the corpse. One success allows
					for an instantaneous movement, such as a twitch, while
					five allow the vampire to set up specific conditions under which the body animates (“The next time someone
					enters the room, I want the corpse to sit up and open
					its eyes.”). Under no circumstances can Tremens cause
					a dead body to attack or cause damage.`,
				},
			},
			"Apprentice's Brooms": {
				BaseDesc: `With Apprentice’s Brooms, the necromancer can
				make a dead body rise and perform a simple function.
				For example, the corpse could be set to carrying heavy
				objects, digging, or just shambling from place to place.
				The cadavers thus animated do not attack or defend
				themselves if interfered with, but instead attempt to
				carry out their given instructions until such time as
				they’ve been rendered inanimate. Generally it takes dismemberment, flame, or something similar to destroy
				a corpse animated in this way.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `A roll of Wits + Occult (difficulty 7) and
					the expenditure of a point of both blood and Will-
					power are all that is necessary to animate corpses with
					Apprentice’s Brooms. The number of corpses animated
					is equal to the number of successes achieved. The necromancer must then state the task to which he is setting his zombies. The cadavers turn themselves to their
					work until they finish the job (at which point they collapse) or something (including time) destroys them.`,
				},
			},
			"Shambling Hordes": {
				BaseDesc: `Shambling Hordes creates obvious results: reanimated corpses with the ability to attack, albeit neither very
				well nor very quickly. Once primed by this power, the
				corpses wait — for years, if necessary — to fulfill the
				command given them. The orders might be to protect
				a certain site or simply to attack immediately, but they
				will be carried out until every last one of the decom-
				posing monsters is destroyed.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The player spends a point of Willpower.
					The player then must succeed on a Wits + Occult roll
					(difficulty 8). Each success allows the vampire to raise
					another corpse from the grave, and costs one blood
					point. If the player cannot or chooses not to pay the
					blood point cost of additional zombies past a certain
					number, the extra successes are simply lost. Each zombie can follow one simple instruction, such as “Stay
					here and guard this graveyard against any intruders,”
					or “Kill them!”`,
				},
			},
			"Soul Stealing": {
				BaseDesc: `This power affects the living, not the dead. It does,
				however, temporarily turn a living soul into a sort of
				wraith, as it allows a necromancer to strip a soul from
				a living body. A mortal exiled from his body by this
				power becomes a wraith with a single tie to the real
				world: his now-empty body.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The player spends a point of Willpower
					and then makes a contested Willpower roll against
					the intended victim (difficulty 6). Successes indicate
					the number of hours during which the original soul is
					forced out of its housing. The body itself remains autonomically alive but catatonic.
					This power can be used to create suitable hosts for
					Daemonic Possession. It has no effect on Kindred or
					other supernatural creatures (except ghouls) until such
					creatures are dead – in the case of vampires, this means
					Final Death.”`,
				},
			},
			"Daemonic Possession": {
				BaseDesc: `Daemonic Possession lets a vampire insert a soul into
				a freshly dead body. This does not turn the reanimated
				corpse into anything other than a reanimated corpse,
				one that will irrevocably decay after a week, but it does give either a wraith or a free-floating soul (say, that of
				a vampire using Psychic Projection) a temporary home
				in the physical world.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The body in question must be no more than
					30 minutes dead, and the new tenant must agree to inhabit it — a ghost or astral form cannot be forced into a
					new shell. However, most ghosts would gladly seize the
					opportunity. Should the vampire, for whatever reason,
					wish to insert a soul into another vampire’s corpse (before it crumbles to ash), the necromancer must achieve
					five successes on a resisted Willpower roll against the
					original owner of the body. Otherwise, the interloper
					is denied entrance. The soul can use whatever physical abilities
					(Athletics, Brawl, Potence) his new fleshy home pos-
					sesses, and whatever mental abilities (Computer, Law,
					Presence) he already possessed. He cannot use the
					physical abilities of his old form, or the mental abilities
					of his new one.`,
				},
			},
		},
	},
	"theCenotaphPath": {
		Name: "The Cenopath Path",
		Description: `Practitioners of the Cenotaph Path are primarily
		concerned with discovering or forging links between
		the living world and the Shadowlands. It functions
		on the principle that a Kindred, already a corpse, is
		an unnatural bridge between the living and the dead,
		and the necromancer can use this to find other, similar
		linkages. The basic rudiments of the Cenotaph Path
		function easily enough once the Kindred learns to at-
		tune himself to these connections. Advanced mastery
		of the path usually entails some brief ritual to forge ar-
		tificial connections, either through focusing unsavory
		passions or commanding this world and the Shadow-
		lands together.`,
		Abilities: map[string]disciplineAbility{
			"A Touch of Death": {
				BaseDesc: `Just as a necromancer may exert mastery over the
				Shadowlands, so too can some ghosts exert themselves in the mortal world. Whereas obvious displays
				of ghostly power such as bleeding walls or disembodied
				moans certainly won’t be mistaken, some ghostly abilities exert subtle effects that aren’t easily recognized.
				A necromancer sensitized to the residue of the dead,
				though, can feel whether an object has been touched
				by a ghost or sense the recent passage of a wraith.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The necromancer simply touches a person or
					object that he suspects is a victim of ghostly influence.
					The player rolls Perception + Awareness (difficulty 6).
					If successful, the necromancer can determine whether a ghost has exerted any sort of power on the subject, or
					even crossed nearby, to the duration detailed below.
					On a failure, the necromancer receives no impressions. A botch reveals a misleading answer (an object
					may seem tinged with ghostly power when it’s not, or
					vice versa). Should the necromancer succeed in detec-
					tion while touching an object or person that a ghost
					is possessing, he immediately becomes aware that the
					ghost is still inside. The impression gained in such a
					case is sufficient to count as an image of the spirit for
					purposes of the Sepulchre Path’s powers, so the Kindred may be able to (for example) immediately command a ghost to exit a person whom it possesses.
					`,
					"1": `Last turn; detect use of ghostly powers`,
					"2": `Last three turns; detect use of ghostly
					powers`,
					"3": `Last hour; detect ghost’s touch and
					use of ghostly powers`,
					"4": `Last day; detect ghost’s touch and use
					of ghostly power`,
					"5": `Last week; detect nearby passage of
					ghost, ghost’s touch, and use of
					ghostly powers`,
				},
			},
			"Reveal the Catene": {
				BaseDesc: `Necromantic compulsions function much more effectively when the caster uses an object of significance
				to the ghost in question. Such fetters tie the dead to
				the living lands through their remembered importance
				— a favored recliner for relaxing, a reviled piece of art
				foisted off by hated relatives, or some object of similarly
				intense emotion. Many necromancers can detect such
				catene through the use of rituals. With this power, though, the
				necromancer can determine a fetter with just a few moments of handling. The Kindred simply runs his hands
				over the object and concentrates on it. He quickly
				receives an impression of the item’s (or person’s) importance to wraiths, if any; should the wraith be one
				known to the necromancer, he immediately recognizes
				the object as a fetter to that (or those) ghost(s). Successful identification of a connected ghost is not exclusive; that is, if the vampire determines that the object
				is important to a given wraith, he can also determine if
				there are other ghosts tied to the item, though he must
				use the power again to gain their identities.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The necromancer holds and examines the
					object for at least three turns — if it’s an item, this
					means turning it over in his hands, running his fin-
					gers along it, or otherwise giving it a critical eye; with
					a person, this may require a more… invasive… ex-
					amination. The player then spends a blood point and
					rolls Perception + Occult (difficulty 7). If successful,
					the Kindred determines whether the object holds any
					significance to any ghost and, with three or more successes, the identity of at least one such ghost (which
					allows the Kindred to use the Sepulchre Path on that
					wraith, if desired). If the necromancer already knows
					any of the ghosts involved, their ties are revealed with
					their identity — so, if the necromancer already knows
					a wraith well enough to summon and compel it with
					other powers, successful identification of a fetter tells
					whether the object is tied to that ghost, in addition to
					any other impressions gained. If a botch is scored, the necromancer can never successfully use this power on the item being examined.`,
				},
			},
			"Tread Upon the Grave": {
				BaseDesc: `The extended awareness granted with the Cenotaph
				Path allows the necromancer to find locations where
				the Shadowlands and the living world come close.
				Often, the necromancer experiences a chill or shiver
				when stepping into an area where the Underworld lies
				near the living one. With practice, the vampire can
				tell exactly where such locations are.
				Experienced necromancers learn that certain locations are susceptible to ghostly influence; these haunted areas often become homes of a sort for ghosts. A
				knowledgeable vampire can thus discover places where
				the dead are likely to congregate, the better to snare
				them with other Necromancy powers.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The player simply declares intent to sense
					the Shroud in an area and makes a Willpower roll
					(difficulty 8). Success reveals whether the location is
					highly attuned to the Shadowlands, about average (not
					particularly close to the world of the dead), or far removed from the realm of death. A failing attempt at
					using the power has no adverse effect, though it may
					be attempted only once per scene (so the necromancer
					must either wait for a time or move to a different area
					before attempting Tread Upon the Grave once more). A botch stuns the necromancer into inaction for a full
					turn and costs him a temporary Willpower point, as he
					is overcome by shivers and a sense of overwhelming
					despair.
					With three or more successes, the necromancer can
					determine whether the Shroud’s strength has been ar-
					tificially altered in the area.`,
				},
			},
			"Death Knell": {
				BaseDesc: `Not all who die go on to become ghosts — many
				lack the drive to hang on after death or simply have no
				overwhelming needs that compel them to stick around.
				Normally, even necromancers have no way to sort
				those who might become ghosts from the masses who
				go on to whatever rewards await. Over time, though, a
				necromancer can become sensitized to the pull that oc-
				curs when a soul escapes from a body only to hover in
				wait, enslaved by its desires. The weight of desperation
				becomes like a tangible tug, and some necromancers
				savor this emotion even as they follow the sensation to
				find the new ghost.
				Of course, actually discovering the new ghost can be
				problematic. The Kindred may need some means to see
				through the Shroud or may have to send other wraiths
				to look for the new unfortunate, especially if a large
				accident or massacre leaves too many corpses for the
				necromancer to easily discern and test names.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `Whenever someone dies and becomes a
					ghost within a half-mile or kilometer of the necromancer, she automatically senses the demise (though many
					choose to ignore this “always-on” power unless actively
					seeking someone). This power does not automatically
					pinpoint the location of the new ghost or identify it,
					but the player may spend one Willpower point and roll
					Perception + Occult (difficulty 7) for the necromancer
					to gain a vague sense of the distance and direction to
					the new wraith. With one success, the Kindred may
					sense a vague pull in a general direction; with three
					successes, the necromancer can sense the direction and
					guess distance to within a quarter-mile or half a kilometer. With five successes, the necromancer immediately senses the location of the new ghost to within
					one foot or 30 cm. A failure carries no penalty but a
					botched attempt sends the necromancer scurrying off
					in the wrong direction.`,
				},
			},
			"Ephemeral Binding": {
				BaseDesc: `The most puissant necromancers learn not only to
				sense the ties between living and dead, but to forge
				such ties themselves. The master of Ephemeral Bind-
				ing turns an otherwise mundane object or person into a
				depository for his own necromantic energy. The undy-
				ing Curse transforms the subject into a sort of linkage
				between the living and dead. The necromancer smears
				his blood on the item in question, which mystically absorbs the vitae and, in doing so, becomes a vessel to
				anchor a spirit. A fetter created with Ephemeral Binding lasts for
				one night per success scored. The expenditure of an ad-
				ditional point of Willpower increases this duration to
				a week per success, whereas spending a permanent dot
				of Willpower extends the duration to a year and a day.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The necromancer must coat an object with
					his blood (a full blood point’s worth); if the subject is
					a person, then that individual must ingest the vitae.
					The player marks off the blood point, spends a point
					of Willpower, and rolls Manipulation + Occult (difficulty 8). If successful, the item temporarily becomes
					a fetter to one wraith. If the Kindred already knows
					the name of the wraith or has a strong psychic impression, then the object can become a fetter at any range,
					even to a ghost who normally does not come near the
					living world (so long as the ghost still exists). Otherwise, the necromancer must be able to see or sense the
					ghost (with Witness of Death, Shroudsight, or other
					such means). Botching with this power not only causes failure but
					also makes the ghost immediately aware of what the
					necromancer was trying to do. Most ghosts do not take
					kindly to meddling Kindred trying to make artificial
					chains for them.`,
				},
			},
		},
	},
	"theCorpseInTheMonster": {
		Name: "The Corpse in the Monster",
		Description: `This path enhances the necromantic understanding
		of the unliving form and allows the user to fully experience the corpse as a gateway between life and death.
		The path lets the vampire apply some of a corpse’s
		traits to a vampire, and she can enhance or reduce
		these traits at various levels of the power`,
		Abilities: map[string]disciplineAbility{
			"Masque of Death": {
				BaseDesc: `The character with this ability can assume a visage
				of death or inflict that shape on another vampire. The
				victim’s flesh becomes pallid and thin (if it is not already), and skin pulls tight against bone. This ability
				can be very useful, as it allows one to hide in plain sight
				in a tomb or crypt at any time (though the character
				remains as vulnerable to sunlight and fire as ever).
				When a necromancer uses this power on another Kindred, the victim gains the same corpselike demeanor.
				In this sense, the ability works as something of a minor
				curse. If the user inflicts Masque of Death on another vampire, he must spend a blood point, touch the target, and
				then make a Stamina + Medicine roll (difficulty equal
				to the target’s Stamina + 3). The Masque of Death lasts
				until the next sunset, unless the character who created
				the masque wishes to extinguish its effects earlier.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The player spends one blood point for the
					character to gain the form described. Those afflicted with the Masque of Death lose two points of Dexterity and Appearance (minimum of 1 in Dexterity and
					0 in Appearance) for the duration of the power. The
					player also gets two extra dice to his Intimidation dice
					pool, should he wish to terrify any onlookers. Further,
					if the character remains perfectly still, observers must
					roll five successes on a Perception + Medicine roll (difficulty 7) to distinguish the character from a normal
					corpse. The player doesn’t need to roll anything to
					have the character stop moving — vampires have no
					autonomic functions.
					`,
					"1": `Last turn; detect use of ghostly powers`,
					"2": `Last three turns; detect use of ghostly
					powers`,
					"3": `Last hour; detect ghost’s touch and
					use of ghostly powers`,
					"4": `Last day; detect ghost’s touch and use
					of ghostly power`,
					"5": `Last week; detect nearby passage of
					ghost, ghost’s touch, and use of
					ghostly powers`,
				},
			},
			"Cold of the Grave": {
				BaseDesc: `The dead feel no pain, though most undead do. With
				this ability, the character can temporarily take on the
				unfeeling semblance of the dead, in order to protect
				herself from physical and emotional harm. When as-
				suming the Cold of the Grave, the vampire’s skin be-
				comes unusually cold. When she speaks, her breath
				mists even in warm air — those with exceptional sens-
				es might even see a slight red tinge to the breath.
				The power brings a sense of lethargy over the char-
				acter, as a mortal might feel under the influence of a
				mildly unpleasant disease. It becomes difficult to rouse
				oneself to action, and very little seems important
				enough to really worry about. A corpse has no worries,
				after all.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The player spends one Willpower point.
					For the remainder of the scene, the character takes no
					wound penalties, and the player gains an additional die
					to all dice pools that involve resisting emotional manipulation, such as Intimidation or Empathy. However,
					the player also loses a die from dice pools to emotionally manipulate others. The character is a cold fish to
					those she interacts with, and they do not respond readily to her. The Cold of the Grave does not protect the
					character against the depredations of the Beast. She
					may be emotionally cold on the surface, but if others
					taunt and anger her sufficiently, she is still subject to
					frenzy as normal.`,
				},
			},
			"Curse of Life": {
				BaseDesc: `The Curse of Life inflicts some of the undesirable
				traits of the living upon the undead, removing their corpselike nature and creating a false life to remind
				them of the worst things about being alive. Targets of
				this power regain only the unpleasant aspects of life, as
				culled from the memory of the Discipline’s user. This
				may include mundane hunger and thirst, sweat and
				other excretions, the need to urinate and defecate, a
				decrease in sensory acuity, and a particular vulnerability to attacks that the character might normally shrug
				off.`,
				System: map[string]string{
					"sys": `The player spends one Willpower and rolls
					Intelligence + Medicine (difficulty 8) to affect a target
					within line of sight and no farther than 20 yards or meters from the character. If the roll succeeds, the target
					suffers the weaknesses of the living without gaining any
					benefit from that state. He does not become immune
					to sunlight or holy artifacts, for instance. However, he
					does become badly distracted by mundane needs, with
					the net result that his player suffers a +2 difficulty penalty to all rolls. He can ignore these distractions at the
					cost of one Willpower point per scene. Additionally,
					the victim cannot use blood to raise his Physical Attributes while this power is in effect, and Willpower
					cannot eliminate this penalty. The power remains in
					effect until the next sunset.`,
				},
			},
			"Gift of Corpse": {
				BaseDesc: `This power, one of the most potent on the Corpse
				in the Monster path, enables a necromancer to ignore
				most of her race’s inherent weaknesses for a short time.
				A dead body is not particularly vulnerable to sunlight,
				holy artifacts, frenzy, or being staked through the heart,
				after all, and so it is with a vampire using the Gift of
				the Corpse. As with the Cold of the Grave, above,
				the character using this power takes on an even more
				deathlike mien. It lasts for less than a minute, typical-
				ly, but that time may be enough to enable a character
				to charge through a burning building without fearing
				frenzy or instant death.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The player spends one Willpower and rolls
					Stamina + Occult (difficulty 8). For every success, the
					character can spend one turn in a state in which he is
					more akin to an animated corpse than a vampire. Holy
					artifacts and sanctified ground have no effect, and the
					character is immune to frenzy and Rötschreck. Sun-
					light does only bashing damage, and then only if bare
					skin is exposed on a clear day. Being staked through
					the heart is only as much of a danger as getting stabbed
					through his dead spleen would be. Fire harms him only
					as it would a mortal — causing lethal damage instead
					of aggravated. Should the character end the power’s duration while
					exposed to any of the aforementioned harmful things,
					he immediately takes their full effect. If he is staked, he
					become immobilized; if he is on or near fire, he begins
					to take the damage a Cainite should take, and he must
					immediately roll against Rötschreck.`,
				},
			},
			"Gift of Life": {
				BaseDesc: `With the Gift of Life, the character can experience
				the best and most positive things about being alive. The
				overwhelming hunger for blood temporarily abates, al-
				lowing the character to consume and enjoy food and
				drink. She can also enjoy sex as she wishes, and the sun
				does not burn her. The Gift of Life comes with a dark,
				terrible cost, however. Its use is almost sure to result in
				the death of a mortal, as the vampire must expend an
				enormous quantity of vitae in order to initiate it. The
				Discipline’s effects last until the midnight after the
				character uses the power, so it is in her best interests to
				use it just after midnight. After her transformation, the character gains many
				traits of an ordinary human. She is largely immune to
				the scorching effects of the sun (Fortitude difficulties
				to soak damage from direct sunlight are halved, and
				she takes no damage if she is sufficiently covered), and
				she can experience and enjoy many of the fine things
				about human life. She retains a few of her vampiric
				benefits, however. Fortitude and Auspex abilities re-
				main in place if she has either of those Disciplines, and
				the Storyteller may allow her to retain other Disciplines
				as well if he deems them dramatically appropriate. She
				also retains a vampire’s benefits when it comes to han-
				dling bashing damage. However, she is still vulnerable 
				to holy artifacts, human faith, and being staked. Her
				blood remains vitae, not human blood. Use of this ability — which creates a mockery of human life — may
				interfere with a character’s Path advancement.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The player spends 12 blood points, burning
					as much blood as possible each turn until she meets
					that level. She then rolls Stamina + Occult (difficulty
					6) and needs only one success for the power to work.
					A botch has catastrophic effects. The character might
					be instantly killed or might inadvertently Embrace her
					victim, for example. If it takes longer than one turn
					to spend the necessary blood to enact this ability, it
					does not take effect until all 12 points have been spent.
					However, the blood must be spent continuously — the
					vampire cannot burn five, run off and feed, then burn
					seven more an hour later. On the other hand, she may
					feed as she activates the power — in one turn she might
					burn one blood point while drinking three. Since few
					Kindred above the Seventh Generation can easily expend such an amount of blood, the most efficient way
					to activate this power is to have a human nearby who
					can be sacrificed to power the transformation. he vampire is no more vulnerable to fire than any
					other mortal while in this half-alive state, but she still
					suffers somewhat from the Beast. Frenzy and Rötschreck
					difficulties are halved (round up). She can remain active during the day without Humanity or Path-based
					dice pool caps, although she is certainly tired during
					the day, since that is not her usual time of activity.
					Her Beast exacts a dangerous retribution when her
					day of “life” is done. Although its influence is greatly
					suppressed during this power’s duration, the Beast has
					its way with the vampire for the next six nights, as all
					difficulties to resist frenzy increase by three. The wise
					necromancer hides herself away somewhere during
					that period, but, depending on morality and temperament, enforced isolation might drive her to frenzy on
					its own.`,
				},
			},
		},
	},
	"theGravesDecay": {
		Name: "The Grave's Decay",
		Description: `This path is derived from the observation of the
		working of time on all things mortal. Stone crumbles
		and the corpse rots away to nothing, a process of end-
		less fascination to the lost Cainites known as Cappa-
		docians. Indeed, for the undying, the process of decay
		is a fascinating disease that afflicts everyone and ev-
		erything save them. Under this path, a practitioner of
		Necromancy channels that force.`,
		Abilities: map[string]disciplineAbility{
			"Destroy the Husk": {
				BaseDesc: `Cainites who kill their victims, rather than just feed-
				ing upon them, frequently find themselves in need of a
				quick way to dispose of a corpse. While there are many
				ways to make sure that a corpse is not found — feed it
				to a pack of hounds or weigh it down and throw it in a
				river — many of these methods do involve risk to the
				vampire and are not guaranteed to succeed. Destroy
				the Husk, by contrast, is foolproof. Use of this power
				simply turns one human corpse to a pile of about 30
				pounds (13 kilograms) of unremarkable dust, roughly
				the size and shape of that body.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The player spends one blood point as the
					vampire drips her vitae onto the corpse. The player
					then rolls Intelligence + Medicine (difficulty 6). One
					success is all that is needed to render the corpse into
					dust, although the process takes a number of turns
					equal to five minus the successes.`,
				},
			},
			"Rigor Mortis": {
				BaseDesc: `One of the first changes that comes over a dead body
				is rigidity; the corpse becomes stiff as a board, frozen
				in a single pose. The Cainite who wields Rigor Mortis
				is able to push a living or undead body to that frozen
				point using only his will and understanding of the forc-
				es of decay. She forces her target to become rigid and
				unable to move without enormous effort of will, as his
				very muscles betray him.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The player spends a point of Willpower and
					rolls Intelligence + Medicine (difficulty 7). Each success freezes the target in place for one turn. A failure
					simply indicates the loss of the Willpower point, while
					a botch renders the target immune to powers in the
					Grave’s Decay path for the next 24 hours. The target
					must be visible and within about 25 yards or meters for
					this ability to take effect. A frozen target is treated as
					though he has been staked (see p. 280). With a Willpower roll (difficulty 7) and two successes, the target
					can break out of the rigor on her turn. Failure causes
					her a level of bashing damage and means another turn
					wasted and frozen.`,
				},
			},
			"Wither": {
				BaseDesc: `Reminiscent of some of the powers of Vicissitude,
				Wither allows a vampire to cripple an opponent’s limb.
				Whether the foe is living or undead, muscle shrivels
				away, skin peels, and bone becomes brittle. The target is unable to exert any noteworthy strength in the
				crippled limb. This injury lasts for far longer than most
				injuries trouble vampires, and in mortals it simply does
				not heal.
				Wither doesn’t have to be used on a limb, although
				that is its usual purpose. It can also be used simply to
				affect the target’s face and hair, making him appear far
				older than his years. It could also be applied to a target’s eye or ear, killing the sense in that organ (and
				thus requiring two uses to permanently blind or deafen). Wither cannot be used as an “instant-kill” power
				— necromancers cannot wither internal organs — but
				it can inflict a wide variety of injuries on a foe.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The player spends a Willpower point. The
					character chooses a limb on the target and then touch-
					es that limb. If the target is trying to avoid contact, the
					invoker’s player rolls Dexterity + Brawl to hit as normal. If the character succeeds in touching the intended
					limb, the target suffers two aggravated wounds. Unless
					the target soaks both wounds (such as with Fortitude),
					the struck limb is crippled and unusable until both of
					those wounds have healed. Kindred heal the wounds as they would any other aggravated wound.
					Mortals are incapable of healing aggravated wounds, so
					they suffer throughout their lives unless they are healed
					through supernatural means. A withered limb does not
					degenerate further, even on a mortal. The character
					may be crippled for life, but the limb won’t become
					infected or gangrenous. The effects of the withering depend on the affected
					limb. A crippled arm has a Strength of 0, cannot ben-
					efit from Potence, and cannot carry anything heavier
					than about half a pound (200 grams). A crippled leg
					prevents the character from moving faster than a stut-
					tering hop or dragging limp. The character suffers the
					effects of the Lame Flaw. A single withered
					eye or ear imposes a +1 difficulty to relevant Perception
					rolls. Losing both eyes or both ears imposes the effects
					of the Blind or Deaf Flaws. A
					withered tongue imposes the effects of the Mute Flaw
					, while a withered face reduces the target’s Ap-
					pearance by one for each aggravated wound suffered.`,
				},
			},
			"Corrupt the Undead Flesh": {
				BaseDesc: `Corrupt the Undead Flesh blurs the line between life
				and undeath, turning an undead creature into some-
				thing just living enough to carry and suffer from dis-
				ease. The disease inflicts the target, causing lethargy,
				dizziness, loss of strength, clumsiness, and the inabil-
				ity to keep blood in his system. This pernicious influ-
				ence is extremely virulent among mortals. They pick
				the disease up simply by spending a few hours near the
				victim. Other vampires have a harder time acquiring
				the disease. They must consume the victim’s blood to
				do so, but afterward, they suffer just as much as the
				original target — including passing the affliction on
				to others.
				The disease fades after roughly a week. Every evening at sunset, the victim has a chance to
				throw off the plague. The victim’s player rolls Stamina,
				with a difficulty equal to 10 minus the number of sunsets since acquiring the plague. On a successful roll, the
				character fights the disease to a standstill and begins
				to recover. He instantly regains his ability to manage
				blood, and he heals back one lost Attribute point per
				hour until all have returned
				`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The player chooses a target within her
					character’s line of sight and no more than 20 yards or
					meters away. She rolls Intelligence + Medicine (difficulty 6) and spends a point of Willpower. The victim’s
					player must roll Stamina (+ Fortitude, if appropriate)
					against a difficulty equal to the attacker’s Willpower. If
					the player scores more successes than the victim, he acquires a virulent disease immediately. The disease has
					the following effects: The victim’s Strength and Wits are halved (rounddown).
						• The victim loses one point of Dexterity.
						• The victim’s player must spend one additional
						blood point every evening for the vampire to rouse
						himself to consciousness. Mortals lose one health level
						per day instead.
						• The victim’s player must roll Self-Control or Instinct each time the character feeds (difficulty 8). On
						a failure, the vampire cannot keep the blood he just
						ingested inside his body, and he vomits it up in great
						horrifying gouts of gore, losing any benefit the blood
						might have provided. Humans vomit up food.`,
				},
			},
			"Dissolve the Flesh": {
				BaseDesc: `This ability brings the Grave’s Decay path full circle,
				as it causes Destroy the Husk to apply to vampires. Dis-
				solve the Flesh allows a necromancer to attempt to
				turn vampiric flesh to dust or ash, as though the target
				had been burned or left out in the sun. The undead flesh damaged by this power turns to
				dust (gone for the time being), and it must be regenerated painstakingly by the victim, should he survive.
				That dust doubtlessly has mystical properties that various sorcerers might be able to take advantage of. Every 
				wound inflicted by this ability represents the loss of
				about one-eighth of the target’s weight(It might also be
				shed from all over, leaving the victim a bit gaunter or
				missing chunks of flesh.)`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The player spends two blood points and a
					Willpower point as the vampire extracts a quantity of
					her vitae charged with the power of the grave. If she
					drips it onto a single Kindred victim anytime within
					the next few turns (most of the blood must reach the
					victim, so flinging a few drops is ineffective), it causes
					whole chunks of the victim’s body to crumble to ash.
					The player rolls Willpower against a difficulty of the
					victim’s Stamina + 3. For every success, the target
					takes one aggravated wound.`,
				},
			},
		},
	},
	"pathOfTheFourHumors": {
		Name: "Path of the Four Humors",
		Description: `Philosophically, the four humors represent different
		qualities, split along two axes: hot and cold, and wet
		and dry. Blood is hot and wet; phlegm is cold and wet;
		yellow bile is hot and dry; and black bile is cold and dry.
		Historically, when a mortal was out of sorts or ill, it was
		said that his humors were out of balance, and a philoso-
		pher or physician would try to heal him by bringing his
		humors back into balance. Ancient necromancers be-
		lieved that in their undead forms, all four humors were
		held in a mystical stasis, and that they could tap into all
		four of them instead of merely tapping into blood in the
		form of vitae as other vampires did.
		This antiquated path was primarily considered the
		knowledge of the Lamia bloodline, and certainly very
		few necromancers have learned this path without tu-
		toring from a Lamia. Since the loss of the Lamia, elder
		necromancers have searched everywhere (both in this
		world and the next) for clues to its existence.`,
		Abilities: map[string]disciplineAbility{
			"Whispers to the Soul": {
				BaseDesc: `The necromancer with this ability can let slip a little
				of her own undead bilious humor as she speaks to another being (whether mortal or Kindred). The wicked
				vapor slips into the target’s ear and whispers nightmares to the target throughout the day and night. The
				target has a harder time sleeping, and becomes irritable
				and distracted during his waking hours.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The character must whisper the target’s
					name (as she knows it) into his ear. The victim rolls
					Willpower (difficulty 8). If the roll fails, the victim suf-
					fers from nightmares and hears mad, wicked mutter-
					ings while awake, for a number of full days equal to the
					necromancer’s Manipulation. The victim loses one die
					from all dice pools while thus afflicted, and at the Sto-
					ryteller’s discretion, the difficulty to resist Rötschreck
					may be increased by one at the same time.`,
				},
			},
			"Kiss of the Dark Mother": {
				BaseDesc: `Kiss of the Dark Mother allows the necromancer
				who uses it to mix her vitae with black bile, turning it
				into a noxious poison. The necromancer forces it into
				her mouth as saliva might once have come; the vitae
				tastes acrid and bitter, as though it had been scorched.
				Once the necromancer coats her teeth and lips with it,
				she can inflict terrible damage with her bite`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The player spends one blood point; activat-
					ing this power is a reflexive action, but it must be done
					before making a bite attack. If the bite hits, the aggravated damage inflicted by a single bite is doubled before soak is calculated. This power does not affect the
					character’s ability to drain blood from the target, nor
					does it increase the amount of damage done by blood
					loss. The necromancer’s bite remains potent until this
					ability is discharged by a successful hit or she spends
					one turn cleansing the dark blood from her mouth.`,
				},
			},
			"Dark Humors": {
				BaseDesc: `The vampire can exude a coat of a particular humor
				onto her skin, causing all that touch it to experience
				the most intense form of that humor. After a necro-
				mancer has used this power, she generally feels the
				opposite of the sensation the humor usually conveys:
				Using blood leaves her depressed and pessimistic; using
				yellow bile renders her calm and placid; using black
				bile leaves her optimistic; and using phlegm makes her
				aroused and angry.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The player spends two blood points. The
					necromancer chooses which humor she wishes to ex-
					crete. The humor can simply coat the skin — in which
					case touching the victim’s skin lets the humor take ef-
					fect — or it can act as a poison if placed in a bever-
					age (or in vitae). The victim must make a Stamina roll
					(difficulty 8) to resist the effects of the humor`,
					"Phlegm": `Target becomes lethargic; all dice pools
					are reduced by two for the remainder of the scene.`,
					"Blood/vitae": `Target becomes prone to excessive
					bleeding, and any lethal or aggravated wounds he suffers deal an additional health level of damage on the
					turn after they originally occur. Vitae altered by Dark
					Humors will not turn a human into a ghoul if ingested,
					nor will it initiate a blood bond.`,
					"Black Bile": ` Target suffers a number of health levels
					of damage equal to the necromancer’s Stamina. This
					damage is considered lethal and can be soaked (if the
					victim is normally capable of soaking such damage),
					though armor does not protect against it.`,
					"Yellow Bile": `Target becomes melancholic and is
					plagued with visions of death. He cannot spend Will-
					power for the remainder of the scene, and all Willpow-
					er rolls receive a +2 difficulty.`,
				},
			},
			"Clutching the Shroud": {
				BaseDesc: `Blood, the sanguine humor, was regarded by philosophers as being both hot and wet. Blood from a cold
				corpse has been transubstantiated into a dead form —
				a cold incarnation of a hot, wet element. This transformation of the living into death holds great power;
				the necromancer knows how to infuse her own being
				with the blood of a cold corpse and transform herself
				into something not wholly vampiric. Instead, the necromancer edges closer to being an animated corpse in
				fact as well as name. She grows distant and chill, as
				though possessed by the spirit of Death itself; she has to
				work to push her attention into the physical world.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The character must drink, and then spend,
					five blood points from a cold corpse (one dead for 24
					hours or more, but generally less than three days). It
					will generally take at least two turns to consume that
					blood, and the power is not activated until the char-
					acter can spend all of it. For example, if the character
					is Twelfth Generation, Clutching the Shroud takes at
					least seven turns total to activate (two to consume the
					blood and five to spend it).
					After the power is active and for the rest of the scene,
					the necromancer gains several benefits. First, she receives two additional soak dice, which may be used to
					soak any sort of damage, even if the character does not
					possess Fortitude. Second, she gains a mystic sense of
					how far those in the area are from death — whether
					they are healthy or infirm, suffer from diseases, or are
					undead, ghouls, or mortals. Finally, a Manipulation +
					Occult roll lets her speak with ghosts freely. The difficulty for this roll depends on how attuned to death a
					locale is; a cemetery would be difficulty 5, while a cozy
					apartment might be difficulty 7. However, this ability
					makes the necromancer much more susceptible to the
					effects of powers used by ghosts, which means that she
					must act carefully.`,
				},
			},
			"Black Breath": {
				BaseDesc: `A necromancer who has mastered this path can harness the undead black bile that festers at the core of her
				being; she pulls that melancholy to her lungs and lets it
				mingle with her outgoing breath. She then exhales the
				dark mist, letting it engulf those nearby. The necro-
				mancer feels curiously lightheaded and optimistic after
				using this power, as she has forced some of her most
				depressed nature out into the world; those caught in
				the black vapors grow despairing and hopeless.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The player spends one Willpower and one
					blood point, and rolls Stamina + Athletics (difficulty
					7). Black Breath allows the character to exhale a dark
					cloud of vapor that is five yards or meters in diameter per
					success rolled. Those caught in the mists may attempt
					a Dexterity + Athletics roll to escape it if they have an
					available action; otherwise, they may be overwhelmed
					by depression to the point of suicide. Those who cannot escape the mists must immediately roll Willpower(difficulty 8 for mortals, 7 for supernatural beings) and
					achieve more successes than the invoker did. Mortals
					who fail in this actively attempt to kill themselves on
					their next turn. They do not attempt such ludicrous
					suicides as praying for a lightning bolt or holding their
					breath; they use the most effective means at hand to
					end their own lives. If prevented from suicide, they attempt it again as soon as an opportunity presents itself.
					Those who succeed on the Willpower
					roll still become enchanted with the prospect of death,
					whether mortal or Kindred, and lose two dice from all
					dice pools for the rest of the scene.
					Kindred who fail the Willpower roll do not attempt
					suicide; as they are already dead, the malign influences
					of undead humors do not have as strong an effect on
					them. Instead, the affected vampire sinks into torpor.
					The duration of this torpor is based on the vampire’s
					Humanity or Path rating, just as if lethal wounds had
					forced him into it.`,
				},
			},
		},
	},
	"vitreousPath": {
		Name: "Vitreous Path",
		Description: `The Vitreous Path allows a necromancer to control
		and influence the energies pertaining to death. This
		extremely rare path manipulates entropy, a force that
		even most necromancers are uncomfortable harness-
		ing. A development of the Nagaraja bloodline,
		the Vitreous Path makes a formidable complement to
		the necromantic craft, and those obsessed with mas-
		tery over death and souls — such as the Harbingers of
		Skulls — would certainly risk much to uncover this
		path’s secrets.
		Like most necromancers, Nagaraja generally learn
		the Sepulchre Path before any others. The Vitreous
		Path is usually their second focus of study.`,
		Abilities: map[string]disciplineAbility{
			"Eyes of the Dead": {
				BaseDesc: `The necromancer employing the Eyes of the Dead
				can see with the perceptions of the Restless Dead
				(called Deathsight). To such a manipulator of ghostly
				energies, the auras of surrounding beings give off telltale hints as to their health and even their ultimate
				fate; the necromancer can see the energies of death
				flowing through everyone, just as ghosts can. By looking at the entropic markings on a person’s body, the
				necromancer can gain rough knowledge of how far that
				person is from death, how soon that person is likely to
				die, and even what the cause of her death is likely to be. The information thus gained is not exact by any
				means, but it gives the necromancer an edge over those
				she scrutinizes.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The player rolls Perception + Occult, dif-
					ficulty 6. One success lets a necromancer determine
					whether someone is injured, diseased, or dying, as well
					as whether the individual labors under any sort of curse
					or baleful magic.
					Further, the vampire can divine the target’s eventual
					demise, depending on the successes scored. One suc-
					cess means the character can guess how long the tar-
					get has to live to within a few weeks. Three successes
					means the character can estimate how long the target
					has to live and what the probable source of death will
					be, as the entropic markings show the wounds that will
					someday exist on that person. Five successes means the
					character can actually see where and when the event
					will occur by interpreting the black marks on the tar-
					get’s soul.
					This ability lasts for one scene, though the necro-
					mancer may choose to end the power early. It can be
					used to read the fate of only one target at a time. Story-
					tellers should exercise judgment with this power, since
					the markings of death are typically unavoidable. He
					may decide to roll the dice himself, so that the player
					has no way of knowing whether her insight is correct.`,
				},
			},
			"Aura of Decay": {
				BaseDesc: `The necromancer can strengthen the feeling of en-
				tropy around her to the point where it breaks down
				nonliving objects and machines. It can gnarl wood,
				rust metal, crack silicon chips, and erode plastic, glass,
				and dead organic material. This power has a range of
				one yard or meter from the necromancer’s body, but all
				those in the presence of the vampire can feel her cor-
				ruption as an icy wind.`,
				Lvl: 2,
				System: map[string]string{
					"sys": ` No roll is required, but this power does cost
					at least one blood point. Objects subjected to this Aura
					of Decay break down and become useless after being
					targeted. How the object gives out, as well as the exact
					mechanism of failure, is up to the Storyteller. Corro-
					sion, metal fatigue, or sheer brittleness are all suitably
					likely for any given item’s demise, but the in-game ef-
					fect of using a doomed item is as if the owning charac-
					ter rolled a botch. The speed at which an item breaks
					down depends on how many blood points are spent. Note that since this power requires the expenditure
					of blood points, a character cannot cause an Aura of
					Decay while staked.`,
					`One BP`:   `One week`,
					`Two BP`:   `One day`,
					`Three BP`: `End of Scene`,
					`Four BP`:  `Five turns`,
					`Five BP`:  `One turn`,
				},
			},
			"Soul Feast": {
				BaseDesc: `Just as the necromancer can release entropic ener-
				gies from within, she may also pull them into herself
				as a source of power. Soul Feasting allows the caster
				to either draw on the ambient death energies around
				her or to actively feed on a ghost, stealing the wraith’s
				substance and mystically transforming that energy into
				sustenance.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The player spends one Willpower point to
					allow the vampire to feed on the negative energies of
					the dead. If the character is drawing the energies from
					the atmosphere, she must be in a place where death has
					occurred within the hour or in a place where death is
					common, such as a cemetery, a morgue, or the scene of
					a recent murder. Generally, the necromancer can draw
					anywhere from one to four points of entropy from such
					a location, although the difficulty in using all Necromancy and similar deathly powers within the area
					increases by an equal amount for a number of nights
					equal to the points taken. The energies of such an area
					may only be drained once until the area’s entropy replenishes. In cases when the necromancer feeds on a ghost, the
					vampire must actually attack the wraith as if feeding
					normally. Wraiths have up to 10 “blood points” that
					may be taken from them, and they become less and
					less substantial as their spirit essence drains away`,
				},
			},
			"Breath of Thanatos": {
				BaseDesc: `The Breath of Thanatos allows the necromancer to
				draw out entropic energy and focus it upon an area
				or person by taking a deep breath and then forcefully
				exhaling a fog of necromantic energy. This cloud of
				virulence is completely invisible to anyone without the
				ability to see the passing of entropy. The energy of this
				cloud is like a beacon for Spectres, and they are drawn
				to the entropic force like moths to a flame.
				Once the energy is pulled from the necromancer’s
				body, she can either disperse it over a large area as a
				lure for Spectres, or use the mist for more sinister pur-
				poses. Channeled into an object or person, the death-
				mist inflicts the subject with a debilitating, wasting ill-
				ness. Furthermore, the focused energies are tainted and
				eerie, and though generally invisible (except to powers
				such as Aura Perception), they tend to cause people
				and animals to feel uncomfortable around the victim. The player spends one blood point and rolls
				Willpower (difficulty 8). Only one success is needed to
				draw out the Breath of Thanatos. If dispersed to sum-
				mon Spectres, the energies cover roughly one-quarter
				of a mile (400 meters) in radius, centered around the
				necromancer. The range increases by an additional
				one-quarter mile or 400 meters for every additional
				blood point expended.
				Spectres summoned with this power will ignore the
				summoning necromancer for the duration of the pow-
				er unless provoked, but may well go out of their way
				to wreak havoc on anyone else in the vicinity. The
				necromancer can then use other Necromancy powers
				(such as those in the Sepulchre Path) to manipulate
				and affect these Spectres. Ghosts so targeted may then
				interact with the necromancer as normal, although the
				other Spectres in the area will continue to ignore both
				the vampire and the targeted ghost.`,
				Lvl: 4,
				System: map[string]string{
					"sys": ``,
				},
			},
			"Night Cry": {
				BaseDesc: `The breath of entropic energy becomes a scream of
				pure chaos. The necromancer can issue an unearthly cry
				(heard both in the living world and in the Shadowlands).
				The howl pours icy oblivion into a target or group of
				targets — either sweeping away the inherent entropy or
				collecting that destruction and unleashing it.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The vampire chooses a number of targets
					within one yard or meter per dot of Necromancy and
					invokes Night Cry with a terrible scream. The player
					spends a Willpower point and a blood point for each
					target beyond the first. (In other words, she spends
					no blood if only going after one target, or one blood
					for two targets. Generational blood limits apply, and
					the vampire may not “pre-spend” blood prior to using
					Night Cry.)
					The player then chooses whether the vampire will aid
					or harm the targets, and rolls Manipulation + Occult
					(difficulty 6). If she chooses to aid the target or targets,
					each success gives each affected target a -2 difficulty
					modifier to all of his actions for one turn per success. If
					she instead chooses harm, each success causes an aggravated wound to each target. Targets may be any kind of
					living creature, including supernatural ones`,
				},
			},
		},
	},
	"necromanticRituals": {
		Name: "Necromantic Rituals",
		Description: `The rituals connected with Necromancy are a hodgepodge lot. Some have direct relations to the paths, while
		others seem to have been taught by ghosts themselves,
		for whatever twisted reason. All beginning necroman-
		cers gain one Level One ritual automatically, but any
		others learned must be gained through in-game play.
		Necromantic rituals are otherwise identical to Thau-
		maturgy rituals and are learned in similar
		fashion, though the two are not at all compatible. Casting times for necromantic rituals vary
		widely; see the description for particulars. The player
		rolls Intelligence + Occult (difficulty 3 + the level of
		the ritual, maximum 9). Success indicates the ritual
		proceeds smoothly, failure produces no effect, and a
		botch indicates something has gone horribly wrong.`,
		Abilities: map[string]disciplineAbility{
			"Call of the Hungry Dead": {
				BaseDesc: `Call of the Hungry Dead takes only 10 minutes to
				cast and requires a hair from the target’s head. The rit-
				ual climaxes with the burning of that hair in the flame
				of a black candle, after which the victim becomes
				able to hear snatches of conversation from across the
				Shroud. If the target is not prepared, the voices come
				as a confusing welter of howls and unearthly demands;
				he is unable to make out anything intelligible, and may
				go briefly mad.`,
				Lvl: 1,
			},
			"Eldritch Beacon": {
				BaseDesc: `Eldritch Beacon takes 15 minutes to cast. The mate-
				rial component is a green candle, the melted wax from
				which must be collected and molded into a half-inch
				(1.5 cm) sphere. Whoever carries this sphere, whether
				in his hand or in a pocket, is highlighted in the Shad-
				owlands with a sickly-glowing green-white aura. All
				ghostly powers affect this individual with greater ease
				and severity. The sphere retains its power for one hour
				per success on the casting roll.`,
				Lvl: 1,
			},
			"Insight": {
				BaseDesc: `This ritual allows a necromancer to stare into the
				eyes of a corpse and see reflected there the last thing
				the dead man witnessed. The vision appears only in
				the eyes of the cadaver and is visible to no one ex-
				cept the necromancer using Insight.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The player rolls
					as normal as the vampire stares into the target’s eyes
					for five minutes. The number of successes on the roll
					determines the clarity of the vision. A botch shows the
					necromancer his own Final Death, which can provoke
					a Rötschreck roll.
					This power cannot be used on the corpses of vampires
					who have reached Golconda, or on bodies in which
					both eyes are missing or advanced decomposition has
					already occurred`,
					`1 sucess`: `A basic sense of the subject’s death`,
					`2 sucesses`: `A clear image of the subject’s death
					and the seconds preceding ir`,
					`3 sucesses`: `A clear image, with sound, of the
					minutes preceding death`,
					`4 sucesses`: `A clear image, with sound, of the
					half-hour before the subject’s demise`,
					`5 sucesses`: `Full sensory perception of the hour
					leading up to the target’s death`,
				},
			},
			"Knowing Stone": {
				BaseDesc: `By use of her own blood and the proper rituals, a
				necromancer can mark a person’s spirit, allowing the
				vampire to see where her subject is at any time, even
				after he has died. In this fashion many of the spirit-
				haunted vampires keep tabs on their close kin and
				their enemies.
				The necromancer cuts her skin or otherwise bleeds
				herself, and then uses the vitae to paint the name of
				the target on a consecrated stone. If the ritual is suc-
				cessful, she can afterward learn the target’s current
				whereabouts by dancing around the stone in a trance
				state until one of the spirits whispers the desired infor-
				mation into her ear. The stone loses its powers on the
				night of All Saints Day unless the vampire spends a
				blood point`,
				Lvl: 1,
			},
			"Minestra di Morte": {
				BaseDesc: `The necromancer obtains a piece of a dead body and
				simmers it in a pot with half a quart (or half a liter)
				of vampiric vitae. To this stew, the necromancer adds
				rosemary (for remembrance), basil (the funerary herb),
				and salt (the alchemic principle of clarification). After
				bringing the concoction to a full boil, the necromancer
				eats it. If the roll to activate this ritual is successful, the
				character discovers whether the subject of the grisly
				rite became a wraith or Spectre after death, or if indeed
				she became either. Unfortunately, this information can
				be learned only about the person from whose body the
				“stew meat” was taken.
				The blood component is spent progressively through
				the ritual: If the Necromancer takes the blood from
				another Kindred, she doesn’t become partially bound
				from drinking it, nor does she add a point to her blood
				pool. Similarly, if she uses her own blood, her pool decreases by a point but does not increase when she consumes the soup.
				`,
				Lvl: 1,
				System: map[string]string{
					"sys": ``,
				},
			},
			"Ritual of the Smoking Mirror": {
				BaseDesc: `This ritual allows the necromancer to use an obsidian
				mirror to see as ghosts do. By gazing into the mirror’s
				ebony depths, the vampire may discover an object’s
				flaws, assess the general health of mortals, or even read
				a being’s aura.
				At the start of the ritual, the Kindred decides which
				of the ritual’s two aspects she will use — she may not
				use both at the same time. With Lifesight, the nec-
				romancer may read auras as if she had the level two
				Auspex power Aura Perception. Deathsight, on the
				other hand, grants the necromancer the ability to see
				ghosts and the Shadowlands. It also shows the stain of
				oblivion on the living, as per Eyes of the Dead.`,
				Lvl: 1,
			},
			"Eyes of Graves": {
				BaseDesc: `This ritual, which takes two hours to cast, causes the
				target to experience intermittent visions of her death
				over the period of a week. The visions come without
				warning and can last up to a minute. The caster of the
				ritual has no idea what the visions contain, as only the
				victim sees them. Each time a vision manifests, the tar-
				get must roll Courage (difficulty 7) or be reduced to
				quivering panic. The visions, which come randomly,
				can also interfere with activities such as driving, study-
				ing, shooting, and so on.
				Eyes of the Grave requires a pinch of soil from a fresh
				grave.`,
				Lvl: 2,
			},
			"The Hand of Glory": {
				BaseDesc: `The Hand of Glory is a mummified hand used by the
				necromancer to anesthetize a home’s residents and,
				thereby, allow him free rein to do what he will in the
				residence. To create one, the necromancer wraps the
				severed hand of a condemned murderer in a shroud,
				draws it tight to squeeze out any remaining blood, and
				preserves the hand in an earthenware jar with salt, salt-
				peter, and long peppers. After a fortnight, the vampire
				removes the hand and dries it in an oven with vervain
				and fern. At the end of this process, if the roll to ac-
				tivate the ritual garners any successes, the creation is
				viable.
				`,
				Lvl: 2,
				System: map[string]string{
					"sys": `To use the Hand of Glory, the vampire first coats
					the fingertips of the mummified hand with a flammable
					substance derived from the fat of a hanged man and
					sets the fingers alight. The necromancer then recites
					the phrase, “Let all those who are asleep be asleep, and
					let those who are awake be awake.” All mortals within
					a household who are affected fall into a deep sleep and
					cannot be roused (the hand has no effect on supernatu-
					ral creatures). For each unaffected occupant of a home,
					one finger of the hand will refuse to light. Botches may
					result in all of the fingers being lit but no one in the
					home being asleep. The hand may be extinguished at
					any time by the necromancer who created it. Anyone
					else wishing to douse the hand must use milk to do so
					— nothing else works. Once made, the Hand of Glory
					may be reused indefinitely. Effects last for one scene.`,
				},
			},
			"Occhio d’Uomo Morto": {
				BaseDesc: `To cast this ritual, the necromancer needs an eye
				from a corpse whose absent soul became a ghost or
				Spectre. The eye is ritually prepared in a process in-
				volving incense, the new moon, and a period of mid-
				night chanting. The chanting climaxes when the nec-
				romancer removes one of her own eyes and replaces it
				with the one from the corpse (fresher is better). Kin-
				dred healing takes over at that point, sealing the eye
				within the socket.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `If the ritual succeeds, the Necromancer permanently
					gains the Shroudsight ability. This ability
					is always active and does not require a roll.
					Furthermore, if it was a Spectre’s corpse, the vampire
					can hear the vague murmuring of any Spectres in the
					area. This ability isn’t very precise; rather than mind
					reading, it’s more like trying to overhear a low-voiced
					conversation in the next room. With a Perception +
					Occult roll, the Necromancer can glean a very vague
					impression of what nearby Spectres are up to. Botching
					this roll may well earn the necromancer a new derange-
					ment (at the Storyteller’s discretion), as the whispers
					creep into the caster’s subconscious.
					This ritual has some major drawbacks, the first being
					that its proper result is hideously ugly. Unless the vam-
					pire wears sunglasses or finds some other way to con-
					ceal her eye, her Appearance is reduced by one dot.
					Also, dead or rotted tissue is not the best for normal
					perception. Any mundane visual Perception rolls are at
					+1 difficulty (possibly more if the corpse had bad eye-
					sight in life). On the other hand, since the eye offers
					a window into a different soul than the necromancer,
					it offers some protection against powers requiring eye
					contact. These Disciplines are used against the dead-
					eyed necromancer at +1 difficulty.
					Most importantly, however, the ghost whose body
					was desecrated knows it, and very likely hates it. The
					ghost can find the necromancer possessing his eye any-
					where, and all ghostly powers used against the necro-
					mancer by that particular ghost are at –1 difficulty`,
				},
			},
			"Puppet": {
				BaseDesc: `Used primarily to facilitate conversations with the
				recently departed, though also applied as a method of
				psychological torture, Puppet prepares a subject (will-
				ing or unwilling) as a suitable receptacle for ghostly
				possession. Over the course of one hour, the necroman-
				cer smears grave soil across the subject’s eyes, lips, and
				forehead. For the remainder of the night, any wraith
				attempting to take control of the subject gains two au-
				tomatic successes. The ritual’s effects remain even if
				the soil is washed off.`,
				Lvl: 2,
			},
			"The Ritual of Pochtli": {
				BaseDesc: `This ritual cannot be cast by itself, but only in con-
				junction with another Necromantic ritual, or with the
				heavily ritualized use of a Necromantic path. The ac-
				tion of the ritual is this: Two or more Kindred necro-
				mancers restrain a mortal vessel and inflict incisions in
				the shape of blasphemous symbols (typically subverted
				Egyptian hieroglyphs or Aztec symbols). They then
				drink from these injuries. Each participating Necro-
				mancer must make his own cut and drink from no oth-
				er cut. Thereafter, the Necromantic power the Kindred
				seek to employ gains the benefit of all the participants’
				knowledge. This ritual makes it possible for Necroman-
				cers to create truly fearsome feats of death magic.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The player rolls to activate this ritual as normal. If
					the roll succeeds, the Kindred who have participated
					in the ritual may work together on the path or ritual
					the Ritual of Pochtli is intended to assist, and players
					share successes. Note that the primary application of
					Necromancy requires its own roll, and that successes
					(and failures) garnered by the group are pooled. All
					Kindred participating in the ritual must know the Rit-
					ual of Pochtli as well as the ritual or path power the
					group seeks to enact.
					The downside of this power is that a single player’s
					botch negates the successes of the entire group, result-
					ing in a horrific failure for all the ritual workers.`,
				},
			},
			"Two Cetimes": {
				BaseDesc: `The necromancer ceremonially “kills” a mortal, lay-
				ing him out on a pallet and putting pennies on his eyes.
				The mortal’s soul journeys to the Underworld, which
				he perceives, initially at least, as a way-station. The
				mortal can interact with the souls of the dead and trav-
				el elsewhere in the Underworld, while also retaining
				the power to speak to the vampire and describe what
				he’s experiencing. While in the Underworld, however,
				the subject’s soul cannot affect the environment. Al-
				though he may talk to other spirits, he may not physi-
				cally interact with them or their surroundings — he is
				a “ghost among ghosts,” as it were. Minions may vol-
				untarily undergo the ritual to assist necromancers, or
				the vampire may use Two Centimes to terrify unwill-
				ing victims.`,
				Lvl: 2,
			},
			"Blood Dance": {
				BaseDesc: `The Blood Dance allows a ghost to communicate
				with a living relative. Necromancers sometimes perform this ritual for people in exchange for money or
				favors.
				The vampire must dance and chant for two hours,
				calling forth the right spirit and entreating all other
				ghosts to leave the area. While dancing, the vampire
				pours colored sands and ocean salt on the ground in a
				precise pattern and then makes the link between the
				living person and the deceased. If successful, the ghost
				“appears” within the necromancer’s sand-sigil and the
				living person can communicate with her for one hour.
				Failure means the spirit could not be contacted.`,
				Lvl: 3,
			},
			"Divine Sign": {
				BaseDesc: ``,
				Lvl:      3,
				System: map[string]string{
					"sys": `Upon learning a person’s birth date, the necroman-
					cer’s player may roll to activate this ritual. If successful,
					the Kindred may use this to predict the target’s next
					course of action, allowing him to deal with it accord-
					ingly. The effect on ghosts is quite different: Instead,
					the ritual imparts upon the necromancer so intimate
					an understanding of the wraith in question that it acts
					as a connection to the ghost, making it easier to invoke
					other Necromancy effects on that spirit. For story pur-
					poses, it’s the equivalent to holding one of that wraith’s
					fetters.`,
				},
			},
			"Din of the Damned": {
				BaseDesc: `This ritual is similar to the Level One Ritual Call
				of the Hungry Dead in that it makes the
				sounds of the Underworld audible in the physical realm.
				However, Din of the Damned is an area-effecting rit-
				ual used to ward a room against eavesdropping. Over
				the course of half an hour, the necromancer draws an
				unbroken line of ash from a crematorium along the
				room’s walls (this line may pass over doorframes to al-
				low entrance and egress). For the rest of the night, any
				attempt to listen in on events inside the room, whether
				simple (such as a glass to the wall), electronic (like a
				laser microphone), or mystic (including powers such as
				Heightened Senses), requires the eavesdropper to score
				more successes in a Perception + Occult roll (difficulty
				7) than the caster of the ritual scored. Failure to beat
				this mark gives the listener an earful of ghostly wailing
				and moaning and the sound of howling winds; a botch
				deafens him for the rest of the night.`,
				Lvl: 3,
			},
			"Nightmare Drums": {
				BaseDesc: `The necromancer using this ritual sends the dead to
				haunt the dreams of an enemy, using the ghosts to drive
				an opponent slowly insane. Once the ritual is cast, the
				vampire has no control over this power, except to stop
				it from continuing. The shape of the nightmares and
				181VAMPIRE THE MASQUERADE 20th ANNIVERSARY EDITION
				the images that assault the target are not under the
				control of the necromancer; they are under the control
				of the ghosts who actually do the haunting.
				The necromancer uses his own blood and a personal
				possession of the target’s in this ritual. Once the item
				has been coated with blood, the vampire must burn the
				item, sending a ghostly icon of it to the Shadowlands
				both as an identifying badge and as a reward to the
				ghosts who agree to haunt the target. While the item
				burns, the necromancer (and assistants, if available)
				pound out a relentless beat on gigantic drums of hu-
				man skin. The drums are inaudible in this realm but
				thunderous in the home of the dead. To silence the
				deafening drums, the ghosts resignedly agree to negoti-
				ate with the necromancer. They promise to send night-
				mares to the victim for as long as the vampire demands,
				in return for a favor. Their request normally runs along
				the lines of passing a message to a living relative or ex-
				acting revenge against someone who slighted them.`,
				Lvl: 3,
			},
			"Ritual of The Unearthed Fetter": {
				BaseDesc: `This ritual requires that a necromancer have a fin-
				ger bone from the skeleton of the particular ghost he’s
				interested in. When the ritual is cast, the finger bone
				becomes attuned to something vitally important to the
				wraith, the possession of which by the necromancer
				makes the casting of Necromantic powers against that
				ghost much easier (see the Sepulchre Path for
				an example). Most necromancers take the attuned fin-
				ger bone and suspend it from a thread, allowing it to
				act as a sort of supernatural compass and following it to
				the special item in question.
				Ritual of the Unearthed Fetter takes three hours to
				cast properly. It requires both the name of the ghost
				targeted and the finger bone already mentioned, as
				well as a chip knocked off a gravestone or other marker
				(not necessarily the marker of the bone’s former own-
				er). During the course of the ritual the stone crumbles
				to dust, which is then sprinkled over the finger bone.`,
				Lvl: 3,
			},
			"Tempesta Scudo": {
				BaseDesc: `Unlike most rituals, Tempesta Scudo can be cast
				speedily. The necromancer performs a short and awk-
				ward dance that ends with her biting through her own
				lip and spitting the blood in a circle around her. All
				ghosts’ actions within the circle of blood are made at
				+2 difficulty.
				To cast this ritual successfully, the necromancer must
				spend one combat turn performing the dance. At the
				end of the turn, she makes a Dexterity + Performance
				roll against difficulty 7 (if done outside of combat, the
				difficulty is only 6). During the next combat turn, she
				bites through her own lips (taking a level of bashing
				damage) and spits (spending one blood point). Then
				the normal ritual roll is made to see whether the power
				takes effect.`,
				Lvl: 3,
			},
			"Baleful Doll": {
				BaseDesc: `A baleful doll is a powerful figure that is linked di-
				rectly to the spirit of the target. This doll must be hand-
				crafted, and is only finished when it has been painted
				with the blood of the necromancer and dressed in some
				article of clothing from the victim (which should be
				unwashed for a better connection). Once the doll has
				been cursed, the vampire can use it to cause physical
				damage to the target. If the doll is injured (often with
				pins or other items), the victim takes six dice of bash-
				ing damage. If the doll is destroyed, the target suffers
				six dice of lethal damage.
				The necromancer must craft the doll, using ritual
				chants throughout the process. This normally takes
				four to five hours. The player rolls Stamina + Crafts
				(difficulty 8) to succeed in this part of the ritual — a
				doll that does not resemble its victim is useless for the
				purposes of this ritual, though some necromancers sell
				failures as “authentic voodoo dolls” to tourists.`,
				Lvl: 4,
			},
			"Bastone Diabolico": {
				BaseDesc: `Casting this ritual is tricky because it requires the
				removal of a leg bone from a living person. The donor
				must survive the removal, at least for a little while. The
				bone is then submerged in molten lead. Once it cools,
				the thin lead coating is inscribed with various runes.
				The necromancer then uses this metal-shod bone
				to beat its donor to death while repeating a droning
				Greek chant.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `With a successful roll, this ritual produces a bastone
					diabolico or “devil stick.” The stick can be activated by
					anyone who holds it and expends a point of Willpower.
					Activation lasts for a scene, and during that time any
					ghost hit with the devil stick loses a point from its Pas-
					sion pool (see p. 385). In addition to its normal effects,
					this club does an additional die of damage when used
					against the walking dead (not vampires), and such
					damage is aggravated.
					Unfortunately for the necromancer, ghosts can sense
					that the bastone diabolico is bad news, even if they don’t
					know exactly what the thing does. They tend to stay
					away from anybody carrying one, which means that all
					rolls for such a character to use powers that summon or
					attract ghosts occur at +1 difficulty.`,
				},
			},
			"Cadaver's Touch": {
				BaseDesc: `By chanting for three hours and melting a wax doll in
				the shape of the target, the necromancer turns a mortal
				target into a corpselike ruin. As the doll loses the last
				of its form, the target becomes cold and clammy. His
				pulse becomes weak and thready, and his flesh pale and
				chalky. For all intents and purposes, he becomes a rea-
				sonable facsimile of the walking dead. This can have
				some adverse effects in social situations (+2 difficulty
				on all Social rolls). The effects of the ritual wear off
				only when the wax of the doll is permitted to solidify.
				If the wax is allowed to boil off, the spell is broken.`,
				Lvl: 4,
			},
			"Peek Past the Shroud": {
				BaseDesc: `This hour-long ritual enchants a handful of ergot
				fungi mold to act as a catalyst for second sight. By eating a pinch of the mold, a subject gains the benefits of
				Shroudsight for a number of hours equal to the
				necromancer’s Stamina score. Three doses of the enchanted ergot are created for every success on the roll.
				Ergot is normally poisonous to some degree; this ritual
				removes its toxic properties. However, a botch renders
				the ergot highly and instantaneously toxic, inflicting
				eight dice of lethal damage on any subject who ingests
				it — including vampires`,
				Lvl: 4,
			},
			"Ritual of Xipe Totec": {
				BaseDesc: `To perform the ritual, the Kindred removes his vic-
				tim’s top layer of skin with an obsidian dagger, tak-
				ing care to damage the skin as little as possible in the
				process. The victim must survive this process (though
				she may well die of blood loss shortly after the ritual
				if not seen to properly). He then drains the victim’s
				blood into a large ceremonial golden bowl. There the
				blood is mixed with octli, amaranth flower, and other
				ingredients. When imbibed by the necromancer, this
				mixture causes him to sweat a glistening sheen of blood
				(equal to one blood point). `,
				Lvl: 4,
				System: map[string]string{
					"sys": `The Kindred then dons the
					skin of his victim, which on a successful roll absorbs
					the Kindred vitae and begins to heal, forming a second
					skin over the vampire’s own. The victim needs to be of
					similar stature — otherwise, the features become dis-
					torted and the disguise is rendered useless. This power
					also has no effect on supernatural creatures (although
					it can affect ghouls).
					Under normal visual scrutiny, the ruse is flawless. Of
					course, it imparts none of the victim’s knowledge or
					mannerisms (and does nothing to mask the Kindred’s
					own undead nature). Therefore, it works best for situations in which contact with friends and family may
					be minimized. To preserve the skin’s condition, the
					Kindred must bathe it in a blood point’s worth of vi-
					tae nightly. When the necromancer removes the skin
					(which causes one level of unsoakable lethal damage
					to the user and must be done with the same knife used
					to flay the victim in the first place), it is ruined in the
					process.`,
				},
			},
			"Chill of Oblivion": {
				BaseDesc: ``,
				Lvl:      5,
				System: map[string]string{
					"sys": ``,
				},
			},
			"Dead Man’s Hand": {
				BaseDesc: `The necromancer takes a rag stained in the blood,
				sweat, or tears of the intended victim. She takes a
				freshly severed human hand (which can come either
				from a corpse or a living “donor”) and closes it around
				the rag. As the hand decomposes, so does the victim.
				His flesh bloats, turns gray and then green, then starts
				to slough off. The victim’s brain remains fresh until the
				very end, so he can see the maggots writhe in the putrescent rack of meat that once was his healthy body.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The necromancer makes the standard roll and spends
					two blood points for each point of Stamina (and Forti-
					tude) possessed by the victim. The victim loses health
					levels according to the timetable below. Only the re-
					moval of the rag from the hand can stop the process.
					If this happens, health levels return, also according to
					the chart below.`,
					`Bruised`:       ` 12 hours`,
					`Hurt`:          `12 hours`,
					`Injured`:       `6 hours`,
					`Wounded`:       `3 hours`,
					`Mauled`:        `1 hour`,
					`Crippled`:      `30 minutes`,
					`Incapacitated`: `12 minutes`,
				},
			},
			"Esilio": {
				BaseDesc: `Like Tempesta Scudo, Esilio is a quick and dirty ritual.
				The necromancer simply speaks five syllables. No one
				can identify the casting language, but according to the
				ritual’s oblique history, the language is what God gave
				humankind before the confusion of Babel. The legend
				further states that while the particular meaning of the
				words is lost, they are what Caine’s father said to him
				while exiling him to Nod.
				Regardless of the truth of the matter, the Words of
				Exile are not spoken lightly. When the ritual is cast
				successfully, it opens a hole within reality itself — a rip
				between the lands of the living and the darkest depths
				of the Underworld. This tear is invisible to normal vi-
				sion, but to Witness of Death or Shroudsight it looks
				like a black vortex opening within the vampire’s own
				body (the very few unfortunate enough to look into the
				gap with high levels of Auspex are generally unwill-
				ing or unable to discuss it). Any ghost clutched to the
				Kindred’s chest is instantly torn to shreds. Grabbing a
				ghost in this fashion requires a Clinch or Tackle ma-
				neuver. Destroyed spirits don’t come back for at least
				184 CHAPTER FOUR: DISCIPLINES
				a month, if ever. A wraith destroyed in this fashion
				tends to return as a Spectre, if it returns at all`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The necromancer may clutch and destroy a number
					of spirits equal to the number of successes she rolled.
					After that, the vortex closes. It closes at the end of the
					scene if it hasn’t already.
					Of course, using one’s body as a portal between our
					world and what some people might call Hell is neither
					simple nor healthy. For starters, it costs a blood point
					and a point of Willpower (which does not give an au-
					tomatic success on the ritual roll). More importantly,
					each success rolled inflicts a level of unsoakable lethal
					damage on the necromancer. Most significantly, every
					use of Esilio permanently reduces the necromancer’s
					Humanity by one point if he follows that morality, and
					may impact other Paths at the Storyteller’s discretion.`,
				},
			},
			"Grasp the Ghostly": {
				BaseDesc: `Requiring a full six hours of chanting, this ritual al-
				lows a necromancer to bring an object from the Un-
				derworld into the real world. It’s not simple, however
				— a wraith may object to having his possessions stolen
				and fight back. Furthermore, the object taken must be
				replaced by a material item of roughly equal mass, oth-
				erwise the target of the ritual snaps back to its previous,
				ghostly existence.
				Objects taken from the Underworld tend to fade away
				after about a year. Only items recently destroyed in the
				real world (called “relics” by ghosts) may be recaptured
				in this manner. Artifacts created by wraiths themselves
				were never meant to exist outside the Underworld, and
				vanish on contact with the living world`,
				Lvl: 5,
			},
		},
	},
	"obfuscate": {
		Name: "Obfuscate",
		Description: `Obfuscate is the uncanny ability for Kindred to con-
		ceal themselves from sight, sometimes even in full
		view of a crowd. An Obfuscated vampire doesn’t actu-
		ally become invisible, however — rather, he is able to
		delude observers into believing that he has vanished.
		Obfuscate also allows Kindred to change their features
		and conceal other people or objects. Typically vam-
		pires using Obfuscate must be within a short range of
		their witnesses (approximately five yards or meters per
		dot of Wits + Stealth) for their power to be effective.
		Unless the Kindred chooses to make herself seen,
		she can remain obscured for as long as she wills it. At
		higher levels, the vampire can actually fade from sight
		so subtly that those nearby can’t actually recall the mo-
		ment at which she left. Since Obfuscate clouds the mind of the viewer, vam-
		pires can’t use it to hide their presence from electronic
		or mechanical devices. Video and photo cameras, for
		example, capture the vampire’s image accurately. Even
		so, the person using, say, her cell phone to record an
		Obfuscated vampire will still have her mind impacted
		by the power, and she won’t see the Kindred’s im-
		age until she views the video at a later date (if even
		then)`,
		Abilities: map[string]disciplineAbility{
			"Cloak of Shadows": {
				BaseDesc: `At this level, the vampire must rely on nearby shad-
				ows and cover to assist in hiding his presence. He steps
				into an out-of-the-way, shadowed place and eases him-
				self from normal sight. The vampire remains unnoticed
				as long as he stays silent, still, under some degree of
				cover (such as a curtain, bush, door frame, lamppost,
				or alley), and out of direct lighting. The immortal’s
				concealment vanishes if he moves, attacks, or falls un-
				der direct light. Furthermore, the vampire’s deception
				cannot stand up to concentrated observation without
				fading.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `No roll is required as long as the character
					fulfills the criteria described above. So long as he re-
					mains quiet and motionless, virtually no one but an-
					other Kindred with a high enough Auspex rating will
					see him.`,
				},
			},
			"Unseen Presence": {
				BaseDesc: `With experience, the vampire can move around
				without being seen. Shadows seem to shift to cover
				him, and people automatically avert their gazes as he
				passes by. Others move unconsciously to avoid contact
				with the cloaked creature; those with weak wills may
				even scurry away from the area in unacknowledged
				fear. The vampire remains ignored indefinitely unless
				someone deliberately seeks him out or he inadvertently
				reveals himself.
				Since the vampire fully retains his physical sub-
				stance, he must be careful to avoid contact with any-
				thing that may disclose his presence (knocking over a
				vase, bumping into someone). Even a whispered word
				or the scuffing of a shoe against the floor can be enough
				to disrupt the power.`,
				Lvl: 2,
				System: map[string]string{
					"sys": ` No roll is necessary to use this power unless
					the character speaks, attacks, or otherwise draws atten-
					tion to himself. The Storyteller should call for a Wits +
					Stealth roll under any circumstances that might cause
					the character to reveal himself. The difficulty of the
					roll depends on the situation; stepping on a squeaky
					floorboard might be a 5, while walking through a pool
					of water may require a 9. Other acts may require a certain number of successes; speaking quietly without giving away one’s position, for instance, demands at least
					three successes. Upon success, the vampire, all her
					clothing, and objects that could fit into a pocket are
					concealed.
					Some things are beyond the power of Unseen Pres-
					ence to conceal. Although the character is cloaked
					from view while he smashes through a window, yells
					out, or throws someone across the room, the vampire
					becomes visible to all in the aftermath. Bystanders snap
					out of the subtle fugue in which Obfuscate put them.
					Worse still, each viewer can make a Wits + Awareness
					roll (difficulty 7); if successful, the mental haze clears
					completely, so those individuals recall every move the
					character made up until then as if he had been visible
					the entire time.`,
				},
			},
			"Mask of a Thousand Faces": {
				BaseDesc: `The vampire can influence the perception of others,
				causing them to see a face different from his. Although
				the Kindred’s physical form does not change, any observer who cannot sense the truth sees whomever the
				vampire wishes her to see.
				The vampire must have a firm idea of the visage he
				wishes to project. The primary decision is whether to
				create an imaginary face or to superimpose the features
				of another person. Manufactured features are often
				more difficult to compose in believable proportions,
				but such a disguise is easier to maintain than having to
				impersonate someone else. Of course, things get simpler if the Kindred borrows the face but doesn’t bother
				with the personality.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The player rolls Manipulation + Perfor-
					mance (difficulty 7) to determine how well the disguise
					works. If the character tries to impersonate someone,
					he must get a good look at the subject before putting
					on the mask. The Storyteller may raise the difficulty if
					the character catches only a glimpse. The chart below
					lists the degrees of success in manufacturing another
					appearance. Vampires wishing to mask themselves as
					a person more attractive than they are must pay ad-
					ditional blood points equal to the difference between
					the vampire’s Appearance rating and the Appearance
					of the mask (which means that younger vampires may
					need to take longer in order to spend the blood neces-
					sary). Actually posing as someone else carries its own prob-
					lems. The character should know at least basic infor-
					mation about the individual; especially difficult decep-
					tions (fooling a lover or close friend) require at least
					some familiarity with the target in order to succeed.`,
					`1 sucess`: `The vampire retains the same height
					and build, with a few slight
					alterations to his basic features.
					Nosferatu can appear as normal,
					albeit ugly, mortals.`,
					`2 sucesses`: `He looks unlike himself; people don’t
					easily recognize him or agree about
					his appearance.`,
					`3 sucesses`: `He looks the way he wants to appear.`,
					`4 sucesses`: `Complete transformation, including
					gestures, mannerisms, appearance,
					and voice.`,
					`5 sucesses`: `Profound alteration (appear as the
						opposite sex, a vastly different age,
						or an extreme change of size.`,
				},
			},
			"Vanish from the Mind's Eye": {
				BaseDesc: `This potent expression of Obfuscate enables the
				vampire to disappear from plain view. So profound is
				this vanishing that the immortal can fade away even if
				he stands directly in front of someone.
				While the disappearance itself is quietly subtle, its
				impact on those who see it is anything but. Most kine
				panic and flee in the aftermath. Especially weak-willed
				individuals wipe the memory of the Kindred from their
				minds. Although vampires are not shaken so easily,
				even Kindred may be momentarily surprised by a sud-
				den vanishing.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The player rolls Charisma + Stealth; the
					difficulty equals the target’s Wits + Alertness (use the
					highest total in the group if the character disappears
					in front of a crowd). With three or fewer successes,
					the character fades but does not vanish, becoming an
					indistinct, ghostlike figure. With more than three, he
					disappears completely. If the player scores more suc-
					cesses than an observer’s Willpower rating, that person
					forgets that the vampire was there in the first place.
					Tracking the character accurately while he appears
					ghostlike requires a Perception + Alertness roll (dif-
					ficulty 8). A successful roll means the individual can
					interact normally with the vampire (although the
					Kindred looks like a profoundly disturbing ghostly
					shape). A failed roll results in a +2 difficulty modi-
					fier (maximum 10) when attempting to act upon, or
					interact with, the vampire. The Storyteller may call
					for new observation checks if the vampire moves to an
					environment in which he’s difficult to see (heads into
					shadows, crosses behind an obstacle, proceeds through
					a crowd). When fully invisible, the vampire is handled
					as described under Unseen Presence, above.
					A person subject to the vanishing makes a Wits +
					Courage roll (mortals at difficulty 9, vampires at dif-
					ficulty 5). A successful roll means the individual reacts
					immediately (although after the vampire performs his
					action for that turn); failure means the person stands
					uncomprehending for two turns while her mind tries to
					make sense of what she just experienced.`,
				},
			},
			"Cloak the Gathering": {
				BaseDesc: `At this degree of power, the vampire may extend
				his concealing abilities to cover an area. The immor-
				tal may use any Obfuscate power upon those nearby as
				well as upon himself, if he wishes.
				Any protected person who compromises the cloak
				exposes himself to view. Further, if the one who invokes the power gives himself away, the cloak falls
				from everyone. This power is particularly useful if the
				vampire needs to bring his retinue through a secure location without drawing the notice of others.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The character may conceal one extra indi-
					vidual for each dot of Stealth he possesses. He may be-
					stow any single Obfuscate power at a given time to the
					group. While the power applies to everyone under the
					character’s cloak, his player need only make a single
					roll. Each individual must follow the requirements de-
					scribed under the relevant Obfuscate power to remain
					under its effect; any person who fails to do so loses the
					cloak’s protection, but doesn’t expose the others. Only
					if the vampire himself errs does the power drop for ev-
					eryone.`,
				},
			},
			"Conceal": {
				BaseDesc: `The vampire may mask an inanimate object up to
				the size of a house (Obfuscate cannot be used to dis-
				guise inanimate objects without the use of this power).
				If the object is hidden, so are all of its contents. While
				Conceal is in effect, passersby walk around the con-
				cealed object as if it were still visible, but refuse to ac-
				knowledge that they are making any kind of detour.`,
				Lvl: 6,
				System: map[string]string{
					"sys": ` In order to activate this power, a character
					must be within about 30 feet (approximately 10 me-
					ters) of the object to be concealed and the item must
					hold some personal significance for him. The Conceal
					power functions as Unseen Presence for purposes of
					detection, as well as the duration and durability of the
					disguise.
					Conceal can be used on a vehicle in which the char-
					acter is traveling. In this instance, traffic patterns seem
					to flow around the vehicle, and accidents are actually
					less likely as other drivers subconsciously maneuver
					away from the concealed auto. A police radar gun still
					registers a speeding car masked in this fashion, but the
					officer behind the gun is disinclined to make a traffic
					stop of the phantom blip. Using Conceal on aircraft
					is problematic, as the power’s range generally doesn’t
					extend far enough to cover air traffic controllers and
					the like.`,
				},
			},
			"Mind Blank": {
				BaseDesc: `A vampire with this power is able to shrug off tele-
				pathic contact, easily withstanding invasive probes of
				her mind.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `Any attempt to read or probe the character’s
					mind first requires a successful Perception + Empathy
					roll (difficulty equal to the character’s Wits + Stealth).
					Even if a potential intruder does succeed, his dice pool
					for the attempt is then limited to the number of successes he scored on the initial roll.`,
				},
			},
			"Soul Mask": {
				BaseDesc: `In addition to concealing her form, a vampire who
				has developed Soul Mask is able to conceal her aura.
				She may display whatever combination of colors and
				shades she wishes, or may appear to have no aura what-
				soever. This power is of particular use to those of elder
				Generation who have reached such heights of power
				through diablerie.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The use of this power allows the projection of only one aura (or lack thereof) — the vampire
					chooses the precise colors to be displayed when she first
					develops Soul Mask. If the character has no experience
					with the use of Aura Perception, she may not choose
					an alternate aura, as she has no idea what one would
					look like, though she can still choose to display no aura
					whatsoever. Soul Mask can be bought multiple times,
					if desired, in order to give a vampire multiple alternate
					auras from which to choose.
					Unless the player states otherwise, Soul Mask is al-
					ways in effect. If the character has bought Soul Mask
					two or more times, her “default” aura displayed is the
					first one she learned.`,
				},
			},
			"Cache": {
				BaseDesc: `Most Obfuscate powers require the individual using
				them to be within a short distance of the subjects of
				the concealment. Cache extends this range consider-
				ably, allowing an elder with this power to leave people
				or objects safely hidden while he goes about his busi-
				ness elsewhere.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `A character must be within the normal re-
					quired distance to initiate an Obfuscate power. Once
					this is done, the player spends a Willpower point, which
					activates Cache on top of the already functioning use
					of the Discipline. The concealment will now remain
					in effect as long as the vampire is within a distance
					equal to his Wits + Stealth in miles (or one and half
					times that in kilometers) from the object or person he
					wishes to conceal. The enhanced concealment fades at
					the next sunrise, or breaks (as always) if the Obfuscate
					subject reveals himself.`,
				},
			},
			"Veil of Blissful Ignorance": {
				BaseDesc: `This power’s development is attributed to the Mal-
				kavians, but many Nosferatu have also found it to be
				highly useful. The Veil of Blissful Ignorance allows a
				vampire to Obfuscate an unwilling victim, removing
				him from the notice of others. Some Nosferatu use this
				power to teach a humbling lesson to individuals who
				take the presence and aid of others for granted, while
				others utilize it to remove an essential member of a
				group in the midst of a crisis.The victim of Veil of Blissful Ignorance does not necessarily know that he is under the effect of this power.
				He is only aware that everyone around him has sud-
				denly begun acting as if he were not there. The victim
				cannot break this effect, even with violence; if he at-
				tacks someone, the target ascribes the act to the visible
				individual nearest to him. More than one fatal brawl
				has been incited by this side-effect. The Veil persists
				even if the vampire who activated it leaves the area.
				Curiously enough, Veil of Blissful Ignorance can
				never be used on anyone who is ready and willing to
				accept its effects.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `The character must touch the victim to ac-
					tivate this power. The player spends a blood point and
					rolls Wits + Stealth (difficulty equals the victim’s Ap-
					pearance + 3). If the roll is a success, the victim is sub-
					ject to the effects of Vanish From the Mind’s Eye for a
					length of time determined by the number of successes
					the player rolls.`,
					`1 sucess`:   `Three turns`,
					`2 sucess`:   `One minutes(20 turns)`,
					`3 sucess`:   `15 minutes`,
					`4 sucess`:   `One hour`,
					`5 sucesses`: `One night`,
				},
			},
			"Old Friend": {
				BaseDesc: `Many elder Nosferatu have made reputations for om-
				niscience with the secrets they learn through creative
				uses of this power. A variation of Mask of a Thousand
				Faces, Old Friend allows a vampire to probe a subject’s
				subconscious and take the semblance of the individual
				whom that victim trusts over anyone else. Someone
				using this power does not appear as someone who the
				victim is frightened of or awed by, but rather someone
				to whom the victim feels comfortable revealing inti-
				mate secrets. Old Friend doesn’t necessarily make its
				user appear as someone who is still among the living; a
				long- dead friend or relative is just as likely, and in such
				cases the subject remembers the encounter as a dream
				or a ghostly visitation.`,
				Lvl: 8,
				System: map[string]string{
					"sys": `The player rolls Manipulation + Subterfuge
					(difficulty equal to the victim’s Perception + Alert-
					ness or Awareness, maximum 10). The more successes,
					the more convincing the impersonation. This power
					only affects one victim at a time; other observers see
					the vampire as she truly is, unless she also establishes
					a Mask of a Thousand Faces in addition to using Old
					Friend.`,
				},
			},
			"Create Name": {
				BaseDesc: `Some Toreador call this power the ultimate develop-
				ment of method acting. Create Name allows a charac-
				ter to create a completely new identity; face, speech
				pattern, aura, even thought processes are constructed
				according to the vampire’s desired identity. The power
				can be used to impersonate an existing individual, or
				it can project the semblance of a completely fictional
				identity with perfect accuracy.`,
				Lvl: 9,
				System: map[string]string{
					"sys": `A vampire working with Create Name must
					spend three hours a night in relatively uninterrupted
					quiet to establish a new personality by means of this
					power. The player makes an extended roll of Intelligence + Subterfuge (difficulty 8), one roll per night. A
					total of 20 successes are necessary to construct a new
					identity, while a botch removes five successes from the
					vampire’s total. Once a new identity has been successfully created, however, the character can step into it at
					any time without any sort of roll. Any outside observer
					without Auspex 9 or the equivalent sees the artificial
					identity. The character’s face, aura, Nature, Demeanor,
					even thoughts and Psychological Merits and Flaws all
					appear to be those of the persona selected and crafted
					by the character.
					The only way to pierce this disguise, other than Auspex 9, is to notice any discrepancies between the assumed identity and the Abilities it logically ought to
					possess. A character with no dots in Medicine should
					have a hard time pulling off a created identity as a neurosurgeon, for example. The Storyteller should make a
					secret roll of Perception + Alertness (difficulty 9) for
					each character who should catch a slip made by the
					impostor.`,
				},
			},
		},
	},
	"obtenebration": {
		Name: "Obtenebration",
		Description: `The signature power of the Lasombra, Obtenebra-
		tion grants the vampire power over darkness itself. The
		nature of the darkness invoked by Obtenebration is
		a matter of intense debate among Kindred. Some be-
		lieve it to be merely shadows, while others feel that the
		power gives control over the stuff of the vampire’s soul,
		coaxing it tangibly outward.
		Regardless, the effects of Obtenebration are terrify-
		ing, as waves of darkness roil out from the Cainite, en-
		veloping those in their path like an infernal wave. As
		Obtenebration is mostly known as a Sabbat Discipline,
		any Camarilla vampire caught using the power had
		better have a damned good explanation.  Vampires using Obtenebration can see through
		the darkness they control, though other vampires (even
		those that also have Obtenebration) cannot. `,
		Abilities: map[string]disciplineAbility{
			"Shadow Play": {
				BaseDesc: `This power grants the vampire limited control over
				shadows and other ambient darkness. Though the vam-
				pire cannot truly “create” darkness, she can overlap and
				stretch existing shadows, creating patches of gloom.
				This power also allows Kindred to separate shadows
				from their casting bodies and even shape darkness into
				the shadows of things that are not there.
				Once a Kindred takes control of darkness or shadow,
				it gains a mystical tangibility. By varying accounts cold
				or hellishly hot and cloying, the darkness may be used
				to aggravate or even smother victims. Certain callous
				Lasombra claim to have choked mortals to death with
				their own shadows.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `This power requires no roll, but a blood
					point must be spent to activate it. Shadow Play lasts for
					one scene and requires no active concentration. Kindred cloaking themselves in shadow gain an extra die
					in their Stealth dice pools and add one to the difficulties of ranged weapon attacks against them. Vampires
					who use the darkness to make themselves more terrifying add one die to Intimidation dice pools. Opponents
					plagued by flapping shadows and strangling darkness
					subtract one die from all Stamina dice pools (including soak). Mortals, ghouls, and other air-breathers reduced to zero Stamina by strangling shadows begin to
					asphyxiate; vampires lose all appropriate dice but are
					otherwise unaffected. Only one target or subject may
					be affected by this power at any given time, though
					some modicum of concealment is offered to a relatively
					motionless group.
					The unnatural appearance of this power proves extremely disconcerting to mortals and animals (and, at
					the Storyteller’s discretion, Kindred who have never
					seen it before). Whenever this power is invoked within
					a mortal’s vicinity, that individual must make a Courage roll (difficulty 8) or suffer a one-die penalty to all
					dice pools for the remainder of the scene, due to fear of
					the monstrous shadows.`,
				},
			},
			"Shroud of Night": {
				BaseDesc: `The vampire can create a cloud of inky blackness.
				The cloud completely obscures light and even sound
				to some extent. Those who have been trapped within
				it (and survived) describe the cloud as viscous and un-
				nerving. This physical manifestation lends credence to
				those Lasombra who claim that their darkness is some-
				thing other than mere shadow.
				The tenebrous cloud may even move, if the creating
				Kindred wishes, though this requires complete concen-
				tration.`,
				Lvl: 2,
				System: map[string]string{
					"sys": ` The player rolls Manipulation + Occult
					(difficulty 7). Success on the roll generates darkness
					roughly 10 feet (three meters) in diameter, though
					the amorphous cloud constantly shifts and undulates,
					sometimes even extending shadowy tendrils. Each
					additional success doubles the diameter of the cloud
					(though the vampire may voluntarily reduce the area
					she wishes to cover). The cloud may be invoked at
					a distance of up to 50 yards/meters, though creating
					darkness outside the vampire’s line of sight adds two
					to the difficulty of the roll and requires a blood point’s
					expenditure.
					The tarry mass actually extinguishes light sources it
					engulfs (with the exception of fire), and muffles sounds
					until they are indistinguishable. Those within the
					cloud lose all sense of sight and feel as though they’ve
					been immersed in pitch. Sound also warps and distorts
					within the cloud, making it nearly impossible to accomplish anything (+2 difficulty, as per Blind Fighting). Even those possessed of Heightened
					Senses, Eyes of the Beast, Tongue of the Asp, and similar powers suffer the penalty for blindness due to the
					unnatural darkness. Additionally, being surrounded by
					the Shroud of Night reduces Stamina-based dice pools
					by two dice, as the murk smothers and agitates the victims. This effect is not cumulative with Shadow Play,
					although targets asphyxiate as per Shadow Play if they
					reach 0 Stamina; more than one unfortunate mortal
					has “drowned” in darkness.
					Mortals and animals surrounded by the Shroud
					of Night must make Courage rolls per Shadow Play,
					above, or panic and flee.`,
				},
			},
			"Arms of the Abyss": {
				BaseDesc: `Refining his control over darkness, the Kindred can
				create prehensile tentacles that emerge from patches of
				dim lighting. These tentacles may grasp, restrain, and
				constrict foes.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The player spends a blood point and makes
					a simple (never extended) Manipulation + Occult roll
					(difficulty 7); each success enables the creation of a
					single tentacle. Each tentacle is six feet (two meters)
					long and possesses Strength and Dexterity ratings equal
					to the invoking vampire’s Obtenebration Trait — Po-
					tence and Celerity dots are added to these Strength and
					Dexterity ratings, respectively. If the vampire chooses,
					she may spend a blood point either to increase a single
					tentacle’s Strength or Dexterity by one or to extend its
					length by another six feet or two meters. Each tentacle
					has four health levels, is affected by fire and sunlight
					as if it were a vampire, and soaks bashing and lethal
					damage using the vampire’s Stamina + Fortitude. Ag-
					gravated damage may not be soaked.
					Tentacles may constrict foes, inflicting (Strength +1)
					lethal damage per turn. Breaking the grasp of a ten-
					tacle requires the victim to win a resisted Strength roll
					against the tentacle (difficulty 6 for each). However,
					tentacles cannot be used for any kind of manipulation,
					such as typing or driving.`,
				},
			},
			"Black Metamorphosis": {
				BaseDesc: `The Cainite calls upon his inner darkness and in-
				fuses himself with it, becoming a monstrous hybrid of
				matter and shadow. His body becomes mottled with
				spots of tenebrous shade, and wispy tentacles extrude
				from his torso and abdomen. Though still humanoid,
				the vampire takes on an almost demonic appearance,
				as the darkness within him bubbles to the surface.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The player spends two blood points and
					makes a Manipulation + Courage roll (difficulty 7) —
					vampires of lower Generation may have to take two
					turns to make the transition. Failure indicates the vam-
					pire cannot undergo the Black Metamorphosis (though
					he spends the blood points nonetheless). A botch in-
					flicts two unsoakable health levels of lethal damage on
					the vampire as darkness ravages his undead body.While under the effects of the Black Metamorphosis,
					the vampire possesses four tentacles similar to those
					evoked via Arms of the Abyss (though their Strength
					and Dexterity ratings are equal to the vampire’s own
					Attributes, including dice from Celerity and Potence).
					These tentacles, combined with the bands of darkness
					all over the Kindred’s body, subtract two dice from the
					Stamina and soak dice pools of opponents physically
					touched in combat, for as long as the vampire remains
					in contact with the victim. This is not cumulative
					with other powers in Obtenebration, although targets
					can asphyxiate at Stamina 0, as per Shadow Play. The
					vampire may make an additional attack without pen-
					alty by using the tentacles (for a total of two attacks,
					not one additional attack per tentacle). Additionally,
					the vampire can sense his surroundings fully even in
					pitch darkness.
					The vampire’s head and extremities sometimes ap-
					pear to fade away into nothingness, while at other
					times they seem swathed in otherworldly darkness.
					This, combined with the wriggling tentacles writhing
					from his body, creates an unsettling sight. Mortals, ani-
					mals, and other creatures not accustomed to this sort
					of display must make Courage rolls (difficulty 8) or suc-
					cumb to a panic that amounts to Rötschreck (though
					it is inspired by the darkness rather than fire). Many
					Kindred cultivate this devilish aspect, and the Black
					Metamorphosis adds three dice to the invoking Kin-
					dred’s Intimidation dice pools.`,
				},
			},
			"Tenebrous Form": {
				BaseDesc: `At this level, the Kindred’s mastery of darkness is so
				extensive that she may physically become it. Upon ac-
				tivation of this power, the vampire becomes an inky,
				amoeboid patch of shadow. Vampires in this form are
				practically invulnerable and may slither through cracks
				and crevices. In addition, the shadow-vampire gains
				the ability to see in natural darkness.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The transformation costs three blood points
					(which may need to occur over three turns, depending on the vampire’s Generation). The vampire is immune to physical attacks while in the tenebrous form
					(though she still takes aggravated damage from fire and
					sunlight), but may not herself physically attack. She
					may, however, envelop and ooze over others, affecting them in the same manner as a Shroud of Night,
					in addition to using mental Disciplines. Vampires in
					Tenebrous Form may even slither up walls and across
					ceilings or “drip” darkness upward — they have no
					mass and are thus unaffected by gravity. Rötschreck
					difficulties from fire and sunlight do increase by one for
					vampires in this form, as the light is even more painful
					to their shadowy bodies.
					Mortals (and others not used to such displays) who
					witness the vampire transform into unholy shadow re-
					quire Courage rolls (difficulty 8) in order to avoid the de-
					bilitating terror described under Black Metamorphosis.`,
				},
			},
			"The Darkness Within": {
				BaseDesc: `This power allows the Cainite to call forth the dark-
				ness contained in her black soul. This enormous, turbu-
				lent shadow vomits from the vampire’s mouth, though
				some vampires are said to cut themselves and let the
				blackness seep from their veins. The shadow-cloud en-
				gulfs a chosen target, burning it with a soul-scarring
				chill and siphoning its blood away in torrents.`,
				Lvl: 6,
				System: map[string]string{
					"sys": ` The player makes a Willpower roll (difficulty 6) and spends a blood point. The resulting shadow
						envelops the target and, though it does not physically
						harm the victim, it may strike terror into him. Individuals observing the Darkness Within, whether as targets or onlookers, may suffer from the terror described
						under Black Metamorphosis, unless they are already
						familiar with the Kindred’s power.
						Individuals touched by The Darkness Within lose one
						point of blood per turn, though targets may resist this effect by succeeding on a Stamina roll (difficulty 6) each
						turn the target remains in contact with the cloud.
						The Cainite invoking The Darkness Within must
						devote all her attention to maintaining the cloud. If
						the vampire is attacked, the darkness immediately returns to her through whatever orifice it originated. The
						Cainite can summon the darkness back at any time,
						gaining a number of blood points equal to one-half the
						number the shadow siphoned from its victims (round
						up). Taking blood from another in this fashion is similar to drinking from that vampire, and blood bonds
						may result. Additionally, the Darkness Within may
						take blood from only one individual per turn, though it
						may be in contact with many.`,
				},
			},
			"Shadowstep": {
				BaseDesc: `The vampire has such fine control over the darkness
				that he may become it briefly and reform himself from
				other darkness close by. The vampire may Shadowstep
				through walls, floors, and even mystical barriers. The
				Cainite simply steps “into” a shadow and re-emerges
				from another shadow a short distance away (or next to
				the barrier, if there is no shadow on the other side).`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The player rolls Dexterity + Occult, and on
					a successful roll, the character may emerge from anoth-
					er shadow no more than 50 feet (or 15 meters) away.
					Failing the roll means simply that the character cannot
					step through the shadow-realm, while a botch signifies the character has become trapped between shadows (which fiendish Storytellers should have a heyday
					with). Pulling another individual through the shadow
					requires a Strength + Occult roll, with consequences
					for failure similar to failing by oneself.`,
				},
			},
			"Shadow Twin": {
				BaseDesc: `The vampire’s control over darkness has progressed
				to such a degree that he may bestow upon it a lim-
				ited degree of sentience. By animating his own shadow
				or that of another, the Cainite can actually “free” the
				shadow cast by light. While this power is active, the
				subject casts no shadow, as it has left to pursue the
				vampire’s commands.
				This power can unnerve mortals and even a few inex-
				perienced vampires. The Kindred wielding Obtenebra-
				tion commands the individual’s shadow; some vampires
				report having seen mortals literally scared to death, as
				their shadows leapt away to taunt or menace them.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `The player spends a blood point and makes
					a Willpower roll (difficulty 8). If the roll succeeds, the
					shadow springs to unholy freedom for one hour per
					success (though it disappears at sunrise regardless of
					how many successes the vampire had). The Shadow
					Twin has Attribute and Ability ratings equal to half
					those of its parent body; they won’t do much talking
					or thinking, so Mental and Social Traits don’t matter
					much, though Wits may come into play. Additionally,
					the Shadow Twin has an Obtenebration score equal
					to one-half of that of the vampire who animated it
					(rounded down). Mortals and vampires unused to Ob-
					tenebration require a Courage roll upon witnessing
					this, as per Shadow Play.
					The twin may separate itself from the parent and
					travel up to 50 feet or 15 meters away, crawling
					through crevices or sliding up walls. It may attack and
					be attacked, though it takes and does only half dam-
					age (again, round down). Flame and supernatural at-
					tacks (such as vampire fangs, mystical powers, etc.) do
					full damage. If the Shadow Twin is killed, its parent
					loses half her Willpower points and must roll to avoid
					Rötschreck (difficulty 9).`,
				},
			},
			"Oubliette": {
				BaseDesc: `By creating a “chamber” of pure darkness, the Cain-
				ite may entrap or smother her enemies. No air exists
				in this shadow-trap, and mortals suffocate within its
				chilling void. Even vampires have little recourse once
				trapped — they may leave only at their captor’s whim.
				The Oubliette appears as a dense patch of shadow, un-
				affected by ambient light around it.`,
				Lvl: 8,
				System: map[string]string{
					"sys": `The vampire spends a blood point, but no
					roll is necessary to create the Oubliette. To actually
					create the Oubliette around someone requires a contested Wits + Larceny roll against the target’s Dexterity
					+ Occult (difficulty 7 for both rolls). Mortals suffocate
					within a number of minutes equal to their Stamina
					(though the Lasombra may choose to leave their heads
					exposed or trap a quantity of air inside as well), while
					vampires are simply suspended impotently in darkness
					and may not use Disciplines or take other actions. The
					Oubliette vanishes instantly when touched by sunlight
					— which has left more than one vampire exposed to the
					sun’s unforgiving rays — or when the Kindred chooses
					to relax it. A vampire may maintain only one Oubliette
					at a time (which can only contain one target at a time),
					which leads some Cainite philosophers to argue that it
					is a prison created from the vampire’s very soul.`,
				},
			},
			"Ahriman's Demense": {
				BaseDesc: `This power allows the vampire to summon a dark-
				ness so oppressive that it extinguishes the light of life
				— or unlife — of any victim trapped within it. Ahri-
				man’s Demesne creates a 50-foot (or 15-meter) radius
				of void that issues from the Cainite’s hand and takes
				away the bodies of those it claims when it vanishes.
				The overwhelming darkness destroys friend and foe
				alike, claiming anyone unfortunate enough to be with-
				in its circumference.`,
				Lvl: 9,
				System: map[string]string{
					"sys": ` The player spends two points of Willpower
					and concentrates for three turns. During this time, the
					blackness billows out of the character’s hand, growing
					to fill the area. At the end of the third turn, the player
					rolls Manipulation + Occult (difficulty 6). Everyone
					in the darkness’ area suffers that many health levels of
					damage (aggravated, if the victims are vampires) out-
					right — six successes yield six levels of damage, not
					six dice of damage. After Ahriman’s Demesne does its
					damage, it collapses, taking with it the bodies of any
					who died when they came in contact with the dreadful
					shadow.`,
				},
			},
		},
	},
	"potence": {
		Name: "Potence",
		Description: `This Discipline enables vampire to leap
		massive distances, lift tremendous weights, and strike
		opponents with brutal force. Even low ranks of this
		power can give Kindred physical power beyond mor-
		tal bounds. More powerful Kindred can leap so far that
		they appear to be flying, toss cars like soda cans, and
		punch through walls like cardboard. While the more
		subtle mental Disciplines can be awe-inspiring, the
		brutal effectiveness of Potence is formidable in its own
		right.`,
		Abilities: map[string]disciplineAbility{
			"Potence 1-5": {
				BaseDesc: `Kindred endowed with Potence possess unnatural strength.`,
				Lvl:      1,
				System: map[string]string{
					"sys": `: Each dot that the vampire has in Potence
					adds one die to all Strength-related dice rolls. Further,
					the player can spend one blood point and change his
					Potence dice into an equal number of automatic successes to all Strength-related rolls for the turn. In melee and brawling combat, successes from Potence (either rolled or automatic) are applied to the damage roll
					results.`,
				},
			},
			"Imprint": {
				BaseDesc: `A vampire with extensive knowledge of Potence can
				squeeze very, very hard. As a matter of fact, she can
				squeeze (or press, or push) so hard that she can leave
				an imprint of her fingers or hand in any hard surface
				up to and including solid steel. A use of Imprint can
				simply serve as a threat, or it can be used, for example,
				to dig handholds into sheer surfaces for purposes of
				climbing.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `Imprint requires a point of blood to activate.
					The power remains active for the duration of a scene.
					The depth of the imprint the vampire creates with Imprint is up to the Storyteller — decisions should take
					into account how much force the vampire can bring to
					bear, the toughness of the material, and its thickness.
					If the object the vampire grasps is thin enough, at the
					Storyteller’s discretion, the vampire might simply be
					able to push through it (in the case of a wall) or tear it
					off (in the case of a spear or pipe).`,
				},
			},
			"Earthshock": {
				BaseDesc: `According to some pundits, Potence is merely the
				art of hitting something very hard. But what do you
				do when your target is too far away to hit directly? The
				answer is, if you’re sufficiently talented with the Dis-
				cipline, to employ Earthshock. On its simplest level,
				Earthshock is the ability to hit the ground at point A,
				and subsequently have the force of the blow emerge
				from the ground at point B.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `The use of Earthshock requires the expenditure of two blood points, as well as a normal Dexterity
					+ Brawl roll. The vampire punches or stamps on the
					ground, and, if the attack is a success, the force of the
					blow emerges from the ground as a geyser of stone and
					earth directly underneath the target. The attack can be
					dodged at a +2 difficulty.
					Earthshock’s range is 10 feet or three meters for every level of Potence the vampire has, up to the limits
					of visibility. A failure on the attack roll means that
					the strike goes errant and is liable to explode anywhere
					within range; a botch means that the vampire pulverizes the ground beneath him and may well bury himself
					in the process.`,
				},
			},
			"Flick": {
				BaseDesc: `It is a truism that “the great ones always make it look
				easy.” In the case of Flick, that saying stops being a
				truism and becomes literal truth. With this power, a
				master of Potence can make the slightest gesture — a
				wave, a snap of the fingers, the toss of a ball — and
				have it unleash the full, devastating impact of a dead-
				on strike. The attack can come without warning, limiting the target’s ability to dodge or anticipate; this makes
				Flick one of the most feared applications of Potence.`,
				Lvl: 8,
				System: map[string]string{
					"sys": ` Flick costs a point of blood, and requires
					a Dexterity + Brawl roll (difficulty 6). The vampire
					must also make some sort of gesture directing the blow.
					What the gesture is remains up to the player — any-
					thing from a snap of the fingers to a blown kiss has
					worked in the past.
					Flick‘s range is equal to the limit of the Kindred’s
					perception, and the blow struck does damage equal to
					a normal punch (including all bonuses).`,
				},
			},
		},
	},
	"presence": {
		Name: "Presence",
		Description: `Presence is the Discipline of emotional manipula-
		tion. Vampires with this power can inspire passionate
		fervor or unreasoning terror in mortals and Kindred
		alike. In addition, unlike most Disciplines, some of
		Presence’s powers can be used on entire crowds at one
		time. Presence can transcend race, religion, gender,
		class, and (most importantly) supernatural nature. As
		such, this subtle power is one of the most useful Disci-
		plines a vampire can possess. Anyone can resist Presence for one scene by spend-
		ing a Willpower point and succeeding on a Willpower
		roll (difficulty 8), but the affected individual must keep
		spending points until he is no longer in the presence
		of the vampire (or, in the case of Summon, until the
		effect wears off). Vampires three or more Generations
		lower than the wielder need only spend a single Will-
		power to ignore the Presence for an entire night and
		need not roll Willpower to do so.`,
		Abilities: map[string]disciplineAbility{
			"Awe": {
				BaseDesc: `Those near the vampire suddenly desire to be closer
				to her and become receptive to her point of view. Awe
				is extremely useful for mass communication. It matters
				little what is said — the hearts of those affected lean
				toward the vampire’s opinion. The weak want to agree
				with her; even if the strong-willed resist, they soon
				find themselves outnumbered. Awe can turn a chancy
				deliberation into a certain resolution in the vampire’s
				favor almost before her opponents know that the tide
				has turned.
				Despite the intensity of this attraction, those so smitten do not lose their sense of self-preservation. Danger
				breaks the spell of fascination, as does leaving the area.
				Those subject to Awe will remember how they felt in
				the vampire’s presence, however. This will influence
				their reactions should they ever encounter her again.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The player spends a blood point and rolls
					Charisma + Performance (difficulty 7). The number of
					successes rolled determines how many people are affected, as noted on the chart below. If there are more
					people present than the character can influence, Awe
					affects those with lower Willpower ratings first. The
					power stays in effect for the remainder of the scene or
					until the character chooses to drop it.`,
					`1 success`:   `1 person`,
					`2 successes`: `2 people`,
					`3 successes`: `6 people`,
					`4 successes`: `20 people`,
					`5 successes`: `Everyone in the vampire's immediate vicinity (an entire auditorium, a mob)`,
				},
			},
			"Dread Gaze": {
				BaseDesc: `While all Kindred can frighten others by physically
				revealing their true vampiric natures — baring claws
				and fangs, glaring with malevolence, hissing loudly with
				malice — this power focuses these elements to insanely
				terrifying levels. Dread Gaze engenders unbearable ter-
				ror in its victim, stupefying him into madness, immobil-
				ity, or reckless flight. Even the most stalwart individual
				will fall back from the vampire’s horrific visage`,
				Lvl: 2,
				System: map[string]string{
					"sys": ` The player rolls Charisma + Intimidation
					(difficulty equal to the victim’s Wits + Courage). Success indicates the victim is cowed, while failure means
					the target is startled but not terrified by the sight. Three
					or more successes means he runs away in abject fear;
					victims who have nowhere to run claw at the walls,
					hoping to dig a way out rather than face the vampire.
					Moreover, each success subtracts one from the target’s
					action dice pools next turn.
					The character may attempt Dread Gaze once per turn
					against a single target, though she may also perform it
					as an extended action, adding her successes in order to
					subjugate the target completely. Once the target loses
					enough dice that he cannot perform any action, he’s so
					shaken and terrified that he curls up on the ground and
					weeps. Failure during the extended action means the
					attempt falters. The character loses all her collected
					successes and can start over next turn, while the victim
					may act normally again.
					A botch at any time indicates the target is not at all
					impressed — perhaps even finding the vampire’s antics
					comical — and remains immune to any further uses of
					Presence by the character for the rest of the story.`,
				},
			},
			" Entrancement": {
				BaseDesc: `This power bends others’ emotions, making them
				the vampire’s willing servants. Due to what these in-
				dividuals see as true and enduring devotion, they heed
				the vampire’s every desire. Since this is done willingly,
				instead of having their wills sapped, these servants re-
				tain their creativity and individuality.
				While these obedient minions are more personable
				and spirited than the mind-slaves created by Domi-
				nate, they’re also somewhat unpredictable. Further,
				since Entrancement is of a temporary duration, dealing
				with a lapsed servant can be troublesome. A wise Kin-
				dred either disposes of those she Entrances after they
				serve their usefulness, or binds them more securely by a
				blood bond (made much easier by the minion’s willing-
				ness to serve).`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The player spends a blood point and rolls
					Appearance + Empathy (difficulty equal to the target’s
					current Willpower points); the number of successes
					determines how long the subject is Entranced, as per
					the chart below. (Subjects can still spend Willpower to
					temporarily resist, like any other Presence power.) The
					Storyteller may wish to make the roll instead, since the
					character is never certain of the strength of her hold
					on the victim. The vampire may try to keep the subject
					under her thrall, but can do so only after the initial
					Entrancement wears off. Attempting this power while
					Entrancement is already in operation has no effect.`,
					`Botch`: `Subject cannot be entranced
					for the rest of the story`,
					`Failure`: `Subject cannot be entranced
					for the rest of the night`,
					`1 success`:   `1 hour`,
					`2 successes`: `1 day`,
					`3 successes`: `1 week`,
					`4 successes`: `1 month`,
					`5 successes`: `1 year`,
				},
			},
			"Summon": {
				BaseDesc: `This impressive power enables the vampire to call to
				herself any person whom she has ever met. This call
				can go to anyone, mortal or supernatural, across any
				distance within the physical world. The subject of the
				Summons comes as fast as he is able, possibly without
				even knowing why. He knows intuitively how to find
				his Summoner — even if the vampire moves to a new
				location, the subject redirects his own course as soon
				as he can. After all, he’s coming to the vampire herself,
				not to some predetermined site.
				Although this power allows the vampire to call some-
				one across a staggering distance, it is most useful when
				used locally. Even if the desired person books the next
				available flight, getting to Kyoto from Milwaukee can
				still take far longer than the vampire needs. Obviously,
				the individual’s financial resources are a factor; if he
				doesn’t have the money to travel quickly, it will take
				him a far greater time to get there.`,
				Lvl: 4,
				System: map[string]string{
					"sys": ` The player spends a blood point and rolls
					Charisma + Subterfuge. The base difficulty is 5; this in-
					creases to difficulty 7 if the subject was met only briefly.
					If the character used Presence successfully on the target
					in the past, this difficulty drops to 4, but if the attempt
					was unsuccessful, the difficulty rises to 8.
					The number of successes indicates the subject’s speed
					and attitude in responding.`,
					`Botch`: `Subject cannot be Summoned by
					that vampire for the rest of the story.`,
					`Failure`: `Subject cannot be Summoned by that
					vampire for the rest of the night`,
					`1 success`: `Subject approaches slowly
					and hesitantly.`,
					`2 successes`: `Subject approaches reluctantly and
					is easily thwarted by obstacles.`,
					`3 successes`: `Subject approaches with
					reasonable speed.`,
					`4 successes`: `Subject comes with haste,
					overcoming any obstacles in
					his way.`,
					`5 successes`: `Subject rushes to the vampire, doing
					anything to get to her.`,
				},
			},
			"Majesty": {
				BaseDesc: `At this stage, the vampire can augment her super-
				natural mien a thousandfold. The attractive become
				paralyzingly beautiful; the homely become hideously
				twisted. Majesty inspires universal respect, devo-
				tion, fear — or all those emotions at once — in those
				around the vampire. The weak scramble to obey her
				every whim, and even the most dauntless find it almost
				impossible to deny her.
				People affected find the vampire so formidable that
				they dare not risk her displeasure. Raising their voices
				to her is difficult; raising a hand against her is unthink-
				able. Those few who shake off the vampire’s potent
				mystique enough to oppose her are shouted down by
				the many under her thrall, before the immortal need
				even respond.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `No roll is required on the part of the vam-
					pire, but she must spend a Willpower point. A subject
					must make a Courage roll (difficulty equal to the vam-
					pire’s Charisma + Intimidation, to a maximum of 10) if
					he wishes to be rude or simply contrary to the vampire.
					Success allows the individual to act normally for the
					moment, although he feels the weight of the vampire’s
					displeasure crushing down on him. A subject who fails
					the roll aborts his intended action and even goes to
					absurd lengths to humble himself before the vampire,
					no matter who else is watching. The effects of Majesty
					last for one scene.`,
				},
			},
			"Love": {
				BaseDesc: `The blood bond is one of the most powerful tools in
				an elder’s inventory. However, more and more childer
				are aware of how to avoid being bound, so alterna-
				tives are needed. The Presence power called Love is
				one such alternative, as it simulates the effects of the
				bond without any of the messy side effects. While not
				as sure a method of control as a true blood bond, nor
				as long-lasting, Love is still an extremely potent means
				of command.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The player spends a blood point and rolls
					Charisma + Subterfuge (difficulty equal to the target’s
					current Willpower points). Success on the roll indi-
					cates that the victim feels as attached to the character
					as if he were blood bound to her. Each success also re-
					duces the victim’s dice pool by one die for any Social
					rolls to be made against the character. A botch makes
					the target immune to all of the character’s Presence
					powers for the rest of the night. This power lasts for
					one scene and can be applied to the same victim over
					multiple scenes in the same night.`,
				},
			},
			"Paralyzing Glance": {
				BaseDesc: `Some elders have honed their mastery of Dread Gaze
				to such a degree that they are said to be able to para-
				lyze with a look. The power‘s name is something of a
				misnomer, as victims of this power are not precisely
				paralyzed in a physical sense, but rather frozen with
				sheer terror.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `The character must make eye contact
					with her intended victim. The player then
					rolls Manipulation + Intimidation (difficulty equal to
					the target’s current Willpower points). Success renders
					the victim so terrified that he falls into a whimpering,
					catatonic state, unable to take any actions except curling into a fetal position and gibbering incoherently.
					The condition lasts for a length of time determined by
					the number of successes rolled. If the victim’s life is directly threatened (by assault, impending sunrise, etc.),
					the poor wretch may attempt to break out of his paraly-
					sis with a Courage roll (difficulty equal to the charac-
					ter’s Intimidation + 3). One success ends the paralysis.
					A botch sends the victim into a continuous state of
					Rötschreck for the rest of the night.`,
					`1 success`:    `3 turns`,
					`2 successes`:  `5 minutes`,
					`3 successes`:  `Remainder of scene`,
					`4 successes`:  `One hour`,
					`5 successes`:  `Rest of night`,
					`6+ successes`: `A week(or more)`,
				},
			},
			"Spark of Rage": {
				BaseDesc: `A vampire possessing this power can shorten tem-
				pers and bring old grudges and irritations to the boiling
				point with a minimum of effort. Spark of Rage causes
				disagreements and fights, and can even send other
				vampires into frenzy.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The player spends a blood point and rolls
					Manipulation + Subterfuge (difficulty 8). The number of individuals affected is determined by how many
					successes are rolled. If this power is used in a crowd,
					those affected are the people in closest proximity to
					the character. A vampire affected by this power must
					spend a Willpower point or roll Self-Control/Instinct
					(difficulty equal to the character’s Manipulation +
					Subterfuge); failure sends the target into a frenzy. A
					botch by the vampire using Spark of Rage sends the
					invoking character into immediate frenzy.`,

					`1 success`:   `Two people`,
					`2 successes`: `Four people`,
					`3 successes`: `Eight people`,
					`4 successes`: `20 people`,
					`5 successes`: `Everyone in the character’s
					immediate vicinity`,
				},
			},
			"Cooperation": {
				BaseDesc: `Any elder knows that Kindred are the most difficult
				beings in existence to force to work together. Peaceful
				coexistence is not a common tenet of vampiric society.
				With that in mind, this power can be used to nudge
				those affected by it into a fragile spirit of camaraderie.
				Some cynical (or realistic) Ventrue claim that their
				Clan’s mastery of this Presence effect is the sole reason
				that anything is ever accomplished in Camarilla conclaves. Ventrue who voice this opinion too loudly also
				tend to have numerous chances to test just how effective Cooperation is.`,
				Lvl: 7,
				System: map[string]string{
					"sys": ` To invoke Cooperation, the player spends a
					blood point and rolls Charisma + Leadership (difficulty
					8). The number of individuals affected is determined by
					how many successes the player rolls. Cooperation lasts
					for the remainder of the scene in which it is invoked,
					though particularly strong users of Presence may create
					longer-lasting feelings of non-aggression (at Storyteller
					discretion) by spending Willpower. While this power
					is in effect, all those under its influence are more favor-
					ably disposed toward one another and are more willing
					to extend trust or make cooperative plans.`,
					`1 success`:   `Two people`,
					`2 successes`: `Four people`,
					`3 successes`: `Eight people`,
					`4 successes`: `20 people`,
					`5 successes`: `Everyone in the character’s immediate vicinity`,
				},
			},
			"Ironclad Command": {
				BaseDesc: `Any individual can normally resist the powers of
				Presence for a brief time through an effort of will. Some
				elder Toreador and Ventrue have developed such force
				of personality that their powers of Presence cannot be
				resisted without truly heroic efforts.`,
				Lvl: 8,
				System: map[string]string{
					"sys": ` This power is always in effect once it has
					been learned. A mortal may not spend Willpower to resist the character’s Presence (for purposes of this power,
					the definition of “mortal” does not include supernaturally active humans such as ghouls or those who possess
					True Faith). A supernatural being must roll Willpower
					(difficulty of the character’s Willpower + 2; difficulties
					over 10 mean that the roll cannot even be attempted)
					the first time he attempts to spend a Willpower point
					to overcome the character’s Presence. For the rest of
					the night, the maximum number of Willpower points
					he can spend to resist the vampire’s Presence powers
					is equal to the number of successes he rolled. A botch
					doubles the character’s Presence dice pools against the
					hapless victim for the remainder of the night.`,
				},
			},
			"Pulse of the City": {
				BaseDesc: `A vampire who has developed her Presence to this
				terrifying degree can control the emotional climate of
				the entire region around her, up to the size of a large
				city. This power is always in effect on a low level, at-
				tuning those who dwell in the area to the Kindred’s
				mood, but it can also be used to project a specific emo-
				tion into the minds of every being in the area. Pulse
				of the City affects residents much more strongly than
				tourists, and also has a significant impact on those in-
				dividuals who might be elsewhere at the time but who
				still have strong ties to the affected city.`,
				Lvl: 9,
				System: map[string]string{
					"sys": ` The character must be present in the city in
					question, and have at least a casual, personal knowledge of its streets and makeup (maps won’t help). The
					player spends a Willpower point and rolls Charisma +
					Streetwise (difficulty 9, though specializations or Storyteller may decrease this difficulty if the character
					is intimately familiar with a particular city). The number of successes indicates how long mortal residents are
					affected by the particular emotion that the character
					broadcasts; visitors with no ties to the area and supernatural beings are affected for a duration one success
					step lower. The character can choose to terminate this
					effect at any time before it expires.`,
					`1 success`:   `One minute`,
					`2 successes`: `10 minutes`,
					`3 successes`: `One hour`,
					`4 successes`: `One day`,
					`5 successes`: `One week`,
				},
			},
		},
	},
	"protean": {
		Name: "Protean",
		Description: `Protean allows the Kindred the mystical ability to
		manipulate his physical form. Some vampires believe
		the power stems from a heightened connection to the
		natural world, while others consider it to be a magnification of the mark of Caine. Whatever its basis may
		be, those that develop this Discipline can grow bestial
		claws, take on the forms of bats and wolves, turn themselves into mist, and even meld into the very earth itself.
		Transformed Kindred can generally use other Disciplines — vampires in wolf form can still read auras and
		communicate with other animals, for example. However, the Storyteller may rule that certain Disciplines
		may not be used in specific situations. The Kindred’s
		clothes and personal possessions also change when he
		transforms (presumably absorbed within his very substance), although armor and the like do not provide
		any benefit while transformed.`,
		Abilities: map[string]disciplineAbility{
			"Eye of the Beast": {
				BaseDesc: `The vampire sees perfectly well in pitch darkness,
				not requiring a light source to notice details in even
				the darkest basement or cave. The vampire’s Beast is
				evident in his red glowing eyes, a sight sure to disturb
				most mortals.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The character must declare his desire to call
					forth the Eyes. No roll is necessary, but the change re-
					quires a full turn to complete. While manifesting the
					Eyes, the character suffers a +1 difficulty to all Social
					rolls with mortals unless he takes steps to shield his
					eyes (sunglasses are the simplest solution). (A vampire
					without this power who is immersed in total darkness
					suffers blind-fighting penalties)`,
				},
			},
			"Feral Claws": {
				BaseDesc: `The vampire’s nails transform into long, bestial claws.
				These talons are wickedly sharp, able to rend flesh with
				ease and even carve stone and metal with little trouble.
				The Beast is prominent in the claws as well, making
				them fearsome weapons against other immortals. It’s
				rumored that some Gangrel have modified this power
				to change their vampiric fangs into vicious tusks.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The claws grow automatically in response to
					the character’s desire, and can grow from both hands
					and feet. The transformation requires the expenditure
					of a blood point, takes a single turn to complete, and
					lasts for a scene.
					The character attacks normally in combat, but the
					claws inflict Strength + 1 aggravated damage. Other
					supernaturals cannot normally soak this damage, although a power such as Fortitude may be used. Additionally, the difficulties of all climbing rolls are reduced
					by two.`,
				},
			},
			"Earth Meld": {
				BaseDesc: `One of the most prized powers within Protean,
				Earth Meld enables the vampire to become one with
				the earth. The immortal literally sinks into the bare
				ground, transmuting his substance to bond with the
				earth.
				Though a vampire can immerse himself fully into the
				ground, he cannot move around within it. Further, it
				is impossible to meld into earth through another sub-
				stance. Wood slats, blacktop, even artificial turf blocks
				Earth Meld’s effectiveness — then again, it’s a rela-
				tively simple matter for a vampire at this level of power
				to grow claws and rip apart enough of the flooring to
				expose the raw soil beneath.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `No roll is necessary, although the charater must spend a blood point. Sinking into the earth is
					automatic and takes a turn to complete. The character
					falls into a state one step above torpor during this time,
					sensing his surroundings only distantly. The player
					must make a Humanity or Path roll (difficulty 6) for
					the character to rouse himself in response to danger
					prior to his desired time of emergence.
					Since the character is in an in-between state, any attempts to locate him (catching his scent, scanning for
					his aura, traveling astrally, and so on) are made at +2
					difficulty. Astral individuals cannot affect the vampire
					directly, instead meeting with a kind of spongy resistance as their hands pass through him. Similarly, digging in the material world encounters incredibly hardpacked earth, virtually as dense as stone. Attempts at violence upon the submerged vampire
					from either side return him to his physical nature, expelling the soil with which he bonded in a blinding
					spray (all Perception-based rolls are at +2 difficulty for
					the turn). The character himself subtracts two from
					his initiative for the first turn after his restoration, due
					to momentary disorientation. Once expelled from the
					earth, the vampire may act normally.`,
				},
			},
			"Shape of the Beast": {
				BaseDesc: `This endows the vampire with the legendary ability
				to transform into a wolf or bat. A Kindred changed
				in this way is a particularly imposing representative
				of the animal kingdom. Indeed, he is far superior to
				normal animals, even ones possessed by Subsume the
				Spirit. He retains his own psyche and temperament,
				but can still call upon the abilities of the beast form
				— increased senses for the wolf and flight for the bat.
				Gangrel are reputed to change to other animal forms
				better suited to their environment — jackals in Africa,
				dholes in Asia, and even enormous rats in urban envi-
				ronments — a feat that other Clans learning Protean
				cannot seem to duplicate.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The character spends one blood point to assume the desired shape. The transformation requires
					three turns to complete (spending additional blood
					points reduces the time of transformation by one turn
					per point spent, to a minimum of one). The vampire
					remains in his beast form until the next dawn, unless
					he wishes to change back sooner.
					While in the animal’s shape, the vampire can use any
					Discipline he possesses except Necromancy, Serpentis, Thaumaturgy, or Vicissitude (as well as any others
					the Storyteller deems unavailable). Furthermore, each
					form gives the character the abilities of that creature.
					In wolf form, the vampire’s teeth and claws inflict
					Strength + 1 aggravated damage, he can run at double
					speed, and the difficulties of all Perception rolls are reduced by two. In bat form, the vampire’s Strength is reduced to 1, but he can fly at speeds of up to 20 miles per
					hour, difficulties for all hearing-based Perception rolls
					are reduced by three, and attacks made against him are
					at +2 difficulty due to his small size.
					The Storyteller may allow Gangrel to assume a different animal shape, but should establish the natural
					abilities it grants the character.`,
				},
			},
			"Mist Form": {
				BaseDesc: `This truly unsettling power enables the vampire to
				turn into mist. His physical shape disperses into a hazy
				cloud, but one still subject entirely to the immortal’s
				will. He floats at a brisk pace and may slip under doors,
				through screens, down pipes, and through other tiny
				openings. Although strong winds can blow the vampire from his chosen course, even hurricane-force winds
				cannot disperse his mist shape.
				Some Kindred feel that this power is an expression of
				the vampire’s ultimate control over the material world,
				while others believe that it is the immortal’s soul made
				manifest (damned though it may be)`,
				Lvl: 5,
				System: map[string]string{
					"sys": ` No roll is required, although a blood point
					must be spent. The transformation takes three turns to
					complete, although the character may reduce this time
					by one turn for each additional blood point spent (to
					a minimum of one turn). Strong winds may buffet the
					character, although Disciplines such as Potence may
					be used to resist them. Vampires in Mist Form can perceive their surroundings normally, although they cannot use powers that require eye contact.
					The vampire is immune to all mundane physical attacks while in mist form, although supernatural attacks
					affect him normally. Also, the vampire takes one fewer
					die of damage from fire and sunlight. The character may
					not attack others physically while in this state — this
					includes encountering another vampire in mist form.
					He may use Disciplines that do not require physical
					substance, however.`,
				},
			},
			"Earth Control": {
				BaseDesc: `An Earth Melded character with this power is no
				longer confined to the resting place she selected the
				night before. She can pass through the ground as if it
				were water, “swimming” through the earth itself. Some
				elders use this as a means of unobstructed and unobtru-
				sive travel, while others find it a highly effective means
				of maneuvering in combat.`,
				Lvl: 6,
				System: map[string]string{
					"sys": ` This power is in effect whenever a character
					is Earth Melded, with no additional roll or expenditure
					necessary. While in the ground, a vampire can propel
					herself at half of her normal walking speed. She cannot
					see, but gains a supernatural awareness of her underground surroundings out to a range of 50 yards or meters. Water, rock, tree roots, and cement all effectively
					block her progress; she can only move through earth
					and substances of similar consistency, such as sand or
					fine gravel. If two or more vampires attempt to interact underground, only direct physical contact is possible. All damage dice pools in this case are halved,
					and dodge and parry attempts are at -2 difficulty. If
					an underground chase takes place, it is resolved with
					an extended, contested Strength + Athletics roll`,
				},
			},
			"Flesh of Marble": {
				BaseDesc: `Tales have long spoken of the combat prowess of
				Gangrel elders and of their inhuman resilience. Poorly
				informed individuals believe the stories of swords shat-
				tering and bullets flattening against immortal skin
				to be exaggerated reports of the effects of Fortitude.
				Those with more reliable information know that such
				tales result from encounters with vampires who have
				developed Flesh of Marble. The skin of an elder with
				Flesh of Marble becomes in essence a sort of flexible
				stone, although it appears (and feels) no different than
				normal skin and muscle.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The player spends three blood points to activate Flesh of Marble, which goes into effect instantly.
					The effects of the power last for the remainder of the
					scene. While the power is functioning, the damage dice
					pools of all physical attacks made against the character
					are halved (round down). That includes assaults made
					with fists, claws, swords, firearms, and explosions, but
					not fire, sunlight, or supernatural powers (unless the
					effect in question is a direct physical attack, such as a
					rock hurled by means of Movement of the Mind). Additionally, while this power is in effect, a character can
					attempt to parry melee attacks with his bare hands as if
					he were holding some form of weapon.`,
				},
			},
			"Restore the Mortal Visage": {
				BaseDesc: `Cainites are of two opinions regarding this power.
				Those who are politically active, or who associate extensively with mortals, view it as both necessary and
				acceptable. Those Kindred who embrace their more feral sides, however, see it as a disgusting defiance of the
				very nature of vampirism. The schism comes because
				the power allows the elder who possesses it to temporarily return his appearance to what it was before the
				Embrace, removing the bestial features he has accumulated over the centuries. Restore the Mortal Visage
				has only been displayed by Gangrel; several Nosferatu
				elders have attempted to develop it, and it is whispered
				that they met spontaneous, grotesque Final Deaths
				when they attempted to take their mortal forms.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `The player spends three blood points and
					a Willpower point and rolls Willpower (difficulty 8).
					Success restores the character’s appearance to what it
					was just before he was first Embraced, erasing all physical or social animalistic features gained from frenzies. The power also affects any of the character‘s affected Traits relating to social interaction, returning them to their original values (for example, lost
					Appearance dots return, but Humanity points removed
					due to frenzy are not). A botched Willpower roll earns
					the character another feature, as per the Gangrel Clan
					weakness. Once activated, Restore the Mortal Visage
					lasts for the remainder of the scene.`,
				},
			},
			"Shape of the Beast’s Wrath": {
				BaseDesc: `Users of this power are often mistaken for Tzimisce
				employing the Vicissitude power Horrid Form. A vam-
				pire employing this power shifts into a huge, monstrous
				form, gaining half her height again and tripling her
				weight. Her overall shape flows into an unholy amal-
				gamation of her own form and that of the animal she
				feels the closest kinship to (wolves, rats, and great cats
				are the most common manifestations, though ravens,
				serpents, bats, and stranger beasts have been reported).
				The vampire’s new shape does bear some vague resem-
				blance to the war-forms of the werecreatures, but the
				difference quickly becomes apparent.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `The player spends three blood points, the
					expenditure of which triggers the change. The character’s transformation takes three turns (the player may
					spend additional blood points to reduce this time at a
					cost of one point per turn of reduction). Once transformed, the character remains in this form until sunrise
					or until she shifts back voluntarily.
					The precise Traits of this form are determined when
					the character first learns this power, as is the animal
					whose appearance the character takes on. The vampire’s new form adds a total of seven dots to the character’s Physical Attributes. At least one dot must go
					into each Physical Attribute, meaning that no more
					than five can go into any one (so a character could
					have +5 Strength, +1 Dexterity, and +1 Stamina, but
					not +2 Strength and +5 Dexterity). These bonuses are
					always the same once they are selected; a different allocation requires that the character buy this power a
					second time and thus purchase another alternate form.
					Additionally, the character inflicts Strength + 2 dice
					of aggravated damage with both bite and claw attacks
					when in monstrous form. She also gains an extra Hurt
					health level, and doubles her normal running speed.
					Finally, the character’s perceptions are also heightened. She is assumed to have both the Auspex power
					Heightened Senses and the Protean power Eyes of the
					Beast after transformation, with all of the benefits and
					drawbacks of each.
					This form does carry two additional drawbacks. The
					first is a lack of communication ability. The character’s
					Social Attributes all drop to 1, or to 0 if they already
					were 1 (except when making Intimidation rolls) when
					the transformation occurs. The second problem that
					a character in this form encounters is the suddenly
					heightened power of her Beast. All difficulties of rolls
					to resist frenzy are increased by two for the duration of
					the power’s effect, and the player may not spend Willpower points on such rolls.`,
				},
			},
			" Spectral Body": {
				BaseDesc: `This powerful variation on Mist Form allows a vam-
				pire to take a shape with most of the advantages of the
				lesser power but fewer of the disadvantages. A vampire
				who assumes Spectral Form retains his normal appear-
				ance, but becomes completely insubstantial. He walks
				through walls and bullets with equal ease, and can pass
				through the floor on which he stands if he chooses to.
				Although his lungs are no longer solid, the vampire can
				still speak, a fact in which some elders of the Daughters
				of Cacophony bloodline have expressed great interest.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `The player spends three blood points. The
					transformation takes one turn to complete, and lasts
					for the rest of the night unless the character decides to
					reverse it. When the power takes hold, the character
					becomes completely insubstantial, but remains fully
					visible. Henceforth, he is unaffected by any physical
					attacks, and he doubles his dice pool to soak damage
					from fire and sunlight. The vampire may even ignore
					gravity if he chooses to do so, rising and sinking through
					solid objects if he does not wish to stand on them (although he may move no faster than his normal walking speed while “flying” in this manner). While in this
					form, the character may also use any Disciplines that
					do not require physical contact or a physical body. On
					the downside, while in Spectral Form, a vampire can
					physically manipulate his environment only through
					the use of powers such as Movement of the Mind.`,
				},
			},
			" Purify the Impaled Beast": {
				BaseDesc: `Camarilla records indicate that a disproportionately
				small number of Gangrel elders were killed by mortals
				and Anarchs during the Inquisition and the Anarch
				Revolt. This power is one of the primary reasons for
				the survival of these Kindred. An elder with this Pro-
				tean power can expel foreign objects from her body
				with great force, even excising stakes that transfix her
				heart.`,
				Lvl: 8,
				System: map[string]string{
					"sys": `The player spends three blood points and
					rolls Willpower (difficulty 6, or 8 if the vampire is paralyzed by an object impaling her heart). One success is
					sufficient to remove all foreign objects and substances
					from the character’s body. Dirt, bullets, and even stakes
					through the heart are instantly and violently removed.
					The larger the object, the farther away it is hurled by
					this power. Objects expelled thus are considered to
					have an attack dice pool of three for any bystanders,
					and to have a dice pool of one to four (depending on
					size) for damage. If the character wishes to leave an object in his body (such as a prosthetic limb) or partially
					in (expelling a stake from his heart but leaving it sticking out of his breastbone as a ruse), the player must also
					spend a Willpower point when activating this power.`,
				},
			},
			"Inward Focus": {
				BaseDesc: `This power has no outwardly visible effects what-
				soever. Indeed, its very existence is unknown outside
				those Gangrel Methuselahs who have developed it.
				The internal effects of this razor-sharp honing of Pro-
				tean, however, are in some ways more dramatic than
				any external manifestation. A vampire with this power
				can heighten the efficiency of his undead body’s inter-
				nal workings to levels undreamed of by lesser Kindred,
				withstanding inconceivable amounts of injury and
				moving with blinding speed and shattering strength.`,
				Lvl: 9,
				System: map[string]string{
					"sys": ` The player spends four blood points to ac-
					tivate this power and an additional two blood points
					for every turn past the first that Inward Focus is main-
					tained. There are three effects of this power: First, the
					character gains a number of extra actions during each
					turn equal to his unmodified Dexterity score. Second,
					the damage of his physical attacks is increased by three
					dice per dice pool. Finally, all damage inflicted on the
					character is halved and rounded down after the soak
					roll is made (so an attack that inflicts five health levels
					after soak is reduced to two health levels).
					This power may be used in conjunction with other
					Protean powers that modify the character‘s combat
					abilities, such as Shape of the Beast’s Wrath (above).
					It may also be used in conjunction with Celerity, For-
					titude, and Potence, turning a vampire who has mas-
					tered this power into a truly terrifying opponent.`,
				},
			},
		},
	},
	"quietus": {
		Name: "Quietus",
		Description: `The Discipline of silent death, Quietus is practiced
		by those of Clan Assamite. Based on elements of
		blood, poison, vitae control, and pestilence, Quietus
		focuses on the destruction of a target through a variety
		of means. This Discipline doesn’t always cause a quick
		death, but the Assamites rely on its lethality to hide
		their involvement with their victims.`,
		Abilities: map[string]disciplineAbility{
			"Silence of Death": {
				BaseDesc: `Many Assamites claim never to have heard their
				targets’ death screams. Silence of Death imbues the
				vampire with a mystical silence that radiates from her
				body, muting all noise within a certain vicinity. No
				sound occurs inside this zone, though sounds originating outside the area of effect may be heard by anyone
				in it. Rumors abound of certain skilled Assamite viziers
				who have the ability to silence a location rather than a
				circumference that follows them, but no proof of this
				has been forthcoming.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `This power (which costs one blood point to
						activate) maintains a 20-foot (6-meter) radius of utter
						stillness around the Kindred for one hour.`,
				},
			},
			"Scorpion’s Touch": {
				BaseDesc: `By changing the properties of her blood, a vampire
				may create powerful venom that strips her prey of his resilience. This power is greatly feared by other Kindred,
				and all manner of hideous tales concerning methods of
				delivery circulate among trembling coteries. Kindred
				with Quietus are known to deliver the poison by coating their weapons with it, blighting their opponents
				with a touch, or spitting it like a cobra. An apocryphal
				account speaks of a proud Prince who discovered an
				Assamite plotting her exsanguination and began to
				diablerize her would-be assassin. Halfway through the
				act, she learned that she had ingested a dire quantity of
				tainted blood and was then unable to resist the weakened hashashiyyin’s renewed attack.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `To convert a bit of her blood to poison, the
					Kindred’s player spends at least one blood point and
					rolls Willpower (difficulty 6). If this roll is successful,
					and the vampire successfully hits (but not necessarily
					damages) her opponent, the target loses a number of
					Stamina points equal to the number of blood points
					converted into poison — vampires attempting to drink
					the blood of the Kindred with Scorpion’s Touch are
					automatically considered to be “successfully hit.”
					The victim may resist the poison with a Stamina +
					Fortitude roll (difficulty 6); successes achieved on the
					resistance roll subtract from the vampire’s successes.
					The maximum number of blood points a Kindred may
					convert at any one time is equal to her Stamina. The
					number of successes scored indicates the duration of
					the Stamina loss.`,
					`1 success`:   `One turn`,
					`2 successes`: `One hour`,
					`3 successes`: `One day`,
					`4 successes`: `One month`,
					`5 successes`: `Permanently (though Stamina may be bought back up with experience`,
				},
			},
			"Dagon’s Call": {
				BaseDesc: `This terrible power allows a vampire to drown her
				target in his own blood. By concentrating, the Kindred
				bursts her target’s blood vessels and fills his lungs with
				vitae that strangles him from within. The blood actually constricts the target’s body from the inside as it
				floods through his system; thus, it works even on unbreathing Kindred. Until the target collapses in agony
				or death throes, this power has no visible effect, and
				many Kindred like it because it leaves no trace of their
				presence.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The vampire must touch her target prior
					to using Dagon’s Call. Within an hour thereafter, the
					vampire may issue the call, though she need not be in
					the presence or even in the line of sight of her target.
					Invoking the power costs one Willpower point.
					The Kindred’s player makes a contested Stamina roll
					against the target’s Stamina; the difficulty of each roll
					is equal to the opponent’s permanent Willpower rating. The number of successes the vampire using Dagon’s Call achieves is the amount of lethal damage, in
					health levels, the victim suffers. For an additional point
					of Willpower spent in the next turn, the vampire may
					continue using Dagon’s Call by engaging in another
					contested Stamina roll. So long as the Kindred’s player
					continues to spend Willpower, the character may continue rending her opponent from within.`,
				},
			},
			"Baal’s Caress": {
				BaseDesc: `The penultimate use of blood as a weapon (short of
					diablerie itself), Baal’s Caress allows the Kindred to
					transmute her blood into a virulent ichor that destroys
					any living or undead flesh it touches. In nights of yore,
					when Assamites led the charges of Saracen legions, the
					Assassins were often seen licking their blades, slicing
					open their tongues and lubricating their weapons with
					this foul secretion.
					Baal’s Caress may be used to augment any bladed
					weapon; everything from poisoned knives and swords
					to tainted fingernails and claws has been reported.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `Baal’s Caress does not increase the damage
					done by a given weapon, but that weapon inflicts aggravated damage rather than normal. No roll is necessary to activate this power, but one blood point is
					consumed per hit. For example, if a Cainite poisons his
					knife and strikes his opponent (even if he inflicts no
					damage), one blood point’s worth of lubrication disappears. For this reason, many vampires choose to coat
					their weapons with a significant quantity of blood. If
					the vampire misses, no tainted blood is consumed`,
				},
			},
			"Taste of Death": {
				BaseDesc: `A refinement of Baal’s Caress, Taste of Death allows
				the Cainite to spit caustic blood at her target. The
				blood coughed forth with this power burns flesh and
				corrodes bone; some vampires have been reported to
				vomit voluminous streams of vitae that reduce their
				targets to heaps of sludge.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The vampire may spit up to 10 feet (3 meters) for each dot of Strength and Potence he possesses.
						Hitting the target requires a Stamina + Athletics roll
						(difficulty 6). Each blood point spewed at the target
						inflicts two dice of aggravated damage, and there is no
						limit (other than the vampire’s capacity and per-turn
						expenditure maximum) to the quantity of blood with
						which a target may be deluged.`,
				},
			},
			"Blood Sweat": {
				BaseDesc: `Although vampires do not have functioning sebaceous glands, they are still capable of sweating at
				times of extreme stress. This “sweat” is actually a thin
				sheen of blood on the Cainite’s forehead and palms.
				Most Kindred see blood sweat as an admission of fear
				or guilt. The vampire who has mastered Blood Sweat is
				capable of inducing these feelings in a subject to a preternatural degree. The victim experiences a torrential
				outpouring of vitae if he harbors the tiniest shred of
				remorse for any action he has ever undertaken.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The character must be within line of sight
					of the subject and spend three turns concentrating.
					The player spends a Willpower point and rolls Manip-
					ulation + Intimidation (difficulty equal to the target’s
					current Willpower points). The target loses one blood
					point per success. Mortals sustain injury as if they had
					lost blood from being fed upon. The target actually
					“sweats out” the lost blood in a sudden rush of sanguinary perspiration that soaks his clothes. Large amounts
					may even form puddles at his feet. Blood lost through
					this process is considered dead, inert mortal blood, and
					provides minimal nutrition (half normal) for Cainites
					desperate enough to lick it up from the floor or wring
					it out of towels. It provides no sustenance for the individual from whom it emerged.
					In addition to the blood loss, the victim is overcome
					by a sense of remorse and guilt for his past transgressions (if he has a strong conscience) or an overwhelming compulsion to brag (if he is of sufficiently coarse
					moral character). The severity of this impulse depends
					on the number of successes rolled: One success could
					cause a slight twinge of conscience, while five successes
					may result in the subject breaking down and pouring
					out a full confession of his crimes. This effect is more
					story-oriented than mechanical, and the Storyteller is
					the final authority on what the victim feels compelled
					to confess or boast.`,
				},
			},
			"Selective Silence": {
				BaseDesc: `Although Silence of Death is an effective tool for
				the battlefield and the court alike, it is indiscriminate
				in its effects. The assassin who uses it in preparation
				for firing a shotgun also silences the radio over which
				her comrades might warn her of an oncoming guard.
				The courtier who suppresses a room full of dissenting
				voices is likewise unable to speak her own mind. Selective Silence allows the skilled Cainite to overcome
				these limitations by silencing only those individuals or
				objects that she wants to silence.
				When using this power, most individuals exhale a
				thin mist of blood that clings to the selected subjects,
				gradually evaporating as Selective Silence’s effects
				fade. Some vampires also use a similar technique when
				invoking Silence of Death, in which case the mist surrounds them and moves with them.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The player spends two blood points and
					rolls Stamina + Stealth (difficulty 7). Each success allows for one individual or object that the character can
					silence with this use of the power. Each subject must
					be within 20 yards or meters of the character. Objects
					larger than a man count as more than one subject: a
					heavy machine gun counts as two, a car as three, and
					a small aircraft as five. Objects larger than a private jet
					or creatures larger than an elephant cannot be silenced
					through the use of this power.
					Each subject is completely silenced for a number of
					minutes equal to the character’s permanent Willpower.
					Nothing it does generates sound, though the secondary
					effects of its actions will do so normally. For example, a
					gun silenced with this power will not produce an audible explosion when fired, but its bullets still make noise
					as they break the sound barrier. A silenced victim can
					scream all he likes and not make a sound, but may be
					able to summon help by smashing a window.`,
				},
			},
			"Ripples of the Heart": {
				BaseDesc: `According to Assamite lore, this technique origi-
				nated with a Byzantine scholar who wanted to protect
				his herd from the thirst of other Cainites. Ripples of
				the Heart allows a Cainite to leave emotions within
				the bloodstream of any mortal from whom he feeds.
				Any other vampire who subsequently drinks from that
				mortal experiences those emotions as if they were his
				own.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The character drinks at least one blood
					point from the subject mortal then spends a minute in
					physical contact with the subject while concentrating
					on the emotion he wishes to leave in her blood. The
					player spends a point of Willpower and rolls Charisma
					+ Empathy (difficulty 7 under normal circumstances,
					5 if the subject is currently experiencing the intended
					emotion, 9 if he is currently experiencing a strong opposite emotion). The subject’s blood carries the weight
					of the intended emotion for one lunar month per suc-
					cess rolled. A mortal’s blood can only carry one emo-
					tion at a time. Subsequent attempts to use Ripples of
					the Heart on the same individual have no effect until
					the previous application has worn off.
					Any vampire who drinks from a vessel under the ef-
					fects of Ripples of the Heart must succeed in a Self-
					Control/Instinct roll (difficulty of the mortal’s current
					Willpower points) as soon as she swallows the first
					blood point. If she fails this roll, the vitae-borne emotion immediately overtakes her. The strength of the
					emotion depends on how many blood points she drinks.
					One blood point results in a momentary mood swing,
					two causes a significant shift in demeanor, and three
					or more generates a complete change in emotional
					state. Depending on the circumstances and the precise
					emotion, the effects of this may be spectacular or cata-
					strophic. A vampire overtaken by romantic passion
					may temporarily believe she is in love with the mortal
					(or any other convenient bystander). One who drinks
					from a hate-infused vessel may rend her prey to shreds,
					and one who takes a draught of a mortal touched with
					fear may run away screaming. The vampire remains
					subject to the emotion for a number of hours equal to
					the mortal’s permanent Willpower, though she is still
					subject to other feelings after the initial rush of sensation.`,
				},
			},
			"Purification": {
				BaseDesc: `Although most mortal cultures affix negative connotations to the spilling of blood, most Assamites
				— indeed, most vampires — have quite the opposite
				reaction to it. For them, blood is an unlife-affirming
				and reinvigorating substance. Purification works on
				this principle, using the power of vitae to cleanse and
				restore. Rather than purging foreign taints from the
				body, Purification allows its wielder to cleanse other
				individuals’ minds and souls of stains, including those
				left by the mind control of other Kindred. The vampire
				using this power expels his own blood through his skin
				and allows it to soak through his subject, slowly dissipating. As it does so, it carries away spiritual impurities
				and outside influences.`,
				Lvl: 6,
				System: map[string]string{
					"sys": ` The character touches the forehead of her
					intended subject, and both parties spend a minimum of
					five minutes in deep concentration. The player spends
					a number of blood points equal to the subject’s permanent Willpower. The subject rolls Willpower once for
					every external supernatural influence (a vampiric Discipline, usually) to which his mind has been subjected.
					The difficulty equals the level of the power in question
					+4 (or a difficulty of 7 if the power level is unknown,
					such as one used by a different kind of supernatural
					creature). A success nullifies that effect.
					Purification has its limits. It can remove directly intrusive influences such as Dominate-implanted commands, Dementation-generated derangements, or the
					imperatives caused by elder vampires’ Presence. It cannot dispel influences that are transmitted by blood, including a blood bond or the dispositions imparted by
					one’s Clan or bloodline, nor can it erase those caused
					by purely mundane techniques such as persuasion,
					hypnosis, or brainwashing, or genuine emotional states
					such as love and hate. It can remove mind-altering
					blood magic effects, but either the character wielding Purification or the power’s beneficiary must have
					a level of Thaumaturgy equal to or greater than that
					with which the effect was placed. A character cannot
					use Purification on herself.`,
				},
			},
			"Baal’s Bloody Talons": {
				BaseDesc: `The toxin generated by Baal’s Caress is not enough to
				significantly harm some truly fearsome foes. This progressive development of that lesser technique allows its
				user to envenom his weapon with a blood-based poison
				so potent that it corrodes the very weapon that bears
				it, eating away at the strongest metal in a minute or
				less. However, its effects on its victims are spectacular
				enough to make the loss of even the most treasured
				blades worthwhile. This power’s effects are of very limited duration, as the substance it creates will quickly
				evaporate away.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `The character coats an edged weapon with
					her own blood, as per Baal’s Caress. The player spends
					one or more blood points and rolls Willpower (difficulty 7). The weapon now does aggravated damage. It also
					gains a number of additional damage dice equal to the
					number of successes rolled plus the number of blood
					points spent. These extra dice fade at a rate of one per
					turn as the poison dissipates, drips off, and reacts with
					the weapon’s material. Once the extra damage dice
					are all gone, the weapon’s base damage dice begin to
					fade at the same rate. The weapon breaks if used when
					its base damage is reduced to the wielder’s Strength or
					less. The only weapons that can resist this corrosion
					are those created with a supernatural power of a level
					equal to or greater than the character’s Quietus rating,
					though even this is subject to the Storyteller’s discretion.
					Baal’s Bloody Talons is subject to the same limitations as Baal’s Caress, except the limit on the number of successful strikes that do aggravated damage. A
					weapon affected by Baal’s Bloody Talons does aggravated damage with every successful attack until it is
					destroyed.
					At the Storyteller’s discretion, the character may use
					the venom this power produces for other purposes, such
					as burning through a padlock or destroying an incriminating tape. He may not, however, store this poison for
					later use — even if a container proves resistant to it,
					the substance becomes inert within a few minutes of
					leaving its creator’s body.`,
				},
			},
			"Poison the Well of Life": {
				BaseDesc: `Beyond leaving emotional traces in a subject’s blood,
				the master of this Quietus power can now taint that
				same vitae, making it into a deadly poison for any other
				Cainite who drinks from that mortal. Some vampires
				use Poison the Well of Life to guard their own herds
				against “poachers” or to ward specific vessels against
				indiscriminate feeding. Others have been known to
				employ it as a subtle trap for other vampires, turning
				herds against their owners.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `Beyond leaving emotional traces in a subject’s blood,
					the master of this Quietus power can now taint that
					same vitae, making it into a deadly poison for any other
					Cainite who drinks from that mortal. Some vampires
					use Poison the Well of Life to guard their own herds
					against “poachers” or to ward specific vessels against
					indiscriminate feeding. Others have been known to
					employ it as a subtle trap for other vampires, turning
					herds against their owners.`,
				},
			},
			"Songs of Distant Vitae": {
				BaseDesc: `Blood magic practitioners and individuals skilled at
				Auspex have long known that vitae can carry residual
				impressions of emotion and personality. This power
				invokes those impressions to overwhelm its victim
				with “remembered” images and sensations drawn from
				the vessels who held that blood before the vampire
				fed from them. Particularly strong-willed or hardened
				subjects may shrug off these visions as daydreams, but
				those who are less self-possessed can be permanently
				changed by the experience. A side effect of the use of
				this power is the partial destruction of the vitae from
				which the images are drawn. Some viziers theorize that
				this is the result of motes of the vessels’ consciousness
				making an effort to escape their usurper.`,
				Lvl: 8,
				System: map[string]string{
					"sys": `The character touches his target and spends
					a turn in concentration. The player spends four blood
					points and rolls Wits + Intimidation in a contested
					roll against the victim’s permanent Willpower (difficulty 7). If the target has committed diablerie within a
					number of nights equal to the character’s Perception,
					the attacker gains one automatic success. The result
					depends on the number of net successes the attacker
					rolls. Note that in all invocations of this power, the
					sensations that the subject relives are expressly negative and terrifying — for example, he experiences none
					of the pleasure that would normally accompany the
					Kiss when he flashes back to such an event.
					In addition to the effects listed below, the subject of
					a successful attack loses a number of blood points equal
					to the number of successes rolled. This vitae oozes from
					his body in warm red trickles that do no damage but are
					certain to terrify onlookers.`,
					`Botch`:       `The attacker enters a flashback sequence in which he relives his last feeding from the vessel’s point of view. If the player rolled three or more “1”s, the character acquires a permanent derangement related to feeding.`,
					`Failure`:     `The target is unharmed and immune to this power for a number of nights equal to his Willpower.`,
					`1 success`:   `The target experiences a brief (10- second/3-turn) flashback sequence in which he relives his last feeding from the vessel’s point of view. During this time, he is at +2 difficulty to all rolls.`,
					`2 successes`: `The target experiences a brief montage (15 seconds/5 turns) of flashbacks during which his viewpoint flashes between various feedings from the vessels’ points of view. During this time, he is at +3 difficulty to all rolls. Once the initial rush of sensation passes, he is unsettled and at +1 difficulty to all rolls until he succeeds on a Self-Control/Instinct roll (difficulty 8), which he may attempt once per scene.`,
					`3 successes`: `The target experiences a composite memory, assembled by his own subconscious, of the terror that various vessels felt while being stalked and fed upon. He must succeed in a Courage roll (difficulty 8) or instantly enter Rötschreck. If he succeeds in this roll, he is still at +3 difficulty on all actions for the rest of the scene due to the distraction that his visions impose.`,
					`4 successes`: `The target is stunned and completely unable to act for a number of turns equal to 8 - his Self-Control/ Instinct as he is bombarded by a sequence of the most terrifying memories of his various vessels. Once this initial onslaught subsides, he must roll Courage (difficulty 9) or enter Rötschreck. If he fails this roll, he must roll Self-Control/Instinct (difficulty 8) or gain the Sanguinary Animism derangement.`,
					`5 successes`: `The target is thrown into a nightmarish reenactment of the greatest fears of every individual upon whom he has ever fed. He must make a successful Self-Control/ Instinct roll (difficulty 9) or fall into torpor for (10 - the target’s Stamina) nights, at the end of which he loses half his permanent Willpower and gains the derangement Sanguinary Animism. If he succeeds on the Self-Control/Instinct roll, he enters Rötschreck for the rest of the night, during which time his greatest fear is of his own image. At the end of the night, he must succeed at a Willpower roll (difficulty 9) or lose a permanent Willpower point and gain the derangement Sanguinary Animism.`,
				},
			},
			"Condemn the Sins of the Father": {
				BaseDesc: `Although the Second City’s judges recognized that
				heritage does not equal guilt, they also encountered
				many situations in which a vampire’s entire brood had
				committed the same crime. In such cases, the judges
				often decreed the same punishment for all transgressors. This technique, which modern viziers believe to
				have originated at that time, allows its wielder to administer such judgments. Through Condemn the Sins
				of the Father, a Cainite can apply lesser Quietus powers to an entire lineage.`,
				Lvl: 8,
				System: map[string]string{
					"sys": `After successfully using any lesser Quietus
					power on another vampire, the player spends a permanent Willpower point and 10 blood points and rolls
					Stamina + Occult. The difficulty of this roll is equal
					to four plus the number of Generations of the original
					target’s descendants that the player wants to affect, up
					to a maximum difficulty of 10. If the roll succeeds, every descendant of the original target within the specified range of Generations suffers the same effects that
					the original target experienced, resisting with his own
					relevant Traits. The player may exempt a number of
					potential subjects from this effect equal to twice the
					character’s Wits, but the character must know their
					faces or have tasted their vitae.`,
				},
			},
		},
	},
	"serpentis": {
		Name: "Serpentis",
		Description: `Serpentis is believed to be the legacy of Set himself, a
		gift to his children. The Followers of Set are very careful to guard this Discipline’s secrets, only teaching the
		art to those who they deem worthy. Most vampires fear
		the Setites because of the powers of Serpentis and its
		connection to snakes and reptiles; this Discipline can
		evoke a primordial fear in others, particularly those
		who recall the tale of Eden.`,
		Abilities: map[string]disciplineAbility{
			"The Eyes of the Serpent": {
				BaseDesc: `This power grants the vampire the legendary hyp-
				notic gaze of the serpent. The Kindred’s eyes become
				gold with large black irises, and mortals in the char-
				acter’s vicinity find themselves strangely attracted to
				him. A mortal who meets the vampire’s beguiling gaze
				is immobilized. Until the character takes his eyes off
				his victim, the person is frozen in place.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `No roll is required, but this power can be
					avoided if the mortal takes care not to look into the
					vampire’s eyes. Vampires and other supernatural crea-
					tures can also be affected by this power if the Cainite’s
					player succeeds on a Willpower roll (difficulty 9). If atacked or otherwise harmed, supernatural creatures can
					spend a point of Willpower to break the spell.
					Note: The target must be able to see the
					vampire’s eyes for Eyes of the Serpent to work.`,
				},
			},
			"The Tongue of the Asp": {
				BaseDesc: `The vampire may lengthen her tongue at will, splitting it into a fork like that of a serpent. The tongue
				may reach 18 inches or half a meter, and makes a terrifyingly effective weapon in close combat.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The lash of the tongue’s razor fork causes aggravated wounds (difficulty 6, Strength damage). If the
					Kindred wounds her enemy, she may drink blood from
					the target on the next turn as though she had sunk her
					fangs into the victim’s neck. Horrifying though it is, the
					tongue’s caress is very like the Kiss, and strikes mortal
					victims helpless with fear and ecstasy. Additionally,
					the tongue is highly sensitive to vibrations, enabling
					the vampire to function effectively in the darkness the
					Clan prefers. By flicking his tongue in and out of his
					mouth, the vampire can halve any penalties relating to
					darkness.`,
				},
			},
			"The Skin of the Adder": {
				BaseDesc: `By calling upon her Blood, the vampire may trans-
				form her skin into a mottled, scaly hide. A vampire in
				this form becomes more supple and flexible.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The vampire spends one blood point and
					one Willpower point. Her skin becomes scaly and
					mottled; this, combined with the character’s increased
					flexibility, reduces soak difficulties to 5. The vampire
					may use her Stamina to soak aggravated damage from
					claws and fangs, but not from fire, sunlight, or other supernatural energies. The vampire’s mouth widens and
					fangs lengthen, enabling her bite to inflict an extra die
					of damage. Finally, the vampire may slip through any
					opening wide enough to fit her head through.
					The vampire’s Appearance drops to 1, and she is obviously inhuman if observed with any degree of care,
					though casual passersby might not notice, if the vampire is in darkness or wearing heavy clothing.`,
				},
			},
			"The Form of the Cobra": {
				BaseDesc: `The Cainite may change his form into that of a huge
				black cobra. The serpent weighs as much as the vampire’s human form, stretches over 10 feet or three meters long, and is about 20 inches (50 cm) around. The
				Form of the Cobra grants several advantages, including
				a venomous bite, the ability to slither through small
				spaces, and a greatly enhanced sense of smell. The
				character may use any Disciplines while in this form
				save those that require hands (such as Feral Claws).`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The vampire spends one blood point; the
					change is automatic, but takes three turns. Clothing
					and small personal possessions transform with the vampire. The vampire remains in serpent form until the
					next dawn, unless he desires to change back sooner.
					The Storyteller may allow the Setite bonus dice on all
					Perception rolls related to smell, but the difficulties for
					all hearing rolls are increased by two. The cobra’s bite
					inflicts damage equal to the vampire’s, but the vampire
					does not need to grapple his victim; furthermore, the
					poison delivered is fatal to mortals.`,
				},
			},
			"The Heart of Darkness": {
				BaseDesc: `A Kindred with mastery of Serpentis may pull her
				heart from her body. She can even use this ability on
				other Cainites, although this requires several hours of
				gruesome surgery. This power can only be invoked during a new moon. If performed under any other moon,
				the rite fails. Upon removing her heart, the vampire
				places it in a small clay urn, and then carefully hides or
				buries the urn. While her heart is hidden, she cannot
				be staked by any wood that pierces her breast. Moreover, because the heart is the seat of emotion, the difficulties of all her rolls to resist frenzy are two lower
				while this power is in effect. Cainites are careful to keep their hearts safe from
				danger. If someone seizes her heart, the vampire is
				completely at that person’s mercy. The heart can be
				destroyed only by casting it into a fire or exposing it to
				sunlight. If this happens, the Kindred dies where she
				stands, boiling away into a blistering heap of ash and
				blackened bone. Plunging a wooden stake into an exposed heart drives the vampire into instant torpor.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `This power requires no roll. Those who witness a vampire pull his heart from his breast (or cut
					the heart from another vampire) must make Courage
					rolls. Failure indicates anything from strong uneasiness
					to complete revulsion, possibly even Rötschreck.`,
				},
			},
			" Cobra Fangs": {
				BaseDesc: `A character using Form of the Cobra gains a venomous bite along with his serpentine form. Unfortunately,
				huge black cobras tend to make people run away as fast
				as they can. This Serpentis power enables a vampire
				to gain the deadly bite without the full-body transformation, making it more useful for taking victims by
				surprise. The police do ask questions when someone
				dies from a cobra bite under unlikely circumstances, so
				Cobra Fangs still requires some discretion in its use.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `The Kindred expends one blood point, and
					in one turn his fangs become hollow, more slender and
					venomous. The vampire injects venom when he bites.
					He must still grapple with the victim to deliver a bite
					attack, and the bite does the usual amount of damage;
					the venom, however, kills mortals within one minute.
					Bitten vampires or other supernaturally resilient creatures suffer (10 – victim’s Stamina and Fortitude) levels
					of aggravated damage over the course of five minutes.`,
				},
			},
			"Divine Image": {
				BaseDesc: `Many of the low-Generation Setite elders no longer need the illusions of Obfuscate to appear as a god.
				Through this Serpentis power, a vampire can physically metamorphose into the form of a god. Male Kindred
				generally take the form of Set himself: a muscular man
				with the head of the “Typhonic Beast,” an animal with
				a long, narrow snout and upstanding, square-topped
				ears. Less often, they take the form of the crocodile-
				headed god Sobek, whom the Egyptians often linked
				to Set, or the wolf-headed war-god Wepwawet, often
				identified with Set’s son Anubis. Female vampires generally assume the form of the cobra-headed goddess
				Renenet, wife of Sobek, or the hippopotamus-goddess
				Taweret, sometimes considered a consort of Set. Both
				were goddesses of pregnancy and childbirth. Setite
				doctrine labels all four deities as Set’s eldest childer.
				While assuming the Divine Image, the vampire becomes stronger, tougher, and more impressive. More
				importantly, perhaps, the vampire’s will becomes more
				powerful as he identifies with a divine forebear.`,
				Lvl: 7,
				System: map[string]string{
					"sys": `The character expends three blood points
					and transforms into the Divine Image in one turn. In
					the Divine Image, the vampire gains two dots each of
					Strength and Stamina and a dot each of Charisma and
					Manipulation, but her Appearance drops to 1. These
					enhancements can push the vampire over his generational limit. The character also gains two full dots of
					Willpower (to a maximum of 10). The Cainite can stay
					in the Divine Image for a full scene.
					A vampire has only one Divine Image form (unless
					the player buys this power twice). The character does
					not know what Divine Image he will manifest until he
					invokes the power the first time, although the player
					can freely choose the preferred form.`,
				},
			},
			"Heart Thief": {
				BaseDesc: `The Serpentis power Heart of Darkness normally
				takes hours to perform upon other vampires, and only
				works at the dark of the moon. Some elders, however,
				can pull the heart from another vampire’s chest with
				a quick snatch. This does not destroy the victim, unless the attacker then destroys the stolen heart. Heart
				Thief is not an easy power to use despite its speed, but
				few Discipline effects can place one vampire in another’s power so suddenly and completely`,
				Lvl: 8,
				System: map[string]string{
					"sys": `The character must expend one Willpower
					point. Removing the heart of a reluctant vampire is a
					difficult feat, comparable to staking: the attacker must
					garner at least three successes on a Dexterity + Brawl
					attack (difficulty 9). The victim may use Fortitude to
					“soak” the attacker’s successes, but mundane Stamina
					has no effect against this magical attack.
					A vampire who loses his heart this way takes one
					unsoakable level of aggravated damage, and receives
					all the benefits and drawbacks of the Heart of Darkness
					power. Resisting frenzy becomes easier (-2 difficulty)
					and he cannot be staked by wood that impales his
					breast. On the other hand, thrusting a stake through
					the removed heart will instantly force the vampire into
					torpor and exposing the heart to fire or sunlight will
					bum the vampire to ash; even biting into the heart will
					cause aggravated wounds to the vampire in question.`,
				},
			},
			"Shadow of Apep": {
				BaseDesc: `Only Set and Set’s own childer can perform this terrifying power. These ancient monsters can take the form
				of Set’s defeated enemy, Apep. The vampire becomes
				a giant serpent of fluid, glittering Darkness — not mere
				shadow, but anti-light, like the black force commanded
				by Obtenebration. In this form, physical force cannot
				harm the vampire: not claws or fangs, not bullets, not
				explosions, nothing but fire, sunlight, or mystical powers. Physical barriers cannot easily stop the vampire,
				whose shadowy form can seep through even the tiniest
				crack. The vampire, however, can still exert physical
				and supernatural force quite freely.`,
				Lvl: 9,
				System: map[string]string{
					"sys": `Taking the form of Apep costs a Willpower
					point. The transformation takes three turns to complete; once the vampire has transformed, her body remains changed for one scene. In this form, the vampire
					takes no damage from any physical attack: fists, weapons, or falling buildings pass through the vampire as if
					she were a shadow. Fire and sunlight inflict the normal aggravated damage, however, and mystical powers (such as Thaumaturgy) still affect the transformed
					vampire. The vampire’s new body gains three dots in
					each Physical Attribute. Ignore generational limits for
					this purpose. The transformed vampire can use her
					Strength to make normal close combat attacks and
					can bite for Strength + 2 dice of damage. The vampire
					can also employ any Discipline that does not require
					hands.`,
				},
			},
		},
	},
	"thaumaturgy": {
		Name: "Thaumaturgy",
		Description: `Thaumaturgy encompasses blood magic and other
		sorcerous arts available to Kindred. The Tremere Clan
		is best known for their possession (and jealous hoarding) of this Discipline. The Tremere created Thaumaturgy by combining mortal wizardry with the power
		of vampiric vitae, and as a result it is a versatile and
		powerful Discipline. Although there are whispers of
		the existence of Tremere antitribu in the Sabbat, other
		Clans in the Sword of Caine have also researched and
		developed access to such mystical might. Nevertheless,
		the Tremere of the Camarilla remain this Discipline’s
		masters.
		Like Necromancy, the practice of Thaumaturgy is
		divided into paths and rituals. Thaumaturgical paths
		are applications of the vampire’s knowledge of blood
		magic, allowing her to create effects on a whim. Ritu-
		als are more formulaic in nature, most akin to ancient
		magical “spells.” Because so many different paths and
		rituals are available to the arcane Tremere, one never
		knows what to expect when confronted with a practi-
		tioner of this Discipline.`,
		Abilities: map[string]disciplineAbility{
			"Thaumaturgy Base": {
				BaseDesc: `Thaumaturgy encompasses blood magic and other
				sorcerous arts available to Kindred. The Tremere Clan
				is best known for their possession (and jealous hoarding) of this Discipline. The Tremere created Thauma-
				turgy by combining mortal wizardry with the power
				of vampiric vitae, and as a result it is a versatile and
				powerful Discipline. Although there are whispers of
				the existence of Tremere antitribu in the Sabbat, other
				Clans in the Sword of Caine have also researched and
				developed access to such mystical might. Nevertheless,
				the Tremere of the Camarilla remain this Discipline’s
				masters.
				Like Necromancy, the practice of Thaumaturgy is
				divided into paths and rituals. Thaumaturgical paths
				are applications of the vampire’s knowledge of blood
				magic, allowing her to create effects on a whim. Rituals are more formulaic in nature, most akin to ancient
				magical “spells.” Because so many different paths and
				rituals are available to the arcane Tremere, one never
				knows what to expect when confronted with a practitioner of this Discipline`,
				Lvl: 1,
				System: map[string]string{
					"sys": `Paths define the types of magic a vampire can perform. A vampire typically learns his primary path from
					his sire, though it is not unknown for some vampires to
					study under many different tutors.
					As mentioned before, the first path a character learns
					is considered her primary path and increases automatically as the character advances in the Discipline itself.
					Secondary paths may be learned once the character
					has acquired two or more dots in her primary path, and
					they must be raised separately with experience points.
					Furthermore, a character’s rating in her primary path
					must always be at least one dot higher than any of her
					secondary paths until she has mastered her primary
					path. Once the character has achieved the fifth level
					of her primary path, secondary paths may be increased
					to that level.`,
				},
			},
			"The Path of Blood": {
				BaseDesc: `The Path of Blood encompasses some of the most fun-
				damental principles of Thaumaturgy,based as it is on
				the manipulation of Kindred vitae`,
				Lvl: 1,
			},
			"Elemental Mastery": {
				BaseDesc: `This path allows a vampire limited control over and
				communion with inanimate objects.`,
				Lvl: 1,
			},
			"The Green Path": {
				BaseDesc: `The Green Path deals with the manipulation of
				plant matter of all sorts`,
				Lvl: 1,
			},
			"Hands of Destruction": {
				BaseDesc: `The Hands of Destruction has an infamous history, and some Tremere refuse to practice it due to rumors that it is demonic in origin.`,
				Lvl:      1,
				System: map[string]string{
					"sys": ``,
				},
			},
			"The Lure of Flames": {
				BaseDesc: `This path grants the thaumaturge the ability to conjure forth mystical flames — small fires at first, but
				skilled magicians may create great conflagrations. Fire
				created by this path is not “natural.” In fact, many vampires believe the flames to be conjured from Hell itself.
				The Lure of Flames is greatly feared, as fire is one of the
				surest ways to bring Final Death upon a vampire.
				Fire conjured by The Lure of Flames must be released
				for it to have any effect. Thus, a “palm of flame” does
				not burn the vampire’s hand and cause an aggravated
				wound (nor does it cause the caster to frenzy) — it
				merely produces light. Once the flame has been re-
				leased, however, it burns normally and the character
				has no control over it.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The number of successes determines how
					accurately the vampire places the flame in his desired
					location (declared before the roll is made). One success is all that is necessary to conjure a flame in one’s
					hand, while five successes place a flame anywhere in
					the Kindred’s line of sight. Less successes mean that
					the flame appears somewhere at the Storyteller’s discretion — as a rough rule of thumb, the thaumaturge
					can accurately place a flame within 10 yards or meters
					of themselves per success.
					Individual descriptions are not provided for each
					level of this path — fire is fire, after all (including potentially causing frenzy in other vampires witnessing
					it). The chart below describes the path level required
					to generate a specific amount of flame. To soak the
					damage at all, a vampire must have the Fortitude Discipline. Fire under the caster’s control does not harm
					the vampire or cause him to frenzy, but fires started as
					a result of the unnatural flame affect the thaumaturge
					normally.`,
				},
			},
			"Neptune’s Might": {
				BaseDesc: ` This path is
				based primarily around the manipulation of standing
				water.`,
				Lvl: 1,
			},
			"Movement of the Mind": {
				BaseDesc: `This path gives the thaumaturge the ability to move
				objects telekinetically through the mystic power of
				blood. At higher levels, even flight is possible (but be
				careful who sees you…). Objects under the character’s
				control may be manipulated as if she held them — they
				may be lifted, spun, juggled, or even “thrown,” though
				creating enough force to inflict actual damage requires
				mastery of at least the fourth level of this path. Some
				casters skilled in this path even use it to guard their
				havens, animating swords, axes, and firearms to ward
				off intruders. This path may frighten and disconcert
				onlookers.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The number of successes indicates the duration of the caster’s control over the object (or subject).
					Each success allows one turn of manipulation, though
					the Kindred may attempt to maintain control after
					this time by making a new roll (she need not spend
					additional blood to maintain control). If the roll is successful, control is maintained. If a thaumaturge loses or
					relaxes control over an object and later manipulates it
					again, her player must spend another blood point, as
					a new attempt is being made. Five or more successes
					on the initial roll means the vampire can control the
					object for duration of the scene.
					If this power is used to manipulate a living being, the
					subject may attempt to resist. In this case, the caster
					and the subject make opposed Willpower rolls each
					turn the control is exercised.
					Like The Lure of Flames, individual power levels are
					not provided for this path — consult the chart below
					to see how much weight a thaumaturge may control.
					Once a Kindred reaches a rating of 3, she may levitate herself and “fly” at approximately running speed,
					no matter how much she weighs, though the weight
					restrictions apply if she manipulates other objects or
					subjects. Once a Kindred achieves 4, she may “throw”
					objects at a Strength equal to her level of mastery of
					this path.`,
					`1`: `One pound/one-half kilogram`,
					`2`: `20 pounds/10 kilograms`,
					`3`: `200 pounds/100 kilograms`,
					`4`: `500 pounds/250 kilograms`,
					`5`: `1000 pounds/500 kilograms`,
				},
			},
			"The Path of Conjuring": {
				BaseDesc: `This Thaumaturgical path enables
				powerful conjurations limited only by the mind of the
				practitioner.`,
				Lvl: 1,
				System: map[string]string{
					"sys": ``,
				},
			},
			"The Path of Corruption": {
				BaseDesc: `The Path of Corruption is primarily a mentally and
				spiritually oriented path centered on influencing the
				psyches of other individuals`,
				Lvl: 1,
			},
			"The Path of Mars": {
				BaseDesc: `This path has proven useful, turn-
				ing the tides of several confrontations with elder vam-
				pires`,
				Lvl: 1,
			},
			"The Path of Technomancy": {
				BaseDesc: `The path focuses on the control of electronic
				devices, from cellphones to laptops, and its proponents
				maintain that it is a prime example of the versatility of
				Thaumaturgy with regards to a changing world.`,
				Lvl: 1,
				System: map[string]string{
					"sys": ``,
				},
			},
			"The Path of the Father’s Vengeance": {
				BaseDesc: `The power of this path comes not only from the
				magic of blood, but also incantation of verses from the
				Book of Nod.`,
				Lvl: 1,
			},
			"Thaumaturgical Countermagic": {
				BaseDesc: `This is less of a path than it is a separate Discipline,
				as the power to resist Thaumaturgy can be taught independently of Thaumaturgy, even to those Kindred who
				are incapable of mastering the simplest ritual. Though
				the techniques of Thaumaturgical Countermagic are
				not officially taught outside Clan Tremere, unofficial
				methods are likely to exist. Any non-Tremere who
				displays the ability to resist Thaumaturgy quickly becomes the subject of potentially fatal scrutiny from the
				entirety of Clan Tremere.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `Thaumaturgical Countermagic is treated as
					a separate Discipline, although it uses the usual rules
					for Thaumaturgy (including experience costs and the
					fact that it is limited to only five levels). It cannot be
					taken as a character’s primary path, and a rating in it
					does not allow the character to perform rituals.
					The use of Thaumaturgical Countermagic is treated
					as a free action in combat and does not require a split
					dice pool. To oppose a Thaumaturgy power or ritual, a
					character must have a Thaumaturgical Countermagic
					rating equal to or greater than the rating of that power
					or ritual. The player spends a blood point and rolls the
					number of dice indicated by the character’s Thaumaturgical Countermagic rating (difficulty equal to the
					difficulty of the power in use). Each success cancels one
					of the opposing thaumaturge’s successes.
					Thaumaturgical Countermagic is only at full effectiveness when used against Thaumaturgy. It works with
					halved dice pools against Necromancy and other mystical Disciplines, and is completely ineffective against
					non-vampiric magics and powers.
					Thaumaturgical Countermagic can be learned by
					characters who are unable to learn Thaumaturgy (e.g.,
					those with the Merit Magic Resistance). Any non-
					Tremere character with a rating in this power automatically gains the Flaw Clan Enmity (Tremere), receiv-
					ing no freebie points for it. This power cannot be taken
					during character creation and cannot be spontaneously
					developed. Whether the character has Thaumaturgy as
					an in-Clan power or not, it costs the same as any other
					non-Clan Discipline to learn.`,
					`1`: `Two dice of countermagic. The character can attempt to cancel only those powers and rituals that directly affect him, his garments, and objects on his person.`,
					`2`: `Four dice of countermagic.`,
					`3`: `Six dice of countermagic. The character can attempt to cancel a Thaumaturgy power that affects anyone or anything in physical contact with him.`,
					`4`: `Eight dice of countermagic.`,
					`5`: `Ten dice of countermagic. The character can now attempt to cancel a power or ritual that targets anything within a radius equal to his Willpower in yards or meters, or one that is being used or performed within that same radius.`,
				},
			},
			"Weather Control": {
				BaseDesc: `Command over the weather has long been a staple
				power of wizards both mortal and immortal, and this
				path is said to predate the Tremere by millennia. The
				proliferation of usage of this path outside the Clan tends
				to confirm this theory; Weather Control is quite common outside the Tremere, and even outside the Camarilla. Lower levels of this path allow subtle manipulations, while higher stages of mastery allow a vampire to
				call up raging storms. The area affected by this power is
				usually rather small, no more than three or four miles
				(five to six kilometers) in diameter, and the changes
				the power wreaks are not always immediate.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The number of successes rolled indicates
					how long it takes the weather magic to take effect.
					One success indicates an entire day may pass before the
					weather changes to the thaumaturge’s liking, while a
					roll with five successes brings an almost instant effect.
					The difficulty of the Willpower roll necessary to invoke this power may change depending on the current
					local weather conditions and the weather the character is attempting to create. The Storyteller should impose a bonus (-1 or -2 difficulty) for relatively minor
					shifts, such as clearing away a light drizzle or calling
					lightning when a severe thunderstorm is already raging. Conversely, a penalty (+1 or +2 difficulty) should
					be applied when the desired change is at odds with the
					current conditions, such as summoning the same light
					drizzle in the middle of the Sahara Desert or calling
					down lightning from a cloudless sky.
					If the character tries to strike a specific target with
					lightning, the player must roll Perception + Occult
					(difficulty 6 if the target is standing in open terrain,
					8 if he is under shelter, or 10 if he is inside but near a
					window) in addition to the base roll to use Thaumaturgy. Otherwise the bolt goes astray, with the relative
					degree of failure of the roll determining where exactly
					the lightning strikes. Individual power descriptions are not provided for
					this path, as the general principle is fairly consistent.
					Instead, the strongest weather phenomenon possible at
					each level is listed.`,
					`1`: `Fog: Vision is slightly impaired and sounds are muffled; a +1 difficulty is imposed on all Perception rolls that involve sight and hearing, and the effective range of all ranged attacks are halved.
					Light breeze: A +1 difficulty is imposed on all Perception rolls that involve smell.
					Minor temperature change: It is possible to raise or lower the local temperature by up to 10 degrees Fahrenheit or 5 degrees Celsius.`,
					`2`: `Rain or snow: As Fog, but Perception rolls are impaired to a much greater extent; the difficulty modifier for all such rolls rises to +2. In addition, the difficulty on all Drive rolls increases by two.`,
					`3`: `High Winds: The wind speed rises to around 30 miles per hour or 50 kilometers per hour, with gusts of up to twice that. Ranged attacks are much more difficult: + 1 to firearm attacks and +2 to thrown weapons and archery. In addition, during fierce gusts, Dexterity roll (difficulty 6) may be required to keep characters from being knocked over by the winds. When gale-force winds are in effect, papers go flying, objects get picked up by the winds and hurled with abandon, and other suitably cinematic effects are likely.
					Moderate temperature change: The local temperature can be raised or lowered by up to 20 degrees Fahrenheit or 10 degrees Celsius.`,
					`4`: `Storm: This has the effects of both Rain and High Winds.`,
					`5`: `Lightning Strike: This attack inflicts 10 dice of lethal damage. Body armor does not add to the target’s dice pool to soak this attack.`,
				},
			},
			"Thaumaturgical Rituals": {
				BaseDesc: `Rituals are Thaumaturgical formulas, meticulously researched and prepared, that create powerful magical effects.`,
				Lvl:      1,
			},
		},
	},
	"thePathofBlood": {
		Name: "The Path of Blood",
		Description: `Almost every Tremere studies the Path of Blood as
		her primary path. It encompasses some of the most fun-
		damental principles of Thaumaturgy, based as it is on
		the manipulation of Kindred vitae. If a player wishes to
		select another path as her character’s primary path, the
		Storyteller may require additional reasoning (though
		choosing a different path is by no means unheard of)`,
		Abilities: map[string]disciplineAbility{
			"A Taste of Blood": {
				BaseDesc: `This power was developed as a means of testing a
				foe’s might — an extremely important ability in the
				tumultuous early nights of Clan Tremere. By merely
				touching the blood of his subject, the caster may determine how much vitae remains in the subject and, if
				the subject is a vampire, how recently he has fed, his
				approximate Generation and, with three or more successes, whether he has ever committed diablerie`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The number of successes achieved on the
					roll determines how much information the thauma-
					turge gleans and how accurate it is.`,
				},
			},
			"Blood Rage": {
				BaseDesc: `This power allows a vampire to force another Kin-
				dred to expend blood against his will. The caster must
				touch her subject for this power to work, though only
				the lightest contact is necessary. A vampire affected
				by this power might feel a physical rush as the thau-
				maturge heightens his Physical Attributes, might find
				himself suddenly looking more human, or may even
				find himself on the brink of frenzy as his stores of vitae
				are mystically depleted.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `  Each success forces the subject to spend one
					blood point immediately in the way the caster desires
					(which must go towards some logical expenditure the
					target vampire could make, such as increasing Physical
					Attributes or powering Disciplines). Note that blood
					points forcibly spent in this manner may exceed the
					normal “per turn” maximum indicated by the victim’s
					Generation. Each success gained also increases the
					subject’s difficulty to resist frenzy by one. The thauma-
					turge may not use Blood Rage on herself to circumvent
					generational limits.`,
				},
			},
			"Blood of Potency": {
				BaseDesc: `The thaumaturge gains such control over his own
				blood that he may effectively “concentrate” it, mak-
				ing it more powerful for a short time. In effect, he may
				temporarily lower his own Generation with this power.
				This power may be used only once per night.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `One success on the Willpower roll allows
					the character to lower his Generation by one step for
					one hour. Each additional success grants the Kindred
					either one step down in Generation or one hour of ef-
					fect. Successes earned must be spent both to decrease
					the vampire’s Generation and to maintain the change
					(this power cannot be activated again until the origi-
					nal application wears off). If the vampire is diablerized
					while this power is in effect, it wears off immediately
					and the diablerist gains power appropriate to the cast-
					er’s actual Generation. Furthermore, any mortals Em-
					braced by the thaumaturge are born to the Generation
					appropriate to their sire’s original Generation (e.g., a
					Tenth-Generation Tremere who has reduced his ef-
					fective Generation to Eighth still produces Eleventh-
					Generation childer).
					Once the effect wears off, any blood over the char-
					acter’s blood pool maximum dilutes, leaving the char-
					acter at his regular blood pool maximum. Thus, if a
					Twelfth-Generation Tremere (maximum blood pool
					of 11) decreased his Generation to Ninth (maximum
					blood pool 14), ingested 14 blood points, and had this
					much vitae in his system when the power wore off, his
					blood pool would immediately drop to 11.`,
				},
			},
			"Theft of Vitae": {
				BaseDesc: `A thaumaturge using this power siphons vitae from
				her subject. She need never come in contact with the
				subject — blood literally streams out in a physical torrent from the subject to the Kindred (though it is often
				mystically absorbed and need not enter through the
				mouth).`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The number of successes determines how
					many blood points the caster transfers from the suject. The subject must be visible to the thaumaturge
					and within 50 feet (15 meters). Using this power prvents the caster from being blood-bound, but otherwise counts as if the vampire ingested the blood herself. This power is spectacularly obvious, and Camarilla
					princes justifiably consider its public use a breach of
					the Masquerade.`,
				},
			},
			"Cauldron of Blood": {
				BaseDesc: `A thaumaturge using this power boils her subject’s
				blood in his veins like water on a stove. The Kindred
				must touch her subject, and it is this contact that sim-
				mers the subject’s blood. This power is always fatal to
				mortals, and causes great damage to even the mightiest
				vampires.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The number of successes gained determines
					how many blood points are brought to boil. The sub-
					ject suffers one health level of aggravated damage for
					each point boiled (individuals with Fortitude may soak
					this damage using only their Fortitude dice). A single
					success kills any mortal, though some ghouls with ac-
					cess to Fortitude are said to have survived after soaking
					all of the aggravated damage.`,
				},
			},
		},
	},
	"elementalMastery": {
		Name: "Elemental Mastery",
		Description: `This path allows a vampire limited control over and
		communion with inanimate objects. Elemental Mas-
		tery can only be used to affect the unliving — a vam-
		pire could not cause a tree to walk by using Animate
		the Unmoving, for instance. Thaumaturges who seek
		mastery over living things generally study paths such
		as The Green Path.`,
		Abilities: map[string]disciplineAbility{
			"Elemental Strength": {
				BaseDesc: `The vampire can draw upon the strength and resil-
				ience of the earth, or of the objects around him, to in-
				crease his physical prowess without the need for large
				amounts of blood.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The player allocates a total of three tem-
					porary bonus dots between the character’s Strength
					and Stamina. The number of successes on the roll to
					activate the power is the number of turns these dots
					remain. The player may spend a Willpower point to
					increase this duration by one turn. This power cannot
					be “stacked” — one application must expire before the
					next can be made.`,
				},
			},
			"Wooden Tongues": {
				BaseDesc: `A vampire may speak, albeit in limited fashion, with
				the spirit of any inanimate object. The conversation
				may not be incredibly interesting, as most rocks and
				chairs have limited concern for what occurs around
				them, but the vampire can get at least a general impression of what the subject has “experienced.” Note
				that events which are significant to a vampire may not
				be the same events that interest a lawn gnome.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The number of successes dictates the amount
					and relevance of the information that the character receives. One success may yield a boulder’s memory of a
					forest fire, while three may indicate that it remembers
					a shadowy figure running past, and five will cause the
					rock to relate a precise description of a local Gangrel`,
				},
			},
			"Animate the Unmoving": {
				BaseDesc: `Objects affected by this power move as the vampire
				using it dictates. An object cannot take an action that
				would be completely inconceivable for something with
				its form — for instance, a door could not leap from
				its hinges and carry someone across a street. However,
				seemingly solid objects can become flexible within rea-
				son: Barstools can run with their legs, guns can twist
				out of their owners’ hands or fire while holstered, and
				humanoid statues can move like normal humans.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `This power requires the expenditure of a
					Willpower point with less than four successes on the
					roll. Each use of this power animates one object no
					larger than human-sized; the caster may simultaneously control a number of animate objects equal to his Intelligence rating. Objects animated by this power stay
					animated as long as they are within the caster’s line of
					sight or up to an hour, although the thaumaturge can
					take other actions during that time.`,
				},
			},
			"Elemental Form": {
				BaseDesc: `The vampire can take the shape of any inanimate
				object of a mass roughly equal to her own. A desk, a
				statue, or a bicycle would be feasible, but a house or a
				pen would be beyond this power’s capacity.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The number of successes determines how
					completely the character takes the shape she wishes
					to counterfeit. At least three successes are required for
					the character to use her senses or Disciplines while in
					her altered form. This power lasts for the remainder of
					the night, although the character may return to her
					normal form at will.`,
				},
			},
			"Summon Elemental": {
				BaseDesc: `A vampire may summon one of the traditional spirits of the elements: a salamander (fire), a sylph (air),
				a gnome (earth), or an undine (water). Some thaumaturges claim to have contacted elemental spirits of
				glass, electricity, blood, and even atomic energy, but
				such reports remain unconfirmed (even as their authors are summoned to Vienna for questioning). The
				caster may choose what type of elemental he wishes to
				summon and command.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The character must be near some quantity
					of the classical element corresponding to the spirit
					he wishes to invoke. The spirit invoked may or may
					not actually follow the caster’s instructions once summoned, but generally will at least pay rough attention
					to what it’s being told to do. The number of successes
					gained on the Willpower roll determines the power
					level of the elemental.
					The elemental has three dots in all Physical and
					Mental Attributes. One dot may be added to one of
					the elemental’s Physical Attributes for each success
					gained by the caster on the initial roll. The Storyteller
					should determine the elemental’s Abilities, attacks,
					and damage, and any special powers it has related to
					its element.
					Once the elemental has been summoned, the thaumaturge must exert control over it. The more powerful the elemental, the more difficult a task this is. The
					player rolls Manipulation + Occult (difficulty of the
					number of successes scored on the casting roll + 4),
					and the number of successes determines the degree of
					control`,
					`Botch`:       `The elemental immediately attacks the thaumaturge.`,
					`Failure`:     `The elemental goes free and may attack anyone or leave the scene at the Storyteller’s discretion.`,
					`1 success`:   `The elemental does not attack its summoner.`,
					`2 successes`: `The elemental behaves favorably toward the summoner and may perform a service in exchange for payment (determined by the Storyteller).`,
					`3 successes`: `The elemental performs one service, within reason.`,
					`4 successes`: `The elemental performs any one task for the caster that does not jeopardize its own existence.`,
					`5 successes`: `The elemental performs any task that the caster sets for it, even one that may take several nights to complete or that places its existence at risk.`,
				},
			},
		},
	},
	"theGreenPath": {
		Name: "The Green Path",
		Description: `The Green Path deals with the manipulation of
		plant matter of all sorts. Anything more complex than
		an algae bloom can theoretically be controlled through
		the appropriate application of this path. Ferns, roses,
		dandelions, and even ancient redwoods are all valid
		targets for this path’s powers, and living and dead plant
		matter are equally affected. While not as immediately
		impressive as some other more widely practiced paths,
		the Green Path (sometimes disparagingly referred to as
		“Botanical Mastery”) is as subtle and powerful as the
		natural world which it affects.`,
		Abilities: map[string]disciplineAbility{
			"Herbal Wisdom": {
				BaseDesc: `With a touch, a vampire can commune with the
				spirit of a plant. Conversations held in this manner are
				often cryptic but rewarding — the wisdom and experi-
				ence of the spirits of some trees surpasses that of the
				oracles of legend. Crabgrass, on the other hand, rarely
				has much insight to offer, but might reveal the appear-
				ance of the last person who trod upon it.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The number of successes rolled determines
					the amount of information that can be gained from the
					contact. Depending on the precise information that
					the vampire seeks, the Storyteller might require the
					player to roll Intelligence + Occult in order to inter-
					pret the results of the communication.`,
					`1 success`:   `Fleeting cryptic impressions`,
					`2 successes`: `One or two clear images`,
					`3 successes`: `A concise answer to a simple query`,
					`4 successes`: `A detailed response to one or more complex questions`,
					`5 successes`: `The sum total of the plant-spirit’s knowledge on a given subject`,
				},
			},
			"Speed the Season’s Passing": {
				BaseDesc: `This power allows a thaumaturge to accelerate a plant’s
				growth, causing roses to bloom in a matter of minutes or
				trees to shoot up from saplings overnight. Alternately,
				she can speed a plant’s death and decay, withering grass
				and crumbling wooden stakes with but a touch.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The character touches the target plant. The
					player rolls normally, and the number of successes de-
					termines the amount of growth or decay. One success
					gives the plant a brief growth spurt or simulates the ef-
					fects of harsh weather, while three noticeably enlarge
					or wither it. With five successes, a full-grown plant
					springs from a seed or crumbles to dust in a few min-
					utes, and a tree sprouts fruit or begins decaying almost
					instantaneously. If this power is used in combat, three
					successes are needed to render a wooden weapon com-
					pletely useless. Two successes suffice to weaken it, while
					five cause it to disintegrate in the wielder’s hand.`,
				},
			},
			"Dance of Vines": {
				BaseDesc: `The thaumaturge can animate a mass of vegetation
				up to his own size, using it for utilitarian or combat pur-
				poses with equal ease. Leaves can walk along a desktop,
				ivy can act as a scribe, and jungle creepers can strangle
				opponents. Intruders should beware of Tremere work-
				shops that harbor potted rowan saplings.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `Any total amount of vegetation with a mass
					less than or equal to the character’s own may be animat-
					ed through this power. The plants stay active for one turn
					per success scored on the roll, and are under the complete
					control of the character. If used for combat purposes, the
					plants have Strength and Dexterity ratings each equal
					to half the character’s Willpower (rounded down) and
					Brawl ratings one lower than that of the character.
					Dance of Vines cannot make plants uproot them-
					selves and go stomping about. Even the most energetic
					vegetation is incapable of pulling out of the soil and
					walking under the effect of this power. However, 200
					pounds (100 kilograms) of kudzu can cover a consider-
					able area all by itself...`,
				},
			},
			"Verdant Haven": {
				BaseDesc: `This power weaves a temporary shelter out of a sufficient
				amount of plant matter. In addition to providing physical
				protection from the elements (and even sunlight), the
				Verdant Haven also establishes a mystical barrier which
				is nearly impassable to anyone the caster wishes to exclude. A Verdant Haven appears as a six-foot-tall (two-meter-tall) hemisphere of interlocked branches, leaves,
				and vines with no discernible opening, and even to the
				casual observer it appears to be an unnatural construction. Verdant Havens are rumored to have supernatural
				healing properties, hut no Kindred have reported experiencing such benefits from a stay in one.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `A character must be standing in a heavily
					vegetated area to use this power. The Verdant Haven
					springs up around the character over the course of three
					turns. Once the haven is established, anyone wishing
					to enter the haven without the caster’s permission
					must achieve more than the caster’s original number of
					successes on a single roll of Wits + Survival (difficulty
					equal to the caster’s Willpower). The haven lasts until
					the next sunset, or until the caster dispels or leaves it.
					If the caster scored four or more successes, the haven is
					impenetrable to sunlight unless physically breached.`,
				},
			},
			"Awaken the Forest Giants": {
				BaseDesc: `Entire trees can be animated by a master of the
				Green Path. Ancient oaks can be temporarily given
				the gift of movement, pulling their roots from the soil
				and shaking the ground with their steps. While not as
				versatile as elementals or other summoned spirits, trees
				brought to ponderous life via this power display awe-
				some strength and resilience.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The character touches the tree to be animated. The player spends a blood point and rolls normally. If the roll succeeds, the player must spend a blood
					point for every success. The tree stays animated for one
					turn per success rolled; once this time expires, the tree
					puts its roots down wherever it stands and cannot be
					animated again until the next night. While animated,
					the tree follows the character’s verbal commands to
					the best of its ability. An animated tree has Strength
					and Stamina equal to the caster’s Thaumaturgy rating,
					Dexterity 2, and a Brawl rating equal to the caster’s
					own. It is immune to bashing damage, and all lethal
					damage dice pools are halved due to its size.
					Once the animating energy leaves a tree, it puts
					down roots immediately, regardless of what it is currently standing on. A tree re-establishing itself in the
					soil can punch through concrete and asphalt to find
					nourishing dirt and water underneath, meaning that it
					is entirely possible for a sycamore to root itself in the
					middle of a road without any warning.`,
				},
			},
		},
	},
	"handsofDestruction": {
		Name: "Hands of Destruction",
		Description: `This Path is practiced most commonly by the various
		thaumaturges of the Sabbat. Though it is not widely
		seen outside that Sect, a few Camarilla Tremere have
		managed to learn the secrets of this path over the centuries. The Hands of Destruction has an infamous history, and some Tremere refuse to practice it due to rumors that it is demonic in origin.`,
		Abilities: map[string]disciplineAbility{
			"Decay": {
				BaseDesc: `This power accelerates the decrepitude of its target,
				causing it to wither, rot, or otherwise break down. The
				target must be inanimate, though dead organic matter
				can be affected.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `If the roll is successful, the inanimate ob-
					ject touched by the thaumaturge ages 10 years for every
					minute the Kindred touches it. If the vampire breaks
					physical contact and wishes to age the object again,
					another blood point must be spent and another roll
					must be made. This power does not affect vampires.`,
				},
			},
			"Gnarl Wood": {
				BaseDesc: `This power warps and bends wooden objects.
				Though the wood is otherwise undamaged, this power
				often leaves the objects completely useless. This power
				may also be used to swell or contract wood, in addi-
				tion to bending it into unwholesome shapes. Unlike
				other powers of this path, Gnarl Wood requires merely
				a glance rather than physical contact.`,
				Lvl: 2,
				System: map[string]string{
					"sys": ` Fifty pounds or twenty-five kilograms of vis-
					ible wood may be gnarled for each blood point spent
					on this power (the thaumaturge may expend as much
					blood as she likes on this power, up to her per-turn
					generational maximum). It is also possible to warp
					multiple visible objects — like all the stakes a team of
					vampire-hunters wields.`,
				},
			},
			"Acidic Touch": {
				BaseDesc: `The vampire secretes a bilious, acidic fluid from any
				portion of his body. The viscous acid corrodes metal,
				destroys wood, and causes horrendous chemical burns
				to living tissue.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The player spends one blood point to create
					the acid — the blood literally transmutes into the vola-
					tile secretion. One blood point creates enough acid to
					burn through a quarter-inch or half a centimeter of steel
					plate or three inches (seven centimeters) of wood. The
					damage from an acid-augmented hand-to-hand attack
					is aggravated and costs one blood point per turn to use.
					A thaumaturge is immune to her own acidic touch.`,
				},
			},
			"Atrophy": {
				BaseDesc: `This power withers a victim’s limb, leaving only a desiccated, almost mummified husk of bone and skin. The effects
				are instantaneous; in mortals, they are also irreversible.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The victim may resist the effects of Atrophy
					by scoring three or more successes on a Stamina + Ath-
					letics roll (difficulty 8). Failure means the limb is permanently and completely crippled. Partial resistance
					is possible: One success indicates that the difficulty of
					any roll involving the use of the arm increases by two,
					though these effects are still permanent with regard to
					mortals. Two successes signify that difficulties increase
					by one. Vampires afflicted by this power may spend five
					blood points to rejuvenate atrophied limbs. Mortals are
					permanently crippled. This power affects only limbs or
					parts of limbs (arms, legs, hands); it does not work on
					victims’ heads, torsos, etc.`,
				},
			},
			"Turn to Dust": {
				BaseDesc: `This fearsome power accelerates decrepitude in its
				victims. Mortals literally age at the mere touch of a
				skilled thaumaturge, gaining decades in moments.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `Each success on the roll ages the victim by
					10 years. A potential victim may resist with a Stami-
					na + Courage roll (difficulty 8), but must accumulate
					more successes than the caster’s activation roll — it’s
					an all-or-nothing affair. If the victim succeeds, he does
					not age at all. If he does not acquire more successes
					than the thaumaturge, he ages the full amount. Obvi-
					ously, this power, while it affects vampires, has no det-
					rimental effect on them (they’re immortal). At most, a
					Kindred victim grows paler and withers slightly (-1 to
					Appearance) for one night.`,
				},
			},
		},
	},
	"neptunesMight": {
		Name: "Neptune’s Might",
		Description: `Vampires are rarely associated with the ocean in
		most mythologies, and most Kindred have nothing
		to do with water in large quantities simply because
		they have no reason to do so. Nevertheless, Neptune’s
		Might has enjoyed a small, but devoted, following for
		centuries among Camarilla thaumaturges. This path is
		based primarily around the manipulation of standing
		water, although some of its more disturbing effects depart from this principle.
		Once a character reaches the third level of Neptune’s
		Might, the player may choose to specialize in either
		fresh water or salt water. Such specialization lowers all
		Neptune’s Might difficulties by one when dealing with
		the chosen medium but raises them by one when dealing with the opposite. Blood is considered neither fresh
		nor salty for this purpose, and difficulties in manipulating it are unaffected.`,
		Abilities: map[string]disciplineAbility{
			"Eyes of the Sea": {
				BaseDesc: `The thaumaturge may peer into a body of water and
				view events that have transpired on, in, or around it
				from the water’s perspective. Some older practitioners
				of this art claim that the vampire communes with the
				spirits of the waters when using this power; younger
				Kindred scoff at such claims.`,
				Lvl: 1,
				System: map[string]string{
					"sys": ` The number of successes rolled determines
					how far into the past the character can look.`,
					`1 success`:   `One day`,
					`2 successes`: `One week`,
					`3 successes`: `One month`,
					`4 successes`: `One year`,
					`5 successes`: `10 years`,
				},
			},
			"Prison of Water": {
				BaseDesc: `The thaumaturge can command a sufficiently large
				quantity of water to animate itself and imprison a sub-
				ject. This power requires a significant amount of fluid
				to be fully effective, although even a few gallons can be
				used to shape chains of animated water.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The number of successes scored on the roll
					is the number of successes the victim must score on a
					Strength roll (difficulty 8; Potence can add to this roll)
					to break free. A subject may be held in only one prison
					at a time, although the caster is free to invoke multiple
					uses of this power upon separate victims and may dis-
					solve these prisons at will. If a sufficient quantity of
					water (at least a bathtub’s worth) is not present, the
					difficulty of the Willpower roll to activate this power
					is raised by one.`,
				},
			},
			"Blood to Water": {
				BaseDesc: `TThe thaumaturge has now attained enough power
				over water that she can transmute other liquids to this
				basic element. The most commonly seen use of this
				power is as an assault; with but a touch, the victim’s
				blood transforms to water, weakening vampires and
				killing mortals in moments.`,
				Lvl: 3,
				System: map[string]string{
					"sys": ` The character must touch her intended vic-
					tim. The player rolls Willpower normally. Each success
					converts one of the victim’s blood points to water. One
					success kills a mortal within minutes. Vampires who
					lose blood points to this power also suffer dice pool
					penalties as if they had received an equivalent number
					of health levels of injury. The water left in the target’s
					system by this attack evaporates out at a rate of one
					blood point’s worth per hour, but the lost blood does
					not return.
					At the Storyteller’s discretion, other liquids may be
					turned to water with this power (the difficulty for such
					an action is reduced by one unless the substance is particularly dangerous or magical in nature). The character must still touch the substance or its container to use
					this power.`,
				},
			},
			"Flowing Wall": {
				BaseDesc: `Tales of vampires’ inability to cross running water
				may have derived in part from garbled accounts of this
				power in action. The thaumaturge can animate water
				to an even greater degree than is possible with the use
				of Prison of Water, commanding it to rise up to form a
				barrier impassable to almost any being.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The character touches the surface of a
					standing body of water; the player spends three Willpower points and the normal required blood point and
					rolls normally. Successes are applied to both width and
					height of the wall; each success “buys” 10 feet/three
					meters in one dimension. The wall may be placed anywhere within the character’s line of sight, and must be
					formed in a straight line. The wall lasts until the next
					sunrise. It cannot be climbed, though it can be flown
					over. To pass through the barrier, any supernatural being (including beings trying to pass the wall on other
					levels of existence, such as ghosts) must score at least
					three successes on a single Willpower roll (difficulty
					9).`,
				},
			},
			"Dehydrate": {
				BaseDesc: `At this level of mastery, the thaumaturge can directly attack living and unliving targets by removing the
				water from their bodies. Victims killed by this power
				leave behind hideous mummified corpses. This power
				can also be used for less aggressive purposes, such as
				drying out wet clothes — or evaporating puddles to
				keep other practitioners of this path from using them.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `This power can be used on any target in the
					character’s line of sight. The player rolls normally; the
					victim resists with a roll of Stamina + Fortitude (difficulty 9). Each success gained by the caster translates
					into one health level of lethal damage inflicted on the
					victim. This injury cannot be soaked (the resistance
					roll replaces soak for this attack) but can be healed
					normally. Vampires lose blood points instead of health
					levels, though if a vampire has no blood points this attack inflicts health level loss as it would against a mortal. The victim of this attack must also roll Courage
					(difficulty equal to the number of successes scored by
					the caster + 3) to be able to act on the turn following
					the attack; failure means he is overcome with agony
					and can do nothing.`,
				},
			},
		},
	},
	"thePathofConjuring": {
		Name: "The Path of Conjuring",
		Description: `Invoking objects “out of thin air” has been a staple
		of occult and supernatural legend since long before the
		rise of the Tremere. This Thaumaturgical path enables
		powerful conjurations limited only by the mind of the
		practitioner.
		Objects summoned via this path bear two distinct
		characteristics. They are uniformly “generic” in that
		each object summoned, if summoned again, would look
		exactly as it did at first. For example, a knife would be
		precisely the same knife if created twice; the two would
		be indistinguishable. Even a specific knife — the one a
		character’s father used to threaten her — would appear
		identical every time it was conjured. A rat would have
		repeated “tiled” patterns over its fur, and a garbage can
		would have a completely uniform fluted texture over
		its surface. Additionally, conjured objects bear no
		flaws: Weapons have no dents or scratches, tools have
		no distinguishing marks, and cellphones all look like
		they just came out of their packaging.
		`,
		Abilities: map[string]disciplineAbility{
			"Summon the Simple Form": {
				BaseDesc: `At this level of mastery, the conjurer may create
				simple, inanimate objects. The object cannot have any
				moving parts and may not be made of multiple materials. For example, the conjurer may summon a steel
				baton, a lead pipe, a wooden stake, or a chunk of granite.`,
				Lvl: 1,
				System: map[string]string{
					"sys": ` Each turn the conjurer wishes to keep the
					object in existence, another Willpower point must be
					spent or the object vanishes.`,
				},
			},
			"Permanency": {
				BaseDesc: `At this level, the conjurer no longer needs to pay
				Willpower costs to keep an object in existence. The
				object is permanent, though simple objects are still all
				that may be created.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The player must invest three blood points
					in an object to make it real.`,
				},
			},
			"Magic of the Smith": {
				BaseDesc: `The Kindred may now conjure complex objects of
				multiple components and with moving parts. For ex-
				ample, the thaumaturge can create guns, bicycles,
				chainsaws, or cellphones.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `Objects created via Magic of the Smith are
					automatically permanent and cost five blood points
					to conjure. Particularly complex items often require a
					Knowledge roll (Crafts, Science, Technology, etc.) in
					addition to the basic roll.`,
				},
			},
			"Reverse Conjuration": {
				BaseDesc: `This power allows the conjurer to “banish” into nonexistence any object previously called forth via this
				path.`,
				Lvl: 4,
				System: map[string]string{
					"sys": ` This is an extended success roll. The conjurer must accumulate as many successes as the original
					caster received when creating the object in question.
					This can also be used by the thaumaturge to banish
					object she created herself with this Path.`,
				},
			},
			"Power Over Life": {
				BaseDesc: `This power cannot create true life, though it can
				summon forth impressive simulacra. Creatures (and
				people) summoned with this power lack the free will
				to act on their own, mindlessly following the simple
				instructions of their conjurer instead. People created
				in this way can be subject to the use of the Dominate
				power Possession, if desired.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The player spends 10 blood points. Imperfect and impermanent, creatures summoned via this
					path are too complex to exist for long. Within a week
					after their conjuration, the simulacra vanish into insubstantiality.`,
				},
			},
		},
	},
	"thePathofCorruption": {
		Name: "The Path of Corruption",
		Description: `The origins of this path are hotly debated among
		those who are familiar with its intricacies. One theory
		holds that its secrets were taught to the Tremere by demons and that use of it brings the practitioner dangerously close to the infernal powers. A second opinion
		has been advanced that the Path of Corruption is a
		holdover from the days when Clan Tremere was still
		mortal. The third theory, and the most disturbing to
		the Tremere, is that the path originated with the Followers of Set, and that knowledge of its workings was
		sold to the Tremere for an unspecified price. This last
		rumor is vehemently denied by the Tremere, which automatically makes it a favorite topic of discussion when
		the matter comes up.
		The Path of Corruption is primarily a mentally and
		spiritually oriented path centered on influencing the
		psyches of other individuals. It can be used neither to
		issue commands like Dominate nor to change emotions in the moment like Presence. Rather, it produces
		a gradual and subtle twisting of the subject’s actions,
		morals, and thought processes. This path deals intimately with deception and dark desires, and those who
		work through it must understand the hidden places
		of the heart. Accordingly, no character may have a
		higher rating in the Path of Corruption than he has in
		Subterfuge.`,
		Abilities: map[string]disciplineAbility{
			"Contradict": {
				BaseDesc: `The vampire can interrupt a subject’s thought pro-
				cesses, forcing the victim to reverse his current course
				of action. An Archon may be caused to execute a pris-
				oner she was about to exonerate and release; a mortal
				lover might switch from gentle and caring to sadistic
				and demanding in the middle of an encounter. The re-
				sults of Contradict are never precisely known to the
				thaumaturge in advance, but they always take the form
				of a more negative action than the subject had origi-
				nally intended to perform.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `This power may be used on any subject
					within the character’s line of sight. The player rolls as
					per normal. The target rolls Perception + Subterfuge
					(difficulty equal to the number of successes scored by
					the caster + 2). Two successes allow the subject to realize that she is being influenced by some outside source.
					Three successes let her pinpoint the source of the effect. Four successes give her a moment of hesitation,
					neither performing her original action nor its inverse,
					while five allow her to carry through with the original
					action.
					The Storyteller dictates what the subject’s precise
					reaction to this power is. Contradict cannot be used in
					combat or to affect other actions (at the Storyteller’s
					discretion) that are mainly physical and reflexive.`,
				},
			},
			"Subvert": {
				BaseDesc: `This power follows the same principle as does Contradict, the release of a subject’s dark, self-destructive
				side. However, Subvert’s effects are longer-lasting than
				the momentary flare of Contradiction. Under the influence of this power, victims act on their own suppressed temptations, pursuing agendas that their morals or self-control would forbid them to follow under
				normal circumstances.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `This power requires the character to make
					eye contact with the intended victim. The
					player rolls normally. The target resists with a roll of
					Perception + Subterfuge (difficulty equal to the target’s
					Manipulation + Subterfuge). If the thaumaturge scores
					more successes, the victim becomes inclined to follow
					a repressed, shameful desire for the length of time de-
					scribed below. The Storyteller determines the precise desire or
					agenda that the victim follows. It should be in keep-
					ing with the Psychological Flaws that she possesses or
					with the negative aspects of her Nature (for example,
					a Loner desiring isolation to such an extent that she
					becomes violent if forced to attend a social function).
					The subject should not become fixated on following
					this new agenda at all times, but should occasionally be
					forced to spend a Willpower point if the opportunity to
					succumb arises and she wishes to resist the impulse`,
					`1 success`:   `Five minutes`,
					`2 successes`: `One hour`,
					`3 successes`: `One night`,
					`4 successes`: `Three nights`,
					`5 successes`: `One week`,
				},
			},
			"Dissociate": {
				BaseDesc: `“Divide and conquer” is a maxim that is well-understood by the Tremere, and Dissociate is a powerful tool
				with which to divide the Clan’s enemies. This power
				is used to break the social ties of interpersonal relationships. Even the most passionate affair or the oldest
				friendship can be cooled through use of Dissociate, and
				weaker personal ties can be destroyed altogether.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The character must touch the target. The
					player rolls normally. The target resists with a Willpower roll (difficulty of the thaumaturge’s Manipulation + Empathy). The victim loses three dice from
					all Social rolls for a period of time determined by the
					number of successes gained by the caster. This penalty applies to all rolls that rely on Social
					Attributes, even those required for the use of Disciplines. If this power is used on a character who has
					participated in the Vaulderie or a similar ritual, that
					character’s Vinculum ratings are reduced by three for
					the duration of Dissociate’s effect.
					Dissociate’s primary effect falls under roleplaying
					rather than game mechanics. Victims of this power
					should be played as withdrawn, suspicious, and emotionally distant. The Storyteller should feel free to require a Willpower point expenditure for a player who
					does not follow these guidelines.`,
					`1 success`:   `Five minutes`,
					`2 successes`: `One hour`,
					`3 successes`: `One night`,
					`4 successes`: `Three nights`,
					`5 successes`: `One week`,
				},
			},
			"Addiction": {
				BaseDesc: `This power is a much stronger and more potentially
				damaging form of Subvert. Addiction creates just that
				in the victim. By simply exposing the target to a par-
				ticular sensation, substance, or action, the caster cre-
				ates a powerful psychological dependence. Many thau-
				maturges ensure that their victims become addicted to
				substances or thrills that only the mystic can provide,
				thus creating both a source of income and potential
				blackmail material.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The subject must encounter or be exposed
					to the sensation, substance, or action to which the
					character wants to addict him. The thaumaturge then
					touches his target. The player rolls normally; the victim resists with a Self-Control/Instinct roll (difficulty
					equal to the number of successes scored by the caster
					+ 3). Failure gives the subject an instant addiction to
					that object.
					An addicted character must get his fix at least once
					a night. Every night that he goes without satisfying his
					desire imposes a cumulative penalty of one die on all of
					his dice pools (to a minimum pool of one die). The victim must roll Self-Control/Instinct (difficulty 8) every
					time he is confronted with the object of his addiction
					and wishes to keep from indulging. Addiction lasts for
					a number of weeks equal to the thaumaturge’s Manipulation score.
					An individual may try to break the effects of Addiction. This requires an extended Self-Control/Instinct
					roll (difficulty of the caster’s Manipulation + Subterfuge), with one roll made per night. The addict must
					accumulate a number of successes equal to three times
					the number of successes scored by the caster. The victim may not indulge in his addiction over the time
					needed to accumulate these successes. If he does so, all
					accumulated successes are lost and he must begin anew
					on the next night. Note that the Self-Control/Instinct
					dice pool is reduced every night that the victim goes
					without feeding his addiction.`,
				},
			},
			"Dependence": {
				BaseDesc: `Many former pawns of Clan Tremere claim to have
				felt a strange sensation similar to depression when not
				in the presence of their masters. This is usually attributed to the blood bond, but is sometimes the result of
				the vampire’s mastery of Dependence. The final power
				of the Path of Corruption enables the vampire to tie
				her victim’s soul to her own, engendering feelings of
				lethargy and helplessness when the victim is not in her
				presence or acting to further her desires.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The character engages the target in conversation. The player rolls normally. The victim rolls
					Self-Control/Instinct (difficulty equals the number of
					successes scored by the caster + 3). Failure means that
					the victim’s psyche has been subtly bonded to that of
					the thaumaturge for one night per success rolled by the
					caster.
					A bonded victim is no less likely to attack his controller, and feels no particular positive emotions toward her. However, he is psychologically addicted to
					her presence, and suffers a one-die penalty to all rolls
					when he is not around her or performing tasks for her.
					Additionally, he is much less resistant to her commands, and his dice pools are halved when he attempts
					to resist her Dominate, Presence (or other mental or
					emotional control powers), or mundane Social rolls.
					Finally, he is unable to regain Willpower when he is
					not in the thaumaturge’s presence.`,
				},
			},
		},
	},
	"thePathofMars": {
		Name: "The Path of Mars",
		Description: `Those rare Sabbat who have retained Thaumaturgical
		talents have turned their focus to the assistance of the
		Sect in times of war. This path has proven useful, turning the tides of several confrontations with elder vampires. The path adopts a very martial stance, whereas
		other blood magics tend to have subtler, less violent effects. It is rumored that some Camarilla Tremere have
		learned this path, but very few of them have the right
		temperament to wield this path effectively.`,
		Abilities: map[string]disciplineAbility{
			"War Cry": {
				BaseDesc: `A vampire on the attack can focus his will, making
				him less susceptible to battle fear or the powers of the
				undead. The vampire shouts a primal scream to start the
				effect, though some thaumaturges have been known to
				paint their faces or cut themselves open instead.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `For the duration of one scene, the vampire
					adds one to his Courage Trait. Additionally, for the
					purposes of hostile effects, his Willpower is considered
					to be one higher (though this bonus applies only to the
					Trait itself, not the Willpower pool). A character may
					only gain the benefits of War Cry once per scene.`,
				},
			},
			"Strike True": {
				BaseDesc: `The vampire makes a single attack, guided by the
				unholy power of her Blood. This attack strikes its foe
				infallibly.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `By invoking this power, the player need
					not roll to see if the vampire’s attack hits — it does,
					automatically. Only Melee or Brawl attacks may be
					made in this manner. These attacks are considered to
					be one-success attacks; they offer no additional damage dice. Also, they may be dodged, blocked, or parried
					normally, and the defender needs only one success (as
					the attacks’ number of success is assumed to be one).
					Strike True has no effect if attempted on multiple attacks (dice pool splits) in a single turn from one character.`,
				},
			},
			"Wind Dance": {
				BaseDesc: `The thaumaturge invokes the power of the winds,
				moving in a blur. She gains a preternatural edge in
				avoiding her enemies’ blows, moving out of their way
				before the enemy has a chance to throw them.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The player can dodge any number of attacks
					with her full dice pool in a single turn. This advantage
					applies only to dodges — if the character wishes to attack and dodge, the player must still split her dice pool.
					This power lasts for one scene.`,
				},
			},
			"Fearless Heart": {
				BaseDesc: `The vampire temporarily augments his abilities as a
				warrior. Through the mystical powers of blood magic,
				the character becomes a potent fighting force`,
				Lvl: 4,
				System: map[string]string{
					"sys": `Fearless Heart grants the vampire an extra
					point in each of the Physical Attributes (Strength,
					Dexterity, and Stamina). These Traits may not exceed
					their generational maximums, though the player may
					use blood points to push the character’s Traits even
					higher. The effects last for one scene, and a character
					may gain its benefits only once per scene. The vampire
					must spend two hours in a calm and restful state following the use of Fearless Heart, or lose a blood point
					every 15 minutes until he rests.`,
				},
			},
			"Comrades at Arms": {
				BaseDesc: `This ability extends the power of the previous abilities in the path. It allows any of the earlier effects to be
				applied to a group such as a pack or War Party.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The player chooses one of the lower-level
					powers in the path, invoking it as normal. Afterward,
					he touches another character and (if the roll for Comrades at Arms is successful) bestows the benefit on her
					as well. The same power may be delivered to any number of packmates, as long as the rolls for Comrades at
					Arms are successful and the thaumaturge pays the appropriate blood costs.`,
				},
			},
		},
	},
	"thePathofTechnomancy": {
		Name: "The Path of Technomancy",
		Description: `The newest path to be accepted by the Tremere
		hierarchy as part of the Clan’s official body of knowledge, the Path of Technomancy is a relatively recent
		innovation, developed in the latter half of the 20th
		century. The path focuses on the control of electronic
		devices, from cellphones to laptops, and its proponents
		maintain that it is a prime example of the versatility of
		Thaumaturgy with regards to a changing world. More
		conservative Tremere, however, state that mixing
		Tremere magic with mortal science borders on treason
		or even blasphemy, and some European Regents have
		gone so far as to declare knowledge of Technomancy
		grounds for expulsion from their chantries. The Inner
		Council did approve the introduction of the path into
		the Clan’s grimoires, but has yet to voice any opinion
		on the conservative opposition to Technomancy.`,
		Abilities: map[string]disciplineAbility{
			"Analyze": {
				BaseDesc: `Mortals are constantly developing new innovations,
				and any vampire who would work Technomancy must
				be able to understand that upon which he practices his
				magic. The most basic power of this path allows the
				thaumaturge to project his perceptions into a device,
				granting him a temporary understanding of its purpose,
				the principles of its functioning, and its means of operation. This does not grant permanent knowledge,
				only a momentary flash of insight which fades within
				minutes.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `A character must touch the device in order
					to apply this power. The number of successes rolled determines how well the character understands this particular piece of equipment. One success allows a basic
					knowledge (on/off and simple functions), while three
					successes grant competence in operating the device,
					and five successes show the character the full range of
					the device’s potential. The knowledge lasts for a number of minutes equal to the character’s Intelligence.
					This power can also be used to understand a nonphysical technological innovation — generally a piece
					of software — at +2 difficulty. The character must
					touch the computer on which the software is installed
					— simply holding the flash drive or CD-ROM is not
					enough. Software applied remotely to a device (such
					as through an app store) also cannot be analyzed until
					it is installed.`,
				},
			},
			"Burnout": {
				BaseDesc: `It is usually easier to destroy than to create, and sensitive electronics are no exception to this rule. Burnout
				is used to cause a device’s power supply (either internal or external) to surge, damaging or destroying the
				target. Burnout cannot be used to directly injure another individual, although the sudden destruction of
				a pacemaker or a car’s fuel injection control chip can
				certainly create a health hazard. Large enough systems, such as a server cluster or
				a passenger aircraft, impose a +2 to +4 difficulty (at
				Storyteller discretion) to affect with this power. Additionally, some systems, such as military and banking
				networks, may be protected against power surges and
				spikes, and thus possess one to five dice (Storyteller
				discretion again) to roll to resist this power. Each success on this roll (difficulty 6) takes away one success
				from the Thaumaturgy roll.
				Burnout may be used to destroy electronic data storage, in which case three successes destroy all information on the target item, and five erase it beyond any
				hope of non-magical recovery.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `A character can use this power at a range
					of up to 10 times her Willpower in yards or meters,
					although a +1 difficulty is applied if she is not touching
					the target item. The number of successes determines
					the extent of the damage`,
					`1 success`: `Momentary interruption of operation
					(one turn), but no permanent
					damage.`,
					`2 successes`: `Significant loss of function; +1
					difficulty to use using the device for
					the rest of the scene.`,
					`3 successes`: `The device breaks and is inoperable
					until repaired.`,
					`4 successes`: `Even after repairs, the device’s
					capabilities are diminished
					(permanent +1 difficulty to use).`,
					`5 successes`: `The equipment is a total write-off;
					completely unsalvageable.`,
				},
			},
			"Encrypt/Decrypt": {
				BaseDesc: `Electronic security is a paramount concern of governments and corporations alike. Those thaumaturges
				who are techno-savvy enough to understand the issues
				at stake have become quite enamored of this power,
				which allows them to scramble a device’s controls mystically, making it inaccessible to anyone else. Encrypt/
				Decrypt also works on electronic media; a DVD under
				the influence of this power displays just snow and static
				if played back without the owner’s approval. Some neonates have taken to calling this power “DRM.”`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The character touches the device or data
					container that he wishes to encrypt. The player rolls
					normally. The number of successes scored is applied as
					a difficulty modifier for anyone who attempts to use the
					protected equipment or access the scrambled information without the assistance of the character. The caster
					can dispel the effect at any time by touching the target
					item and spending a point of Willpower.
					This power may also be used to counter another
					thaumaturge’s use of Encrypt/Decrypt. The player rolls
					at +1 difficulty; each success negates one of the “owner’s.”
					The effects of Encrypt/Decrypt last for a number of
					weeks equal to the character’s permanent Willpower
					rating.`,
				},
			},
			"Remote Access": {
				BaseDesc: `With this power, a skilled thaumaturge can bypass
				the need for physical contact to operate a device. This
				is not a form of telekinesis; the vampire does not manipulate the item’s controls, but rather touches it directly with the power of his mind.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `This power may be used on any electronic
					device within the character’s line of sight. The number of successes rolled is the maximum number of dice
					from any relevant Ability that the character may use
					while remotely controlling the device. (For instance, if
					Fritz has Technology 5 and scores three successes while
					using Remote Access on a keypad lock, he can only
					apply three dots of his Technology rating to any rolls
					that he makes through any use of the power.) Remote
					Access lasts for a number of turns equal to the number
					of successes rolled, and can only be used on one item
					at a time.
					If an item is destroyed while under the effects of Re-
					mote Access, the character takes five dice of bashing
					damage due to the shock of having his perceptions
					rudely shunted back into his own body.`,
				},
			},
			"Telecommute": {
				BaseDesc: `A progressive derivation of Remote Access, Telecommute allows a thaumaturge to project her consciousness into the Internet, sending her mind through
				network connections as fast as they can transfer her.
				While immersed in the network, she can use any other
				Technomancy power on the devices with which she
				makes contact.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The character touches any form of communications device: a cellphone, 3G-equipped netbook,
					Wi-Fi tablet, or anything else that is connected directly
					or indirectly to the Internet. The player rolls normally
					and spends a Willpower point. Telecommute lasts for
					five minutes per success rolled, and may be extended by
					10 minutes with the expenditure of another Willpower
					point. The number of successes indicates the maximum
					range that the character can project her consciousness
					away from her body. While in the network, the character can apply any
					other Path of Technomancy power to any device or
					data with which she comes in contact. A loss of connection, either through the destruction of a part of the
					network or simply a loss of cell signal, hurls her consciousness back to her body and inflicts eight dice of
					bashing damage.
					A character traveling through the Internet by means
					of this power can use her Path of Technomancy powers
					at a normal difficulty. Using any other abilities or powers while engaged thus is done at a +2 difficulty.`,
					`1 success`:   `25 miles/40 kilometers`,
					`2 successes`: `250 miles/400 kilometers`,
					`3 successes`: `1000 miles/1500 kilometers`,
					`4 successes`: `5000 miles/8000 kilometers`,
					`5 successes`: `Anywhere in the world`,
				},
			},
		},
	},
	"thePathoftheFathersVengeance": {
		Name: "The Path of the Father’s Vengeance",
		Description: `This path, based loosely on a powerful thaumaturge’s
		interpretations of the Book of Nod, devotes itself to delivering justice to the race of Cainites. Each power sup-
		posedly has some precedent in the parables of the ancient book, and focuses on teaching the lessons of Caine
		via the power of blood magic. Use of this path is hotly
		debated in the Sabbat, as some consider it tantamount
		to claiming to hold Caine’s right over all vampires oneself. Camarilla vampires don’t have the same knowledge
		of the Book of Nod that the Sabbat do, but the path is
		not entirely unheard of in Tremere chantries.
		The power of this path comes not only from the
		magic of blood, but also incantation of verses from the
		Book of Nod. For any of these powers to take effect,
		the caster must speak the actual condemnation. For example, to invoke the third-level power, the caster must
		state plainly to his target that she may eat only ashes.
		The subject must generally be able to hear the caster
		for these powers to take effect, though writing them
		and showing them to the subject will do.
		These powers apply to vampires only. They do not affect mortals, ghouls, or any other supernatural creatures.`,
		Abilities: map[string]disciplineAbility{
			"Zillah’s Litany": {
				BaseDesc: `Zillah, the wife of Caine, unknowingly drank from
				her husband and sire three times, thus becoming bonded to him. This power reveals existing blood bonds and
				Vinculi to the thaumaturge.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `If the subject has any blood bonds or Vinculi to other vampires, this power reveals them to the
					caster. Although the caster may not know the vampires in question, this power does reveal the names and
					gives rough psychic impressions of the individuals in
					question.`,
				},
			},
			"The Crone’s Pride": {
				BaseDesc: `This power inflicts the curse of the crone, who bound
				Caine to her as he fled his wife’s spurning. Hideously
				ugly, the crone had to resort to trickery to get others to
				help or serve her.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `This power reduces the target’s Appearance
					to zero. All Social rolls during this time generally fail,
					unless the character attempts to intimidate or browbeat the subject. This power lasts for one night.`,
				},
			},
			"Feast of Ashes": {
				BaseDesc: `Primarily used against wanton or excessive vampires,
				this power temporarily removes a vampire’s dependency on blood. While some would say this negates the
				Curse of Caine, it reduces the vampire to little more
				than a wretched scavenger, as he must consume literal
				ashes, though he gains little sustenance from them.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The victim of this power can no longer consume blood, vomiting it up as he would mortal food or
					drink. Instead, the victim can eat only ashes, and the
					“blood points” he gains from this may be used only to
					rise each night. Ashen “blood points” may not be used
					to power Disciplines, raise Attributes, or feed ghouls
					(though actual blood points in the character’s body at
					the time this power is invoked may still be used for
					such). One blood point’s worth of ash is roughly one
					pint or half-liter, and any ash will do — cigarette ash,
					campfire leftovers, or vampire corpses destroyed by fire
					or sunlight. This power lasts for one week.`,
				},
			},
			"Uriel’s Disfavor": {
				BaseDesc: `This power invokes the darkness of the Angel of
				Death. All but the dimmest of light causes the subject
				excruciating pain, and some artificial forms of bright
				light may even damage the vampire. Uriel delivered
				God’s curse on Caine, shielding him in the blackness
				of his wings.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The presence of any light makes the subject
					uncomfortable, and bright light of any kind — flashlights, headlights, etc. — inflict one health level of aggravated damage on the character for every turn he remains under its direct focus. Most vampires who suffer
					this curse elect to sleep for the duration, hiding away in
					the darkness of their havens until they can walk again
					among the living without pain. This power lasts for
					one week.`,
				},
			},
			"Valediction": {
				BaseDesc: `Many Sabbat rightfully fear this power, though very
				few have ever seen it used. It levies a punishment for
				breaking one of Caine’s greatest commandments —
				the ban against diablerie. As most Sabbat attain their
				power and station through some measure of diablerie,
				they must reconcile their beliefs with the admonitions
				of Caine, and this power engenders a great sense of humility.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `When this power takes effect, the subject
					immediately reverts to her original Generation. This
					change may entail losing points in certain Traits due to
					generational maximums. This power lasts for one week,
					after which any Traits reduced to higher-Generation
					maximums return to normal. It takes three turns to speak
					the full verse that implements this power’s effects.`,
				},
			},
		},
	},
	"thaumaturgicalRituals": {
		Name: "The Path of Blood",
		Description: `Rituals are Thaumaturgical formulas, meticulously
		researched and prepared, that create powerful magical
		effects. Rituals are less versatile than paths, as their effects are singular and straightforward, but they are bet-
		ter suited to specific ends.
		All thaumaturges have the ability to use rituals,
		though each individual ritual must be learned separately. By acquainting herself with the arcane practice
		of blood magic, the caster gains the capacity to manipulate these focused effects. Casting rituals requires a successful Intelligence
		+ Occult roll, for which the difficulty equals 3 + the
		level of the ritual (maximum 9). Only one success is
		required for a ritual to work, though certain spells may
		require more successes or have variable effects based
		on how well the caster’s roll goes. Should a roll to activate a ritual fail, the Storyteller is encouraged to cre-
		ate strange occurrences or side effects, or even make it
		appear that the ritual was successful, only to reveal its
		failure at a later time. A botched ritual roll may even
		indicate a catastrophic failure or summon an ill-tem-
		pered demon.
		Rituals sometimes require special ingredients or reagents to work — these are noted in each ritual’s description. Common components include herbs, animal
		bones, ceremonial items, feathers, eye of newt, tongue
		of toad, etc. Acquiring magical components for a powerful ritual may form the basis for an entire story.
		At the first level of Thaumaturgy, the vampire automatically gains a single Level One ritual. To learn
		further rituals, the thaumaturge must find someone
		to teach him, or learn the ritual from a scroll, tome,
		or other archive. Learning a new ritual can take anywhere from a few nights (Level One ritual) to months
		or years (Level Five ritual). Some mystics have studied
		individual rituals for decades, or even centuries.`,
		Abilities: map[string]disciplineAbility{
			"Bind the Accusing Tongue": {
				BaseDesc: `This ancient ritual is said to have been one of the
				first developed by the Tremere and a primary reason
				for the lack of cohesive opposition to their expansion.
				Bind the Accusing Tongue lays a compulsion upon
				the subject that prevents him from speaking ill of the
				caster, allowing the thaumaturge to commit literally
				unspeakable acts without fear of reprisal.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The caster must have a picture or other image or effigy of the ritual’s target, a lock of the target‘s
					hair, and a black silken cord. The caster winds the cord
					around the hair and image while intoning the ritual’s
					vocal component. Once the ritual is complete, the target must score more successes on a Willpower roll (difficulty of the caster‘s Thaumaturgy rating + 3) than the
					caster scored in order to say anything negative about
					the caster. The ritual lasts until the target succeeds at
					this roll or the silk cord is unwound, at which point the
					image and the lock of hair crumble to dust.`,
				},
			},
			"Blood Rush": {
				BaseDesc: `This ritual allows the vampire to create the sensation
				of drinking blood in himself without actually feeding.
				The ritual can be used for pleasure, but it is more often used to prevent frenzy when confronted with fresh
				blood. The vampire must carry the fang of a predatory
				animal on his person for this ritual to work`,
				Lvl: 1,
				System: map[string]string{
					"sys": `Performance of the ritual results in the
					Beast being kept in check automatically. Blood Rush
					allows the vampire to resist hunger-based frenzy for up
					to one hour, at which point the Cainite feels hungry
					again (assuming he did before). This ritual takes only
					one turn to enact.`,
				},
			},
			"Communicate with Kindred Sire": {
				BaseDesc: `By enacting this ritual, the caster may join minds
				with her sire, speaking telepathically with him over
				any distance. The communication may continue until
				the ritual expires or until either party ends the conversation. The caster must possess an item once owned by
				her sire for the ritual to work.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The caster must meditate for 30 minutes
					to create the connection. Conversation may be maintained for 10 minutes per success on the activation
					roll.`,
				},
			},
			"Defense of the Sacred Haven": {
				BaseDesc: `This ritual prevents sunlight from entering an area
				within 20 feet (six meters) of this ritual’s casting. A
				mystical darkness blankets the area, keeping the bale-
				ful light at bay. Sunlight reflects off windows or magically fails to pass through doors or other portals. To
				invoke this ritual’s protection, the caster draws sigils in
				her own blood on all the affected windows and doors.
				The ritual lasts as long as the thaumaturge stays within
				the 20-foot (6-meter) radius.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `This ritual requires one hour to perform,
					during which the caster recites incantations and in-
					scribes glyphs. One blood point is required for this
					ritual to work.`,
				},
			},
			"Deflection of Wooden Doom": {
				BaseDesc: `This ritual protects the caster from being staked,
				whether she is resting or active. While this ritual is in
				effect, the first stake that would pierce the vampire’s
				heart disintegrates in the attacker’s hand. A stake
				merely held near the caster is unaffected; for this ritual
				to work, the stake must actively be used in an attempt
				to impale the vampire.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The caster must surround herself with a circle of wood for a full hour. Any wood will work: furniture, sawdust, raw timber, 2’ x 4’s, whatever. The circle
					must remain unbroken, however. At the end of the
					hour, the vampire places a wooden splinter under her
					tongue. If this splinter is removed, the ritual is nullified. This ritual lasts until the following dawn or dusk.`,
				},
			},
			"Devil’s Touch": {
				BaseDesc: `Thaumaturges use this ritual to place curses upon
				mortals who earn their ire. Using this ritual marks an
				individual invisibly, causing all those who come in
				contact with him to receive him poorly. The mortal is
				treated as the most loathsome individual conceivable,
				and all who deal with him do everything in their power
				to make him miserable. Even bums spit at an afflicted
				individual, and children taunt him and barrage him
				with vulgarities.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The effects of this ritual last one night, disappearing as the sun rises. The mortal (it doesn’t work
					on vampires) must be present when the ritual is invoked, and a penny must be placed somewhere on his
					person (in a pocket, shoe, etc.).`,
				},
			},
			"Domino of Life": {
				BaseDesc: `A vampire wanting or needing to simulate a human
				Characteristic can do so once Domino of Life is cast. For
				one entire night, the vampire can eat, breathe, main-
				tain a normal body temperature, assume a human flesh
				tone, or display some other single trait of humankind
				she desires. Note that only one trait can be replicated
				in this fashion. The vampire must have a vial of fresh
				human blood on his person to maintain this ritual.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `Using this ritual adds one die to the caster’s
					dice pools when attempting to pass as human. Unless onlookers are especially wary, the Domino of Life
					should fool them into thinking the caster is mortal —
					not that they should have any reason to suspect otherwise.`,
				},
			},
			"Engaging the Vessel of Transference": {
				BaseDesc: `This ritual enchants a container to fill itself with
				blood from any living or unliving being who holds it,
				replacing the volume of blood taken with an equal
				amount previously held inside the container. When
				the ritual is enacted, the vessel (which must be between
				the size of a small cup and a one-gallon/four-liter jug) is
				sealed full of the caster’s blood and inscribed with the
				Hermetic sigil which empowers the ritual. Whenever
				an individual touches the container with his bare skin,
				he feels a slight chill against his flesh but no further
				discomfort. The container continues to exchange the
				blood it contains until it is opened. The two most common uses of this ritual are to covertly create a blood
				bond and to obtain a sample of a subject’s blood for
				ritual or experimental purposes.`,
				Lvl: 1,
				System: map[string]string{
					"sys": ` This ritual takes three hours to enact (reduced by 15 minutes for each success on the casting
						roll) and requires one blood point (although not necessarily the caster’s blood), which is sealed inside the
						container. The ritual only switches blood between itself and a subject if it is touched with bare skin — even
						thin cotton gloves keep it from activating.
						Individuals with at least four dots in Occult recog-
						nize the Hermetic sigil with two successes on an Intelligence + Occult roll (difficulty 8).`,
				},
			},
			"Illuminate the Trail of Prey": {
				BaseDesc: `This ritual causes the path of the subject’s passing
				to glow in a manner that only the vampire can see.
				The tracks shine distinctly, but only to the eyes of the
				caster. Even airplane trajectories and animal tracks
				shine with unhealthy light. The ritual is nullified if the
				target wades through or immerses himself in water, or
				if he reaches the destination of his journey. The caster
				must burn a length of white satin ribbon that has been
				in her possession for at least 24 hours for this ritual to
				take effect.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The thaumaturge must have a mental picture of or know the name of her prey. The individual’s
					wake glows with a level of brightness dependent on
					how long it has been since he passed that way — old
					tracks burn less brightly, while fresh tracks blaze`,
				},
			},
			"Incantation of the Shepherd": {
				BaseDesc: `This ritual enables the caster to mystically locate all
				members of his herd. While intoning the ritual’s vocal
				component, he spins in a slow circle with a glass object
				of some sort held to each of his eyes. At the end of the
				ritual, he has a subliminal sense of the direction and
				distance to each of his regular vessels.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `This ritual gives the character the location
					(relative to him) of every member of his Herd. If he
					does not have the Herd Background, Incantation of the
					Shepherd locates the closest three mortals from whom
					the caster has fed at least three times each. This ritual
					has a maximum range of 10 miles or 15 kilometers times
					the character’s Herd Background, or five miles (eight
					kilometers) if he has no points in that Background.`,
				},
			},
			"Purity of Flesh": {
				BaseDesc: `The caster cleanses her body of all foreign material with this ritual. To perform it, she meditates on
				bare earth or stone while surrounded by a circle of 13
				sharp stones. Over the course of the ritual, the caster
				is slowly purged of all physical impurities: dirt, alcohol,
				drugs, poison, bullets lodged in the flesh, and tattoo
				ink are equally affected, slowly rising to the surface of
				the caster’s skin and flaking away as a gritty gray film
				that settles within the circle. Any jewelry, makeup, or
				clothes that the caster is wearing are also dissolved`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The player spends one blood point before
					rolling. Purity of Flesh removes all physical items from
					the caster’s body, but does not remove enchantments,
					mind control, or diseases of the blood.`,
				},
			},
			"Wake with Evening’s Freshness": {
				BaseDesc: `This ritual allows a vampire to awaken at any sign
				of danger, especially during the day. If any potentially
				harmful circumstances arise, the caster immediately
				rises, ready to face the problem. This ritual requires the
				ashes of burned feathers to be spread over the area in
				which the Kindred wishes to sleep.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `This ritual must be performed immediately
					before the vampire goes to sleep for the day. Any interruption to the ceremonial casting renders the ritual ineffective. If danger arises, the caster awakens and may
					ignore the Humanity/Path dice pool limit rule for the
					first two turns of consciousness. Thereafter, the penalty
					takes effect, but the thaumaturge will have already risen and will be able to address problematic situations`,
				},
			},
			"Widow’s Spite": {
				BaseDesc: `This ritual causes a pain, itch, or other significant
				(but not deadly) sensation in the subject. Similar in
				effect to legendary “voodoo doll” effects, this ritual is
				used more out of scorn or malice than actual enmity. In
				fact, it requires a wax or cloth doll that resembles the
				target, which bleeds when the power takes effect.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The ceremonial doll must resemble, however rudely, the victim of the ritual. It produces no mechanical effect, other than a simple physical stimulus.
					The caster may determine where on the subject’s body
					the pain or itch appears.`,
				},
			},
			"Blood Walk": {
				BaseDesc: `Developed during Clan Tremere’s troubled inception, Burning Blade allows a caster to temporarily enchant a melee weapon to inflict unhealable wounds on
				supernatural creatures. While this ritual is in effect, the
				weapon flickers with an unholy greenish flame.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `This ritual can only be cast on melee weapons. The caster must cut the palm of her weapon hand
					during the ritual — with the weapon if it is edged, otherwise with a sharp stone. This inflicts a single health
					level of lethal damage, which cannot be soaked but
					may be healed normally. The player spends three blood
					points, which are absorbed by the weapon. Once the
					ritual is cast, the weapon inflicts aggravated damage on
					all supernatural creatures for the next few successful attacks, one per success rolled. Multiple castings of Burning Blade cannot be “stacked” for longer durations.
					Furthermore, the wielder of the weapon cannot choose
					to do normal damage and “save up” aggravated strikes
					— each successful attack uses one aggravated strike
					until there are none left, at which point the weapon
					reverts to inflicting normal damage.`,
				},
			},
			"Donning the Mask of Shadows": {
				BaseDesc: `This ritual renders its subject translucent; her form
				appears dark and smoky, and the sounds of her footsteps are muffled. While it does not create true invisibility, the Mask of Shadows makes the subject much
				less likely to be detected by sight or hearing.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `This ritual may be simultaneously cast on a
					number of subjects equal to the caster’s Occult rating;
					each individual past the first adds five minutes to the
					base casting time. Individuals under the Mask of Shadows can only be detected if the observer possesses a
					power (such as Auspex) sufficient to penetrate Obfuscate 3. The Mask of Shadows lasts a number of hours
					equal to the number of successes rolled when it is cast
					or until the caster voluntarily lowers it.`,
				},
			},
			"Eyes of the Night Hawk": {
				BaseDesc: `This ritual allows the vampire to see through the
				eyes of a bird, and to hear through its ears. The bird
				chosen must be a predatory bird, and the vampire must
				touch it when initiating this ritual. At the end of this
				ritual, the caster must put out the bird’s eyes, or suffer
				blindness herself.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The vampire is able to mentally control
					where the bird travels for the duration of the ritual.
					The bird will not necessarily perform any other action
					than flight — the caster cannot command it to fight,
					pick up and return an object, or scratch a target. The
					bird returns to the vampire after finishing its flight. If
					the vampire does not put out the bird’s eyes, she suffers
					a three-night period of blindness. This ritual ceases effect at sunrise.`,
				},
			},
			"Machine Blitz": {
				BaseDesc: `Machines go haywire when this ritual is cast. It
				takes effect instantly and lasts as long as the vampire
				concentrates on it. This ritual may be used to kill car
				engines, erase flash drives, drain the battery of a cellphone, stop life-support machines, et cetera. Essentially, Machine Blitz stops any machine more complex
				than a rope-and-pulley. The thaumaturge must have a
				scrap of rusted metal in her possession for this ritual to
				work, though some vampires use a variant that requires
				a knot steeped in human saliva to be untied.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `This ritual only stops machines; it does not
					grant any control over them. The effects of this ritual
					are invisible and appear to be coincidental.`,
				},
			},
			"Principal Focus of Vitae Infusion": {
				BaseDesc: `This ritual imbues a quantity of blood within an
				object small enough for the vampire to carry in both
				hands. (The object may not be any larger than this,
				though it may be as small as a dime.) After the ritual
				is conducted, the object takes on a reddish hue and
				becomes slick to the touch. At a mental command, the
				caster may release the object from its enchantment,
				causing it to break down into a pool of blood. This
				blood may serve whatever purpose the vampire desires;
				many thaumaturges wear enchanted baubles to ensure
				they have emergency supplies of vitae`,
				Lvl: 2,
				System: map[string]string{
					"sys": `An object may store only one blood point of
					vitae. If a Kindred wishes to make an infused focus for
					an ally, she may do so, but the blood contained within
					must be her own. (If the ally then drinks the blood, he
					is one step closer to the blood bond). The ally must be
					present at the creation of the focus.`,
				},
			},
			"Recure of the Homeland": {
				BaseDesc: `The vampire calls on the power of the earth to heal
				grave wounds she may have received. The thaumaturge
				must use at least a handful of dirt from the city or town
				of her mortal birth and recite a litany of her mortal
				family tree as she casts this ritual.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The Cainite must mix the earth with two
					points of her own blood to make a healing paste. One
					handful will heal one aggravated wound, and only one
					handful can be used per night. This ritual can only be
					used on the vampire who knows it.`,
				},
			},
			"Ward Versus Ghouls": {
				BaseDesc: `Wary Tremere created this ritual to protect themselves from the minions of vengeful rivals. By invoking
				this ritual, the caster creates a glyph that causes great
				pain to any ghouls who come in contact with it. The
				Kindred pours a point’s worth of blood over the object he wishes to ward (a piece of parchment, a coin,
				a doorknob, etc.), and recites the incantation, which
				takes 10 minutes. In 10 hours, the magical ward is complete, and will inflict excruciating pain on any ghoul
				unfortunate enough to touch the warded object.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `Ghouls who touch warded objects suffer
					three dice of lethal damage. This damage occurs again
					if the ghoul touches the object further; indeed, a ghoul
					who consciously wishes to touch a warded object must
					spend a point of Willpower to do so. This ritual wards only one object — if inscribed on
					the side of a car, the ward affects only that door or
					fender, not the whole car. Wards may be placed on
					weapons, even bullets, though this usually works best
					on small-caliber weapons. Bullets often warp upon firing, however, and for a ward to remain intact on a fired
					round, the player needs five successes on the Firearms
					roll.`,
				},
			},
			"Warding Circle versus Ghouls": {
				BaseDesc: `This ritual is enacted in a manner similar to that of
				Ward versus Ghouls, but creates a circle centered on
				the caster into which a ghoul cannot pass without being burned. The circle can be made as large and as permanent as the caster desires, as long as she is willing to
				pay the necessary price. Many Tremere chantries and
				havens are protected by this and other Warding Circle
				rituals.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `The ritual requires three blood points of
					mortal blood. The caster determines the size of the
					warding circle when it is cast; the default radius is 10
					feet/3 meters, and every 10-foot/3-meter increase raises
					the difficulty by one, to a maximum of 9 (one additional success is required for every increase past the number
					necessary to raise the difficulty to 9). The player spends
					one blood point for every 10 feet/3 meters of radius and
					rolls. The ritual takes the normal casting time if it is to
					be short-term (lasting for the rest of the night) or one
					night if it is to be long-term (lasting a year and a day).
					Once the warding circle is established, any ghoul
					who attempts to cross its boundary feels a tingle on his skin and a slight breeze on his face — a successful In-
					telligence + Occult roll (difficulty 8) identifies this as
					a warding circle. If the ghoul attempts to press on, he
					must roll more successes on a Willpower roll (difficulty
					of the caster’s Thaumaturgy rating + 3) than the caster
					rolled when establishing the ward. Failure indicates
					that the ward blocks his passage and inflicts three dice
					of bashing damage, and his next roll to attempt to enter the circle is at +1 difficulty. If the ghoul leaves the
					circle and attempts to enter it again, he must repeat
					the roll. Attempts to leave the circle are not blocked.`,
				},
			},
			"Clinging of the Insect": {
				BaseDesc: `This ritual allows the caster to cling to walls or ceilings, as would a spider. She may even crawl along these
				surfaces (as long as they can support her). Use of this
				power seriously discomfits mortal onlookers. The character must place a live spider under her tongue for the
				duration of the ritual (though the spider may die while
				in the thaumaturge’s mouth without canceling the
				power).`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The character may move at half her nor-
					mal rate while climbing walls or ceilings. This power
					lasts for one scene, or until the vampire spits out the
					spider.`,
				},
			},
			"Flesh of Fiery Touch": {
				BaseDesc: `This defensive ritual inflicts painful burns on anyone
				who deliberately touches the subject’s skin. It requires
				the subject to swallow a small glowing ember, which
				does put off some vampires with low pain thresholds.
				Some vain thaumaturges use this ritual purely for its
				subsidiary effect of darkening the subject’s skin to a
				healthy sun-bronzed hue.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `Flesh of Fiery Touch takes two hours to cast
					(reduced by 10 minutes per success). It requires a small
					piece of wood, coal, or other common fuel source,
					which ignites and is swallowed at the end of the ritual.
					The subject who swallows the red-hot ember receives
					a single aggravated health level of damage (difficulty 6
					to soak with Fortitude). Until the next sunset, anyone
					who touches the subject’s flesh receives a burn that inflicts a single aggravated health level of damage (again,
					difficulty 6 to soak with Fortitude). The victim must
					voluntarily touch the subject; this damage is not inflicted if the victim is touched or accidentally comes in
					contact with the subject.
					This ritual darkens the subject’s skin to that which
					would be obtained by long-term exposure to the sun in
					a mortal. The tone is slightly unnatural and metallic,
					and is clearly artificial to any observer who succeeds in
					a Perception + Medicine roll (difficulty 8).`,
				},
			},
			"Incorporeal Passage": {
				BaseDesc: `Use of this ritual allows the thaumaturge to make
				herself insubstantial. The caster becomes completely
				immaterial and thus is able to walk through walls, pass
				through closed doors, escape manacles, etc. The caster
				also becomes invulnerable to physical attacks for the
				duration of the ritual. The caster must follow a straight
				path through any physical objects, and may not draw
				back. Thus, a Kindred may walk through a solid wall,
				but may not walk down through the earth (as it would
				be impossible to reach the other side before the ritual
				lapsed). This ritual requires that the caster carry a shard
				from a shattered mirror to hold her image.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `This ritual lasts a number of hours equal to
					the number of successes scored on a Wits + Survival
					roll (difficulty 6). The caster may prematurely end the
					ritual (and, thus, her incorporeality) by turning the
					mirror shard away so that it no longer reflects her image.`,
				},
			},
			"Mirror of Second Sight": {
				BaseDesc: `This object is an oval mirror no less than four inches
				(10 cm) wide and no more than 18 inches (45 cm) in
				length. It looks like a normal mirror, but once created,
				the vampire can use it to see the supernatural: It reflects the true form of Lupines and faeries, and enables
				the owner to see ghosts as they move though the Underworld. The caster creates the mirror by bathing an
				ordinary mirror in a quantity of her own blood while
				reciting a ritual incantation.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The ritual requires one point of the vampire’s blood. Thereafter, the mirror reflects images of
					other supernatural creatures’ true forms — werewolves
					appear in their hulking man-wolf shapes, magi glow in
					a scintillating nimbus, ghosts become visible (in the
					mirror), and so on. Sometimes, the mirror also reveals
					those possessed of True Faith in clouds of golden light.`,
				},
			},
			"Pavis of Foul Presence": {
				BaseDesc: `The Tremere joke privately that this is their “ritual
				for the Ventrue.” Kindred who invoke the Presence
				Discipline on the subject of this ritual find the effects of
				their Discipline reversed, as if they had used the power
				on themselves. For example, a vampire using Presence
				to instill fear in a Kindred under the influence of this
				ritual feels the fear herself; a vampire summoning the
				caster is instead drawn to the thaumaturge’s location.
				This ritual is an unbroken secret among the Tremere,
				and the Warlocks maintain that its use is unknown
				outside their Clan. The magical component for this
				ritual is a length of blue silken cord, which must be
				worn around the caster’s neck.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `This ritual lasts against a number of effects
					equal to the successes rolled, or until the sunrise after it is enacted. Note that the Presence Discipline
					power must actually succeed before being reversed by
					the ritual. As such, only powers that specifically target
					the caster (and thus, require a roll to succeed) can be
					reversed — “passive” powers such as Majesty are not
					affected.`,
				},
			},
			"Sanguine Assistant": {
				BaseDesc: `Thaumaturges often need laboratory assistants whom
				they can trust implicitly. This ritual allows the intrepid
				vampire to conjure a temporary servant. To cast the
				ritual, the thaumaturge slices open his arm and bleeds
				into a specially prepared earthen bowl. The ritual sucks
				in and animates whatever random unimportant items
				the wizard happens to have lying around his workshop
				— glass beakers, dissection tools, pencils, crumpled papers, semiprecious stones — and binds the materials
				together into a small humanoid form animated by the
				power of the ritual and the blood. Oddly enough, this
				ritual almost never takes in any tool that the caster
				finds himself needing during the assistant’s lifespan,
				nor does it take the physical components of any other
				ritual nor any living thing. The servant has no personality to speak of at first, but gradually adopts the
				mannerisms and thought processes that the thaumaturge desires in an ideal servant. Sanguine Assistants
				are temporary creations, but some vampires become
				fond of their tiny accomplices and create the same one
				whenever the need arises.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The player spends five blood points and
					rolls. The servant created by the ritual stands a foot
					(30 cm) high and appears as a roughly humanoid shape
					composed of whatever the ritual sucked in for its own
					use. It lasts for one night per success rolled. At the end
					of the last night, the assistant crawls into the bowl used
					for its creation and falls apart. The assistant can be reanimated through another application of this ritual; if
					the caster so desires, it re-forms from the same materials with the same memories and personality.
					A Sanguine Assistant has Strength and Stamina of
					1, and Dexterity and Mental Attributes equal to those
					of the caster. It begins with no Social Attributes to
					speak of, but gains one dot per night in Charisma and
					Manipulation until its ratings are equal to those of the
					caster. It has all of the caster’s Abilities at one dot lower than his own. A Sanguine Assistant is a naturally
					timid creature and flees if attacked, though it will try to
					defend its master’s life at the cost of its own. It has no
					Disciplines of its own, but has a full understanding of
					all of its master’s Thaumaturgical knowledge and can
					instruct others if so commanded. A Sanguine Assistant
					is impervious to any mind-controlling Disciplines or
					magic, so completely is it bound to its creator’s will`,
				},
			},
			"Shaft of Belated Quiescence": {
				BaseDesc: `This ritual turns an ordinary stake of rowan wood
				into a particularly vicious weapon. When the stake
				penetrates a vampire’s body, the tip breaks off and begins working its way through the victim’s flesh to his
				heart. The trip may take several minutes or several
				nights, depending on where the stake struck. The stake
				eludes attempts to dig it out, burrowing farther into the
				victim’s body to escape surgery. The only Kindred who
				are immune to this internal attack are those who have
				had their hearts removed by Serpentis.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The ritual takes five hours to enact, minus
					30 minutes per success. The stake must be carved of
					rowan wood, coated with three blood points of the
					caster’s blood, and blackened in an oak-wood fire.
					When the ritual is complete, the stake is enchanted to
					act as described above.
					An attack with a Shaft of Belated Quiescence is performed as with a normal stake: a Dexterity + Melee
					roll (difficulty 6, modified as per the normal combat
					rules; the attack does not need to specifically target the
					heart) with a lethal damage rating of Strength + 1. If
					at least one health level of damage is inflicted after the
					target rolls to soak, the tip of the stake breaks off and
					begins burrowing. If not, the stake may be used to make
					subsequent attacks until it strikes deep enough to ac-
					tivate.
					Once the tip of the stake is in the victim’s body, the
					Storyteller begins an extended roll of the caster’s Thaumaturgy rating (difficulty 9), rolling once per hour of
					game time. Successes on this roll are added to the successes scored in the initial attack. This represents the
					tip’s progress toward the victim’s heart. A botch indicates that the tip has struck a bone and all accumulated
					successes are lost (including those from the initial attack roll). Removing the part of the body where the tip
					impacted (such as a Tzimisce turning into blood or a
					vampire cutting off their arm) may stop the tip’s progress, depending on the number of successes acquired
					and the Storyteller’s discretion. When the shaft accu-
					mulates a total of 15 successes, it reaches the victim’s
					heart. This paralyzes Kindred and is instantly fatal to
					mortals and ghouls.
					Attempts to surgically remove the tip of the shaft
					can be made with an extended Dexterity + Medicine
					roll made once per hour (difficulty 7 for Kindred and 8
					for mortals). The surgeon must accumulate a number
					of successes equal to those currently held by the shaft
					in order to remove the tip. Once surgery begins, how-
					ever, the shaft begins actively evading the surgeon’s
					probes, and its rolls are made once every 30 minutes for
					the duration of the surgery attempt. Each individual
					surgery roll that scores less than three successes inflicts
					an additional unsoakable level of lethal damage on the
					patient.
					Shaft of Belated Quiescence may be performed on
					other wooden impaling weapons, such as spears, ar-
					rows, practice swords, and pool cues, provided that
					they are made of rowan wood. It may not, however,
					create a Bullet of Belated Quiescence.`,
				},
			},
			"Ward versus Lupines": {
				BaseDesc: `This wards an object in a manner identical to that of
				the Level Two ritual Ward Versus Ghouls, except that
				it affects werewolves`,
				Lvl: 3,
				System: map[string]string{
					"sys": `Ward versus Lupines behaves exactly as does
					Ward versus Ghouls, but it affects werewolves rather
					than ghouls. The ritual requires a handful of silver dust
					rather than a blood point.`,
				},
			},
			"Bone of Lies": {
				BaseDesc: `This ritual enchants a mortal bone so that anyone
				who holds it must tell the truth. The bone in question
				is often a skull, though any part of the skeleton will
				do — some casters use strings of teeth, necklaces of
				finger joints, or wands fashioned from ribs or arms. The
				bone grows blacker as it compels its holder to tell the
				truth, until it has turned completely ebony and has no
				magic left.
				This ritual binds the spirit of the individual to whom
				the bone belonged in life; it is this spirit who wrests
				the truth from the potential liar. The spirit absorbs
				the lies intended to be told by the bone’s holder, and
				as it compels more truth, it becomes more and more
				corrupt. If summoned forth, this spirit reflects the sins it has siphoned from the defeated liar (in addition to
				anger over its unwilling servitude). For this reason,
				anonymous bones are often used in the ritual, and the
				bone is commonly buried after it has been used to its
				full extent. A specific bone may never be used twice
				for this ritual, though different bones from the same
				corpse can.
				`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The bone imbued with this magical power
					must be at least 200 years old and must absorb 10 blood
					points on the night that the ritual is cast. Each lie
					the holder wishes to tell consumes one of these blood
					points, and the holder must speak the truth immediately thereafter. When all 10 blood points have been
					consumed, the bone’s magic ceases to work.`,
				},
			},
			"Firewalker": {
				BaseDesc: `This ritual imbues the vampire with an unnatural resistance to fire. Only a foolish vampire would actually
				attempt to walk on or through fire, but this ritual does
				grant an advanced tolerance to flame. Some Sabbat use
				this ritual to show off, while other thaumaturges use
				it only for martial concerns. To enact the ritual, the
				caster must cut off the end of one of his fingers and
				burn it in a Thaumaturgical circle.
				`,
				Lvl: 4,
				System: map[string]string{
					"sys": `Cutting off one’s finger does not do any
					health levels of damage, but it hurts like hell and requires a Willpower roll to perform. This ritual may be
					cast on other vampires (at the expense of the caster’s
					fingertips...). If the subject has no Fortitude, he may
					soak fire with his Stamina for the duration of this ritu-
					al. If the vampire has Fortitude, he may soak fire with
					his Stamina + Fortitude for the duration of the ritual.
					This ritual lasts one hour.`,
				},
			},
			"Heart of Stone": {
				BaseDesc: `A vampire under the effect of this ritual experiences
				the transformation suggested by the ritual’s name: his
				heart is completely transmuted to solid rock, render-
				ing him virtually impervious to staking. The subsidiary
				effects of the transformation, however, seem to follow
				the Hermetic laws of sympathetic magic: The vam-
				pire’s emotional capacity becomes almost nonexistent,
				and his ability to relate to others suffers as well.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `This ritual requires nine hours (reduced by
						one hour for every success). It can only be cast on oneself. The caster lies naked on a flat stone surface and
						places a bare candle over his heart. The candle burns
						down to nothing over the course of the ritual, causing
						one aggravated health level of damage (difficulty 5 to
						soak with Fortitude).
						At the end of the ritual, the caster’s heart hardens
						to stone. The caster gains a number of additional dice
						equal to twice his Thaumaturgy rating to soak any attack that aims for his heart and is completely impervious to the effects of a Shaft of Belated Quiescence (see
						p. 237). Additionally, the difficulty to use all Presence
						or other emotionally manipulative powers on him is
						increased by three due to his emotional isolation. The
						drawbacks are as follows: the caster’s Conscience/Conviction and Empathy scores drop to 1 (or to 0 if they
						already were at 1) and all dice pools for Social rolls except those involving Intimidation are halved (including those required to use Disciplines). All Merits that
						the character has pertaining to positive social interaction are neutralized. Heart of Stone lasts as long as the
						caster wishes it to.`,
				},
			},
			"Splinter Servant": {
				BaseDesc: `Another ritual designed to enchant a stake, Splinter
				Servant is a progressive development of Shaft of Belated Quiescence. (The two rituals, however, are mutually exclusive.) A Splinter Servant consists of a stake
				carved from a tree which has nourished itself on the
				dead. The stake must be bound in wax-sealed night-
				shade twine. When the binding is torn off, the Splin-
				ter Servant leaps to life, animating itself and attacking
				whomever the wielder commands — or the wielder, if
				she is too slow in assigning a target. The servant splits
				itself into a roughly humanoid form and begins singlemindedly trying to impale the target’s heart. Its exertions tear it apart within a few minutes, but if it pierces
				its victim’s heart before it destroys itself, it is remarkably difficult to remove, as pieces tend to remain behind if the main portion is yanked out.`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The ritual requires 12 hours to cast, minus one per success, and the servant must be created
					as described above. When the binding is torn off, the
					character who holds it must point the servant at its
					target and verbally command it to attack during the
					same turn. If this command is not given, the servant
					attacks the closest living or unliving being, usually the
					unfortunate individual who currently carries it.
					A Splinter Servant always aims for the heart. It has
					an attack dice pool of the caster’s Wits + Occult, a
					damage dice pool of the caster’s Thaumaturgy rating,
					and a maximum movement rate of 30 yards or meters
					per turn. Note that these values are those of the caster
					who created the servant, not the individual who activates it. A Splinter Servant cannot fly, but can leap
					its full movement rating every turn. Every action it
					takes is to attack or move toward its target; it cannot
					dodge or split its dice pool to perform multiple attacks.
					The servant makes normal stake attacks that aim for
					the heart (difficulty 9), and its success is judged as per
					the rules for a normal staking. A Splinter
					Servant has three health levels, and attacks directed
					against it are made at +3 difficulty due to its small size
					and erratic movement patterns.
					A Splinter Servant has an effective life of five com-
					bat turns per success rolled in its creation. If it has not
					impaled its victim by the last turn of its life, the servant
					collapses into a pile of ordinary, inanimate splinters.
					Three successes on a Dexterity roll (difficulty 8) are
					required to remove a Splinter Servant from a victim’s
					heart without leaving behind shards of the stake.`,
				},
			},
			"Ward versus Kindred": {
				BaseDesc: `This warding ritual functions exactly as do Ward
				versus Ghouls and Ward versus Lupines, but it inflicts
				injury upon Cainites`,
				Lvl: 4,
				System: map[string]string{
					"sys": `Ward versus Kindred behaves exactly as
					does Ward versus Ghouls, but it affects vampires rather
					than ghouls. The ritual requires a blood point of the
					caster’s own blood and does not affect the caster.`,
				},
			},
			"Blood Contract": {
				BaseDesc: `This ritual creates an unbreakable agreement between the two parties who sign it. The contract must
				be written in the caster’s blood and signed in the blood
				of whoever applies their name to the document. This
				ritual takes three nights to enact fully, after which both
				parties are compelled to fulfill the terms of the contract.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `This ritual is best handled by the Storyteller, who may bring those who sign the blood contract
					into compliance by whatever means necessary (it is
					not unknown for demons to materialize and enforce
					adherence to certain blood contracts). The only way
					to terminate the ritual is to complete the terms of the
					contract or to burn the document itself — attempts to
					add a clause forbidding burning the contract have resulted in the contract spontaneously combusting upon
					completion of the ritual. One blood point is consumed
					in the creation of the document, and an additional
					blood point is consumed by those who sign it.`,
				},
			},
			"Ward versus Spirits": {
				BaseDesc: `This warding ritual functions exactly as Ward versus
				Ghouls, but it inflicts injury upon spirits. Several other
				versions of this ward exist, each geared toward a particular type of non-physical being.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `Ward versus Spirits behaves exactly as does
					Ward versus Ghouls, but it affects spirits (including those
					summoned or given physical form by Thaumaturgy Paths
					such as Elemental Mastery). The material component
					for Ward versus Spirits is a handful of pure sea salt.
					The other versions of this ward, also Level Five rituals, are Ward versus Ghosts and Ward versus Demons.
					Each of these three Level Five wards affects its respective target on both the physical and spiritual planes.
					Ward versus Ghosts requires a handful of powdered
					marble from a tombstone, while Ward versus Demons
					requires a vial of holy water.`,
				},
			},
			"Paper Flesh": {
				BaseDesc: `This dreadful ritual enfeebles the subject, making
				her skin brittle and weak. Humors rise to the surface
				and flesh tightens around bones and scales away at the
				slightest touch. Used against physically tough opponents, this ritual strips away the inherent resilience of
				the vampiric body, leaving it a fragile, dry husk. The
				thaumaturge must inscribe his subject’s true name
				(which is much harder to discern for elders than it is
				for young vampires) on a piece of paper, which he uses
				to cut himself and then burns to cinders.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `This ritual causes the subject’s Stamina and
					Fortitude (if any) to drop to 1 each. For every Generation below Eighth, the subject retains one extra point
					of Stamina or Fortitude (keeping Fortitude first, though
					she may not exceed her original scores). For example, a
					vampire of the Fourth Generation with Fortitude targeted by Paper Flesh would drop to a Stamina + Fortitude score of 6 (assuming the score was 6 or more to
					begin with). This ritual lasts one night.`,
				},
			},
			"Escape to a True Friend": {
				BaseDesc: `Escape to a True Friend allows the caster to travel to
				the person whose friendship and trust she most values.
				The ritual has a physical component of a yard-wide/
				meter-wide circle charred into the bare ground or
				floor. The caster may step into the circle at any time
				and speak the true name of her friend. She is instantly
				transported to that individual, wherever he may be at
				the moment. She does not appear directly in front of
				him, but materializes in a location within a few minutes’ walk that is out of sight of any observer. The circle
				may be reused indefinitely, as long as it is unmarred.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `This ritual takes six hours a night for six
					nights to cast, reduced by one night for every two successes. Each night requires the sacrifice of three of the
					caster’s own blood points, which are poured into the
					circle. Once the circle is complete, the transport may
					be attempted at any time. The caster may take one
					other individual with her when she travels, or a maximum amount of “cargo” equal to her own weight`,
				},
			},
			"Enchant Talisman": {
				BaseDesc: `This ritual is the first taught to most Tremere once
				they have attained mastery of their first path. Create
				Talisman allows the caster to enchant a personal magical item (the fabled wizard’s staff) to act as an amplifier
				for her will and thaumaturgical might. Many talismans
				are laden with additional rituals (such as every ward
				known to the caster). The physical appearance of a
				talisman varies, but it must be a rigid object close to
				a yard or a meter long. Swords and walking sticks are
				the most common talismans, but some innovative or
				eccentric thaumaturges have enchanted violins, shotguns, pool cues, and classroom pointers.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `This ritual requires six hours per night for
					one complete cycle of the moon, beginning and ending
					on the new moon. Over this time, the vampire care-
					fully prepares her talisman, carving it with Hermetic
					runes that signify her true name and the sum total of
					her thaumaturgical knowledge. The player spends one
					blood point per night and makes an extended roll of
					Intelligence + Occult (difficulty 8), one roll per week.
					If a night’s work is missed or if the four rolls do not
					accumulate at least 20 net successes, the talisman is
					ruined and the process must be begun again.
					A completed talisman gives the caster several advan-
					tages. When the character is holding the talisman, the
					difficulty of all magic that targets her is increased by
					one. The player receives two extra dice when rolling
					for uses of the character’s primary path and one extra
					die when rolling for the character’s ritual castings. If
					the talisman is used as a weapon, it gives the player
					an additional die to roll to hit. If the thaumaturge is
					separated from her talisman, a successful Perception +
					Occult roll (difficulty 7) gives her its location.
					If a talisman is in the possession of another individu-
					al, it gives that individual three additional dice to roll
					when using any form of magic against the talisman’s
					owner. At the Storyteller’s discretion, rituals that tar-
					get the caster and use her talisman as a physical com-
					ponent may have greatly increased effects.
					A thaumaturge may only have one talisman at a
					time. Ownership of a talisman may not be transferred
					— each individual must create her own.`,
				},
			},
		},
	},
	"vicissitude": {
		Name: "Vicissitude",
		Description: `Vicissitude is the signature power of the Tzimisce,
		and is rarely shared outside the Clan (though it is
		known to some other Cainites of the Sabbat). Similar
		to Protean in some ways, Vicissitude allows vampires
		to shape and sculpt flesh and bone. When a Kindred
		uses Vicissitude to alter mortals, ghouls, and vampires
		of higher Generation, the effects of the power are permanent; vampires of equal or lower Generation can
		choose to heal the effects of Vicissitude as though they
		were aggravated wounds. A wielder of Vicissitude can
		always reshape her own flesh.
		The wielder must establish skin-to-skin contact and
		must often manually sculpt the desired result for these
		powers to take effect. This also applies to the use of the
		power on oneself. Tzimisce skilled in Vicissitude are
		often inhumanly beautiful; those less skilled are simply
		inhuman.
		There are rumors that Vicissitude is a disease rather
		than a “normal” Discipline, but only the Fiends know
		for sure, and they aren’t talking.
		Note: Nosferatu always “heal” Vicissitude alterations, at least the ones that make them better-looking.
		The ancient curse of the Clan may not be circumvented through Vicissitude. The same applies to physical
		deformities from the Gangrel Clan weakness.`,
		Abilities: map[string]disciplineAbility{
			"Malleable Visage": {
				BaseDesc: `A vampire with this power may alter her own bodily
				parameters: height, build, voice, facial features, and
				skin tone, among other things. Such changes are cosmetic and minor in scope — no more than a foot (30
				cm) of height gained or lost, for example. She must
				physically mold the alteration, literally shaping her
				flesh into the desired result.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `The player must spend a blood point for
					each body part to be changed, then roll Intelligence +
					Medicine (difficulty 6). To duplicate another person or
					voice requires a Perception + Medicine roll (difficulty
					8), and five successes are required for a flawless copy;
					fewer successes leave minute (or not-so-minute) flaws.
					Increasing one’s Appearance Trait has a difficulty of 9,
					and the vampire must spend an additional blood point
					for each dot of Appearance increased beyond their natural total. A botch permanently reduces the Attribute
					by one.`,
				},
			},
			"Fleshcraft": {
				BaseDesc: `This power is similar to Malleable Visage, above, but
				allows the vampire to perform drastic, grotesque alterations on other creatures. Tzimisce often use this power
				to transform their servitors into monstrous guards, the
				better to frighten foes. Only flesh (skin, muscle, fat,
				and cartilage, but not bone) may be transformed.`,
				Lvl: 2,
				System: map[string]string{
					"sys": `After spending a blood point, the vampire must grapple the intended victim. The player of
					the Flescrafting vampire makes a successful Dexterity + Medicine roll (difficulty variable: 5 for a crude
					yank-and-tuck, up to 9 for precise transformations). A
					vampire who wishes to increase another’s Appearance
					Trait does so as described under Malleable Visage; reducing the Attribute is considerably easier (difficulty
					5), though truly inspired disfigurement may dictate a
					higher difficulty. In either case, each success increases
					or reduces the Attribute by one.
					A vampire may use this power to move clumps of
					skin, fat, and muscle tissue, thus providing additional
					padding where needed. For each success scored on a
					Dexterity + Medicine roll (difficulty 8), the vampire
					may increase the subject’s soak dice pool by one, at the
					expense of either a point of Strength or a health level
					(vampire’s choice).`,
				},
			},
			"Bonecraft": {
				BaseDesc: `This terrible power allows a vampire to manipulate
				bone in the same manner that flesh is shaped. In conjunction with Fleshcraft, above, this power enables a
				Vicissitude practitioner to deform a victim (or herself)
				beyond recognition. This power should be used in conjunction with the flesh-shaping arts, unless the vampire wants to inflict injury on the victim.`,
				Lvl: 3,
				System: map[string]string{
					"sys": `The vampire’s player must spend a blood
					point and make a Strength + Medicine roll (difficulties
					as above). Bonecraft may be used without the fleshshaping arts, as an offensive weapon. Each success
					scored on the Strength + Medicine roll (difficulty 7)
					inflicts one health level of lethal damage on the victim, as his bones rip, puncture, and slice their way out
					of his skin.
					The vampire may utilize this power (on herself or
					others) to form spikes or talons of bone, either on the
					knuckles as an offensive weapon or all over the body
					as defensive “quills.” If bone spikes are used, the vampire or victim takes one health level of lethal damage
					(the vampire’s comes from having the very sharp bone
					pierce through his skin — this weaponry doesn’t come
					cheaply). In the case of quills, the subject takes a number of health levels equal to five minus the number of
					successes (a botch kills the subject or sends the vampire
					into torpor). These health levels may be healed normally. Knuckle spikes inflict Strength +1 lethal damage. Defensive quills inflict a hand-to-hand attacker’s
					Strength in lethal damage unless the attacker scores
					three or more successes on the attack roll (in which
					case the attacker takes no damage); the defender still
					takes damage normally. Quills also enable the vampire
					or altered subject to add two to all damage inflicted via
					holds, clinches, or tackles.
					A vampire who scores five or more successes on the
					Strength + Medicine roll may cause a rival vampire’s
					rib cage to curve inward and pierce the heart. While
					this does not send a vampire into torpor, it does cause
					the affected vampire to lose half his blood points, as
					the seat of his vitae ruptures in a shower of gore.`,
				},
			},
			"Horrid Form": {
				BaseDesc: `Kindred use this power to become hideous and deadly monsters. The vampire’s stature increases to a full
				eight feet (two and a half meters), the skin becomes a
				sickly greenish-gray or grayish-black chitin, the arms
				become apelike and ropy with ragged black nails, and
				the face warps into something out of a nightmare. A
				row of spines sprouts from the vertebrae, and the external carapace exudes a foul-smelling grease`,
				Lvl: 4,
				System: map[string]string{
					"sys": `The Horrid Form costs two blood points to
					awaken. All Physical Attributes increase by three, but
					all Social Attributes drop to zero, except when dealing
					with others also in Horrid Form. However, a vampire
					in Horrid Form who is trying to intimidate someone
					may substitute Strength for a Social Attribute. Damage inflicted in brawling combat increases by one due
					to the jagged ridges and bony knobs creasing the creature’s hands.`,
				},
			},
			"Bloodform": {
				BaseDesc: `A vampire with this power can physically transform
				all or part of her body into sentient vitae. This blood
				is in all respects identical to the vampire’s normal vitae; she can use it to nourish herself or others, create
				ghouls, or establish blood bonds. If all this blood is imbibed or otherwise destroyed, the vampire meets Final
				Death.`,
				Lvl: 5,
				System: map[string]string{
					"sys": `The vampire may transform all or part of
					herself as she deems fit. Each leg can turn into two
					blood points worth of vitae, as can the torso; each arm,
					the head, and the abdomen convert to one blood point
					each. The blood can be reconverted to the body part,
					provided it is in contact with the vampire. If the blood
					has been utilized or destroyed, the vampire must spend
					a number of blood points equal to what was originally
					created to regrow the missing body part.
					A vampire entirely in this form may not be staked,
					cut, bludgeoned, or pierced, but can be burned or exposed to the sun. The vampire may ooze along, drip up
					walls, and flow through the narrowest cracks, as though
					she were in Tenebrous Form.
					Mental Disciplines may be used, provided no eye
					contact or vocal utterance is necessary, although the
					vampire can perceive her surroundings just fine (but
					the perceptions are always centered on the largest pool
					of blood). If a vampire in this form “washes” over a
					mortal or animal, that mortal must make a Courage
					roll (difficulty 8) or fly into a panic.`,
				},
			},
			"Chiropteran Marauder": {
				BaseDesc: `Similar to Horrid Form, the Chiropteran Marauder
				is a terrifying, bipedal bat, bearing a wickedly fanged
				maw and veined, leathery wings. This power confers all
				of the benefits of the Horrid Form, in addition to a few
				others. The mere sight of the marauder is enough to
				make mortals or weak-willed vampires flee in horror.`,
				Lvl: 6,
				System: map[string]string{
					"sys": ` The vampire gains all the effects of the Horrid Form. Further, the Marauder’s fluted wings allow
					flight at 25 mph (40 kph), during which the vampire
					may carry, but not manipulate, objects no larger than
					it can hold in its hands. If the player chooses to, she
					may make a Strength + Medicine roll (difficulty 6) to
					extend bony claws at the ends of the vampire’s wings.
					These claws inflict Strength +2 aggravated damage.
					Also, the vampire subtracts two from all hearing-based
					Perception rolls (though he adds one to vision- based
					Wits and Perception rolls). Assuming the mantle of
					the Chiropteran Marauder costs three blood points.`,
				},
			},
			"Blood of Acid": {
				BaseDesc: `At this level of mastery, the vampire has converted
				his blood to a viscous acid. Any blood he consumes
				likewise becomes acid, which is corrosive enough to
				burn human (and vampiric flesh) as well as wood. This
				effect is particularly potent when the vampire assumes
				the Bloodform. One of the side effects of this power
				is the complete inability to create new vampires and
				ghouls, or give blood to another vampire — the acid
				would corrode them as they drank it. The obvious benefit, however, is that would-be diablerists are likewise
				unable to devour the Cainite’s blood.`,
				Lvl: 6,
				System: map[string]string{
					"sys": `Each acidified blood point that comes in
					contact with something other than the vampire himself does five dice worth of aggravated damage. If the
					vampire is injured in combat, his blood may spatter on
					an opponent — foes must make Dexterity + Athletics
					rolls to avoid the blood, which must be accomplished by
					splitting their dice pools. (Unless an opponent knows
					the vampire has this power, she’s unlikely to split her
					dice pool on her first attack, which causes many Tzimisce to cackle with glee.)`,
				},
			},
			"Cocoon": {
				BaseDesc: `The Cainite can form an opaque cocoon from blood
				and other fluids excreted from her body. The cocoon
				hardens after a few moments, turning into a tough,
				white shell. This cocoon provides considerable protection to the vampire, even shielding her from sunlight
				and, to a limited degree, fire.`,
				Lvl: 1,
				System: map[string]string{
					"sys": `A vampire may only cocoon herself, and the
					process takes 10 minutes. Additionally, creating a cocoon costs three blood points. The cocoon offers complete protection from sunlight, and provides a number
					of dice of soak equal to twice the vampire’s unmodified
					Stamina against all damage, aggravated or otherwise. It lasts as long as the Cainite wishes, and she may dissolve it at her whim, emerging from a pulpy, bloody
					paste. A vampire contained within a cocoon may still
					use mental Disciplines that do not require eye contact
					or other conditions to be met.`,
				},
			},
			"Breath of the Dragon": {
				BaseDesc: `The vampire becomes like one of the terrible draculs
				of the Old World, able to exhale a deadly gout of flame.
				This flame does not hurt the vampire himself, though
				he may become trapped in flames if it engulfs flammable objects.`,
				Lvl: 8,
				System: map[string]string{
					"sys": `The flaming cloud affects a six-foot area, doing two dice of aggravated damage to anyone in the
					flames’ circumference. This fire may cause combustible
					items to ignite, and it may ignite victims who suffer fire
					damage.`,
				},
			},
			"Command": {
				BaseDesc: `This power, developed in the nights when the Tzimisce were the terrible masters of Eastern Europe,
				allows the vampire to sink into and disperse himself
				through the ground. Unlike the Protean power of Earth
				Meld, however, the vampire actually dissolves his body
				into the ground; nothing short of a wide-area explosion can affect him, and he may not be dug up bodily.
				In addition, from sunset to sunrise, the vampire may
				see and hear everything happening in his environment
				through his mystical connection to the land. The mere
				fact that this power exists terrifies many Tzimisce, who
				are secretly unsure of whether or not the diablerie of
				their Antediluvian was successful.`,
				Lvl: 9,
				System: map[string]string{
					"sys": `This power costs six blood points to activate, and it lasts as long as the vampire wishes to remain contained within the soil. As per the Cocoon,
					the vampire may use mental Disciplines that do not
					require physical solvency or eye contact. He may communicate mentally with anyone who wanders into the
					area under which he rests.`,
				},
			},
		},
	},
}

var InfoMap = map[string]map[string]Characteristic{
	"attribute": {
		"strength": {
			BaseDesc: ` Strength is the raw, brute power of a character. It
			governs how much weight a character can lift, how much he can
			physically push, and how hard he can hit another character or
			object. This trait is added to a character’s damage dice pool
			when they hit an opponent in hand-to-hand combat. It is also
			used when a character wishes to break, lift, or carry
			something,as well as when a character tries to jump a distance.
			`,
			ValDesc: map[string]string{
				"Poor":        "You can lift 40 lbs",
				"Average":     "You can lift 100 lbs",
				"Good":        "You can lift 250 lbs",
				"Exceptional": "You can lift 400 lbs",
				"Outstanding": "You can lift 650 lbs and crush skulls like grapes",
			},
		},
		"dexterity": {
			BaseDesc: ` Dexterity  measures a character’s general physical
			prowess.  It encompasses the character’s speed, agility, and
			overall quickness, as well as indicating the character’s ability
			to manipulate objects with control and precision. Also included
			under Dexterity’s heading are hand-eye coordination, reflexes,
			and bodily grace.  `,
			ValDesc: map[string]string{
				"Poor":        "You are clumsy and awkward. Put that gun down before you hurt yourself",
				"Average":     "You’re no klutz, but you’re no ballerina, either",
				"Good":        "You possess some degree of athletic potential.",
				"Exceptional": "You could be an acrobat if you wished",
				"Outstanding": "Your movements are liquid and hypnotic — almost superhuman",
			},
		},
		"stamina": {
			BaseDesc: ` Stamina reflects a character’s health, toughness,
			and resilience. It indicates how long a character can exert
			themselves and how much punishment they can withstand before
			suffering physical trauma. Stamina also includes a bit of
			psychic fortitude, indicating a character’s grit and tenacity.
			`,
			ValDesc: map[string]string{
				"Poor":        "You have glass skin and paper bones.",
				"Average":     "You are moderately healthy and can take a punch or two.",
				"Good":        "You are in good shape and rarely fall ill.",
				"Exceptional": "You can run — and perhaps win — any marathon you choose.",
				"Outstanding": "Your constitution is truly herculean.",
			},
		},
		"charisma": {
			BaseDesc: ` Charisma is a character’s ability to entice and
			please others through their personality. Charisma comes into
			question when a character tries to win another character’s
			sympathies or encourage others to trust her. Charisma reflects
			the power of a character’s charm and influence. It governs a
			character’s ability to convince others to see her point of view.
			This Attribute doesn’t necessarily indicate how the character is
			charismatic, whether she’s a silver-tongued charmer or a
			grinning bully.  `,
			ValDesc: map[string]string{
				"Poor":        "There’s something a little sketchy about you",
				"Average":     "You are generally likable and have several friends",
				"Good":        "People trust you implicitly",
				"Exceptional": "You have significant personal magnetism",
				"Outstanding": "Entire cultures could follow your lead",
			},
		},
		"manipulation": {
			BaseDesc: ` Manipulation measures a character’s ability for self
			expression in the interests of getting others to share their
			outlook or follow her whims. Manipulation is used to trick,
			bluff, fast-talk, and railroad other characters. Whether or not
			the characters in question actually like the manipulator is
			irrelevant (this is why Manipulation differs from
			Charisma).Botching a Manipulation roll may add a name to the
			character’s list of enemies.People are manipulated every day,
			and typically ignore it. If the fact isbrought to their
			attention, however, many people become quite
			defensive.Characters with high Manipulation ratings are often
			distrusted by those around them.  `,
			ValDesc: map[string]string{
				"Poor":        "A person of few (often ineffectual) words",
				"Average":     "You can fool some of the people some of the time, just like anybody else",
				"Good":        "You never pay full price.",
				"Exceptional": "You could be a politician or cult leader",
				"Outstanding": "You can convince people die in your stead",
			},
		},
		"appearance": {
			BaseDesc: ` Appearance is a measure of how well a character
			makes a first impression. This may be conventional
			“attractiveness,” but it can also be the effect of distinctive
			features, an exotic mien, an air of confidence, distinctive
			posture, a flair for dressing well — anything remarkable upon
			initial observation can contribute to a character’s Appearance.
			Appearance is subconscious and instinctual, it shapes first
			impressions and the nature of memories thereafter. This Trait is
			useful for getting potential vessels to heed your beckoning
			across a crowded dance floor. A character may have no more dice
			in a Social dice pool than her Appearance rating. Thus, it is
			critically important to either look your best or get to know
			people before you start trying to convince them to firebomb the
			Justicar’s haven.  `,
			ValDesc: map[string]string{
				"Poor":        "Your clothes stink, you turn people off right away, or you’re just damned ugly",
				"Average":     "You don’t stand out in a crowd, for better or for worse",
				"Good":        "People sometimes approach you and buy you drinks",
				"Exceptional": "People go out of their way to make your acquaintance",
				"Outstanding": "People never forget you",
			},
		},
		"perception": {
			BaseDesc: ` Perception measures a character’s ability to observe
			their environment. This may involve a conscious effort, such as
			searching an area, but it is more often intuitive, as the
			character’s keen senses notice something out of the ordinary.
			Perception is a sensitivity to one’s surroundings, and is seldom
			present in the cynical or jaded (who have seen it all
			before).Perception is used to determine whether or not a
			character understands a given situation or detects an
			environmental stimulus. It can warn a character of
			ambushes,distinguish a clue from a pile of refuse, or uncover
			any other hidden or overlookable detail, whether physical or
			otherwise.  `,
			ValDesc: map[string]string{
				"Poor":        "Perhaps you are absurdly self-absorbed, perhaps merely an airhead.",
				"Average":     "he very subtle evades you, but you’re aware of the bigger picture.",
				"Good":        "You perceive moods, textures, and small changes in your environment.",
				"Exceptional": "Almost nothing escapes your notice.",
				"Outstanding": "You instantly observe things almost imperceptible to human senses.",
			},
		},
		"intelligence": {
			BaseDesc: ` Intelligence refers to a character’s grasp of facts
			and knowledge. It also governs a character’s ability to reason,
			solve problems, and evaluate situations.Intelligence also
			includes critical thinking and flexibility of thought.
			Intelligence does not include savvy, wisdom, or common sense.
			Even the smartest character may be too foolish to realize the
			thugs who want to “borrow” her car keys are up to no
			good.Characters with low Intelligence aren’t necessarily stupid
			(though they might be); they are just uneducated or simple
			thinkers. Likewise, characters with high Intelligence aren’t all
			Einsteins; they may be better at rote memorization or have
			particularly keen judgement.  `,
			ValDesc: map[string]string{
				"Poor":        "Not the brightest crayon in the box",
				"Average":     "Smart enough to realize you’re normal.",
				"Good":        "More enlightened than the masses.",
				"Exceptional": "You’re not just bright, you’re downright brilliant.",
				"Outstanding": "Certified genius.",
			},
		},
		"wits": {
			BaseDesc: `
				Wit measures the charactedisr’s ability to think on their feet
				and react quickly. It also reflects a character’s general
				cleverness. Characters with low Wits ratings are thick and
				mentally lethargic, or maybe gullible and unsophisticated.
				By contrast, characters with high Wits Traits almost always
				have a plan immediately and adapt to their surroundings with
				striking expedience. Characters with high Wits also manage
				to keep their cool in stressful situations.
				`,
			ValDesc: map[string]string{
				"Poor":        "Your brains a little smooth.",
				"Average":     "You know when to bet or fold in poker.",
				"Good":        "You are seldom surprised or left speechless.",
				"Exceptional": "You’re one of the people who make others think, “I should have said-” the next day",
				"Outstanding": " You think and respond almost more quickly than you can act",
			},
		},
	},
	"talents": {
		"alertness": {
			BaseDesc: ` Alertness describes the attention you pay to the
			outside world, whether otherwise occupied or not.  This Talent
			is typically paired with Perception, and is best used when
			sensing physical stimuli (as opposed to moods or clues).				`,
			ValDesc: map[string]string{
				"Novice":    "You’re no mindless drone.",
				"Practiced": "Habitual eavesdropper",
				"Competent": "You keep a sharp eye on your surroundings.",
				"Expert":    "Whether from paranoia or good sense, you are rarely caught off-guard.",
				"Master":    "Your senses are on par with those of a wild animal.",
			},
		},
		"athletics": {
			BaseDesc: ` Athletics represents your basic athletic ability, as
			well as any training you might have had in sports or other
			rigorous activities. Athletics concerns all forms of running,
			jumping, throwing, swimming, sports, and the like. However,
			Athletics doesn’t cover basic motor actions such as lifting
			weights, nor does it govern athletic feats covered by another
			Ability (ex. Melee).`,
			ValDesc: map[string]string{
				"Novice":    "You had an active childhood.",
				"Practiced": "High-school athletes",
				"Competent": "Talented lifelong amateur",
				"Expert":    "Professional athletes",
				"Master":    "Olympic medalist",
			},
		},
		"awareness": {
			BaseDesc: ` Awareness is an instinctual reaction to the presence
			of the supernatural. It differs from Alertness (which measures
			sensitivity to mundane events) and Occult (which covers actual
			knowledge about the supernatural). Characters with Awareness
			sometimes get hunches, chills, or sudden flashes of inspiration
			when they are near supernatural creatures, objects, or events.
			This insight is purely subconscious, and knowing that
			something’s wrong doesn’t mean that the character knows what it
			is.`,
			ValDesc: map[string]string{
				"Novice":    "Once in a while, you get the feeling that something isn’t right.",
				"Practiced": "You sometimes get strange vibes from a particular direction or vague area (like a building).",
				"Competent": "You can walk into a room and know that something unusual is going on within.",
				"Expert":    "If you concentrate, you can sense whether a someone in a group of people or a collection of objects is supernatural.",
				"Master":    "You instinctually know if something or someone is mundane or supernatural.",
			},
		},
		"brawl": {
			BaseDesc: ` Brawl represents how well you fight in
			tooth-and-nail situations. This Talent represents skill in
			unarmed combat, whether from formal martial arts training or
			simply from plenty of experience. Effective brawlers are
			coordinated, resistant to pain, quick, strong, and mean. The
			willingness to do whatever it takes to hurt your opponent wins
			plenty of fights.`,
			ValDesc: map[string]string{
				"Novice":    "You were picked on as a kid.",
				"Practiced": "You’ve participated in the occasional barroom tussle.",
				"Competent": "You’ve fought regularly generally walked away in better shape than your opponents.",
				"Expert":    "You could be a serious contender on the MMA circuit.",
				"Master":    "There’s a video of you taking down three men in four seconds.",
			},
		},
		"empathy": {
			BaseDesc: ` With Empathy you understand the emotions of others,
			and can sympathize with, feign sympathy for, or play on such
			emotions as you see fit. You are adept at discerning motive, and
			might be able to discern when someone’s lying to you. However,
			you may be so in tune with other people’s feelings that your own
			emotions are affected.`,
			ValDesc: map[string]string{
				"Novice":    "You lend the occasional shoulder to cry on.",
				"Practiced": "You can sometimes literally feel someone else’s suffering.",
				"Competent": "You have a keen insight into other people’s motivations.",
				"Expert":    "It’s almost impossible to lie to you.",
				"Master":    "The human soul conceals no mysteries from you.",
			},
		},
		"expression": {
			BaseDesc: ` Expression is your ability to get your point across
			clearly, whether through conversation, poetry, or even in "1""4"0
			characters or fewer. Characters with high Expression can phrase
			their opinions or beliefs in a manner that cannot be ignored
			(even if their opinions are misinformed or worthless). They
			might also be talented actors, skilled at conveying moods or
			communicating emotion with every gesture. Additionally, this
			Talent represents your ability for poetry, creative writing, or
			other literary art forms. You can choose a specialty in
			Expression, even if the skill is below expert.  `,
			ValDesc: map[string]string{
				"Novice":    "Your talent has matured past crude poetry on notebook paper. ",
				"Practiced": "You could lead a collegedebate team.",
				"Competent": "You could be a successful writer.",
				"Expert":    "Your work is Pulitzer material.",
				"Master":    "You could be the next Steve Jobs.",
			},
		},
		"intimidation": {
			BaseDesc: `Intimidation takes many forms, from outright threats
			and physical violence to mere force of personality. It needn’t
			be course or callous, and a well-placed intimidating word under
			the right circumstances might well be called “diplomacy” in
			certain circles. You know the right method for each occasion,
			and can be very… persuasive.  `,
			ValDesc: map[string]string{
				"Novice":    "Shady teenager",
				"Practiced": "Skinhead thug",
				"Competent": "Drill sergeant",
				"Expert":    "Your air of authority cows casual passersby.",
				"Master":    "You can frighten off vicious animals.",
			},
		},
		"leadership": {
			BaseDesc: ` You are an example to others and can inspire them to
			do what you want. Leadership has less to do with manipulating
			people’s desires than it does with presenting yourself as the
			sort of person they want to follow. Anyone can lead a group into
			some sort of conflict; a good leader can get them back out
			intact. This Talent is usually paired with Charisma rather than
			Manipulation.  `,
			ValDesc: map[string]string{
				"Novice":    "Captain of your Little League team",
				"Practiced": "Student body president",
				"Competent": "An effective CEO",
				"Expert":    "Presidential material",
				"Master":    "You could be beloved dictator of a nation.",
			},
		},
		"streetwise": {
			BaseDesc: `The streets can provide a lot of information or money
			to those who know the language. Streetwise allows you to blend
			in unobtrusively with the local scene, pick up gossip,
			understand slang, or even dabble in criminal doings.  `,
			ValDesc: map[string]string{
				"Novice":    " You know who’s holding.",
				"Practiced": "You’re accorded respect on the street.",
				"Competent": "You could head your own gang.",
				"Expert":    "You have little to fear in even the worst neighborhoods.",
				"Master":    "If you haven’t heard it, it hasn’t been said. ",
			},
		},
		"subterfuge": {
			BaseDesc: ` You know how to conceal your own motives and project
			what you wish. Furthermore, if you can root out other people’s
			motives, you can then use those motives against them. This
			Talent defines your talent for intrigue, secrets, and
			double-dealing. Mastery of Subterfuge can make you the ultimate
			seducer or a brilliant spy.  `,
			ValDesc: map[string]string{
				"Novice":    "You tell the occasional white lie.",
				"Practiced": "You can spin elaborate lies",
				"Competent": "Criminal lawyer",
				"Expert":    "Deep-cover agent",
				"Master":    "You’re the very last person anyone would suspect.",
			},
		},
		"hobby": {
			BaseDesc: ` This category encompasses anything that the
			Storyteller deems to be mainly self-taught and is usually 
			(though not always) more active than intellectual.  `,
			ValDesc: map[string]string{
				"Novice":    "You’ve dabbled.",
				"Practiced": "You’ve got a good grasp of your hobby’s basics",
				"Competent": "Other practitioners regard you as fairly skilled and competent.",
				"Expert":    "You are familiar with the subtle nuances and applications of your Talent.",
				"Master":    "You could write a book on what you do. Perhaps you already have.",
			},
		},
		"animalken": {
			BaseDesc: `You can understand animals’ behavior patterns. This 
			Skill allows you to predict how an animal might react in a given
			situation, train a domesticated creature, or even try to calm or
			enrage animals `,
			ValDesc: map[string]string{
				"Novice":    "You can get a domesticated animals to let you pet them.",
				"Practiced": "You can housebreak a puppy",
				"Competent": "You could train a seeing-eye dog.",
				"Expert":    "Circus trainer",
				"Master":    " You can tame wild beasts without benefit of supernatural powers.",
			},
		},
		"crafts": {
			BaseDesc: `Crafts covers your ability to make or fix things with
			your hands. Crafts allows you to work in fields such as
			carpentry, leatherworking, weaving, or even mechanical expertise
			such as car repair. You can even create lasting works of art
			with this Skill, depending on the number of successes you
			achieve. You must always choose a specialization in Crafts, even
			though you retain some skill in multiple fields.  `,
			ValDesc: map[string]string{
				"Novice":    "High school wood shop",
				"Practiced": "You’re starting to develop your own style.",
				"Competent": "You could start your own shop.",
				"Expert":    "You wrote instruction manuals on your field of specialization",
				"Master":    "Your craftsmanship and insight is virtually without peer",
			},
		},
		"drive": {
			BaseDesc: `You can drive a car, and maybe other vehicles as
			well.  This Skill does not automatically entail familiarity with
			complicated vehicles such as tanks or "1"8-wheelers, and 
			difficulties may vary depending on your experience with
			individual automobiles.  `,
			ValDesc: map[string]string{
				"Novice":    "You know how to work an automatic transmission",
				"Practiced": "You can drive a stick shift.",
				"Competent": "Professional trucker.",
				"Expert":    " NASCAR daredevil or tank pilot.",
				"Master":    ": Whether it’s a Fiat or a Ferrari, you can make it sing.",
			},
		},
		"etiquette": {
			BaseDesc: ` You understand the nuances of proper behavior, in 
			both mortal society and Kindred culture. In many cases, knowing
			how to broach a topic is as important as the discussion itself,
			and a person with poor etiquette will never have an opportunity
			to make herself heard because she doesn’t know when or how to
			interject.  This Skill is used during meetings, haggling,
			seduction, dancing, dinner etiquette, and all forms of
			diplomacy.  `,
			ValDesc: map[string]string{
				"Novice":    "You know when to keep your mouth shut",
				"Practiced": "You’ve been to a black-tie event or two.",
				"Competent": "You know your way around even obscure silverware ",
				"Expert":    "Her Majesty would consider you charming. ",
				"Master":    "If the right people came to dinner, you could end wars — or start them.",
			},
		},
		"firearms": {
			BaseDesc: ` This Skill represents familiarity with a range of
			firearms, from holdout pistols to heavy machine guns. Further, 
			someone skilled in Firearms can clean, repair, recognize, and
			accurately fire most forms of small arms. This Skill is also
			used to unjam guns (Wits + Firearms).`,
			ValDesc: map[string]string{
				"Novice":    "You had a BB gun as a kid",
				"Practiced": "You while away the occasional hour at the gun club.",
				"Competent": "You’ve survived a firefight or two",
				"Expert":    "You could pick off people for a living",
				"Master":    " You’ve been practicing since the debut of the Winchester",
			},
		},
		"larceny": {
			BaseDesc: `This Skill entails familiarity with the tools and techniques for the sorts of physical manipulation typically 
			associated with criminal activity. Picking locks, manual 
			forgery, safecracking, simple hotwiring, various forms 
			of breaking and entering, and even sleight-of-hand all 
			fall under the auspices of Larceny. Larceny is useful not 
			only for theft, but also for setting up “the unbeatable 
			system” or deducing where a thief broke in. This skill 
			does not confer any aptitude with advanced security or 
			anti-crime technologies such as video surveillance or 
			alarm systems — those are covered by the Technology 
			Knowledge.`,
			ValDesc: map[string]string{
				"Novice":    "You can pick a simple lock.",
				"Practiced": "You could run a shell game hustle on the corner.",
				"Competent": "You can open a standard locked window from the outside.",
				"Expert":    "You can “retool” a passport or ID card.",
				"Master":    "You could get into (or out of…) a multinational bank’s central vault.",
			},
		},
		"melee": {
			BaseDesc: `As the Kindred maxim runs, “Guns mean nothing to a
			lifeless heart”. A blade is often worth far more, as is the
			skill to use it properly. Melee covers your ability to use
			hand-to-hand weapons of all forms, from swords and clubs to
			esoteric martial-arts paraphernalia such as sai or nunchaku.
			And, of course, there is always the utility of the wooden stake.
			`,
			ValDesc: map[string]string{
				"Novice":    "You know the right way to hold a knife.",
				"Practiced": "You may have been in the occasional street fight.",
				"Competent": "You could make a college fencing team.",
				"Expert":    "You could keep order in the Prince’s court.",
				"Master":    "Your enemies would rather face a SWAT team than your blade.",
			},
		},
		"performance": {
			BaseDesc: ` The Performance Skill governs your ability to
			perform artistic endeavors such as singing, dancing, acting, or
			playing a musical instrument. You are almost certainly
			specialized in one field, although true virtuosos may be
			talented in many forms of performance. This Skill represents not
			only technical know-how, but the ability to work an audience and
			enrapture them with your show. As with Crafts, you must choose a
			specialty, even though this Skill also imparts a general sense
			for watching and responding to your audience’s mood regardless
			of medium.`,
			ValDesc: map[string]string{
				"Novice":    "You could sing in the church choir.",
				"Practiced": "Your Internet videos have over a hundred thousand views.",
				"Competent": "You almost always have a gig booked.",
				"Expert":    "You have the talent to be a national sensation.",
				"Master":    "You are a virtuoso without peer.",
			},
		},
		"stealth": {
			BaseDesc: `This Skill is the ability to avoid being detected, 
			whether you’re hiding or moving at the time. Stealth is often
			tested against someone else’s Perception + Alertness. This
			Ability is, for obvious reasons, highly useful in stalking prey.
			In many cases, Stealth is also used to conceal items, whether on
			one’s person or somewhere in the environment.  `,
			ValDesc: map[string]string{
				"Novice":    "You can hide in a darkened room.",
				"Practiced": "You can shadow someone from streetlight to streetlight.",
				"Competent": "You have little difficulty finding prey from evening to evening.",
				"Expert":    "You can move quietly over dry leaves.",
				"Master":    "Nosferatu elder",
			},
		},
		"survival": {
			BaseDesc: `Although vampires have little to fear from starvation
			and exposure, the wilderness can still be dangerous to a
			Cainite. This Skill allows you to find shelter, navigate your
			way to civilization, track prey, establish a makeshift haven,
			and possibly even avoid supernatural threats like werewolves
			that also prowl the World of Darkness. Note that Survival need
			not be used only in areas considered “wilderness.” There’s
			plenty of Survival that goes into getting by in various parts of
			modern cities.
			`,
			ValDesc: map[string]string{
				"Novice":    "You can survive a night spent outside.",
				"Practiced": "You’ve “roughed it” on a regular basis.",
				"Competent": "You can separate poison or spoilage from edible forage.",
				"Expert":    "You could live for months in the challenging environment of your choice.",
				"Master":    "You could get dropped naked into the Andes and do all right for yourself.",
			},
		},
		"professional": {
			BaseDesc: `This category encompasses anything that the
			Storyteller deems to be a taught Ability and is primarily active
			in application. Storytellers should first examine the list of
			existing Skills to determine if a particular task might fall
			under one of those (e.g. Tracking would be a specialty of
			Survival).`,
			ValDesc: map[string]string{
				"Novice":    "You’ve apprenticed.",
				"Practiced": "You have a handle on the basics.",
				"Competent": "You could make a living, although not a fortune, doing what you do.",
				"Expert":    "You know the more esoteric uses of your Skill, and are rarely at a loss.",
				"Master":    "You are an acknowledged authority in your chosen field of endeavor.",
			},
		},
		"academics": {
			BaseDesc: `This catchall Knowledge covers the character’s
			erudition in thehumanities: literature, history, art,
			philosophy, and other “liberal” arts and sciences. A character
			with dots in Academics is generally well rounded in these
			fields, and at high levels may be considered an expert in one or
			more areas of study. Not only can this Knowledge impress at
			salons and other Elysium functions, but it can also offer
			valuable clues to certain past — and future — movements in the
			Jyhad.`,
			ValDesc: map[string]string{
				"student":   "You’re aware that 1066 isn’t a Beverly Hills area code.",
				"college":   "You can quote from the classics and identify major cultural movements.",
				"masters":   "You could get a paper published in a scholarly journal.",
				"doctorate": "Professor emeritus",
				"scholar":   "Scholars worldwide acknowledge you as one of the foremost experts of your time.",
			},
		},
		"computer": {
			BaseDesc: `This Knowledge represents the ability to operate and 
		program computers, including mobile devices. Most Computer use also
		imparts a degree of Internet awareness (if not savvy).`,
			ValDesc: map[string]string{
				"student":   "You can navigate touch-screen and traditional point-and-click ",
				"college":   "You know your way around various applications and the Internet.",
				"masters":   "You know what to do with a text command prompt.",
				"doctorate": " You can make a very comfortable living as a consultant.",
				"scholar":   "You have all the SDKs and comprehend data structures for a stunning variety of programming languages",
			},
		},
		"finance": {
			BaseDesc: `You know the ins and outs of commerce, from evaluating an
		item’s relative worth to keeping up with currency exchange rates.
		This Knowledge can be invaluable when brokering items, running
		numbers, or playing the stock market. Sufficiently high levels in 
		Finance allow you to raise your standards of living to a very
		comfortable level.`,
			ValDesc: map[string]string{
				"student":   "You’ve taken a few business classes.",
				"college":   "You have some practical experience and can keep your books fairly neat.",
				"masters":   "You’d make a fine stockbroker.",
				"doctorate": "Corporations follow your financial lead.",
				"scholar":   "You could turn a $20 bill into a fortune.",
			},
		},
		"investigation": {
			BaseDesc: `You’ve learned to notice details others might overlook,
		and might make an admirable detective. This Knowledge represents not
		only a good eye for detail, but also an ability to do research and
		follow leads. Such research may include Internet searches or more
		specific research techniques like hitting the law books and
		periodicals archives at the library.`,
			ValDesc: map[string]string{
				"student":   "You can parse a broad Web search for clues.",
				"college":   "Police officer",
				"masters":   "Private detective",
				"doctorate": "Federal agent",
				"scholar":   "Sherlock Holmes",
			},
		},
		"law": {
			BaseDesc: `The Law Knowledge represents a knowledge of both legal
		statutes and proper procedures for enforcing them. Law can be useful
		for filing suits, avoiding lawsuits, or getting out of jail. What’s
		more, the Kindred keep their own laws, and more than one vampire has
		saved his own unlife by deftly exploiting a loophole in one of the
		Traditions.`,
			ValDesc: map[string]string{
				"student":   " You know whether to plead guilty, not guilty, or nolo contendere.",
				"college":   "You’re either studying for or just passed the bar exam.",
				"masters":   "You can make a living of the practice, and probably do.",
				"doctorate": "If you’re not partner yet, you will be soon.",
				"scholar":   "You could find the loopholes in the Devil’s contracts.",
			},
		},
		"medicine": {
			BaseDesc: `You have an understanding of how the human body — and to
		a lesser extent the vampiric body — works.  This Ability covers
		knowledge of medicines, ailments, first-aid procedures, and
		diagnosis or treatment of disease. Medicine is of great use to those
		Kindred with an interest in repairing, damaging, or reworking the
		human body.`,
			ValDesc: map[string]string{
				"student":   "You’ve taken a CPR course.",
				"college":   "Premed or paramedic",
				"masters":   "General practitioner",
				"doctorate": "You can perform transplants.",
				"scholar":   "You are respected by the world’s medical community as a pioneer.",
			},
		},
		"occult": {
			BaseDesc: `You are knowledgeable in occult areas such as mysticism,
		curses, magic, folklore, and particularly vampire lore. Unlike most
		other Knowledges, Occult does not imply a command of hard facts.
		Much of what you know may well be rumor, myth, speculation, or
		hearsay. However, the secrets to be learned in this field are worth
		centuries of sifting legend from fact. High levels of Occult imply a
		deep understanding of vampire lore, as well as a good grounding in
		other aspects of the occult. At the very least, you can discern what
		is patently false.`,
			ValDesc: map[string]string{
				"student":   "You’ve got a blog about the eerie and the disturbing.",
				"college":   "There seems to be some unsettling truth to some of the rumors you’ve heard.",
				"masters":   "You’ve heard a lot and actually seen a little for yourself.",
				"doctorate": "You can recognize blatantly false sources and make educated guesses about the rest.",
				"scholar":   "You know most of the basic truths about the hidden world.",
			},
		},
		"politics": {
			BaseDesc: `You are familiar with the politics of the moment,
			including the people in charge and how they got there.  This
			Knowledge can aid you in dealing with or influencing mortal
			politicians, or even offer some insight into the local Cainite
			power structure. The Politics Knowledge includes the ability to
			practically navigate various bureaucracies, as it assumes that
			certain organizational structures and relationship currencies
			are universal.`,
			ValDesc: map[string]string{
				"student":   "Activist; you can pay aspeeding ticket online.",
				"college":   "Political science major; you know how to file a request for information.",
				"masters":   "Campaign manager or talk-radio host; the clerk will help you navigate the forms you need to complete and tell you who needs the duplicates.",
				"doctorate": "Senator; “We’re not supposed to show this to anyone without press credentials, so don’t quote me.”",
				"scholar":   "You could choose the next President of the United States.",
			},
		},
		"science": {
			BaseDesc: `You have at least a basic understanding of most of the 
			physical sciences, such as chemistry, biology, physics, 
			and geology. This Knowledge can be put to all forms 
			of practical use.`,
			ValDesc: map[string]string{
				"student":   "You know most of the high-school basics.",
				"college":   "You’re familiar with the major theories.",
				"masters":   "You could teach high-school science.",
				"doctorate": "You’re fully capable of advancing the knowledge in your field.",
				"scholar":   "Your Nobel Prize is waiting for you.",
			},
		},
		"technology": {
			BaseDesc: `The Technology Knowledge represents a broad acumen
			with electronics, computer hardware, and devices more elaborate
			than “machines,” which fall under the Crafts Skill. If it has a
			processor, a transistor, or an integrated circuit — if it’s
			electronic rather than electrical — manipulating it uses the
			Technology Knowledge.  This is the wide-ranging Ability used to
			build one’s own computer, install (or subvert) a security
			system, repair a mobile phone, or kitbash a shortwave radio.`,
			ValDesc: map[string]string{
				"student":   "You can perform simple modifications or repairs.",
				"college":   "You could make your living in assembly or repair.",
				"masters":   "You can design new technologies from a set of objective requirements.",
				"doctorate": "For you, it’s not, “Can this be done?” but “How can this be done?”",
				"scholar":   "A visionary in the field; you shape how people interact with their world through devices.",
			},
		},
		"expert knowledge": {
			BaseDesc: `Like Hobby Talent and Professional Skill, this is a 
			catchall category. An Expert Knowledge is anything that is
			primarily intellectual or mental in nature and must be
			studied.`,
			ValDesc: map[string]string{
				"student":   "You’ve taken an undergraduate course or read a few books.",
				"college":   "You may have minored in the field.",
				"masters":   "You might hold a degree and are well versed in what’s been written.",
				"doctorate": "You are well-versed in what hasn’t been written.",
				"scholar":   "You know the hidden mysteries of your field and are a veritable font of information.",
			},
		},
	},
	"background": {
		"allies": {
			BaseDesc: `Allies are mortals who support and help you — fami-
			ly, friends, or even a mortal organization that owes you
			some loyalty. Although allies aid you willingly, with-
			out coaxing or coercion, they are not always available
			to offer assistance; they have their own concerns and
			can do only so much for the sake of your relationship.
			However, they might have some useful Background
			Traits of their own, and could provide you with indi-
			rect access to their contacts, influence, or resources.`,
			ValDesc: map[string]string{
				"1": "One ally of moderate influence and power",
				"2": "Two allies, both of moderate power",
				"3": "Three allies, one of whom is quite influential",
				"4": "Four allies, one of whom is very influential",
				"5": "Five allies, one of whom is extremely influential",
			},
		},
		"alternateIdentity": {
			BaseDesc: `You maintain an alternate identity, complete with
			papers, birth certificates, or any other documentation
			you desire. Only a few may know your real name or
			identity. Your alternate persona may be highly in-
			volved in organized crime, a member of the opposite
			Sect, a con artist who uses alternate identities for her
			game, or you may simply gather information about the
			enemy. Indeed, some vampires may know you as one
			individual while others believe you to be someone else
			entirely.`,
			ValDesc: map[string]string{
				"1": `You are new at this identity game.
				Sometimes you slip and forget your
				other persona.`,
				"2": `You are well grounded in your
				alternate identity. You are
				convincing enough to play the part
				of a doctor, lawyer, funeral salesman,
				drug-smuggler, or a capable spy.`,
				"3": `You have a fair reputation as your
				alternate persona and get name-
				recognition in the area where you
				have infiltrated`,
				"4": `Your alternate identity has respect
				and trust within your area of
				infiltration.`,
				"5": `You command respect in your area of
				infiltration, and you may even have
				accumulated a bit of influence. You
				have the trust (or at least the
				recognition) of many powerful
				individuals within your area.`,
			},
		},
		"blackHandMembership": {
			BaseDesc: `You are a member of the feared Black Hand, the
			body of soldiers and assassins that serves the Sabbat
			fervently. You may call upon members of the Black Hand to
			aid you, should you ever need it. You may find yourself
			assigned to perform assassinations, lend martial aid, or
			even further the political ends of the Hand as a diplo-
			mat or spy. You may also be required to attend crusades
			that take you away from your pack. All members of the
			Black Hand must heed the call of another Hand mem-
			ber, especially the superiors of the faction.
			Being a member of the Black Hand is a prestigious
			matter, when dealing with other Sabbat, should
			you choose to reveal your affiliation you may add your rating in this Background to any
			Social dice pools. Most Hand members,
			however, choose not to reveal their allegiance.`,
			ValDesc: map[string]string{
				"1": `You are a grunt; you may call upon
				one Black Hand member once per
				story.`,
				"2": `You are known and respected in the
				Black Hand; you may call upon two
				Black Hand members once per story.`,
				"3": `You are held in the Black Hand’s
				regard; you may call upon five Black
				Hand members once per story.`,
				"4": `You are a hero among members of
				the Black Hand; you may call upon
				seven Black Hand members twice per
				story (if it seems you’re becoming
				soft, you may lose points in this
				Background). You may also lead large
				numbers of Hand members into action
				should it ever become necessary.`,
				"5": `You are part of Black Hand legend;
				you may call upon 12 Black Hand
				members twice per story (but see the
				preceding caution). You may also lead
				large numbers of Hand members into
				action should it ever become
				necessary. The Seraphim may even
				seek your counsel on matters of
				import.`,
			},
		},
		"contacts": {
			BaseDesc: `You know people all over the city. When you start
			making phone calls around your network, the amount
			of information you can dig up is impressive. Rather
			than friends you can rely on to help you, like Allies,
			Contacts are largely people whom you can bribe, ma-
			nipulate, or coerce into offering information. You also
			have a few major Contacts — associates who can give
			you accurate information in their fields of expertise.
			You should describe each major contact in some detail
			before the game begins.
			In addition to your major contacts, you also have a
			number of minor contacts spread throughout the city.
			Your major contact might be in the district attorney’s
			office, while your minor contacts might include beat
			cops, DMV clerks, club bouncers, or members of an
			online social network. You don’t need to detail these
			various “passing acquaintances” before play.`,
			ValDesc: map[string]string{
				"1": `One major contact`,
				"2": `Two major contacts`,
				"3": `Three major contacts`,
				"4": `Four major contacts`,
				"5": ` Five major contacts`,
			},
		},
		"domain": {
			BaseDesc: `Domain is physical territory (usually within the
				chronicle’s central city) to which your character con-
				trols access for the purpose of feeding. Some Kindred
				refer to their domain as hunting grounds, and most
				jealously guard their domains, even invoking the Tra-
				dition of the same name to protect their claims. As
				part of this Background, the character’s claim to the
				domain is recognized by the Prince or some other Kin-
				dred authority in the city where it is located.
				The Kindred who claims the domain can’t keep the
				living inhabitants from going about their business,
				nor does she exercise any direct influence over them,
				but she can keep watch herself and mind their com-
				ings and goings. She can also have Allies or Retainers
				specifically look for unfamiliar vampires and alert her
				when they find some.Each level of Domain reduces the difficulty of hunt-
				ing checks by one for your character and those whom
				the character allows in.`,
			ValDesc: map[string]string{
				"1": "A single small building, such as a single-family home or a social establishment — enough for a basic haven.",
				"2": `A church, factory, warehouse,
				mid-rise, or other large structure —
				a location with ready but easily
				controllable access to the outside
				world.`,
				"3": `A high-rise, city block, or an
				important intersection — a
				location or area that offers areas for
				concealment as well as controlled
				access.`,
				"4": `A sewer subsection, a network of
				service tunnels, the enclave of homes
				on a hill overlooking the city —
				a place with inherently protective
				features, such as an isolated mountain
				road, bridge-only access, or vigilant
				private security force.`,
				"5": `An entire neighborhood, an ethnic
				subdivision like “Chinatown” or
				“Little Italy,” or a whole suburb`,
			},
		},
		"fame": {
			BaseDesc: `You enjoy widespread recognition in mortal society,
			perhaps as an entertainer, writer, or athlete. People
			may enjoy just being seen with you. This gives you all
			manner of privileges when moving in mortal society,
			but can also attract an unwanted amount of attention
			now that you’re no longer alive. The greatest weapon
			fame has to offer is the ability to sway public opin-
			ion — as modern media constantly proves. Fame isn’t
			always tied to entertainment: A heinous criminal in
			a high-profile trial probably has a certain amount of
			fame, as do a lawmaker and a scientist who has made a
			popularized discovery.`,
			ValDesc: map[string]string{
				"1": `You’re known to a select subculture
				— local club-goers, industry bloggers,
				or the Park Avenue set, for instance`,
				"2": `Random people start to recognize
				your face; you’re a minor celebrity
				such as a small-time criminal or a
				local news anchor.`,
				"3": `You have greater renown; perhaps
				you’re a senator or an entertainer who
				regularly gets hundreds of thousands
				of YouTube hits.`,
				"4": `A full-blown celebrity; your name is
				often recognized by the average
				person on the street.`,
				"5": `You’re a household word. People
				name their children after you.`,
			},
		},
		"generation": {
			BaseDesc: `This Background represents your Generation: the
			purity of your blood, and your proximity to the First
			Vampire. A high Generation rating may represent
			a powerful sire or a decidedly dangerous taste for di-
			ablerie. If you don’t take any dots in this Trait, you
			begin play as a Thirteenth Generation vampire.`,
			ValDesc: map[string]string{
				"1": `Twelfth Generation: 11 blood pool,
				can spend 1 blood point per turn`,
				"2": `Eleventh Generation: 12 blood pool,
				can spend 1 blood point per turn`,
				"3": `Tenth Generation: 13 blood pool,
				can spend 1 blood point per turn`,
				"4": `Ninth Generation: 14 blood pool,
				can spend 2 blood points per turn`,
				"5": `Eighth Generation: 15 blood pool,
				can spend 3 blood points per turn`,
			},
		},
		"herd": {
			BaseDesc: `You have built a group of mortals from whom you can
			feed without fear. A herd may take many forms, from
			circles of kinky clubgoers to actual cults built around
			you as a god-figure. In addition to providing nourishment, your herd might come in handy for minor tasks,
			though they are typically not very controllable, closely
			connected to you, or particularly skilled (for more effective pawns, purchase Allies or Retainers). Your
			Herd rating adds dice to your rolls for hunting`,
			ValDesc: map[string]string{
				"1": `Three vessels`,
				"2": `Seven vessels`,
				"3": `15 vessels`,
				"4": `30 vessels`,
				"5": `60 vessels`,
			},
		},
		"influence": {
			BaseDesc: `You have pull in the mortal community, whether
			through wealth, prestige, political office, blackmail, or
			supernatural manipulation. Kindred with high Influence can sway, and in rare cases even control, the political and social processes of human society. Influence
			represents the sum of your opinion- or policy-swaying
			power in your community, particularly among the police and bureaucracy. In some cases, cultivating Influence is a path to generating Resources (see below).
			Some rolls may require you to use Influence in place
			of an Ability, particularly when attempting to sway minor bureaucrats. It’s easier to institute sweeping chang-
			es on a local level than a worldwide scale (e.g., having
			an “abandoned” building demolished is relatively easy,
			while starting a war is a bit more difficult).`,
			ValDesc: map[string]string{
				"1": `Moderately influential; a factor in city
				politics`,
				"2": `Well-connected; a force in state
				politics`,
				"3": `Position of influence; a factor in
				regional politics`,
				"4": `Broad personal power; a force in
				national politics`,
				"5": `Vastly influential; a factor in global
				politics`,
			},
		},
		"mentor": {
			BaseDesc: `This Trait represents a Kindred or group of Kindred
			who looks out for you, offering guidance or aid once in
			a while. A mentor may be powerful, but his power need
			not be direct. Depending on the number of dots in this
			Background, your mentor might be nothing more than
			a vampire with a remarkable information network, or
			might be a centuries-old creature with tremendous in-
			fluence and supernatural power. He may offer advice,
			speak to the Prince or Archbishop on your behalf,
			steer other elders clear of you, or warn you when you’re
			walking into situations you don’t understand.
			Most often your mentor is your sire, but it could well
			be any Cainite with an interest in your wellbeing. A
			high Mentor rating could even represent a group of
			like-minded vampires, such as the elders of the city’s
			Tremere chantry or a Black Hand cell.
			`,
			ValDesc: map[string]string{
				"1": `Mentor is an ancilla of little
				influence, or a Ductus or Pack Priest.
				•• Mentor is respected: an elder or
				highly-decorated veteran,
				for instance.`,
				"2": `Mentor is heavily influential, such as
				a member of the Primogen or
				a Bishop.`,
				"3": `Mentor is heavily influential, such as
				a member of the Primogen or
				a Bishop.`,
				"4": `Mentor has a great deal of power over
				the city: a Prince or Archbishop,
				for example.`,
				"5": `Mentor is extraordinarily powerful,
				perhaps even a Justicar or Cardinal.`,
			},
		},
		"resources": {
			BaseDesc: `Resources are valuable goods whose disposition your
			character controls. These assets may be actual cash,
			but as this Background increases, they’re more likely
			to be investments, property, or earning capital of some
			sort — land, industrial assets, stocks and bonds, com-
			mercial inventories, criminal infrastructure, contra-
			band, even taxes or tithes. Remember that vampires
			don’t need to arrange for any food except blood and
			their actual needs (as opposed to wants) for shelter are
			very easily accommodated. Resources for vampires go
			mostly to pay for luxuries and the associated expenses
			of developing and maintaining Status, Influence, and
			other Backgrounds. A character with no dots in Re-
			sources may have enough clothing and supplies to get
			by, or she may be destitute and squatting in a refrigera-
			tor box under an overpass.`,
			ValDesc: map[string]string{
				"1": `Sufficient. You can maintain a typical
				residence in the style of the working
				class with stability, even if spending
				sprees come seldom.`,
				"2": `Moderate. You can display yourself
				as a member in good standing of the
				middle class, with the occasional gift
				and indulgence seemly for a person
				of even higher station. You can
				maintain a servant or hire specific
				help as necessary. A fraction of your
				resources are available in cash,
				readily portable property (like jewelry
				or furniture), and other valuables
				(such as a car or modest home) that
				let you maintain a standard of living
				at the one-dot level wherever you
				happen to be, for up to six months.`,
				"3": `Comfortable. You are a prominent
				and established member of your
				community, with land and an owned
				dwelling, and you have a reputation
				that lets you draw on credit at very
				generous terms. You likely have more
				tied up in equity and property than
				you do in ready cash. You can
				maintain a one-dot quality of
				existence wherever you are without
				difficulty, for as long as you choose.`,
				"4": `Wealthy. You rarely touch cash, as
				most of your assets exist in tangible
				forms that are themselves more
				valuable and stable than paper
				money. You hold more wealth than
				many of your local peers (if they can
				be called such a thing). When
				earning your Resources doesn’t
				enjoy your usual degree of attention,
				you can maintain a three-dot
				existence for up to a year, and a two-
				dot existence indefinitely.`,
				"5": `Extremely Wealthy. You are the
				model to which others strive to
				achieve, at least in the popular mind.
				Television shows, magazine spreads,
				and gossip websites speculate about
				your clothing, the appointments of
				your numerous homes, and the luxury
				of your modes of transportation. You
				have vast and widely distributed
				assets, perhaps tied to the fates of
				nations, each with huge staffs and
				connections to every level of
				society through a region. You
				travel with a minimum of three-dot
				comforts, more with a little effort.
				Corporations and governments
				sometimes come to you to buy into
				stocks or bond programs.`,
			},
		},
		"retainers": {
			BaseDesc: `Not precisely Allies or Contacts, your retainers are ser-
			vants, assistants, or other people who are your loyal and
			steadfast companions. Many vampires’ servants are ghouls their supernatural powers and blood bond-en-
			forced loyalty make them the servants of choice. Retainers
			may also be people whom you’ve repeatedly Dominated
			until they have no free will left, or followers so enthralled
			with your Presence that their loyalty borders on blind fa-
			naticism. Some vampires, particularly those with the Ani-
			malism Discipline, use animal ghouls as retainers.
			You must maintain some control over your retainers,
			whether through a salary, the gift of your vitae, or the
			use of Disciplines. Retainers are never “blindly loyal no
			matter what” — if you treat them poorly without exer-
			cising strict control, they might well turn on you.
			Retainers may be useful, but they should never be
			flawless.`,
			ValDesc: map[string]string{
				"1": "One retainer",
				"2": "Two retainers",
				"3": "Three retainers",
				"4": "Four retainers",
				"5": "Five retainers",
			},
		},
		"rituals": {
			BaseDesc: `This Background is for Sabbat characters only.
			You know the ritae and rituals of the Sabbat, and you
			can enact many of them. This Background is vital to be-
			ing a Pack Priest — without this Background, ritae will
			not function. This Background is actually a supernatural
			investment, drawing on the magic of the eldest Tzimisce
			sorcerers. Sabbat vampires who are not their pack‘s
			priests should have an outstanding reason for acquiring
			this Background, as Pack Priests are loath to share their
			secrets with more secular members of the Sect.`,
			ValDesc: map[string]string{
				"1": `You know a few of the auctoritas ritae
				(your choice).`,
				"2": `You know some of the auctoritas ritae
				(your choice) and a few ignoblis ritae
				(your choice).`,
				"3": `You know all of the auctoritas ritae
				and some ignoblis ritae (your choice).
				Also, you may create your own
				ignoblis ritae, given enough time
				(consult your Storyteller for
				development time and game effects).`,
				"4": `You know all the auctoritas ritae
				and many ignoblis ritae (your
				choice). You may create your own
				ignoblis ritae, given enough time
				(consult your Storyteller for
				development time and game effects).
				You are also familiar with the
				functions of numerous regional and
				pack-specific ignoblis ritae, even if
				you cannot perform them.`,
				"5": `You know all the auctoritas ritae and
				dozens of ignoblis ritae (your choice).
				You may create your own ignoblis
				ritae, given enough time (consult
				your Storyteller for development time
				and game effects). You are also
				familiar with the functions of almost
				all regional and pack-specific
				ignoblis ritae, even if you cannot
				perform them; if it’s been written
				down or passed around in lore,
				you’ve heard of it`,
			},
		},
		"status": {
			BaseDesc: `You have something of a reputation and standing
			(earned or unearned) within the local community of
			Kindred. Status among Camarilla society is as often
			derived from your sire’s status and the respect due your
			particular bloodline as it is by personal achievement.
			Among the Sabbat, status is more likely to stem from
			the reputation of your pack or the zeal of your outlook.
			Elders are known for having little respect for their ju-
			niors; this Background can mitigate that somewhat`,
			ValDesc: map[string]string{
				"1": "Known: a neonate/Pack Priest",
				"2": "Respected: an ancilla/respected Ductus",
				"3": "Influential: an elder/Templar",
				"4": `Powerful: a member of the Primogen/
				a Bishop`,
				"5": "Luminary: a Prince/Archbishop",
			},
		},
	},
}

var VirtueMap = map[string]Characteristic{
	"conscience": {
		BaseDesc: `Conscience is a Trait that allows characters to evaluate their conduct with relation to what is “right” and
		“wrong.” A character’s moral judgment with Conscience
		stems from her attitude and outlook. Conscience is
		what prevents a vampire from succumbing to the Beast,
		by defining the Beast’s urges as unacceptable.`,
		ValDesc: map[string]string{
			"1": "Uncaring",
			"2": "Normal",
			"3": "Ethical",
			"4": "Righteous",
			"5": "Remorseful",
		},
	},
	"selfControl": {
		BaseDesc: `Self-Control defines a character’s discipline and mas-
		tery over the Beast. Characters with high Self-Control
		rarely succumb to emotional urges, and are thus able to
		restrain their darker sides more readily than characters
		with low Self-Control.`,
		ValDesc: map[string]string{
			"1": "Unstable",
			"2": "Normal",
			"3": "Temperate",
			"4": "Hardened",
			"5": "Total self-mastery",
		},
	},
	"courage": {
		BaseDesc: `All characters have a Courage Trait, regardless of
		the Path they follow. Courage is the quality that allows
		characters to stand in the face of fear or daunting ad-
		versity. It is bravery, mettle, and stoicism combined. A
		character with high Courage meets her fears head-on,
		while a character of lesser Courage may flee in terror.
		Kindred use the Courage Virtue when faced with cir-
		cumstances they endemically dread: fire, sunlight, True
		Faith.`,
		ValDesc: map[string]string{
			"1": "Timid",
			"2": "Normal",
			"3": "Bold",
			"4": "Resolute",
			"5": "Heroic",
		},
	},
}

var EntryMap = map[string]Entry{
	"lextalionis": {
		Name: "Lextalionis",
		Description: `The dreaded blood hunt. A Camarilla Prince declares the hunt in his court, and the word filters down through his Primogen to the Clans. All who hear the call must participate (although “participation” can simply mean staying out of the way whenthe hunters come through). Once in a while such a
		hunt is called to hound the criminal Kindred into exile outside the city’s borders, but more often than not
		the hunt only ends when the hunted suffers the Final
		Death. Some Princes even turn a blind eye to the act
		of Amaranth during such a hunt.`,
	},
	"shovelHead": {
		Name: "Shovel Head Method",
		Description: `Used during times of war, this infamous method
		consists of collecting a number of victims, Embracing
		them with the tiniest quantity of blood possible, bashing them over the head with a shovel (to knock them unconscious before they frenzy), and burying them in a mass grave. The newly Embraced Cainites rouse quickly, and they must dig themselves out of the grave to sate their frenzy, often at the expense of the weaker vampires entombed with them. The Sabbat does not even consider vampires made this way true vampires yet, and it has little reservation against throwing legions of these frenzied horrors against their foes.
		`,
	},
	"rotschreck": {
		Name:        "Rötschreck",
		Description: ``,
	},
	"Whither Mortis": {
		Name: "Rötschreck",
		Description: `Lost Clans and bloodlines, such as the Cap-
		padocians and the Lamia, had access to an
		ancient Discipline known as Mortis. Some
		Kindred scholars claim that Mortis and Nec-
		romancy are distinct Disciplines, but for ease
		the three Mortis paths presented in this book
		are listed as Necromancy paths. The Cappa-
		docians specialized in either the
		Corpse in the Monster or the Grave’s Decay,
		while the Lamia generally took
		the Path of the Four Humors as their primary
		path. Rumors are that the Harbingers of Skulls
		in the Sabbat have relearned the
		Corpse in the Monster and the Grave’s Decay,
		claiming them collectively as the “Mortuus
		Path,” but they still tend to follow most mod-
		ern necromancers and choose Sepulchre Path
		as their primary path before learning Grave’s
		Decay or the Corpse in the Monster.`,
	},
	"eyeContact": {
		Name: "Eye Contact",
		Description: `The need for eye contact stems from the aggressor Kindred’s need to see his victim’s soul,
		and the eyes are the traditionally known as
		the windows to the soul. While the vampire
		needs to capture his target’s attention, the
		target’s eyes need not be present for such
		a power to work (although the arts of the
		Tzimisce make this somewhat challenging at
		times) — they only need to find the soul of
		his victim laid bare.
		A target trying to avoid eye contact can make
		a Willpower roll against a difficulty equal to
		Dominate user’s Manipulation + Intimida-
		tion (or other appropriate combination for
		other Disciplines or specific situations). The difficulty may be
		reduced for mitigating factors: -1 in the case
		of the target obscuring his eyes slightly (such
		as closing her eyes or wearing dark sunglasses)
		up to a -3 for the eyes being completely un-
		seen (such as with a thick blindfold or having
		her eyes torn out).`,
	},
	"vaulderie": {
		Name: "The Vaulderie",
		Description: `The vampires of the Sabbat take their nightly struggle seriously — so seriously that they tolerate no dissent in their ranks. From the lowliest new recruit to the most
		exalted Cainite, the Sabbat ensure loyalty to one another through a bloody rite known as the Vaulderie.
		The Vaulderie is similar to a blood bond, though it
		differs in intent and function. No Sabbat would ever
		voluntarily succumb to a blood bond, reasoning that
		such bonds are the tools the elders use to enslave their
		childer. Rather, the Sabbat swear the Vaulderie to
		each other, bonding themselves to the pack instead of
		an individual, and, thus, to the Sabbat’s greater cause.`,
	},
}

func ValDescString(val map[string]string) []string {
	var arr []string

	for k, v := range val {
		arr = append(arr, fmt.Sprintf("%s//%s", k, v))
	}

	return arr
}

func CleanDesc(val string) string {
	val = strings.ReplaceAll(val, "\n", " ")
	val = strings.ReplaceAll(val, "\t", "")
	return val
}

