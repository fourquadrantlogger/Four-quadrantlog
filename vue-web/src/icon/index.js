import { reactive } from 'vue'

import {
  CloseBold,
  Close,
  Plus,
  Star,
  UserFilled,
  Loading,
  Connection,
  Edit,
  FolderOpened
} from '@element-plus/icons-vue'

const dictIcon = reactive({
  'CloseBold': CloseBold,
  'Close': Close,
  'Plus': Plus,
  'Star': Star,
  'UserFilled': UserFilled,
  'Loading': Loading,
  'Connection': Connection,
  'Edit': Edit,
  'FolderOpened': FolderOpened
})

const installIcon = (app) => {
  // 便于模板获取
  app.config.globalProperties.$icon = dictIcon
}

export default installIcon