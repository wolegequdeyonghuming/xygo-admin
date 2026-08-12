<script>
	export default {
		onLaunch: function() {
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
		/* wot-design-uni 主题主色 */
		--wot-color-theme: #336FFF;
		--wot-color-theme-light: #336FFF;
		--wot-color-theme-dark: #2a5bd6;

		background-color: #f5f6f8;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
		font-size: 28rpx;
		color: #333333;
	}
</style>
