package main

import (
	"github.com/marlin5555/dddgen/pkg/code-gen/util"
	"github.com/marlin5555/dddgen/pkg/template/factory"
)

func (g *Generator) generateFactory(genDir string) {
	w, f := getWriter(genDir, util.FactoryPath, util.Persistence)
	factory.WritePersistence(w, g.pkg.fullEntities, g.pkg.common)
	w.Flush()
	f.Close()

	w, f = getWriter(genDir, util.FactoryPath, util.RepoPackage)
	factory.WriteRepo(w, g.pkg.groups.GroupMap, g.pkg.common)
	w.Flush()
	f.Close()
}
