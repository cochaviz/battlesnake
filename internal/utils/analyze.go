package utils

import (
	"github.com/gocarina/gocsv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"log"
	"os"
)

type Measurement struct {
	Turn   int `csv:"Turn"`
	Length int `csv:"Length"`
}

func (measurement Measurement) AppendToFile(filename string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err, "Trying to create file")

		f, err = os.Create(filename)
		if err != nil {
			return err
		}
		return gocsv.MarshalFile([]Measurement{measurement}, f)
	}
	if err != nil {
		log.Println(err)
	}
	return gocsv.MarshalCSVWithoutHeaders([]Measurement{measurement}, gocsv.DefaultCSVWriter(f))
}

func ReadFromFile(filename string) ([]Measurement, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 664)

	if err != nil {
		return []Measurement{}, err
	}
	var measurements = []Measurement{}
	return measurements, gocsv.UnmarshalFile(f, &measurements)
}

func PlotMeasurements(inputFilename string, outputFilename string) {
	measurements, err := ReadFromFile(inputFilename)

	if err != nil {
		log.Print(err)
		return
	}
	p := plot.New()

	var values plotter.XYs

	for _, measurement := range measurements {
		values = append(values, plotter.XY{Y: float64(measurement.Length), X: float64(measurement.Turn)})
	}
	p.Title.Text = "Length over time"

	scatter, err := plotter.NewLine(values)

	if err != nil {
		log.Print(err)
		return
	}
	p.Add(scatter)
	err = p.Save(500, 300, outputFilename)

	if err != nil {
		log.Print(err)
	}
}
