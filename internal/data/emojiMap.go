// Package data provide emoji data
package data

// NOTE: THIS FILE WAS CREATED BY
// EMOJIDATA GENERATION TOOL. (github.com/zajann/emoji/cmd/emojiDataGenerator)
// DO NOT EDIT. IF EDIT THIS FILE THE PROGRAM
// MIGHT NOT WORK PROPERLY.

type EmojiMap map[string]Code

type Code struct {
	Unicode string
	TagMap  map[string]struct{}
}

var Emojis = EmojiMap{
	":grinning:": Code{
		"\U0001f600",
		map[string]struct{}{
			"face":          struct{}{},
			"grin":          struct{}{},
			"grinning face": struct{}{},
		},
	},
	":grin:": Code{
		"\U0001f601",
		map[string]struct{}{
			"beaming face with smiling eyes": struct{}{},
			"eye":                            struct{}{},
			"face":                           struct{}{},
			"grin":                           struct{}{},
			"smile":                          struct{}{},
		},
	},
	":joy:": Code{
		"\U0001f602",
		map[string]struct{}{
			"face":                   struct{}{},
			"face with tears of joy": struct{}{},
			"joy":                    struct{}{},
			"laugh":                  struct{}{},
			"tear":                   struct{}{},
		},
	},
	":rofl:": Code{
		"\U0001f923",
		map[string]struct{}{
			"face":                          struct{}{},
			"floor":                         struct{}{},
			"laugh":                         struct{}{},
			"rolling":                       struct{}{},
			"rolling on the floor laughing": struct{}{},
		},
	},
	":smiley:": Code{
		"\U0001f603",
		map[string]struct{}{
			"face":                        struct{}{},
			"grinning face with big eyes": struct{}{},
			"mouth":                       struct{}{},
			"open":                        struct{}{},
			"smile":                       struct{}{},
		},
	},
	":sweat_smile:": Code{
		"\U0001f605",
		map[string]struct{}{
			"cold":                     struct{}{},
			"face":                     struct{}{},
			"grinning face with sweat": struct{}{},
			"open":                     struct{}{},
			"smile":                    struct{}{},
			"sweat":                    struct{}{},
		},
	},
	":laughing:": Code{
		"\U0001f606",
		map[string]struct{}{
			"face":                    struct{}{},
			"grinning squinting face": struct{}{},
			"laugh":                   struct{}{},
			"mouth":                   struct{}{},
			"open":                    struct{}{},
			"satisfied":               struct{}{},
			"smile":                   struct{}{},
		},
	},
	":wink:": Code{
		"\U0001f609",
		map[string]struct{}{
			"face":         struct{}{},
			"wink":         struct{}{},
			"winking face": struct{}{},
		},
	},
	":blush:": Code{
		"\U0001f60a",
		map[string]struct{}{
			"blush":                          struct{}{},
			"eye":                            struct{}{},
			"face":                           struct{}{},
			"smile":                          struct{}{},
			"smiling face with smiling eyes": struct{}{},
		},
	},
	":yum:": Code{
		"\U0001f60b",
		map[string]struct{}{
			"delicious":          struct{}{},
			"face":               struct{}{},
			"face savoring food": struct{}{},
			"savouring":          struct{}{},
			"smile":              struct{}{},
			"um":                 struct{}{},
			"yum":                struct{}{},
		},
	},
	":sunglasses:": Code{
		"\U0001f60e",
		map[string]struct{}{
			"bright":                       struct{}{},
			"cool":                         struct{}{},
			"eye":                          struct{}{},
			"eyewear":                      struct{}{},
			"face":                         struct{}{},
			"glasses":                      struct{}{},
			"smile":                        struct{}{},
			"smiling face with sunglasses": struct{}{},
			"sun":                          struct{}{},
			"sunglasses":                   struct{}{},
		},
	},
	":heart_eyes:": Code{
		"\U0001f60d",
		map[string]struct{}{
			"eye":                          struct{}{},
			"face":                         struct{}{},
			"love":                         struct{}{},
			"smile":                        struct{}{},
			"smiling face with heart-eyes": struct{}{},
		},
	},
	":kissing_heart:": Code{
		"\U0001f618",
		map[string]struct{}{
			"face":                struct{}{},
			"face blowing a kiss": struct{}{},
			"kiss":                struct{}{},
		},
	},
	":kissing:": Code{
		"\U0001f617",
		map[string]struct{}{
			"face":         struct{}{},
			"kiss":         struct{}{},
			"kissing face": struct{}{},
		},
	},
	":kissing_smiling_eyes:": Code{
		"\U0001f619",
		map[string]struct{}{
			"eye":                            struct{}{},
			"face":                           struct{}{},
			"kiss":                           struct{}{},
			"kissing face with smiling eyes": struct{}{},
			"smile":                          struct{}{},
		},
	},
	":kissing_closed_eyes:": Code{
		"\U0001f61a",
		map[string]struct{}{
			"closed":                        struct{}{},
			"eye":                           struct{}{},
			"face":                          struct{}{},
			"kiss":                          struct{}{},
			"kissing face with closed eyes": struct{}{},
		},
	},
	":relaxed:": Code{
		"\u263a",
		map[string]struct{}{
			"face":         struct{}{},
			"outlined":     struct{}{},
			"relaxed":      struct{}{},
			"smile":        struct{}{},
			"smiling face": struct{}{},
		},
	},
	":slight_smile:": Code{
		"\U0001f642",
		map[string]struct{}{
			"face":                  struct{}{},
			"slightly smiling face": struct{}{},
			"smile":                 struct{}{},
		},
	},
	":hugging:": Code{
		"\U0001f917",
		map[string]struct{}{
			"face":         struct{}{},
			"hug":          struct{}{},
			"hugging":      struct{}{},
			"hugging face": struct{}{},
		},
	},
	":star_struck:": Code{
		"\U0001f929",
		map[string]struct{}{
			"eyes":        struct{}{},
			"face":        struct{}{},
			"grinning":    struct{}{},
			"star":        struct{}{},
			"star-struck": struct{}{},
			"starry-eyed": struct{}{},
		},
	},
	":thinking:": Code{
		"\U0001f914",
		map[string]struct{}{
			"face":          struct{}{},
			"thinking":      struct{}{},
			"thinking face": struct{}{},
		},
	},
	":face_with_raised_eyebrow:": Code{
		"\U0001f928",
		map[string]struct{}{
			"distrust":                 struct{}{},
			"face with raised eyebrow": struct{}{},
			"skeptic":                  struct{}{},
			"disapproval":              struct{}{},
			"disbelief":                struct{}{},
			"mild surprise":            struct{}{},
			"scepticism":               struct{}{},
		},
	},
	":neutral_face:": Code{
		"\U0001f610",
		map[string]struct{}{
			"deadpan":      struct{}{},
			"face":         struct{}{},
			"neutral":      struct{}{},
			"neutral face": struct{}{},
		},
	},
	":expressionless:": Code{
		"\U0001f611",
		map[string]struct{}{
			"expressionless":      struct{}{},
			"expressionless face": struct{}{},
			"face":                struct{}{},
			"inexpressive":        struct{}{},
			"unexpressive":        struct{}{},
		},
	},
	":no_mouth:": Code{
		"\U0001f636",
		map[string]struct{}{
			"face":               struct{}{},
			"face without mouth": struct{}{},
			"mouth":              struct{}{},
			"quiet":              struct{}{},
			"silent":             struct{}{},
		},
	},
	":rolling_eyes:": Code{
		"\U0001f644",
		map[string]struct{}{
			"eyes":                   struct{}{},
			"face":                   struct{}{},
			"face with rolling eyes": struct{}{},
			"rolling":                struct{}{},
		},
	},
	":smirk:": Code{
		"\U0001f60f",
		map[string]struct{}{
			"face":          struct{}{},
			"smirk":         struct{}{},
			"smirking face": struct{}{},
		},
	},
	":persevere:": Code{
		"\U0001f623",
		map[string]struct{}{
			"face":             struct{}{},
			"persevere":        struct{}{},
			"persevering face": struct{}{},
		},
	},
	":disappointed_relieved:": Code{
		"\U0001f625",
		map[string]struct{}{
			"disappointed":          struct{}{},
			"face":                  struct{}{},
			"relieved":              struct{}{},
			"sad but relieved face": struct{}{},
			"whew":                  struct{}{},
		},
	},
	":open_mouth:": Code{
		"\U0001f62e",
		map[string]struct{}{
			"face":                 struct{}{},
			"face with open mouth": struct{}{},
			"mouth":                struct{}{},
			"open":                 struct{}{},
			"sympathy":             struct{}{},
		},
	},
	":zipper_mouth:": Code{
		"\U0001f910",
		map[string]struct{}{
			"face":              struct{}{},
			"mouth":             struct{}{},
			"zipper":            struct{}{},
			"zipper-mouth face": struct{}{},
		},
	},
	":hushed:": Code{
		"\U0001f62f",
		map[string]struct{}{
			"face":        struct{}{},
			"hushed":      struct{}{},
			"hushed face": struct{}{},
			"stunned":     struct{}{},
			"surprised":   struct{}{},
		},
	},
	":sleepy:": Code{
		"\U0001f62a",
		map[string]struct{}{
			"face":        struct{}{},
			"sleep":       struct{}{},
			"sleepy face": struct{}{},
		},
	},
	":tired_face:": Code{
		"\U0001f62b",
		map[string]struct{}{
			"face":       struct{}{},
			"tired":      struct{}{},
			"tired face": struct{}{},
		},
	},
	":sleeping:": Code{
		"\U0001f634",
		map[string]struct{}{
			"face":          struct{}{},
			"sleep":         struct{}{},
			"sleeping face": struct{}{},
			"zzz":           struct{}{},
		},
	},
	":relieved:": Code{
		"\U0001f60c",
		map[string]struct{}{
			"face":          struct{}{},
			"relieved":      struct{}{},
			"relieved face": struct{}{},
		},
	},
	":stuck_out_tongue:": Code{
		"\U0001f61b",
		map[string]struct{}{
			"face":             struct{}{},
			"face with tongue": struct{}{},
			"tongue":           struct{}{},
		},
	},
	":stuck_out_tongue_winking_eye:": Code{
		"\U0001f61c",
		map[string]struct{}{
			"eye":                      struct{}{},
			"face":                     struct{}{},
			"joke":                     struct{}{},
			"tongue":                   struct{}{},
			"wink":                     struct{}{},
			"winking face with tongue": struct{}{},
		},
	},
	":stuck_out_tongue_closed_eyes:": Code{
		"\U0001f61d",
		map[string]struct{}{
			"eye":                        struct{}{},
			"face":                       struct{}{},
			"horrible":                   struct{}{},
			"squinting face with tongue": struct{}{},
			"taste":                      struct{}{},
			"tongue":                     struct{}{},
		},
	},
	":drooling_face:": Code{
		"\U0001f924",
		map[string]struct{}{
			"drooling":      struct{}{},
			"drooling face": struct{}{},
			"face":          struct{}{},
		},
	},
	":unamused:": Code{
		"\U0001f612",
		map[string]struct{}{
			"face":          struct{}{},
			"unamused":      struct{}{},
			"unamused face": struct{}{},
			"unhappy":       struct{}{},
		},
	},
	":sweat:": Code{
		"\U0001f613",
		map[string]struct{}{
			"cold":                     struct{}{},
			"downcast face with sweat": struct{}{},
			"face":                     struct{}{},
			"sweat":                    struct{}{},
		},
	},
	":pensive:": Code{
		"\U0001f614",
		map[string]struct{}{
			"dejected":     struct{}{},
			"face":         struct{}{},
			"pensive":      struct{}{},
			"pensive face": struct{}{},
		},
	},
	":confused:": Code{
		"\U0001f615",
		map[string]struct{}{
			"confused":      struct{}{},
			"confused face": struct{}{},
			"face":          struct{}{},
		},
	},
	":upside_down:": Code{
		"\U0001f643",
		map[string]struct{}{
			"face":             struct{}{},
			"upside-down":      struct{}{},
			"upside-down face": struct{}{},
		},
	},
	":money_mouth:": Code{
		"\U0001f911",
		map[string]struct{}{
			"face":             struct{}{},
			"money":            struct{}{},
			"money-mouth face": struct{}{},
			"mouth":            struct{}{},
		},
	},
	":astonished:": Code{
		"\U0001f632",
		map[string]struct{}{
			"astonished":      struct{}{},
			"astonished face": struct{}{},
			"face":            struct{}{},
			"shocked":         struct{}{},
			"totally":         struct{}{},
		},
	},
	":frowning2:": Code{
		"\u2639",
		map[string]struct{}{
			"face":          struct{}{},
			"frown":         struct{}{},
			"frowning face": struct{}{},
		},
	},
	":slight_frown:": Code{
		"\U0001f641",
		map[string]struct{}{
			"face":                   struct{}{},
			"frown":                  struct{}{},
			"slightly frowning face": struct{}{},
		},
	},
	":confounded:": Code{
		"\U0001f616",
		map[string]struct{}{
			"confounded":      struct{}{},
			"confounded face": struct{}{},
			"face":            struct{}{},
		},
	},
	":disappointed:": Code{
		"\U0001f61e",
		map[string]struct{}{
			"disappointed":      struct{}{},
			"disappointed face": struct{}{},
			"face":              struct{}{},
		},
	},
	":worried:": Code{
		"\U0001f61f",
		map[string]struct{}{
			"face":         struct{}{},
			"worried":      struct{}{},
			"worried face": struct{}{},
		},
	},
	":triumph:": Code{
		"\U0001f624",
		map[string]struct{}{
			"face":                      struct{}{},
			"face with steam from nose": struct{}{},
			"triumph":                   struct{}{},
			"won":                       struct{}{},
		},
	},
	":cry:": Code{
		"\U0001f622",
		map[string]struct{}{
			"cry":         struct{}{},
			"crying face": struct{}{},
			"face":        struct{}{},
			"sad":         struct{}{},
			"tear":        struct{}{},
		},
	},
	":sob:": Code{
		"\U0001f62d",
		map[string]struct{}{
			"cry":                struct{}{},
			"face":               struct{}{},
			"loudly crying face": struct{}{},
			"sad":                struct{}{},
			"sob":                struct{}{},
			"tear":               struct{}{},
		},
	},
	":frowning:": Code{
		"\U0001f626",
		map[string]struct{}{
			"face":                          struct{}{},
			"frown":                         struct{}{},
			"frowning face with open mouth": struct{}{},
			"mouth":                         struct{}{},
			"open":                          struct{}{},
		},
	},
	":anguished:": Code{
		"\U0001f627",
		map[string]struct{}{
			"anguished":      struct{}{},
			"anguished face": struct{}{},
			"face":           struct{}{},
		},
	},
	":fearful:": Code{
		"\U0001f628",
		map[string]struct{}{
			"face":         struct{}{},
			"fear":         struct{}{},
			"fearful":      struct{}{},
			"fearful face": struct{}{},
			"scared":       struct{}{},
		},
	},
	":weary:": Code{
		"\U0001f629",
		map[string]struct{}{
			"face":       struct{}{},
			"tired":      struct{}{},
			"weary":      struct{}{},
			"weary face": struct{}{},
		},
	},
	":exploding_head:": Code{
		"\U0001f92f",
		map[string]struct{}{
			"exploding head": struct{}{},
			"shocked":        struct{}{},
		},
	},
	":grimacing:": Code{
		"\U0001f62c",
		map[string]struct{}{
			"face":           struct{}{},
			"grimace":        struct{}{},
			"grimacing face": struct{}{},
		},
	},
	":cold_sweat:": Code{
		"\U0001f630",
		map[string]struct{}{
			"anxious face with sweat": struct{}{},
			"blue":                    struct{}{},
			"cold":                    struct{}{},
			"face":                    struct{}{},
			"mouth":                   struct{}{},
			"open":                    struct{}{},
			"rushed":                  struct{}{},
			"sweat":                   struct{}{},
		},
	},
	":scream:": Code{
		"\U0001f631",
		map[string]struct{}{
			"face":                   struct{}{},
			"face screaming in fear": struct{}{},
			"fear":                   struct{}{},
			"fearful":                struct{}{},
			"munch":                  struct{}{},
			"scared":                 struct{}{},
			"scream":                 struct{}{},
		},
	},
	":flushed:": Code{
		"\U0001f633",
		map[string]struct{}{
			"dazed":        struct{}{},
			"face":         struct{}{},
			"flushed":      struct{}{},
			"flushed face": struct{}{},
		},
	},
	":crazy_face:": Code{
		"\U0001f92a",
		map[string]struct{}{
			"eye":       struct{}{},
			"goofy":     struct{}{},
			"large":     struct{}{},
			"small":     struct{}{},
			"zany face": struct{}{},
		},
	},
	":dizzy_face:": Code{
		"\U0001f635",
		map[string]struct{}{
			"dizzy":      struct{}{},
			"dizzy face": struct{}{},
			"face":       struct{}{},
		},
	},
	":rage:": Code{
		"\U0001f621",
		map[string]struct{}{
			"angry":        struct{}{},
			"face":         struct{}{},
			"mad":          struct{}{},
			"pouting":      struct{}{},
			"pouting face": struct{}{},
			"rage":         struct{}{},
			"red":          struct{}{},
		},
	},
	":angry:": Code{
		"\U0001f620",
		map[string]struct{}{
			"angry":      struct{}{},
			"angry face": struct{}{},
			"face":       struct{}{},
			"mad":        struct{}{},
		},
	},
	":face_with_symbols_over_mouth:": Code{
		"\U0001f92c",
		map[string]struct{}{
			"face with symbols on mouth": struct{}{},
			"swearing":                   struct{}{},
			"cursing":                    struct{}{},
		},
	},
	":mask:": Code{
		"\U0001f637",
		map[string]struct{}{
			"cold":                   struct{}{},
			"doctor":                 struct{}{},
			"face":                   struct{}{},
			"face with medical mask": struct{}{},
			"mask":                   struct{}{},
			"medicine":               struct{}{},
			"sick":                   struct{}{},
		},
	},
	":thermometer_face:": Code{
		"\U0001f912",
		map[string]struct{}{
			"face":                  struct{}{},
			"face with thermometer": struct{}{},
			"ill":                   struct{}{},
			"sick":                  struct{}{},
			"thermometer":           struct{}{},
		},
	},
	":head_bandage:": Code{
		"\U0001f915",
		map[string]struct{}{
			"bandage":                struct{}{},
			"face":                   struct{}{},
			"face with head-bandage": struct{}{},
			"hurt":                   struct{}{},
			"injury":                 struct{}{},
		},
	},
	":nauseated_face:": Code{
		"\U0001f922",
		map[string]struct{}{
			"face":           struct{}{},
			"nauseated":      struct{}{},
			"nauseated face": struct{}{},
			"vomit":          struct{}{},
		},
	},
	":face_vomiting:": Code{
		"\U0001f92e",
		map[string]struct{}{
			"face vomiting": struct{}{},
			"sick":          struct{}{},
			"vomit":         struct{}{},
		},
	},
	":sneezing_face:": Code{
		"\U0001f927",
		map[string]struct{}{
			"face":          struct{}{},
			"gesundheit":    struct{}{},
			"sneeze":        struct{}{},
			"sneezing face": struct{}{},
		},
	},
	":innocent:": Code{
		"\U0001f607",
		map[string]struct{}{
			"angel":                  struct{}{},
			"face":                   struct{}{},
			"fairy tale":             struct{}{},
			"fantasy":                struct{}{},
			"halo":                   struct{}{},
			"innocent":               struct{}{},
			"smile":                  struct{}{},
			"smiling face with halo": struct{}{},
		},
	},
	":cowboy:": Code{
		"\U0001f920",
		map[string]struct{}{
			"cowboy":          struct{}{},
			"cowboy hat face": struct{}{},
			"cowgirl":         struct{}{},
			"face":            struct{}{},
			"hat":             struct{}{},
		},
	},
	":clown:": Code{
		"\U0001f921",
		map[string]struct{}{
			"clown":      struct{}{},
			"clown face": struct{}{},
			"face":       struct{}{},
		},
	},
	":lying_face:": Code{
		"\U0001f925",
		map[string]struct{}{
			"face":       struct{}{},
			"lie":        struct{}{},
			"lying face": struct{}{},
			"pinocchio":  struct{}{},
		},
	},
	":shushing_face:": Code{
		"\U0001f92b",
		map[string]struct{}{
			"quiet":         struct{}{},
			"shush":         struct{}{},
			"shushing face": struct{}{},
		},
	},
	":face_with_hand_over_mouth:": Code{
		"\U0001f92d",
		map[string]struct{}{
			"face with hand over mouth": struct{}{},
			"whoops":                    struct{}{},
			"shock":                     struct{}{},
			"sudden realization":        struct{}{},
			"surprise":                  struct{}{},
		},
	},
	":face_with_monocle:": Code{
		"\U0001f9d0",
		map[string]struct{}{
			"face with monocle": struct{}{},
			"stuffy":            struct{}{},
			"wealthy":           struct{}{},
		},
	},
	":nerd:": Code{
		"\U0001f913",
		map[string]struct{}{
			"face":      struct{}{},
			"geek":      struct{}{},
			"nerd":      struct{}{},
			"nerd face": struct{}{},
		},
	},
	":smiling_imp:": Code{
		"\U0001f608",
		map[string]struct{}{
			"face":                    struct{}{},
			"fairy tale":              struct{}{},
			"fantasy":                 struct{}{},
			"horns":                   struct{}{},
			"smile":                   struct{}{},
			"smiling face with horns": struct{}{},
		},
	},
	":imp:": Code{
		"\U0001f47f",
		map[string]struct{}{
			"angry face with horns": struct{}{},
			"demon":                 struct{}{},
			"devil":                 struct{}{},
			"face":                  struct{}{},
			"fairy tale":            struct{}{},
			"fantasy":               struct{}{},
			"imp":                   struct{}{},
		},
	},
	":japanese_ogre:": Code{
		"\U0001f479",
		map[string]struct{}{
			"creature":   struct{}{},
			"face":       struct{}{},
			"fairy tale": struct{}{},
			"fantasy":    struct{}{},
			"monster":    struct{}{},
			"ogre":       struct{}{},
			"troll":      struct{}{},
		},
	},
	":japanese_goblin:": Code{
		"\U0001f47a",
		map[string]struct{}{
			"creature":   struct{}{},
			"face":       struct{}{},
			"fairy tale": struct{}{},
			"fantasy":    struct{}{},
			"goblin":     struct{}{},
			"monster":    struct{}{},
		},
	},
	":skull:": Code{
		"\U0001f480",
		map[string]struct{}{
			"death":      struct{}{},
			"face":       struct{}{},
			"fairy tale": struct{}{},
			"monster":    struct{}{},
			"skull":      struct{}{},
		},
	},
	":skull_crossbones:": Code{
		"\u2620",
		map[string]struct{}{
			"crossbones":           struct{}{},
			"death":                struct{}{},
			"face":                 struct{}{},
			"monster":              struct{}{},
			"skull":                struct{}{},
			"skull and crossbones": struct{}{},
		},
	},
	":ghost:": Code{
		"\U0001f47b",
		map[string]struct{}{
			"creature":   struct{}{},
			"face":       struct{}{},
			"fairy tale": struct{}{},
			"fantasy":    struct{}{},
			"ghost":      struct{}{},
			"monster":    struct{}{},
		},
	},
	":alien:": Code{
		"\U0001f47d",
		map[string]struct{}{
			"alien":            struct{}{},
			"creature":         struct{}{},
			"extraterrestrial": struct{}{},
			"face":             struct{}{},
			"fairy tale":       struct{}{},
			"fantasy":          struct{}{},
			"monster":          struct{}{},
			"ufo":              struct{}{},
		},
	},
	":space_invader:": Code{
		"\U0001f47e",
		map[string]struct{}{
			"alien":            struct{}{},
			"alien monster":    struct{}{},
			"creature":         struct{}{},
			"extraterrestrial": struct{}{},
			"face":             struct{}{},
			"fairy tale":       struct{}{},
			"fantasy":          struct{}{},
			"monster":          struct{}{},
			"ufo":              struct{}{},
		},
	},
	":robot:": Code{
		"\U0001f916",
		map[string]struct{}{
			"face":       struct{}{},
			"monster":    struct{}{},
			"robot":      struct{}{},
			"robot face": struct{}{},
		},
	},
	":poop:": Code{
		"\U0001f4a9",
		map[string]struct{}{
			"comic":       struct{}{},
			"dung":        struct{}{},
			"face":        struct{}{},
			"monster":     struct{}{},
			"pile of poo": struct{}{},
			"poo":         struct{}{},
			"poop":        struct{}{},
		},
	},
	":smiley_cat:": Code{
		"\U0001f63a",
		map[string]struct{}{
			"cat":               struct{}{},
			"face":              struct{}{},
			"grinning cat face": struct{}{},
			"mouth":             struct{}{},
			"open":              struct{}{},
			"smile":             struct{}{},
		},
	},
	":smile_cat:": Code{
		"\U0001f638",
		map[string]struct{}{
			"cat":                                 struct{}{},
			"eye":                                 struct{}{},
			"face":                                struct{}{},
			"grin":                                struct{}{},
			"grinning cat face with smiling eyes": struct{}{},
			"smile":                               struct{}{},
		},
	},
	":joy_cat:": Code{
		"\U0001f639",
		map[string]struct{}{
			"cat":                        struct{}{},
			"cat face with tears of joy": struct{}{},
			"face":                       struct{}{},
			"joy":                        struct{}{},
			"tear":                       struct{}{},
		},
	},
	":heart_eyes_cat:": Code{
		"\U0001f63b",
		map[string]struct{}{
			"cat":                              struct{}{},
			"eye":                              struct{}{},
			"face":                             struct{}{},
			"love":                             struct{}{},
			"smile":                            struct{}{},
			"smiling cat face with heart-eyes": struct{}{},
		},
	},
	":smirk_cat:": Code{
		"\U0001f63c",
		map[string]struct{}{
			"cat":                     struct{}{},
			"cat face with wry smile": struct{}{},
			"face":                    struct{}{},
			"ironic":                  struct{}{},
			"smile":                   struct{}{},
			"wry":                     struct{}{},
		},
	},
	":kissing_cat:": Code{
		"\U0001f63d",
		map[string]struct{}{
			"cat":              struct{}{},
			"eye":              struct{}{},
			"face":             struct{}{},
			"kiss":             struct{}{},
			"kissing cat face": struct{}{},
		},
	},
	":scream_cat:": Code{
		"\U0001f640",
		map[string]struct{}{
			"cat":            struct{}{},
			"face":           struct{}{},
			"oh":             struct{}{},
			"surprised":      struct{}{},
			"weary":          struct{}{},
			"weary cat face": struct{}{},
		},
	},
	":crying_cat_face:": Code{
		"\U0001f63f",
		map[string]struct{}{
			"cat":             struct{}{},
			"cry":             struct{}{},
			"crying cat face": struct{}{},
			"face":            struct{}{},
			"sad":             struct{}{},
			"tear":            struct{}{},
		},
	},
	":pouting_cat:": Code{
		"\U0001f63e",
		map[string]struct{}{
			"cat":              struct{}{},
			"face":             struct{}{},
			"pouting":          struct{}{},
			"pouting cat face": struct{}{},
		},
	},
	":see_no_evil:": Code{
		"\U0001f648",
		map[string]struct{}{
			"evil":               struct{}{},
			"face":               struct{}{},
			"forbidden":          struct{}{},
			"gesture":            struct{}{},
			"monkey":             struct{}{},
			"no":                 struct{}{},
			"not":                struct{}{},
			"prohibited":         struct{}{},
			"see":                struct{}{},
			"see-no-evil monkey": struct{}{},
		},
	},
	":hear_no_evil:": Code{
		"\U0001f649",
		map[string]struct{}{
			"evil":                struct{}{},
			"face":                struct{}{},
			"forbidden":           struct{}{},
			"gesture":             struct{}{},
			"hear":                struct{}{},
			"hear-no-evil monkey": struct{}{},
			"monkey":              struct{}{},
			"no":                  struct{}{},
			"not":                 struct{}{},
			"prohibited":          struct{}{},
		},
	},
	":speak_no_evil:": Code{
		"\U0001f64a",
		map[string]struct{}{
			"evil":                 struct{}{},
			"face":                 struct{}{},
			"forbidden":            struct{}{},
			"gesture":              struct{}{},
			"monkey":               struct{}{},
			"no":                   struct{}{},
			"not":                  struct{}{},
			"prohibited":           struct{}{},
			"speak":                struct{}{},
			"speak-no-evil monkey": struct{}{},
		},
	},
	":baby:": Code{
		"\U0001f476",
		map[string]struct{}{
			"baby":  struct{}{},
			"young": struct{}{},
		},
	},
	":baby_tone1:": Code{
		"\U0001f476\U0001f3fb",
		map[string]struct{}{
			"baby":            struct{}{},
			"light skin tone": struct{}{},
			"young":           struct{}{},
		},
	},
	":baby_tone2:": Code{
		"\U0001f476\U0001f3fc",
		map[string]struct{}{
			"baby":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"young":                  struct{}{},
		},
	},
	":baby_tone3:": Code{
		"\U0001f476\U0001f3fd",
		map[string]struct{}{
			"baby":             struct{}{},
			"medium skin tone": struct{}{},
			"young":            struct{}{},
		},
	},
	":baby_tone4:": Code{
		"\U0001f476\U0001f3fe",
		map[string]struct{}{
			"baby":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"young":                 struct{}{},
		},
	},
	":baby_tone5:": Code{
		"\U0001f476\U0001f3ff",
		map[string]struct{}{
			"baby":           struct{}{},
			"dark skin tone": struct{}{},
			"young":          struct{}{},
		},
	},
	":child:": Code{
		"\U0001f9d2",
		map[string]struct{}{
			"child":                struct{}{},
			"gender-neutral":       struct{}{},
			"young":                struct{}{},
			"(no specific gender)": struct{}{},
		},
	},
	":child_tone1:": Code{
		"\U0001f9d2\U0001f3fb",
		map[string]struct{}{
			"child":           struct{}{},
			"gender-neutral":  struct{}{},
			"light skin tone": struct{}{},
			"young":           struct{}{},
		},
	},
	":child_tone2:": Code{
		"\U0001f9d2\U0001f3fc",
		map[string]struct{}{
			"child":                  struct{}{},
			"gender-neutral":         struct{}{},
			"medium-light skin tone": struct{}{},
			"young":                  struct{}{},
		},
	},
	":child_tone3:": Code{
		"\U0001f9d2\U0001f3fd",
		map[string]struct{}{
			"child":            struct{}{},
			"gender-neutral":   struct{}{},
			"medium skin tone": struct{}{},
			"young":            struct{}{},
		},
	},
	":child_tone4:": Code{
		"\U0001f9d2\U0001f3fe",
		map[string]struct{}{
			"child":                 struct{}{},
			"gender-neutral":        struct{}{},
			"medium-dark skin tone": struct{}{},
			"young":                 struct{}{},
		},
	},
	":child_tone5:": Code{
		"\U0001f9d2\U0001f3ff",
		map[string]struct{}{
			"child":          struct{}{},
			"dark skin tone": struct{}{},
			"gender-neutral": struct{}{},
			"young":          struct{}{},
		},
	},
	":boy:": Code{
		"\U0001f466",
		map[string]struct{}{
			"boy":   struct{}{},
			"young": struct{}{},
		},
	},
	":boy_tone1:": Code{
		"\U0001f466\U0001f3fb",
		map[string]struct{}{
			"boy":             struct{}{},
			"light skin tone": struct{}{},
			"young":           struct{}{},
		},
	},
	":boy_tone2:": Code{
		"\U0001f466\U0001f3fc",
		map[string]struct{}{
			"boy":                    struct{}{},
			"medium-light skin tone": struct{}{},
			"young":                  struct{}{},
		},
	},
	":boy_tone3:": Code{
		"\U0001f466\U0001f3fd",
		map[string]struct{}{
			"boy":              struct{}{},
			"medium skin tone": struct{}{},
			"young":            struct{}{},
		},
	},
	":boy_tone4:": Code{
		"\U0001f466\U0001f3fe",
		map[string]struct{}{
			"boy":                   struct{}{},
			"medium-dark skin tone": struct{}{},
			"young":                 struct{}{},
		},
	},
	":boy_tone5:": Code{
		"\U0001f466\U0001f3ff",
		map[string]struct{}{
			"boy":            struct{}{},
			"dark skin tone": struct{}{},
			"young":          struct{}{},
		},
	},
	":girl:": Code{
		"\U0001f467",
		map[string]struct{}{
			"girl":   struct{}{},
			"Virgo":  struct{}{},
			"young":  struct{}{},
			"zodiac": struct{}{},
		},
	},
	":girl_tone1:": Code{
		"\U0001f467\U0001f3fb",
		map[string]struct{}{
			"girl":            struct{}{},
			"light skin tone": struct{}{},
			"Virgo":           struct{}{},
			"young":           struct{}{},
			"zodiac":          struct{}{},
		},
	},
	":girl_tone2:": Code{
		"\U0001f467\U0001f3fc",
		map[string]struct{}{
			"girl":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"Virgo":                  struct{}{},
			"young":                  struct{}{},
			"zodiac":                 struct{}{},
		},
	},
	":girl_tone3:": Code{
		"\U0001f467\U0001f3fd",
		map[string]struct{}{
			"girl":             struct{}{},
			"medium skin tone": struct{}{},
			"Virgo":            struct{}{},
			"young":            struct{}{},
			"zodiac":           struct{}{},
		},
	},
	":girl_tone4:": Code{
		"\U0001f467\U0001f3fe",
		map[string]struct{}{
			"girl":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"Virgo":                 struct{}{},
			"young":                 struct{}{},
			"zodiac":                struct{}{},
		},
	},
	":girl_tone5:": Code{
		"\U0001f467\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"girl":           struct{}{},
			"Virgo":          struct{}{},
			"young":          struct{}{},
			"zodiac":         struct{}{},
		},
	},
	":adult:": Code{
		"\U0001f9d1",
		map[string]struct{}{
			"adult":                struct{}{},
			"gender-neutral":       struct{}{},
			"(no specific gender)": struct{}{},
		},
	},
	":adult_tone1:": Code{
		"\U0001f9d1\U0001f3fb",
		map[string]struct{}{
			"adult":           struct{}{},
			"gender-neutral":  struct{}{},
			"light skin tone": struct{}{},
		},
	},
	":adult_tone2:": Code{
		"\U0001f9d1\U0001f3fc",
		map[string]struct{}{
			"adult":                  struct{}{},
			"gender-neutral":         struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":adult_tone3:": Code{
		"\U0001f9d1\U0001f3fd",
		map[string]struct{}{
			"adult":            struct{}{},
			"gender-neutral":   struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":adult_tone4:": Code{
		"\U0001f9d1\U0001f3fe",
		map[string]struct{}{
			"adult":                 struct{}{},
			"gender-neutral":        struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":adult_tone5:": Code{
		"\U0001f9d1\U0001f3ff",
		map[string]struct{}{
			"adult":          struct{}{},
			"dark skin tone": struct{}{},
			"gender-neutral": struct{}{},
		},
	},
	":man:": Code{
		"\U0001f468",
		map[string]struct{}{
			"man": struct{}{},
		},
	},
	":man_tone1:": Code{
		"\U0001f468\U0001f3fb",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"man":             struct{}{},
		},
	},
	":man_tone2:": Code{
		"\U0001f468\U0001f3fc",
		map[string]struct{}{
			"man":                    struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_tone3:": Code{
		"\U0001f468\U0001f3fd",
		map[string]struct{}{
			"man":              struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":man_tone4:": Code{
		"\U0001f468\U0001f3fe",
		map[string]struct{}{
			"man":                   struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_tone5:": Code{
		"\U0001f468\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"man":            struct{}{},
		},
	},
	":woman:": Code{
		"\U0001f469",
		map[string]struct{}{
			"woman": struct{}{},
		},
	},
	":woman_tone1:": Code{
		"\U0001f469\U0001f3fb",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"woman":           struct{}{},
		},
	},
	":woman_tone2:": Code{
		"\U0001f469\U0001f3fc",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
		},
	},
	":woman_tone3:": Code{
		"\U0001f469\U0001f3fd",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"woman":            struct{}{},
		},
	},
	":woman_tone4:": Code{
		"\U0001f469\U0001f3fe",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
		},
	},
	":woman_tone5:": Code{
		"\U0001f469\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"woman":          struct{}{},
		},
	},
	":older_adult:": Code{
		"\U0001f9d3",
		map[string]struct{}{
			"gender-neutral":       struct{}{},
			"old":                  struct{}{},
			"older adult":          struct{}{},
			"(no specific gender)": struct{}{},
		},
	},
	":older_adult_tone1:": Code{
		"\U0001f9d3\U0001f3fb",
		map[string]struct{}{
			"gender-neutral":  struct{}{},
			"light skin tone": struct{}{},
			"old":             struct{}{},
			"older adult":     struct{}{},
		},
	},
	":older_adult_tone2:": Code{
		"\U0001f9d3\U0001f3fc",
		map[string]struct{}{
			"gender-neutral":         struct{}{},
			"medium-light skin tone": struct{}{},
			"old":                    struct{}{},
			"older adult":            struct{}{},
		},
	},
	":older_adult_tone3:": Code{
		"\U0001f9d3\U0001f3fd",
		map[string]struct{}{
			"gender-neutral":   struct{}{},
			"medium skin tone": struct{}{},
			"old":              struct{}{},
			"older adult":      struct{}{},
		},
	},
	":older_adult_tone4:": Code{
		"\U0001f9d3\U0001f3fe",
		map[string]struct{}{
			"gender-neutral":        struct{}{},
			"medium-dark skin tone": struct{}{},
			"old":                   struct{}{},
			"older adult":           struct{}{},
		},
	},
	":older_adult_tone5:": Code{
		"\U0001f9d3\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"gender-neutral": struct{}{},
			"old":            struct{}{},
			"older adult":    struct{}{},
		},
	},
	":older_man:": Code{
		"\U0001f474",
		map[string]struct{}{
			"man":     struct{}{},
			"old":     struct{}{},
			"old man": struct{}{},
		},
	},
	":older_man_tone1:": Code{
		"\U0001f474\U0001f3fb",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"old":             struct{}{},
			"old man":         struct{}{},
		},
	},
	":older_man_tone2:": Code{
		"\U0001f474\U0001f3fc",
		map[string]struct{}{
			"man":                    struct{}{},
			"medium-light skin tone": struct{}{},
			"old":                    struct{}{},
			"old man":                struct{}{},
		},
	},
	":older_man_tone3:": Code{
		"\U0001f474\U0001f3fd",
		map[string]struct{}{
			"man":              struct{}{},
			"medium skin tone": struct{}{},
			"old":              struct{}{},
			"old man":          struct{}{},
		},
	},
	":older_man_tone4:": Code{
		"\U0001f474\U0001f3fe",
		map[string]struct{}{
			"man":                   struct{}{},
			"medium-dark skin tone": struct{}{},
			"old":                   struct{}{},
			"old man":               struct{}{},
		},
	},
	":older_man_tone5:": Code{
		"\U0001f474\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"man":            struct{}{},
			"old":            struct{}{},
			"old man":        struct{}{},
		},
	},
	":older_woman:": Code{
		"\U0001f475",
		map[string]struct{}{
			"old":       struct{}{},
			"old woman": struct{}{},
			"woman":     struct{}{},
		},
	},
	":older_woman_tone1:": Code{
		"\U0001f475\U0001f3fb",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"old":             struct{}{},
			"old woman":       struct{}{},
			"woman":           struct{}{},
		},
	},
	":older_woman_tone2:": Code{
		"\U0001f475\U0001f3fc",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"old":                    struct{}{},
			"old woman":              struct{}{},
			"woman":                  struct{}{},
		},
	},
	":older_woman_tone3:": Code{
		"\U0001f475\U0001f3fd",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"old":              struct{}{},
			"old woman":        struct{}{},
			"woman":            struct{}{},
		},
	},
	":older_woman_tone4:": Code{
		"\U0001f475\U0001f3fe",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"old":                   struct{}{},
			"old woman":             struct{}{},
			"woman":                 struct{}{},
		},
	},
	":older_woman_tone5:": Code{
		"\U0001f475\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"old":            struct{}{},
			"old woman":      struct{}{},
			"woman":          struct{}{},
		},
	},
	":man_health_worker:": Code{
		"\U0001f468\u200d\u2695\ufe0f",
		map[string]struct{}{
			"doctor":            struct{}{},
			"healthcare":        struct{}{},
			"man":               struct{}{},
			"man health worker": struct{}{},
			"nurse":             struct{}{},
			"therapist":         struct{}{},
		},
	},
	":man_health_worker_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\u2695\ufe0f",
		map[string]struct{}{
			"doctor":            struct{}{},
			"healthcare":        struct{}{},
			"light skin tone":   struct{}{},
			"man":               struct{}{},
			"man health worker": struct{}{},
			"nurse":             struct{}{},
			"therapist":         struct{}{},
		},
	},
	":man_health_worker_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\u2695\ufe0f",
		map[string]struct{}{
			"doctor":                 struct{}{},
			"healthcare":             struct{}{},
			"man":                    struct{}{},
			"man health worker":      struct{}{},
			"medium-light skin tone": struct{}{},
			"nurse":                  struct{}{},
			"therapist":              struct{}{},
		},
	},
	":man_health_worker_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\u2695\ufe0f",
		map[string]struct{}{
			"doctor":            struct{}{},
			"healthcare":        struct{}{},
			"man":               struct{}{},
			"man health worker": struct{}{},
			"medium skin tone":  struct{}{},
			"nurse":             struct{}{},
			"therapist":         struct{}{},
		},
	},
	":man_health_worker_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\u2695\ufe0f",
		map[string]struct{}{
			"doctor":                struct{}{},
			"healthcare":            struct{}{},
			"man":                   struct{}{},
			"man health worker":     struct{}{},
			"medium-dark skin tone": struct{}{},
			"nurse":                 struct{}{},
			"therapist":             struct{}{},
		},
	},
	":man_health_worker_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\u2695\ufe0f",
		map[string]struct{}{
			"dark skin tone":    struct{}{},
			"doctor":            struct{}{},
			"healthcare":        struct{}{},
			"man":               struct{}{},
			"man health worker": struct{}{},
			"nurse":             struct{}{},
			"therapist":         struct{}{},
		},
	},
	":woman_health_worker:": Code{
		"\U0001f469\u200d\u2695\ufe0f",
		map[string]struct{}{
			"doctor":              struct{}{},
			"healthcare":          struct{}{},
			"nurse":               struct{}{},
			"therapist":           struct{}{},
			"woman":               struct{}{},
			"woman health worker": struct{}{},
		},
	},
	":woman_health_worker_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\u2695\ufe0f",
		map[string]struct{}{
			"doctor":              struct{}{},
			"healthcare":          struct{}{},
			"light skin tone":     struct{}{},
			"nurse":               struct{}{},
			"therapist":           struct{}{},
			"woman":               struct{}{},
			"woman health worker": struct{}{},
		},
	},
	":woman_health_worker_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\u2695\ufe0f",
		map[string]struct{}{
			"doctor":                 struct{}{},
			"healthcare":             struct{}{},
			"medium-light skin tone": struct{}{},
			"nurse":                  struct{}{},
			"therapist":              struct{}{},
			"woman":                  struct{}{},
			"woman health worker":    struct{}{},
		},
	},
	":woman_health_worker_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\u2695\ufe0f",
		map[string]struct{}{
			"doctor":              struct{}{},
			"healthcare":          struct{}{},
			"medium skin tone":    struct{}{},
			"nurse":               struct{}{},
			"therapist":           struct{}{},
			"woman":               struct{}{},
			"woman health worker": struct{}{},
		},
	},
	":woman_health_worker_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\u2695\ufe0f",
		map[string]struct{}{
			"doctor":                struct{}{},
			"healthcare":            struct{}{},
			"medium-dark skin tone": struct{}{},
			"nurse":                 struct{}{},
			"therapist":             struct{}{},
			"woman":                 struct{}{},
			"woman health worker":   struct{}{},
		},
	},
	":woman_health_worker_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\u2695\ufe0f",
		map[string]struct{}{
			"dark skin tone":      struct{}{},
			"doctor":              struct{}{},
			"healthcare":          struct{}{},
			"nurse":               struct{}{},
			"therapist":           struct{}{},
			"woman":               struct{}{},
			"woman health worker": struct{}{},
		},
	},
	":man_student:": Code{
		"\U0001f468\u200d\U0001f393",
		map[string]struct{}{
			"graduate":    struct{}{},
			"man":         struct{}{},
			"man student": struct{}{},
			"student":     struct{}{},
		},
	},
	":man_student_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\U0001f393",
		map[string]struct{}{
			"graduate":        struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man student":     struct{}{},
			"student":         struct{}{},
		},
	},
	":man_student_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\U0001f393",
		map[string]struct{}{
			"graduate":               struct{}{},
			"man":                    struct{}{},
			"man student":            struct{}{},
			"medium-light skin tone": struct{}{},
			"student":                struct{}{},
		},
	},
	":man_student_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\U0001f393",
		map[string]struct{}{
			"graduate":         struct{}{},
			"man":              struct{}{},
			"man student":      struct{}{},
			"medium skin tone": struct{}{},
			"student":          struct{}{},
		},
	},
	":man_student_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\U0001f393",
		map[string]struct{}{
			"graduate":              struct{}{},
			"man":                   struct{}{},
			"man student":           struct{}{},
			"medium-dark skin tone": struct{}{},
			"student":               struct{}{},
		},
	},
	":man_student_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\U0001f393",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"graduate":       struct{}{},
			"man":            struct{}{},
			"man student":    struct{}{},
			"student":        struct{}{},
		},
	},
	":woman_student:": Code{
		"\U0001f469\u200d\U0001f393",
		map[string]struct{}{
			"graduate":      struct{}{},
			"student":       struct{}{},
			"woman":         struct{}{},
			"woman student": struct{}{},
		},
	},
	":woman_student_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\U0001f393",
		map[string]struct{}{
			"graduate":        struct{}{},
			"light skin tone": struct{}{},
			"student":         struct{}{},
			"woman":           struct{}{},
			"woman student":   struct{}{},
		},
	},
	":woman_student_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\U0001f393",
		map[string]struct{}{
			"graduate":               struct{}{},
			"medium-light skin tone": struct{}{},
			"student":                struct{}{},
			"woman":                  struct{}{},
			"woman student":          struct{}{},
		},
	},
	":woman_student_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\U0001f393",
		map[string]struct{}{
			"graduate":         struct{}{},
			"medium skin tone": struct{}{},
			"student":          struct{}{},
			"woman":            struct{}{},
			"woman student":    struct{}{},
		},
	},
	":woman_student_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\U0001f393",
		map[string]struct{}{
			"graduate":              struct{}{},
			"medium-dark skin tone": struct{}{},
			"student":               struct{}{},
			"woman":                 struct{}{},
			"woman student":         struct{}{},
		},
	},
	":woman_student_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\U0001f393",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"graduate":       struct{}{},
			"student":        struct{}{},
			"woman":          struct{}{},
			"woman student":  struct{}{},
		},
	},
	":man_teacher:": Code{
		"\U0001f468\u200d\U0001f3eb",
		map[string]struct{}{
			"instructor":  struct{}{},
			"man":         struct{}{},
			"man teacher": struct{}{},
			"professor":   struct{}{},
			"teacher":     struct{}{},
		},
	},
	":man_teacher_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\U0001f3eb",
		map[string]struct{}{
			"instructor":      struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man teacher":     struct{}{},
			"professor":       struct{}{},
			"teacher":         struct{}{},
		},
	},
	":man_teacher_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\U0001f3eb",
		map[string]struct{}{
			"instructor":             struct{}{},
			"man":                    struct{}{},
			"man teacher":            struct{}{},
			"medium-light skin tone": struct{}{},
			"professor":              struct{}{},
			"teacher":                struct{}{},
		},
	},
	":man_teacher_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\U0001f3eb",
		map[string]struct{}{
			"instructor":       struct{}{},
			"man":              struct{}{},
			"man teacher":      struct{}{},
			"medium skin tone": struct{}{},
			"professor":        struct{}{},
			"teacher":          struct{}{},
		},
	},
	":man_teacher_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\U0001f3eb",
		map[string]struct{}{
			"instructor":            struct{}{},
			"man":                   struct{}{},
			"man teacher":           struct{}{},
			"medium-dark skin tone": struct{}{},
			"professor":             struct{}{},
			"teacher":               struct{}{},
		},
	},
	":man_teacher_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\U0001f3eb",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"instructor":     struct{}{},
			"man":            struct{}{},
			"man teacher":    struct{}{},
			"professor":      struct{}{},
			"teacher":        struct{}{},
		},
	},
	":woman_teacher:": Code{
		"\U0001f469\u200d\U0001f3eb",
		map[string]struct{}{
			"instructor":    struct{}{},
			"professor":     struct{}{},
			"teacher":       struct{}{},
			"woman":         struct{}{},
			"woman teacher": struct{}{},
		},
	},
	":woman_teacher_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\U0001f3eb",
		map[string]struct{}{
			"instructor":      struct{}{},
			"light skin tone": struct{}{},
			"professor":       struct{}{},
			"teacher":         struct{}{},
			"woman":           struct{}{},
			"woman teacher":   struct{}{},
		},
	},
	":woman_teacher_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\U0001f3eb",
		map[string]struct{}{
			"instructor":             struct{}{},
			"medium-light skin tone": struct{}{},
			"professor":              struct{}{},
			"teacher":                struct{}{},
			"woman":                  struct{}{},
			"woman teacher":          struct{}{},
		},
	},
	":woman_teacher_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\U0001f3eb",
		map[string]struct{}{
			"instructor":       struct{}{},
			"medium skin tone": struct{}{},
			"professor":        struct{}{},
			"teacher":          struct{}{},
			"woman":            struct{}{},
			"woman teacher":    struct{}{},
		},
	},
	":woman_teacher_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\U0001f3eb",
		map[string]struct{}{
			"instructor":            struct{}{},
			"medium-dark skin tone": struct{}{},
			"professor":             struct{}{},
			"teacher":               struct{}{},
			"woman":                 struct{}{},
			"woman teacher":         struct{}{},
		},
	},
	":woman_teacher_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\U0001f3eb",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"instructor":     struct{}{},
			"professor":      struct{}{},
			"teacher":        struct{}{},
			"woman":          struct{}{},
			"woman teacher":  struct{}{},
		},
	},
	":man_judge:": Code{
		"\U0001f468\u200d\u2696\ufe0f",
		map[string]struct{}{
			"justice":   struct{}{},
			"man":       struct{}{},
			"man judge": struct{}{},
			"scales":    struct{}{},
		},
	},
	":man_judge_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\u2696\ufe0f",
		map[string]struct{}{
			"justice":         struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man judge":       struct{}{},
			"scales":          struct{}{},
		},
	},
	":man_judge_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\u2696\ufe0f",
		map[string]struct{}{
			"justice":                struct{}{},
			"man":                    struct{}{},
			"man judge":              struct{}{},
			"medium-light skin tone": struct{}{},
			"scales":                 struct{}{},
		},
	},
	":man_judge_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\u2696\ufe0f",
		map[string]struct{}{
			"justice":          struct{}{},
			"man":              struct{}{},
			"man judge":        struct{}{},
			"medium skin tone": struct{}{},
			"scales":           struct{}{},
		},
	},
	":man_judge_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\u2696\ufe0f",
		map[string]struct{}{
			"justice":               struct{}{},
			"man":                   struct{}{},
			"man judge":             struct{}{},
			"medium-dark skin tone": struct{}{},
			"scales":                struct{}{},
		},
	},
	":man_judge_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\u2696\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"justice":        struct{}{},
			"man":            struct{}{},
			"man judge":      struct{}{},
			"scales":         struct{}{},
		},
	},
	":woman_judge:": Code{
		"\U0001f469\u200d\u2696\ufe0f",
		map[string]struct{}{
			"judge":       struct{}{},
			"scales":      struct{}{},
			"woman":       struct{}{},
			"woman judge": struct{}{},
		},
	},
	":woman_judge_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\u2696\ufe0f",
		map[string]struct{}{
			"judge":           struct{}{},
			"light skin tone": struct{}{},
			"scales":          struct{}{},
			"woman":           struct{}{},
			"woman judge":     struct{}{},
		},
	},
	":woman_judge_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\u2696\ufe0f",
		map[string]struct{}{
			"judge":                  struct{}{},
			"medium-light skin tone": struct{}{},
			"scales":                 struct{}{},
			"woman":                  struct{}{},
			"woman judge":            struct{}{},
		},
	},
	":woman_judge_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\u2696\ufe0f",
		map[string]struct{}{
			"judge":            struct{}{},
			"medium skin tone": struct{}{},
			"scales":           struct{}{},
			"woman":            struct{}{},
			"woman judge":      struct{}{},
		},
	},
	":woman_judge_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\u2696\ufe0f",
		map[string]struct{}{
			"judge":                 struct{}{},
			"medium-dark skin tone": struct{}{},
			"scales":                struct{}{},
			"woman":                 struct{}{},
			"woman judge":           struct{}{},
		},
	},
	":woman_judge_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\u2696\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"judge":          struct{}{},
			"scales":         struct{}{},
			"woman":          struct{}{},
			"woman judge":    struct{}{},
		},
	},
	":man_farmer:": Code{
		"\U0001f468\u200d\U0001f33e",
		map[string]struct{}{
			"farmer":     struct{}{},
			"gardener":   struct{}{},
			"man":        struct{}{},
			"man farmer": struct{}{},
			"rancher":    struct{}{},
		},
	},
	":man_farmer_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\U0001f33e",
		map[string]struct{}{
			"farmer":          struct{}{},
			"gardener":        struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man farmer":      struct{}{},
			"rancher":         struct{}{},
		},
	},
	":man_farmer_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\U0001f33e",
		map[string]struct{}{
			"farmer":                 struct{}{},
			"gardener":               struct{}{},
			"man":                    struct{}{},
			"man farmer":             struct{}{},
			"medium-light skin tone": struct{}{},
			"rancher":                struct{}{},
		},
	},
	":man_farmer_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\U0001f33e",
		map[string]struct{}{
			"farmer":           struct{}{},
			"gardener":         struct{}{},
			"man":              struct{}{},
			"man farmer":       struct{}{},
			"medium skin tone": struct{}{},
			"rancher":          struct{}{},
		},
	},
	":man_farmer_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\U0001f33e",
		map[string]struct{}{
			"farmer":                struct{}{},
			"gardener":              struct{}{},
			"man":                   struct{}{},
			"man farmer":            struct{}{},
			"medium-dark skin tone": struct{}{},
			"rancher":               struct{}{},
		},
	},
	":man_farmer_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\U0001f33e",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"farmer":         struct{}{},
			"gardener":       struct{}{},
			"man":            struct{}{},
			"man farmer":     struct{}{},
			"rancher":        struct{}{},
		},
	},
	":woman_farmer:": Code{
		"\U0001f469\u200d\U0001f33e",
		map[string]struct{}{
			"farmer":       struct{}{},
			"gardener":     struct{}{},
			"rancher":      struct{}{},
			"woman":        struct{}{},
			"woman farmer": struct{}{},
		},
	},
	":woman_farmer_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\U0001f33e",
		map[string]struct{}{
			"farmer":          struct{}{},
			"gardener":        struct{}{},
			"light skin tone": struct{}{},
			"rancher":         struct{}{},
			"woman":           struct{}{},
			"woman farmer":    struct{}{},
		},
	},
	":woman_farmer_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\U0001f33e",
		map[string]struct{}{
			"farmer":                 struct{}{},
			"gardener":               struct{}{},
			"medium-light skin tone": struct{}{},
			"rancher":                struct{}{},
			"woman":                  struct{}{},
			"woman farmer":           struct{}{},
		},
	},
	":woman_farmer_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\U0001f33e",
		map[string]struct{}{
			"farmer":           struct{}{},
			"gardener":         struct{}{},
			"medium skin tone": struct{}{},
			"rancher":          struct{}{},
			"woman":            struct{}{},
			"woman farmer":     struct{}{},
		},
	},
	":woman_farmer_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\U0001f33e",
		map[string]struct{}{
			"farmer":                struct{}{},
			"gardener":              struct{}{},
			"medium-dark skin tone": struct{}{},
			"rancher":               struct{}{},
			"woman":                 struct{}{},
			"woman farmer":          struct{}{},
		},
	},
	":woman_farmer_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\U0001f33e",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"farmer":         struct{}{},
			"gardener":       struct{}{},
			"rancher":        struct{}{},
			"woman":          struct{}{},
			"woman farmer":   struct{}{},
		},
	},
	":man_cook:": Code{
		"\U0001f468\u200d\U0001f373",
		map[string]struct{}{
			"chef":     struct{}{},
			"cook":     struct{}{},
			"man":      struct{}{},
			"man cook": struct{}{},
		},
	},
	":man_cook_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\U0001f373",
		map[string]struct{}{
			"chef":            struct{}{},
			"cook":            struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man cook":        struct{}{},
		},
	},
	":man_cook_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\U0001f373",
		map[string]struct{}{
			"chef":                   struct{}{},
			"cook":                   struct{}{},
			"man":                    struct{}{},
			"man cook":               struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_cook_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\U0001f373",
		map[string]struct{}{
			"chef":             struct{}{},
			"cook":             struct{}{},
			"man":              struct{}{},
			"man cook":         struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":man_cook_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\U0001f373",
		map[string]struct{}{
			"chef":                  struct{}{},
			"cook":                  struct{}{},
			"man":                   struct{}{},
			"man cook":              struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_cook_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\U0001f373",
		map[string]struct{}{
			"chef":           struct{}{},
			"cook":           struct{}{},
			"dark skin tone": struct{}{},
			"man":            struct{}{},
			"man cook":       struct{}{},
		},
	},
	":woman_cook:": Code{
		"\U0001f469\u200d\U0001f373",
		map[string]struct{}{
			"chef":       struct{}{},
			"cook":       struct{}{},
			"woman":      struct{}{},
			"woman cook": struct{}{},
		},
	},
	":woman_cook_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\U0001f373",
		map[string]struct{}{
			"chef":            struct{}{},
			"cook":            struct{}{},
			"light skin tone": struct{}{},
			"woman":           struct{}{},
			"woman cook":      struct{}{},
		},
	},
	":woman_cook_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\U0001f373",
		map[string]struct{}{
			"chef":                   struct{}{},
			"cook":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
			"woman cook":             struct{}{},
		},
	},
	":woman_cook_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\U0001f373",
		map[string]struct{}{
			"chef":             struct{}{},
			"cook":             struct{}{},
			"medium skin tone": struct{}{},
			"woman":            struct{}{},
			"woman cook":       struct{}{},
		},
	},
	":woman_cook_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\U0001f373",
		map[string]struct{}{
			"chef":                  struct{}{},
			"cook":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
			"woman cook":            struct{}{},
		},
	},
	":woman_cook_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\U0001f373",
		map[string]struct{}{
			"chef":           struct{}{},
			"cook":           struct{}{},
			"dark skin tone": struct{}{},
			"woman":          struct{}{},
			"woman cook":     struct{}{},
		},
	},
	":man_mechanic:": Code{
		"\U0001f468\u200d\U0001f527",
		map[string]struct{}{
			"electrician":  struct{}{},
			"man":          struct{}{},
			"man mechanic": struct{}{},
			"mechanic":     struct{}{},
			"plumber":      struct{}{},
			"tradesperson": struct{}{},
		},
	},
	":man_mechanic_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\U0001f527",
		map[string]struct{}{
			"electrician":     struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man mechanic":    struct{}{},
			"mechanic":        struct{}{},
			"plumber":         struct{}{},
			"tradesperson":    struct{}{},
		},
	},
	":man_mechanic_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\U0001f527",
		map[string]struct{}{
			"electrician":            struct{}{},
			"man":                    struct{}{},
			"man mechanic":           struct{}{},
			"mechanic":               struct{}{},
			"medium-light skin tone": struct{}{},
			"plumber":                struct{}{},
			"tradesperson":           struct{}{},
		},
	},
	":man_mechanic_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\U0001f527",
		map[string]struct{}{
			"electrician":      struct{}{},
			"man":              struct{}{},
			"man mechanic":     struct{}{},
			"mechanic":         struct{}{},
			"medium skin tone": struct{}{},
			"plumber":          struct{}{},
			"tradesperson":     struct{}{},
		},
	},
	":man_mechanic_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\U0001f527",
		map[string]struct{}{
			"electrician":           struct{}{},
			"man":                   struct{}{},
			"man mechanic":          struct{}{},
			"mechanic":              struct{}{},
			"medium-dark skin tone": struct{}{},
			"plumber":               struct{}{},
			"tradesperson":          struct{}{},
		},
	},
	":man_mechanic_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\U0001f527",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"electrician":    struct{}{},
			"man":            struct{}{},
			"man mechanic":   struct{}{},
			"mechanic":       struct{}{},
			"plumber":        struct{}{},
			"tradesperson":   struct{}{},
		},
	},
	":woman_mechanic:": Code{
		"\U0001f469\u200d\U0001f527",
		map[string]struct{}{
			"electrician":    struct{}{},
			"mechanic":       struct{}{},
			"plumber":        struct{}{},
			"tradesperson":   struct{}{},
			"woman":          struct{}{},
			"woman mechanic": struct{}{},
		},
	},
	":woman_mechanic_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\U0001f527",
		map[string]struct{}{
			"electrician":     struct{}{},
			"light skin tone": struct{}{},
			"mechanic":        struct{}{},
			"plumber":         struct{}{},
			"tradesperson":    struct{}{},
			"woman":           struct{}{},
			"woman mechanic":  struct{}{},
		},
	},
	":woman_mechanic_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\U0001f527",
		map[string]struct{}{
			"electrician":            struct{}{},
			"mechanic":               struct{}{},
			"medium-light skin tone": struct{}{},
			"plumber":                struct{}{},
			"tradesperson":           struct{}{},
			"woman":                  struct{}{},
			"woman mechanic":         struct{}{},
		},
	},
	":woman_mechanic_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\U0001f527",
		map[string]struct{}{
			"electrician":      struct{}{},
			"mechanic":         struct{}{},
			"medium skin tone": struct{}{},
			"plumber":          struct{}{},
			"tradesperson":     struct{}{},
			"woman":            struct{}{},
			"woman mechanic":   struct{}{},
		},
	},
	":woman_mechanic_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\U0001f527",
		map[string]struct{}{
			"electrician":           struct{}{},
			"mechanic":              struct{}{},
			"medium-dark skin tone": struct{}{},
			"plumber":               struct{}{},
			"tradesperson":          struct{}{},
			"woman":                 struct{}{},
			"woman mechanic":        struct{}{},
		},
	},
	":woman_mechanic_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\U0001f527",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"electrician":    struct{}{},
			"mechanic":       struct{}{},
			"plumber":        struct{}{},
			"tradesperson":   struct{}{},
			"woman":          struct{}{},
			"woman mechanic": struct{}{},
		},
	},
	":man_factory_worker:": Code{
		"\U0001f468\u200d\U0001f3ed",
		map[string]struct{}{
			"assembly":           struct{}{},
			"factory":            struct{}{},
			"industrial":         struct{}{},
			"man":                struct{}{},
			"man factory worker": struct{}{},
			"worker":             struct{}{},
		},
	},
	":man_factory_worker_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\U0001f3ed",
		map[string]struct{}{
			"assembly":           struct{}{},
			"factory":            struct{}{},
			"industrial":         struct{}{},
			"light skin tone":    struct{}{},
			"man":                struct{}{},
			"man factory worker": struct{}{},
			"worker":             struct{}{},
		},
	},
	":man_factory_worker_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\U0001f3ed",
		map[string]struct{}{
			"assembly":               struct{}{},
			"factory":                struct{}{},
			"industrial":             struct{}{},
			"man":                    struct{}{},
			"man factory worker":     struct{}{},
			"medium-light skin tone": struct{}{},
			"worker":                 struct{}{},
		},
	},
	":man_factory_worker_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\U0001f3ed",
		map[string]struct{}{
			"assembly":           struct{}{},
			"factory":            struct{}{},
			"industrial":         struct{}{},
			"man":                struct{}{},
			"man factory worker": struct{}{},
			"medium skin tone":   struct{}{},
			"worker":             struct{}{},
		},
	},
	":man_factory_worker_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\U0001f3ed",
		map[string]struct{}{
			"assembly":              struct{}{},
			"factory":               struct{}{},
			"industrial":            struct{}{},
			"man":                   struct{}{},
			"man factory worker":    struct{}{},
			"medium-dark skin tone": struct{}{},
			"worker":                struct{}{},
		},
	},
	":man_factory_worker_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\U0001f3ed",
		map[string]struct{}{
			"assembly":           struct{}{},
			"dark skin tone":     struct{}{},
			"factory":            struct{}{},
			"industrial":         struct{}{},
			"man":                struct{}{},
			"man factory worker": struct{}{},
			"worker":             struct{}{},
		},
	},
	":woman_factory_worker:": Code{
		"\U0001f469\u200d\U0001f3ed",
		map[string]struct{}{
			"assembly":             struct{}{},
			"factory":              struct{}{},
			"industrial":           struct{}{},
			"woman":                struct{}{},
			"woman factory worker": struct{}{},
			"worker":               struct{}{},
		},
	},
	":woman_factory_worker_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\U0001f3ed",
		map[string]struct{}{
			"assembly":             struct{}{},
			"factory":              struct{}{},
			"industrial":           struct{}{},
			"light skin tone":      struct{}{},
			"woman":                struct{}{},
			"woman factory worker": struct{}{},
			"worker":               struct{}{},
		},
	},
	":woman_factory_worker_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\U0001f3ed",
		map[string]struct{}{
			"assembly":               struct{}{},
			"factory":                struct{}{},
			"industrial":             struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
			"woman factory worker":   struct{}{},
			"worker":                 struct{}{},
		},
	},
	":woman_factory_worker_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\U0001f3ed",
		map[string]struct{}{
			"assembly":             struct{}{},
			"factory":              struct{}{},
			"industrial":           struct{}{},
			"medium skin tone":     struct{}{},
			"woman":                struct{}{},
			"woman factory worker": struct{}{},
			"worker":               struct{}{},
		},
	},
	":woman_factory_worker_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\U0001f3ed",
		map[string]struct{}{
			"assembly":              struct{}{},
			"factory":               struct{}{},
			"industrial":            struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
			"woman factory worker":  struct{}{},
			"worker":                struct{}{},
		},
	},
	":woman_factory_worker_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\U0001f3ed",
		map[string]struct{}{
			"assembly":             struct{}{},
			"dark skin tone":       struct{}{},
			"factory":              struct{}{},
			"industrial":           struct{}{},
			"woman":                struct{}{},
			"woman factory worker": struct{}{},
			"worker":               struct{}{},
		},
	},
	":man_office_worker:": Code{
		"\U0001f468\u200d\U0001f4bc",
		map[string]struct{}{
			"architect":         struct{}{},
			"business":          struct{}{},
			"man":               struct{}{},
			"man office worker": struct{}{},
			"manager":           struct{}{},
			"office":            struct{}{},
			"white-collar":      struct{}{},
		},
	},
	":man_office_worker_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\U0001f4bc",
		map[string]struct{}{
			"architect":         struct{}{},
			"business":          struct{}{},
			"light skin tone":   struct{}{},
			"man":               struct{}{},
			"man office worker": struct{}{},
			"manager":           struct{}{},
			"office":            struct{}{},
			"white-collar":      struct{}{},
		},
	},
	":man_office_worker_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\U0001f4bc",
		map[string]struct{}{
			"architect":              struct{}{},
			"business":               struct{}{},
			"man":                    struct{}{},
			"man office worker":      struct{}{},
			"manager":                struct{}{},
			"medium-light skin tone": struct{}{},
			"office":                 struct{}{},
			"white-collar":           struct{}{},
		},
	},
	":man_office_worker_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\U0001f4bc",
		map[string]struct{}{
			"architect":         struct{}{},
			"business":          struct{}{},
			"man":               struct{}{},
			"man office worker": struct{}{},
			"manager":           struct{}{},
			"medium skin tone":  struct{}{},
			"office":            struct{}{},
			"white-collar":      struct{}{},
		},
	},
	":man_office_worker_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\U0001f4bc",
		map[string]struct{}{
			"architect":             struct{}{},
			"business":              struct{}{},
			"man":                   struct{}{},
			"man office worker":     struct{}{},
			"manager":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"office":                struct{}{},
			"white-collar":          struct{}{},
		},
	},
	":man_office_worker_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\U0001f4bc",
		map[string]struct{}{
			"architect":         struct{}{},
			"business":          struct{}{},
			"dark skin tone":    struct{}{},
			"man":               struct{}{},
			"man office worker": struct{}{},
			"manager":           struct{}{},
			"office":            struct{}{},
			"white-collar":      struct{}{},
		},
	},
	":woman_office_worker:": Code{
		"\U0001f469\u200d\U0001f4bc",
		map[string]struct{}{
			"architect":           struct{}{},
			"business":            struct{}{},
			"manager":             struct{}{},
			"office":              struct{}{},
			"white-collar":        struct{}{},
			"woman":               struct{}{},
			"woman office worker": struct{}{},
		},
	},
	":woman_office_worker_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\U0001f4bc",
		map[string]struct{}{
			"architect":           struct{}{},
			"business":            struct{}{},
			"light skin tone":     struct{}{},
			"manager":             struct{}{},
			"office":              struct{}{},
			"white-collar":        struct{}{},
			"woman":               struct{}{},
			"woman office worker": struct{}{},
		},
	},
	":woman_office_worker_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\U0001f4bc",
		map[string]struct{}{
			"architect":              struct{}{},
			"business":               struct{}{},
			"manager":                struct{}{},
			"medium-light skin tone": struct{}{},
			"office":                 struct{}{},
			"white-collar":           struct{}{},
			"woman":                  struct{}{},
			"woman office worker":    struct{}{},
		},
	},
	":woman_office_worker_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\U0001f4bc",
		map[string]struct{}{
			"architect":           struct{}{},
			"business":            struct{}{},
			"manager":             struct{}{},
			"medium skin tone":    struct{}{},
			"office":              struct{}{},
			"white-collar":        struct{}{},
			"woman":               struct{}{},
			"woman office worker": struct{}{},
		},
	},
	":woman_office_worker_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\U0001f4bc",
		map[string]struct{}{
			"architect":             struct{}{},
			"business":              struct{}{},
			"manager":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"office":                struct{}{},
			"white-collar":          struct{}{},
			"woman":                 struct{}{},
			"woman office worker":   struct{}{},
		},
	},
	":woman_office_worker_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\U0001f4bc",
		map[string]struct{}{
			"architect":           struct{}{},
			"business":            struct{}{},
			"dark skin tone":      struct{}{},
			"manager":             struct{}{},
			"office":              struct{}{},
			"white-collar":        struct{}{},
			"woman":               struct{}{},
			"woman office worker": struct{}{},
		},
	},
	":man_scientist:": Code{
		"\U0001f468\u200d\U0001f52c",
		map[string]struct{}{
			"biologist":     struct{}{},
			"chemist":       struct{}{},
			"engineer":      struct{}{},
			"man":           struct{}{},
			"man scientist": struct{}{},
			"mathematician": struct{}{},
			"physicist":     struct{}{},
			"scientist":     struct{}{},
		},
	},
	":man_scientist_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\U0001f52c",
		map[string]struct{}{
			"biologist":       struct{}{},
			"chemist":         struct{}{},
			"engineer":        struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man scientist":   struct{}{},
			"mathematician":   struct{}{},
			"physicist":       struct{}{},
			"scientist":       struct{}{},
		},
	},
	":man_scientist_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\U0001f52c",
		map[string]struct{}{
			"biologist":              struct{}{},
			"chemist":                struct{}{},
			"engineer":               struct{}{},
			"man":                    struct{}{},
			"man scientist":          struct{}{},
			"mathematician":          struct{}{},
			"medium-light skin tone": struct{}{},
			"physicist":              struct{}{},
			"scientist":              struct{}{},
		},
	},
	":man_scientist_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\U0001f52c",
		map[string]struct{}{
			"biologist":        struct{}{},
			"chemist":          struct{}{},
			"engineer":         struct{}{},
			"man":              struct{}{},
			"man scientist":    struct{}{},
			"mathematician":    struct{}{},
			"medium skin tone": struct{}{},
			"physicist":        struct{}{},
			"scientist":        struct{}{},
		},
	},
	":man_scientist_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\U0001f52c",
		map[string]struct{}{
			"biologist":             struct{}{},
			"chemist":               struct{}{},
			"engineer":              struct{}{},
			"man":                   struct{}{},
			"man scientist":         struct{}{},
			"mathematician":         struct{}{},
			"medium-dark skin tone": struct{}{},
			"physicist":             struct{}{},
			"scientist":             struct{}{},
		},
	},
	":man_scientist_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\U0001f52c",
		map[string]struct{}{
			"biologist":      struct{}{},
			"chemist":        struct{}{},
			"dark skin tone": struct{}{},
			"engineer":       struct{}{},
			"man":            struct{}{},
			"man scientist":  struct{}{},
			"mathematician":  struct{}{},
			"physicist":      struct{}{},
			"scientist":      struct{}{},
		},
	},
	":woman_scientist:": Code{
		"\U0001f469\u200d\U0001f52c",
		map[string]struct{}{
			"biologist":       struct{}{},
			"chemist":         struct{}{},
			"engineer":        struct{}{},
			"mathematician":   struct{}{},
			"physicist":       struct{}{},
			"scientist":       struct{}{},
			"woman":           struct{}{},
			"woman scientist": struct{}{},
		},
	},
	":woman_scientist_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\U0001f52c",
		map[string]struct{}{
			"biologist":       struct{}{},
			"chemist":         struct{}{},
			"engineer":        struct{}{},
			"light skin tone": struct{}{},
			"mathematician":   struct{}{},
			"physicist":       struct{}{},
			"scientist":       struct{}{},
			"woman":           struct{}{},
			"woman scientist": struct{}{},
		},
	},
	":woman_scientist_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\U0001f52c",
		map[string]struct{}{
			"biologist":              struct{}{},
			"chemist":                struct{}{},
			"engineer":               struct{}{},
			"mathematician":          struct{}{},
			"medium-light skin tone": struct{}{},
			"physicist":              struct{}{},
			"scientist":              struct{}{},
			"woman":                  struct{}{},
			"woman scientist":        struct{}{},
		},
	},
	":woman_scientist_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\U0001f52c",
		map[string]struct{}{
			"biologist":        struct{}{},
			"chemist":          struct{}{},
			"engineer":         struct{}{},
			"mathematician":    struct{}{},
			"medium skin tone": struct{}{},
			"physicist":        struct{}{},
			"scientist":        struct{}{},
			"woman":            struct{}{},
			"woman scientist":  struct{}{},
		},
	},
	":woman_scientist_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\U0001f52c",
		map[string]struct{}{
			"biologist":             struct{}{},
			"chemist":               struct{}{},
			"engineer":              struct{}{},
			"mathematician":         struct{}{},
			"medium-dark skin tone": struct{}{},
			"physicist":             struct{}{},
			"scientist":             struct{}{},
			"woman":                 struct{}{},
			"woman scientist":       struct{}{},
		},
	},
	":woman_scientist_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\U0001f52c",
		map[string]struct{}{
			"biologist":       struct{}{},
			"chemist":         struct{}{},
			"dark skin tone":  struct{}{},
			"engineer":        struct{}{},
			"mathematician":   struct{}{},
			"physicist":       struct{}{},
			"scientist":       struct{}{},
			"woman":           struct{}{},
			"woman scientist": struct{}{},
		},
	},
	":man_technologist:": Code{
		"\U0001f468\u200d\U0001f4bb",
		map[string]struct{}{
			"coder":            struct{}{},
			"developer":        struct{}{},
			"inventor":         struct{}{},
			"man":              struct{}{},
			"man technologist": struct{}{},
			"software":         struct{}{},
			"technologist":     struct{}{},
		},
	},
	":man_technologist_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\U0001f4bb",
		map[string]struct{}{
			"coder":            struct{}{},
			"developer":        struct{}{},
			"inventor":         struct{}{},
			"light skin tone":  struct{}{},
			"man":              struct{}{},
			"man technologist": struct{}{},
			"software":         struct{}{},
			"technologist":     struct{}{},
		},
	},
	":man_technologist_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\U0001f4bb",
		map[string]struct{}{
			"coder":                  struct{}{},
			"developer":              struct{}{},
			"inventor":               struct{}{},
			"man":                    struct{}{},
			"man technologist":       struct{}{},
			"medium-light skin tone": struct{}{},
			"software":               struct{}{},
			"technologist":           struct{}{},
		},
	},
	":man_technologist_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\U0001f4bb",
		map[string]struct{}{
			"coder":            struct{}{},
			"developer":        struct{}{},
			"inventor":         struct{}{},
			"man":              struct{}{},
			"man technologist": struct{}{},
			"medium skin tone": struct{}{},
			"software":         struct{}{},
			"technologist":     struct{}{},
		},
	},
	":man_technologist_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\U0001f4bb",
		map[string]struct{}{
			"coder":                 struct{}{},
			"developer":             struct{}{},
			"inventor":              struct{}{},
			"man":                   struct{}{},
			"man technologist":      struct{}{},
			"medium-dark skin tone": struct{}{},
			"software":              struct{}{},
			"technologist":          struct{}{},
		},
	},
	":man_technologist_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\U0001f4bb",
		map[string]struct{}{
			"coder":            struct{}{},
			"dark skin tone":   struct{}{},
			"developer":        struct{}{},
			"inventor":         struct{}{},
			"man":              struct{}{},
			"man technologist": struct{}{},
			"software":         struct{}{},
			"technologist":     struct{}{},
		},
	},
	":woman_technologist:": Code{
		"\U0001f469\u200d\U0001f4bb",
		map[string]struct{}{
			"coder":              struct{}{},
			"developer":          struct{}{},
			"inventor":           struct{}{},
			"software":           struct{}{},
			"technologist":       struct{}{},
			"woman":              struct{}{},
			"woman technologist": struct{}{},
		},
	},
	":woman_technologist_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\U0001f4bb",
		map[string]struct{}{
			"coder":              struct{}{},
			"developer":          struct{}{},
			"inventor":           struct{}{},
			"light skin tone":    struct{}{},
			"software":           struct{}{},
			"technologist":       struct{}{},
			"woman":              struct{}{},
			"woman technologist": struct{}{},
		},
	},
	":woman_technologist_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\U0001f4bb",
		map[string]struct{}{
			"coder":                  struct{}{},
			"developer":              struct{}{},
			"inventor":               struct{}{},
			"medium-light skin tone": struct{}{},
			"software":               struct{}{},
			"technologist":           struct{}{},
			"woman":                  struct{}{},
			"woman technologist":     struct{}{},
		},
	},
	":woman_technologist_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\U0001f4bb",
		map[string]struct{}{
			"coder":              struct{}{},
			"developer":          struct{}{},
			"inventor":           struct{}{},
			"medium skin tone":   struct{}{},
			"software":           struct{}{},
			"technologist":       struct{}{},
			"woman":              struct{}{},
			"woman technologist": struct{}{},
		},
	},
	":woman_technologist_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\U0001f4bb",
		map[string]struct{}{
			"coder":                 struct{}{},
			"developer":             struct{}{},
			"inventor":              struct{}{},
			"medium-dark skin tone": struct{}{},
			"software":              struct{}{},
			"technologist":          struct{}{},
			"woman":                 struct{}{},
			"woman technologist":    struct{}{},
		},
	},
	":woman_technologist_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\U0001f4bb",
		map[string]struct{}{
			"coder":              struct{}{},
			"dark skin tone":     struct{}{},
			"developer":          struct{}{},
			"inventor":           struct{}{},
			"software":           struct{}{},
			"technologist":       struct{}{},
			"woman":              struct{}{},
			"woman technologist": struct{}{},
		},
	},
	":man_singer:": Code{
		"\U0001f468\u200d\U0001f3a4",
		map[string]struct{}{
			"actor":       struct{}{},
			"entertainer": struct{}{},
			"man":         struct{}{},
			"man singer":  struct{}{},
			"rock":        struct{}{},
			"singer":      struct{}{},
			"star":        struct{}{},
		},
	},
	":man_singer_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\U0001f3a4",
		map[string]struct{}{
			"actor":           struct{}{},
			"entertainer":     struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man singer":      struct{}{},
			"rock":            struct{}{},
			"singer":          struct{}{},
			"star":            struct{}{},
		},
	},
	":man_singer_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\U0001f3a4",
		map[string]struct{}{
			"actor":                  struct{}{},
			"entertainer":            struct{}{},
			"man":                    struct{}{},
			"man singer":             struct{}{},
			"medium-light skin tone": struct{}{},
			"rock":                   struct{}{},
			"singer":                 struct{}{},
			"star":                   struct{}{},
		},
	},
	":man_singer_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\U0001f3a4",
		map[string]struct{}{
			"actor":            struct{}{},
			"entertainer":      struct{}{},
			"man":              struct{}{},
			"man singer":       struct{}{},
			"medium skin tone": struct{}{},
			"rock":             struct{}{},
			"singer":           struct{}{},
			"star":             struct{}{},
		},
	},
	":man_singer_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\U0001f3a4",
		map[string]struct{}{
			"actor":                 struct{}{},
			"entertainer":           struct{}{},
			"man":                   struct{}{},
			"man singer":            struct{}{},
			"medium-dark skin tone": struct{}{},
			"rock":                  struct{}{},
			"singer":                struct{}{},
			"star":                  struct{}{},
		},
	},
	":man_singer_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\U0001f3a4",
		map[string]struct{}{
			"actor":          struct{}{},
			"dark skin tone": struct{}{},
			"entertainer":    struct{}{},
			"man":            struct{}{},
			"man singer":     struct{}{},
			"rock":           struct{}{},
			"singer":         struct{}{},
			"star":           struct{}{},
		},
	},
	":woman_singer:": Code{
		"\U0001f469\u200d\U0001f3a4",
		map[string]struct{}{
			"actor":        struct{}{},
			"entertainer":  struct{}{},
			"rock":         struct{}{},
			"singer":       struct{}{},
			"star":         struct{}{},
			"woman":        struct{}{},
			"woman singer": struct{}{},
		},
	},
	":woman_singer_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\U0001f3a4",
		map[string]struct{}{
			"actor":           struct{}{},
			"entertainer":     struct{}{},
			"light skin tone": struct{}{},
			"rock":            struct{}{},
			"singer":          struct{}{},
			"star":            struct{}{},
			"woman":           struct{}{},
			"woman singer":    struct{}{},
		},
	},
	":woman_singer_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\U0001f3a4",
		map[string]struct{}{
			"actor":                  struct{}{},
			"entertainer":            struct{}{},
			"medium-light skin tone": struct{}{},
			"rock":                   struct{}{},
			"singer":                 struct{}{},
			"star":                   struct{}{},
			"woman":                  struct{}{},
			"woman singer":           struct{}{},
		},
	},
	":woman_singer_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\U0001f3a4",
		map[string]struct{}{
			"actor":            struct{}{},
			"entertainer":      struct{}{},
			"medium skin tone": struct{}{},
			"rock":             struct{}{},
			"singer":           struct{}{},
			"star":             struct{}{},
			"woman":            struct{}{},
			"woman singer":     struct{}{},
		},
	},
	":woman_singer_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\U0001f3a4",
		map[string]struct{}{
			"actor":                 struct{}{},
			"entertainer":           struct{}{},
			"medium-dark skin tone": struct{}{},
			"rock":                  struct{}{},
			"singer":                struct{}{},
			"star":                  struct{}{},
			"woman":                 struct{}{},
			"woman singer":          struct{}{},
		},
	},
	":woman_singer_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\U0001f3a4",
		map[string]struct{}{
			"actor":          struct{}{},
			"dark skin tone": struct{}{},
			"entertainer":    struct{}{},
			"rock":           struct{}{},
			"singer":         struct{}{},
			"star":           struct{}{},
			"woman":          struct{}{},
			"woman singer":   struct{}{},
		},
	},
	":man_artist:": Code{
		"\U0001f468\u200d\U0001f3a8",
		map[string]struct{}{
			"artist":     struct{}{},
			"man":        struct{}{},
			"man artist": struct{}{},
			"palette":    struct{}{},
		},
	},
	":man_artist_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\U0001f3a8",
		map[string]struct{}{
			"artist":          struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man artist":      struct{}{},
			"palette":         struct{}{},
		},
	},
	":man_artist_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\U0001f3a8",
		map[string]struct{}{
			"artist":                 struct{}{},
			"man":                    struct{}{},
			"man artist":             struct{}{},
			"medium-light skin tone": struct{}{},
			"palette":                struct{}{},
		},
	},
	":man_artist_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\U0001f3a8",
		map[string]struct{}{
			"artist":           struct{}{},
			"man":              struct{}{},
			"man artist":       struct{}{},
			"medium skin tone": struct{}{},
			"palette":          struct{}{},
		},
	},
	":man_artist_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\U0001f3a8",
		map[string]struct{}{
			"artist":                struct{}{},
			"man":                   struct{}{},
			"man artist":            struct{}{},
			"medium-dark skin tone": struct{}{},
			"palette":               struct{}{},
		},
	},
	":man_artist_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\U0001f3a8",
		map[string]struct{}{
			"artist":         struct{}{},
			"dark skin tone": struct{}{},
			"man":            struct{}{},
			"man artist":     struct{}{},
			"palette":        struct{}{},
		},
	},
	":woman_artist:": Code{
		"\U0001f469\u200d\U0001f3a8",
		map[string]struct{}{
			"artist":       struct{}{},
			"palette":      struct{}{},
			"woman":        struct{}{},
			"woman artist": struct{}{},
		},
	},
	":woman_artist_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\U0001f3a8",
		map[string]struct{}{
			"artist":          struct{}{},
			"light skin tone": struct{}{},
			"palette":         struct{}{},
			"woman":           struct{}{},
			"woman artist":    struct{}{},
		},
	},
	":woman_artist_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\U0001f3a8",
		map[string]struct{}{
			"artist":                 struct{}{},
			"medium-light skin tone": struct{}{},
			"palette":                struct{}{},
			"woman":                  struct{}{},
			"woman artist":           struct{}{},
		},
	},
	":woman_artist_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\U0001f3a8",
		map[string]struct{}{
			"artist":           struct{}{},
			"medium skin tone": struct{}{},
			"palette":          struct{}{},
			"woman":            struct{}{},
			"woman artist":     struct{}{},
		},
	},
	":woman_artist_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\U0001f3a8",
		map[string]struct{}{
			"artist":                struct{}{},
			"medium-dark skin tone": struct{}{},
			"palette":               struct{}{},
			"woman":                 struct{}{},
			"woman artist":          struct{}{},
		},
	},
	":woman_artist_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\U0001f3a8",
		map[string]struct{}{
			"artist":         struct{}{},
			"dark skin tone": struct{}{},
			"palette":        struct{}{},
			"woman":          struct{}{},
			"woman artist":   struct{}{},
		},
	},
	":man_pilot:": Code{
		"\U0001f468\u200d\u2708\ufe0f",
		map[string]struct{}{
			"man":       struct{}{},
			"man pilot": struct{}{},
			"pilot":     struct{}{},
			"plane":     struct{}{},
		},
	},
	":man_pilot_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\u2708\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man pilot":       struct{}{},
			"pilot":           struct{}{},
			"plane":           struct{}{},
		},
	},
	":man_pilot_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\u2708\ufe0f",
		map[string]struct{}{
			"man":                    struct{}{},
			"man pilot":              struct{}{},
			"medium-light skin tone": struct{}{},
			"pilot":                  struct{}{},
			"plane":                  struct{}{},
		},
	},
	":man_pilot_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\u2708\ufe0f",
		map[string]struct{}{
			"man":              struct{}{},
			"man pilot":        struct{}{},
			"medium skin tone": struct{}{},
			"pilot":            struct{}{},
			"plane":            struct{}{},
		},
	},
	":man_pilot_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\u2708\ufe0f",
		map[string]struct{}{
			"man":                   struct{}{},
			"man pilot":             struct{}{},
			"medium-dark skin tone": struct{}{},
			"pilot":                 struct{}{},
			"plane":                 struct{}{},
		},
	},
	":man_pilot_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\u2708\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"man":            struct{}{},
			"man pilot":      struct{}{},
			"pilot":          struct{}{},
			"plane":          struct{}{},
		},
	},
	":woman_pilot:": Code{
		"\U0001f469\u200d\u2708\ufe0f",
		map[string]struct{}{
			"pilot":       struct{}{},
			"plane":       struct{}{},
			"woman":       struct{}{},
			"woman pilot": struct{}{},
		},
	},
	":woman_pilot_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\u2708\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"pilot":           struct{}{},
			"plane":           struct{}{},
			"woman":           struct{}{},
			"woman pilot":     struct{}{},
		},
	},
	":woman_pilot_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\u2708\ufe0f",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"pilot":                  struct{}{},
			"plane":                  struct{}{},
			"woman":                  struct{}{},
			"woman pilot":            struct{}{},
		},
	},
	":woman_pilot_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\u2708\ufe0f",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"pilot":            struct{}{},
			"plane":            struct{}{},
			"woman":            struct{}{},
			"woman pilot":      struct{}{},
		},
	},
	":woman_pilot_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\u2708\ufe0f",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"pilot":                 struct{}{},
			"plane":                 struct{}{},
			"woman":                 struct{}{},
			"woman pilot":           struct{}{},
		},
	},
	":woman_pilot_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\u2708\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"pilot":          struct{}{},
			"plane":          struct{}{},
			"woman":          struct{}{},
			"woman pilot":    struct{}{},
		},
	},
	":man_astronaut:": Code{
		"\U0001f468\u200d\U0001f680",
		map[string]struct{}{
			"astronaut":     struct{}{},
			"man":           struct{}{},
			"man astronaut": struct{}{},
			"rocket":        struct{}{},
		},
	},
	":man_astronaut_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\U0001f680",
		map[string]struct{}{
			"astronaut":       struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man astronaut":   struct{}{},
			"rocket":          struct{}{},
		},
	},
	":man_astronaut_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\U0001f680",
		map[string]struct{}{
			"astronaut":              struct{}{},
			"man":                    struct{}{},
			"man astronaut":          struct{}{},
			"medium-light skin tone": struct{}{},
			"rocket":                 struct{}{},
		},
	},
	":man_astronaut_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\U0001f680",
		map[string]struct{}{
			"astronaut":        struct{}{},
			"man":              struct{}{},
			"man astronaut":    struct{}{},
			"medium skin tone": struct{}{},
			"rocket":           struct{}{},
		},
	},
	":man_astronaut_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\U0001f680",
		map[string]struct{}{
			"astronaut":             struct{}{},
			"man":                   struct{}{},
			"man astronaut":         struct{}{},
			"medium-dark skin tone": struct{}{},
			"rocket":                struct{}{},
		},
	},
	":man_astronaut_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\U0001f680",
		map[string]struct{}{
			"astronaut":      struct{}{},
			"dark skin tone": struct{}{},
			"man":            struct{}{},
			"man astronaut":  struct{}{},
			"rocket":         struct{}{},
		},
	},
	":woman_astronaut:": Code{
		"\U0001f469\u200d\U0001f680",
		map[string]struct{}{
			"astronaut":       struct{}{},
			"rocket":          struct{}{},
			"woman":           struct{}{},
			"woman astronaut": struct{}{},
		},
	},
	":woman_astronaut_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\U0001f680",
		map[string]struct{}{
			"astronaut":       struct{}{},
			"light skin tone": struct{}{},
			"rocket":          struct{}{},
			"woman":           struct{}{},
			"woman astronaut": struct{}{},
		},
	},
	":woman_astronaut_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\U0001f680",
		map[string]struct{}{
			"astronaut":              struct{}{},
			"medium-light skin tone": struct{}{},
			"rocket":                 struct{}{},
			"woman":                  struct{}{},
			"woman astronaut":        struct{}{},
		},
	},
	":woman_astronaut_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\U0001f680",
		map[string]struct{}{
			"astronaut":        struct{}{},
			"medium skin tone": struct{}{},
			"rocket":           struct{}{},
			"woman":            struct{}{},
			"woman astronaut":  struct{}{},
		},
	},
	":woman_astronaut_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\U0001f680",
		map[string]struct{}{
			"astronaut":             struct{}{},
			"medium-dark skin tone": struct{}{},
			"rocket":                struct{}{},
			"woman":                 struct{}{},
			"woman astronaut":       struct{}{},
		},
	},
	":woman_astronaut_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\U0001f680",
		map[string]struct{}{
			"astronaut":       struct{}{},
			"dark skin tone":  struct{}{},
			"rocket":          struct{}{},
			"woman":           struct{}{},
			"woman astronaut": struct{}{},
		},
	},
	":man_firefighter:": Code{
		"\U0001f468\u200d\U0001f692",
		map[string]struct{}{
			"firefighter":     struct{}{},
			"firetruck":       struct{}{},
			"man":             struct{}{},
			"man firefighter": struct{}{},
		},
	},
	":man_firefighter_tone1:": Code{
		"\U0001f468\U0001f3fb\u200d\U0001f692",
		map[string]struct{}{
			"firefighter":     struct{}{},
			"firetruck":       struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man firefighter": struct{}{},
		},
	},
	":man_firefighter_tone2:": Code{
		"\U0001f468\U0001f3fc\u200d\U0001f692",
		map[string]struct{}{
			"firefighter":            struct{}{},
			"firetruck":              struct{}{},
			"man":                    struct{}{},
			"man firefighter":        struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_firefighter_tone3:": Code{
		"\U0001f468\U0001f3fd\u200d\U0001f692",
		map[string]struct{}{
			"firefighter":      struct{}{},
			"firetruck":        struct{}{},
			"man":              struct{}{},
			"man firefighter":  struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":man_firefighter_tone4:": Code{
		"\U0001f468\U0001f3fe\u200d\U0001f692",
		map[string]struct{}{
			"firefighter":           struct{}{},
			"firetruck":             struct{}{},
			"man":                   struct{}{},
			"man firefighter":       struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_firefighter_tone5:": Code{
		"\U0001f468\U0001f3ff\u200d\U0001f692",
		map[string]struct{}{
			"dark skin tone":  struct{}{},
			"firefighter":     struct{}{},
			"firetruck":       struct{}{},
			"man":             struct{}{},
			"man firefighter": struct{}{},
		},
	},
	":woman_firefighter:": Code{
		"\U0001f469\u200d\U0001f692",
		map[string]struct{}{
			"firefighter":       struct{}{},
			"firetruck":         struct{}{},
			"woman":             struct{}{},
			"woman firefighter": struct{}{},
		},
	},
	":woman_firefighter_tone1:": Code{
		"\U0001f469\U0001f3fb\u200d\U0001f692",
		map[string]struct{}{
			"firefighter":       struct{}{},
			"firetruck":         struct{}{},
			"light skin tone":   struct{}{},
			"woman":             struct{}{},
			"woman firefighter": struct{}{},
		},
	},
	":woman_firefighter_tone2:": Code{
		"\U0001f469\U0001f3fc\u200d\U0001f692",
		map[string]struct{}{
			"firefighter":            struct{}{},
			"firetruck":              struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
			"woman firefighter":      struct{}{},
		},
	},
	":woman_firefighter_tone3:": Code{
		"\U0001f469\U0001f3fd\u200d\U0001f692",
		map[string]struct{}{
			"firefighter":       struct{}{},
			"firetruck":         struct{}{},
			"medium skin tone":  struct{}{},
			"woman":             struct{}{},
			"woman firefighter": struct{}{},
		},
	},
	":woman_firefighter_tone4:": Code{
		"\U0001f469\U0001f3fe\u200d\U0001f692",
		map[string]struct{}{
			"firefighter":           struct{}{},
			"firetruck":             struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
			"woman firefighter":     struct{}{},
		},
	},
	":woman_firefighter_tone5:": Code{
		"\U0001f469\U0001f3ff\u200d\U0001f692",
		map[string]struct{}{
			"dark skin tone":    struct{}{},
			"firefighter":       struct{}{},
			"firetruck":         struct{}{},
			"woman":             struct{}{},
			"woman firefighter": struct{}{},
		},
	},
	":police_officer:": Code{
		"\U0001f46e",
		map[string]struct{}{
			"cop":            struct{}{},
			"officer":        struct{}{},
			"police":         struct{}{},
			"police officer": struct{}{},
		},
	},
	":police_officer_tone1:": Code{
		"\U0001f46e\U0001f3fb",
		map[string]struct{}{
			"cop":             struct{}{},
			"light skin tone": struct{}{},
			"officer":         struct{}{},
			"police":          struct{}{},
			"police officer":  struct{}{},
		},
	},
	":police_officer_tone2:": Code{
		"\U0001f46e\U0001f3fc",
		map[string]struct{}{
			"cop":                    struct{}{},
			"medium-light skin tone": struct{}{},
			"officer":                struct{}{},
			"police":                 struct{}{},
			"police officer":         struct{}{},
		},
	},
	":police_officer_tone3:": Code{
		"\U0001f46e\U0001f3fd",
		map[string]struct{}{
			"cop":              struct{}{},
			"medium skin tone": struct{}{},
			"officer":          struct{}{},
			"police":           struct{}{},
			"police officer":   struct{}{},
		},
	},
	":police_officer_tone4:": Code{
		"\U0001f46e\U0001f3fe",
		map[string]struct{}{
			"cop":                   struct{}{},
			"medium-dark skin tone": struct{}{},
			"officer":               struct{}{},
			"police":                struct{}{},
			"police officer":        struct{}{},
		},
	},
	":police_officer_tone5:": Code{
		"\U0001f46e\U0001f3ff",
		map[string]struct{}{
			"cop":            struct{}{},
			"dark skin tone": struct{}{},
			"officer":        struct{}{},
			"police":         struct{}{},
			"police officer": struct{}{},
		},
	},
	":man_police_officer:": Code{
		"\U0001f46e\u200d\u2642\ufe0f",
		map[string]struct{}{
			"cop":                struct{}{},
			"man":                struct{}{},
			"man police officer": struct{}{},
			"officer":            struct{}{},
			"police":             struct{}{},
		},
	},
	":man_police_officer_tone1:": Code{
		"\U0001f46e\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"cop":                struct{}{},
			"light skin tone":    struct{}{},
			"man":                struct{}{},
			"man police officer": struct{}{},
			"officer":            struct{}{},
			"police":             struct{}{},
		},
	},
	":man_police_officer_tone2:": Code{
		"\U0001f46e\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"cop":                    struct{}{},
			"man":                    struct{}{},
			"man police officer":     struct{}{},
			"medium-light skin tone": struct{}{},
			"officer":                struct{}{},
			"police":                 struct{}{},
		},
	},
	":man_police_officer_tone3:": Code{
		"\U0001f46e\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"cop":                struct{}{},
			"man":                struct{}{},
			"man police officer": struct{}{},
			"medium skin tone":   struct{}{},
			"officer":            struct{}{},
			"police":             struct{}{},
		},
	},
	":man_police_officer_tone4:": Code{
		"\U0001f46e\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"cop":                   struct{}{},
			"man":                   struct{}{},
			"man police officer":    struct{}{},
			"medium-dark skin tone": struct{}{},
			"officer":               struct{}{},
			"police":                struct{}{},
		},
	},
	":man_police_officer_tone5:": Code{
		"\U0001f46e\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"cop":                struct{}{},
			"dark skin tone":     struct{}{},
			"man":                struct{}{},
			"man police officer": struct{}{},
			"officer":            struct{}{},
			"police":             struct{}{},
		},
	},
	":woman_police_officer:": Code{
		"\U0001f46e\u200d\u2640\ufe0f",
		map[string]struct{}{
			"cop":                  struct{}{},
			"officer":              struct{}{},
			"police":               struct{}{},
			"woman":                struct{}{},
			"woman police officer": struct{}{},
		},
	},
	":woman_police_officer_tone1:": Code{
		"\U0001f46e\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"cop":                  struct{}{},
			"light skin tone":      struct{}{},
			"officer":              struct{}{},
			"police":               struct{}{},
			"woman":                struct{}{},
			"woman police officer": struct{}{},
		},
	},
	":woman_police_officer_tone2:": Code{
		"\U0001f46e\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"cop":                    struct{}{},
			"medium-light skin tone": struct{}{},
			"officer":                struct{}{},
			"police":                 struct{}{},
			"woman":                  struct{}{},
			"woman police officer":   struct{}{},
		},
	},
	":woman_police_officer_tone3:": Code{
		"\U0001f46e\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"cop":                  struct{}{},
			"medium skin tone":     struct{}{},
			"officer":              struct{}{},
			"police":               struct{}{},
			"woman":                struct{}{},
			"woman police officer": struct{}{},
		},
	},
	":woman_police_officer_tone4:": Code{
		"\U0001f46e\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"cop":                   struct{}{},
			"medium-dark skin tone": struct{}{},
			"officer":               struct{}{},
			"police":                struct{}{},
			"woman":                 struct{}{},
			"woman police officer":  struct{}{},
		},
	},
	":woman_police_officer_tone5:": Code{
		"\U0001f46e\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"cop":                  struct{}{},
			"dark skin tone":       struct{}{},
			"officer":              struct{}{},
			"police":               struct{}{},
			"woman":                struct{}{},
			"woman police officer": struct{}{},
		},
	},
	":detective:": Code{
		"\U0001f575",
		map[string]struct{}{
			"detective": struct{}{},
			"sleuth":    struct{}{},
			"spy":       struct{}{},
		},
	},
	":detective_tone1:": Code{
		"\U0001f575\U0001f3fb",
		map[string]struct{}{
			"detective":       struct{}{},
			"light skin tone": struct{}{},
			"sleuth":          struct{}{},
			"spy":             struct{}{},
		},
	},
	":detective_tone2:": Code{
		"\U0001f575\U0001f3fc",
		map[string]struct{}{
			"detective":              struct{}{},
			"medium-light skin tone": struct{}{},
			"sleuth":                 struct{}{},
			"spy":                    struct{}{},
		},
	},
	":detective_tone3:": Code{
		"\U0001f575\U0001f3fd",
		map[string]struct{}{
			"detective":        struct{}{},
			"medium skin tone": struct{}{},
			"sleuth":           struct{}{},
			"spy":              struct{}{},
		},
	},
	":detective_tone4:": Code{
		"\U0001f575\U0001f3fe",
		map[string]struct{}{
			"detective":             struct{}{},
			"medium-dark skin tone": struct{}{},
			"sleuth":                struct{}{},
			"spy":                   struct{}{},
		},
	},
	":detective_tone5:": Code{
		"\U0001f575\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"detective":      struct{}{},
			"sleuth":         struct{}{},
			"spy":            struct{}{},
		},
	},
	":man_detective:": Code{
		"\U0001f575\ufe0f\u200d\u2642\ufe0f",
		map[string]struct{}{
			"detective":     struct{}{},
			"man":           struct{}{},
			"man detective": struct{}{},
			"sleuth":        struct{}{},
			"spy":           struct{}{},
		},
	},
	":man_detective_tone1:": Code{
		"\U0001f575\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"detective":       struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man detective":   struct{}{},
			"sleuth":          struct{}{},
			"spy":             struct{}{},
		},
	},
	":man_detective_tone2:": Code{
		"\U0001f575\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"detective":              struct{}{},
			"man":                    struct{}{},
			"man detective":          struct{}{},
			"medium-light skin tone": struct{}{},
			"sleuth":                 struct{}{},
			"spy":                    struct{}{},
		},
	},
	":man_detective_tone3:": Code{
		"\U0001f575\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"detective":        struct{}{},
			"man":              struct{}{},
			"man detective":    struct{}{},
			"medium skin tone": struct{}{},
			"sleuth":           struct{}{},
			"spy":              struct{}{},
		},
	},
	":man_detective_tone4:": Code{
		"\U0001f575\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"detective":             struct{}{},
			"man":                   struct{}{},
			"man detective":         struct{}{},
			"medium-dark skin tone": struct{}{},
			"sleuth":                struct{}{},
			"spy":                   struct{}{},
		},
	},
	":man_detective_tone5:": Code{
		"\U0001f575\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"detective":      struct{}{},
			"man":            struct{}{},
			"man detective":  struct{}{},
			"sleuth":         struct{}{},
			"spy":            struct{}{},
		},
	},
	":woman_detective:": Code{
		"\U0001f575\ufe0f\u200d\u2640\ufe0f",
		map[string]struct{}{
			"detective":       struct{}{},
			"sleuth":          struct{}{},
			"spy":             struct{}{},
			"woman":           struct{}{},
			"woman detective": struct{}{},
		},
	},
	":woman_detective_tone1:": Code{
		"\U0001f575\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"detective":       struct{}{},
			"light skin tone": struct{}{},
			"sleuth":          struct{}{},
			"spy":             struct{}{},
			"woman":           struct{}{},
			"woman detective": struct{}{},
		},
	},
	":woman_detective_tone2:": Code{
		"\U0001f575\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"detective":              struct{}{},
			"medium-light skin tone": struct{}{},
			"sleuth":                 struct{}{},
			"spy":                    struct{}{},
			"woman":                  struct{}{},
			"woman detective":        struct{}{},
		},
	},
	":woman_detective_tone3:": Code{
		"\U0001f575\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"detective":        struct{}{},
			"medium skin tone": struct{}{},
			"sleuth":           struct{}{},
			"spy":              struct{}{},
			"woman":            struct{}{},
			"woman detective":  struct{}{},
		},
	},
	":woman_detective_tone4:": Code{
		"\U0001f575\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"detective":             struct{}{},
			"medium-dark skin tone": struct{}{},
			"sleuth":                struct{}{},
			"spy":                   struct{}{},
			"woman":                 struct{}{},
			"woman detective":       struct{}{},
		},
	},
	":woman_detective_tone5:": Code{
		"\U0001f575\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":  struct{}{},
			"detective":       struct{}{},
			"sleuth":          struct{}{},
			"spy":             struct{}{},
			"woman":           struct{}{},
			"woman detective": struct{}{},
		},
	},
	":guard:": Code{
		"\U0001f482",
		map[string]struct{}{
			"guard": struct{}{},
		},
	},
	":guard_tone1:": Code{
		"\U0001f482\U0001f3fb",
		map[string]struct{}{
			"guard":           struct{}{},
			"light skin tone": struct{}{},
		},
	},
	":guard_tone2:": Code{
		"\U0001f482\U0001f3fc",
		map[string]struct{}{
			"guard":                  struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":guard_tone3:": Code{
		"\U0001f482\U0001f3fd",
		map[string]struct{}{
			"guard":            struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":guard_tone4:": Code{
		"\U0001f482\U0001f3fe",
		map[string]struct{}{
			"guard":                 struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":guard_tone5:": Code{
		"\U0001f482\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"guard":          struct{}{},
		},
	},
	":man_guard:": Code{
		"\U0001f482\u200d\u2642\ufe0f",
		map[string]struct{}{
			"guard":     struct{}{},
			"man":       struct{}{},
			"man guard": struct{}{},
		},
	},
	":man_guard_tone1:": Code{
		"\U0001f482\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"guard":           struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man guard":       struct{}{},
		},
	},
	":man_guard_tone2:": Code{
		"\U0001f482\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"guard":                  struct{}{},
			"man":                    struct{}{},
			"man guard":              struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_guard_tone3:": Code{
		"\U0001f482\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"guard":            struct{}{},
			"man":              struct{}{},
			"man guard":        struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":man_guard_tone4:": Code{
		"\U0001f482\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"guard":                 struct{}{},
			"man":                   struct{}{},
			"man guard":             struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_guard_tone5:": Code{
		"\U0001f482\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"guard":          struct{}{},
			"man":            struct{}{},
			"man guard":      struct{}{},
		},
	},
	":woman_guard:": Code{
		"\U0001f482\u200d\u2640\ufe0f",
		map[string]struct{}{
			"guard":       struct{}{},
			"woman":       struct{}{},
			"woman guard": struct{}{},
		},
	},
	":woman_guard_tone1:": Code{
		"\U0001f482\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"guard":           struct{}{},
			"light skin tone": struct{}{},
			"woman":           struct{}{},
			"woman guard":     struct{}{},
		},
	},
	":woman_guard_tone2:": Code{
		"\U0001f482\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"guard":                  struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
			"woman guard":            struct{}{},
		},
	},
	":woman_guard_tone3:": Code{
		"\U0001f482\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"guard":            struct{}{},
			"medium skin tone": struct{}{},
			"woman":            struct{}{},
			"woman guard":      struct{}{},
		},
	},
	":woman_guard_tone4:": Code{
		"\U0001f482\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"guard":                 struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
			"woman guard":           struct{}{},
		},
	},
	":woman_guard_tone5:": Code{
		"\U0001f482\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"guard":          struct{}{},
			"woman":          struct{}{},
			"woman guard":    struct{}{},
		},
	},
	":construction_worker:": Code{
		"\U0001f477",
		map[string]struct{}{
			"construction":        struct{}{},
			"construction worker": struct{}{},
			"hat":                 struct{}{},
			"worker":              struct{}{},
		},
	},
	":construction_worker_tone1:": Code{
		"\U0001f477\U0001f3fb",
		map[string]struct{}{
			"construction":        struct{}{},
			"construction worker": struct{}{},
			"hat":                 struct{}{},
			"light skin tone":     struct{}{},
			"worker":              struct{}{},
		},
	},
	":construction_worker_tone2:": Code{
		"\U0001f477\U0001f3fc",
		map[string]struct{}{
			"construction":           struct{}{},
			"construction worker":    struct{}{},
			"hat":                    struct{}{},
			"medium-light skin tone": struct{}{},
			"worker":                 struct{}{},
		},
	},
	":construction_worker_tone3:": Code{
		"\U0001f477\U0001f3fd",
		map[string]struct{}{
			"construction":        struct{}{},
			"construction worker": struct{}{},
			"hat":                 struct{}{},
			"medium skin tone":    struct{}{},
			"worker":              struct{}{},
		},
	},
	":construction_worker_tone4:": Code{
		"\U0001f477\U0001f3fe",
		map[string]struct{}{
			"construction":          struct{}{},
			"construction worker":   struct{}{},
			"hat":                   struct{}{},
			"medium-dark skin tone": struct{}{},
			"worker":                struct{}{},
		},
	},
	":construction_worker_tone5:": Code{
		"\U0001f477\U0001f3ff",
		map[string]struct{}{
			"construction":        struct{}{},
			"construction worker": struct{}{},
			"dark skin tone":      struct{}{},
			"hat":                 struct{}{},
			"worker":              struct{}{},
		},
	},
	":man_construction_worker:": Code{
		"\U0001f477\u200d\u2642\ufe0f",
		map[string]struct{}{
			"construction":            struct{}{},
			"man":                     struct{}{},
			"man construction worker": struct{}{},
			"worker":                  struct{}{},
		},
	},
	":man_construction_worker_tone1:": Code{
		"\U0001f477\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"construction":            struct{}{},
			"light skin tone":         struct{}{},
			"man":                     struct{}{},
			"man construction worker": struct{}{},
			"worker":                  struct{}{},
		},
	},
	":man_construction_worker_tone2:": Code{
		"\U0001f477\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"construction":            struct{}{},
			"man":                     struct{}{},
			"man construction worker": struct{}{},
			"medium-light skin tone":  struct{}{},
			"worker":                  struct{}{},
		},
	},
	":man_construction_worker_tone3:": Code{
		"\U0001f477\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"construction":            struct{}{},
			"man":                     struct{}{},
			"man construction worker": struct{}{},
			"medium skin tone":        struct{}{},
			"worker":                  struct{}{},
		},
	},
	":man_construction_worker_tone4:": Code{
		"\U0001f477\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"construction":            struct{}{},
			"man":                     struct{}{},
			"man construction worker": struct{}{},
			"medium-dark skin tone":   struct{}{},
			"worker":                  struct{}{},
		},
	},
	":man_construction_worker_tone5:": Code{
		"\U0001f477\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"construction":            struct{}{},
			"dark skin tone":          struct{}{},
			"man":                     struct{}{},
			"man construction worker": struct{}{},
			"worker":                  struct{}{},
		},
	},
	":woman_construction_worker:": Code{
		"\U0001f477\u200d\u2640\ufe0f",
		map[string]struct{}{
			"construction":              struct{}{},
			"woman":                     struct{}{},
			"woman construction worker": struct{}{},
			"worker":                    struct{}{},
		},
	},
	":woman_construction_worker_tone1:": Code{
		"\U0001f477\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"construction":              struct{}{},
			"light skin tone":           struct{}{},
			"woman":                     struct{}{},
			"woman construction worker": struct{}{},
			"worker":                    struct{}{},
		},
	},
	":woman_construction_worker_tone2:": Code{
		"\U0001f477\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"construction":              struct{}{},
			"medium-light skin tone":    struct{}{},
			"woman":                     struct{}{},
			"woman construction worker": struct{}{},
			"worker":                    struct{}{},
		},
	},
	":woman_construction_worker_tone3:": Code{
		"\U0001f477\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"construction":              struct{}{},
			"medium skin tone":          struct{}{},
			"woman":                     struct{}{},
			"woman construction worker": struct{}{},
			"worker":                    struct{}{},
		},
	},
	":woman_construction_worker_tone4:": Code{
		"\U0001f477\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"construction":              struct{}{},
			"medium-dark skin tone":     struct{}{},
			"woman":                     struct{}{},
			"woman construction worker": struct{}{},
			"worker":                    struct{}{},
		},
	},
	":woman_construction_worker_tone5:": Code{
		"\U0001f477\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"construction":              struct{}{},
			"dark skin tone":            struct{}{},
			"woman":                     struct{}{},
			"woman construction worker": struct{}{},
			"worker":                    struct{}{},
		},
	},
	":prince:": Code{
		"\U0001f934",
		map[string]struct{}{
			"prince": struct{}{},
		},
	},
	":prince_tone1:": Code{
		"\U0001f934\U0001f3fb",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"prince":          struct{}{},
		},
	},
	":prince_tone2:": Code{
		"\U0001f934\U0001f3fc",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"prince":                 struct{}{},
		},
	},
	":prince_tone3:": Code{
		"\U0001f934\U0001f3fd",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"prince":           struct{}{},
		},
	},
	":prince_tone4:": Code{
		"\U0001f934\U0001f3fe",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"prince":                struct{}{},
		},
	},
	":prince_tone5:": Code{
		"\U0001f934\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"prince":         struct{}{},
		},
	},
	":princess:": Code{
		"\U0001f478",
		map[string]struct{}{
			"fairy tale": struct{}{},
			"fantasy":    struct{}{},
			"princess":   struct{}{},
		},
	},
	":princess_tone1:": Code{
		"\U0001f478\U0001f3fb",
		map[string]struct{}{
			"fairy tale":      struct{}{},
			"fantasy":         struct{}{},
			"light skin tone": struct{}{},
			"princess":        struct{}{},
		},
	},
	":princess_tone2:": Code{
		"\U0001f478\U0001f3fc",
		map[string]struct{}{
			"fairy tale":             struct{}{},
			"fantasy":                struct{}{},
			"medium-light skin tone": struct{}{},
			"princess":               struct{}{},
		},
	},
	":princess_tone3:": Code{
		"\U0001f478\U0001f3fd",
		map[string]struct{}{
			"fairy tale":       struct{}{},
			"fantasy":          struct{}{},
			"medium skin tone": struct{}{},
			"princess":         struct{}{},
		},
	},
	":princess_tone4:": Code{
		"\U0001f478\U0001f3fe",
		map[string]struct{}{
			"fairy tale":            struct{}{},
			"fantasy":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"princess":              struct{}{},
		},
	},
	":princess_tone5:": Code{
		"\U0001f478\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"fairy tale":     struct{}{},
			"fantasy":        struct{}{},
			"princess":       struct{}{},
		},
	},
	":person_wearing_turban:": Code{
		"\U0001f473",
		map[string]struct{}{
			"person wearing turban": struct{}{},
			"turban":                struct{}{},
		},
	},
	":person_wearing_turban_tone1:": Code{
		"\U0001f473\U0001f3fb",
		map[string]struct{}{
			"light skin tone":       struct{}{},
			"person wearing turban": struct{}{},
			"turban":                struct{}{},
		},
	},
	":person_wearing_turban_tone2:": Code{
		"\U0001f473\U0001f3fc",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"person wearing turban":  struct{}{},
			"turban":                 struct{}{},
		},
	},
	":person_wearing_turban_tone3:": Code{
		"\U0001f473\U0001f3fd",
		map[string]struct{}{
			"medium skin tone":      struct{}{},
			"person wearing turban": struct{}{},
			"turban":                struct{}{},
		},
	},
	":person_wearing_turban_tone4:": Code{
		"\U0001f473\U0001f3fe",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"person wearing turban": struct{}{},
			"turban":                struct{}{},
		},
	},
	":person_wearing_turban_tone5:": Code{
		"\U0001f473\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":        struct{}{},
			"person wearing turban": struct{}{},
			"turban":                struct{}{},
		},
	},
	":man_wearing_turban:": Code{
		"\U0001f473\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                struct{}{},
			"man wearing turban": struct{}{},
			"turban":             struct{}{},
		},
	},
	":man_wearing_turban_tone1:": Code{
		"\U0001f473\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"light skin tone":    struct{}{},
			"man":                struct{}{},
			"man wearing turban": struct{}{},
			"turban":             struct{}{},
		},
	},
	":man_wearing_turban_tone2:": Code{
		"\U0001f473\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                    struct{}{},
			"man wearing turban":     struct{}{},
			"medium-light skin tone": struct{}{},
			"turban":                 struct{}{},
		},
	},
	":man_wearing_turban_tone3:": Code{
		"\U0001f473\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                struct{}{},
			"man wearing turban": struct{}{},
			"medium skin tone":   struct{}{},
			"turban":             struct{}{},
		},
	},
	":man_wearing_turban_tone4:": Code{
		"\U0001f473\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                   struct{}{},
			"man wearing turban":    struct{}{},
			"medium-dark skin tone": struct{}{},
			"turban":                struct{}{},
		},
	},
	":man_wearing_turban_tone5:": Code{
		"\U0001f473\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone":     struct{}{},
			"man":                struct{}{},
			"man wearing turban": struct{}{},
			"turban":             struct{}{},
		},
	},
	":woman_wearing_turban:": Code{
		"\U0001f473\u200d\u2640\ufe0f",
		map[string]struct{}{
			"turban":               struct{}{},
			"woman":                struct{}{},
			"woman wearing turban": struct{}{},
		},
	},
	":woman_wearing_turban_tone1:": Code{
		"\U0001f473\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"light skin tone":      struct{}{},
			"turban":               struct{}{},
			"woman":                struct{}{},
			"woman wearing turban": struct{}{},
		},
	},
	":woman_wearing_turban_tone2:": Code{
		"\U0001f473\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"turban":                 struct{}{},
			"woman":                  struct{}{},
			"woman wearing turban":   struct{}{},
		},
	},
	":woman_wearing_turban_tone3:": Code{
		"\U0001f473\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium skin tone":     struct{}{},
			"turban":               struct{}{},
			"woman":                struct{}{},
			"woman wearing turban": struct{}{},
		},
	},
	":woman_wearing_turban_tone4:": Code{
		"\U0001f473\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"turban":                struct{}{},
			"woman":                 struct{}{},
			"woman wearing turban":  struct{}{},
		},
	},
	":woman_wearing_turban_tone5:": Code{
		"\U0001f473\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":       struct{}{},
			"turban":               struct{}{},
			"woman":                struct{}{},
			"woman wearing turban": struct{}{},
		},
	},
	":man_with_chinese_cap:": Code{
		"\U0001f472",
		map[string]struct{}{
			"gua pi mao":           struct{}{},
			"hat":                  struct{}{},
			"man":                  struct{}{},
			"man with Chinese cap": struct{}{},
		},
	},
	":man_with_chinese_cap_tone1:": Code{
		"\U0001f472\U0001f3fb",
		map[string]struct{}{
			"gua pi mao":           struct{}{},
			"hat":                  struct{}{},
			"light skin tone":      struct{}{},
			"man":                  struct{}{},
			"man with Chinese cap": struct{}{},
		},
	},
	":man_with_chinese_cap_tone2:": Code{
		"\U0001f472\U0001f3fc",
		map[string]struct{}{
			"gua pi mao":             struct{}{},
			"hat":                    struct{}{},
			"man":                    struct{}{},
			"man with Chinese cap":   struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_with_chinese_cap_tone3:": Code{
		"\U0001f472\U0001f3fd",
		map[string]struct{}{
			"gua pi mao":           struct{}{},
			"hat":                  struct{}{},
			"man":                  struct{}{},
			"man with Chinese cap": struct{}{},
			"medium skin tone":     struct{}{},
		},
	},
	":man_with_chinese_cap_tone4:": Code{
		"\U0001f472\U0001f3fe",
		map[string]struct{}{
			"gua pi mao":            struct{}{},
			"hat":                   struct{}{},
			"man":                   struct{}{},
			"man with Chinese cap":  struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_with_chinese_cap_tone5:": Code{
		"\U0001f472\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":       struct{}{},
			"gua pi mao":           struct{}{},
			"hat":                  struct{}{},
			"man":                  struct{}{},
			"man with Chinese cap": struct{}{},
		},
	},
	":woman_with_headscarf:": Code{
		"\U0001f9d5",
		map[string]struct{}{
			"headscarf":            struct{}{},
			"hijab":                struct{}{},
			"mantilla":             struct{}{},
			"tichel":               struct{}{},
			"woman with headscarf": struct{}{},
			"bandana":              struct{}{},
			"head kerchief":        struct{}{},
		},
	},
	":woman_with_headscarf_tone1:": Code{
		"\U0001f9d5\U0001f3fb",
		map[string]struct{}{
			"headscarf":            struct{}{},
			"hijab":                struct{}{},
			"light skin tone":      struct{}{},
			"mantilla":             struct{}{},
			"tichel":               struct{}{},
			"woman with headscarf": struct{}{},
		},
	},
	":woman_with_headscarf_tone2:": Code{
		"\U0001f9d5\U0001f3fc",
		map[string]struct{}{
			"headscarf":              struct{}{},
			"hijab":                  struct{}{},
			"mantilla":               struct{}{},
			"medium-light skin tone": struct{}{},
			"tichel":                 struct{}{},
			"woman with headscarf":   struct{}{},
		},
	},
	":woman_with_headscarf_tone3:": Code{
		"\U0001f9d5\U0001f3fd",
		map[string]struct{}{
			"headscarf":            struct{}{},
			"hijab":                struct{}{},
			"mantilla":             struct{}{},
			"medium skin tone":     struct{}{},
			"tichel":               struct{}{},
			"woman with headscarf": struct{}{},
		},
	},
	":woman_with_headscarf_tone4:": Code{
		"\U0001f9d5\U0001f3fe",
		map[string]struct{}{
			"headscarf":             struct{}{},
			"hijab":                 struct{}{},
			"mantilla":              struct{}{},
			"medium-dark skin tone": struct{}{},
			"tichel":                struct{}{},
			"woman with headscarf":  struct{}{},
		},
	},
	":woman_with_headscarf_tone5:": Code{
		"\U0001f9d5\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":       struct{}{},
			"headscarf":            struct{}{},
			"hijab":                struct{}{},
			"mantilla":             struct{}{},
			"tichel":               struct{}{},
			"woman with headscarf": struct{}{},
		},
	},
	":bearded_person:": Code{
		"\U0001f9d4",
		map[string]struct{}{
			"beard":          struct{}{},
			"bearded person": struct{}{},
			"bewhiskered":    struct{}{},
		},
	},
	":bearded_person_tone1:": Code{
		"\U0001f9d4\U0001f3fb",
		map[string]struct{}{
			"beard":           struct{}{},
			"bearded person":  struct{}{},
			"light skin tone": struct{}{},
		},
	},
	":bearded_person_tone2:": Code{
		"\U0001f9d4\U0001f3fc",
		map[string]struct{}{
			"beard":                  struct{}{},
			"bearded person":         struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":bearded_person_tone3:": Code{
		"\U0001f9d4\U0001f3fd",
		map[string]struct{}{
			"beard":            struct{}{},
			"bearded person":   struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":bearded_person_tone4:": Code{
		"\U0001f9d4\U0001f3fe",
		map[string]struct{}{
			"beard":                 struct{}{},
			"bearded person":        struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":bearded_person_tone5:": Code{
		"\U0001f9d4\U0001f3ff",
		map[string]struct{}{
			"beard":          struct{}{},
			"bearded person": struct{}{},
			"dark skin tone": struct{}{},
		},
	},
	":blond_haired_person:": Code{
		"\U0001f471",
		map[string]struct{}{
			"blond":               struct{}{},
			"blond-haired person": struct{}{},
		},
	},
	":blond_haired_person_tone1:": Code{
		"\U0001f471\U0001f3fb",
		map[string]struct{}{
			"blond":               struct{}{},
			"blond-haired person": struct{}{},
			"light skin tone":     struct{}{},
		},
	},
	":blond_haired_person_tone2:": Code{
		"\U0001f471\U0001f3fc",
		map[string]struct{}{
			"blond":                  struct{}{},
			"blond-haired person":    struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":blond_haired_person_tone3:": Code{
		"\U0001f471\U0001f3fd",
		map[string]struct{}{
			"blond":               struct{}{},
			"blond-haired person": struct{}{},
			"medium skin tone":    struct{}{},
		},
	},
	":blond_haired_person_tone4:": Code{
		"\U0001f471\U0001f3fe",
		map[string]struct{}{
			"blond":                 struct{}{},
			"blond-haired person":   struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":blond_haired_person_tone5:": Code{
		"\U0001f471\U0001f3ff",
		map[string]struct{}{
			"blond":               struct{}{},
			"blond-haired person": struct{}{},
			"dark skin tone":      struct{}{},
		},
	},
	":blond-haired_man:": Code{
		"\U0001f471\u200d\u2642\ufe0f",
		map[string]struct{}{
			"blond":            struct{}{},
			"blond-haired man": struct{}{},
			"man":              struct{}{},
		},
	},
	":blond-haired_man_tone1:": Code{
		"\U0001f471\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"blond":            struct{}{},
			"blond-haired man": struct{}{},
			"light skin tone":  struct{}{},
			"man":              struct{}{},
		},
	},
	":blond-haired_man_tone2:": Code{
		"\U0001f471\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"blond":                  struct{}{},
			"blond-haired man":       struct{}{},
			"man":                    struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":blond-haired_man_tone3:": Code{
		"\U0001f471\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"blond":            struct{}{},
			"blond-haired man": struct{}{},
			"man":              struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":blond-haired_man_tone4:": Code{
		"\U0001f471\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"blond":                 struct{}{},
			"blond-haired man":      struct{}{},
			"man":                   struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":blond-haired_man_tone5:": Code{
		"\U0001f471\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"blond":            struct{}{},
			"blond-haired man": struct{}{},
			"dark skin tone":   struct{}{},
			"man":              struct{}{},
		},
	},
	":blond-haired_woman:": Code{
		"\U0001f471\u200d\u2640\ufe0f",
		map[string]struct{}{
			"blond-haired woman": struct{}{},
			"blonde":             struct{}{},
			"woman":              struct{}{},
		},
	},
	":blond-haired_woman_tone1:": Code{
		"\U0001f471\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"blond-haired woman": struct{}{},
			"blonde":             struct{}{},
			"light skin tone":    struct{}{},
			"woman":              struct{}{},
		},
	},
	":blond-haired_woman_tone2:": Code{
		"\U0001f471\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"blond-haired woman":     struct{}{},
			"blonde":                 struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
		},
	},
	":blond-haired_woman_tone3:": Code{
		"\U0001f471\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"blond-haired woman": struct{}{},
			"blonde":             struct{}{},
			"medium skin tone":   struct{}{},
			"woman":              struct{}{},
		},
	},
	":blond-haired_woman_tone4:": Code{
		"\U0001f471\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"blond-haired woman":    struct{}{},
			"blonde":                struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
		},
	},
	":blond-haired_woman_tone5:": Code{
		"\U0001f471\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"blond-haired woman": struct{}{},
			"blonde":             struct{}{},
			"dark skin tone":     struct{}{},
			"woman":              struct{}{},
		},
	},
	":man_in_tuxedo:": Code{
		"\U0001f935",
		map[string]struct{}{
			"groom":         struct{}{},
			"man":           struct{}{},
			"man in tuxedo": struct{}{},
			"tuxedo":        struct{}{},
		},
	},
	":man_in_tuxedo_tone1:": Code{
		"\U0001f935\U0001f3fb",
		map[string]struct{}{
			"groom":           struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man in tuxedo":   struct{}{},
			"tuxedo":          struct{}{},
		},
	},
	":man_in_tuxedo_tone2:": Code{
		"\U0001f935\U0001f3fc",
		map[string]struct{}{
			"groom":                  struct{}{},
			"man":                    struct{}{},
			"man in tuxedo":          struct{}{},
			"medium-light skin tone": struct{}{},
			"tuxedo":                 struct{}{},
		},
	},
	":man_in_tuxedo_tone3:": Code{
		"\U0001f935\U0001f3fd",
		map[string]struct{}{
			"groom":            struct{}{},
			"man":              struct{}{},
			"man in tuxedo":    struct{}{},
			"medium skin tone": struct{}{},
			"tuxedo":           struct{}{},
		},
	},
	":man_in_tuxedo_tone4:": Code{
		"\U0001f935\U0001f3fe",
		map[string]struct{}{
			"groom":                 struct{}{},
			"man":                   struct{}{},
			"man in tuxedo":         struct{}{},
			"medium-dark skin tone": struct{}{},
			"tuxedo":                struct{}{},
		},
	},
	":man_in_tuxedo_tone5:": Code{
		"\U0001f935\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"groom":          struct{}{},
			"man":            struct{}{},
			"man in tuxedo":  struct{}{},
			"tuxedo":         struct{}{},
		},
	},
	":bride_with_veil:": Code{
		"\U0001f470",
		map[string]struct{}{
			"bride":           struct{}{},
			"bride with veil": struct{}{},
			"veil":            struct{}{},
			"wedding":         struct{}{},
		},
	},
	":bride_with_veil_tone1:": Code{
		"\U0001f470\U0001f3fb",
		map[string]struct{}{
			"bride":           struct{}{},
			"bride with veil": struct{}{},
			"light skin tone": struct{}{},
			"veil":            struct{}{},
			"wedding":         struct{}{},
		},
	},
	":bride_with_veil_tone2:": Code{
		"\U0001f470\U0001f3fc",
		map[string]struct{}{
			"bride":                  struct{}{},
			"bride with veil":        struct{}{},
			"medium-light skin tone": struct{}{},
			"veil":                   struct{}{},
			"wedding":                struct{}{},
		},
	},
	":bride_with_veil_tone3:": Code{
		"\U0001f470\U0001f3fd",
		map[string]struct{}{
			"bride":            struct{}{},
			"bride with veil":  struct{}{},
			"medium skin tone": struct{}{},
			"veil":             struct{}{},
			"wedding":          struct{}{},
		},
	},
	":bride_with_veil_tone4:": Code{
		"\U0001f470\U0001f3fe",
		map[string]struct{}{
			"bride":                 struct{}{},
			"bride with veil":       struct{}{},
			"medium-dark skin tone": struct{}{},
			"veil":                  struct{}{},
			"wedding":               struct{}{},
		},
	},
	":bride_with_veil_tone5:": Code{
		"\U0001f470\U0001f3ff",
		map[string]struct{}{
			"bride":           struct{}{},
			"bride with veil": struct{}{},
			"dark skin tone":  struct{}{},
			"veil":            struct{}{},
			"wedding":         struct{}{},
		},
	},
	":pregnant_woman:": Code{
		"\U0001f930",
		map[string]struct{}{
			"pregnant":       struct{}{},
			"pregnant woman": struct{}{},
			"woman":          struct{}{},
		},
	},
	":pregnant_woman_tone1:": Code{
		"\U0001f930\U0001f3fb",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"pregnant":        struct{}{},
			"pregnant woman":  struct{}{},
			"woman":           struct{}{},
		},
	},
	":pregnant_woman_tone2:": Code{
		"\U0001f930\U0001f3fc",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"pregnant":               struct{}{},
			"pregnant woman":         struct{}{},
			"woman":                  struct{}{},
		},
	},
	":pregnant_woman_tone3:": Code{
		"\U0001f930\U0001f3fd",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"pregnant":         struct{}{},
			"pregnant woman":   struct{}{},
			"woman":            struct{}{},
		},
	},
	":pregnant_woman_tone4:": Code{
		"\U0001f930\U0001f3fe",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"pregnant":              struct{}{},
			"pregnant woman":        struct{}{},
			"woman":                 struct{}{},
		},
	},
	":pregnant_woman_tone5:": Code{
		"\U0001f930\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"pregnant":       struct{}{},
			"pregnant woman": struct{}{},
			"woman":          struct{}{},
		},
	},
	":breast_feeding:": Code{
		"\U0001f931",
		map[string]struct{}{
			"baby":           struct{}{},
			"breast":         struct{}{},
			"breast-feeding": struct{}{},
			"nursing":        struct{}{},
		},
	},
	":breast_feeding_tone1:": Code{
		"\U0001f931\U0001f3fb",
		map[string]struct{}{
			"baby":            struct{}{},
			"breast":          struct{}{},
			"breast-feeding":  struct{}{},
			"light skin tone": struct{}{},
			"nursing":         struct{}{},
		},
	},
	":breast_feeding_tone2:": Code{
		"\U0001f931\U0001f3fc",
		map[string]struct{}{
			"baby":                   struct{}{},
			"breast":                 struct{}{},
			"breast-feeding":         struct{}{},
			"medium-light skin tone": struct{}{},
			"nursing":                struct{}{},
		},
	},
	":breast_feeding_tone3:": Code{
		"\U0001f931\U0001f3fd",
		map[string]struct{}{
			"baby":             struct{}{},
			"breast":           struct{}{},
			"breast-feeding":   struct{}{},
			"medium skin tone": struct{}{},
			"nursing":          struct{}{},
		},
	},
	":breast_feeding_tone4:": Code{
		"\U0001f931\U0001f3fe",
		map[string]struct{}{
			"baby":                  struct{}{},
			"breast":                struct{}{},
			"breast-feeding":        struct{}{},
			"medium-dark skin tone": struct{}{},
			"nursing":               struct{}{},
		},
	},
	":breast_feeding_tone5:": Code{
		"\U0001f931\U0001f3ff",
		map[string]struct{}{
			"baby":           struct{}{},
			"breast":         struct{}{},
			"breast-feeding": struct{}{},
			"dark skin tone": struct{}{},
			"nursing":        struct{}{},
		},
	},
	":angel:": Code{
		"\U0001f47c",
		map[string]struct{}{
			"angel":      struct{}{},
			"baby":       struct{}{},
			"baby angel": struct{}{},
			"face":       struct{}{},
			"fairy tale": struct{}{},
			"fantasy":    struct{}{},
		},
	},
	":angel_tone1:": Code{
		"\U0001f47c\U0001f3fb",
		map[string]struct{}{
			"angel":           struct{}{},
			"baby":            struct{}{},
			"baby angel":      struct{}{},
			"face":            struct{}{},
			"fairy tale":      struct{}{},
			"fantasy":         struct{}{},
			"light skin tone": struct{}{},
		},
	},
	":angel_tone2:": Code{
		"\U0001f47c\U0001f3fc",
		map[string]struct{}{
			"angel":                  struct{}{},
			"baby":                   struct{}{},
			"baby angel":             struct{}{},
			"face":                   struct{}{},
			"fairy tale":             struct{}{},
			"fantasy":                struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":angel_tone3:": Code{
		"\U0001f47c\U0001f3fd",
		map[string]struct{}{
			"angel":            struct{}{},
			"baby":             struct{}{},
			"baby angel":       struct{}{},
			"face":             struct{}{},
			"fairy tale":       struct{}{},
			"fantasy":          struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":angel_tone4:": Code{
		"\U0001f47c\U0001f3fe",
		map[string]struct{}{
			"angel":                 struct{}{},
			"baby":                  struct{}{},
			"baby angel":            struct{}{},
			"face":                  struct{}{},
			"fairy tale":            struct{}{},
			"fantasy":               struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":angel_tone5:": Code{
		"\U0001f47c\U0001f3ff",
		map[string]struct{}{
			"angel":          struct{}{},
			"baby":           struct{}{},
			"baby angel":     struct{}{},
			"dark skin tone": struct{}{},
			"face":           struct{}{},
			"fairy tale":     struct{}{},
			"fantasy":        struct{}{},
		},
	},
	":santa:": Code{
		"\U0001f385",
		map[string]struct{}{
			"celebration": struct{}{},
			"Christmas":   struct{}{},
			"claus":       struct{}{},
			"father":      struct{}{},
			"santa":       struct{}{},
			"Santa Claus": struct{}{},
		},
	},
	":santa_tone1:": Code{
		"\U0001f385\U0001f3fb",
		map[string]struct{}{
			"celebration":     struct{}{},
			"Christmas":       struct{}{},
			"claus":           struct{}{},
			"father":          struct{}{},
			"light skin tone": struct{}{},
			"santa":           struct{}{},
			"Santa Claus":     struct{}{},
		},
	},
	":santa_tone2:": Code{
		"\U0001f385\U0001f3fc",
		map[string]struct{}{
			"celebration":            struct{}{},
			"Christmas":              struct{}{},
			"claus":                  struct{}{},
			"father":                 struct{}{},
			"medium-light skin tone": struct{}{},
			"santa":                  struct{}{},
			"Santa Claus":            struct{}{},
		},
	},
	":santa_tone3:": Code{
		"\U0001f385\U0001f3fd",
		map[string]struct{}{
			"celebration":      struct{}{},
			"Christmas":        struct{}{},
			"claus":            struct{}{},
			"father":           struct{}{},
			"medium skin tone": struct{}{},
			"santa":            struct{}{},
			"Santa Claus":      struct{}{},
		},
	},
	":santa_tone4:": Code{
		"\U0001f385\U0001f3fe",
		map[string]struct{}{
			"celebration":           struct{}{},
			"Christmas":             struct{}{},
			"claus":                 struct{}{},
			"father":                struct{}{},
			"medium-dark skin tone": struct{}{},
			"santa":                 struct{}{},
			"Santa Claus":           struct{}{},
		},
	},
	":santa_tone5:": Code{
		"\U0001f385\U0001f3ff",
		map[string]struct{}{
			"celebration":    struct{}{},
			"Christmas":      struct{}{},
			"claus":          struct{}{},
			"dark skin tone": struct{}{},
			"father":         struct{}{},
			"santa":          struct{}{},
			"Santa Claus":    struct{}{},
		},
	},
	":mrs_claus:": Code{
		"\U0001f936",
		map[string]struct{}{
			"celebration": struct{}{},
			"Christmas":   struct{}{},
			"claus":       struct{}{},
			"mother":      struct{}{},
			"Mrs.":        struct{}{},
			"Mrs. Claus":  struct{}{},
		},
	},
	":mrs_claus_tone1:": Code{
		"\U0001f936\U0001f3fb",
		map[string]struct{}{
			"celebration":     struct{}{},
			"Christmas":       struct{}{},
			"claus":           struct{}{},
			"light skin tone": struct{}{},
			"mother":          struct{}{},
			"Mrs.":            struct{}{},
			"Mrs. Claus":      struct{}{},
		},
	},
	":mrs_claus_tone2:": Code{
		"\U0001f936\U0001f3fc",
		map[string]struct{}{
			"celebration":            struct{}{},
			"Christmas":              struct{}{},
			"claus":                  struct{}{},
			"medium-light skin tone": struct{}{},
			"mother":                 struct{}{},
			"Mrs.":                   struct{}{},
			"Mrs. Claus":             struct{}{},
		},
	},
	":mrs_claus_tone3:": Code{
		"\U0001f936\U0001f3fd",
		map[string]struct{}{
			"celebration":      struct{}{},
			"Christmas":        struct{}{},
			"claus":            struct{}{},
			"medium skin tone": struct{}{},
			"mother":           struct{}{},
			"Mrs.":             struct{}{},
			"Mrs. Claus":       struct{}{},
		},
	},
	":mrs_claus_tone4:": Code{
		"\U0001f936\U0001f3fe",
		map[string]struct{}{
			"celebration":           struct{}{},
			"Christmas":             struct{}{},
			"claus":                 struct{}{},
			"medium-dark skin tone": struct{}{},
			"mother":                struct{}{},
			"Mrs.":                  struct{}{},
			"Mrs. Claus":            struct{}{},
		},
	},
	":mrs_claus_tone5:": Code{
		"\U0001f936\U0001f3ff",
		map[string]struct{}{
			"celebration":    struct{}{},
			"Christmas":      struct{}{},
			"claus":          struct{}{},
			"dark skin tone": struct{}{},
			"mother":         struct{}{},
			"Mrs.":           struct{}{},
			"Mrs. Claus":     struct{}{},
		},
	},
	":mage:": Code{
		"\U0001f9d9",
		map[string]struct{}{
			"mage":      struct{}{},
			"sorcerer":  struct{}{},
			"sorceress": struct{}{},
			"witch":     struct{}{},
			"wizard":    struct{}{},
		},
	},
	":mage_tone1:": Code{
		"\U0001f9d9\U0001f3fb",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"mage":            struct{}{},
			"sorcerer":        struct{}{},
			"sorceress":       struct{}{},
			"witch":           struct{}{},
			"wizard":          struct{}{},
		},
	},
	":mage_tone2:": Code{
		"\U0001f9d9\U0001f3fc",
		map[string]struct{}{
			"mage":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"sorcerer":               struct{}{},
			"sorceress":              struct{}{},
			"witch":                  struct{}{},
			"wizard":                 struct{}{},
		},
	},
	":mage_tone3:": Code{
		"\U0001f9d9\U0001f3fd",
		map[string]struct{}{
			"mage":             struct{}{},
			"medium skin tone": struct{}{},
			"sorcerer":         struct{}{},
			"sorceress":        struct{}{},
			"witch":            struct{}{},
			"wizard":           struct{}{},
		},
	},
	":mage_tone4:": Code{
		"\U0001f9d9\U0001f3fe",
		map[string]struct{}{
			"mage":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"sorcerer":              struct{}{},
			"sorceress":             struct{}{},
			"witch":                 struct{}{},
			"wizard":                struct{}{},
		},
	},
	":mage_tone5:": Code{
		"\U0001f9d9\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"mage":           struct{}{},
			"sorcerer":       struct{}{},
			"sorceress":      struct{}{},
			"witch":          struct{}{},
			"wizard":         struct{}{},
		},
	},
	":woman_mage:": Code{
		"\U0001f9d9\u200d\u2640\ufe0f",
		map[string]struct{}{
			"sorceress":  struct{}{},
			"witch":      struct{}{},
			"woman mage": struct{}{},
		},
	},
	":woman_mage_tone1:": Code{
		"\U0001f9d9\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"sorceress":       struct{}{},
			"witch":           struct{}{},
			"woman mage":      struct{}{},
		},
	},
	":woman_mage_tone2:": Code{
		"\U0001f9d9\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"sorceress":              struct{}{},
			"witch":                  struct{}{},
			"woman mage":             struct{}{},
		},
	},
	":woman_mage_tone3:": Code{
		"\U0001f9d9\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"sorceress":        struct{}{},
			"witch":            struct{}{},
			"woman mage":       struct{}{},
		},
	},
	":woman_mage_tone4:": Code{
		"\U0001f9d9\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"sorceress":             struct{}{},
			"witch":                 struct{}{},
			"woman mage":            struct{}{},
		},
	},
	":woman_mage_tone5:": Code{
		"\U0001f9d9\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"sorceress":      struct{}{},
			"witch":          struct{}{},
			"woman mage":     struct{}{},
		},
	},
	":man_mage:": Code{
		"\U0001f9d9\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man mage": struct{}{},
			"sorcerer": struct{}{},
			"wizard":   struct{}{},
		},
	},
	":man_mage_tone1:": Code{
		"\U0001f9d9\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"man mage":        struct{}{},
			"sorcerer":        struct{}{},
			"wizard":          struct{}{},
		},
	},
	":man_mage_tone2:": Code{
		"\U0001f9d9\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man mage":               struct{}{},
			"medium-light skin tone": struct{}{},
			"sorcerer":               struct{}{},
			"wizard":                 struct{}{},
		},
	},
	":man_mage_tone3:": Code{
		"\U0001f9d9\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man mage":         struct{}{},
			"medium skin tone": struct{}{},
			"sorcerer":         struct{}{},
			"wizard":           struct{}{},
		},
	},
	":man_mage_tone4:": Code{
		"\U0001f9d9\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man mage":              struct{}{},
			"medium-dark skin tone": struct{}{},
			"sorcerer":              struct{}{},
			"wizard":                struct{}{},
		},
	},
	":man_mage_tone5:": Code{
		"\U0001f9d9\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"man mage":       struct{}{},
			"sorcerer":       struct{}{},
			"wizard":         struct{}{},
		},
	},
	":fairy:": Code{
		"\U0001f9da",
		map[string]struct{}{
			"fairy":   struct{}{},
			"Oberon":  struct{}{},
			"Puck":    struct{}{},
			"Titania": struct{}{},
		},
	},
	":fairy_tone1:": Code{
		"\U0001f9da\U0001f3fb",
		map[string]struct{}{
			"fairy":           struct{}{},
			"light skin tone": struct{}{},
			"Oberon":          struct{}{},
			"Puck":            struct{}{},
			"Titania":         struct{}{},
		},
	},
	":fairy_tone2:": Code{
		"\U0001f9da\U0001f3fc",
		map[string]struct{}{
			"fairy":                  struct{}{},
			"medium-light skin tone": struct{}{},
			"Oberon":                 struct{}{},
			"Puck":                   struct{}{},
			"Titania":                struct{}{},
		},
	},
	":fairy_tone3:": Code{
		"\U0001f9da\U0001f3fd",
		map[string]struct{}{
			"fairy":            struct{}{},
			"medium skin tone": struct{}{},
			"Oberon":           struct{}{},
			"Puck":             struct{}{},
			"Titania":          struct{}{},
		},
	},
	":fairy_tone4:": Code{
		"\U0001f9da\U0001f3fe",
		map[string]struct{}{
			"fairy":                 struct{}{},
			"medium-dark skin tone": struct{}{},
			"Oberon":                struct{}{},
			"Puck":                  struct{}{},
			"Titania":               struct{}{},
		},
	},
	":fairy_tone5:": Code{
		"\U0001f9da\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"fairy":          struct{}{},
			"Oberon":         struct{}{},
			"Puck":           struct{}{},
			"Titania":        struct{}{},
		},
	},
	":woman_fairy:": Code{
		"\U0001f9da\u200d\u2640\ufe0f",
		map[string]struct{}{
			"Titania":     struct{}{},
			"woman fairy": struct{}{},
		},
	},
	":woman_fairy_tone1:": Code{
		"\U0001f9da\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"Titania":         struct{}{},
			"woman fairy":     struct{}{},
		},
	},
	":woman_fairy_tone2:": Code{
		"\U0001f9da\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"Titania":                struct{}{},
			"woman fairy":            struct{}{},
		},
	},
	":woman_fairy_tone3:": Code{
		"\U0001f9da\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"Titania":          struct{}{},
			"woman fairy":      struct{}{},
		},
	},
	":woman_fairy_tone4:": Code{
		"\U0001f9da\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"Titania":               struct{}{},
			"woman fairy":           struct{}{},
		},
	},
	":woman_fairy_tone5:": Code{
		"\U0001f9da\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"Titania":        struct{}{},
			"woman fairy":    struct{}{},
		},
	},
	":man_fairy:": Code{
		"\U0001f9da\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man fairy": struct{}{},
			"Oberon":    struct{}{},
			"Puck":      struct{}{},
		},
	},
	":man_fairy_tone1:": Code{
		"\U0001f9da\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"man fairy":       struct{}{},
			"Oberon":          struct{}{},
			"Puck":            struct{}{},
		},
	},
	":man_fairy_tone2:": Code{
		"\U0001f9da\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man fairy":              struct{}{},
			"medium-light skin tone": struct{}{},
			"Oberon":                 struct{}{},
			"Puck":                   struct{}{},
		},
	},
	":man_fairy_tone3:": Code{
		"\U0001f9da\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man fairy":        struct{}{},
			"medium skin tone": struct{}{},
			"Oberon":           struct{}{},
			"Puck":             struct{}{},
		},
	},
	":man_fairy_tone4:": Code{
		"\U0001f9da\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man fairy":             struct{}{},
			"medium-dark skin tone": struct{}{},
			"Oberon":                struct{}{},
			"Puck":                  struct{}{},
		},
	},
	":man_fairy_tone5:": Code{
		"\U0001f9da\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"man fairy":      struct{}{},
			"Oberon":         struct{}{},
			"Puck":           struct{}{},
		},
	},
	":vampire:": Code{
		"\U0001f9db",
		map[string]struct{}{
			"Dracula": struct{}{},
			"undead":  struct{}{},
			"vampire": struct{}{},
		},
	},
	":vampire_tone1:": Code{
		"\U0001f9db\U0001f3fb",
		map[string]struct{}{
			"Dracula":         struct{}{},
			"light skin tone": struct{}{},
			"undead":          struct{}{},
			"vampire":         struct{}{},
		},
	},
	":vampire_tone2:": Code{
		"\U0001f9db\U0001f3fc",
		map[string]struct{}{
			"Dracula":                struct{}{},
			"medium-light skin tone": struct{}{},
			"undead":                 struct{}{},
			"vampire":                struct{}{},
		},
	},
	":vampire_tone3:": Code{
		"\U0001f9db\U0001f3fd",
		map[string]struct{}{
			"Dracula":          struct{}{},
			"medium skin tone": struct{}{},
			"undead":           struct{}{},
			"vampire":          struct{}{},
		},
	},
	":vampire_tone4:": Code{
		"\U0001f9db\U0001f3fe",
		map[string]struct{}{
			"Dracula":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"undead":                struct{}{},
			"vampire":               struct{}{},
		},
	},
	":vampire_tone5:": Code{
		"\U0001f9db\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"Dracula":        struct{}{},
			"undead":         struct{}{},
			"vampire":        struct{}{},
		},
	},
	":woman_vampire:": Code{
		"\U0001f9db\u200d\u2640\ufe0f",
		map[string]struct{}{
			"undead":        struct{}{},
			"woman vampire": struct{}{},
		},
	},
	":woman_vampire_tone1:": Code{
		"\U0001f9db\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"undead":          struct{}{},
			"woman vampire":   struct{}{},
		},
	},
	":woman_vampire_tone2:": Code{
		"\U0001f9db\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"undead":                 struct{}{},
			"woman vampire":          struct{}{},
		},
	},
	":woman_vampire_tone3:": Code{
		"\U0001f9db\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"undead":           struct{}{},
			"woman vampire":    struct{}{},
		},
	},
	":woman_vampire_tone4:": Code{
		"\U0001f9db\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"undead":                struct{}{},
			"woman vampire":         struct{}{},
		},
	},
	":woman_vampire_tone5:": Code{
		"\U0001f9db\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"undead":         struct{}{},
			"woman vampire":  struct{}{},
		},
	},
	":man_vampire:": Code{
		"\U0001f9db\u200d\u2642\ufe0f",
		map[string]struct{}{
			"Dracula":     struct{}{},
			"man vampire": struct{}{},
			"undead":      struct{}{},
		},
	},
	":man_vampire_tone1:": Code{
		"\U0001f9db\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"Dracula":         struct{}{},
			"light skin tone": struct{}{},
			"man vampire":     struct{}{},
			"undead":          struct{}{},
		},
	},
	":man_vampire_tone2:": Code{
		"\U0001f9db\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"Dracula":                struct{}{},
			"man vampire":            struct{}{},
			"medium-light skin tone": struct{}{},
			"undead":                 struct{}{},
		},
	},
	":man_vampire_tone3:": Code{
		"\U0001f9db\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"Dracula":          struct{}{},
			"man vampire":      struct{}{},
			"medium skin tone": struct{}{},
			"undead":           struct{}{},
		},
	},
	":man_vampire_tone4:": Code{
		"\U0001f9db\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"Dracula":               struct{}{},
			"man vampire":           struct{}{},
			"medium-dark skin tone": struct{}{},
			"undead":                struct{}{},
		},
	},
	":man_vampire_tone5:": Code{
		"\U0001f9db\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"Dracula":        struct{}{},
			"man vampire":    struct{}{},
			"undead":         struct{}{},
		},
	},
	":merperson:": Code{
		"\U0001f9dc",
		map[string]struct{}{
			"mermaid":   struct{}{},
			"merman":    struct{}{},
			"merperson": struct{}{},
			"merwoman":  struct{}{},
		},
	},
	":merperson_tone1:": Code{
		"\U0001f9dc\U0001f3fb",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"mermaid":         struct{}{},
			"merman":          struct{}{},
			"merperson":       struct{}{},
			"merwoman":        struct{}{},
		},
	},
	":merperson_tone2:": Code{
		"\U0001f9dc\U0001f3fc",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"mermaid":                struct{}{},
			"merman":                 struct{}{},
			"merperson":              struct{}{},
			"merwoman":               struct{}{},
		},
	},
	":merperson_tone3:": Code{
		"\U0001f9dc\U0001f3fd",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"mermaid":          struct{}{},
			"merman":           struct{}{},
			"merperson":        struct{}{},
			"merwoman":         struct{}{},
		},
	},
	":merperson_tone4:": Code{
		"\U0001f9dc\U0001f3fe",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"mermaid":               struct{}{},
			"merman":                struct{}{},
			"merperson":             struct{}{},
			"merwoman":              struct{}{},
		},
	},
	":merperson_tone5:": Code{
		"\U0001f9dc\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"mermaid":        struct{}{},
			"merman":         struct{}{},
			"merperson":      struct{}{},
			"merwoman":       struct{}{},
		},
	},
	":mermaid:": Code{
		"\U0001f9dc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"mermaid":  struct{}{},
			"merwoman": struct{}{},
		},
	},
	":mermaid_tone1:": Code{
		"\U0001f9dc\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"mermaid":         struct{}{},
			"merwoman":        struct{}{},
		},
	},
	":mermaid_tone2:": Code{
		"\U0001f9dc\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"mermaid":                struct{}{},
			"merwoman":               struct{}{},
		},
	},
	":mermaid_tone3:": Code{
		"\U0001f9dc\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"mermaid":          struct{}{},
			"merwoman":         struct{}{},
		},
	},
	":mermaid_tone4:": Code{
		"\U0001f9dc\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"mermaid":               struct{}{},
			"merwoman":              struct{}{},
		},
	},
	":mermaid_tone5:": Code{
		"\U0001f9dc\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"mermaid":        struct{}{},
			"merwoman":       struct{}{},
		},
	},
	":merman:": Code{
		"\U0001f9dc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"merman": struct{}{},
			"Triton": struct{}{},
		},
	},
	":merman_tone1:": Code{
		"\U0001f9dc\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"merman":          struct{}{},
			"Triton":          struct{}{},
		},
	},
	":merman_tone2:": Code{
		"\U0001f9dc\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"merman":                 struct{}{},
			"Triton":                 struct{}{},
		},
	},
	":merman_tone3:": Code{
		"\U0001f9dc\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"merman":           struct{}{},
			"Triton":           struct{}{},
		},
	},
	":merman_tone4:": Code{
		"\U0001f9dc\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"merman":                struct{}{},
			"Triton":                struct{}{},
		},
	},
	":merman_tone5:": Code{
		"\U0001f9dc\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"merman":         struct{}{},
			"Triton":         struct{}{},
		},
	},
	":elf:": Code{
		"\U0001f9dd",
		map[string]struct{}{
			"elf":        struct{}{},
			"magical":    struct{}{},
			"LOTR style": struct{}{},
		},
	},
	":elf_tone1:": Code{
		"\U0001f9dd\U0001f3fb",
		map[string]struct{}{
			"elf":             struct{}{},
			"light skin tone": struct{}{},
			"magical":         struct{}{},
		},
	},
	":elf_tone2:": Code{
		"\U0001f9dd\U0001f3fc",
		map[string]struct{}{
			"elf":                    struct{}{},
			"magical":                struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":elf_tone3:": Code{
		"\U0001f9dd\U0001f3fd",
		map[string]struct{}{
			"elf":              struct{}{},
			"magical":          struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":elf_tone4:": Code{
		"\U0001f9dd\U0001f3fe",
		map[string]struct{}{
			"elf":                   struct{}{},
			"magical":               struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":elf_tone5:": Code{
		"\U0001f9dd\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"elf":            struct{}{},
			"magical":        struct{}{},
		},
	},
	":woman_elf:": Code{
		"\U0001f9dd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"magical":   struct{}{},
			"woman elf": struct{}{},
		},
	},
	":woman_elf_tone1:": Code{
		"\U0001f9dd\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"magical":         struct{}{},
			"woman elf":       struct{}{},
		},
	},
	":woman_elf_tone2:": Code{
		"\U0001f9dd\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"magical":                struct{}{},
			"medium-light skin tone": struct{}{},
			"woman elf":              struct{}{},
		},
	},
	":woman_elf_tone3:": Code{
		"\U0001f9dd\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"magical":          struct{}{},
			"medium skin tone": struct{}{},
			"woman elf":        struct{}{},
		},
	},
	":woman_elf_tone4:": Code{
		"\U0001f9dd\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"magical":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman elf":             struct{}{},
		},
	},
	":woman_elf_tone5:": Code{
		"\U0001f9dd\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"magical":        struct{}{},
			"woman elf":      struct{}{},
		},
	},
	":man_elf:": Code{
		"\U0001f9dd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"magical": struct{}{},
			"man elf": struct{}{},
		},
	},
	":man_elf_tone1:": Code{
		"\U0001f9dd\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"magical":         struct{}{},
			"man elf":         struct{}{},
		},
	},
	":man_elf_tone2:": Code{
		"\U0001f9dd\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"magical":                struct{}{},
			"man elf":                struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_elf_tone3:": Code{
		"\U0001f9dd\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"magical":          struct{}{},
			"man elf":          struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":man_elf_tone4:": Code{
		"\U0001f9dd\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"magical":               struct{}{},
			"man elf":               struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_elf_tone5:": Code{
		"\U0001f9dd\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"magical":        struct{}{},
			"man elf":        struct{}{},
		},
	},
	":genie:": Code{
		"\U0001f9de",
		map[string]struct{}{
			"djinn":             struct{}{},
			"genie":             struct{}{},
			"(non-human color)": struct{}{},
		},
	},
	":woman_genie:": Code{
		"\U0001f9de\u200d\u2640\ufe0f",
		map[string]struct{}{
			"djinn":       struct{}{},
			"woman genie": struct{}{},
		},
	},
	":man_genie:": Code{
		"\U0001f9de\u200d\u2642\ufe0f",
		map[string]struct{}{
			"djinn":     struct{}{},
			"man genie": struct{}{},
		},
	},
	":zombie:": Code{
		"\U0001f9df",
		map[string]struct{}{
			"undead":            struct{}{},
			"walking dead":      struct{}{},
			"zombie":            struct{}{},
			"(non-human color)": struct{}{},
		},
	},
	":woman_zombie:": Code{
		"\U0001f9df\u200d\u2640\ufe0f",
		map[string]struct{}{
			"undead":       struct{}{},
			"walking dead": struct{}{},
			"woman zombie": struct{}{},
		},
	},
	":man_zombie:": Code{
		"\U0001f9df\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man zombie":   struct{}{},
			"undead":       struct{}{},
			"walking dead": struct{}{},
		},
	},
	":person_frowning:": Code{
		"\U0001f64d",
		map[string]struct{}{
			"frown":           struct{}{},
			"gesture":         struct{}{},
			"person frowning": struct{}{},
		},
	},
	":person_frowning_tone1:": Code{
		"\U0001f64d\U0001f3fb",
		map[string]struct{}{
			"frown":           struct{}{},
			"gesture":         struct{}{},
			"light skin tone": struct{}{},
			"person frowning": struct{}{},
		},
	},
	":person_frowning_tone2:": Code{
		"\U0001f64d\U0001f3fc",
		map[string]struct{}{
			"frown":                  struct{}{},
			"gesture":                struct{}{},
			"medium-light skin tone": struct{}{},
			"person frowning":        struct{}{},
		},
	},
	":person_frowning_tone3:": Code{
		"\U0001f64d\U0001f3fd",
		map[string]struct{}{
			"frown":            struct{}{},
			"gesture":          struct{}{},
			"medium skin tone": struct{}{},
			"person frowning":  struct{}{},
		},
	},
	":person_frowning_tone4:": Code{
		"\U0001f64d\U0001f3fe",
		map[string]struct{}{
			"frown":                 struct{}{},
			"gesture":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"person frowning":       struct{}{},
		},
	},
	":person_frowning_tone5:": Code{
		"\U0001f64d\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":  struct{}{},
			"frown":           struct{}{},
			"gesture":         struct{}{},
			"person frowning": struct{}{},
		},
	},
	":man_frowning:": Code{
		"\U0001f64d\u200d\u2642\ufe0f",
		map[string]struct{}{
			"frowning":     struct{}{},
			"gesture":      struct{}{},
			"man":          struct{}{},
			"man frowning": struct{}{},
		},
	},
	":man_frowning_tone1:": Code{
		"\U0001f64d\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"frowning":        struct{}{},
			"gesture":         struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man frowning":    struct{}{},
		},
	},
	":man_frowning_tone2:": Code{
		"\U0001f64d\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"frowning":               struct{}{},
			"gesture":                struct{}{},
			"man":                    struct{}{},
			"man frowning":           struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_frowning_tone3:": Code{
		"\U0001f64d\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"frowning":         struct{}{},
			"gesture":          struct{}{},
			"man":              struct{}{},
			"man frowning":     struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":man_frowning_tone4:": Code{
		"\U0001f64d\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"frowning":              struct{}{},
			"gesture":               struct{}{},
			"man":                   struct{}{},
			"man frowning":          struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_frowning_tone5:": Code{
		"\U0001f64d\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"frowning":       struct{}{},
			"gesture":        struct{}{},
			"man":            struct{}{},
			"man frowning":   struct{}{},
		},
	},
	":woman_frowning:": Code{
		"\U0001f64d\u200d\u2640\ufe0f",
		map[string]struct{}{
			"frowning":       struct{}{},
			"gesture":        struct{}{},
			"woman":          struct{}{},
			"woman frowning": struct{}{},
		},
	},
	":woman_frowning_tone1:": Code{
		"\U0001f64d\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"frowning":        struct{}{},
			"gesture":         struct{}{},
			"light skin tone": struct{}{},
			"woman":           struct{}{},
			"woman frowning":  struct{}{},
		},
	},
	":woman_frowning_tone2:": Code{
		"\U0001f64d\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"frowning":               struct{}{},
			"gesture":                struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
			"woman frowning":         struct{}{},
		},
	},
	":woman_frowning_tone3:": Code{
		"\U0001f64d\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"frowning":         struct{}{},
			"gesture":          struct{}{},
			"medium skin tone": struct{}{},
			"woman":            struct{}{},
			"woman frowning":   struct{}{},
		},
	},
	":woman_frowning_tone4:": Code{
		"\U0001f64d\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"frowning":              struct{}{},
			"gesture":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
			"woman frowning":        struct{}{},
		},
	},
	":woman_frowning_tone5:": Code{
		"\U0001f64d\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"frowning":       struct{}{},
			"gesture":        struct{}{},
			"woman":          struct{}{},
			"woman frowning": struct{}{},
		},
	},
	":person_pouting:": Code{
		"\U0001f64e",
		map[string]struct{}{
			"gesture":        struct{}{},
			"person pouting": struct{}{},
			"pouting":        struct{}{},
		},
	},
	":person_pouting_tone1:": Code{
		"\U0001f64e\U0001f3fb",
		map[string]struct{}{
			"gesture":         struct{}{},
			"light skin tone": struct{}{},
			"person pouting":  struct{}{},
			"pouting":         struct{}{},
		},
	},
	":person_pouting_tone2:": Code{
		"\U0001f64e\U0001f3fc",
		map[string]struct{}{
			"gesture":                struct{}{},
			"medium-light skin tone": struct{}{},
			"person pouting":         struct{}{},
			"pouting":                struct{}{},
		},
	},
	":person_pouting_tone3:": Code{
		"\U0001f64e\U0001f3fd",
		map[string]struct{}{
			"gesture":          struct{}{},
			"medium skin tone": struct{}{},
			"person pouting":   struct{}{},
			"pouting":          struct{}{},
		},
	},
	":person_pouting_tone4:": Code{
		"\U0001f64e\U0001f3fe",
		map[string]struct{}{
			"gesture":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"person pouting":        struct{}{},
			"pouting":               struct{}{},
		},
	},
	":person_pouting_tone5:": Code{
		"\U0001f64e\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"gesture":        struct{}{},
			"person pouting": struct{}{},
			"pouting":        struct{}{},
		},
	},
	":man_pouting:": Code{
		"\U0001f64e\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":     struct{}{},
			"man":         struct{}{},
			"man pouting": struct{}{},
			"pouting":     struct{}{},
		},
	},
	":man_pouting_tone1:": Code{
		"\U0001f64e\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":         struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man pouting":     struct{}{},
			"pouting":         struct{}{},
		},
	},
	":man_pouting_tone2:": Code{
		"\U0001f64e\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":                struct{}{},
			"man":                    struct{}{},
			"man pouting":            struct{}{},
			"medium-light skin tone": struct{}{},
			"pouting":                struct{}{},
		},
	},
	":man_pouting_tone3:": Code{
		"\U0001f64e\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":          struct{}{},
			"man":              struct{}{},
			"man pouting":      struct{}{},
			"medium skin tone": struct{}{},
			"pouting":          struct{}{},
		},
	},
	":man_pouting_tone4:": Code{
		"\U0001f64e\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":               struct{}{},
			"man":                   struct{}{},
			"man pouting":           struct{}{},
			"medium-dark skin tone": struct{}{},
			"pouting":               struct{}{},
		},
	},
	":man_pouting_tone5:": Code{
		"\U0001f64e\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"gesture":        struct{}{},
			"man":            struct{}{},
			"man pouting":    struct{}{},
			"pouting":        struct{}{},
		},
	},
	":woman_pouting:": Code{
		"\U0001f64e\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":       struct{}{},
			"pouting":       struct{}{},
			"woman":         struct{}{},
			"woman pouting": struct{}{},
		},
	},
	":woman_pouting_tone1:": Code{
		"\U0001f64e\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":         struct{}{},
			"light skin tone": struct{}{},
			"pouting":         struct{}{},
			"woman":           struct{}{},
			"woman pouting":   struct{}{},
		},
	},
	":woman_pouting_tone2:": Code{
		"\U0001f64e\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":                struct{}{},
			"medium-light skin tone": struct{}{},
			"pouting":                struct{}{},
			"woman":                  struct{}{},
			"woman pouting":          struct{}{},
		},
	},
	":woman_pouting_tone3:": Code{
		"\U0001f64e\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":          struct{}{},
			"medium skin tone": struct{}{},
			"pouting":          struct{}{},
			"woman":            struct{}{},
			"woman pouting":    struct{}{},
		},
	},
	":woman_pouting_tone4:": Code{
		"\U0001f64e\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"pouting":               struct{}{},
			"woman":                 struct{}{},
			"woman pouting":         struct{}{},
		},
	},
	":woman_pouting_tone5:": Code{
		"\U0001f64e\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"gesture":        struct{}{},
			"pouting":        struct{}{},
			"woman":          struct{}{},
			"woman pouting":  struct{}{},
		},
	},
	":person_gesturing_no:": Code{
		"\U0001f645",
		map[string]struct{}{
			"forbidden":           struct{}{},
			"gesture":             struct{}{},
			"hand":                struct{}{},
			"no":                  struct{}{},
			"not":                 struct{}{},
			"person gesturing NO": struct{}{},
			"prohibited":          struct{}{},
		},
	},
	":person_gesturing_no_tone1:": Code{
		"\U0001f645\U0001f3fb",
		map[string]struct{}{
			"forbidden":           struct{}{},
			"gesture":             struct{}{},
			"hand":                struct{}{},
			"light skin tone":     struct{}{},
			"no":                  struct{}{},
			"not":                 struct{}{},
			"person gesturing NO": struct{}{},
			"prohibited":          struct{}{},
		},
	},
	":person_gesturing_no_tone2:": Code{
		"\U0001f645\U0001f3fc",
		map[string]struct{}{
			"forbidden":              struct{}{},
			"gesture":                struct{}{},
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"no":                     struct{}{},
			"not":                    struct{}{},
			"person gesturing NO":    struct{}{},
			"prohibited":             struct{}{},
		},
	},
	":person_gesturing_no_tone3:": Code{
		"\U0001f645\U0001f3fd",
		map[string]struct{}{
			"forbidden":           struct{}{},
			"gesture":             struct{}{},
			"hand":                struct{}{},
			"medium skin tone":    struct{}{},
			"no":                  struct{}{},
			"not":                 struct{}{},
			"person gesturing NO": struct{}{},
			"prohibited":          struct{}{},
		},
	},
	":person_gesturing_no_tone4:": Code{
		"\U0001f645\U0001f3fe",
		map[string]struct{}{
			"forbidden":             struct{}{},
			"gesture":               struct{}{},
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"no":                    struct{}{},
			"not":                   struct{}{},
			"person gesturing NO":   struct{}{},
			"prohibited":            struct{}{},
		},
	},
	":person_gesturing_no_tone5:": Code{
		"\U0001f645\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":      struct{}{},
			"forbidden":           struct{}{},
			"gesture":             struct{}{},
			"hand":                struct{}{},
			"no":                  struct{}{},
			"not":                 struct{}{},
			"person gesturing NO": struct{}{},
			"prohibited":          struct{}{},
		},
	},
	":man_gesturing_no:": Code{
		"\U0001f645\u200d\u2642\ufe0f",
		map[string]struct{}{
			"forbidden":        struct{}{},
			"gesture":          struct{}{},
			"hand":             struct{}{},
			"man":              struct{}{},
			"man gesturing NO": struct{}{},
			"no":               struct{}{},
			"prohibited":       struct{}{},
		},
	},
	":man_gesturing_no_tone1:": Code{
		"\U0001f645\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"forbidden":        struct{}{},
			"gesture":          struct{}{},
			"hand":             struct{}{},
			"light skin tone":  struct{}{},
			"man":              struct{}{},
			"man gesturing NO": struct{}{},
			"no":               struct{}{},
			"prohibited":       struct{}{},
		},
	},
	":man_gesturing_no_tone2:": Code{
		"\U0001f645\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"forbidden":              struct{}{},
			"gesture":                struct{}{},
			"hand":                   struct{}{},
			"man":                    struct{}{},
			"man gesturing NO":       struct{}{},
			"medium-light skin tone": struct{}{},
			"no":                     struct{}{},
			"prohibited":             struct{}{},
		},
	},
	":man_gesturing_no_tone3:": Code{
		"\U0001f645\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"forbidden":        struct{}{},
			"gesture":          struct{}{},
			"hand":             struct{}{},
			"man":              struct{}{},
			"man gesturing NO": struct{}{},
			"medium skin tone": struct{}{},
			"no":               struct{}{},
			"prohibited":       struct{}{},
		},
	},
	":man_gesturing_no_tone4:": Code{
		"\U0001f645\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"forbidden":             struct{}{},
			"gesture":               struct{}{},
			"hand":                  struct{}{},
			"man":                   struct{}{},
			"man gesturing NO":      struct{}{},
			"medium-dark skin tone": struct{}{},
			"no":                    struct{}{},
			"prohibited":            struct{}{},
		},
	},
	":man_gesturing_no_tone5:": Code{
		"\U0001f645\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone":   struct{}{},
			"forbidden":        struct{}{},
			"gesture":          struct{}{},
			"hand":             struct{}{},
			"man":              struct{}{},
			"man gesturing NO": struct{}{},
			"no":               struct{}{},
			"prohibited":       struct{}{},
		},
	},
	":woman_gesturing_no:": Code{
		"\U0001f645\u200d\u2640\ufe0f",
		map[string]struct{}{
			"forbidden":          struct{}{},
			"gesture":            struct{}{},
			"hand":               struct{}{},
			"no":                 struct{}{},
			"prohibited":         struct{}{},
			"woman":              struct{}{},
			"woman gesturing NO": struct{}{},
		},
	},
	":woman_gesturing_no_tone1:": Code{
		"\U0001f645\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"forbidden":          struct{}{},
			"gesture":            struct{}{},
			"hand":               struct{}{},
			"light skin tone":    struct{}{},
			"no":                 struct{}{},
			"prohibited":         struct{}{},
			"woman":              struct{}{},
			"woman gesturing NO": struct{}{},
		},
	},
	":woman_gesturing_no_tone2:": Code{
		"\U0001f645\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"forbidden":              struct{}{},
			"gesture":                struct{}{},
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"no":                     struct{}{},
			"prohibited":             struct{}{},
			"woman":                  struct{}{},
			"woman gesturing NO":     struct{}{},
		},
	},
	":woman_gesturing_no_tone3:": Code{
		"\U0001f645\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"forbidden":          struct{}{},
			"gesture":            struct{}{},
			"hand":               struct{}{},
			"medium skin tone":   struct{}{},
			"no":                 struct{}{},
			"prohibited":         struct{}{},
			"woman":              struct{}{},
			"woman gesturing NO": struct{}{},
		},
	},
	":woman_gesturing_no_tone4:": Code{
		"\U0001f645\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"forbidden":             struct{}{},
			"gesture":               struct{}{},
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"no":                    struct{}{},
			"prohibited":            struct{}{},
			"woman":                 struct{}{},
			"woman gesturing NO":    struct{}{},
		},
	},
	":woman_gesturing_no_tone5:": Code{
		"\U0001f645\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":     struct{}{},
			"forbidden":          struct{}{},
			"gesture":            struct{}{},
			"hand":               struct{}{},
			"no":                 struct{}{},
			"prohibited":         struct{}{},
			"woman":              struct{}{},
			"woman gesturing NO": struct{}{},
		},
	},
	":person_gesturing_ok:": Code{
		"\U0001f646",
		map[string]struct{}{
			"gesture":             struct{}{},
			"hand":                struct{}{},
			"OK":                  struct{}{},
			"person gesturing OK": struct{}{},
		},
	},
	":person_gesturing_ok_tone1:": Code{
		"\U0001f646\U0001f3fb",
		map[string]struct{}{
			"gesture":             struct{}{},
			"hand":                struct{}{},
			"light skin tone":     struct{}{},
			"OK":                  struct{}{},
			"person gesturing OK": struct{}{},
		},
	},
	":person_gesturing_ok_tone2:": Code{
		"\U0001f646\U0001f3fc",
		map[string]struct{}{
			"gesture":                struct{}{},
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"OK":                     struct{}{},
			"person gesturing OK":    struct{}{},
		},
	},
	":person_gesturing_ok_tone3:": Code{
		"\U0001f646\U0001f3fd",
		map[string]struct{}{
			"gesture":             struct{}{},
			"hand":                struct{}{},
			"medium skin tone":    struct{}{},
			"OK":                  struct{}{},
			"person gesturing OK": struct{}{},
		},
	},
	":person_gesturing_ok_tone4:": Code{
		"\U0001f646\U0001f3fe",
		map[string]struct{}{
			"gesture":               struct{}{},
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"OK":                    struct{}{},
			"person gesturing OK":   struct{}{},
		},
	},
	":person_gesturing_ok_tone5:": Code{
		"\U0001f646\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":      struct{}{},
			"gesture":             struct{}{},
			"hand":                struct{}{},
			"OK":                  struct{}{},
			"person gesturing OK": struct{}{},
		},
	},
	":man_gesturing_ok:": Code{
		"\U0001f646\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":          struct{}{},
			"hand":             struct{}{},
			"man":              struct{}{},
			"man gesturing OK": struct{}{},
			"OK":               struct{}{},
		},
	},
	":man_gesturing_ok_tone1:": Code{
		"\U0001f646\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":          struct{}{},
			"hand":             struct{}{},
			"light skin tone":  struct{}{},
			"man":              struct{}{},
			"man gesturing OK": struct{}{},
			"OK":               struct{}{},
		},
	},
	":man_gesturing_ok_tone2:": Code{
		"\U0001f646\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":                struct{}{},
			"hand":                   struct{}{},
			"man":                    struct{}{},
			"man gesturing OK":       struct{}{},
			"medium-light skin tone": struct{}{},
			"OK":                     struct{}{},
		},
	},
	":man_gesturing_ok_tone3:": Code{
		"\U0001f646\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":          struct{}{},
			"hand":             struct{}{},
			"man":              struct{}{},
			"man gesturing OK": struct{}{},
			"medium skin tone": struct{}{},
			"OK":               struct{}{},
		},
	},
	":man_gesturing_ok_tone4:": Code{
		"\U0001f646\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":               struct{}{},
			"hand":                  struct{}{},
			"man":                   struct{}{},
			"man gesturing OK":      struct{}{},
			"medium-dark skin tone": struct{}{},
			"OK":                    struct{}{},
		},
	},
	":man_gesturing_ok_tone5:": Code{
		"\U0001f646\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone":   struct{}{},
			"gesture":          struct{}{},
			"hand":             struct{}{},
			"man":              struct{}{},
			"man gesturing OK": struct{}{},
			"OK":               struct{}{},
		},
	},
	":woman_gesturing_ok:": Code{
		"\U0001f646\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":            struct{}{},
			"hand":               struct{}{},
			"OK":                 struct{}{},
			"woman":              struct{}{},
			"woman gesturing OK": struct{}{},
		},
	},
	":woman_gesturing_ok_tone1:": Code{
		"\U0001f646\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":            struct{}{},
			"hand":               struct{}{},
			"light skin tone":    struct{}{},
			"OK":                 struct{}{},
			"woman":              struct{}{},
			"woman gesturing OK": struct{}{},
		},
	},
	":woman_gesturing_ok_tone2:": Code{
		"\U0001f646\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":                struct{}{},
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"OK":                     struct{}{},
			"woman":                  struct{}{},
			"woman gesturing OK":     struct{}{},
		},
	},
	":woman_gesturing_ok_tone3:": Code{
		"\U0001f646\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":            struct{}{},
			"hand":               struct{}{},
			"medium skin tone":   struct{}{},
			"OK":                 struct{}{},
			"woman":              struct{}{},
			"woman gesturing OK": struct{}{},
		},
	},
	":woman_gesturing_ok_tone4:": Code{
		"\U0001f646\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":               struct{}{},
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"OK":                    struct{}{},
			"woman":                 struct{}{},
			"woman gesturing OK":    struct{}{},
		},
	},
	":woman_gesturing_ok_tone5:": Code{
		"\U0001f646\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":     struct{}{},
			"gesture":            struct{}{},
			"hand":               struct{}{},
			"OK":                 struct{}{},
			"woman":              struct{}{},
			"woman gesturing OK": struct{}{},
		},
	},
	":person_tipping_hand:": Code{
		"\U0001f481",
		map[string]struct{}{
			"hand":                struct{}{},
			"help":                struct{}{},
			"information":         struct{}{},
			"person tipping hand": struct{}{},
			"sassy":               struct{}{},
			"tipping":             struct{}{},
		},
	},
	":person_tipping_hand_tone1:": Code{
		"\U0001f481\U0001f3fb",
		map[string]struct{}{
			"hand":                struct{}{},
			"help":                struct{}{},
			"information":         struct{}{},
			"light skin tone":     struct{}{},
			"person tipping hand": struct{}{},
			"sassy":               struct{}{},
			"tipping":             struct{}{},
		},
	},
	":person_tipping_hand_tone2:": Code{
		"\U0001f481\U0001f3fc",
		map[string]struct{}{
			"hand":                   struct{}{},
			"help":                   struct{}{},
			"information":            struct{}{},
			"medium-light skin tone": struct{}{},
			"person tipping hand":    struct{}{},
			"sassy":                  struct{}{},
			"tipping":                struct{}{},
		},
	},
	":person_tipping_hand_tone3:": Code{
		"\U0001f481\U0001f3fd",
		map[string]struct{}{
			"hand":                struct{}{},
			"help":                struct{}{},
			"information":         struct{}{},
			"medium skin tone":    struct{}{},
			"person tipping hand": struct{}{},
			"sassy":               struct{}{},
			"tipping":             struct{}{},
		},
	},
	":person_tipping_hand_tone4:": Code{
		"\U0001f481\U0001f3fe",
		map[string]struct{}{
			"hand":                  struct{}{},
			"help":                  struct{}{},
			"information":           struct{}{},
			"medium-dark skin tone": struct{}{},
			"person tipping hand":   struct{}{},
			"sassy":                 struct{}{},
			"tipping":               struct{}{},
		},
	},
	":person_tipping_hand_tone5:": Code{
		"\U0001f481\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":      struct{}{},
			"hand":                struct{}{},
			"help":                struct{}{},
			"information":         struct{}{},
			"person tipping hand": struct{}{},
			"sassy":               struct{}{},
			"tipping":             struct{}{},
		},
	},
	":man_tipping_hand:": Code{
		"\U0001f481\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":              struct{}{},
			"man tipping hand": struct{}{},
			"sassy":            struct{}{},
			"tipping hand":     struct{}{},
		},
	},
	":man_tipping_hand_tone1:": Code{
		"\U0001f481\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"light skin tone":  struct{}{},
			"man":              struct{}{},
			"man tipping hand": struct{}{},
			"sassy":            struct{}{},
			"tipping hand":     struct{}{},
		},
	},
	":man_tipping_hand_tone2:": Code{
		"\U0001f481\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                    struct{}{},
			"man tipping hand":       struct{}{},
			"medium-light skin tone": struct{}{},
			"sassy":                  struct{}{},
			"tipping hand":           struct{}{},
		},
	},
	":man_tipping_hand_tone3:": Code{
		"\U0001f481\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":              struct{}{},
			"man tipping hand": struct{}{},
			"medium skin tone": struct{}{},
			"sassy":            struct{}{},
			"tipping hand":     struct{}{},
		},
	},
	":man_tipping_hand_tone4:": Code{
		"\U0001f481\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                   struct{}{},
			"man tipping hand":      struct{}{},
			"medium-dark skin tone": struct{}{},
			"sassy":                 struct{}{},
			"tipping hand":          struct{}{},
		},
	},
	":man_tipping_hand_tone5:": Code{
		"\U0001f481\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone":   struct{}{},
			"man":              struct{}{},
			"man tipping hand": struct{}{},
			"sassy":            struct{}{},
			"tipping hand":     struct{}{},
		},
	},
	":woman_tipping_hand:": Code{
		"\U0001f481\u200d\u2640\ufe0f",
		map[string]struct{}{
			"sassy":              struct{}{},
			"tipping hand":       struct{}{},
			"woman":              struct{}{},
			"woman tipping hand": struct{}{},
		},
	},
	":woman_tipping_hand_tone1:": Code{
		"\U0001f481\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"light skin tone":    struct{}{},
			"sassy":              struct{}{},
			"tipping hand":       struct{}{},
			"woman":              struct{}{},
			"woman tipping hand": struct{}{},
		},
	},
	":woman_tipping_hand_tone2:": Code{
		"\U0001f481\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"sassy":                  struct{}{},
			"tipping hand":           struct{}{},
			"woman":                  struct{}{},
			"woman tipping hand":     struct{}{},
		},
	},
	":woman_tipping_hand_tone3:": Code{
		"\U0001f481\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium skin tone":   struct{}{},
			"sassy":              struct{}{},
			"tipping hand":       struct{}{},
			"woman":              struct{}{},
			"woman tipping hand": struct{}{},
		},
	},
	":woman_tipping_hand_tone4:": Code{
		"\U0001f481\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"sassy":                 struct{}{},
			"tipping hand":          struct{}{},
			"woman":                 struct{}{},
			"woman tipping hand":    struct{}{},
		},
	},
	":woman_tipping_hand_tone5:": Code{
		"\U0001f481\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":     struct{}{},
			"sassy":              struct{}{},
			"tipping hand":       struct{}{},
			"woman":              struct{}{},
			"woman tipping hand": struct{}{},
		},
	},
	":person_raising_hand:": Code{
		"\U0001f64b",
		map[string]struct{}{
			"gesture":             struct{}{},
			"hand":                struct{}{},
			"happy":               struct{}{},
			"person raising hand": struct{}{},
			"raised":              struct{}{},
		},
	},
	":person_raising_hand_tone1:": Code{
		"\U0001f64b\U0001f3fb",
		map[string]struct{}{
			"gesture":             struct{}{},
			"hand":                struct{}{},
			"happy":               struct{}{},
			"light skin tone":     struct{}{},
			"person raising hand": struct{}{},
			"raised":              struct{}{},
		},
	},
	":person_raising_hand_tone2:": Code{
		"\U0001f64b\U0001f3fc",
		map[string]struct{}{
			"gesture":                struct{}{},
			"hand":                   struct{}{},
			"happy":                  struct{}{},
			"medium-light skin tone": struct{}{},
			"person raising hand":    struct{}{},
			"raised":                 struct{}{},
		},
	},
	":person_raising_hand_tone3:": Code{
		"\U0001f64b\U0001f3fd",
		map[string]struct{}{
			"gesture":             struct{}{},
			"hand":                struct{}{},
			"happy":               struct{}{},
			"medium skin tone":    struct{}{},
			"person raising hand": struct{}{},
			"raised":              struct{}{},
		},
	},
	":person_raising_hand_tone4:": Code{
		"\U0001f64b\U0001f3fe",
		map[string]struct{}{
			"gesture":               struct{}{},
			"hand":                  struct{}{},
			"happy":                 struct{}{},
			"medium-dark skin tone": struct{}{},
			"person raising hand":   struct{}{},
			"raised":                struct{}{},
		},
	},
	":person_raising_hand_tone5:": Code{
		"\U0001f64b\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":      struct{}{},
			"gesture":             struct{}{},
			"hand":                struct{}{},
			"happy":               struct{}{},
			"person raising hand": struct{}{},
			"raised":              struct{}{},
		},
	},
	":man_raising_hand:": Code{
		"\U0001f64b\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":          struct{}{},
			"man":              struct{}{},
			"man raising hand": struct{}{},
			"raising hand":     struct{}{},
		},
	},
	":man_raising_hand_tone1:": Code{
		"\U0001f64b\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":          struct{}{},
			"light skin tone":  struct{}{},
			"man":              struct{}{},
			"man raising hand": struct{}{},
			"raising hand":     struct{}{},
		},
	},
	":man_raising_hand_tone2:": Code{
		"\U0001f64b\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":                struct{}{},
			"man":                    struct{}{},
			"man raising hand":       struct{}{},
			"medium-light skin tone": struct{}{},
			"raising hand":           struct{}{},
		},
	},
	":man_raising_hand_tone3:": Code{
		"\U0001f64b\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":          struct{}{},
			"man":              struct{}{},
			"man raising hand": struct{}{},
			"medium skin tone": struct{}{},
			"raising hand":     struct{}{},
		},
	},
	":man_raising_hand_tone4:": Code{
		"\U0001f64b\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gesture":               struct{}{},
			"man":                   struct{}{},
			"man raising hand":      struct{}{},
			"medium-dark skin tone": struct{}{},
			"raising hand":          struct{}{},
		},
	},
	":man_raising_hand_tone5:": Code{
		"\U0001f64b\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone":   struct{}{},
			"gesture":          struct{}{},
			"man":              struct{}{},
			"man raising hand": struct{}{},
			"raising hand":     struct{}{},
		},
	},
	":woman_raising_hand:": Code{
		"\U0001f64b\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":            struct{}{},
			"raising hand":       struct{}{},
			"woman":              struct{}{},
			"woman raising hand": struct{}{},
		},
	},
	":woman_raising_hand_tone1:": Code{
		"\U0001f64b\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":            struct{}{},
			"light skin tone":    struct{}{},
			"raising hand":       struct{}{},
			"woman":              struct{}{},
			"woman raising hand": struct{}{},
		},
	},
	":woman_raising_hand_tone2:": Code{
		"\U0001f64b\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":                struct{}{},
			"medium-light skin tone": struct{}{},
			"raising hand":           struct{}{},
			"woman":                  struct{}{},
			"woman raising hand":     struct{}{},
		},
	},
	":woman_raising_hand_tone3:": Code{
		"\U0001f64b\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":            struct{}{},
			"medium skin tone":   struct{}{},
			"raising hand":       struct{}{},
			"woman":              struct{}{},
			"woman raising hand": struct{}{},
		},
	},
	":woman_raising_hand_tone4:": Code{
		"\U0001f64b\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gesture":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"raising hand":          struct{}{},
			"woman":                 struct{}{},
			"woman raising hand":    struct{}{},
		},
	},
	":woman_raising_hand_tone5:": Code{
		"\U0001f64b\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":     struct{}{},
			"gesture":            struct{}{},
			"raising hand":       struct{}{},
			"woman":              struct{}{},
			"woman raising hand": struct{}{},
		},
	},
	":person_bowing:": Code{
		"\U0001f647",
		map[string]struct{}{
			"apology":       struct{}{},
			"bow":           struct{}{},
			"gesture":       struct{}{},
			"person bowing": struct{}{},
			"sorry":         struct{}{},
		},
	},
	":person_bowing_tone1:": Code{
		"\U0001f647\U0001f3fb",
		map[string]struct{}{
			"apology":         struct{}{},
			"bow":             struct{}{},
			"gesture":         struct{}{},
			"light skin tone": struct{}{},
			"person bowing":   struct{}{},
			"sorry":           struct{}{},
		},
	},
	":person_bowing_tone2:": Code{
		"\U0001f647\U0001f3fc",
		map[string]struct{}{
			"apology":                struct{}{},
			"bow":                    struct{}{},
			"gesture":                struct{}{},
			"medium-light skin tone": struct{}{},
			"person bowing":          struct{}{},
			"sorry":                  struct{}{},
		},
	},
	":person_bowing_tone3:": Code{
		"\U0001f647\U0001f3fd",
		map[string]struct{}{
			"apology":          struct{}{},
			"bow":              struct{}{},
			"gesture":          struct{}{},
			"medium skin tone": struct{}{},
			"person bowing":    struct{}{},
			"sorry":            struct{}{},
		},
	},
	":person_bowing_tone4:": Code{
		"\U0001f647\U0001f3fe",
		map[string]struct{}{
			"apology":               struct{}{},
			"bow":                   struct{}{},
			"gesture":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"person bowing":         struct{}{},
			"sorry":                 struct{}{},
		},
	},
	":person_bowing_tone5:": Code{
		"\U0001f647\U0001f3ff",
		map[string]struct{}{
			"apology":        struct{}{},
			"bow":            struct{}{},
			"dark skin tone": struct{}{},
			"gesture":        struct{}{},
			"person bowing":  struct{}{},
			"sorry":          struct{}{},
		},
	},
	":man_bowing:": Code{
		"\U0001f647\u200d\u2642\ufe0f",
		map[string]struct{}{
			"apology":    struct{}{},
			"bowing":     struct{}{},
			"favor":      struct{}{},
			"gesture":    struct{}{},
			"man":        struct{}{},
			"man bowing": struct{}{},
			"sorry":      struct{}{},
		},
	},
	":man_bowing_tone1:": Code{
		"\U0001f647\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"apology":         struct{}{},
			"bowing":          struct{}{},
			"favor":           struct{}{},
			"gesture":         struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man bowing":      struct{}{},
			"sorry":           struct{}{},
		},
	},
	":man_bowing_tone2:": Code{
		"\U0001f647\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"apology":                struct{}{},
			"bowing":                 struct{}{},
			"favor":                  struct{}{},
			"gesture":                struct{}{},
			"man":                    struct{}{},
			"man bowing":             struct{}{},
			"medium-light skin tone": struct{}{},
			"sorry":                  struct{}{},
		},
	},
	":man_bowing_tone3:": Code{
		"\U0001f647\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"apology":          struct{}{},
			"bowing":           struct{}{},
			"favor":            struct{}{},
			"gesture":          struct{}{},
			"man":              struct{}{},
			"man bowing":       struct{}{},
			"medium skin tone": struct{}{},
			"sorry":            struct{}{},
		},
	},
	":man_bowing_tone4:": Code{
		"\U0001f647\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"apology":               struct{}{},
			"bowing":                struct{}{},
			"favor":                 struct{}{},
			"gesture":               struct{}{},
			"man":                   struct{}{},
			"man bowing":            struct{}{},
			"medium-dark skin tone": struct{}{},
			"sorry":                 struct{}{},
		},
	},
	":man_bowing_tone5:": Code{
		"\U0001f647\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"apology":        struct{}{},
			"bowing":         struct{}{},
			"dark skin tone": struct{}{},
			"favor":          struct{}{},
			"gesture":        struct{}{},
			"man":            struct{}{},
			"man bowing":     struct{}{},
			"sorry":          struct{}{},
		},
	},
	":woman_bowing:": Code{
		"\U0001f647\u200d\u2640\ufe0f",
		map[string]struct{}{
			"apology":      struct{}{},
			"bowing":       struct{}{},
			"favor":        struct{}{},
			"gesture":      struct{}{},
			"sorry":        struct{}{},
			"woman":        struct{}{},
			"woman bowing": struct{}{},
		},
	},
	":woman_bowing_tone1:": Code{
		"\U0001f647\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"apology":         struct{}{},
			"bowing":          struct{}{},
			"favor":           struct{}{},
			"gesture":         struct{}{},
			"light skin tone": struct{}{},
			"sorry":           struct{}{},
			"woman":           struct{}{},
			"woman bowing":    struct{}{},
		},
	},
	":woman_bowing_tone2:": Code{
		"\U0001f647\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"apology":                struct{}{},
			"bowing":                 struct{}{},
			"favor":                  struct{}{},
			"gesture":                struct{}{},
			"medium-light skin tone": struct{}{},
			"sorry":                  struct{}{},
			"woman":                  struct{}{},
			"woman bowing":           struct{}{},
		},
	},
	":woman_bowing_tone3:": Code{
		"\U0001f647\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"apology":          struct{}{},
			"bowing":           struct{}{},
			"favor":            struct{}{},
			"gesture":          struct{}{},
			"medium skin tone": struct{}{},
			"sorry":            struct{}{},
			"woman":            struct{}{},
			"woman bowing":     struct{}{},
		},
	},
	":woman_bowing_tone4:": Code{
		"\U0001f647\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"apology":               struct{}{},
			"bowing":                struct{}{},
			"favor":                 struct{}{},
			"gesture":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"sorry":                 struct{}{},
			"woman":                 struct{}{},
			"woman bowing":          struct{}{},
		},
	},
	":woman_bowing_tone5:": Code{
		"\U0001f647\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"apology":        struct{}{},
			"bowing":         struct{}{},
			"dark skin tone": struct{}{},
			"favor":          struct{}{},
			"gesture":        struct{}{},
			"sorry":          struct{}{},
			"woman":          struct{}{},
			"woman bowing":   struct{}{},
		},
	},
	":person_facepalming:": Code{
		"\U0001f926",
		map[string]struct{}{
			"disbelief":          struct{}{},
			"exasperation":       struct{}{},
			"face":               struct{}{},
			"palm":               struct{}{},
			"person facepalming": struct{}{},
		},
	},
	":person_facepalming_tone1:": Code{
		"\U0001f926\U0001f3fb",
		map[string]struct{}{
			"disbelief":          struct{}{},
			"exasperation":       struct{}{},
			"face":               struct{}{},
			"light skin tone":    struct{}{},
			"palm":               struct{}{},
			"person facepalming": struct{}{},
		},
	},
	":person_facepalming_tone2:": Code{
		"\U0001f926\U0001f3fc",
		map[string]struct{}{
			"disbelief":              struct{}{},
			"exasperation":           struct{}{},
			"face":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"palm":                   struct{}{},
			"person facepalming":     struct{}{},
		},
	},
	":person_facepalming_tone3:": Code{
		"\U0001f926\U0001f3fd",
		map[string]struct{}{
			"disbelief":          struct{}{},
			"exasperation":       struct{}{},
			"face":               struct{}{},
			"medium skin tone":   struct{}{},
			"palm":               struct{}{},
			"person facepalming": struct{}{},
		},
	},
	":person_facepalming_tone4:": Code{
		"\U0001f926\U0001f3fe",
		map[string]struct{}{
			"disbelief":             struct{}{},
			"exasperation":          struct{}{},
			"face":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"palm":                  struct{}{},
			"person facepalming":    struct{}{},
		},
	},
	":person_facepalming_tone5:": Code{
		"\U0001f926\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":     struct{}{},
			"disbelief":          struct{}{},
			"exasperation":       struct{}{},
			"face":               struct{}{},
			"palm":               struct{}{},
			"person facepalming": struct{}{},
		},
	},
	":man_facepalming:": Code{
		"\U0001f926\u200d\u2642\ufe0f",
		map[string]struct{}{
			"disbelief":       struct{}{},
			"exasperation":    struct{}{},
			"facepalm":        struct{}{},
			"man":             struct{}{},
			"man facepalming": struct{}{},
		},
	},
	":man_facepalming_tone1:": Code{
		"\U0001f926\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"disbelief":       struct{}{},
			"exasperation":    struct{}{},
			"facepalm":        struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man facepalming": struct{}{},
		},
	},
	":man_facepalming_tone2:": Code{
		"\U0001f926\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"disbelief":              struct{}{},
			"exasperation":           struct{}{},
			"facepalm":               struct{}{},
			"man":                    struct{}{},
			"man facepalming":        struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_facepalming_tone3:": Code{
		"\U0001f926\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"disbelief":        struct{}{},
			"exasperation":     struct{}{},
			"facepalm":         struct{}{},
			"man":              struct{}{},
			"man facepalming":  struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":man_facepalming_tone4:": Code{
		"\U0001f926\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"disbelief":             struct{}{},
			"exasperation":          struct{}{},
			"facepalm":              struct{}{},
			"man":                   struct{}{},
			"man facepalming":       struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_facepalming_tone5:": Code{
		"\U0001f926\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone":  struct{}{},
			"disbelief":       struct{}{},
			"exasperation":    struct{}{},
			"facepalm":        struct{}{},
			"man":             struct{}{},
			"man facepalming": struct{}{},
		},
	},
	":woman_facepalming:": Code{
		"\U0001f926\u200d\u2640\ufe0f",
		map[string]struct{}{
			"disbelief":         struct{}{},
			"exasperation":      struct{}{},
			"facepalm":          struct{}{},
			"woman":             struct{}{},
			"woman facepalming": struct{}{},
		},
	},
	":woman_facepalming_tone1:": Code{
		"\U0001f926\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"disbelief":         struct{}{},
			"exasperation":      struct{}{},
			"facepalm":          struct{}{},
			"light skin tone":   struct{}{},
			"woman":             struct{}{},
			"woman facepalming": struct{}{},
		},
	},
	":woman_facepalming_tone2:": Code{
		"\U0001f926\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"disbelief":              struct{}{},
			"exasperation":           struct{}{},
			"facepalm":               struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
			"woman facepalming":      struct{}{},
		},
	},
	":woman_facepalming_tone3:": Code{
		"\U0001f926\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"disbelief":         struct{}{},
			"exasperation":      struct{}{},
			"facepalm":          struct{}{},
			"medium skin tone":  struct{}{},
			"woman":             struct{}{},
			"woman facepalming": struct{}{},
		},
	},
	":woman_facepalming_tone4:": Code{
		"\U0001f926\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"disbelief":             struct{}{},
			"exasperation":          struct{}{},
			"facepalm":              struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
			"woman facepalming":     struct{}{},
		},
	},
	":woman_facepalming_tone5:": Code{
		"\U0001f926\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":    struct{}{},
			"disbelief":         struct{}{},
			"exasperation":      struct{}{},
			"facepalm":          struct{}{},
			"woman":             struct{}{},
			"woman facepalming": struct{}{},
		},
	},
	":person_shrugging:": Code{
		"\U0001f937",
		map[string]struct{}{
			"doubt":            struct{}{},
			"ignorance":        struct{}{},
			"indifference":     struct{}{},
			"person shrugging": struct{}{},
			"shrug":            struct{}{},
		},
	},
	":person_shrugging_tone1:": Code{
		"\U0001f937\U0001f3fb",
		map[string]struct{}{
			"doubt":            struct{}{},
			"ignorance":        struct{}{},
			"indifference":     struct{}{},
			"light skin tone":  struct{}{},
			"person shrugging": struct{}{},
			"shrug":            struct{}{},
		},
	},
	":person_shrugging_tone2:": Code{
		"\U0001f937\U0001f3fc",
		map[string]struct{}{
			"doubt":                  struct{}{},
			"ignorance":              struct{}{},
			"indifference":           struct{}{},
			"medium-light skin tone": struct{}{},
			"person shrugging":       struct{}{},
			"shrug":                  struct{}{},
		},
	},
	":person_shrugging_tone3:": Code{
		"\U0001f937\U0001f3fd",
		map[string]struct{}{
			"doubt":            struct{}{},
			"ignorance":        struct{}{},
			"indifference":     struct{}{},
			"medium skin tone": struct{}{},
			"person shrugging": struct{}{},
			"shrug":            struct{}{},
		},
	},
	":person_shrugging_tone4:": Code{
		"\U0001f937\U0001f3fe",
		map[string]struct{}{
			"doubt":                 struct{}{},
			"ignorance":             struct{}{},
			"indifference":          struct{}{},
			"medium-dark skin tone": struct{}{},
			"person shrugging":      struct{}{},
			"shrug":                 struct{}{},
		},
	},
	":person_shrugging_tone5:": Code{
		"\U0001f937\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":   struct{}{},
			"doubt":            struct{}{},
			"ignorance":        struct{}{},
			"indifference":     struct{}{},
			"person shrugging": struct{}{},
			"shrug":            struct{}{},
		},
	},
	":man_shrugging:": Code{
		"\U0001f937\u200d\u2642\ufe0f",
		map[string]struct{}{
			"doubt":         struct{}{},
			"ignorance":     struct{}{},
			"indifference":  struct{}{},
			"man":           struct{}{},
			"man shrugging": struct{}{},
			"shrug":         struct{}{},
		},
	},
	":man_shrugging_tone1:": Code{
		"\U0001f937\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"doubt":           struct{}{},
			"ignorance":       struct{}{},
			"indifference":    struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man shrugging":   struct{}{},
			"shrug":           struct{}{},
		},
	},
	":man_shrugging_tone2:": Code{
		"\U0001f937\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"doubt":                  struct{}{},
			"ignorance":              struct{}{},
			"indifference":           struct{}{},
			"man":                    struct{}{},
			"man shrugging":          struct{}{},
			"medium-light skin tone": struct{}{},
			"shrug":                  struct{}{},
		},
	},
	":man_shrugging_tone3:": Code{
		"\U0001f937\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"doubt":            struct{}{},
			"ignorance":        struct{}{},
			"indifference":     struct{}{},
			"man":              struct{}{},
			"man shrugging":    struct{}{},
			"medium skin tone": struct{}{},
			"shrug":            struct{}{},
		},
	},
	":man_shrugging_tone4:": Code{
		"\U0001f937\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"doubt":                 struct{}{},
			"ignorance":             struct{}{},
			"indifference":          struct{}{},
			"man":                   struct{}{},
			"man shrugging":         struct{}{},
			"medium-dark skin tone": struct{}{},
			"shrug":                 struct{}{},
		},
	},
	":man_shrugging_tone5:": Code{
		"\U0001f937\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"doubt":          struct{}{},
			"ignorance":      struct{}{},
			"indifference":   struct{}{},
			"man":            struct{}{},
			"man shrugging":  struct{}{},
			"shrug":          struct{}{},
		},
	},
	":woman_shrugging:": Code{
		"\U0001f937\u200d\u2640\ufe0f",
		map[string]struct{}{
			"doubt":           struct{}{},
			"ignorance":       struct{}{},
			"indifference":    struct{}{},
			"shrug":           struct{}{},
			"woman":           struct{}{},
			"woman shrugging": struct{}{},
		},
	},
	":woman_shrugging_tone1:": Code{
		"\U0001f937\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"doubt":           struct{}{},
			"ignorance":       struct{}{},
			"indifference":    struct{}{},
			"light skin tone": struct{}{},
			"shrug":           struct{}{},
			"woman":           struct{}{},
			"woman shrugging": struct{}{},
		},
	},
	":woman_shrugging_tone2:": Code{
		"\U0001f937\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"doubt":                  struct{}{},
			"ignorance":              struct{}{},
			"indifference":           struct{}{},
			"medium-light skin tone": struct{}{},
			"shrug":                  struct{}{},
			"woman":                  struct{}{},
			"woman shrugging":        struct{}{},
		},
	},
	":woman_shrugging_tone3:": Code{
		"\U0001f937\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"doubt":            struct{}{},
			"ignorance":        struct{}{},
			"indifference":     struct{}{},
			"medium skin tone": struct{}{},
			"shrug":            struct{}{},
			"woman":            struct{}{},
			"woman shrugging":  struct{}{},
		},
	},
	":woman_shrugging_tone4:": Code{
		"\U0001f937\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"doubt":                 struct{}{},
			"ignorance":             struct{}{},
			"indifference":          struct{}{},
			"medium-dark skin tone": struct{}{},
			"shrug":                 struct{}{},
			"woman":                 struct{}{},
			"woman shrugging":       struct{}{},
		},
	},
	":woman_shrugging_tone5:": Code{
		"\U0001f937\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":  struct{}{},
			"doubt":           struct{}{},
			"ignorance":       struct{}{},
			"indifference":    struct{}{},
			"shrug":           struct{}{},
			"woman":           struct{}{},
			"woman shrugging": struct{}{},
		},
	},
	":person_getting_massage:": Code{
		"\U0001f486",
		map[string]struct{}{
			"face":                   struct{}{},
			"massage":                struct{}{},
			"person getting massage": struct{}{},
			"salon":                  struct{}{},
		},
	},
	":person_getting_massage_tone1:": Code{
		"\U0001f486\U0001f3fb",
		map[string]struct{}{
			"face":                   struct{}{},
			"light skin tone":        struct{}{},
			"massage":                struct{}{},
			"person getting massage": struct{}{},
			"salon":                  struct{}{},
		},
	},
	":person_getting_massage_tone2:": Code{
		"\U0001f486\U0001f3fc",
		map[string]struct{}{
			"face":                   struct{}{},
			"massage":                struct{}{},
			"medium-light skin tone": struct{}{},
			"person getting massage": struct{}{},
			"salon":                  struct{}{},
		},
	},
	":person_getting_massage_tone3:": Code{
		"\U0001f486\U0001f3fd",
		map[string]struct{}{
			"face":                   struct{}{},
			"massage":                struct{}{},
			"medium skin tone":       struct{}{},
			"person getting massage": struct{}{},
			"salon":                  struct{}{},
		},
	},
	":person_getting_massage_tone4:": Code{
		"\U0001f486\U0001f3fe",
		map[string]struct{}{
			"face":                   struct{}{},
			"massage":                struct{}{},
			"medium-dark skin tone":  struct{}{},
			"person getting massage": struct{}{},
			"salon":                  struct{}{},
		},
	},
	":person_getting_massage_tone5:": Code{
		"\U0001f486\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":         struct{}{},
			"face":                   struct{}{},
			"massage":                struct{}{},
			"person getting massage": struct{}{},
			"salon":                  struct{}{},
		},
	},
	":man_getting_face_massage:": Code{
		"\U0001f486\u200d\u2642\ufe0f",
		map[string]struct{}{
			"face":                struct{}{},
			"man":                 struct{}{},
			"man getting massage": struct{}{},
			"massage":             struct{}{},
		},
	},
	":man_getting_face_massage_tone1:": Code{
		"\U0001f486\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"face":                struct{}{},
			"light skin tone":     struct{}{},
			"man":                 struct{}{},
			"man getting massage": struct{}{},
			"massage":             struct{}{},
		},
	},
	":man_getting_face_massage_tone2:": Code{
		"\U0001f486\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"face":                   struct{}{},
			"man":                    struct{}{},
			"man getting massage":    struct{}{},
			"massage":                struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_getting_face_massage_tone3:": Code{
		"\U0001f486\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"face":                struct{}{},
			"man":                 struct{}{},
			"man getting massage": struct{}{},
			"massage":             struct{}{},
			"medium skin tone":    struct{}{},
		},
	},
	":man_getting_face_massage_tone4:": Code{
		"\U0001f486\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"face":                  struct{}{},
			"man":                   struct{}{},
			"man getting massage":   struct{}{},
			"massage":               struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_getting_face_massage_tone5:": Code{
		"\U0001f486\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone":      struct{}{},
			"face":                struct{}{},
			"man":                 struct{}{},
			"man getting massage": struct{}{},
			"massage":             struct{}{},
		},
	},
	":woman_getting_face_massage:": Code{
		"\U0001f486\u200d\u2640\ufe0f",
		map[string]struct{}{
			"face":                  struct{}{},
			"massage":               struct{}{},
			"woman":                 struct{}{},
			"woman getting massage": struct{}{},
		},
	},
	":woman_getting_face_massage_tone1:": Code{
		"\U0001f486\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"face":                  struct{}{},
			"light skin tone":       struct{}{},
			"massage":               struct{}{},
			"woman":                 struct{}{},
			"woman getting massage": struct{}{},
		},
	},
	":woman_getting_face_massage_tone2:": Code{
		"\U0001f486\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"face":                   struct{}{},
			"massage":                struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
			"woman getting massage":  struct{}{},
		},
	},
	":woman_getting_face_massage_tone3:": Code{
		"\U0001f486\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"face":                  struct{}{},
			"massage":               struct{}{},
			"medium skin tone":      struct{}{},
			"woman":                 struct{}{},
			"woman getting massage": struct{}{},
		},
	},
	":woman_getting_face_massage_tone4:": Code{
		"\U0001f486\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"face":                  struct{}{},
			"massage":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
			"woman getting massage": struct{}{},
		},
	},
	":woman_getting_face_massage_tone5:": Code{
		"\U0001f486\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":        struct{}{},
			"face":                  struct{}{},
			"massage":               struct{}{},
			"woman":                 struct{}{},
			"woman getting massage": struct{}{},
		},
	},
	":person_getting_haircut:": Code{
		"\U0001f487",
		map[string]struct{}{
			"barber":                 struct{}{},
			"beauty":                 struct{}{},
			"haircut":                struct{}{},
			"parlor":                 struct{}{},
			"person getting haircut": struct{}{},
		},
	},
	":person_getting_haircut_tone1:": Code{
		"\U0001f487\U0001f3fb",
		map[string]struct{}{
			"barber":                 struct{}{},
			"beauty":                 struct{}{},
			"haircut":                struct{}{},
			"light skin tone":        struct{}{},
			"parlor":                 struct{}{},
			"person getting haircut": struct{}{},
		},
	},
	":person_getting_haircut_tone2:": Code{
		"\U0001f487\U0001f3fc",
		map[string]struct{}{
			"barber":                 struct{}{},
			"beauty":                 struct{}{},
			"haircut":                struct{}{},
			"medium-light skin tone": struct{}{},
			"parlor":                 struct{}{},
			"person getting haircut": struct{}{},
		},
	},
	":person_getting_haircut_tone3:": Code{
		"\U0001f487\U0001f3fd",
		map[string]struct{}{
			"barber":                 struct{}{},
			"beauty":                 struct{}{},
			"haircut":                struct{}{},
			"medium skin tone":       struct{}{},
			"parlor":                 struct{}{},
			"person getting haircut": struct{}{},
		},
	},
	":person_getting_haircut_tone4:": Code{
		"\U0001f487\U0001f3fe",
		map[string]struct{}{
			"barber":                 struct{}{},
			"beauty":                 struct{}{},
			"haircut":                struct{}{},
			"medium-dark skin tone":  struct{}{},
			"parlor":                 struct{}{},
			"person getting haircut": struct{}{},
		},
	},
	":person_getting_haircut_tone5:": Code{
		"\U0001f487\U0001f3ff",
		map[string]struct{}{
			"barber":                 struct{}{},
			"beauty":                 struct{}{},
			"dark skin tone":         struct{}{},
			"haircut":                struct{}{},
			"parlor":                 struct{}{},
			"person getting haircut": struct{}{},
		},
	},
	":man_getting_haircut:": Code{
		"\U0001f487\u200d\u2642\ufe0f",
		map[string]struct{}{
			"haircut":             struct{}{},
			"man":                 struct{}{},
			"man getting haircut": struct{}{},
		},
	},
	":man_getting_haircut_tone1:": Code{
		"\U0001f487\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"haircut":             struct{}{},
			"light skin tone":     struct{}{},
			"man":                 struct{}{},
			"man getting haircut": struct{}{},
		},
	},
	":man_getting_haircut_tone2:": Code{
		"\U0001f487\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"haircut":                struct{}{},
			"man":                    struct{}{},
			"man getting haircut":    struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_getting_haircut_tone3:": Code{
		"\U0001f487\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"haircut":             struct{}{},
			"man":                 struct{}{},
			"man getting haircut": struct{}{},
			"medium skin tone":    struct{}{},
		},
	},
	":man_getting_haircut_tone4:": Code{
		"\U0001f487\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"haircut":               struct{}{},
			"man":                   struct{}{},
			"man getting haircut":   struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_getting_haircut_tone5:": Code{
		"\U0001f487\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone":      struct{}{},
			"haircut":             struct{}{},
			"man":                 struct{}{},
			"man getting haircut": struct{}{},
		},
	},
	":woman_getting_haircut:": Code{
		"\U0001f487\u200d\u2640\ufe0f",
		map[string]struct{}{
			"haircut":               struct{}{},
			"woman":                 struct{}{},
			"woman getting haircut": struct{}{},
		},
	},
	":woman_getting_haircut_tone1:": Code{
		"\U0001f487\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"haircut":               struct{}{},
			"light skin tone":       struct{}{},
			"woman":                 struct{}{},
			"woman getting haircut": struct{}{},
		},
	},
	":woman_getting_haircut_tone2:": Code{
		"\U0001f487\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"haircut":                struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
			"woman getting haircut":  struct{}{},
		},
	},
	":woman_getting_haircut_tone3:": Code{
		"\U0001f487\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"haircut":               struct{}{},
			"medium skin tone":      struct{}{},
			"woman":                 struct{}{},
			"woman getting haircut": struct{}{},
		},
	},
	":woman_getting_haircut_tone4:": Code{
		"\U0001f487\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"haircut":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
			"woman getting haircut": struct{}{},
		},
	},
	":woman_getting_haircut_tone5:": Code{
		"\U0001f487\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":        struct{}{},
			"haircut":               struct{}{},
			"woman":                 struct{}{},
			"woman getting haircut": struct{}{},
		},
	},
	":person_walking:": Code{
		"\U0001f6b6",
		map[string]struct{}{
			"hike":           struct{}{},
			"person walking": struct{}{},
			"walk":           struct{}{},
			"walking":        struct{}{},
		},
	},
	":person_walking_tone1:": Code{
		"\U0001f6b6\U0001f3fb",
		map[string]struct{}{
			"hike":            struct{}{},
			"light skin tone": struct{}{},
			"person walking":  struct{}{},
			"walk":            struct{}{},
			"walking":         struct{}{},
		},
	},
	":person_walking_tone2:": Code{
		"\U0001f6b6\U0001f3fc",
		map[string]struct{}{
			"hike":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"person walking":         struct{}{},
			"walk":                   struct{}{},
			"walking":                struct{}{},
		},
	},
	":person_walking_tone3:": Code{
		"\U0001f6b6\U0001f3fd",
		map[string]struct{}{
			"hike":             struct{}{},
			"medium skin tone": struct{}{},
			"person walking":   struct{}{},
			"walk":             struct{}{},
			"walking":          struct{}{},
		},
	},
	":person_walking_tone4:": Code{
		"\U0001f6b6\U0001f3fe",
		map[string]struct{}{
			"hike":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"person walking":        struct{}{},
			"walk":                  struct{}{},
			"walking":               struct{}{},
		},
	},
	":person_walking_tone5:": Code{
		"\U0001f6b6\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"hike":           struct{}{},
			"person walking": struct{}{},
			"walk":           struct{}{},
			"walking":        struct{}{},
		},
	},
	":man_walking:": Code{
		"\U0001f6b6\u200d\u2642\ufe0f",
		map[string]struct{}{
			"hike":        struct{}{},
			"man":         struct{}{},
			"man walking": struct{}{},
			"walk":        struct{}{},
		},
	},
	":man_walking_tone1:": Code{
		"\U0001f6b6\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"hike":            struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man walking":     struct{}{},
			"walk":            struct{}{},
		},
	},
	":man_walking_tone2:": Code{
		"\U0001f6b6\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"hike":                   struct{}{},
			"man":                    struct{}{},
			"man walking":            struct{}{},
			"medium-light skin tone": struct{}{},
			"walk":                   struct{}{},
		},
	},
	":man_walking_tone3:": Code{
		"\U0001f6b6\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"hike":             struct{}{},
			"man":              struct{}{},
			"man walking":      struct{}{},
			"medium skin tone": struct{}{},
			"walk":             struct{}{},
		},
	},
	":man_walking_tone4:": Code{
		"\U0001f6b6\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"hike":                  struct{}{},
			"man":                   struct{}{},
			"man walking":           struct{}{},
			"medium-dark skin tone": struct{}{},
			"walk":                  struct{}{},
		},
	},
	":man_walking_tone5:": Code{
		"\U0001f6b6\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"hike":           struct{}{},
			"man":            struct{}{},
			"man walking":    struct{}{},
			"walk":           struct{}{},
		},
	},
	":woman_walking:": Code{
		"\U0001f6b6\u200d\u2640\ufe0f",
		map[string]struct{}{
			"hike":          struct{}{},
			"walk":          struct{}{},
			"woman":         struct{}{},
			"woman walking": struct{}{},
		},
	},
	":woman_walking_tone1:": Code{
		"\U0001f6b6\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"hike":            struct{}{},
			"light skin tone": struct{}{},
			"walk":            struct{}{},
			"woman":           struct{}{},
			"woman walking":   struct{}{},
		},
	},
	":woman_walking_tone2:": Code{
		"\U0001f6b6\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"hike":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"walk":                   struct{}{},
			"woman":                  struct{}{},
			"woman walking":          struct{}{},
		},
	},
	":woman_walking_tone3:": Code{
		"\U0001f6b6\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"hike":             struct{}{},
			"medium skin tone": struct{}{},
			"walk":             struct{}{},
			"woman":            struct{}{},
			"woman walking":    struct{}{},
		},
	},
	":woman_walking_tone4:": Code{
		"\U0001f6b6\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"hike":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"walk":                  struct{}{},
			"woman":                 struct{}{},
			"woman walking":         struct{}{},
		},
	},
	":woman_walking_tone5:": Code{
		"\U0001f6b6\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"hike":           struct{}{},
			"walk":           struct{}{},
			"woman":          struct{}{},
			"woman walking":  struct{}{},
		},
	},
	":person_running:": Code{
		"\U0001f3c3",
		map[string]struct{}{
			"marathon":       struct{}{},
			"person running": struct{}{},
			"running":        struct{}{},
		},
	},
	":person_running_tone1:": Code{
		"\U0001f3c3\U0001f3fb",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"marathon":        struct{}{},
			"person running":  struct{}{},
			"running":         struct{}{},
		},
	},
	":person_running_tone2:": Code{
		"\U0001f3c3\U0001f3fc",
		map[string]struct{}{
			"marathon":               struct{}{},
			"medium-light skin tone": struct{}{},
			"person running":         struct{}{},
			"running":                struct{}{},
		},
	},
	":person_running_tone3:": Code{
		"\U0001f3c3\U0001f3fd",
		map[string]struct{}{
			"marathon":         struct{}{},
			"medium skin tone": struct{}{},
			"person running":   struct{}{},
			"running":          struct{}{},
		},
	},
	":person_running_tone4:": Code{
		"\U0001f3c3\U0001f3fe",
		map[string]struct{}{
			"marathon":              struct{}{},
			"medium-dark skin tone": struct{}{},
			"person running":        struct{}{},
			"running":               struct{}{},
		},
	},
	":person_running_tone5:": Code{
		"\U0001f3c3\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"marathon":       struct{}{},
			"person running": struct{}{},
			"running":        struct{}{},
		},
	},
	":man_running:": Code{
		"\U0001f3c3\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":         struct{}{},
			"man running": struct{}{},
			"marathon":    struct{}{},
			"racing":      struct{}{},
			"running":     struct{}{},
		},
	},
	":man_running_tone1:": Code{
		"\U0001f3c3\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man running":     struct{}{},
			"marathon":        struct{}{},
			"racing":          struct{}{},
			"running":         struct{}{},
		},
	},
	":man_running_tone2:": Code{
		"\U0001f3c3\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                    struct{}{},
			"man running":            struct{}{},
			"marathon":               struct{}{},
			"medium-light skin tone": struct{}{},
			"racing":                 struct{}{},
			"running":                struct{}{},
		},
	},
	":man_running_tone3:": Code{
		"\U0001f3c3\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":              struct{}{},
			"man running":      struct{}{},
			"marathon":         struct{}{},
			"medium skin tone": struct{}{},
			"racing":           struct{}{},
			"running":          struct{}{},
		},
	},
	":man_running_tone4:": Code{
		"\U0001f3c3\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                   struct{}{},
			"man running":           struct{}{},
			"marathon":              struct{}{},
			"medium-dark skin tone": struct{}{},
			"racing":                struct{}{},
			"running":               struct{}{},
		},
	},
	":man_running_tone5:": Code{
		"\U0001f3c3\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"man":            struct{}{},
			"man running":    struct{}{},
			"marathon":       struct{}{},
			"racing":         struct{}{},
			"running":        struct{}{},
		},
	},
	":woman_running:": Code{
		"\U0001f3c3\u200d\u2640\ufe0f",
		map[string]struct{}{
			"marathon":      struct{}{},
			"racing":        struct{}{},
			"running":       struct{}{},
			"woman":         struct{}{},
			"woman running": struct{}{},
		},
	},
	":woman_running_tone1:": Code{
		"\U0001f3c3\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"marathon":        struct{}{},
			"racing":          struct{}{},
			"running":         struct{}{},
			"woman":           struct{}{},
			"woman running":   struct{}{},
		},
	},
	":woman_running_tone2:": Code{
		"\U0001f3c3\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"marathon":               struct{}{},
			"medium-light skin tone": struct{}{},
			"racing":                 struct{}{},
			"running":                struct{}{},
			"woman":                  struct{}{},
			"woman running":          struct{}{},
		},
	},
	":woman_running_tone3:": Code{
		"\U0001f3c3\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"marathon":         struct{}{},
			"medium skin tone": struct{}{},
			"racing":           struct{}{},
			"running":          struct{}{},
			"woman":            struct{}{},
			"woman running":    struct{}{},
		},
	},
	":woman_running_tone4:": Code{
		"\U0001f3c3\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"marathon":              struct{}{},
			"medium-dark skin tone": struct{}{},
			"racing":                struct{}{},
			"running":               struct{}{},
			"woman":                 struct{}{},
			"woman running":         struct{}{},
		},
	},
	":woman_running_tone5:": Code{
		"\U0001f3c3\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"marathon":       struct{}{},
			"racing":         struct{}{},
			"running":        struct{}{},
			"woman":          struct{}{},
			"woman running":  struct{}{},
		},
	},
	":dancer:": Code{
		"\U0001f483",
		map[string]struct{}{
			"dancing":       struct{}{},
			"woman":         struct{}{},
			"woman dancing": struct{}{},
		},
	},
	":dancer_tone1:": Code{
		"\U0001f483\U0001f3fb",
		map[string]struct{}{
			"dancing":         struct{}{},
			"light skin tone": struct{}{},
			"woman":           struct{}{},
			"woman dancing":   struct{}{},
		},
	},
	":dancer_tone2:": Code{
		"\U0001f483\U0001f3fc",
		map[string]struct{}{
			"dancing":                struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
			"woman dancing":          struct{}{},
		},
	},
	":dancer_tone3:": Code{
		"\U0001f483\U0001f3fd",
		map[string]struct{}{
			"dancing":          struct{}{},
			"medium skin tone": struct{}{},
			"woman":            struct{}{},
			"woman dancing":    struct{}{},
		},
	},
	":dancer_tone4:": Code{
		"\U0001f483\U0001f3fe",
		map[string]struct{}{
			"dancing":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
			"woman dancing":         struct{}{},
		},
	},
	":dancer_tone5:": Code{
		"\U0001f483\U0001f3ff",
		map[string]struct{}{
			"dancing":        struct{}{},
			"dark skin tone": struct{}{},
			"woman":          struct{}{},
			"woman dancing":  struct{}{},
		},
	},
	":man_dancing:": Code{
		"\U0001f57a",
		map[string]struct{}{
			"dance":       struct{}{},
			"man":         struct{}{},
			"man dancing": struct{}{},
		},
	},
	":man_dancing_tone1:": Code{
		"\U0001f57a\U0001f3fb",
		map[string]struct{}{
			"dance":           struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man dancing":     struct{}{},
		},
	},
	":man_dancing_tone2:": Code{
		"\U0001f57a\U0001f3fc",
		map[string]struct{}{
			"dance":                  struct{}{},
			"man":                    struct{}{},
			"man dancing":            struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_dancing_tone3:": Code{
		"\U0001f57a\U0001f3fd",
		map[string]struct{}{
			"dance":            struct{}{},
			"man":              struct{}{},
			"man dancing":      struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":man_dancing_tone4:": Code{
		"\U0001f57a\U0001f3fe",
		map[string]struct{}{
			"dance":                 struct{}{},
			"man":                   struct{}{},
			"man dancing":           struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_dancing_tone5:": Code{
		"\U0001f57a\U0001f3ff",
		map[string]struct{}{
			"dance":          struct{}{},
			"dark skin tone": struct{}{},
			"man":            struct{}{},
			"man dancing":    struct{}{},
		},
	},
	":people_with_bunny_ears_partying:": Code{
		"\U0001f46f",
		map[string]struct{}{
			"bunny ear":              struct{}{},
			"dancer":                 struct{}{},
			"partying":               struct{}{},
			"people with bunny ears": struct{}{},
		},
	},
	":men_with_bunny_ears_partying:": Code{
		"\U0001f46f\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bunny ear":           struct{}{},
			"dancer":              struct{}{},
			"men":                 struct{}{},
			"men with bunny ears": struct{}{},
			"partying":            struct{}{},
		},
	},
	":women_with_bunny_ears_partying:": Code{
		"\U0001f46f\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bunny ear":             struct{}{},
			"dancer":                struct{}{},
			"partying":              struct{}{},
			"women":                 struct{}{},
			"women with bunny ears": struct{}{},
		},
	},
	":person_in_steamy_room:": Code{
		"\U0001f9d6",
		map[string]struct{}{
			"person in steamy room": struct{}{},
			"sauna":                 struct{}{},
			"steam room":            struct{}{},
			"hamam":                 struct{}{},
			"steambath":             struct{}{},
		},
	},
	":person_in_steamy_room_tone1:": Code{
		"\U0001f9d6\U0001f3fb",
		map[string]struct{}{
			"light skin tone":       struct{}{},
			"person in steamy room": struct{}{},
			"sauna":                 struct{}{},
			"steam room":            struct{}{},
		},
	},
	":person_in_steamy_room_tone2:": Code{
		"\U0001f9d6\U0001f3fc",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"person in steamy room":  struct{}{},
			"sauna":                  struct{}{},
			"steam room":             struct{}{},
		},
	},
	":person_in_steamy_room_tone3:": Code{
		"\U0001f9d6\U0001f3fd",
		map[string]struct{}{
			"medium skin tone":      struct{}{},
			"person in steamy room": struct{}{},
			"sauna":                 struct{}{},
			"steam room":            struct{}{},
		},
	},
	":person_in_steamy_room_tone4:": Code{
		"\U0001f9d6\U0001f3fe",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"person in steamy room": struct{}{},
			"sauna":                 struct{}{},
			"steam room":            struct{}{},
		},
	},
	":person_in_steamy_room_tone5:": Code{
		"\U0001f9d6\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":        struct{}{},
			"person in steamy room": struct{}{},
			"sauna":                 struct{}{},
			"steam room":            struct{}{},
		},
	},
	":woman_in_steamy_room:": Code{
		"\U0001f9d6\u200d\u2640\ufe0f",
		map[string]struct{}{
			"sauna":                struct{}{},
			"steam room":           struct{}{},
			"woman in steamy room": struct{}{},
		},
	},
	":woman_in_steamy_room_tone1:": Code{
		"\U0001f9d6\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"light skin tone":      struct{}{},
			"sauna":                struct{}{},
			"steam room":           struct{}{},
			"woman in steamy room": struct{}{},
		},
	},
	":woman_in_steamy_room_tone2:": Code{
		"\U0001f9d6\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"sauna":                  struct{}{},
			"steam room":             struct{}{},
			"woman in steamy room":   struct{}{},
		},
	},
	":woman_in_steamy_room_tone3:": Code{
		"\U0001f9d6\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium skin tone":     struct{}{},
			"sauna":                struct{}{},
			"steam room":           struct{}{},
			"woman in steamy room": struct{}{},
		},
	},
	":woman_in_steamy_room_tone4:": Code{
		"\U0001f9d6\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"sauna":                 struct{}{},
			"steam room":            struct{}{},
			"woman in steamy room":  struct{}{},
		},
	},
	":woman_in_steamy_room_tone5:": Code{
		"\U0001f9d6\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":       struct{}{},
			"sauna":                struct{}{},
			"steam room":           struct{}{},
			"woman in steamy room": struct{}{},
		},
	},
	":man_in_steamy_room:": Code{
		"\U0001f9d6\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man in steamy room": struct{}{},
			"sauna":              struct{}{},
			"steam room":         struct{}{},
		},
	},
	":man_in_steamy_room_tone1:": Code{
		"\U0001f9d6\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"light skin tone":    struct{}{},
			"man in steamy room": struct{}{},
			"sauna":              struct{}{},
			"steam room":         struct{}{},
		},
	},
	":man_in_steamy_room_tone2:": Code{
		"\U0001f9d6\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man in steamy room":     struct{}{},
			"medium-light skin tone": struct{}{},
			"sauna":                  struct{}{},
			"steam room":             struct{}{},
		},
	},
	":man_in_steamy_room_tone3:": Code{
		"\U0001f9d6\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man in steamy room": struct{}{},
			"medium skin tone":   struct{}{},
			"sauna":              struct{}{},
			"steam room":         struct{}{},
		},
	},
	":man_in_steamy_room_tone4:": Code{
		"\U0001f9d6\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man in steamy room":    struct{}{},
			"medium-dark skin tone": struct{}{},
			"sauna":                 struct{}{},
			"steam room":            struct{}{},
		},
	},
	":man_in_steamy_room_tone5:": Code{
		"\U0001f9d6\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone":     struct{}{},
			"man in steamy room": struct{}{},
			"sauna":              struct{}{},
			"steam room":         struct{}{},
		},
	},
	":person_climbing:": Code{
		"\U0001f9d7",
		map[string]struct{}{
			"climber":         struct{}{},
			"person climbing": struct{}{},
		},
	},
	":person_climbing_tone1:": Code{
		"\U0001f9d7\U0001f3fb",
		map[string]struct{}{
			"climber":         struct{}{},
			"light skin tone": struct{}{},
			"person climbing": struct{}{},
		},
	},
	":person_climbing_tone2:": Code{
		"\U0001f9d7\U0001f3fc",
		map[string]struct{}{
			"climber":                struct{}{},
			"medium-light skin tone": struct{}{},
			"person climbing":        struct{}{},
		},
	},
	":person_climbing_tone3:": Code{
		"\U0001f9d7\U0001f3fd",
		map[string]struct{}{
			"climber":          struct{}{},
			"medium skin tone": struct{}{},
			"person climbing":  struct{}{},
		},
	},
	":person_climbing_tone4:": Code{
		"\U0001f9d7\U0001f3fe",
		map[string]struct{}{
			"climber":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"person climbing":       struct{}{},
		},
	},
	":person_climbing_tone5:": Code{
		"\U0001f9d7\U0001f3ff",
		map[string]struct{}{
			"climber":         struct{}{},
			"dark skin tone":  struct{}{},
			"person climbing": struct{}{},
		},
	},
	":woman_climbing:": Code{
		"\U0001f9d7\u200d\u2640\ufe0f",
		map[string]struct{}{
			"climber":        struct{}{},
			"woman climbing": struct{}{},
		},
	},
	":woman_climbing_tone1:": Code{
		"\U0001f9d7\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"climber":         struct{}{},
			"light skin tone": struct{}{},
			"woman climbing":  struct{}{},
		},
	},
	":woman_climbing_tone2:": Code{
		"\U0001f9d7\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"climber":                struct{}{},
			"medium-light skin tone": struct{}{},
			"woman climbing":         struct{}{},
		},
	},
	":woman_climbing_tone3:": Code{
		"\U0001f9d7\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"climber":          struct{}{},
			"medium skin tone": struct{}{},
			"woman climbing":   struct{}{},
		},
	},
	":woman_climbing_tone4:": Code{
		"\U0001f9d7\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"climber":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman climbing":        struct{}{},
		},
	},
	":woman_climbing_tone5:": Code{
		"\U0001f9d7\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"climber":        struct{}{},
			"dark skin tone": struct{}{},
			"woman climbing": struct{}{},
		},
	},
	":man_climbing:": Code{
		"\U0001f9d7\u200d\u2642\ufe0f",
		map[string]struct{}{
			"climber":      struct{}{},
			"man climbing": struct{}{},
		},
	},
	":man_climbing_tone1:": Code{
		"\U0001f9d7\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"climber":         struct{}{},
			"light skin tone": struct{}{},
			"man climbing":    struct{}{},
		},
	},
	":man_climbing_tone2:": Code{
		"\U0001f9d7\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"climber":                struct{}{},
			"man climbing":           struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_climbing_tone3:": Code{
		"\U0001f9d7\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"climber":          struct{}{},
			"man climbing":     struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":man_climbing_tone4:": Code{
		"\U0001f9d7\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"climber":               struct{}{},
			"man climbing":          struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_climbing_tone5:": Code{
		"\U0001f9d7\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"climber":        struct{}{},
			"dark skin tone": struct{}{},
			"man climbing":   struct{}{},
		},
	},
	":person_in_lotus_position:": Code{
		"\U0001f9d8",
		map[string]struct{}{
			"meditation":               struct{}{},
			"person in lotus position": struct{}{},
			"yoga":                     struct{}{},
			"serenity":                 struct{}{},
		},
	},
	":person_in_lotus_position_tone1:": Code{
		"\U0001f9d8\U0001f3fb",
		map[string]struct{}{
			"light skin tone":          struct{}{},
			"meditation":               struct{}{},
			"person in lotus position": struct{}{},
			"yoga":                     struct{}{},
		},
	},
	":person_in_lotus_position_tone2:": Code{
		"\U0001f9d8\U0001f3fc",
		map[string]struct{}{
			"meditation":               struct{}{},
			"medium-light skin tone":   struct{}{},
			"person in lotus position": struct{}{},
			"yoga":                     struct{}{},
		},
	},
	":person_in_lotus_position_tone3:": Code{
		"\U0001f9d8\U0001f3fd",
		map[string]struct{}{
			"meditation":               struct{}{},
			"medium skin tone":         struct{}{},
			"person in lotus position": struct{}{},
			"yoga":                     struct{}{},
		},
	},
	":person_in_lotus_position_tone4:": Code{
		"\U0001f9d8\U0001f3fe",
		map[string]struct{}{
			"meditation":               struct{}{},
			"medium-dark skin tone":    struct{}{},
			"person in lotus position": struct{}{},
			"yoga":                     struct{}{},
		},
	},
	":person_in_lotus_position_tone5:": Code{
		"\U0001f9d8\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":           struct{}{},
			"meditation":               struct{}{},
			"person in lotus position": struct{}{},
			"yoga":                     struct{}{},
		},
	},
	":woman_in_lotus_position:": Code{
		"\U0001f9d8\u200d\u2640\ufe0f",
		map[string]struct{}{
			"meditation":              struct{}{},
			"woman in lotus position": struct{}{},
			"yoga":                    struct{}{},
		},
	},
	":woman_in_lotus_position_tone1:": Code{
		"\U0001f9d8\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"light skin tone":         struct{}{},
			"meditation":              struct{}{},
			"woman in lotus position": struct{}{},
			"yoga":                    struct{}{},
		},
	},
	":woman_in_lotus_position_tone2:": Code{
		"\U0001f9d8\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"meditation":              struct{}{},
			"medium-light skin tone":  struct{}{},
			"woman in lotus position": struct{}{},
			"yoga":                    struct{}{},
		},
	},
	":woman_in_lotus_position_tone3:": Code{
		"\U0001f9d8\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"meditation":              struct{}{},
			"medium skin tone":        struct{}{},
			"woman in lotus position": struct{}{},
			"yoga":                    struct{}{},
		},
	},
	":woman_in_lotus_position_tone4:": Code{
		"\U0001f9d8\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"meditation":              struct{}{},
			"medium-dark skin tone":   struct{}{},
			"woman in lotus position": struct{}{},
			"yoga":                    struct{}{},
		},
	},
	":woman_in_lotus_position_tone5:": Code{
		"\U0001f9d8\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":          struct{}{},
			"meditation":              struct{}{},
			"woman in lotus position": struct{}{},
			"yoga":                    struct{}{},
		},
	},
	":man_in_lotus_position:": Code{
		"\U0001f9d8\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man in lotus position": struct{}{},
			"meditation":            struct{}{},
			"yoga":                  struct{}{},
		},
	},
	":man_in_lotus_position_tone1:": Code{
		"\U0001f9d8\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"light skin tone":       struct{}{},
			"man in lotus position": struct{}{},
			"meditation":            struct{}{},
			"yoga":                  struct{}{},
		},
	},
	":man_in_lotus_position_tone2:": Code{
		"\U0001f9d8\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man in lotus position":  struct{}{},
			"meditation":             struct{}{},
			"medium-light skin tone": struct{}{},
			"yoga":                   struct{}{},
		},
	},
	":man_in_lotus_position_tone3:": Code{
		"\U0001f9d8\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man in lotus position": struct{}{},
			"meditation":            struct{}{},
			"medium skin tone":      struct{}{},
			"yoga":                  struct{}{},
		},
	},
	":man_in_lotus_position_tone4:": Code{
		"\U0001f9d8\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man in lotus position": struct{}{},
			"meditation":            struct{}{},
			"medium-dark skin tone": struct{}{},
			"yoga":                  struct{}{},
		},
	},
	":man_in_lotus_position_tone5:": Code{
		"\U0001f9d8\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone":        struct{}{},
			"man in lotus position": struct{}{},
			"meditation":            struct{}{},
			"yoga":                  struct{}{},
		},
	},
	":bath:": Code{
		"\U0001f6c0",
		map[string]struct{}{
			"bath":               struct{}{},
			"bathtub":            struct{}{},
			"person taking bath": struct{}{},
		},
	},
	":bath_tone1:": Code{
		"\U0001f6c0\U0001f3fb",
		map[string]struct{}{
			"bath":               struct{}{},
			"bathtub":            struct{}{},
			"light skin tone":    struct{}{},
			"person taking bath": struct{}{},
		},
	},
	":bath_tone2:": Code{
		"\U0001f6c0\U0001f3fc",
		map[string]struct{}{
			"bath":                   struct{}{},
			"bathtub":                struct{}{},
			"medium-light skin tone": struct{}{},
			"person taking bath":     struct{}{},
		},
	},
	":bath_tone3:": Code{
		"\U0001f6c0\U0001f3fd",
		map[string]struct{}{
			"bath":               struct{}{},
			"bathtub":            struct{}{},
			"medium skin tone":   struct{}{},
			"person taking bath": struct{}{},
		},
	},
	":bath_tone4:": Code{
		"\U0001f6c0\U0001f3fe",
		map[string]struct{}{
			"bath":                  struct{}{},
			"bathtub":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"person taking bath":    struct{}{},
		},
	},
	":bath_tone5:": Code{
		"\U0001f6c0\U0001f3ff",
		map[string]struct{}{
			"bath":               struct{}{},
			"bathtub":            struct{}{},
			"dark skin tone":     struct{}{},
			"person taking bath": struct{}{},
		},
	},
	":sleeping_accommodation:": Code{
		"\U0001f6cc",
		map[string]struct{}{
			"hotel":         struct{}{},
			"person in bed": struct{}{},
			"sleep":         struct{}{},
		},
	},
	":person_in_bed_tone1:": Code{
		"\U0001f6cc\U0001f3fb",
		map[string]struct{}{
			"hotel":           struct{}{},
			"light skin tone": struct{}{},
			"person in bed":   struct{}{},
			"sleep":           struct{}{},
		},
	},
	":person_in_bed_tone2:": Code{
		"\U0001f6cc\U0001f3fc",
		map[string]struct{}{
			"hotel":                  struct{}{},
			"medium-light skin tone": struct{}{},
			"person in bed":          struct{}{},
			"sleep":                  struct{}{},
		},
	},
	":person_in_bed_tone3:": Code{
		"\U0001f6cc\U0001f3fd",
		map[string]struct{}{
			"hotel":            struct{}{},
			"medium skin tone": struct{}{},
			"person in bed":    struct{}{},
			"sleep":            struct{}{},
		},
	},
	":person_in_bed_tone4:": Code{
		"\U0001f6cc\U0001f3fe",
		map[string]struct{}{
			"hotel":                 struct{}{},
			"medium-dark skin tone": struct{}{},
			"person in bed":         struct{}{},
			"sleep":                 struct{}{},
		},
	},
	":person_in_bed_tone5:": Code{
		"\U0001f6cc\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"hotel":          struct{}{},
			"person in bed":  struct{}{},
			"sleep":          struct{}{},
		},
	},
	":man_in_business_suit_levitating_tone1:": Code{
		"\U0001f574\U0001f3fb",
		map[string]struct{}{
			"business":               struct{}{},
			"light skin tone":        struct{}{},
			"man":                    struct{}{},
			"man in suit levitating": struct{}{},
			"suit":                   struct{}{},
		},
	},
	":man_in_business_suit_levitating_tone2:": Code{
		"\U0001f574\U0001f3fc",
		map[string]struct{}{
			"business":               struct{}{},
			"man":                    struct{}{},
			"man in suit levitating": struct{}{},
			"medium-light skin tone": struct{}{},
			"suit":                   struct{}{},
		},
	},
	":man_in_business_suit_levitating_tone3:": Code{
		"\U0001f574\U0001f3fd",
		map[string]struct{}{
			"business":               struct{}{},
			"man":                    struct{}{},
			"man in suit levitating": struct{}{},
			"medium skin tone":       struct{}{},
			"suit":                   struct{}{},
		},
	},
	":man_in_business_suit_levitating_tone4:": Code{
		"\U0001f574\U0001f3fe",
		map[string]struct{}{
			"business":               struct{}{},
			"man":                    struct{}{},
			"man in suit levitating": struct{}{},
			"medium-dark skin tone":  struct{}{},
			"suit":                   struct{}{},
		},
	},
	":man_in_business_suit_levitating_tone5:": Code{
		"\U0001f574\U0001f3ff",
		map[string]struct{}{
			"business":               struct{}{},
			"dark skin tone":         struct{}{},
			"man":                    struct{}{},
			"man in suit levitating": struct{}{},
			"suit":                   struct{}{},
		},
	},
	":speaking_head:": Code{
		"\U0001f5e3",
		map[string]struct{}{
			"face":          struct{}{},
			"head":          struct{}{},
			"silhouette":    struct{}{},
			"speak":         struct{}{},
			"speaking":      struct{}{},
			"speaking head": struct{}{},
		},
	},
	":bust_in_silhouette:": Code{
		"\U0001f464",
		map[string]struct{}{
			"bust":               struct{}{},
			"bust in silhouette": struct{}{},
			"silhouette":         struct{}{},
		},
	},
	":busts_in_silhouette:": Code{
		"\U0001f465",
		map[string]struct{}{
			"bust":                struct{}{},
			"busts in silhouette": struct{}{},
			"silhouette":          struct{}{},
		},
	},
	":person_fencing:": Code{
		"\U0001f93a",
		map[string]struct{}{
			"fencer":         struct{}{},
			"fencing":        struct{}{},
			"person fencing": struct{}{},
			"sword":          struct{}{},
		},
	},
	":horse_racing:": Code{
		"\U0001f3c7",
		map[string]struct{}{
			"horse":        struct{}{},
			"horse racing": struct{}{},
			"jockey":       struct{}{},
			"racehorse":    struct{}{},
			"racing":       struct{}{},
		},
	},
	":horse_racing_tone1:": Code{
		"\U0001f3c7\U0001f3fb",
		map[string]struct{}{
			"horse":           struct{}{},
			"horse racing":    struct{}{},
			"jockey":          struct{}{},
			"light skin tone": struct{}{},
			"racehorse":       struct{}{},
			"racing":          struct{}{},
		},
	},
	":horse_racing_tone2:": Code{
		"\U0001f3c7\U0001f3fc",
		map[string]struct{}{
			"horse":                  struct{}{},
			"horse racing":           struct{}{},
			"jockey":                 struct{}{},
			"medium-light skin tone": struct{}{},
			"racehorse":              struct{}{},
			"racing":                 struct{}{},
		},
	},
	":horse_racing_tone3:": Code{
		"\U0001f3c7\U0001f3fd",
		map[string]struct{}{
			"horse":            struct{}{},
			"horse racing":     struct{}{},
			"jockey":           struct{}{},
			"medium skin tone": struct{}{},
			"racehorse":        struct{}{},
			"racing":           struct{}{},
		},
	},
	":horse_racing_tone4:": Code{
		"\U0001f3c7\U0001f3fe",
		map[string]struct{}{
			"horse":                 struct{}{},
			"horse racing":          struct{}{},
			"jockey":                struct{}{},
			"medium-dark skin tone": struct{}{},
			"racehorse":             struct{}{},
			"racing":                struct{}{},
		},
	},
	":horse_racing_tone5:": Code{
		"\U0001f3c7\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"horse":          struct{}{},
			"horse racing":   struct{}{},
			"jockey":         struct{}{},
			"racehorse":      struct{}{},
			"racing":         struct{}{},
		},
	},
	":skier:": Code{
		"\u26f7",
		map[string]struct{}{
			"ski":   struct{}{},
			"skier": struct{}{},
			"snow":  struct{}{},
		},
	},
	":snowboarder:": Code{
		"\U0001f3c2",
		map[string]struct{}{
			"ski":         struct{}{},
			"snow":        struct{}{},
			"snowboard":   struct{}{},
			"snowboarder": struct{}{},
		},
	},
	":snowboarder_tone1:": Code{
		"\U0001f3c2\U0001f3fb",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"ski":             struct{}{},
			"snow":            struct{}{},
			"snowboard":       struct{}{},
			"snowboarder":     struct{}{},
		},
	},
	":snowboarder_tone2:": Code{
		"\U0001f3c2\U0001f3fc",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"ski":                    struct{}{},
			"snow":                   struct{}{},
			"snowboard":              struct{}{},
			"snowboarder":            struct{}{},
		},
	},
	":snowboarder_tone3:": Code{
		"\U0001f3c2\U0001f3fd",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"ski":              struct{}{},
			"snow":             struct{}{},
			"snowboard":        struct{}{},
			"snowboarder":      struct{}{},
		},
	},
	":snowboarder_tone4:": Code{
		"\U0001f3c2\U0001f3fe",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"ski":                   struct{}{},
			"snow":                  struct{}{},
			"snowboard":             struct{}{},
			"snowboarder":           struct{}{},
		},
	},
	":snowboarder_tone5:": Code{
		"\U0001f3c2\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"ski":            struct{}{},
			"snow":           struct{}{},
			"snowboard":      struct{}{},
			"snowboarder":    struct{}{},
		},
	},
	":person_golfing:": Code{
		"\U0001f3cc",
		map[string]struct{}{
			"ball":           struct{}{},
			"golf":           struct{}{},
			"person golfing": struct{}{},
		},
	},
	":person_golfing_tone1:": Code{
		"\U0001f3cc\U0001f3fb",
		map[string]struct{}{
			"ball":            struct{}{},
			"golf":            struct{}{},
			"light skin tone": struct{}{},
			"person golfing":  struct{}{},
		},
	},
	":person_golfing_tone2:": Code{
		"\U0001f3cc\U0001f3fc",
		map[string]struct{}{
			"ball":                   struct{}{},
			"golf":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"person golfing":         struct{}{},
		},
	},
	":person_golfing_tone3:": Code{
		"\U0001f3cc\U0001f3fd",
		map[string]struct{}{
			"ball":             struct{}{},
			"golf":             struct{}{},
			"medium skin tone": struct{}{},
			"person golfing":   struct{}{},
		},
	},
	":person_golfing_tone4:": Code{
		"\U0001f3cc\U0001f3fe",
		map[string]struct{}{
			"ball":                  struct{}{},
			"golf":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"person golfing":        struct{}{},
		},
	},
	":person_golfing_tone5:": Code{
		"\U0001f3cc\U0001f3ff",
		map[string]struct{}{
			"ball":           struct{}{},
			"dark skin tone": struct{}{},
			"golf":           struct{}{},
			"person golfing": struct{}{},
		},
	},
	":man_golfing:": Code{
		"\U0001f3cc\ufe0f\u200d\u2642\ufe0f",
		map[string]struct{}{
			"golf":        struct{}{},
			"man":         struct{}{},
			"man golfing": struct{}{},
		},
	},
	":man_golfing_tone1:": Code{
		"\U0001f3cc\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"golf":            struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man golfing":     struct{}{},
		},
	},
	":man_golfing_tone2:": Code{
		"\U0001f3cc\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"golf":                   struct{}{},
			"man":                    struct{}{},
			"man golfing":            struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_golfing_tone3:": Code{
		"\U0001f3cc\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"golf":             struct{}{},
			"man":              struct{}{},
			"man golfing":      struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":man_golfing_tone4:": Code{
		"\U0001f3cc\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"golf":                  struct{}{},
			"man":                   struct{}{},
			"man golfing":           struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_golfing_tone5:": Code{
		"\U0001f3cc\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"golf":           struct{}{},
			"man":            struct{}{},
			"man golfing":    struct{}{},
		},
	},
	":woman_golfing:": Code{
		"\U0001f3cc\ufe0f\u200d\u2640\ufe0f",
		map[string]struct{}{
			"golf":          struct{}{},
			"woman":         struct{}{},
			"woman golfing": struct{}{},
		},
	},
	":woman_golfing_tone1:": Code{
		"\U0001f3cc\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"golf":            struct{}{},
			"light skin tone": struct{}{},
			"woman":           struct{}{},
			"woman golfing":   struct{}{},
		},
	},
	":woman_golfing_tone2:": Code{
		"\U0001f3cc\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"golf":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
			"woman golfing":          struct{}{},
		},
	},
	":woman_golfing_tone3:": Code{
		"\U0001f3cc\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"golf":             struct{}{},
			"medium skin tone": struct{}{},
			"woman":            struct{}{},
			"woman golfing":    struct{}{},
		},
	},
	":woman_golfing_tone4:": Code{
		"\U0001f3cc\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"golf":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
			"woman golfing":         struct{}{},
		},
	},
	":woman_golfing_tone5:": Code{
		"\U0001f3cc\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"golf":           struct{}{},
			"woman":          struct{}{},
			"woman golfing":  struct{}{},
		},
	},
	":person_surfing:": Code{
		"\U0001f3c4",
		map[string]struct{}{
			"person surfing": struct{}{},
			"surfing":        struct{}{},
		},
	},
	":person_surfing_tone1:": Code{
		"\U0001f3c4\U0001f3fb",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"person surfing":  struct{}{},
			"surfing":         struct{}{},
		},
	},
	":person_surfing_tone2:": Code{
		"\U0001f3c4\U0001f3fc",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"person surfing":         struct{}{},
			"surfing":                struct{}{},
		},
	},
	":person_surfing_tone3:": Code{
		"\U0001f3c4\U0001f3fd",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"person surfing":   struct{}{},
			"surfing":          struct{}{},
		},
	},
	":person_surfing_tone4:": Code{
		"\U0001f3c4\U0001f3fe",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"person surfing":        struct{}{},
			"surfing":               struct{}{},
		},
	},
	":person_surfing_tone5:": Code{
		"\U0001f3c4\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"person surfing": struct{}{},
			"surfing":        struct{}{},
		},
	},
	":man_surfing:": Code{
		"\U0001f3c4\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":         struct{}{},
			"man surfing": struct{}{},
			"surfing":     struct{}{},
		},
	},
	":man_surfing_tone1:": Code{
		"\U0001f3c4\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man surfing":     struct{}{},
			"surfing":         struct{}{},
		},
	},
	":man_surfing_tone2:": Code{
		"\U0001f3c4\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                    struct{}{},
			"man surfing":            struct{}{},
			"medium-light skin tone": struct{}{},
			"surfing":                struct{}{},
		},
	},
	":man_surfing_tone3:": Code{
		"\U0001f3c4\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":              struct{}{},
			"man surfing":      struct{}{},
			"medium skin tone": struct{}{},
			"surfing":          struct{}{},
		},
	},
	":man_surfing_tone4:": Code{
		"\U0001f3c4\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                   struct{}{},
			"man surfing":           struct{}{},
			"medium-dark skin tone": struct{}{},
			"surfing":               struct{}{},
		},
	},
	":man_surfing_tone5:": Code{
		"\U0001f3c4\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"man":            struct{}{},
			"man surfing":    struct{}{},
			"surfing":        struct{}{},
		},
	},
	":woman_surfing:": Code{
		"\U0001f3c4\u200d\u2640\ufe0f",
		map[string]struct{}{
			"surfing":       struct{}{},
			"woman":         struct{}{},
			"woman surfing": struct{}{},
		},
	},
	":woman_surfing_tone1:": Code{
		"\U0001f3c4\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"surfing":         struct{}{},
			"woman":           struct{}{},
			"woman surfing":   struct{}{},
		},
	},
	":woman_surfing_tone2:": Code{
		"\U0001f3c4\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"surfing":                struct{}{},
			"woman":                  struct{}{},
			"woman surfing":          struct{}{},
		},
	},
	":woman_surfing_tone3:": Code{
		"\U0001f3c4\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"surfing":          struct{}{},
			"woman":            struct{}{},
			"woman surfing":    struct{}{},
		},
	},
	":woman_surfing_tone4:": Code{
		"\U0001f3c4\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"surfing":               struct{}{},
			"woman":                 struct{}{},
			"woman surfing":         struct{}{},
		},
	},
	":woman_surfing_tone5:": Code{
		"\U0001f3c4\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"surfing":        struct{}{},
			"woman":          struct{}{},
			"woman surfing":  struct{}{},
		},
	},
	":person_rowing_boat:": Code{
		"\U0001f6a3",
		map[string]struct{}{
			"boat":               struct{}{},
			"person rowing boat": struct{}{},
			"rowboat":            struct{}{},
		},
	},
	":person_rowing_boat_tone1:": Code{
		"\U0001f6a3\U0001f3fb",
		map[string]struct{}{
			"boat":               struct{}{},
			"light skin tone":    struct{}{},
			"person rowing boat": struct{}{},
			"rowboat":            struct{}{},
		},
	},
	":person_rowing_boat_tone2:": Code{
		"\U0001f6a3\U0001f3fc",
		map[string]struct{}{
			"boat":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"person rowing boat":     struct{}{},
			"rowboat":                struct{}{},
		},
	},
	":person_rowing_boat_tone3:": Code{
		"\U0001f6a3\U0001f3fd",
		map[string]struct{}{
			"boat":               struct{}{},
			"medium skin tone":   struct{}{},
			"person rowing boat": struct{}{},
			"rowboat":            struct{}{},
		},
	},
	":person_rowing_boat_tone4:": Code{
		"\U0001f6a3\U0001f3fe",
		map[string]struct{}{
			"boat":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"person rowing boat":    struct{}{},
			"rowboat":               struct{}{},
		},
	},
	":person_rowing_boat_tone5:": Code{
		"\U0001f6a3\U0001f3ff",
		map[string]struct{}{
			"boat":               struct{}{},
			"dark skin tone":     struct{}{},
			"person rowing boat": struct{}{},
			"rowboat":            struct{}{},
		},
	},
	":man_rowing_boat:": Code{
		"\U0001f6a3\u200d\u2642\ufe0f",
		map[string]struct{}{
			"boat":            struct{}{},
			"man":             struct{}{},
			"man rowing boat": struct{}{},
			"rowboat":         struct{}{},
		},
	},
	":man_rowing_boat_tone1:": Code{
		"\U0001f6a3\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"boat":            struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man rowing boat": struct{}{},
			"rowboat":         struct{}{},
		},
	},
	":man_rowing_boat_tone2:": Code{
		"\U0001f6a3\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"boat":                   struct{}{},
			"man":                    struct{}{},
			"man rowing boat":        struct{}{},
			"medium-light skin tone": struct{}{},
			"rowboat":                struct{}{},
		},
	},
	":man_rowing_boat_tone3:": Code{
		"\U0001f6a3\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"boat":             struct{}{},
			"man":              struct{}{},
			"man rowing boat":  struct{}{},
			"medium skin tone": struct{}{},
			"rowboat":          struct{}{},
		},
	},
	":man_rowing_boat_tone4:": Code{
		"\U0001f6a3\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"boat":                  struct{}{},
			"man":                   struct{}{},
			"man rowing boat":       struct{}{},
			"medium-dark skin tone": struct{}{},
			"rowboat":               struct{}{},
		},
	},
	":man_rowing_boat_tone5:": Code{
		"\U0001f6a3\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"boat":            struct{}{},
			"dark skin tone":  struct{}{},
			"man":             struct{}{},
			"man rowing boat": struct{}{},
			"rowboat":         struct{}{},
		},
	},
	":woman_rowing_boat:": Code{
		"\U0001f6a3\u200d\u2640\ufe0f",
		map[string]struct{}{
			"boat":              struct{}{},
			"rowboat":           struct{}{},
			"woman":             struct{}{},
			"woman rowing boat": struct{}{},
		},
	},
	":woman_rowing_boat_tone1:": Code{
		"\U0001f6a3\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"boat":              struct{}{},
			"light skin tone":   struct{}{},
			"rowboat":           struct{}{},
			"woman":             struct{}{},
			"woman rowing boat": struct{}{},
		},
	},
	":woman_rowing_boat_tone2:": Code{
		"\U0001f6a3\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"boat":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"rowboat":                struct{}{},
			"woman":                  struct{}{},
			"woman rowing boat":      struct{}{},
		},
	},
	":woman_rowing_boat_tone3:": Code{
		"\U0001f6a3\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"boat":              struct{}{},
			"medium skin tone":  struct{}{},
			"rowboat":           struct{}{},
			"woman":             struct{}{},
			"woman rowing boat": struct{}{},
		},
	},
	":woman_rowing_boat_tone4:": Code{
		"\U0001f6a3\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"boat":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"rowboat":               struct{}{},
			"woman":                 struct{}{},
			"woman rowing boat":     struct{}{},
		},
	},
	":woman_rowing_boat_tone5:": Code{
		"\U0001f6a3\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"boat":              struct{}{},
			"dark skin tone":    struct{}{},
			"rowboat":           struct{}{},
			"woman":             struct{}{},
			"woman rowing boat": struct{}{},
		},
	},
	":person_swimming:": Code{
		"\U0001f3ca",
		map[string]struct{}{
			"person swimming": struct{}{},
			"swim":            struct{}{},
		},
	},
	":person_swimming_tone1:": Code{
		"\U0001f3ca\U0001f3fb",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"person swimming": struct{}{},
			"swim":            struct{}{},
		},
	},
	":person_swimming_tone2:": Code{
		"\U0001f3ca\U0001f3fc",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"person swimming":        struct{}{},
			"swim":                   struct{}{},
		},
	},
	":person_swimming_tone3:": Code{
		"\U0001f3ca\U0001f3fd",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"person swimming":  struct{}{},
			"swim":             struct{}{},
		},
	},
	":person_swimming_tone4:": Code{
		"\U0001f3ca\U0001f3fe",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"person swimming":       struct{}{},
			"swim":                  struct{}{},
		},
	},
	":person_swimming_tone5:": Code{
		"\U0001f3ca\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":  struct{}{},
			"person swimming": struct{}{},
			"swim":            struct{}{},
		},
	},
	":man_swimming:": Code{
		"\U0001f3ca\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":          struct{}{},
			"man swimming": struct{}{},
			"swim":         struct{}{},
		},
	},
	":man_swimming_tone1:": Code{
		"\U0001f3ca\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man swimming":    struct{}{},
			"swim":            struct{}{},
		},
	},
	":man_swimming_tone2:": Code{
		"\U0001f3ca\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                    struct{}{},
			"man swimming":           struct{}{},
			"medium-light skin tone": struct{}{},
			"swim":                   struct{}{},
		},
	},
	":man_swimming_tone3:": Code{
		"\U0001f3ca\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":              struct{}{},
			"man swimming":     struct{}{},
			"medium skin tone": struct{}{},
			"swim":             struct{}{},
		},
	},
	":man_swimming_tone4:": Code{
		"\U0001f3ca\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                   struct{}{},
			"man swimming":          struct{}{},
			"medium-dark skin tone": struct{}{},
			"swim":                  struct{}{},
		},
	},
	":man_swimming_tone5:": Code{
		"\U0001f3ca\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"man":            struct{}{},
			"man swimming":   struct{}{},
			"swim":           struct{}{},
		},
	},
	":woman_swimming:": Code{
		"\U0001f3ca\u200d\u2640\ufe0f",
		map[string]struct{}{
			"swim":           struct{}{},
			"woman":          struct{}{},
			"woman swimming": struct{}{},
		},
	},
	":woman_swimming_tone1:": Code{
		"\U0001f3ca\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"light skin tone": struct{}{},
			"swim":            struct{}{},
			"woman":           struct{}{},
			"woman swimming":  struct{}{},
		},
	},
	":woman_swimming_tone2:": Code{
		"\U0001f3ca\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"swim":                   struct{}{},
			"woman":                  struct{}{},
			"woman swimming":         struct{}{},
		},
	},
	":woman_swimming_tone3:": Code{
		"\U0001f3ca\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium skin tone": struct{}{},
			"swim":             struct{}{},
			"woman":            struct{}{},
			"woman swimming":   struct{}{},
		},
	},
	":woman_swimming_tone4:": Code{
		"\U0001f3ca\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"swim":                  struct{}{},
			"woman":                 struct{}{},
			"woman swimming":        struct{}{},
		},
	},
	":woman_swimming_tone5:": Code{
		"\U0001f3ca\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"swim":           struct{}{},
			"woman":          struct{}{},
			"woman swimming": struct{}{},
		},
	},
	":person_bouncing_ball:": Code{
		"\u26f9",
		map[string]struct{}{
			"ball":                 struct{}{},
			"person bouncing ball": struct{}{},
		},
	},
	":person_bouncing_ball_tone1:": Code{
		"\u26f9\U0001f3fb",
		map[string]struct{}{
			"ball":                 struct{}{},
			"light skin tone":      struct{}{},
			"person bouncing ball": struct{}{},
		},
	},
	":person_bouncing_ball_tone2:": Code{
		"\u26f9\U0001f3fc",
		map[string]struct{}{
			"ball":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"person bouncing ball":   struct{}{},
		},
	},
	":person_bouncing_ball_tone3:": Code{
		"\u26f9\U0001f3fd",
		map[string]struct{}{
			"ball":                 struct{}{},
			"medium skin tone":     struct{}{},
			"person bouncing ball": struct{}{},
		},
	},
	":person_bouncing_ball_tone4:": Code{
		"\u26f9\U0001f3fe",
		map[string]struct{}{
			"ball":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"person bouncing ball":  struct{}{},
		},
	},
	":person_bouncing_ball_tone5:": Code{
		"\u26f9\U0001f3ff",
		map[string]struct{}{
			"ball":                 struct{}{},
			"dark skin tone":       struct{}{},
			"person bouncing ball": struct{}{},
		},
	},
	":man_bouncing_ball:": Code{
		"\u26f9\ufe0f\u200d\u2642\ufe0f",
		map[string]struct{}{
			"ball":              struct{}{},
			"man":               struct{}{},
			"man bouncing ball": struct{}{},
		},
	},
	":man_bouncing_ball_tone1:": Code{
		"\u26f9\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"ball":              struct{}{},
			"light skin tone":   struct{}{},
			"man":               struct{}{},
			"man bouncing ball": struct{}{},
		},
	},
	":man_bouncing_ball_tone2:": Code{
		"\u26f9\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"ball":                   struct{}{},
			"man":                    struct{}{},
			"man bouncing ball":      struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_bouncing_ball_tone3:": Code{
		"\u26f9\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"ball":              struct{}{},
			"man":               struct{}{},
			"man bouncing ball": struct{}{},
			"medium skin tone":  struct{}{},
		},
	},
	":man_bouncing_ball_tone4:": Code{
		"\u26f9\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"ball":                  struct{}{},
			"man":                   struct{}{},
			"man bouncing ball":     struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_bouncing_ball_tone5:": Code{
		"\u26f9\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"ball":              struct{}{},
			"dark skin tone":    struct{}{},
			"man":               struct{}{},
			"man bouncing ball": struct{}{},
		},
	},
	":woman_bouncing_ball:": Code{
		"\u26f9\ufe0f\u200d\u2640\ufe0f",
		map[string]struct{}{
			"ball":                struct{}{},
			"woman":               struct{}{},
			"woman bouncing ball": struct{}{},
		},
	},
	":woman_bouncing_ball_tone1:": Code{
		"\u26f9\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"ball":                struct{}{},
			"light skin tone":     struct{}{},
			"woman":               struct{}{},
			"woman bouncing ball": struct{}{},
		},
	},
	":woman_bouncing_ball_tone2:": Code{
		"\u26f9\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"ball":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
			"woman bouncing ball":    struct{}{},
		},
	},
	":woman_bouncing_ball_tone3:": Code{
		"\u26f9\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"ball":                struct{}{},
			"medium skin tone":    struct{}{},
			"woman":               struct{}{},
			"woman bouncing ball": struct{}{},
		},
	},
	":woman_bouncing_ball_tone4:": Code{
		"\u26f9\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"ball":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
			"woman bouncing ball":   struct{}{},
		},
	},
	":woman_bouncing_ball_tone5:": Code{
		"\u26f9\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"ball":                struct{}{},
			"dark skin tone":      struct{}{},
			"woman":               struct{}{},
			"woman bouncing ball": struct{}{},
		},
	},
	":person_lifting_weights:": Code{
		"\U0001f3cb",
		map[string]struct{}{
			"lifter":                 struct{}{},
			"person lifting weights": struct{}{},
			"weight":                 struct{}{},
		},
	},
	":person_lifting_weights_tone1:": Code{
		"\U0001f3cb\U0001f3fb",
		map[string]struct{}{
			"lifter":                 struct{}{},
			"light skin tone":        struct{}{},
			"person lifting weights": struct{}{},
			"weight":                 struct{}{},
		},
	},
	":person_lifting_weights_tone2:": Code{
		"\U0001f3cb\U0001f3fc",
		map[string]struct{}{
			"lifter":                 struct{}{},
			"medium-light skin tone": struct{}{},
			"person lifting weights": struct{}{},
			"weight":                 struct{}{},
		},
	},
	":person_lifting_weights_tone3:": Code{
		"\U0001f3cb\U0001f3fd",
		map[string]struct{}{
			"lifter":                 struct{}{},
			"medium skin tone":       struct{}{},
			"person lifting weights": struct{}{},
			"weight":                 struct{}{},
		},
	},
	":person_lifting_weights_tone4:": Code{
		"\U0001f3cb\U0001f3fe",
		map[string]struct{}{
			"lifter":                 struct{}{},
			"medium-dark skin tone":  struct{}{},
			"person lifting weights": struct{}{},
			"weight":                 struct{}{},
		},
	},
	":person_lifting_weights_tone5:": Code{
		"\U0001f3cb\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":         struct{}{},
			"lifter":                 struct{}{},
			"person lifting weights": struct{}{},
			"weight":                 struct{}{},
		},
	},
	":man_lifting_weights:": Code{
		"\U0001f3cb\ufe0f\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                 struct{}{},
			"man lifting weights": struct{}{},
			"weight lifter":       struct{}{},
		},
	},
	":man_lifting_weights_tone1:": Code{
		"\U0001f3cb\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"light skin tone":     struct{}{},
			"man":                 struct{}{},
			"man lifting weights": struct{}{},
			"weight lifter":       struct{}{},
		},
	},
	":man_lifting_weights_tone2:": Code{
		"\U0001f3cb\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                    struct{}{},
			"man lifting weights":    struct{}{},
			"medium-light skin tone": struct{}{},
			"weight lifter":          struct{}{},
		},
	},
	":man_lifting_weights_tone3:": Code{
		"\U0001f3cb\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                 struct{}{},
			"man lifting weights": struct{}{},
			"medium skin tone":    struct{}{},
			"weight lifter":       struct{}{},
		},
	},
	":man_lifting_weights_tone4:": Code{
		"\U0001f3cb\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                   struct{}{},
			"man lifting weights":   struct{}{},
			"medium-dark skin tone": struct{}{},
			"weight lifter":         struct{}{},
		},
	},
	":man_lifting_weights_tone5:": Code{
		"\U0001f3cb\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone":      struct{}{},
			"man":                 struct{}{},
			"man lifting weights": struct{}{},
			"weight lifter":       struct{}{},
		},
	},
	":woman_lifting_weights:": Code{
		"\U0001f3cb\ufe0f\u200d\u2640\ufe0f",
		map[string]struct{}{
			"weight lifter":         struct{}{},
			"woman":                 struct{}{},
			"woman lifting weights": struct{}{},
		},
	},
	":woman_lifting_weights_tone1:": Code{
		"\U0001f3cb\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"light skin tone":       struct{}{},
			"weight lifter":         struct{}{},
			"woman":                 struct{}{},
			"woman lifting weights": struct{}{},
		},
	},
	":woman_lifting_weights_tone2:": Code{
		"\U0001f3cb\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"weight lifter":          struct{}{},
			"woman":                  struct{}{},
			"woman lifting weights":  struct{}{},
		},
	},
	":woman_lifting_weights_tone3:": Code{
		"\U0001f3cb\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium skin tone":      struct{}{},
			"weight lifter":         struct{}{},
			"woman":                 struct{}{},
			"woman lifting weights": struct{}{},
		},
	},
	":woman_lifting_weights_tone4:": Code{
		"\U0001f3cb\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"weight lifter":         struct{}{},
			"woman":                 struct{}{},
			"woman lifting weights": struct{}{},
		},
	},
	":woman_lifting_weights_tone5:": Code{
		"\U0001f3cb\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":        struct{}{},
			"weight lifter":         struct{}{},
			"woman":                 struct{}{},
			"woman lifting weights": struct{}{},
		},
	},
	":person_biking:": Code{
		"\U0001f6b4",
		map[string]struct{}{
			"bicycle":       struct{}{},
			"biking":        struct{}{},
			"cyclist":       struct{}{},
			"person biking": struct{}{},
		},
	},
	":person_biking_tone1:": Code{
		"\U0001f6b4\U0001f3fb",
		map[string]struct{}{
			"bicycle":         struct{}{},
			"biking":          struct{}{},
			"cyclist":         struct{}{},
			"light skin tone": struct{}{},
			"person biking":   struct{}{},
		},
	},
	":person_biking_tone2:": Code{
		"\U0001f6b4\U0001f3fc",
		map[string]struct{}{
			"bicycle":                struct{}{},
			"biking":                 struct{}{},
			"cyclist":                struct{}{},
			"medium-light skin tone": struct{}{},
			"person biking":          struct{}{},
		},
	},
	":person_biking_tone3:": Code{
		"\U0001f6b4\U0001f3fd",
		map[string]struct{}{
			"bicycle":          struct{}{},
			"biking":           struct{}{},
			"cyclist":          struct{}{},
			"medium skin tone": struct{}{},
			"person biking":    struct{}{},
		},
	},
	":person_biking_tone4:": Code{
		"\U0001f6b4\U0001f3fe",
		map[string]struct{}{
			"bicycle":               struct{}{},
			"biking":                struct{}{},
			"cyclist":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"person biking":         struct{}{},
		},
	},
	":person_biking_tone5:": Code{
		"\U0001f6b4\U0001f3ff",
		map[string]struct{}{
			"bicycle":        struct{}{},
			"biking":         struct{}{},
			"cyclist":        struct{}{},
			"dark skin tone": struct{}{},
			"person biking":  struct{}{},
		},
	},
	":man_biking:": Code{
		"\U0001f6b4\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bicycle":    struct{}{},
			"biking":     struct{}{},
			"cyclist":    struct{}{},
			"man":        struct{}{},
			"man biking": struct{}{},
		},
	},
	":man_biking_tone1:": Code{
		"\U0001f6b4\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bicycle":         struct{}{},
			"biking":          struct{}{},
			"cyclist":         struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man biking":      struct{}{},
		},
	},
	":man_biking_tone2:": Code{
		"\U0001f6b4\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bicycle":                struct{}{},
			"biking":                 struct{}{},
			"cyclist":                struct{}{},
			"man":                    struct{}{},
			"man biking":             struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_biking_tone3:": Code{
		"\U0001f6b4\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bicycle":          struct{}{},
			"biking":           struct{}{},
			"cyclist":          struct{}{},
			"man":              struct{}{},
			"man biking":       struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":man_biking_tone4:": Code{
		"\U0001f6b4\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bicycle":               struct{}{},
			"biking":                struct{}{},
			"cyclist":               struct{}{},
			"man":                   struct{}{},
			"man biking":            struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_biking_tone5:": Code{
		"\U0001f6b4\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bicycle":        struct{}{},
			"biking":         struct{}{},
			"cyclist":        struct{}{},
			"dark skin tone": struct{}{},
			"man":            struct{}{},
			"man biking":     struct{}{},
		},
	},
	":woman_biking:": Code{
		"\U0001f6b4\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bicycle":      struct{}{},
			"biking":       struct{}{},
			"cyclist":      struct{}{},
			"woman":        struct{}{},
			"woman biking": struct{}{},
		},
	},
	":woman_biking_tone1:": Code{
		"\U0001f6b4\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bicycle":         struct{}{},
			"biking":          struct{}{},
			"cyclist":         struct{}{},
			"light skin tone": struct{}{},
			"woman":           struct{}{},
			"woman biking":    struct{}{},
		},
	},
	":woman_biking_tone2:": Code{
		"\U0001f6b4\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bicycle":                struct{}{},
			"biking":                 struct{}{},
			"cyclist":                struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
			"woman biking":           struct{}{},
		},
	},
	":woman_biking_tone3:": Code{
		"\U0001f6b4\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bicycle":          struct{}{},
			"biking":           struct{}{},
			"cyclist":          struct{}{},
			"medium skin tone": struct{}{},
			"woman":            struct{}{},
			"woman biking":     struct{}{},
		},
	},
	":woman_biking_tone4:": Code{
		"\U0001f6b4\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bicycle":               struct{}{},
			"biking":                struct{}{},
			"cyclist":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
			"woman biking":          struct{}{},
		},
	},
	":woman_biking_tone5:": Code{
		"\U0001f6b4\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bicycle":        struct{}{},
			"biking":         struct{}{},
			"cyclist":        struct{}{},
			"dark skin tone": struct{}{},
			"woman":          struct{}{},
			"woman biking":   struct{}{},
		},
	},
	":person_mountain_biking:": Code{
		"\U0001f6b5",
		map[string]struct{}{
			"bicycle":                struct{}{},
			"bicyclist":              struct{}{},
			"bike":                   struct{}{},
			"cyclist":                struct{}{},
			"mountain":               struct{}{},
			"person mountain biking": struct{}{},
		},
	},
	":person_mountain_biking_tone1:": Code{
		"\U0001f6b5\U0001f3fb",
		map[string]struct{}{
			"bicycle":                struct{}{},
			"bicyclist":              struct{}{},
			"bike":                   struct{}{},
			"cyclist":                struct{}{},
			"light skin tone":        struct{}{},
			"mountain":               struct{}{},
			"person mountain biking": struct{}{},
		},
	},
	":person_mountain_biking_tone2:": Code{
		"\U0001f6b5\U0001f3fc",
		map[string]struct{}{
			"bicycle":                struct{}{},
			"bicyclist":              struct{}{},
			"bike":                   struct{}{},
			"cyclist":                struct{}{},
			"medium-light skin tone": struct{}{},
			"mountain":               struct{}{},
			"person mountain biking": struct{}{},
		},
	},
	":person_mountain_biking_tone3:": Code{
		"\U0001f6b5\U0001f3fd",
		map[string]struct{}{
			"bicycle":                struct{}{},
			"bicyclist":              struct{}{},
			"bike":                   struct{}{},
			"cyclist":                struct{}{},
			"medium skin tone":       struct{}{},
			"mountain":               struct{}{},
			"person mountain biking": struct{}{},
		},
	},
	":person_mountain_biking_tone4:": Code{
		"\U0001f6b5\U0001f3fe",
		map[string]struct{}{
			"bicycle":                struct{}{},
			"bicyclist":              struct{}{},
			"bike":                   struct{}{},
			"cyclist":                struct{}{},
			"medium-dark skin tone":  struct{}{},
			"mountain":               struct{}{},
			"person mountain biking": struct{}{},
		},
	},
	":person_mountain_biking_tone5:": Code{
		"\U0001f6b5\U0001f3ff",
		map[string]struct{}{
			"bicycle":                struct{}{},
			"bicyclist":              struct{}{},
			"bike":                   struct{}{},
			"cyclist":                struct{}{},
			"dark skin tone":         struct{}{},
			"mountain":               struct{}{},
			"person mountain biking": struct{}{},
		},
	},
	":man_mountain_biking:": Code{
		"\U0001f6b5\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bicycle":             struct{}{},
			"bike":                struct{}{},
			"cyclist":             struct{}{},
			"man":                 struct{}{},
			"man mountain biking": struct{}{},
			"mountain":            struct{}{},
		},
	},
	":man_mountain_biking_tone1:": Code{
		"\U0001f6b5\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bicycle":             struct{}{},
			"bike":                struct{}{},
			"cyclist":             struct{}{},
			"light skin tone":     struct{}{},
			"man":                 struct{}{},
			"man mountain biking": struct{}{},
			"mountain":            struct{}{},
		},
	},
	":man_mountain_biking_tone2:": Code{
		"\U0001f6b5\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bicycle":                struct{}{},
			"bike":                   struct{}{},
			"cyclist":                struct{}{},
			"man":                    struct{}{},
			"man mountain biking":    struct{}{},
			"medium-light skin tone": struct{}{},
			"mountain":               struct{}{},
		},
	},
	":man_mountain_biking_tone3:": Code{
		"\U0001f6b5\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bicycle":             struct{}{},
			"bike":                struct{}{},
			"cyclist":             struct{}{},
			"man":                 struct{}{},
			"man mountain biking": struct{}{},
			"medium skin tone":    struct{}{},
			"mountain":            struct{}{},
		},
	},
	":man_mountain_biking_tone4:": Code{
		"\U0001f6b5\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bicycle":               struct{}{},
			"bike":                  struct{}{},
			"cyclist":               struct{}{},
			"man":                   struct{}{},
			"man mountain biking":   struct{}{},
			"medium-dark skin tone": struct{}{},
			"mountain":              struct{}{},
		},
	},
	":man_mountain_biking_tone5:": Code{
		"\U0001f6b5\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bicycle":             struct{}{},
			"bike":                struct{}{},
			"cyclist":             struct{}{},
			"dark skin tone":      struct{}{},
			"man":                 struct{}{},
			"man mountain biking": struct{}{},
			"mountain":            struct{}{},
		},
	},
	":woman_mountain_biking:": Code{
		"\U0001f6b5\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bicycle":               struct{}{},
			"bike":                  struct{}{},
			"biking":                struct{}{},
			"cyclist":               struct{}{},
			"mountain":              struct{}{},
			"woman":                 struct{}{},
			"woman mountain biking": struct{}{},
		},
	},
	":woman_mountain_biking_tone1:": Code{
		"\U0001f6b5\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bicycle":               struct{}{},
			"bike":                  struct{}{},
			"biking":                struct{}{},
			"cyclist":               struct{}{},
			"light skin tone":       struct{}{},
			"mountain":              struct{}{},
			"woman":                 struct{}{},
			"woman mountain biking": struct{}{},
		},
	},
	":woman_mountain_biking_tone2:": Code{
		"\U0001f6b5\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bicycle":                struct{}{},
			"bike":                   struct{}{},
			"biking":                 struct{}{},
			"cyclist":                struct{}{},
			"medium-light skin tone": struct{}{},
			"mountain":               struct{}{},
			"woman":                  struct{}{},
			"woman mountain biking":  struct{}{},
		},
	},
	":woman_mountain_biking_tone3:": Code{
		"\U0001f6b5\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bicycle":               struct{}{},
			"bike":                  struct{}{},
			"biking":                struct{}{},
			"cyclist":               struct{}{},
			"medium skin tone":      struct{}{},
			"mountain":              struct{}{},
			"woman":                 struct{}{},
			"woman mountain biking": struct{}{},
		},
	},
	":woman_mountain_biking_tone4:": Code{
		"\U0001f6b5\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bicycle":               struct{}{},
			"bike":                  struct{}{},
			"biking":                struct{}{},
			"cyclist":               struct{}{},
			"medium-dark skin tone": struct{}{},
			"mountain":              struct{}{},
			"woman":                 struct{}{},
			"woman mountain biking": struct{}{},
		},
	},
	":woman_mountain_biking_tone5:": Code{
		"\U0001f6b5\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bicycle":               struct{}{},
			"bike":                  struct{}{},
			"biking":                struct{}{},
			"cyclist":               struct{}{},
			"dark skin tone":        struct{}{},
			"mountain":              struct{}{},
			"woman":                 struct{}{},
			"woman mountain biking": struct{}{},
		},
	},
	":race_car:": Code{
		"\U0001f3ce",
		map[string]struct{}{
			"car":        struct{}{},
			"racing":     struct{}{},
			"racing car": struct{}{},
		},
	},
	":motorcycle:": Code{
		"\U0001f3cd",
		map[string]struct{}{
			"motorcycle": struct{}{},
			"racing":     struct{}{},
		},
	},
	":person_doing_cartwheel:": Code{
		"\U0001f938",
		map[string]struct{}{
			"cartwheel":           struct{}{},
			"gymnastics":          struct{}{},
			"person cartwheeling": struct{}{},
		},
	},
	":person_doing_cartwheel_tone1:": Code{
		"\U0001f938\U0001f3fb",
		map[string]struct{}{
			"cartwheel":           struct{}{},
			"gymnastics":          struct{}{},
			"light skin tone":     struct{}{},
			"person cartwheeling": struct{}{},
		},
	},
	":person_doing_cartwheel_tone2:": Code{
		"\U0001f938\U0001f3fc",
		map[string]struct{}{
			"cartwheel":              struct{}{},
			"gymnastics":             struct{}{},
			"medium-light skin tone": struct{}{},
			"person cartwheeling":    struct{}{},
		},
	},
	":person_doing_cartwheel_tone3:": Code{
		"\U0001f938\U0001f3fd",
		map[string]struct{}{
			"cartwheel":           struct{}{},
			"gymnastics":          struct{}{},
			"medium skin tone":    struct{}{},
			"person cartwheeling": struct{}{},
		},
	},
	":person_doing_cartwheel_tone4:": Code{
		"\U0001f938\U0001f3fe",
		map[string]struct{}{
			"cartwheel":             struct{}{},
			"gymnastics":            struct{}{},
			"medium-dark skin tone": struct{}{},
			"person cartwheeling":   struct{}{},
		},
	},
	":person_doing_cartwheel_tone5:": Code{
		"\U0001f938\U0001f3ff",
		map[string]struct{}{
			"cartwheel":           struct{}{},
			"dark skin tone":      struct{}{},
			"gymnastics":          struct{}{},
			"person cartwheeling": struct{}{},
		},
	},
	":man_cartwheeling:": Code{
		"\U0001f938\u200d\u2642\ufe0f",
		map[string]struct{}{
			"cartwheel":        struct{}{},
			"gymnastics":       struct{}{},
			"man":              struct{}{},
			"man cartwheeling": struct{}{},
		},
	},
	":man_cartwheeling_tone1:": Code{
		"\U0001f938\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"cartwheel":        struct{}{},
			"gymnastics":       struct{}{},
			"light skin tone":  struct{}{},
			"man":              struct{}{},
			"man cartwheeling": struct{}{},
		},
	},
	":man_cartwheeling_tone2:": Code{
		"\U0001f938\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"cartwheel":              struct{}{},
			"gymnastics":             struct{}{},
			"man":                    struct{}{},
			"man cartwheeling":       struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_cartwheeling_tone3:": Code{
		"\U0001f938\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"cartwheel":        struct{}{},
			"gymnastics":       struct{}{},
			"man":              struct{}{},
			"man cartwheeling": struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":man_cartwheeling_tone4:": Code{
		"\U0001f938\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"cartwheel":             struct{}{},
			"gymnastics":            struct{}{},
			"man":                   struct{}{},
			"man cartwheeling":      struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_cartwheeling_tone5:": Code{
		"\U0001f938\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"cartwheel":        struct{}{},
			"dark skin tone":   struct{}{},
			"gymnastics":       struct{}{},
			"man":              struct{}{},
			"man cartwheeling": struct{}{},
		},
	},
	":woman_cartwheeling:": Code{
		"\U0001f938\u200d\u2640\ufe0f",
		map[string]struct{}{
			"cartwheel":          struct{}{},
			"gymnastics":         struct{}{},
			"woman":              struct{}{},
			"woman cartwheeling": struct{}{},
		},
	},
	":woman_cartwheeling_tone1:": Code{
		"\U0001f938\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"cartwheel":          struct{}{},
			"gymnastics":         struct{}{},
			"light skin tone":    struct{}{},
			"woman":              struct{}{},
			"woman cartwheeling": struct{}{},
		},
	},
	":woman_cartwheeling_tone2:": Code{
		"\U0001f938\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"cartwheel":              struct{}{},
			"gymnastics":             struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
			"woman cartwheeling":     struct{}{},
		},
	},
	":woman_cartwheeling_tone3:": Code{
		"\U0001f938\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"cartwheel":          struct{}{},
			"gymnastics":         struct{}{},
			"medium skin tone":   struct{}{},
			"woman":              struct{}{},
			"woman cartwheeling": struct{}{},
		},
	},
	":woman_cartwheeling_tone4:": Code{
		"\U0001f938\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"cartwheel":             struct{}{},
			"gymnastics":            struct{}{},
			"medium-dark skin tone": struct{}{},
			"woman":                 struct{}{},
			"woman cartwheeling":    struct{}{},
		},
	},
	":woman_cartwheeling_tone5:": Code{
		"\U0001f938\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"cartwheel":          struct{}{},
			"dark skin tone":     struct{}{},
			"gymnastics":         struct{}{},
			"woman":              struct{}{},
			"woman cartwheeling": struct{}{},
		},
	},
	":people_wrestling:": Code{
		"\U0001f93c",
		map[string]struct{}{
			"people wrestling": struct{}{},
			"wrestle":          struct{}{},
			"wrestler":         struct{}{},
		},
	},
	":men_wrestling:": Code{
		"\U0001f93c\u200d\u2642\ufe0f",
		map[string]struct{}{
			"men":           struct{}{},
			"men wrestling": struct{}{},
			"wrestle":       struct{}{},
		},
	},
	":women_wrestling:": Code{
		"\U0001f93c\u200d\u2640\ufe0f",
		map[string]struct{}{
			"women":           struct{}{},
			"women wrestling": struct{}{},
			"wrestle":         struct{}{},
		},
	},
	":person_playing_water_polo:": Code{
		"\U0001f93d",
		map[string]struct{}{
			"person playing water polo": struct{}{},
			"polo":                      struct{}{},
			"water":                     struct{}{},
		},
	},
	":person_playing_water_polo_tone1:": Code{
		"\U0001f93d\U0001f3fb",
		map[string]struct{}{
			"light skin tone":           struct{}{},
			"person playing water polo": struct{}{},
			"polo":                      struct{}{},
			"water":                     struct{}{},
		},
	},
	":person_playing_water_polo_tone2:": Code{
		"\U0001f93d\U0001f3fc",
		map[string]struct{}{
			"medium-light skin tone":    struct{}{},
			"person playing water polo": struct{}{},
			"polo":                      struct{}{},
			"water":                     struct{}{},
		},
	},
	":person_playing_water_polo_tone3:": Code{
		"\U0001f93d\U0001f3fd",
		map[string]struct{}{
			"medium skin tone":          struct{}{},
			"person playing water polo": struct{}{},
			"polo":                      struct{}{},
			"water":                     struct{}{},
		},
	},
	":person_playing_water_polo_tone4:": Code{
		"\U0001f93d\U0001f3fe",
		map[string]struct{}{
			"medium-dark skin tone":     struct{}{},
			"person playing water polo": struct{}{},
			"polo":                      struct{}{},
			"water":                     struct{}{},
		},
	},
	":person_playing_water_polo_tone5:": Code{
		"\U0001f93d\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":            struct{}{},
			"person playing water polo": struct{}{},
			"polo":                      struct{}{},
			"water":                     struct{}{},
		},
	},
	":man_playing_water_polo:": Code{
		"\U0001f93d\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                    struct{}{},
			"man playing water polo": struct{}{},
			"water polo":             struct{}{},
		},
	},
	":man_playing_water_polo_tone1:": Code{
		"\U0001f93d\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"light skin tone":        struct{}{},
			"man":                    struct{}{},
			"man playing water polo": struct{}{},
			"water polo":             struct{}{},
		},
	},
	":man_playing_water_polo_tone2:": Code{
		"\U0001f93d\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                    struct{}{},
			"man playing water polo": struct{}{},
			"medium-light skin tone": struct{}{},
			"water polo":             struct{}{},
		},
	},
	":man_playing_water_polo_tone3:": Code{
		"\U0001f93d\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                    struct{}{},
			"man playing water polo": struct{}{},
			"medium skin tone":       struct{}{},
			"water polo":             struct{}{},
		},
	},
	":man_playing_water_polo_tone4:": Code{
		"\U0001f93d\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"man":                    struct{}{},
			"man playing water polo": struct{}{},
			"medium-dark skin tone":  struct{}{},
			"water polo":             struct{}{},
		},
	},
	":man_playing_water_polo_tone5:": Code{
		"\U0001f93d\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone":         struct{}{},
			"man":                    struct{}{},
			"man playing water polo": struct{}{},
			"water polo":             struct{}{},
		},
	},
	":woman_playing_water_polo:": Code{
		"\U0001f93d\u200d\u2640\ufe0f",
		map[string]struct{}{
			"water polo":               struct{}{},
			"woman":                    struct{}{},
			"woman playing water polo": struct{}{},
		},
	},
	":woman_playing_water_polo_tone1:": Code{
		"\U0001f93d\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"light skin tone":          struct{}{},
			"water polo":               struct{}{},
			"woman":                    struct{}{},
			"woman playing water polo": struct{}{},
		},
	},
	":woman_playing_water_polo_tone2:": Code{
		"\U0001f93d\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-light skin tone":   struct{}{},
			"water polo":               struct{}{},
			"woman":                    struct{}{},
			"woman playing water polo": struct{}{},
		},
	},
	":woman_playing_water_polo_tone3:": Code{
		"\U0001f93d\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium skin tone":         struct{}{},
			"water polo":               struct{}{},
			"woman":                    struct{}{},
			"woman playing water polo": struct{}{},
		},
	},
	":woman_playing_water_polo_tone4:": Code{
		"\U0001f93d\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"medium-dark skin tone":    struct{}{},
			"water polo":               struct{}{},
			"woman":                    struct{}{},
			"woman playing water polo": struct{}{},
		},
	},
	":woman_playing_water_polo_tone5:": Code{
		"\U0001f93d\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":           struct{}{},
			"water polo":               struct{}{},
			"woman":                    struct{}{},
			"woman playing water polo": struct{}{},
		},
	},
	":person_playing_handball:": Code{
		"\U0001f93e",
		map[string]struct{}{
			"ball":                    struct{}{},
			"handball":                struct{}{},
			"person playing handball": struct{}{},
		},
	},
	":person_playing_handball_tone1:": Code{
		"\U0001f93e\U0001f3fb",
		map[string]struct{}{
			"ball":                    struct{}{},
			"handball":                struct{}{},
			"light skin tone":         struct{}{},
			"person playing handball": struct{}{},
		},
	},
	":person_playing_handball_tone2:": Code{
		"\U0001f93e\U0001f3fc",
		map[string]struct{}{
			"ball":                    struct{}{},
			"handball":                struct{}{},
			"medium-light skin tone":  struct{}{},
			"person playing handball": struct{}{},
		},
	},
	":person_playing_handball_tone3:": Code{
		"\U0001f93e\U0001f3fd",
		map[string]struct{}{
			"ball":                    struct{}{},
			"handball":                struct{}{},
			"medium skin tone":        struct{}{},
			"person playing handball": struct{}{},
		},
	},
	":person_playing_handball_tone4:": Code{
		"\U0001f93e\U0001f3fe",
		map[string]struct{}{
			"ball":                    struct{}{},
			"handball":                struct{}{},
			"medium-dark skin tone":   struct{}{},
			"person playing handball": struct{}{},
		},
	},
	":person_playing_handball_tone5:": Code{
		"\U0001f93e\U0001f3ff",
		map[string]struct{}{
			"ball":                    struct{}{},
			"dark skin tone":          struct{}{},
			"handball":                struct{}{},
			"person playing handball": struct{}{},
		},
	},
	":man_playing_handball:": Code{
		"\U0001f93e\u200d\u2642\ufe0f",
		map[string]struct{}{
			"handball":             struct{}{},
			"man":                  struct{}{},
			"man playing handball": struct{}{},
		},
	},
	":man_playing_handball_tone1:": Code{
		"\U0001f93e\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"handball":             struct{}{},
			"light skin tone":      struct{}{},
			"man":                  struct{}{},
			"man playing handball": struct{}{},
		},
	},
	":man_playing_handball_tone2:": Code{
		"\U0001f93e\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"handball":               struct{}{},
			"man":                    struct{}{},
			"man playing handball":   struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":man_playing_handball_tone3:": Code{
		"\U0001f93e\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"handball":             struct{}{},
			"man":                  struct{}{},
			"man playing handball": struct{}{},
			"medium skin tone":     struct{}{},
		},
	},
	":man_playing_handball_tone4:": Code{
		"\U0001f93e\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"handball":              struct{}{},
			"man":                   struct{}{},
			"man playing handball":  struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":man_playing_handball_tone5:": Code{
		"\U0001f93e\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone":       struct{}{},
			"handball":             struct{}{},
			"man":                  struct{}{},
			"man playing handball": struct{}{},
		},
	},
	":woman_playing_handball:": Code{
		"\U0001f93e\u200d\u2640\ufe0f",
		map[string]struct{}{
			"handball":               struct{}{},
			"woman":                  struct{}{},
			"woman playing handball": struct{}{},
		},
	},
	":woman_playing_handball_tone1:": Code{
		"\U0001f93e\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"handball":               struct{}{},
			"light skin tone":        struct{}{},
			"woman":                  struct{}{},
			"woman playing handball": struct{}{},
		},
	},
	":woman_playing_handball_tone2:": Code{
		"\U0001f93e\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"handball":               struct{}{},
			"medium-light skin tone": struct{}{},
			"woman":                  struct{}{},
			"woman playing handball": struct{}{},
		},
	},
	":woman_playing_handball_tone3:": Code{
		"\U0001f93e\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"handball":               struct{}{},
			"medium skin tone":       struct{}{},
			"woman":                  struct{}{},
			"woman playing handball": struct{}{},
		},
	},
	":woman_playing_handball_tone4:": Code{
		"\U0001f93e\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"handball":               struct{}{},
			"medium-dark skin tone":  struct{}{},
			"woman":                  struct{}{},
			"woman playing handball": struct{}{},
		},
	},
	":woman_playing_handball_tone5:": Code{
		"\U0001f93e\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone":         struct{}{},
			"handball":               struct{}{},
			"woman":                  struct{}{},
			"woman playing handball": struct{}{},
		},
	},
	":person_juggling:": Code{
		"\U0001f939",
		map[string]struct{}{
			"balance":         struct{}{},
			"juggle":          struct{}{},
			"multitask":       struct{}{},
			"person juggling": struct{}{},
			"skill":           struct{}{},
		},
	},
	":person_juggling_tone1:": Code{
		"\U0001f939\U0001f3fb",
		map[string]struct{}{
			"balance":         struct{}{},
			"juggle":          struct{}{},
			"light skin tone": struct{}{},
			"multitask":       struct{}{},
			"person juggling": struct{}{},
			"skill":           struct{}{},
		},
	},
	":person_juggling_tone2:": Code{
		"\U0001f939\U0001f3fc",
		map[string]struct{}{
			"balance":                struct{}{},
			"juggle":                 struct{}{},
			"medium-light skin tone": struct{}{},
			"multitask":              struct{}{},
			"person juggling":        struct{}{},
			"skill":                  struct{}{},
		},
	},
	":person_juggling_tone3:": Code{
		"\U0001f939\U0001f3fd",
		map[string]struct{}{
			"balance":          struct{}{},
			"juggle":           struct{}{},
			"medium skin tone": struct{}{},
			"multitask":        struct{}{},
			"person juggling":  struct{}{},
			"skill":            struct{}{},
		},
	},
	":person_juggling_tone4:": Code{
		"\U0001f939\U0001f3fe",
		map[string]struct{}{
			"balance":               struct{}{},
			"juggle":                struct{}{},
			"medium-dark skin tone": struct{}{},
			"multitask":             struct{}{},
			"person juggling":       struct{}{},
			"skill":                 struct{}{},
		},
	},
	":person_juggling_tone5:": Code{
		"\U0001f939\U0001f3ff",
		map[string]struct{}{
			"balance":         struct{}{},
			"dark skin tone":  struct{}{},
			"juggle":          struct{}{},
			"multitask":       struct{}{},
			"person juggling": struct{}{},
			"skill":           struct{}{},
		},
	},
	":man_juggling:": Code{
		"\U0001f939\u200d\u2642\ufe0f",
		map[string]struct{}{
			"juggling":     struct{}{},
			"man":          struct{}{},
			"man juggling": struct{}{},
			"multitask":    struct{}{},
		},
	},
	":man_juggling_tone1:": Code{
		"\U0001f939\U0001f3fb\u200d\u2642\ufe0f",
		map[string]struct{}{
			"juggling":        struct{}{},
			"light skin tone": struct{}{},
			"man":             struct{}{},
			"man juggling":    struct{}{},
			"multitask":       struct{}{},
		},
	},
	":man_juggling_tone2:": Code{
		"\U0001f939\U0001f3fc\u200d\u2642\ufe0f",
		map[string]struct{}{
			"juggling":               struct{}{},
			"man":                    struct{}{},
			"man juggling":           struct{}{},
			"medium-light skin tone": struct{}{},
			"multitask":              struct{}{},
		},
	},
	":man_juggling_tone3:": Code{
		"\U0001f939\U0001f3fd\u200d\u2642\ufe0f",
		map[string]struct{}{
			"juggling":         struct{}{},
			"man":              struct{}{},
			"man juggling":     struct{}{},
			"medium skin tone": struct{}{},
			"multitask":        struct{}{},
		},
	},
	":man_juggling_tone4:": Code{
		"\U0001f939\U0001f3fe\u200d\u2642\ufe0f",
		map[string]struct{}{
			"juggling":              struct{}{},
			"man":                   struct{}{},
			"man juggling":          struct{}{},
			"medium-dark skin tone": struct{}{},
			"multitask":             struct{}{},
		},
	},
	":man_juggling_tone5:": Code{
		"\U0001f939\U0001f3ff\u200d\u2642\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"juggling":       struct{}{},
			"man":            struct{}{},
			"man juggling":   struct{}{},
			"multitask":      struct{}{},
		},
	},
	":woman_juggling:": Code{
		"\U0001f939\u200d\u2640\ufe0f",
		map[string]struct{}{
			"juggling":       struct{}{},
			"multitask":      struct{}{},
			"woman":          struct{}{},
			"woman juggling": struct{}{},
		},
	},
	":woman_juggling_tone1:": Code{
		"\U0001f939\U0001f3fb\u200d\u2640\ufe0f",
		map[string]struct{}{
			"juggling":        struct{}{},
			"light skin tone": struct{}{},
			"multitask":       struct{}{},
			"woman":           struct{}{},
			"woman juggling":  struct{}{},
		},
	},
	":woman_juggling_tone2:": Code{
		"\U0001f939\U0001f3fc\u200d\u2640\ufe0f",
		map[string]struct{}{
			"juggling":               struct{}{},
			"medium-light skin tone": struct{}{},
			"multitask":              struct{}{},
			"woman":                  struct{}{},
			"woman juggling":         struct{}{},
		},
	},
	":woman_juggling_tone3:": Code{
		"\U0001f939\U0001f3fd\u200d\u2640\ufe0f",
		map[string]struct{}{
			"juggling":         struct{}{},
			"medium skin tone": struct{}{},
			"multitask":        struct{}{},
			"woman":            struct{}{},
			"woman juggling":   struct{}{},
		},
	},
	":woman_juggling_tone4:": Code{
		"\U0001f939\U0001f3fe\u200d\u2640\ufe0f",
		map[string]struct{}{
			"juggling":              struct{}{},
			"medium-dark skin tone": struct{}{},
			"multitask":             struct{}{},
			"woman":                 struct{}{},
			"woman juggling":        struct{}{},
		},
	},
	":woman_juggling_tone5:": Code{
		"\U0001f939\U0001f3ff\u200d\u2640\ufe0f",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"juggling":       struct{}{},
			"multitask":      struct{}{},
			"woman":          struct{}{},
			"woman juggling": struct{}{},
		},
	},
	":couple:": Code{
		"\U0001f46b",
		map[string]struct{}{
			"couple":                      struct{}{},
			"hand":                        struct{}{},
			"hold":                        struct{}{},
			"man":                         struct{}{},
			"man and woman holding hands": struct{}{},
			"woman":                       struct{}{},
		},
	},
	":two_men_holding_hands:": Code{
		"\U0001f46c",
		map[string]struct{}{
			"couple":                struct{}{},
			"Gemini":                struct{}{},
			"hand":                  struct{}{},
			"hold":                  struct{}{},
			"man":                   struct{}{},
			"twins":                 struct{}{},
			"two men holding hands": struct{}{},
			"zodiac":                struct{}{},
		},
	},
	":two_women_holding_hands:": Code{
		"\U0001f46d",
		map[string]struct{}{
			"couple":                  struct{}{},
			"hand":                    struct{}{},
			"hold":                    struct{}{},
			"two women holding hands": struct{}{},
			"woman":                   struct{}{},
		},
	},
	":couplekiss:": Code{
		"\U0001f48f",
		map[string]struct{}{
			"couple": struct{}{},
			"kiss":   struct{}{},
		},
	},
	":kiss_woman_man:": Code{
		"\U0001f469\u200d\u2764\ufe0f\u200d\U0001f48b\u200d\U0001f468",
		map[string]struct{}{
			"couple": struct{}{},
			"kiss":   struct{}{},
			"man":    struct{}{},
			"woman":  struct{}{},
		},
	},
	":kiss_mm:": Code{
		"\U0001f468\u200d\u2764\ufe0f\u200d\U0001f48b\u200d\U0001f468",
		map[string]struct{}{
			"couple": struct{}{},
			"kiss":   struct{}{},
			"man":    struct{}{},
		},
	},
	":kiss_ww:": Code{
		"\U0001f469\u200d\u2764\ufe0f\u200d\U0001f48b\u200d\U0001f469",
		map[string]struct{}{
			"couple": struct{}{},
			"kiss":   struct{}{},
			"woman":  struct{}{},
		},
	},
	":couple_with_heart:": Code{
		"\U0001f491",
		map[string]struct{}{
			"couple":            struct{}{},
			"couple with heart": struct{}{},
			"love":              struct{}{},
		},
	},
	":couple_with_heart_woman_man:": Code{
		"\U0001f469\u200d\u2764\ufe0f\u200d\U0001f468",
		map[string]struct{}{
			"couple":            struct{}{},
			"couple with heart": struct{}{},
			"love":              struct{}{},
			"man":               struct{}{},
			"woman":             struct{}{},
		},
	},
	":couple_mm:": Code{
		"\U0001f468\u200d\u2764\ufe0f\u200d\U0001f468",
		map[string]struct{}{
			"couple":            struct{}{},
			"couple with heart": struct{}{},
			"love":              struct{}{},
			"man":               struct{}{},
		},
	},
	":couple_ww:": Code{
		"\U0001f469\u200d\u2764\ufe0f\u200d\U0001f469",
		map[string]struct{}{
			"couple":            struct{}{},
			"couple with heart": struct{}{},
			"love":              struct{}{},
			"woman":             struct{}{},
		},
	},
	":family:": Code{
		"\U0001f46a",
		map[string]struct{}{
			"family": struct{}{},
		},
	},
	":family_man_woman_boy:": Code{
		"\U0001f468\u200d\U0001f469\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"man":    struct{}{},
			"woman":  struct{}{},
		},
	},
	":family_mwg:": Code{
		"\U0001f468\u200d\U0001f469\u200d\U0001f467",
		map[string]struct{}{
			"family": struct{}{},
			"girl":   struct{}{},
			"man":    struct{}{},
			"woman":  struct{}{},
		},
	},
	":family_mwgb:": Code{
		"\U0001f468\u200d\U0001f469\u200d\U0001f467\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"girl":   struct{}{},
			"man":    struct{}{},
			"woman":  struct{}{},
		},
	},
	":family_mwbb:": Code{
		"\U0001f468\u200d\U0001f469\u200d\U0001f466\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"man":    struct{}{},
			"woman":  struct{}{},
		},
	},
	":family_mwgg:": Code{
		"\U0001f468\u200d\U0001f469\u200d\U0001f467\u200d\U0001f467",
		map[string]struct{}{
			"family": struct{}{},
			"girl":   struct{}{},
			"man":    struct{}{},
			"woman":  struct{}{},
		},
	},
	":family_mmb:": Code{
		"\U0001f468\u200d\U0001f468\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"man":    struct{}{},
		},
	},
	":family_mmg:": Code{
		"\U0001f468\u200d\U0001f468\u200d\U0001f467",
		map[string]struct{}{
			"family": struct{}{},
			"girl":   struct{}{},
			"man":    struct{}{},
		},
	},
	":family_mmgb:": Code{
		"\U0001f468\u200d\U0001f468\u200d\U0001f467\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"girl":   struct{}{},
			"man":    struct{}{},
		},
	},
	":family_mmbb:": Code{
		"\U0001f468\u200d\U0001f468\u200d\U0001f466\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"man":    struct{}{},
		},
	},
	":family_mmgg:": Code{
		"\U0001f468\u200d\U0001f468\u200d\U0001f467\u200d\U0001f467",
		map[string]struct{}{
			"family": struct{}{},
			"girl":   struct{}{},
			"man":    struct{}{},
		},
	},
	":family_wwb:": Code{
		"\U0001f469\u200d\U0001f469\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"woman":  struct{}{},
		},
	},
	":family_wwg:": Code{
		"\U0001f469\u200d\U0001f469\u200d\U0001f467",
		map[string]struct{}{
			"family": struct{}{},
			"girl":   struct{}{},
			"woman":  struct{}{},
		},
	},
	":family_wwgb:": Code{
		"\U0001f469\u200d\U0001f469\u200d\U0001f467\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"girl":   struct{}{},
			"woman":  struct{}{},
		},
	},
	":family_wwbb:": Code{
		"\U0001f469\u200d\U0001f469\u200d\U0001f466\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"woman":  struct{}{},
		},
	},
	":family_wwgg:": Code{
		"\U0001f469\u200d\U0001f469\u200d\U0001f467\u200d\U0001f467",
		map[string]struct{}{
			"family": struct{}{},
			"girl":   struct{}{},
			"woman":  struct{}{},
		},
	},
	":family_man_boy:": Code{
		"\U0001f468\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"man":    struct{}{},
		},
	},
	":family_man_boy_boy:": Code{
		"\U0001f468\u200d\U0001f466\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"man":    struct{}{},
		},
	},
	":family_man_girl:": Code{
		"\U0001f468\u200d\U0001f467",
		map[string]struct{}{
			"family": struct{}{},
			"girl":   struct{}{},
			"man":    struct{}{},
		},
	},
	":family_man_girl_boy:": Code{
		"\U0001f468\u200d\U0001f467\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"girl":   struct{}{},
			"man":    struct{}{},
		},
	},
	":family_man_girl_girl:": Code{
		"\U0001f468\u200d\U0001f467\u200d\U0001f467",
		map[string]struct{}{
			"family": struct{}{},
			"girl":   struct{}{},
			"man":    struct{}{},
		},
	},
	":family_woman_boy:": Code{
		"\U0001f469\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"woman":  struct{}{},
		},
	},
	":family_woman_boy_boy:": Code{
		"\U0001f469\u200d\U0001f466\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"woman":  struct{}{},
		},
	},
	":family_woman_girl:": Code{
		"\U0001f469\u200d\U0001f467",
		map[string]struct{}{
			"family": struct{}{},
			"girl":   struct{}{},
			"woman":  struct{}{},
		},
	},
	":family_woman_girl_boy:": Code{
		"\U0001f469\u200d\U0001f467\u200d\U0001f466",
		map[string]struct{}{
			"boy":    struct{}{},
			"family": struct{}{},
			"girl":   struct{}{},
			"woman":  struct{}{},
		},
	},
	":family_woman_girl_girl:": Code{
		"\U0001f469\u200d\U0001f467\u200d\U0001f467",
		map[string]struct{}{
			"family": struct{}{},
			"girl":   struct{}{},
			"woman":  struct{}{},
		},
	},
	":selfie:": Code{
		"\U0001f933",
		map[string]struct{}{
			"camera": struct{}{},
			"phone":  struct{}{},
			"selfie": struct{}{},
		},
	},
	":selfie_tone1:": Code{
		"\U0001f933\U0001f3fb",
		map[string]struct{}{
			"camera":          struct{}{},
			"light skin tone": struct{}{},
			"phone":           struct{}{},
			"selfie":          struct{}{},
		},
	},
	":selfie_tone2:": Code{
		"\U0001f933\U0001f3fc",
		map[string]struct{}{
			"camera":                 struct{}{},
			"medium-light skin tone": struct{}{},
			"phone":                  struct{}{},
			"selfie":                 struct{}{},
		},
	},
	":selfie_tone3:": Code{
		"\U0001f933\U0001f3fd",
		map[string]struct{}{
			"camera":           struct{}{},
			"medium skin tone": struct{}{},
			"phone":            struct{}{},
			"selfie":           struct{}{},
		},
	},
	":selfie_tone4:": Code{
		"\U0001f933\U0001f3fe",
		map[string]struct{}{
			"camera":                struct{}{},
			"medium-dark skin tone": struct{}{},
			"phone":                 struct{}{},
			"selfie":                struct{}{},
		},
	},
	":selfie_tone5:": Code{
		"\U0001f933\U0001f3ff",
		map[string]struct{}{
			"camera":         struct{}{},
			"dark skin tone": struct{}{},
			"phone":          struct{}{},
			"selfie":         struct{}{},
		},
	},
	":muscle:": Code{
		"\U0001f4aa",
		map[string]struct{}{
			"biceps":        struct{}{},
			"comic":         struct{}{},
			"flex":          struct{}{},
			"flexed biceps": struct{}{},
			"muscle":        struct{}{},
		},
	},
	":muscle_tone1:": Code{
		"\U0001f4aa\U0001f3fb",
		map[string]struct{}{
			"biceps":          struct{}{},
			"comic":           struct{}{},
			"flex":            struct{}{},
			"flexed biceps":   struct{}{},
			"light skin tone": struct{}{},
			"muscle":          struct{}{},
		},
	},
	":muscle_tone2:": Code{
		"\U0001f4aa\U0001f3fc",
		map[string]struct{}{
			"biceps":                 struct{}{},
			"comic":                  struct{}{},
			"flex":                   struct{}{},
			"flexed biceps":          struct{}{},
			"medium-light skin tone": struct{}{},
			"muscle":                 struct{}{},
		},
	},
	":muscle_tone3:": Code{
		"\U0001f4aa\U0001f3fd",
		map[string]struct{}{
			"biceps":           struct{}{},
			"comic":            struct{}{},
			"flex":             struct{}{},
			"flexed biceps":    struct{}{},
			"medium skin tone": struct{}{},
			"muscle":           struct{}{},
		},
	},
	":muscle_tone4:": Code{
		"\U0001f4aa\U0001f3fe",
		map[string]struct{}{
			"biceps":                struct{}{},
			"comic":                 struct{}{},
			"flex":                  struct{}{},
			"flexed biceps":         struct{}{},
			"medium-dark skin tone": struct{}{},
			"muscle":                struct{}{},
		},
	},
	":muscle_tone5:": Code{
		"\U0001f4aa\U0001f3ff",
		map[string]struct{}{
			"biceps":         struct{}{},
			"comic":          struct{}{},
			"dark skin tone": struct{}{},
			"flex":           struct{}{},
			"flexed biceps":  struct{}{},
			"muscle":         struct{}{},
		},
	},
	":point_left:": Code{
		"\U0001f448",
		map[string]struct{}{
			"backhand":                     struct{}{},
			"backhand index pointing left": struct{}{},
			"finger":                       struct{}{},
			"hand":                         struct{}{},
			"index":                        struct{}{},
			"point":                        struct{}{},
		},
	},
	":point_left_tone1:": Code{
		"\U0001f448\U0001f3fb",
		map[string]struct{}{
			"backhand":                     struct{}{},
			"backhand index pointing left": struct{}{},
			"finger":                       struct{}{},
			"hand":                         struct{}{},
			"index":                        struct{}{},
			"light skin tone":              struct{}{},
			"point":                        struct{}{},
		},
	},
	":point_left_tone2:": Code{
		"\U0001f448\U0001f3fc",
		map[string]struct{}{
			"backhand":                     struct{}{},
			"backhand index pointing left": struct{}{},
			"finger":                       struct{}{},
			"hand":                         struct{}{},
			"index":                        struct{}{},
			"medium-light skin tone":       struct{}{},
			"point":                        struct{}{},
		},
	},
	":point_left_tone3:": Code{
		"\U0001f448\U0001f3fd",
		map[string]struct{}{
			"backhand":                     struct{}{},
			"backhand index pointing left": struct{}{},
			"finger":                       struct{}{},
			"hand":                         struct{}{},
			"index":                        struct{}{},
			"medium skin tone":             struct{}{},
			"point":                        struct{}{},
		},
	},
	":point_left_tone4:": Code{
		"\U0001f448\U0001f3fe",
		map[string]struct{}{
			"backhand":                     struct{}{},
			"backhand index pointing left": struct{}{},
			"finger":                       struct{}{},
			"hand":                         struct{}{},
			"index":                        struct{}{},
			"medium-dark skin tone":        struct{}{},
			"point":                        struct{}{},
		},
	},
	":point_left_tone5:": Code{
		"\U0001f448\U0001f3ff",
		map[string]struct{}{
			"backhand":                     struct{}{},
			"backhand index pointing left": struct{}{},
			"dark skin tone":               struct{}{},
			"finger":                       struct{}{},
			"hand":                         struct{}{},
			"index":                        struct{}{},
			"point":                        struct{}{},
		},
	},
	":point_right:": Code{
		"\U0001f449",
		map[string]struct{}{
			"backhand":                      struct{}{},
			"backhand index pointing right": struct{}{},
			"finger":                        struct{}{},
			"hand":                          struct{}{},
			"index":                         struct{}{},
			"point":                         struct{}{},
		},
	},
	":point_right_tone1:": Code{
		"\U0001f449\U0001f3fb",
		map[string]struct{}{
			"backhand":                      struct{}{},
			"backhand index pointing right": struct{}{},
			"finger":                        struct{}{},
			"hand":                          struct{}{},
			"index":                         struct{}{},
			"light skin tone":               struct{}{},
			"point":                         struct{}{},
		},
	},
	":point_right_tone2:": Code{
		"\U0001f449\U0001f3fc",
		map[string]struct{}{
			"backhand":                      struct{}{},
			"backhand index pointing right": struct{}{},
			"finger":                        struct{}{},
			"hand":                          struct{}{},
			"index":                         struct{}{},
			"medium-light skin tone":        struct{}{},
			"point":                         struct{}{},
		},
	},
	":point_right_tone3:": Code{
		"\U0001f449\U0001f3fd",
		map[string]struct{}{
			"backhand":                      struct{}{},
			"backhand index pointing right": struct{}{},
			"finger":                        struct{}{},
			"hand":                          struct{}{},
			"index":                         struct{}{},
			"medium skin tone":              struct{}{},
			"point":                         struct{}{},
		},
	},
	":point_right_tone4:": Code{
		"\U0001f449\U0001f3fe",
		map[string]struct{}{
			"backhand":                      struct{}{},
			"backhand index pointing right": struct{}{},
			"finger":                        struct{}{},
			"hand":                          struct{}{},
			"index":                         struct{}{},
			"medium-dark skin tone":         struct{}{},
			"point":                         struct{}{},
		},
	},
	":point_right_tone5:": Code{
		"\U0001f449\U0001f3ff",
		map[string]struct{}{
			"backhand":                      struct{}{},
			"backhand index pointing right": struct{}{},
			"dark skin tone":                struct{}{},
			"finger":                        struct{}{},
			"hand":                          struct{}{},
			"index":                         struct{}{},
			"point":                         struct{}{},
		},
	},
	":point_up:": Code{
		"\u261d",
		map[string]struct{}{
			"finger":            struct{}{},
			"hand":              struct{}{},
			"index":             struct{}{},
			"index pointing up": struct{}{},
			"point":             struct{}{},
			"up":                struct{}{},
		},
	},
	":point_up_tone1:": Code{
		"\u261d\U0001f3fb",
		map[string]struct{}{
			"finger":            struct{}{},
			"hand":              struct{}{},
			"index":             struct{}{},
			"index pointing up": struct{}{},
			"light skin tone":   struct{}{},
			"point":             struct{}{},
			"up":                struct{}{},
		},
	},
	":point_up_tone2:": Code{
		"\u261d\U0001f3fc",
		map[string]struct{}{
			"finger":                 struct{}{},
			"hand":                   struct{}{},
			"index":                  struct{}{},
			"index pointing up":      struct{}{},
			"medium-light skin tone": struct{}{},
			"point":                  struct{}{},
			"up":                     struct{}{},
		},
	},
	":point_up_tone3:": Code{
		"\u261d\U0001f3fd",
		map[string]struct{}{
			"finger":            struct{}{},
			"hand":              struct{}{},
			"index":             struct{}{},
			"index pointing up": struct{}{},
			"medium skin tone":  struct{}{},
			"point":             struct{}{},
			"up":                struct{}{},
		},
	},
	":point_up_tone4:": Code{
		"\u261d\U0001f3fe",
		map[string]struct{}{
			"finger":                struct{}{},
			"hand":                  struct{}{},
			"index":                 struct{}{},
			"index pointing up":     struct{}{},
			"medium-dark skin tone": struct{}{},
			"point":                 struct{}{},
			"up":                    struct{}{},
		},
	},
	":point_up_tone5:": Code{
		"\u261d\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":    struct{}{},
			"finger":            struct{}{},
			"hand":              struct{}{},
			"index":             struct{}{},
			"index pointing up": struct{}{},
			"point":             struct{}{},
			"up":                struct{}{},
		},
	},
	":point_up_2:": Code{
		"\U0001f446",
		map[string]struct{}{
			"backhand":                   struct{}{},
			"backhand index pointing up": struct{}{},
			"finger":                     struct{}{},
			"hand":                       struct{}{},
			"index":                      struct{}{},
			"point":                      struct{}{},
			"up":                         struct{}{},
		},
	},
	":point_up_2_tone1:": Code{
		"\U0001f446\U0001f3fb",
		map[string]struct{}{
			"backhand":                   struct{}{},
			"backhand index pointing up": struct{}{},
			"finger":                     struct{}{},
			"hand":                       struct{}{},
			"index":                      struct{}{},
			"light skin tone":            struct{}{},
			"point":                      struct{}{},
			"up":                         struct{}{},
		},
	},
	":point_up_2_tone2:": Code{
		"\U0001f446\U0001f3fc",
		map[string]struct{}{
			"backhand":                   struct{}{},
			"backhand index pointing up": struct{}{},
			"finger":                     struct{}{},
			"hand":                       struct{}{},
			"index":                      struct{}{},
			"medium-light skin tone":     struct{}{},
			"point":                      struct{}{},
			"up":                         struct{}{},
		},
	},
	":point_up_2_tone3:": Code{
		"\U0001f446\U0001f3fd",
		map[string]struct{}{
			"backhand":                   struct{}{},
			"backhand index pointing up": struct{}{},
			"finger":                     struct{}{},
			"hand":                       struct{}{},
			"index":                      struct{}{},
			"medium skin tone":           struct{}{},
			"point":                      struct{}{},
			"up":                         struct{}{},
		},
	},
	":point_up_2_tone4:": Code{
		"\U0001f446\U0001f3fe",
		map[string]struct{}{
			"backhand":                   struct{}{},
			"backhand index pointing up": struct{}{},
			"finger":                     struct{}{},
			"hand":                       struct{}{},
			"index":                      struct{}{},
			"medium-dark skin tone":      struct{}{},
			"point":                      struct{}{},
			"up":                         struct{}{},
		},
	},
	":point_up_2_tone5:": Code{
		"\U0001f446\U0001f3ff",
		map[string]struct{}{
			"backhand":                   struct{}{},
			"backhand index pointing up": struct{}{},
			"dark skin tone":             struct{}{},
			"finger":                     struct{}{},
			"hand":                       struct{}{},
			"index":                      struct{}{},
			"point":                      struct{}{},
			"up":                         struct{}{},
		},
	},
	":middle_finger:": Code{
		"\U0001f595",
		map[string]struct{}{
			"finger":        struct{}{},
			"hand":          struct{}{},
			"middle finger": struct{}{},
		},
	},
	":middle_finger_tone1:": Code{
		"\U0001f595\U0001f3fb",
		map[string]struct{}{
			"finger":          struct{}{},
			"hand":            struct{}{},
			"light skin tone": struct{}{},
			"middle finger":   struct{}{},
		},
	},
	":middle_finger_tone2:": Code{
		"\U0001f595\U0001f3fc",
		map[string]struct{}{
			"finger":                 struct{}{},
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"middle finger":          struct{}{},
		},
	},
	":middle_finger_tone3:": Code{
		"\U0001f595\U0001f3fd",
		map[string]struct{}{
			"finger":           struct{}{},
			"hand":             struct{}{},
			"medium skin tone": struct{}{},
			"middle finger":    struct{}{},
		},
	},
	":middle_finger_tone4:": Code{
		"\U0001f595\U0001f3fe",
		map[string]struct{}{
			"finger":                struct{}{},
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"middle finger":         struct{}{},
		},
	},
	":middle_finger_tone5:": Code{
		"\U0001f595\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"finger":         struct{}{},
			"hand":           struct{}{},
			"middle finger":  struct{}{},
		},
	},
	":point_down:": Code{
		"\U0001f447",
		map[string]struct{}{
			"backhand":                     struct{}{},
			"backhand index pointing down": struct{}{},
			"down":                         struct{}{},
			"finger":                       struct{}{},
			"hand":                         struct{}{},
			"index":                        struct{}{},
			"point":                        struct{}{},
		},
	},
	":point_down_tone1:": Code{
		"\U0001f447\U0001f3fb",
		map[string]struct{}{
			"backhand":                     struct{}{},
			"backhand index pointing down": struct{}{},
			"down":                         struct{}{},
			"finger":                       struct{}{},
			"hand":                         struct{}{},
			"index":                        struct{}{},
			"light skin tone":              struct{}{},
			"point":                        struct{}{},
		},
	},
	":point_down_tone2:": Code{
		"\U0001f447\U0001f3fc",
		map[string]struct{}{
			"backhand":                     struct{}{},
			"backhand index pointing down": struct{}{},
			"down":                         struct{}{},
			"finger":                       struct{}{},
			"hand":                         struct{}{},
			"index":                        struct{}{},
			"medium-light skin tone":       struct{}{},
			"point":                        struct{}{},
		},
	},
	":point_down_tone3:": Code{
		"\U0001f447\U0001f3fd",
		map[string]struct{}{
			"backhand":                     struct{}{},
			"backhand index pointing down": struct{}{},
			"down":                         struct{}{},
			"finger":                       struct{}{},
			"hand":                         struct{}{},
			"index":                        struct{}{},
			"medium skin tone":             struct{}{},
			"point":                        struct{}{},
		},
	},
	":point_down_tone4:": Code{
		"\U0001f447\U0001f3fe",
		map[string]struct{}{
			"backhand":                     struct{}{},
			"backhand index pointing down": struct{}{},
			"down":                         struct{}{},
			"finger":                       struct{}{},
			"hand":                         struct{}{},
			"index":                        struct{}{},
			"medium-dark skin tone":        struct{}{},
			"point":                        struct{}{},
		},
	},
	":point_down_tone5:": Code{
		"\U0001f447\U0001f3ff",
		map[string]struct{}{
			"backhand":                     struct{}{},
			"backhand index pointing down": struct{}{},
			"dark skin tone":               struct{}{},
			"down":                         struct{}{},
			"finger":                       struct{}{},
			"hand":                         struct{}{},
			"index":                        struct{}{},
			"point":                        struct{}{},
		},
	},
	":v:": Code{
		"\u270c",
		map[string]struct{}{
			"hand":         struct{}{},
			"v":            struct{}{},
			"victory":      struct{}{},
			"victory hand": struct{}{},
		},
	},
	":v_tone1:": Code{
		"\u270c\U0001f3fb",
		map[string]struct{}{
			"hand":            struct{}{},
			"light skin tone": struct{}{},
			"v":               struct{}{},
			"victory":         struct{}{},
			"victory hand":    struct{}{},
		},
	},
	":v_tone2:": Code{
		"\u270c\U0001f3fc",
		map[string]struct{}{
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"v":                      struct{}{},
			"victory":                struct{}{},
			"victory hand":           struct{}{},
		},
	},
	":v_tone3:": Code{
		"\u270c\U0001f3fd",
		map[string]struct{}{
			"hand":             struct{}{},
			"medium skin tone": struct{}{},
			"v":                struct{}{},
			"victory":          struct{}{},
			"victory hand":     struct{}{},
		},
	},
	":v_tone4:": Code{
		"\u270c\U0001f3fe",
		map[string]struct{}{
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"v":                     struct{}{},
			"victory":               struct{}{},
			"victory hand":          struct{}{},
		},
	},
	":v_tone5:": Code{
		"\u270c\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"hand":           struct{}{},
			"v":              struct{}{},
			"victory":        struct{}{},
			"victory hand":   struct{}{},
		},
	},
	":fingers_crossed:": Code{
		"\U0001f91e",
		map[string]struct{}{
			"cross":           struct{}{},
			"crossed fingers": struct{}{},
			"finger":          struct{}{},
			"hand":            struct{}{},
			"luck":            struct{}{},
		},
	},
	":fingers_crossed_tone1:": Code{
		"\U0001f91e\U0001f3fb",
		map[string]struct{}{
			"cross":           struct{}{},
			"crossed fingers": struct{}{},
			"finger":          struct{}{},
			"hand":            struct{}{},
			"light skin tone": struct{}{},
			"luck":            struct{}{},
		},
	},
	":fingers_crossed_tone2:": Code{
		"\U0001f91e\U0001f3fc",
		map[string]struct{}{
			"cross":                  struct{}{},
			"crossed fingers":        struct{}{},
			"finger":                 struct{}{},
			"hand":                   struct{}{},
			"luck":                   struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":fingers_crossed_tone3:": Code{
		"\U0001f91e\U0001f3fd",
		map[string]struct{}{
			"cross":            struct{}{},
			"crossed fingers":  struct{}{},
			"finger":           struct{}{},
			"hand":             struct{}{},
			"luck":             struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":fingers_crossed_tone4:": Code{
		"\U0001f91e\U0001f3fe",
		map[string]struct{}{
			"cross":                 struct{}{},
			"crossed fingers":       struct{}{},
			"finger":                struct{}{},
			"hand":                  struct{}{},
			"luck":                  struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":fingers_crossed_tone5:": Code{
		"\U0001f91e\U0001f3ff",
		map[string]struct{}{
			"cross":           struct{}{},
			"crossed fingers": struct{}{},
			"dark skin tone":  struct{}{},
			"finger":          struct{}{},
			"hand":            struct{}{},
			"luck":            struct{}{},
		},
	},
	":vulcan:": Code{
		"\U0001f596",
		map[string]struct{}{
			"finger":        struct{}{},
			"hand":          struct{}{},
			"spock":         struct{}{},
			"vulcan":        struct{}{},
			"vulcan salute": struct{}{},
		},
	},
	":vulcan_tone1:": Code{
		"\U0001f596\U0001f3fb",
		map[string]struct{}{
			"finger":          struct{}{},
			"hand":            struct{}{},
			"light skin tone": struct{}{},
			"spock":           struct{}{},
			"vulcan":          struct{}{},
			"vulcan salute":   struct{}{},
		},
	},
	":vulcan_tone2:": Code{
		"\U0001f596\U0001f3fc",
		map[string]struct{}{
			"finger":                 struct{}{},
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"spock":                  struct{}{},
			"vulcan":                 struct{}{},
			"vulcan salute":          struct{}{},
		},
	},
	":vulcan_tone3:": Code{
		"\U0001f596\U0001f3fd",
		map[string]struct{}{
			"finger":           struct{}{},
			"hand":             struct{}{},
			"medium skin tone": struct{}{},
			"spock":            struct{}{},
			"vulcan":           struct{}{},
			"vulcan salute":    struct{}{},
		},
	},
	":vulcan_tone4:": Code{
		"\U0001f596\U0001f3fe",
		map[string]struct{}{
			"finger":                struct{}{},
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"spock":                 struct{}{},
			"vulcan":                struct{}{},
			"vulcan salute":         struct{}{},
		},
	},
	":vulcan_tone5:": Code{
		"\U0001f596\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"finger":         struct{}{},
			"hand":           struct{}{},
			"spock":          struct{}{},
			"vulcan":         struct{}{},
			"vulcan salute":  struct{}{},
		},
	},
	":metal:": Code{
		"\U0001f918",
		map[string]struct{}{
			"finger":            struct{}{},
			"hand":              struct{}{},
			"horns":             struct{}{},
			"rock-on":           struct{}{},
			"sign of the horns": struct{}{},
		},
	},
	":metal_tone1:": Code{
		"\U0001f918\U0001f3fb",
		map[string]struct{}{
			"finger":            struct{}{},
			"hand":              struct{}{},
			"horns":             struct{}{},
			"light skin tone":   struct{}{},
			"rock-on":           struct{}{},
			"sign of the horns": struct{}{},
		},
	},
	":metal_tone2:": Code{
		"\U0001f918\U0001f3fc",
		map[string]struct{}{
			"finger":                 struct{}{},
			"hand":                   struct{}{},
			"horns":                  struct{}{},
			"medium-light skin tone": struct{}{},
			"rock-on":                struct{}{},
			"sign of the horns":      struct{}{},
		},
	},
	":metal_tone3:": Code{
		"\U0001f918\U0001f3fd",
		map[string]struct{}{
			"finger":            struct{}{},
			"hand":              struct{}{},
			"horns":             struct{}{},
			"medium skin tone":  struct{}{},
			"rock-on":           struct{}{},
			"sign of the horns": struct{}{},
		},
	},
	":metal_tone4:": Code{
		"\U0001f918\U0001f3fe",
		map[string]struct{}{
			"finger":                struct{}{},
			"hand":                  struct{}{},
			"horns":                 struct{}{},
			"medium-dark skin tone": struct{}{},
			"rock-on":               struct{}{},
			"sign of the horns":     struct{}{},
		},
	},
	":metal_tone5:": Code{
		"\U0001f918\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":    struct{}{},
			"finger":            struct{}{},
			"hand":              struct{}{},
			"horns":             struct{}{},
			"rock-on":           struct{}{},
			"sign of the horns": struct{}{},
		},
	},
	":call_me:": Code{
		"\U0001f919",
		map[string]struct{}{
			"call":         struct{}{},
			"call me hand": struct{}{},
			"hand":         struct{}{},
		},
	},
	":call_me_tone1:": Code{
		"\U0001f919\U0001f3fb",
		map[string]struct{}{
			"call":            struct{}{},
			"call me hand":    struct{}{},
			"hand":            struct{}{},
			"light skin tone": struct{}{},
		},
	},
	":call_me_tone2:": Code{
		"\U0001f919\U0001f3fc",
		map[string]struct{}{
			"call":                   struct{}{},
			"call me hand":           struct{}{},
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":call_me_tone3:": Code{
		"\U0001f919\U0001f3fd",
		map[string]struct{}{
			"call":             struct{}{},
			"call me hand":     struct{}{},
			"hand":             struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":call_me_tone4:": Code{
		"\U0001f919\U0001f3fe",
		map[string]struct{}{
			"call":                  struct{}{},
			"call me hand":          struct{}{},
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":call_me_tone5:": Code{
		"\U0001f919\U0001f3ff",
		map[string]struct{}{
			"call":           struct{}{},
			"call me hand":   struct{}{},
			"dark skin tone": struct{}{},
			"hand":           struct{}{},
		},
	},
	":hand_splayed_tone1:": Code{
		"\U0001f590\U0001f3fb",
		map[string]struct{}{
			"finger":                    struct{}{},
			"hand":                      struct{}{},
			"hand with fingers splayed": struct{}{},
			"light skin tone":           struct{}{},
			"splayed":                   struct{}{},
		},
	},
	":hand_splayed_tone2:": Code{
		"\U0001f590\U0001f3fc",
		map[string]struct{}{
			"finger":                    struct{}{},
			"hand":                      struct{}{},
			"hand with fingers splayed": struct{}{},
			"medium-light skin tone":    struct{}{},
			"splayed":                   struct{}{},
		},
	},
	":hand_splayed_tone3:": Code{
		"\U0001f590\U0001f3fd",
		map[string]struct{}{
			"finger":                    struct{}{},
			"hand":                      struct{}{},
			"hand with fingers splayed": struct{}{},
			"medium skin tone":          struct{}{},
			"splayed":                   struct{}{},
		},
	},
	":hand_splayed_tone4:": Code{
		"\U0001f590\U0001f3fe",
		map[string]struct{}{
			"finger":                    struct{}{},
			"hand":                      struct{}{},
			"hand with fingers splayed": struct{}{},
			"medium-dark skin tone":     struct{}{},
			"splayed":                   struct{}{},
		},
	},
	":hand_splayed_tone5:": Code{
		"\U0001f590\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":            struct{}{},
			"finger":                    struct{}{},
			"hand":                      struct{}{},
			"hand with fingers splayed": struct{}{},
			"splayed":                   struct{}{},
		},
	},
	":raised_hand:": Code{
		"\u270b",
		map[string]struct{}{
			"hand":        struct{}{},
			"raised hand": struct{}{},
		},
	},
	":raised_hand_tone1:": Code{
		"\u270b\U0001f3fb",
		map[string]struct{}{
			"hand":            struct{}{},
			"light skin tone": struct{}{},
			"raised hand":     struct{}{},
		},
	},
	":raised_hand_tone2:": Code{
		"\u270b\U0001f3fc",
		map[string]struct{}{
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"raised hand":            struct{}{},
		},
	},
	":raised_hand_tone3:": Code{
		"\u270b\U0001f3fd",
		map[string]struct{}{
			"hand":             struct{}{},
			"medium skin tone": struct{}{},
			"raised hand":      struct{}{},
		},
	},
	":raised_hand_tone4:": Code{
		"\u270b\U0001f3fe",
		map[string]struct{}{
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"raised hand":           struct{}{},
		},
	},
	":raised_hand_tone5:": Code{
		"\u270b\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"hand":           struct{}{},
			"raised hand":    struct{}{},
		},
	},
	":ok_hand:": Code{
		"\U0001f44c",
		map[string]struct{}{
			"hand":    struct{}{},
			"OK":      struct{}{},
			"OK hand": struct{}{},
		},
	},
	":ok_hand_tone1:": Code{
		"\U0001f44c\U0001f3fb",
		map[string]struct{}{
			"hand":            struct{}{},
			"light skin tone": struct{}{},
			"OK":              struct{}{},
			"OK hand":         struct{}{},
		},
	},
	":ok_hand_tone2:": Code{
		"\U0001f44c\U0001f3fc",
		map[string]struct{}{
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"OK":                     struct{}{},
			"OK hand":                struct{}{},
		},
	},
	":ok_hand_tone3:": Code{
		"\U0001f44c\U0001f3fd",
		map[string]struct{}{
			"hand":             struct{}{},
			"medium skin tone": struct{}{},
			"OK":               struct{}{},
			"OK hand":          struct{}{},
		},
	},
	":ok_hand_tone4:": Code{
		"\U0001f44c\U0001f3fe",
		map[string]struct{}{
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"OK":                    struct{}{},
			"OK hand":               struct{}{},
		},
	},
	":ok_hand_tone5:": Code{
		"\U0001f44c\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"hand":           struct{}{},
			"OK":             struct{}{},
			"OK hand":        struct{}{},
		},
	},
	":thumbsup:": Code{
		"\U0001f44d",
		map[string]struct{}{
			"#ERROR!": struct{}{},
		},
	},
	":thumbsup_tone1:": Code{
		"\U0001f44d\U0001f3fb",
		map[string]struct{}{
			"#ERROR!": struct{}{},
		},
	},
	":thumbsup_tone2:": Code{
		"\U0001f44d\U0001f3fc",
		map[string]struct{}{
			"#ERROR!": struct{}{},
		},
	},
	":thumbsup_tone3:": Code{
		"\U0001f44d\U0001f3fd",
		map[string]struct{}{
			"#ERROR!": struct{}{},
		},
	},
	":thumbsup_tone4:": Code{
		"\U0001f44d\U0001f3fe",
		map[string]struct{}{
			"#ERROR!": struct{}{},
		},
	},
	":thumbsup_tone5:": Code{
		"\U0001f44d\U0001f3ff",
		map[string]struct{}{
			"#ERROR!": struct{}{},
		},
	},
	":thumbsdown:": Code{
		"\U0001f44e",
		map[string]struct{}{
			"-1":          struct{}{},
			"down":        struct{}{},
			"hand":        struct{}{},
			"thumb":       struct{}{},
			"thumbs down": struct{}{},
		},
	},
	":thumbsdown_tone1:": Code{
		"\U0001f44e\U0001f3fb",
		map[string]struct{}{
			"-1":              struct{}{},
			"down":            struct{}{},
			"hand":            struct{}{},
			"light skin tone": struct{}{},
			"thumb":           struct{}{},
			"thumbs down":     struct{}{},
		},
	},
	":thumbsdown_tone2:": Code{
		"\U0001f44e\U0001f3fc",
		map[string]struct{}{
			"-1":                     struct{}{},
			"down":                   struct{}{},
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"thumb":                  struct{}{},
			"thumbs down":            struct{}{},
		},
	},
	":thumbsdown_tone3:": Code{
		"\U0001f44e\U0001f3fd",
		map[string]struct{}{
			"-1":               struct{}{},
			"down":             struct{}{},
			"hand":             struct{}{},
			"medium skin tone": struct{}{},
			"thumb":            struct{}{},
			"thumbs down":      struct{}{},
		},
	},
	":thumbsdown_tone4:": Code{
		"\U0001f44e\U0001f3fe",
		map[string]struct{}{
			"-1":                    struct{}{},
			"down":                  struct{}{},
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"thumb":                 struct{}{},
			"thumbs down":           struct{}{},
		},
	},
	":thumbsdown_tone5:": Code{
		"\U0001f44e\U0001f3ff",
		map[string]struct{}{
			"-1":             struct{}{},
			"dark skin tone": struct{}{},
			"down":           struct{}{},
			"hand":           struct{}{},
			"thumb":          struct{}{},
			"thumbs down":    struct{}{},
		},
	},
	":fist:": Code{
		"\u270a",
		map[string]struct{}{
			"clenched":    struct{}{},
			"fist":        struct{}{},
			"hand":        struct{}{},
			"punch":       struct{}{},
			"raised fist": struct{}{},
		},
	},
	":fist_tone1:": Code{
		"\u270a\U0001f3fb",
		map[string]struct{}{
			"clenched":        struct{}{},
			"fist":            struct{}{},
			"hand":            struct{}{},
			"light skin tone": struct{}{},
			"punch":           struct{}{},
			"raised fist":     struct{}{},
		},
	},
	":fist_tone2:": Code{
		"\u270a\U0001f3fc",
		map[string]struct{}{
			"clenched":               struct{}{},
			"fist":                   struct{}{},
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"punch":                  struct{}{},
			"raised fist":            struct{}{},
		},
	},
	":fist_tone3:": Code{
		"\u270a\U0001f3fd",
		map[string]struct{}{
			"clenched":         struct{}{},
			"fist":             struct{}{},
			"hand":             struct{}{},
			"medium skin tone": struct{}{},
			"punch":            struct{}{},
			"raised fist":      struct{}{},
		},
	},
	":fist_tone4:": Code{
		"\u270a\U0001f3fe",
		map[string]struct{}{
			"clenched":              struct{}{},
			"fist":                  struct{}{},
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"punch":                 struct{}{},
			"raised fist":           struct{}{},
		},
	},
	":fist_tone5:": Code{
		"\u270a\U0001f3ff",
		map[string]struct{}{
			"clenched":       struct{}{},
			"dark skin tone": struct{}{},
			"fist":           struct{}{},
			"hand":           struct{}{},
			"punch":          struct{}{},
			"raised fist":    struct{}{},
		},
	},
	":punch:": Code{
		"\U0001f44a",
		map[string]struct{}{
			"clenched":      struct{}{},
			"fist":          struct{}{},
			"hand":          struct{}{},
			"oncoming fist": struct{}{},
			"punch":         struct{}{},
		},
	},
	":punch_tone1:": Code{
		"\U0001f44a\U0001f3fb",
		map[string]struct{}{
			"clenched":        struct{}{},
			"fist":            struct{}{},
			"hand":            struct{}{},
			"light skin tone": struct{}{},
			"oncoming fist":   struct{}{},
			"punch":           struct{}{},
		},
	},
	":punch_tone2:": Code{
		"\U0001f44a\U0001f3fc",
		map[string]struct{}{
			"clenched":               struct{}{},
			"fist":                   struct{}{},
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"oncoming fist":          struct{}{},
			"punch":                  struct{}{},
		},
	},
	":punch_tone3:": Code{
		"\U0001f44a\U0001f3fd",
		map[string]struct{}{
			"clenched":         struct{}{},
			"fist":             struct{}{},
			"hand":             struct{}{},
			"medium skin tone": struct{}{},
			"oncoming fist":    struct{}{},
			"punch":            struct{}{},
		},
	},
	":punch_tone4:": Code{
		"\U0001f44a\U0001f3fe",
		map[string]struct{}{
			"clenched":              struct{}{},
			"fist":                  struct{}{},
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"oncoming fist":         struct{}{},
			"punch":                 struct{}{},
		},
	},
	":punch_tone5:": Code{
		"\U0001f44a\U0001f3ff",
		map[string]struct{}{
			"clenched":       struct{}{},
			"dark skin tone": struct{}{},
			"fist":           struct{}{},
			"hand":           struct{}{},
			"oncoming fist":  struct{}{},
			"punch":          struct{}{},
		},
	},
	":left_facing_fist:": Code{
		"\U0001f91b",
		map[string]struct{}{
			"fist":             struct{}{},
			"left-facing fist": struct{}{},
			"leftwards":        struct{}{},
		},
	},
	":left_facing_fist_tone1:": Code{
		"\U0001f91b\U0001f3fb",
		map[string]struct{}{
			"fist":             struct{}{},
			"left-facing fist": struct{}{},
			"leftwards":        struct{}{},
			"light skin tone":  struct{}{},
		},
	},
	":left_facing_fist_tone2:": Code{
		"\U0001f91b\U0001f3fc",
		map[string]struct{}{
			"fist":                   struct{}{},
			"left-facing fist":       struct{}{},
			"leftwards":              struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":left_facing_fist_tone3:": Code{
		"\U0001f91b\U0001f3fd",
		map[string]struct{}{
			"fist":             struct{}{},
			"left-facing fist": struct{}{},
			"leftwards":        struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":left_facing_fist_tone4:": Code{
		"\U0001f91b\U0001f3fe",
		map[string]struct{}{
			"fist":                  struct{}{},
			"left-facing fist":      struct{}{},
			"leftwards":             struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":left_facing_fist_tone5:": Code{
		"\U0001f91b\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":   struct{}{},
			"fist":             struct{}{},
			"left-facing fist": struct{}{},
			"leftwards":        struct{}{},
		},
	},
	":right_facing_fist:": Code{
		"\U0001f91c",
		map[string]struct{}{
			"fist":              struct{}{},
			"right-facing fist": struct{}{},
			"rightwards":        struct{}{},
		},
	},
	":right_facing_fist_tone1:": Code{
		"\U0001f91c\U0001f3fb",
		map[string]struct{}{
			"fist":              struct{}{},
			"light skin tone":   struct{}{},
			"right-facing fist": struct{}{},
			"rightwards":        struct{}{},
		},
	},
	":right_facing_fist_tone2:": Code{
		"\U0001f91c\U0001f3fc",
		map[string]struct{}{
			"fist":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"right-facing fist":      struct{}{},
			"rightwards":             struct{}{},
		},
	},
	":right_facing_fist_tone3:": Code{
		"\U0001f91c\U0001f3fd",
		map[string]struct{}{
			"fist":              struct{}{},
			"medium skin tone":  struct{}{},
			"right-facing fist": struct{}{},
			"rightwards":        struct{}{},
		},
	},
	":right_facing_fist_tone4:": Code{
		"\U0001f91c\U0001f3fe",
		map[string]struct{}{
			"fist":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"right-facing fist":     struct{}{},
			"rightwards":            struct{}{},
		},
	},
	":right_facing_fist_tone5:": Code{
		"\U0001f91c\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":    struct{}{},
			"fist":              struct{}{},
			"right-facing fist": struct{}{},
			"rightwards":        struct{}{},
		},
	},
	":raised_back_of_hand:": Code{
		"\U0001f91a",
		map[string]struct{}{
			"backhand":            struct{}{},
			"raised":              struct{}{},
			"raised back of hand": struct{}{},
		},
	},
	":raised_back_of_hand_tone1:": Code{
		"\U0001f91a\U0001f3fb",
		map[string]struct{}{
			"backhand":            struct{}{},
			"light skin tone":     struct{}{},
			"raised":              struct{}{},
			"raised back of hand": struct{}{},
		},
	},
	":raised_back_of_hand_tone2:": Code{
		"\U0001f91a\U0001f3fc",
		map[string]struct{}{
			"backhand":               struct{}{},
			"medium-light skin tone": struct{}{},
			"raised":                 struct{}{},
			"raised back of hand":    struct{}{},
		},
	},
	":raised_back_of_hand_tone3:": Code{
		"\U0001f91a\U0001f3fd",
		map[string]struct{}{
			"backhand":            struct{}{},
			"medium skin tone":    struct{}{},
			"raised":              struct{}{},
			"raised back of hand": struct{}{},
		},
	},
	":raised_back_of_hand_tone4:": Code{
		"\U0001f91a\U0001f3fe",
		map[string]struct{}{
			"backhand":              struct{}{},
			"medium-dark skin tone": struct{}{},
			"raised":                struct{}{},
			"raised back of hand":   struct{}{},
		},
	},
	":raised_back_of_hand_tone5:": Code{
		"\U0001f91a\U0001f3ff",
		map[string]struct{}{
			"backhand":            struct{}{},
			"dark skin tone":      struct{}{},
			"raised":              struct{}{},
			"raised back of hand": struct{}{},
		},
	},
	":wave:": Code{
		"\U0001f44b",
		map[string]struct{}{
			"hand":        struct{}{},
			"wave":        struct{}{},
			"waving":      struct{}{},
			"waving hand": struct{}{},
		},
	},
	":wave_tone1:": Code{
		"\U0001f44b\U0001f3fb",
		map[string]struct{}{
			"hand":            struct{}{},
			"light skin tone": struct{}{},
			"wave":            struct{}{},
			"waving":          struct{}{},
			"waving hand":     struct{}{},
		},
	},
	":wave_tone2:": Code{
		"\U0001f44b\U0001f3fc",
		map[string]struct{}{
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"wave":                   struct{}{},
			"waving":                 struct{}{},
			"waving hand":            struct{}{},
		},
	},
	":wave_tone3:": Code{
		"\U0001f44b\U0001f3fd",
		map[string]struct{}{
			"hand":             struct{}{},
			"medium skin tone": struct{}{},
			"wave":             struct{}{},
			"waving":           struct{}{},
			"waving hand":      struct{}{},
		},
	},
	":wave_tone4:": Code{
		"\U0001f44b\U0001f3fe",
		map[string]struct{}{
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"wave":                  struct{}{},
			"waving":                struct{}{},
			"waving hand":           struct{}{},
		},
	},
	":wave_tone5:": Code{
		"\U0001f44b\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"hand":           struct{}{},
			"wave":           struct{}{},
			"waving":         struct{}{},
			"waving hand":    struct{}{},
		},
	},
	":love_you_gesture:": Code{
		"\U0001f91f",
		map[string]struct{}{
			"hand":             struct{}{},
			"ILY":              struct{}{},
			"love-you gesture": struct{}{},
		},
	},
	":love_you_gesture_tone1:": Code{
		"\U0001f91f\U0001f3fb",
		map[string]struct{}{
			"hand":             struct{}{},
			"ILY":              struct{}{},
			"light skin tone":  struct{}{},
			"love-you gesture": struct{}{},
		},
	},
	":love_you_gesture_tone2:": Code{
		"\U0001f91f\U0001f3fc",
		map[string]struct{}{
			"hand":                   struct{}{},
			"ILY":                    struct{}{},
			"love-you gesture":       struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":love_you_gesture_tone3:": Code{
		"\U0001f91f\U0001f3fd",
		map[string]struct{}{
			"hand":             struct{}{},
			"ILY":              struct{}{},
			"love-you gesture": struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":love_you_gesture_tone4:": Code{
		"\U0001f91f\U0001f3fe",
		map[string]struct{}{
			"hand":                  struct{}{},
			"ILY":                   struct{}{},
			"love-you gesture":      struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":love_you_gesture_tone5:": Code{
		"\U0001f91f\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":   struct{}{},
			"hand":             struct{}{},
			"ILY":              struct{}{},
			"love-you gesture": struct{}{},
		},
	},
	":writing_hand:": Code{
		"\u270d",
		map[string]struct{}{
			"hand":         struct{}{},
			"write":        struct{}{},
			"writing hand": struct{}{},
		},
	},
	":writing_hand_tone1:": Code{
		"\u270d\U0001f3fb",
		map[string]struct{}{
			"hand":            struct{}{},
			"light skin tone": struct{}{},
			"write":           struct{}{},
			"writing hand":    struct{}{},
		},
	},
	":writing_hand_tone2:": Code{
		"\u270d\U0001f3fc",
		map[string]struct{}{
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"write":                  struct{}{},
			"writing hand":           struct{}{},
		},
	},
	":writing_hand_tone3:": Code{
		"\u270d\U0001f3fd",
		map[string]struct{}{
			"hand":             struct{}{},
			"medium skin tone": struct{}{},
			"write":            struct{}{},
			"writing hand":     struct{}{},
		},
	},
	":writing_hand_tone4:": Code{
		"\u270d\U0001f3fe",
		map[string]struct{}{
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"write":                 struct{}{},
			"writing hand":          struct{}{},
		},
	},
	":writing_hand_tone5:": Code{
		"\u270d\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"hand":           struct{}{},
			"write":          struct{}{},
			"writing hand":   struct{}{},
		},
	},
	":clap:": Code{
		"\U0001f44f",
		map[string]struct{}{
			"clap":           struct{}{},
			"clapping hands": struct{}{},
			"hand":           struct{}{},
		},
	},
	":clap_tone1:": Code{
		"\U0001f44f\U0001f3fb",
		map[string]struct{}{
			"clap":            struct{}{},
			"clapping hands":  struct{}{},
			"hand":            struct{}{},
			"light skin tone": struct{}{},
		},
	},
	":clap_tone2:": Code{
		"\U0001f44f\U0001f3fc",
		map[string]struct{}{
			"clap":                   struct{}{},
			"clapping hands":         struct{}{},
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":clap_tone3:": Code{
		"\U0001f44f\U0001f3fd",
		map[string]struct{}{
			"clap":             struct{}{},
			"clapping hands":   struct{}{},
			"hand":             struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":clap_tone4:": Code{
		"\U0001f44f\U0001f3fe",
		map[string]struct{}{
			"clap":                  struct{}{},
			"clapping hands":        struct{}{},
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":clap_tone5:": Code{
		"\U0001f44f\U0001f3ff",
		map[string]struct{}{
			"clap":           struct{}{},
			"clapping hands": struct{}{},
			"dark skin tone": struct{}{},
			"hand":           struct{}{},
		},
	},
	":open_hands:": Code{
		"\U0001f450",
		map[string]struct{}{
			"hand":       struct{}{},
			"open":       struct{}{},
			"open hands": struct{}{},
		},
	},
	":open_hands_tone1:": Code{
		"\U0001f450\U0001f3fb",
		map[string]struct{}{
			"hand":            struct{}{},
			"light skin tone": struct{}{},
			"open":            struct{}{},
			"open hands":      struct{}{},
		},
	},
	":open_hands_tone2:": Code{
		"\U0001f450\U0001f3fc",
		map[string]struct{}{
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"open":                   struct{}{},
			"open hands":             struct{}{},
		},
	},
	":open_hands_tone3:": Code{
		"\U0001f450\U0001f3fd",
		map[string]struct{}{
			"hand":             struct{}{},
			"medium skin tone": struct{}{},
			"open":             struct{}{},
			"open hands":       struct{}{},
		},
	},
	":open_hands_tone4:": Code{
		"\U0001f450\U0001f3fe",
		map[string]struct{}{
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"open":                  struct{}{},
			"open hands":            struct{}{},
		},
	},
	":open_hands_tone5:": Code{
		"\U0001f450\U0001f3ff",
		map[string]struct{}{
			"dark skin tone": struct{}{},
			"hand":           struct{}{},
			"open":           struct{}{},
			"open hands":     struct{}{},
		},
	},
	":raised_hands:": Code{
		"\U0001f64c",
		map[string]struct{}{
			"celebration":   struct{}{},
			"gesture":       struct{}{},
			"hand":          struct{}{},
			"hooray":        struct{}{},
			"raised":        struct{}{},
			"raising hands": struct{}{},
		},
	},
	":raised_hands_tone1:": Code{
		"\U0001f64c\U0001f3fb",
		map[string]struct{}{
			"celebration":     struct{}{},
			"gesture":         struct{}{},
			"hand":            struct{}{},
			"hooray":          struct{}{},
			"light skin tone": struct{}{},
			"raised":          struct{}{},
			"raising hands":   struct{}{},
		},
	},
	":raised_hands_tone2:": Code{
		"\U0001f64c\U0001f3fc",
		map[string]struct{}{
			"celebration":            struct{}{},
			"gesture":                struct{}{},
			"hand":                   struct{}{},
			"hooray":                 struct{}{},
			"medium-light skin tone": struct{}{},
			"raised":                 struct{}{},
			"raising hands":          struct{}{},
		},
	},
	":raised_hands_tone3:": Code{
		"\U0001f64c\U0001f3fd",
		map[string]struct{}{
			"celebration":      struct{}{},
			"gesture":          struct{}{},
			"hand":             struct{}{},
			"hooray":           struct{}{},
			"medium skin tone": struct{}{},
			"raised":           struct{}{},
			"raising hands":    struct{}{},
		},
	},
	":raised_hands_tone4:": Code{
		"\U0001f64c\U0001f3fe",
		map[string]struct{}{
			"celebration":           struct{}{},
			"gesture":               struct{}{},
			"hand":                  struct{}{},
			"hooray":                struct{}{},
			"medium-dark skin tone": struct{}{},
			"raised":                struct{}{},
			"raising hands":         struct{}{},
		},
	},
	":raised_hands_tone5:": Code{
		"\U0001f64c\U0001f3ff",
		map[string]struct{}{
			"celebration":    struct{}{},
			"dark skin tone": struct{}{},
			"gesture":        struct{}{},
			"hand":           struct{}{},
			"hooray":         struct{}{},
			"raised":         struct{}{},
			"raising hands":  struct{}{},
		},
	},
	":palms_up_together:": Code{
		"\U0001f932",
		map[string]struct{}{
			"palms up together": struct{}{},
			"prayer":            struct{}{},
			"cupped hands":      struct{}{},
		},
	},
	":palms_up_together_tone1:": Code{
		"\U0001f932\U0001f3fb",
		map[string]struct{}{
			"light skin tone":   struct{}{},
			"palms up together": struct{}{},
			"prayer":            struct{}{},
		},
	},
	":palms_up_together_tone2:": Code{
		"\U0001f932\U0001f3fc",
		map[string]struct{}{
			"medium-light skin tone": struct{}{},
			"palms up together":      struct{}{},
			"prayer":                 struct{}{},
		},
	},
	":palms_up_together_tone3:": Code{
		"\U0001f932\U0001f3fd",
		map[string]struct{}{
			"medium skin tone":  struct{}{},
			"palms up together": struct{}{},
			"prayer":            struct{}{},
		},
	},
	":palms_up_together_tone4:": Code{
		"\U0001f932\U0001f3fe",
		map[string]struct{}{
			"medium-dark skin tone": struct{}{},
			"palms up together":     struct{}{},
			"prayer":                struct{}{},
		},
	},
	":palms_up_together_tone5:": Code{
		"\U0001f932\U0001f3ff",
		map[string]struct{}{
			"dark skin tone":    struct{}{},
			"palms up together": struct{}{},
			"prayer":            struct{}{},
		},
	},
	":pray:": Code{
		"\U0001f64f",
		map[string]struct{}{
			"ask":          struct{}{},
			"bow":          struct{}{},
			"folded":       struct{}{},
			"folded hands": struct{}{},
			"gesture":      struct{}{},
			"hand":         struct{}{},
			"please":       struct{}{},
			"pray":         struct{}{},
			"thanks":       struct{}{},
		},
	},
	":pray_tone1:": Code{
		"\U0001f64f\U0001f3fb",
		map[string]struct{}{
			"ask":             struct{}{},
			"bow":             struct{}{},
			"folded":          struct{}{},
			"folded hands":    struct{}{},
			"gesture":         struct{}{},
			"hand":            struct{}{},
			"light skin tone": struct{}{},
			"please":          struct{}{},
			"pray":            struct{}{},
			"thanks":          struct{}{},
		},
	},
	":pray_tone2:": Code{
		"\U0001f64f\U0001f3fc",
		map[string]struct{}{
			"ask":                    struct{}{},
			"bow":                    struct{}{},
			"folded":                 struct{}{},
			"folded hands":           struct{}{},
			"gesture":                struct{}{},
			"hand":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"please":                 struct{}{},
			"pray":                   struct{}{},
			"thanks":                 struct{}{},
		},
	},
	":pray_tone3:": Code{
		"\U0001f64f\U0001f3fd",
		map[string]struct{}{
			"ask":              struct{}{},
			"bow":              struct{}{},
			"folded":           struct{}{},
			"folded hands":     struct{}{},
			"gesture":          struct{}{},
			"hand":             struct{}{},
			"medium skin tone": struct{}{},
			"please":           struct{}{},
			"pray":             struct{}{},
			"thanks":           struct{}{},
		},
	},
	":pray_tone4:": Code{
		"\U0001f64f\U0001f3fe",
		map[string]struct{}{
			"ask":                   struct{}{},
			"bow":                   struct{}{},
			"folded":                struct{}{},
			"folded hands":          struct{}{},
			"gesture":               struct{}{},
			"hand":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"please":                struct{}{},
			"pray":                  struct{}{},
			"thanks":                struct{}{},
		},
	},
	":pray_tone5:": Code{
		"\U0001f64f\U0001f3ff",
		map[string]struct{}{
			"ask":            struct{}{},
			"bow":            struct{}{},
			"dark skin tone": struct{}{},
			"folded":         struct{}{},
			"folded hands":   struct{}{},
			"gesture":        struct{}{},
			"hand":           struct{}{},
			"please":         struct{}{},
			"pray":           struct{}{},
			"thanks":         struct{}{},
		},
	},
	":handshake:": Code{
		"\U0001f91d",
		map[string]struct{}{
			"agreement": struct{}{},
			"hand":      struct{}{},
			"handshake": struct{}{},
			"meeting":   struct{}{},
			"shake":     struct{}{},
		},
	},
	":nail_care:": Code{
		"\U0001f485",
		map[string]struct{}{
			"care":        struct{}{},
			"cosmetics":   struct{}{},
			"manicure":    struct{}{},
			"nail":        struct{}{},
			"nail polish": struct{}{},
			"polish":      struct{}{},
		},
	},
	":nail_care_tone1:": Code{
		"\U0001f485\U0001f3fb",
		map[string]struct{}{
			"care":            struct{}{},
			"cosmetics":       struct{}{},
			"light skin tone": struct{}{},
			"manicure":        struct{}{},
			"nail":            struct{}{},
			"nail polish":     struct{}{},
			"polish":          struct{}{},
		},
	},
	":nail_care_tone2:": Code{
		"\U0001f485\U0001f3fc",
		map[string]struct{}{
			"care":                   struct{}{},
			"cosmetics":              struct{}{},
			"manicure":               struct{}{},
			"medium-light skin tone": struct{}{},
			"nail":                   struct{}{},
			"nail polish":            struct{}{},
			"polish":                 struct{}{},
		},
	},
	":nail_care_tone3:": Code{
		"\U0001f485\U0001f3fd",
		map[string]struct{}{
			"care":             struct{}{},
			"cosmetics":        struct{}{},
			"manicure":         struct{}{},
			"medium skin tone": struct{}{},
			"nail":             struct{}{},
			"nail polish":      struct{}{},
			"polish":           struct{}{},
		},
	},
	":nail_care_tone4:": Code{
		"\U0001f485\U0001f3fe",
		map[string]struct{}{
			"care":                  struct{}{},
			"cosmetics":             struct{}{},
			"manicure":              struct{}{},
			"medium-dark skin tone": struct{}{},
			"nail":                  struct{}{},
			"nail polish":           struct{}{},
			"polish":                struct{}{},
		},
	},
	":nail_care_tone5:": Code{
		"\U0001f485\U0001f3ff",
		map[string]struct{}{
			"care":           struct{}{},
			"cosmetics":      struct{}{},
			"dark skin tone": struct{}{},
			"manicure":       struct{}{},
			"nail":           struct{}{},
			"nail polish":    struct{}{},
			"polish":         struct{}{},
		},
	},
	":ear:": Code{
		"\U0001f442",
		map[string]struct{}{
			"body": struct{}{},
			"ear":  struct{}{},
		},
	},
	":ear_tone1:": Code{
		"\U0001f442\U0001f3fb",
		map[string]struct{}{
			"body":            struct{}{},
			"ear":             struct{}{},
			"light skin tone": struct{}{},
		},
	},
	":ear_tone2:": Code{
		"\U0001f442\U0001f3fc",
		map[string]struct{}{
			"body":                   struct{}{},
			"ear":                    struct{}{},
			"medium-light skin tone": struct{}{},
		},
	},
	":ear_tone3:": Code{
		"\U0001f442\U0001f3fd",
		map[string]struct{}{
			"body":             struct{}{},
			"ear":              struct{}{},
			"medium skin tone": struct{}{},
		},
	},
	":ear_tone4:": Code{
		"\U0001f442\U0001f3fe",
		map[string]struct{}{
			"body":                  struct{}{},
			"ear":                   struct{}{},
			"medium-dark skin tone": struct{}{},
		},
	},
	":ear_tone5:": Code{
		"\U0001f442\U0001f3ff",
		map[string]struct{}{
			"body":           struct{}{},
			"dark skin tone": struct{}{},
			"ear":            struct{}{},
		},
	},
	":nose:": Code{
		"\U0001f443",
		map[string]struct{}{
			"body": struct{}{},
			"nose": struct{}{},
		},
	},
	":nose_tone1:": Code{
		"\U0001f443\U0001f3fb",
		map[string]struct{}{
			"body":            struct{}{},
			"light skin tone": struct{}{},
			"nose":            struct{}{},
		},
	},
	":nose_tone2:": Code{
		"\U0001f443\U0001f3fc",
		map[string]struct{}{
			"body":                   struct{}{},
			"medium-light skin tone": struct{}{},
			"nose":                   struct{}{},
		},
	},
	":nose_tone3:": Code{
		"\U0001f443\U0001f3fd",
		map[string]struct{}{
			"body":             struct{}{},
			"medium skin tone": struct{}{},
			"nose":             struct{}{},
		},
	},
	":nose_tone4:": Code{
		"\U0001f443\U0001f3fe",
		map[string]struct{}{
			"body":                  struct{}{},
			"medium-dark skin tone": struct{}{},
			"nose":                  struct{}{},
		},
	},
	":nose_tone5:": Code{
		"\U0001f443\U0001f3ff",
		map[string]struct{}{
			"body":           struct{}{},
			"dark skin tone": struct{}{},
			"nose":           struct{}{},
		},
	},
	":footprints:": Code{
		"\U0001f463",
		map[string]struct{}{
			"clothing":   struct{}{},
			"footprint":  struct{}{},
			"footprints": struct{}{},
			"print":      struct{}{},
		},
	},
	":eyes:": Code{
		"\U0001f440",
		map[string]struct{}{
			"eye":  struct{}{},
			"eyes": struct{}{},
			"face": struct{}{},
		},
	},
	":eye:": Code{
		"\U0001f441",
		map[string]struct{}{
			"body": struct{}{},
			"eye":  struct{}{},
		},
	},
	":eye_in_speech_bubble:": Code{
		"\U0001f441\ufe0f\u200d\U0001f5e8\ufe0f",
		map[string]struct{}{
			"eye":                  struct{}{},
			"eye in speech bubble": struct{}{},
			"speech bubble":        struct{}{},
			"witness":              struct{}{},
		},
	},
	":brain:": Code{
		"\U0001f9e0",
		map[string]struct{}{
			"brain":       struct{}{},
			"intelligent": struct{}{},
		},
	},
	":tongue:": Code{
		"\U0001f445",
		map[string]struct{}{
			"body":   struct{}{},
			"tongue": struct{}{},
		},
	},
	":lips:": Code{
		"\U0001f444",
		map[string]struct{}{
			"lips":  struct{}{},
			"mouth": struct{}{},
		},
	},
	":kiss:": Code{
		"\U0001f48b",
		map[string]struct{}{
			"kiss":      struct{}{},
			"kiss mark": struct{}{},
			"lips":      struct{}{},
		},
	},
	":cupid:": Code{
		"\U0001f498",
		map[string]struct{}{
			"arrow":            struct{}{},
			"cupid":            struct{}{},
			"heart with arrow": struct{}{},
		},
	},
	":heart:": Code{
		"\u2764",
		map[string]struct{}{
			"heart":     struct{}{},
			"red heart": struct{}{},
		},
	},
	":heartbeat:": Code{
		"\U0001f493",
		map[string]struct{}{
			"beating":       struct{}{},
			"beating heart": struct{}{},
			"heartbeat":     struct{}{},
			"pulsating":     struct{}{},
		},
	},
	":broken_heart:": Code{
		"\U0001f494",
		map[string]struct{}{
			"break":        struct{}{},
			"broken":       struct{}{},
			"broken heart": struct{}{},
		},
	},
	":two_hearts:": Code{
		"\U0001f495",
		map[string]struct{}{
			"love":       struct{}{},
			"two hearts": struct{}{},
		},
	},
	":sparkling_heart:": Code{
		"\U0001f496",
		map[string]struct{}{
			"excited":         struct{}{},
			"sparkle":         struct{}{},
			"sparkling heart": struct{}{},
		},
	},
	":heartpulse:": Code{
		"\U0001f497",
		map[string]struct{}{
			"excited":       struct{}{},
			"growing":       struct{}{},
			"growing heart": struct{}{},
			"nervous":       struct{}{},
			"pulse":         struct{}{},
		},
	},
	":blue_heart:": Code{
		"\U0001f499",
		map[string]struct{}{
			"blue":       struct{}{},
			"blue heart": struct{}{},
		},
	},
	":green_heart:": Code{
		"\U0001f49a",
		map[string]struct{}{
			"green":       struct{}{},
			"green heart": struct{}{},
		},
	},
	":yellow_heart:": Code{
		"\U0001f49b",
		map[string]struct{}{
			"yellow":       struct{}{},
			"yellow heart": struct{}{},
		},
	},
	":orange_heart:": Code{
		"\U0001f9e1",
		map[string]struct{}{
			"orange":       struct{}{},
			"orange heart": struct{}{},
		},
	},
	":purple_heart:": Code{
		"\U0001f49c",
		map[string]struct{}{
			"purple":       struct{}{},
			"purple heart": struct{}{},
		},
	},
	":black_heart:": Code{
		"\U0001f5a4",
		map[string]struct{}{
			"black":       struct{}{},
			"black heart": struct{}{},
			"evil":        struct{}{},
			"wicked":      struct{}{},
		},
	},
	":gift_heart:": Code{
		"\U0001f49d",
		map[string]struct{}{
			"heart with ribbon": struct{}{},
			"ribbon":            struct{}{},
			"valentine":         struct{}{},
		},
	},
	":revolving_hearts:": Code{
		"\U0001f49e",
		map[string]struct{}{
			"revolving":        struct{}{},
			"revolving hearts": struct{}{},
		},
	},
	":heart_decoration:": Code{
		"\U0001f49f",
		map[string]struct{}{
			"heart":            struct{}{},
			"heart decoration": struct{}{},
		},
	},
	":heart_exclamation:": Code{
		"\u2763",
		map[string]struct{}{
			"exclamation":             struct{}{},
			"heavy heart exclamation": struct{}{},
			"mark":                    struct{}{},
			"punctuation":             struct{}{},
		},
	},
	":love_letter:": Code{
		"\U0001f48c",
		map[string]struct{}{
			"heart":       struct{}{},
			"letter":      struct{}{},
			"love":        struct{}{},
			"love letter": struct{}{},
			"mail":        struct{}{},
		},
	},
	":zzz:": Code{
		"\U0001f4a4",
		map[string]struct{}{
			"comic": struct{}{},
			"sleep": struct{}{},
			"zzz":   struct{}{},
		},
	},
	":anger:": Code{
		"\U0001f4a2",
		map[string]struct{}{
			"anger symbol": struct{}{},
			"angry":        struct{}{},
			"comic":        struct{}{},
			"mad":          struct{}{},
		},
	},
	":bomb:": Code{
		"\U0001f4a3",
		map[string]struct{}{
			"bomb":  struct{}{},
			"comic": struct{}{},
		},
	},
	":boom:": Code{
		"\U0001f4a5",
		map[string]struct{}{
			"boom":      struct{}{},
			"collision": struct{}{},
			"comic":     struct{}{},
		},
	},
	":sweat_drops:": Code{
		"\U0001f4a6",
		map[string]struct{}{
			"comic":          struct{}{},
			"splashing":      struct{}{},
			"sweat":          struct{}{},
			"sweat droplets": struct{}{},
		},
	},
	":dash:": Code{
		"\U0001f4a8",
		map[string]struct{}{
			"comic":        struct{}{},
			"dash":         struct{}{},
			"dashing away": struct{}{},
			"running":      struct{}{},
		},
	},
	":dizzy:": Code{
		"\U0001f4ab",
		map[string]struct{}{
			"comic": struct{}{},
			"dizzy": struct{}{},
			"star":  struct{}{},
		},
	},
	":speech_balloon:": Code{
		"\U0001f4ac",
		map[string]struct{}{
			"balloon":        struct{}{},
			"bubble":         struct{}{},
			"comic":          struct{}{},
			"dialog":         struct{}{},
			"speech":         struct{}{},
			"speech balloon": struct{}{},
		},
	},
	":speech_left:": Code{
		"\U0001f5e8",
		map[string]struct{}{
			"dialog":             struct{}{},
			"left speech bubble": struct{}{},
			"speech":             struct{}{},
		},
	},
	":anger_right:": Code{
		"\U0001f5ef",
		map[string]struct{}{
			"angry":              struct{}{},
			"balloon":            struct{}{},
			"bubble":             struct{}{},
			"mad":                struct{}{},
			"right anger bubble": struct{}{},
		},
	},
	":thought_balloon:": Code{
		"\U0001f4ad",
		map[string]struct{}{
			"balloon":         struct{}{},
			"bubble":          struct{}{},
			"comic":           struct{}{},
			"thought":         struct{}{},
			"thought balloon": struct{}{},
		},
	},
	":hole:": Code{
		"\U0001f573",
		map[string]struct{}{
			"hole": struct{}{},
		},
	},
	":eyeglasses:": Code{
		"\U0001f453",
		map[string]struct{}{
			"clothing":   struct{}{},
			"eye":        struct{}{},
			"eyeglasses": struct{}{},
			"eyewear":    struct{}{},
			"glasses":    struct{}{},
		},
	},
	":dark_sunglasses:": Code{
		"\U0001f576",
		map[string]struct{}{
			"dark":       struct{}{},
			"eye":        struct{}{},
			"eyewear":    struct{}{},
			"glasses":    struct{}{},
			"sunglasses": struct{}{},
		},
	},
	":necktie:": Code{
		"\U0001f454",
		map[string]struct{}{
			"clothing": struct{}{},
			"necktie":  struct{}{},
		},
	},
	":shirt:": Code{
		"\U0001f455",
		map[string]struct{}{
			"clothing": struct{}{},
			"shirt":    struct{}{},
			"t-shirt":  struct{}{},
			"tshirt":   struct{}{},
		},
	},
	":jeans:": Code{
		"\U0001f456",
		map[string]struct{}{
			"clothing": struct{}{},
			"jeans":    struct{}{},
			"pants":    struct{}{},
			"trousers": struct{}{},
		},
	},
	":scarf:": Code{
		"\U0001f9e3",
		map[string]struct{}{
			"neck":  struct{}{},
			"scarf": struct{}{},
		},
	},
	":gloves:": Code{
		"\U0001f9e4",
		map[string]struct{}{
			"gloves": struct{}{},
			"hand":   struct{}{},
		},
	},
	":coat:": Code{
		"\U0001f9e5",
		map[string]struct{}{
			"coat":   struct{}{},
			"jacket": struct{}{},
		},
	},
	":socks:": Code{
		"\U0001f9e6",
		map[string]struct{}{
			"socks":    struct{}{},
			"stocking": struct{}{},
		},
	},
	":dress:": Code{
		"\U0001f457",
		map[string]struct{}{
			"clothing": struct{}{},
			"dress":    struct{}{},
		},
	},
	":kimono:": Code{
		"\U0001f458",
		map[string]struct{}{
			"clothing": struct{}{},
			"kimono":   struct{}{},
		},
	},
	":bikini:": Code{
		"\U0001f459",
		map[string]struct{}{
			"bikini":   struct{}{},
			"clothing": struct{}{},
			"swim":     struct{}{},
		},
	},
	":womans_clothes:": Code{
		"\U0001f45a",
		map[string]struct{}{
			"clothing":        struct{}{},
			"woman":           struct{}{},
			"woman’s clothes": struct{}{},
		},
	},
	":purse:": Code{
		"\U0001f45b",
		map[string]struct{}{
			"clothing": struct{}{},
			"coin":     struct{}{},
			"purse":    struct{}{},
		},
	},
	":handbag:": Code{
		"\U0001f45c",
		map[string]struct{}{
			"bag":      struct{}{},
			"clothing": struct{}{},
			"handbag":  struct{}{},
			"purse":    struct{}{},
		},
	},
	":pouch:": Code{
		"\U0001f45d",
		map[string]struct{}{
			"bag":        struct{}{},
			"clothing":   struct{}{},
			"clutch bag": struct{}{},
			"pouch":      struct{}{},
		},
	},
	":shopping_bags:": Code{
		"\U0001f6cd",
		map[string]struct{}{
			"bag":           struct{}{},
			"hotel":         struct{}{},
			"shopping":      struct{}{},
			"shopping bags": struct{}{},
		},
	},
	":school_satchel:": Code{
		"\U0001f392",
		map[string]struct{}{
			"bag":             struct{}{},
			"satchel":         struct{}{},
			"school":          struct{}{},
			"school backpack": struct{}{},
		},
	},
	":mans_shoe:": Code{
		"\U0001f45e",
		map[string]struct{}{
			"clothing":   struct{}{},
			"man":        struct{}{},
			"man’s shoe": struct{}{},
			"shoe":       struct{}{},
		},
	},
	":athletic_shoe:": Code{
		"\U0001f45f",
		map[string]struct{}{
			"athletic":     struct{}{},
			"clothing":     struct{}{},
			"running shoe": struct{}{},
			"shoe":         struct{}{},
			"sneaker":      struct{}{},
		},
	},
	":high_heel:": Code{
		"\U0001f460",
		map[string]struct{}{
			"clothing":         struct{}{},
			"heel":             struct{}{},
			"high-heeled shoe": struct{}{},
			"shoe":             struct{}{},
			"woman":            struct{}{},
		},
	},
	":sandal:": Code{
		"\U0001f461",
		map[string]struct{}{
			"clothing":       struct{}{},
			"sandal":         struct{}{},
			"shoe":           struct{}{},
			"woman":          struct{}{},
			"woman’s sandal": struct{}{},
		},
	},
	":boot:": Code{
		"\U0001f462",
		map[string]struct{}{
			"boot":         struct{}{},
			"clothing":     struct{}{},
			"shoe":         struct{}{},
			"woman":        struct{}{},
			"woman’s boot": struct{}{},
		},
	},
	":crown:": Code{
		"\U0001f451",
		map[string]struct{}{
			"clothing": struct{}{},
			"crown":    struct{}{},
			"king":     struct{}{},
			"queen":    struct{}{},
		},
	},
	":womans_hat:": Code{
		"\U0001f452",
		map[string]struct{}{
			"clothing":    struct{}{},
			"hat":         struct{}{},
			"woman":       struct{}{},
			"woman’s hat": struct{}{},
		},
	},
	":tophat:": Code{
		"\U0001f3a9",
		map[string]struct{}{
			"clothing": struct{}{},
			"hat":      struct{}{},
			"top":      struct{}{},
			"top hat":  struct{}{},
			"tophat":   struct{}{},
		},
	},
	":mortar_board:": Code{
		"\U0001f393",
		map[string]struct{}{
			"cap":            struct{}{},
			"celebration":    struct{}{},
			"clothing":       struct{}{},
			"graduation":     struct{}{},
			"graduation cap": struct{}{},
			"hat":            struct{}{},
		},
	},
	":billed_cap:": Code{
		"\U0001f9e2",
		map[string]struct{}{
			"baseball cap": struct{}{},
			"billed cap":   struct{}{},
		},
	},
	":helmet_with_cross:": Code{
		"\u26d1",
		map[string]struct{}{
			"aid":                    struct{}{},
			"cross":                  struct{}{},
			"face":                   struct{}{},
			"hat":                    struct{}{},
			"helmet":                 struct{}{},
			"rescue worker’s helmet": struct{}{},
		},
	},
	":prayer_beads:": Code{
		"\U0001f4ff",
		map[string]struct{}{
			"beads":        struct{}{},
			"clothing":     struct{}{},
			"necklace":     struct{}{},
			"prayer":       struct{}{},
			"prayer beads": struct{}{},
			"religion":     struct{}{},
		},
	},
	":lipstick:": Code{
		"\U0001f484",
		map[string]struct{}{
			"cosmetics": struct{}{},
			"lipstick":  struct{}{},
			"makeup":    struct{}{},
		},
	},
	":ring:": Code{
		"\U0001f48d",
		map[string]struct{}{
			"diamond": struct{}{},
			"ring":    struct{}{},
		},
	},
	":gem:": Code{
		"\U0001f48e",
		map[string]struct{}{
			"diamond":   struct{}{},
			"gem":       struct{}{},
			"gem stone": struct{}{},
			"jewel":     struct{}{},
		},
	},
	":monkey_face:": Code{
		"\U0001f435",
		map[string]struct{}{
			"face":        struct{}{},
			"monkey":      struct{}{},
			"monkey face": struct{}{},
		},
	},
	":monkey:": Code{
		"\U0001f412",
		map[string]struct{}{
			"monkey": struct{}{},
		},
	},
	":gorilla:": Code{
		"\U0001f98d",
		map[string]struct{}{
			"gorilla": struct{}{},
		},
	},
	":dog:": Code{
		"\U0001f436",
		map[string]struct{}{
			"dog":      struct{}{},
			"dog face": struct{}{},
			"face":     struct{}{},
			"pet":      struct{}{},
		},
	},
	":dog2:": Code{
		"\U0001f415",
		map[string]struct{}{
			"dog": struct{}{},
			"pet": struct{}{},
		},
	},
	":poodle:": Code{
		"\U0001f429",
		map[string]struct{}{
			"dog":    struct{}{},
			"poodle": struct{}{},
		},
	},
	":wolf:": Code{
		"\U0001f43a",
		map[string]struct{}{
			"face":      struct{}{},
			"wolf":      struct{}{},
			"wolf face": struct{}{},
		},
	},
	":fox:": Code{
		"\U0001f98a",
		map[string]struct{}{
			"face":     struct{}{},
			"fox":      struct{}{},
			"fox face": struct{}{},
		},
	},
	":cat:": Code{
		"\U0001f431",
		map[string]struct{}{
			"cat":      struct{}{},
			"cat face": struct{}{},
			"face":     struct{}{},
			"pet":      struct{}{},
		},
	},
	":cat2:": Code{
		"\U0001f408",
		map[string]struct{}{
			"cat": struct{}{},
			"pet": struct{}{},
		},
	},
	":lion_face:": Code{
		"\U0001f981",
		map[string]struct{}{
			"face":      struct{}{},
			"Leo":       struct{}{},
			"lion":      struct{}{},
			"lion face": struct{}{},
			"zodiac":    struct{}{},
		},
	},
	":tiger:": Code{
		"\U0001f42f",
		map[string]struct{}{
			"face":       struct{}{},
			"tiger":      struct{}{},
			"tiger face": struct{}{},
		},
	},
	":tiger2:": Code{
		"\U0001f405",
		map[string]struct{}{
			"tiger": struct{}{},
		},
	},
	":leopard:": Code{
		"\U0001f406",
		map[string]struct{}{
			"leopard": struct{}{},
		},
	},
	":horse:": Code{
		"\U0001f434",
		map[string]struct{}{
			"face":       struct{}{},
			"horse":      struct{}{},
			"horse face": struct{}{},
		},
	},
	":racehorse:": Code{
		"\U0001f40e",
		map[string]struct{}{
			"equestrian": struct{}{},
			"horse":      struct{}{},
			"racehorse":  struct{}{},
			"racing":     struct{}{},
		},
	},
	":unicorn:": Code{
		"\U0001f984",
		map[string]struct{}{
			"face":         struct{}{},
			"unicorn":      struct{}{},
			"unicorn face": struct{}{},
		},
	},
	":zebra:": Code{
		"\U0001f993",
		map[string]struct{}{
			"stripe": struct{}{},
			"zebra":  struct{}{},
		},
	},
	":deer:": Code{
		"\U0001f98c",
		map[string]struct{}{
			"deer": struct{}{},
		},
	},
	":cow:": Code{
		"\U0001f42e",
		map[string]struct{}{
			"cow":      struct{}{},
			"cow face": struct{}{},
			"face":     struct{}{},
		},
	},
	":ox:": Code{
		"\U0001f402",
		map[string]struct{}{
			"bull":   struct{}{},
			"ox":     struct{}{},
			"Taurus": struct{}{},
			"zodiac": struct{}{},
		},
	},
	":water_buffalo:": Code{
		"\U0001f403",
		map[string]struct{}{
			"buffalo":       struct{}{},
			"water":         struct{}{},
			"water buffalo": struct{}{},
		},
	},
	":cow2:": Code{
		"\U0001f404",
		map[string]struct{}{
			"cow": struct{}{},
		},
	},
	":pig:": Code{
		"\U0001f437",
		map[string]struct{}{
			"face":     struct{}{},
			"pig":      struct{}{},
			"pig face": struct{}{},
		},
	},
	":pig2:": Code{
		"\U0001f416",
		map[string]struct{}{
			"pig": struct{}{},
			"sow": struct{}{},
		},
	},
	":boar:": Code{
		"\U0001f417",
		map[string]struct{}{
			"boar": struct{}{},
			"pig":  struct{}{},
		},
	},
	":pig_nose:": Code{
		"\U0001f43d",
		map[string]struct{}{
			"face":     struct{}{},
			"nose":     struct{}{},
			"pig":      struct{}{},
			"pig nose": struct{}{},
		},
	},
	":ram:": Code{
		"\U0001f40f",
		map[string]struct{}{
			"Aries":  struct{}{},
			"male":   struct{}{},
			"ram":    struct{}{},
			"sheep":  struct{}{},
			"zodiac": struct{}{},
		},
	},
	":sheep:": Code{
		"\U0001f411",
		map[string]struct{}{
			"ewe":    struct{}{},
			"female": struct{}{},
			"sheep":  struct{}{},
		},
	},
	":goat:": Code{
		"\U0001f410",
		map[string]struct{}{
			"Capricorn": struct{}{},
			"goat":      struct{}{},
			"zodiac":    struct{}{},
		},
	},
	":dromedary_camel:": Code{
		"\U0001f42a",
		map[string]struct{}{
			"camel":     struct{}{},
			"dromedary": struct{}{},
			"hump":      struct{}{},
		},
	},
	":camel:": Code{
		"\U0001f42b",
		map[string]struct{}{
			"bactrian":       struct{}{},
			"camel":          struct{}{},
			"hump":           struct{}{},
			"two-hump camel": struct{}{},
		},
	},
	":giraffe:": Code{
		"\U0001f992",
		map[string]struct{}{
			"giraffe": struct{}{},
			"spots":   struct{}{},
		},
	},
	":elephant:": Code{
		"\U0001f418",
		map[string]struct{}{
			"elephant": struct{}{},
		},
	},
	":rhino:": Code{
		"\U0001f98f",
		map[string]struct{}{
			"rhinoceros": struct{}{},
		},
	},
	":mouse:": Code{
		"\U0001f42d",
		map[string]struct{}{
			"face":       struct{}{},
			"mouse":      struct{}{},
			"mouse face": struct{}{},
		},
	},
	":mouse2:": Code{
		"\U0001f401",
		map[string]struct{}{
			"mouse": struct{}{},
		},
	},
	":rat:": Code{
		"\U0001f400",
		map[string]struct{}{
			"rat": struct{}{},
		},
	},
	":hamster:": Code{
		"\U0001f439",
		map[string]struct{}{
			"face":         struct{}{},
			"hamster":      struct{}{},
			"hamster face": struct{}{},
			"pet":          struct{}{},
		},
	},
	":rabbit:": Code{
		"\U0001f430",
		map[string]struct{}{
			"bunny":       struct{}{},
			"face":        struct{}{},
			"pet":         struct{}{},
			"rabbit":      struct{}{},
			"rabbit face": struct{}{},
		},
	},
	":rabbit2:": Code{
		"\U0001f407",
		map[string]struct{}{
			"bunny":  struct{}{},
			"pet":    struct{}{},
			"rabbit": struct{}{},
		},
	},
	":chipmunk:": Code{
		"\U0001f43f",
		map[string]struct{}{
			"chipmunk": struct{}{},
		},
	},
	":hedgehog:": Code{
		"\U0001f994",
		map[string]struct{}{
			"hedgehog": struct{}{},
			"spiny":    struct{}{},
		},
	},
	":bat:": Code{
		"\U0001f987",
		map[string]struct{}{
			"bat":     struct{}{},
			"vampire": struct{}{},
		},
	},
	":bear:": Code{
		"\U0001f43b",
		map[string]struct{}{
			"bear":      struct{}{},
			"bear face": struct{}{},
			"face":      struct{}{},
		},
	},
	":koala:": Code{
		"\U0001f428",
		map[string]struct{}{
			"bear":  struct{}{},
			"koala": struct{}{},
		},
	},
	":panda_face:": Code{
		"\U0001f43c",
		map[string]struct{}{
			"face":       struct{}{},
			"panda":      struct{}{},
			"panda face": struct{}{},
		},
	},
	":feet:": Code{
		"\U0001f43e",
		map[string]struct{}{
			"feet":       struct{}{},
			"paw":        struct{}{},
			"paw prints": struct{}{},
			"print":      struct{}{},
		},
	},
	":turkey:": Code{
		"\U0001f983",
		map[string]struct{}{
			"bird":   struct{}{},
			"turkey": struct{}{},
		},
	},
	":chicken:": Code{
		"\U0001f414",
		map[string]struct{}{
			"bird":    struct{}{},
			"chicken": struct{}{},
		},
	},
	":rooster:": Code{
		"\U0001f413",
		map[string]struct{}{
			"bird":    struct{}{},
			"rooster": struct{}{},
		},
	},
	":hatching_chick:": Code{
		"\U0001f423",
		map[string]struct{}{
			"baby":           struct{}{},
			"bird":           struct{}{},
			"chick":          struct{}{},
			"hatching":       struct{}{},
			"hatching chick": struct{}{},
		},
	},
	":baby_chick:": Code{
		"\U0001f424",
		map[string]struct{}{
			"baby":       struct{}{},
			"baby chick": struct{}{},
			"bird":       struct{}{},
			"chick":      struct{}{},
		},
	},
	":hatched_chick:": Code{
		"\U0001f425",
		map[string]struct{}{
			"baby":                    struct{}{},
			"bird":                    struct{}{},
			"chick":                   struct{}{},
			"front-facing baby chick": struct{}{},
		},
	},
	":bird:": Code{
		"\U0001f426",
		map[string]struct{}{
			"bird": struct{}{},
		},
	},
	":penguin:": Code{
		"\U0001f427",
		map[string]struct{}{
			"bird":    struct{}{},
			"penguin": struct{}{},
		},
	},
	":dove:": Code{
		"\U0001f54a",
		map[string]struct{}{
			"bird":  struct{}{},
			"dove":  struct{}{},
			"fly":   struct{}{},
			"peace": struct{}{},
		},
	},
	":eagle:": Code{
		"\U0001f985",
		map[string]struct{}{
			"bird":  struct{}{},
			"eagle": struct{}{},
		},
	},
	":duck:": Code{
		"\U0001f986",
		map[string]struct{}{
			"bird": struct{}{},
			"duck": struct{}{},
		},
	},
	":owl:": Code{
		"\U0001f989",
		map[string]struct{}{
			"bird": struct{}{},
			"owl":  struct{}{},
			"wise": struct{}{},
		},
	},
	":frog:": Code{
		"\U0001f438",
		map[string]struct{}{
			"face":      struct{}{},
			"frog":      struct{}{},
			"frog face": struct{}{},
		},
	},
	":crocodile:": Code{
		"\U0001f40a",
		map[string]struct{}{
			"crocodile": struct{}{},
		},
	},
	":turtle:": Code{
		"\U0001f422",
		map[string]struct{}{
			"terrapin": struct{}{},
			"tortoise": struct{}{},
			"turtle":   struct{}{},
		},
	},
	":lizard:": Code{
		"\U0001f98e",
		map[string]struct{}{
			"lizard":  struct{}{},
			"reptile": struct{}{},
		},
	},
	":snake:": Code{
		"\U0001f40d",
		map[string]struct{}{
			"bearer":    struct{}{},
			"Ophiuchus": struct{}{},
			"serpent":   struct{}{},
			"snake":     struct{}{},
			"zodiac":    struct{}{},
		},
	},
	":dragon_face:": Code{
		"\U0001f432",
		map[string]struct{}{
			"dragon":      struct{}{},
			"dragon face": struct{}{},
			"face":        struct{}{},
			"fairy tale":  struct{}{},
		},
	},
	":dragon:": Code{
		"\U0001f409",
		map[string]struct{}{
			"dragon":     struct{}{},
			"fairy tale": struct{}{},
		},
	},
	":sauropod:": Code{
		"\U0001f995",
		map[string]struct{}{
			"brachiosaurus": struct{}{},
			"brontosaurus":  struct{}{},
			"diplodocus":    struct{}{},
			"sauropod":      struct{}{},
		},
	},
	":t_rex:": Code{
		"\U0001f996",
		map[string]struct{}{
			"T-Rex":             struct{}{},
			"Tyrannosaurus Rex": struct{}{},
		},
	},
	":whale:": Code{
		"\U0001f433",
		map[string]struct{}{
			"face":           struct{}{},
			"spouting":       struct{}{},
			"spouting whale": struct{}{},
			"whale":          struct{}{},
		},
	},
	":whale2:": Code{
		"\U0001f40b",
		map[string]struct{}{
			"whale": struct{}{},
		},
	},
	":dolphin:": Code{
		"\U0001f42c",
		map[string]struct{}{
			"dolphin": struct{}{},
			"flipper": struct{}{},
		},
	},
	":fish:": Code{
		"\U0001f41f",
		map[string]struct{}{
			"fish":   struct{}{},
			"Pisces": struct{}{},
			"zodiac": struct{}{},
		},
	},
	":tropical_fish:": Code{
		"\U0001f420",
		map[string]struct{}{
			"fish":          struct{}{},
			"tropical":      struct{}{},
			"tropical fish": struct{}{},
		},
	},
	":blowfish:": Code{
		"\U0001f421",
		map[string]struct{}{
			"blowfish": struct{}{},
			"fish":     struct{}{},
		},
	},
	":shark:": Code{
		"\U0001f988",
		map[string]struct{}{
			"fish":  struct{}{},
			"shark": struct{}{},
		},
	},
	":octopus:": Code{
		"\U0001f419",
		map[string]struct{}{
			"octopus": struct{}{},
		},
	},
	":shell:": Code{
		"\U0001f41a",
		map[string]struct{}{
			"shell":        struct{}{},
			"spiral":       struct{}{},
			"spiral shell": struct{}{},
		},
	},
	":crab:": Code{
		"\U0001f980",
		map[string]struct{}{
			"Cancer": struct{}{},
			"crab":   struct{}{},
			"zodiac": struct{}{},
		},
	},
	":shrimp:": Code{
		"\U0001f990",
		map[string]struct{}{
			"food":      struct{}{},
			"shellfish": struct{}{},
			"shrimp":    struct{}{},
			"small":     struct{}{},
		},
	},
	":squid:": Code{
		"\U0001f991",
		map[string]struct{}{
			"food":   struct{}{},
			"molusc": struct{}{},
			"squid":  struct{}{},
		},
	},
	":snail:": Code{
		"\U0001f40c",
		map[string]struct{}{
			"snail": struct{}{},
		},
	},
	":butterfly:": Code{
		"\U0001f98b",
		map[string]struct{}{
			"butterfly": struct{}{},
			"insect":    struct{}{},
			"pretty":    struct{}{},
		},
	},
	":bug:": Code{
		"\U0001f41b",
		map[string]struct{}{
			"bug":    struct{}{},
			"insect": struct{}{},
		},
	},
	":ant:": Code{
		"\U0001f41c",
		map[string]struct{}{
			"ant":    struct{}{},
			"insect": struct{}{},
		},
	},
	":bee:": Code{
		"\U0001f41d",
		map[string]struct{}{
			"bee":      struct{}{},
			"honeybee": struct{}{},
			"insect":   struct{}{},
		},
	},
	":beetle:": Code{
		"\U0001f41e",
		map[string]struct{}{
			"beetle":      struct{}{},
			"insect":      struct{}{},
			"lady beetle": struct{}{},
			"ladybird":    struct{}{},
			"ladybug":     struct{}{},
		},
	},
	":cricket:": Code{
		"\U0001f997",
		map[string]struct{}{
			"cricket":     struct{}{},
			"grasshopper": struct{}{},
			"Orthoptera":  struct{}{},
		},
	},
	":spider:": Code{
		"\U0001f577",
		map[string]struct{}{
			"insect": struct{}{},
			"spider": struct{}{},
		},
	},
	":spider_web:": Code{
		"\U0001f578",
		map[string]struct{}{
			"spider":     struct{}{},
			"spider web": struct{}{},
			"web":        struct{}{},
		},
	},
	":scorpion:": Code{
		"\U0001f982",
		map[string]struct{}{
			"scorpio":  struct{}{},
			"Scorpio":  struct{}{},
			"scorpion": struct{}{},
			"zodiac":   struct{}{},
		},
	},
	":bouquet:": Code{
		"\U0001f490",
		map[string]struct{}{
			"bouquet": struct{}{},
			"flower":  struct{}{},
		},
	},
	":cherry_blossom:": Code{
		"\U0001f338",
		map[string]struct{}{
			"blossom":        struct{}{},
			"cherry":         struct{}{},
			"cherry blossom": struct{}{},
			"flower":         struct{}{},
		},
	},
	":white_flower:": Code{
		"\U0001f4ae",
		map[string]struct{}{
			"flower":       struct{}{},
			"white flower": struct{}{},
		},
	},
	":rosette:": Code{
		"\U0001f3f5",
		map[string]struct{}{
			"plant":   struct{}{},
			"rosette": struct{}{},
		},
	},
	":rose:": Code{
		"\U0001f339",
		map[string]struct{}{
			"flower": struct{}{},
			"rose":   struct{}{},
		},
	},
	":wilted_rose:": Code{
		"\U0001f940",
		map[string]struct{}{
			"flower":        struct{}{},
			"wilted":        struct{}{},
			"wilted flower": struct{}{},
		},
	},
	":hibiscus:": Code{
		"\U0001f33a",
		map[string]struct{}{
			"flower":   struct{}{},
			"hibiscus": struct{}{},
		},
	},
	":sunflower:": Code{
		"\U0001f33b",
		map[string]struct{}{
			"flower":    struct{}{},
			"sun":       struct{}{},
			"sunflower": struct{}{},
		},
	},
	":blossom:": Code{
		"\U0001f33c",
		map[string]struct{}{
			"blossom": struct{}{},
			"flower":  struct{}{},
		},
	},
	":tulip:": Code{
		"\U0001f337",
		map[string]struct{}{
			"flower": struct{}{},
			"tulip":  struct{}{},
		},
	},
	":seedling:": Code{
		"\U0001f331",
		map[string]struct{}{
			"seedling": struct{}{},
			"young":    struct{}{},
		},
	},
	":evergreen_tree:": Code{
		"\U0001f332",
		map[string]struct{}{
			"evergreen tree": struct{}{},
			"tree":           struct{}{},
		},
	},
	":deciduous_tree:": Code{
		"\U0001f333",
		map[string]struct{}{
			"deciduous":      struct{}{},
			"deciduous tree": struct{}{},
			"shedding":       struct{}{},
			"tree":           struct{}{},
		},
	},
	":palm_tree:": Code{
		"\U0001f334",
		map[string]struct{}{
			"palm":      struct{}{},
			"palm tree": struct{}{},
			"tree":      struct{}{},
		},
	},
	":cactus:": Code{
		"\U0001f335",
		map[string]struct{}{
			"cactus": struct{}{},
			"plant":  struct{}{},
		},
	},
	":ear_of_rice:": Code{
		"\U0001f33e",
		map[string]struct{}{
			"ear":           struct{}{},
			"grain":         struct{}{},
			"rice":          struct{}{},
			"sheaf of rice": struct{}{},
		},
	},
	":herb:": Code{
		"\U0001f33f",
		map[string]struct{}{
			"herb": struct{}{},
			"leaf": struct{}{},
		},
	},
	":shamrock:": Code{
		"\u2618",
		map[string]struct{}{
			"plant":    struct{}{},
			"shamrock": struct{}{},
		},
	},
	":four_leaf_clover:": Code{
		"\U0001f340",
		map[string]struct{}{
			"4":                struct{}{},
			"clover":           struct{}{},
			"four":             struct{}{},
			"four leaf clover": struct{}{},
			"leaf":             struct{}{},
		},
	},
	":maple_leaf:": Code{
		"\U0001f341",
		map[string]struct{}{
			"falling":    struct{}{},
			"leaf":       struct{}{},
			"maple":      struct{}{},
			"maple leaf": struct{}{},
		},
	},
	":fallen_leaf:": Code{
		"\U0001f342",
		map[string]struct{}{
			"fallen leaf": struct{}{},
			"falling":     struct{}{},
			"leaf":        struct{}{},
		},
	},
	":leaves:": Code{
		"\U0001f343",
		map[string]struct{}{
			"blow":                    struct{}{},
			"flutter":                 struct{}{},
			"leaf":                    struct{}{},
			"leaf fluttering in wind": struct{}{},
			"wind":                    struct{}{},
		},
	},
	":grapes:": Code{
		"\U0001f347",
		map[string]struct{}{
			"fruit":  struct{}{},
			"grape":  struct{}{},
			"grapes": struct{}{},
		},
	},
	":melon:": Code{
		"\U0001f348",
		map[string]struct{}{
			"fruit": struct{}{},
			"melon": struct{}{},
		},
	},
	":watermelon:": Code{
		"\U0001f349",
		map[string]struct{}{
			"fruit":      struct{}{},
			"watermelon": struct{}{},
		},
	},
	":tangerine:": Code{
		"\U0001f34a",
		map[string]struct{}{
			"fruit":     struct{}{},
			"orange":    struct{}{},
			"tangerine": struct{}{},
		},
	},
	":lemon:": Code{
		"\U0001f34b",
		map[string]struct{}{
			"citrus": struct{}{},
			"fruit":  struct{}{},
			"lemon":  struct{}{},
		},
	},
	":banana:": Code{
		"\U0001f34c",
		map[string]struct{}{
			"banana": struct{}{},
			"fruit":  struct{}{},
		},
	},
	":pineapple:": Code{
		"\U0001f34d",
		map[string]struct{}{
			"fruit":     struct{}{},
			"pineapple": struct{}{},
		},
	},
	":apple:": Code{
		"\U0001f34e",
		map[string]struct{}{
			"apple":     struct{}{},
			"fruit":     struct{}{},
			"red":       struct{}{},
			"red apple": struct{}{},
		},
	},
	":green_apple:": Code{
		"\U0001f34f",
		map[string]struct{}{
			"apple":       struct{}{},
			"fruit":       struct{}{},
			"green":       struct{}{},
			"green apple": struct{}{},
		},
	},
	":pear:": Code{
		"\U0001f350",
		map[string]struct{}{
			"fruit": struct{}{},
			"pear":  struct{}{},
		},
	},
	":peach:": Code{
		"\U0001f351",
		map[string]struct{}{
			"fruit": struct{}{},
			"peach": struct{}{},
		},
	},
	":cherries:": Code{
		"\U0001f352",
		map[string]struct{}{
			"berries":  struct{}{},
			"cherries": struct{}{},
			"cherry":   struct{}{},
			"fruit":    struct{}{},
			"red":      struct{}{},
		},
	},
	":strawberry:": Code{
		"\U0001f353",
		map[string]struct{}{
			"berry":      struct{}{},
			"fruit":      struct{}{},
			"strawberry": struct{}{},
		},
	},
	":kiwi:": Code{
		"\U0001f95d",
		map[string]struct{}{
			"food":       struct{}{},
			"fruit":      struct{}{},
			"kiwi":       struct{}{},
			"kiwi fruit": struct{}{},
		},
	},
	":tomato:": Code{
		"\U0001f345",
		map[string]struct{}{
			"fruit":     struct{}{},
			"tomato":    struct{}{},
			"vegetable": struct{}{},
		},
	},
	":coconut:": Code{
		"\U0001f965",
		map[string]struct{}{
			"coconut":     struct{}{},
			"palm":        struct{}{},
			"piña colada": struct{}{},
		},
	},
	":avocado:": Code{
		"\U0001f951",
		map[string]struct{}{
			"avocado": struct{}{},
			"food":    struct{}{},
			"fruit":   struct{}{},
		},
	},
	":eggplant:": Code{
		"\U0001f346",
		map[string]struct{}{
			"aubergine": struct{}{},
			"eggplant":  struct{}{},
			"vegetable": struct{}{},
		},
	},
	":potato:": Code{
		"\U0001f954",
		map[string]struct{}{
			"food":      struct{}{},
			"potato":    struct{}{},
			"vegetable": struct{}{},
		},
	},
	":carrot:": Code{
		"\U0001f955",
		map[string]struct{}{
			"carrot":    struct{}{},
			"food":      struct{}{},
			"vegetable": struct{}{},
		},
	},
	":corn:": Code{
		"\U0001f33d",
		map[string]struct{}{
			"corn":        struct{}{},
			"ear":         struct{}{},
			"ear of corn": struct{}{},
			"maize":       struct{}{},
			"maze":        struct{}{},
		},
	},
	":hot_pepper:": Code{
		"\U0001f336",
		map[string]struct{}{
			"hot":        struct{}{},
			"hot pepper": struct{}{},
			"pepper":     struct{}{},
		},
	},
	":cucumber:": Code{
		"\U0001f952",
		map[string]struct{}{
			"cucumber":  struct{}{},
			"food":      struct{}{},
			"pickle":    struct{}{},
			"vegetable": struct{}{},
		},
	},
	":broccoli:": Code{
		"\U0001f966",
		map[string]struct{}{
			"broccoli":     struct{}{},
			"wild cabbage": struct{}{},
		},
	},
	":mushroom:": Code{
		"\U0001f344",
		map[string]struct{}{
			"mushroom":  struct{}{},
			"toadstool": struct{}{},
		},
	},
	":peanuts:": Code{
		"\U0001f95c",
		map[string]struct{}{
			"food":      struct{}{},
			"nut":       struct{}{},
			"peanut":    struct{}{},
			"peanuts":   struct{}{},
			"vegetable": struct{}{},
		},
	},
	":chestnut:": Code{
		"\U0001f330",
		map[string]struct{}{
			"chestnut": struct{}{},
			"plant":    struct{}{},
		},
	},
	":bread:": Code{
		"\U0001f35e",
		map[string]struct{}{
			"bread": struct{}{},
			"loaf":  struct{}{},
		},
	},
	":croissant:": Code{
		"\U0001f950",
		map[string]struct{}{
			"bread":         struct{}{},
			"crescent roll": struct{}{},
			"croissant":     struct{}{},
			"food":          struct{}{},
			"french":        struct{}{},
		},
	},
	":french_bread:": Code{
		"\U0001f956",
		map[string]struct{}{
			"baguette":       struct{}{},
			"baguette bread": struct{}{},
			"bread":          struct{}{},
			"food":           struct{}{},
			"french":         struct{}{},
		},
	},
	":pretzel:": Code{
		"\U0001f968",
		map[string]struct{}{
			"pretzel":    struct{}{},
			"twisted":    struct{}{},
			"convoluted": struct{}{},
		},
	},
	":pancakes:": Code{
		"\U0001f95e",
		map[string]struct{}{
			"crêpe":    struct{}{},
			"food":     struct{}{},
			"hotcake":  struct{}{},
			"pancake":  struct{}{},
			"pancakes": struct{}{},
		},
	},
	":cheese:": Code{
		"\U0001f9c0",
		map[string]struct{}{
			"cheese":       struct{}{},
			"cheese wedge": struct{}{},
		},
	},
	":meat_on_bone:": Code{
		"\U0001f356",
		map[string]struct{}{
			"bone":         struct{}{},
			"meat":         struct{}{},
			"meat on bone": struct{}{},
		},
	},
	":poultry_leg:": Code{
		"\U0001f357",
		map[string]struct{}{
			"bone":        struct{}{},
			"chicken":     struct{}{},
			"leg":         struct{}{},
			"poultry":     struct{}{},
			"poultry leg": struct{}{},
		},
	},
	":cut_of_meat:": Code{
		"\U0001f969",
		map[string]struct{}{
			"chop":        struct{}{},
			"cut of meat": struct{}{},
			"lambchop":    struct{}{},
			"porkchop":    struct{}{},
			"steak":       struct{}{},
		},
	},
	":bacon:": Code{
		"\U0001f953",
		map[string]struct{}{
			"bacon": struct{}{},
			"food":  struct{}{},
			"meat":  struct{}{},
		},
	},
	":hamburger:": Code{
		"\U0001f354",
		map[string]struct{}{
			"burger":    struct{}{},
			"hamburger": struct{}{},
		},
	},
	":fries:": Code{
		"\U0001f35f",
		map[string]struct{}{
			"french":       struct{}{},
			"french fries": struct{}{},
			"fries":        struct{}{},
		},
	},
	":pizza:": Code{
		"\U0001f355",
		map[string]struct{}{
			"cheese": struct{}{},
			"pizza":  struct{}{},
			"slice":  struct{}{},
		},
	},
	":hotdog:": Code{
		"\U0001f32d",
		map[string]struct{}{
			"frankfurter": struct{}{},
			"hot dog":     struct{}{},
			"hotdog":      struct{}{},
			"sausage":     struct{}{},
		},
	},
	":sandwich:": Code{
		"\U0001f96a",
		map[string]struct{}{
			"bread":    struct{}{},
			"sandwich": struct{}{},
		},
	},
	":taco:": Code{
		"\U0001f32e",
		map[string]struct{}{
			"mexican": struct{}{},
			"taco":    struct{}{},
		},
	},
	":burrito:": Code{
		"\U0001f32f",
		map[string]struct{}{
			"burrito": struct{}{},
			"mexican": struct{}{},
			"wrap":    struct{}{},
		},
	},
	":stuffed_flatbread:": Code{
		"\U0001f959",
		map[string]struct{}{
			"falafel":           struct{}{},
			"flatbread":         struct{}{},
			"food":              struct{}{},
			"gyro":              struct{}{},
			"kebab":             struct{}{},
			"stuffed":           struct{}{},
			"stuffed flatbread": struct{}{},
		},
	},
	":egg:": Code{
		"\U0001f95a",
		map[string]struct{}{
			"egg":  struct{}{},
			"food": struct{}{},
		},
	},
	":cooking:": Code{
		"\U0001f373",
		map[string]struct{}{
			"cooking": struct{}{},
			"egg":     struct{}{},
			"frying":  struct{}{},
			"pan":     struct{}{},
		},
	},
	":shallow_pan_of_food:": Code{
		"\U0001f958",
		map[string]struct{}{
			"casserole":           struct{}{},
			"food":                struct{}{},
			"paella":              struct{}{},
			"pan":                 struct{}{},
			"shallow":             struct{}{},
			"shallow pan of food": struct{}{},
		},
	},
	":stew:": Code{
		"\U0001f372",
		map[string]struct{}{
			"pot":         struct{}{},
			"pot of food": struct{}{},
			"stew":        struct{}{},
		},
	},
	":bowl_with_spoon:": Code{
		"\U0001f963",
		map[string]struct{}{
			"bowl with spoon": struct{}{},
			"breakfast":       struct{}{},
			"cereal":          struct{}{},
			"congee":          struct{}{},
			"oatmeal":         struct{}{},
			"porridge":        struct{}{},
		},
	},
	":salad:": Code{
		"\U0001f957",
		map[string]struct{}{
			"food":        struct{}{},
			"green":       struct{}{},
			"green salad": struct{}{},
			"salad":       struct{}{},
		},
	},
	":popcorn:": Code{
		"\U0001f37f",
		map[string]struct{}{
			"popcorn": struct{}{},
		},
	},
	":canned_food:": Code{
		"\U0001f96b",
		map[string]struct{}{
			"can":         struct{}{},
			"canned food": struct{}{},
		},
	},
	":bento:": Code{
		"\U0001f371",
		map[string]struct{}{
			"bento":     struct{}{},
			"bento box": struct{}{},
			"box":       struct{}{},
		},
	},
	":rice_cracker:": Code{
		"\U0001f358",
		map[string]struct{}{
			"cracker":      struct{}{},
			"rice":         struct{}{},
			"rice cracker": struct{}{},
		},
	},
	":rice_ball:": Code{
		"\U0001f359",
		map[string]struct{}{
			"ball":      struct{}{},
			"Japanese":  struct{}{},
			"rice":      struct{}{},
			"rice ball": struct{}{},
		},
	},
	":rice:": Code{
		"\U0001f35a",
		map[string]struct{}{
			"cooked":      struct{}{},
			"cooked rice": struct{}{},
			"rice":        struct{}{},
		},
	},
	":curry:": Code{
		"\U0001f35b",
		map[string]struct{}{
			"curry":      struct{}{},
			"curry rice": struct{}{},
			"rice":       struct{}{},
		},
	},
	":ramen:": Code{
		"\U0001f35c",
		map[string]struct{}{
			"bowl":          struct{}{},
			"noodle":        struct{}{},
			"ramen":         struct{}{},
			"steaming":      struct{}{},
			"steaming bowl": struct{}{},
		},
	},
	":spaghetti:": Code{
		"\U0001f35d",
		map[string]struct{}{
			"pasta":     struct{}{},
			"spaghetti": struct{}{},
		},
	},
	":sweet_potato:": Code{
		"\U0001f360",
		map[string]struct{}{
			"potato":               struct{}{},
			"roasted":              struct{}{},
			"roasted sweet potato": struct{}{},
			"sweet":                struct{}{},
		},
	},
	":oden:": Code{
		"\U0001f362",
		map[string]struct{}{
			"kebab":   struct{}{},
			"oden":    struct{}{},
			"seafood": struct{}{},
			"skewer":  struct{}{},
			"stick":   struct{}{},
		},
	},
	":sushi:": Code{
		"\U0001f363",
		map[string]struct{}{
			"sushi": struct{}{},
		},
	},
	":fried_shrimp:": Code{
		"\U0001f364",
		map[string]struct{}{
			"fried":        struct{}{},
			"fried shrimp": struct{}{},
			"prawn":        struct{}{},
			"shrimp":       struct{}{},
			"tempura":      struct{}{},
		},
	},
	":fish_cake:": Code{
		"\U0001f365",
		map[string]struct{}{
			"cake":                 struct{}{},
			"fish":                 struct{}{},
			"fish cake with swirl": struct{}{},
			"pastry":               struct{}{},
			"swirl":                struct{}{},
		},
	},
	":dango:": Code{
		"\U0001f361",
		map[string]struct{}{
			"dango":    struct{}{},
			"dessert":  struct{}{},
			"Japanese": struct{}{},
			"skewer":   struct{}{},
			"stick":    struct{}{},
			"sweet":    struct{}{},
		},
	},
	":dumpling:": Code{
		"\U0001f95f",
		map[string]struct{}{
			"dumpling":   struct{}{},
			"empanada":   struct{}{},
			"gyōza":      struct{}{},
			"jiaozi":     struct{}{},
			"pierogi":    struct{}{},
			"potsticker": struct{}{},
		},
	},
	":fortune_cookie:": Code{
		"\U0001f960",
		map[string]struct{}{
			"fortune cookie": struct{}{},
			"prophecy":       struct{}{},
		},
	},
	":takeout_box:": Code{
		"\U0001f961",
		map[string]struct{}{
			"oyster pail": struct{}{},
			"takeout box": struct{}{},
		},
	},
	":icecream:": Code{
		"\U0001f366",
		map[string]struct{}{
			"cream":          struct{}{},
			"dessert":        struct{}{},
			"ice":            struct{}{},
			"icecream":       struct{}{},
			"soft":           struct{}{},
			"soft ice cream": struct{}{},
			"sweet":          struct{}{},
		},
	},
	":shaved_ice:": Code{
		"\U0001f367",
		map[string]struct{}{
			"dessert":    struct{}{},
			"ice":        struct{}{},
			"shaved":     struct{}{},
			"shaved ice": struct{}{},
			"sweet":      struct{}{},
		},
	},
	":ice_cream:": Code{
		"\U0001f368",
		map[string]struct{}{
			"cream":     struct{}{},
			"dessert":   struct{}{},
			"ice":       struct{}{},
			"ice cream": struct{}{},
			"sweet":     struct{}{},
		},
	},
	":doughnut:": Code{
		"\U0001f369",
		map[string]struct{}{
			"dessert":  struct{}{},
			"donut":    struct{}{},
			"doughnut": struct{}{},
			"sweet":    struct{}{},
		},
	},
	":cookie:": Code{
		"\U0001f36a",
		map[string]struct{}{
			"cookie":  struct{}{},
			"dessert": struct{}{},
			"sweet":   struct{}{},
		},
	},
	":birthday:": Code{
		"\U0001f382",
		map[string]struct{}{
			"birthday":      struct{}{},
			"birthday cake": struct{}{},
			"cake":          struct{}{},
			"celebration":   struct{}{},
			"dessert":       struct{}{},
			"pastry":        struct{}{},
			"sweet":         struct{}{},
		},
	},
	":cake:": Code{
		"\U0001f370",
		map[string]struct{}{
			"cake":      struct{}{},
			"dessert":   struct{}{},
			"pastry":    struct{}{},
			"shortcake": struct{}{},
			"slice":     struct{}{},
			"sweet":     struct{}{},
		},
	},
	":pie:": Code{
		"\U0001f967",
		map[string]struct{}{
			"filling": struct{}{},
			"pastry":  struct{}{},
			"pie":     struct{}{},
			"fruit":   struct{}{},
			"meat":    struct{}{},
		},
	},
	":chocolate_bar:": Code{
		"\U0001f36b",
		map[string]struct{}{
			"bar":           struct{}{},
			"chocolate":     struct{}{},
			"chocolate bar": struct{}{},
			"dessert":       struct{}{},
			"sweet":         struct{}{},
		},
	},
	":candy:": Code{
		"\U0001f36c",
		map[string]struct{}{
			"candy":   struct{}{},
			"dessert": struct{}{},
			"sweet":   struct{}{},
		},
	},
	":lollipop:": Code{
		"\U0001f36d",
		map[string]struct{}{
			"candy":    struct{}{},
			"dessert":  struct{}{},
			"lollipop": struct{}{},
			"sweet":    struct{}{},
		},
	},
	":custard:": Code{
		"\U0001f36e",
		map[string]struct{}{
			"custard": struct{}{},
			"dessert": struct{}{},
			"pudding": struct{}{},
			"sweet":   struct{}{},
		},
	},
	":honey_pot:": Code{
		"\U0001f36f",
		map[string]struct{}{
			"honey":     struct{}{},
			"honey pot": struct{}{},
			"honeypot":  struct{}{},
			"pot":       struct{}{},
			"sweet":     struct{}{},
		},
	},
	":baby_bottle:": Code{
		"\U0001f37c",
		map[string]struct{}{
			"baby":        struct{}{},
			"baby bottle": struct{}{},
			"bottle":      struct{}{},
			"drink":       struct{}{},
			"milk":        struct{}{},
		},
	},
	":milk:": Code{
		"\U0001f95b",
		map[string]struct{}{
			"drink":         struct{}{},
			"glass":         struct{}{},
			"glass of milk": struct{}{},
			"milk":          struct{}{},
		},
	},
	":coffee:": Code{
		"\u2615",
		map[string]struct{}{
			"beverage":     struct{}{},
			"coffee":       struct{}{},
			"drink":        struct{}{},
			"hot":          struct{}{},
			"hot beverage": struct{}{},
			"steaming":     struct{}{},
			"tea":          struct{}{},
		},
	},
	":tea:": Code{
		"\U0001f375",
		map[string]struct{}{
			"beverage":              struct{}{},
			"cup":                   struct{}{},
			"drink":                 struct{}{},
			"tea":                   struct{}{},
			"teacup":                struct{}{},
			"teacup without handle": struct{}{},
		},
	},
	":sake:": Code{
		"\U0001f376",
		map[string]struct{}{
			"bar":      struct{}{},
			"beverage": struct{}{},
			"bottle":   struct{}{},
			"cup":      struct{}{},
			"drink":    struct{}{},
			"sake":     struct{}{},
		},
	},
	":champagne:": Code{
		"\U0001f37e",
		map[string]struct{}{
			"bar":                      struct{}{},
			"bottle":                   struct{}{},
			"bottle with popping cork": struct{}{},
			"cork":                     struct{}{},
			"drink":                    struct{}{},
			"popping":                  struct{}{},
		},
	},
	":wine_glass:": Code{
		"\U0001f377",
		map[string]struct{}{
			"bar":        struct{}{},
			"beverage":   struct{}{},
			"drink":      struct{}{},
			"glass":      struct{}{},
			"wine":       struct{}{},
			"wine glass": struct{}{},
		},
	},
	":cocktail:": Code{
		"\U0001f378",
		map[string]struct{}{
			"bar":            struct{}{},
			"cocktail":       struct{}{},
			"cocktail glass": struct{}{},
			"drink":          struct{}{},
			"glass":          struct{}{},
		},
	},
	":tropical_drink:": Code{
		"\U0001f379",
		map[string]struct{}{
			"bar":            struct{}{},
			"drink":          struct{}{},
			"tropical":       struct{}{},
			"tropical drink": struct{}{},
		},
	},
	":beer:": Code{
		"\U0001f37a",
		map[string]struct{}{
			"bar":      struct{}{},
			"beer":     struct{}{},
			"beer mug": struct{}{},
			"drink":    struct{}{},
			"mug":      struct{}{},
		},
	},
	":beers:": Code{
		"\U0001f37b",
		map[string]struct{}{
			"bar":                struct{}{},
			"beer":               struct{}{},
			"clink":              struct{}{},
			"clinking beer mugs": struct{}{},
			"drink":              struct{}{},
			"mug":                struct{}{},
		},
	},
	":champagne_glass:": Code{
		"\U0001f942",
		map[string]struct{}{
			"celebrate":        struct{}{},
			"clink":            struct{}{},
			"clinking glasses": struct{}{},
			"drink":            struct{}{},
			"glass":            struct{}{},
		},
	},
	":tumbler_glass:": Code{
		"\U0001f943",
		map[string]struct{}{
			"glass":         struct{}{},
			"liquor":        struct{}{},
			"shot":          struct{}{},
			"tumbler":       struct{}{},
			"tumbler glass": struct{}{},
			"whisky":        struct{}{},
		},
	},
	":cup_with_straw:": Code{
		"\U0001f964",
		map[string]struct{}{
			"cup with straw": struct{}{},
			"juice":          struct{}{},
			"soda":           struct{}{},
			"malt":           struct{}{},
			"soft drink":     struct{}{},
			"water":          struct{}{},
		},
	},
	":chopsticks:": Code{
		"\U0001f962",
		map[string]struct{}{
			"chopsticks": struct{}{},
			"hashi":      struct{}{},
			"jeotgarak":  struct{}{},
			"kuaizi":     struct{}{},
		},
	},
	":fork_knife_plate:": Code{
		"\U0001f37d",
		map[string]struct{}{
			"cooking":                   struct{}{},
			"fork":                      struct{}{},
			"fork and knife with plate": struct{}{},
			"knife":                     struct{}{},
			"plate":                     struct{}{},
		},
	},
	":fork_and_knife:": Code{
		"\U0001f374",
		map[string]struct{}{
			"cooking":        struct{}{},
			"fork":           struct{}{},
			"fork and knife": struct{}{},
			"knife":          struct{}{},
		},
	},
	":spoon:": Code{
		"\U0001f944",
		map[string]struct{}{
			"spoon":     struct{}{},
			"tableware": struct{}{},
		},
	},
	":knife:": Code{
		"\U0001f52a",
		map[string]struct{}{
			"cooking":       struct{}{},
			"hocho":         struct{}{},
			"kitchen knife": struct{}{},
			"knife":         struct{}{},
			"tool":          struct{}{},
			"weapon":        struct{}{},
		},
	},
	":amphora:": Code{
		"\U0001f3fa",
		map[string]struct{}{
			"amphora":  struct{}{},
			"Aquarius": struct{}{},
			"cooking":  struct{}{},
			"drink":    struct{}{},
			"jug":      struct{}{},
			"tool":     struct{}{},
			"weapon":   struct{}{},
			"zodiac":   struct{}{},
		},
	},
	":earth_africa:": Code{
		"\U0001f30d",
		map[string]struct{}{
			"Africa":                      struct{}{},
			"earth":                       struct{}{},
			"Europe":                      struct{}{},
			"globe":                       struct{}{},
			"globe showing Europe-Africa": struct{}{},
			"world":                       struct{}{},
		},
	},
	":earth_americas:": Code{
		"\U0001f30e",
		map[string]struct{}{
			"Americas":               struct{}{},
			"earth":                  struct{}{},
			"globe":                  struct{}{},
			"globe showing Americas": struct{}{},
			"world":                  struct{}{},
		},
	},
	":earth_asia:": Code{
		"\U0001f30f",
		map[string]struct{}{
			"Asia":                         struct{}{},
			"Australia":                    struct{}{},
			"earth":                        struct{}{},
			"globe":                        struct{}{},
			"globe showing Asia-Australia": struct{}{},
			"world":                        struct{}{},
		},
	},
	":globe_with_meridians:": Code{
		"\U0001f310",
		map[string]struct{}{
			"earth":                struct{}{},
			"globe":                struct{}{},
			"globe with meridians": struct{}{},
			"meridians":            struct{}{},
			"world":                struct{}{},
		},
	},
	":map:": Code{
		"\U0001f5fa",
		map[string]struct{}{
			"map":       struct{}{},
			"world":     struct{}{},
			"world map": struct{}{},
		},
	},
	":japan:": Code{
		"\U0001f5fe",
		map[string]struct{}{
			"Japan":        struct{}{},
			"map":          struct{}{},
			"map of Japan": struct{}{},
		},
	},
	":mountain_snow:": Code{
		"\U0001f3d4",
		map[string]struct{}{
			"cold":                 struct{}{},
			"mountain":             struct{}{},
			"snow":                 struct{}{},
			"snow-capped mountain": struct{}{},
		},
	},
	":mountain:": Code{
		"\u26f0",
		map[string]struct{}{
			"mountain": struct{}{},
		},
	},
	":volcano:": Code{
		"\U0001f30b",
		map[string]struct{}{
			"eruption": struct{}{},
			"mountain": struct{}{},
			"volcano":  struct{}{},
		},
	},
	":mount_fuji:": Code{
		"\U0001f5fb",
		map[string]struct{}{
			"fuji":       struct{}{},
			"mount fuji": struct{}{},
			"mountain":   struct{}{},
		},
	},
	":camping:": Code{
		"\U0001f3d5",
		map[string]struct{}{
			"camping": struct{}{},
		},
	},
	":beach:": Code{
		"\U0001f3d6",
		map[string]struct{}{
			"beach":               struct{}{},
			"beach with umbrella": struct{}{},
			"umbrella":            struct{}{},
		},
	},
	":desert:": Code{
		"\U0001f3dc",
		map[string]struct{}{
			"desert": struct{}{},
		},
	},
	":island:": Code{
		"\U0001f3dd",
		map[string]struct{}{
			"desert":        struct{}{},
			"desert island": struct{}{},
			"island":        struct{}{},
		},
	},
	":park:": Code{
		"\U0001f3de",
		map[string]struct{}{
			"national park": struct{}{},
			"park":          struct{}{},
		},
	},
	":stadium:": Code{
		"\U0001f3df",
		map[string]struct{}{
			"stadium": struct{}{},
		},
	},
	":classical_building:": Code{
		"\U0001f3db",
		map[string]struct{}{
			"classical":          struct{}{},
			"classical building": struct{}{},
		},
	},
	":construction_site:": Code{
		"\U0001f3d7",
		map[string]struct{}{
			"building construction": struct{}{},
			"construction":          struct{}{},
		},
	},
	":homes:": Code{
		"\U0001f3d8",
		map[string]struct{}{
			"houses": struct{}{},
		},
	},
	":cityscape:": Code{
		"\U0001f3d9",
		map[string]struct{}{
			"city":      struct{}{},
			"cityscape": struct{}{},
		},
	},
	":house_abandoned:": Code{
		"\U0001f3da",
		map[string]struct{}{
			"derelict":       struct{}{},
			"derelict house": struct{}{},
			"house":          struct{}{},
		},
	},
	":house:": Code{
		"\U0001f3e0",
		map[string]struct{}{
			"home":  struct{}{},
			"house": struct{}{},
		},
	},
	":house_with_garden:": Code{
		"\U0001f3e1",
		map[string]struct{}{
			"garden":            struct{}{},
			"home":              struct{}{},
			"house":             struct{}{},
			"house with garden": struct{}{},
		},
	},
	":office:": Code{
		"\U0001f3e2",
		map[string]struct{}{
			"building":        struct{}{},
			"office building": struct{}{},
		},
	},
	":post_office:": Code{
		"\U0001f3e3",
		map[string]struct{}{
			"Japanese":             struct{}{},
			"Japanese post office": struct{}{},
			"post":                 struct{}{},
		},
	},
	":european_post_office:": Code{
		"\U0001f3e4",
		map[string]struct{}{
			"European":    struct{}{},
			"post":        struct{}{},
			"post office": struct{}{},
		},
	},
	":hospital:": Code{
		"\U0001f3e5",
		map[string]struct{}{
			"doctor":   struct{}{},
			"hospital": struct{}{},
			"medicine": struct{}{},
		},
	},
	":bank:": Code{
		"\U0001f3e6",
		map[string]struct{}{
			"bank":     struct{}{},
			"building": struct{}{},
		},
	},
	":hotel:": Code{
		"\U0001f3e8",
		map[string]struct{}{
			"building": struct{}{},
			"hotel":    struct{}{},
		},
	},
	":love_hotel:": Code{
		"\U0001f3e9",
		map[string]struct{}{
			"hotel":      struct{}{},
			"love":       struct{}{},
			"love hotel": struct{}{},
		},
	},
	":convenience_store:": Code{
		"\U0001f3ea",
		map[string]struct{}{
			"convenience":       struct{}{},
			"convenience store": struct{}{},
			"store":             struct{}{},
		},
	},
	":school:": Code{
		"\U0001f3eb",
		map[string]struct{}{
			"building": struct{}{},
			"school":   struct{}{},
		},
	},
	":department_store:": Code{
		"\U0001f3ec",
		map[string]struct{}{
			"department":       struct{}{},
			"department store": struct{}{},
			"store":            struct{}{},
		},
	},
	":factory:": Code{
		"\U0001f3ed",
		map[string]struct{}{
			"building": struct{}{},
			"factory":  struct{}{},
		},
	},
	":japanese_castle:": Code{
		"\U0001f3ef",
		map[string]struct{}{
			"castle":          struct{}{},
			"Japanese":        struct{}{},
			"Japanese castle": struct{}{},
		},
	},
	":european_castle:": Code{
		"\U0001f3f0",
		map[string]struct{}{
			"castle":   struct{}{},
			"European": struct{}{},
		},
	},
	":wedding:": Code{
		"\U0001f492",
		map[string]struct{}{
			"chapel":  struct{}{},
			"romance": struct{}{},
			"wedding": struct{}{},
		},
	},
	":tokyo_tower:": Code{
		"\U0001f5fc",
		map[string]struct{}{
			"Tokyo":       struct{}{},
			"Tokyo tower": struct{}{},
			"tower":       struct{}{},
		},
	},
	":statue_of_liberty:": Code{
		"\U0001f5fd",
		map[string]struct{}{
			"liberty":           struct{}{},
			"statue":            struct{}{},
			"Statue of Liberty": struct{}{},
		},
	},
	":church:": Code{
		"\u26ea",
		map[string]struct{}{
			"Christian": struct{}{},
			"church":    struct{}{},
			"cross":     struct{}{},
			"religion":  struct{}{},
		},
	},
	":mosque:": Code{
		"\U0001f54c",
		map[string]struct{}{
			"islam":    struct{}{},
			"mosque":   struct{}{},
			"Muslim":   struct{}{},
			"religion": struct{}{},
		},
	},
	":synagogue:": Code{
		"\U0001f54d",
		map[string]struct{}{
			"Jew":       struct{}{},
			"Jewish":    struct{}{},
			"religion":  struct{}{},
			"synagogue": struct{}{},
			"temple":    struct{}{},
		},
	},
	":shinto_shrine:": Code{
		"\u26e9",
		map[string]struct{}{
			"religion":      struct{}{},
			"shinto":        struct{}{},
			"shinto shrine": struct{}{},
			"shrine":        struct{}{},
		},
	},
	":kaaba:": Code{
		"\U0001f54b",
		map[string]struct{}{
			"islam":    struct{}{},
			"kaaba":    struct{}{},
			"Muslim":   struct{}{},
			"religion": struct{}{},
		},
	},
	":fountain:": Code{
		"\u26f2",
		map[string]struct{}{
			"fountain": struct{}{},
		},
	},
	":tent:": Code{
		"\u26fa",
		map[string]struct{}{
			"camping": struct{}{},
			"tent":    struct{}{},
		},
	},
	":foggy:": Code{
		"\U0001f301",
		map[string]struct{}{
			"fog":   struct{}{},
			"foggy": struct{}{},
		},
	},
	":night_with_stars:": Code{
		"\U0001f303",
		map[string]struct{}{
			"night":            struct{}{},
			"night with stars": struct{}{},
			"star":             struct{}{},
		},
	},
	":sunrise_over_mountains:": Code{
		"\U0001f304",
		map[string]struct{}{
			"morning":                struct{}{},
			"mountain":               struct{}{},
			"sun":                    struct{}{},
			"sunrise":                struct{}{},
			"sunrise over mountains": struct{}{},
		},
	},
	":sunrise:": Code{
		"\U0001f305",
		map[string]struct{}{
			"morning": struct{}{},
			"sun":     struct{}{},
			"sunrise": struct{}{},
		},
	},
	":city_dusk:": Code{
		"\U0001f306",
		map[string]struct{}{
			"city":              struct{}{},
			"cityscape at dusk": struct{}{},
			"dusk":              struct{}{},
			"evening":           struct{}{},
			"landscape":         struct{}{},
			"sun":               struct{}{},
			"sunset":            struct{}{},
		},
	},
	":city_sunset:": Code{
		"\U0001f307",
		map[string]struct{}{
			"dusk":   struct{}{},
			"sun":    struct{}{},
			"sunset": struct{}{},
		},
	},
	":bridge_at_night:": Code{
		"\U0001f309",
		map[string]struct{}{
			"bridge":          struct{}{},
			"bridge at night": struct{}{},
			"night":           struct{}{},
		},
	},
	":hotsprings:": Code{
		"\u2668",
		map[string]struct{}{
			"hot":         struct{}{},
			"hot springs": struct{}{},
			"hotsprings":  struct{}{},
			"springs":     struct{}{},
			"steaming":    struct{}{},
		},
	},
	":milky_way:": Code{
		"\U0001f30c",
		map[string]struct{}{
			"milky way": struct{}{},
			"space":     struct{}{},
		},
	},
	":carousel_horse:": Code{
		"\U0001f3a0",
		map[string]struct{}{
			"carousel":       struct{}{},
			"carousel horse": struct{}{},
			"horse":          struct{}{},
		},
	},
	":ferris_wheel:": Code{
		"\U0001f3a1",
		map[string]struct{}{
			"amusement park": struct{}{},
			"ferris":         struct{}{},
			"ferris wheel":   struct{}{},
			"wheel":          struct{}{},
		},
	},
	":roller_coaster:": Code{
		"\U0001f3a2",
		map[string]struct{}{
			"amusement park": struct{}{},
			"coaster":        struct{}{},
			"roller":         struct{}{},
			"roller coaster": struct{}{},
		},
	},
	":barber:": Code{
		"\U0001f488",
		map[string]struct{}{
			"barber":      struct{}{},
			"barber pole": struct{}{},
			"haircut":     struct{}{},
			"pole":        struct{}{},
		},
	},
	":circus_tent:": Code{
		"\U0001f3aa",
		map[string]struct{}{
			"circus":      struct{}{},
			"circus tent": struct{}{},
			"tent":        struct{}{},
		},
	},
	":performing_arts:": Code{
		"\U0001f3ad",
		map[string]struct{}{
			"art":             struct{}{},
			"mask":            struct{}{},
			"performing":      struct{}{},
			"performing arts": struct{}{},
			"theater":         struct{}{},
			"theatre":         struct{}{},
		},
	},
	":frame_photo:": Code{
		"\U0001f5bc",
		map[string]struct{}{
			"art":            struct{}{},
			"frame":          struct{}{},
			"framed picture": struct{}{},
			"museum":         struct{}{},
			"painting":       struct{}{},
			"picture":        struct{}{},
		},
	},
	":art:": Code{
		"\U0001f3a8",
		map[string]struct{}{
			"art":            struct{}{},
			"artist palette": struct{}{},
			"museum":         struct{}{},
			"painting":       struct{}{},
			"palette":        struct{}{},
		},
	},
	":slot_machine:": Code{
		"\U0001f3b0",
		map[string]struct{}{
			"game":         struct{}{},
			"slot":         struct{}{},
			"slot machine": struct{}{},
		},
	},
	":steam_locomotive:": Code{
		"\U0001f682",
		map[string]struct{}{
			"engine":     struct{}{},
			"locomotive": struct{}{},
			"railway":    struct{}{},
			"steam":      struct{}{},
			"train":      struct{}{},
		},
	},
	":railway_car:": Code{
		"\U0001f683",
		map[string]struct{}{
			"car":         struct{}{},
			"electric":    struct{}{},
			"railway":     struct{}{},
			"railway car": struct{}{},
			"train":       struct{}{},
			"tram":        struct{}{},
			"trolleybus":  struct{}{},
		},
	},
	":bullettrain_side:": Code{
		"\U0001f684",
		map[string]struct{}{
			"high-speed train": struct{}{},
			"railway":          struct{}{},
			"shinkansen":       struct{}{},
			"speed":            struct{}{},
			"train":            struct{}{},
		},
	},
	":bullettrain_front:": Code{
		"\U0001f685",
		map[string]struct{}{
			"bullet":       struct{}{},
			"bullet train": struct{}{},
			"railway":      struct{}{},
			"shinkansen":   struct{}{},
			"speed":        struct{}{},
			"train":        struct{}{},
		},
	},
	":train2:": Code{
		"\U0001f686",
		map[string]struct{}{
			"railway": struct{}{},
			"train":   struct{}{},
		},
	},
	":metro:": Code{
		"\U0001f687",
		map[string]struct{}{
			"metro":  struct{}{},
			"subway": struct{}{},
		},
	},
	":light_rail:": Code{
		"\U0001f688",
		map[string]struct{}{
			"light rail": struct{}{},
			"railway":    struct{}{},
		},
	},
	":station:": Code{
		"\U0001f689",
		map[string]struct{}{
			"railway": struct{}{},
			"station": struct{}{},
			"train":   struct{}{},
		},
	},
	":tram:": Code{
		"\U0001f68a",
		map[string]struct{}{
			"tram":       struct{}{},
			"trolleybus": struct{}{},
		},
	},
	":monorail:": Code{
		"\U0001f69d",
		map[string]struct{}{
			"monorail": struct{}{},
			"vehicle":  struct{}{},
		},
	},
	":mountain_railway:": Code{
		"\U0001f69e",
		map[string]struct{}{
			"car":              struct{}{},
			"mountain":         struct{}{},
			"mountain railway": struct{}{},
			"railway":          struct{}{},
		},
	},
	":train:": Code{
		"\U0001f68b",
		map[string]struct{}{
			"car":        struct{}{},
			"tram":       struct{}{},
			"tram car":   struct{}{},
			"trolleybus": struct{}{},
		},
	},
	":bus:": Code{
		"\U0001f68c",
		map[string]struct{}{
			"bus":     struct{}{},
			"vehicle": struct{}{},
		},
	},
	":oncoming_bus:": Code{
		"\U0001f68d",
		map[string]struct{}{
			"bus":          struct{}{},
			"oncoming":     struct{}{},
			"oncoming bus": struct{}{},
		},
	},
	":trolleybus:": Code{
		"\U0001f68e",
		map[string]struct{}{
			"bus":        struct{}{},
			"tram":       struct{}{},
			"trolley":    struct{}{},
			"trolleybus": struct{}{},
		},
	},
	":minibus:": Code{
		"\U0001f690",
		map[string]struct{}{
			"bus":     struct{}{},
			"minibus": struct{}{},
		},
	},
	":ambulance:": Code{
		"\U0001f691",
		map[string]struct{}{
			"ambulance": struct{}{},
			"vehicle":   struct{}{},
		},
	},
	":fire_engine:": Code{
		"\U0001f692",
		map[string]struct{}{
			"engine":      struct{}{},
			"fire":        struct{}{},
			"fire engine": struct{}{},
			"truck":       struct{}{},
		},
	},
	":police_car:": Code{
		"\U0001f693",
		map[string]struct{}{
			"car":        struct{}{},
			"patrol":     struct{}{},
			"police":     struct{}{},
			"police car": struct{}{},
		},
	},
	":oncoming_police_car:": Code{
		"\U0001f694",
		map[string]struct{}{
			"car":                 struct{}{},
			"oncoming":            struct{}{},
			"oncoming police car": struct{}{},
			"police":              struct{}{},
		},
	},
	":taxi:": Code{
		"\U0001f695",
		map[string]struct{}{
			"taxi":    struct{}{},
			"vehicle": struct{}{},
		},
	},
	":oncoming_taxi:": Code{
		"\U0001f696",
		map[string]struct{}{
			"oncoming":      struct{}{},
			"oncoming taxi": struct{}{},
			"taxi":          struct{}{},
		},
	},
	":red_car:": Code{
		"\U0001f697",
		map[string]struct{}{
			"automobile": struct{}{},
			"car":        struct{}{},
		},
	},
	":oncoming_automobile:": Code{
		"\U0001f698",
		map[string]struct{}{
			"automobile":          struct{}{},
			"car":                 struct{}{},
			"oncoming":            struct{}{},
			"oncoming automobile": struct{}{},
		},
	},
	":blue_car:": Code{
		"\U0001f699",
		map[string]struct{}{
			"recreational":          struct{}{},
			"sport utility":         struct{}{},
			"sport utility vehicle": struct{}{},
		},
	},
	":truck:": Code{
		"\U0001f69a",
		map[string]struct{}{
			"delivery":       struct{}{},
			"delivery truck": struct{}{},
			"truck":          struct{}{},
		},
	},
	":articulated_lorry:": Code{
		"\U0001f69b",
		map[string]struct{}{
			"articulated lorry": struct{}{},
			"lorry":             struct{}{},
			"semi":              struct{}{},
			"truck":             struct{}{},
		},
	},
	":tractor:": Code{
		"\U0001f69c",
		map[string]struct{}{
			"tractor": struct{}{},
			"vehicle": struct{}{},
		},
	},
	":bike:": Code{
		"\U0001f6b2",
		map[string]struct{}{
			"bicycle": struct{}{},
			"bike":    struct{}{},
		},
	},
	":scooter:": Code{
		"\U0001f6f4",
		map[string]struct{}{
			"kick":         struct{}{},
			"kick scooter": struct{}{},
			"scooter":      struct{}{},
		},
	},
	":motor_scooter:": Code{
		"\U0001f6f5",
		map[string]struct{}{
			"motor":         struct{}{},
			"motor scooter": struct{}{},
			"scooter":       struct{}{},
		},
	},
	":busstop:": Code{
		"\U0001f68f",
		map[string]struct{}{
			"bus":      struct{}{},
			"bus stop": struct{}{},
			"busstop":  struct{}{},
			"stop":     struct{}{},
		},
	},
	":motorway:": Code{
		"\U0001f6e3",
		map[string]struct{}{
			"highway":  struct{}{},
			"motorway": struct{}{},
			"road":     struct{}{},
		},
	},
	":railway_track:": Code{
		"\U0001f6e4",
		map[string]struct{}{
			"railway":       struct{}{},
			"railway track": struct{}{},
			"train":         struct{}{},
		},
	},
	":fuelpump:": Code{
		"\u26fd",
		map[string]struct{}{
			"fuel":      struct{}{},
			"fuel pump": struct{}{},
			"fuelpump":  struct{}{},
			"gas":       struct{}{},
			"pump":      struct{}{},
			"station":   struct{}{},
		},
	},
	":rotating_light:": Code{
		"\U0001f6a8",
		map[string]struct{}{
			"beacon":           struct{}{},
			"car":              struct{}{},
			"light":            struct{}{},
			"police":           struct{}{},
			"police car light": struct{}{},
			"revolving":        struct{}{},
		},
	},
	":traffic_light:": Code{
		"\U0001f6a5",
		map[string]struct{}{
			"horizontal traffic light": struct{}{},
			"light":                    struct{}{},
			"signal":                   struct{}{},
			"traffic":                  struct{}{},
		},
	},
	":vertical_traffic_light:": Code{
		"\U0001f6a6",
		map[string]struct{}{
			"light":                  struct{}{},
			"signal":                 struct{}{},
			"traffic":                struct{}{},
			"vertical traffic light": struct{}{},
		},
	},
	":construction:": Code{
		"\U0001f6a7",
		map[string]struct{}{
			"barrier":      struct{}{},
			"construction": struct{}{},
		},
	},
	":octagonal_sign:": Code{
		"\U0001f6d1",
		map[string]struct{}{
			"octagonal": struct{}{},
			"sign":      struct{}{},
			"stop":      struct{}{},
			"stop sign": struct{}{},
		},
	},
	":anchor:": Code{
		"\u2693",
		map[string]struct{}{
			"anchor": struct{}{},
			"ship":   struct{}{},
			"tool":   struct{}{},
		},
	},
	":sailboat:": Code{
		"\u26f5",
		map[string]struct{}{
			"boat":     struct{}{},
			"resort":   struct{}{},
			"sailboat": struct{}{},
			"sea":      struct{}{},
			"yacht":    struct{}{},
		},
	},
	":canoe:": Code{
		"\U0001f6f6",
		map[string]struct{}{
			"boat":  struct{}{},
			"canoe": struct{}{},
		},
	},
	":speedboat:": Code{
		"\U0001f6a4",
		map[string]struct{}{
			"boat":      struct{}{},
			"speedboat": struct{}{},
		},
	},
	":cruise_ship:": Code{
		"\U0001f6f3",
		map[string]struct{}{
			"passenger":      struct{}{},
			"passenger ship": struct{}{},
			"ship":           struct{}{},
		},
	},
	":ferry:": Code{
		"\u26f4",
		map[string]struct{}{
			"boat":      struct{}{},
			"ferry":     struct{}{},
			"passenger": struct{}{},
		},
	},
	":motorboat:": Code{
		"\U0001f6e5",
		map[string]struct{}{
			"boat":       struct{}{},
			"motor boat": struct{}{},
			"motorboat":  struct{}{},
		},
	},
	":ship:": Code{
		"\U0001f6a2",
		map[string]struct{}{
			"boat":      struct{}{},
			"passenger": struct{}{},
			"ship":      struct{}{},
		},
	},
	":airplane:": Code{
		"\u2708",
		map[string]struct{}{
			"aeroplane": struct{}{},
			"airplane":  struct{}{},
		},
	},
	":airplane_small:": Code{
		"\U0001f6e9",
		map[string]struct{}{
			"aeroplane":      struct{}{},
			"airplane":       struct{}{},
			"small airplane": struct{}{},
		},
	},
	":airplane_departure:": Code{
		"\U0001f6eb",
		map[string]struct{}{
			"aeroplane":          struct{}{},
			"airplane":           struct{}{},
			"airplane departure": struct{}{},
			"check-in":           struct{}{},
			"departure":          struct{}{},
			"departures":         struct{}{},
		},
	},
	":airplane_arriving:": Code{
		"\U0001f6ec",
		map[string]struct{}{
			"aeroplane":        struct{}{},
			"airplane":         struct{}{},
			"airplane arrival": struct{}{},
			"arrivals":         struct{}{},
			"arriving":         struct{}{},
			"landing":          struct{}{},
		},
	},
	":seat:": Code{
		"\U0001f4ba",
		map[string]struct{}{
			"chair": struct{}{},
			"seat":  struct{}{},
		},
	},
	":helicopter:": Code{
		"\U0001f681",
		map[string]struct{}{
			"helicopter": struct{}{},
			"vehicle":    struct{}{},
		},
	},
	":suspension_railway:": Code{
		"\U0001f69f",
		map[string]struct{}{
			"railway":            struct{}{},
			"suspension":         struct{}{},
			"suspension railway": struct{}{},
		},
	},
	":mountain_cableway:": Code{
		"\U0001f6a0",
		map[string]struct{}{
			"cable":             struct{}{},
			"gondola":           struct{}{},
			"mountain":          struct{}{},
			"mountain cableway": struct{}{},
		},
	},
	":aerial_tramway:": Code{
		"\U0001f6a1",
		map[string]struct{}{
			"aerial":         struct{}{},
			"aerial tramway": struct{}{},
			"cable":          struct{}{},
			"car":            struct{}{},
			"gondola":        struct{}{},
			"tramway":        struct{}{},
		},
	},
	":satellite_orbital:": Code{
		"\U0001f6f0",
		map[string]struct{}{
			"satellite": struct{}{},
			"space":     struct{}{},
		},
	},
	":rocket:": Code{
		"\U0001f680",
		map[string]struct{}{
			"rocket": struct{}{},
			"space":  struct{}{},
		},
	},
	":flying_saucer:": Code{
		"\U0001f6f8",
		map[string]struct{}{
			"flying saucer": struct{}{},
			"UFO":           struct{}{},
		},
	},
	":bellhop:": Code{
		"\U0001f6ce",
		map[string]struct{}{
			"bell":         struct{}{},
			"bellhop":      struct{}{},
			"bellhop bell": struct{}{},
			"hotel":        struct{}{},
		},
	},
	":door:": Code{
		"\U0001f6aa",
		map[string]struct{}{
			"door": struct{}{},
		},
	},
	":bed:": Code{
		"\U0001f6cf",
		map[string]struct{}{
			"bed":   struct{}{},
			"hotel": struct{}{},
			"sleep": struct{}{},
		},
	},
	":couch:": Code{
		"\U0001f6cb",
		map[string]struct{}{
			"couch":          struct{}{},
			"couch and lamp": struct{}{},
			"hotel":          struct{}{},
			"lamp":           struct{}{},
		},
	},
	":toilet:": Code{
		"\U0001f6bd",
		map[string]struct{}{
			"toilet": struct{}{},
		},
	},
	":shower:": Code{
		"\U0001f6bf",
		map[string]struct{}{
			"shower": struct{}{},
			"water":  struct{}{},
		},
	},
	":bathtub:": Code{
		"\U0001f6c1",
		map[string]struct{}{
			"bath":    struct{}{},
			"bathtub": struct{}{},
		},
	},
	":hourglass:": Code{
		"\u231b",
		map[string]struct{}{
			"hourglass done": struct{}{},
			"sand":           struct{}{},
			"timer":          struct{}{},
		},
	},
	":hourglass_flowing_sand:": Code{
		"\u23f3",
		map[string]struct{}{
			"hourglass":          struct{}{},
			"hourglass not done": struct{}{},
			"sand":               struct{}{},
			"timer":              struct{}{},
		},
	},
	":watch:": Code{
		"\u231a",
		map[string]struct{}{
			"clock": struct{}{},
			"watch": struct{}{},
		},
	},
	":alarm_clock:": Code{
		"\u23f0",
		map[string]struct{}{
			"alarm":       struct{}{},
			"alarm clock": struct{}{},
			"clock":       struct{}{},
		},
	},
	":stopwatch:": Code{
		"\u23f1",
		map[string]struct{}{
			"clock":     struct{}{},
			"stopwatch": struct{}{},
		},
	},
	":timer:": Code{
		"\u23f2",
		map[string]struct{}{
			"clock":       struct{}{},
			"timer":       struct{}{},
			"timer clock": struct{}{},
		},
	},
	":clock:": Code{
		"\U0001f570",
		map[string]struct{}{
			"clock":             struct{}{},
			"mantelpiece clock": struct{}{},
		},
	},
	":clock12:": Code{
		"\U0001f55b",
		map[string]struct{}{
			"00":             struct{}{},
			"12":             struct{}{},
			"12:00":          struct{}{},
			"clock":          struct{}{},
			"o’clock":        struct{}{},
			"twelve":         struct{}{},
			"twelve o’clock": struct{}{},
		},
	},
	":clock1230:": Code{
		"\U0001f567",
		map[string]struct{}{
			"12":            struct{}{},
			"12:30":         struct{}{},
			"30":            struct{}{},
			"clock":         struct{}{},
			"thirty":        struct{}{},
			"twelve":        struct{}{},
			"twelve-thirty": struct{}{},
		},
	},
	":clock1:": Code{
		"\U0001f550",
		map[string]struct{}{
			"00":          struct{}{},
			"1":           struct{}{},
			"1:00":        struct{}{},
			"clock":       struct{}{},
			"o’clock":     struct{}{},
			"one":         struct{}{},
			"one o’clock": struct{}{},
		},
	},
	":clock130:": Code{
		"\U0001f55c",
		map[string]struct{}{
			"1":          struct{}{},
			"1:30":       struct{}{},
			"30":         struct{}{},
			"clock":      struct{}{},
			"one":        struct{}{},
			"one-thirty": struct{}{},
			"thirty":     struct{}{},
		},
	},
	":clock2:": Code{
		"\U0001f551",
		map[string]struct{}{
			"00":          struct{}{},
			"2":           struct{}{},
			"2:00":        struct{}{},
			"clock":       struct{}{},
			"o’clock":     struct{}{},
			"two":         struct{}{},
			"two o’clock": struct{}{},
		},
	},
	":clock230:": Code{
		"\U0001f55d",
		map[string]struct{}{
			"2":          struct{}{},
			"2:30":       struct{}{},
			"30":         struct{}{},
			"clock":      struct{}{},
			"thirty":     struct{}{},
			"two":        struct{}{},
			"two-thirty": struct{}{},
		},
	},
	":clock3:": Code{
		"\U0001f552",
		map[string]struct{}{
			"00":            struct{}{},
			"3":             struct{}{},
			"3:00":          struct{}{},
			"clock":         struct{}{},
			"o’clock":       struct{}{},
			"three":         struct{}{},
			"three o’clock": struct{}{},
		},
	},
	":clock330:": Code{
		"\U0001f55e",
		map[string]struct{}{
			"3":            struct{}{},
			"3:30":         struct{}{},
			"30":           struct{}{},
			"clock":        struct{}{},
			"thirty":       struct{}{},
			"three":        struct{}{},
			"three-thirty": struct{}{},
		},
	},
	":clock4:": Code{
		"\U0001f553",
		map[string]struct{}{
			"00":           struct{}{},
			"4":            struct{}{},
			"4:00":         struct{}{},
			"clock":        struct{}{},
			"four":         struct{}{},
			"four o’clock": struct{}{},
			"o’clock":      struct{}{},
		},
	},
	":clock430:": Code{
		"\U0001f55f",
		map[string]struct{}{
			"30":          struct{}{},
			"4":           struct{}{},
			"4:30":        struct{}{},
			"clock":       struct{}{},
			"four":        struct{}{},
			"four-thirty": struct{}{},
			"thirty":      struct{}{},
		},
	},
	":clock5:": Code{
		"\U0001f554",
		map[string]struct{}{
			"00":           struct{}{},
			"5":            struct{}{},
			"5:00":         struct{}{},
			"clock":        struct{}{},
			"five":         struct{}{},
			"five o’clock": struct{}{},
			"o’clock":      struct{}{},
		},
	},
	":clock530:": Code{
		"\U0001f560",
		map[string]struct{}{
			"30":          struct{}{},
			"5":           struct{}{},
			"5:30":        struct{}{},
			"clock":       struct{}{},
			"five":        struct{}{},
			"five-thirty": struct{}{},
			"thirty":      struct{}{},
		},
	},
	":clock6:": Code{
		"\U0001f555",
		map[string]struct{}{
			"00":          struct{}{},
			"6":           struct{}{},
			"6:00":        struct{}{},
			"clock":       struct{}{},
			"o’clock":     struct{}{},
			"six":         struct{}{},
			"six o’clock": struct{}{},
		},
	},
	":clock630:": Code{
		"\U0001f561",
		map[string]struct{}{
			"30":         struct{}{},
			"6":          struct{}{},
			"6:30":       struct{}{},
			"clock":      struct{}{},
			"six":        struct{}{},
			"six-thirty": struct{}{},
			"thirty":     struct{}{},
		},
	},
	":clock7:": Code{
		"\U0001f556",
		map[string]struct{}{
			"00":            struct{}{},
			"7":             struct{}{},
			"7:00":          struct{}{},
			"clock":         struct{}{},
			"o’clock":       struct{}{},
			"seven":         struct{}{},
			"seven o’clock": struct{}{},
		},
	},
	":clock730:": Code{
		"\U0001f562",
		map[string]struct{}{
			"30":           struct{}{},
			"7":            struct{}{},
			"7:30":         struct{}{},
			"clock":        struct{}{},
			"seven":        struct{}{},
			"seven-thirty": struct{}{},
			"thirty":       struct{}{},
		},
	},
	":clock8:": Code{
		"\U0001f557",
		map[string]struct{}{
			"00":            struct{}{},
			"8":             struct{}{},
			"8:00":          struct{}{},
			"clock":         struct{}{},
			"eight":         struct{}{},
			"eight o’clock": struct{}{},
			"o’clock":       struct{}{},
		},
	},
	":clock830:": Code{
		"\U0001f563",
		map[string]struct{}{
			"30":           struct{}{},
			"8":            struct{}{},
			"8:30":         struct{}{},
			"clock":        struct{}{},
			"eight":        struct{}{},
			"eight-thirty": struct{}{},
			"thirty":       struct{}{},
		},
	},
	":clock9:": Code{
		"\U0001f558",
		map[string]struct{}{
			"00":           struct{}{},
			"9":            struct{}{},
			"9:00":         struct{}{},
			"clock":        struct{}{},
			"nine":         struct{}{},
			"nine o’clock": struct{}{},
			"o’clock":      struct{}{},
		},
	},
	":clock930:": Code{
		"\U0001f564",
		map[string]struct{}{
			"30":          struct{}{},
			"9":           struct{}{},
			"9:30":        struct{}{},
			"clock":       struct{}{},
			"nine":        struct{}{},
			"nine-thirty": struct{}{},
			"thirty":      struct{}{},
		},
	},
	":clock10:": Code{
		"\U0001f559",
		map[string]struct{}{
			"00":          struct{}{},
			"10":          struct{}{},
			"10:00":       struct{}{},
			"clock":       struct{}{},
			"o’clock":     struct{}{},
			"ten":         struct{}{},
			"ten o’clock": struct{}{},
		},
	},
	":clock1030:": Code{
		"\U0001f565",
		map[string]struct{}{
			"10":         struct{}{},
			"10:30":      struct{}{},
			"30":         struct{}{},
			"clock":      struct{}{},
			"ten":        struct{}{},
			"ten-thirty": struct{}{},
			"thirty":     struct{}{},
		},
	},
	":clock11:": Code{
		"\U0001f55a",
		map[string]struct{}{
			"00":             struct{}{},
			"11":             struct{}{},
			"11:00":          struct{}{},
			"clock":          struct{}{},
			"eleven":         struct{}{},
			"eleven o’clock": struct{}{},
			"o’clock":        struct{}{},
		},
	},
	":clock1130:": Code{
		"\U0001f566",
		map[string]struct{}{
			"11":            struct{}{},
			"11:30":         struct{}{},
			"30":            struct{}{},
			"clock":         struct{}{},
			"eleven":        struct{}{},
			"eleven-thirty": struct{}{},
			"thirty":        struct{}{},
		},
	},
	":new_moon:": Code{
		"\U0001f311",
		map[string]struct{}{
			"dark":     struct{}{},
			"moon":     struct{}{},
			"new moon": struct{}{},
		},
	},
	":waxing_crescent_moon:": Code{
		"\U0001f312",
		map[string]struct{}{
			"crescent":             struct{}{},
			"moon":                 struct{}{},
			"waxing":               struct{}{},
			"waxing crescent moon": struct{}{},
		},
	},
	":first_quarter_moon:": Code{
		"\U0001f313",
		map[string]struct{}{
			"first quarter moon": struct{}{},
			"moon":               struct{}{},
			"quarter":            struct{}{},
		},
	},
	":waxing_gibbous_moon:": Code{
		"\U0001f314",
		map[string]struct{}{
			"gibbous":             struct{}{},
			"moon":                struct{}{},
			"waxing":              struct{}{},
			"waxing gibbous moon": struct{}{},
		},
	},
	":full_moon:": Code{
		"\U0001f315",
		map[string]struct{}{
			"full":      struct{}{},
			"full moon": struct{}{},
			"moon":      struct{}{},
		},
	},
	":waning_gibbous_moon:": Code{
		"\U0001f316",
		map[string]struct{}{
			"gibbous":             struct{}{},
			"moon":                struct{}{},
			"waning":              struct{}{},
			"waning gibbous moon": struct{}{},
		},
	},
	":last_quarter_moon:": Code{
		"\U0001f317",
		map[string]struct{}{
			"last quarter moon": struct{}{},
			"moon":              struct{}{},
			"quarter":           struct{}{},
		},
	},
	":waning_crescent_moon:": Code{
		"\U0001f318",
		map[string]struct{}{
			"crescent":             struct{}{},
			"moon":                 struct{}{},
			"waning":               struct{}{},
			"waning crescent moon": struct{}{},
		},
	},
	":crescent_moon:": Code{
		"\U0001f319",
		map[string]struct{}{
			"crescent":      struct{}{},
			"crescent moon": struct{}{},
			"moon":          struct{}{},
		},
	},
	":new_moon_with_face:": Code{
		"\U0001f31a",
		map[string]struct{}{
			"face":          struct{}{},
			"moon":          struct{}{},
			"new moon face": struct{}{},
		},
	},
	":first_quarter_moon_with_face:": Code{
		"\U0001f31b",
		map[string]struct{}{
			"face":                    struct{}{},
			"first quarter moon face": struct{}{},
			"moon":                    struct{}{},
			"quarter":                 struct{}{},
		},
	},
	":last_quarter_moon_with_face:": Code{
		"\U0001f31c",
		map[string]struct{}{
			"face":                   struct{}{},
			"last quarter moon face": struct{}{},
			"moon":                   struct{}{},
			"quarter":                struct{}{},
		},
	},
	":thermometer:": Code{
		"\U0001f321",
		map[string]struct{}{
			"thermometer": struct{}{},
			"weather":     struct{}{},
		},
	},
	":sunny:": Code{
		"\u2600",
		map[string]struct{}{
			"bright": struct{}{},
			"rays":   struct{}{},
			"sun":    struct{}{},
			"sunny":  struct{}{},
		},
	},
	":full_moon_with_face:": Code{
		"\U0001f31d",
		map[string]struct{}{
			"bright":         struct{}{},
			"face":           struct{}{},
			"full":           struct{}{},
			"full moon face": struct{}{},
			"moon":           struct{}{},
		},
	},
	":sun_with_face:": Code{
		"\U0001f31e",
		map[string]struct{}{
			"bright":        struct{}{},
			"face":          struct{}{},
			"sun":           struct{}{},
			"sun with face": struct{}{},
		},
	},
	":star:": Code{
		"\u2b50",
		map[string]struct{}{
			"star":              struct{}{},
			"white medium star": struct{}{},
		},
	},
	":star2:": Code{
		"\U0001f31f",
		map[string]struct{}{
			"glittery":     struct{}{},
			"glow":         struct{}{},
			"glowing star": struct{}{},
			"shining":      struct{}{},
			"sparkle":      struct{}{},
			"star":         struct{}{},
		},
	},
	":stars:": Code{
		"\U0001f320",
		map[string]struct{}{
			"falling":       struct{}{},
			"shooting":      struct{}{},
			"shooting star": struct{}{},
			"star":          struct{}{},
		},
	},
	":cloud:": Code{
		"\u2601",
		map[string]struct{}{
			"cloud":   struct{}{},
			"weather": struct{}{},
		},
	},
	":partly_sunny:": Code{
		"\u26c5",
		map[string]struct{}{
			"cloud":            struct{}{},
			"sun":              struct{}{},
			"sun behind cloud": struct{}{},
		},
	},
	":thunder_cloud_rain:": Code{
		"\u26c8",
		map[string]struct{}{
			"cloud":                         struct{}{},
			"cloud with lightning and rain": struct{}{},
			"rain":                          struct{}{},
			"thunder":                       struct{}{},
		},
	},
	":white_sun_small_cloud:": Code{
		"\U0001f324",
		map[string]struct{}{
			"cloud":                  struct{}{},
			"sun":                    struct{}{},
			"sun behind small cloud": struct{}{},
		},
	},
	":white_sun_cloud:": Code{
		"\U0001f325",
		map[string]struct{}{
			"cloud":                  struct{}{},
			"sun":                    struct{}{},
			"sun behind large cloud": struct{}{},
		},
	},
	":white_sun_rain_cloud:": Code{
		"\U0001f326",
		map[string]struct{}{
			"cloud":                 struct{}{},
			"rain":                  struct{}{},
			"sun":                   struct{}{},
			"sun behind rain cloud": struct{}{},
		},
	},
	":cloud_rain:": Code{
		"\U0001f327",
		map[string]struct{}{
			"cloud":           struct{}{},
			"cloud with rain": struct{}{},
			"rain":            struct{}{},
		},
	},
	":cloud_snow:": Code{
		"\U0001f328",
		map[string]struct{}{
			"cloud":           struct{}{},
			"cloud with snow": struct{}{},
			"cold":            struct{}{},
			"snow":            struct{}{},
		},
	},
	":cloud_lightning:": Code{
		"\U0001f329",
		map[string]struct{}{
			"cloud":                struct{}{},
			"cloud with lightning": struct{}{},
			"lightning":            struct{}{},
		},
	},
	":cloud_tornado:": Code{
		"\U0001f32a",
		map[string]struct{}{
			"cloud":     struct{}{},
			"tornado":   struct{}{},
			"whirlwind": struct{}{},
		},
	},
	":fog:": Code{
		"\U0001f32b",
		map[string]struct{}{
			"cloud": struct{}{},
			"fog":   struct{}{},
		},
	},
	":wind_blowing_face:": Code{
		"\U0001f32c",
		map[string]struct{}{
			"blow":      struct{}{},
			"cloud":     struct{}{},
			"face":      struct{}{},
			"wind":      struct{}{},
			"wind face": struct{}{},
		},
	},
	":cyclone:": Code{
		"\U0001f300",
		map[string]struct{}{
			"cyclone": struct{}{},
			"dizzy":   struct{}{},
			"twister": struct{}{},
			"typhoon": struct{}{},
		},
	},
	":rainbow:": Code{
		"\U0001f308",
		map[string]struct{}{
			"rain":    struct{}{},
			"rainbow": struct{}{},
		},
	},
	":closed_umbrella:": Code{
		"\U0001f302",
		map[string]struct{}{
			"closed umbrella": struct{}{},
			"clothing":        struct{}{},
			"rain":            struct{}{},
			"umbrella":        struct{}{},
		},
	},
	":umbrella2:": Code{
		"\u2602",
		map[string]struct{}{
			"clothing": struct{}{},
			"rain":     struct{}{},
			"umbrella": struct{}{},
		},
	},
	":umbrella:": Code{
		"\u2614",
		map[string]struct{}{
			"clothing":                 struct{}{},
			"drop":                     struct{}{},
			"rain":                     struct{}{},
			"umbrella":                 struct{}{},
			"umbrella with rain drops": struct{}{},
		},
	},
	":beach_umbrella:": Code{
		"\u26f1",
		map[string]struct{}{
			"rain":               struct{}{},
			"sun":                struct{}{},
			"umbrella":           struct{}{},
			"umbrella on ground": struct{}{},
		},
	},
	":zap:": Code{
		"\u26a1",
		map[string]struct{}{
			"danger":       struct{}{},
			"electric":     struct{}{},
			"electricity":  struct{}{},
			"high voltage": struct{}{},
			"lightning":    struct{}{},
			"voltage":      struct{}{},
			"zap":          struct{}{},
		},
	},
	":snowflake:": Code{
		"\u2744",
		map[string]struct{}{
			"cold":      struct{}{},
			"snow":      struct{}{},
			"snowflake": struct{}{},
		},
	},
	":snowman2:": Code{
		"\u2603",
		map[string]struct{}{
			"cold":    struct{}{},
			"snow":    struct{}{},
			"snowman": struct{}{},
		},
	},
	":snowman:": Code{
		"\u26c4",
		map[string]struct{}{
			"cold":                 struct{}{},
			"snow":                 struct{}{},
			"snowman":              struct{}{},
			"snowman without snow": struct{}{},
		},
	},
	":comet:": Code{
		"\u2604",
		map[string]struct{}{
			"comet": struct{}{},
			"space": struct{}{},
		},
	},
	":fire:": Code{
		"\U0001f525",
		map[string]struct{}{
			"fire":  struct{}{},
			"flame": struct{}{},
			"tool":  struct{}{},
		},
	},
	":droplet:": Code{
		"\U0001f4a7",
		map[string]struct{}{
			"cold":    struct{}{},
			"comic":   struct{}{},
			"drop":    struct{}{},
			"droplet": struct{}{},
			"sweat":   struct{}{},
		},
	},
	":ocean:": Code{
		"\U0001f30a",
		map[string]struct{}{
			"ocean":      struct{}{},
			"water":      struct{}{},
			"water wave": struct{}{},
			"wave":       struct{}{},
		},
	},
	":jack_o_lantern:": Code{
		"\U0001f383",
		map[string]struct{}{
			"celebration":    struct{}{},
			"halloween":      struct{}{},
			"jack":           struct{}{},
			"jack-o-lantern": struct{}{},
			"lantern":        struct{}{},
		},
	},
	":christmas_tree:": Code{
		"\U0001f384",
		map[string]struct{}{
			"celebration":    struct{}{},
			"Christmas":      struct{}{},
			"Christmas tree": struct{}{},
			"tree":           struct{}{},
		},
	},
	":fireworks:": Code{
		"\U0001f386",
		map[string]struct{}{
			"celebration": struct{}{},
			"fireworks":   struct{}{},
		},
	},
	":sparkler:": Code{
		"\U0001f387",
		map[string]struct{}{
			"celebration": struct{}{},
			"fireworks":   struct{}{},
			"sparkle":     struct{}{},
			"sparkler":    struct{}{},
		},
	},
	":sparkles:": Code{
		"\u2728",
		map[string]struct{}{
			"sparkle":  struct{}{},
			"sparkles": struct{}{},
			"star":     struct{}{},
		},
	},
	":balloon:": Code{
		"\U0001f388",
		map[string]struct{}{
			"balloon":     struct{}{},
			"celebration": struct{}{},
		},
	},
	":tada:": Code{
		"\U0001f389",
		map[string]struct{}{
			"celebration":  struct{}{},
			"party":        struct{}{},
			"party popper": struct{}{},
			"popper":       struct{}{},
			"tada":         struct{}{},
		},
	},
	":confetti_ball:": Code{
		"\U0001f38a",
		map[string]struct{}{
			"ball":          struct{}{},
			"celebration":   struct{}{},
			"confetti":      struct{}{},
			"confetti ball": struct{}{},
		},
	},
	":tanabata_tree:": Code{
		"\U0001f38b",
		map[string]struct{}{
			"banner":        struct{}{},
			"celebration":   struct{}{},
			"Japanese":      struct{}{},
			"tanabata tree": struct{}{},
			"tree":          struct{}{},
		},
	},
	":bamboo:": Code{
		"\U0001f38d",
		map[string]struct{}{
			"bamboo":          struct{}{},
			"celebration":     struct{}{},
			"Japanese":        struct{}{},
			"pine":            struct{}{},
			"pine decoration": struct{}{},
		},
	},
	":dolls:": Code{
		"\U0001f38e",
		map[string]struct{}{
			"celebration":    struct{}{},
			"doll":           struct{}{},
			"festival":       struct{}{},
			"Japanese":       struct{}{},
			"Japanese dolls": struct{}{},
		},
	},
	":flags:": Code{
		"\U0001f38f",
		map[string]struct{}{
			"carp":          struct{}{},
			"carp streamer": struct{}{},
			"celebration":   struct{}{},
			"streamer":      struct{}{},
		},
	},
	":wind_chime:": Code{
		"\U0001f390",
		map[string]struct{}{
			"bell":        struct{}{},
			"celebration": struct{}{},
			"chime":       struct{}{},
			"wind":        struct{}{},
			"wind chime":  struct{}{},
		},
	},
	":rice_scene:": Code{
		"\U0001f391",
		map[string]struct{}{
			"celebration":           struct{}{},
			"ceremony":              struct{}{},
			"moon":                  struct{}{},
			"moon viewing ceremony": struct{}{},
		},
	},
	":ribbon:": Code{
		"\U0001f380",
		map[string]struct{}{
			"celebration": struct{}{},
			"ribbon":      struct{}{},
		},
	},
	":gift:": Code{
		"\U0001f381",
		map[string]struct{}{
			"box":          struct{}{},
			"celebration":  struct{}{},
			"gift":         struct{}{},
			"present":      struct{}{},
			"wrapped":      struct{}{},
			"wrapped gift": struct{}{},
		},
	},
	":reminder_ribbon:": Code{
		"\U0001f397",
		map[string]struct{}{
			"celebration":     struct{}{},
			"reminder":        struct{}{},
			"reminder ribbon": struct{}{},
			"ribbon":          struct{}{},
		},
	},
	":tickets:": Code{
		"\U0001f39f",
		map[string]struct{}{
			"admission":         struct{}{},
			"admission tickets": struct{}{},
			"ticket":            struct{}{},
		},
	},
	":ticket:": Code{
		"\U0001f3ab",
		map[string]struct{}{
			"admission": struct{}{},
			"ticket":    struct{}{},
		},
	},
	":military_medal:": Code{
		"\U0001f396",
		map[string]struct{}{
			"celebration":    struct{}{},
			"medal":          struct{}{},
			"military":       struct{}{},
			"military medal": struct{}{},
		},
	},
	":trophy:": Code{
		"\U0001f3c6",
		map[string]struct{}{
			"prize":  struct{}{},
			"trophy": struct{}{},
		},
	},
	":medal:": Code{
		"\U0001f3c5",
		map[string]struct{}{
			"medal":        struct{}{},
			"sports medal": struct{}{},
		},
	},
	":first_place:": Code{
		"\U0001f947",
		map[string]struct{}{
			"1st place medal": struct{}{},
			"first":           struct{}{},
			"gold":            struct{}{},
			"medal":           struct{}{},
		},
	},
	":second_place:": Code{
		"\U0001f948",
		map[string]struct{}{
			"2nd place medal": struct{}{},
			"medal":           struct{}{},
			"second":          struct{}{},
			"silver":          struct{}{},
		},
	},
	":third_place:": Code{
		"\U0001f949",
		map[string]struct{}{
			"3rd place medal": struct{}{},
			"bronze":          struct{}{},
			"medal":           struct{}{},
			"third":           struct{}{},
		},
	},
	":soccer:": Code{
		"\u26bd",
		map[string]struct{}{
			"ball":        struct{}{},
			"football":    struct{}{},
			"soccer":      struct{}{},
			"soccer ball": struct{}{},
		},
	},
	":baseball:": Code{
		"\u26be",
		map[string]struct{}{
			"ball":     struct{}{},
			"baseball": struct{}{},
		},
	},
	":basketball:": Code{
		"\U0001f3c0",
		map[string]struct{}{
			"ball":       struct{}{},
			"basketball": struct{}{},
			"hoop":       struct{}{},
		},
	},
	":volleyball:": Code{
		"\U0001f3d0",
		map[string]struct{}{
			"ball":       struct{}{},
			"game":       struct{}{},
			"volleyball": struct{}{},
		},
	},
	":football:": Code{
		"\U0001f3c8",
		map[string]struct{}{
			"american":          struct{}{},
			"american football": struct{}{},
			"ball":              struct{}{},
			"football":          struct{}{},
		},
	},
	":rugby_football:": Code{
		"\U0001f3c9",
		map[string]struct{}{
			"ball":           struct{}{},
			"football":       struct{}{},
			"rugby":          struct{}{},
			"rugby football": struct{}{},
		},
	},
	":tennis:": Code{
		"\U0001f3be",
		map[string]struct{}{
			"ball":    struct{}{},
			"racquet": struct{}{},
			"tennis":  struct{}{},
		},
	},
	":8ball:": Code{
		"\U0001f3b1",
		map[string]struct{}{
			"8":           struct{}{},
			"8 ball":      struct{}{},
			"ball":        struct{}{},
			"billiard":    struct{}{},
			"eight":       struct{}{},
			"game":        struct{}{},
			"pool 8 ball": struct{}{},
		},
	},
	":bowling:": Code{
		"\U0001f3b3",
		map[string]struct{}{
			"ball":    struct{}{},
			"bowling": struct{}{},
			"game":    struct{}{},
		},
	},
	":cricket_game:": Code{
		"\U0001f3cf",
		map[string]struct{}{
			"ball":         struct{}{},
			"bat":          struct{}{},
			"cricket game": struct{}{},
			"game":         struct{}{},
		},
	},
	":field_hockey:": Code{
		"\U0001f3d1",
		map[string]struct{}{
			"ball":         struct{}{},
			"field":        struct{}{},
			"field hockey": struct{}{},
			"game":         struct{}{},
			"hockey":       struct{}{},
			"stick":        struct{}{},
		},
	},
	":hockey:": Code{
		"\U0001f3d2",
		map[string]struct{}{
			"game":       struct{}{},
			"hockey":     struct{}{},
			"ice":        struct{}{},
			"ice hockey": struct{}{},
			"puck":       struct{}{},
			"stick":      struct{}{},
		},
	},
	":ping_pong:": Code{
		"\U0001f3d3",
		map[string]struct{}{
			"ball":         struct{}{},
			"bat":          struct{}{},
			"game":         struct{}{},
			"paddle":       struct{}{},
			"ping pong":    struct{}{},
			"table tennis": struct{}{},
		},
	},
	":badminton:": Code{
		"\U0001f3f8",
		map[string]struct{}{
			"badminton":   struct{}{},
			"birdie":      struct{}{},
			"game":        struct{}{},
			"racquet":     struct{}{},
			"shuttlecock": struct{}{},
		},
	},
	":boxing_glove:": Code{
		"\U0001f94a",
		map[string]struct{}{
			"boxing":       struct{}{},
			"boxing glove": struct{}{},
			"glove":        struct{}{},
		},
	},
	":martial_arts_uniform:": Code{
		"\U0001f94b",
		map[string]struct{}{
			"judo":                 struct{}{},
			"karate":               struct{}{},
			"martial arts":         struct{}{},
			"martial arts uniform": struct{}{},
			"taekwondo":            struct{}{},
			"uniform":              struct{}{},
		},
	},
	":goal:": Code{
		"\U0001f945",
		map[string]struct{}{
			"goal":     struct{}{},
			"goal net": struct{}{},
			"net":      struct{}{},
		},
	},
	":dart:": Code{
		"\U0001f3af",
		map[string]struct{}{
			"bull":       struct{}{},
			"bullseye":   struct{}{},
			"dart":       struct{}{},
			"direct hit": struct{}{},
			"eye":        struct{}{},
			"game":       struct{}{},
			"hit":        struct{}{},
			"target":     struct{}{},
		},
	},
	":golf:": Code{
		"\u26f3",
		map[string]struct{}{
			"flag in hole": struct{}{},
			"golf":         struct{}{},
			"hole":         struct{}{},
		},
	},
	":ice_skate:": Code{
		"\u26f8",
		map[string]struct{}{
			"ice":       struct{}{},
			"ice skate": struct{}{},
			"skate":     struct{}{},
		},
	},
	":fishing_pole_and_fish:": Code{
		"\U0001f3a3",
		map[string]struct{}{
			"fish":         struct{}{},
			"fishing pole": struct{}{},
			"pole":         struct{}{},
		},
	},
	":running_shirt_with_sash:": Code{
		"\U0001f3bd",
		map[string]struct{}{
			"athletics":     struct{}{},
			"running":       struct{}{},
			"running shirt": struct{}{},
			"sash":          struct{}{},
			"shirt":         struct{}{},
		},
	},
	":ski:": Code{
		"\U0001f3bf",
		map[string]struct{}{
			"ski":  struct{}{},
			"skis": struct{}{},
			"snow": struct{}{},
		},
	},
	":sled:": Code{
		"\U0001f6f7",
		map[string]struct{}{
			"sled":     struct{}{},
			"sledge":   struct{}{},
			"sleigh":   struct{}{},
			"luge":     struct{}{},
			"toboggan": struct{}{},
		},
	},
	":curling_stone:": Code{
		"\U0001f94c",
		map[string]struct{}{
			"curling stone": struct{}{},
			"game":          struct{}{},
			"rock":          struct{}{},
		},
	},
	":video_game:": Code{
		"\U0001f3ae",
		map[string]struct{}{
			"controller": struct{}{},
			"game":       struct{}{},
			"video game": struct{}{},
		},
	},
	":joystick:": Code{
		"\U0001f579",
		map[string]struct{}{
			"game":       struct{}{},
			"joystick":   struct{}{},
			"video game": struct{}{},
		},
	},
	":game_die:": Code{
		"\U0001f3b2",
		map[string]struct{}{
			"dice":     struct{}{},
			"die":      struct{}{},
			"game":     struct{}{},
			"game die": struct{}{},
		},
	},
	":spades:": Code{
		"\u2660",
		map[string]struct{}{
			"card":       struct{}{},
			"game":       struct{}{},
			"spade suit": struct{}{},
		},
	},
	":hearts:": Code{
		"\u2665",
		map[string]struct{}{
			"card":       struct{}{},
			"game":       struct{}{},
			"heart suit": struct{}{},
		},
	},
	":diamonds:": Code{
		"\u2666",
		map[string]struct{}{
			"card":         struct{}{},
			"diamond suit": struct{}{},
			"game":         struct{}{},
		},
	},
	":clubs:": Code{
		"\u2663",
		map[string]struct{}{
			"card":      struct{}{},
			"club suit": struct{}{},
			"game":      struct{}{},
		},
	},
	":black_joker:": Code{
		"\U0001f0cf",
		map[string]struct{}{
			"card":     struct{}{},
			"game":     struct{}{},
			"joker":    struct{}{},
			"wildcard": struct{}{},
		},
	},
	":mahjong:": Code{
		"\U0001f004",
		map[string]struct{}{
			"game":               struct{}{},
			"mahjong":            struct{}{},
			"mahjong red dragon": struct{}{},
			"red":                struct{}{},
		},
	},
	":flower_playing_cards:": Code{
		"\U0001f3b4",
		map[string]struct{}{
			"card":                 struct{}{},
			"flower":               struct{}{},
			"flower playing cards": struct{}{},
			"game":                 struct{}{},
			"Japanese":             struct{}{},
			"playing":              struct{}{},
		},
	},
	":mute:": Code{
		"\U0001f507",
		map[string]struct{}{
			"mute":          struct{}{},
			"muted speaker": struct{}{},
			"quiet":         struct{}{},
			"silent":        struct{}{},
			"speaker":       struct{}{},
		},
	},
	":speaker:": Code{
		"\U0001f508",
		map[string]struct{}{
			"soft":               struct{}{},
			"speaker low volume": struct{}{},
		},
	},
	":sound:": Code{
		"\U0001f509",
		map[string]struct{}{
			"medium":                struct{}{},
			"speaker medium volume": struct{}{},
		},
	},
	":loud_sound:": Code{
		"\U0001f50a",
		map[string]struct{}{
			"loud":                struct{}{},
			"speaker high volume": struct{}{},
		},
	},
	":loudspeaker:": Code{
		"\U0001f4e2",
		map[string]struct{}{
			"loud":           struct{}{},
			"loudspeaker":    struct{}{},
			"public address": struct{}{},
		},
	},
	":mega:": Code{
		"\U0001f4e3",
		map[string]struct{}{
			"cheering":  struct{}{},
			"megaphone": struct{}{},
		},
	},
	":postal_horn:": Code{
		"\U0001f4ef",
		map[string]struct{}{
			"horn":        struct{}{},
			"post":        struct{}{},
			"postal":      struct{}{},
			"postal horn": struct{}{},
		},
	},
	":bell:": Code{
		"\U0001f514",
		map[string]struct{}{
			"bell": struct{}{},
		},
	},
	":no_bell:": Code{
		"\U0001f515",
		map[string]struct{}{
			"bell":            struct{}{},
			"bell with slash": struct{}{},
			"forbidden":       struct{}{},
			"mute":            struct{}{},
			"no":              struct{}{},
			"not":             struct{}{},
			"prohibited":      struct{}{},
			"quiet":           struct{}{},
			"silent":          struct{}{},
		},
	},
	":musical_score:": Code{
		"\U0001f3bc",
		map[string]struct{}{
			"music":         struct{}{},
			"musical score": struct{}{},
			"score":         struct{}{},
		},
	},
	":musical_note:": Code{
		"\U0001f3b5",
		map[string]struct{}{
			"music":        struct{}{},
			"musical note": struct{}{},
			"note":         struct{}{},
		},
	},
	":notes:": Code{
		"\U0001f3b6",
		map[string]struct{}{
			"music":         struct{}{},
			"musical notes": struct{}{},
			"note":          struct{}{},
			"notes":         struct{}{},
		},
	},
	":microphone2:": Code{
		"\U0001f399",
		map[string]struct{}{
			"mic":               struct{}{},
			"microphone":        struct{}{},
			"music":             struct{}{},
			"studio":            struct{}{},
			"studio microphone": struct{}{},
		},
	},
	":level_slider:": Code{
		"\U0001f39a",
		map[string]struct{}{
			"level":        struct{}{},
			"level slider": struct{}{},
			"music":        struct{}{},
			"slider":       struct{}{},
		},
	},
	":control_knobs:": Code{
		"\U0001f39b",
		map[string]struct{}{
			"control":       struct{}{},
			"control knobs": struct{}{},
			"knobs":         struct{}{},
			"music":         struct{}{},
		},
	},
	":microphone:": Code{
		"\U0001f3a4",
		map[string]struct{}{
			"karaoke":    struct{}{},
			"mic":        struct{}{},
			"microphone": struct{}{},
		},
	},
	":headphones:": Code{
		"\U0001f3a7",
		map[string]struct{}{
			"earbud":    struct{}{},
			"headphone": struct{}{},
		},
	},
	":radio:": Code{
		"\U0001f4fb",
		map[string]struct{}{
			"radio": struct{}{},
			"video": struct{}{},
		},
	},
	":saxophone:": Code{
		"\U0001f3b7",
		map[string]struct{}{
			"instrument": struct{}{},
			"music":      struct{}{},
			"sax":        struct{}{},
			"saxophone":  struct{}{},
		},
	},
	":guitar:": Code{
		"\U0001f3b8",
		map[string]struct{}{
			"guitar":     struct{}{},
			"instrument": struct{}{},
			"music":      struct{}{},
		},
	},
	":musical_keyboard:": Code{
		"\U0001f3b9",
		map[string]struct{}{
			"instrument":       struct{}{},
			"keyboard":         struct{}{},
			"music":            struct{}{},
			"musical keyboard": struct{}{},
			"piano":            struct{}{},
		},
	},
	":trumpet:": Code{
		"\U0001f3ba",
		map[string]struct{}{
			"instrument": struct{}{},
			"music":      struct{}{},
			"trumpet":    struct{}{},
		},
	},
	":violin:": Code{
		"\U0001f3bb",
		map[string]struct{}{
			"instrument": struct{}{},
			"music":      struct{}{},
			"violin":     struct{}{},
		},
	},
	":drum:": Code{
		"\U0001f941",
		map[string]struct{}{
			"drum":       struct{}{},
			"drumsticks": struct{}{},
			"music":      struct{}{},
		},
	},
	":iphone:": Code{
		"\U0001f4f1",
		map[string]struct{}{
			"cell":         struct{}{},
			"mobile":       struct{}{},
			"mobile phone": struct{}{},
			"phone":        struct{}{},
			"telephone":    struct{}{},
		},
	},
	":calling:": Code{
		"\U0001f4f2",
		map[string]struct{}{
			"arrow":                   struct{}{},
			"call":                    struct{}{},
			"cell":                    struct{}{},
			"mobile":                  struct{}{},
			"mobile phone with arrow": struct{}{},
			"phone":                   struct{}{},
			"receive":                 struct{}{},
			"telephone":               struct{}{},
		},
	},
	":telephone:": Code{
		"\u260e",
		map[string]struct{}{
			"phone":     struct{}{},
			"telephone": struct{}{},
		},
	},
	":telephone_receiver:": Code{
		"\U0001f4de",
		map[string]struct{}{
			"phone":              struct{}{},
			"receiver":           struct{}{},
			"telephone":          struct{}{},
			"telephone receiver": struct{}{},
		},
	},
	":pager:": Code{
		"\U0001f4df",
		map[string]struct{}{
			"pager": struct{}{},
		},
	},
	":fax:": Code{
		"\U0001f4e0",
		map[string]struct{}{
			"fax":         struct{}{},
			"fax machine": struct{}{},
		},
	},
	":battery:": Code{
		"\U0001f50b",
		map[string]struct{}{
			"battery": struct{}{},
		},
	},
	":electric_plug:": Code{
		"\U0001f50c",
		map[string]struct{}{
			"electric":      struct{}{},
			"electric plug": struct{}{},
			"electricity":   struct{}{},
			"plug":          struct{}{},
		},
	},
	":computer:": Code{
		"\U0001f4bb",
		map[string]struct{}{
			"computer":        struct{}{},
			"laptop computer": struct{}{},
			"pc":              struct{}{},
			"personal":        struct{}{},
		},
	},
	":desktop:": Code{
		"\U0001f5a5",
		map[string]struct{}{
			"computer":         struct{}{},
			"desktop":          struct{}{},
			"desktop computer": struct{}{},
		},
	},
	":printer:": Code{
		"\U0001f5a8",
		map[string]struct{}{
			"computer": struct{}{},
			"printer":  struct{}{},
		},
	},
	":keyboard:": Code{
		"\u2328",
		map[string]struct{}{
			"computer": struct{}{},
			"keyboard": struct{}{},
		},
	},
	":mouse_three_button:": Code{
		"\U0001f5b1",
		map[string]struct{}{
			"computer":       struct{}{},
			"computer mouse": struct{}{},
		},
	},
	":trackball:": Code{
		"\U0001f5b2",
		map[string]struct{}{
			"computer":  struct{}{},
			"trackball": struct{}{},
		},
	},
	":minidisc:": Code{
		"\U0001f4bd",
		map[string]struct{}{
			"computer":      struct{}{},
			"computer disk": struct{}{},
			"disk":          struct{}{},
			"minidisk":      struct{}{},
			"optical":       struct{}{},
		},
	},
	":floppy_disk:": Code{
		"\U0001f4be",
		map[string]struct{}{
			"computer":    struct{}{},
			"disk":        struct{}{},
			"floppy":      struct{}{},
			"floppy disk": struct{}{},
		},
	},
	":cd:": Code{
		"\U0001f4bf",
		map[string]struct{}{
			"cd":           struct{}{},
			"computer":     struct{}{},
			"disk":         struct{}{},
			"optical":      struct{}{},
			"optical disk": struct{}{},
		},
	},
	":dvd:": Code{
		"\U0001f4c0",
		map[string]struct{}{
			"blu-ray":  struct{}{},
			"computer": struct{}{},
			"disk":     struct{}{},
			"dvd":      struct{}{},
			"optical":  struct{}{},
		},
	},
	":movie_camera:": Code{
		"\U0001f3a5",
		map[string]struct{}{
			"camera":       struct{}{},
			"cinema":       struct{}{},
			"movie":        struct{}{},
			"movie camera": struct{}{},
		},
	},
	":film_frames:": Code{
		"\U0001f39e",
		map[string]struct{}{
			"cinema":      struct{}{},
			"film":        struct{}{},
			"film frames": struct{}{},
			"frames":      struct{}{},
			"movie":       struct{}{},
		},
	},
	":projector:": Code{
		"\U0001f4fd",
		map[string]struct{}{
			"cinema":         struct{}{},
			"film":           struct{}{},
			"film projector": struct{}{},
			"movie":          struct{}{},
			"projector":      struct{}{},
			"video":          struct{}{},
		},
	},
	":clapper:": Code{
		"\U0001f3ac",
		map[string]struct{}{
			"clapper":       struct{}{},
			"clapper board": struct{}{},
			"movie":         struct{}{},
		},
	},
	":tv:": Code{
		"\U0001f4fa",
		map[string]struct{}{
			"television": struct{}{},
			"tv":         struct{}{},
			"video":      struct{}{},
		},
	},
	":camera:": Code{
		"\U0001f4f7",
		map[string]struct{}{
			"camera": struct{}{},
			"video":  struct{}{},
		},
	},
	":camera_with_flash:": Code{
		"\U0001f4f8",
		map[string]struct{}{
			"camera":            struct{}{},
			"camera with flash": struct{}{},
			"flash":             struct{}{},
			"video":             struct{}{},
		},
	},
	":video_camera:": Code{
		"\U0001f4f9",
		map[string]struct{}{
			"camera":       struct{}{},
			"video":        struct{}{},
			"video camera": struct{}{},
		},
	},
	":vhs:": Code{
		"\U0001f4fc",
		map[string]struct{}{
			"tape":          struct{}{},
			"vhs":           struct{}{},
			"video":         struct{}{},
			"videocassette": struct{}{},
		},
	},
	":mag:": Code{
		"\U0001f50d",
		map[string]struct{}{
			"glass":                        struct{}{},
			"magnifying":                   struct{}{},
			"magnifying glass tilted left": struct{}{},
			"search":                       struct{}{},
			"tool":                         struct{}{},
		},
	},
	":mag_right:": Code{
		"\U0001f50e",
		map[string]struct{}{
			"glass":                         struct{}{},
			"magnifying":                    struct{}{},
			"magnifying glass tilted right": struct{}{},
			"search":                        struct{}{},
			"tool":                          struct{}{},
		},
	},
	":microscope:": Code{
		"\U0001f52c",
		map[string]struct{}{
			"microscope": struct{}{},
			"science":    struct{}{},
			"tool":       struct{}{},
		},
	},
	":telescope:": Code{
		"\U0001f52d",
		map[string]struct{}{
			"science":   struct{}{},
			"telescope": struct{}{},
			"tool":      struct{}{},
		},
	},
	":satellite:": Code{
		"\U0001f4e1",
		map[string]struct{}{
			"antenna":           struct{}{},
			"dish":              struct{}{},
			"satellite":         struct{}{},
			"satellite antenna": struct{}{},
		},
	},
	":candle:": Code{
		"\U0001f56f",
		map[string]struct{}{
			"candle": struct{}{},
			"light":  struct{}{},
		},
	},
	":bulb:": Code{
		"\U0001f4a1",
		map[string]struct{}{
			"bulb":       struct{}{},
			"comic":      struct{}{},
			"electric":   struct{}{},
			"idea":       struct{}{},
			"light":      struct{}{},
			"light bulb": struct{}{},
		},
	},
	":flashlight:": Code{
		"\U0001f526",
		map[string]struct{}{
			"electric":   struct{}{},
			"flashlight": struct{}{},
			"light":      struct{}{},
			"tool":       struct{}{},
			"torch":      struct{}{},
		},
	},
	":izakaya_lantern:": Code{
		"\U0001f3ee",
		map[string]struct{}{
			"bar":               struct{}{},
			"lantern":           struct{}{},
			"light":             struct{}{},
			"red":               struct{}{},
			"red paper lantern": struct{}{},
		},
	},
	":notebook_with_decorative_cover:": Code{
		"\U0001f4d4",
		map[string]struct{}{
			"book":                           struct{}{},
			"cover":                          struct{}{},
			"decorated":                      struct{}{},
			"notebook":                       struct{}{},
			"notebook with decorative cover": struct{}{},
		},
	},
	":closed_book:": Code{
		"\U0001f4d5",
		map[string]struct{}{
			"book":        struct{}{},
			"closed":      struct{}{},
			"closed book": struct{}{},
		},
	},
	":book:": Code{
		"\U0001f4d6",
		map[string]struct{}{
			"book":      struct{}{},
			"open":      struct{}{},
			"open book": struct{}{},
		},
	},
	":green_book:": Code{
		"\U0001f4d7",
		map[string]struct{}{
			"book":       struct{}{},
			"green":      struct{}{},
			"green book": struct{}{},
		},
	},
	":blue_book:": Code{
		"\U0001f4d8",
		map[string]struct{}{
			"blue":      struct{}{},
			"blue book": struct{}{},
			"book":      struct{}{},
		},
	},
	":orange_book:": Code{
		"\U0001f4d9",
		map[string]struct{}{
			"book":        struct{}{},
			"orange":      struct{}{},
			"orange book": struct{}{},
		},
	},
	":books:": Code{
		"\U0001f4da",
		map[string]struct{}{
			"book":  struct{}{},
			"books": struct{}{},
		},
	},
	":notebook:": Code{
		"\U0001f4d3",
		map[string]struct{}{
			"notebook": struct{}{},
		},
	},
	":ledger:": Code{
		"\U0001f4d2",
		map[string]struct{}{
			"ledger":   struct{}{},
			"notebook": struct{}{},
		},
	},
	":page_with_curl:": Code{
		"\U0001f4c3",
		map[string]struct{}{
			"curl":           struct{}{},
			"document":       struct{}{},
			"page":           struct{}{},
			"page with curl": struct{}{},
		},
	},
	":scroll:": Code{
		"\U0001f4dc",
		map[string]struct{}{
			"paper":  struct{}{},
			"scroll": struct{}{},
		},
	},
	":page_facing_up:": Code{
		"\U0001f4c4",
		map[string]struct{}{
			"document":       struct{}{},
			"page":           struct{}{},
			"page facing up": struct{}{},
		},
	},
	":newspaper:": Code{
		"\U0001f4f0",
		map[string]struct{}{
			"news":      struct{}{},
			"newspaper": struct{}{},
			"paper":     struct{}{},
		},
	},
	":newspaper2:": Code{
		"\U0001f5de",
		map[string]struct{}{
			"news":                struct{}{},
			"newspaper":           struct{}{},
			"paper":               struct{}{},
			"rolled":              struct{}{},
			"rolled-up newspaper": struct{}{},
		},
	},
	":bookmark_tabs:": Code{
		"\U0001f4d1",
		map[string]struct{}{
			"bookmark":      struct{}{},
			"bookmark tabs": struct{}{},
			"mark":          struct{}{},
			"marker":        struct{}{},
			"tabs":          struct{}{},
		},
	},
	":bookmark:": Code{
		"\U0001f516",
		map[string]struct{}{
			"bookmark": struct{}{},
			"mark":     struct{}{},
		},
	},
	":label:": Code{
		"\U0001f3f7",
		map[string]struct{}{
			"label": struct{}{},
		},
	},
	":moneybag:": Code{
		"\U0001f4b0",
		map[string]struct{}{
			"bag":       struct{}{},
			"dollar":    struct{}{},
			"money":     struct{}{},
			"money bag": struct{}{},
			"moneybag":  struct{}{},
		},
	},
	":yen:": Code{
		"\U0001f4b4",
		map[string]struct{}{
			"bank":         struct{}{},
			"banknote":     struct{}{},
			"bill":         struct{}{},
			"currency":     struct{}{},
			"money":        struct{}{},
			"note":         struct{}{},
			"yen":          struct{}{},
			"yen banknote": struct{}{},
		},
	},
	":dollar:": Code{
		"\U0001f4b5",
		map[string]struct{}{
			"bank":            struct{}{},
			"banknote":        struct{}{},
			"bill":            struct{}{},
			"currency":        struct{}{},
			"dollar":          struct{}{},
			"dollar banknote": struct{}{},
			"money":           struct{}{},
			"note":            struct{}{},
		},
	},
	":euro:": Code{
		"\U0001f4b6",
		map[string]struct{}{
			"bank":          struct{}{},
			"banknote":      struct{}{},
			"bill":          struct{}{},
			"currency":      struct{}{},
			"euro":          struct{}{},
			"euro banknote": struct{}{},
			"money":         struct{}{},
			"note":          struct{}{},
		},
	},
	":pound:": Code{
		"\U0001f4b7",
		map[string]struct{}{
			"bank":           struct{}{},
			"banknote":       struct{}{},
			"bill":           struct{}{},
			"currency":       struct{}{},
			"money":          struct{}{},
			"note":           struct{}{},
			"pound":          struct{}{},
			"pound banknote": struct{}{},
		},
	},
	":money_with_wings:": Code{
		"\U0001f4b8",
		map[string]struct{}{
			"bank":             struct{}{},
			"banknote":         struct{}{},
			"bill":             struct{}{},
			"dollar":           struct{}{},
			"fly":              struct{}{},
			"money":            struct{}{},
			"money with wings": struct{}{},
			"note":             struct{}{},
			"wings":            struct{}{},
		},
	},
	":credit_card:": Code{
		"\U0001f4b3",
		map[string]struct{}{
			"bank":        struct{}{},
			"card":        struct{}{},
			"credit":      struct{}{},
			"credit card": struct{}{},
			"money":       struct{}{},
		},
	},
	":chart:": Code{
		"\U0001f4b9",
		map[string]struct{}{
			"bank":                      struct{}{},
			"chart":                     struct{}{},
			"chart increasing with yen": struct{}{},
			"currency":                  struct{}{},
			"graph":                     struct{}{},
			"growth":                    struct{}{},
			"market":                    struct{}{},
			"money":                     struct{}{},
			"rise":                      struct{}{},
			"trend":                     struct{}{},
			"upward":                    struct{}{},
			"yen":                       struct{}{},
		},
	},
	":currency_exchange:": Code{
		"\U0001f4b1",
		map[string]struct{}{
			"bank":              struct{}{},
			"currency":          struct{}{},
			"currency exchange": struct{}{},
			"exchange":          struct{}{},
			"money":             struct{}{},
		},
	},
	":heavy_dollar_sign:": Code{
		"\U0001f4b2",
		map[string]struct{}{
			"currency":          struct{}{},
			"dollar":            struct{}{},
			"heavy dollar sign": struct{}{},
			"money":             struct{}{},
		},
	},
	":envelope:": Code{
		"\u2709",
		map[string]struct{}{
			"email":    struct{}{},
			"envelope": struct{}{},
			"letter":   struct{}{},
		},
	},
	":e-mail:": Code{
		"\U0001f4e7",
		map[string]struct{}{
			"e-mail": struct{}{},
			"email":  struct{}{},
			"letter": struct{}{},
			"mail":   struct{}{},
		},
	},
	":incoming_envelope:": Code{
		"\U0001f4e8",
		map[string]struct{}{
			"e-mail":            struct{}{},
			"email":             struct{}{},
			"envelope":          struct{}{},
			"incoming":          struct{}{},
			"incoming envelope": struct{}{},
			"letter":            struct{}{},
			"mail":              struct{}{},
			"receive":           struct{}{},
		},
	},
	":envelope_with_arrow:": Code{
		"\U0001f4e9",
		map[string]struct{}{
			"arrow":               struct{}{},
			"down":                struct{}{},
			"e-mail":              struct{}{},
			"email":               struct{}{},
			"envelope":            struct{}{},
			"envelope with arrow": struct{}{},
			"letter":              struct{}{},
			"mail":                struct{}{},
			"outgoing":            struct{}{},
			"sent":                struct{}{},
		},
	},
	":outbox_tray:": Code{
		"\U0001f4e4",
		map[string]struct{}{
			"box":         struct{}{},
			"letter":      struct{}{},
			"mail":        struct{}{},
			"outbox":      struct{}{},
			"outbox tray": struct{}{},
			"sent":        struct{}{},
			"tray":        struct{}{},
		},
	},
	":inbox_tray:": Code{
		"\U0001f4e5",
		map[string]struct{}{
			"box":        struct{}{},
			"inbox":      struct{}{},
			"inbox tray": struct{}{},
			"letter":     struct{}{},
			"mail":       struct{}{},
			"receive":    struct{}{},
			"tray":       struct{}{},
		},
	},
	":package:": Code{
		"\U0001f4e6",
		map[string]struct{}{
			"box":     struct{}{},
			"package": struct{}{},
			"parcel":  struct{}{},
		},
	},
	":mailbox:": Code{
		"\U0001f4eb",
		map[string]struct{}{
			"closed":                          struct{}{},
			"closed mailbox with raised flag": struct{}{},
			"mail":                            struct{}{},
			"mailbox":                         struct{}{},
			"postbox":                         struct{}{},
		},
	},
	":mailbox_closed:": Code{
		"\U0001f4ea",
		map[string]struct{}{
			"closed":                           struct{}{},
			"closed mailbox with lowered flag": struct{}{},
			"lowered":                          struct{}{},
			"mail":                             struct{}{},
			"mailbox":                          struct{}{},
			"postbox":                          struct{}{},
		},
	},
	":mailbox_with_mail:": Code{
		"\U0001f4ec",
		map[string]struct{}{
			"mail":                          struct{}{},
			"mailbox":                       struct{}{},
			"open":                          struct{}{},
			"open mailbox with raised flag": struct{}{},
			"postbox":                       struct{}{},
		},
	},
	":mailbox_with_no_mail:": Code{
		"\U0001f4ed",
		map[string]struct{}{
			"lowered":                        struct{}{},
			"mail":                           struct{}{},
			"mailbox":                        struct{}{},
			"open":                           struct{}{},
			"open mailbox with lowered flag": struct{}{},
			"postbox":                        struct{}{},
		},
	},
	":postbox:": Code{
		"\U0001f4ee",
		map[string]struct{}{
			"mail":    struct{}{},
			"mailbox": struct{}{},
			"postbox": struct{}{},
		},
	},
	":ballot_box:": Code{
		"\U0001f5f3",
		map[string]struct{}{
			"ballot":                 struct{}{},
			"ballot box with ballot": struct{}{},
			"box":                    struct{}{},
		},
	},
	":pencil2:": Code{
		"\u270f",
		map[string]struct{}{
			"pencil": struct{}{},
		},
	},
	":black_nib:": Code{
		"\u2712",
		map[string]struct{}{
			"black nib": struct{}{},
			"nib":       struct{}{},
			"pen":       struct{}{},
		},
	},
	":pen_fountain:": Code{
		"\U0001f58b",
		map[string]struct{}{
			"fountain":     struct{}{},
			"fountain pen": struct{}{},
			"pen":          struct{}{},
		},
	},
	":pen_ballpoint:": Code{
		"\U0001f58a",
		map[string]struct{}{
			"ballpoint": struct{}{},
			"pen":       struct{}{},
		},
	},
	":paintbrush:": Code{
		"\U0001f58c",
		map[string]struct{}{
			"paintbrush": struct{}{},
			"painting":   struct{}{},
		},
	},
	":crayon:": Code{
		"\U0001f58d",
		map[string]struct{}{
			"crayon": struct{}{},
		},
	},
	":pencil:": Code{
		"\U0001f4dd",
		map[string]struct{}{
			"memo":   struct{}{},
			"pencil": struct{}{},
		},
	},
	":briefcase:": Code{
		"\U0001f4bc",
		map[string]struct{}{
			"briefcase": struct{}{},
		},
	},
	":file_folder:": Code{
		"\U0001f4c1",
		map[string]struct{}{
			"file":        struct{}{},
			"file folder": struct{}{},
			"folder":      struct{}{},
		},
	},
	":open_file_folder:": Code{
		"\U0001f4c2",
		map[string]struct{}{
			"file":             struct{}{},
			"folder":           struct{}{},
			"open":             struct{}{},
			"open file folder": struct{}{},
		},
	},
	":dividers:": Code{
		"\U0001f5c2",
		map[string]struct{}{
			"card":                struct{}{},
			"card index dividers": struct{}{},
			"dividers":            struct{}{},
			"index":               struct{}{},
		},
	},
	":date:": Code{
		"\U0001f4c5",
		map[string]struct{}{
			"calendar": struct{}{},
			"date":     struct{}{},
		},
	},
	":calendar:": Code{
		"\U0001f4c6",
		map[string]struct{}{
			"calendar":          struct{}{},
			"tear-off calendar": struct{}{},
		},
	},
	":notepad_spiral:": Code{
		"\U0001f5d2",
		map[string]struct{}{
			"note":           struct{}{},
			"pad":            struct{}{},
			"spiral":         struct{}{},
			"spiral notepad": struct{}{},
		},
	},
	":calendar_spiral:": Code{
		"\U0001f5d3",
		map[string]struct{}{
			"calendar":        struct{}{},
			"pad":             struct{}{},
			"spiral":          struct{}{},
			"spiral calendar": struct{}{},
		},
	},
	":card_index:": Code{
		"\U0001f4c7",
		map[string]struct{}{
			"card":       struct{}{},
			"card index": struct{}{},
			"index":      struct{}{},
			"rolodex":    struct{}{},
		},
	},
	":chart_with_upwards_trend:": Code{
		"\U0001f4c8",
		map[string]struct{}{
			"chart":            struct{}{},
			"chart increasing": struct{}{},
			"graph":            struct{}{},
			"growth":           struct{}{},
			"trend":            struct{}{},
			"upward":           struct{}{},
		},
	},
	":chart_with_downwards_trend:": Code{
		"\U0001f4c9",
		map[string]struct{}{
			"chart":            struct{}{},
			"chart decreasing": struct{}{},
			"down":             struct{}{},
			"graph":            struct{}{},
			"trend":            struct{}{},
		},
	},
	":bar_chart:": Code{
		"\U0001f4ca",
		map[string]struct{}{
			"bar":       struct{}{},
			"bar chart": struct{}{},
			"chart":     struct{}{},
			"graph":     struct{}{},
		},
	},
	":clipboard:": Code{
		"\U0001f4cb",
		map[string]struct{}{
			"clipboard": struct{}{},
		},
	},
	":pushpin:": Code{
		"\U0001f4cc",
		map[string]struct{}{
			"pin":     struct{}{},
			"pushpin": struct{}{},
		},
	},
	":round_pushpin:": Code{
		"\U0001f4cd",
		map[string]struct{}{
			"pin":           struct{}{},
			"pushpin":       struct{}{},
			"round pushpin": struct{}{},
		},
	},
	":paperclip:": Code{
		"\U0001f4ce",
		map[string]struct{}{
			"paperclip": struct{}{},
		},
	},
	":paperclips:": Code{
		"\U0001f587",
		map[string]struct{}{
			"link":              struct{}{},
			"linked paperclips": struct{}{},
			"paperclip":         struct{}{},
		},
	},
	":straight_ruler:": Code{
		"\U0001f4cf",
		map[string]struct{}{
			"ruler":          struct{}{},
			"straight edge":  struct{}{},
			"straight ruler": struct{}{},
		},
	},
	":triangular_ruler:": Code{
		"\U0001f4d0",
		map[string]struct{}{
			"ruler":            struct{}{},
			"set":              struct{}{},
			"triangle":         struct{}{},
			"triangular ruler": struct{}{},
		},
	},
	":scissors:": Code{
		"\u2702",
		map[string]struct{}{
			"cutting":  struct{}{},
			"scissors": struct{}{},
			"tool":     struct{}{},
		},
	},
	":card_box:": Code{
		"\U0001f5c3",
		map[string]struct{}{
			"box":           struct{}{},
			"card":          struct{}{},
			"card file box": struct{}{},
			"file":          struct{}{},
		},
	},
	":file_cabinet:": Code{
		"\U0001f5c4",
		map[string]struct{}{
			"cabinet":      struct{}{},
			"file":         struct{}{},
			"file cabinet": struct{}{},
			"filing":       struct{}{},
		},
	},
	":wastebasket:": Code{
		"\U0001f5d1",
		map[string]struct{}{
			"wastebasket": struct{}{},
		},
	},
	":lock:": Code{
		"\U0001f512",
		map[string]struct{}{
			"closed": struct{}{},
			"locked": struct{}{},
		},
	},
	":unlock:": Code{
		"\U0001f513",
		map[string]struct{}{
			"lock":     struct{}{},
			"open":     struct{}{},
			"unlock":   struct{}{},
			"unlocked": struct{}{},
		},
	},
	":lock_with_ink_pen:": Code{
		"\U0001f50f",
		map[string]struct{}{
			"ink":             struct{}{},
			"lock":            struct{}{},
			"locked with pen": struct{}{},
			"nib":             struct{}{},
			"pen":             struct{}{},
			"privacy":         struct{}{},
		},
	},
	":closed_lock_with_key:": Code{
		"\U0001f510",
		map[string]struct{}{
			"closed":          struct{}{},
			"key":             struct{}{},
			"lock":            struct{}{},
			"locked with key": struct{}{},
			"secure":          struct{}{},
		},
	},
	":key:": Code{
		"\U0001f511",
		map[string]struct{}{
			"key":      struct{}{},
			"lock":     struct{}{},
			"password": struct{}{},
		},
	},
	":key2:": Code{
		"\U0001f5dd",
		map[string]struct{}{
			"clue":    struct{}{},
			"key":     struct{}{},
			"lock":    struct{}{},
			"old":     struct{}{},
			"old key": struct{}{},
		},
	},
	":hammer:": Code{
		"\U0001f528",
		map[string]struct{}{
			"hammer": struct{}{},
			"tool":   struct{}{},
		},
	},
	":pick:": Code{
		"\u26cf",
		map[string]struct{}{
			"mining": struct{}{},
			"pick":   struct{}{},
			"tool":   struct{}{},
		},
	},
	":hammer_pick:": Code{
		"\u2692",
		map[string]struct{}{
			"hammer":          struct{}{},
			"hammer and pick": struct{}{},
			"pick":            struct{}{},
			"tool":            struct{}{},
		},
	},
	":tools:": Code{
		"\U0001f6e0",
		map[string]struct{}{
			"hammer":            struct{}{},
			"hammer and wrench": struct{}{},
			"spanner":           struct{}{},
			"tool":              struct{}{},
			"wrench":            struct{}{},
		},
	},
	":dagger:": Code{
		"\U0001f5e1",
		map[string]struct{}{
			"dagger": struct{}{},
			"knife":  struct{}{},
			"weapon": struct{}{},
		},
	},
	":crossed_swords:": Code{
		"\u2694",
		map[string]struct{}{
			"crossed":        struct{}{},
			"crossed swords": struct{}{},
			"swords":         struct{}{},
			"weapon":         struct{}{},
		},
	},
	":gun:": Code{
		"\U0001f52b",
		map[string]struct{}{
			"gun":      struct{}{},
			"handgun":  struct{}{},
			"pistol":   struct{}{},
			"revolver": struct{}{},
			"tool":     struct{}{},
			"weapon":   struct{}{},
		},
	},
	":bow_and_arrow:": Code{
		"\U0001f3f9",
		map[string]struct{}{
			"archer":        struct{}{},
			"archery":       struct{}{},
			"arrow":         struct{}{},
			"bow":           struct{}{},
			"bow and arrow": struct{}{},
			"Sagittarius":   struct{}{},
			"tool":          struct{}{},
			"weapon":        struct{}{},
			"zodiac":        struct{}{},
		},
	},
	":shield:": Code{
		"\U0001f6e1",
		map[string]struct{}{
			"shield": struct{}{},
			"weapon": struct{}{},
		},
	},
	":wrench:": Code{
		"\U0001f527",
		map[string]struct{}{
			"spanner": struct{}{},
			"tool":    struct{}{},
			"wrench":  struct{}{},
		},
	},
	":nut_and_bolt:": Code{
		"\U0001f529",
		map[string]struct{}{
			"bolt":         struct{}{},
			"nut":          struct{}{},
			"nut and bolt": struct{}{},
			"tool":         struct{}{},
		},
	},
	":gear:": Code{
		"\u2699",
		map[string]struct{}{
			"gear": struct{}{},
			"tool": struct{}{},
		},
	},
	":compression:": Code{
		"\U0001f5dc",
		map[string]struct{}{
			"clamp":    struct{}{},
			"compress": struct{}{},
			"tool":     struct{}{},
			"vice":     struct{}{},
		},
	},
	":alembic:": Code{
		"\u2697",
		map[string]struct{}{
			"alembic":   struct{}{},
			"chemistry": struct{}{},
			"tool":      struct{}{},
		},
	},
	":scales:": Code{
		"\u2696",
		map[string]struct{}{
			"balance":       struct{}{},
			"balance scale": struct{}{},
			"justice":       struct{}{},
			"Libra":         struct{}{},
			"scales":        struct{}{},
			"tool":          struct{}{},
			"weight":        struct{}{},
			"zodiac":        struct{}{},
		},
	},
	":link:": Code{
		"\U0001f517",
		map[string]struct{}{
			"link": struct{}{},
		},
	},
	":chains:": Code{
		"\u26d3",
		map[string]struct{}{
			"chain":  struct{}{},
			"chains": struct{}{},
		},
	},
	":syringe:": Code{
		"\U0001f489",
		map[string]struct{}{
			"doctor":   struct{}{},
			"medicine": struct{}{},
			"needle":   struct{}{},
			"shot":     struct{}{},
			"sick":     struct{}{},
			"syringe":  struct{}{},
			"tool":     struct{}{},
		},
	},
	":pill:": Code{
		"\U0001f48a",
		map[string]struct{}{
			"doctor":   struct{}{},
			"medicine": struct{}{},
			"pill":     struct{}{},
			"sick":     struct{}{},
		},
	},
	":smoking:": Code{
		"\U0001f6ac",
		map[string]struct{}{
			"cigarette": struct{}{},
			"smoking":   struct{}{},
		},
	},
	":coffin:": Code{
		"\u26b0",
		map[string]struct{}{
			"coffin": struct{}{},
			"death":  struct{}{},
		},
	},
	":urn:": Code{
		"\u26b1",
		map[string]struct{}{
			"ashes":       struct{}{},
			"death":       struct{}{},
			"funeral":     struct{}{},
			"funeral urn": struct{}{},
			"urn":         struct{}{},
		},
	},
	":moyai:": Code{
		"\U0001f5ff",
		map[string]struct{}{
			"face":   struct{}{},
			"moai":   struct{}{},
			"moyai":  struct{}{},
			"statue": struct{}{},
		},
	},
	":oil:": Code{
		"\U0001f6e2",
		map[string]struct{}{
			"drum":     struct{}{},
			"oil":      struct{}{},
			"oil drum": struct{}{},
		},
	},
	":crystal_ball:": Code{
		"\U0001f52e",
		map[string]struct{}{
			"ball":         struct{}{},
			"crystal":      struct{}{},
			"crystal ball": struct{}{},
			"fairy tale":   struct{}{},
			"fantasy":      struct{}{},
			"fortune":      struct{}{},
			"tool":         struct{}{},
		},
	},
	":shopping_cart:": Code{
		"\U0001f6d2",
		map[string]struct{}{
			"cart":          struct{}{},
			"shopping":      struct{}{},
			"shopping cart": struct{}{},
			"trolley":       struct{}{},
		},
	},
	":atm:": Code{
		"\U0001f3e7",
		map[string]struct{}{
			"atm":       struct{}{},
			"ATM sign":  struct{}{},
			"automated": struct{}{},
			"bank":      struct{}{},
			"teller":    struct{}{},
		},
	},
	":put_litter_in_its_place:": Code{
		"\U0001f6ae",
		map[string]struct{}{
			"litter":             struct{}{},
			"litter bin":         struct{}{},
			"litter in bin sign": struct{}{},
		},
	},
	":potable_water:": Code{
		"\U0001f6b0",
		map[string]struct{}{
			"drinking":      struct{}{},
			"potable":       struct{}{},
			"potable water": struct{}{},
			"water":         struct{}{},
		},
	},
	":wheelchair:": Code{
		"\u267f",
		map[string]struct{}{
			"access":            struct{}{},
			"wheelchair symbol": struct{}{},
		},
	},
	":mens:": Code{
		"\U0001f6b9",
		map[string]struct{}{
			"lavatory":   struct{}{},
			"man":        struct{}{},
			"men’s room": struct{}{},
			"restroom":   struct{}{},
			"wc":         struct{}{},
		},
	},
	":womens:": Code{
		"\U0001f6ba",
		map[string]struct{}{
			"lavatory":     struct{}{},
			"restroom":     struct{}{},
			"wc":           struct{}{},
			"woman":        struct{}{},
			"women’s room": struct{}{},
		},
	},
	":restroom:": Code{
		"\U0001f6bb",
		map[string]struct{}{
			"lavatory": struct{}{},
			"restroom": struct{}{},
			"WC":       struct{}{},
		},
	},
	":baby_symbol:": Code{
		"\U0001f6bc",
		map[string]struct{}{
			"baby":        struct{}{},
			"baby symbol": struct{}{},
			"changing":    struct{}{},
		},
	},
	":wc:": Code{
		"\U0001f6be",
		map[string]struct{}{
			"closet":       struct{}{},
			"lavatory":     struct{}{},
			"restroom":     struct{}{},
			"water":        struct{}{},
			"water closet": struct{}{},
			"wc":           struct{}{},
		},
	},
	":passport_control:": Code{
		"\U0001f6c2",
		map[string]struct{}{
			"control":          struct{}{},
			"passport":         struct{}{},
			"passport control": struct{}{},
		},
	},
	":customs:": Code{
		"\U0001f6c3",
		map[string]struct{}{
			"customs": struct{}{},
		},
	},
	":baggage_claim:": Code{
		"\U0001f6c4",
		map[string]struct{}{
			"baggage":       struct{}{},
			"baggage claim": struct{}{},
			"claim":         struct{}{},
		},
	},
	":left_luggage:": Code{
		"\U0001f6c5",
		map[string]struct{}{
			"baggage":      struct{}{},
			"left luggage": struct{}{},
			"locker":       struct{}{},
			"luggage":      struct{}{},
		},
	},
	":warning:": Code{
		"\u26a0",
		map[string]struct{}{
			"warning": struct{}{},
		},
	},
	":children_crossing:": Code{
		"\U0001f6b8",
		map[string]struct{}{
			"child":             struct{}{},
			"children crossing": struct{}{},
			"crossing":          struct{}{},
			"pedestrian":        struct{}{},
			"traffic":           struct{}{},
		},
	},
	":no_entry:": Code{
		"\u26d4",
		map[string]struct{}{
			"entry":      struct{}{},
			"forbidden":  struct{}{},
			"no":         struct{}{},
			"no entry":   struct{}{},
			"not":        struct{}{},
			"prohibited": struct{}{},
			"traffic":    struct{}{},
		},
	},
	":no_entry_sign:": Code{
		"\U0001f6ab",
		map[string]struct{}{
			"entry":      struct{}{},
			"forbidden":  struct{}{},
			"no":         struct{}{},
			"not":        struct{}{},
			"prohibited": struct{}{},
		},
	},
	":no_bicycles:": Code{
		"\U0001f6b3",
		map[string]struct{}{
			"bicycle":     struct{}{},
			"bike":        struct{}{},
			"forbidden":   struct{}{},
			"no":          struct{}{},
			"no bicycles": struct{}{},
			"not":         struct{}{},
			"prohibited":  struct{}{},
		},
	},
	":no_smoking:": Code{
		"\U0001f6ad",
		map[string]struct{}{
			"forbidden":  struct{}{},
			"no":         struct{}{},
			"no smoking": struct{}{},
			"not":        struct{}{},
			"prohibited": struct{}{},
			"smoking":    struct{}{},
		},
	},
	":do_not_litter:": Code{
		"\U0001f6af",
		map[string]struct{}{
			"forbidden":    struct{}{},
			"litter":       struct{}{},
			"no":           struct{}{},
			"no littering": struct{}{},
			"not":          struct{}{},
			"prohibited":   struct{}{},
		},
	},
	":non-potable_water:": Code{
		"\U0001f6b1",
		map[string]struct{}{
			"non-drinking":      struct{}{},
			"non-potable":       struct{}{},
			"non-potable water": struct{}{},
			"water":             struct{}{},
		},
	},
	":no_pedestrians:": Code{
		"\U0001f6b7",
		map[string]struct{}{
			"forbidden":      struct{}{},
			"no":             struct{}{},
			"no pedestrians": struct{}{},
			"not":            struct{}{},
			"pedestrian":     struct{}{},
			"prohibited":     struct{}{},
		},
	},
	":no_mobile_phones:": Code{
		"\U0001f4f5",
		map[string]struct{}{
			"cell":             struct{}{},
			"forbidden":        struct{}{},
			"mobile":           struct{}{},
			"no":               struct{}{},
			"no mobile phones": struct{}{},
			"not":              struct{}{},
			"phone":            struct{}{},
			"prohibited":       struct{}{},
			"telephone":        struct{}{},
		},
	},
	":underage:": Code{
		"\U0001f51e",
		map[string]struct{}{
			"18":                    struct{}{},
			"age restriction":       struct{}{},
			"eighteen":              struct{}{},
			"forbidden":             struct{}{},
			"no":                    struct{}{},
			"no one under eighteen": struct{}{},
			"not":                   struct{}{},
			"prohibited":            struct{}{},
			"underage":              struct{}{},
		},
	},
	":radioactive:": Code{
		"\u2622",
		map[string]struct{}{
			"radioactive": struct{}{},
			"sign":        struct{}{},
		},
	},
	":biohazard:": Code{
		"\u2623",
		map[string]struct{}{
			"biohazard": struct{}{},
			"sign":      struct{}{},
		},
	},
	":arrow_up:": Code{
		"\u2b06",
		map[string]struct{}{
			"arrow":     struct{}{},
			"cardinal":  struct{}{},
			"direction": struct{}{},
			"north":     struct{}{},
			"up arrow":  struct{}{},
		},
	},
	":arrow_upper_right:": Code{
		"\u2197",
		map[string]struct{}{
			"arrow":          struct{}{},
			"direction":      struct{}{},
			"intercardinal":  struct{}{},
			"northeast":      struct{}{},
			"up-right arrow": struct{}{},
		},
	},
	":arrow_right:": Code{
		"\u27a1",
		map[string]struct{}{
			"arrow":       struct{}{},
			"cardinal":    struct{}{},
			"direction":   struct{}{},
			"east":        struct{}{},
			"right arrow": struct{}{},
		},
	},
	":arrow_lower_right:": Code{
		"\u2198",
		map[string]struct{}{
			"arrow":            struct{}{},
			"direction":        struct{}{},
			"down-right arrow": struct{}{},
			"intercardinal":    struct{}{},
			"southeast":        struct{}{},
		},
	},
	":arrow_down:": Code{
		"\u2b07",
		map[string]struct{}{
			"arrow":      struct{}{},
			"cardinal":   struct{}{},
			"direction":  struct{}{},
			"down":       struct{}{},
			"down arrow": struct{}{},
			"south":      struct{}{},
		},
	},
	":arrow_lower_left:": Code{
		"\u2199",
		map[string]struct{}{
			"arrow":           struct{}{},
			"direction":       struct{}{},
			"down-left arrow": struct{}{},
			"intercardinal":   struct{}{},
			"southwest":       struct{}{},
		},
	},
	":arrow_left:": Code{
		"\u2b05",
		map[string]struct{}{
			"arrow":      struct{}{},
			"cardinal":   struct{}{},
			"direction":  struct{}{},
			"left arrow": struct{}{},
			"west":       struct{}{},
		},
	},
	":arrow_upper_left:": Code{
		"\u2196",
		map[string]struct{}{
			"arrow":         struct{}{},
			"direction":     struct{}{},
			"intercardinal": struct{}{},
			"northwest":     struct{}{},
			"up-left arrow": struct{}{},
		},
	},
	":arrow_up_down:": Code{
		"\u2195",
		map[string]struct{}{
			"arrow":         struct{}{},
			"up-down arrow": struct{}{},
		},
	},
	":left_right_arrow:": Code{
		"\u2194",
		map[string]struct{}{
			"arrow":            struct{}{},
			"left-right arrow": struct{}{},
		},
	},
	":leftwards_arrow_with_hook:": Code{
		"\u21a9",
		map[string]struct{}{
			"arrow":                    struct{}{},
			"right arrow curving left": struct{}{},
		},
	},
	":arrow_right_hook:": Code{
		"\u21aa",
		map[string]struct{}{
			"arrow":                    struct{}{},
			"left arrow curving right": struct{}{},
		},
	},
	":arrow_heading_up:": Code{
		"\u2934",
		map[string]struct{}{
			"arrow":                  struct{}{},
			"right arrow curving up": struct{}{},
		},
	},
	":arrow_heading_down:": Code{
		"\u2935",
		map[string]struct{}{
			"arrow":                    struct{}{},
			"down":                     struct{}{},
			"right arrow curving down": struct{}{},
		},
	},
	":arrows_clockwise:": Code{
		"\U0001f503",
		map[string]struct{}{
			"arrow":                     struct{}{},
			"clockwise":                 struct{}{},
			"clockwise vertical arrows": struct{}{},
			"reload":                    struct{}{},
		},
	},
	":arrows_counterclockwise:": Code{
		"\U0001f504",
		map[string]struct{}{
			"anticlockwise":                  struct{}{},
			"arrow":                          struct{}{},
			"counterclockwise":               struct{}{},
			"counterclockwise arrows button": struct{}{},
			"withershins":                    struct{}{},
		},
	},
	":back:": Code{
		"\U0001f519",
		map[string]struct{}{
			"arrow":      struct{}{},
			"back":       struct{}{},
			"BACK arrow": struct{}{},
		},
	},
	":end:": Code{
		"\U0001f51a",
		map[string]struct{}{
			"arrow":     struct{}{},
			"end":       struct{}{},
			"END arrow": struct{}{},
		},
	},
	":on:": Code{
		"\U0001f51b",
		map[string]struct{}{
			"arrow":     struct{}{},
			"mark":      struct{}{},
			"on":        struct{}{},
			"ON! arrow": struct{}{},
		},
	},
	":soon:": Code{
		"\U0001f51c",
		map[string]struct{}{
			"arrow":      struct{}{},
			"soon":       struct{}{},
			"SOON arrow": struct{}{},
		},
	},
	":top:": Code{
		"\U0001f51d",
		map[string]struct{}{
			"arrow":     struct{}{},
			"top":       struct{}{},
			"TOP arrow": struct{}{},
			"up":        struct{}{},
		},
	},
	":place_of_worship:": Code{
		"\U0001f6d0",
		map[string]struct{}{
			"place of worship": struct{}{},
			"religion":         struct{}{},
			"worship":          struct{}{},
		},
	},
	":atom:": Code{
		"\u269b",
		map[string]struct{}{
			"atheist":     struct{}{},
			"atom":        struct{}{},
			"atom symbol": struct{}{},
		},
	},
	":om_symbol:": Code{
		"\U0001f549",
		map[string]struct{}{
			"Hindu":    struct{}{},
			"om":       struct{}{},
			"religion": struct{}{},
		},
	},
	":star_of_david:": Code{
		"\u2721",
		map[string]struct{}{
			"David":         struct{}{},
			"Jew":           struct{}{},
			"Jewish":        struct{}{},
			"religion":      struct{}{},
			"star":          struct{}{},
			"star of David": struct{}{},
		},
	},
	":wheel_of_dharma:": Code{
		"\u2638",
		map[string]struct{}{
			"Buddhist":        struct{}{},
			"dharma":          struct{}{},
			"religion":        struct{}{},
			"wheel":           struct{}{},
			"wheel of dharma": struct{}{},
		},
	},
	":yin_yang:": Code{
		"\u262f",
		map[string]struct{}{
			"religion": struct{}{},
			"tao":      struct{}{},
			"taoist":   struct{}{},
			"yang":     struct{}{},
			"yin":      struct{}{},
			"yin yang": struct{}{},
		},
	},
	":cross:": Code{
		"\u271d",
		map[string]struct{}{
			"Christian":   struct{}{},
			"cross":       struct{}{},
			"latin cross": struct{}{},
			"religion":    struct{}{},
		},
	},
	":orthodox_cross:": Code{
		"\u2626",
		map[string]struct{}{
			"Christian":      struct{}{},
			"cross":          struct{}{},
			"orthodox cross": struct{}{},
			"religion":       struct{}{},
		},
	},
	":star_and_crescent:": Code{
		"\u262a",
		map[string]struct{}{
			"islam":             struct{}{},
			"Muslim":            struct{}{},
			"religion":          struct{}{},
			"star and crescent": struct{}{},
		},
	},
	":peace:": Code{
		"\u262e",
		map[string]struct{}{
			"peace":        struct{}{},
			"peace symbol": struct{}{},
		},
	},
	":menorah:": Code{
		"\U0001f54e",
		map[string]struct{}{
			"candelabrum": struct{}{},
			"candlestick": struct{}{},
			"menorah":     struct{}{},
			"religion":    struct{}{},
		},
	},
	":six_pointed_star:": Code{
		"\U0001f52f",
		map[string]struct{}{
			"dotted six-pointed star": struct{}{},
			"fortune":                 struct{}{},
			"star":                    struct{}{},
		},
	},
	":aries:": Code{
		"\u2648",
		map[string]struct{}{
			"Aries":  struct{}{},
			"ram":    struct{}{},
			"zodiac": struct{}{},
		},
	},
	":taurus:": Code{
		"\u2649",
		map[string]struct{}{
			"bull":   struct{}{},
			"ox":     struct{}{},
			"Taurus": struct{}{},
			"zodiac": struct{}{},
		},
	},
	":gemini:": Code{
		"\u264a",
		map[string]struct{}{
			"Gemini": struct{}{},
			"twins":  struct{}{},
			"zodiac": struct{}{},
		},
	},
	":cancer:": Code{
		"\u264b",
		map[string]struct{}{
			"Cancer": struct{}{},
			"crab":   struct{}{},
			"zodiac": struct{}{},
		},
	},
	":leo:": Code{
		"\u264c",
		map[string]struct{}{
			"Leo":    struct{}{},
			"lion":   struct{}{},
			"zodiac": struct{}{},
		},
	},
	":virgo:": Code{
		"\u264d",
		map[string]struct{}{
			"Virgo":  struct{}{},
			"zodiac": struct{}{},
		},
	},
	":libra:": Code{
		"\u264e",
		map[string]struct{}{
			"balance": struct{}{},
			"justice": struct{}{},
			"Libra":   struct{}{},
			"scales":  struct{}{},
			"zodiac":  struct{}{},
		},
	},
	":scorpius:": Code{
		"\u264f",
		map[string]struct{}{
			"Scorpio":  struct{}{},
			"scorpion": struct{}{},
			"scorpius": struct{}{},
			"zodiac":   struct{}{},
		},
	},
	":sagittarius:": Code{
		"\u2650",
		map[string]struct{}{
			"archer":      struct{}{},
			"Sagittarius": struct{}{},
			"zodiac":      struct{}{},
		},
	},
	":capricorn:": Code{
		"\u2651",
		map[string]struct{}{
			"Capricorn": struct{}{},
			"goat":      struct{}{},
			"zodiac":    struct{}{},
		},
	},
	":aquarius:": Code{
		"\u2652",
		map[string]struct{}{
			"Aquarius": struct{}{},
			"bearer":   struct{}{},
			"water":    struct{}{},
			"zodiac":   struct{}{},
		},
	},
	":pisces:": Code{
		"\u2653",
		map[string]struct{}{
			"fish":   struct{}{},
			"Pisces": struct{}{},
			"zodiac": struct{}{},
		},
	},
	":ophiuchus:": Code{
		"\u26ce",
		map[string]struct{}{
			"bearer":    struct{}{},
			"Ophiuchus": struct{}{},
			"serpent":   struct{}{},
			"snake":     struct{}{},
			"zodiac":    struct{}{},
		},
	},
	":twisted_rightwards_arrows:": Code{
		"\U0001f500",
		map[string]struct{}{
			"arrow":                 struct{}{},
			"crossed":               struct{}{},
			"shuffle tracks button": struct{}{},
		},
	},
	":repeat:": Code{
		"\U0001f501",
		map[string]struct{}{
			"arrow":         struct{}{},
			"clockwise":     struct{}{},
			"repeat":        struct{}{},
			"repeat button": struct{}{},
		},
	},
	":repeat_one:": Code{
		"\U0001f502",
		map[string]struct{}{
			"arrow":                struct{}{},
			"clockwise":            struct{}{},
			"once":                 struct{}{},
			"repeat single button": struct{}{},
		},
	},
	":arrow_forward:": Code{
		"\u25b6",
		map[string]struct{}{
			"arrow":       struct{}{},
			"play":        struct{}{},
			"play button": struct{}{},
			"right":       struct{}{},
			"triangle":    struct{}{},
		},
	},
	":fast_forward:": Code{
		"\u23e9",
		map[string]struct{}{
			"arrow":               struct{}{},
			"double":              struct{}{},
			"fast":                struct{}{},
			"fast-forward button": struct{}{},
			"forward":             struct{}{},
		},
	},
	":track_next:": Code{
		"\u23ed",
		map[string]struct{}{
			"arrow":             struct{}{},
			"next scene":        struct{}{},
			"next track":        struct{}{},
			"next track button": struct{}{},
			"triangle":          struct{}{},
		},
	},
	":play_pause:": Code{
		"\u23ef",
		map[string]struct{}{
			"arrow":                struct{}{},
			"pause":                struct{}{},
			"play":                 struct{}{},
			"play or pause button": struct{}{},
			"right":                struct{}{},
			"triangle":             struct{}{},
		},
	},
	":arrow_backward:": Code{
		"\u25c0",
		map[string]struct{}{
			"arrow":          struct{}{},
			"left":           struct{}{},
			"reverse":        struct{}{},
			"reverse button": struct{}{},
			"triangle":       struct{}{},
		},
	},
	":rewind:": Code{
		"\u23ea",
		map[string]struct{}{
			"arrow":               struct{}{},
			"double":              struct{}{},
			"fast reverse button": struct{}{},
			"rewind":              struct{}{},
		},
	},
	":track_previous:": Code{
		"\u23ee",
		map[string]struct{}{
			"arrow":             struct{}{},
			"last track button": struct{}{},
			"previous scene":    struct{}{},
			"previous track":    struct{}{},
			"triangle":          struct{}{},
		},
	},
	":arrow_up_small:": Code{
		"\U0001f53c",
		map[string]struct{}{
			"arrow":          struct{}{},
			"button":         struct{}{},
			"red":            struct{}{},
			"upwards button": struct{}{},
		},
	},
	":arrow_double_up:": Code{
		"\u23eb",
		map[string]struct{}{
			"arrow":          struct{}{},
			"double":         struct{}{},
			"fast up button": struct{}{},
		},
	},
	":arrow_down_small:": Code{
		"\U0001f53d",
		map[string]struct{}{
			"arrow":            struct{}{},
			"button":           struct{}{},
			"down":             struct{}{},
			"downwards button": struct{}{},
			"red":              struct{}{},
		},
	},
	":arrow_double_down:": Code{
		"\u23ec",
		map[string]struct{}{
			"arrow":            struct{}{},
			"double":           struct{}{},
			"down":             struct{}{},
			"fast down button": struct{}{},
		},
	},
	":pause_button:": Code{
		"\u23f8",
		map[string]struct{}{
			"bar":          struct{}{},
			"double":       struct{}{},
			"pause":        struct{}{},
			"pause button": struct{}{},
			"vertical":     struct{}{},
		},
	},
	":stop_button:": Code{
		"\u23f9",
		map[string]struct{}{
			"square":      struct{}{},
			"stop":        struct{}{},
			"stop button": struct{}{},
		},
	},
	":record_button:": Code{
		"\u23fa",
		map[string]struct{}{
			"circle":        struct{}{},
			"record":        struct{}{},
			"record button": struct{}{},
		},
	},
	":eject:": Code{
		"\u23cf",
		map[string]struct{}{
			"eject":        struct{}{},
			"eject button": struct{}{},
		},
	},
	":cinema:": Code{
		"\U0001f3a6",
		map[string]struct{}{
			"camera": struct{}{},
			"cinema": struct{}{},
			"film":   struct{}{},
			"movie":  struct{}{},
		},
	},
	":low_brightness:": Code{
		"\U0001f505",
		map[string]struct{}{
			"brightness": struct{}{},
			"dim":        struct{}{},
			"dim button": struct{}{},
			"low":        struct{}{},
		},
	},
	":high_brightness:": Code{
		"\U0001f506",
		map[string]struct{}{
			"bright":        struct{}{},
			"bright button": struct{}{},
			"brightness":    struct{}{},
		},
	},
	":signal_strength:": Code{
		"\U0001f4f6",
		map[string]struct{}{
			"antenna":      struct{}{},
			"antenna bars": struct{}{},
			"bar":          struct{}{},
			"cell":         struct{}{},
			"mobile":       struct{}{},
			"phone":        struct{}{},
			"signal":       struct{}{},
			"telephone":    struct{}{},
		},
	},
	":vibration_mode:": Code{
		"\U0001f4f3",
		map[string]struct{}{
			"cell":           struct{}{},
			"mobile":         struct{}{},
			"mode":           struct{}{},
			"phone":          struct{}{},
			"telephone":      struct{}{},
			"vibration":      struct{}{},
			"vibration mode": struct{}{},
		},
	},
	":mobile_phone_off:": Code{
		"\U0001f4f4",
		map[string]struct{}{
			"cell":             struct{}{},
			"mobile":           struct{}{},
			"mobile phone off": struct{}{},
			"off":              struct{}{},
			"phone":            struct{}{},
			"telephone":        struct{}{},
		},
	},
	":female_sign:": Code{
		"\u2640",
		map[string]struct{}{
			"female sign": struct{}{},
			"woman":       struct{}{},
		},
	},
	":male_sign:": Code{
		"\u2642",
		map[string]struct{}{
			"male sign": struct{}{},
			"man":       struct{}{},
		},
	},
	":medical_symbol:": Code{
		"\u2695",
		map[string]struct{}{
			"aesculapius":    struct{}{},
			"medical symbol": struct{}{},
			"medicine":       struct{}{},
			"staff":          struct{}{},
		},
	},
	":recycle:": Code{
		"\u267b",
		map[string]struct{}{
			"recycle":          struct{}{},
			"recycling symbol": struct{}{},
		},
	},
	":fleur-de-lis:": Code{
		"\u269c",
		map[string]struct{}{
			"fleur-de-lis": struct{}{},
		},
	},
	":trident:": Code{
		"\U0001f531",
		map[string]struct{}{
			"anchor":         struct{}{},
			"emblem":         struct{}{},
			"ship":           struct{}{},
			"tool":           struct{}{},
			"trident":        struct{}{},
			"trident emblem": struct{}{},
		},
	},
	":name_badge:": Code{
		"\U0001f4db",
		map[string]struct{}{
			"badge":      struct{}{},
			"name":       struct{}{},
			"name badge": struct{}{},
		},
	},
	":beginner:": Code{
		"\U0001f530",
		map[string]struct{}{
			"beginner":                     struct{}{},
			"chevron":                      struct{}{},
			"green":                        struct{}{},
			"Japanese":                     struct{}{},
			"Japanese symbol for beginner": struct{}{},
			"leaf":                         struct{}{},
			"tool":                         struct{}{},
			"yellow":                       struct{}{},
		},
	},
	":o:": Code{
		"\u2b55",
		map[string]struct{}{
			"circle":             struct{}{},
			"heavy large circle": struct{}{},
			"o":                  struct{}{},
		},
	},
	":white_check_mark:": Code{
		"\u2705",
		map[string]struct{}{
			"check":                  struct{}{},
			"mark":                   struct{}{},
			"white heavy check mark": struct{}{},
		},
	},
	":ballot_box_with_check:": Code{
		"\u2611",
		map[string]struct{}{
			"ballot":                struct{}{},
			"ballot box with check": struct{}{},
			"box":                   struct{}{},
			"check":                 struct{}{},
		},
	},
	":heavy_check_mark:": Code{
		"\u2714",
		map[string]struct{}{
			"check":            struct{}{},
			"heavy check mark": struct{}{},
			"mark":             struct{}{},
		},
	},
	":heavy_multiplication_x:": Code{
		"\u2716",
		map[string]struct{}{
			"cancel":                 struct{}{},
			"heavy multiplication x": struct{}{},
			"multiplication":         struct{}{},
			"multiply":               struct{}{},
			"x":                      struct{}{},
		},
	},
	":x:": Code{
		"\u274c",
		map[string]struct{}{
			"cancel":         struct{}{},
			"cross mark":     struct{}{},
			"mark":           struct{}{},
			"multiplication": struct{}{},
			"multiply":       struct{}{},
			"x":              struct{}{},
		},
	},
	":negative_squared_cross_mark:": Code{
		"\u274e",
		map[string]struct{}{
			"cross mark button": struct{}{},
			"mark":              struct{}{},
			"square":            struct{}{},
		},
	},
	":heavy_plus_sign:": Code{
		"\u2795",
		map[string]struct{}{
			"heavy plus sign": struct{}{},
			"math":            struct{}{},
			"plus":            struct{}{},
		},
	},
	":heavy_minus_sign:": Code{
		"\u2796",
		map[string]struct{}{
			"heavy minus sign": struct{}{},
			"math":             struct{}{},
			"minus":            struct{}{},
		},
	},
	":heavy_division_sign:": Code{
		"\u2797",
		map[string]struct{}{
			"division":            struct{}{},
			"heavy division sign": struct{}{},
			"math":                struct{}{},
		},
	},
	":curly_loop:": Code{
		"\u27b0",
		map[string]struct{}{
			"curl":       struct{}{},
			"curly loop": struct{}{},
			"loop":       struct{}{},
		},
	},
	":loop:": Code{
		"\u27bf",
		map[string]struct{}{
			"curl":              struct{}{},
			"double":            struct{}{},
			"double curly loop": struct{}{},
			"loop":              struct{}{},
		},
	},
	":part_alternation_mark:": Code{
		"\u303d",
		map[string]struct{}{
			"mark":                  struct{}{},
			"part":                  struct{}{},
			"part alternation mark": struct{}{},
		},
	},
	":eight_spoked_asterisk:": Code{
		"\u2733",
		map[string]struct{}{
			"asterisk":              struct{}{},
			"eight-spoked asterisk": struct{}{},
		},
	},
	":eight_pointed_black_star:": Code{
		"\u2734",
		map[string]struct{}{
			"eight-pointed star": struct{}{},
			"star":               struct{}{},
		},
	},
	":sparkle:": Code{
		"\u2747",
		map[string]struct{}{
			"sparkle": struct{}{},
		},
	},
	":bangbang:": Code{
		"\u203c",
		map[string]struct{}{
			"bangbang":                struct{}{},
			"double exclamation mark": struct{}{},
			"exclamation":             struct{}{},
			"mark":                    struct{}{},
			"punctuation":             struct{}{},
		},
	},
	":interrobang:": Code{
		"\u2049",
		map[string]struct{}{
			"exclamation":               struct{}{},
			"exclamation question mark": struct{}{},
			"interrobang":               struct{}{},
			"mark":                      struct{}{},
			"punctuation":               struct{}{},
			"question":                  struct{}{},
		},
	},
	":question:": Code{
		"\u2753",
		map[string]struct{}{
			"mark":          struct{}{},
			"punctuation":   struct{}{},
			"question":      struct{}{},
			"question mark": struct{}{},
		},
	},
	":grey_question:": Code{
		"\u2754",
		map[string]struct{}{
			"mark":                struct{}{},
			"outlined":            struct{}{},
			"punctuation":         struct{}{},
			"question":            struct{}{},
			"white question mark": struct{}{},
		},
	},
	":grey_exclamation:": Code{
		"\u2755",
		map[string]struct{}{
			"exclamation":            struct{}{},
			"mark":                   struct{}{},
			"outlined":               struct{}{},
			"punctuation":            struct{}{},
			"white exclamation mark": struct{}{},
		},
	},
	":exclamation:": Code{
		"\u2757",
		map[string]struct{}{
			"exclamation":      struct{}{},
			"exclamation mark": struct{}{},
			"mark":             struct{}{},
			"punctuation":      struct{}{},
		},
	},
	":wavy_dash:": Code{
		"\u3030",
		map[string]struct{}{
			"dash":        struct{}{},
			"punctuation": struct{}{},
			"wavy":        struct{}{},
			"wavy dash":   struct{}{},
		},
	},
	":copyright:": Code{
		"\u00a9",
		map[string]struct{}{
			"copyright": struct{}{},
		},
	},
	":registered:": Code{
		"\u00ae",
		map[string]struct{}{
			"registered": struct{}{},
		},
	},
	":tm:": Code{
		"\u2122",
		map[string]struct{}{
			"mark":       struct{}{},
			"tm":         struct{}{},
			"trade mark": struct{}{},
			"trademark":  struct{}{},
		},
	},
	":hash:": Code{
		"#\ufe0f\u20e3",
		map[string]struct{}{
			"keycap": struct{}{},
		},
	},
	":asterisk:": Code{
		"*\ufe0f\u20e3",
		map[string]struct{}{
			"keycap": struct{}{},
		},
	},
	":zero:": Code{
		"0\ufe0f\u20e3",
		map[string]struct{}{
			"keycap": struct{}{},
		},
	},
	":one:": Code{
		"1\ufe0f\u20e3",
		map[string]struct{}{
			"keycap": struct{}{},
		},
	},
	":two:": Code{
		"2\ufe0f\u20e3",
		map[string]struct{}{
			"keycap": struct{}{},
		},
	},
	":three:": Code{
		"3\ufe0f\u20e3",
		map[string]struct{}{
			"keycap": struct{}{},
		},
	},
	":four:": Code{
		"4\ufe0f\u20e3",
		map[string]struct{}{
			"keycap": struct{}{},
		},
	},
	":five:": Code{
		"5\ufe0f\u20e3",
		map[string]struct{}{
			"keycap": struct{}{},
		},
	},
	":six:": Code{
		"6\ufe0f\u20e3",
		map[string]struct{}{
			"keycap": struct{}{},
		},
	},
	":seven:": Code{
		"7\ufe0f\u20e3",
		map[string]struct{}{
			"keycap": struct{}{},
		},
	},
	":eight:": Code{
		"8\ufe0f\u20e3",
		map[string]struct{}{
			"keycap": struct{}{},
		},
	},
	":nine:": Code{
		"9\ufe0f\u20e3",
		map[string]struct{}{
			"keycap": struct{}{},
		},
	},
	":keycap_ten:": Code{
		"\U0001f51f",
		map[string]struct{}{
			"keycap": struct{}{},
		},
	},
	":100:": Code{
		"\U0001f4af",
		map[string]struct{}{
			"100":            struct{}{},
			"full":           struct{}{},
			"hundred":        struct{}{},
			"hundred points": struct{}{},
			"score":          struct{}{},
		},
	},
	":capital_abcd:": Code{
		"\U0001f520",
		map[string]struct{}{
			"ABCD":                  struct{}{},
			"input":                 struct{}{},
			"input latin uppercase": struct{}{},
			"latin":                 struct{}{},
			"letters":               struct{}{},
			"uppercase":             struct{}{},
		},
	},
	":abcd:": Code{
		"\U0001f521",
		map[string]struct{}{
			"abcd":                  struct{}{},
			"input":                 struct{}{},
			"input latin lowercase": struct{}{},
			"latin":                 struct{}{},
			"letters":               struct{}{},
			"lowercase":             struct{}{},
		},
	},
	":1234:": Code{
		"\U0001f522",
		map[string]struct{}{
			"1234":          struct{}{},
			"input":         struct{}{},
			"input numbers": struct{}{},
			"numbers":       struct{}{},
		},
	},
	":symbols:": Code{
		"\U0001f523",
		map[string]struct{}{
			"〒♪&%":          struct{}{},
			"input":         struct{}{},
			"input symbols": struct{}{},
		},
	},
	":abc:": Code{
		"\U0001f524",
		map[string]struct{}{
			"abc":                 struct{}{},
			"alphabet":            struct{}{},
			"input":               struct{}{},
			"input latin letters": struct{}{},
			"latin":               struct{}{},
			"letters":             struct{}{},
		},
	},
	":a:": Code{
		"\U0001f170",
		map[string]struct{}{
			"a":                     struct{}{},
			"A button (blood type)": struct{}{},
			"blood type":            struct{}{},
		},
	},
	":ab:": Code{
		"\U0001f18e",
		map[string]struct{}{
			"ab":                     struct{}{},
			"AB button (blood type)": struct{}{},
			"blood type":             struct{}{},
		},
	},
	":b:": Code{
		"\U0001f171",
		map[string]struct{}{
			"b":                     struct{}{},
			"B button (blood type)": struct{}{},
			"blood type":            struct{}{},
		},
	},
	":cl:": Code{
		"\U0001f191",
		map[string]struct{}{
			"cl":        struct{}{},
			"CL button": struct{}{},
		},
	},
	":cool:": Code{
		"\U0001f192",
		map[string]struct{}{
			"cool":        struct{}{},
			"COOL button": struct{}{},
		},
	},
	":free:": Code{
		"\U0001f193",
		map[string]struct{}{
			"free":        struct{}{},
			"FREE button": struct{}{},
		},
	},
	":information_source:": Code{
		"\u2139",
		map[string]struct{}{
			"i":           struct{}{},
			"information": struct{}{},
		},
	},
	":id:": Code{
		"\U0001f194",
		map[string]struct{}{
			"id":        struct{}{},
			"ID button": struct{}{},
			"identity":  struct{}{},
		},
	},
	":m:": Code{
		"\u24dc",
		map[string]struct{}{
			"circle":    struct{}{},
			"circled M": struct{}{},
			"m":         struct{}{},
		},
	},
	":new:": Code{
		"\U0001f195",
		map[string]struct{}{
			"new":        struct{}{},
			"NEW button": struct{}{},
		},
	},
	":ng:": Code{
		"\U0001f196",
		map[string]struct{}{
			"ng":        struct{}{},
			"NG button": struct{}{},
		},
	},
	":o2:": Code{
		"\U0001f17e",
		map[string]struct{}{
			"blood type":            struct{}{},
			"o":                     struct{}{},
			"O button (blood type)": struct{}{},
		},
	},
	":ok:": Code{
		"\U0001f197",
		map[string]struct{}{
			"OK":        struct{}{},
			"OK button": struct{}{},
		},
	},
	":parking:": Code{
		"\U0001f17f",
		map[string]struct{}{
			"P button": struct{}{},
			"parking":  struct{}{},
		},
	},
	":sos:": Code{
		"\U0001f198",
		map[string]struct{}{
			"help":       struct{}{},
			"sos":        struct{}{},
			"SOS button": struct{}{},
		},
	},
	":up:": Code{
		"\U0001f199",
		map[string]struct{}{
			"mark":       struct{}{},
			"up":         struct{}{},
			"UP! button": struct{}{},
		},
	},
	":vs:": Code{
		"\U0001f19a",
		map[string]struct{}{
			"versus":    struct{}{},
			"vs":        struct{}{},
			"VS button": struct{}{},
		},
	},
	":koko:": Code{
		"\U0001f201",
		map[string]struct{}{
			"“here”":                 struct{}{},
			"Japanese":               struct{}{},
			"Japanese “here” button": struct{}{},
			"katakana":               struct{}{},
			"ココ":                     struct{}{},
		},
	},
	":sa:": Code{
		"\U0001f202",
		map[string]struct{}{
			"“service charge”":                 struct{}{},
			"Japanese":                         struct{}{},
			"Japanese “service charge” button": struct{}{},
			"katakana":                         struct{}{},
			"サ":                                struct{}{},
		},
	},
	":u6708:": Code{
		"\U0001f237",
		map[string]struct{}{
			"“monthly amount”":                 struct{}{},
			"ideograph":                        struct{}{},
			"Japanese":                         struct{}{},
			"Japanese “monthly amount” button": struct{}{},
			"月":                                struct{}{},
		},
	},
	":u6709:": Code{
		"\U0001f236",
		map[string]struct{}{
			"“not free of charge”": struct{}{},
			"ideograph":            struct{}{},
			"Japanese":             struct{}{},
			"Japanese “not free of charge” button": struct{}{},
			"有": struct{}{},
		},
	},
	":u6307:": Code{
		"\U0001f22f",
		map[string]struct{}{
			"“reserved”":                 struct{}{},
			"ideograph":                  struct{}{},
			"Japanese":                   struct{}{},
			"Japanese “reserved” button": struct{}{},
			"指":                          struct{}{},
		},
	},
	":ideograph_advantage:": Code{
		"\U0001f250",
		map[string]struct{}{
			"“bargain”":                 struct{}{},
			"ideograph":                 struct{}{},
			"Japanese":                  struct{}{},
			"Japanese “bargain” button": struct{}{},
			"得":                         struct{}{},
		},
	},
	":u5272:": Code{
		"\U0001f239",
		map[string]struct{}{
			"“discount”":                 struct{}{},
			"ideograph":                  struct{}{},
			"Japanese":                   struct{}{},
			"Japanese “discount” button": struct{}{},
			"割":                          struct{}{},
		},
	},
	":u7121:": Code{
		"\U0001f21a",
		map[string]struct{}{
			"“free of charge”":                 struct{}{},
			"ideograph":                        struct{}{},
			"Japanese":                         struct{}{},
			"Japanese “free of charge” button": struct{}{},
			"無":                                struct{}{},
		},
	},
	":u7981:": Code{
		"\U0001f232",
		map[string]struct{}{
			"“prohibited”":                 struct{}{},
			"ideograph":                    struct{}{},
			"Japanese":                     struct{}{},
			"Japanese “prohibited” button": struct{}{},
			"禁":                            struct{}{},
		},
	},
	":accept:": Code{
		"\U0001f251",
		map[string]struct{}{
			"“acceptable”":                 struct{}{},
			"ideograph":                    struct{}{},
			"Japanese":                     struct{}{},
			"Japanese “acceptable” button": struct{}{},
			"可":                            struct{}{},
		},
	},
	":u7533:": Code{
		"\U0001f238",
		map[string]struct{}{
			"“application”":                 struct{}{},
			"ideograph":                     struct{}{},
			"Japanese":                      struct{}{},
			"Japanese “application” button": struct{}{},
			"申":                             struct{}{},
		},
	},
	":u5408:": Code{
		"\U0001f234",
		map[string]struct{}{
			"“passing grade”":                 struct{}{},
			"ideograph":                       struct{}{},
			"Japanese":                        struct{}{},
			"Japanese “passing grade” button": struct{}{},
			"合":                               struct{}{},
		},
	},
	":u7a7a:": Code{
		"\U0001f233",
		map[string]struct{}{
			"“vacancy”":                 struct{}{},
			"ideograph":                 struct{}{},
			"Japanese":                  struct{}{},
			"Japanese “vacancy” button": struct{}{},
			"空":                         struct{}{},
		},
	},
	":congratulations:": Code{
		"\u3297",
		map[string]struct{}{
			"“congratulations”":                 struct{}{},
			"ideograph":                         struct{}{},
			"Japanese":                          struct{}{},
			"Japanese “congratulations” button": struct{}{},
			"祝":                                 struct{}{},
		},
	},
	":secret:": Code{
		"\u3299",
		map[string]struct{}{
			"“secret”":                 struct{}{},
			"ideograph":                struct{}{},
			"Japanese":                 struct{}{},
			"Japanese “secret” button": struct{}{},
			"秘":                        struct{}{},
		},
	},
	":u55b6:": Code{
		"\U0001f23a",
		map[string]struct{}{
			"“open for business”": struct{}{},
			"ideograph":           struct{}{},
			"Japanese":            struct{}{},
			"Japanese “open for business” button": struct{}{},
			"営": struct{}{},
		},
	},
	":u6e80:": Code{
		"\U0001f235",
		map[string]struct{}{
			"“no vacancy”":                 struct{}{},
			"ideograph":                    struct{}{},
			"Japanese":                     struct{}{},
			"Japanese “no vacancy” button": struct{}{},
			"満":                            struct{}{},
		},
	},
	":black_small_square:": Code{
		"\u25aa",
		map[string]struct{}{
			"black small square": struct{}{},
			"geometric":          struct{}{},
			"square":             struct{}{},
		},
	},
	":white_small_square:": Code{
		"\u25ab",
		map[string]struct{}{
			"geometric":          struct{}{},
			"square":             struct{}{},
			"white small square": struct{}{},
		},
	},
	":white_medium_square:": Code{
		"\u25fb",
		map[string]struct{}{
			"geometric":           struct{}{},
			"square":              struct{}{},
			"white medium square": struct{}{},
		},
	},
	":black_medium_square:": Code{
		"\u25fc",
		map[string]struct{}{
			"black medium square": struct{}{},
			"geometric":           struct{}{},
			"square":              struct{}{},
		},
	},
	":white_medium_small_square:": Code{
		"\u25fd",
		map[string]struct{}{
			"geometric":                 struct{}{},
			"square":                    struct{}{},
			"white medium-small square": struct{}{},
		},
	},
	":black_medium_small_square:": Code{
		"\u25fe",
		map[string]struct{}{
			"black medium-small square": struct{}{},
			"geometric":                 struct{}{},
			"square":                    struct{}{},
		},
	},
	":black_large_square:": Code{
		"\u2b1b",
		map[string]struct{}{
			"black large square": struct{}{},
			"geometric":          struct{}{},
			"square":             struct{}{},
		},
	},
	":white_large_square:": Code{
		"\u2b1c",
		map[string]struct{}{
			"geometric":          struct{}{},
			"square":             struct{}{},
			"white large square": struct{}{},
		},
	},
	":large_orange_diamond:": Code{
		"\U0001f536",
		map[string]struct{}{
			"diamond":              struct{}{},
			"geometric":            struct{}{},
			"large orange diamond": struct{}{},
			"orange":               struct{}{},
		},
	},
	":large_blue_diamond:": Code{
		"\U0001f537",
		map[string]struct{}{
			"blue":               struct{}{},
			"diamond":            struct{}{},
			"geometric":          struct{}{},
			"large blue diamond": struct{}{},
		},
	},
	":small_orange_diamond:": Code{
		"\U0001f538",
		map[string]struct{}{
			"diamond":              struct{}{},
			"geometric":            struct{}{},
			"orange":               struct{}{},
			"small orange diamond": struct{}{},
		},
	},
	":small_blue_diamond:": Code{
		"\U0001f539",
		map[string]struct{}{
			"blue":               struct{}{},
			"diamond":            struct{}{},
			"geometric":          struct{}{},
			"small blue diamond": struct{}{},
		},
	},
	":small_red_triangle:": Code{
		"\U0001f53a",
		map[string]struct{}{
			"geometric":               struct{}{},
			"red":                     struct{}{},
			"red triangle pointed up": struct{}{},
		},
	},
	":small_red_triangle_down:": Code{
		"\U0001f53b",
		map[string]struct{}{
			"down":                      struct{}{},
			"geometric":                 struct{}{},
			"red":                       struct{}{},
			"red triangle pointed down": struct{}{},
		},
	},
	":diamond_shape_with_a_dot_inside:": Code{
		"\U0001f4a0",
		map[string]struct{}{
			"comic":              struct{}{},
			"diamond":            struct{}{},
			"diamond with a dot": struct{}{},
			"geometric":          struct{}{},
			"inside":             struct{}{},
		},
	},
	":radio_button:": Code{
		"\U0001f518",
		map[string]struct{}{
			"button":       struct{}{},
			"geometric":    struct{}{},
			"radio":        struct{}{},
			"radio button": struct{}{},
		},
	},
	":black_square_button:": Code{
		"\U0001f532",
		map[string]struct{}{
			"black square button": struct{}{},
			"button":              struct{}{},
			"geometric":           struct{}{},
			"square":              struct{}{},
		},
	},
	":white_square_button:": Code{
		"\U0001f533",
		map[string]struct{}{
			"button":              struct{}{},
			"geometric":           struct{}{},
			"outlined":            struct{}{},
			"square":              struct{}{},
			"white square button": struct{}{},
		},
	},
	":white_circle:": Code{
		"\u26aa",
		map[string]struct{}{
			"circle":       struct{}{},
			"geometric":    struct{}{},
			"white circle": struct{}{},
		},
	},
	":black_circle:": Code{
		"\u26ab",
		map[string]struct{}{
			"black circle": struct{}{},
			"circle":       struct{}{},
			"geometric":    struct{}{},
		},
	},
	":red_circle:": Code{
		"\U0001f534",
		map[string]struct{}{
			"circle":     struct{}{},
			"geometric":  struct{}{},
			"red":        struct{}{},
			"red circle": struct{}{},
		},
	},
	":blue_circle:": Code{
		"\U0001f535",
		map[string]struct{}{
			"blue":        struct{}{},
			"blue circle": struct{}{},
			"circle":      struct{}{},
			"geometric":   struct{}{},
		},
	},
	":checkered_flag:": Code{
		"\U0001f3c1",
		map[string]struct{}{
			"checkered":      struct{}{},
			"chequered":      struct{}{},
			"chequered flag": struct{}{},
			"racing":         struct{}{},
		},
	},
	":triangular_flag_on_post:": Code{
		"\U0001f6a9",
		map[string]struct{}{
			"post":            struct{}{},
			"triangular flag": struct{}{},
		},
	},
	":crossed_flags:": Code{
		"\U0001f38c",
		map[string]struct{}{
			"celebration":   struct{}{},
			"cross":         struct{}{},
			"crossed":       struct{}{},
			"crossed flags": struct{}{},
			"Japanese":      struct{}{},
		},
	},
	":flag_black:": Code{
		"\U0001f3f4",
		map[string]struct{}{
			"black flag": struct{}{},
			"waving":     struct{}{},
		},
	},
	":flag_white:": Code{
		"\U0001f3f3",
		map[string]struct{}{
			"waving":     struct{}{},
			"white flag": struct{}{},
		},
	},
	":rainbow_flag:": Code{
		"\U0001f3f3\ufe0f\u200d\U0001f308",
		map[string]struct{}{
			"rainbow":      struct{}{},
			"rainbow flag": struct{}{},
		},
	},
	":flag_ac:": Code{
		"\U0001f1e6\U0001f1e8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ad:": Code{
		"\U0001f1e6\U0001f1e9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ae:": Code{
		"\U0001f1e6\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_af:": Code{
		"\U0001f1e6\U0001f1eb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ag:": Code{
		"\U0001f1e6\U0001f1ec",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ai:": Code{
		"\U0001f1e6\U0001f1ee",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_al:": Code{
		"\U0001f1e6\U0001f1f1",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_am:": Code{
		"\U0001f1e6\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ao:": Code{
		"\U0001f1e6\U0001f1f4",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_aq:": Code{
		"\U0001f1e6\U0001f1f6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ar:": Code{
		"\U0001f1e6\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_as:": Code{
		"\U0001f1e6\U0001f1f8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_at:": Code{
		"\U0001f1e6\U0001f1f9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_au:": Code{
		"\U0001f1e6\U0001f1fa",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_aw:": Code{
		"\U0001f1e6\U0001f1fc",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ax:": Code{
		"\U0001f1e6\U0001f1fd",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_az:": Code{
		"\U0001f1e6\U0001f1ff",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ba:": Code{
		"\U0001f1e7\U0001f1e6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bb:": Code{
		"\U0001f1e7\U0001f1e7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bd:": Code{
		"\U0001f1e7\U0001f1e9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_be:": Code{
		"\U0001f1e7\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bf:": Code{
		"\U0001f1e7\U0001f1eb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bg:": Code{
		"\U0001f1e7\U0001f1ec",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bh:": Code{
		"\U0001f1e7\U0001f1ed",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bi:": Code{
		"\U0001f1e7\U0001f1ee",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bj:": Code{
		"\U0001f1e7\U0001f1ef",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bl:": Code{
		"\U0001f1e7\U0001f1f1",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bm:": Code{
		"\U0001f1e7\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bn:": Code{
		"\U0001f1e7\U0001f1f3",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bo:": Code{
		"\U0001f1e7\U0001f1f4",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bq:": Code{
		"\U0001f1e7\U0001f1f6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_br:": Code{
		"\U0001f1e7\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bs:": Code{
		"\U0001f1e7\U0001f1f8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bt:": Code{
		"\U0001f1e7\U0001f1f9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bv:": Code{
		"\U0001f1e7\U0001f1fb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bw:": Code{
		"\U0001f1e7\U0001f1fc",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_by:": Code{
		"\U0001f1e7\U0001f1fe",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_bz:": Code{
		"\U0001f1e7\U0001f1ff",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ca:": Code{
		"\U0001f1e8\U0001f1e6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cc:": Code{
		"\U0001f1e8\U0001f1e8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cd:": Code{
		"\U0001f1e8\U0001f1e9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cf:": Code{
		"\U0001f1e8\U0001f1eb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cg:": Code{
		"\U0001f1e8\U0001f1ec",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ch:": Code{
		"\U0001f1e8\U0001f1ed",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ci:": Code{
		"\U0001f1e8\U0001f1ee",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ck:": Code{
		"\U0001f1e8\U0001f1f0",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cl:": Code{
		"\U0001f1e8\U0001f1f1",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cm:": Code{
		"\U0001f1e8\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cn:": Code{
		"\U0001f1e8\U0001f1f3",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_co:": Code{
		"\U0001f1e8\U0001f1f4",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cp:": Code{
		"\U0001f1e8\U0001f1f5",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cr:": Code{
		"\U0001f1e8\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cu:": Code{
		"\U0001f1e8\U0001f1fa",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cv:": Code{
		"\U0001f1e8\U0001f1fb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cw:": Code{
		"\U0001f1e8\U0001f1fc",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cx:": Code{
		"\U0001f1e8\U0001f1fd",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cy:": Code{
		"\U0001f1e8\U0001f1fe",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_cz:": Code{
		"\U0001f1e8\U0001f1ff",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_de:": Code{
		"\U0001f1e9\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_dg:": Code{
		"\U0001f1e9\U0001f1ec",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_dj:": Code{
		"\U0001f1e9\U0001f1ef",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_dk:": Code{
		"\U0001f1e9\U0001f1f0",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_dm:": Code{
		"\U0001f1e9\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_do:": Code{
		"\U0001f1e9\U0001f1f4",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_dz:": Code{
		"\U0001f1e9\U0001f1ff",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ea:": Code{
		"\U0001f1ea\U0001f1e6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ec:": Code{
		"\U0001f1ea\U0001f1e8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ee:": Code{
		"\U0001f1ea\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_eg:": Code{
		"\U0001f1ea\U0001f1ec",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_eh:": Code{
		"\U0001f1ea\U0001f1ed",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_er:": Code{
		"\U0001f1ea\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_es:": Code{
		"\U0001f1ea\U0001f1f8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_et:": Code{
		"\U0001f1ea\U0001f1f9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_eu:": Code{
		"\U0001f1ea\U0001f1fa",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_fi:": Code{
		"\U0001f1eb\U0001f1ee",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_fj:": Code{
		"\U0001f1eb\U0001f1ef",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_fk:": Code{
		"\U0001f1eb\U0001f1f0",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_fm:": Code{
		"\U0001f1eb\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_fo:": Code{
		"\U0001f1eb\U0001f1f4",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_fr:": Code{
		"\U0001f1eb\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ga:": Code{
		"\U0001f1ec\U0001f1e6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gb:": Code{
		"\U0001f1ec\U0001f1e7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gd:": Code{
		"\U0001f1ec\U0001f1e9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ge:": Code{
		"\U0001f1ec\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gf:": Code{
		"\U0001f1ec\U0001f1eb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gg:": Code{
		"\U0001f1ec\U0001f1ec",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gh:": Code{
		"\U0001f1ec\U0001f1ed",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gi:": Code{
		"\U0001f1ec\U0001f1ee",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gl:": Code{
		"\U0001f1ec\U0001f1f1",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gm:": Code{
		"\U0001f1ec\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gn:": Code{
		"\U0001f1ec\U0001f1f3",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gp:": Code{
		"\U0001f1ec\U0001f1f5",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gq:": Code{
		"\U0001f1ec\U0001f1f6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gr:": Code{
		"\U0001f1ec\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gs:": Code{
		"\U0001f1ec\U0001f1f8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gt:": Code{
		"\U0001f1ec\U0001f1f9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gu:": Code{
		"\U0001f1ec\U0001f1fa",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gw:": Code{
		"\U0001f1ec\U0001f1fc",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_gy:": Code{
		"\U0001f1ec\U0001f1fe",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_hk:": Code{
		"\U0001f1ed\U0001f1f0",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_hm:": Code{
		"\U0001f1ed\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_hn:": Code{
		"\U0001f1ed\U0001f1f3",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_hr:": Code{
		"\U0001f1ed\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ht:": Code{
		"\U0001f1ed\U0001f1f9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_hu:": Code{
		"\U0001f1ed\U0001f1fa",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ic:": Code{
		"\U0001f1ee\U0001f1e8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_id:": Code{
		"\U0001f1ee\U0001f1e9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ie:": Code{
		"\U0001f1ee\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_il:": Code{
		"\U0001f1ee\U0001f1f1",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_im:": Code{
		"\U0001f1ee\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_in:": Code{
		"\U0001f1ee\U0001f1f3",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_io:": Code{
		"\U0001f1ee\U0001f1f4",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_iq:": Code{
		"\U0001f1ee\U0001f1f6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ir:": Code{
		"\U0001f1ee\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_is:": Code{
		"\U0001f1ee\U0001f1f8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_it:": Code{
		"\U0001f1ee\U0001f1f9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_je:": Code{
		"\U0001f1ef\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_jm:": Code{
		"\U0001f1ef\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_jo:": Code{
		"\U0001f1ef\U0001f1f4",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_jp:": Code{
		"\U0001f1ef\U0001f1f5",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ke:": Code{
		"\U0001f1f0\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_kg:": Code{
		"\U0001f1f0\U0001f1ec",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_kh:": Code{
		"\U0001f1f0\U0001f1ed",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ki:": Code{
		"\U0001f1f0\U0001f1ee",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_km:": Code{
		"\U0001f1f0\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_kn:": Code{
		"\U0001f1f0\U0001f1f3",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_kp:": Code{
		"\U0001f1f0\U0001f1f5",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_kr:": Code{
		"\U0001f1f0\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_kw:": Code{
		"\U0001f1f0\U0001f1fc",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ky:": Code{
		"\U0001f1f0\U0001f1fe",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_kz:": Code{
		"\U0001f1f0\U0001f1ff",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_la:": Code{
		"\U0001f1f1\U0001f1e6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_lb:": Code{
		"\U0001f1f1\U0001f1e7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_lc:": Code{
		"\U0001f1f1\U0001f1e8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_li:": Code{
		"\U0001f1f1\U0001f1ee",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_lk:": Code{
		"\U0001f1f1\U0001f1f0",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_lr:": Code{
		"\U0001f1f1\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ls:": Code{
		"\U0001f1f1\U0001f1f8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_lt:": Code{
		"\U0001f1f1\U0001f1f9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_lu:": Code{
		"\U0001f1f1\U0001f1fa",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_lv:": Code{
		"\U0001f1f1\U0001f1fb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ly:": Code{
		"\U0001f1f1\U0001f1fe",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ma:": Code{
		"\U0001f1f2\U0001f1e6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mc:": Code{
		"\U0001f1f2\U0001f1e8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_md:": Code{
		"\U0001f1f2\U0001f1e9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_me:": Code{
		"\U0001f1f2\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mf:": Code{
		"\U0001f1f2\U0001f1eb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mg:": Code{
		"\U0001f1f2\U0001f1ec",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mh:": Code{
		"\U0001f1f2\U0001f1ed",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mk:": Code{
		"\U0001f1f2\U0001f1f0",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ml:": Code{
		"\U0001f1f2\U0001f1f1",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mm:": Code{
		"\U0001f1f2\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mn:": Code{
		"\U0001f1f2\U0001f1f3",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mo:": Code{
		"\U0001f1f2\U0001f1f4",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mp:": Code{
		"\U0001f1f2\U0001f1f5",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mq:": Code{
		"\U0001f1f2\U0001f1f6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mr:": Code{
		"\U0001f1f2\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ms:": Code{
		"\U0001f1f2\U0001f1f8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mt:": Code{
		"\U0001f1f2\U0001f1f9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mu:": Code{
		"\U0001f1f2\U0001f1fa",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mv:": Code{
		"\U0001f1f2\U0001f1fb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mw:": Code{
		"\U0001f1f2\U0001f1fc",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mx:": Code{
		"\U0001f1f2\U0001f1fd",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_my:": Code{
		"\U0001f1f2\U0001f1fe",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_mz:": Code{
		"\U0001f1f2\U0001f1ff",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_na:": Code{
		"\U0001f1f3\U0001f1e6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_nc:": Code{
		"\U0001f1f3\U0001f1e8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ne:": Code{
		"\U0001f1f3\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_nf:": Code{
		"\U0001f1f3\U0001f1eb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ng:": Code{
		"\U0001f1f3\U0001f1ec",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ni:": Code{
		"\U0001f1f3\U0001f1ee",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_nl:": Code{
		"\U0001f1f3\U0001f1f1",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_no:": Code{
		"\U0001f1f3\U0001f1f4",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_np:": Code{
		"\U0001f1f3\U0001f1f5",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_nr:": Code{
		"\U0001f1f3\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_nu:": Code{
		"\U0001f1f3\U0001f1fa",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_nz:": Code{
		"\U0001f1f3\U0001f1ff",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_om:": Code{
		"\U0001f1f4\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_pa:": Code{
		"\U0001f1f5\U0001f1e6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_pe:": Code{
		"\U0001f1f5\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_pf:": Code{
		"\U0001f1f5\U0001f1eb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_pg:": Code{
		"\U0001f1f5\U0001f1ec",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ph:": Code{
		"\U0001f1f5\U0001f1ed",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_pk:": Code{
		"\U0001f1f5\U0001f1f0",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_pl:": Code{
		"\U0001f1f5\U0001f1f1",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_pm:": Code{
		"\U0001f1f5\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_pn:": Code{
		"\U0001f1f5\U0001f1f3",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_pr:": Code{
		"\U0001f1f5\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ps:": Code{
		"\U0001f1f5\U0001f1f8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_pt:": Code{
		"\U0001f1f5\U0001f1f9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_pw:": Code{
		"\U0001f1f5\U0001f1fc",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_py:": Code{
		"\U0001f1f5\U0001f1fe",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_qa:": Code{
		"\U0001f1f6\U0001f1e6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_re:": Code{
		"\U0001f1f7\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ro:": Code{
		"\U0001f1f7\U0001f1f4",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_rs:": Code{
		"\U0001f1f7\U0001f1f8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ru:": Code{
		"\U0001f1f7\U0001f1fa",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_rw:": Code{
		"\U0001f1f7\U0001f1fc",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sa:": Code{
		"\U0001f1f8\U0001f1e6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sb:": Code{
		"\U0001f1f8\U0001f1e7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sc:": Code{
		"\U0001f1f8\U0001f1e8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sd:": Code{
		"\U0001f1f8\U0001f1e9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_se:": Code{
		"\U0001f1f8\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sg:": Code{
		"\U0001f1f8\U0001f1ec",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sh:": Code{
		"\U0001f1f8\U0001f1ed",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_si:": Code{
		"\U0001f1f8\U0001f1ee",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sj:": Code{
		"\U0001f1f8\U0001f1ef",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sk:": Code{
		"\U0001f1f8\U0001f1f0",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sl:": Code{
		"\U0001f1f8\U0001f1f1",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sm:": Code{
		"\U0001f1f8\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sn:": Code{
		"\U0001f1f8\U0001f1f3",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_so:": Code{
		"\U0001f1f8\U0001f1f4",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sr:": Code{
		"\U0001f1f8\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ss:": Code{
		"\U0001f1f8\U0001f1f8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_st:": Code{
		"\U0001f1f8\U0001f1f9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sv:": Code{
		"\U0001f1f8\U0001f1fb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sx:": Code{
		"\U0001f1f8\U0001f1fd",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sy:": Code{
		"\U0001f1f8\U0001f1fe",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_sz:": Code{
		"\U0001f1f8\U0001f1ff",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ta:": Code{
		"\U0001f1f9\U0001f1e6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_tc:": Code{
		"\U0001f1f9\U0001f1e8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_td:": Code{
		"\U0001f1f9\U0001f1e9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_tf:": Code{
		"\U0001f1f9\U0001f1eb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_tg:": Code{
		"\U0001f1f9\U0001f1ec",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_th:": Code{
		"\U0001f1f9\U0001f1ed",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_tj:": Code{
		"\U0001f1f9\U0001f1ef",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_tk:": Code{
		"\U0001f1f9\U0001f1f0",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_tl:": Code{
		"\U0001f1f9\U0001f1f1",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_tm:": Code{
		"\U0001f1f9\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_tn:": Code{
		"\U0001f1f9\U0001f1f3",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_to:": Code{
		"\U0001f1f9\U0001f1f4",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_tr:": Code{
		"\U0001f1f9\U0001f1f7",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_tt:": Code{
		"\U0001f1f9\U0001f1f9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_tv:": Code{
		"\U0001f1f9\U0001f1fb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_tw:": Code{
		"\U0001f1f9\U0001f1fc",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_tz:": Code{
		"\U0001f1f9\U0001f1ff",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ua:": Code{
		"\U0001f1fa\U0001f1e6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ug:": Code{
		"\U0001f1fa\U0001f1ec",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_um:": Code{
		"\U0001f1fa\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":united_nations:": Code{
		"\U0001f1fa\U0001f1f3",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_us:": Code{
		"\U0001f1fa\U0001f1f8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_uy:": Code{
		"\U0001f1fa\U0001f1fe",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_uz:": Code{
		"\U0001f1fa\U0001f1ff",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_va:": Code{
		"\U0001f1fb\U0001f1e6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_vc:": Code{
		"\U0001f1fb\U0001f1e8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ve:": Code{
		"\U0001f1fb\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_vg:": Code{
		"\U0001f1fb\U0001f1ec",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_vi:": Code{
		"\U0001f1fb\U0001f1ee",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_vn:": Code{
		"\U0001f1fb\U0001f1f3",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_vu:": Code{
		"\U0001f1fb\U0001f1fa",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_wf:": Code{
		"\U0001f1fc\U0001f1eb",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ws:": Code{
		"\U0001f1fc\U0001f1f8",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_xk:": Code{
		"\U0001f1fd\U0001f1f0",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_ye:": Code{
		"\U0001f1fe\U0001f1ea",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_yt:": Code{
		"\U0001f1fe\U0001f1f9",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_za:": Code{
		"\U0001f1ff\U0001f1e6",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_zm:": Code{
		"\U0001f1ff\U0001f1f2",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":flag_zw:": Code{
		"\U0001f1ff\U0001f1fc",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":england:": Code{
		"\U0001f3f4\U000e0067\U000e0062\U000e0065\U000e006e\U000e0067\U000e007f",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":scotland:": Code{
		"\U0001f3f4\U000e0067\U000e0062\U000e0073\U000e0063\U000e0074\U000e007f",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":wales:": Code{
		"\U0001f3f4\U000e0067\U000e0062\U000e0077\U000e006c\U000e0073\U000e007f",
		map[string]struct{}{
			"flag": struct{}{},
		},
	},
	":smile:": Code{
		"\U0001f604",
		map[string]struct{}{
			"happy":   struct{}{},
			"joy":     struct{}{},
			"laugh":   struct{}{},
			"pleased": struct{}{},
		},
	},
	":satisfied:": Code{
		"\U0001f606",
		map[string]struct{}{
			"happy": struct{}{},
			"haha":  struct{}{},
		},
	},
	":slightly_smiling_face:": Code{
		"\U0001f642",
		map[string]struct{}{},
	},
	":upside_down_face:": Code{
		"\U0001f643",
		map[string]struct{}{},
	},
	":smiling_face_with_three_hearts:": Code{
		"\U0001f970",
		map[string]struct{}{
			"love": struct{}{},
		},
	},
	":smiling_face_with_tear:": Code{
		"\U0001f972",
		map[string]struct{}{},
	},
	":zany_face:": Code{
		"\U0001f92a",
		map[string]struct{}{
			"goofy": struct{}{},
			"wacky": struct{}{},
		},
	},
	":money_mouth_face:": Code{
		"\U0001f911",
		map[string]struct{}{
			"rich": struct{}{},
		},
	},
	":hugs:": Code{
		"\U0001f917",
		map[string]struct{}{},
	},
	":hand_over_mouth:": Code{
		"\U0001f92d",
		map[string]struct{}{
			"quiet":  struct{}{},
			"whoops": struct{}{},
		},
	},
	":zipper_mouth_face:": Code{
		"\U0001f910",
		map[string]struct{}{
			"silence": struct{}{},
			"hush":    struct{}{},
		},
	},
	":raised_eyebrow:": Code{
		"\U0001f928",
		map[string]struct{}{
			"suspicious": struct{}{},
		},
	},
	":roll_eyes:": Code{
		"\U0001f644",
		map[string]struct{}{},
	},
	":face_with_thermometer:": Code{
		"\U0001f912",
		map[string]struct{}{
			"sick": struct{}{},
		},
	},
	":face_with_head_bandage:": Code{
		"\U0001f915",
		map[string]struct{}{
			"hurt": struct{}{},
		},
	},
	":vomiting_face:": Code{
		"\U0001f92e",
		map[string]struct{}{
			"barf": struct{}{},
			"sick": struct{}{},
		},
	},
	":hot_face:": Code{
		"\U0001f975",
		map[string]struct{}{
			"heat":     struct{}{},
			"sweating": struct{}{},
		},
	},
	":cold_face:": Code{
		"\U0001f976",
		map[string]struct{}{
			"freezing": struct{}{},
			"ice":      struct{}{},
		},
	},
	":woozy_face:": Code{
		"\U0001f974",
		map[string]struct{}{
			"groggy": struct{}{},
		},
	},
	":cowboy_hat_face:": Code{
		"\U0001f920",
		map[string]struct{}{},
	},
	":partying_face:": Code{
		"\U0001f973",
		map[string]struct{}{
			"celebration": struct{}{},
			"birthday":    struct{}{},
		},
	},
	":disguised_face:": Code{
		"\U0001f978",
		map[string]struct{}{},
	},
	":nerd_face:": Code{
		"\U0001f913",
		map[string]struct{}{
			"geek":    struct{}{},
			"glasses": struct{}{},
		},
	},
	":monocle_face:": Code{
		"\U0001f9d0",
		map[string]struct{}{},
	},
	":slightly_frowning_face:": Code{
		"\U0001f641",
		map[string]struct{}{},
	},
	":frowning_face:": Code{
		"\u2639\ufe0f",
		map[string]struct{}{},
	},
	":pleading_face:": Code{
		"\U0001f97a",
		map[string]struct{}{
			"puppy": struct{}{},
			"eyes":  struct{}{},
		},
	},
	":yawning_face:": Code{
		"\U0001f971",
		map[string]struct{}{},
	},
	":pout:": Code{
		"\U0001f621",
		map[string]struct{}{
			"angry": struct{}{},
		},
	},
	":cursing_face:": Code{
		"\U0001f92c",
		map[string]struct{}{
			"foul": struct{}{},
		},
	},
	":skull_and_crossbones:": Code{
		"\u2620\ufe0f",
		map[string]struct{}{
			"danger": struct{}{},
			"pirate": struct{}{},
		},
	},
	":hankey:": Code{
		"\U0001f4a9",
		map[string]struct{}{
			"crap": struct{}{},
		},
	},
	":shit:": Code{
		"\U0001f4a9",
		map[string]struct{}{
			"crap": struct{}{},
		},
	},
	":clown_face:": Code{
		"\U0001f921",
		map[string]struct{}{},
	},
	":heavy_heart_exclamation:": Code{
		"\u2763\ufe0f",
		map[string]struct{}{},
	},
	":brown_heart:": Code{
		"\U0001f90e",
		map[string]struct{}{},
	},
	":white_heart:": Code{
		"\U0001f90d",
		map[string]struct{}{},
	},
	":collision:": Code{
		"\U0001f4a5",
		map[string]struct{}{
			"explode": struct{}{},
		},
	},
	":eye_speech_bubble:": Code{
		"\U0001f441\ufe0f\u200d\U0001f5e8\ufe0f",
		map[string]struct{}{},
	},
	":left_speech_bubble:": Code{
		"\U0001f5e8\ufe0f",
		map[string]struct{}{},
	},
	":right_anger_bubble:": Code{
		"\U0001f5ef\ufe0f",
		map[string]struct{}{},
	},
	":raised_hand_with_fingers_splayed:": Code{
		"\U0001f590\ufe0f",
		map[string]struct{}{},
	},
	":hand:": Code{
		"\u270b",
		map[string]struct{}{
			"highfive": struct{}{},
			"stop":     struct{}{},
		},
	},
	":vulcan_salute:": Code{
		"\U0001f596",
		map[string]struct{}{
			"prosper": struct{}{},
			"spock":   struct{}{},
		},
	},
	":pinched_fingers:": Code{
		"\U0001f90c",
		map[string]struct{}{},
	},
	":pinching_hand:": Code{
		"\U0001f90f",
		map[string]struct{}{},
	},
	":crossed_fingers:": Code{
		"\U0001f91e",
		map[string]struct{}{
			"luck":    struct{}{},
			"hopeful": struct{}{},
		},
	},
	":call_me_hand:": Code{
		"\U0001f919",
		map[string]struct{}{},
	},
	":fu:": Code{
		"\U0001f595",
		map[string]struct{}{},
	},
	":+1:": Code{
		"\U0001f44d",
		map[string]struct{}{
			"approve": struct{}{},
			"ok":      struct{}{},
		},
	},
	":-1:": Code{
		"\U0001f44e",
		map[string]struct{}{
			"disapprove": struct{}{},
			"bury":       struct{}{},
		},
	},
	":fist_raised:": Code{
		"\u270a",
		map[string]struct{}{
			"power": struct{}{},
		},
	},
	":fist_oncoming:": Code{
		"\U0001f44a",
		map[string]struct{}{
			"attack": struct{}{},
		},
	},
	":facepunch:": Code{
		"\U0001f44a",
		map[string]struct{}{
			"attack": struct{}{},
		},
	},
	":fist_left:": Code{
		"\U0001f91b",
		map[string]struct{}{},
	},
	":fist_right:": Code{
		"\U0001f91c",
		map[string]struct{}{},
	},
	":mechanical_arm:": Code{
		"\U0001f9be",
		map[string]struct{}{},
	},
	":mechanical_leg:": Code{
		"\U0001f9bf",
		map[string]struct{}{},
	},
	":leg:": Code{
		"\U0001f9b5",
		map[string]struct{}{},
	},
	":foot:": Code{
		"\U0001f9b6",
		map[string]struct{}{},
	},
	":ear_with_hearing_aid:": Code{
		"\U0001f9bb",
		map[string]struct{}{},
	},
	":anatomical_heart:": Code{
		"\U0001fac0",
		map[string]struct{}{},
	},
	":lungs:": Code{
		"\U0001fac1",
		map[string]struct{}{},
	},
	":tooth:": Code{
		"\U0001f9b7",
		map[string]struct{}{},
	},
	":bone:": Code{
		"\U0001f9b4",
		map[string]struct{}{},
	},
	":red_haired_man:": Code{
		"\U0001f468\u200d\U0001f9b0",
		map[string]struct{}{},
	},
	":curly_haired_man:": Code{
		"\U0001f468\u200d\U0001f9b1",
		map[string]struct{}{},
	},
	":white_haired_man:": Code{
		"\U0001f468\u200d\U0001f9b3",
		map[string]struct{}{},
	},
	":bald_man:": Code{
		"\U0001f468\u200d\U0001f9b2",
		map[string]struct{}{},
	},
	":red_haired_woman:": Code{
		"\U0001f469\u200d\U0001f9b0",
		map[string]struct{}{},
	},
	":person_red_hair:": Code{
		"\U0001f9d1\u200d\U0001f9b0",
		map[string]struct{}{},
	},
	":curly_haired_woman:": Code{
		"\U0001f469\u200d\U0001f9b1",
		map[string]struct{}{},
	},
	":person_curly_hair:": Code{
		"\U0001f9d1\u200d\U0001f9b1",
		map[string]struct{}{},
	},
	":white_haired_woman:": Code{
		"\U0001f469\u200d\U0001f9b3",
		map[string]struct{}{},
	},
	":person_white_hair:": Code{
		"\U0001f9d1\u200d\U0001f9b3",
		map[string]struct{}{},
	},
	":bald_woman:": Code{
		"\U0001f469\u200d\U0001f9b2",
		map[string]struct{}{},
	},
	":person_bald:": Code{
		"\U0001f9d1\u200d\U0001f9b2",
		map[string]struct{}{},
	},
	":blond_haired_woman:": Code{
		"\U0001f471\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":blonde_woman:": Code{
		"\U0001f471\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":blond_haired_man:": Code{
		"\U0001f471\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":frowning_person:": Code{
		"\U0001f64d",
		map[string]struct{}{},
	},
	":frowning_man:": Code{
		"\U0001f64d\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":frowning_woman:": Code{
		"\U0001f64d\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":pouting_face:": Code{
		"\U0001f64e",
		map[string]struct{}{},
	},
	":pouting_man:": Code{
		"\U0001f64e\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":pouting_woman:": Code{
		"\U0001f64e\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":no_good:": Code{
		"\U0001f645",
		map[string]struct{}{
			"stop":   struct{}{},
			"halt":   struct{}{},
			"denied": struct{}{},
		},
	},
	":no_good_man:": Code{
		"\U0001f645\u200d\u2642\ufe0f",
		map[string]struct{}{
			"stop":   struct{}{},
			"halt":   struct{}{},
			"denied": struct{}{},
		},
	},
	":ng_man:": Code{
		"\U0001f645\u200d\u2642\ufe0f",
		map[string]struct{}{
			"stop":   struct{}{},
			"halt":   struct{}{},
			"denied": struct{}{},
		},
	},
	":no_good_woman:": Code{
		"\U0001f645\u200d\u2640\ufe0f",
		map[string]struct{}{
			"stop":   struct{}{},
			"halt":   struct{}{},
			"denied": struct{}{},
		},
	},
	":ng_woman:": Code{
		"\U0001f645\u200d\u2640\ufe0f",
		map[string]struct{}{
			"stop":   struct{}{},
			"halt":   struct{}{},
			"denied": struct{}{},
		},
	},
	":ok_person:": Code{
		"\U0001f646",
		map[string]struct{}{},
	},
	":ok_man:": Code{
		"\U0001f646\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":ok_woman:": Code{
		"\U0001f646\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":tipping_hand_person:": Code{
		"\U0001f481",
		map[string]struct{}{},
	},
	":information_desk_person:": Code{
		"\U0001f481",
		map[string]struct{}{},
	},
	":tipping_hand_man:": Code{
		"\U0001f481\u200d\u2642\ufe0f",
		map[string]struct{}{
			"information": struct{}{},
		},
	},
	":sassy_man:": Code{
		"\U0001f481\u200d\u2642\ufe0f",
		map[string]struct{}{
			"information": struct{}{},
		},
	},
	":tipping_hand_woman:": Code{
		"\U0001f481\u200d\u2640\ufe0f",
		map[string]struct{}{
			"information": struct{}{},
		},
	},
	":sassy_woman:": Code{
		"\U0001f481\u200d\u2640\ufe0f",
		map[string]struct{}{
			"information": struct{}{},
		},
	},
	":raising_hand:": Code{
		"\U0001f64b",
		map[string]struct{}{},
	},
	":raising_hand_man:": Code{
		"\U0001f64b\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":raising_hand_woman:": Code{
		"\U0001f64b\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":deaf_person:": Code{
		"\U0001f9cf",
		map[string]struct{}{},
	},
	":deaf_man:": Code{
		"\U0001f9cf\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":deaf_woman:": Code{
		"\U0001f9cf\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":bow:": Code{
		"\U0001f647",
		map[string]struct{}{
			"respect": struct{}{},
			"thanks":  struct{}{},
		},
	},
	":bowing_man:": Code{
		"\U0001f647\u200d\u2642\ufe0f",
		map[string]struct{}{
			"respect": struct{}{},
			"thanks":  struct{}{},
		},
	},
	":bowing_woman:": Code{
		"\U0001f647\u200d\u2640\ufe0f",
		map[string]struct{}{
			"respect": struct{}{},
			"thanks":  struct{}{},
		},
	},
	":facepalm:": Code{
		"\U0001f926",
		map[string]struct{}{},
	},
	":shrug:": Code{
		"\U0001f937",
		map[string]struct{}{},
	},
	":health_worker:": Code{
		"\U0001f9d1\u200d\u2695\ufe0f",
		map[string]struct{}{},
	},
	":student:": Code{
		"\U0001f9d1\u200d\U0001f393",
		map[string]struct{}{},
	},
	":teacher:": Code{
		"\U0001f9d1\u200d\U0001f3eb",
		map[string]struct{}{},
	},
	":judge:": Code{
		"\U0001f9d1\u200d\u2696\ufe0f",
		map[string]struct{}{},
	},
	":farmer:": Code{
		"\U0001f9d1\u200d\U0001f33e",
		map[string]struct{}{},
	},
	":cook:": Code{
		"\U0001f9d1\u200d\U0001f373",
		map[string]struct{}{},
	},
	":mechanic:": Code{
		"\U0001f9d1\u200d\U0001f527",
		map[string]struct{}{},
	},
	":factory_worker:": Code{
		"\U0001f9d1\u200d\U0001f3ed",
		map[string]struct{}{},
	},
	":office_worker:": Code{
		"\U0001f9d1\u200d\U0001f4bc",
		map[string]struct{}{},
	},
	":scientist:": Code{
		"\U0001f9d1\u200d\U0001f52c",
		map[string]struct{}{},
	},
	":technologist:": Code{
		"\U0001f9d1\u200d\U0001f4bb",
		map[string]struct{}{},
	},
	":singer:": Code{
		"\U0001f9d1\u200d\U0001f3a4",
		map[string]struct{}{},
	},
	":artist:": Code{
		"\U0001f9d1\u200d\U0001f3a8",
		map[string]struct{}{},
	},
	":pilot:": Code{
		"\U0001f9d1\u200d\u2708\ufe0f",
		map[string]struct{}{},
	},
	":astronaut:": Code{
		"\U0001f9d1\u200d\U0001f680",
		map[string]struct{}{},
	},
	":firefighter:": Code{
		"\U0001f9d1\u200d\U0001f692",
		map[string]struct{}{},
	},
	":cop:": Code{
		"\U0001f46e",
		map[string]struct{}{
			"law": struct{}{},
		},
	},
	":policeman:": Code{
		"\U0001f46e\u200d\u2642\ufe0f",
		map[string]struct{}{
			"law": struct{}{},
			"cop": struct{}{},
		},
	},
	":policewoman:": Code{
		"\U0001f46e\u200d\u2640\ufe0f",
		map[string]struct{}{
			"law": struct{}{},
			"cop": struct{}{},
		},
	},
	":male_detective:": Code{
		"\U0001f575\ufe0f\u200d\u2642\ufe0f",
		map[string]struct{}{
			"sleuth": struct{}{},
		},
	},
	":female_detective:": Code{
		"\U0001f575\ufe0f\u200d\u2640\ufe0f",
		map[string]struct{}{
			"sleuth": struct{}{},
		},
	},
	":guardsman:": Code{
		"\U0001f482\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":guardswoman:": Code{
		"\U0001f482\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":ninja:": Code{
		"\U0001f977",
		map[string]struct{}{},
	},
	":construction_worker_man:": Code{
		"\U0001f477\u200d\u2642\ufe0f",
		map[string]struct{}{
			"helmet": struct{}{},
		},
	},
	":construction_worker_woman:": Code{
		"\U0001f477\u200d\u2640\ufe0f",
		map[string]struct{}{
			"helmet": struct{}{},
		},
	},
	":person_with_turban:": Code{
		"\U0001f473",
		map[string]struct{}{},
	},
	":man_with_turban:": Code{
		"\U0001f473\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":woman_with_turban:": Code{
		"\U0001f473\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":man_with_gua_pi_mao:": Code{
		"\U0001f472",
		map[string]struct{}{},
	},
	":person_in_tuxedo:": Code{
		"\U0001f935",
		map[string]struct{}{
			"groom":    struct{}{},
			"marriage": struct{}{},
			"wedding":  struct{}{},
		},
	},
	":woman_in_tuxedo:": Code{
		"\U0001f935\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":person_with_veil:": Code{
		"\U0001f470",
		map[string]struct{}{
			"marriage": struct{}{},
			"wedding":  struct{}{},
		},
	},
	":man_with_veil:": Code{
		"\U0001f470\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":woman_with_veil:": Code{
		"\U0001f470\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":woman_feeding_baby:": Code{
		"\U0001f469\u200d\U0001f37c",
		map[string]struct{}{},
	},
	":man_feeding_baby:": Code{
		"\U0001f468\u200d\U0001f37c",
		map[string]struct{}{},
	},
	":person_feeding_baby:": Code{
		"\U0001f9d1\u200d\U0001f37c",
		map[string]struct{}{},
	},
	":mx_claus:": Code{
		"\U0001f9d1\u200d\U0001f384",
		map[string]struct{}{},
	},
	":superhero:": Code{
		"\U0001f9b8",
		map[string]struct{}{},
	},
	":superhero_man:": Code{
		"\U0001f9b8\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":superhero_woman:": Code{
		"\U0001f9b8\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":supervillain:": Code{
		"\U0001f9b9",
		map[string]struct{}{},
	},
	":supervillain_man:": Code{
		"\U0001f9b9\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":supervillain_woman:": Code{
		"\U0001f9b9\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":mage_man:": Code{
		"\U0001f9d9\u200d\u2642\ufe0f",
		map[string]struct{}{
			"wizard": struct{}{},
		},
	},
	":mage_woman:": Code{
		"\U0001f9d9\u200d\u2640\ufe0f",
		map[string]struct{}{
			"wizard": struct{}{},
		},
	},
	":fairy_man:": Code{
		"\U0001f9da\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":fairy_woman:": Code{
		"\U0001f9da\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":vampire_man:": Code{
		"\U0001f9db\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":vampire_woman:": Code{
		"\U0001f9db\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":elf_man:": Code{
		"\U0001f9dd\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":elf_woman:": Code{
		"\U0001f9dd\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":genie_man:": Code{
		"\U0001f9de\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":genie_woman:": Code{
		"\U0001f9de\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":zombie_man:": Code{
		"\U0001f9df\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":zombie_woman:": Code{
		"\U0001f9df\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":massage:": Code{
		"\U0001f486",
		map[string]struct{}{
			"spa": struct{}{},
		},
	},
	":massage_man:": Code{
		"\U0001f486\u200d\u2642\ufe0f",
		map[string]struct{}{
			"spa": struct{}{},
		},
	},
	":massage_woman:": Code{
		"\U0001f486\u200d\u2640\ufe0f",
		map[string]struct{}{
			"spa": struct{}{},
		},
	},
	":haircut:": Code{
		"\U0001f487",
		map[string]struct{}{
			"beauty": struct{}{},
		},
	},
	":haircut_man:": Code{
		"\U0001f487\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":haircut_woman:": Code{
		"\U0001f487\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":walking:": Code{
		"\U0001f6b6",
		map[string]struct{}{},
	},
	":walking_man:": Code{
		"\U0001f6b6\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":walking_woman:": Code{
		"\U0001f6b6\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":standing_person:": Code{
		"\U0001f9cd",
		map[string]struct{}{},
	},
	":standing_man:": Code{
		"\U0001f9cd\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":standing_woman:": Code{
		"\U0001f9cd\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":kneeling_person:": Code{
		"\U0001f9ce",
		map[string]struct{}{},
	},
	":kneeling_man:": Code{
		"\U0001f9ce\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":kneeling_woman:": Code{
		"\U0001f9ce\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":person_with_probing_cane:": Code{
		"\U0001f9d1\u200d\U0001f9af",
		map[string]struct{}{},
	},
	":man_with_probing_cane:": Code{
		"\U0001f468\u200d\U0001f9af",
		map[string]struct{}{},
	},
	":woman_with_probing_cane:": Code{
		"\U0001f469\u200d\U0001f9af",
		map[string]struct{}{},
	},
	":person_in_motorized_wheelchair:": Code{
		"\U0001f9d1\u200d\U0001f9bc",
		map[string]struct{}{},
	},
	":man_in_motorized_wheelchair:": Code{
		"\U0001f468\u200d\U0001f9bc",
		map[string]struct{}{},
	},
	":woman_in_motorized_wheelchair:": Code{
		"\U0001f469\u200d\U0001f9bc",
		map[string]struct{}{},
	},
	":person_in_manual_wheelchair:": Code{
		"\U0001f9d1\u200d\U0001f9bd",
		map[string]struct{}{},
	},
	":man_in_manual_wheelchair:": Code{
		"\U0001f468\u200d\U0001f9bd",
		map[string]struct{}{},
	},
	":woman_in_manual_wheelchair:": Code{
		"\U0001f469\u200d\U0001f9bd",
		map[string]struct{}{},
	},
	":runner:": Code{
		"\U0001f3c3",
		map[string]struct{}{
			"exercise": struct{}{},
			"workout":  struct{}{},
			"marathon": struct{}{},
		},
	},
	":running:": Code{
		"\U0001f3c3",
		map[string]struct{}{
			"exercise": struct{}{},
			"workout":  struct{}{},
			"marathon": struct{}{},
		},
	},
	":running_man:": Code{
		"\U0001f3c3\u200d\u2642\ufe0f",
		map[string]struct{}{
			"exercise": struct{}{},
			"workout":  struct{}{},
			"marathon": struct{}{},
		},
	},
	":running_woman:": Code{
		"\U0001f3c3\u200d\u2640\ufe0f",
		map[string]struct{}{
			"exercise": struct{}{},
			"workout":  struct{}{},
			"marathon": struct{}{},
		},
	},
	":woman_dancing:": Code{
		"\U0001f483",
		map[string]struct{}{
			"dress": struct{}{},
		},
	},
	":business_suit_levitating:": Code{
		"\U0001f574\ufe0f",
		map[string]struct{}{},
	},
	":dancers:": Code{
		"\U0001f46f",
		map[string]struct{}{
			"bunny": struct{}{},
		},
	},
	":dancing_men:": Code{
		"\U0001f46f\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bunny": struct{}{},
		},
	},
	":dancing_women:": Code{
		"\U0001f46f\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bunny": struct{}{},
		},
	},
	":sauna_person:": Code{
		"\U0001f9d6",
		map[string]struct{}{
			"steamy": struct{}{},
		},
	},
	":sauna_man:": Code{
		"\U0001f9d6\u200d\u2642\ufe0f",
		map[string]struct{}{
			"steamy": struct{}{},
		},
	},
	":sauna_woman:": Code{
		"\U0001f9d6\u200d\u2640\ufe0f",
		map[string]struct{}{
			"steamy": struct{}{},
		},
	},
	":climbing:": Code{
		"\U0001f9d7",
		map[string]struct{}{
			"bouldering": struct{}{},
		},
	},
	":climbing_man:": Code{
		"\U0001f9d7\u200d\u2642\ufe0f",
		map[string]struct{}{
			"bouldering": struct{}{},
		},
	},
	":climbing_woman:": Code{
		"\U0001f9d7\u200d\u2640\ufe0f",
		map[string]struct{}{
			"bouldering": struct{}{},
		},
	},
	":golfing:": Code{
		"\U0001f3cc\ufe0f",
		map[string]struct{}{},
	},
	":golfing_man:": Code{
		"\U0001f3cc\ufe0f\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":golfing_woman:": Code{
		"\U0001f3cc\ufe0f\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":surfer:": Code{
		"\U0001f3c4",
		map[string]struct{}{},
	},
	":surfing_man:": Code{
		"\U0001f3c4\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":surfing_woman:": Code{
		"\U0001f3c4\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":rowboat:": Code{
		"\U0001f6a3",
		map[string]struct{}{},
	},
	":rowing_man:": Code{
		"\U0001f6a3\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":rowing_woman:": Code{
		"\U0001f6a3\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":swimmer:": Code{
		"\U0001f3ca",
		map[string]struct{}{},
	},
	":swimming_man:": Code{
		"\U0001f3ca\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":swimming_woman:": Code{
		"\U0001f3ca\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":bouncing_ball_person:": Code{
		"\u26f9\ufe0f",
		map[string]struct{}{
			"basketball": struct{}{},
		},
	},
	":bouncing_ball_man:": Code{
		"\u26f9\ufe0f\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":basketball_man:": Code{
		"\u26f9\ufe0f\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":bouncing_ball_woman:": Code{
		"\u26f9\ufe0f\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":basketball_woman:": Code{
		"\u26f9\ufe0f\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":weight_lifting:": Code{
		"\U0001f3cb\ufe0f",
		map[string]struct{}{
			"gym":     struct{}{},
			"workout": struct{}{},
		},
	},
	":weight_lifting_man:": Code{
		"\U0001f3cb\ufe0f\u200d\u2642\ufe0f",
		map[string]struct{}{
			"gym":     struct{}{},
			"workout": struct{}{},
		},
	},
	":weight_lifting_woman:": Code{
		"\U0001f3cb\ufe0f\u200d\u2640\ufe0f",
		map[string]struct{}{
			"gym":     struct{}{},
			"workout": struct{}{},
		},
	},
	":bicyclist:": Code{
		"\U0001f6b4",
		map[string]struct{}{},
	},
	":biking_man:": Code{
		"\U0001f6b4\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":biking_woman:": Code{
		"\U0001f6b4\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":mountain_bicyclist:": Code{
		"\U0001f6b5",
		map[string]struct{}{},
	},
	":mountain_biking_man:": Code{
		"\U0001f6b5\u200d\u2642\ufe0f",
		map[string]struct{}{},
	},
	":mountain_biking_woman:": Code{
		"\U0001f6b5\u200d\u2640\ufe0f",
		map[string]struct{}{},
	},
	":cartwheeling:": Code{
		"\U0001f938",
		map[string]struct{}{},
	},
	":wrestling:": Code{
		"\U0001f93c",
		map[string]struct{}{},
	},
	":water_polo:": Code{
		"\U0001f93d",
		map[string]struct{}{},
	},
	":handball_person:": Code{
		"\U0001f93e",
		map[string]struct{}{},
	},
	":juggling_person:": Code{
		"\U0001f939",
		map[string]struct{}{},
	},
	":lotus_position:": Code{
		"\U0001f9d8",
		map[string]struct{}{
			"meditation": struct{}{},
		},
	},
	":lotus_position_man:": Code{
		"\U0001f9d8\u200d\u2642\ufe0f",
		map[string]struct{}{
			"meditation": struct{}{},
		},
	},
	":lotus_position_woman:": Code{
		"\U0001f9d8\u200d\u2640\ufe0f",
		map[string]struct{}{
			"meditation": struct{}{},
		},
	},
	":sleeping_bed:": Code{
		"\U0001f6cc",
		map[string]struct{}{},
	},
	":people_holding_hands:": Code{
		"\U0001f9d1\u200d\U0001f91d\u200d\U0001f9d1",
		map[string]struct{}{
			"couple": struct{}{},
			"date":   struct{}{},
		},
	},
	":couplekiss_man_woman:": Code{
		"\U0001f469\u200d\u2764\ufe0f\u200d\U0001f48b\u200d\U0001f468",
		map[string]struct{}{},
	},
	":couplekiss_man_man:": Code{
		"\U0001f468\u200d\u2764\ufe0f\u200d\U0001f48b\u200d\U0001f468",
		map[string]struct{}{},
	},
	":couplekiss_woman_woman:": Code{
		"\U0001f469\u200d\u2764\ufe0f\u200d\U0001f48b\u200d\U0001f469",
		map[string]struct{}{},
	},
	":couple_with_heart_man_man:": Code{
		"\U0001f468\u200d\u2764\ufe0f\u200d\U0001f468",
		map[string]struct{}{},
	},
	":couple_with_heart_woman_woman:": Code{
		"\U0001f469\u200d\u2764\ufe0f\u200d\U0001f469",
		map[string]struct{}{},
	},
	":family_man_woman_girl:": Code{
		"\U0001f468\u200d\U0001f469\u200d\U0001f467",
		map[string]struct{}{},
	},
	":family_man_woman_girl_boy:": Code{
		"\U0001f468\u200d\U0001f469\u200d\U0001f467\u200d\U0001f466",
		map[string]struct{}{},
	},
	":family_man_woman_boy_boy:": Code{
		"\U0001f468\u200d\U0001f469\u200d\U0001f466\u200d\U0001f466",
		map[string]struct{}{},
	},
	":family_man_woman_girl_girl:": Code{
		"\U0001f468\u200d\U0001f469\u200d\U0001f467\u200d\U0001f467",
		map[string]struct{}{},
	},
	":family_man_man_boy:": Code{
		"\U0001f468\u200d\U0001f468\u200d\U0001f466",
		map[string]struct{}{},
	},
	":family_man_man_girl:": Code{
		"\U0001f468\u200d\U0001f468\u200d\U0001f467",
		map[string]struct{}{},
	},
	":family_man_man_girl_boy:": Code{
		"\U0001f468\u200d\U0001f468\u200d\U0001f467\u200d\U0001f466",
		map[string]struct{}{},
	},
	":family_man_man_boy_boy:": Code{
		"\U0001f468\u200d\U0001f468\u200d\U0001f466\u200d\U0001f466",
		map[string]struct{}{},
	},
	":family_man_man_girl_girl:": Code{
		"\U0001f468\u200d\U0001f468\u200d\U0001f467\u200d\U0001f467",
		map[string]struct{}{},
	},
	":family_woman_woman_boy:": Code{
		"\U0001f469\u200d\U0001f469\u200d\U0001f466",
		map[string]struct{}{},
	},
	":family_woman_woman_girl:": Code{
		"\U0001f469\u200d\U0001f469\u200d\U0001f467",
		map[string]struct{}{},
	},
	":family_woman_woman_girl_boy:": Code{
		"\U0001f469\u200d\U0001f469\u200d\U0001f467\u200d\U0001f466",
		map[string]struct{}{},
	},
	":family_woman_woman_boy_boy:": Code{
		"\U0001f469\u200d\U0001f469\u200d\U0001f466\u200d\U0001f466",
		map[string]struct{}{},
	},
	":family_woman_woman_girl_girl:": Code{
		"\U0001f469\u200d\U0001f469\u200d\U0001f467\u200d\U0001f467",
		map[string]struct{}{},
	},
	":people_hugging:": Code{
		"\U0001fac2",
		map[string]struct{}{},
	},
	":orangutan:": Code{
		"\U0001f9a7",
		map[string]struct{}{},
	},
	":guide_dog:": Code{
		"\U0001f9ae",
		map[string]struct{}{},
	},
	":service_dog:": Code{
		"\U0001f415\u200d\U0001f9ba",
		map[string]struct{}{},
	},
	":fox_face:": Code{
		"\U0001f98a",
		map[string]struct{}{},
	},
	":raccoon:": Code{
		"\U0001f99d",
		map[string]struct{}{},
	},
	":black_cat:": Code{
		"\U0001f408\u200d\u2b1b",
		map[string]struct{}{},
	},
	":lion:": Code{
		"\U0001f981",
		map[string]struct{}{},
	},
	":bison:": Code{
		"\U0001f9ac",
		map[string]struct{}{},
	},
	":llama:": Code{
		"\U0001f999",
		map[string]struct{}{},
	},
	":mammoth:": Code{
		"\U0001f9a3",
		map[string]struct{}{},
	},
	":rhinoceros:": Code{
		"\U0001f98f",
		map[string]struct{}{},
	},
	":hippopotamus:": Code{
		"\U0001f99b",
		map[string]struct{}{},
	},
	":beaver:": Code{
		"\U0001f9ab",
		map[string]struct{}{},
	},
	":polar_bear:": Code{
		"\U0001f43b\u200d\u2744\ufe0f",
		map[string]struct{}{},
	},
	":sloth:": Code{
		"\U0001f9a5",
		map[string]struct{}{},
	},
	":otter:": Code{
		"\U0001f9a6",
		map[string]struct{}{},
	},
	":skunk:": Code{
		"\U0001f9a8",
		map[string]struct{}{},
	},
	":kangaroo:": Code{
		"\U0001f998",
		map[string]struct{}{},
	},
	":badger:": Code{
		"\U0001f9a1",
		map[string]struct{}{},
	},
	":paw_prints:": Code{
		"\U0001f43e",
		map[string]struct{}{},
	},
	":swan:": Code{
		"\U0001f9a2",
		map[string]struct{}{},
	},
	":dodo:": Code{
		"\U0001f9a4",
		map[string]struct{}{},
	},
	":feather:": Code{
		"\U0001fab6",
		map[string]struct{}{},
	},
	":flamingo:": Code{
		"\U0001f9a9",
		map[string]struct{}{},
	},
	":peacock:": Code{
		"\U0001f99a",
		map[string]struct{}{},
	},
	":parrot:": Code{
		"\U0001f99c",
		map[string]struct{}{},
	},
	":t-rex:": Code{
		"\U0001f996",
		map[string]struct{}{
			"dinosaur": struct{}{},
		},
	},
	":flipper:": Code{
		"\U0001f42c",
		map[string]struct{}{},
	},
	":seal:": Code{
		"\U0001f9ad",
		map[string]struct{}{},
	},
	":honeybee:": Code{
		"\U0001f41d",
		map[string]struct{}{},
	},
	":lady_beetle:": Code{
		"\U0001f41e",
		map[string]struct{}{
			"bug": struct{}{},
		},
	},
	":cockroach:": Code{
		"\U0001fab3",
		map[string]struct{}{},
	},
	":mosquito:": Code{
		"\U0001f99f",
		map[string]struct{}{},
	},
	":fly:": Code{
		"\U0001fab0",
		map[string]struct{}{},
	},
	":worm:": Code{
		"\U0001fab1",
		map[string]struct{}{},
	},
	":microbe:": Code{
		"\U0001f9a0",
		map[string]struct{}{
			"germ": struct{}{},
		},
	},
	":wilted_flower:": Code{
		"\U0001f940",
		map[string]struct{}{},
	},
	":potted_plant:": Code{
		"\U0001fab4",
		map[string]struct{}{},
	},
	":orange:": Code{
		"\U0001f34a",
		map[string]struct{}{},
	},
	":mandarin:": Code{
		"\U0001f34a",
		map[string]struct{}{},
	},
	":mango:": Code{
		"\U0001f96d",
		map[string]struct{}{},
	},
	":blueberries:": Code{
		"\U0001fad0",
		map[string]struct{}{},
	},
	":kiwi_fruit:": Code{
		"\U0001f95d",
		map[string]struct{}{},
	},
	":olive:": Code{
		"\U0001fad2",
		map[string]struct{}{},
	},
	":bell_pepper:": Code{
		"\U0001fad1",
		map[string]struct{}{},
	},
	":leafy_green:": Code{
		"\U0001f96c",
		map[string]struct{}{},
	},
	":garlic:": Code{
		"\U0001f9c4",
		map[string]struct{}{},
	},
	":onion:": Code{
		"\U0001f9c5",
		map[string]struct{}{},
	},
	":baguette_bread:": Code{
		"\U0001f956",
		map[string]struct{}{},
	},
	":flatbread:": Code{
		"\U0001fad3",
		map[string]struct{}{},
	},
	":bagel:": Code{
		"\U0001f96f",
		map[string]struct{}{},
	},
	":waffle:": Code{
		"\U0001f9c7",
		map[string]struct{}{},
	},
	":tamale:": Code{
		"\U0001fad4",
		map[string]struct{}{},
	},
	":falafel:": Code{
		"\U0001f9c6",
		map[string]struct{}{},
	},
	":fried_egg:": Code{
		"\U0001f373",
		map[string]struct{}{
			"breakfast": struct{}{},
		},
	},
	":fondue:": Code{
		"\U0001fad5",
		map[string]struct{}{},
	},
	":green_salad:": Code{
		"\U0001f957",
		map[string]struct{}{},
	},
	":butter:": Code{
		"\U0001f9c8",
		map[string]struct{}{},
	},
	":salt:": Code{
		"\U0001f9c2",
		map[string]struct{}{},
	},
	":moon_cake:": Code{
		"\U0001f96e",
		map[string]struct{}{},
	},
	":lobster:": Code{
		"\U0001f99e",
		map[string]struct{}{},
	},
	":oyster:": Code{
		"\U0001f9aa",
		map[string]struct{}{},
	},
	":cupcake:": Code{
		"\U0001f9c1",
		map[string]struct{}{},
	},
	":milk_glass:": Code{
		"\U0001f95b",
		map[string]struct{}{},
	},
	":teapot:": Code{
		"\U0001fad6",
		map[string]struct{}{},
	},
	":clinking_glasses:": Code{
		"\U0001f942",
		map[string]struct{}{
			"cheers": struct{}{},
			"toast":  struct{}{},
		},
	},
	":bubble_tea:": Code{
		"\U0001f9cb",
		map[string]struct{}{},
	},
	":beverage_box:": Code{
		"\U0001f9c3",
		map[string]struct{}{},
	},
	":mate:": Code{
		"\U0001f9c9",
		map[string]struct{}{},
	},
	":ice_cube:": Code{
		"\U0001f9ca",
		map[string]struct{}{},
	},
	":plate_with_cutlery:": Code{
		"\U0001f37d\ufe0f",
		map[string]struct{}{
			"dining": struct{}{},
			"dinner": struct{}{},
		},
	},
	":hocho:": Code{
		"\U0001f52a",
		map[string]struct{}{
			"cut":  struct{}{},
			"chop": struct{}{},
		},
	},
	":world_map:": Code{
		"\U0001f5fa\ufe0f",
		map[string]struct{}{
			"travel": struct{}{},
		},
	},
	":compass:": Code{
		"\U0001f9ed",
		map[string]struct{}{},
	},
	":desert_island:": Code{
		"\U0001f3dd\ufe0f",
		map[string]struct{}{},
	},
	":national_park:": Code{
		"\U0001f3de\ufe0f",
		map[string]struct{}{},
	},
	":building_construction:": Code{
		"\U0001f3d7\ufe0f",
		map[string]struct{}{},
	},
	":bricks:": Code{
		"\U0001f9f1",
		map[string]struct{}{},
	},
	":rock:": Code{
		"\U0001faa8",
		map[string]struct{}{},
	},
	":wood:": Code{
		"\U0001fab5",
		map[string]struct{}{},
	},
	":hut:": Code{
		"\U0001f6d6",
		map[string]struct{}{},
	},
	":houses:": Code{
		"\U0001f3d8\ufe0f",
		map[string]struct{}{},
	},
	":derelict_house:": Code{
		"\U0001f3da\ufe0f",
		map[string]struct{}{},
	},
	":hindu_temple:": Code{
		"\U0001f6d5",
		map[string]struct{}{},
	},
	":city_sunrise:": Code{
		"\U0001f307",
		map[string]struct{}{},
	},
	":car:": Code{
		"\U0001f697",
		map[string]struct{}{},
	},
	":pickup_truck:": Code{
		"\U0001f6fb",
		map[string]struct{}{},
	},
	":racing_car:": Code{
		"\U0001f3ce\ufe0f",
		map[string]struct{}{},
	},
	":manual_wheelchair:": Code{
		"\U0001f9bd",
		map[string]struct{}{},
	},
	":motorized_wheelchair:": Code{
		"\U0001f9bc",
		map[string]struct{}{},
	},
	":auto_rickshaw:": Code{
		"\U0001f6fa",
		map[string]struct{}{},
	},
	":kick_scooter:": Code{
		"\U0001f6f4",
		map[string]struct{}{},
	},
	":skateboard:": Code{
		"\U0001f6f9",
		map[string]struct{}{},
	},
	":roller_skate:": Code{
		"\U0001f6fc",
		map[string]struct{}{},
	},
	":oil_drum:": Code{
		"\U0001f6e2\ufe0f",
		map[string]struct{}{},
	},
	":stop_sign:": Code{
		"\U0001f6d1",
		map[string]struct{}{},
	},
	":boat:": Code{
		"\u26f5",
		map[string]struct{}{},
	},
	":passenger_ship:": Code{
		"\U0001f6f3\ufe0f",
		map[string]struct{}{
			"cruise": struct{}{},
		},
	},
	":motor_boat:": Code{
		"\U0001f6e5\ufe0f",
		map[string]struct{}{},
	},
	":small_airplane:": Code{
		"\U0001f6e9\ufe0f",
		map[string]struct{}{
			"flight": struct{}{},
		},
	},
	":flight_departure:": Code{
		"\U0001f6eb",
		map[string]struct{}{},
	},
	":flight_arrival:": Code{
		"\U0001f6ec",
		map[string]struct{}{},
	},
	":parachute:": Code{
		"\U0001fa82",
		map[string]struct{}{},
	},
	":artificial_satellite:": Code{
		"\U0001f6f0\ufe0f",
		map[string]struct{}{
			"orbit": struct{}{},
			"space": struct{}{},
		},
	},
	":bellhop_bell:": Code{
		"\U0001f6ce\ufe0f",
		map[string]struct{}{},
	},
	":luggage:": Code{
		"\U0001f9f3",
		map[string]struct{}{},
	},
	":timer_clock:": Code{
		"\u23f2\ufe0f",
		map[string]struct{}{},
	},
	":mantelpiece_clock:": Code{
		"\U0001f570\ufe0f",
		map[string]struct{}{},
	},
	":moon:": Code{
		"\U0001f314",
		map[string]struct{}{},
	},
	":ringed_planet:": Code{
		"\U0001fa90",
		map[string]struct{}{},
	},
	":cloud_with_lightning_and_rain:": Code{
		"\u26c8\ufe0f",
		map[string]struct{}{},
	},
	":sun_behind_small_cloud:": Code{
		"\U0001f324\ufe0f",
		map[string]struct{}{},
	},
	":sun_behind_large_cloud:": Code{
		"\U0001f325\ufe0f",
		map[string]struct{}{},
	},
	":sun_behind_rain_cloud:": Code{
		"\U0001f326\ufe0f",
		map[string]struct{}{},
	},
	":cloud_with_rain:": Code{
		"\U0001f327\ufe0f",
		map[string]struct{}{},
	},
	":cloud_with_snow:": Code{
		"\U0001f328\ufe0f",
		map[string]struct{}{},
	},
	":cloud_with_lightning:": Code{
		"\U0001f329\ufe0f",
		map[string]struct{}{},
	},
	":tornado:": Code{
		"\U0001f32a\ufe0f",
		map[string]struct{}{},
	},
	":wind_face:": Code{
		"\U0001f32c\ufe0f",
		map[string]struct{}{},
	},
	":open_umbrella:": Code{
		"\u2602\ufe0f",
		map[string]struct{}{},
	},
	":parasol_on_ground:": Code{
		"\u26f1\ufe0f",
		map[string]struct{}{
			"beach_umbrella": struct{}{},
		},
	},
	":snowman_with_snow:": Code{
		"\u2603\ufe0f",
		map[string]struct{}{
			"winter":    struct{}{},
			"christmas": struct{}{},
		},
	},
	":firecracker:": Code{
		"\U0001f9e8",
		map[string]struct{}{},
	},
	":red_envelope:": Code{
		"\U0001f9e7",
		map[string]struct{}{},
	},
	":medal_military:": Code{
		"\U0001f396\ufe0f",
		map[string]struct{}{},
	},
	":medal_sports:": Code{
		"\U0001f3c5",
		map[string]struct{}{
			"gold":   struct{}{},
			"winner": struct{}{},
		},
	},
	":1st_place_medal:": Code{
		"\U0001f947",
		map[string]struct{}{
			"gold": struct{}{},
		},
	},
	":2nd_place_medal:": Code{
		"\U0001f948",
		map[string]struct{}{
			"silver": struct{}{},
		},
	},
	":3rd_place_medal:": Code{
		"\U0001f949",
		map[string]struct{}{
			"bronze": struct{}{},
		},
	},
	":softball:": Code{
		"\U0001f94e",
		map[string]struct{}{},
	},
	":flying_disc:": Code{
		"\U0001f94f",
		map[string]struct{}{},
	},
	":ice_hockey:": Code{
		"\U0001f3d2",
		map[string]struct{}{},
	},
	":lacrosse:": Code{
		"\U0001f94d",
		map[string]struct{}{},
	},
	":goal_net:": Code{
		"\U0001f945",
		map[string]struct{}{},
	},
	":diving_mask:": Code{
		"\U0001f93f",
		map[string]struct{}{},
	},
	":yo_yo:": Code{
		"\U0001fa80",
		map[string]struct{}{},
	},
	":kite:": Code{
		"\U0001fa81",
		map[string]struct{}{},
	},
	":magic_wand:": Code{
		"\U0001fa84",
		map[string]struct{}{},
	},
	":nazar_amulet:": Code{
		"\U0001f9ff",
		map[string]struct{}{},
	},
	":jigsaw:": Code{
		"\U0001f9e9",
		map[string]struct{}{},
	},
	":teddy_bear:": Code{
		"\U0001f9f8",
		map[string]struct{}{},
	},
	":pinata:": Code{
		"\U0001fa85",
		map[string]struct{}{},
	},
	":nesting_dolls:": Code{
		"\U0001fa86",
		map[string]struct{}{},
	},
	":chess_pawn:": Code{
		"\u265f\ufe0f",
		map[string]struct{}{},
	},
	":framed_picture:": Code{
		"\U0001f5bc\ufe0f",
		map[string]struct{}{},
	},
	":thread:": Code{
		"\U0001f9f5",
		map[string]struct{}{},
	},
	":sewing_needle:": Code{
		"\U0001faa1",
		map[string]struct{}{},
	},
	":yarn:": Code{
		"\U0001f9f6",
		map[string]struct{}{},
	},
	":knot:": Code{
		"\U0001faa2",
		map[string]struct{}{},
	},
	":goggles:": Code{
		"\U0001f97d",
		map[string]struct{}{},
	},
	":lab_coat:": Code{
		"\U0001f97c",
		map[string]struct{}{},
	},
	":safety_vest:": Code{
		"\U0001f9ba",
		map[string]struct{}{},
	},
	":tshirt:": Code{
		"\U0001f455",
		map[string]struct{}{},
	},
	":sari:": Code{
		"\U0001f97b",
		map[string]struct{}{},
	},
	":one_piece_swimsuit:": Code{
		"\U0001fa71",
		map[string]struct{}{},
	},
	":swim_brief:": Code{
		"\U0001fa72",
		map[string]struct{}{},
	},
	":shorts:": Code{
		"\U0001fa73",
		map[string]struct{}{},
	},
	":shopping:": Code{
		"\U0001f6cd\ufe0f",
		map[string]struct{}{
			"bags": struct{}{},
		},
	},
	":thong_sandal:": Code{
		"\U0001fa74",
		map[string]struct{}{},
	},
	":shoe:": Code{
		"\U0001f45e",
		map[string]struct{}{},
	},
	":hiking_boot:": Code{
		"\U0001f97e",
		map[string]struct{}{},
	},
	":flat_shoe:": Code{
		"\U0001f97f",
		map[string]struct{}{},
	},
	":ballet_shoes:": Code{
		"\U0001fa70",
		map[string]struct{}{},
	},
	":military_helmet:": Code{
		"\U0001fa96",
		map[string]struct{}{},
	},
	":rescue_worker_helmet:": Code{
		"\u26d1\ufe0f",
		map[string]struct{}{},
	},
	":studio_microphone:": Code{
		"\U0001f399\ufe0f",
		map[string]struct{}{
			"podcast": struct{}{},
		},
	},
	":accordion:": Code{
		"\U0001fa97",
		map[string]struct{}{},
	},
	":banjo:": Code{
		"\U0001fa95",
		map[string]struct{}{},
	},
	":long_drum:": Code{
		"\U0001fa98",
		map[string]struct{}{},
	},
	":phone:": Code{
		"\u260e\ufe0f",
		map[string]struct{}{},
	},
	":desktop_computer:": Code{
		"\U0001f5a5\ufe0f",
		map[string]struct{}{},
	},
	":computer_mouse:": Code{
		"\U0001f5b1\ufe0f",
		map[string]struct{}{},
	},
	":abacus:": Code{
		"\U0001f9ee",
		map[string]struct{}{},
	},
	":film_strip:": Code{
		"\U0001f39e\ufe0f",
		map[string]struct{}{},
	},
	":film_projector:": Code{
		"\U0001f4fd\ufe0f",
		map[string]struct{}{},
	},
	":camera_flash:": Code{
		"\U0001f4f8",
		map[string]struct{}{
			"photo": struct{}{},
		},
	},
	":lantern:": Code{
		"\U0001f3ee",
		map[string]struct{}{},
	},
	":diya_lamp:": Code{
		"\U0001fa94",
		map[string]struct{}{},
	},
	":open_book:": Code{
		"\U0001f4d6",
		map[string]struct{}{},
	},
	":newspaper_roll:": Code{
		"\U0001f5de\ufe0f",
		map[string]struct{}{
			"press": struct{}{},
		},
	},
	":coin:": Code{
		"\U0001fa99",
		map[string]struct{}{},
	},
	":receipt:": Code{
		"\U0001f9fe",
		map[string]struct{}{},
	},
	":email:": Code{
		"\U0001f4e7",
		map[string]struct{}{},
	},
	":fountain_pen:": Code{
		"\U0001f58b\ufe0f",
		map[string]struct{}{},
	},
	":pen:": Code{
		"\U0001f58a\ufe0f",
		map[string]struct{}{},
	},
	":memo:": Code{
		"\U0001f4dd",
		map[string]struct{}{
			"document": struct{}{},
			"note":     struct{}{},
		},
	},
	":card_index_dividers:": Code{
		"\U0001f5c2\ufe0f",
		map[string]struct{}{},
	},
	":spiral_notepad:": Code{
		"\U0001f5d2\ufe0f",
		map[string]struct{}{},
	},
	":spiral_calendar:": Code{
		"\U0001f5d3\ufe0f",
		map[string]struct{}{},
	},
	":card_file_box:": Code{
		"\U0001f5c3\ufe0f",
		map[string]struct{}{},
	},
	":old_key:": Code{
		"\U0001f5dd\ufe0f",
		map[string]struct{}{},
	},
	":axe:": Code{
		"\U0001fa93",
		map[string]struct{}{},
	},
	":hammer_and_pick:": Code{
		"\u2692\ufe0f",
		map[string]struct{}{},
	},
	":hammer_and_wrench:": Code{
		"\U0001f6e0\ufe0f",
		map[string]struct{}{},
	},
	":boomerang:": Code{
		"\U0001fa83",
		map[string]struct{}{},
	},
	":carpentry_saw:": Code{
		"\U0001fa9a",
		map[string]struct{}{},
	},
	":screwdriver:": Code{
		"\U0001fa9b",
		map[string]struct{}{},
	},
	":clamp:": Code{
		"\U0001f5dc\ufe0f",
		map[string]struct{}{},
	},
	":balance_scale:": Code{
		"\u2696\ufe0f",
		map[string]struct{}{},
	},
	":probing_cane:": Code{
		"\U0001f9af",
		map[string]struct{}{},
	},
	":hook:": Code{
		"\U0001fa9d",
		map[string]struct{}{},
	},
	":toolbox:": Code{
		"\U0001f9f0",
		map[string]struct{}{},
	},
	":magnet:": Code{
		"\U0001f9f2",
		map[string]struct{}{},
	},
	":ladder:": Code{
		"\U0001fa9c",
		map[string]struct{}{},
	},
	":test_tube:": Code{
		"\U0001f9ea",
		map[string]struct{}{},
	},
	":petri_dish:": Code{
		"\U0001f9eb",
		map[string]struct{}{},
	},
	":dna:": Code{
		"\U0001f9ec",
		map[string]struct{}{},
	},
	":drop_of_blood:": Code{
		"\U0001fa78",
		map[string]struct{}{},
	},
	":adhesive_bandage:": Code{
		"\U0001fa79",
		map[string]struct{}{},
	},
	":stethoscope:": Code{
		"\U0001fa7a",
		map[string]struct{}{},
	},
	":elevator:": Code{
		"\U0001f6d7",
		map[string]struct{}{},
	},
	":mirror:": Code{
		"\U0001fa9e",
		map[string]struct{}{},
	},
	":window:": Code{
		"\U0001fa9f",
		map[string]struct{}{},
	},
	":couch_and_lamp:": Code{
		"\U0001f6cb\ufe0f",
		map[string]struct{}{},
	},
	":chair:": Code{
		"\U0001fa91",
		map[string]struct{}{},
	},
	":plunger:": Code{
		"\U0001faa0",
		map[string]struct{}{},
	},
	":mouse_trap:": Code{
		"\U0001faa4",
		map[string]struct{}{},
	},
	":razor:": Code{
		"\U0001fa92",
		map[string]struct{}{},
	},
	":lotion_bottle:": Code{
		"\U0001f9f4",
		map[string]struct{}{},
	},
	":safety_pin:": Code{
		"\U0001f9f7",
		map[string]struct{}{},
	},
	":broom:": Code{
		"\U0001f9f9",
		map[string]struct{}{},
	},
	":basket:": Code{
		"\U0001f9fa",
		map[string]struct{}{},
	},
	":roll_of_paper:": Code{
		"\U0001f9fb",
		map[string]struct{}{
			"toilet": struct{}{},
		},
	},
	":bucket:": Code{
		"\U0001faa3",
		map[string]struct{}{},
	},
	":soap:": Code{
		"\U0001f9fc",
		map[string]struct{}{},
	},
	":toothbrush:": Code{
		"\U0001faa5",
		map[string]struct{}{},
	},
	":sponge:": Code{
		"\U0001f9fd",
		map[string]struct{}{},
	},
	":fire_extinguisher:": Code{
		"\U0001f9ef",
		map[string]struct{}{},
	},
	":headstone:": Code{
		"\U0001faa6",
		map[string]struct{}{},
	},
	":funeral_urn:": Code{
		"\u26b1\ufe0f",
		map[string]struct{}{},
	},
	":placard:": Code{
		"\U0001faa7",
		map[string]struct{}{},
	},
	":atom_symbol:": Code{
		"\u269b\ufe0f",
		map[string]struct{}{},
	},
	":om:": Code{
		"\U0001f549\ufe0f",
		map[string]struct{}{},
	},
	":latin_cross:": Code{
		"\u271d\ufe0f",
		map[string]struct{}{},
	},
	":peace_symbol:": Code{
		"\u262e\ufe0f",
		map[string]struct{}{},
	},
	":next_track_button:": Code{
		"\u23ed\ufe0f",
		map[string]struct{}{},
	},
	":play_or_pause_button:": Code{
		"\u23ef\ufe0f",
		map[string]struct{}{},
	},
	":previous_track_button:": Code{
		"\u23ee\ufe0f",
		map[string]struct{}{},
	},
	":eject_button:": Code{
		"\u23cf\ufe0f",
		map[string]struct{}{},
	},
	":transgender_symbol:": Code{
		"\u26a7\ufe0f",
		map[string]struct{}{},
	},
	":infinity:": Code{
		"\u267e\ufe0f",
		map[string]struct{}{},
	},
	":heavy_exclamation_mark:": Code{
		"\u2757",
		map[string]struct{}{
			"bang": struct{}{},
		},
	},
	":fleur_de_lis:": Code{
		"\u269c\ufe0f",
		map[string]struct{}{},
	},
	":orange_circle:": Code{
		"\U0001f7e0",
		map[string]struct{}{},
	},
	":yellow_circle:": Code{
		"\U0001f7e1",
		map[string]struct{}{},
	},
	":green_circle:": Code{
		"\U0001f7e2",
		map[string]struct{}{},
	},
	":large_blue_circle:": Code{
		"\U0001f535",
		map[string]struct{}{},
	},
	":purple_circle:": Code{
		"\U0001f7e3",
		map[string]struct{}{},
	},
	":brown_circle:": Code{
		"\U0001f7e4",
		map[string]struct{}{},
	},
	":red_square:": Code{
		"\U0001f7e5",
		map[string]struct{}{},
	},
	":orange_square:": Code{
		"\U0001f7e7",
		map[string]struct{}{},
	},
	":yellow_square:": Code{
		"\U0001f7e8",
		map[string]struct{}{},
	},
	":green_square:": Code{
		"\U0001f7e9",
		map[string]struct{}{},
	},
	":blue_square:": Code{
		"\U0001f7e6",
		map[string]struct{}{},
	},
	":purple_square:": Code{
		"\U0001f7ea",
		map[string]struct{}{},
	},
	":brown_square:": Code{
		"\U0001f7eb",
		map[string]struct{}{},
	},
	":black_flag:": Code{
		"\U0001f3f4",
		map[string]struct{}{},
	},
	":white_flag:": Code{
		"\U0001f3f3\ufe0f",
		map[string]struct{}{},
	},
	":transgender_flag:": Code{
		"\U0001f3f3\ufe0f\u200d\u26a7\ufe0f",
		map[string]struct{}{},
	},
	":pirate_flag:": Code{
		"\U0001f3f4\u200d\u2620\ufe0f",
		map[string]struct{}{},
	},
	":ascension_island:": Code{
		"\U0001f1e6\U0001f1e8",
		map[string]struct{}{},
	},
	":andorra:": Code{
		"\U0001f1e6\U0001f1e9",
		map[string]struct{}{},
	},
	":united_arab_emirates:": Code{
		"\U0001f1e6\U0001f1ea",
		map[string]struct{}{},
	},
	":afghanistan:": Code{
		"\U0001f1e6\U0001f1eb",
		map[string]struct{}{},
	},
	":antigua_barbuda:": Code{
		"\U0001f1e6\U0001f1ec",
		map[string]struct{}{},
	},
	":anguilla:": Code{
		"\U0001f1e6\U0001f1ee",
		map[string]struct{}{},
	},
	":albania:": Code{
		"\U0001f1e6\U0001f1f1",
		map[string]struct{}{},
	},
	":armenia:": Code{
		"\U0001f1e6\U0001f1f2",
		map[string]struct{}{},
	},
	":angola:": Code{
		"\U0001f1e6\U0001f1f4",
		map[string]struct{}{},
	},
	":antarctica:": Code{
		"\U0001f1e6\U0001f1f6",
		map[string]struct{}{},
	},
	":argentina:": Code{
		"\U0001f1e6\U0001f1f7",
		map[string]struct{}{},
	},
	":american_samoa:": Code{
		"\U0001f1e6\U0001f1f8",
		map[string]struct{}{},
	},
	":austria:": Code{
		"\U0001f1e6\U0001f1f9",
		map[string]struct{}{},
	},
	":australia:": Code{
		"\U0001f1e6\U0001f1fa",
		map[string]struct{}{},
	},
	":aruba:": Code{
		"\U0001f1e6\U0001f1fc",
		map[string]struct{}{},
	},
	":aland_islands:": Code{
		"\U0001f1e6\U0001f1fd",
		map[string]struct{}{},
	},
	":azerbaijan:": Code{
		"\U0001f1e6\U0001f1ff",
		map[string]struct{}{},
	},
	":bosnia_herzegovina:": Code{
		"\U0001f1e7\U0001f1e6",
		map[string]struct{}{},
	},
	":barbados:": Code{
		"\U0001f1e7\U0001f1e7",
		map[string]struct{}{},
	},
	":bangladesh:": Code{
		"\U0001f1e7\U0001f1e9",
		map[string]struct{}{},
	},
	":belgium:": Code{
		"\U0001f1e7\U0001f1ea",
		map[string]struct{}{},
	},
	":burkina_faso:": Code{
		"\U0001f1e7\U0001f1eb",
		map[string]struct{}{},
	},
	":bulgaria:": Code{
		"\U0001f1e7\U0001f1ec",
		map[string]struct{}{},
	},
	":bahrain:": Code{
		"\U0001f1e7\U0001f1ed",
		map[string]struct{}{},
	},
	":burundi:": Code{
		"\U0001f1e7\U0001f1ee",
		map[string]struct{}{},
	},
	":benin:": Code{
		"\U0001f1e7\U0001f1ef",
		map[string]struct{}{},
	},
	":st_barthelemy:": Code{
		"\U0001f1e7\U0001f1f1",
		map[string]struct{}{},
	},
	":bermuda:": Code{
		"\U0001f1e7\U0001f1f2",
		map[string]struct{}{},
	},
	":brunei:": Code{
		"\U0001f1e7\U0001f1f3",
		map[string]struct{}{},
	},
	":bolivia:": Code{
		"\U0001f1e7\U0001f1f4",
		map[string]struct{}{},
	},
	":caribbean_netherlands:": Code{
		"\U0001f1e7\U0001f1f6",
		map[string]struct{}{},
	},
	":brazil:": Code{
		"\U0001f1e7\U0001f1f7",
		map[string]struct{}{},
	},
	":bahamas:": Code{
		"\U0001f1e7\U0001f1f8",
		map[string]struct{}{},
	},
	":bhutan:": Code{
		"\U0001f1e7\U0001f1f9",
		map[string]struct{}{},
	},
	":bouvet_island:": Code{
		"\U0001f1e7\U0001f1fb",
		map[string]struct{}{},
	},
	":botswana:": Code{
		"\U0001f1e7\U0001f1fc",
		map[string]struct{}{},
	},
	":belarus:": Code{
		"\U0001f1e7\U0001f1fe",
		map[string]struct{}{},
	},
	":belize:": Code{
		"\U0001f1e7\U0001f1ff",
		map[string]struct{}{},
	},
	":canada:": Code{
		"\U0001f1e8\U0001f1e6",
		map[string]struct{}{},
	},
	":cocos_islands:": Code{
		"\U0001f1e8\U0001f1e8",
		map[string]struct{}{
			"keeling": struct{}{},
		},
	},
	":congo_kinshasa:": Code{
		"\U0001f1e8\U0001f1e9",
		map[string]struct{}{},
	},
	":central_african_republic:": Code{
		"\U0001f1e8\U0001f1eb",
		map[string]struct{}{},
	},
	":congo_brazzaville:": Code{
		"\U0001f1e8\U0001f1ec",
		map[string]struct{}{},
	},
	":switzerland:": Code{
		"\U0001f1e8\U0001f1ed",
		map[string]struct{}{},
	},
	":cote_divoire:": Code{
		"\U0001f1e8\U0001f1ee",
		map[string]struct{}{
			"ivory": struct{}{},
		},
	},
	":cook_islands:": Code{
		"\U0001f1e8\U0001f1f0",
		map[string]struct{}{},
	},
	":chile:": Code{
		"\U0001f1e8\U0001f1f1",
		map[string]struct{}{},
	},
	":cameroon:": Code{
		"\U0001f1e8\U0001f1f2",
		map[string]struct{}{},
	},
	":cn:": Code{
		"\U0001f1e8\U0001f1f3",
		map[string]struct{}{
			"china": struct{}{},
		},
	},
	":colombia:": Code{
		"\U0001f1e8\U0001f1f4",
		map[string]struct{}{},
	},
	":clipperton_island:": Code{
		"\U0001f1e8\U0001f1f5",
		map[string]struct{}{},
	},
	":costa_rica:": Code{
		"\U0001f1e8\U0001f1f7",
		map[string]struct{}{},
	},
	":cuba:": Code{
		"\U0001f1e8\U0001f1fa",
		map[string]struct{}{},
	},
	":cape_verde:": Code{
		"\U0001f1e8\U0001f1fb",
		map[string]struct{}{},
	},
	":curacao:": Code{
		"\U0001f1e8\U0001f1fc",
		map[string]struct{}{},
	},
	":christmas_island:": Code{
		"\U0001f1e8\U0001f1fd",
		map[string]struct{}{},
	},
	":cyprus:": Code{
		"\U0001f1e8\U0001f1fe",
		map[string]struct{}{},
	},
	":czech_republic:": Code{
		"\U0001f1e8\U0001f1ff",
		map[string]struct{}{},
	},
	":de:": Code{
		"\U0001f1e9\U0001f1ea",
		map[string]struct{}{
			"flag":    struct{}{},
			"germany": struct{}{},
		},
	},
	":diego_garcia:": Code{
		"\U0001f1e9\U0001f1ec",
		map[string]struct{}{},
	},
	":djibouti:": Code{
		"\U0001f1e9\U0001f1ef",
		map[string]struct{}{},
	},
	":denmark:": Code{
		"\U0001f1e9\U0001f1f0",
		map[string]struct{}{},
	},
	":dominica:": Code{
		"\U0001f1e9\U0001f1f2",
		map[string]struct{}{},
	},
	":dominican_republic:": Code{
		"\U0001f1e9\U0001f1f4",
		map[string]struct{}{},
	},
	":algeria:": Code{
		"\U0001f1e9\U0001f1ff",
		map[string]struct{}{},
	},
	":ceuta_melilla:": Code{
		"\U0001f1ea\U0001f1e6",
		map[string]struct{}{},
	},
	":ecuador:": Code{
		"\U0001f1ea\U0001f1e8",
		map[string]struct{}{},
	},
	":estonia:": Code{
		"\U0001f1ea\U0001f1ea",
		map[string]struct{}{},
	},
	":egypt:": Code{
		"\U0001f1ea\U0001f1ec",
		map[string]struct{}{},
	},
	":western_sahara:": Code{
		"\U0001f1ea\U0001f1ed",
		map[string]struct{}{},
	},
	":eritrea:": Code{
		"\U0001f1ea\U0001f1f7",
		map[string]struct{}{},
	},
	":es:": Code{
		"\U0001f1ea\U0001f1f8",
		map[string]struct{}{
			"spain": struct{}{},
		},
	},
	":ethiopia:": Code{
		"\U0001f1ea\U0001f1f9",
		map[string]struct{}{},
	},
	":eu:": Code{
		"\U0001f1ea\U0001f1fa",
		map[string]struct{}{},
	},
	":european_union:": Code{
		"\U0001f1ea\U0001f1fa",
		map[string]struct{}{},
	},
	":finland:": Code{
		"\U0001f1eb\U0001f1ee",
		map[string]struct{}{},
	},
	":fiji:": Code{
		"\U0001f1eb\U0001f1ef",
		map[string]struct{}{},
	},
	":falkland_islands:": Code{
		"\U0001f1eb\U0001f1f0",
		map[string]struct{}{},
	},
	":micronesia:": Code{
		"\U0001f1eb\U0001f1f2",
		map[string]struct{}{},
	},
	":faroe_islands:": Code{
		"\U0001f1eb\U0001f1f4",
		map[string]struct{}{},
	},
	":fr:": Code{
		"\U0001f1eb\U0001f1f7",
		map[string]struct{}{
			"france": struct{}{},
			"french": struct{}{},
		},
	},
	":gabon:": Code{
		"\U0001f1ec\U0001f1e6",
		map[string]struct{}{},
	},
	":gb:": Code{
		"\U0001f1ec\U0001f1e7",
		map[string]struct{}{
			"flag":    struct{}{},
			"british": struct{}{},
		},
	},
	":uk:": Code{
		"\U0001f1ec\U0001f1e7",
		map[string]struct{}{
			"flag":    struct{}{},
			"british": struct{}{},
		},
	},
	":grenada:": Code{
		"\U0001f1ec\U0001f1e9",
		map[string]struct{}{},
	},
	":georgia:": Code{
		"\U0001f1ec\U0001f1ea",
		map[string]struct{}{},
	},
	":french_guiana:": Code{
		"\U0001f1ec\U0001f1eb",
		map[string]struct{}{},
	},
	":guernsey:": Code{
		"\U0001f1ec\U0001f1ec",
		map[string]struct{}{},
	},
	":ghana:": Code{
		"\U0001f1ec\U0001f1ed",
		map[string]struct{}{},
	},
	":gibraltar:": Code{
		"\U0001f1ec\U0001f1ee",
		map[string]struct{}{},
	},
	":greenland:": Code{
		"\U0001f1ec\U0001f1f1",
		map[string]struct{}{},
	},
	":gambia:": Code{
		"\U0001f1ec\U0001f1f2",
		map[string]struct{}{},
	},
	":guinea:": Code{
		"\U0001f1ec\U0001f1f3",
		map[string]struct{}{},
	},
	":guadeloupe:": Code{
		"\U0001f1ec\U0001f1f5",
		map[string]struct{}{},
	},
	":equatorial_guinea:": Code{
		"\U0001f1ec\U0001f1f6",
		map[string]struct{}{},
	},
	":greece:": Code{
		"\U0001f1ec\U0001f1f7",
		map[string]struct{}{},
	},
	":south_georgia_south_sandwich_islands:": Code{
		"\U0001f1ec\U0001f1f8",
		map[string]struct{}{},
	},
	":guatemala:": Code{
		"\U0001f1ec\U0001f1f9",
		map[string]struct{}{},
	},
	":guam:": Code{
		"\U0001f1ec\U0001f1fa",
		map[string]struct{}{},
	},
	":guinea_bissau:": Code{
		"\U0001f1ec\U0001f1fc",
		map[string]struct{}{},
	},
	":guyana:": Code{
		"\U0001f1ec\U0001f1fe",
		map[string]struct{}{},
	},
	":hong_kong:": Code{
		"\U0001f1ed\U0001f1f0",
		map[string]struct{}{},
	},
	":heard_mcdonald_islands:": Code{
		"\U0001f1ed\U0001f1f2",
		map[string]struct{}{},
	},
	":honduras:": Code{
		"\U0001f1ed\U0001f1f3",
		map[string]struct{}{},
	},
	":croatia:": Code{
		"\U0001f1ed\U0001f1f7",
		map[string]struct{}{},
	},
	":haiti:": Code{
		"\U0001f1ed\U0001f1f9",
		map[string]struct{}{},
	},
	":hungary:": Code{
		"\U0001f1ed\U0001f1fa",
		map[string]struct{}{},
	},
	":canary_islands:": Code{
		"\U0001f1ee\U0001f1e8",
		map[string]struct{}{},
	},
	":indonesia:": Code{
		"\U0001f1ee\U0001f1e9",
		map[string]struct{}{},
	},
	":ireland:": Code{
		"\U0001f1ee\U0001f1ea",
		map[string]struct{}{},
	},
	":israel:": Code{
		"\U0001f1ee\U0001f1f1",
		map[string]struct{}{},
	},
	":isle_of_man:": Code{
		"\U0001f1ee\U0001f1f2",
		map[string]struct{}{},
	},
	":india:": Code{
		"\U0001f1ee\U0001f1f3",
		map[string]struct{}{},
	},
	":british_indian_ocean_territory:": Code{
		"\U0001f1ee\U0001f1f4",
		map[string]struct{}{},
	},
	":iraq:": Code{
		"\U0001f1ee\U0001f1f6",
		map[string]struct{}{},
	},
	":iran:": Code{
		"\U0001f1ee\U0001f1f7",
		map[string]struct{}{},
	},
	":iceland:": Code{
		"\U0001f1ee\U0001f1f8",
		map[string]struct{}{},
	},
	":it:": Code{
		"\U0001f1ee\U0001f1f9",
		map[string]struct{}{
			"italy": struct{}{},
		},
	},
	":jersey:": Code{
		"\U0001f1ef\U0001f1ea",
		map[string]struct{}{},
	},
	":jamaica:": Code{
		"\U0001f1ef\U0001f1f2",
		map[string]struct{}{},
	},
	":jordan:": Code{
		"\U0001f1ef\U0001f1f4",
		map[string]struct{}{},
	},
	":jp:": Code{
		"\U0001f1ef\U0001f1f5",
		map[string]struct{}{
			"japan": struct{}{},
		},
	},
	":kenya:": Code{
		"\U0001f1f0\U0001f1ea",
		map[string]struct{}{},
	},
	":kyrgyzstan:": Code{
		"\U0001f1f0\U0001f1ec",
		map[string]struct{}{},
	},
	":cambodia:": Code{
		"\U0001f1f0\U0001f1ed",
		map[string]struct{}{},
	},
	":kiribati:": Code{
		"\U0001f1f0\U0001f1ee",
		map[string]struct{}{},
	},
	":comoros:": Code{
		"\U0001f1f0\U0001f1f2",
		map[string]struct{}{},
	},
	":st_kitts_nevis:": Code{
		"\U0001f1f0\U0001f1f3",
		map[string]struct{}{},
	},
	":north_korea:": Code{
		"\U0001f1f0\U0001f1f5",
		map[string]struct{}{},
	},
	":kr:": Code{
		"\U0001f1f0\U0001f1f7",
		map[string]struct{}{
			"korea": struct{}{},
		},
	},
	":kuwait:": Code{
		"\U0001f1f0\U0001f1fc",
		map[string]struct{}{},
	},
	":cayman_islands:": Code{
		"\U0001f1f0\U0001f1fe",
		map[string]struct{}{},
	},
	":kazakhstan:": Code{
		"\U0001f1f0\U0001f1ff",
		map[string]struct{}{},
	},
	":laos:": Code{
		"\U0001f1f1\U0001f1e6",
		map[string]struct{}{},
	},
	":lebanon:": Code{
		"\U0001f1f1\U0001f1e7",
		map[string]struct{}{},
	},
	":st_lucia:": Code{
		"\U0001f1f1\U0001f1e8",
		map[string]struct{}{},
	},
	":liechtenstein:": Code{
		"\U0001f1f1\U0001f1ee",
		map[string]struct{}{},
	},
	":sri_lanka:": Code{
		"\U0001f1f1\U0001f1f0",
		map[string]struct{}{},
	},
	":liberia:": Code{
		"\U0001f1f1\U0001f1f7",
		map[string]struct{}{},
	},
	":lesotho:": Code{
		"\U0001f1f1\U0001f1f8",
		map[string]struct{}{},
	},
	":lithuania:": Code{
		"\U0001f1f1\U0001f1f9",
		map[string]struct{}{},
	},
	":luxembourg:": Code{
		"\U0001f1f1\U0001f1fa",
		map[string]struct{}{},
	},
	":latvia:": Code{
		"\U0001f1f1\U0001f1fb",
		map[string]struct{}{},
	},
	":libya:": Code{
		"\U0001f1f1\U0001f1fe",
		map[string]struct{}{},
	},
	":morocco:": Code{
		"\U0001f1f2\U0001f1e6",
		map[string]struct{}{},
	},
	":monaco:": Code{
		"\U0001f1f2\U0001f1e8",
		map[string]struct{}{},
	},
	":moldova:": Code{
		"\U0001f1f2\U0001f1e9",
		map[string]struct{}{},
	},
	":montenegro:": Code{
		"\U0001f1f2\U0001f1ea",
		map[string]struct{}{},
	},
	":st_martin:": Code{
		"\U0001f1f2\U0001f1eb",
		map[string]struct{}{},
	},
	":madagascar:": Code{
		"\U0001f1f2\U0001f1ec",
		map[string]struct{}{},
	},
	":marshall_islands:": Code{
		"\U0001f1f2\U0001f1ed",
		map[string]struct{}{},
	},
	":macedonia:": Code{
		"\U0001f1f2\U0001f1f0",
		map[string]struct{}{},
	},
	":mali:": Code{
		"\U0001f1f2\U0001f1f1",
		map[string]struct{}{},
	},
	":myanmar:": Code{
		"\U0001f1f2\U0001f1f2",
		map[string]struct{}{
			"burma": struct{}{},
		},
	},
	":mongolia:": Code{
		"\U0001f1f2\U0001f1f3",
		map[string]struct{}{},
	},
	":macau:": Code{
		"\U0001f1f2\U0001f1f4",
		map[string]struct{}{},
	},
	":northern_mariana_islands:": Code{
		"\U0001f1f2\U0001f1f5",
		map[string]struct{}{},
	},
	":martinique:": Code{
		"\U0001f1f2\U0001f1f6",
		map[string]struct{}{},
	},
	":mauritania:": Code{
		"\U0001f1f2\U0001f1f7",
		map[string]struct{}{},
	},
	":montserrat:": Code{
		"\U0001f1f2\U0001f1f8",
		map[string]struct{}{},
	},
	":malta:": Code{
		"\U0001f1f2\U0001f1f9",
		map[string]struct{}{},
	},
	":mauritius:": Code{
		"\U0001f1f2\U0001f1fa",
		map[string]struct{}{},
	},
	":maldives:": Code{
		"\U0001f1f2\U0001f1fb",
		map[string]struct{}{},
	},
	":malawi:": Code{
		"\U0001f1f2\U0001f1fc",
		map[string]struct{}{},
	},
	":mexico:": Code{
		"\U0001f1f2\U0001f1fd",
		map[string]struct{}{},
	},
	":malaysia:": Code{
		"\U0001f1f2\U0001f1fe",
		map[string]struct{}{},
	},
	":mozambique:": Code{
		"\U0001f1f2\U0001f1ff",
		map[string]struct{}{},
	},
	":namibia:": Code{
		"\U0001f1f3\U0001f1e6",
		map[string]struct{}{},
	},
	":new_caledonia:": Code{
		"\U0001f1f3\U0001f1e8",
		map[string]struct{}{},
	},
	":niger:": Code{
		"\U0001f1f3\U0001f1ea",
		map[string]struct{}{},
	},
	":norfolk_island:": Code{
		"\U0001f1f3\U0001f1eb",
		map[string]struct{}{},
	},
	":nigeria:": Code{
		"\U0001f1f3\U0001f1ec",
		map[string]struct{}{},
	},
	":nicaragua:": Code{
		"\U0001f1f3\U0001f1ee",
		map[string]struct{}{},
	},
	":netherlands:": Code{
		"\U0001f1f3\U0001f1f1",
		map[string]struct{}{},
	},
	":norway:": Code{
		"\U0001f1f3\U0001f1f4",
		map[string]struct{}{},
	},
	":nepal:": Code{
		"\U0001f1f3\U0001f1f5",
		map[string]struct{}{},
	},
	":nauru:": Code{
		"\U0001f1f3\U0001f1f7",
		map[string]struct{}{},
	},
	":niue:": Code{
		"\U0001f1f3\U0001f1fa",
		map[string]struct{}{},
	},
	":new_zealand:": Code{
		"\U0001f1f3\U0001f1ff",
		map[string]struct{}{},
	},
	":oman:": Code{
		"\U0001f1f4\U0001f1f2",
		map[string]struct{}{},
	},
	":panama:": Code{
		"\U0001f1f5\U0001f1e6",
		map[string]struct{}{},
	},
	":peru:": Code{
		"\U0001f1f5\U0001f1ea",
		map[string]struct{}{},
	},
	":french_polynesia:": Code{
		"\U0001f1f5\U0001f1eb",
		map[string]struct{}{},
	},
	":papua_new_guinea:": Code{
		"\U0001f1f5\U0001f1ec",
		map[string]struct{}{},
	},
	":philippines:": Code{
		"\U0001f1f5\U0001f1ed",
		map[string]struct{}{},
	},
	":pakistan:": Code{
		"\U0001f1f5\U0001f1f0",
		map[string]struct{}{},
	},
	":poland:": Code{
		"\U0001f1f5\U0001f1f1",
		map[string]struct{}{},
	},
	":st_pierre_miquelon:": Code{
		"\U0001f1f5\U0001f1f2",
		map[string]struct{}{},
	},
	":pitcairn_islands:": Code{
		"\U0001f1f5\U0001f1f3",
		map[string]struct{}{},
	},
	":puerto_rico:": Code{
		"\U0001f1f5\U0001f1f7",
		map[string]struct{}{},
	},
	":palestinian_territories:": Code{
		"\U0001f1f5\U0001f1f8",
		map[string]struct{}{},
	},
	":portugal:": Code{
		"\U0001f1f5\U0001f1f9",
		map[string]struct{}{},
	},
	":palau:": Code{
		"\U0001f1f5\U0001f1fc",
		map[string]struct{}{},
	},
	":paraguay:": Code{
		"\U0001f1f5\U0001f1fe",
		map[string]struct{}{},
	},
	":qatar:": Code{
		"\U0001f1f6\U0001f1e6",
		map[string]struct{}{},
	},
	":reunion:": Code{
		"\U0001f1f7\U0001f1ea",
		map[string]struct{}{},
	},
	":romania:": Code{
		"\U0001f1f7\U0001f1f4",
		map[string]struct{}{},
	},
	":serbia:": Code{
		"\U0001f1f7\U0001f1f8",
		map[string]struct{}{},
	},
	":ru:": Code{
		"\U0001f1f7\U0001f1fa",
		map[string]struct{}{
			"russia": struct{}{},
		},
	},
	":rwanda:": Code{
		"\U0001f1f7\U0001f1fc",
		map[string]struct{}{},
	},
	":saudi_arabia:": Code{
		"\U0001f1f8\U0001f1e6",
		map[string]struct{}{},
	},
	":solomon_islands:": Code{
		"\U0001f1f8\U0001f1e7",
		map[string]struct{}{},
	},
	":seychelles:": Code{
		"\U0001f1f8\U0001f1e8",
		map[string]struct{}{},
	},
	":sudan:": Code{
		"\U0001f1f8\U0001f1e9",
		map[string]struct{}{},
	},
	":sweden:": Code{
		"\U0001f1f8\U0001f1ea",
		map[string]struct{}{},
	},
	":singapore:": Code{
		"\U0001f1f8\U0001f1ec",
		map[string]struct{}{},
	},
	":st_helena:": Code{
		"\U0001f1f8\U0001f1ed",
		map[string]struct{}{},
	},
	":slovenia:": Code{
		"\U0001f1f8\U0001f1ee",
		map[string]struct{}{},
	},
	":svalbard_jan_mayen:": Code{
		"\U0001f1f8\U0001f1ef",
		map[string]struct{}{},
	},
	":slovakia:": Code{
		"\U0001f1f8\U0001f1f0",
		map[string]struct{}{},
	},
	":sierra_leone:": Code{
		"\U0001f1f8\U0001f1f1",
		map[string]struct{}{},
	},
	":san_marino:": Code{
		"\U0001f1f8\U0001f1f2",
		map[string]struct{}{},
	},
	":senegal:": Code{
		"\U0001f1f8\U0001f1f3",
		map[string]struct{}{},
	},
	":somalia:": Code{
		"\U0001f1f8\U0001f1f4",
		map[string]struct{}{},
	},
	":suriname:": Code{
		"\U0001f1f8\U0001f1f7",
		map[string]struct{}{},
	},
	":south_sudan:": Code{
		"\U0001f1f8\U0001f1f8",
		map[string]struct{}{},
	},
	":sao_tome_principe:": Code{
		"\U0001f1f8\U0001f1f9",
		map[string]struct{}{},
	},
	":el_salvador:": Code{
		"\U0001f1f8\U0001f1fb",
		map[string]struct{}{},
	},
	":sint_maarten:": Code{
		"\U0001f1f8\U0001f1fd",
		map[string]struct{}{},
	},
	":syria:": Code{
		"\U0001f1f8\U0001f1fe",
		map[string]struct{}{},
	},
	":swaziland:": Code{
		"\U0001f1f8\U0001f1ff",
		map[string]struct{}{},
	},
	":tristan_da_cunha:": Code{
		"\U0001f1f9\U0001f1e6",
		map[string]struct{}{},
	},
	":turks_caicos_islands:": Code{
		"\U0001f1f9\U0001f1e8",
		map[string]struct{}{},
	},
	":chad:": Code{
		"\U0001f1f9\U0001f1e9",
		map[string]struct{}{},
	},
	":french_southern_territories:": Code{
		"\U0001f1f9\U0001f1eb",
		map[string]struct{}{},
	},
	":togo:": Code{
		"\U0001f1f9\U0001f1ec",
		map[string]struct{}{},
	},
	":thailand:": Code{
		"\U0001f1f9\U0001f1ed",
		map[string]struct{}{},
	},
	":tajikistan:": Code{
		"\U0001f1f9\U0001f1ef",
		map[string]struct{}{},
	},
	":tokelau:": Code{
		"\U0001f1f9\U0001f1f0",
		map[string]struct{}{},
	},
	":timor_leste:": Code{
		"\U0001f1f9\U0001f1f1",
		map[string]struct{}{},
	},
	":turkmenistan:": Code{
		"\U0001f1f9\U0001f1f2",
		map[string]struct{}{},
	},
	":tunisia:": Code{
		"\U0001f1f9\U0001f1f3",
		map[string]struct{}{},
	},
	":tonga:": Code{
		"\U0001f1f9\U0001f1f4",
		map[string]struct{}{},
	},
	":tr:": Code{
		"\U0001f1f9\U0001f1f7",
		map[string]struct{}{
			"turkey": struct{}{},
		},
	},
	":trinidad_tobago:": Code{
		"\U0001f1f9\U0001f1f9",
		map[string]struct{}{},
	},
	":tuvalu:": Code{
		"\U0001f1f9\U0001f1fb",
		map[string]struct{}{},
	},
	":taiwan:": Code{
		"\U0001f1f9\U0001f1fc",
		map[string]struct{}{},
	},
	":tanzania:": Code{
		"\U0001f1f9\U0001f1ff",
		map[string]struct{}{},
	},
	":ukraine:": Code{
		"\U0001f1fa\U0001f1e6",
		map[string]struct{}{},
	},
	":uganda:": Code{
		"\U0001f1fa\U0001f1ec",
		map[string]struct{}{},
	},
	":us_outlying_islands:": Code{
		"\U0001f1fa\U0001f1f2",
		map[string]struct{}{},
	},
	":us:": Code{
		"\U0001f1fa\U0001f1f8",
		map[string]struct{}{
			"flag":    struct{}{},
			"united":  struct{}{},
			"america": struct{}{},
		},
	},
	":uruguay:": Code{
		"\U0001f1fa\U0001f1fe",
		map[string]struct{}{},
	},
	":uzbekistan:": Code{
		"\U0001f1fa\U0001f1ff",
		map[string]struct{}{},
	},
	":vatican_city:": Code{
		"\U0001f1fb\U0001f1e6",
		map[string]struct{}{},
	},
	":st_vincent_grenadines:": Code{
		"\U0001f1fb\U0001f1e8",
		map[string]struct{}{},
	},
	":venezuela:": Code{
		"\U0001f1fb\U0001f1ea",
		map[string]struct{}{},
	},
	":british_virgin_islands:": Code{
		"\U0001f1fb\U0001f1ec",
		map[string]struct{}{},
	},
	":us_virgin_islands:": Code{
		"\U0001f1fb\U0001f1ee",
		map[string]struct{}{},
	},
	":vietnam:": Code{
		"\U0001f1fb\U0001f1f3",
		map[string]struct{}{},
	},
	":vanuatu:": Code{
		"\U0001f1fb\U0001f1fa",
		map[string]struct{}{},
	},
	":wallis_futuna:": Code{
		"\U0001f1fc\U0001f1eb",
		map[string]struct{}{},
	},
	":samoa:": Code{
		"\U0001f1fc\U0001f1f8",
		map[string]struct{}{},
	},
	":kosovo:": Code{
		"\U0001f1fd\U0001f1f0",
		map[string]struct{}{},
	},
	":yemen:": Code{
		"\U0001f1fe\U0001f1ea",
		map[string]struct{}{},
	},
	":mayotte:": Code{
		"\U0001f1fe\U0001f1f9",
		map[string]struct{}{},
	},
	":south_africa:": Code{
		"\U0001f1ff\U0001f1e6",
		map[string]struct{}{},
	},
	":zambia:": Code{
		"\U0001f1ff\U0001f1f2",
		map[string]struct{}{},
	},
	":zimbabwe:": Code{
		"\U0001f1ff\U0001f1fc",
		map[string]struct{}{},
	},
}
