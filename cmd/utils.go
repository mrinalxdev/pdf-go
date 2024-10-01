 package cmd

 import (
     "fmt"
     "os"
 )

 func CheckFileExists(filePath string) error {
    _, err := os.Stat(filePath)
     if os.IsNotExist(err){
         return fmt.Errorf("file does not exist : %s", filePath)
     }

     return err
 }
