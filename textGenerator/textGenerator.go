package textgenerator

import (
	"bufio"
	"fmt"
	"image/png"
	"os"
	"strings"

	gim "github.com/ozankasikci/go-image-merge"
)

func Generate() string {
	// Scanner to read STDinput
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//grid from de go-image-merge package
	grids := []*gim.Grid{}
	// get each char from the input and append the addres of the iamge to the grid

	for _, char := range scanner.Text() {
		ch := &gim.Grid{ImageFilePath: "../images/" + strings.ToUpper(string(char)) + ".png"}
		grids = append(grids, ch)
	}

	// merge the grid
	rgba, err := gim.New(grids, len(scanner.Text()), 1).Merge()
	if err != nil {
		fmt.Println("error:", err)
	}

	// save the output to jpg or png
	file, err := os.Create("../images/output.png")
	if err != nil {
		fmt.Println("Error creating the png file ", err)
	}
	err = png.Encode(file, rgba)
	if err != nil {
		fmt.Println("Error writing the output file: ", err)
	}
	return scanner.Text()
}
