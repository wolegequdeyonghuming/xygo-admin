/**
 * useOrderPerm - 订单角色/步骤权限判断（与后端 perm.go 保持一致）
 *
 * 前端 roles 来自 userStore.info.roles（角色 key 数组），可能存在多个角色，
 * 这里按优先级取「订单相关角色」用于页面按钮/操作列的显隐判断。
 * 后端仍会二次校验，前端仅做交互层控制。
 */
import { computed } from 'vue'
import { useUserStore } from '@/store/modules/user'

/** 订单状态常量（与后端 consts/order.go 一致） */
export const ORDER_STATUS = {
  NOT_DEAL: 0, // 未成交
  RECORDED: 1, // 已录单
  ASSIGNED: 2, // 已接单
  APPOINTED: 3, // 已预约
  VISITED: 4, // 已上门
  FOLLOWED: 5, // 已回访
  COMPLETED: 6 // 已完工
} as const

/** 订单步骤（与后端 step 一致） */
export const ORDER_STEP = {
  RECORD: 1, // 录单
  ASSIGN: 2, // 派单
  APPOINT: 3, // 预约
  VISIT: 4, // 上门
  FOLLOW: 5, // 回访
  COMPLETE: 6 // 完成
} as const

/** 订单相关角色 key（集中管理） */
const ORDER_ROLES: Record<string, string> = {
  ADMIN: 'admin',
  TELEMARKETER_MANAGER: 'telemarketer_manager',
  AGENT_MANAGER: 'agent_manager',
  DOCUMENTARY: 'documentary',
  TELEMARKETER: 'telemarketer',
  AGENT: 'agent',
  SUPER_ADMIN: 'super_admin'
}

/** 角色优先级：值越小优先级越高（决定「当前订单角色」取哪一个） */
const ROLE_PRIORITY: string[] = [
  ORDER_ROLES.ADMIN,
  ORDER_ROLES.SUPER_ADMIN,
  ORDER_ROLES.AGENT_MANAGER,
  ORDER_ROLES.TELEMARKETER_MANAGER,
  ORDER_ROLES.DOCUMENTARY,
  ORDER_ROLES.TELEMARKETER,
  ORDER_ROLES.AGENT
]

export const useOrderPerm = () => {
  const userStore = useUserStore()

  /** 当前用户的订单角色 key（多角色时取优先级最高的订单角色） */
  const orderRole = computed<string>(() => {
    const roles: string[] = userStore.info.roles ?? []
    for (const key of ROLE_PRIORITY) {
      if (roles.includes(key)) return key
    }
    return roles[0] ?? ''
  })

  const role = () => orderRole.value

  /** 是否可查看全部订单（收单员管理员/财务文员/管理员/超管） */
  const canViewAll = (): boolean =>
    [
      ORDER_ROLES.AGENT_MANAGER,
      ORDER_ROLES.DOCUMENTARY,
      ORDER_ROLES.ADMIN,
      ORDER_ROLES.SUPER_ADMIN
    ].includes(role())

  /** 是否可查看全部订单（computed 版本，供模板使用） */
  const canViewAllOrders = computed(() => canViewAll())

  /** 当前用户 ID（用于操作本人名下订单判断） */
  const currentUserId = (): number => userStore.info.id ?? 0

  /**
   * 是否允许该角色执行某步骤（不含状态推进校验，仅角色维度）
   * step 对应 ORDER_STEP
   */
  const canStep = (step: number): boolean => {
    return canStepForRole(role(), step)
  }

  /**
   * 判断用户任一角色是否允许执行某步骤（含管理员/超管）。
   * 用于顶部操作按钮等需要「任一角色具备即可」的场景（兼容多角色账号）。
   */
  const anyRoleCanStep = (step: number): boolean => {
    const roles: string[] = userStore.info.roles ?? []
    return roles.some((key) => canStepForRole(key, step))
  }

  /** 按具体角色 key 判断步骤权限（canStep 的单角色版本，供 anyRoleCanStep 复用） */
  const canStepForRole = (roleKey: string, step: number): boolean => {
    if ([ORDER_ROLES.ADMIN, ORDER_ROLES.SUPER_ADMIN].includes(roleKey)) {
      return step >= ORDER_STEP.RECORD && step <= ORDER_STEP.COMPLETE
    }
    switch (step) {
      case ORDER_STEP.RECORD:
        return roleKey === ORDER_ROLES.TELEMARKETER || roleKey === ORDER_ROLES.TELEMARKETER_MANAGER
      case ORDER_STEP.ASSIGN:
        return roleKey === ORDER_ROLES.AGENT_MANAGER
      case ORDER_STEP.APPOINT:
        return roleKey === ORDER_ROLES.AGENT || roleKey === ORDER_ROLES.AGENT_MANAGER
      case ORDER_STEP.VISIT:
        return roleKey === ORDER_ROLES.AGENT || roleKey === ORDER_ROLES.AGENT_MANAGER
      case ORDER_STEP.FOLLOW:
        return roleKey === ORDER_ROLES.TELEMARKETER || roleKey === ORDER_ROLES.TELEMARKETER_MANAGER
      case ORDER_STEP.COMPLETE:
        return roleKey === ORDER_ROLES.DOCUMENTARY
    }
    return false
  }

  /**
   * 判断对处于 status 状态的订单执行 step 步骤是否被允许（含状态推进）。
   * 允许：原地保存（step === status）或正常推进（step === status + 1）。
   */
  const canStepFromStatus = (status: number, step: number): boolean => {
    if (!canStep(step)) return false
    if (step === status + 1) return true
    // 录单步骤：仅未成交/初始状态可录入
    if (step === ORDER_STEP.RECORD && status !== ORDER_STATUS.NOT_DEAL) return false
    return false
  }

  /**
   * 是否有权操作该订单：
   * - 话务员/话务员管理员：只能操作自己录的单（createdBy）
   * - 收单员：只能操作派给自己的单（agentId）
   * - 其余（收单员管理员/财务文员/管理员/超管）：不限
   */
  const canOperateOrder = (order: Record<string, any>): boolean => {
    if (canViewAll()) return true
    const r = role()
    const uid = currentUserId()
    switch (r) {
      case ORDER_ROLES.TELEMARKETER:
      case ORDER_ROLES.TELEMARKETER_MANAGER:
        return Number(order.createdBy) === uid
      case ORDER_ROLES.AGENT:
        return Number(order.agentId) === uid
    }
    return false
  }

  /**
   * 是否允许删除某状态订单：
   * - 话务员/话务员管理员：仅已录单（status=1）
   * - 管理员/超管：任意
   */
  const canDeleteOrder = (order: Record<string, any>): boolean => {
    const r = role()
    if ([ORDER_ROLES.ADMIN, ORDER_ROLES.SUPER_ADMIN].includes(r)) return true
    if (
      [ORDER_ROLES.TELEMARKETER, ORDER_ROLES.TELEMARKETER_MANAGER].includes(r) &&
      Number(order.orderStatus) === ORDER_STATUS.RECORDED
    ) {
      return true
    }
    return false
  }

  /**
   * 该订单当前应展示的「下一步」步骤（0 表示无下一步）。
   * 规则：管理员看状态推进，其余角色按 canStepFromStatus 判定可执行的下一步。
   */
  const nextStepForOrder = (order: Record<string, any>): number => {
    const status = Number(order.orderStatus)
    const next = status + 1
    if (next > ORDER_STATUS.COMPLETED) return 0
    // 未成交状态时，话务员/管理员执行录单；其余状态按角色推进
    if (canStepFromStatus(status, next) && canOperateOrder(order)) return next
    return 0
  }

  /** 该订单是否允许执行「录单」步骤（用于列表行内补录/编辑录单） */
  const canRecordOrder = (order: Record<string, any>): boolean => {
    const status = Number(order.orderStatus)
    if (status !== ORDER_STATUS.NOT_DEAL && status !== ORDER_STATUS.RECORDED) return false
    return canStep(ORDER_STEP.RECORD) && canOperateOrder(order)
  }

  return {
    orderRole,
    canViewAll,
    canViewAllOrders,
    canStep,
    anyRoleCanStep,
    canStepFromStatus,
    canOperateOrder,
    canDeleteOrder,
    nextStepForOrder,
    canRecordOrder,
    currentUserId
  }
}
