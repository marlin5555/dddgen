package main

import (
	"github.com/marlin5555/dddgen/pkg/code-gen/util"
	"github.com/marlin5555/dddgen/pkg/template/repo"
)

func (g *Generator) generateRepoPorts(genDir string) {
	for gName, group := range g.pkg.groups.GroupMap {
		w, f := getWriter(genDir, util.RepoPortsPath, gName)
		repo.WriteRepoPorts(w, group)
		w.Flush()
		f.Close()
	}
}

func (g *Generator) generateRepo(genDir string) {
	for gName, group := range g.pkg.groups.GroupMap {
		w, f := getWriter(genDir, util.RepoPath, gName)
		repo.WriteRepo(w, group)
		w.Flush()
		f.Close()
	}
}
