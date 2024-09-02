--
-- Data for Name: ingredient_categories; Type: TABLE DATA; Schema: public; Owner: james
--

COPY public.ingredient_categories (name) FROM stdin;
dairy
meat
deli
alcohol
produce
baking
bakery
bread
frozen foods
spices
condiments
other
asian
latin
\.


--
-- Data for Name: ingredients; Type: TABLE DATA; Schema: public; Owner: james
--

COPY public.ingredients (name, category) FROM stdin;
boneless skinless chicken thighs	meat
olive oil	other
dijon mustard	condiments
honey	condiments
salt	baking
black pepper	spices
oregano	spices
cayenne pepper	spices
frozen fruit	frozen foods
fresh fruit	produce
coconut cream	asian
spinach	produce
protein powder	other
almonds	produce
cereal	other
milk	dairy
sourdough bread	bread
deli meat	deli
deli cheese	deli
lettuce	produce
red onion	produce
mayonaise	condiments
coffee grounds	other
coffee creamer	dairy
balsamic vinegar	condiments
brussel sprouts	produce
brown sugar	baking
brioche bread	bread
eggs	dairy
cinnamon	spices
vanilla	baking
yellow onion	produce
carrot	produce
celery	produce
tomato paste	other
peas	produce
green beans	produce
squash	produce
potato	produce
garlic	produce
thyme	spices
diced tomatoes	other
vegetable broth	other
bay leaf	spices
red pepper flakes	spices
shell pasta	other
cannellini beans	other
lemon juice	other
parmesan	dairy
vegan chicken	produce
dried basil	spices
gluten free pasta	other
parsley	produce
chicken	meat
egg noodles	other
chicken broth	other
bacon	meat
butter	dairy
flour	baking
golden potato	produce
heavy cream	dairy
pepper	spices
ancho chili powder	spices
sour cream	dairy
shredded cheese	dairy
chives	produce
ground beef	meat
chili powder	spices
cumin	spices
sugar	baking
garlic powder	spices
beef broth	other
kidney beans	other
tomato sauce	other
cheddar cheese	dairy
black beans	other
egg	dairy
seasonal vegetable	produce
plum tomatoes	produce
unsalted butter	dairy
canned plum tomatoes	other
basil leaves	produce
veggie stock	other
vegetable stock	other
onion	produce
canned tomatoes	other
worcestershire sauce	condiments
italian seasoning	spices
bay leaves	spices
macaroni	other
bagel	bread
cream cheese	dairy
mac and cheese	other
sausage	meat
mushrooms	produce
french bread	bread
garlic spread	deli
butternut squash	produce
heavy whipping cream	dairy
pasta	other
whole milk	dairy
half and half	dairy
gruyere cheese	dairy
paprika	spices
green chile	other
panko	other
carrots	produce
crushed tomatoes	other
tortellini	other
shredded parmesan	dairy
basil	produce
coriander	spices
shrimp	meat
cilantro	produce
avocado	produce
turnip	produce
broth	other
white fish	meat
clams	other
coconut flour	other
corn	other
leek	produce
oyster crackers	other
cajun seasoning	spices
salmon	meat
lemon	produce
dijon	condiments
smoked paprika	spices
chicken thighs	meat
shallot	produce
pesto	condiments
orzo	other
chicken stock	other
panko bread crumbs	other
mayo	condiments
dill weed	spices
lemon pepper	spices
broccoli	produce
feta	dairy
sweet potatoes	produce
sweet italian sausage	meat
lean ground beef	meat
fennel seeds	spices
lasagna noodles	other
ricotta	dairy
mozzarella	dairy
smoked salmon	meat
chickpeas	other
lime	produce
greek yogurt	dairy
chorizo	meat
tostada shells	other
zucchini	produce
kosher salt	spices
onion powder	spices
cotija	dairy
white miso	asian
shaoxing wine	asian
corn starch	baking
soy sauce	asian
sichuan peppercorns	spices
bell pepper	produce
neutral oil	other
peanuts	other
dried red chiles	produce
scallions	produce
anchovy fillets	other
\.


--
-- Data for Name: recipes; Type: TABLE DATA; Schema: public; Owner: james
--

COPY public.recipes (name, serves, url) FROM stdin;
Honey Mustard Chicken Thighs	\N	\N
Breakfast Smoothie	\N	\N
Breakfast Cereal	\N	\N
Sandwich	\N	\N
Coffee	\N	\N
Balsamic Glazed Brussies	\N	\N
Brookes French Toast	\N	\N
Veggie Minestrone	\N	\N
Gentle Chicken Noodle Soup	\N	\N
Chicken Noodle Soup	\N	\N
Potato Soup	\N	\N
Gentle Potato Soup	\N	\N
Chili	\N	\N
Gentle Chili	\N	\N
Omelet	\N	\N
Tomato Soup	\N	\N
Goulash	\N	\N
Bagel	\N	\N
Annies Delight	\N	\N
Garlic Bread	\N	\N
Butternut Squash Soup	\N	\N
Baked Mac And Cheese	\N	\N
Tomato Basil Tortellini Soup	\N	\N
Mexican Shrimp Bisque	\N	\N
Seafood Chowder	\N	\N
Cajun Oyster Crackers	\N	\N
Southwest Carrot Roast	\N	\N
Honey Mustard Salmon	\N	\N
Garlic Parm Green Beans	\N	\N
Lemon Pesto Chicken Thighs	\N	\N
Salmon Burgers	4	https://saltandbaker.com/easy-salmon-burgers/
Feta Broccoli	2	\N
Honey Cinnamon Sweet Potatoes	4	https://www.foodnetwork.com/recipes/tyler-florence/roasted-sweet-potatoes-with-honey-butter-recipe-1946538
Worlds Best Lasagna	8	https://www.allrecipes.com/recipe/23600/worlds-best-lasagna/
Chickpea Chorizo Tostadas	4	https://www.bonappetit.com/recipe/chickpea-and-chorizo-tostadas
Black Beans	2	https://www.allrecipes.com/recipe/63113/best-black-beans
Roasted Mexican Zucchini	4	https://www.isabeleats.com/roasted-mexican-zucchini/
Miso Butter Roasted Brocolli	4	https://cooking.nytimes.com/recipes/1023496-miso-butter-roasted-broccoli?campaign_id=58&emc=edit_ck_20221014&instance_id=74535&nl=cooking&regi_id=158234325&segment_id=109959&te=1&user_id=8f72391c4a635e63ccfdd0bf05e5a318
Gong Bao Shrimp	4	https://cooking.nytimes.com/recipes/1023350-kung-pao-shrimp?campaign_id=58&emc=edit_ck_20221014&instance_id=74535&nl=cooking&regi_id=158234325&segment_id=109959&te=1&user_id=8f72391c4a635e63ccfdd0bf05e5a318
Caramelized Shallot Pasta	4	https://cooking.nytimes.com/recipes/1020830-caramelized-shallot-pasta
\.


--
-- Data for Name: units; Type: TABLE DATA; Schema: public; Owner: james
--

COPY public.units (name) FROM stdin;
tbsp
cup
lb
tsp
oz
scoop
slice
loaf
cloves
clove
cups
qt
slices
can
pinch
heads
\.


--
-- Data for Name: recipe_ingredients; Type: TABLE DATA; Schema: public; Owner: james
--

COPY public.recipe_ingredients (recipe, ingredient, measurement, units) FROM stdin;
Honey Mustard Chicken Thighs	boneless skinless chicken thighs	2	lb
Honey Mustard Chicken Thighs	olive oil	1	tbsp
Honey Mustard Chicken Thighs	dijon mustard	0.25	cup
Honey Mustard Chicken Thighs	honey	0.25	cup
Honey Mustard Chicken Thighs	salt	0.5	tsp
Honey Mustard Chicken Thighs	black pepper	0.25	tsp
Honey Mustard Chicken Thighs	oregano	0.5	tsp
Honey Mustard Chicken Thighs	cayenne pepper	0.0625	tsp
Breakfast Smoothie	frozen fruit	1	cup
Breakfast Smoothie	fresh fruit	1	cup
Breakfast Smoothie	coconut cream	0.25	cup
Breakfast Smoothie	honey	0.25	cup
Breakfast Smoothie	spinach	0.25	cup
Breakfast Smoothie	protein powder	1	scoop
Breakfast Smoothie	almonds	0.125	cup
Breakfast Cereal	cereal	1	cup
Breakfast Cereal	milk	1	cup
Sandwich	sourdough bread	2	slice
Sandwich	deli meat	3	slice
Sandwich	deli cheese	3	slice
Sandwich	lettuce	\N	\N
Sandwich	red onion	\N	\N
Sandwich	mayonaise	1	tbsp
Sandwich	dijon mustard	1	tbsp
Coffee	coffee grounds	4	tbsp
Coffee	coffee creamer	0.125	cup
Balsamic Glazed Brussies	balsamic vinegar	2	cup
Balsamic Glazed Brussies	brussel sprouts	2	cup
Balsamic Glazed Brussies	brown sugar	0.5	cup
Brookes French Toast	brioche bread	1	loaf
Brookes French Toast	eggs	4	\N
Brookes French Toast	cinnamon	1	tbsp
Brookes French Toast	vanilla	1	tbsp
Veggie Minestrone	yellow onion	1	\N
Veggie Minestrone	carrot	2	\N
Veggie Minestrone	celery	2	\N
Veggie Minestrone	tomato paste	0.25	cup
Veggie Minestrone	peas	0.5	cup
Veggie Minestrone	green beans	0.5	cup
Veggie Minestrone	squash	1	\N
Veggie Minestrone	potato	2	\N
Veggie Minestrone	garlic	4	cloves
Veggie Minestrone	oregano	0.5	tsp
Veggie Minestrone	thyme	0.5	tsp
Veggie Minestrone	diced tomatoes	28	oz
Veggie Minestrone	vegetable broth	32	oz
Veggie Minestrone	bay leaf	2	\N
Veggie Minestrone	red pepper flakes	0.125	tsp
Veggie Minestrone	shell pasta	1	cup
Veggie Minestrone	cannellini beans	1.5	cup
Veggie Minestrone	spinach	2	cup
Veggie Minestrone	lemon juice	2	tsp
Veggie Minestrone	parmesan	1	cup
Gentle Chicken Noodle Soup	yellow onion	1	\N
Gentle Chicken Noodle Soup	carrot	3	\N
Gentle Chicken Noodle Soup	celery	3	\N
Gentle Chicken Noodle Soup	garlic	4	cloves
Gentle Chicken Noodle Soup	vegan chicken	2	cup
Gentle Chicken Noodle Soup	dried basil	1	tsp
Gentle Chicken Noodle Soup	thyme	0.5	tsp
Gentle Chicken Noodle Soup	gluten free pasta	12	oz
Gentle Chicken Noodle Soup	vegetable broth	64	oz
Gentle Chicken Noodle Soup	bay leaf	2	\N
Gentle Chicken Noodle Soup	parsley	1	tbsp
Gentle Chicken Noodle Soup	lemon juice	1	tbsp
Chicken Noodle Soup	yellow onion	1	\N
Chicken Noodle Soup	carrot	3	\N
Chicken Noodle Soup	celery	3	\N
Chicken Noodle Soup	garlic	4	cloves
Chicken Noodle Soup	chicken	2.5	lb
Chicken Noodle Soup	dried basil	1	tsp
Chicken Noodle Soup	thyme	0.5	tsp
Chicken Noodle Soup	egg noodles	12	oz
Chicken Noodle Soup	chicken broth	64	oz
Chicken Noodle Soup	bay leaf	2	\N
Chicken Noodle Soup	parsley	1	tbsp
Chicken Noodle Soup	lemon juice	1	tbsp
Potato Soup	bacon	1	lb
Potato Soup	butter	3	tbsp
Potato Soup	yellow onion	1	\N
Potato Soup	garlic	3	cloves
Potato Soup	flour	0.33	cup
Potato Soup	golden potato	2.5	lb
Potato Soup	chicken broth	4	cup
Potato Soup	milk	2	cup
Potato Soup	heavy cream	0.66	cup
Potato Soup	salt	1.5	tsp
Potato Soup	pepper	1	tsp
Potato Soup	ancho chili powder	0.25	tsp
Potato Soup	sour cream	0.66	cup
Potato Soup	shredded cheese	1	\N
Potato Soup	chives	1	\N
Gentle Potato Soup	butter	3	tbsp
Gentle Potato Soup	yellow onion	1	\N
Gentle Potato Soup	garlic	3	cloves
Gentle Potato Soup	flour	0.33	cup
Gentle Potato Soup	golden potato	2.5	lb
Gentle Potato Soup	vegetable broth	4	cup
Gentle Potato Soup	milk	2	cup
Gentle Potato Soup	heavy cream	0.66	cup
Gentle Potato Soup	salt	1.5	tsp
Gentle Potato Soup	pepper	1	tsp
Gentle Potato Soup	ancho chili powder	0.25	tsp
Gentle Potato Soup	sour cream	0.66	cup
Gentle Potato Soup	shredded cheese	1	\N
Gentle Potato Soup	chives	1	\N
Chili	olive oil	1	tbsp
Chili	yellow onion	1	\N
Chili	ground beef	1	lb
Chili	chili powder	2.5	tbsp
Chili	cumin	2	tbsp
Chili	sugar	2	tbsp
Chili	tomato paste	2	tbsp
Chili	garlic powder	1	tbsp
Chili	salt	1.5	tsp
Chili	pepper	0.5	tsp
Chili	cayenne pepper	0.25	tsp
Chili	beef broth	1.5	cup
Chili	diced tomatoes	15	oz
Chili	kidney beans	16	oz
Chili	tomato sauce	8	oz
Chili	sour cream	\N	\N
Chili	cheddar cheese	\N	\N
Chili	chives	\N	\N
Gentle Chili	olive oil	1	tbsp
Gentle Chili	yellow onion	1	\N
Gentle Chili	black beans	15	oz
Gentle Chili	chili powder	2.5	tbsp
Gentle Chili	cumin	2	tbsp
Gentle Chili	sugar	2	tbsp
Gentle Chili	tomato paste	2	tbsp
Gentle Chili	garlic powder	1	tbsp
Gentle Chili	salt	1.5	tsp
Gentle Chili	pepper	0.5	tsp
Gentle Chili	cayenne pepper	0.25	tsp
Gentle Chili	vegetable broth	1.5	cup
Gentle Chili	diced tomatoes	15	oz
Gentle Chili	kidney beans	16	oz
Gentle Chili	tomato sauce	8	oz
Gentle Chili	carrot	2	\N
Gentle Chili	sour cream	\N	\N
Gentle Chili	cheddar cheese	\N	\N
Gentle Chili	chives	\N	\N
Omelet	egg	3	\N
Omelet	shredded cheese	0.25	cup
Omelet	heavy cream	2	tbsp
Omelet	seasonal vegetable	\N	\N
Chili	carrot	2	\N
Tomato Soup	plum tomatoes	3	lb
Tomato Soup	olive oil	0.25	cup
Tomato Soup	salt	1	tbsp
Tomato Soup	pepper	1.5	tsp
Tomato Soup	spinach	0.25	cup
Tomato Soup	yellow onion	2	cup
Tomato Soup	garlic	6	clove
Tomato Soup	unsalted butter	2	tbsp
Tomato Soup	red pepper flakes	0.25	tsp
Tomato Soup	canned plum tomatoes	28	oz
Tomato Soup	basil leaves	4	cups
Tomato Soup	thyme	1	tsp
Tomato Soup	vegetable broth	1	qt
Goulash	olive oil	1	tbsp
Goulash	onion	1	cup
Goulash	ground beef	1	lb
Goulash	garlic	3	tsp
Goulash	tomato sauce	15	oz
Goulash	canned tomatoes	15	oz
Goulash	beef broth	3	cup
Goulash	worcestershire sauce	3	tbsp
Goulash	salt	2	tsp
Goulash	italian seasoning	2	tbsp
Goulash	bay leaves	3	\N
Goulash	macaroni	2	cup
Goulash	cheddar cheese	1	cup
Bagel	bagel	1	\N
Bagel	cream cheese	1	tbsp
Annies Delight	mac and cheese	1	\N
Annies Delight	sausage	1	\N
Annies Delight	mushrooms	1	\N
Garlic Bread	french bread	1	loaf
Garlic Bread	garlic spread	1	tbsp
Butternut Squash Soup	onion	1	\N
Butternut Squash Soup	garlic	3	clove
Butternut Squash Soup	butternut squash	1	\N
Butternut Squash Soup	vegetable broth	4	cup
Butternut Squash Soup	heavy whipping cream	0.5	cup
Baked Mac And Cheese	pasta	1	lb
Baked Mac And Cheese	butter	0.5	cup
Baked Mac And Cheese	flour	0.5	cup
Baked Mac And Cheese	whole milk	1.5	cup
Baked Mac And Cheese	half and half	2.5	cup
Baked Mac And Cheese	cheddar cheese	4	cup
Baked Mac And Cheese	gruyere cheese	2	cup
Baked Mac And Cheese	paprika	0.25	tsp
Baked Mac And Cheese	bacon	0.5	lb
Baked Mac And Cheese	green chile	1	\N
Baked Mac And Cheese	panko	0.5	cup
Tomato Basil Tortellini Soup	onion	1	\N
Tomato Basil Tortellini Soup	carrots	3	\N
Tomato Basil Tortellini Soup	garlic	5	clove
Tomato Basil Tortellini Soup	crushed tomatoes	84	oz
Tomato Basil Tortellini Soup	vegetable broth	32	oz
Tomato Basil Tortellini Soup	sugar	1	tbsp
Tomato Basil Tortellini Soup	dried basil	1	tsp
Tomato Basil Tortellini Soup	bay leaf	1	\N
Tomato Basil Tortellini Soup	tortellini	27	oz
Tomato Basil Tortellini Soup	half and half	0.75	cup
Tomato Basil Tortellini Soup	shredded parmesan	\N	\N
Tomato Basil Tortellini Soup	basil	\N	\N
Mexican Shrimp Bisque	onion	1	\N
Mexican Shrimp Bisque	olive oil	3	tbsp
Mexican Shrimp Bisque	garlic	2	clove
Mexican Shrimp Bisque	flour	1	tbsp
Mexican Shrimp Bisque	heavy whipping cream	0.5	cup
Mexican Shrimp Bisque	vegetable broth	0.5	cup
Mexican Shrimp Bisque	chili powder	1	tbsp
Mexican Shrimp Bisque	cumin	0.5	tsp
Mexican Shrimp Bisque	coriander	0.5	tsp
Mexican Shrimp Bisque	shrimp	0.5	lb
Mexican Shrimp Bisque	sour cream	0.5	cup
Mexican Shrimp Bisque	cilantro	\N	\N
Mexican Shrimp Bisque	avocado	\N	\N
Seafood Chowder	bacon	12	slices
Seafood Chowder	onion	1	\N
Seafood Chowder	turnip	6	\N
Seafood Chowder	broth	6	cup
Seafood Chowder	thyme	1.5	tsp
Seafood Chowder	salt	2.5	tsp
Seafood Chowder	pepper	1.5	tsp
Seafood Chowder	heavy whipping cream	4.5	cup
Seafood Chowder	white fish	2	lb
Seafood Chowder	clams	2	can
Seafood Chowder	butter	6	tbsp
Seafood Chowder	coconut flour	2.25	tsp
Seafood Chowder	corn	2	can
Seafood Chowder	shrimp	0.5	lb
Seafood Chowder	leek	3	cup
Cajun Oyster Crackers	oyster crackers	9	oz
Cajun Oyster Crackers	cajun seasoning	2	tsp
Cajun Oyster Crackers	butter	2	tbsp
Southwest Carrot Roast	carrot	1.5	lb
Southwest Carrot Roast	olive oil	1.5	tbsp
Southwest Carrot Roast	salt	0.25	tsp
Southwest Carrot Roast	cumin	0.25	tsp
Southwest Carrot Roast	chili powder	0.125	tsp
Honey Mustard Salmon	salmon	1	lb
Honey Mustard Salmon	olive oil	1	tbsp
Honey Mustard Salmon	lemon	2	\N
Honey Mustard Salmon	honey	4	tbsp
Honey Mustard Salmon	dijon	4	tsp
Honey Mustard Salmon	garlic	2	clove
Honey Mustard Salmon	smoked paprika	0.25	tsp
Garlic Parm Green Beans	green beans	1.5	lb
Garlic Parm Green Beans	garlic	4	clove
Garlic Parm Green Beans	red pepper flakes	0.5	tsp
Garlic Parm Green Beans	parmesan	\N	\N
Lemon Pesto Chicken Thighs	chicken thighs	\N	\N
Lemon Pesto Chicken Thighs	butter	4	tbsp
Lemon Pesto Chicken Thighs	lemon	2	\N
Lemon Pesto Chicken Thighs	shallot	2	\N
Lemon Pesto Chicken Thighs	pesto	3	tbsp
Lemon Pesto Chicken Thighs	orzo	1.5	cup
Lemon Pesto Chicken Thighs	chicken stock	2.5	cup
Salmon Burgers	salmon	1.5	lb
Salmon Burgers	panko bread crumbs	0.5	cup
Salmon Burgers	egg	1	\N
Salmon Burgers	parsley	3	tbsp
Salmon Burgers	salt	0.5	tsp
Salmon Burgers	garlic powder	0.5	tsp
Salmon Burgers	lemon juice	1	tbsp
Salmon Burgers	mayo	0.5	cup
Salmon Burgers	garlic	1	clove
Salmon Burgers	dill weed	0.5	tsp
Salmon Burgers	dijon mustard	0.5	tsp
Salmon Burgers	lemon pepper	1	pinch
Feta Broccoli	broccoli	2	heads
Feta Broccoli	lemon juice	1	tbsp
Feta Broccoli	feta	0.5	cup
Feta Broccoli	dill weed	1	tsp
Honey Cinnamon Sweet Potatoes	sweet potatoes	4	\N
Honey Cinnamon Sweet Potatoes	olive oil	0.25	cup
Honey Cinnamon Sweet Potatoes	honey	0.25	cup
Honey Cinnamon Sweet Potatoes	cinnamon	2	tsp
Honey Cinnamon Sweet Potatoes	salt	\N	\N
Honey Cinnamon Sweet Potatoes	pepper	\N	\N
Worlds Best Lasagna	sweet italian sausage	1	lb
Worlds Best Lasagna	lean ground beef	0.75	lb
Worlds Best Lasagna	onion	1	\N
Worlds Best Lasagna	garlic	2	cloves
Worlds Best Lasagna	crushed tomatoes	28	oz
Worlds Best Lasagna	tomato paste	12	oz
Worlds Best Lasagna	tomato sauce	13	oz
Worlds Best Lasagna	sugar	2	tbsp
Worlds Best Lasagna	dried basil	1.5	tsp
Worlds Best Lasagna	fennel seeds	0.5	tsp
Worlds Best Lasagna	italian seasoning	1	tsp
Worlds Best Lasagna	salt	1.5	tsp
Worlds Best Lasagna	black pepper	0.25	tsp
Worlds Best Lasagna	parsley	4	tbsp
Worlds Best Lasagna	lasagna noodles	12	\N
Worlds Best Lasagna	ricotta	16	oz
Worlds Best Lasagna	egg	1	\N
Worlds Best Lasagna	mozzarella	0.75	lb
Worlds Best Lasagna	parmesan	0.75	cup
Chickpea Chorizo Tostadas	chickpeas	30	oz
Chickpea Chorizo Tostadas	lime	2	\N
Chickpea Chorizo Tostadas	greek yogurt	0.5	cup
Chickpea Chorizo Tostadas	cilantro	1	\N
Chickpea Chorizo Tostadas	chorizo	16	oz
Chickpea Chorizo Tostadas	tostada shells	8	\N
Black Beans	black beans	16	oz
Black Beans	onion	1	\N
Black Beans	garlic	1	clove
Black Beans	cilantro	1	tbsp
Black Beans	cayenne pepper	0.25	tsp
Roasted Mexican Zucchini	zucchini	4	\N
Roasted Mexican Zucchini	olive oil	1.5	tbsp
Roasted Mexican Zucchini	chili powder	0.5	tsp
Roasted Mexican Zucchini	garlic powder	0.5	tsp
Roasted Mexican Zucchini	kosher salt	1	tsp
Roasted Mexican Zucchini	pepper	0.5	tsp
Roasted Mexican Zucchini	onion powder	0.5	tsp
Roasted Mexican Zucchini	cumin	0.25	tsp
Roasted Mexican Zucchini	cotija	2	tbsp
Roasted Mexican Zucchini	lime	1	\N
Roasted Mexican Zucchini	cilantro	2	tbsp
Bagel	smoked salmon	4	oz
Miso Butter Roasted Brocolli	olive oil	2	tbsp
Miso Butter Roasted Brocolli	kosher salt	1	\N
Miso Butter Roasted Brocolli	unsalted butter	1	tbsp
Miso Butter Roasted Brocolli	lime	1	\N
Miso Butter Roasted Brocolli	white miso	2	tsp
Gong Bao Shrimp	shrimp	1	lb
Gong Bao Shrimp	shaoxing wine	2	tbsp
Gong Bao Shrimp	corn starch	2	tsp
Gong Bao Shrimp	salt	\N	\N
Gong Bao Shrimp	sugar	2	tbsp
Gong Bao Shrimp	soy sauce	3	tbsp
Gong Bao Shrimp	balsamic vinegar	2	tbsp
Gong Bao Shrimp	sichuan peppercorns	0.5	tsp
Gong Bao Shrimp	bell pepper	1	\N
Gong Bao Shrimp	garlic	5	cloves
Gong Bao Shrimp	neutral oil	0.25	cup
Gong Bao Shrimp	peanuts	0.5	cup
Gong Bao Shrimp	dried red chiles	0.75	cup
Gong Bao Shrimp	scallions	3	\N
Miso Butter Roasted Brocolli	broccoli	1	lb
Caramelized Shallot Pasta	olive oil	0.25	cup
Caramelized Shallot Pasta	shallot	6	\N
Caramelized Shallot Pasta	garlic	5	cloves
Caramelized Shallot Pasta	red pepper flakes	1	tsp
Caramelized Shallot Pasta	salt	\N	\N
Caramelized Shallot Pasta	pepper	\N	\N
Caramelized Shallot Pasta	anchovy fillets	2	oz
Caramelized Shallot Pasta	tomato paste	6	oz
Caramelized Shallot Pasta	pasta	10	oz
Caramelized Shallot Pasta	parsley	1	cup
\.


--
-- PostgreSQL database dump complete
--

