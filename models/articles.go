package models

type Article struct {
	Title   string
	Image   string
	Article string
}

var Articles = map[string][]Article{
	"Perfume": {
		{
			Title:   "Hugo Boss Infinite",
			Image:   "https://images.unsplash.com/photo-1638551442447-085a2d42918f?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			Article: "Introducing Hugo Boss Infinite, a fragrance that embodies timeless sophistication and modern elegance. Crafted for the confident man who values quality and refinement, Hugo Boss Infinite is a captivating blend of invigorating freshness and warm sensuality. Its top notes of refreshing mandarin and energizing apple create a vibrant opening, while a heart of spicy cinnamon and aromatic sage adds depth and complexity. The base notes of patchouli and rosemary lend a lingering allure, making Hugo Boss Infinite a fragrance that leaves a lasting impression. Designed to complement the dynamic lifestyle of the modern man, Hugo Boss Infinite is a versatile scent that transitions seamlessly from day to night. Whether worn at the office or during an evening out, its sophisticated aroma exudes confidence and sophistication. Embrace the infinite possibilities that Hugo Boss Infinite offers and experience a fragrance that embodies the essence of timeless masculinity.",
		},
		{
			Title:   "Gentleman Givenchy",
			Image:   "https://images.unsplash.com/photo-1588514912908-8f5891714f8d?q=80&w=2069&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			Article: "Gentleman Givenchy is a fragrance that embodies the essence of sophistication and masculinity. With its timeless blend of aromatic notes, it has become a staple choice for modern men who appreciate classic elegance with a contemporary twist. The fragrance opens with fresh and vibrant top notes of bergamot, pear, and cardamom, creating an invigorating and memorable first impression. As it settles, the heart notes of lavender and geranium add a refined and aromatic touch, while the base notes of patchouli and leather provide depth and warmth, leaving a lasting and memorable trail. This iconic scent is not just a fragrance; it's a statement of style and refinement. Whether worn during the day for a touch of understated elegance or in the evening for a sophisticated aura, Gentleman Givenchy is a versatile choice for the modern gentleman who values timeless classics and contemporary flair.",
		},
	},
}
