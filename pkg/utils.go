package pkg

import (
	"context"
	"math/rand"
	"time"

	"github.com/joaopedrosgs/loucore/ent/city"
)

var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomCityCoordinates() (int, int) {
	x, y := 0, 0
	exists := true
	for exists == true {
		x = randomGenerator.Intn(500)
		y = randomGenerator.Intn(500)
		exists, _ = client.City.Query().Where(city.XEQ(x), city.YEQ(y)).Exist(context.Background())
		if !exists {
			break
		}
	}
	return x, y
}
