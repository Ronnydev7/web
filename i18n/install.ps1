function New-Directory {
    param (
        [string]$folder
    )

    # Check if the folder exists, and create it if not
    if (!(Test-Path -Path $folder)) {
        New-Item -ItemType Directory -Force -Path $folder
    }
}

function Remove-Directory {
    param (
        [string]$folder
    )

    # Check if the folder exists, and remove it if it does
    if (Test-Path -Path $folder) {
        Remove-Item -Recurse -Force -Path $folder
    }
}

Write-Host "++++++++++++++++++++"
Write-Host "installing i18n to app"

Write-Host "Installing intl..."
Remove-Directory -folder "..\app\src\generated\intl\"
New-Directory -folder "..\app\src\generated"
Copy-Item -Recurse -Force -Path "intl" -Destination "..\app\src\generated\intl\"

Write-Host "Installing translations..."
Remove-Directory -folder "..\app\src\generated\translations\"
New-Directory -folder "..\app\src\generated"
Copy-Item -Recurse -Force -Path "translations" -Destination "..\app\src\generated\translations\"
Write-Host "Successfully installed i18n to app"
Write-Host "---------------------"

Write-Host "++++++++++++++++++++"
Write-Host "installing i18n to api"
Set-Location -Path "..\api"
# You need to replace the following line with the equivalent command in PowerShell
go generate ./generators/i18n
Set-Location -Path "..\i18n"
Write-Host "Successfully installed i18n to api"
Write-Host "---------------------"
