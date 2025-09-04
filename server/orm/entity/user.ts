import "reflect-metadata";
import {role} from "./relationship.js"

import { Entity, PrimaryGeneratedColumn, Column, CreateDateColumn, UpdateDateColumn } from "typeorm";


@Entity()
export class user {
  // 数据库ID
  @PrimaryGeneratedColumn()
  id: number = 0; 

  // 用户可见ID
  @Column({ type: "varchar", length: 16,nullable: false  })
  uid: string = ""  ;

  // 用户名
  @Column({ type: "varchar", length: 30,nullable: false  })
  name: string = "";

  // 角色
  @Column({ type: "enum",  enum: role,  default: role.user,nullable: false })
  role: role = role.user;

  // 密码
  @Column({ type: "varchar", length: 100 ,nullable: false })
  password: string = "";

  // 手机号
  @Column({ unique: true,  type: "varchar", length: 20, nullable: true })
  phone: string | null = null;

  // 邮箱
  @Column({ unique: true, type: "varchar",  length: 100, nullable: true })
  email: string | null = null;

  // 微信唯一ID (openId/unionId)
  @Column({ unique: true, type: "varchar",  length: 100, nullable: true })
  wechatId: string | null = null;

  // GitHub 唯一ID
  @Column({ unique: true, type: "varchar",  length: 100, nullable: true })
  githubId: string | null = null;

  // 最新活跃时间
  @Column({ type: "datetime", nullable: false})
  lastLoginAt?: Date;

  // 创建时间
  @CreateDateColumn()
  createdAt!: Date;

  // 账号数据更新时间
  @UpdateDateColumn()
  updatedAt!: Date;
}