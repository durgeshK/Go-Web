package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type MenuItem struct {
	Name        string
	Description string
	Price       float64
}

type RestaurantMenu struct {
	RestaurantName string
	Menu           map[string][]MenuItem
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	rajdhaniRestaurant := new(RestaurantMenu)
	rajdhaniRestaurant.RestaurantName = "Rajdhani Restaurant"
	rajdhaniRestaurant.Menu = make(map[string][]MenuItem)

	rajdhaniRestaurant.Menu["Breakfast"] = []MenuItem{
		MenuItem{"Aloo Paratha", "Indian Bread stuffed with Aloo masala", 100},
		MenuItem{"Puri Sabji", "Indian Bread deep fried with Aloo gravy", 150},
	}

	rajdhaniRestaurant.Menu["Lunch"] = []MenuItem{
		MenuItem{"Rajma Chawal", "Kidney Beans gravy with Rice", 100},
		MenuItem{"Indian Thali", "Chapatti, Rice, Daal", 150},
	}

	rajdhaniRestaurant.Menu["Dinner"] = []MenuItem{
		MenuItem{"Indian Thali", "Chapatti, Rice, Daal", 150},
		MenuItem{"Paneer Thali", "Indian Thali with Paneer gravy", 200},
	}

	err := tpl.Execute(os.Stdout, rajdhaniRestaurant)
	if err != nil {
		log.Fatalln(err)
	}

}
