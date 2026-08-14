<script>
import { initSystem } from '@/utils/system'
export default {
	onLaunch: function() {
		initSystem()
		console.log('App Launch')
	},
	onShow: function() {
		console.log('App Show')
	},
	onHide: function() {
		console.log('App Hide')
	}
}
</script>

<style lang="scss">
	page {
		/* wot-design-uni 主题主色（钴蓝：主操作） */
		--wot-color-theme: #2563EB;
		--wot-color-theme-light: #3B82F6;
		--wot-color-theme-dark: #1D4ED8;

		/* Cobalt Ops · 清爽白 + 钴蓝 令牌 */
		--bg: #F4F6FA;
		--card: #FFFFFF;
		--cobalt: #2563EB;
		--cobalt-dark: #1D4ED8;
		--cobalt-soft: #E8EFFD;

		--ink: #111827;
		--text-2: #4B5563;
		--text-3: #9CA3AF;

		--status-gray: #8A8F98;
		--status-amber: #E8930C;
		--status-red: #E5484D;
		--status-green: #30A46C;

		--card-shadow: 0 6rpx 24rpx rgba(17, 24, 39, 0.06);
		--card-shadow-lg: 0 16rpx 40rpx rgba(37, 99, 235, 0.12);
		--hairline: 1rpx solid rgba(17, 24, 39, 0.06);

		background-color: #F4F6FA;
		font-family: 'PingFang SC', 'HarmonyOS Sans SC', 'MiSans', 'Helvetica Neue', Arial, sans-serif;
		font-size: 28rpx;
		color: #111827;
	}
</style>
