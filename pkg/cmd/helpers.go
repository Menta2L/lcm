package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"strings"
)

var FileExtensions = []string{".json", ".yaml", ".yml"}
var InputExtensions = append(FileExtensions, "stdin")

type FilenameOptions struct {
	Filenames []string
	Kustomize string
	Recursive bool
}

func (o *FilenameOptions) validate() []error {
	var errs []error
	if len(o.Filenames) > 0 && len(o.Kustomize) > 0 {
		errs = append(errs, fmt.Errorf("only one of -f or -k can be specified"))

	}
	if len(o.Kustomize) > 0 && o.Recursive {
		errs = append(errs, fmt.Errorf("the -k flag can't be used with -f or -R"))

	}
	return errs

}

func (o *FilenameOptions) RequireFilenameOrKustomize() error {
	if len(o.Filenames) == 0 && len(o.Kustomize) == 0 {
		return fmt.Errorf("must specify one of -f and -k")

	}
	return nil

}

func AddFilenameOptionFlags(cmd *cobra.Command, options *FilenameOptions, usage string) {
	AddJsonFilenameFlag(cmd.Flags(), &options.Filenames, "Filename, directory, or URL to files "+usage)
	cmd.Flags().BoolVarP(&options.Recursive, "recursive", "R", options.Recursive, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")

}
func AddJsonFilenameFlag(flags *pflag.FlagSet, value *[]string, usage string) {
	flags.StringSliceVarP(value, "filename", "f", *value, usage)
	annotations := make([]string, 0, len(FileExtensions))
	for _, ext := range FileExtensions {
		annotations = append(annotations, strings.TrimLeft(ext, "."))

	}
	flags.SetAnnotation("filename", cobra.BashCompFilenameExt, annotations)

}
