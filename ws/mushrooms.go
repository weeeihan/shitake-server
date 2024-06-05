package ws

var GPTmushrooms = map[int]Mushroom{
	1: {
		Name:        "White button",
		Damage:      2,
		Description: "Agaricus bisporus, commonly known as the cultivated mushroom, is a basidiomycete mushroom native to grasslands in Eurasia and North America. It is cultivated in more than 70 countries and is one of the most commonly and widely consumed mushrooms in the world.",
		Color:       "white",
	},
	2: {
		Name:        "Enoki",
		Damage:      2,
		Description: "Flammulina filiformis is a species of edible agaric in the family Physalacriaceae. It is widely cultivated in East Asia, and well known for its role in Japanese and Chinese cuisine.",
		Color:       "white",
	},
	3: {
		Name:        "Morel",
		Damage:      2,
		Description: "Morchella, the true morels, is a genus of edible sac fungi closely related to anatomically simpler cup fungi in the order Pezizales. These distinctive fungi have a honeycomb appearance due to the network of ridges with pits composing their caps.",
		Color:       "black",
	},
	4: {
		Name:        "Shiitake",
		Damage:      2,
		Description: "Lentinula edodes is a species of edible mushroom native to East Asia, which is cultivated and consumed in many Asian countries. It is considered a medicinal mushroom in some forms of traditional medicine.",
		Color:       "brown",
	},
	5: {
		Name:        "Oyster",
		Damage:      2,
		Description: "Pleurotus ostreatus, the oyster mushroom, is a common edible mushroom. It was first cultivated in Germany as a subsistence measure during World War I and is now grown commercially around the world for food.",
		Color:       "white",
	},
	6: {
		Name:        "Porcini",
		Damage:      2,
		Description: "Boletus edulis is a basidiomycete fungus, and the type species of the genus Boletus. Widely distributed in the Northern Hemisphere across Europe, Asia, and North America, it does not occur naturally in the Southern Hemisphere, although it has been introduced to southern Africa, Australia, and New Zealand.",
		Color:       "brown",
	},
	7: {
		Name:        "Chanterelle",
		Damage:      2,
		Description: "Cantharellus cibarius, commonly known as the chanterelle, golden chanterelle or girolle, is a fungus. It is probably the best known species of the genus Cantharellus, if not the entire family of Cantharellaceae.",
		Color:       "yellow",
	},
	8: {
		Name:        "Lion's Mane",
		Damage:      2,
		Description: "Hericium erinaceus is a species of tooth fungus in the family Hericiaceae. It is native to North America, Europe, and Asia. It can be mistaken for other species of Hericium, all popular edibles, which grow across the same range.",
		Color:       "white",
	},
	9: {
		Name:        "Reishi",
		Damage:      2,
		Description: "Ganoderma lucidum is a species of bracket fungus, and the type species of the genus Ganoderma. It lives on deadwood, especially dead trees, and is generally considered to be a saprotroph, rather than a parasite.",
		Color:       "red",
	},
	10: {
		Name:        "Maitake",
		Damage:      2,
		Description: "Grifola frondosa is a polypore mushroom that grows in clusters at the base of trees, particularly oaks. The mushroom is commonly known among English speakers as hen of the woods, ram's head, and sheep's head.",
		Color:       "white",
	},
	11: {
		Name:        "King Oyster",
		Damage:      2,
		Description: "Pleurotus eryngii, also known as king trumpet mushroom, is an edible mushroom native to Mediterranean regions of Europe, the Middle East, and North Africa.",
		Color:       "brown",
	},
	12: {
		Name:        "Shimeji",
		Damage:      2,
		Description: "Hypsizygus tessellatus, also known as shimeji, is an edible mushroom native to East Asia. It has a slightly nutty flavor and is used in soups, stews, and stir-fries.",
		Color:       "brown",
	},
	13: {
		Name:        "Beech",
		Damage:      2,
		Description: "Hypsizygus marmoreus, known as beech mushroom, is native to East Asia and known for its crunchy texture and sweet, nutty flavor.",
		Color:       "white",
	},
	14: {
		Name:        "Wood Ear",
		Damage:      2,
		Description: "Auricularia auricula-judae, known as wood ear or jelly ear, is a species of edible fungus found worldwide. It is commonly used in Chinese cuisine for its unique texture.",
		Color:       "brown",
	},
	15: {
		Name:        "Chicken of the Woods",
		Damage:      2,
		Description: "Laetiporus sulphureus, commonly known as chicken of the woods, is an edible polypore mushroom found primarily in North America. It has a texture similar to chicken when cooked.",
		Color:       "yellow",
	},
	16: {
		Name:        "Puffball",
		Damage:      2,
		Description: "Lycoperdon perlatum, known as the common puffball, is an edible mushroom with a round shape and a white to brownish color. It releases spores in a puff when mature.",
		Color:       "white",
	},
	17: {
		Name:        "Caesar's Mushroom",
		Damage:      2,
		Description: "Amanita caesarea, also known as Caesar's mushroom, is a highly regarded edible mushroom native to southern Europe and North Africa. It has an orange cap and a rich, nutty flavor.",
		Color:       "orange",
	},
	18: {
		Name:        "Black Trumpet",
		Damage:      2,
		Description: "Craterellus cornucopioides, known as black trumpet or horn of plenty, is an edible mushroom with a dark, trumpet-shaped cap. It is prized for its rich, smoky flavor.",
		Color:       "black",
	},
	19: {
		Name:        "Turkey Tail",
		Damage:      2,
		Description: "Trametes versicolor, commonly known as turkey tail, is a polypore mushroom found throughout the world. It is known for its colorful, fan-shaped fruiting bodies and medicinal properties.",
		Color:       "multi-colored",
	},
	20: {
		Name:        "King Bolete",
		Damage:      2,
		Description: "Boletus edulis, also known as the king bolete, is a highly sought-after edible mushroom found in Europe, Asia, and North America. It has a large, brown cap and a thick, white stalk.",
		Color:       "brown",
	},
	21: {
		Name:        "Matsutake",
		Damage:      2,
		Description: "Tricholoma matsutake is an aromatic mushroom highly prized in Japan, Korea, and China. It grows under trees and is harvested in the wild.",
		Color:       "brown",
	},
	22: {
		Name:        "Yellowfoot",
		Damage:      2,
		Description: "Craterellus tubaeformis, known as yellowfoot or winter mushroom, is an edible fungus found in coniferous forests. It has a yellow stem and a brown cap.",
		Color:       "yellow",
	},
	23: {
		Name:        "Lobster Mushroom",
		Damage:      2,
		Description: "Hypomyces lactifluorum, commonly known as lobster mushroom, is a parasitic fungus that grows on certain species of mushrooms, transforming them into bright orange, edible mushrooms with a seafood-like flavor.",
		Color:       "orange",
	},
	24: {
		Name:        "Candy Cap",
		Damage:      2,
		Description: "Lactarius rubidus, known as candy cap, is an edible mushroom known for its sweet, maple syrup-like aroma. It is used in desserts and baking.",
		Color:       "red",
	},
	25: {
		Name:        "Bluefoot",
		Damage:      2,
		Description: "Clitocybe nuda, also known as bluefoot, is an edible mushroom with a bluish-purple cap and stem. It has a mild flavor and is used in various culinary dishes.",
		Color:       "blue",
	},
	26: {
		Name:        "Shrimp Mushroom",
		Damage:      2,
		Description: "Russula xerampelina, commonly known as shrimp mushroom, is an edible fungus with a reddish cap and a shrimp-like odor. It is found in coniferous forests.",
		Color:       "red",
	},
	27: {
		Name:        "St. George's Mushroom",
		Damage:      2,
		Description: "Calocybe gambosa, known as St. George's mushroom, is an edible species found in Europe. It has a white cap and a strong, mealy odor.",
		Color:       "white",
	},
	28: {
		Name:        "Blewit",
		Damage:      2,
		Description: "Clitocybe nuda, also known as blewit or wood blewit, is an edible mushroom with a purple cap and stem. It is commonly found in leaf litter and compost.",
		Color:       "purple",
	},
	29: {
		Name:        "Fairy Ring Mushroom",
		Damage:      2,
		Description: "Marasmius oreades, commonly known as the fairy ring mushroom, is an edible species that forms rings in grassy areas. It has a sweet flavor and a tough texture.",
		Color:       "brown",
	},
	30: {
		Name:        "Hedgehog Mushroom",
		Damage:      2,
		Description: "Hydnum repandum, known as the hedgehog mushroom, is an edible species with a yellow to orange cap and spines instead of gills. It has a nutty flavor.",
		Color:       "orange",
	},
	31: {
		Name:        "Cauliflower Mushroom",
		Damage:      2,
		Description: "Sparassis crispa, known as the cauliflower mushroom, is an edible species with a distinctive, frilly appearance. It grows at the base of trees and has a mild flavor.",
		Color:       "white",
	},
	32: {
		Name:        "Milk Cap",
		Damage:      2,
		Description: "Lactarius deliciosus, commonly known as the saffron milk cap, is an edible mushroom with a bright orange cap and stem. It exudes a carrot-colored milk when cut.",
		Color:       "orange",
	},
	33: {
		Name:        "False Morel",
		Damage:      2,
		Description: "Gyromitra esculenta, known as the false morel, is an edible mushroom that must be properly cooked to remove toxins. It has a brain-like appearance.",
		Color:       "brown",
	},
	34: {
		Name:        "Brick Cap",
		Damage:      2,
		Description: "Hypholoma sublateritium, known as the brick cap, is an edible mushroom with a reddish-brown cap. It grows in clusters on decaying wood.",
		Color:       "red",
	},
	35: {
		Name:        "Parasol Mushroom",
		Damage:      2,
		Description: "Macrolepiota procera, commonly known as the parasol mushroom, is an edible species with a large, umbrella-like cap and a scaly surface.",
		Color:       "brown",
	},
	36: {
		Name:        "Honey Fungus",
		Damage:      2,
		Description: "Armillaria mellea, known as the honey fungus, is an edible mushroom that grows in clusters on wood. It has a sweet flavor and a yellow to brown cap.",
		Color:       "yellow",
	},
	37: {
		Name:        "Deer Mushroom",
		Damage:      2,
		Description: "Pluteus cervinus, known as the deer mushroom, is an edible species with a brown cap and pink gills. It grows on decaying wood and has a mild taste.",
		Color:       "brown",
	},
	38: {
		Name:        "Snow Fungus",
		Damage:      2,
		Description: "Tremella fuciformis, known as snow fungus or white jelly mushroom, is an edible species used in Chinese cuisine. It has a gelatinous texture and is often used in desserts.",
		Color:       "white",
	},
	39: {
		Name:        "Pink Oyster",
		Damage:      2,
		Description: "Pleurotus djamor, known as pink oyster mushroom, is an edible species with a vibrant pink color. It has a delicate texture and a mild, seafood-like flavor.",
		Color:       "pink",
	},
	40: {
		Name:        "Elm Oyster",
		Damage:      2,
		Description: "Hypsizygus ulmarius, known as the elm oyster, is an edible mushroom that grows on elm trees. It has a white to pale brown cap and a firm texture.",
		Color:       "white",
	},
	41: {
		Name:        "Chestnut Mushroom",
		Damage:      2,
		Description: "Pholiota adiposa, known as the chestnut mushroom, is an edible species with a reddish-brown cap and a nutty flavor. It grows in clusters on decaying wood.",
		Color:       "brown",
	},
	42: {
		Name:        "Amber Jelly",
		Damage:      2,
		Description: "Exidia recisa, known as amber jelly, is an edible jelly fungus found on dead deciduous trees. It has a translucent, amber color and a gelatinous texture.",
		Color:       "amber",
	},
	43: {
		Name:        "Red-Capped Scaber Stalk",
		Damage:      2,
		Description: "Leccinum aurantiacum, known as red-capped scaber stalk, is an edible bolete with a red cap and a scaly stem. It is found in forests of birch trees.",
		Color:       "red",
	},
	44: {
		Name:        "Winter Mushroom",
		Damage:      2,
		Description: "Flammulina velutipes, known as the winter mushroom, is an edible species with a yellow to orange cap and a dark stem. It grows during cold weather.",
		Color:       "orange",
	},
	45: {
		Name:        "Velvet Foot",
		Damage:      2,
		Description: "Flammulina velutipes, also known as velvet foot, is an edible mushroom with a yellow to brown cap and a velvety stem. It grows on dead wood in cold weather.",
		Color:       "brown",
	},
	46: {
		Name:        "Sulfur Tuft",
		Damage:      2,
		Description: "Hypholoma fasciculare, known as sulfur tuft, is an inedible mushroom with a yellow cap and greenish gills. It grows in clusters on dead wood.",
		Color:       "yellow",
	},
	47: {
		Name:        "Fly Agaric",
		Damage:      2,
		Description: "Amanita muscaria, known as fly agaric, is a highly recognizable mushroom with a red cap and white spots. It is toxic and hallucinogenic.",
		Color:       "red",
	},
	48: {
		Name:        "Death Cap",
		Damage:      2,
		Description: "Amanita phalloides, known as the death cap, is a highly toxic mushroom with a greenish cap. It is responsible for the majority of mushroom poisoning deaths.",
		Color:       "green",
	},
	49: {
		Name:        "Destroying Angel",
		Damage:      2,
		Description: "Amanita bisporigera, known as the destroying angel, is a deadly poisonous mushroom with a white cap and stem. It is found in North America.",
		Color:       "white",
	},
	50: {
		Name:        "Panther Cap",
		Damage:      2,
		Description: "Amanita pantherina, known as the panther cap, is a toxic mushroom with a brown cap and white spots. It contains psychoactive compounds.",
		Color:       "brown",
	},
	51: {
		Name:        "Jack-o'-Lantern",
		Damage:      2,
		Description: "Omphalotus olearius, known as the jack-o'-lantern mushroom, is a toxic species with a bright orange color and bioluminescent properties.",
		Color:       "orange",
	},
	52: {
		Name:        "Green-Spored Parasol",
		Damage:      2,
		Description: "Chlorophyllum molybdites, known as the green-spored parasol, is a toxic mushroom with a large, white cap and greenish spores.",
		Color:       "white",
	},
	53: {
		Name:        "Ivory Funnel",
		Damage:      2,
		Description: "Clitocybe dealbata, known as the ivory funnel, is a poisonous mushroom with a white to ivory cap and a funnel shape. It contains muscarine.",
		Color:       "white",
	},
	54: {
		Name:        "Deadly Webcap",
		Damage:      2,
		Description: "Cortinarius rubellus, known as the deadly webcap, is a toxic mushroom with a reddish-brown cap and a web-like partial veil. It contains orellanine.",
		Color:       "brown",
	},
	55: {
		Name:        "Autumn Skullcap",
		Damage:      2,
		Description: "Galerina marginata, known as the autumn skullcap, is a toxic mushroom with a brown cap and gills. It contains amatoxins similar to those in the death cap.",
		Color:       "brown",
	},
	56: {
		Name:        "Conocybe Filaris",
		Damage:      2,
		Description: "Conocybe filaris is a highly toxic mushroom with a conical brown cap. It contains deadly amatoxins and is found in grassy areas and woodchip mulch.",
		Color:       "brown",
	},
	57: {
		Name:        "Yellow-Staining Mushroom",
		Damage:      2,
		Description: "Agaricus xanthodermus, known as the yellow-staining mushroom, is a toxic species with a white cap that turns yellow when bruised. It has a phenolic odor.",
		Color:       "white",
	},
	58: {
		Name:        "Brown Roll-Rim",
		Damage:      2,
		Description: "Paxillus involutus, known as the brown roll-rim, is a toxic mushroom with a brown cap and inrolled margins. It causes immune-mediated hemolysis.",
		Color:       "brown",
	},
	59: {
		Name:        "False Parasol",
		Damage:      2,
		Description: "Chlorophyllum molybdites, known as the false parasol, is a toxic mushroom with a large, white cap and greenish gills. It is often confused with edible species.",
		Color:       "white",
	},
	60: {
		Name:        "Angel Wings",
		Damage:      2,
		Description: "Pleurocybella porrigens, known as angel wings, is a toxic mushroom with a white, fan-shaped cap. It grows on dead conifer wood and causes neurological symptoms.",
		Color:       "white",
	},
	61: {
		Name:        "Lilac Bonnet",
		Damage:      2,
		Description: "Mycena pura, known as the lilac bonnet, is a toxic mushroom with a lilac-colored cap and gills. It contains muscarine and other toxic compounds.",
		Color:       "purple",
	},
	62: {
		Name:        "Yellow Knight",
		Damage:      2,
		Description: "Tricholoma equestre, known as the yellow knight, is a toxic mushroom with a yellow cap and gills. It causes rhabdomyolysis and kidney failure.",
		Color:       "yellow",
	},
	63: {
		Name:        "Sulfur Shelf",
		Damage:      2,
		Description: "Laetiporus sulphureus, known as sulfur shelf or chicken of the woods, is an edible mushroom with a bright orange color. It has a texture similar to chicken when cooked.",
		Color:       "orange",
	},
	64: {
		Name:        "Shaggy Ink Cap",
		Damage:      2,
		Description: "Coprinus comatus, known as the shaggy ink cap or lawyer's wig, is an edible mushroom that dissolves into a black ink when mature.",
		Color:       "white",
	},
	65: {
		Name:        "Gem-Studded Puffball",
		Damage:      2,
		Description: "Lycoperdon perlatum, known as the gem-studded puffball, is an edible mushroom with a round shape and a white to brownish color. It releases spores in a puff when mature.",
		Color:       "white",
	},
	66: {
		Name:        "Purple-Spored Puffball",
		Damage:      2,
		Description: "Calvatia cyathiformis, known as the purple-spored puffball, is an edible mushroom with a large, round shape and purple spores. It grows in open fields and meadows.",
		Color:       "purple",
	},
	67: {
		Name:        "Smooth Earthball",
		Damage:      2,
		Description: "Scleroderma citrinum, known as the smooth earthball, is a toxic mushroom with a yellow to brown, rough outer surface. It contains toxic compounds that cause gastrointestinal distress.",
		Color:       "brown",
	},
	68: {
		Name:        "Pigskin Poison Puffball",
		Damage:      2,
		Description: "Scleroderma citrinum, known as the pigskin poison puffball, is a toxic mushroom with a thick, rough outer surface. It causes severe gastrointestinal symptoms if ingested.",
		Color:       "brown",
	},
	69: {
		Name:        "Witch's Hat",
		Damage:      2,
		Description: "Hygrocybe conica, known as the witch's hat, is an inedible mushroom with a conical cap that changes color from yellow to black. It grows in grasslands and wooded areas.",
		Color:       "yellow",
	},
	70: {
		Name:        "Earthstar",
		Damage:      2,
		Description: "Geastrum triplex, known as the collared earthstar, is an inedible mushroom with a star-shaped fruiting body. It is found in forests and grassy areas.",
		Color:       "brown",
	},
	71: {
		Name:        "Dog Stinkhorn",
		Damage:      2,
		Description: "Mutinus caninus, known as the dog stinkhorn, is an inedible mushroom with a foul odor. It has a phallic shape and attracts flies to disperse its spores.",
		Color:       "white",
	},
	72: {
		Name:        "Elegant Stinkhorn",
		Damage:      2,
		Description: "Mutinus elegans, known as the elegant stinkhorn, is an inedible mushroom with a foul odor. It has a reddish cap and a white stem.",
		Color:       "red",
	},
	73: {
		Name:        "Stinkhorn",
		Damage:      2,
		Description: "Phallus impudicus, known as the common stinkhorn, is an inedible mushroom with a phallic shape and a foul odor. It grows in forests and gardens.",
		Color:       "white",
	},
	74: {
		Name:        "Dead Man's Fingers",
		Damage:      2,
		Description: "Xylaria polymorpha, known as dead man's fingers, is an inedible fungus with black, finger-like fruiting bodies. It grows on decaying wood.",
		Color:       "black",
	},
	75: {
		Name:        "Devil's Urn",
		Damage:      2,
		Description: "Urnula craterium, known as the devil's urn, is an inedible fungus with dark, urn-shaped fruiting bodies. It grows on decaying wood in forests.",
		Color:       "black",
	},
	76: {
		Name:        "Elfin Saddle",
		Damage:      2,
		Description: "Helvella lacunosa, known as the elfin saddle, is an inedible mushroom with a dark, irregularly shaped cap. It grows in forests and wooded areas.",
		Color:       "black",
	},
	77: {
		Name:        "Scaly Vase Chanterelle",
		Damage:      2,
		Description: "Turbinellus floccosus, known as the scaly vase chanterelle, is an inedible mushroom with a vase-shaped cap covered in scales. It grows in coniferous forests.",
		Color:       "orange",
	},
	78: {
		Name:        "Hygrocybe conica",
		Damage:      2,
		Description: "Hygrocybe conica, known as the witch's hat, is an inedible mushroom with a conical cap that changes color from yellow to black. It grows in grasslands and wooded areas.",
		Color:       "yellow",
	},
	79: {
		Name:        "Scaly Chanterelle",
		Damage:      2,
		Description: "Gomphus floccosus, known as the scaly chanterelle, is an inedible mushroom with a vase-shaped cap covered in scales. It grows in coniferous forests.",
		Color:       "orange",
	},
	80: {
		Name:        "Penny Bun",
		Damage:      2,
		Description: "Boletus edulis, known as the penny bun, is an edible mushroom with a large, brown cap and a thick, white stem. It is found in forests of birch trees.",
		Color:       "brown",
	},
	81: {
		Name:        "Red-Capped Scaber Stalk",
		Damage:      2,
		Description: "Leccinum aurantiacum, known as the red-capped scaber stalk, is an edible bolete with a red cap and a scaly stem. It is found in forests of birch trees.",
		Color:       "red",
	},
	82: {
		Name:        "Orange Birch Bolete",
		Damage:      2,
		Description: "Leccinum versipelle, known as the orange birch bolete, is an edible mushroom with an orange cap and a scaly stem. It grows in birch forests.",
		Color:       "orange",
	},
	83: {
		Name:        "Scaly Chanterelle",
		Damage:      2,
		Description: "Gomphus floccosus, known as the scaly chanterelle, is an inedible mushroom with a vase-shaped cap covered in scales. It grows in coniferous forests.",
		Color:       "orange",
	},
	84: {
		Name:        "Woolly Milkcap",
		Damage:      2,
		Description: "Lactarius torminosus, known as the woolly milkcap, is an inedible mushroom with a pinkish cap covered in woolly hairs. It exudes a white, acrid latex when cut.",
		Color:       "pink",
	},
	85: {
		Name:        "Liberty Cap",
		Damage:      2,
		Description: "Psilocybe semilanceata, known as the liberty cap, is a psychoactive mushroom with a conical cap and a slender stem. It contains the hallucinogen psilocybin.",
		Color:       "brown",
	},
	86: {
		Name:        "Fly Agaric",
		Damage:      2,
		Description: "Amanita muscaria, known as fly agaric, is a highly recognizable mushroom with a red cap and white spots. It is toxic and hallucinogenic.",
		Color:       "red",
	},
	87: {
		Name:        "Yellow Stainer",
		Damage:      2,
		Description: "Agaricus xanthodermus, known as the yellow stainer, is a toxic mushroom with a white cap that turns yellow when bruised. It has a phenolic odor.",
		Color:       "white",
	},
	88: {
		Name:        "Sulphur Tuft",
		Damage:      2,
		Description: "Hypholoma fasciculare, known as the sulphur tuft, is an inedible mushroom with a yellow cap and greenish gills. It grows in clusters on dead wood.",
		Color:       "yellow",
	},
	89: {
		Name:        "Bleeding Tooth Fungus",
		Damage:      2,
		Description: "Hydnellum peckii, known as the bleeding tooth fungus, is a unique mushroom with a white cap that exudes red fluid, giving it a 'bleeding' appearance.",
		Color:       "white",
	},
	90: {
		Name:        "Black Trumpet",
		Damage:      2,
		Description: "Craterellus cornucopioides, known as the black trumpet, is an edible mushroom with a dark, trumpet-shaped cap and a rich, smoky flavor.",
		Color:       "black",
	},
	91: {
		Name:        "Puffball",
		Damage:      2,
		Description: "Calvatia gigantea, known as the giant puffball, is an edible mushroom that can grow to be quite large. When mature, it releases a cloud of spores.",
		Color:       "white",
	},
	92: {
		Name:        "Shaggy Mane",
		Damage:      2,
		Description: "Coprinus comatus, known as the shaggy mane, is an edible mushroom with a tall, white cap that turns black and inky as it matures.",
		Color:       "white",
	},
	93: {
		Name:        "Red Cage",
		Damage:      2,
		Description: "Clathrus ruber, known as the red cage, is an inedible mushroom with a striking red, lattice-like structure and a foul odor.",
		Color:       "red",
	},
	94: {
		Name:        "Jack O'Lantern",
		Damage:      2,
		Description: "Omphalotus illudens, known as the jack o'lantern mushroom, is a toxic mushroom that glows in the dark due to bioluminescence.",
		Color:       "orange",
	},
	95: {
		Name:        "Parrot Waxcap",
		Damage:      2,
		Description: "Gliophorus psittacinus, known as the parrot waxcap, is an edible mushroom with a vibrant green cap that turns yellow with age.",
		Color:       "green",
	},
	96: {
		Name:        "Candy Cap",
		Damage:      2,
		Description: "Lactarius rubidus, known as the candy cap, is an edible mushroom with a sweet, maple syrup-like scent when dried.",
		Color:       "orange",
	},
	97: {
		Name:        "Indigo Milk Cap",
		Damage:      2,
		Description: "Lactarius indigo, known as the indigo milk cap, is an edible mushroom with a distinctive blue color and milk-like latex.",
		Color:       "blue",
	},
	98: {
		Name:        "Amethyst Deceiver",
		Damage:      2,
		Description: "Laccaria amethystina, known as the amethyst deceiver, is an edible mushroom with a purple cap that can fade with age.",
		Color:       "purple",
	},
	99: {
		Name:        "Green-Spored Parasol",
		Damage:      2,
		Description: "Chlorophyllum molybdites, known as the green-spored parasol, is a toxic mushroom with a large white cap and green spores.",
		Color:       "white",
	},
	100: {
		Name:        "Caesar's Mushroom",
		Damage:      2,
		Description: "Amanita caesarea, known as Caesar's mushroom, is an edible mushroom with an orange cap and a long history of consumption in Roman cuisine.",
		Color:       "orange",
	},
	101: {
		Name:        "Amethyst Deceiver",
		Damage:      2,
		Description: "Laccaria amethystina, known as the amethyst deceiver, is an edible mushroom with a striking purple color. It is found in deciduous and coniferous forests and is notable for its deep purple cap and stem.",
		Color:       "Purple",
	},
	102: {
		Name:        "Golden Chanterelle",
		Damage:      2,
		Description: "Cantharellus cibarius, commonly known as the golden chanterelle, is an edible mushroom valued for its rich flavor and aroma. It has a bright yellow to orange color and is found in hardwood forests.",
		Color:       "Yellow",
	},
	103: {
		Name:        "Cauliflower Fungus",
		Damage:      2,
		Description: "Sparassis spathulata, known as the cauliflower fungus, is an edible mushroom with a distinctive, frilly appearance resembling a cauliflower. It grows at the base of trees and has a mild flavor.",
		Color:       "White",
	},
	104: {
		Name:        "Red-Cracking Bolete",
		Damage:      2,
		Description: "Xerocomellus chrysenteron, known as the red-cracking bolete, is an edible mushroom with a reddish-brown cap that cracks to reveal a yellow flesh underneath. It is found in coniferous and deciduous forests.",
		Color:       "Red",
	},
}