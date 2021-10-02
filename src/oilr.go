package main

import (
  "fmt"
  
  "go.pkg.littleman.co/oilr/internal/boilr"
  "go.pkg.littleman.co/oilr/internal/cmd"
  "go.pkg.littleman.co/oilr/internal/util/exit"
  "go.pkg.littleman.co/oilr/internal/util/osutil"
)

func main() {
  if exists, err := osutil.DirExists(boilr.Configuration.TemplateDirPath); ! exists {
    if err := osutil.CreateDirs(boilr.Configuration.TemplateDirPath); err != nil {
      exit.Error(fmt.Errorf("Tried to initialise your template directory, but it has failed: %s", err))
    }
  } else if err != nil {
    exit.Error(fmt.Errorf("Failed to init: %s", err))
  }

  cmd.Run()
}

