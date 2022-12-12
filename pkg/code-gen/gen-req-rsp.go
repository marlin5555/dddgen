package main

import (
	"github.com/marlin5555/dddgen/pkg/code-gen/util"
	"github.com/marlin5555/dddgen/pkg/template/entity"
)

func (g *Generator) generateReq(genDir string) {
	for gName, group := range g.pkg.groups.GroupMap {
		w, f := getWriter(genDir, util.EntityReqPath, gName)
		entity.WriteReq(w, group)
		w.Flush()
		f.Close()
	}
}

func (g *Generator) generateRsp(genDir string) {
	commonGroup := g.pkg.groups.ShareColumn.GetGroup()
	w, f := getWriter(genDir, util.EntityRspPath, commonGroup.Name)
	entity.WriteCommonRsp(w, commonGroup)
	w.Flush()
	f.Close()

	for gName, group := range g.pkg.groups.GroupMap {
		w, f := getWriter(genDir, util.EntityRspPath, gName)
		entity.WriteRsp(w, group)
		w.Flush()
		f.Close()
	}
}

func (g *Generator) generateEntity(genDir string) {
	commonGroup := g.pkg.groups.ShareColumn.GetGroup()
	w, f := getWriter(genDir, util.EntityPath, commonGroup.Name)
	entity.WriteCommonEntity(w, commonGroup)
	w.Flush()
	f.Close()

	for gName, group := range g.pkg.groups.GroupMap {
		w, f := getWriter(genDir, util.EntityPath, gName)
		entity.WriteEntity(w, group)
		w.Flush()
		f.Close()
	}
}
