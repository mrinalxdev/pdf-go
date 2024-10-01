package cmd

import "fmt"


func RenderParsedText(doc DocumentStructure){
    fmt.Printf("Title : %s\n\n", doc.Title)
    fmt.Println("MetaData :")
    for key, value := range doc.MetaData {
        fmt.Printf(" %s: %s\n", key, value)
    }
    fmt.Println()

    for _, section := range doc.Sections {
        fmt.Printf("Section : %s\n", section.Header)
        for _, paragraph := range section.Paragraphs {
            fmt.Printf(" %s\n", paragraph)
        }

        fmt.Println()
    }
}
