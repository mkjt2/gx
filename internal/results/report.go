package results

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/montanaflynn/stats"
	"github.com/rodaine/table"
)

type Metric[V int | float64] struct {
	Points []V
}

func Average(m *Metric[float64]) (float64, error) {
	return stats.Mean(m.Points)
}

func Stdev(m *Metric[float64]) (float64, error) {
	return stats.StandardDeviation(m.Points)
}

func MetricFormatted(includeStdev bool, f func() *Metric[float64]) string {
	av, err := Average(f())
	if err != nil {
		panic(err)
	}
	if !includeStdev {
		return fmt.Sprintf("%.2f", av)
	}
	stdev, err := Stdev(f())
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%.2f (%.3f)", av, stdev)
}

func (m *Metric[V]) Add(v V) {
	m.Points = append(m.Points, v)
}

type Result struct {
	Algorithm                string
	Package                  string
	OriginalSize             Metric[int]
	CompressedSize           Metric[int]
	CompressionTimeSeconds   Metric[float64]
	DecompressionTimeSeconds Metric[float64]
	Errors                   []string
}

func (r *Result) CompressionRatio() *Metric[float64] {
	m := Metric[float64]{}
	for i := range r.OriginalSize.Points {
		m.Add(float64(r.OriginalSize.Points[i]) / float64(r.CompressedSize.Points[i]))
	}
	return &m
}

func (r *Result) GetCompressionTimeSeconds() *Metric[float64] {
	return &r.CompressionTimeSeconds
}

func (r *Result) GetDecompressionTimeSeconds() *Metric[float64] {
	return &r.DecompressionTimeSeconds
}

func GetReportTable(results []*Result) table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	// TODO indicate stdev in brackets clearly
	tbl := table.New("Algo", "Package", "CX Ratio", "CX Time (s)", "DX Time (s)")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	// TODO make report prettier (e.g. numeric formatting, etc)

	for _, result := range results {
		tbl.AddRow(
			result.Algorithm,
			result.Package,
			fmt.Sprintf(MetricFormatted(false, result.CompressionRatio)),
			fmt.Sprintf(MetricFormatted(true, result.GetCompressionTimeSeconds)),
			fmt.Sprintf(MetricFormatted(true, result.GetDecompressionTimeSeconds)),
		)
	}
	return tbl
}
