package main 

import (
	"fmt"
	"strings"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/gocolly/colly"

)


type OrderCard struct {
	Title string  
	Price string
	CreatedAt string
	Link string 
}


func RemoveLetters(s string) string {
	// Takes a string and returns only numbers
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	delete_str := re.ReplaceAllString(s, "")
	return strings.Replace(delete_str, " ", "", 4)
}


func ReformatPrice(price string) (string, error){
	// Takes a price and deletes all spaces and letters.
	if price == "договорная"{
		return price, nil
	} else{
	cleanedPrice := RemoveLetters(price)
	s, _ := strconv.Atoi(cleanedPrice);
	
	if s >= 50000 {
		return cleanedPrice, nil 
	}
	return "", nil
	}
}	
  

func Scraper(page_num int) ([]OrderCard, error){
	// Function for parsing pages.

	var cards []OrderCard
	envFile, _ := godotenv.Read(".env")

	for page := 1; page < page_num +1 ; page++ {

		url := envFile["URL"] + "page=" + strconv.Itoa(page)
		market_url := envFile["MARKET_URL"]
		c := colly.NewCollector()
		c.UserAgent = envFile["USER_AGENT"]
		
		c.OnHTML(".task_list", func(h *colly.HTMLElement) {
			title := h.ChildText(".task__title")
			price := h.ChildText(".task__price")
			created_at := h.ChildText(".params__published-at")
			link := market_url + h.ChildAttr(".task__title a", "href")
			formattedPrice, err := ReformatPrice(price)

			if err != nil {
				fmt.Println("Error reformatting price:", err)
				return
			}
			if formattedPrice != ""{
				card := OrderCard{
					Title: title,
					Price: price,
					CreatedAt: created_at,
					Link: link,
				}
				cards = append(cards, card)
			}
		})
		c.Visit(url)
		}
	return cards, nil
}

