
// 访问状态
import { store } from '@naturefw/nf-state'

/**
 * 从 json 加载meta
 */

let baseUrl = ''

const urls = {
  button: (key) => `${baseUrl}json/${key}_button.json`,
  pager: (key) => `${baseUrl}json/${key}_pager.json`,
  grid: (key) => `${baseUrl}json/${key}_grid.json`,
  form: (key) => `${baseUrl}json/${key}_form.json`,
  find: (key) => `${baseUrl}json/${key}_find.json`
}

const loadMeta = async (key) => {
  baseUrl = store.web.baseUrl

  const meta = {
    moduleId: key,
    gridMeta: (await axios.get(urls.grid(key))).data,
    // pager: await await axios.get(urls.pager(key)),
    formMeta: (await axios.get(urls.form(key))).data,
    findMeta: (await axios.get(urls.find(key))).data,
    buttonMeta: (await axios.get(urls.button(key))).data
  }

  store.meta[key] = meta
  return meta
}

export {
  loadMeta
}
