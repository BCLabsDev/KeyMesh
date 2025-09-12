package models

import (
	"gorm.io/gorm"
)

// 登录方式枚举
type LoginProvider string

const (
	LoginProviderEmail LoginProvider = "email"
	LoginProviderPhone LoginProvider = "phone"
)

// 用户表
type User struct {
	gorm.Model
	Name     string `gorm:"not null;unique"`
	UID      string `gorm:"not null;unique"`
	Password string

	Enable bool `gorm:"default:false"`
}

// 登录和授权记录
type Record struct {
	gorm.Model
	UserID   uint          `gorm:"not null;index"`                                // 外键
	Provider LoginProvider `gorm:"type:enum('email', 'phone');not null"`          // 登录类型
	User     User          `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // 外键关联
}

// 授权平台
type OAuthProvider struct {
	gorm.Model
	UserID   uint          `gorm:"not null;index"`                                // 外键，关联 users 表
	Provider LoginProvider `gorm:"type:enum('weixin', 'github');not null"`        // 第三方平台
	OpenID   string        `gorm:"type:varchar(100);unique;not null"`             // 第三方用户标识
	User     User          `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // 外键关联
}
