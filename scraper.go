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


func GetPagesAmount(url string) []string {
	// Takes colly instance and url. Then just takes pages amount from website
	var amount []string
	envFile, _ := godotenv.Read(".env")
	c := colly.NewCollector()
	c.UserAgent = envFile["USER_AGENT"]

    c.OnHTML(".pagination", func(c *colly.HTMLElement) {
        req := c.ChildText(".pagination a")[4:]
		remove_str := RemoveLetters(req)
        amount = append(amount, remove_str)
    })
    c.Visit(url)
    c.Wait()
    return amount
  }

  
func Scraper() {
	// Function for parsing pages.
	envFile, _ := godotenv.Read(".env")
	var cards []OrderCard

	
	pagesAmount := GetPagesAmount(envFile["URL"])

	if len(pagesAmount) > 0 {
        maxPages, err := strconv.Atoi(pagesAmount[0])
        if err != nil {
            fmt.Println("Invalid page number:", pagesAmount[0])
            return
        }
	
	for page := 1; page < maxPages+1; page++ {
		fmt.Println(page)
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
					Link: fmt.Sprintf(market_url, link),
				}
				cards = append(cards, card)
			}
			// output. TODO: save in json or send to TGbot
			fmt.Println(cards) 
		})
		c.Visit(url)
		}
	
}}


func main() {
	Scraper()
}