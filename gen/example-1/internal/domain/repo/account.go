// Package repo Code generated, DO NOT EDIT.
package repo

import (
	"context"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/req"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/zero"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/ports/infra/persistence"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/ports/repo"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/po"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/potoentity"
)

// AccountRepo ...
type AccountRepo struct {
	accountDAO      persistence.AccountDAO
	passportDAO     persistence.PassportDAO
	secretDAO       persistence.SecretDAO
	techRelationDAO persistence.TechRelationDAO
}

// NewAccountRepo ...
func NewAccountRepo(accountDAO persistence.AccountDAO, passportDAO persistence.PassportDAO, secretDAO persistence.SecretDAO, techRelationDAO persistence.TechRelationDAO) repo.AccountRepository {
	return &AccountRepo{
		accountDAO:      accountDAO,
		passportDAO:     passportDAO,
		secretDAO:       secretDAO,
		techRelationDAO: techRelationDAO,
	}
}

// CreateAccount create Account
func (r *AccountRepo) CreateAccount(ctx context.Context, req req.CreateAccountReq) (string, error) {
	return r.accountDAO.Create(ctx, req)
}

// UpdateAccount update Account
func (r *AccountRepo) UpdateAccount(ctx context.Context, req req.UpdateAccountReq) error {
	return r.accountDAO.Update(ctx, req)
}

// DeleteAccount delete Account
func (r *AccountRepo) DeleteAccount(ctx context.Context, req req.DeleteAccountReq) error {
	return r.accountDAO.Delete(ctx, req)
}

// GetAccounts general get Account method
func (r *AccountRepo) GetAccounts(ctx context.Context, rq req.GetAccountsReq) (entity.Accounts, uint32, error) {
	// 做主表查询，获取到主表记录
	result, total, err := r.accountDAO.Get(ctx, struct {
		req.GetAccountsNoForeign
		base.PageOrderOperator // 分页查询信息
	}{
		GetAccountsNoForeign: rq,
		PageOrderOperator:    rq,
	})
	// 关联表查询
	builder := potoentity.NewBuilder(rq,
		potoentity.WithAccounts(base.ToStringKeyMap(result,
			func(i interface{}) string { return i.(base.AccountID).GetAccountID() })),
		potoentity.WithFields(r.accountDAO.BaseFields()))
	builder.Build()
	return builder.GetAccounts(), total, err
}

func needAccountIDs(req req.GetAccountsPure) bool {
	if !base.IsZero(req.GetFuzzyAccountName()) {
		return true
	}
	if !base.IsZero(req.GetFuzzyAccountNickname()) {
		return true
	}
	if !base.IsZero(req.GetAccountID()) {
		return true
	}
	if !base.IsZero(req.GetAccountIDs()) {
		return true
	}
	if !base.IsZero(req.GetAccountName()) {
		return true
	}
	if !base.IsZero(req.GetAccountNames()) {
		return true
	}
	if !base.IsZero(req.GetAccountNickname()) {
		return true
	}
	return false
}

func getAccountIDsByRef(ctx context.Context, dao persistence.AccountDAO, pure req.GetAccountsPure) (base.AccountIDs, bool, error) {
	if !needAccountIDs(pure) {
		return zero.AccountIDs{}, false, nil
	}
	var accounts po.Accounts
	var err error
	accounts, _, err = dao.Get(ctx, struct {
		zero.PageOrderOperator
		req.GetAccountsPure
	}{GetAccountsPure: pure})
	if len(accounts.GetAccountIDs()) == 0 || err != nil {
		return zero.AccountIDs{}, true, nil
	}
	return zero.AccountIDs{Value: accounts.GetAccountIDs()}, false, nil
}

func getAccountsByIDsIfRelated(ctx context.Context, dao persistence.AccountDAO, ids base.AccountIDs, fields base.Fields) (po.Accounts, error) {
	if !dao.IsRelated(fields) {
		return nil, nil
	}
	return getAccountsByIDs(ctx, dao, ids)
}

func getAccountsByIDs(ctx context.Context, dao persistence.AccountDAO, ids base.AccountIDs) (po.Accounts, error) {
	var accounts po.Accounts
	var err error
	if accounts, _, err = dao.Get(ctx, struct {
		zero.FuzzyAccountName
		zero.FuzzyAccountNickname
		zero.AccountID
		base.AccountIDs
		zero.AccountName
		zero.AccountNames
		zero.AccountNickname
		zero.PageOrderOperator
	}{AccountIDs: ids}); err != nil {
		return nil, err
	}
	return accounts, err
}

// CreatePassport create Passport
func (r *AccountRepo) CreatePassport(ctx context.Context, req req.CreatePassportReq) (string, error) {
	return r.passportDAO.Create(ctx, req)
}

// UpdatePassport update Passport
func (r *AccountRepo) UpdatePassport(ctx context.Context, req req.UpdatePassportReq) error {
	return r.passportDAO.Update(ctx, req)
}

// DeletePassport delete Passport
func (r *AccountRepo) DeletePassport(ctx context.Context, req req.DeletePassportReq) error {
	return r.passportDAO.Delete(ctx, req)
}

// GetPassports general get Passport method
func (r *AccountRepo) GetPassports(ctx context.Context, rq req.GetPassportsReq) (entity.Passports, uint32, error) {
	var err error
	var isStop bool
	var accountIDs base.AccountIDs
	if accountIDs, isStop, err = getAccountIDsByRef(ctx, r.accountDAO, rq); isStop {
		return nil, 0, err
	}
	// 做主表查询，获取到主表记录
	result, total, err := r.passportDAO.Get(ctx, struct {
		req.GetPassportsNoForeign
		base.PageOrderOperator // 分页查询信息
		zero.AccountID
		base.AccountIDs
	}{
		GetPassportsNoForeign: rq,
		PageOrderOperator:     rq,
		AccountIDs:            zero.AccountIDs{Value: accountIDs.GetAccountIDs()},
	})
	// 关联表查询
	var accounts po.Accounts
	if accounts, err = getAccountsByIDsIfRelated(ctx, r.accountDAO, result, rq); err != nil {
		return nil, 0, err
	}
	builder := potoentity.NewBuilder(rq,
		potoentity.WithPassports(base.ToStringKeyMap(result,
			func(i interface{}) string { return i.(base.PassportID).GetPassportID() })),
		potoentity.WithAccounts(base.ToStringKeyMap(accounts,
			func(i interface{}) string { return i.(base.AccountID).GetAccountID() })),
		potoentity.WithFields(r.passportDAO.BaseFields()))
	builder.Build()
	return builder.GetPassports(), total, err
}

func needPassportIDs(req req.GetPassportsPure) bool {
	if !base.IsZero(req.GetAccountID()) {
		return true
	}
	if !base.IsZero(req.GetAccountIDs()) {
		return true
	}
	if !base.IsZero(req.GetFuzzyPassportPostalAddress()) {
		return true
	}
	if !base.IsZero(req.GetPassportID()) {
		return true
	}
	if !base.IsZero(req.GetPassportIDs()) {
		return true
	}
	if !base.IsZero(req.GetPassportNationality()) {
		return true
	}
	if !base.IsZero(req.GetPassportNationalitys()) {
		return true
	}
	return false
}

func getPassportIDsByRef(ctx context.Context, dao persistence.PassportDAO, pure req.GetPassportsPure) (base.PassportIDs, bool, error) {
	if !needPassportIDs(pure) {
		return zero.PassportIDs{}, false, nil
	}
	var passports po.Passports
	var err error
	passports, _, err = dao.Get(ctx, struct {
		zero.PageOrderOperator
		req.GetPassportsPure
	}{GetPassportsPure: pure})
	if len(passports.GetPassportIDs()) == 0 || err != nil {
		return zero.PassportIDs{}, true, nil
	}
	return zero.PassportIDs{Value: passports.GetPassportIDs()}, false, nil
}

func getPassportsByIDsIfRelated(ctx context.Context, dao persistence.PassportDAO, ids base.PassportIDs, fields base.Fields) (po.Passports, error) {
	if !dao.IsRelated(fields) {
		return nil, nil
	}
	return getPassportsByIDs(ctx, dao, ids)
}

func getPassportsByIDs(ctx context.Context, dao persistence.PassportDAO, ids base.PassportIDs) (po.Passports, error) {
	var passports po.Passports
	var err error
	if passports, _, err = dao.Get(ctx, struct {
		zero.AccountID
		zero.AccountIDs
		zero.FuzzyPassportPostalAddress
		zero.PassportID
		base.PassportIDs
		zero.PassportNationality
		zero.PassportNationalitys
		zero.PageOrderOperator
	}{PassportIDs: ids}); err != nil {
		return nil, err
	}
	return passports, err
}

func getPassportsByAccountIDsIfRelated(ctx context.Context, dao persistence.PassportDAO, ids base.AccountIDs, fields base.Fields) (po.Passports, error) {
	if !dao.IsRelated(fields) {
		return nil, nil
	}
	return getPassportsByAccountIDs(ctx, dao, ids)
}

func getPassportsByAccountIDs(ctx context.Context, dao persistence.PassportDAO, ids base.AccountIDs) (po.Passports, error) {
	var passports po.Passports
	var err error
	if passports, _, err = dao.Get(ctx, struct {
		zero.AccountID
		base.AccountIDs
		zero.FuzzyPassportPostalAddress
		zero.PassportID
		zero.PassportIDs
		zero.PassportNationality
		zero.PassportNationalitys
		zero.PageOrderOperator
	}{AccountIDs: ids}); err != nil {
		return nil, err
	}
	return passports, err
}

// CreateSecret create Secret
func (r *AccountRepo) CreateSecret(ctx context.Context, req req.CreateSecretReq) (string, error) {
	return r.secretDAO.Create(ctx, req)
}

// UpdateSecret update Secret
func (r *AccountRepo) UpdateSecret(ctx context.Context, req req.UpdateSecretReq) error {
	return r.secretDAO.Update(ctx, req)
}

// DeleteSecret delete Secret
func (r *AccountRepo) DeleteSecret(ctx context.Context, req req.DeleteSecretReq) error {
	return r.secretDAO.Delete(ctx, req)
}

// GetSecrets general get Secret method
func (r *AccountRepo) GetSecrets(ctx context.Context, rq req.GetSecretsReq) (entity.Secrets, uint32, error) {
	var err error
	var isStop bool
	var accountIDs base.AccountIDs
	if accountIDs, isStop, err = getAccountIDsByRef(ctx, r.accountDAO, rq); isStop {
		return nil, 0, err
	}
	// 做主表查询，获取到主表记录
	result, total, err := r.secretDAO.Get(ctx, struct {
		req.GetSecretsNoForeign
		base.PageOrderOperator // 分页查询信息
		zero.AccountID
		base.AccountIDs
	}{
		GetSecretsNoForeign: rq,
		PageOrderOperator:   rq,
		AccountIDs:          zero.AccountIDs{Value: accountIDs.GetAccountIDs()},
	})
	// 关联表查询
	var accounts po.Accounts
	if accounts, err = getAccountsByIDsIfRelated(ctx, r.accountDAO, result, rq); err != nil {
		return nil, 0, err
	}
	builder := potoentity.NewBuilder(rq,
		potoentity.WithSecrets(base.ToStringKeyMap(result,
			func(i interface{}) string { return i.(base.SecretID).GetSecretID() })),
		potoentity.WithAccounts(base.ToStringKeyMap(accounts,
			func(i interface{}) string { return i.(base.AccountID).GetAccountID() })),
		potoentity.WithFields(r.secretDAO.BaseFields()))
	builder.Build()
	return builder.GetSecrets(), total, err
}

func needSecretIDs(req req.GetSecretsPure) bool {
	if !base.IsZero(req.GetAccountID()) {
		return true
	}
	if !base.IsZero(req.GetSecretID()) {
		return true
	}
	if !base.IsZero(req.GetSecretIDs()) {
		return true
	}
	return false
}

func getSecretIDsByRef(ctx context.Context, dao persistence.SecretDAO, pure req.GetSecretsPure) (base.SecretIDs, bool, error) {
	if !needSecretIDs(pure) {
		return zero.SecretIDs{}, false, nil
	}
	var secrets po.Secrets
	var err error
	secrets, _, err = dao.Get(ctx, struct {
		zero.PageOrderOperator
		req.GetSecretsPure
	}{GetSecretsPure: pure})
	if len(secrets.GetSecretIDs()) == 0 || err != nil {
		return zero.SecretIDs{}, true, nil
	}
	return zero.SecretIDs{Value: secrets.GetSecretIDs()}, false, nil
}

func getSecretsByIDsIfRelated(ctx context.Context, dao persistence.SecretDAO, ids base.SecretIDs, fields base.Fields) (po.Secrets, error) {
	if !dao.IsRelated(fields) {
		return nil, nil
	}
	return getSecretsByIDs(ctx, dao, ids)
}

func getSecretsByIDs(ctx context.Context, dao persistence.SecretDAO, ids base.SecretIDs) (po.Secrets, error) {
	var secrets po.Secrets
	var err error
	if secrets, _, err = dao.Get(ctx, struct {
		zero.AccountID
		zero.SecretID
		base.SecretIDs
		zero.PageOrderOperator
	}{SecretIDs: ids}); err != nil {
		return nil, err
	}
	return secrets, err
}

// CreateTechRelation create TechRelation
func (r *AccountRepo) CreateTechRelation(ctx context.Context, req req.CreateTechRelationReq) (string, error) {
	return r.techRelationDAO.Create(ctx, req)
}

// UpdateTechRelation update TechRelation
func (r *AccountRepo) UpdateTechRelation(ctx context.Context, req req.UpdateTechRelationReq) error {
	return r.techRelationDAO.Update(ctx, req)
}

// DeleteTechRelation delete TechRelation
func (r *AccountRepo) DeleteTechRelation(ctx context.Context, req req.DeleteTechRelationReq) error {
	return r.techRelationDAO.Delete(ctx, req)
}

// GetTechRelations general get TechRelation method
func (r *AccountRepo) GetTechRelations(ctx context.Context, rq req.GetTechRelationsReq) (entity.TechRelations, uint32, error) {
	var err error
	var isStop bool
	var studentIDs base.AccountIDs
	if studentIDs, isStop, err = getAccountIDsByRef(ctx, r.accountDAO, rq); isStop {
		return nil, 0, err
	}
	var teacherIDs base.AccountIDs
	if teacherIDs, isStop, err = getAccountIDsByRef(ctx, r.accountDAO, rq); isStop {
		return nil, 0, err
	}
	// 做主表查询，获取到主表记录
	result, total, err := r.techRelationDAO.Get(ctx, struct {
		req.GetTechRelationsNoForeign
		base.PageOrderOperator // 分页查询信息
		zero.StudentID
		base.StudentIDs
		zero.TeacherID
		base.TeacherIDs
	}{
		GetTechRelationsNoForeign: rq,
		PageOrderOperator:         rq,
		StudentIDs:                zero.StudentIDs{Value: studentIDs.GetAccountIDs()},
		TeacherIDs:                zero.TeacherIDs{Value: teacherIDs.GetAccountIDs()},
	})
	// 关联表查询
	var students po.Accounts
	if students, err = getAccountsByIDsIfRelated(ctx, r.accountDAO, base.StudentIDsToAccountIDs(result), rq); err != nil {
		return nil, 0, err
	}
	var teachers po.Accounts
	if teachers, err = getAccountsByIDsIfRelated(ctx, r.accountDAO, base.TeacherIDsToAccountIDs(result), rq); err != nil {
		return nil, 0, err
	}
	builder := potoentity.NewBuilder(rq,
		potoentity.WithTechRelations(base.ToStringKeyMap(result,
			func(i interface{}) string { return i.(base.TechRelationID).GetTechRelationID() })),
		potoentity.WithAccounts(base.ToStringKeyMap(students,
			func(i interface{}) string { return i.(base.AccountID).GetAccountID() })),
		potoentity.WithAccounts(base.ToStringKeyMap(teachers,
			func(i interface{}) string { return i.(base.AccountID).GetAccountID() })),
		potoentity.WithFields(r.techRelationDAO.BaseFields()))
	builder.Build()
	return builder.GetTechRelations(), total, err
}

func needTechRelationIDs(req req.GetTechRelationsPure) bool {
	if !base.IsZero(req.GetTechRelationID()) {
		return true
	}
	if !base.IsZero(req.GetTechRelationIDs()) {
		return true
	}
	return false
}

func getTechRelationIDsByRef(ctx context.Context, dao persistence.TechRelationDAO, pure req.GetTechRelationsPure) (base.TechRelationIDs, bool, error) {
	if !needTechRelationIDs(pure) {
		return zero.TechRelationIDs{}, false, nil
	}
	var techRelations po.TechRelations
	var err error
	techRelations, _, err = dao.Get(ctx, struct {
		zero.PageOrderOperator
		req.GetTechRelationsPure
	}{GetTechRelationsPure: pure})
	if len(techRelations.GetTechRelationIDs()) == 0 || err != nil {
		return zero.TechRelationIDs{}, true, nil
	}
	return zero.TechRelationIDs{Value: techRelations.GetTechRelationIDs()}, false, nil
}

func getTechRelationsByIDsIfRelated(ctx context.Context, dao persistence.TechRelationDAO, ids base.TechRelationIDs, fields base.Fields) (po.TechRelations, error) {
	if !dao.IsRelated(fields) {
		return nil, nil
	}
	return getTechRelationsByIDs(ctx, dao, ids)
}

func getTechRelationsByIDs(ctx context.Context, dao persistence.TechRelationDAO, ids base.TechRelationIDs) (po.TechRelations, error) {
	var techRelations po.TechRelations
	var err error
	if techRelations, _, err = dao.Get(ctx, struct {
		zero.TechRelationID
		base.TechRelationIDs
		zero.PageOrderOperator
	}{TechRelationIDs: ids}); err != nil {
		return nil, err
	}
	return techRelations, err
}
