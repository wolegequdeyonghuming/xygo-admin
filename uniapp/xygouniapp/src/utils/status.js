// 订单状态配置（收单员视角，Field Deck 色板）
export const ORDER_STATUS = {
  0: { text: '未成交', color: '#8A8F98', bg: '#E6E8EB' },
  1: { text: '已录单', color: '#E8930C', bg: '#FBEED7' },
  2: { text: '待预约', color: '#E8930C', bg: '#FBEED7' },
  3: { text: '待收单', color: '#E5484D', bg: '#FCE3E4' },
  4: { text: '已收单', color: '#30A46C', bg: '#DCF0E6' },
  5: { text: '已回访', color: '#30A46C', bg: '#DCF0E6' },
  6: { text: '已完工', color: '#30A46C', bg: '#DCF0E6' }
}

export function orderStatus(status) {
  return ORDER_STATUS[String(status)] || { text: '-', color: '#8A8F98', bg: '#E6E8EB' }
}
