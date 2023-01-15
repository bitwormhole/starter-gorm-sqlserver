// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfgsqlserver

import (
	driver0x11d6c2 "github.com/bitwormhole/starter-gorm-sqlserver/driver"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: com0-driver0x11d6c2.SQLServerDriver
	cominfobuilder.Next()
	cominfobuilder.ID("com0-driver0x11d6c2.SQLServerDriver").Class("starter-gorm-driver-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComSQLServerDriver{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComSQLServerDriver : the factory of component: com0-driver0x11d6c2.SQLServerDriver
type comFactory4pComSQLServerDriver struct {

    mPrototype * driver0x11d6c2.SQLServerDriver

	

}

func (inst * comFactory4pComSQLServerDriver) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComSQLServerDriver) newObject() * driver0x11d6c2.SQLServerDriver {
	return & driver0x11d6c2.SQLServerDriver {}
}

func (inst * comFactory4pComSQLServerDriver) castObject(instance application.ComponentInstance) * driver0x11d6c2.SQLServerDriver {
	return instance.Get().(*driver0x11d6c2.SQLServerDriver)
}

func (inst * comFactory4pComSQLServerDriver) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComSQLServerDriver) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComSQLServerDriver) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComSQLServerDriver) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSQLServerDriver) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSQLServerDriver) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}




