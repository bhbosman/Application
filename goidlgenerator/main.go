package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/Extensions"
	"github.com/bhbosman/Application/goidlgenerator/Publish"
	_ "github.com/bhbosman/Application/goidlgenerator/Publish/json"
	_ "github.com/bhbosman/Application/goidlgenerator/Publish/publishGo"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//go:generate goyacc -o yacc/idl.go -p "IdlExpr" yacc/idl.y
//go:generate go clean
//go:generate go build
//go:generate go install

func main() {
	//inputFile := flag.String("in", "", "inputFile")
	outputFile := flag.String("out", "", "outputFile")
	outputType := flag.String("outType", "json", "outputType")
	verbose := flag.Bool("verbose", false, "verbose")
	packageName := flag.String("packageName", "default", "packageName")
	defaultTypes := flag.Bool("defaultTypes", false, "Generate the built in types for the type in use")
	typesToUseAsString := flag.String("typesToUse", "IdlNative", `values "IdlNative" or "Mitch"`)
	showHelp := flag.Bool("help", false, "Show this")

	flag.Parse()

	if *showHelp {
		_, _ = fmt.Fprintln(os.Stderr, "goidlgenerator params args\n\twhere params is")
		flag.PrintDefaults()

		fmt.Println(">>>>>>>>>>>>>")

		if *verbose {
			flag.VisitAll(func(i *flag.Flag) {
				fmt.Println(i.Name, i.Value)
			})
		}
		fmt.Println("<<<<<<<<<<<<<")

		os.Exit(1)
		return

	}

	outStream, err := GetOutput(*outputFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Output file could not be created.\n")
		os.Exit(2)
		return
	}

	defer func() {
		_ = outStream.Close()
	}()

	typesInUse, err := yacc.StringToIDlBaseType(*typesToUseAsString)
	if err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("invalid type in use (%v).", typesToUseAsString))
		os.Exit(7)
	}

	var definitionDeclarations []interfaces.IDefinitionDeclaration = nil
	if !*defaultTypes {
		definitionDeclarations = ReadFromSource(*verbose, typesInUse)

	} else {
		definitionDeclarations, err = typesInUse.DefaultDecls()
		if err != nil {
			_, _ = os.Stderr.Write([]byte(fmt.Sprintf("Error: %v\n", err.Error())))
			os.Exit(10)
		}

	}
	publisher, err := Publish.HasOutputType(Publish.ToOutputType(*outputType))
	if err != nil {
		_, _ = os.Stderr.Write([]byte(fmt.Sprintf("Error: %v\n", err.Error())))
		os.Exit(5)
	}

	err = publisher.Export(
		typesInUse,
		Extensions.DefaultTypeValueHelper(),
		Publish.ExportParams{
			OutputStream:  outStream,
			PackageName:   *packageName,
			DeclaredTypes: definitionDeclarations})
	if err != nil {
		_, _ = os.Stderr.Write([]byte(fmt.Sprintf("Error: %v\n", err.Error())))
		os.Exit(6)
	}

}

func ReadFromSource(verbose bool, typesInUse interfaces.IBaseTypeInformation) []interfaces.IDefinitionDeclaration {
	inStream, err := GetInput(flag.Args())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "No input file. See help:\n")
		os.Exit(3)
		return nil
	}
	for _, closer := range inStream {
		defer func(closer io.Closer) {
			_ = closer.Close()
		}(closer)
	}

	_, definitionDeclarations, err := func(inStream []io.ReadCloser) ([]int, []interfaces.IDefinitionDeclaration, error) {
		result := make([]int, len(inStream), len(inStream))

		definitionDeclarations := make([]interfaces.IDefinitionDeclaration, 0, 16)

		idlExprContext := yacc.NewIdlExprContext()
		for i, reader := range inStream {
			lex, err := yacc.NewIdlExprLex(
				bufio.NewReader(reader),
				typesInUse,
				yacc.NewIdlExprLexParams{
					IdlExprContext: idlExprContext,
					Verbose:        verbose,
				})
			if err != nil {
				continue
			}
			resultInstance := yacc.IdlExprParse(lex)
			if resultInstance == 0 {
				for _, definitionDeclaration := range idlExprContext.GetSpecification() {
					definitionDeclarations = append(definitionDeclarations, definitionDeclaration)
				}
			}
			result[i] = resultInstance
		}

		return result, definitionDeclarations, nil

	}(inStream)
	if err != nil {
		os.Exit(4)
	}

	return definitionDeclarations
}

type WriterNoCloser struct {
	Writer io.Writer
}

func (self *WriterNoCloser) Close() error {
	return nil
}

func (self *WriterNoCloser) Write(p []byte) (n int, err error) {
	return self.Writer.Write(p)
}

func GetOutput(fileName string) (io.WriteCloser, error) {
	if fileName == "" {
		wc := &WriterNoCloser{
			Writer: os.Stdout,
		}
		return wc, nil
	}
	return os.Create(fileName)
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GetInput(args []string) ([]io.ReadCloser, error) {
	if len(args) > 0 {
		list := make([]io.ReadCloser, 0, len(args))
		for _, item := range args {
			if Exists(item) {
				file, err := os.Open(item)
				if err != nil {
					return nil, err
				}
				list = append(list, file)
				continue
			} else {
				readerCloser := ioutil.NopCloser(strings.NewReader(item))
				list = append(list, readerCloser)
			}
		}
		return list, nil
	}
	return nil, errors.New("no input files or strings")
}
