package main

import (
	"github.com/marlin5555/dddgen/pkg/code-gen/util"
	persist "github.com/marlin5555/dddgen/pkg/template/persistence"
)

func (g *Generator) generatePersistenceConstant(genDir string) {
	w, f := getWriter(genDir, util.PersistencePath, util.ConstantFile)
	persist.WriteConstant(w, g.pkg.fullEntities, g.pkg.common)
	w.Flush()
	f.Close()
}

func (g *Generator) generatePO(genDir string) {
	for gName, group := range g.pkg.groups.GroupMap {
		w, f := getWriter(genDir, util.PoPath, gName)
		persist.WritePO(w, group)
		w.Flush()
		f.Close()
	}
}

func (g *Generator) generateSQL(genDir string) {
	for eName, entity := range g.pkg.tables {
		w, f := getWriter(genDir, util.SqlPath, eName)
		persist.WriteSQL(w, entity)
		w.Flush()
		f.Close()
	}
	w, f := getWriter(genDir, util.SqlPath, util.CommonFile)
	persist.WriteCommonSQL(w, g.pkg.common)
	w.Flush()
	f.Close()
}

func (g *Generator) generatePersistencePorts(genDir string) {
	for gName, group := range g.pkg.groups.GroupMap {
		w, f := getWriter(genDir, util.PersistencePortsPath, gName)
		persist.WritePersistencePorts(w, group)
		w.Flush()
		f.Close()
	}
	w, f := getWriter(genDir, util.PersistencePortsPath, util.CommonFile)
	persist.WritePersistencePortsCommon(w, g.pkg.common)
	w.Flush()
	f.Close()
}

func (g *Generator) generatePOToEntity(genDir string) {
	for gName, group := range g.pkg.groups.GroupMap {
		w, f := getWriter(genDir, util.PoToEntityPath, gName)
		persist.WritePOToEntity(w, group)
		w.Flush()
		f.Close()
	}
	w, f := getWriter(genDir, util.PoToEntityPath, util.BuilderFile)
	persist.WriteBuilder(w, g.pkg.fullEntities, g.pkg.common)
	w.Flush()
	f.Close()

	w, f = getWriter(genDir, util.PoToEntityPath, util.CommonFile)
	persist.WriteBuilderCommon(w, g.pkg.common)
	w.Flush()
	f.Close()

}
