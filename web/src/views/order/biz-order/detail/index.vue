<!-- 订单主表 详情页（全屏，按步骤分区展示） -->
<template>
  <div class="biz-order-detail art-full-height">
    <ElCard shadow="never" class="art-table-card">
      <template #header>
        <div class="flex items-center justify-between">
          <span class="font-bold text-lg">订单详情</span>
          <ElButton @click="goBack">
            <ArtSvgIcon icon="ri:arrow-left-line" class="text-sm mr-1" />
            返回列表
          </ElButton>
        </div>
      </template>

      <div v-if="loading" class="flex justify-center py-20">
        <ElIcon class="is-loading" :size="32"><Loading /></ElIcon>
      </div>

      <div v-else-if="detail" style="height: 100%; overflow-y: auto">
        <!-- 顶部概览 -->
        <ElDescriptions :column="5" border class="detail-descriptions overview-descriptions">
          <ElDescriptionsItem label="订单状态" :span="1">
            <DictLabel :dict-type="'order_status'" :value="detail.orderStatus" />
          </ElDescriptionsItem>
          <ElDescriptionsItem label="话务员" :span="2">{{
            detail.telemarketer_real_name ?? '-'
          }}</ElDescriptionsItem>
          <ElDescriptionsItem label="收单员" :span="2">{{
            detail.agent_real_name ?? '-'
          }}</ElDescriptionsItem>
          <ElDescriptionsItem label="创建时间" :span="3">{{
            formatTimestamp(detail.createdAt)
          }}</ElDescriptionsItem>
          <ElDescriptionsItem label="更新时间" :span="2">{{
            formatTimestamp(detail.updatedAt)
          }}</ElDescriptionsItem>
        </ElDescriptions>

        <!-- 1 录单信息 -->
        <ElDescriptions :column="2" border class="detail-descriptions" title="录单信息">
          <ElDescriptionsItem label="排单日期">{{
            formatTimestamp(detail.scheduleDate, 'date')
          }}</ElDescriptionsItem>
          <ElDescriptionsItem label="客户姓名">{{ detail.customerName ?? '-' }}</ElDescriptionsItem>
          <ElDescriptionsItem label="可联系时间">{{
            detail.availableTimeDesc ?? '-'
          }}</ElDescriptionsItem>
          <ElDescriptionsItem label="联系电话">{{ detail.contactPhone ?? '-' }}</ElDescriptionsItem>
          <ElDescriptionsItem label="区县">
            <DictLabel :dict-type="'area'" :value="detail.area" />
          </ElDescriptionsItem>
          <ElDescriptionsItem label="安装地址">{{
            detail.installAddress ?? '-'
          }}</ElDescriptionsItem>
          <ElDescriptionsItem label="预约业务">{{ detail.businessType ?? '-' }}</ElDescriptionsItem>
          <ElDescriptionsItem label="到期时间">{{ detail.expiryDate ?? '-' }}</ElDescriptionsItem>
        </ElDescriptions>

        <!-- 3 预约信息 -->
        <StepSection title="预约信息">
          <ElDescriptions :column="2" border class="detail-descriptions">
            <ElDescriptionsItem label="收单预约情况">{{
              detail.appointmentDesc ?? '-'
            }}</ElDescriptionsItem>
          </ElDescriptions>
        </StepSection>

        <!-- 4 收单信息 -->
        <StepSection title="收单信息">
          <ElDescriptions :column="2" border class="detail-descriptions">
            <ElDescriptionsItem label="上门日期">{{
              formatTimestamp(detail.visitDate)
            }}</ElDescriptionsItem>
            <ElDescriptionsItem label="成交业务">{{
              detail.dealtBusinessType ?? '-'
            }}</ElDescriptionsItem>
            <ElDescriptionsItem label="携转情况">{{
              detail.portingStatus ?? '-'
            }}</ElDescriptionsItem>
            <ElDescriptionsItem label="客户实际姓名">{{
              detail.customerRealName ?? '-'
            }}</ElDescriptionsItem>
            <ElDescriptionsItem label="客户身份证号">{{
              detail.customerIdNumber ?? '-'
            }}</ElDescriptionsItem>
            <ElDescriptionsItem label="实缴额度">{{ detail.paidAmount ?? '-' }}</ElDescriptionsItem>
            <ElDescriptionsItem label="是否乡下单">
              <ElTag size="small">{{
                { '0': '否', '1': '是' }[String(detail.isRuralOrder)] || detail.isRuralOrder
              }}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem label="主卡号码">{{
              detail.primaryTelNo ?? '-'
            }}</ElDescriptionsItem>
            <ElDescriptionsItem label="新开号码">{{ detail.newPhoneNo ?? '-' }}</ElDescriptionsItem>
            <ElDescriptionsItem label="终端串码">{{
              detail.deviceSerial ?? '-'
            }}</ElDescriptionsItem>
            <ElDescriptionsItem label="话补">{{ detail.subsidyAmount ?? '-' }}</ElDescriptionsItem>
            <ElDescriptionsItem label="是否纯新增">{{ detail.isNew ?? '-' }}</ElDescriptionsItem>
            <ElDescriptionsItem label="附件">
              <a
                v-if="detail.attachmentId"
                :href="detail.attachmentId"
                target="_blank"
                style="color: var(--el-color-primary)"
                >{{ detail.attachmentId }}</a
              >
              <span v-else>-</span>
            </ElDescriptionsItem>
          </ElDescriptions>
        </StepSection>

        <!-- 5 回访信息 -->
        <StepSection title="回访信息">
          <ElDescriptions :column="2" border class="detail-descriptions">
            <ElDescriptionsItem label="话务二次回访情况">{{
              detail.followUpDesc ?? '-'
            }}</ElDescriptionsItem>
          </ElDescriptions>
        </StepSection>

        <!-- 6 完成信息 -->
        <StepSection title="完成信息">
          <ElDescriptions :column="3" border class="detail-descriptions">
            <ElDescriptionsItem label="宽带账号">{{
              detail.broadbandAccount ?? '-'
            }}</ElDescriptionsItem>
            <ElDescriptionsItem label="工号">{{ detail.agencyNo ?? '-' }}</ElDescriptionsItem>
            <ElDescriptionsItem label="是否完工">
              <ElTag size="small">{{
                { '0': '未完工', '1': '已完工' }[String(detail.isCompleted)] || detail.isCompleted
              }}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem label="备注" :span="3">{{ detail.remark ?? '-' }}</ElDescriptionsItem>
          </ElDescriptions>
        </StepSection>

        <!-- 详细情况（评论区） -->
        <ElCard shadow="never" class="comment-card">
          <OrderComment :order-id="Number(detail.id)" />
        </ElCard>
      </div>

      <div v-else class="py-20 text-center text-gray-400"> 数据加载失败 </div>
    </ElCard>
  </div>
</template>

<script setup lang="ts">
  import { Loading } from '@element-plus/icons-vue'
  import ArtSvgIcon from '@/components/core/base/art-svg-icon/index.vue'
  import { fetchBizOrderView } from '@/api/order/biz-order'
  import { formatTimestamp } from '@/utils/time'
  import DictLabel from '@/components/DictLabel/index.vue'
  import OrderComment from '../components/order-comment.vue'
  import StepSection from './step-section.vue'
  import { useRoute, useRouter } from 'vue-router'

  defineOptions({ name: 'BizOrderDetail' })

  const route = useRoute()
  const router = useRouter()
  const loading = ref(false)
  const detail = ref<Record<string, any> | null>(null)

  const goBack = () => {
    router.back()
  }

  onMounted(async () => {
    const id = Number(route.query.id || route.params.id)
    if (!id) return
    loading.value = true
    try {
      detail.value = (await fetchBizOrderView(id)) as any
    } catch {
      detail.value = null
    }
    loading.value = false
  })
</script>

<style scoped>
    :deep(.el-descriptions__label) {
      width: 140px;
      font-weight: 600;
    }
  .overview-descriptions {
    margin-bottom: 16px;
  }
  .comment-card {
    margin-top: 16px;
    :deep(.el-card__body) {
      padding: 16px;
    }
  }
</style>
