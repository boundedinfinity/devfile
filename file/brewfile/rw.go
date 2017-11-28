package brewfile

import (
    "fmt"
	"bufio"
	"os"
	"strings"
	"github.com/pkg/errors"
)

func (this *BrewFileProcessor) write() error {
    f, err := os.Create(this.Path)

    if err != nil {
        return err
    }

    defer f.Close()

    w := bufio.NewWriter(f)

    defer w.Flush()

    for _, l := range this.File.Lines {
        var t string

        if l.Name == BrewfileLineType_Comment {
            t = l.Value
        } else {
            if l.Comment == "" {
                t = fmt.Sprintf("%s \"%s\"", l.Name, l.Value)
            } else {
                t = fmt.Sprintf("%s \"%s\" # %s", l.Name, l.Value, l.Comment)
            }
        }

        _, err := w.WriteString(fmt.Sprintf("%s\n", t))

        if err != nil {
            return err
        }
    }

    return nil
}

func (this *BrewFileProcessor) lex() error {
	if this.Debug {
		this.Logger.Printf("Processing: %s", this.Path)
	}

	if _, err := os.Stat(this.Path); os.IsNotExist(err) {
		return  nil
	}

	file, err := os.Open(this.Path)

	if err != nil {
		return err
	}

	defer file.Close()

	lines := make([]BrewFileLine, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		token := scanner.Text()

		if strings.HasPrefix(token, string(BrewfileLineType_Comment)) {
			if this.Debug {
				this.Logger.Println("Processed comment")
			}

			lines = append(lines, BrewFileLine{
				Name:  BrewfileLineType_Comment,
				Value: token,
			})

			this.File.Comments += 1
			continue
		}

		if strings.HasPrefix(token, string(BrewfileLineType_Brew)) || strings.HasPrefix(token, string(BrewfileLineType_Tap)) || strings.HasPrefix(token, string(BrewfileLineType_Cask)) {
			line, err := this.lexLine(token)

			if err != nil {
				return err
			}

			lines = append(lines, line)

			if this.Debug {
				this.Logger.Printf("Processed line: %s\n", token)
			}

			this.File.Actions += 1
			continue
		}

		this.File.Ignored += 1

		if this.Debug {
			this.Logger.Printf("Ignored line: %s\n", token)
		}
	}

	this.File.Lines = lines
	return nil
}


func (this *BrewFileProcessor) lexLine(line string) (BrewFileLine, error) {
	reader := strings.NewReader(line)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	tokens := make([]string, 0)

	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	if len(tokens) < 2 {
		return BrewFileLine{}, errors.Errorf("error parsing [%s]", line)
	}

    l := &BrewFileLine{
        Name: String2BrewfileLineType(tokens[0]),
        Value: strings.Replace(tokens[1], "\"", "", -1),
    }

    if len(tokens) > 2 {
        l.Comment = strings.Join(tokens[3:], " ")
    }

	return *l, nil
}
