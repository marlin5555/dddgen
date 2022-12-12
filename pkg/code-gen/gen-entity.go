package main

import (
	"github.com/marlin5555/dddgen/pkg/code-gen/util"
	tmplentity "github.com/marlin5555/dddgen/pkg/template/entity"
)

func (g *Generator) generateEntityInterface(genDir string) {
	for dName, d := range g.pkg.groups.GroupMap {
		w, f := getWriter(genDir, util.EntityBasePath, dName)
		tmplentity.WriteInterface(w, d)
		w.Flush()
		f.Close()
	}
	commonGroup := g.pkg.groups.ShareColumn.GetGroup()
	w, f := getWriter(genDir, util.EntityBasePath, commonGroup.Name)
	tmplentity.WriteCommonInterface(w, commonGroup)
	w.Flush()
	f.Close()
	w, f = getWriter(genDir, util.EntityBasePath, util.UtilFile)
	tmplentity.WriteUtilInterface(w, commonGroup)
	w.Flush()
	f.Close()
}

func (g *Generator) generateEntityZero(genDir string) {
	for dName, d := range g.pkg.groups.GroupMap {
		w, f := getWriter(genDir, util.EntityZeroPath, dName)
		tmplentity.WriteZero(w, d)
		w.Flush()
		f.Close()
	}
	commonGroup := g.pkg.groups.ShareColumn.GetGroup()
	w, f := getWriter(genDir, util.EntityZeroPath, commonGroup.Name)
	tmplentity.WriteCommonZero(w, commonGroup)
	w.Flush()
	f.Close()
}
