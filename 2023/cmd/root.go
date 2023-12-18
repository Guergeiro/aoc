package cmd

import (
	"fmt"
	"os"

	"github.com/guergeiro/aoc/2023/cmd/aoc01"
	"github.com/guergeiro/aoc/2023/cmd/aoc02"
	"github.com/guergeiro/aoc/2023/cmd/aoc03"
	"github.com/guergeiro/aoc/2023/cmd/aoc04"
	"github.com/guergeiro/aoc/2023/cmd/aoc05"
	"github.com/guergeiro/aoc/2023/cmd/aoc06"
	"github.com/guergeiro/aoc/2023/cmd/aoc07"
	"github.com/guergeiro/aoc/2023/cmd/aoc08"
	"github.com/guergeiro/aoc/2023/cmd/aoc09"
	"github.com/spf13/cobra"
)

func Execute() {

	rootCmd := &cobra.Command{
		Use:   "aoc2023",
		Short: "AoC 2023 solutions",
		Run: func(cmd *cobra.Command, args []string) {
			aoc01.Solve()
			aoc02.Solve()
			aoc03.Solve()
			aoc04.Solve()
			aoc05.Solve()
			aoc06.Solve()
			aoc07.Solve()
			aoc08.Solve()
			aoc09.Solve()
		},
	}
	rootCmd.AddCommand(
		&cobra.Command{
			Use: "1",
			Run: func(cmd *cobra.Command, args []string) {
				aoc01.Solve()
			},
		},
		&cobra.Command{
			Use: "2",
			Run: func(cmd *cobra.Command, args []string) {
				aoc02.Solve()
			},
		},
		&cobra.Command{
			Use: "3",
			Run: func(cmd *cobra.Command, args []string) {
				aoc03.Solve()
			},
		},
		&cobra.Command{
			Use: "4",
			Run: func(cmd *cobra.Command, args []string) {
				aoc04.Solve()
			},
		},
		&cobra.Command{
			Use: "5",
			Run: func(cmd *cobra.Command, args []string) {
				aoc05.Solve()
			},
		},
		&cobra.Command{
			Use: "6",
			Run: func(cmd *cobra.Command, args []string) {
				aoc06.Solve()
			},
		},
		&cobra.Command{
			Use: "7",
			Run: func(cmd *cobra.Command, args []string) {
				aoc07.Solve()
			},
		},
		&cobra.Command{
			Use: "8",
			Run: func(cmd *cobra.Command, args []string) {
				aoc08.Solve()
			},
		},
		&cobra.Command{
			Use: "9",
			Run: func(cmd *cobra.Command, args []string) {
				aoc09.Solve()
			},
		},
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
