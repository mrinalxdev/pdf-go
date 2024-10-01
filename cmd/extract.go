package cmd

import (
	"github.com/ledongthuc/pdf"
)

func ExtractText(path string) (string, error) {
    f, r, err := pdf.Open(path)

    if err != nil {
        return " ", err
    }
    defer f.Close()

    var text string
    for i := 1; i <= r.NumPage(); i++ {
        p := r.Page(i)
        if p.V.IsNull(){
            continue
        }
        t, err := p.GetPlainText(nil)
        if err != nil {
            return " ", err
        }
        text += t
    }
    return text, nil
}
