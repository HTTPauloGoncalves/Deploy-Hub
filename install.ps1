$ErrorActionPreference = "Stop"

$InstallDir = "$env:USERPROFILE\go\bin"
$ConfigDir = "$env:APPDATA\deployhub"
$Url = "https://github.com/HTTPauloGoncalves/Deploy-Hub/releases/download/v0.1.1/deployhub-windows-amd64.exe"

New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
New-Item -ItemType Directory -Force -Path $ConfigDir | Out-Null
Invoke-WebRequest -Uri $Url -OutFile "$InstallDir\deployhub.exe"

$userPath = [Environment]::GetEnvironmentVariable("Path", "User")

if ($userPath -notlike "*$InstallDir*") {
    [Environment]::SetEnvironmentVariable("Path", "$InstallDir;$userPath", "User")
    Write-Host "Feche e abra o terminal para atualizar o PATH."
}

Write-Host "DeployHub instalado!"
Write-Host "Configuracao padrao em: $ConfigDir\deploy.yaml"
Write-Host "Teste com: deployhub --help"
