package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	cli "github.com/spf13/cobra"

	"go.pkg.littleman.co/oilr/internal/boilr"
	"go.pkg.littleman.co/oilr/internal/util/osutil"
	"go.pkg.littleman.co/oilr/internal/util/tlog"
	"go.pkg.littleman.co/oilr/internal/util/validate"
)

// Delete contains the cli-command for deleting templates.
var Delete = &cli.Command{
	Use:   "delete <template-tag>",
	Short: "Delete a project template from the template registry",
	Run: func(c *cli.Command, args []string) {
		MustValidateVarArgs(args, validate.Argument{"template-path", validate.AlphanumericExt})

		MustValidateTemplateDir()

		for _, templateName := range args {
			targetDir := filepath.Join(boilr.Configuration.TemplateDirPath, templateName)

			switch exists, err := osutil.DirExists(targetDir); {
			case err != nil:
				tlog.Error(fmt.Sprintf("delete: %s", err))
				continue
			case !exists:
				tlog.Error(fmt.Sprintf("Template %v doesn't exist", templateName))
				continue
			}

			if err := os.RemoveAll(targetDir); err != nil {
				tlog.Error(fmt.Sprintf("delete: %v", err))
				continue
			}

			tlog.Success(fmt.Sprintf("Successfully deleted the template %v", templateName))
		}
	},
}