package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
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

//go:generate goyacc -o yacc/idl.publishGo -p "IdlExpr" yacc/idl.y
//go:generate go clean
//go:generate go build
//go:generate go install
//go:generate rm -rf goidlgenerator

func main() {
	//inputFile := flag.String("in", "", "inputFile")
	outputFile := flag.String("out", "", "outputFile")
	outputType := flag.String("outType", "json", "outputType")
	verbose := flag.Bool("verbose", false, "verbose")
	packageName := flag.String("packageName", "default", "packageName")
	showHelp := flag.Bool("help", false, "Show this")

	flag.Parse()

	if *showHelp {
		_, _ = fmt.Fprintln(os.Stderr, "goidlgenerator params args\n\twhere params is")
		flag.PrintDefaults()
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

	inStream, err := GetInput(flag.Args())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "No input file. See help:\n")
		os.Exit(3)
		return
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
				idlExprContext,
				*verbose)
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
	jsonPublisher, err := Publish.HasOutputType(Publish.ToOutputType(*outputType))
	if err != nil {
		_, _ = os.Stderr.Write([]byte(fmt.Sprintf("Error: %v\n", err.Error())))
		os.Exit(5)
	}
	err = jsonPublisher.Export(outStream, *packageName, definitionDeclarations)
	if err != nil {
		_, _ = os.Stderr.Write([]byte(fmt.Sprintf("Error: %v\n", err.Error())))
		os.Exit(6)
	}

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
