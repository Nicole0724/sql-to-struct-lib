
package entityModel

type DictBizBase struct{
	//主键 
	Id	int32 `gorm:"id" form:"id" json:"id"` 
	//租户ID 
	TenantId	string `gorm:"tenant_id" form:"tenantId" json:"tenantId"` 
	//父主键 
	ParentId	int64 `gorm:"parent_id" form:"parentId" json:"parentId"` 
	//字典码 
	Code	string `gorm:"code" form:"code" json:"code"` 
	//字典值 
	DictKey	string `gorm:"dict_key" form:"dictKey" json:"dictKey"` 
	//字典名称 
	DictValue	string `gorm:"dict_value" form:"dictValue" json:"dictValue"` 
	//排序 
	Sort	int32 `gorm:"sort" form:"sort" json:"sort"` 
	//字典备注 
	Remark	string `gorm:"remark" form:"remark" json:"remark"` 
	//是否已封存 
	IsSealed	int32 `gorm:"is_sealed" form:"isSealed" json:"isSealed"` 
}

type DictBiz struct{
 	//主键  
	Id	int32 `gorm:"id" form:"id" json:"id"` 
 	//租户ID  
	TenantId	string `gorm:"tenant_id" form:"tenantId" json:"tenantId"` 
 	//父主键  
	ParentId	int64 `gorm:"parent_id" form:"parentId" json:"parentId"` 
 	//字典码  
	Code	string `gorm:"code" form:"code" json:"code"` 
 	//字典值  
	DictKey	string `gorm:"dict_key" form:"dictKey" json:"dictKey"` 
 	//字典名称  
	DictValue	string `gorm:"dict_value" form:"dictValue" json:"dictValue"` 
 	//排序  
	Sort	int32 `gorm:"sort" form:"sort" json:"sort"` 
 	//字典备注  
	Remark	string `gorm:"remark" form:"remark" json:"remark"` 
 	//是否已封存  
	IsSealed	int32 `gorm:"is_sealed" form:"isSealed" json:"isSealed"` 
 	//创建人  
	CreateUser	int32 `gorm:"create_user" form:"createUser" json:"createUser"` 
 	//创建时间  
	CreateTime	string `gorm:"create_time" form:"createTime" json:"createTime"` 
 	//修改人  
	UpdateUser	int32 `gorm:"update_user" form:"updateUser" json:"updateUser"` 
 	//修改时间  
	UpdateTime	string `gorm:"update_time" form:"updateTime" json:"updateTime"` 
 	//是否已删除  
	IsDeleted	int32 `gorm:"is_deleted" form:"isDeleted" json:"isDeleted"` 

}
