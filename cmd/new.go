package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"goignore/helpers"
	"io/ioutil"
	"log"
	"strings"
)

type IgnoreFile struct {
	Techs []string
	Content *bytes.Buffer
}

var newCmd = &cobra.Command{
	Use: "new",
	Short: "Generates a new .gitignore, defaults to the local directory",
	Long: "Creates a new .gitignore file and configures it based on your config and args",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Building ignore file for: %v", args)
		ignoreFile := NewIgnoreFile(args)
		ignoreFile.GenerateContent()
		ignoreFile.WriteOut()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func NewIgnoreFile(techs []string) *IgnoreFile {
	b := new(bytes.Buffer)
	ignoreFile := &IgnoreFile{
		techs,
		b,
	}

	return ignoreFile
}

func (ign *IgnoreFile) GenerateContent() {
	for _, tech := range ign.Techs {
		tech = strings.Title(tech)
		scrapedContent, err := helpers.GetIgnore(tech)
		log.Print(string(scrapedContent))
		if err != nil {
			msg := fmt.Sprintf("Error getting ignore config for '$s'!\n %v", err)
			exit(msg)
		}
		header := fmt.Sprintf("\n#----Ignore rules for: %s\n", tech)
		ign.Content.Write([]byte(header))
		ign.Content.Write(scrapedContent)
	}
}

func (ign *IgnoreFile) WriteOut() {
	fileContent := ign.Content.Bytes()
	log.Print("####")
	log.Printf(string(fileContent))
	err := ioutil.WriteFile(".gitignore", fileContent, 0664)
	if err != nil {
		msg := fmt.Sprintf("Failed to write to file!\n%v", err)
		exit(msg)
	}
}

