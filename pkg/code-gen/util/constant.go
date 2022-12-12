package util

import "path"

const (
	FUZZY  = "FUZZY"
	Fuzzy  = "Fuzzy"
	BATCH  = "BATCH"
	EXACT  = "EXACT"
	UPSERT = "UPSERT"
	INSERT = "INSERT"
	UPDATE = "UPDATE"
	SIMPLE = "SIMPLE"
	DETAIL = "DETAIL"

	ROLE   = "ROLE" // 关系中的角色，用于外键引用
	Module = "github.com/marlin5555/dddgen/"
)

const (
	REF = "REF"
	FER = "FER"
	// 1 - 1
	MUSTA = "MUST_A"
	TSUMA = "TSUM_A"
	MAYA  = "MAY_A"
	YAMA  = "YAM_A"
	// 1 - n
	MUSTS = "MUST_S"
	TSUMS = "TSUM_S"
	MAYS  = "MAY_S"
	YAMS  = "YAM_S"
	// m - n
	SMAYMUSTS = "S_MAY_MUST_S"
	SMUSTMAYS = "S_MUST_MAY_S"
	SMAYS     = "S_MAY_S"
	SYAMS     = "S_YAM_S"
	SMUSTS    = "S_MUST_S"
	STSUMS    = "S_TSUM_S"
)

const (
	ID    = "ID"
	AUTO  = "AUTO"
	GIVEN = "GIVEN"
	NAMED = "NAMED"
)

var (
	pkg           = "pkg"
	LogPackage    = "log"
	LogPath       = path.Join(pkg, LogPackage)
	ConfigPackage = "conf"
	ConfigPath    = path.Join(pkg, ConfigPackage)

	internal = "internal"

	InfraPath         = path.Join(internal, "infra")
	Persistence       = "persistence"
	PersistencePath   = path.Join(InfraPath, Persistence)
	PoPackage         = "po"
	PoPath            = path.Join(PersistencePath, PoPackage)
	PoToEntityPackage = "potoentity"
	PoToEntityPath    = path.Join(PersistencePath, PoToEntityPackage)
	SqlPackage        = "db"
	SqlPath           = path.Join(PersistencePath, SqlPackage)
	CommonFile        = "common"
	UtilFile          = "util"
	BuilderFile       = "builder"
	ConstantFile      = "constant"

	DomainPath        = path.Join(internal, "domain")
	EntityPackage     = "entity"
	EntityPath        = path.Join(DomainPath, EntityPackage)     // entity interface and getter
	EntityBasePackage = "base"                                   // package: interface
	EntityBasePath    = path.Join(EntityPath, EntityBasePackage) // interface
	EntityZeroPackage = "zero"                                   // package: struct impl interface
	EntityZeroPath    = path.Join(EntityPath, EntityZeroPackage) // struct impl interface
	EntityReqPackage  = "req"                                    // package: req
	EntityReqPath     = path.Join(EntityPath, EntityReqPackage)  // req interface
	EntityRspPackage  = "rsp"                                    // Package: rsp
	EntityRspPath     = path.Join(EntityPath, EntityRspPackage)  // rsp interface

	PortsPath            = path.Join(DomainPath, "ports")
	PortsInfraPath       = path.Join(PortsPath, "infra")
	PersistencePortsPath = path.Join(PortsInfraPath, Persistence)
	RepoPackage          = "repo"
	RepoPortsPath        = path.Join(PortsPath, RepoPackage)
	RepoPath             = path.Join(DomainPath, RepoPackage)

	FactoryPackage = "factory"
	FactoryPath    = path.Join(internal, FactoryPackage)
)

const poPath = "internal/dao/storage/po"
const sqlPath = "internal/dao/storage/sql"
const poToEntity = "internal/dao/convertor/entity"
const poEntityBuilder = "EntityBuilder"
const idKey = "id"
const idType = "uint64"

const entityReqPath = "internal/domain/entity/req" // request interface
const entityRspPath = "internal/domain/entity/rsp" // response interface

const portsRepoPath = "internal/domain/ports/repo"       // repo interface
const portsServicePath = "internal/domain/ports/service" // service interface

const servicePath = "internal/domain/service" // service impl
const repoPath = "internal/domain/repo"       // repo impl
