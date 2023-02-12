package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/briandowns/spinner"
	"os"
	"time"
)

func drawTable() {
	b := table.NewWriter()
	b.SetOutputMirror(os.Stdout)
	b.AppendHeader(table.Row{"АКТИВЫ", "тыс.руб.", "штуки", "ПАССИВЫ", "тыс.руб"})
	b.AppendRows([]table.Row{
		{"Станки", CurrentAssets.Machine.Money, CurrentAssets.Machine.Physical, "Собственные средства", CurrentLiability.Equity},
		{"Сырье", CurrentAssets.Raw.Money, CurrentAssets.Raw.Physical, "Займы", CurrentLiability.Debt},
		{"Товар", CurrentAssets.Product.Money, CurrentAssets.Product.Physical},
		{"Деньги", CurrentAssets.Money},
	})
	b.AppendSeparator()

	b.AppendFooter(table.Row{"АКТИВЫ", assetsCalc(), "", "ПАССИВЫ", equityCalc()})

	//Визуал
	b.SetStyle(table.StyleColoredYellowWhiteOnBlack)
	b.Render()
}

func showSpinner() {
	fmt.Println()
	s := spinner.New(spinner.CharSets[33], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	time.Sleep(1 * time.Second)                                  // Run for some time to simulate work
	s.Stop()
}
