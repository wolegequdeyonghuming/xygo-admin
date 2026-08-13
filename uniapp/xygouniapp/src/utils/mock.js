// 收单小程序 mock 数据（后端 /staff/* 接口就绪前用于 UI 联调）
// TODO: 后端 M1 完成后删除本文件，改为调用 api/staff.js 真实接口

export const mockOrders = [
  {
    id: 1,
    orderNo: '202608110001',
    orderStatus: 2,
    customerName: '王小明',
    contactPhone: '13800001234',
    businessType: '携号转网',
    visitDate: '2026-08-11',
    scheduleDate: '2026-08-10',
    installAddress: '大连市沙河口区星海街道星海广场6号',
    availableTimeDesc: '9:00-12:00'
  },
  {
    id: 2,
    orderNo: '202608110002',
    orderStatus: 3,
    customerName: '李四',
    contactPhone: '13900005678',
    businessType: '新装宽带',
    visitDate: '2026-08-11',
    scheduleDate: '2026-08-10',
    installAddress: '大连市甘井子区华北路88号',
    availableTimeDesc: '14:00-18:00',
    appointmentDesc: '客户要求下午上门，套餐保持宽带100M'
  },
  {
    id: 3,
    orderNo: '202608110003',
    orderStatus: 2,
    customerName: '张伟',
    contactPhone: '13700001111',
    businessType: '话费充值',
    visitDate: '2026-08-11',
    scheduleDate: '2026-08-11',
    installAddress: '大连市西岗区黄河路50号',
    availableTimeDesc: '全天'
  },
  {
    id: 4,
    orderNo: '202608110004',
    orderStatus: 4,
    customerName: '刘芳',
    contactPhone: '13600002222',
    businessType: '携号转网',
    visitDate: '2026-08-11',
    scheduleDate: '2026-08-10',
    installAddress: '大连市中山区人民路10号',
    availableTimeDesc: '9:00-11:00',
    appointmentDesc: '已预约上午上门',
    dealtBusinessType: '携转成功',
    paidAmount: 100
  },
  {
    id: 5,
    orderNo: '202608110005',
    orderStatus: 5,
    customerName: '陈杰',
    contactPhone: '13500003333',
    businessType: '新装宽带',
    visitDate: '2026-08-11',
    scheduleDate: '2026-08-09',
    installAddress: '大连市金州区光明街道',
    availableTimeDesc: '下午',
    appointmentDesc: '已上门，回访中'
  },
  {
    id: 6,
    orderNo: '202608100001',
    orderStatus: 6,
    customerName: '周涛',
    contactPhone: '13400004444',
    businessType: '话费充值',
    visitDate: '2026-08-10',
    scheduleDate: '2026-08-09',
    installAddress: '大连市旅顺口区白玉山',
    availableTimeDesc: '全天',
    appointmentDesc: '已完工',
    dealtBusinessType: '充值完成'
  }
]

// 待办 = 待预约(2) + 待收单(3)；已办 = 已上门(4)/已回访(5)/已完工(6)
export const STAT_TABS = [
  { key: 'todo', label: '待收单', color: '#D92400', statuses: [2, 3], title: '待收单订单' },
  { key: 'today', label: '今日收单', color: '#299A0C', statuses: [4, 5, 6], title: '今日收单订单' },
  { key: 'done', label: '已收单', color: '#1F61FF', statuses: [4, 5, 6], title: '已收单订单' }
]

export function filterOrders(tabIndex, list, keyword = '', month = '') {
  const tab = STAT_TABS[tabIndex]
  let arr = list.filter((o) => tab.statuses.includes(Number(o.orderStatus)))
  if (keyword) {
    arr = arr.filter(
      (o) =>
        (o.customerName || '').includes(keyword) || (o.contactPhone || '').includes(keyword)
    )
  }
  // 已收单 tab：按上门日期所属月份过滤
  if (tabIndex === 2 && month) {
    arr = arr.filter((o) => (o.visitDate || '').startsWith(month))
  }
  return arr
}

// mock 预览：当前收单员资料（后端就绪后由 /staff/user/profile 返回）
export const mockProfile = {
  id: 9,
  username: 'wanbing',
  realName: '万冰',
  nickname: '万冰',
  mobile: '13888886666',
  postName: '收单员',
  deptName: '收单部'
}

// mock 预览：详细情况（后端就绪后由 /staff/order/detailList 返回）
export const mockDetails = [
  {
    id: 1,
    orderId: 2,
    userId: 1,
    userName: '未知用户',
    createdAt: '2026-08-09 10:53:14',
    content: '客户要求套餐不变'
  }
]
