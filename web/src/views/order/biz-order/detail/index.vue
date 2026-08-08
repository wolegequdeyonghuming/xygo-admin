<!-- 订单主表 详情页（全屏） -->
<template>
  <div class="biz-order-detail art-full-height">
    <ElCard shadow="never" class="art-table-card">
      <template #header>
        <div class="flex items-center justify-between">
          <span class="font-bold text-lg">订单主表详情</span>
          <ElButton @click="goBack">
            <ArtSvgIcon icon="ri:arrow-left-line" class="text-sm mr-1" />
            返回列表
          </ElButton>
        </div>
      </template>

      <div v-if="loading" class="flex justify-center py-20">
        <ElIcon class="is-loading" :size="32"><Loading /></ElIcon>
      </div>

      <ElDescriptions v-else-if="detail" :column="2" border class="detail-descriptions">
        <ElDescriptionsItem label="订单状态">
          <DictLabel :dict-type="'order_status'" :value="detail.orderStatus" />
        </ElDescriptionsItem>
        <ElDescriptionsItem label="排单日期">{{ formatTimestamp(detail.scheduleDate) }}</ElDescriptionsItem>
        <ElDescriptionsItem label="上门日期">{{ formatTimestamp(detail.visitDate) }}</ElDescriptionsItem>
        <ElDescriptionsItem label="客户姓名">{{ detail.customerName ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="可联系时间">{{ detail.availableTimeDesc ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="区县">
          <DictLabel :dict-type="'area'" :value="detail.area" />
        </ElDescriptionsItem>
        <ElDescriptionsItem label="安装地址">{{ detail.installAddress ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="联系电话">{{ detail.contactPhone ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="预约业务">{{ detail.businessType ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="到期时间">{{ detail.expiryDate ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="主卡号码">{{ detail.primaryTelNo ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="详细情况">{{ detail.details ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="话务员">{{ detail.telemarketer_real_name ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="收单员">{{ detail.agent_real_name ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="收单预约情况">{{ detail.appointmentDesc ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="话务二次回放情况">{{ detail.followUpDesc ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="成交业务">{{ detail.dealtBusinessType ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="携转情况">{{ detail.portingStatus ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="客户实际姓名">{{ detail.customerRealName ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="客户身份证号">{{ detail.customerIdNumber ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="实缴额度">{{ detail.paidAmount ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="是否乡下单">
          <ElTag size="small">{{ ({ '0': '否', '1': '是',  })[String(detail.isRuralOrder)] || detail.isRuralOrder }}</ElTag>
        </ElDescriptionsItem>
        <ElDescriptionsItem label="新开号码">{{ detail.newPhoneNo ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="终端串码">{{ detail.deviceSerial ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="宽带账号">{{ detail.broadbandAccount ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="是否完工">
          <ElTag size="small">{{ ({ '0': '未完工', '1': '已完工',  })[String(detail.isCompleted)] || detail.isCompleted }}</ElTag>
        </ElDescriptionsItem>
        <ElDescriptionsItem label="话补">{{ detail.subsidyAmount ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="工号">{{ detail.agencyNo ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="是否纯新增">{{ detail.isNew ?? '-' }}</ElDescriptionsItem>
        <ElDescriptionsItem label="附件">
          <a v-if="detail.attachmentId" :href="detail.attachmentId" target="_blank" style="color:var(--el-color-primary)">{{ detail.attachmentId }}</a>
          <span v-else>-</span>
        </ElDescriptionsItem>
        <ElDescriptionsItem label="备注">{{ detail.remark ?? '-' }}</ElDescriptionsItem>
      </ElDescriptions>

      <div v-else class="py-20 text-center text-gray-400">
        数据加载失败
      </div>
    </ElCard>
  </div>
</template>

<script setup lang="ts">
  import { Loading } from '@element-plus/icons-vue'
  import ArtSvgIcon from '@/components/core/base/art-svg-icon/index.vue'
  import { fetchBizOrderView } from '@/api/order/biz-order'
  import { formatTimestamp } from '@/utils/time'
  import { useDictStore } from '@/store/modules/dict'
  import DictLabel from '@/components/DictLabel/index.vue'
  import { useRoute, useRouter } from 'vue-router'

  const dictStore = useDictStore()

  defineOptions({ name: 'BizOrderDetail' })

  const route = useRoute()
  const router = useRouter()
  const loading = ref(false)
  const detail = ref<Record<string, any> | null>(null)

  const goBack = () => {
    router.back()
  }

  onMounted(async () => {
    const id = Number(route.params.id)
    if (!id) return
    loading.value = true
    try {
      detail.value = await fetchBizOrderView(id) as any
    } catch { detail.value = null }
    loading.value = false
  })
</script>

<style scoped>
  .detail-descriptions {
    :deep(.el-descriptions__label) {
      width: 140px;
      font-weight: 600;
    }
  }
</style>
