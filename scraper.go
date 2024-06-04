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
	Title, Price, CreatedAt, Link string 
}


func ReformatPrice(price string) (string, error){
	// Reformat price. Delete all spaces and letters.
	if price == "договорная"{
		return price, nil
	} else{
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	cleanedPrice := re.ReplaceAllString(price, "")
	cleanedPrice = strings.Replace(cleanedPrice, " ", "", 4)
	
	s, _ := strconv.Atoi(cleanedPrice);
	
	if s >= 50000 {
		return cleanedPrice, nil 
	}
	return "", nil
	}
}	


func Scraper() {
	// Function for parsing page
	envFile, _ := godotenv.Read(".env")

	// TODO: Change page number for pagination. 
	var cards []OrderCard
	url := envFile["URL"]
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
				Link: fmt.Sprintf(market_url, link),
			}
			cards = append(cards, card)
		}
		// output. TODO: save in json or send to TGbot
		fmt.Println(cards) 
	})
	c.Visit(url)
}


func main() {
	Scraper()
}