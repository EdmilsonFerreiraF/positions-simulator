package route

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string
	ClientID  string
	Positions []Position
}

type Position struct {
	Lat  float64
	Long float64
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route id not informed")
	}

	f, err := os.Open("destinations/" + r.ID + ".txt")

	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")

		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}

		long, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}

		r.Positions = append(r.Positions, Position{
			Lat:  lat,
			Long: long,
		})
	}

	return nil
}
