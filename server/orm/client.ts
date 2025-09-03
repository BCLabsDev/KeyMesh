import "reflect-metadata";
import { DataSource } from "typeorm";
import { User } from "./entity/user.js";
import { logger } from "../utils/log.js";

export const db = new DataSource({
    type: "mysql",
    host: process.env.DB_HOST!,
    port: Number(process.env.DB_PORT),
    username: process.env.DB_USER!,
    password: process.env.DB_PASS!,
    database: process.env.DB_NAME!,
    entities: [User],
    synchronize: true,
});

// 初始化DB连接
export async function InitDBClient() {
    if (!db.isInitialized) {
        await db.initialize();
        logger.info("数据库连接成功 ✅");
    }
}
