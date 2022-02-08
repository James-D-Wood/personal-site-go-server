package valuesort

const InitialData = `
  [
	{
	  "title": "Unsorted",
	  "cards": [
		{ "body": "ACCEPTANCE", "details": "to be accepted as I am" },
		{ "body": "ACCURACY", "details": "to be correct in my opinions and beliefs" },
		{ "body": "ACHIEVEMENT", "details": "to have important accomplishments" },
		{ "body": "ADVENTURE", "details": "to have new and exciting experiences" },
		{ "body": "ART", "details": "to appreciate or express myself in art" },
		{ "body": "ATTRACTIVENESS", "details": "to be physically attractive" },
		{ "body": "AUTHORITY", "details": "to be in charge of others" },
		{ "body": "AUTONOMY", "details": "to be self-determined and independent" },
		{ "body": "BEAUTY", "details": "to appreciate beauty around me" },
		{
		  "body": "BELONGING",
		  "details": "to have a sense of belonging, being part of"
		},
		{ "body": "CARING", "details": "to take care of others" },
		{ "body": "CHALLENGE", "details": "to take on difficult tasks and problems" },
		{ "body": "COMFORT", "details": "to have a pleasant and comfortable life" },
		{
		  "body": "COMMITMENT",
		  "details": "to make enduring, meaningful commitments"
		},
		{ "body": "COMPASSION", "details": "to feel and act on concern for others" },
		{ "body": "COMPLEXITY", "details": "to embrace the intricacies of life" },
		{
		  "body": "COMPROMISE",
		  "details": "to be willing to give and take in reaching agreements"
		},
		{
		  "body": "CONTRIBUTION",
		  "details": "to make a lasting contribution in the world"
		},
		{ "body": "COOPERATION", "details": "to work collaboratively with others" },
		{
		  "body": "COURAGE",
		  "details": "to be brave and strong in the face of adversity"
		},
		{
		  "body": "COURTESY",
		  "details": "to be considerate and polite toward others"
		},
		{ "body": "CREATIVITY", "details": "to create new things or ideas" },
		{
		  "body": "CURIOSITY",
		  "details": "to seek out, experience, and learn new things"
		},
		{ "body": "DEPENDABILITY", "details": "to be reliable and trustworthy" },
		{
		  "body": "DILIGENCE",
		  "details": "to be thorough and conscientious in whatever I do"
		},
		{ "body": "DUTY", "details": "to carry out my duties and obligations" },
		{ "body": "ECOLOGY", "details": "to live in harmony with the environment" },
		{
		  "body": "EXCITEMENT",
		  "details": "to have a life full of thrills and stimulation"
		},
		{
		  "body": "FAITHFULNESS",
		  "details": "to be loyal and true in relationships"
		},
		{ "body": "FAME", "details": "to be known and recognized" },
		{
		  "body": "FAMILY",
		  "details":
			"to have a happy, loving family"
		},
		{ "body": "FITNESS", "details": "to be physically fit and strong" },
		{
		  "body": "FLEXIBILITY",
		  "details":
			"to adjust to new circumstances easily to be forgiving of others"
		},
		{ "body": "FORGIVENESS", "details": "to be forgiving of others" },
		{
		  "body": "FREEDOM",
		  "details": "to be free from undue restrictions and limitations"
		},
		{ "body": "FRIENDSHIP", "details": "to have close, supportive friends" },
		{ "body": "FUN", "details": "to play and have fun" },
		{ "body": "GENEROSITY", "details": "to give what I have to others" },
		{ "body": "GENUINENESS", "details": "to act in a manner that is true" },
		{ "body": "GOD\"S WILL", "details": "to seek and obey the will of God" },
		{ "body": "GRATITUDE", "details": "to be thankful and appreciative" },
		{ "body": "GROWTH", "details": "to keep changing and growing" },
		{ "body": "HEALTH", "details": "to be physically well and health" },
		{ "body": "HONESTY", "details": "to be honest and truthful" },
		{
		  "body": "HOPE",
		  "details": "to maintain a positive and optimistic outlook"
		},
		{ "body": "HUMILITY", "details": "to be modest and unassuming" },
		{ "body": "HUMOR", "details": "to see the humorous side of myself" },
		{ "body": "IMAGINATION", "details": "to have dreams and see possibilities" },
		{
		  "body": "INDEPENDENCE",
		  "details": "to be free from depending on others"
		},
		{ "body": "INDUSTRY", "details": "to work hard and well at my life tasks" },
		{ "body": "INNER PEACE", "details": "to experience personal peace" },
		{
		  "body": "INTEGRITY",
		  "details":
			"to live my daily life in a way that is consistent with my values"
		},
		{ "body": "INTELLIGENCE", "details": "to keep my mind sharp and active" },
		{
		  "body": "INTIMACY",
		  "details": "to share my innermost experiences with others"
		},
		{
		  "body": "JUSTICE",
		  "details": "to promote fair and equal treatment for all"
		},
		{
		  "body": "KNOWLEDGE",
		  "details": "to learn and contribute valuable knowledge"
		},
		{ "body": "LEADERSHIP", "details": "to inspire and guide others" },
		{ "body": "LEISURE", "details": "to take time to relax and enjoy" },
		{ "body": "LOVED", "details": "to be loved by those close to me" },
		{ "body": "LOVING", "details": "to give love to others" },
		{ "body": "MASTERY", "details": "to be competent in my everyday activities" },
		{
		  "body": "MINDFULNESS",
		  "details": "to live conscious and mindful of the present moment"
		},
		{
		  "body": "MODERATION",
		  "details": "to avoid excesses and find a middle ground"
		},
		{ "body": "MONOGAMY", "details": "to have one close, loving relationship" },
		{ "body": "MUSIC", "details": "to enjoy or express myself in music" },
		{
		  "body": "NON-CONFORMITY",
		  "details": "to question and challenge authority and norms"
		},
		{ "body": "NOVELTY", "details": "to have a life full of change and variety" },
		{ "body": "NURTURANCE", "details": "to encourage and support others" },
		{
		  "body": "OPENNESS",
		  "details": "to be open to new experiences, ideas, and options"
		},
		{ "body": "ORDER", "details": "to have a life that is well and organized" },
		{
		  "body": "PASSION",
		  "details": "to have deep feelings about ideas, activities, or people"
		},
		{ "body": "PATRIOTISM", "details": "to love, serve, and protect my country" },
		{ "body": "PLEASURE", "details": "to feel good" },
		{ "body": "POPULARITY", "details": "to be well-liked by many people" },
		{ "body": "POWER", "details": "to have control over others" },
		{
		  "body": "PRACTICALITY",
		  "details": "to focus on what is practical, prudent, and sensible"
		},
		{ "body": "PROTECT", "details": "to protect and keep safe those I love" },
		{ "body": "PROVIDE", "details": "to provide for and take care of my family" },
		{ "body": "PURPOSE", "details": "to have meaning and direction in my life" },
		{
		  "body": "RATIONALITY",
		  "details": "to be guided by reason, logic, and evidence"
		},
		{
		  "body": "REALISM",
		  "details": "to see and act realistically and practically"
		},
		{
		  "body": "RESPONSIBILITY",
		  "details": "to make and carry out responsible decisions"
		},
		{ "body": "RISK", "details": "to take risks and chances" },
		{ "body": "ROMANCE", "details": "to have intense, exciting love in my life" },
		{ "body": "SAFETY", "details": "to be safe and secure" },
		{ "body": "SELF-ACCEPTANCE", "details": "to accept myself as I am" },
		{ "body": "SELF-CONTROL", "details": "to be disciplined in my own actions" },
		{ "body": "SELF-ESTEEM", "details": "to feel good about myself" },
		{
		  "body": "SELF-KNOWLEDGE",
		  "details": "to have a deep and honest understanding of myself"
		},
		{ "body": "SERVICE", "details": "to be helpful and of service to others" },
		{
		  "body": "SEXUALITY",
		  "details": "to have an active and satisfying sex life"
		},
		{
		  "body": "SIMPLICITY",
		  "details": "to live life simply, with minimal needs"
		},
		{
		  "body": "SOLITUDE",
		  "details": "to have time and space where be apart from others"
		},
		{ "body": "SPIRITUALITY", "details": "to grow and mature spiritually" },
		{
		  "body": "STABILITY",
		  "details": "to have a life that stays fairly consistent"
		},
		{
		  "body": "TOLERANCE",
		  "details": "to accept and respect those who differ from me"
		},
		{
		  "body": "TRADITION",
		  "details": "to follow respected patterns of the past"
		},
		{ "body": "VIRTUE", "details": "to live a morally pure and excellent life" },
		{ "body": "WEALTH", "details": "to have plenty of money" },
		{ "body": "WORLD", "details": "to work to promote peace in the world" }
	  ]
	},
	{
	  "title": "Not Important",
	  "cards": []
	},
	{
	  "title": "Somewhat Important",
	  "cards": []
	},
	{
	  "title": "Important",
	  "cards": []
	},
	{
	  "title": "Very Important",
	  "cards": []
	},
	{
	  "title": "Most Important",
	  "cards": []
	}
  ]
`
