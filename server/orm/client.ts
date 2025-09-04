import "reflect-metadata";
import { DataSource } from "typeorm";
import { role } from "./entity/relationship.js";
import { user } from "./entity/user.js";
import { logger } from "../utils/log.js";
import bcrypt from "bcryptjs";
import {generateUserId} from "../utils/general.js"


export const db = new DataSource({
    type: "mysql",
    host: process.env.DB_HOST!,
    port: Number(process.env.DB_PORT),
    username: process.env.DB_USER!,
    password: process.env.DB_PASS!,
    database: process.env.DB_NAME!,
    entities: [user],
    synchronize: true,
});

// 初始化DB连接
export async function InitDBClient() {
    if (!db.isInitialized) {
        await db.initialize();
        logger.info("数据库连接成功 ✅");

        const userRepository = db.getRepository(user);
        const count = await userRepository.count();
        
        if (count === 0) {
            if (!process.env.ADMIN_NAME || !process.env.ADMIN_PASS) {
                throw new Error("未知的管理员用户或者密码,请检查ADMIN_NAME 和 ADMIN_PASS相关配置");
            }
            const admin = userRepository.create({
                name: process.env.ADMIN_NAME!,
                uid: generateUserId(),
                password: await bcrypt.hash(process.env.ADMIN_PASS!, 10), 
                role: role.admin,
                lastLoginAt: new Date()
            });
            await userRepository.save(admin);
            logger.info("第一次启动,管理员账号已自动创建 ✅");
        }
    }
}