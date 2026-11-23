# ============================================================
# 构建 daily 二进制（linux/amd64 静态 ELF），产物：deploy\daily
# 用法：powershell -File deploy\build-daily.ps1
# 每次改后端代码后运行本脚本，并提交 deploy\daily
# 注意：必须在同一会话设置 GOOS/GOARCH 后再 go build，否则 Windows 默认产出 PE（MZ）
# ============================================================
$ErrorActionPreference = 'Stop'

$root = Split-Path -Parent $PSScriptRoot
$serverDir = Join-Path $root 'server'
$out = Join-Path $root 'deploy\fpd'

$env:CGO_ENABLED = '0'
$env:GOOS = 'linux'
$env:GOARCH = 'amd64'

Push-Location $serverDir
try {
    go build -o $out .
} finally {
    Pop-Location
}

# 核验魔数：应为 7F 45 4C 46（ELF），若是 4D 5A（MZ）则失败
$magic = [System.IO.File]::ReadAllBytes($out)[0..3] | ForEach-Object { $_.ToString('X2') }
$magicStr = ($magic -join ' ')
if ($magicStr -ne '7F 45 4C 46') {
    Write-Host "ERROR: 产物非 ELF（magic=$magicStr），可能未正确设置 GOOS=linux"
    exit 1
}
$size = (Get-Item $out).Length
Write-Host "OK: $out (ELF, $([math]::Round($size / 1MB, 1)) MB)"
