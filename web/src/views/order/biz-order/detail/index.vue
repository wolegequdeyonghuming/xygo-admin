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
          <ElDescriptionsItem label="订单编号" :span="3">{{
            detail.orderNo ?? '-'
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
              <div v-if="attachmentList.length" class="attachment-list">
                <template v-for="(url, idx) in attachmentList" :key="idx">
                  <ElImage
                    v-if="isImage(url)"
                    :src="url"
                    :preview-src-list="imageList"
                    :initial-index="imageList.indexOf(url)"
                    fit="cover"
                    preview-teleported
                    class="attachment-thumb"
                  />
                  <a v-else :href="url" target="_blank" class="attachment-file">
                    <ArtSvgIcon :icon="getFileTypeIcon(url)" class="text-xl" />
                    <span class="attachment-ext">{{ getExt(url) }}</span>
                  </a>
                </template>
              </div>
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
            <ElDescriptionsItem label="备注" :span="3">{{
              detail.remark ?? '-'
            }}</ElDescriptionsItem>
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

  /** 附件 URL 列表（兼容逗号分隔的多个附件） */
  const attachmentList = computed<string[]>(() => {
    const raw = detail.value?.attachmentId
    if (!raw) return []
    return String(raw)
      .split(',')
      .map((s) => s.trim())
      .filter(Boolean)
  })

  /** 图片类型的附件 URL 列表（供 ElImage 预览） */
  const imageList = computed<string[]>(() => attachmentList.value.filter((url) => isImage(url)))

  /** 判断是否为图片（与上传预览组件逻辑一致） */
  const isImage = (url: string) => /\.(jpg|jpeg|png|gif|webp|bmp|svg)(\?.*)?$/i.test(url)

  /** 获取文件扩展名 */
  const getExt = (url: string) => url.split('.').pop()?.split('?')[0]?.toUpperCase() || 'FILE'

  /** 根据扩展名返回文件类型图标（与上传预览组件逻辑一致） */
  const getFileTypeIcon = (url: string) => {
    if (isImage(url)) return 'ri:image-line'
    const ext = getExt(url).toLowerCase()
    if (['pdf'].includes(ext)) return 'ri:file-pdf-2-line'
    if (['doc', 'docx'].includes(ext)) return 'ri:file-word-line'
    if (['xls', 'xlsx'].includes(ext)) return 'ri:file-excel-line'
    if (['zip', 'rar', '7z'].includes(ext)) return 'ri:file-zip-line'
    if (['mp4', 'avi', 'mov'].includes(ext)) return 'ri:video-line'
    if (['mp3', 'wav', 'flac'].includes(ext)) return 'ri:music-line'
    return 'ri:file-line'
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
  .attachment-list {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
  }
  .attachment-thumb {
    width: 80px;
    height: 80px;
    border-radius: 6px;
    border: 1px solid var(--el-border-color-lighter);
    display: block;
    cursor: pointer;
  }
  .attachment-file {
    width: 80px;
    height: 80px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    border-radius: 6px;
    border: 1px solid var(--el-border-color-lighter);
    background: var(--el-fill-color-lighter);
    color: var(--el-text-color-secondary);
    cursor: pointer;
    text-decoration: none;
  }
  .attachment-ext {
    font-size: 10px;
    margin-top: 2px;
  }
</style>
