$InstallDir = "$env:USERPROFILE\go\bin"
$Url = "https://github.com/HTTPauloGoncalves/Deploy-Hub/releases/latest/download/deployhub-windows-amd64.exe"

New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
Invoke-WebRequest -Uri $Url -OutFile "$InstallDir\deployhub.exe"

$userPath = [Environment]::GetEnvironmentVariable("Path", "User")

if ($userPath -notlike "*$InstallDir*") {
    [Environment]::SetEnvironmentVariable("Path", "$InstallDir;$userPath", "User")
    Write-Host "Feche e abra o terminal para atualizar o PATH."
}

Write-Host "DeployHub instalado!"