package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mrinalxdev/pdf-go/cmd"
)

 func main () {
     if len(os.Args) < 2 {
        fmt.Println("Please provide a pdf file path as an argument")
        os.Exit(1)
    }

    pdfPath := os.Args[1]
    absPath, _ := filepath.Abs(pdfPath)

    if err := cmd.CheckFileExists(absPath); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    text, err := cmd.ExtractText(absPath)
    if err != nil {
        fmt.Printf("Error extracting text : %v\n", err)
        os.Exit(1)
    }

    parsedDoc := cmd.ParseText(text)
    cmd.RenderParsedText(parsedDoc)
 }