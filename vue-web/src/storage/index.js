
// 创建help
import { dbCreateHelp } from '@naturefw/storage'

// 访问状态
import { store } from '@naturefw/nf-state'

// 设置数据库名称和版本
const db = {
  dbName: 'nf-four-quadrant-cus',
  ver: 1
}

/**
 * 客户项目的 meta 的 db 缓存
 */
export default async function createDBHelp (callback) {
  // 设置基础路径
  // config.baseUrl = url

  const help = dbCreateHelp({
    // dbFlag: 'project-meta-db',
    dbConfig: db,
    stores: { // 数据库里的对象仓库
      db_fq_log: { // 可以存放多个项目的文档，包含导航信息和项目信息
        id: 'id',
        index: {},
        isClear: false
        /**
         */
      }
    },
    // 设置初始数据
    async init (help) {
      const { web } = store
      // 把 help 存入状态
      web.dbHelp = help
    }
  })
  return help
}
