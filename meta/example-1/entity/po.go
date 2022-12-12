// condition = <entity2>.<attr>=<value>
// <entity1> -condition[ AND condition ]*-> <entity alias>
// e.g. <application> -[operation.type = 1]-> <subscriber>
// e.g. operation.app_unq_id ref application.id

// id: auto/named/given
// auto-int autoincrement; auto-string uuid
// named - 由外部设置
// given - 通过接口获取

// req: update - update related entity attribute(column)
// req: insert - create related entity attribute(column)
// req: upsert - create/update related entity attribute(column)
// req: batch - in
// req: fuzzy - like
// req: exact - eq
// req: filter - join
// status: 后续扩展，支持做审批流程
// role: 用于定义 relationship 中的角色
// ref: <table>.<primaryKey> - entity ref
// fer: <table>.<foreignKey> - entity be reffered

// must_a(primaryKey): <table>.<foreignKey> - one-to-one relationship                          account.id          - must_a - secret.account_id   (账户必须有且只有一个密钥)
// tsum_a(foreignKey): <table>.<primaryKey> - one-to-one relationship                          secret.account_id   - tsum_a - account.id
// may_a(primaryKey): <table>.<foreignKey> - one(this)-to-one(that optional) relationship      account.id          - may_a  - passport.account_id (账户可以只有/没有一个护照)
// yam_a(foreignKey): <table>.<primaryKey> - one(this optional)-to-one(that) relationship      passport.account_id - yam_a  - account.id          (护照必须关联一个账户)

// account 1 --> 0/n book [假设并非所有人都有书]
// may_s(primaryKey): <table>.<foreignKey> - one(this)-to-many(that optional) relationship     account.id          - may_s  - book.owner_id       (账户可以有/没有多本书) 说明：这里的书是具体的某一本书
//                                                                                             this account is owner for that book
// yam_s(foreignKey): <table>.<primaryKey> - many(this optional)-to-one(that) relationship     book.owner_id       - yam_s  - account.id          (一本书必须关联一个所有者)
//                                                                                             this book's owner is that account

// house 1 --> 1/n account [假设符合所有房子都有人住]
// must_s(primaryKey): <table>.<foreignKey> - one(this)-to-many(that) relationship             house.id                  - must_s - account.living_address_id   (房屋必须有一或多个账户/住户)
//                                                                                             this house is living-address for that account
// tsum_s(foreignKey): <table>.<primaryKey> - many(this)-to-one(that) relationship             account.living_address_id - tsum_s - house.id   (账户必须关联一个现住址)
//                                                                                             this account's living-address is that house

// course 0/n --> 1/m account [假设所有课程都有人学，但并非人人都是学生]
// s_may_must_s: <rel-table>.<foreignKey>  - many(this optional)-to-many(that) relationship    account.id      - s_may_must_s(study)      - study_record.student_id
//                                                                                             this account is student for that course
// s_must_may_s: <rel-table>.<foreignKey>  - many(this)-to-many(that optional) relationship    course.id       - s_must_may_s(is learned) - study_record.course_id
//                                                                                             this course is learned by that student

// course 0/n --> 0/m account [假设并非所有直播都有人看，并非所有人都看直播]
// s_may_s: <rel-table>.<foreignKey>  - many(this optional)-to-many(that optional) relationship    account.id      - s_may_s(watch)      - watch_record.audience_id
//                                                                                                 this account is audience for that play
// s_may_s: <rel-table>.<foreignKey>  - many(this optional)-to-many(that optional) relationship    play.id         - s_may_s(is watched) - watch_record.play_id
//                                                                                                 this course is learned by that student

// course n --> m account [假设所有商品都有人买，所有人都买商品]
// s_must_s: <rel-table>.<foreignKey>  - many(this)-to-many(that) relationship                 account.id      - s_must_s(buy)      - shopping_record.consumer_id
//                                                                                             this account is the consumer for that goods
// s_must_s: <rel-table>.<foreignKey>  - many(this)-to-many(that) relationship                 product.id      - s_must_s - shopping_record.goods_id
//                                                                                             this product is bought as goods by that consumer

// this account is student in that course
// this course is
// rsp: detail
// rsp: simple
package entity
