package models

type permisionsArray struct {
	Reason     string `form:"reason" json:"reason" binding:"required"`
	Permission string `form:"permission" json:"permission" binding:"required"`
}

type AmeConf struct {
	Email                string            `form:"email" json:"email" binding:"required"`
	Name                 string            `form:"name" json:"name" binding:"required"`
	Title                string            `form:"title" json:"title" binding:"required"`
	Slug                 string            `form:"slug" json:"slug" binding:"required"`
	Version              string            `form:"version" json:"version" binding:"required"`
	PublicKey            string            `form:"public-key" json:"public-key" binding:"required"`
	AmeMiniappComponents string            `form:"ame-miniapp-components" json:"ame-miniapp-components" binding:"required"`
	AmeSuperAppClient    string            `form:"ame-super-app-client" json:"ame-super-app-client" binding:"required"`
	Permissions          []permisionsArray `form:"permissions" json:"permissions"`
}
