package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"unicode"
)

type Options struct {
	From      string
	To        string
	Offset    uint64
	Limit     int64
	BlockSize int
	Conv      []string
}

func ParseFlags() (*Options, error) {
	var opts Options

	flag.StringVar(&opts.From, "from", "stdin", "path to the source file, if not specified, use stdin")
	flag.StringVar(&opts.To, "to", "stdout", "path to the copy, if not specified, use stdout")
	flag.Uint64Var(&opts.Offset, "offset", 0, "number of bytes to skip at the beginning of input")
	flag.Int64Var(&opts.Limit, "limit", -1, "maximum number of bytes to copy from input, -1 means all")
	flag.IntVar(&opts.BlockSize, "block-size", 4096, "size of one block in bytes when reading and writing")
	conv := flag.String("conv", "", "comma-separated list of text transformations to apply")
	flag.Parse()

	if *conv != "" {
		convList := strings.Split(*conv, ",")
		for _, c := range convList {
			switch c {
			case "upper_case", "lower_case", "trim_spaces":
				opts.Conv = append(opts.Conv, c)
			default:
				return nil, fmt.Errorf("unknown conversion: %s", c)
			}
		}
		if opts.has("upper_case") && opts.has("lower_case") {
			return nil, fmt.Errorf("cannot convert to both upper and lower case")
		}
	}

	return &opts, nil
}

func (opts *Options) has(str string) bool {
	for i := range opts.Conv {
		if opts.Conv[i] == str {
			return true
		}
	}
	return false
}

func main() {
	opts, err := ParseFlags()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "can not parse flags:", err)
		os.Exit(1)
	}

	var from io.Reader
	if opts.From == "stdin" {
		from = os.Stdin
	} else {
		f, err := os.Stat(opts.From)
		if os.IsNotExist(err) {
			_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		if uint64(f.Size()) < opts.Offset {
			_, _ = fmt.Fprintln(os.Stderr, "Error:", "offset is greater than input size")
			os.Exit(1)
		}
		file, err := os.Open(opts.From)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		from = file
	}

	var to io.Writer
	if opts.To == "stdout" {
		to = os.Stdout
	} else {
		_, err := os.Stat(opts.To)
		if os.IsNotExist(err) {
			file, err := os.Create(opts.To)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			to = file
		} else {
			_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	}

	if opts.Offset > 0 {
		buf := make([]byte, opts.Offset)
		if _, err := io.ReadFull(from, buf); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	}

	if opts.Limit < 0 {
		opts.Limit = math.MaxInt64
	}
	if opts.Limit > 0 {
		from = io.LimitReader(from, opts.Limit)
	}

	if opts.Conv != nil {
		bytes, err := io.ReadAll(from)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		text := string(bytes)
		for _, t := range opts.Conv {
			switch t {
			case "upper_case":
				text = strings.ToUpper(text)
			case "lower_case":
				text = strings.ToLower(text)
			case "trim_spaces":
				text = strings.TrimFunc(text, unicode.IsSpace)
			}
		}
		if _, err := to.Write([]byte(text)); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	} else {
		if _, err := io.Copy(to, from); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}

}
