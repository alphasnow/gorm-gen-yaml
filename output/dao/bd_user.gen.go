// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/we7coreteam/gorm-gen-yaml/output/entity"
)

func newUser(db *gorm.DB, opts ...gen.DOOption) user {
	_user := user{}

	_user.userDo.UseDB(db, opts...)
	_user.userDo.UseModel(&entity.User{})

	tableName := _user.userDo.TableName()
	_user.ALL = field.NewAsterisk(tableName)
	_user.ID = field.NewInt32(tableName, "id")
	_user.Username = field.NewString(tableName, "username")
	_user.Nickname = field.NewString(tableName, "nickname")
	_user.Mobile = field.NewString(tableName, "mobile")
	_user.Avatar = field.NewString(tableName, "avatar")
	_user.RealName = field.NewString(tableName, "real_name")
	_user.WxNumber = field.NewString(tableName, "wx_number")
	_user.LocationAddress = field.NewString(tableName, "location_address")
	_user.Address = field.NewString(tableName, "address")
	_user.Model = field.NewString(tableName, "model")
	_user.Lng = field.NewFloat64(tableName, "lng")
	_user.Lat = field.NewFloat64(tableName, "lat")
	_user.LoginTime = field.NewInt32(tableName, "login_time")
	_user.LoginCount = field.NewInt32(tableName, "login_count")
	_user.Gender = field.NewInt32(tableName, "gender")
	_user.CreateTime = field.NewInt32(tableName, "create_time")
	_user.UpdateTime = field.NewInt32(tableName, "update_time")
	_user.AlipayAccount = field.NewString(tableName, "alipay_account")
	_user.AlipayName = field.NewString(tableName, "alipay_name")
	_user.AlipayQrcode = field.NewString(tableName, "alipay_qrcode")
	_user.UserOauth = userHasOneUserOauth{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("UserOauth", "entity.UserOauth"),
	}

	_user.fillFieldMap()

	return _user
}

type user struct {
	userDo

	ALL             field.Asterisk
	ID              field.Int32
	Username        field.String  // 用户名
	Nickname        field.String  // 昵称
	Mobile          field.String  // 手机号
	Avatar          field.String  // 头像地址
	RealName        field.String  // 真实姓名
	WxNumber        field.String  // 微信号码
	LocationAddress field.String  // 定位地址
	Address         field.String  // 详细地址
	Model           field.String  // 手机型号
	Lng             field.Float64 // 经度
	Lat             field.Float64 // 纬度
	LoginTime       field.Int32   // 用户登录时间
	LoginCount      field.Int32   // 用户登录次数
	Gender          field.Int32   // 用户性别:0=未知,1=男性,2=女性
	CreateTime      field.Int32   // 添加时间
	UpdateTime      field.Int32   // 修改时间
	AlipayAccount   field.String  // 支付宝帐号
	AlipayName      field.String  // 支付宝名称
	AlipayQrcode    field.String  // 支付宝二维码
	UserOauth       userHasOneUserOauth

	fieldMap map[string]field.Expr
}

func (u user) Table(newTableName string) *user {
	u.userDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u user) As(alias string) *user {
	u.userDo.DO = *(u.userDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *user) updateTableName(table string) *user {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.Username = field.NewString(table, "username")
	u.Nickname = field.NewString(table, "nickname")
	u.Mobile = field.NewString(table, "mobile")
	u.Avatar = field.NewString(table, "avatar")
	u.RealName = field.NewString(table, "real_name")
	u.WxNumber = field.NewString(table, "wx_number")
	u.LocationAddress = field.NewString(table, "location_address")
	u.Address = field.NewString(table, "address")
	u.Model = field.NewString(table, "model")
	u.Lng = field.NewFloat64(table, "lng")
	u.Lat = field.NewFloat64(table, "lat")
	u.LoginTime = field.NewInt32(table, "login_time")
	u.LoginCount = field.NewInt32(table, "login_count")
	u.Gender = field.NewInt32(table, "gender")
	u.CreateTime = field.NewInt32(table, "create_time")
	u.UpdateTime = field.NewInt32(table, "update_time")
	u.AlipayAccount = field.NewString(table, "alipay_account")
	u.AlipayName = field.NewString(table, "alipay_name")
	u.AlipayQrcode = field.NewString(table, "alipay_qrcode")

	u.fillFieldMap()

	return u
}

func (u *user) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *user) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 21)
	u.fieldMap["id"] = u.ID
	u.fieldMap["username"] = u.Username
	u.fieldMap["nickname"] = u.Nickname
	u.fieldMap["mobile"] = u.Mobile
	u.fieldMap["avatar"] = u.Avatar
	u.fieldMap["real_name"] = u.RealName
	u.fieldMap["wx_number"] = u.WxNumber
	u.fieldMap["location_address"] = u.LocationAddress
	u.fieldMap["address"] = u.Address
	u.fieldMap["model"] = u.Model
	u.fieldMap["lng"] = u.Lng
	u.fieldMap["lat"] = u.Lat
	u.fieldMap["login_time"] = u.LoginTime
	u.fieldMap["login_count"] = u.LoginCount
	u.fieldMap["gender"] = u.Gender
	u.fieldMap["create_time"] = u.CreateTime
	u.fieldMap["update_time"] = u.UpdateTime
	u.fieldMap["alipay_account"] = u.AlipayAccount
	u.fieldMap["alipay_name"] = u.AlipayName
	u.fieldMap["alipay_qrcode"] = u.AlipayQrcode

}

func (u user) clone(db *gorm.DB) user {
	u.userDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u user) replaceDB(db *gorm.DB) user {
	u.userDo.ReplaceDB(db)
	return u
}

type userHasOneUserOauth struct {
	db *gorm.DB

	field.RelationField
}

func (a userHasOneUserOauth) Where(conds ...field.Expr) *userHasOneUserOauth {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a userHasOneUserOauth) WithContext(ctx context.Context) *userHasOneUserOauth {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userHasOneUserOauth) Session(session *gorm.Session) *userHasOneUserOauth {
	a.db = a.db.Session(session)
	return &a
}

func (a userHasOneUserOauth) Model(m *entity.User) *userHasOneUserOauthTx {
	return &userHasOneUserOauthTx{a.db.Model(m).Association(a.Name())}
}

type userHasOneUserOauthTx struct{ tx *gorm.Association }

func (a userHasOneUserOauthTx) Find() (result *entity.UserOauth, err error) {
	return result, a.tx.Find(&result)
}

func (a userHasOneUserOauthTx) Append(values ...*entity.UserOauth) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userHasOneUserOauthTx) Replace(values ...*entity.UserOauth) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userHasOneUserOauthTx) Delete(values ...*entity.UserOauth) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userHasOneUserOauthTx) Clear() error {
	return a.tx.Clear()
}

func (a userHasOneUserOauthTx) Count() int64 {
	return a.tx.Count()
}

type userDo struct{ gen.DO }

type IUserDo interface {
	gen.SubQuery
	Debug() IUserDo
	WithContext(ctx context.Context) IUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserDo
	WriteDB() IUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserDo
	Not(conds ...gen.Condition) IUserDo
	Or(conds ...gen.Condition) IUserDo
	Select(conds ...field.Expr) IUserDo
	Where(conds ...gen.Condition) IUserDo
	Order(conds ...field.Expr) IUserDo
	Distinct(cols ...field.Expr) IUserDo
	Omit(cols ...field.Expr) IUserDo
	Join(table schema.Tabler, on ...field.Expr) IUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserDo
	Group(cols ...field.Expr) IUserDo
	Having(conds ...gen.Condition) IUserDo
	Limit(limit int) IUserDo
	Offset(offset int) IUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserDo
	Unscoped() IUserDo
	Create(values ...*entity.User) error
	CreateInBatches(values []*entity.User, batchSize int) error
	Save(values ...*entity.User) error
	First() (*entity.User, error)
	Take() (*entity.User, error)
	Last() (*entity.User, error)
	Find() ([]*entity.User, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.User, err error)
	FindInBatches(result *[]*entity.User, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.User) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserDo
	Assign(attrs ...field.AssignExpr) IUserDo
	Joins(fields ...field.RelationField) IUserDo
	Preload(fields ...field.RelationField) IUserDo
	FirstOrInit() (*entity.User, error)
	FirstOrCreate() (*entity.User, error)
	FindByPage(offset int, limit int) (result []*entity.User, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userDo) Debug() IUserDo {
	return u.withDO(u.DO.Debug())
}

func (u userDo) WithContext(ctx context.Context) IUserDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userDo) ReadDB() IUserDo {
	return u.Clauses(dbresolver.Read)
}

func (u userDo) WriteDB() IUserDo {
	return u.Clauses(dbresolver.Write)
}

func (u userDo) Session(config *gorm.Session) IUserDo {
	return u.withDO(u.DO.Session(config))
}

func (u userDo) Clauses(conds ...clause.Expression) IUserDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userDo) Returning(value interface{}, columns ...string) IUserDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userDo) Not(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userDo) Or(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userDo) Select(conds ...field.Expr) IUserDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userDo) Where(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userDo) Order(conds ...field.Expr) IUserDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userDo) Distinct(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userDo) Omit(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userDo) Join(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userDo) Group(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userDo) Having(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userDo) Limit(limit int) IUserDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userDo) Offset(offset int) IUserDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userDo) Unscoped() IUserDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userDo) Create(values ...*entity.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userDo) CreateInBatches(values []*entity.User, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userDo) Save(values ...*entity.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userDo) First() (*entity.User, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.User), nil
	}
}

func (u userDo) Take() (*entity.User, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.User), nil
	}
}

func (u userDo) Last() (*entity.User, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.User), nil
	}
}

func (u userDo) Find() ([]*entity.User, error) {
	result, err := u.DO.Find()
	return result.([]*entity.User), err
}

func (u userDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.User, err error) {
	buf := make([]*entity.User, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userDo) FindInBatches(result *[]*entity.User, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userDo) Attrs(attrs ...field.AssignExpr) IUserDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userDo) Assign(attrs ...field.AssignExpr) IUserDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userDo) Joins(fields ...field.RelationField) IUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userDo) Preload(fields ...field.RelationField) IUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userDo) FirstOrInit() (*entity.User, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.User), nil
	}
}

func (u userDo) FirstOrCreate() (*entity.User, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.User), nil
	}
}

func (u userDo) FindByPage(offset int, limit int) (result []*entity.User, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userDo) Delete(models ...*entity.User) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userDo) withDO(do gen.Dao) *userDo {
	u.DO = *do.(*gen.DO)
	return u
}
